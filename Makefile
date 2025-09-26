.PHONY: server client 

server:
	@go run ./znet/cmd/main.go

client:
	@go run ./client/client.go