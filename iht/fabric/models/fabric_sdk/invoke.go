package fabric_sdk

import (
	"errors"

	"ihtPrivateSDK/share/logging"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

// Invoke Operate chaincode
func Invoke(request channel.Request) (*channel.Response, error) {
	client := createClient()
	if client == nil {
		return nil, errors.New("Invoke: Create channel client failed")
	}

	response, err := client.Execute(request, channel.WithRetry(retry.DefaultChannelOpts))
	if err != nil {
		logging.Error("Invoke: %v", err)
		return nil, err
	}
	// fmt.Println("chaincodeStatus: ", response.ChaincodeStatus)
	// fmt.Println("Proposal.TxnID: ", response.Proposal.TxnID)
	// fmt.Println("Proposal Header: ", string(response.Proposal.Header))
	// fmt.Println("Proposal.Payload: ", string(response.Proposal.Payload))
	// fmt.Println("Proposal.Extension: ", string(response.Proposal.Extension))
	logging.Info("TransactionID: %s", response.TransactionID)
	logging.Info("TxValidationCode: %d", response.TxValidationCode)
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
func Disposal(response *channel.Response) (*fab.TransactionProposalResponse, error) {

	if response.ChaincodeStatus == 200 && len(response.Responses) == 1 {
		return (response.Responses)[0], nil
	}
	if response.ChaincodeStatus == 200 && len(response.Responses) > 1 {
		return nil, errors.New("Invoke: Return too much response.Responses. But In theory, It's not to be there")
	}
	return (response.Responses)[0], errors.New("Invoke: Return too little response.Responses")
}
