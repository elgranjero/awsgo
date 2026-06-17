package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sesCmd represents the ses command
var _sesCmd = &cobra.Command{
	Use:   "ses",
	Short: "AWS ses CLI",
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
		client := ses.NewFromConfig(cfg)
		if _sesCloneReceiptRuleSet {
			ses_CloneReceiptRuleSet(cfg, client)
			return
		}
		if _sesCreateConfigurationSet {
			ses_CreateConfigurationSet(cfg, client)
			return
		}
		if _sesCreateConfigurationSetEventDestination {
			ses_CreateConfigurationSetEventDestination(cfg, client)
			return
		}
		if _sesCreateConfigurationSetTrackingOptions {
			ses_CreateConfigurationSetTrackingOptions(cfg, client)
			return
		}
		if _sesCreateCustomVerificationEmailTemplate {
			ses_CreateCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesCreateReceiptFilter {
			ses_CreateReceiptFilter(cfg, client)
			return
		}
		if _sesCreateReceiptRule {
			ses_CreateReceiptRule(cfg, client)
			return
		}
		if _sesCreateReceiptRuleSet {
			ses_CreateReceiptRuleSet(cfg, client)
			return
		}
		if _sesCreateTemplate {
			ses_CreateTemplate(cfg, client)
			return
		}
		if _sesDeleteConfigurationSet {
			ses_DeleteConfigurationSet(cfg, client)
			return
		}
		if _sesDeleteConfigurationSetEventDestination {
			ses_DeleteConfigurationSetEventDestination(cfg, client)
			return
		}
		if _sesDeleteConfigurationSetTrackingOptions {
			ses_DeleteConfigurationSetTrackingOptions(cfg, client)
			return
		}
		if _sesDeleteCustomVerificationEmailTemplate {
			ses_DeleteCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesDeleteIdentity {
			ses_DeleteIdentity(cfg, client)
			return
		}
		if _sesDeleteIdentityPolicy {
			ses_DeleteIdentityPolicy(cfg, client)
			return
		}
		if _sesDeleteReceiptFilter {
			ses_DeleteReceiptFilter(cfg, client)
			return
		}
		if _sesDeleteReceiptRule {
			ses_DeleteReceiptRule(cfg, client)
			return
		}
		if _sesDeleteReceiptRuleSet {
			ses_DeleteReceiptRuleSet(cfg, client)
			return
		}
		if _sesDeleteTemplate {
			ses_DeleteTemplate(cfg, client)
			return
		}
		if _sesDeleteVerifiedEmailAddress {
			ses_DeleteVerifiedEmailAddress(cfg, client)
			return
		}
		if _sesDescribeActiveReceiptRuleSet {
			ses_DescribeActiveReceiptRuleSet(cfg, client)
			return
		}
		if _sesDescribeConfigurationSet {
			ses_DescribeConfigurationSet(cfg, client)
			return
		}
		if _sesDescribeReceiptRule {
			ses_DescribeReceiptRule(cfg, client)
			return
		}
		if _sesDescribeReceiptRuleSet {
			ses_DescribeReceiptRuleSet(cfg, client)
			return
		}
		if _sesGetAccountSendingEnabled {
			ses_GetAccountSendingEnabled(cfg, client)
			return
		}
		if _sesGetCustomVerificationEmailTemplate {
			ses_GetCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesGetIdentityDkimAttributes {
			ses_GetIdentityDkimAttributes(cfg, client)
			return
		}
		if _sesGetIdentityMailFromDomainAttributes {
			ses_GetIdentityMailFromDomainAttributes(cfg, client)
			return
		}
		if _sesGetIdentityNotificationAttributes {
			ses_GetIdentityNotificationAttributes(cfg, client)
			return
		}
		if _sesGetIdentityPolicies {
			ses_GetIdentityPolicies(cfg, client)
			return
		}
		if _sesGetIdentityVerificationAttributes {
			ses_GetIdentityVerificationAttributes(cfg, client)
			return
		}
		if _sesGetSendQuota {
			ses_GetSendQuota(cfg, client)
			return
		}
		if _sesGetSendStatistics {
			ses_GetSendStatistics(cfg, client)
			return
		}
		if _sesGetTemplate {
			ses_GetTemplate(cfg, client)
			return
		}
		if _sesListConfigurationSets {
			ses_ListConfigurationSets(cfg, client)
			return
		}
		if _sesListCustomVerificationEmailTemplates {
			ses_ListCustomVerificationEmailTemplates(cfg, client)
			return
		}
		if _sesListIdentities {
			ses_ListIdentities(cfg, client)
			return
		}
		if _sesListIdentityPolicies {
			ses_ListIdentityPolicies(cfg, client)
			return
		}
		if _sesListReceiptFilters {
			ses_ListReceiptFilters(cfg, client)
			return
		}
		if _sesListReceiptRuleSets {
			ses_ListReceiptRuleSets(cfg, client)
			return
		}
		if _sesListTemplates {
			ses_ListTemplates(cfg, client)
			return
		}
		if _sesListVerifiedEmailAddresses {
			ses_ListVerifiedEmailAddresses(cfg, client)
			return
		}
		if _sesPutConfigurationSetDeliveryOptions {
			ses_PutConfigurationSetDeliveryOptions(cfg, client)
			return
		}
		if _sesPutIdentityPolicy {
			ses_PutIdentityPolicy(cfg, client)
			return
		}
		if _sesReorderReceiptRuleSet {
			ses_ReorderReceiptRuleSet(cfg, client)
			return
		}
		if _sesSendBounce {
			ses_SendBounce(cfg, client)
			return
		}
		if _sesSendBulkTemplatedEmail {
			ses_SendBulkTemplatedEmail(cfg, client)
			return
		}
		if _sesSendCustomVerificationEmail {
			ses_SendCustomVerificationEmail(cfg, client)
			return
		}
		if _sesSendEmail {
			ses_SendEmail(cfg, client)
			return
		}
		if _sesSendRawEmail {
			ses_SendRawEmail(cfg, client)
			return
		}
		if _sesSendTemplatedEmail {
			ses_SendTemplatedEmail(cfg, client)
			return
		}
		if _sesSetActiveReceiptRuleSet {
			ses_SetActiveReceiptRuleSet(cfg, client)
			return
		}
		if _sesSetIdentityDkimEnabled {
			ses_SetIdentityDkimEnabled(cfg, client)
			return
		}
		if _sesSetIdentityFeedbackForwardingEnabled {
			ses_SetIdentityFeedbackForwardingEnabled(cfg, client)
			return
		}
		if _sesSetIdentityHeadersInNotificationsEnabled {
			ses_SetIdentityHeadersInNotificationsEnabled(cfg, client)
			return
		}
		if _sesSetIdentityMailFromDomain {
			ses_SetIdentityMailFromDomain(cfg, client)
			return
		}
		if _sesSetIdentityNotificationTopic {
			ses_SetIdentityNotificationTopic(cfg, client)
			return
		}
		if _sesSetReceiptRulePosition {
			ses_SetReceiptRulePosition(cfg, client)
			return
		}
		if _sesTestRenderTemplate {
			ses_TestRenderTemplate(cfg, client)
			return
		}
		if _sesUpdateAccountSendingEnabled {
			ses_UpdateAccountSendingEnabled(cfg, client)
			return
		}
		if _sesUpdateConfigurationSetEventDestination {
			ses_UpdateConfigurationSetEventDestination(cfg, client)
			return
		}
		if _sesUpdateConfigurationSetReputationMetricsEnabled {
			ses_UpdateConfigurationSetReputationMetricsEnabled(cfg, client)
			return
		}
		if _sesUpdateConfigurationSetSendingEnabled {
			ses_UpdateConfigurationSetSendingEnabled(cfg, client)
			return
		}
		if _sesUpdateConfigurationSetTrackingOptions {
			ses_UpdateConfigurationSetTrackingOptions(cfg, client)
			return
		}
		if _sesUpdateCustomVerificationEmailTemplate {
			ses_UpdateCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesUpdateReceiptRule {
			ses_UpdateReceiptRule(cfg, client)
			return
		}
		if _sesUpdateTemplate {
			ses_UpdateTemplate(cfg, client)
			return
		}
		if _sesVerifyDomainDkim {
			ses_VerifyDomainDkim(cfg, client)
			return
		}
		if _sesVerifyDomainIdentity {
			ses_VerifyDomainIdentity(cfg, client)
			return
		}
		if _sesVerifyEmailAddress {
			ses_VerifyEmailAddress(cfg, client)
			return
		}
		if _sesVerifyEmailIdentity {
			ses_VerifyEmailIdentity(cfg, client)
			return
		}

	},
}

var (
	_sesCloneReceiptRuleSet                            bool
	_sesCreateConfigurationSet                         bool
	_sesCreateConfigurationSetEventDestination         bool
	_sesCreateConfigurationSetTrackingOptions          bool
	_sesCreateCustomVerificationEmailTemplate          bool
	_sesCreateReceiptFilter                            bool
	_sesCreateReceiptRule                              bool
	_sesCreateReceiptRuleSet                           bool
	_sesCreateTemplate                                 bool
	_sesDeleteConfigurationSet                         bool
	_sesDeleteConfigurationSetEventDestination         bool
	_sesDeleteConfigurationSetTrackingOptions          bool
	_sesDeleteCustomVerificationEmailTemplate          bool
	_sesDeleteIdentity                                 bool
	_sesDeleteIdentityPolicy                           bool
	_sesDeleteReceiptFilter                            bool
	_sesDeleteReceiptRule                              bool
	_sesDeleteReceiptRuleSet                           bool
	_sesDeleteTemplate                                 bool
	_sesDeleteVerifiedEmailAddress                     bool
	_sesDescribeActiveReceiptRuleSet                   bool
	_sesDescribeConfigurationSet                       bool
	_sesDescribeReceiptRule                            bool
	_sesDescribeReceiptRuleSet                         bool
	_sesGetAccountSendingEnabled                       bool
	_sesGetCustomVerificationEmailTemplate             bool
	_sesGetIdentityDkimAttributes                      bool
	_sesGetIdentityMailFromDomainAttributes            bool
	_sesGetIdentityNotificationAttributes              bool
	_sesGetIdentityPolicies                            bool
	_sesGetIdentityVerificationAttributes              bool
	_sesGetSendQuota                                   bool
	_sesGetSendStatistics                              bool
	_sesGetTemplate                                    bool
	_sesListConfigurationSets                          bool
	_sesListCustomVerificationEmailTemplates           bool
	_sesListIdentities                                 bool
	_sesListIdentityPolicies                           bool
	_sesListReceiptFilters                             bool
	_sesListReceiptRuleSets                            bool
	_sesListTemplates                                  bool
	_sesListVerifiedEmailAddresses                     bool
	_sesPutConfigurationSetDeliveryOptions             bool
	_sesPutIdentityPolicy                              bool
	_sesReorderReceiptRuleSet                          bool
	_sesSendBounce                                     bool
	_sesSendBulkTemplatedEmail                         bool
	_sesSendCustomVerificationEmail                    bool
	_sesSendEmail                                      bool
	_sesSendRawEmail                                   bool
	_sesSendTemplatedEmail                             bool
	_sesSetActiveReceiptRuleSet                        bool
	_sesSetIdentityDkimEnabled                         bool
	_sesSetIdentityFeedbackForwardingEnabled           bool
	_sesSetIdentityHeadersInNotificationsEnabled       bool
	_sesSetIdentityMailFromDomain                      bool
	_sesSetIdentityNotificationTopic                   bool
	_sesSetReceiptRulePosition                         bool
	_sesTestRenderTemplate                             bool
	_sesUpdateAccountSendingEnabled                    bool
	_sesUpdateConfigurationSetEventDestination         bool
	_sesUpdateConfigurationSetReputationMetricsEnabled bool
	_sesUpdateConfigurationSetSendingEnabled           bool
	_sesUpdateConfigurationSetTrackingOptions          bool
	_sesUpdateCustomVerificationEmailTemplate          bool
	_sesUpdateReceiptRule                              bool
	_sesUpdateTemplate                                 bool
	_sesVerifyDomainDkim                               bool
	_sesVerifyDomainIdentity                           bool
	_sesVerifyEmailAddress                             bool
	_sesVerifyEmailIdentity                            bool

	_sesAfter                          string
	_sesBehaviorOnMXFailure            string
	_sesBounceSender                   string
	_sesBounceSenderArn                string
	_sesBouncedRecipientInfoList       string
	_sesConfigurationSet               string
	_sesConfigurationSetAttributeNames string
	_sesConfigurationSetName           string
	_sesDefaultTags                    string
	_sesDefaultTemplateData            string
	_sesDeliveryOptions                string
	_sesDestination                    string
	_sesDestinations                   []string
	_sesDkimEnabled                    string
	_sesDomain                         string
	_sesEmailAddress                   string
	_sesEnabled                        string
	_sesEventDestination               string
	_sesEventDestinationName           string
	_sesExplanation                    string
	_sesFailureRedirectionURL          string
	_sesFilter                         string
	_sesFilterName                     string
	_sesForwardingEnabled              string
	_sesFromArn                        string
	_sesFromEmailAddress               string
	_sesIdentities                     []string
	_sesIdentity                       string
	_sesIdentityType                   string
	_sesMailFromDomain                 string
	_sesMaxItems                       string
	_sesMaxResults                     string
	_sesMessage                        string
	_sesMessageDsn                     string
	_sesNextToken                      string
	_sesNotificationType               string
	_sesOriginalMessageId              string
	_sesOriginalRuleSetName            string
	_sesPolicy                         string
	_sesPolicyName                     string
	_sesPolicyNames                    []string
	_sesRawMessage                     string
	_sesReplyToAddresses               []string
	_sesReturnPath                     string
	_sesReturnPathArn                  string
	_sesRule                           string
	_sesRuleName                       string
	_sesRuleNames                      []string
	_sesRuleSetName                    string
	_sesSnsTopic                       string
	_sesSource                         string
	_sesSourceArn                      string
	_sesSuccessRedirectionURL          string
	_sesTags                           string
	_sesTemplate                       string
	_sesTemplateArn                    string
	_sesTemplateContent                string
	_sesTemplateData                   string
	_sesTemplateName                   string
	_sesTemplateSubject                string
	_sesTrackingOptions                string
)

// Creates a receipt rule set by cloning an existing one. All receipt rules and
// configurations are copied to the new receipt rule set and are completely
// independent of the source rule set.
//
// For information about setting up rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html#receiving-email-concepts-rules
func ses_CloneReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.CloneReceiptRuleSetInput{
		// OriginalRuleSetName: *string, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesOriginalRuleSetName) > 0 {
		input.OriginalRuleSetName = aws.String(_sesOriginalRuleSetName)
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.CloneReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a configuration set.
// Configuration sets enable you to publish email sending events. For information
// about using configuration sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_CreateConfigurationSet(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateConfigurationSetInput{
		// ConfigurationSet: *types.ConfigurationSet, // Required
	}

	if len(_sesConfigurationSet) > 0 {
		if err := assignInputField(input, "ConfigurationSet", _sesConfigurationSet); err != nil {
			log.Errorf("invalid --configuration-set: %s", err.Error())
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

// Creates a configuration set event destination.
// When you create or update an event destination, you must provide one, and only
// one, destination. The destination can be CloudWatch, Amazon Kinesis Firehose, or
// Amazon Simple Notification Service (Amazon SNS).
//
// An event destination is the Amazon Web Services service to which Amazon SES
// publishes the email sending events associated with a configuration set. For
// information about using configuration sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_CreateConfigurationSetEventDestination(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestination: *types.EventDestination, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesEventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _sesEventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between a configuration set and a custom domain for open
// and click event tracking.
//
// By default, images and links used for tracking open and click events are hosted
// on domains operated by Amazon SES. You can configure a subdomain of your own to
// handle these events. For information about using custom domains, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html
func ses_CreateConfigurationSetTrackingOptions(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateConfigurationSetTrackingOptionsInput{
		// ConfigurationSetName: *string, // Required
		// TrackingOptions: *types.TrackingOptions, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesTrackingOptions) > 0 {
		if err := assignInputField(input, "TrackingOptions", _sesTrackingOptions); err != nil {
			log.Errorf("invalid --tracking-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateConfigurationSetTrackingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom verification email template.
// For more information about custom verification email templates, see [Using Custom Verification Email Templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func ses_CreateCustomVerificationEmailTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateCustomVerificationEmailTemplateInput{
		// FailureRedirectionURL: *string, // Required
		// FromEmailAddress: *string, // Required
		// SuccessRedirectionURL: *string, // Required
		// TemplateContent: *string, // Required
		// TemplateName: *string, // Required
		// TemplateSubject: *string, // Required
	}

	if len(_sesFailureRedirectionURL) > 0 {
		input.FailureRedirectionURL = aws.String(_sesFailureRedirectionURL)
	}
	if len(_sesFromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesFromEmailAddress)
	}
	if len(_sesSuccessRedirectionURL) > 0 {
		input.SuccessRedirectionURL = aws.String(_sesSuccessRedirectionURL)
	}
	if len(_sesTemplateContent) > 0 {
		input.TemplateContent = aws.String(_sesTemplateContent)
	}
	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}
	if len(_sesTemplateSubject) > 0 {
		input.TemplateSubject = aws.String(_sesTemplateSubject)
	}

	if resp, err := client.CreateCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new IP address filter.
// For information about setting up IP address filters, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-ip-filtering-console-walkthrough.html
func ses_CreateReceiptFilter(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateReceiptFilterInput{
		// Filter: *types.ReceiptFilter, // Required
	}

	if len(_sesFilter) > 0 {
		if err := assignInputField(input, "Filter", _sesFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateReceiptFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a receipt rule.
// For information about setting up receipt rules, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_CreateReceiptRule(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateReceiptRuleInput{
		// Rule: *types.ReceiptRule, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesRule) > 0 {
		if err := assignInputField(input, "Rule", _sesRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}
	if len(_sesAfter) > 0 {
		input.After = aws.String(_sesAfter)
	}

	if resp, err := client.CreateReceiptRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an empty receipt rule set.
// For information about setting up receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html#receiving-email-concepts-rules
func ses_CreateReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateReceiptRuleSetInput{
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.CreateReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an email template. Email templates enable you to send personalized
// email to one or more destinations in a single operation. For more information,
// see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-personalized-email-api.html
func ses_CreateTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.CreateTemplateInput{
		// Template: *types.Template, // Required
	}

	if len(_sesTemplate) > 0 {
		if err := assignInputField(input, "Template", _sesTemplate); err != nil {
			log.Errorf("invalid --template: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration set. Configuration sets enable you to publish email
// sending events. For information about using configuration sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_DeleteConfigurationSet(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}

	if resp, err := client.DeleteConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a configuration set event destination. Configuration set event
// destinations are associated with configuration sets, which enable you to publish
// email sending events. For information about using configuration sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_DeleteConfigurationSetEventDestination(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesEventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_sesEventDestinationName)
	}

	if resp, err := client.DeleteConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an association between a configuration set and a custom domain for open
// and click event tracking.
//
// By default, images and links used for tracking open and click events are hosted
// on domains operated by Amazon SES. You can configure a subdomain of your own to
// handle these events. For information about using custom domains, see the [Amazon SES Developer Guide].
//
// Deleting this kind of association results in emails sent using the specified
// configuration set to capture open and click events using the standard, Amazon
// SES-operated domains.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html
func ses_DeleteConfigurationSetTrackingOptions(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteConfigurationSetTrackingOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}

	if resp, err := client.DeleteConfigurationSetTrackingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing custom verification email template.
// For more information about custom verification email templates, see [Using Custom Verification Email Templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func ses_DeleteCustomVerificationEmailTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteCustomVerificationEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}

	if resp, err := client.DeleteCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified identity (an email address or a domain) from the list of
// verified identities.
//
// You can execute this operation no more than once per second.
func ses_DeleteIdentity(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteIdentityInput{
		// Identity: *string, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}

	if resp, err := client.DeleteIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified sending authorization policy for the given identity (an
// email address or a domain). This operation returns successfully even if a policy
// with the specified name does not exist.
//
// This operation is for the identity owner only. If you have not verified the
// identity, it returns an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
func ses_DeleteIdentityPolicy(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteIdentityPolicyInput{
		// Identity: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}
	if len(_sesPolicyName) > 0 {
		input.PolicyName = aws.String(_sesPolicyName)
	}

	if resp, err := client.DeleteIdentityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified IP address filter.
// For information about managing IP address filters, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-ip-filtering-console-walkthrough.html
func ses_DeleteReceiptFilter(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteReceiptFilterInput{
		// FilterName: *string, // Required
	}

	if len(_sesFilterName) > 0 {
		input.FilterName = aws.String(_sesFilterName)
	}

	if resp, err := client.DeleteReceiptFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified receipt rule.
// For information about managing receipt rules, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_DeleteReceiptRule(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteReceiptRuleInput{
		// RuleName: *string, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleName) > 0 {
		input.RuleName = aws.String(_sesRuleName)
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.DeleteReceiptRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified receipt rule set and all of the receipt rules it contains.
// The currently active rule set cannot be deleted.
//
// For information about managing receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_DeleteReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteReceiptRuleSetInput{
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.DeleteReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an email template.
// You can execute this operation no more than once per second.
func ses_DeleteTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}

	if resp, err := client.DeleteTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use the DeleteIdentity operation to delete email addresses and
// domains.
func ses_DeleteVerifiedEmailAddress(cfg aws.Config, client *ses.Client) {
	input := &ses.DeleteVerifiedEmailAddressInput{
		// EmailAddress: *string, // Required
	}

	if len(_sesEmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesEmailAddress)
	}

	if resp, err := client.DeleteVerifiedEmailAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the metadata and receipt rules for the receipt rule set that is
// currently active.
//
// For information about setting up receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-concepts.html#receiving-email-concepts-rules
func ses_DescribeActiveReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.DescribeActiveReceiptRuleSetInput{}

	if resp, err := client.DescribeActiveReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the specified configuration set. For information about
// using configuration sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_DescribeConfigurationSet(cfg aws.Config, client *ses.Client) {
	input := &ses.DescribeConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesConfigurationSetAttributeNames) > 0 {
		if err := assignInputField(input, "ConfigurationSetAttributeNames", _sesConfigurationSetAttributeNames); err != nil {
			log.Errorf("invalid --configuration-set-attribute-names: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the specified receipt rule.
// For information about setting up receipt rules, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_DescribeReceiptRule(cfg aws.Config, client *ses.Client) {
	input := &ses.DescribeReceiptRuleInput{
		// RuleName: *string, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleName) > 0 {
		input.RuleName = aws.String(_sesRuleName)
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.DescribeReceiptRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of the specified receipt rule set.
// For information about managing receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_DescribeReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.DescribeReceiptRuleSetInput{
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.DescribeReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the email sending status of the Amazon SES account for the current
// Region.
//
// You can execute this operation no more than once per second.
func ses_GetAccountSendingEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.GetAccountSendingEnabledInput{}

	if resp, err := client.GetAccountSendingEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the custom email verification template for the template name you
// specify.
//
// For more information about custom verification email templates, see [Using Custom Verification Email Templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func ses_GetCustomVerificationEmailTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.GetCustomVerificationEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}

	if resp, err := client.GetCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current status of Easy DKIM signing for an entity. For domain name
// identities, this operation also returns the DKIM tokens that are required for
// Easy DKIM signing, and whether Amazon SES has successfully verified that these
// tokens have been published.
//
// This operation takes a list of identities as input and returns the following
// information for each:
//
// - Whether Easy DKIM signing is enabled or disabled.
//
// - A set of DKIM tokens that represent the identity. If the identity is an
// email address, the tokens represent the domain of that address.
//
// - Whether Amazon SES has successfully verified the DKIM tokens published in
// the domain's DNS. This information is only returned for domain name identities,
// not for email addresses.
//
// This operation is throttled at one request per second and can only get DKIM
// attributes for up to 100 identities at a time.
//
// For more information about creating DNS records using DKIM tokens, go to the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy-managing.html
func ses_GetIdentityDkimAttributes(cfg aws.Config, client *ses.Client) {
	input := &ses.GetIdentityDkimAttributesInput{
		// Identities: []string, // Required
	}

	if len(_sesIdentities) > 0 {
		input.Identities = append([]string(nil), _sesIdentities...)
	}

	if resp, err := client.GetIdentityDkimAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the custom MAIL FROM attributes for a list of identities (email
// addresses : domains).
//
// This operation is throttled at one request per second and can only get custom
// MAIL FROM attributes for up to 100 identities at a time.
func ses_GetIdentityMailFromDomainAttributes(cfg aws.Config, client *ses.Client) {
	input := &ses.GetIdentityMailFromDomainAttributesInput{
		// Identities: []string, // Required
	}

	if len(_sesIdentities) > 0 {
		input.Identities = append([]string(nil), _sesIdentities...)
	}

	if resp, err := client.GetIdentityMailFromDomainAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a list of verified identities (email addresses and/or domains), returns a
// structure describing identity notification attributes.
//
// This operation is throttled at one request per second and can only get
// notification attributes for up to 100 identities at a time.
//
// For more information about using notifications with Amazon SES, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
func ses_GetIdentityNotificationAttributes(cfg aws.Config, client *ses.Client) {
	input := &ses.GetIdentityNotificationAttributesInput{
		// Identities: []string, // Required
	}

	if len(_sesIdentities) > 0 {
		input.Identities = append([]string(nil), _sesIdentities...)
	}

	if resp, err := client.GetIdentityNotificationAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the requested sending authorization policies for the given identity (an
// email address or a domain). The policies are returned as a map of policy names
// to policy contents. You can retrieve a maximum of 20 policies at a time.
//
// This operation is for the identity owner only. If you have not verified the
// identity, it returns an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
func ses_GetIdentityPolicies(cfg aws.Config, client *ses.Client) {
	input := &ses.GetIdentityPoliciesInput{
		// Identity: *string, // Required
		// PolicyNames: []string, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}
	if len(_sesPolicyNames) > 0 {
		input.PolicyNames = append([]string(nil), _sesPolicyNames...)
	}

	if resp, err := client.GetIdentityPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given a list of identities (email addresses and/or domains), returns the
// verification status and (for domain identities) the verification token for each
// identity.
//
// The verification status of an email address is "Pending" until the email
// address owner clicks the link within the verification email that Amazon SES sent
// to that address. If the email address owner clicks the link within 24 hours, the
// verification status of the email address changes to "Success". If the link is
// not clicked within 24 hours, the verification status changes to "Failed." In
// that case, to verify the email address, you must restart the verification
// process from the beginning.
//
// For domain identities, the domain's verification status is "Pending" as Amazon
// SES searches for the required TXT record in the DNS settings of the domain. When
// Amazon SES detects the record, the domain's verification status changes to
// "Success". If Amazon SES is unable to detect the record within 72 hours, the
// domain's verification status changes to "Failed." In that case, to verify the
// domain, you must restart the verification process from the beginning.
//
// This operation is throttled at one request per second and can only get
// verification attributes for up to 100 identities at a time.
func ses_GetIdentityVerificationAttributes(cfg aws.Config, client *ses.Client) {
	input := &ses.GetIdentityVerificationAttributesInput{
		// Identities: []string, // Required
	}

	if len(_sesIdentities) > 0 {
		input.Identities = append([]string(nil), _sesIdentities...)
	}

	if resp, err := client.GetIdentityVerificationAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides the sending limits for the Amazon SES account.
// You can execute this operation no more than once per second.
func ses_GetSendQuota(cfg aws.Config, client *ses.Client) {
	input := &ses.GetSendQuotaInput{}

	if resp, err := client.GetSendQuota(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides sending statistics for the current Amazon Web Services Region. The
// result is a list of data points, representing the last two weeks of sending
// activity. Each data point in the list contains statistics for a 15-minute period
// of time.
//
// You can execute this operation no more than once per second.
func ses_GetSendStatistics(cfg aws.Config, client *ses.Client) {
	input := &ses.GetSendStatisticsInput{}

	if resp, err := client.GetSendStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the template object (which includes the Subject line, HTML part and
// text part) for the template you specify.
//
// You can execute this operation no more than once per second.
func ses_GetTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.GetTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}

	if resp, err := client.GetTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides a list of the configuration sets associated with your Amazon SES
// account in the current Amazon Web Services Region. For information about using
// configuration sets, see [Monitoring Your Amazon SES Sending Activity]in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second. This operation
// returns up to 1,000 configuration sets each time it is run. If your Amazon SES
// account has more than 1,000 configuration sets, this operation also returns
// NextToken . You can then execute the ListConfigurationSets operation again,
// passing the NextToken parameter and the value of the NextToken element to
// retrieve additional results.
//
// [Monitoring Your Amazon SES Sending Activity]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_ListConfigurationSets(cfg aws.Config, client *ses.Client) {
	input := &ses.ListConfigurationSetsInput{}

	if len(_sesMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _sesMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_sesNextToken) > 0 {
		input.NextToken = aws.String(_sesNextToken)
	}

	if resp, err := client.ListConfigurationSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the existing custom verification email templates for your account in the
// current Amazon Web Services Region.
//
// For more information about custom verification email templates, see [Using Custom Verification Email Templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func ses_ListCustomVerificationEmailTemplates(cfg aws.Config, client *ses.Client) {
	input := &ses.ListCustomVerificationEmailTemplatesInput{}

	if len(_sesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _sesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_sesNextToken) > 0 {
		input.NextToken = aws.String(_sesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomVerificationEmailTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ses.ListCustomVerificationEmailTemplatesOutput
	p := ses.NewListCustomVerificationEmailTemplatesPaginator(client, input)
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

// Returns a list containing all of the identities (email addresses and domains)
// for your Amazon Web Services account in the current Amazon Web Services Region,
// regardless of verification status.
//
// You can execute this operation no more than once per second.
//
// It's recommended that for successive pagination calls of this API, you continue
// to the use the same parameter/value pairs as used in the original call, e.g., if
// you used IdentityType=Domain in the the original call and received a NextToken
// in the response, you should continue providing the IdentityType=Domain
// parameter for further NextToken calls; however, if you didn't provide the
// IdentityType parameter in the original call, then continue to not provide it for
// successive pagination calls. Using this protocol will ensure consistent results.
func ses_ListIdentities(cfg aws.Config, client *ses.Client) {
	input := &ses.ListIdentitiesInput{}

	if len(_sesIdentityType) > 0 {
		if err := assignInputField(input, "IdentityType", _sesIdentityType); err != nil {
			log.Errorf("invalid --identity-type: %s", err.Error())
			return
		}
	}
	if len(_sesMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _sesMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_sesNextToken) > 0 {
		input.NextToken = aws.String(_sesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ses.ListIdentitiesOutput
	p := ses.NewListIdentitiesPaginator(client, input)
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

// Returns a list of sending authorization policies that are attached to the given
// identity (an email address or a domain). This operation returns only a list. To
// get the actual policy content, use GetIdentityPolicies .
//
// This operation is for the identity owner only. If you have not verified the
// identity, it returns an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
func ses_ListIdentityPolicies(cfg aws.Config, client *ses.Client) {
	input := &ses.ListIdentityPoliciesInput{
		// Identity: *string, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}

	if resp, err := client.ListIdentityPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the IP address filters associated with your Amazon Web Services account
// in the current Amazon Web Services Region.
//
// For information about managing IP address filters, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-ip-filtering-console-walkthrough.html
func ses_ListReceiptFilters(cfg aws.Config, client *ses.Client) {
	input := &ses.ListReceiptFiltersInput{}

	if resp, err := client.ListReceiptFilters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the receipt rule sets that exist under your Amazon Web Services account
// in the current Amazon Web Services Region. If there are additional receipt rule
// sets to be retrieved, you receive a NextToken that you can provide to the next
// call to ListReceiptRuleSets to retrieve the additional entries.
//
// For information about managing receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_ListReceiptRuleSets(cfg aws.Config, client *ses.Client) {
	input := &ses.ListReceiptRuleSetsInput{}

	if len(_sesNextToken) > 0 {
		input.NextToken = aws.String(_sesNextToken)
	}

	if resp, err := client.ListReceiptRuleSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the email templates present in your Amazon SES account in the current
// Amazon Web Services Region.
//
// You can execute this operation no more than once per second.
func ses_ListTemplates(cfg aws.Config, client *ses.Client) {
	input := &ses.ListTemplatesInput{}

	if len(_sesMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _sesMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_sesNextToken) > 0 {
		input.NextToken = aws.String(_sesNextToken)
	}

	if resp, err := client.ListTemplates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use the ListIdentities operation to list the email addresses and
// domains associated with your account.
func ses_ListVerifiedEmailAddresses(cfg aws.Config, client *ses.Client) {
	input := &ses.ListVerifiedEmailAddressesInput{}

	if resp, err := client.ListVerifiedEmailAddresses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the delivery options for a configuration set.
func ses_PutConfigurationSetDeliveryOptions(cfg aws.Config, client *ses.Client) {
	input := &ses.PutConfigurationSetDeliveryOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesDeliveryOptions) > 0 {
		if err := assignInputField(input, "DeliveryOptions", _sesDeliveryOptions); err != nil {
			log.Errorf("invalid --delivery-options: %s", err.Error())
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

// Adds or updates a sending authorization policy for the specified identity (an
// email address or a domain).
//
// This operation is for the identity owner only. If you have not verified the
// identity, it returns an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
func ses_PutIdentityPolicy(cfg aws.Config, client *ses.Client) {
	input := &ses.PutIdentityPolicyInput{
		// Identity: *string, // Required
		// Policy: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}
	if len(_sesPolicy) > 0 {
		input.Policy = aws.String(_sesPolicy)
	}
	if len(_sesPolicyName) > 0 {
		input.PolicyName = aws.String(_sesPolicyName)
	}

	if resp, err := client.PutIdentityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reorders the receipt rules within a receipt rule set.
// All of the rules in the rule set must be represented in this request. That is,
// it is error if the reorder request doesn't explicitly position all of the rules.
//
// For information about managing receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_ReorderReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.ReorderReceiptRuleSetInput{
		// RuleNames: []string, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleNames) > 0 {
		input.RuleNames = append([]string(nil), _sesRuleNames...)
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.ReorderReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates and sends a bounce message to the sender of an email you received
// through Amazon SES. You can only use this operation on an email up to 24 hours
// after you receive it.
//
// You cannot use this operation to send generic bounces for mail that was not
// received by Amazon SES.
//
// For information about receiving email through Amazon SES, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email.html
func ses_SendBounce(cfg aws.Config, client *ses.Client) {
	input := &ses.SendBounceInput{
		// BounceSender: *string, // Required
		// BouncedRecipientInfoList: []types.BouncedRecipientInfo, // Required
		// OriginalMessageId: *string, // Required
	}

	if len(_sesBounceSender) > 0 {
		input.BounceSender = aws.String(_sesBounceSender)
	}
	if len(_sesBouncedRecipientInfoList) > 0 {
		if err := assignInputField(input, "BouncedRecipientInfoList", _sesBouncedRecipientInfoList); err != nil {
			log.Errorf("invalid --bounced-recipient-info-list: %s", err.Error())
			return
		}
	}
	if len(_sesOriginalMessageId) > 0 {
		input.OriginalMessageId = aws.String(_sesOriginalMessageId)
	}
	if len(_sesBounceSenderArn) > 0 {
		input.BounceSenderArn = aws.String(_sesBounceSenderArn)
	}
	if len(_sesExplanation) > 0 {
		input.Explanation = aws.String(_sesExplanation)
	}
	if len(_sesMessageDsn) > 0 {
		if err := assignInputField(input, "MessageDsn", _sesMessageDsn); err != nil {
			log.Errorf("invalid --message-dsn: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendBounce(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Composes an email message to multiple destinations. The message body is created
// using an email template.
//
// To send email using this operation, your call must meet the following
// requirements:
//
// - The call must refer to an existing email template. You can create email
// templates using CreateTemplate.
//
// - The message must be sent from a verified email address or domain.
//
// - If your account is still in the Amazon SES sandbox, you may send only to
// verified addresses or domains, or to email addresses associated with the Amazon
// SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
// Guide.
//
// - The maximum message size is 10 MB.
//
// - Each Destination parameter must include at least one recipient email
// address. The recipient address can be a To: address, a CC: address, or a BCC:
// address. If a recipient email address is invalid (that is, it is not in the
// format UserName(at)[SubDomain.]Domain.TopLevelDomain), the entire message is
// rejected, even if the message contains other recipients that are valid.
//
// - The message may not include more than 50 recipients, across the To:, CC:
// and BCC: fields. If you need to send an email message to a larger audience, you
// can divide your recipient list into groups of 50 or fewer, and then call the
// SendBulkTemplatedEmail operation several times to send the message to each
// group.
//
// - The number of destinations you can contact in a single call can be limited
// by your account's maximum sending rate.
//
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
func ses_SendBulkTemplatedEmail(cfg aws.Config, client *ses.Client) {
	input := &ses.SendBulkTemplatedEmailInput{
		// DefaultTemplateData: *string, // Required
		// Destinations: []types.BulkEmailDestination, // Required
		// Source: *string, // Required
		// Template: *string, // Required
	}

	if len(_sesDefaultTemplateData) > 0 {
		input.DefaultTemplateData = aws.String(_sesDefaultTemplateData)
	}
	if len(_sesDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _sesDestinations[0]); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_sesSource) > 0 {
		input.Source = aws.String(_sesSource)
	}
	if len(_sesTemplate) > 0 {
		input.Template = aws.String(_sesTemplate)
	}
	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesDefaultTags) > 0 {
		if err := assignInputField(input, "DefaultTags", _sesDefaultTags); err != nil {
			log.Errorf("invalid --default-tags: %s", err.Error())
			return
		}
	}
	if len(_sesReplyToAddresses) > 0 {
		input.ReplyToAddresses = append([]string(nil), _sesReplyToAddresses...)
	}
	if len(_sesReturnPath) > 0 {
		input.ReturnPath = aws.String(_sesReturnPath)
	}
	if len(_sesReturnPathArn) > 0 {
		input.ReturnPathArn = aws.String(_sesReturnPathArn)
	}
	if len(_sesSourceArn) > 0 {
		input.SourceArn = aws.String(_sesSourceArn)
	}
	if len(_sesTemplateArn) > 0 {
		input.TemplateArn = aws.String(_sesTemplateArn)
	}

	if resp, err := client.SendBulkTemplatedEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an email address to the list of identities for your Amazon SES account in
// the current Amazon Web Services Region and attempts to verify it. As a result of
// executing this operation, a customized verification email is sent to the
// specified address.
//
// To use this operation, you must first create a custom verification email
// template. For more information about creating and using custom verification
// email templates, see [Using Custom Verification Email Templates]in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func ses_SendCustomVerificationEmail(cfg aws.Config, client *ses.Client) {
	input := &ses.SendCustomVerificationEmailInput{
		// EmailAddress: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_sesEmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesEmailAddress)
	}
	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}
	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}

	if resp, err := client.SendCustomVerificationEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Composes an email message and immediately queues it for sending. To send email
// using this operation, your message must meet the following requirements:
//
// - The message must be sent from a verified email address or domain. If you
// attempt to send email using a non-verified address or domain, the operation
// results in an "Email address not verified" error.
//
// - If your account is still in the Amazon SES sandbox, you may only send to
// verified addresses or domains, or to email addresses associated with the Amazon
// SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
// Guide.
//
// - The maximum message size is 10 MB.
//
// - The message must include at least one recipient email address. The
// recipient address can be a To: address, a CC: address, or a BCC: address. If a
// recipient email address is invalid (that is, it is not in the format
// UserName(at)[SubDomain.]Domain.TopLevelDomain), the entire message is rejected,
// even if the message contains other recipients that are valid.
//
// - The message may not include more than 50 recipients, across the To:, CC:
// and BCC: fields. If you need to send an email message to a larger audience, you
// can divide your recipient list into groups of 50 or fewer, and then call the
// SendEmail operation several times to send the message to each group.
//
// For every message that you send, the total number of recipients (including each
// recipient in the To:, CC: and BCC: fields) is counted against the maximum number
// of emails you can send in a 24-hour period (your sending quota). For more
// information about sending quotas in Amazon SES, see [Managing Your Amazon SES Sending Limits]in the Amazon SES Developer
// Guide.
//
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
// [Managing Your Amazon SES Sending Limits]: https://docs.aws.amazon.com/ses/latest/dg/manage-sending-quotas.html
func ses_SendEmail(cfg aws.Config, client *ses.Client) {
	input := &ses.SendEmailInput{
		// Destination: *types.Destination, // Required
		// Message: *types.Message, // Required
		// Source: *string, // Required
	}

	if len(_sesDestination) > 0 {
		if err := assignInputField(input, "Destination", _sesDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_sesMessage) > 0 {
		if err := assignInputField(input, "Message", _sesMessage); err != nil {
			log.Errorf("invalid --message: %s", err.Error())
			return
		}
	}
	if len(_sesSource) > 0 {
		input.Source = aws.String(_sesSource)
	}
	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesReplyToAddresses) > 0 {
		input.ReplyToAddresses = append([]string(nil), _sesReplyToAddresses...)
	}
	if len(_sesReturnPath) > 0 {
		input.ReturnPath = aws.String(_sesReturnPath)
	}
	if len(_sesReturnPathArn) > 0 {
		input.ReturnPathArn = aws.String(_sesReturnPathArn)
	}
	if len(_sesSourceArn) > 0 {
		input.SourceArn = aws.String(_sesSourceArn)
	}
	if len(_sesTags) > 0 {
		if err := assignInputField(input, "Tags", _sesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Composes an email message and immediately queues it for sending.
// This operation is more flexible than the SendEmail operation. When you use the
// SendRawEmail operation, you can specify the headers of the message as well as
// its content. This flexibility is useful, for example, when you need to send a
// multipart MIME email (such a message that contains both a text and an HTML
// version). You can also use this operation to send messages that include
// attachments.
//
// The SendRawEmail operation has the following requirements:
//
// - You can only send email from [verified email addresses or domains]. If you try to send email from an address
// that isn't verified, the operation results in an "Email address not verified"
// error.
//
// - If your account is still in the [Amazon SES sandbox], you can only send email to other verified
// addresses in your account, or to addresses that are associated with the [Amazon SES mailbox simulator].
//
// - The maximum message size, including attachments, is 10 MB.
//
// - Each message has to include at least one recipient address. A recipient
// address includes any address on the To:, CC:, or BCC: lines.
//
// - If you send a single message to more than one recipient address, and one of
// the recipient addresses isn't in a valid format (that is, it's not in the format
// UserName(at)[SubDomain.]Domain.TopLevelDomain), Amazon SES rejects the entire
// message, even if the other addresses are valid.
//
// - Each message can include up to 50 recipient addresses across the To:, CC:,
// or BCC: lines. If you need to send a single message to more than 50 recipients,
// you have to split the list of recipient addresses into groups of less than 50
// recipients, and send separate messages to each group.
//
// - Amazon SES allows you to specify 8-bit Content-Transfer-Encoding for MIME
// message parts. However, if Amazon SES has to modify the contents of your message
// (for example, if you use open and click tracking), 8-bit content isn't
// preserved. For this reason, we highly recommend that you encode all content that
// isn't 7-bit ASCII. For more information, see [MIME Encoding]in the Amazon SES Developer
// Guide.
//
// Additionally, keep the following considerations in mind when using the
// SendRawEmail operation:
//
// - Although you can customize the message headers when using the SendRawEmail
// operation, Amazon SES automatically applies its own Message-ID and Date
// headers; if you passed these headers when creating the message, they are
// overwritten by the values that Amazon SES provides.
//
// - If you are using sending authorization to send on behalf of another user,
// SendRawEmail enables you to specify the cross-account identity for the email's
// Source, From, and Return-Path parameters in one of two ways: you can pass
// optional parameters SourceArn , FromArn , and/or ReturnPathArn , or you can
// include the following X-headers in the header of your raw email:
//
// - X-SES-SOURCE-ARN
//
// - X-SES-FROM-ARN
//
// - X-SES-RETURN-PATH-ARN
//
// Don't include these X-headers in the DKIM signature. Amazon SES removes these
//
// before it sends the email.
//
// # If you only specify the SourceIdentityArn parameter, Amazon SES sets the From
//
// and Return-Path addresses to the same identity that you specified.
//
// For more information about sending authorization, see the [Using Sending Authorization with Amazon SES]in the Amazon SES
//
// Developer Guide.
//
// - For every message that you send, the total number of recipients (including
// each recipient in the To:, CC: and BCC: fields) is counted against the maximum
// number of emails you can send in a 24-hour period (your sending quota). For more
// information about sending quotas in Amazon SES, see [Managing Your Amazon SES Sending Limits]in the Amazon SES
// Developer Guide.
//
// [Using Sending Authorization with Amazon SES]: https://docs.aws.amazon.com/ses/latest/dg/sending-authorization.html
// [MIME Encoding]: https://docs.aws.amazon.com/ses/latest/dg/send-email-raw.html#send-email-mime-encoding
// [Amazon SES mailbox simulator]: https://docs.aws.amazon.com/ses/latest/dg/send-an-email-from-console.html
// [Amazon SES sandbox]: https://docs.aws.amazon.com/ses/latest/dg/request-production-access.html
// [verified email addresses or domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
// [Managing Your Amazon SES Sending Limits]: https://docs.aws.amazon.com/ses/latest/dg/manage-sending-quotas.html
func ses_SendRawEmail(cfg aws.Config, client *ses.Client) {
	input := &ses.SendRawEmailInput{
		// RawMessage: *types.RawMessage, // Required
	}

	if len(_sesRawMessage) > 0 {
		if err := assignInputField(input, "RawMessage", _sesRawMessage); err != nil {
			log.Errorf("invalid --raw-message: %s", err.Error())
			return
		}
	}
	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesDestinations) > 0 {
		input.Destinations = append([]string(nil), _sesDestinations...)
	}
	if len(_sesFromArn) > 0 {
		input.FromArn = aws.String(_sesFromArn)
	}
	if len(_sesReturnPathArn) > 0 {
		input.ReturnPathArn = aws.String(_sesReturnPathArn)
	}
	if len(_sesSource) > 0 {
		input.Source = aws.String(_sesSource)
	}
	if len(_sesSourceArn) > 0 {
		input.SourceArn = aws.String(_sesSourceArn)
	}
	if len(_sesTags) > 0 {
		if err := assignInputField(input, "Tags", _sesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendRawEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Composes an email message using an email template and immediately queues it for
// sending.
//
// To send email using this operation, your call must meet the following
// requirements:
//
// - The call must refer to an existing email template. You can create email
// templates using the CreateTemplateoperation.
//
// - The message must be sent from a verified email address or domain.
//
// - If your account is still in the Amazon SES sandbox, you may only send to
// verified addresses or domains, or to email addresses associated with the Amazon
// SES Mailbox Simulator. For more information, see [Verifying Email Addresses and Domains]in the Amazon SES Developer
// Guide.
//
// - The maximum message size is 10 MB.
//
// - Calls to the SendTemplatedEmail operation may only include one Destination
// parameter. A destination is a set of recipients that receives the same version
// of the email. The Destination parameter can include up to 50 recipients,
// across the To:, CC: and BCC: fields.
//
// - The Destination parameter must include at least one recipient email address.
// The recipient address can be a To: address, a CC: address, or a BCC: address. If
// a recipient email address is invalid (that is, it is not in the format
// UserName(at)[SubDomain.]Domain.TopLevelDomain), the entire message is rejected,
// even if the message contains other recipients that are valid.
//
// If your call to the SendTemplatedEmail operation includes all of the required
// parameters, Amazon SES accepts it and returns a Message ID. However, if Amazon
// SES can't render the email because the template contains errors, it doesn't send
// the email. Additionally, because it already accepted the message, Amazon SES
// doesn't return a message stating that it was unable to send the email.
//
// For these reasons, we highly recommend that you set up Amazon SES to send you
// notifications when Rendering Failure events occur. For more information, see [Sending Personalized Email Using the Amazon SES API]in
// the Amazon Simple Email Service Developer Guide.
//
// [Sending Personalized Email Using the Amazon SES API]: https://docs.aws.amazon.com/ses/latest/dg/send-personalized-email-api.html
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
func ses_SendTemplatedEmail(cfg aws.Config, client *ses.Client) {
	input := &ses.SendTemplatedEmailInput{
		// Destination: *types.Destination, // Required
		// Source: *string, // Required
		// Template: *string, // Required
		// TemplateData: *string, // Required
	}

	if len(_sesDestination) > 0 {
		if err := assignInputField(input, "Destination", _sesDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_sesSource) > 0 {
		input.Source = aws.String(_sesSource)
	}
	if len(_sesTemplate) > 0 {
		input.Template = aws.String(_sesTemplate)
	}
	if len(_sesTemplateData) > 0 {
		input.TemplateData = aws.String(_sesTemplateData)
	}
	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesReplyToAddresses) > 0 {
		input.ReplyToAddresses = append([]string(nil), _sesReplyToAddresses...)
	}
	if len(_sesReturnPath) > 0 {
		input.ReturnPath = aws.String(_sesReturnPath)
	}
	if len(_sesReturnPathArn) > 0 {
		input.ReturnPathArn = aws.String(_sesReturnPathArn)
	}
	if len(_sesSourceArn) > 0 {
		input.SourceArn = aws.String(_sesSourceArn)
	}
	if len(_sesTags) > 0 {
		if err := assignInputField(input, "Tags", _sesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sesTemplateArn) > 0 {
		input.TemplateArn = aws.String(_sesTemplateArn)
	}

	if resp, err := client.SendTemplatedEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the specified receipt rule set as the active receipt rule set.
// To disable your email-receiving through Amazon SES completely, you can call
// this operation with RuleSetName set to null.
//
// For information about managing receipt rule sets, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_SetActiveReceiptRuleSet(cfg aws.Config, client *ses.Client) {
	input := &ses.SetActiveReceiptRuleSetInput{}

	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.SetActiveReceiptRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables Easy DKIM signing of email sent from an identity. If Easy
// DKIM signing is enabled for a domain, then Amazon SES uses DKIM to sign all
// email that it sends from addresses on that domain. If Easy DKIM signing is
// enabled for an email address, then Amazon SES uses DKIM to sign all email it
// sends from that address.
//
// For email addresses (for example, user(at)example.com ), you can only enable DKIM
// signing if the corresponding domain (in this case, example.com ) has been set up
// to use Easy DKIM.
//
// You can enable DKIM signing for an identity at any time after you start the
// verification process for the identity, even if the verification process isn't
// complete.
//
// You can execute this operation no more than once per second.
//
// For more information about Easy DKIM signing, go to the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html
func ses_SetIdentityDkimEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.SetIdentityDkimEnabledInput{
		// DkimEnabled: bool, // Required
		// Identity: *string, // Required
	}

	if len(_sesDkimEnabled) > 0 {
		if err := assignInputField(input, "DkimEnabled", _sesDkimEnabled); err != nil {
			log.Errorf("invalid --dkim-enabled: %s", err.Error())
			return
		}
	}
	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}

	if resp, err := client.SetIdentityDkimEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given an identity (an email address or a domain), enables or disables whether
// Amazon SES forwards bounce and complaint notifications as email. Feedback
// forwarding can only be disabled when Amazon Simple Notification Service (Amazon
// SNS) topics are specified for both bounces and complaints.
//
// Feedback forwarding does not apply to delivery notifications. Delivery
// notifications are only available through Amazon SNS.
//
// You can execute this operation no more than once per second.
//
// For more information about using notifications with Amazon SES, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
func ses_SetIdentityFeedbackForwardingEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.SetIdentityFeedbackForwardingEnabledInput{
		// ForwardingEnabled: bool, // Required
		// Identity: *string, // Required
	}

	if len(_sesForwardingEnabled) > 0 {
		if err := assignInputField(input, "ForwardingEnabled", _sesForwardingEnabled); err != nil {
			log.Errorf("invalid --forwarding-enabled: %s", err.Error())
			return
		}
	}
	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}

	if resp, err := client.SetIdentityFeedbackForwardingEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Given an identity (an email address or a domain), sets whether Amazon SES
// includes the original email headers in the Amazon Simple Notification Service
// (Amazon SNS) notifications of a specified type.
//
// You can execute this operation no more than once per second.
//
// For more information about using notifications with Amazon SES, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
func ses_SetIdentityHeadersInNotificationsEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.SetIdentityHeadersInNotificationsEnabledInput{
		// Enabled: bool, // Required
		// Identity: *string, // Required
		// NotificationType: types.NotificationType, // Required
	}

	if len(_sesEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _sesEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}
	if len(_sesNotificationType) > 0 {
		if err := assignInputField(input, "NotificationType", _sesNotificationType); err != nil {
			log.Errorf("invalid --notification-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetIdentityHeadersInNotificationsEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables the custom MAIL FROM domain setup for a verified identity
// (an email address or a domain).
//
// To send emails using the specified MAIL FROM domain, you must add an MX record
// to your MAIL FROM domain's DNS settings. To ensure that your emails pass Sender
// Policy Framework (SPF) checks, you must also add or update an SPF record. For
// more information, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/mail-from.html
func ses_SetIdentityMailFromDomain(cfg aws.Config, client *ses.Client) {
	input := &ses.SetIdentityMailFromDomainInput{
		// Identity: *string, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}
	if len(_sesBehaviorOnMXFailure) > 0 {
		if err := assignInputField(input, "BehaviorOnMXFailure", _sesBehaviorOnMXFailure); err != nil {
			log.Errorf("invalid --behavior-on-mx-failure: %s", err.Error())
			return
		}
	}
	if len(_sesMailFromDomain) > 0 {
		input.MailFromDomain = aws.String(_sesMailFromDomain)
	}

	if resp, err := client.SetIdentityMailFromDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets an Amazon Simple Notification Service (Amazon SNS) topic to use when
// delivering notifications. When you use this operation, you specify a verified
// identity, such as an email address or domain. When you send an email that uses
// the chosen identity in the Source field, Amazon SES sends notifications to the
// topic you specified. You can send bounce, complaint, or delivery notifications
// (or any combination of the three) to the Amazon SNS topic that you specify.
//
// You can execute this operation no more than once per second.
//
// For more information about feedback notification, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity-using-notifications.html
func ses_SetIdentityNotificationTopic(cfg aws.Config, client *ses.Client) {
	input := &ses.SetIdentityNotificationTopicInput{
		// Identity: *string, // Required
		// NotificationType: types.NotificationType, // Required
	}

	if len(_sesIdentity) > 0 {
		input.Identity = aws.String(_sesIdentity)
	}
	if len(_sesNotificationType) > 0 {
		if err := assignInputField(input, "NotificationType", _sesNotificationType); err != nil {
			log.Errorf("invalid --notification-type: %s", err.Error())
			return
		}
	}
	if len(_sesSnsTopic) > 0 {
		input.SnsTopic = aws.String(_sesSnsTopic)
	}

	if resp, err := client.SetIdentityNotificationTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the position of the specified receipt rule in the receipt rule set.
// For information about managing receipt rules, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_SetReceiptRulePosition(cfg aws.Config, client *ses.Client) {
	input := &ses.SetReceiptRulePositionInput{
		// RuleName: *string, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesRuleName) > 0 {
		input.RuleName = aws.String(_sesRuleName)
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}
	if len(_sesAfter) > 0 {
		input.After = aws.String(_sesAfter)
	}

	if resp, err := client.SetReceiptRulePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a preview of the MIME content of an email when provided with a template
// and a set of replacement data.
//
// You can execute this operation no more than once per second.
func ses_TestRenderTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.TestRenderTemplateInput{
		// TemplateData: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_sesTemplateData) > 0 {
		input.TemplateData = aws.String(_sesTemplateData)
	}
	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}

	if resp, err := client.TestRenderTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables email sending across your entire Amazon SES account in the
// current Amazon Web Services Region. You can use this operation in conjunction
// with Amazon CloudWatch alarms to temporarily pause email sending across your
// Amazon SES account in a given Amazon Web Services Region when reputation metrics
// (such as your bounce or complaint rates) reach certain thresholds.
//
// You can execute this operation no more than once per second.
func ses_UpdateAccountSendingEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateAccountSendingEnabledInput{}

	if len(_sesEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _sesEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSendingEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the event destination of a configuration set. Event destinations are
// associated with configuration sets, which enable you to publish email sending
// events to Amazon CloudWatch, Amazon Kinesis Firehose, or Amazon Simple
// Notification Service (Amazon SNS). For information about using configuration
// sets, see [Monitoring Your Amazon SES Sending Activity]in the Amazon SES Developer Guide.
//
// When you create or update an event destination, you must provide one, and only
// one, destination. The destination can be Amazon CloudWatch, Amazon Kinesis
// Firehose, or Amazon Simple Notification Service (Amazon SNS).
//
// You can execute this operation no more than once per second.
//
// [Monitoring Your Amazon SES Sending Activity]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
func ses_UpdateConfigurationSetEventDestination(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestination: *types.EventDestination, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesEventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _sesEventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables the publishing of reputation metrics for emails sent using
// a specific configuration set in a given Amazon Web Services Region. Reputation
// metrics include bounce and complaint rates. These metrics are published to
// Amazon CloudWatch. By using CloudWatch, you can create alarms when bounce or
// complaint rates exceed certain thresholds.
//
// You can execute this operation no more than once per second.
func ses_UpdateConfigurationSetReputationMetricsEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateConfigurationSetReputationMetricsEnabledInput{
		// ConfigurationSetName: *string, // Required
		// Enabled: bool, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _sesEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationSetReputationMetricsEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables email sending for messages sent using a specific
// configuration set in a given Amazon Web Services Region. You can use this
// operation in conjunction with Amazon CloudWatch alarms to temporarily pause
// email sending for a configuration set when the reputation metrics for that
// configuration set (such as your bounce on complaint rate) exceed certain
// thresholds.
//
// You can execute this operation no more than once per second.
func ses_UpdateConfigurationSetSendingEnabled(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateConfigurationSetSendingEnabledInput{
		// ConfigurationSetName: *string, // Required
		// Enabled: bool, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _sesEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationSetSendingEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an association between a configuration set and a custom domain for
// open and click event tracking.
//
// By default, images and links used for tracking open and click events are hosted
// on domains operated by Amazon SES. You can configure a subdomain of your own to
// handle these events. For information about using custom domains, see the [Amazon SES Developer Guide].
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/configure-custom-open-click-domains.html
func ses_UpdateConfigurationSetTrackingOptions(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateConfigurationSetTrackingOptionsInput{
		// ConfigurationSetName: *string, // Required
		// TrackingOptions: *types.TrackingOptions, // Required
	}

	if len(_sesConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesConfigurationSetName)
	}
	if len(_sesTrackingOptions) > 0 {
		if err := assignInputField(input, "TrackingOptions", _sesTrackingOptions); err != nil {
			log.Errorf("invalid --tracking-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateConfigurationSetTrackingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing custom verification email template.
// For more information about custom verification email templates, see [Using Custom Verification Email Templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using Custom Verification Email Templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func ses_UpdateCustomVerificationEmailTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateCustomVerificationEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesTemplateName) > 0 {
		input.TemplateName = aws.String(_sesTemplateName)
	}
	if len(_sesFailureRedirectionURL) > 0 {
		input.FailureRedirectionURL = aws.String(_sesFailureRedirectionURL)
	}
	if len(_sesFromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesFromEmailAddress)
	}
	if len(_sesSuccessRedirectionURL) > 0 {
		input.SuccessRedirectionURL = aws.String(_sesSuccessRedirectionURL)
	}
	if len(_sesTemplateContent) > 0 {
		input.TemplateContent = aws.String(_sesTemplateContent)
	}
	if len(_sesTemplateSubject) > 0 {
		input.TemplateSubject = aws.String(_sesTemplateSubject)
	}

	if resp, err := client.UpdateCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a receipt rule.
// For information about managing receipt rules, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/receiving-email-receipt-rules-console-walkthrough.html
func ses_UpdateReceiptRule(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateReceiptRuleInput{
		// Rule: *types.ReceiptRule, // Required
		// RuleSetName: *string, // Required
	}

	if len(_sesRule) > 0 {
		if err := assignInputField(input, "Rule", _sesRule); err != nil {
			log.Errorf("invalid --rule: %s", err.Error())
			return
		}
	}
	if len(_sesRuleSetName) > 0 {
		input.RuleSetName = aws.String(_sesRuleSetName)
	}

	if resp, err := client.UpdateReceiptRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an email template. Email templates enable you to send personalized
// email to one or more destinations in a single operation. For more information,
// see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/dg/send-personalized-email-api.html
func ses_UpdateTemplate(cfg aws.Config, client *ses.Client) {
	input := &ses.UpdateTemplateInput{
		// Template: *types.Template, // Required
	}

	if len(_sesTemplate) > 0 {
		if err := assignInputField(input, "Template", _sesTemplate); err != nil {
			log.Errorf("invalid --template: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a set of DKIM tokens for a domain identity.
// When you execute the VerifyDomainDkim operation, the domain that you specify is
// added to the list of identities that are associated with your account. This is
// true even if you haven't already associated the domain with your account by
// using the VerifyDomainIdentity operation. However, you can't send email from
// the domain until you either successfully [verify it]or you successfully [set up DKIM for it].
//
// You use the tokens that are generated by this operation to create CNAME
// records. When Amazon SES detects that you've added these records to the DNS
// configuration for a domain, you can start sending email from that domain. You
// can start sending email even if you haven't added the TXT record provided by the
// VerifyDomainIdentity operation to the DNS configuration for your domain. All
// email that you send from the domain is authenticated using DKIM.
//
// To create the CNAME records for DKIM authentication, use the following values:
//
// - Name: token._domainkey.example.com
//
// - Type: CNAME
//
// - Value: token.dkim.amazonses.com
//
// In the preceding example, replace token with one of the tokens that are
// generated when you execute this operation. Replace example.com with your domain.
// Repeat this process for each token that's generated by this operation.
//
// You can execute this operation no more than once per second.
//
// [verify it]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#verify-domain-procedure
// [set up DKIM for it]: https://docs.aws.amazon.com/ses/latest/dg/send-email-authentication-dkim-easy.html
func ses_VerifyDomainDkim(cfg aws.Config, client *ses.Client) {
	input := &ses.VerifyDomainDkimInput{
		// Domain: *string, // Required
	}

	if len(_sesDomain) > 0 {
		input.Domain = aws.String(_sesDomain)
	}

	if resp, err := client.VerifyDomainDkim(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a domain to the list of identities for your Amazon SES account in the
// current Amazon Web Services Region and attempts to verify it. For more
// information about verifying domains, see [Verifying Email Addresses and Domains]in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Verifying Email Addresses and Domains]: https://docs.aws.amazon.com/ses/latest/dg/verify-addresses-and-domains.html
func ses_VerifyDomainIdentity(cfg aws.Config, client *ses.Client) {
	input := &ses.VerifyDomainIdentityInput{
		// Domain: *string, // Required
	}

	if len(_sesDomain) > 0 {
		input.Domain = aws.String(_sesDomain)
	}

	if resp, err := client.VerifyDomainIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecated. Use the VerifyEmailIdentity operation to verify a new email address.
func ses_VerifyEmailAddress(cfg aws.Config, client *ses.Client) {
	input := &ses.VerifyEmailAddressInput{
		// EmailAddress: *string, // Required
	}

	if len(_sesEmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesEmailAddress)
	}

	if resp, err := client.VerifyEmailAddress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an email address to the list of identities for your Amazon SES account in
// the current Amazon Web Services Region and attempts to verify it. As a result of
// executing this operation, a verification email is sent to the specified address.
//
// You can execute this operation no more than once per second.
func ses_VerifyEmailIdentity(cfg aws.Config, client *ses.Client) {
	input := &ses.VerifyEmailIdentityInput{
		// EmailAddress: *string, // Required
	}

	if len(_sesEmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesEmailAddress)
	}

	if resp, err := client.VerifyEmailIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sesCmd)
	_sesCmd.Flags().SortFlags = false

	_sesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_sesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_sesCmd.Flags().StringVarP(&_sesAfter, "after", "", "", "After")
	_sesCmd.Flags().StringVarP(&_sesBehaviorOnMXFailure, "behavior-on-mx-failure", "", "", "Behavior On Mx Failure")
	_sesCmd.Flags().StringVarP(&_sesBounceSender, "bounce-sender", "", "", "Bounce Sender")
	_sesCmd.Flags().StringVarP(&_sesBounceSenderArn, "bounce-sender-arn", "", "", "Bounce Sender ARN")
	_sesCmd.Flags().StringVarP(&_sesBouncedRecipientInfoList, "bounced-recipient-info-list", "", "", "Bounced Recipient Info List")
	_sesCmd.Flags().StringVarP(&_sesConfigurationSet, "configuration-set", "", "", "Configuration Set")
	_sesCmd.Flags().StringVarP(&_sesConfigurationSetAttributeNames, "configuration-set-attribute-names", "", "", "Configuration Set Attribute Names")
	_sesCmd.Flags().StringVarP(&_sesConfigurationSetName, "configuration-set-name", "", "", "Configuration Set Name")
	_sesCmd.Flags().StringVarP(&_sesDefaultTags, "default-tags", "", "", "Default Tags")
	_sesCmd.Flags().StringVarP(&_sesDefaultTemplateData, "default-template-data", "", "", "Default Template Data")
	_sesCmd.Flags().StringVarP(&_sesDeliveryOptions, "delivery-options", "", "", "Delivery Options")
	_sesCmd.Flags().StringVarP(&_sesDestination, "destination", "", "", "Destination")
	_sesCmd.Flags().StringSliceVarP(&_sesDestinations, "destinations", "", nil, "Destinations")
	_sesCmd.Flags().StringVarP(&_sesDkimEnabled, "dkim-enabled", "", "", "Dkim Enabled")
	_sesCmd.Flags().StringVarP(&_sesDomain, "domain", "", "", "Domain")
	_sesCmd.Flags().StringVarP(&_sesEmailAddress, "email-address", "", "", "Email Address")
	_sesCmd.Flags().StringVarP(&_sesEnabled, "enabled", "", "", "Enabled")
	_sesCmd.Flags().StringVarP(&_sesEventDestination, "event-destination", "", "", "Event Destination")
	_sesCmd.Flags().StringVarP(&_sesEventDestinationName, "event-destination-name", "", "", "Event Destination Name")
	_sesCmd.Flags().StringVarP(&_sesExplanation, "explanation", "", "", "Explanation")
	_sesCmd.Flags().StringVarP(&_sesFailureRedirectionURL, "failure-redirection-url", "", "", "Failure Redirection URL")
	_sesCmd.Flags().StringVarP(&_sesFilter, "filter", "", "", "Filter")
	_sesCmd.Flags().StringVarP(&_sesFilterName, "filter-name", "", "", "Filter Name")
	_sesCmd.Flags().StringVarP(&_sesForwardingEnabled, "forwarding-enabled", "", "", "Forwarding Enabled")
	_sesCmd.Flags().StringVarP(&_sesFromArn, "from-arn", "", "", "From ARN")
	_sesCmd.Flags().StringVarP(&_sesFromEmailAddress, "from-email-address", "", "", "From Email Address")
	_sesCmd.Flags().StringSliceVarP(&_sesIdentities, "identities", "", nil, "Identities")
	_sesCmd.Flags().StringVarP(&_sesIdentity, "identity", "", "", "Identity")
	_sesCmd.Flags().StringVarP(&_sesIdentityType, "identity-type", "", "", "Identity Type")
	_sesCmd.Flags().StringVarP(&_sesMailFromDomain, "mail-from-domain", "", "", "Mail From Domain")
	_sesCmd.Flags().StringVarP(&_sesMaxItems, "max-items", "", "", "Max Items")
	_sesCmd.Flags().StringVarP(&_sesMaxResults, "max-results", "", "", "Max Results")
	_sesCmd.Flags().StringVarP(&_sesMessage, "message", "", "", "Message")
	_sesCmd.Flags().StringVarP(&_sesMessageDsn, "message-dsn", "", "", "Message Dsn")
	_sesCmd.Flags().StringVarP(&_sesNextToken, "next-token", "", "", "Next Token")
	_sesCmd.Flags().StringVarP(&_sesNotificationType, "notification-type", "", "", "Notification Type")
	_sesCmd.Flags().StringVarP(&_sesOriginalMessageId, "original-message-id", "", "", "Original Message ID")
	_sesCmd.Flags().StringVarP(&_sesOriginalRuleSetName, "original-rule-set-name", "", "", "Original Rule Set Name")
	_sesCmd.Flags().StringVarP(&_sesPolicy, "policy", "", "", "Policy")
	_sesCmd.Flags().StringVarP(&_sesPolicyName, "policy-name", "", "", "Policy Name")
	_sesCmd.Flags().StringSliceVarP(&_sesPolicyNames, "policy-names", "", nil, "Policy Names")
	_sesCmd.Flags().StringVarP(&_sesRawMessage, "raw-message", "", "", "Raw Message")
	_sesCmd.Flags().StringSliceVarP(&_sesReplyToAddresses, "reply-to-addresses", "", nil, "Reply To Addresses")
	_sesCmd.Flags().StringVarP(&_sesReturnPath, "return-path", "", "", "Return Path")
	_sesCmd.Flags().StringVarP(&_sesReturnPathArn, "return-path-arn", "", "", "Return Path ARN")
	_sesCmd.Flags().StringVarP(&_sesRule, "rule", "", "", "Rule")
	_sesCmd.Flags().StringVarP(&_sesRuleName, "rule-name", "", "", "Rule Name")
	_sesCmd.Flags().StringSliceVarP(&_sesRuleNames, "rule-names", "", nil, "Rule Names")
	_sesCmd.Flags().StringVarP(&_sesRuleSetName, "rule-set-name", "", "", "Rule Set Name")
	_sesCmd.Flags().StringVarP(&_sesSnsTopic, "sns-topic", "", "", "SNS Topic")
	_sesCmd.Flags().StringVarP(&_sesSource, "source", "", "", "Source")
	_sesCmd.Flags().StringVarP(&_sesSourceArn, "source-arn", "", "", "Source ARN")
	_sesCmd.Flags().StringVarP(&_sesSuccessRedirectionURL, "success-redirection-url", "", "", "Success Redirection URL")
	_sesCmd.Flags().StringVarP(&_sesTags, "tags", "", "", "Tags")
	_sesCmd.Flags().StringVarP(&_sesTemplate, "template", "", "", "Template")
	_sesCmd.Flags().StringVarP(&_sesTemplateArn, "template-arn", "", "", "Template ARN")
	_sesCmd.Flags().StringVarP(&_sesTemplateContent, "template-content", "", "", "Template Content")
	_sesCmd.Flags().StringVarP(&_sesTemplateData, "template-data", "", "", "Template Data")
	_sesCmd.Flags().StringVarP(&_sesTemplateName, "template-name", "", "", "Template Name")
	_sesCmd.Flags().StringVarP(&_sesTemplateSubject, "template-subject", "", "", "Template Subject")
	_sesCmd.Flags().StringVarP(&_sesTrackingOptions, "tracking-options", "", "", "Tracking Options")

	_sesCmd.Flags().BoolVarP(&_sesCloneReceiptRuleSet, "clone-receipt-rule-set", "", false, "Clone Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesCreateConfigurationSet, "create-configuration-set", "", false, "Create Configuration Set")
	_sesCmd.Flags().BoolVarP(&_sesCreateConfigurationSetEventDestination, "create-configuration-set-event-destination", "", false, "Create Configuration Set Event Destination")
	_sesCmd.Flags().BoolVarP(&_sesCreateConfigurationSetTrackingOptions, "create-configuration-set-tracking-options", "", false, "Create Configuration Set Tracking Options")
	_sesCmd.Flags().BoolVarP(&_sesCreateCustomVerificationEmailTemplate, "create-custom-verification-email-template", "", false, "Create Custom Verification Email Template")
	_sesCmd.Flags().BoolVarP(&_sesCreateReceiptFilter, "create-receipt-filter", "", false, "Create Receipt Filter")
	_sesCmd.Flags().BoolVarP(&_sesCreateReceiptRule, "create-receipt-rule", "", false, "Create Receipt Rule")
	_sesCmd.Flags().BoolVarP(&_sesCreateReceiptRuleSet, "create-receipt-rule-set", "", false, "Create Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesCreateTemplate, "create-template", "", false, "Create Template")
	_sesCmd.Flags().BoolVarP(&_sesDeleteConfigurationSet, "delete-configuration-set", "", false, "Delete Configuration Set")
	_sesCmd.Flags().BoolVarP(&_sesDeleteConfigurationSetEventDestination, "delete-configuration-set-event-destination", "", false, "Delete Configuration Set Event Destination")
	_sesCmd.Flags().BoolVarP(&_sesDeleteConfigurationSetTrackingOptions, "delete-configuration-set-tracking-options", "", false, "Delete Configuration Set Tracking Options")
	_sesCmd.Flags().BoolVarP(&_sesDeleteCustomVerificationEmailTemplate, "delete-custom-verification-email-template", "", false, "Delete Custom Verification Email Template")
	_sesCmd.Flags().BoolVarP(&_sesDeleteIdentity, "delete-identity", "", false, "Delete Identity")
	_sesCmd.Flags().BoolVarP(&_sesDeleteIdentityPolicy, "delete-identity-policy", "", false, "Delete Identity Policy")
	_sesCmd.Flags().BoolVarP(&_sesDeleteReceiptFilter, "delete-receipt-filter", "", false, "Delete Receipt Filter")
	_sesCmd.Flags().BoolVarP(&_sesDeleteReceiptRule, "delete-receipt-rule", "", false, "Delete Receipt Rule")
	_sesCmd.Flags().BoolVarP(&_sesDeleteReceiptRuleSet, "delete-receipt-rule-set", "", false, "Delete Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesDeleteTemplate, "delete-template", "", false, "Delete Template")
	_sesCmd.Flags().BoolVarP(&_sesDeleteVerifiedEmailAddress, "delete-verified-email-address", "", false, "Delete Verified Email Address")
	_sesCmd.Flags().BoolVarP(&_sesDescribeActiveReceiptRuleSet, "describe-active-receipt-rule-set", "", false, "Describe Active Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesDescribeConfigurationSet, "describe-configuration-set", "", false, "Describe Configuration Set")
	_sesCmd.Flags().BoolVarP(&_sesDescribeReceiptRule, "describe-receipt-rule", "", false, "Describe Receipt Rule")
	_sesCmd.Flags().BoolVarP(&_sesDescribeReceiptRuleSet, "describe-receipt-rule-set", "", false, "Describe Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesGetAccountSendingEnabled, "get-account-sending-enabled", "", false, "Get Account Sending Enabled")
	_sesCmd.Flags().BoolVarP(&_sesGetCustomVerificationEmailTemplate, "get-custom-verification-email-template", "", false, "Get Custom Verification Email Template")
	_sesCmd.Flags().BoolVarP(&_sesGetIdentityDkimAttributes, "get-identity-dkim-attributes", "", false, "Get Identity Dkim Attributes")
	_sesCmd.Flags().BoolVarP(&_sesGetIdentityMailFromDomainAttributes, "get-identity-mail-from-domain-attributes", "", false, "Get Identity Mail From Domain Attributes")
	_sesCmd.Flags().BoolVarP(&_sesGetIdentityNotificationAttributes, "get-identity-notification-attributes", "", false, "Get Identity Notification Attributes")
	_sesCmd.Flags().BoolVarP(&_sesGetIdentityPolicies, "get-identity-policies", "", false, "Get Identity Policies")
	_sesCmd.Flags().BoolVarP(&_sesGetIdentityVerificationAttributes, "get-identity-verification-attributes", "", false, "Get Identity Verification Attributes")
	_sesCmd.Flags().BoolVarP(&_sesGetSendQuota, "get-send-quota", "", false, "Get Send Quota")
	_sesCmd.Flags().BoolVarP(&_sesGetSendStatistics, "get-send-statistics", "", false, "Get Send Statistics")
	_sesCmd.Flags().BoolVarP(&_sesGetTemplate, "get-template", "", false, "Get Template")
	_sesCmd.Flags().BoolVarP(&_sesListConfigurationSets, "list-configuration-sets", "", false, "List Configuration Sets")
	_sesCmd.Flags().BoolVarP(&_sesListCustomVerificationEmailTemplates, "list-custom-verification-email-templates", "", false, "List Custom Verification Email Templates")
	_sesCmd.Flags().BoolVarP(&_sesListIdentities, "list-identities", "", false, "List Identities")
	_sesCmd.Flags().BoolVarP(&_sesListIdentityPolicies, "list-identity-policies", "", false, "List Identity Policies")
	_sesCmd.Flags().BoolVarP(&_sesListReceiptFilters, "list-receipt-filters", "", false, "List Receipt Filters")
	_sesCmd.Flags().BoolVarP(&_sesListReceiptRuleSets, "list-receipt-rule-sets", "", false, "List Receipt Rule Sets")
	_sesCmd.Flags().BoolVarP(&_sesListTemplates, "list-templates", "", false, "List Templates")
	_sesCmd.Flags().BoolVarP(&_sesListVerifiedEmailAddresses, "list-verified-email-addresses", "", false, "List Verified Email Addresses")
	_sesCmd.Flags().BoolVarP(&_sesPutConfigurationSetDeliveryOptions, "put-configuration-set-delivery-options", "", false, "Put Configuration Set Delivery Options")
	_sesCmd.Flags().BoolVarP(&_sesPutIdentityPolicy, "put-identity-policy", "", false, "Put Identity Policy")
	_sesCmd.Flags().BoolVarP(&_sesReorderReceiptRuleSet, "reorder-receipt-rule-set", "", false, "Reorder Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesSendBounce, "send-bounce", "", false, "Send Bounce")
	_sesCmd.Flags().BoolVarP(&_sesSendBulkTemplatedEmail, "send-bulk-templated-email", "", false, "Send Bulk Templated Email")
	_sesCmd.Flags().BoolVarP(&_sesSendCustomVerificationEmail, "send-custom-verification-email", "", false, "Send Custom Verification Email")
	_sesCmd.Flags().BoolVarP(&_sesSendEmail, "send-email", "", false, "Send Email")
	_sesCmd.Flags().BoolVarP(&_sesSendRawEmail, "send-raw-email", "", false, "Send Raw Email")
	_sesCmd.Flags().BoolVarP(&_sesSendTemplatedEmail, "send-templated-email", "", false, "Send Templated Email")
	_sesCmd.Flags().BoolVarP(&_sesSetActiveReceiptRuleSet, "set-active-receipt-rule-set", "", false, "Set Active Receipt Rule Set")
	_sesCmd.Flags().BoolVarP(&_sesSetIdentityDkimEnabled, "set-identity-dkim-enabled", "", false, "Set Identity Dkim Enabled")
	_sesCmd.Flags().BoolVarP(&_sesSetIdentityFeedbackForwardingEnabled, "set-identity-feedback-forwarding-enabled", "", false, "Set Identity Feedback Forwarding Enabled")
	_sesCmd.Flags().BoolVarP(&_sesSetIdentityHeadersInNotificationsEnabled, "set-identity-headers-in-notifications-enabled", "", false, "Set Identity Headers In Notifications Enabled")
	_sesCmd.Flags().BoolVarP(&_sesSetIdentityMailFromDomain, "set-identity-mail-from-domain", "", false, "Set Identity Mail From Domain")
	_sesCmd.Flags().BoolVarP(&_sesSetIdentityNotificationTopic, "set-identity-notification-topic", "", false, "Set Identity Notification Topic")
	_sesCmd.Flags().BoolVarP(&_sesSetReceiptRulePosition, "set-receipt-rule-position", "", false, "Set Receipt Rule Position")
	_sesCmd.Flags().BoolVarP(&_sesTestRenderTemplate, "test-render-template", "", false, "Test Render Template")
	_sesCmd.Flags().BoolVarP(&_sesUpdateAccountSendingEnabled, "update-account-sending-enabled", "", false, "Update Account Sending Enabled")
	_sesCmd.Flags().BoolVarP(&_sesUpdateConfigurationSetEventDestination, "update-configuration-set-event-destination", "", false, "Update Configuration Set Event Destination")
	_sesCmd.Flags().BoolVarP(&_sesUpdateConfigurationSetReputationMetricsEnabled, "update-configuration-set-reputation-metrics-enabled", "", false, "Update Configuration Set Reputation Metrics Enabled")
	_sesCmd.Flags().BoolVarP(&_sesUpdateConfigurationSetSendingEnabled, "update-configuration-set-sending-enabled", "", false, "Update Configuration Set Sending Enabled")
	_sesCmd.Flags().BoolVarP(&_sesUpdateConfigurationSetTrackingOptions, "update-configuration-set-tracking-options", "", false, "Update Configuration Set Tracking Options")
	_sesCmd.Flags().BoolVarP(&_sesUpdateCustomVerificationEmailTemplate, "update-custom-verification-email-template", "", false, "Update Custom Verification Email Template")
	_sesCmd.Flags().BoolVarP(&_sesUpdateReceiptRule, "update-receipt-rule", "", false, "Update Receipt Rule")
	_sesCmd.Flags().BoolVarP(&_sesUpdateTemplate, "update-template", "", false, "Update Template")
	_sesCmd.Flags().BoolVarP(&_sesVerifyDomainDkim, "verify-domain-dkim", "", false, "Verify Domain Dkim")
	_sesCmd.Flags().BoolVarP(&_sesVerifyDomainIdentity, "verify-domain-identity", "", false, "Verify Domain Identity")
	_sesCmd.Flags().BoolVarP(&_sesVerifyEmailAddress, "verify-email-address", "", false, "Verify Email Address")
	_sesCmd.Flags().BoolVarP(&_sesVerifyEmailIdentity, "verify-email-identity", "", false, "Verify Email Identity")

}
