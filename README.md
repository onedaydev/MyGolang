# MyGolang
## gRPC
- grpc 폴더 내 README.md 참고

## Cmd Command
### cmd
- flag모듈 사용한 cmd 명령어 사용 예시. 
- test 파일 포함.

### data-downloader
- cmd에서 명령어로 특정 url body 읽어오기.

### goroutine_example
- cmd에서 명령어로 파일 내 특정 단어 검색하는 모듈.

### sub_cmd
- cmd 명령어 예시.
- test 포함.

### pkgquery
- json 데이터 언마셜링 예시.

### pkgregister
- json 요청 응답 예시.
- encoding/json의 marshal, unmarshal 과정.
- 테스트 포함.

## http
### client-slow-write 
- io.Pipe()를 이용한 느린 요청 클라이언트 예시.

### complex-server
- http 서버를 복합적으로 사용한 예시. 
- config, handlers, middleware 모듈로 분류해서 사용함. 
- test 파일 포함.

### graceful-shutdown
- shutdown 고루틴 생성.
- context와 Channel 활용해서 모든 서비스 종료 후에 서버 종료.

### handle_func_timeout
- http.TimeoutHandler 사용 예시.

### http-handler-type
- Config 값을 사용한 설정이 필요한 Handler를 다루는 예시.

### http-serve-mux
- 멀티플렉서 사용 예시.

### logging-middleware
- 로깅 미들웨어 예시.

### middleware-chaining
- 미들웨어 체이닝 예시.

### server-timeout
- http.Server 커스텀(ReadTimeout, WriteTimeout) 예시.

### streaming-response
- io.Pipe() 사용한 streaming-response 예시.

### tls-server
- tls 서버 예시.
- tls 인증서 생성 과정 생략됨.

### todo
- todo 백엔드.

### user-input-timeout
- context.WithTimeout 이용 예시.

### user-signal
- SIGINT(일반적으로 Ctrl+C에 의해 발생) 또는 SIGTERM 시그널 발생 시 종료하는 예시.
