package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"weather-api/internal/dto"
	"weather-api/internal/entity"
)

var WeatherAPIBaseURL = "https://api.weatherapi.com/v1/current.json"

func WeatherHandler(city string) (*entity.Weather, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")

	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", WeatherAPIBaseURL, apiKey, url.QueryEscape(city))
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	if req.StatusCode == http.StatusNotFound {
		return nil, ErrNotFound
	}

	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var data dto.WeatherAPIOutput
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return entity.NewWeather(data.Current.TempC), nil
}
