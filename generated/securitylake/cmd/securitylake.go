package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securitylake"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// securitylakeCmd represents the securitylake command
var _securitylakeCmd = &cobra.Command{
	Use:   "securitylake",
	Short: "AWS securitylake CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := securitylake.NewFromConfig(cfg)
		if _securitylakeCreateAwsLogSource {
			securitylake_CreateAwsLogSource(cfg, client)
			return
		}
		if _securitylakeCreateCustomLogSource {
			securitylake_CreateCustomLogSource(cfg, client)
			return
		}
		if _securitylakeCreateDataLake {
			securitylake_CreateDataLake(cfg, client)
			return
		}
		if _securitylakeCreateDataLakeExceptionSubscription {
			securitylake_CreateDataLakeExceptionSubscription(cfg, client)
			return
		}
		if _securitylakeCreateDataLakeOrganizationConfiguration {
			securitylake_CreateDataLakeOrganizationConfiguration(cfg, client)
			return
		}
		if _securitylakeCreateSubscriber {
			securitylake_CreateSubscriber(cfg, client)
			return
		}
		if _securitylakeCreateSubscriberNotification {
			securitylake_CreateSubscriberNotification(cfg, client)
			return
		}
		if _securitylakeDeleteAwsLogSource {
			securitylake_DeleteAwsLogSource(cfg, client)
			return
		}
		if _securitylakeDeleteCustomLogSource {
			securitylake_DeleteCustomLogSource(cfg, client)
			return
		}
		if _securitylakeDeleteDataLake {
			securitylake_DeleteDataLake(cfg, client)
			return
		}
		if _securitylakeDeleteDataLakeExceptionSubscription {
			securitylake_DeleteDataLakeExceptionSubscription(cfg, client)
			return
		}
		if _securitylakeDeleteDataLakeOrganizationConfiguration {
			securitylake_DeleteDataLakeOrganizationConfiguration(cfg, client)
			return
		}
		if _securitylakeDeleteSubscriber {
			securitylake_DeleteSubscriber(cfg, client)
			return
		}
		if _securitylakeDeleteSubscriberNotification {
			securitylake_DeleteSubscriberNotification(cfg, client)
			return
		}
		if _securitylakeDeregisterDataLakeDelegatedAdministrator {
			securitylake_DeregisterDataLakeDelegatedAdministrator(cfg, client)
			return
		}
		if _securitylakeGetDataLakeExceptionSubscription {
			securitylake_GetDataLakeExceptionSubscription(cfg, client)
			return
		}
		if _securitylakeGetDataLakeOrganizationConfiguration {
			securitylake_GetDataLakeOrganizationConfiguration(cfg, client)
			return
		}
		if _securitylakeGetDataLakeSources {
			securitylake_GetDataLakeSources(cfg, client)
			return
		}
		if _securitylakeGetSubscriber {
			securitylake_GetSubscriber(cfg, client)
			return
		}
		if _securitylakeListDataLakeExceptions {
			securitylake_ListDataLakeExceptions(cfg, client)
			return
		}
		if _securitylakeListDataLakes {
			securitylake_ListDataLakes(cfg, client)
			return
		}
		if _securitylakeListLogSources {
			securitylake_ListLogSources(cfg, client)
			return
		}
		if _securitylakeListSubscribers {
			securitylake_ListSubscribers(cfg, client)
			return
		}
		if _securitylakeListTagsForResource {
			securitylake_ListTagsForResource(cfg, client)
			return
		}
		if _securitylakeRegisterDataLakeDelegatedAdministrator {
			securitylake_RegisterDataLakeDelegatedAdministrator(cfg, client)
			return
		}
		if _securitylakeTagResource {
			securitylake_TagResource(cfg, client)
			return
		}
		if _securitylakeUntagResource {
			securitylake_UntagResource(cfg, client)
			return
		}
		if _securitylakeUpdateDataLake {
			securitylake_UpdateDataLake(cfg, client)
			return
		}
		if _securitylakeUpdateDataLakeExceptionSubscription {
			securitylake_UpdateDataLakeExceptionSubscription(cfg, client)
			return
		}
		if _securitylakeUpdateSubscriber {
			securitylake_UpdateSubscriber(cfg, client)
			return
		}
		if _securitylakeUpdateSubscriberNotification {
			securitylake_UpdateSubscriberNotification(cfg, client)
			return
		}

	},
}

