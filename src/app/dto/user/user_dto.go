package user

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type UserReqDTOInterface interface {
	Validate() error
}

type RegisterReqDTO struct {
	UserName string `json:"username"`
}

func (dto *RegisterReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type RegisterRespDTO struct {
	ID       int64  `json:"id" db:"id"`
	UserName string `json:"username" db:"username"`
	WalletID int64  `json:"wallet_id" db:"wallet_id"`
	Token    string `json:"token"`
}

type LoginReqDTO struct {
	UserName string `json:"username"`
}

func (dto *LoginReqDTO) Validate() error {
	if err := validation.ValidateStruct(
		dto,
		validation.Field(&dto.UserName, validation.Required),
	); err != nil {
		return err
	}
	return nil
}

type LoginSocialMediaRespDTO struct {
	Url string `json:"url"`
}

type UserInfoGoogleDTO struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
	PicURL        string `json:"picture"`
	VerifiedEmail bool   `json:"verified_email"`
}
