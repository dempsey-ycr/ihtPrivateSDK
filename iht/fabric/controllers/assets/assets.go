package assets

import (
	"encoding/json"
	"ihtPrivateSDK/share/lib"

	sdk "ihtPrivateSDK/iht/fabric/models/fabric_sdk"
	pro "protobuf/projects/go/protocol/asset"

	"ihtPrivateSDK/share/logging"

	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

// Response ...
type Response struct {
	Code   int         `json:"code"`   // 返回的code
	ErrMsg string      `json:"error"`  // 失败描述
	Data   interface{} `json:"data"`   // 成功有可能会返回这个，json对象或数组或字符串
	TxID   string      `json:"txId"`   // 交易ID
	TxCode interface{} `json:"txCode"` // 交易id类型
}

// AssetManages ...
type AssetManages struct{}

// NewAssets ...
func NewAssets() *AssetManages {
	return &AssetManages{}
}

// POST asset payload
func (m *AssetManages) POST(c *gin.Context) {
	asset := &pro.RequestAssetPut{}
	if code, err := lib.RecvAndUnmarshalJSON(c, 1024, asset); err != nil {
		logging.Error("post json %v", err)
		lib.WriteString(c, code, nil)
		return
	}

	value, err := json.Marshal(asset.Value)
	if err != nil {
		logging.Error("Marshal: %v", err)
		lib.WriteString(c, 40002, nil)
		return
	}

	req := channel.Request{
		ChaincodeID: "byfn",
		Fcn:         asset.Fun,
		Args:        [][]byte{[]byte(asset.Key), value},
	}
	res, err := sdk.Invoke(req)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	v := &Response{
		TxID:   string(res.TransactionID),
		TxCode: res.TxValidationCode,
	}

	_, err = sdk.Disposal(res)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	if err = json.Unmarshal(nil, v); err != nil {
		logging.Error("%v", err)
		lib.WriteString(c, 40002, nil)
		return
	}

	lib.WriteJSON(c, v)
}

// GET asset
func (m *AssetManages) GET(c *gin.Context) {
	fcn := c.Query("func")
	key := c.Query("key")

	req := channel.Request{
		ChaincodeID: "byfn",
		Fcn:         fcn,
		Args:        [][]byte{[]byte(key)},
	}

	res, err := sdk.Query(req)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	response, err := sdk.DisposalQuery(res)
	if err != nil {
		lib.WriteString(c, 40002, nil)
		return
	}

	lib.WriteJSON(c, response)
}
