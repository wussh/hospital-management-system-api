package util

import (
	"log"
	"os"
	"path/filepath"
)

func CreateStorageDirectory() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, "storage_data", "avatar")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.MkdirAll(path, os.ModeAppend)
	}
}
