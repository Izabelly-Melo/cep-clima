package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather-api/internal/dto"
	"weather-api/internal/entity"
)

func ViaCepHandler(cep string) (*entity.CEP, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data dto.ViaCEPOutput
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	if data.Erro {
		return nil, ErrNotFound
	}

	responseCep := entity.NewCEP("Via CEP", data.Cep, data.Logradouro, data.Bairro, data.Localidade, data.UF)

	return responseCep, nil
}
