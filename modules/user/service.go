package user

import (
  "errors"
  "fmt"
"github.com/mrdulin/go-rpc-cnode/utils/http"
)

type GetUserByLoginnameArgs struct {
  Loginname string
}
type ValidateAccessTokenArgs struct {
  AccessToken string
}

type validateAccessTokenRequestPayload struct {
  AccessToken string `json:"accesstoken"`
}

type Service struct {
  HttpClient http.Client
  BaseURL    string
}

type IService interface {
  GetUserByLoginname(args *GetUserByLoginnameArgs, res *UserDetail) error 
  ValidateAccessToken(args *ValidateAccessTokenArgs, res *UserEntity) error 
}

func (u *Service) GetUserByLoginname(args *GetUserByLoginnameArgs, res *UserDetail) error {
  endpoint := u.BaseURL + "/user/" + args.Loginname
  err := u.HttpClient.Get(endpoint, &res)
  if err != nil {
    fmt.Println(err)
    return errors.New("Get user by loginname")
  }
  return nil
}

func (u *Service) ValidateAccessToken(args *ValidateAccessTokenArgs, res *UserEntity) error {
  url := u.BaseURL + "/accesstoken"
  err := u.HttpClient.Post(url, &validateAccessTokenRequestPayload{AccessToken: args.AccessToken}, &res)
  if err != nil {
    fmt.Println(err)
    return errors.New("Validate accessToken")
  }
  return nil
}