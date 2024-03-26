package tleeinterface

import (
	subjectivelogic "github.com/vs-uulm/taf-tlee-interface/pkg/subjectivelogic"
	"github.com/vs-uulm/taf-tlee-interface/pkg/trustmodelstructure"
)

func RunTLEE(trustmodelID string, version int, fingerprint uint32, structure trustmodelstructure.Structure, values map[string]subjectivelogic.Opinion) map[string]subjectivelogic.Opinion {
	//TODO: Replace with functional TLEE implementation
	return make(map[string]subjectivelogic.Opinion)
}
