
generate:
	remoto generate ./definition ../../templates/go/remotohttp/server.go.plush -o ./dist/go/server/server.go && go fmt ./dist/go/server/server.go
	remoto generate ./definition ../../templates/go/remotohttp/client.go.plush -o ./dist/go/client/client.go && go fmt ./dist/go/client/client.go
	remoto generate ./definition ../../templates/go/cli/cobra-cli.go.plush -o ./dist/go/cli/greeter/main.go && go fmt ./dist/go/cli/greeter/main.go
	remoto generate ./definition ../../templates/js/remotohttp/jquery3-client.js.plush -o ./dist/js/jquery-client.js
	remoto generate ./definition ../../templates/js/remotohttp/jquery3-client-ui.html.plush -o ./dist/js/index.html

install:
	cd ../.. ; go install
