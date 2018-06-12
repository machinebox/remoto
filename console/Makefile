serve:
	dev_appserver.py dispatch.yaml www/app.yaml api/app.yaml

build:
	cd www ; npm run build # see command in www/package.json

copytemplates:
	rm -rf ./api/templates
	cp -R ../templates ./api/templates
