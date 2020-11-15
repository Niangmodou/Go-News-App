package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

//Search Handler to search for news
func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.Url.String())

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.WriteHeader([]byte("Internal Server Error"))

		return
	}

	params := u.Query()

	searchKey = params.Get("q")
	page = params.Get("page")

	if page == "" {
		page = "1"
	}

	fmt.Println("Search Key: ", searchKey)
	fmt.Println("PageNum: ", page)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}
	//Creation of a HTTP request multiplexer
	mux := http.NewServeMux()

	//Handler functions for the paths
	mux.HandleFunc("/", indexHandler)
	mux.HandlerFunc("/search", searchHandler)

	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//Serving the API at port 3000
	http.ListenAndServe(":"+port, mux)
}
