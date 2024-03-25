package main

import (
	"fmt"

	"gitlab-vs.informatik.uni-ulm.de/connect/tlee-interface/pkg/testsetgenerator"
	"gitlab-vs.informatik.uni-ulm.de/connect/tlee-interface/pkg/tlee"
)

func main() {

	// Generate test sets
	testSet, err := testsetgenerator.GenerateTestSet()

	if err != nil {
		fmt.Println("Error generating test sets: ", err)
		return
	}

	// Print test set
	fmt.Println("Generated test set:")
	fmt.Printf("TrustmodelID: %s\n", testSet.TrustmodelID)
	fmt.Printf("Version: %d\n", testSet.Version)
	fmt.Printf("Fingerprint: %d\n", testSet.Fingerprint)
	fmt.Println("Structure:")
	fmt.Println(testSet.Structure.ToString())
	fmt.Println("Values:")
	for key, value := range testSet.Values {
		fmt.Printf("%s: %s\n", key, value.ToString())
	}

	// Run TLEE
	fmt.Printf("Running TLEE for trust model structure %s at version %d\n", testSet.TrustmodelID, testSet.Version)

	// Print results
	fmt.Println("Results:")
	for key, value := range tlee.RunTLEE(testSet.TrustmodelID, testSet.Version, testSet.Fingerprint, testSet.Structure, testSet.Values) {
		fmt.Printf("%s: %s\n", key, value.ToString())
		fmt.Printf("Sum of belief, disbelief, and uncertainty: %f\n", value.Belief+value.Disbelief+value.Uncertainty)
	}

}
