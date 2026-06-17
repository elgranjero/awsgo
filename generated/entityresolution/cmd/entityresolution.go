package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/entityresolution"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// entityresolutionCmd represents the entityresolution command
var _entityresolutionCmd = &cobra.Command{
	Use:   "entityresolution",
	Short: "AWS entityresolution CLI",
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
		client := entityresolution.NewFromConfig(cfg)
		if _entityresolutionAddPolicyStatement {
			entityresolution_AddPolicyStatement(cfg, client)
			return
		}
		if _entityresolutionBatchDeleteUniqueId {
			entityresolution_BatchDeleteUniqueId(cfg, client)
			return
		}
		if _entityresolutionCreateIdMappingWorkflow {
			entityresolution_CreateIdMappingWorkflow(cfg, client)
			return
		}
		if _entityresolutionCreateIdNamespace {
			entityresolution_CreateIdNamespace(cfg, client)
			return
		}
		if _entityresolutionCreateMatchingWorkflow {
			entityresolution_CreateMatchingWorkflow(cfg, client)
			return
		}
		if _entityresolutionCreateSchemaMapping {
			entityresolution_CreateSchemaMapping(cfg, client)
			return
		}
		if _entityresolutionDeleteIdMappingWorkflow {
			entityresolution_DeleteIdMappingWorkflow(cfg, client)
			return
		}
		if _entityresolutionDeleteIdNamespace {
			entityresolution_DeleteIdNamespace(cfg, client)
			return
		}
		if _entityresolutionDeleteMatchingWorkflow {
			entityresolution_DeleteMatchingWorkflow(cfg, client)
			return
		}
		if _entityresolutionDeletePolicyStatement {
			entityresolution_DeletePolicyStatement(cfg, client)
			return
		}
		if _entityresolutionDeleteSchemaMapping {
			entityresolution_DeleteSchemaMapping(cfg, client)
			return
		}
		if _entityresolutionGenerateMatchId {
			entityresolution_GenerateMatchId(cfg, client)
			return
		}
		if _entityresolutionGetIdMappingJob {
			entityresolution_GetIdMappingJob(cfg, client)
			return
		}
		if _entityresolutionGetIdMappingWorkflow {
			entityresolution_GetIdMappingWorkflow(cfg, client)
			return
		}
		if _entityresolutionGetIdNamespace {
			entityresolution_GetIdNamespace(cfg, client)
			return
		}
		if _entityresolutionGetMatchId {
			entityresolution_GetMatchId(cfg, client)
			return
		}
		if _entityresolutionGetMatchingJob {
			entityresolution_GetMatchingJob(cfg, client)
			return
		}
		if _entityresolutionGetMatchingWorkflow {
			entityresolution_GetMatchingWorkflow(cfg, client)
			return
		}
		if _entityresolutionGetPolicy {
			entityresolution_GetPolicy(cfg, client)
			return
		}
		if _entityresolutionGetProviderService {
			entityresolution_GetProviderService(cfg, client)
			return
		}
		if _entityresolutionGetSchemaMapping {
			entityresolution_GetSchemaMapping(cfg, client)
			return
		}
		if _entityresolutionListIdMappingJobs {
			entityresolution_ListIdMappingJobs(cfg, client)
			return
		}
		if _entityresolutionListIdMappingWorkflows {
			entityresolution_ListIdMappingWorkflows(cfg, client)
			return
		}
		if _entityresolutionListIdNamespaces {
			entityresolution_ListIdNamespaces(cfg, client)
			return
		}
		if _entityresolutionListMatchingJobs {
			entityresolution_ListMatchingJobs(cfg, client)
			return
		}
		if _entityresolutionListMatchingWorkflows {
			entityresolution_ListMatchingWorkflows(cfg, client)
			return
		}
		if _entityresolutionListProviderServices {
			entityresolution_ListProviderServices(cfg, client)
			return
		}
		if _entityresolutionListSchemaMappings {
			entityresolution_ListSchemaMappings(cfg, client)
			return
		}
		if _entityresolutionListTagsForResource {
			entityresolution_ListTagsForResource(cfg, client)
			return
		}
		if _entityresolutionPutPolicy {
			entityresolution_PutPolicy(cfg, client)
			return
		}
		if _entityresolutionStartIdMappingJob {
			entityresolution_StartIdMappingJob(cfg, client)
			return
		}
		if _entityresolutionStartMatchingJob {
			entityresolution_StartMatchingJob(cfg, client)
			return
		}
		if _entityresolutionTagResource {
			entityresolution_TagResource(cfg, client)
			return
		}
		if _entityresolutionUntagResource {
			entityresolution_UntagResource(cfg, client)
			return
		}
		if _entityresolutionUpdateIdMappingWorkflow {
			entityresolution_UpdateIdMappingWorkflow(cfg, client)
			return
		}
		if _entityresolutionUpdateIdNamespace {
			entityresolution_UpdateIdNamespace(cfg, client)
			return
		}
		if _entityresolutionUpdateMatchingWorkflow {
			entityresolution_UpdateMatchingWorkflow(cfg, client)
			return
		}
		if _entityresolutionUpdateSchemaMapping {
			entityresolution_UpdateSchemaMapping(cfg, client)
			return
		}

	},
}

