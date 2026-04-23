/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package createResponse

import (
	"encoding/base64"
	"fmt"
	"os"
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

//VectorDBURL:   "localhost:50051",

// func guessMimeType(path string, data []byte) string {
// 	ext := strings.ToLower(filepath.Ext(path))
// 	switch ext {
// 	case ".png":
// 		return "image/png"
// 	case ".jpg", ".jpeg":
// 		return "image/jpeg"
// 	case ".webp":
// 		return "image/webp"
// 	case ".gif":
// 		return "image/gif"
// 	default:
// 		// Fallback: sniff content
// 		return http.DetectContentType(data)
// 	}
// }

func TestDescribeImage(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:      os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL: os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Model:       "gpt-4.1-mini",
		InputFormat: "image/png;base64",
		MaxRetries:  1,
	}

	act := &Activity{
		Settings: &s,
	}

	imagePath := "../../../testdata/images/man-cannon.png"

	imgBytes, err := os.ReadFile(imagePath)
	if err != nil {
		t.Logf("read image: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(imgBytes)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "Describe the  image?")
	tc.SetInput("base64String", b64)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	response := tc.GetOutput("response")
	if response == nil {
		assert.Fail(t, "Expected a response but got nil")
	} else {
		fmt.Printf("Response: %v\n", response)
		assert.True(t, done)
	}

}

func TestDescribeImageOllama(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:      os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL: os.Getenv("http://localhost:11434/v1/"),
		Model:       "qwen3.5:latest",
		InputFormat: "image/png;base64",
		MaxRetries:  1,
	}

	act := &Activity{
		Settings: &s,
	}

	imagePath := "../../../testdata/images/man-cannon.png"

	imgBytes, err := os.ReadFile(imagePath)
	if err != nil {
		t.Logf("read image: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(imgBytes)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "Describe the image?")
	tc.SetInput("base64String", b64)

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	response := tc.GetOutput("response")
	if response == nil {
		assert.Fail(t, "Expected a response but got nil")
	} else {
		fmt.Printf("Response: %v\n", response)
		assert.True(t, done)
	}

}
