package dispatcher

import (
	"net/http"
	Controller "github.com/uguisu/controller"
)

/**
 * Declare DefaultDispatcher
 */
type DefaultDispatcher struct {
    
    // TODO
	AllowRedirect   bool
	
	// TODO
	AllowASync      bool
}

func (dp DefaultDispatcher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	
// 	Controller.HelloServer(res, req)

    callHandler := make(chan bool)
    
    switch req.URL.Path {
        case "/hello":
            go Controller.HelloController(res, req, dp.AllowASync, callHandler)
    }
	
	<-callHandler
}

func Create() *DefaultDispatcher {
	return &DefaultDispatcher {
		AllowRedirect: true,
		AllowASync:    true,
	}
}