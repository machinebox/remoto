package example

import (
	"context"

	"github.com/machinebox/remoto/remototypes"
)

// Facebox provides facial detection and recognition capabilities.
type Facebox interface {
	Teach(context.Context, *TeachRequest) (*TeachResponse, error)
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

type TeachRequest struct {
	Name       string
	TeachFiles []TeachFile
}

type TeachFile struct {
	Image remototypes.File
}

type TeachResponse struct {
}

type CheckRequest struct {
	Image remototypes.File
}

type CheckResponse struct {
	Faces []Faces
}

type Faces struct {
	Name    string
	Matched bool
}
