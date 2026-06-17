package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resiliencehub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resiliencehubCmd represents the resiliencehub command
var _resiliencehubCmd = &cobra.Command{
	Use:   "resiliencehub",
	Short: "AWS resiliencehub CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := resiliencehub.NewFromConfig(cfg)
		if _resiliencehubAcceptResourceGroupingRecommendations {
			resiliencehub_AcceptResourceGroupingRecommendations(cfg, client)
			return
		}
		if _resiliencehubAddDraftAppVersionResourceMappings {
			resiliencehub_AddDraftAppVersionResourceMappings(cfg, client)
			return
		}
		if _resiliencehubBatchUpdateRecommendationStatus {
			resiliencehub_BatchUpdateRecommendationStatus(cfg, client)
			return
		}
		if _resiliencehubCreateApp {
			resiliencehub_CreateApp(cfg, client)
			return
		}
		if _resiliencehubCreateAppVersionAppComponent {
			resiliencehub_CreateAppVersionAppComponent(cfg, client)
			return
		}
		if _resiliencehubCreateAppVersionResource {
			resiliencehub_CreateAppVersionResource(cfg, client)
			return
		}
		if _resiliencehubCreateRecommendationTemplate {
			resiliencehub_CreateRecommendationTemplate(cfg, client)
			return
		}
		if _resiliencehubCreateResiliencyPolicy {
			resiliencehub_CreateResiliencyPolicy(cfg, client)
			return
		}
		if _resiliencehubDeleteApp {
			resiliencehub_DeleteApp(cfg, client)
			return
		}
		if _resiliencehubDeleteAppAssessment {
			resiliencehub_DeleteAppAssessment(cfg, client)
			return
		}
		if _resiliencehubDeleteAppInputSource {
			resiliencehub_DeleteAppInputSource(cfg, client)
			return
		}
		if _resiliencehubDeleteAppVersionAppComponent {
			resiliencehub_DeleteAppVersionAppComponent(cfg, client)
			return
		}
		if _resiliencehubDeleteAppVersionResource {
			resiliencehub_DeleteAppVersionResource(cfg, client)
			return
		}
		if _resiliencehubDeleteRecommendationTemplate {
			resiliencehub_DeleteRecommendationTemplate(cfg, client)
			return
		}
		if _resiliencehubDeleteResiliencyPolicy {
			resiliencehub_DeleteResiliencyPolicy(cfg, client)
			return
		}
		if _resiliencehubDescribeApp {
			resiliencehub_DescribeApp(cfg, client)
			return
		}
		if _resiliencehubDescribeAppAssessment {
			resiliencehub_DescribeAppAssessment(cfg, client)
			return
		}
		if _resiliencehubDescribeAppVersion {
			resiliencehub_DescribeAppVersion(cfg, client)
			return
		}
		if _resiliencehubDescribeAppVersionAppComponent {
			resiliencehub_DescribeAppVersionAppComponent(cfg, client)
			return
		}
		if _resiliencehubDescribeAppVersionResource {
			resiliencehub_DescribeAppVersionResource(cfg, client)
			return
		}
		if _resiliencehubDescribeAppVersionResourcesResolutionStatus {
			resiliencehub_DescribeAppVersionResourcesResolutionStatus(cfg, client)
			return
		}
		if _resiliencehubDescribeAppVersionTemplate {
			resiliencehub_DescribeAppVersionTemplate(cfg, client)
			return
		}
		if _resiliencehubDescribeDraftAppVersionResourcesImportStatus {
			resiliencehub_DescribeDraftAppVersionResourcesImportStatus(cfg, client)
			return
		}
		if _resiliencehubDescribeMetricsExport {
			resiliencehub_DescribeMetricsExport(cfg, client)
			return
		}
		if _resiliencehubDescribeResiliencyPolicy {
			resiliencehub_DescribeResiliencyPolicy(cfg, client)
			return
		}
		if _resiliencehubDescribeResourceGroupingRecommendationTask {
			resiliencehub_DescribeResourceGroupingRecommendationTask(cfg, client)
			return
		}
		if _resiliencehubImportResourcesToDraftAppVersion {
			resiliencehub_ImportResourcesToDraftAppVersion(cfg, client)
			return
		}
		if _resiliencehubListAlarmRecommendations {
			resiliencehub_ListAlarmRecommendations(cfg, client)
			return
		}
		if _resiliencehubListAppAssessmentComplianceDrifts {
			resiliencehub_ListAppAssessmentComplianceDrifts(cfg, client)
			return
		}
		if _resiliencehubListAppAssessmentResourceDrifts {
			resiliencehub_ListAppAssessmentResourceDrifts(cfg, client)
			return
		}
		if _resiliencehubListAppAssessments {
			resiliencehub_ListAppAssessments(cfg, client)
			return
		}
		if _resiliencehubListAppComponentCompliances {
			resiliencehub_ListAppComponentCompliances(cfg, client)
			return
		}
		if _resiliencehubListAppComponentRecommendations {
			resiliencehub_ListAppComponentRecommendations(cfg, client)
			return
		}
		if _resiliencehubListAppInputSources {
			resiliencehub_ListAppInputSources(cfg, client)
			return
		}
		if _resiliencehubListAppVersionAppComponents {
			resiliencehub_ListAppVersionAppComponents(cfg, client)
			return
		}
		if _resiliencehubListAppVersionResourceMappings {
			resiliencehub_ListAppVersionResourceMappings(cfg, client)
			return
		}
		if _resiliencehubListAppVersionResources {
			resiliencehub_ListAppVersionResources(cfg, client)
			return
		}
		if _resiliencehubListAppVersions {
			resiliencehub_ListAppVersions(cfg, client)
			return
		}
		if _resiliencehubListApps {
			resiliencehub_ListApps(cfg, client)
			return
		}
		if _resiliencehubListMetrics {
			resiliencehub_ListMetrics(cfg, client)
			return
		}
		if _resiliencehubListRecommendationTemplates {
			resiliencehub_ListRecommendationTemplates(cfg, client)
			return
		}
		if _resiliencehubListResiliencyPolicies {
			resiliencehub_ListResiliencyPolicies(cfg, client)
			return
		}
		if _resiliencehubListResourceGroupingRecommendations {
			resiliencehub_ListResourceGroupingRecommendations(cfg, client)
			return
		}
		if _resiliencehubListSopRecommendations {
			resiliencehub_ListSopRecommendations(cfg, client)
			return
		}
		if _resiliencehubListSuggestedResiliencyPolicies {
			resiliencehub_ListSuggestedResiliencyPolicies(cfg, client)
			return
		}
		if _resiliencehubListTagsForResource {
			resiliencehub_ListTagsForResource(cfg, client)
			return
		}
		if _resiliencehubListTestRecommendations {
			resiliencehub_ListTestRecommendations(cfg, client)
			return
		}
		if _resiliencehubListUnsupportedAppVersionResources {
			resiliencehub_ListUnsupportedAppVersionResources(cfg, client)
			return
		}
		if _resiliencehubPublishAppVersion {
			resiliencehub_PublishAppVersion(cfg, client)
			return
		}
		if _resiliencehubPutDraftAppVersionTemplate {
			resiliencehub_PutDraftAppVersionTemplate(cfg, client)
			return
		}
		if _resiliencehubRejectResourceGroupingRecommendations {
			resiliencehub_RejectResourceGroupingRecommendations(cfg, client)
			return
		}
		if _resiliencehubRemoveDraftAppVersionResourceMappings {
			resiliencehub_RemoveDraftAppVersionResourceMappings(cfg, client)
			return
		}
		if _resiliencehubResolveAppVersionResources {
			resiliencehub_ResolveAppVersionResources(cfg, client)
			return
		}
		if _resiliencehubStartAppAssessment {
			resiliencehub_StartAppAssessment(cfg, client)
			return
		}
		if _resiliencehubStartMetricsExport {
			resiliencehub_StartMetricsExport(cfg, client)
			return
		}
		if _resiliencehubStartResourceGroupingRecommendationTask {
			resiliencehub_StartResourceGroupingRecommendationTask(cfg, client)
			return
		}
		if _resiliencehubTagResource {
			resiliencehub_TagResource(cfg, client)
			return
		}
		if _resiliencehubUntagResource {
			resiliencehub_UntagResource(cfg, client)
			return
		}
		if _resiliencehubUpdateApp {
			resiliencehub_UpdateApp(cfg, client)
			return
		}
		if _resiliencehubUpdateAppVersion {
			resiliencehub_UpdateAppVersion(cfg, client)
			return
		}
		if _resiliencehubUpdateAppVersionAppComponent {
			resiliencehub_UpdateAppVersionAppComponent(cfg, client)
			return
		}
		if _resiliencehubUpdateAppVersionResource {
			resiliencehub_UpdateAppVersionResource(cfg, client)
			return
		}
		if _resiliencehubUpdateResiliencyPolicy {
			resiliencehub_UpdateResiliencyPolicy(cfg, client)
			return
		}

	},
}

