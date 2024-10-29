[EN](README.md) | **[PT-BR](README_pt-BR.md)**

# Teste de estresse com GO

O projeto é um testador de estresse que utiliza multithreading para enviar múltiplas requisições a uma url/endpoint determinado.
Este projeto também possui uma aplicação de servidor disponível para ser usada como um endpoint para o teste de estresse, caso desejado. Ele retorna códigos de status aleatórios, dando prioridade ao código de status 200.

## Requisitos

[Docker](https://www.docker.com/)

## Instruções para executar:

#### Servidor (opcional):
```
docker build --build-arg APP_NAME=server -t stress-test-server:latest .
docker run stress-test-server
```
A API estará disponível com um endpoint em http://localhost:8080/
Será necessário abrir outro terminal para o teste de estresse.

#### Teste de estresse:
```
docker build -t stress-test:latest .
docker run stress-test --url=[url] --requests=[requests] --concurrency=[concurrency]
```
##### Parâmetros:
- `url`: A URL para o serviço a ser testado. *Obrigatório*
- `requests`: Define o número total de requisições. *Opcional. Padrão: 10*
- `concurrency`: Define o número de requisições simultâneas. *Opcional. Padrão: 2*