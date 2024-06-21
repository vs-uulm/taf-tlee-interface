package trustmodelstructure

import "github.com/vs-uulm/go-subjectivelogic/pkg/subjectivelogic"

// The TrustRelationship between a Trustor (getSource()) and the Trustee (getDestination()) and its Trust Opinion. A TrustRelationship is always bound to a specific scope it is defined in.
type TrustRelationship interface {
	Source() string
	Destination() string
	Opinion() subjectivelogic.QueryableOpinion
}
