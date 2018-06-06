
generate:
	remoto generate ./suggestionbox ../../templates/go/remotohttp/server.go.plush -o ./dist/go/server/server.go && go fmt ./dist/go/server/server.go
	remoto generate ./suggestionbox ../../templates/go/remotohttp/client.go.plush -o ./dist/go/client/client.go && go fmt ./dist/go/client/client.go
	remoto generate ./suggestionbox ../../templates/go/cli/cobra-cli.go.plush -o ./dist/go/cli/example/main.go && go fmt ./dist/go/cli/example/main.go

install:
	cd ../../ ; go install