package engine

import (
	"errors"

	"github.com/JoStMc/kundokubungo/internal/models"
)

// To be refactored

// kaeriten = characters
// order = desired output order
// positions = index of where each number is in order[]
type kaeri struct {
	kaeriten []string
	order []int
	positions []int

	sequenceEnds []int
	sequenceStarts []int
	sequenceDepths []int

	sequenceParents []int
	sequenceChildren []int
} 

func GetKaeriten(order []int) ([]string, error) {
	k := kaeri{
		kaeriten: make([]string, len(order)),
		order: order,
	}
	err := k.getPositionsAsIndex()
	if err != nil {
		return []string{}, err
	}

	k.getSequences()
	k.addRes()
	k.getSequenceDepths()
	k.pushDepth()
	k.setSequenceKaeri()

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

// Loops backwards through the sequences to see if any of the later sequences
// begin or end within the sequence.
// Sets the parent and child of each sequence accordingly to the index of the
// the sequence which is the parent/child
func (k *kaeri) getSequenceDepths() {
	k.sequenceDepths = make([]int, len(k.sequenceEnds))
	k.sequenceChildren = make([]int, len(k.sequenceEnds))
	k.sequenceParents = make([]int, len(k.sequenceEnds))

	for i := len(k.sequenceDepths) - 1; i >= 0; i-- {
		k.sequenceParents[i] = -1
		currentStartPos := k.positions[k.sequenceStarts[i]]
		currentEndPos := k.positions[k.sequenceEnds[i]]

		deepest := 0
		deepestIdx := -1
		for j := i+1; j < len(k.sequenceDepths); j++ {
			jStartPos := k.positions[k.sequenceStarts[j]]
			jEndPos := k.positions[k.sequenceEnds[j]]
			if (jStartPos >= currentStartPos && jStartPos <= currentEndPos) ||
			(jEndPos >= currentStartPos && jEndPos <= currentEndPos) {
				if deepest < k.sequenceDepths[j] + 1 {
					deepest = k.sequenceDepths[j] + 1
					deepestIdx = j
				}
			}
		}
		k.sequenceDepths[i] = deepest
		if deepestIdx == -1 {
		    k.sequenceChildren[i] = -1
		    continue
		} 
		k.sequenceParents[deepestIdx] = i
	}
}

// Increases depths of any sequence which is too long for the specific depth
// e.g. level 2 is 上中下, but if the sequence is 4 characters, it should be 
// moved to 甲乙丙丁, and 天地人 (level 4) to 元亨利貞
func (k *kaeri) pushDepth() {
	for i, depth := range k.sequenceDepths {
		seqLen := k.sequenceStarts[i] - k.sequenceEnds[i] + 1
		if depth == 2 && seqLen < 4 {
			k.sequenceDepths[i]++
			parent := k.sequenceParents[i]
			for parent != -1 {
				k.sequenceDepths[parent]++
				parent = k.sequenceParents[parent]
			}
		}
	}
	for i, depth := range k.sequenceDepths {
		seqLen := k.sequenceStarts[i] - k.sequenceEnds[i] + 1
		if depth == 4 && seqLen < 4 {
			k.sequenceDepths[i]++
			parent := k.sequenceParents[i]
			for parent != -1 {
				k.sequenceDepths[parent]++
				parent = k.sequenceParents[parent]
			}
		}
	}
}

func (k *kaeri) setSequenceKaeri() {
	for i, seqSt := range k.sequenceStarts {
		seqEnd := k.sequenceEnds[i]
		if seqSt == seqEnd {
		    continue
		} 

		curMark := ""
		switch k.sequenceDepths[i] {
		case 1: curMark = "一"
		case 2: curMark = "上"
		case 3: curMark = "甲"
		case 4: curMark = "天"
		case 5: curMark = "元"
		case 6: curMark = "乾"
		}

		// ADD: check for if sequence is too long
		for s := seqSt; s <= seqEnd; s++ {
			k.kaeriten[k.positions[s]] = curMark
			curMark = nextMarks[curMark]
		} 
	} 
} 
