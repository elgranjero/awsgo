package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// wafv2Cmd represents the wafv2 command
var _wafv2Cmd = &cobra.Command{
	Use:   "wafv2",
	Short: "AWS wafv2 CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_wafv2Region) > 0 {
			cfg.Region = _wafv2Region
		}
		client := wafv2.NewFromConfig(cfg)
		if _wafv2AssociateWebACL {
			wafv2_AssociateWebACL(cfg, client)
			return
		}
		if _wafv2CheckCapacity {
			wafv2_CheckCapacity(cfg, client)
			return
		}
		if _wafv2CreateAPIKey {
			wafv2_CreateAPIKey(cfg, client)
			return
		}
		if _wafv2CreateIPSet {
			wafv2_CreateIPSet(cfg, client)
			return
		}
		if _wafv2CreateRegexPatternSet {
			wafv2_CreateRegexPatternSet(cfg, client)
			return
		}
		if _wafv2CreateRuleGroup {
			wafv2_CreateRuleGroup(cfg, client)
			return
		}
		if _wafv2CreateWebACL {
			wafv2_CreateWebACL(cfg, client)
			return
		}
		if _wafv2DeleteAPIKey {
			wafv2_DeleteAPIKey(cfg, client)
			return
		}
		if _wafv2DeleteFirewallManagerRuleGroups {
			wafv2_DeleteFirewallManagerRuleGroups(cfg, client)
			return
		}
		if _wafv2DeleteIPSet {
			wafv2_DeleteIPSet(cfg, client)
			return
		}
		if _wafv2DeleteLoggingConfiguration {
			wafv2_DeleteLoggingConfiguration(cfg, client)
			return
		}
		if _wafv2DeletePermissionPolicy {
			wafv2_DeletePermissionPolicy(cfg, client)
			return
		}
		if _wafv2DeleteRegexPatternSet {
			wafv2_DeleteRegexPatternSet(cfg, client)
			return
		}
		if _wafv2DeleteRuleGroup {
			wafv2_DeleteRuleGroup(cfg, client)
			return
		}
		if _wafv2DeleteWebACL {
			wafv2_DeleteWebACL(cfg, client)
			return
		}
		if _wafv2DescribeAllManagedProducts {
			wafv2_DescribeAllManagedProducts(cfg, client)
			return
		}
		if _wafv2DescribeManagedProductsByVendor {
			wafv2_DescribeManagedProductsByVendor(cfg, client)
			return
		}
		if _wafv2DescribeManagedRuleGroup {
			wafv2_DescribeManagedRuleGroup(cfg, client)
			return
		}
		if _wafv2DisassociateWebACL {
			wafv2_DisassociateWebACL(cfg, client)
			return
		}
		if _wafv2GenerateMobileSdkReleaseUrl {
			wafv2_GenerateMobileSdkReleaseUrl(cfg, client)
			return
		}
		if _wafv2GetDecryptedAPIKey {
			wafv2_GetDecryptedAPIKey(cfg, client)
			return
		}
		if _wafv2GetIPSet {
			wafv2_GetIPSet(cfg, client)
			return
		}
		if _wafv2GetLoggingConfiguration {
			wafv2_GetLoggingConfiguration(cfg, client)
			return
		}
		if _wafv2GetManagedRuleSet {
			wafv2_GetManagedRuleSet(cfg, client)
			return
		}
		if _wafv2GetMobileSdkRelease {
			wafv2_GetMobileSdkRelease(cfg, client)
			return
		}
		if _wafv2GetPermissionPolicy {
			wafv2_GetPermissionPolicy(cfg, client)
			return
		}
		if _wafv2GetRateBasedStatementManagedKeys {
			wafv2_GetRateBasedStatementManagedKeys(cfg, client)
			return
		}
		if _wafv2GetRegexPatternSet {
			wafv2_GetRegexPatternSet(cfg, client)
			return
		}
		if _wafv2GetRuleGroup {
			wafv2_GetRuleGroup(cfg, client)
			return
		}
		if _wafv2GetSampledRequests {
			wafv2_GetSampledRequests(cfg, client)
			return
		}
		if _wafv2GetTopPathStatisticsByTraffic {
			wafv2_GetTopPathStatisticsByTraffic(cfg, client)
			return
		}
		if _wafv2GetWebACL {
			wafv2_GetWebACL(cfg, client)
			return
		}
		if _wafv2GetWebACLForResource {
			wafv2_GetWebACLForResource(cfg, client)
			return
		}
		if _wafv2ListAPIKeys {
			wafv2_ListAPIKeys(cfg, client)
			return
		}
		if _wafv2ListAvailableManagedRuleGroupVersions {
			wafv2_ListAvailableManagedRuleGroupVersions(cfg, client)
			return
		}
		if _wafv2ListAvailableManagedRuleGroups {
			wafv2_ListAvailableManagedRuleGroups(cfg, client)
			return
		}
		if _wafv2ListIPSets {
			wafv2_ListIPSets(cfg, client)
			return
		}
		if _wafv2ListLoggingConfigurations {
			wafv2_ListLoggingConfigurations(cfg, client)
			return
		}
		if _wafv2ListManagedRuleSets {
			wafv2_ListManagedRuleSets(cfg, client)
			return
		}
		if _wafv2ListMobileSdkReleases {
			wafv2_ListMobileSdkReleases(cfg, client)
			return
		}
		if _wafv2ListRegexPatternSets {
			wafv2_ListRegexPatternSets(cfg, client)
			return
		}
		if _wafv2ListResourcesForWebACL {
			wafv2_ListResourcesForWebACL(cfg, client)
			return
		}
		if _wafv2ListRuleGroups {
			wafv2_ListRuleGroups(cfg, client)
			return
		}
		if _wafv2ListTagsForResource {
			wafv2_ListTagsForResource(cfg, client)
			return
		}
		if _wafv2ListWebACLs {
			wafv2_ListWebACLs(cfg, client)
			return
		}
		if _wafv2PutLoggingConfiguration {
			wafv2_PutLoggingConfiguration(cfg, client)
			return
		}
		if _wafv2PutManagedRuleSetVersions {
			wafv2_PutManagedRuleSetVersions(cfg, client)
			return
		}
		if _wafv2PutPermissionPolicy {
			wafv2_PutPermissionPolicy(cfg, client)
			return
		}
		if _wafv2TagResource {
			wafv2_TagResource(cfg, client)
			return
		}
		if _wafv2UntagResource {
			wafv2_UntagResource(cfg, client)
			return
		}
		if _wafv2UpdateIPSet {
			wafv2_UpdateIPSet(cfg, client)
			return
		}
		if _wafv2UpdateManagedRuleSetVersionExpiryDate {
			wafv2_UpdateManagedRuleSetVersionExpiryDate(cfg, client)
			return
		}
		if _wafv2UpdateRegexPatternSet {
			wafv2_UpdateRegexPatternSet(cfg, client)
			return
		}
		if _wafv2UpdateRuleGroup {
			wafv2_UpdateRuleGroup(cfg, client)
			return
		}
		if _wafv2UpdateWebACL {
			wafv2_UpdateWebACL(cfg, client)
			return
		}

	},
}

