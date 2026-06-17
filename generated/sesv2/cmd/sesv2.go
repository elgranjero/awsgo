package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sesv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// sesv2Cmd represents the sesv2 command
var _sesv2Cmd = &cobra.Command{
	Use:   "sesv2",
	Short: "AWS sesv2 CLI",
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
		client := sesv2.NewFromConfig(cfg)
		if _sesv2BatchGetMetricData {
			sesv2_BatchGetMetricData(cfg, client)
			return
		}
		if _sesv2CancelExportJob {
			sesv2_CancelExportJob(cfg, client)
			return
		}
		if _sesv2CreateConfigurationSet {
			sesv2_CreateConfigurationSet(cfg, client)
			return
		}
		if _sesv2CreateConfigurationSetEventDestination {
			sesv2_CreateConfigurationSetEventDestination(cfg, client)
			return
		}
		if _sesv2CreateContact {
			sesv2_CreateContact(cfg, client)
			return
		}
		if _sesv2CreateContactList {
			sesv2_CreateContactList(cfg, client)
			return
		}
		if _sesv2CreateCustomVerificationEmailTemplate {
			sesv2_CreateCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesv2CreateDedicatedIpPool {
			sesv2_CreateDedicatedIpPool(cfg, client)
			return
		}
		if _sesv2CreateDeliverabilityTestReport {
			sesv2_CreateDeliverabilityTestReport(cfg, client)
			return
		}
		if _sesv2CreateEmailIdentity {
			sesv2_CreateEmailIdentity(cfg, client)
			return
		}
		if _sesv2CreateEmailIdentityPolicy {
			sesv2_CreateEmailIdentityPolicy(cfg, client)
			return
		}
		if _sesv2CreateEmailTemplate {
			sesv2_CreateEmailTemplate(cfg, client)
			return
		}
		if _sesv2CreateExportJob {
			sesv2_CreateExportJob(cfg, client)
			return
		}
		if _sesv2CreateImportJob {
			sesv2_CreateImportJob(cfg, client)
			return
		}
		if _sesv2CreateMultiRegionEndpoint {
			sesv2_CreateMultiRegionEndpoint(cfg, client)
			return
		}
		if _sesv2CreateTenant {
			sesv2_CreateTenant(cfg, client)
			return
		}
		if _sesv2CreateTenantResourceAssociation {
			sesv2_CreateTenantResourceAssociation(cfg, client)
			return
		}
		if _sesv2DeleteConfigurationSet {
			sesv2_DeleteConfigurationSet(cfg, client)
			return
		}
		if _sesv2DeleteConfigurationSetEventDestination {
			sesv2_DeleteConfigurationSetEventDestination(cfg, client)
			return
		}
		if _sesv2DeleteContact {
			sesv2_DeleteContact(cfg, client)
			return
		}
		if _sesv2DeleteContactList {
			sesv2_DeleteContactList(cfg, client)
			return
		}
		if _sesv2DeleteCustomVerificationEmailTemplate {
			sesv2_DeleteCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesv2DeleteDedicatedIpPool {
			sesv2_DeleteDedicatedIpPool(cfg, client)
			return
		}
		if _sesv2DeleteEmailIdentity {
			sesv2_DeleteEmailIdentity(cfg, client)
			return
		}
		if _sesv2DeleteEmailIdentityPolicy {
			sesv2_DeleteEmailIdentityPolicy(cfg, client)
			return
		}
		if _sesv2DeleteEmailTemplate {
			sesv2_DeleteEmailTemplate(cfg, client)
			return
		}
		if _sesv2DeleteMultiRegionEndpoint {
			sesv2_DeleteMultiRegionEndpoint(cfg, client)
			return
		}
		if _sesv2DeleteSuppressedDestination {
			sesv2_DeleteSuppressedDestination(cfg, client)
			return
		}
		if _sesv2DeleteTenant {
			sesv2_DeleteTenant(cfg, client)
			return
		}
		if _sesv2DeleteTenantResourceAssociation {
			sesv2_DeleteTenantResourceAssociation(cfg, client)
			return
		}
		if _sesv2GetAccount {
			sesv2_GetAccount(cfg, client)
			return
		}
		if _sesv2GetBlacklistReports {
			sesv2_GetBlacklistReports(cfg, client)
			return
		}
		if _sesv2GetConfigurationSet {
			sesv2_GetConfigurationSet(cfg, client)
			return
		}
		if _sesv2GetConfigurationSetEventDestinations {
			sesv2_GetConfigurationSetEventDestinations(cfg, client)
			return
		}
		if _sesv2GetContact {
			sesv2_GetContact(cfg, client)
			return
		}
		if _sesv2GetContactList {
			sesv2_GetContactList(cfg, client)
			return
		}
		if _sesv2GetCustomVerificationEmailTemplate {
			sesv2_GetCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesv2GetDedicatedIp {
			sesv2_GetDedicatedIp(cfg, client)
			return
		}
		if _sesv2GetDedicatedIpPool {
			sesv2_GetDedicatedIpPool(cfg, client)
			return
		}
		if _sesv2GetDedicatedIps {
			sesv2_GetDedicatedIps(cfg, client)
			return
		}
		if _sesv2GetDeliverabilityDashboardOptions {
			sesv2_GetDeliverabilityDashboardOptions(cfg, client)
			return
		}
		if _sesv2GetDeliverabilityTestReport {
			sesv2_GetDeliverabilityTestReport(cfg, client)
			return
		}
		if _sesv2GetDomainDeliverabilityCampaign {
			sesv2_GetDomainDeliverabilityCampaign(cfg, client)
			return
		}
		if _sesv2GetDomainStatisticsReport {
			sesv2_GetDomainStatisticsReport(cfg, client)
			return
		}
		if _sesv2GetEmailAddressInsights {
			sesv2_GetEmailAddressInsights(cfg, client)
			return
		}
		if _sesv2GetEmailIdentity {
			sesv2_GetEmailIdentity(cfg, client)
			return
		}
		if _sesv2GetEmailIdentityPolicies {
			sesv2_GetEmailIdentityPolicies(cfg, client)
			return
		}
		if _sesv2GetEmailTemplate {
			sesv2_GetEmailTemplate(cfg, client)
			return
		}
		if _sesv2GetExportJob {
			sesv2_GetExportJob(cfg, client)
			return
		}
		if _sesv2GetImportJob {
			sesv2_GetImportJob(cfg, client)
			return
		}
		if _sesv2GetMessageInsights {
			sesv2_GetMessageInsights(cfg, client)
			return
		}
		if _sesv2GetMultiRegionEndpoint {
			sesv2_GetMultiRegionEndpoint(cfg, client)
			return
		}
		if _sesv2GetReputationEntity {
			sesv2_GetReputationEntity(cfg, client)
			return
		}
		if _sesv2GetSuppressedDestination {
			sesv2_GetSuppressedDestination(cfg, client)
			return
		}
		if _sesv2GetTenant {
			sesv2_GetTenant(cfg, client)
			return
		}
		if _sesv2ListConfigurationSets {
			sesv2_ListConfigurationSets(cfg, client)
			return
		}
		if _sesv2ListContactLists {
			sesv2_ListContactLists(cfg, client)
			return
		}
		if _sesv2ListContacts {
			sesv2_ListContacts(cfg, client)
			return
		}
		if _sesv2ListCustomVerificationEmailTemplates {
			sesv2_ListCustomVerificationEmailTemplates(cfg, client)
			return
		}
		if _sesv2ListDedicatedIpPools {
			sesv2_ListDedicatedIpPools(cfg, client)
			return
		}
		if _sesv2ListDeliverabilityTestReports {
			sesv2_ListDeliverabilityTestReports(cfg, client)
			return
		}
		if _sesv2ListDomainDeliverabilityCampaigns {
			sesv2_ListDomainDeliverabilityCampaigns(cfg, client)
			return
		}
		if _sesv2ListEmailIdentities {
			sesv2_ListEmailIdentities(cfg, client)
			return
		}
		if _sesv2ListEmailTemplates {
			sesv2_ListEmailTemplates(cfg, client)
			return
		}
		if _sesv2ListExportJobs {
			sesv2_ListExportJobs(cfg, client)
			return
		}
		if _sesv2ListImportJobs {
			sesv2_ListImportJobs(cfg, client)
			return
		}
		if _sesv2ListMultiRegionEndpoints {
			sesv2_ListMultiRegionEndpoints(cfg, client)
			return
		}
		if _sesv2ListRecommendations {
			sesv2_ListRecommendations(cfg, client)
			return
		}
		if _sesv2ListReputationEntities {
			sesv2_ListReputationEntities(cfg, client)
			return
		}
		if _sesv2ListResourceTenants {
			sesv2_ListResourceTenants(cfg, client)
			return
		}
		if _sesv2ListSuppressedDestinations {
			sesv2_ListSuppressedDestinations(cfg, client)
			return
		}
		if _sesv2ListTagsForResource {
			sesv2_ListTagsForResource(cfg, client)
			return
		}
		if _sesv2ListTenantResources {
			sesv2_ListTenantResources(cfg, client)
			return
		}
		if _sesv2ListTenants {
			sesv2_ListTenants(cfg, client)
			return
		}
		if _sesv2PutAccountDedicatedIpWarmupAttributes {
			sesv2_PutAccountDedicatedIpWarmupAttributes(cfg, client)
			return
		}
		if _sesv2PutAccountDetails {
			sesv2_PutAccountDetails(cfg, client)
			return
		}
		if _sesv2PutAccountSendingAttributes {
			sesv2_PutAccountSendingAttributes(cfg, client)
			return
		}
		if _sesv2PutAccountSuppressionAttributes {
			sesv2_PutAccountSuppressionAttributes(cfg, client)
			return
		}
		if _sesv2PutAccountVdmAttributes {
			sesv2_PutAccountVdmAttributes(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetArchivingOptions {
			sesv2_PutConfigurationSetArchivingOptions(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetDeliveryOptions {
			sesv2_PutConfigurationSetDeliveryOptions(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetReputationOptions {
			sesv2_PutConfigurationSetReputationOptions(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetSendingOptions {
			sesv2_PutConfigurationSetSendingOptions(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetSuppressionOptions {
			sesv2_PutConfigurationSetSuppressionOptions(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetTrackingOptions {
			sesv2_PutConfigurationSetTrackingOptions(cfg, client)
			return
		}
		if _sesv2PutConfigurationSetVdmOptions {
			sesv2_PutConfigurationSetVdmOptions(cfg, client)
			return
		}
		if _sesv2PutDedicatedIpInPool {
			sesv2_PutDedicatedIpInPool(cfg, client)
			return
		}
		if _sesv2PutDedicatedIpPoolScalingAttributes {
			sesv2_PutDedicatedIpPoolScalingAttributes(cfg, client)
			return
		}
		if _sesv2PutDedicatedIpWarmupAttributes {
			sesv2_PutDedicatedIpWarmupAttributes(cfg, client)
			return
		}
		if _sesv2PutDeliverabilityDashboardOption {
			sesv2_PutDeliverabilityDashboardOption(cfg, client)
			return
		}
		if _sesv2PutEmailIdentityConfigurationSetAttributes {
			sesv2_PutEmailIdentityConfigurationSetAttributes(cfg, client)
			return
		}
		if _sesv2PutEmailIdentityDkimAttributes {
			sesv2_PutEmailIdentityDkimAttributes(cfg, client)
			return
		}
		if _sesv2PutEmailIdentityDkimSigningAttributes {
			sesv2_PutEmailIdentityDkimSigningAttributes(cfg, client)
			return
		}
		if _sesv2PutEmailIdentityFeedbackAttributes {
			sesv2_PutEmailIdentityFeedbackAttributes(cfg, client)
			return
		}
		if _sesv2PutEmailIdentityMailFromAttributes {
			sesv2_PutEmailIdentityMailFromAttributes(cfg, client)
			return
		}
		if _sesv2PutSuppressedDestination {
			sesv2_PutSuppressedDestination(cfg, client)
			return
		}
		if _sesv2SendBulkEmail {
			sesv2_SendBulkEmail(cfg, client)
			return
		}
		if _sesv2SendCustomVerificationEmail {
			sesv2_SendCustomVerificationEmail(cfg, client)
			return
		}
		if _sesv2SendEmail {
			sesv2_SendEmail(cfg, client)
			return
		}
		if _sesv2TagResource {
			sesv2_TagResource(cfg, client)
			return
		}
		if _sesv2TestRenderEmailTemplate {
			sesv2_TestRenderEmailTemplate(cfg, client)
			return
		}
		if _sesv2UntagResource {
			sesv2_UntagResource(cfg, client)
			return
		}
		if _sesv2UpdateConfigurationSetEventDestination {
			sesv2_UpdateConfigurationSetEventDestination(cfg, client)
			return
		}
		if _sesv2UpdateContact {
			sesv2_UpdateContact(cfg, client)
			return
		}
		if _sesv2UpdateContactList {
			sesv2_UpdateContactList(cfg, client)
			return
		}
		if _sesv2UpdateCustomVerificationEmailTemplate {
			sesv2_UpdateCustomVerificationEmailTemplate(cfg, client)
			return
		}
		if _sesv2UpdateEmailIdentityPolicy {
			sesv2_UpdateEmailIdentityPolicy(cfg, client)
			return
		}
		if _sesv2UpdateEmailTemplate {
			sesv2_UpdateEmailTemplate(cfg, client)
			return
		}
		if _sesv2UpdateReputationEntityCustomerManagedStatus {
			sesv2_UpdateReputationEntityCustomerManagedStatus(cfg, client)
			return
		}
		if _sesv2UpdateReputationEntityPolicy {
			sesv2_UpdateReputationEntityPolicy(cfg, client)
			return
		}

	},
}

var (
	_sesv2BatchGetMetricData                          bool
	_sesv2CancelExportJob                             bool
	_sesv2CreateConfigurationSet                      bool
	_sesv2CreateConfigurationSetEventDestination      bool
	_sesv2CreateContact                               bool
	_sesv2CreateContactList                           bool
	_sesv2CreateCustomVerificationEmailTemplate       bool
	_sesv2CreateDedicatedIpPool                       bool
	_sesv2CreateDeliverabilityTestReport              bool
	_sesv2CreateEmailIdentity                         bool
	_sesv2CreateEmailIdentityPolicy                   bool
	_sesv2CreateEmailTemplate                         bool
	_sesv2CreateExportJob                             bool
	_sesv2CreateImportJob                             bool
	_sesv2CreateMultiRegionEndpoint                   bool
	_sesv2CreateTenant                                bool
	_sesv2CreateTenantResourceAssociation             bool
	_sesv2DeleteConfigurationSet                      bool
	_sesv2DeleteConfigurationSetEventDestination      bool
	_sesv2DeleteContact                               bool
	_sesv2DeleteContactList                           bool
	_sesv2DeleteCustomVerificationEmailTemplate       bool
	_sesv2DeleteDedicatedIpPool                       bool
	_sesv2DeleteEmailIdentity                         bool
	_sesv2DeleteEmailIdentityPolicy                   bool
	_sesv2DeleteEmailTemplate                         bool
	_sesv2DeleteMultiRegionEndpoint                   bool
	_sesv2DeleteSuppressedDestination                 bool
	_sesv2DeleteTenant                                bool
	_sesv2DeleteTenantResourceAssociation             bool
	_sesv2GetAccount                                  bool
	_sesv2GetBlacklistReports                         bool
	_sesv2GetConfigurationSet                         bool
	_sesv2GetConfigurationSetEventDestinations        bool
	_sesv2GetContact                                  bool
	_sesv2GetContactList                              bool
	_sesv2GetCustomVerificationEmailTemplate          bool
	_sesv2GetDedicatedIp                              bool
	_sesv2GetDedicatedIpPool                          bool
	_sesv2GetDedicatedIps                             bool
	_sesv2GetDeliverabilityDashboardOptions           bool
	_sesv2GetDeliverabilityTestReport                 bool
	_sesv2GetDomainDeliverabilityCampaign             bool
	_sesv2GetDomainStatisticsReport                   bool
	_sesv2GetEmailAddressInsights                     bool
	_sesv2GetEmailIdentity                            bool
	_sesv2GetEmailIdentityPolicies                    bool
	_sesv2GetEmailTemplate                            bool
	_sesv2GetExportJob                                bool
	_sesv2GetImportJob                                bool
	_sesv2GetMessageInsights                          bool
	_sesv2GetMultiRegionEndpoint                      bool
	_sesv2GetReputationEntity                         bool
	_sesv2GetSuppressedDestination                    bool
	_sesv2GetTenant                                   bool
	_sesv2ListConfigurationSets                       bool
	_sesv2ListContactLists                            bool
	_sesv2ListContacts                                bool
	_sesv2ListCustomVerificationEmailTemplates        bool
	_sesv2ListDedicatedIpPools                        bool
	_sesv2ListDeliverabilityTestReports               bool
	_sesv2ListDomainDeliverabilityCampaigns           bool
	_sesv2ListEmailIdentities                         bool
	_sesv2ListEmailTemplates                          bool
	_sesv2ListExportJobs                              bool
	_sesv2ListImportJobs                              bool
	_sesv2ListMultiRegionEndpoints                    bool
	_sesv2ListRecommendations                         bool
	_sesv2ListReputationEntities                      bool
	_sesv2ListResourceTenants                         bool
	_sesv2ListSuppressedDestinations                  bool
	_sesv2ListTagsForResource                         bool
	_sesv2ListTenantResources                         bool
	_sesv2ListTenants                                 bool
	_sesv2PutAccountDedicatedIpWarmupAttributes       bool
	_sesv2PutAccountDetails                           bool
	_sesv2PutAccountSendingAttributes                 bool
	_sesv2PutAccountSuppressionAttributes             bool
	_sesv2PutAccountVdmAttributes                     bool
	_sesv2PutConfigurationSetArchivingOptions         bool
	_sesv2PutConfigurationSetDeliveryOptions          bool
	_sesv2PutConfigurationSetReputationOptions        bool
	_sesv2PutConfigurationSetSendingOptions           bool
	_sesv2PutConfigurationSetSuppressionOptions       bool
	_sesv2PutConfigurationSetTrackingOptions          bool
	_sesv2PutConfigurationSetVdmOptions               bool
	_sesv2PutDedicatedIpInPool                        bool
	_sesv2PutDedicatedIpPoolScalingAttributes         bool
	_sesv2PutDedicatedIpWarmupAttributes              bool
	_sesv2PutDeliverabilityDashboardOption            bool
	_sesv2PutEmailIdentityConfigurationSetAttributes  bool
	_sesv2PutEmailIdentityDkimAttributes              bool
	_sesv2PutEmailIdentityDkimSigningAttributes       bool
	_sesv2PutEmailIdentityFeedbackAttributes          bool
	_sesv2PutEmailIdentityMailFromAttributes          bool
	_sesv2PutSuppressedDestination                    bool
	_sesv2SendBulkEmail                               bool
	_sesv2SendCustomVerificationEmail                 bool
	_sesv2SendEmail                                   bool
	_sesv2TagResource                                 bool
	_sesv2TestRenderEmailTemplate                     bool
	_sesv2UntagResource                               bool
	_sesv2UpdateConfigurationSetEventDestination      bool
	_sesv2UpdateContact                               bool
	_sesv2UpdateContactList                           bool
	_sesv2UpdateCustomVerificationEmailTemplate       bool
	_sesv2UpdateEmailIdentityPolicy                   bool
	_sesv2UpdateEmailTemplate                         bool
	_sesv2UpdateReputationEntityCustomerManagedStatus bool
	_sesv2UpdateReputationEntityPolicy                bool

	_sesv2AdditionalContactEmailAddresses           []string
	_sesv2ArchiveArn                                string
	_sesv2ArchivingOptions                          string
	_sesv2AttributesData                            string
	_sesv2AutoWarmupEnabled                         string
	_sesv2BehaviorOnMxFailure                       string
	_sesv2BlacklistItemNames                        []string
	_sesv2BulkEmailEntries                          string
	_sesv2CampaignId                                string
	_sesv2ConfigurationSetName                      string
	_sesv2ContactLanguage                           string
	_sesv2ContactListName                           string
	_sesv2Content                                   string
	_sesv2CustomRedirectDomain                      string
	_sesv2DashboardEnabled                          string
	_sesv2DefaultContent                            string
	_sesv2DefaultEmailTags                          string
	_sesv2DeliveryOptions                           string
	_sesv2Description                               string
	_sesv2Destination                               string
	_sesv2DestinationPoolName                       string
	_sesv2Details                                   string
	_sesv2DkimSigningAttributes                     string
	_sesv2Domain                                    string
	_sesv2EmailAddress                              string
	_sesv2EmailForwardingEnabled                    string
	_sesv2EmailIdentity                             string
	_sesv2EmailTags                                 string
	_sesv2EndDate                                   string
	_sesv2EndpointId                                string
	_sesv2EndpointName                              string
	_sesv2EventDestination                          string
	_sesv2EventDestinationName                      string
	_sesv2ExportDataSource                          string
	_sesv2ExportDestination                         string
	_sesv2ExportSourceType                          string
	_sesv2FailureRedirectionURL                     string
	_sesv2FeedbackForwardingEmailAddress            string
	_sesv2FeedbackForwardingEmailAddressIdentityArn string
	_sesv2Filter                                    string
	_sesv2FromEmailAddress                          string
	_sesv2FromEmailAddressIdentityArn               string
	_sesv2HttpsPolicy                               string
	_sesv2ImportDataSource                          string
	_sesv2ImportDestination                         string
	_sesv2ImportDestinationType                     string
	_sesv2Ip                                        string
	_sesv2JobId                                     string
	_sesv2JobStatus                                 string
	_sesv2ListManagementOptions                     string
	_sesv2MailFromDomain                            string
	_sesv2MailType                                  string
	_sesv2MaxDeliverySeconds                        string
	_sesv2MessageId                                 string
	_sesv2NextToken                                 string
	_sesv2PageSize                                  string
	_sesv2Policy                                    string
	_sesv2PolicyName                                string
	_sesv2PoolName                                  string
	_sesv2ProductionAccessEnabled                   string
	_sesv2Queries                                   string
	_sesv2Reason                                    string
	_sesv2Reasons                                   string
	_sesv2ReplyToAddresses                          []string
	_sesv2ReportId                                  string
	_sesv2ReportName                                string
	_sesv2ReputationEntityPolicy                    string
	_sesv2ReputationEntityReference                 string
	_sesv2ReputationEntityType                      string
	_sesv2ReputationMetricsEnabled                  string
	_sesv2ReputationOptions                         string
	_sesv2ResourceArn                               string
	_sesv2ScalingMode                               string
	_sesv2SendingEnabled                            string
	_sesv2SendingOptions                            string
	_sesv2SendingPoolName                           string
	_sesv2SendingStatus                             string
	_sesv2SigningAttributes                         string
	_sesv2SigningAttributesOrigin                   string
	_sesv2SigningEnabled                            string
	_sesv2StartDate                                 string
	_sesv2SubscribedDomain                          string
	_sesv2SubscribedDomains                         string
	_sesv2SuccessRedirectionURL                     string
	_sesv2SuppressedReasons                         string
	_sesv2SuppressionOptions                        string
	_sesv2TagKeys                                   []string
	_sesv2Tags                                      string
	_sesv2TemplateContent                           string
	_sesv2TemplateData                              string
	_sesv2TemplateName                              string
	_sesv2TemplateSubject                           string
	_sesv2TenantName                                string
	_sesv2TlsPolicy                                 string
	_sesv2TopicPreferences                          string
	_sesv2Topics                                    string
	_sesv2TrackingOptions                           string
	_sesv2UnsubscribeAll                            string
	_sesv2UseCaseDescription                        string
	_sesv2ValidationAttributes                      string
	_sesv2ValidationOptions                         string
	_sesv2VdmAttributes                             string
	_sesv2VdmOptions                                string
	_sesv2WarmupPercentage                          string
	_sesv2WebsiteURL                                string
)

// Retrieves batches of metric data collected based on your sending activity.
// You can execute this operation no more than 16 times per second, and with at
// most 160 queries from the batches per second (cumulative).
func sesv2_BatchGetMetricData(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.BatchGetMetricDataInput{
		// Queries: []types.BatchGetMetricDataQuery, // Required
	}

	if len(_sesv2Queries) > 0 {
		if err := assignInputField(input, "Queries", _sesv2Queries); err != nil {
			log.Errorf("invalid --queries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetMetricData(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an export job.
func sesv2_CancelExportJob(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CancelExportJobInput{
		// JobId: *string, // Required
	}

	if len(_sesv2JobId) > 0 {
		input.JobId = aws.String(_sesv2JobId)
	}

	if resp, err := client.CancelExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a configuration set. Configuration sets are groups of rules that you can
// apply to the emails that you send. You apply a configuration set to an email by
// specifying the name of the configuration set when you call the Amazon SES API
// v2. When you apply a configuration set to an email, all of the rules in that
// configuration set are applied to the email.
func sesv2_CreateConfigurationSet(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2ArchivingOptions) > 0 {
		if err := assignInputField(input, "ArchivingOptions", _sesv2ArchivingOptions); err != nil {
			log.Errorf("invalid --archiving-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2DeliveryOptions) > 0 {
		if err := assignInputField(input, "DeliveryOptions", _sesv2DeliveryOptions); err != nil {
			log.Errorf("invalid --delivery-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2ReputationOptions) > 0 {
		if err := assignInputField(input, "ReputationOptions", _sesv2ReputationOptions); err != nil {
			log.Errorf("invalid --reputation-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2SendingOptions) > 0 {
		if err := assignInputField(input, "SendingOptions", _sesv2SendingOptions); err != nil {
			log.Errorf("invalid --sending-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2SuppressionOptions) > 0 {
		if err := assignInputField(input, "SuppressionOptions", _sesv2SuppressionOptions); err != nil {
			log.Errorf("invalid --suppression-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sesv2TrackingOptions) > 0 {
		if err := assignInputField(input, "TrackingOptions", _sesv2TrackingOptions); err != nil {
			log.Errorf("invalid --tracking-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2VdmOptions) > 0 {
		if err := assignInputField(input, "VdmOptions", _sesv2VdmOptions); err != nil {
			log.Errorf("invalid --vdm-options: %s", err.Error())
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

// Create an event destination. Events include message sends, deliveries, opens,
// clicks, bounces, and complaints. Event destinations are places that you can send
// information about these events to. For example, you can send event data to
// Amazon EventBridge and associate a rule to send the event to the specified
// target.
//
// A single configuration set can include more than one event destination.
func sesv2_CreateConfigurationSetEventDestination(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestination: *types.EventDestinationDefinition, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2EventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _sesv2EventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}
	if len(_sesv2EventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_sesv2EventDestinationName)
	}

	if resp, err := client.CreateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a contact, which is an end-user who is receiving the email, and adds
// them to a contact list.
func sesv2_CreateContact(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateContactInput{
		// ContactListName: *string, // Required
		// EmailAddress: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}
	if len(_sesv2AttributesData) > 0 {
		input.AttributesData = aws.String(_sesv2AttributesData)
	}
	if len(_sesv2TopicPreferences) > 0 {
		if err := assignInputField(input, "TopicPreferences", _sesv2TopicPreferences); err != nil {
			log.Errorf("invalid --topic-preferences: %s", err.Error())
			return
		}
	}
	if len(_sesv2UnsubscribeAll) > 0 {
		if err := assignInputField(input, "UnsubscribeAll", _sesv2UnsubscribeAll); err != nil {
			log.Errorf("invalid --unsubscribe-all: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a contact list.
func sesv2_CreateContactList(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateContactListInput{
		// ContactListName: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2Description) > 0 {
		input.Description = aws.String(_sesv2Description)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_sesv2Topics) > 0 {
		if err := assignInputField(input, "Topics", _sesv2Topics); err != nil {
			log.Errorf("invalid --topics: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateContactList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom verification email template.
// For more information about custom verification email templates, see [Using custom verification email templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func sesv2_CreateCustomVerificationEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateCustomVerificationEmailTemplateInput{
		// FailureRedirectionURL: *string, // Required
		// FromEmailAddress: *string, // Required
		// SuccessRedirectionURL: *string, // Required
		// TemplateContent: *string, // Required
		// TemplateName: *string, // Required
		// TemplateSubject: *string, // Required
	}

	if len(_sesv2FailureRedirectionURL) > 0 {
		input.FailureRedirectionURL = aws.String(_sesv2FailureRedirectionURL)
	}
	if len(_sesv2FromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesv2FromEmailAddress)
	}
	if len(_sesv2SuccessRedirectionURL) > 0 {
		input.SuccessRedirectionURL = aws.String(_sesv2SuccessRedirectionURL)
	}
	if len(_sesv2TemplateContent) > 0 {
		input.TemplateContent = aws.String(_sesv2TemplateContent)
	}
	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}
	if len(_sesv2TemplateSubject) > 0 {
		input.TemplateSubject = aws.String(_sesv2TemplateSubject)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new pool of dedicated IP addresses. A pool can include one or more
// dedicated IP addresses that are associated with your Amazon Web Services
// account. You can associate a pool with a configuration set. When you send an
// email that uses that configuration set, the message is sent from one of the
// addresses in the associated pool.
func sesv2_CreateDedicatedIpPool(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateDedicatedIpPoolInput{
		// PoolName: *string, // Required
	}

	if len(_sesv2PoolName) > 0 {
		input.PoolName = aws.String(_sesv2PoolName)
	}
	if len(_sesv2ScalingMode) > 0 {
		if err := assignInputField(input, "ScalingMode", _sesv2ScalingMode); err != nil {
			log.Errorf("invalid --scaling-mode: %s", err.Error())
			return
		}
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
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
// your customers. Amazon SES then sends that message to special email addresses
// spread across several major email providers. After about 24 hours, the test is
// complete, and you can use the GetDeliverabilityTestReport operation to view the
// results of the test.
func sesv2_CreateDeliverabilityTestReport(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateDeliverabilityTestReportInput{
		// Content: *types.EmailContent, // Required
		// FromEmailAddress: *string, // Required
	}

	if len(_sesv2Content) > 0 {
		if err := assignInputField(input, "Content", _sesv2Content); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_sesv2FromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesv2FromEmailAddress)
	}
	if len(_sesv2ReportName) > 0 {
		input.ReportName = aws.String(_sesv2ReportName)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
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

// Starts the process of verifying an email identity. An identity is an email
// address or domain that you use when you send email. Before you can use an
// identity to send email, you first have to verify it. By verifying an identity,
// you demonstrate that you're the owner of the identity, and that you've given
// Amazon SES API v2 permission to send email from the identity.
//
// When you verify an email address, Amazon SES sends an email to the address.
// Your email address is verified as soon as you follow the link in the
// verification email.
//
// When you verify a domain without specifying the DkimSigningAttributes object,
// this operation provides a set of DKIM tokens. You can convert these tokens into
// CNAME records, which you then add to the DNS configuration for your domain. Your
// domain is verified when Amazon SES detects these records in the DNS
// configuration for your domain. This verification method is known as [Easy DKIM].
//
// Alternatively, you can perform the verification process by providing your own
// public-private key pair. This verification method is known as Bring Your Own
// DKIM (BYODKIM). To use BYODKIM, your call to the CreateEmailIdentity operation
// has to include the DkimSigningAttributes object. When you specify this object,
// you provide a selector (a component of the DNS record name that identifies the
// public key to use for DKIM authentication) and a private key.
//
// When you verify a domain, this operation provides a set of DKIM tokens, which
// you can convert into CNAME tokens. You add these CNAME tokens to the DNS
// configuration for your domain. Your domain is verified when Amazon SES detects
// these records in the DNS configuration for your domain. For some DNS providers,
// it can take 72 hours or more to complete the domain verification process.
//
// Additionally, you can associate an existing configuration set with the email
// identity that you're verifying.
//
// [Easy DKIM]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/easy-dkim.html
func sesv2_CreateEmailIdentity(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateEmailIdentityInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2DkimSigningAttributes) > 0 {
		if err := assignInputField(input, "DkimSigningAttributes", _sesv2DkimSigningAttributes); err != nil {
			log.Errorf("invalid --dkim-signing-attributes: %s", err.Error())
			return
		}
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
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

// Creates the specified sending authorization policy for the given identity (an
// email address or a domain).
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html
func sesv2_CreateEmailIdentityPolicy(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateEmailIdentityPolicyInput{
		// EmailIdentity: *string, // Required
		// Policy: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2Policy) > 0 {
		input.Policy = aws.String(_sesv2Policy)
	}
	if len(_sesv2PolicyName) > 0 {
		input.PolicyName = aws.String(_sesv2PolicyName)
	}

	if resp, err := client.CreateEmailIdentityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an email template. Email templates enable you to send personalized
// email to one or more destinations in a single API operation. For more
// information, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html
func sesv2_CreateEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateEmailTemplateInput{
		// TemplateContent: *types.EmailTemplateContent, // Required
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateContent) > 0 {
		if err := assignInputField(input, "TemplateContent", _sesv2TemplateContent); err != nil {
			log.Errorf("invalid --template-content: %s", err.Error())
			return
		}
	}
	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an export job for a data source and destination.
// You can execute this operation no more than once per second.
func sesv2_CreateExportJob(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateExportJobInput{
		// ExportDataSource: *types.ExportDataSource, // Required
		// ExportDestination: *types.ExportDestination, // Required
	}

	if len(_sesv2ExportDataSource) > 0 {
		if err := assignInputField(input, "ExportDataSource", _sesv2ExportDataSource); err != nil {
			log.Errorf("invalid --export-data-source: %s", err.Error())
			return
		}
	}
	if len(_sesv2ExportDestination) > 0 {
		if err := assignInputField(input, "ExportDestination", _sesv2ExportDestination); err != nil {
			log.Errorf("invalid --export-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an import job for a data destination.
func sesv2_CreateImportJob(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateImportJobInput{
		// ImportDataSource: *types.ImportDataSource, // Required
		// ImportDestination: *types.ImportDestination, // Required
	}

	if len(_sesv2ImportDataSource) > 0 {
		if err := assignInputField(input, "ImportDataSource", _sesv2ImportDataSource); err != nil {
			log.Errorf("invalid --import-data-source: %s", err.Error())
			return
		}
	}
	if len(_sesv2ImportDestination) > 0 {
		if err := assignInputField(input, "ImportDestination", _sesv2ImportDestination); err != nil {
			log.Errorf("invalid --import-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a multi-region endpoint (global-endpoint).
// The primary region is going to be the AWS-Region where the operation is
// executed. The secondary region has to be provided in request's parameters. From
// the data flow standpoint there is no difference between primary and secondary
// regions - sending traffic will be split equally between the two. The primary
// region is the region where the resource has been created and where it can be
// managed.
func sesv2_CreateMultiRegionEndpoint(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateMultiRegionEndpointInput{
		// Details: *types.Details, // Required
		// EndpointName: *string, // Required
	}

	if len(_sesv2Details) > 0 {
		if err := assignInputField(input, "Details", _sesv2Details); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}
	if len(_sesv2EndpointName) > 0 {
		input.EndpointName = aws.String(_sesv2EndpointName)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMultiRegionEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a tenant.
// Tenants are logical containers that group related SES resources together. Each
// tenant can have its own set of resources like email identities, configuration
// sets, and templates, along with reputation metrics and sending status. This
// helps isolate and manage email sending for different customers or business units
// within your Amazon SES API v2 account.
func sesv2_CreateTenant(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateTenantInput{
		// TenantName: *string, // Required
	}

	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a resource with a tenant.
// Resources can be email identities, configuration sets, or email templates. When
// you associate a resource with a tenant, you can use that resource when sending
// emails on behalf of that tenant.
//
// A single resource can be associated with multiple tenants, allowing for
// resource sharing across different tenants while maintaining isolation in email
// sending operations.
func sesv2_CreateTenantResourceAssociation(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.CreateTenantResourceAssociationInput{
		// ResourceArn: *string, // Required
		// TenantName: *string, // Required
	}

	if len(_sesv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_sesv2ResourceArn)
	}
	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}

	if resp, err := client.CreateTenantResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing configuration set.
// Configuration sets are groups of rules that you can apply to the emails you
// send. You apply a configuration set to an email by including a reference to the
// configuration set in the headers of the email. When you apply a configuration
// set to an email, all of the rules in that configuration set are applied to the
// email.
func sesv2_DeleteConfigurationSet(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}

	if resp, err := client.DeleteConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an event destination.
// Events include message sends, deliveries, opens, clicks, bounces, and
// complaints. Event destinations are places that you can send information about
// these events to. For example, you can send event data to Amazon EventBridge and
// associate a rule to send the event to the specified target.
func sesv2_DeleteConfigurationSetEventDestination(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2EventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_sesv2EventDestinationName)
	}

	if resp, err := client.DeleteConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a contact from a contact list.
func sesv2_DeleteContact(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteContactInput{
		// ContactListName: *string, // Required
		// EmailAddress: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}

	if resp, err := client.DeleteContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a contact list and all of the contacts on that list.
func sesv2_DeleteContactList(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteContactListInput{
		// ContactListName: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}

	if resp, err := client.DeleteContactList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing custom verification email template.
// For more information about custom verification email templates, see [Using custom verification email templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func sesv2_DeleteCustomVerificationEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteCustomVerificationEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}

	if resp, err := client.DeleteCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a dedicated IP pool.
func sesv2_DeleteDedicatedIpPool(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteDedicatedIpPoolInput{
		// PoolName: *string, // Required
	}

	if len(_sesv2PoolName) > 0 {
		input.PoolName = aws.String(_sesv2PoolName)
	}

	if resp, err := client.DeleteDedicatedIpPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an email identity. An identity can be either an email address or a
// domain name.
func sesv2_DeleteEmailIdentity(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteEmailIdentityInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}

	if resp, err := client.DeleteEmailIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified sending authorization policy for the given identity (an
// email address or a domain). This API returns successfully even if a policy with
// the specified name does not exist.
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html
func sesv2_DeleteEmailIdentityPolicy(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteEmailIdentityPolicyInput{
		// EmailIdentity: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2PolicyName) > 0 {
		input.PolicyName = aws.String(_sesv2PolicyName)
	}

	if resp, err := client.DeleteEmailIdentityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an email template.
// You can execute this operation no more than once per second.
func sesv2_DeleteEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}

	if resp, err := client.DeleteEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a multi-region endpoint (global-endpoint).
// Only multi-region endpoints (global-endpoints) whose primary region is the
// AWS-Region where operation is executed can be deleted.
func sesv2_DeleteMultiRegionEndpoint(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteMultiRegionEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_sesv2EndpointName) > 0 {
		input.EndpointName = aws.String(_sesv2EndpointName)
	}

	if resp, err := client.DeleteMultiRegionEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an email address from the suppression list for your account.
func sesv2_DeleteSuppressedDestination(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteSuppressedDestinationInput{
		// EmailAddress: *string, // Required
	}

	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}

	if resp, err := client.DeleteSuppressedDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an existing tenant.
// When you delete a tenant, its associations with resources are removed, but the
// resources themselves are not deleted.
func sesv2_DeleteTenant(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteTenantInput{
		// TenantName: *string, // Required
	}

	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}

	if resp, err := client.DeleteTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an association between a tenant and a resource.
// When you delete a tenant-resource association, the resource itself is not
// deleted, only its association with the specific tenant is removed. After
// removal, the resource will no longer be available for use with that tenant's
// email sending operations.
func sesv2_DeleteTenantResourceAssociation(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.DeleteTenantResourceAssociationInput{
		// ResourceArn: *string, // Required
		// TenantName: *string, // Required
	}

	if len(_sesv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_sesv2ResourceArn)
	}
	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}

	if resp, err := client.DeleteTenantResourceAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Obtain information about the email-sending status and capabilities of your
// Amazon SES account in the current Amazon Web Services Region.
func sesv2_GetAccount(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetAccountInput{}

	if resp, err := client.GetAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve a list of the blacklists that your dedicated IP addresses appear on.
func sesv2_GetBlacklistReports(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetBlacklistReportsInput{
		// BlacklistItemNames: []string, // Required
	}

	if len(_sesv2BlacklistItemNames) > 0 {
		input.BlacklistItemNames = append([]string(nil), _sesv2BlacklistItemNames...)
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
// Configuration sets are groups of rules that you can apply to the emails you
// send. You apply a configuration set to an email by including a reference to the
// configuration set in the headers of the email. When you apply a configuration
// set to an email, all of the rules in that configuration set are applied to the
// email.
func sesv2_GetConfigurationSet(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
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
// Events include message sends, deliveries, opens, clicks, bounces, and
// complaints. Event destinations are places that you can send information about
// these events to. For example, you can send event data to Amazon EventBridge and
// associate a rule to send the event to the specified target.
func sesv2_GetConfigurationSetEventDestinations(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetConfigurationSetEventDestinationsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}

	if resp, err := client.GetConfigurationSetEventDestinations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a contact from a contact list.
func sesv2_GetContact(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetContactInput{
		// ContactListName: *string, // Required
		// EmailAddress: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}

	if resp, err := client.GetContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns contact list metadata. It does not return any information about the
// contacts present in the list.
func sesv2_GetContactList(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetContactListInput{
		// ContactListName: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}

	if resp, err := client.GetContactList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the custom email verification template for the template name you
// specify.
//
// For more information about custom verification email templates, see [Using custom verification email templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func sesv2_GetCustomVerificationEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetCustomVerificationEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}

	if resp, err := client.GetCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about a dedicated IP address, including the name of the
// dedicated IP pool that it's associated with, as well information about the
// automatic warm-up process for the address.
func sesv2_GetDedicatedIp(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDedicatedIpInput{
		// Ip: *string, // Required
	}

	if len(_sesv2Ip) > 0 {
		input.Ip = aws.String(_sesv2Ip)
	}

	if resp, err := client.GetDedicatedIp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve information about the dedicated pool.
func sesv2_GetDedicatedIpPool(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDedicatedIpPoolInput{
		// PoolName: *string, // Required
	}

	if len(_sesv2PoolName) > 0 {
		input.PoolName = aws.String(_sesv2PoolName)
	}

	if resp, err := client.GetDedicatedIpPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the dedicated IP addresses that are associated with your Amazon Web
// Services account.
func sesv2_GetDedicatedIps(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDedicatedIpsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_sesv2PoolName) > 0 {
		input.PoolName = aws.String(_sesv2PoolName)
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

	var results []*sesv2.GetDedicatedIpsOutput
	p := sesv2.NewGetDedicatedIpsPaginator(client, input)
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
// account. When the Deliverability dashboard is enabled, you gain access to
// reputation, deliverability, and other metrics for the domains that you use to
// send email. You also gain the ability to perform predictive inbox placement
// tests.
//
// When you use the Deliverability dashboard, you pay a monthly subscription
// charge, in addition to any other fees that you accrue by using Amazon SES and
// other Amazon Web Services services. For more information about the features and
// cost of a Deliverability dashboard subscription, see [Amazon SES Pricing].
//
// [Amazon SES Pricing]: http://aws.amazon.com/ses/pricing/
func sesv2_GetDeliverabilityDashboardOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDeliverabilityDashboardOptionsInput{}

	if resp, err := client.GetDeliverabilityDashboardOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the results of a predictive inbox placement test.
func sesv2_GetDeliverabilityTestReport(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDeliverabilityTestReportInput{
		// ReportId: *string, // Required
	}

	if len(_sesv2ReportId) > 0 {
		input.ReportId = aws.String(_sesv2ReportId)
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
// the Deliverability dashboard is enabled for.
func sesv2_GetDomainDeliverabilityCampaign(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDomainDeliverabilityCampaignInput{
		// CampaignId: *string, // Required
	}

	if len(_sesv2CampaignId) > 0 {
		input.CampaignId = aws.String(_sesv2CampaignId)
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
func sesv2_GetDomainStatisticsReport(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetDomainStatisticsReportInput{
		// Domain: *string, // Required
		// EndDate: *time.Time, // Required
		// StartDate: *time.Time, // Required
	}

	if len(_sesv2Domain) > 0 {
		input.Domain = aws.String(_sesv2Domain)
	}
	if len(_sesv2EndDate) > 0 {
		if err := assignInputField(input, "EndDate", _sesv2EndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_sesv2StartDate) > 0 {
		if err := assignInputField(input, "StartDate", _sesv2StartDate); err != nil {
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

// Provides validation insights about a specific email address, including syntax
// validation, DNS record checks, mailbox existence, and other deliverability
// factors.
func sesv2_GetEmailAddressInsights(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetEmailAddressInsightsInput{
		// EmailAddress: *string, // Required
	}

	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}

	if resp, err := client.GetEmailAddressInsights(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a specific identity, including the identity's
// verification status, sending authorization policies, its DKIM authentication
// status, and its custom Mail-From settings.
func sesv2_GetEmailIdentity(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetEmailIdentityInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}

	if resp, err := client.GetEmailIdentity(context.TODO(), input); err != nil {
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
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html
func sesv2_GetEmailIdentityPolicies(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetEmailIdentityPoliciesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}

	if resp, err := client.GetEmailIdentityPolicies(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the template object (which includes the subject line, HTML part and
// text part) for the template you specify.
//
// You can execute this operation no more than once per second.
func sesv2_GetEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetEmailTemplateInput{
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}

	if resp, err := client.GetEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about an export job.
func sesv2_GetExportJob(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetExportJobInput{
		// JobId: *string, // Required
	}

	if len(_sesv2JobId) > 0 {
		input.JobId = aws.String(_sesv2JobId)
	}

	if resp, err := client.GetExportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about an import job.
func sesv2_GetImportJob(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetImportJobInput{
		// JobId: *string, // Required
	}

	if len(_sesv2JobId) > 0 {
		input.JobId = aws.String(_sesv2JobId)
	}

	if resp, err := client.GetImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about a specific message, including the from address, the
// subject, the recipient address, email tags, as well as events associated with
// the message.
//
// You can execute this operation no more than once per second.
func sesv2_GetMessageInsights(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetMessageInsightsInput{
		// MessageId: *string, // Required
	}

	if len(_sesv2MessageId) > 0 {
		input.MessageId = aws.String(_sesv2MessageId)
	}

	if resp, err := client.GetMessageInsights(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays the multi-region endpoint (global-endpoint) configuration.
// Only multi-region endpoints (global-endpoints) whose primary region is the
// AWS-Region where operation is executed can be displayed.
func sesv2_GetMultiRegionEndpoint(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetMultiRegionEndpointInput{
		// EndpointName: *string, // Required
	}

	if len(_sesv2EndpointName) > 0 {
		input.EndpointName = aws.String(_sesv2EndpointName)
	}

	if resp, err := client.GetMultiRegionEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve information about a specific reputation entity, including its
// reputation management policy, customer-managed status, Amazon Web Services
// Amazon SES-managed status, and aggregate sending status.
//
// Reputation entities represent resources in your Amazon SES account that have
// reputation tracking and management capabilities. The reputation impact reflects
// the highest impact reputation finding for the entity. Reputation findings can be
// retrieved using the ListRecommendations operation.
func sesv2_GetReputationEntity(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetReputationEntityInput{
		// ReputationEntityReference: *string, // Required
		// ReputationEntityType: types.ReputationEntityType, // Required
	}

	if len(_sesv2ReputationEntityReference) > 0 {
		input.ReputationEntityReference = aws.String(_sesv2ReputationEntityReference)
	}
	if len(_sesv2ReputationEntityType) > 0 {
		if err := assignInputField(input, "ReputationEntityType", _sesv2ReputationEntityType); err != nil {
			log.Errorf("invalid --reputation-entity-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetReputationEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific email address that's on the suppression
// list for your account.
func sesv2_GetSuppressedDestination(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetSuppressedDestinationInput{
		// EmailAddress: *string, // Required
	}

	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}

	if resp, err := client.GetSuppressedDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get information about a specific tenant, including the tenant's name, ID, ARN,
// creation timestamp, tags, and sending status.
func sesv2_GetTenant(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.GetTenantInput{
		// TenantName: *string, // Required
	}

	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}

	if resp, err := client.GetTenant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all of the configuration sets associated with your account in the current
// region.
//
// Configuration sets are groups of rules that you can apply to the emails you
// send. You apply a configuration set to an email by including a reference to the
// configuration set in the headers of the email. When you apply a configuration
// set to an email, all of the rules in that configuration set are applied to the
// email.
func sesv2_ListConfigurationSets(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListConfigurationSetsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
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

	var results []*sesv2.ListConfigurationSetsOutput
	p := sesv2.NewListConfigurationSetsPaginator(client, input)
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

// Lists all of the contact lists available.
// If your output includes a "NextToken" field with a string value, this indicates
// there may be additional contacts on the filtered list - regardless of the number
// of contacts returned.
func sesv2_ListContactLists(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListContactListsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListContactLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListContactListsOutput
	p := sesv2.NewListContactListsPaginator(client, input)
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

// Lists the contacts present in a specific contact list.
func sesv2_ListContacts(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListContactsInput{
		// ContactListName: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2Filter) > 0 {
		if err := assignInputField(input, "Filter", _sesv2Filter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListContacts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListContactsOutput
	p := sesv2.NewListContactsPaginator(client, input)
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

// Lists the existing custom verification email templates for your account in the
// current Amazon Web Services Region.
//
// For more information about custom verification email templates, see [Using custom verification email templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func sesv2_ListCustomVerificationEmailTemplates(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListCustomVerificationEmailTemplatesInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
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

	var results []*sesv2.ListCustomVerificationEmailTemplatesOutput
	p := sesv2.NewListCustomVerificationEmailTemplatesPaginator(client, input)
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

// List all of the dedicated IP pools that exist in your Amazon Web Services
// account in the current Region.
func sesv2_ListDedicatedIpPools(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListDedicatedIpPoolsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
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

	var results []*sesv2.ListDedicatedIpPoolsOutput
	p := sesv2.NewListDedicatedIpPoolsPaginator(client, input)
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
func sesv2_ListDeliverabilityTestReports(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListDeliverabilityTestReportsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
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

	var results []*sesv2.ListDeliverabilityTestReportsOutput
	p := sesv2.NewListDeliverabilityTestReportsPaginator(client, input)
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
// only if you enabled the Deliverability dashboard for the domain.
func sesv2_ListDomainDeliverabilityCampaigns(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListDomainDeliverabilityCampaignsInput{
		// EndDate: *time.Time, // Required
		// StartDate: *time.Time, // Required
		// SubscribedDomain: *string, // Required
	}

	if len(_sesv2EndDate) > 0 {
		if err := assignInputField(input, "EndDate", _sesv2EndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_sesv2StartDate) > 0 {
		if err := assignInputField(input, "StartDate", _sesv2StartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}
	if len(_sesv2SubscribedDomain) > 0 {
		input.SubscribedDomain = aws.String(_sesv2SubscribedDomain)
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
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

	var results []*sesv2.ListDomainDeliverabilityCampaignsOutput
	p := sesv2.NewListDomainDeliverabilityCampaignsPaginator(client, input)
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
// Amazon Web Services account. An identity can be either an email address or a
// domain. This operation returns identities that are verified as well as those
// that aren't. This operation returns identities that are associated with Amazon
// SES and Amazon Pinpoint.
func sesv2_ListEmailIdentities(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListEmailIdentitiesInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
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

	var results []*sesv2.ListEmailIdentitiesOutput
	p := sesv2.NewListEmailIdentitiesPaginator(client, input)
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

// Lists the email templates present in your Amazon SES account in the current
// Amazon Web Services Region.
//
// You can execute this operation no more than once per second.
func sesv2_ListEmailTemplates(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListEmailTemplatesInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListEmailTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListEmailTemplatesOutput
	p := sesv2.NewListEmailTemplatesPaginator(client, input)
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

// Lists all of the export jobs.
func sesv2_ListExportJobs(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListExportJobsInput{}

	if len(_sesv2ExportSourceType) > 0 {
		if err := assignInputField(input, "ExportSourceType", _sesv2ExportSourceType); err != nil {
			log.Errorf("invalid --export-source-type: %s", err.Error())
			return
		}
	}
	if len(_sesv2JobStatus) > 0 {
		if err := assignInputField(input, "JobStatus", _sesv2JobStatus); err != nil {
			log.Errorf("invalid --job-status: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListExportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListExportJobsOutput
	p := sesv2.NewListExportJobsPaginator(client, input)
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

// Lists all of the import jobs.
func sesv2_ListImportJobs(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListImportJobsInput{}

	if len(_sesv2ImportDestinationType) > 0 {
		if err := assignInputField(input, "ImportDestinationType", _sesv2ImportDestinationType); err != nil {
			log.Errorf("invalid --import-destination-type: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListImportJobsOutput
	p := sesv2.NewListImportJobsPaginator(client, input)
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

// List the multi-region endpoints (global-endpoints).
// Only multi-region endpoints (global-endpoints) whose primary region is the
// AWS-Region where operation is executed will be listed.
func sesv2_ListMultiRegionEndpoints(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListMultiRegionEndpointsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMultiRegionEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListMultiRegionEndpointsOutput
	p := sesv2.NewListMultiRegionEndpointsPaginator(client, input)
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

// Lists the recommendations present in your Amazon SES account in the current
// Amazon Web Services Region.
//
// You can execute this operation no more than once per second.
func sesv2_ListRecommendations(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListRecommendationsInput{}

	if len(_sesv2Filter) > 0 {
		if err := assignInputField(input, "Filter", _sesv2Filter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListRecommendationsOutput
	p := sesv2.NewListRecommendationsPaginator(client, input)
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

// List reputation entities in your Amazon SES account in the current Amazon Web
// Services Region. You can filter the results by entity type, reputation impact,
// sending status, or entity reference prefix.
//
// Reputation entities represent resources in your account that have reputation
// tracking and management capabilities. Use this operation to get an overview of
// all entities and their current reputation status.
func sesv2_ListReputationEntities(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListReputationEntitiesInput{}

	if len(_sesv2Filter) > 0 {
		if err := assignInputField(input, "Filter", _sesv2Filter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListReputationEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListReputationEntitiesOutput
	p := sesv2.NewListReputationEntitiesPaginator(client, input)
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

// List all tenants associated with a specific resource.
// This operation returns a list of tenants that are associated with the specified
// resource. This is useful for understanding which tenants are currently using a
// particular resource such as an email identity, configuration set, or email
// template.
func sesv2_ListResourceTenants(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListResourceTenantsInput{
		// ResourceArn: *string, // Required
	}

	if len(_sesv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_sesv2ResourceArn)
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceTenants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListResourceTenantsOutput
	p := sesv2.NewListResourceTenantsPaginator(client, input)
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

// Retrieves a list of email addresses that are on the suppression list for your
// account.
func sesv2_ListSuppressedDestinations(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListSuppressedDestinationsInput{}

	if len(_sesv2EndDate) > 0 {
		if err := assignInputField(input, "EndDate", _sesv2EndDate); err != nil {
			log.Errorf("invalid --end-date: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_sesv2Reasons) > 0 {
		if err := assignInputField(input, "Reasons", _sesv2Reasons); err != nil {
			log.Errorf("invalid --reasons: %s", err.Error())
			return
		}
	}
	if len(_sesv2StartDate) > 0 {
		if err := assignInputField(input, "StartDate", _sesv2StartDate); err != nil {
			log.Errorf("invalid --start-date: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSuppressedDestinations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListSuppressedDestinationsOutput
	p := sesv2.NewListSuppressedDestinationsPaginator(client, input)
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
// with a resource. Each tag consists of a required tag key and an optional
// associated tag value. A tag key is a general label that acts as a category for
// more specific tag values. A tag value acts as a descriptor within a tag key.
func sesv2_ListTagsForResource(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_sesv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_sesv2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List all resources associated with a specific tenant.
// This operation returns a list of resources (email identities, configuration
// sets, or email templates) that are associated with the specified tenant. You can
// optionally filter the results by resource type.
func sesv2_ListTenantResources(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListTenantResourcesInput{
		// TenantName: *string, // Required
	}

	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}
	if len(_sesv2Filter) > 0 {
		if err := assignInputField(input, "Filter", _sesv2Filter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTenantResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListTenantResourcesOutput
	p := sesv2.NewListTenantResourcesPaginator(client, input)
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

// List all tenants associated with your account in the current Amazon Web
// Services Region.
//
// This operation returns basic information about each tenant, such as tenant
// name, ID, ARN, and creation timestamp.
func sesv2_ListTenants(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.ListTenantsInput{}

	if len(_sesv2NextToken) > 0 {
		input.NextToken = aws.String(_sesv2NextToken)
	}
	if len(_sesv2PageSize) > 0 {
		if err := assignInputField(input, "PageSize", _sesv2PageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTenants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sesv2.ListTenantsOutput
	p := sesv2.NewListTenantsPaginator(client, input)
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

// Enable or disable the automatic warm-up feature for dedicated IP addresses.
func sesv2_PutAccountDedicatedIpWarmupAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutAccountDedicatedIpWarmupAttributesInput{}

	if len(_sesv2AutoWarmupEnabled) > 0 {
		if err := assignInputField(input, "AutoWarmupEnabled", _sesv2AutoWarmupEnabled); err != nil {
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

// Update your Amazon SES account details.
func sesv2_PutAccountDetails(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutAccountDetailsInput{
		// MailType: types.MailType, // Required
		// WebsiteURL: *string, // Required
	}

	if len(_sesv2MailType) > 0 {
		if err := assignInputField(input, "MailType", _sesv2MailType); err != nil {
			log.Errorf("invalid --mail-type: %s", err.Error())
			return
		}
	}
	if len(_sesv2WebsiteURL) > 0 {
		input.WebsiteURL = aws.String(_sesv2WebsiteURL)
	}
	if len(_sesv2AdditionalContactEmailAddresses) > 0 {
		input.AdditionalContactEmailAddresses = append([]string(nil), _sesv2AdditionalContactEmailAddresses...)
	}
	if len(_sesv2ContactLanguage) > 0 {
		if err := assignInputField(input, "ContactLanguage", _sesv2ContactLanguage); err != nil {
			log.Errorf("invalid --contact-language: %s", err.Error())
			return
		}
	}
	if len(_sesv2ProductionAccessEnabled) > 0 {
		if err := assignInputField(input, "ProductionAccessEnabled", _sesv2ProductionAccessEnabled); err != nil {
			log.Errorf("invalid --production-access-enabled: %s", err.Error())
			return
		}
	}
	if len(_sesv2UseCaseDescription) > 0 {
		input.UseCaseDescription = aws.String(_sesv2UseCaseDescription)
	}

	if resp, err := client.PutAccountDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable or disable the ability of your account to send email.
func sesv2_PutAccountSendingAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutAccountSendingAttributesInput{}

	if len(_sesv2SendingEnabled) > 0 {
		if err := assignInputField(input, "SendingEnabled", _sesv2SendingEnabled); err != nil {
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

// Change the settings for the account-level suppression list.
func sesv2_PutAccountSuppressionAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutAccountSuppressionAttributesInput{}

	if len(_sesv2SuppressedReasons) > 0 {
		if err := assignInputField(input, "SuppressedReasons", _sesv2SuppressedReasons); err != nil {
			log.Errorf("invalid --suppressed-reasons: %s", err.Error())
			return
		}
	}
	if len(_sesv2ValidationAttributes) > 0 {
		if err := assignInputField(input, "ValidationAttributes", _sesv2ValidationAttributes); err != nil {
			log.Errorf("invalid --validation-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountSuppressionAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update your Amazon SES account VDM attributes.
// You can execute this operation no more than once per second.
func sesv2_PutAccountVdmAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutAccountVdmAttributesInput{
		// VdmAttributes: *types.VdmAttributes, // Required
	}

	if len(_sesv2VdmAttributes) > 0 {
		if err := assignInputField(input, "VdmAttributes", _sesv2VdmAttributes); err != nil {
			log.Errorf("invalid --vdm-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountVdmAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate the configuration set with a MailManager archive. When you send email
// using the SendEmail or SendBulkEmail operations the message as it will be given
// to the receiving SMTP server will be archived, along with the recipient
// information.
func sesv2_PutConfigurationSetArchivingOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetArchivingOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2ArchiveArn) > 0 {
		input.ArchiveArn = aws.String(_sesv2ArchiveArn)
	}

	if resp, err := client.PutConfigurationSetArchivingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a configuration set with a dedicated IP pool. You can use dedicated
// IP pools to create groups of dedicated IP addresses for sending specific types
// of email.
func sesv2_PutConfigurationSetDeliveryOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetDeliveryOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2MaxDeliverySeconds) > 0 {
		if err := assignInputField(input, "MaxDeliverySeconds", _sesv2MaxDeliverySeconds); err != nil {
			log.Errorf("invalid --max-delivery-seconds: %s", err.Error())
			return
		}
	}
	if len(_sesv2SendingPoolName) > 0 {
		input.SendingPoolName = aws.String(_sesv2SendingPoolName)
	}
	if len(_sesv2TlsPolicy) > 0 {
		if err := assignInputField(input, "TlsPolicy", _sesv2TlsPolicy); err != nil {
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
// using a particular configuration set in a specific Amazon Web Services Region.
func sesv2_PutConfigurationSetReputationOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetReputationOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2ReputationMetricsEnabled) > 0 {
		if err := assignInputField(input, "ReputationMetricsEnabled", _sesv2ReputationMetricsEnabled); err != nil {
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
// configuration set in a specific Amazon Web Services Region.
func sesv2_PutConfigurationSetSendingOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetSendingOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2SendingEnabled) > 0 {
		if err := assignInputField(input, "SendingEnabled", _sesv2SendingEnabled); err != nil {
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

// Specify the account suppression list preferences for a configuration set.
func sesv2_PutConfigurationSetSuppressionOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetSuppressionOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2SuppressedReasons) > 0 {
		if err := assignInputField(input, "SuppressedReasons", _sesv2SuppressedReasons); err != nil {
			log.Errorf("invalid --suppressed-reasons: %s", err.Error())
			return
		}
	}
	if len(_sesv2ValidationOptions) > 0 {
		if err := assignInputField(input, "ValidationOptions", _sesv2ValidationOptions); err != nil {
			log.Errorf("invalid --validation-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationSetSuppressionOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify a custom domain to use for open and click tracking elements in email
// that you send.
func sesv2_PutConfigurationSetTrackingOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetTrackingOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2CustomRedirectDomain) > 0 {
		input.CustomRedirectDomain = aws.String(_sesv2CustomRedirectDomain)
	}
	if len(_sesv2HttpsPolicy) > 0 {
		if err := assignInputField(input, "HttpsPolicy", _sesv2HttpsPolicy); err != nil {
			log.Errorf("invalid --https-policy: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationSetTrackingOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specify VDM preferences for email that you send using the configuration set.
// You can execute this operation no more than once per second.
func sesv2_PutConfigurationSetVdmOptions(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutConfigurationSetVdmOptionsInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2VdmOptions) > 0 {
		if err := assignInputField(input, "VdmOptions", _sesv2VdmOptions); err != nil {
			log.Errorf("invalid --vdm-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutConfigurationSetVdmOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Move a dedicated IP address to an existing dedicated IP pool.
// The dedicated IP address that you specify must already exist, and must be
// associated with your Amazon Web Services account.
//
// The dedicated IP pool you specify must already exist. You can create a new pool
// by using the CreateDedicatedIpPool operation.
func sesv2_PutDedicatedIpInPool(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutDedicatedIpInPoolInput{
		// DestinationPoolName: *string, // Required
		// Ip: *string, // Required
	}

	if len(_sesv2DestinationPoolName) > 0 {
		input.DestinationPoolName = aws.String(_sesv2DestinationPoolName)
	}
	if len(_sesv2Ip) > 0 {
		input.Ip = aws.String(_sesv2Ip)
	}

	if resp, err := client.PutDedicatedIpInPool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to convert a dedicated IP pool to a different scaling mode.
// MANAGED pools cannot be converted to STANDARD scaling mode.
func sesv2_PutDedicatedIpPoolScalingAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutDedicatedIpPoolScalingAttributesInput{
		// PoolName: *string, // Required
		// ScalingMode: types.ScalingMode, // Required
	}

	if len(_sesv2PoolName) > 0 {
		input.PoolName = aws.String(_sesv2PoolName)
	}
	if len(_sesv2ScalingMode) > 0 {
		if err := assignInputField(input, "ScalingMode", _sesv2ScalingMode); err != nil {
			log.Errorf("invalid --scaling-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDedicatedIpPoolScalingAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func sesv2_PutDedicatedIpWarmupAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutDedicatedIpWarmupAttributesInput{
		// Ip: *string, // Required
		// WarmupPercentage: *int32, // Required
	}

	if len(_sesv2Ip) > 0 {
		input.Ip = aws.String(_sesv2Ip)
	}
	if len(_sesv2WarmupPercentage) > 0 {
		if err := assignInputField(input, "WarmupPercentage", _sesv2WarmupPercentage); err != nil {
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

// Enable or disable the Deliverability dashboard. When you enable the
// Deliverability dashboard, you gain access to reputation, deliverability, and
// other metrics for the domains that you use to send email. You also gain the
// ability to perform predictive inbox placement tests.
//
// When you use the Deliverability dashboard, you pay a monthly subscription
// charge, in addition to any other fees that you accrue by using Amazon SES and
// other Amazon Web Services services. For more information about the features and
// cost of a Deliverability dashboard subscription, see [Amazon SES Pricing].
//
// [Amazon SES Pricing]: http://aws.amazon.com/ses/pricing/
func sesv2_PutDeliverabilityDashboardOption(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutDeliverabilityDashboardOptionInput{
		// DashboardEnabled: bool, // Required
	}

	if len(_sesv2DashboardEnabled) > 0 {
		if err := assignInputField(input, "DashboardEnabled", _sesv2DashboardEnabled); err != nil {
			log.Errorf("invalid --dashboard-enabled: %s", err.Error())
			return
		}
	}
	if len(_sesv2SubscribedDomains) > 0 {
		if err := assignInputField(input, "SubscribedDomains", _sesv2SubscribedDomains); err != nil {
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

// Used to associate a configuration set with an email identity.
func sesv2_PutEmailIdentityConfigurationSetAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutEmailIdentityConfigurationSetAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}

	if resp, err := client.PutEmailIdentityConfigurationSetAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Used to enable or disable DKIM authentication for an email identity.
func sesv2_PutEmailIdentityDkimAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutEmailIdentityDkimAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2SigningEnabled) > 0 {
		if err := assignInputField(input, "SigningEnabled", _sesv2SigningEnabled); err != nil {
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

// Used to configure or change the DKIM authentication settings for an email
// domain identity. You can use this operation to do any of the following:
//
// - Update the signing attributes for an identity that uses Bring Your Own DKIM
// (BYODKIM).
//
// - Update the key length that should be used for Easy DKIM.
//
// - Change from using no DKIM authentication to using Easy DKIM.
//
// - Change from using no DKIM authentication to using BYODKIM.
//
// - Change from using Easy DKIM to using BYODKIM.
//
// - Change from using BYODKIM to using Easy DKIM.
func sesv2_PutEmailIdentityDkimSigningAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutEmailIdentityDkimSigningAttributesInput{
		// EmailIdentity: *string, // Required
		// SigningAttributesOrigin: types.DkimSigningAttributesOrigin, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2SigningAttributesOrigin) > 0 {
		if err := assignInputField(input, "SigningAttributesOrigin", _sesv2SigningAttributesOrigin); err != nil {
			log.Errorf("invalid --signing-attributes-origin: %s", err.Error())
			return
		}
	}
	if len(_sesv2SigningAttributes) > 0 {
		if err := assignInputField(input, "SigningAttributes", _sesv2SigningAttributes); err != nil {
			log.Errorf("invalid --signing-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutEmailIdentityDkimSigningAttributes(context.TODO(), input); err != nil {
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
// If the value is true , you receive email notifications when bounce or complaint
// events occur. These notifications are sent to the address that you specified in
// the Return-Path header of the original email.
//
// You're required to have a method of tracking bounces and complaints. If you
// haven't set up another mechanism for receiving bounce or complaint notifications
// (for example, by setting up an event destination), you receive an email
// notification when these events occur (even if this setting is disabled).
func sesv2_PutEmailIdentityFeedbackAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutEmailIdentityFeedbackAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2EmailForwardingEnabled) > 0 {
		if err := assignInputField(input, "EmailForwardingEnabled", _sesv2EmailForwardingEnabled); err != nil {
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
func sesv2_PutEmailIdentityMailFromAttributes(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutEmailIdentityMailFromAttributesInput{
		// EmailIdentity: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2BehaviorOnMxFailure) > 0 {
		if err := assignInputField(input, "BehaviorOnMxFailure", _sesv2BehaviorOnMxFailure); err != nil {
			log.Errorf("invalid --behavior-on-mx-failure: %s", err.Error())
			return
		}
	}
	if len(_sesv2MailFromDomain) > 0 {
		input.MailFromDomain = aws.String(_sesv2MailFromDomain)
	}

	if resp, err := client.PutEmailIdentityMailFromAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an email address to the suppression list for your account.
func sesv2_PutSuppressedDestination(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.PutSuppressedDestinationInput{
		// EmailAddress: *string, // Required
		// Reason: types.SuppressionListReason, // Required
	}

	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}
	if len(_sesv2Reason) > 0 {
		if err := assignInputField(input, "Reason", _sesv2Reason); err != nil {
			log.Errorf("invalid --reason: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutSuppressedDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Composes an email message to multiple destinations.
func sesv2_SendBulkEmail(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.SendBulkEmailInput{
		// BulkEmailEntries: []types.BulkEmailEntry, // Required
		// DefaultContent: *types.BulkEmailContent, // Required
	}

	if len(_sesv2BulkEmailEntries) > 0 {
		if err := assignInputField(input, "BulkEmailEntries", _sesv2BulkEmailEntries); err != nil {
			log.Errorf("invalid --bulk-email-entries: %s", err.Error())
			return
		}
	}
	if len(_sesv2DefaultContent) > 0 {
		if err := assignInputField(input, "DefaultContent", _sesv2DefaultContent); err != nil {
			log.Errorf("invalid --default-content: %s", err.Error())
			return
		}
	}
	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2DefaultEmailTags) > 0 {
		if err := assignInputField(input, "DefaultEmailTags", _sesv2DefaultEmailTags); err != nil {
			log.Errorf("invalid --default-email-tags: %s", err.Error())
			return
		}
	}
	if len(_sesv2EndpointId) > 0 {
		input.EndpointId = aws.String(_sesv2EndpointId)
	}
	if len(_sesv2FeedbackForwardingEmailAddress) > 0 {
		input.FeedbackForwardingEmailAddress = aws.String(_sesv2FeedbackForwardingEmailAddress)
	}
	if len(_sesv2FeedbackForwardingEmailAddressIdentityArn) > 0 {
		input.FeedbackForwardingEmailAddressIdentityArn = aws.String(_sesv2FeedbackForwardingEmailAddressIdentityArn)
	}
	if len(_sesv2FromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesv2FromEmailAddress)
	}
	if len(_sesv2FromEmailAddressIdentityArn) > 0 {
		input.FromEmailAddressIdentityArn = aws.String(_sesv2FromEmailAddressIdentityArn)
	}
	if len(_sesv2ReplyToAddresses) > 0 {
		input.ReplyToAddresses = append([]string(nil), _sesv2ReplyToAddresses...)
	}
	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}

	if resp, err := client.SendBulkEmail(context.TODO(), input); err != nil {
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
// email templates, see [Using custom verification email templates]in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func sesv2_SendCustomVerificationEmail(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.SendCustomVerificationEmailInput{
		// EmailAddress: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}
	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}
	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}

	if resp, err := client.SendCustomVerificationEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an email message. You can use the Amazon SES API v2 to send the following
// types of messages:
//
// - Simple – A standard email message. When you create this type of message,
// you specify the sender, the recipient, and the message body, and Amazon SES
// assembles the message for you.
//
// - Raw – A raw, MIME-formatted email message. When you send this type of
// email, you have to specify all of the message headers, as well as the message
// body. You can use this message type to send messages that contain attachments.
// The message that you specify has to be a valid MIME message.
//
// - Templated – A message that contains personalization tags. When you send
// this type of email, Amazon SES API v2 automatically replaces the tags with
// values that you specify.
func sesv2_SendEmail(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.SendEmailInput{
		// Content: *types.EmailContent, // Required
	}

	if len(_sesv2Content) > 0 {
		if err := assignInputField(input, "Content", _sesv2Content); err != nil {
			log.Errorf("invalid --content: %s", err.Error())
			return
		}
	}
	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2Destination) > 0 {
		if err := assignInputField(input, "Destination", _sesv2Destination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_sesv2EmailTags) > 0 {
		if err := assignInputField(input, "EmailTags", _sesv2EmailTags); err != nil {
			log.Errorf("invalid --email-tags: %s", err.Error())
			return
		}
	}
	if len(_sesv2EndpointId) > 0 {
		input.EndpointId = aws.String(_sesv2EndpointId)
	}
	if len(_sesv2FeedbackForwardingEmailAddress) > 0 {
		input.FeedbackForwardingEmailAddress = aws.String(_sesv2FeedbackForwardingEmailAddress)
	}
	if len(_sesv2FeedbackForwardingEmailAddressIdentityArn) > 0 {
		input.FeedbackForwardingEmailAddressIdentityArn = aws.String(_sesv2FeedbackForwardingEmailAddressIdentityArn)
	}
	if len(_sesv2FromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesv2FromEmailAddress)
	}
	if len(_sesv2FromEmailAddressIdentityArn) > 0 {
		input.FromEmailAddressIdentityArn = aws.String(_sesv2FromEmailAddressIdentityArn)
	}
	if len(_sesv2ListManagementOptions) > 0 {
		if err := assignInputField(input, "ListManagementOptions", _sesv2ListManagementOptions); err != nil {
			log.Errorf("invalid --list-management-options: %s", err.Error())
			return
		}
	}
	if len(_sesv2ReplyToAddresses) > 0 {
		input.ReplyToAddresses = append([]string(nil), _sesv2ReplyToAddresses...)
	}
	if len(_sesv2TenantName) > 0 {
		input.TenantName = aws.String(_sesv2TenantName)
	}

	if resp, err := client.SendEmail(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add one or more tags (keys and values) to a specified resource. A tag is a
// label that you optionally define and associate with a resource. Tags can help
// you categorize and manage resources in different ways, such as by purpose,
// owner, environment, or other criteria. A resource can have as many as 50 tags.
//
// Each tag consists of a required tag key and an associated tag value, both of
// which you define. A tag key is a general label that acts as a category for more
// specific tag values. A tag value acts as a descriptor within a tag key.
func sesv2_TagResource(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_sesv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_sesv2ResourceArn)
	}
	if len(_sesv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _sesv2Tags); err != nil {
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

// Creates a preview of the MIME content of an email when provided with a template
// and a set of replacement data.
//
// You can execute this operation no more than once per second.
func sesv2_TestRenderEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.TestRenderEmailTemplateInput{
		// TemplateData: *string, // Required
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateData) > 0 {
		input.TemplateData = aws.String(_sesv2TemplateData)
	}
	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}

	if resp, err := client.TestRenderEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove one or more tags (keys and values) from a specified resource.
func sesv2_UntagResource(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_sesv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_sesv2ResourceArn)
	}
	if len(_sesv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _sesv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the configuration of an event destination for a configuration set.
// Events include message sends, deliveries, opens, clicks, bounces, and
// complaints. Event destinations are places that you can send information about
// these events to. For example, you can send event data to Amazon EventBridge and
// associate a rule to send the event to the specified target.
func sesv2_UpdateConfigurationSetEventDestination(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateConfigurationSetEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestination: *types.EventDestinationDefinition, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_sesv2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_sesv2ConfigurationSetName)
	}
	if len(_sesv2EventDestination) > 0 {
		if err := assignInputField(input, "EventDestination", _sesv2EventDestination); err != nil {
			log.Errorf("invalid --event-destination: %s", err.Error())
			return
		}
	}
	if len(_sesv2EventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_sesv2EventDestinationName)
	}

	if resp, err := client.UpdateConfigurationSetEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a contact's preferences for a list.
// You must specify all existing topic preferences in the TopicPreferences object,
// not just the ones that need updating; otherwise, all your existing preferences
// will be removed.
func sesv2_UpdateContact(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateContactInput{
		// ContactListName: *string, // Required
		// EmailAddress: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2EmailAddress) > 0 {
		input.EmailAddress = aws.String(_sesv2EmailAddress)
	}
	if len(_sesv2AttributesData) > 0 {
		input.AttributesData = aws.String(_sesv2AttributesData)
	}
	if len(_sesv2TopicPreferences) > 0 {
		if err := assignInputField(input, "TopicPreferences", _sesv2TopicPreferences); err != nil {
			log.Errorf("invalid --topic-preferences: %s", err.Error())
			return
		}
	}
	if len(_sesv2UnsubscribeAll) > 0 {
		if err := assignInputField(input, "UnsubscribeAll", _sesv2UnsubscribeAll); err != nil {
			log.Errorf("invalid --unsubscribe-all: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates contact list metadata. This operation does a complete replacement.
func sesv2_UpdateContactList(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateContactListInput{
		// ContactListName: *string, // Required
	}

	if len(_sesv2ContactListName) > 0 {
		input.ContactListName = aws.String(_sesv2ContactListName)
	}
	if len(_sesv2Description) > 0 {
		input.Description = aws.String(_sesv2Description)
	}
	if len(_sesv2Topics) > 0 {
		if err := assignInputField(input, "Topics", _sesv2Topics); err != nil {
			log.Errorf("invalid --topics: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateContactList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing custom verification email template.
// For more information about custom verification email templates, see [Using custom verification email templates] in the
// Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second.
//
// [Using custom verification email templates]: https://docs.aws.amazon.com/ses/latest/dg/creating-identities.html#send-email-verify-address-custom
func sesv2_UpdateCustomVerificationEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateCustomVerificationEmailTemplateInput{
		// FailureRedirectionURL: *string, // Required
		// FromEmailAddress: *string, // Required
		// SuccessRedirectionURL: *string, // Required
		// TemplateContent: *string, // Required
		// TemplateName: *string, // Required
		// TemplateSubject: *string, // Required
	}

	if len(_sesv2FailureRedirectionURL) > 0 {
		input.FailureRedirectionURL = aws.String(_sesv2FailureRedirectionURL)
	}
	if len(_sesv2FromEmailAddress) > 0 {
		input.FromEmailAddress = aws.String(_sesv2FromEmailAddress)
	}
	if len(_sesv2SuccessRedirectionURL) > 0 {
		input.SuccessRedirectionURL = aws.String(_sesv2SuccessRedirectionURL)
	}
	if len(_sesv2TemplateContent) > 0 {
		input.TemplateContent = aws.String(_sesv2TemplateContent)
	}
	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}
	if len(_sesv2TemplateSubject) > 0 {
		input.TemplateSubject = aws.String(_sesv2TemplateSubject)
	}

	if resp, err := client.UpdateCustomVerificationEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified sending authorization policy for the given identity (an
// email address or a domain). This API returns successfully even if a policy with
// the specified name does not exist.
//
// This API is for the identity owner only. If you have not verified the identity,
// this API will return an error.
//
// Sending authorization is a feature that enables an identity owner to authorize
// other senders to use its identities. For information about using sending
// authorization, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html
func sesv2_UpdateEmailIdentityPolicy(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateEmailIdentityPolicyInput{
		// EmailIdentity: *string, // Required
		// Policy: *string, // Required
		// PolicyName: *string, // Required
	}

	if len(_sesv2EmailIdentity) > 0 {
		input.EmailIdentity = aws.String(_sesv2EmailIdentity)
	}
	if len(_sesv2Policy) > 0 {
		input.Policy = aws.String(_sesv2Policy)
	}
	if len(_sesv2PolicyName) > 0 {
		input.PolicyName = aws.String(_sesv2PolicyName)
	}

	if resp, err := client.UpdateEmailIdentityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an email template. Email templates enable you to send personalized
// email to one or more destinations in a single API operation. For more
// information, see the [Amazon SES Developer Guide].
//
// You can execute this operation no more than once per second.
//
// [Amazon SES Developer Guide]: https://docs.aws.amazon.com/ses/latest/DeveloperGuide/send-personalized-email-api.html
func sesv2_UpdateEmailTemplate(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateEmailTemplateInput{
		// TemplateContent: *types.EmailTemplateContent, // Required
		// TemplateName: *string, // Required
	}

	if len(_sesv2TemplateContent) > 0 {
		if err := assignInputField(input, "TemplateContent", _sesv2TemplateContent); err != nil {
			log.Errorf("invalid --template-content: %s", err.Error())
			return
		}
	}
	if len(_sesv2TemplateName) > 0 {
		input.TemplateName = aws.String(_sesv2TemplateName)
	}

	if resp, err := client.UpdateEmailTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the customer-managed sending status for a reputation entity. This allows
// you to enable, disable, or reinstate sending for the entity.
//
// The customer-managed status works in conjunction with the Amazon Web Services
// Amazon SES-managed status to determine the overall sending capability. When you
// update the customer-managed status, the Amazon Web Services Amazon SES-managed
// status remains unchanged. If Amazon Web Services Amazon SES has disabled the
// entity, it will not be allowed to send regardless of the customer-managed status
// setting. When you reinstate an entity through the customer-managed status, it
// can continue sending only if the Amazon Web Services Amazon SES-managed status
// also permits sending, even if there are active reputation findings, until the
// findings are resolved or new violations occur.
func sesv2_UpdateReputationEntityCustomerManagedStatus(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateReputationEntityCustomerManagedStatusInput{
		// ReputationEntityReference: *string, // Required
		// ReputationEntityType: types.ReputationEntityType, // Required
		// SendingStatus: types.SendingStatus, // Required
	}

	if len(_sesv2ReputationEntityReference) > 0 {
		input.ReputationEntityReference = aws.String(_sesv2ReputationEntityReference)
	}
	if len(_sesv2ReputationEntityType) > 0 {
		if err := assignInputField(input, "ReputationEntityType", _sesv2ReputationEntityType); err != nil {
			log.Errorf("invalid --reputation-entity-type: %s", err.Error())
			return
		}
	}
	if len(_sesv2SendingStatus) > 0 {
		if err := assignInputField(input, "SendingStatus", _sesv2SendingStatus); err != nil {
			log.Errorf("invalid --sending-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReputationEntityCustomerManagedStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the reputation management policy for a reputation entity. The policy
// determines how the entity responds to reputation findings, such as automatically
// pausing sending when certain thresholds are exceeded.
//
// Reputation management policies are Amazon Web Services Amazon SES-managed
// (predefined policies). You can select from none, standard, and strict policies.
func sesv2_UpdateReputationEntityPolicy(cfg aws.Config, client *sesv2.Client) {
	input := &sesv2.UpdateReputationEntityPolicyInput{
		// ReputationEntityPolicy: *string, // Required
		// ReputationEntityReference: *string, // Required
		// ReputationEntityType: types.ReputationEntityType, // Required
	}

	if len(_sesv2ReputationEntityPolicy) > 0 {
		input.ReputationEntityPolicy = aws.String(_sesv2ReputationEntityPolicy)
	}
	if len(_sesv2ReputationEntityReference) > 0 {
		input.ReputationEntityReference = aws.String(_sesv2ReputationEntityReference)
	}
	if len(_sesv2ReputationEntityType) > 0 {
		if err := assignInputField(input, "ReputationEntityType", _sesv2ReputationEntityType); err != nil {
			log.Errorf("invalid --reputation-entity-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateReputationEntityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_sesv2Cmd)
	_sesv2Cmd.Flags().SortFlags = false

	_sesv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_sesv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_sesv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_sesv2Cmd.Flags().StringSliceVarP(&_sesv2AdditionalContactEmailAddresses, "additional-contact-email-addresses", "", nil, "Additional Contact Email Addresses")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ArchiveArn, "archive-arn", "", "", "Archive ARN")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ArchivingOptions, "archiving-options", "", "", "Archiving Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2AttributesData, "attributes-data", "", "", "Attributes Data")
	_sesv2Cmd.Flags().StringVarP(&_sesv2AutoWarmupEnabled, "auto-warmup-enabled", "", "", "Auto Warmup Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2BehaviorOnMxFailure, "behavior-on-mx-failure", "", "", "Behavior On Mx Failure")
	_sesv2Cmd.Flags().StringSliceVarP(&_sesv2BlacklistItemNames, "blacklist-item-names", "", nil, "Blacklist Item Names")
	_sesv2Cmd.Flags().StringVarP(&_sesv2BulkEmailEntries, "bulk-email-entries", "", "", "Bulk Email Entries")
	_sesv2Cmd.Flags().StringVarP(&_sesv2CampaignId, "campaign-id", "", "", "Campaign ID")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ConfigurationSetName, "configuration-set-name", "", "", "Configuration Set Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ContactLanguage, "contact-language", "", "", "Contact Language")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ContactListName, "contact-list-name", "", "", "Contact List Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Content, "content", "", "", "Content")
	_sesv2Cmd.Flags().StringVarP(&_sesv2CustomRedirectDomain, "custom-redirect-domain", "", "", "Custom Redirect Domain")
	_sesv2Cmd.Flags().StringVarP(&_sesv2DashboardEnabled, "dashboard-enabled", "", "", "Dashboard Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2DefaultContent, "default-content", "", "", "Default Content")
	_sesv2Cmd.Flags().StringVarP(&_sesv2DefaultEmailTags, "default-email-tags", "", "", "Default Email Tags")
	_sesv2Cmd.Flags().StringVarP(&_sesv2DeliveryOptions, "delivery-options", "", "", "Delivery Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Description, "description", "", "", "Description")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Destination, "destination", "", "", "Destination")
	_sesv2Cmd.Flags().StringVarP(&_sesv2DestinationPoolName, "destination-pool-name", "", "", "Destination Pool Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Details, "details", "", "", "Details")
	_sesv2Cmd.Flags().StringVarP(&_sesv2DkimSigningAttributes, "dkim-signing-attributes", "", "", "Dkim Signing Attributes")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Domain, "domain", "", "", "Domain")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EmailAddress, "email-address", "", "", "Email Address")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EmailForwardingEnabled, "email-forwarding-enabled", "", "", "Email Forwarding Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EmailIdentity, "email-identity", "", "", "Email Identity")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EmailTags, "email-tags", "", "", "Email Tags")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EndDate, "end-date", "", "", "End Date")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EndpointId, "endpoint-id", "", "", "Endpoint ID")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EndpointName, "endpoint-name", "", "", "Endpoint Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EventDestination, "event-destination", "", "", "Event Destination")
	_sesv2Cmd.Flags().StringVarP(&_sesv2EventDestinationName, "event-destination-name", "", "", "Event Destination Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ExportDataSource, "export-data-source", "", "", "Export Data Source")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ExportDestination, "export-destination", "", "", "Export Destination")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ExportSourceType, "export-source-type", "", "", "Export Source Type")
	_sesv2Cmd.Flags().StringVarP(&_sesv2FailureRedirectionURL, "failure-redirection-url", "", "", "Failure Redirection URL")
	_sesv2Cmd.Flags().StringVarP(&_sesv2FeedbackForwardingEmailAddress, "feedback-forwarding-email-address", "", "", "Feedback Forwarding Email Address")
	_sesv2Cmd.Flags().StringVarP(&_sesv2FeedbackForwardingEmailAddressIdentityArn, "feedback-forwarding-email-address-identity-arn", "", "", "Feedback Forwarding Email Address Identity ARN")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Filter, "filter", "", "", "Filter")
	_sesv2Cmd.Flags().StringVarP(&_sesv2FromEmailAddress, "from-email-address", "", "", "From Email Address")
	_sesv2Cmd.Flags().StringVarP(&_sesv2FromEmailAddressIdentityArn, "from-email-address-identity-arn", "", "", "From Email Address Identity ARN")
	_sesv2Cmd.Flags().StringVarP(&_sesv2HttpsPolicy, "https-policy", "", "", "HTTPS Policy")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ImportDataSource, "import-data-source", "", "", "Import Data Source")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ImportDestination, "import-destination", "", "", "Import Destination")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ImportDestinationType, "import-destination-type", "", "", "Import Destination Type")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Ip, "ip", "", "", "IP")
	_sesv2Cmd.Flags().StringVarP(&_sesv2JobId, "job-id", "", "", "Job ID")
	_sesv2Cmd.Flags().StringVarP(&_sesv2JobStatus, "job-status", "", "", "Job Status")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ListManagementOptions, "list-management-options", "", "", "List Management Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2MailFromDomain, "mail-from-domain", "", "", "Mail From Domain")
	_sesv2Cmd.Flags().StringVarP(&_sesv2MailType, "mail-type", "", "", "Mail Type")
	_sesv2Cmd.Flags().StringVarP(&_sesv2MaxDeliverySeconds, "max-delivery-seconds", "", "", "Max Delivery Seconds")
	_sesv2Cmd.Flags().StringVarP(&_sesv2MessageId, "message-id", "", "", "Message ID")
	_sesv2Cmd.Flags().StringVarP(&_sesv2NextToken, "next-token", "", "", "Next Token")
	_sesv2Cmd.Flags().StringVarP(&_sesv2PageSize, "page-size", "", "", "Page Size")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Policy, "policy", "", "", "Policy")
	_sesv2Cmd.Flags().StringVarP(&_sesv2PolicyName, "policy-name", "", "", "Policy Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2PoolName, "pool-name", "", "", "Pool Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ProductionAccessEnabled, "production-access-enabled", "", "", "Production Access Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Queries, "queries", "", "", "Queries")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Reason, "reason", "", "", "Reason")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Reasons, "reasons", "", "", "Reasons")
	_sesv2Cmd.Flags().StringSliceVarP(&_sesv2ReplyToAddresses, "reply-to-addresses", "", nil, "Reply To Addresses")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReportId, "report-id", "", "", "Report ID")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReportName, "report-name", "", "", "Report Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReputationEntityPolicy, "reputation-entity-policy", "", "", "Reputation Entity Policy")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReputationEntityReference, "reputation-entity-reference", "", "", "Reputation Entity Reference")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReputationEntityType, "reputation-entity-type", "", "", "Reputation Entity Type")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReputationMetricsEnabled, "reputation-metrics-enabled", "", "", "Reputation Metrics Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ReputationOptions, "reputation-options", "", "", "Reputation Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ScalingMode, "scaling-mode", "", "", "Scaling Mode")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SendingEnabled, "sending-enabled", "", "", "Sending Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SendingOptions, "sending-options", "", "", "Sending Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SendingPoolName, "sending-pool-name", "", "", "Sending Pool Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SendingStatus, "sending-status", "", "", "Sending Status")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SigningAttributes, "signing-attributes", "", "", "Signing Attributes")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SigningAttributesOrigin, "signing-attributes-origin", "", "", "Signing Attributes Origin")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SigningEnabled, "signing-enabled", "", "", "Signing Enabled")
	_sesv2Cmd.Flags().StringVarP(&_sesv2StartDate, "start-date", "", "", "Start Date")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SubscribedDomain, "subscribed-domain", "", "", "Subscribed Domain")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SubscribedDomains, "subscribed-domains", "", "", "Subscribed Domains")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SuccessRedirectionURL, "success-redirection-url", "", "", "Success Redirection URL")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SuppressedReasons, "suppressed-reasons", "", "", "Suppressed Reasons")
	_sesv2Cmd.Flags().StringVarP(&_sesv2SuppressionOptions, "suppression-options", "", "", "Suppression Options")
	_sesv2Cmd.Flags().StringSliceVarP(&_sesv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Tags, "tags", "", "", "Tags")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TemplateContent, "template-content", "", "", "Template Content")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TemplateData, "template-data", "", "", "Template Data")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TemplateName, "template-name", "", "", "Template Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TemplateSubject, "template-subject", "", "", "Template Subject")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TenantName, "tenant-name", "", "", "Tenant Name")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TlsPolicy, "tls-policy", "", "", "TLS Policy")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TopicPreferences, "topic-preferences", "", "", "Topic Preferences")
	_sesv2Cmd.Flags().StringVarP(&_sesv2Topics, "topics", "", "", "Topics")
	_sesv2Cmd.Flags().StringVarP(&_sesv2TrackingOptions, "tracking-options", "", "", "Tracking Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2UnsubscribeAll, "unsubscribe-all", "", "", "Unsubscribe All")
	_sesv2Cmd.Flags().StringVarP(&_sesv2UseCaseDescription, "use-case-description", "", "", "Use Case Description")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ValidationAttributes, "validation-attributes", "", "", "Validation Attributes")
	_sesv2Cmd.Flags().StringVarP(&_sesv2ValidationOptions, "validation-options", "", "", "Validation Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2VdmAttributes, "vdm-attributes", "", "", "Vdm Attributes")
	_sesv2Cmd.Flags().StringVarP(&_sesv2VdmOptions, "vdm-options", "", "", "Vdm Options")
	_sesv2Cmd.Flags().StringVarP(&_sesv2WarmupPercentage, "warmup-percentage", "", "", "Warmup Percentage")
	_sesv2Cmd.Flags().StringVarP(&_sesv2WebsiteURL, "website-url", "", "", "Website URL")

	_sesv2Cmd.Flags().BoolVarP(&_sesv2BatchGetMetricData, "batch-get-metric-data", "", false, "Batch Get Metric Data")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CancelExportJob, "cancel-export-job", "", false, "Cancel Export Job")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateConfigurationSet, "create-configuration-set", "", false, "Create Configuration Set")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateConfigurationSetEventDestination, "create-configuration-set-event-destination", "", false, "Create Configuration Set Event Destination")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateContact, "create-contact", "", false, "Create Contact")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateContactList, "create-contact-list", "", false, "Create Contact List")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateCustomVerificationEmailTemplate, "create-custom-verification-email-template", "", false, "Create Custom Verification Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateDedicatedIpPool, "create-dedicated-ip-pool", "", false, "Create Dedicated IP Pool")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateDeliverabilityTestReport, "create-deliverability-test-report", "", false, "Create Deliverability Test Report")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateEmailIdentity, "create-email-identity", "", false, "Create Email Identity")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateEmailIdentityPolicy, "create-email-identity-policy", "", false, "Create Email Identity Policy")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateEmailTemplate, "create-email-template", "", false, "Create Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateExportJob, "create-export-job", "", false, "Create Export Job")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateImportJob, "create-import-job", "", false, "Create Import Job")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateMultiRegionEndpoint, "create-multi-region-endpoint", "", false, "Create Multi Region Endpoint")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateTenant, "create-tenant", "", false, "Create Tenant")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2CreateTenantResourceAssociation, "create-tenant-resource-association", "", false, "Create Tenant Resource Association")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteConfigurationSet, "delete-configuration-set", "", false, "Delete Configuration Set")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteConfigurationSetEventDestination, "delete-configuration-set-event-destination", "", false, "Delete Configuration Set Event Destination")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteContact, "delete-contact", "", false, "Delete Contact")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteContactList, "delete-contact-list", "", false, "Delete Contact List")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteCustomVerificationEmailTemplate, "delete-custom-verification-email-template", "", false, "Delete Custom Verification Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteDedicatedIpPool, "delete-dedicated-ip-pool", "", false, "Delete Dedicated IP Pool")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteEmailIdentity, "delete-email-identity", "", false, "Delete Email Identity")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteEmailIdentityPolicy, "delete-email-identity-policy", "", false, "Delete Email Identity Policy")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteEmailTemplate, "delete-email-template", "", false, "Delete Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteMultiRegionEndpoint, "delete-multi-region-endpoint", "", false, "Delete Multi Region Endpoint")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteSuppressedDestination, "delete-suppressed-destination", "", false, "Delete Suppressed Destination")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteTenant, "delete-tenant", "", false, "Delete Tenant")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2DeleteTenantResourceAssociation, "delete-tenant-resource-association", "", false, "Delete Tenant Resource Association")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetAccount, "get-account", "", false, "Get Account")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetBlacklistReports, "get-blacklist-reports", "", false, "Get Blacklist Reports")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetConfigurationSet, "get-configuration-set", "", false, "Get Configuration Set")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetConfigurationSetEventDestinations, "get-configuration-set-event-destinations", "", false, "Get Configuration Set Event Destinations")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetContact, "get-contact", "", false, "Get Contact")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetContactList, "get-contact-list", "", false, "Get Contact List")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetCustomVerificationEmailTemplate, "get-custom-verification-email-template", "", false, "Get Custom Verification Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDedicatedIp, "get-dedicated-ip", "", false, "Get Dedicated IP")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDedicatedIpPool, "get-dedicated-ip-pool", "", false, "Get Dedicated IP Pool")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDedicatedIps, "get-dedicated-ips", "", false, "Get Dedicated Ips")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDeliverabilityDashboardOptions, "get-deliverability-dashboard-options", "", false, "Get Deliverability Dashboard Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDeliverabilityTestReport, "get-deliverability-test-report", "", false, "Get Deliverability Test Report")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDomainDeliverabilityCampaign, "get-domain-deliverability-campaign", "", false, "Get Domain Deliverability Campaign")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetDomainStatisticsReport, "get-domain-statistics-report", "", false, "Get Domain Statistics Report")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetEmailAddressInsights, "get-email-address-insights", "", false, "Get Email Address Insights")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetEmailIdentity, "get-email-identity", "", false, "Get Email Identity")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetEmailIdentityPolicies, "get-email-identity-policies", "", false, "Get Email Identity Policies")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetEmailTemplate, "get-email-template", "", false, "Get Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetExportJob, "get-export-job", "", false, "Get Export Job")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetImportJob, "get-import-job", "", false, "Get Import Job")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetMessageInsights, "get-message-insights", "", false, "Get Message Insights")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetMultiRegionEndpoint, "get-multi-region-endpoint", "", false, "Get Multi Region Endpoint")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetReputationEntity, "get-reputation-entity", "", false, "Get Reputation Entity")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetSuppressedDestination, "get-suppressed-destination", "", false, "Get Suppressed Destination")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2GetTenant, "get-tenant", "", false, "Get Tenant")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListConfigurationSets, "list-configuration-sets", "", false, "List Configuration Sets")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListContactLists, "list-contact-lists", "", false, "List Contact Lists")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListContacts, "list-contacts", "", false, "List Contacts")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListCustomVerificationEmailTemplates, "list-custom-verification-email-templates", "", false, "List Custom Verification Email Templates")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListDedicatedIpPools, "list-dedicated-ip-pools", "", false, "List Dedicated IP Pools")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListDeliverabilityTestReports, "list-deliverability-test-reports", "", false, "List Deliverability Test Reports")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListDomainDeliverabilityCampaigns, "list-domain-deliverability-campaigns", "", false, "List Domain Deliverability Campaigns")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListEmailIdentities, "list-email-identities", "", false, "List Email Identities")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListEmailTemplates, "list-email-templates", "", false, "List Email Templates")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListExportJobs, "list-export-jobs", "", false, "List Export Jobs")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListImportJobs, "list-import-jobs", "", false, "List Import Jobs")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListMultiRegionEndpoints, "list-multi-region-endpoints", "", false, "List Multi Region Endpoints")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListRecommendations, "list-recommendations", "", false, "List Recommendations")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListReputationEntities, "list-reputation-entities", "", false, "List Reputation Entities")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListResourceTenants, "list-resource-tenants", "", false, "List Resource Tenants")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListSuppressedDestinations, "list-suppressed-destinations", "", false, "List Suppressed Destinations")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListTenantResources, "list-tenant-resources", "", false, "List Tenant Resources")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2ListTenants, "list-tenants", "", false, "List Tenants")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutAccountDedicatedIpWarmupAttributes, "put-account-dedicated-ip-warmup-attributes", "", false, "Put Account Dedicated IP Warmup Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutAccountDetails, "put-account-details", "", false, "Put Account Details")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutAccountSendingAttributes, "put-account-sending-attributes", "", false, "Put Account Sending Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutAccountSuppressionAttributes, "put-account-suppression-attributes", "", false, "Put Account Suppression Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutAccountVdmAttributes, "put-account-vdm-attributes", "", false, "Put Account Vdm Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetArchivingOptions, "put-configuration-set-archiving-options", "", false, "Put Configuration Set Archiving Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetDeliveryOptions, "put-configuration-set-delivery-options", "", false, "Put Configuration Set Delivery Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetReputationOptions, "put-configuration-set-reputation-options", "", false, "Put Configuration Set Reputation Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetSendingOptions, "put-configuration-set-sending-options", "", false, "Put Configuration Set Sending Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetSuppressionOptions, "put-configuration-set-suppression-options", "", false, "Put Configuration Set Suppression Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetTrackingOptions, "put-configuration-set-tracking-options", "", false, "Put Configuration Set Tracking Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutConfigurationSetVdmOptions, "put-configuration-set-vdm-options", "", false, "Put Configuration Set Vdm Options")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutDedicatedIpInPool, "put-dedicated-ip-in-pool", "", false, "Put Dedicated IP In Pool")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutDedicatedIpPoolScalingAttributes, "put-dedicated-ip-pool-scaling-attributes", "", false, "Put Dedicated IP Pool Scaling Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutDedicatedIpWarmupAttributes, "put-dedicated-ip-warmup-attributes", "", false, "Put Dedicated IP Warmup Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutDeliverabilityDashboardOption, "put-deliverability-dashboard-option", "", false, "Put Deliverability Dashboard Option")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutEmailIdentityConfigurationSetAttributes, "put-email-identity-configuration-set-attributes", "", false, "Put Email Identity Configuration Set Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutEmailIdentityDkimAttributes, "put-email-identity-dkim-attributes", "", false, "Put Email Identity Dkim Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutEmailIdentityDkimSigningAttributes, "put-email-identity-dkim-signing-attributes", "", false, "Put Email Identity Dkim Signing Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutEmailIdentityFeedbackAttributes, "put-email-identity-feedback-attributes", "", false, "Put Email Identity Feedback Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutEmailIdentityMailFromAttributes, "put-email-identity-mail-from-attributes", "", false, "Put Email Identity Mail From Attributes")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2PutSuppressedDestination, "put-suppressed-destination", "", false, "Put Suppressed Destination")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2SendBulkEmail, "send-bulk-email", "", false, "Send Bulk Email")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2SendCustomVerificationEmail, "send-custom-verification-email", "", false, "Send Custom Verification Email")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2SendEmail, "send-email", "", false, "Send Email")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2TagResource, "tag-resource", "", false, "Tag Resource")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2TestRenderEmailTemplate, "test-render-email-template", "", false, "Test Render Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateConfigurationSetEventDestination, "update-configuration-set-event-destination", "", false, "Update Configuration Set Event Destination")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateContact, "update-contact", "", false, "Update Contact")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateContactList, "update-contact-list", "", false, "Update Contact List")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateCustomVerificationEmailTemplate, "update-custom-verification-email-template", "", false, "Update Custom Verification Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateEmailIdentityPolicy, "update-email-identity-policy", "", false, "Update Email Identity Policy")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateEmailTemplate, "update-email-template", "", false, "Update Email Template")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateReputationEntityCustomerManagedStatus, "update-reputation-entity-customer-managed-status", "", false, "Update Reputation Entity Customer Managed Status")
	_sesv2Cmd.Flags().BoolVarP(&_sesv2UpdateReputationEntityPolicy, "update-reputation-entity-policy", "", false, "Update Reputation Entity Policy")

}
