# Spotify: "Discover Weekly" playlist backup

Backup **Discover Weekly** playlist.

## Requirements

- [dep](https://golang.github.io/dep/)

## Usage

```shell
# Install dependencies
dep ensure -v

# Complete your .env file
cp .env.dist .env

# Obtain ACCESS_TOKEN and REFRESH_TOKEN and copy to .env file
go build -o bin/credentials cmd/credentials/main.go
./bin/credentials

# Backup your playlist
go build -o bin/backup cmd/backup/main.go
./bin/backup
```
