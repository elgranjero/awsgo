package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/appconfigdata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// appconfigdataCmd represents the appconfigdata command
var _appconfigdataCmd = &cobra.Command{
	Use:   "appconfigdata",
	Short: "AWS appconfigdata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := appconfigdata.NewFromConfig(cfg)
		if _appconfigdataGetLatestConfiguration {
			appconfigdata_GetLatestConfiguration(cfg, client)
			return
		}
		if _appconfigdataStartConfigurationSession {
			appconfigdata_StartConfigurationSession(cfg, client)
			return
		}

	},
}

var (
	_appconfigdataGetLatestConfiguration    bool
	_appconfigdataStartConfigurationSession bool

	_appconfigdataApplicationIdentifier                string
	_appconfigdataConfigurationProfileIdentifier       string
	_appconfigdataConfigurationToken                   string
	_appconfigdataEnvironmentIdentifier                string
	_appconfigdataRequiredMinimumPollIntervalInSeconds string
)

// Retrieves the latest deployed configuration. This API may return empty
// configuration data if the client already has the latest version. For more
// information about this API action and to view example CLI commands that show how
// to use it with the StartConfigurationSessionAPI action, see [Retrieving the configuration] in the AppConfig User Guide.
//
// Note the following important information.
//
// - Each configuration token is only valid for one call to
// GetLatestConfiguration . The GetLatestConfiguration response includes a
// NextPollConfigurationToken that should always replace the token used for the
// just-completed call in preparation for the next one.
//
// - GetLatestConfiguration is a priced call. For more information, see [Pricing].
//
// [Pricing]: https://aws.amazon.com/systems-manager/pricing/
// [Retrieving the configuration]: http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration
func appconfigdata_GetLatestConfiguration(cfg aws.Config, client *appconfigdata.Client) {
	input := &appconfigdata.GetLatestConfigurationInput{
		// ConfigurationToken: *string, // Required
	}

	if len(_appconfigdataConfigurationToken) > 0 {
		input.ConfigurationToken = aws.String(_appconfigdataConfigurationToken)
	}

	if resp, err := client.GetLatestConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a configuration session used to retrieve a deployed configuration. For
// more information about this API action and to view example CLI commands that
// show how to use it with the GetLatestConfigurationAPI action, see [Retrieving the configuration] in the AppConfig User Guide.
//
// [Retrieving the configuration]: http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration
func appconfigdata_StartConfigurationSession(cfg aws.Config, client *appconfigdata.Client) {
	input := &appconfigdata.StartConfigurationSessionInput{
		// ApplicationIdentifier: *string, // Required
		// ConfigurationProfileIdentifier: *string, // Required
		// EnvironmentIdentifier: *string, // Required
	}

	if len(_appconfigdataApplicationIdentifier) > 0 {
		input.ApplicationIdentifier = aws.String(_appconfigdataApplicationIdentifier)
	}
	if len(_appconfigdataConfigurationProfileIdentifier) > 0 {
		input.ConfigurationProfileIdentifier = aws.String(_appconfigdataConfigurationProfileIdentifier)
	}
	if len(_appconfigdataEnvironmentIdentifier) > 0 {
		input.EnvironmentIdentifier = aws.String(_appconfigdataEnvironmentIdentifier)
	}
	if len(_appconfigdataRequiredMinimumPollIntervalInSeconds) > 0 {
		if err := assignInputField(input, "RequiredMinimumPollIntervalInSeconds", _appconfigdataRequiredMinimumPollIntervalInSeconds); err != nil {
			log.Errorf("invalid --required-minimum-poll-interval-in-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartConfigurationSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_appconfigdataCmd)
	_appconfigdataCmd.Flags().SortFlags = false

	_appconfigdataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_appconfigdataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_appconfigdataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_appconfigdataCmd.Flags().StringVarP(&_appconfigdataApplicationIdentifier, "application-identifier", "", "", "Application Identifier")
	_appconfigdataCmd.Flags().StringVarP(&_appconfigdataConfigurationProfileIdentifier, "configuration-profile-identifier", "", "", "Configuration Profile Identifier")
	_appconfigdataCmd.Flags().StringVarP(&_appconfigdataConfigurationToken, "configuration-token", "", "", "Configuration Token")
	_appconfigdataCmd.Flags().StringVarP(&_appconfigdataEnvironmentIdentifier, "environment-identifier", "", "", "Environment Identifier")
	_appconfigdataCmd.Flags().StringVarP(&_appconfigdataRequiredMinimumPollIntervalInSeconds, "required-minimum-poll-interval-in-seconds", "", "", "Required Minimum Poll Interval In Seconds")

	_appconfigdataCmd.Flags().BoolVarP(&_appconfigdataGetLatestConfiguration, "get-latest-configuration", "", false, "Get Latest Configuration")
	_appconfigdataCmd.Flags().BoolVarP(&_appconfigdataStartConfigurationSession, "start-configuration-session", "", false, "Start Configuration Session")

}
