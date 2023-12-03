package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	bs := make([]byte, 99999)

	//Passing a slice of bytes to the Read function,
	//and let the read function fill up the slice with the data it reads
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
