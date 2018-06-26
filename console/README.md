# Remoto Console

The Remote Console is an online application for writing Remoto definition files, and generating documentation, client libraries and SDKs.

## API

### Parse a definition

To parse a Remoto definition file, post the source to the following endpoint:

```
POST /api/define
definition=...
```

* `definition` field should contain the source

Response will be a JSON object containing the definition or an error message if something went wrong.

Unsuccessful response:

```json
{
	"ok": false,
	"error": "io.Reader.go:24:18: expected '}', found 'EOF'"
}
```

Successful response:

```json
{
	"ok": true,
	"definition": {
		"services": [{
			"name": "Greeter",
			"comment": "Greeter is a friendly service.",
			"methods": [{
				"name": "Greet",
				"comment": "Greet makes a greeting.",
				"requestStructure": {
					"name": "GreetRequest",
					"comment": "GreetRequest is the request object for Greeter.Greet.",
					"fields": [{
						"name": "Name",
						"comment": "Name is the person to greet.",
						"type": {
							"name": "string",
							"isMultiple": false,
							"isStruct": false,
							"isImported": false
						}
					}],
					"isImported": false,
					"isRequestObject": true,
					"isResponseObject": false
				},
				"responseStructure": {
					"name": "GreetResponse",
					"comment": "GreetResponse is the response object for Greeter.Greet.",
					"fields": [{
						"name": "Greeting",
						"comment": "Greeting is a personalized message",
						"type": {
							"name": "string",
							"isMultiple": false,
							"isStruct": false,
							"isImported": false
						}
					}, {
						"name": "Error",
						"comment": "Error is an error message if one occurred.",
						"type": {
							"name": "string",
							"isMultiple": false,
							"isStruct": false,
							"isImported": false
						}
					}],
					"isImported": false,
					"isRequestObject": false,
					"isResponseObject": true
				}
			}],
			"structures": [{
				"name": "GreetRequest",
				"comment": "GreetRequest is the request object for Greeter.Greet.",
				"fields": [{
					"name": "Name",
					"comment": "Name is the person to greet.",
					"type": {
						"name": "string",
						"isMultiple": false,
						"isStruct": false,
						"isImported": false
					}
				}],
				"isImported": false,
				"isRequestObject": true,
				"isResponseObject": false
			}, {
				"name": "GreetResponse",
				"comment": "GreetResponse is the response object for Greeter.Greet.",
				"fields": [{
					"name": "Greeting",
					"comment": "Greeting is a personalized message",
					"type": {
						"name": "string",
						"isMultiple": false,
						"isStruct": false,
						"isImported": false
					}
				}, {
					"name": "Error",
					"comment": "Error is an error message if one occurred.",
					"type": {
						"name": "string",
						"isMultiple": false,
						"isStruct": false,
						"isImported": false
					}
				}],
				"isImported": false,
				"isRequestObject": false,
				"isResponseObject": true
			}]
		}],
		"packageName": "remoto",
		"packageComment": "Remoto service is an example service."
	}
}
```

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
