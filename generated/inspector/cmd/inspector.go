package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// inspectorCmd represents the inspector command
var _inspectorCmd = &cobra.Command{
	Use:   "inspector",
	Short: "AWS inspector CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := inspector.NewFromConfig(cfg)
		if _inspectorAddAttributesToFindings {
			inspector_AddAttributesToFindings(cfg, client)
			return
		}
		if _inspectorCreateAssessmentTarget {
			inspector_CreateAssessmentTarget(cfg, client)
			return
		}
		if _inspectorCreateAssessmentTemplate {
			inspector_CreateAssessmentTemplate(cfg, client)
			return
		}
		if _inspectorCreateExclusionsPreview {
			inspector_CreateExclusionsPreview(cfg, client)
			return
		}
		if _inspectorCreateResourceGroup {
			inspector_CreateResourceGroup(cfg, client)
			return
		}
		if _inspectorDeleteAssessmentRun {
			inspector_DeleteAssessmentRun(cfg, client)
			return
		}
		if _inspectorDeleteAssessmentTarget {
			inspector_DeleteAssessmentTarget(cfg, client)
			return
		}
		if _inspectorDeleteAssessmentTemplate {
			inspector_DeleteAssessmentTemplate(cfg, client)
			return
		}
		if _inspectorDescribeAssessmentRuns {
			inspector_DescribeAssessmentRuns(cfg, client)
			return
		}
		if _inspectorDescribeAssessmentTargets {
			inspector_DescribeAssessmentTargets(cfg, client)
			return
		}
		if _inspectorDescribeAssessmentTemplates {
			inspector_DescribeAssessmentTemplates(cfg, client)
			return
		}
		if _inspectorDescribeCrossAccountAccessRole {
			inspector_DescribeCrossAccountAccessRole(cfg, client)
			return
		}
		if _inspectorDescribeExclusions {
			inspector_DescribeExclusions(cfg, client)
			return
		}
		if _inspectorDescribeFindings {
			inspector_DescribeFindings(cfg, client)
			return
		}
		if _inspectorDescribeResourceGroups {
			inspector_DescribeResourceGroups(cfg, client)
			return
		}
		if _inspectorDescribeRulesPackages {
			inspector_DescribeRulesPackages(cfg, client)
			return
		}
		if _inspectorGetAssessmentReport {
			inspector_GetAssessmentReport(cfg, client)
			return
		}
		if _inspectorGetExclusionsPreview {
			inspector_GetExclusionsPreview(cfg, client)
			return
		}
		if _inspectorGetTelemetryMetadata {
			inspector_GetTelemetryMetadata(cfg, client)
			return
		}
		if _inspectorListAssessmentRunAgents {
			inspector_ListAssessmentRunAgents(cfg, client)
			return
		}
		if _inspectorListAssessmentRuns {
			inspector_ListAssessmentRuns(cfg, client)
			return
		}
		if _inspectorListAssessmentTargets {
			inspector_ListAssessmentTargets(cfg, client)
			return
		}
		if _inspectorListAssessmentTemplates {
			inspector_ListAssessmentTemplates(cfg, client)
			return
		}
		if _inspectorListEventSubscriptions {
			inspector_ListEventSubscriptions(cfg, client)
			return
		}
		if _inspectorListExclusions {
			inspector_ListExclusions(cfg, client)
			return
		}
		if _inspectorListFindings {
			inspector_ListFindings(cfg, client)
			return
		}
		if _inspectorListRulesPackages {
			inspector_ListRulesPackages(cfg, client)
			return
		}
		if _inspectorListTagsForResource {
			inspector_ListTagsForResource(cfg, client)
			return
		}
		if _inspectorPreviewAgents {
			inspector_PreviewAgents(cfg, client)
			return
		}
		if _inspectorRegisterCrossAccountAccessRole {
			inspector_RegisterCrossAccountAccessRole(cfg, client)
			return
		}
		if _inspectorRemoveAttributesFromFindings {
			inspector_RemoveAttributesFromFindings(cfg, client)
			return
		}
		if _inspectorSetTagsForResource {
			inspector_SetTagsForResource(cfg, client)
			return
		}
		if _inspectorStartAssessmentRun {
			inspector_StartAssessmentRun(cfg, client)
			return
		}
		if _inspectorStopAssessmentRun {
			inspector_StopAssessmentRun(cfg, client)
			return
		}
		if _inspectorSubscribeToEvent {
			inspector_SubscribeToEvent(cfg, client)
			return
		}
		if _inspectorUnsubscribeFromEvent {
			inspector_UnsubscribeFromEvent(cfg, client)
			return
		}
		if _inspectorUpdateAssessmentTarget {
			inspector_UpdateAssessmentTarget(cfg, client)
			return
		}

	},
}

