引言

仿网盘，总体来说还是不错。



# 接口API

| URL           | method | 功能               |
| ------------- | ------ | ------------------ |
| /login        | POST   | 用户注册           |
| /register     | POST   | 用户登录           |
| /uploadfile   | POST   | 上传文件           |
| /downloadfile | POST   | 下载文件           |
| /shareQRcode  | POST   | 创建二维码分享链接 |
| /sharesecret  | POST   | 加密分享链接       |



# 功能

## 文件上传

```
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
```

思路很简单，拿到想要上传的参数文件，然后上传保存至指定的文件夹

## 文件下载

```
func Downloadfile(c *gin.Context)  {
   //下载的话  拿到路径就可以了
   	filename:=c.PostForm("filename")
	paths:=c.PostForm("path")
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File(paths)//文件当前所在目录

	msg:="Download seccess!"
	resps.Ok(c,msg)
}
```

仔细找了找，也看了Go文档，发现官方有实现上传文件的功能却没有下载文件的功能。花了老半天时间才发现c.File这个好东西。

## 登陆注册

老朋友了，这里我就不上代码了。

## 权限管理

用的是jwt，下面是中间件

```
router.POST("/register", controllers.Register)
router.POST("/login", controllers.Login)
router.Use(middlewares.User)
{
   router.POST("/uploadfile", controllers.Uploadfile)
   router.POST("/downloadfile", controllers.Downloadfile)
   //生成二维码分享链接
   router.POST("/shareQRcode", controllers.ShareQRcode)
   //加密分享
   router.POST("/sharesecret", controllers.Sharesecret)
}
```

登录的时候会创建token，后续的操作会有中间件来识别token是否正确



```
func User(c *gin.Context) {
   auth:= c.GetHeader("Authorization")
   fmt.Println(auth)
   if len(auth)<7 {
      resps.Error(c, "token error")
      c.Abort()
      return
   }
   token := auth[7:]
    err := jwt.CheckToken(token)
   if err != nil {
      resps.Error(c, "token error")
      c.Abort()
      return
   }
   c.Next()
   return
}
```

jwt的代码太长了我就不上代码了。

## 文件分享





## 进阶

### 二维码分享

调用了github大牛的生成二维码库。

```
func ShareQRcode(c *gin.Context)  {
   //这个是拿到想要分享的链接的路径
   path:=c.PostForm("path")
   _ = qrcode.WriteFile(path, qrcode.Medium, 256, "qr.png")
}
```

下面是自己调试的时候生成的二维码。

![image-20200729213328449](C:\Users\Mechrevo\AppData\Roaming\Typora\typora-user-images\image-20200729213328449.png)

## 加密分享

用的是MD5加密

```
func Sharesecret(c *gin.Context){
   //输入想要分享的东西的路径
   path:=c.PostForm("path")
   s:=md5.New()
   s.Write([]byte(path))   // 带加密数据
   signData:=s.Sum(nil)

   c.JSON(200,gin.H{"status":http.StatusOK,"":signData})
}
```



# 项目目录

```
├──cmd
├──controllers
├──jwt
├──middlewares
├──model
├──resps
├──router
└──service
```



# 最后想说的

这次实现的代码，自我最满意的是设计的分布式设计。修改代码的时候也确实非常方便修改，