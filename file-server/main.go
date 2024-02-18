package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	muxServer := http.NewServeMux()
	muxServer.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8080", muxServer))
}
