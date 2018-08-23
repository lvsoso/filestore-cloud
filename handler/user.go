package handler

import (
	"net/http"

	dblayer "github.com/moxiaomomo/distributed-fileserver/db"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	user := r.Form.Get("user")
	pwd := r.Form.Get("pwd")

	res := dblayer.UserRegister(user, pwd)
	if res {
		w.Write([]byte("{\"code\":0,\"msg\":\"user register succeeded.\"}"))
	} else {
		w.Write([]byte("{\"code\":-1,\"msg\":\"user register failed.\"}"))
	}
}
