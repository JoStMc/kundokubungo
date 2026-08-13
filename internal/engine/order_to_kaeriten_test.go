package engine

import (
	"log"
	"reflect"
	"testing"
)

func TestGetSequences(t *testing.T) {
	tests := map[string]struct {
		input []int
		output []int
	}{
		"one": {input: []int{1}, output: []int{1}},
		"in order": {input: []int{1, 2, 3}, output: []int{3, 2, 1}},
		"reverse order": {input: []int{3, 2, 1}, output: []int{1}},
		"long": {input: []int{1, 4, 3, 2, 5, 9, 8, 6, 7}, output: []int{7, 6, 5, 2, 1}},
	} 

	for name, tc := range tests {
		k := kaeri{
		    order: tc.input,
		} 
		err := k.getPositionsAsIndex()
		if err != nil {
			log.Fatal(err)
		}
		k.getSequences()

		if !reflect.DeepEqual(tc.output, k.sequenceEnds) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.output, k.sequenceEnds)
		} 
	} 
} 


func TestAddRes(t *testing.T) {
	tests := map[string]struct {
		input []int
		output []int
	}{
		"one": {input: []int{1}, output: []int{1}},
		"in order": {input: []int{1, 2, 3}, output: []int{3, 2, 1}},
		"reverse order": {input: []int{3, 2, 1}, output: []int{1}},
		"long": {input: []int{1, 4, 3, 2, 5, 9, 8, 6, 7}, output: []int{8, 6, 5, 2, 1}},
		"several": {input: []int{6, 5, 4, 2, 1, 3}, output: []int{4, 1}},
	}

	for name, tc := range tests {
		k := kaeri{
			kaeriten: make([]string, len(tc.input)),
		    order: tc.input,
		}
		err := k.getPositionsAsIndex()
		if err != nil {
			log.Fatal(err)
		}
		k.getSequences()
		k.addRes()

		if !reflect.DeepEqual(tc.output, k.sequenceStarts) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.output, k.sequenceStarts)
		}
	}
} 

func TestGetSequenceDepths(t *testing.T) {
	tests := map[string]struct {
		input []int
		output []int
	}{
		"one": {input: []int{1}, output: []int{0}},
		"in order": {input: []int{1, 2, 3}, output: []int{0, 0, 0}},
		"reverse order": {input: []int{3, 2, 1}, output: []int{0}},
		"layered": {input: []int{10, 9, 7, 5, 3, 2, 1, 4, 6, 8}, output: []int{3, 2, 1, 0}},
		"separated": {input: []int{3, 1, 2, 6, 4, 5}, output: []int{1, 0, 1, 0}},
		"end before": {input: []int{5, 2, 3, 4, 1}, output: []int{1, 0, 0}}, 
		"start after": {input: []int{2, 4, 1, 3}, output: []int{1, 0}}, 
	}

	for name, tc := range tests {
		k := kaeri{
			kaeriten: make([]string, len(tc.input)),
		    order: tc.input,
		} 
		err := k.getPositionsAsIndex()
		if err != nil {
			log.Fatal(err)
		}
		k.getSequences()
		k.addRes()
		k.getSequenceDepths()

		if !reflect.DeepEqual(tc.output, k.sequenceDepths) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.output, k.sequenceDepths)
		} 
	} 

}
