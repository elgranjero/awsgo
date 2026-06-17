package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// accessanalyzerCmd represents the accessanalyzer command
var _accessanalyzerCmd = &cobra.Command{
	Use:   "accessanalyzer",
	Short: "AWS accessanalyzer CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := accessanalyzer.NewFromConfig(cfg)
		if _accessanalyzerApplyArchiveRule {
			accessanalyzer_ApplyArchiveRule(cfg, client)
			return
		}
		if _accessanalyzerCancelPolicyGeneration {
			accessanalyzer_CancelPolicyGeneration(cfg, client)
			return
		}
		if _accessanalyzerCheckAccessNotGranted {
			accessanalyzer_CheckAccessNotGranted(cfg, client)
			return
		}
		if _accessanalyzerCheckNoNewAccess {
			accessanalyzer_CheckNoNewAccess(cfg, client)
			return
		}
		if _accessanalyzerCheckNoPublicAccess {
			accessanalyzer_CheckNoPublicAccess(cfg, client)
			return
		}
		if _accessanalyzerCreateAccessPreview {
			accessanalyzer_CreateAccessPreview(cfg, client)
			return
		}
		if _accessanalyzerCreateAnalyzer {
			accessanalyzer_CreateAnalyzer(cfg, client)
			return
		}
		if _accessanalyzerCreateArchiveRule {
			accessanalyzer_CreateArchiveRule(cfg, client)
			return
		}
		if _accessanalyzerDeleteAnalyzer {
			accessanalyzer_DeleteAnalyzer(cfg, client)
			return
		}
		if _accessanalyzerDeleteArchiveRule {
			accessanalyzer_DeleteArchiveRule(cfg, client)
			return
		}
		if _accessanalyzerGenerateFindingRecommendation {
			accessanalyzer_GenerateFindingRecommendation(cfg, client)
			return
		}
		if _accessanalyzerGetAccessPreview {
			accessanalyzer_GetAccessPreview(cfg, client)
			return
		}
		if _accessanalyzerGetAnalyzedResource {
			accessanalyzer_GetAnalyzedResource(cfg, client)
			return
		}
		if _accessanalyzerGetAnalyzer {
			accessanalyzer_GetAnalyzer(cfg, client)
			return
		}
		if _accessanalyzerGetArchiveRule {
			accessanalyzer_GetArchiveRule(cfg, client)
			return
		}
		if _accessanalyzerGetFinding {
			accessanalyzer_GetFinding(cfg, client)
			return
		}
		if _accessanalyzerGetFindingRecommendation {
			accessanalyzer_GetFindingRecommendation(cfg, client)
			return
		}
		if _accessanalyzerGetFindingV2 {
			accessanalyzer_GetFindingV2(cfg, client)
			return
		}
		if _accessanalyzerGetFindingsStatistics {
			accessanalyzer_GetFindingsStatistics(cfg, client)
			return
		}
		if _accessanalyzerGetGeneratedPolicy {
			accessanalyzer_GetGeneratedPolicy(cfg, client)
			return
		}
		if _accessanalyzerListAccessPreviewFindings {
			accessanalyzer_ListAccessPreviewFindings(cfg, client)
			return
		}
		if _accessanalyzerListAccessPreviews {
			accessanalyzer_ListAccessPreviews(cfg, client)
			return
		}
		if _accessanalyzerListAnalyzedResources {
			accessanalyzer_ListAnalyzedResources(cfg, client)
			return
		}
		if _accessanalyzerListAnalyzers {
			accessanalyzer_ListAnalyzers(cfg, client)
			return
		}
		if _accessanalyzerListArchiveRules {
			accessanalyzer_ListArchiveRules(cfg, client)
			return
		}
		if _accessanalyzerListFindings {
			accessanalyzer_ListFindings(cfg, client)
			return
		}
		if _accessanalyzerListFindingsV2 {
			accessanalyzer_ListFindingsV2(cfg, client)
			return
		}
		if _accessanalyzerListPolicyGenerations {
			accessanalyzer_ListPolicyGenerations(cfg, client)
			return
		}
		if _accessanalyzerListTagsForResource {
			accessanalyzer_ListTagsForResource(cfg, client)
			return
		}
		if _accessanalyzerStartPolicyGeneration {
			accessanalyzer_StartPolicyGeneration(cfg, client)
			return
		}
		if _accessanalyzerStartResourceScan {
			accessanalyzer_StartResourceScan(cfg, client)
			return
		}
		if _accessanalyzerTagResource {
			accessanalyzer_TagResource(cfg, client)
			return
		}
		if _accessanalyzerUntagResource {
			accessanalyzer_UntagResource(cfg, client)
			return
		}
		if _accessanalyzerUpdateAnalyzer {
			accessanalyzer_UpdateAnalyzer(cfg, client)
			return
		}
		if _accessanalyzerUpdateArchiveRule {
			accessanalyzer_UpdateArchiveRule(cfg, client)
			return
		}
		if _accessanalyzerUpdateFindings {
			accessanalyzer_UpdateFindings(cfg, client)
			return
		}
		if _accessanalyzerValidatePolicy {
			accessanalyzer_ValidatePolicy(cfg, client)
			return
		}

	},
}

