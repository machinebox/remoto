# Writing templates

Remoto templates are how code is generated. Templates are written using [Plush](https://github.com/gobuffalo/plush).

### Basics

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

In Go you might have template code that looks like this:

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

### Dealing with repsonse files

Some endpoints can return a file, these have the `*remototypes.FileResponse` return type.

Plush templates can check if a method has this return type:

```c
// general code

<%= if (method.ResponseType.Name == "remototypes.FileResponse") { %>
	// file specific response code
<% } else { %>
	// Data structure response code
<% } %>
```
