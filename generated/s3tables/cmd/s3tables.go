package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3tables"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// s3tablesCmd represents the s3tables command
var _s3tablesCmd = &cobra.Command{
	Use:   "s3tables",
	Short: "AWS s3tables CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := s3tables.NewFromConfig(cfg)
		if _s3tablesCreateNamespace {
			s3tables_CreateNamespace(cfg, client)
			return
		}
		if _s3tablesCreateTable {
			s3tables_CreateTable(cfg, client)
			return
		}
		if _s3tablesCreateTableBucket {
			s3tables_CreateTableBucket(cfg, client)
			return
		}
		if _s3tablesDeleteNamespace {
			s3tables_DeleteNamespace(cfg, client)
			return
		}
		if _s3tablesDeleteTable {
			s3tables_DeleteTable(cfg, client)
			return
		}
		if _s3tablesDeleteTableBucket {
			s3tables_DeleteTableBucket(cfg, client)
			return
		}
		if _s3tablesDeleteTableBucketEncryption {
			s3tables_DeleteTableBucketEncryption(cfg, client)
			return
		}
		if _s3tablesDeleteTableBucketMetricsConfiguration {
			s3tables_DeleteTableBucketMetricsConfiguration(cfg, client)
			return
		}
		if _s3tablesDeleteTableBucketPolicy {
			s3tables_DeleteTableBucketPolicy(cfg, client)
			return
		}
		if _s3tablesDeleteTableBucketReplication {
			s3tables_DeleteTableBucketReplication(cfg, client)
			return
		}
		if _s3tablesDeleteTablePolicy {
			s3tables_DeleteTablePolicy(cfg, client)
			return
		}
		if _s3tablesDeleteTableReplication {
			s3tables_DeleteTableReplication(cfg, client)
			return
		}
		if _s3tablesGetNamespace {
			s3tables_GetNamespace(cfg, client)
			return
		}
		if _s3tablesGetTable {
			s3tables_GetTable(cfg, client)
			return
		}
		if _s3tablesGetTableBucket {
			s3tables_GetTableBucket(cfg, client)
			return
		}
		if _s3tablesGetTableBucketEncryption {
			s3tables_GetTableBucketEncryption(cfg, client)
			return
		}
		if _s3tablesGetTableBucketMaintenanceConfiguration {
			s3tables_GetTableBucketMaintenanceConfiguration(cfg, client)
			return
		}
		if _s3tablesGetTableBucketMetricsConfiguration {
			s3tables_GetTableBucketMetricsConfiguration(cfg, client)
			return
		}
		if _s3tablesGetTableBucketPolicy {
			s3tables_GetTableBucketPolicy(cfg, client)
			return
		}
		if _s3tablesGetTableBucketReplication {
			s3tables_GetTableBucketReplication(cfg, client)
			return
		}
		if _s3tablesGetTableBucketStorageClass {
			s3tables_GetTableBucketStorageClass(cfg, client)
			return
		}
		if _s3tablesGetTableEncryption {
			s3tables_GetTableEncryption(cfg, client)
			return
		}
		if _s3tablesGetTableMaintenanceConfiguration {
			s3tables_GetTableMaintenanceConfiguration(cfg, client)
			return
		}
		if _s3tablesGetTableMaintenanceJobStatus {
			s3tables_GetTableMaintenanceJobStatus(cfg, client)
			return
		}
		if _s3tablesGetTableMetadataLocation {
			s3tables_GetTableMetadataLocation(cfg, client)
			return
		}
		if _s3tablesGetTablePolicy {
			s3tables_GetTablePolicy(cfg, client)
			return
		}
		if _s3tablesGetTableRecordExpirationConfiguration {
			s3tables_GetTableRecordExpirationConfiguration(cfg, client)
			return
		}
		if _s3tablesGetTableRecordExpirationJobStatus {
			s3tables_GetTableRecordExpirationJobStatus(cfg, client)
			return
		}
		if _s3tablesGetTableReplication {
			s3tables_GetTableReplication(cfg, client)
			return
		}
		if _s3tablesGetTableReplicationStatus {
			s3tables_GetTableReplicationStatus(cfg, client)
			return
		}
		if _s3tablesGetTableStorageClass {
			s3tables_GetTableStorageClass(cfg, client)
			return
		}
		if _s3tablesListNamespaces {
			s3tables_ListNamespaces(cfg, client)
			return
		}
		if _s3tablesListTableBuckets {
			s3tables_ListTableBuckets(cfg, client)
			return
		}
		if _s3tablesListTables {
			s3tables_ListTables(cfg, client)
			return
		}
		if _s3tablesListTagsForResource {
			s3tables_ListTagsForResource(cfg, client)
			return
		}
		if _s3tablesPutTableBucketEncryption {
			s3tables_PutTableBucketEncryption(cfg, client)
			return
		}
		if _s3tablesPutTableBucketMaintenanceConfiguration {
			s3tables_PutTableBucketMaintenanceConfiguration(cfg, client)
			return
		}
		if _s3tablesPutTableBucketMetricsConfiguration {
			s3tables_PutTableBucketMetricsConfiguration(cfg, client)
			return
		}
		if _s3tablesPutTableBucketPolicy {
			s3tables_PutTableBucketPolicy(cfg, client)
			return
		}
		if _s3tablesPutTableBucketReplication {
			s3tables_PutTableBucketReplication(cfg, client)
			return
		}
		if _s3tablesPutTableBucketStorageClass {
			s3tables_PutTableBucketStorageClass(cfg, client)
			return
		}
		if _s3tablesPutTableMaintenanceConfiguration {
			s3tables_PutTableMaintenanceConfiguration(cfg, client)
			return
		}
		if _s3tablesPutTablePolicy {
			s3tables_PutTablePolicy(cfg, client)
			return
		}
		if _s3tablesPutTableRecordExpirationConfiguration {
			s3tables_PutTableRecordExpirationConfiguration(cfg, client)
			return
		}
		if _s3tablesPutTableReplication {
			s3tables_PutTableReplication(cfg, client)
			return
		}
		if _s3tablesRenameTable {
			s3tables_RenameTable(cfg, client)
			return
		}
		if _s3tablesTagResource {
			s3tables_TagResource(cfg, client)
			return
		}
		if _s3tablesUntagResource {
			s3tables_UntagResource(cfg, client)
			return
		}
		if _s3tablesUpdateTableMetadataLocation {
			s3tables_UpdateTableMetadataLocation(cfg, client)
			return
		}

	},
}