var (
	_securitylakeCreateAwsLogSource                       bool
	_securitylakeCreateCustomLogSource                    bool
	_securitylakeCreateDataLake                           bool
	_securitylakeCreateDataLakeExceptionSubscription      bool
	_securitylakeCreateDataLakeOrganizationConfiguration  bool
	_securitylakeCreateSubscriber                         bool
	_securitylakeCreateSubscriberNotification             bool
	_securitylakeDeleteAwsLogSource                       bool
	_securitylakeDeleteCustomLogSource                    bool
	_securitylakeDeleteDataLake                           bool
	_securitylakeDeleteDataLakeExceptionSubscription      bool
	_securitylakeDeleteDataLakeOrganizationConfiguration  bool
	_securitylakeDeleteSubscriber                         bool
	_securitylakeDeleteSubscriberNotification             bool
	_securitylakeDeregisterDataLakeDelegatedAdministrator bool
	_securitylakeGetDataLakeExceptionSubscription         bool
	_securitylakeGetDataLakeOrganizationConfiguration     bool
	_securitylakeGetDataLakeSources                       bool
	_securitylakeGetSubscriber                            bool
	_securitylakeListDataLakeExceptions                   bool
	_securitylakeListDataLakes                            bool
	_securitylakeListLogSources                           bool
	_securitylakeListSubscribers                          bool
	_securitylakeListTagsForResource                      bool
	_securitylakeRegisterDataLakeDelegatedAdministrator   bool
	_securitylakeTagResource                              bool
	_securitylakeUntagResource                            bool
	_securitylakeUpdateDataLake                           bool
	_securitylakeUpdateDataLakeExceptionSubscription      bool
	_securitylakeUpdateSubscriber                         bool
	_securitylakeUpdateSubscriberNotification             bool

	_securitylakeAccessTypes             string
	_securitylakeAccountId               string
	_securitylakeAccounts                []string
	_securitylakeAutoEnableNewAccount    string
	_securitylakeConfiguration           string
	_securitylakeConfigurations          string
	_securitylakeEventClasses            []string
	_securitylakeExceptionTimeToLive     string
	_securitylakeMaxResults              string
	_securitylakeMetaStoreManagerRoleArn string
	_securitylakeNextToken               string
	_securitylakeNotificationEndpoint    string
	_securitylakeRegions                 []string
	_securitylakeResourceArn             string
	_securitylakeSourceName              string
	_securitylakeSourceVersion           string
	_securitylakeSources                 string
	_securitylakeSubscriberDescription   string
	_securitylakeSubscriberId            string
	_securitylakeSubscriberIdentity      string
	_securitylakeSubscriberName          string
	_securitylakeSubscriptionProtocol    string
	_securitylakeTagKeys                 []string
	_securitylakeTags                    string
)

