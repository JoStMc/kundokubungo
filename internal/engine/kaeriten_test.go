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
		"re and nums": {input: "レ二 一レ二 一", output: []int{3,2,0,1,7,6,4,5}}, 
		"kououtheitei basic": {input: "丁 丙 乙 甲", output: []int{6,0,5,1,4,2,3}}, 
		"kouotsuheitei w/ ichinisan": {input: "丁       丙  二 一  乙 甲", output: []int{18,0,1,2,3,4,5,6,17,7,8,11,9,10,12,13,16,14,15}}, 
		"jouchuuge": {input: "下 中  上", output: []int{5,0,4,1,2,3}},
		"jouge": {input: "下  上", output: []int{3,0,1,2}},
		"koujouichi seq": {input: "乙下二 一上甲 ", output: []int{6,4,2,0,1,3,5,7}},
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
