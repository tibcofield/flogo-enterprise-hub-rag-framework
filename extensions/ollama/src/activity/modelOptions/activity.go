package modelOptions

/*
* Copyright © 2023 - 2025. Cloud Software Group, Inc.
* This file is subject to the license terms contained
* in the license file that is distributed with this file.
 */

import (
	ollama "github.com/ollama/ollama/api"
	"github.com/project-flogo/core/activity"
)

// activityMd is the metadata for the activity.
var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

// Metadata returns the activity's metadata.
func (a *Activity) Metadata() *activity.Metadata {

	return activityMd
}

func init() {
	_ = activity.Register(&Activity{}, New)
}

// Activity is a Ollama API activity
type Activity struct {
	settings *Settings
}

// New creates a new instance of the Activity.
func New(ctx activity.InitContext) (activity.Activity, error) {
	s := &Settings{}
	err := s.FromMap(ctx.Settings())

	if err != nil {
		return nil, err
	}

	return &Activity{
		settings: s,
	}, nil
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	var modelOptions = ollama.DefaultOptions()
	if ctx.GetInput(iNumCTX) != nil {
		modelOptions.NumCtx = ctx.GetInput(iNumCTX).(int)
	}
	if ctx.GetInput(iRepeastLastN) != nil {
		modelOptions.RepeatLastN = ctx.GetInput(iRepeastLastN).(int)
	}
	if ctx.GetInput(iRepeatPenalty) != nil {
		modelOptions.RepeatPenalty = ctx.GetInput(iRepeatPenalty).(float32)
	}
	if ctx.GetInput(iTemperature) != nil {
		modelOptions.Temperature = ctx.GetInput(iTemperature).(float32)
	}
	if ctx.GetInput(iSeed) != nil {
		modelOptions.Seed = ctx.GetInput(iSeed).(int)
	}
	if ctx.GetInput(iStop) != nil {
		modelOptions.Stop = ctx.GetInput(iStop).([]string)
	}
	if ctx.GetInput(iNumPredict) != nil {
		modelOptions.NumPredict = ctx.GetInput(iNumPredict).(int)
	}
	if ctx.GetInput(iTopK) != nil {
		modelOptions.TopK = ctx.GetInput(iTopK).(int)
	}
	if ctx.GetInput(itopP) != nil {
		modelOptions.TopP = ctx.GetInput(itopP).(float32)
	}
	if ctx.GetInput(iminP) != nil {
		modelOptions.MinP = ctx.GetInput(iminP).(float32)
	}

	ctx.SetOutput(oModelOptions, modelOptions)

	return true, nil
}
