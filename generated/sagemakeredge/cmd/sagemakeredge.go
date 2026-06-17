package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sagemakeredge"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sagemakeredgeCmd represents the sagemakeredge command
var _sagemakeredgeCmd = &cobra.Command{
	Use:   "sagemakeredge",
	Short: "AWS sagemakeredge CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sagemakeredge.NewFromConfig(cfg)
		if _sagemakeredgeGetDeployments {
			sagemakeredge_GetDeployments(cfg, client)
			return
		}
		if _sagemakeredgeGetDeviceRegistration {
			sagemakeredge_GetDeviceRegistration(cfg, client)
			return
		}
		if _sagemakeredgeSendHeartbeat {
			sagemakeredge_SendHeartbeat(cfg, client)
			return
		}

	},
}

var (
	_sagemakeredgeGetDeployments        bool
	_sagemakeredgeGetDeviceRegistration bool
	_sagemakeredgeSendHeartbeat         bool

	_sagemakeredgeAgentMetrics     string
	_sagemakeredgeAgentVersion     string
	_sagemakeredgeDeploymentResult string
	_sagemakeredgeDeviceFleetName  string
	_sagemakeredgeDeviceName       string
	_sagemakeredgeModels           string
)

// Use to get the active deployments from a device.
func sagemakeredge_GetDeployments(cfg aws.Config, client *sagemakeredge.Client) {
	input := &sagemakeredge.GetDeploymentsInput{
		// DeviceFleetName: *string, // Required
		// DeviceName: *string, // Required
	}

	if len(_sagemakeredgeDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakeredgeDeviceFleetName)
	}
	if len(_sagemakeredgeDeviceName) > 0 {
		input.DeviceName = aws.String(_sagemakeredgeDeviceName)
	}

	if resp, err := client.GetDeployments(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to check if a device is registered with SageMaker Edge Manager.
func sagemakeredge_GetDeviceRegistration(cfg aws.Config, client *sagemakeredge.Client) {
	input := &sagemakeredge.GetDeviceRegistrationInput{
		// DeviceFleetName: *string, // Required
		// DeviceName: *string, // Required
	}

	if len(_sagemakeredgeDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakeredgeDeviceFleetName)
	}
	if len(_sagemakeredgeDeviceName) > 0 {
		input.DeviceName = aws.String(_sagemakeredgeDeviceName)
	}

	if resp, err := client.GetDeviceRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use to get the current status of devices registered on SageMaker Edge Manager.
func sagemakeredge_SendHeartbeat(cfg aws.Config, client *sagemakeredge.Client) {
	input := &sagemakeredge.SendHeartbeatInput{
		// AgentVersion: *string, // Required
		// DeviceFleetName: *string, // Required
		// DeviceName: *string, // Required
	}

	if len(_sagemakeredgeAgentVersion) > 0 {
		input.AgentVersion = aws.String(_sagemakeredgeAgentVersion)
	}
	if len(_sagemakeredgeDeviceFleetName) > 0 {
		input.DeviceFleetName = aws.String(_sagemakeredgeDeviceFleetName)
	}
	if len(_sagemakeredgeDeviceName) > 0 {
		input.DeviceName = aws.String(_sagemakeredgeDeviceName)
	}
	if len(_sagemakeredgeAgentMetrics) > 0 {
		if err := assignInputField(input, "AgentMetrics", _sagemakeredgeAgentMetrics); err != nil {
			log.Errorf("invalid --agent-metrics: %s", err.Error())
			return
		}
	}
	if len(_sagemakeredgeDeploymentResult) > 0 {
		if err := assignInputField(input, "DeploymentResult", _sagemakeredgeDeploymentResult); err != nil {
			log.Errorf("invalid --deployment-result: %s", err.Error())
			return
		}
	}
	if len(_sagemakeredgeModels) > 0 {
		if err := assignInputField(input, "Models", _sagemakeredgeModels); err != nil {
			log.Errorf("invalid --models: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendHeartbeat(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sagemakeredgeCmd)
	_sagemakeredgeCmd.Flags().SortFlags = false

	_sagemakeredgeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_sagemakeredgeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sagemakeredgeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_sagemakeredgeCmd.Flags().StringVarP(&_sagemakeredgeAgentMetrics, "agent-metrics", "", "", "Agent Metrics")
	_sagemakeredgeCmd.Flags().StringVarP(&_sagemakeredgeAgentVersion, "agent-version", "", "", "Agent Version")
	_sagemakeredgeCmd.Flags().StringVarP(&_sagemakeredgeDeploymentResult, "deployment-result", "", "", "Deployment Result")
	_sagemakeredgeCmd.Flags().StringVarP(&_sagemakeredgeDeviceFleetName, "device-fleet-name", "", "", "Device Fleet Name")
	_sagemakeredgeCmd.Flags().StringVarP(&_sagemakeredgeDeviceName, "device-name", "", "", "Device Name")
	_sagemakeredgeCmd.Flags().StringVarP(&_sagemakeredgeModels, "models", "", "", "Models")

	_sagemakeredgeCmd.Flags().BoolVarP(&_sagemakeredgeGetDeployments, "get-deployments", "", false, "Get Deployments")
	_sagemakeredgeCmd.Flags().BoolVarP(&_sagemakeredgeGetDeviceRegistration, "get-device-registration", "", false, "Get Device Registration")
	_sagemakeredgeCmd.Flags().BoolVarP(&_sagemakeredgeSendHeartbeat, "send-heartbeat", "", false, "Send Heartbeat")

}
