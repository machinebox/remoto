
generate:
	remoto generate files.remoto.go ../../templates/remotohttp/server.go.plush -o ./server/files/server.go && gofmt -w ./server/files/server.go
	remoto generate files.remoto.go ../../templates/remotohttp/client.go.plush -o ./client/files/client.go && gofmt -w ./client/files/client.go
	remoto generate files.remoto.go ../../templates/remotohttp/client.es6.js.plush -o ./client/files/files.es6.js 

install:
	cd ../.. ; go install
