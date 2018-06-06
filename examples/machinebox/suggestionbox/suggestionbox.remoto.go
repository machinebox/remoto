package suggestionbox

import (
	"context"

	"github.com/machinebox/remoto/remototypes"
)

type Suggestionbox interface {
	CreateModel(context.Context, *CreateModelRequest) (*CreateModelResponse, error)
	Predict(context.Context, *PredictRequest) (*PredictResponse, error)
	Reward(context.Context, *RewardRequest) (*RewardResponse, error)

	ListModels(context.Context, *ListModelsRequest) (*ListModelsResponse, error)
	DeleteModel(context.Context, *DeleteModelRequest) (*DeleteModelResponse, error)

	GetState(context.Context, *GetStateRequest) (*GetStateResponse, error)
	PutState(context.Context, *PutStateRequest) (*PutStateResponse, error)
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