var (
	_entityresolutionAddPolicyStatement      bool
	_entityresolutionBatchDeleteUniqueId     bool
	_entityresolutionCreateIdMappingWorkflow bool
	_entityresolutionCreateIdNamespace       bool
	_entityresolutionCreateMatchingWorkflow  bool
	_entityresolutionCreateSchemaMapping     bool
	_entityresolutionDeleteIdMappingWorkflow bool
	_entityresolutionDeleteIdNamespace       bool
	_entityresolutionDeleteMatchingWorkflow  bool
	_entityresolutionDeletePolicyStatement   bool
	_entityresolutionDeleteSchemaMapping     bool
	_entityresolutionGenerateMatchId         bool
	_entityresolutionGetIdMappingJob         bool
	_entityresolutionGetIdMappingWorkflow    bool
	_entityresolutionGetIdNamespace          bool
	_entityresolutionGetMatchId              bool
	_entityresolutionGetMatchingJob          bool
	_entityresolutionGetMatchingWorkflow     bool
	_entityresolutionGetPolicy               bool
	_entityresolutionGetProviderService      bool
	_entityresolutionGetSchemaMapping        bool
	_entityresolutionListIdMappingJobs       bool
	_entityresolutionListIdMappingWorkflows  bool
	_entityresolutionListIdNamespaces        bool
	_entityresolutionListMatchingJobs        bool
	_entityresolutionListMatchingWorkflows   bool
	_entityresolutionListProviderServices    bool
	_entityresolutionListSchemaMappings      bool
	_entityresolutionListTagsForResource     bool
	_entityresolutionPutPolicy               bool
	_entityresolutionStartIdMappingJob       bool
	_entityresolutionStartMatchingJob        bool
	_entityresolutionTagResource             bool
	_entityresolutionUntagResource           bool
	_entityresolutionUpdateIdMappingWorkflow bool
	_entityresolutionUpdateIdNamespace       bool
	_entityresolutionUpdateMatchingWorkflow  bool
	_entityresolutionUpdateSchemaMapping     bool

	_entityresolutionAction                      []string
	_entityresolutionApplyNormalization          string
	_entityresolutionArn                         string
	_entityresolutionCondition                   string
	_entityresolutionDescription                 string
	_entityresolutionEffect                      string
	_entityresolutionIdMappingTechniques         string
	_entityresolutionIdMappingWorkflowProperties string
	_entityresolutionIdNamespaceName             string
	_entityresolutionIncrementalRunConfig        string
	_entityresolutionInputSource                 string
	_entityresolutionInputSourceConfig           string
	_entityresolutionJobId                       string
	_entityresolutionJobType                     string
	_entityresolutionMappedInputFields           string
	_entityresolutionMaxResults                  string
	_entityresolutionNextToken                   string
	_entityresolutionOutputSourceConfig          string
	_entityresolutionPolicy                      string
	_entityresolutionPrincipal                   []string
	_entityresolutionProcessingType              string
	_entityresolutionProviderName                string
	_entityresolutionProviderServiceName         string
	_entityresolutionRecord                      string
	_entityresolutionRecords                     string
	_entityresolutionResolutionTechniques        string
	_entityresolutionResourceArn                 string
	_entityresolutionRoleArn                     string
	_entityresolutionSchemaName                  string
	_entityresolutionStatementId                 string
	_entityresolutionTagKeys                     []string
	_entityresolutionTags                        string
	_entityresolutionToken                       string
	_entityresolutionType                        string
	_entityresolutionUniqueIds                   []string
	_entityresolutionWorkflowName                string
)

