package tlee

import (
	subjectivelogic "gitlab-vs.informatik.uni-ulm.de/connect/tlee-interface/pkg/subjectivelogic"
	"gitlab-vs.informatik.uni-ulm.de/connect/tlee-interface/pkg/trustmodelstructure"
)

func RunTLEE(trustmodelID string, version int, fingerprint uint32, structure trustmodelstructure.Structure, values map[string]subjectivelogic.Opinion) map[string]subjectivelogic.Opinion {
	//TODO: Replace with functional TLEE implementation
	return make(map[string]subjectivelogic.Opinion)
}
