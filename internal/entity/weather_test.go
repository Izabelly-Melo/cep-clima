package entity

import "testing"

func TestNewWeather(t *testing.T) {
	w := NewWeather(28.5)

	if w.TempC != 28.5 {
		t.Errorf("esperado TempC 28.5, obtido %v", w.TempC)
	}
	if w.TempF != 83.3 {
		t.Errorf("esperado TempF 83.3, obtido %v", w.TempF)
	}
	if w.TempK != 301.5 {
		t.Errorf("esperado TempK 301.5, obtido %v", w.TempK)
	}
}

func TestNewWeatherZero(t *testing.T) {
	w := NewWeather(0)

	if w.TempF != 32 {
		t.Errorf("esperado TempF 32, obtido %v", w.TempF)
	}
	if w.TempK != 273 {
		t.Errorf("esperado TempK 273, obtido %v", w.TempK)
	}
}
