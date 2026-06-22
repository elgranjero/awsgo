package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mq"
)

var fields_create_broker = []leanruntime.Field{
	{Name: "AuthenticationStrategy", Flag: "authentication-strategy", Type: "types.AuthenticationStrategy", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "BrokerName", Flag: "broker-name", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ConfigurationId", Required: false},
	{Name: "CreatorRequestId", Flag: "creator-request-id", Type: "*string", Required: false},
	{Name: "DataReplicationMode", Flag: "data-replication-mode", Type: "types.DataReplicationMode", Required: false},
	{Name: "DataReplicationPrimaryBrokerArn", Flag: "data-replication-primary-broker-arn", Type: "*string", Required: false},
	{Name: "DeploymentMode", Flag: "deployment-mode", Type: "types.DeploymentMode", Required: true},
	{Name: "EncryptionOptions", Flag: "encryption-options", Type: "*types.EncryptionOptions", Required: false},
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "HostInstanceType", Flag: "host-instance-type", Type: "*string", Required: true},
	{Name: "LdapServerMetadata", Flag: "ldap-server-metadata", Type: "*types.LdapServerMetadataInput", Required: false},
	{Name: "Logs", Flag: "logs", Type: "*types.Logs", Required: false},
	{Name: "MaintenanceWindowStartTime", Flag: "maintenance-window-start-time", Type: "*types.WeeklyStartTime", Required: false},
	{Name: "PubliclyAccessible", Flag: "publicly-accessible", Type: "*bool", Required: true},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "types.BrokerStorageType", Required: false},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Users", Flag: "users", Type: "[]types.User", Required: false},
}

var fields_create_configuration = []leanruntime.Field{
	{Name: "AuthenticationStrategy", Flag: "authentication-strategy", Type: "types.AuthenticationStrategy", Required: false},
	{Name: "EngineType", Flag: "engine-type", Type: "types.EngineType", Required: true},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_user = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "ConsoleAccess", Flag: "console-access", Type: "*bool", Required: false},
	{Name: "Groups", Flag: "groups", Type: "[]string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
	{Name: "ReplicationUser", Flag: "replication-user", Type: "*bool", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_delete_broker = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
}

var fields_delete_configuration = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
}

var fields_delete_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_delete_user = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_describe_broker = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
}

var fields_describe_broker_engine_types = []leanruntime.Field{
	{Name: "EngineType", Flag: "engine-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_describe_broker_instance_options = []leanruntime.Field{
	{Name: "EngineType", Flag: "engine-type", Type: "*string", Required: false},
	{Name: "HostInstanceType", Flag: "host-instance-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StorageType", Flag: "storage-type", Type: "*string", Required: false},
}

var fields_describe_configuration = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
}

var fields_describe_configuration_revision = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
	{Name: "ConfigurationRevision", Flag: "configuration-revision", Type: "*string", Required: true},
}

var fields_describe_user = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

var fields_list_brokers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configuration_revisions = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_users = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_promote = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "Mode", Flag: "mode", Type: "types.PromoteMode", Required: true},
}

var fields_reboot_broker = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
}

var fields_update_broker = []leanruntime.Field{
	{Name: "AuthenticationStrategy", Flag: "authentication-strategy", Type: "types.AuthenticationStrategy", Required: false},
	{Name: "AutoMinorVersionUpgrade", Flag: "auto-minor-version-upgrade", Type: "*bool", Required: false},
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "*types.ConfigurationId", Required: false},
	{Name: "DataReplicationMode", Flag: "data-replication-mode", Type: "types.DataReplicationMode", Required: false},
	{Name: "EngineVersion", Flag: "engine-version", Type: "*string", Required: false},
	{Name: "HostInstanceType", Flag: "host-instance-type", Type: "*string", Required: false},
	{Name: "LdapServerMetadata", Flag: "ldap-server-metadata", Type: "*types.LdapServerMetadataInput", Required: false},
	{Name: "Logs", Flag: "logs", Type: "*types.Logs", Required: false},
	{Name: "MaintenanceWindowStartTime", Flag: "maintenance-window-start-time", Type: "*types.WeeklyStartTime", Required: false},
	{Name: "SecurityGroups", Flag: "security-groups", Type: "[]string", Required: false},
}

