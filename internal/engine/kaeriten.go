package engine

import (
	"slices"

	"github.com/JoStMc/kundokubungo/internal/models"
)

// Here we define what to do for each kaeriten
// Recently everything has been generalised, so this may be redundant now
var kaeritenTypes = map[string]func(*config, int) {
	models.MarkRe: (*config).reten,
	"": (*config).allChars,
	"一": (*config).recursivePull,
	"上": (*config).recursivePull,
	"甲": (*config).recursivePull,
	"天": (*config).recursivePull,
	"元": (*config).recursivePull,
	"乾": (*config).recursivePull,
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
		order: slices.Repeat([]int{-1}, len(characters)),
		currentChar: 0,
	} 

	for i, char := range characters {
		kaeriFunc, ok := kaeritenTypes[char.Kaeriten]
		if !ok {
			kaeriFunc = (*config).saveCharPos
		} 
		kaeriFunc(&cfg, i)
	} 

	return cfg.order, nil
} 

func (cfg *config) allChars(index int) {
	// cfg.order[index] = cfg.currentChar
	cfg.order[cfg.currentChar] = index
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

// The next function is a generic function for 二三, 乙丙丁, 亨利貞 or
// other kaeriten which work by returning to those characters sequentially
func (cfg *config) saveCharPos(index int) {
	curMark := cfg.sentence[index].Kaeriten
	cfg.marks[curMark] = index
} 

// This sets off the loop scanning back over characters
// 上中下 need special treatment because of the potential for just 上下
func (cfg *config) recursivePull(index int) {
	cfg.allChars(index)

	curMark := cfg.sentence[index].Kaeriten
	nextMark := nextMarks[curMark]

	// For when the mark is 上 and there is no 中
	if _, ok := cfg.marks[nextMark]; !ok {
	    nextMark = nextMarks[nextMark]
	} 

	nextIndex, ok := cfg.marks[nextMark]
	if ok {
		cfg.recursivePull(nextIndex)
	} 
	delete(cfg.marks, curMark)
} 

var nextMarks = map[string]string{
	// 一二
	"一":"二",
	"二":"三",
	"三": "四", 
	"四": "五",
	"五": "六",
	"六": "七",
	"七": "八",
	"八": "九",
	"九": "十",

	// 上下
	"上": "中",
	"中": "下",

	// 十干
	"甲":"乙",
	"乙":"丙",
	"丙":"丁",
	"丁": "戊",
	"戊": "己",
	// (庚辛壬癸)

	// 天地人
	"天": "地",
	"地": "人",

	// 四徳
	"元": "亨",
	"亨": "利",
	"利": "貞",

	// 乾坤
	"乾": "坤",
} 
