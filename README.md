# GoExpert - CEP2Weather

![test](https://github.com/lmtani/learning-current-city-weather/actions/workflows/test.yml/badge.svg)
![deploy](https://github.com/lmtani/learning-current-city-weather/actions/workflows/deploy.yml/badge.svg)

**CEP2Weather** é um sistema em Go que recebe um CEP, identifica a cidade correspondente e retorna a temperatura atual em graus Celsius, Fahrenheit e Kelvin.

## Demonstração

Acesse em produção (Google Cloud Run):  
**[CEP2Weather Demo](https://cep2weather-425952763790.us-central1.run.app?cep=13280001)**

Basta informar o CEP na query string, por exemplo:

```http
https://cep2weather-425952763790.us-central1.run.app?cep=13280001
```

## Funcionalidades

- **Busca de CEP**: Utiliza a API [BrasilAPI](https://brasilapi.com.br/) para converter um CEP válido em uma cidade.
- **Busca de clima**: Utiliza a API [WeatherAPI](https://www.weatherapi.com/) para obter a temperatura na cidade.
- **Conversões de temperatura**:  
  - Celsius para Fahrenheit: `F = C * 1.8 + 32`  
  - Celsius para Kelvin: `K = C + 273`
- **Tratamento de erros**:
  - CEP inválido (8 dígitos incorretos): retorna código HTTP **422** e mensagem `"invalid zipcode"`.
  - CEP não encontrado: retorna código HTTP **404** e mensagem `"can not find zipcode"`.
  - Timeout após 3 tentativas: empiricamente foi verificado que é um erro comum quando o CEP não é encontrado. Por isso retorna código HTTP **404** e mensagem `"can not find zipcode"`.

## Requisitos

- **CEP**: Obrigatoriamente 8 dígitos.
- **Formato de resposta em caso de sucesso** (HTTP 200):
  
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.5
  }
  ```

## Pré-requisitos

- **Docker** e **Docker Compose** instalados em seu ambiente.
- **Chave de API** da [WeatherAPI](https://www.weatherapi.com/).

## Configuração

1. **Clonar o repositório**:

   ```bash
   git clone https://github.com/lmtani/learning-current-city-weather.git
   cd learning-current-city-weather
   ```

2. **Criar arquivo `.env`** com a variável de ambiente de API Key no diretório `config/`:

   ```bash
   echo WEATHER_API_KEY=<SUA_CHAVE> > config/.env
   ```

3. **Iniciar o sistema** com Docker Compose:

   ```bash
   docker compose up
   ```

4. **Acessar a aplicação** (por padrão, via porta 8080):

   ```http
   http://localhost:8080/?cep=13280001
   ```

## Testes

- **Testes unitários**: Validam o caso de uso de busca de clima por CEP.
- **HTTP tests**: Arquivos `.http` para validação de chamadas REST:
  - `success.http`: Teste de sucesso com um CEP válido (expectativa de código **200**).
  - `invalid.http`: Teste com CEP inválido (expectativa de código **422**).
  - `not-found.http`: Teste com CEP não encontrado (expectativa de código **404**).

## Deploy

Este projeto está configurado para deploy automático no **Google Cloud Run** por meio de GitHub Actions. Para visualizar a configuração, consulte o arquivo `deploy.yml` no diretório `.github/workflows/`.

1. **Pipeline de CI/CD**: A cada push na branch `main`, o teste é executado e, em seguida, o deploy é realizado no Cloud Run.
2. **Imagem Docker**: A imagem é construída e enviada para o [Google Artifact Registry](https://cloud.google.com/artifact-registry), e então é feita a implantação no Cloud Run.
3. **Novas revisões**: A cada novo deploy, uma nova revisão do serviço é criada automaticamente.
