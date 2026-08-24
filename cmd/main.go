package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"regexp"
	"weather-api/internal/api"

	"github.com/joho/godotenv"
)

var cepRegex = regexp.MustCompile(`^\d{8}$`)

type ErrorResponse struct {
	Message string `json:"message"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: arquivo .env não encontrado, seguindo com variáveis de ambiente do sistema")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{cep}", weatherHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.ListenAndServe(":"+port, mux)
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	cep := r.PathValue("cep")
	w.Header().Set("Content-Type", "application/json")

	if !cepRegex.MatchString(cep) {
		writeError(w, http.StatusUnprocessableEntity, "invalid zipcode")
		return
	}

	address, err := api.ViaCepHandler(cep)
	if err != nil {
		if errors.Is(err, api.ErrNotFound) {
			writeError(w, http.StatusNotFound, "can not find zipcode")
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	weather, err := api.WeatherHandler(address.Cidade)
	if err != nil {
		if errors.Is(err, api.ErrNotFound) {
			writeError(w, http.StatusNotFound, "can not find zipcode")
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(weather)
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}
