package main

import (
	HandlerRoot "MelancholyChill/server"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	HandlerRoot.Handler(r)
}
