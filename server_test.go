package main_test

import (
  "fmt"
  "github.com/mrdulin/go-rpc-cnode/modules/user"
  "log"
  "net/rpc"
  "testing"
)

const (
  serverAddress string = "localhost:3000"
)

func TestUser_GetUserByLoginname(t *testing.T) {
  client, err := rpc.DialHTTP("tcp", serverAddress)
  if err != nil {
    log.Fatal("dialing:", err)
  }
  args := user.GetUserByLoginnameArgs{Loginname: "mrdulin"}
  var res user.UserDetail
  err = client.Call("UserService.GetUserByLoginname", args, &res)
  if err != nil {
    log.Fatal("arith error:", err)
  }
  fmt.Printf("res: %+v", res)
} 