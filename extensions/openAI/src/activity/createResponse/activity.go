package createResponse

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"log"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
	"github.com/project-flogo/core/activity"
)

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

	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return false, err
	}

	prompt := input.Prompt
	base64String := input.Base64String

	if a.Settings.ApiKey == "" {
		log.Fatal("Missing openAPI key")
	}

	oaiClient := openai.NewClient(
		option.WithAPIKey(a.Settings.ApiKey),
		option.WithBaseURL(a.Settings.EndPointURL),
		option.WithMaxRetries(a.Settings.MaxRetries),
	)

	clientCtx := context.Background()

	params := responses.ResponseNewParams{
		Model: a.Settings.Model, // openai.ChatModelGPT4_1, // or another supported model
		Input: responses.ResponseNewParamsInputUnion{
			OfInputItemList: responses.ResponseInputParam{
				responses.ResponseInputItemParamOfMessage(
					responses.ResponseInputMessageContentListParam{
						responses.ResponseInputContentUnionParam{
							OfInputImage: &responses.ResponseInputImageParam{
								ImageURL: openai.String("data:" + a.Settings.InputFormat + "," + base64String),
								Type:     "input_image",
							},
						},
						responses.ResponseInputContentUnionParam{
							OfInputText: &responses.ResponseInputTextParam{
								Text: prompt,
								Type: "input_text",
							},
						},
					},
					"user",
				),
			},
		},
		MaxOutputTokens: openai.Int(256),
		Store:           openai.Bool(false),
	}

	// Send the request
	resp, err := oaiClient.Responses.New(clientCtx, params)

	if err != nil {
		log.Fatalf("Responses.New error: %v", err)
	}

	// Display the output
	log.Println("Response ID:", resp.ID)
	log.Println("Model:", resp.Model)
	log.Println("prompt:", prompt)

	outputString := ""

	for _, output := range resp.Output {
		if output.Type == "message" {
			for _, content := range output.Content {
				if content.Type == "output_text" {
					outputString += content.Text
				}
			}
		}
	}
	log.Println("Output Text:", outputString)
	ctx.SetOutput(oResponse, outputString)
	return true, nil
}
