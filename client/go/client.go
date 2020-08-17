package main

import (
  "fmt"
  "github.com/mrdulin/go-rpc-cnode/modules/user"
  "log"
  "net/rpc/jsonrpc"
  "reflect"
)

const (
  serverAddress string = "localhost:3000"
)

func main() {
  client, err := jsonrpc.Dial("tcp", serverAddress)
  if err != nil {
    log.Fatal("dialing:", err)
  }
  args := user.GetUserByLoginnameArgs{Loginname: "mrdulin"}
  var res user.UserDetail
  err = client.Call("UserService.GetUserByLoginname", args, &res)
  if err != nil {
    log.Fatal("arith error:", err)
  }
  fmt.Printf("res: %+v, type: %+v", res, reflect.TypeOf(res).Kind())
} 