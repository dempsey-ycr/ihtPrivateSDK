package routes

import (
	"ihtPrivateSDK/iht/ipfs/controllers/ipfs"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	engine.GET("/", ipfs.NewTestInfo().GET)

	rg := engine.Group("/api/v0")

	// publish
	RegPublish(rg)
}
