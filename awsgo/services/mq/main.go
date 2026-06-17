package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/mq/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-broker", "create-configuration", "create-tags", "create-user", "delete-broker", "delete-configuration", "delete-tags", "delete-user", "describe-broker", "describe-broker-engine-types", "describe-broker-instance-options", "describe-configuration", "describe-configuration-revision", "describe-user", "list-brokers", "list-configuration-revisions", "list-configurations", "list-tags", "list-users", "promote", "reboot-broker", "update-broker", "update-configuration", "update-user"},
		OperationSet: map[string]bool{"create-broker": true, "create-configuration": true, "create-tags": true, "create-user": true, "delete-broker": true, "delete-configuration": true, "delete-tags": true, "delete-user": true, "describe-broker": true, "describe-broker-engine-types": true, "describe-broker-instance-options": true, "describe-configuration": true, "describe-configuration-revision": true, "describe-user": true, "list-brokers": true, "list-configuration-revisions": true, "list-configurations": true, "list-tags": true, "list-users": true, "promote": true, "reboot-broker": true, "update-broker": true, "update-configuration": true, "update-user": true},
		OperationInputs: map[string][]string{
			"create-broker":                    {"AuthenticationStrategy", "AutoMinorVersionUpgrade", "BrokerName", "Configuration", "CreatorRequestId", "DataReplicationMode", "DataReplicationPrimaryBrokerArn", "DeploymentMode", "EncryptionOptions", "EngineType", "EngineVersion", "HostInstanceType", "LdapServerMetadata", "Logs", "MaintenanceWindowStartTime", "PubliclyAccessible", "SecurityGroups", "StorageType", "SubnetIds", "Tags", "Users"},
			"create-configuration":             {"AuthenticationStrategy", "EngineType", "EngineVersion", "Name", "Tags"},
			"create-tags":                      {"ResourceArn", "Tags"},
			"create-user":                      {"BrokerId", "ConsoleAccess", "Groups", "Password", "ReplicationUser", "Username"},
			"delete-broker":                    {"BrokerId"},
			"delete-configuration":             {"ConfigurationId"},
			"delete-tags":                      {"ResourceArn", "TagKeys"},
			"delete-user":                      {"BrokerId", "Username"},
			"describe-broker":                  {"BrokerId"},
			"describe-broker-engine-types":     {"EngineType", "MaxResults", "NextToken"},
			"describe-broker-instance-options": {"EngineType", "HostInstanceType", "MaxResults", "NextToken", "StorageType"},
			"describe-configuration":           {"ConfigurationId"},
			"describe-configuration-revision":  {"ConfigurationId", "ConfigurationRevision"},
			"describe-user":                    {"BrokerId", "Username"},
			"list-brokers":                     {"MaxResults", "NextToken"},
			"list-configuration-revisions":     {"ConfigurationId", "MaxResults", "NextToken"},
			"list-configurations":              {"MaxResults", "NextToken"},
			"list-tags":                        {"ResourceArn"},
			"list-users":                       {"BrokerId", "MaxResults", "NextToken"},
			"promote":                          {"BrokerId", "Mode"},
			"reboot-broker":                    {"BrokerId"},
			"update-broker":                    {"AuthenticationStrategy", "AutoMinorVersionUpgrade", "BrokerId", "Configuration", "DataReplicationMode", "EngineVersion", "HostInstanceType", "LdapServerMetadata", "Logs", "MaintenanceWindowStartTime", "SecurityGroups"},
			"update-configuration":             {"ConfigurationId", "Data", "Description"},
			"update-user":                      {"BrokerId", "ConsoleAccess", "Groups", "Password", "ReplicationUser", "Username"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-broker":                    {"AuthenticationStrategy": "types.AuthenticationStrategy", "AutoMinorVersionUpgrade": "*bool", "BrokerName": "*string", "Configuration": "*types.ConfigurationId", "CreatorRequestId": "*string", "DataReplicationMode": "types.DataReplicationMode", "DataReplicationPrimaryBrokerArn": "*string", "DeploymentMode": "types.DeploymentMode", "EncryptionOptions": "*types.EncryptionOptions", "EngineType": "types.EngineType", "EngineVersion": "*string", "HostInstanceType": "*string", "LdapServerMetadata": "*types.LdapServerMetadataInput", "Logs": "*types.Logs", "MaintenanceWindowStartTime": "*types.WeeklyStartTime", "PubliclyAccessible": "*bool", "SecurityGroups": "[]string", "StorageType": "types.BrokerStorageType", "SubnetIds": "[]string", "Tags": "map[string]string", "Users": "[]types.User"},
			"create-configuration":             {"AuthenticationStrategy": "types.AuthenticationStrategy", "EngineType": "types.EngineType", "EngineVersion": "*string", "Name": "*string", "Tags": "map[string]string"},
			"create-tags":                      {"ResourceArn": "*string", "Tags": "map[string]string"},
			"create-user":                      {"BrokerId": "*string", "ConsoleAccess": "*bool", "Groups": "[]string", "Password": "*string", "ReplicationUser": "*bool", "Username": "*string"},
			"delete-broker":                    {"BrokerId": "*string"},
			"delete-configuration":             {"ConfigurationId": "*string"},
			"delete-tags":                      {"ResourceArn": "*string", "TagKeys": "[]string"},
			"delete-user":                      {"BrokerId": "*string", "Username": "*string"},
			"describe-broker":                  {"BrokerId": "*string"},
			"describe-broker-engine-types":     {"EngineType": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"describe-broker-instance-options": {"EngineType": "*string", "HostInstanceType": "*string", "MaxResults": "*int32", "NextToken": "*string", "StorageType": "*string"},
			"describe-configuration":           {"ConfigurationId": "*string"},
			"describe-configuration-revision":  {"ConfigurationId": "*string", "ConfigurationRevision": "*string"},
			"describe-user":                    {"BrokerId": "*string", "Username": "*string"},
			"list-brokers":                     {"MaxResults": "*int32", "NextToken": "*string"},
			"list-configuration-revisions":     {"ConfigurationId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-configurations":              {"MaxResults": "*int32", "NextToken": "*string"},
			"list-tags":                        {"ResourceArn": "*string"},
			"list-users":                       {"BrokerId": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"promote":                          {"BrokerId": "*string", "Mode": "types.PromoteMode"},
			"reboot-broker":                    {"BrokerId": "*string"},
			"update-broker":                    {"AuthenticationStrategy": "types.AuthenticationStrategy", "AutoMinorVersionUpgrade": "*bool", "BrokerId": "*string", "Configuration": "*types.ConfigurationId", "DataReplicationMode": "types.DataReplicationMode", "EngineVersion": "*string", "HostInstanceType": "*string", "LdapServerMetadata": "*types.LdapServerMetadataInput", "Logs": "*types.Logs", "MaintenanceWindowStartTime": "*types.WeeklyStartTime", "SecurityGroups": "[]string"},
			"update-configuration":             {"ConfigurationId": "*string", "Data": "*string", "Description": "*string"},
			"update-user":                      {"BrokerId": "*string", "ConsoleAccess": "*bool", "Groups": "[]string", "Password": "*string", "ReplicationUser": "*bool", "Username": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-broker":                    {"BrokerName", "DeploymentMode", "EngineType", "HostInstanceType", "PubliclyAccessible"},
			"create-configuration":             {"EngineType", "Name"},
			"create-tags":                      {"ResourceArn"},
			"create-user":                      {"BrokerId", "Password", "Username"},
			"delete-broker":                    {"BrokerId"},
			"delete-configuration":             {"ConfigurationId"},
			"delete-tags":                      {"ResourceArn", "TagKeys"},
			"delete-user":                      {"BrokerId", "Username"},
			"describe-broker":                  {"BrokerId"},
			"describe-broker-engine-types":     {},
			"describe-broker-instance-options": {},
			"describe-configuration":           {"ConfigurationId"},
			"describe-configuration-revision":  {"ConfigurationId", "ConfigurationRevision"},
			"describe-user":                    {"BrokerId", "Username"},
			"list-brokers":                     {},
			"list-configuration-revisions":     {"ConfigurationId"},
			"list-configurations":              {},
			"list-tags":                        {"ResourceArn"},
			"list-users":                       {"BrokerId"},
			"promote":                          {"BrokerId", "Mode"},
			"reboot-broker":                    {"BrokerId"},
			"update-broker":                    {"BrokerId"},
			"update-configuration":             {"ConfigurationId", "Data"},
			"update-user":                      {"BrokerId", "Username"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("mq", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
