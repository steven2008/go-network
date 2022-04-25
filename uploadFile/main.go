package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("listen err", err)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	var result = make(map[string]string)
	for k, v := range r.MultipartForm.Value {
		result[k] = v[0]
	}
	fmt.Println("result", result)

	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println("uploadfile", err)
		return
	}
	defer file.Close()

	fh, err := handler.Open()
	if err != nil {
		fmt.Println("fh err", err)
	}
	defer fh.Close()
	imageBuf, err := ioutil.ReadAll(fh)
	if err != nil {
		fmt.Println("imageBuf err", err)
	}
	path := "/Volumes/Macintosh HD/Users/rockonterol/Desktop/ss/"
	folderName := time.Now().Format("2006-01-02")
	folderPath := filepath.Join(path, folderName)
	//golang判断文件或文件夹是否存在的方法为使用os.Stat()函数返回的错误值进行判断:
	//如果返回的错误为nil,说明文件或文件夹存在
	//如果返回的错误类型使用os.IsNotExist()判断为true,说明文件或文件夹不存在
	//如果返回的错误为其它类型,则不确定是否在存在
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		//先创建文件，然后赋予操作权限
		os.Mkdir(folderPath, 0777)
		os.Chmod(folderPath, 0777)
		fmt.Println("=============111111")
	}
	err = ioutil.WriteFile(folderPath+"/"+handler.Filename+".jpg", imageBuf, 0777)
	if err != nil {
		fmt.Println("WriteFile err", err)
	}
	//第二种方式保存图片
	f, err := os.OpenFile("/Volumes/Macintosh HD/Users/rockonterol/Desktop/ss/111.jpg", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		fmt.Println("OpenFile err", err)
	}
	defer f.Close()
	io.Copy(f, file)
}
