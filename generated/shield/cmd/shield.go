package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/shield"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// shieldCmd represents the shield command
var _shieldCmd = &cobra.Command{
	Use:   "shield",
	Short: "AWS shield CLI",
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
		client := shield.NewFromConfig(cfg)
		if _shieldAssociateDRTLogBucket {
			shield_AssociateDRTLogBucket(cfg, client)
			return
		}
		if _shieldAssociateDRTRole {
			shield_AssociateDRTRole(cfg, client)
			return
		}
		if _shieldAssociateHealthCheck {
			shield_AssociateHealthCheck(cfg, client)
			return
		}
		if _shieldAssociateProactiveEngagementDetails {
			shield_AssociateProactiveEngagementDetails(cfg, client)
			return
		}
		if _shieldCreateProtection {
			shield_CreateProtection(cfg, client)
			return
		}
		if _shieldCreateProtectionGroup {
			shield_CreateProtectionGroup(cfg, client)
			return
		}
		if _shieldCreateSubscription {
			shield_CreateSubscription(cfg, client)
			return
		}
		if _shieldDeleteProtection {
			shield_DeleteProtection(cfg, client)
			return
		}
		if _shieldDeleteProtectionGroup {
			shield_DeleteProtectionGroup(cfg, client)
			return
		}
		if _shieldDeleteSubscription {
			shield_DeleteSubscription(cfg, client)
			return
		}
		if _shieldDescribeAttack {
			shield_DescribeAttack(cfg, client)
			return
		}
		if _shieldDescribeAttackStatistics {
			shield_DescribeAttackStatistics(cfg, client)
			return
		}
		if _shieldDescribeDRTAccess {
			shield_DescribeDRTAccess(cfg, client)
			return
		}
		if _shieldDescribeEmergencyContactSettings {
			shield_DescribeEmergencyContactSettings(cfg, client)
			return
		}
		if _shieldDescribeProtection {
			shield_DescribeProtection(cfg, client)
			return
		}
		if _shieldDescribeProtectionGroup {
			shield_DescribeProtectionGroup(cfg, client)
			return
		}
		if _shieldDescribeSubscription {
			shield_DescribeSubscription(cfg, client)
			return
		}
		if _shieldDisableApplicationLayerAutomaticResponse {
			shield_DisableApplicationLayerAutomaticResponse(cfg, client)
			return
		}
		if _shieldDisableProactiveEngagement {
			shield_DisableProactiveEngagement(cfg, client)
			return
		}
		if _shieldDisassociateDRTLogBucket {
			shield_DisassociateDRTLogBucket(cfg, client)
			return
		}
		if _shieldDisassociateDRTRole {
			shield_DisassociateDRTRole(cfg, client)
			return
		}
		if _shieldDisassociateHealthCheck {
			shield_DisassociateHealthCheck(cfg, client)
			return
		}
		if _shieldEnableApplicationLayerAutomaticResponse {
			shield_EnableApplicationLayerAutomaticResponse(cfg, client)
			return
		}
		if _shieldEnableProactiveEngagement {
			shield_EnableProactiveEngagement(cfg, client)
			return
		}
		if _shieldGetSubscriptionState {
			shield_GetSubscriptionState(cfg, client)
			return
		}
		if _shieldListAttacks {
			shield_ListAttacks(cfg, client)
			return
		}
		if _shieldListProtectionGroups {
			shield_ListProtectionGroups(cfg, client)
			return
		}
		if _shieldListProtections {
			shield_ListProtections(cfg, client)
			return
		}
		if _shieldListResourcesInProtectionGroup {
			shield_ListResourcesInProtectionGroup(cfg, client)
			return
		}
		if _shieldListTagsForResource {
			shield_ListTagsForResource(cfg, client)
			return
		}
		if _shieldTagResource {
			shield_TagResource(cfg, client)
			return
		}
		if _shieldUntagResource {
			shield_UntagResource(cfg, client)
			return
		}
		if _shieldUpdateApplicationLayerAutomaticResponse {
			shield_UpdateApplicationLayerAutomaticResponse(cfg, client)
			return
		}
		if _shieldUpdateEmergencyContactSettings {
			shield_UpdateEmergencyContactSettings(cfg, client)
			return
		}
		if _shieldUpdateProtectionGroup {
			shield_UpdateProtectionGroup(cfg, client)
			return
		}
		if _shieldUpdateSubscription {
			shield_UpdateSubscription(cfg, client)
			return
		}

	},
}

