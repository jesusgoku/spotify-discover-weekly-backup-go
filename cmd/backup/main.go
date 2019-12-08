package main

import (
	"fmt"
	"os"
	"time"

	"github.com/snabb/isoweek"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	accessToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	token := new(oauth2.Token)

	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	token.TokenType = "Bearer"
	token.Expiry = time.Date(2009, time.October, 7, 23, 0, 0, 0, time.UTC)

	client := spotify.NewAuthenticator("").NewClient(token)

	user, _ := client.CurrentUser()

	result, _ := client.Search("Discover Weekly", spotify.SearchTypePlaylist)

	discoverWeeklyID := result.Playlists.Playlists[0].ID
	discoverWeeklyPlaylist, _ := client.GetPlaylist(discoverWeeklyID)

	tracksID := make([]spotify.ID, 0, len(discoverWeeklyPlaylist.Tracks.Tracks))

	for _, track := range discoverWeeklyPlaylist.Tracks.Tracks {
		tracksID = append(tracksID, track.Track.SimpleTrack.ID)
	}

	now := time.Now()
	year, week := now.ISOWeek()
	firstDayOfWeek := isoweek.StartTime(year, week, now.Location())
	backupPlaylistName := fmt.Sprintf("Lista Semanal %s", firstDayOfWeek.Format("20060102"))

	backupPlaylist, _ := client.CreatePlaylistForUser(user.ID, backupPlaylistName, "", false)

	client.AddTracksToPlaylist(backupPlaylist.ID, tracksID...)
}
