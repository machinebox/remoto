package files

import (
	"github.com/machinebox/remoto/remototypes"
)

// Images provides image services.
type Images interface {
	Flip(FlipRequest) remototypes.FileResponse
}

// FlipRequest is the request for Images.Flip.
type FlipRequest struct {
	Image remototypes.File
}
