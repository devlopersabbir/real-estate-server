BINARY_NAME=juan-server

docker-build:
	docker compose -f docker-compose.yaml --profile prod build --no-cache

docker-dev:
	docker compose -f docker-compose.yaml --profile dev up -d

docker-up:
	docker compose -f docker-compose.yaml --profile prod up -d

docker-down:
	docker compose -f docker-compose.yaml --profile prod down

start-prod:
	./bin/$(BINARY_NAME)

build:
	go build -o bin/$(BINARY_NAME) ./cmd/main.go

deps:
	go mod download
	go mod tidy
dev:
	air

swagger:
	@echo "Generating Swagger docs..."
	$(shell go env GOPATH)/bin/swag fmt
	$(shell go env GOPATH)/bin/swag init \
	-g startup/server.go \
	-d ./ \
	--parseDependency \
	--parseInternal