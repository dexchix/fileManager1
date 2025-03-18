package filesystem

import (
	"fileManager/models"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

const DefaultPath string = "C:\\"

func GetFiles(path string) (result []models.File, err error) {
	files, err := os.ReadDir(path)
	if err != nil {
		return result, fmt.Errorf("error: directory is not found or not accessible")
	}

	for _, r := range files {
		fileInfo, err := r.Info()
		if err != nil {
			log.Printf("error: file [%s] is not accessible", r.Name())
			continue
		}

		filePath := filepath.Join(path, fileInfo.Name())

		result = append(result, models.File{
			Name:        fileInfo.Name(),
			IsDirectory: fileInfo.IsDir(),
			LastUpdate:  fileInfo.ModTime(),
			Size:        fileInfo.Size(),
			Path:        filePath,
		})
	}

	return result, nil
}

func DeleteFile(file models.File) error {
	if file.IsDirectory {
		err := os.RemoveAll(file.Path)
		if err != nil {
			return fmt.Errorf("Не удалось удалить папку %s", file.Path)
		}
	} else {
		err := os.Remove(file.Path)
		if err != nil {
			return fmt.Errorf("Не удалось удалить файл %s", file.Path)
		}
	}
	return nil
}

func Copy(src, dstDir string) error {
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("не удалось получить информацию об исходном элементе: %w", err)
	}

	if srcInfo.IsDir() {
		return copyDir(src, dstDir, srcInfo)
	}

	return copyFile(src, dstDir, srcInfo)
}

func copyDir(src, dstDir string, srcInfo os.FileInfo) error {
	dirName := filepath.Base(src)

	dst := filepath.Join(dstDir, dirName)

	err := os.MkdirAll(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("не удалось создать целевую папку: %w", err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("не удалось прочитать содержимое папки: %w", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			err = Copy(srcPath, dstPath)
			if err != nil {
				return err
			}
		} else {
			err = copyFile(srcPath, dstPath, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func copyFile(src, dstDir string, srcInfo os.FileInfo) error {
	if srcInfo == nil {
		var err error
		srcInfo, err = os.Stat(src)
		if err != nil {
			return fmt.Errorf("не удалось получить информацию о файле: %w", err)
		}
	}

	fileName := filepath.Base(src)

	dst := filepath.Join(dstDir, fileName)

	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("не удалось открыть исходный файл: %w", err)
	}
	defer srcFile.Close()

	err = os.MkdirAll(dstDir, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("не удалось создать целевую директорию: %w", err)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("не удалось создать целевой файл: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("не удалось скопировать содержимое файла: %w", err)
	}

	err = os.Chmod(dst, srcInfo.Mode())
	if err != nil {
		return fmt.Errorf("не удалось установить права доступа для целевого файла: %w", err)
	}

	return nil
}
