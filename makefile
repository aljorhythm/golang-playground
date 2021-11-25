microtest:
	go test ./...

format:
	go fmt ./...

run:
	go run .

all: format microtest