package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/redshiftserverless"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// redshiftserverlessCmd represents the redshiftserverless command
var _redshiftserverlessCmd = &cobra.Command{
	Use:   "redshiftserverless",
	Short: "AWS redshiftserverless CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := redshiftserverless.NewFromConfig(cfg)
		if _redshiftserverlessConvertRecoveryPointToSnapshot {
			redshiftserverless_ConvertRecoveryPointToSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessCreateCustomDomainAssociation {
			redshiftserverless_CreateCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftserverlessCreateEndpointAccess {
			redshiftserverless_CreateEndpointAccess(cfg, client)
			return
		}
		if _redshiftserverlessCreateNamespace {
			redshiftserverless_CreateNamespace(cfg, client)
			return
		}
		if _redshiftserverlessCreateReservation {
			redshiftserverless_CreateReservation(cfg, client)
			return
		}
		if _redshiftserverlessCreateScheduledAction {
			redshiftserverless_CreateScheduledAction(cfg, client)
			return
		}
		if _redshiftserverlessCreateSnapshot {
			redshiftserverless_CreateSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessCreateSnapshotCopyConfiguration {
			redshiftserverless_CreateSnapshotCopyConfiguration(cfg, client)
			return
		}
		if _redshiftserverlessCreateUsageLimit {
			redshiftserverless_CreateUsageLimit(cfg, client)
			return
		}
		if _redshiftserverlessCreateWorkgroup {
			redshiftserverless_CreateWorkgroup(cfg, client)
			return
		}
		if _redshiftserverlessDeleteCustomDomainAssociation {
			redshiftserverless_DeleteCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftserverlessDeleteEndpointAccess {
			redshiftserverless_DeleteEndpointAccess(cfg, client)
			return
		}
		if _redshiftserverlessDeleteNamespace {
			redshiftserverless_DeleteNamespace(cfg, client)
			return
		}
		if _redshiftserverlessDeleteResourcePolicy {
			redshiftserverless_DeleteResourcePolicy(cfg, client)
			return
		}
		if _redshiftserverlessDeleteScheduledAction {
			redshiftserverless_DeleteScheduledAction(cfg, client)
			return
		}
		if _redshiftserverlessDeleteSnapshot {
			redshiftserverless_DeleteSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessDeleteSnapshotCopyConfiguration {
			redshiftserverless_DeleteSnapshotCopyConfiguration(cfg, client)
			return
		}
		if _redshiftserverlessDeleteUsageLimit {
			redshiftserverless_DeleteUsageLimit(cfg, client)
			return
		}
		if _redshiftserverlessDeleteWorkgroup {
			redshiftserverless_DeleteWorkgroup(cfg, client)
			return
		}
		if _redshiftserverlessGetCredentials {
			redshiftserverless_GetCredentials(cfg, client)
			return
		}
		if _redshiftserverlessGetCustomDomainAssociation {
			redshiftserverless_GetCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftserverlessGetEndpointAccess {
			redshiftserverless_GetEndpointAccess(cfg, client)
			return
		}
		if _redshiftserverlessGetIdentityCenterAuthToken {
			redshiftserverless_GetIdentityCenterAuthToken(cfg, client)
			return
		}
		if _redshiftserverlessGetNamespace {
			redshiftserverless_GetNamespace(cfg, client)
			return
		}
		if _redshiftserverlessGetRecoveryPoint {
			redshiftserverless_GetRecoveryPoint(cfg, client)
			return
		}
		if _redshiftserverlessGetReservation {
			redshiftserverless_GetReservation(cfg, client)
			return
		}
		if _redshiftserverlessGetReservationOffering {
			redshiftserverless_GetReservationOffering(cfg, client)
			return
		}
		if _redshiftserverlessGetResourcePolicy {
			redshiftserverless_GetResourcePolicy(cfg, client)
			return
		}
		if _redshiftserverlessGetScheduledAction {
			redshiftserverless_GetScheduledAction(cfg, client)
			return
		}
		if _redshiftserverlessGetSnapshot {
			redshiftserverless_GetSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessGetTableRestoreStatus {
			redshiftserverless_GetTableRestoreStatus(cfg, client)
			return
		}
		if _redshiftserverlessGetTrack {
			redshiftserverless_GetTrack(cfg, client)
			return
		}
		if _redshiftserverlessGetUsageLimit {
			redshiftserverless_GetUsageLimit(cfg, client)
			return
		}
		if _redshiftserverlessGetWorkgroup {
			redshiftserverless_GetWorkgroup(cfg, client)
			return
		}
		if _redshiftserverlessListCustomDomainAssociations {
			redshiftserverless_ListCustomDomainAssociations(cfg, client)
			return
		}
		if _redshiftserverlessListEndpointAccess {
			redshiftserverless_ListEndpointAccess(cfg, client)
			return
		}
		if _redshiftserverlessListManagedWorkgroups {
			redshiftserverless_ListManagedWorkgroups(cfg, client)
			return
		}
		if _redshiftserverlessListNamespaces {
			redshiftserverless_ListNamespaces(cfg, client)
			return
		}
		if _redshiftserverlessListRecoveryPoints {
			redshiftserverless_ListRecoveryPoints(cfg, client)
			return
		}
		if _redshiftserverlessListReservationOfferings {
			redshiftserverless_ListReservationOfferings(cfg, client)
			return
		}
		if _redshiftserverlessListReservations {
			redshiftserverless_ListReservations(cfg, client)
			return
		}
		if _redshiftserverlessListScheduledActions {
			redshiftserverless_ListScheduledActions(cfg, client)
			return
		}
		if _redshiftserverlessListSnapshotCopyConfigurations {
			redshiftserverless_ListSnapshotCopyConfigurations(cfg, client)
			return
		}
		if _redshiftserverlessListSnapshots {
			redshiftserverless_ListSnapshots(cfg, client)
			return
		}
		if _redshiftserverlessListTableRestoreStatus {
			redshiftserverless_ListTableRestoreStatus(cfg, client)
			return
		}
		if _redshiftserverlessListTagsForResource {
			redshiftserverless_ListTagsForResource(cfg, client)
			return
		}
		if _redshiftserverlessListTracks {
			redshiftserverless_ListTracks(cfg, client)
			return
		}
		if _redshiftserverlessListUsageLimits {
			redshiftserverless_ListUsageLimits(cfg, client)
			return
		}
		if _redshiftserverlessListWorkgroups {
			redshiftserverless_ListWorkgroups(cfg, client)
			return
		}
		if _redshiftserverlessPutResourcePolicy {
			redshiftserverless_PutResourcePolicy(cfg, client)
			return
		}
		if _redshiftserverlessRestoreFromRecoveryPoint {
			redshiftserverless_RestoreFromRecoveryPoint(cfg, client)
			return
		}
		if _redshiftserverlessRestoreFromSnapshot {
			redshiftserverless_RestoreFromSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessRestoreTableFromRecoveryPoint {
			redshiftserverless_RestoreTableFromRecoveryPoint(cfg, client)
			return
		}
		if _redshiftserverlessRestoreTableFromSnapshot {
			redshiftserverless_RestoreTableFromSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessTagResource {
			redshiftserverless_TagResource(cfg, client)
			return
		}
		if _redshiftserverlessUntagResource {
			redshiftserverless_UntagResource(cfg, client)
			return
		}
		if _redshiftserverlessUpdateCustomDomainAssociation {
			redshiftserverless_UpdateCustomDomainAssociation(cfg, client)
			return
		}
		if _redshiftserverlessUpdateEndpointAccess {
			redshiftserverless_UpdateEndpointAccess(cfg, client)
			return
		}
		if _redshiftserverlessUpdateLakehouseConfiguration {
			redshiftserverless_UpdateLakehouseConfiguration(cfg, client)
			return
		}
		if _redshiftserverlessUpdateNamespace {
			redshiftserverless_UpdateNamespace(cfg, client)
			return
		}
		if _redshiftserverlessUpdateScheduledAction {
			redshiftserverless_UpdateScheduledAction(cfg, client)
			return
		}
		if _redshiftserverlessUpdateSnapshot {
			redshiftserverless_UpdateSnapshot(cfg, client)
			return
		}
		if _redshiftserverlessUpdateSnapshotCopyConfiguration {
			redshiftserverless_UpdateSnapshotCopyConfiguration(cfg, client)
			return
		}
		if _redshiftserverlessUpdateUsageLimit {
			redshiftserverless_UpdateUsageLimit(cfg, client)
			return
		}
		if _redshiftserverlessUpdateWorkgroup {
			redshiftserverless_UpdateWorkgroup(cfg, client)
			return
		}

	},
}

var (
	_redshiftserverlessConvertRecoveryPointToSnapshot  bool
	_redshiftserverlessCreateCustomDomainAssociation   bool
	_redshiftserverlessCreateEndpointAccess            bool
	_redshiftserverlessCreateNamespace                 bool
	_redshiftserverlessCreateReservation               bool
	_redshiftserverlessCreateScheduledAction           bool
	_redshiftserverlessCreateSnapshot                  bool
	_redshiftserverlessCreateSnapshotCopyConfiguration bool
	_redshiftserverlessCreateUsageLimit                bool
	_redshiftserverlessCreateWorkgroup                 bool
	_redshiftserverlessDeleteCustomDomainAssociation   bool
	_redshiftserverlessDeleteEndpointAccess            bool
	_redshiftserverlessDeleteNamespace                 bool
	_redshiftserverlessDeleteResourcePolicy            bool
	_redshiftserverlessDeleteScheduledAction           bool
	_redshiftserverlessDeleteSnapshot                  bool
	_redshiftserverlessDeleteSnapshotCopyConfiguration bool
	_redshiftserverlessDeleteUsageLimit                bool
	_redshiftserverlessDeleteWorkgroup                 bool
	_redshiftserverlessGetCredentials                  bool
	_redshiftserverlessGetCustomDomainAssociation      bool
	_redshiftserverlessGetEndpointAccess               bool
	_redshiftserverlessGetIdentityCenterAuthToken      bool
	_redshiftserverlessGetNamespace                    bool
	_redshiftserverlessGetRecoveryPoint                bool
	_redshiftserverlessGetReservation                  bool
	_redshiftserverlessGetReservationOffering          bool
	_redshiftserverlessGetResourcePolicy               bool
	_redshiftserverlessGetScheduledAction              bool
	_redshiftserverlessGetSnapshot                     bool
	_redshiftserverlessGetTableRestoreStatus           bool
	_redshiftserverlessGetTrack                        bool
	_redshiftserverlessGetUsageLimit                   bool
	_redshiftserverlessGetWorkgroup                    bool
	_redshiftserverlessListCustomDomainAssociations    bool
	_redshiftserverlessListEndpointAccess              bool
	_redshiftserverlessListManagedWorkgroups           bool
	_redshiftserverlessListNamespaces                  bool
	_redshiftserverlessListRecoveryPoints              bool
	_redshiftserverlessListReservationOfferings        bool
	_redshiftserverlessListReservations                bool
	_redshiftserverlessListScheduledActions            bool
	_redshiftserverlessListSnapshotCopyConfigurations  bool
	_redshiftserverlessListSnapshots                   bool
	_redshiftserverlessListTableRestoreStatus          bool
	_redshiftserverlessListTagsForResource             bool
	_redshiftserverlessListTracks                      bool
	_redshiftserverlessListUsageLimits                 bool
	_redshiftserverlessListWorkgroups                  bool
	_redshiftserverlessPutResourcePolicy               bool
	_redshiftserverlessRestoreFromRecoveryPoint        bool
	_redshiftserverlessRestoreFromSnapshot             bool
	_redshiftserverlessRestoreTableFromRecoveryPoint   bool
	_redshiftserverlessRestoreTableFromSnapshot        bool
	_redshiftserverlessTagResource                     bool
	_redshiftserverlessUntagResource                   bool
	_redshiftserverlessUpdateCustomDomainAssociation   bool
	_redshiftserverlessUpdateEndpointAccess            bool
	_redshiftserverlessUpdateLakehouseConfiguration    bool
	_redshiftserverlessUpdateNamespace                 bool
	_redshiftserverlessUpdateScheduledAction           bool
	_redshiftserverlessUpdateSnapshot                  bool
	_redshiftserverlessUpdateSnapshotCopyConfiguration bool
	_redshiftserverlessUpdateUsageLimit                bool
	_redshiftserverlessUpdateWorkgroup                 bool

	_redshiftserverlessActivateCaseSensitiveIdentifier      string
	_redshiftserverlessAdminPasswordSecretKmsKeyId          string
	_redshiftserverlessAdminUserPassword                    string
	_redshiftserverlessAdminUsername                        string
	_redshiftserverlessAmount                               string
	_redshiftserverlessBaseCapacity                         string
	_redshiftserverlessBreachAction                         string
	_redshiftserverlessCapacity                             string
	_redshiftserverlessCatalogName                          string
	_redshiftserverlessClientToken                          string
	_redshiftserverlessConfigParameters                     string
	_redshiftserverlessCustomDomainCertificateArn           string
	_redshiftserverlessCustomDomainName                     string
	_redshiftserverlessDbName                               string
	_redshiftserverlessDefaultIamRoleArn                    string
	_redshiftserverlessDestinationKmsKeyId                  string
	_redshiftserverlessDestinationRegion                    string
	_redshiftserverlessDryRun                               string
	_redshiftserverlessDurationSeconds                      string
	_redshiftserverlessEnabled                              string
	_redshiftserverlessEndTime                              string
	_redshiftserverlessEndpointName                         string
	_redshiftserverlessEnhancedVpcRouting                   string
	_redshiftserverlessExtraComputeForAutomaticOptimization string
	_redshiftserverlessFinalSnapshotName                    string
	_redshiftserverlessFinalSnapshotRetentionPeriod         string
	_redshiftserverlessIamRoles                             []string
	_redshiftserverlessIpAddressType                        string
	_redshiftserverlessKmsKeyId                             string
	_redshiftserverlessLakehouseIdcApplicationArn           string
	_redshiftserverlessLakehouseIdcRegistration             string
	_redshiftserverlessLakehouseRegistration                string
	_redshiftserverlessLogExports                           string
	_redshiftserverlessManageAdminPassword                  string
	_redshiftserverlessMaxCapacity                          string
	_redshiftserverlessMaxResults                           string
	_redshiftserverlessNamespaceArn                         string
	_redshiftserverlessNamespaceName                        string
	_redshiftserverlessNewTableName                         string
	_redshiftserverlessNextToken                            string
	_redshiftserverlessOfferingId                           string
	_redshiftserverlessOwnerAccount                         string
	_redshiftserverlessPeriod                               string
	_redshiftserverlessPolicy                               string
	_redshiftserverlessPort                                 string
	_redshiftserverlessPricePerformanceTarget               string
	_redshiftserverlessPubliclyAccessible                   string
	_redshiftserverlessRecoveryPointId                      string
	_redshiftserverlessRedshiftIdcApplicationArn            string
	_redshiftserverlessReservationId                        string
	_redshiftserverlessResourceArn                          string
	_redshiftserverlessRetentionPeriod                      string
	_redshiftserverlessRoleArn                              string
	_redshiftserverlessSchedule                             string
	_redshiftserverlessScheduledActionDescription           string
	_redshiftserverlessScheduledActionName                  string
	_redshiftserverlessSecurityGroupIds                     []string
	_redshiftserverlessSnapshotArn                          string
	_redshiftserverlessSnapshotCopyConfigurationId          string
	_redshiftserverlessSnapshotName                         string
	_redshiftserverlessSnapshotRetentionPeriod              string
	_redshiftserverlessSourceArn                            string
	_redshiftserverlessSourceDatabaseName                   string
	_redshiftserverlessSourceSchemaName                     string
	_redshiftserverlessSourceTableName                      string
	_redshiftserverlessStartTime                            string
	_redshiftserverlessSubnetIds                            []string
	_redshiftserverlessTableRestoreRequestId                string
	_redshiftserverlessTagKeys                              []string
	_redshiftserverlessTags                                 string
	_redshiftserverlessTargetAction                         string
	_redshiftserverlessTargetDatabaseName                   string
	_redshiftserverlessTargetSchemaName                     string
	_redshiftserverlessTrackName                            string
	_redshiftserverlessUsageLimitId                         string
	_redshiftserverlessUsageType                            string
	_redshiftserverlessVpcId                                string
	_redshiftserverlessVpcSecurityGroupIds                  []string
	_redshiftserverlessWorkgroupName                        string
	_redshiftserverlessWorkgroupNames                       []string
)

// Converts a recovery point to a snapshot. For more information about recovery
// points and snapshots, see [Working with snapshots and recovery points].
//
// [Working with snapshots and recovery points]: https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-snapshots-recovery-points.html
func redshiftserverless_ConvertRecoveryPointToSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ConvertRecoveryPointToSnapshotInput{
		// RecoveryPointId: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_redshiftserverlessRecoveryPointId) > 0 {
		input.RecoveryPointId = aws.String(_redshiftserverlessRecoveryPointId)
	}
	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}
	if len(_redshiftserverlessRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _redshiftserverlessRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ConvertRecoveryPointToSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a custom domain association for Amazon Redshift Serverless.
func redshiftserverless_CreateCustomDomainAssociation(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateCustomDomainAssociationInput{
		// CustomDomainCertificateArn: *string, // Required
		// CustomDomainName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessCustomDomainCertificateArn) > 0 {
		input.CustomDomainCertificateArn = aws.String(_redshiftserverlessCustomDomainCertificateArn)
	}
	if len(_redshiftserverlessCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftserverlessCustomDomainName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.CreateCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Redshift Serverless managed VPC endpoint.
func redshiftserverless_CreateEndpointAccess(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateEndpointAccessInput{
		// EndpointName: *string, // Required
		// SubnetIds: []string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftserverlessEndpointName)
	}
	if len(_redshiftserverlessSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _redshiftserverlessSubnetIds...)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}
	if len(_redshiftserverlessOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftserverlessOwnerAccount)
	}
	if len(_redshiftserverlessVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftserverlessVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a namespace in Amazon Redshift Serverless.
func redshiftserverless_CreateNamespace(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateNamespaceInput{
		// NamespaceName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessAdminPasswordSecretKmsKeyId) > 0 {
		input.AdminPasswordSecretKmsKeyId = aws.String(_redshiftserverlessAdminPasswordSecretKmsKeyId)
	}
	if len(_redshiftserverlessAdminUserPassword) > 0 {
		input.AdminUserPassword = aws.String(_redshiftserverlessAdminUserPassword)
	}
	if len(_redshiftserverlessAdminUsername) > 0 {
		input.AdminUsername = aws.String(_redshiftserverlessAdminUsername)
	}
	if len(_redshiftserverlessDbName) > 0 {
		input.DbName = aws.String(_redshiftserverlessDbName)
	}
	if len(_redshiftserverlessDefaultIamRoleArn) > 0 {
		input.DefaultIamRoleArn = aws.String(_redshiftserverlessDefaultIamRoleArn)
	}
	if len(_redshiftserverlessIamRoles) > 0 {
		input.IamRoles = append([]string(nil), _redshiftserverlessIamRoles...)
	}
	if len(_redshiftserverlessKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_redshiftserverlessKmsKeyId)
	}
	if len(_redshiftserverlessLogExports) > 0 {
		if err := assignInputField(input, "LogExports", _redshiftserverlessLogExports); err != nil {
			log.Errorf("invalid --log-exports: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessManageAdminPassword) > 0 {
		if err := assignInputField(input, "ManageAdminPassword", _redshiftserverlessManageAdminPassword); err != nil {
			log.Errorf("invalid --manage-admin-password: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessRedshiftIdcApplicationArn) > 0 {
		input.RedshiftIdcApplicationArn = aws.String(_redshiftserverlessRedshiftIdcApplicationArn)
	}
	if len(_redshiftserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon Redshift Serverless reservation, which gives you the option
// to commit to a specified number of Redshift Processing Units (RPUs) for a year
// at a discount from Serverless on-demand (OD) rates.
func redshiftserverless_CreateReservation(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateReservationInput{
		// Capacity: int32, // Required
		// OfferingId: *string, // Required
	}

	if len(_redshiftserverlessCapacity) > 0 {
		if err := assignInputField(input, "Capacity", _redshiftserverlessCapacity); err != nil {
			log.Errorf("invalid --capacity: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessOfferingId) > 0 {
		input.OfferingId = aws.String(_redshiftserverlessOfferingId)
	}
	if len(_redshiftserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_redshiftserverlessClientToken)
	}

	if resp, err := client.CreateReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a scheduled action. A scheduled action contains a schedule and an
// Amazon Redshift API action. For example, you can create a schedule of when to
// run the CreateSnapshot API operation.
func redshiftserverless_CreateScheduledAction(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateScheduledActionInput{
		// NamespaceName: *string, // Required
		// RoleArn: *string, // Required
		// Schedule: types.Schedule, // Required
		// ScheduledActionName: *string, // Required
		// TargetAction: types.TargetAction, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessRoleArn) > 0 {
		input.RoleArn = aws.String(_redshiftserverlessRoleArn)
	}
	if len(_redshiftserverlessSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _redshiftserverlessSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftserverlessScheduledActionName)
	}
	if len(_redshiftserverlessTargetAction) > 0 {
		if err := assignInputField(input, "TargetAction", _redshiftserverlessTargetAction); err != nil {
			log.Errorf("invalid --target-action: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _redshiftserverlessEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftserverlessEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessScheduledActionDescription) > 0 {
		input.ScheduledActionDescription = aws.String(_redshiftserverlessScheduledActionDescription)
	}
	if len(_redshiftserverlessStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftserverlessStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of all databases in a namespace. For more information about
// snapshots, see [Working with snapshots and recovery points].
//
// [Working with snapshots and recovery points]: https://docs.aws.amazon.com/redshift/latest/mgmt/serverless-snapshots-recovery-points.html
func redshiftserverless_CreateSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateSnapshotInput{
		// NamespaceName: *string, // Required
		// SnapshotName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}
	if len(_redshiftserverlessRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _redshiftserverlessRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot copy configuration that lets you copy snapshots to another
// Amazon Web Services Region.
func redshiftserverless_CreateSnapshotCopyConfiguration(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateSnapshotCopyConfigurationInput{
		// DestinationRegion: *string, // Required
		// NamespaceName: *string, // Required
	}

	if len(_redshiftserverlessDestinationRegion) > 0 {
		input.DestinationRegion = aws.String(_redshiftserverlessDestinationRegion)
	}
	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessDestinationKmsKeyId) > 0 {
		input.DestinationKmsKeyId = aws.String(_redshiftserverlessDestinationKmsKeyId)
	}
	if len(_redshiftserverlessSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "SnapshotRetentionPeriod", _redshiftserverlessSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --snapshot-retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSnapshotCopyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a usage limit for a specified Amazon Redshift Serverless usage type.
// The usage limit is identified by the returned usage limit identifier.
func redshiftserverless_CreateUsageLimit(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateUsageLimitInput{
		// Amount: *int64, // Required
		// ResourceArn: *string, // Required
		// UsageType: types.UsageLimitUsageType, // Required
	}

	if len(_redshiftserverlessAmount) > 0 {
		if err := assignInputField(input, "Amount", _redshiftserverlessAmount); err != nil {
			log.Errorf("invalid --amount: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}
	if len(_redshiftserverlessUsageType) > 0 {
		if err := assignInputField(input, "UsageType", _redshiftserverlessUsageType); err != nil {
			log.Errorf("invalid --usage-type: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessBreachAction) > 0 {
		if err := assignInputField(input, "BreachAction", _redshiftserverlessBreachAction); err != nil {
			log.Errorf("invalid --breach-action: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPeriod) > 0 {
		if err := assignInputField(input, "Period", _redshiftserverlessPeriod); err != nil {
			log.Errorf("invalid --period: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an workgroup in Amazon Redshift Serverless.
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// workgroup is in an account with VPC BPA turned on, the following capabilities
// are blocked:
//
// - Creating a public access workgroup
//
// - Modifying a private workgroup to public
//
// - Adding a subnet with VPC BPA turned on to the workgroup when the workgroup
// is public
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
func redshiftserverless_CreateWorkgroup(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.CreateWorkgroupInput{
		// NamespaceName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}
	if len(_redshiftserverlessBaseCapacity) > 0 {
		if err := assignInputField(input, "BaseCapacity", _redshiftserverlessBaseCapacity); err != nil {
			log.Errorf("invalid --base-capacity: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessConfigParameters) > 0 {
		if err := assignInputField(input, "ConfigParameters", _redshiftserverlessConfigParameters); err != nil {
			log.Errorf("invalid --config-parameters: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessEnhancedVpcRouting) > 0 {
		if err := assignInputField(input, "EnhancedVpcRouting", _redshiftserverlessEnhancedVpcRouting); err != nil {
			log.Errorf("invalid --enhanced-vpc-routing: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessExtraComputeForAutomaticOptimization) > 0 {
		if err := assignInputField(input, "ExtraComputeForAutomaticOptimization", _redshiftserverlessExtraComputeForAutomaticOptimization); err != nil {
			log.Errorf("invalid --extra-compute-for-automatic-optimization: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessIpAddressType) > 0 {
		input.IpAddressType = aws.String(_redshiftserverlessIpAddressType)
	}
	if len(_redshiftserverlessMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _redshiftserverlessMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPort) > 0 {
		if err := assignInputField(input, "Port", _redshiftserverlessPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPricePerformanceTarget) > 0 {
		if err := assignInputField(input, "PricePerformanceTarget", _redshiftserverlessPricePerformanceTarget); err != nil {
			log.Errorf("invalid --price-performance-target: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _redshiftserverlessPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _redshiftserverlessSecurityGroupIds...)
	}
	if len(_redshiftserverlessSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _redshiftserverlessSubnetIds...)
	}
	if len(_redshiftserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessTrackName) > 0 {
		input.TrackName = aws.String(_redshiftserverlessTrackName)
	}

	if resp, err := client.CreateWorkgroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a custom domain association for Amazon Redshift Serverless.
func redshiftserverless_DeleteCustomDomainAssociation(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteCustomDomainAssociationInput{
		// CustomDomainName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftserverlessCustomDomainName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.DeleteCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon Redshift Serverless managed VPC endpoint.
func redshiftserverless_DeleteEndpointAccess(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteEndpointAccessInput{
		// EndpointName: *string, // Required
	}

	if len(_redshiftserverlessEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftserverlessEndpointName)
	}

	if resp, err := client.DeleteEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a namespace from Amazon Redshift Serverless. Before you delete the
// namespace, you can create a final snapshot that has all of the data within the
// namespace.
func redshiftserverless_DeleteNamespace(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteNamespaceInput{
		// NamespaceName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessFinalSnapshotName) > 0 {
		input.FinalSnapshotName = aws.String(_redshiftserverlessFinalSnapshotName)
	}
	if len(_redshiftserverlessFinalSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "FinalSnapshotRetentionPeriod", _redshiftserverlessFinalSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --final-snapshot-retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource policy.
func redshiftserverless_DeleteResourcePolicy(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a scheduled action.
func redshiftserverless_DeleteScheduledAction(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteScheduledActionInput{
		// ScheduledActionName: *string, // Required
	}

	if len(_redshiftserverlessScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftserverlessScheduledActionName)
	}

	if resp, err := client.DeleteScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a snapshot from Amazon Redshift Serverless.
func redshiftserverless_DeleteSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteSnapshotInput{
		// SnapshotName: *string, // Required
	}

	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}

	if resp, err := client.DeleteSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a snapshot copy configuration
func redshiftserverless_DeleteSnapshotCopyConfiguration(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteSnapshotCopyConfigurationInput{
		// SnapshotCopyConfigurationId: *string, // Required
	}

	if len(_redshiftserverlessSnapshotCopyConfigurationId) > 0 {
		input.SnapshotCopyConfigurationId = aws.String(_redshiftserverlessSnapshotCopyConfigurationId)
	}

	if resp, err := client.DeleteSnapshotCopyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a usage limit from Amazon Redshift Serverless.
func redshiftserverless_DeleteUsageLimit(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteUsageLimitInput{
		// UsageLimitId: *string, // Required
	}

	if len(_redshiftserverlessUsageLimitId) > 0 {
		input.UsageLimitId = aws.String(_redshiftserverlessUsageLimitId)
	}

	if resp, err := client.DeleteUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a workgroup.
func redshiftserverless_DeleteWorkgroup(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.DeleteWorkgroupInput{
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.DeleteWorkgroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a database user name and temporary password with temporary
// authorization to log in to Amazon Redshift Serverless.
//
// By default, the temporary credentials expire in 900 seconds. You can optionally
// specify a duration between 900 seconds (15 minutes) and 3600 seconds (60
// minutes).
//
// The Identity and Access Management (IAM) user or role that runs GetCredentials
// must have an IAM policy attached that allows access to all necessary actions and
// resources.
//
// If the DbName parameter is specified, the IAM policy must allow access to the
// resource dbname for the specified database name.
func redshiftserverless_GetCredentials(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetCredentialsInput{}

	if len(_redshiftserverlessCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftserverlessCustomDomainName)
	}
	if len(_redshiftserverlessDbName) > 0 {
		input.DbName = aws.String(_redshiftserverlessDbName)
	}
	if len(_redshiftserverlessDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _redshiftserverlessDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.GetCredentials(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about a specific custom domain association.
func redshiftserverless_GetCustomDomainAssociation(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetCustomDomainAssociationInput{
		// CustomDomainName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftserverlessCustomDomainName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.GetCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information, such as the name, about a VPC endpoint.
func redshiftserverless_GetEndpointAccess(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetEndpointAccessInput{
		// EndpointName: *string, // Required
	}

	if len(_redshiftserverlessEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftserverlessEndpointName)
	}

	if resp, err := client.GetEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an Identity Center authentication token for accessing Amazon Redshift
// Serverless workgroups.
//
// The token provides secure access to data within the specified workgroups using
// Identity Center identity propagation. The token expires after a specified
// duration and must be refreshed for continued access.
//
// The Identity and Access Management (IAM) user or role that runs
// GetIdentityCenterAuthToken must have appropriate permissions to access the
// specified workgroups and Identity Center integration must be configured for the
// workgroups.
func redshiftserverless_GetIdentityCenterAuthToken(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetIdentityCenterAuthTokenInput{
		// WorkgroupNames: []string, // Required
	}

	if len(_redshiftserverlessWorkgroupNames) > 0 {
		input.WorkgroupNames = append([]string(nil), _redshiftserverlessWorkgroupNames...)
	}

	if resp, err := client.GetIdentityCenterAuthToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a namespace in Amazon Redshift Serverless.
func redshiftserverless_GetNamespace(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetNamespaceInput{
		// NamespaceName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}

	if resp, err := client.GetNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a recovery point.
func redshiftserverless_GetRecoveryPoint(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetRecoveryPointInput{
		// RecoveryPointId: *string, // Required
	}

	if len(_redshiftserverlessRecoveryPointId) > 0 {
		input.RecoveryPointId = aws.String(_redshiftserverlessRecoveryPointId)
	}

	if resp, err := client.GetRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon Redshift Serverless reservation. A reservation gives you the
// option to commit to a specified number of Redshift Processing Units (RPUs) for a
// year at a discount from Serverless on-demand (OD) rates.
func redshiftserverless_GetReservation(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetReservationInput{
		// ReservationId: *string, // Required
	}

	if len(_redshiftserverlessReservationId) > 0 {
		input.ReservationId = aws.String(_redshiftserverlessReservationId)
	}

	if resp, err := client.GetReservation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the reservation offering. The offering determines the payment schedule
// for the reservation.
func redshiftserverless_GetReservationOffering(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetReservationOfferingInput{
		// OfferingId: *string, // Required
	}

	if len(_redshiftserverlessOfferingId) > 0 {
		input.OfferingId = aws.String(_redshiftserverlessOfferingId)
	}

	if resp, err := client.GetReservationOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a resource policy.
func redshiftserverless_GetResourcePolicy(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetResourcePolicyInput{
		// ResourceArn: *string, // Required
	}

	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a scheduled action.
func redshiftserverless_GetScheduledAction(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetScheduledActionInput{
		// ScheduledActionName: *string, // Required
	}

	if len(_redshiftserverlessScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftserverlessScheduledActionName)
	}

	if resp, err := client.GetScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific snapshot.
func redshiftserverless_GetSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetSnapshotInput{}

	if len(_redshiftserverlessOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftserverlessOwnerAccount)
	}
	if len(_redshiftserverlessSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftserverlessSnapshotArn)
	}
	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}

	if resp, err := client.GetSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a TableRestoreStatus object.
func redshiftserverless_GetTableRestoreStatus(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetTableRestoreStatusInput{
		// TableRestoreRequestId: *string, // Required
	}

	if len(_redshiftserverlessTableRestoreRequestId) > 0 {
		input.TableRestoreRequestId = aws.String(_redshiftserverlessTableRestoreRequestId)
	}

	if resp, err := client.GetTableRestoreStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the Redshift Serverless version for a specified track.
func redshiftserverless_GetTrack(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetTrackInput{
		// TrackName: *string, // Required
	}

	if len(_redshiftserverlessTrackName) > 0 {
		input.TrackName = aws.String(_redshiftserverlessTrackName)
	}

	if resp, err := client.GetTrack(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a usage limit.
func redshiftserverless_GetUsageLimit(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetUsageLimitInput{
		// UsageLimitId: *string, // Required
	}

	if len(_redshiftserverlessUsageLimitId) > 0 {
		input.UsageLimitId = aws.String(_redshiftserverlessUsageLimitId)
	}

	if resp, err := client.GetUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific workgroup.
func redshiftserverless_GetWorkgroup(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.GetWorkgroupInput{
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.GetWorkgroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists custom domain associations for Amazon Redshift Serverless.
func redshiftserverless_ListCustomDomainAssociations(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListCustomDomainAssociationsInput{}

	if len(_redshiftserverlessCustomDomainCertificateArn) > 0 {
		input.CustomDomainCertificateArn = aws.String(_redshiftserverlessCustomDomainCertificateArn)
	}
	if len(_redshiftserverlessCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftserverlessCustomDomainName)
	}
	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCustomDomainAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListCustomDomainAssociationsOutput
	p := redshiftserverless.NewListCustomDomainAssociationsPaginator(client, input)
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

// Returns an array of EndpointAccess objects and relevant information.
func redshiftserverless_ListEndpointAccess(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListEndpointAccessInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftserverlessOwnerAccount)
	}
	if len(_redshiftserverlessVpcId) > 0 {
		input.VpcId = aws.String(_redshiftserverlessVpcId)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if disablePaginator() {
		if resp, err := client.ListEndpointAccess(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListEndpointAccessOutput
	p := redshiftserverless.NewListEndpointAccessPaginator(client, input)
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

// Returns information about a list of specified managed workgroups in your
// account.
func redshiftserverless_ListManagedWorkgroups(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListManagedWorkgroupsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessSourceArn) > 0 {
		input.SourceArn = aws.String(_redshiftserverlessSourceArn)
	}

	if disablePaginator() {
		if resp, err := client.ListManagedWorkgroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListManagedWorkgroupsOutput
	p := redshiftserverless.NewListManagedWorkgroupsPaginator(client, input)
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

// Returns information about a list of specified namespaces.
func redshiftserverless_ListNamespaces(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListNamespacesInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
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

	var results []*redshiftserverless.ListNamespacesOutput
	p := redshiftserverless.NewListNamespacesPaginator(client, input)
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

// Returns an array of recovery points.
func redshiftserverless_ListRecoveryPoints(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListRecoveryPointsInput{}

	if len(_redshiftserverlessEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftserverlessEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNamespaceArn) > 0 {
		input.NamespaceArn = aws.String(_redshiftserverlessNamespaceArn)
	}
	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftserverlessStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRecoveryPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListRecoveryPointsOutput
	p := redshiftserverless.NewListRecoveryPointsPaginator(client, input)
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

// Returns the current reservation offerings in your account.
func redshiftserverless_ListReservationOfferings(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListReservationOfferingsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReservationOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListReservationOfferingsOutput
	p := redshiftserverless.NewListReservationOfferingsPaginator(client, input)
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

// Returns a list of Reservation objects.
func redshiftserverless_ListReservations(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListReservationsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListReservations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListReservationsOutput
	p := redshiftserverless.NewListReservationsPaginator(client, input)
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

// Returns a list of scheduled actions. You can use the flags to filter the list
// of returned scheduled actions.
func redshiftserverless_ListScheduledActions(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListScheduledActionsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListScheduledActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListScheduledActionsOutput
	p := redshiftserverless.NewListScheduledActionsPaginator(client, input)
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

// Returns a list of snapshot copy configurations.
func redshiftserverless_ListSnapshotCopyConfigurations(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListSnapshotCopyConfigurationsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSnapshotCopyConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListSnapshotCopyConfigurationsOutput
	p := redshiftserverless.NewListSnapshotCopyConfigurationsPaginator(client, input)
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

// Returns a list of snapshots.
func redshiftserverless_ListSnapshots(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListSnapshotsInput{}

	if len(_redshiftserverlessEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftserverlessEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNamespaceArn) > 0 {
		input.NamespaceArn = aws.String(_redshiftserverlessNamespaceArn)
	}
	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftserverlessOwnerAccount)
	}
	if len(_redshiftserverlessStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftserverlessStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListSnapshotsOutput
	p := redshiftserverless.NewListSnapshotsPaginator(client, input)
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

// Returns information about an array of TableRestoreStatus objects.
func redshiftserverless_ListTableRestoreStatus(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListTableRestoreStatusInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if disablePaginator() {
		if resp, err := client.ListTableRestoreStatus(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListTableRestoreStatusOutput
	p := redshiftserverless.NewListTableRestoreStatusPaginator(client, input)
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

// Lists the tags assigned to a resource.
func redshiftserverless_ListTagsForResource(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List the Amazon Redshift Serverless versions.
func redshiftserverless_ListTracks(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListTracksInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTracks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListTracksOutput
	p := redshiftserverless.NewListTracksPaginator(client, input)
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

// Lists all usage limits within Amazon Redshift Serverless.
func redshiftserverless_ListUsageLimits(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListUsageLimitsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}
	if len(_redshiftserverlessUsageType) > 0 {
		if err := assignInputField(input, "UsageType", _redshiftserverlessUsageType); err != nil {
			log.Errorf("invalid --usage-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListUsageLimits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListUsageLimitsOutput
	p := redshiftserverless.NewListUsageLimitsPaginator(client, input)
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

// Returns information about a list of specified workgroups.
func redshiftserverless_ListWorkgroups(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.ListWorkgroupsInput{}

	if len(_redshiftserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _redshiftserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessNextToken) > 0 {
		input.NextToken = aws.String(_redshiftserverlessNextToken)
	}
	if len(_redshiftserverlessOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftserverlessOwnerAccount)
	}

	if disablePaginator() {
		if resp, err := client.ListWorkgroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*redshiftserverless.ListWorkgroupsOutput
	p := redshiftserverless.NewListWorkgroupsPaginator(client, input)
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

// Creates or updates a resource policy. Currently, you can use policies to share
// snapshots across Amazon Web Services accounts.
func redshiftserverless_PutResourcePolicy(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.PutResourcePolicyInput{
		// Policy: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_redshiftserverlessPolicy) > 0 {
		input.Policy = aws.String(_redshiftserverlessPolicy)
	}
	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restore the data from a recovery point.
func redshiftserverless_RestoreFromRecoveryPoint(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.RestoreFromRecoveryPointInput{
		// NamespaceName: *string, // Required
		// RecoveryPointId: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessRecoveryPointId) > 0 {
		input.RecoveryPointId = aws.String(_redshiftserverlessRecoveryPointId)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.RestoreFromRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a namespace from a snapshot.
func redshiftserverless_RestoreFromSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.RestoreFromSnapshotInput{
		// NamespaceName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}
	if len(_redshiftserverlessAdminPasswordSecretKmsKeyId) > 0 {
		input.AdminPasswordSecretKmsKeyId = aws.String(_redshiftserverlessAdminPasswordSecretKmsKeyId)
	}
	if len(_redshiftserverlessManageAdminPassword) > 0 {
		if err := assignInputField(input, "ManageAdminPassword", _redshiftserverlessManageAdminPassword); err != nil {
			log.Errorf("invalid --manage-admin-password: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessOwnerAccount) > 0 {
		input.OwnerAccount = aws.String(_redshiftserverlessOwnerAccount)
	}
	if len(_redshiftserverlessSnapshotArn) > 0 {
		input.SnapshotArn = aws.String(_redshiftserverlessSnapshotArn)
	}
	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}

	if resp, err := client.RestoreFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a table from a recovery point to your Amazon Redshift Serverless
// instance. You can't use this operation to restore tables with interleaved sort
// keys.
func redshiftserverless_RestoreTableFromRecoveryPoint(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.RestoreTableFromRecoveryPointInput{
		// NamespaceName: *string, // Required
		// NewTableName: *string, // Required
		// RecoveryPointId: *string, // Required
		// SourceDatabaseName: *string, // Required
		// SourceTableName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNewTableName) > 0 {
		input.NewTableName = aws.String(_redshiftserverlessNewTableName)
	}
	if len(_redshiftserverlessRecoveryPointId) > 0 {
		input.RecoveryPointId = aws.String(_redshiftserverlessRecoveryPointId)
	}
	if len(_redshiftserverlessSourceDatabaseName) > 0 {
		input.SourceDatabaseName = aws.String(_redshiftserverlessSourceDatabaseName)
	}
	if len(_redshiftserverlessSourceTableName) > 0 {
		input.SourceTableName = aws.String(_redshiftserverlessSourceTableName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}
	if len(_redshiftserverlessActivateCaseSensitiveIdentifier) > 0 {
		if err := assignInputField(input, "ActivateCaseSensitiveIdentifier", _redshiftserverlessActivateCaseSensitiveIdentifier); err != nil {
			log.Errorf("invalid --activate-case-sensitive-identifier: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessSourceSchemaName) > 0 {
		input.SourceSchemaName = aws.String(_redshiftserverlessSourceSchemaName)
	}
	if len(_redshiftserverlessTargetDatabaseName) > 0 {
		input.TargetDatabaseName = aws.String(_redshiftserverlessTargetDatabaseName)
	}
	if len(_redshiftserverlessTargetSchemaName) > 0 {
		input.TargetSchemaName = aws.String(_redshiftserverlessTargetSchemaName)
	}

	if resp, err := client.RestoreTableFromRecoveryPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a table from a snapshot to your Amazon Redshift Serverless instance.
// You can't use this operation to restore tables with [interleaved sort keys].
//
// [interleaved sort keys]: https://docs.aws.amazon.com/redshift/latest/dg/t_Sorting_data.html#t_Sorting_data-interleaved
func redshiftserverless_RestoreTableFromSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.RestoreTableFromSnapshotInput{
		// NamespaceName: *string, // Required
		// NewTableName: *string, // Required
		// SnapshotName: *string, // Required
		// SourceDatabaseName: *string, // Required
		// SourceTableName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessNewTableName) > 0 {
		input.NewTableName = aws.String(_redshiftserverlessNewTableName)
	}
	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}
	if len(_redshiftserverlessSourceDatabaseName) > 0 {
		input.SourceDatabaseName = aws.String(_redshiftserverlessSourceDatabaseName)
	}
	if len(_redshiftserverlessSourceTableName) > 0 {
		input.SourceTableName = aws.String(_redshiftserverlessSourceTableName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}
	if len(_redshiftserverlessActivateCaseSensitiveIdentifier) > 0 {
		if err := assignInputField(input, "ActivateCaseSensitiveIdentifier", _redshiftserverlessActivateCaseSensitiveIdentifier); err != nil {
			log.Errorf("invalid --activate-case-sensitive-identifier: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessSourceSchemaName) > 0 {
		input.SourceSchemaName = aws.String(_redshiftserverlessSourceSchemaName)
	}
	if len(_redshiftserverlessTargetDatabaseName) > 0 {
		input.TargetDatabaseName = aws.String(_redshiftserverlessTargetDatabaseName)
	}
	if len(_redshiftserverlessTargetSchemaName) > 0 {
		input.TargetSchemaName = aws.String(_redshiftserverlessTargetSchemaName)
	}

	if resp, err := client.RestoreTableFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags to a resource.
func redshiftserverless_TagResource(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}
	if len(_redshiftserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _redshiftserverlessTags); err != nil {
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

// Removes a tag or set of tags from a resource.
func redshiftserverless_UntagResource(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_redshiftserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_redshiftserverlessResourceArn)
	}
	if len(_redshiftserverlessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _redshiftserverlessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Redshift Serverless certificate associated with a custom
// domain.
func redshiftserverless_UpdateCustomDomainAssociation(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateCustomDomainAssociationInput{
		// CustomDomainCertificateArn: *string, // Required
		// CustomDomainName: *string, // Required
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessCustomDomainCertificateArn) > 0 {
		input.CustomDomainCertificateArn = aws.String(_redshiftserverlessCustomDomainCertificateArn)
	}
	if len(_redshiftserverlessCustomDomainName) > 0 {
		input.CustomDomainName = aws.String(_redshiftserverlessCustomDomainName)
	}
	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}

	if resp, err := client.UpdateCustomDomainAssociation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an Amazon Redshift Serverless managed endpoint.
func redshiftserverless_UpdateEndpointAccess(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateEndpointAccessInput{
		// EndpointName: *string, // Required
	}

	if len(_redshiftserverlessEndpointName) > 0 {
		input.EndpointName = aws.String(_redshiftserverlessEndpointName)
	}
	if len(_redshiftserverlessVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _redshiftserverlessVpcSecurityGroupIds...)
	}

	if resp, err := client.UpdateEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the lakehouse configuration for a namespace. This operation allows you
// to manage Amazon Redshift federated permissions and Amazon Web Services IAM
// Identity Center trusted identity propagation.
func redshiftserverless_UpdateLakehouseConfiguration(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateLakehouseConfigurationInput{
		// NamespaceName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessCatalogName) > 0 {
		input.CatalogName = aws.String(_redshiftserverlessCatalogName)
	}
	if len(_redshiftserverlessDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _redshiftserverlessDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessLakehouseIdcApplicationArn) > 0 {
		input.LakehouseIdcApplicationArn = aws.String(_redshiftserverlessLakehouseIdcApplicationArn)
	}
	if len(_redshiftserverlessLakehouseIdcRegistration) > 0 {
		if err := assignInputField(input, "LakehouseIdcRegistration", _redshiftserverlessLakehouseIdcRegistration); err != nil {
			log.Errorf("invalid --lakehouse-idc-registration: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessLakehouseRegistration) > 0 {
		if err := assignInputField(input, "LakehouseRegistration", _redshiftserverlessLakehouseRegistration); err != nil {
			log.Errorf("invalid --lakehouse-registration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLakehouseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a namespace with the specified settings. Unless required, you can't
// update multiple parameters in one request. For example, you must specify both
// adminUsername and adminUserPassword to update either field, but you can't
// update both kmsKeyId and logExports in a single request.
func redshiftserverless_UpdateNamespace(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateNamespaceInput{
		// NamespaceName: *string, // Required
	}

	if len(_redshiftserverlessNamespaceName) > 0 {
		input.NamespaceName = aws.String(_redshiftserverlessNamespaceName)
	}
	if len(_redshiftserverlessAdminPasswordSecretKmsKeyId) > 0 {
		input.AdminPasswordSecretKmsKeyId = aws.String(_redshiftserverlessAdminPasswordSecretKmsKeyId)
	}
	if len(_redshiftserverlessAdminUserPassword) > 0 {
		input.AdminUserPassword = aws.String(_redshiftserverlessAdminUserPassword)
	}
	if len(_redshiftserverlessAdminUsername) > 0 {
		input.AdminUsername = aws.String(_redshiftserverlessAdminUsername)
	}
	if len(_redshiftserverlessDefaultIamRoleArn) > 0 {
		input.DefaultIamRoleArn = aws.String(_redshiftserverlessDefaultIamRoleArn)
	}
	if len(_redshiftserverlessIamRoles) > 0 {
		input.IamRoles = append([]string(nil), _redshiftserverlessIamRoles...)
	}
	if len(_redshiftserverlessKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_redshiftserverlessKmsKeyId)
	}
	if len(_redshiftserverlessLogExports) > 0 {
		if err := assignInputField(input, "LogExports", _redshiftserverlessLogExports); err != nil {
			log.Errorf("invalid --log-exports: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessManageAdminPassword) > 0 {
		if err := assignInputField(input, "ManageAdminPassword", _redshiftserverlessManageAdminPassword); err != nil {
			log.Errorf("invalid --manage-admin-password: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateNamespace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a scheduled action.
func redshiftserverless_UpdateScheduledAction(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateScheduledActionInput{
		// ScheduledActionName: *string, // Required
	}

	if len(_redshiftserverlessScheduledActionName) > 0 {
		input.ScheduledActionName = aws.String(_redshiftserverlessScheduledActionName)
	}
	if len(_redshiftserverlessEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _redshiftserverlessEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _redshiftserverlessEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessRoleArn) > 0 {
		input.RoleArn = aws.String(_redshiftserverlessRoleArn)
	}
	if len(_redshiftserverlessSchedule) > 0 {
		if err := assignInputField(input, "Schedule", _redshiftserverlessSchedule); err != nil {
			log.Errorf("invalid --schedule: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessScheduledActionDescription) > 0 {
		input.ScheduledActionDescription = aws.String(_redshiftserverlessScheduledActionDescription)
	}
	if len(_redshiftserverlessStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _redshiftserverlessStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessTargetAction) > 0 {
		if err := assignInputField(input, "TargetAction", _redshiftserverlessTargetAction); err != nil {
			log.Errorf("invalid --target-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateScheduledAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a snapshot.
func redshiftserverless_UpdateSnapshot(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateSnapshotInput{
		// SnapshotName: *string, // Required
	}

	if len(_redshiftserverlessSnapshotName) > 0 {
		input.SnapshotName = aws.String(_redshiftserverlessSnapshotName)
	}
	if len(_redshiftserverlessRetentionPeriod) > 0 {
		if err := assignInputField(input, "RetentionPeriod", _redshiftserverlessRetentionPeriod); err != nil {
			log.Errorf("invalid --retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a snapshot copy configuration.
func redshiftserverless_UpdateSnapshotCopyConfiguration(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateSnapshotCopyConfigurationInput{
		// SnapshotCopyConfigurationId: *string, // Required
	}

	if len(_redshiftserverlessSnapshotCopyConfigurationId) > 0 {
		input.SnapshotCopyConfigurationId = aws.String(_redshiftserverlessSnapshotCopyConfigurationId)
	}
	if len(_redshiftserverlessSnapshotRetentionPeriod) > 0 {
		if err := assignInputField(input, "SnapshotRetentionPeriod", _redshiftserverlessSnapshotRetentionPeriod); err != nil {
			log.Errorf("invalid --snapshot-retention-period: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSnapshotCopyConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update a usage limit in Amazon Redshift Serverless. You can't update the usage
// type or period of a usage limit.
func redshiftserverless_UpdateUsageLimit(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateUsageLimitInput{
		// UsageLimitId: *string, // Required
	}

	if len(_redshiftserverlessUsageLimitId) > 0 {
		input.UsageLimitId = aws.String(_redshiftserverlessUsageLimitId)
	}
	if len(_redshiftserverlessAmount) > 0 {
		if err := assignInputField(input, "Amount", _redshiftserverlessAmount); err != nil {
			log.Errorf("invalid --amount: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessBreachAction) > 0 {
		if err := assignInputField(input, "BreachAction", _redshiftserverlessBreachAction); err != nil {
			log.Errorf("invalid --breach-action: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUsageLimit(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a workgroup with the specified configuration settings. You can't update
// multiple parameters in one request. For example, you can update baseCapacity or
// port in a single request, but you can't update both in the same request.
//
// VPC Block Public Access (BPA) enables you to block resources in VPCs and
// subnets that you own in a Region from reaching or being reached from the
// internet through internet gateways and egress-only internet gateways. If a
// workgroup is in an account with VPC BPA turned on, the following capabilities
// are blocked:
//
// - Creating a public access workgroup
//
// - Modifying a private workgroup to public
//
// - Adding a subnet with VPC BPA turned on to the workgroup when the workgroup
// is public
//
// For more information about VPC BPA, see [Block public access to VPCs and subnets] in the Amazon VPC User Guide.
//
// [Block public access to VPCs and subnets]: https://docs.aws.amazon.com/vpc/latest/userguide/security-vpc-bpa.html
func redshiftserverless_UpdateWorkgroup(cfg aws.Config, client *redshiftserverless.Client) {
	input := &redshiftserverless.UpdateWorkgroupInput{
		// WorkgroupName: *string, // Required
	}

	if len(_redshiftserverlessWorkgroupName) > 0 {
		input.WorkgroupName = aws.String(_redshiftserverlessWorkgroupName)
	}
	if len(_redshiftserverlessBaseCapacity) > 0 {
		if err := assignInputField(input, "BaseCapacity", _redshiftserverlessBaseCapacity); err != nil {
			log.Errorf("invalid --base-capacity: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessConfigParameters) > 0 {
		if err := assignInputField(input, "ConfigParameters", _redshiftserverlessConfigParameters); err != nil {
			log.Errorf("invalid --config-parameters: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessEnhancedVpcRouting) > 0 {
		if err := assignInputField(input, "EnhancedVpcRouting", _redshiftserverlessEnhancedVpcRouting); err != nil {
			log.Errorf("invalid --enhanced-vpc-routing: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessExtraComputeForAutomaticOptimization) > 0 {
		if err := assignInputField(input, "ExtraComputeForAutomaticOptimization", _redshiftserverlessExtraComputeForAutomaticOptimization); err != nil {
			log.Errorf("invalid --extra-compute-for-automatic-optimization: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessIpAddressType) > 0 {
		input.IpAddressType = aws.String(_redshiftserverlessIpAddressType)
	}
	if len(_redshiftserverlessMaxCapacity) > 0 {
		if err := assignInputField(input, "MaxCapacity", _redshiftserverlessMaxCapacity); err != nil {
			log.Errorf("invalid --max-capacity: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPort) > 0 {
		if err := assignInputField(input, "Port", _redshiftserverlessPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPricePerformanceTarget) > 0 {
		if err := assignInputField(input, "PricePerformanceTarget", _redshiftserverlessPricePerformanceTarget); err != nil {
			log.Errorf("invalid --price-performance-target: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _redshiftserverlessPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_redshiftserverlessSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _redshiftserverlessSecurityGroupIds...)
	}
	if len(_redshiftserverlessSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _redshiftserverlessSubnetIds...)
	}
	if len(_redshiftserverlessTrackName) > 0 {
		input.TrackName = aws.String(_redshiftserverlessTrackName)
	}

	if resp, err := client.UpdateWorkgroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_redshiftserverlessCmd)
	_redshiftserverlessCmd.Flags().SortFlags = false

	_redshiftserverlessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_redshiftserverlessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_redshiftserverlessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessActivateCaseSensitiveIdentifier, "activate-case-sensitive-identifier", "", "", "Activate Case Sensitive Identifier")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessAdminPasswordSecretKmsKeyId, "admin-password-secret-kms-key-id", "", "", "Admin Password Secret KMS Key ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessAdminUserPassword, "admin-user-password", "", "", "Admin User Password")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessAdminUsername, "admin-username", "", "", "Admin Username")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessAmount, "amount", "", "", "Amount")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessBaseCapacity, "base-capacity", "", "", "Base Capacity")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessBreachAction, "breach-action", "", "", "Breach Action")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessCapacity, "capacity", "", "", "Capacity")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessCatalogName, "catalog-name", "", "", "Catalog Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessClientToken, "client-token", "", "", "Client Token")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessConfigParameters, "config-parameters", "", "", "Config Parameters")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessCustomDomainCertificateArn, "custom-domain-certificate-arn", "", "", "Custom Domain Certificate ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessCustomDomainName, "custom-domain-name", "", "", "Custom Domain Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessDbName, "db-name", "", "", "DB Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessDefaultIamRoleArn, "default-iam-role-arn", "", "", "Default IAM Role ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessDestinationKmsKeyId, "destination-kms-key-id", "", "", "Destination KMS Key ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessDestinationRegion, "destination-region", "", "", "Destination Region")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessDryRun, "dry-run", "", "", "Dry Run")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessEnabled, "enabled", "", "", "Enabled")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessEndTime, "end-time", "", "", "End Time")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessEndpointName, "endpoint-name", "", "", "Endpoint Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessEnhancedVpcRouting, "enhanced-vpc-routing", "", "", "Enhanced VPC Routing")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessExtraComputeForAutomaticOptimization, "extra-compute-for-automatic-optimization", "", "", "Extra Compute For Automatic Optimization")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessFinalSnapshotName, "final-snapshot-name", "", "", "Final Snapshot Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessFinalSnapshotRetentionPeriod, "final-snapshot-retention-period", "", "", "Final Snapshot Retention Period")
	_redshiftserverlessCmd.Flags().StringSliceVarP(&_redshiftserverlessIamRoles, "iam-roles", "", nil, "IAM Roles")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessLakehouseIdcApplicationArn, "lakehouse-idc-application-arn", "", "", "Lakehouse Idc Application ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessLakehouseIdcRegistration, "lakehouse-idc-registration", "", "", "Lakehouse Idc Registration")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessLakehouseRegistration, "lakehouse-registration", "", "", "Lakehouse Registration")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessLogExports, "log-exports", "", "", "Log Exports")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessManageAdminPassword, "manage-admin-password", "", "", "Manage Admin Password")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessMaxCapacity, "max-capacity", "", "", "Max Capacity")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessMaxResults, "max-results", "", "", "Max Results")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessNamespaceArn, "namespace-arn", "", "", "Namespace ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessNamespaceName, "namespace-name", "", "", "Namespace Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessNewTableName, "new-table-name", "", "", "New Table Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessNextToken, "next-token", "", "", "Next Token")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessOfferingId, "offering-id", "", "", "Offering ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessOwnerAccount, "owner-account", "", "", "Owner Account")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessPeriod, "period", "", "", "Period")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessPolicy, "policy", "", "", "Policy")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessPort, "port", "", "", "Port")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessPricePerformanceTarget, "price-performance-target", "", "", "Price Performance Target")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessPubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessRecoveryPointId, "recovery-point-id", "", "", "Recovery Point ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessRedshiftIdcApplicationArn, "redshift-idc-application-arn", "", "", "Redshift Idc Application ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessReservationId, "reservation-id", "", "", "Reservation ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessResourceArn, "resource-arn", "", "", "Resource ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessRetentionPeriod, "retention-period", "", "", "Retention Period")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessRoleArn, "role-arn", "", "", "Role ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSchedule, "schedule", "", "", "Schedule")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessScheduledActionDescription, "scheduled-action-description", "", "", "Scheduled Action Description")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessScheduledActionName, "scheduled-action-name", "", "", "Scheduled Action Name")
	_redshiftserverlessCmd.Flags().StringSliceVarP(&_redshiftserverlessSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSnapshotArn, "snapshot-arn", "", "", "Snapshot ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSnapshotCopyConfigurationId, "snapshot-copy-configuration-id", "", "", "Snapshot Copy Configuration ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSnapshotName, "snapshot-name", "", "", "Snapshot Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSnapshotRetentionPeriod, "snapshot-retention-period", "", "", "Snapshot Retention Period")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSourceArn, "source-arn", "", "", "Source ARN")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSourceDatabaseName, "source-database-name", "", "", "Source Database Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSourceSchemaName, "source-schema-name", "", "", "Source Schema Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessSourceTableName, "source-table-name", "", "", "Source Table Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessStartTime, "start-time", "", "", "Start Time")
	_redshiftserverlessCmd.Flags().StringSliceVarP(&_redshiftserverlessSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessTableRestoreRequestId, "table-restore-request-id", "", "", "Table Restore Request ID")
	_redshiftserverlessCmd.Flags().StringSliceVarP(&_redshiftserverlessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessTags, "tags", "", "", "Tags")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessTargetAction, "target-action", "", "", "Target Action")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessTargetDatabaseName, "target-database-name", "", "", "Target Database Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessTargetSchemaName, "target-schema-name", "", "", "Target Schema Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessTrackName, "track-name", "", "", "Track Name")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessUsageLimitId, "usage-limit-id", "", "", "Usage Limit ID")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessUsageType, "usage-type", "", "", "Usage Type")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessVpcId, "vpc-id", "", "", "VPC ID")
	_redshiftserverlessCmd.Flags().StringSliceVarP(&_redshiftserverlessVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")
	_redshiftserverlessCmd.Flags().StringVarP(&_redshiftserverlessWorkgroupName, "workgroup-name", "", "", "Workgroup Name")
	_redshiftserverlessCmd.Flags().StringSliceVarP(&_redshiftserverlessWorkgroupNames, "workgroup-names", "", nil, "Workgroup Names")

	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessConvertRecoveryPointToSnapshot, "convert-recovery-point-to-snapshot", "", false, "Convert Recovery Point To Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateCustomDomainAssociation, "create-custom-domain-association", "", false, "Create Custom Domain Association")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateEndpointAccess, "create-endpoint-access", "", false, "Create Endpoint Access")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateNamespace, "create-namespace", "", false, "Create Namespace")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateReservation, "create-reservation", "", false, "Create Reservation")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateScheduledAction, "create-scheduled-action", "", false, "Create Scheduled Action")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateSnapshot, "create-snapshot", "", false, "Create Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateSnapshotCopyConfiguration, "create-snapshot-copy-configuration", "", false, "Create Snapshot Copy Configuration")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateUsageLimit, "create-usage-limit", "", false, "Create Usage Limit")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessCreateWorkgroup, "create-workgroup", "", false, "Create Workgroup")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteCustomDomainAssociation, "delete-custom-domain-association", "", false, "Delete Custom Domain Association")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteEndpointAccess, "delete-endpoint-access", "", false, "Delete Endpoint Access")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteNamespace, "delete-namespace", "", false, "Delete Namespace")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteScheduledAction, "delete-scheduled-action", "", false, "Delete Scheduled Action")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteSnapshot, "delete-snapshot", "", false, "Delete Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteSnapshotCopyConfiguration, "delete-snapshot-copy-configuration", "", false, "Delete Snapshot Copy Configuration")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteUsageLimit, "delete-usage-limit", "", false, "Delete Usage Limit")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessDeleteWorkgroup, "delete-workgroup", "", false, "Delete Workgroup")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetCredentials, "get-credentials", "", false, "Get Credentials")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetCustomDomainAssociation, "get-custom-domain-association", "", false, "Get Custom Domain Association")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetEndpointAccess, "get-endpoint-access", "", false, "Get Endpoint Access")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetIdentityCenterAuthToken, "get-identity-center-auth-token", "", false, "Get Identity Center Auth Token")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetNamespace, "get-namespace", "", false, "Get Namespace")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetRecoveryPoint, "get-recovery-point", "", false, "Get Recovery Point")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetReservation, "get-reservation", "", false, "Get Reservation")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetReservationOffering, "get-reservation-offering", "", false, "Get Reservation Offering")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetScheduledAction, "get-scheduled-action", "", false, "Get Scheduled Action")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetSnapshot, "get-snapshot", "", false, "Get Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetTableRestoreStatus, "get-table-restore-status", "", false, "Get Table Restore Status")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetTrack, "get-track", "", false, "Get Track")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetUsageLimit, "get-usage-limit", "", false, "Get Usage Limit")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessGetWorkgroup, "get-workgroup", "", false, "Get Workgroup")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListCustomDomainAssociations, "list-custom-domain-associations", "", false, "List Custom Domain Associations")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListEndpointAccess, "list-endpoint-access", "", false, "List Endpoint Access")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListManagedWorkgroups, "list-managed-workgroups", "", false, "List Managed Workgroups")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListNamespaces, "list-namespaces", "", false, "List Namespaces")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListRecoveryPoints, "list-recovery-points", "", false, "List Recovery Points")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListReservationOfferings, "list-reservation-offerings", "", false, "List Reservation Offerings")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListReservations, "list-reservations", "", false, "List Reservations")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListScheduledActions, "list-scheduled-actions", "", false, "List Scheduled Actions")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListSnapshotCopyConfigurations, "list-snapshot-copy-configurations", "", false, "List Snapshot Copy Configurations")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListSnapshots, "list-snapshots", "", false, "List Snapshots")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListTableRestoreStatus, "list-table-restore-status", "", false, "List Table Restore Status")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListTracks, "list-tracks", "", false, "List Tracks")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListUsageLimits, "list-usage-limits", "", false, "List Usage Limits")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessListWorkgroups, "list-workgroups", "", false, "List Workgroups")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessPutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessRestoreFromRecoveryPoint, "restore-from-recovery-point", "", false, "Restore From Recovery Point")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessRestoreFromSnapshot, "restore-from-snapshot", "", false, "Restore From Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessRestoreTableFromRecoveryPoint, "restore-table-from-recovery-point", "", false, "Restore Table From Recovery Point")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessRestoreTableFromSnapshot, "restore-table-from-snapshot", "", false, "Restore Table From Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessTagResource, "tag-resource", "", false, "Tag Resource")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUntagResource, "untag-resource", "", false, "Untag Resource")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateCustomDomainAssociation, "update-custom-domain-association", "", false, "Update Custom Domain Association")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateEndpointAccess, "update-endpoint-access", "", false, "Update Endpoint Access")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateLakehouseConfiguration, "update-lakehouse-configuration", "", false, "Update Lakehouse Configuration")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateNamespace, "update-namespace", "", false, "Update Namespace")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateScheduledAction, "update-scheduled-action", "", false, "Update Scheduled Action")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateSnapshot, "update-snapshot", "", false, "Update Snapshot")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateSnapshotCopyConfiguration, "update-snapshot-copy-configuration", "", false, "Update Snapshot Copy Configuration")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateUsageLimit, "update-usage-limit", "", false, "Update Usage Limit")
	_redshiftserverlessCmd.Flags().BoolVarP(&_redshiftserverlessUpdateWorkgroup, "update-workgroup", "", false, "Update Workgroup")

}
