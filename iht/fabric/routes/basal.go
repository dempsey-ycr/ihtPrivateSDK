package routes

import (
	"ihtPrivateSDK/iht/fabric/controllers/assets"
	"ihtPrivateSDK/iht/fabric/controllers/basal"

	"github.com/gin-gonic/gin"
)

// RegPublish register ipfs routes
func RegPublish(rg *gin.RouterGroup) {

	rg.GET("/sdk/query", basal.NewTestInfo().Query)
	rg.POST("/sdk/invoke", basal.NewTestInfo().Invoke)

	// 资产信息
	rg.GET("/asset/get", assets.NewAssets().GET)
	rg.POST("/asset/put", assets.NewAssets().POST)

	// Some general methods for manipulating chained objects
	rg.POST("/verse/insert", basal.NewMetaHandle().Insert)
	rg.POST("/verse/change", basal.NewMetaHandle().Change)
	rg.POST("/verse/delete", basal.NewMetaHandle().DeleteSingle)
	rg.POST("/verse/read_single", basal.NewMetaHandle().ReadSingle)
	rg.POST("/verse/trace_history", basal.NewMetaHandle().TraceHistory)
}