var (
	_inspectorAddAttributesToFindings        bool
	_inspectorCreateAssessmentTarget         bool
	_inspectorCreateAssessmentTemplate       bool
	_inspectorCreateExclusionsPreview        bool
	_inspectorCreateResourceGroup            bool
	_inspectorDeleteAssessmentRun            bool
	_inspectorDeleteAssessmentTarget         bool
	_inspectorDeleteAssessmentTemplate       bool
	_inspectorDescribeAssessmentRuns         bool
	_inspectorDescribeAssessmentTargets      bool
	_inspectorDescribeAssessmentTemplates    bool
	_inspectorDescribeCrossAccountAccessRole bool
	_inspectorDescribeExclusions             bool
	_inspectorDescribeFindings               bool
	_inspectorDescribeResourceGroups         bool
	_inspectorDescribeRulesPackages          bool
	_inspectorGetAssessmentReport            bool
	_inspectorGetExclusionsPreview           bool
	_inspectorGetTelemetryMetadata           bool
	_inspectorListAssessmentRunAgents        bool
	_inspectorListAssessmentRuns             bool
	_inspectorListAssessmentTargets          bool
	_inspectorListAssessmentTemplates        bool
	_inspectorListEventSubscriptions         bool
	_inspectorListExclusions                 bool
	_inspectorListFindings                   bool
	_inspectorListRulesPackages              bool
	_inspectorListTagsForResource            bool
	_inspectorPreviewAgents                  bool
	_inspectorRegisterCrossAccountAccessRole bool
	_inspectorRemoveAttributesFromFindings   bool
	_inspectorSetTagsForResource             bool
	_inspectorStartAssessmentRun             bool
	_inspectorStopAssessmentRun              bool
	_inspectorSubscribeToEvent               bool
	_inspectorUnsubscribeFromEvent           bool
	_inspectorUpdateAssessmentTarget         bool

	_inspectorAssessmentRunArn          string
	_inspectorAssessmentRunArns         []string
	_inspectorAssessmentRunName         string
	_inspectorAssessmentTargetArn       string
	_inspectorAssessmentTargetArns      []string
	_inspectorAssessmentTargetName      string
	_inspectorAssessmentTemplateArn     string
	_inspectorAssessmentTemplateArns    []string
	_inspectorAssessmentTemplateName    string
	_inspectorAttributeKeys             []string
	_inspectorAttributes                string
	_inspectorDurationInSeconds         string
	_inspectorEvent                     string
	_inspectorExclusionArns             []string
	_inspectorFilter                    string
	_inspectorFindingArns               []string
	_inspectorLocale                    string
	_inspectorMaxResults                string
	_inspectorNextToken                 string
	_inspectorPreviewAgentsArn          string
	_inspectorPreviewToken              string
	_inspectorReportFileFormat          string
	_inspectorReportType                string
	_inspectorResourceArn               string
	_inspectorResourceGroupArn          string
	_inspectorResourceGroupArns         []string
	_inspectorResourceGroupTags         string
	_inspectorRoleArn                   string
	_inspectorRulesPackageArns          []string
	_inspectorStopAction                string
	_inspectorTags                      string
	_inspectorTopicArn                  string
	_inspectorUserAttributesForFindings string
)

