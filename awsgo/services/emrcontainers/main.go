package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/emrcontainers/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"cancel-job-run", "create-job-template", "create-managed-endpoint", "create-security-configuration", "create-virtual-cluster", "delete-job-template", "delete-managed-endpoint", "delete-virtual-cluster", "describe-job-run", "describe-job-template", "describe-managed-endpoint", "describe-security-configuration", "describe-virtual-cluster", "get-managed-endpoint-session-credentials", "list-job-runs", "list-job-templates", "list-managed-endpoints", "list-security-configurations", "list-tags-for-resource", "list-virtual-clusters", "start-job-run", "tag-resource", "untag-resource"},
		OperationSet: map[string]bool{"cancel-job-run": true, "create-job-template": true, "create-managed-endpoint": true, "create-security-configuration": true, "create-virtual-cluster": true, "delete-job-template": true, "delete-managed-endpoint": true, "delete-virtual-cluster": true, "describe-job-run": true, "describe-job-template": true, "describe-managed-endpoint": true, "describe-security-configuration": true, "describe-virtual-cluster": true, "get-managed-endpoint-session-credentials": true, "list-job-runs": true, "list-job-templates": true, "list-managed-endpoints": true, "list-security-configurations": true, "list-tags-for-resource": true, "list-virtual-clusters": true, "start-job-run": true, "tag-resource": true, "untag-resource": true},
		OperationInputs: map[string][]string{
			"cancel-job-run":                           {"Id", "VirtualClusterId"},
			"create-job-template":                      {"ClientToken", "JobTemplateData", "KmsKeyArn", "Name", "Tags"},
			"create-managed-endpoint":                  {"CertificateArn", "ClientToken", "ConfigurationOverrides", "ExecutionRoleArn", "Name", "ReleaseLabel", "Tags", "Type", "VirtualClusterId"},
			"create-security-configuration":            {"ClientToken", "ContainerProvider", "Name", "SecurityConfigurationData", "Tags"},
			"create-virtual-cluster":                   {"ClientToken", "ContainerProvider", "Name", "SecurityConfigurationId", "Tags"},
			"delete-job-template":                      {"Id"},
			"delete-managed-endpoint":                  {"Id", "VirtualClusterId"},
			"delete-virtual-cluster":                   {"Id"},
			"describe-job-run":                         {"Id", "VirtualClusterId"},
			"describe-job-template":                    {"Id"},
			"describe-managed-endpoint":                {"Id", "VirtualClusterId"},
			"describe-security-configuration":          {"Id"},
			"describe-virtual-cluster":                 {"Id"},
			"get-managed-endpoint-session-credentials": {"ClientToken", "CredentialType", "DurationInSeconds", "EndpointIdentifier", "ExecutionRoleArn", "LogContext", "VirtualClusterIdentifier"},
			"list-job-runs":                            {"CreatedAfter", "CreatedBefore", "MaxResults", "Name", "NextToken", "States", "VirtualClusterId"},
			"list-job-templates":                       {"CreatedAfter", "CreatedBefore", "MaxResults", "NextToken"},
			"list-managed-endpoints":                   {"CreatedAfter", "CreatedBefore", "MaxResults", "NextToken", "States", "Types", "VirtualClusterId"},
			"list-security-configurations":             {"CreatedAfter", "CreatedBefore", "MaxResults", "NextToken"},
			"list-tags-for-resource":                   {"ResourceArn"},
			"list-virtual-clusters":                    {"ContainerProviderId", "ContainerProviderType", "CreatedAfter", "CreatedBefore", "EksAccessEntryIntegrated", "MaxResults", "NextToken", "States"},
			"start-job-run":                            {"ClientToken", "ConfigurationOverrides", "ExecutionRoleArn", "JobDriver", "JobTemplateId", "JobTemplateParameters", "Name", "ReleaseLabel", "RetryPolicyConfiguration", "Tags", "VirtualClusterId"},
			"tag-resource":                             {"ResourceArn", "Tags"},
			"untag-resource":                           {"ResourceArn", "TagKeys"},
		},
		OperationInputTypes: map[string]map[string]string{
			"cancel-job-run":                           {"Id": "*string", "VirtualClusterId": "*string"},
			"create-job-template":                      {"ClientToken": "*string", "JobTemplateData": "*types.JobTemplateData", "KmsKeyArn": "*string", "Name": "*string", "Tags": "map[string]string"},
			"create-managed-endpoint":                  {"CertificateArn": "*string", "ClientToken": "*string", "ConfigurationOverrides": "*types.ConfigurationOverrides", "ExecutionRoleArn": "*string", "Name": "*string", "ReleaseLabel": "*string", "Tags": "map[string]string", "Type": "*string", "VirtualClusterId": "*string"},
			"create-security-configuration":            {"ClientToken": "*string", "ContainerProvider": "*types.ContainerProvider", "Name": "*string", "SecurityConfigurationData": "*types.SecurityConfigurationData", "Tags": "map[string]string"},
			"create-virtual-cluster":                   {"ClientToken": "*string", "ContainerProvider": "*types.ContainerProvider", "Name": "*string", "SecurityConfigurationId": "*string", "Tags": "map[string]string"},
			"delete-job-template":                      {"Id": "*string"},
			"delete-managed-endpoint":                  {"Id": "*string", "VirtualClusterId": "*string"},
			"delete-virtual-cluster":                   {"Id": "*string"},
			"describe-job-run":                         {"Id": "*string", "VirtualClusterId": "*string"},
			"describe-job-template":                    {"Id": "*string"},
			"describe-managed-endpoint":                {"Id": "*string", "VirtualClusterId": "*string"},
			"describe-security-configuration":          {"Id": "*string"},
			"describe-virtual-cluster":                 {"Id": "*string"},
			"get-managed-endpoint-session-credentials": {"ClientToken": "*string", "CredentialType": "*string", "DurationInSeconds": "*int32", "EndpointIdentifier": "*string", "ExecutionRoleArn": "*string", "LogContext": "*string", "VirtualClusterIdentifier": "*string"},
			"list-job-runs":                            {"CreatedAfter": "*time.Time", "CreatedBefore": "*time.Time", "MaxResults": "*int32", "Name": "*string", "NextToken": "*string", "States": "[]types.JobRunState", "VirtualClusterId": "*string"},
			"list-job-templates":                       {"CreatedAfter": "*time.Time", "CreatedBefore": "*time.Time", "MaxResults": "*int32", "NextToken": "*string"},
			"list-managed-endpoints":                   {"CreatedAfter": "*time.Time", "CreatedBefore": "*time.Time", "MaxResults": "*int32", "NextToken": "*string", "States": "[]types.EndpointState", "Types": "[]string", "VirtualClusterId": "*string"},
			"list-security-configurations":             {"CreatedAfter": "*time.Time", "CreatedBefore": "*time.Time", "MaxResults": "*int32", "NextToken": "*string"},
			"list-tags-for-resource":                   {"ResourceArn": "*string"},
			"list-virtual-clusters":                    {"ContainerProviderId": "*string", "ContainerProviderType": "types.ContainerProviderType", "CreatedAfter": "*time.Time", "CreatedBefore": "*time.Time", "EksAccessEntryIntegrated": "*bool", "MaxResults": "*int32", "NextToken": "*string", "States": "[]types.VirtualClusterState"},
			"start-job-run":                            {"ClientToken": "*string", "ConfigurationOverrides": "*types.ConfigurationOverrides", "ExecutionRoleArn": "*string", "JobDriver": "*types.JobDriver", "JobTemplateId": "*string", "JobTemplateParameters": "map[string]string", "Name": "*string", "ReleaseLabel": "*string", "RetryPolicyConfiguration": "*types.RetryPolicyConfiguration", "Tags": "map[string]string", "VirtualClusterId": "*string"},
			"tag-resource":                             {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":                           {"ResourceArn": "*string", "TagKeys": "[]string"},
		},
		OperationInputRequired: map[string][]string{
			"cancel-job-run":                           {"Id", "VirtualClusterId"},
			"create-job-template":                      {"ClientToken", "JobTemplateData", "Name"},
			"create-managed-endpoint":                  {"ClientToken", "ExecutionRoleArn", "Name", "ReleaseLabel", "Type", "VirtualClusterId"},
			"create-security-configuration":            {"ClientToken", "Name", "SecurityConfigurationData"},
			"create-virtual-cluster":                   {"ClientToken", "ContainerProvider", "Name"},
			"delete-job-template":                      {"Id"},
			"delete-managed-endpoint":                  {"Id", "VirtualClusterId"},
			"delete-virtual-cluster":                   {"Id"},
			"describe-job-run":                         {"Id", "VirtualClusterId"},
			"describe-job-template":                    {"Id"},
			"describe-managed-endpoint":                {"Id", "VirtualClusterId"},
			"describe-security-configuration":          {"Id"},
			"describe-virtual-cluster":                 {"Id"},
			"get-managed-endpoint-session-credentials": {"CredentialType", "EndpointIdentifier", "ExecutionRoleArn", "VirtualClusterIdentifier"},
			"list-job-runs":                            {"VirtualClusterId"},
			"list-job-templates":                       {},
			"list-managed-endpoints":                   {"VirtualClusterId"},
			"list-security-configurations":             {},
			"list-tags-for-resource":                   {"ResourceArn"},
			"list-virtual-clusters":                    {},
			"start-job-run":                            {"ClientToken", "VirtualClusterId"},
			"tag-resource":                             {"ResourceArn", "Tags"},
			"untag-resource":                           {"ResourceArn", "TagKeys"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("emrcontainers", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
