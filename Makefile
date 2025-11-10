run:
	go run cmd/main.go

test:
	go test ./tests -v

build:
	go build -o taskify cmd/main.go

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

clean:
	rm -f taskify

.PHONY: run test build docker-up docker-down docker-logs clean
