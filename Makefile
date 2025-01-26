HOST = http://localhost:8080

GREEN = \033[0;32m
NC = \033[0m # No Color

.PHONY: health

health: 
	@echo "${GREEN}How's that thang goin'?${NC}"
	curl ${HOST}/health