package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
    "io/ioutil"
)

const (
	upload_path string = "/tmp/"
)

func load_success(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	// GET 方法获取上传主页
	if r.Method == "GET" {
        b, err := ioutil.ReadFile("./static/view/index.html")
        if err != nil {
            return
        }
		io.WriteString(w, string(b))

    // POST 方法获取文件上传内容
	} else {
		file, head, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fW, err := os.Create(upload_path + head.Filename)
		if err != nil {
			fmt.Println("failed to create a file.")
			return
		}
		defer fW.Close()
		_, err = io.Copy(fW, file)
		if err != nil {
			fmt.Println("failed to save the file.")
			return
		}
		http.Redirect(w, r, "/success", http.StatusFound)
	}
}

func main() {
	fmt.Println("to start upload server...")
	http.HandleFunc("/success", load_success)
	http.HandleFunc("/upload", uploadHandle)

	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Printf("failed to start server, err: %v\n", err)
		return
	}
	fmt.Println("server eixted.")
}
