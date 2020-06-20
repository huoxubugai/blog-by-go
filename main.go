package main

import (
	"blog/dao/db"
	"blog/helper"
	"github.com/gin-gonic/gin"
	"html/template"
	"os"
	"path/filepath"
	"strings"
)
import "blog/controllers"

func main() {
	router := gin.Default()
	dns := "root:wang6490918@tcp(localhost:3306)/blog?charset=utf8&parseTime=True&loc=Local"
	db, err := db.InitDb(dns)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//router.Static("/static", filepath.Join(getCurrentDirectory(), "./static"))
	router.Static("/static/", "./static")
	setTemplate(router)
	router.GET("/", controllers.IndexGet)
	router.GET("/blog/:id", controllers.BlogGet)
	router.GET("/about", controllers.AboutMe)
	router.GET("/types/:id", controllers.TypesGet)
	router.GET("/tags/:id", controllers.TagsGet)
	router.Run()
}

func setTemplate(engine *gin.Engine) {
	funcMap := template.FuncMap{
		"dateFormat": helper.DateFormat,
	}
	engine.SetFuncMap(funcMap)
	engine.LoadHTMLGlob("views/*")
}
func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		//seelog.Critical(err)
		panic(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
