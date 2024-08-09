package trustmodelstructure

type FusionOperator int

// Define the enum values using iota
const (
	AveragingFusion FusionOperator = iota
	ConstraintFusion
	CumulativeFusion
	WeightedFusion
	ConsensusAndCompromiseFusion
)
