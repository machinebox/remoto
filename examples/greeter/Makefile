
generate:
	remoto generate . ../../templates/go/remotohttp/server.go.plush -o ./server/greeter/server.go && gofmt -w ./server/greeter/server.go
	remoto generate . ../../templates/go/remotohttp/client.go.plush -o ./client/greeter/client.go && gofmt -w ./client/greeter/client.go

install:
	cd ../.. ; go install
