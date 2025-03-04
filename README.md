# GoExpert - From Brazilian CEP to Current Temperature

![test](https://github.com/lmtani/learning-current-city-weather/actions/workflows/test.yml/badge.svg)

Sistema em Go que recebe um CEP, identifica a cidade e retorna o clima atual temperatura em graus celsius (temp_C), fahrenheit (temp_F) e kelvin (temp_K).

**Demo:** https://cep2weather-425952763790.us-central1.run.app?cep=13280001

## Quickstart

**Pré-requisitos:** docker e docker compose

- Crie um arquivo .env dentro do diretório config/ com sua credencial WEATHER_API_KEY:

   ```bash
   echo WEATHER_API_KEY=<WEATHER_API_KEY> > config/.env
   ```

- Inicie o sistema com `docker compose up`

## Testes

- Testes unitários para o caso de uso de busca de clima por CEP.
- HTTP test para o endpoint de busca de clima por CEP.
  - *success.http*: Teste de sucesso com um CEP valido (200)
  - *invalid.http*: Teste com um CEP invalido (422)
  - *not-found.http*: Teste com um CEP não encontrado (404)
