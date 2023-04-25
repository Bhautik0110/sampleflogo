package sampleflogo

import "github.com/project-flogo/core/data/coerce"

type Settings struct {
}

type Input struct {
	A int `md:"a,required"`
	B int `md:"b,required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {
	a, _ := coerce.ToInt(values["a"])
	b, _ := coerce.ToInt(values["b"])
	r.A = a
	r.B = b
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"a": r.A,
		"b": r.B,
	}
}

type Output struct {
	Ans int `md:"ans"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	ans, _ := coerce.ToInt(values["ans"])
	o.Ans = ans
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"ans": o.Ans,
	}
}
