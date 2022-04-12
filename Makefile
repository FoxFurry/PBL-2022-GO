run:
	go run main.go serve

build:
	go build -o bin/petfeeder

test:
	go test ./...

migrate:
	./migrations/migrate.sh