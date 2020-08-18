package services_test

import (
	"errors"
	"reflect"
	"strconv"
	"testing"

	"github.com/mrdulin/go-rpc-cnode/mocks"
	"github.com/mrdulin/go-rpc-cnode/models"
	"github.com/mrdulin/go-rpc-cnode/services"
	"github.com/stretchr/testify/mock"
)

const (
	baseurl     string = "http://localhost:3000"
	accesstoken string = "123"
)

func TestMessageService_MarkOneMessage(t *testing.T) {
	markedMsgId := "666"

	t.Run("should mark one message", func(t *testing.T) {
		testHttp := new(mocks.MockedHttp)
		var r models.MarkOneMessageResponse
		var res string
		args := services.MarkOneMessageArgs{ID: "1", Accesstoken: accesstoken}
		testHttp.
			On("Post", baseurl+"/message/mark_one/"+args.ID, &services.MarkOneMessageRequestPayload{Accesstoken: accesstoken}, &r).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(2).(*models.MarkOneMessageResponse)
				arg.MarkedMsgId = &markedMsgId
			})
		svc := services.NewMessageService(testHttp, baseurl)
		err := svc.MarkOneMessage(&args, &res)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(res, markedMsgId) {
			t.Errorf("got %#v, want: %#v", res, markedMsgId)
		}
	})
}

func TestMessageService_MarkAllMessages(t *testing.T) {
	var testMarkedMessage []models.MarkedMessage
	for i := 1; i <= 3; i++ {
		testMarkedMessage = append(testMarkedMessage, models.MarkedMessage{ID: strconv.Itoa(i)})
	}
	t.Run("should mark all messages", func(t *testing.T) {
		testHttp := new(mocks.MockedHttp)
		var r models.MarkAllMessagesResponse
		var res []models.MarkedMessage
		args := services.MarkAllMessagesArgs{Accesstoken: accesstoken}
		testHttp.
			On("Post", baseurl+"/message/mark_all", &services.MarkAllMessagesRequestPayload{Accesstoken: args.Accesstoken}, &r).
			Return(nil).
			Run(func(args mock.Arguments) {
				arg := args.Get(2).(*models.MarkAllMessagesResponse)
				arg.MarkedMsgs = &testMarkedMessage
			})
		svc := services.NewMessageService(testHttp, baseurl)
		err := svc.MarkAllMessages(&args, &res)
		testHttp.AssertExpectations(t)
		if !reflect.DeepEqual(err, nil) {
			t.Error(err)
		}
		if !reflect.DeepEqual(res, testMarkedMessage) {
			t.Errorf("got: %+v, want: %+v", res, testMarkedMessage)
		}
	})

	t.Run("should return error", func(t *testing.T) {
		testHttp := new(mocks.MockedHttp)
		var r models.MarkAllMessagesResponse
		var res []models.MarkedMessage
		args := services.MarkAllMessagesArgs{Accesstoken: accesstoken}
		testHttp.
			On("Post", baseurl+"/message/mark_all", &services.MarkAllMessagesRequestPayload{Accesstoken: args.Accesstoken}, &r).
			Return(errors.New("network"))
		svc := services.NewMessageService(testHttp, baseurl)
		err := svc.MarkAllMessages(&args, &res)
		testHttp.AssertExpectations(t)
		if !reflect.DeepEqual(err, services.ErrMarkAllMessages) {
			t.Errorf("got: %+v, want: %+v", err, services.ErrMarkAllMessages)
		}
	})
}
