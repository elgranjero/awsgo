package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/codecommit"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// codecommitCmd represents the codecommit command
var _codecommitCmd = &cobra.Command{
	Use:   "codecommit",
	Short: "AWS codecommit CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := codecommit.NewFromConfig(cfg)
		if _codecommitAssociateApprovalRuleTemplateWithRepository {
			codecommit_AssociateApprovalRuleTemplateWithRepository(cfg, client)
			return
		}
		if _codecommitBatchAssociateApprovalRuleTemplateWithRepositories {
			codecommit_BatchAssociateApprovalRuleTemplateWithRepositories(cfg, client)
			return
		}
		if _codecommitBatchDescribeMergeConflicts {
			codecommit_BatchDescribeMergeConflicts(cfg, client)
			return
		}
		if _codecommitBatchDisassociateApprovalRuleTemplateFromRepositories {
			codecommit_BatchDisassociateApprovalRuleTemplateFromRepositories(cfg, client)
			return
		}
		if _codecommitBatchGetCommits {
			codecommit_BatchGetCommits(cfg, client)
			return
		}
		if _codecommitBatchGetRepositories {
			codecommit_BatchGetRepositories(cfg, client)
			return
		}
		if _codecommitCreateApprovalRuleTemplate {
			codecommit_CreateApprovalRuleTemplate(cfg, client)
			return
		}
		if _codecommitCreateBranch {
			codecommit_CreateBranch(cfg, client)
			return
		}
		if _codecommitCreateCommit {
			codecommit_CreateCommit(cfg, client)
			return
		}
		if _codecommitCreatePullRequest {
			codecommit_CreatePullRequest(cfg, client)
			return
		}
		if _codecommitCreatePullRequestApprovalRule {
			codecommit_CreatePullRequestApprovalRule(cfg, client)
			return
		}
		if _codecommitCreateRepository {
			codecommit_CreateRepository(cfg, client)
			return
		}
		if _codecommitCreateUnreferencedMergeCommit {
			codecommit_CreateUnreferencedMergeCommit(cfg, client)
			return
		}
		if _codecommitDeleteApprovalRuleTemplate {
			codecommit_DeleteApprovalRuleTemplate(cfg, client)
			return
		}
		if _codecommitDeleteBranch {
			codecommit_DeleteBranch(cfg, client)
			return
		}
		if _codecommitDeleteCommentContent {
			codecommit_DeleteCommentContent(cfg, client)
			return
		}
		if _codecommitDeleteFile {
			codecommit_DeleteFile(cfg, client)
			return
		}
		if _codecommitDeletePullRequestApprovalRule {
			codecommit_DeletePullRequestApprovalRule(cfg, client)
			return
		}
		if _codecommitDeleteRepository {
			codecommit_DeleteRepository(cfg, client)
			return
		}
		if _codecommitDescribeMergeConflicts {
			codecommit_DescribeMergeConflicts(cfg, client)
			return
		}
		if _codecommitDescribePullRequestEvents {
			codecommit_DescribePullRequestEvents(cfg, client)
			return
		}
		if _codecommitDisassociateApprovalRuleTemplateFromRepository {
			codecommit_DisassociateApprovalRuleTemplateFromRepository(cfg, client)
			return
		}
		if _codecommitEvaluatePullRequestApprovalRules {
			codecommit_EvaluatePullRequestApprovalRules(cfg, client)
			return
		}
		if _codecommitGetApprovalRuleTemplate {
			codecommit_GetApprovalRuleTemplate(cfg, client)
			return
		}
		if _codecommitGetBlob {
			codecommit_GetBlob(cfg, client)
			return
		}
		if _codecommitGetBranch {
			codecommit_GetBranch(cfg, client)
			return
		}
		if _codecommitGetComment {
			codecommit_GetComment(cfg, client)
			return
		}
		if _codecommitGetCommentReactions {
			codecommit_GetCommentReactions(cfg, client)
			return
		}
		if _codecommitGetCommentsForComparedCommit {
			codecommit_GetCommentsForComparedCommit(cfg, client)
			return
		}
		if _codecommitGetCommentsForPullRequest {
			codecommit_GetCommentsForPullRequest(cfg, client)
			return
		}
		if _codecommitGetCommit {
			codecommit_GetCommit(cfg, client)
			return
		}
		if _codecommitGetDifferences {
			codecommit_GetDifferences(cfg, client)
			return
		}
		if _codecommitGetFile {
			codecommit_GetFile(cfg, client)
			return
		}
		if _codecommitGetFolder {
			codecommit_GetFolder(cfg, client)
			return
		}
		if _codecommitGetMergeCommit {
			codecommit_GetMergeCommit(cfg, client)
			return
		}
		if _codecommitGetMergeConflicts {
			codecommit_GetMergeConflicts(cfg, client)
			return
		}
		if _codecommitGetMergeOptions {
			codecommit_GetMergeOptions(cfg, client)
			return
		}
		if _codecommitGetPullRequest {
			codecommit_GetPullRequest(cfg, client)
			return
		}
		if _codecommitGetPullRequestApprovalStates {
			codecommit_GetPullRequestApprovalStates(cfg, client)
			return
		}
		if _codecommitGetPullRequestOverrideState {
			codecommit_GetPullRequestOverrideState(cfg, client)
			return
		}
		if _codecommitGetRepository {
			codecommit_GetRepository(cfg, client)
			return
		}
		if _codecommitGetRepositoryTriggers {
			codecommit_GetRepositoryTriggers(cfg, client)
			return
		}
		if _codecommitListApprovalRuleTemplates {
			codecommit_ListApprovalRuleTemplates(cfg, client)
			return
		}
		if _codecommitListAssociatedApprovalRuleTemplatesForRepository {
			codecommit_ListAssociatedApprovalRuleTemplatesForRepository(cfg, client)
			return
		}
		if _codecommitListBranches {
			codecommit_ListBranches(cfg, client)
			return
		}
		if _codecommitListFileCommitHistory {
			codecommit_ListFileCommitHistory(cfg, client)
			return
		}
		if _codecommitListPullRequests {
			codecommit_ListPullRequests(cfg, client)
			return
		}
		if _codecommitListRepositories {
			codecommit_ListRepositories(cfg, client)
			return
		}
		if _codecommitListRepositoriesForApprovalRuleTemplate {
			codecommit_ListRepositoriesForApprovalRuleTemplate(cfg, client)
			return
		}
		if _codecommitListTagsForResource {
			codecommit_ListTagsForResource(cfg, client)
			return
		}
		if _codecommitMergeBranchesByFastForward {
			codecommit_MergeBranchesByFastForward(cfg, client)
			return
		}
		if _codecommitMergeBranchesBySquash {
			codecommit_MergeBranchesBySquash(cfg, client)
			return
		}
		if _codecommitMergeBranchesByThreeWay {
			codecommit_MergeBranchesByThreeWay(cfg, client)
			return
		}
		if _codecommitMergePullRequestByFastForward {
			codecommit_MergePullRequestByFastForward(cfg, client)
			return
		}
		if _codecommitMergePullRequestBySquash {
			codecommit_MergePullRequestBySquash(cfg, client)
			return
		}
		if _codecommitMergePullRequestByThreeWay {
			codecommit_MergePullRequestByThreeWay(cfg, client)
			return
		}
		if _codecommitOverridePullRequestApprovalRules {
			codecommit_OverridePullRequestApprovalRules(cfg, client)
			return
		}
		if _codecommitPostCommentForComparedCommit {
			codecommit_PostCommentForComparedCommit(cfg, client)
			return
		}
		if _codecommitPostCommentForPullRequest {
			codecommit_PostCommentForPullRequest(cfg, client)
			return
		}
		if _codecommitPostCommentReply {
			codecommit_PostCommentReply(cfg, client)
			return
		}
		if _codecommitPutCommentReaction {
			codecommit_PutCommentReaction(cfg, client)
			return
		}
		if _codecommitPutFile {
			codecommit_PutFile(cfg, client)
			return
		}
		if _codecommitPutRepositoryTriggers {
			codecommit_PutRepositoryTriggers(cfg, client)
			return
		}
		if _codecommitTagResource {
			codecommit_TagResource(cfg, client)
			return
		}
		if _codecommitTestRepositoryTriggers {
			codecommit_TestRepositoryTriggers(cfg, client)
			return
		}
		if _codecommitUntagResource {
			codecommit_UntagResource(cfg, client)
			return
		}
		if _codecommitUpdateApprovalRuleTemplateContent {
			codecommit_UpdateApprovalRuleTemplateContent(cfg, client)
			return
		}
		if _codecommitUpdateApprovalRuleTemplateDescription {
			codecommit_UpdateApprovalRuleTemplateDescription(cfg, client)
			return
		}
		if _codecommitUpdateApprovalRuleTemplateName {
			codecommit_UpdateApprovalRuleTemplateName(cfg, client)
			return
		}
		if _codecommitUpdateComment {
			codecommit_UpdateComment(cfg, client)
			return
		}
		if _codecommitUpdateDefaultBranch {
			codecommit_UpdateDefaultBranch(cfg, client)
			return
		}
		if _codecommitUpdatePullRequestApprovalRuleContent {
			codecommit_UpdatePullRequestApprovalRuleContent(cfg, client)
			return
		}
		if _codecommitUpdatePullRequestApprovalState {
			codecommit_UpdatePullRequestApprovalState(cfg, client)
			return
		}
		if _codecommitUpdatePullRequestDescription {
			codecommit_UpdatePullRequestDescription(cfg, client)
			return
		}
		if _codecommitUpdatePullRequestStatus {
			codecommit_UpdatePullRequestStatus(cfg, client)
			return
		}
		if _codecommitUpdatePullRequestTitle {
			codecommit_UpdatePullRequestTitle(cfg, client)
			return
		}
		if _codecommitUpdateRepositoryDescription {
			codecommit_UpdateRepositoryDescription(cfg, client)
			return
		}
		if _codecommitUpdateRepositoryEncryptionKey {
			codecommit_UpdateRepositoryEncryptionKey(cfg, client)
			return
		}
		if _codecommitUpdateRepositoryName {
			codecommit_UpdateRepositoryName(cfg, client)
			return
		}

	},
}

