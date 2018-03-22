all:
	go build
	
test: 
	go test -race ./...

run:
	go run ./cmd/main.go