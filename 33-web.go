package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "apa kabar!")
}

func templateExample(w http.ResponseWriter, r *http.Request) {
	var data = map[string]string{
		"Name":    "john wick",
		"Message": "have a nice day",
	}

	var t, err = template.ParseFiles("template.html")
	if err != nil {
		fmt.Println(err.Error())
	}

	t.Execute(w, data)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "halo!")
	})

	http.HandleFunc("/index", index)
	// penggunaan tempate web
	http.HandleFunc("/template-example", templateExample)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