var fields_update_configuration = []leanruntime.Field{
	{Name: "ConfigurationId", Flag: "configuration-id", Type: "*string", Required: true},
	{Name: "Data", Flag: "data", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
}

var fields_update_user = []leanruntime.Field{
	{Name: "BrokerId", Flag: "broker-id", Type: "*string", Required: true},
	{Name: "ConsoleAccess", Flag: "console-access", Type: "*bool", Required: false},
	{Name: "Groups", Flag: "groups", Type: "[]string", Required: false},
	{Name: "Password", Flag: "password", Type: "*string", Required: false},
	{Name: "ReplicationUser", Flag: "replication-user", Type: "*bool", Required: false},
	{Name: "Username", Flag: "username", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-broker": {
			Name:   "create-broker",
			Fields: fields_create_broker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBrokerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_broker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBroker(ctx, input)
			},
		},
		"create-configuration": {
			Name:   "create-configuration",
			Fields: fields_create_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConfiguration(ctx, input)
			},
		},
		"create-tags": {
			Name:   "create-tags",
			Fields: fields_create_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTags(ctx, input)
			},
		},
		"create-user": {
			Name:   "create-user",
			Fields: fields_create_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUser(ctx, input)
			},
		},
		"delete-broker": {
			Name:   "delete-broker",
			Fields: fields_delete_broker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBrokerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_broker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBroker(ctx, input)
			},
		},
		"delete-configuration": {
			Name:   "delete-configuration",
			Fields: fields_delete_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConfiguration(ctx, input)
			},
		},
		"delete-tags": {
			Name:   "delete-tags",
			Fields: fields_delete_tags,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTags(ctx, input)
			},
		},
		"delete-user": {
			Name:   "delete-user",
			Fields: fields_delete_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUser(ctx, input)
			},
		},
		"describe-broker": {
			Name:   "describe-broker",
			Fields: fields_describe_broker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBrokerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_broker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBroker(ctx, input)
			},
		},
		"describe-broker-engine-types": {
			Name:   "describe-broker-engine-types",
			Fields: fields_describe_broker_engine_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBrokerEngineTypesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_broker_engine_types, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBrokerEngineTypes(ctx, input)
			},
		},
		"describe-broker-instance-options": {
			Name:   "describe-broker-instance-options",
			Fields: fields_describe_broker_instance_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeBrokerInstanceOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_broker_instance_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeBrokerInstanceOptions(ctx, input)
			},
		},
		"describe-configuration": {
			Name:   "describe-configuration",
			Fields: fields_describe_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfiguration(ctx, input)
			},
		},
		"describe-configuration-revision": {
			Name:   "describe-configuration-revision",
			Fields: fields_describe_configuration_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConfigurationRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_configuration_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConfigurationRevision(ctx, input)
			},
		},
		"describe-user": {
			Name:   "describe-user",
			Fields: fields_describe_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeUser(ctx, input)
			},
		},
		"list-brokers": {
			Name:   "list-brokers",
			Fields: fields_list_brokers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBrokersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_brokers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBrokers(ctx, input)
				}
				var results []*svc.ListBrokersOutput
				p := svc.NewListBrokersPaginator(client, input)
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
		"list-configuration-revisions": {
			Name:   "list-configuration-revisions",
			Fields: fields_list_configuration_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationRevisionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_configuration_revisions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConfigurationRevisions(ctx, input)
			},
		},
		"list-configurations": {
			Name:   "list-configurations",
			Fields: fields_list_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListConfigurations(ctx, input)
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
		"list-users": {
			Name:   "list-users",
			Fields: fields_list_users,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_users, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListUsers(ctx, input)
			},
		},
		"promote": {
			Name:   "promote",
			Fields: fields_promote,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PromoteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_promote, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.Promote(ctx, input)
			},
		},
		"reboot-broker": {
			Name:   "reboot-broker",
			Fields: fields_reboot_broker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RebootBrokerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reboot_broker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RebootBroker(ctx, input)
			},
		},
		"update-broker": {
			Name:   "update-broker",
			Fields: fields_update_broker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBrokerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_broker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBroker(ctx, input)
			},
		},
		"update-configuration": {
			Name:   "update-configuration",
			Fields: fields_update_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConfiguration(ctx, input)
			},
		},
		"update-user": {
			Name:   "update-user",
			Fields: fields_update_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUser(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mq", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
