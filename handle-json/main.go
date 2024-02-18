package main

import (
	"encoding/json"
	"os"
)

type Saldo struct {
	Number int `json:"n"`
	Amount int `json:"s"`
}

func main() {

	saldo := Saldo{Number: 1, Amount: 1000}
	res, err := json.Marshal(saldo)

	if err != nil {
		panic(err)
	}

	println(string(res))

	json.NewEncoder(os.Stdout).Encode(saldo)

	conta2 := []byte(`{"n":2,"s":2000}`)

	var saldoInStruc Saldo
	err = json.Unmarshal(conta2, &saldoInStruc)

	if err != nil {
		panic(err)
	}

	println(saldoInStruc.Amount)

}
