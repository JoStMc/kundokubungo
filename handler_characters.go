package main

import (
	"context"
	"fmt"
	"net/http"
)

type lookupResponse struct{
	Kanji   string `json:"kanji"`
	Onyomi  string `json:"onyomi"`
	Kunyomi string `json:"kunyomi"`
	Imi     string `json:"imi"`
	Itaiji  string `json:"itaiji"`
	Bushu   string `json:"bushu"`
} 

func (cfg *config) handlerCharacterLookup(w http.ResponseWriter, r *http.Request) {
	character := r.PathValue("kanji")

	charInfo, err := cfg.dbQueries.GetKanji(context.TODO(), character)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error looking up character: %v", err))
		return
	}

	response := lookupResponse{
		Kanji: charInfo.Kanji,
		Onyomi: charInfo.Onyomi.String,
		Kunyomi: charInfo.Kunyomi.String,
		Imi: charInfo.Imi.String,
		Itaiji: charInfo.Itaiji.String,
		Bushu: charInfo.Bushu.String,
	} 
	respondWithJSON(w, http.StatusOK, response)
}
