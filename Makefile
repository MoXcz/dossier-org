run: build
	@./bin/dossier-org

build:
	@go build -o bin/dossier-org .

test:
	@go test -v ./...

migrations:
	@goose reset && goose up

seed:
	@go run ./scripts/seed.go
