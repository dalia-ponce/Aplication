package main

import (
    "fmt"
    "html/template"
    "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
    template, err := template.ParseFiles("Templates/index.html")
    if err != nil {
        fmt.Fprintf(w, "Pagina no encontrada")
    } else {
        template.Execute(w, nil)
    }
}

func main() {

    http.HandleFunc("/", index)

    fmt.Println("Hola Mundo!")
    http.ListenAndServe(":3000", nil)
}