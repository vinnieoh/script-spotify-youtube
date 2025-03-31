# script-spotify-youtube


Sincronize suas playlists do **Spotify** com o **YouTube** automaticamente usando Go!

Este script busca músicas de uma playlist pública no Spotify, pesquisa os vídeos correspondentes no YouTube e cria (ou atualiza) uma playlist com esses vídeos no YouTube.

---

## 🚀 Funcionalidades

- 🔍 Busca músicas de uma playlist do Spotify
- 🎥 Encontra os vídeos correspondentes no YouTube
- 📺 Cria uma nova playlist no YouTube com os resultados
- 🌐 Interface web simples para inserir o ID da playlist

---

## 🛠 Requisitos

- Go 1.23 ou superior
- Conta de desenvolvedor no Spotify e Google
- Chaves de API:
  - `Spotify`: Client ID e Secret
  - `YouTube`: API Key (YouTube Data API v3 ativada)

---

## ⚙️ Configuração

Crie o arquivo de variáveis de ambiente em:

```
./dotenv_files/.env
```

Com o seguinte conteúdo:

```env
SPOTIFY_CLIENT_ID=seu_spotify_client_id
SPOTIFY_CLIENT_SECRET=seu_spotify_client_secret
YOUTUBE_API_KEY=sua_youtube_api_key
```

> ✅ O script já carrega esse arquivo automaticamente ao iniciar, usando `godotenv`.

---

## ▶️ Como Usar

### ✅ Rodando no terminal

Edite o ID da playlist diretamente no código (`main.go`) ou use a interface web.

```bash
make run
```

---

### 🌐 Rodando com interface web

1. Execute:

```bash
make run
```

2. Acesse: [http://localhost:8080](http://localhost:8080)

3. Insira o ID da playlist do Spotify no campo e clique em **Sincronizar**.

---

## 🧪 Testes

```bash
make test
```

---

## 📦 Compilar o binário

```bash
make build
```

O executável será gerado em `./bin/script-spotify-youtube`.

---

## 📁 Estrutura do Projeto

```
script-spotify-youtube/
│
├── main.go                 # Entrada principal da aplicação
├── spotify/                # Lógica de conexão com o Spotify
├── youtube/                # Lógica de conexão com o YouTube
├── config/                 # Carregamento de variáveis e utilitários
├── templates/              # HTML da interface web
├── dotenv_files/           # Pasta onde fica o arquivo .env
├── Makefile                # Comandos úteis (run, build, test, etc)
└── README.md               # Este arquivo :)
```



## 📄 Licença

MIT © [vinnieoh](https://github.com/vinnieoh)



## 💡 Dica

> Você pode testar com playlists públicas do Spotify — basta copiar o ID do final da URL ou URI e colar na interface.

