
generate:
	remoto generate . ../../templates/go/remotohttp/server.go.plush -o ./server/files/server.go && gofmt -w ./server/files/server.go
	remoto generate . ../../templates/go/remotohttp/client.go.plush -o ./client/files/client.go && gofmt -w ./client/files/client.go

install:
	cd ../.. ; go install
