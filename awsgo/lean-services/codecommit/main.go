package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/codecommit"
)

var fields_associate_approval_rule_template_with_repository = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_batch_associate_approval_rule_template_with_repositories = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: true},
}

var fields_batch_describe_merge_conflicts = []leanruntime.Field{
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "FilePaths", Flag: "file-paths", Type: "[]string", Required: false},
	{Name: "MaxConflictFiles", Flag: "max-conflict-files", Type: "*int32", Required: false},
	{Name: "MaxMergeHunks", Flag: "max-merge-hunks", Type: "*int32", Required: false},
	{Name: "MergeOption", Flag: "merge-option", Type: "types.MergeOptionTypeEnum", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
}

var fields_batch_disassociate_approval_rule_template_from_repositories = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: true},
}

var fields_batch_get_commits = []leanruntime.Field{
	{Name: "CommitIds", Flag: "commit-ids", Type: "[]string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_batch_get_repositories = []leanruntime.Field{
	{Name: "RepositoryNames", Flag: "repository-names", Type: "[]string", Required: true},
}

var fields_create_approval_rule_template = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateContent", Flag: "approval-rule-template-content", Type: "*string", Required: true},
	{Name: "ApprovalRuleTemplateDescription", Flag: "approval-rule-template-description", Type: "*string", Required: false},
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
}

var fields_create_branch = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "CommitId", Flag: "commit-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_create_commit = []leanruntime.Field{
	{Name: "AuthorName", Flag: "author-name", Type: "*string", Required: false},
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "DeleteFiles", Flag: "delete-files", Type: "[]types.DeleteFileEntry", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "ParentCommitId", Flag: "parent-commit-id", Type: "*string", Required: false},
	{Name: "PutFiles", Flag: "put-files", Type: "[]types.PutFileEntry", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SetFileModes", Flag: "set-file-modes", Type: "[]types.SetFileModeEntry", Required: false},
}

var fields_create_pull_request = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Targets", Flag: "targets", Type: "[]types.Target", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_create_pull_request_approval_rule = []leanruntime.Field{
	{Name: "ApprovalRuleContent", Flag: "approval-rule-content", Type: "*string", Required: true},
	{Name: "ApprovalRuleName", Flag: "approval-rule-name", Type: "*string", Required: true},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
}

var fields_create_repository = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "RepositoryDescription", Flag: "repository-description", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_unreferenced_merge_commit = []leanruntime.Field{
	{Name: "AuthorName", Flag: "author-name", Type: "*string", Required: false},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolution", Flag: "conflict-resolution", Type: "*types.ConflictResolution", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "MergeOption", Flag: "merge-option", Type: "types.MergeOptionTypeEnum", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
}

var fields_delete_approval_rule_template = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
}

var fields_delete_branch = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_delete_comment_content = []leanruntime.Field{
	{Name: "CommentId", Flag: "comment-id", Type: "*string", Required: true},
}

var fields_delete_file = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentCommitId", Flag: "parent-commit-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_delete_pull_request_approval_rule = []leanruntime.Field{
	{Name: "ApprovalRuleName", Flag: "approval-rule-name", Type: "*string", Required: true},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
}

var fields_delete_repository = []leanruntime.Field{
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_describe_merge_conflicts = []leanruntime.Field{
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "MaxMergeHunks", Flag: "max-merge-hunks", Type: "*int32", Required: false},
	{Name: "MergeOption", Flag: "merge-option", Type: "types.MergeOptionTypeEnum", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
}

var fields_describe_pull_request_events = []leanruntime.Field{
	{Name: "ActorArn", Flag: "actor-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PullRequestEventType", Flag: "pull-request-event-type", Type: "types.PullRequestEventType", Required: false},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
}

var fields_disassociate_approval_rule_template_from_repository = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_evaluate_pull_request_approval_rules = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_get_approval_rule_template = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
}

