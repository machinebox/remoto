serve:
	dev_appserver.py dispatch.yaml www/app.yaml api/app.yaml

build: copyassets
	cd www ; npm run build # see build command in www/package.json

copyassets:
	rm -rf ./api/templates
	cp -R ../templates ./api/templates
