package data

import "fmt"

type Image struct {
	Imageurl string
}

//新建图片
func (img *Image) Create() (err error) {
	statement := "insert image set imageurl=?"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(img.Imageurl)
	if err != nil {
		return err
	} else {
		return nil
	}
}
//根据图片名字返回图片