var (
	_accessanalyzerApplyArchiveRule              bool
	_accessanalyzerCancelPolicyGeneration        bool
	_accessanalyzerCheckAccessNotGranted         bool
	_accessanalyzerCheckNoNewAccess              bool
	_accessanalyzerCheckNoPublicAccess           bool
	_accessanalyzerCreateAccessPreview           bool
	_accessanalyzerCreateAnalyzer                bool
	_accessanalyzerCreateArchiveRule             bool
	_accessanalyzerDeleteAnalyzer                bool
	_accessanalyzerDeleteArchiveRule             bool
	_accessanalyzerGenerateFindingRecommendation bool
	_accessanalyzerGetAccessPreview              bool
	_accessanalyzerGetAnalyzedResource           bool
	_accessanalyzerGetAnalyzer                   bool
	_accessanalyzerGetArchiveRule                bool
	_accessanalyzerGetFinding                    bool
	_accessanalyzerGetFindingRecommendation      bool
	_accessanalyzerGetFindingV2                  bool
	_accessanalyzerGetFindingsStatistics         bool
	_accessanalyzerGetGeneratedPolicy            bool
	_accessanalyzerListAccessPreviewFindings     bool
	_accessanalyzerListAccessPreviews            bool
	_accessanalyzerListAnalyzedResources         bool
	_accessanalyzerListAnalyzers                 bool
	_accessanalyzerListArchiveRules              bool
	_accessanalyzerListFindings                  bool
	_accessanalyzerListFindingsV2                bool
	_accessanalyzerListPolicyGenerations         bool
	_accessanalyzerListTagsForResource           bool
	_accessanalyzerStartPolicyGeneration         bool
	_accessanalyzerStartResourceScan             bool
	_accessanalyzerTagResource                   bool
	_accessanalyzerUntagResource                 bool
	_accessanalyzerUpdateAnalyzer                bool
	_accessanalyzerUpdateArchiveRule             bool
	_accessanalyzerUpdateFindings                bool
	_accessanalyzerValidatePolicy                bool

	_accessanalyzerAccess                      string
	_accessanalyzerAccessPreviewId             string
	_accessanalyzerAnalyzerArn                 string
	_accessanalyzerAnalyzerName                string
	_accessanalyzerArchiveRules                string
	_accessanalyzerClientToken                 string
	_accessanalyzerCloudTrailDetails           string
	_accessanalyzerConfiguration               string
	_accessanalyzerConfigurations              string
	_accessanalyzerExistingPolicyDocument      string
	_accessanalyzerFilter                      string
	_accessanalyzerId                          string
	_accessanalyzerIds                         []string
	_accessanalyzerIncludeResourcePlaceholders string
	_accessanalyzerIncludeServiceLevelTemplate string
	_accessanalyzerJobId                       string
	_accessanalyzerLocale                      string
	_accessanalyzerMaxResults                  string
	_accessanalyzerNewPolicyDocument           string
	_accessanalyzerNextToken                   string
	_accessanalyzerPolicyDocument              string
	_accessanalyzerPolicyGenerationDetails     string
	_accessanalyzerPolicyType                  string
	_accessanalyzerPrincipalArn                string
	_accessanalyzerResourceArn                 string
	_accessanalyzerResourceOwnerAccount        string
	_accessanalyzerResourceType                string
	_accessanalyzerRuleName                    string
	_accessanalyzerSort                        string
	_accessanalyzerStatus                      string
	_accessanalyzerTagKeys                     []string
	_accessanalyzerTags                        string
	_accessanalyzerType                        string
	_accessanalyzerValidatePolicyResourceType  string
)

