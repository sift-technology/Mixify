package spotify

import (
	"context"
	"log"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
)

func Authenticate() {

	authConfig := &clientcredentials.Config{
		ClientID:     "6277022508444e50870dde60e2131882",
		ClientSecret: "78bde7df219c400bb0bcd6e658e76d06",
		TokenURL:     spotify.TokenURL,
	}

	accessToken, err := authConfig.Token(context.Background())
	if err != nil {
		log.Fatalf("error retrieve access token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(accessToken)
	Recommend(&client)

}
