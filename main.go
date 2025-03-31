package main

import (
	"fmt"

	"github.com/vinnieoh/script-spotify-youtube/app/config"
	"github.com/vinnieoh/script-spotify-youtube/app/spotify"
	"github.com/vinnieoh/script-spotify-youtube/app/youtube"
)



func main(){

	cfg := config.EnvConfig()

	spotifyClientID := cfg.SPOTIFY_CLIENT_ID
	spotifyClientSecret := cfg.SPOTIFY_CLIENT_SECRET
	spotifyPlaylistID := cfg.SPOTIFY_PLAYLIST_ID
	youtubeAPIKey := cfg.YOUTUBE_API_KEY


	// Spotify
	sp := spotify.NewSpotifyClient(spotifyClientID, spotifyClientSecret)
	tracks := sp.GetTrackNamesFromPlaylist(spotifyPlaylistID)

	// YouTube
	yt := youtube.NewYouTubeClient(youtubeAPIKey)
	playlistID := yt.CreatePlaylist("Playlist migrada do Spotify")

	for _, track := range tracks {
		fmt.Println("Adicionando:", track)
		yt.AddTrackToPlaylist(playlistID, track)
	}

	fmt.Println("✅ Playlist criada com sucesso!")
	


}