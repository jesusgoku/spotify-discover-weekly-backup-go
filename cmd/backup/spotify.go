package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

// Context for application
type Context struct {
	Client spotify.Client
	User   *spotify.PrivateUser
}

// CurrentUser retrieve current user for client
func (c *Context) CurrentUser() (*spotify.PrivateUser, error) {
	return c.Client.CurrentUser()
}

// SetCurrentUser check and set User in Context
func (c *Context) SetCurrentUser() {
	if c.User != nil {
		return
	}

	user, _ := c.CurrentUser()

	c.User = user
}

// DiscoverWeeklyPlaylistTracks retrieve tracks ids for "Discover Weekly" playlist
func (c *Context) DiscoverWeeklyPlaylistTracks() []spotify.ID {
	result, _ := c.Client.Search("Discover Weekly", spotify.SearchTypePlaylist)

	discoverWeeklyID := result.Playlists.Playlists[0].ID
	discoverWeeklyPlaylist, _ := c.Client.GetPlaylist(discoverWeeklyID)

	tracksID := make([]spotify.ID, 0, len(discoverWeeklyPlaylist.Tracks.Tracks))

	for _, track := range discoverWeeklyPlaylist.Tracks.Tracks {
		tracksID = append(tracksID, track.Track.SimpleTrack.ID)
	}

	return tracksID
}

// BackupDiscoverWeeklyPlaylist backup "Discover Weekly" user playlist
func (c *Context) BackupDiscoverWeeklyPlaylist() {
	c.SetCurrentUser()

	tracksIDs := c.DiscoverWeeklyPlaylistTracks()

	firstDayOfWeek := GetFirstDayForCurrentWeek()
	backupPlaylistName := fmt.Sprintf("Lista Semanal %s", firstDayOfWeek.Format("20060102"))

	backupPlaylist, _ := c.Client.CreatePlaylistForUser(c.User.ID, backupPlaylistName, "", false)

	c.Client.AddTracksToPlaylist(backupPlaylist.ID, tracksIDs...)
}

// GetClient retrieve Spotify Client from ACCESS_TOKEN and REFRESH_TOKEN
func GetClient() spotify.Client {
	accessToken := os.Getenv("SPOTIFY_ACCESS_TOKEN")
	refreshToken := os.Getenv("SPOTIFY_REFRESH_TOKEN")

	token := &oauth2.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		Expiry:       time.Date(2009, time.October, 7, 23, 0, 0, 0, time.UTC),
	}

	return spotify.NewAuthenticator("").NewClient(token)
}
