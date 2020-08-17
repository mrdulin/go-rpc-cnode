package main

import (
  "github.com/mrdulin/go-rpc-cnode/modules/user"
  api "github.com/mrdulin/go-rpc-cnode/utils/http"
  "log"
  "net"
  "net/http"
  "net/rpc"
  
)

const (
  port string = "3000"
  baseurl string = "https://cnodejs.org/api/v1"
)

func main() {
  httpClient := api.NewClient()
  userService := user.Service{httpClient, baseurl}
  e := rpc.RegisterName("UserService", &userService)
  if e != nil {
    log.Fatal("RegisterName error:", e) 
  }
  rpc.HandleHTTP()
  l, e := net.Listen("tcp", ":" + port)
  if e != nil {
    log.Fatal("listen error:", e)
  }
  log.Fatalln(http.Serve(l, nil))
}

