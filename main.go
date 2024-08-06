package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", getRoot)

	err := http.ListenAndServe(":8081", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("received request: %s\n", r.Host)
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "error parsing template", http.StatusInternalServerError)
	}

	data := struct {
		Name string
	}{
		Name: "Simon",
	}

	tmpl.Execute(w, data)

}
