package models

import (
	"time"
)

type File struct {
	Name        string
	Path        string
	IsDirectory bool
	LastUpdate  time.Time
	Size        int64
}
