package fabric_sdk

import (
	"errors"

	"ihtPrivateSDK/share/logging"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
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
	logging.Info("TransactionID: %s", response.TransactionID)
	logging.Info("TxValidationCode: %d", response.TxValidationCode)

	// Notifier(client, getEventID())
	return &response, nil
}

// DisposalQuery Get accurate data by processing
func DisposalQuery(response *channel.Response) (*fab.TransactionProposalResponse, error) {
	if response.ChaincodeStatus == 200 && len(response.Responses) == 1 {
		return (response.Responses)[0], nil
	}
	if response.ChaincodeStatus == 200 && len(response.Responses) > 1 {
		return nil, errors.New("Query: Return too much response.Responses. But In theory, It's not to be there")
	}
	return (response.Responses)[0], errors.New("Query: Return too little response.Responses")
}
