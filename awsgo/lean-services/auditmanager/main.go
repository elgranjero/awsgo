package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/auditmanager"
)

var fields_associate_assessment_report_evidence_folder = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
}

var fields_batch_associate_assessment_report_evidence = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
	{Name: "EvidenceIds", Flag: "evidence-ids", Type: "[]string", Required: true},
}

var fields_batch_create_delegation_by_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "CreateDelegationRequests", Flag: "create-delegation-requests", Type: "[]types.CreateDelegationRequest", Required: true},
}

var fields_batch_delete_delegation_by_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "DelegationIds", Flag: "delegation-ids", Type: "[]string", Required: true},
}

var fields_batch_disassociate_assessment_report_evidence = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
	{Name: "EvidenceIds", Flag: "evidence-ids", Type: "[]string", Required: true},
}

var fields_batch_import_evidence_to_assessment_control = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "ManualEvidence", Flag: "manual-evidence", Type: "[]types.ManualEvidence", Required: true},
}

var fields_create_assessment = []leanruntime.Field{
	{Name: "AssessmentReportsDestination", Flag: "assessment-reports-destination", Type: "*types.AssessmentReportsDestination", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FrameworkId", Flag: "framework-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Roles", Flag: "roles", Type: "[]types.Role", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_assessment_framework = []leanruntime.Field{
	{Name: "ComplianceType", Flag: "compliance-type", Type: "*string", Required: false},
	{Name: "ControlSets", Flag: "control-sets", Type: "[]types.CreateAssessmentFrameworkControlSet", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_assessment_report = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QueryStatement", Flag: "query-statement", Type: "*string", Required: false},
}

var fields_create_control = []leanruntime.Field{
	{Name: "ActionPlanInstructions", Flag: "action-plan-instructions", Type: "*string", Required: false},
	{Name: "ActionPlanTitle", Flag: "action-plan-title", Type: "*string", Required: false},
	{Name: "ControlMappingSources", Flag: "control-mapping-sources", Type: "[]types.CreateControlMappingSource", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TestingInformation", Flag: "testing-information", Type: "*string", Required: false},
}

var fields_delete_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
}

var fields_delete_assessment_framework = []leanruntime.Field{
	{Name: "FrameworkId", Flag: "framework-id", Type: "*string", Required: true},
}

var fields_delete_assessment_framework_share = []leanruntime.Field{
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "RequestType", Flag: "request-type", Type: "types.ShareRequestType", Required: true},
}

var fields_delete_assessment_report = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "AssessmentReportId", Flag: "assessment-report-id", Type: "*string", Required: true},
}

var fields_delete_control = []leanruntime.Field{
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
}

var fields_deregister_account = []leanruntime.Field{}

var fields_deregister_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: false},
}

var fields_disassociate_assessment_report_evidence_folder = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
}

var fields_get_account_status = []leanruntime.Field{}

var fields_get_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
}

var fields_get_assessment_framework = []leanruntime.Field{
	{Name: "FrameworkId", Flag: "framework-id", Type: "*string", Required: true},
}

var fields_get_assessment_report_url = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "AssessmentReportId", Flag: "assessment-report-id", Type: "*string", Required: true},
}

var fields_get_change_logs = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: false},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_control = []leanruntime.Field{
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
}

var fields_get_delegations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_evidence = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
	{Name: "EvidenceId", Flag: "evidence-id", Type: "*string", Required: true},
}

var fields_get_evidence_by_evidence_folder = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_evidence_file_upload_url = []leanruntime.Field{
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: true},
}

var fields_get_evidence_folder = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "EvidenceFolderId", Flag: "evidence-folder-id", Type: "*string", Required: true},
}

var fields_get_evidence_folders_by_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_evidence_folders_by_assessment_control = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_insights = []leanruntime.Field{}

var fields_get_insights_by_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
}

var fields_get_organization_admin_account = []leanruntime.Field{}

var fields_get_services_in_scope = []leanruntime.Field{}

