package basal

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
		chaincodeID: "byfn",
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
	Fun   string     `json:"method"`
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
		lib.WriteError(c, 40002, err.Error(), nil)
		return
	}

	payload, _ := sdk.DisposalQuery(res)
	if payload == nil {
		logging.Error("DisposalQuery error...")
		lib.WriteError(c, 40002, "", nil)
		return
	}

	var metadata map[string]interface{} // 把[]byte按照其原来结构解析出来，不关心其类型
	data := (payload.Payload).([]byte)
	if len(data) != 0 {
		if err = json.Unmarshal(data, &metadata); err != nil {
			logging.Error("json.Unmarshal:%v", err)
			payload.Code = 40002
			lib.WriteJSON(c, payload.FabricPrivate)
			return
		}
		payload.Payload = &metadata
	}

	lib.WriteJSON(c, payload)
}

// Invoke chaincode invoke test
func (t *TestInfo) Invoke(c *gin.Context) {
	r := &ReqInvokeTest{}
	if code, err := lib.RecvAndUnmarshalJSON(c, 1024, r); err != nil {
		lib.WriteError(c, code, err.Error(), nil)
		return
	}

	value, err := json.Marshal(r.Value)
	if err != nil {
		lib.WriteError(c, 40002, err.Error(), nil)
		return
	}
	req := channel.Request{
		ChaincodeID: t.chaincodeID,
		Fcn:         r.Fun,
		Args:        [][]byte{[]byte(r.Key), value},
	}
	res, err := sdk.Invoke(req)
	if err != nil {
		lib.WriteError(c, 40002, err.Error(), nil)
		return
	}

	payload, _ := sdk.Disposal(res)
	if payload == nil {
		logging.Error("DisposalQuery error...")
		lib.WriteError(c, 40002, "", nil)
		return
	}

	var metadata map[string]interface{} // 把[]byte按照其原来结构解析出来，不关心其类型
	data := (payload.Payload).([]byte)
	if len(data) != 0 {
		if err = json.Unmarshal(data, &metadata); err != nil {
			logging.Error("json.Unmarshal:%v", err)
			payload.Code = 40002
			lib.WriteJSON(c, payload.FabricPrivate)
			return
		}
		payload.Payload = &metadata
	}

	lib.WriteJSON(c, payload)
}
