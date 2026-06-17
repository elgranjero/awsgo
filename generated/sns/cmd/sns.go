package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// snsCmd represents the sns command
var _snsCmd = &cobra.Command{
	Use:   "sns",
	Short: "AWS sns CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := sns.NewFromConfig(cfg)
		if _snsAddPermission {
			sns_AddPermission(cfg, client)
			return
		}
		if _snsCheckIfPhoneNumberIsOptedOut {
			sns_CheckIfPhoneNumberIsOptedOut(cfg, client)
			return
		}
		if _snsConfirmSubscription {
			sns_ConfirmSubscription(cfg, client)
			return
		}
		if _snsCreatePlatformApplication {
			sns_CreatePlatformApplication(cfg, client)
			return
		}
		if _snsCreatePlatformEndpoint {
			sns_CreatePlatformEndpoint(cfg, client)
			return
		}
		if _snsCreateSMSSandboxPhoneNumber {
			sns_CreateSMSSandboxPhoneNumber(cfg, client)
			return
		}
		if _snsCreateTopic {
			sns_CreateTopic(cfg, client)
			return
		}
		if _snsDeleteEndpoint {
			sns_DeleteEndpoint(cfg, client)
			return
		}
		if _snsDeletePlatformApplication {
			sns_DeletePlatformApplication(cfg, client)
			return
		}
		if _snsDeleteSMSSandboxPhoneNumber {
			sns_DeleteSMSSandboxPhoneNumber(cfg, client)
			return
		}
		if _snsDeleteTopic {
			sns_DeleteTopic(cfg, client)
			return
		}
		if _snsGetDataProtectionPolicy {
			sns_GetDataProtectionPolicy(cfg, client)
			return
		}
		if _snsGetEndpointAttributes {
			sns_GetEndpointAttributes(cfg, client)
			return
		}
		if _snsGetPlatformApplicationAttributes {
			sns_GetPlatformApplicationAttributes(cfg, client)
			return
		}
		if _snsGetSMSAttributes {
			sns_GetSMSAttributes(cfg, client)
			return
		}
		if _snsGetSMSSandboxAccountStatus {
			sns_GetSMSSandboxAccountStatus(cfg, client)
			return
		}
		if _snsGetSubscriptionAttributes {
			sns_GetSubscriptionAttributes(cfg, client)
			return
		}
		if _snsGetTopicAttributes {
			sns_GetTopicAttributes(cfg, client)
			return
		}
		if _snsListEndpointsByPlatformApplication {
			sns_ListEndpointsByPlatformApplication(cfg, client)
			return
		}
		if _snsListOriginationNumbers {
			sns_ListOriginationNumbers(cfg, client)
			return
		}
		if _snsListPhoneNumbersOptedOut {
			sns_ListPhoneNumbersOptedOut(cfg, client)
			return
		}
		if _snsListPlatformApplications {
			sns_ListPlatformApplications(cfg, client)
			return
		}
		if _snsListSMSSandboxPhoneNumbers {
			sns_ListSMSSandboxPhoneNumbers(cfg, client)
			return
		}
		if _snsListSubscriptions {
			sns_ListSubscriptions(cfg, client)
			return
		}
		if _snsListSubscriptionsByTopic {
			sns_ListSubscriptionsByTopic(cfg, client)
			return
		}
		if _snsListTagsForResource {
			sns_ListTagsForResource(cfg, client)
			return
		}
		if _snsListTopics {
			sns_ListTopics(cfg, client)
			return
		}
		if _snsOptInPhoneNumber {
			sns_OptInPhoneNumber(cfg, client)
			return
		}
		if _snsPublish {
			sns_Publish(cfg, client)
			return
		}
		if _snsPublishBatch {
			sns_PublishBatch(cfg, client)
			return
		}
		if _snsPutDataProtectionPolicy {
			sns_PutDataProtectionPolicy(cfg, client)
			return
		}
		if _snsRemovePermission {
			sns_RemovePermission(cfg, client)
			return
		}
		if _snsSetEndpointAttributes {
			sns_SetEndpointAttributes(cfg, client)
			return
		}
		if _snsSetPlatformApplicationAttributes {
			sns_SetPlatformApplicationAttributes(cfg, client)
			return
		}
		if _snsSetSMSAttributes {
			sns_SetSMSAttributes(cfg, client)
			return
		}
		if _snsSetSubscriptionAttributes {
			sns_SetSubscriptionAttributes(cfg, client)
			return
		}
		if _snsSetTopicAttributes {
			sns_SetTopicAttributes(cfg, client)
			return
		}
		if _snsSubscribe {
			sns_Subscribe(cfg, client)
			return
		}
		if _snsTagResource {
			sns_TagResource(cfg, client)
			return
		}
		if _snsUnsubscribe {
			sns_Unsubscribe(cfg, client)
			return
		}
		if _snsUntagResource {
			sns_UntagResource(cfg, client)
			return
		}
		if _snsVerifySMSSandboxPhoneNumber {
			sns_VerifySMSSandboxPhoneNumber(cfg, client)
			return
		}

	},
}