var fields_get_settings = []leanruntime.Field{
	{Name: "Attribute", Flag: "attribute", Type: "types.SettingAttribute", Required: true},
}

var fields_list_assessment_control_insights_by_control_domain = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "ControlDomainId", Flag: "control-domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assessment_framework_share_requests = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RequestType", Flag: "request-type", Type: "types.ShareRequestType", Required: true},
}

var fields_list_assessment_frameworks = []leanruntime.Field{
	{Name: "FrameworkType", Flag: "framework-type", Type: "types.FrameworkType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assessment_reports = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_assessments = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.AssessmentStatus", Required: false},
}

var fields_list_control_domain_insights = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_control_domain_insights_by_assessment = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_control_insights_by_control_domain = []leanruntime.Field{
	{Name: "ControlDomainId", Flag: "control-domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_controls = []leanruntime.Field{
	{Name: "ControlCatalogId", Flag: "control-catalog-id", Type: "*string", Required: false},
	{Name: "ControlType", Flag: "control-type", Type: "types.ControlType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_keywords_for_data_source = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "types.DataSourceType", Required: true},
}

var fields_list_notifications = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_register_account = []leanruntime.Field{
	{Name: "DelegatedAdminAccount", Flag: "delegated-admin-account", Type: "*string", Required: false},
	{Name: "KmsKey", Flag: "kms-key", Type: "*string", Required: false},
}

var fields_register_organization_admin_account = []leanruntime.Field{
	{Name: "AdminAccountId", Flag: "admin-account-id", Type: "*string", Required: true},
}

var fields_start_assessment_framework_share = []leanruntime.Field{
	{Name: "Comment", Flag: "comment", Type: "*string", Required: false},
	{Name: "DestinationAccount", Flag: "destination-account", Type: "*string", Required: true},
	{Name: "DestinationRegion", Flag: "destination-region", Type: "*string", Required: true},
	{Name: "FrameworkId", Flag: "framework-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_assessment = []leanruntime.Field{
	{Name: "AssessmentDescription", Flag: "assessment-description", Type: "*string", Required: false},
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "AssessmentName", Flag: "assessment-name", Type: "*string", Required: false},
	{Name: "AssessmentReportsDestination", Flag: "assessment-reports-destination", Type: "*types.AssessmentReportsDestination", Required: false},
	{Name: "Roles", Flag: "roles", Type: "[]types.Role", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*types.Scope", Required: true},
}

var fields_update_assessment_control = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "CommentBody", Flag: "comment-body", Type: "*string", Required: false},
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "ControlStatus", Flag: "control-status", Type: "types.ControlStatus", Required: false},
}

var fields_update_assessment_control_set_status = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "Comment", Flag: "comment", Type: "*string", Required: true},
	{Name: "ControlSetId", Flag: "control-set-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.ControlSetStatus", Required: true},
}

