package main

import (
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	client := GetClient()
	context := &Context{Client: client}

	context.BackupDiscoverWeeklyPlaylist()
}
