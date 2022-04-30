package HandlerRoot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func Handler(r *gin.Engine) {
	r = gin.New()
	r.GET("/", index)
	r.GET("/playlist", playlist)
	r.POST("/register", nil)
	r.POST("/login", auth)
	r.GET("/library", library)
	r.Static("/assets", "client/assets")
	r.LoadHTMLGlob("client/templates/**/*.tmpl")
	r.Run()
}
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Главная",
	})
}

var aqua []string

func playlist(c *gin.Context) {
	files, _ := ioutil.ReadDir("client/assets/files/")
	for _, file := range files {
		fmt.Println(file.Name())
		fileName := strings.Trim(file.Name(), ".mp3")
		aqua = append(aqua, fileName)
	}
	c.HTML(http.StatusOK, "playlist.tmpl", gin.H{
		"title":   "Плейлисты",
		"clients": aqua,
	})
}

func library(c *gin.Context) {
	files, _ := ioutil.ReadDir("client/assets/files/")
	for _, file := range files {
		fmt.Println(file.Name())
		fileName := strings.Trim(file.Name(), ".mp3")
		aqua = append(aqua, fileName)
	}
	c.HTML(http.StatusOK, "library.tmpl", gin.H{
		"title":   "Библиотека",
		"clients": aqua,
	})
}
func auth(c *gin.Context) {

}
