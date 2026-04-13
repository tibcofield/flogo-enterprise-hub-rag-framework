package modelOptions

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
	iNumCTX        = "numCTX"
	iRepeastLastN  = "repeastLastN"
	iRepeatPenalty = "repeatPenalty"
	iTemperature   = "temperature"
	iSeed          = "seed"
	iStop          = "stop"
	iNumPredict    = "numPredict"
	iTopK          = "topK"
	itopP          = "topP"
	iminP          = "minP"
	oModelOptions  = "modelOptions"
)

// Settings defines configuration options for your activity
type Settings struct {
}

// FromMap populates the settings struct from a map.
func (s *Settings) FromMap(values map[string]interface{}) error {

	return nil
}

// Input defines what data the activity receives
type Input struct {
	NumCTX        int           `md:"numCTX"`
	RepeastLastN  int           `md:"repeastLastN"`
	RepeatPenalty float32       `md:"repeatPenalty"`
	Temperature   float32       `md:"temperature"`
	Seed          int           `md:"seed"`
	Stop          []interface{} `md:"stop"`
	NumPredict    int           `md:"numPredict"`
	TopK          int           `md:"topK"`
	TopP          float32       `md:"topP"`
	MinP          float32       `md:"minP"`
}

// FromMap populates the struct from the activity's inputs.
func (i *Input) FromMap(values map[string]interface{}) error {

	if values == nil {
		return nil
	}

	// Todo Refactor this code to make efficient.
	var err error

	if val, ok := values[iNumCTX]; ok && val != nil {
		i.NumCTX, err = coerce.ToInt(values[iNumCTX])
		if err != nil {
			return err
		}
	}

	if val, ok := values[iRepeastLastN]; ok && val != nil {
		i.RepeastLastN, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iRepeatPenalty]; ok && val != nil {
		i.RepeatPenalty, err = coerce.ToFloat32(val)

		if err != nil {
			return err
		}
	}
	if val, ok := values[iTemperature]; ok && val != nil {
		i.Temperature, err = coerce.ToFloat32(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iSeed]; ok && val != nil {
		i.Seed, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iStop]; ok && val != nil {
		i.Stop, err = coerce.ToArray(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iNumPredict]; ok && val != nil {
		i.NumPredict, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iTopK]; ok && val != nil {
		i.TopK, err = coerce.ToInt(val)
		if err != nil {
			return err
		}
	}

	if val, ok := values[itopP]; ok && val != nil {

		i.TopP, err = coerce.ToFloat32(val)
		if err != nil {
			return err
		}
	}
	if val, ok := values[iminP]; ok && val != nil {
		i.MinP, err = coerce.ToFloat32(val)
		if err != nil {
			return err
		}
	}
	return nil
}

// ToMap converts the struct to a map.
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		iNumCTX:        i.NumCTX,
		iRepeastLastN:  i.RepeastLastN,
		iRepeatPenalty: i.RepeatPenalty,
		iTemperature:   i.Temperature,
		iSeed:          i.Seed,
		iStop:          i.Stop,
		iNumPredict:    i.NumPredict,
		iTopK:          i.TopK,
		itopP:          i.TopP,
		iminP:          i.MinP,
	}
}

// Output defines what data the activity returns
type Output struct {
	ModelOptions string `md:"modelOptions"`
}

// ToMap converts the struct to a map.
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		oModelOptions: o.ModelOptions,
	}
}

// FromMap populates the struct from a map.
func (o *Output) FromMap(values map[string]interface{}) error {
	if values == nil {
		return nil
	}

	var err error
	o.ModelOptions, err = coerce.ToString(values["oModelOptions"])
	if err != nil {
		return err
	}

	return nil
}
