package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iottwinmaker"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// iottwinmakerCmd represents the iottwinmaker command
var _iottwinmakerCmd = &cobra.Command{
	Use:   "iottwinmaker",
	Short: "AWS iottwinmaker CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := iottwinmaker.NewFromConfig(cfg)
		if _iottwinmakerBatchPutPropertyValues {
			iottwinmaker_BatchPutPropertyValues(cfg, client)
			return
		}
		if _iottwinmakerCancelMetadataTransferJob {
			iottwinmaker_CancelMetadataTransferJob(cfg, client)
			return
		}
		if _iottwinmakerCreateComponentType {
			iottwinmaker_CreateComponentType(cfg, client)
			return
		}
		if _iottwinmakerCreateEntity {
			iottwinmaker_CreateEntity(cfg, client)
			return
		}
		if _iottwinmakerCreateMetadataTransferJob {
			iottwinmaker_CreateMetadataTransferJob(cfg, client)
			return
		}
		if _iottwinmakerCreateScene {
			iottwinmaker_CreateScene(cfg, client)
			return
		}
		if _iottwinmakerCreateSyncJob {
			iottwinmaker_CreateSyncJob(cfg, client)
			return
		}
		if _iottwinmakerCreateWorkspace {
			iottwinmaker_CreateWorkspace(cfg, client)
			return
		}
		if _iottwinmakerDeleteComponentType {
			iottwinmaker_DeleteComponentType(cfg, client)
			return
		}
		if _iottwinmakerDeleteEntity {
			iottwinmaker_DeleteEntity(cfg, client)
			return
		}
		if _iottwinmakerDeleteScene {
			iottwinmaker_DeleteScene(cfg, client)
			return
		}
		if _iottwinmakerDeleteSyncJob {
			iottwinmaker_DeleteSyncJob(cfg, client)
			return
		}
		if _iottwinmakerDeleteWorkspace {
			iottwinmaker_DeleteWorkspace(cfg, client)
			return
		}
		if _iottwinmakerExecuteQuery {
			iottwinmaker_ExecuteQuery(cfg, client)
			return
		}
		if _iottwinmakerGetComponentType {
			iottwinmaker_GetComponentType(cfg, client)
			return
		}
		if _iottwinmakerGetEntity {
			iottwinmaker_GetEntity(cfg, client)
			return
		}
		if _iottwinmakerGetMetadataTransferJob {
			iottwinmaker_GetMetadataTransferJob(cfg, client)
			return
		}
		if _iottwinmakerGetPricingPlan {
			iottwinmaker_GetPricingPlan(cfg, client)
			return
		}
		if _iottwinmakerGetPropertyValue {
			iottwinmaker_GetPropertyValue(cfg, client)
			return
		}
		if _iottwinmakerGetPropertyValueHistory {
			iottwinmaker_GetPropertyValueHistory(cfg, client)
			return
		}
		if _iottwinmakerGetScene {
			iottwinmaker_GetScene(cfg, client)
			return
		}
		if _iottwinmakerGetSyncJob {
			iottwinmaker_GetSyncJob(cfg, client)
			return
		}
		if _iottwinmakerGetWorkspace {
			iottwinmaker_GetWorkspace(cfg, client)
			return
		}
		if _iottwinmakerListComponentTypes {
			iottwinmaker_ListComponentTypes(cfg, client)
			return
		}
		if _iottwinmakerListComponents {
			iottwinmaker_ListComponents(cfg, client)
			return
		}
		if _iottwinmakerListEntities {
			iottwinmaker_ListEntities(cfg, client)
			return
		}
		if _iottwinmakerListMetadataTransferJobs {
			iottwinmaker_ListMetadataTransferJobs(cfg, client)
			return
		}
		if _iottwinmakerListProperties {
			iottwinmaker_ListProperties(cfg, client)
			return
		}
		if _iottwinmakerListScenes {
			iottwinmaker_ListScenes(cfg, client)
			return
		}
		if _iottwinmakerListSyncJobs {
			iottwinmaker_ListSyncJobs(cfg, client)
			return
		}
		if _iottwinmakerListSyncResources {
			iottwinmaker_ListSyncResources(cfg, client)
			return
		}
		if _iottwinmakerListTagsForResource {
			iottwinmaker_ListTagsForResource(cfg, client)
			return
		}
		if _iottwinmakerListWorkspaces {
			iottwinmaker_ListWorkspaces(cfg, client)
			return
		}
		if _iottwinmakerTagResource {
			iottwinmaker_TagResource(cfg, client)
			return
		}
		if _iottwinmakerUntagResource {
			iottwinmaker_UntagResource(cfg, client)
			return
		}
		if _iottwinmakerUpdateComponentType {
			iottwinmaker_UpdateComponentType(cfg, client)
			return
		}
		if _iottwinmakerUpdateEntity {
			iottwinmaker_UpdateEntity(cfg, client)
			return
		}
		if _iottwinmakerUpdatePricingPlan {
			iottwinmaker_UpdatePricingPlan(cfg, client)
			return
		}
		if _iottwinmakerUpdateScene {
			iottwinmaker_UpdateScene(cfg, client)
			return
		}
		if _iottwinmakerUpdateWorkspace {
			iottwinmaker_UpdateWorkspace(cfg, client)
			return
		}

	},
}

