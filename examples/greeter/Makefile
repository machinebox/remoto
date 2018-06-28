
generate:
	remoto generate example.remoto.go ../../templates/remotohttp/server.go.plush -o ./server/greeter/server.go && gofmt -w ./server/greeter/server.go
	remoto generate example.remoto.go ../../templates/remotohttp/client.go.plush -o ./client/greeter/client.go && gofmt -w ./client/greeter/client.go

install:
	cd ../.. ; go install
