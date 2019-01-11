# remototypes

The `remototypes` package contains specially handled data types.

## Files

Remoto provides a simple solution for returning a file (download), and submitting
many files in a request (upload).

### Return a file

To return a file, use `*remototypes.FileResponse` as the return type instead of a custom structure of your own.

The `GetCatPic` method will return a file.

### Submitting files

The `remototypes.File` type indicates a file to upload. It is versatile enough to be used
just like a normal type in your request structures, in arrays or sub-structures, or anywhere
that is valid in Go.

### Example code for files

The following `cat.remoto.go` file demonstrates how to both upload and download files.

```go
package cats

import (
	"github.com/machinebox/remoto/remototypes"
)

// CatService provides cat picture services.
type CatService interface {
	// GetCatPic gets a random cat picture.
	GetCatPic(*GetCatPicRequest) *remototypes.FileResponse
	// SubmitCatPic submits one or more cat pictures for consideration.
	SubmitCatPic(*SubmitCatPicRequest) *SubmitCatPicResponse
}

type SubmitCatPicRequest struct {
	// Pics are the pictures to submit.
	Pics []CatPic
}

// CatPic is a picture of a cat.
type CatPic struct {
	// Caption is a funny quip about the image.
	Caption string
	// Image is the image file.
	Image remototypes.File
}

type SubmitCatPicResponse struct {
	// TotalCatPics is the total number of cat pics
	// in the service.
	TotalCatPics int
}
```
