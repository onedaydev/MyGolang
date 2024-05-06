# GRPC

### user-service
- grpc 서버 기초 예시. .proto 파일부터 server, client 그리고 테스트까지.

### multiple-service
- 여러 서비스 제공하는 gRPC 예시.




protocol compile
```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
users.proto repositories.proto
```