package Controller

import (
    "io"
    "net/http"
    "log"
)


func HelloController(res http.ResponseWriter, req *http.Request, isASynchronized bool, callHandler chan bool) {
    if !isASynchronized {
        go helloServer_sync(res, req, callHandler)
    } else {
        go helloServer_async(res, req)
        callHandler <- true
    }
}

/**
 * Handle URL /hello
 */
func helloServer_async(res http.ResponseWriter, req *http.Request) {

    helloController_Process(res, req)

}

func helloServer_sync(res http.ResponseWriter, req *http.Request, callHandler chan bool) {
    helloController_Process(res, req)
    callHandler <- true
}

func helloController_Process(res http.ResponseWriter, req *http.Request) {
    // TODO
}