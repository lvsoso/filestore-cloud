package mpupload

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func InitiateMultipartUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	filesize, err := strconv.ParseInt(r.Form.Get("filesize"), 10, 64)
	if err != nil {
		w.Write([]byte("{\"code\":-1,\"msg\":\"params invalid.\"}"))
		return
	}

	upInfo := InitiateMultipartUpload(filesize)
	data, err := json.Marshal(upInfo)
	if err != nil {
		w.Write([]byte("{\"code\":-2,\"msg\":\"internel server error.\"}"))
		return
	}
	w.Write([]byte("{\"code\":0,\"msg\":\"ok\",\"data\":" + string(data) + "}"))
}
