package engine

import (
	"reflect"
	"testing"
	"unicode/utf8"

	"github.com/JoStMc/kundokubungo/internal/models"
)

func TestGetCharOrder(t *testing.T) {
	tests := map[string]struct {
		input string
		output []int
	}{
		"no kaeriten":	{input: "   ", output: []int{0,1,2}},
		"single reten":	{input: "レ ", output: []int{1,0}},
		"single ichiniten":	{input: "二 一", output: []int{2,0,1}},
		"ichinisan":	{input: "三   二 一", output: []int{6,0,1,2,5,3,4}},
		"alternative ichinisan": {input: "  三 二 一", output: []int{0,1,6,2,5,3,4}},
		"re and nums": {input: "レ二 一レ二 一", output: []int{3,2,0,1,7,6,4,5}}, 
	} 

	for name, tc := range tests {
		sentence := makeSentence(&tc.input)
		got, err := getCharOrder(&sentence)
		if err != nil {
		    t.Fatal("Error parsing sentence", err)
		} 
		if !reflect.DeepEqual(tc.output, got) {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.output, got)
		} 
	} 
} 


func makeSentence(input *string) models.Sentence {
    runeCount := utf8.RuneCountInString(*input)
    sentence := models.Sentence{
        Characters: make([]models.Character, runeCount),
    }

    runeIndex := 0
    for _, char := range *input {
        if char == ' ' {
            sentence.Characters[runeIndex].Kaeriten = ""
        } else {
            sentence.Characters[runeIndex].Kaeriten = string(char)
        }
        runeIndex++
    }

    return sentence
}
