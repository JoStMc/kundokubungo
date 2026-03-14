package engine

import (
	"strings"

	"github.com/JoStMc/kundokubungo/internal/models"
)

func ToKakikudashi(sentence *models.Sentence) (string, error) {
	order, err := getCharOrder(sentence)
	if err != nil {
	    return "", err
	} 
	var output strings.Builder

	saidokuParsed := make(map[int]struct{})

	for _, charIdx := range order {
		char := sentence.Characters[charIdx]
		kanji, okurigana := char.Kanji, char.Okurigana
		if _, ok := saidokuParsed[charIdx]; ok {
			okurigana = char.SecondOkurigana
		} 
		output.WriteString(kanji)
		output.WriteString(okurigana)
	} 
	return output.String(), nil
} 
