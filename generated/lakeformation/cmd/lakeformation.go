package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// lakeformationCmd represents the lakeformation command
var _lakeformationCmd = &cobra.Command{
	Use:   "lakeformation",
	Short: "AWS lakeformation CLI",
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
		client := lakeformation.NewFromConfig(cfg)
		if _lakeformationAddLFTagsToResource {
			lakeformation_AddLFTagsToResource(cfg, client)
			return
		}
		if _lakeformationAssumeDecoratedRoleWithSAML {
			lakeformation_AssumeDecoratedRoleWithSAML(cfg, client)
			return
		}
		if _lakeformationBatchGrantPermissions {
			lakeformation_BatchGrantPermissions(cfg, client)
			return
		}
		if _lakeformationBatchRevokePermissions {
			lakeformation_BatchRevokePermissions(cfg, client)
			return
		}
		if _lakeformationCancelTransaction {
			lakeformation_CancelTransaction(cfg, client)
			return
		}
		if _lakeformationCommitTransaction {
			lakeformation_CommitTransaction(cfg, client)
			return
		}
		if _lakeformationCreateDataCellsFilter {
			lakeformation_CreateDataCellsFilter(cfg, client)
			return
		}
		if _lakeformationCreateLakeFormationIdentityCenterConfiguration {
			lakeformation_CreateLakeFormationIdentityCenterConfiguration(cfg, client)
			return
		}
		if _lakeformationCreateLakeFormationOptIn {
			lakeformation_CreateLakeFormationOptIn(cfg, client)
			return
		}
		if _lakeformationCreateLFTag {
			lakeformation_CreateLFTag(cfg, client)
			return
		}
		if _lakeformationCreateLFTagExpression {
			lakeformation_CreateLFTagExpression(cfg, client)
			return
		}
		if _lakeformationDeleteDataCellsFilter {
			lakeformation_DeleteDataCellsFilter(cfg, client)
			return
		}
		if _lakeformationDeleteLakeFormationIdentityCenterConfiguration {
			lakeformation_DeleteLakeFormationIdentityCenterConfiguration(cfg, client)
			return
		}
		if _lakeformationDeleteLakeFormationOptIn {
			lakeformation_DeleteLakeFormationOptIn(cfg, client)
			return
		}
		if _lakeformationDeleteLFTag {
			lakeformation_DeleteLFTag(cfg, client)
			return
		}
		if _lakeformationDeleteLFTagExpression {
			lakeformation_DeleteLFTagExpression(cfg, client)
			return
		}
		if _lakeformationDeleteObjectsOnCancel {
			lakeformation_DeleteObjectsOnCancel(cfg, client)
			return
		}
		if _lakeformationDeregisterResource {
			lakeformation_DeregisterResource(cfg, client)
			return
		}
		if _lakeformationDescribeLakeFormationIdentityCenterConfiguration {
			lakeformation_DescribeLakeFormationIdentityCenterConfiguration(cfg, client)
			return
		}
		if _lakeformationDescribeResource {
			lakeformation_DescribeResource(cfg, client)
			return
		}
		if _lakeformationDescribeTransaction {
			lakeformation_DescribeTransaction(cfg, client)
			return
		}
		if _lakeformationExtendTransaction {
			lakeformation_ExtendTransaction(cfg, client)
			return
		}
		if _lakeformationGetDataCellsFilter {
			lakeformation_GetDataCellsFilter(cfg, client)
			return
		}
		if _lakeformationGetDataLakePrincipal {
			lakeformation_GetDataLakePrincipal(cfg, client)
			return
		}
		if _lakeformationGetDataLakeSettings {
			lakeformation_GetDataLakeSettings(cfg, client)
			return
		}
		if _lakeformationGetEffectivePermissionsForPath {
			lakeformation_GetEffectivePermissionsForPath(cfg, client)
			return
		}
		if _lakeformationGetLFTag {
			lakeformation_GetLFTag(cfg, client)
			return
		}
		if _lakeformationGetLFTagExpression {
			lakeformation_GetLFTagExpression(cfg, client)
			return
		}
		if _lakeformationGetQueryState {
			lakeformation_GetQueryState(cfg, client)
			return
		}
		if _lakeformationGetQueryStatistics {
			lakeformation_GetQueryStatistics(cfg, client)
			return
		}
		if _lakeformationGetResourceLFTags {
			lakeformation_GetResourceLFTags(cfg, client)
			return
		}
		if _lakeformationGetTableObjects {
			lakeformation_GetTableObjects(cfg, client)
			return
		}
		if _lakeformationGetTemporaryDataLocationCredentials {
			lakeformation_GetTemporaryDataLocationCredentials(cfg, client)
			return
		}
		if _lakeformationGetTemporaryGluePartitionCredentials {
			lakeformation_GetTemporaryGluePartitionCredentials(cfg, client)
			return
		}
		if _lakeformationGetTemporaryGlueTableCredentials {
			lakeformation_GetTemporaryGlueTableCredentials(cfg, client)
			return
		}
		if _lakeformationGetWorkUnitResults {
			lakeformation_GetWorkUnitResults(cfg, client)
			return
		}
		if _lakeformationGetWorkUnits {
			lakeformation_GetWorkUnits(cfg, client)
			return
		}
		if _lakeformationGrantPermissions {
			lakeformation_GrantPermissions(cfg, client)
			return
		}
		if _lakeformationListDataCellsFilter {
			lakeformation_ListDataCellsFilter(cfg, client)
			return
		}
		if _lakeformationListLakeFormationOptIns {
			lakeformation_ListLakeFormationOptIns(cfg, client)
			return
		}
		if _lakeformationListLFTagExpressions {
			lakeformation_ListLFTagExpressions(cfg, client)
			return
		}
		if _lakeformationListLFTags {
			lakeformation_ListLFTags(cfg, client)
			return
		}
		if _lakeformationListPermissions {
			lakeformation_ListPermissions(cfg, client)
			return
		}
		if _lakeformationListResources {
			lakeformation_ListResources(cfg, client)
			return
		}
		if _lakeformationListTableStorageOptimizers {
			lakeformation_ListTableStorageOptimizers(cfg, client)
			return
		}
		if _lakeformationListTransactions {
			lakeformation_ListTransactions(cfg, client)
			return
		}
		if _lakeformationPutDataLakeSettings {
			lakeformation_PutDataLakeSettings(cfg, client)
			return
		}
		if _lakeformationRegisterResource {
			lakeformation_RegisterResource(cfg, client)
			return
		}
		if _lakeformationRemoveLFTagsFromResource {
			lakeformation_RemoveLFTagsFromResource(cfg, client)
			return
		}
		if _lakeformationRevokePermissions {
			lakeformation_RevokePermissions(cfg, client)
			return
		}
		if _lakeformationSearchDatabasesByLFTags {
			lakeformation_SearchDatabasesByLFTags(cfg, client)
			return
		}
		if _lakeformationSearchTablesByLFTags {
			lakeformation_SearchTablesByLFTags(cfg, client)
			return
		}
		if _lakeformationStartQueryPlanning {
			lakeformation_StartQueryPlanning(cfg, client)
			return
		}
		if _lakeformationStartTransaction {
			lakeformation_StartTransaction(cfg, client)
			return
		}
		if _lakeformationUpdateDataCellsFilter {
			lakeformation_UpdateDataCellsFilter(cfg, client)
			return
		}
		if _lakeformationUpdateLakeFormationIdentityCenterConfiguration {
			lakeformation_UpdateLakeFormationIdentityCenterConfiguration(cfg, client)
			return
		}
		if _lakeformationUpdateLFTag {
			lakeformation_UpdateLFTag(cfg, client)
			return
		}
		if _lakeformationUpdateLFTagExpression {
			lakeformation_UpdateLFTagExpression(cfg, client)
			return
		}
		if _lakeformationUpdateResource {
			lakeformation_UpdateResource(cfg, client)
			return
		}
		if _lakeformationUpdateTableObjects {
			lakeformation_UpdateTableObjects(cfg, client)
			return
		}
		if _lakeformationUpdateTableStorageOptimizer {
			lakeformation_UpdateTableStorageOptimizer(cfg, client)
			return
		}

	},
}

