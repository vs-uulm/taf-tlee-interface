package tleeinterface

import (
	"github.com/vs-uulm/go-subjectivelogic/pkg/subjectivelogic"
	"github.com/vs-uulm/taf-tlee-interface/pkg/trustmodelstructure"
)

/*
The RunTLEE function provided by the TLEE to be used by the TAF.
*/
type TLEE interface {
	RunTLEE(trustmodelID string, version int, fingerprint uint32, structure trustmodelstructure.TrustGraphStructure, values map[string][]trustmodelstructure.TrustRelationship) map[string]subjectivelogic.QueryableOpinion
}
