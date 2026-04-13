/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package vectorSearch

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

//VectorDBURL:   "localhost:50051",

func TestSearchDocumentsDefault(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		VectorStoreID:      os.Getenv("VECTOR_STORE_ID"),
		ApiKey:             os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:        os.Getenv("OPENAI_API_ENDPOINT_URL"),
		MaxNumberOfResults: 10,
		RewriteQuery:       false,
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("searchString", "tibco ems versions")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	success := tc.GetOutput("searchResultRows")

	fmt.Printf("Results: %s", success)

	if success != 0 {
		assert.True(t, done)

	}
}

// func TestUpdateDocumentrow(t *testing.T) {
// 	s := Settings{
// 		APIKey:        "MySecretKey",
// 		VectorDBURL:   "192.168.1.50:50052",
// 		DocumentStore: "ds_store_tdks",
// 		Action:        "Update",
// 	}

// 	act := &Activity{
// 		Settings: &s,
// 	}

// 	tc := test.NewActivityContext(act.Metadata())

// 	tc.SetInput("docID", "9144897421204059669")
// 	tc.SetInput("status", "IT WORKS..3.")
// 	tc.SetInput("docTitle", "Updated Activity Test Doc 2")
// 	tc.SetInput("chunks", 0)
// 	tc.SetInput("vectorStore", "vs_store_tdks_doc")
// 	tc.SetInput("metaData", "Pages 44")
// 	done, err := act.Eval(tc)
// 	if !done {
// 		fmt.Println(err)
// 	}

// 	success := tc.GetOutput("success")

// 	fmt.Printf("Document Row Created: %d", tc.GetOutput("docID"))

// 	if success != 0 {
// 		assert.True(t, done)
// 	}
// }

// func TestUpdateROWSTATUS(t *testing.T) {
// 	s := Settings{
// 		APIKey:        "MySecretKey",
// 		VectorDBURL:   "192.168.1.50:50052",
// 		DocumentStore: "ds_store_tdks",
// 		Action:        "Update",
// 	}

// 	act := &Activity{
// 		Settings: &s,
// 	}

// 	tc := test.NewActivityContext(act.Metadata())

// 	tc.SetInput("docID", "9144897421204059669")
// 	tc.SetInput("status", "IT WORKS..3.")

// 	done, err := act.Eval(tc)
// 	if !done {
// 		fmt.Println(err)
// 	}

// 	success := tc.GetOutput("success")

// 	fmt.Printf("Document Row Created: %d", tc.GetOutput("docID"))

// 	if success != 0 {
// 		assert.True(t, done)
// 	}
// }

// func TestGetROW(t *testing.T) {
// 	s := Settings{
// 		APIKey:        "MySecretKey",
// 		VectorDBURL:   "192.168.1.50:50052",
// 		DocumentStore: "ds_store_tdks",
// 		Action:        "Get",
// 	}

// 	act := &Activity{
// 		Settings: &s,
// 	}

// 	tc := test.NewActivityContext(act.Metadata())

// 	tc.SetInput("docID", "-3818324397740905472")

// 	done, err := act.Eval(tc)
// 	if !done {
// 		fmt.Println(err)
// 	}

// 	success := tc.GetOutput("success")

// 	fmt.Printf("Document Row Recieved: %d", tc.GetOutput("docID"))
// 	fmt.Printf("Document Row Recieved: %f", tc.GetOutput("rows"))

// 	if success != 0 {
// 		assert.True(t, done)
// 	}
// }