var (
	_wafv2AssociateWebACL                       bool
	_wafv2CheckCapacity                         bool
	_wafv2CreateAPIKey                          bool
	_wafv2CreateIPSet                           bool
	_wafv2CreateRegexPatternSet                 bool
	_wafv2CreateRuleGroup                       bool
	_wafv2CreateWebACL                          bool
	_wafv2DeleteAPIKey                          bool
	_wafv2DeleteFirewallManagerRuleGroups       bool
	_wafv2DeleteIPSet                           bool
	_wafv2DeleteLoggingConfiguration            bool
	_wafv2DeletePermissionPolicy                bool
	_wafv2DeleteRegexPatternSet                 bool
	_wafv2DeleteRuleGroup                       bool
	_wafv2DeleteWebACL                          bool
	_wafv2DescribeAllManagedProducts            bool
	_wafv2DescribeManagedProductsByVendor       bool
	_wafv2DescribeManagedRuleGroup              bool
	_wafv2DisassociateWebACL                    bool
	_wafv2GenerateMobileSdkReleaseUrl           bool
	_wafv2GetDecryptedAPIKey                    bool
	_wafv2GetIPSet                              bool
	_wafv2GetLoggingConfiguration               bool
	_wafv2GetManagedRuleSet                     bool
	_wafv2GetMobileSdkRelease                   bool
	_wafv2GetPermissionPolicy                   bool
	_wafv2GetRateBasedStatementManagedKeys      bool
	_wafv2GetRegexPatternSet                    bool
	_wafv2GetRuleGroup                          bool
	_wafv2GetSampledRequests                    bool
	_wafv2GetTopPathStatisticsByTraffic         bool
	_wafv2GetWebACL                             bool
	_wafv2GetWebACLForResource                  bool
	_wafv2ListAPIKeys                           bool
	_wafv2ListAvailableManagedRuleGroupVersions bool
	_wafv2ListAvailableManagedRuleGroups        bool
	_wafv2ListIPSets                            bool
	_wafv2ListLoggingConfigurations             bool
	_wafv2ListManagedRuleSets                   bool
	_wafv2ListMobileSdkReleases                 bool
	_wafv2ListRegexPatternSets                  bool
	_wafv2ListResourcesForWebACL                bool
	_wafv2ListRuleGroups                        bool
	_wafv2ListTagsForResource                   bool
	_wafv2ListWebACLs                           bool
	_wafv2PutLoggingConfiguration               bool
	_wafv2PutManagedRuleSetVersions             bool
	_wafv2PutPermissionPolicy                   bool
	_wafv2TagResource                           bool
	_wafv2UntagResource                         bool
	_wafv2UpdateIPSet                           bool
	_wafv2UpdateManagedRuleSetVersionExpiryDate bool
	_wafv2UpdateRegexPatternSet                 bool
	_wafv2UpdateRuleGroup                       bool
	_wafv2UpdateWebACL                          bool

	_wafv2Region string

	_wafv2Addresses                     []string
	_wafv2APIKey                        string
	_wafv2ApplicationConfig             string
	_wafv2ARN                           string
	_wafv2AssociationConfig             string
	_wafv2BotCategory                   string
	_wafv2BotName                       string
	_wafv2BotOrganization               string
	_wafv2Capacity                      string
	_wafv2CaptchaConfig                 string
	_wafv2ChallengeConfig               string
	_wafv2CustomResponseBodies          string
	_wafv2DataProtectionConfig          string
	_wafv2DefaultAction                 string
	_wafv2Description                   string
	_wafv2ExpiryTimestamp               string
	_wafv2Id                            string
	_wafv2IPAddressVersion              string
	_wafv2Limit                         string
	_wafv2LockToken                     string
	_wafv2LogScope                      string
	_wafv2LogType                       string
	_wafv2LoggingConfiguration          string
	_wafv2MaxItems                      string
	_wafv2Name                          string
	_wafv2NextMarker                    string
	_wafv2NumberOfTopTrafficBotsPerPath string
	_wafv2OnSourceDDoSProtectionConfig  string
	_wafv2Platform                      string
	_wafv2Policy                        string
	_wafv2RecommendedVersion            string
	_wafv2RegularExpressionList         string
	_wafv2ReleaseVersion                string
	_wafv2ResourceARN                   string
	_wafv2ResourceType                  string
	_wafv2RuleGroupRuleName             string
	_wafv2RuleMetricName                string
	_wafv2RuleName                      string
	_wafv2Rules                         string
	_wafv2Scope                         string
	_wafv2TagKeys                       []string
	_wafv2Tags                          string
	_wafv2TimeWindow                    string
	_wafv2TokenDomains                  []string
	_wafv2UriPathPrefix                 string
	_wafv2VendorName                    string
	_wafv2VersionName                   string
	_wafv2VersionToExpire               string
	_wafv2VersionsToPublish             string
	_wafv2VisibilityConfig              string
	_wafv2WebACLArn                     string
	_wafv2WebACLLockToken               string
	_wafv2WebACLName                    string
	_wafv2WebACLId                      string
)