var (
	_lakeformationAddLFTagsToResource                              bool
	_lakeformationAssumeDecoratedRoleWithSAML                      bool
	_lakeformationBatchGrantPermissions                            bool
	_lakeformationBatchRevokePermissions                           bool
	_lakeformationCancelTransaction                                bool
	_lakeformationCommitTransaction                                bool
	_lakeformationCreateDataCellsFilter                            bool
	_lakeformationCreateLakeFormationIdentityCenterConfiguration   bool
	_lakeformationCreateLakeFormationOptIn                         bool
	_lakeformationCreateLFTag                                      bool
	_lakeformationCreateLFTagExpression                            bool
	_lakeformationDeleteDataCellsFilter                            bool
	_lakeformationDeleteLakeFormationIdentityCenterConfiguration   bool
	_lakeformationDeleteLakeFormationOptIn                         bool
	_lakeformationDeleteLFTag                                      bool
	_lakeformationDeleteLFTagExpression                            bool
	_lakeformationDeleteObjectsOnCancel                            bool
	_lakeformationDeregisterResource                               bool
	_lakeformationDescribeLakeFormationIdentityCenterConfiguration bool
	_lakeformationDescribeResource                                 bool
	_lakeformationDescribeTransaction                              bool
	_lakeformationExtendTransaction                                bool
	_lakeformationGetDataCellsFilter                               bool
	_lakeformationGetDataLakePrincipal                             bool
	_lakeformationGetDataLakeSettings                              bool
	_lakeformationGetEffectivePermissionsForPath                   bool
	_lakeformationGetLFTag                                         bool
	_lakeformationGetLFTagExpression                               bool
	_lakeformationGetQueryState                                    bool
	_lakeformationGetQueryStatistics                               bool
	_lakeformationGetResourceLFTags                                bool
	_lakeformationGetTableObjects                                  bool
	_lakeformationGetTemporaryDataLocationCredentials              bool
	_lakeformationGetTemporaryGluePartitionCredentials             bool
	_lakeformationGetTemporaryGlueTableCredentials                 bool
	_lakeformationGetWorkUnitResults                               bool
	_lakeformationGetWorkUnits                                     bool
	_lakeformationGrantPermissions                                 bool
	_lakeformationListDataCellsFilter                              bool
	_lakeformationListLakeFormationOptIns                          bool
	_lakeformationListLFTagExpressions                             bool
	_lakeformationListLFTags                                       bool
	_lakeformationListPermissions                                  bool
	_lakeformationListResources                                    bool
	_lakeformationListTableStorageOptimizers                       bool
	_lakeformationListTransactions                                 bool
	_lakeformationPutDataLakeSettings                              bool
	_lakeformationRegisterResource                                 bool
	_lakeformationRemoveLFTagsFromResource                         bool
	_lakeformationRevokePermissions                                bool
	_lakeformationSearchDatabasesByLFTags                          bool
	_lakeformationSearchTablesByLFTags                             bool
	_lakeformationStartQueryPlanning                               bool
	_lakeformationStartTransaction                                 bool
	_lakeformationUpdateDataCellsFilter                            bool
	_lakeformationUpdateLakeFormationIdentityCenterConfiguration   bool
	_lakeformationUpdateLFTag                                      bool
	_lakeformationUpdateLFTagExpression                            bool
	_lakeformationUpdateResource                                   bool
	_lakeformationUpdateTableObjects                               bool
	_lakeformationUpdateTableStorageOptimizer                      bool

	_lakeformationApplicationStatus            string
	_lakeformationAuditContext                 string
	_lakeformationCatalogId                    string
	_lakeformationCondition                    string
	_lakeformationCredentialsScope             string
	_lakeformationDataLakeSettings             string
	_lakeformationDataLocations                []string
	_lakeformationDatabaseName                 string
	_lakeformationDescription                  string
	_lakeformationDurationSeconds              string
	_lakeformationEntries                      string
	_lakeformationExpectedResourceOwnerAccount string
	_lakeformationExpression                   string
	_lakeformationExternalFiltering            string
	_lakeformationFilterConditionList          string
	_lakeformationHybridAccessEnabled          string
	_lakeformationIncludeRelated               string
	_lakeformationInstanceArn                  string
	_lakeformationLFTags                       string
	_lakeformationMaxResults                   string
	_lakeformationName                         string
	_lakeformationNextToken                    string
	_lakeformationObjects                      string
	_lakeformationPageSize                     string
	_lakeformationPartition                    string
	_lakeformationPartitionPredicate           string
	_lakeformationPermissions                  string
	_lakeformationPermissionsWithGrantOption   string
	_lakeformationPrincipal                    string
	_lakeformationPrincipalArn                 string
	_lakeformationQueryAsOfTime                string
	_lakeformationQueryId                      string
	_lakeformationQueryPlanningContext         string
	_lakeformationQuerySessionContext          string
	_lakeformationQueryString                  string
	_lakeformationResource                     string
	_lakeformationResourceArn                  string
	_lakeformationResourceShareType            string
	_lakeformationResourceType                 string
	_lakeformationRoleArn                      string
	_lakeformationS3Path                       string
	_lakeformationSAMLAssertion                string
	_lakeformationServiceIntegrations          string
	_lakeformationShareRecipients              string
	_lakeformationShowAssignedLFTags           string
	_lakeformationStatusFilter                 string
	_lakeformationStorageOptimizerConfig       string
	_lakeformationStorageOptimizerType         string
	_lakeformationSupportedPermissionTypes     string
	_lakeformationTable                        string
	_lakeformationTableArn                     string
	_lakeformationTableCatalogId               string
	_lakeformationTableData                    string
	_lakeformationTableName                    string
	_lakeformationTagKey                       string
	_lakeformationTagValues                    []string
	_lakeformationTagValuesToAdd               []string
	_lakeformationTagValuesToDelete            []string
	_lakeformationTransactionId                string
	_lakeformationTransactionType              string
	_lakeformationUseServiceLinkedRole         string
	_lakeformationWithFederation               string
	_lakeformationWithPrivilegedAccess         string
	_lakeformationWorkUnitId                   string
	_lakeformationWorkUnitToken                string
	_lakeformationWriteOperations              string
)

