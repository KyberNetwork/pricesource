test: 
	go test -race -vet=off ./...

lint: 
	golangci-lint run ./...
