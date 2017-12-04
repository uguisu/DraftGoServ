package dispatcher

import (
	"net/http"
	Controller "github.com/uguisu/controller"
)

/**
 * Declare DefaultDispatcher
 */
type DefaultDispatcher struct {
    
    // Re-direct flag. Allow URL re-direct to a new target
	AllowRedirect   bool
	
	// Async flag. This will call target function in async way.
	// The result may could not return to http client
	AllowASync      bool
}

func (dp DefaultDispatcher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	
	// declare a caller channel
    callHandler := make(chan bool)
    
    // dispatch target URL
    // TODO: use reflect to call method automatically
    switch req.URL.Path {
        case "/hello":
            go Controller.HelloController(res, req, dp.AllowASync, callHandler)
    }
	
	// wait for function finished
	<-callHandler
}

/**
 * Create new DefaultDispatcher object
 */
func Create() *DefaultDispatcher {
	return &DefaultDispatcher {
		AllowRedirect: true,
		AllowASync:    false,
	}
}