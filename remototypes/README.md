# remototypes

The `remototypes` package contains specially handled data types.

## Files

Remoto provides a simple solution for returning a file (download), and submitting
many files in a request (upload).

### Return a file

To return a file, use the `*remototypes.FileResponse` type, instead of a custom structure.

### Submitting files

The `remototypes.File` type indicates a file to upload. It is versatile enough to be used
just like a normal type in your request structures, in arrays or sub-structures, or anywhere
that is valid in Go.
