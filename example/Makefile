
generate:
	remoto generate ./definition ../templates/go/remotohttp/server.go.plush -o ./go/server/server.go && go fmt ./go/client/client.go
	remoto generate ./definition ../templates/go/remotohttp/client.go.plush -o ./go/client/client.go && go fmt ./go/client/client.go
	remoto generate ./definition ../templates/go/cli/cobra-cli.go.plush -o ./go/cli/example/main.go && go fmt ./go/cli/example/main.go
	#remoto generate ./definition ../templates/js/remotohttp/jquery3-client.js.plush -o ./js/jquery-client.js
	#remoto generate ./definition ../templates/js/remotohttp/jquery3-client-ui.html.plush -o ./js/index.html

install:
	cd .. ; go install
