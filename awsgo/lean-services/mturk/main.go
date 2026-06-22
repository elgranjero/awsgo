package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/mturk"
)

var fields_accept_qualification_request = []leanruntime.Field{
	{Name: "IntegerValue", Flag: "integer-value", Type: "*int32", Required: false},
	{Name: "QualificationRequestId", Flag: "qualification-request-id", Type: "*string", Required: true},
}

var fields_approve_assignment = []leanruntime.Field{
	{Name: "AssignmentId", Flag: "assignment-id", Type: "*string", Required: true},
	{Name: "OverrideRejection", Flag: "override-rejection", Type: "*bool", Required: false},
	{Name: "RequesterFeedback", Flag: "requester-feedback", Type: "*string", Required: false},
}

var fields_associate_qualification_with_worker = []leanruntime.Field{
	{Name: "IntegerValue", Flag: "integer-value", Type: "*int32", Required: false},
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
	{Name: "SendNotification", Flag: "send-notification", Type: "*bool", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_create_additional_assignments_for_hit = []leanruntime.Field{
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
	{Name: "NumberOfAdditionalAssignments", Flag: "number-of-additional-assignments", Type: "*int32", Required: true},
	{Name: "UniqueRequestToken", Flag: "unique-request-token", Type: "*string", Required: false},
}

var fields_create_hit = []leanruntime.Field{
	{Name: "AssignmentDurationInSeconds", Flag: "assignment-duration-in-seconds", Type: "*int64", Required: true},
	{Name: "AssignmentReviewPolicy", Flag: "assignment-review-policy", Type: "*types.ReviewPolicy", Required: false},
	{Name: "AutoApprovalDelayInSeconds", Flag: "auto-approval-delay-in-seconds", Type: "*int64", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "HITLayoutId", Flag: "hit-layout-id", Type: "*string", Required: false},
	{Name: "HITLayoutParameters", Flag: "hit-layout-parameters", Type: "[]types.HITLayoutParameter", Required: false},
	{Name: "HITReviewPolicy", Flag: "hit-review-policy", Type: "*types.ReviewPolicy", Required: false},
	{Name: "Keywords", Flag: "keywords", Type: "*string", Required: false},
	{Name: "LifetimeInSeconds", Flag: "lifetime-in-seconds", Type: "*int64", Required: true},
	{Name: "MaxAssignments", Flag: "max-assignments", Type: "*int32", Required: false},
	{Name: "QualificationRequirements", Flag: "qualification-requirements", Type: "[]types.QualificationRequirement", Required: false},
	{Name: "Question", Flag: "question", Type: "*string", Required: false},
	{Name: "RequesterAnnotation", Flag: "requester-annotation", Type: "*string", Required: false},
	{Name: "Reward", Flag: "reward", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
	{Name: "UniqueRequestToken", Flag: "unique-request-token", Type: "*string", Required: false},
}

var fields_create_hit_type = []leanruntime.Field{
	{Name: "AssignmentDurationInSeconds", Flag: "assignment-duration-in-seconds", Type: "*int64", Required: true},
	{Name: "AutoApprovalDelayInSeconds", Flag: "auto-approval-delay-in-seconds", Type: "*int64", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Keywords", Flag: "keywords", Type: "*string", Required: false},
	{Name: "QualificationRequirements", Flag: "qualification-requirements", Type: "[]types.QualificationRequirement", Required: false},
	{Name: "Reward", Flag: "reward", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_create_hit_with_hit_type = []leanruntime.Field{
	{Name: "AssignmentReviewPolicy", Flag: "assignment-review-policy", Type: "*types.ReviewPolicy", Required: false},
	{Name: "HITLayoutId", Flag: "hit-layout-id", Type: "*string", Required: false},
	{Name: "HITLayoutParameters", Flag: "hit-layout-parameters", Type: "[]types.HITLayoutParameter", Required: false},
	{Name: "HITReviewPolicy", Flag: "hit-review-policy", Type: "*types.ReviewPolicy", Required: false},
	{Name: "HITTypeId", Flag: "hit-type-id", Type: "*string", Required: true},
	{Name: "LifetimeInSeconds", Flag: "lifetime-in-seconds", Type: "*int64", Required: true},
	{Name: "MaxAssignments", Flag: "max-assignments", Type: "*int32", Required: false},
	{Name: "Question", Flag: "question", Type: "*string", Required: false},
	{Name: "RequesterAnnotation", Flag: "requester-annotation", Type: "*string", Required: false},
	{Name: "UniqueRequestToken", Flag: "unique-request-token", Type: "*string", Required: false},
}

var fields_create_qualification_type = []leanruntime.Field{
	{Name: "AnswerKey", Flag: "answer-key", Type: "*string", Required: false},
	{Name: "AutoGranted", Flag: "auto-granted", Type: "*bool", Required: false},
	{Name: "AutoGrantedValue", Flag: "auto-granted-value", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Keywords", Flag: "keywords", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "QualificationTypeStatus", Flag: "qualification-type-status", Type: "types.QualificationTypeStatus", Required: true},
	{Name: "RetryDelayInSeconds", Flag: "retry-delay-in-seconds", Type: "*int64", Required: false},
	{Name: "Test", Flag: "test", Type: "*string", Required: false},
	{Name: "TestDurationInSeconds", Flag: "test-duration-in-seconds", Type: "*int64", Required: false},
}

var fields_create_worker_block = []leanruntime.Field{
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_delete_hit = []leanruntime.Field{
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
}

var fields_delete_qualification_type = []leanruntime.Field{
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
}

var fields_delete_worker_block = []leanruntime.Field{
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_disassociate_qualification_from_worker = []leanruntime.Field{
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_get_account_balance = []leanruntime.Field{}

var fields_get_assignment = []leanruntime.Field{
	{Name: "AssignmentId", Flag: "assignment-id", Type: "*string", Required: true},
}

var fields_get_file_upload_url = []leanruntime.Field{
	{Name: "AssignmentId", Flag: "assignment-id", Type: "*string", Required: true},
	{Name: "QuestionIdentifier", Flag: "question-identifier", Type: "*string", Required: true},
}

var fields_get_hit = []leanruntime.Field{
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
}

var fields_get_qualification_score = []leanruntime.Field{
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_get_qualification_type = []leanruntime.Field{
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
}

var fields_list_assignments_for_hit = []leanruntime.Field{
	{Name: "AssignmentStatuses", Flag: "assignment-statuses", Type: "[]types.AssignmentStatus", Required: false},
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_bonus_payments = []leanruntime.Field{
	{Name: "AssignmentId", Flag: "assignment-id", Type: "*string", Required: false},
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hits = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_hits_for_qualification_type = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
}

var fields_list_qualification_requests = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: false},
}

var fields_list_qualification_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MustBeOwnedByCaller", Flag: "must-be-owned-by-caller", Type: "*bool", Required: false},
	{Name: "MustBeRequestable", Flag: "must-be-requestable", Type: "*bool", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Query", Flag: "query", Type: "*string", Required: false},
}

var fields_list_review_policy_results_for_hit = []leanruntime.Field{
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyLevels", Flag: "policy-levels", Type: "[]types.ReviewPolicyLevel", Required: false},
	{Name: "RetrieveActions", Flag: "retrieve-actions", Type: "*bool", Required: false},
	{Name: "RetrieveResults", Flag: "retrieve-results", Type: "*bool", Required: false},
}

var fields_list_reviewable_hits = []leanruntime.Field{
	{Name: "HITTypeId", Flag: "hit-type-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ReviewableHITStatus", Required: false},
}

var fields_list_worker_blocks = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_workers_with_qualification_type = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.QualificationStatus", Required: false},
}

var fields_notify_workers = []leanruntime.Field{
	{Name: "MessageText", Flag: "message-text", Type: "*string", Required: true},
	{Name: "Subject", Flag: "subject", Type: "*string", Required: true},
	{Name: "WorkerIds", Flag: "worker-ids", Type: "[]string", Required: true},
}

var fields_reject_assignment = []leanruntime.Field{
	{Name: "AssignmentId", Flag: "assignment-id", Type: "*string", Required: true},
	{Name: "RequesterFeedback", Flag: "requester-feedback", Type: "*string", Required: true},
}

var fields_reject_qualification_request = []leanruntime.Field{
	{Name: "QualificationRequestId", Flag: "qualification-request-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_send_bonus = []leanruntime.Field{
	{Name: "AssignmentId", Flag: "assignment-id", Type: "*string", Required: true},
	{Name: "BonusAmount", Flag: "bonus-amount", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
	{Name: "UniqueRequestToken", Flag: "unique-request-token", Type: "*string", Required: false},
	{Name: "WorkerId", Flag: "worker-id", Type: "*string", Required: true},
}

var fields_send_test_event_notification = []leanruntime.Field{
	{Name: "Notification", Flag: "notification", Type: "*types.NotificationSpecification", Required: true},
	{Name: "TestEventType", Flag: "test-event-type", Type: "types.EventType", Required: true},
}

var fields_update_expiration_for_hit = []leanruntime.Field{
	{Name: "ExpireAt", Flag: "expire-at", Type: "*time.Time", Required: true},
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
}

var fields_update_hit_review_status = []leanruntime.Field{
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
	{Name: "Revert", Flag: "revert", Type: "*bool", Required: false},
}

var fields_update_hit_type_of_hit = []leanruntime.Field{
	{Name: "HITId", Flag: "hitid", Type: "*string", Required: true},
	{Name: "HITTypeId", Flag: "hit-type-id", Type: "*string", Required: true},
}

var fields_update_notification_settings = []leanruntime.Field{
	{Name: "Active", Flag: "active", Type: "*bool", Required: false},
	{Name: "HITTypeId", Flag: "hit-type-id", Type: "*string", Required: true},
	{Name: "Notification", Flag: "notification", Type: "*types.NotificationSpecification", Required: false},
}

var fields_update_qualification_type = []leanruntime.Field{
	{Name: "AnswerKey", Flag: "answer-key", Type: "*string", Required: false},
	{Name: "AutoGranted", Flag: "auto-granted", Type: "*bool", Required: false},
	{Name: "AutoGrantedValue", Flag: "auto-granted-value", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "QualificationTypeId", Flag: "qualification-type-id", Type: "*string", Required: true},
	{Name: "QualificationTypeStatus", Flag: "qualification-type-status", Type: "types.QualificationTypeStatus", Required: false},
	{Name: "RetryDelayInSeconds", Flag: "retry-delay-in-seconds", Type: "*int64", Required: false},
	{Name: "Test", Flag: "test", Type: "*string", Required: false},
	{Name: "TestDurationInSeconds", Flag: "test-duration-in-seconds", Type: "*int64", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-qualification-request": {
			Name:   "accept-qualification-request",
			Fields: fields_accept_qualification_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptQualificationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_qualification_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptQualificationRequest(ctx, input)
			},
		},
		"approve-assignment": {
			Name:   "approve-assignment",
			Fields: fields_approve_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ApproveAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_approve_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ApproveAssignment(ctx, input)
			},
		},
		"associate-qualification-with-worker": {
			Name:   "associate-qualification-with-worker",
			Fields: fields_associate_qualification_with_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateQualificationWithWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_qualification_with_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateQualificationWithWorker(ctx, input)
			},
		},
		"create-additional-assignments-for-hit": {
			Name:   "create-additional-assignments-for-hit",
			Fields: fields_create_additional_assignments_for_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAdditionalAssignmentsForHITInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_additional_assignments_for_hit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAdditionalAssignmentsForHIT(ctx, input)
			},
		},
		"create-hit": {
			Name:   "create-hit",
			Fields: fields_create_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHITInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHIT(ctx, input)
			},
		},
		"create-hit-type": {
			Name:   "create-hit-type",
			Fields: fields_create_hit_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHITTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hit_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHITType(ctx, input)
			},
		},
		"create-hit-with-hit-type": {
			Name:   "create-hit-with-hit-type",
			Fields: fields_create_hit_with_hit_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateHITWithHITTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_hit_with_hit_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateHITWithHITType(ctx, input)
			},
		},
		"create-qualification-type": {
			Name:   "create-qualification-type",
			Fields: fields_create_qualification_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateQualificationTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_qualification_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateQualificationType(ctx, input)
			},
		},
		"create-worker-block": {
			Name:   "create-worker-block",
			Fields: fields_create_worker_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWorkerBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_worker_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWorkerBlock(ctx, input)
			},
		},
		"delete-hit": {
			Name:   "delete-hit",
			Fields: fields_delete_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteHITInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_hit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteHIT(ctx, input)
			},
		},
		"delete-qualification-type": {
			Name:   "delete-qualification-type",
			Fields: fields_delete_qualification_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteQualificationTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_qualification_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteQualificationType(ctx, input)
			},
		},
		"delete-worker-block": {
			Name:   "delete-worker-block",
			Fields: fields_delete_worker_block,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWorkerBlockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_worker_block, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWorkerBlock(ctx, input)
			},
		},
		"disassociate-qualification-from-worker": {
			Name:   "disassociate-qualification-from-worker",
			Fields: fields_disassociate_qualification_from_worker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateQualificationFromWorkerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_qualification_from_worker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateQualificationFromWorker(ctx, input)
			},
		},
		"get-account-balance": {
			Name:   "get-account-balance",
			Fields: fields_get_account_balance,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountBalanceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_balance, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountBalance(ctx, input)
			},
		},
		"get-assignment": {
			Name:   "get-assignment",
			Fields: fields_get_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssignment(ctx, input)
			},
		},
		"get-file-upload-url": {
			Name:   "get-file-upload-url",
			Fields: fields_get_file_upload_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFileUploadURLInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_file_upload_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFileUploadURL(ctx, input)
			},
		},
		"get-hit": {
			Name:   "get-hit",
			Fields: fields_get_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetHITInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_hit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetHIT(ctx, input)
			},
		},
		"get-qualification-score": {
			Name:   "get-qualification-score",
			Fields: fields_get_qualification_score,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQualificationScoreInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_qualification_score, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQualificationScore(ctx, input)
			},
		},
		"get-qualification-type": {
			Name:   "get-qualification-type",
			Fields: fields_get_qualification_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetQualificationTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_qualification_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetQualificationType(ctx, input)
			},
		},
		"list-assignments-for-hit": {
			Name:   "list-assignments-for-hit",
			Fields: fields_list_assignments_for_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssignmentsForHITInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_assignments_for_hit, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssignmentsForHIT(ctx, input)
				}
				var results []*svc.ListAssignmentsForHITOutput
				p := svc.NewListAssignmentsForHITPaginator(client, input)
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
		"list-bonus-payments": {
			Name:   "list-bonus-payments",
			Fields: fields_list_bonus_payments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBonusPaymentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_bonus_payments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBonusPayments(ctx, input)
				}
				var results []*svc.ListBonusPaymentsOutput
				p := svc.NewListBonusPaymentsPaginator(client, input)
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
		"list-hits": {
			Name:   "list-hits",
			Fields: fields_list_hits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHITsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHITs(ctx, input)
				}
				var results []*svc.ListHITsOutput
				p := svc.NewListHITsPaginator(client, input)
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
		"list-hits-for-qualification-type": {
			Name:   "list-hits-for-qualification-type",
			Fields: fields_list_hits_for_qualification_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHITsForQualificationTypeInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_hits_for_qualification_type, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHITsForQualificationType(ctx, input)
				}
				var results []*svc.ListHITsForQualificationTypeOutput
				p := svc.NewListHITsForQualificationTypePaginator(client, input)
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
		"list-qualification-requests": {
			Name:   "list-qualification-requests",
			Fields: fields_list_qualification_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQualificationRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_qualification_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQualificationRequests(ctx, input)
				}
				var results []*svc.ListQualificationRequestsOutput
				p := svc.NewListQualificationRequestsPaginator(client, input)
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
		"list-qualification-types": {
			Name:   "list-qualification-types",
			Fields: fields_list_qualification_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListQualificationTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_qualification_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListQualificationTypes(ctx, input)
				}
				var results []*svc.ListQualificationTypesOutput
				p := svc.NewListQualificationTypesPaginator(client, input)
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
		"list-review-policy-results-for-hit": {
			Name:   "list-review-policy-results-for-hit",
			Fields: fields_list_review_policy_results_for_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReviewPolicyResultsForHITInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_review_policy_results_for_hit, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReviewPolicyResultsForHIT(ctx, input)
				}
				var results []*svc.ListReviewPolicyResultsForHITOutput
				p := svc.NewListReviewPolicyResultsForHITPaginator(client, input)
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
		"list-reviewable-hits": {
			Name:   "list-reviewable-hits",
			Fields: fields_list_reviewable_hits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReviewableHITsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_reviewable_hits, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListReviewableHITs(ctx, input)
				}
				var results []*svc.ListReviewableHITsOutput
				p := svc.NewListReviewableHITsPaginator(client, input)
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
		"list-worker-blocks": {
			Name:   "list-worker-blocks",
			Fields: fields_list_worker_blocks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkerBlocksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_worker_blocks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkerBlocks(ctx, input)
				}
				var results []*svc.ListWorkerBlocksOutput
				p := svc.NewListWorkerBlocksPaginator(client, input)
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
		"list-workers-with-qualification-type": {
			Name:   "list-workers-with-qualification-type",
			Fields: fields_list_workers_with_qualification_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWorkersWithQualificationTypeInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_workers_with_qualification_type, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWorkersWithQualificationType(ctx, input)
				}
				var results []*svc.ListWorkersWithQualificationTypeOutput
				p := svc.NewListWorkersWithQualificationTypePaginator(client, input)
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
		"notify-workers": {
			Name:   "notify-workers",
			Fields: fields_notify_workers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyWorkersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_workers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyWorkers(ctx, input)
			},
		},
		"reject-assignment": {
			Name:   "reject-assignment",
			Fields: fields_reject_assignment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectAssignmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_assignment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectAssignment(ctx, input)
			},
		},
		"reject-qualification-request": {
			Name:   "reject-qualification-request",
			Fields: fields_reject_qualification_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectQualificationRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_qualification_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectQualificationRequest(ctx, input)
			},
		},
		"send-bonus": {
			Name:   "send-bonus",
			Fields: fields_send_bonus,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendBonusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_bonus, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendBonus(ctx, input)
			},
		},
		"send-test-event-notification": {
			Name:   "send-test-event-notification",
			Fields: fields_send_test_event_notification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendTestEventNotificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_test_event_notification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendTestEventNotification(ctx, input)
			},
		},
		"update-expiration-for-hit": {
			Name:   "update-expiration-for-hit",
			Fields: fields_update_expiration_for_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateExpirationForHITInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_expiration_for_hit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateExpirationForHIT(ctx, input)
			},
		},
		"update-hit-review-status": {
			Name:   "update-hit-review-status",
			Fields: fields_update_hit_review_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHITReviewStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hit_review_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHITReviewStatus(ctx, input)
			},
		},
		"update-hit-type-of-hit": {
			Name:   "update-hit-type-of-hit",
			Fields: fields_update_hit_type_of_hit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateHITTypeOfHITInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_hit_type_of_hit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateHITTypeOfHIT(ctx, input)
			},
		},
		"update-notification-settings": {
			Name:   "update-notification-settings",
			Fields: fields_update_notification_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNotificationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_notification_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNotificationSettings(ctx, input)
			},
		},
		"update-qualification-type": {
			Name:   "update-qualification-type",
			Fields: fields_update_qualification_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateQualificationTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_qualification_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateQualificationType(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("mturk", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
