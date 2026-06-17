package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/waf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// wafCmd represents the waf command
var _wafCmd = &cobra.Command{
	Use:   "waf",
	Short: "AWS waf CLI",
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
		client := waf.NewFromConfig(cfg)
		if _wafCreateByteMatchSet {
			waf_CreateByteMatchSet(cfg, client)
			return
		}
		if _wafCreateGeoMatchSet {
			waf_CreateGeoMatchSet(cfg, client)
			return
		}
		if _wafCreateIPSet {
			waf_CreateIPSet(cfg, client)
			return
		}
		if _wafCreateRateBasedRule {
			waf_CreateRateBasedRule(cfg, client)
			return
		}
		if _wafCreateRegexMatchSet {
			waf_CreateRegexMatchSet(cfg, client)
			return
		}
		if _wafCreateRegexPatternSet {
			waf_CreateRegexPatternSet(cfg, client)
			return
		}
		if _wafCreateRule {
			waf_CreateRule(cfg, client)
			return
		}
		if _wafCreateRuleGroup {
			waf_CreateRuleGroup(cfg, client)
			return
		}
		if _wafCreateSizeConstraintSet {
			waf_CreateSizeConstraintSet(cfg, client)
			return
		}
		if _wafCreateSqlInjectionMatchSet {
			waf_CreateSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafCreateWebACL {
			waf_CreateWebACL(cfg, client)
			return
		}
		if _wafCreateWebACLMigrationStack {
			waf_CreateWebACLMigrationStack(cfg, client)
			return
		}
		if _wafCreateXssMatchSet {
			waf_CreateXssMatchSet(cfg, client)
			return
		}
		if _wafDeleteByteMatchSet {
			waf_DeleteByteMatchSet(cfg, client)
			return
		}
		if _wafDeleteGeoMatchSet {
			waf_DeleteGeoMatchSet(cfg, client)
			return
		}
		if _wafDeleteIPSet {
			waf_DeleteIPSet(cfg, client)
			return
		}
		if _wafDeleteLoggingConfiguration {
			waf_DeleteLoggingConfiguration(cfg, client)
			return
		}
		if _wafDeletePermissionPolicy {
			waf_DeletePermissionPolicy(cfg, client)
			return
		}
		if _wafDeleteRateBasedRule {
			waf_DeleteRateBasedRule(cfg, client)
			return
		}
		if _wafDeleteRegexMatchSet {
			waf_DeleteRegexMatchSet(cfg, client)
			return
		}
		if _wafDeleteRegexPatternSet {
			waf_DeleteRegexPatternSet(cfg, client)
			return
		}
		if _wafDeleteRule {
			waf_DeleteRule(cfg, client)
			return
		}
		if _wafDeleteRuleGroup {
			waf_DeleteRuleGroup(cfg, client)
			return
		}
		if _wafDeleteSizeConstraintSet {
			waf_DeleteSizeConstraintSet(cfg, client)
			return
		}
		if _wafDeleteSqlInjectionMatchSet {
			waf_DeleteSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafDeleteWebACL {
			waf_DeleteWebACL(cfg, client)
			return
		}
		if _wafDeleteXssMatchSet {
			waf_DeleteXssMatchSet(cfg, client)
			return
		}
		if _wafGetByteMatchSet {
			waf_GetByteMatchSet(cfg, client)
			return
		}
		if _wafGetChangeToken {
			waf_GetChangeToken(cfg, client)
			return
		}
		if _wafGetChangeTokenStatus {
			waf_GetChangeTokenStatus(cfg, client)
			return
		}
		if _wafGetGeoMatchSet {
			waf_GetGeoMatchSet(cfg, client)
			return
		}
		if _wafGetIPSet {
			waf_GetIPSet(cfg, client)
			return
		}
		if _wafGetLoggingConfiguration {
			waf_GetLoggingConfiguration(cfg, client)
			return
		}
		if _wafGetPermissionPolicy {
			waf_GetPermissionPolicy(cfg, client)
			return
		}
		if _wafGetRateBasedRule {
			waf_GetRateBasedRule(cfg, client)
			return
		}
		if _wafGetRateBasedRuleManagedKeys {
			waf_GetRateBasedRuleManagedKeys(cfg, client)
			return
		}
		if _wafGetRegexMatchSet {
			waf_GetRegexMatchSet(cfg, client)
			return
		}
		if _wafGetRegexPatternSet {
			waf_GetRegexPatternSet(cfg, client)
			return
		}
		if _wafGetRule {
			waf_GetRule(cfg, client)
			return
		}
		if _wafGetRuleGroup {
			waf_GetRuleGroup(cfg, client)
			return
		}
		if _wafGetSampledRequests {
			waf_GetSampledRequests(cfg, client)
			return
		}
		if _wafGetSizeConstraintSet {
			waf_GetSizeConstraintSet(cfg, client)
			return
		}
		if _wafGetSqlInjectionMatchSet {
			waf_GetSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafGetWebACL {
			waf_GetWebACL(cfg, client)
			return
		}
		if _wafGetXssMatchSet {
			waf_GetXssMatchSet(cfg, client)
			return
		}
		if _wafListActivatedRulesInRuleGroup {
			waf_ListActivatedRulesInRuleGroup(cfg, client)
			return
		}
		if _wafListByteMatchSets {
			waf_ListByteMatchSets(cfg, client)
			return
		}
		if _wafListGeoMatchSets {
			waf_ListGeoMatchSets(cfg, client)
			return
		}
		if _wafListIPSets {
			waf_ListIPSets(cfg, client)
			return
		}
		if _wafListLoggingConfigurations {
			waf_ListLoggingConfigurations(cfg, client)
			return
		}
		if _wafListRateBasedRules {
			waf_ListRateBasedRules(cfg, client)
			return
		}
		if _wafListRegexMatchSets {
			waf_ListRegexMatchSets(cfg, client)
			return
		}
		if _wafListRegexPatternSets {
			waf_ListRegexPatternSets(cfg, client)
			return
		}
		if _wafListRuleGroups {
			waf_ListRuleGroups(cfg, client)
			return
		}
		if _wafListRules {
			waf_ListRules(cfg, client)
			return
		}
		if _wafListSizeConstraintSets {
			waf_ListSizeConstraintSets(cfg, client)
			return
		}
		if _wafListSqlInjectionMatchSets {
			waf_ListSqlInjectionMatchSets(cfg, client)
			return
		}
		if _wafListSubscribedRuleGroups {
			waf_ListSubscribedRuleGroups(cfg, client)
			return
		}
		if _wafListTagsForResource {
			waf_ListTagsForResource(cfg, client)
			return
		}
		if _wafListWebACLs {
			waf_ListWebACLs(cfg, client)
			return
		}
		if _wafListXssMatchSets {
			waf_ListXssMatchSets(cfg, client)
			return
		}
		if _wafPutLoggingConfiguration {
			waf_PutLoggingConfiguration(cfg, client)
			return
		}
		if _wafPutPermissionPolicy {
			waf_PutPermissionPolicy(cfg, client)
			return
		}
		if _wafTagResource {
			waf_TagResource(cfg, client)
			return
		}
		if _wafUntagResource {
			waf_UntagResource(cfg, client)
			return
		}
		if _wafUpdateByteMatchSet {
			waf_UpdateByteMatchSet(cfg, client)
			return
		}
		if _wafUpdateGeoMatchSet {
			waf_UpdateGeoMatchSet(cfg, client)
			return
		}
		if _wafUpdateIPSet {
			waf_UpdateIPSet(cfg, client)
			return
		}
		if _wafUpdateRateBasedRule {
			waf_UpdateRateBasedRule(cfg, client)
			return
		}
		if _wafUpdateRegexMatchSet {
			waf_UpdateRegexMatchSet(cfg, client)
			return
		}
		if _wafUpdateRegexPatternSet {
			waf_UpdateRegexPatternSet(cfg, client)
			return
		}
		if _wafUpdateRule {
			waf_UpdateRule(cfg, client)
			return
		}
		if _wafUpdateRuleGroup {
			waf_UpdateRuleGroup(cfg, client)
			return
		}
		if _wafUpdateSizeConstraintSet {
			waf_UpdateSizeConstraintSet(cfg, client)
			return
		}
		if _wafUpdateSqlInjectionMatchSet {
			waf_UpdateSqlInjectionMatchSet(cfg, client)
			return
		}
		if _wafUpdateWebACL {
			waf_UpdateWebACL(cfg, client)
			return
		}
		if _wafUpdateXssMatchSet {
			waf_UpdateXssMatchSet(cfg, client)
			return
		}

	},
}

var (
	_wafCreateByteMatchSet            bool
	_wafCreateGeoMatchSet             bool
	_wafCreateIPSet                   bool
	_wafCreateRateBasedRule           bool
	_wafCreateRegexMatchSet           bool
	_wafCreateRegexPatternSet         bool
	_wafCreateRule                    bool
	_wafCreateRuleGroup               bool
	_wafCreateSizeConstraintSet       bool
	_wafCreateSqlInjectionMatchSet    bool
	_wafCreateWebACL                  bool
	_wafCreateWebACLMigrationStack    bool
	_wafCreateXssMatchSet             bool
	_wafDeleteByteMatchSet            bool
	_wafDeleteGeoMatchSet             bool
	_wafDeleteIPSet                   bool
	_wafDeleteLoggingConfiguration    bool
	_wafDeletePermissionPolicy        bool
	_wafDeleteRateBasedRule           bool
	_wafDeleteRegexMatchSet           bool
	_wafDeleteRegexPatternSet         bool
	_wafDeleteRule                    bool
	_wafDeleteRuleGroup               bool
	_wafDeleteSizeConstraintSet       bool
	_wafDeleteSqlInjectionMatchSet    bool
	_wafDeleteWebACL                  bool
	_wafDeleteXssMatchSet             bool
	_wafGetByteMatchSet               bool
	_wafGetChangeToken                bool
	_wafGetChangeTokenStatus          bool
	_wafGetGeoMatchSet                bool
	_wafGetIPSet                      bool
	_wafGetLoggingConfiguration       bool
	_wafGetPermissionPolicy           bool
	_wafGetRateBasedRule              bool
	_wafGetRateBasedRuleManagedKeys   bool
	_wafGetRegexMatchSet              bool
	_wafGetRegexPatternSet            bool
	_wafGetRule                       bool
	_wafGetRuleGroup                  bool
	_wafGetSampledRequests            bool
	_wafGetSizeConstraintSet          bool
	_wafGetSqlInjectionMatchSet       bool
	_wafGetWebACL                     bool
	_wafGetXssMatchSet                bool
	_wafListActivatedRulesInRuleGroup bool
	_wafListByteMatchSets             bool
	_wafListGeoMatchSets              bool
	_wafListIPSets                    bool
	_wafListLoggingConfigurations     bool
	_wafListRateBasedRules            bool
	_wafListRegexMatchSets            bool
	_wafListRegexPatternSets          bool
	_wafListRuleGroups                bool
	_wafListRules                     bool
	_wafListSizeConstraintSets        bool
	_wafListSqlInjectionMatchSets     bool
	_wafListSubscribedRuleGroups      bool
	_wafListTagsForResource           bool
	_wafListWebACLs                   bool
	_wafListXssMatchSets              bool
	_wafPutLoggingConfiguration       bool
	_wafPutPermissionPolicy           bool
	_wafTagResource                   bool
	_wafUntagResource                 bool
	_wafUpdateByteMatchSet            bool
	_wafUpdateGeoMatchSet             bool
	_wafUpdateIPSet                   bool
	_wafUpdateRateBasedRule           bool
	_wafUpdateRegexMatchSet           bool
	_wafUpdateRegexPatternSet         bool
	_wafUpdateRule                    bool
	_wafUpdateRuleGroup               bool
	_wafUpdateSizeConstraintSet       bool
	_wafUpdateSqlInjectionMatchSet    bool
	_wafUpdateWebACL                  bool
	_wafUpdateXssMatchSet             bool

	_wafByteMatchSetId         string
	_wafChangeToken            string
	_wafDefaultAction          string
	_wafGeoMatchSetId          string
	_wafIgnoreUnsupportedType  string
	_wafIPSetId                string
	_wafLimit                  string
	_wafLoggingConfiguration   string
	_wafMaxItems               string
	_wafMetricName             string
	_wafName                   string
	_wafNextMarker             string
	_wafPolicy                 string
	_wafRateKey                string
	_wafRateLimit              string
	_wafRegexMatchSetId        string
	_wafRegexPatternSetId      string
	_wafResourceARN            string
	_wafRuleGroupId            string
	_wafRuleId                 string
	_wafS3BucketName           string
	_wafSizeConstraintSetId    string
	_wafSqlInjectionMatchSetId string
	_wafTagKeys                []string
	_wafTags                   string
	_wafTimeWindow             string
	_wafUpdates                string
	_wafWebAclId               string
	_wafWebACLId               string
	_wafXssMatchSetId          string
)

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
func waf_CreateByteMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateByteMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateGeoMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateGeoMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateIPSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateIPSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateRateBasedRule(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateRateBasedRuleInput{
		// ChangeToken: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
		// RateKey: types.RateKey, // Required
		// RateLimit: *int64, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafMetricName) > 0 {
		input.MetricName = aws.String(_wafMetricName)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
	}
	if len(_wafRateKey) > 0 {
		if err := assignInputField(input, "RateKey", _wafRateKey); err != nil {
			log.Errorf("invalid --rate-key: %s", err.Error())
			return
		}
	}
	if len(_wafRateLimit) > 0 {
		if err := assignInputField(input, "RateLimit", _wafRateLimit); err != nil {
			log.Errorf("invalid --rate-limit: %s", err.Error())
			return
		}
	}
	if len(_wafTags) > 0 {
		if err := assignInputField(input, "Tags", _wafTags); err != nil {
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
func waf_CreateRegexMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateRegexMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateRegexPatternSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateRegexPatternSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateRule(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateRuleInput{
		// ChangeToken: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafMetricName) > 0 {
		input.MetricName = aws.String(_wafMetricName)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
	}
	if len(_wafTags) > 0 {
		if err := assignInputField(input, "Tags", _wafTags); err != nil {
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
func waf_CreateRuleGroup(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateRuleGroupInput{
		// ChangeToken: *string, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafMetricName) > 0 {
		input.MetricName = aws.String(_wafMetricName)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
	}
	if len(_wafTags) > 0 {
		if err := assignInputField(input, "Tags", _wafTags); err != nil {
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
func waf_CreateSizeConstraintSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateSizeConstraintSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateSqlInjectionMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateSqlInjectionMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_CreateWebACL(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateWebACLInput{
		// ChangeToken: *string, // Required
		// DefaultAction: *types.WafAction, // Required
		// MetricName: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _wafDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_wafMetricName) > 0 {
		input.MetricName = aws.String(_wafMetricName)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
	}
	if len(_wafTags) > 0 {
		if err := assignInputField(input, "Tags", _wafTags); err != nil {
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
func waf_CreateWebACLMigrationStack(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateWebACLMigrationStackInput{
		// IgnoreUnsupportedType: *bool, // Required
		// S3BucketName: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafIgnoreUnsupportedType) > 0 {
		if err := assignInputField(input, "IgnoreUnsupportedType", _wafIgnoreUnsupportedType); err != nil {
			log.Errorf("invalid --ignore-unsupported-type: %s", err.Error())
			return
		}
	}
	if len(_wafS3BucketName) > 0 {
		input.S3BucketName = aws.String(_wafS3BucketName)
	}
	if len(_wafWebACLId) > 0 {
		input.WebACLId = aws.String(_wafWebACLId)
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
func waf_CreateXssMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.CreateXssMatchSetInput{
		// ChangeToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafName) > 0 {
		input.Name = aws.String(_wafName)
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
func waf_DeleteByteMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteByteMatchSetInput{
		// ByteMatchSetId: *string, // Required
		// ChangeToken: *string, // Required
	}

	if len(_wafByteMatchSetId) > 0 {
		input.ByteMatchSetId = aws.String(_wafByteMatchSetId)
	}
	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
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
func waf_DeleteGeoMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteGeoMatchSetInput{
		// ChangeToken: *string, // Required
		// GeoMatchSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafGeoMatchSetId) > 0 {
		input.GeoMatchSetId = aws.String(_wafGeoMatchSetId)
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
func waf_DeleteIPSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteIPSetInput{
		// ChangeToken: *string, // Required
		// IPSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafIPSetId) > 0 {
		input.IPSetId = aws.String(_wafIPSetId)
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
func waf_DeleteLoggingConfiguration(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteLoggingConfigurationInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafResourceARN)
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
func waf_DeletePermissionPolicy(cfg aws.Config, client *waf.Client) {
	input := &waf.DeletePermissionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafResourceARN)
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
func waf_DeleteRateBasedRule(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteRateBasedRuleInput{
		// ChangeToken: *string, // Required
		// RuleId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
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
func waf_DeleteRegexMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteRegexMatchSetInput{
		// ChangeToken: *string, // Required
		// RegexMatchSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRegexMatchSetId) > 0 {
		input.RegexMatchSetId = aws.String(_wafRegexMatchSetId)
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
func waf_DeleteRegexPatternSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteRegexPatternSetInput{
		// ChangeToken: *string, // Required
		// RegexPatternSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRegexPatternSetId) > 0 {
		input.RegexPatternSetId = aws.String(_wafRegexPatternSetId)
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
func waf_DeleteRule(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteRuleInput{
		// ChangeToken: *string, // Required
		// RuleId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
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
func waf_DeleteRuleGroup(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteRuleGroupInput{
		// ChangeToken: *string, // Required
		// RuleGroupId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafRuleGroupId)
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
func waf_DeleteSizeConstraintSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteSizeConstraintSetInput{
		// ChangeToken: *string, // Required
		// SizeConstraintSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafSizeConstraintSetId) > 0 {
		input.SizeConstraintSetId = aws.String(_wafSizeConstraintSetId)
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
func waf_DeleteSqlInjectionMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteSqlInjectionMatchSetInput{
		// ChangeToken: *string, // Required
		// SqlInjectionMatchSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafSqlInjectionMatchSetId) > 0 {
		input.SqlInjectionMatchSetId = aws.String(_wafSqlInjectionMatchSetId)
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
func waf_DeleteWebACL(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteWebACLInput{
		// ChangeToken: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafWebACLId) > 0 {
		input.WebACLId = aws.String(_wafWebACLId)
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
func waf_DeleteXssMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.DeleteXssMatchSetInput{
		// ChangeToken: *string, // Required
		// XssMatchSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafXssMatchSetId) > 0 {
		input.XssMatchSetId = aws.String(_wafXssMatchSetId)
	}

	if resp, err := client.DeleteXssMatchSet(context.TODO(), input); err != nil {
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
func waf_GetByteMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetByteMatchSetInput{
		// ByteMatchSetId: *string, // Required
	}

	if len(_wafByteMatchSetId) > 0 {
		input.ByteMatchSetId = aws.String(_wafByteMatchSetId)
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
func waf_GetChangeToken(cfg aws.Config, client *waf.Client) {
	input := &waf.GetChangeTokenInput{}

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
func waf_GetChangeTokenStatus(cfg aws.Config, client *waf.Client) {
	input := &waf.GetChangeTokenStatusInput{
		// ChangeToken: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
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
func waf_GetGeoMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetGeoMatchSetInput{
		// GeoMatchSetId: *string, // Required
	}

	if len(_wafGeoMatchSetId) > 0 {
		input.GeoMatchSetId = aws.String(_wafGeoMatchSetId)
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
func waf_GetIPSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetIPSetInput{
		// IPSetId: *string, // Required
	}

	if len(_wafIPSetId) > 0 {
		input.IPSetId = aws.String(_wafIPSetId)
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
func waf_GetLoggingConfiguration(cfg aws.Config, client *waf.Client) {
	input := &waf.GetLoggingConfigurationInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafResourceARN)
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
func waf_GetPermissionPolicy(cfg aws.Config, client *waf.Client) {
	input := &waf.GetPermissionPolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafResourceARN)
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
func waf_GetRateBasedRule(cfg aws.Config, client *waf.Client) {
	input := &waf.GetRateBasedRuleInput{
		// RuleId: *string, // Required
	}

	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
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
func waf_GetRateBasedRuleManagedKeys(cfg aws.Config, client *waf.Client) {
	input := &waf.GetRateBasedRuleManagedKeysInput{
		// RuleId: *string, // Required
	}

	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_GetRegexMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetRegexMatchSetInput{
		// RegexMatchSetId: *string, // Required
	}

	if len(_wafRegexMatchSetId) > 0 {
		input.RegexMatchSetId = aws.String(_wafRegexMatchSetId)
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
func waf_GetRegexPatternSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetRegexPatternSetInput{
		// RegexPatternSetId: *string, // Required
	}

	if len(_wafRegexPatternSetId) > 0 {
		input.RegexPatternSetId = aws.String(_wafRegexPatternSetId)
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
func waf_GetRule(cfg aws.Config, client *waf.Client) {
	input := &waf.GetRuleInput{
		// RuleId: *string, // Required
	}

	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
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
func waf_GetRuleGroup(cfg aws.Config, client *waf.Client) {
	input := &waf.GetRuleGroupInput{
		// RuleGroupId: *string, // Required
	}

	if len(_wafRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafRuleGroupId)
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
func waf_GetSampledRequests(cfg aws.Config, client *waf.Client) {
	input := &waf.GetSampledRequestsInput{
		// MaxItems: *int64, // Required
		// RuleId: *string, // Required
		// TimeWindow: *types.TimeWindow, // Required
		// WebAclId: *string, // Required
	}

	if len(_wafMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _wafMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
	}
	if len(_wafTimeWindow) > 0 {
		if err := assignInputField(input, "TimeWindow", _wafTimeWindow); err != nil {
			log.Errorf("invalid --time-window: %s", err.Error())
			return
		}
	}
	if len(_wafWebAclId) > 0 {
		input.WebAclId = aws.String(_wafWebAclId)
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
func waf_GetSizeConstraintSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetSizeConstraintSetInput{
		// SizeConstraintSetId: *string, // Required
	}

	if len(_wafSizeConstraintSetId) > 0 {
		input.SizeConstraintSetId = aws.String(_wafSizeConstraintSetId)
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
func waf_GetSqlInjectionMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetSqlInjectionMatchSetInput{
		// SqlInjectionMatchSetId: *string, // Required
	}

	if len(_wafSqlInjectionMatchSetId) > 0 {
		input.SqlInjectionMatchSetId = aws.String(_wafSqlInjectionMatchSetId)
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
func waf_GetWebACL(cfg aws.Config, client *waf.Client) {
	input := &waf.GetWebACLInput{
		// WebACLId: *string, // Required
	}

	if len(_wafWebACLId) > 0 {
		input.WebACLId = aws.String(_wafWebACLId)
	}

	if resp, err := client.GetWebACL(context.TODO(), input); err != nil {
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
func waf_GetXssMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.GetXssMatchSetInput{
		// XssMatchSetId: *string, // Required
	}

	if len(_wafXssMatchSetId) > 0 {
		input.XssMatchSetId = aws.String(_wafXssMatchSetId)
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
func waf_ListActivatedRulesInRuleGroup(cfg aws.Config, client *waf.Client) {
	input := &waf.ListActivatedRulesInRuleGroupInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
	}
	if len(_wafRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafRuleGroupId)
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
func waf_ListByteMatchSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListByteMatchSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListGeoMatchSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListGeoMatchSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListIPSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListIPSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListLoggingConfigurations(cfg aws.Config, client *waf.Client) {
	input := &waf.ListLoggingConfigurationsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListRateBasedRules(cfg aws.Config, client *waf.Client) {
	input := &waf.ListRateBasedRulesInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListRegexMatchSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListRegexMatchSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListRegexPatternSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListRegexPatternSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
	}

	if resp, err := client.ListRegexPatternSets(context.TODO(), input); err != nil {
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
func waf_ListRuleGroups(cfg aws.Config, client *waf.Client) {
	input := &waf.ListRuleGroupsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListRules(cfg aws.Config, client *waf.Client) {
	input := &waf.ListRulesInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListSizeConstraintSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListSizeConstraintSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListSqlInjectionMatchSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListSqlInjectionMatchSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListSubscribedRuleGroups(cfg aws.Config, client *waf.Client) {
	input := &waf.ListSubscribedRuleGroupsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListTagsForResource(cfg aws.Config, client *waf.Client) {
	input := &waf.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafResourceARN)
	}
	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListWebACLs(cfg aws.Config, client *waf.Client) {
	input := &waf.ListWebACLsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_ListXssMatchSets(cfg aws.Config, client *waf.Client) {
	input := &waf.ListXssMatchSetsInput{}

	if len(_wafLimit) > 0 {
		if err := assignInputField(input, "Limit", _wafLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_wafNextMarker) > 0 {
		input.NextMarker = aws.String(_wafNextMarker)
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
func waf_PutLoggingConfiguration(cfg aws.Config, client *waf.Client) {
	input := &waf.PutLoggingConfigurationInput{
		// LoggingConfiguration: *types.LoggingConfiguration, // Required
	}

	if len(_wafLoggingConfiguration) > 0 {
		if err := assignInputField(input, "LoggingConfiguration", _wafLoggingConfiguration); err != nil {
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
func waf_PutPermissionPolicy(cfg aws.Config, client *waf.Client) {
	input := &waf.PutPermissionPolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_wafPolicy) > 0 {
		input.Policy = aws.String(_wafPolicy)
	}
	if len(_wafResourceARN) > 0 {
		input.ResourceArn = aws.String(_wafResourceARN)
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
func waf_TagResource(cfg aws.Config, client *waf.Client) {
	input := &waf.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafResourceARN)
	}
	if len(_wafTags) > 0 {
		if err := assignInputField(input, "Tags", _wafTags); err != nil {
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
func waf_UntagResource(cfg aws.Config, client *waf.Client) {
	input := &waf.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_wafResourceARN) > 0 {
		input.ResourceARN = aws.String(_wafResourceARN)
	}
	if len(_wafTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _wafTagKeys...)
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
func waf_UpdateByteMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateByteMatchSetInput{
		// ByteMatchSetId: *string, // Required
		// ChangeToken: *string, // Required
		// Updates: []types.ByteMatchSetUpdate, // Required
	}

	if len(_wafByteMatchSetId) > 0 {
		input.ByteMatchSetId = aws.String(_wafByteMatchSetId)
	}
	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateGeoMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateGeoMatchSetInput{
		// ChangeToken: *string, // Required
		// GeoMatchSetId: *string, // Required
		// Updates: []types.GeoMatchSetUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafGeoMatchSetId) > 0 {
		input.GeoMatchSetId = aws.String(_wafGeoMatchSetId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateIPSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateIPSetInput{
		// ChangeToken: *string, // Required
		// IPSetId: *string, // Required
		// Updates: []types.IPSetUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafIPSetId) > 0 {
		input.IPSetId = aws.String(_wafIPSetId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateRateBasedRule(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateRateBasedRuleInput{
		// ChangeToken: *string, // Required
		// RateLimit: *int64, // Required
		// RuleId: *string, // Required
		// Updates: []types.RuleUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRateLimit) > 0 {
		if err := assignInputField(input, "RateLimit", _wafRateLimit); err != nil {
			log.Errorf("invalid --rate-limit: %s", err.Error())
			return
		}
	}
	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateRegexMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateRegexMatchSetInput{
		// ChangeToken: *string, // Required
		// RegexMatchSetId: *string, // Required
		// Updates: []types.RegexMatchSetUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRegexMatchSetId) > 0 {
		input.RegexMatchSetId = aws.String(_wafRegexMatchSetId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateRegexPatternSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateRegexPatternSetInput{
		// ChangeToken: *string, // Required
		// RegexPatternSetId: *string, // Required
		// Updates: []types.RegexPatternSetUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRegexPatternSetId) > 0 {
		input.RegexPatternSetId = aws.String(_wafRegexPatternSetId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateRule(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateRuleInput{
		// ChangeToken: *string, // Required
		// RuleId: *string, // Required
		// Updates: []types.RuleUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRuleId) > 0 {
		input.RuleId = aws.String(_wafRuleId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateRuleGroup(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateRuleGroupInput{
		// ChangeToken: *string, // Required
		// RuleGroupId: *string, // Required
		// Updates: []types.RuleGroupUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafRuleGroupId) > 0 {
		input.RuleGroupId = aws.String(_wafRuleGroupId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateSizeConstraintSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateSizeConstraintSetInput{
		// ChangeToken: *string, // Required
		// SizeConstraintSetId: *string, // Required
		// Updates: []types.SizeConstraintSetUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafSizeConstraintSetId) > 0 {
		input.SizeConstraintSetId = aws.String(_wafSizeConstraintSetId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateSqlInjectionMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateSqlInjectionMatchSetInput{
		// ChangeToken: *string, // Required
		// SqlInjectionMatchSetId: *string, // Required
		// Updates: []types.SqlInjectionMatchSetUpdate, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafSqlInjectionMatchSetId) > 0 {
		input.SqlInjectionMatchSetId = aws.String(_wafSqlInjectionMatchSetId)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateWebACL(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateWebACLInput{
		// ChangeToken: *string, // Required
		// WebACLId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafWebACLId) > 0 {
		input.WebACLId = aws.String(_wafWebACLId)
	}
	if len(_wafDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _wafDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
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
func waf_UpdateXssMatchSet(cfg aws.Config, client *waf.Client) {
	input := &waf.UpdateXssMatchSetInput{
		// ChangeToken: *string, // Required
		// Updates: []types.XssMatchSetUpdate, // Required
		// XssMatchSetId: *string, // Required
	}

	if len(_wafChangeToken) > 0 {
		input.ChangeToken = aws.String(_wafChangeToken)
	}
	if len(_wafUpdates) > 0 {
		if err := assignInputField(input, "Updates", _wafUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}
	if len(_wafXssMatchSetId) > 0 {
		input.XssMatchSetId = aws.String(_wafXssMatchSetId)
	}

	if resp, err := client.UpdateXssMatchSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_wafCmd)
	_wafCmd.Flags().SortFlags = false

	_wafCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_wafCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_wafCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_wafCmd.Flags().StringVarP(&_wafByteMatchSetId, "byte-match-set-id", "", "", "Byte Match Set ID")
	_wafCmd.Flags().StringVarP(&_wafChangeToken, "change-token", "", "", "Change Token")
	_wafCmd.Flags().StringVarP(&_wafDefaultAction, "default-action", "", "", "Default Action")
	_wafCmd.Flags().StringVarP(&_wafGeoMatchSetId, "geo-match-set-id", "", "", "Geo Match Set ID")
	_wafCmd.Flags().StringVarP(&_wafIgnoreUnsupportedType, "ignore-unsupported-type", "", "", "Ignore Unsupported Type")
	_wafCmd.Flags().StringVarP(&_wafIPSetId, "ip-set-id", "", "", "IP Set ID")
	_wafCmd.Flags().StringVarP(&_wafLimit, "limit", "", "", "Limit")
	_wafCmd.Flags().StringVarP(&_wafLoggingConfiguration, "logging-configuration", "", "", "Logging Configuration")
	_wafCmd.Flags().StringVarP(&_wafMaxItems, "max-items", "", "", "Max Items")
	_wafCmd.Flags().StringVarP(&_wafMetricName, "metric-name", "", "", "Metric Name")
	_wafCmd.Flags().StringVarP(&_wafName, "name", "", "", "Name")
	_wafCmd.Flags().StringVarP(&_wafNextMarker, "next-marker", "", "", "Next Marker")
	_wafCmd.Flags().StringVarP(&_wafPolicy, "policy", "", "", "Policy")
	_wafCmd.Flags().StringVarP(&_wafRateKey, "rate-key", "", "", "Rate Key")
	_wafCmd.Flags().StringVarP(&_wafRateLimit, "rate-limit", "", "", "Rate Limit")
	_wafCmd.Flags().StringVarP(&_wafRegexMatchSetId, "regex-match-set-id", "", "", "Regex Match Set ID")
	_wafCmd.Flags().StringVarP(&_wafRegexPatternSetId, "regex-pattern-set-id", "", "", "Regex Pattern Set ID")
	_wafCmd.Flags().StringVarP(&_wafResourceARN, "resource-arn", "", "", "Resource ARN")
	_wafCmd.Flags().StringVarP(&_wafRuleGroupId, "rule-group-id", "", "", "Rule Group ID")
	_wafCmd.Flags().StringVarP(&_wafRuleId, "rule-id", "", "", "Rule ID")
	_wafCmd.Flags().StringVarP(&_wafS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_wafCmd.Flags().StringVarP(&_wafSizeConstraintSetId, "size-constraint-set-id", "", "", "Size Constraint Set ID")
	_wafCmd.Flags().StringVarP(&_wafSqlInjectionMatchSetId, "sql-injection-match-set-id", "", "", "Sql Injection Match Set ID")
	_wafCmd.Flags().StringSliceVarP(&_wafTagKeys, "tag-keys", "", nil, "Tag Keys")
	_wafCmd.Flags().StringVarP(&_wafTags, "tags", "", "", "Tags")
	_wafCmd.Flags().StringVarP(&_wafTimeWindow, "time-window", "", "", "Time Window")
	_wafCmd.Flags().StringVarP(&_wafUpdates, "updates", "", "", "Updates")
	_wafCmd.Flags().StringVarP(&_wafWebAclId, "web-acl-id", "", "", "Web ACL ID")
	_wafCmd.Flags().StringVarP(&_wafWebACLId, "web-aclid", "", "", "Web Aclid")
	_wafCmd.Flags().StringVarP(&_wafXssMatchSetId, "xss-match-set-id", "", "", "Xss Match Set ID")

	_wafCmd.Flags().BoolVarP(&_wafCreateByteMatchSet, "create-byte-match-set", "", false, "Create Byte Match Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateGeoMatchSet, "create-geo-match-set", "", false, "Create Geo Match Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateIPSet, "create-ip-set", "", false, "Create IP Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateRateBasedRule, "create-rate-based-rule", "", false, "Create Rate Based Rule")
	_wafCmd.Flags().BoolVarP(&_wafCreateRegexMatchSet, "create-regex-match-set", "", false, "Create Regex Match Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateRegexPatternSet, "create-regex-pattern-set", "", false, "Create Regex Pattern Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateRule, "create-rule", "", false, "Create Rule")
	_wafCmd.Flags().BoolVarP(&_wafCreateRuleGroup, "create-rule-group", "", false, "Create Rule Group")
	_wafCmd.Flags().BoolVarP(&_wafCreateSizeConstraintSet, "create-size-constraint-set", "", false, "Create Size Constraint Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateSqlInjectionMatchSet, "create-sql-injection-match-set", "", false, "Create Sql Injection Match Set")
	_wafCmd.Flags().BoolVarP(&_wafCreateWebACL, "create-web-acl", "", false, "Create Web ACL")
	_wafCmd.Flags().BoolVarP(&_wafCreateWebACLMigrationStack, "create-web-acl-migration-stack", "", false, "Create Web ACL Migration Stack")
	_wafCmd.Flags().BoolVarP(&_wafCreateXssMatchSet, "create-xss-match-set", "", false, "Create Xss Match Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteByteMatchSet, "delete-byte-match-set", "", false, "Delete Byte Match Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteGeoMatchSet, "delete-geo-match-set", "", false, "Delete Geo Match Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteIPSet, "delete-ip-set", "", false, "Delete IP Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteLoggingConfiguration, "delete-logging-configuration", "", false, "Delete Logging Configuration")
	_wafCmd.Flags().BoolVarP(&_wafDeletePermissionPolicy, "delete-permission-policy", "", false, "Delete Permission Policy")
	_wafCmd.Flags().BoolVarP(&_wafDeleteRateBasedRule, "delete-rate-based-rule", "", false, "Delete Rate Based Rule")
	_wafCmd.Flags().BoolVarP(&_wafDeleteRegexMatchSet, "delete-regex-match-set", "", false, "Delete Regex Match Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteRegexPatternSet, "delete-regex-pattern-set", "", false, "Delete Regex Pattern Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteRule, "delete-rule", "", false, "Delete Rule")
	_wafCmd.Flags().BoolVarP(&_wafDeleteRuleGroup, "delete-rule-group", "", false, "Delete Rule Group")
	_wafCmd.Flags().BoolVarP(&_wafDeleteSizeConstraintSet, "delete-size-constraint-set", "", false, "Delete Size Constraint Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteSqlInjectionMatchSet, "delete-sql-injection-match-set", "", false, "Delete Sql Injection Match Set")
	_wafCmd.Flags().BoolVarP(&_wafDeleteWebACL, "delete-web-acl", "", false, "Delete Web ACL")
	_wafCmd.Flags().BoolVarP(&_wafDeleteXssMatchSet, "delete-xss-match-set", "", false, "Delete Xss Match Set")
	_wafCmd.Flags().BoolVarP(&_wafGetByteMatchSet, "get-byte-match-set", "", false, "Get Byte Match Set")
	_wafCmd.Flags().BoolVarP(&_wafGetChangeToken, "get-change-token", "", false, "Get Change Token")
	_wafCmd.Flags().BoolVarP(&_wafGetChangeTokenStatus, "get-change-token-status", "", false, "Get Change Token Status")
	_wafCmd.Flags().BoolVarP(&_wafGetGeoMatchSet, "get-geo-match-set", "", false, "Get Geo Match Set")
	_wafCmd.Flags().BoolVarP(&_wafGetIPSet, "get-ip-set", "", false, "Get IP Set")
	_wafCmd.Flags().BoolVarP(&_wafGetLoggingConfiguration, "get-logging-configuration", "", false, "Get Logging Configuration")
	_wafCmd.Flags().BoolVarP(&_wafGetPermissionPolicy, "get-permission-policy", "", false, "Get Permission Policy")
	_wafCmd.Flags().BoolVarP(&_wafGetRateBasedRule, "get-rate-based-rule", "", false, "Get Rate Based Rule")
	_wafCmd.Flags().BoolVarP(&_wafGetRateBasedRuleManagedKeys, "get-rate-based-rule-managed-keys", "", false, "Get Rate Based Rule Managed Keys")
	_wafCmd.Flags().BoolVarP(&_wafGetRegexMatchSet, "get-regex-match-set", "", false, "Get Regex Match Set")
	_wafCmd.Flags().BoolVarP(&_wafGetRegexPatternSet, "get-regex-pattern-set", "", false, "Get Regex Pattern Set")
	_wafCmd.Flags().BoolVarP(&_wafGetRule, "get-rule", "", false, "Get Rule")
	_wafCmd.Flags().BoolVarP(&_wafGetRuleGroup, "get-rule-group", "", false, "Get Rule Group")
	_wafCmd.Flags().BoolVarP(&_wafGetSampledRequests, "get-sampled-requests", "", false, "Get Sampled Requests")
	_wafCmd.Flags().BoolVarP(&_wafGetSizeConstraintSet, "get-size-constraint-set", "", false, "Get Size Constraint Set")
	_wafCmd.Flags().BoolVarP(&_wafGetSqlInjectionMatchSet, "get-sql-injection-match-set", "", false, "Get Sql Injection Match Set")
	_wafCmd.Flags().BoolVarP(&_wafGetWebACL, "get-web-acl", "", false, "Get Web ACL")
	_wafCmd.Flags().BoolVarP(&_wafGetXssMatchSet, "get-xss-match-set", "", false, "Get Xss Match Set")
	_wafCmd.Flags().BoolVarP(&_wafListActivatedRulesInRuleGroup, "list-activated-rules-in-rule-group", "", false, "List Activated Rules In Rule Group")
	_wafCmd.Flags().BoolVarP(&_wafListByteMatchSets, "list-byte-match-sets", "", false, "List Byte Match Sets")
	_wafCmd.Flags().BoolVarP(&_wafListGeoMatchSets, "list-geo-match-sets", "", false, "List Geo Match Sets")
	_wafCmd.Flags().BoolVarP(&_wafListIPSets, "list-ip-sets", "", false, "List IP Sets")
	_wafCmd.Flags().BoolVarP(&_wafListLoggingConfigurations, "list-logging-configurations", "", false, "List Logging Configurations")
	_wafCmd.Flags().BoolVarP(&_wafListRateBasedRules, "list-rate-based-rules", "", false, "List Rate Based Rules")
	_wafCmd.Flags().BoolVarP(&_wafListRegexMatchSets, "list-regex-match-sets", "", false, "List Regex Match Sets")
	_wafCmd.Flags().BoolVarP(&_wafListRegexPatternSets, "list-regex-pattern-sets", "", false, "List Regex Pattern Sets")
	_wafCmd.Flags().BoolVarP(&_wafListRuleGroups, "list-rule-groups", "", false, "List Rule Groups")
	_wafCmd.Flags().BoolVarP(&_wafListRules, "list-rules", "", false, "List Rules")
	_wafCmd.Flags().BoolVarP(&_wafListSizeConstraintSets, "list-size-constraint-sets", "", false, "List Size Constraint Sets")
	_wafCmd.Flags().BoolVarP(&_wafListSqlInjectionMatchSets, "list-sql-injection-match-sets", "", false, "List Sql Injection Match Sets")
	_wafCmd.Flags().BoolVarP(&_wafListSubscribedRuleGroups, "list-subscribed-rule-groups", "", false, "List Subscribed Rule Groups")
	_wafCmd.Flags().BoolVarP(&_wafListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_wafCmd.Flags().BoolVarP(&_wafListWebACLs, "list-web-acls", "", false, "List Web Acls")
	_wafCmd.Flags().BoolVarP(&_wafListXssMatchSets, "list-xss-match-sets", "", false, "List Xss Match Sets")
	_wafCmd.Flags().BoolVarP(&_wafPutLoggingConfiguration, "put-logging-configuration", "", false, "Put Logging Configuration")
	_wafCmd.Flags().BoolVarP(&_wafPutPermissionPolicy, "put-permission-policy", "", false, "Put Permission Policy")
	_wafCmd.Flags().BoolVarP(&_wafTagResource, "tag-resource", "", false, "Tag Resource")
	_wafCmd.Flags().BoolVarP(&_wafUntagResource, "untag-resource", "", false, "Untag Resource")
	_wafCmd.Flags().BoolVarP(&_wafUpdateByteMatchSet, "update-byte-match-set", "", false, "Update Byte Match Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateGeoMatchSet, "update-geo-match-set", "", false, "Update Geo Match Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateIPSet, "update-ip-set", "", false, "Update IP Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateRateBasedRule, "update-rate-based-rule", "", false, "Update Rate Based Rule")
	_wafCmd.Flags().BoolVarP(&_wafUpdateRegexMatchSet, "update-regex-match-set", "", false, "Update Regex Match Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateRegexPatternSet, "update-regex-pattern-set", "", false, "Update Regex Pattern Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateRule, "update-rule", "", false, "Update Rule")
	_wafCmd.Flags().BoolVarP(&_wafUpdateRuleGroup, "update-rule-group", "", false, "Update Rule Group")
	_wafCmd.Flags().BoolVarP(&_wafUpdateSizeConstraintSet, "update-size-constraint-set", "", false, "Update Size Constraint Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateSqlInjectionMatchSet, "update-sql-injection-match-set", "", false, "Update Sql Injection Match Set")
	_wafCmd.Flags().BoolVarP(&_wafUpdateWebACL, "update-web-acl", "", false, "Update Web ACL")
	_wafCmd.Flags().BoolVarP(&_wafUpdateXssMatchSet, "update-xss-match-set", "", false, "Update Xss Match Set")

}