// Adds a natively supported Amazon Web Services service as an Amazon Security
// Lake source. Enables source types for member accounts in required Amazon Web
// Services Regions, based on the parameters you specify. You can choose any source
// type in any Region for either accounts that are part of a trusted organization
// or standalone accounts. Once you add an Amazon Web Services service as a source,
// Security Lake starts collecting logs and events from it.
//
// You can use this API only to enable natively supported Amazon Web Services
// services as a source. Use CreateCustomLogSource to enable data collection from
// a custom source.
func securitylake_CreateAwsLogSource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateAwsLogSourceInput{
		// Sources: []types.AwsLogSourceConfiguration, // Required
	}

	if len(_securitylakeSources) > 0 {
		if err := assignInputField(input, "Sources", _securitylakeSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAwsLogSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a third-party custom source in Amazon Security Lake, from the Amazon Web
// Services Region where you want to create a custom source. Security Lake can
// collect logs and events from third-party custom sources. After creating the
// appropriate IAM role to invoke Glue crawler, use this API to add a custom source
// name in Security Lake. This operation creates a partition in the Amazon S3
// bucket for Security Lake as the target location for log files from the custom
// source. In addition, this operation also creates an associated Glue table and an
// Glue crawler.
func securitylake_CreateCustomLogSource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateCustomLogSourceInput{
		// Configuration: *types.CustomLogSourceConfiguration, // Required
		// SourceName: *string, // Required
	}

	if len(_securitylakeConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _securitylakeConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSourceName) > 0 {
		input.SourceName = aws.String(_securitylakeSourceName)
	}
	if len(_securitylakeEventClasses) > 0 {
		input.EventClasses = append([]string(nil), _securitylakeEventClasses...)
	}
	if len(_securitylakeSourceVersion) > 0 {
		input.SourceVersion = aws.String(_securitylakeSourceVersion)
	}

	if resp, err := client.CreateCustomLogSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initializes an Amazon Security Lake instance with the provided (or default)
// configuration. You can enable Security Lake in Amazon Web Services Regions with
// customized settings before enabling log collection in Regions. To specify
// particular Regions, configure these Regions using the configurations parameter.
// If you have already enabled Security Lake in a Region when you call this
// command, the command will update the Region if you provide new configuration
// parameters. If you have not already enabled Security Lake in the Region when you
// call this API, it will set up the data lake in the Region with the specified
// configurations.
//
// When you enable Security Lake, it starts ingesting security data after the
// CreateAwsLogSource call and after you create subscribers using the
// CreateSubscriber API. This includes ingesting security data from sources,
// storing data, and making data accessible to subscribers. Security Lake also
// enables all the existing settings and resources that it stores or maintains for
// your Amazon Web Services account in the current Region, including security log
// and event data. For more information, see the [Amazon Security Lake User Guide].
//
// [Amazon Security Lake User Guide]: https://docs.aws.amazon.com/security-lake/latest/userguide/what-is-security-lake.html
func securitylake_CreateDataLake(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateDataLakeInput{
		// Configurations: []types.DataLakeConfiguration, // Required
		// MetaStoreManagerRoleArn: *string, // Required
	}

	if len(_securitylakeConfigurations) > 0 {
		if err := assignInputField(input, "Configurations", _securitylakeConfigurations); err != nil {
			log.Errorf("invalid --configurations: %s", err.Error())
			return
		}
	}
	if len(_securitylakeMetaStoreManagerRoleArn) > 0 {
		input.MetaStoreManagerRoleArn = aws.String(_securitylakeMetaStoreManagerRoleArn)
	}
	if len(_securitylakeTags) > 0 {
		if err := assignInputField(input, "Tags", _securitylakeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataLake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates the specified notification subscription in Amazon Security Lake for the
// organization you specify. The notification subscription is created for
// exceptions that cannot be resolved by Security Lake automatically.
func securitylake_CreateDataLakeExceptionSubscription(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateDataLakeExceptionSubscriptionInput{
		// NotificationEndpoint: *string, // Required
		// SubscriptionProtocol: *string, // Required
	}

	if len(_securitylakeNotificationEndpoint) > 0 {
		input.NotificationEndpoint = aws.String(_securitylakeNotificationEndpoint)
	}
	if len(_securitylakeSubscriptionProtocol) > 0 {
		input.SubscriptionProtocol = aws.String(_securitylakeSubscriptionProtocol)
	}
	if len(_securitylakeExceptionTimeToLive) > 0 {
		if err := assignInputField(input, "ExceptionTimeToLive", _securitylakeExceptionTimeToLive); err != nil {
			log.Errorf("invalid --exception-time-to-live: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataLakeExceptionSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Automatically enables Amazon Security Lake for new member accounts in your
// organization. Security Lake is not automatically enabled for any existing member
// accounts in your organization.
//
// This operation merges the new data lake organization configuration with the
// existing configuration for Security Lake in your organization. If you want to
// create a new data lake organization configuration, you must delete the existing
// one using [DeleteDataLakeOrganizationConfiguration].
//
// [DeleteDataLakeOrganizationConfiguration]: https://docs.aws.amazon.com/security-lake/latest/APIReference/API_DeleteDataLakeOrganizationConfiguration.html
func securitylake_CreateDataLakeOrganizationConfiguration(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateDataLakeOrganizationConfigurationInput{}

	if len(_securitylakeAutoEnableNewAccount) > 0 {
		if err := assignInputField(input, "AutoEnableNewAccount", _securitylakeAutoEnableNewAccount); err != nil {
			log.Errorf("invalid --auto-enable-new-account: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataLakeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscriber for accounts that are already enabled in Amazon Security
// Lake. You can create a subscriber with access to data in the current Amazon Web
// Services Region.
func securitylake_CreateSubscriber(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateSubscriberInput{
		// Sources: []types.LogSourceResource, // Required
		// SubscriberIdentity: *types.AwsIdentity, // Required
		// SubscriberName: *string, // Required
	}

	if len(_securitylakeSources) > 0 {
		if err := assignInputField(input, "Sources", _securitylakeSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberIdentity) > 0 {
		if err := assignInputField(input, "SubscriberIdentity", _securitylakeSubscriberIdentity); err != nil {
			log.Errorf("invalid --subscriber-identity: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberName) > 0 {
		input.SubscriberName = aws.String(_securitylakeSubscriberName)
	}
	if len(_securitylakeAccessTypes) > 0 {
		if err := assignInputField(input, "AccessTypes", _securitylakeAccessTypes); err != nil {
			log.Errorf("invalid --access-types: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberDescription) > 0 {
		input.SubscriberDescription = aws.String(_securitylakeSubscriberDescription)
	}
	if len(_securitylakeTags) > 0 {
		if err := assignInputField(input, "Tags", _securitylakeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies the subscriber when new data is written to the data lake for the
// sources that the subscriber consumes in Security Lake. You can create only one
// subscriber notification per subscriber.
func securitylake_CreateSubscriberNotification(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.CreateSubscriberNotificationInput{
		// Configuration: types.NotificationConfiguration, // Required
		// SubscriberId: *string, // Required
	}

	if len(_securitylakeConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _securitylakeConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberId) > 0 {
		input.SubscriberId = aws.String(_securitylakeSubscriberId)
	}

	if resp, err := client.CreateSubscriberNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a natively supported Amazon Web Services service as an Amazon Security
// Lake source. You can remove a source for one or more Regions. When you remove
// the source, Security Lake stops collecting data from that source in the
// specified Regions and accounts, and subscribers can no longer consume new data
// from the source. However, subscribers can still consume data that Security Lake
// collected from the source before removal.
//
// You can choose any source type in any Amazon Web Services Region for either
// accounts that are part of a trusted organization or standalone accounts.
func securitylake_DeleteAwsLogSource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteAwsLogSourceInput{
		// Sources: []types.AwsLogSourceConfiguration, // Required
	}

	if len(_securitylakeSources) > 0 {
		if err := assignInputField(input, "Sources", _securitylakeSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAwsLogSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a custom log source from Amazon Security Lake, to stop sending data
// from the custom source to Security Lake.
func securitylake_DeleteCustomLogSource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteCustomLogSourceInput{
		// SourceName: *string, // Required
	}

	if len(_securitylakeSourceName) > 0 {
		input.SourceName = aws.String(_securitylakeSourceName)
	}
	if len(_securitylakeSourceVersion) > 0 {
		input.SourceVersion = aws.String(_securitylakeSourceVersion)
	}

	if resp, err := client.DeleteCustomLogSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you disable Amazon Security Lake from your account, Security Lake is
// disabled in all Amazon Web Services Regions and it stops collecting data from
// your sources. Also, this API automatically takes steps to remove the account
// from Security Lake. However, Security Lake retains all of your existing settings
// and the resources that it created in your Amazon Web Services account in the
// current Amazon Web Services Region.
//
// The DeleteDataLake operation does not delete the data that is stored in your
// Amazon S3 bucket, which is owned by your Amazon Web Services account. For more
// information, see the [Amazon Security Lake User Guide].
//
// [Amazon Security Lake User Guide]: https://docs.aws.amazon.com/security-lake/latest/userguide/disable-security-lake.html
func securitylake_DeleteDataLake(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteDataLakeInput{
		// Regions: []string, // Required
	}

	if len(_securitylakeRegions) > 0 {
		input.Regions = append([]string(nil), _securitylakeRegions...)
	}

	if resp, err := client.DeleteDataLake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified notification subscription in Amazon Security Lake for the
// organization you specify.
func securitylake_DeleteDataLakeExceptionSubscription(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteDataLakeExceptionSubscriptionInput{}

	if resp, err := client.DeleteDataLakeExceptionSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Turns off automatic enablement of Amazon Security Lake for member accounts that
// are added to an organization in Organizations. Only the delegated Security Lake
// administrator for an organization can perform this operation. If the delegated
// Security Lake administrator performs this operation, new member accounts won't
// automatically contribute data to the data lake.
func securitylake_DeleteDataLakeOrganizationConfiguration(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteDataLakeOrganizationConfigurationInput{}

	if len(_securitylakeAutoEnableNewAccount) > 0 {
		if err := assignInputField(input, "AutoEnableNewAccount", _securitylakeAutoEnableNewAccount); err != nil {
			log.Errorf("invalid --auto-enable-new-account: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDataLakeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the subscription permission and all notification settings for accounts
// that are already enabled in Amazon Security Lake. When you run DeleteSubscriber
// , the subscriber will no longer consume data from Security Lake and the
// subscriber is removed. This operation deletes the subscriber and removes access
// to data in the current Amazon Web Services Region.
func securitylake_DeleteSubscriber(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteSubscriberInput{
		// SubscriberId: *string, // Required
	}

	if len(_securitylakeSubscriberId) > 0 {
		input.SubscriberId = aws.String(_securitylakeSubscriberId)
	}

	if resp, err := client.DeleteSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified subscription notification in Amazon Security Lake for the
// organization you specify.
func securitylake_DeleteSubscriberNotification(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeleteSubscriberNotificationInput{
		// SubscriberId: *string, // Required
	}

	if len(_securitylakeSubscriberId) > 0 {
		input.SubscriberId = aws.String(_securitylakeSubscriberId)
	}

	if resp, err := client.DeleteSubscriberNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the Amazon Security Lake delegated administrator account for the
// organization. This API can only be called by the organization management
// account. The organization management account cannot be the delegated
// administrator account.
func securitylake_DeregisterDataLakeDelegatedAdministrator(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.DeregisterDataLakeDelegatedAdministratorInput{}

	if resp, err := client.DeregisterDataLakeDelegatedAdministrator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the protocol and endpoint that were provided when subscribing to
// Amazon SNS topics for exception notifications.
func securitylake_GetDataLakeExceptionSubscription(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.GetDataLakeExceptionSubscriptionInput{}

	if resp, err := client.GetDataLakeExceptionSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration that will be automatically set up for accounts
// added to the organization after the organization has onboarded to Amazon
// Security Lake. This API does not take input parameters.
func securitylake_GetDataLakeOrganizationConfiguration(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.GetDataLakeOrganizationConfigurationInput{}

	if resp, err := client.GetDataLakeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a snapshot of the current Region, including whether Amazon Security
// Lake is enabled for those accounts and which sources Security Lake is collecting
// data from.
func securitylake_GetDataLakeSources(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.GetDataLakeSourcesInput{}

	if len(_securitylakeAccounts) > 0 {
		input.Accounts = append([]string(nil), _securitylakeAccounts...)
	}
	if len(_securitylakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securitylakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securitylakeNextToken) > 0 {
		input.NextToken = aws.String(_securitylakeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetDataLakeSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securitylake.GetDataLakeSourcesOutput
	p := securitylake.NewGetDataLakeSourcesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the subscription information for the specified subscription ID. You
// can get information about a specific subscriber.
func securitylake_GetSubscriber(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.GetSubscriberInput{
		// SubscriberId: *string, // Required
	}

	if len(_securitylakeSubscriberId) > 0 {
		input.SubscriberId = aws.String(_securitylakeSubscriberId)
	}

	if resp, err := client.GetSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the Amazon Security Lake exceptions that you can use to find the source
// of problems and fix them.
func securitylake_ListDataLakeExceptions(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.ListDataLakeExceptionsInput{}

	if len(_securitylakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securitylakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securitylakeNextToken) > 0 {
		input.NextToken = aws.String(_securitylakeNextToken)
	}
	if len(_securitylakeRegions) > 0 {
		input.Regions = append([]string(nil), _securitylakeRegions...)
	}

	if disablePaginator() {
		if resp, err := client.ListDataLakeExceptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securitylake.ListDataLakeExceptionsOutput
	p := securitylake.NewListDataLakeExceptionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the Amazon Security Lake configuration object for the specified
// Amazon Web Services Regions. You can use this operation to determine whether
// Security Lake is enabled for a Region.
func securitylake_ListDataLakes(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.ListDataLakesInput{}

	if len(_securitylakeRegions) > 0 {
		input.Regions = append([]string(nil), _securitylakeRegions...)
	}

	if resp, err := client.ListDataLakes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the log sources.
func securitylake_ListLogSources(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.ListLogSourcesInput{}

	if len(_securitylakeAccounts) > 0 {
		input.Accounts = append([]string(nil), _securitylakeAccounts...)
	}
	if len(_securitylakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securitylakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securitylakeNextToken) > 0 {
		input.NextToken = aws.String(_securitylakeNextToken)
	}
	if len(_securitylakeRegions) > 0 {
		input.Regions = append([]string(nil), _securitylakeRegions...)
	}
	if len(_securitylakeSources) > 0 {
		if err := assignInputField(input, "Sources", _securitylakeSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLogSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securitylake.ListLogSourcesOutput
	p := securitylake.NewListLogSourcesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all subscribers for the specific Amazon Security Lake account ID. You can
// retrieve a list of subscriptions associated with a specific organization or
// Amazon Web Services account.
func securitylake_ListSubscribers(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.ListSubscribersInput{}

	if len(_securitylakeMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securitylakeMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securitylakeNextToken) > 0 {
		input.NextToken = aws.String(_securitylakeNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscribers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securitylake.ListSubscribersOutput
	p := securitylake.NewListSubscribersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the tags (keys and values) that are associated with an Amazon
// Security Lake resource: a subscriber, or the data lake configuration for your
// Amazon Web Services account in a particular Amazon Web Services Region.
func securitylake_ListTagsForResource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_securitylakeResourceArn) > 0 {
		input.ResourceArn = aws.String(_securitylakeResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates the Amazon Security Lake delegated administrator account for the
// organization. This API can only be called by the organization management
// account. The organization management account cannot be the delegated
// administrator account.
func securitylake_RegisterDataLakeDelegatedAdministrator(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.RegisterDataLakeDelegatedAdministratorInput{
		// AccountId: *string, // Required
	}

	if len(_securitylakeAccountId) > 0 {
		input.AccountId = aws.String(_securitylakeAccountId)
	}

	if resp, err := client.RegisterDataLakeDelegatedAdministrator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates one or more tags that are associated with an Amazon Security
// Lake resource: a subscriber, or the data lake configuration for your Amazon Web
// Services account in a particular Amazon Web Services Region. A tag is a label
// that you can define and associate with Amazon Web Services resources. Each tag
// consists of a required tag key and an associated tag value. A tag key is a
// general label that acts as a category for a more specific tag value. A tag value
// acts as a descriptor for a tag key. Tags can help you identify, categorize, and
// manage resources in different ways, such as by owner, environment, or other
// criteria. For more information, see [Tagging Amazon Security Lake resources]in the Amazon Security Lake User Guide.
//
// [Tagging Amazon Security Lake resources]: https://docs.aws.amazon.com/security-lake/latest/userguide/tagging-resources.html
func securitylake_TagResource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_securitylakeResourceArn) > 0 {
		input.ResourceArn = aws.String(_securitylakeResourceArn)
	}
	if len(_securitylakeTags) > 0 {
		if err := assignInputField(input, "Tags", _securitylakeTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags (keys and values) from an Amazon Security Lake
// resource: a subscriber, or the data lake configuration for your Amazon Web
// Services account in a particular Amazon Web Services Region.
func securitylake_UntagResource(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_securitylakeResourceArn) > 0 {
		input.ResourceArn = aws.String(_securitylakeResourceArn)
	}
	if len(_securitylakeTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _securitylakeTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use UpdateDataLake to specify where to store your security data, how it
// should be encrypted at rest and for how long. You can add a [Rollup Region]to consolidate data
// from multiple Amazon Web Services Regions, replace default encryption (SSE-S3)
// with [Customer Manged Key], or specify transition and expiration actions through storage [Lifecycle management]. The
// UpdateDataLake API works as an "upsert" operation that performs an insert if the
// specified item or record does not exist, or an update if it already exists.
// Security Lake securely stores your data at rest using Amazon Web Services
// encryption solutions. For more details, see [Data protection in Amazon Security Lake].
//
// For example, omitting the key encryptionConfiguration from a Region that is
// included in an update call that currently uses KMS will leave that Region's KMS
// key in place, but specifying encryptionConfiguration: {kmsKeyId:
// 'S3_MANAGED_KEY'} for that same Region will reset the key to S3-managed .
//
// For more details about lifecycle management and how to update retention
// settings for one or more Regions after enabling Security Lake, see the [Amazon Security Lake User Guide].
//
// [Lifecycle management]: https://docs.aws.amazon.com/security-lake/latest/userguide/lifecycle-management.html
// [Rollup Region]: https://docs.aws.amazon.com/security-lake/latest/userguide/manage-regions.html#add-rollup-region
// [Data protection in Amazon Security Lake]: https://docs.aws.amazon.com/security-lake/latest/userguide/data-protection.html
// [Amazon Security Lake User Guide]: https://docs.aws.amazon.com/security-lake/latest/userguide/lifecycle-management.html
// [Customer Manged Key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
func securitylake_UpdateDataLake(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.UpdateDataLakeInput{
		// Configurations: []types.DataLakeConfiguration, // Required
	}

	if len(_securitylakeConfigurations) > 0 {
		if err := assignInputField(input, "Configurations", _securitylakeConfigurations); err != nil {
			log.Errorf("invalid --configurations: %s", err.Error())
			return
		}
	}
	if len(_securitylakeMetaStoreManagerRoleArn) > 0 {
		input.MetaStoreManagerRoleArn = aws.String(_securitylakeMetaStoreManagerRoleArn)
	}

	if resp, err := client.UpdateDataLake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified notification subscription in Amazon Security Lake for the
// organization you specify.
func securitylake_UpdateDataLakeExceptionSubscription(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.UpdateDataLakeExceptionSubscriptionInput{
		// NotificationEndpoint: *string, // Required
		// SubscriptionProtocol: *string, // Required
	}

	if len(_securitylakeNotificationEndpoint) > 0 {
		input.NotificationEndpoint = aws.String(_securitylakeNotificationEndpoint)
	}
	if len(_securitylakeSubscriptionProtocol) > 0 {
		input.SubscriptionProtocol = aws.String(_securitylakeSubscriptionProtocol)
	}
	if len(_securitylakeExceptionTimeToLive) > 0 {
		if err := assignInputField(input, "ExceptionTimeToLive", _securitylakeExceptionTimeToLive); err != nil {
			log.Errorf("invalid --exception-time-to-live: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataLakeExceptionSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing subscription for the given Amazon Security Lake account ID.
// You can update a subscriber by changing the sources that the subscriber consumes
// data from.
func securitylake_UpdateSubscriber(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.UpdateSubscriberInput{
		// SubscriberId: *string, // Required
	}

	if len(_securitylakeSubscriberId) > 0 {
		input.SubscriberId = aws.String(_securitylakeSubscriberId)
	}
	if len(_securitylakeSources) > 0 {
		if err := assignInputField(input, "Sources", _securitylakeSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberDescription) > 0 {
		input.SubscriberDescription = aws.String(_securitylakeSubscriberDescription)
	}
	if len(_securitylakeSubscriberIdentity) > 0 {
		if err := assignInputField(input, "SubscriberIdentity", _securitylakeSubscriberIdentity); err != nil {
			log.Errorf("invalid --subscriber-identity: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberName) > 0 {
		input.SubscriberName = aws.String(_securitylakeSubscriberName)
	}

	if resp, err := client.UpdateSubscriber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing notification method for the subscription (SQS or HTTPs
// endpoint) or switches the notification subscription endpoint for a subscriber.
func securitylake_UpdateSubscriberNotification(cfg aws.Config, client *securitylake.Client) {
	input := &securitylake.UpdateSubscriberNotificationInput{
		// Configuration: types.NotificationConfiguration, // Required
		// SubscriberId: *string, // Required
	}

	if len(_securitylakeConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _securitylakeConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_securitylakeSubscriberId) > 0 {
		input.SubscriberId = aws.String(_securitylakeSubscriberId)
	}

	if resp, err := client.UpdateSubscriberNotification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_securitylakeCmd)
	_securitylakeCmd.Flags().SortFlags = false

	_securitylakeCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_securitylakeCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_securitylakeCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_securitylakeCmd.Flags().StringVarP(&_securitylakeAccessTypes, "access-types", "", "", "Access Types")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeAccountId, "account-id", "", "", "Account ID")
	_securitylakeCmd.Flags().StringSliceVarP(&_securitylakeAccounts, "accounts", "", nil, "Accounts")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeAutoEnableNewAccount, "auto-enable-new-account", "", "", "Auto Enable New Account")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeConfiguration, "configuration", "", "", "Configuration")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeConfigurations, "configurations", "", "", "Configurations")
	_securitylakeCmd.Flags().StringSliceVarP(&_securitylakeEventClasses, "event-classes", "", nil, "Event Classes")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeExceptionTimeToLive, "exception-time-to-live", "", "", "Exception Time To Live")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeMaxResults, "max-results", "", "", "Max Results")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeMetaStoreManagerRoleArn, "meta-store-manager-role-arn", "", "", "Meta Store Manager Role ARN")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeNextToken, "next-token", "", "", "Next Token")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeNotificationEndpoint, "notification-endpoint", "", "", "Notification Endpoint")
	_securitylakeCmd.Flags().StringSliceVarP(&_securitylakeRegions, "regions", "", nil, "Regions")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeResourceArn, "resource-arn", "", "", "Resource ARN")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSourceName, "source-name", "", "", "Source Name")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSourceVersion, "source-version", "", "", "Source Version")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSources, "sources", "", "", "Sources")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSubscriberDescription, "subscriber-description", "", "", "Subscriber Description")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSubscriberId, "subscriber-id", "", "", "Subscriber ID")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSubscriberIdentity, "subscriber-identity", "", "", "Subscriber Identity")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSubscriberName, "subscriber-name", "", "", "Subscriber Name")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeSubscriptionProtocol, "subscription-protocol", "", "", "Subscription Protocol")
	_securitylakeCmd.Flags().StringSliceVarP(&_securitylakeTagKeys, "tag-keys", "", nil, "Tag Keys")
	_securitylakeCmd.Flags().StringVarP(&_securitylakeTags, "tags", "", "", "Tags")

	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateAwsLogSource, "create-aws-log-source", "", false, "Create AWS Log Source")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateCustomLogSource, "create-custom-log-source", "", false, "Create Custom Log Source")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateDataLake, "create-data-lake", "", false, "Create Data Lake")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateDataLakeExceptionSubscription, "create-data-lake-exception-subscription", "", false, "Create Data Lake Exception Subscription")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateDataLakeOrganizationConfiguration, "create-data-lake-organization-configuration", "", false, "Create Data Lake Organization Configuration")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateSubscriber, "create-subscriber", "", false, "Create Subscriber")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeCreateSubscriberNotification, "create-subscriber-notification", "", false, "Create Subscriber Notification")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteAwsLogSource, "delete-aws-log-source", "", false, "Delete AWS Log Source")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteCustomLogSource, "delete-custom-log-source", "", false, "Delete Custom Log Source")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteDataLake, "delete-data-lake", "", false, "Delete Data Lake")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteDataLakeExceptionSubscription, "delete-data-lake-exception-subscription", "", false, "Delete Data Lake Exception Subscription")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteDataLakeOrganizationConfiguration, "delete-data-lake-organization-configuration", "", false, "Delete Data Lake Organization Configuration")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteSubscriber, "delete-subscriber", "", false, "Delete Subscriber")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeleteSubscriberNotification, "delete-subscriber-notification", "", false, "Delete Subscriber Notification")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeDeregisterDataLakeDelegatedAdministrator, "deregister-data-lake-delegated-administrator", "", false, "Deregister Data Lake Delegated Administrator")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeGetDataLakeExceptionSubscription, "get-data-lake-exception-subscription", "", false, "Get Data Lake Exception Subscription")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeGetDataLakeOrganizationConfiguration, "get-data-lake-organization-configuration", "", false, "Get Data Lake Organization Configuration")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeGetDataLakeSources, "get-data-lake-sources", "", false, "Get Data Lake Sources")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeGetSubscriber, "get-subscriber", "", false, "Get Subscriber")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeListDataLakeExceptions, "list-data-lake-exceptions", "", false, "List Data Lake Exceptions")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeListDataLakes, "list-data-lakes", "", false, "List Data Lakes")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeListLogSources, "list-log-sources", "", false, "List Log Sources")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeListSubscribers, "list-subscribers", "", false, "List Subscribers")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeRegisterDataLakeDelegatedAdministrator, "register-data-lake-delegated-administrator", "", false, "Register Data Lake Delegated Administrator")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeTagResource, "tag-resource", "", false, "Tag Resource")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeUntagResource, "untag-resource", "", false, "Untag Resource")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeUpdateDataLake, "update-data-lake", "", false, "Update Data Lake")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeUpdateDataLakeExceptionSubscription, "update-data-lake-exception-subscription", "", false, "Update Data Lake Exception Subscription")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeUpdateSubscriber, "update-subscriber", "", false, "Update Subscriber")
	_securitylakeCmd.Flags().BoolVarP(&_securitylakeUpdateSubscriberNotification, "update-subscriber-notification", "", false, "Update Subscriber Notification")

}
