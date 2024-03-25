/*
Package subjectivelogic provides a struct to represent opinions in subjective logic.
*/
package subjectivelogic

import "fmt"

// Opinion represents an opinion in subjective logic.
type Opinion struct {
	Belief      float64
	Disbelief   float64
	Uncertainty float64
	BaseRate    float64
}

func (o Opinion) ToString() string {
	str := fmt.Sprintf("Belief: %f, Disbelief: %f, Uncertainty: %f, BaseRate: %f", o.Belief, o.Disbelief, o.Uncertainty, o.BaseRate)
	return str
}
