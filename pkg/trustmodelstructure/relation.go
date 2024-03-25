package trustmodelstructure

// Relation represents a trust relation one trust object has to another
//
// ID: The unique identifier of the relation
// Target: The unique identifier of the target trust object
type Relation struct {
	ID     string
	Target string
}
