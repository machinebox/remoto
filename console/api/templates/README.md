# Remoto templates

* [Introduction](#introduction)
* [Writing templates](#writing-templates)

## Introduction

Remoto generates code by processing a [Plush](https://github.com/gobuffalo/plush) template with a data structure describing the services.

A goal of the project is to keep and maintain all Remoto templates in this repository, although it is
possible to render with any template for advanced cases, the wider community will benefit from a
carefully maintained experience.

# Writing templates

Remoto templates are used to generate code; clients, server stubs, SDKs, etc.

### Introduction

You will generate code for one or more services, accessible via `def.Services`:

```c
<%= for (service) in def.Services { %>
	// rendered for each service
	<%= for (method) in service.Methods %>
		// rendered for each method
	<% } %>
<% } %>
```

You will also generate code for the list of structures, and the `unique_structures` helper 
gives you the complete list:

```c
<%= for (structure) in unique_structures(def) { %>
	// rendered for each structure
<% } %>
```

The structure comes with a list of fields.

Templates aren't the nicest of things to look at and work with, but the pain here means
we can generate human-readable code for our users.

For example, in Go you might have template code that looks like this:

```go
<%= for (structure) in unique_structures(def) { %>
<%= print_comment(structure.Comment) %>type <%= structure.Name %> struct {
	<%= for (field) in structure.Fields { %>
	<%= print_comment(field.Comment) %><%= field.Name %> <%= go_type_string(field.Type) %> `json:"<%= underscore(field.Name) %>"`
	<% } %>
}
<% } %>
```

The above template will generate a Go `struct` matching the data structure of the object.

### Templating language

Templates are written using [Plush](https://github.com/gobuffalo/plush), a templating package
from the [Buffalo project](https://gobuffalo.io/).

Plush uses `<%= tagsLikeThese %>` to inject data and provide conditional output and loops.

* See [Usage of the Plush Package](https://github.com/gobuffalo/plush#usage)

### Template helpers

Remoto inherits all of the [Plush helpers](https://github.com/gobuffalo/plush#helpers) and adds some
specific ones in the [generator/template_helpers.go](https://github.com/machinebox/remoto/blob/master/generator/template_helpers.go) file.

### Template data structure

The data structure for the templates is best expressed through the godoc online documentation:

* [https://godoc.org/github.com/machinebox/remoto/generator/definition](https://godoc.org/github.com/machinebox/remoto/generator/definition)

### Dealing with repsonse files

Some endpoints can return a file, these have the `*remototypes.FileResponse` return type.

Plush templates can check if a method has this return type:

```c
// general code

<%= if (method.ResponseStructure.Name == "remototypes.FileResponse") { %>
	// file specific response code
<% } else { %>
	// Data structure response code
<% } %>
```
