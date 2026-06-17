package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/wafregional"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// wafregionalCmd represents the wafregional command
var _wafregionalCmd = &cobra.Command{
	Use:   "wafregional",
	Short: "AWS wafregional CLI",
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
		client := wafregional.NewFromConfig(cfg)
		if _wafregionalAssociateWebACL {
			wafregional_AssociateWebACL(cfg, client)
			return
		}
		if _wafregionalCreateByteMatchSet {
			wafregional_CreateByteMatchSet(cfg, client)
			return
		}
		if _wafregionalCreateGeoMatchSet {
			wafregional_CreateGeoMatchSet(cfg, client)
			return
		}
		if _wafregionalCreateIPSet {
			wafregional_CreateIPSet(cfg, client)
			return
		}
		if _wafregionalCreateRateBasedRule {
			wafregional_CreateRateBasedRule(cfg, client)
			return
		}
		if _wafregionalCreateRegexMatchSet {
			wafregional_CreateRegexMatchSet(cfg, client)
			return
		}
		if _wafregionalCreateRegexPatternSet {
			wafregional_CreateRegexPatternSet(cfg, client)
			return
		}
		if _wafregionalCreateRule {
			wafregional_CreateRule(cfg, client)
			return
		}
		if _wafregionalCreateRuleGroup {
			wafregional_CreateRuleGroup(cfg, client)
			return
		}
		if _wafregionalCreateSizeConstraintSet {
			wafregional_CreateSizeConstraintSet(cfg, client)
			return
		}
		if _wafregionalCreateSqlInjectionMatchSet {
			wafregional_CreateSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafregionalCreateWebACL {
			wafregional_CreateWebACL(cfg, client)
			return
		}
		if _wafregionalCreateWebACLMigrationStack {
			wafregional_CreateWebACLMigrationStack(cfg, client)
			return
		}
		if _wafregionalCreateXssMatchSet {
			wafregional_CreateXssMatchSet(cfg, client)
			return
		}
		if _wafregionalDeleteByteMatchSet {
			wafregional_DeleteByteMatchSet(cfg, client)
			return
		}
		if _wafregionalDeleteGeoMatchSet {
			wafregional_DeleteGeoMatchSet(cfg, client)
			return
		}
		if _wafregionalDeleteIPSet {
			wafregional_DeleteIPSet(cfg, client)
			return
		}
		if _wafregionalDeleteLoggingConfiguration {
			wafregional_DeleteLoggingConfiguration(cfg, client)
			return
		}
		if _wafregionalDeletePermissionPolicy {
			wafregional_DeletePermissionPolicy(cfg, client)
			return
		}
		if _wafregionalDeleteRateBasedRule {
			wafregional_DeleteRateBasedRule(cfg, client)
			return
		}
		if _wafregionalDeleteRegexMatchSet {
			wafregional_DeleteRegexMatchSet(cfg, client)
			return
		}
		if _wafregionalDeleteRegexPatternSet {
			wafregional_DeleteRegexPatternSet(cfg, client)
			return
		}
		if _wafregionalDeleteRule {
			wafregional_DeleteRule(cfg, client)
			return
		}
		if _wafregionalDeleteRuleGroup {
			wafregional_DeleteRuleGroup(cfg, client)
			return
		}
		if _wafregionalDeleteSizeConstraintSet {
			wafregional_DeleteSizeConstraintSet(cfg, client)
			return
		}
		if _wafregionalDeleteSqlInjectionMatchSet {
			wafregional_DeleteSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafregionalDeleteWebACL {
			wafregional_DeleteWebACL(cfg, client)
			return
		}
		if _wafregionalDeleteXssMatchSet {
			wafregional_DeleteXssMatchSet(cfg, client)
			return
		}
		if _wafregionalDisassociateWebACL {
			wafregional_DisassociateWebACL(cfg, client)
			return
		}
		if _wafregionalGetByteMatchSet {
			wafregional_GetByteMatchSet(cfg, client)
			return
		}
		if _wafregionalGetChangeToken {
			wafregional_GetChangeToken(cfg, client)
			return
		}
		if _wafregionalGetChangeTokenStatus {
			wafregional_GetChangeTokenStatus(cfg, client)
			return
		}
		if _wafregionalGetGeoMatchSet {
			wafregional_GetGeoMatchSet(cfg, client)
			return
		}
		if _wafregionalGetIPSet {
			wafregional_GetIPSet(cfg, client)
			return
		}
		if _wafregionalGetLoggingConfiguration {
			wafregional_GetLoggingConfiguration(cfg, client)
			return
		}
		if _wafregionalGetPermissionPolicy {
			wafregional_GetPermissionPolicy(cfg, client)
			return
		}
		if _wafregionalGetRateBasedRule {
			wafregional_GetRateBasedRule(cfg, client)
			return
		}
		if _wafregionalGetRateBasedRuleManagedKeys {
			wafregional_GetRateBasedRuleManagedKeys(cfg, client)
			return
		}
		if _wafregionalGetRegexMatchSet {
			wafregional_GetRegexMatchSet(cfg, client)
			return
		}
		if _wafregionalGetRegexPatternSet {
			wafregional_GetRegexPatternSet(cfg, client)
			return
		}
		if _wafregionalGetRule {
			wafregional_GetRule(cfg, client)
			return
		}
		if _wafregionalGetRuleGroup {
			wafregional_GetRuleGroup(cfg, client)
			return
		}
		if _wafregionalGetSampledRequests {
			wafregional_GetSampledRequests(cfg, client)
			return
		}
		if _wafregionalGetSizeConstraintSet {
			wafregional_GetSizeConstraintSet(cfg, client)
			return
		}
		if _wafregionalGetSqlInjectionMatchSet {
			wafregional_GetSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafregionalGetWebACL {
			wafregional_GetWebACL(cfg, client)
			return
		}
		if _wafregionalGetWebACLForResource {
			wafregional_GetWebACLForResource(cfg, client)
			return
		}
		if _wafregionalGetXssMatchSet {
			wafregional_GetXssMatchSet(cfg, client)
			return
		}
		if _wafregionalListActivatedRulesInRuleGroup {
			wafregional_ListActivatedRulesInRuleGroup(cfg, client)
			return
		}
		if _wafregionalListByteMatchSets {
			wafregional_ListByteMatchSets(cfg, client)
			return
		}
		if _wafregionalListGeoMatchSets {
			wafregional_ListGeoMatchSets(cfg, client)
			return
		}
		if _wafregionalListIPSets {
			wafregional_ListIPSets(cfg, client)
			return
		}
		if _wafregionalListLoggingConfigurations {
			wafregional_ListLoggingConfigurations(cfg, client)
			return
		}
		if _wafregionalListRateBasedRules {
			wafregional_ListRateBasedRules(cfg, client)
			return
		}
		if _wafregionalListRegexMatchSets {
			wafregional_ListRegexMatchSets(cfg, client)
			return
		}
		if _wafregionalListRegexPatternSets {
			wafregional_ListRegexPatternSets(cfg, client)
			return
		}
		if _wafregionalListResourcesForWebACL {
			wafregional_ListResourcesForWebACL(cfg, client)
			return
		}
		if _wafregionalListRuleGroups {
			wafregional_ListRuleGroups(cfg, client)
			return
		}
		if _wafregionalListRules {
			wafregional_ListRules(cfg, client)
			return
		}
		if _wafregionalListSizeConstraintSets {
			wafregional_ListSizeConstraintSets(cfg, client)
			return
		}
		if _wafregionalListSqlInjectionMatchSets {
			wafregional_ListSqlInjectionMatchSets(cfg, client)
			return
		}
		if _wafregionalListSubscribedRuleGroups {
			wafregional_ListSubscribedRuleGroups(cfg, client)
			return
		}
		if _wafregionalListTagsForResource {
			wafregional_ListTagsForResource(cfg, client)
			return
		}
		if _wafregionalListWebACLs {
			wafregional_ListWebACLs(cfg, client)
			return
		}
		if _wafregionalListXssMatchSets {
			wafregional_ListXssMatchSets(cfg, client)
			return
		}
		if _wafregionalPutLoggingConfiguration {
			wafregional_PutLoggingConfiguration(cfg, client)
			return
		}
		if _wafregionalPutPermissionPolicy {
			wafregional_PutPermissionPolicy(cfg, client)
			return
		}
		if _wafregionalTagResource {
			wafregional_TagResource(cfg, client)
			return
		}
		if _wafregionalUntagResource {
			wafregional_UntagResource(cfg, client)
			return
		}
		if _wafregionalUpdateByteMatchSet {
			wafregional_UpdateByteMatchSet(cfg, client)
			return
		}
		if _wafregionalUpdateGeoMatchSet {
			wafregional_UpdateGeoMatchSet(cfg, client)
			return
		}
		if _wafregionalUpdateIPSet {
			wafregional_UpdateIPSet(cfg, client)
			return
		}
		if _wafregionalUpdateRateBasedRule {
			wafregional_UpdateRateBasedRule(cfg, client)
			return
		}
		if _wafregionalUpdateRegexMatchSet {
			wafregional_UpdateRegexMatchSet(cfg, client)
			return
		}
		if _wafregionalUpdateRegexPatternSet {
			wafregional_UpdateRegexPatternSet(cfg, client)
			return
		}
		if _wafregionalUpdateRule {
			wafregional_UpdateRule(cfg, client)
			return
		}
		if _wafregionalUpdateRuleGroup {
			wafregional_UpdateRuleGroup(cfg, client)
			return
		}
		if _wafregionalUpdateSizeConstraintSet {
			wafregional_UpdateSizeConstraintSet(cfg, client)
			return
		}
		if _wafregionalUpdateSqlInjectionMatchSet {
			wafregional_UpdateSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafregionalUpdateWebACL {
			wafregional_UpdateWebACL(cfg, client)
			return
		}
		if _wafregionalUpdateXssMatchSet {
			wafregional_UpdateXssMatchSet(cfg, client)
			return
		}

	},
}

var (
	_wafregionalAssociateWebACL               bool
	_wafregionalCreateByteMatchSet            bool
	_wafregionalCreateGeoMatchSet             bool
	_wafregionalCreateIPSet                   bool
	_wafregionalCreateRateBasedRule           bool
	_wafregionalCreateRegexMatchSet           bool
	_wafregionalCreateRegexPatternSet         bool
	_wafregionalCreateRule                    bool
	_wafregionalCreateRuleGroup               bool
	_wafregionalCreateSizeConstraintSet       bool
	_wafregionalCreateSqlInjectionMatchSet    bool
	_wafregionalCreateWebACL                  bool
	_wafregionalCreateWebACLMigrationStack    bool
	_wafregionalCreateXssMatchSet             bool
	_wafregionalDeleteByteMatchSet            bool
	_wafregionalDeleteGeoMatchSet             bool
	_wafregionalDeleteIPSet                   bool
	_wafregionalDeleteLoggingConfiguration    bool
	_wafregionalDeletePermissionPolicy        bool
	_wafregionalDeleteRateBasedRule           bool
	_wafregionalDeleteRegexMatchSet           bool
	_wafregionalDeleteRegexPatternSet         bool
	_wafregionalDeleteRule                    bool
	_wafregionalDeleteRuleGroup               bool
	_wafregionalDeleteSizeConstraintSet       bool
	_wafregionalDeleteSqlInjectionMatchSet    bool
	_wafregionalDeleteWebACL                  bool
	_wafregionalDeleteXssMatchSet             bool
	_wafregionalDisassociateWebACL            bool
	_wafregionalGetByteMatchSet               bool
	_wafregionalGetChangeToken                bool
	_wafregionalGetChangeTokenStatus          bool
	_wafregionalGetGeoMatchSet                bool
	_wafregionalGetIPSet                      bool
	_wafregionalGetLoggingConfiguration       bool
	_wafregionalGetPermissionPolicy           bool
	_wafregionalGetRateBasedRule              bool
	_wafregionalGetRateBasedRuleManagedKeys   bool
	_wafregionalGetRegexMatchSet              bool
	_wafregionalGetRegexPatternSet            bool
	_wafregionalGetRule                       bool
	_wafregionalGetRuleGroup                  bool
	_wafregionalGetSampledRequests            bool
	_wafregionalGetSizeConstraintSet          bool
	_wafregionalGetSqlInjectionMatchSet       bool
	_wafregionalGetWebACL                     bool
	_wafregionalGetWebACLForResource          bool
	_wafregionalGetXssMatchSet                bool
	_wafregionalListActivatedRulesInRuleGroup bool
	_wafregionalListByteMatchSets             bool
	_wafregionalListGeoMatchSets              bool
	_wafregionalListIPSets                    bool
	_wafregionalListLoggingConfigurations     bool
	_wafregionalListRateBasedRules            bool
	_wafregionalListRegexMatchSets            bool
	_wafregionalListRegexPatternSets          bool
	_wafregionalListResourcesForWebACL        bool
	_wafregionalListRuleGroups                bool
	_wafregionalListRules                     bool
	_wafregionalListSizeConstraintSets        bool
	_wafregionalListSqlInjectionMatchSets     bool
	_wafregionalListSubscribedRuleGroups      bool
	_wafregionalListTagsForResource           bool
	_wafregionalListWebACLs                   bool
	_wafregionalListXssMatchSets              bool
	_wafregionalPutLoggingConfiguration       bool
	_wafregionalPutPermissionPolicy           bool
	_wafregionalTagResource                   bool
	_wafregionalUntagResource                 bool
	_wafregionalUpdateByteMatchSet            bool
	_wafregionalUpdateGeoMatchSet             bool
	_wafregionalUpdateIPSet                   bool
	_wafregionalUpdateRateBasedRule           bool
	_wafregionalUpdateRegexMatchSet           bool
	_wafregionalUpdateRegexPatternSet         bool
	_wafregionalUpdateRule                    bool
	_wafregionalUpdateRuleGroup               bool
	_wafregionalUpdateSizeConstraintSet       bool
	_wafregionalUpdateSqlInjectionMatchSet    bool
	_wafregionalUpdateWebACL                  bool
	_wafregionalUpdateXssMatchSet             bool

	_wafregionalByteMatchSetId         string
	_wafregionalChangeToken            string
	_wafregionalDefaultAction          string
	_wafregionalGeoMatchSetId          string
	_wafregionalIgnoreUnsupportedType  string
	_wafregionalIPSetId                string
	_wafregionalLimit                  string
	_wafregionalLoggingConfiguration   string
	_wafregionalMaxItems               string
	_wafregionalMetricName             string
	_wafregionalName                   string
	_wafregionalNextMarker             string
	_wafregionalPolicy                 string
	_wafregionalRateKey                string
	_wafregionalRateLimit              string
	_wafregionalRegexMatchSetId        string
	_wafregionalRegexPatternSetId      string
	_wafregionalResourceARN            string
	_wafregionalResourceType           string
	_wafregionalRuleGroupId            string
	_wafregionalRuleId                 string
	_wafregionalS3BucketName           string
	_wafregionalSizeConstraintSetId    string
	_wafregionalSqlInjectionMatchSetId string
	_wafregionalTagKeys                []string
	_wafregionalTags                   string
	_wafregionalTimeWindow             string
	_wafregionalUpdates                string
	_wafregionalWebAclId               string
	_wafregionalWebACLId               string
	_wafregionalXssMatchSetId          string
)

// This is AWS WAF Classic Regional documentation. For more information, see [AWS WAF Classic] in
// the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Associates a web ACL with a resource, either an application load balancer or
// Amazon API Gateway stage.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_AssociateWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.AssociateWebACLInput{
		// ResourceArn: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}
	if len(_wafregionalWebACLId) > 0 {
		input.WebACLId = aws.String(_wafregionalWebACLId)
	}

	if resp, err := client.AssociateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a ByteMatchSet . You then use UpdateByteMatchSet to identify the part of a web request
// that you want AWS WAF to inspect, such as the values of the User-Agent header
// or the query string. For example, you can create a ByteMatchSet that matches
// any requests with User-Agent headers that contain the string BadBot . You can
// then configure AWS WAF to reject those requests.
//
// To create and configure a ByteMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateByteMatchSet request.
//
// - Submit a CreateByteMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateByteMatchSet request.
//
// - Submit an UpdateByteMatchSetrequest to specify the part of the request that you want AWS WAF
// to inspect (for example, the header or the URI) and the value that you want AWS
// WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateByteMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateByteMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateByteMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates an GeoMatchSet, which you use to specify which web requests you want to allow or
// block based on the country that the requests originate from. For example, if
// you're receiving a lot of requests from one or more countries and you want to
// block the requests, you can create an GeoMatchSet that contains those countries
// and then configure AWS WAF to block the requests.
//
// To create and configure a GeoMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateGeoMatchSet request.
//
// - Submit a CreateGeoMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateGeoMatchSetrequest.
//
// - Submit an UpdateGeoMatchSetSet request to specify the countries that you
// want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateGeoMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateGeoMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateGeoMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates an IPSet, which you use to specify which web requests that you want to allow
// or block based on the IP addresses that the requests originate from. For
// example, if you're receiving a lot of requests from one or more individual IP
// addresses or one or more ranges of IP addresses and you want to block the
// requests, you can create an IPSet that contains those IP addresses and then
// configure AWS WAF to block the requests.
//
// To create and configure an IPSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateIPSet request.
//
// - Submit a CreateIPSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateIPSetrequest.
//
// - Submit an UpdateIPSet request to specify the IP addresses that you want AWS
// WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateIPSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateIPSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a RateBasedRule. The RateBasedRule contains a RateLimit , which specifies the maximum
// number of requests that AWS WAF allows from a specified IP address in a
// five-minute period. The RateBasedRule also contains the IPSet objects,
// ByteMatchSet objects, and other predicates that identify the requests that you
// want to count or block if these requests exceed the RateLimit .
//
// If you add more than one predicate to a RateBasedRule , a request not only must
// exceed the RateLimit , but it also must match all the conditions to be counted
// or blocked. For example, suppose you add the following to a RateBasedRule :
//
// - An IPSet that matches the IP address 192.0.2.44/32
//
// - A ByteMatchSet that matches BadBot in the User-Agent header
//
// Further, you specify a RateLimit of 1,000.
//
// You then add the RateBasedRule to a WebACL and specify that you want to block
// requests that meet the conditions in the rule. For a request to be blocked, it
// must come from the IP address 192.0.2.44 and the User-Agent header in the
// request must contain the value BadBot . Further, requests that match these two
// conditions must be received at a rate of more than 1,000 requests every five
// minutes. If both conditions are met and the rate is exceeded, AWS WAF blocks the
// requests. If the rate drops below 1,000 for a five-minute period, AWS WAF no
// longer blocks the requests.
//
// As a second example, suppose you want to limit requests to a particular page on
// your site. To do this, you could add the following to a RateBasedRule :
//
// - A ByteMatchSet with FieldToMatch of URI
//
// - A PositionalConstraint of STARTS_WITH
//
// - A TargetString of login
//
// Further, you specify a RateLimit of 1,000.
//
// By adding this RateBasedRule to a WebACL , you could limit requests to your
// login page without affecting the rest of your site.
//
// To create and configure a RateBasedRule , perform the following steps:
//
// - Create and update the predicates that you want to include in the rule. For
// more information, see CreateByteMatchSet, CreateIPSet, and CreateSqlInjectionMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateRule request.
//
// - Submit a CreateRateBasedRule request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRulerequest.
//
// - Submit an UpdateRateBasedRule request to specify the predicates that you
// want to include in the rule.
//
// - Create and update a WebACL that contains the RateBasedRule . For more
// information, see CreateWebACL.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateRateBasedRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateRateBasedRuleInput{
		// ChangeToken: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
		// RateKey: types.RateKey, // Required
		// RateLimit: *int64, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalMetricName) > 0 {
		input.MetricName = aws.String(_wafregionalMetricName)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}
	if len(_wafregionalRateKey) > 0 {
		if err := assignInputField(input, "RateKey", _wafregionalRateKey); err != nil {
			log.Errorf("invalid --rate-key: %s", err.Error())
			return
		}
	}
	if len(_wafregionalRateLimit) > 0 {
		if err := assignInputField(input, "RateLimit", _wafregionalRateLimit); err != nil {
			log.Errorf("invalid --rate-limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalTags) > 0 {
		if err := assignInputField(input, "Tags", _wafregionalTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRateBasedRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a RegexMatchSet. You then use UpdateRegexMatchSet to identify the part of a web request that you want
// AWS WAF to inspect, such as the values of the User-Agent header or the query
// string. For example, you can create a RegexMatchSet that contains a
// RegexMatchTuple that looks for any requests with User-Agent headers that match
// a RegexPatternSet with pattern B[a(at)]dB[o0]t . You can then configure AWS WAF to
// reject those requests.
//
// To create and configure a RegexMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateRegexMatchSet request.
//
// - Submit a CreateRegexMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRegexMatchSet request.
//
// - Submit an UpdateRegexMatchSetrequest to specify the part of the request that you want AWS WAF
// to inspect (for example, the header or the URI) and the value, using a
// RegexPatternSet , that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateRegexMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateRegexMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateRegexMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a RegexPatternSet . You then use UpdateRegexPatternSet to specify the regular expression
// (regex) pattern that you want AWS WAF to search for, such as B[a(at)]dB[o0]t . You
// can then configure AWS WAF to reject those requests.
//
// To create and configure a RegexPatternSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateRegexPatternSet request.
//
// - Submit a CreateRegexPatternSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRegexPatternSet request.
//
// - Submit an UpdateRegexPatternSetrequest to specify the string that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateRegexPatternSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateRegexPatternSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a Rule , which contains the IPSet objects, ByteMatchSet objects, and
// other predicates that identify the requests that you want to block. If you add
// more than one predicate to a Rule , a request must match all of the
// specifications to be allowed or blocked. For example, suppose that you add the
// following to a Rule :
//
// - An IPSet that matches the IP address 192.0.2.44/32
//
// - A ByteMatchSet that matches BadBot in the User-Agent header
//
// You then add the Rule to a WebACL and specify that you want to blocks requests
// that satisfy the Rule . For a request to be blocked, it must come from the IP
// address 192.0.2.44 and the User-Agent header in the request must contain the
// value BadBot .
//
// To create and configure a Rule , perform the following steps:
//
// - Create and update the predicates that you want to include in the Rule . For
// more information, see CreateByteMatchSet, CreateIPSet, and CreateSqlInjectionMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateRule request.
//
// - Submit a CreateRule request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRulerequest.
//
// - Submit an UpdateRule request to specify the predicates that you want to
// include in the Rule .
//
// - Create and update a WebACL that contains the Rule . For more information,
// see CreateWebACL.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateRuleInput{
		// ChangeToken: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalMetricName) > 0 {
		input.MetricName = aws.String(_wafregionalMetricName)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}
	if len(_wafregionalTags) > 0 {
		if err := assignInputField(input, "Tags", _wafregionalTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a RuleGroup . A rule group is a collection of predefined rules that you
// add to a web ACL. You use UpdateRuleGroupto add rules to the rule group.
//
// Rule groups are subject to the following limits:
//
// - Three rule groups per account. You can request an increase to this limit by
// contacting customer support.
//
// - One rule group per web ACL.
//
// - Ten rules per rule group.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateRuleGroup(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateRuleGroupInput{
		// ChangeToken: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalMetricName) > 0 {
		input.MetricName = aws.String(_wafregionalMetricName)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}
	if len(_wafregionalTags) > 0 {
		if err := assignInputField(input, "Tags", _wafregionalTags); err != nil {
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

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a SizeConstraintSet . You then use UpdateSizeConstraintSet to identify the part of a web
// request that you want AWS WAF to check for length, such as the length of the
// User-Agent header or the length of the query string. For example, you can create
// a SizeConstraintSet that matches any requests that have a query string that is
// longer than 100 bytes. You can then configure AWS WAF to reject those requests.
//
// To create and configure a SizeConstraintSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateSizeConstraintSet request.
//
// - Submit a CreateSizeConstraintSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateSizeConstraintSet request.
//
// - Submit an UpdateSizeConstraintSetrequest to specify the part of the request that you want AWS WAF
// to inspect (for example, the header or the URI) and the value that you want AWS
// WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateSizeConstraintSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateSizeConstraintSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateSizeConstraintSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a SqlInjectionMatchSet, which you use to allow, block, or count requests that contain
// snippets of SQL code in a specified part of web requests. AWS WAF searches for
// character sequences that are likely to be malicious strings.
//
// To create and configure a SqlInjectionMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateSqlInjectionMatchSet request.
//
// - Submit a CreateSqlInjectionMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateSqlInjectionMatchSetrequest.
//
// - Submit an UpdateSqlInjectionMatchSetrequest to specify the parts of web requests in which you want to
// allow, block, or count malicious SQL code.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateSqlInjectionMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateSqlInjectionMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateSqlInjectionMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates a WebACL , which contains the Rules that identify the CloudFront web
// requests that you want to allow, block, or count. AWS WAF evaluates Rules in
// order based on the value of Priority for each Rule .
//
// You also specify a default action, either ALLOW or BLOCK . If a web request
// doesn't match any of the Rules in a WebACL , AWS WAF responds to the request
// with the default action.
//
// To create and configure a WebACL , perform the following steps:
//
// - Create and update the ByteMatchSet objects and other predicates that you
// want to include in Rules . For more information, see CreateByteMatchSet, UpdateByteMatchSet, CreateIPSet, UpdateIPSet, CreateSqlInjectionMatchSet, and UpdateSqlInjectionMatchSet.
//
// - Create and update the Rules that you want to include in the WebACL . For
// more information, see CreateRuleand UpdateRule.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateWebACL request.
//
// - Submit a CreateWebACL request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateWebACLrequest.
//
// - Submit an UpdateWebACLrequest to specify the Rules that you want to include in the
// WebACL , to specify the default action, and to associate the WebACL with a
// CloudFront distribution.
//
// For more information about how to use the AWS WAF API, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateWebACLInput{
		// ChangeToken: *string, // Required
		// DefaultAction: *types.WafAction, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _wafregionalDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_wafregionalMetricName) > 0 {
		input.MetricName = aws.String(_wafregionalMetricName)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}
	if len(_wafregionalTags) > 0 {
		if err := assignInputField(input, "Tags", _wafregionalTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an AWS CloudFormation WAFV2 template for the specified web ACL in the
// specified Amazon S3 bucket. Then, in CloudFormation, you create a stack from the
// template, to create the web ACL and its resources in AWS WAFV2. Use this to
// migrate your AWS WAF Classic web ACL to the latest version of AWS WAF.
//
// This is part of a larger migration procedure for web ACLs from AWS WAF Classic
// to the latest version of AWS WAF. For the full procedure, including caveats and
// manual steps to complete the migration and switch over to the new web ACL, see [Migrating your AWS WAF Classic resources to AWS WAF]
// in the [AWS WAF Developer Guide].
//
// [Migrating your AWS WAF Classic resources to AWS WAF]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-migrating-from-classic.html
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
func wafregional_CreateWebACLMigrationStack(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateWebACLMigrationStackInput{
		// IgnoreUnsupportedType: *bool, // Required
		// S3BucketName: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafregionalIgnoreUnsupportedType) > 0 {
		if err := assignInputField(input, "IgnoreUnsupportedType", _wafregionalIgnoreUnsupportedType); err != nil {
			log.Errorf("invalid --ignore-unsupported-type: %s", err.Error())
			return
		}
	}
	if len(_wafregionalS3BucketName) > 0 {
		input.S3BucketName = aws.String(_wafregionalS3BucketName)
	}
	if len(_wafregionalWebACLId) > 0 {
		input.WebACLId = aws.String(_wafregionalWebACLId)
	}

	if resp, err := client.CreateWebACLMigrationStack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Creates an XssMatchSet, which you use to allow, block, or count requests that contain
// cross-site scripting attacks in the specified part of web requests. AWS WAF
// searches for character sequences that are likely to be malicious strings.
//
// To create and configure an XssMatchSet , perform the following steps:
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a CreateXssMatchSet request.
//
// - Submit a CreateXssMatchSet request.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateXssMatchSetrequest.
//
// - Submit an UpdateXssMatchSetrequest to specify the parts of web requests in which you want to
// allow, block, or count cross-site scripting attacks.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_CreateXssMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.CreateXssMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalName) > 0 {
		input.Name = aws.String(_wafregionalName)
	}

	if resp, err := client.CreateXssMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a ByteMatchSet. You can't delete a ByteMatchSet if it's still used in
// any Rules or if it still includes any ByteMatchTuple objects (any filters).
//
// If you just want to remove a ByteMatchSet from a Rule , use UpdateRule.
//
// To permanently delete a ByteMatchSet , perform the following steps:
//
// - Update the ByteMatchSet to remove filters, if any. For more information, see UpdateByteMatchSet
// .
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteByteMatchSet request.
//
// - Submit a DeleteByteMatchSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteByteMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteByteMatchSetInput{
		// ByteMatchSetId: *string, // Required
		// ChangeToken: *string, // Required
	}

	if len(_wafregionalByteMatchSetId) > 0 {
		input.ByteMatchSetId = aws.String(_wafregionalByteMatchSetId)
	}
	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}

	if resp, err := client.DeleteByteMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a GeoMatchSet. You can't delete a GeoMatchSet if it's still used in
// any Rules or if it still includes any countries.
//
// If you just want to remove a GeoMatchSet from a Rule , use UpdateRule.
//
// To permanently delete a GeoMatchSet from AWS WAF, perform the following steps:
//
// - Update the GeoMatchSet to remove any countries. For more information, see UpdateGeoMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteGeoMatchSet request.
//
// - Submit a DeleteGeoMatchSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteGeoMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteGeoMatchSetInput{
		// ChangeToken: *string, // Required
		// GeoMatchSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalGeoMatchSetId) > 0 {
		input.GeoMatchSetId = aws.String(_wafregionalGeoMatchSetId)
	}

	if resp, err := client.DeleteGeoMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes an IPSet. You can't delete an IPSet if it's still used in any
// Rules or if it still includes any IP addresses.
//
// If you just want to remove an IPSet from a Rule , use UpdateRule.
//
// To permanently delete an IPSet from AWS WAF, perform the following steps:
//
// - Update the IPSet to remove IP address ranges, if any. For more information,
// see UpdateIPSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteIPSet request.
//
// - Submit a DeleteIPSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteIPSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteIPSetInput{
		// ChangeToken: *string, // Required
		// IPSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalIPSetId) > 0 {
		input.IPSetId = aws.String(_wafregionalIPSetId)
	}

	if resp, err := client.DeleteIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes the LoggingConfiguration from the specified web ACL.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteLoggingConfiguration(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteLoggingConfigurationInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.DeleteLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes an IAM policy from the specified RuleGroup.
//
// The user making the request must be the owner of the RuleGroup.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeletePermissionPolicy(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeletePermissionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.DeletePermissionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a RateBasedRule. You can't delete a rule if it's still used in any WebACL
// objects or if it still includes any predicates, such as ByteMatchSet objects.
//
// If you just want to remove a rule from a WebACL , use UpdateWebACL.
//
// To permanently delete a RateBasedRule from AWS WAF, perform the following steps:
//
// - Update the RateBasedRule to remove predicates, if any. For more information,
// see UpdateRateBasedRule.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteRateBasedRule request.
//
// - Submit a DeleteRateBasedRule request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteRateBasedRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteRateBasedRuleInput{
		// ChangeToken: *string, // Required
		// RuleId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}

	if resp, err := client.DeleteRateBasedRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a RegexMatchSet. You can't delete a RegexMatchSet if it's still used in
// any Rules or if it still includes any RegexMatchTuples objects (any filters).
//
// If you just want to remove a RegexMatchSet from a Rule , use UpdateRule.
//
// To permanently delete a RegexMatchSet , perform the following steps:
//
// - Update the RegexMatchSet to remove filters, if any. For more information,
// see UpdateRegexMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteRegexMatchSet request.
//
// - Submit a DeleteRegexMatchSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteRegexMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteRegexMatchSetInput{
		// ChangeToken: *string, // Required
		// RegexMatchSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRegexMatchSetId) > 0 {
		input.RegexMatchSetId = aws.String(_wafregionalRegexMatchSetId)
	}

	if resp, err := client.DeleteRegexMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a RegexPatternSet. You can't delete a RegexPatternSet if it's still used
// in any RegexMatchSet or if the RegexPatternSet is not empty.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteRegexPatternSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteRegexPatternSetInput{
		// ChangeToken: *string, // Required
		// RegexPatternSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRegexPatternSetId) > 0 {
		input.RegexPatternSetId = aws.String(_wafregionalRegexPatternSetId)
	}

	if resp, err := client.DeleteRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a Rule. You can't delete a Rule if it's still used in any WebACL
// objects or if it still includes any predicates, such as ByteMatchSet objects.
//
// If you just want to remove a Rule from a WebACL , use UpdateWebACL.
//
// To permanently delete a Rule from AWS WAF, perform the following steps:
//
// - Update the Rule to remove predicates, if any. For more information, see UpdateRule.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteRule request.
//
// - Submit a DeleteRule request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteRuleInput{
		// ChangeToken: *string, // Required
		// RuleId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}

	if resp, err := client.DeleteRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a RuleGroup. You can't delete a RuleGroup if it's still used in any
// WebACL objects or if it still includes any rules.
//
// If you just want to remove a RuleGroup from a WebACL , use UpdateWebACL.
//
// To permanently delete a RuleGroup from AWS WAF, perform the following steps:
//
// - Update the RuleGroup to remove rules, if any. For more information, see UpdateRuleGroup.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteRuleGroup request.
//
// - Submit a DeleteRuleGroup request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteRuleGroup(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteRuleGroupInput{
		// ChangeToken: *string, // Required
		// RuleGroupId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafregionalRuleGroupId)
	}

	if resp, err := client.DeleteRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a SizeConstraintSet. You can't delete a SizeConstraintSet if it's still used
// in any Rules or if it still includes any SizeConstraint objects (any filters).
//
// If you just want to remove a SizeConstraintSet from a Rule , use UpdateRule.
//
// To permanently delete a SizeConstraintSet , perform the following steps:
//
// - Update the SizeConstraintSet to remove filters, if any. For more
// information, see UpdateSizeConstraintSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteSizeConstraintSet request.
//
// - Submit a DeleteSizeConstraintSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteSizeConstraintSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteSizeConstraintSetInput{
		// ChangeToken: *string, // Required
		// SizeConstraintSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalSizeConstraintSetId) > 0 {
		input.SizeConstraintSetId = aws.String(_wafregionalSizeConstraintSetId)
	}

	if resp, err := client.DeleteSizeConstraintSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a SqlInjectionMatchSet. You can't delete a SqlInjectionMatchSet if it's still
// used in any Rules or if it still contains any SqlInjectionMatchTuple objects.
//
// If you just want to remove a SqlInjectionMatchSet from a Rule , use UpdateRule.
//
// To permanently delete a SqlInjectionMatchSet from AWS WAF, perform the
// following steps:
//
// - Update the SqlInjectionMatchSet to remove filters, if any. For more
// information, see UpdateSqlInjectionMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteSqlInjectionMatchSet request.
//
// - Submit a DeleteSqlInjectionMatchSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteSqlInjectionMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteSqlInjectionMatchSetInput{
		// ChangeToken: *string, // Required
		// SqlInjectionMatchSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalSqlInjectionMatchSetId) > 0 {
		input.SqlInjectionMatchSetId = aws.String(_wafregionalSqlInjectionMatchSetId)
	}

	if resp, err := client.DeleteSqlInjectionMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes a WebACL. You can't delete a WebACL if it still contains any Rules
// .
//
// To delete a WebACL , perform the following steps:
//
// - Update the WebACL to remove Rules , if any. For more information, see UpdateWebACL.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteWebACL request.
//
// - Submit a DeleteWebACL request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteWebACLInput{
		// ChangeToken: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalWebACLId) > 0 {
		input.WebACLId = aws.String(_wafregionalWebACLId)
	}

	if resp, err := client.DeleteWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Permanently deletes an XssMatchSet. You can't delete an XssMatchSet if it's still used in
