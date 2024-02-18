package main

import (
	"io"
	"net/http"
)

func main() {
	c := http.Client{}

	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("User-Agent", "Gustavo Henrique")
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	println(string(body))
}
