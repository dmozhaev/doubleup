package main

import (
    "fmt"
    "net/http"
    "double_up/controller"
    "double_up/utils"
)

func main() {
    // Establish a connection to the database
    db, err := utils.Connect()
    if err != nil {
        fmt.Println("Error connecting to database:", err)
        return
    }
    defer db.Close()

    // API routes
    http.HandleFunc("/play/start", func(w http.ResponseWriter, r *http.Request) {
        controller.PlayStartHandler(db, w, r)
    })
    http.HandleFunc("/play/continue", func(w http.ResponseWriter, r *http.Request) {
        controller.PlayContinueHandler(db, w, r)
    })

    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}
