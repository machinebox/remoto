// Package facebox provides access to Facebox Remoto services.
package facebox

import (
	"github.com/machinebox/remoto/remototypes"
)

// Facebox provides facial detection and recognition in images.
type Facebox interface {
	TeachFile(*TeachFileRequest) *TeachFileResponse
	TeachURL(*TeachURLRequest) *TeachURLResponse
	TeachFaceprint(*TeachFaceprintRequest) *TeachFaceprintResponse

	CheckFile(*CheckFileRequest) *CheckFileResponse
	CheckURL(*CheckURLRequest) *CheckURLResponse

	SimilarID(*SimilarIDRequest) *SimilarIDResponse
	SimilarFile(*SimilarFileRequest) *SimilarFileResponse
	SimilarURL(*SimilarURLRequest) *SimilarURLResponse

	Rename(*RenameRequest) *RenameResponse
	RenameID(*RenameIDRequest) *RenameIDResponse

	RemoveID(*RemoveIDRequest) *RemoveIDResponse

	FaceprintCompare(*FaceprintCompareRequest) *FaceprintCompareResponse
	CheckFaceprint(*CheckFaceprintRequest) *CheckFaceprintResponse

	GetState(*GetStateRequest) *remototypes.FileResponse
	PutState(*PutStateRequest) *PutStateResponse
}

// Suggestionbox provides facial detection and recognition in images.
type Suggestionbox interface {
	TeachFile(*TeachFileRequest) *TeachFileResponse
	TeachURL(*TeachURLRequest) *TeachURLResponse
	TeachFaceprint(*TeachFaceprintRequest) *TeachFaceprintResponse

	CheckFile(*CheckFileRequest) *CheckFileResponse
	CheckURL(*CheckURLRequest) *CheckURLResponse

	SimilarID(*SimilarIDRequest) *SimilarIDResponse
	SimilarFile(*SimilarFileRequest) *SimilarFileResponse
	SimilarURL(*SimilarURLRequest) *SimilarURLResponse

	Rename(*RenameRequest) *RenameResponse
	RenameID(*RenameIDRequest) *RenameIDResponse

	RemoveID(*RemoveIDRequest) *RemoveIDResponse

	FaceprintCompare(*FaceprintCompareRequest) *FaceprintCompareResponse
	CheckFaceprint(*CheckFaceprintRequest) *CheckFaceprintResponse

	GetState(*GetStateRequest) *remototypes.FileResponse
	PutState(*PutStateRequest) *PutStateResponse
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
