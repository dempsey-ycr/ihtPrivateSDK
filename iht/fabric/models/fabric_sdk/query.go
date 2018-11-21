package fabric_sdk

import (
	"errors"
	"ihtPrivateSDK/share/logging"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"

	. "ihtPrivateSDK/iht/fabric/models"
	pb "protobuf/projects/go/protocol/common"
)

// Query query chaincode
func Query(request channel.Request) (*channel.Response, error) {
	client := createClient()
	if client == nil {
		return nil, errors.New("Query: Create channel client failed")
	}

	response, err := client.Query(request, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		logging.Error("Query: %v", err)
		return nil, err
	}
	// Notifier(client, getEventID())
	return &response, nil
}

// DisposalQuery Get accurate data by processing
func DisposalQuery(response *channel.Response) (*Response, error) {
	var endorsements []*pb.Endorsement
	for _, v := range response.Responses {
		logging.Debug(v.Endorser)
		endorser := &pb.Endorsement{
			Status:      v.Status,
			EndorseAddr: v.Endorser,
			Signature:   v.Endorsement.Signature,
			Version:     v.Version,
			// v.Endorsement.Endorser 背书证书
		}
		if v.Timestamp != nil {
			endorser.Timestamp = v.Timestamp.Seconds
		}
		endorsements = append(endorsements, endorser)
	}

	private := &pb.FabricPrivate{
		ChaincodeStatus:  response.ChaincodeStatus,
		TransactionID:    string(response.TransactionID),
		TxValidationCode: int32(response.TxValidationCode),
		Endorsement:      endorsements,
	}

	res := &Response{
		Code:          (response.Responses)[0].Response.Status,
		Message:       (response.Responses)[0].Response.Message,
		Payload:       (response.Responses)[0].Response.Payload,
		FabricPrivate: private,
	}

	return res, nil
}
