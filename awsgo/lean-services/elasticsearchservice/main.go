package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
)

var fields_accept_inbound_cross_cluster_search_connection = []leanruntime.Field{
	{Name: "CrossClusterSearchConnectionId", Flag: "cross-cluster-search-connection-id", Type: "*string", Required: true},
}

var fields_add_tags = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: true},
}

var fields_associate_package = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_authorize_vpc_endpoint_access = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_cancel_domain_config_change = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
}

var fields_cancel_elasticsearch_service_software_update = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_create_elasticsearch_domain = []leanruntime.Field{
	{Name: "AccessPolicies", Flag: "access-policies", Type: "*string", Required: false},
	{Name: "AdvancedOptions", Flag: "advanced-options", Type: "map[string]string", Required: false},
	{Name: "AdvancedSecurityOptions", Flag: "advanced-security-options", Type: "*types.AdvancedSecurityOptionsInput", Required: false},
	{Name: "AutoTuneOptions", Flag: "auto-tune-options", Type: "*types.AutoTuneOptionsInput", Required: false},
	{Name: "CognitoOptions", Flag: "cognito-options", Type: "*types.CognitoOptions", Required: false},
	{Name: "DomainEndpointOptions", Flag: "domain-endpoint-options", Type: "*types.DomainEndpointOptions", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "EBSOptions", Flag: "ebs-options", Type: "*types.EBSOptions", Required: false},
	{Name: "ElasticsearchClusterConfig", Flag: "elasticsearch-cluster-config", Type: "*types.ElasticsearchClusterConfig", Required: false},
	{Name: "ElasticsearchVersion", Flag: "elasticsearch-version", Type: "*string", Required: false},
	{Name: "EncryptionAtRestOptions", Flag: "encryption-at-rest-options", Type: "*types.EncryptionAtRestOptions", Required: false},
	{Name: "LogPublishingOptions", Flag: "log-publishing-options", Type: "map[string]types.LogPublishingOption", Required: false},
	{Name: "NodeToNodeEncryptionOptions", Flag: "node-to-node-encryption-options", Type: "*types.NodeToNodeEncryptionOptions", Required: false},
	{Name: "SnapshotOptions", Flag: "snapshot-options", Type: "*types.SnapshotOptions", Required: false},
	{Name: "TagList", Flag: "tag-list", Type: "[]types.Tag", Required: false},
	{Name: "VPCOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: false},
}

var fields_create_outbound_cross_cluster_search_connection = []leanruntime.Field{
	{Name: "ConnectionAlias", Flag: "connection-alias", Type: "*string", Required: true},
	{Name: "DestinationDomainInfo", Flag: "destination-domain-info", Type: "*types.DomainInformation", Required: true},
	{Name: "SourceDomainInfo", Flag: "source-domain-info", Type: "*types.DomainInformation", Required: true},
}

var fields_create_package = []leanruntime.Field{
	{Name: "PackageDescription", Flag: "package-description", Type: "*string", Required: false},
	{Name: "PackageName", Flag: "package-name", Type: "*string", Required: true},
	{Name: "PackageSource", Flag: "package-source", Type: "*types.PackageSource", Required: true},
	{Name: "PackageType", Flag: "package-type", Type: "types.PackageType", Required: true},
}

var fields_create_vpc_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainArn", Flag: "domain-arn", Type: "*string", Required: true},
	{Name: "VpcOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: true},
}

var fields_delete_elasticsearch_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_elasticsearch_service_role = []leanruntime.Field{}

var fields_delete_inbound_cross_cluster_search_connection = []leanruntime.Field{
	{Name: "CrossClusterSearchConnectionId", Flag: "cross-cluster-search-connection-id", Type: "*string", Required: true},
}

var fields_delete_outbound_cross_cluster_search_connection = []leanruntime.Field{
	{Name: "CrossClusterSearchConnectionId", Flag: "cross-cluster-search-connection-id", Type: "*string", Required: true},
}

var fields_delete_package = []leanruntime.Field{
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_delete_vpc_endpoint = []leanruntime.Field{
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: true},
}

var fields_describe_domain_auto_tunes = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_domain_change_progress = []leanruntime.Field{
	{Name: "ChangeId", Flag: "change-id", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_elasticsearch_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_elasticsearch_domain_config = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_describe_elasticsearch_domains = []leanruntime.Field{
	{Name: "DomainNames", Flag: "domain-names", Type: "[]string", Required: true},
}

var fields_describe_elasticsearch_instance_type_limits = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "ElasticsearchVersion", Flag: "elasticsearch-version", Type: "*string", Required: true},
	{Name: "InstanceType", Flag: "instance-type", Type: "types.ESPartitionInstanceType", Required: true},
}

var fields_describe_inbound_cross_cluster_search_connections = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_outbound_cross_cluster_search_connections = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_packages = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.DescribePackagesFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_reserved_elasticsearch_instance_offerings = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReservedElasticsearchInstanceOfferingId", Flag: "reserved-elasticsearch-instance-offering-id", Type: "*string", Required: false},
}

