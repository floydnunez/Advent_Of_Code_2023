package util

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileIntoArray(filename string) ([]string, bool) {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	errclose := readFile.Close()
	return fileLines, errclose != nil
}
