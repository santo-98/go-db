package main

import (
	"fmt"
	"os"

	"github.com/santo-98/godb/utils/fileoperations"
)

func main() {
	f, err := os.Create("test.godb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Created test file")

	_, writeErr := f.WriteString("testing")
	if writeErr != nil {
		fmt.Println(writeErr)
	}

	fmt.Println(fileoperations.ReadEncryptedFile)
	defer f.Close()
}
