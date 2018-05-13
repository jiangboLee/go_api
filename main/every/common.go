package every

import (
	"net/http"
	"html/template"
	"go_api/main/utils"
	"fmt"
	"go_api/main/data"
	"os"
	"io"
	"strings"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func SaveImage(w http.ResponseWriter, r *http.Request) {
	//fmt.Printf(GetCurrPath())
	//fill_path, _ := filepath.Rel(GetCurrPath(),"/go_api/main/html/common/uploadImg.html")
	t, err := template.ParseFiles(GetCurrPath() + "/html/common/uploadImg.html")
	if err != nil {
		utils.Danger(err, "模板有错")
		fmt.Fprintf(w, "%s%s", GetCurrPath(), err)
		return
	}
	t.Execute(w, nil)
}
//返回当前的绝对路径
func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

const UPLOAD_PATH string = utils.SYSTEM_PATH + "/imgs/"

func UploadImage(w http.ResponseWriter, r *http.Request) {
	var img data.Image
	r.ParseMultipartForm(1024)
	imgFile, imgHead, imgErr := r.FormFile("img")
	if imgErr != nil {
		 fmt.Println(imgErr)
		 return
	}
	defer imgFile.Close()

	img.Imageurl = imgHead.Filename
	image, err := os.Create(UPLOAD_PATH + img.Imageurl)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer image.Close()

	_, err = io.Copy(image, imgFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = img.Create()
	if err != nil {
		utils.Danger(err, "插入图片失败")
		fmt.Fprintf(w,"%s,%s",err,"上传失败")
		return
	}
	fmt.Fprintf(w,"上传成功: www.ljbniubi.top/img/look/%s",imgHead.Filename)
}
//查看图片
func LookImg(w http.ResponseWriter, r *http.Request)  {
	fmt.Printf(r.URL.Scheme,r.URL.User)
	path := strings.Split(r.URL.Path,"/")
	var name string
	if len(path) > 1 {
		name = path[len(path) - 1]
	} else {
		name = "null.png"
	}
	fmt.Printf(name)
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Disposition",fmt.Sprintf("inline; filename=\"%s\"",name))

	file, err := ioutil.ReadFile(UPLOAD_PATH + name)
	if err != nil {
		fmt.Fprintf(w,"查无此图片")
		return
	}
	w.Write(file)
}
