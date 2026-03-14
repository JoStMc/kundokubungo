package engine

import (
	"testing"

	"github.com/JoStMc/kundokubungo/internal/models"
)

func TestKakikudashi(t *testing.T) {
	tests := map[string]struct {
		input models.Sentence
		output string
	}{
		"re": {input: models.Sentence{Characters: []models.Character{
							{Kanji: "登", Kaeriten: "レ", Okurigana: "ル"}, 
							{Kanji: "山", Okurigana: "二"}}},
						output: "山二登ル"},

		"ichini": {input: models.Sentence{Characters: []models.Character{
							{Kanji: "見", Kaeriten: "二", Okurigana: "ル"},
							{Kanji: "南"},
							{Kanji: "山", Kaeriten: "一", Okurigana: "ヲ"}}}, 
						output: "南山ヲ見ル"}, 
		"hyphen": {input: models.Sentence{Characters: []models.Character{
							{Kanji: "三", IsJukugoHead: true, Kaeriten: "二"},
							{Kanji: "省", IsJukugoTail: true, Okurigana: "ス"},
							{Kanji: "吾", Okurigana: "ガ"},
							{Kanji: "身" , Kaeriten: "一", Okurigana: "ヲ"}}}, 
						output: "吾ガ身ヲ三省ス"},
		"jouge": {input: models.Sentence{Characters: []models.Character{
							{Kanji: "悪", Kaeriten: "下", Okurigana: "ム"},
							{Kanji: "称", Kaeriten: "二", Okurigana: "スル"},
							{Kanji: "人"},
							{Kanji: "之"},
							{Kanji: "悪", Kaeriten: "一", Okurigana: "ヲ"}, 
							{Kanji: "者", Kaeriten: "上", Okurigana: "ヲ"}}},
						output: "人之悪ヲ称スル者ヲ悪ム"},
		"kouotsu": {input: models.Sentence{Characters: []models.Character{
							{Kanji: "有", Kaeriten: "乙", Okurigana: "リ"},
							{Kanji: "難", Kaeriten: "下", Okurigana: "キ"},
							{Kanji: "以", Kaeriten: "二", Okurigana: "テ"},
							{Kanji: "里",},
							{Kanji: "数", Kaeriten: "一", Okurigana: "ヲ"},
							{Kanji: "限", Kaeriten: "上", Okurigana: "リ"},
							{Kanji: "者", Kaeriten: "甲"},
							{Kanji: "矣",},
		}},
						output: "里数ヲ以テ限リ難キ者有リ矣",
		},
		"ichire": {input: models.Sentence{Characters: []models.Character{
							{Kanji: "璧", Okurigana: "ハ"},
							{Kanji: "不", Kaeriten: "レ"},
							{Kanji: "可", Kaeriten: "二", Okurigana: "カラ"},
							{Kanji: "以", Okurigana: "テ"},
							{Kanji: "禦", Kaeriten: "一レ", Okurigana: "グ"},
							{Kanji: "寒", Okurigana: "キヲ"}}},
						output: "璧ハ以テ寒キヲ禦グ可カラ不"}, 
	} 

	for name, tc := range tests {
		got, err := ToKakikudashi(&tc.input)
		if err != nil {
		    t.Fatal("Error parsing sentence", err)
		} 
		if tc.output != got {
			t.Fatalf("%s: expected: %v, got: %v", name, tc.output, got)
		} 
	} 
} 
