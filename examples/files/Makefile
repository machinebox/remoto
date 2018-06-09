
generate:
	remoto generate . ../../templates/remotohttp/server.go.plush -o ./server/files/server.go && gofmt -w ./server/files/server.go
	remoto generate . ../../templates/remotohttp/client.go.plush -o ./client/files/client.go && gofmt -w ./client/files/client.go
	remoto generate . ../../templates/remotohttp/client.es6.js.plush -o ./client/files/files.es6.js 

install:
	cd ../.. ; go install
