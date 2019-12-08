# Spotify: "Discover Weekly" playlist backup

Backup **Discover Weekly** playlist.

## Requirements

- [dep](https://golang.github.io/dep/)

## Setup

```shell
# Install dependencies
dep ensure -v

# Complete your .env file or export vars to environment
cp .env.dist .env
```

## Usage

```shell
# Obtain ACCESS_TOKEN and REFRESH_TOKEN and copy to .env file
go run ./cmd/credentials


# Backup your playlist
go run ./cmd/backup
```

## Build

```shell
# Obtain credentials
go build -o bin/credentials ./cmd/credentials
./bin/credentials

# Backup playlist
go build -o bin/backup cmd/backup/main.go
./bin/backup
```

## Tests

```shell
go test ./cmd/backup
```
