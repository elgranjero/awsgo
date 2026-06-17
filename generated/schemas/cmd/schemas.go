package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/schemas"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// schemasCmd represents the schemas command
var _schemasCmd = &cobra.Command{
	Use:   "schemas",
	Short: "AWS schemas CLI",
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
		client := schemas.NewFromConfig(cfg)
		if _schemasCreateDiscoverer {
			schemas_CreateDiscoverer(cfg, client)
			return
		}
		if _schemasCreateRegistry {
			schemas_CreateRegistry(cfg, client)
			return
		}
		if _schemasCreateSchema {
			schemas_CreateSchema(cfg, client)
			return
		}
		if _schemasDeleteDiscoverer {
			schemas_DeleteDiscoverer(cfg, client)
			return
		}
		if _schemasDeleteRegistry {
			schemas_DeleteRegistry(cfg, client)
			return
		}
		if _schemasDeleteResourcePolicy {
			schemas_DeleteResourcePolicy(cfg, client)
			return
		}
		if _schemasDeleteSchema {
			schemas_DeleteSchema(cfg, client)
			return
		}
		if _schemasDeleteSchemaVersion {
			schemas_DeleteSchemaVersion(cfg, client)
			return
		}
		if _schemasDescribeCodeBinding {
			schemas_DescribeCodeBinding(cfg, client)
			return
		}
		if _schemasDescribeDiscoverer {
			schemas_DescribeDiscoverer(cfg, client)
			return
		}
		if _schemasDescribeRegistry {
			schemas_DescribeRegistry(cfg, client)
			return
		}
		if _schemasDescribeSchema {
			schemas_DescribeSchema(cfg, client)
			return
		}
		if _schemasExportSchema {
			schemas_ExportSchema(cfg, client)
			return
		}
		if _schemasGetCodeBindingSource {
			schemas_GetCodeBindingSource(cfg, client)
			return
		}
		if _schemasGetDiscoveredSchema {
			schemas_GetDiscoveredSchema(cfg, client)
			return
		}
		if _schemasGetResourcePolicy {
			schemas_GetResourcePolicy(cfg, client)
			return
		}
		if _schemasListDiscoverers {
			schemas_ListDiscoverers(cfg, client)
			return
		}
		if _schemasListRegistries {
			schemas_ListRegistries(cfg, client)
			return
		}
		if _schemasListSchemaVersions {
			schemas_ListSchemaVersions(cfg, client)
			return
		}
		if _schemasListSchemas {
			schemas_ListSchemas(cfg, client)
			return
		}
		if _schemasListTagsForResource {
			schemas_ListTagsForResource(cfg, client)
			return
		}
		if _schemasPutCodeBinding {
			schemas_PutCodeBinding(cfg, client)
			return
		}
		if _schemasPutResourcePolicy {
			schemas_PutResourcePolicy(cfg, client)
			return
		}
		if _schemasSearchSchemas {
			schemas_SearchSchemas(cfg, client)
			return
		}
		if _schemasStartDiscoverer {
			schemas_StartDiscoverer(cfg, client)
			return
		}
		if _schemasStopDiscoverer {
			schemas_StopDiscoverer(cfg, client)
			return
		}
		if _schemasTagResource {
			schemas_TagResource(cfg, client)
			return
		}
		if _schemasUntagResource {
			schemas_UntagResource(cfg, client)
			return
		}
		if _schemasUpdateDiscoverer {
			schemas_UpdateDiscoverer(cfg, client)
			return
		}
		if _schemasUpdateRegistry {
			schemas_UpdateRegistry(cfg, client)
			return
		}
		if _schemasUpdateSchema {
			schemas_UpdateSchema(cfg, client)
			return
		}

	},
}

