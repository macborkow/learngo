package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your request: %s\n", r.URL.Path)
  })

  r.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello there!")
  })

  r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    title := vars["title"]
    page := vars["page"]

    fmt.Fprintf(w, "You requested book %s on page %s.\n", title, page)
  })

  fs := http.FileServer(http.Dir("static/"))
  r.Handle("/static/", http.StripPrefix("/static/", fs))

  http.ListenAndServe(":1234", r)
}