var (
	_iottwinmakerBatchPutPropertyValues    bool
	_iottwinmakerCancelMetadataTransferJob bool
	_iottwinmakerCreateComponentType       bool
	_iottwinmakerCreateEntity              bool
	_iottwinmakerCreateMetadataTransferJob bool
	_iottwinmakerCreateScene               bool
	_iottwinmakerCreateSyncJob             bool
	_iottwinmakerCreateWorkspace           bool
	_iottwinmakerDeleteComponentType       bool
	_iottwinmakerDeleteEntity              bool
	_iottwinmakerDeleteScene               bool
	_iottwinmakerDeleteSyncJob             bool
	_iottwinmakerDeleteWorkspace           bool
	_iottwinmakerExecuteQuery              bool
	_iottwinmakerGetComponentType          bool
	_iottwinmakerGetEntity                 bool
	_iottwinmakerGetMetadataTransferJob    bool
	_iottwinmakerGetPricingPlan            bool
	_iottwinmakerGetPropertyValue          bool
	_iottwinmakerGetPropertyValueHistory   bool
	_iottwinmakerGetScene                  bool
	_iottwinmakerGetSyncJob                bool
	_iottwinmakerGetWorkspace              bool
	_iottwinmakerListComponentTypes        bool
	_iottwinmakerListComponents            bool
	_iottwinmakerListEntities              bool
	_iottwinmakerListMetadataTransferJobs  bool
	_iottwinmakerListProperties            bool
	_iottwinmakerListScenes                bool
	_iottwinmakerListSyncJobs              bool
	_iottwinmakerListSyncResources         bool
	_iottwinmakerListTagsForResource       bool
	_iottwinmakerListWorkspaces            bool
	_iottwinmakerTagResource               bool
	_iottwinmakerUntagResource             bool
	_iottwinmakerUpdateComponentType       bool
	_iottwinmakerUpdateEntity              bool
	_iottwinmakerUpdatePricingPlan         bool
	_iottwinmakerUpdateScene               bool
	_iottwinmakerUpdateWorkspace           bool

	_iottwinmakerBundleNames               []string
	_iottwinmakerCapabilities              []string
	_iottwinmakerComponentName             string
	_iottwinmakerComponentPath             string
	_iottwinmakerComponentTypeId           string
	_iottwinmakerComponentTypeName         string
	_iottwinmakerComponentUpdates          string
	_iottwinmakerComponents                string
	_iottwinmakerCompositeComponentTypes   string
	_iottwinmakerCompositeComponentUpdates string
	_iottwinmakerCompositeComponents       string
	_iottwinmakerContentLocation           string
	_iottwinmakerDescription               string
	_iottwinmakerDestination               string
	_iottwinmakerDestinationType           string
	_iottwinmakerEndDateTime               string
	_iottwinmakerEndTime                   string
	_iottwinmakerEntityId                  string
	_iottwinmakerEntityName                string
	_iottwinmakerEntries                   string
	_iottwinmakerExtendsFrom               []string
	_iottwinmakerFilters                   string
	_iottwinmakerFunctions                 string
	_iottwinmakerInterpolation             string
	_iottwinmakerIsRecursive               string
	_iottwinmakerIsSingleton               string
	_iottwinmakerMaxResults                string
	_iottwinmakerMetadataTransferJobId     string
	_iottwinmakerNextToken                 string
	_iottwinmakerOrderByTime               string
	_iottwinmakerParentEntityId            string
	_iottwinmakerParentEntityUpdate        string
	_iottwinmakerPricingMode               string
	_iottwinmakerPropertyDefinitions       string
	_iottwinmakerPropertyFilters           string
	_iottwinmakerPropertyGroupName         string
	_iottwinmakerPropertyGroups            string
	_iottwinmakerQueryStatement            string
	_iottwinmakerResourceARN               string
	_iottwinmakerRole                      string
	_iottwinmakerS3Location                string
	_iottwinmakerSceneId                   string
	_iottwinmakerSceneMetadata             string
	_iottwinmakerSelectedProperties        []string
	_iottwinmakerSourceType                string
	_iottwinmakerSources                   string
	_iottwinmakerStartDateTime             string
	_iottwinmakerStartTime                 string
	_iottwinmakerSyncRole                  string
	_iottwinmakerSyncSource                string
	_iottwinmakerTabularConditions         string
	_iottwinmakerTagKeys                   []string
	_iottwinmakerTags                      string
	_iottwinmakerWorkspaceId               string
)

