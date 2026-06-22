package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/wafregional"
)

var fields_associate_web_acl = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_create_byte_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_geo_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_ip_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_rate_based_rule = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RateKey", Flag: "rate-key", Type: "types.RateKey", Required: true},
	{Name: "RateLimit", Flag: "rate-limit", Type: "*int64", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_regex_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_regex_pattern_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_rule = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_rule_group = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_size_constraint_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_sql_injection_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_web_acl = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "DefaultAction", Flag: "default-action", Type: "*types.WafAction", Required: true},
	{Name: "MetricName", Flag: "metric-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_web_acl_migration_stack = []leanruntime.Field{
	{Name: "IgnoreUnsupportedType", Flag: "ignore-unsupported-type", Type: "*bool", Required: true},
	{Name: "S3BucketName", Flag: "s3-bucket-name", Type: "*string", Required: true},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_create_xss_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_byte_match_set = []leanruntime.Field{
	{Name: "ByteMatchSetId", Flag: "byte-match-set-id", Type: "*string", Required: true},
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
}

var fields_delete_geo_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "GeoMatchSetId", Flag: "geo-match-set-id", Type: "*string", Required: true},
}

var fields_delete_ip_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "IPSetId", Flag: "ip-set-id", Type: "*string", Required: true},
}

var fields_delete_logging_configuration = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_permission_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_delete_rate_based_rule = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_delete_regex_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RegexMatchSetId", Flag: "regex-match-set-id", Type: "*string", Required: true},
}

var fields_delete_regex_pattern_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RegexPatternSetId", Flag: "regex-pattern-set-id", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_delete_rule_group = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RuleGroupId", Flag: "rule-group-id", Type: "*string", Required: true},
}

var fields_delete_size_constraint_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "SizeConstraintSetId", Flag: "size-constraint-set-id", Type: "*string", Required: true},
}

var fields_delete_sql_injection_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "SqlInjectionMatchSetId", Flag: "sql-injection-match-set-id", Type: "*string", Required: true},
}

var fields_delete_web_acl = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_delete_xss_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "XssMatchSetId", Flag: "xss-match-set-id", Type: "*string", Required: true},
}

var fields_disassociate_web_acl = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_byte_match_set = []leanruntime.Field{
	{Name: "ByteMatchSetId", Flag: "byte-match-set-id", Type: "*string", Required: true},
}

var fields_get_change_token = []leanruntime.Field{}

var fields_get_change_token_status = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
}

var fields_get_geo_match_set = []leanruntime.Field{
	{Name: "GeoMatchSetId", Flag: "geo-match-set-id", Type: "*string", Required: true},
}

var fields_get_ip_set = []leanruntime.Field{
	{Name: "IPSetId", Flag: "ip-set-id", Type: "*string", Required: true},
}

var fields_get_logging_configuration = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_permission_policy = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_rate_based_rule = []leanruntime.Field{
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_get_rate_based_rule_managed_keys = []leanruntime.Field{
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_get_regex_match_set = []leanruntime.Field{
	{Name: "RegexMatchSetId", Flag: "regex-match-set-id", Type: "*string", Required: true},
}

var fields_get_regex_pattern_set = []leanruntime.Field{
	{Name: "RegexPatternSetId", Flag: "regex-pattern-set-id", Type: "*string", Required: true},
}

var fields_get_rule = []leanruntime.Field{
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
}

var fields_get_rule_group = []leanruntime.Field{
	{Name: "RuleGroupId", Flag: "rule-group-id", Type: "*string", Required: true},
}

var fields_get_sampled_requests = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int64", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
	{Name: "TimeWindow", Flag: "time-window", Type: "*types.TimeWindow", Required: true},
	{Name: "WebAclId", Flag: "web-acl-id", Type: "*string", Required: true},
}

var fields_get_size_constraint_set = []leanruntime.Field{
	{Name: "SizeConstraintSetId", Flag: "size-constraint-set-id", Type: "*string", Required: true},
}

var fields_get_sql_injection_match_set = []leanruntime.Field{
	{Name: "SqlInjectionMatchSetId", Flag: "sql-injection-match-set-id", Type: "*string", Required: true},
}

var fields_get_web_acl = []leanruntime.Field{
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_get_web_acl_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_get_xss_match_set = []leanruntime.Field{
	{Name: "XssMatchSetId", Flag: "xss-match-set-id", Type: "*string", Required: true},
}

var fields_list_activated_rules_in_rule_group = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "RuleGroupId", Flag: "rule-group-id", Type: "*string", Required: false},
}

