# TLEE Interface

Interface for defining TLEE/TAF interface.

## Implementing the Interfaces

As Go interfaces are implemented implicitly, the following steps would be necessary to get the interface running:

Add this module to the project in which the TLEE is implemented. Given the fact that the module is not published under its name, you can do so by checking out this repository next to the tlee-implementation: 

```
.
├── tlee-implementation
└── tlee-interface
```

Then, use the following code in the `go.mod` file of your tlee-implementation project:
```
replace (
	github.com/vs-uulm/taf-tlee-interface => ../tlee-interface
)
```

This will now allow you to access the packages `tleeinterface` and `trustmodelstructure` from within your code.

Next, you need to add functions to your structs so they will implicitly implement the interfaces. 
This will probably include the following:

| Generic Interface in tlee-interface                            | Struct in tlee-implementation |
|----------------------------------------------------------------| -- |
| `trustmodelstructure.TrustGraphStructure`                      |  `StructureGraphTAFMultipleProp` |
| `trustmodelstructure.AdjacencyListEntry`                       |  `VertexEdgeDTO` |
| `trustmodelstructure.TrustRelationship`                        |  `OpinionDTO` |
| `subjectivelogic.QueryableOpinion` from [go-subjectivelogic](https://github.com/vs-uulm/go-subjectivelogic/blob/0eed3084529279af0be08f5bc0de270f3028ae7b/pkg/subjectivelogic/Opinion.go#L44) |  `OpinionDTOValue` |

Furthermore, you will need to adapt your code to handle the input parameters that it will no longer by the structs on the right side that are used, but rather the interfaces from left side.
Alternatively, you can implement some wrapper code that converts the interface-typed inputs into your own internal DTOs at the beginning of the `RunTLEE()` implementation.