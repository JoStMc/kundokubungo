package engine

import (
	"github.com/JoStMc/kundokubungo/internal/models"
)

// Here we define what to do for each kaeriten
// Recently everything has been generalised, so this may be redundant now
var kaeritenTypes = map[string]func(*config, int) {
	models.MarkRe: (*config).reten,
	"": (*config).allChars,
} 


type config struct {
	sentence []models.Character
	marks map[string]int
	order []int
	currentChar int
} 

// This simply loops through the characters to define what their position
// in the sentence should be. 
// If characters [A, B, C, D] were to return [1, 0, 3, 2], the character 
// order should be BADC. 
func getCharOrder(sentence *models.Sentence) ([]int, error) {
	characters := sentence.Characters

	cfg := config{
		sentence: characters,
		marks: make(map[string]int),
		order: make([]int, len(characters)),
		currentChar: 0,
	} 

	for i, char := range characters {
		kaeriFunc, ok := kaeritenTypes[char.Kaeriten]
		if !ok {
			kaeriFunc = (*config).sequenceFunc
		} 
		kaeriFunc(&cfg, i)
	} 

	return cfg.order, nil
} 

func (cfg *config) allChars(index int) {
	cfg.order[index] = cfg.currentChar
	cfg.currentChar++
	if index != 0 && cfg.sentence[index-1].Kaeriten == models.MarkRe {
		cfg.allChars(index-1)
	} 
} 

func (cfg *config) reten(index int) {
	// This function should do nothing, because allChars should
	// catch reten by checking if the previous char is reten.
	// We don't want to add 1 to the current char
} 

// The next function is a generic function for 一二三, 甲乙丙丁, 元亨利貞 or
// other kaeriten which work by returning to those characters sequentially
// 上中下 need special treatment because of the potential for just 上下
func (cfg *config) sequenceFunc(index int) {
	curMark := cfg.sentence[index].Kaeriten
	cfg.marks[curMark] = index
	prevMark, notFirst := previousMarks[curMark]
	nextMark, notLast := nextMarks[curMark]

	// Following: if 上 and no 中, check 下
	// Then if 下 and no 中, check 上
	// I do not believe there are other marks which skip like this,
	// but recursive checking could be implemented if need arises
	if _, ok := cfg.marks[nextMark]; !ok {
		nextMark = nextMarks[nextMark]
	} 
	if _, ok := cfg.marks[prevMark]; !ok {
		prevMark = previousMarks[prevMark]
	} 

	if !notFirst {
		cfg.allChars(index)
		cfg.sequenceFunc(cfg.marks[nextMark])
		delete(cfg.marks, curMark)
	} else {
		if _, ok := cfg.marks[prevMark]; ok {
			cfg.allChars(index)
			if notLast {
				if nextIndex, ok := cfg.marks[nextMark]; ok {
					cfg.sequenceFunc(nextIndex)
				} 
			} 
			delete(cfg.marks, curMark)
		} 
	}
} 

var previousMarks = map[string]string{
	"二": "一",
	"三": "二",

	"乙": "甲",
	"丙": "乙",
	"丁": "丙",

	"中": "上",
	"下": "中",
}
var nextMarks = map[string]string{
	"一":"二",
	"二":"三",

	"甲":"乙",
	"乙":"丙",
	"丙":"丁",

	"上": "中",
	"中": "下",
} 