// Associates a web ACL with a resource, to protect the resource.
// Use this for all resource types except for Amazon CloudFront distributions. For
// Amazon CloudFront, call UpdateDistribution for the distribution and provide the
// Amazon Resource Name (ARN) of the web ACL in the web ACL ID. For information,
// see [UpdateDistribution]in the Amazon CloudFront Developer Guide.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for AssociateWebACL]in the WAF Developer Guide.
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
//
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
// [Permissions for AssociateWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-AssociateWebACL
func wafv2_AssociateWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.AssociateWebACLInput{
		// ResourceArn: *string, // Required
		// WebACLArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}
	if len(_wafv2WebACLArn) > 0 {
		input.WebACLArn = aws.String(_wafv2WebACLArn)
	}

	if resp, err := client.AssociateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the web ACL capacity unit (WCU) requirements for a specified scope and
// set of rules. You can use this to check the capacity requirements for the rules
// you want to use in a RuleGroupor WebACL.
//
// WAF uses WCUs to calculate and control the operating resources that are used to
// run your rules, rule groups, and web ACLs. WAF calculates capacity differently
// for each rule type, to reflect the relative cost of each rule. Simple rules that
// cost little to run use fewer WCUs than more complex rules that use more
// processing power. Rule group capacity is fixed at creation, which helps users
// plan their web ACL WCU usage when they use a rule group. For more information,
// see [WAF web ACL capacity units (WCU)]in the WAF Developer Guide.
//
// [WAF web ACL capacity units (WCU)]: https://docs.aws.amazon.com/waf/latest/developerguide/aws-waf-capacity-units.html
func wafv2_CheckCapacity(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.CheckCapacityInput{
		// Rules: []types.Rule, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Rules) > 0 {
		if err := assignInputField(input, "Rules", _wafv2Rules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.CheckCapacity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an API key that contains a set of token domains.
// API keys are required for the integration of the CAPTCHA API in your JavaScript
// client applications. The API lets you customize the placement and
// characteristics of the CAPTCHA puzzle for your end users. For more information
// about the CAPTCHA JavaScript integration, see [WAF client application integration]in the WAF Developer Guide.
//
// You can use a single key for up to 5 domains. After you generate a key, you can
// copy it for use in your JavaScript integration.
//
// [WAF client application integration]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html
func wafv2_CreateAPIKey(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.CreateAPIKeyInput{
		// Scope: types.Scope, // Required
		// TokenDomains: []string, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2TokenDomains) > 0 {
		input.TokenDomains = append([]string(nil), _wafv2TokenDomains...)
	}

	if resp, err := client.CreateAPIKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IPSet, which you use to identify web requests that originate from
// specific IP addresses or ranges of IP addresses. For example, if you're
// receiving a lot of requests from a ranges of IP addresses, you can configure WAF
// to block them using an IPSet that lists those IP addresses.
func wafv2_CreateIPSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.CreateIPSetInput{
		// Addresses: []string, // Required
		// IPAddressVersion: types.IPAddressVersion, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Addresses) > 0 {
		input.Addresses = append([]string(nil), _wafv2Addresses...)
	}
	if len(_wafv2IPAddressVersion) > 0 {
		if err := assignInputField(input, "IPAddressVersion", _wafv2IPAddressVersion); err != nil {
			log.Errorf("invalid --ip-address-version: %s", err.Error())
			return
		}
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}
	if len(_wafv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _wafv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a RegexPatternSet, which you reference in a RegexPatternSetReferenceStatement, to have WAF inspect a web request
// component for the specified patterns.
func wafv2_CreateRegexPatternSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.CreateRegexPatternSetInput{
		// Name: *string, // Required
		// RegularExpressionList: []types.Regex, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2RegularExpressionList) > 0 {
		if err := assignInputField(input, "RegularExpressionList", _wafv2RegularExpressionList); err != nil {
			log.Errorf("invalid --regular-expression-list: %s", err.Error())
			return
		}
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}
	if len(_wafv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _wafv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a RuleGroup per the specifications provided.
// A rule group defines a collection of rules to inspect and control web requests
// that you can use in a WebACL. When you create a rule group, you define an immutable
// capacity limit. If you update a rule group, you must stay within the capacity.
// This allows others to reuse the rule group with confidence in its capacity
// requirements.
func wafv2_CreateRuleGroup(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.CreateRuleGroupInput{
		// Capacity: *int64, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VisibilityConfig: *types.VisibilityConfig, // Required
	}

	if len(_wafv2Capacity) > 0 {
		if err := assignInputField(input, "Capacity", _wafv2Capacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VisibilityConfig) > 0 {
		if err := assignInputField(input, "VisibilityConfig", _wafv2VisibilityConfig); err != nil {
			log.Errorf("invalid --visibility-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2CustomResponseBodies) > 0 {
		if err := assignInputField(input, "CustomResponseBodies", _wafv2CustomResponseBodies); err != nil {
			log.Errorf("invalid --custom-response-bodies: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}
	if len(_wafv2Rules) > 0 {
		if err := assignInputField(input, "Rules", _wafv2Rules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_wafv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _wafv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a WebACL per the specifications provided.
// A web ACL defines a collection of rules to use to inspect and control web
// requests. Each rule has a statement that defines what to look for in web
// requests and an action that WAF applies to requests that match the statement. In
// the web ACL, you assign a default action to take (allow, block) for any request
// that does not match any of the rules. The rules in a web ACL can be a
// combination of the types Rule, RuleGroup, and managed rule group. You can associate a web
// ACL with one or more Amazon Web Services resources to protect. The resource
// types include Amazon CloudFront distribution, Amazon API Gateway REST API,
// Application Load Balancer, AppSync GraphQL API, Amazon Cognito user pool, App
// Runner service, Amplify application, and Amazon Web Services Verified Access
// instance.
func wafv2_CreateWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.CreateWebACLInput{
		// DefaultAction: *types.DefaultAction, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VisibilityConfig: *types.VisibilityConfig, // Required
	}

	if len(_wafv2DefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _wafv2DefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VisibilityConfig) > 0 {
		if err := assignInputField(input, "VisibilityConfig", _wafv2VisibilityConfig); err != nil {
			log.Errorf("invalid --visibility-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2ApplicationConfig) > 0 {
		if err := assignInputField(input, "ApplicationConfig", _wafv2ApplicationConfig); err != nil {
			log.Errorf("invalid --application-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2AssociationConfig) > 0 {
		if err := assignInputField(input, "AssociationConfig", _wafv2AssociationConfig); err != nil {
			log.Errorf("invalid --association-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2CaptchaConfig) > 0 {
		if err := assignInputField(input, "CaptchaConfig", _wafv2CaptchaConfig); err != nil {
			log.Errorf("invalid --captcha-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2ChallengeConfig) > 0 {
		if err := assignInputField(input, "ChallengeConfig", _wafv2ChallengeConfig); err != nil {
			log.Errorf("invalid --challenge-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2CustomResponseBodies) > 0 {
		if err := assignInputField(input, "CustomResponseBodies", _wafv2CustomResponseBodies); err != nil {
			log.Errorf("invalid --custom-response-bodies: %s", err.Error())
			return
		}
	}
	if len(_wafv2DataProtectionConfig) > 0 {
		if err := assignInputField(input, "DataProtectionConfig", _wafv2DataProtectionConfig); err != nil {
			log.Errorf("invalid --data-protection-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}
	if len(_wafv2OnSourceDDoSProtectionConfig) > 0 {
		if err := assignInputField(input, "OnSourceDDoSProtectionConfig", _wafv2OnSourceDDoSProtectionConfig); err != nil {
			log.Errorf("invalid --on-source-ddos-protection-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2Rules) > 0 {
		if err := assignInputField(input, "Rules", _wafv2Rules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_wafv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _wafv2Tags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_wafv2TokenDomains) > 0 {
		input.TokenDomains = append([]string(nil), _wafv2TokenDomains...)
	}

	if resp, err := client.CreateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified API key.
// After you delete a key, it can take up to 24 hours for WAF to disallow use of
// the key in all regions.
func wafv2_DeleteAPIKey(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteAPIKeyInput{
		// APIKey: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2APIKey) > 0 {
		input.APIKey = aws.String(_wafv2APIKey)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAPIKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes all rule groups that are managed by Firewall Manager from the specified WebACL
// .
//
// You can only use this if ManagedByFirewallManager and
// RetrofittedByFirewallManager are both false in the web ACL.
func wafv2_DeleteFirewallManagerRuleGroups(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteFirewallManagerRuleGroupsInput{
		// WebACLArn: *string, // Required
		// WebACLLockToken: *string, // Required
	}

	if len(_wafv2WebACLArn) > 0 {
		input.WebACLArn = aws.String(_wafv2WebACLArn)
	}
	if len(_wafv2WebACLLockToken) > 0 {
		input.WebACLLockToken = aws.String(_wafv2WebACLLockToken)
	}

	if resp, err := client.DeleteFirewallManagerRuleGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified IPSet.
func wafv2_DeleteIPSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteIPSetInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the LoggingConfiguration from the specified web ACL.
func wafv2_DeleteLoggingConfiguration(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteLoggingConfigurationInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}
	if len(_wafv2LogScope) > 0 {
		if err := assignInputField(input, "LogScope", _wafv2LogScope); err != nil {
			log.Errorf("invalid --log-scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2LogType) > 0 {
		if err := assignInputField(input, "LogType", _wafv2LogType); err != nil {
			log.Errorf("invalid --log-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes an IAM policy from the specified rule group.
// You must be the owner of the rule group to perform this operation.
func wafv2_DeletePermissionPolicy(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeletePermissionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}

	if resp, err := client.DeletePermissionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified RegexPatternSet.
func wafv2_DeleteRegexPatternSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteRegexPatternSetInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified RuleGroup.
func wafv2_DeleteRuleGroup(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteRuleGroupInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified WebACL.
// You can only use this if ManagedByFirewallManager is false in the web ACL.
//
// Before deleting any web ACL, first disassociate it from all resources.
//
// - To retrieve a list of the resources that are associated with a web ACL, use
// the following calls:
//
// - For Amazon CloudFront distributions, use the CloudFront call
// ListDistributionsByWebACLId . For information, see [ListDistributionsByWebACLId]in the Amazon CloudFront
// API Reference.
//
// - For all other resources, call ListResourcesForWebACL.
//
// - To disassociate a resource from a web ACL, use the following calls:
//
// - For Amazon CloudFront distributions, provide an empty web ACL ID in the
// CloudFront call UpdateDistribution . For information, see [UpdateDistribution]in the Amazon
// CloudFront API Reference.
//
// - For all other resources, call DisassociateWebACL.
//
// [ListDistributionsByWebACLId]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListDistributionsByWebACLId.html
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
func wafv2_DeleteWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DeleteWebACLInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides high-level information for the Amazon Web Services Managed Rules rule
// groups and Amazon Web Services Marketplace managed rule groups.
func wafv2_DescribeAllManagedProducts(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DescribeAllManagedProductsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeAllManagedProducts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides high-level information for the managed rule groups owned by a specific
// vendor.
func wafv2_DescribeManagedProductsByVendor(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DescribeManagedProductsByVendorInput{
		// Scope: types.Scope, // Required
		// VendorName: *string, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VendorName) > 0 {
		input.VendorName = aws.String(_wafv2VendorName)
	}

	if resp, err := client.DescribeManagedProductsByVendor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides high-level information for a managed rule group, including
// descriptions of the rules.
func wafv2_DescribeManagedRuleGroup(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DescribeManagedRuleGroupInput{
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VendorName: *string, // Required
	}

	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VendorName) > 0 {
		input.VendorName = aws.String(_wafv2VendorName)
	}
	if len(_wafv2VersionName) > 0 {
		input.VersionName = aws.String(_wafv2VersionName)
	}

	if resp, err := client.DescribeManagedRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified resource from its web ACL association, if it has
// one.
//
// Use this for all resource types except for Amazon CloudFront distributions. For
// Amazon CloudFront, call UpdateDistribution for the distribution and provide an
// empty web ACL ID. For information, see [UpdateDistribution]in the Amazon CloudFront API Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for DisassociateWebACL]in the WAF Developer Guide.
//
// [Permissions for DisassociateWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-DisassociateWebACL
// [UpdateDistribution]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_UpdateDistribution.html
func wafv2_DisassociateWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.DisassociateWebACLInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}

	if resp, err := client.DisassociateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates a presigned download URL for the specified release of the mobile SDK.
// The mobile SDK is not generally available. Customers who have access to the
// mobile SDK can use it to establish and manage WAF tokens for use in HTTP(S)
// requests from a mobile device to WAF. For more information, see [WAF client application integration]in the WAF
// Developer Guide.
//
// [WAF client application integration]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html
func wafv2_GenerateMobileSdkReleaseUrl(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GenerateMobileSdkReleaseUrlInput{
		// Platform: types.Platform, // Required
		// ReleaseVersion: *string, // Required
	}

	if len(_wafv2Platform) > 0 {
		if err := assignInputField(input, "Platform", _wafv2Platform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_wafv2ReleaseVersion) > 0 {
		input.ReleaseVersion = aws.String(_wafv2ReleaseVersion)
	}

	if resp, err := client.GenerateMobileSdkReleaseUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns your API key in decrypted form. Use this to check the token domains
// that you have defined for the key.
//
// API keys are required for the integration of the CAPTCHA API in your JavaScript
// client applications. The API lets you customize the placement and
// characteristics of the CAPTCHA puzzle for your end users. For more information
// about the CAPTCHA JavaScript integration, see [WAF client application integration]in the WAF Developer Guide.
//
// [WAF client application integration]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html
func wafv2_GetDecryptedAPIKey(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetDecryptedAPIKeyInput{
		// APIKey: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2APIKey) > 0 {
		input.APIKey = aws.String(_wafv2APIKey)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDecryptedAPIKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified IPSet.
func wafv2_GetIPSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetIPSetInput{
		// Id: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the LoggingConfiguration for the specified web ACL.
func wafv2_GetLoggingConfiguration(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetLoggingConfigurationInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}
	if len(_wafv2LogScope) > 0 {
		if err := assignInputField(input, "LogScope", _wafv2LogScope); err != nil {
			log.Errorf("invalid --log-scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2LogType) > 0 {
		if err := assignInputField(input, "LogType", _wafv2LogType); err != nil {
			log.Errorf("invalid --log-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified managed rule set.
// This is intended for use only by vendors of managed rule sets. Vendors are
// Amazon Web Services and Amazon Web Services Marketplace sellers.
//
// Vendors, you can use the managed rule set APIs to provide controlled rollout of
// your versioned managed rule group offerings for your customers. The APIs are
// ListManagedRuleSets , GetManagedRuleSet , PutManagedRuleSetVersions , and
// UpdateManagedRuleSetVersionExpiryDate .
func wafv2_GetManagedRuleSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetManagedRuleSetInput{
		// Id: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetManagedRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information for the specified mobile SDK release, including release
// notes and tags.
//
// The mobile SDK is not generally available. Customers who have access to the
// mobile SDK can use it to establish and manage WAF tokens for use in HTTP(S)
// requests from a mobile device to WAF. For more information, see [WAF client application integration]in the WAF
// Developer Guide.
//
// [WAF client application integration]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html
func wafv2_GetMobileSdkRelease(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetMobileSdkReleaseInput{
		// Platform: types.Platform, // Required
		// ReleaseVersion: *string, // Required
	}

	if len(_wafv2Platform) > 0 {
		if err := assignInputField(input, "Platform", _wafv2Platform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_wafv2ReleaseVersion) > 0 {
		input.ReleaseVersion = aws.String(_wafv2ReleaseVersion)
	}

	if resp, err := client.GetMobileSdkRelease(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the IAM policy that is attached to the specified rule group.
// You must be the owner of the rule group to perform this operation.
func wafv2_GetPermissionPolicy(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetPermissionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}

	if resp, err := client.GetPermissionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the IP addresses that are currently blocked by a rate-based rule
// instance. This is only available for rate-based rules that aggregate solely on
// the IP address or on the forwarded IP address.
//
// The maximum number of addresses that can be blocked for a single rate-based
// rule instance is 10,000. If more than 10,000 addresses exceed the rate limit,
// those with the highest rates are blocked.
//
// For a rate-based rule that you've defined inside a rule group, provide the name
// of the rule group reference statement in your request, in addition to the
// rate-based rule name and the web ACL name.
//
// WAF monitors web requests and manages keys independently for each unique
// combination of web ACL, optional rule group, and rate-based rule. For example,
// if you define a rate-based rule inside a rule group, and then use the rule group
// in a web ACL, WAF monitors web requests and manages keys for that web ACL, rule
// group reference statement, and rate-based rule instance. If you use the same
// rule group in a second web ACL, WAF monitors web requests and manages keys for
// this second usage completely independent of your first.
func wafv2_GetRateBasedStatementManagedKeys(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetRateBasedStatementManagedKeysInput{
		// RuleName: *string, // Required
		// Scope: types.Scope, // Required
		// WebACLId: *string, // Required
		// WebACLName: *string, // Required
	}

	if len(_wafv2RuleName) > 0 {
		input.RuleName = aws.String(_wafv2RuleName)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2WebACLId) > 0 {
		input.WebACLId = aws.String(_wafv2WebACLId)
	}
	if len(_wafv2WebACLName) > 0 {
		input.WebACLName = aws.String(_wafv2WebACLName)
	}
	if len(_wafv2RuleGroupRuleName) > 0 {
		input.RuleGroupRuleName = aws.String(_wafv2RuleGroupRuleName)
	}

	if resp, err := client.GetRateBasedStatementManagedKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified RegexPatternSet.
func wafv2_GetRegexPatternSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetRegexPatternSetInput{
		// Id: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified RuleGroup.
func wafv2_GetRuleGroup(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetRuleGroupInput{}

	if len(_wafv2ARN) > 0 {
		input.ARN = aws.String(_wafv2ARN)
	}
	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about a specified number of requests--a sample--that
// WAF randomly selects from among the first 5,000 requests that your Amazon Web
// Services resource received during a time range that you choose. You can specify
// a sample size of up to 500 requests, and you can specify any time range in the
// previous three hours.
//
// GetSampledRequests returns a time range, which is usually the time range that
// you specified. However, if your resource (such as a CloudFront distribution)
// received 5,000 requests before the specified time range elapsed,
// GetSampledRequests returns an updated time range. This new time range indicates
// the actual period during which WAF selected the requests in the sample.
func wafv2_GetSampledRequests(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetSampledRequestsInput{
		// MaxItems: *int64, // Required
		// RuleMetricName: *string, // Required
		// Scope: types.Scope, // Required
		// TimeWindow: *types.TimeWindow, // Required
		// WebAclArn: *string, // Required
	}

	if len(_wafv2MaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _wafv2MaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_wafv2RuleMetricName) > 0 {
		input.RuleMetricName = aws.String(_wafv2RuleMetricName)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2TimeWindow) > 0 {
		if err := assignInputField(input, "TimeWindow", _wafv2TimeWindow); err != nil {
			log.Errorf("invalid --time-window: %s", err.Error())
			return
		}
	}
	if len(_wafv2WebACLArn) > 0 {
		input.WebAclArn = aws.String(_wafv2WebACLArn)
	}

	if resp, err := client.GetSampledRequests(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves aggregated statistics about the top URI paths accessed by bot traffic
// for a specified web ACL and time window. You can use this operation to analyze
// which paths on your web application receive the most bot traffic and identify
// the specific bots accessing those paths. The operation supports filtering by bot
// category, organization, or name, and allows you to drill down into specific path
// prefixes to view detailed URI-level statistics.
func wafv2_GetTopPathStatisticsByTraffic(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetTopPathStatisticsByTrafficInput{
		// Limit: *int32, // Required
		// NumberOfTopTrafficBotsPerPath: *int32, // Required
		// Scope: types.Scope, // Required
		// TimeWindow: *types.TimeWindow, // Required
		// WebAclArn: *string, // Required
	}

	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NumberOfTopTrafficBotsPerPath) > 0 {
		if err := assignInputField(input, "NumberOfTopTrafficBotsPerPath", _wafv2NumberOfTopTrafficBotsPerPath); err != nil {
			log.Errorf("invalid --number-of-top-traffic-bots-per-path: %s", err.Error())
			return
		}
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2TimeWindow) > 0 {
		if err := assignInputField(input, "TimeWindow", _wafv2TimeWindow); err != nil {
			log.Errorf("invalid --time-window: %s", err.Error())
			return
		}
	}
	if len(_wafv2WebACLArn) > 0 {
		input.WebAclArn = aws.String(_wafv2WebACLArn)
	}
	if len(_wafv2BotCategory) > 0 {
		input.BotCategory = aws.String(_wafv2BotCategory)
	}
	if len(_wafv2BotName) > 0 {
		input.BotName = aws.String(_wafv2BotName)
	}
	if len(_wafv2BotOrganization) > 0 {
		input.BotOrganization = aws.String(_wafv2BotOrganization)
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}
	if len(_wafv2UriPathPrefix) > 0 {
		input.UriPathPrefix = aws.String(_wafv2UriPathPrefix)
	}

	if resp, err := client.GetTopPathStatisticsByTraffic(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the specified WebACL.
func wafv2_GetWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetWebACLInput{}

	if len(_wafv2ARN) > 0 {
		input.ARN = aws.String(_wafv2ARN)
	}
	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the WebACL for the specified resource.
// This call uses GetWebACL , to verify that your account has permission to access
// the retrieved web ACL. If you get an error that indicates that your account
// isn't authorized to perform wafv2:GetWebACL on the resource, that error won't
// be included in your CloudTrail event history.
//
// For Amazon CloudFront, don't use this call. Instead, call the CloudFront action
// GetDistributionConfig . For information, see [GetDistributionConfig] in the Amazon CloudFront API
// Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for GetWebACLForResource]in the WAF Developer Guide.
//
// [GetDistributionConfig]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_GetDistributionConfig.html
// [Permissions for GetWebACLForResource]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-GetWebACLForResource
func wafv2_GetWebACLForResource(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.GetWebACLForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}

	if resp, err := client.GetWebACLForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the API keys that you've defined for the specified scope.
// API keys are required for the integration of the CAPTCHA API in your JavaScript
// client applications. The API lets you customize the placement and
// characteristics of the CAPTCHA puzzle for your end users. For more information
// about the CAPTCHA JavaScript integration, see [WAF client application integration]in the WAF Developer Guide.
//
// [WAF client application integration]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html
func wafv2_ListAPIKeys(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListAPIKeysInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListAPIKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the available versions for the specified managed rule group.
func wafv2_ListAvailableManagedRuleGroupVersions(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListAvailableManagedRuleGroupVersionsInput{
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VendorName: *string, // Required
	}

	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VendorName) > 0 {
		input.VendorName = aws.String(_wafv2VendorName)
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListAvailableManagedRuleGroupVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of managed rule groups that are available for you to use.
// This list includes all Amazon Web Services Managed Rules rule groups and all of
// the Amazon Web Services Marketplace managed rule groups that you're subscribed
// to.
func wafv2_ListAvailableManagedRuleGroups(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListAvailableManagedRuleGroupsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListAvailableManagedRuleGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of IPSetSummary objects for the IP sets that you manage.
func wafv2_ListIPSets(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListIPSetsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListIPSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of your LoggingConfiguration objects.
func wafv2_ListLoggingConfigurations(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListLoggingConfigurationsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2LogScope) > 0 {
		if err := assignInputField(input, "LogScope", _wafv2LogScope); err != nil {
			log.Errorf("invalid --log-scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListLoggingConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the managed rule sets that you own.
// This is intended for use only by vendors of managed rule sets. Vendors are
// Amazon Web Services and Amazon Web Services Marketplace sellers.
//
// Vendors, you can use the managed rule set APIs to provide controlled rollout of
// your versioned managed rule group offerings for your customers. The APIs are
// ListManagedRuleSets , GetManagedRuleSet , PutManagedRuleSetVersions , and
// UpdateManagedRuleSetVersionExpiryDate .
func wafv2_ListManagedRuleSets(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListManagedRuleSetsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListManagedRuleSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of the available releases for the mobile SDK and the specified
// device platform.
//
// The mobile SDK is not generally available. Customers who have access to the
// mobile SDK can use it to establish and manage WAF tokens for use in HTTP(S)
// requests from a mobile device to WAF. For more information, see [WAF client application integration]in the WAF
// Developer Guide.
//
// [WAF client application integration]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-application-integration.html
func wafv2_ListMobileSdkReleases(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListMobileSdkReleasesInput{
		// Platform: types.Platform, // Required
	}

	if len(_wafv2Platform) > 0 {
		if err := assignInputField(input, "Platform", _wafv2Platform); err != nil {
			log.Errorf("invalid --platform: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListMobileSdkReleases(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of RegexPatternSetSummary objects for the regex pattern sets that you manage.
func wafv2_ListRegexPatternSets(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListRegexPatternSetsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListRegexPatternSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of the Amazon Resource Names (ARNs) for the resources that
// are associated with the specified web ACL.
//
// For Amazon CloudFront, don't use this call. Instead, use the CloudFront call
// ListDistributionsByWebACLId . For information, see [ListDistributionsByWebACLId] in the Amazon CloudFront
// API Reference.
//
// # Required permissions for customer-managed IAM policies
//
// This call requires permissions that are specific to the protected resource
// type. For details, see [Permissions for ListResourcesForWebACL]in the WAF Developer Guide.
//
// [Permissions for ListResourcesForWebACL]: https://docs.aws.amazon.com/waf/latest/developerguide/security_iam_service-with-iam.html#security_iam_action-ListResourcesForWebACL
// [ListDistributionsByWebACLId]: https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_ListDistributionsByWebACLId.html
func wafv2_ListResourcesForWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListResourcesForWebACLInput{
		// WebACLArn: *string, // Required
	}

	if len(_wafv2WebACLArn) > 0 {
		input.WebACLArn = aws.String(_wafv2WebACLArn)
	}
	if len(_wafv2ResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _wafv2ResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListResourcesForWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of RuleGroupSummary objects for the rule groups that you manage.
func wafv2_ListRuleGroups(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListRuleGroupsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListRuleGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the TagInfoForResource for the specified resource. Tags are key:value pairs that you
// can use to categorize and manage your resources, for purposes like billing. For
// example, you might set the tag key to "customer" and the value to the customer
// name or ID. You can specify one or more tags to add to each Amazon Web Services
// resource, up to 50 tags for a resource.
//
// You can tag the Amazon Web Services resources that you manage through WAF: web
// ACLs, rule groups, IP sets, and regex pattern sets. You can't manage or view
// tags through the WAF console.
func wafv2_ListTagsForResource(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafv2ResourceARN)
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an array of WebACLSummary objects for the web ACLs that you manage.
func wafv2_ListWebACLs(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.ListWebACLsInput{
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Limit) > 0 {
		if err := assignInputField(input, "Limit", _wafv2Limit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafv2NextMarker) > 0 {
		input.NextMarker = aws.String(_wafv2NextMarker)
	}

	if resp, err := client.ListWebACLs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the specified LoggingConfiguration, to start logging from a web ACL, according to the
// configuration provided.
//
// If you configure data protection for the web ACL, the protection applies to the
// data that WAF sends to the logs.
//
// This operation completely replaces any mutable specifications that you already
// have for a logging configuration with the ones that you provide to this call.
//
// To modify an existing logging configuration, do the following:
//
// - Retrieve it by calling GetLoggingConfiguration
//
// - Update its settings as needed
//
// - Provide the complete logging configuration specification to this call
//
// You can define one logging destination per web ACL.
//
// You can access information about the traffic that WAF inspects using the
// following steps:
//
// - Create your logging destination. You can use an Amazon CloudWatch Logs log
// group, an Amazon Simple Storage Service (Amazon S3) bucket, or an Amazon Kinesis
// Data Firehose.
//
// The name that you give the destination must start with aws-waf-logs- . Depending
//
// on the type of destination, you might need to configure additional settings or
// permissions.
//
// # For configuration requirements and pricing information for each destination
//
// type, see [Logging web ACL traffic]in the WAF Developer Guide.
//
// - Associate your logging destination to your web ACL using a
// PutLoggingConfiguration request.
//
// When you successfully enable logging using a PutLoggingConfiguration request,
// WAF creates an additional role or policy that is required to write logs to the
// logging destination. For an Amazon CloudWatch Logs log group, WAF creates a
// resource policy on the log group. For an Amazon S3 bucket, WAF creates a bucket
// policy. For an Amazon Kinesis Data Firehose, WAF creates a service-linked role.
//
// For additional information about web ACL logging, see [Logging web ACL traffic information] in the WAF Developer
// Guide.
//
// [Logging web ACL traffic information]: https://docs.aws.amazon.com/waf/latest/developerguide/logging.html
// [Logging web ACL traffic]: https://docs.aws.amazon.com/waf/latest/developerguide/logging.html
func wafv2_PutLoggingConfiguration(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.PutLoggingConfigurationInput{
		// LoggingConfiguration: *types.LoggingConfiguration, // Required
	}

	if len(_wafv2LoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _wafv2LoggingConfiguration); err != nil {
			log.Errorf("invalid --logging-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Defines the versions of your managed rule set that you are offering to the
// customers. Customers see your offerings as managed rule groups with versioning.
//
// This is intended for use only by vendors of managed rule sets. Vendors are
// Amazon Web Services and Amazon Web Services Marketplace sellers.
//
// Vendors, you can use the managed rule set APIs to provide controlled rollout of
// your versioned managed rule group offerings for your customers. The APIs are
// ListManagedRuleSets , GetManagedRuleSet , PutManagedRuleSetVersions , and
// UpdateManagedRuleSetVersionExpiryDate .
//
// Customers retrieve their managed rule group list by calling ListAvailableManagedRuleGroups. The name that you
// provide here for your managed rule set is the name the customer sees for the
// corresponding managed rule group. Customers can retrieve the available versions
// for a managed rule group by calling ListAvailableManagedRuleGroupVersions. You provide a rule group specification
// for each version. For each managed rule set, you must specify a version that you
// recommend using.
//
// To initiate the expiration of a managed rule group version, use UpdateManagedRuleSetVersionExpiryDate.
func wafv2_PutManagedRuleSetVersions(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.PutManagedRuleSetVersionsInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2RecommendedVersion) > 0 {
		input.RecommendedVersion = aws.String(_wafv2RecommendedVersion)
	}
	if len(_wafv2VersionsToPublish) > 0 {
		if err := assignInputField(input, "VersionsToPublish", _wafv2VersionsToPublish); err != nil {
			log.Errorf("invalid --versions-to-publish: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutManagedRuleSetVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this to share a rule group with other accounts.
// This action attaches an IAM policy to the specified resource. You must be the
// owner of the rule group to perform this operation.
//
// This action is subject to the following restrictions:
//
// - You can attach only one policy with each PutPermissionPolicy request.
//
// - The ARN in the request must be a valid WAF RuleGroupARN and the rule group must
// exist in the same Region.
//
// - The user making the request must be the owner of the rule group.
//
// If a rule group has been shared with your account, you can access it through
// the call GetRuleGroup , and you can reference it in CreateWebACL and
// UpdateWebACL . Rule groups that are shared with you don't appear in your WAF
// console rule groups listing.
func wafv2_PutPermissionPolicy(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.PutPermissionPolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_wafv2Policy) > 0 {
		input.Policy = aws.String(_wafv2Policy)
	}
	if len(_wafv2ResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafv2ResourceARN)
	}

	if resp, err := client.PutPermissionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates tags with the specified Amazon Web Services resource. Tags are
// key:value pairs that you can use to categorize and manage your resources, for
// purposes like billing. For example, you might set the tag key to "customer" and
// the value to the customer name or ID. You can specify one or more tags to add to
// each Amazon Web Services resource, up to 50 tags for a resource.
//
// You can tag the Amazon Web Services resources that you manage through WAF: web
// ACLs, rule groups, IP sets, and regex pattern sets. You can't manage or view
// tags through the WAF console.
func wafv2_TagResource(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafv2ResourceARN)
	}
	if len(_wafv2Tags) > 0 {
		if err := assignInputField(input, "Tags", _wafv2Tags); err != nil {
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

// Disassociates tags from an Amazon Web Services resource. Tags are key:value
// pairs that you can associate with Amazon Web Services resources. For example,
// the tag key might be "customer" and the tag value might be "companyA." You can
// specify one or more tags to add to each container. You can add up to 50 tags to
// each Amazon Web Services resource.
func wafv2_UntagResource(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_wafv2ResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafv2ResourceARN)
	}
	if len(_wafv2TagKeys) > 0 {
		input.TagKeys = append([]string(nil), _wafv2TagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified IPSet.
// This operation completely replaces the mutable specifications that you already
// have for the IP set with the ones that you provide to this call.
//
// To modify an IP set, do the following:
//
// - Retrieve it by calling GetIPSet
//
// - Update its settings as needed
//
// - Provide the complete IP set specification to this call
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
func wafv2_UpdateIPSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.UpdateIPSetInput{
		// Addresses: []string, // Required
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Addresses) > 0 {
		input.Addresses = append([]string(nil), _wafv2Addresses...)
	}
	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}

	if resp, err := client.UpdateIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the expiration information for your managed rule set. Use this to
// initiate the expiration of a managed rule group version. After you initiate
// expiration for a version, WAF excludes it from the response to ListAvailableManagedRuleGroupVersionsfor the managed
// rule group.
//
// This is intended for use only by vendors of managed rule sets. Vendors are
// Amazon Web Services and Amazon Web Services Marketplace sellers.
//
// Vendors, you can use the managed rule set APIs to provide controlled rollout of
// your versioned managed rule group offerings for your customers. The APIs are
// ListManagedRuleSets , GetManagedRuleSet , PutManagedRuleSetVersions , and
// UpdateManagedRuleSetVersionExpiryDate .
func wafv2_UpdateManagedRuleSetVersionExpiryDate(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.UpdateManagedRuleSetVersionExpiryDateInput{
		// ExpiryTimestamp: *time.Time, // Required
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VersionToExpire: *string, // Required
	}

	if len(_wafv2ExpiryTimestamp) > 0 {
		if err := assignInputField(input, "ExpiryTimestamp", _wafv2ExpiryTimestamp); err != nil {
			log.Errorf("invalid --expiry-timestamp: %s", err.Error())
			return
		}
	}
	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VersionToExpire) > 0 {
		input.VersionToExpire = aws.String(_wafv2VersionToExpire)
	}

	if resp, err := client.UpdateManagedRuleSetVersionExpiryDate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified RegexPatternSet.
// This operation completely replaces the mutable specifications that you already
// have for the regex pattern set with the ones that you provide to this call.
//
// To modify a regex pattern set, do the following:
//
// - Retrieve it by calling GetRegexPatternSet
//
// - Update its settings as needed
//
// - Provide the complete regex pattern set specification to this call
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
func wafv2_UpdateRegexPatternSet(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.UpdateRegexPatternSetInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// RegularExpressionList: []types.Regex, // Required
		// Scope: types.Scope, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2RegularExpressionList) > 0 {
		if err := assignInputField(input, "RegularExpressionList", _wafv2RegularExpressionList); err != nil {
			log.Errorf("invalid --regular-expression-list: %s", err.Error())
			return
		}
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}

	if resp, err := client.UpdateRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified RuleGroup.
// This operation completely replaces the mutable specifications that you already
// have for the rule group with the ones that you provide to this call.
//
// To modify a rule group, do the following:
//
// - Retrieve it by calling GetRuleGroup
//
// - Update its settings as needed
//
// - Provide the complete rule group specification to this call
//
// A rule group defines a collection of rules to inspect and control web requests
// that you can use in a WebACL. When you create a rule group, you define an immutable
// capacity limit. If you update a rule group, you must stay within the capacity.
// This allows others to reuse the rule group with confidence in its capacity
// requirements.
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
func wafv2_UpdateRuleGroup(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.UpdateRuleGroupInput{
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VisibilityConfig: *types.VisibilityConfig, // Required
	}

	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VisibilityConfig) > 0 {
		if err := assignInputField(input, "VisibilityConfig", _wafv2VisibilityConfig); err != nil {
			log.Errorf("invalid --visibility-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2CustomResponseBodies) > 0 {
		if err := assignInputField(input, "CustomResponseBodies", _wafv2CustomResponseBodies); err != nil {
			log.Errorf("invalid --custom-response-bodies: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}
	if len(_wafv2Rules) > 0 {
		if err := assignInputField(input, "Rules", _wafv2Rules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified WebACL. While updating a web ACL, WAF provides continuous
// coverage to the resources that you have associated with the web ACL.
//
// This operation completely replaces the mutable specifications that you already
// have for the web ACL with the ones that you provide to this call.
//
// To modify a web ACL, do the following:
//
// - Retrieve it by calling GetWebACL
//
// - Update its settings as needed
//
// - Provide the complete web ACL specification to this call
//
// A web ACL defines a collection of rules to use to inspect and control web
// requests. Each rule has a statement that defines what to look for in web
// requests and an action that WAF applies to requests that match the statement. In
// the web ACL, you assign a default action to take (allow, block) for any request
// that does not match any of the rules. The rules in a web ACL can be a
// combination of the types Rule, RuleGroup, and managed rule group. You can associate a web
// ACL with one or more Amazon Web Services resources to protect. The resource
// types include Amazon CloudFront distribution, Amazon API Gateway REST API,
// Application Load Balancer, AppSync GraphQL API, Amazon Cognito user pool, App
// Runner service, Amplify application, and Amazon Web Services Verified Access
// instance.
//
// # Temporary inconsistencies during updates
//
// When you create or change a web ACL or other WAF resources, the changes take a
// small amount of time to propagate to all areas where the resources are stored.
// The propagation time can be from a few seconds to a number of minutes.
//
// The following are examples of the temporary inconsistencies that you might
// notice during change propagation:
//
// - After you create a web ACL, if you try to associate it with a resource, you
// might get an exception indicating that the web ACL is unavailable.
//
// - After you add a rule group to a web ACL, the new rule group rules might be
// in effect in one area where the web ACL is used and not in another.
//
// - After you change a rule action setting, you might see the old action in
// some places and the new action in others.
//
// - After you add an IP address to an IP set that is in use in a blocking rule,
// the new address might be blocked in one area while still allowed in another.
func wafv2_UpdateWebACL(cfg aws.Config, client *wafv2.Client) {
	input := &wafv2.UpdateWebACLInput{
		// DefaultAction: *types.DefaultAction, // Required
		// Id: *string, // Required
		// LockToken: *string, // Required
		// Name: *string, // Required
		// Scope: types.Scope, // Required
		// VisibilityConfig: *types.VisibilityConfig, // Required
	}

	if len(_wafv2DefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _wafv2DefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_wafv2Id) > 0 {
		input.Id = aws.String(_wafv2Id)
	}
	if len(_wafv2LockToken) > 0 {
		input.LockToken = aws.String(_wafv2LockToken)
	}
	if len(_wafv2Name) > 0 {
		input.Name = aws.String(_wafv2Name)
	}
	if len(_wafv2Scope) > 0 {
		if err := assignInputField(input, "Scope", _wafv2Scope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_wafv2VisibilityConfig) > 0 {
		if err := assignInputField(input, "VisibilityConfig", _wafv2VisibilityConfig); err != nil {
			log.Errorf("invalid --visibility-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2ApplicationConfig) > 0 {
		if err := assignInputField(input, "ApplicationConfig", _wafv2ApplicationConfig); err != nil {
			log.Errorf("invalid --application-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2AssociationConfig) > 0 {
		if err := assignInputField(input, "AssociationConfig", _wafv2AssociationConfig); err != nil {
			log.Errorf("invalid --association-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2CaptchaConfig) > 0 {
		if err := assignInputField(input, "CaptchaConfig", _wafv2CaptchaConfig); err != nil {
			log.Errorf("invalid --captcha-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2ChallengeConfig) > 0 {
		if err := assignInputField(input, "ChallengeConfig", _wafv2ChallengeConfig); err != nil {
			log.Errorf("invalid --challenge-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2CustomResponseBodies) > 0 {
		if err := assignInputField(input, "CustomResponseBodies", _wafv2CustomResponseBodies); err != nil {
			log.Errorf("invalid --custom-response-bodies: %s", err.Error())
			return
		}
	}
	if len(_wafv2DataProtectionConfig) > 0 {
		if err := assignInputField(input, "DataProtectionConfig", _wafv2DataProtectionConfig); err != nil {
			log.Errorf("invalid --data-protection-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2Description) > 0 {
		input.Description = aws.String(_wafv2Description)
	}
	if len(_wafv2OnSourceDDoSProtectionConfig) > 0 {
		if err := assignInputField(input, "OnSourceDDoSProtectionConfig", _wafv2OnSourceDDoSProtectionConfig); err != nil {
			log.Errorf("invalid --on-source-ddos-protection-config: %s", err.Error())
			return
		}
	}
	if len(_wafv2Rules) > 0 {
		if err := assignInputField(input, "Rules", _wafv2Rules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_wafv2TokenDomains) > 0 {
		input.TokenDomains = append([]string(nil), _wafv2TokenDomains...)
	}

	if resp, err := client.UpdateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_wafv2Cmd)
	_wafv2Cmd.Flags().SortFlags = false

	_wafv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Region, "region", "r", "us-east-1", "Set AWS Region")

	_wafv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_wafv2Cmd.Flags().StringSliceVarP(&_wafv2Addresses, "addresses", "", nil, "Addresses")
	_wafv2Cmd.Flags().StringVarP(&_wafv2APIKey, "api-key", "", "", "API Key")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ApplicationConfig, "application-config", "", "", "Application Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ARN, "arn", "", "", "ARN")
	_wafv2Cmd.Flags().StringVarP(&_wafv2AssociationConfig, "association-config", "", "", "Association Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2BotCategory, "bot-category", "", "", "Bot Category")
	_wafv2Cmd.Flags().StringVarP(&_wafv2BotName, "bot-name", "", "", "Bot Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2BotOrganization, "bot-organization", "", "", "Bot Organization")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Capacity, "capacity", "", "", "Capacity")
	_wafv2Cmd.Flags().StringVarP(&_wafv2CaptchaConfig, "captcha-config", "", "", "Captcha Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ChallengeConfig, "challenge-config", "", "", "Challenge Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2CustomResponseBodies, "custom-response-bodies", "", "", "Custom Response Bodies")
	_wafv2Cmd.Flags().StringVarP(&_wafv2DataProtectionConfig, "data-protection-config", "", "", "Data Protection Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2DefaultAction, "default-action", "", "", "Default Action")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Description, "description", "", "", "Description")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ExpiryTimestamp, "expiry-timestamp", "", "", "Expiry Timestamp")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Id, "id", "", "", "ID")
	_wafv2Cmd.Flags().StringVarP(&_wafv2IPAddressVersion, "ip-address-version", "", "", "IP Address Version")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Limit, "limit", "", "", "Limit")
	_wafv2Cmd.Flags().StringVarP(&_wafv2LockToken, "lock-token", "", "", "Lock Token")
	_wafv2Cmd.Flags().StringVarP(&_wafv2LogScope, "log-scope", "", "", "Log Scope")
	_wafv2Cmd.Flags().StringVarP(&_wafv2LogType, "log-type", "", "", "Log Type")
	_wafv2Cmd.Flags().StringVarP(&_wafv2LoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_wafv2Cmd.Flags().StringVarP(&_wafv2MaxItems, "max-items", "", "", "Max Items")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Name, "name", "", "", "Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2NextMarker, "next-marker", "", "", "Next Marker")
	_wafv2Cmd.Flags().StringVarP(&_wafv2NumberOfTopTrafficBotsPerPath, "number-of-top-traffic-bots-per-path", "", "", "Number Of Top Traffic Bots Per Path")
	_wafv2Cmd.Flags().StringVarP(&_wafv2OnSourceDDoSProtectionConfig, "on-source-ddos-protection-config", "", "", "On Source DDOS Protection Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Platform, "platform", "", "", "Platform")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Policy, "policy", "", "", "Policy")
	_wafv2Cmd.Flags().StringVarP(&_wafv2RecommendedVersion, "recommended-version", "", "", "Recommended Version")
	_wafv2Cmd.Flags().StringVarP(&_wafv2RegularExpressionList, "regular-expression-list", "", "", "Regular Expression List")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ReleaseVersion, "release-version", "", "", "Release Version")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ResourceARN, "resource-arn", "", "", "Resource ARN")
	_wafv2Cmd.Flags().StringVarP(&_wafv2ResourceType, "resource-type", "", "", "Resource Type")
	_wafv2Cmd.Flags().StringVarP(&_wafv2RuleGroupRuleName, "rule-group-rule-name", "", "", "Rule Group Rule Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2RuleMetricName, "rule-metric-name", "", "", "Rule Metric Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2RuleName, "rule-name", "", "", "Rule Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Rules, "rules", "", "", "Rules")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Scope, "scope", "", "", "Scope")
	_wafv2Cmd.Flags().StringSliceVarP(&_wafv2TagKeys, "tag-keys", "", nil, "Tag Keys")
	_wafv2Cmd.Flags().StringVarP(&_wafv2Tags, "tags", "", "", "Tags")
	_wafv2Cmd.Flags().StringVarP(&_wafv2TimeWindow, "time-window", "", "", "Time Window")
	_wafv2Cmd.Flags().StringSliceVarP(&_wafv2TokenDomains, "token-domains", "", nil, "Token Domains")
	_wafv2Cmd.Flags().StringVarP(&_wafv2UriPathPrefix, "uri-path-prefix", "", "", "URI Path Prefix")
	_wafv2Cmd.Flags().StringVarP(&_wafv2VendorName, "vendor-name", "", "", "Vendor Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2VersionName, "version-name", "", "", "Version Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2VersionToExpire, "version-to-expire", "", "", "Version To Expire")
	_wafv2Cmd.Flags().StringVarP(&_wafv2VersionsToPublish, "versions-to-publish", "", "", "Versions To Publish")
	_wafv2Cmd.Flags().StringVarP(&_wafv2VisibilityConfig, "visibility-config", "", "", "Visibility Config")
	_wafv2Cmd.Flags().StringVarP(&_wafv2WebACLArn, "web-acl-arn", "", "", "Web ACL ARN")
	_wafv2Cmd.Flags().StringVarP(&_wafv2WebACLLockToken, "web-acl-lock-token", "", "", "Web ACL Lock Token")
	_wafv2Cmd.Flags().StringVarP(&_wafv2WebACLName, "web-acl-name", "", "", "Web ACL Name")
	_wafv2Cmd.Flags().StringVarP(&_wafv2WebACLId, "web-aclid", "", "", "Web Aclid")

	_wafv2Cmd.Flags().BoolVarP(&_wafv2AssociateWebACL, "associate-web-acl", "", false, "Associate Web ACL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2CheckCapacity, "check-capacity", "", false, "Check Capacity")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2CreateAPIKey, "create-api-key", "", false, "Create API Key")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2CreateIPSet, "create-ip-set", "", false, "Create IP Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2CreateRegexPatternSet, "create-regex-pattern-set", "", false, "Create Regex Pattern Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2CreateRuleGroup, "create-rule-group", "", false, "Create Rule Group")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2CreateWebACL, "create-web-acl", "", false, "Create Web ACL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteAPIKey, "delete-api-key", "", false, "Delete API Key")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteFirewallManagerRuleGroups, "delete-firewall-manager-rule-groups", "", false, "Delete Firewall Manager Rule Groups")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteIPSet, "delete-ip-set", "", false, "Delete IP Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteLoggingConfiguration, "delete-logging-configuration", "", false, "Delete Logging Configuration")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeletePermissionPolicy, "delete-permission-policy", "", false, "Delete Permission Policy")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteRegexPatternSet, "delete-regex-pattern-set", "", false, "Delete Regex Pattern Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteRuleGroup, "delete-rule-group", "", false, "Delete Rule Group")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DeleteWebACL, "delete-web-acl", "", false, "Delete Web ACL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DescribeAllManagedProducts, "describe-all-managed-products", "", false, "Describe All Managed Products")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DescribeManagedProductsByVendor, "describe-managed-products-by-vendor", "", false, "Describe Managed Products By Vendor")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DescribeManagedRuleGroup, "describe-managed-rule-group", "", false, "Describe Managed Rule Group")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2DisassociateWebACL, "disassociate-web-acl", "", false, "Disassociate Web ACL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GenerateMobileSdkReleaseUrl, "generate-mobile-sdk-release-url", "", false, "Generate Mobile Sdk Release URL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetDecryptedAPIKey, "get-decrypted-api-key", "", false, "Get Decrypted API Key")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetIPSet, "get-ip-set", "", false, "Get IP Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetLoggingConfiguration, "get-logging-configuration", "", false, "Get Logging Configuration")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetManagedRuleSet, "get-managed-rule-set", "", false, "Get Managed Rule Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetMobileSdkRelease, "get-mobile-sdk-release", "", false, "Get Mobile Sdk Release")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetPermissionPolicy, "get-permission-policy", "", false, "Get Permission Policy")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetRateBasedStatementManagedKeys, "get-rate-based-statement-managed-keys", "", false, "Get Rate Based Statement Managed Keys")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetRegexPatternSet, "get-regex-pattern-set", "", false, "Get Regex Pattern Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetRuleGroup, "get-rule-group", "", false, "Get Rule Group")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetSampledRequests, "get-sampled-requests", "", false, "Get Sampled Requests")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetTopPathStatisticsByTraffic, "get-top-path-statistics-by-traffic", "", false, "Get Top Path Statistics By Traffic")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetWebACL, "get-web-acl", "", false, "Get Web ACL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2GetWebACLForResource, "get-web-acl-for-resource", "", false, "Get Web ACL For Resource")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListAPIKeys, "list-api-keys", "", false, "List API Keys")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListAvailableManagedRuleGroupVersions, "list-available-managed-rule-group-versions", "", false, "List Available Managed Rule Group Versions")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListAvailableManagedRuleGroups, "list-available-managed-rule-groups", "", false, "List Available Managed Rule Groups")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListIPSets, "list-ip-sets", "", false, "List IP Sets")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListLoggingConfigurations, "list-logging-configurations", "", false, "List Logging Configurations")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListManagedRuleSets, "list-managed-rule-sets", "", false, "List Managed Rule Sets")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListMobileSdkReleases, "list-mobile-sdk-releases", "", false, "List Mobile Sdk Releases")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListRegexPatternSets, "list-regex-pattern-sets", "", false, "List Regex Pattern Sets")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListResourcesForWebACL, "list-resources-for-web-acl", "", false, "List Resources For Web ACL")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListRuleGroups, "list-rule-groups", "", false, "List Rule Groups")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2ListWebACLs, "list-web-acls", "", false, "List Web Acls")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2PutLoggingConfiguration, "put-logging-configuration", "", false, "Put Logging Configuration")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2PutManagedRuleSetVersions, "put-managed-rule-set-versions", "", false, "Put Managed Rule Set Versions")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2PutPermissionPolicy, "put-permission-policy", "", false, "Put Permission Policy")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2TagResource, "tag-resource", "", false, "Tag Resource")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2UntagResource, "untag-resource", "", false, "Untag Resource")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2UpdateIPSet, "update-ip-set", "", false, "Update IP Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2UpdateManagedRuleSetVersionExpiryDate, "update-managed-rule-set-version-expiry-date", "", false, "Update Managed Rule Set Version Expiry Date")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2UpdateRegexPatternSet, "update-regex-pattern-set", "", false, "Update Regex Pattern Set")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2UpdateRuleGroup, "update-rule-group", "", false, "Update Rule Group")
	_wafv2Cmd.Flags().BoolVarP(&_wafv2UpdateWebACL, "update-web-acl", "", false, "Update Web ACL")

}
