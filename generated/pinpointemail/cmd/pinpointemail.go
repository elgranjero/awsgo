package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointemail"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pinpointemailCmd represents the pinpointemail command
var _pinpointemailCmd = &cobra.Command{
	Use:   "pinpointemail",
	Short: "AWS pinpointemail CLI",
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
		client := pinpointemail.NewFromConfig(cfg)
		if _pinpointemailCreateConfigurationSet {
			pinpointemail_CreateConfigurationSet(cfg, client)
			return
		}
		if _pinpointemailCreateConfigurationSetEventDestination {
			pinpointemail_CreateConfigurationSetEventDestination(cfg, client)
			return
		}
		if _pinpointemailCreateDedicatedIpPool {
			pinpointemail_CreateDedicatedIpPool(cfg, client)
			return
		}
		if _pinpointemailCreateDeliverabilityTestReport {
			pinpointemail_CreateDeliverabilityTestReport(cfg, client)
			return
		}
		if _pinpointemailCreateEmailIdentity {
			pinpointemail_CreateEmailIdentity(cfg, client)
			return
		}
		if _pinpointemailDeleteConfigurationSet {
			pinpointemail_DeleteConfigurationSet(cfg, client)
			return
		}
		if _pinpointemailDeleteConfigurationSetEventDestination {
			pinpointemail_DeleteConfigurationSetEventDestination(cfg, client)
			return
		}
		if _pinpointemailDeleteDedicatedIpPool {
			pinpointemail_DeleteDedicatedIpPool(cfg, client)
			return
		}
		if _pinpointemailDeleteEmailIdentity {
			pinpointemail_DeleteEmailIdentity(cfg, client)
			return
		}
		if _pinpointemailGetAccount {
			pinpointemail_GetAccount(cfg, client)
			return
		}
		if _pinpointemailGetBlacklistReports {
			pinpointemail_GetBlacklistReports(cfg, client)
			return
		}
		if _pinpointemailGetConfigurationSet {
			pinpointemail_GetConfigurationSet(cfg, client)
			return
		}
		if _pinpointemailGetConfigurationSetEventDestinations {
			pinpointemail_GetConfigurationSetEventDestinations(cfg, client)
			return
		}
		if _pinpointemailGetDedicatedIp {
			pinpointemail_GetDedicatedIp(cfg, client)
			return
		}
		if _pinpointemailGetDedicatedIps {
			pinpointemail_GetDedicatedIps(cfg, client)
			return
		}
		if _pinpointemailGetDeliverabilityDashboardOptions {
			pinpointemail_GetDeliverabilityDashboardOptions(cfg, client)
			return
		}
		if _pinpointemailGetDeliverabilityTestReport {
			pinpointemail_GetDeliverabilityTestReport(cfg, client)
			return
		}
		if _pinpointemailGetDomainDeliverabilityCampaign {
			pinpointemail_GetDomainDeliverabilityCampaign(cfg, client)
			return
		}
		if _pinpointemailGetDomainStatisticsReport {
			pinpointemail_GetDomainStatisticsReport(cfg, client)
			return
		}
		if _pinpointemailGetEmailIdentity {
			pinpointemail_GetEmailIdentity(cfg, client)
			return
		}
		if _pinpointemailListConfigurationSets {
			pinpointemail_ListConfigurationSets(cfg, client)
			return
		}
		if _pinpointemailListDedicatedIpPools {
			pinpointemail_ListDedicatedIpPools(cfg, client)
			return
		}
		if _pinpointemailListDeliverabilityTestReports {
			pinpointemail_ListDeliverabilityTestReports(cfg, client)
			return
		}
		if _pinpointemailListDomainDeliverabilityCampaigns {
			pinpointemail_ListDomainDeliverabilityCampaigns(cfg, client)
			return
		}
		if _pinpointemailListEmailIdentities {
			pinpointemail_ListEmailIdentities(cfg, client)
			return
		}
		if _pinpointemailListTagsForResource {
			pinpointemail_ListTagsForResource(cfg, client)
			return
		}
		if _pinpointemailPutAccountDedicatedIpWarmupAttributes {
			pinpointemail_PutAccountDedicatedIpWarmupAttributes(cfg, client)
			return
		}
		if _pinpointemailPutAccountSendingAttributes {
			pinpointemail_PutAccountSendingAttributes(cfg, client)
			return
		}
		if _pinpointemailPutConfigurationSetDeliveryOptions {
			pinpointemail_PutConfigurationSetDeliveryOptions(cfg, client)
			return
		}
		if _pinpointemailPutConfigurationSetReputationOptions {
			pinpointemail_PutConfigurationSetReputationOptions(cfg, client)
			return
		}
		if _pinpointemailPutConfigurationSetSendingOptions {
			pinpointemail_PutConfigurationSetSendingOptions(cfg, client)
			return
		}
		if _pinpointemailPutConfigurationSetTrackingOptions {
			pinpointemail_PutConfigurationSetTrackingOptions(cfg, client)
			return
		}
		if _pinpointemailPutDedicatedIpInPool {
			pinpointemail_PutDedicatedIpInPool(cfg, client)
			return
		}
		if _pinpointemailPutDedicatedIpWarmupAttributes {
			pinpointemail_PutDedicatedIpWarmupAttributes(cfg, client)
			return
		}
		if _pinpointemailPutDeliverabilityDashboardOption {
			pinpointemail_PutDeliverabilityDashboardOption(cfg, client)
			return
		}
		if _pinpointemailPutEmailIdentityDkimAttributes {
			pinpointemail_PutEmailIdentityDkimAttributes(cfg, client)
			return
		}
		if _pinpointemailPutEmailIdentityFeedbackAttributes {
			pinpointemail_PutEmailIdentityFeedbackAttributes(cfg, client)
			return
		}
		if _pinpointemailPutEmailIdentityMailFromAttributes {
			pinpointemail_PutEmailIdentityMailFromAttributes(cfg, client)
			return
		}
		if _pinpointemailSendEmail {
			pinpointemail_SendEmail(cfg, client)
			return
		}
		if _pinpointemailTagResource {
			pinpointemail_TagResource(cfg, client)
			return
		}
		if _pinpointemailUntagResource {
			pinpointemail_UntagResource(cfg, client)
			return
		}
		if _pinpointemailUpdateConfigurationSetEventDestination {
			pinpointemail_UpdateConfigurationSetEventDestination(cfg, client)
			return
		}

	},
}

