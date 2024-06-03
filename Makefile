build:
	go build -o bin/gokeys cmd/gokeys/main.go

run: build
	./bin/gokeys

test:
	go test -v ./... -count=1