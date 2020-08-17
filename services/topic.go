package services

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mrdulin/go-rpc-cnode/models"
	"github.com/mrdulin/go-rpc-cnode/utils/http"
)

var (
	ErrGetTopicsByPage = errors.New("get topics by page")
	ErrGetTopicById    = errors.New("get topic by id")
)

type (
	GetTopicByIdArgs struct {
		ID          string `json:"id" url:"id,omitempty"`
		Accesstoken string `json:"accesstoken" url:"accesstoken,omitempty"`
		Mdrender    string `json:"mdrender" url:"mdrender,omitempty"`
	}
	GetTopicsByPageArgs struct {
		Page     int             `json:"page" url:"page,omitempty"`
		Tab      models.TopicTab `json:"tab" url:"tab,omitempty"`
		Limit    int             `json:"limit" url:"limit,omitempty"`
		Mdrender string          `json:"mdrender" url:"mdrender,omitempty"`
	}
)

type topicService struct {
	HttpClient http.Client
	BaseURL    string
}

type TopicService interface {
	GetTopicsByPage(args *GetTopicByIdArgs, res *[]models.Topic) error
	GetTopicById(args *GetTopicsByPageArgs, res *models.TopicDetail) error
}

func NewTopicService(httpClient http.Client, BaseURL string) *topicService {
	return &topicService{HttpClient: httpClient, BaseURL: BaseURL}
}
func (t *topicService) GetTopicsByPage(args *GetTopicsByPageArgs, res *[]models.Topic) error {
	base, err := url.Parse(t.BaseURL + "/topics")
	if err != nil {
		fmt.Println("parse url.", err)
		return ErrGetTopicsByPage
	}
	v, err := query.Values(args)
	if err != nil {
		fmt.Printf("query.Values(args) error. args: %+v, error: %s", args, err)
		return ErrGetTopicsByPage
	}
	base.RawQuery = v.Encode()
	err = t.HttpClient.Get(base.String(), res)
	if err != nil {
		fmt.Println(err)
		return ErrGetTopicsByPage
	}
	return nil
}

func (t *topicService) GetTopicById(args *GetTopicByIdArgs, res *models.TopicDetail) error {
	base, err := url.Parse(t.BaseURL + "/topic/" + args.ID)
	if err != nil {
		fmt.Println("parse url.", err)
		return ErrGetTopicById
	}
	urlValues := url.Values{}
	if args.Accesstoken != "" {
		urlValues.Add("accesstoken", args.Accesstoken)
	}
	urlValues.Add("mdrender", args.Mdrender)
	base.RawQuery = urlValues.Encode()
	err = t.HttpClient.Get(base.String(), res)
	if err != nil {
		fmt.Println(err)
		return ErrGetTopicById
	}
	return nil
}
