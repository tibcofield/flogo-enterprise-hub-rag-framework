package generateEmbeding

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"encoding/json"
	"time"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

// activityMd is the metadata for the activity.
var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})
var logger = log.ChildLogger(log.RootLogger(), "opeanAI-generate-embedding")

// Metadata returns the activity's metadata.
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

type ResponseMeta struct {
	Model           string
	PromptTokens    int64
	TotalTokens     int64
	EmbeddingLength int
	ExecutionTime   string
	ExecutedByUser  string
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

// Activity is a ChatGPT API activity
type Activity struct {
	Settings *Settings
}

// New creates a new instance of the Activity.
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := s.FromMap(ctx.Settings())
	if err != nil {
		return nil, err
	}

	return &Activity{
		Settings: s,
	}, nil

}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	embeddingText := ctx.GetInput(iPrompt).(string)

	if a.Settings.ApiKey == "" {
		logger.Error("Missing openAPI key")
		return false, nil
	}

	if a.Settings.Model == "" {
		logger.Error("Missing Model")
		return false, nil
	}

	if a.Settings.SafetyIdentifier == "" {
		logger.Error("Missing Safety Identifier")
		return false, nil
	}

	oaiClient := openai.NewClient(
		option.WithAPIKey(a.Settings.ApiKey),
		option.WithBaseURL(a.Settings.EndPointURL),
		option.WithMaxRetries(a.Settings.MaxRetries),
	)

	resp, err := oaiClient.Embeddings.New(
		context.Background(),
		openai.EmbeddingNewParams{
			Input: openai.EmbeddingNewParamsInputUnion{
				OfString: openai.String(embeddingText),
			},
			Model:          a.Settings.Model,
			Dimensions:     openai.Int(a.Settings.Dimensions),
			User:           openai.String(a.Settings.SafetyIdentifier),
			EncodingFormat: openai.EmbeddingNewParamsEncodingFormat(a.Settings.EmbeddingFormat),
		},
	)

	if err != nil {
		logger.Error("Embedding creation error: %v", err)
		return false, nil
	}
	currentTime := time.Now()

	var responseMeta ResponseMeta

	responseMeta.EmbeddingLength = len(resp.Data[0].Embedding)
	responseMeta.ExecutedByUser = a.Settings.SafetyIdentifier
	responseMeta.ExecutionTime = currentTime.Format(time.RFC1123Z)
	responseMeta.Model = resp.Model
	responseMeta.PromptTokens = resp.Usage.PromptTokens
	responseMeta.TotalTokens = resp.Usage.TotalTokens

	jsonMetaData, err := json.Marshal(responseMeta)
	if err != nil {
		logger.Error("Error generating Meta Data %v", err)
		return false, nil
	}

	vector := resp.Data[0].Embedding

	ctx.SetOutput(oEmbedding, vector)
	ctx.SetOutput(oEmbeddingMetadata, jsonMetaData)
	return true, nil
}
