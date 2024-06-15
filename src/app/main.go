package main

import (
  "fmt"
  "net/http"
)

func login(username string, password string) string {
  return "login ! username: " + username + "; " + "password: " + password + "; \n"
}

func createuser(token string, username string, host string) string {
  return "createuser ! token: " + token + "; " + "username: " + username + "; " + "host: " + host + "; \n"
}

func http_login(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
  username := query["username"][0]
	password := query["password"][0]

  fmt.Fprintf(w, login(username, password))
}

func http_createuser(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
  token := query["token"][0]
  username := query["username"][0]
  host := query["host"][0]

  fmt.Fprintf(w, createuser(token, username, host))
}

func http_describeuser(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "describeuser\n")
}

func http_grantuser(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "grantuser\n")
}

func http_revokeuser(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "revokeuser\n")
}

func http_deleteuser(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "deleteuser\n")
}

func http_deleteallusers(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "deleteallusers\n")
}

func http_listallusers(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "listallusers\n")
}

func http_createdir(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "createdir\n")
}

func http_listdir(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "listdir\n")
}

func http_publicdir(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "publicdir\n")
}

func http_privatedir(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "privatedir\n")
}

func http_deletedir(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "deletedir\n")
}

func http_createfile(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "createfile\n")
}

func http_listfile(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "listfile\n")
}

func http_readfile(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "readfile\n")
}

func http_writefile(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "writefile\n")
}

func http_deletefile(w http.ResponseWriter, req *http.Request) {
  fmt.Println(w, "deletefile\n")
}

func main() {
  http.HandleFunc("/login", http_login)
  http.HandleFunc("/createuser", http_createuser)
  http.HandleFunc("/describeuser", http_describeuser)
  http.HandleFunc("/grantuser", http_grantuser)
  http.HandleFunc("/revokeuser", http_revokeuser)
  http.HandleFunc("/deleteuser", http_deleteuser)
  http.HandleFunc("/deleteallusers", http_deleteallusers)
  http.HandleFunc("/listallusers", http_listallusers)
  http.HandleFunc("/createdir", http_createdir)
  http.HandleFunc("/listdir", http_listdir)
  http.HandleFunc("/publicdir", http_publicdir)
  http.HandleFunc("/privatedir", http_privatedir)
  http.HandleFunc("/deletedir", http_deletedir)
  http.HandleFunc("/createfile", http_createfile)
  http.HandleFunc("/listfile", http_listfile)
  http.HandleFunc("/readfile", http_readfile)
  http.HandleFunc("/writefile", http_writefile)
  http.HandleFunc("/deletefile", http_deletefile)

  http.ListenAndServe(":80", nil)
}
