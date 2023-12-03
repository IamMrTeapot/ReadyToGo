package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 99999)

	// //Passing a slice of bytes to the Read function,
	// //and let the read function fill up the slice with the data it reads
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	//io.Copy takes a writer and a reader
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
