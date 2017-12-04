package Controller

import (
    "io"
    "net/http"
    "log"
)

/**
 * HelloController
 */
func HelloController(res http.ResponseWriter, req *http.Request, isASynchronized bool, callHandler chan bool) {
	
	// make sure the call is sync or async
    if !isASynchronized {
        go helloServer_sync(res, req, callHandler)
    } else {
        go helloServer_async(res, req)
        // update caller channel
        callHandler <- true
    }
}

/**
 * Handle URL "/hello", async
 */
func helloServer_async(res http.ResponseWriter, req *http.Request) {
    helloController_Process(res, req)
}

/**
 * Handle URL "/hello", sync
 */
func helloServer_sync(res http.ResponseWriter, req *http.Request, callHandler chan bool) {
    helloController_Process(res, req)
    // update caller channel
    callHandler <- true
}

/**
 * Business logic
 */
func helloController_Process(res http.ResponseWriter, req *http.Request) {
    log.Println("An incomming request")
    io.WriteString(res, "hello, world!\n")
}