var fields_describe_reserved_elasticsearch_instances = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReservedElasticsearchInstanceId", Flag: "reserved-elasticsearch-instance-id", Type: "*string", Required: false},
}

var fields_describe_vpc_endpoints = []leanruntime.Field{
	{Name: "VpcEndpointIds", Flag: "vpc-endpoint-ids", Type: "[]string", Required: true},
}

var fields_dissociate_package = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_get_compatible_elasticsearch_versions = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
}

var fields_get_package_version_history = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_get_upgrade_history = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_upgrade_status = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_list_domain_names = []leanruntime.Field{
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: false},
}

var fields_list_domains_for_package = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
}

var fields_list_elasticsearch_instance_types = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
	{Name: "ElasticsearchVersion", Flag: "elasticsearch-version", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_elasticsearch_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_packages_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: true},
}

var fields_list_vpc_endpoint_access = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_endpoints = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vpc_endpoints_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_purchase_reserved_elasticsearch_instance_offering = []leanruntime.Field{
	{Name: "InstanceCount", Flag: "instance-count", Type: "*int32", Required: false},
	{Name: "ReservationName", Flag: "reservation-name", Type: "*string", Required: true},
	{Name: "ReservedElasticsearchInstanceOfferingId", Flag: "reserved-elasticsearch-instance-offering-id", Type: "*string", Required: true},
}

var fields_reject_inbound_cross_cluster_search_connection = []leanruntime.Field{
	{Name: "CrossClusterSearchConnectionId", Flag: "cross-cluster-search-connection-id", Type: "*string", Required: true},
}

