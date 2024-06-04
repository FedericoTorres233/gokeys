build:
	go build -o bin/gokeys cmd/main.go

run: build
	./bin/gokeys

test:
	go test -v ./... -count=1

tidy:
	go mod tidy -v

format:
	go run mvdan.cc/gofumpt@latest -w -l .