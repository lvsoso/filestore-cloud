package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	futil "github.com/moxiaomomo/filestore-cloud/util"
)

const (
	upload_path string = "/tmp/files/"
)

func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Upload finished!")
}

func FileListHandler(w http.ResponseWriter, r *http.Request) {
	flist, err := ioutil.ReadDir(upload_path)
	if err != nil {
		io.WriteString(w, "Get filelist error.")
	} else {
		files := ""
		for _, v := range flist {
			files += (v.Name() + "&#10;")
		}
		io.WriteString(w, files)
	}
}

func FileDelHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	filehash := r.Form.Get("filehash")

	flist, err := ioutil.ReadDir(upload_path)
	if err != nil {
		io.WriteString(w, "Get filelist error.")
	} else {
		for _, v := range flist {
			file, err := os.Open(upload_path + v.Name())
			if err != nil {
				continue
			}

			tmpSha1 := futil.FileSha1(file)
			file.Close()

			fmt.Println(v.Name() + " " + tmpSha1 + "  " + filehash)
			if tmpSha1 == filehash {
				os.Remove(upload_path + v.Name())
			}
		}
		io.WriteString(w, "File removed.")
	}
}

func FileUploadHandle(w http.ResponseWriter, r *http.Request) {
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
