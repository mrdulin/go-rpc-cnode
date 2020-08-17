package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/mrdulin/go-rpc-cnode/models"
	"github.com/mrdulin/go-rpc-cnode/services"
)

const (
	serverAddress string = "localhost:3000"
)

func testGetUserByLoginname(rpcclient *rpc.Client) {
	args := services.GetUserByLoginnameArgs{Loginname: "mrdulin"}
	var res models.UserDetail
	err := rpcclient.Call("UserService.GetUserByLoginname", args, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("testGetUserByLoginname: %+v\n", res)
}

func testGetTopicsByPage(rpcclient *rpc.Client) {
	args := services.GetTopicsByPageArgs{Page: 1, Limit: 2, Mdrender: "false"}
	var res []models.Topic
	err := rpcclient.Call("TopicService.GetTopicsByPage", args, &res)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("testGetTopicsByPage: %+v\n", res)
}

func main() {
	client, err := jsonrpc.Dial("tcp", serverAddress)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	testGetUserByLoginname(client)
	testGetTopicsByPage(client)
}
