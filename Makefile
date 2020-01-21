.PHONY: build_all build_credentials build_backup run_credentials run_backup credentials backup clean

build_all: build_credentials build_backup

build_credentials:
	@go build -o bin/credentials ./cmd/credentials

build_backup:
	@go build -o bin/backup ./cmd/backup

run_credentials:
	@go run ./cmd/credentials

run_backup:
	@go run ./cmd/backup

credentials: build_credentials
	@./bin/credentials

backup: build_backup
	@./bin/backup

dependencies:
	@dep ensure -v

test:
	@go test ./cmd/backup

clean:
	@rm -rf bin/credentials bin/backup
