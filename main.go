package main

import (
	"emailapi/app"
	"emailapi/app/database"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

// entry point of the app is constructed
func main() {
	println("app running...")
	app := app.New()
	// initiate connection to database
	app.DB = &database.DB{}
	err := app.DB.Open()
	check(err)
	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)
	err = http.ListenAndServe(":8000", nil)
	check(err)
}

func check(e error){
	if e != nil{
		log.Println(e)
		os.Exit(1)
	}
}