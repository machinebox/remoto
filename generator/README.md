# Remoto generators

Remoto uses [Plush](https://github.com/gobuffalo/plush) templates to generate code.

## Built in functions

As well as the [Plush built-in helpers](https://github.com/gobuffalo/plush#builtin-helpers), Remoto also provides:

* `unique_structures(definition)` - Get a list of all structures in the entire definition
* `has_field_type(definition|structure, type)` - Gets whether any field in the definition or structure matches the specified type (useful for determining if file support is required)
