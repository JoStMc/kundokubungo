package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"unicode/utf8"

	"github.com/JoStMc/kundokubungo/internal/engine"
	"github.com/JoStMc/kundokubungo/internal/models"
)


type createRequest struct {
	Text string `json:"text"`
} 

type createResponse struct {
	Sentence models.Sentence `json:"sentence"`
	Id 		 int			 `json:"id"`
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

	sentenceStore = sentence

	respondWithJSON(w, http.StatusOK, createResponse{
	    Sentence: sentence,
		Id: 1,
	})
}


type updateRequest struct {
	Text       string `json:"text"`
	Index 	   int	  `json:"index"`
	SentenceId int    `json:"sentence_id"`
	UpdateType string `json:"type"`
} 

type updateResponse struct {
	Text string `json:"text"`
} 

func handlerUpdate(w http.ResponseWriter, r *http.Request) {
	var update updateRequest
	err := json.NewDecoder(r.Body).Decode(&update)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Unable to decode request", err))
		return
	}

	character := &sentenceStore.Characters[update.Index]
	switch update.UpdateType {
	case "kaeri":
		character.Kaeriten = update.Text
	case "okuri":
		character.Okurigana = update.Text
	case "okuri2":
		character.SecondOkurigana = update.Text
	case "saidoku":
		character.IsSaidokumoji = !character.IsSaidokumoji
	case "juku":
		if len(sentenceStore.Characters) == update.Index + 1 {
		    respondWithError(w, http.StatusBadRequest, "Cannot set last character as jukugo")
			return
		} 
		character.IsJukugoHead = !character.IsJukugoHead
		nextChar := &sentenceStore.Characters[update.Index + 1]
		nextChar.IsJukugoTail = !nextChar.IsJukugoTail
	} 

	kakikudashi, err := engine.ToKakikudashi(&sentenceStore)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprint("Unable to convert sentence:", err))
		return
	}

	respondWithJSON(w, http.StatusOK, updateResponse{
		Text: kakikudashi,
	})
} 