var (
	_snsAddPermission                      bool
	_snsCheckIfPhoneNumberIsOptedOut       bool
	_snsConfirmSubscription                bool
	_snsCreatePlatformApplication          bool
	_snsCreatePlatformEndpoint             bool
	_snsCreateSMSSandboxPhoneNumber        bool
	_snsCreateTopic                        bool
	_snsDeleteEndpoint                     bool
	_snsDeletePlatformApplication          bool
	_snsDeleteSMSSandboxPhoneNumber        bool
	_snsDeleteTopic                        bool
	_snsGetDataProtectionPolicy            bool
	_snsGetEndpointAttributes              bool
	_snsGetPlatformApplicationAttributes   bool
	_snsGetSMSAttributes                   bool
	_snsGetSMSSandboxAccountStatus         bool
	_snsGetSubscriptionAttributes          bool
	_snsGetTopicAttributes                 bool
	_snsListEndpointsByPlatformApplication bool
	_snsListOriginationNumbers             bool
	_snsListPhoneNumbersOptedOut           bool
	_snsListPlatformApplications           bool
	_snsListSMSSandboxPhoneNumbers         bool
	_snsListSubscriptions                  bool
	_snsListSubscriptionsByTopic           bool
	_snsListTagsForResource                bool
	_snsListTopics                         bool
	_snsOptInPhoneNumber                   bool
	_snsPublish                            bool
	_snsPublishBatch                       bool
	_snsPutDataProtectionPolicy            bool
	_snsRemovePermission                   bool
	_snsSetEndpointAttributes              bool
	_snsSetPlatformApplicationAttributes   bool
	_snsSetSMSAttributes                   bool
	_snsSetSubscriptionAttributes          bool
	_snsSetTopicAttributes                 bool
	_snsSubscribe                          bool
	_snsTagResource                        bool
	_snsUnsubscribe                        bool
	_snsUntagResource                      bool
	_snsVerifySMSSandboxPhoneNumber        bool

	_snsActionName                 []string
	_snsAttributeName              string
	_snsAttributeValue             string
	_snsAttributes                 string
	_snsAuthenticateOnUnsubscribe  string
	_snsAWSAccountId               []string
	_snsCustomUserData             string
	_snsDataProtectionPolicy       string
	_snsEndpoint                   string
	_snsEndpointArn                string
	_snsLabel                      string
	_snsLanguageCode               string
	_snsMaxResults                 string
	_snsMessage                    string
	_snsMessageAttributes          string
	_snsMessageDeduplicationId     string
	_snsMessageGroupId             string
	_snsMessageStructure           string
	_snsName                       string
	_snsNextToken                  string
	_snsOneTimePassword            string
	_snsPhoneNumber                string
	_snsPlatform                   string
	_snsPlatformApplicationArn     string
	_snsProtocol                   string
	_snsPublishBatchRequestEntries string
	_snsResourceArn                string
	_snsReturnSubscriptionArn      string
	_snsSubject                    string
	_snsSubscriptionArn            string
	_snsTagKeys                    []string
	_snsTags                       string
	_snsTargetArn                  string
	_snsToken                      string
	_snsTopicArn                   string
)

