package main

import (
  "fmt"
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Halaman Awal")
}

func info(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Ini kayaknya Halaman Info")
}

func main(){
  http.HandleFunc("/", home)
  http.HandleFunc("/kampret", info)
func contact(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Halaman Kontak")
}


func main(){
  http.HandleFunc("/", home)
  http.HandleFunc("/contact", contact)
  http.ListenAndServe(":8080", nil)
}
