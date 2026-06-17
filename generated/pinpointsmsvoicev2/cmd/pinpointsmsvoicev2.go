package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// pinpointsmsvoicev2Cmd represents the pinpointsmsvoicev2 command
var _pinpointsmsvoicev2Cmd = &cobra.Command{
	Use:   "pinpointsmsvoicev2",
	Short: "AWS pinpointsmsvoicev2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := pinpointsmsvoicev2.NewFromConfig(cfg)
		if _pinpointsmsvoicev2AssociateOriginationIdentity {
			pinpointsmsvoicev2_AssociateOriginationIdentity(cfg, client)
			return
		}
		if _pinpointsmsvoicev2AssociateProtectConfiguration {
			pinpointsmsvoicev2_AssociateProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CarrierLookup {
			pinpointsmsvoicev2_CarrierLookup(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateConfigurationSet {
			pinpointsmsvoicev2_CreateConfigurationSet(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateEventDestination {
			pinpointsmsvoicev2_CreateEventDestination(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateOptOutList {
			pinpointsmsvoicev2_CreateOptOutList(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreatePool {
			pinpointsmsvoicev2_CreatePool(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateProtectConfiguration {
			pinpointsmsvoicev2_CreateProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateRegistration {
			pinpointsmsvoicev2_CreateRegistration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateRegistrationAssociation {
			pinpointsmsvoicev2_CreateRegistrationAssociation(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateRegistrationAttachment {
			pinpointsmsvoicev2_CreateRegistrationAttachment(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateRegistrationVersion {
			pinpointsmsvoicev2_CreateRegistrationVersion(cfg, client)
			return
		}
		if _pinpointsmsvoicev2CreateVerifiedDestinationNumber {
			pinpointsmsvoicev2_CreateVerifiedDestinationNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteAccountDefaultProtectConfiguration {
			pinpointsmsvoicev2_DeleteAccountDefaultProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteConfigurationSet {
			pinpointsmsvoicev2_DeleteConfigurationSet(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteDefaultMessageType {
			pinpointsmsvoicev2_DeleteDefaultMessageType(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteDefaultSenderId {
			pinpointsmsvoicev2_DeleteDefaultSenderId(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteEventDestination {
			pinpointsmsvoicev2_DeleteEventDestination(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteKeyword {
			pinpointsmsvoicev2_DeleteKeyword(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteMediaMessageSpendLimitOverride {
			pinpointsmsvoicev2_DeleteMediaMessageSpendLimitOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteOptOutList {
			pinpointsmsvoicev2_DeleteOptOutList(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteOptedOutNumber {
			pinpointsmsvoicev2_DeleteOptedOutNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeletePool {
			pinpointsmsvoicev2_DeletePool(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteProtectConfiguration {
			pinpointsmsvoicev2_DeleteProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteProtectConfigurationRuleSetNumberOverride {
			pinpointsmsvoicev2_DeleteProtectConfigurationRuleSetNumberOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteRegistration {
			pinpointsmsvoicev2_DeleteRegistration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteRegistrationAttachment {
			pinpointsmsvoicev2_DeleteRegistrationAttachment(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteRegistrationFieldValue {
			pinpointsmsvoicev2_DeleteRegistrationFieldValue(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteResourcePolicy {
			pinpointsmsvoicev2_DeleteResourcePolicy(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteTextMessageSpendLimitOverride {
			pinpointsmsvoicev2_DeleteTextMessageSpendLimitOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteVerifiedDestinationNumber {
			pinpointsmsvoicev2_DeleteVerifiedDestinationNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DeleteVoiceMessageSpendLimitOverride {
			pinpointsmsvoicev2_DeleteVoiceMessageSpendLimitOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeAccountAttributes {
			pinpointsmsvoicev2_DescribeAccountAttributes(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeAccountLimits {
			pinpointsmsvoicev2_DescribeAccountLimits(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeConfigurationSets {
			pinpointsmsvoicev2_DescribeConfigurationSets(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeKeywords {
			pinpointsmsvoicev2_DescribeKeywords(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeOptOutLists {
			pinpointsmsvoicev2_DescribeOptOutLists(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeOptedOutNumbers {
			pinpointsmsvoicev2_DescribeOptedOutNumbers(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribePhoneNumbers {
			pinpointsmsvoicev2_DescribePhoneNumbers(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribePools {
			pinpointsmsvoicev2_DescribePools(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeProtectConfigurations {
			pinpointsmsvoicev2_DescribeProtectConfigurations(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrationAttachments {
			pinpointsmsvoicev2_DescribeRegistrationAttachments(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrationFieldDefinitions {
			pinpointsmsvoicev2_DescribeRegistrationFieldDefinitions(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrationFieldValues {
			pinpointsmsvoicev2_DescribeRegistrationFieldValues(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrationSectionDefinitions {
			pinpointsmsvoicev2_DescribeRegistrationSectionDefinitions(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrationTypeDefinitions {
			pinpointsmsvoicev2_DescribeRegistrationTypeDefinitions(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrationVersions {
			pinpointsmsvoicev2_DescribeRegistrationVersions(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeRegistrations {
			pinpointsmsvoicev2_DescribeRegistrations(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeSenderIds {
			pinpointsmsvoicev2_DescribeSenderIds(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeSpendLimits {
			pinpointsmsvoicev2_DescribeSpendLimits(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DescribeVerifiedDestinationNumbers {
			pinpointsmsvoicev2_DescribeVerifiedDestinationNumbers(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DisassociateOriginationIdentity {
			pinpointsmsvoicev2_DisassociateOriginationIdentity(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DisassociateProtectConfiguration {
			pinpointsmsvoicev2_DisassociateProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2DiscardRegistrationVersion {
			pinpointsmsvoicev2_DiscardRegistrationVersion(cfg, client)
			return
		}
		if _pinpointsmsvoicev2GetProtectConfigurationCountryRuleSet {
			pinpointsmsvoicev2_GetProtectConfigurationCountryRuleSet(cfg, client)
			return
		}
		if _pinpointsmsvoicev2GetResourcePolicy {
			pinpointsmsvoicev2_GetResourcePolicy(cfg, client)
			return
		}
		if _pinpointsmsvoicev2ListPoolOriginationIdentities {
			pinpointsmsvoicev2_ListPoolOriginationIdentities(cfg, client)
			return
		}
		if _pinpointsmsvoicev2ListProtectConfigurationRuleSetNumberOverrides {
			pinpointsmsvoicev2_ListProtectConfigurationRuleSetNumberOverrides(cfg, client)
			return
		}
		if _pinpointsmsvoicev2ListRegistrationAssociations {
			pinpointsmsvoicev2_ListRegistrationAssociations(cfg, client)
			return
		}
		if _pinpointsmsvoicev2ListTagsForResource {
			pinpointsmsvoicev2_ListTagsForResource(cfg, client)
			return
		}
		if _pinpointsmsvoicev2PutKeyword {
			pinpointsmsvoicev2_PutKeyword(cfg, client)
			return
		}
		if _pinpointsmsvoicev2PutMessageFeedback {
			pinpointsmsvoicev2_PutMessageFeedback(cfg, client)
			return
		}
		if _pinpointsmsvoicev2PutOptedOutNumber {
			pinpointsmsvoicev2_PutOptedOutNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2PutProtectConfigurationRuleSetNumberOverride {
			pinpointsmsvoicev2_PutProtectConfigurationRuleSetNumberOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2PutRegistrationFieldValue {
			pinpointsmsvoicev2_PutRegistrationFieldValue(cfg, client)
			return
		}
		if _pinpointsmsvoicev2PutResourcePolicy {
			pinpointsmsvoicev2_PutResourcePolicy(cfg, client)
			return
		}
		if _pinpointsmsvoicev2ReleasePhoneNumber {
			pinpointsmsvoicev2_ReleasePhoneNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2ReleaseSenderId {
			pinpointsmsvoicev2_ReleaseSenderId(cfg, client)
			return
		}
		if _pinpointsmsvoicev2RequestPhoneNumber {
			pinpointsmsvoicev2_RequestPhoneNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2RequestSenderId {
			pinpointsmsvoicev2_RequestSenderId(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SendDestinationNumberVerificationCode {
			pinpointsmsvoicev2_SendDestinationNumberVerificationCode(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SendMediaMessage {
			pinpointsmsvoicev2_SendMediaMessage(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SendTextMessage {
			pinpointsmsvoicev2_SendTextMessage(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SendVoiceMessage {
			pinpointsmsvoicev2_SendVoiceMessage(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetAccountDefaultProtectConfiguration {
			pinpointsmsvoicev2_SetAccountDefaultProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetDefaultMessageFeedbackEnabled {
			pinpointsmsvoicev2_SetDefaultMessageFeedbackEnabled(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetDefaultMessageType {
			pinpointsmsvoicev2_SetDefaultMessageType(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetDefaultSenderId {
			pinpointsmsvoicev2_SetDefaultSenderId(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetMediaMessageSpendLimitOverride {
			pinpointsmsvoicev2_SetMediaMessageSpendLimitOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetTextMessageSpendLimitOverride {
			pinpointsmsvoicev2_SetTextMessageSpendLimitOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SetVoiceMessageSpendLimitOverride {
			pinpointsmsvoicev2_SetVoiceMessageSpendLimitOverride(cfg, client)
			return
		}
		if _pinpointsmsvoicev2SubmitRegistrationVersion {
			pinpointsmsvoicev2_SubmitRegistrationVersion(cfg, client)
			return
		}
		if _pinpointsmsvoicev2TagResource {
			pinpointsmsvoicev2_TagResource(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UntagResource {
			pinpointsmsvoicev2_UntagResource(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UpdateEventDestination {
			pinpointsmsvoicev2_UpdateEventDestination(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UpdatePhoneNumber {
			pinpointsmsvoicev2_UpdatePhoneNumber(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UpdatePool {
			pinpointsmsvoicev2_UpdatePool(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UpdateProtectConfiguration {
			pinpointsmsvoicev2_UpdateProtectConfiguration(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UpdateProtectConfigurationCountryRuleSet {
			pinpointsmsvoicev2_UpdateProtectConfigurationCountryRuleSet(cfg, client)
			return
		}
		if _pinpointsmsvoicev2UpdateSenderId {
			pinpointsmsvoicev2_UpdateSenderId(cfg, client)
			return
		}
		if _pinpointsmsvoicev2VerifyDestinationNumber {
			pinpointsmsvoicev2_VerifyDestinationNumber(cfg, client)
			return
		}

	},
}

var (
	_pinpointsmsvoicev2AssociateOriginationIdentity                    bool
	_pinpointsmsvoicev2AssociateProtectConfiguration                   bool
	_pinpointsmsvoicev2CarrierLookup                                   bool
	_pinpointsmsvoicev2CreateConfigurationSet                          bool
	_pinpointsmsvoicev2CreateEventDestination                          bool
	_pinpointsmsvoicev2CreateOptOutList                                bool
	_pinpointsmsvoicev2CreatePool                                      bool
	_pinpointsmsvoicev2CreateProtectConfiguration                      bool
	_pinpointsmsvoicev2CreateRegistration                              bool
	_pinpointsmsvoicev2CreateRegistrationAssociation                   bool
	_pinpointsmsvoicev2CreateRegistrationAttachment                    bool
	_pinpointsmsvoicev2CreateRegistrationVersion                       bool
	_pinpointsmsvoicev2CreateVerifiedDestinationNumber                 bool
	_pinpointsmsvoicev2DeleteAccountDefaultProtectConfiguration        bool
	_pinpointsmsvoicev2DeleteConfigurationSet                          bool
	_pinpointsmsvoicev2DeleteDefaultMessageType                        bool
	_pinpointsmsvoicev2DeleteDefaultSenderId                           bool
	_pinpointsmsvoicev2DeleteEventDestination                          bool
	_pinpointsmsvoicev2DeleteKeyword                                   bool
	_pinpointsmsvoicev2DeleteMediaMessageSpendLimitOverride            bool
	_pinpointsmsvoicev2DeleteOptOutList                                bool
	_pinpointsmsvoicev2DeleteOptedOutNumber                            bool
	_pinpointsmsvoicev2DeletePool                                      bool
	_pinpointsmsvoicev2DeleteProtectConfiguration                      bool
	_pinpointsmsvoicev2DeleteProtectConfigurationRuleSetNumberOverride bool
	_pinpointsmsvoicev2DeleteRegistration                              bool
	_pinpointsmsvoicev2DeleteRegistrationAttachment                    bool
	_pinpointsmsvoicev2DeleteRegistrationFieldValue                    bool
	_pinpointsmsvoicev2DeleteResourcePolicy                            bool
	_pinpointsmsvoicev2DeleteTextMessageSpendLimitOverride             bool
	_pinpointsmsvoicev2DeleteVerifiedDestinationNumber                 bool
	_pinpointsmsvoicev2DeleteVoiceMessageSpendLimitOverride            bool
	_pinpointsmsvoicev2DescribeAccountAttributes                       bool
	_pinpointsmsvoicev2DescribeAccountLimits                           bool
	_pinpointsmsvoicev2DescribeConfigurationSets                       bool
	_pinpointsmsvoicev2DescribeKeywords                                bool
	_pinpointsmsvoicev2DescribeOptOutLists                             bool
	_pinpointsmsvoicev2DescribeOptedOutNumbers                         bool
	_pinpointsmsvoicev2DescribePhoneNumbers                            bool
	_pinpointsmsvoicev2DescribePools                                   bool
	_pinpointsmsvoicev2DescribeProtectConfigurations                   bool
	_pinpointsmsvoicev2DescribeRegistrationAttachments                 bool
	_pinpointsmsvoicev2DescribeRegistrationFieldDefinitions            bool
	_pinpointsmsvoicev2DescribeRegistrationFieldValues                 bool
	_pinpointsmsvoicev2DescribeRegistrationSectionDefinitions          bool
	_pinpointsmsvoicev2DescribeRegistrationTypeDefinitions             bool
	_pinpointsmsvoicev2DescribeRegistrationVersions                    bool
	_pinpointsmsvoicev2DescribeRegistrations                           bool
	_pinpointsmsvoicev2DescribeSenderIds                               bool
	_pinpointsmsvoicev2DescribeSpendLimits                             bool
	_pinpointsmsvoicev2DescribeVerifiedDestinationNumbers              bool
	_pinpointsmsvoicev2DisassociateOriginationIdentity                 bool
	_pinpointsmsvoicev2DisassociateProtectConfiguration                bool
	_pinpointsmsvoicev2DiscardRegistrationVersion                      bool
	_pinpointsmsvoicev2GetProtectConfigurationCountryRuleSet           bool
	_pinpointsmsvoicev2GetResourcePolicy                               bool
	_pinpointsmsvoicev2ListPoolOriginationIdentities                   bool
	_pinpointsmsvoicev2ListProtectConfigurationRuleSetNumberOverrides  bool
	_pinpointsmsvoicev2ListRegistrationAssociations                    bool
	_pinpointsmsvoicev2ListTagsForResource                             bool
	_pinpointsmsvoicev2PutKeyword                                      bool
	_pinpointsmsvoicev2PutMessageFeedback                              bool
	_pinpointsmsvoicev2PutOptedOutNumber                               bool
	_pinpointsmsvoicev2PutProtectConfigurationRuleSetNumberOverride    bool
	_pinpointsmsvoicev2PutRegistrationFieldValue                       bool
	_pinpointsmsvoicev2PutResourcePolicy                               bool
	_pinpointsmsvoicev2ReleasePhoneNumber                              bool
	_pinpointsmsvoicev2ReleaseSenderId                                 bool
	_pinpointsmsvoicev2RequestPhoneNumber                              bool
	_pinpointsmsvoicev2RequestSenderId                                 bool
	_pinpointsmsvoicev2SendDestinationNumberVerificationCode           bool
	_pinpointsmsvoicev2SendMediaMessage                                bool
	_pinpointsmsvoicev2SendTextMessage                                 bool
	_pinpointsmsvoicev2SendVoiceMessage                                bool
	_pinpointsmsvoicev2SetAccountDefaultProtectConfiguration           bool
	_pinpointsmsvoicev2SetDefaultMessageFeedbackEnabled                bool
	_pinpointsmsvoicev2SetDefaultMessageType                           bool
	_pinpointsmsvoicev2SetDefaultSenderId                              bool
	_pinpointsmsvoicev2SetMediaMessageSpendLimitOverride               bool
	_pinpointsmsvoicev2SetTextMessageSpendLimitOverride                bool
	_pinpointsmsvoicev2SetVoiceMessageSpendLimitOverride               bool
	_pinpointsmsvoicev2SubmitRegistrationVersion                       bool
	_pinpointsmsvoicev2TagResource                                     bool
	_pinpointsmsvoicev2UntagResource                                   bool
	_pinpointsmsvoicev2UpdateEventDestination                          bool
	_pinpointsmsvoicev2UpdatePhoneNumber                               bool
	_pinpointsmsvoicev2UpdatePool                                      bool
	_pinpointsmsvoicev2UpdateProtectConfiguration                      bool
	_pinpointsmsvoicev2UpdateProtectConfigurationCountryRuleSet        bool
	_pinpointsmsvoicev2UpdateSenderId                                  bool
	_pinpointsmsvoicev2VerifyDestinationNumber                         bool

	_pinpointsmsvoicev2Action                       string
	_pinpointsmsvoicev2AttachmentBody               string
	_pinpointsmsvoicev2AttachmentUrl                string
	_pinpointsmsvoicev2AwsReview                    string
	_pinpointsmsvoicev2ClientToken                  string
	_pinpointsmsvoicev2CloudWatchLogsDestination    string
	_pinpointsmsvoicev2ConfigurationSetName         string
	_pinpointsmsvoicev2ConfigurationSetNames        []string
	_pinpointsmsvoicev2Context                      string
	_pinpointsmsvoicev2CountryRuleSetUpdates        string
	_pinpointsmsvoicev2DeletionProtectionEnabled    string
	_pinpointsmsvoicev2DestinationCountryParameters string
	_pinpointsmsvoicev2DestinationPhoneNumber       string
	_pinpointsmsvoicev2DestinationPhoneNumbers      []string
	_pinpointsmsvoicev2DryRun                       string
	_pinpointsmsvoicev2Enabled                      string
	_pinpointsmsvoicev2EventDestinationName         string
	_pinpointsmsvoicev2ExpirationTimestamp          string
	_pinpointsmsvoicev2FieldPath                    string
	_pinpointsmsvoicev2FieldPaths                   []string
	_pinpointsmsvoicev2Filters                      string
	_pinpointsmsvoicev2InternationalSendingEnabled  string
	_pinpointsmsvoicev2IsoCountryCode               string
	_pinpointsmsvoicev2Keyword                      string
	_pinpointsmsvoicev2KeywordAction                string
	_pinpointsmsvoicev2KeywordMessage               string
	_pinpointsmsvoicev2Keywords                     []string
	_pinpointsmsvoicev2KinesisFirehoseDestination   string
	_pinpointsmsvoicev2LanguageCode                 string
	_pinpointsmsvoicev2MatchingEventTypes           string
	_pinpointsmsvoicev2MaxPrice                     string
	_pinpointsmsvoicev2MaxPricePerMinute            string
	_pinpointsmsvoicev2MaxResults                   string
	_pinpointsmsvoicev2MediaUrls                    []string
	_pinpointsmsvoicev2MessageBody                  string
	_pinpointsmsvoicev2MessageBodyTextType          string
	_pinpointsmsvoicev2MessageFeedbackEnabled       string
	_pinpointsmsvoicev2MessageFeedbackStatus        string
	_pinpointsmsvoicev2MessageId                    string
	_pinpointsmsvoicev2MessageType                  string
	_pinpointsmsvoicev2MessageTypes                 string
	_pinpointsmsvoicev2MonthlyLimit                 string
	_pinpointsmsvoicev2NextToken                    string
	_pinpointsmsvoicev2NumberCapabilities           string
	_pinpointsmsvoicev2NumberCapability             string
	_pinpointsmsvoicev2NumberType                   string
	_pinpointsmsvoicev2OptOutListName               string
	_pinpointsmsvoicev2OptOutListNames              []string
	_pinpointsmsvoicev2OptedOutNumber               string
	_pinpointsmsvoicev2OptedOutNumbers              []string
	_pinpointsmsvoicev2OriginationIdentity          string
	_pinpointsmsvoicev2Owner                        string
	_pinpointsmsvoicev2PhoneNumber                  string
	_pinpointsmsvoicev2PhoneNumberId                string
	_pinpointsmsvoicev2PhoneNumberIds               []string
	_pinpointsmsvoicev2Policy                       string
	_pinpointsmsvoicev2PoolId                       string
	_pinpointsmsvoicev2PoolIds                      []string
	_pinpointsmsvoicev2ProtectConfigurationId       string
	_pinpointsmsvoicev2ProtectConfigurationIds      []string
	_pinpointsmsvoicev2RegistrationAttachmentId     string
	_pinpointsmsvoicev2RegistrationAttachmentIds    []string
	_pinpointsmsvoicev2RegistrationId               string
	_pinpointsmsvoicev2RegistrationIds              []string
	_pinpointsmsvoicev2RegistrationType             string
	_pinpointsmsvoicev2RegistrationTypes            []string
	_pinpointsmsvoicev2ResourceArn                  string
	_pinpointsmsvoicev2ResourceId                   string
	_pinpointsmsvoicev2SectionPath                  string
	_pinpointsmsvoicev2SectionPaths                 []string
	_pinpointsmsvoicev2SelectChoices                []string
	_pinpointsmsvoicev2SelfManagedOptOutsEnabled    string
	_pinpointsmsvoicev2SenderId                     string
	_pinpointsmsvoicev2SenderIds                    string
	_pinpointsmsvoicev2SharedRoutesEnabled          string
	_pinpointsmsvoicev2SnsDestination               string
	_pinpointsmsvoicev2TagKeys                      []string
	_pinpointsmsvoicev2Tags                         string
	_pinpointsmsvoicev2TextValue                    string
	_pinpointsmsvoicev2TimeToLive                   string
	_pinpointsmsvoicev2TwoWayChannelArn             string
	_pinpointsmsvoicev2TwoWayChannelRole            string
	_pinpointsmsvoicev2TwoWayEnabled                string
	_pinpointsmsvoicev2VerificationChannel          string
	_pinpointsmsvoicev2VerificationCode             string
	_pinpointsmsvoicev2VerifiedDestinationNumberId  string
	_pinpointsmsvoicev2VerifiedDestinationNumberIds []string
	_pinpointsmsvoicev2VersionNumber                string
	_pinpointsmsvoicev2VersionNumbers               string
	_pinpointsmsvoicev2VoiceId                      string
)

// Associates the specified origination identity with a pool.
// If the origination identity is a phone number and is already associated with
// another pool, an error is returned. A sender ID can be associated with multiple
// pools.
//
// If the origination identity configuration doesn't match the pool's
// configuration, an error is returned.
func pinpointsmsvoicev2_AssociateOriginationIdentity(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.AssociateOriginationIdentityInput{
		// IsoCountryCode: *string, // Required
		// OriginationIdentity: *string, // Required
		// PoolId: *string, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2PoolId) > 0 {
		input.PoolId = aws.String(_pinpointsmsvoicev2PoolId)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}

	if resp, err := client.AssociateOriginationIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate a protect configuration with a configuration set. This replaces the
// configuration sets current protect configuration. A configuration set can only
// be associated with one protect configuration at a time. A protect configuration
// can be associated with multiple configuration sets.
func pinpointsmsvoicev2_AssociateProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.AssociateProtectConfigurationInput{
		// ConfigurationSetName: *string, // Required
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.AssociateProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a destination phone number, including whether the
// number type and whether it is valid, the carrier, and more.
func pinpointsmsvoicev2_CarrierLookup(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CarrierLookupInput{
		// PhoneNumber: *string, // Required
	}

	if len(_pinpointsmsvoicev2PhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_pinpointsmsvoicev2PhoneNumber)
	}

	if resp, err := client.CarrierLookup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new configuration set. After you create the configuration set, you
// can add one or more event destinations to it.
//
// A configuration set is a set of rules that you apply to the SMS and voice
// messages that you send.
//
// When you send a message, you can optionally specify a single configuration set.
func pinpointsmsvoicev2_CreateConfigurationSet(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a new event destination in a configuration set.
// An event destination is a location where you send message events. The event
// options are Amazon CloudWatch, Amazon Data Firehose, or Amazon SNS. For example,
// when a message is delivered successfully, you can send information about that
// event to an event destination, or send notifications to endpoints that are
// subscribed to an Amazon SNS topic.
//
// You can only create one event destination at a time. You must provide a value
// for a single event destination using either CloudWatchLogsDestination ,
// KinesisFirehoseDestination or SnsDestination . If an event destination isn't
// provided then an exception is returned.
//
// Each configuration set can contain between 0 and 5 event destinations. Each
// event destination can contain a reference to a single destination, such as a
// CloudWatch or Firehose destination.
func pinpointsmsvoicev2_CreateEventDestination(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
		// MatchingEventTypes: []types.EventType, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2EventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointsmsvoicev2EventDestinationName)
	}
	if len(_pinpointsmsvoicev2MatchingEventTypes) > 0 {
		if err := assignInputField(input, "MatchingEventTypes", _pinpointsmsvoicev2MatchingEventTypes); err != nil {
			log.Errorf("invalid --matching-event-types: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2CloudWatchLogsDestination) > 0 {
		if err := assignInputField(input, "CloudWatchLogsDestination", _pinpointsmsvoicev2CloudWatchLogsDestination); err != nil {
			log.Errorf("invalid --cloud-watch-logs-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2KinesisFirehoseDestination) > 0 {
		if err := assignInputField(input, "KinesisFirehoseDestination", _pinpointsmsvoicev2KinesisFirehoseDestination); err != nil {
			log.Errorf("invalid --kinesis-firehose-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2SnsDestination) > 0 {
		if err := assignInputField(input, "SnsDestination", _pinpointsmsvoicev2SnsDestination); err != nil {
			log.Errorf("invalid --sns-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new opt-out list.
// If the opt-out list name already exists, an error is returned.
//
// An opt-out list is a list of phone numbers that are opted out, meaning you
// can't send SMS or voice messages to them. If end user replies with the keyword
// "STOP," an entry for the phone number is added to the opt-out list. In addition
// to STOP, your recipients can use any supported opt-out keyword, such as CANCEL
// or OPTOUT. For a list of supported opt-out keywords, see [SMS opt out]in the End User
// Messaging SMS User Guide.
//
// [SMS opt out]: https://docs.aws.amazon.com/pinpoint/latest/userguide/channels-sms-manage.html#channels-sms-manage-optout
func pinpointsmsvoicev2_CreateOptOutList(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateOptOutListInput{
		// OptOutListName: *string, // Required
	}

	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOptOutList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new pool and associates the specified origination identity to the
// pool. A pool can include one or more phone numbers and SenderIds that are
// associated with your Amazon Web Services account.
//
// The new pool inherits its configuration from the specified origination
// identity. This includes keywords, message type, opt-out list, two-way
// configuration, and self-managed opt-out configuration. Deletion protection isn't
// inherited from the origination identity and defaults to false.
//
// If the origination identity is a phone number and is already associated with
// another pool, an error is returned. A sender ID can be associated with multiple
// pools.
func pinpointsmsvoicev2_CreatePool(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreatePoolInput{
		// IsoCountryCode: *string, // Required
		// MessageType: types.MessageType, // Required
		// OriginationIdentity: *string, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2MessageType) > 0 {
		if err := assignInputField(input, "MessageType", _pinpointsmsvoicev2MessageType); err != nil {
			log.Errorf("invalid --message-type: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new protect configuration. By default all country rule sets for each
// capability are set to ALLOW . Update the country rule sets using
// UpdateProtectConfigurationCountryRuleSet . A protect configurations name is
// stored as a Tag with the key set to Name and value as the name of the protect
// configuration.
func pinpointsmsvoicev2_CreateProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateProtectConfigurationInput{}

	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new registration based on the RegistrationType field.
func pinpointsmsvoicev2_CreateRegistration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateRegistrationInput{
		// RegistrationType: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationType) > 0 {
		input.RegistrationType = aws.String(_pinpointsmsvoicev2RegistrationType)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate the registration with an origination identity such as a phone number
// or sender ID.
func pinpointsmsvoicev2_CreateRegistrationAssociation(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateRegistrationAssociationInput{
		// RegistrationId: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2ResourceId) > 0 {
		input.ResourceId = aws.String(_pinpointsmsvoicev2ResourceId)
	}

	if resp, err := client.CreateRegistrationAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new registration attachment to use for uploading a file or a URL to a
// file. The maximum file size is 500KB and valid file extensions are PDF, JPEG and
// PNG. For example, many sender ID registrations require a signed “letter of
// authorization” (LOA) to be submitted.
//
// Use either AttachmentUrl or AttachmentBody to upload your attachment. If both
// are specified then an exception is returned.
func pinpointsmsvoicev2_CreateRegistrationAttachment(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateRegistrationAttachmentInput{}

	if len(_pinpointsmsvoicev2AttachmentBody) > 0 {
		if err := assignInputField(input, "AttachmentBody", _pinpointsmsvoicev2AttachmentBody); err != nil {
			log.Errorf("invalid --attachment-body: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2AttachmentUrl) > 0 {
		input.AttachmentUrl = aws.String(_pinpointsmsvoicev2AttachmentUrl)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRegistrationAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a new version of the registration and increase the VersionNumber. The
// previous version of the registration becomes read-only.
func pinpointsmsvoicev2_CreateRegistrationVersion(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateRegistrationVersionInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}

	if resp, err := client.CreateRegistrationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can only send messages to verified destination numbers when your account is
// in the sandbox. You can add up to 10 verified destination numbers.
func pinpointsmsvoicev2_CreateVerifiedDestinationNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.CreateVerifiedDestinationNumberInput{
		// DestinationPhoneNumber: *string, // Required
	}

	if len(_pinpointsmsvoicev2DestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoicev2DestinationPhoneNumber)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVerifiedDestinationNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the current account default protect configuration.
func pinpointsmsvoicev2_DeleteAccountDefaultProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteAccountDefaultProtectConfigurationInput{}

	if resp, err := client.DeleteAccountDefaultProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing configuration set.
// A configuration set is a set of rules that you apply to voice and SMS messages
// that you send. In a configuration set, you can specify a destination for
// specific types of events related to voice and SMS messages.
func pinpointsmsvoicev2_DeleteConfigurationSet(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteConfigurationSetInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}

	if resp, err := client.DeleteConfigurationSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing default message type on a configuration set.
// A message type is a type of messages that you plan to send. If you send
// account-related messages or time-sensitive messages such as one-time passcodes,
// choose Transactional. If you plan to send messages that contain marketing
// material or other promotional content, choose Promotional. This setting applies
// to your entire Amazon Web Services account.
func pinpointsmsvoicev2_DeleteDefaultMessageType(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteDefaultMessageTypeInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}

	if resp, err := client.DeleteDefaultMessageType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing default sender ID on a configuration set.
// A default sender ID is the identity that appears on recipients' devices when
// they receive SMS messages. Support for sender ID capabilities varies by country
// or region.
func pinpointsmsvoicev2_DeleteDefaultSenderId(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteDefaultSenderIdInput{
		// ConfigurationSetName: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}

	if resp, err := client.DeleteDefaultSenderId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing event destination.
// An event destination is a location where you send response information about
// the messages that you send. For example, when a message is delivered
// successfully, you can send information about that event to an Amazon CloudWatch
// destination, or send notifications to endpoints that are subscribed to an Amazon
// SNS topic.
func pinpointsmsvoicev2_DeleteEventDestination(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2EventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointsmsvoicev2EventDestinationName)
	}

	if resp, err := client.DeleteEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing keyword from an origination phone number or pool.
// A keyword is a word that you can search for on a particular phone number or
// pool. It is also a specific word or phrase that an end user can send to your
// number to elicit a response, such as an informational message or a special
// offer. When your number receives a message that begins with a keyword, End User
// Messaging SMS responds with a customizable message.
//
// Keywords "HELP" and "STOP" can't be deleted or modified.
func pinpointsmsvoicev2_DeleteKeyword(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteKeywordInput{
		// Keyword: *string, // Required
		// OriginationIdentity: *string, // Required
	}

	if len(_pinpointsmsvoicev2Keyword) > 0 {
		input.Keyword = aws.String(_pinpointsmsvoicev2Keyword)
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}

	if resp, err := client.DeleteKeyword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an account-level monthly spending limit override for sending multimedia
// messages (MMS). Deleting a spend limit override will set the EnforcedLimit to
// equal the MaxLimit , which is controlled by Amazon Web Services. For more
// information on spend limits (quotas) see [Quotas for Server Migration Service]in the Server Migration Service User
// Guide.
//
// [Quotas for Server Migration Service]: https://docs.aws.amazon.com/sms-voice/latest/userguide/quotas.html
func pinpointsmsvoicev2_DeleteMediaMessageSpendLimitOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteMediaMessageSpendLimitOverrideInput{}

	if resp, err := client.DeleteMediaMessageSpendLimitOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing opt-out list. All opted out phone numbers in the opt-out
// list are deleted.
//
// If the specified opt-out list name doesn't exist or is in-use by an origination
// phone number or pool, an error is returned.
func pinpointsmsvoicev2_DeleteOptOutList(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteOptOutListInput{
		// OptOutListName: *string, // Required
	}

	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}

	if resp, err := client.DeleteOptOutList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing opted out destination phone number from the specified
// opt-out list.
//
// Each destination phone number can only be deleted once every 30 days.
//
// If the specified destination phone number doesn't exist or if the opt-out list
// doesn't exist, an error is returned.
func pinpointsmsvoicev2_DeleteOptedOutNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteOptedOutNumberInput{
		// OptOutListName: *string, // Required
		// OptedOutNumber: *string, // Required
	}

	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2OptedOutNumber) > 0 {
		input.OptedOutNumber = aws.String(_pinpointsmsvoicev2OptedOutNumber)
	}

	if resp, err := client.DeleteOptedOutNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing pool. Deleting a pool disassociates all origination
// identities from that pool.
//
// If the pool status isn't active or if deletion protection is enabled, an error
// is returned.
//
// A pool is a collection of phone numbers and SenderIds. A pool can include one
// or more phone numbers and SenderIds that are associated with your Amazon Web
// Services account.
func pinpointsmsvoicev2_DeletePool(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeletePoolInput{
		// PoolId: *string, // Required
	}

	if len(_pinpointsmsvoicev2PoolId) > 0 {
		input.PoolId = aws.String(_pinpointsmsvoicev2PoolId)
	}

	if resp, err := client.DeletePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete the protect configuration. The protect configuration must
// have deletion protection disabled and must not be associated as the account
// default protect configuration or associated with a configuration set.
func pinpointsmsvoicev2_DeleteProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteProtectConfigurationInput{
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.DeleteProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete the protect configuration rule set number override.
func pinpointsmsvoicev2_DeleteProtectConfigurationRuleSetNumberOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteProtectConfigurationRuleSetNumberOverrideInput{
		// DestinationPhoneNumber: *string, // Required
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2DestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoicev2DestinationPhoneNumber)
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.DeleteProtectConfigurationRuleSetNumberOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete an existing registration from your account.
func pinpointsmsvoicev2_DeleteRegistration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteRegistrationInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}

	if resp, err := client.DeleteRegistration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently delete the specified registration attachment.
func pinpointsmsvoicev2_DeleteRegistrationAttachment(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteRegistrationAttachmentInput{
		// RegistrationAttachmentId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationAttachmentId) > 0 {
		input.RegistrationAttachmentId = aws.String(_pinpointsmsvoicev2RegistrationAttachmentId)
	}

	if resp, err := client.DeleteRegistrationAttachment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the value in a registration form field.
func pinpointsmsvoicev2_DeleteRegistrationFieldValue(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteRegistrationFieldValueInput{
		// FieldPath: *string, // Required
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2FieldPath) > 0 {
		input.FieldPath = aws.String(_pinpointsmsvoicev2FieldPath)
	}
	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}

	if resp, err := client.DeleteRegistrationFieldValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy document attached to the End User Messaging
// SMS resource. A shared resource can be a Pool, Opt-out list, Sender Id, or Phone
// number.
func pinpointsmsvoicev2_DeleteResourcePolicy(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_pinpointsmsvoicev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointsmsvoicev2ResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an account-level monthly spending limit override for sending text
// messages. Deleting a spend limit override will set the EnforcedLimit to equal
// the MaxLimit , which is controlled by Amazon Web Services. For more information
// on spend limits (quotas) see [Quotas]in the End User Messaging SMS User Guide.
//
// [Quotas]: https://docs.aws.amazon.com/sms-voice/latest/userguide/quotas.html
func pinpointsmsvoicev2_DeleteTextMessageSpendLimitOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteTextMessageSpendLimitOverrideInput{}

	if resp, err := client.DeleteTextMessageSpendLimitOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a verified destination phone number.
func pinpointsmsvoicev2_DeleteVerifiedDestinationNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteVerifiedDestinationNumberInput{
		// VerifiedDestinationNumberId: *string, // Required
	}

	if len(_pinpointsmsvoicev2VerifiedDestinationNumberId) > 0 {
		input.VerifiedDestinationNumberId = aws.String(_pinpointsmsvoicev2VerifiedDestinationNumberId)
	}

	if resp, err := client.DeleteVerifiedDestinationNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an account level monthly spend limit override for sending voice
// messages. Deleting a spend limit override sets the EnforcedLimit equal to the
// MaxLimit , which is controlled by Amazon Web Services. For more information on
// spending limits (quotas) see [Quotas]in the End User Messaging SMS User Guide.
//
// [Quotas]: https://docs.aws.amazon.com/sms-voice/latest/userguide/quotas.html
func pinpointsmsvoicev2_DeleteVoiceMessageSpendLimitOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DeleteVoiceMessageSpendLimitOverrideInput{}

	if resp, err := client.DeleteVoiceMessageSpendLimitOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes attributes of your Amazon Web Services account. The supported account
// attributes include account tier, which indicates whether your account is in the
// sandbox or production environment. When you're ready to move your account out of
// the sandbox, create an Amazon Web Services Support case for a service limit
// increase request.
//
// New accounts are placed into an SMS or voice sandbox. The sandbox protects both
// Amazon Web Services end recipients and SMS or voice recipients from fraud and
// abuse.
func pinpointsmsvoicev2_DescribeAccountAttributes(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeAccountAttributesInput{}

	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAccountAttributes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeAccountAttributesOutput
	p := pinpointsmsvoicev2.NewDescribeAccountAttributesPaginator(client, input)
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

// Describes the current End User Messaging SMS SMS Voice V2 resource quotas for
// your account. The description for a quota includes the quota name, current usage
// toward that quota, and the quota's maximum value.
//
// When you establish an Amazon Web Services account, the account has initial
// quotas on the maximum number of configuration sets, opt-out lists, phone
// numbers, and pools that you can create in a given Region. For more information
// see [Quotas]in the End User Messaging SMS User Guide.
//
// [Quotas]: https://docs.aws.amazon.com/sms-voice/latest/userguide/quotas.html
func pinpointsmsvoicev2_DescribeAccountLimits(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeAccountLimitsInput{}

	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAccountLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeAccountLimitsOutput
	p := pinpointsmsvoicev2.NewDescribeAccountLimitsPaginator(client, input)
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

// Describes the specified configuration sets or all in your account.
// If you specify configuration set names, the output includes information for
// only the specified configuration sets. If you specify filters, the output
// includes information for only those configuration sets that meet the filter
// criteria. If you don't specify configuration set names or filters, the output
// includes information for all configuration sets.
//
// If you specify a configuration set name that isn't valid, an error is returned.
func pinpointsmsvoicev2_DescribeConfigurationSets(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeConfigurationSetsInput{}

	if len(_pinpointsmsvoicev2ConfigurationSetNames) > 0 {
		input.ConfigurationSetNames = append([]string(nil), _pinpointsmsvoicev2ConfigurationSetNames...)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeConfigurationSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeConfigurationSetsOutput
	p := pinpointsmsvoicev2.NewDescribeConfigurationSetsPaginator(client, input)
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

// Describes the specified keywords or all keywords on your origination phone
// number or pool.
//
// A keyword is a word that you can search for on a particular phone number or
// pool. It is also a specific word or phrase that an end user can send to your
// number to elicit a response, such as an informational message or a special
// offer. When your number receives a message that begins with a keyword, End User
// Messaging SMS responds with a customizable message.
//
// If you specify a keyword that isn't valid, an error is returned.
func pinpointsmsvoicev2_DescribeKeywords(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeKeywordsInput{
		// OriginationIdentity: *string, // Required
	}

	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2Keywords) > 0 {
		input.Keywords = append([]string(nil), _pinpointsmsvoicev2Keywords...)
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeKeywords(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeKeywordsOutput
	p := pinpointsmsvoicev2.NewDescribeKeywordsPaginator(client, input)
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

// Describes the specified opt-out list or all opt-out lists in your account.
// If you specify opt-out list names, the output includes information for only the
// specified opt-out lists. Opt-out lists include only those that meet the filter
// criteria. If you don't specify opt-out list names or filters, the output
// includes information for all opt-out lists.
//
// If you specify an opt-out list name that isn't valid, an error is returned.
func pinpointsmsvoicev2_DescribeOptOutLists(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeOptOutListsInput{}

	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2OptOutListNames) > 0 {
		input.OptOutListNames = append([]string(nil), _pinpointsmsvoicev2OptOutListNames...)
	}
	if len(_pinpointsmsvoicev2Owner) > 0 {
		if err := assignInputField(input, "Owner", _pinpointsmsvoicev2Owner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeOptOutLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeOptOutListsOutput
	p := pinpointsmsvoicev2.NewDescribeOptOutListsPaginator(client, input)
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

// Describes the specified opted out destination numbers or all opted out
// destination numbers in an opt-out list.
//
// If you specify opted out numbers, the output includes information for only the
// specified opted out numbers. If you specify filters, the output includes
// information for only those opted out numbers that meet the filter criteria. If
// you don't specify opted out numbers or filters, the output includes information
// for all opted out destination numbers in your opt-out list.
//
// If you specify an opted out number that isn't valid, an exception is returned.
func pinpointsmsvoicev2_DescribeOptedOutNumbers(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeOptedOutNumbersInput{
		// OptOutListName: *string, // Required
	}

	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2OptedOutNumbers) > 0 {
		input.OptedOutNumbers = append([]string(nil), _pinpointsmsvoicev2OptedOutNumbers...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOptedOutNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeOptedOutNumbersOutput
	p := pinpointsmsvoicev2.NewDescribeOptedOutNumbersPaginator(client, input)
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

// Describes the specified origination phone number, or all the phone numbers in
// your account.
//
// If you specify phone number IDs, the output includes information for only the
// specified phone numbers. If you specify filters, the output includes information
// for only those phone numbers that meet the filter criteria. If you don't specify
// phone number IDs or filters, the output includes information for all phone
// numbers.
//
// If you specify a phone number ID that isn't valid, an error is returned.
func pinpointsmsvoicev2_DescribePhoneNumbers(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribePhoneNumbersInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2Owner) > 0 {
		if err := assignInputField(input, "Owner", _pinpointsmsvoicev2Owner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2PhoneNumberIds) > 0 {
		input.PhoneNumberIds = append([]string(nil), _pinpointsmsvoicev2PhoneNumberIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribePhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribePhoneNumbersOutput
	p := pinpointsmsvoicev2.NewDescribePhoneNumbersPaginator(client, input)
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

// Retrieves the specified pools or all pools associated with your Amazon Web
// Services account.
//
// If you specify pool IDs, the output includes information for only the specified
// pools. If you specify filters, the output includes information for only those
// pools that meet the filter criteria. If you don't specify pool IDs or filters,
// the output includes information for all pools.
//
// If you specify a pool ID that isn't valid, an error is returned.
//
// A pool is a collection of phone numbers and SenderIds. A pool can include one
// or more phone numbers and SenderIds that are associated with your Amazon Web
// Services account.
func pinpointsmsvoicev2_DescribePools(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribePoolsInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2Owner) > 0 {
		if err := assignInputField(input, "Owner", _pinpointsmsvoicev2Owner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2PoolIds) > 0 {
		input.PoolIds = append([]string(nil), _pinpointsmsvoicev2PoolIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribePools(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribePoolsOutput
	p := pinpointsmsvoicev2.NewDescribePoolsPaginator(client, input)
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

// Retrieves the protect configurations that match any of filters. If a filter
// isn’t provided then all protect configurations are returned.
func pinpointsmsvoicev2_DescribeProtectConfigurations(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeProtectConfigurationsInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationIds) > 0 {
		input.ProtectConfigurationIds = append([]string(nil), _pinpointsmsvoicev2ProtectConfigurationIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeProtectConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeProtectConfigurationsOutput
	p := pinpointsmsvoicev2.NewDescribeProtectConfigurationsPaginator(client, input)
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

// Retrieves the specified registration attachments or all registration
// attachments associated with your Amazon Web Services account.
func pinpointsmsvoicev2_DescribeRegistrationAttachments(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationAttachmentsInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2RegistrationAttachmentIds) > 0 {
		input.RegistrationAttachmentIds = append([]string(nil), _pinpointsmsvoicev2RegistrationAttachmentIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrationAttachments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationAttachmentsOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationAttachmentsPaginator(client, input)
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

// Retrieves the specified registration type field definitions. You can use
// DescribeRegistrationFieldDefinitions to view the requirements for creating,
// filling out, and submitting each registration type.
func pinpointsmsvoicev2_DescribeRegistrationFieldDefinitions(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationFieldDefinitionsInput{
		// RegistrationType: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationType) > 0 {
		input.RegistrationType = aws.String(_pinpointsmsvoicev2RegistrationType)
	}
	if len(_pinpointsmsvoicev2FieldPaths) > 0 {
		input.FieldPaths = append([]string(nil), _pinpointsmsvoicev2FieldPaths...)
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2SectionPath) > 0 {
		input.SectionPath = aws.String(_pinpointsmsvoicev2SectionPath)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrationFieldDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationFieldDefinitionsOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationFieldDefinitionsPaginator(client, input)
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

// Retrieves the specified registration field values.
func pinpointsmsvoicev2_DescribeRegistrationFieldValues(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationFieldValuesInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2FieldPaths) > 0 {
		input.FieldPaths = append([]string(nil), _pinpointsmsvoicev2FieldPaths...)
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2SectionPath) > 0 {
		input.SectionPath = aws.String(_pinpointsmsvoicev2SectionPath)
	}
	if len(_pinpointsmsvoicev2VersionNumber) > 0 {
		if err := assignInputField(input, "VersionNumber", _pinpointsmsvoicev2VersionNumber); err != nil {
			log.Errorf("invalid --version-number: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrationFieldValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationFieldValuesOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationFieldValuesPaginator(client, input)
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

// Retrieves the specified registration section definitions. You can use
// DescribeRegistrationSectionDefinitions to view the requirements for creating,
// filling out, and submitting each registration type.
func pinpointsmsvoicev2_DescribeRegistrationSectionDefinitions(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationSectionDefinitionsInput{
		// RegistrationType: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationType) > 0 {
		input.RegistrationType = aws.String(_pinpointsmsvoicev2RegistrationType)
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2SectionPaths) > 0 {
		input.SectionPaths = append([]string(nil), _pinpointsmsvoicev2SectionPaths...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrationSectionDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationSectionDefinitionsOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationSectionDefinitionsPaginator(client, input)
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

// Retrieves the specified registration type definitions. You can use
// DescribeRegistrationTypeDefinitions to view the requirements for creating,
// filling out, and submitting each registration type.
func pinpointsmsvoicev2_DescribeRegistrationTypeDefinitions(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationTypeDefinitionsInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2RegistrationTypes) > 0 {
		input.RegistrationTypes = append([]string(nil), _pinpointsmsvoicev2RegistrationTypes...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrationTypeDefinitions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationTypeDefinitionsOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationTypeDefinitionsPaginator(client, input)
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

// Retrieves the specified registration version.
func pinpointsmsvoicev2_DescribeRegistrationVersions(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationVersionsInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2VersionNumbers) > 0 {
		if err := assignInputField(input, "VersionNumbers", _pinpointsmsvoicev2VersionNumbers); err != nil {
			log.Errorf("invalid --version-numbers: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrationVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationVersionsOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationVersionsPaginator(client, input)
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

// Retrieves the specified registrations.
func pinpointsmsvoicev2_DescribeRegistrations(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeRegistrationsInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2RegistrationIds) > 0 {
		input.RegistrationIds = append([]string(nil), _pinpointsmsvoicev2RegistrationIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeRegistrations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeRegistrationsOutput
	p := pinpointsmsvoicev2.NewDescribeRegistrationsPaginator(client, input)
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

// Describes the specified SenderIds or all SenderIds associated with your Amazon
// Web Services account.
//
// If you specify SenderIds, the output includes information for only the
// specified SenderIds. If you specify filters, the output includes information for
// only those SenderIds that meet the filter criteria. If you don't specify
// SenderIds or filters, the output includes information for all SenderIds.
//
// f you specify a sender ID that isn't valid, an error is returned.
func pinpointsmsvoicev2_DescribeSenderIds(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeSenderIdsInput{}

	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2Owner) > 0 {
		if err := assignInputField(input, "Owner", _pinpointsmsvoicev2Owner); err != nil {
			log.Errorf("invalid --owner: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2SenderIds) > 0 {
		if err := assignInputField(input, "SenderIds", _pinpointsmsvoicev2SenderIds); err != nil {
			log.Errorf("invalid --sender-ids: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeSenderIds(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeSenderIdsOutput
	p := pinpointsmsvoicev2.NewDescribeSenderIdsPaginator(client, input)
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

// Describes the current monthly spend limits for sending voice and text messages.
// When you establish an Amazon Web Services account, the account has initial
// monthly spend limit in a given Region. For more information on increasing your
// monthly spend limit, see [Requesting increases to your monthly SMS, MMS, or Voice spending quota]in the End User Messaging SMS User Guide.
//
// [Requesting increases to your monthly SMS, MMS, or Voice spending quota]: https://docs.aws.amazon.com/sms-voice/latest/userguide/awssupport-spend-threshold.html
func pinpointsmsvoicev2_DescribeSpendLimits(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeSpendLimitsInput{}

	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeSpendLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeSpendLimitsOutput
	p := pinpointsmsvoicev2.NewDescribeSpendLimitsPaginator(client, input)
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

// Retrieves the specified verified destination numbers.
func pinpointsmsvoicev2_DescribeVerifiedDestinationNumbers(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DescribeVerifiedDestinationNumbersInput{}

	if len(_pinpointsmsvoicev2DestinationPhoneNumbers) > 0 {
		input.DestinationPhoneNumbers = append([]string(nil), _pinpointsmsvoicev2DestinationPhoneNumbers...)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}
	if len(_pinpointsmsvoicev2VerifiedDestinationNumberIds) > 0 {
		input.VerifiedDestinationNumberIds = append([]string(nil), _pinpointsmsvoicev2VerifiedDestinationNumberIds...)
	}

	if disablePaginator() {
		if resp, err := client.DescribeVerifiedDestinationNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.DescribeVerifiedDestinationNumbersOutput
	p := pinpointsmsvoicev2.NewDescribeVerifiedDestinationNumbersPaginator(client, input)
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

// Removes the specified origination identity from an existing pool.
// If the origination identity isn't associated with the specified pool, an error
// is returned.
func pinpointsmsvoicev2_DisassociateOriginationIdentity(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DisassociateOriginationIdentityInput{
		// IsoCountryCode: *string, // Required
		// OriginationIdentity: *string, // Required
		// PoolId: *string, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2PoolId) > 0 {
		input.PoolId = aws.String(_pinpointsmsvoicev2PoolId)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}

	if resp, err := client.DisassociateOriginationIdentity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociate a protect configuration from a configuration set.
func pinpointsmsvoicev2_DisassociateProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DisassociateProtectConfigurationInput{
		// ConfigurationSetName: *string, // Required
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.DisassociateProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Discard the current version of the registration.
func pinpointsmsvoicev2_DiscardRegistrationVersion(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.DiscardRegistrationVersionInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}

	if resp, err := client.DiscardRegistrationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the CountryRuleSet for the specified NumberCapability from a protect
// configuration.
func pinpointsmsvoicev2_GetProtectConfigurationCountryRuleSet(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.GetProtectConfigurationCountryRuleSetInput{
		// NumberCapability: types.NumberCapability, // Required
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2NumberCapability) > 0 {
		if err := assignInputField(input, "NumberCapability", _pinpointsmsvoicev2NumberCapability); err != nil {
			log.Errorf("invalid --number-capability: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.GetProtectConfigurationCountryRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the JSON text of the resource-based policy document attached to the
// End User Messaging SMS resource. A shared resource can be a Pool, Opt-out list,
// Sender Id, or Phone number.
func pinpointsmsvoicev2_GetResourcePolicy(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_pinpointsmsvoicev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointsmsvoicev2ResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all associated origination identities in your pool.
// If you specify filters, the output includes information for only those
// origination identities that meet the filter criteria.
func pinpointsmsvoicev2_ListPoolOriginationIdentities(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.ListPoolOriginationIdentitiesInput{
		// PoolId: *string, // Required
	}

	if len(_pinpointsmsvoicev2PoolId) > 0 {
		input.PoolId = aws.String(_pinpointsmsvoicev2PoolId)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPoolOriginationIdentities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.ListPoolOriginationIdentitiesOutput
	p := pinpointsmsvoicev2.NewListPoolOriginationIdentitiesPaginator(client, input)
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

// Retrieve all of the protect configuration rule set number overrides that match
// the filters.
func pinpointsmsvoicev2_ListProtectConfigurationRuleSetNumberOverrides(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.ListProtectConfigurationRuleSetNumberOverridesInput{
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProtectConfigurationRuleSetNumberOverrides(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.ListProtectConfigurationRuleSetNumberOverridesOutput
	p := pinpointsmsvoicev2.NewListProtectConfigurationRuleSetNumberOverridesPaginator(client, input)
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

// Retrieve all of the origination identities that are associated with a
// registration.
func pinpointsmsvoicev2_ListRegistrationAssociations(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.ListRegistrationAssociationsInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2Filters) > 0 {
		if err := assignInputField(input, "Filters", _pinpointsmsvoicev2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _pinpointsmsvoicev2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NextToken) > 0 {
		input.NextToken = aws.String(_pinpointsmsvoicev2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRegistrationAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*pinpointsmsvoicev2.ListRegistrationAssociationsOutput
	p := pinpointsmsvoicev2.NewListRegistrationAssociationsPaginator(client, input)
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

// List all tags associated with a resource.
func pinpointsmsvoicev2_ListTagsForResource(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_pinpointsmsvoicev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointsmsvoicev2ResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a keyword configuration on an origination phone number or
// pool.
//
// A keyword is a word that you can search for on a particular phone number or
// pool. It is also a specific word or phrase that an end user can send to your
// number to elicit a response, such as an informational message or a special
// offer. When your number receives a message that begins with a keyword, End User
// Messaging SMS responds with a customizable message.
//
// If you specify a keyword that isn't valid, an error is returned.
func pinpointsmsvoicev2_PutKeyword(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.PutKeywordInput{
		// Keyword: *string, // Required
		// KeywordMessage: *string, // Required
		// OriginationIdentity: *string, // Required
	}

	if len(_pinpointsmsvoicev2Keyword) > 0 {
		input.Keyword = aws.String(_pinpointsmsvoicev2Keyword)
	}
	if len(_pinpointsmsvoicev2KeywordMessage) > 0 {
		input.KeywordMessage = aws.String(_pinpointsmsvoicev2KeywordMessage)
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2KeywordAction) > 0 {
		if err := assignInputField(input, "KeywordAction", _pinpointsmsvoicev2KeywordAction); err != nil {
			log.Errorf("invalid --keyword-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutKeyword(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set the MessageFeedbackStatus as RECEIVED or FAILED for the passed in
// MessageId.
//
// If you use message feedback then you must update message feedback record. When
// you receive a signal that a user has received the message you must use
// PutMessageFeedback to set the message feedback record as RECEIVED ; Otherwise,
// an hour after the message feedback record is set to FAILED .
func pinpointsmsvoicev2_PutMessageFeedback(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.PutMessageFeedbackInput{
		// MessageFeedbackStatus: types.MessageFeedbackStatus, // Required
		// MessageId: *string, // Required
	}

	if len(_pinpointsmsvoicev2MessageFeedbackStatus) > 0 {
		if err := assignInputField(input, "MessageFeedbackStatus", _pinpointsmsvoicev2MessageFeedbackStatus); err != nil {
			log.Errorf("invalid --message-feedback-status: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MessageId) > 0 {
		input.MessageId = aws.String(_pinpointsmsvoicev2MessageId)
	}

	if resp, err := client.PutMessageFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an opted out destination phone number in the opt-out list.
// If the destination phone number isn't valid or if the specified opt-out list
// doesn't exist, an error is returned.
func pinpointsmsvoicev2_PutOptedOutNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.PutOptedOutNumberInput{
		// OptOutListName: *string, // Required
		// OptedOutNumber: *string, // Required
	}

	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2OptedOutNumber) > 0 {
		input.OptedOutNumber = aws.String(_pinpointsmsvoicev2OptedOutNumber)
	}

	if resp, err := client.PutOptedOutNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create or update a phone number rule override and associate it with a protect
// configuration.
func pinpointsmsvoicev2_PutProtectConfigurationRuleSetNumberOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.PutProtectConfigurationRuleSetNumberOverrideInput{
		// Action: types.ProtectConfigurationRuleOverrideAction, // Required
		// DestinationPhoneNumber: *string, // Required
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2Action) > 0 {
		if err := assignInputField(input, "Action", _pinpointsmsvoicev2Action); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2DestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoicev2DestinationPhoneNumber)
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2ExpirationTimestamp) > 0 {
		if err := assignInputField(input, "ExpirationTimestamp", _pinpointsmsvoicev2ExpirationTimestamp); err != nil {
			log.Errorf("invalid --expiration-timestamp: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutProtectConfigurationRuleSetNumberOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a field value for a registration.
func pinpointsmsvoicev2_PutRegistrationFieldValue(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.PutRegistrationFieldValueInput{
		// FieldPath: *string, // Required
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2FieldPath) > 0 {
		input.FieldPath = aws.String(_pinpointsmsvoicev2FieldPath)
	}
	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2RegistrationAttachmentId) > 0 {
		input.RegistrationAttachmentId = aws.String(_pinpointsmsvoicev2RegistrationAttachmentId)
	}
	if len(_pinpointsmsvoicev2SelectChoices) > 0 {
		input.SelectChoices = append([]string(nil), _pinpointsmsvoicev2SelectChoices...)
	}
	if len(_pinpointsmsvoicev2TextValue) > 0 {
		input.TextValue = aws.String(_pinpointsmsvoicev2TextValue)
	}

	if resp, err := client.PutRegistrationFieldValue(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based policy to a End User Messaging SMS resource(phone
// number, sender Id, phone poll, or opt-out list) that is used for sharing the
// resource. A shared resource can be a Pool, Opt-out list, Sender Id, or Phone
// number. For more information about resource-based policies, see [Working with shared resources]in the End User
// Messaging SMS User Guide.
//
// [Working with shared resources]: https://docs.aws.amazon.com/sms-voice/latest/userguide/shared-resources.html
func pinpointsmsvoicev2_PutResourcePolicy(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_pinpointsmsvoicev2Policy) > 0 {
		input.Policy = aws.String(_pinpointsmsvoicev2Policy)
	}
	if len(_pinpointsmsvoicev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointsmsvoicev2ResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Releases an existing origination phone number in your account. Once released, a
// phone number is no longer available for sending messages.
//
// If the origination phone number has deletion protection enabled or is
// associated with a pool, an error is returned.
func pinpointsmsvoicev2_ReleasePhoneNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.ReleasePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_pinpointsmsvoicev2PhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_pinpointsmsvoicev2PhoneNumberId)
	}

	if resp, err := client.ReleasePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Releases an existing sender ID in your account.
func pinpointsmsvoicev2_ReleaseSenderId(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.ReleaseSenderIdInput{
		// IsoCountryCode: *string, // Required
		// SenderId: *string, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2SenderId) > 0 {
		input.SenderId = aws.String(_pinpointsmsvoicev2SenderId)
	}

	if resp, err := client.ReleaseSenderId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request an origination phone number for use in your account. For more
// information on phone number request see [Request a phone number]in the End User Messaging SMS User
// Guide.
//
// [Request a phone number]: https://docs.aws.amazon.com/sms-voice/latest/userguide/phone-numbers-request.html
func pinpointsmsvoicev2_RequestPhoneNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.RequestPhoneNumberInput{
		// IsoCountryCode: *string, // Required
		// MessageType: types.MessageType, // Required
		// NumberCapabilities: []types.NumberCapability, // Required
		// NumberType: types.RequestableNumberType, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2MessageType) > 0 {
		if err := assignInputField(input, "MessageType", _pinpointsmsvoicev2MessageType); err != nil {
			log.Errorf("invalid --message-type: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NumberCapabilities) > 0 {
		if err := assignInputField(input, "NumberCapabilities", _pinpointsmsvoicev2NumberCapabilities); err != nil {
			log.Errorf("invalid --number-capabilities: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NumberType) > 0 {
		if err := assignInputField(input, "NumberType", _pinpointsmsvoicev2NumberType); err != nil {
			log.Errorf("invalid --number-type: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2InternationalSendingEnabled) > 0 {
		if err := assignInputField(input, "InternationalSendingEnabled", _pinpointsmsvoicev2InternationalSendingEnabled); err != nil {
			log.Errorf("invalid --international-sending-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2PoolId) > 0 {
		input.PoolId = aws.String(_pinpointsmsvoicev2PoolId)
	}
	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RequestPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Request a new sender ID that doesn't require registration.
func pinpointsmsvoicev2_RequestSenderId(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.RequestSenderIdInput{
		// IsoCountryCode: *string, // Required
		// SenderId: *string, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2SenderId) > 0 {
		input.SenderId = aws.String(_pinpointsmsvoicev2SenderId)
	}
	if len(_pinpointsmsvoicev2ClientToken) > 0 {
		input.ClientToken = aws.String(_pinpointsmsvoicev2ClientToken)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MessageTypes) > 0 {
		if err := assignInputField(input, "MessageTypes", _pinpointsmsvoicev2MessageTypes); err != nil {
			log.Errorf("invalid --message-types: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.RequestSenderId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Before you can send test messages to a verified destination phone number you
// need to opt-in the verified destination phone number. Creates a new text message
// with a verification code and send it to a verified destination phone number.
// Once you have the verification code use VerifyDestinationNumberto opt-in the verified destination
// phone number to receive messages.
func pinpointsmsvoicev2_SendDestinationNumberVerificationCode(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SendDestinationNumberVerificationCodeInput{
		// VerificationChannel: types.VerificationChannel, // Required
		// VerifiedDestinationNumberId: *string, // Required
	}

	if len(_pinpointsmsvoicev2VerificationChannel) > 0 {
		if err := assignInputField(input, "VerificationChannel", _pinpointsmsvoicev2VerificationChannel); err != nil {
			log.Errorf("invalid --verification-channel: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2VerifiedDestinationNumberId) > 0 {
		input.VerifiedDestinationNumberId = aws.String(_pinpointsmsvoicev2VerifiedDestinationNumberId)
	}
	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2Context) > 0 {
		if err := assignInputField(input, "Context", _pinpointsmsvoicev2Context); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2DestinationCountryParameters) > 0 {
		if err := assignInputField(input, "DestinationCountryParameters", _pinpointsmsvoicev2DestinationCountryParameters); err != nil {
			log.Errorf("invalid --destination-country-parameters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2LanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _pinpointsmsvoicev2LanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}

	if resp, err := client.SendDestinationNumberVerificationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new multimedia message (MMS) and sends it to a recipient's phone
// number.
func pinpointsmsvoicev2_SendMediaMessage(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SendMediaMessageInput{
		// DestinationPhoneNumber: *string, // Required
		// OriginationIdentity: *string, // Required
	}

	if len(_pinpointsmsvoicev2DestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoicev2DestinationPhoneNumber)
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2Context) > 0 {
		if err := assignInputField(input, "Context", _pinpointsmsvoicev2Context); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2DryRun) > 0 {
		if err := assignInputField(input, "DryRun", _pinpointsmsvoicev2DryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxPrice) > 0 {
		input.MaxPrice = aws.String(_pinpointsmsvoicev2MaxPrice)
	}
	if len(_pinpointsmsvoicev2MediaUrls) > 0 {
		input.MediaUrls = append([]string(nil), _pinpointsmsvoicev2MediaUrls...)
	}
	if len(_pinpointsmsvoicev2MessageBody) > 0 {
		input.MessageBody = aws.String(_pinpointsmsvoicev2MessageBody)
	}
	if len(_pinpointsmsvoicev2MessageFeedbackEnabled) > 0 {
		if err := assignInputField(input, "MessageFeedbackEnabled", _pinpointsmsvoicev2MessageFeedbackEnabled); err != nil {
			log.Errorf("invalid --message-feedback-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}
	if len(_pinpointsmsvoicev2TimeToLive) > 0 {
		if err := assignInputField(input, "TimeToLive", _pinpointsmsvoicev2TimeToLive); err != nil {
			log.Errorf("invalid --time-to-live: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendMediaMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new text message and sends it to a recipient's phone number.
// SendTextMessage only sends an SMS message to one recipient each time it is
// invoked.
//
// SMS throughput limits are measured in Message Parts per Second (MPS). Your MPS
// limit depends on the destination country of your messages, as well as the type
// of phone number (origination number) that you use to send the message. For more
// information about MPS, see [Message Parts per Second (MPS) limits]in the End User Messaging SMS User Guide.
//
// [Message Parts per Second (MPS) limits]: https://docs.aws.amazon.com/sms-voice/latest/userguide/sms-limitations-mps.html
func pinpointsmsvoicev2_SendTextMessage(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SendTextMessageInput{
		// DestinationPhoneNumber: *string, // Required
	}

	if len(_pinpointsmsvoicev2DestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoicev2DestinationPhoneNumber)
	}
	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2Context) > 0 {
		if err := assignInputField(input, "Context", _pinpointsmsvoicev2Context); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2DestinationCountryParameters) > 0 {
		if err := assignInputField(input, "DestinationCountryParameters", _pinpointsmsvoicev2DestinationCountryParameters); err != nil {
			log.Errorf("invalid --destination-country-parameters: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2DryRun) > 0 {
		if err := assignInputField(input, "DryRun", _pinpointsmsvoicev2DryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2Keyword) > 0 {
		input.Keyword = aws.String(_pinpointsmsvoicev2Keyword)
	}
	if len(_pinpointsmsvoicev2MaxPrice) > 0 {
		input.MaxPrice = aws.String(_pinpointsmsvoicev2MaxPrice)
	}
	if len(_pinpointsmsvoicev2MessageBody) > 0 {
		input.MessageBody = aws.String(_pinpointsmsvoicev2MessageBody)
	}
	if len(_pinpointsmsvoicev2MessageFeedbackEnabled) > 0 {
		if err := assignInputField(input, "MessageFeedbackEnabled", _pinpointsmsvoicev2MessageFeedbackEnabled); err != nil {
			log.Errorf("invalid --message-feedback-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MessageType) > 0 {
		if err := assignInputField(input, "MessageType", _pinpointsmsvoicev2MessageType); err != nil {
			log.Errorf("invalid --message-type: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}
	if len(_pinpointsmsvoicev2TimeToLive) > 0 {
		if err := assignInputField(input, "TimeToLive", _pinpointsmsvoicev2TimeToLive); err != nil {
			log.Errorf("invalid --time-to-live: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendTextMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to send a request that sends a voice message. This operation uses [Amazon Polly]
// to convert a text script into a voice message.
//
// [Amazon Polly]: http://aws.amazon.com/polly/
func pinpointsmsvoicev2_SendVoiceMessage(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SendVoiceMessageInput{
		// DestinationPhoneNumber: *string, // Required
		// OriginationIdentity: *string, // Required
	}

	if len(_pinpointsmsvoicev2DestinationPhoneNumber) > 0 {
		input.DestinationPhoneNumber = aws.String(_pinpointsmsvoicev2DestinationPhoneNumber)
	}
	if len(_pinpointsmsvoicev2OriginationIdentity) > 0 {
		input.OriginationIdentity = aws.String(_pinpointsmsvoicev2OriginationIdentity)
	}
	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2Context) > 0 {
		if err := assignInputField(input, "Context", _pinpointsmsvoicev2Context); err != nil {
			log.Errorf("invalid --context: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2DryRun) > 0 {
		if err := assignInputField(input, "DryRun", _pinpointsmsvoicev2DryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MaxPricePerMinute) > 0 {
		input.MaxPricePerMinute = aws.String(_pinpointsmsvoicev2MaxPricePerMinute)
	}
	if len(_pinpointsmsvoicev2MessageBody) > 0 {
		input.MessageBody = aws.String(_pinpointsmsvoicev2MessageBody)
	}
	if len(_pinpointsmsvoicev2MessageBodyTextType) > 0 {
		if err := assignInputField(input, "MessageBodyTextType", _pinpointsmsvoicev2MessageBodyTextType); err != nil {
			log.Errorf("invalid --message-body-text-type: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MessageFeedbackEnabled) > 0 {
		if err := assignInputField(input, "MessageFeedbackEnabled", _pinpointsmsvoicev2MessageFeedbackEnabled); err != nil {
			log.Errorf("invalid --message-feedback-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}
	if len(_pinpointsmsvoicev2TimeToLive) > 0 {
		if err := assignInputField(input, "TimeToLive", _pinpointsmsvoicev2TimeToLive); err != nil {
			log.Errorf("invalid --time-to-live: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2VoiceId) > 0 {
		if err := assignInputField(input, "VoiceId", _pinpointsmsvoicev2VoiceId); err != nil {
			log.Errorf("invalid --voice-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.SendVoiceMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Set a protect configuration as your account default. You can only have one
// account default protect configuration at a time. The current account default
// protect configuration is replaced with the provided protect configuration.
func pinpointsmsvoicev2_SetAccountDefaultProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetAccountDefaultProtectConfigurationInput{
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.SetAccountDefaultProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets a configuration set's default for message feedback.
func pinpointsmsvoicev2_SetDefaultMessageFeedbackEnabled(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetDefaultMessageFeedbackEnabledInput{
		// ConfigurationSetName: *string, // Required
		// MessageFeedbackEnabled: *bool, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2MessageFeedbackEnabled) > 0 {
		if err := assignInputField(input, "MessageFeedbackEnabled", _pinpointsmsvoicev2MessageFeedbackEnabled); err != nil {
			log.Errorf("invalid --message-feedback-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetDefaultMessageFeedbackEnabled(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the default message type on a configuration set.
// Choose the category of SMS messages that you plan to send from this account. If
// you send account-related messages or time-sensitive messages such as one-time
// passcodes, choose Transactional. If you plan to send messages that contain
// marketing material or other promotional content, choose Promotional. This
// setting applies to your entire Amazon Web Services account.
func pinpointsmsvoicev2_SetDefaultMessageType(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetDefaultMessageTypeInput{
		// ConfigurationSetName: *string, // Required
		// MessageType: types.MessageType, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2MessageType) > 0 {
		if err := assignInputField(input, "MessageType", _pinpointsmsvoicev2MessageType); err != nil {
			log.Errorf("invalid --message-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetDefaultMessageType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets default sender ID on a configuration set.
// When sending a text message to a destination country that supports sender IDs,
// the default sender ID on the configuration set specified will be used if no
// dedicated origination phone numbers or registered sender IDs are available in
// your account.
func pinpointsmsvoicev2_SetDefaultSenderId(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetDefaultSenderIdInput{
		// ConfigurationSetName: *string, // Required
		// SenderId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2SenderId) > 0 {
		input.SenderId = aws.String(_pinpointsmsvoicev2SenderId)
	}

	if resp, err := client.SetDefaultSenderId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets an account level monthly spend limit override for sending MMS messages.
// The requested spend limit must be less than or equal to the MaxLimit , which is
// set by Amazon Web Services.
func pinpointsmsvoicev2_SetMediaMessageSpendLimitOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetMediaMessageSpendLimitOverrideInput{
		// MonthlyLimit: *int64, // Required
	}

	if len(_pinpointsmsvoicev2MonthlyLimit) > 0 {
		if err := assignInputField(input, "MonthlyLimit", _pinpointsmsvoicev2MonthlyLimit); err != nil {
			log.Errorf("invalid --monthly-limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetMediaMessageSpendLimitOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets an account level monthly spend limit override for sending text messages.
// The requested spend limit must be less than or equal to the MaxLimit , which is
// set by Amazon Web Services.
func pinpointsmsvoicev2_SetTextMessageSpendLimitOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetTextMessageSpendLimitOverrideInput{
		// MonthlyLimit: *int64, // Required
	}

	if len(_pinpointsmsvoicev2MonthlyLimit) > 0 {
		if err := assignInputField(input, "MonthlyLimit", _pinpointsmsvoicev2MonthlyLimit); err != nil {
			log.Errorf("invalid --monthly-limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetTextMessageSpendLimitOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets an account level monthly spend limit override for sending voice messages.
// The requested spend limit must be less than or equal to the MaxLimit , which is
// set by Amazon Web Services.
func pinpointsmsvoicev2_SetVoiceMessageSpendLimitOverride(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SetVoiceMessageSpendLimitOverrideInput{
		// MonthlyLimit: *int64, // Required
	}

	if len(_pinpointsmsvoicev2MonthlyLimit) > 0 {
		if err := assignInputField(input, "MonthlyLimit", _pinpointsmsvoicev2MonthlyLimit); err != nil {
			log.Errorf("invalid --monthly-limit: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetVoiceMessageSpendLimitOverride(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Submit the specified registration for review and approval.
func pinpointsmsvoicev2_SubmitRegistrationVersion(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.SubmitRegistrationVersionInput{
		// RegistrationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2RegistrationId) > 0 {
		input.RegistrationId = aws.String(_pinpointsmsvoicev2RegistrationId)
	}
	if len(_pinpointsmsvoicev2AwsReview) > 0 {
		if err := assignInputField(input, "AwsReview", _pinpointsmsvoicev2AwsReview); err != nil {
			log.Errorf("invalid --aws-review: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitRegistrationVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites only the specified tags for the specified resource. When you
// specify an existing tag key, the value is overwritten with the new value. Each
// tag consists of a key and an optional value. Tag keys must be unique per
// resource. For more information about tags, see [Tags]in the End User Messaging SMS
// User Guide.
//
// [Tags]: https://docs.aws.amazon.com/sms-voice/latest/userguide/phone-numbers-tags.html
func pinpointsmsvoicev2_TagResource(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_pinpointsmsvoicev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointsmsvoicev2ResourceArn)
	}
	if len(_pinpointsmsvoicev2Tags) > 0 {
		if err := assignInputField(input, "Tags", _pinpointsmsvoicev2Tags); err != nil {
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

// Removes the association of the specified tags from a resource. For more
// information on tags see [Tags]in the End User Messaging SMS User Guide.
//
// [Tags]: https://docs.aws.amazon.com/sms-voice/latest/userguide/phone-numbers-tags.html
func pinpointsmsvoicev2_UntagResource(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_pinpointsmsvoicev2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_pinpointsmsvoicev2ResourceArn)
	}
	if len(_pinpointsmsvoicev2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _pinpointsmsvoicev2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing event destination in a configuration set. You can update
// the IAM role ARN for CloudWatch Logs and Firehose. You can also enable or
// disable the event destination.
//
// You may want to update an event destination to change its matching event types
// or updating the destination resource ARN. You can't change an event
// destination's type between CloudWatch Logs, Firehose, and Amazon SNS.
func pinpointsmsvoicev2_UpdateEventDestination(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UpdateEventDestinationInput{
		// ConfigurationSetName: *string, // Required
		// EventDestinationName: *string, // Required
	}

	if len(_pinpointsmsvoicev2ConfigurationSetName) > 0 {
		input.ConfigurationSetName = aws.String(_pinpointsmsvoicev2ConfigurationSetName)
	}
	if len(_pinpointsmsvoicev2EventDestinationName) > 0 {
		input.EventDestinationName = aws.String(_pinpointsmsvoicev2EventDestinationName)
	}
	if len(_pinpointsmsvoicev2CloudWatchLogsDestination) > 0 {
		if err := assignInputField(input, "CloudWatchLogsDestination", _pinpointsmsvoicev2CloudWatchLogsDestination); err != nil {
			log.Errorf("invalid --cloud-watch-logs-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2Enabled) > 0 {
		if err := assignInputField(input, "Enabled", _pinpointsmsvoicev2Enabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2KinesisFirehoseDestination) > 0 {
		if err := assignInputField(input, "KinesisFirehoseDestination", _pinpointsmsvoicev2KinesisFirehoseDestination); err != nil {
			log.Errorf("invalid --kinesis-firehose-destination: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2MatchingEventTypes) > 0 {
		if err := assignInputField(input, "MatchingEventTypes", _pinpointsmsvoicev2MatchingEventTypes); err != nil {
			log.Errorf("invalid --matching-event-types: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2SnsDestination) > 0 {
		if err := assignInputField(input, "SnsDestination", _pinpointsmsvoicev2SnsDestination); err != nil {
			log.Errorf("invalid --sns-destination: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEventDestination(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing origination phone number. You can
// update the opt-out list, enable or disable two-way messaging, change the
// TwoWayChannelArn, enable or disable self-managed opt-outs, and enable or disable
// deletion protection.
//
// If the origination phone number is associated with a pool, an error is returned.
func pinpointsmsvoicev2_UpdatePhoneNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UpdatePhoneNumberInput{
		// PhoneNumberId: *string, // Required
	}

	if len(_pinpointsmsvoicev2PhoneNumberId) > 0 {
		input.PhoneNumberId = aws.String(_pinpointsmsvoicev2PhoneNumberId)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2InternationalSendingEnabled) > 0 {
		if err := assignInputField(input, "InternationalSendingEnabled", _pinpointsmsvoicev2InternationalSendingEnabled); err != nil {
			log.Errorf("invalid --international-sending-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2SelfManagedOptOutsEnabled) > 0 {
		if err := assignInputField(input, "SelfManagedOptOutsEnabled", _pinpointsmsvoicev2SelfManagedOptOutsEnabled); err != nil {
			log.Errorf("invalid --self-managed-opt-outs-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2TwoWayChannelArn) > 0 {
		input.TwoWayChannelArn = aws.String(_pinpointsmsvoicev2TwoWayChannelArn)
	}
	if len(_pinpointsmsvoicev2TwoWayChannelRole) > 0 {
		input.TwoWayChannelRole = aws.String(_pinpointsmsvoicev2TwoWayChannelRole)
	}
	if len(_pinpointsmsvoicev2TwoWayEnabled) > 0 {
		if err := assignInputField(input, "TwoWayEnabled", _pinpointsmsvoicev2TwoWayEnabled); err != nil {
			log.Errorf("invalid --two-way-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing pool. You can update the opt-out list,
// enable or disable two-way messaging, change the TwoWayChannelArn , enable or
// disable self-managed opt-outs, enable or disable deletion protection, and enable
// or disable shared routes.
func pinpointsmsvoicev2_UpdatePool(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UpdatePoolInput{
		// PoolId: *string, // Required
	}

	if len(_pinpointsmsvoicev2PoolId) > 0 {
		input.PoolId = aws.String(_pinpointsmsvoicev2PoolId)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2OptOutListName) > 0 {
		input.OptOutListName = aws.String(_pinpointsmsvoicev2OptOutListName)
	}
	if len(_pinpointsmsvoicev2SelfManagedOptOutsEnabled) > 0 {
		if err := assignInputField(input, "SelfManagedOptOutsEnabled", _pinpointsmsvoicev2SelfManagedOptOutsEnabled); err != nil {
			log.Errorf("invalid --self-managed-opt-outs-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2SharedRoutesEnabled) > 0 {
		if err := assignInputField(input, "SharedRoutesEnabled", _pinpointsmsvoicev2SharedRoutesEnabled); err != nil {
			log.Errorf("invalid --shared-routes-enabled: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2TwoWayChannelArn) > 0 {
		input.TwoWayChannelArn = aws.String(_pinpointsmsvoicev2TwoWayChannelArn)
	}
	if len(_pinpointsmsvoicev2TwoWayChannelRole) > 0 {
		input.TwoWayChannelRole = aws.String(_pinpointsmsvoicev2TwoWayChannelRole)
	}
	if len(_pinpointsmsvoicev2TwoWayEnabled) > 0 {
		if err := assignInputField(input, "TwoWayEnabled", _pinpointsmsvoicev2TwoWayEnabled); err != nil {
			log.Errorf("invalid --two-way-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePool(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the setting for an existing protect configuration.
func pinpointsmsvoicev2_UpdateProtectConfiguration(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UpdateProtectConfigurationInput{
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProtectConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a country rule set to ALLOW , BLOCK , MONITOR , or FILTER messages to be
// sent to the specified destination counties. You can update one or multiple
// countries at a time. The updates are only applied to the specified
// NumberCapability type.
func pinpointsmsvoicev2_UpdateProtectConfigurationCountryRuleSet(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UpdateProtectConfigurationCountryRuleSetInput{
		// CountryRuleSetUpdates: map[string]types.ProtectConfigurationCountryRuleSetInformation, // Required
		// NumberCapability: types.NumberCapability, // Required
		// ProtectConfigurationId: *string, // Required
	}

	if len(_pinpointsmsvoicev2CountryRuleSetUpdates) > 0 {
		if err := assignInputField(input, "CountryRuleSetUpdates", _pinpointsmsvoicev2CountryRuleSetUpdates); err != nil {
			log.Errorf("invalid --country-rule-set-updates: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2NumberCapability) > 0 {
		if err := assignInputField(input, "NumberCapability", _pinpointsmsvoicev2NumberCapability); err != nil {
			log.Errorf("invalid --number-capability: %s", err.Error())
			return
		}
	}
	if len(_pinpointsmsvoicev2ProtectConfigurationId) > 0 {
		input.ProtectConfigurationId = aws.String(_pinpointsmsvoicev2ProtectConfigurationId)
	}

	if resp, err := client.UpdateProtectConfigurationCountryRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of an existing sender ID.
func pinpointsmsvoicev2_UpdateSenderId(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.UpdateSenderIdInput{
		// IsoCountryCode: *string, // Required
		// SenderId: *string, // Required
	}

	if len(_pinpointsmsvoicev2IsoCountryCode) > 0 {
		input.IsoCountryCode = aws.String(_pinpointsmsvoicev2IsoCountryCode)
	}
	if len(_pinpointsmsvoicev2SenderId) > 0 {
		input.SenderId = aws.String(_pinpointsmsvoicev2SenderId)
	}
	if len(_pinpointsmsvoicev2DeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _pinpointsmsvoicev2DeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSenderId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use the verification code that was received by the verified destination phone
// number to opt-in the verified destination phone number to receive more messages.
func pinpointsmsvoicev2_VerifyDestinationNumber(cfg aws.Config, client *pinpointsmsvoicev2.Client) {
	input := &pinpointsmsvoicev2.VerifyDestinationNumberInput{
		// VerificationCode: *string, // Required
		// VerifiedDestinationNumberId: *string, // Required
	}

	if len(_pinpointsmsvoicev2VerificationCode) > 0 {
		input.VerificationCode = aws.String(_pinpointsmsvoicev2VerificationCode)
	}
	if len(_pinpointsmsvoicev2VerifiedDestinationNumberId) > 0 {
		input.VerifiedDestinationNumberId = aws.String(_pinpointsmsvoicev2VerifiedDestinationNumberId)
	}

	if resp, err := client.VerifyDestinationNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_pinpointsmsvoicev2Cmd)
	_pinpointsmsvoicev2Cmd.Flags().SortFlags = false

	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Action, "action", "", "", "Action")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2AttachmentBody, "attachment-body", "", "", "Attachment Body")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2AttachmentUrl, "attachment-url", "", "", "Attachment URL")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2AwsReview, "aws-review", "", "", "AWS Review")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2ClientToken, "client-token", "", "", "Client Token")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2CloudWatchLogsDestination, "cloud-watch-logs-destination", "", "", "Cloud Watch Logs Destination")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2ConfigurationSetName, "configuration-set-name", "", "", "Configuration Set Name")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2ConfigurationSetNames, "configuration-set-names", "", nil, "Configuration Set Names")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Context, "context", "", "", "Context")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2CountryRuleSetUpdates, "country-rule-set-updates", "", "", "Country Rule Set Updates")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2DeletionProtectionEnabled, "deletion-protection-enabled", "", "", "Deletion Protection Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2DestinationCountryParameters, "destination-country-parameters", "", "", "Destination Country Parameters")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2DestinationPhoneNumber, "destination-phone-number", "", "", "Destination Phone Number")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2DestinationPhoneNumbers, "destination-phone-numbers", "", nil, "Destination Phone Numbers")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2DryRun, "dry-run", "", "", "Dry Run")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Enabled, "enabled", "", "", "Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2EventDestinationName, "event-destination-name", "", "", "Event Destination Name")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2ExpirationTimestamp, "expiration-timestamp", "", "", "Expiration Timestamp")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2FieldPath, "field-path", "", "", "Field Path")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2FieldPaths, "field-paths", "", nil, "Field Paths")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Filters, "filters", "", "", "Filters")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2InternationalSendingEnabled, "international-sending-enabled", "", "", "International Sending Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2IsoCountryCode, "iso-country-code", "", "", "Iso Country Code")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Keyword, "keyword", "", "", "Keyword")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2KeywordAction, "keyword-action", "", "", "Keyword Action")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2KeywordMessage, "keyword-message", "", "", "Keyword Message")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2Keywords, "keywords", "", nil, "Keywords")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2KinesisFirehoseDestination, "kinesis-firehose-destination", "", "", "Kinesis Firehose Destination")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2LanguageCode, "language-code", "", "", "Language Code")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MatchingEventTypes, "matching-event-types", "", "", "Matching Event Types")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MaxPrice, "max-price", "", "", "Max Price")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MaxPricePerMinute, "max-price-per-minute", "", "", "Max Price Per Minute")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MaxResults, "max-results", "", "", "Max Results")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2MediaUrls, "media-urls", "", nil, "Media Urls")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageBody, "message-body", "", "", "Message Body")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageBodyTextType, "message-body-text-type", "", "", "Message Body Text Type")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageFeedbackEnabled, "message-feedback-enabled", "", "", "Message Feedback Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageFeedbackStatus, "message-feedback-status", "", "", "Message Feedback Status")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageId, "message-id", "", "", "Message ID")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageType, "message-type", "", "", "Message Type")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MessageTypes, "message-types", "", "", "Message Types")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2MonthlyLimit, "monthly-limit", "", "", "Monthly Limit")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2NextToken, "next-token", "", "", "Next Token")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2NumberCapabilities, "number-capabilities", "", "", "Number Capabilities")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2NumberCapability, "number-capability", "", "", "Number Capability")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2NumberType, "number-type", "", "", "Number Type")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2OptOutListName, "opt-out-list-name", "", "", "Opt Out List Name")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2OptOutListNames, "opt-out-list-names", "", nil, "Opt Out List Names")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2OptedOutNumber, "opted-out-number", "", "", "Opted Out Number")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2OptedOutNumbers, "opted-out-numbers", "", nil, "Opted Out Numbers")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2OriginationIdentity, "origination-identity", "", "", "Origination Identity")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Owner, "owner", "", "", "Owner")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2PhoneNumber, "phone-number", "", "", "Phone Number")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2PhoneNumberId, "phone-number-id", "", "", "Phone Number ID")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2PhoneNumberIds, "phone-number-ids", "", nil, "Phone Number Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Policy, "policy", "", "", "Policy")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2PoolId, "pool-id", "", "", "Pool ID")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2PoolIds, "pool-ids", "", nil, "Pool Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2ProtectConfigurationId, "protect-configuration-id", "", "", "Protect Configuration ID")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2ProtectConfigurationIds, "protect-configuration-ids", "", nil, "Protect Configuration Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2RegistrationAttachmentId, "registration-attachment-id", "", "", "Registration Attachment ID")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2RegistrationAttachmentIds, "registration-attachment-ids", "", nil, "Registration Attachment Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2RegistrationId, "registration-id", "", "", "Registration ID")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2RegistrationIds, "registration-ids", "", nil, "Registration Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2RegistrationType, "registration-type", "", "", "Registration Type")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2RegistrationTypes, "registration-types", "", nil, "Registration Types")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2ResourceId, "resource-id", "", "", "Resource ID")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2SectionPath, "section-path", "", "", "Section Path")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2SectionPaths, "section-paths", "", nil, "Section Paths")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2SelectChoices, "select-choices", "", nil, "Select Choices")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2SelfManagedOptOutsEnabled, "self-managed-opt-outs-enabled", "", "", "Self Managed Opt Outs Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2SenderId, "sender-id", "", "", "Sender ID")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2SenderIds, "sender-ids", "", "", "Sender Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2SharedRoutesEnabled, "shared-routes-enabled", "", "", "Shared Routes Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2SnsDestination, "sns-destination", "", "", "SNS Destination")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2Tags, "tags", "", "", "Tags")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2TextValue, "text-value", "", "", "Text Value")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2TimeToLive, "time-to-live", "", "", "Time To Live")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2TwoWayChannelArn, "two-way-channel-arn", "", "", "Two Way Channel ARN")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2TwoWayChannelRole, "two-way-channel-role", "", "", "Two Way Channel Role")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2TwoWayEnabled, "two-way-enabled", "", "", "Two Way Enabled")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2VerificationChannel, "verification-channel", "", "", "Verification Channel")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2VerificationCode, "verification-code", "", "", "Verification Code")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2VerifiedDestinationNumberId, "verified-destination-number-id", "", "", "Verified Destination Number ID")
	_pinpointsmsvoicev2Cmd.Flags().StringSliceVarP(&_pinpointsmsvoicev2VerifiedDestinationNumberIds, "verified-destination-number-ids", "", nil, "Verified Destination Number Ids")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2VersionNumber, "version-number", "", "", "Version Number")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2VersionNumbers, "version-numbers", "", "", "Version Numbers")
	_pinpointsmsvoicev2Cmd.Flags().StringVarP(&_pinpointsmsvoicev2VoiceId, "voice-id", "", "", "Voice ID")

	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2AssociateOriginationIdentity, "associate-origination-identity", "", false, "Associate Origination Identity")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2AssociateProtectConfiguration, "associate-protect-configuration", "", false, "Associate Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CarrierLookup, "carrier-lookup", "", false, "Carrier Lookup")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateConfigurationSet, "create-configuration-set", "", false, "Create Configuration Set")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateEventDestination, "create-event-destination", "", false, "Create Event Destination")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateOptOutList, "create-opt-out-list", "", false, "Create Opt Out List")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreatePool, "create-pool", "", false, "Create Pool")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateProtectConfiguration, "create-protect-configuration", "", false, "Create Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateRegistration, "create-registration", "", false, "Create Registration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateRegistrationAssociation, "create-registration-association", "", false, "Create Registration Association")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateRegistrationAttachment, "create-registration-attachment", "", false, "Create Registration Attachment")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateRegistrationVersion, "create-registration-version", "", false, "Create Registration Version")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2CreateVerifiedDestinationNumber, "create-verified-destination-number", "", false, "Create Verified Destination Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteAccountDefaultProtectConfiguration, "delete-account-default-protect-configuration", "", false, "Delete Account Default Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteConfigurationSet, "delete-configuration-set", "", false, "Delete Configuration Set")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteDefaultMessageType, "delete-default-message-type", "", false, "Delete Default Message Type")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteDefaultSenderId, "delete-default-sender-id", "", false, "Delete Default Sender ID")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteEventDestination, "delete-event-destination", "", false, "Delete Event Destination")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteKeyword, "delete-keyword", "", false, "Delete Keyword")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteMediaMessageSpendLimitOverride, "delete-media-message-spend-limit-override", "", false, "Delete Media Message Spend Limit Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteOptOutList, "delete-opt-out-list", "", false, "Delete Opt Out List")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteOptedOutNumber, "delete-opted-out-number", "", false, "Delete Opted Out Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeletePool, "delete-pool", "", false, "Delete Pool")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteProtectConfiguration, "delete-protect-configuration", "", false, "Delete Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteProtectConfigurationRuleSetNumberOverride, "delete-protect-configuration-rule-set-number-override", "", false, "Delete Protect Configuration Rule Set Number Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteRegistration, "delete-registration", "", false, "Delete Registration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteRegistrationAttachment, "delete-registration-attachment", "", false, "Delete Registration Attachment")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteRegistrationFieldValue, "delete-registration-field-value", "", false, "Delete Registration Field Value")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteTextMessageSpendLimitOverride, "delete-text-message-spend-limit-override", "", false, "Delete Text Message Spend Limit Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteVerifiedDestinationNumber, "delete-verified-destination-number", "", false, "Delete Verified Destination Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DeleteVoiceMessageSpendLimitOverride, "delete-voice-message-spend-limit-override", "", false, "Delete Voice Message Spend Limit Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeAccountAttributes, "describe-account-attributes", "", false, "Describe Account Attributes")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeAccountLimits, "describe-account-limits", "", false, "Describe Account Limits")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeConfigurationSets, "describe-configuration-sets", "", false, "Describe Configuration Sets")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeKeywords, "describe-keywords", "", false, "Describe Keywords")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeOptOutLists, "describe-opt-out-lists", "", false, "Describe Opt Out Lists")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeOptedOutNumbers, "describe-opted-out-numbers", "", false, "Describe Opted Out Numbers")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribePhoneNumbers, "describe-phone-numbers", "", false, "Describe Phone Numbers")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribePools, "describe-pools", "", false, "Describe Pools")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeProtectConfigurations, "describe-protect-configurations", "", false, "Describe Protect Configurations")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrationAttachments, "describe-registration-attachments", "", false, "Describe Registration Attachments")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrationFieldDefinitions, "describe-registration-field-definitions", "", false, "Describe Registration Field Definitions")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrationFieldValues, "describe-registration-field-values", "", false, "Describe Registration Field Values")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrationSectionDefinitions, "describe-registration-section-definitions", "", false, "Describe Registration Section Definitions")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrationTypeDefinitions, "describe-registration-type-definitions", "", false, "Describe Registration Type Definitions")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrationVersions, "describe-registration-versions", "", false, "Describe Registration Versions")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeRegistrations, "describe-registrations", "", false, "Describe Registrations")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeSenderIds, "describe-sender-ids", "", false, "Describe Sender Ids")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeSpendLimits, "describe-spend-limits", "", false, "Describe Spend Limits")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DescribeVerifiedDestinationNumbers, "describe-verified-destination-numbers", "", false, "Describe Verified Destination Numbers")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DisassociateOriginationIdentity, "disassociate-origination-identity", "", false, "Disassociate Origination Identity")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DisassociateProtectConfiguration, "disassociate-protect-configuration", "", false, "Disassociate Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2DiscardRegistrationVersion, "discard-registration-version", "", false, "Discard Registration Version")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2GetProtectConfigurationCountryRuleSet, "get-protect-configuration-country-rule-set", "", false, "Get Protect Configuration Country Rule Set")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2GetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2ListPoolOriginationIdentities, "list-pool-origination-identities", "", false, "List Pool Origination Identities")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2ListProtectConfigurationRuleSetNumberOverrides, "list-protect-configuration-rule-set-number-overrides", "", false, "List Protect Configuration Rule Set Number Overrides")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2ListRegistrationAssociations, "list-registration-associations", "", false, "List Registration Associations")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2PutKeyword, "put-keyword", "", false, "Put Keyword")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2PutMessageFeedback, "put-message-feedback", "", false, "Put Message Feedback")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2PutOptedOutNumber, "put-opted-out-number", "", false, "Put Opted Out Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2PutProtectConfigurationRuleSetNumberOverride, "put-protect-configuration-rule-set-number-override", "", false, "Put Protect Configuration Rule Set Number Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2PutRegistrationFieldValue, "put-registration-field-value", "", false, "Put Registration Field Value")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2PutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2ReleasePhoneNumber, "release-phone-number", "", false, "Release Phone Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2ReleaseSenderId, "release-sender-id", "", false, "Release Sender ID")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2RequestPhoneNumber, "request-phone-number", "", false, "Request Phone Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2RequestSenderId, "request-sender-id", "", false, "Request Sender ID")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SendDestinationNumberVerificationCode, "send-destination-number-verification-code", "", false, "Send Destination Number Verification Code")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SendMediaMessage, "send-media-message", "", false, "Send Media Message")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SendTextMessage, "send-text-message", "", false, "Send Text Message")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SendVoiceMessage, "send-voice-message", "", false, "Send Voice Message")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetAccountDefaultProtectConfiguration, "set-account-default-protect-configuration", "", false, "Set Account Default Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetDefaultMessageFeedbackEnabled, "set-default-message-feedback-enabled", "", false, "Set Default Message Feedback Enabled")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetDefaultMessageType, "set-default-message-type", "", false, "Set Default Message Type")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetDefaultSenderId, "set-default-sender-id", "", false, "Set Default Sender ID")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetMediaMessageSpendLimitOverride, "set-media-message-spend-limit-override", "", false, "Set Media Message Spend Limit Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetTextMessageSpendLimitOverride, "set-text-message-spend-limit-override", "", false, "Set Text Message Spend Limit Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SetVoiceMessageSpendLimitOverride, "set-voice-message-spend-limit-override", "", false, "Set Voice Message Spend Limit Override")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2SubmitRegistrationVersion, "submit-registration-version", "", false, "Submit Registration Version")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2TagResource, "tag-resource", "", false, "Tag Resource")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UntagResource, "untag-resource", "", false, "Untag Resource")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UpdateEventDestination, "update-event-destination", "", false, "Update Event Destination")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UpdatePhoneNumber, "update-phone-number", "", false, "Update Phone Number")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UpdatePool, "update-pool", "", false, "Update Pool")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UpdateProtectConfiguration, "update-protect-configuration", "", false, "Update Protect Configuration")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UpdateProtectConfigurationCountryRuleSet, "update-protect-configuration-country-rule-set", "", false, "Update Protect Configuration Country Rule Set")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2UpdateSenderId, "update-sender-id", "", false, "Update Sender ID")
	_pinpointsmsvoicev2Cmd.Flags().BoolVarP(&_pinpointsmsvoicev2VerifyDestinationNumber, "verify-destination-number", "", false, "Verify Destination Number")

}
