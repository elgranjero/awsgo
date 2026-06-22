package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/deadline"
)

var fields_associate_member_to_farm = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MembershipLevel", Flag: "membership-level", Type: "types.MembershipLevel", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.DeadlinePrincipalType", Required: true},
}

var fields_associate_member_to_fleet = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MembershipLevel", Flag: "membership-level", Type: "types.MembershipLevel", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.DeadlinePrincipalType", Required: true},
}

var fields_associate_member_to_job = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MembershipLevel", Flag: "membership-level", Type: "types.MembershipLevel", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.DeadlinePrincipalType", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_associate_member_to_queue = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "IdentityStoreId", Flag: "identity-store-id", Type: "*string", Required: true},
	{Name: "MembershipLevel", Flag: "membership-level", Type: "types.MembershipLevel", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.DeadlinePrincipalType", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_assume_fleet_role_for_read = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_assume_fleet_role_for_worker = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_assume_queue_role_for_read = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_assume_queue_role_for_user = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_assume_queue_role_for_worker = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_batch_get_job_entity = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "Identifiers", Flag: "identifiers", Type: "[]types.JobEntityIdentifiersUnion", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_copy_job_template = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "TargetS3Location", Flag: "target-s3-location", Type: "*types.S3Location", Required: true},
}

var fields_create_budget = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "[]types.BudgetActionToAdd", Required: true},
	{Name: "ApproximateDollarLimit", Flag: "approximate-dollar-limit", Type: "*float32", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "types.BudgetSchedule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UsageTrackingResource", Flag: "usage-tracking-resource", Type: "types.UsageTrackingResource", Required: true},
}

var fields_create_farm = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "KmsKeyArn", Flag: "kms-key-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_fleet = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.FleetConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "HostConfiguration", Flag: "host-configuration", Type: "*types.HostConfiguration", Required: false},
	{Name: "MaxWorkerCount", Flag: "max-worker-count", Type: "*int32", Required: true},
	{Name: "MinWorkerCount", Flag: "min-worker-count", Type: "int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_job = []leanruntime.Field{
	{Name: "Attachments", Flag: "attachments", Type: "*types.Attachments", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DescriptionOverride", Flag: "description-override", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxFailedTasksCount", Flag: "max-failed-tasks-count", Type: "*int32", Required: false},
	{Name: "MaxRetriesPerTask", Flag: "max-retries-per-task", Type: "*int32", Required: false},
	{Name: "MaxWorkerCount", Flag: "max-worker-count", Type: "*int32", Required: false},
	{Name: "NameOverride", Flag: "name-override", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "map[string]types.JobParameter", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "SourceJobId", Flag: "source-job-id", Type: "*string", Required: false},
	{Name: "StorageProfileId", Flag: "storage-profile-id", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TargetTaskRunStatus", Flag: "target-task-run-status", Type: "types.CreateJobTargetTaskRunStatus", Required: false},
	{Name: "Template", Flag: "template", Type: "*string", Required: false},
	{Name: "TemplateType", Flag: "template-type", Type: "types.JobTemplateType", Required: false},
}

var fields_create_license_endpoint = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "SecurityGroupIds", Flag: "security-group-ids", Type: "[]string", Required: true},
	{Name: "SubnetIds", Flag: "subnet-ids", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VpcId", Flag: "vpc-id", Type: "*string", Required: true},
}

var fields_create_limit = []leanruntime.Field{
	{Name: "AmountRequirementName", Flag: "amount-requirement-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxCount", Flag: "max-count", Type: "*int32", Required: true},
}

var fields_create_monitor = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "IdentityCenterInstanceArn", Flag: "identity-center-instance-arn", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: true},
	{Name: "Subdomain", Flag: "subdomain", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_queue = []leanruntime.Field{
	{Name: "AllowedStorageProfileIds", Flag: "allowed-storage-profile-ids", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultBudgetAction", Flag: "default-budget-action", Type: "types.DefaultQueueBudgetAction", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobAttachmentSettings", Flag: "job-attachment-settings", Type: "*types.JobAttachmentSettings", Required: false},
	{Name: "JobRunAsUser", Flag: "job-run-as-user", Type: "*types.JobRunAsUser", Required: false},
	{Name: "RequiredFileSystemLocationNames", Flag: "required-file-system-location-names", Type: "[]string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_queue_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "Template", Flag: "template", Type: "*string", Required: true},
	{Name: "TemplateType", Flag: "template-type", Type: "types.EnvironmentTemplateType", Required: true},
}

var fields_create_queue_fleet_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_create_queue_limit_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_create_storage_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FileSystemLocations", Flag: "file-system-locations", Type: "[]types.FileSystemLocation", Required: false},
	{Name: "OsFamily", Flag: "os-family", Type: "types.StorageProfileOperatingSystemFamily", Required: true},
}

var fields_create_worker = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "HostProperties", Flag: "host-properties", Type: "*types.HostPropertiesRequest", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_budget = []leanruntime.Field{
	{Name: "BudgetId", Flag: "budget-id", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
}

var fields_delete_farm = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
}

var fields_delete_fleet = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_delete_license_endpoint = []leanruntime.Field{
	{Name: "LicenseEndpointId", Flag: "license-endpoint-id", Type: "*string", Required: true},
}

var fields_delete_limit = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
}

