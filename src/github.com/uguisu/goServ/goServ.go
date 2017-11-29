package main

import (
    "net/http"
    "log"
    Config "github.com/uguisu/config"
)

/**
 * Main
 */
func main() {

    // Load config
    configFinished := Config.LoadConfig()
    if <- configFinished {
        log.Println("Load config finished.")
    } else {
        log.Fatal("Unknow excetion whien loading config")
    }

    log.Println("Server started.")

    // Start server
    err := http.ListenAndServe(Config.GetPort(), nil)
    if err != nil {
        log.Fatal("Unknow excetion: ", err)
    }
}
