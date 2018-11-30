/*
Create a program that reads the contents of a text file then prints its contents to the terminal.
e.g. 'go run print_file.go ../../README.md'
 */

package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fileName := os.Args[1]; fmt.Println(fileName)
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}
