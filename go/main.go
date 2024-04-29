package main

import (
    "fmt"
    "net/http"
    "double_up/controller"
)

func main() {
    http.HandleFunc("/play/start", controller.PlayStartHandler)

    fmt.Println("Server listening on port 8080...")
    http.ListenAndServe(":8080", nil)
}

