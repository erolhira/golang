package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/**
https://pkg.go.dev/net/http
 */

func makeHttpCall() {

	resp, err := http.Get("https://www.google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Println(resp)

	//1.read from Reader interface of Body
	bs := make([]byte, 99999) //create an empty byte slice having an initial capacity of 99999 bytes
	resp.Body.Read(bs)
	fmt.Println(string(bs))

	//2. alternative to 1
	io.Copy(os.Stdout, resp.Body)
}