package routes

import (
	"ihtPrivateSDK/iht/fabric/controllers/assets"
	"ihtPrivateSDK/iht/fabric/controllers/fabric"

	"github.com/gin-gonic/gin"
)

// RegPublish register ipfs routes
func RegPublish(rg *gin.RouterGroup) {

	rg.GET("/query", fabric.NewTestInfo().Query)
	rg.POST("/invoke", fabric.NewTestInfo().Invoke)

	// 资产信息
	rg.GET("/asset/get", assets.NewAssets().GET)
	rg.POST("/asset/put", assets.NewAssets().POST)
}
