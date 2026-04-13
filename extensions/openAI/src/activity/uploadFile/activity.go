package uploadFile

/*
* Copyright © 2023 - 2026. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"context"
	"mime"
	"os"
	"path/filepath"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/packages/param"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
)

var logger = log.ChildLogger(log.RootLogger(), "opeanAI-upload-file")

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

	if a.Settings.ApiKey == "" {
		logger.Error("Missing openAPI key")
		return false, err
	}

	oaiClient := openai.NewClient(
		option.WithAPIKey(a.Settings.ApiKey),
		option.WithBaseURL(a.Settings.EndPointURL),
	)

	clientCtx := context.Background()

	fileReader, err := os.Open(input.FileName)
	if err != nil {
		logger.Error(err.Error())
		return false, err
	}

	fileName := filepath.Base(input.FileName)

	mimeType := mime.TypeByExtension(filepath.Ext(input.FileName))
	logger.Infof("MimeType: %s", mimeType)

	inputFile := openai.File(fileReader, fileName, mimeType)

	fileResp, err := oaiClient.Files.New(clientCtx, openai.FileNewParams{
		File:    inputFile,
		Purpose: openai.FilePurpose(a.Settings.Purpose),
	})
	if err != nil {
		logger.Errorf("File upload error: %v\n", err)
		return false, err
	}

	logger.Info("Populating custom metadata from source data...")
	customMetadata := make(map[string]openai.VectorStoreFileNewParamsAttributeUnion)
	for _, item := range input.FileAttributes {
		customMetadata[item.Key] = openai.VectorStoreFileNewParamsAttributeUnion{
			OfString: param.NewOpt(item.Value),
		}
		logger.Info("adding" + item.Key)
	}

	if a.Settings.VectorStoreID != "" {
		_, err = oaiClient.VectorStores.Files.New(clientCtx, a.Settings.VectorStoreID,
			openai.VectorStoreFileNewParams{
				FileID:     fileResp.ID,
				Attributes: customMetadata,
				ChunkingStrategy: openai.FileChunkingStrategyParamOfStatic(openai.StaticFileChunkingStrategyParam{
					MaxChunkSizeTokens: a.Settings.MaxChunkSizeTokens,
					ChunkOverlapTokens: a.Settings.ChunkOverlapTokens,
				}),
			},
		)

		if err != nil {
			logger.Errorf("Error generating vector store embedding: %v\n", err)
			return false, err
		}
	}

	ctx.SetOutput(oMetaData, "File Uploaded ID:"+fileResp.ID)
	return true, nil
}
