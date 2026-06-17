package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// apigatewaymanagementapiCmd represents the apigatewaymanagementapi command
var _apigatewaymanagementapiCmd = &cobra.Command{
	Use:   "apigatewaymanagementapi",
	Short: "AWS apigatewaymanagementapi CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := apigatewaymanagementapi.NewFromConfig(cfg)
		if _apigatewaymanagementapiDeleteConnection {
			apigatewaymanagementapi_DeleteConnection(cfg, client)
			return
		}
		if _apigatewaymanagementapiGetConnection {
			apigatewaymanagementapi_GetConnection(cfg, client)
			return
		}
		if _apigatewaymanagementapiPostToConnection {
			apigatewaymanagementapi_PostToConnection(cfg, client)
			return
		}

	},
}

var (
	_apigatewaymanagementapiDeleteConnection bool
	_apigatewaymanagementapiGetConnection    bool
	_apigatewaymanagementapiPostToConnection bool

	_apigatewaymanagementapiConnectionId string
	_apigatewaymanagementapiData         string
)

// Delete the connection with the provided id.
func apigatewaymanagementapi_DeleteConnection(cfg aws.Config, client *apigatewaymanagementapi.Client) {
	input := &apigatewaymanagementapi.DeleteConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_apigatewaymanagementapiConnectionId) > 0 {
		input.ConnectionId = aws.String(_apigatewaymanagementapiConnectionId)
	}

	if resp, err := client.DeleteConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about the connection with the provided id.
func apigatewaymanagementapi_GetConnection(cfg aws.Config, client *apigatewaymanagementapi.Client) {
	input := &apigatewaymanagementapi.GetConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_apigatewaymanagementapiConnectionId) > 0 {
		input.ConnectionId = aws.String(_apigatewaymanagementapiConnectionId)
	}

	if resp, err := client.GetConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends the provided data to the specified connection.
func apigatewaymanagementapi_PostToConnection(cfg aws.Config, client *apigatewaymanagementapi.Client) {
	input := &apigatewaymanagementapi.PostToConnectionInput{
		// ConnectionId: *string, // Required
		// Data: []byte, // Required
	}

	if len(_apigatewaymanagementapiConnectionId) > 0 {
		input.ConnectionId = aws.String(_apigatewaymanagementapiConnectionId)
	}
	if len(_apigatewaymanagementapiData) > 0 {
		if err := assignInputField(input, "Data", _apigatewaymanagementapiData); err != nil {
			log.Errorf("invalid --data: %s", err.Error())
			return
		}
	}

	if resp, err := client.PostToConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_apigatewaymanagementapiCmd)
	_apigatewaymanagementapiCmd.Flags().SortFlags = false

	_apigatewaymanagementapiCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_apigatewaymanagementapiCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_apigatewaymanagementapiCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_apigatewaymanagementapiCmd.Flags().StringVarP(&_apigatewaymanagementapiConnectionId, "connection-id", "", "", "Connection ID")
	_apigatewaymanagementapiCmd.Flags().StringVarP(&_apigatewaymanagementapiData, "data", "", "", "Data")

	_apigatewaymanagementapiCmd.Flags().BoolVarP(&_apigatewaymanagementapiDeleteConnection, "delete-connection", "", false, "Delete Connection")
	_apigatewaymanagementapiCmd.Flags().BoolVarP(&_apigatewaymanagementapiGetConnection, "get-connection", "", false, "Get Connection")
	_apigatewaymanagementapiCmd.Flags().BoolVarP(&_apigatewaymanagementapiPostToConnection, "post-to-connection", "", false, "Post To Connection")

}
