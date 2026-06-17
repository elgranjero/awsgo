package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iotthingsgraphCmd represents the iotthingsgraph command
var _iotthingsgraphCmd = &cobra.Command{
	Use:   "iotthingsgraph",
	Short: "AWS iotthingsgraph CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iotthingsgraph.NewFromConfig(cfg)
		if _iotthingsgraphAssociateEntityToThing {
			iotthingsgraph_AssociateEntityToThing(cfg, client)
			return
		}
		if _iotthingsgraphCreateFlowTemplate {
			iotthingsgraph_CreateFlowTemplate(cfg, client)
			return
		}
		if _iotthingsgraphCreateSystemInstance {
			iotthingsgraph_CreateSystemInstance(cfg, client)
			return
		}
		if _iotthingsgraphCreateSystemTemplate {
			iotthingsgraph_CreateSystemTemplate(cfg, client)
			return
		}
		if _iotthingsgraphDeleteFlowTemplate {
			iotthingsgraph_DeleteFlowTemplate(cfg, client)
			return
		}
		if _iotthingsgraphDeleteNamespace {
			iotthingsgraph_DeleteNamespace(cfg, client)
			return
		}
		if _iotthingsgraphDeleteSystemInstance {
			iotthingsgraph_DeleteSystemInstance(cfg, client)
			return
		}
		if _iotthingsgraphDeleteSystemTemplate {
			iotthingsgraph_DeleteSystemTemplate(cfg, client)
			return
		}
		if _iotthingsgraphDeploySystemInstance {
			iotthingsgraph_DeploySystemInstance(cfg, client)
			return
		}
		if _iotthingsgraphDeprecateFlowTemplate {
			iotthingsgraph_DeprecateFlowTemplate(cfg, client)
			return
		}
		if _iotthingsgraphDeprecateSystemTemplate {
			iotthingsgraph_DeprecateSystemTemplate(cfg, client)
			return
		}
		if _iotthingsgraphDescribeNamespace {
			iotthingsgraph_DescribeNamespace(cfg, client)
			return
		}
		if _iotthingsgraphDissociateEntityFromThing {
			iotthingsgraph_DissociateEntityFromThing(cfg, client)
			return
		}
		if _iotthingsgraphGetEntities {
			iotthingsgraph_GetEntities(cfg, client)
			return
		}
		if _iotthingsgraphGetFlowTemplate {
			iotthingsgraph_GetFlowTemplate(cfg, client)
			return
		}
		if _iotthingsgraphGetFlowTemplateRevisions {
			iotthingsgraph_GetFlowTemplateRevisions(cfg, client)
			return
		}
		if _iotthingsgraphGetNamespaceDeletionStatus {
			iotthingsgraph_GetNamespaceDeletionStatus(cfg, client)
			return
		}
		if _iotthingsgraphGetSystemInstance {
			iotthingsgraph_GetSystemInstance(cfg, client)
			return
		}
		if _iotthingsgraphGetSystemTemplate {
			iotthingsgraph_GetSystemTemplate(cfg, client)
			return
		}
		if _iotthingsgraphGetSystemTemplateRevisions {
			iotthingsgraph_GetSystemTemplateRevisions(cfg, client)
			return
		}
		if _iotthingsgraphGetUploadStatus {
			iotthingsgraph_GetUploadStatus(cfg, client)
			return
		}
		if _iotthingsgraphListFlowExecutionMessages {
			iotthingsgraph_ListFlowExecutionMessages(cfg, client)
			return
		}
		if _iotthingsgraphListTagsForResource {
			iotthingsgraph_ListTagsForResource(cfg, client)
			return
		}
		if _iotthingsgraphSearchEntities {
			iotthingsgraph_SearchEntities(cfg, client)
			return
		}
		if _iotthingsgraphSearchFlowExecutions {
			iotthingsgraph_SearchFlowExecutions(cfg, client)
			return
		}
		if _iotthingsgraphSearchFlowTemplates {
			iotthingsgraph_SearchFlowTemplates(cfg, client)
			return
		}
		if _iotthingsgraphSearchSystemInstances {
			iotthingsgraph_SearchSystemInstances(cfg, client)
			return
		}
		if _iotthingsgraphSearchSystemTemplates {
			iotthingsgraph_SearchSystemTemplates(cfg, client)
			return
		}
		if _iotthingsgraphSearchThings {
			iotthingsgraph_SearchThings(cfg, client)
			return
		}
		if _iotthingsgraphTagResource {
			iotthingsgraph_TagResource(cfg, client)
			return
		}
		if _iotthingsgraphUndeploySystemInstance {
			iotthingsgraph_UndeploySystemInstance(cfg, client)
			return
		}
		if _iotthingsgraphUntagResource {
			iotthingsgraph_UntagResource(cfg, client)
			return
		}
		if _iotthingsgraphUpdateFlowTemplate {
			iotthingsgraph_UpdateFlowTemplate(cfg, client)
			return
		}
		if _iotthingsgraphUpdateSystemTemplate {
			iotthingsgraph_UpdateSystemTemplate(cfg, client)
			return
		}
		if _iotthingsgraphUploadEntityDefinitions {
			iotthingsgraph_UploadEntityDefinitions(cfg, client)
			return
		}

	},
}