var fields_list_byte_match_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_geo_match_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_ip_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_logging_configurations = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_rate_based_rules = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_regex_match_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_regex_pattern_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_resources_for_web_acl = []leanruntime.Field{
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: false},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_list_rule_groups = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_rules = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_size_constraint_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_sql_injection_match_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_subscribed_rule_groups = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_web_acls = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_list_xss_match_sets = []leanruntime.Field{
	{Name: "Limit", Flag: "limit", Type: "int32", Required: false},
	{Name: "NextMarker", Flag: "next-marker", Type: "*string", Required: false},
}

var fields_put_logging_configuration = []leanruntime.Field{
	{Name: "LoggingConfiguration", Flag: "logging-configuration", Type: "*types.LoggingConfiguration", Required: true},
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

var fields_update_byte_match_set = []leanruntime.Field{
	{Name: "ByteMatchSetId", Flag: "byte-match-set-id", Type: "*string", Required: true},
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.ByteMatchSetUpdate", Required: true},
}

var fields_update_geo_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "GeoMatchSetId", Flag: "geo-match-set-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.GeoMatchSetUpdate", Required: true},
}

var fields_update_ip_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "IPSetId", Flag: "ip-set-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.IPSetUpdate", Required: true},
}

var fields_update_rate_based_rule = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RateLimit", Flag: "rate-limit", Type: "*int64", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.RuleUpdate", Required: true},
}

var fields_update_regex_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RegexMatchSetId", Flag: "regex-match-set-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.RegexMatchSetUpdate", Required: true},
}

var fields_update_regex_pattern_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RegexPatternSetId", Flag: "regex-pattern-set-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.RegexPatternSetUpdate", Required: true},
}

var fields_update_rule = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RuleId", Flag: "rule-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.RuleUpdate", Required: true},
}

var fields_update_rule_group = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "RuleGroupId", Flag: "rule-group-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.RuleGroupUpdate", Required: true},
}

var fields_update_size_constraint_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "SizeConstraintSetId", Flag: "size-constraint-set-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.SizeConstraintSetUpdate", Required: true},
}

var fields_update_sql_injection_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "SqlInjectionMatchSetId", Flag: "sql-injection-match-set-id", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.SqlInjectionMatchSetUpdate", Required: true},
}

var fields_update_web_acl = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "DefaultAction", Flag: "default-action", Type: "*types.WafAction", Required: false},
	{Name: "Updates", Flag: "updates", Type: "[]types.WebACLUpdate", Required: false},
	{Name: "WebACLId", Flag: "web-aclid", Type: "*string", Required: true},
}