var (
	_schemasCreateDiscoverer     bool
	_schemasCreateRegistry       bool
	_schemasCreateSchema         bool
	_schemasDeleteDiscoverer     bool
	_schemasDeleteRegistry       bool
	_schemasDeleteResourcePolicy bool
	_schemasDeleteSchema         bool
	_schemasDeleteSchemaVersion  bool
	_schemasDescribeCodeBinding  bool
	_schemasDescribeDiscoverer   bool
	_schemasDescribeRegistry     bool
	_schemasDescribeSchema       bool
	_schemasExportSchema         bool
	_schemasGetCodeBindingSource bool
	_schemasGetDiscoveredSchema  bool
	_schemasGetResourcePolicy    bool
	_schemasListDiscoverers      bool
	_schemasListRegistries       bool
	_schemasListSchemaVersions   bool
	_schemasListSchemas          bool
	_schemasListTagsForResource  bool
	_schemasPutCodeBinding       bool
	_schemasPutResourcePolicy    bool
	_schemasSearchSchemas        bool
	_schemasStartDiscoverer      bool
	_schemasStopDiscoverer       bool
	_schemasTagResource          bool
	_schemasUntagResource        bool
	_schemasUpdateDiscoverer     bool
	_schemasUpdateRegistry       bool
	_schemasUpdateSchema         bool

	_schemasClientTokenId      string
	_schemasContent            string
	_schemasCrossAccount       string
	_schemasDescription        string
	_schemasDiscovererId       string
	_schemasDiscovererIdPrefix string
	_schemasEvents             []string
	_schemasKeywords           string
	_schemasLanguage           string
	_schemasLimit              string
	_schemasNextToken          string
	_schemasPolicy             string
	_schemasRegistryName       string
	_schemasRegistryNamePrefix string
	_schemasResourceArn        string
	_schemasRevisionId         string
	_schemasSchemaName         string
	_schemasSchemaNamePrefix   string
	_schemasSchemaVersion      string
	_schemasScope              string
	_schemasSourceArn          string
	_schemasSourceArnPrefix    string
	_schemasTagKeys            []string
	_schemasTags               string
	_schemasType               string
)

