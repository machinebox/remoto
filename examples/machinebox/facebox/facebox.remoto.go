// Package facebox provides access to Facebox Remoto services.
package facebox

import (
	"github.com/machinebox/remoto/remototypes"
)

// Facebox provides facial detection and recognition in images.
type Facebox interface {

	// TeachFile teaches Facebox a new face from an image file.
	TeachFile(*TeachFileRequest) *TeachFileResponse

	// TeachURL teaches Facebox a new face from an image on the web.
	TeachURL(*TeachURLRequest) *TeachURLResponse

	// TeachFaceprint teaches Facebox about a face from a Faceprint.
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

// TeachFileRequest is the request object for TeachFile calls.
type TeachFileRequest struct {
	ID   string
	Name string
	File remototypes.File
}

// TeachFileResponse is the response object for TeachFile calls.
type TeachFileResponse struct{}

// TeachURLRequest is the request object for TeachURL calls.
type TeachURLRequest struct {
	ID   string
	Name string
	URL  string
}

// TeachURLResponse is the response object for TeachURL calls.
type TeachURLResponse struct{}

// TeachFaceprintRequest is the request object for TeachFaceprint calls.
type TeachFaceprintRequest struct {
	ID        string
	Name      string
	Faceprint string
}

// TeachFaceprintResponse is the response object for TeachFaceprint calls.
type TeachFaceprintResponse struct{}

// CheckFileRequest is the request object for CheckFile calls.
type CheckFileRequest struct {
	File remototypes.File
}

// CheckFileResponse is the response object for CheckFile calls.
type CheckFileResponse struct {
	Faces []Face
}

// CheckURLRequest is the request object for CheckURL calls.
type CheckURLRequest struct {
	File remototypes.File
}

// CheckURLResponse is the response object for CheckURL calls.
type CheckURLResponse struct {
	Faces []Face
}

// Face describes a face.
type Face struct {
	ID        string
	Name      string
	Matched   bool
	Faceprint string
	Rect      Rect
}

// Rect is a bounding box describing a rectangle of an image.
type Rect struct {
	Top    int
	Left   int
	Width  int
	Height int
}

// SimilarIDRequest is the request object for SimilarID calls.
type SimilarIDRequest struct {
	ID string
}

// SimilarIDResponse is the response object for SimilarID calls.
type SimilarIDResponse struct {
	Faces []SimilarFace
}

// SimilarFileRequest is the request object for SimilarFile calls.
type SimilarFileRequest struct {
	File remototypes.File
}

// SimilarFileResponse is the response object for SimilarFile calls.
type SimilarFileResponse struct {
	Faces []SimilarFace
}

// SimilarURLRequest is the request object for SimilarURL calls.
type SimilarURLRequest struct {
	URL string
}

// SimilarURLResponse is the response object for SimilarURL calls.
type SimilarURLResponse struct {
	Faces []SimilarFace
}

// SimilarFace is a detected face with similar matching faces.
type SimilarFace struct {
	Rect         Rect
	SimilarFaces []Face
}

// RenameRequest is the request object for Rename calls.
type RenameRequest struct {
	From string
	To   string
}

// RenameResponse is the response object for Rename calls.
type RenameResponse struct{}

// RenameIDRequest is the request object for RenameID calls.
type RenameIDRequest struct {
	ID   string
	Name string
}

// RenameIDResponse is the response object for RenameID calls.
type RenameIDResponse struct{}

// RemoveIDRequest is the request object for RemoveID calls.
type RemoveIDRequest struct {
	ID string
}

// RemoveIDResponse is the response object for RemoveID calls.
type RemoveIDResponse struct{}

// FaceprintCompareRequest is the request object for FaceprintCompare calls.
type FaceprintCompareRequest struct {
	Target     string
	Faceprints []string
}

// FaceprintCompareResponse is the response object for FaceprintCompare calls.
type FaceprintCompareResponse struct {
	Confidences []float64
}

// CheckFaceprintRequest is the request object for CheckFaceprint calls.
type CheckFaceprintRequest struct {
	Faceprints []string
}

// CheckFaceprintResponse is the response object for CheckFaceprint calls.
type CheckFaceprintResponse struct {
	Faces []FaceprintFace
}

// FaceprintFace is a face.
type FaceprintFace struct {
	Matched    bool
	Confidence float64
	ID         string
	Name       string
}

// GetStateRequest is the request object for GetState calls.
type GetStateRequest struct{}

// GetStateResponse is the response object for GetState calls.
type GetStateResponse struct {
	StateFile remototypes.File
}

// PutStateRequest is the request object for PutState calls.
type PutStateRequest struct {
	StateFile remototypes.File
}

// PutStateResponse is the response object for PutState calls.
type PutStateResponse struct{}
