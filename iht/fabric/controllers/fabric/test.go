package fabric

import (
	"encoding/json"
	"ihtPrivateSDK/share/lib"

	"ihtPrivateSDK/share/logging"

	sdk "ihtPrivateSDK/iht/fabric/models/fabric_sdk"

	"github.com/gin-gonic/gin"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// TestInfo test info
type TestInfo struct {
	chaincodeID string
}

// NewTestInfo new a testInfo
func NewTestInfo() *TestInfo {
	return &TestInfo{
		chaincodeID: "app",
	}
}

// GET test
func (t *TestInfo) GET(c *gin.Context) {
	lib.WriteString(c, 200, "hello world...")
}

//------------------------------------ chaincode operate test --------------------------------------------//

// ReqInvokeTest request invoke test
type ReqInvokeTest struct {
	Key   string     `json:"key"`
	Fun   string     `json:"fun"`
	Value InvokeInfo `json:"value"`
}

// InvokeInfo info
type InvokeInfo struct {
	User    string
	Age     int
	Message string
}

// Query chaincode query test
func (t *TestInfo) Query(c *gin.Context) {
	fcn := c.Query("func")
	key := c.Query("key")

	req := channel.Request{
		ChaincodeID: t.chaincodeID,
		Fcn:         fcn,
		Args:        [][]byte{[]byte(key)},
	}

	res, err := sdk.Query(req)
	if err != nil {
		lib.WriteSDK(c, 40002, "", nil)
		return
	}

	v := &InvokeInfo{}
	dis, err := sdk.DisposalQuery(res)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	if err = json.Unmarshal(dis.Response.Payload, v); err != nil {
		logging.Error("%v", err)
		lib.WriteSDK(c, 40002, "", nil)
		return
	}

	lib.WriteSDK(c, dis.Response.Status, dis.Response.Message, v)
}

// Invoke chaincode invoke test
func (t *TestInfo) Invoke(c *gin.Context) {
	r := &ReqInvokeTest{}
	if code, err := lib.RecvAndUnmarshalJSON(c, 1024, r); err != nil {
		logging.Error("post json %v", err)
		lib.WriteString(c, code, nil)
		return
	}

	value, err := json.Marshal(r.Value)
	if err != nil {
		logging.Error("Marshal: %v", err)
		lib.WriteString(c, 40002, nil)
		return
	}
	req := channel.Request{
		ChaincodeID: t.chaincodeID,
		Fcn:         r.Fun,
		Args:        [][]byte{[]byte(r.Key), value},
	}
	res, err := sdk.Invoke(req)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	dis, err := sdk.Disposal(res)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	lib.WriteSDK(c, dis.Response.Status, dis.Response.Message, string(dis.Response.Payload))
}