// Retroactively applies the archive rule to existing findings that meet the
// archive rule criteria.
func accessanalyzer_ApplyArchiveRule(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ApplyArchiveRuleInput{
		// AnalyzerArn: *string, // Required
		// RuleName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerRuleName) > 0 {
		input.RuleName = aws.String(_accessanalyzerRuleName)
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}

	if resp, err := client.ApplyArchiveRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the requested policy generation.
func accessanalyzer_CancelPolicyGeneration(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CancelPolicyGenerationInput{
		// JobId: *string, // Required
	}

	if len(_accessanalyzerJobId) > 0 {
		input.JobId = aws.String(_accessanalyzerJobId)
	}

	if resp, err := client.CancelPolicyGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks whether the specified access isn't allowed by a policy.
func accessanalyzer_CheckAccessNotGranted(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CheckAccessNotGrantedInput{
		// Access: []types.Access, // Required
		// PolicyDocument: *string, // Required
		// PolicyType: types.AccessCheckPolicyType, // Required
	}

	if len(_accessanalyzerAccess) > 0 {
		if err := assignInputField(input, "Access", _accessanalyzerAccess); err != nil {
			log.Errorf("invalid --access: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_accessanalyzerPolicyDocument)
	}
	if len(_accessanalyzerPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _accessanalyzerPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CheckAccessNotGranted(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks whether new access is allowed for an updated policy when compared to the
// existing policy.
//
// You can find examples for reference policies and learn how to set up and run a
// custom policy check for new access in the [IAM Access Analyzer custom policy checks samples]repository on GitHub. The reference
// policies in this repository are meant to be passed to the existingPolicyDocument
// request parameter.
//
// [IAM Access Analyzer custom policy checks samples]: https://github.com/aws-samples/iam-access-analyzer-custom-policy-check-samples
func accessanalyzer_CheckNoNewAccess(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CheckNoNewAccessInput{
		// ExistingPolicyDocument: *string, // Required
		// NewPolicyDocument: *string, // Required
		// PolicyType: types.AccessCheckPolicyType, // Required
	}

	if len(_accessanalyzerExistingPolicyDocument) > 0 {
		input.ExistingPolicyDocument = aws.String(_accessanalyzerExistingPolicyDocument)
	}
	if len(_accessanalyzerNewPolicyDocument) > 0 {
		input.NewPolicyDocument = aws.String(_accessanalyzerNewPolicyDocument)
	}
	if len(_accessanalyzerPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _accessanalyzerPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CheckNoNewAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks whether a resource policy can grant public access to the specified
// resource type.
func accessanalyzer_CheckNoPublicAccess(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CheckNoPublicAccessInput{
		// PolicyDocument: *string, // Required
		// ResourceType: types.AccessCheckResourceType, // Required
	}

	if len(_accessanalyzerPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_accessanalyzerPolicyDocument)
	}
	if len(_accessanalyzerResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _accessanalyzerResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.CheckNoPublicAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access preview that allows you to preview IAM Access Analyzer
// findings for your resource before deploying resource permissions.
func accessanalyzer_CreateAccessPreview(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CreateAccessPreviewInput{
		// AnalyzerArn: *string, // Required
		// Configurations: map[string]types.Configuration, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerConfigurations) > 0 {
		if err := assignInputField(input, "Configurations", _accessanalyzerConfigurations); err != nil {
			log.Errorf("invalid --configurations: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}

	if resp, err := client.CreateAccessPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an analyzer for your account.
func accessanalyzer_CreateAnalyzer(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CreateAnalyzerInput{
		// AnalyzerName: *string, // Required
		// Type: types.Type, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerType) > 0 {
		if err := assignInputField(input, "Type", _accessanalyzerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerArchiveRules) > 0 {
		if err := assignInputField(input, "ArchiveRules", _accessanalyzerArchiveRules); err != nil {
			log.Errorf("invalid --archive-rules: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}
	if len(_accessanalyzerConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _accessanalyzerConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerTags) > 0 {
		if err := assignInputField(input, "Tags", _accessanalyzerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAnalyzer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an archive rule for the specified analyzer. Archive rules automatically
// archive new findings that meet the criteria you define when you create the rule.
//
// To learn about filter keys that you can use to create an archive rule, see [IAM Access Analyzer filter keys] in
// the IAM User Guide.
//
// [IAM Access Analyzer filter keys]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-filter-keys.html
func accessanalyzer_CreateArchiveRule(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.CreateArchiveRuleInput{
		// AnalyzerName: *string, // Required
		// Filter: map[string]types.Criterion, // Required
		// RuleName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerFilter) > 0 {
		if err := assignInputField(input, "Filter", _accessanalyzerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerRuleName) > 0 {
		input.RuleName = aws.String(_accessanalyzerRuleName)
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}

	if resp, err := client.CreateArchiveRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified analyzer. When you delete an analyzer, IAM Access
// Analyzer is disabled for the account or organization in the current or specific
// Region. All findings that were generated by the analyzer are deleted. You cannot
// undo this action.
func accessanalyzer_DeleteAnalyzer(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.DeleteAnalyzerInput{
		// AnalyzerName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}

	if resp, err := client.DeleteAnalyzer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified archive rule.
func accessanalyzer_DeleteArchiveRule(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.DeleteArchiveRuleInput{
		// AnalyzerName: *string, // Required
		// RuleName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerRuleName) > 0 {
		input.RuleName = aws.String(_accessanalyzerRuleName)
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}

	if resp, err := client.DeleteArchiveRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a recommendation for an unused permissions finding.
func accessanalyzer_GenerateFindingRecommendation(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GenerateFindingRecommendationInput{
		// AnalyzerArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerId) > 0 {
		input.Id = aws.String(_accessanalyzerId)
	}

	if resp, err := client.GenerateFindingRecommendation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an access preview for the specified analyzer.
func accessanalyzer_GetAccessPreview(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetAccessPreviewInput{
		// AccessPreviewId: *string, // Required
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAccessPreviewId) > 0 {
		input.AccessPreviewId = aws.String(_accessanalyzerAccessPreviewId)
	}
	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}

	if resp, err := client.GetAccessPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a resource that was analyzed.
// This action is supported only for external access analyzers.
func accessanalyzer_GetAnalyzedResource(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetAnalyzedResourceInput{
		// AnalyzerArn: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerResourceArn) > 0 {
		input.ResourceArn = aws.String(_accessanalyzerResourceArn)
	}

	if resp, err := client.GetAnalyzedResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified analyzer.
func accessanalyzer_GetAnalyzer(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetAnalyzerInput{
		// AnalyzerName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}

	if resp, err := client.GetAnalyzer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an archive rule.
// To learn about filter keys that you can use to create an archive rule, see [IAM Access Analyzer filter keys] in
// the IAM User Guide.
//
// [IAM Access Analyzer filter keys]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-filter-keys.html
func accessanalyzer_GetArchiveRule(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetArchiveRuleInput{
		// AnalyzerName: *string, // Required
		// RuleName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerRuleName) > 0 {
		input.RuleName = aws.String(_accessanalyzerRuleName)
	}

	if resp, err := client.GetArchiveRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about the specified finding. GetFinding and GetFindingV2
// both use access-analyzer:GetFinding in the Action element of an IAM policy
// statement. You must have permission to perform the access-analyzer:GetFinding
// action.
//
// GetFinding is supported only for external access analyzers. You must use
// GetFindingV2 for internal and unused access analyzers.
func accessanalyzer_GetFinding(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetFindingInput{
		// AnalyzerArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerId) > 0 {
		input.Id = aws.String(_accessanalyzerId)
	}

	if resp, err := client.GetFinding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a finding recommendation for the specified analyzer.
func accessanalyzer_GetFindingRecommendation(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetFindingRecommendationInput{
		// AnalyzerArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerId) > 0 {
		input.Id = aws.String(_accessanalyzerId)
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFindingRecommendation(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.GetFindingRecommendationOutput
	p := accessanalyzer.NewGetFindingRecommendationPaginator(client, input)
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

// Retrieves information about the specified finding. GetFinding and GetFindingV2
// both use access-analyzer:GetFinding in the Action element of an IAM policy
// statement. You must have permission to perform the access-analyzer:GetFinding
// action.
func accessanalyzer_GetFindingV2(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetFindingV2Input{
		// AnalyzerArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerId) > 0 {
		input.Id = aws.String(_accessanalyzerId)
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFindingV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.GetFindingV2Output
	p := accessanalyzer.NewGetFindingV2Paginator(client, input)
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

// Retrieves a list of aggregated finding statistics for an external access or
// unused access analyzer.
func accessanalyzer_GetFindingsStatistics(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetFindingsStatisticsInput{
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}

	if resp, err := client.GetFindingsStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the policy that was generated using StartPolicyGeneration .
func accessanalyzer_GetGeneratedPolicy(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.GetGeneratedPolicyInput{
		// JobId: *string, // Required
	}

	if len(_accessanalyzerJobId) > 0 {
		input.JobId = aws.String(_accessanalyzerJobId)
	}
	if len(_accessanalyzerIncludeResourcePlaceholders) > 0 {
		if err := assignInputField(input, "IncludeResourcePlaceholders", _accessanalyzerIncludeResourcePlaceholders); err != nil {
			log.Errorf("invalid --include-resource-placeholders: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerIncludeServiceLevelTemplate) > 0 {
		if err := assignInputField(input, "IncludeServiceLevelTemplate", _accessanalyzerIncludeServiceLevelTemplate); err != nil {
			log.Errorf("invalid --include-service-level-template: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetGeneratedPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves a list of access preview findings generated by the specified access
// preview.
func accessanalyzer_ListAccessPreviewFindings(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListAccessPreviewFindingsInput{
		// AccessPreviewId: *string, // Required
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAccessPreviewId) > 0 {
		input.AccessPreviewId = aws.String(_accessanalyzerAccessPreviewId)
	}
	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerFilter) > 0 {
		if err := assignInputField(input, "Filter", _accessanalyzerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPreviewFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListAccessPreviewFindingsOutput
	p := accessanalyzer.NewListAccessPreviewFindingsPaginator(client, input)
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

// Retrieves a list of access previews for the specified analyzer.
func accessanalyzer_ListAccessPreviews(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListAccessPreviewsInput{
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPreviews(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListAccessPreviewsOutput
	p := accessanalyzer.NewListAccessPreviewsPaginator(client, input)
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

// Retrieves a list of resources of the specified type that have been analyzed by
// the specified analyzer.
func accessanalyzer_ListAnalyzedResources(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListAnalyzedResourcesInput{
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}
	if len(_accessanalyzerResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _accessanalyzerResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAnalyzedResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListAnalyzedResourcesOutput
	p := accessanalyzer.NewListAnalyzedResourcesPaginator(client, input)
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

// Retrieves a list of analyzers.
func accessanalyzer_ListAnalyzers(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListAnalyzersInput{}

	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}
	if len(_accessanalyzerType) > 0 {
		if err := assignInputField(input, "Type", _accessanalyzerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAnalyzers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListAnalyzersOutput
	p := accessanalyzer.NewListAnalyzersPaginator(client, input)
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

// Retrieves a list of archive rules created for the specified analyzer.
func accessanalyzer_ListArchiveRules(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListArchiveRulesInput{
		// AnalyzerName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListArchiveRules(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListArchiveRulesOutput
	p := accessanalyzer.NewListArchiveRulesPaginator(client, input)
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

// Retrieves a list of findings generated by the specified analyzer. ListFindings
// and ListFindingsV2 both use access-analyzer:ListFindings in the Action element
// of an IAM policy statement. You must have permission to perform the
// access-analyzer:ListFindings action.
//
// To learn about filter keys that you can use to retrieve a list of findings, see [IAM Access Analyzer filter keys]
// in the IAM User Guide.
//
// ListFindings is supported only for external access analyzers. You must use
// ListFindingsV2 for internal and unused access analyzers.
//
// [IAM Access Analyzer filter keys]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-filter-keys.html
func accessanalyzer_ListFindings(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListFindingsInput{
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerFilter) > 0 {
		if err := assignInputField(input, "Filter", _accessanalyzerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}
	if len(_accessanalyzerSort) > 0 {
		if err := assignInputField(input, "Sort", _accessanalyzerSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFindings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListFindingsOutput
	p := accessanalyzer.NewListFindingsPaginator(client, input)
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

// Retrieves a list of findings generated by the specified analyzer. ListFindings
// and ListFindingsV2 both use access-analyzer:ListFindings in the Action element
// of an IAM policy statement. You must have permission to perform the
// access-analyzer:ListFindings action.
//
// To learn about filter keys that you can use to retrieve a list of findings, see [IAM Access Analyzer filter keys]
// in the IAM User Guide.
//
// [IAM Access Analyzer filter keys]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-reference-filter-keys.html
func accessanalyzer_ListFindingsV2(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListFindingsV2Input{
		// AnalyzerArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerFilter) > 0 {
		if err := assignInputField(input, "Filter", _accessanalyzerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}
	if len(_accessanalyzerSort) > 0 {
		if err := assignInputField(input, "Sort", _accessanalyzerSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListFindingsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListFindingsV2Output
	p := accessanalyzer.NewListFindingsV2Paginator(client, input)
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

// Lists all of the policy generations requested in the last seven days.
func accessanalyzer_ListPolicyGenerations(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListPolicyGenerationsInput{}

	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}
	if len(_accessanalyzerPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_accessanalyzerPrincipalArn)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyGenerations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ListPolicyGenerationsOutput
	p := accessanalyzer.NewListPolicyGenerationsPaginator(client, input)
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

// Retrieves a list of tags applied to the specified resource.
func accessanalyzer_ListTagsForResource(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_accessanalyzerResourceArn) > 0 {
		input.ResourceArn = aws.String(_accessanalyzerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the policy generation request.
func accessanalyzer_StartPolicyGeneration(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.StartPolicyGenerationInput{
		// PolicyGenerationDetails: *types.PolicyGenerationDetails, // Required
	}

	if len(_accessanalyzerPolicyGenerationDetails) > 0 {
		if err := assignInputField(input, "PolicyGenerationDetails", _accessanalyzerPolicyGenerationDetails); err != nil {
			log.Errorf("invalid --policy-generation-details: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}
	if len(_accessanalyzerCloudTrailDetails) > 0 {
		if err := assignInputField(input, "CloudTrailDetails", _accessanalyzerCloudTrailDetails); err != nil {
			log.Errorf("invalid --cloud-trail-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartPolicyGeneration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Immediately starts a scan of the policies applied to the specified resource.
// This action is supported only for external access analyzers.
func accessanalyzer_StartResourceScan(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.StartResourceScanInput{
		// AnalyzerArn: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerResourceArn) > 0 {
		input.ResourceArn = aws.String(_accessanalyzerResourceArn)
	}
	if len(_accessanalyzerResourceOwnerAccount) > 0 {
		input.ResourceOwnerAccount = aws.String(_accessanalyzerResourceOwnerAccount)
	}

	if resp, err := client.StartResourceScan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag to the specified resource.
func accessanalyzer_TagResource(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_accessanalyzerResourceArn) > 0 {
		input.ResourceArn = aws.String(_accessanalyzerResourceArn)
	}
	if len(_accessanalyzerTags) > 0 {
		if err := assignInputField(input, "Tags", _accessanalyzerTags); err != nil {
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

// Removes a tag from the specified resource.
func accessanalyzer_UntagResource(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_accessanalyzerResourceArn) > 0 {
		input.ResourceArn = aws.String(_accessanalyzerResourceArn)
	}
	if len(_accessanalyzerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _accessanalyzerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the configuration of an existing analyzer.
// This action is not supported for external access analyzers.
func accessanalyzer_UpdateAnalyzer(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.UpdateAnalyzerInput{
		// AnalyzerName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _accessanalyzerConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAnalyzer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the criteria and values for the specified archive rule.
func accessanalyzer_UpdateArchiveRule(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.UpdateArchiveRuleInput{
		// AnalyzerName: *string, // Required
		// Filter: map[string]types.Criterion, // Required
		// RuleName: *string, // Required
	}

	if len(_accessanalyzerAnalyzerName) > 0 {
		input.AnalyzerName = aws.String(_accessanalyzerAnalyzerName)
	}
	if len(_accessanalyzerFilter) > 0 {
		if err := assignInputField(input, "Filter", _accessanalyzerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerRuleName) > 0 {
		input.RuleName = aws.String(_accessanalyzerRuleName)
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}

	if resp, err := client.UpdateArchiveRule(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status for the specified findings.
func accessanalyzer_UpdateFindings(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.UpdateFindingsInput{
		// AnalyzerArn: *string, // Required
		// Status: types.FindingStatusUpdate, // Required
	}

	if len(_accessanalyzerAnalyzerArn) > 0 {
		input.AnalyzerArn = aws.String(_accessanalyzerAnalyzerArn)
	}
	if len(_accessanalyzerStatus) > 0 {
		if err := assignInputField(input, "Status", _accessanalyzerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerClientToken) > 0 {
		input.ClientToken = aws.String(_accessanalyzerClientToken)
	}
	if len(_accessanalyzerIds) > 0 {
		input.Ids = append([]string(nil), _accessanalyzerIds...)
	}
	if len(_accessanalyzerResourceArn) > 0 {
		input.ResourceArn = aws.String(_accessanalyzerResourceArn)
	}

	if resp, err := client.UpdateFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests the validation of a policy and returns a list of findings. The
// findings help you identify issues and provide actionable recommendations to
// resolve the issue and enable you to author functional policies that meet
// security best practices.
func accessanalyzer_ValidatePolicy(cfg aws.Config, client *accessanalyzer.Client) {
	input := &accessanalyzer.ValidatePolicyInput{
		// PolicyDocument: *string, // Required
		// PolicyType: types.PolicyType, // Required
	}

	if len(_accessanalyzerPolicyDocument) > 0 {
		input.PolicyDocument = aws.String(_accessanalyzerPolicyDocument)
	}
	if len(_accessanalyzerPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _accessanalyzerPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerLocale) > 0 {
		if err := assignInputField(input, "Locale", _accessanalyzerLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _accessanalyzerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_accessanalyzerNextToken) > 0 {
		input.NextToken = aws.String(_accessanalyzerNextToken)
	}
	if len(_accessanalyzerValidatePolicyResourceType) > 0 {
		if err := assignInputField(input, "ValidatePolicyResourceType", _accessanalyzerValidatePolicyResourceType); err != nil {
			log.Errorf("invalid --validate-policy-resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ValidatePolicy(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*accessanalyzer.ValidatePolicyOutput
	p := accessanalyzer.NewValidatePolicyPaginator(client, input)
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

func init() {
	_rootCmd.AddCommand(_accessanalyzerCmd)
	_accessanalyzerCmd.Flags().SortFlags = false

	_accessanalyzerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_accessanalyzerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_accessanalyzerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerAccess, "access", "", "", "Access")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerAccessPreviewId, "access-preview-id", "", "", "Access Preview ID")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerAnalyzerArn, "analyzer-arn", "", "", "Analyzer ARN")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerAnalyzerName, "analyzer-name", "", "", "Analyzer Name")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerArchiveRules, "archive-rules", "", "", "Archive Rules")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerClientToken, "client-token", "", "", "Client Token")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerCloudTrailDetails, "cloud-trail-details", "", "", "Cloud Trail Details")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerConfiguration, "configuration", "", "", "Configuration")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerConfigurations, "configurations", "", "", "Configurations")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerExistingPolicyDocument, "existing-policy-document", "", "", "Existing Policy Document")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerFilter, "filter", "", "", "Filter")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerId, "id", "", "", "ID")
	_accessanalyzerCmd.Flags().StringSliceVarP(&_accessanalyzerIds, "ids", "", nil, "Ids")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerIncludeResourcePlaceholders, "include-resource-placeholders", "", "", "Include Resource Placeholders")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerIncludeServiceLevelTemplate, "include-service-level-template", "", "", "Include Service Level Template")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerJobId, "job-id", "", "", "Job ID")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerLocale, "locale", "", "", "Locale")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerMaxResults, "max-results", "", "", "Max Results")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerNewPolicyDocument, "new-policy-document", "", "", "New Policy Document")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerNextToken, "next-token", "", "", "Next Token")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerPolicyDocument, "policy-document", "", "", "Policy Document")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerPolicyGenerationDetails, "policy-generation-details", "", "", "Policy Generation Details")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerPolicyType, "policy-type", "", "", "Policy Type")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerPrincipalArn, "principal-arn", "", "", "Principal ARN")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerResourceArn, "resource-arn", "", "", "Resource ARN")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerResourceOwnerAccount, "resource-owner-account", "", "", "Resource Owner Account")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerResourceType, "resource-type", "", "", "Resource Type")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerRuleName, "rule-name", "", "", "Rule Name")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerSort, "sort", "", "", "Sort")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerStatus, "status", "", "", "Status")
	_accessanalyzerCmd.Flags().StringSliceVarP(&_accessanalyzerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerTags, "tags", "", "", "Tags")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerType, "type", "", "", "Type")
	_accessanalyzerCmd.Flags().StringVarP(&_accessanalyzerValidatePolicyResourceType, "validate-policy-resource-type", "", "", "Validate Policy Resource Type")

	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerApplyArchiveRule, "apply-archive-rule", "", false, "Apply Archive Rule")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCancelPolicyGeneration, "cancel-policy-generation", "", false, "Cancel Policy Generation")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCheckAccessNotGranted, "check-access-not-granted", "", false, "Check Access Not Granted")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCheckNoNewAccess, "check-no-new-access", "", false, "Check No New Access")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCheckNoPublicAccess, "check-no-public-access", "", false, "Check No Public Access")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCreateAccessPreview, "create-access-preview", "", false, "Create Access Preview")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCreateAnalyzer, "create-analyzer", "", false, "Create Analyzer")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerCreateArchiveRule, "create-archive-rule", "", false, "Create Archive Rule")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerDeleteAnalyzer, "delete-analyzer", "", false, "Delete Analyzer")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerDeleteArchiveRule, "delete-archive-rule", "", false, "Delete Archive Rule")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGenerateFindingRecommendation, "generate-finding-recommendation", "", false, "Generate Finding Recommendation")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetAccessPreview, "get-access-preview", "", false, "Get Access Preview")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetAnalyzedResource, "get-analyzed-resource", "", false, "Get Analyzed Resource")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetAnalyzer, "get-analyzer", "", false, "Get Analyzer")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetArchiveRule, "get-archive-rule", "", false, "Get Archive Rule")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetFinding, "get-finding", "", false, "Get Finding")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetFindingRecommendation, "get-finding-recommendation", "", false, "Get Finding Recommendation")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetFindingV2, "get-finding-v2", "", false, "Get Finding V2")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetFindingsStatistics, "get-findings-statistics", "", false, "Get Findings Statistics")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerGetGeneratedPolicy, "get-generated-policy", "", false, "Get Generated Policy")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListAccessPreviewFindings, "list-access-preview-findings", "", false, "List Access Preview Findings")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListAccessPreviews, "list-access-previews", "", false, "List Access Previews")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListAnalyzedResources, "list-analyzed-resources", "", false, "List Analyzed Resources")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListAnalyzers, "list-analyzers", "", false, "List Analyzers")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListArchiveRules, "list-archive-rules", "", false, "List Archive Rules")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListFindings, "list-findings", "", false, "List Findings")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListFindingsV2, "list-findings-v2", "", false, "List Findings V2")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListPolicyGenerations, "list-policy-generations", "", false, "List Policy Generations")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerStartPolicyGeneration, "start-policy-generation", "", false, "Start Policy Generation")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerStartResourceScan, "start-resource-scan", "", false, "Start Resource Scan")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerTagResource, "tag-resource", "", false, "Tag Resource")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerUntagResource, "untag-resource", "", false, "Untag Resource")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerUpdateAnalyzer, "update-analyzer", "", false, "Update Analyzer")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerUpdateArchiveRule, "update-archive-rule", "", false, "Update Archive Rule")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerUpdateFindings, "update-findings", "", false, "Update Findings")
	_accessanalyzerCmd.Flags().BoolVarP(&_accessanalyzerValidatePolicy, "validate-policy", "", false, "Validate Policy")

}
