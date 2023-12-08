package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileIntoArray(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		panic("could not read the file")
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	errclose := readFile.Close()
	if errclose != nil {
		panic("could not close the file")
	}
	return fileLines
}