var (
	_resiliencehubAcceptResourceGroupingRecommendations        bool
	_resiliencehubAddDraftAppVersionResourceMappings           bool
	_resiliencehubBatchUpdateRecommendationStatus              bool
	_resiliencehubCreateApp                                    bool
	_resiliencehubCreateAppVersionAppComponent                 bool
	_resiliencehubCreateAppVersionResource                     bool
	_resiliencehubCreateRecommendationTemplate                 bool
	_resiliencehubCreateResiliencyPolicy                       bool
	_resiliencehubDeleteApp                                    bool
	_resiliencehubDeleteAppAssessment                          bool
	_resiliencehubDeleteAppInputSource                         bool
	_resiliencehubDeleteAppVersionAppComponent                 bool
	_resiliencehubDeleteAppVersionResource                     bool
	_resiliencehubDeleteRecommendationTemplate                 bool
	_resiliencehubDeleteResiliencyPolicy                       bool
	_resiliencehubDescribeApp                                  bool
	_resiliencehubDescribeAppAssessment                        bool
	_resiliencehubDescribeAppVersion                           bool
	_resiliencehubDescribeAppVersionAppComponent               bool
	_resiliencehubDescribeAppVersionResource                   bool
	_resiliencehubDescribeAppVersionResourcesResolutionStatus  bool
	_resiliencehubDescribeAppVersionTemplate                   bool
	_resiliencehubDescribeDraftAppVersionResourcesImportStatus bool
	_resiliencehubDescribeMetricsExport                        bool
	_resiliencehubDescribeResiliencyPolicy                     bool
	_resiliencehubDescribeResourceGroupingRecommendationTask   bool
	_resiliencehubImportResourcesToDraftAppVersion             bool
	_resiliencehubListAlarmRecommendations                     bool
	_resiliencehubListAppAssessmentComplianceDrifts            bool
	_resiliencehubListAppAssessmentResourceDrifts              bool
	_resiliencehubListAppAssessments                           bool
	_resiliencehubListAppComponentCompliances                  bool
	_resiliencehubListAppComponentRecommendations              bool
	_resiliencehubListAppInputSources                          bool
	_resiliencehubListAppVersionAppComponents                  bool
	_resiliencehubListAppVersionResourceMappings               bool
	_resiliencehubListAppVersionResources                      bool
	_resiliencehubListAppVersions                              bool
	_resiliencehubListApps                                     bool
	_resiliencehubListMetrics                                  bool
	_resiliencehubListRecommendationTemplates                  bool
	_resiliencehubListResiliencyPolicies                       bool
	_resiliencehubListResourceGroupingRecommendations          bool
	_resiliencehubListSopRecommendations                       bool
	_resiliencehubListSuggestedResiliencyPolicies              bool
	_resiliencehubListTagsForResource                          bool
	_resiliencehubListTestRecommendations                      bool
	_resiliencehubListUnsupportedAppVersionResources           bool
	_resiliencehubPublishAppVersion                            bool
	_resiliencehubPutDraftAppVersionTemplate                   bool
	_resiliencehubRejectResourceGroupingRecommendations        bool
	_resiliencehubRemoveDraftAppVersionResourceMappings        bool
	_resiliencehubResolveAppVersionResources                   bool
	_resiliencehubStartAppAssessment                           bool
	_resiliencehubStartMetricsExport                           bool
	_resiliencehubStartResourceGroupingRecommendationTask      bool
	_resiliencehubTagResource                                  bool
	_resiliencehubUntagResource                                bool
	_resiliencehubUpdateApp                                    bool
	_resiliencehubUpdateAppVersion                             bool
	_resiliencehubUpdateAppVersionAppComponent                 bool
	_resiliencehubUpdateAppVersionResource                     bool
	_resiliencehubUpdateResiliencyPolicy                       bool

	_resiliencehubAdditionalInfo            string
	_resiliencehubAppArn                    string
	_resiliencehubAppComponents             []string
	_resiliencehubAppRegistryAppNames       []string
	_resiliencehubAppTemplateBody           string
	_resiliencehubAppVersion                string
	_resiliencehubAssessmentArn             string
	_resiliencehubAssessmentName            string
	_resiliencehubAssessmentSchedule        string
	_resiliencehubAssessmentStatus          string
	_resiliencehubAwsAccountId              string
	_resiliencehubAwsApplicationArn         string
	_resiliencehubAwsRegion                 string
	_resiliencehubBucketName                string
	_resiliencehubClearResiliencyPolicyArn  string
	_resiliencehubClientToken               string
	_resiliencehubComplianceStatus          string
	_resiliencehubConditions                string
	_resiliencehubDataLocationConstraint    string
	_resiliencehubDataSource                string
	_resiliencehubDescription               string
	_resiliencehubEksSourceClusterNamespace string
	_resiliencehubEksSourceNames            []string
	_resiliencehubEksSources                string
	_resiliencehubEndTime                   string
	_resiliencehubEntries                   string
	_resiliencehubEventSubscriptions        string
	_resiliencehubExcluded                  string
	_resiliencehubFields                    string
	_resiliencehubForceDelete               string
	_resiliencehubFormat                    string
	_resiliencehubFromLastAssessmentTime    string
	_resiliencehubGroupingId                string
	_resiliencehubId                        string
	_resiliencehubImportStrategy            string
	_resiliencehubInvoker                   string
	_resiliencehubLogicalResourceId         string
	_resiliencehubLogicalStackNames         []string
	_resiliencehubMaxResults                string
	_resiliencehubMetricsExportId           string
	_resiliencehubName                      string
	_resiliencehubNextToken                 string
	_resiliencehubPermissionModel           string
	_resiliencehubPhysicalResourceId        string
	_resiliencehubPolicy                    string
	_resiliencehubPolicyArn                 string
	_resiliencehubPolicyDescription         string
	_resiliencehubPolicyName                string
	_resiliencehubRecommendationIds         []string
	_resiliencehubRecommendationTemplateArn string
	_resiliencehubRecommendationTypes       string
	_resiliencehubRequestEntries            string
	_resiliencehubResolutionId              string
	_resiliencehubResourceArn               string
	_resiliencehubResourceGroupNames        []string
	_resiliencehubResourceMappings          string
	_resiliencehubResourceName              string
	_resiliencehubResourceNames             []string
	_resiliencehubResourceType              string
	_resiliencehubReverseOrder              string
	_resiliencehubSorts                     string
	_resiliencehubSourceArn                 string
	_resiliencehubSourceArns                []string
	_resiliencehubStartTime                 string
	_resiliencehubStatus                    string
	_resiliencehubTagKeys                   []string
	_resiliencehubTags                      string
	_resiliencehubTerraformSource           string
	_resiliencehubTerraformSourceNames      []string
	_resiliencehubTerraformSources          string
	_resiliencehubTier                      string
	_resiliencehubToLastAssessmentTime      string
	_resiliencehubType                      string
	_resiliencehubVersionName               string
)

