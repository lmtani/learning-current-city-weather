# GoExpert - CEP2Weather

**CEP2Weather** é um sistema para estudos em Go que recebe um CEP, identifica a cidade correspondente e retorna a temperatura atual em graus Celsius, Fahrenheit e Kelvin.

## Funcionalidades

- **Busca de CEP**: Utiliza a API [BrasilAPI](https://brasilapi.com.br/) para converter um CEP válido em uma cidade.
- **Busca de clima**: Utiliza a API [WeatherAPI](https://www.weatherapi.com/) para obter a temperatura na cidade.
- **Conversões de temperatura**:  
  - Celsius para Fahrenheit: `F = C * 1.8 + 32`  
  - Celsius para Kelvin: `K = C + 273`
- **Observabilidade**: 
  - Traces distribuídos com OpenTelemetry
  - Visualização de traces no Zipkin
  - Métricas coletadas via Prometheus

## Observabilidade

O sistema está instrumentado com OpenTelemetry para coleta de métricas e traces distribuídos. Os traces são enviados para o Zipkin via OTLP e podem ser visualizados em http://localhost:9411.

A arquitetura de observabilidade inclui:

- **OpenTelemetry Collector**: Recebe traces e métricas dos serviços
- **Zipkin**: Armazena e visualiza os traces
- **Prometheus**: Coleta métricas do sistema

## Como usar

### Requisitos

- **CEP**: Obrigatoriamente 8 dígitos.
- **Formato de resposta em caso de sucesso** (HTTP 200):
  
  ```json
  {
    "temp_C": 28.5,
    "temp_F": 83.3,
    "temp_K": 301.5
  }
  ```

### Pré-requisitos

- **Docker** e **Docker Compose** instalados em seu ambiente.
- **Chave de API** da [WeatherAPI](https://www.weatherapi.com/).

### Configuração

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

4. **Utilizar a aplicação** (por padrão, via porta 8080):

   ```bash
   curl -X POST -H 'Content-Type: application/json' -d '{"cep": "13280001"}' "http://localhost:8080"
   ```

## Exemplo Prático

1. Fazer uma requisição ao serviço:

   ```bash
   curl -X POST -H 'Content-Type: application/json' -d '{"cep": "13400008"}' "http://localhost:8080"
   ```

2. Visualizar o trace no Zipkin:
   - Acessar http://localhost:9411
   - Procurar pelo trace correspondente (Buscar por "ServiceName": "service-a")
   - Visualizar os spans e o tempo de execução de cada serviço

Exemplo:
![image](https://github.com/user-attachments/assets/fd74fa8c-7479-48dd-b9f5-5e35b5755c74)

## Testes

- **Testes unitários**: Validam o caso de uso de busca de clima por CEP.
- **HTTP tests**: Arquivos `.http` para validação de chamadas REST:
  - `success.http`: Teste de sucesso com um CEP válido (expectativa de código **200**).
  - `invalid.http`: Teste com CEP inválido (expectativa de código **422**).
  - `not-found.http`: Teste com CEP não encontrado (expectativa de código **404**).
