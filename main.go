package main

import (
  "fmt"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Halaman Awal")
}

func contact(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Halaman Kontak")
}


func main(){
  http.HandleFunc("/", home)
  http.HandleFunc("/contact", contact)
  http.ListenAndServe(":8080", nil)
}