// Attaches one or more LF-tags to an existing resource.
func lakeformation_AddLFTagsToResource(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.AddLFTagsToResourceInput{
		// LFTags: []types.LFTagPair, // Required
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationLFTags) > 0 {
		if err := assignInputField(input, "LFTags", _lakeformationLFTags); err != nil {
			log.Errorf("invalid --lf-tags: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.AddLFTagsToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a caller to assume an IAM role decorated as the SAML user specified in
// the SAML assertion included in the request. This decoration allows Lake
// Formation to enforce access policies against the SAML users and groups. This API
// operation requires SAML federation setup in the caller’s account as it can only
// be called with valid SAML assertions. Lake Formation does not scope down the
// permission of the assumed role. All permissions attached to the role via the
// SAML federation setup will be included in the role session.
//
// This decorated role is expected to access data in Amazon S3 by getting
// temporary access from Lake Formation which is authorized via the virtual API
// GetDataAccess . Therefore, all SAML roles that can be assumed via
// AssumeDecoratedRoleWithSAML must at a minimum include
// lakeformation:GetDataAccess in their role policies. A typical IAM policy
// attached to such a role would include the following actions:
//
// - glue:*Database*
//
// - glue:*Table*
//
// - glue:*Partition*
//
// - glue:*UserDefinedFunction*
//
// - lakeformation:GetDataAccess
func lakeformation_AssumeDecoratedRoleWithSAML(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.AssumeDecoratedRoleWithSAMLInput{
		// PrincipalArn: *string, // Required
		// RoleArn: *string, // Required
		// SAMLAssertion: *string, // Required
	}

	if len(_lakeformationPrincipalArn) > 0 {
		input.PrincipalArn = aws.String(_lakeformationPrincipalArn)
	}
	if len(_lakeformationRoleArn) > 0 {
		input.RoleArn = aws.String(_lakeformationRoleArn)
	}
	if len(_lakeformationSAMLAssertion) > 0 {
		input.SAMLAssertion = aws.String(_lakeformationSAMLAssertion)
	}
	if len(_lakeformationDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _lakeformationDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssumeDecoratedRoleWithSAML(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Batch operation to grant permissions to the principal.
func lakeformation_BatchGrantPermissions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.BatchGrantPermissionsInput{
		// Entries: []types.BatchPermissionsRequestEntry, // Required
	}

	if len(_lakeformationEntries) > 0 {
		if err := assignInputField(input, "Entries", _lakeformationEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.BatchGrantPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Batch operation to revoke permissions from the principal.
func lakeformation_BatchRevokePermissions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.BatchRevokePermissionsInput{
		// Entries: []types.BatchPermissionsRequestEntry, // Required
	}

	if len(_lakeformationEntries) > 0 {
		if err := assignInputField(input, "Entries", _lakeformationEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.BatchRevokePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to cancel the specified transaction. Returns an exception if the
// transaction was previously committed.
func lakeformation_CancelTransaction(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CancelTransactionInput{
		// TransactionId: *string, // Required
	}

	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}

	if resp, err := client.CancelTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attempts to commit the specified transaction. Returns an exception if the
// transaction was previously aborted. This API action is idempotent if called
// multiple times for the same transaction.
func lakeformation_CommitTransaction(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CommitTransactionInput{
		// TransactionId: *string, // Required
	}

	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}

	if resp, err := client.CommitTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data cell filter to allow one to grant access to certain columns on
// certain rows.
func lakeformation_CreateDataCellsFilter(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CreateDataCellsFilterInput{
		// TableData: *types.DataCellsFilter, // Required
	}

	if len(_lakeformationTableData) > 0 {
		if err := assignInputField(input, "TableData", _lakeformationTableData); err != nil {
			log.Errorf("invalid --table-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDataCellsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an IAM Identity Center connection with Lake Formation to allow IAM
// Identity Center users and groups to access Data Catalog resources.
func lakeformation_CreateLakeFormationIdentityCenterConfiguration(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CreateLakeFormationIdentityCenterConfigurationInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationExternalFiltering) > 0 {
		if err := assignInputField(input, "ExternalFiltering", _lakeformationExternalFiltering); err != nil {
			log.Errorf("invalid --external-filtering: %s", err.Error())
			return
		}
	}
	if len(_lakeformationInstanceArn) > 0 {
		input.InstanceArn = aws.String(_lakeformationInstanceArn)
	}
	if len(_lakeformationServiceIntegrations) > 0 {
		if err := assignInputField(input, "ServiceIntegrations", _lakeformationServiceIntegrations); err != nil {
			log.Errorf("invalid --service-integrations: %s", err.Error())
			return
		}
	}
	if len(_lakeformationShareRecipients) > 0 {
		if err := assignInputField(input, "ShareRecipients", _lakeformationShareRecipients); err != nil {
			log.Errorf("invalid --share-recipients: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLakeFormationIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enforce Lake Formation permissions for the given databases, tables, and
// principals.
func lakeformation_CreateLakeFormationOptIn(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CreateLakeFormationOptInInput{
		// Principal: *types.DataLakePrincipal, // Required
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _lakeformationPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCondition) > 0 {
		if err := assignInputField(input, "Condition", _lakeformationCondition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLakeFormationOptIn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an LF-tag with the specified name and values.
func lakeformation_CreateLFTag(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CreateLFTagInput{
		// TagKey: *string, // Required
		// TagValues: []string, // Required
	}

	if len(_lakeformationTagKey) > 0 {
		input.TagKey = aws.String(_lakeformationTagKey)
	}
	if len(_lakeformationTagValues) > 0 {
		input.TagValues = append([]string(nil), _lakeformationTagValues...)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.CreateLFTag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new LF-Tag expression with the provided name, description, catalog
// ID, and expression body. This call fails if a LF-Tag expression with the same
// name already exists in the caller’s account or if the underlying LF-Tags don't
// exist. To call this API operation, caller needs the following Lake Formation
// permissions:
//
// CREATE_LF_TAG_EXPRESSION on the root catalog resource.
//
// GRANT_WITH_LF_TAG_EXPRESSION on all underlying LF-Tag key:value pairs included
// in the expression.
func lakeformation_CreateLFTagExpression(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.CreateLFTagExpressionInput{
		// Expression: []types.LFTag, // Required
		// Name: *string, // Required
	}

	if len(_lakeformationExpression) > 0 {
		if err := assignInputField(input, "Expression", _lakeformationExpression); err != nil {
			log.Errorf("invalid --expression: %s", err.Error())
			return
		}
	}
	if len(_lakeformationName) > 0 {
		input.Name = aws.String(_lakeformationName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationDescription) > 0 {
		input.Description = aws.String(_lakeformationDescription)
	}

	if resp, err := client.CreateLFTagExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a data cell filter.
func lakeformation_DeleteDataCellsFilter(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeleteDataCellsFilterInput{}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationName) > 0 {
		input.Name = aws.String(_lakeformationName)
	}
	if len(_lakeformationTableCatalogId) > 0 {
		input.TableCatalogId = aws.String(_lakeformationTableCatalogId)
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}

	if resp, err := client.DeleteDataCellsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an IAM Identity Center connection with Lake Formation.
func lakeformation_DeleteLakeFormationIdentityCenterConfiguration(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeleteLakeFormationIdentityCenterConfigurationInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.DeleteLakeFormationIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Remove the Lake Formation permissions enforcement of the given databases,
// tables, and principals.
func lakeformation_DeleteLakeFormationOptIn(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeleteLakeFormationOptInInput{
		// Principal: *types.DataLakePrincipal, // Required
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _lakeformationPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCondition) > 0 {
		if err := assignInputField(input, "Condition", _lakeformationCondition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteLakeFormationOptIn(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an LF-tag by its key name. The operation fails if the specified tag
// key doesn't exist. When you delete an LF-Tag:
//
// - The associated LF-Tag policy becomes invalid.
//
// - Resources that had this tag assigned will no longer have the tag policy
// applied to them.
func lakeformation_DeleteLFTag(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeleteLFTagInput{
		// TagKey: *string, // Required
	}

	if len(_lakeformationTagKey) > 0 {
		input.TagKey = aws.String(_lakeformationTagKey)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.DeleteLFTag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the LF-Tag expression. The caller must be a data lake admin or have DROP
// permissions on the LF-Tag expression. Deleting a LF-Tag expression will also
// delete all LFTagPolicy permissions referencing the LF-Tag expression.
func lakeformation_DeleteLFTagExpression(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeleteLFTagExpressionInput{
		// Name: *string, // Required
	}

	if len(_lakeformationName) > 0 {
		input.Name = aws.String(_lakeformationName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.DeleteLFTagExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For a specific governed table, provides a list of Amazon S3 objects that will
// be written during the current transaction and that can be automatically deleted
// if the transaction is canceled. Without this call, no Amazon S3 objects are
// automatically deleted when a transaction cancels.
//
// The Glue ETL library function write_dynamic_frame.from_catalog() includes an
// option to automatically call DeleteObjectsOnCancel before writes. For more
// information, see [Rolling Back Amazon S3 Writes].
//
// [Rolling Back Amazon S3 Writes]: https://docs.aws.amazon.com/lake-formation/latest/dg/transactions-data-operations.html#rolling-back-writes
func lakeformation_DeleteObjectsOnCancel(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeleteObjectsOnCancelInput{
		// DatabaseName: *string, // Required
		// Objects: []types.VirtualObject, // Required
		// TableName: *string, // Required
		// TransactionId: *string, // Required
	}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationObjects) > 0 {
		if err := assignInputField(input, "Objects", _lakeformationObjects); err != nil {
			log.Errorf("invalid --objects: %s", err.Error())
			return
		}
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}
	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.DeleteObjectsOnCancel(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters the resource as managed by the Data Catalog.
// When you deregister a path, Lake Formation removes the path from the inline
// policy attached to your service-linked role.
func lakeformation_DeregisterResource(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DeregisterResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_lakeformationResourceArn) > 0 {
		input.ResourceArn = aws.String(_lakeformationResourceArn)
	}

	if resp, err := client.DeregisterResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the instance ARN and application ARN for the connection.
func lakeformation_DescribeLakeFormationIdentityCenterConfiguration(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DescribeLakeFormationIdentityCenterConfigurationInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.DescribeLakeFormationIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current data access role for the given resource registered in
// Lake Formation.
func lakeformation_DescribeResource(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DescribeResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_lakeformationResourceArn) > 0 {
		input.ResourceArn = aws.String(_lakeformationResourceArn)
	}

	if resp, err := client.DescribeResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details of a single transaction.
func lakeformation_DescribeTransaction(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.DescribeTransactionInput{
		// TransactionId: *string, // Required
	}

	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}

	if resp, err := client.DescribeTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Indicates to the service that the specified transaction is still active and
// should not be treated as idle and aborted.
//
// Write transactions that remain idle for a long period are automatically aborted
// unless explicitly extended.
func lakeformation_ExtendTransaction(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ExtendTransactionInput{}

	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}

	if resp, err := client.ExtendTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a data cells filter.
func lakeformation_GetDataCellsFilter(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetDataCellsFilterInput{
		// DatabaseName: *string, // Required
		// Name: *string, // Required
		// TableCatalogId: *string, // Required
		// TableName: *string, // Required
	}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationName) > 0 {
		input.Name = aws.String(_lakeformationName)
	}
	if len(_lakeformationTableCatalogId) > 0 {
		input.TableCatalogId = aws.String(_lakeformationTableCatalogId)
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}

	if resp, err := client.GetDataCellsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the identity of the invoking principal.
func lakeformation_GetDataLakePrincipal(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetDataLakePrincipalInput{}

	if resp, err := client.GetDataLakePrincipal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of the data lake administrators of a Lake Formation-managed
// data lake.
func lakeformation_GetDataLakeSettings(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetDataLakeSettingsInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.GetDataLakeSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the Lake Formation permissions for a specified table or database
// resource located at a path in Amazon S3. GetEffectivePermissionsForPath will
// not return databases and tables if the catalog is encrypted.
func lakeformation_GetEffectivePermissionsForPath(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetEffectivePermissionsForPathInput{
		// ResourceArn: *string, // Required
	}

	if len(_lakeformationResourceArn) > 0 {
		input.ResourceArn = aws.String(_lakeformationResourceArn)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetEffectivePermissionsForPath(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.GetEffectivePermissionsForPathOutput
	p := lakeformation.NewGetEffectivePermissionsForPathPaginator(client, input)
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

// Returns an LF-tag definition.
func lakeformation_GetLFTag(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetLFTagInput{
		// TagKey: *string, // Required
	}

	if len(_lakeformationTagKey) > 0 {
		input.TagKey = aws.String(_lakeformationTagKey)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.GetLFTag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the details about the LF-Tag expression. The caller must be a data lake
// admin or must have DESCRIBE permission on the LF-Tag expression resource.
func lakeformation_GetLFTagExpression(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetLFTagExpressionInput{
		// Name: *string, // Required
	}

	if len(_lakeformationName) > 0 {
		input.Name = aws.String(_lakeformationName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.GetLFTagExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the state of a query previously submitted. Clients are expected to poll
// GetQueryState to monitor the current state of the planning before retrieving the
// work units. A query state is only visible to the principal that made the initial
// call to StartQueryPlanning .
func lakeformation_GetQueryState(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetQueryStateInput{
		// QueryId: *string, // Required
	}

	if len(_lakeformationQueryId) > 0 {
		input.QueryId = aws.String(_lakeformationQueryId)
	}

	if resp, err := client.GetQueryState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves statistics on the planning and execution of a query.
func lakeformation_GetQueryStatistics(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetQueryStatisticsInput{
		// QueryId: *string, // Required
	}

	if len(_lakeformationQueryId) > 0 {
		input.QueryId = aws.String(_lakeformationQueryId)
	}

	if resp, err := client.GetQueryStatistics(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the LF-tags applied to a resource.
func lakeformation_GetResourceLFTags(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetResourceLFTagsInput{
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationShowAssignedLFTags) > 0 {
		if err := assignInputField(input, "ShowAssignedLFTags", _lakeformationShowAssignedLFTags); err != nil {
			log.Errorf("invalid --show-assigned-lf-tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetResourceLFTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the set of Amazon S3 objects that make up the specified governed table.
// A transaction ID or timestamp can be specified for time-travel queries.
func lakeformation_GetTableObjects(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetTableObjectsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationPartitionPredicate) > 0 {
		input.PartitionPredicate = aws.String(_lakeformationPartitionPredicate)
	}
	if len(_lakeformationQueryAsOfTime) > 0 {
		if err := assignInputField(input, "QueryAsOfTime", _lakeformationQueryAsOfTime); err != nil {
			log.Errorf("invalid --query-as-of-time: %s", err.Error())
			return
		}
	}
	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}

	if disablePaginator() {
		if resp, err := client.GetTableObjects(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.GetTableObjectsOutput
	p := lakeformation.NewGetTableObjectsPaginator(client, input)
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

// Allows a user or application in a secure environment to access data in a
// specific Amazon S3 location registered with Lake Formation by providing
// temporary scoped credentials that are limited to the requested data location and
// the caller's authorized access level.
//
// The API operation returns an error in the following scenarios:
//
// - The data location is not registered with Lake Formation.
//
// - No Glue table is associated with the data location.
//
// - The caller doesn't have required permissions on the associated table. The
// caller must have SELECT or SUPER permissions on the associated table, and
// credential vending for full table access must be enabled in the data lake
// settings.
//
// For more information, see [Application integration for full table access].
//
// - The data location is in a different Amazon Web Services Region. Lake
// Formation doesn't support cross-Region access when vending credentials for a
// data location. Lake Formation only supports Amazon S3 paths registered within
// the same Region as the API call.
//
// [Application integration for full table access]: https://docs.aws.amazon.com/lake-formation/latest/dg/full-table-credential-vending.html
func lakeformation_GetTemporaryDataLocationCredentials(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetTemporaryDataLocationCredentialsInput{}

	if len(_lakeformationAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _lakeformationAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCredentialsScope) > 0 {
		if err := assignInputField(input, "CredentialsScope", _lakeformationCredentialsScope); err != nil {
			log.Errorf("invalid --credentials-scope: %s", err.Error())
			return
		}
	}
	if len(_lakeformationDataLocations) > 0 {
		input.DataLocations = append([]string(nil), _lakeformationDataLocations...)
	}
	if len(_lakeformationDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _lakeformationDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTemporaryDataLocationCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API is identical to GetTemporaryTableCredentials except that this is used
// when the target Data Catalog resource is of type Partition. Lake Formation
// restricts the permission of the vended credentials with the same scope down
// policy which restricts access to a single Amazon S3 prefix.
func lakeformation_GetTemporaryGluePartitionCredentials(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetTemporaryGluePartitionCredentialsInput{
		// Partition: *types.PartitionValueList, // Required
		// TableArn: *string, // Required
	}

	if len(_lakeformationPartition) > 0 {
		if err := assignInputField(input, "Partition", _lakeformationPartition); err != nil {
			log.Errorf("invalid --partition: %s", err.Error())
			return
		}
	}
	if len(_lakeformationTableArn) > 0 {
		input.TableArn = aws.String(_lakeformationTableArn)
	}
	if len(_lakeformationAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _lakeformationAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_lakeformationDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _lakeformationDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_lakeformationPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _lakeformationPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_lakeformationSupportedPermissionTypes) > 0 {
		if err := assignInputField(input, "SupportedPermissionTypes", _lakeformationSupportedPermissionTypes); err != nil {
			log.Errorf("invalid --supported-permission-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTemporaryGluePartitionCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows a caller in a secure environment to assume a role with permission to
// access Amazon S3. In order to vend such credentials, Lake Formation assumes the
// role associated with a registered location, for example an Amazon S3 bucket,
// with a scope down policy which restricts the access to a single prefix.
//
// To call this API, the role that the service assumes must have
// lakeformation:GetDataAccess permission on the resource.
func lakeformation_GetTemporaryGlueTableCredentials(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetTemporaryGlueTableCredentialsInput{
		// TableArn: *string, // Required
	}

	if len(_lakeformationTableArn) > 0 {
		input.TableArn = aws.String(_lakeformationTableArn)
	}
	if len(_lakeformationAuditContext) > 0 {
		if err := assignInputField(input, "AuditContext", _lakeformationAuditContext); err != nil {
			log.Errorf("invalid --audit-context: %s", err.Error())
			return
		}
	}
	if len(_lakeformationDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _lakeformationDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_lakeformationPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _lakeformationPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_lakeformationQuerySessionContext) > 0 {
		if err := assignInputField(input, "QuerySessionContext", _lakeformationQuerySessionContext); err != nil {
			log.Errorf("invalid --query-session-context: %s", err.Error())
			return
		}
	}
	if len(_lakeformationS3Path) > 0 {
		input.S3Path = aws.String(_lakeformationS3Path)
	}
	if len(_lakeformationSupportedPermissionTypes) > 0 {
		if err := assignInputField(input, "SupportedPermissionTypes", _lakeformationSupportedPermissionTypes); err != nil {
			log.Errorf("invalid --supported-permission-types: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetTemporaryGlueTableCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the work units resulting from the query. Work units can be executed in
// any order and in parallel.
func lakeformation_GetWorkUnitResults(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetWorkUnitResultsInput{
		// QueryId: *string, // Required
		// WorkUnitId: int64, // Required
		// WorkUnitToken: *string, // Required
	}

	if len(_lakeformationQueryId) > 0 {
		input.QueryId = aws.String(_lakeformationQueryId)
	}
	if len(_lakeformationWorkUnitId) > 0 {
		if err := assignInputField(input, "WorkUnitId", _lakeformationWorkUnitId); err != nil {
			log.Errorf("invalid --work-unit-id: %s", err.Error())
			return
		}
	}
	if len(_lakeformationWorkUnitToken) > 0 {
		input.WorkUnitToken = aws.String(_lakeformationWorkUnitToken)
	}

	if resp, err := client.GetWorkUnitResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the work units generated by the StartQueryPlanning operation.
func lakeformation_GetWorkUnits(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GetWorkUnitsInput{
		// QueryId: *string, // Required
	}

	if len(_lakeformationQueryId) > 0 {
		input.QueryId = aws.String(_lakeformationQueryId)
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _lakeformationPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetWorkUnits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.GetWorkUnitsOutput
	p := lakeformation.NewGetWorkUnitsPaginator(client, input)
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

// Grants permissions to the principal to access metadata in the Data Catalog and
// data organized in underlying data storage such as Amazon S3.
//
// For information about permissions, see [Security and Access Control to Metadata and Data].
//
// [Security and Access Control to Metadata and Data]: https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html
func lakeformation_GrantPermissions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.GrantPermissionsInput{
		// Permissions: []types.Permission, // Required
		// Principal: *types.DataLakePrincipal, // Required
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _lakeformationPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_lakeformationPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _lakeformationPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationCondition) > 0 {
		if err := assignInputField(input, "Condition", _lakeformationCondition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}
	if len(_lakeformationPermissionsWithGrantOption) > 0 {
		if err := assignInputField(input, "PermissionsWithGrantOption", _lakeformationPermissionsWithGrantOption); err != nil {
			log.Errorf("invalid --permissions-with-grant-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.GrantPermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the data cell filters on a table.
func lakeformation_ListDataCellsFilter(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListDataCellsFilterInput{}

	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationTable) > 0 {
		if err := assignInputField(input, "Table", _lakeformationTable); err != nil {
			log.Errorf("invalid --table: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDataCellsFilter(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListDataCellsFilterOutput
	p := lakeformation.NewListDataCellsFilterPaginator(client, input)
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

// Retrieve the current list of resources and principals that are opt in to
// enforce Lake Formation permissions.
func lakeformation_ListLakeFormationOptIns(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListLakeFormationOptInsInput{}

	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _lakeformationPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLakeFormationOptIns(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListLakeFormationOptInsOutput
	p := lakeformation.NewListLakeFormationOptInsPaginator(client, input)
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

// Returns the LF-Tag expressions in caller’s account filtered based on caller's
// permissions. Data Lake and read only admins implicitly can see all tag
// expressions in their account, else caller needs DESCRIBE permissions on tag
// expression.
func lakeformation_ListLFTagExpressions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListLFTagExpressionsInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLFTagExpressions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListLFTagExpressionsOutput
	p := lakeformation.NewListLFTagExpressionsPaginator(client, input)
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

// Lists LF-tags that the requester has permission to view.
func lakeformation_ListLFTags(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListLFTagsInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationResourceShareType) > 0 {
		if err := assignInputField(input, "ResourceShareType", _lakeformationResourceShareType); err != nil {
			log.Errorf("invalid --resource-share-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListLFTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListLFTagsOutput
	p := lakeformation.NewListLFTagsPaginator(client, input)
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

// Returns a list of the principal permissions on the resource, filtered by the
// permissions of the caller. For example, if you are granted an ALTER permission,
// you are able to see only the principal permissions for ALTER.
//
// This operation returns only those permissions that have been explicitly
// granted. If both Principal and Resource parameters are provided, the response
// returns effective permissions rather than the explicitly granted permissions.
//
// For information about permissions, see [Security and Access Control to Metadata and Data].
//
// [Security and Access Control to Metadata and Data]: https://docs.aws.amazon.com/lake-formation/latest/dg/security-data-access.html
func lakeformation_ListPermissions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListPermissionsInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationIncludeRelated) > 0 {
		input.IncludeRelated = aws.String(_lakeformationIncludeRelated)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _lakeformationPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResourceType) > 0 {
		if err := assignInputField(input, "ResourceType", _lakeformationResourceType); err != nil {
			log.Errorf("invalid --resource-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListPermissionsOutput
	p := lakeformation.NewListPermissionsPaginator(client, input)
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

// Lists the resources registered to be managed by the Data Catalog.
func lakeformation_ListResources(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListResourcesInput{}

	if len(_lakeformationFilterConditionList) > 0 {
		if err := assignInputField(input, "FilterConditionList", _lakeformationFilterConditionList); err != nil {
			log.Errorf("invalid --filter-condition-list: %s", err.Error())
			return
		}
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListResourcesOutput
	p := lakeformation.NewListResourcesPaginator(client, input)
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

// Returns the configuration of all storage optimizers associated with a specified
// table.
func lakeformation_ListTableStorageOptimizers(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListTableStorageOptimizersInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
	}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationStorageOptimizerType) > 0 {
		if err := assignInputField(input, "StorageOptimizerType", _lakeformationStorageOptimizerType); err != nil {
			log.Errorf("invalid --storage-optimizer-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTableStorageOptimizers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListTableStorageOptimizersOutput
	p := lakeformation.NewListTableStorageOptimizersPaginator(client, input)
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

// Returns metadata about transactions and their status. To prevent the response
// from growing indefinitely, only uncommitted transactions and those available for
// time-travel queries are returned.
//
// This operation can help you identify uncommitted transactions or to get
// information about transactions.
func lakeformation_ListTransactions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.ListTransactionsInput{}

	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}
	if len(_lakeformationStatusFilter) > 0 {
		if err := assignInputField(input, "StatusFilter", _lakeformationStatusFilter); err != nil {
			log.Errorf("invalid --status-filter: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTransactions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.ListTransactionsOutput
	p := lakeformation.NewListTransactionsPaginator(client, input)
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

// Sets the list of data lake administrators who have admin privileges on all
// resources managed by Lake Formation. For more information on admin privileges,
// see [Granting Lake Formation Permissions].
//
// This API replaces the current list of data lake admins with the new list being
// passed. To add an admin, fetch the current list and add the new admin to that
// list and pass that list in this API.
//
// [Granting Lake Formation Permissions]: https://docs.aws.amazon.com/lake-formation/latest/dg/lake-formation-permissions.html
func lakeformation_PutDataLakeSettings(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.PutDataLakeSettingsInput{
		// DataLakeSettings: *types.DataLakeSettings, // Required
	}

	if len(_lakeformationDataLakeSettings) > 0 {
		if err := assignInputField(input, "DataLakeSettings", _lakeformationDataLakeSettings); err != nil {
			log.Errorf("invalid --data-lake-settings: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.PutDataLakeSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Registers the resource as managed by the Data Catalog.
// To add or update data, Lake Formation needs read/write access to the chosen
// data location. Choose a role that you know has permission to do this, or choose
// the AWSServiceRoleForLakeFormationDataAccess service-linked role. When you
// register the first Amazon S3 path, the service-linked role and a new inline
// policy are created on your behalf. Lake Formation adds the first path to the
// inline policy and attaches it to the service-linked role. When you register
// subsequent paths, Lake Formation adds the path to the existing policy.
//
// The following request registers a new location and gives Lake Formation
// permission to use the service-linked role to access that location.
//
// ResourceArn = arn:aws:s3:::my-bucket/ UseServiceLinkedRole = true
//
// If UseServiceLinkedRole is not set to true, you must provide or set the RoleArn :
//
// arn:aws:iam::12345:role/my-data-access-role
func lakeformation_RegisterResource(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.RegisterResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_lakeformationResourceArn) > 0 {
		input.ResourceArn = aws.String(_lakeformationResourceArn)
	}
	if len(_lakeformationExpectedResourceOwnerAccount) > 0 {
		input.ExpectedResourceOwnerAccount = aws.String(_lakeformationExpectedResourceOwnerAccount)
	}
	if len(_lakeformationHybridAccessEnabled) > 0 {
		if err := assignInputField(input, "HybridAccessEnabled", _lakeformationHybridAccessEnabled); err != nil {
			log.Errorf("invalid --hybrid-access-enabled: %s", err.Error())
			return
		}
	}
	if len(_lakeformationRoleArn) > 0 {
		input.RoleArn = aws.String(_lakeformationRoleArn)
	}
	if len(_lakeformationUseServiceLinkedRole) > 0 {
		if err := assignInputField(input, "UseServiceLinkedRole", _lakeformationUseServiceLinkedRole); err != nil {
			log.Errorf("invalid --use-service-linked-role: %s", err.Error())
			return
		}
	}
	if len(_lakeformationWithFederation) > 0 {
		if err := assignInputField(input, "WithFederation", _lakeformationWithFederation); err != nil {
			log.Errorf("invalid --with-federation: %s", err.Error())
			return
		}
	}
	if len(_lakeformationWithPrivilegedAccess) > 0 {
		if err := assignInputField(input, "WithPrivilegedAccess", _lakeformationWithPrivilegedAccess); err != nil {
			log.Errorf("invalid --with-privileged-access: %s", err.Error())
			return
		}
	}

	if resp, err := client.RegisterResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an LF-tag from the resource. Only database, table, or tableWithColumns
// resource are allowed. To tag columns, use the column inclusion list in
// tableWithColumns to specify column input.
func lakeformation_RemoveLFTagsFromResource(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.RemoveLFTagsFromResourceInput{
		// LFTags: []types.LFTagPair, // Required
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationLFTags) > 0 {
		if err := assignInputField(input, "LFTags", _lakeformationLFTags); err != nil {
			log.Errorf("invalid --lf-tags: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.RemoveLFTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes permissions to the principal to access metadata in the Data Catalog and
// data organized in underlying data storage such as Amazon S3.
func lakeformation_RevokePermissions(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.RevokePermissionsInput{
		// Permissions: []types.Permission, // Required
		// Principal: *types.DataLakePrincipal, // Required
		// Resource: *types.Resource, // Required
	}

	if len(_lakeformationPermissions) > 0 {
		if err := assignInputField(input, "Permissions", _lakeformationPermissions); err != nil {
			log.Errorf("invalid --permissions: %s", err.Error())
			return
		}
	}
	if len(_lakeformationPrincipal) > 0 {
		if err := assignInputField(input, "Principal", _lakeformationPrincipal); err != nil {
			log.Errorf("invalid --principal: %s", err.Error())
			return
		}
	}
	if len(_lakeformationResource) > 0 {
		if err := assignInputField(input, "Resource", _lakeformationResource); err != nil {
			log.Errorf("invalid --resource: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationCondition) > 0 {
		if err := assignInputField(input, "Condition", _lakeformationCondition); err != nil {
			log.Errorf("invalid --condition: %s", err.Error())
			return
		}
	}
	if len(_lakeformationPermissionsWithGrantOption) > 0 {
		if err := assignInputField(input, "PermissionsWithGrantOption", _lakeformationPermissionsWithGrantOption); err != nil {
			log.Errorf("invalid --permissions-with-grant-option: %s", err.Error())
			return
		}
	}

	if resp, err := client.RevokePermissions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation allows a search on DATABASE resources by TagCondition . This
// operation is used by admins who want to grant user permissions on certain
// TagConditions . Before making a grant, the admin can use SearchDatabasesByTags
// to find all resources where the given TagConditions are valid to verify whether
// the returned resources can be shared.
func lakeformation_SearchDatabasesByLFTags(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.SearchDatabasesByLFTagsInput{
		// Expression: []types.LFTag, // Required
	}

	if len(_lakeformationExpression) > 0 {
		if err := assignInputField(input, "Expression", _lakeformationExpression); err != nil {
			log.Errorf("invalid --expression: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchDatabasesByLFTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.SearchDatabasesByLFTagsOutput
	p := lakeformation.NewSearchDatabasesByLFTagsPaginator(client, input)
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

// This operation allows a search on TABLE resources by LFTag s. This will be used
// by admins who want to grant user permissions on certain LF-tags. Before making a
// grant, the admin can use SearchTablesByLFTags to find all resources where the
// given LFTag s are valid to verify whether the returned resources can be shared.
func lakeformation_SearchTablesByLFTags(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.SearchTablesByLFTagsInput{
		// Expression: []types.LFTag, // Required
	}

	if len(_lakeformationExpression) > 0 {
		if err := assignInputField(input, "Expression", _lakeformationExpression); err != nil {
			log.Errorf("invalid --expression: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _lakeformationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_lakeformationNextToken) > 0 {
		input.NextToken = aws.String(_lakeformationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.SearchTablesByLFTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*lakeformation.SearchTablesByLFTagsOutput
	p := lakeformation.NewSearchTablesByLFTagsPaginator(client, input)
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

// Submits a request to process a query statement.
// This operation generates work units that can be retrieved with the GetWorkUnits
// operation as soon as the query state is WORKUNITS_AVAILABLE or FINISHED.
func lakeformation_StartQueryPlanning(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.StartQueryPlanningInput{
		// QueryPlanningContext: *types.QueryPlanningContext, // Required
		// QueryString: *string, // Required
	}

	if len(_lakeformationQueryPlanningContext) > 0 {
		if err := assignInputField(input, "QueryPlanningContext", _lakeformationQueryPlanningContext); err != nil {
			log.Errorf("invalid --query-planning-context: %s", err.Error())
			return
		}
	}
	if len(_lakeformationQueryString) > 0 {
		input.QueryString = aws.String(_lakeformationQueryString)
	}

	if resp, err := client.StartQueryPlanning(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a new transaction and returns its transaction ID. Transaction IDs are
// opaque objects that you can use to identify a transaction.
func lakeformation_StartTransaction(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.StartTransactionInput{}

	if len(_lakeformationTransactionType) > 0 {
		if err := assignInputField(input, "TransactionType", _lakeformationTransactionType); err != nil {
			log.Errorf("invalid --transaction-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartTransaction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a data cell filter.
func lakeformation_UpdateDataCellsFilter(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateDataCellsFilterInput{
		// TableData: *types.DataCellsFilter, // Required
	}

	if len(_lakeformationTableData) > 0 {
		if err := assignInputField(input, "TableData", _lakeformationTableData); err != nil {
			log.Errorf("invalid --table-data: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataCellsFilter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the IAM Identity Center connection parameters.
func lakeformation_UpdateLakeFormationIdentityCenterConfiguration(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateLakeFormationIdentityCenterConfigurationInput{}

	if len(_lakeformationApplicationStatus) > 0 {
		if err := assignInputField(input, "ApplicationStatus", _lakeformationApplicationStatus); err != nil {
			log.Errorf("invalid --application-status: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationExternalFiltering) > 0 {
		if err := assignInputField(input, "ExternalFiltering", _lakeformationExternalFiltering); err != nil {
			log.Errorf("invalid --external-filtering: %s", err.Error())
			return
		}
	}
	if len(_lakeformationServiceIntegrations) > 0 {
		if err := assignInputField(input, "ServiceIntegrations", _lakeformationServiceIntegrations); err != nil {
			log.Errorf("invalid --service-integrations: %s", err.Error())
			return
		}
	}
	if len(_lakeformationShareRecipients) > 0 {
		if err := assignInputField(input, "ShareRecipients", _lakeformationShareRecipients); err != nil {
			log.Errorf("invalid --share-recipients: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLakeFormationIdentityCenterConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the list of possible values for the specified LF-tag key. If the LF-tag
// does not exist, the operation throws an EntityNotFoundException. The values in
// the delete key values will be deleted from list of possible values. If any value
// in the delete key values is attached to a resource, then API errors out with a
// 400 Exception - "Update not allowed". Untag the attribute before deleting the
// LF-tag key's value.
func lakeformation_UpdateLFTag(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateLFTagInput{
		// TagKey: *string, // Required
	}

	if len(_lakeformationTagKey) > 0 {
		input.TagKey = aws.String(_lakeformationTagKey)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationTagValuesToAdd) > 0 {
		input.TagValuesToAdd = append([]string(nil), _lakeformationTagValuesToAdd...)
	}
	if len(_lakeformationTagValuesToDelete) > 0 {
		input.TagValuesToDelete = append([]string(nil), _lakeformationTagValuesToDelete...)
	}

	if resp, err := client.UpdateLFTag(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the name of the LF-Tag expression to the new description and expression
// body provided. Updating a LF-Tag expression immediately changes the permission
// boundaries of all existing LFTagPolicy permission grants that reference the
// given LF-Tag expression.
func lakeformation_UpdateLFTagExpression(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateLFTagExpressionInput{
		// Expression: []types.LFTag, // Required
		// Name: *string, // Required
	}

	if len(_lakeformationExpression) > 0 {
		if err := assignInputField(input, "Expression", _lakeformationExpression); err != nil {
			log.Errorf("invalid --expression: %s", err.Error())
			return
		}
	}
	if len(_lakeformationName) > 0 {
		input.Name = aws.String(_lakeformationName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationDescription) > 0 {
		input.Description = aws.String(_lakeformationDescription)
	}

	if resp, err := client.UpdateLFTagExpression(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the data access role used for vending access to the given (registered)
// resource in Lake Formation.
func lakeformation_UpdateResource(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateResourceInput{
		// ResourceArn: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_lakeformationResourceArn) > 0 {
		input.ResourceArn = aws.String(_lakeformationResourceArn)
	}
	if len(_lakeformationRoleArn) > 0 {
		input.RoleArn = aws.String(_lakeformationRoleArn)
	}
	if len(_lakeformationExpectedResourceOwnerAccount) > 0 {
		input.ExpectedResourceOwnerAccount = aws.String(_lakeformationExpectedResourceOwnerAccount)
	}
	if len(_lakeformationHybridAccessEnabled) > 0 {
		if err := assignInputField(input, "HybridAccessEnabled", _lakeformationHybridAccessEnabled); err != nil {
			log.Errorf("invalid --hybrid-access-enabled: %s", err.Error())
			return
		}
	}
	if len(_lakeformationWithFederation) > 0 {
		if err := assignInputField(input, "WithFederation", _lakeformationWithFederation); err != nil {
			log.Errorf("invalid --with-federation: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the manifest of Amazon S3 objects that make up the specified governed
// table.
func lakeformation_UpdateTableObjects(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateTableObjectsInput{
		// DatabaseName: *string, // Required
		// TableName: *string, // Required
		// WriteOperations: []types.WriteOperation, // Required
	}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}
	if len(_lakeformationWriteOperations) > 0 {
		if err := assignInputField(input, "WriteOperations", _lakeformationWriteOperations); err != nil {
			log.Errorf("invalid --write-operations: %s", err.Error())
			return
		}
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}
	if len(_lakeformationTransactionId) > 0 {
		input.TransactionId = aws.String(_lakeformationTransactionId)
	}

	if resp, err := client.UpdateTableObjects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration of the storage optimizers for a table.
func lakeformation_UpdateTableStorageOptimizer(cfg aws.Config, client *lakeformation.Client) {
	input := &lakeformation.UpdateTableStorageOptimizerInput{
		// DatabaseName: *string, // Required
		// StorageOptimizerConfig: map[string]map[string]string, // Required
		// TableName: *string, // Required
	}

	if len(_lakeformationDatabaseName) > 0 {
		input.DatabaseName = aws.String(_lakeformationDatabaseName)
	}
	if len(_lakeformationStorageOptimizerConfig) > 0 {
		if err := assignInputField(input, "StorageOptimizerConfig", _lakeformationStorageOptimizerConfig); err != nil {
			log.Errorf("invalid --storage-optimizer-config: %s", err.Error())
			return
		}
	}
	if len(_lakeformationTableName) > 0 {
		input.TableName = aws.String(_lakeformationTableName)
	}
	if len(_lakeformationCatalogId) > 0 {
		input.CatalogId = aws.String(_lakeformationCatalogId)
	}

	if resp, err := client.UpdateTableStorageOptimizer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_lakeformationCmd)
	_lakeformationCmd.Flags().SortFlags = false

	_lakeformationCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_lakeformationCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_lakeformationCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_lakeformationCmd.Flags().StringVarP(&_lakeformationApplicationStatus, "application-status", "", "", "Application Status")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationAuditContext, "audit-context", "", "", "Audit Context")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationCatalogId, "catalog-id", "", "", "Catalog ID")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationCondition, "condition", "", "", "Condition")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationCredentialsScope, "credentials-scope", "", "", "Credentials Scope")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationDataLakeSettings, "data-lake-settings", "", "", "Data Lake Settings")
	_lakeformationCmd.Flags().StringSliceVarP(&_lakeformationDataLocations, "data-locations", "", nil, "Data Locations")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationDatabaseName, "database-name", "", "", "Database Name")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationDescription, "description", "", "", "Description")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationEntries, "entries", "", "", "Entries")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationExpectedResourceOwnerAccount, "expected-resource-owner-account", "", "", "Expected Resource Owner Account")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationExpression, "expression", "", "", "Expression")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationExternalFiltering, "external-filtering", "", "", "External Filtering")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationFilterConditionList, "filter-condition-list", "", "", "Filter Condition List")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationHybridAccessEnabled, "hybrid-access-enabled", "", "", "Hybrid Access Enabled")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationIncludeRelated, "include-related", "", "", "Include Related")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationInstanceArn, "instance-arn", "", "", "Instance ARN")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationLFTags, "lf-tags", "", "", "Lf Tags")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationMaxResults, "max-results", "", "", "Max Results")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationName, "name", "", "", "Name")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationNextToken, "next-token", "", "", "Next Token")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationObjects, "objects", "", "", "Objects")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPageSize, "page-size", "", "", "Page Size")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPartition, "partition", "", "", "Partition")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPartitionPredicate, "partition-predicate", "", "", "Partition Predicate")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPermissions, "permissions", "", "", "Permissions")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPermissionsWithGrantOption, "permissions-with-grant-option", "", "", "Permissions With Grant Option")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPrincipal, "principal", "", "", "Principal")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationPrincipalArn, "principal-arn", "", "", "Principal ARN")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationQueryAsOfTime, "query-as-of-time", "", "", "Query As Of Time")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationQueryId, "query-id", "", "", "Query ID")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationQueryPlanningContext, "query-planning-context", "", "", "Query Planning Context")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationQuerySessionContext, "query-session-context", "", "", "Query Session Context")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationQueryString, "query-string", "", "", "Query String")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationResource, "resource", "", "", "Resource")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationResourceArn, "resource-arn", "", "", "Resource ARN")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationResourceShareType, "resource-share-type", "", "", "Resource Share Type")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationResourceType, "resource-type", "", "", "Resource Type")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationRoleArn, "role-arn", "", "", "Role ARN")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationS3Path, "s3-path", "", "", "S3 Path")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationSAMLAssertion, "saml-assertion", "", "", "Saml Assertion")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationServiceIntegrations, "service-integrations", "", "", "Service Integrations")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationShareRecipients, "share-recipients", "", "", "Share Recipients")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationShowAssignedLFTags, "show-assigned-lf-tags", "", "", "Show Assigned Lf Tags")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationStatusFilter, "status-filter", "", "", "Status Filter")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationStorageOptimizerConfig, "storage-optimizer-config", "", "", "Storage Optimizer Config")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationStorageOptimizerType, "storage-optimizer-type", "", "", "Storage Optimizer Type")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationSupportedPermissionTypes, "supported-permission-types", "", "", "Supported Permission Types")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTable, "table", "", "", "Table")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTableArn, "table-arn", "", "", "Table ARN")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTableCatalogId, "table-catalog-id", "", "", "Table Catalog ID")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTableData, "table-data", "", "", "Table Data")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTableName, "table-name", "", "", "Table Name")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTagKey, "tag-key", "", "", "Tag Key")
	_lakeformationCmd.Flags().StringSliceVarP(&_lakeformationTagValues, "tag-values", "", nil, "Tag Values")
	_lakeformationCmd.Flags().StringSliceVarP(&_lakeformationTagValuesToAdd, "tag-values-to-add", "", nil, "Tag Values To Add")
	_lakeformationCmd.Flags().StringSliceVarP(&_lakeformationTagValuesToDelete, "tag-values-to-delete", "", nil, "Tag Values To Delete")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTransactionId, "transaction-id", "", "", "Transaction ID")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationTransactionType, "transaction-type", "", "", "Transaction Type")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationUseServiceLinkedRole, "use-service-linked-role", "", "", "Use Service Linked Role")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationWithFederation, "with-federation", "", "", "With Federation")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationWithPrivilegedAccess, "with-privileged-access", "", "", "With Privileged Access")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationWorkUnitId, "work-unit-id", "", "", "Work Unit ID")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationWorkUnitToken, "work-unit-token", "", "", "Work Unit Token")
	_lakeformationCmd.Flags().StringVarP(&_lakeformationWriteOperations, "write-operations", "", "", "Write Operations")

	_lakeformationCmd.Flags().BoolVarP(&_lakeformationAddLFTagsToResource, "add-lf-tags-to-resource", "", false, "Add Lf Tags To Resource")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationAssumeDecoratedRoleWithSAML, "assume-decorated-role-with-saml", "", false, "Assume Decorated Role With Saml")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationBatchGrantPermissions, "batch-grant-permissions", "", false, "Batch Grant Permissions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationBatchRevokePermissions, "batch-revoke-permissions", "", false, "Batch Revoke Permissions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCancelTransaction, "cancel-transaction", "", false, "Cancel Transaction")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCommitTransaction, "commit-transaction", "", false, "Commit Transaction")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCreateDataCellsFilter, "create-data-cells-filter", "", false, "Create Data Cells Filter")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCreateLakeFormationIdentityCenterConfiguration, "create-lake-formation-identity-center-configuration", "", false, "Create Lake Formation Identity Center Configuration")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCreateLakeFormationOptIn, "create-lake-formation-opt-in", "", false, "Create Lake Formation Opt In")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCreateLFTag, "create-lf-tag", "", false, "Create Lf Tag")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationCreateLFTagExpression, "create-lf-tag-expression", "", false, "Create Lf Tag Expression")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeleteDataCellsFilter, "delete-data-cells-filter", "", false, "Delete Data Cells Filter")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeleteLakeFormationIdentityCenterConfiguration, "delete-lake-formation-identity-center-configuration", "", false, "Delete Lake Formation Identity Center Configuration")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeleteLakeFormationOptIn, "delete-lake-formation-opt-in", "", false, "Delete Lake Formation Opt In")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeleteLFTag, "delete-lf-tag", "", false, "Delete Lf Tag")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeleteLFTagExpression, "delete-lf-tag-expression", "", false, "Delete Lf Tag Expression")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeleteObjectsOnCancel, "delete-objects-on-cancel", "", false, "Delete Objects On Cancel")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDeregisterResource, "deregister-resource", "", false, "Deregister Resource")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDescribeLakeFormationIdentityCenterConfiguration, "describe-lake-formation-identity-center-configuration", "", false, "Describe Lake Formation Identity Center Configuration")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDescribeResource, "describe-resource", "", false, "Describe Resource")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationDescribeTransaction, "describe-transaction", "", false, "Describe Transaction")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationExtendTransaction, "extend-transaction", "", false, "Extend Transaction")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetDataCellsFilter, "get-data-cells-filter", "", false, "Get Data Cells Filter")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetDataLakePrincipal, "get-data-lake-principal", "", false, "Get Data Lake Principal")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetDataLakeSettings, "get-data-lake-settings", "", false, "Get Data Lake Settings")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetEffectivePermissionsForPath, "get-effective-permissions-for-path", "", false, "Get Effective Permissions For Path")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetLFTag, "get-lf-tag", "", false, "Get Lf Tag")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetLFTagExpression, "get-lf-tag-expression", "", false, "Get Lf Tag Expression")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetQueryState, "get-query-state", "", false, "Get Query State")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetQueryStatistics, "get-query-statistics", "", false, "Get Query Statistics")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetResourceLFTags, "get-resource-lf-tags", "", false, "Get Resource Lf Tags")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetTableObjects, "get-table-objects", "", false, "Get Table Objects")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetTemporaryDataLocationCredentials, "get-temporary-data-location-credentials", "", false, "Get Temporary Data Location Credentials")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetTemporaryGluePartitionCredentials, "get-temporary-glue-partition-credentials", "", false, "Get Temporary Glue Partition Credentials")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetTemporaryGlueTableCredentials, "get-temporary-glue-table-credentials", "", false, "Get Temporary Glue Table Credentials")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetWorkUnitResults, "get-work-unit-results", "", false, "Get Work Unit Results")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGetWorkUnits, "get-work-units", "", false, "Get Work Units")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationGrantPermissions, "grant-permissions", "", false, "Grant Permissions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListDataCellsFilter, "list-data-cells-filter", "", false, "List Data Cells Filter")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListLakeFormationOptIns, "list-lake-formation-opt-ins", "", false, "List Lake Formation Opt Ins")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListLFTagExpressions, "list-lf-tag-expressions", "", false, "List Lf Tag Expressions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListLFTags, "list-lf-tags", "", false, "List Lf Tags")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListPermissions, "list-permissions", "", false, "List Permissions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListResources, "list-resources", "", false, "List Resources")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListTableStorageOptimizers, "list-table-storage-optimizers", "", false, "List Table Storage Optimizers")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationListTransactions, "list-transactions", "", false, "List Transactions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationPutDataLakeSettings, "put-data-lake-settings", "", false, "Put Data Lake Settings")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationRegisterResource, "register-resource", "", false, "Register Resource")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationRemoveLFTagsFromResource, "remove-lf-tags-from-resource", "", false, "Remove Lf Tags From Resource")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationRevokePermissions, "revoke-permissions", "", false, "Revoke Permissions")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationSearchDatabasesByLFTags, "search-databases-by-lf-tags", "", false, "Search Databases By Lf Tags")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationSearchTablesByLFTags, "search-tables-by-lf-tags", "", false, "Search Tables By Lf Tags")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationStartQueryPlanning, "start-query-planning", "", false, "Start Query Planning")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationStartTransaction, "start-transaction", "", false, "Start Transaction")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateDataCellsFilter, "update-data-cells-filter", "", false, "Update Data Cells Filter")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateLakeFormationIdentityCenterConfiguration, "update-lake-formation-identity-center-configuration", "", false, "Update Lake Formation Identity Center Configuration")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateLFTag, "update-lf-tag", "", false, "Update Lf Tag")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateLFTagExpression, "update-lf-tag-expression", "", false, "Update Lf Tag Expression")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateResource, "update-resource", "", false, "Update Resource")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateTableObjects, "update-table-objects", "", false, "Update Table Objects")
	_lakeformationCmd.Flags().BoolVarP(&_lakeformationUpdateTableStorageOptimizer, "update-table-storage-optimizer", "", false, "Update Table Storage Optimizer")

}
