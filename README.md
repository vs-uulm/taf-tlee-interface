# TLEE Interface

Interface for testing TLEE functionality.

## Testing a TLEE implementation

**In order to test a TLEE implementation, simply replace the TLEE placeholder with an actual TLEE implementation.** 
This can be done in several ways:
1. Replace the logic inside the *RunTLEE* function in the *tlee* package (Especially suitable for testing very simple, minimalistic TLEE implementations). 
2. Replace the *tlee* package with another tlee package that implements the *RunTLEE* function.

However, the *RunTLEE* function is required to have the following signature:
```go
func RunTLEE(trustmodelID string, version int, fingerprint uint32, structure trustmodelstructure.Structure, values map[string]subjectivelogic.Opinion) map[string]float64

```

## How it works

In `main.go`, three test sets are created to call the *RunTLEE* function with. The values returned by the *RunTLEE* function are currently not inspected automatically, but only printed to console.
 
### trust model ID and version
The trust model ID and the version are random values. The trust model ID is based on a timestamp in order to provide uniqueness.

### fingerprint
The fingerprint is based on the trust model structure topology. The relations in the corresponding trust model structure are expressed in the form "source,target,relationID". These relation strings are then sorted alphabetically and concatenated into a single string, which is then hashed.

### trust model structures
The trust model structures used in these test sets are currently hardcoded.

### values (trust opinions)
For each relation in the corresponding trust model structure, a random trust opinion is generated.



## *trustmodelstructure* package

Provides data structures and methods to represent and manipulate the structure of trust models.
This package **does not contain trust opinions**, but mainly the structure properties of trust models.

### Objects
Objects represent trust objects of a trust model structure. Each object has a unique identifier, a list of relations to other objects, and an operator string to optionally define an operator used for fusing associated trust opinions.

### Relations
Relations represent the trust relations one trust object has to another. Each relation specifies an unique identifyer and its target objects identifyer.

### Structure
The Structure type is a list of trust objects. It provides methods to convert the structure to a string, and to get partial trust model structures per leaf node object.


## *subjectivelogic* package

Provides data structure to represent a subjective logic opinion as defined by [Audun Josang](https://link.springer.com/book/10.1007/978-3-319-42337-1).