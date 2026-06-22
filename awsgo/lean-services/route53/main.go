package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53"
)

var fields_activate_key_signing_key = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_associate_vpc_with_hosted_zone = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "VPC", Flag: "vpc", Type: "*types.VPC", Required: true},
}

var fields_change_cidr_collection = []leanruntime.Field{
	{Name: "Changes", Flag: "changes", Type: "[]types.CidrCollectionChange", Required: true},
	{Name: "CollectionVersion", Flag: "collection-version", Type: "*int64", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_change_resource_record_sets = []leanruntime.Field{
	{Name: "ChangeBatch", Flag: "change-batch", Type: "*types.ChangeBatch", Required: true},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
}

var fields_change_tags_for_resource = []leanruntime.Field{
	{Name: "AddTags", Flag: "add-tags", Type: "[]types.Tag", Required: false},
	{Name: "RemoveTagKeys", Flag: "remove-tag-keys", Type: "[]string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.TagResourceType", Required: true},
}

var fields_create_cidr_collection = []leanruntime.Field{
	{Name: "CallerReference", Flag: "caller-reference", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_health_check = []leanruntime.Field{
	{Name: "CallerReference", Flag: "caller-reference", Type: "*string", Required: true},
	{Name: "HealthCheckConfig", Flag: "health-check-config", Type: "*types.HealthCheckConfig", Required: true},
}

var fields_create_hosted_zone = []leanruntime.Field{
	{Name: "CallerReference", Flag: "caller-reference", Type: "*string", Required: true},
	{Name: "DelegationSetId", Flag: "delegation-set-id", Type: "*string", Required: false},
	{Name: "HostedZoneConfig", Flag: "hosted-zone-config", Type: "*types.HostedZoneConfig", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VPC", Flag: "vpc", Type: "*types.VPC", Required: false},
}

var fields_create_key_signing_key = []leanruntime.Field{
	{Name: "CallerReference", Flag: "caller-reference", Type: "*string", Required: true},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "KeyManagementServiceArn", Flag: "key-management-service-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "*string", Required: true},
}

var fields_create_query_logging_config = []leanruntime.Field{
	{Name: "CloudWatchLogsLogGroupArn", Flag: "cloud-watch-logs-log-group-arn", Type: "*string", Required: true},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
}

var fields_create_reusable_delegation_set = []leanruntime.Field{
	{Name: "CallerReference", Flag: "caller-reference", Type: "*string", Required: true},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: false},
}

var fields_create_traffic_policy = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "Document", Flag: "document", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_traffic_policy_instance = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TTL", Flag: "ttl", Type: "*int64", Required: true},
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
	{Name: "TrafficPolicyVersion", Flag: "traffic-policy-version", Type: "*int32", Required: true},
}

var fields_create_traffic_policy_version = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "Document", Flag: "document", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_create_vpc_association_authorization = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "VPC", Flag: "vpc", Type: "*types.VPC", Required: true},
}

var fields_deactivate_key_signing_key = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_cidr_collection = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_health_check = []leanruntime.Field{
	{Name: "HealthCheckId", Flag: "health-check-id", Type: "*string", Required: true},
}

var fields_delete_hosted_zone = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_key_signing_key = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_query_logging_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_reusable_delegation_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_traffic_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*int32", Required: true},
}

var fields_delete_traffic_policy_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_vpc_association_authorization = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "VPC", Flag: "vpc", Type: "*types.VPC", Required: true},
}

var fields_disable_hosted_zone_dnssec = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
}

var fields_disassociate_vpc_from_hosted_zone = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "VPC", Flag: "vpc", Type: "*types.VPC", Required: true},
}

var fields_enable_hosted_zone_dnssec = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
}

var fields_get_account_limit = []leanruntime.Field{
	{Name: "Type", Flag: "type", Type: "types.AccountLimitType", Required: true},
}

var fields_get_change = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_checker_ip_ranges = []leanruntime.Field{}

var fields_get_dnssec = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
}

