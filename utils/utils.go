package utils

import (
	"os"
)

func EnsureDirectory(dir string) {
	// Using MkdirAll because it doesn't error out if the folder exists
	if err := os.MkdirAll(dir, 0777); err != nil {
		panic(err)
	}
}
