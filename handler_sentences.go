package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/JoStMc/kundokubungo/internal/models"
)


type createRequest struct {
	Text string `json:"text"`
} 

type createResponse struct {
	Sentence models.Sentence `json:"sentence"`
} 

func handlerCreate(w http.ResponseWriter, r *http.Request) {
	var rawInput createRequest
	err := json.NewDecoder(r.Body).Decode(&rawInput)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Unable to decode request", err))
		return
	}

	sentence := models.Sentence{
	    Characters: make([]models.Character, utf8.RuneCountInString(rawInput.Text)),
	} 

	runeCount := 0
	for _, char := range rawInput.Text {
		sentence.Characters[runeCount].Kanji = string(char)
		runeCount++
	} 

	respondWithJSON(w, http.StatusOK, createResponse{
	    Sentence: sentence,
	})
}

func handlerUpdate(w http.ResponseWriter, r *http.Request) {
} 
