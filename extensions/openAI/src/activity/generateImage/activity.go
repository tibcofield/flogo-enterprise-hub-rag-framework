package generateImage

/*
* Copyright © 2023 - 2025. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

var logger = log.ChildLogger(log.RootLogger(), "opeanAI-generate-image")

// activityMd is the metadata for the activity.
var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Metadata returns the activity's metadata.
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
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

	prompt := ctx.GetInput(iPrompt).(string)
	fileName := a.Settings.OutputDirectory + ctx.GetInput(iOutputFileName).(string) + "." + a.Settings.OutputFormat

	if a.Settings.ApiKey == "" {
		logger.Error("Missing openAPI key")

	}

	logger.Info("Model:" + a.Settings.Model)

	oaiClient := openai.NewClient(
		option.WithAPIKey(a.Settings.ApiKey),
		option.WithBaseURL(a.Settings.EndPointURL),
		option.WithMaxRetries(a.Settings.MaxRetries),
	)

	clientCtx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// Request an image
	imgResp, err := oaiClient.Images.Generate(clientCtx, openai.ImageGenerateParams{
		Model:  a.Settings.Model,
		Prompt: prompt,
		Size:   openai.ImageGenerateParamsSize(a.Settings.ImageSize),
	})
	if err != nil {
		logger.Errorf("Image generation error: %v\n", err)
		return false, err
	}

	if len(imgResp.Data) == 0 {
		logger.Error("No image data returned")
		return false, fmt.Errorf("no image data returned from OpenAI API")
	}

	// Get the Base64 data
	b64Data := imgResp.Data[0].B64JSON

	// Decode the Base64 string to bytes
	imgBytes, err := base64.StdEncoding.DecodeString(b64Data)
	if err != nil {
		logger.Errorf("Base64 decode error: %v\n", err)
		return false, err
	}

	// Save to file

	if err := os.WriteFile(fileName, imgBytes, 0644); err != nil {
		logger.Errorf("File save error: %v\n", err)
		return true, err
	}

	logger.Errorf("Image saved as %s\n", fileName)

	ctx.SetOutput(oMetaData, "Image generated successfully")
	return true, nil
}
