![Go Report](https://goreportcard.com/badge/github.com/DeOne4eg/go-event-loop)
![Repository Top Language](https://img.shields.io/github/languages/top/DeOne4eg/go-event-loop)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/DeOne4eg/go-event-loop)
![Codacy Grade](https://img.shields.io/codacy/grade/c9467ed47e064b1981e53862d0286d65)
![Github Repository Size](https://img.shields.io/github/repo-size/DeOne4eg/go-event-loop)
![Github Open Issues](https://img.shields.io/github/issues/DeOne4eg/go-event-loop)
![Lines of code](https://img.shields.io/tokei/lines/github/DeOne4eg/go-event-loop)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/DeOne4eg/go-event-loop)
![GitHub contributors](https://img.shields.io/github/contributors/DeOne4eg/go-event-loop)

<img align="right" src="./images/go.png">

# RabbitMQ Test case

## Task description

МикросервисА генерит 2 события (событие это JSON) и кладет их в очередь.

МикросервисB (инстанс 1) получает 1 событие и обрабатывает.

МикросервисB (инстанс 2) получает 2 событие и обрабатывает.

МикросервисB может отправить Ack, Nack, Reject с параметром requeue=true/false, а также может сгенерировать сообщение и также положить в очередь.

Нужно реализовать клиента к RabbitMQ с общим интерфейсом (без привязки к RabbitMQ, есть например еще ActiveMQ) со всеми методами которые нужны для реализации логики описанной выше.

## Example ENV
```
BROKER_DRIVER=rabbitmq
BROKER_HOST=localhost
BROKER_PORT=5672
BROKER_USERNAME=guest
BROKER_PASSWORD=guest
```

## Solution notes
+ 🔱 clean architecture
+ ✅ tests

## HOWTO
+ run with `make run`
+ run client with `make run_client`
+ run server with `make run_server`
+ run test with `make test`
+ run test coverage with `make test_coverage`