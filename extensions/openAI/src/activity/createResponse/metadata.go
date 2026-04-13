package createResponse

/*
* Copyright © 2023 - 2024. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"github.com/project-flogo/core/data/coerce"
)

// Constants for identifying settings and inputs
const (
	sEnodPointURL = "endPointURL"
	sAPIKey       = "apiKey"
	SMaxRetries   = "maxRetries"
	sInputFormat  = "inputFormat"
	sOutputFormat = "outputFormat"
	sModel        = "model"
	iPrompt       = "prompt"
	iTool         = "tool"
	iBase64String = "base64String"
	oResponse     = "response"
)

// Settings defines configuration options for your activity
type Settings struct {
	EndPointURL  string `md:"endPointURL, required"`
	ApiKey       string `md:"apiKey, required"`
	Model        string `md:"model, required"`
	MaxRetries   int    `md:"maxRetries"`   // Flogo metadata tag
	InputFormat  string `md:"inputFormat"`  // Flogo metadata tag
	OutputFormat string `md:"outputFormat"` // Flogo metadata tag
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.EndPointURL = ""
		s.ApiKey = ""
		s.InputFormat = "text"
		s.OutputFormat = "text"
		return nil
	}

	var err error

	if val, ok := values[sEnodPointURL]; ok && val != nil {
		s.EndPointURL, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[sInputFormat]; ok && val != nil {
		s.InputFormat, err = coerce.ToString(val)

		if err != nil {
			return err
		}

		if s.InputFormat == "" {
			s.InputFormat = "text"
		}
	}

	if val, ok := values[SMaxRetries]; ok && val != nil {
		s.MaxRetries, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[sModel]; ok && val != nil {
		s.Model, err = coerce.ToString(val)
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

	s.ApiKey, err = coerce.ToString(values[sAPIKey])
	if err != nil {
		return err
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	Prompt       string                 `md:"prompt, required"`
	Tool         map[string]interface{} `md:"tool, required"`
	Base64String string                 `md:"base64String, required"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	i.Prompt, err = coerce.ToString(values[iPrompt])
	if err != nil {
		return err
	}

	i.Tool, err = coerce.ToObject(values[iTool])
	if err != nil {
		return err
	}

	i.Base64String, err = coerce.ToString(values[iBase64String])
	if err != nil {
		return err
	}

	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iPrompt:       i.Prompt,
		iTool:         i.Tool,
		iBase64String: i.Base64String,
	}
}

// Output defines what data the activity returns
type Output struct {
	Response string `md:"response"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oResponse: o.Response,
	}
}

// FromMap populates the struct from a map.
func (o *Output) FromMap(values map[string]interface{}) error {
	if values == nil {
		return nil
	}

	var err error
	if val, ok := values[oResponse]; ok && val != nil {
		o.Response, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	return nil
}
