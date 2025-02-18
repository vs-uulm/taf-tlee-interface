package trustmodelstructure

// A TrustGraphStructure defines the graph-structural properties of a trust model. It does not define scopes, as scopes are only defined for the values of a graph (i.e., trust opinions)
type TrustGraphStructure interface {
	Operator() FusionOperator
	DiscountOperator() DiscountOperator
	AdjacencyList() []AdjacencyListEntry
}

// A AdjacencyListEntry defines all outgoing edges of a source node by listing the corresponding target nodes of these edges.
type AdjacencyListEntry interface {
	SourceNode() string
	TargetNodes() []string
}
