package youtube

import (
	"context"
	"log"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTubeClient struct {
	service *youtube.Service
}

func NewYouTubeClient(apiKey string) *YouTubeClient {
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Erro ao inicializar YouTube API: %v", err)
	}
	return &YouTubeClient{service: service}
}

func (yt *YouTubeClient) CreatePlaylist(title string) string {
	playlist := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title: title,
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: "private",
		},
	}

	resp, err := yt.service.Playlists.Insert([]string{"snippet,status"}, playlist).Do()
	if err != nil {
		log.Fatalf("Erro ao criar playlist: %v", err)
	}
	return resp.Id
}

func (yt *YouTubeClient) AddTrackToPlaylist(playlistID, query string) {
	searchCall := yt.service.Search.List([]string{"snippet"}).Q(query).Type("video").MaxResults(1)
	result, err := searchCall.Do()
	if err != nil || len(result.Items) == 0 {
		log.Printf("Erro ao buscar vídeo para '%s': %v", query, err)
		return
	}

	videoID := result.Items[0].Id.VideoId

	item := &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlistID,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: videoID,
			},
		},
	}

	_, err = yt.service.PlaylistItems.Insert([]string{"snippet"}, item).Do()
	if err != nil {
		log.Printf("Erro ao adicionar vídeo '%s': %v", query, err)
	}
}