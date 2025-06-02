package spotify

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

type SpotifyClient struct {
	client *spotify.Client
}

func NewSpotifyClient(clientID, clientSecret string) *SpotifyClient {
	ctx := context.Background()

	config := &clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		log.Fatalf("Erro ao gerar token Spotify: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)
	return &SpotifyClient{client: &client}
}


func (s *SpotifyClient) GetTrackNamesFromPlaylist(playlistID string) []string {
	playlist, err := s.client.GetPlaylist(spotify.ID(playlistID))
	if err != nil {
		log.Fatalf("Erro ao buscar playlist: %v", err)
	}

	var tracks []string
	for _, item := range playlist.Tracks.Tracks {
		track := item.Track
		tracks = append(tracks, fmt.Sprintf("%s %s", track.Name, track.Artists[0].Name))
	}
	return tracks
}