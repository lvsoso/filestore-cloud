package main

import (
	"fmt"
	"net/http"

	hdl "github.com/moxiaomomo/filestore-cloud/handler"
)

func main() {
	fmt.Println("to start upload server...")

	http.HandleFunc("/user/signup", hdl.RegisterHandler)
	http.HandleFunc("/user/signin", hdl.LoginHandler)

	http.HandleFunc("/file/upload/success", hdl.AccessAuth(hdl.UploadSucHandler))
	http.HandleFunc("/file/list", hdl.AccessAuth(hdl.FileListHandler))
	http.HandleFunc("/file/upload", hdl.AccessAuth(hdl.FileUploadHandle))
	http.HandleFunc("/file/delete", hdl.AccessAuth(hdl.FileDelHandler))

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Printf("failed to start server, err: %v\n", err)
		return
	}
	fmt.Println("server eixted.")
}