var (
	_s3tablesCreateNamespace                        bool
	_s3tablesCreateTable                            bool
	_s3tablesCreateTableBucket                      bool
	_s3tablesDeleteNamespace                        bool
	_s3tablesDeleteTable                            bool
	_s3tablesDeleteTableBucket                      bool
	_s3tablesDeleteTableBucketEncryption            bool
	_s3tablesDeleteTableBucketMetricsConfiguration  bool
	_s3tablesDeleteTableBucketPolicy                bool
	_s3tablesDeleteTableBucketReplication           bool
	_s3tablesDeleteTablePolicy                      bool
	_s3tablesDeleteTableReplication                 bool
	_s3tablesGetNamespace                           bool
	_s3tablesGetTable                               bool
	_s3tablesGetTableBucket                         bool
	_s3tablesGetTableBucketEncryption               bool
	_s3tablesGetTableBucketMaintenanceConfiguration bool
	_s3tablesGetTableBucketMetricsConfiguration     bool
	_s3tablesGetTableBucketPolicy                   bool
	_s3tablesGetTableBucketReplication              bool
	_s3tablesGetTableBucketStorageClass             bool
	_s3tablesGetTableEncryption                     bool
	_s3tablesGetTableMaintenanceConfiguration       bool
	_s3tablesGetTableMaintenanceJobStatus           bool
	_s3tablesGetTableMetadataLocation               bool
	_s3tablesGetTablePolicy                         bool
	_s3tablesGetTableRecordExpirationConfiguration  bool
	_s3tablesGetTableRecordExpirationJobStatus      bool
	_s3tablesGetTableReplication                    bool
	_s3tablesGetTableReplicationStatus              bool
	_s3tablesGetTableStorageClass                   bool
	_s3tablesListNamespaces                         bool
	_s3tablesListTableBuckets                       bool
	_s3tablesListTables                             bool
	_s3tablesListTagsForResource                    bool
	_s3tablesPutTableBucketEncryption               bool
	_s3tablesPutTableBucketMaintenanceConfiguration bool
	_s3tablesPutTableBucketMetricsConfiguration     bool
	_s3tablesPutTableBucketPolicy                   bool
	_s3tablesPutTableBucketReplication              bool
	_s3tablesPutTableBucketStorageClass             bool
	_s3tablesPutTableMaintenanceConfiguration       bool
	_s3tablesPutTablePolicy                         bool
	_s3tablesPutTableRecordExpirationConfiguration  bool
	_s3tablesPutTableReplication                    bool
	_s3tablesRenameTable                            bool
	_s3tablesTagResource                            bool
	_s3tablesUntagResource                          bool
	_s3tablesUpdateTableMetadataLocation            bool

	_s3tablesConfiguration             string
	_s3tablesContinuationToken         string
	_s3tablesEncryptionConfiguration   string
	_s3tablesFormat                    string
	_s3tablesMaxBuckets                string
	_s3tablesMaxNamespaces             string
	_s3tablesMaxTables                 string
	_s3tablesMetadata                  string
	_s3tablesMetadataLocation          string
	_s3tablesName                      string
	_s3tablesNamespace                 string
	_s3tablesNewName                   string
	_s3tablesNewNamespaceName          string
	_s3tablesPrefix                    string
	_s3tablesResourceArn               string
	_s3tablesResourcePolicy            string
	_s3tablesStorageClassConfiguration string
	_s3tablesTableArn                  string
	_s3tablesTableBucketARN            string
	_s3tablesTagKeys                   []string
	_s3tablesTags                      string
	_s3tablesType                      string
	_s3tablesValue                     string
	_s3tablesVersionToken              string
)

