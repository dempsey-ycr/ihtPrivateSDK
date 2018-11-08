package routes

import (
	"ihtPrivateSDK/iht/fabric/controllers/basal"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/", basal.NewTestInfo().GET)

	rg := engine.Group("/api")

	// publish
	RegPublish(rg)
}
