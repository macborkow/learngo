package main

import (
  "fmt"
  "net/http"
  "log"
  "time"

  "github.com/gorilla/mux"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)


func main() {
  db, err := sql.Open("mysql", "root:root@(172.17.0.2:3306)/mysql?parseTime=true")
  if err != nil {
    log.Fatal(err)
  }
  if err := db.Ping(); err != nil {
    log.Fatal(err)
  }

/*
  { // create a table
    query := `
      CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
      );`

    if _, err := db.Exec(query); err != nil {
      log.Fatal(err)
    }
  }
*/

  { // insert a user
    username := "john"
    password := "secret"
    createdAt := time.Now()

    result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
    if err != nil {
      log.Fatal(err)
    }

    id, err := result.LastInsertId()
    fmt.Println(id)
  }

  { // query a single user
    var (
      id        int
      username  string
      password  string
      createdAt time.Time
    )

    query := "SELECT id, username, password, created_at FROM users WHERE id = ?"
    if err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt); err != nil {
      log.Fatal(err)
    }

    fmt.Println(id, username, password, createdAt)
  }

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
