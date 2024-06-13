package oauth

import (
	"context"
	"encoding/json"
	dto "ex_service/src/app/dto/user"
	"ex_service/src/infra/config"

	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthGoogleService interface {
	GetUrl(state string) (url string)
	ExchangeCode(ctx context.Context, code string) (token string, err error)
	GetUser(ctx context.Context, token string) (user *dto.UserInfoGoogleDTO, err error)
}

func NewOauthGoogleService() OauthGoogleService {
	var googleOauthCfg = oauth2.Config{
		ClientID:     config.Envs.Oauth.Google.ClientID,
		ClientSecret: config.Envs.Oauth.Google.ClientSecret,
		RedirectURL:  config.Envs.Oauth.Google.CallbackURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return &oauthGoogleService{
		oauthCfg: googleOauthCfg,
	}
}

type oauthGoogleService struct {
	oauthCfg oauth2.Config
}

func (o *oauthGoogleService) GetUrl(state string) (url string) {
	return o.oauthCfg.AuthCodeURL(state)
}

func (o *oauthGoogleService) ExchangeCode(ctx context.Context, code string) (token string, err error) {
	tok, err := o.oauthCfg.Exchange(ctx, code)
	if err != nil {
		logrus.Error("Error exchanging code for token: ", err)
		return "", err
	}

	return tok.AccessToken, nil
}

func (o *oauthGoogleService) GetUser(ctx context.Context, token string) (user *dto.UserInfoGoogleDTO, err error) {
	var (
		userInfo = &dto.UserInfoGoogleDTO{}
		client   = o.oauthCfg.Client(ctx, &oauth2.Token{AccessToken: token})
	)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		logrus.Error("Error getting user info: ", err)
		return userInfo, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(userInfo)
	if err != nil {
		return userInfo, err
	}

	return userInfo, nil
}
