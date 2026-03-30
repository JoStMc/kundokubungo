package main

import (
	"context"
	"fmt"
	"net/http"
)

func (cfg *config) handlerCharacterLookup(w http.ResponseWriter, r *http.Request) {
	character := r.PathValue("kanji")

	charInfo, err := cfg.dbQueries.GetKanji(context.TODO(), character)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error looking up character: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, charInfo)
}
