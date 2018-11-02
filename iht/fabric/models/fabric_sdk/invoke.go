package fabric_sdk

import (
	"errors"

	"ihtPrivateSDK/share/logging"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"

	. "ihtPrivateSDK/iht/fabric/models"
	pb "protobuf/projects/go/protocol/basic"
)

// Invoke Operate chaincode
func Invoke(request channel.Request) (*channel.Response, error) {
	client := createClient()
	if client == nil {
		return nil, errors.New("Invoke: Create channel client failed")
	}

	response, err := client.Execute(request, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		logging.Error("Invoke, client execute error: %v", err)
		return nil, err
	}
	// fmt.Println("chaincodeStatus: ", response.ChaincodeStatus)
	// fmt.Println("Proposal.TxnID: ", response.Proposal.TxnID)
	// fmt.Println("Proposal Header: ", string(response.Proposal.Header))
	// fmt.Println("Proposal.Payload: ", string(response.Proposal.Payload))
	// fmt.Println("Proposal.Extension: ", string(response.Proposal.Extension))
	// logging.Info("TransactionID: %s", response.TransactionID)
	// logging.Info("TxValidationCode: %d", response.TxValidationCode)
	// fmt.Println("Responses: ", (response.Responses)[0].Status)
	// fmt.Println("Responses: ", (response.Responses)[0].Endorser)
	// fmt.Println("Responses: ", (response.Responses)[0].ChaincodeStatus)
	// fmt.Println("Responses: ", (response.Responses)[0].Version)
	// fmt.Println("Responses: ", (response.Responses)[0].Timestamp)
	// fmt.Println("Responses payload: ", string((response.Responses)[0].Payload))
	// fmt.Println("Responses: ", (response.Responses)[0].Endorsement)

	Notifier(client, getEventID())
	return &response, nil
}

// Disposal Get accurate data by processing
func Disposal(response *channel.Response) (*Response, error) {
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
