package files

import (
	"context"

	"github.com/machinebox/remoto/remototypes"
)

// Images provides image services.
type Images interface {
	Flip(context.Context, *FlipRequest) (*FlipResponse, error)
}

// FlipRequest is the request for Images.Flip.
type FlipRequest struct {
	Image remototypes.File
}

// FlipResponse is the response for Images.Flip.
type FlipResponse struct {
	FlippedImage remototypes.File
}
