
generate:
	remoto generate ./definition ../templates/go/remotohttpjson/server.go.plush -o ./go/server/server.go
	remoto generate ./definition ../templates/go/remotohttpjson/client.go.plush -o ./go/client/client.go
	remoto generate ./definition ../templates/js/remotohttpjson/jquery3-client.js.plush -o ./js/jquery-client.js
	remoto generate ./definition ../templates/js/remotohttpjson/jquery3-client-ui.html.plush -o ./js/index.html

install:
	cd .. ; go install