var (
	_pinpointemailCreateConfigurationSet                 bool
	_pinpointemailCreateConfigurationSetEventDestination bool
	_pinpointemailCreateDedicatedIpPool                  bool
	_pinpointemailCreateDeliverabilityTestReport         bool
	_pinpointemailCreateEmailIdentity                    bool
	_pinpointemailDeleteConfigurationSet                 bool
	_pinpointemailDeleteConfigurationSetEventDestination bool
	_pinpointemailDeleteDedicatedIpPool                  bool
	_pinpointemailDeleteEmailIdentity                    bool
	_pinpointemailGetAccount                             bool
	_pinpointemailGetBlacklistReports                    bool
	_pinpointemailGetConfigurationSet                    bool
	_pinpointemailGetConfigurationSetEventDestinations   bool
	_pinpointemailGetDedicatedIp                         bool
	_pinpointemailGetDedicatedIps                        bool
	_pinpointemailGetDeliverabilityDashboardOptions      bool
	_pinpointemailGetDeliverabilityTestReport            bool
	_pinpointemailGetDomainDeliverabilityCampaign        bool
	_pinpointemailGetDomainStatisticsReport              bool
	_pinpointemailGetEmailIdentity                       bool
	_pinpointemailListConfigurationSets                  bool
	_pinpointemailListDedicatedIpPools                   bool
	_pinpointemailListDeliverabilityTestReports          bool
	_pinpointemailListDomainDeliverabilityCampaigns      bool
	_pinpointemailListEmailIdentities                    bool
	_pinpointemailListTagsForResource                    bool
	_pinpointemailPutAccountDedicatedIpWarmupAttributes  bool
	_pinpointemailPutAccountSendingAttributes            bool
	_pinpointemailPutConfigurationSetDeliveryOptions     bool
	_pinpointemailPutConfigurationSetReputationOptions   bool
	_pinpointemailPutConfigurationSetSendingOptions      bool
	_pinpointemailPutConfigurationSetTrackingOptions     bool
	_pinpointemailPutDedicatedIpInPool                   bool
	_pinpointemailPutDedicatedIpWarmupAttributes         bool
	_pinpointemailPutDeliverabilityDashboardOption       bool
	_pinpointemailPutEmailIdentityDkimAttributes         bool
	_pinpointemailPutEmailIdentityFeedbackAttributes     bool
	_pinpointemailPutEmailIdentityMailFromAttributes     bool
	_pinpointemailSendEmail                              bool
	_pinpointemailTagResource                            bool
	_pinpointemailUntagResource                          bool
	_pinpointemailUpdateConfigurationSetEventDestination bool

	_pinpointemailAutoWarmupEnabled              string
	_pinpointemailBehaviorOnMxFailure            string
	_pinpointemailBlacklistItemNames             []string
	_pinpointemailCampaignId                     string
	_pinpointemailConfigurationSetName           string
	_pinpointemailContent                        string
	_pinpointemailCustomRedirectDomain           string
	_pinpointemailDashboardEnabled               string
	_pinpointemailDeliveryOptions                string
	_pinpointemailDestination                    string
	_pinpointemailDestinationPoolName            string
	_pinpointemailDomain                         string
	_pinpointemailEmailForwardingEnabled         string
	_pinpointemailEmailIdentity                  string
	_pinpointemailEmailTags                      string
	_pinpointemailEndDate                        string
	_pinpointemailEventDestination               string
	_pinpointemailEventDestinationName           string
	_pinpointemailFeedbackForwardingEmailAddress string
	_pinpointemailFromEmailAddress               string
	_pinpointemailIp                             string
	_pinpointemailMailFromDomain                 string
	_pinpointemailNextToken                      string
	_pinpointemailPageSize                       string
	_pinpointemailPoolName                       string
	_pinpointemailReplyToAddresses               []string
	_pinpointemailReportId                       string
	_pinpointemailReportName                     string
	_pinpointemailReputationMetricsEnabled       string
	_pinpointemailReputationOptions              string
	_pinpointemailResourceArn                    string
	_pinpointemailSendingEnabled                 string
	_pinpointemailSendingOptions                 string
	_pinpointemailSendingPoolName                string
	_pinpointemailSigningEnabled                 string
	_pinpointemailStartDate                      string
	_pinpointemailSubscribedDomain               string
	_pinpointemailSubscribedDomains              string
	_pinpointemailTagKeys                        []string
	_pinpointemailTags                           string
	_pinpointemailTlsPolicy                      string
	_pinpointemailTrackingOptions                string
	_pinpointemailWarmupPercentage               string
)

