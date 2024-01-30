.PHONY: build
build:
	go build -v -o goserver ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...

.PHONY: db
db:
	migrate -path migrations -database "postgres://localhost:5432/go?sslmode=disable" down 
	migrate -path migrations -database "postgres://localhost:5432/go?sslmode=disable" up

.PHONY: prod
prod:
	@if [ "$(filter windows,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=windows GOARCH=amd64 go build -o goserver-win-x86.exe -v ./cmd/apiserver; \
	elif [ "$(filter macos,$(MAKECMDGOALS))" != "" ]; then \
		GOOS=darwin GOARCH=amd64 go build -o goserver-darwin-amd64 -v ./cmd/apiserver; \
	else \
		GOOS=linux GOARCH=amd64 go build -o goserver-linux-amd64 -v ./cmd/apiserver; \
	fi

.PHONE: run
run:
	go run ./cmd/apiserver/main.go

PHONY: clean
clean:
	rm -f goserver goserver-win-x86.exe goserver-darwin-amd64 goserver-linux-amd64
	
.DEFAULT_GOAL := build
