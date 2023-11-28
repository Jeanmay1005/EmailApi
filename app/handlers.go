package app

import (
	"emailapi/app/models"
	"fmt"
	"log"
	"net/http"
)

// an index handler that returns a function
func (a *App) indexHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to Email API")
	}
}

func (a *App) CreateEmailHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		req := models.EmailRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("cannot parse post body. err = %v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}
		// Create the Email
		e := &models.Email{
			ID: 0,
			Title: req.Title,
			Content: req.Content,
			Author: req.Author,
		}

		// Save in Database
		err = a.DB.CreateEmail(e)
		if err != nil{
			log.Printf("Cannot save post in DB. err = %v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}
		// Create http response with email object
		response := mapEmailToJSON(e)
		sendResponse(w, r, response, http.StatusOK)
	}
}

func (a *App) GetEmailHandler() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		emails, err := a.DB.GetEmail()
		if err != nil{
			log.Printf("Cannot get posts, err = %v\n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
		}

		var response = make([]models.JSONEmail, len(emails))
		for idx, email := range emails{
			response[idx] = mapEmailToJSON(email)
		}

		sendResponse(w, r, response, http.StatusOK)
	}
}


