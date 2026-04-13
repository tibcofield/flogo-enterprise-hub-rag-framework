/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package generateEmbeding

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

// TestCreateEmbeddingOpenAILarge tests the creation of embeddings using the OpenAI platform api and text-embedding-3-large model
func TestCreateEmbeddingOpenAISmall(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:      os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Model:            "text-embedding-3-small",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}

}

// TestCreateEmbeddingOpenAILarge tests the creation of embeddings using OpenAI platform api and text-embedding-3-large model
func TestCreateEmbeddingOpenAILarge(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OPEN_AI_API_KEY"),
		EndPointURL:      os.Getenv("OPENAI_API_ENDPOINT_URL"),
		Model:            "text-embedding-3-large",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}

}

// TestCreateEmbeddingOllamanomicText tests the creation of embeddings using Ollama and nomic-embed-text:latest model
func TestCreateEmbeddingOllamanomicText(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "nomic-embed-text:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}

}

// TestCreateEmbeddingOllamamxbaiembedlarge tests the creation of embeddings using Ollama and mxbai-embed-large model
func TestCreateEmbeddingOllamamxbaiembedlarge(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "mxbai-embed-large:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}

}

// TestCreateEmbeddingOllamaGemma tests the creation of embeddings using the Ollama and embeddinggemma:latest model
func TestCreateEmbeddingOllamaGemma(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "embeddinggemma:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}

// TestCreateEmbeddingOllamabgem3 tests the creation of embeddings using on Ollama and bge-m3:latest model
func TestCreateEmbeddingOllamabgem3(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "bge-m3:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}

// TestCreateEmbeddingOllama tests the creation of embeddings using Ollama and all-minilm:latest model
func TestCreateEmbeddingOllamaallminilm(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "all-minilm:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}

// TestCreateEmbeddingOllamaArticEmbed2 tests the creation of embeddings using Ollama and artic-embed2:latest model
func TestCreateEmbeddingOllamaArticEmbed2(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "snowflake-arctic-embed2:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}

// TestCreateEmbeddingOllamabcelarge tests the creation of embeddings using Ollama and bge-large:latest model
func TestCreateEmbeddingOllamabgelarge(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "bge-large:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}

// TestCreateEmbeddingOllamaParaphraseMultilingual tests the creation of embeddings using Ollama and paraphrase-multilingual:latest model
func TestCreateEmbeddingOllamaParaphraseMultilingual(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "paraphrase-multilingual:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}

// TestCreateEmbeddingOllamagraniteembedding tests the creation of embeddings using Ollama and granite-embedding:latest model
func TestCreateEmbeddingOllamagraniteembedding(t *testing.T) {

	if os.Getenv("RUN_INTEGRATION") != "1" {
		t.Skip("integration tests  disabled")
	}

	s := Settings{
		ApiKey:           os.Getenv("OLLAMA_API_KEY"),
		EndPointURL:      os.Getenv("OLLAMA_OPENAI_API_ENDPOINT_URL"),
		Model:            "granite-embedding:latest",
		MaxRetries:       1,
		Dimensions:       1024,
		EmbeddingFormat:  "float",
		SafetyIdentifier: "User2536",
	}

	act := &Activity{
		Settings: &s,
	}

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("prompt", "This is my text, a man in red, slept in his bed, and then he went to the store")

	done, err := act.Eval(tc)
	if !done {
		fmt.Println(err)
	}

	embedding := tc.GetOutput("embedding")
	if embedding == nil {
		assert.Fail(t, "Expected an embedding   but got nil")
	} else {
		fmt.Printf("Response: %v\n", embedding)
		assert.True(t, done)
	}
}
