package fabric_sdk

import (
	"ihtPrivateSDK/share/lib"
	"ihtPrivateSDK/share/logging"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	// "github.com/hyperledger/fabric-sdk-go/test/integration"
)

const (
	channelID      = "mychannel"
	orgName        = "org1"
	orgAdmin       = "Admin"
	ordererOrgName = "peerOrg1"
	// ordererOrgName = "Org1MSP"
	chaincodeID = "app"

	configPath = "./conf.d/config_test.yaml"
)

var (
	client *channel.Client
)

func createClient(sdkOpts ...fabsdk.Option) *channel.Client {
	if client != nil {
		return client
	}
	if !lib.IsFileExist(configPath) {
		logging.Error("The fabric-go-sdk config file is not exist")
		return nil
	}

	configOpt := config.FromFile(configPath)

	sdk, err := fabsdk.New(configOpt, sdkOpts...)
	if err != nil {
		logging.Error("Failed to create new SDK: %v", err)
		return nil
	}
	// defer sdk.Close()

	//prepare channel client context using client context
	context := sdk.ChannelContext(channelID, fabsdk.WithUser(orgAdmin), fabsdk.WithOrg(orgName))

	client, err := channel.New(context)
	if err != nil {
		logging.Error("Failed to create new channel client: %v", err)
		return nil
	}
	return client
}

// Notifier orderer event callback
func Notifier(client *channel.Client, eventID string) {
	// Register chaincode event (pass in channel which receives event details when the event is complete)
	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		logging.Debug("Failed to register chaincode event: %s", err)
	}
	defer client.UnregisterChaincodeEvent(reg)

	select {
	case ccEvent := <-notifier:
		logging.Debug("Received chaincode event: %s", ccEvent)
	case <-time.After(time.Second * 6):
		logging.Debug("Did't receive chaincode event for eventId(%s)", eventID)
	}
}

func getEventID() string {
	// eventID := integration.GenerateRandomID()
	return "test([a-zA-Z]+)"
}
