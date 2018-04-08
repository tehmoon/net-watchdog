package main

import (
	"os"
)

func openRegularFile(name string) (*os.File, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0400)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		file.Close()
		return nil, err
	}

	if ! stat.Mode().IsRegular() {
		file.Close()
		return nil, ErrNotRegularFile
	}

	return file, nil
}