// any Rules or if it still contains any XssMatchTuple objects.
//
// If you just want to remove an XssMatchSet from a Rule , use UpdateRule.
//
// To permanently delete an XssMatchSet from AWS WAF, perform the following steps:
//
// - Update the XssMatchSet to remove filters, if any. For more information, see UpdateXssMatchSet
// .
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// a DeleteXssMatchSet request.
//
// - Submit a DeleteXssMatchSet request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DeleteXssMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DeleteXssMatchSetInput{
		// ChangeToken: *string, // Required
		// XssMatchSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalXssMatchSetId) > 0 {
		input.XssMatchSetId = aws.String(_wafregionalXssMatchSetId)
	}

	if resp, err := client.DeleteXssMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic Regional documentation. For more information, see [AWS WAF Classic] in
// the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Removes a web ACL from the specified resource, either an application load
// balancer or Amazon API Gateway stage.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_DisassociateWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.DisassociateWebACLInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.DisassociateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the ByteMatchSet specified by ByteMatchSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetByteMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetByteMatchSetInput{
		// ByteMatchSetId: *string, // Required
	}

	if len(_wafregionalByteMatchSetId) > 0 {
		input.ByteMatchSetId = aws.String(_wafregionalByteMatchSetId)
	}

	if resp, err := client.GetByteMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// When you want to create, update, or delete AWS WAF objects, get a change token