var (
	_codecommitAssociateApprovalRuleTemplateWithRepository           bool
	_codecommitBatchAssociateApprovalRuleTemplateWithRepositories    bool
	_codecommitBatchDescribeMergeConflicts                           bool
	_codecommitBatchDisassociateApprovalRuleTemplateFromRepositories bool
	_codecommitBatchGetCommits                                       bool
	_codecommitBatchGetRepositories                                  bool
	_codecommitCreateApprovalRuleTemplate                            bool
	_codecommitCreateBranch                                          bool
	_codecommitCreateCommit                                          bool
	_codecommitCreatePullRequest                                     bool
	_codecommitCreatePullRequestApprovalRule                         bool
	_codecommitCreateRepository                                      bool
	_codecommitCreateUnreferencedMergeCommit                         bool
	_codecommitDeleteApprovalRuleTemplate                            bool
	_codecommitDeleteBranch                                          bool
	_codecommitDeleteCommentContent                                  bool
	_codecommitDeleteFile                                            bool
	_codecommitDeletePullRequestApprovalRule                         bool
	_codecommitDeleteRepository                                      bool
	_codecommitDescribeMergeConflicts                                bool
	_codecommitDescribePullRequestEvents                             bool
	_codecommitDisassociateApprovalRuleTemplateFromRepository        bool
	_codecommitEvaluatePullRequestApprovalRules                      bool
	_codecommitGetApprovalRuleTemplate                               bool
	_codecommitGetBlob                                               bool
	_codecommitGetBranch                                             bool
	_codecommitGetComment                                            bool
	_codecommitGetCommentReactions                                   bool
	_codecommitGetCommentsForComparedCommit                          bool
	_codecommitGetCommentsForPullRequest                             bool
	_codecommitGetCommit                                             bool
	_codecommitGetDifferences                                        bool
	_codecommitGetFile                                               bool
	_codecommitGetFolder                                             bool
	_codecommitGetMergeCommit                                        bool
	_codecommitGetMergeConflicts                                     bool
	_codecommitGetMergeOptions                                       bool
	_codecommitGetPullRequest                                        bool
	_codecommitGetPullRequestApprovalStates                          bool
	_codecommitGetPullRequestOverrideState                           bool
	_codecommitGetRepository                                         bool
	_codecommitGetRepositoryTriggers                                 bool
	_codecommitListApprovalRuleTemplates                             bool
	_codecommitListAssociatedApprovalRuleTemplatesForRepository      bool
	_codecommitListBranches                                          bool
	_codecommitListFileCommitHistory                                 bool
	_codecommitListPullRequests                                      bool
	_codecommitListRepositories                                      bool
	_codecommitListRepositoriesForApprovalRuleTemplate               bool
	_codecommitListTagsForResource                                   bool
	_codecommitMergeBranchesByFastForward                            bool
	_codecommitMergeBranchesBySquash                                 bool
	_codecommitMergeBranchesByThreeWay                               bool
	_codecommitMergePullRequestByFastForward                         bool
	_codecommitMergePullRequestBySquash                              bool
	_codecommitMergePullRequestByThreeWay                            bool
	_codecommitOverridePullRequestApprovalRules                      bool
	_codecommitPostCommentForComparedCommit                          bool
	_codecommitPostCommentForPullRequest                             bool
	_codecommitPostCommentReply                                      bool
	_codecommitPutCommentReaction                                    bool
	_codecommitPutFile                                               bool
	_codecommitPutRepositoryTriggers                                 bool
	_codecommitTagResource                                           bool
	_codecommitTestRepositoryTriggers                                bool
	_codecommitUntagResource                                         bool
	_codecommitUpdateApprovalRuleTemplateContent                     bool
	_codecommitUpdateApprovalRuleTemplateDescription                 bool
	_codecommitUpdateApprovalRuleTemplateName                        bool
	_codecommitUpdateComment                                         bool
	_codecommitUpdateDefaultBranch                                   bool
	_codecommitUpdatePullRequestApprovalRuleContent                  bool
	_codecommitUpdatePullRequestApprovalState                        bool
	_codecommitUpdatePullRequestDescription                          bool
	_codecommitUpdatePullRequestStatus                               bool
	_codecommitUpdatePullRequestTitle                                bool
	_codecommitUpdateRepositoryDescription                           bool
	_codecommitUpdateRepositoryEncryptionKey                         bool
	_codecommitUpdateRepositoryName                                  bool

	_codecommitActorArn                        string
	_codecommitAfterCommitId                   string
	_codecommitAfterCommitSpecifier            string
	_codecommitAfterPath                       string
	_codecommitApprovalRuleContent             string
	_codecommitApprovalRuleName                string
	_codecommitApprovalRuleTemplateContent     string
	_codecommitApprovalRuleTemplateDescription string
	_codecommitApprovalRuleTemplateName        string
	_codecommitApprovalState                   string
	_codecommitAuthorArn                       string
	_codecommitAuthorName                      string
	_codecommitBeforeCommitId                  string
	_codecommitBeforeCommitSpecifier           string
	_codecommitBeforePath                      string
	_codecommitBlobId                          string
	_codecommitBranchName                      string
	_codecommitClientRequestToken              string
	_codecommitCommentId                       string
	_codecommitCommitId                        string
	_codecommitCommitIds                       []string
	_codecommitCommitMessage                   string
	_codecommitCommitSpecifier                 string
	_codecommitConflictDetailLevel             string
	_codecommitConflictResolution              string
	_codecommitConflictResolutionStrategy      string
	_codecommitContent                         string
	_codecommitDefaultBranchName               string
	_codecommitDeleteFiles                     string
	_codecommitDescription                     string
	_codecommitDestinationCommitSpecifier      string
	_codecommitEmail                           string
	_codecommitExistingRuleContentSha256       string
	_codecommitFileContent                     string
	_codecommitFileMode                        string
	_codecommitFilePath                        string
	_codecommitFilePaths                       []string
	_codecommitFolderPath                      string
	_codecommitInReplyTo                       string
	_codecommitKeepEmptyFolders                string
	_codecommitKmsKeyId                        string
	_codecommitLocation                        string
	_codecommitMaxConflictFiles                string
	_codecommitMaxMergeHunks                   string
	_codecommitMaxResults                      string
	_codecommitMergeOption                     string
	_codecommitName                            string
	_codecommitNewApprovalRuleTemplateName     string
	_codecommitNewName                         string
	_codecommitNewRuleContent                  string
	_codecommitNextToken                       string
	_codecommitOldApprovalRuleTemplateName     string
	_codecommitOldName                         string
	_codecommitOrder                           string
	_codecommitOverrideStatus                  string
	_codecommitParentCommitId                  string
	_codecommitPullRequestEventType            string
	_codecommitPullRequestId                   string
	_codecommitPullRequestStatus               string
	_codecommitPutFiles                        string
	_codecommitReactionUserArn                 string
	_codecommitReactionValue                   string
	_codecommitRepositoryDescription           string
	_codecommitRepositoryName                  string
	_codecommitRepositoryNames                 []string
	_codecommitResourceArn                     string
	_codecommitRevisionId                      string
	_codecommitSetFileModes                    string
	_codecommitSortBy                          string
	_codecommitSourceCommitId                  string
	_codecommitSourceCommitSpecifier           string
	_codecommitTagKeys                         []string
	_codecommitTags                            string
	_codecommitTargetBranch                    string
	_codecommitTargets                         string
	_codecommitTitle                           string
	_codecommitTriggers                        string
)

