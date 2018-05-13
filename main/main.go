package main

import (
	"net/http"
	"fmt"
	"go_api/main/wx"
	"go_api/main/every"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"hello https")
}


func main() {

	http.HandleFunc("/", handler)
	http.HandleFunc("/wx/api/appid", wx.GetAppid)
	/*图片上传*/
	http.HandleFunc("/img/upload", every.SaveImage)
	http.HandleFunc("/img/uploading", every.UploadImage)
	http.HandleFunc("/img/look/",every.LookImg)

	//http.ListenAndServe(":8080",nil)
	http.ListenAndServeTLS(":443","1_www.ljbniubi.top_bundle.crt","2_www.ljbniubi.top.key", nil)
}
