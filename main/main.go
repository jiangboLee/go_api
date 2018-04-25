package main

import (
	"net/http"
	"fmt"
	"goapi/wx"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w,"hello lee-go")
}


func main() {

	http.HandleFunc("/",handler)
	http.HandleFunc("/wx/api/appid",wx.GetAppid)
	http.ListenAndServeTLS(":443","1_www.ljbniubi.top_bundle.crt","2_www.ljbniubi.top.key", nil)
}
