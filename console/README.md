# Remoto Console

The Remote Console is an online application for writing Remoto definition files, and generating documentation, client libraries and SDKs.

## API

### Get templates

Get a list of templates:

```
GET /api/templates
```

Returns:

```json
{
	"templates": [{
		"name": "html/docs.bootstrap.html",
		"x": false,
		"dirs": ["html"],
		"label": "docs.bootstrap.html"
	}, {
		"name": "remotohttp/client.es6.js",
		"x": false,
		"dirs": ["remotohttp"],
		"label": "client.es6.js"
	}, {
		"name": "remotohttp/client.go",
		"x": false,
		"dirs": ["remotohttp"],
		"label": "client.go"
	}, {
		"name": "remotohttp/client.jquery.js",
		"x": false,
		"dirs": ["remotohttp"],
		"label": "client.jquery.js"
	}, {
		"name": "remotohttp/server.go",
		"x": false,
		"dirs": ["remotohttp"],
		"label": "server.go"
	}, {
		"name": "x/go/cli/cobra-cli.go",
		"x": true,
		"dirs": ["x", "go", "cli"],
		"label": "cobra-cli.go"
	}]
}
```

## Generate code from a template

```
POST /api/templates/{name}
definition=<remoto defintion>
```

* `{name}` - (string) The template name (looks like a path) to use to generate code
* `definition` - (string) The Remoto definition file to use to generaate the template

## Generare all code

You can use the helpful shortcut to generate a `.zip` file containing all generated code:

```
POST /api/all.zip
definition=<remoto defintion>
```

* `definition` - (string) The Remoto definition file to use to generaate the template

Downloads a ZIP file containing all generated code.
