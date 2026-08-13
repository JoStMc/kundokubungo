package engine

import (
	"errors"

	"github.com/JoStMc/kundokubungo/internal/models"
)

type kaeri struct {
	kaeriten []string
	order []int
	positions []int
	sequenceEnds []int
	sequenceStarts []int
	sequenceDepths []int
} 

/* Plan:
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
	sequenceStarts := []int{}
	sequenceEnds := []int{}
	current := len(k.order)
	for current != 0 {
		sequenceStarts = append(sequenceStarts, current)
		for k.positions[current-1] > k.positions[current] {
		    current -= 1
		} 
		sequenceEnds = append(sequenceEnds, current)
		current -= 1
	} 
	k.sequenceEnds = sequenceEnds
	k.sequenceStarts = sequenceStarts
} 

// Adds レ to all contiguous descending sequences until there is at least 
// one character between them. 21 gives 2 レ, 54312 gives 5 and 4 レ
func (k *kaeri) addRes() {
	for i, start := range k.sequenceStarts {
		end := k.sequenceEnds[i]
		for current := start; current > end; current-- {
			if k.positions[current-1] != k.positions[current] + 1 {
			    break
			} 
			k.kaeriten[k.positions[current]] = models.MarkRe
			k.sequenceStarts[i]--
		} 
	} 
}

func (k *kaeri) getSequenceDepths() {
	k.sequenceDepths = make([]int, len(k.sequenceEnds))
	for i := len(k.sequenceDepths) - 1; i >= 0; i-- {
		currentStartPos := k.positions[k.sequenceStarts[i]]
		currentEndPos := k.positions[k.sequenceEnds[i]]

		deepest := 0
		for j := i+1; j < len(k.sequenceDepths); j++ {
			jStartPos := k.positions[k.sequenceStarts[j]]
			jEndPos := k.positions[k.sequenceEnds[j]]
			if (jStartPos >= currentStartPos && jStartPos <= currentEndPos) ||
			(jEndPos >= currentStartPos && jEndPos <= currentEndPos) {
				if deepest < k.sequenceDepths[j] + 1 {
					deepest = k.sequenceDepths[j] + 1
				} 
			} 
		} 
		k.sequenceDepths[i] = deepest
	} 
}
