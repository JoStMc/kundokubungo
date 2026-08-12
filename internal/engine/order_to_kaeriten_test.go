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

		if !reflect.DeepEqual(tc.output, k.sequences) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.output, k.sequences)
		} 
	} 
} 