var fields_update_xss_match_set = []leanruntime.Field{
	{Name: "ChangeToken", Flag: "change-token", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.XssMatchSetUpdate", Required: true},
	{Name: "XssMatchSetId", Flag: "xss-match-set-id", Type: "*string", Required: true},
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
		"create-byte-match-set": {
			Name:   "create-byte-match-set",
			Fields: fields_create_byte_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateByteMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_byte_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateByteMatchSet(ctx, input)
			},
		},
		"create-geo-match-set": {
			Name:   "create-geo-match-set",
			Fields: fields_create_geo_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGeoMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_geo_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGeoMatchSet(ctx, input)
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
		"create-rate-based-rule": {
			Name:   "create-rate-based-rule",
			Fields: fields_create_rate_based_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRateBasedRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rate_based_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRateBasedRule(ctx, input)
			},
		},
		"create-regex-match-set": {
			Name:   "create-regex-match-set",
			Fields: fields_create_regex_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRegexMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_regex_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRegexMatchSet(ctx, input)
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
		"create-rule": {
			Name:   "create-rule",
			Fields: fields_create_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRule(ctx, input)
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
		"create-size-constraint-set": {
			Name:   "create-size-constraint-set",
			Fields: fields_create_size_constraint_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSizeConstraintSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_size_constraint_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSizeConstraintSet(ctx, input)
			},
		},
		"create-sql-injection-match-set": {
			Name:   "create-sql-injection-match-set",
			Fields: fields_create_sql_injection_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSqlInjectionMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_sql_injection_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSqlInjectionMatchSet(ctx, input)
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
		"create-web-acl-migration-stack": {
			Name:   "create-web-acl-migration-stack",
			Fields: fields_create_web_acl_migration_stack,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWebACLMigrationStackInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_web_acl_migration_stack, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWebACLMigrationStack(ctx, input)
			},
		},
		"create-xss-match-set": {
			Name:   "create-xss-match-set",
			Fields: fields_create_xss_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateXssMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_xss_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateXssMatchSet(ctx, input)
			},
		},
		"delete-byte-match-set": {
			Name:   "delete-byte-match-set",
			Fields: fields_delete_byte_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteByteMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_byte_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteByteMatchSet(ctx, input)
			},
		},
		"delete-geo-match-set": {
			Name:   "delete-geo-match-set",
			Fields: fields_delete_geo_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGeoMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_geo_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGeoMatchSet(ctx, input)
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
		"delete-rate-based-rule": {
			Name:   "delete-rate-based-rule",
			Fields: fields_delete_rate_based_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRateBasedRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rate_based_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRateBasedRule(ctx, input)
			},
		},
		"delete-regex-match-set": {
			Name:   "delete-regex-match-set",
			Fields: fields_delete_regex_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRegexMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_regex_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRegexMatchSet(ctx, input)
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
		"delete-rule": {
			Name:   "delete-rule",
			Fields: fields_delete_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRule(ctx, input)
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
		"delete-size-constraint-set": {
			Name:   "delete-size-constraint-set",
			Fields: fields_delete_size_constraint_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSizeConstraintSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_size_constraint_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSizeConstraintSet(ctx, input)
			},
		},
		"delete-sql-injection-match-set": {
			Name:   "delete-sql-injection-match-set",
			Fields: fields_delete_sql_injection_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSqlInjectionMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_sql_injection_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSqlInjectionMatchSet(ctx, input)
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
		"delete-xss-match-set": {
			Name:   "delete-xss-match-set",
			Fields: fields_delete_xss_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteXssMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_xss_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteXssMatchSet(ctx, input)
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
		"get-byte-match-set": {
			Name:   "get-byte-match-set",
			Fields: fields_get_byte_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetByteMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_byte_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetByteMatchSet(ctx, input)
			},
		},
		"get-change-token": {
			Name:   "get-change-token",
			Fields: fields_get_change_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChangeTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_change_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChangeToken(ctx, input)
			},
		},
		"get-change-token-status": {
			Name:   "get-change-token-status",
			Fields: fields_get_change_token_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChangeTokenStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_change_token_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChangeTokenStatus(ctx, input)
			},
		},
		"get-geo-match-set": {
			Name:   "get-geo-match-set",
			Fields: fields_get_geo_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGeoMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_geo_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGeoMatchSet(ctx, input)
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
		"get-rate-based-rule": {
			Name:   "get-rate-based-rule",
			Fields: fields_get_rate_based_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRateBasedRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rate_based_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRateBasedRule(ctx, input)
			},
		},
		"get-rate-based-rule-managed-keys": {
			Name:   "get-rate-based-rule-managed-keys",
			Fields: fields_get_rate_based_rule_managed_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRateBasedRuleManagedKeysInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rate_based_rule_managed_keys, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRateBasedRuleManagedKeys(ctx, input)
			},
		},
		"get-regex-match-set": {
			Name:   "get-regex-match-set",
			Fields: fields_get_regex_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegexMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_regex_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegexMatchSet(ctx, input)
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
		"get-rule": {
			Name:   "get-rule",
			Fields: fields_get_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRule(ctx, input)
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
		"get-size-constraint-set": {
			Name:   "get-size-constraint-set",
			Fields: fields_get_size_constraint_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSizeConstraintSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_size_constraint_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSizeConstraintSet(ctx, input)
			},
		},
		"get-sql-injection-match-set": {
			Name:   "get-sql-injection-match-set",
			Fields: fields_get_sql_injection_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSqlInjectionMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_sql_injection_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSqlInjectionMatchSet(ctx, input)
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
		"get-xss-match-set": {
			Name:   "get-xss-match-set",
			Fields: fields_get_xss_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetXssMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_xss_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetXssMatchSet(ctx, input)
			},
		},
		"list-activated-rules-in-rule-group": {
			Name:   "list-activated-rules-in-rule-group",
			Fields: fields_list_activated_rules_in_rule_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListActivatedRulesInRuleGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_activated_rules_in_rule_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListActivatedRulesInRuleGroup(ctx, input)
			},
		},
		"list-byte-match-sets": {
			Name:   "list-byte-match-sets",
			Fields: fields_list_byte_match_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListByteMatchSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_byte_match_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListByteMatchSets(ctx, input)
			},
		},
		"list-geo-match-sets": {
			Name:   "list-geo-match-sets",
			Fields: fields_list_geo_match_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGeoMatchSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_geo_match_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGeoMatchSets(ctx, input)
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
		"list-rate-based-rules": {
			Name:   "list-rate-based-rules",
			Fields: fields_list_rate_based_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRateBasedRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rate_based_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRateBasedRules(ctx, input)
			},
		},
		"list-regex-match-sets": {
			Name:   "list-regex-match-sets",
			Fields: fields_list_regex_match_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRegexMatchSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_regex_match_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRegexMatchSets(ctx, input)
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
		"list-rules": {
			Name:   "list-rules",
			Fields: fields_list_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRules(ctx, input)
			},
		},
		"list-size-constraint-sets": {
			Name:   "list-size-constraint-sets",
			Fields: fields_list_size_constraint_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSizeConstraintSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_size_constraint_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSizeConstraintSets(ctx, input)
			},
		},
		"list-sql-injection-match-sets": {
			Name:   "list-sql-injection-match-sets",
			Fields: fields_list_sql_injection_match_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSqlInjectionMatchSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_sql_injection_match_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSqlInjectionMatchSets(ctx, input)
			},
		},
		"list-subscribed-rule-groups": {
			Name:   "list-subscribed-rule-groups",
			Fields: fields_list_subscribed_rule_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscribedRuleGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_subscribed_rule_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListSubscribedRuleGroups(ctx, input)
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
		"list-xss-match-sets": {
			Name:   "list-xss-match-sets",
			Fields: fields_list_xss_match_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListXssMatchSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_xss_match_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListXssMatchSets(ctx, input)
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
		"update-byte-match-set": {
			Name:   "update-byte-match-set",
			Fields: fields_update_byte_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateByteMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_byte_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateByteMatchSet(ctx, input)
			},
		},
		"update-geo-match-set": {
			Name:   "update-geo-match-set",
			Fields: fields_update_geo_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGeoMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_geo_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGeoMatchSet(ctx, input)
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
		"update-rate-based-rule": {
			Name:   "update-rate-based-rule",
			Fields: fields_update_rate_based_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRateBasedRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rate_based_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRateBasedRule(ctx, input)
			},
		},
		"update-regex-match-set": {
			Name:   "update-regex-match-set",
			Fields: fields_update_regex_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRegexMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_regex_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRegexMatchSet(ctx, input)
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
		"update-rule": {
			Name:   "update-rule",
			Fields: fields_update_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRule(ctx, input)
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
		"update-size-constraint-set": {
			Name:   "update-size-constraint-set",
			Fields: fields_update_size_constraint_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSizeConstraintSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_size_constraint_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSizeConstraintSet(ctx, input)
			},
		},
		"update-sql-injection-match-set": {
			Name:   "update-sql-injection-match-set",
			Fields: fields_update_sql_injection_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSqlInjectionMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_sql_injection_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSqlInjectionMatchSet(ctx, input)
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
		"update-xss-match-set": {
			Name:   "update-xss-match-set",
			Fields: fields_update_xss_match_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateXssMatchSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_xss_match_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateXssMatchSet(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("wafregional", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
