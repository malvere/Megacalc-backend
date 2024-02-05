build:
	go build -v -o goserver ./cmd/server

test:
	go test -v -race -timeout 30s ./...

db:
	migrate -path migrations -database "postgres://localhost:5432/go?sslmode=disable" down 
	migrate -path migrations -database "postgres://localhost:5432/go?sslmode=disable" up

prod:
	@if [ "$(filter windows,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=windows GOARCH=amd64 go build -o goserver-win-x86.exe -v ./cmd/apiserver; \
	elif [ "$(filter macos,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=darwin GOARCH=amd64 go build -o goserver-darwin-amd64 -v ./cmd/apiserver; \
	else \
		GOOS=linux GOARCH=amd64 go build -o goserver-linux-amd64 -v ./cmd/apiserver; \
	fi

run:
	go run ./cmd/apiserver/main.go

clean:
	rm -f goserver goserver-win-x86.exe goserver-darwin-amd64 goserver-linux-amd64

sqlc:
	sqlc generate

.PHONY: build, test, db, prod, run, clean
.DEFAULT_GOAL := build
