package main

import (
	"fmt"
	"io"
	"os"
)

func checkNilErr(err error) {
	e
}

func main() {
	fmt.Println("File handling in Go!")

	filePath := "./GoFile.txt"
	fileContent := "Hello ing the GO World!\n This is a file handling in Go.\n"

	createdFile, createFileErr := os.Create(filePath)
	if createFileErr != nil {
		panic(createFileErr)
	}

	fileLength, writeErr := io.WriteString(createdFile, fileContent)
	if writeErr != nil {
		panic(writeErr)
	}

	fmt.Println("File created and written successfully with length", fileLength)
	defer createdFile.Close()

	// reading the file
	readFile(filePath)
}

func readFile(filePath string) {

	dataInByte, fileReadErr := os.ReadFile(filePath)
	if fileReadErr != nil {
		panic(fileReadErr)
	}

	fmt.Println("File content is: \n", string(dataInByte))

}
