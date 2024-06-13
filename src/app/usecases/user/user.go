package user

import (
	"context"
	"errors"
	dto "ex_service/src/app/dto/user"
	helper "ex_service/src/infra/helper"
	integ "ex_service/src/infra/integration/oauthgoogle"
	repo "ex_service/src/infra/persistence/postgres/user"

	"log"
)

type UserUCInterface interface {
	Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error)
	Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error)
	LoginSocialMedia(provider string) (*dto.LoginSocialMediaRespDTO, error)
	ExchangeCodeGoogle(ctx context.Context, code string) (*dto.UserInfoGoogleDTO, error)
}

type userUseCase struct {
	Repo             repo.UserRepository
	OauthGoogleInteg integ.OauthGoogleService
}

func NewUserUseCase(repo repo.UserRepository, o integ.OauthGoogleService) UserUCInterface {
	return &userUseCase{
		Repo:             repo,
		OauthGoogleInteg: o,
	}
}

func (uc *userUseCase) Register(data *dto.RegisterReqDTO) (*dto.RegisterRespDTO, error) {
	result, err := uc.Repo.Register(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result.Token, err = helper.GenerateToken(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (uc *userUseCase) Login(data *dto.LoginReqDTO) (*dto.RegisterRespDTO, error) {

	result, err := uc.Repo.Login(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	result.Token, err = helper.GenerateToken(result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *userUseCase) LoginSocialMedia(provider string) (*dto.LoginSocialMediaRespDTO, error) {
	if provider == "google" {
		url := u.OauthGoogleInteg.GetUrl("google")
		return &dto.LoginSocialMediaRespDTO{Url: url}, nil
	}

	return nil, errors.New("provider not found")
}

func (u *userUseCase) ExchangeCodeGoogle(ctx context.Context, code string) (*dto.UserInfoGoogleDTO, error) {
	token, err := u.OauthGoogleInteg.ExchangeCode(ctx, code)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	userInfo, err := u.OauthGoogleInteg.GetUser(ctx, token)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return userInfo, nil
}
