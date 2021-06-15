COMPOSE_FILE = docker/docker-compose.yml

all: 
	@echo "Hello World"

compose-up:
	docker compose -f ${COMPOSE_FILE} up -d

compose-down:
	docker compose -f ${COMPOSE_FILE} down --rmi all

compose-start:
	docker compose -f ${COMPOSE_FILE} start

compose-stop:
	docker compose -f ${COMPOSE_FILE} stop
