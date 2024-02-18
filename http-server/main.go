package main

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct {
	Status string `json:"status"`
}

func main() {

	http.HandleFunc("/", HealthCheckController)
	http.ListenAndServe(":8080", nil)
}

func HealthCheckController(w http.ResponseWriter, r *http.Request) {
	status := HealthCheck{Status: "UP"}
	res, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	w.Write([]byte(res))
}
