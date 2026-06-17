package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/keyspaces"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// keyspacesCmd represents the keyspaces command
var _keyspacesCmd = &cobra.Command{
	Use:   "keyspaces",
	Short: "AWS keyspaces CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := keyspaces.NewFromConfig(cfg)
		if _keyspacesCreateKeyspace {
			keyspaces_CreateKeyspace(cfg, client)
			return
		}
		if _keyspacesCreateTable {
			keyspaces_CreateTable(cfg, client)
			return
		}
		if _keyspacesCreateType {
			keyspaces_CreateType(cfg, client)
			return
		}
		if _keyspacesDeleteKeyspace {
			keyspaces_DeleteKeyspace(cfg, client)
			return
		}
		if _keyspacesDeleteTable {
			keyspaces_DeleteTable(cfg, client)
			return
		}
		if _keyspacesDeleteType {
			keyspaces_DeleteType(cfg, client)
			return
		}
		if _keyspacesGetKeyspace {
			keyspaces_GetKeyspace(cfg, client)
			return
		}
		if _keyspacesGetTable {
			keyspaces_GetTable(cfg, client)
			return
		}
		if _keyspacesGetTableAutoScalingSettings {
			keyspaces_GetTableAutoScalingSettings(cfg, client)
			return
		}
		if _keyspacesGetType {
			keyspaces_GetType(cfg, client)
			return
		}
		if _keyspacesListKeyspaces {
			keyspaces_ListKeyspaces(cfg, client)
			return
		}
		if _keyspacesListTables {
			keyspaces_ListTables(cfg, client)
			return
		}
		if _keyspacesListTagsForResource {
			keyspaces_ListTagsForResource(cfg, client)
			return
		}
		if _keyspacesListTypes {
			keyspaces_ListTypes(cfg, client)
			return
		}
		if _keyspacesRestoreTable {
			keyspaces_RestoreTable(cfg, client)
			return
		}
		if _keyspacesTagResource {
			keyspaces_TagResource(cfg, client)
			return
		}
		if _keyspacesUntagResource {
			keyspaces_UntagResource(cfg, client)
			return
		}
		if _keyspacesUpdateKeyspace {
			keyspaces_UpdateKeyspace(cfg, client)
			return
		}
		if _keyspacesUpdateTable {
			keyspaces_UpdateTable(cfg, client)
			return
		}

	},
}

var (
	_keyspacesCreateKeyspace              bool
	_keyspacesCreateTable                 bool
	_keyspacesCreateType                  bool
	_keyspacesDeleteKeyspace              bool
	_keyspacesDeleteTable                 bool
	_keyspacesDeleteType                  bool
	_keyspacesGetKeyspace                 bool
	_keyspacesGetTable                    bool
	_keyspacesGetTableAutoScalingSettings bool
	_keyspacesGetType                     bool
	_keyspacesListKeyspaces               bool
	_keyspacesListTables                  bool
	_keyspacesListTagsForResource         bool
	_keyspacesListTypes                   bool
	_keyspacesRestoreTable                bool
	_keyspacesTagResource                 bool
	_keyspacesUntagResource               bool
	_keyspacesUpdateKeyspace              bool
	_keyspacesUpdateTable                 bool

	_keyspacesAddColumns                      string
	_keyspacesAutoScalingSpecification        string
	_keyspacesCapacitySpecification           string
	_keyspacesCapacitySpecificationOverride   string
	_keyspacesCdcSpecification                string
	_keyspacesClientSideTimestamps            string
	_keyspacesComment                         string
	_keyspacesDefaultTimeToLive               string
	_keyspacesEncryptionSpecification         string
	_keyspacesEncryptionSpecificationOverride string
	_keyspacesFieldDefinitions                string
	_keyspacesKeyspaceName                    string
	_keyspacesMaxResults                      string
	_keyspacesNextToken                       string
	_keyspacesPointInTimeRecovery             string
	_keyspacesPointInTimeRecoveryOverride     string
	_keyspacesReplicaSpecifications           string
	_keyspacesReplicationSpecification        string
	_keyspacesResourceArn                     string
	_keyspacesRestoreTimestamp                string
	_keyspacesSchemaDefinition                string
	_keyspacesSourceKeyspaceName              string
	_keyspacesSourceTableName                 string
	_keyspacesTableName                       string
	_keyspacesTags                            string
	_keyspacesTagsOverride                    string
	_keyspacesTargetKeyspaceName              string
	_keyspacesTargetTableName                 string
	_keyspacesTtl                             string
	_keyspacesTypeName                        string
	_keyspacesWarmThroughputSpecification     string
)

