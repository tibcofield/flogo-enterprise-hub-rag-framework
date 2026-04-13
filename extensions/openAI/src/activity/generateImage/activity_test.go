/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package generateImage

import (
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
func TestGenerateImage(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests disabled")
	}

	s := Settings{
		ApiKey:          os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:     os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Model:           "gpt-image-1.5",
		MaxRetries:      1,
		Mode:            "new",
		OutputFormat:    "jpeg",
		OutputDirectory: "/tmp/",
		InputDirectory:  "/tmp/",
		ImageSize:       "auto",
		Quality:         "auto",
		Compression:     50,
		Transparent:     false,
		Moderation:      "none",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "Create a clown being that is being chased by a dog in a park")
	tc.SetInput("outputFilename", "test_clown_image.jpg")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	metaData := tc.GetOutput("metaData")
	if metaData == nil {
		assert.Fail(t, "Expected some metadata but got nil")
	} else {
		fmt.Printf("metaData: %v\n", metaData)
		assert.True(t, done)
	}

}
