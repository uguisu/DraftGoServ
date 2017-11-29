package dispatcher

import (
	"net/http"
	Controller "github.com/uguisu/controller"
)

/**
 * Declare DefaultDispatcher
 */
type DefaultDispatcher struct {
	AllowRedirect   bool
}

func (dp DefaultDispatcher) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	
	Controller.HelloServer(res, req)
	
}

func Create() *DefaultDispatcher {
	return &DefaultDispatcher {
		AllowRedirect: true,
	}
}