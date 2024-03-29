package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonvar := bytes.NewBuffer([]byte(`{"name":"Gustavo Henrique"}`))
	resp, err := c.Post("http://httpbin.org/post", "application/json", jsonvar)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.CopyBuffer(os.Stdout, resp.Body, nil)
}
