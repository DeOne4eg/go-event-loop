run:
	docker-compose up --remove-orphans --build

run_client:
	go run -race cmd/client/main.go

run_server:
	go run -race cmd/server/main.go

test:
	go test -coverprofile=cover.out -v ./...

test_coverage:
	go tool cover -func=cover.out