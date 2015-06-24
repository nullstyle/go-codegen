package foo

import (
	"net/http"
)

// ServeHTTPC is a method for web.Handler
func (action *MyCustomAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	_ = "arg1"
	_ = "second_atg"

	ap := &action.Action
	ap.Prepare(w, r)
	ap.Execute(&action)
}
