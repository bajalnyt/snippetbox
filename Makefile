run:
	go run ./cmd/web -addr=":4000"


test:
	go test -v ./cmd/web
