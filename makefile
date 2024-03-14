run:
	go run ./cmd

test:
	go test ./... -v -cover

coverage:
	go test ./... -v -coverprofile=coverage.out
