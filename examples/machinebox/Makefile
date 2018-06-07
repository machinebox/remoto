
generate:
	remoto generate ./suggestionbox ../../templates/go/remotohttp/server.go.plush -o ./suggestionbox/server/suggestionbox/server.go && gofmt -w ./suggestionbox/server/suggestionbox/server.go
	remoto generate ./suggestionbox ../../templates/go/remotohttp/client.go.plush -o ./suggestionbox/client/suggestionbox/client.go && gofmt -w ./suggestionbox/client/suggestionbox/client.go
	remoto generate ./facebox ../../templates/go/remotohttp/server.go.plush -o ./facebox/server/facebox/server.go && gofmt -w ./facebox/server/facebox/server.go
	remoto generate ./facebox ../../templates/go/remotohttp/client.go.plush -o ./facebox/client/facebox/client.go && gofmt -w ./facebox/client/facebox/client.go

install:
	cd ../../ ; go install
