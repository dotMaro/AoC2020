package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// OpenFile wraps os.Open.
// Panics on error.
func OpenFile(name string) *os.File {
	pwd, _ := os.Getwd()
	path := filepath.Join(pwd, name)
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	return file
}

// ReadFile returns a file's content as a string.
func ReadFile(name string) string {
	file := OpenFile(name)
	defer file.Close()
	inputBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	return string(inputBytes)
}

// Print to console.
func Print(format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
}

// SplitLine splits s into a slice of lines.
// It supports both with carriage return and without.
func SplitLine(s string) []string {
	return strings.Split(strings.ReplaceAll(s, "\r\n", "\n"), "\n")
}
