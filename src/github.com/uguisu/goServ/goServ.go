package main

import (
    "net/http"
    "log"
    Config             "github.com/uguisu/config"
    DefaultDispatcher  "github.com/uguisu/dispatcher"
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

	dp := DefaultDispatcher.Create()

    // Start server
    err := http.ListenAndServe(Config.GetPort(), dp)
    if err != nil {
        log.Fatal("Unknow excetion: ", err)
    }
}
