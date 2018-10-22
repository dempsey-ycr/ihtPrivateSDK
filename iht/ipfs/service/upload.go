package service

import (
	"sync"

	"github.com/gin-gonic/gin"
	shell "github.com/ipfs/go-ipfs-api"

	"ihtPrivateSDK/share/lib"
	"ihtPrivateSDK/share/logging"
)

// ipfs class
var (
	instance *IPFSCtl
	once     sync.Once
)

type IPFSCtl struct {
	ipfs *shell.Shell
}

func GetInstance() *IPFSCtl {
	once.Do(func() {
		instance = &IPFSCtl{
			ipfs: shell.NewShell("localhost:5001"),
		}
	})
	return instance
}

// ipfs operation
type Attachment struct {
	ipfs *shell.Shell
}

func NewAttachment() *Attachment {
	return &Attachment{
		ipfs: GetInstance().ipfs,
	}
}

// response struct
type ResponseJson struct {
	Hash    string `json:"hash"`
	SrcName string `json:"srcName"`
}

func (f *Attachment) POST(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(32 << 20); err != nil {
		logging.Error("%v", err)
		lib.WriteString(c, 40002, err.Error)
		return
	}

	file, header, err := c.Request.FormFile("filename")
	if err != nil {
		logging.Error("FromFile error:%v", err)
		lib.WriteString(c, 40002, err.Error)
		return
	}

	defer file.Close()
	res, err := f.ipfs.Add(file)
	if err != nil {
		logging.Error("%v", err)
		lib.WriteString(c, 40002, err.Error)
		return
	}

	resJson := &ResponseJson{
		Hash:    res,
		SrcName: header.Filename,
	}
	lib.WriteString(c, 200, resJson)
}
