package classificationbox

import (
	"github.com/machinebox/remoto/remototypes"
)

// Classificationbox lets you use machine learning to automatically classify
// various types of data, such as text, images, structured and unstructured data.
type Classificationbox interface {
	CreateModel(CreateModelRequest) CreateModelResponse
	Teach(TeachRequest) TeachResponse
	Predict(PredictRequest) PredictResponse

	ListModels(ListModelsRequest) ListModelsResponse
	DeleteModel(DeleteModelRequest) DeleteModelResponse

	GetState(GetStateRequest) remototypes.FileResponse
	PutState(PutStateRequest) PutStateResponse
}

type CreateModelRequest struct {
	Model Model
}

type CreateModelResponse struct {
}

type Model struct {
	ID      string
	Name    string
	Options ModelOptions
	Classes []string
}

type ModelOptions struct {
	Ngrams    int
	Skipgrams int
}

type TeachRequest struct {
	ModelID string
	Class   string
	Inputs  []Feature
}

type TeachResponse struct {
}

type Feature struct {
	Key   string
	Type  string
	Value string
	File  remototypes.File
}

type PredictRequest struct {
	ModelID string
	Limit   int
	Inputs  []Feature
}

type PredictResponse struct {
	Classes []PredictedClass
}

type PredictedClass struct {
	ID    string
	Score float64
}

type ListModelsRequest struct {
}
type ListModelsResponse struct {
	Models []Model
}

type GetStateRequest struct{}

type GetStateResponse struct {
	StateFile remototypes.File
}

type PutStateRequest struct {
	StateFile remototypes.File
}

type PutStateResponse struct{}

type DeleteModelRequest struct {
	ModelID string
}

type DeleteModelResponse struct{}
