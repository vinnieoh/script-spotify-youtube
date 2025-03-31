package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SPOTIFY_CLIENT_ID     string
	SPOTIFY_CLIENT_SECRET string
	SPOTIFY_PLAYLIST_ID   string
	YOUTUBE_API_KEY       string
}

func EnvConfig() *Config {

	if err := godotenv.Load("./dotenv_files/.env"); err != nil {
		log.Printf("Error loading .env file")
	}

	cfg := &Config{}

	cfg.SPOTIFY_CLIENT_ID = os.Getenv("SPOTIFY_CLIENT_ID")
	cfg.SPOTIFY_CLIENT_SECRET = os.Getenv("SPOTIFY_CLIENT_SECRET")
	cfg.SPOTIFY_PLAYLIST_ID = os.Getenv("SPOTIFY_PLAYLIST_ID")
	cfg.YOUTUBE_API_KEY = os.Getenv("YOUTUBE_API_KEY")

	if cfg.SPOTIFY_CLIENT_ID == "" || cfg.SPOTIFY_CLIENT_SECRET == "" || 
		cfg.SPOTIFY_PLAYLIST_ID == "" || cfg.YOUTUBE_API_KEY == "" {

		log.Fatal("Variáveis de ambiente não definidas.")
	}

	return cfg

}