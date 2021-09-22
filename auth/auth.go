package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
	"watsup.com/envload"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  envload.GoDotEnvVariable("REDIRECTURL"),
		ClientID:     envload.GoDotEnvVariable("GOOGLECLIENTID"),
		ClientSecret: envload.GoDotEnvVariable("GOOGLECLIENTSECRET"),
		Endpoint:     google.Endpoint,
	}
}

type TokenConfig struct {
	Token string `json:"token"`
}

func ValidateGoogleToken(googleToken TokenConfig) (string, error) {
	google_clientid := envload.GoDotEnvVariable("GOOGLECLIENTID")
	payload, err := idtoken.Validate(context.Background(), googleToken.Token, google_clientid)
	if err != nil {
		return "", err
	}
	email := payload.Claims["email"].(string)
	return email, nil
}
