package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"weather-api/internal/dto"
	"weather-api/internal/entity"
)

func BrasilCepHandler(cep string) (*entity.CEP, error) {
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data dto.BrasilAPIOutput
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	responseCep := entity.NewCEP("Brasil CEP", data.Cep, data.Street, data.Neighborhood, data.City, data.State)

	return responseCep, nil
}
