APP_NAME := product-sorter

.PHONY: build run test docker-build docker-run clean lint

build:
	go build -o $(APP_NAME) main.go

run: build
	./$(APP_NAME)

test:
	go test -v ./...

lint:
	golangci-lint run --timeout=3m

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run --rm -it $(APP_NAME):latest

clean:
	rm -f $(APP_NAME)