// Creates a discoverer.
func schemas_CreateDiscoverer(cfg aws.Config, client *schemas.Client) {
	input := &schemas.CreateDiscovererInput{
		// SourceArn: *string, // Required
	}

	if len(_schemasSourceArn) > 0 {
		input.SourceArn = aws.String(_schemasSourceArn)
	}
	if len(_schemasCrossAccount) > 0 {
		if err := assignInputField(input, "CrossAccount", _schemasCrossAccount); err != nil {
			log.Errorf("invalid --cross-account: %s", err.Error())
			return
		}
	}
	if len(_schemasDescription) > 0 {
		input.Description = aws.String(_schemasDescription)
	}
	if len(_schemasTags) > 0 {
		if err := assignInputField(input, "Tags", _schemasTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDiscoverer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a registry.
func schemas_CreateRegistry(cfg aws.Config, client *schemas.Client) {
	input := &schemas.CreateRegistryInput{
		// RegistryName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasDescription) > 0 {
		input.Description = aws.String(_schemasDescription)
	}
	if len(_schemasTags) > 0 {
		if err := assignInputField(input, "Tags", _schemasTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a schema definition.
// Inactive schemas will be deleted after two years.
func schemas_CreateSchema(cfg aws.Config, client *schemas.Client) {
	input := &schemas.CreateSchemaInput{
		// Content: *string, // Required
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
		// Type: types.Type, // Required
	}

	if len(_schemasContent) > 0 {
		input.Content = aws.String(_schemasContent)
	}
	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasType) > 0 {
		if err := assignInputField(input, "Type", _schemasType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_schemasDescription) > 0 {
		input.Description = aws.String(_schemasDescription)
	}
	if len(_schemasTags) > 0 {
		if err := assignInputField(input, "Tags", _schemasTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a discoverer.
func schemas_DeleteDiscoverer(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DeleteDiscovererInput{
		// DiscovererId: *string, // Required
	}

	if len(_schemasDiscovererId) > 0 {
		input.DiscovererId = aws.String(_schemasDiscovererId)
	}

	if resp, err := client.DeleteDiscoverer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Registry.
func schemas_DeleteRegistry(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DeleteRegistryInput{
		// RegistryName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}

	if resp, err := client.DeleteRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the resource-based policy attached to the specified registry.
func schemas_DeleteResourcePolicy(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DeleteResourcePolicyInput{}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a schema definition.
func schemas_DeleteSchema(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DeleteSchemaInput{
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}

	if resp, err := client.DeleteSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the schema version definition
func schemas_DeleteSchemaVersion(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DeleteSchemaVersionInput{
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
		// SchemaVersion: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasSchemaVersion) > 0 {
		input.SchemaVersion = aws.String(_schemasSchemaVersion)
	}

	if resp, err := client.DeleteSchemaVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe the code binding URI.
func schemas_DescribeCodeBinding(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DescribeCodeBindingInput{
		// Language: *string, // Required
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasLanguage) > 0 {
		input.Language = aws.String(_schemasLanguage)
	}
	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasSchemaVersion) > 0 {
		input.SchemaVersion = aws.String(_schemasSchemaVersion)
	}

	if resp, err := client.DescribeCodeBinding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the discoverer.
func schemas_DescribeDiscoverer(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DescribeDiscovererInput{
		// DiscovererId: *string, // Required
	}

	if len(_schemasDiscovererId) > 0 {
		input.DiscovererId = aws.String(_schemasDiscovererId)
	}

	if resp, err := client.DescribeDiscoverer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the registry.
func schemas_DescribeRegistry(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DescribeRegistryInput{
		// RegistryName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}

	if resp, err := client.DescribeRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the schema definition.
func schemas_DescribeSchema(cfg aws.Config, client *schemas.Client) {
	input := &schemas.DescribeSchemaInput{
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasSchemaVersion) > 0 {
		input.SchemaVersion = aws.String(_schemasSchemaVersion)
	}

	if resp, err := client.DescribeSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func schemas_ExportSchema(cfg aws.Config, client *schemas.Client) {
	input := &schemas.ExportSchemaInput{
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
		// Type: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasType) > 0 {
		input.Type = aws.String(_schemasType)
	}
	if len(_schemasSchemaVersion) > 0 {
		input.SchemaVersion = aws.String(_schemasSchemaVersion)
	}

	if resp, err := client.ExportSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the code binding source URI.
func schemas_GetCodeBindingSource(cfg aws.Config, client *schemas.Client) {
	input := &schemas.GetCodeBindingSourceInput{
		// Language: *string, // Required
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasLanguage) > 0 {
		input.Language = aws.String(_schemasLanguage)
	}
	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasSchemaVersion) > 0 {
		input.SchemaVersion = aws.String(_schemasSchemaVersion)
	}

	if resp, err := client.GetCodeBindingSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the discovered schema that was generated based on sampled events.
func schemas_GetDiscoveredSchema(cfg aws.Config, client *schemas.Client) {
	input := &schemas.GetDiscoveredSchemaInput{
		// Events: []string, // Required
		// Type: types.Type, // Required
	}

	if len(_schemasEvents) > 0 {
		input.Events = append([]string(nil), _schemasEvents...)
	}
	if len(_schemasType) > 0 {
		if err := assignInputField(input, "Type", _schemasType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDiscoveredSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource-based policy attached to a given registry.
func schemas_GetResourcePolicy(cfg aws.Config, client *schemas.Client) {
	input := &schemas.GetResourcePolicyInput{}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the discoverers.
func schemas_ListDiscoverers(cfg aws.Config, client *schemas.Client) {
	input := &schemas.ListDiscoverersInput{}

	if len(_schemasDiscovererIdPrefix) > 0 {
		input.DiscovererIdPrefix = aws.String(_schemasDiscovererIdPrefix)
	}
	if len(_schemasLimit) > 0 {
		if err := assignInputField(input, "Limit", _schemasLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_schemasNextToken) > 0 {
		input.NextToken = aws.String(_schemasNextToken)
	}
	if len(_schemasSourceArnPrefix) > 0 {
		input.SourceArnPrefix = aws.String(_schemasSourceArnPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListDiscoverers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*schemas.ListDiscoverersOutput
	p := schemas.NewListDiscoverersPaginator(client, input)
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

// List the registries.
func schemas_ListRegistries(cfg aws.Config, client *schemas.Client) {
	input := &schemas.ListRegistriesInput{}

	if len(_schemasLimit) > 0 {
		if err := assignInputField(input, "Limit", _schemasLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_schemasNextToken) > 0 {
		input.NextToken = aws.String(_schemasNextToken)
	}
	if len(_schemasRegistryNamePrefix) > 0 {
		input.RegistryNamePrefix = aws.String(_schemasRegistryNamePrefix)
	}
	if len(_schemasScope) > 0 {
		input.Scope = aws.String(_schemasScope)
	}

	if disablePaginator() {
		if resp, err := client.ListRegistries(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*schemas.ListRegistriesOutput
	p := schemas.NewListRegistriesPaginator(client, input)
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

// Provides a list of the schema versions and related information.
func schemas_ListSchemaVersions(cfg aws.Config, client *schemas.Client) {
	input := &schemas.ListSchemaVersionsInput{
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasLimit) > 0 {
		if err := assignInputField(input, "Limit", _schemasLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_schemasNextToken) > 0 {
		input.NextToken = aws.String(_schemasNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemaVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*schemas.ListSchemaVersionsOutput
	p := schemas.NewListSchemaVersionsPaginator(client, input)
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

// List the schemas.
func schemas_ListSchemas(cfg aws.Config, client *schemas.Client) {
	input := &schemas.ListSchemasInput{
		// RegistryName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasLimit) > 0 {
		if err := assignInputField(input, "Limit", _schemasLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_schemasNextToken) > 0 {
		input.NextToken = aws.String(_schemasNextToken)
	}
	if len(_schemasSchemaNamePrefix) > 0 {
		input.SchemaNamePrefix = aws.String(_schemasSchemaNamePrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*schemas.ListSchemasOutput
	p := schemas.NewListSchemasPaginator(client, input)
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

// Get tags for resource.
func schemas_ListTagsForResource(cfg aws.Config, client *schemas.Client) {
	input := &schemas.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_schemasResourceArn) > 0 {
		input.ResourceArn = aws.String(_schemasResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Put code binding URI
func schemas_PutCodeBinding(cfg aws.Config, client *schemas.Client) {
	input := &schemas.PutCodeBindingInput{
		// Language: *string, // Required
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasLanguage) > 0 {
		input.Language = aws.String(_schemasLanguage)
	}
	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasSchemaVersion) > 0 {
		input.SchemaVersion = aws.String(_schemasSchemaVersion)
	}

	if resp, err := client.PutCodeBinding(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The name of the policy.
func schemas_PutResourcePolicy(cfg aws.Config, client *schemas.Client) {
	input := &schemas.PutResourcePolicyInput{
		// Policy: *string, // Required
	}

	if len(_schemasPolicy) > 0 {
		input.Policy = aws.String(_schemasPolicy)
	}
	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasRevisionId) > 0 {
		input.RevisionId = aws.String(_schemasRevisionId)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Search the schemas
func schemas_SearchSchemas(cfg aws.Config, client *schemas.Client) {
	input := &schemas.SearchSchemasInput{
		// Keywords: *string, // Required
		// RegistryName: *string, // Required
	}

	if len(_schemasKeywords) > 0 {
		input.Keywords = aws.String(_schemasKeywords)
	}
	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasLimit) > 0 {
		if err := assignInputField(input, "Limit", _schemasLimit); err != nil {
			log.Errorf("invalid --limit: %s", err.Error())
			return
		}
	}
	if len(_schemasNextToken) > 0 {
		input.NextToken = aws.String(_schemasNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchSchemas(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*schemas.SearchSchemasOutput
	p := schemas.NewSearchSchemasPaginator(client, input)
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

// Starts the discoverer
func schemas_StartDiscoverer(cfg aws.Config, client *schemas.Client) {
	input := &schemas.StartDiscovererInput{
		// DiscovererId: *string, // Required
	}

	if len(_schemasDiscovererId) > 0 {
		input.DiscovererId = aws.String(_schemasDiscovererId)
	}

	if resp, err := client.StartDiscoverer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the discoverer
func schemas_StopDiscoverer(cfg aws.Config, client *schemas.Client) {
	input := &schemas.StopDiscovererInput{
		// DiscovererId: *string, // Required
	}

	if len(_schemasDiscovererId) > 0 {
		input.DiscovererId = aws.String(_schemasDiscovererId)
	}

	if resp, err := client.StopDiscoverer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Add tags to a resource.
func schemas_TagResource(cfg aws.Config, client *schemas.Client) {
	input := &schemas.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_schemasResourceArn) > 0 {
		input.ResourceArn = aws.String(_schemasResourceArn)
	}
	if len(_schemasTags) > 0 {
		if err := assignInputField(input, "Tags", _schemasTags); err != nil {
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
func schemas_UntagResource(cfg aws.Config, client *schemas.Client) {
	input := &schemas.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_schemasResourceArn) > 0 {
		input.ResourceArn = aws.String(_schemasResourceArn)
	}
	if len(_schemasTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _schemasTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the discoverer
func schemas_UpdateDiscoverer(cfg aws.Config, client *schemas.Client) {
	input := &schemas.UpdateDiscovererInput{
		// DiscovererId: *string, // Required
	}

	if len(_schemasDiscovererId) > 0 {
		input.DiscovererId = aws.String(_schemasDiscovererId)
	}
	if len(_schemasCrossAccount) > 0 {
		if err := assignInputField(input, "CrossAccount", _schemasCrossAccount); err != nil {
			log.Errorf("invalid --cross-account: %s", err.Error())
			return
		}
	}
	if len(_schemasDescription) > 0 {
		input.Description = aws.String(_schemasDescription)
	}

	if resp, err := client.UpdateDiscoverer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a registry.
func schemas_UpdateRegistry(cfg aws.Config, client *schemas.Client) {
	input := &schemas.UpdateRegistryInput{
		// RegistryName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasDescription) > 0 {
		input.Description = aws.String(_schemasDescription)
	}

	if resp, err := client.UpdateRegistry(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the schema definition
// Inactive schemas will be deleted after two years.
func schemas_UpdateSchema(cfg aws.Config, client *schemas.Client) {
	input := &schemas.UpdateSchemaInput{
		// RegistryName: *string, // Required
		// SchemaName: *string, // Required
	}

	if len(_schemasRegistryName) > 0 {
		input.RegistryName = aws.String(_schemasRegistryName)
	}
	if len(_schemasSchemaName) > 0 {
		input.SchemaName = aws.String(_schemasSchemaName)
	}
	if len(_schemasClientTokenId) > 0 {
		input.ClientTokenId = aws.String(_schemasClientTokenId)
	}
	if len(_schemasContent) > 0 {
		input.Content = aws.String(_schemasContent)
	}
	if len(_schemasDescription) > 0 {
		input.Description = aws.String(_schemasDescription)
	}
	if len(_schemasType) > 0 {
		if err := assignInputField(input, "Type", _schemasType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSchema(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_schemasCmd)
	_schemasCmd.Flags().SortFlags = false

	_schemasCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_schemasCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_schemasCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_schemasCmd.Flags().StringVarP(&_schemasClientTokenId, "client-token-id", "", "", "Client Token ID")
	_schemasCmd.Flags().StringVarP(&_schemasContent, "content", "", "", "Content")
	_schemasCmd.Flags().StringVarP(&_schemasCrossAccount, "cross-account", "", "", "Cross Account")
	_schemasCmd.Flags().StringVarP(&_schemasDescription, "description", "", "", "Description")
	_schemasCmd.Flags().StringVarP(&_schemasDiscovererId, "discoverer-id", "", "", "Discoverer ID")
	_schemasCmd.Flags().StringVarP(&_schemasDiscovererIdPrefix, "discoverer-id-prefix", "", "", "Discoverer ID Prefix")
	_schemasCmd.Flags().StringSliceVarP(&_schemasEvents, "events", "", nil, "Events")
	_schemasCmd.Flags().StringVarP(&_schemasKeywords, "keywords", "", "", "Keywords")
	_schemasCmd.Flags().StringVarP(&_schemasLanguage, "language", "", "", "Language")
	_schemasCmd.Flags().StringVarP(&_schemasLimit, "limit", "", "", "Limit")
	_schemasCmd.Flags().StringVarP(&_schemasNextToken, "next-token", "", "", "Next Token")
	_schemasCmd.Flags().StringVarP(&_schemasPolicy, "policy", "", "", "Policy")
	_schemasCmd.Flags().StringVarP(&_schemasRegistryName, "registry-name", "", "", "Registry Name")
	_schemasCmd.Flags().StringVarP(&_schemasRegistryNamePrefix, "registry-name-prefix", "", "", "Registry Name Prefix")
	_schemasCmd.Flags().StringVarP(&_schemasResourceArn, "resource-arn", "", "", "Resource ARN")
	_schemasCmd.Flags().StringVarP(&_schemasRevisionId, "revision-id", "", "", "Revision ID")
	_schemasCmd.Flags().StringVarP(&_schemasSchemaName, "schema-name", "", "", "Schema Name")
	_schemasCmd.Flags().StringVarP(&_schemasSchemaNamePrefix, "schema-name-prefix", "", "", "Schema Name Prefix")
	_schemasCmd.Flags().StringVarP(&_schemasSchemaVersion, "schema-version", "", "", "Schema Version")
	_schemasCmd.Flags().StringVarP(&_schemasScope, "scope", "", "", "Scope")
	_schemasCmd.Flags().StringVarP(&_schemasSourceArn, "source-arn", "", "", "Source ARN")
	_schemasCmd.Flags().StringVarP(&_schemasSourceArnPrefix, "source-arn-prefix", "", "", "Source ARN Prefix")
	_schemasCmd.Flags().StringSliceVarP(&_schemasTagKeys, "tag-keys", "", nil, "Tag Keys")
	_schemasCmd.Flags().StringVarP(&_schemasTags, "tags", "", "", "Tags")
	_schemasCmd.Flags().StringVarP(&_schemasType, "type", "", "", "Type")

	_schemasCmd.Flags().BoolVarP(&_schemasCreateDiscoverer, "create-discoverer", "", false, "Create Discoverer")
	_schemasCmd.Flags().BoolVarP(&_schemasCreateRegistry, "create-registry", "", false, "Create Registry")
	_schemasCmd.Flags().BoolVarP(&_schemasCreateSchema, "create-schema", "", false, "Create Schema")
	_schemasCmd.Flags().BoolVarP(&_schemasDeleteDiscoverer, "delete-discoverer", "", false, "Delete Discoverer")
	_schemasCmd.Flags().BoolVarP(&_schemasDeleteRegistry, "delete-registry", "", false, "Delete Registry")
	_schemasCmd.Flags().BoolVarP(&_schemasDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_schemasCmd.Flags().BoolVarP(&_schemasDeleteSchema, "delete-schema", "", false, "Delete Schema")
	_schemasCmd.Flags().BoolVarP(&_schemasDeleteSchemaVersion, "delete-schema-version", "", false, "Delete Schema Version")
	_schemasCmd.Flags().BoolVarP(&_schemasDescribeCodeBinding, "describe-code-binding", "", false, "Describe Code Binding")
	_schemasCmd.Flags().BoolVarP(&_schemasDescribeDiscoverer, "describe-discoverer", "", false, "Describe Discoverer")
	_schemasCmd.Flags().BoolVarP(&_schemasDescribeRegistry, "describe-registry", "", false, "Describe Registry")
	_schemasCmd.Flags().BoolVarP(&_schemasDescribeSchema, "describe-schema", "", false, "Describe Schema")
	_schemasCmd.Flags().BoolVarP(&_schemasExportSchema, "export-schema", "", false, "Export Schema")
	_schemasCmd.Flags().BoolVarP(&_schemasGetCodeBindingSource, "get-code-binding-source", "", false, "Get Code Binding Source")
	_schemasCmd.Flags().BoolVarP(&_schemasGetDiscoveredSchema, "get-discovered-schema", "", false, "Get Discovered Schema")
	_schemasCmd.Flags().BoolVarP(&_schemasGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_schemasCmd.Flags().BoolVarP(&_schemasListDiscoverers, "list-discoverers", "", false, "List Discoverers")
	_schemasCmd.Flags().BoolVarP(&_schemasListRegistries, "list-registries", "", false, "List Registries")
	_schemasCmd.Flags().BoolVarP(&_schemasListSchemaVersions, "list-schema-versions", "", false, "List Schema Versions")
	_schemasCmd.Flags().BoolVarP(&_schemasListSchemas, "list-schemas", "", false, "List Schemas")
	_schemasCmd.Flags().BoolVarP(&_schemasListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_schemasCmd.Flags().BoolVarP(&_schemasPutCodeBinding, "put-code-binding", "", false, "Put Code Binding")
	_schemasCmd.Flags().BoolVarP(&_schemasPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_schemasCmd.Flags().BoolVarP(&_schemasSearchSchemas, "search-schemas", "", false, "Search Schemas")
	_schemasCmd.Flags().BoolVarP(&_schemasStartDiscoverer, "start-discoverer", "", false, "Start Discoverer")
	_schemasCmd.Flags().BoolVarP(&_schemasStopDiscoverer, "stop-discoverer", "", false, "Stop Discoverer")
	_schemasCmd.Flags().BoolVarP(&_schemasTagResource, "tag-resource", "", false, "Tag Resource")
	_schemasCmd.Flags().BoolVarP(&_schemasUntagResource, "untag-resource", "", false, "Untag Resource")
	_schemasCmd.Flags().BoolVarP(&_schemasUpdateDiscoverer, "update-discoverer", "", false, "Update Discoverer")
	_schemasCmd.Flags().BoolVarP(&_schemasUpdateRegistry, "update-registry", "", false, "Update Registry")
	_schemasCmd.Flags().BoolVarP(&_schemasUpdateSchema, "update-schema", "", false, "Update Schema")

}
