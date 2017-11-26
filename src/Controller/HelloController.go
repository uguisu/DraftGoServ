package Controller

import (
    "io"
    "net/http"
    "log"
)

/**
 * Handle URL /hello
 */
func HelloServer(w http.ResponseWriter, req *http.Request) {

    log.Println("An incomming request");

    io.WriteString(w, "hello, world!\n")

}
