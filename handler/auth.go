package handler

import (
	"net/http"

	dblayer "github.com/moxiaomomo/filestore-cloud/db"
)

func AccessAuth(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		phone := r.Form.Get("phone")
		token := r.Form.Get("token")

		if IsTokenExpired(token) || !dblayer.TokenValid(phone, token) {
			w.Write([]byte("{\"code\":-1,\"msg\":\"token invalid\"}"))
			return
		}
		h(w, r)
	})
}
