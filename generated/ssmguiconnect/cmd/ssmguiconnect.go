package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssmguiconnect"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ssmguiconnectCmd represents the ssmguiconnect command
var _ssmguiconnectCmd = &cobra.Command{
	Use:   "ssmguiconnect",
	Short: "AWS ssmguiconnect CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ssmguiconnect.NewFromConfig(cfg)
		if _ssmguiconnectDeleteConnectionRecordingPreferences {
			ssmguiconnect_DeleteConnectionRecordingPreferences(cfg, client)
			return
		}
		if _ssmguiconnectGetConnectionRecordingPreferences {
			ssmguiconnect_GetConnectionRecordingPreferences(cfg, client)
			return
		}
		if _ssmguiconnectUpdateConnectionRecordingPreferences {
			ssmguiconnect_UpdateConnectionRecordingPreferences(cfg, client)
			return
		}

	},
}

var (
	_ssmguiconnectDeleteConnectionRecordingPreferences bool
	_ssmguiconnectGetConnectionRecordingPreferences    bool
	_ssmguiconnectUpdateConnectionRecordingPreferences bool

	_ssmguiconnectClientToken                    string
	_ssmguiconnectConnectionRecordingPreferences string
)

// Deletes the preferences for recording RDP connections.
func ssmguiconnect_DeleteConnectionRecordingPreferences(cfg aws.Config, client *ssmguiconnect.Client) {
	input := &ssmguiconnect.DeleteConnectionRecordingPreferencesInput{}

	if len(_ssmguiconnectClientToken) > 0 {
		input.ClientToken = aws.String(_ssmguiconnectClientToken)
	}

	if resp, err := client.DeleteConnectionRecordingPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the preferences specified for recording RDP connections in the
// requesting Amazon Web Services account and Amazon Web Services Region.
func ssmguiconnect_GetConnectionRecordingPreferences(cfg aws.Config, client *ssmguiconnect.Client) {
	input := &ssmguiconnect.GetConnectionRecordingPreferencesInput{}

	if resp, err := client.GetConnectionRecordingPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the preferences for recording RDP connections.
func ssmguiconnect_UpdateConnectionRecordingPreferences(cfg aws.Config, client *ssmguiconnect.Client) {
	input := &ssmguiconnect.UpdateConnectionRecordingPreferencesInput{
		// ConnectionRecordingPreferences: *types.ConnectionRecordingPreferences, // Required
	}

	if len(_ssmguiconnectConnectionRecordingPreferences) > 0 {
		if err := assignInputField(input, "ConnectionRecordingPreferences", _ssmguiconnectConnectionRecordingPreferences); err != nil {
			log.Errorf("invalid --connection-recording-preferences: %s", err.Error())
			return
		}
	}
	if len(_ssmguiconnectClientToken) > 0 {
		input.ClientToken = aws.String(_ssmguiconnectClientToken)
	}

	if resp, err := client.UpdateConnectionRecordingPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ssmguiconnectCmd)
	_ssmguiconnectCmd.Flags().SortFlags = false

	_ssmguiconnectCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ssmguiconnectCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ssmguiconnectCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ssmguiconnectCmd.Flags().StringVarP(&_ssmguiconnectClientToken, "client-token", "", "", "Client Token")
	_ssmguiconnectCmd.Flags().StringVarP(&_ssmguiconnectConnectionRecordingPreferences, "connection-recording-preferences", "", "", "Connection Recording Preferences")

	_ssmguiconnectCmd.Flags().BoolVarP(&_ssmguiconnectDeleteConnectionRecordingPreferences, "delete-connection-recording-preferences", "", false, "Delete Connection Recording Preferences")
	_ssmguiconnectCmd.Flags().BoolVarP(&_ssmguiconnectGetConnectionRecordingPreferences, "get-connection-recording-preferences", "", false, "Get Connection Recording Preferences")
	_ssmguiconnectCmd.Flags().BoolVarP(&_ssmguiconnectUpdateConnectionRecordingPreferences, "update-connection-recording-preferences", "", false, "Update Connection Recording Preferences")

}
