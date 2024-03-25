package trustmodelstructure

// Structure represents a trust model structure
type Structure []Object

// Returns a string representation of the trust model structure in the format of
// "objectID -> targetID, targetID, ..."
func (s Structure) ToString() string {
	var str string
	for _, object := range s {
		str += object.ID + " -> "
		for _, relation := range object.Relations {
			str += relation.Target + ", "
		}
		str += "\n"
	}
	return str
}