var fields_delete_metered_product = []leanruntime.Field{
	{Name: "LicenseEndpointId", Flag: "license-endpoint-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_delete_monitor = []leanruntime.Field{
	{Name: "MonitorId", Flag: "monitor-id", Type: "*string", Required: true},
}

var fields_delete_queue = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_delete_queue_environment = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueEnvironmentId", Flag: "queue-environment-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_delete_queue_fleet_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_delete_queue_limit_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_delete_storage_profile = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "StorageProfileId", Flag: "storage-profile-id", Type: "*string", Required: true},
}

var fields_delete_worker = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_disassociate_member_from_farm = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
}

var fields_disassociate_member_from_fleet = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
}

var fields_disassociate_member_from_job = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_disassociate_member_from_queue = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_get_budget = []leanruntime.Field{
	{Name: "BudgetId", Flag: "budget-id", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
}

var fields_get_farm = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
}

var fields_get_fleet = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_get_job = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_get_license_endpoint = []leanruntime.Field{
	{Name: "LicenseEndpointId", Flag: "license-endpoint-id", Type: "*string", Required: true},
}

var fields_get_limit = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
}

var fields_get_monitor = []leanruntime.Field{
	{Name: "MonitorId", Flag: "monitor-id", Type: "*string", Required: true},
}

var fields_get_queue = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_get_queue_environment = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueEnvironmentId", Flag: "queue-environment-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_get_queue_fleet_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_get_queue_limit_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_get_session = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
}

var fields_get_session_action = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "SessionActionId", Flag: "session-action-id", Type: "*string", Required: true},
}

var fields_get_sessions_statistics_aggregation = []leanruntime.Field{
	{Name: "AggregationId", Flag: "aggregation-id", Type: "*string", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_step = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
}

var fields_get_storage_profile = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "StorageProfileId", Flag: "storage-profile-id", Type: "*string", Required: true},
}

var fields_get_storage_profile_for_queue = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StorageProfileId", Flag: "storage-profile-id", Type: "*string", Required: true},
}

var fields_get_task = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_get_worker = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_list_available_metered_products = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_budgets = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BudgetStatus", Required: false},
}

var fields_list_farm_members = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_farms = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: false},
}

var fields_list_fleet_members = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fleets = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.FleetStatus", Required: false},
}

var fields_list_job_members = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_job_parameter_definitions = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_jobs = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_license_endpoints = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_limits = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_metered_products = []leanruntime.Field{
	{Name: "LicenseEndpointId", Flag: "license-endpoint-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_monitors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_queue_environments = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_queue_fleet_associations = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: false},
}

var fields_list_queue_limit_associations = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: false},
}

var fields_list_queue_members = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_queues = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PrincipalId", Flag: "principal-id", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.QueueStatus", Required: false},
}

var fields_list_session_actions = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: false},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: false},
}

var fields_list_sessions = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_sessions_for_worker = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_list_step_consumers = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
}

var fields_list_step_dependencies = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
}

var fields_list_steps = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_storage_profiles = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_storage_profiles_for_queue = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tasks = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
}

var fields_list_workers = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_metered_product = []leanruntime.Field{
	{Name: "LicenseEndpointId", Flag: "license-endpoint-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_search_jobs = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FilterExpressions", Flag: "filter-expressions", Type: "*types.SearchGroupedFilterExpressions", Required: false},
	{Name: "ItemOffset", Flag: "item-offset", Type: "*int32", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "QueueIds", Flag: "queue-ids", Type: "[]string", Required: true},
	{Name: "SortExpressions", Flag: "sort-expressions", Type: "[]types.SearchSortExpression", Required: false},
}

var fields_search_steps = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FilterExpressions", Flag: "filter-expressions", Type: "*types.SearchGroupedFilterExpressions", Required: false},
	{Name: "ItemOffset", Flag: "item-offset", Type: "*int32", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "QueueIds", Flag: "queue-ids", Type: "[]string", Required: true},
	{Name: "SortExpressions", Flag: "sort-expressions", Type: "[]types.SearchSortExpression", Required: false},
}

var fields_search_tasks = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FilterExpressions", Flag: "filter-expressions", Type: "*types.SearchGroupedFilterExpressions", Required: false},
	{Name: "ItemOffset", Flag: "item-offset", Type: "*int32", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "QueueIds", Flag: "queue-ids", Type: "[]string", Required: true},
	{Name: "SortExpressions", Flag: "sort-expressions", Type: "[]types.SearchSortExpression", Required: false},
}

