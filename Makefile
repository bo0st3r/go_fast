HOST = http://localhost:8080

GREEN = \033[0;32m
NC = \033[0m # No Color

.PHONY: health buildstart db-up db-down clean

dev:
	rm -f ./go-fast
	@echo "${GREEN}Building app...${NC}"
	go build -o go-fast cmd/server/main.go
	@echo "${GREEN}Running app...${NC}"
	./go-fast 

# Clean built artifacts
clean:
	rm -f ./go-fast

db-up:
	docker compose -f docker/docker-compose.yml --env-file .env up -d

db-down:
	docker compose -f docker/docker-compose.yml --env-file .env down

health: 
	@echo "${GREEN}How's that thang goin'?${NC}"
	curl -f ${HOST}/health