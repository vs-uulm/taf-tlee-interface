/*
Package testsetgenerator provides a function to generate test sets for the TLEE.
The test sets are generated based on hardcoded trust model structures and random trust opinions.
TrustmodelID and version are each random values.
The fingerprint is generated based on the trust model structure topology.
*/
package testsetgenerator

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"

	"gitlab-vs.informatik.uni-ulm.de/connect/tlee-interface/pkg/subjectivelogic"
	"gitlab-vs.informatik.uni-ulm.de/connect/tlee-interface/pkg/trustmodelstructure"
)

// TestSet provides all parameters relevant for calling the TLEE
type TestSet struct {
	TrustmodelID string
	Version      int
	Fingerprint  uint32
	Structure    trustmodelstructure.Structure
	Values       map[string]subjectivelogic.Opinion
}

// GenerateTestSets generates test sets for the TLEE.
// The test set is generated based on a hardcoded trust model structure and random trust opinions.
// TrustmodelID and version are each random values, the fingerprint is based on the trust model structure topology.
// If the fingerprint of two test sets is the same, the trust model structure topology is the same.
func GenerateTestSet() (TestSet, error) {

	trustmodelstructure, err := parseTrustModelStructure("assets/trustmodelstructure1.json")
	if err != nil {
		return TestSet{}, err
	}

	opinions, err := generateTrustOpinions(trustmodelstructure)
	if err != nil {
		return TestSet{}, err
	}

	testSet := TestSet{generateUniqueID(), generateVersion(), generateFingerprint(trustmodelstructure), trustmodelstructure, opinions}

	return testSet, nil

}

// Generates a random trust opinion for each relation in the given trust model structure.
// Returns a map of relation IDs to trust opinions.
func generateTrustOpinions(trustmodelstructure trustmodelstructure.Structure) (map[string]subjectivelogic.Opinion, error) {
	opinions := make(map[string]subjectivelogic.Opinion)

	for _, object := range trustmodelstructure {
		for _, relation := range object.Relations {

			opinion, err := generateRndOpinion()
			if err != nil {
				return nil, err
			}

			opinions[relation.ID] = opinion
		}
	}

	return opinions, nil
}

// Generates a single random trust opinion
func generateRndOpinion() (subjectivelogic.Opinion, error) {
	belief, disbelief, uncertainty, err := generateRndBDUVals()
	if err != nil {
		return subjectivelogic.Opinion{}, err
	}

	baseRate, err := generateRndBRVal()
	if err != nil {
		return subjectivelogic.Opinion{}, err
	}

	opinion := subjectivelogic.Opinion{belief, disbelief, uncertainty, baseRate}

	return opinion, nil

}

// Generates a random float64 value between min and max
func generateRndFloat(min float64, max float64) float64 {
	return rand.Float64()*(max-min) + min
}

// Generates random belief, disbelief and uncertainty values that sum up to 1.0
func generateRndBDUVals() (float64, float64, float64, error) {
	b := generateRndFloat(0.0, 1.0)
	d := generateRndFloat(0.0, (1.0 - b))
	u := 1.0 - b - d

	return b, d, u, nil
}

// Generates a random base rate value
func generateRndBRVal() (float64, error) {
	rndBR := generateRndFloat(0.0, 2.0) / 2.0
	return rndBR, nil
}

// Generates a unique ID based on the current timestamp
func generateUniqueID() string {
	timestamp := time.Now().UnixNano()
	time.Sleep(1 * time.Nanosecond)
	return fmt.Sprint(timestamp)
}

// Generates a random version number between 0 and 9999
func generateVersion() int {
	return rand.Intn(10000)
}

// Generates a fingerprint based on the trust model structure topology
// The relations in the trust model structure are expressed in the form "source,target,relationID".
// These relation strings are then sorted alphabetically and concatenated into a single string, which is then hashed and returned.
func generateFingerprint(structure trustmodelstructure.Structure) uint32 {

	topology := make([]string, 0)

	for _, object := range structure {
		for _, relation := range object.Relations {
			topology = append(topology, strings.ToLower(fmt.Sprintf("%s,%s,%s", object.ID, relation.Target, relation.ID)))
		}
	}

	// sort topology slice
	slices.Sort(topology)

	// join topology slice into a single string
	topologyString := strings.Join(topology, "-")

	// return hash of topology string
	return hash(topologyString)

}

// Hashes a string using the FNV-1a algorithm
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

// parses and returns hardcoded trust model structure
func parseTrustModelStructure(path string) (trustmodelstructure.Structure, error) {

	// Parse given json file into trustmodelstructure.Structure using encoding/json
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var structure trustmodelstructure.Structure
	err = decoder.Decode(&structure)
	if err != nil {
		return nil, err
	}

	return structure, nil

}
