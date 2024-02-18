package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	req, err := http.Get("https://www.google.com")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(req.Body)

	fmt.Println("Response: ", string(res))

	req.Body.Close()
}
