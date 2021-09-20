package auth

import (
	"context"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/idtoken"
)

var (
	googleOauthConfig *oauth2.Config
)

func init() {
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/callback",
		ClientID:     "118723902684-csndh1evv68kjef55vi98ukhmsmludk4.apps.googleusercontent.com",
		ClientSecret: "r0-Naj3yvHWWWX4lltmnV7Bs",
		Endpoint:     google.Endpoint,
	}
}

type TokenConfig struct {
	Token string `json:"token"`
}

func ValidateGoogleToken(googleToken TokenConfig) (string, error) {
	payload, err := idtoken.Validate(context.Background(), googleToken.Token, "118723902684-csndh1evv68kjef55vi98ukhmsmludk4.apps.googleusercontent.com")
	if err != nil {
		return "", err
	}
	email := payload.Claims["email"].(string)
	return email, nil
	// response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	// }

	// defer response.Body.Close()
	// contents, err := ioutil.ReadAll(response.Body)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	// }
	// fmt.Println(contents)

	// return contents, nil
}
