package basal

import (
	"encoding/json"
	. "ihtPrivateSDK/iht/fabric/models"
	sdk "ihtPrivateSDK/iht/fabric/models/fabric_sdk"
	"ihtPrivateSDK/share/lib"
	"strconv"

	"haina.com/share/logging"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// ReadSingle ...
func (m *MetadataHandle) DeleteSingle(c *gin.Context) {
	m.chaincodeID = c.Query("chaincodeID")
	m.version = c.Query("version")
	method := "Delete"

	metaReq := &RequestMetadata{}
	if code, err := lib.RecvAndUnmarshalJSON(c, 1024, metaReq); err != nil {
		lib.WriteError(c, code, err.Error(), nil)
		return
	}
	logging.Info("meta cond: %#v", metaReq.MetaData)
	value, err := json.Marshal(metaReq.MetaData)
	if err != nil {
		lib.WriteError(c, 40002, err.Error(), nil)
		return
	}

	req := channel.Request{
		ChaincodeID: m.chaincodeID,
		Fcn:         method,
		Args:        [][]byte{[]byte(strconv.Itoa(metaReq.ObjType)), value},
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
