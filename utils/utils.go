package utils

import (
	"os"
)

func EnsureDirectory(dir string) {
	// Using MkdirAll because it doesn't error out if the folder exists
	err := os.MkdirAll(dir, 0777)
	ErrPanic(err)
}

func ErrPanic(e error) {
	if e != nil {
		panic(e)
	}
}

func ErrPanicB(e error, b bool) {
	if b && e != nil {
		panic(e)
	}
}
