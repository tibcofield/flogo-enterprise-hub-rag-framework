package uploadFile

/*
* Copyright © 2023 - 2025. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	"strconv"

	"github.com/project-flogo/core/data/coerce"
)

// KeyValuePair defines the structure for each item in our input array
type FileAttributeData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Constants for identifying settings and inputs
const (
	sAPIKey             = "apiKey"
	sEnpointURL         = "endPointURL"
	sPurpose            = "purpose"
	sVectorStoreID      = "vectorStoreID"
	sMaxChunkSizeTokens = "maxChunkSizeTokens"
	sChunkOverlapTokens = "chunkOverlapTokens"
	sTimeoutSeconds     = "timeoutSeconds"
	iFilename           = "filename"
	iFileAttributeNames = "fileAttributeNames"
	iFileAttributes     = "fileAttributes"
	oMetaData           = "metaData"
)

// Settings defines configuration options for your activity
type Settings struct {
	ApiKey             string `md:"apiKey, required"`
	EndPointURL        string `md:"endPointURL, required"`
	Purpose            string `md:"purpose, required"`
	VectorStoreID      string `md:"string"`
	MaxChunkSizeTokens int64  `md:"maxChunkSizeTokens"`
	ChunkOverlapTokens int64  `md:"chunkOverlapTokens"`
	TimeoutSeconds     int    `md:"timeoutSeconds"`
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {
	if values == nil {
		s.Purpose = "assistants"
		s.TimeoutSeconds = 300 // Default 5 minutes for large files
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

	if val, ok := values[sPurpose]; ok && val != nil {
		s.Purpose, err = coerce.ToString(val)
		if err != nil {
			return err
		}

		if s.Purpose == "" {
			s.Purpose = "assistants"
		}
	}

	s.VectorStoreID, err = coerce.ToString(values[sVectorStoreID])
	if err != nil {
		return err
	}

	s.MaxChunkSizeTokens, err = coerce.ToInt64(values[sMaxChunkSizeTokens])
	if err != nil {
		return err
	}

	s.ChunkOverlapTokens, err = coerce.ToInt64(values[sChunkOverlapTokens])
	if err != nil {
		return err
	}

	s.TimeoutSeconds, err = coerce.ToInt(values[sTimeoutSeconds])
	if err != nil {
		return err
	}

	return nil
}

// Input defines what data the activity receives
type Input struct {
	FileName           string              `md:"filename"`
	FileAttributeNames []interface{}       `md:"fileAttributeNames"`
	FileAttributes     []FileAttributeData `md:"fileAttributes"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	i.FileName, err = coerce.ToString(values[iFilename])
	if err != nil {
		return err
	}

	i.FileAttributeNames, err = coerce.ToArray(values[iFileAttributeNames])
	if err != nil {
		return err
	}

	var fileAttributeTemp []interface{}

	fileAttributeTemp, err = coerce.ToArray(values[iFileAttributes])

	logger.Info("File has attributes:" + strconv.Itoa((len(fileAttributeTemp))))

	if err != nil {
		return err
	} else {
		i.FileAttributes = make([]FileAttributeData, 0, len(fileAttributeTemp))
		for _, metaRow := range fileAttributeTemp {
			if m, ok := metaRow.(map[string]interface{}); ok {
				key, _ := coerce.ToString(m["key"])
				value, _ := coerce.ToString(m["value"])
				i.FileAttributes = append(i.FileAttributes, FileAttributeData{
					Key:   key,
					Value: value,
				})
				attributekey, _ := coerce.ToString(m["key"])
				logger.Info("Adding " + attributekey)
			}
		}
	}

	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iFilename:           i.FileName,
		iFileAttributeNames: i.FileAttributeNames,
		iFileAttributes:     i.FileAttributes,
	}
}

// Output defines what data the activity returns
type Output struct {
	ID            string `md:"id"`
	MetaData      string `md:"metaData"`
	OutputFileURL string `md:"outputFileURL"`
	ObjectID      string `md:"object"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oMetaData: o.MetaData,
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

	return nil
}
