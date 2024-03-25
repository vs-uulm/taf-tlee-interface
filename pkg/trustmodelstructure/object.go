package trustmodelstructure

// Represents a trust model object
//
// ID: The unique identifier of the object
// Operator: The operator used to fuse multiple trust opinions (corresponding to multiple parallel outgoing relations from the object)
// Relations: The outgoing trust relations to other objects
type Object struct {
	ID        string
	Operator  string
	Relations []Relation
}
