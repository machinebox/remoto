package files

import (
	"context"

	"github.com/machinebox/remoto/remototypes"
)

// Images provides image services.
type Images interface {
	Flip(context.Context, *FlipRequest) (*remototypes.FileResponse, error)
}

// FlipRequest is the request for Images.Flip.
type FlipRequest struct {
	Image remototypes.File
}
