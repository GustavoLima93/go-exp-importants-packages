package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {

	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cep := r.URL.Path[1:]
	if len(cep) != 8 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	viaCep, err := BuscaCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	//res, err := json.Marshal(viaCep)
	//if err != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//
	//}
	//
	//w.Write([]byte(res))

	json.NewEncoder(w).Encode(viaCep)

}

func BuscaCep(cep string) (*ViaCep, error) {
	req, err := http.Get("https://viacep.com.br/ws/" + cep + "/json")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var viaCep ViaCep
	err = json.Unmarshal(res, &viaCep)
	if err != nil {
		return nil, err
	}
	return &viaCep, nil

}
