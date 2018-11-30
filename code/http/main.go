package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

/*	//bs := []byte{}  //zero length, Reader func can't store data in it
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	//fmt.Println(string(bs))
	fmt.Printf("%s\n", string(bs))  //both print trailing null. Why?*/

	//io.Copy(os.Stdout, resp.Body)

	lw := logWriter{}

	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
