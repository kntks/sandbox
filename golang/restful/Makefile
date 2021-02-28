
.PHONY: fmt start build test clean


build:
	echo "compile"

start:
	reflex -r '(\.go$$|go\.mod)' -s go run cmd/server.go

test:
	go test ./... -coverprofile cover.out
	go tool cover -func cover.out | grep total | awk '{print $$3}'
	rm -rf cover.out

fmt:
	go fmt ./...

clean:
	go clean