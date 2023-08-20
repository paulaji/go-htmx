// define a package name
// package refers to lines of code / group of code
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// define a struct
type Movie struct {
	Title string
	Actress string
}

func main(){
	fmt.Println("Namaskaaram!")

	// defining the h1 function, handler 1 (h1) for when we are at the home route
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		movies := map[string][]Movie{
			"Movies": {
				{Title: "Kantaara", Actress: "Saptami Gowda"},
				{Title: "Knock Knock", Actress: "Ana D Armas"},
			},
		}
		tmpl.Execute(w, movies) // the second arguement is nil, which means we are not passing any value into the html page, same as django
	}

	// now define which handler to use in a particular route
	http.HandleFunc("/", h1) // in this case, use the h1 function which we will define
	
	// set a port for server to listen to
	log.Fatal(http.ListenAndServe(":8000", nil))
}