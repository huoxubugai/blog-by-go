package main

import (
	"blog/dao/db"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"strings"
)
import "blog/controllers"

func main(){
	router:=gin.Default()
	dns := "root:wang6490918@tcp(localhost:3306)/blog?parseTime=true"
	db.InitDb(dns)
	//router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.Static("/static/","./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controllers.IndexGet)
	router.Run()
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//seelog.Critical(err)
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
