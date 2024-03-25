/*
Package trustmodelstructure provides data structures and methods to represent and manipulate the structure of trust models.
This package **does not contain trust opinions**, but mainly the structure properties of trust models.

# Objects

Objects represent trust objects of a trust model structure. Each object has a unique identifier, a list of relations to other objects, and an operator string to optionally define an operator used for fusing associated trust opinions.

# Relations

Relations represent the trust relations one trust object has to another. Each relation specifies an unique identifyer and its target objects identifyer.

# Structure

The Structure type is a list of trust objects. It provides methods to convert the structure to a string, and to get partial trust model structures per leaf node object.
*/
package trustmodelstructure