var fields_search_workers = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FilterExpressions", Flag: "filter-expressions", Type: "*types.SearchGroupedFilterExpressions", Required: false},
	{Name: "FleetIds", Flag: "fleet-ids", Type: "[]string", Required: true},
	{Name: "ItemOffset", Flag: "item-offset", Type: "*int32", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
	{Name: "SortExpressions", Flag: "sort-expressions", Type: "[]types.SearchSortExpression", Required: false},
}

var fields_start_sessions_statistics_aggregation = []leanruntime.Field{
	{Name: "EndTime", Flag: "end-time", Type: "*time.Time", Required: true},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "GroupBy", Flag: "group-by", Type: "[]types.UsageGroupByField", Required: true},
	{Name: "Period", Flag: "period", Type: "types.Period", Required: false},
	{Name: "ResourceIds", Flag: "resource-ids", Type: "types.SessionsStatisticsResources", Required: true},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: true},
	{Name: "Statistics", Flag: "statistics", Type: "[]types.UsageStatistic", Required: true},
	{Name: "Timezone", Flag: "timezone", Type: "*string", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_budget = []leanruntime.Field{
	{Name: "ActionsToAdd", Flag: "actions-to-add", Type: "[]types.BudgetActionToAdd", Required: false},
	{Name: "ActionsToRemove", Flag: "actions-to-remove", Type: "[]types.BudgetActionToRemove", Required: false},
	{Name: "ApproximateDollarLimit", Flag: "approximate-dollar-limit", Type: "*float32", Required: false},
	{Name: "BudgetId", Flag: "budget-id", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "Schedule", Flag: "schedule", Type: "types.BudgetSchedule", Required: false},
	{Name: "Status", Flag: "status", Type: "types.BudgetStatus", Required: false},
}

var fields_update_farm = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
}

var fields_update_fleet = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.FleetConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "HostConfiguration", Flag: "host-configuration", Type: "*types.HostConfiguration", Required: false},
	{Name: "MaxWorkerCount", Flag: "max-worker-count", Type: "*int32", Required: false},
	{Name: "MinWorkerCount", Flag: "min-worker-count", Type: "*int32", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "LifecycleStatus", Flag: "lifecycle-status", Type: "types.UpdateJobLifecycleStatus", Required: false},
	{Name: "MaxFailedTasksCount", Flag: "max-failed-tasks-count", Type: "*int32", Required: false},
	{Name: "MaxRetriesPerTask", Flag: "max-retries-per-task", Type: "*int32", Required: false},
	{Name: "MaxWorkerCount", Flag: "max-worker-count", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "TargetTaskRunStatus", Flag: "target-task-run-status", Type: "types.JobTargetTaskRunStatus", Required: false},
}

var fields_update_limit = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
	{Name: "MaxCount", Flag: "max-count", Type: "*int32", Required: false},
}

var fields_update_monitor = []leanruntime.Field{
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "MonitorId", Flag: "monitor-id", Type: "*string", Required: true},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
	{Name: "Subdomain", Flag: "subdomain", Type: "*string", Required: false},
}

var fields_update_queue = []leanruntime.Field{
	{Name: "AllowedStorageProfileIdsToAdd", Flag: "allowed-storage-profile-ids-to-add", Type: "[]string", Required: false},
	{Name: "AllowedStorageProfileIdsToRemove", Flag: "allowed-storage-profile-ids-to-remove", Type: "[]string", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DefaultBudgetAction", Flag: "default-budget-action", Type: "types.DefaultQueueBudgetAction", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobAttachmentSettings", Flag: "job-attachment-settings", Type: "*types.JobAttachmentSettings", Required: false},
	{Name: "JobRunAsUser", Flag: "job-run-as-user", Type: "*types.JobRunAsUser", Required: false},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "RequiredFileSystemLocationNamesToAdd", Flag: "required-file-system-location-names-to-add", Type: "[]string", Required: false},
	{Name: "RequiredFileSystemLocationNamesToRemove", Flag: "required-file-system-location-names-to-remove", Type: "[]string", Required: false},
	{Name: "RoleArn", Flag: "role-arn", Type: "*string", Required: false},
}

var fields_update_queue_environment = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "QueueEnvironmentId", Flag: "queue-environment-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "Template", Flag: "template", Type: "*string", Required: false},
	{Name: "TemplateType", Flag: "template-type", Type: "types.EnvironmentTemplateType", Required: false},
}

var fields_update_queue_fleet_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.UpdateQueueFleetAssociationStatus", Required: true},
}

var fields_update_queue_limit_association = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "LimitId", Flag: "limit-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.UpdateQueueLimitAssociationStatus", Required: true},
}

