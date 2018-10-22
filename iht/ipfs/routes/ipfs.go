package routes

import (
	"ihtPrivateSDK/iht/ipfs/service"

	"github.com/gin-gonic/gin"
)

// RegPublish register ipfs routes
func RegPublish(rg *gin.RouterGroup) {

	rg.POST("/add", service.NewAttachment().POST)
	rg.GET("/get", service.NewAttachment().DownloadFile)
	// rg.GET("/get", service.NewAttachment().GET)
	// rg.GET("/server", service.NewAttachment().FileServer)
}
