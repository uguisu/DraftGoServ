package main

import (
    "net/http"
    "log"
    "Controller"
    "Config"
)

/**
 * Main
 */
func main() {


    go Config.LoadConfig()

    http.HandleFunc("/hello", Controller.HelloServer)
    err := http.ListenAndServe(Config.GetPort(), nil)
    if err != nil {
            log.Fatal("ListenAndServe: ", err)
    }
}
