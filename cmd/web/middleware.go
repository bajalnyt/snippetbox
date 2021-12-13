package main

import (
	"fmt"
	"net/http"
)

// Middleware flow:
// secureHeaders --> servemux --> application handler
func secureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-XSS-Protection", "1, mode=block")
		w.Header().Set("X-Frame-Options", "deny")

		//Any Code here gets executed during the way in
		next.ServeHTTP(w, r)

		//Any Code here gets executed during the way back up
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// recover() is a built in function to check if there is a panic
		// this deferred function will always run when go unwinds the call stack
		defer func() {
			if err := recover(); err != nil {
				//trigger to make Go’s HTTP server automatically close the current connection after a response has been sent.
				w.Header().Set("Conection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
