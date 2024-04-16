# go-bindata-assetfs
静态资源打包组件

###### 首先安装静态资源打包组件
```
go get github.com/go-bindata/go-bindata/...
go get github.com/elazarl/go-bindata-assetfs/...
```

###### 打包静态资源
```
go-bindata-assetfs -pkg catchace static/...
```

```
Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
d-----         2024/4/16     14:28                static
d-----         2024/4/16     14:28                static/test/index.html
d-----         2024/4/16     14:28                static/index.html
-a----         2024/4/16     14:28        2806354 bindata_assetfs.go
-a----         2024/4/16     14:28           1010 route.go
-a----         2024/4/16     14:28           1010 main.go
```

###### 使用
```
func RouteConfig(r *gin.Engine) {
	//r.StaticFS("/game", http.Dir("static"))
	r.StaticFS("/game", assetFS()) //bindata_assetfs.go
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/game/index.html")
        // 127.0.0.1:80/game/index.html  127.0.0.1:80/game/test/index.html  
	})
    //POST
	r.POST("/CatchAce/create", createCatchAce)
	// GET
	r.GET("/CatchAce", listCatchAce)
}
```
