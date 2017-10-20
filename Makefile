protos:
	protoc -I/usr/local/include -I. -I${GOPATH}/src \
		--go_out=plugins=micro:. proto/msgme/service.proto

server: protos
	go run ./server/main.go --registry=mdns

demo: protos
	go run ./demo/main.go