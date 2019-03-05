package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type timeObj struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", renderIndex)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func timeNow() timeObj {
	// find the time right...now!!
	now := time.Now()

	// create time object from now
	time := timeObj{
		// TODO: why are these set up in this way?
		Date: now.Format("02-01-2006"),
		Time: now.Format("15:04:05"),
	}

	return time
}

func renderIndex(rsw http.ResponseWriter, req *http.Request) {
	////////////////////////
	// req is request,    //
	// rsw is writers     //
	//  (& response?)     //
	// ------TODO:------- //
	// need more research //
	// on these meanings. //
	////////////////////////

	// parse the html file index.html
	tmp, err := template.ParseFiles("./html/index.html")
	if err != nil { // log any errors
		log.Print("template parsing error: ", err)
	}

	// execute the template and pass it the indexVars struct to fill in the gaps
	err = tmp.Execute(rsw, timeNow())
	if err != nil { // log any errors
		log.Print("template executing error: ", err)
	}
}
