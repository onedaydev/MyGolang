module github.com/onedaydev/mygolang/grpc/multiple-services/server

go 1.22.1

require github.com/onedaydev/mygolang/grpc/multiple-services/service v0.0.0

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20190311183353-d8887717615a // indirect
	golang.org/x/sys v0.0.0-20190215142949-d0b11bdaac8a // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/grpc v1.37.0 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/onedaydev/mygolang/grpc/multiple-services/service => ../service