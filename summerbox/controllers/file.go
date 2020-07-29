package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"path"
	"summerbox/resps"
)

//上传文件的话  从postman获取路径
func Uploadfile(c *gin.Context)  {
	//思路是从Postman读取文件并将其保存至某个文件夹
	file,_:=c.FormFile("file")
	//接下来就是把下载的文件保存到某个文件夹

	//自动保存在这个文件夹
	dst:=path.Join("./upload",file.Filename)
	// dst := path.Join("./upload", "tupian.jpg")

	//将上传的文件保存到本地服务器的指定位置
	_ = c.SaveUploadedFile(file, dst)
	// 返回数据
	msg:="Upload seccess!"
	resps.Ok(c,msg)
}

func Downloadfile(c *gin.Context)  {

	filename:=c.PostForm("filename")
	paths:=c.PostForm("path")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(paths)//文件当前所在目录

	msg:="Download seccess!"
	resps.Ok(c,msg)
}

