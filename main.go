package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.godb")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Created test file")
	defer f.Close()

}