// Creates a namespace. A namespace is a logical grouping of tables within your
// table bucket, which you can use to organize tables. For more information, see [Create a namespace]
// in the Amazon Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:CreateNamespace permission to use this
// operation.
//
// [Create a namespace]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-namespace-create.html
func s3tables_CreateNamespace(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.CreateNamespaceInput{
		// Namespace: []string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesNamespace) > 0 {
		input.Namespace = []string{_s3tablesNamespace}
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.CreateNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new table associated with the given namespace in a table bucket. For
// more information, see [Creating an Amazon S3 table]in the Amazon Simple Storage Service User Guide.
//
// # Permissions
//
// - You must have the s3tables:CreateTable permission to use this operation.
//
// - If you use this operation with the optional metadata request parameter you
// must have the s3tables:PutTableData permission.
//
// - If you use this operation with the optional encryptionConfiguration request
// parameter you must have the s3tables:PutTableEncryption permission.
//
// - If you use this operation with the storageClassConfiguration request
// parameter, you must have the s3tables:PutTableStorageClass permission.
//
// - To create a table with tags, you must have the s3tables:TagResource
// permission in addition to s3tables:CreateTable permission.
//
// Additionally, If you choose SSE-KMS encryption you must grant the S3 Tables
// maintenance principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption].
//
// [Creating an Amazon S3 table]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-create.html
// [Permissions requirements for S3 Tables SSE-KMS encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-kms-permissions.html
func s3tables_CreateTable(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.CreateTableInput{
		// Format: types.OpenTableFormat, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesFormat) > 0 {
		if err := assignInputField(input, "Format", _s3tablesFormat); err != nil {
			log.Errorf("invalid --format: %s", err.Error())
			return
		}
	}
	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _s3tablesEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesMetadata) > 0 {
		if err := assignInputField(input, "Metadata", _s3tablesMetadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_s3tablesStorageClassConfiguration) > 0 {
		if err := assignInputField(input, "StorageClassConfiguration", _s3tablesStorageClassConfiguration); err != nil {
			log.Errorf("invalid --storage-class-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesTags) > 0 {
		if err := assignInputField(input, "Tags", _s3tablesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Creates a table bucket. For more information, see [Creating a table bucket] in the Amazon Simple Storage
// Service User Guide.
//
// # Permissions
//
// - You must have the s3tables:CreateTableBucket permission to use this
// operation.
//
// - If you use this operation with the optional encryptionConfiguration
// parameter you must have the s3tables:PutTableBucketEncryption permission.
//
// - If you use this operation with the storageClassConfiguration request
// parameter, you must have the s3tables:PutTableBucketStorageClass permission.
//
// - To create a table bucket with tags, you must have the s3tables:TagResource
// permission in addition to s3tables:CreateTableBucket permission.
//
// [Creating a table bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-buckets-create.html
func s3tables_CreateTableBucket(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.CreateTableBucketInput{
		// Name: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _s3tablesEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesStorageClassConfiguration) > 0 {
		if err := assignInputField(input, "StorageClassConfiguration", _s3tablesStorageClassConfiguration); err != nil {
			log.Errorf("invalid --storage-class-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesTags) > 0 {
		if err := assignInputField(input, "Tags", _s3tablesTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTableBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a namespace. For more information, see [Delete a namespace] in the Amazon Simple Storage
// Service User Guide.
//
// Permissions You must have the s3tables:DeleteNamespace permission to use this
// operation.
//
// [Delete a namespace]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-namespace-delete.html
func s3tables_DeleteNamespace(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteNamespaceInput{
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.DeleteNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a table. For more information, see [Deleting an Amazon S3 table] in the Amazon Simple Storage
// Service User Guide.
//
// Permissions You must have the s3tables:DeleteTable permission to use this
// operation.
//
// [Deleting an Amazon S3 table]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-delete.html
func s3tables_DeleteTable(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.DeleteTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a table bucket. For more information, see [Deleting a table bucket] in the Amazon Simple Storage
// Service User Guide.
//
// Permissions You must have the s3tables:DeleteTableBucket permission to use this
// operation.
//
// [Deleting a table bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-buckets-delete.html
func s3tables_DeleteTableBucket(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableBucketInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.DeleteTableBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the encryption configuration for a table bucket.
// Permissions You must have the s3tables:DeleteTableBucketEncryption permission
// to use this operation.
func s3tables_DeleteTableBucketEncryption(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableBucketEncryptionInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.DeleteTableBucketEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the metrics configuration for a table bucket.
// Permissions You must have the s3tables:DeleteTableBucketMetricsConfiguration
// permission to use this operation.
func s3tables_DeleteTableBucketMetricsConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableBucketMetricsConfigurationInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.DeleteTableBucketMetricsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a table bucket policy. For more information, see [Deleting a table bucket policy] in the Amazon Simple
// Storage Service User Guide.
//
// Permissions You must have the s3tables:DeleteTableBucketPolicy permission to
// use this operation.
//
// [Deleting a table bucket policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-bucket-policy.html#table-bucket-policy-delete
func s3tables_DeleteTableBucketPolicy(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableBucketPolicyInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.DeleteTableBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the replication configuration for a table bucket. After deletion, new
// table updates will no longer be replicated to destination buckets, though
// existing replicated tables will remain in destination buckets.
//
// Permissions You must have the s3tables:DeleteTableBucketReplication permission
// to use this operation.
func s3tables_DeleteTableBucketReplication(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableBucketReplicationInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.DeleteTableBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a table policy. For more information, see [Deleting a table policy] in the Amazon Simple Storage
// Service User Guide.
//
// Permissions You must have the s3tables:DeleteTablePolicy permission to use this
// operation.
//
// [Deleting a table policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-table-policy.html#table-policy-delete
func s3tables_DeleteTablePolicy(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTablePolicyInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.DeleteTablePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the replication configuration for a specific table. After deletion, new
// updates to this table will no longer be replicated to destination tables, though
// existing replicated copies will remain in destination buckets.
//
// Permissions You must have the s3tables:DeleteTableReplication permission to use
// this operation.
func s3tables_DeleteTableReplication(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.DeleteTableReplicationInput{
		// TableArn: *string, // Required
		// VersionToken: *string, // Required
	}

	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.DeleteTableReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a namespace. For more information, see [Table namespaces] in the Amazon Simple
// Storage Service User Guide.
//
// Permissions You must have the s3tables:GetNamespace permission to use this
// operation.
//
// [Table namespaces]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-namespace.html
func s3tables_GetNamespace(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetNamespaceInput{
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a table. For more information, see [S3 Tables] in the Amazon Simple
// Storage Service User Guide.
//
// Permissions You must have the s3tables:GetTable permission to use this
// operation.
//
// [S3 Tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-tables.html
func s3tables_GetTable(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableInput{}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details on a table bucket. For more information, see [Viewing details about an Amazon S3 table bucket] in the Amazon Simple
// Storage Service User Guide.
//
// Permissions You must have the s3tables:GetTableBucket permission to use this
// operation.
//
// [Viewing details about an Amazon S3 table bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-buckets-details.html
func s3tables_GetTableBucket(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the encryption configuration for a table bucket.
// Permissions You must have the s3tables:GetTableBucketEncryption permission to
// use this operation.
func s3tables_GetTableBucketEncryption(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketEncryptionInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucketEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a maintenance configuration for a given table bucket. For
// more information, see [Amazon S3 table bucket maintenance]in the Amazon Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:GetTableBucketMaintenanceConfiguration
// permission to use this operation.
//
// [Amazon S3 table bucket maintenance]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-table-buckets-maintenance.html
func s3tables_GetTableBucketMaintenanceConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketMaintenanceConfigurationInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucketMaintenanceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the metrics configuration for a table bucket.
// Permissions You must have the s3tables:GetTableBucketMetricsConfiguration
// permission to use this operation.
func s3tables_GetTableBucketMetricsConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketMetricsConfigurationInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucketMetricsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a table bucket policy. For more information, see [Viewing a table bucket policy] in the
// Amazon Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:GetTableBucketPolicy permission to use
// this operation.
//
// [Viewing a table bucket policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-bucket-policy.html#table-bucket-policy-get
func s3tables_GetTableBucketPolicy(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketPolicyInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the replication configuration for a table bucket.This operation
// returns the IAM role, versionToken , and replication rules that define how
// tables in this bucket are replicated to other buckets.
//
// Permissions You must have the s3tables:GetTableBucketReplication permission to
// use this operation.
func s3tables_GetTableBucketReplication(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketReplicationInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the storage class configuration for a specific table. This allows you
// to view the storage class settings that apply to an individual table, which may
// differ from the table bucket's default configuration.
//
// Permissions You must have the s3tables:GetTableBucketStorageClass permission to
// use this operation.
func s3tables_GetTableBucketStorageClass(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableBucketStorageClassInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableBucketStorageClass(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the encryption configuration for a table.
// Permissions You must have the s3tables:GetTableEncryption permission to use
// this operation.
func s3tables_GetTableEncryption(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableEncryptionInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about the maintenance configuration of a table. For more
// information, see [S3 Tables maintenance]in the Amazon Simple Storage Service User Guide.
//
// # Permissions
//
// - You must have the s3tables:GetTableMaintenanceConfiguration permission to
// use this operation.
//
// - You must have the s3tables:GetTableData permission to use set the compaction
// strategy to sort or zorder .
//
// [S3 Tables maintenance]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-maintenance.html
func s3tables_GetTableMaintenanceConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableMaintenanceConfigurationInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableMaintenanceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of a maintenance job for a table. For more information, see [S3 Tables maintenance] in
// the Amazon Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:GetTableMaintenanceJobStatus permission
// to use this operation.
//
// [S3 Tables maintenance]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-maintenance.html
func s3tables_GetTableMaintenanceJobStatus(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableMaintenanceJobStatusInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableMaintenanceJobStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the location of the table metadata.
// Permissions You must have the s3tables:GetTableMetadataLocation permission to
// use this operation.
func s3tables_GetTableMetadataLocation(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableMetadataLocationInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableMetadataLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a table policy. For more information, see [Viewing a table policy] in the Amazon
// Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:GetTablePolicy permission to use this
// operation.
//
// [Viewing a table policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-table-policy.html#table-policy-get
func s3tables_GetTablePolicy(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTablePolicyInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTablePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the expiration configuration settings for records in a table, and the
// status of the configuration. If the status of the configuration is enabled ,
// records expire and are automatically removed from the table after the specified
// number of days.
//
// Permissions You must have the s3tables:GetTableRecordExpirationConfiguration
// permission to use this operation.
func s3tables_GetTableRecordExpirationConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableRecordExpirationConfigurationInput{
		// TableArn: *string, // Required
	}

	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}

	if resp, err := client.GetTableRecordExpirationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the status, metrics, and details of the latest record expiration job
// for a table. This includes when the job ran, and whether it succeeded or failed.
// If the job ran successfully, this also includes statistics about the records
// that were removed.
//
// Permissions You must have the s3tables:GetTableRecordExpirationJobStatus
// permission to use this operation.
func s3tables_GetTableRecordExpirationJobStatus(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableRecordExpirationJobStatusInput{
		// TableArn: *string, // Required
	}

	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}

	if resp, err := client.GetTableRecordExpirationJobStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the replication configuration for a specific table.
// Permissions You must have the s3tables:GetTableReplication permission to use
// this operation.
func s3tables_GetTableReplication(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableReplicationInput{
		// TableArn: *string, // Required
	}

	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}

	if resp, err := client.GetTableReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the replication status for a table, including the status of
// replication to each destination. This operation provides visibility into
// replication health and progress.
//
// Permissions You must have the s3tables:GetTableReplicationStatus permission to
// use this operation.
func s3tables_GetTableReplicationStatus(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableReplicationStatusInput{
		// TableArn: *string, // Required
	}

	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}

	if resp, err := client.GetTableReplicationStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the storage class configuration for a specific table. This allows you
// to view the storage class settings that apply to an individual table, which may
// differ from the table bucket's default configuration.
//
// Permissions You must have the s3tables:GetTableStorageClass permission to use
// this operation.
func s3tables_GetTableStorageClass(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.GetTableStorageClassInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.GetTableStorageClass(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the namespaces within a table bucket. For more information, see [Table namespaces] in the
// Amazon Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:ListNamespaces permission to use this
// operation.
//
// [Table namespaces]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-namespace.html
func s3tables_ListNamespaces(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.ListNamespacesInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3tablesContinuationToken)
	}
	if len(_s3tablesMaxNamespaces) > 0 {
		if err := assignInputField(input, "MaxNamespaces", _s3tablesMaxNamespaces); err != nil {
			log.Errorf("invalid --max-namespaces: %s", err.Error())
			return
		}
	}
	if len(_s3tablesPrefix) > 0 {
		input.Prefix = aws.String(_s3tablesPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListNamespaces(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3tables.ListNamespacesOutput
	p := s3tables.NewListNamespacesPaginator(client, input)
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

// Lists table buckets for your account. For more information, see [S3 Table buckets] in the Amazon
// Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:ListTableBuckets permission to use this
// operation.
//
// [S3 Table buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-buckets.html
func s3tables_ListTableBuckets(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.ListTableBucketsInput{}

	if len(_s3tablesContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3tablesContinuationToken)
	}
	if len(_s3tablesMaxBuckets) > 0 {
		if err := assignInputField(input, "MaxBuckets", _s3tablesMaxBuckets); err != nil {
			log.Errorf("invalid --max-buckets: %s", err.Error())
			return
		}
	}
	if len(_s3tablesPrefix) > 0 {
		input.Prefix = aws.String(_s3tablesPrefix)
	}
	if len(_s3tablesType) > 0 {
		if err := assignInputField(input, "Type", _s3tablesType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTableBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3tables.ListTableBucketsOutput
	p := s3tables.NewListTableBucketsPaginator(client, input)
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

// List tables in the given table bucket. For more information, see [S3 Tables] in the Amazon
// Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:ListTables permission to use this
// operation.
//
// [S3 Tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-tables.html
func s3tables_ListTables(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.ListTablesInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3tablesContinuationToken)
	}
	if len(_s3tablesMaxTables) > 0 {
		if err := assignInputField(input, "MaxTables", _s3tablesMaxTables); err != nil {
			log.Errorf("invalid --max-tables: %s", err.Error())
			return
		}
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesPrefix) > 0 {
		input.Prefix = aws.String(_s3tablesPrefix)
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

	var results []*s3tables.ListTablesOutput
	p := s3tables.NewListTablesPaginator(client, input)
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

// Lists all of the tags applied to a specified Amazon S3 Tables resource. Each
// tag is a label consisting of a key and value pair. Tags can help you organize,
// track costs for, and control access to resources.
//
// For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources].
//
// Permissions For tables and table buckets, you must have the
// s3tables:ListTagsForResource permission to use this operation.
//
// [Managing tags for Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#manage-tags
func s3tables_ListTagsForResource(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_s3tablesResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3tablesResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the encryption configuration for a table bucket.
// Permissions You must have the s3tables:PutTableBucketEncryption permission to
// use this operation.
//
// If you choose SSE-KMS encryption you must grant the S3 Tables maintenance
// principal access to your KMS key. For more information, see [Permissions requirements for S3 Tables SSE-KMS encryption]in the Amazon
// Simple Storage Service User Guide.
//
// [Permissions requirements for S3 Tables SSE-KMS encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-kms-permissions.html
func s3tables_PutTableBucketEncryption(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableBucketEncryptionInput{
		// EncryptionConfiguration: *types.EncryptionConfiguration, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _s3tablesEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.PutTableBucketEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new maintenance configuration or replaces an existing maintenance
// configuration for a table bucket. For more information, see [Amazon S3 table bucket maintenance]in the Amazon
// Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:PutTableBucketMaintenanceConfiguration
// permission to use this operation.
//
// [Amazon S3 table bucket maintenance]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-table-buckets-maintenance.html
func s3tables_PutTableBucketMaintenanceConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableBucketMaintenanceConfigurationInput{
		// TableBucketARN: *string, // Required
		// Type: types.TableBucketMaintenanceType, // Required
		// Value: *types.TableBucketMaintenanceConfigurationValue, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesType) > 0 {
		if err := assignInputField(input, "Type", _s3tablesType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_s3tablesValue) > 0 {
		if err := assignInputField(input, "Value", _s3tablesValue); err != nil {
			log.Errorf("invalid --value: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutTableBucketMaintenanceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the metrics configuration for a table bucket.
// Permissions You must have the s3tables:PutTableBucketMetricsConfiguration
// permission to use this operation.
func s3tables_PutTableBucketMetricsConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableBucketMetricsConfigurationInput{
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.PutTableBucketMetricsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new table bucket policy or replaces an existing table bucket policy
// for a table bucket. For more information, see [Adding a table bucket policy]in the Amazon Simple Storage
// Service User Guide.
//
// Permissions You must have the s3tables:PutTableBucketPolicy permission to use
// this operation.
//
// [Adding a table bucket policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-bucket-policy.html#table-bucket-policy-add
func s3tables_PutTableBucketPolicy(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableBucketPolicyInput{
		// ResourcePolicy: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_s3tablesResourcePolicy)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.PutTableBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the replication configuration for a table bucket. This
// operation defines how tables in the source bucket are replicated to destination
// buckets. Replication helps ensure data availability and disaster recovery across
// regions or accounts.
//
// # Permissions
//
// - You must have the s3tables:PutTableBucketReplication permission to use this
// operation. The IAM role specified in the configuration must have permissions to
// read from the source bucket and write permissions to all destination buckets.
//
// - You must also have the following permissions:
//
// - s3tables:GetTable permission on the source table.
//
// - s3tables:ListTables permission on the bucket containing the table.
//
// - s3tables:CreateTable permission for the destination.
//
// - s3tables:CreateNamespace permission for the destination.
//
// - s3tables:GetTableMaintenanceConfig permission for the source bucket.
//
// - s3tables:PutTableMaintenanceConfig permission for the destination bucket.
//
// - You must have iam:PassRole permission with condition allowing roles to be
// passed to replication.s3tables.amazonaws.com .
func s3tables_PutTableBucketReplication(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableBucketReplicationInput{
		// Configuration: *types.TableBucketReplicationConfiguration, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _s3tablesConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.PutTableBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets or updates the storage class configuration for a table bucket. This
// configuration serves as the default storage class for all new tables created in
// the bucket, allowing you to optimize storage costs at the bucket level.
//
// Permissions You must have the s3tables:PutTableBucketStorageClass permission to
// use this operation.
func s3tables_PutTableBucketStorageClass(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableBucketStorageClassInput{
		// StorageClassConfiguration: *types.StorageClassConfiguration, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesStorageClassConfiguration) > 0 {
		if err := assignInputField(input, "StorageClassConfiguration", _s3tablesStorageClassConfiguration); err != nil {
			log.Errorf("invalid --storage-class-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.PutTableBucketStorageClass(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new maintenance configuration or replaces an existing maintenance
// configuration for a table. For more information, see [S3 Tables maintenance]in the Amazon Simple
// Storage Service User Guide.
//
// Permissions You must have the s3tables:PutTableMaintenanceConfiguration
// permission to use this operation.
//
// [S3 Tables maintenance]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-maintenance.html
func s3tables_PutTableMaintenanceConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableMaintenanceConfigurationInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
		// Type: types.TableMaintenanceType, // Required
		// Value: *types.TableMaintenanceConfigurationValue, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesType) > 0 {
		if err := assignInputField(input, "Type", _s3tablesType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_s3tablesValue) > 0 {
		if err := assignInputField(input, "Value", _s3tablesValue); err != nil {
			log.Errorf("invalid --value: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutTableMaintenanceConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new table policy or replaces an existing table policy for a table.
// For more information, see [Adding a table policy]in the Amazon Simple Storage Service User Guide.
//
// Permissions You must have the s3tables:PutTablePolicy permission to use this
// operation.
//
// [Adding a table policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-table-policy.html#table-policy-add
func s3tables_PutTablePolicy(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTablePolicyInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// ResourcePolicy: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesResourcePolicy) > 0 {
		input.ResourcePolicy = aws.String(_s3tablesResourcePolicy)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}

	if resp, err := client.PutTablePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the expiration configuration settings for records in a
// table, including the status of the configuration. If you enable record
// expiration for a table, records expire and are automatically removed from the
// table after the number of days that you specify.
//
// Permissions You must have the s3tables:PutTableRecordExpirationConfiguration
// permission to use this operation.
func s3tables_PutTableRecordExpirationConfiguration(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableRecordExpirationConfigurationInput{
		// TableArn: *string, // Required
		// Value: *types.TableRecordExpirationConfigurationValue, // Required
	}

	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}
	if len(_s3tablesValue) > 0 {
		if err := assignInputField(input, "Value", _s3tablesValue); err != nil {
			log.Errorf("invalid --value: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutTableRecordExpirationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the replication configuration for a specific table. This
// operation allows you to define table-level replication independently of
// bucket-level replication, providing granular control over which tables are
// replicated and where.
//
// # Permissions
//
// - You must have the s3tables:PutTableReplication permission to use this
// operation. The IAM role specified in the configuration must have permissions to
// read from the source table and write to all destination tables.
//
// - You must also have the following permissions:
//
// - s3tables:GetTable permission on the source table being replicated.
//
// - s3tables:CreateTable permission for the destination.
//
// - s3tables:CreateNamespace permission for the destination.
//
// - s3tables:GetTableMaintenanceConfig permission for the source table.
//
// - s3tables:PutTableMaintenanceConfig permission for the destination table.
//
// - You must have iam:PassRole permission with condition allowing roles to be
// passed to replication.s3tables.amazonaws.com .
func s3tables_PutTableReplication(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.PutTableReplicationInput{
		// Configuration: *types.TableReplicationConfiguration, // Required
		// TableArn: *string, // Required
	}

	if len(_s3tablesConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _s3tablesConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_s3tablesTableArn) > 0 {
		input.TableArn = aws.String(_s3tablesTableArn)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.PutTableReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renames a table or a namespace. For more information, see [S3 Tables] in the Amazon Simple
// Storage Service User Guide.
//
// Permissions You must have the s3tables:RenameTable permission to use this
// operation.
//
// [S3 Tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-tables.html
func s3tables_RenameTable(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.RenameTableInput{
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
	}

	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesNewName) > 0 {
		input.NewName = aws.String(_s3tablesNewName)
	}
	if len(_s3tablesNewNamespaceName) > 0 {
		input.NewNamespaceName = aws.String(_s3tablesNewNamespaceName)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.RenameTable(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies one or more user-defined tags to an Amazon S3 Tables resource or
// updates existing tags. Each tag is a label consisting of a key and value pair.
// Tags can help you organize, track costs for, and control access to your
// resources. You can add up to 50 tags for each S3 resource.
//
// For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources].
//
// Permissions For tables and table buckets, you must have the s3tables:TagResource
// permission to use this operation.
//
// [Managing tags for Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#manage-tags
func s3tables_TagResource(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_s3tablesResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3tablesResourceArn)
	}
	if len(_s3tablesTags) > 0 {
		if err := assignInputField(input, "Tags", _s3tablesTags); err != nil {
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

// Removes the specified user-defined tags from an Amazon S3 Tables resource. You
// can pass one or more tag keys.
//
// For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources].
//
// Permissions For tables and table buckets, you must have the
// s3tables:UntagResource permission to use this operation.
//
// [Managing tags for Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#manage-tags
func s3tables_UntagResource(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_s3tablesResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3tablesResourceArn)
	}
	if len(_s3tablesTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _s3tablesTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the metadata location for a table. The metadata location of a table
// must be an S3 URI that begins with the table's warehouse location. The metadata
// location for an Apache Iceberg table must end with .metadata.json , or if the
// metadata file is Gzip-compressed, .metadata.json.gz .
//
// Permissions You must have the s3tables:UpdateTableMetadataLocation permission
// to use this operation.
func s3tables_UpdateTableMetadataLocation(cfg aws.Config, client *s3tables.Client) {
	input := &s3tables.UpdateTableMetadataLocationInput{
		// MetadataLocation: *string, // Required
		// Name: *string, // Required
		// Namespace: *string, // Required
		// TableBucketARN: *string, // Required
		// VersionToken: *string, // Required
	}

	if len(_s3tablesMetadataLocation) > 0 {
		input.MetadataLocation = aws.String(_s3tablesMetadataLocation)
	}
	if len(_s3tablesName) > 0 {
		input.Name = aws.String(_s3tablesName)
	}
	if len(_s3tablesNamespace) > 0 {
		input.Namespace = aws.String(_s3tablesNamespace)
	}
	if len(_s3tablesTableBucketARN) > 0 {
		input.TableBucketARN = aws.String(_s3tablesTableBucketARN)
	}
	if len(_s3tablesVersionToken) > 0 {
		input.VersionToken = aws.String(_s3tablesVersionToken)
	}

	if resp, err := client.UpdateTableMetadataLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_s3tablesCmd)
	_s3tablesCmd.Flags().SortFlags = false

	_s3tablesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_s3tablesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_s3tablesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_s3tablesCmd.Flags().StringVarP(&_s3tablesConfiguration, "configuration", "", "", "Configuration")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesContinuationToken, "continuation-token", "", "", "Continuation Token")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesFormat, "format", "", "", "Format")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesMaxBuckets, "max-buckets", "", "", "Max Buckets")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesMaxNamespaces, "max-namespaces", "", "", "Max Namespaces")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesMaxTables, "max-tables", "", "", "Max Tables")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesMetadata, "metadata", "", "", "Metadata")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesMetadataLocation, "metadata-location", "", "", "Metadata Location")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesName, "name", "", "", "Name")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesNamespace, "namespace", "", "", "Namespace")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesNewName, "new-name", "", "", "New Name")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesNewNamespaceName, "new-namespace-name", "", "", "New Namespace Name")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesPrefix, "prefix", "", "", "Prefix")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesResourceArn, "resource-arn", "", "", "Resource ARN")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesResourcePolicy, "resource-policy", "", "", "Resource Policy")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesStorageClassConfiguration, "storage-class-configuration", "", "", "Storage Class Configuration")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesTableArn, "table-arn", "", "", "Table ARN")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesTableBucketARN, "table-bucket-arn", "", "", "Table Bucket ARN")
	_s3tablesCmd.Flags().StringSliceVarP(&_s3tablesTagKeys, "tag-keys", "", nil, "Tag Keys")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesTags, "tags", "", "", "Tags")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesType, "type", "", "", "Type")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesValue, "value", "", "", "Value")
	_s3tablesCmd.Flags().StringVarP(&_s3tablesVersionToken, "version-token", "", "", "Version Token")

	_s3tablesCmd.Flags().BoolVarP(&_s3tablesCreateNamespace, "create-namespace", "", false, "Create Namespace")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesCreateTable, "create-table", "", false, "Create Table")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesCreateTableBucket, "create-table-bucket", "", false, "Create Table Bucket")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteNamespace, "delete-namespace", "", false, "Delete Namespace")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTable, "delete-table", "", false, "Delete Table")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTableBucket, "delete-table-bucket", "", false, "Delete Table Bucket")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTableBucketEncryption, "delete-table-bucket-encryption", "", false, "Delete Table Bucket Encryption")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTableBucketMetricsConfiguration, "delete-table-bucket-metrics-configuration", "", false, "Delete Table Bucket Metrics Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTableBucketPolicy, "delete-table-bucket-policy", "", false, "Delete Table Bucket Policy")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTableBucketReplication, "delete-table-bucket-replication", "", false, "Delete Table Bucket Replication")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTablePolicy, "delete-table-policy", "", false, "Delete Table Policy")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesDeleteTableReplication, "delete-table-replication", "", false, "Delete Table Replication")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetNamespace, "get-namespace", "", false, "Get Namespace")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTable, "get-table", "", false, "Get Table")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucket, "get-table-bucket", "", false, "Get Table Bucket")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucketEncryption, "get-table-bucket-encryption", "", false, "Get Table Bucket Encryption")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucketMaintenanceConfiguration, "get-table-bucket-maintenance-configuration", "", false, "Get Table Bucket Maintenance Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucketMetricsConfiguration, "get-table-bucket-metrics-configuration", "", false, "Get Table Bucket Metrics Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucketPolicy, "get-table-bucket-policy", "", false, "Get Table Bucket Policy")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucketReplication, "get-table-bucket-replication", "", false, "Get Table Bucket Replication")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableBucketStorageClass, "get-table-bucket-storage-class", "", false, "Get Table Bucket Storage Class")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableEncryption, "get-table-encryption", "", false, "Get Table Encryption")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableMaintenanceConfiguration, "get-table-maintenance-configuration", "", false, "Get Table Maintenance Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableMaintenanceJobStatus, "get-table-maintenance-job-status", "", false, "Get Table Maintenance Job Status")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableMetadataLocation, "get-table-metadata-location", "", false, "Get Table Metadata Location")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTablePolicy, "get-table-policy", "", false, "Get Table Policy")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableRecordExpirationConfiguration, "get-table-record-expiration-configuration", "", false, "Get Table Record Expiration Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableRecordExpirationJobStatus, "get-table-record-expiration-job-status", "", false, "Get Table Record Expiration Job Status")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableReplication, "get-table-replication", "", false, "Get Table Replication")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableReplicationStatus, "get-table-replication-status", "", false, "Get Table Replication Status")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesGetTableStorageClass, "get-table-storage-class", "", false, "Get Table Storage Class")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesListNamespaces, "list-namespaces", "", false, "List Namespaces")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesListTableBuckets, "list-table-buckets", "", false, "List Table Buckets")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesListTables, "list-tables", "", false, "List Tables")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableBucketEncryption, "put-table-bucket-encryption", "", false, "Put Table Bucket Encryption")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableBucketMaintenanceConfiguration, "put-table-bucket-maintenance-configuration", "", false, "Put Table Bucket Maintenance Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableBucketMetricsConfiguration, "put-table-bucket-metrics-configuration", "", false, "Put Table Bucket Metrics Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableBucketPolicy, "put-table-bucket-policy", "", false, "Put Table Bucket Policy")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableBucketReplication, "put-table-bucket-replication", "", false, "Put Table Bucket Replication")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableBucketStorageClass, "put-table-bucket-storage-class", "", false, "Put Table Bucket Storage Class")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableMaintenanceConfiguration, "put-table-maintenance-configuration", "", false, "Put Table Maintenance Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTablePolicy, "put-table-policy", "", false, "Put Table Policy")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableRecordExpirationConfiguration, "put-table-record-expiration-configuration", "", false, "Put Table Record Expiration Configuration")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesPutTableReplication, "put-table-replication", "", false, "Put Table Replication")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesRenameTable, "rename-table", "", false, "Rename Table")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesTagResource, "tag-resource", "", false, "Tag Resource")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesUntagResource, "untag-resource", "", false, "Untag Resource")
	_s3tablesCmd.Flags().BoolVarP(&_s3tablesUpdateTableMetadataLocation, "update-table-metadata-location", "", false, "Update Table Metadata Location")

}
