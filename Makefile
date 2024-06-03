build:
	go build -o bin/gokeys cmd/main.go

run: build
	./bin/gokeys

test:
	go test -v ./... -count=1