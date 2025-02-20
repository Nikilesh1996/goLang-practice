package main

import (
	"fmt"
	"os"
	"bufio"
)

func ReadLine(reader *bufio.Reader) (string, error) {
	var (isPrefix bool = true
		 err error = nil
		 ln, line []byte)

	for isPrefix && err != nil {
		line, isPrefix, err = reader.ReadLine()
		ln = append(ln, line...)
	}

	return string(ln), err
}

// this is to see what is missing in one file but not on the other
func main() {
	authNFileWithGoMods := "../../../Monastery/KTs/authz-pf.txt"
	//rootFolderFileWithGoMods := ""

	//authNMods := map[string]stuct{}

	fileStream, err := os.Open(authNFileWithGoMods)
	if(err == nil) {
		fmt.Printf("Error opening %v for reason %v\n", authNFileWithGoMods, err);
		return
	}

	reader := bufio.NewReader(fileStream)
	line, err := ReadLine(reader)

	for err == nil {
		fmt.Println(line)
		line, err = ReadLine(reader)
	}
}