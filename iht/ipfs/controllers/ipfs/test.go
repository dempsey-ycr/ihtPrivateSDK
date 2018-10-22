package ipfs

import (
	"ihtPrivateSDK/share/lib"

	"github.com/gin-gonic/gin"
)

type TestInfo struct{}

func NewTestInfo() *TestInfo {
	return &TestInfo{}
}

func (this *TestInfo) GET(c *gin.Context) {
	lib.WriteString(c, 200, "hello world...")
}
