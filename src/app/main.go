package main

import (
  "fmt"
  "net/http"
)

func login(username string, password string) string {
  return "username: " + username + "; " + "password: " + password + "; \n"
}

func http_login(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
  username := query["username"][0]
	password := query["password"][0]

  fmt.Fprintf(w, login(username, password))
}

func http_createuser(w http.ResponseWriter, req *http.Request) {
  fmt.Fprintf(w, "createuser\n")
}

func main() {
  http.HandleFunc("/login", http_login)
  http.HandleFunc("/createuser", http_createuser)

  http.ListenAndServe(":80", nil)
}
