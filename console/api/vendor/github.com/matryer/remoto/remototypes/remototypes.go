package remototypes

// File describes a binary file.
// This type is only allowed in requests, for responses RPC methods should
// return a FileResponse.
type File struct{}

// FileResponse is response type for a file.
type FileResponse struct{}