var fields_get_geo_location = []leanruntime.Field{
	{Name: "ContinentCode", Flag: "continent-code", Type: "*string", Required: false},
	{Name: "CountryCode", Flag: "country-code", Type: "*string", Required: false},
	{Name: "SubdivisionCode", Flag: "subdivision-code", Type: "*string", Required: false},
}

var fields_get_health_check = []leanruntime.Field{
	{Name: "HealthCheckId", Flag: "health-check-id", Type: "*string", Required: true},
}

var fields_get_health_check_count = []leanruntime.Field{}

var fields_get_health_check_last_failure_reason = []leanruntime.Field{
	{Name: "HealthCheckId", Flag: "health-check-id", Type: "*string", Required: true},
}

var fields_get_health_check_status = []leanruntime.Field{
	{Name: "HealthCheckId", Flag: "health-check-id", Type: "*string", Required: true},
}

var fields_get_hosted_zone = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_hosted_zone_count = []leanruntime.Field{}

var fields_get_hosted_zone_limit = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.HostedZoneLimitType", Required: true},
}

var fields_get_query_logging_config = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_reusable_delegation_set = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_reusable_delegation_set_limit = []leanruntime.Field{
	{Name: "DelegationSetId", Flag: "delegation-set-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ReusableDelegationSetLimitType", Required: true},
}

var fields_get_traffic_policy = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*int32", Required: true},
}

var fields_get_traffic_policy_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_traffic_policy_instance_count = []leanruntime.Field{}

var fields_list_cidr_blocks = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "LocationName", Flag: "location-name", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cidr_collections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_cidr_locations = []leanruntime.Field{
	{Name: "CollectionId", Flag: "collection-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_geo_locations = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "StartContinentCode", Flag: "start-continent-code", Type: "*string", Required: false},
	{Name: "StartCountryCode", Flag: "start-country-code", Type: "*string", Required: false},
	{Name: "StartSubdivisionCode", Flag: "start-subdivision-code", Type: "*string", Required: false},
}

var fields_list_health_checks = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_hosted_zones = []leanruntime.Field{
	{Name: "DelegationSetId", Flag: "delegation-set-id", Type: "*string", Required: false},
	{Name: "HostedZoneType", Flag: "hosted-zone-type", Type: "types.HostedZoneType", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_hosted_zones_by_name = []leanruntime.Field{
	{Name: "DNSName", Flag: "dns-name", Type: "*string", Required: false},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_hosted_zones_by_vpc = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VPCId", Flag: "vpcid", Type: "*string", Required: true},
	{Name: "VPCRegion", Flag: "vpc-region", Type: "types.VPCRegion", Required: true},
}

var fields_list_query_logging_configs = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_record_sets = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "StartRecordIdentifier", Flag: "start-record-identifier", Type: "*string", Required: false},
	{Name: "StartRecordName", Flag: "start-record-name", Type: "*string", Required: false},
	{Name: "StartRecordType", Flag: "start-record-type", Type: "types.RRType", Required: false},
}

var fields_list_reusable_delegation_sets = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.TagResourceType", Required: true},
}

var fields_list_tags_for_resources = []leanruntime.Field{
	{Name: "ResourceIds", Flag: "resource-ids", Type: "[]string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.TagResourceType", Required: true},
}

var fields_list_traffic_policies = []leanruntime.Field{
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "TrafficPolicyIdMarker", Flag: "traffic-policy-id-marker", Type: "*string", Required: false},
}

var fields_list_traffic_policy_instances = []leanruntime.Field{
	{Name: "HostedZoneIdMarker", Flag: "hosted-zone-id-marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "TrafficPolicyInstanceNameMarker", Flag: "traffic-policy-instance-name-marker", Type: "*string", Required: false},
	{Name: "TrafficPolicyInstanceTypeMarker", Flag: "traffic-policy-instance-type-marker", Type: "types.RRType", Required: false},
}

var fields_list_traffic_policy_instances_by_hosted_zone = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "TrafficPolicyInstanceNameMarker", Flag: "traffic-policy-instance-name-marker", Type: "*string", Required: false},
	{Name: "TrafficPolicyInstanceTypeMarker", Flag: "traffic-policy-instance-type-marker", Type: "types.RRType", Required: false},
}