var fields_remove_tags = []leanruntime.Field{
	{Name: "ARN", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_revoke_vpc_endpoint_access = []leanruntime.Field{
	{Name: "Account", Flag: "account", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_start_elasticsearch_service_software_update = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_update_elasticsearch_domain_config = []leanruntime.Field{
	{Name: "AccessPolicies", Flag: "access-policies", Type: "*string", Required: false},
	{Name: "AdvancedOptions", Flag: "advanced-options", Type: "map[string]string", Required: false},
	{Name: "AdvancedSecurityOptions", Flag: "advanced-security-options", Type: "*types.AdvancedSecurityOptionsInput", Required: false},
	{Name: "AutoTuneOptions", Flag: "auto-tune-options", Type: "*types.AutoTuneOptions", Required: false},
	{Name: "CognitoOptions", Flag: "cognito-options", Type: "*types.CognitoOptions", Required: false},
	{Name: "DomainEndpointOptions", Flag: "domain-endpoint-options", Type: "*types.DomainEndpointOptions", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DryRun", Flag: "dry-run", Type: "*bool", Required: false},
	{Name: "EBSOptions", Flag: "ebs-options", Type: "*types.EBSOptions", Required: false},
	{Name: "ElasticsearchClusterConfig", Flag: "elasticsearch-cluster-config", Type: "*types.ElasticsearchClusterConfig", Required: false},
	{Name: "EncryptionAtRestOptions", Flag: "encryption-at-rest-options", Type: "*types.EncryptionAtRestOptions", Required: false},
	{Name: "LogPublishingOptions", Flag: "log-publishing-options", Type: "map[string]types.LogPublishingOption", Required: false},
	{Name: "NodeToNodeEncryptionOptions", Flag: "node-to-node-encryption-options", Type: "*types.NodeToNodeEncryptionOptions", Required: false},
	{Name: "SnapshotOptions", Flag: "snapshot-options", Type: "*types.SnapshotOptions", Required: false},
	{Name: "VPCOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: false},
}

var fields_update_package = []leanruntime.Field{
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "PackageDescription", Flag: "package-description", Type: "*string", Required: false},
	{Name: "PackageID", Flag: "package-id", Type: "*string", Required: true},
	{Name: "PackageSource", Flag: "package-source", Type: "*types.PackageSource", Required: true},
}

var fields_update_vpc_endpoint = []leanruntime.Field{
	{Name: "VpcEndpointId", Flag: "vpc-endpoint-id", Type: "*string", Required: true},
	{Name: "VpcOptions", Flag: "vpc-options", Type: "*types.VPCOptions", Required: true},
}

var fields_upgrade_elasticsearch_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "PerformCheckOnly", Flag: "perform-check-only", Type: "*bool", Required: false},
	{Name: "TargetVersion", Flag: "target-version", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-inbound-cross-cluster-search-connection": {
			Name:   "accept-inbound-cross-cluster-search-connection",
			Fields: fields_accept_inbound_cross_cluster_search_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptInboundCrossClusterSearchConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_inbound_cross_cluster_search_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptInboundCrossClusterSearchConnection(ctx, input)
			},
		},
		"add-tags": {
			Name:   "add-tags",
			Fields: fields_add_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddTags(ctx, input)
			},
		},
		"associate-package": {
			Name:   "associate-package",
			Fields: fields_associate_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePackage(ctx, input)
			},
		},
		"authorize-vpc-endpoint-access": {
			Name:   "authorize-vpc-endpoint-access",
			Fields: fields_authorize_vpc_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AuthorizeVpcEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_authorize_vpc_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AuthorizeVpcEndpointAccess(ctx, input)
			},
		},
		"cancel-domain-config-change": {
			Name:   "cancel-domain-config-change",
			Fields: fields_cancel_domain_config_change,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDomainConfigChangeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_domain_config_change, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDomainConfigChange(ctx, input)
			},
		},
		"cancel-elasticsearch-service-software-update": {
			Name:   "cancel-elasticsearch-service-software-update",
			Fields: fields_cancel_elasticsearch_service_software_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelElasticsearchServiceSoftwareUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_elasticsearch_service_software_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelElasticsearchServiceSoftwareUpdate(ctx, input)
			},
		},
		"create-elasticsearch-domain": {
			Name:   "create-elasticsearch-domain",
			Fields: fields_create_elasticsearch_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateElasticsearchDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_elasticsearch_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateElasticsearchDomain(ctx, input)
			},
		},
		"create-outbound-cross-cluster-search-connection": {
			Name:   "create-outbound-cross-cluster-search-connection",
			Fields: fields_create_outbound_cross_cluster_search_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOutboundCrossClusterSearchConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_outbound_cross_cluster_search_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOutboundCrossClusterSearchConnection(ctx, input)
			},
		},
		"create-package": {
			Name:   "create-package",
			Fields: fields_create_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePackage(ctx, input)
			},
		},
		"create-vpc-endpoint": {
			Name:   "create-vpc-endpoint",
			Fields: fields_create_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVpcEndpoint(ctx, input)
			},
		},
		"delete-elasticsearch-domain": {
			Name:   "delete-elasticsearch-domain",
			Fields: fields_delete_elasticsearch_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteElasticsearchDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_elasticsearch_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteElasticsearchDomain(ctx, input)
			},
		},
		"delete-elasticsearch-service-role": {
			Name:   "delete-elasticsearch-service-role",
			Fields: fields_delete_elasticsearch_service_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteElasticsearchServiceRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_elasticsearch_service_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteElasticsearchServiceRole(ctx, input)
			},
		},
		"delete-inbound-cross-cluster-search-connection": {
			Name:   "delete-inbound-cross-cluster-search-connection",
			Fields: fields_delete_inbound_cross_cluster_search_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInboundCrossClusterSearchConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_inbound_cross_cluster_search_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInboundCrossClusterSearchConnection(ctx, input)
			},
		},
		"delete-outbound-cross-cluster-search-connection": {
			Name:   "delete-outbound-cross-cluster-search-connection",
			Fields: fields_delete_outbound_cross_cluster_search_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOutboundCrossClusterSearchConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_outbound_cross_cluster_search_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOutboundCrossClusterSearchConnection(ctx, input)
			},
		},
		"delete-package": {
			Name:   "delete-package",
			Fields: fields_delete_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePackage(ctx, input)
			},
		},
		"delete-vpc-endpoint": {
			Name:   "delete-vpc-endpoint",
			Fields: fields_delete_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVpcEndpoint(ctx, input)
			},
		},
		"describe-domain-auto-tunes": {
			Name:   "describe-domain-auto-tunes",
			Fields: fields_describe_domain_auto_tunes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainAutoTunesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_domain_auto_tunes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeDomainAutoTunes(ctx, input)
				}
				var results []*svc.DescribeDomainAutoTunesOutput
				p := svc.NewDescribeDomainAutoTunesPaginator(client, input)
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
		"describe-domain-change-progress": {
			Name:   "describe-domain-change-progress",
			Fields: fields_describe_domain_change_progress,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainChangeProgressInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain_change_progress, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomainChangeProgress(ctx, input)
			},
		},
		"describe-elasticsearch-domain": {
			Name:   "describe-elasticsearch-domain",
			Fields: fields_describe_elasticsearch_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeElasticsearchDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_elasticsearch_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeElasticsearchDomain(ctx, input)
			},
		},
		"describe-elasticsearch-domain-config": {
			Name:   "describe-elasticsearch-domain-config",
			Fields: fields_describe_elasticsearch_domain_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeElasticsearchDomainConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_elasticsearch_domain_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeElasticsearchDomainConfig(ctx, input)
			},
		},
		"describe-elasticsearch-domains": {
			Name:   "describe-elasticsearch-domains",
			Fields: fields_describe_elasticsearch_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeElasticsearchDomainsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_elasticsearch_domains, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeElasticsearchDomains(ctx, input)
			},
		},
		"describe-elasticsearch-instance-type-limits": {
			Name:   "describe-elasticsearch-instance-type-limits",
			Fields: fields_describe_elasticsearch_instance_type_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeElasticsearchInstanceTypeLimitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_elasticsearch_instance_type_limits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeElasticsearchInstanceTypeLimits(ctx, input)
			},
		},
		"describe-inbound-cross-cluster-search-connections": {
			Name:   "describe-inbound-cross-cluster-search-connections",
			Fields: fields_describe_inbound_cross_cluster_search_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeInboundCrossClusterSearchConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_inbound_cross_cluster_search_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeInboundCrossClusterSearchConnections(ctx, input)
				}
				var results []*svc.DescribeInboundCrossClusterSearchConnectionsOutput
				p := svc.NewDescribeInboundCrossClusterSearchConnectionsPaginator(client, input)
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
		"describe-outbound-cross-cluster-search-connections": {
			Name:   "describe-outbound-cross-cluster-search-connections",
			Fields: fields_describe_outbound_cross_cluster_search_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOutboundCrossClusterSearchConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_outbound_cross_cluster_search_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeOutboundCrossClusterSearchConnections(ctx, input)
				}
				var results []*svc.DescribeOutboundCrossClusterSearchConnectionsOutput
				p := svc.NewDescribeOutboundCrossClusterSearchConnectionsPaginator(client, input)
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
		"describe-packages": {
			Name:   "describe-packages",
			Fields: fields_describe_packages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePackagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_packages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePackages(ctx, input)
				}
				var results []*svc.DescribePackagesOutput
				p := svc.NewDescribePackagesPaginator(client, input)
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
		"describe-reserved-elasticsearch-instance-offerings": {
			Name:   "describe-reserved-elasticsearch-instance-offerings",
			Fields: fields_describe_reserved_elasticsearch_instance_offerings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedElasticsearchInstanceOfferingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_elasticsearch_instance_offerings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedElasticsearchInstanceOfferings(ctx, input)
				}
				var results []*svc.DescribeReservedElasticsearchInstanceOfferingsOutput
				p := svc.NewDescribeReservedElasticsearchInstanceOfferingsPaginator(client, input)
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
		"describe-reserved-elasticsearch-instances": {
			Name:   "describe-reserved-elasticsearch-instances",
			Fields: fields_describe_reserved_elasticsearch_instances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeReservedElasticsearchInstancesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_reserved_elasticsearch_instances, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeReservedElasticsearchInstances(ctx, input)
				}
				var results []*svc.DescribeReservedElasticsearchInstancesOutput
				p := svc.NewDescribeReservedElasticsearchInstancesPaginator(client, input)
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
		"describe-vpc-endpoints": {
			Name:   "describe-vpc-endpoints",
			Fields: fields_describe_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeVpcEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_vpc_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeVpcEndpoints(ctx, input)
			},
		},
		"dissociate-package": {
			Name:   "dissociate-package",
			Fields: fields_dissociate_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DissociatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_dissociate_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DissociatePackage(ctx, input)
			},
		},
		"get-compatible-elasticsearch-versions": {
			Name:   "get-compatible-elasticsearch-versions",
			Fields: fields_get_compatible_elasticsearch_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCompatibleElasticsearchVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_compatible_elasticsearch_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCompatibleElasticsearchVersions(ctx, input)
			},
		},
		"get-package-version-history": {
			Name:   "get-package-version-history",
			Fields: fields_get_package_version_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPackageVersionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_package_version_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetPackageVersionHistory(ctx, input)
				}
				var results []*svc.GetPackageVersionHistoryOutput
				p := svc.NewGetPackageVersionHistoryPaginator(client, input)
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
		"get-upgrade-history": {
			Name:   "get-upgrade-history",
			Fields: fields_get_upgrade_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUpgradeHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_upgrade_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetUpgradeHistory(ctx, input)
				}
				var results []*svc.GetUpgradeHistoryOutput
				p := svc.NewGetUpgradeHistoryPaginator(client, input)
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
		"get-upgrade-status": {
			Name:   "get-upgrade-status",
			Fields: fields_get_upgrade_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUpgradeStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_upgrade_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUpgradeStatus(ctx, input)
			},
		},
		"list-domain-names": {
			Name:   "list-domain-names",
			Fields: fields_list_domain_names,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainNamesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_domain_names, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDomainNames(ctx, input)
			},
		},
		"list-domains-for-package": {
			Name:   "list-domains-for-package",
			Fields: fields_list_domains_for_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsForPackageInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains_for_package, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainsForPackage(ctx, input)
				}
				var results []*svc.ListDomainsForPackageOutput
				p := svc.NewListDomainsForPackagePaginator(client, input)
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
		"list-elasticsearch-instance-types": {
			Name:   "list-elasticsearch-instance-types",
			Fields: fields_list_elasticsearch_instance_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListElasticsearchInstanceTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_elasticsearch_instance_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListElasticsearchInstanceTypes(ctx, input)
				}
				var results []*svc.ListElasticsearchInstanceTypesOutput
				p := svc.NewListElasticsearchInstanceTypesPaginator(client, input)
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
		"list-elasticsearch-versions": {
			Name:   "list-elasticsearch-versions",
			Fields: fields_list_elasticsearch_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListElasticsearchVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_elasticsearch_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListElasticsearchVersions(ctx, input)
				}
				var results []*svc.ListElasticsearchVersionsOutput
				p := svc.NewListElasticsearchVersionsPaginator(client, input)
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
		"list-packages-for-domain": {
			Name:   "list-packages-for-domain",
			Fields: fields_list_packages_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPackagesForDomainInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_packages_for_domain, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPackagesForDomain(ctx, input)
				}
				var results []*svc.ListPackagesForDomainOutput
				p := svc.NewListPackagesForDomainPaginator(client, input)
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
		"list-tags": {
			Name:   "list-tags",
			Fields: fields_list_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTags(ctx, input)
			},
		},
		"list-vpc-endpoint-access": {
			Name:   "list-vpc-endpoint-access",
			Fields: fields_list_vpc_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcEndpointAccess(ctx, input)
			},
		},
		"list-vpc-endpoints": {
			Name:   "list-vpc-endpoints",
			Fields: fields_list_vpc_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcEndpoints(ctx, input)
			},
		},
		"list-vpc-endpoints-for-domain": {
			Name:   "list-vpc-endpoints-for-domain",
			Fields: fields_list_vpc_endpoints_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVpcEndpointsForDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_vpc_endpoints_for_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListVpcEndpointsForDomain(ctx, input)
			},
		},
		"purchase-reserved-elasticsearch-instance-offering": {
			Name:   "purchase-reserved-elasticsearch-instance-offering",
			Fields: fields_purchase_reserved_elasticsearch_instance_offering,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PurchaseReservedElasticsearchInstanceOfferingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_purchase_reserved_elasticsearch_instance_offering, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PurchaseReservedElasticsearchInstanceOffering(ctx, input)
			},
		},
		"reject-inbound-cross-cluster-search-connection": {
			Name:   "reject-inbound-cross-cluster-search-connection",
			Fields: fields_reject_inbound_cross_cluster_search_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectInboundCrossClusterSearchConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_inbound_cross_cluster_search_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectInboundCrossClusterSearchConnection(ctx, input)
			},
		},
		"remove-tags": {
			Name:   "remove-tags",
			Fields: fields_remove_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveTags(ctx, input)
			},
		},
		"revoke-vpc-endpoint-access": {
			Name:   "revoke-vpc-endpoint-access",
			Fields: fields_revoke_vpc_endpoint_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeVpcEndpointAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_vpc_endpoint_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeVpcEndpointAccess(ctx, input)
			},
		},
		"start-elasticsearch-service-software-update": {
			Name:   "start-elasticsearch-service-software-update",
			Fields: fields_start_elasticsearch_service_software_update,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartElasticsearchServiceSoftwareUpdateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_elasticsearch_service_software_update, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartElasticsearchServiceSoftwareUpdate(ctx, input)
			},
		},
		"update-elasticsearch-domain-config": {
			Name:   "update-elasticsearch-domain-config",
			Fields: fields_update_elasticsearch_domain_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateElasticsearchDomainConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_elasticsearch_domain_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateElasticsearchDomainConfig(ctx, input)
			},
		},
		"update-package": {
			Name:   "update-package",
			Fields: fields_update_package,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePackageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_package, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePackage(ctx, input)
			},
		},
		"update-vpc-endpoint": {
			Name:   "update-vpc-endpoint",
			Fields: fields_update_vpc_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVpcEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vpc_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVpcEndpoint(ctx, input)
			},
		},
		"upgrade-elasticsearch-domain": {
			Name:   "upgrade-elasticsearch-domain",
			Fields: fields_upgrade_elasticsearch_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpgradeElasticsearchDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_upgrade_elasticsearch_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpgradeElasticsearchDomain(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("elasticsearchservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
