serve:  
	go run ./cmd/server/main.go

serve_prod:  
	go run ./cmd/server/main.go -env=pro

build:
	go build -o bin/lofola ./cmd/server/main.go

run_prod:
	./bin/lofola -env=pro

build_docker:
	docker build -f deployments/Dockerfile -t lofola:0.0.1 .

compose-up:
	docker compose -f deployments/docker-compose.yaml up

compose-down:
	docker compose -f deployments/docker-compose.yaml down

wire:
	wire ./internal/wire

.PHONY: wire