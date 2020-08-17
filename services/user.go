package services

import (
	"errors"
	"fmt"

	"github.com/mrdulin/go-rpc-cnode/models"
	"github.com/mrdulin/go-rpc-cnode/utils/http"
)

var (
	ErrGetUserByLoginname  = errors.New("get user by login name")
	ErrValidateAccessToken = errors.New("Validate accessToken")
)

type (
	GetUserByLoginnameArgs struct {
		Loginname string
	}
	ValidateAccessTokenArgs struct {
		AccessToken string
	}
	validateAccessTokenRequestPayload struct {
		AccessToken string `json:"accesstoken"`
	}
)

type userService struct {
	HttpClient http.Client
	BaseURL    string
}

type UserService interface {
	GetUserByLoginname(args *GetUserByLoginnameArgs, res *models.UserDetail) error
	ValidateAccessToken(args *ValidateAccessTokenArgs, res *models.UserEntity) error
}

func NewUserService(httpClient http.Client, baseURL string) *userService {
	return &userService{HttpClient: httpClient, BaseURL: baseURL}
}

func (u *userService) GetUserByLoginname(args *GetUserByLoginnameArgs, res *models.UserDetail) error {
	endpoint := u.BaseURL + "/user/" + args.Loginname
	err := u.HttpClient.Get(endpoint, &res)
	if err != nil {
		fmt.Println(err)
		return ErrGetUserByLoginname
	}
	return nil
}

func (u *userService) ValidateAccessToken(args *ValidateAccessTokenArgs, res *models.UserEntity) error {
	url := u.BaseURL + "/accesstoken"
	err := u.HttpClient.Post(url, &validateAccessTokenRequestPayload{AccessToken: args.AccessToken}, &res)
	if err != nil {
		fmt.Println(err)
		return ErrValidateAccessToken
	}
	return nil
}
