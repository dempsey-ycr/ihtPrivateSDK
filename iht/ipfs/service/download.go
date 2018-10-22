package service

import (
	"ihtPrivateSDK/share/logging"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"ihtPrivateSDK/share/lib"

	"github.com/gin-gonic/gin"
)

// Download The struct for download file
type Download struct {
	Hash string `json:"hash"`
	Data []byte `json:"bytes"`
}

// GET The download method
func (f *Attachment) GET(c *gin.Context) {
	dir := "./third_party/filetmp"
	hash := c.Query("hash")
	if err := f.ipfs.Get(hash, dir); err != nil {
		logging.Error("ipfs get: %v", err.Error)
		lib.WriteString(c, 40002, nil)
		return
	}

	data, err := ioutil.ReadFile(dir + "/" + hash)
	if err != nil {
		logging.Debug("ipfs get: %v", err.Error)
		lib.WriteString(c, 40002, nil)
		return
	}
	r := &Download{
		Hash: hash,
		Data: data,
	}
	// lib.DeleteFile(dir + "/" + hash)

	lib.WriteString(c, 200, r)
}

// FileServer test
func (f *Attachment) FileServer(c *gin.Context) {

	dir := "./third_party/filetmp"
	hash := c.Query("hash")

	if err := f.ipfs.Get(hash, dir); err != nil {
		logging.Error("ipfs get: %v", err.Error)
		lib.WriteString(c, 40002, nil)
		return
	}

	http.FileServer(http.Dir(dir + "/" + hash))

}

// DownloadFile test
func (f *Attachment) DownloadFile(c *gin.Context) {
	dir := "./third_party/filetmp"
	hash := c.Query("hash")
	if err := f.ipfs.Get(hash, dir); err != nil {
		logging.Error("ipfs get: %v", err.Error)
		return
	}
	fileFullPath := dir + "/" + hash
	file, err := os.Open(fileFullPath)
	if err != nil {
		logging.Error("Open tmp file: %v", err.Error)
		return
	}

	defer file.Close()
	fileName := path.Base(fileFullPath)
	fileName = url.QueryEscape(fileName) // 防止中文乱码

	c.Header("Content-Type", "application/octet-stream")
	c.Header("content-disposition", "attachment; filename=\""+fileName+"\"")

	_, err = io.Copy(c.Writer, file)
	if err != nil {
		logging.Error("response write:%v", err)
		return
	}
}
