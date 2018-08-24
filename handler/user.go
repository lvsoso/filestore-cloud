package handler

import (
	"fmt"
	"net/http"
	"time"

	dblayer "github.com/moxiaomomo/distributed-fileserver/db"
	"github.com/moxiaomomo/distributed-fileserver/util"
)

const (
	pwd_salt   = "_test"
	token_salt = "_test2"
)

func genLoginToken(user string) string {
	ts := fmt.Sprintf("%x", time.Now().Unix())
	rstr := util.MD5([]byte(user + ts + token_salt))
	return rstr + ts[:8]
}

// handle register
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phone := r.Form.Get("phone")
	pwd := r.Form.Get("pwd")

	if len(phone) < 3 || len(phone) > 32 || len(pwd) != 32 {
		w.Write([]byte("{\"code\":-1,\"msg\":\"params invalid.\"}"))
		return
	}

	enc_pwd := util.Sha1([]byte(pwd + pwd_salt))

	res := dblayer.UserRegister(phone, enc_pwd)
	if res {
		w.Write([]byte("{\"code\":0,\"msg\":\"user register succeeded.\"}"))
	} else {
		w.Write([]byte("{\"code\":-1,\"msg\":\"user register failed.\"}"))
	}
}

// handle login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	phone := r.Form.Get("phone")
	pwd := r.Form.Get("pwd")

	if len(phone) < 3 || len(phone) > 32 || len(pwd) != 32 {
		w.Write([]byte("{\"code\":-1,\"msg\":\"params invalid.\"}"))
		return
	}

	enc_pwd := util.Sha1([]byte(pwd + pwd_salt))

	res := dblayer.UserLogin(phone, enc_pwd)
	if res {
		token := genLoginToken(phone)
		res = dblayer.UserUpdateToken(phone, token)
		if res {
			msg := fmt.Sprintf("{\"code\":0,\"msg\":\"user login succeeded.\",\"token\":\"%s\"}", token)
			w.Write([]byte(msg))
			return
		}
	}
	w.Write([]byte("{\"code\":-1,\"msg\":\"user login failed.\"}"))
}
