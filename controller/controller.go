package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/jalayrupera/toggltest/model"
)

var all_decks = []model.Deck{}

func getDeckFromUUID(uuid string) *model.Deck {
	for _, deck := range all_decks {
		if deck.Uuid == uuid {
			return &deck
		}
	}
	return nil
}

func handleErrorMessage(message string, statusCode int, w *http.ResponseWriter) {
	(*w).WriteHeader(statusCode)
	(*w).Write([]byte(`{"message": "Error: ` + message + `"}`))
}

func CreateDeck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create Deck Request ")
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "POST" {
		handleErrorMessage("method not allowed", http.StatusInternalServerError, &w)
		return
	}
	url := r.URL

	var deck model.Deck

	json.NewDecoder(r.Body).Decode(&deck)

	cardsStr := url.Query().Get("cards")
	newDeckOfCard := model.New_deck(cardsStr)

	if newDeckOfCard == nil {
		handleErrorMessage("could not create new deck of cards. Please check the request", http.StatusInternalServerError, &w)
		return
	}

	if deck.Shuffled {
		model.Shuffle(newDeckOfCard)
	}

	uuid := uuid.NewString()

	newDeck := model.Deck{Cards: &newDeckOfCard, Shuffled: deck.Shuffled, Uuid: uuid}

	all_decks = append(all_decks, newDeck)

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"deck_id": "` + uuid + `", "shuffled": ` + strconv.FormatBool(newDeck.Shuffled) + `, "remaining": ` + strconv.Itoa(len(*newDeck.Cards)) + `}`))
}

func GetDeck(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get Deck Request ")

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET"{
		handleErrorMessage("method not allowed", http.StatusInternalServerError, &w)
		return
	}

	url := r.URL

	uuid := strings.Split(url.String(), "/")[2]

	deck := getDeckFromUUID(uuid)

	if deck == nil {
		handleErrorMessage("Deck not found!", http.StatusBadRequest, &w)
	}

	deckJSON, err := json.Marshal(deck.Cards)

	if err != nil{
		handleErrorMessage(err.Error(), http.StatusInternalServerError, &w)
	}

	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"deck_id": "` + uuid + `", "shuffled": ` + strconv.FormatBool(deck.Shuffled) + `,"remaining": ` + strconv.Itoa(len(*deck.Cards)) + `, "cards": ` + string(deckJSON) + `}`))
}

func DrawCard(w http.ResponseWriter, r *http.Request){
	fmt.Println("Draw Deck Request ")

	w.Header().Set("Content-Type", "application/json")

	if r.Method != "PUT" {
		handleErrorMessage("method not allowed", http.StatusInternalServerError, &w)
		return
	}

	url := r.URL
	uuid := strings.Split(url.String(), "/")[2]

	var numOfCardsToDrawn model.DrawCardsReqBody

	json.NewDecoder(r.Body).Decode(&numOfCardsToDrawn)

	deck := getDeckFromUUID(uuid)

	if deck == nil{
		handleErrorMessage("deck not found!", http.StatusNotFound, &w)
		return
	}

	if numOfCardsToDrawn.NumberOfCards > len(*deck.Cards){
		handleErrorMessage("cannot draw that many cards! deck only have "+strconv.Itoa(len(*deck.Cards)), http.StatusInternalServerError, &w)
		return
	}

	drawnCards := model.Draw_cards(deck.Cards, numOfCardsToDrawn.NumberOfCards)

	deckJSON, err := json.Marshal(drawnCards)

	if err != nil{
		handleErrorMessage(err.Error(), http.StatusInternalServerError, &w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"cards": ` + string(deckJSON) + `}`))
}