// and include the change token in the create, update, or delete request. Change
// tokens ensure that your application doesn't submit conflicting requests to AWS
// WAF.
//
// Each create, update, or delete request must use a unique change token. If your
// application submits a GetChangeToken request and then submits a second
// GetChangeToken request before submitting a create, update, or delete request,
// the second GetChangeToken request returns the same value as the first
// GetChangeToken request.
//
// When you use a change token in a create, update, or delete request, the status
// of the change token changes to PENDING , which indicates that AWS WAF is
// propagating the change to all AWS WAF servers. Use GetChangeTokenStatus to
// determine the status of your change token.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetChangeToken(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetChangeTokenInput{}

	if resp, err := client.GetChangeToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the status of a ChangeToken that you got by calling GetChangeToken. ChangeTokenStatus
// is one of the following values:
//
// - PROVISIONED : You requested the change token by calling GetChangeToken , but
// you haven't used it yet in a call to create, update, or delete an AWS WAF
// object.
//
// - PENDING : AWS WAF is propagating the create, update, or delete request to
// all AWS WAF servers.
//
// - INSYNC : Propagation is complete.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetChangeTokenStatus(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetChangeTokenStatusInput{
		// ChangeToken: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}

	if resp, err := client.GetChangeTokenStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the GeoMatchSet that is specified by GeoMatchSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetGeoMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetGeoMatchSetInput{
		// GeoMatchSetId: *string, // Required
	}

	if len(_wafregionalGeoMatchSetId) > 0 {
		input.GeoMatchSetId = aws.String(_wafregionalGeoMatchSetId)
	}

	if resp, err := client.GetGeoMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the IPSet that is specified by IPSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetIPSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetIPSetInput{
		// IPSetId: *string, // Required
	}

	if len(_wafregionalIPSetId) > 0 {
		input.IPSetId = aws.String(_wafregionalIPSetId)
	}

	if resp, err := client.GetIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the LoggingConfiguration for the specified web ACL.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetLoggingConfiguration(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetLoggingConfigurationInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.GetLoggingConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the IAM policy attached to the RuleGroup.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetPermissionPolicy(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetPermissionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.GetPermissionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the RateBasedRule that is specified by the RuleId that you included in the
// GetRateBasedRule request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetRateBasedRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetRateBasedRuleInput{
		// RuleId: *string, // Required
	}

	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}

	if resp, err := client.GetRateBasedRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of IP addresses currently being blocked by the RateBasedRule that is
