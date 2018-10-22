package routes

import (
	"ihtPrivateSDK/iht/fabric/controllers/fabric"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/", fabric.NewTestInfo().GET)

	rg := engine.Group("/api/sdk")

	// publish
	RegPublish(rg)
}
