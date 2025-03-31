package main

import (
	"log"
	"net/http"
	"text/template"

	"github.com/vinnieoh/script-spotify-youtube/app/config"
	"github.com/vinnieoh/script-spotify-youtube/app/spotify"
	"github.com/vinnieoh/script-spotify-youtube/app/youtube"
)



func main(){

	http.HandleFunc("/", formHandler)
	http.HandleFunc("/sync", syncHandler)
	log.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func syncHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	playlistID := r.FormValue("playlistID")
	if playlistID == "" {
		http.Error(w, "ID da playlist é obrigatório", http.StatusBadRequest)
		return
	}

	cfg := config.EnvConfig()

	spotifyClientID := cfg.SPOTIFY_CLIENT_ID
	spotifyClientSecret := cfg.SPOTIFY_CLIENT_SECRET
	youtubeAPIKey := cfg.YOUTUBE_API_KEY

	// Lógica de sincronização
	sp := spotify.NewSpotifyClient(spotifyClientID, spotifyClientSecret)
	tracks := sp.GetTrackNamesFromPlaylist(playlistID)

	yt := youtube.NewYouTubeClient(youtubeAPIKey)
	ytPlaylistID := yt.CreatePlaylist("Migrada do Spotify")

	var addedTracks []string

	for _, track := range tracks {
		log.Println("Adicionando:", track)
		yt.AddTrackToPlaylist(ytPlaylistID, track)
		addedTracks = append(addedTracks, track)
	}

	// Renderiza o resultado no template
	tmpl := template.Must(template.ParseFiles("templates/result.html"))
	data := struct {
		PlaylistID   string
		Tracks       []string
		Total        int
	}{
		PlaylistID: playlistID,
		Tracks:     addedTracks,
		Total:      len(addedTracks),
	}

	tmpl.Execute(w, data)

	w.Write([]byte("✅ Playlist criada no YouTube com sucesso!"))
}