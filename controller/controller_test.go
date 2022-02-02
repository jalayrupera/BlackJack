package controller

import (
	"bytes"
	"net/http"
	"testing"
)
 
func TestCreateDeck_Correct(t *testing.T) {

	var jsonStr = []byte(`{"shuffled":true,"cards":"AH,AS,AD"}`)

	res, err := http.Post("http://localhost:8000/deck", "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code %v returened", status)
	}

}

func TestCreateDeck_Wrong(t *testing.T) {

	var jsonStr = []byte(`{"shuffled":"true","cards":"A,AS,AD"}`)

	res, err := http.Post("http://localhost:8000/deck", "application/json", bytes.NewBuffer(jsonStr))

	if err != nil {
		panic(err)
	}

	if status := res.StatusCode; status == http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

}

func TestGetDeckByUUID_Correct(t *testing.T) {
	res, err := http.Get("http://localhost:8000/deck/cf0d8007-b208-4603-a6f4-240aa2606e63")

	if err != nil {
		panic(err)
	}

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGetDeckByUUID_Wrong(t *testing.T) {
	res, err := http.Get("http://localhost:8000/deck/cf0d8007-b208-4603-a6f4-240aa2606e63")

	if err != nil {
		panic(err)
	}

	if status := res.StatusCode; status == http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
}
