package sampleflogo

import (
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&SumActivity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type SumActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new TwilioActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &SumActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *SumActivity) Metadata() *activity.Metadata {
	return activity.ToMetadata(&Settings{}, &Input{}, &Output{})
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *SumActivity) Eval(ctx activity.Context) (done bool, err error) {
	input := &Input{}
	err = ctx.GetInputObject(input)
	if err != nil {
		return true, err
	}

	ans := input.A + input.B

	output := &Output{
		Ans: ans,
	}
	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	return true, nil
}