var fields_update_session = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "TargetLifecycleStatus", Flag: "target-lifecycle-status", Type: "types.SessionLifecycleTargetStatus", Required: true},
}

var fields_update_step = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
	{Name: "TargetTaskRunStatus", Flag: "target-task-run-status", Type: "types.StepTargetTaskRunStatus", Required: true},
}

var fields_update_storage_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FileSystemLocationsToAdd", Flag: "file-system-locations-to-add", Type: "[]types.FileSystemLocation", Required: false},
	{Name: "FileSystemLocationsToRemove", Flag: "file-system-locations-to-remove", Type: "[]types.FileSystemLocation", Required: false},
	{Name: "OsFamily", Flag: "os-family", Type: "types.StorageProfileOperatingSystemFamily", Required: false},
	{Name: "StorageProfileId", Flag: "storage-profile-id", Type: "*string", Required: true},
}

var fields_update_task = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "QueueId", Flag: "queue-id", Type: "*string", Required: true},
	{Name: "StepId", Flag: "step-id", Type: "*string", Required: true},
	{Name: "TargetRunStatus", Flag: "target-run-status", Type: "types.TaskTargetRunStatus", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_update_worker = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "*types.WorkerCapabilities", Required: false},
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "HostProperties", Flag: "host-properties", Type: "*types.HostPropertiesRequest", Required: false},
	{Name: "Status", Flag: "status", Type: "types.UpdatedWorkerStatus", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_update_worker_schedule = []leanruntime.Field{
	{Name: "FarmId", Flag: "farm-id", Type: "*string", Required: true},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "UpdatedSessionActions", Flag: "updated-session-actions", Type: "map[string]types.UpdatedSessionActionInfo", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-member-to-farm": {
			Name:   "associate-member-to-farm",
			Fields: fields_associate_member_to_farm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMemberToFarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_member_to_farm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMemberToFarm(ctx, input)
			},
		},
		"associate-member-to-fleet": {
			Name:   "associate-member-to-fleet",
			Fields: fields_associate_member_to_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMemberToFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_member_to_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMemberToFleet(ctx, input)
			},
		},
		"associate-member-to-job": {
			Name:   "associate-member-to-job",
			Fields: fields_associate_member_to_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMemberToJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_member_to_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMemberToJob(ctx, input)
			},
		},
		"associate-member-to-queue": {
			Name:   "associate-member-to-queue",
			Fields: fields_associate_member_to_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateMemberToQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_member_to_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateMemberToQueue(ctx, input)
			},
		},
		"assume-fleet-role-for-read": {
			Name:   "assume-fleet-role-for-read",
			Fields: fields_assume_fleet_role_for_read,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeFleetRoleForReadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_fleet_role_for_read, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeFleetRoleForRead(ctx, input)
			},
		},
		"assume-fleet-role-for-worker": {
			Name:   "assume-fleet-role-for-worker",
			Fields: fields_assume_fleet_role_for_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeFleetRoleForWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_fleet_role_for_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeFleetRoleForWorker(ctx, input)
			},
		},
		"assume-queue-role-for-read": {
			Name:   "assume-queue-role-for-read",
			Fields: fields_assume_queue_role_for_read,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeQueueRoleForReadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_queue_role_for_read, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeQueueRoleForRead(ctx, input)
			},
		},
		"assume-queue-role-for-user": {
			Name:   "assume-queue-role-for-user",
			Fields: fields_assume_queue_role_for_user,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeQueueRoleForUserInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_queue_role_for_user, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeQueueRoleForUser(ctx, input)
			},
		},
		"assume-queue-role-for-worker": {
			Name:   "assume-queue-role-for-worker",
			Fields: fields_assume_queue_role_for_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssumeQueueRoleForWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_assume_queue_role_for_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssumeQueueRoleForWorker(ctx, input)
			},
		},
		"batch-get-job-entity": {
			Name:   "batch-get-job-entity",
			Fields: fields_batch_get_job_entity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetJobEntityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_job_entity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetJobEntity(ctx, input)
			},
		},
		"copy-job-template": {
			Name:   "copy-job-template",
			Fields: fields_copy_job_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyJobTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_job_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyJobTemplate(ctx, input)
			},
		},
		"create-budget": {
			Name:   "create-budget",
			Fields: fields_create_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBudget(ctx, input)
			},
		},
		"create-farm": {
			Name:   "create-farm",
			Fields: fields_create_farm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_farm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFarm(ctx, input)
			},
		},
		"create-fleet": {
			Name:   "create-fleet",
			Fields: fields_create_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleet(ctx, input)
			},
		},
		"create-job": {
			Name:   "create-job",
			Fields: fields_create_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateJob(ctx, input)
			},
		},
		"create-license-endpoint": {
			Name:   "create-license-endpoint",
			Fields: fields_create_license_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseEndpoint(ctx, input)
			},
		},
		"create-limit": {
			Name:   "create-limit",
			Fields: fields_create_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLimit(ctx, input)
			},
		},
		"create-monitor": {
			Name:   "create-monitor",
			Fields: fields_create_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMonitor(ctx, input)
			},
		},
		"create-queue": {
			Name:   "create-queue",
			Fields: fields_create_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueue(ctx, input)
			},
		},
		"create-queue-environment": {
			Name:   "create-queue-environment",
			Fields: fields_create_queue_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueueEnvironment(ctx, input)
			},
		},
		"create-queue-fleet-association": {
			Name:   "create-queue-fleet-association",
			Fields: fields_create_queue_fleet_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueFleetAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue_fleet_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueueFleetAssociation(ctx, input)
			},
		},
		"create-queue-limit-association": {
			Name:   "create-queue-limit-association",
			Fields: fields_create_queue_limit_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQueueLimitAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_queue_limit_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQueueLimitAssociation(ctx, input)
			},
		},
		"create-storage-profile": {
			Name:   "create-storage-profile",
			Fields: fields_create_storage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStorageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_storage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStorageProfile(ctx, input)
			},
		},
		"create-worker": {
			Name:   "create-worker",
			Fields: fields_create_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorker(ctx, input)
			},
		},
		"delete-budget": {
			Name:   "delete-budget",
			Fields: fields_delete_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBudget(ctx, input)
			},
		},
		"delete-farm": {
			Name:   "delete-farm",
			Fields: fields_delete_farm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_farm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFarm(ctx, input)
			},
		},
		"delete-fleet": {
			Name:   "delete-fleet",
			Fields: fields_delete_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleet(ctx, input)
			},
		},
		"delete-license-endpoint": {
			Name:   "delete-license-endpoint",
			Fields: fields_delete_license_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicenseEndpoint(ctx, input)
			},
		},
		"delete-limit": {
			Name:   "delete-limit",
			Fields: fields_delete_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLimit(ctx, input)
			},
		},
		"delete-metered-product": {
			Name:   "delete-metered-product",
			Fields: fields_delete_metered_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMeteredProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_metered_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMeteredProduct(ctx, input)
			},
		},
		"delete-monitor": {
			Name:   "delete-monitor",
			Fields: fields_delete_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMonitor(ctx, input)
			},
		},
		"delete-queue": {
			Name:   "delete-queue",
			Fields: fields_delete_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueue(ctx, input)
			},
		},
		"delete-queue-environment": {
			Name:   "delete-queue-environment",
			Fields: fields_delete_queue_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueueEnvironment(ctx, input)
			},
		},
		"delete-queue-fleet-association": {
			Name:   "delete-queue-fleet-association",
			Fields: fields_delete_queue_fleet_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueFleetAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue_fleet_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueueFleetAssociation(ctx, input)
			},
		},
		"delete-queue-limit-association": {
			Name:   "delete-queue-limit-association",
			Fields: fields_delete_queue_limit_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQueueLimitAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_queue_limit_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQueueLimitAssociation(ctx, input)
			},
		},
		"delete-storage-profile": {
			Name:   "delete-storage-profile",
			Fields: fields_delete_storage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStorageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_storage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStorageProfile(ctx, input)
			},
		},
		"delete-worker": {
			Name:   "delete-worker",
			Fields: fields_delete_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorker(ctx, input)
			},
		},
		"disassociate-member-from-farm": {
			Name:   "disassociate-member-from-farm",
			Fields: fields_disassociate_member_from_farm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMemberFromFarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_member_from_farm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMemberFromFarm(ctx, input)
			},
		},
		"disassociate-member-from-fleet": {
			Name:   "disassociate-member-from-fleet",
			Fields: fields_disassociate_member_from_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMemberFromFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_member_from_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMemberFromFleet(ctx, input)
			},
		},
		"disassociate-member-from-job": {
			Name:   "disassociate-member-from-job",
			Fields: fields_disassociate_member_from_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMemberFromJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_member_from_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMemberFromJob(ctx, input)
			},
		},
		"disassociate-member-from-queue": {
			Name:   "disassociate-member-from-queue",
			Fields: fields_disassociate_member_from_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateMemberFromQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_member_from_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateMemberFromQueue(ctx, input)
			},
		},
		"get-budget": {
			Name:   "get-budget",
			Fields: fields_get_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBudget(ctx, input)
			},
		},
		"get-farm": {
			Name:   "get-farm",
			Fields: fields_get_farm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_farm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFarm(ctx, input)
			},
		},
		"get-fleet": {
			Name:   "get-fleet",
			Fields: fields_get_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFleet(ctx, input)
			},
		},
		"get-job": {
			Name:   "get-job",
			Fields: fields_get_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJob(ctx, input)
			},
		},
		"get-license-endpoint": {
			Name:   "get-license-endpoint",
			Fields: fields_get_license_endpoint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseEndpointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_endpoint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseEndpoint(ctx, input)
			},
		},
		"get-limit": {
			Name:   "get-limit",
			Fields: fields_get_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLimit(ctx, input)
			},
		},
		"get-monitor": {
			Name:   "get-monitor",
			Fields: fields_get_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMonitor(ctx, input)
			},
		},
		"get-queue": {
			Name:   "get-queue",
			Fields: fields_get_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueue(ctx, input)
			},
		},
		"get-queue-environment": {
			Name:   "get-queue-environment",
			Fields: fields_get_queue_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueueEnvironment(ctx, input)
			},
		},
		"get-queue-fleet-association": {
			Name:   "get-queue-fleet-association",
			Fields: fields_get_queue_fleet_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueFleetAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue_fleet_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueueFleetAssociation(ctx, input)
			},
		},
		"get-queue-limit-association": {
			Name:   "get-queue-limit-association",
			Fields: fields_get_queue_limit_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQueueLimitAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_queue_limit_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQueueLimitAssociation(ctx, input)
			},
		},
		"get-session": {
			Name:   "get-session",
			Fields: fields_get_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSession(ctx, input)
			},
		},
		"get-session-action": {
			Name:   "get-session-action",
			Fields: fields_get_session_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_session_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSessionAction(ctx, input)
			},
		},
		"get-sessions-statistics-aggregation": {
			Name:   "get-sessions-statistics-aggregation",
			Fields: fields_get_sessions_statistics_aggregation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSessionsStatisticsAggregationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_sessions_statistics_aggregation, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSessionsStatisticsAggregation(ctx, input)
				}
				var results []*svc.GetSessionsStatisticsAggregationOutput
				p := svc.NewGetSessionsStatisticsAggregationPaginator(client, input)
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
		"get-step": {
			Name:   "get-step",
			Fields: fields_get_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStep(ctx, input)
			},
		},
		"get-storage-profile": {
			Name:   "get-storage-profile",
			Fields: fields_get_storage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStorageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_storage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStorageProfile(ctx, input)
			},
		},
		"get-storage-profile-for-queue": {
			Name:   "get-storage-profile-for-queue",
			Fields: fields_get_storage_profile_for_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStorageProfileForQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_storage_profile_for_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStorageProfileForQueue(ctx, input)
			},
		},
		"get-task": {
			Name:   "get-task",
			Fields: fields_get_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTask(ctx, input)
			},
		},
		"get-worker": {
			Name:   "get-worker",
			Fields: fields_get_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetWorker(ctx, input)
			},
		},
		"list-available-metered-products": {
			Name:   "list-available-metered-products",
			Fields: fields_list_available_metered_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAvailableMeteredProductsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_available_metered_products, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAvailableMeteredProducts(ctx, input)
				}
				var results []*svc.ListAvailableMeteredProductsOutput
				p := svc.NewListAvailableMeteredProductsPaginator(client, input)
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
		"list-budgets": {
			Name:   "list-budgets",
			Fields: fields_list_budgets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBudgetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_budgets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBudgets(ctx, input)
				}
				var results []*svc.ListBudgetsOutput
				p := svc.NewListBudgetsPaginator(client, input)
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
		"list-farm-members": {
			Name:   "list-farm-members",
			Fields: fields_list_farm_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFarmMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_farm_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFarmMembers(ctx, input)
				}
				var results []*svc.ListFarmMembersOutput
				p := svc.NewListFarmMembersPaginator(client, input)
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
		"list-farms": {
			Name:   "list-farms",
			Fields: fields_list_farms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFarmsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_farms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFarms(ctx, input)
				}
				var results []*svc.ListFarmsOutput
				p := svc.NewListFarmsPaginator(client, input)
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
		"list-fleet-members": {
			Name:   "list-fleet-members",
			Fields: fields_list_fleet_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleet_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleetMembers(ctx, input)
				}
				var results []*svc.ListFleetMembersOutput
				p := svc.NewListFleetMembersPaginator(client, input)
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
		"list-fleets": {
			Name:   "list-fleets",
			Fields: fields_list_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleets(ctx, input)
				}
				var results []*svc.ListFleetsOutput
				p := svc.NewListFleetsPaginator(client, input)
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
		"list-job-members": {
			Name:   "list-job-members",
			Fields: fields_list_job_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobMembers(ctx, input)
				}
				var results []*svc.ListJobMembersOutput
				p := svc.NewListJobMembersPaginator(client, input)
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
		"list-job-parameter-definitions": {
			Name:   "list-job-parameter-definitions",
			Fields: fields_list_job_parameter_definitions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobParameterDefinitionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_parameter_definitions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobParameterDefinitions(ctx, input)
				}
				var results []*svc.ListJobParameterDefinitionsOutput
				p := svc.NewListJobParameterDefinitionsPaginator(client, input)
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
		"list-jobs": {
			Name:   "list-jobs",
			Fields: fields_list_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobs(ctx, input)
				}
				var results []*svc.ListJobsOutput
				p := svc.NewListJobsPaginator(client, input)
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
		"list-license-endpoints": {
			Name:   "list-license-endpoints",
			Fields: fields_list_license_endpoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseEndpointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_license_endpoints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLicenseEndpoints(ctx, input)
				}
				var results []*svc.ListLicenseEndpointsOutput
				p := svc.NewListLicenseEndpointsPaginator(client, input)
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
		"list-limits": {
			Name:   "list-limits",
			Fields: fields_list_limits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLimitsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_limits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLimits(ctx, input)
				}
				var results []*svc.ListLimitsOutput
				p := svc.NewListLimitsPaginator(client, input)
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
		"list-metered-products": {
			Name:   "list-metered-products",
			Fields: fields_list_metered_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMeteredProductsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metered_products, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMeteredProducts(ctx, input)
				}
				var results []*svc.ListMeteredProductsOutput
				p := svc.NewListMeteredProductsPaginator(client, input)
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
		"list-monitors": {
			Name:   "list-monitors",
			Fields: fields_list_monitors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMonitorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_monitors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMonitors(ctx, input)
				}
				var results []*svc.ListMonitorsOutput
				p := svc.NewListMonitorsPaginator(client, input)
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
		"list-queue-environments": {
			Name:   "list-queue-environments",
			Fields: fields_list_queue_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueueEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queue_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueueEnvironments(ctx, input)
				}
				var results []*svc.ListQueueEnvironmentsOutput
				p := svc.NewListQueueEnvironmentsPaginator(client, input)
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
		"list-queue-fleet-associations": {
			Name:   "list-queue-fleet-associations",
			Fields: fields_list_queue_fleet_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueueFleetAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queue_fleet_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueueFleetAssociations(ctx, input)
				}
				var results []*svc.ListQueueFleetAssociationsOutput
				p := svc.NewListQueueFleetAssociationsPaginator(client, input)
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
		"list-queue-limit-associations": {
			Name:   "list-queue-limit-associations",
			Fields: fields_list_queue_limit_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueueLimitAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queue_limit_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueueLimitAssociations(ctx, input)
				}
				var results []*svc.ListQueueLimitAssociationsOutput
				p := svc.NewListQueueLimitAssociationsPaginator(client, input)
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
		"list-queue-members": {
			Name:   "list-queue-members",
			Fields: fields_list_queue_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueueMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queue_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueueMembers(ctx, input)
				}
				var results []*svc.ListQueueMembersOutput
				p := svc.NewListQueueMembersPaginator(client, input)
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
		"list-queues": {
			Name:   "list-queues",
			Fields: fields_list_queues,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQueuesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_queues, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQueues(ctx, input)
				}
				var results []*svc.ListQueuesOutput
				p := svc.NewListQueuesPaginator(client, input)
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
		"list-session-actions": {
			Name:   "list-session-actions",
			Fields: fields_list_session_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_session_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessionActions(ctx, input)
				}
				var results []*svc.ListSessionActionsOutput
				p := svc.NewListSessionActionsPaginator(client, input)
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
		"list-sessions": {
			Name:   "list-sessions",
			Fields: fields_list_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessions(ctx, input)
				}
				var results []*svc.ListSessionsOutput
				p := svc.NewListSessionsPaginator(client, input)
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
		"list-sessions-for-worker": {
			Name:   "list-sessions-for-worker",
			Fields: fields_list_sessions_for_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSessionsForWorkerInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_sessions_for_worker, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSessionsForWorker(ctx, input)
				}
				var results []*svc.ListSessionsForWorkerOutput
				p := svc.NewListSessionsForWorkerPaginator(client, input)
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
		"list-step-consumers": {
			Name:   "list-step-consumers",
			Fields: fields_list_step_consumers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStepConsumersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_step_consumers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStepConsumers(ctx, input)
				}
				var results []*svc.ListStepConsumersOutput
				p := svc.NewListStepConsumersPaginator(client, input)
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
		"list-step-dependencies": {
			Name:   "list-step-dependencies",
			Fields: fields_list_step_dependencies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStepDependenciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_step_dependencies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStepDependencies(ctx, input)
				}
				var results []*svc.ListStepDependenciesOutput
				p := svc.NewListStepDependenciesPaginator(client, input)
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
		"list-steps": {
			Name:   "list-steps",
			Fields: fields_list_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStepsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_steps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSteps(ctx, input)
				}
				var results []*svc.ListStepsOutput
				p := svc.NewListStepsPaginator(client, input)
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
		"list-storage-profiles": {
			Name:   "list-storage-profiles",
			Fields: fields_list_storage_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStorageProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_storage_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStorageProfiles(ctx, input)
				}
				var results []*svc.ListStorageProfilesOutput
				p := svc.NewListStorageProfilesPaginator(client, input)
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
		"list-storage-profiles-for-queue": {
			Name:   "list-storage-profiles-for-queue",
			Fields: fields_list_storage_profiles_for_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStorageProfilesForQueueInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_storage_profiles_for_queue, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStorageProfilesForQueue(ctx, input)
				}
				var results []*svc.ListStorageProfilesForQueueOutput
				p := svc.NewListStorageProfilesForQueuePaginator(client, input)
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
		"list-tasks": {
			Name:   "list-tasks",
			Fields: fields_list_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTasksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tasks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTasks(ctx, input)
				}
				var results []*svc.ListTasksOutput
				p := svc.NewListTasksPaginator(client, input)
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
		"list-workers": {
			Name:   "list-workers",
			Fields: fields_list_workers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkers(ctx, input)
				}
				var results []*svc.ListWorkersOutput
				p := svc.NewListWorkersPaginator(client, input)
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
		"put-metered-product": {
			Name:   "put-metered-product",
			Fields: fields_put_metered_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutMeteredProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_metered_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutMeteredProduct(ctx, input)
			},
		},
		"search-jobs": {
			Name:   "search-jobs",
			Fields: fields_search_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchJobsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_jobs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchJobs(ctx, input)
			},
		},
		"search-steps": {
			Name:   "search-steps",
			Fields: fields_search_steps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchStepsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_steps, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchSteps(ctx, input)
			},
		},
		"search-tasks": {
			Name:   "search-tasks",
			Fields: fields_search_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchTasks(ctx, input)
			},
		},
		"search-workers": {
			Name:   "search-workers",
			Fields: fields_search_workers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchWorkersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_workers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchWorkers(ctx, input)
			},
		},
		"start-sessions-statistics-aggregation": {
			Name:   "start-sessions-statistics-aggregation",
			Fields: fields_start_sessions_statistics_aggregation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSessionsStatisticsAggregationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_sessions_statistics_aggregation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSessionsStatisticsAggregation(ctx, input)
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
		"update-budget": {
			Name:   "update-budget",
			Fields: fields_update_budget,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateBudgetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_budget, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateBudget(ctx, input)
			},
		},
		"update-farm": {
			Name:   "update-farm",
			Fields: fields_update_farm,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFarmInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_farm, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFarm(ctx, input)
			},
		},
		"update-fleet": {
			Name:   "update-fleet",
			Fields: fields_update_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleet(ctx, input)
			},
		},
		"update-job": {
			Name:   "update-job",
			Fields: fields_update_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateJob(ctx, input)
			},
		},
		"update-limit": {
			Name:   "update-limit",
			Fields: fields_update_limit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLimitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_limit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLimit(ctx, input)
			},
		},
		"update-monitor": {
			Name:   "update-monitor",
			Fields: fields_update_monitor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMonitorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_monitor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMonitor(ctx, input)
			},
		},
		"update-queue": {
			Name:   "update-queue",
			Fields: fields_update_queue,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueue(ctx, input)
			},
		},
		"update-queue-environment": {
			Name:   "update-queue-environment",
			Fields: fields_update_queue_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueEnvironment(ctx, input)
			},
		},
		"update-queue-fleet-association": {
			Name:   "update-queue-fleet-association",
			Fields: fields_update_queue_fleet_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueFleetAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_fleet_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueFleetAssociation(ctx, input)
			},
		},
		"update-queue-limit-association": {
			Name:   "update-queue-limit-association",
			Fields: fields_update_queue_limit_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQueueLimitAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_queue_limit_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQueueLimitAssociation(ctx, input)
			},
		},
		"update-session": {
			Name:   "update-session",
			Fields: fields_update_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSession(ctx, input)
			},
		},
		"update-step": {
			Name:   "update-step",
			Fields: fields_update_step,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStepInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_step, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStep(ctx, input)
			},
		},
		"update-storage-profile": {
			Name:   "update-storage-profile",
			Fields: fields_update_storage_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStorageProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_storage_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStorageProfile(ctx, input)
			},
		},
		"update-task": {
			Name:   "update-task",
			Fields: fields_update_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTask(ctx, input)
			},
		},
		"update-worker": {
			Name:   "update-worker",
			Fields: fields_update_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorker(ctx, input)
			},
		},
		"update-worker-schedule": {
			Name:   "update-worker-schedule",
			Fields: fields_update_worker_schedule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWorkerScheduleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_worker_schedule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWorkerSchedule(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("deadline", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
