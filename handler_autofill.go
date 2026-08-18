package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"

	"github.com/JoStMc/kundokubungo/internal/engine"
	"github.com/JoStMc/kundokubungo/internal/models"
)

type autofillRequest struct {
	Text string `json:"order"`
}

type autofillResponse struct {
	Sentence 	models.Sentence `json:"sentence"`
	Kakikudashi string 			`json:"kakikudashi"`
	Id 		 	int			 	`json:"id"`
}

func handlerAutofill(w http.ResponseWriter, r *http.Request) {
	var rawInput autofillRequest
	err := json.NewDecoder(r.Body).Decode(&rawInput)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Unable to decode request", err))
		return
	}

	numbers := extractNumbers(rawInput.Text)

	kaeriten, err := engine.GetKaeriten(numbers)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprint("Could not set order", err))
		return
	}

	for i := range sentenceStore.Characters {
	    sentenceStore.Characters[i].Kaeriten = kaeriten[i]
	} 

	kakikudashi := engine.ToKakikudashi(&sentenceStore)

	respondWithJSON(w, http.StatusOK, autofillResponse{
	    Sentence: sentenceStore,
		Kakikudashi: kakikudashi,
		Id: 1,
	})
}

func extractNumbers(s string) []int {
	re := regexp.MustCompile(`[0-9]+`)
	matches := re.FindAllString(s, -1)
	numbers := make([]int, len(matches))
	for i, match := range matches {
		numbers[i], _ = strconv.Atoi(match)
	}
	return numbers
}
