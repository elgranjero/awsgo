package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ec2instanceconnectCmd represents the ec2instanceconnect command
var _ec2instanceconnectCmd = &cobra.Command{
	Use:   "ec2instanceconnect",
	Short: "AWS ec2instanceconnect CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ec2instanceconnect.NewFromConfig(cfg)
		if _ec2instanceconnectSendSerialConsoleSSHPublicKey {
			ec2instanceconnect_SendSerialConsoleSSHPublicKey(cfg, client)
			return
		}
		if _ec2instanceconnectSendSSHPublicKey {
			ec2instanceconnect_SendSSHPublicKey(cfg, client)
			return
		}

	},
}

var (
	_ec2instanceconnectSendSerialConsoleSSHPublicKey bool
	_ec2instanceconnectSendSSHPublicKey              bool

	_ec2instanceconnectAvailabilityZone string
	_ec2instanceconnectInstanceId       string
	_ec2instanceconnectInstanceOSUser   string
	_ec2instanceconnectSerialPort       string
	_ec2instanceconnectSSHPublicKey     string
)

// Pushes an SSH public key to the specified EC2 instance. The key remains for 60
// seconds, which gives you 60 seconds to establish a serial console connection to
// the instance using SSH. For more information, see [EC2 Serial Console]in the Amazon EC2 User Guide.
//
// [EC2 Serial Console]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-serial-console.html
func ec2instanceconnect_SendSerialConsoleSSHPublicKey(cfg aws.Config, client *ec2instanceconnect.Client) {
	input := &ec2instanceconnect.SendSerialConsoleSSHPublicKeyInput{
		// InstanceId: *string, // Required
		// SSHPublicKey: *string, // Required
	}

	if len(_ec2instanceconnectInstanceId) > 0 {
		input.InstanceId = aws.String(_ec2instanceconnectInstanceId)
	}
	if len(_ec2instanceconnectSSHPublicKey) > 0 {
		input.SSHPublicKey = aws.String(_ec2instanceconnectSSHPublicKey)
	}
	if len(_ec2instanceconnectSerialPort) > 0 {
		if err := assignInputField(input, "SerialPort", _ec2instanceconnectSerialPort); err != nil {
			log.Errorf("invalid --serial-port: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendSerialConsoleSSHPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Pushes an SSH public key to the specified EC2 instance for use by the specified
// user. The key remains for 60 seconds. For more information, see [Connect to your Linux instance using EC2 Instance Connect]in the Amazon
// EC2 User Guide.
//
// [Connect to your Linux instance using EC2 Instance Connect]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Connect-using-EC2-Instance-Connect.html
func ec2instanceconnect_SendSSHPublicKey(cfg aws.Config, client *ec2instanceconnect.Client) {
	input := &ec2instanceconnect.SendSSHPublicKeyInput{
		// InstanceId: *string, // Required
		// InstanceOSUser: *string, // Required
		// SSHPublicKey: *string, // Required
	}

	if len(_ec2instanceconnectInstanceId) > 0 {
		input.InstanceId = aws.String(_ec2instanceconnectInstanceId)
	}
	if len(_ec2instanceconnectInstanceOSUser) > 0 {
		input.InstanceOSUser = aws.String(_ec2instanceconnectInstanceOSUser)
	}
	if len(_ec2instanceconnectSSHPublicKey) > 0 {
		input.SSHPublicKey = aws.String(_ec2instanceconnectSSHPublicKey)
	}
	if len(_ec2instanceconnectAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_ec2instanceconnectAvailabilityZone)
	}

	if resp, err := client.SendSSHPublicKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ec2instanceconnectCmd)
	_ec2instanceconnectCmd.Flags().SortFlags = false

	_ec2instanceconnectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ec2instanceconnectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ec2instanceconnectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ec2instanceconnectCmd.Flags().StringVarP(&_ec2instanceconnectAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_ec2instanceconnectCmd.Flags().StringVarP(&_ec2instanceconnectInstanceId, "instance-id", "", "", "Instance ID")
	_ec2instanceconnectCmd.Flags().StringVarP(&_ec2instanceconnectInstanceOSUser, "instance-os-user", "", "", "Instance OS User")
	_ec2instanceconnectCmd.Flags().StringVarP(&_ec2instanceconnectSerialPort, "serial-port", "", "", "Serial Port")
	_ec2instanceconnectCmd.Flags().StringVarP(&_ec2instanceconnectSSHPublicKey, "ssh-public-key", "", "", "SSH Public Key")

	_ec2instanceconnectCmd.Flags().BoolVarP(&_ec2instanceconnectSendSerialConsoleSSHPublicKey, "send-serial-console-ssh-public-key", "", false, "Send Serial Console SSH Public Key")
	_ec2instanceconnectCmd.Flags().BoolVarP(&_ec2instanceconnectSendSSHPublicKey, "send-ssh-public-key", "", false, "Send SSH Public Key")

}
