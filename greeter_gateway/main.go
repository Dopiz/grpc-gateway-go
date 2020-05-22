package main

import (
  "flag"
  "log"
  "net/http"
   
  "github.com/golang/glog"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"
   	
  gw "examples/helloworld"
)
   
var (
  echoEndpoint = flag.String("echo_endpoint", "localhost:50052", "endpoint of YourService")
)
   
func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()
   
  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}
  err := gw.RegisterGreeterHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
  log.Printf(*echoEndpoint)
  if err != nil {
    return err
  }
   
  return http.ListenAndServe(":8087", mux)
}
   
func main() {
  flag.Parse()
  defer glog.Flush()
   
  if err := run(); err != nil {
    glog.Fatal(err)
  }
}