# grpc-gateway-go

- Environment Variables

```bash
export GOBIN=$(go env GOPATH)/bin
export PATH=$PATH:$(go env GOPATH):$(go env GOBIN)
```

- Generate Server pb

```bash
$ cd helloworld
$ protoc -I. --go_out=plugins=grpc,paths=source_relative:. helloworld.proto
```

- Generate Gateway pb

```bash
$ cd helloworld
$ protoc -I. --grpc-gateway_out=logtostderr=true,paths=source_relative:. helloworld.proto
```

- Run Server

```bash
$ go run ./greeter_server/main.go
```

- Run Gateway Server

```bash
$ go run ./greeter_gateway/main.go
```