// Sets values for multiple time series properties.
func iottwinmaker_BatchPutPropertyValues(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.BatchPutPropertyValuesInput{
		// Entries: []types.PropertyValueEntry, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntries) > 0 {
		if err := assignInputField(input, "Entries", _iottwinmakerEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.BatchPutPropertyValues(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels the metadata transfer job.
func iottwinmaker_CancelMetadataTransferJob(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CancelMetadataTransferJobInput{
		// MetadataTransferJobId: *string, // Required
	}

	if len(_iottwinmakerMetadataTransferJobId) > 0 {
		input.MetadataTransferJobId = aws.String(_iottwinmakerMetadataTransferJobId)
	}

	if resp, err := client.CancelMetadataTransferJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a component type.
func iottwinmaker_CreateComponentType(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CreateComponentTypeInput{
		// ComponentTypeId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerComponentTypeId) > 0 {
		input.ComponentTypeId = aws.String(_iottwinmakerComponentTypeId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentTypeName) > 0 {
		input.ComponentTypeName = aws.String(_iottwinmakerComponentTypeName)
	}
	if len(_iottwinmakerCompositeComponentTypes) > 0 {
		if err := assignInputField(input, "CompositeComponentTypes", _iottwinmakerCompositeComponentTypes); err != nil {
			log.Errorf("invalid --composite-component-types: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerExtendsFrom) > 0 {
		input.ExtendsFrom = append([]string(nil), _iottwinmakerExtendsFrom...)
	}
	if len(_iottwinmakerFunctions) > 0 {
		if err := assignInputField(input, "Functions", _iottwinmakerFunctions); err != nil {
			log.Errorf("invalid --functions: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerIsSingleton) > 0 {
		if err := assignInputField(input, "IsSingleton", _iottwinmakerIsSingleton); err != nil {
			log.Errorf("invalid --is-singleton: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerPropertyDefinitions) > 0 {
		if err := assignInputField(input, "PropertyDefinitions", _iottwinmakerPropertyDefinitions); err != nil {
			log.Errorf("invalid --property-definitions: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerPropertyGroups) > 0 {
		if err := assignInputField(input, "PropertyGroups", _iottwinmakerPropertyGroups); err != nil {
			log.Errorf("invalid --property-groups: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerTags) > 0 {
		if err := assignInputField(input, "Tags", _iottwinmakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateComponentType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an entity.
func iottwinmaker_CreateEntity(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CreateEntityInput{
		// EntityName: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntityName) > 0 {
		input.EntityName = aws.String(_iottwinmakerEntityName)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponents) > 0 {
		if err := assignInputField(input, "Components", _iottwinmakerComponents); err != nil {
			log.Errorf("invalid --components: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerCompositeComponents) > 0 {
		if err := assignInputField(input, "CompositeComponents", _iottwinmakerCompositeComponents); err != nil {
			log.Errorf("invalid --composite-components: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerParentEntityId) > 0 {
		input.ParentEntityId = aws.String(_iottwinmakerParentEntityId)
	}
	if len(_iottwinmakerTags) > 0 {
		if err := assignInputField(input, "Tags", _iottwinmakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new metadata transfer job.
func iottwinmaker_CreateMetadataTransferJob(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CreateMetadataTransferJobInput{
		// Destination: *types.DestinationConfiguration, // Required
		// Sources: []types.SourceConfiguration, // Required
	}

	if len(_iottwinmakerDestination) > 0 {
		if err := assignInputField(input, "Destination", _iottwinmakerDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerSources) > 0 {
		if err := assignInputField(input, "Sources", _iottwinmakerSources); err != nil {
			log.Errorf("invalid --sources: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerMetadataTransferJobId) > 0 {
		input.MetadataTransferJobId = aws.String(_iottwinmakerMetadataTransferJobId)
	}

	if resp, err := client.CreateMetadataTransferJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a scene.
func iottwinmaker_CreateScene(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CreateSceneInput{
		// ContentLocation: *string, // Required
		// SceneId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerContentLocation) > 0 {
		input.ContentLocation = aws.String(_iottwinmakerContentLocation)
	}
	if len(_iottwinmakerSceneId) > 0 {
		input.SceneId = aws.String(_iottwinmakerSceneId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerCapabilities) > 0 {
		input.Capabilities = append([]string(nil), _iottwinmakerCapabilities...)
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerSceneMetadata) > 0 {
		if err := assignInputField(input, "SceneMetadata", _iottwinmakerSceneMetadata); err != nil {
			log.Errorf("invalid --scene-metadata: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerTags) > 0 {
		if err := assignInputField(input, "Tags", _iottwinmakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScene(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action creates a SyncJob.
func iottwinmaker_CreateSyncJob(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CreateSyncJobInput{
		// SyncRole: *string, // Required
		// SyncSource: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSyncRole) > 0 {
		input.SyncRole = aws.String(_iottwinmakerSyncRole)
	}
	if len(_iottwinmakerSyncSource) > 0 {
		input.SyncSource = aws.String(_iottwinmakerSyncSource)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerTags) > 0 {
		if err := assignInputField(input, "Tags", _iottwinmakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a workplace.
func iottwinmaker_CreateWorkspace(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.CreateWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerRole) > 0 {
		input.Role = aws.String(_iottwinmakerRole)
	}
	if len(_iottwinmakerS3Location) > 0 {
		input.S3Location = aws.String(_iottwinmakerS3Location)
	}
	if len(_iottwinmakerTags) > 0 {
		if err := assignInputField(input, "Tags", _iottwinmakerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a component type.
func iottwinmaker_DeleteComponentType(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.DeleteComponentTypeInput{
		// ComponentTypeId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerComponentTypeId) > 0 {
		input.ComponentTypeId = aws.String(_iottwinmakerComponentTypeId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.DeleteComponentType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an entity.
func iottwinmaker_DeleteEntity(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.DeleteEntityInput{
		// EntityId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerIsRecursive) > 0 {
		if err := assignInputField(input, "IsRecursive", _iottwinmakerIsRecursive); err != nil {
			log.Errorf("invalid --is-recursive: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a scene.
func iottwinmaker_DeleteScene(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.DeleteSceneInput{
		// SceneId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSceneId) > 0 {
		input.SceneId = aws.String(_iottwinmakerSceneId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.DeleteScene(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the SyncJob.
func iottwinmaker_DeleteSyncJob(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.DeleteSyncJobInput{
		// SyncSource: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSyncSource) > 0 {
		input.SyncSource = aws.String(_iottwinmakerSyncSource)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.DeleteSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workspace.
func iottwinmaker_DeleteWorkspace(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.DeleteWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.DeleteWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Run queries to access information from your knowledge graph of entities within
// individual workspaces.
//
// The ExecuteQuery action only works with [Amazon Web Services Java SDK2]. ExecuteQuery will not work with any
// Amazon Web Services Java SDK version < 2.x.
//
// [Amazon Web Services Java SDK2]: https://docs.aws.amazon.com/sdk-for-java/latest/developer-guide/home.html
func iottwinmaker_ExecuteQuery(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ExecuteQueryInput{
		// QueryStatement: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerQueryStatement) > 0 {
		input.QueryStatement = aws.String(_iottwinmakerQueryStatement)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ExecuteQuery(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ExecuteQueryOutput
	p := iottwinmaker.NewExecuteQueryPaginator(client, input)
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

// Retrieves information about a component type.
func iottwinmaker_GetComponentType(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetComponentTypeInput{
		// ComponentTypeId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerComponentTypeId) > 0 {
		input.ComponentTypeId = aws.String(_iottwinmakerComponentTypeId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.GetComponentType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an entity.
func iottwinmaker_GetEntity(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetEntityInput{
		// EntityId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.GetEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a nmetadata transfer job.
func iottwinmaker_GetMetadataTransferJob(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetMetadataTransferJobInput{
		// MetadataTransferJobId: *string, // Required
	}

	if len(_iottwinmakerMetadataTransferJobId) > 0 {
		input.MetadataTransferJobId = aws.String(_iottwinmakerMetadataTransferJobId)
	}

	if resp, err := client.GetMetadataTransferJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the pricing plan.
func iottwinmaker_GetPricingPlan(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetPricingPlanInput{}

	if resp, err := client.GetPricingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the property values for a component, component type, entity, or workspace.
// You must specify a value for either componentName , componentTypeId , entityId ,
// or workspaceId .
func iottwinmaker_GetPropertyValue(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetPropertyValueInput{
		// SelectedProperties: []string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSelectedProperties) > 0 {
		input.SelectedProperties = append([]string(nil), _iottwinmakerSelectedProperties...)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentName) > 0 {
		input.ComponentName = aws.String(_iottwinmakerComponentName)
	}
	if len(_iottwinmakerComponentPath) > 0 {
		input.ComponentPath = aws.String(_iottwinmakerComponentPath)
	}
	if len(_iottwinmakerComponentTypeId) > 0 {
		input.ComponentTypeId = aws.String(_iottwinmakerComponentTypeId)
	}
	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}
	if len(_iottwinmakerPropertyGroupName) > 0 {
		input.PropertyGroupName = aws.String(_iottwinmakerPropertyGroupName)
	}
	if len(_iottwinmakerTabularConditions) > 0 {
		if err := assignInputField(input, "TabularConditions", _iottwinmakerTabularConditions); err != nil {
			log.Errorf("invalid --tabular-conditions: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetPropertyValue(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.GetPropertyValueOutput
	p := iottwinmaker.NewGetPropertyValuePaginator(client, input)
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

// Retrieves information about the history of a time series property value for a
// component, component type, entity, or workspace.
//
// You must specify a value for workspaceId . For entity-specific queries, specify
// values for componentName and entityId . For cross-entity quries, specify a value
// for componentTypeId .
func iottwinmaker_GetPropertyValueHistory(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetPropertyValueHistoryInput{
		// SelectedProperties: []string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSelectedProperties) > 0 {
		input.SelectedProperties = append([]string(nil), _iottwinmakerSelectedProperties...)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentName) > 0 {
		input.ComponentName = aws.String(_iottwinmakerComponentName)
	}
	if len(_iottwinmakerComponentPath) > 0 {
		input.ComponentPath = aws.String(_iottwinmakerComponentPath)
	}
	if len(_iottwinmakerComponentTypeId) > 0 {
		input.ComponentTypeId = aws.String(_iottwinmakerComponentTypeId)
	}
	if len(_iottwinmakerEndDateTime) > 0 {
		if err := assignInputField(input, "EndDateTime", _iottwinmakerEndDateTime); err != nil {
			log.Errorf("invalid --end-date-time: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerEndTime) > 0 {
		input.EndTime = aws.String(_iottwinmakerEndTime)
	}
	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerInterpolation) > 0 {
		if err := assignInputField(input, "Interpolation", _iottwinmakerInterpolation); err != nil {
			log.Errorf("invalid --interpolation: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}
	if len(_iottwinmakerOrderByTime) > 0 {
		if err := assignInputField(input, "OrderByTime", _iottwinmakerOrderByTime); err != nil {
			log.Errorf("invalid --order-by-time: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerPropertyFilters) > 0 {
		if err := assignInputField(input, "PropertyFilters", _iottwinmakerPropertyFilters); err != nil {
			log.Errorf("invalid --property-filters: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerStartDateTime) > 0 {
		if err := assignInputField(input, "StartDateTime", _iottwinmakerStartDateTime); err != nil {
			log.Errorf("invalid --start-date-time: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerStartTime) > 0 {
		input.StartTime = aws.String(_iottwinmakerStartTime)
	}

	if disablePaginator() {
		if resp, err := client.GetPropertyValueHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.GetPropertyValueHistoryOutput
	p := iottwinmaker.NewGetPropertyValueHistoryPaginator(client, input)
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

// Retrieves information about a scene.
func iottwinmaker_GetScene(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetSceneInput{
		// SceneId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSceneId) > 0 {
		input.SceneId = aws.String(_iottwinmakerSceneId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.GetScene(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the SyncJob.
func iottwinmaker_GetSyncJob(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetSyncJobInput{
		// SyncSource: *string, // Required
	}

	if len(_iottwinmakerSyncSource) > 0 {
		input.SyncSource = aws.String(_iottwinmakerSyncSource)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.GetSyncJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a workspace.
func iottwinmaker_GetWorkspace(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.GetWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}

	if resp, err := client.GetWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all component types in a workspace.
func iottwinmaker_ListComponentTypes(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListComponentTypesInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerFilters) > 0 {
		if err := assignInputField(input, "Filters", _iottwinmakerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponentTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListComponentTypesOutput
	p := iottwinmaker.NewListComponentTypesPaginator(client, input)
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

// This API lists the components of an entity.
func iottwinmaker_ListComponents(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListComponentsInput{
		// EntityId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentPath) > 0 {
		input.ComponentPath = aws.String(_iottwinmakerComponentPath)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComponents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListComponentsOutput
	p := iottwinmaker.NewListComponentsPaginator(client, input)
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

// Lists all entities in a workspace.
func iottwinmaker_ListEntities(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListEntitiesInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerFilters) > 0 {
		if err := assignInputField(input, "Filters", _iottwinmakerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListEntities(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListEntitiesOutput
	p := iottwinmaker.NewListEntitiesPaginator(client, input)
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

// Lists the metadata transfer jobs.
func iottwinmaker_ListMetadataTransferJobs(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListMetadataTransferJobsInput{
		// DestinationType: types.DestinationType, // Required
		// SourceType: types.SourceType, // Required
	}

	if len(_iottwinmakerDestinationType) > 0 {
		if err := assignInputField(input, "DestinationType", _iottwinmakerDestinationType); err != nil {
			log.Errorf("invalid --destination-type: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _iottwinmakerSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerFilters) > 0 {
		if err := assignInputField(input, "Filters", _iottwinmakerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMetadataTransferJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListMetadataTransferJobsOutput
	p := iottwinmaker.NewListMetadataTransferJobsPaginator(client, input)
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

// This API lists the properties of a component.
func iottwinmaker_ListProperties(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListPropertiesInput{
		// EntityId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentName) > 0 {
		input.ComponentName = aws.String(_iottwinmakerComponentName)
	}
	if len(_iottwinmakerComponentPath) > 0 {
		input.ComponentPath = aws.String(_iottwinmakerComponentPath)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProperties(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListPropertiesOutput
	p := iottwinmaker.NewListPropertiesPaginator(client, input)
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

// Lists all scenes in a workspace.
func iottwinmaker_ListScenes(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListScenesInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScenes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListScenesOutput
	p := iottwinmaker.NewListScenesPaginator(client, input)
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

// List all SyncJobs.
func iottwinmaker_ListSyncJobs(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListSyncJobsInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSyncJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListSyncJobsOutput
	p := iottwinmaker.NewListSyncJobsPaginator(client, input)
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

// Lists the sync resources.
func iottwinmaker_ListSyncResources(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListSyncResourcesInput{
		// SyncSource: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSyncSource) > 0 {
		input.SyncSource = aws.String(_iottwinmakerSyncSource)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerFilters) > 0 {
		if err := assignInputField(input, "Filters", _iottwinmakerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSyncResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListSyncResourcesOutput
	p := iottwinmaker.NewListSyncResourcesPaginator(client, input)
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

// Lists all tags associated with a resource.
func iottwinmaker_ListTagsForResource(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListTagsForResourceInput{
		// ResourceARN: *string, // Required
	}

	if len(_iottwinmakerResourceARN) > 0 {
		input.ResourceARN = aws.String(_iottwinmakerResourceARN)
	}
	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about workspaces in the current account.
func iottwinmaker_ListWorkspaces(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.ListWorkspacesInput{}

	if len(_iottwinmakerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _iottwinmakerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerNextToken) > 0 {
		input.NextToken = aws.String(_iottwinmakerNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*iottwinmaker.ListWorkspacesOutput
	p := iottwinmaker.NewListWorkspacesPaginator(client, input)
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

// Adds tags to a resource.
func iottwinmaker_TagResource(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.TagResourceInput{
		// ResourceARN: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_iottwinmakerResourceARN) > 0 {
		input.ResourceARN = aws.String(_iottwinmakerResourceARN)
	}
	if len(_iottwinmakerTags) > 0 {
		if err := assignInputField(input, "Tags", _iottwinmakerTags); err != nil {
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

// Removes tags from a resource.
func iottwinmaker_UntagResource(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.UntagResourceInput{
		// ResourceARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_iottwinmakerResourceARN) > 0 {
		input.ResourceARN = aws.String(_iottwinmakerResourceARN)
	}
	if len(_iottwinmakerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _iottwinmakerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates information in a component type.
func iottwinmaker_UpdateComponentType(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.UpdateComponentTypeInput{
		// ComponentTypeId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerComponentTypeId) > 0 {
		input.ComponentTypeId = aws.String(_iottwinmakerComponentTypeId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentTypeName) > 0 {
		input.ComponentTypeName = aws.String(_iottwinmakerComponentTypeName)
	}
	if len(_iottwinmakerCompositeComponentTypes) > 0 {
		if err := assignInputField(input, "CompositeComponentTypes", _iottwinmakerCompositeComponentTypes); err != nil {
			log.Errorf("invalid --composite-component-types: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerExtendsFrom) > 0 {
		input.ExtendsFrom = append([]string(nil), _iottwinmakerExtendsFrom...)
	}
	if len(_iottwinmakerFunctions) > 0 {
		if err := assignInputField(input, "Functions", _iottwinmakerFunctions); err != nil {
			log.Errorf("invalid --functions: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerIsSingleton) > 0 {
		if err := assignInputField(input, "IsSingleton", _iottwinmakerIsSingleton); err != nil {
			log.Errorf("invalid --is-singleton: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerPropertyDefinitions) > 0 {
		if err := assignInputField(input, "PropertyDefinitions", _iottwinmakerPropertyDefinitions); err != nil {
			log.Errorf("invalid --property-definitions: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerPropertyGroups) > 0 {
		if err := assignInputField(input, "PropertyGroups", _iottwinmakerPropertyGroups); err != nil {
			log.Errorf("invalid --property-groups: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateComponentType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an entity.
func iottwinmaker_UpdateEntity(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.UpdateEntityInput{
		// EntityId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerEntityId) > 0 {
		input.EntityId = aws.String(_iottwinmakerEntityId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerComponentUpdates) > 0 {
		if err := assignInputField(input, "ComponentUpdates", _iottwinmakerComponentUpdates); err != nil {
			log.Errorf("invalid --component-updates: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerCompositeComponentUpdates) > 0 {
		if err := assignInputField(input, "CompositeComponentUpdates", _iottwinmakerCompositeComponentUpdates); err != nil {
			log.Errorf("invalid --composite-component-updates: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerEntityName) > 0 {
		input.EntityName = aws.String(_iottwinmakerEntityName)
	}
	if len(_iottwinmakerParentEntityUpdate) > 0 {
		if err := assignInputField(input, "ParentEntityUpdate", _iottwinmakerParentEntityUpdate); err != nil {
			log.Errorf("invalid --parent-entity-update: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateEntity(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the pricing plan.
func iottwinmaker_UpdatePricingPlan(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.UpdatePricingPlanInput{
		// PricingMode: types.PricingMode, // Required
	}

	if len(_iottwinmakerPricingMode) > 0 {
		if err := assignInputField(input, "PricingMode", _iottwinmakerPricingMode); err != nil {
			log.Errorf("invalid --pricing-mode: %s", err.Error())
			return
		}
	}
	if len(_iottwinmakerBundleNames) > 0 {
		input.BundleNames = append([]string(nil), _iottwinmakerBundleNames...)
	}

	if resp, err := client.UpdatePricingPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a scene.
func iottwinmaker_UpdateScene(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.UpdateSceneInput{
		// SceneId: *string, // Required
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerSceneId) > 0 {
		input.SceneId = aws.String(_iottwinmakerSceneId)
	}
	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerCapabilities) > 0 {
		input.Capabilities = append([]string(nil), _iottwinmakerCapabilities...)
	}
	if len(_iottwinmakerContentLocation) > 0 {
		input.ContentLocation = aws.String(_iottwinmakerContentLocation)
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerSceneMetadata) > 0 {
		if err := assignInputField(input, "SceneMetadata", _iottwinmakerSceneMetadata); err != nil {
			log.Errorf("invalid --scene-metadata: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScene(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a workspace.
func iottwinmaker_UpdateWorkspace(cfg aws.Config, client *iottwinmaker.Client) {
	input := &iottwinmaker.UpdateWorkspaceInput{
		// WorkspaceId: *string, // Required
	}

	if len(_iottwinmakerWorkspaceId) > 0 {
		input.WorkspaceId = aws.String(_iottwinmakerWorkspaceId)
	}
	if len(_iottwinmakerDescription) > 0 {
		input.Description = aws.String(_iottwinmakerDescription)
	}
	if len(_iottwinmakerRole) > 0 {
		input.Role = aws.String(_iottwinmakerRole)
	}
	if len(_iottwinmakerS3Location) > 0 {
		input.S3Location = aws.String(_iottwinmakerS3Location)
	}

	if resp, err := client.UpdateWorkspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_iottwinmakerCmd)
	_iottwinmakerCmd.Flags().SortFlags = false

	_iottwinmakerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_iottwinmakerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_iottwinmakerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_iottwinmakerCmd.Flags().StringSliceVarP(&_iottwinmakerBundleNames, "bundle-names", "", nil, "Bundle Names")
	_iottwinmakerCmd.Flags().StringSliceVarP(&_iottwinmakerCapabilities, "capabilities", "", nil, "Capabilities")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerComponentName, "component-name", "", "", "Component Name")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerComponentPath, "component-path", "", "", "Component Path")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerComponentTypeId, "component-type-id", "", "", "Component Type ID")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerComponentTypeName, "component-type-name", "", "", "Component Type Name")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerComponentUpdates, "component-updates", "", "", "Component Updates")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerComponents, "components", "", "", "Components")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerCompositeComponentTypes, "composite-component-types", "", "", "Composite Component Types")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerCompositeComponentUpdates, "composite-component-updates", "", "", "Composite Component Updates")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerCompositeComponents, "composite-components", "", "", "Composite Components")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerContentLocation, "content-location", "", "", "Content Location")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerDescription, "description", "", "", "Description")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerDestination, "destination", "", "", "Destination")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerDestinationType, "destination-type", "", "", "Destination Type")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerEndDateTime, "end-date-time", "", "", "End Date Time")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerEndTime, "end-time", "", "", "End Time")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerEntityId, "entity-id", "", "", "Entity ID")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerEntityName, "entity-name", "", "", "Entity Name")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerEntries, "entries", "", "", "Entries")
	_iottwinmakerCmd.Flags().StringSliceVarP(&_iottwinmakerExtendsFrom, "extends-from", "", nil, "Extends From")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerFilters, "filters", "", "", "Filters")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerFunctions, "functions", "", "", "Functions")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerInterpolation, "interpolation", "", "", "Interpolation")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerIsRecursive, "is-recursive", "", "", "Is Recursive")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerIsSingleton, "is-singleton", "", "", "Is Singleton")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerMaxResults, "max-results", "", "", "Max Results")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerMetadataTransferJobId, "metadata-transfer-job-id", "", "", "Metadata Transfer Job ID")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerNextToken, "next-token", "", "", "Next Token")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerOrderByTime, "order-by-time", "", "", "Order By Time")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerParentEntityId, "parent-entity-id", "", "", "Parent Entity ID")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerParentEntityUpdate, "parent-entity-update", "", "", "Parent Entity Update")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerPricingMode, "pricing-mode", "", "", "Pricing Mode")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerPropertyDefinitions, "property-definitions", "", "", "Property Definitions")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerPropertyFilters, "property-filters", "", "", "Property Filters")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerPropertyGroupName, "property-group-name", "", "", "Property Group Name")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerPropertyGroups, "property-groups", "", "", "Property Groups")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerQueryStatement, "query-statement", "", "", "Query Statement")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerResourceARN, "resource-arn", "", "", "Resource ARN")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerRole, "role", "", "", "Role")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerS3Location, "s3-location", "", "", "S3 Location")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerSceneId, "scene-id", "", "", "Scene ID")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerSceneMetadata, "scene-metadata", "", "", "Scene Metadata")
	_iottwinmakerCmd.Flags().StringSliceVarP(&_iottwinmakerSelectedProperties, "selected-properties", "", nil, "Selected Properties")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerSourceType, "source-type", "", "", "Source Type")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerSources, "sources", "", "", "Sources")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerStartDateTime, "start-date-time", "", "", "Start Date Time")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerStartTime, "start-time", "", "", "Start Time")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerSyncRole, "sync-role", "", "", "Sync Role")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerSyncSource, "sync-source", "", "", "Sync Source")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerTabularConditions, "tabular-conditions", "", "", "Tabular Conditions")
	_iottwinmakerCmd.Flags().StringSliceVarP(&_iottwinmakerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerTags, "tags", "", "", "Tags")
	_iottwinmakerCmd.Flags().StringVarP(&_iottwinmakerWorkspaceId, "workspace-id", "", "", "Workspace ID")

	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerBatchPutPropertyValues, "batch-put-property-values", "", false, "Batch Put Property Values")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCancelMetadataTransferJob, "cancel-metadata-transfer-job", "", false, "Cancel Metadata Transfer Job")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCreateComponentType, "create-component-type", "", false, "Create Component Type")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCreateEntity, "create-entity", "", false, "Create Entity")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCreateMetadataTransferJob, "create-metadata-transfer-job", "", false, "Create Metadata Transfer Job")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCreateScene, "create-scene", "", false, "Create Scene")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCreateSyncJob, "create-sync-job", "", false, "Create Sync Job")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerCreateWorkspace, "create-workspace", "", false, "Create Workspace")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerDeleteComponentType, "delete-component-type", "", false, "Delete Component Type")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerDeleteEntity, "delete-entity", "", false, "Delete Entity")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerDeleteScene, "delete-scene", "", false, "Delete Scene")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerDeleteSyncJob, "delete-sync-job", "", false, "Delete Sync Job")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerDeleteWorkspace, "delete-workspace", "", false, "Delete Workspace")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerExecuteQuery, "execute-query", "", false, "Execute Query")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetComponentType, "get-component-type", "", false, "Get Component Type")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetEntity, "get-entity", "", false, "Get Entity")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetMetadataTransferJob, "get-metadata-transfer-job", "", false, "Get Metadata Transfer Job")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetPricingPlan, "get-pricing-plan", "", false, "Get Pricing Plan")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetPropertyValue, "get-property-value", "", false, "Get Property Value")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetPropertyValueHistory, "get-property-value-history", "", false, "Get Property Value History")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetScene, "get-scene", "", false, "Get Scene")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetSyncJob, "get-sync-job", "", false, "Get Sync Job")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerGetWorkspace, "get-workspace", "", false, "Get Workspace")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListComponentTypes, "list-component-types", "", false, "List Component Types")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListComponents, "list-components", "", false, "List Components")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListEntities, "list-entities", "", false, "List Entities")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListMetadataTransferJobs, "list-metadata-transfer-jobs", "", false, "List Metadata Transfer Jobs")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListProperties, "list-properties", "", false, "List Properties")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListScenes, "list-scenes", "", false, "List Scenes")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListSyncJobs, "list-sync-jobs", "", false, "List Sync Jobs")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListSyncResources, "list-sync-resources", "", false, "List Sync Resources")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerListWorkspaces, "list-workspaces", "", false, "List Workspaces")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerTagResource, "tag-resource", "", false, "Tag Resource")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerUntagResource, "untag-resource", "", false, "Untag Resource")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerUpdateComponentType, "update-component-type", "", false, "Update Component Type")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerUpdateEntity, "update-entity", "", false, "Update Entity")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerUpdatePricingPlan, "update-pricing-plan", "", false, "Update Pricing Plan")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerUpdateScene, "update-scene", "", false, "Update Scene")
	_iottwinmakerCmd.Flags().BoolVarP(&_iottwinmakerUpdateWorkspace, "update-workspace", "", false, "Update Workspace")

}