// Creates an association between an approval rule template and a specified
// repository. Then, the next time a pull request is created in the repository
// where the destination reference (if specified) matches the destination reference
// (branch) for the pull request, an approval rule that matches the template
// conditions is automatically created for that pull request. If no destination
// references are specified in the template, an approval rule that matches the
// template contents is created for all pull requests in that repository.
func codecommit_AssociateApprovalRuleTemplateWithRepository(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.AssociateApprovalRuleTemplateWithRepositoryInput{
		// ApprovalRuleTemplateName: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.AssociateApprovalRuleTemplateWithRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an association between an approval rule template and one or more
// specified repositories.
func codecommit_BatchAssociateApprovalRuleTemplateWithRepositories(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.BatchAssociateApprovalRuleTemplateWithRepositoriesInput{
		// ApprovalRuleTemplateName: *string, // Required
		// RepositoryNames: []string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _codecommitRepositoryNames...)
	}

	if resp, err := client.BatchAssociateApprovalRuleTemplateWithRepositories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more merge conflicts in the attempted merge of
// two commit specifiers using the squash or three-way merge strategy.
func codecommit_BatchDescribeMergeConflicts(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.BatchDescribeMergeConflictsInput{
		// DestinationCommitSpecifier: *string, // Required
		// MergeOption: types.MergeOptionTypeEnum, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitMergeOption) > 0 {
		if err := assignInputField(input, "MergeOption", _codecommitMergeOption); err != nil {
			log.Errorf("invalid --merge-option: %s", err.Error())
			return
		}
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitFilePaths) > 0 {
		input.FilePaths = append([]string(nil), _codecommitFilePaths...)
	}
	if len(_codecommitMaxConflictFiles) > 0 {
		if err := assignInputField(input, "MaxConflictFiles", _codecommitMaxConflictFiles); err != nil {
			log.Errorf("invalid --max-conflict-files: %s", err.Error())
			return
		}
	}
	if len(_codecommitMaxMergeHunks) > 0 {
		if err := assignInputField(input, "MaxMergeHunks", _codecommitMaxMergeHunks); err != nil {
			log.Errorf("invalid --max-merge-hunks: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if resp, err := client.BatchDescribeMergeConflicts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between an approval rule template and one or more
// specified repositories.
func codecommit_BatchDisassociateApprovalRuleTemplateFromRepositories(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.BatchDisassociateApprovalRuleTemplateFromRepositoriesInput{
		// ApprovalRuleTemplateName: *string, // Required
		// RepositoryNames: []string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _codecommitRepositoryNames...)
	}

	if resp, err := client.BatchDisassociateApprovalRuleTemplateFromRepositories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the contents of one or more commits in a repository.
func codecommit_BatchGetCommits(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.BatchGetCommitsInput{
		// CommitIds: []string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitCommitIds) > 0 {
		input.CommitIds = append([]string(nil), _codecommitCommitIds...)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.BatchGetCommits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more repositories.
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a webpage can expose users to potentially malicious code. Make
// sure that you HTML-encode the description field in any application that uses
// this API to display the repository description on a webpage.
func codecommit_BatchGetRepositories(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.BatchGetRepositoriesInput{
		// RepositoryNames: []string, // Required
	}

	if len(_codecommitRepositoryNames) > 0 {
		input.RepositoryNames = append([]string(nil), _codecommitRepositoryNames...)
	}

	if resp, err := client.BatchGetRepositories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a template for approval rules that can then be associated with one or
// more repositories in your Amazon Web Services account. When you associate a
// template with a repository, CodeCommit creates an approval rule that matches the
// conditions of the template for all pull requests that meet the conditions of the
// template. For more information, see AssociateApprovalRuleTemplateWithRepository.
func codecommit_CreateApprovalRuleTemplate(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreateApprovalRuleTemplateInput{
		// ApprovalRuleTemplateContent: *string, // Required
		// ApprovalRuleTemplateName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateContent) > 0 {
		input.ApprovalRuleTemplateContent = aws.String(_codecommitApprovalRuleTemplateContent)
	}
	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitApprovalRuleTemplateDescription) > 0 {
		input.ApprovalRuleTemplateDescription = aws.String(_codecommitApprovalRuleTemplateDescription)
	}

	if resp, err := client.CreateApprovalRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a branch in a repository and points the branch to a commit.
// Calling the create branch operation does not set a repository's default branch.
// To do this, call the update default branch operation.
func codecommit_CreateBranch(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreateBranchInput{
		// BranchName: *string, // Required
		// CommitId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitBranchName) > 0 {
		input.BranchName = aws.String(_codecommitBranchName)
	}
	if len(_codecommitCommitId) > 0 {
		input.CommitId = aws.String(_codecommitCommitId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.CreateBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a commit for a repository on the tip of a specified branch.
func codecommit_CreateCommit(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreateCommitInput{
		// BranchName: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitBranchName) > 0 {
		input.BranchName = aws.String(_codecommitBranchName)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitAuthorName) > 0 {
		input.AuthorName = aws.String(_codecommitAuthorName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitDeleteFiles) > 0 {
		if err := assignInputField(input, "DeleteFiles", _codecommitDeleteFiles); err != nil {
			log.Errorf("invalid --delete-files: %s", err.Error())
			return
		}
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}
	if len(_codecommitParentCommitId) > 0 {
		input.ParentCommitId = aws.String(_codecommitParentCommitId)
	}
	if len(_codecommitPutFiles) > 0 {
		if err := assignInputField(input, "PutFiles", _codecommitPutFiles); err != nil {
			log.Errorf("invalid --put-files: %s", err.Error())
			return
		}
	}
	if len(_codecommitSetFileModes) > 0 {
		if err := assignInputField(input, "SetFileModes", _codecommitSetFileModes); err != nil {
			log.Errorf("invalid --set-file-modes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCommit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a pull request in the specified repository.
func codecommit_CreatePullRequest(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreatePullRequestInput{
		// Targets: []types.Target, // Required
		// Title: *string, // Required
	}

	if len(_codecommitTargets) > 0 {
		if err := assignInputField(input, "Targets", _codecommitTargets); err != nil {
			log.Errorf("invalid --targets: %s", err.Error())
			return
		}
	}
	if len(_codecommitTitle) > 0 {
		input.Title = aws.String(_codecommitTitle)
	}
	if len(_codecommitClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codecommitClientRequestToken)
	}
	if len(_codecommitDescription) > 0 {
		input.Description = aws.String(_codecommitDescription)
	}

	if resp, err := client.CreatePullRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an approval rule for a pull request.
func codecommit_CreatePullRequestApprovalRule(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreatePullRequestApprovalRuleInput{
		// ApprovalRuleContent: *string, // Required
		// ApprovalRuleName: *string, // Required
		// PullRequestId: *string, // Required
	}

	if len(_codecommitApprovalRuleContent) > 0 {
		input.ApprovalRuleContent = aws.String(_codecommitApprovalRuleContent)
	}
	if len(_codecommitApprovalRuleName) > 0 {
		input.ApprovalRuleName = aws.String(_codecommitApprovalRuleName)
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}

	if resp, err := client.CreatePullRequestApprovalRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new, empty repository.
func codecommit_CreateRepository(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreateRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_codecommitKmsKeyId)
	}
	if len(_codecommitRepositoryDescription) > 0 {
		input.RepositoryDescription = aws.String(_codecommitRepositoryDescription)
	}
	if len(_codecommitTags) > 0 {
		if err := assignInputField(input, "Tags", _codecommitTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an unreferenced commit that represents the result of merging two
// branches using a specified merge strategy. This can help you determine the
// outcome of a potential merge. This API cannot be used with the fast-forward
// merge strategy because that strategy does not create a merge commit.
//
// This unreferenced merge commit can only be accessed using the GetCommit API or
// through git commands such as git fetch. To retrieve this commit, you must
// specify its commit ID or otherwise reference it.
func codecommit_CreateUnreferencedMergeCommit(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.CreateUnreferencedMergeCommitInput{
		// DestinationCommitSpecifier: *string, // Required
		// MergeOption: types.MergeOptionTypeEnum, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitMergeOption) > 0 {
		if err := assignInputField(input, "MergeOption", _codecommitMergeOption); err != nil {
			log.Errorf("invalid --merge-option: %s", err.Error())
			return
		}
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitAuthorName) > 0 {
		input.AuthorName = aws.String(_codecommitAuthorName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolution) > 0 {
		if err := assignInputField(input, "ConflictResolution", _codecommitConflictResolution); err != nil {
			log.Errorf("invalid --conflict-resolution: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUnreferencedMergeCommit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified approval rule template. Deleting a template does not remove
// approval rules on pull requests already created with the template.
func codecommit_DeleteApprovalRuleTemplate(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DeleteApprovalRuleTemplateInput{
		// ApprovalRuleTemplateName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}

	if resp, err := client.DeleteApprovalRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a branch from a repository, unless that branch is the default branch
// for the repository.
func codecommit_DeleteBranch(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DeleteBranchInput{
		// BranchName: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitBranchName) > 0 {
		input.BranchName = aws.String(_codecommitBranchName)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.DeleteBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the content of a comment made on a change, file, or commit in a
// repository.
func codecommit_DeleteCommentContent(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DeleteCommentContentInput{
		// CommentId: *string, // Required
	}

	if len(_codecommitCommentId) > 0 {
		input.CommentId = aws.String(_codecommitCommentId)
	}

	if resp, err := client.DeleteCommentContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified file from a specified branch. A commit is created on the
// branch that contains the revision. The file still exists in the commits earlier
// to the commit that contains the deletion.
func codecommit_DeleteFile(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DeleteFileInput{
		// BranchName: *string, // Required
		// FilePath: *string, // Required
		// ParentCommitId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitBranchName) > 0 {
		input.BranchName = aws.String(_codecommitBranchName)
	}
	if len(_codecommitFilePath) > 0 {
		input.FilePath = aws.String(_codecommitFilePath)
	}
	if len(_codecommitParentCommitId) > 0 {
		input.ParentCommitId = aws.String(_codecommitParentCommitId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}
	if len(_codecommitName) > 0 {
		input.Name = aws.String(_codecommitName)
	}

	if resp, err := client.DeleteFile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an approval rule from a specified pull request. Approval rules can be
// deleted from a pull request only if the pull request is open, and if the
// approval rule was created specifically for a pull request and not generated from
// an approval rule template associated with the repository where the pull request
// was created. You cannot delete an approval rule from a merged or closed pull
// request.
func codecommit_DeletePullRequestApprovalRule(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DeletePullRequestApprovalRuleInput{
		// ApprovalRuleName: *string, // Required
		// PullRequestId: *string, // Required
	}

	if len(_codecommitApprovalRuleName) > 0 {
		input.ApprovalRuleName = aws.String(_codecommitApprovalRuleName)
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}

	if resp, err := client.DeletePullRequestApprovalRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a repository. If a specified repository was already deleted, a null
// repository ID is returned.
//
// Deleting a repository also deletes all associated objects and metadata. After a
// repository is deleted, all future push calls to the deleted repository fail.
func codecommit_DeleteRepository(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DeleteRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.DeleteRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about one or more merge conflicts in the attempted merge of
// two commit specifiers using the squash or three-way merge strategy. If the merge
// option for the attempted merge is specified as FAST_FORWARD_MERGE, an exception
// is thrown.
func codecommit_DescribeMergeConflicts(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DescribeMergeConflictsInput{
		// DestinationCommitSpecifier: *string, // Required
		// FilePath: *string, // Required
		// MergeOption: types.MergeOptionTypeEnum, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitFilePath) > 0 {
		input.FilePath = aws.String(_codecommitFilePath)
	}
	if len(_codecommitMergeOption) > 0 {
		if err := assignInputField(input, "MergeOption", _codecommitMergeOption); err != nil {
			log.Errorf("invalid --merge-option: %s", err.Error())
			return
		}
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitMaxMergeHunks) > 0 {
		if err := assignInputField(input, "MaxMergeHunks", _codecommitMaxMergeHunks); err != nil {
			log.Errorf("invalid --max-merge-hunks: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMergeConflicts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.DescribeMergeConflictsOutput
	p := codecommit.NewDescribeMergeConflictsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about one or more pull request events.
func codecommit_DescribePullRequestEvents(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DescribePullRequestEventsInput{
		// PullRequestId: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitActorArn) > 0 {
		input.ActorArn = aws.String(_codecommitActorArn)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}
	if len(_codecommitPullRequestEventType) > 0 {
		if err := assignInputField(input, "PullRequestEventType", _codecommitPullRequestEventType); err != nil {
			log.Errorf("invalid --pull-request-event-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribePullRequestEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.DescribePullRequestEventsOutput
	p := codecommit.NewDescribePullRequestEventsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Removes the association between a template and a repository so that approval
// rules based on the template are not automatically created when pull requests are
// created in the specified repository. This does not delete any approval rules
// previously created for pull requests through the template association.
func codecommit_DisassociateApprovalRuleTemplateFromRepository(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.DisassociateApprovalRuleTemplateFromRepositoryInput{
		// ApprovalRuleTemplateName: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.DisassociateApprovalRuleTemplateFromRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates whether a pull request has met all the conditions specified in its
// associated approval rules.
func codecommit_EvaluatePullRequestApprovalRules(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.EvaluatePullRequestApprovalRulesInput{
		// PullRequestId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRevisionId) > 0 {
		input.RevisionId = aws.String(_codecommitRevisionId)
	}

	if resp, err := client.EvaluatePullRequestApprovalRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified approval rule template.
func codecommit_GetApprovalRuleTemplate(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetApprovalRuleTemplateInput{
		// ApprovalRuleTemplateName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}

	if resp, err := client.GetApprovalRuleTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the base-64 encoded content of an individual blob in a repository.
func codecommit_GetBlob(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetBlobInput{
		// BlobId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitBlobId) > 0 {
		input.BlobId = aws.String(_codecommitBlobId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.GetBlob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a repository branch, including its name and the last
// commit ID.
func codecommit_GetBranch(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetBranchInput{}

	if len(_codecommitBranchName) > 0 {
		input.BranchName = aws.String(_codecommitBranchName)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.GetBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the content of a comment made on a change, file, or commit in a
// repository.
//
// Reaction counts might include numbers from user identities who were deleted
// after the reaction was made. For a count of reactions from active identities,
// use GetCommentReactions.
func codecommit_GetComment(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetCommentInput{
		// CommentId: *string, // Required
	}

	if len(_codecommitCommentId) > 0 {
		input.CommentId = aws.String(_codecommitCommentId)
	}

	if resp, err := client.GetComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about reactions to a specified comment ID. Reactions from
// users who have been deleted will not be included in the count.
func codecommit_GetCommentReactions(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetCommentReactionsInput{
		// CommentId: *string, // Required
	}

	if len(_codecommitCommentId) > 0 {
		input.CommentId = aws.String(_codecommitCommentId)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}
	if len(_codecommitReactionUserArn) > 0 {
		input.ReactionUserArn = aws.String(_codecommitReactionUserArn)
	}

	if disablePaginator() {
		if resp, err := client.GetCommentReactions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.GetCommentReactionsOutput
	p := codecommit.NewGetCommentReactionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about comments made on the comparison between two commits.
// Reaction counts might include numbers from user identities who were deleted
// after the reaction was made. For a count of reactions from active identities,
// use GetCommentReactions.
func codecommit_GetCommentsForComparedCommit(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetCommentsForComparedCommitInput{
		// AfterCommitId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitAfterCommitId) > 0 {
		input.AfterCommitId = aws.String(_codecommitAfterCommitId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitBeforeCommitId) > 0 {
		input.BeforeCommitId = aws.String(_codecommitBeforeCommitId)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetCommentsForComparedCommit(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.GetCommentsForComparedCommitOutput
	p := codecommit.NewGetCommentsForComparedCommitPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns comments made on a pull request.
// Reaction counts might include numbers from user identities who were deleted
// after the reaction was made. For a count of reactions from active identities,
// use GetCommentReactions.
func codecommit_GetCommentsForPullRequest(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetCommentsForPullRequestInput{
		// PullRequestId: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitAfterCommitId) > 0 {
		input.AfterCommitId = aws.String(_codecommitAfterCommitId)
	}
	if len(_codecommitBeforeCommitId) > 0 {
		input.BeforeCommitId = aws.String(_codecommitBeforeCommitId)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if disablePaginator() {
		if resp, err := client.GetCommentsForPullRequest(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.GetCommentsForPullRequestOutput
	p := codecommit.NewGetCommentsForPullRequestPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about a commit, including commit message and committer
// information.
func codecommit_GetCommit(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetCommitInput{
		// CommitId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitCommitId) > 0 {
		input.CommitId = aws.String(_codecommitCommitId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.GetCommit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the differences in a valid commit specifier (such as
// a branch, tag, HEAD, commit ID, or other fully qualified reference). Results can
// be limited to a specified path.
func codecommit_GetDifferences(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetDifferencesInput{
		// AfterCommitSpecifier: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitAfterCommitSpecifier) > 0 {
		input.AfterCommitSpecifier = aws.String(_codecommitAfterCommitSpecifier)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitAfterPath) > 0 {
		input.AfterPath = aws.String(_codecommitAfterPath)
	}
	if len(_codecommitBeforeCommitSpecifier) > 0 {
		input.BeforeCommitSpecifier = aws.String(_codecommitBeforeCommitSpecifier)
	}
	if len(_codecommitBeforePath) > 0 {
		input.BeforePath = aws.String(_codecommitBeforePath)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetDifferences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.GetDifferencesOutput
	p := codecommit.NewGetDifferencesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns the base-64 encoded contents of a specified file and its metadata.
func codecommit_GetFile(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetFileInput{
		// FilePath: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitFilePath) > 0 {
		input.FilePath = aws.String(_codecommitFilePath)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitCommitSpecifier) > 0 {
		input.CommitSpecifier = aws.String(_codecommitCommitSpecifier)
	}

	if resp, err := client.GetFile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the contents of a specified folder in a repository.
func codecommit_GetFolder(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetFolderInput{
		// FolderPath: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitFolderPath) > 0 {
		input.FolderPath = aws.String(_codecommitFolderPath)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitCommitSpecifier) > 0 {
		input.CommitSpecifier = aws.String(_codecommitCommitSpecifier)
	}

	if resp, err := client.GetFolder(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specified merge commit.
func codecommit_GetMergeCommit(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetMergeCommitInput{
		// DestinationCommitSpecifier: *string, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMergeCommit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about merge conflicts between the before and after commit
// IDs for a pull request in a repository.
func codecommit_GetMergeConflicts(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetMergeConflictsInput{
		// DestinationCommitSpecifier: *string, // Required
		// MergeOption: types.MergeOptionTypeEnum, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitMergeOption) > 0 {
		if err := assignInputField(input, "MergeOption", _codecommitMergeOption); err != nil {
			log.Errorf("invalid --merge-option: %s", err.Error())
			return
		}
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitMaxConflictFiles) > 0 {
		if err := assignInputField(input, "MaxConflictFiles", _codecommitMaxConflictFiles); err != nil {
			log.Errorf("invalid --max-conflict-files: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetMergeConflicts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.GetMergeConflictsOutput
	p := codecommit.NewGetMergeConflictsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns information about the merge options available for merging two specified
// branches. For details about why a merge option is not available, use
// GetMergeConflicts or DescribeMergeConflicts.
func codecommit_GetMergeOptions(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetMergeOptionsInput{
		// DestinationCommitSpecifier: *string, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMergeOptions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a pull request in a specified repository.
func codecommit_GetPullRequest(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetPullRequestInput{
		// PullRequestId: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}

	if resp, err := client.GetPullRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the approval states for a specified pull request.
// Approval states only apply to pull requests that have one or more approval rules
// applied to them.
func codecommit_GetPullRequestApprovalStates(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetPullRequestApprovalStatesInput{
		// PullRequestId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRevisionId) > 0 {
		input.RevisionId = aws.String(_codecommitRevisionId)
	}

	if resp, err := client.GetPullRequestApprovalStates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about whether approval rules have been set aside
// (overridden) for a pull request, and if so, the Amazon Resource Name (ARN) of
// the user or identity that overrode the rules and their requirements for the pull
// request.
func codecommit_GetPullRequestOverrideState(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetPullRequestOverrideStateInput{
		// PullRequestId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRevisionId) > 0 {
		input.RevisionId = aws.String(_codecommitRevisionId)
	}

	if resp, err := client.GetPullRequestOverrideState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a repository.
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a webpage can expose users to potentially malicious code. Make
// sure that you HTML-encode the description field in any application that uses
// this API to display the repository description on a webpage.
func codecommit_GetRepository(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.GetRepository(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about triggers configured for a repository.
func codecommit_GetRepositoryTriggers(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.GetRepositoryTriggersInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.GetRepositoryTriggers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all approval rule templates in the specified Amazon Web Services Region
// in your Amazon Web Services account. If an Amazon Web Services Region is not
// specified, the Amazon Web Services Region where you are signed in is used.
func codecommit_ListApprovalRuleTemplates(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListApprovalRuleTemplatesInput{}

	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApprovalRuleTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListApprovalRuleTemplatesOutput
	p := codecommit.NewListApprovalRuleTemplatesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all approval rule templates that are associated with a specified
// repository.
func codecommit_ListAssociatedApprovalRuleTemplatesForRepository(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssociatedApprovalRuleTemplatesForRepository(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListAssociatedApprovalRuleTemplatesForRepositoryOutput
	p := codecommit.NewListAssociatedApprovalRuleTemplatesForRepositoryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Gets information about one or more branches in a repository.
func codecommit_ListBranches(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListBranchesInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBranches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListBranchesOutput
	p := codecommit.NewListBranchesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves a list of commits and changes to a specified file.
func codecommit_ListFileCommitHistory(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListFileCommitHistoryInput{
		// FilePath: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitFilePath) > 0 {
		input.FilePath = aws.String(_codecommitFilePath)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitCommitSpecifier) > 0 {
		input.CommitSpecifier = aws.String(_codecommitCommitSpecifier)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFileCommitHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListFileCommitHistoryOutput
	p := codecommit.NewListFileCommitHistoryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of pull requests for a specified repository. The return list can
// be refined by pull request status or pull request author ARN.
func codecommit_ListPullRequests(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListPullRequestsInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitAuthorArn) > 0 {
		input.AuthorArn = aws.String(_codecommitAuthorArn)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}
	if len(_codecommitPullRequestStatus) > 0 {
		if err := assignInputField(input, "PullRequestStatus", _codecommitPullRequestStatus); err != nil {
			log.Errorf("invalid --pull-request-status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPullRequests(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListPullRequestsOutput
	p := codecommit.NewListPullRequestsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Gets information about one or more repositories.
func codecommit_ListRepositories(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListRepositoriesInput{}

	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}
	if len(_codecommitOrder) > 0 {
		if err := assignInputField(input, "Order", _codecommitOrder); err != nil {
			log.Errorf("invalid --order: %s", err.Error())
			return
		}
	}
	if len(_codecommitSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _codecommitSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRepositories(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListRepositoriesOutput
	p := codecommit.NewListRepositoriesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all repositories associated with the specified approval rule template.
func codecommit_ListRepositoriesForApprovalRuleTemplate(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListRepositoriesForApprovalRuleTemplateInput{
		// ApprovalRuleTemplateName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _codecommitMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRepositoriesForApprovalRuleTemplate(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*codecommit.ListRepositoriesForApprovalRuleTemplateOutput
	p := codecommit.NewListRepositoriesForApprovalRuleTemplatePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Gets information about Amazon Web Servicestags for a specified Amazon Resource
// Name (ARN) in CodeCommit. For a list of valid resources in CodeCommit, see [CodeCommit Resources and Operations]in
// the CodeCommit User Guide.
//
// [CodeCommit Resources and Operations]: https://docs.aws.amazon.com/codecommit/latest/userguide/auth-and-access-control-iam-access-control-identity-based.html#arn-formats
func codecommit_ListTagsForResource(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_codecommitResourceArn) > 0 {
		input.ResourceArn = aws.String(_codecommitResourceArn)
	}
	if len(_codecommitNextToken) > 0 {
		input.NextToken = aws.String(_codecommitNextToken)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Merges two branches using the fast-forward merge strategy.
func codecommit_MergeBranchesByFastForward(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.MergeBranchesByFastForwardInput{
		// DestinationCommitSpecifier: *string, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitTargetBranch) > 0 {
		input.TargetBranch = aws.String(_codecommitTargetBranch)
	}

	if resp, err := client.MergeBranchesByFastForward(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Merges two branches using the squash merge strategy.
func codecommit_MergeBranchesBySquash(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.MergeBranchesBySquashInput{
		// DestinationCommitSpecifier: *string, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitAuthorName) > 0 {
		input.AuthorName = aws.String(_codecommitAuthorName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolution) > 0 {
		if err := assignInputField(input, "ConflictResolution", _codecommitConflictResolution); err != nil {
			log.Errorf("invalid --conflict-resolution: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}
	if len(_codecommitTargetBranch) > 0 {
		input.TargetBranch = aws.String(_codecommitTargetBranch)
	}

	if resp, err := client.MergeBranchesBySquash(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Merges two specified branches using the three-way merge strategy.
func codecommit_MergeBranchesByThreeWay(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.MergeBranchesByThreeWayInput{
		// DestinationCommitSpecifier: *string, // Required
		// RepositoryName: *string, // Required
		// SourceCommitSpecifier: *string, // Required
	}

	if len(_codecommitDestinationCommitSpecifier) > 0 {
		input.DestinationCommitSpecifier = aws.String(_codecommitDestinationCommitSpecifier)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitSpecifier) > 0 {
		input.SourceCommitSpecifier = aws.String(_codecommitSourceCommitSpecifier)
	}
	if len(_codecommitAuthorName) > 0 {
		input.AuthorName = aws.String(_codecommitAuthorName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolution) > 0 {
		if err := assignInputField(input, "ConflictResolution", _codecommitConflictResolution); err != nil {
			log.Errorf("invalid --conflict-resolution: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}
	if len(_codecommitTargetBranch) > 0 {
		input.TargetBranch = aws.String(_codecommitTargetBranch)
	}

	if resp, err := client.MergeBranchesByThreeWay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// fast-forward merge strategy. If the merge is successful, it closes the pull
// request.
func codecommit_MergePullRequestByFastForward(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.MergePullRequestByFastForwardInput{
		// PullRequestId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitSourceCommitId) > 0 {
		input.SourceCommitId = aws.String(_codecommitSourceCommitId)
	}

	if resp, err := client.MergePullRequestByFastForward(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// squash merge strategy. If the merge is successful, it closes the pull request.
func codecommit_MergePullRequestBySquash(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.MergePullRequestBySquashInput{
		// PullRequestId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitAuthorName) > 0 {
		input.AuthorName = aws.String(_codecommitAuthorName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolution) > 0 {
		if err := assignInputField(input, "ConflictResolution", _codecommitConflictResolution); err != nil {
			log.Errorf("invalid --conflict-resolution: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}
	if len(_codecommitSourceCommitId) > 0 {
		input.SourceCommitId = aws.String(_codecommitSourceCommitId)
	}

	if resp, err := client.MergePullRequestBySquash(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to merge the source commit of a pull request into the specified
// destination branch for that pull request at the specified commit using the
// three-way merge strategy. If the merge is successful, it closes the pull
// request.
func codecommit_MergePullRequestByThreeWay(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.MergePullRequestByThreeWayInput{
		// PullRequestId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitAuthorName) > 0 {
		input.AuthorName = aws.String(_codecommitAuthorName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitConflictDetailLevel) > 0 {
		if err := assignInputField(input, "ConflictDetailLevel", _codecommitConflictDetailLevel); err != nil {
			log.Errorf("invalid --conflict-detail-level: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolution) > 0 {
		if err := assignInputField(input, "ConflictResolution", _codecommitConflictResolution); err != nil {
			log.Errorf("invalid --conflict-resolution: %s", err.Error())
			return
		}
	}
	if len(_codecommitConflictResolutionStrategy) > 0 {
		if err := assignInputField(input, "ConflictResolutionStrategy", _codecommitConflictResolutionStrategy); err != nil {
			log.Errorf("invalid --conflict-resolution-strategy: %s", err.Error())
			return
		}
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitKeepEmptyFolders) > 0 {
		if err := assignInputField(input, "KeepEmptyFolders", _codecommitKeepEmptyFolders); err != nil {
			log.Errorf("invalid --keep-empty-folders: %s", err.Error())
			return
		}
	}
	if len(_codecommitSourceCommitId) > 0 {
		input.SourceCommitId = aws.String(_codecommitSourceCommitId)
	}

	if resp, err := client.MergePullRequestByThreeWay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets aside (overrides) all approval rule requirements for a specified pull
// request.
func codecommit_OverridePullRequestApprovalRules(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.OverridePullRequestApprovalRulesInput{
		// OverrideStatus: types.OverrideStatus, // Required
		// PullRequestId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_codecommitOverrideStatus) > 0 {
		if err := assignInputField(input, "OverrideStatus", _codecommitOverrideStatus); err != nil {
			log.Errorf("invalid --override-status: %s", err.Error())
			return
		}
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRevisionId) > 0 {
		input.RevisionId = aws.String(_codecommitRevisionId)
	}

	if resp, err := client.OverridePullRequestApprovalRules(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Posts a comment on the comparison between two commits.
func codecommit_PostCommentForComparedCommit(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.PostCommentForComparedCommitInput{
		// AfterCommitId: *string, // Required
		// Content: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitAfterCommitId) > 0 {
		input.AfterCommitId = aws.String(_codecommitAfterCommitId)
	}
	if len(_codecommitContent) > 0 {
		input.Content = aws.String(_codecommitContent)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitBeforeCommitId) > 0 {
		input.BeforeCommitId = aws.String(_codecommitBeforeCommitId)
	}
	if len(_codecommitClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codecommitClientRequestToken)
	}
	if len(_codecommitLocation) > 0 {
		if err := assignInputField(input, "Location", _codecommitLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}

	if resp, err := client.PostCommentForComparedCommit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Posts a comment on a pull request.
func codecommit_PostCommentForPullRequest(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.PostCommentForPullRequestInput{
		// AfterCommitId: *string, // Required
		// BeforeCommitId: *string, // Required
		// Content: *string, // Required
		// PullRequestId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitAfterCommitId) > 0 {
		input.AfterCommitId = aws.String(_codecommitAfterCommitId)
	}
	if len(_codecommitBeforeCommitId) > 0 {
		input.BeforeCommitId = aws.String(_codecommitBeforeCommitId)
	}
	if len(_codecommitContent) > 0 {
		input.Content = aws.String(_codecommitContent)
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codecommitClientRequestToken)
	}
	if len(_codecommitLocation) > 0 {
		if err := assignInputField(input, "Location", _codecommitLocation); err != nil {
			log.Errorf("invalid --location: %s", err.Error())
			return
		}
	}

	if resp, err := client.PostCommentForPullRequest(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Posts a comment in reply to an existing comment on a comparison between commits
// or a pull request.
func codecommit_PostCommentReply(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.PostCommentReplyInput{
		// Content: *string, // Required
		// InReplyTo: *string, // Required
	}

	if len(_codecommitContent) > 0 {
		input.Content = aws.String(_codecommitContent)
	}
	if len(_codecommitInReplyTo) > 0 {
		input.InReplyTo = aws.String(_codecommitInReplyTo)
	}
	if len(_codecommitClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_codecommitClientRequestToken)
	}

	if resp, err := client.PostCommentReply(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates a reaction to a specified comment for the user whose identity
// is used to make the request. You can only add or update a reaction for yourself.
// You cannot add, modify, or delete a reaction for another user.
func codecommit_PutCommentReaction(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.PutCommentReactionInput{
		// CommentId: *string, // Required
		// ReactionValue: *string, // Required
	}

	if len(_codecommitCommentId) > 0 {
		input.CommentId = aws.String(_codecommitCommentId)
	}
	if len(_codecommitReactionValue) > 0 {
		input.ReactionValue = aws.String(_codecommitReactionValue)
	}

	if resp, err := client.PutCommentReaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates a file in a branch in an CodeCommit repository, and generates a
// commit for the addition in the specified branch.
func codecommit_PutFile(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.PutFileInput{
		// BranchName: *string, // Required
		// FileContent: []byte, // Required
		// FilePath: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitBranchName) > 0 {
		input.BranchName = aws.String(_codecommitBranchName)
	}
	if len(_codecommitFileContent) > 0 {
		if err := assignInputField(input, "FileContent", _codecommitFileContent); err != nil {
			log.Errorf("invalid --file-content: %s", err.Error())
			return
		}
	}
	if len(_codecommitFilePath) > 0 {
		input.FilePath = aws.String(_codecommitFilePath)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitCommitMessage) > 0 {
		input.CommitMessage = aws.String(_codecommitCommitMessage)
	}
	if len(_codecommitEmail) > 0 {
		input.Email = aws.String(_codecommitEmail)
	}
	if len(_codecommitFileMode) > 0 {
		if err := assignInputField(input, "FileMode", _codecommitFileMode); err != nil {
			log.Errorf("invalid --file-mode: %s", err.Error())
			return
		}
	}
	if len(_codecommitName) > 0 {
		input.Name = aws.String(_codecommitName)
	}
	if len(_codecommitParentCommitId) > 0 {
		input.ParentCommitId = aws.String(_codecommitParentCommitId)
	}

	if resp, err := client.PutFile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces all triggers for a repository. Used to create or delete triggers.
func codecommit_PutRepositoryTriggers(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.PutRepositoryTriggersInput{
		// RepositoryName: *string, // Required
		// Triggers: []types.RepositoryTrigger, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitTriggers) > 0 {
		if err := assignInputField(input, "Triggers", _codecommitTriggers); err != nil {
			log.Errorf("invalid --triggers: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutRepositoryTriggers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a resource in CodeCommit. For a list of valid
// resources in CodeCommit, see [CodeCommit Resources and Operations]in the CodeCommit User Guide.
//
// [CodeCommit Resources and Operations]: https://docs.aws.amazon.com/codecommit/latest/userguide/auth-and-access-control-iam-access-control-identity-based.html#arn-formats
func codecommit_TagResource(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_codecommitResourceArn) > 0 {
		input.ResourceArn = aws.String(_codecommitResourceArn)
	}
	if len(_codecommitTags) > 0 {
		if err := assignInputField(input, "Tags", _codecommitTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tests the functionality of repository triggers by sending information to the
// trigger target. If real data is available in the repository, the test sends data
// from the last commit. If no data is available, sample data is generated.
func codecommit_TestRepositoryTriggers(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.TestRepositoryTriggersInput{
		// RepositoryName: *string, // Required
		// Triggers: []types.RepositoryTrigger, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitTriggers) > 0 {
		if err := assignInputField(input, "Triggers", _codecommitTriggers); err != nil {
			log.Errorf("invalid --triggers: %s", err.Error())
			return
		}
	}

	if resp, err := client.TestRepositoryTriggers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags for a resource in CodeCommit. For a list of valid resources in
// CodeCommit, see [CodeCommit Resources and Operations]in the CodeCommit User Guide.
//
// [CodeCommit Resources and Operations]: https://docs.aws.amazon.com/codecommit/latest/userguide/auth-and-access-control-iam-access-control-identity-based.html#arn-formats
func codecommit_UntagResource(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_codecommitResourceArn) > 0 {
		input.ResourceArn = aws.String(_codecommitResourceArn)
	}
	if len(_codecommitTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _codecommitTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the content of an approval rule template. You can change the number of
// required approvals, the membership of the approval rule, and whether an approval
// pool is defined.
func codecommit_UpdateApprovalRuleTemplateContent(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateApprovalRuleTemplateContentInput{
		// ApprovalRuleTemplateName: *string, // Required
		// NewRuleContent: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}
	if len(_codecommitNewRuleContent) > 0 {
		input.NewRuleContent = aws.String(_codecommitNewRuleContent)
	}
	if len(_codecommitExistingRuleContentSha256) > 0 {
		input.ExistingRuleContentSha256 = aws.String(_codecommitExistingRuleContentSha256)
	}

	if resp, err := client.UpdateApprovalRuleTemplateContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description for a specified approval rule template.
func codecommit_UpdateApprovalRuleTemplateDescription(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateApprovalRuleTemplateDescriptionInput{
		// ApprovalRuleTemplateDescription: *string, // Required
		// ApprovalRuleTemplateName: *string, // Required
	}

	if len(_codecommitApprovalRuleTemplateDescription) > 0 {
		input.ApprovalRuleTemplateDescription = aws.String(_codecommitApprovalRuleTemplateDescription)
	}
	if len(_codecommitApprovalRuleTemplateName) > 0 {
		input.ApprovalRuleTemplateName = aws.String(_codecommitApprovalRuleTemplateName)
	}

	if resp, err := client.UpdateApprovalRuleTemplateDescription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of a specified approval rule template.
func codecommit_UpdateApprovalRuleTemplateName(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateApprovalRuleTemplateNameInput{
		// NewApprovalRuleTemplateName: *string, // Required
		// OldApprovalRuleTemplateName: *string, // Required
	}

	if len(_codecommitNewApprovalRuleTemplateName) > 0 {
		input.NewApprovalRuleTemplateName = aws.String(_codecommitNewApprovalRuleTemplateName)
	}
	if len(_codecommitOldApprovalRuleTemplateName) > 0 {
		input.OldApprovalRuleTemplateName = aws.String(_codecommitOldApprovalRuleTemplateName)
	}

	if resp, err := client.UpdateApprovalRuleTemplateName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the contents of a comment.
func codecommit_UpdateComment(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateCommentInput{
		// CommentId: *string, // Required
		// Content: *string, // Required
	}

	if len(_codecommitCommentId) > 0 {
		input.CommentId = aws.String(_codecommitCommentId)
	}
	if len(_codecommitContent) > 0 {
		input.Content = aws.String(_codecommitContent)
	}

	if resp, err := client.UpdateComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets or changes the default branch name for the specified repository.
// If you use this operation to change the default branch name to the current
// default branch name, a success message is returned even though the default
// branch did not change.
func codecommit_UpdateDefaultBranch(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateDefaultBranchInput{
		// DefaultBranchName: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitDefaultBranchName) > 0 {
		input.DefaultBranchName = aws.String(_codecommitDefaultBranchName)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.UpdateDefaultBranch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the structure of an approval rule created specifically for a pull
// request. For example, you can change the number of required approvers and the
// approval pool for approvers.
func codecommit_UpdatePullRequestApprovalRuleContent(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdatePullRequestApprovalRuleContentInput{
		// ApprovalRuleName: *string, // Required
		// NewRuleContent: *string, // Required
		// PullRequestId: *string, // Required
	}

	if len(_codecommitApprovalRuleName) > 0 {
		input.ApprovalRuleName = aws.String(_codecommitApprovalRuleName)
	}
	if len(_codecommitNewRuleContent) > 0 {
		input.NewRuleContent = aws.String(_codecommitNewRuleContent)
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitExistingRuleContentSha256) > 0 {
		input.ExistingRuleContentSha256 = aws.String(_codecommitExistingRuleContentSha256)
	}

	if resp, err := client.UpdatePullRequestApprovalRuleContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the state of a user's approval on a pull request. The user is derived
// from the signed-in account when the request is made.
func codecommit_UpdatePullRequestApprovalState(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdatePullRequestApprovalStateInput{
		// ApprovalState: types.ApprovalState, // Required
		// PullRequestId: *string, // Required
		// RevisionId: *string, // Required
	}

	if len(_codecommitApprovalState) > 0 {
		if err := assignInputField(input, "ApprovalState", _codecommitApprovalState); err != nil {
			log.Errorf("invalid --approval-state: %s", err.Error())
			return
		}
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitRevisionId) > 0 {
		input.RevisionId = aws.String(_codecommitRevisionId)
	}

	if resp, err := client.UpdatePullRequestApprovalState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the contents of the description of a pull request.
func codecommit_UpdatePullRequestDescription(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdatePullRequestDescriptionInput{
		// Description: *string, // Required
		// PullRequestId: *string, // Required
	}

	if len(_codecommitDescription) > 0 {
		input.Description = aws.String(_codecommitDescription)
	}
	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}

	if resp, err := client.UpdatePullRequestDescription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status of a pull request.
func codecommit_UpdatePullRequestStatus(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdatePullRequestStatusInput{
		// PullRequestId: *string, // Required
		// PullRequestStatus: types.PullRequestStatusEnum, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitPullRequestStatus) > 0 {
		if err := assignInputField(input, "PullRequestStatus", _codecommitPullRequestStatus); err != nil {
			log.Errorf("invalid --pull-request-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePullRequestStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Replaces the title of a pull request.
func codecommit_UpdatePullRequestTitle(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdatePullRequestTitleInput{
		// PullRequestId: *string, // Required
		// Title: *string, // Required
	}

	if len(_codecommitPullRequestId) > 0 {
		input.PullRequestId = aws.String(_codecommitPullRequestId)
	}
	if len(_codecommitTitle) > 0 {
		input.Title = aws.String(_codecommitTitle)
	}

	if resp, err := client.UpdatePullRequestTitle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets or changes the comment or description for a repository.
// The description field for a repository accepts all HTML characters and all
// valid Unicode characters. Applications that do not HTML-encode the description
// and display it in a webpage can expose users to potentially malicious code. Make
// sure that you HTML-encode the description field in any application that uses
// this API to display the repository description on a webpage.
func codecommit_UpdateRepositoryDescription(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateRepositoryDescriptionInput{
		// RepositoryName: *string, // Required
	}

	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}
	if len(_codecommitRepositoryDescription) > 0 {
		input.RepositoryDescription = aws.String(_codecommitRepositoryDescription)
	}

	if resp, err := client.UpdateRepositoryDescription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Key Management Service encryption key used to encrypt and decrypt a
// CodeCommit repository.
func codecommit_UpdateRepositoryEncryptionKey(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateRepositoryEncryptionKeyInput{
		// KmsKeyId: *string, // Required
		// RepositoryName: *string, // Required
	}

	if len(_codecommitKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_codecommitKmsKeyId)
	}
	if len(_codecommitRepositoryName) > 0 {
		input.RepositoryName = aws.String(_codecommitRepositoryName)
	}

	if resp, err := client.UpdateRepositoryEncryptionKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renames a repository. The repository name must be unique across the calling
// Amazon Web Services account. Repository names are limited to 100 alphanumeric,
// dash, and underscore characters, and cannot include certain characters. The
// suffix .git is prohibited. For more information about the limits on repository
// names, see [Quotas]in the CodeCommit User Guide.
//
// [Quotas]: https://docs.aws.amazon.com/codecommit/latest/userguide/limits.html
func codecommit_UpdateRepositoryName(cfg aws.Config, client *codecommit.Client) {
	input := &codecommit.UpdateRepositoryNameInput{
		// NewName: *string, // Required
		// OldName: *string, // Required
	}

	if len(_codecommitNewName) > 0 {
		input.NewName = aws.String(_codecommitNewName)
	}
	if len(_codecommitOldName) > 0 {
		input.OldName = aws.String(_codecommitOldName)
	}

	if resp, err := client.UpdateRepositoryName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_codecommitCmd)
	_codecommitCmd.Flags().SortFlags = false

	_codecommitCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_codecommitCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_codecommitCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_codecommitCmd.Flags().StringVarP(&_codecommitActorArn, "actor-arn", "", "", "Actor ARN")
	_codecommitCmd.Flags().StringVarP(&_codecommitAfterCommitId, "after-commit-id", "", "", "After Commit ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitAfterCommitSpecifier, "after-commit-specifier", "", "", "After Commit Specifier")
	_codecommitCmd.Flags().StringVarP(&_codecommitAfterPath, "after-path", "", "", "After Path")
	_codecommitCmd.Flags().StringVarP(&_codecommitApprovalRuleContent, "approval-rule-content", "", "", "Approval Rule Content")
	_codecommitCmd.Flags().StringVarP(&_codecommitApprovalRuleName, "approval-rule-name", "", "", "Approval Rule Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitApprovalRuleTemplateContent, "approval-rule-template-content", "", "", "Approval Rule Template Content")
	_codecommitCmd.Flags().StringVarP(&_codecommitApprovalRuleTemplateDescription, "approval-rule-template-description", "", "", "Approval Rule Template Description")
	_codecommitCmd.Flags().StringVarP(&_codecommitApprovalRuleTemplateName, "approval-rule-template-name", "", "", "Approval Rule Template Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitApprovalState, "approval-state", "", "", "Approval State")
	_codecommitCmd.Flags().StringVarP(&_codecommitAuthorArn, "author-arn", "", "", "Author ARN")
	_codecommitCmd.Flags().StringVarP(&_codecommitAuthorName, "author-name", "", "", "Author Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitBeforeCommitId, "before-commit-id", "", "", "Before Commit ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitBeforeCommitSpecifier, "before-commit-specifier", "", "", "Before Commit Specifier")
	_codecommitCmd.Flags().StringVarP(&_codecommitBeforePath, "before-path", "", "", "Before Path")
	_codecommitCmd.Flags().StringVarP(&_codecommitBlobId, "blob-id", "", "", "Blob ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitBranchName, "branch-name", "", "", "Branch Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_codecommitCmd.Flags().StringVarP(&_codecommitCommentId, "comment-id", "", "", "Comment ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitCommitId, "commit-id", "", "", "Commit ID")
	_codecommitCmd.Flags().StringSliceVarP(&_codecommitCommitIds, "commit-ids", "", nil, "Commit Ids")
	_codecommitCmd.Flags().StringVarP(&_codecommitCommitMessage, "commit-message", "", "", "Commit Message")
	_codecommitCmd.Flags().StringVarP(&_codecommitCommitSpecifier, "commit-specifier", "", "", "Commit Specifier")
	_codecommitCmd.Flags().StringVarP(&_codecommitConflictDetailLevel, "conflict-detail-level", "", "", "Conflict Detail Level")
	_codecommitCmd.Flags().StringVarP(&_codecommitConflictResolution, "conflict-resolution", "", "", "Conflict Resolution")
	_codecommitCmd.Flags().StringVarP(&_codecommitConflictResolutionStrategy, "conflict-resolution-strategy", "", "", "Conflict Resolution Strategy")
	_codecommitCmd.Flags().StringVarP(&_codecommitContent, "content", "", "", "Content")
	_codecommitCmd.Flags().StringVarP(&_codecommitDefaultBranchName, "default-branch-name", "", "", "Default Branch Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitDeleteFiles, "delete-files", "", "", "Delete Files")
	_codecommitCmd.Flags().StringVarP(&_codecommitDescription, "description", "", "", "Description")
	_codecommitCmd.Flags().StringVarP(&_codecommitDestinationCommitSpecifier, "destination-commit-specifier", "", "", "Destination Commit Specifier")
	_codecommitCmd.Flags().StringVarP(&_codecommitEmail, "email", "", "", "Email")
	_codecommitCmd.Flags().StringVarP(&_codecommitExistingRuleContentSha256, "existing-rule-content-sha256", "", "", "Existing Rule Content SHA256")
	_codecommitCmd.Flags().StringVarP(&_codecommitFileContent, "file-content", "", "", "File Content")
	_codecommitCmd.Flags().StringVarP(&_codecommitFileMode, "file-mode", "", "", "File Mode")
	_codecommitCmd.Flags().StringVarP(&_codecommitFilePath, "file-path", "", "", "File Path")
	_codecommitCmd.Flags().StringSliceVarP(&_codecommitFilePaths, "file-paths", "", nil, "File Paths")
	_codecommitCmd.Flags().StringVarP(&_codecommitFolderPath, "folder-path", "", "", "Folder Path")
	_codecommitCmd.Flags().StringVarP(&_codecommitInReplyTo, "in-reply-to", "", "", "In Reply To")
	_codecommitCmd.Flags().StringVarP(&_codecommitKeepEmptyFolders, "keep-empty-folders", "", "", "Keep Empty Folders")
	_codecommitCmd.Flags().StringVarP(&_codecommitKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitLocation, "location", "", "", "Location")
	_codecommitCmd.Flags().StringVarP(&_codecommitMaxConflictFiles, "max-conflict-files", "", "", "Max Conflict Files")
	_codecommitCmd.Flags().StringVarP(&_codecommitMaxMergeHunks, "max-merge-hunks", "", "", "Max Merge Hunks")
	_codecommitCmd.Flags().StringVarP(&_codecommitMaxResults, "max-results", "", "", "Max Results")
	_codecommitCmd.Flags().StringVarP(&_codecommitMergeOption, "merge-option", "", "", "Merge Option")
	_codecommitCmd.Flags().StringVarP(&_codecommitName, "name", "", "", "Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitNewApprovalRuleTemplateName, "new-approval-rule-template-name", "", "", "New Approval Rule Template Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitNewName, "new-name", "", "", "New Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitNewRuleContent, "new-rule-content", "", "", "New Rule Content")
	_codecommitCmd.Flags().StringVarP(&_codecommitNextToken, "next-token", "", "", "Next Token")
	_codecommitCmd.Flags().StringVarP(&_codecommitOldApprovalRuleTemplateName, "old-approval-rule-template-name", "", "", "Old Approval Rule Template Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitOldName, "old-name", "", "", "Old Name")
	_codecommitCmd.Flags().StringVarP(&_codecommitOrder, "order", "", "", "Order")
	_codecommitCmd.Flags().StringVarP(&_codecommitOverrideStatus, "override-status", "", "", "Override Status")
	_codecommitCmd.Flags().StringVarP(&_codecommitParentCommitId, "parent-commit-id", "", "", "Parent Commit ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitPullRequestEventType, "pull-request-event-type", "", "", "Pull Request Event Type")
	_codecommitCmd.Flags().StringVarP(&_codecommitPullRequestId, "pull-request-id", "", "", "Pull Request ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitPullRequestStatus, "pull-request-status", "", "", "Pull Request Status")
	_codecommitCmd.Flags().StringVarP(&_codecommitPutFiles, "put-files", "", "", "Put Files")
	_codecommitCmd.Flags().StringVarP(&_codecommitReactionUserArn, "reaction-user-arn", "", "", "Reaction User ARN")
	_codecommitCmd.Flags().StringVarP(&_codecommitReactionValue, "reaction-value", "", "", "Reaction Value")
	_codecommitCmd.Flags().StringVarP(&_codecommitRepositoryDescription, "repository-description", "", "", "Repository Description")
	_codecommitCmd.Flags().StringVarP(&_codecommitRepositoryName, "repository-name", "", "", "Repository Name")
	_codecommitCmd.Flags().StringSliceVarP(&_codecommitRepositoryNames, "repository-names", "", nil, "Repository Names")
	_codecommitCmd.Flags().StringVarP(&_codecommitResourceArn, "resource-arn", "", "", "Resource ARN")
	_codecommitCmd.Flags().StringVarP(&_codecommitRevisionId, "revision-id", "", "", "Revision ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitSetFileModes, "set-file-modes", "", "", "Set File Modes")
	_codecommitCmd.Flags().StringVarP(&_codecommitSortBy, "sort-by", "", "", "Sort By")
	_codecommitCmd.Flags().StringVarP(&_codecommitSourceCommitId, "source-commit-id", "", "", "Source Commit ID")
	_codecommitCmd.Flags().StringVarP(&_codecommitSourceCommitSpecifier, "source-commit-specifier", "", "", "Source Commit Specifier")
	_codecommitCmd.Flags().StringSliceVarP(&_codecommitTagKeys, "tag-keys", "", nil, "Tag Keys")
	_codecommitCmd.Flags().StringVarP(&_codecommitTags, "tags", "", "", "Tags")
	_codecommitCmd.Flags().StringVarP(&_codecommitTargetBranch, "target-branch", "", "", "Target Branch")
	_codecommitCmd.Flags().StringVarP(&_codecommitTargets, "targets", "", "", "Targets")
	_codecommitCmd.Flags().StringVarP(&_codecommitTitle, "title", "", "", "Title")
	_codecommitCmd.Flags().StringVarP(&_codecommitTriggers, "triggers", "", "", "Triggers")

	_codecommitCmd.Flags().BoolVarP(&_codecommitAssociateApprovalRuleTemplateWithRepository, "associate-approval-rule-template-with-repository", "", false, "Associate Approval Rule Template With Repository")
	_codecommitCmd.Flags().BoolVarP(&_codecommitBatchAssociateApprovalRuleTemplateWithRepositories, "batch-associate-approval-rule-template-with-repositories", "", false, "Batch Associate Approval Rule Template With Repositories")
	_codecommitCmd.Flags().BoolVarP(&_codecommitBatchDescribeMergeConflicts, "batch-describe-merge-conflicts", "", false, "Batch Describe Merge Conflicts")
	_codecommitCmd.Flags().BoolVarP(&_codecommitBatchDisassociateApprovalRuleTemplateFromRepositories, "batch-disassociate-approval-rule-template-from-repositories", "", false, "Batch Disassociate Approval Rule Template From Repositories")
	_codecommitCmd.Flags().BoolVarP(&_codecommitBatchGetCommits, "batch-get-commits", "", false, "Batch Get Commits")
	_codecommitCmd.Flags().BoolVarP(&_codecommitBatchGetRepositories, "batch-get-repositories", "", false, "Batch Get Repositories")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreateApprovalRuleTemplate, "create-approval-rule-template", "", false, "Create Approval Rule Template")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreateBranch, "create-branch", "", false, "Create Branch")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreateCommit, "create-commit", "", false, "Create Commit")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreatePullRequest, "create-pull-request", "", false, "Create Pull Request")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreatePullRequestApprovalRule, "create-pull-request-approval-rule", "", false, "Create Pull Request Approval Rule")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreateRepository, "create-repository", "", false, "Create Repository")
	_codecommitCmd.Flags().BoolVarP(&_codecommitCreateUnreferencedMergeCommit, "create-unreferenced-merge-commit", "", false, "Create Unreferenced Merge Commit")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDeleteApprovalRuleTemplate, "delete-approval-rule-template", "", false, "Delete Approval Rule Template")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDeleteBranch, "delete-branch", "", false, "Delete Branch")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDeleteCommentContent, "delete-comment-content", "", false, "Delete Comment Content")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDeleteFile, "delete-file", "", false, "Delete File")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDeletePullRequestApprovalRule, "delete-pull-request-approval-rule", "", false, "Delete Pull Request Approval Rule")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDeleteRepository, "delete-repository", "", false, "Delete Repository")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDescribeMergeConflicts, "describe-merge-conflicts", "", false, "Describe Merge Conflicts")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDescribePullRequestEvents, "describe-pull-request-events", "", false, "Describe Pull Request Events")
	_codecommitCmd.Flags().BoolVarP(&_codecommitDisassociateApprovalRuleTemplateFromRepository, "disassociate-approval-rule-template-from-repository", "", false, "Disassociate Approval Rule Template From Repository")
	_codecommitCmd.Flags().BoolVarP(&_codecommitEvaluatePullRequestApprovalRules, "evaluate-pull-request-approval-rules", "", false, "Evaluate Pull Request Approval Rules")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetApprovalRuleTemplate, "get-approval-rule-template", "", false, "Get Approval Rule Template")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetBlob, "get-blob", "", false, "Get Blob")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetBranch, "get-branch", "", false, "Get Branch")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetComment, "get-comment", "", false, "Get Comment")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetCommentReactions, "get-comment-reactions", "", false, "Get Comment Reactions")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetCommentsForComparedCommit, "get-comments-for-compared-commit", "", false, "Get Comments For Compared Commit")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetCommentsForPullRequest, "get-comments-for-pull-request", "", false, "Get Comments For Pull Request")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetCommit, "get-commit", "", false, "Get Commit")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetDifferences, "get-differences", "", false, "Get Differences")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetFile, "get-file", "", false, "Get File")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetFolder, "get-folder", "", false, "Get Folder")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetMergeCommit, "get-merge-commit", "", false, "Get Merge Commit")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetMergeConflicts, "get-merge-conflicts", "", false, "Get Merge Conflicts")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetMergeOptions, "get-merge-options", "", false, "Get Merge Options")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetPullRequest, "get-pull-request", "", false, "Get Pull Request")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetPullRequestApprovalStates, "get-pull-request-approval-states", "", false, "Get Pull Request Approval States")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetPullRequestOverrideState, "get-pull-request-override-state", "", false, "Get Pull Request Override State")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetRepository, "get-repository", "", false, "Get Repository")
	_codecommitCmd.Flags().BoolVarP(&_codecommitGetRepositoryTriggers, "get-repository-triggers", "", false, "Get Repository Triggers")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListApprovalRuleTemplates, "list-approval-rule-templates", "", false, "List Approval Rule Templates")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListAssociatedApprovalRuleTemplatesForRepository, "list-associated-approval-rule-templates-for-repository", "", false, "List Associated Approval Rule Templates For Repository")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListBranches, "list-branches", "", false, "List Branches")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListFileCommitHistory, "list-file-commit-history", "", false, "List File Commit History")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListPullRequests, "list-pull-requests", "", false, "List Pull Requests")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListRepositories, "list-repositories", "", false, "List Repositories")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListRepositoriesForApprovalRuleTemplate, "list-repositories-for-approval-rule-template", "", false, "List Repositories For Approval Rule Template")
	_codecommitCmd.Flags().BoolVarP(&_codecommitListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_codecommitCmd.Flags().BoolVarP(&_codecommitMergeBranchesByFastForward, "merge-branches-by-fast-forward", "", false, "Merge Branches By Fast Forward")
	_codecommitCmd.Flags().BoolVarP(&_codecommitMergeBranchesBySquash, "merge-branches-by-squash", "", false, "Merge Branches By Squash")
	_codecommitCmd.Flags().BoolVarP(&_codecommitMergeBranchesByThreeWay, "merge-branches-by-three-way", "", false, "Merge Branches By Three Way")
	_codecommitCmd.Flags().BoolVarP(&_codecommitMergePullRequestByFastForward, "merge-pull-request-by-fast-forward", "", false, "Merge Pull Request By Fast Forward")
	_codecommitCmd.Flags().BoolVarP(&_codecommitMergePullRequestBySquash, "merge-pull-request-by-squash", "", false, "Merge Pull Request By Squash")
	_codecommitCmd.Flags().BoolVarP(&_codecommitMergePullRequestByThreeWay, "merge-pull-request-by-three-way", "", false, "Merge Pull Request By Three Way")
	_codecommitCmd.Flags().BoolVarP(&_codecommitOverridePullRequestApprovalRules, "override-pull-request-approval-rules", "", false, "Override Pull Request Approval Rules")
	_codecommitCmd.Flags().BoolVarP(&_codecommitPostCommentForComparedCommit, "post-comment-for-compared-commit", "", false, "Post Comment For Compared Commit")
	_codecommitCmd.Flags().BoolVarP(&_codecommitPostCommentForPullRequest, "post-comment-for-pull-request", "", false, "Post Comment For Pull Request")
	_codecommitCmd.Flags().BoolVarP(&_codecommitPostCommentReply, "post-comment-reply", "", false, "Post Comment Reply")
	_codecommitCmd.Flags().BoolVarP(&_codecommitPutCommentReaction, "put-comment-reaction", "", false, "Put Comment Reaction")
	_codecommitCmd.Flags().BoolVarP(&_codecommitPutFile, "put-file", "", false, "Put File")
	_codecommitCmd.Flags().BoolVarP(&_codecommitPutRepositoryTriggers, "put-repository-triggers", "", false, "Put Repository Triggers")
	_codecommitCmd.Flags().BoolVarP(&_codecommitTagResource, "tag-resource", "", false, "Tag Resource")
	_codecommitCmd.Flags().BoolVarP(&_codecommitTestRepositoryTriggers, "test-repository-triggers", "", false, "Test Repository Triggers")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUntagResource, "untag-resource", "", false, "Untag Resource")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateApprovalRuleTemplateContent, "update-approval-rule-template-content", "", false, "Update Approval Rule Template Content")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateApprovalRuleTemplateDescription, "update-approval-rule-template-description", "", false, "Update Approval Rule Template Description")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateApprovalRuleTemplateName, "update-approval-rule-template-name", "", false, "Update Approval Rule Template Name")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateComment, "update-comment", "", false, "Update Comment")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateDefaultBranch, "update-default-branch", "", false, "Update Default Branch")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdatePullRequestApprovalRuleContent, "update-pull-request-approval-rule-content", "", false, "Update Pull Request Approval Rule Content")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdatePullRequestApprovalState, "update-pull-request-approval-state", "", false, "Update Pull Request Approval State")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdatePullRequestDescription, "update-pull-request-description", "", false, "Update Pull Request Description")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdatePullRequestStatus, "update-pull-request-status", "", false, "Update Pull Request Status")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdatePullRequestTitle, "update-pull-request-title", "", false, "Update Pull Request Title")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateRepositoryDescription, "update-repository-description", "", false, "Update Repository Description")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateRepositoryEncryptionKey, "update-repository-encryption-key", "", false, "Update Repository Encryption Key")
	_codecommitCmd.Flags().BoolVarP(&_codecommitUpdateRepositoryName, "update-repository-name", "", false, "Update Repository Name")

}
