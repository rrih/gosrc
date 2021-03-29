package main

import (
	"io"
	"os"
	"strings"
)

// see: https://ascii.jp/elem/000/001/260/1260449/
func main() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}