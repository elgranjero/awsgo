package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/kafkaconnect/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-connector", "create-custom-plugin", "create-worker-configuration", "delete-connector", "delete-custom-plugin", "delete-worker-configuration", "describe-connector", "describe-connector-operation", "describe-custom-plugin", "describe-worker-configuration", "list-connector-operations", "list-connectors", "list-custom-plugins", "list-tags-for-resource", "list-worker-configurations", "tag-resource", "untag-resource", "update-connector"},
		OperationSet: map[string]bool{"create-connector": true, "create-custom-plugin": true, "create-worker-configuration": true, "delete-connector": true, "delete-custom-plugin": true, "delete-worker-configuration": true, "describe-connector": true, "describe-connector-operation": true, "describe-custom-plugin": true, "describe-worker-configuration": true, "list-connector-operations": true, "list-connectors": true, "list-custom-plugins": true, "list-tags-for-resource": true, "list-worker-configurations": true, "tag-resource": true, "untag-resource": true, "update-connector": true},
		OperationInputs: map[string][]string{
			"create-connector":              {"Capacity", "ConnectorConfiguration", "ConnectorDescription", "ConnectorName", "KafkaCluster", "KafkaClusterClientAuthentication", "KafkaClusterEncryptionInTransit", "KafkaConnectVersion", "LogDelivery", "NetworkType", "Plugins", "ServiceExecutionRoleArn", "Tags", "WorkerConfiguration"},
			"create-custom-plugin":          {"ContentType", "Description", "Location", "Name", "Tags"},
			"create-worker-configuration":   {"Description", "Name", "PropertiesFileContent", "Tags"},
			"delete-connector":              {"ConnectorArn", "CurrentVersion"},
			"delete-custom-plugin":          {"CustomPluginArn"},
			"delete-worker-configuration":   {"WorkerConfigurationArn"},
			"describe-connector":            {"ConnectorArn"},
			"describe-connector-operation":  {"ConnectorOperationArn"},
			"describe-custom-plugin":        {"CustomPluginArn"},
			"describe-worker-configuration": {"WorkerConfigurationArn"},
			"list-connector-operations":     {"ConnectorArn", "MaxResults", "NextToken"},
			"list-connectors":               {"ConnectorNamePrefix", "MaxResults", "NextToken"},
			"list-custom-plugins":           {"MaxResults", "NamePrefix", "NextToken"},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-worker-configurations":    {"MaxResults", "NamePrefix", "NextToken"},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-connector":              {"Capacity", "ConnectorArn", "ConnectorConfiguration", "CurrentVersion"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-connector":              {"Capacity": "*types.Capacity", "ConnectorConfiguration": "map[string]string", "ConnectorDescription": "*string", "ConnectorName": "*string", "KafkaCluster": "*types.KafkaCluster", "KafkaClusterClientAuthentication": "*types.KafkaClusterClientAuthentication", "KafkaClusterEncryptionInTransit": "*types.KafkaClusterEncryptionInTransit", "KafkaConnectVersion": "*string", "LogDelivery": "*types.LogDelivery", "NetworkType": "types.NetworkType", "Plugins": "[]types.Plugin", "ServiceExecutionRoleArn": "*string", "Tags": "map[string]string", "WorkerConfiguration": "*types.WorkerConfiguration"},
			"create-custom-plugin":          {"ContentType": "types.CustomPluginContentType", "Description": "*string", "Location": "*types.CustomPluginLocation", "Name": "*string", "Tags": "map[string]string"},
			"create-worker-configuration":   {"Description": "*string", "Name": "*string", "PropertiesFileContent": "*string", "Tags": "map[string]string"},
			"delete-connector":              {"ConnectorArn": "*string", "CurrentVersion": "*string"},
			"delete-custom-plugin":          {"CustomPluginArn": "*string"},
			"delete-worker-configuration":   {"WorkerConfigurationArn": "*string"},
			"describe-connector":            {"ConnectorArn": "*string"},
			"describe-connector-operation":  {"ConnectorOperationArn": "*string"},
			"describe-custom-plugin":        {"CustomPluginArn": "*string"},
			"describe-worker-configuration": {"WorkerConfigurationArn": "*string"},
			"list-connector-operations":     {"ConnectorArn": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-connectors":               {"ConnectorNamePrefix": "*string", "MaxResults": "*int32", "NextToken": "*string"},
			"list-custom-plugins":           {"MaxResults": "*int32", "NamePrefix": "*string", "NextToken": "*string"},
			"list-tags-for-resource":        {"ResourceArn": "*string"},
			"list-worker-configurations":    {"MaxResults": "*int32", "NamePrefix": "*string", "NextToken": "*string"},
			"tag-resource":                  {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-connector":              {"Capacity": "*types.CapacityUpdate", "ConnectorArn": "*string", "ConnectorConfiguration": "map[string]string", "CurrentVersion": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-connector":              {"Capacity", "ConnectorConfiguration", "ConnectorName", "KafkaCluster", "KafkaClusterClientAuthentication", "KafkaClusterEncryptionInTransit", "KafkaConnectVersion", "Plugins", "ServiceExecutionRoleArn"},
			"create-custom-plugin":          {"ContentType", "Location", "Name"},
			"create-worker-configuration":   {"Name", "PropertiesFileContent"},
			"delete-connector":              {"ConnectorArn"},
			"delete-custom-plugin":          {"CustomPluginArn"},
			"delete-worker-configuration":   {"WorkerConfigurationArn"},
			"describe-connector":            {"ConnectorArn"},
			"describe-connector-operation":  {"ConnectorOperationArn"},
			"describe-custom-plugin":        {"CustomPluginArn"},
			"describe-worker-configuration": {"WorkerConfigurationArn"},
			"list-connector-operations":     {"ConnectorArn"},
			"list-connectors":               {},
			"list-custom-plugins":           {},
			"list-tags-for-resource":        {"ResourceArn"},
			"list-worker-configurations":    {},
			"tag-resource":                  {"ResourceArn", "Tags"},
			"untag-resource":                {"ResourceArn", "TagKeys"},
			"update-connector":              {"ConnectorArn", "CurrentVersion"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("kafkaconnect", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
