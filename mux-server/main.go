package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", heatlhCheck)
	mux.Handle("/blog", &blog{title: "My Blog"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/health", heatlhCheck)
	mux2.Handle("/blog", &blog{title: "My Blog"})
	http.ListenAndServe(":8081", mux2)

}

func heatlhCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("UP"))
}

type blog struct {
	title string
}

func (b *blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
