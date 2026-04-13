package generateEmbeding

/*
* Copyright © 2023 - 2025. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"github.com/project-flogo/core/data/coerce"
)

// Constants for identifying settings and inputs
const (
	sEmbeddingFormat   = "ebeddingFormat"
	sAPIKey            = "apiKey"
	sEnpointURL        = "endPointURL"
	sMaxRetries        = "maxRetries"
	sDimensions        = "dimensions"
	sSafetyIdentifier  = "safetyIdentifier"
	sModel             = "embeddingModel"
	iPrompt            = "prompt"
	iTool              = "tool"
	oEmbedding         = "embedding"
	oEmbeddingMetadata = "embeddingMetadata"
)

// Settings defines configuration options for your activity
type Settings struct {
	ApiKey           string `md:"apiKey, required"`
	EndPointURL      string `md:"endPointURL"`
	MaxRetries       int    `md:"maxRetries, required"`
	Model            string `md:"model, required"`
	Dimensions       int64  `md:"dimensions, required"`
	EmbeddingFormat  string `md:"embeddingFormat"`
	SafetyIdentifier string `md:"safetyIdentifier"`
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.MaxRetries = 2
		s.Model = "text-embedding-3-small"
		s.EmbeddingFormat = "float"
		s.Dimensions = 1536
		return nil
	}

	var err error

	s.ApiKey, err = coerce.ToString(values[sAPIKey])
	if err != nil {
		return err
	}
	s.EndPointURL, err = coerce.ToString(values[sEnpointURL])
	if err != nil {
		return err
	}

	s.MaxRetries, err = coerce.ToInt(values[sMaxRetries])
	if err != nil {
		return err
	}

	if val, ok := values[sModel]; ok && val != nil {
		s.Model, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Model == "" {
			s.Model = "text-embedding-3-small"
		}
	}

	s.Dimensions, err = coerce.ToInt64(values[sDimensions])
	if err != nil {
		return err
	}

	if val, ok := values[sEmbeddingFormat]; ok && val != nil {
		s.EmbeddingFormat, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.EmbeddingFormat == "" {
			s.EmbeddingFormat = "float"
		}
	}

	s.SafetyIdentifier, err = coerce.ToString(values[sSafetyIdentifier])
	if err != nil {
		return err
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	Prompt map[string]interface{} `md:"prompt, required"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	i.Prompt, err = coerce.ToObject(values[iPrompt])
	if err != nil {
		return err
	}
	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iPrompt: i.Prompt,
	}
}

// Output defines what data the activity returns
type Output struct {
	Embedding         string `md:"embedding"`
	EmbeddingMetadata string `md:"embeddingMetadata"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oEmbedding:         o.Embedding,
		oEmbeddingMetadata: o.EmbeddingMetadata,
	}
}

// FromMap populates the struct from a map.
func (o *Output) FromMap(values map[string]interface{}) error {
	if values == nil {
		return nil
	}

	var err error
	if val, ok := values[oEmbedding]; ok && val != nil {
		o.Embedding, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[oEmbeddingMetadata]; ok && val != nil {
		o.Embedding, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	return nil
}
