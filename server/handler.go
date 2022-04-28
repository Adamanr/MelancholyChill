package HandlerRoot

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler(r *gin.Engine) {
	r = gin.New()
	r.GET("/", index)
	r.Static("/assets", "client/assets")
	r.LoadHTMLGlob("client/templates/**/*.tmpl")
	r.Run()
}
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Главная",
	})
}
