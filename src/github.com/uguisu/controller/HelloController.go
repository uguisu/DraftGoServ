package Controller

import (
    "io"
    "net/http"
    "log"
)

/**
 * Handle URL /hello
 */
func HelloServer(res http.ResponseWriter, req *http.Request) {

    log.Println("An incomming request");

    io.WriteString(res, "hello, world!\n")

}