// specified by the RuleId . The maximum number of managed keys that will be
// blocked is 10,000. If more than 10,000 addresses exceed the rate limit, the
// 10,000 addresses with the highest rates will be blocked.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetRateBasedRuleManagedKeys(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetRateBasedRuleManagedKeysInput{
		// RuleId: *string, // Required
	}

	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.GetRateBasedRuleManagedKeys(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the RegexMatchSet specified by RegexMatchSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetRegexMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetRegexMatchSetInput{
		// RegexMatchSetId: *string, // Required
	}

	if len(_wafregionalRegexMatchSetId) > 0 {
		input.RegexMatchSetId = aws.String(_wafregionalRegexMatchSetId)
	}

	if resp, err := client.GetRegexMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the RegexPatternSet specified by RegexPatternSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetRegexPatternSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetRegexPatternSetInput{
		// RegexPatternSetId: *string, // Required
	}

	if len(_wafregionalRegexPatternSetId) > 0 {
		input.RegexPatternSetId = aws.String(_wafregionalRegexPatternSetId)
	}

	if resp, err := client.GetRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the Rule that is specified by the RuleId that you included in the GetRule
// request.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetRuleInput{
		// RuleId: *string, // Required
	}

	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}

	if resp, err := client.GetRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the RuleGroup that is specified by the RuleGroupId that you included in the
