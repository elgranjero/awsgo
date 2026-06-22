package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/wafv2"
)

var fields_associate_web_acl = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "WebACLArn", Flag: "web-acl-arn", Type: "*string", Required: true},
}

var fields_check_capacity = []leanruntime.Field{
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_create_api_key = []leanruntime.Field{
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "TokenDomains", Flag: "token-domains", Type: "[]string", Required: true},
}

var fields_create_ip_set = []leanruntime.Field{
	{Name: "Addresses", Flag: "addresses", Type: "[]string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IPAddressVersion", Flag: "ip-address-version", Type: "types.IPAddressVersion", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_regex_pattern_set = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RegularExpressionList", Flag: "regular-expression-list", Type: "[]types.Regex", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_rule_group = []leanruntime.Field{
	{Name: "Capacity", Flag: "capacity", Type: "*int64", Required: true},
	{Name: "CustomResponseBodies", Flag: "custom-response-bodies", Type: "map[string]types.CustomResponseBody", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VisibilityConfig", Flag: "visibility-config", Type: "*types.VisibilityConfig", Required: true},
}

var fields_create_web_acl = []leanruntime.Field{
	{Name: "ApplicationConfig", Flag: "application-config", Type: "*types.ApplicationConfig", Required: false},
	{Name: "AssociationConfig", Flag: "association-config", Type: "*types.AssociationConfig", Required: false},
	{Name: "CaptchaConfig", Flag: "captcha-config", Type: "*types.CaptchaConfig", Required: false},
	{Name: "ChallengeConfig", Flag: "challenge-config", Type: "*types.ChallengeConfig", Required: false},
	{Name: "CustomResponseBodies", Flag: "custom-response-bodies", Type: "map[string]types.CustomResponseBody", Required: false},
	{Name: "DataProtectionConfig", Flag: "data-protection-config", Type: "*types.DataProtectionConfig", Required: false},
	{Name: "DefaultAction", Flag: "default-action", Type: "*types.DefaultAction", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OnSourceDDoSProtectionConfig", Flag: "on-source-ddos-protection-config", Type: "*types.OnSourceDDoSProtectionConfig", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TokenDomains", Flag: "token-domains", Type: "[]string", Required: false},
	{Name: "VisibilityConfig", Flag: "visibility-config", Type: "*types.VisibilityConfig", Required: true},
}

var fields_delete_api_key = []leanruntime.Field{
	{Name: "APIKey", Flag: "api-key", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_delete_firewall_manager_rule_groups = []leanruntime.Field{
	{Name: "WebACLArn", Flag: "web-acl-arn", Type: "*string", Required: true},
	{Name: "WebACLLockToken", Flag: "web-acl-lock-token", Type: "*string", Required: true},
}

var fields_delete_ip_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_delete_logging_configuration = []leanruntime.Field{
	{Name: "LogScope", Flag: "log-scope", Type: "types.LogScope", Required: false},
	{Name: "LogType", Flag: "log-type", Type: "types.LogType", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_permission_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_regex_pattern_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_delete_rule_group = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_delete_web_acl = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_describe_all_managed_products = []leanruntime.Field{
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_describe_managed_products_by_vendor = []leanruntime.Field{
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "VendorName", Flag: "vendor-name", Type: "*string", Required: true},
}

var fields_describe_managed_rule_group = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "VendorName", Flag: "vendor-name", Type: "*string", Required: true},
	{Name: "VersionName", Flag: "version-name", Type: "*string", Required: false},
}

var fields_disassociate_web_acl = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_generate_mobile_sdk_release_url = []leanruntime.Field{
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "ReleaseVersion", Flag: "release-version", Type: "*string", Required: true},
}

var fields_get_decrypted_api_key = []leanruntime.Field{
	{Name: "APIKey", Flag: "api-key", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_get_ip_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_get_logging_configuration = []leanruntime.Field{
	{Name: "LogScope", Flag: "log-scope", Type: "types.LogScope", Required: false},
	{Name: "LogType", Flag: "log-type", Type: "types.LogType", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_managed_rule_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_get_mobile_sdk_release = []leanruntime.Field{
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
	{Name: "ReleaseVersion", Flag: "release-version", Type: "*string", Required: true},
}

var fields_get_permission_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_rate_based_statement_managed_keys = []leanruntime.Field{
	{Name: "RuleGroupRuleName", Flag: "rule-group-rule-name", Type: "*string", Required: false},
	{Name: "RuleName", Flag: "rule-name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
	{Name: "WebACLName", Flag: "web-acl-name", Type: "*string", Required: true},
}

var fields_get_regex_pattern_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_get_rule_group = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: false},
}

var fields_get_sampled_requests = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int64", Required: true},
	{Name: "RuleMetricName", Flag: "rule-metric-name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "TimeWindow", Flag: "time-window", Type: "*types.TimeWindow", Required: true},
	{Name: "WebAclArn", Flag: "web-acl-arn", Type: "*string", Required: true},
}

var fields_get_top_path_statistics_by_traffic = []leanruntime.Field{
	{Name: "BotCategory", Flag: "bot-category", Type: "*string", Required: false},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: false},
	{Name: "BotOrganization", Flag: "bot-organization", Type: "*string", Required: false},
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: true},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "NumberOfTopTrafficBotsPerPath", Flag: "number-of-top-traffic-bots-per-path", Type: "*int32", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "TimeWindow", Flag: "time-window", Type: "*types.TimeWindow", Required: true},
	{Name: "UriPathPrefix", Flag: "uri-path-prefix", Type: "*string", Required: false},
	{Name: "WebAclArn", Flag: "web-acl-arn", Type: "*string", Required: true},
}

var fields_get_web_acl = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: false},
}

var fields_get_web_acl_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_api_keys = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_available_managed_rule_group_versions = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "VendorName", Flag: "vendor-name", Type: "*string", Required: true},
}