// Adds a policy statement object. To retrieve a list of existing policy
// statements, use the GetPolicy API.
func entityresolution_AddPolicyStatement(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.AddPolicyStatementInput{
		// Action: []string, // Required
		// Arn: *string, // Required
		// Effect: types.StatementEffect, // Required
		// Principal: []string, // Required
		// StatementId: *string, // Required
	}

	if len(_entityresolutionAction) > 0 {
		input.Action = append([]string(nil), _entityresolutionAction...)
	}
	if len(_entityresolutionArn) > 0 {
		input.Arn = aws.String(_entityresolutionArn)
	}
	if len(_entityresolutionEffect) > 0 {
		if err := assignInputField(input, "Effect", _entityresolutionEffect); err != nil {
			log.Errorf("invalid --effect: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionPrincipal) > 0 {
		input.Principal = append([]string(nil), _entityresolutionPrincipal...)
	}
	if len(_entityresolutionStatementId) > 0 {
		input.StatementId = aws.String(_entityresolutionStatementId)
	}
	if len(_entityresolutionCondition) > 0 {
		input.Condition = aws.String(_entityresolutionCondition)
	}

	if resp, err := client.AddPolicyStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes multiple unique IDs in a matching workflow.
func entityresolution_BatchDeleteUniqueId(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.BatchDeleteUniqueIdInput{
		// UniqueIds: []string, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionUniqueIds) > 0 {
		input.UniqueIds = append([]string(nil), _entityresolutionUniqueIds...)
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionInputSource) > 0 {
		input.InputSource = aws.String(_entityresolutionInputSource)
	}

	if resp, err := client.BatchDeleteUniqueId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IdMappingWorkflow object which stores the configuration of the data
// processing job to be run. Each IdMappingWorkflow must have a unique workflow
// name. To modify an existing workflow, use the UpdateIdMappingWorkflow API.
//
// Incremental processing is not supported for ID mapping workflows.
func entityresolution_CreateIdMappingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.CreateIdMappingWorkflowInput{
		// IdMappingTechniques: *types.IdMappingTechniques, // Required
		// InputSourceConfig: []types.IdMappingWorkflowInputSource, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionIdMappingTechniques) > 0 {
		if err := assignInputField(input, "IdMappingTechniques", _entityresolutionIdMappingTechniques); err != nil {
			log.Errorf("invalid --id-mapping-techniques: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionInputSourceConfig) > 0 {
		if err := assignInputField(input, "InputSourceConfig", _entityresolutionInputSourceConfig); err != nil {
			log.Errorf("invalid --input-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionIncrementalRunConfig) > 0 {
		if err := assignInputField(input, "IncrementalRunConfig", _entityresolutionIncrementalRunConfig); err != nil {
			log.Errorf("invalid --incremental-run-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionOutputSourceConfig) > 0 {
		if err := assignInputField(input, "OutputSourceConfig", _entityresolutionOutputSourceConfig); err != nil {
			log.Errorf("invalid --output-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionRoleArn) > 0 {
		input.RoleArn = aws.String(_entityresolutionRoleArn)
	}
	if len(_entityresolutionTags) > 0 {
		if err := assignInputField(input, "Tags", _entityresolutionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdMappingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an ID namespace object which will help customers provide metadata
// explaining their dataset and how to use it. Each ID namespace must have a unique
// name. To modify an existing ID namespace, use the UpdateIdNamespace API.
func entityresolution_CreateIdNamespace(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.CreateIdNamespaceInput{
		// IdNamespaceName: *string, // Required
		// Type: types.IdNamespaceType, // Required
	}

	if len(_entityresolutionIdNamespaceName) > 0 {
		input.IdNamespaceName = aws.String(_entityresolutionIdNamespaceName)
	}
	if len(_entityresolutionType) > 0 {
		if err := assignInputField(input, "Type", _entityresolutionType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionIdMappingWorkflowProperties) > 0 {
		if err := assignInputField(input, "IdMappingWorkflowProperties", _entityresolutionIdMappingWorkflowProperties); err != nil {
			log.Errorf("invalid --id-mapping-workflow-properties: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionInputSourceConfig) > 0 {
		if err := assignInputField(input, "InputSourceConfig", _entityresolutionInputSourceConfig); err != nil {
			log.Errorf("invalid --input-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionRoleArn) > 0 {
		input.RoleArn = aws.String(_entityresolutionRoleArn)
	}
	if len(_entityresolutionTags) > 0 {
		if err := assignInputField(input, "Tags", _entityresolutionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a matching workflow that defines the configuration for a data
// processing job. The workflow name must be unique. To modify an existing
// workflow, use UpdateMatchingWorkflow .
//
// For workflows where resolutionType is ML_MATCHING or PROVIDER , incremental
// processing is not supported.
func entityresolution_CreateMatchingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.CreateMatchingWorkflowInput{
		// InputSourceConfig: []types.InputSource, // Required
		// OutputSourceConfig: []types.OutputSource, // Required
		// ResolutionTechniques: *types.ResolutionTechniques, // Required
		// RoleArn: *string, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionInputSourceConfig) > 0 {
		if err := assignInputField(input, "InputSourceConfig", _entityresolutionInputSourceConfig); err != nil {
			log.Errorf("invalid --input-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionOutputSourceConfig) > 0 {
		if err := assignInputField(input, "OutputSourceConfig", _entityresolutionOutputSourceConfig); err != nil {
			log.Errorf("invalid --output-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionResolutionTechniques) > 0 {
		if err := assignInputField(input, "ResolutionTechniques", _entityresolutionResolutionTechniques); err != nil {
			log.Errorf("invalid --resolution-techniques: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionRoleArn) > 0 {
		input.RoleArn = aws.String(_entityresolutionRoleArn)
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionIncrementalRunConfig) > 0 {
		if err := assignInputField(input, "IncrementalRunConfig", _entityresolutionIncrementalRunConfig); err != nil {
			log.Errorf("invalid --incremental-run-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionTags) > 0 {
		if err := assignInputField(input, "Tags", _entityresolutionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMatchingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a schema mapping, which defines the schema of the input customer
// records table. The SchemaMapping also provides Entity Resolution with some
// metadata about the table, such as the attribute types of the columns and which
// columns to match on.
func entityresolution_CreateSchemaMapping(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.CreateSchemaMappingInput{
		// MappedInputFields: []types.SchemaInputAttribute, // Required
		// SchemaName: *string, // Required
	}

	if len(_entityresolutionMappedInputFields) > 0 {
		if err := assignInputField(input, "MappedInputFields", _entityresolutionMappedInputFields); err != nil {
			log.Errorf("invalid --mapped-input-fields: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionSchemaName) > 0 {
		input.SchemaName = aws.String(_entityresolutionSchemaName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionTags) > 0 {
		if err := assignInputField(input, "Tags", _entityresolutionTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchemaMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the IdMappingWorkflow with a given name. This operation will succeed
// even if a workflow with the given name does not exist.
func entityresolution_DeleteIdMappingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.DeleteIdMappingWorkflowInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.DeleteIdMappingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the IdNamespace with a given name.
func entityresolution_DeleteIdNamespace(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.DeleteIdNamespaceInput{
		// IdNamespaceName: *string, // Required
	}

	if len(_entityresolutionIdNamespaceName) > 0 {
		input.IdNamespaceName = aws.String(_entityresolutionIdNamespaceName)
	}

	if resp, err := client.DeleteIdNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the MatchingWorkflow with a given name. This operation will succeed
// even if a workflow with the given name does not exist.
func entityresolution_DeleteMatchingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.DeleteMatchingWorkflowInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.DeleteMatchingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the policy statement.
func entityresolution_DeletePolicyStatement(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.DeletePolicyStatementInput{
		// Arn: *string, // Required
		// StatementId: *string, // Required
	}

	if len(_entityresolutionArn) > 0 {
		input.Arn = aws.String(_entityresolutionArn)
	}
	if len(_entityresolutionStatementId) > 0 {
		input.StatementId = aws.String(_entityresolutionStatementId)
	}

	if resp, err := client.DeletePolicyStatement(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the SchemaMapping with a given name. This operation will succeed even
// if a schema with the given name does not exist. This operation will fail if
// there is a MatchingWorkflow object that references the SchemaMapping in the
// workflow's InputSourceConfig .
func entityresolution_DeleteSchemaMapping(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.DeleteSchemaMappingInput{
		// SchemaName: *string, // Required
	}

	if len(_entityresolutionSchemaName) > 0 {
		input.SchemaName = aws.String(_entityresolutionSchemaName)
	}

	if resp, err := client.DeleteSchemaMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Generates or retrieves Match IDs for records using a rule-based matching
// workflow. When you call this operation, it processes your records against the
// workflow's matching rules to identify potential matches. For existing records,
// it retrieves their Match IDs and associated rules. For records without matches,
// it generates new Match IDs. The operation saves results to Amazon S3.
//
// The processing type ( processingType ) you choose affects both the accuracy and
// response time of the operation. Additional charges apply for each API call,
// whether made through the Entity Resolution console or directly via the API. The
// rule-based matching workflow must exist and be active before calling this
// operation.
func entityresolution_GenerateMatchId(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GenerateMatchIdInput{
		// Records: []types.Record, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionRecords) > 0 {
		if err := assignInputField(input, "Records", _entityresolutionRecords); err != nil {
			log.Errorf("invalid --records: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionProcessingType) > 0 {
		if err := assignInputField(input, "ProcessingType", _entityresolutionProcessingType); err != nil {
			log.Errorf("invalid --processing-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GenerateMatchId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status, metrics, and errors (if there are any) that are associated
// with a job.
func entityresolution_GetIdMappingJob(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetIdMappingJobInput{
		// JobId: *string, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionJobId) > 0 {
		input.JobId = aws.String(_entityresolutionJobId)
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.GetIdMappingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the IdMappingWorkflow with a given name, if it exists.
func entityresolution_GetIdMappingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetIdMappingWorkflowInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.GetIdMappingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the IdNamespace with a given name, if it exists.
func entityresolution_GetIdNamespace(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetIdNamespaceInput{
		// IdNamespaceName: *string, // Required
	}

	if len(_entityresolutionIdNamespaceName) > 0 {
		input.IdNamespaceName = aws.String(_entityresolutionIdNamespaceName)
	}

	if resp, err := client.GetIdNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the corresponding Match ID of a customer record if the record has been
// processed in a rule-based matching workflow.
//
// You can call this API as a dry run of an incremental load on the rule-based
// matching workflow.
func entityresolution_GetMatchId(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetMatchIdInput{
		// Record: map[string]string, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionRecord) > 0 {
		if err := assignInputField(input, "Record", _entityresolutionRecord); err != nil {
			log.Errorf("invalid --record: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionApplyNormalization) > 0 {
		if err := assignInputField(input, "ApplyNormalization", _entityresolutionApplyNormalization); err != nil {
			log.Errorf("invalid --apply-normalization: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetMatchId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the status, metrics, and errors (if there are any) that are associated
// with a job.
func entityresolution_GetMatchingJob(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetMatchingJobInput{
		// JobId: *string, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionJobId) > 0 {
		input.JobId = aws.String(_entityresolutionJobId)
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.GetMatchingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the MatchingWorkflow with a given name, if it exists.
func entityresolution_GetMatchingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetMatchingWorkflowInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.GetMatchingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource-based policy.
func entityresolution_GetPolicy(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetPolicyInput{
		// Arn: *string, // Required
	}

	if len(_entityresolutionArn) > 0 {
		input.Arn = aws.String(_entityresolutionArn)
	}

	if resp, err := client.GetPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the ProviderService of a given name.
func entityresolution_GetProviderService(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetProviderServiceInput{
		// ProviderName: *string, // Required
		// ProviderServiceName: *string, // Required
	}

	if len(_entityresolutionProviderName) > 0 {
		input.ProviderName = aws.String(_entityresolutionProviderName)
	}
	if len(_entityresolutionProviderServiceName) > 0 {
		input.ProviderServiceName = aws.String(_entityresolutionProviderServiceName)
	}

	if resp, err := client.GetProviderService(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the SchemaMapping of a given name.
func entityresolution_GetSchemaMapping(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.GetSchemaMappingInput{
		// SchemaName: *string, // Required
	}

	if len(_entityresolutionSchemaName) > 0 {
		input.SchemaName = aws.String(_entityresolutionSchemaName)
	}

	if resp, err := client.GetSchemaMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all ID mapping jobs for a given workflow.
func entityresolution_ListIdMappingJobs(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListIdMappingJobsInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdMappingJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListIdMappingJobsOutput
	p := entityresolution.NewListIdMappingJobsPaginator(client, input)
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

// Returns a list of all the IdMappingWorkflows that have been created for an
// Amazon Web Services account.
func entityresolution_ListIdMappingWorkflows(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListIdMappingWorkflowsInput{}

	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdMappingWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListIdMappingWorkflowsOutput
	p := entityresolution.NewListIdMappingWorkflowsPaginator(client, input)
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

// Returns a list of all ID namespaces.
func entityresolution_ListIdNamespaces(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListIdNamespacesInput{}

	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListIdNamespacesOutput
	p := entityresolution.NewListIdNamespacesPaginator(client, input)
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

// Lists all jobs for a given workflow.
func entityresolution_ListMatchingJobs(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListMatchingJobsInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMatchingJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListMatchingJobsOutput
	p := entityresolution.NewListMatchingJobsPaginator(client, input)
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

// Returns a list of all the MatchingWorkflows that have been created for an
// Amazon Web Services account.
func entityresolution_ListMatchingWorkflows(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListMatchingWorkflowsInput{}

	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMatchingWorkflows(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListMatchingWorkflowsOutput
	p := entityresolution.NewListMatchingWorkflowsPaginator(client, input)
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

// Returns a list of all the ProviderServices that are available in this Amazon
// Web Services Region.
func entityresolution_ListProviderServices(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListProviderServicesInput{}

	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}
	if len(_entityresolutionProviderName) > 0 {
		input.ProviderName = aws.String(_entityresolutionProviderName)
	}

	if disablePaginator() {
		if resp, err := client.ListProviderServices(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListProviderServicesOutput
	p := entityresolution.NewListProviderServicesPaginator(client, input)
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

// Returns a list of all the SchemaMappings that have been created for an Amazon
// Web Services account.
func entityresolution_ListSchemaMappings(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListSchemaMappingsInput{}

	if len(_entityresolutionMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _entityresolutionMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionNextToken) > 0 {
		input.NextToken = aws.String(_entityresolutionNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemaMappings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*entityresolution.ListSchemaMappingsOutput
	p := entityresolution.NewListSchemaMappingsPaginator(client, input)
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

// Displays the tags associated with an Entity Resolution resource. In Entity
// Resolution, SchemaMapping , and MatchingWorkflow can be tagged.
func entityresolution_ListTagsForResource(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_entityresolutionResourceArn) > 0 {
		input.ResourceArn = aws.String(_entityresolutionResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource-based policy.
func entityresolution_PutPolicy(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.PutPolicyInput{
		// Arn: *string, // Required
		// Policy: *string, // Required
	}

	if len(_entityresolutionArn) > 0 {
		input.Arn = aws.String(_entityresolutionArn)
	}
	if len(_entityresolutionPolicy) > 0 {
		input.Policy = aws.String(_entityresolutionPolicy)
	}
	if len(_entityresolutionToken) > 0 {
		input.Token = aws.String(_entityresolutionToken)
	}

	if resp, err := client.PutPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the IdMappingJob of a workflow. The workflow must have previously been
// created using the CreateIdMappingWorkflow endpoint.
func entityresolution_StartIdMappingJob(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.StartIdMappingJobInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionJobType) > 0 {
		if err := assignInputField(input, "JobType", _entityresolutionJobType); err != nil {
			log.Errorf("invalid --job-type: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionOutputSourceConfig) > 0 {
		if err := assignInputField(input, "OutputSourceConfig", _entityresolutionOutputSourceConfig); err != nil {
			log.Errorf("invalid --output-source-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartIdMappingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the MatchingJob of a workflow. The workflow must have previously been
// created using the CreateMatchingWorkflow endpoint.
func entityresolution_StartMatchingJob(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.StartMatchingJobInput{
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}

	if resp, err := client.StartMatchingJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified Entity Resolution
// resource. Tags can help you organize and categorize your resources. You can also
// use them to scope user permissions by granting a user permission to access or
// change only resources with certain tag values. In Entity Resolution,
// SchemaMapping and MatchingWorkflow can be tagged. Tags don't have any semantic
// meaning to Amazon Web Services and are interpreted strictly as strings of
// characters. You can use the TagResource action with a resource that already has
// tags. If you specify a new tag key, this tag is appended to the list of tags
// associated with the resource. If you specify a tag key that is already
// associated with the resource, the new tag value that you specify replaces the
// previous value for that tag.
func entityresolution_TagResource(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_entityresolutionResourceArn) > 0 {
		input.ResourceArn = aws.String(_entityresolutionResourceArn)
	}
	if len(_entityresolutionTags) > 0 {
		if err := assignInputField(input, "Tags", _entityresolutionTags); err != nil {
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

// Removes one or more tags from the specified Entity Resolution resource. In
// Entity Resolution, SchemaMapping , and MatchingWorkflow can be tagged.
func entityresolution_UntagResource(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_entityresolutionResourceArn) > 0 {
		input.ResourceArn = aws.String(_entityresolutionResourceArn)
	}
	if len(_entityresolutionTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _entityresolutionTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing IdMappingWorkflow . This method is identical to
// CreateIdMappingWorkflow, except it uses an HTTP PUT request instead of a POST
// request, and the IdMappingWorkflow must already exist for the method to succeed.
//
// Incremental processing is not supported for ID mapping workflows.
func entityresolution_UpdateIdMappingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.UpdateIdMappingWorkflowInput{
		// IdMappingTechniques: *types.IdMappingTechniques, // Required
		// InputSourceConfig: []types.IdMappingWorkflowInputSource, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionIdMappingTechniques) > 0 {
		if err := assignInputField(input, "IdMappingTechniques", _entityresolutionIdMappingTechniques); err != nil {
			log.Errorf("invalid --id-mapping-techniques: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionInputSourceConfig) > 0 {
		if err := assignInputField(input, "InputSourceConfig", _entityresolutionInputSourceConfig); err != nil {
			log.Errorf("invalid --input-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionIncrementalRunConfig) > 0 {
		if err := assignInputField(input, "IncrementalRunConfig", _entityresolutionIncrementalRunConfig); err != nil {
			log.Errorf("invalid --incremental-run-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionOutputSourceConfig) > 0 {
		if err := assignInputField(input, "OutputSourceConfig", _entityresolutionOutputSourceConfig); err != nil {
			log.Errorf("invalid --output-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionRoleArn) > 0 {
		input.RoleArn = aws.String(_entityresolutionRoleArn)
	}

	if resp, err := client.UpdateIdMappingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing ID namespace.
func entityresolution_UpdateIdNamespace(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.UpdateIdNamespaceInput{
		// IdNamespaceName: *string, // Required
	}

	if len(_entityresolutionIdNamespaceName) > 0 {
		input.IdNamespaceName = aws.String(_entityresolutionIdNamespaceName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionIdMappingWorkflowProperties) > 0 {
		if err := assignInputField(input, "IdMappingWorkflowProperties", _entityresolutionIdMappingWorkflowProperties); err != nil {
			log.Errorf("invalid --id-mapping-workflow-properties: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionInputSourceConfig) > 0 {
		if err := assignInputField(input, "InputSourceConfig", _entityresolutionInputSourceConfig); err != nil {
			log.Errorf("invalid --input-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionRoleArn) > 0 {
		input.RoleArn = aws.String(_entityresolutionRoleArn)
	}

	if resp, err := client.UpdateIdNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing matching workflow. The workflow must already exist for this
// operation to succeed.
//
// For workflows where resolutionType is ML_MATCHING or PROVIDER , incremental
// processing is not supported.
func entityresolution_UpdateMatchingWorkflow(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.UpdateMatchingWorkflowInput{
		// InputSourceConfig: []types.InputSource, // Required
		// OutputSourceConfig: []types.OutputSource, // Required
		// ResolutionTechniques: *types.ResolutionTechniques, // Required
		// RoleArn: *string, // Required
		// WorkflowName: *string, // Required
	}

	if len(_entityresolutionInputSourceConfig) > 0 {
		if err := assignInputField(input, "InputSourceConfig", _entityresolutionInputSourceConfig); err != nil {
			log.Errorf("invalid --input-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionOutputSourceConfig) > 0 {
		if err := assignInputField(input, "OutputSourceConfig", _entityresolutionOutputSourceConfig); err != nil {
			log.Errorf("invalid --output-source-config: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionResolutionTechniques) > 0 {
		if err := assignInputField(input, "ResolutionTechniques", _entityresolutionResolutionTechniques); err != nil {
			log.Errorf("invalid --resolution-techniques: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionRoleArn) > 0 {
		input.RoleArn = aws.String(_entityresolutionRoleArn)
	}
	if len(_entityresolutionWorkflowName) > 0 {
		input.WorkflowName = aws.String(_entityresolutionWorkflowName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}
	if len(_entityresolutionIncrementalRunConfig) > 0 {
		if err := assignInputField(input, "IncrementalRunConfig", _entityresolutionIncrementalRunConfig); err != nil {
			log.Errorf("invalid --incremental-run-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMatchingWorkflow(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a schema mapping.
// A schema is immutable if it is being used by a workflow. Therefore, you can't
// update a schema mapping if it's associated with a workflow.
func entityresolution_UpdateSchemaMapping(cfg aws.Config, client *entityresolution.Client) {
	input := &entityresolution.UpdateSchemaMappingInput{
		// MappedInputFields: []types.SchemaInputAttribute, // Required
		// SchemaName: *string, // Required
	}

	if len(_entityresolutionMappedInputFields) > 0 {
		if err := assignInputField(input, "MappedInputFields", _entityresolutionMappedInputFields); err != nil {
			log.Errorf("invalid --mapped-input-fields: %s", err.Error())
			return
		}
	}
	if len(_entityresolutionSchemaName) > 0 {
		input.SchemaName = aws.String(_entityresolutionSchemaName)
	}
	if len(_entityresolutionDescription) > 0 {
		input.Description = aws.String(_entityresolutionDescription)
	}

	if resp, err := client.UpdateSchemaMapping(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_entityresolutionCmd)
	_entityresolutionCmd.Flags().SortFlags = false

	_entityresolutionCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_entityresolutionCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_entityresolutionCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_entityresolutionCmd.Flags().StringSliceVarP(&_entityresolutionAction, "action", "", nil, "Action")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionApplyNormalization, "apply-normalization", "", "", "Apply Normalization")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionArn, "arn", "", "", "ARN")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionCondition, "condition", "", "", "Condition")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionDescription, "description", "", "", "Description")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionEffect, "effect", "", "", "Effect")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionIdMappingTechniques, "id-mapping-techniques", "", "", "ID Mapping Techniques")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionIdMappingWorkflowProperties, "id-mapping-workflow-properties", "", "", "ID Mapping Workflow Properties")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionIdNamespaceName, "id-namespace-name", "", "", "ID Namespace Name")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionIncrementalRunConfig, "incremental-run-config", "", "", "Incremental Run Config")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionInputSource, "input-source", "", "", "Input Source")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionInputSourceConfig, "input-source-config", "", "", "Input Source Config")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionJobId, "job-id", "", "", "Job ID")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionJobType, "job-type", "", "", "Job Type")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionMappedInputFields, "mapped-input-fields", "", "", "Mapped Input Fields")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionMaxResults, "max-results", "", "", "Max Results")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionNextToken, "next-token", "", "", "Next Token")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionOutputSourceConfig, "output-source-config", "", "", "Output Source Config")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionPolicy, "policy", "", "", "Policy")
	_entityresolutionCmd.Flags().StringSliceVarP(&_entityresolutionPrincipal, "principal", "", nil, "Principal")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionProcessingType, "processing-type", "", "", "Processing Type")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionProviderName, "provider-name", "", "", "Provider Name")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionProviderServiceName, "provider-service-name", "", "", "Provider Service Name")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionRecord, "record", "", "", "Record")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionRecords, "records", "", "", "Records")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionResolutionTechniques, "resolution-techniques", "", "", "Resolution Techniques")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionResourceArn, "resource-arn", "", "", "Resource ARN")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionRoleArn, "role-arn", "", "", "Role ARN")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionSchemaName, "schema-name", "", "", "Schema Name")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionStatementId, "statement-id", "", "", "Statement ID")
	_entityresolutionCmd.Flags().StringSliceVarP(&_entityresolutionTagKeys, "tag-keys", "", nil, "Tag Keys")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionTags, "tags", "", "", "Tags")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionToken, "token", "", "", "Token")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionType, "type", "", "", "Type")
	_entityresolutionCmd.Flags().StringSliceVarP(&_entityresolutionUniqueIds, "unique-ids", "", nil, "Unique Ids")
	_entityresolutionCmd.Flags().StringVarP(&_entityresolutionWorkflowName, "workflow-name", "", "", "Workflow Name")

	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionAddPolicyStatement, "add-policy-statement", "", false, "Add Policy Statement")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionBatchDeleteUniqueId, "batch-delete-unique-id", "", false, "Batch Delete Unique ID")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionCreateIdMappingWorkflow, "create-id-mapping-workflow", "", false, "Create ID Mapping Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionCreateIdNamespace, "create-id-namespace", "", false, "Create ID Namespace")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionCreateMatchingWorkflow, "create-matching-workflow", "", false, "Create Matching Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionCreateSchemaMapping, "create-schema-mapping", "", false, "Create Schema Mapping")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionDeleteIdMappingWorkflow, "delete-id-mapping-workflow", "", false, "Delete ID Mapping Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionDeleteIdNamespace, "delete-id-namespace", "", false, "Delete ID Namespace")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionDeleteMatchingWorkflow, "delete-matching-workflow", "", false, "Delete Matching Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionDeletePolicyStatement, "delete-policy-statement", "", false, "Delete Policy Statement")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionDeleteSchemaMapping, "delete-schema-mapping", "", false, "Delete Schema Mapping")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGenerateMatchId, "generate-match-id", "", false, "Generate Match ID")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetIdMappingJob, "get-id-mapping-job", "", false, "Get ID Mapping Job")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetIdMappingWorkflow, "get-id-mapping-workflow", "", false, "Get ID Mapping Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetIdNamespace, "get-id-namespace", "", false, "Get ID Namespace")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetMatchId, "get-match-id", "", false, "Get Match ID")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetMatchingJob, "get-matching-job", "", false, "Get Matching Job")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetMatchingWorkflow, "get-matching-workflow", "", false, "Get Matching Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetPolicy, "get-policy", "", false, "Get Policy")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetProviderService, "get-provider-service", "", false, "Get Provider Service")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionGetSchemaMapping, "get-schema-mapping", "", false, "Get Schema Mapping")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListIdMappingJobs, "list-id-mapping-jobs", "", false, "List ID Mapping Jobs")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListIdMappingWorkflows, "list-id-mapping-workflows", "", false, "List ID Mapping Workflows")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListIdNamespaces, "list-id-namespaces", "", false, "List ID Namespaces")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListMatchingJobs, "list-matching-jobs", "", false, "List Matching Jobs")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListMatchingWorkflows, "list-matching-workflows", "", false, "List Matching Workflows")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListProviderServices, "list-provider-services", "", false, "List Provider Services")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListSchemaMappings, "list-schema-mappings", "", false, "List Schema Mappings")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionPutPolicy, "put-policy", "", false, "Put Policy")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionStartIdMappingJob, "start-id-mapping-job", "", false, "Start ID Mapping Job")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionStartMatchingJob, "start-matching-job", "", false, "Start Matching Job")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionTagResource, "tag-resource", "", false, "Tag Resource")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionUntagResource, "untag-resource", "", false, "Untag Resource")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionUpdateIdMappingWorkflow, "update-id-mapping-workflow", "", false, "Update ID Mapping Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionUpdateIdNamespace, "update-id-namespace", "", false, "Update ID Namespace")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionUpdateMatchingWorkflow, "update-matching-workflow", "", false, "Update Matching Workflow")
	_entityresolutionCmd.Flags().BoolVarP(&_entityresolutionUpdateSchemaMapping, "update-schema-mapping", "", false, "Update Schema Mapping")

}
