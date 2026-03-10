package engine

import (
	"errors"
	"fmt"

	"github.com/JoStMc/kundokubungo/internal/models"
)

// Here we define what to do for each kaeriten
var kaeritenTypes = map[string]func(*config, *models.Character, int) {
	models.MarkRe: (*config).reten,
    models.MarkIchi: (*config).ichiten,
    models.MarkNi: (*config).niten,
    models.MarkSan: (*config).santen, 
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
		    return []int{}, errors.New("unknown kaeriten")
		} 
		kaeriFunc(&cfg, &char, i)
	} 

	return cfg.order, nil
} 

func (cfg *config) allChars(char *models.Character, index int) {
	cfg.order[index] = cfg.currentChar
	fmt.Printf("Updated char '%v', position %d, to value %d - currentChar now %d\n", char.Kaeriten, index, cfg.currentChar, cfg.currentChar+1)
	cfg.currentChar++
	if index != 0 && cfg.sentence[index-1].Kaeriten == models.MarkRe {
		cfg.allChars(&cfg.sentence[index-1], index-1)
	} 
} 

func (cfg *config) reten(char *models.Character, index int) {
	// This function should do nothing, because allChars should
	// catch reten by checking if the previous char is reten.
	// We don't want to add 1 to the current char
} 

func (cfg *config) ichiten(char *models.Character, index int) {
	cfg.marks[models.MarkIchi] = index
	cfg.allChars(char, index)
	if nimark, ok := cfg.marks[models.MarkNi]; ok {
		cfg.niten(&cfg.sentence[nimark], nimark)
		delete(cfg.marks, models.MarkIchi)
	}
} 

func (cfg *config) niten(char *models.Character, index int) {
	cfg.marks[models.MarkNi] = index
	if _, ok := cfg.marks[models.MarkIchi]; ok {
		cfg.allChars(char, index)
		if sanmark, ok := cfg.marks[models.MarkSan]; ok {
			cfg.santen(&cfg.sentence[sanmark], sanmark)
			delete(cfg.marks, models.MarkNi)
		}
	} 
} 

func (cfg *config) santen(char *models.Character, index int) {
	cfg.marks[models.MarkSan] = index
	if _, ok := cfg.marks[models.MarkNi]; ok {
		cfg.allChars(char, index)
		delete(cfg.marks, models.MarkSan)
	} 
} 
