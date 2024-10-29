**[EN](README.md)** | [PT-BR](README_pt-BR.md)

# Stress test with GO

The project is a stress tester that uses multithreading to throw multiple requests to a determined url/endpoint.
This project also has a server app available to be used as a endpoint for the stress test if wanted. It returns random status codes, giving priority to status code 200.

## Requirements

[Docker](https://www.docker.com/)

## Instructions to run:

#### Server (optional):
```
docker build --build-arg APP_NAME=server -t stress-test-server:latest .
docker run stress-test-server
```
The api will run with an endpoint at http://localhost:8080/
You will need to open another terminal for the stress tester.

#### Stress tester:
```
docker build -t stress-test:latest .
docker run stress-test --url=[url] --requests=[requests] --concurrency=[concurrency]
```
##### Parameters:
- `url`: The url for the tested service. *Required*
- `requests`: Defines the total number of requests. *Optional. Default: 10*
- `concurrency`: Defines the number of simultaneous requests. *Optional. Default: 2*