// Adds a statement to a topic's access control policy, granting access for the
// specified Amazon Web Services accounts to the specified actions.
//
// To remove the ability to change topic permissions, you must deny permissions to
// the AddPermission , RemovePermission , and SetTopicAttributes actions in your
// IAM policy.
func sns_AddPermission(cfg aws.Config, client *sns.Client) {
	input := &sns.AddPermissionInput{
		// AWSAccountId: []string, // Required
		// ActionName: []string, // Required
		// Label: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_snsAWSAccountId) > 0 {
		input.AWSAccountId = append([]string(nil), _snsAWSAccountId...)
	}
	if len(_snsActionName) > 0 {
		input.ActionName = append([]string(nil), _snsActionName...)
	}
	if len(_snsLabel) > 0 {
		input.Label = aws.String(_snsLabel)
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}

	if resp, err := client.AddPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Accepts a phone number and indicates whether the phone holder has opted out of
// receiving SMS messages from your Amazon Web Services account. You cannot send
// SMS messages to a number that is opted out.
//
// To resume sending messages, you can opt in the number by using the
// OptInPhoneNumber action.
func sns_CheckIfPhoneNumberIsOptedOut(cfg aws.Config, client *sns.Client) {
	input := &sns.CheckIfPhoneNumberIsOptedOutInput{
		// PhoneNumber: *string, // Required
	}

	if len(_snsPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_snsPhoneNumber)
	}

	if resp, err := client.CheckIfPhoneNumberIsOptedOut(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies an endpoint owner's intent to receive messages by validating the token
// sent to the endpoint by an earlier Subscribe action. If the token is valid, the
// action creates a new subscription and returns its Amazon Resource Name (ARN).
// This call requires an AWS signature only when the AuthenticateOnUnsubscribe
// flag is set to "true".
func sns_ConfirmSubscription(cfg aws.Config, client *sns.Client) {
	input := &sns.ConfirmSubscriptionInput{
		// Token: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_snsToken) > 0 {
		input.Token = aws.String(_snsToken)
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}
	if len(_snsAuthenticateOnUnsubscribe) > 0 {
		input.AuthenticateOnUnsubscribe = aws.String(_snsAuthenticateOnUnsubscribe)
	}

	if resp, err := client.ConfirmSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a platform application object for one of the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging), to which
// devices and mobile apps may register. You must specify PlatformPrincipal and
// PlatformCredential attributes when using the CreatePlatformApplication action.
//
// PlatformPrincipal and PlatformCredential are received from the notification
// service.
//
// - For ADM, PlatformPrincipal is client id and PlatformCredential is client
// secret .
//
// - For APNS and APNS_SANDBOX using certificate credentials, PlatformPrincipal
// is SSL certificate and PlatformCredential is private key .
//
// - For APNS and APNS_SANDBOX using token credentials, PlatformPrincipal is
// signing key ID and PlatformCredential is signing key .
//
// - For Baidu, PlatformPrincipal is API key and PlatformCredential is secret key
// .
//
// - For GCM (Firebase Cloud Messaging) using key credentials, there is no
// PlatformPrincipal . The PlatformCredential is API key .
//
// - For GCM (Firebase Cloud Messaging) using token credentials, there is no
// PlatformPrincipal . The PlatformCredential is a JSON formatted private key
// file. When using the Amazon Web Services CLI or Amazon Web Services SDKs, the
// file must be in string format and special characters must be ignored. To format
// the file correctly, Amazon SNS recommends using the following command:
// SERVICE_JSON=$(jq (at)json < service.json) .
//
// - For MPNS, PlatformPrincipal is TLS certificate and PlatformCredential is
// private key .
//
// - For WNS, PlatformPrincipal is Package Security Identifier and
// PlatformCredential is secret key .
//
// You can use the returned PlatformApplicationArn as an attribute for the
// CreatePlatformEndpoint action.
func sns_CreatePlatformApplication(cfg aws.Config, client *sns.Client) {
	input := &sns.CreatePlatformApplicationInput{
		// Attributes: map[string]string, // Required
		// Name: *string, // Required
		// Platform: *string, // Required
	}

	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_snsName) > 0 {
		input.Name = aws.String(_snsName)
	}
	if len(_snsPlatform) > 0 {
		input.Platform = aws.String(_snsPlatform)
	}

	if resp, err := client.CreatePlatformApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an endpoint for a device and mobile app on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS.
// CreatePlatformEndpoint requires the PlatformApplicationArn that is returned
// from CreatePlatformApplication . You can use the returned EndpointArn to send a
// message to a mobile app or by the Subscribe action for subscription to a topic.
// The CreatePlatformEndpoint action is idempotent, so if the requester already
// owns an endpoint with the same device token and attributes, that endpoint's ARN
// is returned without creating a new endpoint. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// When using CreatePlatformEndpoint with Baidu, two attributes must be provided:
// ChannelId and UserId. The token field must also contain the ChannelId. For more
// information, see [Creating an Amazon SNS Endpoint for Baidu].
//
// [Creating an Amazon SNS Endpoint for Baidu]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePushBaiduEndpoint.html
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_CreatePlatformEndpoint(cfg aws.Config, client *sns.Client) {
	input := &sns.CreatePlatformEndpointInput{
		// PlatformApplicationArn: *string, // Required
		// Token: *string, // Required
	}

	if len(_snsPlatformApplicationArn) > 0 {
		input.PlatformApplicationArn = aws.String(_snsPlatformApplicationArn)
	}
	if len(_snsToken) > 0 {
		input.Token = aws.String(_snsToken)
	}
	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_snsCustomUserData) > 0 {
		input.CustomUserData = aws.String(_snsCustomUserData)
	}

	if resp, err := client.CreatePlatformEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a destination phone number to an Amazon Web Services account in the SMS
// sandbox and sends a one-time password (OTP) to that phone number.
//
// When you start using Amazon SNS to send SMS messages, your Amazon Web Services
// account is in the SMS sandbox. The SMS sandbox provides a safe environment for
// you to try Amazon SNS features without risking your reputation as an SMS sender.
// While your Amazon Web Services account is in the SMS sandbox, you can use all of
// the features of Amazon SNS. However, you can send SMS messages only to verified
// destination phone numbers. For more information, including how to move out of
// the sandbox to send messages without restrictions, see [SMS sandbox]in the Amazon SNS
// Developer Guide.
//
// [SMS sandbox]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
func sns_CreateSMSSandboxPhoneNumber(cfg aws.Config, client *sns.Client) {
	input := &sns.CreateSMSSandboxPhoneNumberInput{
		// PhoneNumber: *string, // Required
	}

	if len(_snsPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_snsPhoneNumber)
	}
	if len(_snsLanguageCode) > 0 {
		if err := assignInputField(input, "LanguageCode", _snsLanguageCode); err != nil {
			log.Errorf("invalid --language-code: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSMSSandboxPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a topic to which notifications can be published. Users can create at
// most 100,000 standard topics (at most 1,000 FIFO topics). For more information,
// see [Creating an Amazon SNS topic]in the Amazon SNS Developer Guide. This action is idempotent, so if the
// requester already owns a topic with the specified name, that topic's ARN is
// returned without creating a new topic.
//
// [Creating an Amazon SNS topic]: https://docs.aws.amazon.com/sns/latest/dg/sns-create-topic.html
func sns_CreateTopic(cfg aws.Config, client *sns.Client) {
	input := &sns.CreateTopicInput{
		// Name: *string, // Required
	}

	if len(_snsName) > 0 {
		input.Name = aws.String(_snsName)
	}
	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_snsDataProtectionPolicy) > 0 {
		input.DataProtectionPolicy = aws.String(_snsDataProtectionPolicy)
	}
	if len(_snsTags) > 0 {
		if err := assignInputField(input, "Tags", _snsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the endpoint for a device and mobile app from Amazon SNS. This action
// is idempotent. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// When you delete an endpoint that is also subscribed to a topic, then you must
// also unsubscribe the endpoint from the topic.
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_DeleteEndpoint(cfg aws.Config, client *sns.Client) {
	input := &sns.DeleteEndpointInput{
		// EndpointArn: *string, // Required
	}

	if len(_snsEndpointArn) > 0 {
		input.EndpointArn = aws.String(_snsEndpointArn)
	}

	if resp, err := client.DeleteEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a platform application object for one of the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging). For more
// information, see [Using Amazon SNS Mobile Push Notifications].
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_DeletePlatformApplication(cfg aws.Config, client *sns.Client) {
	input := &sns.DeletePlatformApplicationInput{
		// PlatformApplicationArn: *string, // Required
	}

	if len(_snsPlatformApplicationArn) > 0 {
		input.PlatformApplicationArn = aws.String(_snsPlatformApplicationArn)
	}

	if resp, err := client.DeletePlatformApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Web Services account's verified or pending phone number from
// the SMS sandbox.
//
// When you start using Amazon SNS to send SMS messages, your Amazon Web Services
// account is in the SMS sandbox. The SMS sandbox provides a safe environment for
// you to try Amazon SNS features without risking your reputation as an SMS sender.
// While your Amazon Web Services account is in the SMS sandbox, you can use all of
// the features of Amazon SNS. However, you can send SMS messages only to verified
// destination phone numbers. For more information, including how to move out of
// the sandbox to send messages without restrictions, see [SMS sandbox]in the Amazon SNS
// Developer Guide.
//
// [SMS sandbox]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
func sns_DeleteSMSSandboxPhoneNumber(cfg aws.Config, client *sns.Client) {
	input := &sns.DeleteSMSSandboxPhoneNumberInput{
		// PhoneNumber: *string, // Required
	}

	if len(_snsPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_snsPhoneNumber)
	}

	if resp, err := client.DeleteSMSSandboxPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a topic and all its subscriptions. Deleting a topic might prevent some
// messages previously sent to the topic from being delivered to subscribers. This
// action is idempotent, so deleting a topic that does not exist does not result in
// an error.
func sns_DeleteTopic(cfg aws.Config, client *sns.Client) {
	input := &sns.DeleteTopicInput{
		// TopicArn: *string, // Required
	}

	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}

	if resp, err := client.DeleteTopic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified inline DataProtectionPolicy document that is stored in
// the specified Amazon SNS topic.
func sns_GetDataProtectionPolicy(cfg aws.Config, client *sns.Client) {
	input := &sns.GetDataProtectionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_snsResourceArn) > 0 {
		input.ResourceArn = aws.String(_snsResourceArn)
	}

	if resp, err := client.GetDataProtectionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the endpoint attributes for a device on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
// information, see [Using Amazon SNS Mobile Push Notifications].
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_GetEndpointAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.GetEndpointAttributesInput{
		// EndpointArn: *string, // Required
	}

	if len(_snsEndpointArn) > 0 {
		input.EndpointArn = aws.String(_snsEndpointArn)
	}

	if resp, err := client.GetEndpointAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the attributes of the platform application object for the supported
// push notification services, such as APNS and GCM (Firebase Cloud Messaging). For
// more information, see [Using Amazon SNS Mobile Push Notifications].
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_GetPlatformApplicationAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.GetPlatformApplicationAttributesInput{
		// PlatformApplicationArn: *string, // Required
	}

	if len(_snsPlatformApplicationArn) > 0 {
		input.PlatformApplicationArn = aws.String(_snsPlatformApplicationArn)
	}

	if resp, err := client.GetPlatformApplicationAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the settings for sending SMS messages from your Amazon Web Services
// account.
//
// These settings are set with the SetSMSAttributes action.
func sns_GetSMSAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.GetSMSAttributesInput{}

	if len(_snsAttributes) > 0 {
		input.Attributes = []string{_snsAttributes}
	}

	if resp, err := client.GetSMSAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the SMS sandbox status for the calling Amazon Web Services account in
// the target Amazon Web Services Region.
//
// When you start using Amazon SNS to send SMS messages, your Amazon Web Services
// account is in the SMS sandbox. The SMS sandbox provides a safe environment for
// you to try Amazon SNS features without risking your reputation as an SMS sender.
// While your Amazon Web Services account is in the SMS sandbox, you can use all of
// the features of Amazon SNS. However, you can send SMS messages only to verified
// destination phone numbers. For more information, including how to move out of
// the sandbox to send messages without restrictions, see [SMS sandbox]in the Amazon SNS
// Developer Guide.
//
// [SMS sandbox]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
func sns_GetSMSSandboxAccountStatus(cfg aws.Config, client *sns.Client) {
	input := &sns.GetSMSSandboxAccountStatusInput{}

	if resp, err := client.GetSMSSandboxAccountStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all of the properties of a subscription.
func sns_GetSubscriptionAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.GetSubscriptionAttributesInput{
		// SubscriptionArn: *string, // Required
	}

	if len(_snsSubscriptionArn) > 0 {
		input.SubscriptionArn = aws.String(_snsSubscriptionArn)
	}

	if resp, err := client.GetSubscriptionAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all of the properties of a topic. Topic properties returned might
// differ based on the authorization of the user.
func sns_GetTopicAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.GetTopicAttributesInput{
		// TopicArn: *string, // Required
	}

	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}

	if resp, err := client.GetTopicAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the endpoints and endpoint attributes for devices in a supported push
// notification service, such as GCM (Firebase Cloud Messaging) and APNS. The
// results for ListEndpointsByPlatformApplication are paginated and return a
// limited list of endpoints, up to 100. If additional records are available after
// the first page results, then a NextToken string will be returned. To receive the
// next page, you call ListEndpointsByPlatformApplication again using the
// NextToken string received from the previous call. When there are no more records
// to return, NextToken will be null. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// This action is throttled at 30 transactions per second (TPS).
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_ListEndpointsByPlatformApplication(cfg aws.Config, client *sns.Client) {
	input := &sns.ListEndpointsByPlatformApplicationInput{
		// PlatformApplicationArn: *string, // Required
	}

	if len(_snsPlatformApplicationArn) > 0 {
		input.PlatformApplicationArn = aws.String(_snsPlatformApplicationArn)
	}
	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEndpointsByPlatformApplication(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListEndpointsByPlatformApplicationOutput
	p := sns.NewListEndpointsByPlatformApplicationPaginator(client, input)
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

// Lists the calling Amazon Web Services account's dedicated origination numbers
// and their metadata. For more information about origination numbers, see [Origination numbers]in the
// Amazon SNS Developer Guide.
//
// [Origination numbers]: https://docs.aws.amazon.com/sns/latest/dg/channels-sms-originating-identities-origination-numbers.html
func sns_ListOriginationNumbers(cfg aws.Config, client *sns.Client) {
	input := &sns.ListOriginationNumbersInput{}

	if len(_snsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOriginationNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListOriginationNumbersOutput
	p := sns.NewListOriginationNumbersPaginator(client, input)
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

// Returns a list of phone numbers that are opted out, meaning you cannot send SMS
// messages to them.
//
// The results for ListPhoneNumbersOptedOut are paginated, and each page returns
// up to 100 phone numbers. If additional phone numbers are available after the
// first page of results, then a NextToken string will be returned. To receive the
// next page, you call ListPhoneNumbersOptedOut again using the NextToken string
// received from the previous call. When there are no more records to return,
// NextToken will be null.
func sns_ListPhoneNumbersOptedOut(cfg aws.Config, client *sns.Client) {
	input := &sns.ListPhoneNumbersOptedOutInput{}

	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPhoneNumbersOptedOut(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListPhoneNumbersOptedOutOutput
	p := sns.NewListPhoneNumbersOptedOutPaginator(client, input)
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

// Lists the platform application objects for the supported push notification
// services, such as APNS and GCM (Firebase Cloud Messaging). The results for
// ListPlatformApplications are paginated and return a limited list of
// applications, up to 100. If additional records are available after the first
// page results, then a NextToken string will be returned. To receive the next
// page, you call ListPlatformApplications using the NextToken string received
// from the previous call. When there are no more records to return, NextToken
// will be null. For more information, see [Using Amazon SNS Mobile Push Notifications].
//
// This action is throttled at 15 transactions per second (TPS).
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_ListPlatformApplications(cfg aws.Config, client *sns.Client) {
	input := &sns.ListPlatformApplicationsInput{}

	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlatformApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListPlatformApplicationsOutput
	p := sns.NewListPlatformApplicationsPaginator(client, input)
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

// Lists the calling Amazon Web Services account's current verified and pending
// destination phone numbers in the SMS sandbox.
//
// When you start using Amazon SNS to send SMS messages, your Amazon Web Services
// account is in the SMS sandbox. The SMS sandbox provides a safe environment for
// you to try Amazon SNS features without risking your reputation as an SMS sender.
// While your Amazon Web Services account is in the SMS sandbox, you can use all of
// the features of Amazon SNS. However, you can send SMS messages only to verified
// destination phone numbers. For more information, including how to move out of
// the sandbox to send messages without restrictions, see [SMS sandbox]in the Amazon SNS
// Developer Guide.
//
// [SMS sandbox]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
func sns_ListSMSSandboxPhoneNumbers(cfg aws.Config, client *sns.Client) {
	input := &sns.ListSMSSandboxPhoneNumbersInput{}

	if len(_snsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _snsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSMSSandboxPhoneNumbers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListSMSSandboxPhoneNumbersOutput
	p := sns.NewListSMSSandboxPhoneNumbersPaginator(client, input)
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

// Returns a list of the requester's subscriptions. Each call returns a limited
// list of subscriptions, up to 100. If there are more subscriptions, a NextToken
// is also returned. Use the NextToken parameter in a new ListSubscriptions call
// to get further results.
//
// This action is throttled at 30 transactions per second (TPS).
func sns_ListSubscriptions(cfg aws.Config, client *sns.Client) {
	input := &sns.ListSubscriptionsInput{}

	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListSubscriptionsOutput
	p := sns.NewListSubscriptionsPaginator(client, input)
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

// Returns a list of the subscriptions to a specific topic. Each call returns a
// limited list of subscriptions, up to 100. If there are more subscriptions, a
// NextToken is also returned. Use the NextToken parameter in a new
// ListSubscriptionsByTopic call to get further results.
//
// This action is throttled at 30 transactions per second (TPS).
func sns_ListSubscriptionsByTopic(cfg aws.Config, client *sns.Client) {
	input := &sns.ListSubscriptionsByTopicInput{
		// TopicArn: *string, // Required
	}

	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}
	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSubscriptionsByTopic(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListSubscriptionsByTopicOutput
	p := sns.NewListSubscriptionsByTopicPaginator(client, input)
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

// List all tags added to the specified Amazon SNS topic. For an overview, see [Amazon SNS Tags] in
// the Amazon Simple Notification Service Developer Guide.
//
// [Amazon SNS Tags]: https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html
func sns_ListTagsForResource(cfg aws.Config, client *sns.Client) {
	input := &sns.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_snsResourceArn) > 0 {
		input.ResourceArn = aws.String(_snsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the requester's topics. Each call returns a limited list of
// topics, up to 100. If there are more topics, a NextToken is also returned. Use
// the NextToken parameter in a new ListTopics call to get further results.
//
// This action is throttled at 30 transactions per second (TPS).
func sns_ListTopics(cfg aws.Config, client *sns.Client) {
	input := &sns.ListTopicsInput{}

	if len(_snsNextToken) > 0 {
		input.NextToken = aws.String(_snsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTopics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*sns.ListTopicsOutput
	p := sns.NewListTopicsPaginator(client, input)
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

// Use this request to opt in a phone number that is opted out, which enables you
// to resume sending SMS messages to the number.
//
// You can opt in a phone number only once every 30 days.
func sns_OptInPhoneNumber(cfg aws.Config, client *sns.Client) {
	input := &sns.OptInPhoneNumberInput{
		// PhoneNumber: *string, // Required
	}

	if len(_snsPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_snsPhoneNumber)
	}

	if resp, err := client.OptInPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a message to an Amazon SNS topic, a text message (SMS message) directly
// to a phone number, or a message to a mobile platform endpoint (when you specify
// the TargetArn ).
//
// If you send a message to a topic, Amazon SNS delivers the message to each
// endpoint that is subscribed to the topic. The format of the message depends on
// the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the message is saved and Amazon SNS immediately
// delivers it to subscribers.
//
// To use the Publish action for publishing a message to a mobile endpoint, such
// as an app on a Kindle device or mobile phone, you must specify the EndpointArn
// for the TargetArn parameter. The EndpointArn is returned when making a call with
// the CreatePlatformEndpoint action.
//
// For more information about formatting messages, see [Send Custom Platform-Specific Payloads in Messages to Mobile Devices].
//
// You can publish messages only to topics and endpoints in the same Amazon Web
// Services Region.
//
// [Send Custom Platform-Specific Payloads in Messages to Mobile Devices]: https://docs.aws.amazon.com/sns/latest/dg/mobile-push-send-custommessage.html
func sns_Publish(cfg aws.Config, client *sns.Client) {
	input := &sns.PublishInput{
		// Message: *string, // Required
	}

	if len(_snsMessage) > 0 {
		input.Message = aws.String(_snsMessage)
	}
	if len(_snsMessageAttributes) > 0 {
		if err := assignInputField(input, "MessageAttributes", _snsMessageAttributes); err != nil {
			log.Errorf("invalid --message-attributes: %s", err.Error())
			return
		}
	}
	if len(_snsMessageDeduplicationId) > 0 {
		input.MessageDeduplicationId = aws.String(_snsMessageDeduplicationId)
	}
	if len(_snsMessageGroupId) > 0 {
		input.MessageGroupId = aws.String(_snsMessageGroupId)
	}
	if len(_snsMessageStructure) > 0 {
		input.MessageStructure = aws.String(_snsMessageStructure)
	}
	if len(_snsPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_snsPhoneNumber)
	}
	if len(_snsSubject) > 0 {
		input.Subject = aws.String(_snsSubject)
	}
	if len(_snsTargetArn) > 0 {
		input.TargetArn = aws.String(_snsTargetArn)
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}

	if resp, err := client.Publish(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Publishes up to 10 messages to the specified topic in a single batch. This is a
// batch version of the Publish API. If you try to send more than 10 messages in a
// single batch request, you will receive a TooManyEntriesInBatchRequest exception.
//
// For FIFO topics, multiple messages within a single batch are published in the
// order they are sent, and messages are deduplicated within the batch and across
// batches for five minutes.
//
// The result of publishing each message is reported individually in the response.
// Because the batch request can result in a combination of successful and
// unsuccessful actions, you should check for batch errors even when the call
// returns an HTTP status code of 200.
//
// The maximum allowed individual message size and the maximum total payload size
// (the sum of the individual lengths of all of the batched messages) are both 256
// KB (262,144 bytes).
//
// The PublishBatch API can send up to 10 messages at a time. If you attempt to
// send more than 10 messages in one request, you will encounter a
// TooManyEntriesInBatchRequest exception. In such cases, split your messages into
// multiple requests, each containing no more than 10 messages.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example, a
// parameter list with two elements looks like this:
//
// &AttributeName.1=first
//
// &AttributeName.2=second
//
// If you send a batch message to a topic, Amazon SNS publishes the batch message
// to each endpoint that is subscribed to the topic. The format of the batch
// message depends on the notification protocol for each subscribed endpoint.
//
// When a messageId is returned, the batch message is saved, and Amazon SNS
// immediately delivers the message to subscribers.
func sns_PublishBatch(cfg aws.Config, client *sns.Client) {
	input := &sns.PublishBatchInput{
		// PublishBatchRequestEntries: []types.PublishBatchRequestEntry, // Required
		// TopicArn: *string, // Required
	}

	if len(_snsPublishBatchRequestEntries) > 0 {
		if err := assignInputField(input, "PublishBatchRequestEntries", _snsPublishBatchRequestEntries); err != nil {
			log.Errorf("invalid --publish-batch-request-entries: %s", err.Error())
			return
		}
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}

	if resp, err := client.PublishBatch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates an inline policy document that is stored in the specified
// Amazon SNS topic.
func sns_PutDataProtectionPolicy(cfg aws.Config, client *sns.Client) {
	input := &sns.PutDataProtectionPolicyInput{
		// DataProtectionPolicy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_snsDataProtectionPolicy) > 0 {
		input.DataProtectionPolicy = aws.String(_snsDataProtectionPolicy)
	}
	if len(_snsResourceArn) > 0 {
		input.ResourceArn = aws.String(_snsResourceArn)
	}

	if resp, err := client.PutDataProtectionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a statement from a topic's access control policy.
// To remove the ability to change topic permissions, you must deny permissions to
// the AddPermission , RemovePermission , and SetTopicAttributes actions in your
// IAM policy.
func sns_RemovePermission(cfg aws.Config, client *sns.Client) {
	input := &sns.RemovePermissionInput{
		// Label: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_snsLabel) > 0 {
		input.Label = aws.String(_snsLabel)
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}

	if resp, err := client.RemovePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the attributes for an endpoint for a device on one of the supported push
// notification services, such as GCM (Firebase Cloud Messaging) and APNS. For more
// information, see [Using Amazon SNS Mobile Push Notifications].
//
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_SetEndpointAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.SetEndpointAttributesInput{
		// Attributes: map[string]string, // Required
		// EndpointArn: *string, // Required
	}

	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_snsEndpointArn) > 0 {
		input.EndpointArn = aws.String(_snsEndpointArn)
	}

	if resp, err := client.SetEndpointAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the attributes of the platform application object for the supported push
// notification services, such as APNS and GCM (Firebase Cloud Messaging). For more
// information, see [Using Amazon SNS Mobile Push Notifications]. For information on configuring attributes for message
// delivery status, see [Using Amazon SNS Application Attributes for Message Delivery Status].
//
// [Using Amazon SNS Application Attributes for Message Delivery Status]: https://docs.aws.amazon.com/sns/latest/dg/sns-msg-status.html
// [Using Amazon SNS Mobile Push Notifications]: https://docs.aws.amazon.com/sns/latest/dg/SNSMobilePush.html
func sns_SetPlatformApplicationAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.SetPlatformApplicationAttributesInput{
		// Attributes: map[string]string, // Required
		// PlatformApplicationArn: *string, // Required
	}

	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_snsPlatformApplicationArn) > 0 {
		input.PlatformApplicationArn = aws.String(_snsPlatformApplicationArn)
	}

	if resp, err := client.SetPlatformApplicationAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this request to set the default settings for sending SMS messages and
// receiving daily SMS usage reports.
//
// You can override some of these settings for a single message when you use the
// Publish action with the MessageAttributes.entry.N parameter. For more
// information, see [Publishing to a mobile phone]in the Amazon SNS Developer Guide.
//
// To use this operation, you must grant the Amazon SNS service principal (
// sns.amazonaws.com ) permission to perform the s3:ListBucket action.
//
// [Publishing to a mobile phone]: https://docs.aws.amazon.com/sns/latest/dg/sms_publish-to-phone.html
func sns_SetSMSAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.SetSMSAttributesInput{
		// Attributes: map[string]string, // Required
	}

	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetSMSAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a subscription owner to set an attribute of the subscription to a new
// value.
func sns_SetSubscriptionAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.SetSubscriptionAttributesInput{
		// AttributeName: *string, // Required
		// SubscriptionArn: *string, // Required
	}

	if len(_snsAttributeName) > 0 {
		input.AttributeName = aws.String(_snsAttributeName)
	}
	if len(_snsSubscriptionArn) > 0 {
		input.SubscriptionArn = aws.String(_snsSubscriptionArn)
	}
	if len(_snsAttributeValue) > 0 {
		input.AttributeValue = aws.String(_snsAttributeValue)
	}

	if resp, err := client.SetSubscriptionAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a topic owner to set an attribute of the topic to a new value.
// To remove the ability to change topic permissions, you must deny permissions to
// the AddPermission , RemovePermission , and SetTopicAttributes actions in your
// IAM policy.
func sns_SetTopicAttributes(cfg aws.Config, client *sns.Client) {
	input := &sns.SetTopicAttributesInput{
		// AttributeName: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_snsAttributeName) > 0 {
		input.AttributeName = aws.String(_snsAttributeName)
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}
	if len(_snsAttributeValue) > 0 {
		input.AttributeValue = aws.String(_snsAttributeValue)
	}

	if resp, err := client.SetTopicAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Subscribes an endpoint to an Amazon SNS topic. If the endpoint type is HTTP/S
// or email, or if the endpoint and the topic are not in the same Amazon Web
// Services account, the endpoint owner must run the ConfirmSubscription action to
// confirm the subscription.
//
// You call the ConfirmSubscription action with the token from the subscription
// response. Confirmation tokens are valid for two days.
//
// This action is throttled at 100 transactions per second (TPS).
func sns_Subscribe(cfg aws.Config, client *sns.Client) {
	input := &sns.SubscribeInput{
		// Protocol: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_snsProtocol) > 0 {
		input.Protocol = aws.String(_snsProtocol)
	}
	if len(_snsTopicArn) > 0 {
		input.TopicArn = aws.String(_snsTopicArn)
	}
	if len(_snsAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _snsAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_snsEndpoint) > 0 {
		input.Endpoint = aws.String(_snsEndpoint)
	}
	if len(_snsReturnSubscriptionArn) > 0 {
		if err := assignInputField(input, "ReturnSubscriptionArn", _snsReturnSubscriptionArn); err != nil {
			log.Errorf("invalid --return-subscription-arn: %s", err.Error())
			return
		}
	}

	if resp, err := client.Subscribe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add tags to the specified Amazon SNS topic. For an overview, see [Amazon SNS Tags] in the Amazon
// SNS Developer Guide.
//
// When you use topic tags, keep the following guidelines in mind:
//
// - Adding more than 50 tags to a topic isn't recommended.
//
// - Tags don't have any semantic meaning. Amazon SNS interprets tags as
// character strings.
//
// - Tags are case-sensitive.
//
// - A new tag with a key identical to that of an existing tag overwrites the
// existing tag.
//
// - Tagging actions are limited to 10 TPS per Amazon Web Services account, per
// Amazon Web Services Region. If your application requires a higher throughput,
// file a [technical support request].
//
// [technical support request]: https://console.aws.amazon.com/support/home#/case/create?issueType=technical
// [Amazon SNS Tags]: https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html
func sns_TagResource(cfg aws.Config, client *sns.Client) {
	input := &sns.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_snsResourceArn) > 0 {
		input.ResourceArn = aws.String(_snsResourceArn)
	}
	if len(_snsTags) > 0 {
		if err := assignInputField(input, "Tags", _snsTags); err != nil {
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

// Deletes a subscription. If the subscription requires authentication for
// deletion, only the owner of the subscription or the topic's owner can
// unsubscribe, and an Amazon Web Services signature is required. If the
// Unsubscribe call does not require authentication and the requester is not the
// subscription owner, a final cancellation message is delivered to the endpoint,
// so that the endpoint owner can easily resubscribe to the topic if the
// Unsubscribe request was unintended.
//
// This action is throttled at 100 transactions per second (TPS).
func sns_Unsubscribe(cfg aws.Config, client *sns.Client) {
	input := &sns.UnsubscribeInput{
		// SubscriptionArn: *string, // Required
	}

	if len(_snsSubscriptionArn) > 0 {
		input.SubscriptionArn = aws.String(_snsSubscriptionArn)
	}

	if resp, err := client.Unsubscribe(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove tags from the specified Amazon SNS topic. For an overview, see [Amazon SNS Tags] in the
// Amazon SNS Developer Guide.
//
// [Amazon SNS Tags]: https://docs.aws.amazon.com/sns/latest/dg/sns-tags.html
func sns_UntagResource(cfg aws.Config, client *sns.Client) {
	input := &sns.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_snsResourceArn) > 0 {
		input.ResourceArn = aws.String(_snsResourceArn)
	}
	if len(_snsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _snsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies a destination phone number with a one-time password (OTP) for the
// calling Amazon Web Services account.
//
// When you start using Amazon SNS to send SMS messages, your Amazon Web Services
// account is in the SMS sandbox. The SMS sandbox provides a safe environment for
// you to try Amazon SNS features without risking your reputation as an SMS sender.
// While your Amazon Web Services account is in the SMS sandbox, you can use all of
// the features of Amazon SNS. However, you can send SMS messages only to verified
// destination phone numbers. For more information, including how to move out of
// the sandbox to send messages without restrictions, see [SMS sandbox]in the Amazon SNS
// Developer Guide.
//
// [SMS sandbox]: https://docs.aws.amazon.com/sns/latest/dg/sns-sms-sandbox.html
func sns_VerifySMSSandboxPhoneNumber(cfg aws.Config, client *sns.Client) {
	input := &sns.VerifySMSSandboxPhoneNumberInput{
		// OneTimePassword: *string, // Required
		// PhoneNumber: *string, // Required
	}

	if len(_snsOneTimePassword) > 0 {
		input.OneTimePassword = aws.String(_snsOneTimePassword)
	}
	if len(_snsPhoneNumber) > 0 {
		input.PhoneNumber = aws.String(_snsPhoneNumber)
	}

	if resp, err := client.VerifySMSSandboxPhoneNumber(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_snsCmd)
	_snsCmd.Flags().SortFlags = false

	_snsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_snsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_snsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_snsCmd.Flags().StringSliceVarP(&_snsActionName, "action-name", "", nil, "Action Name")
	_snsCmd.Flags().StringVarP(&_snsAttributeName, "attribute-name", "", "", "Attribute Name")
	_snsCmd.Flags().StringVarP(&_snsAttributeValue, "attribute-value", "", "", "Attribute Value")
	_snsCmd.Flags().StringVarP(&_snsAttributes, "attributes", "", "", "Attributes")
	_snsCmd.Flags().StringVarP(&_snsAuthenticateOnUnsubscribe, "authenticate-on-unsubscribe", "", "", "Authenticate On Unsubscribe")
	_snsCmd.Flags().StringSliceVarP(&_snsAWSAccountId, "aws-account-id", "", nil, "AWS Account ID")
	_snsCmd.Flags().StringVarP(&_snsCustomUserData, "custom-user-data", "", "", "Custom User Data")
	_snsCmd.Flags().StringVarP(&_snsDataProtectionPolicy, "data-protection-policy", "", "", "Data Protection Policy")
	_snsCmd.Flags().StringVarP(&_snsEndpoint, "endpoint", "", "", "Endpoint")
	_snsCmd.Flags().StringVarP(&_snsEndpointArn, "endpoint-arn", "", "", "Endpoint ARN")
	_snsCmd.Flags().StringVarP(&_snsLabel, "label", "", "", "Label")
	_snsCmd.Flags().StringVarP(&_snsLanguageCode, "language-code", "", "", "Language Code")
	_snsCmd.Flags().StringVarP(&_snsMaxResults, "max-results", "", "", "Max Results")
	_snsCmd.Flags().StringVarP(&_snsMessage, "message", "", "", "Message")
	_snsCmd.Flags().StringVarP(&_snsMessageAttributes, "message-attributes", "", "", "Message Attributes")
	_snsCmd.Flags().StringVarP(&_snsMessageDeduplicationId, "message-deduplication-id", "", "", "Message Deduplication ID")
	_snsCmd.Flags().StringVarP(&_snsMessageGroupId, "message-group-id", "", "", "Message Group ID")
	_snsCmd.Flags().StringVarP(&_snsMessageStructure, "message-structure", "", "", "Message Structure")
	_snsCmd.Flags().StringVarP(&_snsName, "name", "", "", "Name")
	_snsCmd.Flags().StringVarP(&_snsNextToken, "next-token", "", "", "Next Token")
	_snsCmd.Flags().StringVarP(&_snsOneTimePassword, "one-time-password", "", "", "One Time Password")
	_snsCmd.Flags().StringVarP(&_snsPhoneNumber, "phone-number", "", "", "Phone Number")
	_snsCmd.Flags().StringVarP(&_snsPlatform, "platform", "", "", "Platform")
	_snsCmd.Flags().StringVarP(&_snsPlatformApplicationArn, "platform-application-arn", "", "", "Platform Application ARN")
	_snsCmd.Flags().StringVarP(&_snsProtocol, "protocol", "", "", "Protocol")
	_snsCmd.Flags().StringVarP(&_snsPublishBatchRequestEntries, "publish-batch-request-entries", "", "", "Publish Batch Request Entries")
	_snsCmd.Flags().StringVarP(&_snsResourceArn, "resource-arn", "", "", "Resource ARN")
	_snsCmd.Flags().StringVarP(&_snsReturnSubscriptionArn, "return-subscription-arn", "", "", "Return Subscription ARN")
	_snsCmd.Flags().StringVarP(&_snsSubject, "subject", "", "", "Subject")
	_snsCmd.Flags().StringVarP(&_snsSubscriptionArn, "subscription-arn", "", "", "Subscription ARN")
	_snsCmd.Flags().StringSliceVarP(&_snsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_snsCmd.Flags().StringVarP(&_snsTags, "tags", "", "", "Tags")
	_snsCmd.Flags().StringVarP(&_snsTargetArn, "target-arn", "", "", "Target ARN")
	_snsCmd.Flags().StringVarP(&_snsToken, "token", "", "", "Token")
	_snsCmd.Flags().StringVarP(&_snsTopicArn, "topic-arn", "", "", "Topic ARN")

	_snsCmd.Flags().BoolVarP(&_snsAddPermission, "add-permission", "", false, "Add Permission")
	_snsCmd.Flags().BoolVarP(&_snsCheckIfPhoneNumberIsOptedOut, "check-if-phone-number-is-opted-out", "", false, "Check If Phone Number Is Opted Out")
	_snsCmd.Flags().BoolVarP(&_snsConfirmSubscription, "confirm-subscription", "", false, "Confirm Subscription")
	_snsCmd.Flags().BoolVarP(&_snsCreatePlatformApplication, "create-platform-application", "", false, "Create Platform Application")
	_snsCmd.Flags().BoolVarP(&_snsCreatePlatformEndpoint, "create-platform-endpoint", "", false, "Create Platform Endpoint")
	_snsCmd.Flags().BoolVarP(&_snsCreateSMSSandboxPhoneNumber, "create-sms-sandbox-phone-number", "", false, "Create Sms Sandbox Phone Number")
	_snsCmd.Flags().BoolVarP(&_snsCreateTopic, "create-topic", "", false, "Create Topic")
	_snsCmd.Flags().BoolVarP(&_snsDeleteEndpoint, "delete-endpoint", "", false, "Delete Endpoint")
	_snsCmd.Flags().BoolVarP(&_snsDeletePlatformApplication, "delete-platform-application", "", false, "Delete Platform Application")
	_snsCmd.Flags().BoolVarP(&_snsDeleteSMSSandboxPhoneNumber, "delete-sms-sandbox-phone-number", "", false, "Delete Sms Sandbox Phone Number")
	_snsCmd.Flags().BoolVarP(&_snsDeleteTopic, "delete-topic", "", false, "Delete Topic")
	_snsCmd.Flags().BoolVarP(&_snsGetDataProtectionPolicy, "get-data-protection-policy", "", false, "Get Data Protection Policy")
	_snsCmd.Flags().BoolVarP(&_snsGetEndpointAttributes, "get-endpoint-attributes", "", false, "Get Endpoint Attributes")
	_snsCmd.Flags().BoolVarP(&_snsGetPlatformApplicationAttributes, "get-platform-application-attributes", "", false, "Get Platform Application Attributes")
	_snsCmd.Flags().BoolVarP(&_snsGetSMSAttributes, "get-sms-attributes", "", false, "Get Sms Attributes")
	_snsCmd.Flags().BoolVarP(&_snsGetSMSSandboxAccountStatus, "get-sms-sandbox-account-status", "", false, "Get Sms Sandbox Account Status")
	_snsCmd.Flags().BoolVarP(&_snsGetSubscriptionAttributes, "get-subscription-attributes", "", false, "Get Subscription Attributes")
	_snsCmd.Flags().BoolVarP(&_snsGetTopicAttributes, "get-topic-attributes", "", false, "Get Topic Attributes")
	_snsCmd.Flags().BoolVarP(&_snsListEndpointsByPlatformApplication, "list-endpoints-by-platform-application", "", false, "List Endpoints By Platform Application")
	_snsCmd.Flags().BoolVarP(&_snsListOriginationNumbers, "list-origination-numbers", "", false, "List Origination Numbers")
	_snsCmd.Flags().BoolVarP(&_snsListPhoneNumbersOptedOut, "list-phone-numbers-opted-out", "", false, "List Phone Numbers Opted Out")
	_snsCmd.Flags().BoolVarP(&_snsListPlatformApplications, "list-platform-applications", "", false, "List Platform Applications")
	_snsCmd.Flags().BoolVarP(&_snsListSMSSandboxPhoneNumbers, "list-sms-sandbox-phone-numbers", "", false, "List Sms Sandbox Phone Numbers")
	_snsCmd.Flags().BoolVarP(&_snsListSubscriptions, "list-subscriptions", "", false, "List Subscriptions")
	_snsCmd.Flags().BoolVarP(&_snsListSubscriptionsByTopic, "list-subscriptions-by-topic", "", false, "List Subscriptions By Topic")
	_snsCmd.Flags().BoolVarP(&_snsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_snsCmd.Flags().BoolVarP(&_snsListTopics, "list-topics", "", false, "List Topics")
	_snsCmd.Flags().BoolVarP(&_snsOptInPhoneNumber, "opt-in-phone-number", "", false, "Opt In Phone Number")
	_snsCmd.Flags().BoolVarP(&_snsPublish, "publish", "", false, "Publish")
	_snsCmd.Flags().BoolVarP(&_snsPublishBatch, "publish-batch", "", false, "Publish Batch")
	_snsCmd.Flags().BoolVarP(&_snsPutDataProtectionPolicy, "put-data-protection-policy", "", false, "Put Data Protection Policy")
	_snsCmd.Flags().BoolVarP(&_snsRemovePermission, "remove-permission", "", false, "Remove Permission")
	_snsCmd.Flags().BoolVarP(&_snsSetEndpointAttributes, "set-endpoint-attributes", "", false, "Set Endpoint Attributes")
	_snsCmd.Flags().BoolVarP(&_snsSetPlatformApplicationAttributes, "set-platform-application-attributes", "", false, "Set Platform Application Attributes")
	_snsCmd.Flags().BoolVarP(&_snsSetSMSAttributes, "set-sms-attributes", "", false, "Set Sms Attributes")
	_snsCmd.Flags().BoolVarP(&_snsSetSubscriptionAttributes, "set-subscription-attributes", "", false, "Set Subscription Attributes")
	_snsCmd.Flags().BoolVarP(&_snsSetTopicAttributes, "set-topic-attributes", "", false, "Set Topic Attributes")
	_snsCmd.Flags().BoolVarP(&_snsSubscribe, "subscribe", "", false, "Subscribe")
	_snsCmd.Flags().BoolVarP(&_snsTagResource, "tag-resource", "", false, "Tag Resource")
	_snsCmd.Flags().BoolVarP(&_snsUnsubscribe, "unsubscribe", "", false, "Unsubscribe")
	_snsCmd.Flags().BoolVarP(&_snsUntagResource, "untag-resource", "", false, "Untag Resource")
	_snsCmd.Flags().BoolVarP(&_snsVerifySMSSandboxPhoneNumber, "verify-sms-sandbox-phone-number", "", false, "Verify Sms Sandbox Phone Number")

}
