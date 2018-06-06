package facebox

import (
	"context"

	"github.com/machinebox/remoto/remototypes"
)

// Facebox provides facial detection and recognition in images.
type Facebox interface {
	TeachFile(context.Context, *TeachFileRequest) (*TeachFileResponse, error)
	TeachURL(context.Context, *TeachURLRequest) (*TeachURLResponse, error)
	TeachFaceprint(context.Context, *TeachFaceprintRequest) (*TeachFaceprintResponse, error)

	CheckFile(context.Context, *CheckFileRequest) (*CheckFileResponse, error)
	CheckURL(context.Context, *CheckURLRequest) (*CheckURLResponse, error)

	SimilarID(context.Context, *SimilarIDRequest) (*SimilarIDResponse, error)
	SimilarFile(context.Context, *SimilarFileRequest) (*SimilarFileResponse, error)
	SimilarURL(context.Context, *SimilarURLRequest) (*SimilarURLResponse, error)

	Rename(context.Context, *RenameRequest) (*RenameResponse, error)
	RenameID(context.Context, *RenameIDRequest) (*RenameIDResponse, error)

	RemoveID(context.Context, *RemoveIDRequest) (*RemoveIDResponse, error)

	FaceprintCompare(context.Context, *FaceprintCompareRequest) (*FaceprintCompareResponse, error)
	CheckFaceprint(context.Context, *CheckFaceprintRequest) (*CheckFaceprintResponse, error)

	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	PutState(context.Context, *PutStateRequest) (*PutStateResponse, error)
}

type TeachFileRequest struct {
	ID   string
	Name string
	File remototypes.File
}

type TeachFileResponse struct{}

type TeachURLRequest struct {
	ID   string
	Name string
	URL  string
}

type TeachURLResponse struct{}

type TeachFaceprintRequest struct {
	ID        string
	Name      string
	Faceprint string
}

type TeachFaceprintResponse struct{}

type CheckFileRequest struct {
	File remototypes.File
}

type CheckFileResponse struct {
	Faces []Face
}

type CheckURLRequest struct {
	File remototypes.File
}

type CheckURLResponse struct {
	Faces []Face
}

type Face struct {
	ID        string
	Name      string
	Matched   bool
	Faceprint string
	Rect      Rect
}

type Rect struct {
	Top    int
	Left   int
	Width  int
	Height int
}

type SimilarIDRequest struct {
	ID string
}

type SimilarIDResponse struct {
	Faces []SimilarFace
}

type SimilarFileRequest struct {
	File remototypes.File
}

type SimilarFileResponse struct {
	Faces []SimilarFace
}

type SimilarURLRequest struct {
	URL string
}

type SimilarURLResponse struct {
	Faces []SimilarFace
}

type SimilarFace struct {
	Rect         Rect
	SimilarFaces []Face
}

type RenameRequest struct {
	From string
	To   string
}

type RenameResponse struct{}

type RenameIDRequest struct {
	ID   string
	Name string
}

type RenameIDResponse struct{}

type RemoveIDRequest struct {
	ID string
}
type RemoveIDResponse struct{}

type FaceprintCompareRequest struct {
	Target     string
	Faceprints []string
}

type FaceprintCompareResponse struct {
	Confidences []float64
}

type CheckFaceprintRequest struct {
	Faceprints []string
}

type CheckFaceprintResponse struct {
	Faces []FaceprintFace
}

type FaceprintFace struct {
	Matched    bool
	Confidence float64
	ID         string
	Name       string
}

type GetStateRequest struct{}

type GetStateResponse struct {
	StateFile remototypes.File
}

type PutStateRequest struct {
	StateFile remototypes.File
}

type PutStateResponse struct{}