var (
	_shieldAssociateDRTLogBucket                    bool
	_shieldAssociateDRTRole                         bool
	_shieldAssociateHealthCheck                     bool
	_shieldAssociateProactiveEngagementDetails      bool
	_shieldCreateProtection                         bool
	_shieldCreateProtectionGroup                    bool
	_shieldCreateSubscription                       bool
	_shieldDeleteProtection                         bool
	_shieldDeleteProtectionGroup                    bool
	_shieldDeleteSubscription                       bool
	_shieldDescribeAttack                           bool
	_shieldDescribeAttackStatistics                 bool
	_shieldDescribeDRTAccess                        bool
	_shieldDescribeEmergencyContactSettings         bool
	_shieldDescribeProtection                       bool
	_shieldDescribeProtectionGroup                  bool
	_shieldDescribeSubscription                     bool
	_shieldDisableApplicationLayerAutomaticResponse bool
	_shieldDisableProactiveEngagement               bool
	_shieldDisassociateDRTLogBucket                 bool
	_shieldDisassociateDRTRole                      bool
	_shieldDisassociateHealthCheck                  bool
	_shieldEnableApplicationLayerAutomaticResponse  bool
	_shieldEnableProactiveEngagement                bool
	_shieldGetSubscriptionState                     bool
	_shieldListAttacks                              bool
	_shieldListProtectionGroups                     bool
	_shieldListProtections                          bool
	_shieldListResourcesInProtectionGroup           bool
	_shieldListTagsForResource                      bool
	_shieldTagResource                              bool
	_shieldUntagResource                            bool
	_shieldUpdateApplicationLayerAutomaticResponse  bool
	_shieldUpdateEmergencyContactSettings           bool
	_shieldUpdateProtectionGroup                    bool
	_shieldUpdateSubscription                       bool

	_shieldAction               string
	_shieldAggregation          string
	_shieldAttackId             string
	_shieldAutoRenew            string
	_shieldEmergencyContactList string
	_shieldEndTime              string
	_shieldHealthCheckArn       string
	_shieldInclusionFilters     string
	_shieldLogBucket            string
	_shieldMaxResults           string
	_shieldMembers              []string
	_shieldName                 string
	_shieldNextToken            string
	_shieldPattern              string
	_shieldProtectionGroupId    string
	_shieldProtectionId         string
	_shieldResourceARN          string
	_shieldResourceArns         []string
	_shieldResourceType         string
	_shieldRoleArn              string
	_shieldStartTime            string
	_shieldTagKeys              []string
	_shieldTags                 string
)

