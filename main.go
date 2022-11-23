package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/addproject", addproject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project/{id}", project).Methods("GET")
	route.HandleFunc("/submitproject", submitproject)

	port := "5000"

	fmt.Print("Server sedang berjalan di port " + port + "\n")
	http.ListenAndServe("localhost:"+port, route)
}

// Home
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

// Add Project
func addproject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/add-project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

// Contact
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}

// Project
func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf=8")
	tmpt, err := template.ParseFiles("views/project.html")

	if err != nil {
		w.Write([]byte("Message : " + err.Error()))
		return
	}

	// Id
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// Object
	data := map[string]interface{}{
		"title":   "Judul",
		"content": "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		"id":      id,
	}

	tmpt.Execute(w, data)
}

// Project Form Submit
func submitproject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(
		"Title: "+r.PostForm.Get("title"),
		"\nContent: "+r.PostForm.Get("content"),
		"\nDate Start: "+r.PostForm.Get("datestart"),
		"\nDate Start: "+r.PostForm.Get("datestart"),
		"\nTechnologies: ",
		"\n Node Js: "+r.PostForm.Get("nodejs"),
		"\n React Js: "+r.PostForm.Get("reactjs"),
		"\n Next Js: "+r.PostForm.Get("nextjs"),
		"\n TypeScript: "+r.PostForm.Get("typescript"),
	)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
