package middleware

import (
	"fmt"
	"net/http"
)

func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer internalRecover(w)
		next.ServeHTTP(w, r)
	})
}

func internalRecover(w http.ResponseWriter) {
	err := recover()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