// GetRuleGroup request.
//
// To view the rules in a rule group, use ListActivatedRulesInRuleGroup.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetRuleGroup(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetRuleGroupInput{
		// RuleGroupId: *string, // Required
	}

	if len(_wafregionalRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafregionalRuleGroupId)
	}

	if resp, err := client.GetRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Gets detailed information about a specified number of requests--a sample--that
// AWS WAF randomly selects from among the first 5,000 requests that your AWS
// resource received during a time range that you choose. You can specify a sample
// size of up to 500 requests, and you can specify any time range in the previous
// three hours.
//
// GetSampledRequests returns a time range, which is usually the time range that
// you specified. However, if your resource (such as a CloudFront distribution)
// received 5,000 requests before the specified time range elapsed,
// GetSampledRequests returns an updated time range. This new time range indicates
// the actual period during which AWS WAF selected the requests in the sample.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetSampledRequests(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetSampledRequestsInput{
		// MaxItems: *int64, // Required
		// RuleId: *string, // Required
		// TimeWindow: *types.TimeWindow, // Required
		// WebAclId: *string, // Required
	}

	if len(_wafregionalMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _wafregionalMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}
	if len(_wafregionalTimeWindow) > 0 {
		if err := assignInputField(input, "TimeWindow", _wafregionalTimeWindow); err != nil {
			log.Errorf("invalid --time-window: %s", err.Error())
			return
		}
	}
	if len(_wafregionalWebAclId) > 0 {
		input.WebAclId = aws.String(_wafregionalWebAclId)
	}

	if resp, err := client.GetSampledRequests(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the SizeConstraintSet specified by SizeConstraintSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetSizeConstraintSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetSizeConstraintSetInput{
		// SizeConstraintSetId: *string, // Required
	}

	if len(_wafregionalSizeConstraintSetId) > 0 {
		input.SizeConstraintSetId = aws.String(_wafregionalSizeConstraintSetId)
	}

	if resp, err := client.GetSizeConstraintSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the SqlInjectionMatchSet that is specified by SqlInjectionMatchSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetSqlInjectionMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetSqlInjectionMatchSetInput{
		// SqlInjectionMatchSetId: *string, // Required
	}

	if len(_wafregionalSqlInjectionMatchSetId) > 0 {
		input.SqlInjectionMatchSetId = aws.String(_wafregionalSqlInjectionMatchSetId)
	}

	if resp, err := client.GetSqlInjectionMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the WebACL that is specified by WebACLId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetWebACLInput{
		// WebACLId: *string, // Required
	}

	if len(_wafregionalWebACLId) > 0 {
		input.WebACLId = aws.String(_wafregionalWebACLId)
	}

	if resp, err := client.GetWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic Regional documentation. For more information, see [AWS WAF Classic] in
// the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the web ACL for the specified resource, either an application load
// balancer or Amazon API Gateway stage.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetWebACLForResource(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetWebACLForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.GetWebACLForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns the XssMatchSet that is specified by XssMatchSetId .
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_GetXssMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.GetXssMatchSetInput{
		// XssMatchSetId: *string, // Required
	}

	if len(_wafregionalXssMatchSetId) > 0 {
		input.XssMatchSetId = aws.String(_wafregionalXssMatchSetId)
	}

	if resp, err := client.GetXssMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of ActivatedRule objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListActivatedRulesInRuleGroup(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListActivatedRulesInRuleGroupInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}
	if len(_wafregionalRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafregionalRuleGroupId)
	}

	if resp, err := client.ListActivatedRulesInRuleGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of ByteMatchSetSummary objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListByteMatchSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListByteMatchSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListByteMatchSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of GeoMatchSetSummary objects in the response.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListGeoMatchSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListGeoMatchSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListGeoMatchSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of IPSetSummary objects in the response.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListIPSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListIPSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListIPSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of LoggingConfiguration objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListLoggingConfigurations(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListLoggingConfigurationsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListLoggingConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of RuleSummary objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListRateBasedRules(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListRateBasedRulesInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListRateBasedRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of RegexMatchSetSummary objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListRegexMatchSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListRegexMatchSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListRegexMatchSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of RegexPatternSetSummary objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListRegexPatternSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListRegexPatternSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListRegexPatternSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic Regional documentation. For more information, see [AWS WAF Classic] in
// the developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of resources associated with the specified web ACL.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListResourcesForWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListResourcesForWebACLInput{
		// WebACLId: *string, // Required
	}

	if len(_wafregionalWebACLId) > 0 {
		input.WebACLId = aws.String(_wafregionalWebACLId)
	}
	if len(_wafregionalResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _wafregionalResourceType); err != nil {
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

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of RuleGroup objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListRuleGroups(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListRuleGroupsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListRuleGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of RuleSummary objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListRules(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListRulesInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of SizeConstraintSetSummary objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListSizeConstraintSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListSizeConstraintSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListSizeConstraintSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of SqlInjectionMatchSet objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListSqlInjectionMatchSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListSqlInjectionMatchSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListSqlInjectionMatchSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of RuleGroup objects that you are subscribed to.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListSubscribedRuleGroups(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListSubscribedRuleGroupsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListSubscribedRuleGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Retrieves the tags associated with the specified AWS resource. Tags are
// key:value pairs that you can use to categorize and manage your resources, for
// purposes like billing. For example, you might set the tag key to "customer" and
// the value to the customer name or ID. You can specify one or more tags to add to
// each AWS resource, up to 50 tags for a resource.
//
// Tagging is only available through the API, SDKs, and CLI. You can't manage or
// view tags through the AWS WAF Classic console. You can tag the AWS resources
// that you manage through AWS WAF Classic: web ACLs, rule groups, and rules.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListTagsForResource(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafregionalResourceARN)
	}
	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of WebACLSummary objects in the response.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListWebACLs(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListWebACLsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListWebACLs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Returns an array of XssMatchSet objects.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_ListXssMatchSets(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.ListXssMatchSetsInput{}

	if len(_wafregionalLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafregionalLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalNextMarker) > 0 {
		input.NextMarker = aws.String(_wafregionalNextMarker)
	}

	if resp, err := client.ListXssMatchSets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Associates a LoggingConfiguration with a specified web ACL.
//
// You can access information about all traffic that AWS WAF inspects using the
// following steps:
//
// - Create an Amazon Kinesis Data Firehose.
//
// # Create the data firehose with a PUT source and in the region that you are
//
// operating. However, if you are capturing logs for Amazon CloudFront, always
// create the firehose in US East (N. Virginia).
//
// Do not create the data firehose using a Kinesis stream as your source.
//
// - Associate that firehose to your web ACL using a PutLoggingConfiguration
// request.
//
// When you successfully enable logging using a PutLoggingConfiguration request,
// AWS WAF will create a service linked role with the necessary permissions to
// write logs to the Amazon Kinesis Data Firehose. For more information, see [Logging Web ACL Traffic Information]in
// the AWS WAF Developer Guide.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [Logging Web ACL Traffic Information]: https://docs.aws.amazon.com/waf/latest/developerguide/logging.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_PutLoggingConfiguration(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.PutLoggingConfigurationInput{
		// LoggingConfiguration: *types.LoggingConfiguration, // Required
	}

	if len(_wafregionalLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _wafregionalLoggingConfiguration); err != nil {
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

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Attaches an IAM policy to the specified resource. The only supported use for
// this action is to share a RuleGroup across accounts.
//
// The PutPermissionPolicy is subject to the following restrictions:
//
// - You can attach only one policy with each PutPermissionPolicy request.
//
// - The policy must include an Effect , Action and Principal .
//
// - Effect must specify Allow .
//
// - The Action in the policy must be waf:UpdateWebACL ,
// waf-regional:UpdateWebACL , waf:GetRuleGroup and waf-regional:GetRuleGroup .
// Any extra or wildcard actions in the policy will be rejected.
//
// - The policy cannot include a Resource parameter.
//
// - The ARN in the request must be a valid WAF RuleGroup ARN and the RuleGroup
// must exist in the same region.
//
// - The user making the request must be the owner of the RuleGroup.
//
// - Your policy must be composed using IAM Policy version 2012-10-17.
//
// For more information, see [IAM Policies].
//
// An example of a valid policy parameter is shown in the Examples section below.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [IAM Policies]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_PutPermissionPolicy(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.PutPermissionPolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_wafregionalPolicy) > 0 {
		input.Policy = aws.String(_wafregionalPolicy)
	}
	if len(_wafregionalResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafregionalResourceARN)
	}

	if resp, err := client.PutPermissionPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Associates tags with the specified AWS resource. Tags are key:value pairs that
// you can use to categorize and manage your resources, for purposes like billing.
// For example, you might set the tag key to "customer" and the value to the
// customer name or ID. You can specify one or more tags to add to each AWS
// resource, up to 50 tags for a resource.
//
// Tagging is only available through the API, SDKs, and CLI. You can't manage or
// view tags through the AWS WAF Classic console. You can use this action to tag
// the AWS resources that you manage through AWS WAF Classic: web ACLs, rule
// groups, and rules.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_TagResource(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafregionalResourceARN)
	}
	if len(_wafregionalTags) > 0 {
		if err := assignInputField(input, "Tags", _wafregionalTags); err != nil {
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

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UntagResource(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_wafregionalResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafregionalResourceARN)
	}
	if len(_wafregionalTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _wafregionalTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ByteMatchTuple objects (filters) in a ByteMatchSet. For each ByteMatchTuple object,
// you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change a ByteMatchSetUpdate object, you delete the existing object and add a
// new one.
//
// - The part of a web request that you want AWS WAF to inspect, such as a query
// string or the value of the User-Agent header.
//
// - The bytes (typically a string that corresponds with ASCII characters) that
// you want AWS WAF to look for. For more information, including how you specify
// the values for the AWS WAF API and the AWS CLI or SDKs, see TargetString in
// the ByteMatchTupledata type.
//
// - Where to look, such as at the beginning or the end of a query string.
//
// - Whether to perform any conversions on the request, such as converting it to
// lowercase, before inspecting it for the specified string.
//
// For example, you can add a ByteMatchSetUpdate object that matches web requests
// in which User-Agent headers contain the string BadBot . You can then configure
// AWS WAF to block those requests.
//
// To create and configure a ByteMatchSet , perform the following steps:
//
// - Create a ByteMatchSet. For more information, see CreateByteMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateByteMatchSet request.
//
// - Submit an UpdateByteMatchSet request to specify the part of the request that
// you want AWS WAF to inspect (for example, the header or the URI) and the value
// that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateByteMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateByteMatchSetInput{
		// ByteMatchSetId: *string, // Required
		// ChangeToken: *string, // Required
		// Updates: []types.ByteMatchSetUpdate, // Required
	}

	if len(_wafregionalByteMatchSetId) > 0 {
		input.ByteMatchSetId = aws.String(_wafregionalByteMatchSetId)
	}
	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateByteMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes GeoMatchConstraint objects in an GeoMatchSet . For each GeoMatchConstraint
// object, you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change an GeoMatchConstraint object, you delete the existing object and add a
// new one.
//
// - The Type . The only valid value for Type is Country .
//
// - The Value , which is a two character code for the country to add to the
// GeoMatchConstraint object. Valid codes are listed in GeoMatchConstraint$Value.
//
// To create and configure an GeoMatchSet , perform the following steps:
//
// - Submit a CreateGeoMatchSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateGeoMatchSetrequest.
//
// - Submit an UpdateGeoMatchSet request to specify the country that you want AWS
// WAF to watch for.
//
// When you update an GeoMatchSet , you specify the country that you want to add
// and/or the country that you want to delete. If you want to change a country, you
// delete the existing country and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateGeoMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateGeoMatchSetInput{
		// ChangeToken: *string, // Required
		// GeoMatchSetId: *string, // Required
		// Updates: []types.GeoMatchSetUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalGeoMatchSetId) > 0 {
		input.GeoMatchSetId = aws.String(_wafregionalGeoMatchSetId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGeoMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes IPSetDescriptor objects in an IPSet . For each IPSetDescriptor object, you
// specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change an IPSetDescriptor object, you delete the existing object and add a new
// one.
//
// - The IP address version, IPv4 or IPv6 .
//
// - The IP address in CIDR notation, for example, 192.0.2.0/24 (for the range of
// IP addresses from 192.0.2.0 to 192.0.2.255 ) or 192.0.2.44/32 (for the
// individual IP address 192.0.2.44 ).
//
// AWS WAF supports IPv4 address ranges: /8 and any range between /16 through /32.
// AWS WAF supports IPv6 address ranges: /24, /32, /48, /56, /64, and /128. For
// more information about CIDR notation, see the Wikipedia entry [Classless Inter-Domain Routing].
//
// IPv6 addresses can be represented using any of the following formats:
//
// - 1111:0000:0000:0000:0000:0000:0000:0111/128
//
// - 1111:0:0:0:0:0:0:0111/128
//
// - 1111::0111/128
//
// - 1111::111/128
//
// You use an IPSet to specify which web requests you want to allow or block based
// on the IP addresses that the requests originated from. For example, if you're
// receiving a lot of requests from one or a small number of IP addresses and you
// want to block the requests, you can create an IPSet that specifies those IP
// addresses, and then configure AWS WAF to block the requests.
//
// To create and configure an IPSet , perform the following steps:
//
// - Submit a CreateIPSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateIPSetrequest.
//
// - Submit an UpdateIPSet request to specify the IP addresses that you want AWS
// WAF to watch for.
//
// When you update an IPSet , you specify the IP addresses that you want to add
// and/or the IP addresses that you want to delete. If you want to change an IP
// address, you delete the existing IP address and add the new one.
//
// You can insert a maximum of 1000 addresses in a single request.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [Classless Inter-Domain Routing]: https://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateIPSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateIPSetInput{
		// ChangeToken: *string, // Required
		// IPSetId: *string, // Required
		// Updates: []types.IPSetUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalIPSetId) > 0 {
		input.IPSetId = aws.String(_wafregionalIPSetId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIPSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes Predicate objects in a rule and updates the RateLimit in the rule.
//
// Each Predicate object identifies a predicate, such as a ByteMatchSet or an IPSet, that specifies
// the web requests that you want to block or count. The RateLimit specifies the
// number of requests every five minutes that triggers the rule.
//
// If you add more than one predicate to a RateBasedRule , a request must match all
// the predicates and exceed the RateLimit to be counted or blocked. For example,
// suppose you add the following to a RateBasedRule :
//
// - An IPSet that matches the IP address 192.0.2.44/32
//
// - A ByteMatchSet that matches BadBot in the User-Agent header
//
// Further, you specify a RateLimit of 1,000.
//
// You then add the RateBasedRule to a WebACL and specify that you want to block
// requests that satisfy the rule. For a request to be blocked, it must come from
// the IP address 192.0.2.44 and the User-Agent header in the request must contain
// the value BadBot . Further, requests that match these two conditions much be
// received at a rate of more than 1,000 every five minutes. If the rate drops
// below this limit, AWS WAF no longer blocks the requests.
//
// As a second example, suppose you want to limit requests to a particular page on
// your site. To do this, you could add the following to a RateBasedRule :
//
// - A ByteMatchSet with FieldToMatch of URI
//
// - A PositionalConstraint of STARTS_WITH
//
// - A TargetString of login
//
// Further, you specify a RateLimit of 1,000.
//
// By adding this RateBasedRule to a WebACL , you could limit requests to your
// login page without affecting the rest of your site.
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/waf-chapter.html
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateRateBasedRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateRateBasedRuleInput{
		// ChangeToken: *string, // Required
		// RateLimit: *int64, // Required
		// RuleId: *string, // Required
		// Updates: []types.RuleUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRateLimit) > 0 {
		if err := assignInputField(input, "RateLimit", _wafregionalRateLimit); err != nil {
			log.Errorf("invalid --rate-limit: %s", err.Error())
			return
		}
	}
	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRateBasedRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes RegexMatchTuple objects (filters) in a RegexMatchSet. For each RegexMatchSetUpdate
// object, you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change a RegexMatchSetUpdate object, you delete the existing object and add a
// new one.
//
// - The part of a web request that you want AWS WAF to inspectupdate, such as a
// query string or the value of the User-Agent header.
//
// - The identifier of the pattern (a regular expression) that you want AWS WAF
// to look for. For more information, see RegexPatternSet.
//
// - Whether to perform any conversions on the request, such as converting it to
// lowercase, before inspecting it for the specified string.
//
// For example, you can create a RegexPatternSet that matches any requests with
// User-Agent headers that contain the string B[a(at)]dB[o0]t . You can then configure
// AWS WAF to reject those requests.
//
// To create and configure a RegexMatchSet , perform the following steps:
//
// - Create a RegexMatchSet. For more information, see CreateRegexMatchSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateRegexMatchSet request.
//
// - Submit an UpdateRegexMatchSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and the
// identifier of the RegexPatternSet that contain the regular expression patters
// you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateRegexMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateRegexMatchSetInput{
		// ChangeToken: *string, // Required
		// RegexMatchSetId: *string, // Required
		// Updates: []types.RegexMatchSetUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRegexMatchSetId) > 0 {
		input.RegexMatchSetId = aws.String(_wafregionalRegexMatchSetId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRegexMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes RegexPatternString objects in a RegexPatternSet. For each RegexPatternString
// object, you specify the following values:
//
// - Whether to insert or delete the RegexPatternString .
//
// - The regular expression pattern that you want to insert or delete. For more
// information, see RegexPatternSet.
//
// For example, you can create a RegexPatternString such as B[a(at)]dB[o0]t . AWS WAF
// will match this RegexPatternString to:
//
// - BadBot
//
// - BadB0t
//
// - B(at)dBot
//
// - B(at)dB0t
//
// To create and configure a RegexPatternSet , perform the following steps:
//
// - Create a RegexPatternSet. For more information, see CreateRegexPatternSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateRegexPatternSet request.
//
// - Submit an UpdateRegexPatternSet request to specify the regular expression
// pattern that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateRegexPatternSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateRegexPatternSetInput{
		// ChangeToken: *string, // Required
		// RegexPatternSetId: *string, // Required
		// Updates: []types.RegexPatternSetUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRegexPatternSetId) > 0 {
		input.RegexPatternSetId = aws.String(_wafregionalRegexPatternSetId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRegexPatternSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes Predicate objects in a Rule . Each Predicate object identifies a
// predicate, such as a ByteMatchSetor an IPSet, that specifies the web requests that you want to
// allow, block, or count. If you add more than one predicate to a Rule , a request
// must match all of the specifications to be allowed, blocked, or counted. For
// example, suppose that you add the following to a Rule :
//
// - A ByteMatchSet that matches the value BadBot in the User-Agent header
//
// - An IPSet that matches the IP address 192.0.2.44
//
// You then add the Rule to a WebACL and specify that you want to block requests
// that satisfy the Rule . For a request to be blocked, the User-Agent header in
// the request must contain the value BadBot and the request must originate from
// the IP address 192.0.2.44.
//
// To create and configure a Rule , perform the following steps:
//
// - Create and update the predicates that you want to include in the Rule .
//
// - Create the Rule . See CreateRule.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRulerequest.
//
// - Submit an UpdateRule request to add predicates to the Rule .
//
// - Create and update a WebACL that contains the Rule . See CreateWebACL.
//
// If you want to replace one ByteMatchSet or IPSet with another, you delete the
// existing one and add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateRule(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateRuleInput{
		// ChangeToken: *string, // Required
		// RuleId: *string, // Required
		// Updates: []types.RuleUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRuleId) > 0 {
		input.RuleId = aws.String(_wafregionalRuleId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ActivatedRule objects in a RuleGroup .
//
// You can only insert REGULAR rules into a rule group.
//
// You can have a maximum of ten rules per rule group.
//
// To create and configure a RuleGroup , perform the following steps:
//
// - Create and update the Rules that you want to include in the RuleGroup . See CreateRule
// .
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateRuleGrouprequest.
//
// - Submit an UpdateRuleGroup request to add Rules to the RuleGroup .
//
// - Create and update a WebACL that contains the RuleGroup . See CreateWebACL.
//
// If you want to replace one Rule with another, you delete the existing one and
// add the new one.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateRuleGroup(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateRuleGroupInput{
		// ChangeToken: *string, // Required
		// RuleGroupId: *string, // Required
		// Updates: []types.RuleGroupUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafregionalRuleGroupId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
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

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes SizeConstraint objects (filters) in a SizeConstraintSet. For each SizeConstraint object,
// you specify the following values:
//
// - Whether to insert or delete the object from the array. If you want to
// change a SizeConstraintSetUpdate object, you delete the existing object and
// add a new one.
//
// - The part of a web request that you want AWS WAF to evaluate, such as the
// length of a query string or the length of the User-Agent header.
//
// - Whether to perform any transformations on the request, such as converting
// it to lowercase, before checking its length. Note that transformations of the
// request body are not supported because the AWS resource forwards only the first
// 8192 bytes of your request to AWS WAF.
//
// You can only specify a single type of TextTransformation.
//
// - A ComparisonOperator used for evaluating the selected part of the request
// against the specified Size , such as equals, greater than, less than, and so
// on.
//
// - The length, in bytes, that you want AWS WAF to watch for in selected part
// of the request. The length is computed after applying the transformation.
//
// For example, you can add a SizeConstraintSetUpdate object that matches web
// requests in which the length of the User-Agent header is greater than 100
// bytes. You can then configure AWS WAF to block those requests.
//
// To create and configure a SizeConstraintSet , perform the following steps:
//
// - Create a SizeConstraintSet. For more information, see CreateSizeConstraintSet.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateSizeConstraintSet request.
//
// - Submit an UpdateSizeConstraintSet request to specify the part of the request
// that you want AWS WAF to inspect (for example, the header or the URI) and the
// value that you want AWS WAF to watch for.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateSizeConstraintSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateSizeConstraintSetInput{
		// ChangeToken: *string, // Required
		// SizeConstraintSetId: *string, // Required
		// Updates: []types.SizeConstraintSetUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalSizeConstraintSetId) > 0 {
		input.SizeConstraintSetId = aws.String(_wafregionalSizeConstraintSetId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSizeConstraintSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes SqlInjectionMatchTuple objects (filters) in a SqlInjectionMatchSet. For each SqlInjectionMatchTuple
// object, you specify the following values:
//
// - Action : Whether to insert the object into or delete the object from the
// array. To change a SqlInjectionMatchTuple , you delete the existing object and
// add a new one.
//
// - FieldToMatch : The part of web requests that you want AWS WAF to inspect
// and, if you want AWS WAF to inspect a header or custom query parameter, the name
// of the header or parameter.
//
// - TextTransformation : Which text transformation, if any, to perform on the
// web request before inspecting the request for snippets of malicious SQL code.
//
// You can only specify a single type of TextTransformation.
//
// You use SqlInjectionMatchSet objects to specify which CloudFront requests that
// you want to allow, block, or count. For example, if you're receiving requests
// that contain snippets of SQL code in the query string and you want to block the
// requests, you can create a SqlInjectionMatchSet with the applicable settings,
// and then configure AWS WAF to block the requests.
//
// To create and configure a SqlInjectionMatchSet , perform the following steps:
//
// - Submit a CreateSqlInjectionMatchSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateIPSetrequest.
//
// - Submit an UpdateSqlInjectionMatchSet request to specify the parts of web
// requests that you want AWS WAF to inspect for snippets of SQL code.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateSqlInjectionMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateSqlInjectionMatchSetInput{
		// ChangeToken: *string, // Required
		// SqlInjectionMatchSetId: *string, // Required
		// Updates: []types.SqlInjectionMatchSetUpdate, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalSqlInjectionMatchSetId) > 0 {
		input.SqlInjectionMatchSetId = aws.String(_wafregionalSqlInjectionMatchSetId)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSqlInjectionMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes ActivatedRule objects in a WebACL . Each Rule identifies web requests
// that you want to allow, block, or count. When you update a WebACL , you specify
// the following values:
//
// - A default action for the WebACL , either ALLOW or BLOCK . AWS WAF performs
// the default action if a request doesn't match the criteria in any of the Rules
// in a WebACL .
//
// - The Rules that you want to add or delete. If you want to replace one Rule
// with another, you delete the existing Rule and add the new one.
//
// - For each Rule , whether you want AWS WAF to allow requests, block requests,
// or count requests that match the conditions in the Rule .
//
// - The order in which you want AWS WAF to evaluate the Rules in a WebACL . If
// you add more than one Rule to a WebACL , AWS WAF evaluates each request
// against the Rules in order based on the value of Priority . (The Rule that has
// the lowest value for Priority is evaluated first.) When a web request matches
// all the predicates (such as ByteMatchSets and IPSets ) in a Rule , AWS WAF
// immediately takes the corresponding action, allow or block, and doesn't evaluate
// the request against the remaining Rules in the WebACL , if any.
//
// To create and configure a WebACL , perform the following steps:
//
// - Create and update the predicates that you want to include in Rules . For
// more information, see CreateByteMatchSet, UpdateByteMatchSet, CreateIPSet, UpdateIPSet, CreateSqlInjectionMatchSet, and UpdateSqlInjectionMatchSet.
//
// - Create and update the Rules that you want to include in the WebACL . For
// more information, see CreateRuleand UpdateRule.
//
// - Create a WebACL . See CreateWebACL.
//
// - Use GetChangeToken to get the change token that you provide in the
// ChangeToken parameter of an UpdateWebACLrequest.
//
// - Submit an UpdateWebACL request to specify the Rules that you want to include
// in the WebACL , to specify the default action, and to associate the WebACL
// with a CloudFront distribution.
//
// The ActivatedRule can be a rule group. If you specify a rule group as your
//
// ActivatedRule , you can exclude specific rules from that rule group.
//
// # If you already have a rule group associated with a web ACL and want to submit
//
// an UpdateWebACL request to exclude certain rules from that rule group, you
// must first remove the rule group from the web ACL, the re-insert it again,
// specifying the excluded rules. For details, see ActivatedRule$ExcludedRules.
//
// Be aware that if you try to add a RATE_BASED rule to a web ACL without setting
// the rule type when first creating the rule, the UpdateWebACLrequest will fail because the
// request tries to add a REGULAR rule (the default rule type) with the specified
// ID, which does not exist.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateWebACL(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateWebACLInput{
		// ChangeToken: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalWebACLId) > 0 {
		input.WebACLId = aws.String(_wafregionalWebACLId)
	}
	if len(_wafregionalDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _wafregionalDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateWebACL(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This is AWS WAF Classic documentation. For more information, see [AWS WAF Classic] in the
// developer guide.
//
// For the latest version of AWS WAF, use the AWS WAFV2 API and see the [AWS WAF Developer Guide]. With the
// latest version, AWS WAF has a single set of endpoints for regional and global
// use.
//
// Inserts or deletes XssMatchTuple objects (filters) in an XssMatchSet. For each XssMatchTuple object,
// you specify the following values:
//
// - Action : Whether to insert the object into or delete the object from the
// array. To change an XssMatchTuple , you delete the existing object and add a
// new one.
//
// - FieldToMatch : The part of web requests that you want AWS WAF to inspect
// and, if you want AWS WAF to inspect a header or custom query parameter, the name
// of the header or parameter.
//
// - TextTransformation : Which text transformation, if any, to perform on the
// web request before inspecting the request for cross-site scripting attacks.
//
// You can only specify a single type of TextTransformation.
//
// You use XssMatchSet objects to specify which CloudFront requests that you want
// to allow, block, or count. For example, if you're receiving requests that
// contain cross-site scripting attacks in the request body and you want to block
// the requests, you can create an XssMatchSet with the applicable settings, and
// then configure AWS WAF to block the requests.
//
// To create and configure an XssMatchSet , perform the following steps:
//
// - Submit a CreateXssMatchSetrequest.
//
// - Use GetChangeTokento get the change token that you provide in the ChangeToken parameter of
// an UpdateIPSetrequest.
//
// - Submit an UpdateXssMatchSet request to specify the parts of web requests
// that you want AWS WAF to inspect for cross-site scripting attacks.
//
// For more information about how to use the AWS WAF API to allow or block HTTP
// requests, see the [AWS WAF Developer Guide].
//
// [AWS WAF Developer Guide]: https://docs.aws.amazon.com/waf/latest/developerguide/
// [AWS WAF Classic]: https://docs.aws.amazon.com/waf/latest/developerguide/classic-waf-chapter.html
func wafregional_UpdateXssMatchSet(cfg aws.Config, client *wafregional.Client) {
	input := &wafregional.UpdateXssMatchSetInput{
		// ChangeToken: *string, // Required
		// Updates: []types.XssMatchSetUpdate, // Required
		// XssMatchSetId: *string, // Required
	}

	if len(_wafregionalChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafregionalChangeToken)
	}
	if len(_wafregionalUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafregionalUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}
	if len(_wafregionalXssMatchSetId) > 0 {
		input.XssMatchSetId = aws.String(_wafregionalXssMatchSetId)
	}

	if resp, err := client.UpdateXssMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_wafregionalCmd)
	_wafregionalCmd.Flags().SortFlags = false

	_wafregionalCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_wafregionalCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_wafregionalCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_wafregionalCmd.Flags().StringVarP(&_wafregionalByteMatchSetId, "byte-match-set-id", "", "", "Byte Match Set ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalChangeToken, "change-token", "", "", "Change Token")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalDefaultAction, "default-action", "", "", "Default Action")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalGeoMatchSetId, "geo-match-set-id", "", "", "Geo Match Set ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalIgnoreUnsupportedType, "ignore-unsupported-type", "", "", "Ignore Unsupported Type")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalIPSetId, "ip-set-id", "", "", "IP Set ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalLimit, "limit", "", "", "Limit")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalMaxItems, "max-items", "", "", "Max Items")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalMetricName, "metric-name", "", "", "Metric Name")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalName, "name", "", "", "Name")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalNextMarker, "next-marker", "", "", "Next Marker")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalPolicy, "policy", "", "", "Policy")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalRateKey, "rate-key", "", "", "Rate Key")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalRateLimit, "rate-limit", "", "", "Rate Limit")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalRegexMatchSetId, "regex-match-set-id", "", "", "Regex Match Set ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalRegexPatternSetId, "regex-pattern-set-id", "", "", "Regex Pattern Set ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalResourceARN, "resource-arn", "", "", "Resource ARN")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalResourceType, "resource-type", "", "", "Resource Type")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalRuleGroupId, "rule-group-id", "", "", "Rule Group ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalRuleId, "rule-id", "", "", "Rule ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalSizeConstraintSetId, "size-constraint-set-id", "", "", "Size Constraint Set ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalSqlInjectionMatchSetId, "sql-injection-match-set-id", "", "", "Sql Injection Match Set ID")
	_wafregionalCmd.Flags().StringSliceVarP(&_wafregionalTagKeys, "tag-keys", "", nil, "Tag Keys")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalTags, "tags", "", "", "Tags")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalTimeWindow, "time-window", "", "", "Time Window")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalUpdates, "updates", "", "", "Updates")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalWebAclId, "web-acl-id", "", "", "Web ACL ID")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalWebACLId, "web-aclid", "", "", "Web Aclid")
	_wafregionalCmd.Flags().StringVarP(&_wafregionalXssMatchSetId, "xss-match-set-id", "", "", "Xss Match Set ID")

	_wafregionalCmd.Flags().BoolVarP(&_wafregionalAssociateWebACL, "associate-web-acl", "", false, "Associate Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateByteMatchSet, "create-byte-match-set", "", false, "Create Byte Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateGeoMatchSet, "create-geo-match-set", "", false, "Create Geo Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateIPSet, "create-ip-set", "", false, "Create IP Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateRateBasedRule, "create-rate-based-rule", "", false, "Create Rate Based Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateRegexMatchSet, "create-regex-match-set", "", false, "Create Regex Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateRegexPatternSet, "create-regex-pattern-set", "", false, "Create Regex Pattern Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateRule, "create-rule", "", false, "Create Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateRuleGroup, "create-rule-group", "", false, "Create Rule Group")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateSizeConstraintSet, "create-size-constraint-set", "", false, "Create Size Constraint Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateSqlInjectionMatchSet, "create-sql-injection-match-set", "", false, "Create Sql Injection Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateWebACL, "create-web-acl", "", false, "Create Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateWebACLMigrationStack, "create-web-acl-migration-stack", "", false, "Create Web ACL Migration Stack")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalCreateXssMatchSet, "create-xss-match-set", "", false, "Create Xss Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteByteMatchSet, "delete-byte-match-set", "", false, "Delete Byte Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteGeoMatchSet, "delete-geo-match-set", "", false, "Delete Geo Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteIPSet, "delete-ip-set", "", false, "Delete IP Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteLoggingConfiguration, "delete-logging-configuration", "", false, "Delete Logging Configuration")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeletePermissionPolicy, "delete-permission-policy", "", false, "Delete Permission Policy")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteRateBasedRule, "delete-rate-based-rule", "", false, "Delete Rate Based Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteRegexMatchSet, "delete-regex-match-set", "", false, "Delete Regex Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteRegexPatternSet, "delete-regex-pattern-set", "", false, "Delete Regex Pattern Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteRule, "delete-rule", "", false, "Delete Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteRuleGroup, "delete-rule-group", "", false, "Delete Rule Group")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteSizeConstraintSet, "delete-size-constraint-set", "", false, "Delete Size Constraint Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteSqlInjectionMatchSet, "delete-sql-injection-match-set", "", false, "Delete Sql Injection Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteWebACL, "delete-web-acl", "", false, "Delete Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDeleteXssMatchSet, "delete-xss-match-set", "", false, "Delete Xss Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalDisassociateWebACL, "disassociate-web-acl", "", false, "Disassociate Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetByteMatchSet, "get-byte-match-set", "", false, "Get Byte Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetChangeToken, "get-change-token", "", false, "Get Change Token")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetChangeTokenStatus, "get-change-token-status", "", false, "Get Change Token Status")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetGeoMatchSet, "get-geo-match-set", "", false, "Get Geo Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetIPSet, "get-ip-set", "", false, "Get IP Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetLoggingConfiguration, "get-logging-configuration", "", false, "Get Logging Configuration")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetPermissionPolicy, "get-permission-policy", "", false, "Get Permission Policy")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetRateBasedRule, "get-rate-based-rule", "", false, "Get Rate Based Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetRateBasedRuleManagedKeys, "get-rate-based-rule-managed-keys", "", false, "Get Rate Based Rule Managed Keys")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetRegexMatchSet, "get-regex-match-set", "", false, "Get Regex Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetRegexPatternSet, "get-regex-pattern-set", "", false, "Get Regex Pattern Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetRule, "get-rule", "", false, "Get Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetRuleGroup, "get-rule-group", "", false, "Get Rule Group")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetSampledRequests, "get-sampled-requests", "", false, "Get Sampled Requests")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetSizeConstraintSet, "get-size-constraint-set", "", false, "Get Size Constraint Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetSqlInjectionMatchSet, "get-sql-injection-match-set", "", false, "Get Sql Injection Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetWebACL, "get-web-acl", "", false, "Get Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetWebACLForResource, "get-web-acl-for-resource", "", false, "Get Web ACL For Resource")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalGetXssMatchSet, "get-xss-match-set", "", false, "Get Xss Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListActivatedRulesInRuleGroup, "list-activated-rules-in-rule-group", "", false, "List Activated Rules In Rule Group")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListByteMatchSets, "list-byte-match-sets", "", false, "List Byte Match Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListGeoMatchSets, "list-geo-match-sets", "", false, "List Geo Match Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListIPSets, "list-ip-sets", "", false, "List IP Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListLoggingConfigurations, "list-logging-configurations", "", false, "List Logging Configurations")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListRateBasedRules, "list-rate-based-rules", "", false, "List Rate Based Rules")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListRegexMatchSets, "list-regex-match-sets", "", false, "List Regex Match Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListRegexPatternSets, "list-regex-pattern-sets", "", false, "List Regex Pattern Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListResourcesForWebACL, "list-resources-for-web-acl", "", false, "List Resources For Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListRuleGroups, "list-rule-groups", "", false, "List Rule Groups")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListRules, "list-rules", "", false, "List Rules")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListSizeConstraintSets, "list-size-constraint-sets", "", false, "List Size Constraint Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListSqlInjectionMatchSets, "list-sql-injection-match-sets", "", false, "List Sql Injection Match Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListSubscribedRuleGroups, "list-subscribed-rule-groups", "", false, "List Subscribed Rule Groups")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListWebACLs, "list-web-acls", "", false, "List Web Acls")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalListXssMatchSets, "list-xss-match-sets", "", false, "List Xss Match Sets")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalPutLoggingConfiguration, "put-logging-configuration", "", false, "Put Logging Configuration")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalPutPermissionPolicy, "put-permission-policy", "", false, "Put Permission Policy")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalTagResource, "tag-resource", "", false, "Tag Resource")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUntagResource, "untag-resource", "", false, "Untag Resource")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateByteMatchSet, "update-byte-match-set", "", false, "Update Byte Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateGeoMatchSet, "update-geo-match-set", "", false, "Update Geo Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateIPSet, "update-ip-set", "", false, "Update IP Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateRateBasedRule, "update-rate-based-rule", "", false, "Update Rate Based Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateRegexMatchSet, "update-regex-match-set", "", false, "Update Regex Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateRegexPatternSet, "update-regex-pattern-set", "", false, "Update Regex Pattern Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateRule, "update-rule", "", false, "Update Rule")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateRuleGroup, "update-rule-group", "", false, "Update Rule Group")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateSizeConstraintSet, "update-size-constraint-set", "", false, "Update Size Constraint Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateSqlInjectionMatchSet, "update-sql-injection-match-set", "", false, "Update Sql Injection Match Set")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateWebACL, "update-web-acl", "", false, "Update Web ACL")
	_wafregionalCmd.Flags().BoolVarP(&_wafregionalUpdateXssMatchSet, "update-xss-match-set", "", false, "Update Xss Match Set")

}
