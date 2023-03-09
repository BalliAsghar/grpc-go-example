# Generate go code from protobuf files

```bash
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
product/product.proto
```

## Run server

```bash
go run main.go
```
