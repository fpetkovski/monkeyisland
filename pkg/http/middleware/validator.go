package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/thedevsaddam/govalidator"
	"io/ioutil"
	"net/http"
	"net/url"
)

func MakeValidator(rules govalidator.MapData) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			payload, _ := ioutil.ReadAll(r.Body)
			firstReader := ioutil.NopCloser(bytes.NewBuffer(payload))
			secondReader := ioutil.NopCloser(bytes.NewBuffer(payload))

			r.Body = firstReader
			data := make(map[string]interface{}, 0)
			opts := govalidator.Options{
				Request: r,
				Rules:   rules,
				Data:    &data,
			}
			v := govalidator.New(opts)
			e := v.ValidateJSON()
			if len(e) != 0 {
				w.WriteHeader(http.StatusUnprocessableEntity)
				sendErrors(e, w)
				return
			}

			r.Body = secondReader
			next.ServeHTTP(w, r)
		})
	}
}

func sendErrors(errors url.Values, w http.ResponseWriter) {
	response := map[string]interface{}{"error": errors}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		panic(err)
	}
}