// Accepts the resource grouping recommendations suggested by Resilience Hub for
// your application.
func resiliencehub_AcceptResourceGroupingRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.AcceptResourceGroupingRecommendationsInput{
		// AppArn: *string, // Required
		// Entries: []types.AcceptGroupingRecommendationEntry, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubEntries) > 0 {
		if err := assignInputField(input, "Entries", _resiliencehubEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.AcceptResourceGroupingRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the source of resource-maps to the draft version of an application. During
// assessment, Resilience Hub will use these resource-maps to resolve the latest
// physical ID for each resource in the application template. For more information
// about different types of resources supported by Resilience Hub and how to add
// them in your application, see [Step 2: How is your application managed?]in the Resilience Hub User Guide.
//
// [Step 2: How is your application managed?]: https://docs.aws.amazon.com/resilience-hub/latest/userguide/how-app-manage.html
func resiliencehub_AddDraftAppVersionResourceMappings(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.AddDraftAppVersionResourceMappingsInput{
		// AppArn: *string, // Required
		// ResourceMappings: []types.ResourceMapping, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubResourceMappings) > 0 {
		if err := assignInputField(input, "ResourceMappings", _resiliencehubResourceMappings); err != nil {
			log.Errorf("invalid --resource-mappings: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddDraftAppVersionResourceMappings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables you to include or exclude one or more operational recommendations.
func resiliencehub_BatchUpdateRecommendationStatus(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.BatchUpdateRecommendationStatusInput{
		// AppArn: *string, // Required
		// RequestEntries: []types.UpdateRecommendationStatusRequestEntry, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubRequestEntries) > 0 {
		if err := assignInputField(input, "RequestEntries", _resiliencehubRequestEntries); err != nil {
			log.Errorf("invalid --request-entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateRecommendationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Resilience Hub application. An Resilience Hub application is a
// collection of Amazon Web Services resources structured to prevent and recover
// Amazon Web Services application disruptions. To describe a Resilience Hub
// application, you provide an application name, resources from one or more
// CloudFormation stacks, Resource Groups, Terraform state files, AppRegistry
// applications, and an appropriate resiliency policy. In addition, you can also
// add resources that are located on Amazon Elastic Kubernetes Service (Amazon EKS)
// clusters as optional resources. For more information about the number of
// resources supported per application, see [Service quotas].
//
// After you create an Resilience Hub application, you publish it so that you can
// run a resiliency assessment on it. You can then use recommendations from the
// assessment to improve resiliency by running another assessment, comparing
// results, and then iterating the process until you achieve your goals for
// recovery time objective (RTO) and recovery point objective (RPO).
//
// [Service quotas]: https://docs.aws.amazon.com/general/latest/gr/resiliencehub.html#limits_resiliencehub
func resiliencehub_CreateApp(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.CreateAppInput{
		// Name: *string, // Required
	}

	if len(_resiliencehubName) > 0 {
		input.Name = aws.String(_resiliencehubName)
	}
	if len(_resiliencehubAssessmentSchedule) > 0 {
		if err := assignInputField(input, "AssessmentSchedule", _resiliencehubAssessmentSchedule); err != nil {
			log.Errorf("invalid --assessment-schedule: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubAwsApplicationArn) > 0 {
		input.AwsApplicationArn = aws.String(_resiliencehubAwsApplicationArn)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubDescription) > 0 {
		input.Description = aws.String(_resiliencehubDescription)
	}
	if len(_resiliencehubEventSubscriptions) > 0 {
		if err := assignInputField(input, "EventSubscriptions", _resiliencehubEventSubscriptions); err != nil {
			log.Errorf("invalid --event-subscriptions: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPermissionModel) > 0 {
		if err := assignInputField(input, "PermissionModel", _resiliencehubPermissionModel); err != nil {
			log.Errorf("invalid --permission-model: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPolicyArn) > 0 {
		input.PolicyArn = aws.String(_resiliencehubPolicyArn)
	}
	if len(_resiliencehubTags) > 0 {
		if err := assignInputField(input, "Tags", _resiliencehubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Application Component in the Resilience Hub application.
// This API updates the Resilience Hub application draft version. To use this
// Application Component for running assessments, you must publish the Resilience
// Hub application using the PublishAppVersion API.
func resiliencehub_CreateAppVersionAppComponent(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.CreateAppVersionAppComponentInput{
		// AppArn: *string, // Required
		// Name: *string, // Required
		// Type: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubName) > 0 {
		input.Name = aws.String(_resiliencehubName)
	}
	if len(_resiliencehubType) > 0 {
		input.Type = aws.String(_resiliencehubType)
	}
	if len(_resiliencehubAdditionalInfo) > 0 {
		if err := assignInputField(input, "AdditionalInfo", _resiliencehubAdditionalInfo); err != nil {
			log.Errorf("invalid --additional-info: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubId) > 0 {
		input.Id = aws.String(_resiliencehubId)
	}

	if resp, err := client.CreateAppVersionAppComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a resource to the Resilience Hub application and assigns it to the
// specified Application Components. If you specify a new Application Component,
// Resilience Hub will automatically create the Application Component.
//
// - This action has no effect outside Resilience Hub.
//
// - This API updates the Resilience Hub application draft version. To use this
// resource for running resiliency assessments, you must publish the Resilience Hub
// application using the PublishAppVersion API.
//
// - To update application version with new physicalResourceID , you must call
// ResolveAppVersionResources API.
func resiliencehub_CreateAppVersionResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.CreateAppVersionResourceInput{
		// AppArn: *string, // Required
		// AppComponents: []string, // Required
		// LogicalResourceId: *types.LogicalResourceId, // Required
		// PhysicalResourceId: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppComponents) > 0 {
		input.AppComponents = append([]string(nil), _resiliencehubAppComponents...)
	}
	if len(_resiliencehubLogicalResourceId) > 0 {
		if err := assignInputField(input, "LogicalResourceId", _resiliencehubLogicalResourceId); err != nil {
			log.Errorf("invalid --logical-resource-id: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPhysicalResourceId) > 0 {
		input.PhysicalResourceId = aws.String(_resiliencehubPhysicalResourceId)
	}
	if len(_resiliencehubResourceType) > 0 {
		input.ResourceType = aws.String(_resiliencehubResourceType)
	}
	if len(_resiliencehubAdditionalInfo) > 0 {
		if err := assignInputField(input, "AdditionalInfo", _resiliencehubAdditionalInfo); err != nil {
			log.Errorf("invalid --additional-info: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_resiliencehubAwsAccountId)
	}
	if len(_resiliencehubAwsRegion) > 0 {
		input.AwsRegion = aws.String(_resiliencehubAwsRegion)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubResourceName) > 0 {
		input.ResourceName = aws.String(_resiliencehubResourceName)
	}

	if resp, err := client.CreateAppVersionResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new recommendation template for the Resilience Hub application.
func resiliencehub_CreateRecommendationTemplate(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.CreateRecommendationTemplateInput{
		// AssessmentArn: *string, // Required
		// Name: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubName) > 0 {
		input.Name = aws.String(_resiliencehubName)
	}
	if len(_resiliencehubBucketName) > 0 {
		input.BucketName = aws.String(_resiliencehubBucketName)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubFormat) > 0 {
		if err := assignInputField(input, "Format", _resiliencehubFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubRecommendationIds) > 0 {
		input.RecommendationIds = append([]string(nil), _resiliencehubRecommendationIds...)
	}
	if len(_resiliencehubRecommendationTypes) > 0 {
		if err := assignInputField(input, "RecommendationTypes", _resiliencehubRecommendationTypes); err != nil {
			log.Errorf("invalid --recommendation-types: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubTags) > 0 {
		if err := assignInputField(input, "Tags", _resiliencehubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRecommendationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resiliency policy for an application.
// Resilience Hub allows you to provide a value of zero for rtoInSecs and rpoInSecs
// of your resiliency policy. But, while assessing your application, the lowest
// possible assessment result is near zero. Hence, if you provide value zero for
// rtoInSecs and rpoInSecs , the estimated workload RTO and estimated workload RPO
// result will be near zero and the Compliance status for your application will be
// set to Policy breached.
func resiliencehub_CreateResiliencyPolicy(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.CreateResiliencyPolicyInput{
		// Policy: map[string]types.FailurePolicy, // Required
		// PolicyName: *string, // Required
		// Tier: types.ResiliencyPolicyTier, // Required
	}

	if len(_resiliencehubPolicy) > 0 {
		if err := assignInputField(input, "Policy", _resiliencehubPolicy); err != nil {
			log.Errorf("invalid --policy: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPolicyName) > 0 {
		input.PolicyName = aws.String(_resiliencehubPolicyName)
	}
	if len(_resiliencehubTier) > 0 {
		if err := assignInputField(input, "Tier", _resiliencehubTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubDataLocationConstraint) > 0 {
		if err := assignInputField(input, "DataLocationConstraint", _resiliencehubDataLocationConstraint); err != nil {
			log.Errorf("invalid --data-location-constraint: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPolicyDescription) > 0 {
		input.PolicyDescription = aws.String(_resiliencehubPolicyDescription)
	}
	if len(_resiliencehubTags) > 0 {
		if err := assignInputField(input, "Tags", _resiliencehubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResiliencyPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Resilience Hub application. This is a destructive action that can't
// be undone.
func resiliencehub_DeleteApp(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteAppInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _resiliencehubForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Resilience Hub application assessment. This is a destructive action
// that can't be undone.
func resiliencehub_DeleteAppAssessment(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteAppAssessmentInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}

	if resp, err := client.DeleteAppAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the input source and all of its imported resources from the Resilience
// Hub application.
func resiliencehub_DeleteAppInputSource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteAppInputSourceInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubEksSourceClusterNamespace) > 0 {
		if err := assignInputField(input, "EksSourceClusterNamespace", _resiliencehubEksSourceClusterNamespace); err != nil {
			log.Errorf("invalid --eks-source-cluster-namespace: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubSourceArn) > 0 {
		input.SourceArn = aws.String(_resiliencehubSourceArn)
	}
	if len(_resiliencehubTerraformSource) > 0 {
		if err := assignInputField(input, "TerraformSource", _resiliencehubTerraformSource); err != nil {
			log.Errorf("invalid --terraform-source: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteAppInputSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Application Component from the Resilience Hub application.
// - This API updates the Resilience Hub application draft version. To use this
// Application Component for running assessments, you must publish the Resilience
// Hub application using the PublishAppVersion API.
//
// - You will not be able to delete an Application Component if it has resources
// associated with it.
func resiliencehub_DeleteAppVersionAppComponent(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteAppVersionAppComponentInput{
		// AppArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubId) > 0 {
		input.Id = aws.String(_resiliencehubId)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}

	if resp, err := client.DeleteAppVersionAppComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resource from the Resilience Hub application.
// - You can only delete a manually added resource. To exclude non-manually
// added resources, use the UpdateAppVersionResource API.
//
// - This action has no effect outside Resilience Hub.
//
// - This API updates the Resilience Hub application draft version. To use this
// resource for running resiliency assessments, you must publish the Resilience Hub
// application using the PublishAppVersion API.
func resiliencehub_DeleteAppVersionResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteAppVersionResourceInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_resiliencehubAwsAccountId)
	}
	if len(_resiliencehubAwsRegion) > 0 {
		input.AwsRegion = aws.String(_resiliencehubAwsRegion)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubLogicalResourceId) > 0 {
		if err := assignInputField(input, "LogicalResourceId", _resiliencehubLogicalResourceId); err != nil {
			log.Errorf("invalid --logical-resource-id: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPhysicalResourceId) > 0 {
		input.PhysicalResourceId = aws.String(_resiliencehubPhysicalResourceId)
	}
	if len(_resiliencehubResourceName) > 0 {
		input.ResourceName = aws.String(_resiliencehubResourceName)
	}

	if resp, err := client.DeleteAppVersionResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a recommendation template. This is a destructive action that can't be
// undone.
func resiliencehub_DeleteRecommendationTemplate(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteRecommendationTemplateInput{
		// RecommendationTemplateArn: *string, // Required
	}

	if len(_resiliencehubRecommendationTemplateArn) > 0 {
		input.RecommendationTemplateArn = aws.String(_resiliencehubRecommendationTemplateArn)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}

	if resp, err := client.DeleteRecommendationTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a resiliency policy. This is a destructive action that can't be undone.
func resiliencehub_DeleteResiliencyPolicy(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DeleteResiliencyPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_resiliencehubPolicyArn) > 0 {
		input.PolicyArn = aws.String(_resiliencehubPolicyArn)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}

	if resp, err := client.DeleteResiliencyPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Resilience Hub application.
func resiliencehub_DescribeApp(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}

	if resp, err := client.DescribeApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an assessment for an Resilience Hub application.
func resiliencehub_DescribeAppAssessment(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppAssessmentInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}

	if resp, err := client.DescribeAppAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the Resilience Hub application version.
func resiliencehub_DescribeAppVersion(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppVersionInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}

	if resp, err := client.DescribeAppVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes an Application Component in the Resilience Hub application.
func resiliencehub_DescribeAppVersionAppComponent(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppVersionAppComponentInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
		// Id: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubId) > 0 {
		input.Id = aws.String(_resiliencehubId)
	}

	if resp, err := client.DescribeAppVersionAppComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a resource of the Resilience Hub application.
// This API accepts only one of the following parameters to describe the resource:
//
// - resourceName
//
// - logicalResourceId
//
// - physicalResourceId (Along with physicalResourceId , you can also provide
// awsAccountId , and awsRegion )
func resiliencehub_DescribeAppVersionResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppVersionResourceInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_resiliencehubAwsAccountId)
	}
	if len(_resiliencehubAwsRegion) > 0 {
		input.AwsRegion = aws.String(_resiliencehubAwsRegion)
	}
	if len(_resiliencehubLogicalResourceId) > 0 {
		if err := assignInputField(input, "LogicalResourceId", _resiliencehubLogicalResourceId); err != nil {
			log.Errorf("invalid --logical-resource-id: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPhysicalResourceId) > 0 {
		input.PhysicalResourceId = aws.String(_resiliencehubPhysicalResourceId)
	}
	if len(_resiliencehubResourceName) > 0 {
		input.ResourceName = aws.String(_resiliencehubResourceName)
	}

	if resp, err := client.DescribeAppVersionResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resolution status for the specified resolution identifier for an
// application version. If resolutionId is not specified, the current resolution
// status is returned.
func resiliencehub_DescribeAppVersionResourcesResolutionStatus(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppVersionResourcesResolutionStatusInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubResolutionId) > 0 {
		input.ResolutionId = aws.String(_resiliencehubResolutionId)
	}

	if resp, err := client.DescribeAppVersionResourcesResolutionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes details about an Resilience Hub application.
func resiliencehub_DescribeAppVersionTemplate(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeAppVersionTemplateInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}

	if resp, err := client.DescribeAppVersionTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the status of importing resources to an application version.
// If you get a 404 error with ResourceImportStatusNotFoundAppMetadataException ,
// you must call importResourcesToDraftAppVersion after creating the application
// and before calling describeDraftAppVersionResourcesImportStatus to obtain the
// status.
func resiliencehub_DescribeDraftAppVersionResourcesImportStatus(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeDraftAppVersionResourcesImportStatusInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}

	if resp, err := client.DescribeDraftAppVersionResourcesImportStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the metrics of the application configuration being exported.
func resiliencehub_DescribeMetricsExport(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeMetricsExportInput{
		// MetricsExportId: *string, // Required
	}

	if len(_resiliencehubMetricsExportId) > 0 {
		input.MetricsExportId = aws.String(_resiliencehubMetricsExportId)
	}

	if resp, err := client.DescribeMetricsExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a specified resiliency policy for an Resilience Hub application. The
// returned policy object includes creation time, data location constraints, the
// Amazon Resource Name (ARN) for the policy, tags, tier, and more.
func resiliencehub_DescribeResiliencyPolicy(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeResiliencyPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_resiliencehubPolicyArn) > 0 {
		input.PolicyArn = aws.String(_resiliencehubPolicyArn)
	}

	if resp, err := client.DescribeResiliencyPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the resource grouping recommendation tasks run by Resilience Hub for
// your application.
func resiliencehub_DescribeResourceGroupingRecommendationTask(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.DescribeResourceGroupingRecommendationTaskInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubGroupingId) > 0 {
		input.GroupingId = aws.String(_resiliencehubGroupingId)
	}

	if resp, err := client.DescribeResourceGroupingRecommendationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Imports resources to Resilience Hub application draft version from different
// input sources. For more information about the input sources supported by
// Resilience Hub, see [Discover the structure and describe your Resilience Hub application].
//
// [Discover the structure and describe your Resilience Hub application]: https://docs.aws.amazon.com/resilience-hub/latest/userguide/discover-structure.html
func resiliencehub_ImportResourcesToDraftAppVersion(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ImportResourcesToDraftAppVersionInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubEksSources) > 0 {
		if err := assignInputField(input, "EksSources", _resiliencehubEksSources); err != nil {
			log.Errorf("invalid --eks-sources: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubImportStrategy) > 0 {
		if err := assignInputField(input, "ImportStrategy", _resiliencehubImportStrategy); err != nil {
			log.Errorf("invalid --import-strategy: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubSourceArns) > 0 {
		input.SourceArns = append([]string(nil), _resiliencehubSourceArns...)
	}
	if len(_resiliencehubTerraformSources) > 0 {
		if err := assignInputField(input, "TerraformSources", _resiliencehubTerraformSources); err != nil {
			log.Errorf("invalid --terraform-sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.ImportResourcesToDraftAppVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the alarm recommendations for an Resilience Hub application.
func resiliencehub_ListAlarmRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAlarmRecommendationsInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAlarmRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAlarmRecommendationsOutput
	p := resiliencehub.NewListAlarmRecommendationsPaginator(client, input)
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

// List of compliance drifts that were detected while running an assessment.
func resiliencehub_ListAppAssessmentComplianceDrifts(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppAssessmentComplianceDriftsInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppAssessmentComplianceDrifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppAssessmentComplianceDriftsOutput
	p := resiliencehub.NewListAppAssessmentComplianceDriftsPaginator(client, input)
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

// List of resource drifts that were detected while running an assessment.
func resiliencehub_ListAppAssessmentResourceDrifts(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppAssessmentResourceDriftsInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppAssessmentResourceDrifts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppAssessmentResourceDriftsOutput
	p := resiliencehub.NewListAppAssessmentResourceDriftsPaginator(client, input)
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

// Lists the assessments for an Resilience Hub application. You can use request
// parameters to refine the results for the response object.
func resiliencehub_ListAppAssessments(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppAssessmentsInput{}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAssessmentName) > 0 {
		input.AssessmentName = aws.String(_resiliencehubAssessmentName)
	}
	if len(_resiliencehubAssessmentStatus) > 0 {
		if err := assignInputField(input, "AssessmentStatus", _resiliencehubAssessmentStatus); err != nil {
			log.Errorf("invalid --assessment-status: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubComplianceStatus) > 0 {
		if err := assignInputField(input, "ComplianceStatus", _resiliencehubComplianceStatus); err != nil {
			log.Errorf("invalid --compliance-status: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubInvoker) > 0 {
		if err := assignInputField(input, "Invoker", _resiliencehubInvoker); err != nil {
			log.Errorf("invalid --invoker: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _resiliencehubReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAppAssessments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppAssessmentsOutput
	p := resiliencehub.NewListAppAssessmentsPaginator(client, input)
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

// Lists the compliances for an Resilience Hub Application Component.
func resiliencehub_ListAppComponentCompliances(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppComponentCompliancesInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppComponentCompliances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppComponentCompliancesOutput
	p := resiliencehub.NewListAppComponentCompliancesPaginator(client, input)
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

// Lists the recommendations for an Resilience Hub Application Component.
func resiliencehub_ListAppComponentRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppComponentRecommendationsInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppComponentRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppComponentRecommendationsOutput
	p := resiliencehub.NewListAppComponentRecommendationsPaginator(client, input)
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

// Lists all the input sources of the Resilience Hub application. For more
// information about the input sources supported by Resilience Hub, see [Discover the structure and describe your Resilience Hub application].
//
// [Discover the structure and describe your Resilience Hub application]: https://docs.aws.amazon.com/resilience-hub/latest/userguide/discover-structure.html
func resiliencehub_ListAppInputSources(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppInputSourcesInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppInputSources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppInputSourcesOutput
	p := resiliencehub.NewListAppInputSourcesPaginator(client, input)
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

// Lists all the Application Components in the Resilience Hub application.
func resiliencehub_ListAppVersionAppComponents(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppVersionAppComponentsInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppVersionAppComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppVersionAppComponentsOutput
	p := resiliencehub.NewListAppVersionAppComponentsPaginator(client, input)
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

// Lists how the resources in an application version are mapped/sourced from.
// Mappings can be physical resource identifiers, CloudFormation stacks,
// resource-groups, or an application registry app.
func resiliencehub_ListAppVersionResourceMappings(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppVersionResourceMappingsInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAppVersionResourceMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppVersionResourceMappingsOutput
	p := resiliencehub.NewListAppVersionResourceMappingsPaginator(client, input)
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

// Lists all the resources in an Resilience Hub application.
func resiliencehub_ListAppVersionResources(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppVersionResourcesInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubResolutionId) > 0 {
		input.ResolutionId = aws.String(_resiliencehubResolutionId)
	}

	if disablePaginator() {
		if resp, err := client.ListAppVersionResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppVersionResourcesOutput
	p := resiliencehub.NewListAppVersionResourcesPaginator(client, input)
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

// Lists the different versions for the Resilience Hub applications.
func resiliencehub_ListAppVersions(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppVersionsInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _resiliencehubEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _resiliencehubStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAppVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppVersionsOutput
	p := resiliencehub.NewListAppVersionsPaginator(client, input)
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

// Lists your Resilience Hub applications.
// You can filter applications using only one filter at a time or without using
// any filter. If you try to filter applications using multiple filters, you will
// get the following error:
//
// An error occurred (ValidationException) when calling the ListApps operation:
// Only one filter is supported for this operation.
func resiliencehub_ListApps(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListAppsInput{}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAwsApplicationArn) > 0 {
		input.AwsApplicationArn = aws.String(_resiliencehubAwsApplicationArn)
	}
	if len(_resiliencehubFromLastAssessmentTime) > 0 {
		if err := assignInputField(input, "FromLastAssessmentTime", _resiliencehubFromLastAssessmentTime); err != nil {
			log.Errorf("invalid --from-last-assessment-time: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubName) > 0 {
		input.Name = aws.String(_resiliencehubName)
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _resiliencehubReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubToLastAssessmentTime) > 0 {
		if err := assignInputField(input, "ToLastAssessmentTime", _resiliencehubToLastAssessmentTime); err != nil {
			log.Errorf("invalid --to-last-assessment-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListApps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListAppsOutput
	p := resiliencehub.NewListAppsPaginator(client, input)
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

// Lists the metrics that can be exported.
func resiliencehub_ListMetrics(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListMetricsInput{}

	if len(_resiliencehubConditions) > 0 {
		if err := assignInputField(input, "Conditions", _resiliencehubConditions); err != nil {
			log.Errorf("invalid --conditions: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubDataSource) > 0 {
		input.DataSource = aws.String(_resiliencehubDataSource)
	}
	if len(_resiliencehubFields) > 0 {
		if err := assignInputField(input, "Fields", _resiliencehubFields); err != nil {
			log.Errorf("invalid --fields: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubSorts) > 0 {
		if err := assignInputField(input, "Sorts", _resiliencehubSorts); err != nil {
			log.Errorf("invalid --sorts: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMetrics(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListMetricsOutput
	p := resiliencehub.NewListMetricsPaginator(client, input)
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

// Lists the recommendation templates for the Resilience Hub applications.
func resiliencehub_ListRecommendationTemplates(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListRecommendationTemplatesInput{}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubName) > 0 {
		input.Name = aws.String(_resiliencehubName)
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubRecommendationTemplateArn) > 0 {
		input.RecommendationTemplateArn = aws.String(_resiliencehubRecommendationTemplateArn)
	}
	if len(_resiliencehubReverseOrder) > 0 {
		if err := assignInputField(input, "ReverseOrder", _resiliencehubReverseOrder); err != nil {
			log.Errorf("invalid --reverse-order: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubStatus) > 0 {
		if err := assignInputField(input, "Status", _resiliencehubStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRecommendationTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListRecommendationTemplatesOutput
	p := resiliencehub.NewListRecommendationTemplatesPaginator(client, input)
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

// Lists the resiliency policies for the Resilience Hub applications.
func resiliencehub_ListResiliencyPolicies(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListResiliencyPoliciesInput{}

	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubPolicyName) > 0 {
		input.PolicyName = aws.String(_resiliencehubPolicyName)
	}

	if disablePaginator() {
		if resp, err := client.ListResiliencyPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListResiliencyPoliciesOutput
	p := resiliencehub.NewListResiliencyPoliciesPaginator(client, input)
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

// Lists the resource grouping recommendations suggested by Resilience Hub for
// your application.
func resiliencehub_ListResourceGroupingRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListResourceGroupingRecommendationsInput{}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceGroupingRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListResourceGroupingRecommendationsOutput
	p := resiliencehub.NewListResourceGroupingRecommendationsPaginator(client, input)
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

// Lists the standard operating procedure (SOP) recommendations for the Resilience
// Hub applications.
func resiliencehub_ListSopRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListSopRecommendationsInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSopRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListSopRecommendationsOutput
	p := resiliencehub.NewListSopRecommendationsPaginator(client, input)
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

// Lists the suggested resiliency policies for the Resilience Hub applications.
func resiliencehub_ListSuggestedResiliencyPolicies(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListSuggestedResiliencyPoliciesInput{}

	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSuggestedResiliencyPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListSuggestedResiliencyPoliciesOutput
	p := resiliencehub.NewListSuggestedResiliencyPoliciesPaginator(client, input)
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

// Lists the tags for your resources in your Resilience Hub applications.
func resiliencehub_ListTagsForResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_resiliencehubResourceArn) > 0 {
		input.ResourceArn = aws.String(_resiliencehubResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the test recommendations for the Resilience Hub application.
func resiliencehub_ListTestRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListTestRecommendationsInput{
		// AssessmentArn: *string, // Required
	}

	if len(_resiliencehubAssessmentArn) > 0 {
		input.AssessmentArn = aws.String(_resiliencehubAssessmentArn)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTestRecommendations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListTestRecommendationsOutput
	p := resiliencehub.NewListTestRecommendationsPaginator(client, input)
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

// Lists the resources that are not currently supported in Resilience Hub. An
// unsupported resource is a resource that exists in the object that was used to
// create an app, but is not supported by Resilience Hub.
func resiliencehub_ListUnsupportedAppVersionResources(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ListUnsupportedAppVersionResourcesInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resiliencehubMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubNextToken) > 0 {
		input.NextToken = aws.String(_resiliencehubNextToken)
	}
	if len(_resiliencehubResolutionId) > 0 {
		input.ResolutionId = aws.String(_resiliencehubResolutionId)
	}

	if disablePaginator() {
		if resp, err := client.ListUnsupportedAppVersionResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resiliencehub.ListUnsupportedAppVersionResourcesOutput
	p := resiliencehub.NewListUnsupportedAppVersionResourcesPaginator(client, input)
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

// Publishes a new version of a specific Resilience Hub application.
func resiliencehub_PublishAppVersion(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.PublishAppVersionInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubVersionName) > 0 {
		input.VersionName = aws.String(_resiliencehubVersionName)
	}

	if resp, err := client.PublishAppVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates the app template for an Resilience Hub application draft
// version.
func resiliencehub_PutDraftAppVersionTemplate(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.PutDraftAppVersionTemplateInput{
		// AppArn: *string, // Required
		// AppTemplateBody: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppTemplateBody) > 0 {
		input.AppTemplateBody = aws.String(_resiliencehubAppTemplateBody)
	}

	if resp, err := client.PutDraftAppVersionTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects resource grouping recommendations.
func resiliencehub_RejectResourceGroupingRecommendations(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.RejectResourceGroupingRecommendationsInput{
		// AppArn: *string, // Required
		// Entries: []types.RejectGroupingRecommendationEntry, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubEntries) > 0 {
		if err := assignInputField(input, "Entries", _resiliencehubEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.RejectResourceGroupingRecommendations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes resource mappings from a draft application version.
func resiliencehub_RemoveDraftAppVersionResourceMappings(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.RemoveDraftAppVersionResourceMappingsInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppRegistryAppNames) > 0 {
		input.AppRegistryAppNames = append([]string(nil), _resiliencehubAppRegistryAppNames...)
	}
	if len(_resiliencehubEksSourceNames) > 0 {
		input.EksSourceNames = append([]string(nil), _resiliencehubEksSourceNames...)
	}
	if len(_resiliencehubLogicalStackNames) > 0 {
		input.LogicalStackNames = append([]string(nil), _resiliencehubLogicalStackNames...)
	}
	if len(_resiliencehubResourceGroupNames) > 0 {
		input.ResourceGroupNames = append([]string(nil), _resiliencehubResourceGroupNames...)
	}
	if len(_resiliencehubResourceNames) > 0 {
		input.ResourceNames = append([]string(nil), _resiliencehubResourceNames...)
	}
	if len(_resiliencehubTerraformSourceNames) > 0 {
		input.TerraformSourceNames = append([]string(nil), _resiliencehubTerraformSourceNames...)
	}

	if resp, err := client.RemoveDraftAppVersionResourceMappings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Resolves the resources for an application version.
func resiliencehub_ResolveAppVersionResources(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.ResolveAppVersionResourcesInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}

	if resp, err := client.ResolveAppVersionResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new application assessment for an application.
func resiliencehub_StartAppAssessment(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.StartAppAssessmentInput{
		// AppArn: *string, // Required
		// AppVersion: *string, // Required
		// AssessmentName: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAppVersion) > 0 {
		input.AppVersion = aws.String(_resiliencehubAppVersion)
	}
	if len(_resiliencehubAssessmentName) > 0 {
		input.AssessmentName = aws.String(_resiliencehubAssessmentName)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}
	if len(_resiliencehubTags) > 0 {
		if err := assignInputField(input, "Tags", _resiliencehubTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartAppAssessment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the export task of metrics.
func resiliencehub_StartMetricsExport(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.StartMetricsExportInput{}

	if len(_resiliencehubBucketName) > 0 {
		input.BucketName = aws.String(_resiliencehubBucketName)
	}
	if len(_resiliencehubClientToken) > 0 {
		input.ClientToken = aws.String(_resiliencehubClientToken)
	}

	if resp, err := client.StartMetricsExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts grouping recommendation task.
func resiliencehub_StartResourceGroupingRecommendationTask(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.StartResourceGroupingRecommendationTaskInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}

	if resp, err := client.StartResourceGroupingRecommendationTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies one or more tags to a resource.
func resiliencehub_TagResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_resiliencehubResourceArn) > 0 {
		input.ResourceArn = aws.String(_resiliencehubResourceArn)
	}
	if len(_resiliencehubTags) > 0 {
		if err := assignInputField(input, "Tags", _resiliencehubTags); err != nil {
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

// Removes one or more tags from a resource.
func resiliencehub_UntagResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_resiliencehubResourceArn) > 0 {
		input.ResourceArn = aws.String(_resiliencehubResourceArn)
	}
	if len(_resiliencehubTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _resiliencehubTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an application.
func resiliencehub_UpdateApp(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.UpdateAppInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAssessmentSchedule) > 0 {
		if err := assignInputField(input, "AssessmentSchedule", _resiliencehubAssessmentSchedule); err != nil {
			log.Errorf("invalid --assessment-schedule: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubClearResiliencyPolicyArn) > 0 {
		if err := assignInputField(input, "ClearResiliencyPolicyArn", _resiliencehubClearResiliencyPolicyArn); err != nil {
			log.Errorf("invalid --clear-resiliency-policy-arn: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubDescription) > 0 {
		input.Description = aws.String(_resiliencehubDescription)
	}
	if len(_resiliencehubEventSubscriptions) > 0 {
		if err := assignInputField(input, "EventSubscriptions", _resiliencehubEventSubscriptions); err != nil {
			log.Errorf("invalid --event-subscriptions: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPermissionModel) > 0 {
		if err := assignInputField(input, "PermissionModel", _resiliencehubPermissionModel); err != nil {
			log.Errorf("invalid --permission-model: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPolicyArn) > 0 {
		input.PolicyArn = aws.String(_resiliencehubPolicyArn)
	}

	if resp, err := client.UpdateApp(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the Resilience Hub application version.
// This API updates the Resilience Hub application draft version. To use this
// information for running resiliency assessments, you must publish the Resilience
// Hub application using the PublishAppVersion API.
func resiliencehub_UpdateAppVersion(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.UpdateAppVersionInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAdditionalInfo) > 0 {
		if err := assignInputField(input, "AdditionalInfo", _resiliencehubAdditionalInfo); err != nil {
			log.Errorf("invalid --additional-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAppVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing Application Component in the Resilience Hub application.
// This API updates the Resilience Hub application draft version. To use this
// Application Component for running assessments, you must publish the Resilience
// Hub application using the PublishAppVersion API.
func resiliencehub_UpdateAppVersionAppComponent(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.UpdateAppVersionAppComponentInput{
		// AppArn: *string, // Required
		// Id: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubId) > 0 {
		input.Id = aws.String(_resiliencehubId)
	}
	if len(_resiliencehubAdditionalInfo) > 0 {
		if err := assignInputField(input, "AdditionalInfo", _resiliencehubAdditionalInfo); err != nil {
			log.Errorf("invalid --additional-info: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubName) > 0 {
		input.Name = aws.String(_resiliencehubName)
	}
	if len(_resiliencehubType) > 0 {
		input.Type = aws.String(_resiliencehubType)
	}

	if resp, err := client.UpdateAppVersionAppComponent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource details in the Resilience Hub application.
// - This action has no effect outside Resilience Hub.
//
// - This API updates the Resilience Hub application draft version. To use this
// resource for running resiliency assessments, you must publish the Resilience Hub
// application using the PublishAppVersion API.
//
// - To update application version with new physicalResourceID , you must call
// ResolveAppVersionResources API.
func resiliencehub_UpdateAppVersionResource(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.UpdateAppVersionResourceInput{
		// AppArn: *string, // Required
	}

	if len(_resiliencehubAppArn) > 0 {
		input.AppArn = aws.String(_resiliencehubAppArn)
	}
	if len(_resiliencehubAdditionalInfo) > 0 {
		if err := assignInputField(input, "AdditionalInfo", _resiliencehubAdditionalInfo); err != nil {
			log.Errorf("invalid --additional-info: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubAppComponents) > 0 {
		input.AppComponents = append([]string(nil), _resiliencehubAppComponents...)
	}
	if len(_resiliencehubAwsAccountId) > 0 {
		input.AwsAccountId = aws.String(_resiliencehubAwsAccountId)
	}
	if len(_resiliencehubAwsRegion) > 0 {
		input.AwsRegion = aws.String(_resiliencehubAwsRegion)
	}
	if len(_resiliencehubExcluded) > 0 {
		if err := assignInputField(input, "Excluded", _resiliencehubExcluded); err != nil {
			log.Errorf("invalid --excluded: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubLogicalResourceId) > 0 {
		if err := assignInputField(input, "LogicalResourceId", _resiliencehubLogicalResourceId); err != nil {
			log.Errorf("invalid --logical-resource-id: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPhysicalResourceId) > 0 {
		input.PhysicalResourceId = aws.String(_resiliencehubPhysicalResourceId)
	}
	if len(_resiliencehubResourceName) > 0 {
		input.ResourceName = aws.String(_resiliencehubResourceName)
	}
	if len(_resiliencehubResourceType) > 0 {
		input.ResourceType = aws.String(_resiliencehubResourceType)
	}

	if resp, err := client.UpdateAppVersionResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a resiliency policy.
// Resilience Hub allows you to provide a value of zero for rtoInSecs and rpoInSecs
// of your resiliency policy. But, while assessing your application, the lowest
// possible assessment result is near zero. Hence, if you provide value zero for
// rtoInSecs and rpoInSecs , the estimated workload RTO and estimated workload RPO
// result will be near zero and the Compliance status for your application will be
// set to Policy breached.
func resiliencehub_UpdateResiliencyPolicy(cfg aws.Config, client *resiliencehub.Client) {
	input := &resiliencehub.UpdateResiliencyPolicyInput{
		// PolicyArn: *string, // Required
	}

	if len(_resiliencehubPolicyArn) > 0 {
		input.PolicyArn = aws.String(_resiliencehubPolicyArn)
	}
	if len(_resiliencehubDataLocationConstraint) > 0 {
		if err := assignInputField(input, "DataLocationConstraint", _resiliencehubDataLocationConstraint); err != nil {
			log.Errorf("invalid --data-location-constraint: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPolicy) > 0 {
		if err := assignInputField(input, "Policy", _resiliencehubPolicy); err != nil {
			log.Errorf("invalid --policy: %s", err.Error())
			return
		}
	}
	if len(_resiliencehubPolicyDescription) > 0 {
		input.PolicyDescription = aws.String(_resiliencehubPolicyDescription)
	}
	if len(_resiliencehubPolicyName) > 0 {
		input.PolicyName = aws.String(_resiliencehubPolicyName)
	}
	if len(_resiliencehubTier) > 0 {
		if err := assignInputField(input, "Tier", _resiliencehubTier); err != nil {
			log.Errorf("invalid --tier: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResiliencyPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_resiliencehubCmd)
	_resiliencehubCmd.Flags().SortFlags = false

	_resiliencehubCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_resiliencehubCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_resiliencehubCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAdditionalInfo, "additional-info", "", "", "Additional Info")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAppArn, "app-arn", "", "", "App ARN")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubAppComponents, "app-components", "", nil, "App Components")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubAppRegistryAppNames, "app-registry-app-names", "", nil, "App Registry App Names")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAppTemplateBody, "app-template-body", "", "", "App Template Body")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAppVersion, "app-version", "", "", "App Version")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAssessmentArn, "assessment-arn", "", "", "Assessment ARN")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAssessmentName, "assessment-name", "", "", "Assessment Name")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAssessmentSchedule, "assessment-schedule", "", "", "Assessment Schedule")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAssessmentStatus, "assessment-status", "", "", "Assessment Status")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAwsAccountId, "aws-account-id", "", "", "AWS Account ID")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAwsApplicationArn, "aws-application-arn", "", "", "AWS Application ARN")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubAwsRegion, "aws-region", "", "", "AWS Region")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubBucketName, "bucket-name", "", "", "Bucket Name")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubClearResiliencyPolicyArn, "clear-resiliency-policy-arn", "", "", "Clear Resiliency Policy ARN")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubClientToken, "client-token", "", "", "Client Token")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubComplianceStatus, "compliance-status", "", "", "Compliance Status")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubConditions, "conditions", "", "", "Conditions")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubDataLocationConstraint, "data-location-constraint", "", "", "Data Location Constraint")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubDataSource, "data-source", "", "", "Data Source")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubDescription, "description", "", "", "Description")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubEksSourceClusterNamespace, "eks-source-cluster-namespace", "", "", "Eks Source Cluster Namespace")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubEksSourceNames, "eks-source-names", "", nil, "Eks Source Names")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubEksSources, "eks-sources", "", "", "Eks Sources")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubEndTime, "end-time", "", "", "End Time")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubEntries, "entries", "", "", "Entries")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubEventSubscriptions, "event-subscriptions", "", "", "Event Subscriptions")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubExcluded, "excluded", "", "", "Excluded")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubFields, "fields", "", "", "Fields")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubForceDelete, "force-delete", "", "", "Force Delete")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubFormat, "format", "", "", "Format")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubFromLastAssessmentTime, "from-last-assessment-time", "", "", "From Last Assessment Time")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubGroupingId, "grouping-id", "", "", "Grouping ID")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubId, "id", "", "", "ID")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubImportStrategy, "import-strategy", "", "", "Import Strategy")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubInvoker, "invoker", "", "", "Invoker")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubLogicalResourceId, "logical-resource-id", "", "", "Logical Resource ID")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubLogicalStackNames, "logical-stack-names", "", nil, "Logical Stack Names")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubMaxResults, "max-results", "", "", "Max Results")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubMetricsExportId, "metrics-export-id", "", "", "Metrics Export ID")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubName, "name", "", "", "Name")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubNextToken, "next-token", "", "", "Next Token")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubPermissionModel, "permission-model", "", "", "Permission Model")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubPhysicalResourceId, "physical-resource-id", "", "", "Physical Resource ID")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubPolicy, "policy", "", "", "Policy")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubPolicyArn, "policy-arn", "", "", "Policy ARN")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubPolicyDescription, "policy-description", "", "", "Policy Description")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubPolicyName, "policy-name", "", "", "Policy Name")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubRecommendationIds, "recommendation-ids", "", nil, "Recommendation Ids")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubRecommendationTemplateArn, "recommendation-template-arn", "", "", "Recommendation Template ARN")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubRecommendationTypes, "recommendation-types", "", "", "Recommendation Types")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubRequestEntries, "request-entries", "", "", "Request Entries")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubResolutionId, "resolution-id", "", "", "Resolution ID")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubResourceArn, "resource-arn", "", "", "Resource ARN")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubResourceGroupNames, "resource-group-names", "", nil, "Resource Group Names")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubResourceMappings, "resource-mappings", "", "", "Resource Mappings")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubResourceName, "resource-name", "", "", "Resource Name")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubResourceNames, "resource-names", "", nil, "Resource Names")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubResourceType, "resource-type", "", "", "Resource Type")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubReverseOrder, "reverse-order", "", "", "Reverse Order")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubSorts, "sorts", "", "", "Sorts")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubSourceArn, "source-arn", "", "", "Source ARN")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubSourceArns, "source-arns", "", nil, "Source Arns")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubStartTime, "start-time", "", "", "Start Time")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubStatus, "status", "", "", "Status")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubTagKeys, "tag-keys", "", nil, "Tag Keys")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubTags, "tags", "", "", "Tags")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubTerraformSource, "terraform-source", "", "", "Terraform Source")
	_resiliencehubCmd.Flags().StringSliceVarP(&_resiliencehubTerraformSourceNames, "terraform-source-names", "", nil, "Terraform Source Names")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubTerraformSources, "terraform-sources", "", "", "Terraform Sources")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubTier, "tier", "", "", "Tier")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubToLastAssessmentTime, "to-last-assessment-time", "", "", "To Last Assessment Time")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubType, "type", "", "", "Type")
	_resiliencehubCmd.Flags().StringVarP(&_resiliencehubVersionName, "version-name", "", "", "Version Name")

	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubAcceptResourceGroupingRecommendations, "accept-resource-grouping-recommendations", "", false, "Accept Resource Grouping Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubAddDraftAppVersionResourceMappings, "add-draft-app-version-resource-mappings", "", false, "Add Draft App Version Resource Mappings")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubBatchUpdateRecommendationStatus, "batch-update-recommendation-status", "", false, "Batch Update Recommendation Status")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubCreateApp, "create-app", "", false, "Create App")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubCreateAppVersionAppComponent, "create-app-version-app-component", "", false, "Create App Version App Component")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubCreateAppVersionResource, "create-app-version-resource", "", false, "Create App Version Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubCreateRecommendationTemplate, "create-recommendation-template", "", false, "Create Recommendation Template")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubCreateResiliencyPolicy, "create-resiliency-policy", "", false, "Create Resiliency Policy")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteApp, "delete-app", "", false, "Delete App")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteAppAssessment, "delete-app-assessment", "", false, "Delete App Assessment")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteAppInputSource, "delete-app-input-source", "", false, "Delete App Input Source")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteAppVersionAppComponent, "delete-app-version-app-component", "", false, "Delete App Version App Component")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteAppVersionResource, "delete-app-version-resource", "", false, "Delete App Version Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteRecommendationTemplate, "delete-recommendation-template", "", false, "Delete Recommendation Template")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDeleteResiliencyPolicy, "delete-resiliency-policy", "", false, "Delete Resiliency Policy")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeApp, "describe-app", "", false, "Describe App")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeAppAssessment, "describe-app-assessment", "", false, "Describe App Assessment")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeAppVersion, "describe-app-version", "", false, "Describe App Version")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeAppVersionAppComponent, "describe-app-version-app-component", "", false, "Describe App Version App Component")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeAppVersionResource, "describe-app-version-resource", "", false, "Describe App Version Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeAppVersionResourcesResolutionStatus, "describe-app-version-resources-resolution-status", "", false, "Describe App Version Resources Resolution Status")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeAppVersionTemplate, "describe-app-version-template", "", false, "Describe App Version Template")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeDraftAppVersionResourcesImportStatus, "describe-draft-app-version-resources-import-status", "", false, "Describe Draft App Version Resources Import Status")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeMetricsExport, "describe-metrics-export", "", false, "Describe Metrics Export")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeResiliencyPolicy, "describe-resiliency-policy", "", false, "Describe Resiliency Policy")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubDescribeResourceGroupingRecommendationTask, "describe-resource-grouping-recommendation-task", "", false, "Describe Resource Grouping Recommendation Task")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubImportResourcesToDraftAppVersion, "import-resources-to-draft-app-version", "", false, "Import Resources To Draft App Version")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAlarmRecommendations, "list-alarm-recommendations", "", false, "List Alarm Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppAssessmentComplianceDrifts, "list-app-assessment-compliance-drifts", "", false, "List App Assessment Compliance Drifts")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppAssessmentResourceDrifts, "list-app-assessment-resource-drifts", "", false, "List App Assessment Resource Drifts")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppAssessments, "list-app-assessments", "", false, "List App Assessments")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppComponentCompliances, "list-app-component-compliances", "", false, "List App Component Compliances")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppComponentRecommendations, "list-app-component-recommendations", "", false, "List App Component Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppInputSources, "list-app-input-sources", "", false, "List App Input Sources")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppVersionAppComponents, "list-app-version-app-components", "", false, "List App Version App Components")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppVersionResourceMappings, "list-app-version-resource-mappings", "", false, "List App Version Resource Mappings")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppVersionResources, "list-app-version-resources", "", false, "List App Version Resources")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListAppVersions, "list-app-versions", "", false, "List App Versions")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListApps, "list-apps", "", false, "List Apps")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListMetrics, "list-metrics", "", false, "List Metrics")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListRecommendationTemplates, "list-recommendation-templates", "", false, "List Recommendation Templates")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListResiliencyPolicies, "list-resiliency-policies", "", false, "List Resiliency Policies")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListResourceGroupingRecommendations, "list-resource-grouping-recommendations", "", false, "List Resource Grouping Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListSopRecommendations, "list-sop-recommendations", "", false, "List Sop Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListSuggestedResiliencyPolicies, "list-suggested-resiliency-policies", "", false, "List Suggested Resiliency Policies")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListTestRecommendations, "list-test-recommendations", "", false, "List Test Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubListUnsupportedAppVersionResources, "list-unsupported-app-version-resources", "", false, "List Unsupported App Version Resources")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubPublishAppVersion, "publish-app-version", "", false, "Publish App Version")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubPutDraftAppVersionTemplate, "put-draft-app-version-template", "", false, "Put Draft App Version Template")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubRejectResourceGroupingRecommendations, "reject-resource-grouping-recommendations", "", false, "Reject Resource Grouping Recommendations")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubRemoveDraftAppVersionResourceMappings, "remove-draft-app-version-resource-mappings", "", false, "Remove Draft App Version Resource Mappings")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubResolveAppVersionResources, "resolve-app-version-resources", "", false, "Resolve App Version Resources")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubStartAppAssessment, "start-app-assessment", "", false, "Start App Assessment")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubStartMetricsExport, "start-metrics-export", "", false, "Start Metrics Export")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubStartResourceGroupingRecommendationTask, "start-resource-grouping-recommendation-task", "", false, "Start Resource Grouping Recommendation Task")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubTagResource, "tag-resource", "", false, "Tag Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubUntagResource, "untag-resource", "", false, "Untag Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubUpdateApp, "update-app", "", false, "Update App")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubUpdateAppVersion, "update-app-version", "", false, "Update App Version")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubUpdateAppVersionAppComponent, "update-app-version-app-component", "", false, "Update App Version App Component")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubUpdateAppVersionResource, "update-app-version-resource", "", false, "Update App Version Resource")
	_resiliencehubCmd.Flags().BoolVarP(&_resiliencehubUpdateResiliencyPolicy, "update-resiliency-policy", "", false, "Update Resiliency Policy")

}