// Authorizes the Shield Response Team (SRT) to access the specified Amazon S3
// bucket containing log data such as Application Load Balancer access logs,
// CloudFront logs, or logs from third party sources. You can associate up to 10
// Amazon S3 buckets with your subscription.
//
// To use the services of the SRT and make an AssociateDRTLogBucket request, you
// must be subscribed to the [Business Support plan]or the [Enterprise Support plan].
//
// [Enterprise Support plan]: http://aws.amazon.com/premiumsupport/enterprise-support/
// [Business Support plan]: http://aws.amazon.com/premiumsupport/business-support/
func shield_AssociateDRTLogBucket(cfg aws.Config, client *shield.Client) {
	input := &shield.AssociateDRTLogBucketInput{
		// LogBucket: *string, // Required
	}

	if len(_shieldLogBucket) > 0 {
		input.LogBucket = aws.String(_shieldLogBucket)
	}

	if resp, err := client.AssociateDRTLogBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Authorizes the Shield Response Team (SRT) using the specified role, to access
// your Amazon Web Services account to assist with DDoS attack mitigation during
// potential attacks. This enables the SRT to inspect your WAF configuration and
// create or update WAF rules and web ACLs.
//
// You can associate only one RoleArn with your subscription. If you submit an
// AssociateDRTRole request for an account that already has an associated role, the
// new RoleArn will replace the existing RoleArn .
//
// Prior to making the AssociateDRTRole request, you must attach the
// AWSShieldDRTAccessPolicy managed policy to the role that you'll specify in the
// request. You can access this policy in the IAM console at [AWSShieldDRTAccessPolicy]. For more
// information see [Adding and removing IAM identity permissions]. The role must also trust the service principal
// drt.shield.amazonaws.com . For more information, see [IAM JSON policy elements: Principal].
//
// The SRT will have access only to your WAF and Shield resources. By submitting
// this request, you authorize the SRT to inspect your WAF and Shield configuration
// and create and update WAF rules and web ACLs on your behalf. The SRT takes these
// actions only if explicitly authorized by you.
//
// You must have the iam:PassRole permission to make an AssociateDRTRole request.
// For more information, see [Granting a user permissions to pass a role to an Amazon Web Services service].
//
// To use the services of the SRT and make an AssociateDRTRole request, you must
// be subscribed to the [Business Support plan]or the [Enterprise Support plan].
//
// [Adding and removing IAM identity permissions]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_manage-attach-detach.html
// [Enterprise Support plan]: http://aws.amazon.com/premiumsupport/enterprise-support/
// [AWSShieldDRTAccessPolicy]: https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSShieldDRTAccessPolicy
// [IAM JSON policy elements: Principal]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements_principal.html
// [Granting a user permissions to pass a role to an Amazon Web Services service]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [Business Support plan]: http://aws.amazon.com/premiumsupport/business-support/
func shield_AssociateDRTRole(cfg aws.Config, client *shield.Client) {
	input := &shield.AssociateDRTRoleInput{
		// RoleArn: *string, // Required
	}

	if len(_shieldRoleArn) > 0 {
		input.RoleArn = aws.String(_shieldRoleArn)
	}

	if resp, err := client.AssociateDRTRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds health-based detection to the Shield Advanced protection for a resource.
// Shield Advanced health-based detection uses the health of your Amazon Web
// Services resource to improve responsiveness and accuracy in attack detection and
// response.
//
// You define the health check in Route 53 and then associate it with your Shield
// Advanced protection. For more information, see [Shield Advanced Health-Based Detection]in the WAF Developer Guide.
//
// [Shield Advanced Health-Based Detection]: https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html#ddos-advanced-health-check-option
func shield_AssociateHealthCheck(cfg aws.Config, client *shield.Client) {
	input := &shield.AssociateHealthCheckInput{
		// HealthCheckArn: *string, // Required
		// ProtectionId: *string, // Required
	}

	if len(_shieldHealthCheckArn) > 0 {
		input.HealthCheckArn = aws.String(_shieldHealthCheckArn)
	}
	if len(_shieldProtectionId) > 0 {
		input.ProtectionId = aws.String(_shieldProtectionId)
	}

	if resp, err := client.AssociateHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initializes proactive engagement and sets the list of contacts for the Shield
// Response Team (SRT) to use. You must provide at least one phone number in the
// emergency contact list.
//
// After you have initialized proactive engagement using this call, to disable or
// enable proactive engagement, use the calls DisableProactiveEngagement and
// EnableProactiveEngagement .
//
// This call defines the list of email addresses and phone numbers that the SRT
// can use to contact you for escalations to the SRT and to initiate proactive
// customer support.
//
// The contacts that you provide in the request replace any contacts that were
// already defined. If you already have contacts defined and want to use them,
// retrieve the list using DescribeEmergencyContactSettings and then provide it to
// this call.
func shield_AssociateProactiveEngagementDetails(cfg aws.Config, client *shield.Client) {
	input := &shield.AssociateProactiveEngagementDetailsInput{
		// EmergencyContactList: []types.EmergencyContact, // Required
	}

	if len(_shieldEmergencyContactList) > 0 {
		if err := assignInputField(input, "EmergencyContactList", _shieldEmergencyContactList); err != nil {
			log.Errorf("invalid --emergency-contact-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateProactiveEngagementDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables Shield Advanced for a specific Amazon Web Services resource. The
// resource can be an Amazon CloudFront distribution, Amazon Route 53 hosted zone,
// Global Accelerator standard accelerator, Elastic IP Address, Application Load
// Balancer, or a Classic Load Balancer. You can protect Amazon EC2 instances and
// Network Load Balancers by association with protected Amazon EC2 Elastic IP
// addresses.
//
// You can add protection to only a single resource with each CreateProtection
// request. You can add protection to multiple resources at once through the Shield
// Advanced console at [https://console.aws.amazon.com/wafv2/shieldv2#/]. For more information see [Getting Started with Shield Advanced] and [Adding Shield Advanced protection to Amazon Web Services resources].
//
// [Adding Shield Advanced protection to Amazon Web Services resources]: https://docs.aws.amazon.com/waf/latest/developerguide/configure-new-protection.html
// [https://console.aws.amazon.com/wafv2/shieldv2#/]: https://console.aws.amazon.com/wafv2/shieldv2#/
// [Getting Started with Shield Advanced]: https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html
func shield_CreateProtection(cfg aws.Config, client *shield.Client) {
	input := &shield.CreateProtectionInput{
		// Name: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_shieldName) > 0 {
		input.Name = aws.String(_shieldName)
	}
	if len(_shieldResourceARN) > 0 {
		input.ResourceArn = aws.String(_shieldResourceARN)
	}
	if len(_shieldTags) > 0 {
		if err := assignInputField(input, "Tags", _shieldTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a grouping of protected resources so they can be handled as a
// collective. This resource grouping improves the accuracy of detection and
// reduces false positives.
func shield_CreateProtectionGroup(cfg aws.Config, client *shield.Client) {
	input := &shield.CreateProtectionGroupInput{
		// Aggregation: types.ProtectionGroupAggregation, // Required
		// Pattern: types.ProtectionGroupPattern, // Required
		// ProtectionGroupId: *string, // Required
	}

	if len(_shieldAggregation) > 0 {
		if err := assignInputField(input, "Aggregation", _shieldAggregation); err != nil {
			log.Errorf("invalid --aggregation: %s", err.Error())
			return
		}
	}
	if len(_shieldPattern) > 0 {
		if err := assignInputField(input, "Pattern", _shieldPattern); err != nil {
			log.Errorf("invalid --pattern: %s", err.Error())
			return
		}
	}
	if len(_shieldProtectionGroupId) > 0 {
		input.ProtectionGroupId = aws.String(_shieldProtectionGroupId)
	}
	if len(_shieldMembers) > 0 {
		input.Members = append([]string(nil), _shieldMembers...)
	}
	if len(_shieldResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _shieldResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}
	if len(_shieldTags) > 0 {
		if err := assignInputField(input, "Tags", _shieldTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProtectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Activates Shield Advanced for an account.
// For accounts that are members of an Organizations organization, Shield Advanced
// subscriptions are billed against the organization's payer account, regardless of
// whether the payer account itself is subscribed.
//
// When you initially create a subscription, your subscription is set to be
// automatically renewed at the end of the existing subscription period. You can
// change this by submitting an UpdateSubscription request.
func shield_CreateSubscription(cfg aws.Config, client *shield.Client) {
	input := &shield.CreateSubscriptionInput{}

	if resp, err := client.CreateSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Shield Advanced Protection.
func shield_DeleteProtection(cfg aws.Config, client *shield.Client) {
	input := &shield.DeleteProtectionInput{
		// ProtectionId: *string, // Required
	}

	if len(_shieldProtectionId) > 0 {
		input.ProtectionId = aws.String(_shieldProtectionId)
	}

	if resp, err := client.DeleteProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified protection group.
func shield_DeleteProtectionGroup(cfg aws.Config, client *shield.Client) {
	input := &shield.DeleteProtectionGroupInput{
		// ProtectionGroupId: *string, // Required
	}

	if len(_shieldProtectionGroupId) > 0 {
		input.ProtectionGroupId = aws.String(_shieldProtectionGroupId)
	}

	if resp, err := client.DeleteProtectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes Shield Advanced from an account. Shield Advanced requires a 1-year
// subscription commitment. You cannot delete a subscription prior to the
// completion of that commitment.
//
// Deprecated: This operation has been deprecated.
func shield_DeleteSubscription(cfg aws.Config, client *shield.Client) {
	input := &shield.DeleteSubscriptionInput{}

	if resp, err := client.DeleteSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the details of a DDoS attack.
func shield_DescribeAttack(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeAttackInput{
		// AttackId: *string, // Required
	}

	if len(_shieldAttackId) > 0 {
		input.AttackId = aws.String(_shieldAttackId)
	}

	if resp, err := client.DescribeAttack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides information about the number and type of attacks Shield has detected
// in the last year for all resources that belong to your account, regardless of
// whether you've defined Shield protections for them. This operation is available
// to Shield customers as well as to Shield Advanced customers.
//
// The operation returns data for the time range of midnight UTC, one year ago, to
// midnight UTC, today. For example, if the current time is 2020-10-26 15:39:32 PDT
// , equal to 2020-10-26 22:39:32 UTC , then the time range for the attack data
// returned is from 2019-10-26 00:00:00 UTC to 2020-10-26 00:00:00 UTC .
//
// The time range indicates the period covered by the attack statistics data items.
func shield_DescribeAttackStatistics(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeAttackStatisticsInput{}

	if resp, err := client.DescribeAttackStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the current role and list of Amazon S3 log buckets used by the Shield
// Response Team (SRT) to access your Amazon Web Services account while assisting
// with attack mitigation.
func shield_DescribeDRTAccess(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeDRTAccessInput{}

	if resp, err := client.DescribeDRTAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of email addresses and phone numbers that the Shield Response Team (SRT)
// can use to contact you if you have proactive engagement enabled, for escalations
// to the SRT and to initiate proactive customer support.
func shield_DescribeEmergencyContactSettings(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeEmergencyContactSettingsInput{}

	if resp, err := client.DescribeEmergencyContactSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the details of a Protection object.
func shield_DescribeProtection(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeProtectionInput{}

	if len(_shieldProtectionId) > 0 {
		input.ProtectionId = aws.String(_shieldProtectionId)
	}
	if len(_shieldResourceARN) > 0 {
		input.ResourceArn = aws.String(_shieldResourceARN)
	}

	if resp, err := client.DescribeProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the specification for the specified protection group.
func shield_DescribeProtectionGroup(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeProtectionGroupInput{
		// ProtectionGroupId: *string, // Required
	}

	if len(_shieldProtectionGroupId) > 0 {
		input.ProtectionGroupId = aws.String(_shieldProtectionGroupId)
	}

	if resp, err := client.DescribeProtectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides details about the Shield Advanced subscription for an account.
func shield_DescribeSubscription(cfg aws.Config, client *shield.Client) {
	input := &shield.DescribeSubscriptionInput{}

	if resp, err := client.DescribeSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disable the Shield Advanced automatic application layer DDoS mitigation feature
// for the protected resource. This stops Shield Advanced from creating, verifying,
// and applying WAF rules for attacks that it detects for the resource.
func shield_DisableApplicationLayerAutomaticResponse(cfg aws.Config, client *shield.Client) {
	input := &shield.DisableApplicationLayerAutomaticResponseInput{
		// ResourceArn: *string, // Required
	}

	if len(_shieldResourceARN) > 0 {
		input.ResourceArn = aws.String(_shieldResourceARN)
	}

	if resp, err := client.DisableApplicationLayerAutomaticResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes authorization from the Shield Response Team (SRT) to notify contacts
// about escalations to the SRT and to initiate proactive customer support.
func shield_DisableProactiveEngagement(cfg aws.Config, client *shield.Client) {
	input := &shield.DisableProactiveEngagementInput{}

	if resp, err := client.DisableProactiveEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the Shield Response Team's (SRT) access to the specified Amazon S3
// bucket containing the logs that you shared previously.
func shield_DisassociateDRTLogBucket(cfg aws.Config, client *shield.Client) {
	input := &shield.DisassociateDRTLogBucketInput{
		// LogBucket: *string, // Required
	}

	if len(_shieldLogBucket) > 0 {
		input.LogBucket = aws.String(_shieldLogBucket)
	}

	if resp, err := client.DisassociateDRTLogBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the Shield Response Team's (SRT) access to your Amazon Web Services
// account.
func shield_DisassociateDRTRole(cfg aws.Config, client *shield.Client) {
	input := &shield.DisassociateDRTRoleInput{}

	if resp, err := client.DisassociateDRTRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes health-based detection from the Shield Advanced protection for a
// resource. Shield Advanced health-based detection uses the health of your Amazon
// Web Services resource to improve responsiveness and accuracy in attack detection
// and response.
//
// You define the health check in Route 53 and then associate or disassociate it
// with your Shield Advanced protection. For more information, see [Shield Advanced Health-Based Detection]in the WAF
// Developer Guide.
//
// [Shield Advanced Health-Based Detection]: https://docs.aws.amazon.com/waf/latest/developerguide/ddos-overview.html#ddos-advanced-health-check-option
func shield_DisassociateHealthCheck(cfg aws.Config, client *shield.Client) {
	input := &shield.DisassociateHealthCheckInput{
		// HealthCheckArn: *string, // Required
		// ProtectionId: *string, // Required
	}

	if len(_shieldHealthCheckArn) > 0 {
		input.HealthCheckArn = aws.String(_shieldHealthCheckArn)
	}
	if len(_shieldProtectionId) > 0 {
		input.ProtectionId = aws.String(_shieldProtectionId)
	}

	if resp, err := client.DisassociateHealthCheck(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable the Shield Advanced automatic application layer DDoS mitigation for the
// protected resource.
//
// This feature is available for Amazon CloudFront distributions and Application
// Load Balancers only.
//
// This causes Shield Advanced to create, verify, and apply WAF rules for DDoS
// attacks that it detects for the resource. Shield Advanced applies the rules in a
// Shield rule group inside the web ACL that you've associated with the resource.
// For information about how automatic mitigation works and the requirements for
// using it, see [Shield Advanced automatic application layer DDoS mitigation].
//
// Don't use this action to make changes to automatic mitigation settings when
// it's already enabled for a resource. Instead, use UpdateApplicationLayerAutomaticResponse.
//
// To use this feature, you must associate a web ACL with the protected resource.
// The web ACL must be created using the latest version of WAF (v2). You can
// associate the web ACL through the Shield Advanced console at [https://console.aws.amazon.com/wafv2/shieldv2#/]. For more
// information, see [Getting Started with Shield Advanced]. You can also associate the web ACL to the resource through
// the WAF console or the WAF API, but you must manage Shield Advanced automatic
// mitigation through Shield Advanced. For information about WAF, see [WAF Developer Guide].
//
// [https://console.aws.amazon.com/wafv2/shieldv2#/]: https://console.aws.amazon.com/wafv2/shieldv2#/
// [Getting Started with Shield Advanced]: https://docs.aws.amazon.com/waf/latest/developerguide/getting-started-ddos.html
// [WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [Shield Advanced automatic application layer DDoS mitigation]: https://docs.aws.amazon.com/waf/latest/developerguide/ddos-advanced-automatic-app-layer-response.html
func shield_EnableApplicationLayerAutomaticResponse(cfg aws.Config, client *shield.Client) {
	input := &shield.EnableApplicationLayerAutomaticResponseInput{
		// Action: *types.ResponseAction, // Required
		// ResourceArn: *string, // Required
	}

	if len(_shieldAction) > 0 {
		if err := assignInputField(input, "Action", _shieldAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_shieldResourceARN) > 0 {
		input.ResourceArn = aws.String(_shieldResourceARN)
	}

	if resp, err := client.EnableApplicationLayerAutomaticResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Authorizes the Shield Response Team (SRT) to use email and phone to notify
// contacts about escalations to the SRT and to initiate proactive customer
// support.
func shield_EnableProactiveEngagement(cfg aws.Config, client *shield.Client) {
	input := &shield.EnableProactiveEngagementInput{}

	if resp, err := client.EnableProactiveEngagement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the SubscriptionState , either Active or Inactive .
func shield_GetSubscriptionState(cfg aws.Config, client *shield.Client) {
	input := &shield.GetSubscriptionStateInput{}

	if resp, err := client.GetSubscriptionState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns all ongoing DDoS attacks or all DDoS attacks during a specified time
// period.
func shield_ListAttacks(cfg aws.Config, client *shield.Client) {
	input := &shield.ListAttacksInput{}

	if len(_shieldEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _shieldEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_shieldMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _shieldMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_shieldNextToken) > 0 {
		input.NextToken = aws.String(_shieldNextToken)
	}
	if len(_shieldResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _shieldResourceArns...)
	}
	if len(_shieldStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _shieldStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAttacks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*shield.ListAttacksOutput
	p := shield.NewListAttacksPaginator(client, input)
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

// Retrieves ProtectionGroup objects for the account. You can retrieve all protection groups or
// you can provide filtering criteria and retrieve just the subset of protection
// groups that match the criteria.
func shield_ListProtectionGroups(cfg aws.Config, client *shield.Client) {
	input := &shield.ListProtectionGroupsInput{}

	if len(_shieldInclusionFilters) > 0 {
		if err := assignInputField(input, "InclusionFilters", _shieldInclusionFilters); err != nil {
			log.Errorf("invalid --inclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_shieldMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _shieldMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_shieldNextToken) > 0 {
		input.NextToken = aws.String(_shieldNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProtectionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*shield.ListProtectionGroupsOutput
	p := shield.NewListProtectionGroupsPaginator(client, input)
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

// Retrieves Protection objects for the account. You can retrieve all protections or you can
// provide filtering criteria and retrieve just the subset of protections that
// match the criteria.
func shield_ListProtections(cfg aws.Config, client *shield.Client) {
	input := &shield.ListProtectionsInput{}

	if len(_shieldInclusionFilters) > 0 {
		if err := assignInputField(input, "InclusionFilters", _shieldInclusionFilters); err != nil {
			log.Errorf("invalid --inclusion-filters: %s", err.Error())
			return
		}
	}
	if len(_shieldMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _shieldMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_shieldNextToken) > 0 {
		input.NextToken = aws.String(_shieldNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProtections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*shield.ListProtectionsOutput
	p := shield.NewListProtectionsPaginator(client, input)
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

// Retrieves the resources that are included in the protection group.
func shield_ListResourcesInProtectionGroup(cfg aws.Config, client *shield.Client) {
	input := &shield.ListResourcesInProtectionGroupInput{
		// ProtectionGroupId: *string, // Required
	}

	if len(_shieldProtectionGroupId) > 0 {
		input.ProtectionGroupId = aws.String(_shieldProtectionGroupId)
	}
	if len(_shieldMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _shieldMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_shieldNextToken) > 0 {
		input.NextToken = aws.String(_shieldNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourcesInProtectionGroup(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*shield.ListResourcesInProtectionGroupOutput
	p := shield.NewListResourcesInProtectionGroupPaginator(client, input)
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

// Gets information about Amazon Web Services tags for a specified Amazon Resource
// Name (ARN) in Shield.
func shield_ListTagsForResource(cfg aws.Config, client *shield.Client) {
	input := &shield.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_shieldResourceARN) > 0 {
		input.ResourceARN = aws.String(_shieldResourceARN)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a resource in Shield.
func shield_TagResource(cfg aws.Config, client *shield.Client) {
	input := &shield.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_shieldResourceARN) > 0 {
		input.ResourceARN = aws.String(_shieldResourceARN)
	}
	if len(_shieldTags) > 0 {
		if err := assignInputField(input, "Tags", _shieldTags); err != nil {
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

// Removes tags from a resource in Shield.
func shield_UntagResource(cfg aws.Config, client *shield.Client) {
	input := &shield.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_shieldResourceARN) > 0 {
		input.ResourceARN = aws.String(_shieldResourceARN)
	}
	if len(_shieldTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _shieldTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Shield Advanced automatic application layer DDoS mitigation
// configuration for the specified resource.
func shield_UpdateApplicationLayerAutomaticResponse(cfg aws.Config, client *shield.Client) {
	input := &shield.UpdateApplicationLayerAutomaticResponseInput{
		// Action: *types.ResponseAction, // Required
		// ResourceArn: *string, // Required
	}

	if len(_shieldAction) > 0 {
		if err := assignInputField(input, "Action", _shieldAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_shieldResourceARN) > 0 {
		input.ResourceArn = aws.String(_shieldResourceARN)
	}

	if resp, err := client.UpdateApplicationLayerAutomaticResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of the list of email addresses and phone numbers that the
// Shield Response Team (SRT) can use to contact you if you have proactive
// engagement enabled, for escalations to the SRT and to initiate proactive
// customer support.
func shield_UpdateEmergencyContactSettings(cfg aws.Config, client *shield.Client) {
	input := &shield.UpdateEmergencyContactSettingsInput{}

	if len(_shieldEmergencyContactList) > 0 {
		if err := assignInputField(input, "EmergencyContactList", _shieldEmergencyContactList); err != nil {
			log.Errorf("invalid --emergency-contact-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEmergencyContactSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing protection group. A protection group is a grouping of
// protected resources so they can be handled as a collective. This resource
// grouping improves the accuracy of detection and reduces false positives.
func shield_UpdateProtectionGroup(cfg aws.Config, client *shield.Client) {
	input := &shield.UpdateProtectionGroupInput{
		// Aggregation: types.ProtectionGroupAggregation, // Required
		// Pattern: types.ProtectionGroupPattern, // Required
		// ProtectionGroupId: *string, // Required
	}

	if len(_shieldAggregation) > 0 {
		if err := assignInputField(input, "Aggregation", _shieldAggregation); err != nil {
			log.Errorf("invalid --aggregation: %s", err.Error())
			return
		}
	}
	if len(_shieldPattern) > 0 {
		if err := assignInputField(input, "Pattern", _shieldPattern); err != nil {
			log.Errorf("invalid --pattern: %s", err.Error())
			return
		}
	}
	if len(_shieldProtectionGroupId) > 0 {
		input.ProtectionGroupId = aws.String(_shieldProtectionGroupId)
	}
	if len(_shieldMembers) > 0 {
		input.Members = append([]string(nil), _shieldMembers...)
	}
	if len(_shieldResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _shieldResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProtectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the details of an existing subscription. Only enter values for
// parameters you want to change. Empty parameters are not updated.
//
// For accounts that are members of an Organizations organization, Shield Advanced
// subscriptions are billed against the organization's payer account, regardless of
// whether the payer account itself is subscribed.
func shield_UpdateSubscription(cfg aws.Config, client *shield.Client) {
	input := &shield.UpdateSubscriptionInput{}

	if len(_shieldAutoRenew) > 0 {
		if err := assignInputField(input, "AutoRenew", _shieldAutoRenew); err != nil {
			log.Errorf("invalid --auto-renew: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_shieldCmd)
	_shieldCmd.Flags().SortFlags = false

	_shieldCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_shieldCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_shieldCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_shieldCmd.Flags().StringVarP(&_shieldAction, "action", "", "", "Action")
	_shieldCmd.Flags().StringVarP(&_shieldAggregation, "aggregation", "", "", "Aggregation")
	_shieldCmd.Flags().StringVarP(&_shieldAttackId, "attack-id", "", "", "Attack ID")
	_shieldCmd.Flags().StringVarP(&_shieldAutoRenew, "auto-renew", "", "", "Auto Renew")
	_shieldCmd.Flags().StringVarP(&_shieldEmergencyContactList, "emergency-contact-list", "", "", "Emergency Contact List")
	_shieldCmd.Flags().StringVarP(&_shieldEndTime, "end-time", "", "", "End Time")
	_shieldCmd.Flags().StringVarP(&_shieldHealthCheckArn, "health-check-arn", "", "", "Health Check ARN")
	_shieldCmd.Flags().StringVarP(&_shieldInclusionFilters, "inclusion-filters", "", "", "Inclusion Filters")
	_shieldCmd.Flags().StringVarP(&_shieldLogBucket, "log-bucket", "", "", "Log Bucket")
	_shieldCmd.Flags().StringVarP(&_shieldMaxResults, "max-results", "", "", "Max Results")
	_shieldCmd.Flags().StringSliceVarP(&_shieldMembers, "members", "", nil, "Members")
	_shieldCmd.Flags().StringVarP(&_shieldName, "name", "", "", "Name")
	_shieldCmd.Flags().StringVarP(&_shieldNextToken, "next-token", "", "", "Next Token")
	_shieldCmd.Flags().StringVarP(&_shieldPattern, "pattern", "", "", "Pattern")
	_shieldCmd.Flags().StringVarP(&_shieldProtectionGroupId, "protection-group-id", "", "", "Protection Group ID")
	_shieldCmd.Flags().StringVarP(&_shieldProtectionId, "protection-id", "", "", "Protection ID")
	_shieldCmd.Flags().StringVarP(&_shieldResourceARN, "resource-arn", "", "", "Resource ARN")
	_shieldCmd.Flags().StringSliceVarP(&_shieldResourceArns, "resource-arns", "", nil, "Resource Arns")
	_shieldCmd.Flags().StringVarP(&_shieldResourceType, "resource-type", "", "", "Resource Type")
	_shieldCmd.Flags().StringVarP(&_shieldRoleArn, "role-arn", "", "", "Role ARN")
	_shieldCmd.Flags().StringVarP(&_shieldStartTime, "start-time", "", "", "Start Time")
	_shieldCmd.Flags().StringSliceVarP(&_shieldTagKeys, "tag-keys", "", nil, "Tag Keys")
	_shieldCmd.Flags().StringVarP(&_shieldTags, "tags", "", "", "Tags")

	_shieldCmd.Flags().BoolVarP(&_shieldAssociateDRTLogBucket, "associate-drt-log-bucket", "", false, "Associate Drt Log Bucket")
	_shieldCmd.Flags().BoolVarP(&_shieldAssociateDRTRole, "associate-drt-role", "", false, "Associate Drt Role")
	_shieldCmd.Flags().BoolVarP(&_shieldAssociateHealthCheck, "associate-health-check", "", false, "Associate Health Check")
	_shieldCmd.Flags().BoolVarP(&_shieldAssociateProactiveEngagementDetails, "associate-proactive-engagement-details", "", false, "Associate Proactive Engagement Details")
	_shieldCmd.Flags().BoolVarP(&_shieldCreateProtection, "create-protection", "", false, "Create Protection")
	_shieldCmd.Flags().BoolVarP(&_shieldCreateProtectionGroup, "create-protection-group", "", false, "Create Protection Group")
	_shieldCmd.Flags().BoolVarP(&_shieldCreateSubscription, "create-subscription", "", false, "Create Subscription")
	_shieldCmd.Flags().BoolVarP(&_shieldDeleteProtection, "delete-protection", "", false, "Delete Protection")
	_shieldCmd.Flags().BoolVarP(&_shieldDeleteProtectionGroup, "delete-protection-group", "", false, "Delete Protection Group")
	_shieldCmd.Flags().BoolVarP(&_shieldDeleteSubscription, "delete-subscription", "", false, "Delete Subscription")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeAttack, "describe-attack", "", false, "Describe Attack")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeAttackStatistics, "describe-attack-statistics", "", false, "Describe Attack Statistics")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeDRTAccess, "describe-drt-access", "", false, "Describe Drt Access")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeEmergencyContactSettings, "describe-emergency-contact-settings", "", false, "Describe Emergency Contact Settings")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeProtection, "describe-protection", "", false, "Describe Protection")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeProtectionGroup, "describe-protection-group", "", false, "Describe Protection Group")
	_shieldCmd.Flags().BoolVarP(&_shieldDescribeSubscription, "describe-subscription", "", false, "Describe Subscription")
	_shieldCmd.Flags().BoolVarP(&_shieldDisableApplicationLayerAutomaticResponse, "disable-application-layer-automatic-response", "", false, "Disable Application Layer Automatic Response")
	_shieldCmd.Flags().BoolVarP(&_shieldDisableProactiveEngagement, "disable-proactive-engagement", "", false, "Disable Proactive Engagement")
	_shieldCmd.Flags().BoolVarP(&_shieldDisassociateDRTLogBucket, "disassociate-drt-log-bucket", "", false, "Disassociate Drt Log Bucket")
	_shieldCmd.Flags().BoolVarP(&_shieldDisassociateDRTRole, "disassociate-drt-role", "", false, "Disassociate Drt Role")
	_shieldCmd.Flags().BoolVarP(&_shieldDisassociateHealthCheck, "disassociate-health-check", "", false, "Disassociate Health Check")
	_shieldCmd.Flags().BoolVarP(&_shieldEnableApplicationLayerAutomaticResponse, "enable-application-layer-automatic-response", "", false, "Enable Application Layer Automatic Response")
	_shieldCmd.Flags().BoolVarP(&_shieldEnableProactiveEngagement, "enable-proactive-engagement", "", false, "Enable Proactive Engagement")
	_shieldCmd.Flags().BoolVarP(&_shieldGetSubscriptionState, "get-subscription-state", "", false, "Get Subscription State")
	_shieldCmd.Flags().BoolVarP(&_shieldListAttacks, "list-attacks", "", false, "List Attacks")
	_shieldCmd.Flags().BoolVarP(&_shieldListProtectionGroups, "list-protection-groups", "", false, "List Protection Groups")
	_shieldCmd.Flags().BoolVarP(&_shieldListProtections, "list-protections", "", false, "List Protections")
	_shieldCmd.Flags().BoolVarP(&_shieldListResourcesInProtectionGroup, "list-resources-in-protection-group", "", false, "List Resources In Protection Group")
	_shieldCmd.Flags().BoolVarP(&_shieldListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_shieldCmd.Flags().BoolVarP(&_shieldTagResource, "tag-resource", "", false, "Tag Resource")
	_shieldCmd.Flags().BoolVarP(&_shieldUntagResource, "untag-resource", "", false, "Untag Resource")
	_shieldCmd.Flags().BoolVarP(&_shieldUpdateApplicationLayerAutomaticResponse, "update-application-layer-automatic-response", "", false, "Update Application Layer Automatic Response")
	_shieldCmd.Flags().BoolVarP(&_shieldUpdateEmergencyContactSettings, "update-emergency-contact-settings", "", false, "Update Emergency Contact Settings")
	_shieldCmd.Flags().BoolVarP(&_shieldUpdateProtectionGroup, "update-protection-group", "", false, "Update Protection Group")
	_shieldCmd.Flags().BoolVarP(&_shieldUpdateSubscription, "update-subscription", "", false, "Update Subscription")

}