var fields_get_blob = []leanruntime.Field{
	{Name: "BlobId", Flag: "blob-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_branch = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: false},
}

var fields_get_comment = []leanruntime.Field{
	{Name: "CommentId", Flag: "comment-id", Type: "*string", Required: true},
}

var fields_get_comment_reactions = []leanruntime.Field{
	{Name: "CommentId", Flag: "comment-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ReactionUserArn", Flag: "reaction-user-arn", Type: "*string", Required: false},
}

var fields_get_comments_for_compared_commit = []leanruntime.Field{
	{Name: "AfterCommitId", Flag: "after-commit-id", Type: "*string", Required: true},
	{Name: "BeforeCommitId", Flag: "before-commit-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_comments_for_pull_request = []leanruntime.Field{
	{Name: "AfterCommitId", Flag: "after-commit-id", Type: "*string", Required: false},
	{Name: "BeforeCommitId", Flag: "before-commit-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: false},
}

var fields_get_commit = []leanruntime.Field{
	{Name: "CommitId", Flag: "commit-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_differences = []leanruntime.Field{
	{Name: "AfterCommitSpecifier", Flag: "after-commit-specifier", Type: "*string", Required: true},
	{Name: "AfterPath", Flag: "after-path", Type: "*string", Required: false},
	{Name: "BeforeCommitSpecifier", Flag: "before-commit-specifier", Type: "*string", Required: false},
	{Name: "BeforePath", Flag: "before-path", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_file = []leanruntime.Field{
	{Name: "CommitSpecifier", Flag: "commit-specifier", Type: "*string", Required: false},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_folder = []leanruntime.Field{
	{Name: "CommitSpecifier", Flag: "commit-specifier", Type: "*string", Required: false},
	{Name: "FolderPath", Flag: "folder-path", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_merge_commit = []leanruntime.Field{
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
}

var fields_get_merge_conflicts = []leanruntime.Field{
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "MaxConflictFiles", Flag: "max-conflict-files", Type: "*int32", Required: false},
	{Name: "MergeOption", Flag: "merge-option", Type: "types.MergeOptionTypeEnum", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
}

var fields_get_merge_options = []leanruntime.Field{
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
}

var fields_get_pull_request = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
}

var fields_get_pull_request_approval_states = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_get_pull_request_override_state = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_get_repository = []leanruntime.Field{
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_get_repository_triggers = []leanruntime.Field{
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_approval_rule_templates = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associated_approval_rule_templates_for_repository = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_branches = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_file_commit_history = []leanruntime.Field{
	{Name: "CommitSpecifier", Flag: "commit-specifier", Type: "*string", Required: false},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_pull_requests = []leanruntime.Field{
	{Name: "AuthorArn", Flag: "author-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PullRequestStatus", Flag: "pull-request-status", Type: "types.PullRequestStatusEnum", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_list_repositories = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Order", Flag: "order", Type: "types.OrderEnum", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortByEnum", Required: false},
}

var fields_list_repositories_for_approval_rule_template = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_merge_branches_by_fast_forward = []leanruntime.Field{
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
	{Name: "TargetBranch", Flag: "target-branch", Type: "*string", Required: false},
}

var fields_merge_branches_by_squash = []leanruntime.Field{
	{Name: "AuthorName", Flag: "author-name", Type: "*string", Required: false},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolution", Flag: "conflict-resolution", Type: "*types.ConflictResolution", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
	{Name: "TargetBranch", Flag: "target-branch", Type: "*string", Required: false},
}

var fields_merge_branches_by_three_way = []leanruntime.Field{
	{Name: "AuthorName", Flag: "author-name", Type: "*string", Required: false},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolution", Flag: "conflict-resolution", Type: "*types.ConflictResolution", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "DestinationCommitSpecifier", Flag: "destination-commit-specifier", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitSpecifier", Flag: "source-commit-specifier", Type: "*string", Required: true},
	{Name: "TargetBranch", Flag: "target-branch", Type: "*string", Required: false},
}

var fields_merge_pull_request_by_fast_forward = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitId", Flag: "source-commit-id", Type: "*string", Required: false},
}

var fields_merge_pull_request_by_squash = []leanruntime.Field{
	{Name: "AuthorName", Flag: "author-name", Type: "*string", Required: false},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolution", Flag: "conflict-resolution", Type: "*types.ConflictResolution", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitId", Flag: "source-commit-id", Type: "*string", Required: false},
}

var fields_merge_pull_request_by_three_way = []leanruntime.Field{
	{Name: "AuthorName", Flag: "author-name", Type: "*string", Required: false},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "ConflictDetailLevel", Flag: "conflict-detail-level", Type: "types.ConflictDetailLevelTypeEnum", Required: false},
	{Name: "ConflictResolution", Flag: "conflict-resolution", Type: "*types.ConflictResolution", Required: false},
	{Name: "ConflictResolutionStrategy", Flag: "conflict-resolution-strategy", Type: "types.ConflictResolutionStrategyTypeEnum", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "KeepEmptyFolders", Flag: "keep-empty-folders", Type: "bool", Required: false},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "SourceCommitId", Flag: "source-commit-id", Type: "*string", Required: false},
}

var fields_override_pull_request_approval_rules = []leanruntime.Field{
	{Name: "OverrideStatus", Flag: "override-status", Type: "types.OverrideStatus", Required: true},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_post_comment_for_compared_commit = []leanruntime.Field{
	{Name: "AfterCommitId", Flag: "after-commit-id", Type: "*string", Required: true},
	{Name: "BeforeCommitId", Flag: "before-commit-id", Type: "*string", Required: false},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_post_comment_for_pull_request = []leanruntime.Field{
	{Name: "AfterCommitId", Flag: "after-commit-id", Type: "*string", Required: true},
	{Name: "BeforeCommitId", Flag: "before-commit-id", Type: "*string", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Location", Flag: "location", Type: "*types.Location", Required: false},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_post_comment_reply = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: false},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "InReplyTo", Flag: "in-reply-to", Type: "*string", Required: true},
}

var fields_put_comment_reaction = []leanruntime.Field{
	{Name: "CommentId", Flag: "comment-id", Type: "*string", Required: true},
	{Name: "ReactionValue", Flag: "reaction-value", Type: "*string", Required: true},
}

var fields_put_file = []leanruntime.Field{
	{Name: "BranchName", Flag: "branch-name", Type: "*string", Required: true},
	{Name: "CommitMessage", Flag: "commit-message", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "FileContent", Flag: "file-content", Type: "[]byte", Required: true},
	{Name: "FileMode", Flag: "file-mode", Type: "types.FileModeTypeEnum", Required: false},
	{Name: "FilePath", Flag: "file-path", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParentCommitId", Flag: "parent-commit-id", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_put_repository_triggers = []leanruntime.Field{
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "Triggers", Flag: "triggers", Type: "[]types.RepositoryTrigger", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_test_repository_triggers = []leanruntime.Field{
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
	{Name: "Triggers", Flag: "triggers", Type: "[]types.RepositoryTrigger", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_approval_rule_template_content = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
	{Name: "ExistingRuleContentSha256", Flag: "existing-rule-content-sha256", Type: "*string", Required: false},
	{Name: "NewRuleContent", Flag: "new-rule-content", Type: "*string", Required: true},
}

var fields_update_approval_rule_template_description = []leanruntime.Field{
	{Name: "ApprovalRuleTemplateDescription", Flag: "approval-rule-template-description", Type: "*string", Required: true},
	{Name: "ApprovalRuleTemplateName", Flag: "approval-rule-template-name", Type: "*string", Required: true},
}

var fields_update_approval_rule_template_name = []leanruntime.Field{
	{Name: "NewApprovalRuleTemplateName", Flag: "new-approval-rule-template-name", Type: "*string", Required: true},
	{Name: "OldApprovalRuleTemplateName", Flag: "old-approval-rule-template-name", Type: "*string", Required: true},
}

var fields_update_comment = []leanruntime.Field{
	{Name: "CommentId", Flag: "comment-id", Type: "*string", Required: true},
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
}

var fields_update_default_branch = []leanruntime.Field{
	{Name: "DefaultBranchName", Flag: "default-branch-name", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_update_pull_request_approval_rule_content = []leanruntime.Field{
	{Name: "ApprovalRuleName", Flag: "approval-rule-name", Type: "*string", Required: true},
	{Name: "ExistingRuleContentSha256", Flag: "existing-rule-content-sha256", Type: "*string", Required: false},
	{Name: "NewRuleContent", Flag: "new-rule-content", Type: "*string", Required: true},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
}

var fields_update_pull_request_approval_state = []leanruntime.Field{
	{Name: "ApprovalState", Flag: "approval-state", Type: "types.ApprovalState", Required: true},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
}

var fields_update_pull_request_description = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
}

var fields_update_pull_request_status = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "PullRequestStatus", Flag: "pull-request-status", Type: "types.PullRequestStatusEnum", Required: true},
}

var fields_update_pull_request_title = []leanruntime.Field{
	{Name: "PullRequestId", Flag: "pull-request-id", Type: "*string", Required: true},
	{Name: "Title", Flag: "title", Type: "*string", Required: true},
}

var fields_update_repository_description = []leanruntime.Field{
	{Name: "RepositoryDescription", Flag: "repository-description", Type: "*string", Required: false},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_update_repository_encryption_key = []leanruntime.Field{
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: true},
	{Name: "RepositoryName", Flag: "repository-name", Type: "*string", Required: true},
}

var fields_update_repository_name = []leanruntime.Field{
	{Name: "NewName", Flag: "new-name", Type: "*string", Required: true},
	{Name: "OldName", Flag: "old-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-approval-rule-template-with-repository": {
			Name:   "associate-approval-rule-template-with-repository",
			Fields: fields_associate_approval_rule_template_with_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateApprovalRuleTemplateWithRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_approval_rule_template_with_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateApprovalRuleTemplateWithRepository(ctx, input)
			},
		},
		"batch-associate-approval-rule-template-with-repositories": {
			Name:   "batch-associate-approval-rule-template-with-repositories",
			Fields: fields_batch_associate_approval_rule_template_with_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateApprovalRuleTemplateWithRepositoriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_approval_rule_template_with_repositories, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateApprovalRuleTemplateWithRepositories(ctx, input)
			},
		},
		"batch-describe-merge-conflicts": {
			Name:   "batch-describe-merge-conflicts",
			Fields: fields_batch_describe_merge_conflicts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDescribeMergeConflictsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_describe_merge_conflicts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDescribeMergeConflicts(ctx, input)
			},
		},
		"batch-disassociate-approval-rule-template-from-repositories": {
			Name:   "batch-disassociate-approval-rule-template-from-repositories",
			Fields: fields_batch_disassociate_approval_rule_template_from_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_approval_rule_template_from_repositories, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateApprovalRuleTemplateFromRepositories(ctx, input)
			},
		},
		"batch-get-commits": {
			Name:   "batch-get-commits",
			Fields: fields_batch_get_commits,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetCommitsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_commits, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetCommits(ctx, input)
			},
		},
		"batch-get-repositories": {
			Name:   "batch-get-repositories",
			Fields: fields_batch_get_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetRepositoriesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_repositories, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetRepositories(ctx, input)
			},
		},
		"create-approval-rule-template": {
			Name:   "create-approval-rule-template",
			Fields: fields_create_approval_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateApprovalRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_approval_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateApprovalRuleTemplate(ctx, input)
			},
		},
		"create-branch": {
			Name:   "create-branch",
			Fields: fields_create_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBranch(ctx, input)
			},
		},
		"create-commit": {
			Name:   "create-commit",
			Fields: fields_create_commit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCommitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_commit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCommit(ctx, input)
			},
		},
		"create-pull-request": {
			Name:   "create-pull-request",
			Fields: fields_create_pull_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePullRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pull_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePullRequest(ctx, input)
			},
		},
		"create-pull-request-approval-rule": {
			Name:   "create-pull-request-approval-rule",
			Fields: fields_create_pull_request_approval_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePullRequestApprovalRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_pull_request_approval_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePullRequestApprovalRule(ctx, input)
			},
		},
		"create-repository": {
			Name:   "create-repository",
			Fields: fields_create_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRepository(ctx, input)
			},
		},
		"create-unreferenced-merge-commit": {
			Name:   "create-unreferenced-merge-commit",
			Fields: fields_create_unreferenced_merge_commit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUnreferencedMergeCommitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_unreferenced_merge_commit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUnreferencedMergeCommit(ctx, input)
			},
		},
		"delete-approval-rule-template": {
			Name:   "delete-approval-rule-template",
			Fields: fields_delete_approval_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteApprovalRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_approval_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteApprovalRuleTemplate(ctx, input)
			},
		},
		"delete-branch": {
			Name:   "delete-branch",
			Fields: fields_delete_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBranch(ctx, input)
			},
		},
		"delete-comment-content": {
			Name:   "delete-comment-content",
			Fields: fields_delete_comment_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCommentContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_comment_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCommentContent(ctx, input)
			},
		},
		"delete-file": {
			Name:   "delete-file",
			Fields: fields_delete_file,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_file, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFile(ctx, input)
			},
		},
		"delete-pull-request-approval-rule": {
			Name:   "delete-pull-request-approval-rule",
			Fields: fields_delete_pull_request_approval_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePullRequestApprovalRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_pull_request_approval_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePullRequestApprovalRule(ctx, input)
			},
		},
		"delete-repository": {
			Name:   "delete-repository",
			Fields: fields_delete_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRepository(ctx, input)
			},
		},
		"describe-merge-conflicts": {
			Name:   "describe-merge-conflicts",
			Fields: fields_describe_merge_conflicts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMergeConflictsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_merge_conflicts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribeMergeConflicts(ctx, input)
				}
				var results []*svc.DescribeMergeConflictsOutput
				p := svc.NewDescribeMergeConflictsPaginator(client, input)
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
		"describe-pull-request-events": {
			Name:   "describe-pull-request-events",
			Fields: fields_describe_pull_request_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePullRequestEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_pull_request_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePullRequestEvents(ctx, input)
				}
				var results []*svc.DescribePullRequestEventsOutput
				p := svc.NewDescribePullRequestEventsPaginator(client, input)
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
		"disassociate-approval-rule-template-from-repository": {
			Name:   "disassociate-approval-rule-template-from-repository",
			Fields: fields_disassociate_approval_rule_template_from_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateApprovalRuleTemplateFromRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_approval_rule_template_from_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateApprovalRuleTemplateFromRepository(ctx, input)
			},
		},
		"evaluate-pull-request-approval-rules": {
			Name:   "evaluate-pull-request-approval-rules",
			Fields: fields_evaluate_pull_request_approval_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluatePullRequestApprovalRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evaluate_pull_request_approval_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EvaluatePullRequestApprovalRules(ctx, input)
			},
		},
		"get-approval-rule-template": {
			Name:   "get-approval-rule-template",
			Fields: fields_get_approval_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetApprovalRuleTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_approval_rule_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetApprovalRuleTemplate(ctx, input)
			},
		},
		"get-blob": {
			Name:   "get-blob",
			Fields: fields_get_blob,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBlobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_blob, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBlob(ctx, input)
			},
		},
		"get-branch": {
			Name:   "get-branch",
			Fields: fields_get_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBranch(ctx, input)
			},
		},
		"get-comment": {
			Name:   "get-comment",
			Fields: fields_get_comment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_comment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComment(ctx, input)
			},
		},
		"get-comment-reactions": {
			Name:   "get-comment-reactions",
			Fields: fields_get_comment_reactions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommentReactionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_comment_reactions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCommentReactions(ctx, input)
				}
				var results []*svc.GetCommentReactionsOutput
				p := svc.NewGetCommentReactionsPaginator(client, input)
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
		"get-comments-for-compared-commit": {
			Name:   "get-comments-for-compared-commit",
			Fields: fields_get_comments_for_compared_commit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommentsForComparedCommitInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_comments_for_compared_commit, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCommentsForComparedCommit(ctx, input)
				}
				var results []*svc.GetCommentsForComparedCommitOutput
				p := svc.NewGetCommentsForComparedCommitPaginator(client, input)
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
		"get-comments-for-pull-request": {
			Name:   "get-comments-for-pull-request",
			Fields: fields_get_comments_for_pull_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommentsForPullRequestInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_comments_for_pull_request, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetCommentsForPullRequest(ctx, input)
				}
				var results []*svc.GetCommentsForPullRequestOutput
				p := svc.NewGetCommentsForPullRequestPaginator(client, input)
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
		"get-commit": {
			Name:   "get-commit",
			Fields: fields_get_commit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCommitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_commit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCommit(ctx, input)
			},
		},
		"get-differences": {
			Name:   "get-differences",
			Fields: fields_get_differences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDifferencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_differences, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDifferences(ctx, input)
				}
				var results []*svc.GetDifferencesOutput
				p := svc.NewGetDifferencesPaginator(client, input)
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
		"get-file": {
			Name:   "get-file",
			Fields: fields_get_file,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_file, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFile(ctx, input)
			},
		},
		"get-folder": {
			Name:   "get-folder",
			Fields: fields_get_folder,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFolderInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_folder, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFolder(ctx, input)
			},
		},
		"get-merge-commit": {
			Name:   "get-merge-commit",
			Fields: fields_get_merge_commit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMergeCommitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_merge_commit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMergeCommit(ctx, input)
			},
		},
		"get-merge-conflicts": {
			Name:   "get-merge-conflicts",
			Fields: fields_get_merge_conflicts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMergeConflictsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_merge_conflicts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMergeConflicts(ctx, input)
				}
				var results []*svc.GetMergeConflictsOutput
				p := svc.NewGetMergeConflictsPaginator(client, input)
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
		"get-merge-options": {
			Name:   "get-merge-options",
			Fields: fields_get_merge_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMergeOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_merge_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMergeOptions(ctx, input)
			},
		},
		"get-pull-request": {
			Name:   "get-pull-request",
			Fields: fields_get_pull_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPullRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pull_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPullRequest(ctx, input)
			},
		},
		"get-pull-request-approval-states": {
			Name:   "get-pull-request-approval-states",
			Fields: fields_get_pull_request_approval_states,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPullRequestApprovalStatesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pull_request_approval_states, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPullRequestApprovalStates(ctx, input)
			},
		},
		"get-pull-request-override-state": {
			Name:   "get-pull-request-override-state",
			Fields: fields_get_pull_request_override_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPullRequestOverrideStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_pull_request_override_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPullRequestOverrideState(ctx, input)
			},
		},
		"get-repository": {
			Name:   "get-repository",
			Fields: fields_get_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepository(ctx, input)
			},
		},
		"get-repository-triggers": {
			Name:   "get-repository-triggers",
			Fields: fields_get_repository_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRepositoryTriggersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_repository_triggers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRepositoryTriggers(ctx, input)
			},
		},
		"list-approval-rule-templates": {
			Name:   "list-approval-rule-templates",
			Fields: fields_list_approval_rule_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListApprovalRuleTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_approval_rule_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListApprovalRuleTemplates(ctx, input)
				}
				var results []*svc.ListApprovalRuleTemplatesOutput
				p := svc.NewListApprovalRuleTemplatesPaginator(client, input)
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
		"list-associated-approval-rule-templates-for-repository": {
			Name:   "list-associated-approval-rule-templates-for-repository",
			Fields: fields_list_associated_approval_rule_templates_for_repository,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociatedApprovalRuleTemplatesForRepositoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_associated_approval_rule_templates_for_repository, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssociatedApprovalRuleTemplatesForRepository(ctx, input)
				}
				var results []*svc.ListAssociatedApprovalRuleTemplatesForRepositoryOutput
				p := svc.NewListAssociatedApprovalRuleTemplatesForRepositoryPaginator(client, input)
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
		"list-branches": {
			Name:   "list-branches",
			Fields: fields_list_branches,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBranchesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_branches, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBranches(ctx, input)
				}
				var results []*svc.ListBranchesOutput
				p := svc.NewListBranchesPaginator(client, input)
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
		"list-file-commit-history": {
			Name:   "list-file-commit-history",
			Fields: fields_list_file_commit_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFileCommitHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_file_commit_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFileCommitHistory(ctx, input)
				}
				var results []*svc.ListFileCommitHistoryOutput
				p := svc.NewListFileCommitHistoryPaginator(client, input)
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
		"list-pull-requests": {
			Name:   "list-pull-requests",
			Fields: fields_list_pull_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPullRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_pull_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPullRequests(ctx, input)
				}
				var results []*svc.ListPullRequestsOutput
				p := svc.NewListPullRequestsPaginator(client, input)
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
		"list-repositories": {
			Name:   "list-repositories",
			Fields: fields_list_repositories,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoriesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repositories, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositories(ctx, input)
				}
				var results []*svc.ListRepositoriesOutput
				p := svc.NewListRepositoriesPaginator(client, input)
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
		"list-repositories-for-approval-rule-template": {
			Name:   "list-repositories-for-approval-rule-template",
			Fields: fields_list_repositories_for_approval_rule_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRepositoriesForApprovalRuleTemplateInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_repositories_for_approval_rule_template, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRepositoriesForApprovalRuleTemplate(ctx, input)
				}
				var results []*svc.ListRepositoriesForApprovalRuleTemplateOutput
				p := svc.NewListRepositoriesForApprovalRuleTemplatePaginator(client, input)
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
		"merge-branches-by-fast-forward": {
			Name:   "merge-branches-by-fast-forward",
			Fields: fields_merge_branches_by_fast_forward,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergeBranchesByFastForwardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_branches_by_fast_forward, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergeBranchesByFastForward(ctx, input)
			},
		},
		"merge-branches-by-squash": {
			Name:   "merge-branches-by-squash",
			Fields: fields_merge_branches_by_squash,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergeBranchesBySquashInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_branches_by_squash, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergeBranchesBySquash(ctx, input)
			},
		},
		"merge-branches-by-three-way": {
			Name:   "merge-branches-by-three-way",
			Fields: fields_merge_branches_by_three_way,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergeBranchesByThreeWayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_branches_by_three_way, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergeBranchesByThreeWay(ctx, input)
			},
		},
		"merge-pull-request-by-fast-forward": {
			Name:   "merge-pull-request-by-fast-forward",
			Fields: fields_merge_pull_request_by_fast_forward,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergePullRequestByFastForwardInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_pull_request_by_fast_forward, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergePullRequestByFastForward(ctx, input)
			},
		},
		"merge-pull-request-by-squash": {
			Name:   "merge-pull-request-by-squash",
			Fields: fields_merge_pull_request_by_squash,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergePullRequestBySquashInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_pull_request_by_squash, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergePullRequestBySquash(ctx, input)
			},
		},
		"merge-pull-request-by-three-way": {
			Name:   "merge-pull-request-by-three-way",
			Fields: fields_merge_pull_request_by_three_way,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MergePullRequestByThreeWayInput{}
				if _, err := leanruntime.ApplyInput(input, fields_merge_pull_request_by_three_way, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MergePullRequestByThreeWay(ctx, input)
			},
		},
		"override-pull-request-approval-rules": {
			Name:   "override-pull-request-approval-rules",
			Fields: fields_override_pull_request_approval_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OverridePullRequestApprovalRulesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_override_pull_request_approval_rules, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OverridePullRequestApprovalRules(ctx, input)
			},
		},
		"post-comment-for-compared-commit": {
			Name:   "post-comment-for-compared-commit",
			Fields: fields_post_comment_for_compared_commit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostCommentForComparedCommitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_comment_for_compared_commit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostCommentForComparedCommit(ctx, input)
			},
		},
		"post-comment-for-pull-request": {
			Name:   "post-comment-for-pull-request",
			Fields: fields_post_comment_for_pull_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostCommentForPullRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_comment_for_pull_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostCommentForPullRequest(ctx, input)
			},
		},
		"post-comment-reply": {
			Name:   "post-comment-reply",
			Fields: fields_post_comment_reply,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostCommentReplyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_comment_reply, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostCommentReply(ctx, input)
			},
		},
		"put-comment-reaction": {
			Name:   "put-comment-reaction",
			Fields: fields_put_comment_reaction,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutCommentReactionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_comment_reaction, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutCommentReaction(ctx, input)
			},
		},
		"put-file": {
			Name:   "put-file",
			Fields: fields_put_file,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutFileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_file, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutFile(ctx, input)
			},
		},
		"put-repository-triggers": {
			Name:   "put-repository-triggers",
			Fields: fields_put_repository_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutRepositoryTriggersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_repository_triggers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutRepositoryTriggers(ctx, input)
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
		"test-repository-triggers": {
			Name:   "test-repository-triggers",
			Fields: fields_test_repository_triggers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestRepositoryTriggersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_repository_triggers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestRepositoryTriggers(ctx, input)
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
		"update-approval-rule-template-content": {
			Name:   "update-approval-rule-template-content",
			Fields: fields_update_approval_rule_template_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApprovalRuleTemplateContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_approval_rule_template_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApprovalRuleTemplateContent(ctx, input)
			},
		},
		"update-approval-rule-template-description": {
			Name:   "update-approval-rule-template-description",
			Fields: fields_update_approval_rule_template_description,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApprovalRuleTemplateDescriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_approval_rule_template_description, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApprovalRuleTemplateDescription(ctx, input)
			},
		},
		"update-approval-rule-template-name": {
			Name:   "update-approval-rule-template-name",
			Fields: fields_update_approval_rule_template_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateApprovalRuleTemplateNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_approval_rule_template_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateApprovalRuleTemplateName(ctx, input)
			},
		},
		"update-comment": {
			Name:   "update-comment",
			Fields: fields_update_comment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCommentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_comment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateComment(ctx, input)
			},
		},
		"update-default-branch": {
			Name:   "update-default-branch",
			Fields: fields_update_default_branch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDefaultBranchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_default_branch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDefaultBranch(ctx, input)
			},
		},
		"update-pull-request-approval-rule-content": {
			Name:   "update-pull-request-approval-rule-content",
			Fields: fields_update_pull_request_approval_rule_content,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePullRequestApprovalRuleContentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pull_request_approval_rule_content, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePullRequestApprovalRuleContent(ctx, input)
			},
		},
		"update-pull-request-approval-state": {
			Name:   "update-pull-request-approval-state",
			Fields: fields_update_pull_request_approval_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePullRequestApprovalStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pull_request_approval_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePullRequestApprovalState(ctx, input)
			},
		},
		"update-pull-request-description": {
			Name:   "update-pull-request-description",
			Fields: fields_update_pull_request_description,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePullRequestDescriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pull_request_description, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePullRequestDescription(ctx, input)
			},
		},
		"update-pull-request-status": {
			Name:   "update-pull-request-status",
			Fields: fields_update_pull_request_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePullRequestStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pull_request_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePullRequestStatus(ctx, input)
			},
		},
		"update-pull-request-title": {
			Name:   "update-pull-request-title",
			Fields: fields_update_pull_request_title,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePullRequestTitleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_pull_request_title, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePullRequestTitle(ctx, input)
			},
		},
		"update-repository-description": {
			Name:   "update-repository-description",
			Fields: fields_update_repository_description,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRepositoryDescriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_repository_description, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRepositoryDescription(ctx, input)
			},
		},
		"update-repository-encryption-key": {
			Name:   "update-repository-encryption-key",
			Fields: fields_update_repository_encryption_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRepositoryEncryptionKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_repository_encryption_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRepositoryEncryptionKey(ctx, input)
			},
		},
		"update-repository-name": {
			Name:   "update-repository-name",
			Fields: fields_update_repository_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRepositoryNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_repository_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRepositoryName(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("codecommit", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
