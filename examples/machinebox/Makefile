
generate:
	remoto generate ./suggestionbox ../../templates/go/remotohttp/server.go.plush -o ./suggestionbox/server/suggestionbox/server.go && gofmt -w ./suggestionbox/server/suggestionbox/server.go
	remoto generate ./suggestionbox ../../templates/go/remotohttp/client.go.plush -o ./suggestionbox/client/suggestionbox/client.go && gofmt -w ./suggestionbox/client/suggestionbox/client.go

install:
	cd ../../ ; go install
