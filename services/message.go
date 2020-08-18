package services

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/mrdulin/go-rpc-cnode/models"
	"github.com/mrdulin/go-rpc-cnode/utils/http"
)

var (
	ErrGetMessages      = errors.New("get messages")
	ErrGetUnreadMessage = errors.New("get unread message")
	ErrMarkOneMessage   = errors.New("mark one message")
	ErrMarkAllMessages  = errors.New("mark all messages")
)

type (
	MarkOneMessageRequestPayload struct {
		Accesstoken string `json:"accesstoken"`
	}
	MarkAllMessagesRequestPayload struct {
		Accesstoken string `json:"accesstoken"`
	}

	GetMessagesArgs struct {
		Accesstoken, Mdrender string
	}
	GetUnreadMessageArgs struct {
		Accesstoken string
	}
	MarkOneMessageArgs struct {
		Accesstoken, ID string
	}
	MarkAllMessagesArgs struct {
		Accesstoken string
	}
)

type messageService struct {
	HttpClient http.Client
	BaseURL    string
}

type MessageService interface {
	GetMessages(args *GetMessagesArgs, res *models.MessagesResponse) error
	GetUnreadMessage(args *GetUnreadMessageArgs, res *int) error
	MarkOneMessage(args *MarkOneMessageArgs, res *string) error
	MarkAllMessages(args *MarkAllMessagesArgs, res *[]models.MarkedMessage) error
}

func NewMessageService(httpClient http.Client, BaseURL string) *messageService {
	return &messageService{HttpClient: httpClient, BaseURL: BaseURL}
}

func (m *messageService) GetMessages(args *GetMessagesArgs, res *models.MessagesResponse) error {
	base, err := url.Parse(m.BaseURL + "/messages")
	if err != nil {
		fmt.Println("parse url.", err)
		return ErrGetMessages
	}
	urlValues := url.Values{}
	urlValues.Add("accesstoken", args.Accesstoken)
	urlValues.Add("mdrender", args.Mdrender)
	base.RawQuery = urlValues.Encode()
	if err = m.HttpClient.Get(base.String(), &res); err != nil {
		fmt.Println(err)
		return ErrGetMessages
	}
	return nil
}

func (m *messageService) GetUnreadMessage(args *GetUnreadMessageArgs, res *int) error {
	base, err := url.Parse(m.BaseURL + "/message/count")
	if err != nil {
		fmt.Println("parse url.", err)
		return ErrGetUnreadMessage
	}
	urlValues := url.Values{}
	urlValues.Add("accesstoken", args.Accesstoken)
	base.RawQuery = urlValues.Encode()
	if err = m.HttpClient.Get(base.String(), &res); err != nil {
		fmt.Println(err)
		return ErrGetUnreadMessage
	}
	return nil
}

func (m *messageService) MarkOneMessage(args *MarkOneMessageArgs, res *string) error {
	endpoint := m.BaseURL + "/message/mark_one/" + args.ID
	var r models.MarkOneMessageResponse
	if err := m.HttpClient.Post(endpoint, &MarkOneMessageRequestPayload{Accesstoken: args.Accesstoken}, &r); err != nil {
		fmt.Println(err)
		return ErrMarkOneMessage
	}
	*res = *r.MarkedMsgId
	return nil
}

func (m *messageService) MarkAllMessages(args *MarkAllMessagesArgs, res *[]models.MarkedMessage) error {
	endpoint := m.BaseURL + "/message/mark_all"
	var r models.MarkAllMessagesResponse
	if err := m.HttpClient.Post(endpoint, &MarkAllMessagesRequestPayload{Accesstoken: args.Accesstoken}, &r); err != nil {
		fmt.Println(err)
		return ErrMarkAllMessages
	}
	*res = *r.MarkedMsgs
	return nil
}
