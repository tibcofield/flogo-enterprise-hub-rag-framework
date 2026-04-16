/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package uploadFile

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)
	assert.NotNil(t, act)
}

func TestUploadFileFileAttributeTest(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		VectorStoreID:      os.Getenv("VECTOR_STORE_ID"),
		ApiKey:             os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:        os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Purpose:            "assistants",
		ChunkOverlapTokens: 400,
		MaxChunkSizeTokens: 800,
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())

	relPath := "../../../testdata/tib_ad_sdk_relnotes.pdf"
	fmt.Printf("Relative path: %s\n", relPath)

	tc.SetInput("filename", relPath)

	fileAttributes := []FileAttributeData{
		{
			Key:   "ProductGroup",
			Value: "Integration",
		},
		{
			Key:   "Product",
			Value: "Flogo",
		},
		{
			Key:   "DocumentType",
			Value: "ReleaseNotes",
		},
		{
			Key:   "ProductVersion",
			Value: "5.8.0",
		},
	}
	logger.Infof("File Attributes: %+v", fileAttributes)
	tc.SetInput("fileAttributes", fileAttributes)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
		assert.Fail(t, "activity failed")
	}

	metadata := tc.GetOutput("metaData")

	fmt.Printf("Metadata: %s", metadata)
}

func TestSearchDocumentsDefaultRelPath(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		VectorStoreID:      os.Getenv("VECTOR_STORE_ID"),
		ApiKey:             os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:        os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Purpose:            "assistants",
		ChunkOverlapTokens: 400,
		MaxChunkSizeTokens: 800,
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())

	relPath := "../../../testdata/tib_ad_sdk_relnotes.pdf"
	fmt.Printf("Relative path: %s\n", relPath)

	tc.SetInput("filename", relPath)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
		assert.Fail(t, "activity failed")
	}

	metadata := tc.GetOutput("metaData")

	fmt.Printf("Metadata: %s", metadata)
}

func TestSearchDocumentsDefaultAbsPath(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		VectorStoreID:      os.Getenv("VECTOR_STORE_ID"),
		ApiKey:             os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:        os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Purpose:            "assistants",
		ChunkOverlapTokens: 400,
		MaxChunkSizeTokens: 800,
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())

	relPath := "../../../testdata/tib_ad_sdk_relnotes.pdf"
	fmt.Printf("Relative path: %s\n", relPath)
	absPath, err := filepath.Abs(relPath)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Absolute path: %s\n", absPath)
	tc.SetInput("filename", relPath)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
		assert.Fail(t, "activity failed")
	}

	metadata := tc.GetOutput("metaData")

	fmt.Printf("Metadata: %s", metadata)
}
