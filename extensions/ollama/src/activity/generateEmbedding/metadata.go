package generateEmbedding

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
	sOllamaHostURL         = "ollamaHostURL"
	sEmbeddingModel        = "embeddingModel"
	sTruncate              = "truncate"
	sKeepAlive             = "keepAlive"
	sEmbeddingModelVersion = "embeddingModelVersion"
	iEmbeddingText         = "embeddingText"
	iModelOptions          = "modelOptions"
	oEmbedding             = "embedding"
	oEmbeddingMetadata     = "embeddingMetadata"
)

// Settings defines configuration options for your activity
type Settings struct {
	OllamaHostURL        string `md:"ollamaHostURL,required"`
	EmbedingModel        string `md:"embeddingModel"`
	EmbedingModelVersion string `md:"embeddingModelVersion"`
	Truncate             bool   `md:"truncate"`
	KeepAlive            int    `md:"keepAlive"`
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.OllamaHostURL = "http://localhost:11434"
		s.EmbedingModel = "all-minilm"
		s.EmbedingModelVersion = "latest"
		s.Truncate = false
		s.KeepAlive = 5
		return nil
	}

	var err error
	if val, ok := values[sOllamaHostURL]; ok && val != nil {
		s.OllamaHostURL, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.OllamaHostURL == "" {
			s.OllamaHostURL = "http://localhost:11434"
		}
	}

	if val, ok := values[sEmbeddingModel]; ok && val != nil {
		s.EmbedingModel, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.EmbedingModel == "" {
			s.EmbedingModel = "all-minilm"
		}
	}

	if val, ok := values[sEmbeddingModelVersion]; ok && val != nil {
		s.EmbedingModelVersion, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.EmbedingModel == "" {
			s.EmbedingModel = "latest"
		}
	}

	if val, ok := values[sTruncate]; ok && val != nil {
		s.Truncate = false
		s.Truncate, err = coerce.ToBool(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sKeepAlive]; ok && val != nil {

		s.KeepAlive, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	EmbeddingText string `md:"embeddingText, required"`
	ModelOptions  string `md:"modelOptions"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	i.EmbeddingText, err = coerce.ToString(values[iEmbeddingText])
	if err != nil {
		return err
	}

	if val, ok := values[iModelOptions]; ok && val != nil {
		i.ModelOptions, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iEmbeddingText: i.EmbeddingText,
		iModelOptions:  i.ModelOptions,
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