var (
	_iotthingsgraphAssociateEntityToThing     bool
	_iotthingsgraphCreateFlowTemplate         bool
	_iotthingsgraphCreateSystemInstance       bool
	_iotthingsgraphCreateSystemTemplate       bool
	_iotthingsgraphDeleteFlowTemplate         bool
	_iotthingsgraphDeleteNamespace            bool
	_iotthingsgraphDeleteSystemInstance       bool
	_iotthingsgraphDeleteSystemTemplate       bool
	_iotthingsgraphDeploySystemInstance       bool
	_iotthingsgraphDeprecateFlowTemplate      bool
	_iotthingsgraphDeprecateSystemTemplate    bool
	_iotthingsgraphDescribeNamespace          bool
	_iotthingsgraphDissociateEntityFromThing  bool
	_iotthingsgraphGetEntities                bool
	_iotthingsgraphGetFlowTemplate            bool
	_iotthingsgraphGetFlowTemplateRevisions   bool
	_iotthingsgraphGetNamespaceDeletionStatus bool
	_iotthingsgraphGetSystemInstance          bool
	_iotthingsgraphGetSystemTemplate          bool
	_iotthingsgraphGetSystemTemplateRevisions bool
	_iotthingsgraphGetUploadStatus            bool
	_iotthingsgraphListFlowExecutionMessages  bool
	_iotthingsgraphListTagsForResource        bool
	_iotthingsgraphSearchEntities             bool
	_iotthingsgraphSearchFlowExecutions       bool
	_iotthingsgraphSearchFlowTemplates        bool
	_iotthingsgraphSearchSystemInstances      bool
	_iotthingsgraphSearchSystemTemplates      bool
	_iotthingsgraphSearchThings               bool
	_iotthingsgraphTagResource                bool
	_iotthingsgraphUndeploySystemInstance     bool
	_iotthingsgraphUntagResource              bool
	_iotthingsgraphUpdateFlowTemplate         bool
	_iotthingsgraphUpdateSystemTemplate       bool
	_iotthingsgraphUploadEntityDefinitions    bool

	_iotthingsgraphCompatibleNamespaceVersion string
	_iotthingsgraphDefinition                 string
	_iotthingsgraphDeprecateExistingEntities  string
	_iotthingsgraphDocument                   string
	_iotthingsgraphEndTime                    string
	_iotthingsgraphEntityId                   string
	_iotthingsgraphEntityType                 string
	_iotthingsgraphEntityTypes                string
	_iotthingsgraphFilters                    string
	_iotthingsgraphFlowActionsRoleArn         string
	_iotthingsgraphFlowExecutionId            string
	_iotthingsgraphGreengrassGroupName        string
	_iotthingsgraphId                         string
	_iotthingsgraphIds                        []string
	_iotthingsgraphMaxResults                 string
	_iotthingsgraphMetricsConfiguration       string
	_iotthingsgraphNamespaceName              string
	_iotthingsgraphNamespaceVersion           string
	_iotthingsgraphNextToken                  string
	_iotthingsgraphResourceArn                string
	_iotthingsgraphRevisionNumber             string
	_iotthingsgraphS3BucketName               string
	_iotthingsgraphStartTime                  string
	_iotthingsgraphSyncWithPublicNamespace    string
	_iotthingsgraphSystemInstanceId           string
	_iotthingsgraphTagKeys                    []string
	_iotthingsgraphTags                       string
	_iotthingsgraphTarget                     string
	_iotthingsgraphThingName                  string
	_iotthingsgraphUploadId                   string
)