var fields_list_traffic_policy_instances_by_policy = []leanruntime.Field{
	{Name: "HostedZoneIdMarker", Flag: "hosted-zone-id-marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
	{Name: "TrafficPolicyInstanceNameMarker", Flag: "traffic-policy-instance-name-marker", Type: "*string", Required: false},
	{Name: "TrafficPolicyInstanceTypeMarker", Flag: "traffic-policy-instance-type-marker", Type: "types.RRType", Required: false},
	{Name: "TrafficPolicyVersion", Flag: "traffic-policy-version", Type: "*int32", Required: true},
}

var fields_list_traffic_policy_versions = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "TrafficPolicyVersionMarker", Flag: "traffic-policy-version-marker", Type: "*string", Required: false},
}

var fields_list_vpc_association_authorizations = []leanruntime.Field{
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_test_dns_answer = []leanruntime.Field{
	{Name: "EDNS0ClientSubnetIP", Flag: "edns0-client-subnet-ip", Type: "*string", Required: false},
	{Name: "EDNS0ClientSubnetMask", Flag: "edns0-client-subnet-mask", Type: "*string", Required: false},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
	{Name: "RecordName", Flag: "record-name", Type: "*string", Required: true},
	{Name: "RecordType", Flag: "record-type", Type: "types.RRType", Required: true},
	{Name: "ResolverIP", Flag: "resolver-ip", Type: "*string", Required: false},
}

var fields_update_health_check = []leanruntime.Field{
	{Name: "AlarmIdentifier", Flag: "alarm-identifier", Type: "*types.AlarmIdentifier", Required: false},
	{Name: "ChildHealthChecks", Flag: "child-health-checks", Type: "[]string", Required: false},
	{Name: "Disabled", Flag: "disabled", Type: "*bool", Required: false},
	{Name: "EnableSNI", Flag: "enable-sni", Type: "*bool", Required: false},
	{Name: "FailureThreshold", Flag: "failure-threshold", Type: "*int32", Required: false},
	{Name: "FullyQualifiedDomainName", Flag: "fully-qualified-domain-name", Type: "*string", Required: false},
	{Name: "HealthCheckId", Flag: "health-check-id", Type: "*string", Required: true},
	{Name: "HealthCheckVersion", Flag: "health-check-version", Type: "*int64", Required: false},
	{Name: "HealthThreshold", Flag: "health-threshold", Type: "*int32", Required: false},
	{Name: "IPAddress", Flag: "ip-address", Type: "*string", Required: false},
	{Name: "InsufficientDataHealthStatus", Flag: "insufficient-data-health-status", Type: "types.InsufficientDataHealthStatus", Required: false},
	{Name: "Inverted", Flag: "inverted", Type: "*bool", Required: false},
	{Name: "Port", Flag: "port", Type: "*int32", Required: false},
	{Name: "Regions", Flag: "regions", Type: "[]types.HealthCheckRegion", Required: false},
	{Name: "ResetElements", Flag: "reset-elements", Type: "[]types.ResettableElementName", Required: false},
	{Name: "ResourcePath", Flag: "resource-path", Type: "*string", Required: false},
	{Name: "SearchString", Flag: "search-string", Type: "*string", Required: false},
}

var fields_update_hosted_zone_comment = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_hosted_zone_features = []leanruntime.Field{
	{Name: "EnableAcceleratedRecovery", Flag: "enable-accelerated-recovery", Type: "*bool", Required: false},
	{Name: "HostedZoneId", Flag: "hosted-zone-id", Type: "*string", Required: true},
}

var fields_update_traffic_policy_comment = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*int32", Required: true},
}

