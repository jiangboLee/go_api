package every

import (
	"net/http"
	"log"
	"os"
	"io"
)

func SaveImage(w http.ResponseWriter, r *http.Request)  {

}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	//随机生成一个不存在的fileid
	var imgid string
	for{
		imgid=MakeImageID()
		if !FileExist(ImageID2Path(imgid)){
			break
		}
	}
	//上传参数为uploadfile
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("uploadfile")
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error:Upload Error."))
		return
	}
	defer file.Close()
	//检测文件类型
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error:Upload Error."))
		return
	}
	filetype := http.DetectContentType(buff)
	if filetype!="image/jpeg"{
		w.Write([]byte("Error:Not JPEG."))
		return
	}
	//回绕文件指针
	log.Println(filetype)
	if  _, err = file.Seek(0, 0); err!=nil{
		log.Println(err)
	}
	//提前创建整棵存储树（如果不进行存储树结构创建，下面的文件创建不会成功）
	if err=BuildTree(imgid); err!=nil{
		log.Println(err)
	}
	//将文件写入ImageID指定的位置
	f, err := os.OpenFile(ImageID2Path(imgid), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err)
		w.Write([]byte("Error:Save Error."))
		return
	}
	defer f.Close()
	io.Copy(f, file)
	w.Write([]byte(imgid))
}