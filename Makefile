# * FILE RUN GO
serve:  
	go run ./cmd/server/main.go
build_docker:
	docker build -f deployments/Dockerfile -t lofola:0.0.1 .
compose-up:
	docker compose -f deployments/docker-compose.yaml up
compose-down:
	docker compose -f deployments/docker-compose.yaml down