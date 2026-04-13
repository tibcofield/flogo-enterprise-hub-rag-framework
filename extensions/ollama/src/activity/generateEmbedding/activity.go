package generateEmbedding

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	ollama "github.com/ollama/ollama/api"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

// var ollamaClient *ollama.Client
// var embeddingModel string = "all-minilm:latest"
var ollamaClient *ollama.Client
var logger = log.ChildLogger(log.RootLogger(), "ollama-generate-embedding")

// activityMd is the metadata for the activity.
var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Metadata returns the activity's metadata.
func (a *Activity) Metadata() *activity.Metadata {

	return activityMd
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

// Activity is a Ollama API activity
type Activity struct {
	Model         string
	ModelMetaData string
	Settings      *Settings
}

// New creates a new instance of the Activity.
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := s.FromMap(ctx.Settings())
	if err != nil {
		return nil, err
	}

	model := s.EmbedingModel + ":" + s.EmbedingModelVersion

	return &Activity{
		Model:         model,
		ModelMetaData: "Model: " + model,
		Settings:      s,
	}, nil
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	value_embeddings, embeddings_meta, err := a.getEmbeddings(a.Settings.EmbedingModel, input.EmbeddingText, input.ModelOptions)
	if err != nil {
		fmt.Printf("Error getting embedding: %s\n", err)
		return
	}

	ctx.SetOutput(oEmbedding, value_embeddings)
	ctx.SetOutput(oEmbeddingMetadata, embeddings_meta)
	return true, nil
}

func (a *Activity) getOrInitOllamaClient() error {

	settings := &Settings{}

	if settings.OllamaHostURL != "" {
		return fmt.Errorf("ollama host url not valid setting is %s", settings.OllamaHostURL)
	}

	httpClient := &http.Client{
		Timeout: 100 * time.Second,
	}

	baseURL, err := url.Parse(a.Settings.OllamaHostURL)

	if err != nil {
		return fmt.Errorf("invalid Ollama host URL: %v", err)
	}

	oc := ollama.NewClient(baseURL, httpClient)

	if oc == nil {
		return fmt.Errorf("error creating ollama client")
	} else if modelList, err := oc.List(context.Background()); err != nil {
		return err
	} else {
		modelFound := false
		for _, model := range modelList.Models {
			if model.Name == a.Model {
				modelFound = true
			}
		}
		if !modelFound {
			logger.Info("Model %s not found\n", a.Model)
			logger.Info("Available models:")
			for _, model := range modelList.Models {
				logger.Info("  %s\n", model.Name)
			}
			return fmt.Errorf("model %s not found", a.Model)
		}
		logger.Info("Initializing ollama Embedding using model ", a.Model)
		ollamaClient = oc
	}

	return nil
}

// ToMap converts the struct to a map.
func getOptionsMap(o *ollama.Options) map[string]interface{} {
	return map[string]interface{}{
		"NumCtx":        o.NumCtx,
		"RepeatLastN":   o.RepeatLastN,
		"RepeatPenalty": o.RepeatPenalty,
		"Temperature":   o.Temperature,
		"Seed":          o.Seed,
		"Stop":          o.Stop,
		"NumPredict":    o.NumPredict,
		"TopK":          o.TopK,
		"topP":          o.TopP,
		"minP":          o.MinP,
	}
}

func (a *Activity) getEmbeddings(model string, value string, modelOptions string) ([]float32, string, error) {

	var metadata string = "Model: " + a.Settings.EmbedingModel + " Version: " + a.Settings.EmbedingModelVersion

	if err := a.getOrInitOllamaClient(); err != nil {
		return nil, "", err
	}

	var options ollama.Options

	if modelOptions != "" {
		err := json.Unmarshal([]byte(modelOptions), &options)

		if err != nil {
			return nil, "", fmt.Errorf("error parsing model options: %v", err)
		}

		logger.Debug("Using custom model options")
		logger.Debug("Seed value is set to " + strconv.Itoa(options.Seed))

	} else {
		// Use default options if none are provided
		logger.Debug("Using default model options")

	}
	metadata += fmt.Sprintf(" Model Options: %v", getOptionsMap(&options))

	logger.Debug("Getting embeddings: " + metadata)

	keepAliveDuration := time.Duration(a.Settings.KeepAlive*60) * time.Second
	if resp, err := ollamaClient.Embed(context.Background(), &ollama.EmbedRequest{
		Model:     model,
		Input:     value,
		Options:   getOptionsMap(&options),
		Truncate:  &a.Settings.Truncate,
		KeepAlive: &ollama.Duration{Duration: keepAliveDuration},
	}); err != nil {
		return nil, "", err
	} else {
		return resp.Embeddings[0], metadata, nil
	}
}
