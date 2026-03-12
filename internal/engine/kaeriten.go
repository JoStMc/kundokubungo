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
	// 一二
	"二": "一",
	"三": "二",
	"四": "三", 
	"五": "四",
	"六": "七",
	"七": "八",
	"八": "九",
	"九": "十",

	// 上下
	"中": "上",
	"下": "中",

	// 十干
	"乙": "甲",
	"丙": "乙",
	"丁": "丙",
	"戊": "丁",
	"己": "戊",
	// (庚辛壬癸)

	// 天地人
	"地": "天",
	"人": "地",

	// 四徳
	"亨": "元",
	"利": "亨",
	"貞": "利",

	// 乾坤
	"坤": "乾",
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
