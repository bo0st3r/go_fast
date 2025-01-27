HOST = http://localhost:8080

GREEN = \033[0;32m
NC = \033[0m # No Color

.PHONY: health buildstart db-up db-down

buildstart:
	rm -f ./go-fast
	go build -o go-fast cmd/server/main.go
	./go-fast 

health: 
	@echo "${GREEN}How's that thang goin'?${NC}"
	curl ${HOST}/health

db-up:
	docker compose -f docker/docker-compose.yml --env-file .env up

db-down:
	docker compose -f docker/docker-compose.yml --env-file .env down