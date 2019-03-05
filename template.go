package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type pageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	////////////////////////
	// r is request.      //
	// w is writers       //
	//  (& response?)     //
	// ------------------ //
	// need more research //
	// on these meanings. //
	////////////////////////

	now := time.Now()              // find the time right now
	HomePageVars := pageVariables{ //store the date and time in a struct
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	t, err := template.ParseFiles("./html/index.html") //parse the html file homepage.html
	if err != nil {                                    // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
