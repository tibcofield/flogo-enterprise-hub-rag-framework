package generateCompletion

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
	"os"
	"strconv"
	"time"

	ollama "github.com/ollama/ollama/api"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

var ollamaClient *ollama.Client
var logger = log.ChildLogger(log.RootLogger(), "ollama-generateCompletion")

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
	model         string
	modelMetaData string
	settings      *Settings
	modelResposne string
}

// New creates a new instance of the Activity.
func New(ctx activity.InitContext) (activity.Activity, error) {

	s := &Settings{}
	err := s.FromMap(ctx.Settings())
	// //err := metadata.MapToStruct(ctx.Settings(), s, true)  change in a approach in .14 of core
	if err != nil {
		return nil, err
	}

	model := s.Model + ":" + s.ModelVersion

	return &Activity{
		model:         model,
		modelMetaData: "Model: " + model,
		settings:      s,
		modelResposne: "",
	}, nil
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	modelResponse, err := a.generateCompltion(input.Prompt, input.Suffix, input.FileName, input.ModelOptions)

	if err != nil {
		fmt.Printf("Error getting resposne: %s\n", err)
		return
	}

	if a.settings.OutputFormat == "json" {
		ctx.SetOutput(oOutputJson, modelResponse)
	} else {
		ctx.SetOutput(oOouputText, modelResponse)
	}
	ctx.SetOutput(oMetaData, a.modelMetaData)
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

	baseURL, err := url.Parse(a.settings.OllamaHostURL)

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
			if model.Name == a.model {
				modelFound = true
			}
		}
		if !modelFound {
			fmt.Printf("Model %s not found\n", a.model)
			fmt.Println("Available models:")
			for _, model := range modelList.Models {
				fmt.Printf("  %s\n", model.Name)
			}
			return fmt.Errorf("model %s not found", a.model)
		}
		ollamaClient = oc
	}

	return nil
}

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

func (a *Activity) generateCompltion(prompt string, suffix string, filename string, modelOptions string) (modelResposne string, err error) {

	logger.Info("Generate Completion called with: " + prompt + " Suffix: " + suffix + " ModelOptions: " + modelOptions)

	if err := a.getOrInitOllamaClient(); err != nil {
		return "", err
	}

	var options ollama.Options

	logger.Info()

	if modelOptions != "" {
		err := json.Unmarshal([]byte(modelOptions), &options)

		if err != nil {
			return "", fmt.Errorf("error parsing model options: %v", err)
		}

		logger.Info("Using custom model options")
		logger.Info("Seed value is set to " + strconv.Itoa(options.Seed))

	} else {
		// Use default options if none are provided
		logger.Info("Using default model options")

	}

	req := &ollama.GenerateRequest{
		Model:  a.model,
		Prompt: prompt,
		// Activity cannot support streaming so ensure this is always set to false
		Stream: new(bool),
	}

	if suffix != "" {
		req.Suffix = suffix
	}

	if filename != "" {

		imgData, err := os.ReadFile(filename)

		req.Images = []ollama.ImageData{imgData}
		if err != nil {
			return "", fmt.Errorf("error getting completion: %v", err)
		}
	}

	if a.settings.OutputFormat == "json" {
		req.Format = []byte(`"json"`)
	}

	req.Options = getOptionsMap(&options)

	ctx := context.Background()

	var modelRespone string
	respFunc := func(resp ollama.GenerateResponse) error {

		modelRespone += resp.Response
		return nil
	}

	err = ollamaClient.Generate(ctx, req, respFunc)
	if err != nil {
		return "", fmt.Errorf("error getting completion: %v", err)
	}

	return modelRespone, nil
}
