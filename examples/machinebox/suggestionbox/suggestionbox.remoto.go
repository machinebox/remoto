package suggestionbox

import (
	"github.com/machinebox/remoto/remototypes"
)

type Suggestionbox interface {
	CreateModel(CreateModelRequest) CreateModelResponse
	Predict(PredictRequest) PredictResponse
	Reward(RewardRequest) RewardResponse

	ListModels(ListModelsRequest) ListModelsResponse
	DeleteModel(DeleteModelRequest) DeleteModelResponse

	GetState(GetStateRequest) remototypes.FileResponse
	PutState(PutStateRequest) PutStateResponse
}

type CreateModelRequest struct {
	Model Model
}

type CreateModelResponse struct{}

type Model struct {
	ID      string
	Name    string
	Options ModelOptions
	Choices []Choice
}

type ModelOptions struct {
	RewardExpirationSeconds int
	Ngrams                  int
	Skipgrams               int
	Mode                    string
	Epsilon                 float64
	Cover                   float64
}

type Choice struct {
	ID       string
	Features []Feature
}

type PredictedChoice struct {
	ID       string
	Features []Feature
	RewardID string
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
	Choices []PredictedChoice
}

type RewardRequest struct {
	ModelID  string
	RewardID string
	Value    int
}

type RewardResponse struct{}

type ListModelsRequest struct {
}

type ListModelsResponse struct {
	Models []Model
}

type DeleteModelRequest struct {
	ModelID string
}

type DeleteModelResponse struct{}

type GetStateRequest struct{}

type GetStateResponse struct {
	StateFile remototypes.File
}

type PutStateRequest struct {
	StateFile remototypes.File
}

type PutStateResponse struct{}
