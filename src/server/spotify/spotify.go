package spotify

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func Authenticate() {
	authConfig := &clientcredentials.Config{

		TokenURL: spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)

	playlistID := spotify.ID("37i9dQZF1DXcBWIGoYBM5M")
	playlist, err := client.GetPlaylist(playlistID)
	if err != nil {
		log.Fatalf("error retrieve playlist data: %v", err)
	}

	log.Println("playlist id:", playlist.ID)
	log.Println("playlist name:", playlist.Name)
	log.Println("playlist description:", playlist.Description)
}
