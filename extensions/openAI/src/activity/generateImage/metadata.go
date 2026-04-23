package generateImage

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
	sEnodPointURL    = "endPointURL"
	sAPIKey          = "apiKey"
	sModel           = "model"
	sMode            = "mode"
	sOutputFormat    = "outputFormat"
	sOutputDirectory = "outputDirectory"
	sInputDirectory  = "inputDirectory"
	sImageSize       = "imageSize"
	sQuality         = "quality"
	sCompression     = "compression"
	sTransparent     = "Transparent"
	sModeration      = "Moderation"
	iPrompt          = "prompt"
	iInputFile       = "inputFilename"
	iMaskFile        = "maskFilename"
	iOutputFileName  = "outputFilename"
	oMetaData        = "metaData"
	oOutputFileURL   = "outputFileURL"
)

// Settings defines configuration options for your activity
type Settings struct {
	ApiKey          string `md:"apiKey, required"`
	EndPointURL     string `md:"endPointURL"`
	MaxRetries      int    `md:"maxRetries, required"`
	Model           string `md:"model, required"`
	Mode            string `md:"mode, required"`
	OutputFormat    string `md:"outputFormat"`
	OutputDirectory string `md:"outputDirectory"`
	InputDirectory  string `md:"inputDirectory"`
	ImageSize       string `md:"imageSize"`
	Quality         string `md:"quality"`
	Compression     int    `md:"compression"`
	Transparent     bool   `md:"Transparent"`
	Moderation      string `md:"Moderation"`
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.EndPointURL = ""
		s.Mode = "new"
		s.Model = "gpt-image-1"
		s.OutputFormat = "jpeg"
		s.OutputDirectory = "/tmp/"
		s.InputDirectory = "/tmp/"
		s.ImageSize = "auto"
		s.Quality = "auto"
		s.Compression = 50
		s.Transparent = false
		s.Moderation = "auto"
		return nil
	}

	var err error

	s.EndPointURL, err = coerce.ToString(values[sEnodPointURL])
	if err != nil {
		return err
	}

	s.ApiKey, err = coerce.ToString(values[sAPIKey])
	if err != nil {
		return err
	}

	if val, ok := values[sModel]; ok && val != nil {
		s.Model, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Model == "" {
			s.Model = "gpt-image-1"
		}
	}

	if val, ok := values[sMode]; ok && val != nil {
		s.Mode, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Mode == "" {
			s.Mode = "new"
		}
	}

	if val, ok := values[sOutputFormat]; ok && val != nil {
		s.OutputFormat, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.OutputFormat == "" {
			s.OutputFormat = "jpeg"
		}
	}

	if val, ok := values[sOutputFormat]; ok && val != nil {
		s.OutputFormat, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.OutputFormat == "" {
			s.OutputFormat = "jpeg"
		}
	}

	if val, ok := values[sOutputDirectory]; ok && val != nil {
		s.OutputDirectory, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.OutputDirectory == "" {
			s.OutputDirectory = "/tmp/"
		}
	}

	if val, ok := values[sInputDirectory]; ok && val != nil {
		s.InputDirectory, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.InputDirectory == "" {
			s.InputDirectory = "/tmp/"
		}
	}

	if val, ok := values[sImageSize]; ok && val != nil {
		s.ImageSize, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.ImageSize == "" {
			s.ImageSize = "auto"
		}
	}

	if val, ok := values[sQuality]; ok && val != nil {
		s.Quality, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Quality == "" {
			s.Quality = "auto"
		}
	}

	if val, ok := values[sCompression]; ok && val != nil {
		s.Compression, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sTransparent]; ok && val != nil {
		s.Transparent, err = coerce.ToBool(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[sModeration]; ok && val != nil {
		s.Moderation, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Moderation == "" {
			s.Moderation = "auto"
		}
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	Prompt         string `md:"prompt, required"`
	InputFileName  string `md:"inputFilename"`
	MaskFileName   string `md:"maskFilename"`
	OutputFileName string `md:"outputFilename"`
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

	i.InputFileName, err = coerce.ToString(values[iInputFile])
	if err != nil {
		return err
	}

	i.MaskFileName, err = coerce.ToString(values[iMaskFile])
	if err != nil {
		return err
	}

	i.OutputFileName, err = coerce.ToString(values[iOutputFileName])
	if err != nil {
		return err
	}

	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iPrompt:         i.Prompt,
		iInputFile:      i.InputFileName,
		iMaskFile:       i.MaskFileName,
		iOutputFileName: i.OutputFileName,
	}
}

// Output defines what data the activity returns
type Output struct {
	MetaData      string `md:"metaData"`
	OutputFileURL string `md:"outputFileURL"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oMetaData:      o.MetaData,
		oOutputFileURL: o.OutputFileURL,
	}
}

// FromMap populates the struct from a map.
func (o *Output) FromMap(values map[string]interface{}) error {
	if values == nil {
		return nil
	}

	var err error
	if val, ok := values[oMetaData]; ok && val != nil {
		o.MetaData, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[oOutputFileURL]; ok && val != nil {
		o.OutputFileURL, err = coerce.ToString(val)
		if err != nil {
			return err
		}
	}
	return nil
}
