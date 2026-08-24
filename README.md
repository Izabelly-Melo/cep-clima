# Clima por CEP — Go Expert

Sistema em Go que recebe um CEP, identifica a cidade correspondente (ViaCEP) e retorna a temperatura atual em Celsius, Fahrenheit e Kelvin (WeatherAPI).

## URL em produção (Cloud Run)

> https://cep-clima-337022985738.southamerica-east1.run.app

Exemplo de uso:

```bash
curl https://cep-clima-337022985738.southamerica-east1.run.app/01310100
```

## Tecnologias

- Go 1.25
- `net/http` (sem frameworks externos)
- Docker (multi-stage build)

## APIs consumidas

- ViaCEP: `http://viacep.com.br/ws/{cep}/json/` — identifica a cidade a partir do CEP
- WeatherAPI: `https://api.weatherapi.com/v1/current.json` — retorna a temperatura atual da cidade

## Contrato da API

### Sucesso — `GET /{cep}` — 200 OK

```json
{
  "temp_C": 28.5,
  "temp_F": 83.3,
  "temp_K": 301.5
}
```

### Falhas

| Cenário | Condição | Status | Corpo |
|---|---|---|---|
| Formato inválido | CEP sem 8 dígitos ou com caracteres inválidos | 422 | `{"message": "invalid zipcode"}` |
| CEP não encontrado | CEP com formato correto, mas inexistente | 404 | `{"message": "can not find zipcode"}` |

## Variáveis de ambiente

| Variável | Descrição | Obrigatória |
|---|---|---|
| `WEATHER_API_KEY` | Chave de API do [WeatherAPI](https://www.weatherapi.com/) | Sim |
| `PORT` | Porta HTTP do servidor (Cloud Run injeta automaticamente) | Não (padrão `8080`) |

## Como rodar localmente

Pré-requisito: Go instalado (versão 1.25 ou superior).

```bash
git clone <url-do-repositorio>
cd <pasta-do-projeto>
export WEATHER_API_KEY=sua_chave_aqui
go run cmd/main.go
```

O servidor sobe na porta `8080`.

```bash
curl http://localhost:8080/01310100
```

## Como rodar com Docker

```bash
docker build -t clima-cep .
docker run -p 8080:8080 -e WEATHER_API_KEY=sua_chave_aqui clima-cep
```

```bash
curl http://localhost:8080/01310100
```

## Como rodar os testes

```bash
go test ./... -v
```

Os testes cobrem:
- Conversão de temperatura (Celsius → Fahrenheit e Celsius → Kelvin)
- Cenário de sucesso (200)
- CEP com formato inválido (422)
- CEP não encontrado (404)

As chamadas às APIs externas são substituídas por servidores HTTP de teste (`httptest`), então os testes não dependem de rede.

## Deploy no Google Cloud Run

```bash
gcloud builds submit --tag gcr.io/SEU_PROJETO/clima-cep

gcloud run deploy clima-cep \
  --image gcr.io/SEU_PROJETO/clima-cep \
  --platform managed \
  --region us-central1 \
  --allow-unauthenticated \
  --set-env-vars WEATHER_API_KEY=sua_chave_aqui
```

Após o deploy, copie a URL gerada pelo Cloud Run e atualize a seção **URL em produção** deste README.