var fields_update_assessment_framework = []leanruntime.Field{
	{Name: "ComplianceType", Flag: "compliance-type", Type: "*string", Required: false},
	{Name: "ControlSets", Flag: "control-sets", Type: "[]types.UpdateAssessmentFrameworkControlSet", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FrameworkId", Flag: "framework-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_assessment_framework_share = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.ShareRequestAction", Required: true},
	{Name: "RequestId", Flag: "request-id", Type: "*string", Required: true},
	{Name: "RequestType", Flag: "request-type", Type: "types.ShareRequestType", Required: true},
}

var fields_update_assessment_status = []leanruntime.Field{
	{Name: "AssessmentId", Flag: "assessment-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.AssessmentStatus", Required: true},
}

var fields_update_control = []leanruntime.Field{
	{Name: "ActionPlanInstructions", Flag: "action-plan-instructions", Type: "*string", Required: false},
	{Name: "ActionPlanTitle", Flag: "action-plan-title", Type: "*string", Required: false},
	{Name: "ControlId", Flag: "control-id", Type: "*string", Required: true},
	{Name: "ControlMappingSources", Flag: "control-mapping-sources", Type: "[]types.ControlMappingSource", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "TestingInformation", Flag: "testing-information", Type: "*string", Required: false},
}

var fields_update_settings = []leanruntime.Field{
	{Name: "DefaultAssessmentReportsDestination", Flag: "default-assessment-reports-destination", Type: "*types.AssessmentReportsDestination", Required: false},
	{Name: "DefaultExportDestination", Flag: "default-export-destination", Type: "*types.DefaultExportDestination", Required: false},
	{Name: "DefaultProcessOwners", Flag: "default-process-owners", Type: "[]types.Role", Required: false},
	{Name: "DeregistrationPolicy", Flag: "deregistration-policy", Type: "*types.DeregistrationPolicy", Required: false},
	{Name: "EvidenceFinderEnabled", Flag: "evidence-finder-enabled", Type: "*bool", Required: false},
	{Name: "KmsKey", Flag: "kms-key", Type: "*string", Required: false},
	{Name: "SnsTopic", Flag: "sns-topic", Type: "*string", Required: false},
}

var fields_validate_assessment_report_integrity = []leanruntime.Field{
	{Name: "S3RelativePath", Flag: "s3-relative-path", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-assessment-report-evidence-folder": {
			Name:   "associate-assessment-report-evidence-folder",
			Fields: fields_associate_assessment_report_evidence_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAssessmentReportEvidenceFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_assessment_report_evidence_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAssessmentReportEvidenceFolder(ctx, input)
			},
		},
		"batch-associate-assessment-report-evidence": {
			Name:   "batch-associate-assessment-report-evidence",
			Fields: fields_batch_associate_assessment_report_evidence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateAssessmentReportEvidenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_assessment_report_evidence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateAssessmentReportEvidence(ctx, input)
			},
		},
		"batch-create-delegation-by-assessment": {
			Name:   "batch-create-delegation-by-assessment",
			Fields: fields_batch_create_delegation_by_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateDelegationByAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_delegation_by_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateDelegationByAssessment(ctx, input)
			},
		},
		"batch-delete-delegation-by-assessment": {
			Name:   "batch-delete-delegation-by-assessment",
			Fields: fields_batch_delete_delegation_by_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteDelegationByAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_delegation_by_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteDelegationByAssessment(ctx, input)
			},
		},
		"batch-disassociate-assessment-report-evidence": {
			Name:   "batch-disassociate-assessment-report-evidence",
			Fields: fields_batch_disassociate_assessment_report_evidence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateAssessmentReportEvidenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_assessment_report_evidence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateAssessmentReportEvidence(ctx, input)
			},
		},
		"batch-import-evidence-to-assessment-control": {
			Name:   "batch-import-evidence-to-assessment-control",
			Fields: fields_batch_import_evidence_to_assessment_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchImportEvidenceToAssessmentControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_import_evidence_to_assessment_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchImportEvidenceToAssessmentControl(ctx, input)
			},
		},
		"create-assessment": {
			Name:   "create-assessment",
			Fields: fields_create_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssessment(ctx, input)
			},
		},
		"create-assessment-framework": {
			Name:   "create-assessment-framework",
			Fields: fields_create_assessment_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssessmentFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assessment_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssessmentFramework(ctx, input)
			},
		},
		"create-assessment-report": {
			Name:   "create-assessment-report",
			Fields: fields_create_assessment_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssessmentReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_assessment_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssessmentReport(ctx, input)
			},
		},
		"create-control": {
			Name:   "create-control",
			Fields: fields_create_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateControl(ctx, input)
			},
		},
		"delete-assessment": {
			Name:   "delete-assessment",
			Fields: fields_delete_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessment(ctx, input)
			},
		},
		"delete-assessment-framework": {
			Name:   "delete-assessment-framework",
			Fields: fields_delete_assessment_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessmentFramework(ctx, input)
			},
		},
		"delete-assessment-framework-share": {
			Name:   "delete-assessment-framework-share",
			Fields: fields_delete_assessment_framework_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentFrameworkShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment_framework_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessmentFrameworkShare(ctx, input)
			},
		},
		"delete-assessment-report": {
			Name:   "delete-assessment-report",
			Fields: fields_delete_assessment_report,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssessmentReportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_assessment_report, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssessmentReport(ctx, input)
			},
		},
		"delete-control": {
			Name:   "delete-control",
			Fields: fields_delete_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteControl(ctx, input)
			},
		},
		"deregister-account": {
			Name:   "deregister-account",
			Fields: fields_deregister_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterAccount(ctx, input)
			},
		},
		"deregister-organization-admin-account": {
			Name:   "deregister-organization-admin-account",
			Fields: fields_deregister_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterOrganizationAdminAccount(ctx, input)
			},
		},
		"disassociate-assessment-report-evidence-folder": {
			Name:   "disassociate-assessment-report-evidence-folder",
			Fields: fields_disassociate_assessment_report_evidence_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAssessmentReportEvidenceFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_assessment_report_evidence_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAssessmentReportEvidenceFolder(ctx, input)
			},
		},
		"get-account-status": {
			Name:   "get-account-status",
			Fields: fields_get_account_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountStatus(ctx, input)
			},
		},
		"get-assessment": {
			Name:   "get-assessment",
			Fields: fields_get_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssessment(ctx, input)
			},
		},
		"get-assessment-framework": {
			Name:   "get-assessment-framework",
			Fields: fields_get_assessment_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssessmentFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assessment_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssessmentFramework(ctx, input)
			},
		},
		"get-assessment-report-url": {
			Name:   "get-assessment-report-url",
			Fields: fields_get_assessment_report_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssessmentReportUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assessment_report_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssessmentReportUrl(ctx, input)
			},
		},
		"get-change-logs": {
			Name:   "get-change-logs",
			Fields: fields_get_change_logs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetChangeLogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_change_logs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetChangeLogs(ctx, input)
				}
				var results []*svc.GetChangeLogsOutput
				p := svc.NewGetChangeLogsPaginator(client, input)
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
		"get-control": {
			Name:   "get-control",
			Fields: fields_get_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetControl(ctx, input)
			},
		},
		"get-delegations": {
			Name:   "get-delegations",
			Fields: fields_get_delegations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDelegationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_delegations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDelegations(ctx, input)
				}
				var results []*svc.GetDelegationsOutput
				p := svc.NewGetDelegationsPaginator(client, input)
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
		"get-evidence": {
			Name:   "get-evidence",
			Fields: fields_get_evidence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvidenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_evidence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvidence(ctx, input)
			},
		},
		"get-evidence-by-evidence-folder": {
			Name:   "get-evidence-by-evidence-folder",
			Fields: fields_get_evidence_by_evidence_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvidenceByEvidenceFolderInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_evidence_by_evidence_folder, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEvidenceByEvidenceFolder(ctx, input)
				}
				var results []*svc.GetEvidenceByEvidenceFolderOutput
				p := svc.NewGetEvidenceByEvidenceFolderPaginator(client, input)
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
		"get-evidence-file-upload-url": {
			Name:   "get-evidence-file-upload-url",
			Fields: fields_get_evidence_file_upload_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvidenceFileUploadUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_evidence_file_upload_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvidenceFileUploadUrl(ctx, input)
			},
		},
		"get-evidence-folder": {
			Name:   "get-evidence-folder",
			Fields: fields_get_evidence_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvidenceFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_evidence_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEvidenceFolder(ctx, input)
			},
		},
		"get-evidence-folders-by-assessment": {
			Name:   "get-evidence-folders-by-assessment",
			Fields: fields_get_evidence_folders_by_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvidenceFoldersByAssessmentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_evidence_folders_by_assessment, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEvidenceFoldersByAssessment(ctx, input)
				}
				var results []*svc.GetEvidenceFoldersByAssessmentOutput
				p := svc.NewGetEvidenceFoldersByAssessmentPaginator(client, input)
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
		"get-evidence-folders-by-assessment-control": {
			Name:   "get-evidence-folders-by-assessment-control",
			Fields: fields_get_evidence_folders_by_assessment_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEvidenceFoldersByAssessmentControlInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_evidence_folders_by_assessment_control, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetEvidenceFoldersByAssessmentControl(ctx, input)
				}
				var results []*svc.GetEvidenceFoldersByAssessmentControlOutput
				p := svc.NewGetEvidenceFoldersByAssessmentControlPaginator(client, input)
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
		"get-insights": {
			Name:   "get-insights",
			Fields: fields_get_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insights, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsights(ctx, input)
			},
		},
		"get-insights-by-assessment": {
			Name:   "get-insights-by-assessment",
			Fields: fields_get_insights_by_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInsightsByAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_insights_by_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInsightsByAssessment(ctx, input)
			},
		},
		"get-organization-admin-account": {
			Name:   "get-organization-admin-account",
			Fields: fields_get_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOrganizationAdminAccount(ctx, input)
			},
		},
		"get-services-in-scope": {
			Name:   "get-services-in-scope",
			Fields: fields_get_services_in_scope,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServicesInScopeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_services_in_scope, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServicesInScope(ctx, input)
			},
		},
		"get-settings": {
			Name:   "get-settings",
			Fields: fields_get_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSettings(ctx, input)
			},
		},
		"list-assessment-control-insights-by-control-domain": {
			Name:   "list-assessment-control-insights-by-control-domain",
			Fields: fields_list_assessment_control_insights_by_control_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentControlInsightsByControlDomainInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_control_insights_by_control_domain, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentControlInsightsByControlDomain(ctx, input)
				}
				var results []*svc.ListAssessmentControlInsightsByControlDomainOutput
				p := svc.NewListAssessmentControlInsightsByControlDomainPaginator(client, input)
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
		"list-assessment-framework-share-requests": {
			Name:   "list-assessment-framework-share-requests",
			Fields: fields_list_assessment_framework_share_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentFrameworkShareRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_framework_share_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentFrameworkShareRequests(ctx, input)
				}
				var results []*svc.ListAssessmentFrameworkShareRequestsOutput
				p := svc.NewListAssessmentFrameworkShareRequestsPaginator(client, input)
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
		"list-assessment-frameworks": {
			Name:   "list-assessment-frameworks",
			Fields: fields_list_assessment_frameworks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentFrameworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_frameworks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentFrameworks(ctx, input)
				}
				var results []*svc.ListAssessmentFrameworksOutput
				p := svc.NewListAssessmentFrameworksPaginator(client, input)
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
		"list-assessment-reports": {
			Name:   "list-assessment-reports",
			Fields: fields_list_assessment_reports,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentReportsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessment_reports, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessmentReports(ctx, input)
				}
				var results []*svc.ListAssessmentReportsOutput
				p := svc.NewListAssessmentReportsPaginator(client, input)
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
		"list-assessments": {
			Name:   "list-assessments",
			Fields: fields_list_assessments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssessmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assessments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssessments(ctx, input)
				}
				var results []*svc.ListAssessmentsOutput
				p := svc.NewListAssessmentsPaginator(client, input)
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
		"list-control-domain-insights": {
			Name:   "list-control-domain-insights",
			Fields: fields_list_control_domain_insights,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlDomainInsightsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_control_domain_insights, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControlDomainInsights(ctx, input)
				}
				var results []*svc.ListControlDomainInsightsOutput
				p := svc.NewListControlDomainInsightsPaginator(client, input)
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
		"list-control-domain-insights-by-assessment": {
			Name:   "list-control-domain-insights-by-assessment",
			Fields: fields_list_control_domain_insights_by_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlDomainInsightsByAssessmentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_control_domain_insights_by_assessment, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControlDomainInsightsByAssessment(ctx, input)
				}
				var results []*svc.ListControlDomainInsightsByAssessmentOutput
				p := svc.NewListControlDomainInsightsByAssessmentPaginator(client, input)
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
		"list-control-insights-by-control-domain": {
			Name:   "list-control-insights-by-control-domain",
			Fields: fields_list_control_insights_by_control_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlInsightsByControlDomainInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_control_insights_by_control_domain, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControlInsightsByControlDomain(ctx, input)
				}
				var results []*svc.ListControlInsightsByControlDomainOutput
				p := svc.NewListControlInsightsByControlDomainPaginator(client, input)
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
		"list-controls": {
			Name:   "list-controls",
			Fields: fields_list_controls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListControlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_controls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListControls(ctx, input)
				}
				var results []*svc.ListControlsOutput
				p := svc.NewListControlsPaginator(client, input)
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
		"list-keywords-for-data-source": {
			Name:   "list-keywords-for-data-source",
			Fields: fields_list_keywords_for_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeywordsForDataSourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_keywords_for_data_source, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeywordsForDataSource(ctx, input)
				}
				var results []*svc.ListKeywordsForDataSourceOutput
				p := svc.NewListKeywordsForDataSourcePaginator(client, input)
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
		"list-notifications": {
			Name:   "list-notifications",
			Fields: fields_list_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotifications(ctx, input)
				}
				var results []*svc.ListNotificationsOutput
				p := svc.NewListNotificationsPaginator(client, input)
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
		"register-account": {
			Name:   "register-account",
			Fields: fields_register_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterAccount(ctx, input)
			},
		},
		"register-organization-admin-account": {
			Name:   "register-organization-admin-account",
			Fields: fields_register_organization_admin_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterOrganizationAdminAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_organization_admin_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterOrganizationAdminAccount(ctx, input)
			},
		},
		"start-assessment-framework-share": {
			Name:   "start-assessment-framework-share",
			Fields: fields_start_assessment_framework_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartAssessmentFrameworkShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_assessment_framework_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartAssessmentFrameworkShare(ctx, input)
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
		"update-assessment": {
			Name:   "update-assessment",
			Fields: fields_update_assessment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessment(ctx, input)
			},
		},
		"update-assessment-control": {
			Name:   "update-assessment-control",
			Fields: fields_update_assessment_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessmentControl(ctx, input)
			},
		},
		"update-assessment-control-set-status": {
			Name:   "update-assessment-control-set-status",
			Fields: fields_update_assessment_control_set_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentControlSetStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment_control_set_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessmentControlSetStatus(ctx, input)
			},
		},
		"update-assessment-framework": {
			Name:   "update-assessment-framework",
			Fields: fields_update_assessment_framework,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentFrameworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment_framework, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessmentFramework(ctx, input)
			},
		},
		"update-assessment-framework-share": {
			Name:   "update-assessment-framework-share",
			Fields: fields_update_assessment_framework_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentFrameworkShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment_framework_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessmentFrameworkShare(ctx, input)
			},
		},
		"update-assessment-status": {
			Name:   "update-assessment-status",
			Fields: fields_update_assessment_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssessmentStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_assessment_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssessmentStatus(ctx, input)
			},
		},
		"update-control": {
			Name:   "update-control",
			Fields: fields_update_control,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateControlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_control, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateControl(ctx, input)
			},
		},
		"update-settings": {
			Name:   "update-settings",
			Fields: fields_update_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSettings(ctx, input)
			},
		},
		"validate-assessment-report-integrity": {
			Name:   "validate-assessment-report-integrity",
			Fields: fields_validate_assessment_report_integrity,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ValidateAssessmentReportIntegrityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_validate_assessment_report_integrity, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ValidateAssessmentReportIntegrity(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("auditmanager", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
