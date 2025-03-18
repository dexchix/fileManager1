package ui

import (
	"fileManager/models"
	"fmt"
	"math"
)

type FileView struct {
	Name       string
	Info       string
	UpdateTime string
}

func ConvertToFileView(file models.File) FileView {
	result := FileView{}
	result.Name = file.Name

	if file.IsDirectory {
		result.Info = "<DIR>"
	} else {
		result.Info = formatFileSize(file.Size)
	}
	result.UpdateTime = file.LastUpdate.Format("2006-01-02 15:04")
	return result
}

func ConvertToFileViewArr(files []models.File) []FileView {
	result := []FileView{}

	for _, f := range files {
		result = append(result, ConvertToFileView(f))
	}
	return result
}

func formatFileSize(size int64) string {
	if size == 0 {
		return "0 B"
	}
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	order := int(math.Log2(float64(size)) / 10)
	if order >= len(units) {
		order = len(units) - 1
	}
	formattedSize := float64(size) / math.Pow(1024, float64(order))
	return fmt.Sprintf("%.2f %s", formattedSize, units[order])
}
