
all:
	remoto generate ./suggestionbox/suggestionbox.remoto.go ../../templates/remotohttp/server.go.plush -o ./suggestionbox/server/suggestionbox/server.go && gofmt -w ./suggestionbox/server/suggestionbox/server.go
	remoto generate ./suggestionbox/suggestionbox.remoto.go ../../templates/remotohttp/client.go.plush -o ./suggestionbox/client/suggestionbox/client.go && gofmt -w ./suggestionbox/client/suggestionbox/client.go
	remoto generate ./facebox/facebox.remoto.go ../../templates/remotohttp/server.go.plush -o ./facebox/server/facebox/server.go && gofmt -w ./facebox/server/facebox/server.go
	remoto generate ./facebox/facebox.remoto.go ../../templates/remotohttp/client.go.plush -o ./facebox/client/facebox/client.go && gofmt -w ./facebox/client/facebox/client.go
	remoto generate ./facebox/facebox.remoto.go ../../templates/remotohttp/client.jquery.js.plush -o ./facebox/client/web/facebox.jquery.js 
	remoto generate ./facebox/facebox.remoto.go ../../templates/html/docs.bootstrap.html.plush -o ./facebox/client/web/docs.bootstrap.html 
	remoto generate ./facebox/facebox.remoto.go ../../templates/remotohttp/client.es6.js.plush -o ./facebox/client/web/facebox.es6.js 
	cd ./facebox/client/web && npx babel ./facebox.es6.js --out-file ./dist/facebox.js
	cd ./facebox/client/web && tidy -quiet -utf8 -i -w 160 -o ./dist/docs.html docs.bootstrap.html

install:
	cd ../../ ; go install
