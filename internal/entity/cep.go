package entity

import (
	"strings"
)

type CEP struct {
	API        string
	Cep        string
	Logradouro string
	Bairro     string
	Cidade     string
	UF         string
}

func NewCEP(api, cep, logradouro, bairro, cidade, uf string) *CEP {
	return &CEP{
		API:        api,
		Cep:        cepFormatador(cep),
		Logradouro: logradouro,
		Bairro:     bairro,
		Cidade:     cidade,
		UF:         uf,
	}
}

func cepFormatador(cep string) string {
	cep = strings.TrimSpace(cep)
	cep = strings.ReplaceAll(cep, "-", "")

	if len(cep) != 8 {
		return cep
	}

	return cep[:5] + "-" + cep[5:]
}
