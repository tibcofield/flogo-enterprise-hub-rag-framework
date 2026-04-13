package generateCompletion

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
	sOllamaHostURL = "ollamaHostURL"
	sModel         = "model"
	sModelVersion  = "modelVersion"
	sSystemMsg     = "systemMsg"
	sOutputFormat  = "outputFormat"
	sTemplate      = "template"
	sContext       = "context"
	sStream        = "stream"
	sRaw           = "raw"
	sKeepAlive     = "keepAlive"
	iPrompt        = "prompt"
	iSuffix        = "suffix"
	iFileName      = "fileName"
	iModelOptions  = "modelOptions"
	oOutputJson    = "outputJSON"
	oOouputText    = "outputText"
	oMetaData      = "metaData"
)

// Settings defines configuration options for your activity
type Settings struct {
	OllamaHostURL string `md:"ollamaHostURL,required"`
	Model         string `md:"model"`
	ModelVersion  string `md:"modelVersion"`
	SystemMsg     string `md:"systemMsg"`
	OutputFormat  string `md:"outputFormat"`
	Template      string `md:"template"`
	Context       string `md:"context"`
	Stream        bool   `md:"stream"`
	Raw           bool   `md:"raw"`
	KeepAlive     int    `md:"keepAlive"`
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.OllamaHostURL = "http://localhost:11434"
		s.Model = "llama3.2"
		s.OutputFormat = "text"
		s.KeepAlive = 5
		s.Stream = false
		s.Raw = false

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

	if val, ok := values[sModel]; ok && val != nil {
		s.Model, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Model == "" {
			s.Model = "llama3.2"
		}
	}

	if val, ok := values[sModelVersion]; ok && val != nil {
		s.ModelVersion, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.ModelVersion == "" {
			s.ModelVersion = "latest"
		}
	}

	if val, ok := values[sSystemMsg]; ok && val != nil {

		s.SystemMsg, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sOutputFormat]; ok && val != nil {

		s.OutputFormat, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.OutputFormat == "" {
			s.OutputFormat = "text"
		}
	}

	if val, ok := values[sTemplate]; ok && val != nil {
		s.Template, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sContext]; ok && val != nil {
		s.Context, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sRaw]; ok && val != nil {
		s.Raw, err = coerce.ToBool(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sStream]; ok && val != nil {
		s.Stream, err = coerce.ToBool(val)
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
	Prompt       string `md:"prompt, required"`
	Suffix       string `md:"suffix"`
	FileName     string `md:"fileName"`
	ModelOptions string `md:"modelOptions"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	if val, ok := values[iPrompt]; ok && val != nil {
		i.Prompt, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[iSuffix]; ok && val != nil {
		i.Suffix, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iFileName]; ok && val != nil {
		i.FileName, err = coerce.ToString(val)
		if err != nil {
			return err
		}
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
		iFileName:     i.FileName,
		iPrompt:       i.Prompt,
		iSuffix:       i.Suffix,
		iModelOptions: i.ModelOptions,
	}
}

// Output defines what data the activity returns
type Output struct {
	OutputJSON string `md:"outputJSON"`
	OutputText string `md:"outputText"`
	MetaData   string `md:"metaData"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oOutputJson: o.OutputJSON,
		oOouputText: o.OutputText,
		oMetaData:   o.MetaData,
	}
}

// FromMap populates the struct from a map.
func (o *Output) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}
	var err error
	if val, ok := values[oOutputJson]; ok && val != nil {
		o.OutputJSON, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[oOouputText]; ok && val != nil {
		o.OutputText, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[oMetaData]; ok && val != nil {
		o.MetaData, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	return nil
}
