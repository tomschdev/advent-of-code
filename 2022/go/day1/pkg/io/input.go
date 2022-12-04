package io

import (
	"fmt"
	"os"
	"strings"
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

func BytesToTrimmedString(contents []byte, cut string) string {
	return strings.Trim(string(contents), cut)
}
