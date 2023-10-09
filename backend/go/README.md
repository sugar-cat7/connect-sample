# Connect-Go

https://connectrpc.com/docs/go/getting-started

- proto からコード生成
  - `protoc` or `protoc-gen-connect-go`
  - `buf`を使用して lint とコード生成を行う

ex)サーバーを起動してリクエストを送る

1. HTTP/1.1 POST

```
curl \
    --header "Content-Type: application/json" \
    --data '{"name": "Jane"}' \
    http://localhost:8080/proto.v1.GreetService/Greet

2023/10/09 17:35:39 Request headers: map[Accept:[*/*] Content-Length:[16] Content-Type:[application/json] Host:[localhost:8080] User-Agent:[curl/8.1.2]]
```

2. gRPC

```
grpcurl \
    -protoset <(buf build -o -) -plaintext \
    -d '{"name": "Jane"}' \
    localhost:8080 proto.v1.GreetService/Greet

2023/10/09 17:37:57 Request headers:  map[Content-Type:[application/grpc] Grpc-Accept-Encoding:[gzip] Host:[localhost:8080] Te:[trailers] User-Agent:[grpcurl/dev-build (no version set) grpc-go/1.57.0]]
2.gRPC requests
```

1. Connect's generated client

```
go run ./cmd/client/main.go
```

## Serverless 環境の gRPC 対応状況

https://developers.cloudflare.com/support/other-languages/%E6%97%A5%E6%9C%AC%E8%AA%9E/cloudflare%E3%81%AEgrpc%E3%82%B5%E3%83%9D%E3%83%BC%E3%83%88%E3%82%92%E7%90%86%E8%A7%A3%E3%81%99%E3%82%8B/