// Create a configuration set. Configuration sets are groups of rules that you can
// apply to the emails you send using Amazon Pinpoint. You apply a configuration
// set to an email by including a reference to the configuration set in the headers
// of the email. When you apply a configuration set to an email, all of the rules
// in that configuration set are applied to the email.
func pinpointemail_CreateConfigurationSet(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.CreateConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailDeliveryOptions) > 0 {
		if err := assignInputField(input, "DeliveryOptions", _pinpointemailDeliveryOptions); err != nil {
			log.Errorf("invalid --delivery-options: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailReputationOptions) > 0 {
		if err := assignInputField(input, "ReputationOptions", _pinpointemailReputationOptions); err != nil {
			log.Errorf("invalid --reputation-options: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailSendingOptions) > 0 {
		if err := assignInputField(input, "SendingOptions", _pinpointemailSendingOptions); err != nil {
			log.Errorf("invalid --sending-options: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailTags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointemailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailTrackingOptions) > 0 {
		if err := assignInputField(input, "TrackingOptions", _pinpointemailTrackingOptions); err != nil {
			log.Errorf("invalid --tracking-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create an event destination. In Amazon Pinpoint, events include message sends,
// deliveries, opens, clicks, bounces, and complaints. Event destinations are
// places that you can send information about these events to. For example, you can
// send event data to Amazon SNS to receive notifications when you receive bounces
// or complaints, or you can use Amazon Kinesis Data Firehose to stream data to
// Amazon S3 for long-term storage.
//
// A single configuration set can include more than one event destination.
func pinpointemail_CreateConfigurationSetEventDestination(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.CreateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestination: *types.EventDestinationDefinition, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailEventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _pinpointemailEventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointemailEventDestinationName)
	}

	if resp, err := client.CreateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new pool of dedicated IP addresses. A pool can include one or more
// dedicated IP addresses that are associated with your Amazon Pinpoint account.
// You can associate a pool with a configuration set. When you send an email that
// uses that configuration set, Amazon Pinpoint sends it using only the IP
// addresses in the associated pool.
func pinpointemail_CreateDedicatedIpPool(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.CreateDedicatedIpPoolInput{
		// PoolName: *string, // Required
	}

	if len(_pinpointemailPoolName) > 0 {
		input.PoolName = aws.String(_pinpointemailPoolName)
	}
	if len(_pinpointemailTags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointemailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDedicatedIpPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new predictive inbox placement test. Predictive inbox placement tests
// can help you predict how your messages will be handled by various email
// providers around the world. When you perform a predictive inbox placement test,
// you provide a sample message that contains the content that you plan to send to
// your customers. Amazon Pinpoint then sends that message to special email
// addresses spread across several major email providers. After about 24 hours, the
// test is complete, and you can use the GetDeliverabilityTestReport operation to
// view the results of the test.
func pinpointemail_CreateDeliverabilityTestReport(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.CreateDeliverabilityTestReportInput{
		// Content: *types.EmailContent, // Required
		// FromEmailAddress: *string, // Required
	}

	if len(_pinpointemailContent) > 0 {
		if err := assignInputField(input, "Content", _pinpointemailContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailFromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_pinpointemailFromEmailAddress)
	}
	if len(_pinpointemailReportName) > 0 {
		input.ReportName = aws.String(_pinpointemailReportName)
	}
	if len(_pinpointemailTags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointemailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDeliverabilityTestReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies an email identity for use with Amazon Pinpoint. In Amazon Pinpoint, an
// identity is an email address or domain that you use when you send email. Before
// you can use an identity to send email with Amazon Pinpoint, you first have to
// verify it. By verifying an address, you demonstrate that you're the owner of the
// address, and that you've given Amazon Pinpoint permission to send email from the
// address.
//
// When you verify an email address, Amazon Pinpoint sends an email to the
// address. Your email address is verified as soon as you follow the link in the
// verification email.
//
// When you verify a domain, this operation provides a set of DKIM tokens, which
// you can convert into CNAME tokens. You add these CNAME tokens to the DNS
// configuration for your domain. Your domain is verified when Amazon Pinpoint
// detects these records in the DNS configuration for your domain. It usually takes
// around 72 hours to complete the domain verification process.
func pinpointemail_CreateEmailIdentity(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.CreateEmailIdentityInput{
		// EmailIdentity: *string, // Required
	}

	if len(_pinpointemailEmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_pinpointemailEmailIdentity)
	}
	if len(_pinpointemailTags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointemailTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEmailIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing configuration set.
// In Amazon Pinpoint, configuration sets are groups of rules that you can apply
// to the emails you send. You apply a configuration set to an email by including a
// reference to the configuration set in the headers of the email. When you apply a
// configuration set to an email, all of the rules in that configuration set are
// applied to the email.
func pinpointemail_DeleteConfigurationSet(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.DeleteConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}

	if resp, err := client.DeleteConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an event destination.
// In Amazon Pinpoint, events include message sends, deliveries, opens, clicks,
// bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon SNS to receive notifications when you receive bounces or complaints, or
// you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for
// long-term storage.
func pinpointemail_DeleteConfigurationSetEventDestination(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.DeleteConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointemailEventDestinationName)
	}

	if resp, err := client.DeleteConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a dedicated IP pool.
func pinpointemail_DeleteDedicatedIpPool(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.DeleteDedicatedIpPoolInput{
		// PoolName: *string, // Required
	}

	if len(_pinpointemailPoolName) > 0 {
		input.PoolName = aws.String(_pinpointemailPoolName)
	}

	if resp, err := client.DeleteDedicatedIpPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an email identity that you previously verified for use with Amazon
// Pinpoint. An identity can be either an email address or a domain name.
func pinpointemail_DeleteEmailIdentity(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.DeleteEmailIdentityInput{
		// EmailIdentity: *string, // Required
	}

	if len(_pinpointemailEmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_pinpointemailEmailIdentity)
	}

	if resp, err := client.DeleteEmailIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtain information about the email-sending status and capabilities of your
// Amazon Pinpoint account in the current AWS Region.
func pinpointemail_GetAccount(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetAccountInput{}

	if resp, err := client.GetAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a list of the blacklists that your dedicated IP addresses appear on.
func pinpointemail_GetBlacklistReports(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetBlacklistReportsInput{
		// BlacklistItemNames: []string, // Required
	}

	if len(_pinpointemailBlacklistItemNames) > 0 {
		input.BlacklistItemNames = append([]string(nil), _pinpointemailBlacklistItemNames...)
	}

	if resp, err := client.GetBlacklistReports(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about an existing configuration set, including the dedicated IP
// pool that it's associated with, whether or not it's enabled for sending email,
// and more.
//
// In Amazon Pinpoint, configuration sets are groups of rules that you can apply
// to the emails you send. You apply a configuration set to an email by including a
// reference to the configuration set in the headers of the email. When you apply a
// configuration set to an email, all of the rules in that configuration set are
// applied to the email.
func pinpointemail_GetConfigurationSet(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}

	if resp, err := client.GetConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a list of event destinations that are associated with a configuration
// set.
//
// In Amazon Pinpoint, events include message sends, deliveries, opens, clicks,
// bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon SNS to receive notifications when you receive bounces or complaints, or
// you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for
// long-term storage.
func pinpointemail_GetConfigurationSetEventDestinations(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetConfigurationSetEventDestinationsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}

	if resp, err := client.GetConfigurationSetEventDestinations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about a dedicated IP address, including the name of the
// dedicated IP pool that it's associated with, as well information about the
// automatic warm-up process for the address.
func pinpointemail_GetDedicatedIp(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetDedicatedIpInput{
		// Ip: *string, // Required
	}

	if len(_pinpointemailIp) > 0 {
		input.Ip = aws.String(_pinpointemailIp)
	}

	if resp, err := client.GetDedicatedIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the dedicated IP addresses that are associated with your Amazon Pinpoint
// account.
func pinpointemail_GetDedicatedIps(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetDedicatedIpsInput{}

	if len(_pinpointemailNextToken) > 0 {
		input.NextToken = aws.String(_pinpointemailNextToken)
	}
	if len(_pinpointemailPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _pinpointemailPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailPoolName) > 0 {
		input.PoolName = aws.String(_pinpointemailPoolName)
	}

	if disablePaginator() {
		if resp, err := client.GetDedicatedIps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointemail.GetDedicatedIpsOutput
	p := pinpointemail.NewGetDedicatedIpsPaginator(client, input)
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

// Retrieve information about the status of the Deliverability dashboard for your
// Amazon Pinpoint account. When the Deliverability dashboard is enabled, you gain
// access to reputation, deliverability, and other metrics for the domains that you
// use to send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests.
//
// When you use the Deliverability dashboard, you pay a monthly subscription
// charge, in addition to any other fees that you accrue by using Amazon Pinpoint.
// For more information about the features and cost of a Deliverability dashboard
// subscription, see [Amazon Pinpoint Pricing].
//
// [Amazon Pinpoint Pricing]: http://aws.amazon.com/pinpoint/pricing/
func pinpointemail_GetDeliverabilityDashboardOptions(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetDeliverabilityDashboardOptionsInput{}

	if resp, err := client.GetDeliverabilityDashboardOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the results of a predictive inbox placement test.
func pinpointemail_GetDeliverabilityTestReport(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetDeliverabilityTestReportInput{
		// ReportId: *string, // Required
	}

	if len(_pinpointemailReportId) > 0 {
		input.ReportId = aws.String(_pinpointemailReportId)
	}

	if resp, err := client.GetDeliverabilityTestReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve all the deliverability data for a specific campaign. This data is
// available for a campaign only if the campaign sent email by using a domain that
// the Deliverability dashboard is enabled for ( PutDeliverabilityDashboardOption
// operation).
func pinpointemail_GetDomainDeliverabilityCampaign(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetDomainDeliverabilityCampaignInput{
		// CampaignId: *string, // Required
	}

	if len(_pinpointemailCampaignId) > 0 {
		input.CampaignId = aws.String(_pinpointemailCampaignId)
	}

	if resp, err := client.GetDomainDeliverabilityCampaign(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve inbox placement and engagement rates for the domains that you use to
// send email.
func pinpointemail_GetDomainStatisticsReport(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetDomainStatisticsReportInput{
		// Domain: *string, // Required
		// EndDate: *time.Time, // Required
		// StartDate: *time.Time, // Required
	}

	if len(_pinpointemailDomain) > 0 {
		input.Domain = aws.String(_pinpointemailDomain)
	}
	if len(_pinpointemailEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _pinpointemailEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _pinpointemailStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDomainStatisticsReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a specific identity associated with your Amazon
// Pinpoint account, including the identity's verification status, its DKIM
// authentication status, and its custom Mail-From settings.
func pinpointemail_GetEmailIdentity(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.GetEmailIdentityInput{
		// EmailIdentity: *string, // Required
	}

	if len(_pinpointemailEmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_pinpointemailEmailIdentity)
	}

	if resp, err := client.GetEmailIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all of the configuration sets associated with your Amazon Pinpoint account
// in the current region.
//
// In Amazon Pinpoint, configuration sets are groups of rules that you can apply
// to the emails you send. You apply a configuration set to an email by including a
// reference to the configuration set in the headers of the email. When you apply a
// configuration set to an email, all of the rules in that configuration set are
// applied to the email.
func pinpointemail_ListConfigurationSets(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.ListConfigurationSetsInput{}

	if len(_pinpointemailNextToken) > 0 {
		input.NextToken = aws.String(_pinpointemailNextToken)
	}
	if len(_pinpointemailPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _pinpointemailPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConfigurationSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointemail.ListConfigurationSetsOutput
	p := pinpointemail.NewListConfigurationSetsPaginator(client, input)
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

// List all of the dedicated IP pools that exist in your Amazon Pinpoint account
// in the current AWS Region.
func pinpointemail_ListDedicatedIpPools(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.ListDedicatedIpPoolsInput{}

	if len(_pinpointemailNextToken) > 0 {
		input.NextToken = aws.String(_pinpointemailNextToken)
	}
	if len(_pinpointemailPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _pinpointemailPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDedicatedIpPools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointemail.ListDedicatedIpPoolsOutput
	p := pinpointemail.NewListDedicatedIpPoolsPaginator(client, input)
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

// Show a list of the predictive inbox placement tests that you've performed,
// regardless of their statuses. For predictive inbox placement tests that are
// complete, you can use the GetDeliverabilityTestReport operation to view the
// results.
func pinpointemail_ListDeliverabilityTestReports(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.ListDeliverabilityTestReportsInput{}

	if len(_pinpointemailNextToken) > 0 {
		input.NextToken = aws.String(_pinpointemailNextToken)
	}
	if len(_pinpointemailPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _pinpointemailPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDeliverabilityTestReports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointemail.ListDeliverabilityTestReportsOutput
	p := pinpointemail.NewListDeliverabilityTestReportsPaginator(client, input)
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

// Retrieve deliverability data for all the campaigns that used a specific domain
// to send email during a specified time range. This data is available for a domain
// only if you enabled the Deliverability dashboard (
// PutDeliverabilityDashboardOption operation) for the domain.
func pinpointemail_ListDomainDeliverabilityCampaigns(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.ListDomainDeliverabilityCampaignsInput{
		// EndDate: *time.Time, // Required
		// StartDate: *time.Time, // Required
		// SubscribedDomain: *string, // Required
	}

	if len(_pinpointemailEndDate) > 0 {
		if err := assignInputField(input, "EndDate", _pinpointemailEndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailStartDate) > 0 {
		if err := assignInputField(input, "StartDate", _pinpointemailStartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailSubscribedDomain) > 0 {
		input.SubscribedDomain = aws.String(_pinpointemailSubscribedDomain)
	}
	if len(_pinpointemailNextToken) > 0 {
		input.NextToken = aws.String(_pinpointemailNextToken)
	}
	if len(_pinpointemailPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _pinpointemailPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDomainDeliverabilityCampaigns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointemail.ListDomainDeliverabilityCampaignsOutput
	p := pinpointemail.NewListDomainDeliverabilityCampaignsPaginator(client, input)
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

// Returns a list of all of the email identities that are associated with your
// Amazon Pinpoint account. An identity can be either an email address or a domain.
// This operation returns identities that are verified as well as those that
// aren't.
func pinpointemail_ListEmailIdentities(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.ListEmailIdentitiesInput{}

	if len(_pinpointemailNextToken) > 0 {
		input.NextToken = aws.String(_pinpointemailNextToken)
	}
	if len(_pinpointemailPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _pinpointemailPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEmailIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointemail.ListEmailIdentitiesOutput
	p := pinpointemail.NewListEmailIdentitiesPaginator(client, input)
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

// Retrieve a list of the tags (keys and values) that are associated with a
// specified resource. A tag is a label that you optionally define and associate
// with a resource in Amazon Pinpoint. Each tag consists of a required tag key and
// an optional associated tag value. A tag key is a general label that acts as a
// category for more specific tag values. A tag value acts as a descriptor within a
// tag key.
func pinpointemail_ListTagsForResource(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pinpointemailResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointemailResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable or disable the automatic warm-up feature for dedicated IP addresses.
func pinpointemail_PutAccountDedicatedIpWarmupAttributes(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutAccountDedicatedIpWarmupAttributesInput{}

	if len(_pinpointemailAutoWarmupEnabled) > 0 {
		if err := assignInputField(input, "AutoWarmupEnabled", _pinpointemailAutoWarmupEnabled); err != nil {
			log.Errorf("invalid --auto-warmup-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountDedicatedIpWarmupAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable or disable the ability of your account to send email.
func pinpointemail_PutAccountSendingAttributes(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutAccountSendingAttributesInput{}

	if len(_pinpointemailSendingEnabled) > 0 {
		if err := assignInputField(input, "SendingEnabled", _pinpointemailSendingEnabled); err != nil {
			log.Errorf("invalid --sending-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountSendingAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a configuration set with a dedicated IP pool. You can use dedicated
// IP pools to create groups of dedicated IP addresses for sending specific types
// of email.
func pinpointemail_PutConfigurationSetDeliveryOptions(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutConfigurationSetDeliveryOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailSendingPoolName) > 0 {
		input.SendingPoolName = aws.String(_pinpointemailSendingPoolName)
	}
	if len(_pinpointemailTlsPolicy) > 0 {
		if err := assignInputField(input, "TlsPolicy", _pinpointemailTlsPolicy); err != nil {
			log.Errorf("invalid --tls-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationSetDeliveryOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable or disable collection of reputation metrics for emails that you send
// using a particular configuration set in a specific AWS Region.
func pinpointemail_PutConfigurationSetReputationOptions(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutConfigurationSetReputationOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailReputationMetricsEnabled) > 0 {
		if err := assignInputField(input, "ReputationMetricsEnabled", _pinpointemailReputationMetricsEnabled); err != nil {
			log.Errorf("invalid --reputation-metrics-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationSetReputationOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable or disable email sending for messages that use a particular
// configuration set in a specific AWS Region.
func pinpointemail_PutConfigurationSetSendingOptions(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutConfigurationSetSendingOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailSendingEnabled) > 0 {
		if err := assignInputField(input, "SendingEnabled", _pinpointemailSendingEnabled); err != nil {
			log.Errorf("invalid --sending-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationSetSendingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify a custom domain to use for open and click tracking elements in email
// that you send using Amazon Pinpoint.
func pinpointemail_PutConfigurationSetTrackingOptions(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutConfigurationSetTrackingOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailCustomRedirectDomain) > 0 {
		input.CustomRedirectDomain = aws.String(_pinpointemailCustomRedirectDomain)
	}

	if resp, err := client.PutConfigurationSetTrackingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Move a dedicated IP address to an existing dedicated IP pool.
// The dedicated IP address that you specify must already exist, and must be
// associated with your Amazon Pinpoint account.
//
// The dedicated IP pool you specify must already exist. You can create a new pool
// by using the CreateDedicatedIpPool operation.
func pinpointemail_PutDedicatedIpInPool(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutDedicatedIpInPoolInput{
		// DestinationPoolName: *string, // Required
		// Ip: *string, // Required
	}

	if len(_pinpointemailDestinationPoolName) > 0 {
		input.DestinationPoolName = aws.String(_pinpointemailDestinationPoolName)
	}
	if len(_pinpointemailIp) > 0 {
		input.Ip = aws.String(_pinpointemailIp)
	}

	if resp, err := client.PutDedicatedIpInPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func pinpointemail_PutDedicatedIpWarmupAttributes(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutDedicatedIpWarmupAttributesInput{
		// Ip: *string, // Required
		// WarmupPercentage: *int32, // Required
	}

	if len(_pinpointemailIp) > 0 {
		input.Ip = aws.String(_pinpointemailIp)
	}
	if len(_pinpointemailWarmupPercentage) > 0 {
		if err := assignInputField(input, "WarmupPercentage", _pinpointemailWarmupPercentage); err != nil {
			log.Errorf("invalid --warmup-percentage: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDedicatedIpWarmupAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable or disable the Deliverability dashboard for your Amazon Pinpoint
// account. When you enable the Deliverability dashboard, you gain access to
// reputation, deliverability, and other metrics for the domains that you use to
// send email using Amazon Pinpoint. You also gain the ability to perform
// predictive inbox placement tests.
//
// When you use the Deliverability dashboard, you pay a monthly subscription
// charge, in addition to any other fees that you accrue by using Amazon Pinpoint.
// For more information about the features and cost of a Deliverability dashboard
// subscription, see [Amazon Pinpoint Pricing].
//
// [Amazon Pinpoint Pricing]: http://aws.amazon.com/pinpoint/pricing/
func pinpointemail_PutDeliverabilityDashboardOption(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutDeliverabilityDashboardOptionInput{
		// DashboardEnabled: bool, // Required
	}

	if len(_pinpointemailDashboardEnabled) > 0 {
		if err := assignInputField(input, "DashboardEnabled", _pinpointemailDashboardEnabled); err != nil {
			log.Errorf("invalid --dashboard-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailSubscribedDomains) > 0 {
		if err := assignInputField(input, "SubscribedDomains", _pinpointemailSubscribedDomains); err != nil {
			log.Errorf("invalid --subscribed-domains: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDeliverabilityDashboardOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to enable or disable DKIM authentication for an email identity.
func pinpointemail_PutEmailIdentityDkimAttributes(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutEmailIdentityDkimAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_pinpointemailEmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_pinpointemailEmailIdentity)
	}
	if len(_pinpointemailSigningEnabled) > 0 {
		if err := assignInputField(input, "SigningEnabled", _pinpointemailSigningEnabled); err != nil {
			log.Errorf("invalid --signing-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEmailIdentityDkimAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to enable or disable feedback forwarding for an identity. This setting
// determines what happens when an identity is used to send an email that results
// in a bounce or complaint event.
//
// When you enable feedback forwarding, Amazon Pinpoint sends you email
// notifications when bounce or complaint events occur. Amazon Pinpoint sends this
// notification to the address that you specified in the Return-Path header of the
// original email.
//
// When you disable feedback forwarding, Amazon Pinpoint sends notifications
// through other mechanisms, such as by notifying an Amazon SNS topic. You're
// required to have a method of tracking bounces and complaints. If you haven't set
// up another mechanism for receiving bounce or complaint notifications, Amazon
// Pinpoint sends an email notification when these events occur (even if this
// setting is disabled).
func pinpointemail_PutEmailIdentityFeedbackAttributes(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutEmailIdentityFeedbackAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_pinpointemailEmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_pinpointemailEmailIdentity)
	}
	if len(_pinpointemailEmailForwardingEnabled) > 0 {
		if err := assignInputField(input, "EmailForwardingEnabled", _pinpointemailEmailForwardingEnabled); err != nil {
			log.Errorf("invalid --email-forwarding-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEmailIdentityFeedbackAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to enable or disable the custom Mail-From domain configuration for an
// email identity.
func pinpointemail_PutEmailIdentityMailFromAttributes(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.PutEmailIdentityMailFromAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_pinpointemailEmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_pinpointemailEmailIdentity)
	}
	if len(_pinpointemailBehaviorOnMxFailure) > 0 {
		if err := assignInputField(input, "BehaviorOnMxFailure", _pinpointemailBehaviorOnMxFailure); err != nil {
			log.Errorf("invalid --behavior-on-mx-failure: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailMailFromDomain) > 0 {
		input.MailFromDomain = aws.String(_pinpointemailMailFromDomain)
	}

	if resp, err := client.PutEmailIdentityMailFromAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an email message. You can use the Amazon Pinpoint Email API to send two
// types of messages:
//
// - Simple – A standard email message. When you create this type of message,
// you specify the sender, the recipient, and the message body, and Amazon Pinpoint
// assembles the message for you.
//
// - Raw – A raw, MIME-formatted email message. When you send this type of
// email, you have to specify all of the message headers, as well as the message
// body. You can use this message type to send messages that contain attachments.
// The message that you specify has to be a valid MIME message.
func pinpointemail_SendEmail(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.SendEmailInput{
		// Content: *types.EmailContent, // Required
		// Destination: *types.Destination, // Required
	}

	if len(_pinpointemailContent) > 0 {
		if err := assignInputField(input, "Content", _pinpointemailContent); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailDestination) > 0 {
		if err := assignInputField(input, "Destination", _pinpointemailDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailEmailTags) > 0 {
		if err := assignInputField(input, "EmailTags", _pinpointemailEmailTags); err != nil {
			log.Errorf("invalid --email-tags: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailFeedbackForwardingEmailAddress) > 0 {
		input.FeedbackForwardingEmailAddress = aws.String(_pinpointemailFeedbackForwardingEmailAddress)
	}
	if len(_pinpointemailFromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_pinpointemailFromEmailAddress)
	}
	if len(_pinpointemailReplyToAddresses) > 0 {
		input.ReplyToAddresses = append([]string(nil), _pinpointemailReplyToAddresses...)
	}

	if resp, err := client.SendEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add one or more tags (keys and values) to a specified resource. A tag is a
// label that you optionally define and associate with a resource in Amazon
// Pinpoint. Tags can help you categorize and manage resources in different ways,
// such as by purpose, owner, environment, or other criteria. A resource can have
// as many as 50 tags.
//
// Each tag consists of a required tag key and an associated tag value, both of
// which you define. A tag key is a general label that acts as a category for more
// specific tag values. A tag value acts as a descriptor within a tag key.
func pinpointemail_TagResource(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_pinpointemailResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointemailResourceArn)
	}
	if len(_pinpointemailTags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointemailTags); err != nil {
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

// Remove one or more tags (keys and values) from a specified resource.
func pinpointemail_UntagResource(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pinpointemailResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointemailResourceArn)
	}
	if len(_pinpointemailTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pinpointemailTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the configuration of an event destination for a configuration set.
// In Amazon Pinpoint, events include message sends, deliveries, opens, clicks,
// bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon SNS to receive notifications when you receive bounces or complaints, or
// you can use Amazon Kinesis Data Firehose to stream data to Amazon S3 for
// long-term storage.
func pinpointemail_UpdateConfigurationSetEventDestination(cfg aws.Config, client *pinpointemail.Client) {
	input := &pinpointemail.UpdateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestination: *types.EventDestinationDefinition, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointemailConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointemailConfigurationSetName)
	}
	if len(_pinpointemailEventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _pinpointemailEventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointemailEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointemailEventDestinationName)
	}

	if resp, err := client.UpdateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pinpointemailCmd)
	_pinpointemailCmd.Flags().SortFlags = false

	_pinpointemailCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_pinpointemailCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pinpointemailCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailAutoWarmupEnabled, "auto-warmup-enabled", "", "", "Auto Warmup Enabled")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailBehaviorOnMxFailure, "behavior-on-mx-failure", "", "", "Behavior On Mx Failure")
	_pinpointemailCmd.Flags().StringSliceVarP(&_pinpointemailBlacklistItemNames, "blacklist-item-names", "", nil, "Blacklist Item Names")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailCampaignId, "campaign-id", "", "", "Campaign ID")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailConfigurationSetName, "configuration-set-name", "", "", "Configuration Set Name")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailContent, "content", "", "", "Content")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailCustomRedirectDomain, "custom-redirect-domain", "", "", "Custom Redirect Domain")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailDashboardEnabled, "dashboard-enabled", "", "", "Dashboard Enabled")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailDeliveryOptions, "delivery-options", "", "", "Delivery Options")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailDestination, "destination", "", "", "Destination")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailDestinationPoolName, "destination-pool-name", "", "", "Destination Pool Name")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailDomain, "domain", "", "", "Domain")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailEmailForwardingEnabled, "email-forwarding-enabled", "", "", "Email Forwarding Enabled")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailEmailIdentity, "email-identity", "", "", "Email Identity")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailEmailTags, "email-tags", "", "", "Email Tags")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailEndDate, "end-date", "", "", "End Date")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailEventDestination, "event-destination", "", "", "Event Destination")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailEventDestinationName, "event-destination-name", "", "", "Event Destination Name")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailFeedbackForwardingEmailAddress, "feedback-forwarding-email-address", "", "", "Feedback Forwarding Email Address")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailFromEmailAddress, "from-email-address", "", "", "From Email Address")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailIp, "ip", "", "", "IP")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailMailFromDomain, "mail-from-domain", "", "", "Mail From Domain")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailNextToken, "next-token", "", "", "Next Token")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailPageSize, "page-size", "", "", "Page Size")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailPoolName, "pool-name", "", "", "Pool Name")
	_pinpointemailCmd.Flags().StringSliceVarP(&_pinpointemailReplyToAddresses, "reply-to-addresses", "", nil, "Reply To Addresses")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailReportId, "report-id", "", "", "Report ID")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailReportName, "report-name", "", "", "Report Name")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailReputationMetricsEnabled, "reputation-metrics-enabled", "", "", "Reputation Metrics Enabled")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailReputationOptions, "reputation-options", "", "", "Reputation Options")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailResourceArn, "resource-arn", "", "", "Resource ARN")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailSendingEnabled, "sending-enabled", "", "", "Sending Enabled")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailSendingOptions, "sending-options", "", "", "Sending Options")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailSendingPoolName, "sending-pool-name", "", "", "Sending Pool Name")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailSigningEnabled, "signing-enabled", "", "", "Signing Enabled")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailStartDate, "start-date", "", "", "Start Date")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailSubscribedDomain, "subscribed-domain", "", "", "Subscribed Domain")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailSubscribedDomains, "subscribed-domains", "", "", "Subscribed Domains")
	_pinpointemailCmd.Flags().StringSliceVarP(&_pinpointemailTagKeys, "tag-keys", "", nil, "Tag Keys")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailTags, "tags", "", "", "Tags")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailTlsPolicy, "tls-policy", "", "", "TLS Policy")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailTrackingOptions, "tracking-options", "", "", "Tracking Options")
	_pinpointemailCmd.Flags().StringVarP(&_pinpointemailWarmupPercentage, "warmup-percentage", "", "", "Warmup Percentage")

	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailCreateConfigurationSet, "create-configuration-set", "", false, "Create Configuration Set")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailCreateConfigurationSetEventDestination, "create-configuration-set-event-destination", "", false, "Create Configuration Set Event Destination")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailCreateDedicatedIpPool, "create-dedicated-ip-pool", "", false, "Create Dedicated IP Pool")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailCreateDeliverabilityTestReport, "create-deliverability-test-report", "", false, "Create Deliverability Test Report")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailCreateEmailIdentity, "create-email-identity", "", false, "Create Email Identity")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailDeleteConfigurationSet, "delete-configuration-set", "", false, "Delete Configuration Set")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailDeleteConfigurationSetEventDestination, "delete-configuration-set-event-destination", "", false, "Delete Configuration Set Event Destination")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailDeleteDedicatedIpPool, "delete-dedicated-ip-pool", "", false, "Delete Dedicated IP Pool")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailDeleteEmailIdentity, "delete-email-identity", "", false, "Delete Email Identity")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetAccount, "get-account", "", false, "Get Account")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetBlacklistReports, "get-blacklist-reports", "", false, "Get Blacklist Reports")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetConfigurationSet, "get-configuration-set", "", false, "Get Configuration Set")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetConfigurationSetEventDestinations, "get-configuration-set-event-destinations", "", false, "Get Configuration Set Event Destinations")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetDedicatedIp, "get-dedicated-ip", "", false, "Get Dedicated IP")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetDedicatedIps, "get-dedicated-ips", "", false, "Get Dedicated Ips")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetDeliverabilityDashboardOptions, "get-deliverability-dashboard-options", "", false, "Get Deliverability Dashboard Options")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetDeliverabilityTestReport, "get-deliverability-test-report", "", false, "Get Deliverability Test Report")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetDomainDeliverabilityCampaign, "get-domain-deliverability-campaign", "", false, "Get Domain Deliverability Campaign")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetDomainStatisticsReport, "get-domain-statistics-report", "", false, "Get Domain Statistics Report")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailGetEmailIdentity, "get-email-identity", "", false, "Get Email Identity")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailListConfigurationSets, "list-configuration-sets", "", false, "List Configuration Sets")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailListDedicatedIpPools, "list-dedicated-ip-pools", "", false, "List Dedicated IP Pools")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailListDeliverabilityTestReports, "list-deliverability-test-reports", "", false, "List Deliverability Test Reports")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailListDomainDeliverabilityCampaigns, "list-domain-deliverability-campaigns", "", false, "List Domain Deliverability Campaigns")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailListEmailIdentities, "list-email-identities", "", false, "List Email Identities")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutAccountDedicatedIpWarmupAttributes, "put-account-dedicated-ip-warmup-attributes", "", false, "Put Account Dedicated IP Warmup Attributes")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutAccountSendingAttributes, "put-account-sending-attributes", "", false, "Put Account Sending Attributes")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutConfigurationSetDeliveryOptions, "put-configuration-set-delivery-options", "", false, "Put Configuration Set Delivery Options")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutConfigurationSetReputationOptions, "put-configuration-set-reputation-options", "", false, "Put Configuration Set Reputation Options")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutConfigurationSetSendingOptions, "put-configuration-set-sending-options", "", false, "Put Configuration Set Sending Options")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutConfigurationSetTrackingOptions, "put-configuration-set-tracking-options", "", false, "Put Configuration Set Tracking Options")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutDedicatedIpInPool, "put-dedicated-ip-in-pool", "", false, "Put Dedicated IP In Pool")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutDedicatedIpWarmupAttributes, "put-dedicated-ip-warmup-attributes", "", false, "Put Dedicated IP Warmup Attributes")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutDeliverabilityDashboardOption, "put-deliverability-dashboard-option", "", false, "Put Deliverability Dashboard Option")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutEmailIdentityDkimAttributes, "put-email-identity-dkim-attributes", "", false, "Put Email Identity Dkim Attributes")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutEmailIdentityFeedbackAttributes, "put-email-identity-feedback-attributes", "", false, "Put Email Identity Feedback Attributes")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailPutEmailIdentityMailFromAttributes, "put-email-identity-mail-from-attributes", "", false, "Put Email Identity Mail From Attributes")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailSendEmail, "send-email", "", false, "Send Email")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailTagResource, "tag-resource", "", false, "Tag Resource")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailUntagResource, "untag-resource", "", false, "Untag Resource")
	_pinpointemailCmd.Flags().BoolVarP(&_pinpointemailUpdateConfigurationSetEventDestination, "update-configuration-set-event-destination", "", false, "Update Configuration Set Event Destination")

}
