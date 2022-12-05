package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

type JsonResponse struct {
	Value1     string     `json:"key1"`
	Value2     string     `json:"key2"`
	JsonNested JsonNested `json:"JsonNested"`
	JsonNew    JsonNew    `json:"JsonNew"`
}

type JsonNew struct {
	Nested1     string      `json:"FirstName"`
	Nested2     string      `json:"LastName"`
	JsonAddress JsonAddress `json:"JsonAddress"`
}

type JsonAddress struct {
	NewNest1 string `json:"Street"`
	NewNest2 string `json:"City"`
}

type JsonNested struct {
	NestedValue1 string `json:"nestedKey1"`
	NestedValue2 string `json:"nestedKey2"`
}

func main() {

	welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}

	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	nested := JsonNested{
		NestedValue1: "Raj",
		NestedValue2: "Kotak",
	}

	newdouble := JsonAddress{
		NewNest1: "lb nagar, kothapet",
		NewNest2: "Mumbai",
	}

	newnest := JsonNew{
		Nested1:     "rahul",
		Nested2:     "dravid",
		JsonAddress: newdouble,
	}

	jsonResp := JsonResponse{
		Value1:     "some Data",
		Value2:     "other Data",
		JsonNested: nested,
		JsonNew:    newnest,
	}

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/jsonResponse", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(jsonResp)
	})

	http.HandleFunc("/JsonNew", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(newnest)
	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