// The CreateKeyspace operation adds a new keyspace to your account. In an Amazon
// Web Services account, keyspace names must be unique within each Region.
//
// CreateKeyspace is an asynchronous operation. You can monitor the creation
// status of the new keyspace by using the GetKeyspace operation.
//
// For more information, see [Create a keyspace] in the Amazon Keyspaces Developer Guide.
//
// [Create a keyspace]: https://docs.aws.amazon.com/keyspaces/latest/devguide/getting-started.keyspaces.html
func keyspaces_CreateKeyspace(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.CreateKeyspaceInput{
		// KeyspaceName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesReplicationSpecification) > 0 {
		if err := assignInputField(input, "ReplicationSpecification", _keyspacesReplicationSpecification); err != nil {
			log.Errorf("invalid --replication-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _keyspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKeyspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateTable operation adds a new table to the specified keyspace. Within a
// keyspace, table names must be unique.
//
// CreateTable is an asynchronous operation. When the request is received, the
// status of the table is set to CREATING . You can monitor the creation status of
// the new table by using the GetTable operation, which returns the current status
// of the table. You can start using a table when the status is ACTIVE .
//
// For more information, see [Create a table] in the Amazon Keyspaces Developer Guide.
//
// [Create a table]: https://docs.aws.amazon.com/keyspaces/latest/devguide/getting-started.tables.html
func keyspaces_CreateTable(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.CreateTableInput{
		// KeyspaceName: *string, // Required
		// SchemaDefinition: *types.SchemaDefinition, // Required
		// TableName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesSchemaDefinition) > 0 {
		if err := assignInputField(input, "SchemaDefinition", _keyspacesSchemaDefinition); err != nil {
			log.Errorf("invalid --schema-definition: %s", err.Error())
			return
		}
	}
	if len(_keyspacesTableName) > 0 {
		input.TableName = aws.String(_keyspacesTableName)
	}
	if len(_keyspacesAutoScalingSpecification) > 0 {
		if err := assignInputField(input, "AutoScalingSpecification", _keyspacesAutoScalingSpecification); err != nil {
			log.Errorf("invalid --auto-scaling-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesCapacitySpecification) > 0 {
		if err := assignInputField(input, "CapacitySpecification", _keyspacesCapacitySpecification); err != nil {
			log.Errorf("invalid --capacity-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesCdcSpecification) > 0 {
		if err := assignInputField(input, "CdcSpecification", _keyspacesCdcSpecification); err != nil {
			log.Errorf("invalid --cdc-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesClientSideTimestamps) > 0 {
		if err := assignInputField(input, "ClientSideTimestamps", _keyspacesClientSideTimestamps); err != nil {
			log.Errorf("invalid --client-side-timestamps: %s", err.Error())
			return
		}
	}
	if len(_keyspacesComment) > 0 {
		if err := assignInputField(input, "Comment", _keyspacesComment); err != nil {
			log.Errorf("invalid --comment: %s", err.Error())
			return
		}
	}
	if len(_keyspacesDefaultTimeToLive) > 0 {
		if err := assignInputField(input, "DefaultTimeToLive", _keyspacesDefaultTimeToLive); err != nil {
			log.Errorf("invalid --default-time-to-live: %s", err.Error())
			return
		}
	}
	if len(_keyspacesEncryptionSpecification) > 0 {
		if err := assignInputField(input, "EncryptionSpecification", _keyspacesEncryptionSpecification); err != nil {
			log.Errorf("invalid --encryption-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesPointInTimeRecovery) > 0 {
		if err := assignInputField(input, "PointInTimeRecovery", _keyspacesPointInTimeRecovery); err != nil {
			log.Errorf("invalid --point-in-time-recovery: %s", err.Error())
			return
		}
	}
	if len(_keyspacesReplicaSpecifications) > 0 {
		if err := assignInputField(input, "ReplicaSpecifications", _keyspacesReplicaSpecifications); err != nil {
			log.Errorf("invalid --replica-specifications: %s", err.Error())
			return
		}
	}
	if len(_keyspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _keyspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_keyspacesTtl) > 0 {
		if err := assignInputField(input, "Ttl", _keyspacesTtl); err != nil {
			log.Errorf("invalid --ttl: %s", err.Error())
			return
		}
	}
	if len(_keyspacesWarmThroughputSpecification) > 0 {
		if err := assignInputField(input, "WarmThroughputSpecification", _keyspacesWarmThroughputSpecification); err != nil {
			log.Errorf("invalid --warm-throughput-specification: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The CreateType operation creates a new user-defined type in the specified
// keyspace.
//
// To configure the required permissions, see [Permissions to create a UDT] in the Amazon Keyspaces Developer
// Guide.
//
// For more information, see [User-defined types (UDTs)] in the Amazon Keyspaces Developer Guide.
//
// [User-defined types (UDTs)]: https://docs.aws.amazon.com/keyspaces/latest/devguide/udts.html
// [Permissions to create a UDT]: https://docs.aws.amazon.com/keyspaces/latest/devguide/configure-udt-permissions.html#udt-permissions-create
func keyspaces_CreateType(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.CreateTypeInput{
		// FieldDefinitions: []types.FieldDefinition, // Required
		// KeyspaceName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_keyspacesFieldDefinitions) > 0 {
		if err := assignInputField(input, "FieldDefinitions", _keyspacesFieldDefinitions); err != nil {
			log.Errorf("invalid --field-definitions: %s", err.Error())
			return
		}
	}
	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTypeName) > 0 {
		input.TypeName = aws.String(_keyspacesTypeName)
	}

	if resp, err := client.CreateType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteKeyspace operation deletes a keyspace and all of its tables.
func keyspaces_DeleteKeyspace(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.DeleteKeyspaceInput{
		// KeyspaceName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}

	if resp, err := client.DeleteKeyspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteTable operation deletes a table and all of its data. After a
// DeleteTable request is received, the specified table is in the DELETING state
// until Amazon Keyspaces completes the deletion. If the table is in the ACTIVE
// state, you can delete it. If a table is either in the CREATING or UPDATING
// states, then Amazon Keyspaces returns a ResourceInUseException . If the
// specified table does not exist, Amazon Keyspaces returns a
// ResourceNotFoundException . If the table is already in the DELETING state, no
// error is returned.
func keyspaces_DeleteTable(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.DeleteTableInput{
		// KeyspaceName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTableName) > 0 {
		input.TableName = aws.String(_keyspacesTableName)
	}

	if resp, err := client.DeleteTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteType operation deletes a user-defined type (UDT). You can only
// delete a type that is not used in a table or another UDT.
//
// To configure the required permissions, see [Permissions to delete a UDT] in the Amazon Keyspaces Developer
// Guide.
//
// [Permissions to delete a UDT]: https://docs.aws.amazon.com/keyspaces/latest/devguide/configure-udt-permissions.html#udt-permissions-drop
func keyspaces_DeleteType(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.DeleteTypeInput{
		// KeyspaceName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTypeName) > 0 {
		input.TypeName = aws.String(_keyspacesTypeName)
	}

	if resp, err := client.DeleteType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the name of the specified keyspace, the Amazon Resource Name (ARN), the
// replication strategy, the Amazon Web Services Regions of a multi-Region
// keyspace, and the status of newly added Regions after an UpdateKeyspace
// operation.
func keyspaces_GetKeyspace(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.GetKeyspaceInput{
		// KeyspaceName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}

	if resp, err := client.GetKeyspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the table, including the table's name and current
// status, the keyspace name, configuration settings, and metadata.
//
// To read table metadata using GetTable , the IAM principal needs Select action
// permissions for the table and the system keyspace.
func keyspaces_GetTable(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.GetTableInput{
		// KeyspaceName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTableName) > 0 {
		input.TableName = aws.String(_keyspacesTableName)
	}

	if resp, err := client.GetTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns auto scaling related settings of the specified table in JSON format. If
// the table is a multi-Region table, the Amazon Web Services Region specific auto
// scaling settings of the table are included.
//
// Amazon Keyspaces auto scaling helps you provision throughput capacity for
// variable workloads efficiently by increasing and decreasing your table's read
// and write capacity automatically in response to application traffic. For more
// information, see [Managing throughput capacity automatically with Amazon Keyspaces auto scaling]in the Amazon Keyspaces Developer Guide.
//
// GetTableAutoScalingSettings can't be used as an action in an IAM policy.
//
// To define permissions for GetTableAutoScalingSettings , you must allow the
// following two actions in the IAM policy statement's Action element:
//
// - application-autoscaling:DescribeScalableTargets
//
// - application-autoscaling:DescribeScalingPolicies
//
// [Managing throughput capacity automatically with Amazon Keyspaces auto scaling]: https://docs.aws.amazon.com/keyspaces/latest/devguide/autoscaling.html
func keyspaces_GetTableAutoScalingSettings(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.GetTableAutoScalingSettingsInput{
		// KeyspaceName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTableName) > 0 {
		input.TableName = aws.String(_keyspacesTableName)
	}

	if resp, err := client.GetTableAutoScalingSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The GetType operation returns information about the type, for example the
// field definitions, the timestamp when the type was last modified, the level of
// nesting, the status, and details about if the type is used in other types and
// tables.
//
// To read keyspace metadata using GetType , the IAM principal needs Select action
// permissions for the system keyspace. To configure the required permissions, see [Permissions to view a UDT]
// in the Amazon Keyspaces Developer Guide.
//
// [Permissions to view a UDT]: https://docs.aws.amazon.com/keyspaces/latest/devguide/configure-udt-permissions.html#udt-permissions-view
func keyspaces_GetType(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.GetTypeInput{
		// KeyspaceName: *string, // Required
		// TypeName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTypeName) > 0 {
		input.TypeName = aws.String(_keyspacesTypeName)
	}

	if resp, err := client.GetType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The ListKeyspaces operation returns a list of keyspaces.
func keyspaces_ListKeyspaces(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.ListKeyspacesInput{}

	if len(_keyspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_keyspacesNextToken) > 0 {
		input.NextToken = aws.String(_keyspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKeyspaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*keyspaces.ListKeyspacesOutput
	p := keyspaces.NewListKeyspacesPaginator(client, input)
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

// The ListTables operation returns a list of tables for a specified keyspace.
// To read keyspace metadata using ListTables , the IAM principal needs Select
// action permissions for the system keyspace.
func keyspaces_ListTables(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.ListTablesInput{
		// KeyspaceName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_keyspacesNextToken) > 0 {
		input.NextToken = aws.String(_keyspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTables(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*keyspaces.ListTablesOutput
	p := keyspaces.NewListTablesPaginator(client, input)
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

// Returns a list of all tags associated with the specified Amazon Keyspaces
// resource.
//
// To read keyspace metadata using ListTagsForResource , the IAM principal needs
// Select action permissions for the specified resource and the system keyspace.
func keyspaces_ListTagsForResource(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_keyspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_keyspacesResourceArn)
	}
	if len(_keyspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_keyspacesNextToken) > 0 {
		input.NextToken = aws.String(_keyspacesNextToken)
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

	var results []*keyspaces.ListTagsForResourceOutput
	p := keyspaces.NewListTagsForResourcePaginator(client, input)
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

// The ListTypes operation returns a list of types for a specified keyspace.
// To read keyspace metadata using ListTypes , the IAM principal needs Select
// action permissions for the system keyspace. To configure the required
// permissions, see [Permissions to view a UDT]in the Amazon Keyspaces Developer Guide.
//
// [Permissions to view a UDT]: https://docs.aws.amazon.com/keyspaces/latest/devguide/configure-udt-permissions.html#udt-permissions-view
func keyspaces_ListTypes(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.ListTypesInput{
		// KeyspaceName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _keyspacesMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_keyspacesNextToken) > 0 {
		input.NextToken = aws.String(_keyspacesNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*keyspaces.ListTypesOutput
	p := keyspaces.NewListTypesPaginator(client, input)
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

// Restores the table to the specified point in time within the
// earliest_restorable_timestamp and the current time. For more information about
// restore points, see [Time window for PITR continuous backups]in the Amazon Keyspaces Developer Guide.
//
// Any number of users can execute up to 4 concurrent restores (any type of
// restore) in a given account.
//
// When you restore using point in time recovery, Amazon Keyspaces restores your
// source table's schema and data to the state based on the selected timestamp
// (day:hour:minute:second) to a new table. The Time to Live (TTL) settings are
// also restored to the state based on the selected timestamp.
//
// In addition to the table's schema, data, and TTL settings, RestoreTable
// restores the capacity mode, auto scaling settings, encryption settings, and
// point-in-time recovery settings from the source table. Unlike the table's schema
// data and TTL settings, which are restored based on the selected timestamp, these
// settings are always restored based on the table's settings as of the current
// time or when the table was deleted.
//
// You can also overwrite these settings during restore:
//
// - Read/write capacity mode
//
// - Provisioned throughput capacity units
//
// - Auto scaling settings
//
// - Point-in-time (PITR) settings
//
// - Tags
//
// For more information, see [PITR restore settings] in the Amazon Keyspaces Developer Guide.
//
// Note that the following settings are not restored, and you must configure them
// manually for the new table:
//
// - Identity and Access Management (IAM) policies
//
// - Amazon CloudWatch metrics and alarms
//
// [PITR restore settings]: https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery_HowItWorks.html#howitworks_backup_settings
// [Time window for PITR continuous backups]: https://docs.aws.amazon.com/keyspaces/latest/devguide/PointInTimeRecovery_HowItWorks.html#howitworks_backup_window
func keyspaces_RestoreTable(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.RestoreTableInput{
		// SourceKeyspaceName: *string, // Required
		// SourceTableName: *string, // Required
		// TargetKeyspaceName: *string, // Required
		// TargetTableName: *string, // Required
	}

	if len(_keyspacesSourceKeyspaceName) > 0 {
		input.SourceKeyspaceName = aws.String(_keyspacesSourceKeyspaceName)
	}
	if len(_keyspacesSourceTableName) > 0 {
		input.SourceTableName = aws.String(_keyspacesSourceTableName)
	}
	if len(_keyspacesTargetKeyspaceName) > 0 {
		input.TargetKeyspaceName = aws.String(_keyspacesTargetKeyspaceName)
	}
	if len(_keyspacesTargetTableName) > 0 {
		input.TargetTableName = aws.String(_keyspacesTargetTableName)
	}
	if len(_keyspacesAutoScalingSpecification) > 0 {
		if err := assignInputField(input, "AutoScalingSpecification", _keyspacesAutoScalingSpecification); err != nil {
			log.Errorf("invalid --auto-scaling-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesCapacitySpecificationOverride) > 0 {
		if err := assignInputField(input, "CapacitySpecificationOverride", _keyspacesCapacitySpecificationOverride); err != nil {
			log.Errorf("invalid --capacity-specification-override: %s", err.Error())
			return
		}
	}
	if len(_keyspacesEncryptionSpecificationOverride) > 0 {
		if err := assignInputField(input, "EncryptionSpecificationOverride", _keyspacesEncryptionSpecificationOverride); err != nil {
			log.Errorf("invalid --encryption-specification-override: %s", err.Error())
			return
		}
	}
	if len(_keyspacesPointInTimeRecoveryOverride) > 0 {
		if err := assignInputField(input, "PointInTimeRecoveryOverride", _keyspacesPointInTimeRecoveryOverride); err != nil {
			log.Errorf("invalid --point-in-time-recovery-override: %s", err.Error())
			return
		}
	}
	if len(_keyspacesReplicaSpecifications) > 0 {
		if err := assignInputField(input, "ReplicaSpecifications", _keyspacesReplicaSpecifications); err != nil {
			log.Errorf("invalid --replica-specifications: %s", err.Error())
			return
		}
	}
	if len(_keyspacesRestoreTimestamp) > 0 {
		if err := assignInputField(input, "RestoreTimestamp", _keyspacesRestoreTimestamp); err != nil {
			log.Errorf("invalid --restore-timestamp: %s", err.Error())
			return
		}
	}
	if len(_keyspacesTagsOverride) > 0 {
		if err := assignInputField(input, "TagsOverride", _keyspacesTagsOverride); err != nil {
			log.Errorf("invalid --tags-override: %s", err.Error())
			return
		}
	}

	if resp, err := client.RestoreTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a set of tags with a Amazon Keyspaces resource. You can then
// activate these user-defined tags so that they appear on the Cost Management
// Console for cost allocation tracking. For more information, see [Adding tags and labels to Amazon Keyspaces resources]in the Amazon
// Keyspaces Developer Guide.
//
// For IAM policy examples that show how to control access to Amazon Keyspaces
// resources based on tags, see [Amazon Keyspaces resource access based on tags]in the Amazon Keyspaces Developer Guide.
//
// [Amazon Keyspaces resource access based on tags]: https://docs.aws.amazon.com/keyspaces/latest/devguide/security_iam_id-based-policy-examples.html#security_iam_id-based-policy-examples-tags
// [Adding tags and labels to Amazon Keyspaces resources]: https://docs.aws.amazon.com/keyspaces/latest/devguide/tagging-keyspaces.html
func keyspaces_TagResource(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_keyspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_keyspacesResourceArn)
	}
	if len(_keyspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _keyspacesTags); err != nil {
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

// Removes the association of tags from a Amazon Keyspaces resource.
func keyspaces_UntagResource(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.UntagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_keyspacesResourceArn) > 0 {
		input.ResourceArn = aws.String(_keyspacesResourceArn)
	}
	if len(_keyspacesTags) > 0 {
		if err := assignInputField(input, "Tags", _keyspacesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new Amazon Web Services Region to the keyspace. You can add a new
// Region to a keyspace that is either a single or a multi-Region keyspace. Amazon
// Keyspaces is going to replicate all tables in the keyspace to the new Region. To
// successfully replicate all tables to the new Region, they must use client-side
// timestamps for conflict resolution. To enable client-side timestamps, specify
// clientSideTimestamps.status = enabled when invoking the API. For more
// information about client-side timestamps, see [Client-side timestamps in Amazon Keyspaces]in the Amazon Keyspaces Developer
// Guide.
//
// To add a Region to a keyspace using the UpdateKeyspace API, the IAM principal
// needs permissions for the following IAM actions:
//
// - cassandra:Alter
//
// - cassandra:AlterMultiRegionResource
//
// - cassandra:Create
//
// - cassandra:CreateMultiRegionResource
//
// - cassandra:Select
//
// - cassandra:SelectMultiRegionResource
//
// - cassandra:Modify
//
// - cassandra:ModifyMultiRegionResource
//
// If the keyspace contains a table that is configured in provisioned mode with
// auto scaling enabled, the following additional IAM actions need to be allowed.
//
// - application-autoscaling:RegisterScalableTarget
//
// - application-autoscaling:DeregisterScalableTarget
//
// - application-autoscaling:DescribeScalableTargets
//
// - application-autoscaling:PutScalingPolicy
//
// - application-autoscaling:DescribeScalingPolicies
//
// To use the UpdateKeyspace API, the IAM principal also needs permissions to
// create a service-linked role with the following elements:
//
// - iam:CreateServiceLinkedRole - The action the principal can perform.
//
// -
// arn:aws:iam::*:role/aws-service-role/replication.cassandra.amazonaws.com/AWSServiceRoleForKeyspacesReplication
//
// - The resource that the action can be performed on.
//
// - iam:AWSServiceName: replication.cassandra.amazonaws.com - The only Amazon
// Web Services service that this role can be attached to is Amazon Keyspaces.
//
// For more information, see [Configure the IAM permissions required to add an Amazon Web Services Region to a keyspace] in the Amazon Keyspaces Developer Guide.
//
// [Client-side timestamps in Amazon Keyspaces]: https://docs.aws.amazon.com/keyspaces/latest/devguide/client-side-timestamps.html
// [Configure the IAM permissions required to add an Amazon Web Services Region to a keyspace]: https://docs.aws.amazon.com/keyspaces/latest/devguide/howitworks_replication_permissions_addReplica.html
func keyspaces_UpdateKeyspace(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.UpdateKeyspaceInput{
		// KeyspaceName: *string, // Required
		// ReplicationSpecification: *types.ReplicationSpecification, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesReplicationSpecification) > 0 {
		if err := assignInputField(input, "ReplicationSpecification", _keyspacesReplicationSpecification); err != nil {
			log.Errorf("invalid --replication-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesClientSideTimestamps) > 0 {
		if err := assignInputField(input, "ClientSideTimestamps", _keyspacesClientSideTimestamps); err != nil {
			log.Errorf("invalid --client-side-timestamps: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKeyspace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds new columns to the table or updates one of the table's settings, for
// example capacity mode, auto scaling, encryption, point-in-time recovery, or ttl
// settings. Note that you can only update one specific table setting per update
// operation.
func keyspaces_UpdateTable(cfg aws.Config, client *keyspaces.Client) {
	input := &keyspaces.UpdateTableInput{
		// KeyspaceName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_keyspacesKeyspaceName) > 0 {
		input.KeyspaceName = aws.String(_keyspacesKeyspaceName)
	}
	if len(_keyspacesTableName) > 0 {
		input.TableName = aws.String(_keyspacesTableName)
	}
	if len(_keyspacesAddColumns) > 0 {
		if err := assignInputField(input, "AddColumns", _keyspacesAddColumns); err != nil {
			log.Errorf("invalid --add-columns: %s", err.Error())
			return
		}
	}
	if len(_keyspacesAutoScalingSpecification) > 0 {
		if err := assignInputField(input, "AutoScalingSpecification", _keyspacesAutoScalingSpecification); err != nil {
			log.Errorf("invalid --auto-scaling-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesCapacitySpecification) > 0 {
		if err := assignInputField(input, "CapacitySpecification", _keyspacesCapacitySpecification); err != nil {
			log.Errorf("invalid --capacity-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesCdcSpecification) > 0 {
		if err := assignInputField(input, "CdcSpecification", _keyspacesCdcSpecification); err != nil {
			log.Errorf("invalid --cdc-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesClientSideTimestamps) > 0 {
		if err := assignInputField(input, "ClientSideTimestamps", _keyspacesClientSideTimestamps); err != nil {
			log.Errorf("invalid --client-side-timestamps: %s", err.Error())
			return
		}
	}
	if len(_keyspacesDefaultTimeToLive) > 0 {
		if err := assignInputField(input, "DefaultTimeToLive", _keyspacesDefaultTimeToLive); err != nil {
			log.Errorf("invalid --default-time-to-live: %s", err.Error())
			return
		}
	}
	if len(_keyspacesEncryptionSpecification) > 0 {
		if err := assignInputField(input, "EncryptionSpecification", _keyspacesEncryptionSpecification); err != nil {
			log.Errorf("invalid --encryption-specification: %s", err.Error())
			return
		}
	}
	if len(_keyspacesPointInTimeRecovery) > 0 {
		if err := assignInputField(input, "PointInTimeRecovery", _keyspacesPointInTimeRecovery); err != nil {
			log.Errorf("invalid --point-in-time-recovery: %s", err.Error())
			return
		}
	}
	if len(_keyspacesReplicaSpecifications) > 0 {
		if err := assignInputField(input, "ReplicaSpecifications", _keyspacesReplicaSpecifications); err != nil {
			log.Errorf("invalid --replica-specifications: %s", err.Error())
			return
		}
	}
	if len(_keyspacesTtl) > 0 {
		if err := assignInputField(input, "Ttl", _keyspacesTtl); err != nil {
			log.Errorf("invalid --ttl: %s", err.Error())
			return
		}
	}
	if len(_keyspacesWarmThroughputSpecification) > 0 {
		if err := assignInputField(input, "WarmThroughputSpecification", _keyspacesWarmThroughputSpecification); err != nil {
			log.Errorf("invalid --warm-throughput-specification: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_keyspacesCmd)
	_keyspacesCmd.Flags().SortFlags = false

	_keyspacesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_keyspacesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_keyspacesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_keyspacesCmd.Flags().StringVarP(&_keyspacesAddColumns, "add-columns", "", "", "Add Columns")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesAutoScalingSpecification, "auto-scaling-specification", "", "", "Auto Scaling Specification")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesCapacitySpecification, "capacity-specification", "", "", "Capacity Specification")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesCapacitySpecificationOverride, "capacity-specification-override", "", "", "Capacity Specification Override")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesCdcSpecification, "cdc-specification", "", "", "Cdc Specification")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesClientSideTimestamps, "client-side-timestamps", "", "", "Client Side Timestamps")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesComment, "comment", "", "", "Comment")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesDefaultTimeToLive, "default-time-to-live", "", "", "Default Time To Live")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesEncryptionSpecification, "encryption-specification", "", "", "Encryption Specification")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesEncryptionSpecificationOverride, "encryption-specification-override", "", "", "Encryption Specification Override")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesFieldDefinitions, "field-definitions", "", "", "Field Definitions")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesKeyspaceName, "keyspace-name", "", "", "Keyspace Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesMaxResults, "max-results", "", "", "Max Results")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesNextToken, "next-token", "", "", "Next Token")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesPointInTimeRecovery, "point-in-time-recovery", "", "", "Point In Time Recovery")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesPointInTimeRecoveryOverride, "point-in-time-recovery-override", "", "", "Point In Time Recovery Override")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesReplicaSpecifications, "replica-specifications", "", "", "Replica Specifications")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesReplicationSpecification, "replication-specification", "", "", "Replication Specification")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesResourceArn, "resource-arn", "", "", "Resource ARN")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesRestoreTimestamp, "restore-timestamp", "", "", "Restore Timestamp")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesSchemaDefinition, "schema-definition", "", "", "Schema Definition")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesSourceKeyspaceName, "source-keyspace-name", "", "", "Source Keyspace Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesSourceTableName, "source-table-name", "", "", "Source Table Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTableName, "table-name", "", "", "Table Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTags, "tags", "", "", "Tags")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTagsOverride, "tags-override", "", "", "Tags Override")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTargetKeyspaceName, "target-keyspace-name", "", "", "Target Keyspace Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTargetTableName, "target-table-name", "", "", "Target Table Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTtl, "ttl", "", "", "TTL")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesTypeName, "type-name", "", "", "Type Name")
	_keyspacesCmd.Flags().StringVarP(&_keyspacesWarmThroughputSpecification, "warm-throughput-specification", "", "", "Warm Throughput Specification")

	_keyspacesCmd.Flags().BoolVarP(&_keyspacesCreateKeyspace, "create-keyspace", "", false, "Create Keyspace")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesCreateTable, "create-table", "", false, "Create Table")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesCreateType, "create-type", "", false, "Create Type")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesDeleteKeyspace, "delete-keyspace", "", false, "Delete Keyspace")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesDeleteTable, "delete-table", "", false, "Delete Table")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesDeleteType, "delete-type", "", false, "Delete Type")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesGetKeyspace, "get-keyspace", "", false, "Get Keyspace")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesGetTable, "get-table", "", false, "Get Table")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesGetTableAutoScalingSettings, "get-table-auto-scaling-settings", "", false, "Get Table Auto Scaling Settings")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesGetType, "get-type", "", false, "Get Type")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesListKeyspaces, "list-keyspaces", "", false, "List Keyspaces")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesListTables, "list-tables", "", false, "List Tables")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesListTypes, "list-types", "", false, "List Types")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesRestoreTable, "restore-table", "", false, "Restore Table")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesTagResource, "tag-resource", "", false, "Tag Resource")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesUntagResource, "untag-resource", "", false, "Untag Resource")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesUpdateKeyspace, "update-keyspace", "", false, "Update Keyspace")
	_keyspacesCmd.Flags().BoolVarP(&_keyspacesUpdateTable, "update-table", "", false, "Update Table")

}