var fields_list_available_managed_rule_groups = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_ip_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_logging_configurations = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "LogScope", Flag: "log-scope", Type: "types.LogScope", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_managed_rule_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_mobile_sdk_releases = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Platform", Flag: "platform", Type: "types.Platform", Required: true},
}

var fields_list_regex_pattern_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_resources_for_web_acl = []leanruntime.Field{
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: false},
	{Name: "WebACLArn", Flag: "web-acl-arn", Type: "*string", Required: true},
}

var fields_list_rule_groups = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_web_acls = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "*int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_put_logging_configuration = []leanruntime.Field{
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: true},
}

var fields_put_managed_rule_set_versions = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RecommendedVersion", Flag: "recommended-version", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "VersionsToPublish", Flag: "versions-to-publish", Type: "map[string]types.VersionToPublish", Required: false},
}

var fields_put_permission_policy = []leanruntime.Field{
	{Name: "Policy", Flag: "policy", Type: "*string", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_ip_set = []leanruntime.Field{
	{Name: "Addresses", Flag: "addresses", Type: "[]string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_update_managed_rule_set_version_expiry_date = []leanruntime.Field{
	{Name: "ExpiryTimestamp", Flag: "expiry-timestamp", Type: "*time.Time", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "VersionToExpire", Flag: "version-to-expire", Type: "*string", Required: true},
}

var fields_update_regex_pattern_set = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RegularExpressionList", Flag: "regular-expression-list", Type: "[]types.Regex", Required: true},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
}

var fields_update_rule_group = []leanruntime.Field{
	{Name: "CustomResponseBodies", Flag: "custom-response-bodies", Type: "map[string]types.CustomResponseBody", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "VisibilityConfig", Flag: "visibility-config", Type: "*types.VisibilityConfig", Required: true},
}

var fields_update_web_acl = []leanruntime.Field{
	{Name: "ApplicationConfig", Flag: "application-config", Type: "*types.ApplicationConfig", Required: false},
	{Name: "AssociationConfig", Flag: "association-config", Type: "*types.AssociationConfig", Required: false},
	{Name: "CaptchaConfig", Flag: "captcha-config", Type: "*types.CaptchaConfig", Required: false},
	{Name: "ChallengeConfig", Flag: "challenge-config", Type: "*types.ChallengeConfig", Required: false},
	{Name: "CustomResponseBodies", Flag: "custom-response-bodies", Type: "map[string]types.CustomResponseBody", Required: false},
	{Name: "DataProtectionConfig", Flag: "data-protection-config", Type: "*types.DataProtectionConfig", Required: false},
	{Name: "DefaultAction", Flag: "default-action", Type: "*types.DefaultAction", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "LockToken", Flag: "lock-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OnSourceDDoSProtectionConfig", Flag: "on-source-ddos-protection-config", Type: "*types.OnSourceDDoSProtectionConfig", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.Rule", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.Scope", Required: true},
	{Name: "TokenDomains", Flag: "token-domains", Type: "[]string", Required: false},
	{Name: "VisibilityConfig", Flag: "visibility-config", Type: "*types.VisibilityConfig", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-web-acl": {
			Name:   "associate-web-acl",
			Fields: fields_associate_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateWebACL(ctx, input)
			},
		},
		"check-capacity": {
			Name:   "check-capacity",
			Fields: fields_check_capacity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckCapacityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_capacity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckCapacity(ctx, input)
			},
		},
		"create-api-key": {
			Name:   "create-api-key",
			Fields: fields_create_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAPIKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAPIKey(ctx, input)
			},
		},
		"create-ip-set": {
			Name:   "create-ip-set",
			Fields: fields_create_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIPSet(ctx, input)
			},
		},
		"create-regex-pattern-set": {
			Name:   "create-regex-pattern-set",
			Fields: fields_create_regex_pattern_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegexPatternSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_regex_pattern_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegexPatternSet(ctx, input)
			},
		},
		"create-rule-group": {
			Name:   "create-rule-group",
			Fields: fields_create_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRuleGroup(ctx, input)
			},
		},
		"create-web-acl": {
			Name:   "create-web-acl",
			Fields: fields_create_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebACL(ctx, input)
			},
		},
		"delete-api-key": {
			Name:   "delete-api-key",
			Fields: fields_delete_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAPIKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAPIKey(ctx, input)
			},
		},
		"delete-firewall-manager-rule-groups": {
			Name:   "delete-firewall-manager-rule-groups",
			Fields: fields_delete_firewall_manager_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFirewallManagerRuleGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_firewall_manager_rule_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFirewallManagerRuleGroups(ctx, input)
			},
		},
		"delete-ip-set": {
			Name:   "delete-ip-set",
			Fields: fields_delete_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIPSet(ctx, input)
			},
		},
		"delete-logging-configuration": {
			Name:   "delete-logging-configuration",
			Fields: fields_delete_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLoggingConfiguration(ctx, input)
			},
		},
		"delete-permission-policy": {
			Name:   "delete-permission-policy",
			Fields: fields_delete_permission_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePermissionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_permission_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePermissionPolicy(ctx, input)
			},
		},
		"delete-regex-pattern-set": {
			Name:   "delete-regex-pattern-set",
			Fields: fields_delete_regex_pattern_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegexPatternSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_regex_pattern_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegexPatternSet(ctx, input)
			},
		},
		"delete-rule-group": {
			Name:   "delete-rule-group",
			Fields: fields_delete_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRuleGroup(ctx, input)
			},
		},
		"delete-web-acl": {
			Name:   "delete-web-acl",
			Fields: fields_delete_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWebACL(ctx, input)
			},
		},
		"describe-all-managed-products": {
			Name:   "describe-all-managed-products",
			Fields: fields_describe_all_managed_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAllManagedProductsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_all_managed_products, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAllManagedProducts(ctx, input)
			},
		},
		"describe-managed-products-by-vendor": {
			Name:   "describe-managed-products-by-vendor",
			Fields: fields_describe_managed_products_by_vendor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedProductsByVendorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_managed_products_by_vendor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeManagedProductsByVendor(ctx, input)
			},
		},
		"describe-managed-rule-group": {
			Name:   "describe-managed-rule-group",
			Fields: fields_describe_managed_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeManagedRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_managed_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeManagedRuleGroup(ctx, input)
			},
		},
		"disassociate-web-acl": {
			Name:   "disassociate-web-acl",
			Fields: fields_disassociate_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateWebACL(ctx, input)
			},
		},
		"generate-mobile-sdk-release-url": {
			Name:   "generate-mobile-sdk-release-url",
			Fields: fields_generate_mobile_sdk_release_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateMobileSdkReleaseUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_mobile_sdk_release_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateMobileSdkReleaseUrl(ctx, input)
			},
		},
		"get-decrypted-api-key": {
			Name:   "get-decrypted-api-key",
			Fields: fields_get_decrypted_api_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDecryptedAPIKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_decrypted_api_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDecryptedAPIKey(ctx, input)
			},
		},
		"get-ip-set": {
			Name:   "get-ip-set",
			Fields: fields_get_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIPSet(ctx, input)
			},
		},
		"get-logging-configuration": {
			Name:   "get-logging-configuration",
			Fields: fields_get_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoggingConfiguration(ctx, input)
			},
		},
		"get-managed-rule-set": {
			Name:   "get-managed-rule-set",
			Fields: fields_get_managed_rule_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetManagedRuleSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_managed_rule_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetManagedRuleSet(ctx, input)
			},
		},
		"get-mobile-sdk-release": {
			Name:   "get-mobile-sdk-release",
			Fields: fields_get_mobile_sdk_release,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMobileSdkReleaseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_mobile_sdk_release, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMobileSdkRelease(ctx, input)
			},
		},
		"get-permission-policy": {
			Name:   "get-permission-policy",
			Fields: fields_get_permission_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPermissionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_permission_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPermissionPolicy(ctx, input)
			},
		},
		"get-rate-based-statement-managed-keys": {
			Name:   "get-rate-based-statement-managed-keys",
			Fields: fields_get_rate_based_statement_managed_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRateBasedStatementManagedKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rate_based_statement_managed_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRateBasedStatementManagedKeys(ctx, input)
			},
		},
		"get-regex-pattern-set": {
			Name:   "get-regex-pattern-set",
			Fields: fields_get_regex_pattern_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegexPatternSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_regex_pattern_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegexPatternSet(ctx, input)
			},
		},
		"get-rule-group": {
			Name:   "get-rule-group",
			Fields: fields_get_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRuleGroup(ctx, input)
			},
		},
		"get-sampled-requests": {
			Name:   "get-sampled-requests",
			Fields: fields_get_sampled_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSampledRequestsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sampled_requests, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSampledRequests(ctx, input)
			},
		},
		"get-top-path-statistics-by-traffic": {
			Name:   "get-top-path-statistics-by-traffic",
			Fields: fields_get_top_path_statistics_by_traffic,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTopPathStatisticsByTrafficInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_top_path_statistics_by_traffic, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTopPathStatisticsByTraffic(ctx, input)
			},
		},
		"get-web-acl": {
			Name:   "get-web-acl",
			Fields: fields_get_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWebACL(ctx, input)
			},
		},
		"get-web-acl-for-resource": {
			Name:   "get-web-acl-for-resource",
			Fields: fields_get_web_acl_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWebACLForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_web_acl_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWebACLForResource(ctx, input)
			},
		},
		"list-api-keys": {
			Name:   "list-api-keys",
			Fields: fields_list_api_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAPIKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_api_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAPIKeys(ctx, input)
			},
		},
		"list-available-managed-rule-group-versions": {
			Name:   "list-available-managed-rule-group-versions",
			Fields: fields_list_available_managed_rule_group_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableManagedRuleGroupVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_available_managed_rule_group_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAvailableManagedRuleGroupVersions(ctx, input)
			},
		},
		"list-available-managed-rule-groups": {
			Name:   "list-available-managed-rule-groups",
			Fields: fields_list_available_managed_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableManagedRuleGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_available_managed_rule_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAvailableManagedRuleGroups(ctx, input)
			},
		},
		"list-ip-sets": {
			Name:   "list-ip-sets",
			Fields: fields_list_ip_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIPSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_ip_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListIPSets(ctx, input)
			},
		},
		"list-logging-configurations": {
			Name:   "list-logging-configurations",
			Fields: fields_list_logging_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLoggingConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_logging_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLoggingConfigurations(ctx, input)
			},
		},
		"list-managed-rule-sets": {
			Name:   "list-managed-rule-sets",
			Fields: fields_list_managed_rule_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListManagedRuleSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_managed_rule_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListManagedRuleSets(ctx, input)
			},
		},
		"list-mobile-sdk-releases": {
			Name:   "list-mobile-sdk-releases",
			Fields: fields_list_mobile_sdk_releases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMobileSdkReleasesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_mobile_sdk_releases, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListMobileSdkReleases(ctx, input)
			},
		},
		"list-regex-pattern-sets": {
			Name:   "list-regex-pattern-sets",
			Fields: fields_list_regex_pattern_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegexPatternSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_regex_pattern_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRegexPatternSets(ctx, input)
			},
		},
		"list-resources-for-web-acl": {
			Name:   "list-resources-for-web-acl",
			Fields: fields_list_resources_for_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesForWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resources_for_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourcesForWebACL(ctx, input)
			},
		},
		"list-rule-groups": {
			Name:   "list-rule-groups",
			Fields: fields_list_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRuleGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rule_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRuleGroups(ctx, input)
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-web-acls": {
			Name:   "list-web-acls",
			Fields: fields_list_web_acls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWebACLsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_web_acls, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListWebACLs(ctx, input)
			},
		},
		"put-logging-configuration": {
			Name:   "put-logging-configuration",
			Fields: fields_put_logging_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLoggingConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_logging_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLoggingConfiguration(ctx, input)
			},
		},
		"put-managed-rule-set-versions": {
			Name:   "put-managed-rule-set-versions",
			Fields: fields_put_managed_rule_set_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutManagedRuleSetVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_managed_rule_set_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutManagedRuleSetVersions(ctx, input)
			},
		},
		"put-permission-policy": {
			Name:   "put-permission-policy",
			Fields: fields_put_permission_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutPermissionPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_permission_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutPermissionPolicy(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-ip-set": {
			Name:   "update-ip-set",
			Fields: fields_update_ip_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIPSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ip_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIPSet(ctx, input)
			},
		},
		"update-managed-rule-set-version-expiry-date": {
			Name:   "update-managed-rule-set-version-expiry-date",
			Fields: fields_update_managed_rule_set_version_expiry_date,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateManagedRuleSetVersionExpiryDateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_managed_rule_set_version_expiry_date, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateManagedRuleSetVersionExpiryDate(ctx, input)
			},
		},
		"update-regex-pattern-set": {
			Name:   "update-regex-pattern-set",
			Fields: fields_update_regex_pattern_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRegexPatternSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_regex_pattern_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRegexPatternSet(ctx, input)
			},
		},
		"update-rule-group": {
			Name:   "update-rule-group",
			Fields: fields_update_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRuleGroup(ctx, input)
			},
		},
		"update-web-acl": {
			Name:   "update-web-acl",
			Fields: fields_update_web_acl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWebACLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_web_acl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWebACL(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("wafv2", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