var fields_update_traffic_policy_instance = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "TTL", Flag: "ttl", Type: "*int64", Required: true},
	{Name: "TrafficPolicyId", Flag: "traffic-policy-id", Type: "*string", Required: true},
	{Name: "TrafficPolicyVersion", Flag: "traffic-policy-version", Type: "*int32", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"activate-key-signing-key": {
			Name:   "activate-key-signing-key",
			Fields: fields_activate_key_signing_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ActivateKeySigningKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_activate_key_signing_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ActivateKeySigningKey(ctx, input)
			},
		},
		"associate-vpc-with-hosted-zone": {
			Name:   "associate-vpc-with-hosted-zone",
			Fields: fields_associate_vpc_with_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateVPCWithHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_vpc_with_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateVPCWithHostedZone(ctx, input)
			},
		},
		"change-cidr-collection": {
			Name:   "change-cidr-collection",
			Fields: fields_change_cidr_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangeCidrCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_cidr_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangeCidrCollection(ctx, input)
			},
		},
		"change-resource-record-sets": {
			Name:   "change-resource-record-sets",
			Fields: fields_change_resource_record_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangeResourceRecordSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_resource_record_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangeResourceRecordSets(ctx, input)
			},
		},
		"change-tags-for-resource": {
			Name:   "change-tags-for-resource",
			Fields: fields_change_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ChangeTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_change_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ChangeTagsForResource(ctx, input)
			},
		},
		"create-cidr-collection": {
			Name:   "create-cidr-collection",
			Fields: fields_create_cidr_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCidrCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_cidr_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCidrCollection(ctx, input)
			},
		},
		"create-health-check": {
			Name:   "create-health-check",
			Fields: fields_create_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHealthCheck(ctx, input)
			},
		},
		"create-hosted-zone": {
			Name:   "create-hosted-zone",
			Fields: fields_create_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHostedZone(ctx, input)
			},
		},
		"create-key-signing-key": {
			Name:   "create-key-signing-key",
			Fields: fields_create_key_signing_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeySigningKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key_signing_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKeySigningKey(ctx, input)
			},
		},
		"create-query-logging-config": {
			Name:   "create-query-logging-config",
			Fields: fields_create_query_logging_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueryLoggingConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_query_logging_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueryLoggingConfig(ctx, input)
			},
		},
		"create-reusable-delegation-set": {
			Name:   "create-reusable-delegation-set",
			Fields: fields_create_reusable_delegation_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateReusableDelegationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_reusable_delegation_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateReusableDelegationSet(ctx, input)
			},
		},
		"create-traffic-policy": {
			Name:   "create-traffic-policy",
			Fields: fields_create_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficPolicy(ctx, input)
			},
		},
		"create-traffic-policy-instance": {
			Name:   "create-traffic-policy-instance",
			Fields: fields_create_traffic_policy_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficPolicyInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_policy_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficPolicyInstance(ctx, input)
			},
		},
		"create-traffic-policy-version": {
			Name:   "create-traffic-policy-version",
			Fields: fields_create_traffic_policy_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrafficPolicyVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_traffic_policy_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrafficPolicyVersion(ctx, input)
			},
		},
		"create-vpc-association-authorization": {
			Name:   "create-vpc-association-authorization",
			Fields: fields_create_vpc_association_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVPCAssociationAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_association_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVPCAssociationAuthorization(ctx, input)
			},
		},
		"deactivate-key-signing-key": {
			Name:   "deactivate-key-signing-key",
			Fields: fields_deactivate_key_signing_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeactivateKeySigningKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deactivate_key_signing_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeactivateKeySigningKey(ctx, input)
			},
		},
		"delete-cidr-collection": {
			Name:   "delete-cidr-collection",
			Fields: fields_delete_cidr_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCidrCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_cidr_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCidrCollection(ctx, input)
			},
		},
		"delete-health-check": {
			Name:   "delete-health-check",
			Fields: fields_delete_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHealthCheck(ctx, input)
			},
		},
		"delete-hosted-zone": {
			Name:   "delete-hosted-zone",
			Fields: fields_delete_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHostedZone(ctx, input)
			},
		},
		"delete-key-signing-key": {
			Name:   "delete-key-signing-key",
			Fields: fields_delete_key_signing_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeySigningKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key_signing_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKeySigningKey(ctx, input)
			},
		},
		"delete-query-logging-config": {
			Name:   "delete-query-logging-config",
			Fields: fields_delete_query_logging_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueryLoggingConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_query_logging_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueryLoggingConfig(ctx, input)
			},
		},
		"delete-reusable-delegation-set": {
			Name:   "delete-reusable-delegation-set",
			Fields: fields_delete_reusable_delegation_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteReusableDelegationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_reusable_delegation_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteReusableDelegationSet(ctx, input)
			},
		},
		"delete-traffic-policy": {
			Name:   "delete-traffic-policy",
			Fields: fields_delete_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficPolicy(ctx, input)
			},
		},
		"delete-traffic-policy-instance": {
			Name:   "delete-traffic-policy-instance",
			Fields: fields_delete_traffic_policy_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrafficPolicyInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_traffic_policy_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrafficPolicyInstance(ctx, input)
			},
		},
		"delete-vpc-association-authorization": {
			Name:   "delete-vpc-association-authorization",
			Fields: fields_delete_vpc_association_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVPCAssociationAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_association_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVPCAssociationAuthorization(ctx, input)
			},
		},
		"disable-hosted-zone-dnssec": {
			Name:   "disable-hosted-zone-dnssec",
			Fields: fields_disable_hosted_zone_dnssec,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableHostedZoneDNSSECInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_hosted_zone_dnssec, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableHostedZoneDNSSEC(ctx, input)
			},
		},
		"disassociate-vpc-from-hosted-zone": {
			Name:   "disassociate-vpc-from-hosted-zone",
			Fields: fields_disassociate_vpc_from_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateVPCFromHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_vpc_from_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateVPCFromHostedZone(ctx, input)
			},
		},
		"enable-hosted-zone-dnssec": {
			Name:   "enable-hosted-zone-dnssec",
			Fields: fields_enable_hosted_zone_dnssec,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableHostedZoneDNSSECInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_hosted_zone_dnssec, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableHostedZoneDNSSEC(ctx, input)
			},
		},
		"get-account-limit": {
			Name:   "get-account-limit",
			Fields: fields_get_account_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountLimit(ctx, input)
			},
		},
		"get-change": {
			Name:   "get-change",
			Fields: fields_get_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetChange(ctx, input)
			},
		},
		"get-checker-ip-ranges": {
			Name:   "get-checker-ip-ranges",
			Fields: fields_get_checker_ip_ranges,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCheckerIpRangesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_checker_ip_ranges, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCheckerIpRanges(ctx, input)
			},
		},
		"get-dnssec": {
			Name:   "get-dnssec",
			Fields: fields_get_dnssec,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDNSSECInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_dnssec, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDNSSEC(ctx, input)
			},
		},
		"get-geo-location": {
			Name:   "get-geo-location",
			Fields: fields_get_geo_location,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGeoLocationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_geo_location, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGeoLocation(ctx, input)
			},
		},
		"get-health-check": {
			Name:   "get-health-check",
			Fields: fields_get_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHealthCheck(ctx, input)
			},
		},
		"get-health-check-count": {
			Name:   "get-health-check-count",
			Fields: fields_get_health_check_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHealthCheckCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_health_check_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHealthCheckCount(ctx, input)
			},
		},
		"get-health-check-last-failure-reason": {
			Name:   "get-health-check-last-failure-reason",
			Fields: fields_get_health_check_last_failure_reason,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHealthCheckLastFailureReasonInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_health_check_last_failure_reason, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHealthCheckLastFailureReason(ctx, input)
			},
		},
		"get-health-check-status": {
			Name:   "get-health-check-status",
			Fields: fields_get_health_check_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHealthCheckStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_health_check_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHealthCheckStatus(ctx, input)
			},
		},
		"get-hosted-zone": {
			Name:   "get-hosted-zone",
			Fields: fields_get_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHostedZone(ctx, input)
			},
		},
		"get-hosted-zone-count": {
			Name:   "get-hosted-zone-count",
			Fields: fields_get_hosted_zone_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostedZoneCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hosted_zone_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHostedZoneCount(ctx, input)
			},
		},
		"get-hosted-zone-limit": {
			Name:   "get-hosted-zone-limit",
			Fields: fields_get_hosted_zone_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHostedZoneLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hosted_zone_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHostedZoneLimit(ctx, input)
			},
		},
		"get-query-logging-config": {
			Name:   "get-query-logging-config",
			Fields: fields_get_query_logging_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueryLoggingConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_query_logging_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueryLoggingConfig(ctx, input)
			},
		},
		"get-reusable-delegation-set": {
			Name:   "get-reusable-delegation-set",
			Fields: fields_get_reusable_delegation_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReusableDelegationSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reusable_delegation_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReusableDelegationSet(ctx, input)
			},
		},
		"get-reusable-delegation-set-limit": {
			Name:   "get-reusable-delegation-set-limit",
			Fields: fields_get_reusable_delegation_set_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetReusableDelegationSetLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_reusable_delegation_set_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetReusableDelegationSetLimit(ctx, input)
			},
		},
		"get-traffic-policy": {
			Name:   "get-traffic-policy",
			Fields: fields_get_traffic_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrafficPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_traffic_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrafficPolicy(ctx, input)
			},
		},
		"get-traffic-policy-instance": {
			Name:   "get-traffic-policy-instance",
			Fields: fields_get_traffic_policy_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrafficPolicyInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_traffic_policy_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrafficPolicyInstance(ctx, input)
			},
		},
		"get-traffic-policy-instance-count": {
			Name:   "get-traffic-policy-instance-count",
			Fields: fields_get_traffic_policy_instance_count,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrafficPolicyInstanceCountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_traffic_policy_instance_count, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrafficPolicyInstanceCount(ctx, input)
			},
		},
		"list-cidr-blocks": {
			Name:   "list-cidr-blocks",
			Fields: fields_list_cidr_blocks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCidrBlocksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cidr_blocks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCidrBlocks(ctx, input)
				}
				var results []*svc.ListCidrBlocksOutput
				p := svc.NewListCidrBlocksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-cidr-collections": {
			Name:   "list-cidr-collections",
			Fields: fields_list_cidr_collections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCidrCollectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cidr_collections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCidrCollections(ctx, input)
				}
				var results []*svc.ListCidrCollectionsOutput
				p := svc.NewListCidrCollectionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-cidr-locations": {
			Name:   "list-cidr-locations",
			Fields: fields_list_cidr_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCidrLocationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_cidr_locations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCidrLocations(ctx, input)
				}
				var results []*svc.ListCidrLocationsOutput
				p := svc.NewListCidrLocationsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-geo-locations": {
			Name:   "list-geo-locations",
			Fields: fields_list_geo_locations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGeoLocationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_geo_locations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListGeoLocations(ctx, input)
			},
		},
		"list-health-checks": {
			Name:   "list-health-checks",
			Fields: fields_list_health_checks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHealthChecksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_health_checks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHealthChecks(ctx, input)
				}
				var results []*svc.ListHealthChecksOutput
				p := svc.NewListHealthChecksPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-hosted-zones": {
			Name:   "list-hosted-zones",
			Fields: fields_list_hosted_zones,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostedZonesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hosted_zones, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHostedZones(ctx, input)
				}
				var results []*svc.ListHostedZonesOutput
				p := svc.NewListHostedZonesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-hosted-zones-by-name": {
			Name:   "list-hosted-zones-by-name",
			Fields: fields_list_hosted_zones_by_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostedZonesByNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hosted_zones_by_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHostedZonesByName(ctx, input)
			},
		},
		"list-hosted-zones-by-vpc": {
			Name:   "list-hosted-zones-by-vpc",
			Fields: fields_list_hosted_zones_by_vpc,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHostedZonesByVPCInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_hosted_zones_by_vpc, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListHostedZonesByVPC(ctx, input)
			},
		},
		"list-query-logging-configs": {
			Name:   "list-query-logging-configs",
			Fields: fields_list_query_logging_configs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueryLoggingConfigsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_query_logging_configs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueryLoggingConfigs(ctx, input)
				}
				var results []*svc.ListQueryLoggingConfigsOutput
				p := svc.NewListQueryLoggingConfigsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-resource-record-sets": {
			Name:   "list-resource-record-sets",
			Fields: fields_list_resource_record_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceRecordSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resource_record_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourceRecordSets(ctx, input)
			},
		},
		"list-reusable-delegation-sets": {
			Name:   "list-reusable-delegation-sets",
			Fields: fields_list_reusable_delegation_sets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReusableDelegationSetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_reusable_delegation_sets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReusableDelegationSets(ctx, input)
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
		"list-tags-for-resources": {
			Name:   "list-tags-for-resources",
			Fields: fields_list_tags_for_resources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourcesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resources, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResources(ctx, input)
			},
		},
		"list-traffic-policies": {
			Name:   "list-traffic-policies",
			Fields: fields_list_traffic_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficPoliciesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_traffic_policies, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTrafficPolicies(ctx, input)
			},
		},
		"list-traffic-policy-instances": {
			Name:   "list-traffic-policy-instances",
			Fields: fields_list_traffic_policy_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficPolicyInstancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_traffic_policy_instances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTrafficPolicyInstances(ctx, input)
			},
		},
		"list-traffic-policy-instances-by-hosted-zone": {
			Name:   "list-traffic-policy-instances-by-hosted-zone",
			Fields: fields_list_traffic_policy_instances_by_hosted_zone,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficPolicyInstancesByHostedZoneInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_traffic_policy_instances_by_hosted_zone, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTrafficPolicyInstancesByHostedZone(ctx, input)
			},
		},
		"list-traffic-policy-instances-by-policy": {
			Name:   "list-traffic-policy-instances-by-policy",
			Fields: fields_list_traffic_policy_instances_by_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficPolicyInstancesByPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_traffic_policy_instances_by_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTrafficPolicyInstancesByPolicy(ctx, input)
			},
		},
		"list-traffic-policy-versions": {
			Name:   "list-traffic-policy-versions",
			Fields: fields_list_traffic_policy_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrafficPolicyVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_traffic_policy_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTrafficPolicyVersions(ctx, input)
			},
		},
		"list-vpc-association-authorizations": {
			Name:   "list-vpc-association-authorizations",
			Fields: fields_list_vpc_association_authorizations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVPCAssociationAuthorizationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_association_authorizations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVPCAssociationAuthorizations(ctx, input)
			},
		},
		"test-dns-answer": {
			Name:   "test-dns-answer",
			Fields: fields_test_dns_answer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestDNSAnswerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_dns_answer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestDNSAnswer(ctx, input)
			},
		},
		"update-health-check": {
			Name:   "update-health-check",
			Fields: fields_update_health_check,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHealthCheckInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_health_check, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHealthCheck(ctx, input)
			},
		},
		"update-hosted-zone-comment": {
			Name:   "update-hosted-zone-comment",
			Fields: fields_update_hosted_zone_comment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHostedZoneCommentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hosted_zone_comment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHostedZoneComment(ctx, input)
			},
		},
		"update-hosted-zone-features": {
			Name:   "update-hosted-zone-features",
			Fields: fields_update_hosted_zone_features,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHostedZoneFeaturesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hosted_zone_features, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHostedZoneFeatures(ctx, input)
			},
		},
		"update-traffic-policy-comment": {
			Name:   "update-traffic-policy-comment",
			Fields: fields_update_traffic_policy_comment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrafficPolicyCommentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_traffic_policy_comment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrafficPolicyComment(ctx, input)
			},
		},
		"update-traffic-policy-instance": {
			Name:   "update-traffic-policy-instance",
			Fields: fields_update_traffic_policy_instance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrafficPolicyInstanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_traffic_policy_instance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrafficPolicyInstance(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("route53", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
