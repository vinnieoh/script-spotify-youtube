package youtube

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YouTubeClient struct {
	service *youtube.Service
}

// Cria o cliente autenticado via OAuth2
func NewYouTubeClientOAuth(clientID, clientSecret string) *YouTubeClient {
	ctx := context.Background()

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       []string{youtube.YoutubeScope},
		Endpoint:     google.Endpoint,
		RedirectURL:  "urn:ietf:wg:oauth:2.0:oob", // modo CLI
	}

	// Passo 1: abrir link
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("👉 Acesse este link e cole o código de autorização:\n%s\n\n", authURL)

	// Passo 2: usuário cola o código
	var code string
	fmt.Print("🔑 Código: ")
	fmt.Scan(&code)

	// Passo 3: troca código pelo token
	token, err := config.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Erro ao obter token: %v", err)
	}

	client := config.Client(ctx, token)

	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Erro ao criar serviço YouTube: %v", err)
	}

	return &YouTubeClient{service: service}
}

// Cria uma nova playlist no YouTube
func (yt *YouTubeClient) CreatePlaylist(title string) string {
	playlist := &youtube.Playlist{
		Snippet: &youtube.PlaylistSnippet{
			Title:       title,
			Description: "Playlist sincronizada via script",
		},
		Status: &youtube.PlaylistStatus{
			PrivacyStatus: "private", // pode ser "public", "unlisted" ou "private"
		},
	}

	call := yt.service.Playlists.Insert([]string{"snippet", "status"}, playlist)
	created, err := call.Do()
	if err != nil {
		log.Fatalf("Erro ao criar playlist no YouTube: %v", err)
	}
	return created.Id
}

// Busca o vídeo pelo nome e o adiciona à playlist
func (yt *YouTubeClient) AddTrackToPlaylist(playlistID, query string) {
	searchCall := yt.service.Search.List([]string{"snippet"}).
		Q(query).
		Type("video").
		MaxResults(1)

	response, err := searchCall.Do()
	if err != nil {
		log.Printf("Erro ao buscar vídeo '%s': %v", query, err)
		return
	}

	if len(response.Items) == 0 {
		log.Printf("Nenhum resultado encontrado para: %s", query)
		return
	}

	videoID := response.Items[0].Id.VideoId

	item := &youtube.PlaylistItem{
		Snippet: &youtube.PlaylistItemSnippet{
			PlaylistId: playlistID,
			ResourceId: &youtube.ResourceId{
				Kind:    "youtube#video",
				VideoId: videoID,
			},
		},
	}

	insertCall := yt.service.PlaylistItems.Insert([]string{"snippet"}, item)
	_, err = insertCall.Do()
	if err != nil {
		log.Printf("Erro ao adicionar vídeo '%s': %v", query, err)
	}
}