// Assigns attributes (key and value pairs) to the findings that are specified by
// the ARNs of the findings.
func inspector_AddAttributesToFindings(cfg aws.Config, client *inspector.Client) {
	input := &inspector.AddAttributesToFindingsInput{
		// Attributes: []types.Attribute, // Required
		// FindingArns: []string, // Required
	}

	if len(_inspectorAttributes) > 0 {
		if err := assignInputField(input, "Attributes", _inspectorAttributes); err != nil {
			log.Errorf("invalid --attributes: %s", err.Error())
			return
		}
	}
	if len(_inspectorFindingArns) > 0 {
		input.FindingArns = append([]string(nil), _inspectorFindingArns...)
	}

	if resp, err := client.AddAttributesToFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new assessment target using the ARN of the resource group that is
// generated by CreateResourceGroup. If resourceGroupArn is not specified, all EC2 instances in the
// current AWS account and region are included in the assessment target. If the [service-linked role]
// isn’t already registered, this action also creates and registers a
// service-linked role to grant Amazon Inspector access to AWS Services needed to
// perform security assessments. You can create up to 50 assessment targets per AWS
// account. You can run up to 500 concurrent agents per AWS account. For more
// information, see [Amazon Inspector Assessment Targets].
//
// [Amazon Inspector Assessment Targets]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_applications.html
// [service-linked role]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html
func inspector_CreateAssessmentTarget(cfg aws.Config, client *inspector.Client) {
	input := &inspector.CreateAssessmentTargetInput{
		// AssessmentTargetName: *string, // Required
	}

	if len(_inspectorAssessmentTargetName) > 0 {
		input.AssessmentTargetName = aws.String(_inspectorAssessmentTargetName)
	}
	if len(_inspectorResourceGroupArn) > 0 {
		input.ResourceGroupArn = aws.String(_inspectorResourceGroupArn)
	}

	if resp, err := client.CreateAssessmentTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an assessment template for the assessment target that is specified by
// the ARN of the assessment target. If the [service-linked role]isn’t already registered, this action
// also creates and registers a service-linked role to grant Amazon Inspector
// access to AWS Services needed to perform security assessments.
//
// [service-linked role]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_slr.html
func inspector_CreateAssessmentTemplate(cfg aws.Config, client *inspector.Client) {
	input := &inspector.CreateAssessmentTemplateInput{
		// AssessmentTargetArn: *string, // Required
		// AssessmentTemplateName: *string, // Required
		// DurationInSeconds: *int32, // Required
		// RulesPackageArns: []string, // Required
	}

	if len(_inspectorAssessmentTargetArn) > 0 {
		input.AssessmentTargetArn = aws.String(_inspectorAssessmentTargetArn)
	}
	if len(_inspectorAssessmentTemplateName) > 0 {
		input.AssessmentTemplateName = aws.String(_inspectorAssessmentTemplateName)
	}
	if len(_inspectorDurationInSeconds) > 0 {
		if err := assignInputField(input, "DurationInSeconds", _inspectorDurationInSeconds); err != nil {
			log.Errorf("invalid --duration-in-seconds: %s", err.Error())
			return
		}
	}
	if len(_inspectorRulesPackageArns) > 0 {
		input.RulesPackageArns = append([]string(nil), _inspectorRulesPackageArns...)
	}
	if len(_inspectorUserAttributesForFindings) > 0 {
		if err := assignInputField(input, "UserAttributesForFindings", _inspectorUserAttributesForFindings); err != nil {
			log.Errorf("invalid --user-attributes-for-findings: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAssessmentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the generation of an exclusions preview for the specified assessment
// template. The exclusions preview lists the potential exclusions
// (ExclusionPreview) that Inspector can detect before it runs the assessment.
func inspector_CreateExclusionsPreview(cfg aws.Config, client *inspector.Client) {
	input := &inspector.CreateExclusionsPreviewInput{
		// AssessmentTemplateArn: *string, // Required
	}

	if len(_inspectorAssessmentTemplateArn) > 0 {
		input.AssessmentTemplateArn = aws.String(_inspectorAssessmentTemplateArn)
	}

	if resp, err := client.CreateExclusionsPreview(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource group using the specified set of tags (key and value pairs)
// that are used to select the EC2 instances to be included in an Amazon Inspector
// assessment target. The created resource group is then used to create an Amazon
// Inspector assessment target. For more information, see CreateAssessmentTarget.
func inspector_CreateResourceGroup(cfg aws.Config, client *inspector.Client) {
	input := &inspector.CreateResourceGroupInput{
		// ResourceGroupTags: []types.ResourceGroupTag, // Required
	}

	if len(_inspectorResourceGroupTags) > 0 {
		if err := assignInputField(input, "ResourceGroupTags", _inspectorResourceGroupTags); err != nil {
			log.Errorf("invalid --resource-group-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the assessment run that is specified by the ARN of the assessment run.
func inspector_DeleteAssessmentRun(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DeleteAssessmentRunInput{
		// AssessmentRunArn: *string, // Required
	}

	if len(_inspectorAssessmentRunArn) > 0 {
		input.AssessmentRunArn = aws.String(_inspectorAssessmentRunArn)
	}

	if resp, err := client.DeleteAssessmentRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the assessment target that is specified by the ARN of the assessment
// target.
func inspector_DeleteAssessmentTarget(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DeleteAssessmentTargetInput{
		// AssessmentTargetArn: *string, // Required
	}

	if len(_inspectorAssessmentTargetArn) > 0 {
		input.AssessmentTargetArn = aws.String(_inspectorAssessmentTargetArn)
	}

	if resp, err := client.DeleteAssessmentTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the assessment template that is specified by the ARN of the assessment
// template.
func inspector_DeleteAssessmentTemplate(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DeleteAssessmentTemplateInput{
		// AssessmentTemplateArn: *string, // Required
	}

	if len(_inspectorAssessmentTemplateArn) > 0 {
		input.AssessmentTemplateArn = aws.String(_inspectorAssessmentTemplateArn)
	}

	if resp, err := client.DeleteAssessmentTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the assessment runs that are specified by the ARNs of the assessment
// runs.
func inspector_DescribeAssessmentRuns(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeAssessmentRunsInput{
		// AssessmentRunArns: []string, // Required
	}

	if len(_inspectorAssessmentRunArns) > 0 {
		input.AssessmentRunArns = append([]string(nil), _inspectorAssessmentRunArns...)
	}

	if resp, err := client.DescribeAssessmentRuns(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the assessment targets that are specified by the ARNs of the
// assessment targets.
func inspector_DescribeAssessmentTargets(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeAssessmentTargetsInput{
		// AssessmentTargetArns: []string, // Required
	}

	if len(_inspectorAssessmentTargetArns) > 0 {
		input.AssessmentTargetArns = append([]string(nil), _inspectorAssessmentTargetArns...)
	}

	if resp, err := client.DescribeAssessmentTargets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the assessment templates that are specified by the ARNs of the
// assessment templates.
func inspector_DescribeAssessmentTemplates(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeAssessmentTemplatesInput{
		// AssessmentTemplateArns: []string, // Required
	}

	if len(_inspectorAssessmentTemplateArns) > 0 {
		input.AssessmentTemplateArns = append([]string(nil), _inspectorAssessmentTemplateArns...)
	}

	if resp, err := client.DescribeAssessmentTemplates(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the IAM role that enables Amazon Inspector to access your AWS account.
func inspector_DescribeCrossAccountAccessRole(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeCrossAccountAccessRoleInput{}

	if resp, err := client.DescribeCrossAccountAccessRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the exclusions that are specified by the exclusions' ARNs.
func inspector_DescribeExclusions(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeExclusionsInput{
		// ExclusionArns: []string, // Required
	}

	if len(_inspectorExclusionArns) > 0 {
		input.ExclusionArns = append([]string(nil), _inspectorExclusionArns...)
	}
	if len(_inspectorLocale) > 0 {
		if err := assignInputField(input, "Locale", _inspectorLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeExclusions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the findings that are specified by the ARNs of the findings.
func inspector_DescribeFindings(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeFindingsInput{
		// FindingArns: []string, // Required
	}

	if len(_inspectorFindingArns) > 0 {
		input.FindingArns = append([]string(nil), _inspectorFindingArns...)
	}
	if len(_inspectorLocale) > 0 {
		if err := assignInputField(input, "Locale", _inspectorLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the resource groups that are specified by the ARNs of the resource
// groups.
func inspector_DescribeResourceGroups(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeResourceGroupsInput{
		// ResourceGroupArns: []string, // Required
	}

	if len(_inspectorResourceGroupArns) > 0 {
		input.ResourceGroupArns = append([]string(nil), _inspectorResourceGroupArns...)
	}

	if resp, err := client.DescribeResourceGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the rules packages that are specified by the ARNs of the rules
// packages.
func inspector_DescribeRulesPackages(cfg aws.Config, client *inspector.Client) {
	input := &inspector.DescribeRulesPackagesInput{
		// RulesPackageArns: []string, // Required
	}

	if len(_inspectorRulesPackageArns) > 0 {
		input.RulesPackageArns = append([]string(nil), _inspectorRulesPackageArns...)
	}
	if len(_inspectorLocale) > 0 {
		if err := assignInputField(input, "Locale", _inspectorLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeRulesPackages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Produces an assessment report that includes detailed and comprehensive results
// of a specified assessment run.
func inspector_GetAssessmentReport(cfg aws.Config, client *inspector.Client) {
	input := &inspector.GetAssessmentReportInput{
		// AssessmentRunArn: *string, // Required
		// ReportFileFormat: types.ReportFileFormat, // Required
		// ReportType: types.ReportType, // Required
	}

	if len(_inspectorAssessmentRunArn) > 0 {
		input.AssessmentRunArn = aws.String(_inspectorAssessmentRunArn)
	}
	if len(_inspectorReportFileFormat) > 0 {
		if err := assignInputField(input, "ReportFileFormat", _inspectorReportFileFormat); err != nil {
			log.Errorf("invalid --report-file-format: %s", err.Error())
			return
		}
	}
	if len(_inspectorReportType) > 0 {
		if err := assignInputField(input, "ReportType", _inspectorReportType); err != nil {
			log.Errorf("invalid --report-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAssessmentReport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the exclusions preview (a list of ExclusionPreview objects) specified
// by the preview token. You can obtain the preview token by running the
// CreateExclusionsPreview API.
func inspector_GetExclusionsPreview(cfg aws.Config, client *inspector.Client) {
	input := &inspector.GetExclusionsPreviewInput{
		// AssessmentTemplateArn: *string, // Required
		// PreviewToken: *string, // Required
	}

	if len(_inspectorAssessmentTemplateArn) > 0 {
		input.AssessmentTemplateArn = aws.String(_inspectorAssessmentTemplateArn)
	}
	if len(_inspectorPreviewToken) > 0 {
		input.PreviewToken = aws.String(_inspectorPreviewToken)
	}
	if len(_inspectorLocale) > 0 {
		if err := assignInputField(input, "Locale", _inspectorLocale); err != nil {
			log.Errorf("invalid --locale: %s", err.Error())
			return
		}
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetExclusionsPreview(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.GetExclusionsPreviewOutput
	p := inspector.NewGetExclusionsPreviewPaginator(client, input)
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

// Information about the data that is collected for the specified assessment run.
func inspector_GetTelemetryMetadata(cfg aws.Config, client *inspector.Client) {
	input := &inspector.GetTelemetryMetadataInput{
		// AssessmentRunArn: *string, // Required
	}

	if len(_inspectorAssessmentRunArn) > 0 {
		input.AssessmentRunArn = aws.String(_inspectorAssessmentRunArn)
	}

	if resp, err := client.GetTelemetryMetadata(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the agents of the assessment runs that are specified by the ARNs of the
// assessment runs.
func inspector_ListAssessmentRunAgents(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListAssessmentRunAgentsInput{
		// AssessmentRunArn: *string, // Required
	}

	if len(_inspectorAssessmentRunArn) > 0 {
		input.AssessmentRunArn = aws.String(_inspectorAssessmentRunArn)
	}
	if len(_inspectorFilter) > 0 {
		if err := assignInputField(input, "Filter", _inspectorFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentRunAgents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListAssessmentRunAgentsOutput
	p := inspector.NewListAssessmentRunAgentsPaginator(client, input)
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

// Lists the assessment runs that correspond to the assessment templates that are
// specified by the ARNs of the assessment templates.
func inspector_ListAssessmentRuns(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListAssessmentRunsInput{}

	if len(_inspectorAssessmentTemplateArns) > 0 {
		input.AssessmentTemplateArns = append([]string(nil), _inspectorAssessmentTemplateArns...)
	}
	if len(_inspectorFilter) > 0 {
		if err := assignInputField(input, "Filter", _inspectorFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentRuns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListAssessmentRunsOutput
	p := inspector.NewListAssessmentRunsPaginator(client, input)
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

// Lists the ARNs of the assessment targets within this AWS account. For more
// information about assessment targets, see [Amazon Inspector Assessment Targets].
//
// [Amazon Inspector Assessment Targets]: https://docs.aws.amazon.com/inspector/latest/userguide/inspector_applications.html
func inspector_ListAssessmentTargets(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListAssessmentTargetsInput{}

	if len(_inspectorFilter) > 0 {
		if err := assignInputField(input, "Filter", _inspectorFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListAssessmentTargetsOutput
	p := inspector.NewListAssessmentTargetsPaginator(client, input)
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

// Lists the assessment templates that correspond to the assessment targets that
// are specified by the ARNs of the assessment targets.
func inspector_ListAssessmentTemplates(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListAssessmentTemplatesInput{}

	if len(_inspectorAssessmentTargetArns) > 0 {
		input.AssessmentTargetArns = append([]string(nil), _inspectorAssessmentTargetArns...)
	}
	if len(_inspectorFilter) > 0 {
		if err := assignInputField(input, "Filter", _inspectorFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAssessmentTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListAssessmentTemplatesOutput
	p := inspector.NewListAssessmentTemplatesPaginator(client, input)
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

// Lists all the event subscriptions for the assessment template that is specified
// by the ARN of the assessment template. For more information, see SubscribeToEventand UnsubscribeFromEvent.
func inspector_ListEventSubscriptions(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListEventSubscriptionsInput{}

	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}
	if len(_inspectorResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspectorResourceArn)
	}

	if disablePaginator() {
		if resp, err := client.ListEventSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListEventSubscriptionsOutput
	p := inspector.NewListEventSubscriptionsPaginator(client, input)
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

// List exclusions that are generated by the assessment run.
func inspector_ListExclusions(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListExclusionsInput{
		// AssessmentRunArn: *string, // Required
	}

	if len(_inspectorAssessmentRunArn) > 0 {
		input.AssessmentRunArn = aws.String(_inspectorAssessmentRunArn)
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListExclusions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListExclusionsOutput
	p := inspector.NewListExclusionsPaginator(client, input)
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

// Lists findings that are generated by the assessment runs that are specified by
// the ARNs of the assessment runs.
func inspector_ListFindings(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListFindingsInput{}

	if len(_inspectorAssessmentRunArns) > 0 {
		input.AssessmentRunArns = append([]string(nil), _inspectorAssessmentRunArns...)
	}
	if len(_inspectorFilter) > 0 {
		if err := assignInputField(input, "Filter", _inspectorFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
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

	var results []*inspector.ListFindingsOutput
	p := inspector.NewListFindingsPaginator(client, input)
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

// Lists all available Amazon Inspector rules packages.
func inspector_ListRulesPackages(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListRulesPackagesInput{}

	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRulesPackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.ListRulesPackagesOutput
	p := inspector.NewListRulesPackagesPaginator(client, input)
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

// Lists all tags associated with an assessment template.
func inspector_ListTagsForResource(cfg aws.Config, client *inspector.Client) {
	input := &inspector.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_inspectorResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspectorResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Previews the agents installed on the EC2 instances that are part of the
// specified assessment target.
func inspector_PreviewAgents(cfg aws.Config, client *inspector.Client) {
	input := &inspector.PreviewAgentsInput{
		// PreviewAgentsArn: *string, // Required
	}

	if len(_inspectorPreviewAgentsArn) > 0 {
		input.PreviewAgentsArn = aws.String(_inspectorPreviewAgentsArn)
	}
	if len(_inspectorMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _inspectorMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_inspectorNextToken) > 0 {
		input.NextToken = aws.String(_inspectorNextToken)
	}

	if disablePaginator() {
		if resp, err := client.PreviewAgents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*inspector.PreviewAgentsOutput
	p := inspector.NewPreviewAgentsPaginator(client, input)
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

// Registers the IAM role that grants Amazon Inspector access to AWS Services
// needed to perform security assessments.
func inspector_RegisterCrossAccountAccessRole(cfg aws.Config, client *inspector.Client) {
	input := &inspector.RegisterCrossAccountAccessRoleInput{
		// RoleArn: *string, // Required
	}

	if len(_inspectorRoleArn) > 0 {
		input.RoleArn = aws.String(_inspectorRoleArn)
	}

	if resp, err := client.RegisterCrossAccountAccessRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes entire attributes (key and value pairs) from the findings that are
// specified by the ARNs of the findings where an attribute with the specified key
// exists.
func inspector_RemoveAttributesFromFindings(cfg aws.Config, client *inspector.Client) {
	input := &inspector.RemoveAttributesFromFindingsInput{
		// AttributeKeys: []string, // Required
		// FindingArns: []string, // Required
	}

	if len(_inspectorAttributeKeys) > 0 {
		input.AttributeKeys = append([]string(nil), _inspectorAttributeKeys...)
	}
	if len(_inspectorFindingArns) > 0 {
		input.FindingArns = append([]string(nil), _inspectorFindingArns...)
	}

	if resp, err := client.RemoveAttributesFromFindings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets tags (key and value pairs) to the assessment template that is specified by
// the ARN of the assessment template.
func inspector_SetTagsForResource(cfg aws.Config, client *inspector.Client) {
	input := &inspector.SetTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_inspectorResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspectorResourceArn)
	}
	if len(_inspectorTags) > 0 {
		if err := assignInputField(input, "Tags", _inspectorTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.SetTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the assessment run specified by the ARN of the assessment template. For
// this API to function properly, you must not exceed the limit of running up to
// 500 concurrent agents per AWS account.
func inspector_StartAssessmentRun(cfg aws.Config, client *inspector.Client) {
	input := &inspector.StartAssessmentRunInput{
		// AssessmentTemplateArn: *string, // Required
	}

	if len(_inspectorAssessmentTemplateArn) > 0 {
		input.AssessmentTemplateArn = aws.String(_inspectorAssessmentTemplateArn)
	}
	if len(_inspectorAssessmentRunName) > 0 {
		input.AssessmentRunName = aws.String(_inspectorAssessmentRunName)
	}

	if resp, err := client.StartAssessmentRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the assessment run that is specified by the ARN of the assessment run.
func inspector_StopAssessmentRun(cfg aws.Config, client *inspector.Client) {
	input := &inspector.StopAssessmentRunInput{
		// AssessmentRunArn: *string, // Required
	}

	if len(_inspectorAssessmentRunArn) > 0 {
		input.AssessmentRunArn = aws.String(_inspectorAssessmentRunArn)
	}
	if len(_inspectorStopAction) > 0 {
		if err := assignInputField(input, "StopAction", _inspectorStopAction); err != nil {
			log.Errorf("invalid --stop-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.StopAssessmentRun(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables the process of sending Amazon Simple Notification Service (SNS)
// notifications about a specified event to a specified SNS topic.
func inspector_SubscribeToEvent(cfg aws.Config, client *inspector.Client) {
	input := &inspector.SubscribeToEventInput{
		// Event: types.InspectorEvent, // Required
		// ResourceArn: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_inspectorEvent) > 0 {
		if err := assignInputField(input, "Event", _inspectorEvent); err != nil {
			log.Errorf("invalid --event: %s", err.Error())
			return
		}
	}
	if len(_inspectorResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspectorResourceArn)
	}
	if len(_inspectorTopicArn) > 0 {
		input.TopicArn = aws.String(_inspectorTopicArn)
	}

	if resp, err := client.SubscribeToEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the process of sending Amazon Simple Notification Service (SNS)
// notifications about a specified event to a specified SNS topic.
func inspector_UnsubscribeFromEvent(cfg aws.Config, client *inspector.Client) {
	input := &inspector.UnsubscribeFromEventInput{
		// Event: types.InspectorEvent, // Required
		// ResourceArn: *string, // Required
		// TopicArn: *string, // Required
	}

	if len(_inspectorEvent) > 0 {
		if err := assignInputField(input, "Event", _inspectorEvent); err != nil {
			log.Errorf("invalid --event: %s", err.Error())
			return
		}
	}
	if len(_inspectorResourceArn) > 0 {
		input.ResourceArn = aws.String(_inspectorResourceArn)
	}
	if len(_inspectorTopicArn) > 0 {
		input.TopicArn = aws.String(_inspectorTopicArn)
	}

	if resp, err := client.UnsubscribeFromEvent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the assessment target that is specified by the ARN of the assessment
// target.
//
// If resourceGroupArn is not specified, all EC2 instances in the current AWS
// account and region are included in the assessment target.
func inspector_UpdateAssessmentTarget(cfg aws.Config, client *inspector.Client) {
	input := &inspector.UpdateAssessmentTargetInput{
		// AssessmentTargetArn: *string, // Required
		// AssessmentTargetName: *string, // Required
	}

	if len(_inspectorAssessmentTargetArn) > 0 {
		input.AssessmentTargetArn = aws.String(_inspectorAssessmentTargetArn)
	}
	if len(_inspectorAssessmentTargetName) > 0 {
		input.AssessmentTargetName = aws.String(_inspectorAssessmentTargetName)
	}
	if len(_inspectorResourceGroupArn) > 0 {
		input.ResourceGroupArn = aws.String(_inspectorResourceGroupArn)
	}

	if resp, err := client.UpdateAssessmentTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_inspectorCmd)
	_inspectorCmd.Flags().SortFlags = false

	_inspectorCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_inspectorCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_inspectorCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_inspectorCmd.Flags().StringVarP(&_inspectorAssessmentRunArn, "assessment-run-arn", "", "", "Assessment Run ARN")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorAssessmentRunArns, "assessment-run-arns", "", nil, "Assessment Run Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorAssessmentRunName, "assessment-run-name", "", "", "Assessment Run Name")
	_inspectorCmd.Flags().StringVarP(&_inspectorAssessmentTargetArn, "assessment-target-arn", "", "", "Assessment Target ARN")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorAssessmentTargetArns, "assessment-target-arns", "", nil, "Assessment Target Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorAssessmentTargetName, "assessment-target-name", "", "", "Assessment Target Name")
	_inspectorCmd.Flags().StringVarP(&_inspectorAssessmentTemplateArn, "assessment-template-arn", "", "", "Assessment Template ARN")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorAssessmentTemplateArns, "assessment-template-arns", "", nil, "Assessment Template Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorAssessmentTemplateName, "assessment-template-name", "", "", "Assessment Template Name")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorAttributeKeys, "attribute-keys", "", nil, "Attribute Keys")
	_inspectorCmd.Flags().StringVarP(&_inspectorAttributes, "attributes", "", "", "Attributes")
	_inspectorCmd.Flags().StringVarP(&_inspectorDurationInSeconds, "duration-in-seconds", "", "", "Duration In Seconds")
	_inspectorCmd.Flags().StringVarP(&_inspectorEvent, "event", "", "", "Event")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorExclusionArns, "exclusion-arns", "", nil, "Exclusion Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorFilter, "filter", "", "", "Filter")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorFindingArns, "finding-arns", "", nil, "Finding Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorLocale, "locale", "", "", "Locale")
	_inspectorCmd.Flags().StringVarP(&_inspectorMaxResults, "max-results", "", "", "Max Results")
	_inspectorCmd.Flags().StringVarP(&_inspectorNextToken, "next-token", "", "", "Next Token")
	_inspectorCmd.Flags().StringVarP(&_inspectorPreviewAgentsArn, "preview-agents-arn", "", "", "Preview Agents ARN")
	_inspectorCmd.Flags().StringVarP(&_inspectorPreviewToken, "preview-token", "", "", "Preview Token")
	_inspectorCmd.Flags().StringVarP(&_inspectorReportFileFormat, "report-file-format", "", "", "Report File Format")
	_inspectorCmd.Flags().StringVarP(&_inspectorReportType, "report-type", "", "", "Report Type")
	_inspectorCmd.Flags().StringVarP(&_inspectorResourceArn, "resource-arn", "", "", "Resource ARN")
	_inspectorCmd.Flags().StringVarP(&_inspectorResourceGroupArn, "resource-group-arn", "", "", "Resource Group ARN")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorResourceGroupArns, "resource-group-arns", "", nil, "Resource Group Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorResourceGroupTags, "resource-group-tags", "", "", "Resource Group Tags")
	_inspectorCmd.Flags().StringVarP(&_inspectorRoleArn, "role-arn", "", "", "Role ARN")
	_inspectorCmd.Flags().StringSliceVarP(&_inspectorRulesPackageArns, "rules-package-arns", "", nil, "Rules Package Arns")
	_inspectorCmd.Flags().StringVarP(&_inspectorStopAction, "stop-action", "", "", "Stop Action")
	_inspectorCmd.Flags().StringVarP(&_inspectorTags, "tags", "", "", "Tags")
	_inspectorCmd.Flags().StringVarP(&_inspectorTopicArn, "topic-arn", "", "", "Topic ARN")
	_inspectorCmd.Flags().StringVarP(&_inspectorUserAttributesForFindings, "user-attributes-for-findings", "", "", "User Attributes For Findings")

	_inspectorCmd.Flags().BoolVarP(&_inspectorAddAttributesToFindings, "add-attributes-to-findings", "", false, "Add Attributes To Findings")
	_inspectorCmd.Flags().BoolVarP(&_inspectorCreateAssessmentTarget, "create-assessment-target", "", false, "Create Assessment Target")
	_inspectorCmd.Flags().BoolVarP(&_inspectorCreateAssessmentTemplate, "create-assessment-template", "", false, "Create Assessment Template")
	_inspectorCmd.Flags().BoolVarP(&_inspectorCreateExclusionsPreview, "create-exclusions-preview", "", false, "Create Exclusions Preview")
	_inspectorCmd.Flags().BoolVarP(&_inspectorCreateResourceGroup, "create-resource-group", "", false, "Create Resource Group")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDeleteAssessmentRun, "delete-assessment-run", "", false, "Delete Assessment Run")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDeleteAssessmentTarget, "delete-assessment-target", "", false, "Delete Assessment Target")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDeleteAssessmentTemplate, "delete-assessment-template", "", false, "Delete Assessment Template")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeAssessmentRuns, "describe-assessment-runs", "", false, "Describe Assessment Runs")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeAssessmentTargets, "describe-assessment-targets", "", false, "Describe Assessment Targets")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeAssessmentTemplates, "describe-assessment-templates", "", false, "Describe Assessment Templates")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeCrossAccountAccessRole, "describe-cross-account-access-role", "", false, "Describe Cross Account Access Role")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeExclusions, "describe-exclusions", "", false, "Describe Exclusions")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeFindings, "describe-findings", "", false, "Describe Findings")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeResourceGroups, "describe-resource-groups", "", false, "Describe Resource Groups")
	_inspectorCmd.Flags().BoolVarP(&_inspectorDescribeRulesPackages, "describe-rules-packages", "", false, "Describe Rules Packages")
	_inspectorCmd.Flags().BoolVarP(&_inspectorGetAssessmentReport, "get-assessment-report", "", false, "Get Assessment Report")
	_inspectorCmd.Flags().BoolVarP(&_inspectorGetExclusionsPreview, "get-exclusions-preview", "", false, "Get Exclusions Preview")
	_inspectorCmd.Flags().BoolVarP(&_inspectorGetTelemetryMetadata, "get-telemetry-metadata", "", false, "Get Telemetry Metadata")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListAssessmentRunAgents, "list-assessment-run-agents", "", false, "List Assessment Run Agents")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListAssessmentRuns, "list-assessment-runs", "", false, "List Assessment Runs")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListAssessmentTargets, "list-assessment-targets", "", false, "List Assessment Targets")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListAssessmentTemplates, "list-assessment-templates", "", false, "List Assessment Templates")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListEventSubscriptions, "list-event-subscriptions", "", false, "List Event Subscriptions")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListExclusions, "list-exclusions", "", false, "List Exclusions")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListFindings, "list-findings", "", false, "List Findings")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListRulesPackages, "list-rules-packages", "", false, "List Rules Packages")
	_inspectorCmd.Flags().BoolVarP(&_inspectorListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_inspectorCmd.Flags().BoolVarP(&_inspectorPreviewAgents, "preview-agents", "", false, "Preview Agents")
	_inspectorCmd.Flags().BoolVarP(&_inspectorRegisterCrossAccountAccessRole, "register-cross-account-access-role", "", false, "Register Cross Account Access Role")
	_inspectorCmd.Flags().BoolVarP(&_inspectorRemoveAttributesFromFindings, "remove-attributes-from-findings", "", false, "Remove Attributes From Findings")
	_inspectorCmd.Flags().BoolVarP(&_inspectorSetTagsForResource, "set-tags-for-resource", "", false, "Set Tags For Resource")
	_inspectorCmd.Flags().BoolVarP(&_inspectorStartAssessmentRun, "start-assessment-run", "", false, "Start Assessment Run")
	_inspectorCmd.Flags().BoolVarP(&_inspectorStopAssessmentRun, "stop-assessment-run", "", false, "Stop Assessment Run")
	_inspectorCmd.Flags().BoolVarP(&_inspectorSubscribeToEvent, "subscribe-to-event", "", false, "Subscribe To Event")
	_inspectorCmd.Flags().BoolVarP(&_inspectorUnsubscribeFromEvent, "unsubscribe-from-event", "", false, "Unsubscribe From Event")
	_inspectorCmd.Flags().BoolVarP(&_inspectorUpdateAssessmentTarget, "update-assessment-target", "", false, "Update Assessment Target")

}
