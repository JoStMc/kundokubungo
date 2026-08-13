package engine

import (
	"errors"

	"github.com/JoStMc/kundokubungo/internal/models"
)

type kaeri struct {
	kaeriten []string
	order []int
	positions []int
	sequence_ends []int
	sequence_starts []int
} 

/* Plan:
- When the sequence decreases contiguously, レ can be used until it doesn't
- Need to add way to check for embedded
*/
func getKaeriten(order []int) ([]string, error) {
	k := kaeri{
		kaeriten: make([]string, len(order)),
		order: order,
	} 
	err := k.getPositionsAsIndex()
	if err != nil {
		return []string{}, err
	}



	return k.kaeriten, nil
} 

// Reverses the order array, so the position can be found of each number
// e.g. if order[4] is 7, positions[7] is 4
func (k *kaeri) getPositionsAsIndex() error {
	positions := make([]int, len(k.order) + 1)
	for i, pos := range k.order {
		if pos > len(k.order) {
			return errors.New("invalid order: too few characters")
		} 
		positions[pos] = i
	} 
	k.positions = positions
	return nil
}

// Takes the order and returns at what number the sequence ends
// For example, 143259867 would return a list of the last number in 
// each descending sequence (of 987, 6, 5, 432, 1) , 7, 6, 5, 2, 1
func (k *kaeri) getSequences() {
	sequence_starts := []int{}
	sequence_ends := []int{}
	current := len(k.order)
	for current != 0 {
		sequence_starts = append(sequence_starts, current)
		for k.positions[current-1] > k.positions[current] {
		    current -= 1
		} 
		sequence_ends = append(sequence_ends, current)
		current -= 1
	} 
	k.sequence_ends = sequence_ends
	k.sequence_starts = sequence_starts
} 

// Adds レ to all contiguous descending sequences until there is at least 
// one character between them. 21 gives 2 レ, 54312 gives 5 and 4 レ
func (k *kaeri) addRes() {
	for i, start := range k.sequence_starts {
		end := k.sequence_ends[i]
		for current := start; current > end; current-- {
			if k.positions[current-1] != k.positions[current] + 1 {
			    break
			} 
			k.kaeriten[k.positions[current]] = models.MarkRe
			k.sequence_starts[i]--
		} 
	} 
}
