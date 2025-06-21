run: build
	@./bin/dossier-org

build:
	@go build -o bin/dossier-org .

test:
	@go test -v ./...