// Associates a device with a concrete thing that is in the user's registry.
// A thing can be associated with only one device at a time. If you associate a
// thing with a new device id, its previous association will be removed.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_AssociateEntityToThing(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.AssociateEntityToThingInput{
		// EntityId: *string, // Required
		// ThingName: *string, // Required
	}

	if len(_iotthingsgraphEntityId) > 0 {
		input.EntityId = aws.String(_iotthingsgraphEntityId)
	}
	if len(_iotthingsgraphThingName) > 0 {
		input.ThingName = aws.String(_iotthingsgraphThingName)
	}
	if len(_iotthingsgraphNamespaceVersion) > 0 {
		if err := assignInputField(input, "NamespaceVersion", _iotthingsgraphNamespaceVersion); err != nil {
			log.Errorf("invalid --namespace-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateEntityToThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a workflow template. Workflows can be created only in the user's
// namespace. (The public namespace contains only entities.) The workflow can
// contain only entities in the specified namespace. The workflow is validated
// against the entities in the latest version of the user's namespace unless
// another namespace version is specified in the request.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_CreateFlowTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.CreateFlowTemplateInput{
		// Definition: *types.DefinitionDocument, // Required
	}

	if len(_iotthingsgraphDefinition) > 0 {
		if err := assignInputField(input, "Definition", _iotthingsgraphDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphCompatibleNamespaceVersion) > 0 {
		if err := assignInputField(input, "CompatibleNamespaceVersion", _iotthingsgraphCompatibleNamespaceVersion); err != nil {
			log.Errorf("invalid --compatible-namespace-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFlowTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a system instance.
// This action validates the system instance, prepares the deployment-related
// resources. For Greengrass deployments, it updates the Greengrass group that is
// specified by the greengrassGroupName parameter. It also adds a file to the S3
// bucket specified by the s3BucketName parameter. You need to call
// DeploySystemInstance after running this action.
//
// For Greengrass deployments, since this action modifies and adds resources to a
// Greengrass group and an S3 bucket on the caller's behalf, the calling identity
// must have write permissions to both the specified Greengrass group and S3
// bucket. Otherwise, the call will fail with an authorization error.
//
// For cloud deployments, this action requires a flowActionsRoleArn value. This is
// an IAM role that has permissions to access AWS services, such as AWS Lambda and
// AWS IoT, that the flow uses when it executes.
//
// If the definition document doesn't specify a version of the user's namespace,
// the latest version will be used by default.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_CreateSystemInstance(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.CreateSystemInstanceInput{
		// Definition: *types.DefinitionDocument, // Required
		// Target: types.DeploymentTarget, // Required
	}

	if len(_iotthingsgraphDefinition) > 0 {
		if err := assignInputField(input, "Definition", _iotthingsgraphDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphTarget) > 0 {
		if err := assignInputField(input, "Target", _iotthingsgraphTarget); err != nil {
			log.Errorf("invalid --target: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphFlowActionsRoleArn) > 0 {
		input.FlowActionsRoleArn = aws.String(_iotthingsgraphFlowActionsRoleArn)
	}
	if len(_iotthingsgraphGreengrassGroupName) > 0 {
		input.GreengrassGroupName = aws.String(_iotthingsgraphGreengrassGroupName)
	}
	if len(_iotthingsgraphMetricsConfiguration) > 0 {
		if err := assignInputField(input, "MetricsConfiguration", _iotthingsgraphMetricsConfiguration); err != nil {
			log.Errorf("invalid --metrics-configuration: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphS3BucketName) > 0 {
		input.S3BucketName = aws.String(_iotthingsgraphS3BucketName)
	}
	if len(_iotthingsgraphTags) > 0 {
		if err := assignInputField(input, "Tags", _iotthingsgraphTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSystemInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a system. The system is validated against the entities in the latest
// version of the user's namespace unless another namespace version is specified in
// the request.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_CreateSystemTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.CreateSystemTemplateInput{
		// Definition: *types.DefinitionDocument, // Required
	}

	if len(_iotthingsgraphDefinition) > 0 {
		if err := assignInputField(input, "Definition", _iotthingsgraphDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphCompatibleNamespaceVersion) > 0 {
		if err := assignInputField(input, "CompatibleNamespaceVersion", _iotthingsgraphCompatibleNamespaceVersion); err != nil {
			log.Errorf("invalid --compatible-namespace-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSystemTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workflow. Any new system or deployment that contains this workflow
// will fail to update or deploy. Existing deployments that contain the workflow
// will continue to run (since they use a snapshot of the workflow taken at the
// time of deployment).
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DeleteFlowTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeleteFlowTemplateInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.DeleteFlowTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified namespace. This action deletes all of the entities in the
// namespace. Delete the systems and flows that use entities in the namespace
// before performing this action. This action takes no request parameters.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DeleteNamespace(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeleteNamespaceInput{}

	if resp, err := client.DeleteNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a system instance. Only system instances that have never been deployed,
// or that have been undeployed can be deleted.
//
// Users can create a new system instance that has the same ID as a deleted system
// instance.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DeleteSystemInstance(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeleteSystemInstanceInput{}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.DeleteSystemInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a system. New deployments can't contain the system after its deletion.
// Existing deployments that contain the system will continue to work because they
// use a snapshot of the system that is taken when it is deployed.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DeleteSystemTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeleteSystemTemplateInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.DeleteSystemTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Greengrass and Cloud Deployments
// Deploys the system instance to the target specified in CreateSystemInstance .
//
// # Greengrass Deployments
//
// If the system or any workflows and entities have been updated before this
// action is called, then the deployment will create a new Amazon Simple Storage
// Service resource file and then deploy it.
//
// Since this action creates a Greengrass deployment on the caller's behalf, the
// calling identity must have write permissions to the specified Greengrass group.
// Otherwise, the call will fail with an authorization error.
//
// For information about the artifacts that get added to your Greengrass core
// device when you use this API, see [AWS IoT Things Graph and AWS IoT Greengrass].
//
// Deprecated: since: 2022-08-30
//
// [AWS IoT Things Graph and AWS IoT Greengrass]: https://docs.aws.amazon.com/thingsgraph/latest/ug/iot-tg-greengrass.html
func iotthingsgraph_DeploySystemInstance(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeploySystemInstanceInput{}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.DeploySystemInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecates the specified workflow. This action marks the workflow for deletion.
// Deprecated flows can't be deployed, but existing deployments will continue to
// run.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DeprecateFlowTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeprecateFlowTemplateInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.DeprecateFlowTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deprecates the specified system.
// Deprecated: since: 2022-08-30
func iotthingsgraph_DeprecateSystemTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DeprecateSystemTemplateInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.DeprecateSystemTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the latest version of the user's namespace and the public version that it
// is tracking.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DescribeNamespace(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DescribeNamespaceInput{}

	if len(_iotthingsgraphNamespaceName) > 0 {
		input.NamespaceName = aws.String(_iotthingsgraphNamespaceName)
	}

	if resp, err := client.DescribeNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dissociates a device entity from a concrete thing. The action takes only the
// type of the entity that you need to dissociate because only one entity of a
// particular type can be associated with a thing.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_DissociateEntityFromThing(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.DissociateEntityFromThingInput{
		// EntityType: types.EntityType, // Required
		// ThingName: *string, // Required
	}

	if len(_iotthingsgraphEntityType) > 0 {
		if err := assignInputField(input, "EntityType", _iotthingsgraphEntityType); err != nil {
			log.Errorf("invalid --entity-type: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphThingName) > 0 {
		input.ThingName = aws.String(_iotthingsgraphThingName)
	}

	if resp, err := client.DissociateEntityFromThing(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets definitions of the specified entities. Uses the latest version of the
// user's namespace by default. This API returns the following TDM entities.
//
// - Properties
//
// - States
//
// - Events
//
// - Actions
//
// - Capabilities
//
// - Mappings
//
// - Devices
//
// - Device Models
//
// - Services
//
// This action doesn't return definitions for systems, flows, and deployments.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetEntities(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetEntitiesInput{
		// Ids: []string, // Required
	}

	if len(_iotthingsgraphIds) > 0 {
		input.Ids = append([]string(nil), _iotthingsgraphIds...)
	}
	if len(_iotthingsgraphNamespaceVersion) > 0 {
		if err := assignInputField(input, "NamespaceVersion", _iotthingsgraphNamespaceVersion); err != nil {
			log.Errorf("invalid --namespace-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetEntities(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the latest version of the DefinitionDocument and FlowTemplateSummary for
// the specified workflow.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetFlowTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetFlowTemplateInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}
	if len(_iotthingsgraphRevisionNumber) > 0 {
		if err := assignInputField(input, "RevisionNumber", _iotthingsgraphRevisionNumber); err != nil {
			log.Errorf("invalid --revision-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetFlowTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets revisions of the specified workflow. Only the last 100 revisions are
// stored. If the workflow has been deprecated, this action will return revisions
// that occurred before the deprecation. This action won't work for workflows that
// have been deleted.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetFlowTemplateRevisions(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetFlowTemplateRevisionsInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetFlowTemplateRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.GetFlowTemplateRevisionsOutput
	p := iotthingsgraph.NewGetFlowTemplateRevisionsPaginator(client, input)
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

// Gets the status of a namespace deletion task.
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetNamespaceDeletionStatus(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetNamespaceDeletionStatusInput{}

	if resp, err := client.GetNamespaceDeletionStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a system instance.
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetSystemInstance(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetSystemInstanceInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.GetSystemInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a system.
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetSystemTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetSystemTemplateInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}
	if len(_iotthingsgraphRevisionNumber) > 0 {
		if err := assignInputField(input, "RevisionNumber", _iotthingsgraphRevisionNumber); err != nil {
			log.Errorf("invalid --revision-number: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSystemTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets revisions made to the specified system template. Only the previous 100
// revisions are stored. If the system has been deprecated, this action will return
// the revisions that occurred before its deprecation. This action won't work with
// systems that have been deleted.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetSystemTemplateRevisions(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetSystemTemplateRevisionsInput{
		// Id: *string, // Required
	}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetSystemTemplateRevisions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.GetSystemTemplateRevisionsOutput
	p := iotthingsgraph.NewGetSystemTemplateRevisionsPaginator(client, input)
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

// Gets the status of the specified upload.
// Deprecated: since: 2022-08-30
func iotthingsgraph_GetUploadStatus(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.GetUploadStatusInput{
		// UploadId: *string, // Required
	}

	if len(_iotthingsgraphUploadId) > 0 {
		input.UploadId = aws.String(_iotthingsgraphUploadId)
	}

	if resp, err := client.GetUploadStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of objects that contain information about events in a flow
// execution.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_ListFlowExecutionMessages(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.ListFlowExecutionMessagesInput{
		// FlowExecutionId: *string, // Required
	}

	if len(_iotthingsgraphFlowExecutionId) > 0 {
		input.FlowExecutionId = aws.String(_iotthingsgraphFlowExecutionId)
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListFlowExecutionMessages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.ListFlowExecutionMessagesOutput
	p := iotthingsgraph.NewListFlowExecutionMessagesPaginator(client, input)
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

// Lists all tags on an AWS IoT Things Graph resource.
// Deprecated: since: 2022-08-30
func iotthingsgraph_ListTagsForResource(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_iotthingsgraphResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotthingsgraphResourceArn)
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.ListTagsForResourceOutput
	p := iotthingsgraph.NewListTagsForResourcePaginator(client, input)
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

// Searches for entities of the specified type. You can search for entities in
// your namespace and the public namespace that you're tracking.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_SearchEntities(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.SearchEntitiesInput{
		// EntityTypes: []types.EntityType, // Required
	}

	if len(_iotthingsgraphEntityTypes) > 0 {
		if err := assignInputField(input, "EntityTypes", _iotthingsgraphEntityTypes); err != nil {
			log.Errorf("invalid --entity-types: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphFilters) > 0 {
		if err := assignInputField(input, "Filters", _iotthingsgraphFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNamespaceVersion) > 0 {
		if err := assignInputField(input, "NamespaceVersion", _iotthingsgraphNamespaceVersion); err != nil {
			log.Errorf("invalid --namespace-version: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.SearchEntitiesOutput
	p := iotthingsgraph.NewSearchEntitiesPaginator(client, input)
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

// Searches for AWS IoT Things Graph workflow execution instances.
// Deprecated: since: 2022-08-30
func iotthingsgraph_SearchFlowExecutions(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.SearchFlowExecutionsInput{
		// SystemInstanceId: *string, // Required
	}

	if len(_iotthingsgraphSystemInstanceId) > 0 {
		input.SystemInstanceId = aws.String(_iotthingsgraphSystemInstanceId)
	}
	if len(_iotthingsgraphEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _iotthingsgraphEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphFlowExecutionId) > 0 {
		input.FlowExecutionId = aws.String(_iotthingsgraphFlowExecutionId)
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}
	if len(_iotthingsgraphStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _iotthingsgraphStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchFlowExecutions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.SearchFlowExecutionsOutput
	p := iotthingsgraph.NewSearchFlowExecutionsPaginator(client, input)
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

// Searches for summary information about workflows.
// Deprecated: since: 2022-08-30
func iotthingsgraph_SearchFlowTemplates(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.SearchFlowTemplatesInput{}

	if len(_iotthingsgraphFilters) > 0 {
		if err := assignInputField(input, "Filters", _iotthingsgraphFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchFlowTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.SearchFlowTemplatesOutput
	p := iotthingsgraph.NewSearchFlowTemplatesPaginator(client, input)
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

// Searches for system instances in the user's account.
// Deprecated: since: 2022-08-30
func iotthingsgraph_SearchSystemInstances(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.SearchSystemInstancesInput{}

	if len(_iotthingsgraphFilters) > 0 {
		if err := assignInputField(input, "Filters", _iotthingsgraphFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchSystemInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.SearchSystemInstancesOutput
	p := iotthingsgraph.NewSearchSystemInstancesPaginator(client, input)
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

// Searches for summary information about systems in the user's account. You can
// filter by the ID of a workflow to return only systems that use the specified
// workflow.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_SearchSystemTemplates(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.SearchSystemTemplatesInput{}

	if len(_iotthingsgraphFilters) > 0 {
		if err := assignInputField(input, "Filters", _iotthingsgraphFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchSystemTemplates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.SearchSystemTemplatesOutput
	p := iotthingsgraph.NewSearchSystemTemplatesPaginator(client, input)
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

// Searches for things associated with the specified entity. You can search by
// both device and device model.
//
// For example, if two different devices, camera1 and camera2, implement the
// camera device model, the user can associate thing1 to camera1 and thing2 to
// camera2. SearchThings(camera2) will return only thing2, but SearchThings(camera)
// will return both thing1 and thing2.
//
// This action searches for exact matches and doesn't perform partial text
// matching.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_SearchThings(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.SearchThingsInput{
		// EntityId: *string, // Required
	}

	if len(_iotthingsgraphEntityId) > 0 {
		input.EntityId = aws.String(_iotthingsgraphEntityId)
	}
	if len(_iotthingsgraphMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iotthingsgraphMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNamespaceVersion) > 0 {
		if err := assignInputField(input, "NamespaceVersion", _iotthingsgraphNamespaceVersion); err != nil {
			log.Errorf("invalid --namespace-version: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphNextToken) > 0 {
		input.NextToken = aws.String(_iotthingsgraphNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchThings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iotthingsgraph.SearchThingsOutput
	p := iotthingsgraph.NewSearchThingsPaginator(client, input)
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

// Creates a tag for the specified resource.
// Deprecated: since: 2022-08-30
func iotthingsgraph_TagResource(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_iotthingsgraphResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotthingsgraphResourceArn)
	}
	if len(_iotthingsgraphTags) > 0 {
		if err := assignInputField(input, "Tags", _iotthingsgraphTags); err != nil {
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

// Removes a system instance from its target (Cloud or Greengrass).
// Deprecated: since: 2022-08-30
func iotthingsgraph_UndeploySystemInstance(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.UndeploySystemInstanceInput{}

	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}

	if resp, err := client.UndeploySystemInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a tag from the specified resource.
// Deprecated: since: 2022-08-30
func iotthingsgraph_UntagResource(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iotthingsgraphResourceArn) > 0 {
		input.ResourceArn = aws.String(_iotthingsgraphResourceArn)
	}
	if len(_iotthingsgraphTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iotthingsgraphTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified workflow. All deployed systems and system instances that
// use the workflow will see the changes in the flow when it is redeployed. If you
// don't want this behavior, copy the workflow (creating a new workflow with a
// different ID), and update the copy. The workflow can contain only entities in
// the specified namespace.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_UpdateFlowTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.UpdateFlowTemplateInput{
		// Definition: *types.DefinitionDocument, // Required
		// Id: *string, // Required
	}

	if len(_iotthingsgraphDefinition) > 0 {
		if err := assignInputField(input, "Definition", _iotthingsgraphDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}
	if len(_iotthingsgraphCompatibleNamespaceVersion) > 0 {
		if err := assignInputField(input, "CompatibleNamespaceVersion", _iotthingsgraphCompatibleNamespaceVersion); err != nil {
			log.Errorf("invalid --compatible-namespace-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFlowTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified system. You don't need to run this action after updating
// a workflow. Any deployment that uses the system will see the changes in the
// system when it is redeployed.
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_UpdateSystemTemplate(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.UpdateSystemTemplateInput{
		// Definition: *types.DefinitionDocument, // Required
		// Id: *string, // Required
	}

	if len(_iotthingsgraphDefinition) > 0 {
		if err := assignInputField(input, "Definition", _iotthingsgraphDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphId) > 0 {
		input.Id = aws.String(_iotthingsgraphId)
	}
	if len(_iotthingsgraphCompatibleNamespaceVersion) > 0 {
		if err := assignInputField(input, "CompatibleNamespaceVersion", _iotthingsgraphCompatibleNamespaceVersion); err != nil {
			log.Errorf("invalid --compatible-namespace-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSystemTemplate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Asynchronously uploads one or more entity definitions to the user's namespace.
// The document parameter is required if syncWithPublicNamespace and
// deleteExistingEntites are false. If the syncWithPublicNamespace parameter is
// set to true , the user's namespace will synchronize with the latest version of
// the public namespace. If deprecateExistingEntities is set to true, all entities
// in the latest version will be deleted before the new DefinitionDocument is
// uploaded.
//
// When a user uploads entity definitions for the first time, the service creates
// a new namespace for the user. The new namespace tracks the public namespace.
// Currently users can have only one namespace. The namespace version increments
// whenever a user uploads entity definitions that are backwards-incompatible and
// whenever a user sets the syncWithPublicNamespace parameter or the
// deprecateExistingEntities parameter to true .
//
// The IDs for all of the entities should be in URN format. Each entity must be in
// the user's namespace. Users can't create entities in the public namespace, but
// entity definitions can refer to entities in the public namespace.
//
// Valid entities are Device , DeviceModel , Service , Capability , State , Action
// , Event , Property , Mapping , Enum .
//
// Deprecated: since: 2022-08-30
func iotthingsgraph_UploadEntityDefinitions(cfg aws.Config, client *iotthingsgraph.Client) {
	input := &iotthingsgraph.UploadEntityDefinitionsInput{}

	if len(_iotthingsgraphDeprecateExistingEntities) > 0 {
		if err := assignInputField(input, "DeprecateExistingEntities", _iotthingsgraphDeprecateExistingEntities); err != nil {
			log.Errorf("invalid --deprecate-existing-entities: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphDocument) > 0 {
		if err := assignInputField(input, "Document", _iotthingsgraphDocument); err != nil {
			log.Errorf("invalid --document: %s", err.Error())
			return
		}
	}
	if len(_iotthingsgraphSyncWithPublicNamespace) > 0 {
		if err := assignInputField(input, "SyncWithPublicNamespace", _iotthingsgraphSyncWithPublicNamespace); err != nil {
			log.Errorf("invalid --sync-with-public-namespace: %s", err.Error())
			return
		}
	}

	if resp, err := client.UploadEntityDefinitions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iotthingsgraphCmd)
	_iotthingsgraphCmd.Flags().SortFlags = false

	_iotthingsgraphCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_iotthingsgraphCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iotthingsgraphCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphCompatibleNamespaceVersion, "compatible-namespace-version", "", "", "Compatible Namespace Version")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphDefinition, "definition", "", "", "Definition")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphDeprecateExistingEntities, "deprecate-existing-entities", "", "", "Deprecate Existing Entities")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphDocument, "document", "", "", "Document")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphEndTime, "end-time", "", "", "End Time")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphEntityId, "entity-id", "", "", "Entity ID")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphEntityType, "entity-type", "", "", "Entity Type")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphEntityTypes, "entity-types", "", "", "Entity Types")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphFilters, "filters", "", "", "Filters")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphFlowActionsRoleArn, "flow-actions-role-arn", "", "", "Flow Actions Role ARN")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphFlowExecutionId, "flow-execution-id", "", "", "Flow Execution ID")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphGreengrassGroupName, "greengrass-group-name", "", "", "Greengrass Group Name")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphId, "id", "", "", "ID")
	_iotthingsgraphCmd.Flags().StringSliceVarP(&_iotthingsgraphIds, "ids", "", nil, "Ids")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphMaxResults, "max-results", "", "", "Max Results")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphMetricsConfiguration, "metrics-configuration", "", "", "Metrics Configuration")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphNamespaceName, "namespace-name", "", "", "Namespace Name")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphNamespaceVersion, "namespace-version", "", "", "Namespace Version")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphNextToken, "next-token", "", "", "Next Token")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphResourceArn, "resource-arn", "", "", "Resource ARN")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphRevisionNumber, "revision-number", "", "", "Revision Number")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphS3BucketName, "s3-bucket-name", "", "", "S3 Bucket Name")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphStartTime, "start-time", "", "", "Start Time")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphSyncWithPublicNamespace, "sync-with-public-namespace", "", "", "Sync With Public Namespace")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphSystemInstanceId, "system-instance-id", "", "", "System Instance ID")
	_iotthingsgraphCmd.Flags().StringSliceVarP(&_iotthingsgraphTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphTags, "tags", "", "", "Tags")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphTarget, "target", "", "", "Target")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphThingName, "thing-name", "", "", "Thing Name")
	_iotthingsgraphCmd.Flags().StringVarP(&_iotthingsgraphUploadId, "upload-id", "", "", "Upload ID")

	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphAssociateEntityToThing, "associate-entity-to-thing", "", false, "Associate Entity To Thing")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphCreateFlowTemplate, "create-flow-template", "", false, "Create Flow Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphCreateSystemInstance, "create-system-instance", "", false, "Create System Instance")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphCreateSystemTemplate, "create-system-template", "", false, "Create System Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeleteFlowTemplate, "delete-flow-template", "", false, "Delete Flow Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeleteNamespace, "delete-namespace", "", false, "Delete Namespace")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeleteSystemInstance, "delete-system-instance", "", false, "Delete System Instance")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeleteSystemTemplate, "delete-system-template", "", false, "Delete System Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeploySystemInstance, "deploy-system-instance", "", false, "Deploy System Instance")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeprecateFlowTemplate, "deprecate-flow-template", "", false, "Deprecate Flow Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDeprecateSystemTemplate, "deprecate-system-template", "", false, "Deprecate System Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDescribeNamespace, "describe-namespace", "", false, "Describe Namespace")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphDissociateEntityFromThing, "dissociate-entity-from-thing", "", false, "Dissociate Entity From Thing")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetEntities, "get-entities", "", false, "Get Entities")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetFlowTemplate, "get-flow-template", "", false, "Get Flow Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetFlowTemplateRevisions, "get-flow-template-revisions", "", false, "Get Flow Template Revisions")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetNamespaceDeletionStatus, "get-namespace-deletion-status", "", false, "Get Namespace Deletion Status")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetSystemInstance, "get-system-instance", "", false, "Get System Instance")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetSystemTemplate, "get-system-template", "", false, "Get System Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetSystemTemplateRevisions, "get-system-template-revisions", "", false, "Get System Template Revisions")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphGetUploadStatus, "get-upload-status", "", false, "Get Upload Status")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphListFlowExecutionMessages, "list-flow-execution-messages", "", false, "List Flow Execution Messages")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphSearchEntities, "search-entities", "", false, "Search Entities")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphSearchFlowExecutions, "search-flow-executions", "", false, "Search Flow Executions")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphSearchFlowTemplates, "search-flow-templates", "", false, "Search Flow Templates")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphSearchSystemInstances, "search-system-instances", "", false, "Search System Instances")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphSearchSystemTemplates, "search-system-templates", "", false, "Search System Templates")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphSearchThings, "search-things", "", false, "Search Things")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphTagResource, "tag-resource", "", false, "Tag Resource")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphUndeploySystemInstance, "undeploy-system-instance", "", false, "Undeploy System Instance")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphUntagResource, "untag-resource", "", false, "Untag Resource")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphUpdateFlowTemplate, "update-flow-template", "", false, "Update Flow Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphUpdateSystemTemplate, "update-system-template", "", false, "Update System Template")
	_iotthingsgraphCmd.Flags().BoolVarP(&_iotthingsgraphUploadEntityDefinitions, "upload-entity-definitions", "", false, "Upload Entity Definitions")

}
