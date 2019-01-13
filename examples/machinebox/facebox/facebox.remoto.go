// Package machinebox provides access to Facebox Remoto services.
package machinebox

import (
	"github.com/machinebox/remoto/remototypes"
)

// Facebox provides facial detection and recognition in images.
type Facebox interface {

	// TeachFile teaches Facebox a new face from an image file.
	TeachFile(TeachFileRequest) TeachFileResponse

	// TeachURL teaches Facebox a new face from an image on the web.
	TeachURL(TeachURLRequest) TeachURLResponse

	// TeachFaceprint teaches Facebox about a face from a Faceprint.
	TeachFaceprint(TeachFaceprintRequest) TeachFaceprintResponse

	// CheckFile checks an image file for faces.
	CheckFile(CheckFileRequest) CheckFileResponse

	// CheckURL checks a hosted image file for faces.
	CheckURL(CheckURLRequest) CheckURLResponse

	// CheckFaceprint checks to see if a Faceprint matches any known
	// faces.
	CheckFaceprint(CheckFaceprintRequest) CheckFaceprintResponse

	// SimilarID checks for similar faces by ID.
	SimilarID(SimilarIDRequest) SimilarIDResponse

	// SimilarFile checks for similar faces from the face in an image file.
	SimilarFile(SimilarFileRequest) SimilarFileResponse

	// SimilarURL checks for similar faces in a hosted image file.
	SimilarURL(SimilarURLRequest) SimilarURLResponse

	// Rename changes a person's name.
	Rename(RenameRequest) RenameResponse

	// RenameID changes the name of a previously taught face, by ID.
	RenameID(RenameIDRequest) RenameIDResponse

	// RemoveID removes a face with the specified ID.
	RemoveID(RemoveIDRequest) RemoveIDResponse

	// FaceprintCompare compares faceprints to a specified target describing
	// similarity.
	FaceprintCompare(FaceprintCompareRequest) FaceprintCompareResponse

	// GetState gets the Facebox state file.
	GetState(GetStateRequest) remototypes.FileResponse

	// PutState sets the Facebox state file.
	PutState(PutStateRequest) PutStateResponse
}

// TeachFileRequest is the request object for TeachFile calls.
type TeachFileRequest struct {
	// ID is an identifier describing the source, for example the filename.
	ID string
	// Name is the name of the person in the image.
	Name string
	// File is the image containing the face to teach.
	File remototypes.File
}

// TeachFileResponse is the response object for TeachFile calls.
type TeachFileResponse struct{}

// TeachURLRequest is the request object for TeachURL calls.
type TeachURLRequest struct {
	// ID is an identifier describing the source, for example the filename.
	ID string
	// Name is the name of the person in the image.
	Name string
	// URL is the address of the image.
	URL string
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
	// File is the image to check for faces.
	File remototypes.File
}

// CheckFileResponse is the response object for CheckFile calls.
type CheckFileResponse struct {
	// Faces is a list of faces that were found.
	Faces []Face
}

// CheckURLRequest is the request object for CheckURL calls.
type CheckURLRequest struct {
	// URL is the address of the image to check.
	URL string
}

// CheckURLResponse is the response object for CheckURL calls.
type CheckURLResponse struct {
	// Faces is a list of faces that were found.
	Faces []Face
}

// Face describes a face.
type Face struct {
	// ID is the identifier of the source that was matched.
	ID string
	// Name is the name of the identified person.
	Name string
	// Matched is whether the face was recognized or not.
	Matched bool
	// Faceprint is the Facebox Faceprint of this face.
	Faceprint string
	// Rect is where the face appears in the source image.
	Rect Rect
}

// Rect is a bounding box describing a rectangle of an image.
type Rect struct {
	// Top is the starting Y coordinate.
	Top int
	// Left is the starting X coordinate.
	Left int
	// Width is the width.
	Width int
	// Height is the height.
	Height int
}

// SimilarIDRequest is the request object for SimilarID calls.
type SimilarIDRequest struct {
	// ID is the identifier of the source to look for similar faces of.
	ID string
}

// SimilarIDResponse is the response object for SimilarID calls.
type SimilarIDResponse struct {
	// Faces is a list of similar faces.
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
	// Rect is where the face appears in the image.
	Rect Rect
	// SimilarFaces is a list of similar faces.
	SimilarFaces []Face
}

// RenameRequest is the request object for Rename calls.
type RenameRequest struct {
	// From is the original name.
	From string
	// To is the new name.
	To string
}

// RenameResponse is the response object for Rename calls.
type RenameResponse struct{}

// RenameIDRequest is the request object for RenameID calls.
type RenameIDRequest struct {
	// ID is the identifier of the source to rename.
	ID string
	// Name is the new name to assign to the item matching ID.
	Name string
}

// RenameIDResponse is the response object for RenameID calls.
type RenameIDResponse struct{}

// RemoveIDRequest is the request object for RemoveID calls.
type RemoveIDRequest struct {
	// ID is the identifier of the source to remove.
	ID string
}

// RemoveIDResponse is the response object for RemoveID calls.
type RemoveIDResponse struct{}

// FaceprintCompareRequest is the request object for FaceprintCompare calls.
type FaceprintCompareRequest struct {
	// Target is the target Faceprint to which the Faceprints will be compared.
	Target string
	// Faceprints is a list of Faceprints that will be compared to Target.
	Faceprints []string
}

// FaceprintCompareResponse is the response object for FaceprintCompare calls.
type FaceprintCompareResponse struct {
	// Confidences is a list of confidence values.
	// The order matches the order of FaceprintCompareRequest.Faceprints.
	Confidences []float64
}

// CheckFaceprintRequest is the request object for CheckFaceprint calls.
type CheckFaceprintRequest struct {
	// Faceprints is a list of Faceprints to check.
	Faceprints []string
}

// CheckFaceprintResponse is the response object for CheckFaceprint calls.
type CheckFaceprintResponse struct {
	// Faces is a list of faces checked from Faceprints.
	Faces []FaceprintFace
}

// FaceprintFace is a face.
type FaceprintFace struct {
	// Matched is whether the face was recognized or not.
	Matched bool
	// Confidence is a numerical value of how confident the AI
	// is that this is a match.
	Confidence float64
	// ID is the identifier of the source that matched.
	ID string
	// Name is the name of the person recognized.
	Name string
}

// GetStateRequest is the request object for GetState calls.
type GetStateRequest struct{}

// PutStateRequest is the request object for PutState calls.
type PutStateRequest struct {
	// StateFile is the Facebox state file to set.
	StateFile remototypes.File
}

// PutStateResponse is the response object for PutState calls.
type PutStateResponse struct{}
