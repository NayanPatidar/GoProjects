package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	http.Handle("/Hello", http.HandlerFunc(HelloHandler))

	// Handle requests for static files using the file server
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/form", FormHandler)

	port := ":5000"
	fmt.Println("Server is running on port" + port)

	http.ListenAndServe(port, nil)

}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hellooo!!")
}

func FormHandler(w http.ResponseWriter, r *http.Request) {

	temp, _ := template.ParseFiles("static/form.html")
	temp.Execute(w, nil)
}
