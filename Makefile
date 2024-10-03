# Makefile для проекта MenuShka

.PHONY: all build run clean tidy

SERVICES = api-gateway menu-service recipe-service fridge-service auth-service

all: build

build:
	@for service in $(SERVICES); do \
		echo "Building $$service..."; \
		cd $$service && go build -o $$service; \
		cd ..; \
	done

run:
	docker compose up --build

stop:
	docker compose down


clean:
	@for service in $(SERVICES); do \
		echo "Cleaning $$service..."; \
		cd $$service && go clean; \
		rm -f $$service; \
		cd ..; \
	done
	docker compose down --volumes

tidy:
	@for service in $(SERVICES); do \
		echo "Running go mod tidy in $$service..."; \
		cd $$service && go mod tidy; \
		cd ..; \
	done
