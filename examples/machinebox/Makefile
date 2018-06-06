
generate:
	remoto generate ./suggestionbox ../../templates/go/remotohttp/server.go.plush -o ./dist/go/server/server.go && gofmt -w ./dist/go/server/server.go
	remoto generate ./suggestionbox ../../templates/go/remotohttp/client.go.plush -o ./dist/go/client/client.go && gofmt -w ./dist/go/client/client.go
	remoto generate ./suggestionbox ../../templates/go/cli/cobra-cli.go.plush -o ./dist/go/cli/example/main.go && gofmt -w ./dist/go/cli/example/main.go

install:
	cd ../../ ; go install
