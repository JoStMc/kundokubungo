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
		"single ichiniten":	{input: "二 一", output: []int{1,2,0}},
		"ichinisan":	{input: "三   二 一", output: []int{1,2,3,5,6,4,0}},
		"re and nums": {input: "レ二 一レ二 一", output: []int{2,3,1,0,6,7,5,4}}, 
		"kououtheitei basic": {input: "丁 丙 乙 甲", output: []int{1,3,5,6,4,2,0}}, 
		"kouotsuheitei w/ ichinisan": {input: "丁       丙  二 一  乙 甲", output: []int{1,2,3,4,5,6,7,9,10,12,13,11,14,15,17,18,16,8,0}}, 
		"jouchuuge": {input: "下 中  上", output: []int{1,3,4,5,2,0}},
		"jouge": {input: "下  上", output: []int{1,2,3,0}},
		"koujouichi seq": {input: "乙下二 一上甲 ", output: []int{3,4,2,5,1,6,0,7}},
		"tenchi": {input: "人  乙下二 一中 上甲 二 一 地 天  レ 二  一", output: []int{1,2,6,7,5,9,10,8,4,11,3,12,14,15,13,16,18,19,17,0,20,21,23,22,25,26,27,24}},
		"out of order": {input: "十 二 三 四 五 六 七 八 九 一", output: []int{1,3,5,7,9,11,13,15,17,18,2,4,6,8,10,12,14,16,0}},
	} 

	otherTests := map[string]struct {
		input models.Sentence
		output []int
	}{
		"re saidoku": {input: models.Sentence{Characters: []models.Character{
							{}, 
							{IsSaidokumoji: true, Kaeriten: "レ"}, 
							{Kaeriten: "レ"}, 
							{}}},
						output: []int{0,1,3,2,1}},

		"sen saidoku": {input: models.Sentence{Characters: []models.Character{
							{}, {}, 
							{Kaeriten: "レ"}, 
							{}, 
							{IsSaidokumoji: true, Kaeriten: "三"}, 
							{}, {}, 
							{Kaeriten: "二"}, 
							{}, {}, 
							{Kaeriten: "一"}, 
							{}}},
						output: []int{0,1,3,2,4,5,6,8,9,10,7,4,11}}, 
		"basic hyphen": {input: models.Sentence{Characters: []models.Character{
							{IsJukugoHead: true, Kaeriten: "二"},
							{IsJukugoTail: true},
							{},
							{Kaeriten: "一"}}}, 
						output: []int{2,3,0,1}},
		"two hyphens": {input: models.Sentence{Characters: []models.Character{
							{IsJukugoHead: true, Kaeriten: "二"},
							{IsJukugoTail: true},
							{IsJukugoHead: true, Kaeriten: "三"},
							{IsJukugoTail: true},
							{},
							{Kaeriten: "一"}}}, 
						output: []int{4,5,0,1,2,3}},
		"re hyphen": {input: models.Sentence{Characters: []models.Character{
							{IsJukugoHead: true, Kaeriten: "レ"},
							{IsJukugoTail: true},
							{}}},
						output: []int{2,0,1},
		},
		"ichire": {input: models.Sentence{Characters: []models.Character{
							{},
							{Kaeriten: "レ"},
							{Kaeriten: "二"},
							{},
							{Kaeriten: "一レ"},
							{}}},
						output: []int{0,3,5,4,2,1}}, 
		"joure": {input: models.Sentence{Characters: []models.Character{
							{Kaeriten: "下"},
							{Kaeriten: "二"},
							{},
							{Kaeriten: "一"},
							{Kaeriten: "上レ"},
							{}}},
						output: []int{2,3,1,5,4,0}}, 
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

	for name, tc := range otherTests {
		got, err := getCharOrder(&tc.input)
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
