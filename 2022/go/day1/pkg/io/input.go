package io

import (
	"fmt"
	"os"
)

func ReadInputBytes(path string) []byte {
	if path == "" {
		path = "input.txt"
	}
	contents, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("file read failed: ", err)
		os.Exit(1)
	}
	return contents
}
