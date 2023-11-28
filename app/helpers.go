package app

import (
	"emailapi/app/models"
	"encoding/json"
	"log"
	"net/http"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

// Send http status helper function
func sendResponse(w http.ResponseWriter, r *http.Request, data interface{}, status int){
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json. err = %v\n", err)
	}
}

// function that maps email struct to JsonEmail
func mapEmailToJSON(e *models.Email) models.JSONEmail{
	return models.JSONEmail{
		ID: e.ID,
		Title: e.Title,
		Content: e.Content,
		Author: e.Author,
	}
}