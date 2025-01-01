package main

import (
	"fmt"
	"os"
)

func readMarkDown(fileName string) string {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(file)
}

func main() {
	fmt.Println(readMarkDown("md/test.md"))
}
