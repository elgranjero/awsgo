package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/docdb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// docdbCmd represents the docdb command
var _docdbCmd = &cobra.Command{
	Use:   "docdb",
	Short: "AWS docdb CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := docdb.NewFromConfig(cfg)
		if _docdbAddSourceIdentifierToSubscription {
			docdb_AddSourceIdentifierToSubscription(cfg, client)
			return
		}
		if _docdbAddTagsToResource {
			docdb_AddTagsToResource(cfg, client)
			return
		}
		if _docdbApplyPendingMaintenanceAction {
			docdb_ApplyPendingMaintenanceAction(cfg, client)
			return
		}
		if _docdbCopyDBClusterParameterGroup {
			docdb_CopyDBClusterParameterGroup(cfg, client)
			return
		}
		if _docdbCopyDBClusterSnapshot {
			docdb_CopyDBClusterSnapshot(cfg, client)
			return
		}
		if _docdbCreateDBCluster {
			docdb_CreateDBCluster(cfg, client)
			return
		}
		if _docdbCreateDBClusterParameterGroup {
			docdb_CreateDBClusterParameterGroup(cfg, client)
			return
		}
		if _docdbCreateDBClusterSnapshot {
			docdb_CreateDBClusterSnapshot(cfg, client)
			return
		}
		if _docdbCreateDBInstance {
			docdb_CreateDBInstance(cfg, client)
			return
		}
		if _docdbCreateDBSubnetGroup {
			docdb_CreateDBSubnetGroup(cfg, client)
			return
		}
		if _docdbCreateEventSubscription {
			docdb_CreateEventSubscription(cfg, client)
			return
		}
		if _docdbCreateGlobalCluster {
			docdb_CreateGlobalCluster(cfg, client)
			return
		}
		if _docdbDeleteDBCluster {
			docdb_DeleteDBCluster(cfg, client)
			return
		}
		if _docdbDeleteDBClusterParameterGroup {
			docdb_DeleteDBClusterParameterGroup(cfg, client)
			return
		}
		if _docdbDeleteDBClusterSnapshot {
			docdb_DeleteDBClusterSnapshot(cfg, client)
			return
		}
		if _docdbDeleteDBInstance {
			docdb_DeleteDBInstance(cfg, client)
			return
		}
		if _docdbDeleteDBSubnetGroup {
			docdb_DeleteDBSubnetGroup(cfg, client)
			return
		}
		if _docdbDeleteEventSubscription {
			docdb_DeleteEventSubscription(cfg, client)
			return
		}
		if _docdbDeleteGlobalCluster {
			docdb_DeleteGlobalCluster(cfg, client)
			return
		}
		if _docdbDescribeCertificates {
			docdb_DescribeCertificates(cfg, client)
			return
		}
		if _docdbDescribeDBClusterParameterGroups {
			docdb_DescribeDBClusterParameterGroups(cfg, client)
			return
		}
		if _docdbDescribeDBClusterParameters {
			docdb_DescribeDBClusterParameters(cfg, client)
			return
		}
		if _docdbDescribeDBClusterSnapshotAttributes {
			docdb_DescribeDBClusterSnapshotAttributes(cfg, client)
			return
		}
		if _docdbDescribeDBClusterSnapshots {
			docdb_DescribeDBClusterSnapshots(cfg, client)
			return
		}
		if _docdbDescribeDBClusters {
			docdb_DescribeDBClusters(cfg, client)
			return
		}
		if _docdbDescribeDBEngineVersions {
			docdb_DescribeDBEngineVersions(cfg, client)
			return
		}
		if _docdbDescribeDBInstances {
			docdb_DescribeDBInstances(cfg, client)
			return
		}
		if _docdbDescribeDBSubnetGroups {
			docdb_DescribeDBSubnetGroups(cfg, client)
			return
		}
		if _docdbDescribeEngineDefaultClusterParameters {
			docdb_DescribeEngineDefaultClusterParameters(cfg, client)
			return
		}
		if _docdbDescribeEventCategories {
			docdb_DescribeEventCategories(cfg, client)
			return
		}
		if _docdbDescribeEventSubscriptions {
			docdb_DescribeEventSubscriptions(cfg, client)
			return
		}
		if _docdbDescribeEvents {
			docdb_DescribeEvents(cfg, client)
			return
		}
		if _docdbDescribeGlobalClusters {
			docdb_DescribeGlobalClusters(cfg, client)
			return
		}
		if _docdbDescribeOrderableDBInstanceOptions {
			docdb_DescribeOrderableDBInstanceOptions(cfg, client)
			return
		}
		if _docdbDescribePendingMaintenanceActions {
			docdb_DescribePendingMaintenanceActions(cfg, client)
			return
		}
		if _docdbFailoverDBCluster {
			docdb_FailoverDBCluster(cfg, client)
			return
		}
		if _docdbFailoverGlobalCluster {
			docdb_FailoverGlobalCluster(cfg, client)
			return
		}
		if _docdbListTagsForResource {
			docdb_ListTagsForResource(cfg, client)
			return
		}
		if _docdbModifyDBCluster {
			docdb_ModifyDBCluster(cfg, client)
			return
		}
		if _docdbModifyDBClusterParameterGroup {
			docdb_ModifyDBClusterParameterGroup(cfg, client)
			return
		}
		if _docdbModifyDBClusterSnapshotAttribute {
			docdb_ModifyDBClusterSnapshotAttribute(cfg, client)
			return
		}
		if _docdbModifyDBInstance {
			docdb_ModifyDBInstance(cfg, client)
			return
		}
		if _docdbModifyDBSubnetGroup {
			docdb_ModifyDBSubnetGroup(cfg, client)
			return
		}
		if _docdbModifyEventSubscription {
			docdb_ModifyEventSubscription(cfg, client)
			return
		}
		if _docdbModifyGlobalCluster {
			docdb_ModifyGlobalCluster(cfg, client)
			return
		}
		if _docdbRebootDBInstance {
			docdb_RebootDBInstance(cfg, client)
			return
		}
		if _docdbRemoveFromGlobalCluster {
			docdb_RemoveFromGlobalCluster(cfg, client)
			return
		}
		if _docdbRemoveSourceIdentifierFromSubscription {
			docdb_RemoveSourceIdentifierFromSubscription(cfg, client)
			return
		}
		if _docdbRemoveTagsFromResource {
			docdb_RemoveTagsFromResource(cfg, client)
			return
		}
		if _docdbResetDBClusterParameterGroup {
			docdb_ResetDBClusterParameterGroup(cfg, client)
			return
		}
		if _docdbRestoreDBClusterFromSnapshot {
			docdb_RestoreDBClusterFromSnapshot(cfg, client)
			return
		}
		if _docdbRestoreDBClusterToPointInTime {
			docdb_RestoreDBClusterToPointInTime(cfg, client)
			return
		}
		if _docdbStartDBCluster {
			docdb_StartDBCluster(cfg, client)
			return
		}
		if _docdbStopDBCluster {
			docdb_StopDBCluster(cfg, client)
			return
		}
		if _docdbSwitchoverGlobalCluster {
			docdb_SwitchoverGlobalCluster(cfg, client)
			return
		}

	},
}

var (
	_docdbAddSourceIdentifierToSubscription      bool
	_docdbAddTagsToResource                      bool
	_docdbApplyPendingMaintenanceAction          bool
	_docdbCopyDBClusterParameterGroup            bool
	_docdbCopyDBClusterSnapshot                  bool
	_docdbCreateDBCluster                        bool
	_docdbCreateDBClusterParameterGroup          bool
	_docdbCreateDBClusterSnapshot                bool
	_docdbCreateDBInstance                       bool
	_docdbCreateDBSubnetGroup                    bool
	_docdbCreateEventSubscription                bool
	_docdbCreateGlobalCluster                    bool
	_docdbDeleteDBCluster                        bool
	_docdbDeleteDBClusterParameterGroup          bool
	_docdbDeleteDBClusterSnapshot                bool
	_docdbDeleteDBInstance                       bool
	_docdbDeleteDBSubnetGroup                    bool
	_docdbDeleteEventSubscription                bool
	_docdbDeleteGlobalCluster                    bool
	_docdbDescribeCertificates                   bool
	_docdbDescribeDBClusterParameterGroups       bool
	_docdbDescribeDBClusterParameters            bool
	_docdbDescribeDBClusterSnapshotAttributes    bool
	_docdbDescribeDBClusterSnapshots             bool
	_docdbDescribeDBClusters                     bool
	_docdbDescribeDBEngineVersions               bool
	_docdbDescribeDBInstances                    bool
	_docdbDescribeDBSubnetGroups                 bool
	_docdbDescribeEngineDefaultClusterParameters bool
	_docdbDescribeEventCategories                bool
	_docdbDescribeEventSubscriptions             bool
	_docdbDescribeEvents                         bool
	_docdbDescribeGlobalClusters                 bool
	_docdbDescribeOrderableDBInstanceOptions     bool
	_docdbDescribePendingMaintenanceActions      bool
	_docdbFailoverDBCluster                      bool
	_docdbFailoverGlobalCluster                  bool
	_docdbListTagsForResource                    bool
	_docdbModifyDBCluster                        bool
	_docdbModifyDBClusterParameterGroup          bool
	_docdbModifyDBClusterSnapshotAttribute       bool
	_docdbModifyDBInstance                       bool
	_docdbModifyDBSubnetGroup                    bool
	_docdbModifyEventSubscription                bool
	_docdbModifyGlobalCluster                    bool
	_docdbRebootDBInstance                       bool
	_docdbRemoveFromGlobalCluster                bool
	_docdbRemoveSourceIdentifierFromSubscription bool
	_docdbRemoveTagsFromResource                 bool
	_docdbResetDBClusterParameterGroup           bool
	_docdbRestoreDBClusterFromSnapshot           bool
	_docdbRestoreDBClusterToPointInTime          bool
	_docdbStartDBCluster                         bool
	_docdbStopDBCluster                          bool
	_docdbSwitchoverGlobalCluster                bool

	_docdbAllowDataLoss                            string
	_docdbAllowMajorVersionUpgrade                 string
	_docdbApplyAction                              string
	_docdbApplyImmediately                         string
	_docdbAttributeName                            string
	_docdbAutoMinorVersionUpgrade                  string
	_docdbAvailabilityZone                         string
	_docdbAvailabilityZones                        []string
	_docdbBackupRetentionPeriod                    string
	_docdbCACertificateIdentifier                  string
	_docdbCertificateIdentifier                    string
	_docdbCertificateRotationRestart               string
	_docdbCloudwatchLogsExportConfiguration        string
	_docdbCopyTags                                 string
	_docdbCopyTagsToSnapshot                       string
	_docdbDatabaseName                             string
	_docdbDBClusterIdentifier                      string
	_docdbDBClusterParameterGroupName              string
	_docdbDBClusterSnapshotIdentifier              string
	_docdbDBInstanceClass                          string
	_docdbDBInstanceIdentifier                     string
	_docdbDBParameterGroupFamily                   string
	_docdbDBSubnetGroupDescription                 string
	_docdbDBSubnetGroupName                        string
	_docdbDefaultOnly                              string
	_docdbDeletionProtection                       string
	_docdbDescription                              string
	_docdbDuration                                 string
	_docdbEnableCloudwatchLogsExports              []string
	_docdbEnablePerformanceInsights                string
	_docdbEnabled                                  string
	_docdbEndTime                                  string
	_docdbEngine                                   string
	_docdbEngineVersion                            string
	_docdbEventCategories                          []string
	_docdbFilters                                  string
	_docdbFinalDBSnapshotIdentifier                string
	_docdbForceFailover                            string
	_docdbGlobalClusterIdentifier                  string
	_docdbIncludePublic                            string
	_docdbIncludeShared                            string
	_docdbKmsKeyId                                 string
	_docdbLicenseModel                             string
	_docdbListSupportedCharacterSets               string
	_docdbListSupportedTimezones                   string
	_docdbManageMasterUserPassword                 string
	_docdbMarker                                   string
	_docdbMasterUserPassword                       string
	_docdbMasterUserSecretKmsKeyId                 string
	_docdbMasterUsername                           string
	_docdbMaxRecords                               string
	_docdbNetworkType                              string
	_docdbNewDBClusterIdentifier                   string
	_docdbNewDBInstanceIdentifier                  string
	_docdbNewGlobalClusterIdentifier               string
	_docdbOptInType                                string
	_docdbParameters                               string
	_docdbPerformanceInsightsKMSKeyId              string
	_docdbPort                                     string
	_docdbPreSignedUrl                             string
	_docdbPreferredBackupWindow                    string
	_docdbPreferredMaintenanceWindow               string
	_docdbPromotionTier                            string
	_docdbResetAllParameters                       string
	_docdbResourceIdentifier                       string
	_docdbResourceName                             string
	_docdbRestoreToTime                            string
	_docdbRestoreType                              string
	_docdbRotateMasterUserPassword                 string
	_docdbServerlessV2ScalingConfiguration         string
	_docdbSkipFinalSnapshot                        string
	_docdbSnapshotIdentifier                       string
	_docdbSnapshotType                             string
	_docdbSnsTopicArn                              string
	_docdbSource                                   string
	_docdbSourceDBClusterIdentifier                string
	_docdbSourceDBClusterParameterGroupIdentifier  string
	_docdbSourceDBClusterSnapshotIdentifier        string
	_docdbSourceIdentifier                         string
	_docdbSourceIds                                []string
	_docdbSourceRegion                             string
	_docdbSourceType                               string
	_docdbStartTime                                string
	_docdbStorageEncrypted                         string
	_docdbStorageType                              string
	_docdbSubnetIds                                []string
	_docdbSubscriptionName                         string
	_docdbSwitchover                               string
	_docdbTagKeys                                  []string
	_docdbTags                                     string
	_docdbTargetDbClusterIdentifier                string
	_docdbTargetDBClusterParameterGroupDescription string
	_docdbTargetDBClusterParameterGroupIdentifier  string
	_docdbTargetDBClusterSnapshotIdentifier        string
	_docdbTargetDBInstanceIdentifier               string
	_docdbUseLatestRestorableTime                  string
	_docdbValuesToAdd                              []string
	_docdbValuesToRemove                           []string
	_docdbVpc                                      string
	_docdbVpcSecurityGroupIds                      []string
)

// Adds a source identifier to an existing event notification subscription.
func docdb_AddSourceIdentifierToSubscription(cfg aws.Config, client *docdb.Client) {
	input := &docdb.AddSourceIdentifierToSubscriptionInput{
		// SourceIdentifier: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_docdbSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_docdbSourceIdentifier)
	}
	if len(_docdbSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_docdbSubscriptionName)
	}

	if resp, err := client.AddSourceIdentifierToSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to an Amazon DocumentDB resource. You can use these tags
// with cost allocation reporting to track costs that are associated with Amazon
// DocumentDB resources or in a Condition statement in an Identity and Access
// Management (IAM) policy for Amazon DocumentDB.
func docdb_AddTagsToResource(cfg aws.Config, client *docdb.Client) {
	input := &docdb.AddTagsToResourceInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_docdbResourceName) > 0 {
		input.ResourceName = aws.String(_docdbResourceName)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTagsToResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies a pending maintenance action to a resource (for example, to an Amazon
// DocumentDB instance).
func docdb_ApplyPendingMaintenanceAction(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ApplyPendingMaintenanceActionInput{
		// ApplyAction: *string, // Required
		// OptInType: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_docdbApplyAction) > 0 {
		input.ApplyAction = aws.String(_docdbApplyAction)
	}
	if len(_docdbOptInType) > 0 {
		input.OptInType = aws.String(_docdbOptInType)
	}
	if len(_docdbResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_docdbResourceIdentifier)
	}

	if resp, err := client.ApplyPendingMaintenanceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified cluster parameter group.
func docdb_CopyDBClusterParameterGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CopyDBClusterParameterGroupInput{
		// SourceDBClusterParameterGroupIdentifier: *string, // Required
		// TargetDBClusterParameterGroupDescription: *string, // Required
		// TargetDBClusterParameterGroupIdentifier: *string, // Required
	}

	if len(_docdbSourceDBClusterParameterGroupIdentifier) > 0 {
		input.SourceDBClusterParameterGroupIdentifier = aws.String(_docdbSourceDBClusterParameterGroupIdentifier)
	}
	if len(_docdbTargetDBClusterParameterGroupDescription) > 0 {
		input.TargetDBClusterParameterGroupDescription = aws.String(_docdbTargetDBClusterParameterGroupDescription)
	}
	if len(_docdbTargetDBClusterParameterGroupIdentifier) > 0 {
		input.TargetDBClusterParameterGroupIdentifier = aws.String(_docdbTargetDBClusterParameterGroupIdentifier)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies a snapshot of a cluster.
// To copy a cluster snapshot from a shared manual cluster snapshot,
// SourceDBClusterSnapshotIdentifier must be the Amazon Resource Name (ARN) of the
// shared cluster snapshot. You can only copy a shared DB cluster snapshot, whether
// encrypted or not, in the same Amazon Web Services Region.
//
// To cancel the copy operation after it is in progress, delete the target cluster
// snapshot identified by TargetDBClusterSnapshotIdentifier while that cluster
// snapshot is in the copying status.
func docdb_CopyDBClusterSnapshot(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CopyDBClusterSnapshotInput{
		// SourceDBClusterSnapshotIdentifier: *string, // Required
		// TargetDBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_docdbSourceDBClusterSnapshotIdentifier) > 0 {
		input.SourceDBClusterSnapshotIdentifier = aws.String(_docdbSourceDBClusterSnapshotIdentifier)
	}
	if len(_docdbTargetDBClusterSnapshotIdentifier) > 0 {
		input.TargetDBClusterSnapshotIdentifier = aws.String(_docdbTargetDBClusterSnapshotIdentifier)
	}
	if len(_docdbCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _docdbCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_docdbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbKmsKeyId)
	}
	if len(_docdbPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_docdbPreSignedUrl)
	}
	if len(_docdbSourceRegion) > 0 {
		input.SourceRegion = aws.String(_docdbSourceRegion)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon DocumentDB cluster.
func docdb_CreateDBCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbEngine) > 0 {
		input.Engine = aws.String(_docdbEngine)
	}
	if len(_docdbAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _docdbAvailabilityZones...)
	}
	if len(_docdbBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _docdbBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}
	if len(_docdbDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _docdbDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_docdbEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _docdbEnableCloudwatchLogsExports...)
	}
	if len(_docdbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_docdbEngineVersion)
	}
	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}
	if len(_docdbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbKmsKeyId)
	}
	if len(_docdbManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _docdbManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_docdbMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_docdbMasterUserPassword)
	}
	if len(_docdbMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_docdbMasterUserSecretKmsKeyId)
	}
	if len(_docdbMasterUsername) > 0 {
		input.MasterUsername = aws.String(_docdbMasterUsername)
	}
	if len(_docdbNetworkType) > 0 {
		input.NetworkType = aws.String(_docdbNetworkType)
	}
	if len(_docdbPort) > 0 {
		if err := assignInputField(input, "Port", _docdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_docdbPreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_docdbPreSignedUrl)
	}
	if len(_docdbPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_docdbPreferredBackupWindow)
	}
	if len(_docdbPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_docdbPreferredMaintenanceWindow)
	}
	if len(_docdbServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _docdbServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_docdbSourceRegion) > 0 {
		input.SourceRegion = aws.String(_docdbSourceRegion)
	}
	if len(_docdbStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _docdbStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_docdbStorageType) > 0 {
		input.StorageType = aws.String(_docdbStorageType)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_docdbVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cluster parameter group.
// Parameters in a cluster parameter group apply to all of the instances in a
// cluster.
//
// A cluster parameter group is initially created with the default parameters for
// the database engine used by instances in the cluster. In Amazon DocumentDB, you
// cannot make modifications directly to the default.docdb3.6 cluster parameter
// group. If your Amazon DocumentDB cluster is using the default cluster parameter
// group and you want to modify a value in it, you must first [create a new parameter group]or [copy an existing parameter group], modify it, and
// then apply the modified parameter group to your cluster. For the new cluster
// parameter group and associated settings to take effect, you must then reboot the
// instances in the cluster without failover. For more information, see [Modifying Amazon DocumentDB Cluster Parameter Groups].
//
// [create a new parameter group]: https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-create.html
// [Modifying Amazon DocumentDB Cluster Parameter Groups]: https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-modify.html
// [copy an existing parameter group]: https://docs.aws.amazon.com/documentdb/latest/developerguide/cluster_parameter_group-copy.html
func docdb_CreateDBClusterParameterGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
		// DBParameterGroupFamily: *string, // Required
		// Description: *string, // Required
	}

	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_docdbDBParameterGroupFamily)
	}
	if len(_docdbDescription) > 0 {
		input.Description = aws.String(_docdbDescription)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a snapshot of a cluster.
func docdb_CreateDBClusterSnapshot(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateDBClusterSnapshotInput{
		// DBClusterIdentifier: *string, // Required
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_docdbDBClusterSnapshotIdentifier)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new instance.
func docdb_CreateDBInstance(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateDBInstanceInput{
		// DBClusterIdentifier: *string, // Required
		// DBInstanceClass: *string, // Required
		// DBInstanceIdentifier: *string, // Required
		// Engine: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_docdbDBInstanceClass)
	}
	if len(_docdbDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_docdbDBInstanceIdentifier)
	}
	if len(_docdbEngine) > 0 {
		input.Engine = aws.String(_docdbEngine)
	}
	if len(_docdbAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _docdbAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_docdbAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_docdbAvailabilityZone)
	}
	if len(_docdbCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_docdbCACertificateIdentifier)
	}
	if len(_docdbCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _docdbCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_docdbEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _docdbEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_docdbPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_docdbPerformanceInsightsKMSKeyId)
	}
	if len(_docdbPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_docdbPreferredMaintenanceWindow)
	}
	if len(_docdbPromotionTier) > 0 {
		if err := assignInputField(input, "PromotionTier", _docdbPromotionTier); err != nil {
			log.Errorf("invalid --promotion-tier: %s", err.Error())
			return
		}
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new subnet group. subnet groups must contain at least one subnet in
// at least two Availability Zones in the Amazon Web Services Region.
func docdb_CreateDBSubnetGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateDBSubnetGroupInput{
		// DBSubnetGroupDescription: *string, // Required
		// DBSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_docdbDBSubnetGroupDescription) > 0 {
		input.DBSubnetGroupDescription = aws.String(_docdbDBSubnetGroupDescription)
	}
	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}
	if len(_docdbSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _docdbSubnetIds...)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DocumentDB event notification subscription. This action
// requires a topic Amazon Resource Name (ARN) created by using the Amazon
// DocumentDB console, the Amazon SNS console, or the Amazon SNS API. To obtain an
// ARN with Amazon SNS, you must create a topic in Amazon SNS and subscribe to the
// topic. The ARN is displayed in the Amazon SNS console.
//
// You can specify the type of source ( SourceType ) that you want to be notified
// of. You can also provide a list of Amazon DocumentDB sources ( SourceIds ) that
// trigger the events, and you can provide a list of event categories (
// EventCategories ) for events that you want to be notified of. For example, you
// can specify SourceType = db-instance , SourceIds = mydbinstance1, mydbinstance2
// and EventCategories = Availability, Backup .
//
// If you specify both the SourceType and SourceIds (such as SourceType =
// db-instance and SourceIdentifier = myDBInstance1 ), you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify a SourceIdentifier , you receive notice of the events for that
// source type for all your Amazon DocumentDB sources. If you do not specify either
// the SourceType or the SourceIdentifier , you are notified of events generated
// from all Amazon DocumentDB sources belonging to your customer account.
func docdb_CreateEventSubscription(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateEventSubscriptionInput{
		// SnsTopicArn: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_docdbSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_docdbSnsTopicArn)
	}
	if len(_docdbSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_docdbSubscriptionName)
	}
	if len(_docdbEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _docdbEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_docdbEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _docdbEventCategories...)
	}
	if len(_docdbSourceIds) > 0 {
		input.SourceIds = append([]string(nil), _docdbSourceIds...)
	}
	if len(_docdbSourceType) > 0 {
		input.SourceType = aws.String(_docdbSourceType)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon DocumentDB global cluster that can span multiple multiple
// Amazon Web Services Regions. The global cluster contains one primary cluster
// with read-write capability, and up-to 10 read-only secondary clusters. Global
// clusters uses storage-based fast replication across regions with latencies less
// than one second, using dedicated infrastructure with no impact to your
// workload’s performance.
//
// You can create a global cluster that is initially empty, and then add a primary
// and a secondary to it. Or you can specify an existing cluster during the create
// operation, and this cluster becomes the primary of the global cluster.
//
// This action only applies to Amazon DocumentDB clusters.
func docdb_CreateGlobalCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.CreateGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}
	if len(_docdbDatabaseName) > 0 {
		input.DatabaseName = aws.String(_docdbDatabaseName)
	}
	if len(_docdbDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _docdbDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_docdbEngine) > 0 {
		input.Engine = aws.String(_docdbEngine)
	}
	if len(_docdbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_docdbEngineVersion)
	}
	if len(_docdbSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_docdbSourceDBClusterIdentifier)
	}
	if len(_docdbStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _docdbStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously provisioned cluster. When you delete a cluster, all
// automated backups for that cluster are deleted and can't be recovered. Manual DB
// cluster snapshots of the specified cluster are not deleted.
func docdb_DeleteDBCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbFinalDBSnapshotIdentifier) > 0 {
		input.FinalDBSnapshotIdentifier = aws.String(_docdbFinalDBSnapshotIdentifier)
	}
	if len(_docdbSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _docdbSkipFinalSnapshot); err != nil {
			log.Errorf("invalid --skip-final-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified cluster parameter group. The cluster parameter group to be
// deleted can't be associated with any clusters.
func docdb_DeleteDBClusterParameterGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}

	if resp, err := client.DeleteDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a cluster snapshot. If the snapshot is being copied, the copy operation
// is terminated.
//
// The cluster snapshot must be in the available state to be deleted.
func docdb_DeleteDBClusterSnapshot(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteDBClusterSnapshotInput{
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_docdbDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_docdbDBClusterSnapshotIdentifier)
	}

	if resp, err := client.DeleteDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously provisioned instance.
func docdb_DeleteDBInstance(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_docdbDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_docdbDBInstanceIdentifier)
	}

	if resp, err := client.DeleteDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a subnet group.
// The specified database subnet group must not be associated with any DB
// instances.
func docdb_DeleteDBSubnetGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteDBSubnetGroupInput{
		// DBSubnetGroupName: *string, // Required
	}

	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}

	if resp, err := client.DeleteDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon DocumentDB event notification subscription.
func docdb_DeleteEventSubscription(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_docdbSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_docdbSubscriptionName)
	}

	if resp, err := client.DeleteEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a global cluster. The primary and secondary clusters must already be
// detached or deleted before attempting to delete a global cluster.
//
// This action only applies to Amazon DocumentDB clusters.
func docdb_DeleteGlobalCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DeleteGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}

	if resp, err := client.DeleteGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of certificate authority (CA) certificates provided by Amazon
// DocumentDB for this Amazon Web Services account.
func docdb_DescribeCertificates(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeCertificatesInput{}

	if len(_docdbCertificateIdentifier) > 0 {
		input.CertificateIdentifier = aws.String(_docdbCertificateIdentifier)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeCertificates(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeCertificatesOutput
	p := docdb.NewDescribeCertificatesPaginator(client, input)
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

// Returns a list of DBClusterParameterGroup descriptions. If a
// DBClusterParameterGroupName parameter is specified, the list contains only the
// description of the specified cluster parameter group.
func docdb_DescribeDBClusterParameterGroups(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBClusterParameterGroupsInput{}

	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBClusterParameterGroupsOutput
	p := docdb.NewDescribeDBClusterParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular cluster parameter group.
func docdb_DescribeDBClusterParameters(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBClusterParametersInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_docdbSource) > 0 {
		input.Source = aws.String(_docdbSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBClusterParametersOutput
	p := docdb.NewDescribeDBClusterParametersPaginator(client, input)
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

// Returns a list of cluster snapshot attribute names and values for a manual DB
// cluster snapshot.
//
// When you share snapshots with other Amazon Web Services accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the Amazon Web Services accounts that are authorized to copy or restore
// the manual cluster snapshot. If all is included in the list of values for the
// restore attribute, then the manual cluster snapshot is public and can be copied
// or restored by all Amazon Web Services accounts.
func docdb_DescribeDBClusterSnapshotAttributes(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBClusterSnapshotAttributesInput{
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_docdbDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_docdbDBClusterSnapshotIdentifier)
	}

	if resp, err := client.DescribeDBClusterSnapshotAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about cluster snapshots. This API operation supports
// pagination.
func docdb_DescribeDBClusterSnapshots(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBClusterSnapshotsInput{}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_docdbDBClusterSnapshotIdentifier)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbIncludePublic) > 0 {
		if err := assignInputField(input, "IncludePublic", _docdbIncludePublic); err != nil {
			log.Errorf("invalid --include-public: %s", err.Error())
			return
		}
	}
	if len(_docdbIncludeShared) > 0 {
		if err := assignInputField(input, "IncludeShared", _docdbIncludeShared); err != nil {
			log.Errorf("invalid --include-shared: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_docdbSnapshotType) > 0 {
		input.SnapshotType = aws.String(_docdbSnapshotType)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterSnapshots(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBClusterSnapshotsOutput
	p := docdb.NewDescribeDBClusterSnapshotsPaginator(client, input)
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

// Returns information about provisioned Amazon DocumentDB clusters. This API
// operation supports pagination. For certain management features such as cluster
// and instance lifecycle management, Amazon DocumentDB leverages operational
// technology that is shared with Amazon RDS and Amazon Neptune. Use the
// filterName=engine,Values=docdb filter parameter to return only Amazon DocumentDB
// clusters.
func docdb_DescribeDBClusters(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBClustersInput{}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBClustersOutput
	p := docdb.NewDescribeDBClustersPaginator(client, input)
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

// Returns a list of the available engines.
func docdb_DescribeDBEngineVersions(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBEngineVersionsInput{}

	if len(_docdbDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_docdbDBParameterGroupFamily)
	}
	if len(_docdbDefaultOnly) > 0 {
		if err := assignInputField(input, "DefaultOnly", _docdbDefaultOnly); err != nil {
			log.Errorf("invalid --default-only: %s", err.Error())
			return
		}
	}
	if len(_docdbEngine) > 0 {
		input.Engine = aws.String(_docdbEngine)
	}
	if len(_docdbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_docdbEngineVersion)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbListSupportedCharacterSets) > 0 {
		if err := assignInputField(input, "ListSupportedCharacterSets", _docdbListSupportedCharacterSets); err != nil {
			log.Errorf("invalid --list-supported-character-sets: %s", err.Error())
			return
		}
	}
	if len(_docdbListSupportedTimezones) > 0 {
		if err := assignInputField(input, "ListSupportedTimezones", _docdbListSupportedTimezones); err != nil {
			log.Errorf("invalid --list-supported-timezones: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBEngineVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBEngineVersionsOutput
	p := docdb.NewDescribeDBEngineVersionsPaginator(client, input)
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

// Returns information about provisioned Amazon DocumentDB instances. This API
// supports pagination.
func docdb_DescribeDBInstances(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBInstancesInput{}

	if len(_docdbDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_docdbDBInstanceIdentifier)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBInstancesOutput
	p := docdb.NewDescribeDBInstancesPaginator(client, input)
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

// Returns a list of DBSubnetGroup descriptions. If a DBSubnetGroupName is
// specified, the list will contain only the descriptions of the specified
// DBSubnetGroup .
func docdb_DescribeDBSubnetGroups(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeDBSubnetGroupsInput{}

	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBSubnetGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeDBSubnetGroupsOutput
	p := docdb.NewDescribeDBSubnetGroupsPaginator(client, input)
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

// Returns the default engine and system parameter information for the cluster
// database engine.
func docdb_DescribeEngineDefaultClusterParameters(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeEngineDefaultClusterParametersInput{
		// DBParameterGroupFamily: *string, // Required
	}

	if len(_docdbDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_docdbDBParameterGroupFamily)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeEngineDefaultClusterParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Displays a list of categories for all event source types, or, if specified, for
// a specified source type.
func docdb_DescribeEventCategories(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeEventCategoriesInput{}

	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbSourceType) > 0 {
		input.SourceType = aws.String(_docdbSourceType)
	}

	if resp, err := client.DescribeEventCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the subscription descriptions for a customer account. The description
// for a subscription includes SubscriptionName , SNSTopicARN , CustomerID ,
// SourceType , SourceID , CreationTime , and Status .
//
// If you specify a SubscriptionName , lists the description for that subscription.
func docdb_DescribeEventSubscriptions(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeEventSubscriptionsInput{}

	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_docdbSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_docdbSubscriptionName)
	}

	if disablePaginator() {
		if resp, err := client.DescribeEventSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeEventSubscriptionsOutput
	p := docdb.NewDescribeEventSubscriptionsPaginator(client, input)
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

// Returns events related to instances, security groups, snapshots, and DB
// parameter groups for the past 14 days. You can obtain events specific to a
// particular DB instance, security group, snapshot, or parameter group by
// providing the name as a parameter. By default, the events of the past hour are
// returned.
func docdb_DescribeEvents(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeEventsInput{}

	if len(_docdbDuration) > 0 {
		if err := assignInputField(input, "Duration", _docdbDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_docdbEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _docdbEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_docdbEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _docdbEventCategories...)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_docdbSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_docdbSourceIdentifier)
	}
	if len(_docdbSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _docdbSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_docdbStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _docdbStartTime); err != nil {
			log.Errorf("invalid --start-time: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeEventsOutput
	p := docdb.NewDescribeEventsPaginator(client, input)
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

// Returns information about Amazon DocumentDB global clusters. This API supports
// pagination.
//
// This action only applies to Amazon DocumentDB clusters.
func docdb_DescribeGlobalClusters(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeGlobalClustersInput{}

	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeGlobalClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeGlobalClustersOutput
	p := docdb.NewDescribeGlobalClustersPaginator(client, input)
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

// Returns a list of orderable instance options for the specified engine.
func docdb_DescribeOrderableDBInstanceOptions(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribeOrderableDBInstanceOptionsInput{
		// Engine: *string, // Required
	}

	if len(_docdbEngine) > 0 {
		input.Engine = aws.String(_docdbEngine)
	}
	if len(_docdbDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_docdbDBInstanceClass)
	}
	if len(_docdbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_docdbEngineVersion)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbLicenseModel) > 0 {
		input.LicenseModel = aws.String(_docdbLicenseModel)
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_docdbVpc) > 0 {
		if err := assignInputField(input, "Vpc", _docdbVpc); err != nil {
			log.Errorf("invalid --vpc: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeOrderableDBInstanceOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribeOrderableDBInstanceOptionsOutput
	p := docdb.NewDescribeOrderableDBInstanceOptionsPaginator(client, input)
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

// Returns a list of resources (for example, instances) that have at least one
// pending maintenance action.
func docdb_DescribePendingMaintenanceActions(cfg aws.Config, client *docdb.Client) {
	input := &docdb.DescribePendingMaintenanceActionsInput{}

	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_docdbMarker) > 0 {
		input.Marker = aws.String(_docdbMarker)
	}
	if len(_docdbMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _docdbMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_docdbResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_docdbResourceIdentifier)
	}

	if disablePaginator() {
		if resp, err := client.DescribePendingMaintenanceActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*docdb.DescribePendingMaintenanceActionsOutput
	p := docdb.NewDescribePendingMaintenanceActionsPaginator(client, input)
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

// Forces a failover for a cluster.
// A failover for a cluster promotes one of the Amazon DocumentDB replicas
// (read-only instances) in the cluster to be the primary instance (the cluster
// writer).
//
// If the primary instance fails, Amazon DocumentDB automatically fails over to an
// Amazon DocumentDB replica, if one exists. You can force a failover when you want
// to simulate a failure of a primary instance for testing.
func docdb_FailoverDBCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.FailoverDBClusterInput{}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbTargetDBInstanceIdentifier) > 0 {
		input.TargetDBInstanceIdentifier = aws.String(_docdbTargetDBInstanceIdentifier)
	}

	if resp, err := client.FailoverDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Promotes the specified secondary DB cluster to be the primary DB cluster in the
// global cluster when failing over a global cluster occurs.
//
// Use this operation to respond to an unplanned event, such as a regional
// disaster in the primary region. Failing over can result in a loss of write
// transaction data that wasn't replicated to the chosen secondary before the
// failover event occurred. However, the recovery process that promotes a DB
// instance on the chosen seconday DB cluster to be the primary writer DB instance
// guarantees that the data is in a transactionally consistent state.
func docdb_FailoverGlobalCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.FailoverGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
		// TargetDbClusterIdentifier: *string, // Required
	}

	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}
	if len(_docdbTargetDbClusterIdentifier) > 0 {
		input.TargetDbClusterIdentifier = aws.String(_docdbTargetDbClusterIdentifier)
	}
	if len(_docdbAllowDataLoss) > 0 {
		if err := assignInputField(input, "AllowDataLoss", _docdbAllowDataLoss); err != nil {
			log.Errorf("invalid --allow-data-loss: %s", err.Error())
			return
		}
	}
	if len(_docdbSwitchover) > 0 {
		if err := assignInputField(input, "Switchover", _docdbSwitchover); err != nil {
			log.Errorf("invalid --switchover: %s", err.Error())
			return
		}
	}

	if resp, err := client.FailoverGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all tags on an Amazon DocumentDB resource.
func docdb_ListTagsForResource(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ListTagsForResourceInput{
		// ResourceName: *string, // Required
	}

	if len(_docdbResourceName) > 0 {
		input.ResourceName = aws.String(_docdbResourceName)
	}
	if len(_docdbFilters) > 0 {
		if err := assignInputField(input, "Filters", _docdbFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies a setting for an Amazon DocumentDB cluster. You can change one or more
// database configuration parameters by specifying these parameters and the new
// values in the request.
func docdb_ModifyDBCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _docdbAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_docdbApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _docdbApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_docdbBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _docdbBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_docdbCloudwatchLogsExportConfiguration) > 0 {
		if err := assignInputField(input, "CloudwatchLogsExportConfiguration", _docdbCloudwatchLogsExportConfiguration); err != nil {
			log.Errorf("invalid --cloudwatch-logs-export-configuration: %s", err.Error())
			return
		}
	}
	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _docdbDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_docdbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_docdbEngineVersion)
	}
	if len(_docdbManageMasterUserPassword) > 0 {
		if err := assignInputField(input, "ManageMasterUserPassword", _docdbManageMasterUserPassword); err != nil {
			log.Errorf("invalid --manage-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_docdbMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_docdbMasterUserPassword)
	}
	if len(_docdbMasterUserSecretKmsKeyId) > 0 {
		input.MasterUserSecretKmsKeyId = aws.String(_docdbMasterUserSecretKmsKeyId)
	}
	if len(_docdbNetworkType) > 0 {
		input.NetworkType = aws.String(_docdbNetworkType)
	}
	if len(_docdbNewDBClusterIdentifier) > 0 {
		input.NewDBClusterIdentifier = aws.String(_docdbNewDBClusterIdentifier)
	}
	if len(_docdbPort) > 0 {
		if err := assignInputField(input, "Port", _docdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_docdbPreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_docdbPreferredBackupWindow)
	}
	if len(_docdbPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_docdbPreferredMaintenanceWindow)
	}
	if len(_docdbRotateMasterUserPassword) > 0 {
		if err := assignInputField(input, "RotateMasterUserPassword", _docdbRotateMasterUserPassword); err != nil {
			log.Errorf("invalid --rotate-master-user-password: %s", err.Error())
			return
		}
	}
	if len(_docdbServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _docdbServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_docdbStorageType) > 0 {
		input.StorageType = aws.String(_docdbStorageType)
	}
	if len(_docdbVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a cluster parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName , ParameterValue , and
// ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot or maintenance window
//
// before the change can take effect.
//
// After you create a cluster parameter group, you should wait at least 5 minutes
// before creating your first cluster that uses that cluster parameter group as the
// default parameter group. This allows Amazon DocumentDB to fully complete the
// create action before the parameter group is used as the default for a new
// cluster. This step is especially important for parameters that are critical when
// creating the default database for a cluster, such as the character set for the
// default database defined by the character_set_database parameter.
func docdb_ModifyDBClusterParameterGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
		// Parameters: []types.Parameter, // Required
	}

	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbParameters) > 0 {
		if err := assignInputField(input, "Parameters", _docdbParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds an attribute and values to, or removes an attribute and values from, a
// manual cluster snapshot.
//
// To share a manual cluster snapshot with other Amazon Web Services accounts,
// specify restore as the AttributeName , and use the ValuesToAdd parameter to add
// a list of IDs of the Amazon Web Services accounts that are authorized to restore
// the manual cluster snapshot. Use the value all to make the manual cluster
// snapshot public, which means that it can be copied or restored by all Amazon Web
// Services accounts. Do not add the all value for any manual cluster snapshots
// that contain private information that you don't want available to all Amazon Web
// Services accounts. If a manual cluster snapshot is encrypted, it can be shared,
// but only by specifying a list of authorized Amazon Web Services account IDs for
// the ValuesToAdd parameter. You can't use all as a value for that parameter in
// this case.
func docdb_ModifyDBClusterSnapshotAttribute(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyDBClusterSnapshotAttributeInput{
		// AttributeName: *string, // Required
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_docdbAttributeName) > 0 {
		input.AttributeName = aws.String(_docdbAttributeName)
	}
	if len(_docdbDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_docdbDBClusterSnapshotIdentifier)
	}
	if len(_docdbValuesToAdd) > 0 {
		input.ValuesToAdd = append([]string(nil), _docdbValuesToAdd...)
	}
	if len(_docdbValuesToRemove) > 0 {
		input.ValuesToRemove = append([]string(nil), _docdbValuesToRemove...)
	}

	if resp, err := client.ModifyDBClusterSnapshotAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies settings for an instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request.
func docdb_ModifyDBInstance(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_docdbDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_docdbDBInstanceIdentifier)
	}
	if len(_docdbApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _docdbApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_docdbAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _docdbAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_docdbCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_docdbCACertificateIdentifier)
	}
	if len(_docdbCertificateRotationRestart) > 0 {
		if err := assignInputField(input, "CertificateRotationRestart", _docdbCertificateRotationRestart); err != nil {
			log.Errorf("invalid --certificate-rotation-restart: %s", err.Error())
			return
		}
	}
	if len(_docdbCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _docdbCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_docdbDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_docdbDBInstanceClass)
	}
	if len(_docdbEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _docdbEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_docdbNewDBInstanceIdentifier) > 0 {
		input.NewDBInstanceIdentifier = aws.String(_docdbNewDBInstanceIdentifier)
	}
	if len(_docdbPerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_docdbPerformanceInsightsKMSKeyId)
	}
	if len(_docdbPreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_docdbPreferredMaintenanceWindow)
	}
	if len(_docdbPromotionTier) > 0 {
		if err := assignInputField(input, "PromotionTier", _docdbPromotionTier); err != nil {
			log.Errorf("invalid --promotion-tier: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing subnet group. subnet groups must contain at least one
// subnet in at least two Availability Zones in the Amazon Web Services Region.
func docdb_ModifyDBSubnetGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyDBSubnetGroupInput{
		// DBSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}
	if len(_docdbSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _docdbSubnetIds...)
	}
	if len(_docdbDBSubnetGroupDescription) > 0 {
		input.DBSubnetGroupDescription = aws.String(_docdbDBSubnetGroupDescription)
	}

	if resp, err := client.ModifyDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing Amazon DocumentDB event notification subscription.
func docdb_ModifyEventSubscription(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_docdbSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_docdbSubscriptionName)
	}
	if len(_docdbEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _docdbEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_docdbEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _docdbEventCategories...)
	}
	if len(_docdbSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_docdbSnsTopicArn)
	}
	if len(_docdbSourceType) > 0 {
		input.SourceType = aws.String(_docdbSourceType)
	}

	if resp, err := client.ModifyEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify a setting for an Amazon DocumentDB global cluster. You can change one or
// more configuration parameters (for example: deletion protection), or the global
// cluster identifier by specifying these parameters and the new values in the
// request.
//
// This action only applies to Amazon DocumentDB clusters.
func docdb_ModifyGlobalCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ModifyGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}
	if len(_docdbDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _docdbDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_docdbNewGlobalClusterIdentifier) > 0 {
		input.NewGlobalClusterIdentifier = aws.String(_docdbNewGlobalClusterIdentifier)
	}

	if resp, err := client.ModifyGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You might need to reboot your instance, usually for maintenance reasons. For
// example, if you make certain changes, or if you change the cluster parameter
// group that is associated with the instance, you must reboot the instance for the
// changes to take effect.
//
// Rebooting an instance restarts the database engine service. Rebooting an
// instance results in a momentary outage, during which the instance status is set
// to rebooting.
func docdb_RebootDBInstance(cfg aws.Config, client *docdb.Client) {
	input := &docdb.RebootDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_docdbDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_docdbDBInstanceIdentifier)
	}
	if len(_docdbForceFailover) > 0 {
		if err := assignInputField(input, "ForceFailover", _docdbForceFailover); err != nil {
			log.Errorf("invalid --force-failover: %s", err.Error())
			return
		}
	}

	if resp, err := client.RebootDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detaches an Amazon DocumentDB secondary cluster from a global cluster. The
// cluster becomes a standalone cluster with read-write capability instead of being
// read-only and receiving data from a primary in a different region.
//
// This action only applies to Amazon DocumentDB clusters.
func docdb_RemoveFromGlobalCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.RemoveFromGlobalClusterInput{
		// DbClusterIdentifier: *string, // Required
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DbClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}

	if resp, err := client.RemoveFromGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a source identifier from an existing Amazon DocumentDB event
// notification subscription.
func docdb_RemoveSourceIdentifierFromSubscription(cfg aws.Config, client *docdb.Client) {
	input := &docdb.RemoveSourceIdentifierFromSubscriptionInput{
		// SourceIdentifier: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_docdbSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_docdbSourceIdentifier)
	}
	if len(_docdbSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_docdbSubscriptionName)
	}

	if resp, err := client.RemoveSourceIdentifierFromSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes metadata tags from an Amazon DocumentDB resource.
func docdb_RemoveTagsFromResource(cfg aws.Config, client *docdb.Client) {
	input := &docdb.RemoveTagsFromResourceInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_docdbResourceName) > 0 {
		input.ResourceName = aws.String(_docdbResourceName)
	}
	if len(_docdbTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _docdbTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a cluster parameter group to the default value. To
// reset specific parameters, submit a list of the following: ParameterName and
// ApplyMethod . To reset the entire cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When you reset the entire group, dynamic parameters are updated immediately and
// static parameters are set to pending-reboot to take effect on the next DB
// instance reboot.
func docdb_ResetDBClusterParameterGroup(cfg aws.Config, client *docdb.Client) {
	input := &docdb.ResetDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbParameters) > 0 {
		if err := assignInputField(input, "Parameters", _docdbParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_docdbResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _docdbResetAllParameters); err != nil {
			log.Errorf("invalid --reset-all-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cluster from a snapshot or cluster snapshot.
// If a snapshot is specified, the target cluster is created from the source DB
// snapshot with a default configuration and default security group.
//
// If a cluster snapshot is specified, the target cluster is created from the
// source cluster restore point with the same configuration as the original source
// DB cluster, except that the new cluster is created with the default security
// group.
func docdb_RestoreDBClusterFromSnapshot(cfg aws.Config, client *docdb.Client) {
	input := &docdb.RestoreDBClusterFromSnapshotInput{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
		// SnapshotIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbEngine) > 0 {
		input.Engine = aws.String(_docdbEngine)
	}
	if len(_docdbSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_docdbSnapshotIdentifier)
	}
	if len(_docdbAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _docdbAvailabilityZones...)
	}
	if len(_docdbDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_docdbDBClusterParameterGroupName)
	}
	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}
	if len(_docdbDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _docdbDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_docdbEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _docdbEnableCloudwatchLogsExports...)
	}
	if len(_docdbEngineVersion) > 0 {
		input.EngineVersion = aws.String(_docdbEngineVersion)
	}
	if len(_docdbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbKmsKeyId)
	}
	if len(_docdbNetworkType) > 0 {
		input.NetworkType = aws.String(_docdbNetworkType)
	}
	if len(_docdbPort) > 0 {
		if err := assignInputField(input, "Port", _docdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_docdbServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _docdbServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_docdbStorageType) > 0 {
		input.StorageType = aws.String(_docdbStorageType)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_docdbVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target cluster is created from the source cluster with the same
// configuration as the original cluster, except that the new cluster is created
// with the default security group.
func docdb_RestoreDBClusterToPointInTime(cfg aws.Config, client *docdb.Client) {
	input := &docdb.RestoreDBClusterToPointInTimeInput{
		// DBClusterIdentifier: *string, // Required
		// SourceDBClusterIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}
	if len(_docdbSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_docdbSourceDBClusterIdentifier)
	}
	if len(_docdbDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_docdbDBSubnetGroupName)
	}
	if len(_docdbDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _docdbDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_docdbEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _docdbEnableCloudwatchLogsExports...)
	}
	if len(_docdbKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_docdbKmsKeyId)
	}
	if len(_docdbNetworkType) > 0 {
		input.NetworkType = aws.String(_docdbNetworkType)
	}
	if len(_docdbPort) > 0 {
		if err := assignInputField(input, "Port", _docdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_docdbRestoreToTime) > 0 {
		if err := assignInputField(input, "RestoreToTime", _docdbRestoreToTime); err != nil {
			log.Errorf("invalid --restore-to-time: %s", err.Error())
			return
		}
	}
	if len(_docdbRestoreType) > 0 {
		input.RestoreType = aws.String(_docdbRestoreType)
	}
	if len(_docdbServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _docdbServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_docdbStorageType) > 0 {
		input.StorageType = aws.String(_docdbStorageType)
	}
	if len(_docdbTags) > 0 {
		if err := assignInputField(input, "Tags", _docdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_docdbUseLatestRestorableTime) > 0 {
		if err := assignInputField(input, "UseLatestRestorableTime", _docdbUseLatestRestorableTime); err != nil {
			log.Errorf("invalid --use-latest-restorable-time: %s", err.Error())
			return
		}
	}
	if len(_docdbVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _docdbVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterToPointInTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restarts the stopped cluster that is specified by DBClusterIdentifier . For more
// information, see [Stopping and Starting an Amazon DocumentDB Cluster].
//
// [Stopping and Starting an Amazon DocumentDB Cluster]: https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-stop-start.html
func docdb_StartDBCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.StartDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}

	if resp, err := client.StartDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops the running cluster that is specified by DBClusterIdentifier . The cluster
// must be in the available state. For more information, see [Stopping and Starting an Amazon DocumentDB Cluster].
//
// [Stopping and Starting an Amazon DocumentDB Cluster]: https://docs.aws.amazon.com/documentdb/latest/developerguide/db-cluster-stop-start.html
func docdb_StopDBCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.StopDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_docdbDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_docdbDBClusterIdentifier)
	}

	if resp, err := client.StopDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Switches over the specified secondary Amazon DocumentDB cluster to be the new
// primary Amazon DocumentDB cluster in the global database cluster.
func docdb_SwitchoverGlobalCluster(cfg aws.Config, client *docdb.Client) {
	input := &docdb.SwitchoverGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
		// TargetDbClusterIdentifier: *string, // Required
	}

	if len(_docdbGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_docdbGlobalClusterIdentifier)
	}
	if len(_docdbTargetDbClusterIdentifier) > 0 {
		input.TargetDbClusterIdentifier = aws.String(_docdbTargetDbClusterIdentifier)
	}

	if resp, err := client.SwitchoverGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_docdbCmd)
	_docdbCmd.Flags().SortFlags = false

	_docdbCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_docdbCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_docdbCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_docdbCmd.Flags().StringVarP(&_docdbAllowDataLoss, "allow-data-loss", "", "", "Allow Data Loss")
	_docdbCmd.Flags().StringVarP(&_docdbAllowMajorVersionUpgrade, "allow-major-version-upgrade", "", "", "Allow Major Version Upgrade")
	_docdbCmd.Flags().StringVarP(&_docdbApplyAction, "apply-action", "", "", "Apply Action")
	_docdbCmd.Flags().StringVarP(&_docdbApplyImmediately, "apply-immediately", "", "", "Apply Immediately")
	_docdbCmd.Flags().StringVarP(&_docdbAttributeName, "attribute-name", "", "", "Attribute Name")
	_docdbCmd.Flags().StringVarP(&_docdbAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_docdbCmd.Flags().StringVarP(&_docdbAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_docdbCmd.Flags().StringSliceVarP(&_docdbAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_docdbCmd.Flags().StringVarP(&_docdbBackupRetentionPeriod, "backup-retention-period", "", "", "Backup Retention Period")
	_docdbCmd.Flags().StringVarP(&_docdbCACertificateIdentifier, "ca-certificate-identifier", "", "", "Ca Certificate Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbCertificateIdentifier, "certificate-identifier", "", "", "Certificate Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbCertificateRotationRestart, "certificate-rotation-restart", "", "", "Certificate Rotation Restart")
	_docdbCmd.Flags().StringVarP(&_docdbCloudwatchLogsExportConfiguration, "cloudwatch-logs-export-configuration", "", "", "Cloudwatch Logs Export Configuration")
	_docdbCmd.Flags().StringVarP(&_docdbCopyTags, "copy-tags", "", "", "Copy Tags")
	_docdbCmd.Flags().StringVarP(&_docdbCopyTagsToSnapshot, "copy-tags-to-snapshot", "", "", "Copy Tags To Snapshot")
	_docdbCmd.Flags().StringVarP(&_docdbDatabaseName, "database-name", "", "", "Database Name")
	_docdbCmd.Flags().StringVarP(&_docdbDBClusterIdentifier, "db-cluster-identifier", "", "", "DB Cluster Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbDBClusterParameterGroupName, "db-cluster-parameter-group-name", "", "", "DB Cluster Parameter Group Name")
	_docdbCmd.Flags().StringVarP(&_docdbDBClusterSnapshotIdentifier, "db-cluster-snapshot-identifier", "", "", "DB Cluster Snapshot Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbDBInstanceClass, "db-instance-class", "", "", "DB Instance Class")
	_docdbCmd.Flags().StringVarP(&_docdbDBInstanceIdentifier, "db-instance-identifier", "", "", "DB Instance Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbDBParameterGroupFamily, "db-parameter-group-family", "", "", "DB Parameter Group Family")
	_docdbCmd.Flags().StringVarP(&_docdbDBSubnetGroupDescription, "db-subnet-group-description", "", "", "DB Subnet Group Description")
	_docdbCmd.Flags().StringVarP(&_docdbDBSubnetGroupName, "db-subnet-group-name", "", "", "DB Subnet Group Name")
	_docdbCmd.Flags().StringVarP(&_docdbDefaultOnly, "default-only", "", "", "Default Only")
	_docdbCmd.Flags().StringVarP(&_docdbDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_docdbCmd.Flags().StringVarP(&_docdbDescription, "description", "", "", "Description")
	_docdbCmd.Flags().StringVarP(&_docdbDuration, "duration", "", "", "Duration")
	_docdbCmd.Flags().StringSliceVarP(&_docdbEnableCloudwatchLogsExports, "enable-cloudwatch-logs-exports", "", nil, "Enable Cloudwatch Logs Exports")
	_docdbCmd.Flags().StringVarP(&_docdbEnablePerformanceInsights, "enable-performance-insights", "", "", "Enable Performance Insights")
	_docdbCmd.Flags().StringVarP(&_docdbEnabled, "enabled", "", "", "Enabled")
	_docdbCmd.Flags().StringVarP(&_docdbEndTime, "end-time", "", "", "End Time")
	_docdbCmd.Flags().StringVarP(&_docdbEngine, "engine", "", "", "Engine")
	_docdbCmd.Flags().StringVarP(&_docdbEngineVersion, "engine-version", "", "", "Engine Version")
	_docdbCmd.Flags().StringSliceVarP(&_docdbEventCategories, "event-categories", "", nil, "Event Categories")
	_docdbCmd.Flags().StringVarP(&_docdbFilters, "filters", "", "", "Filters")
	_docdbCmd.Flags().StringVarP(&_docdbFinalDBSnapshotIdentifier, "final-db-snapshot-identifier", "", "", "Final DB Snapshot Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbForceFailover, "force-failover", "", "", "Force Failover")
	_docdbCmd.Flags().StringVarP(&_docdbGlobalClusterIdentifier, "global-cluster-identifier", "", "", "Global Cluster Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbIncludePublic, "include-public", "", "", "Include Public")
	_docdbCmd.Flags().StringVarP(&_docdbIncludeShared, "include-shared", "", "", "Include Shared")
	_docdbCmd.Flags().StringVarP(&_docdbKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_docdbCmd.Flags().StringVarP(&_docdbLicenseModel, "license-model", "", "", "License Model")
	_docdbCmd.Flags().StringVarP(&_docdbListSupportedCharacterSets, "list-supported-character-sets", "", "", "List Supported Character Sets")
	_docdbCmd.Flags().StringVarP(&_docdbListSupportedTimezones, "list-supported-timezones", "", "", "List Supported Timezones")
	_docdbCmd.Flags().StringVarP(&_docdbManageMasterUserPassword, "manage-master-user-password", "", "", "Manage Master User Password")
	_docdbCmd.Flags().StringVarP(&_docdbMarker, "marker", "", "", "Marker")
	_docdbCmd.Flags().StringVarP(&_docdbMasterUserPassword, "master-user-password", "", "", "Master User Password")
	_docdbCmd.Flags().StringVarP(&_docdbMasterUserSecretKmsKeyId, "master-user-secret-kms-key-id", "", "", "Master User Secret KMS Key ID")
	_docdbCmd.Flags().StringVarP(&_docdbMasterUsername, "master-username", "", "", "Master Username")
	_docdbCmd.Flags().StringVarP(&_docdbMaxRecords, "max-records", "", "", "Max Records")
	_docdbCmd.Flags().StringVarP(&_docdbNetworkType, "network-type", "", "", "Network Type")
	_docdbCmd.Flags().StringVarP(&_docdbNewDBClusterIdentifier, "new-db-cluster-identifier", "", "", "New DB Cluster Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbNewDBInstanceIdentifier, "new-db-instance-identifier", "", "", "New DB Instance Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbNewGlobalClusterIdentifier, "new-global-cluster-identifier", "", "", "New Global Cluster Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbOptInType, "opt-in-type", "", "", "Opt In Type")
	_docdbCmd.Flags().StringVarP(&_docdbParameters, "parameters", "", "", "Parameters")
	_docdbCmd.Flags().StringVarP(&_docdbPerformanceInsightsKMSKeyId, "performance-insights-kms-key-id", "", "", "Performance Insights KMS Key ID")
	_docdbCmd.Flags().StringVarP(&_docdbPort, "port", "", "", "Port")
	_docdbCmd.Flags().StringVarP(&_docdbPreSignedUrl, "pre-signed-url", "", "", "Pre Signed URL")
	_docdbCmd.Flags().StringVarP(&_docdbPreferredBackupWindow, "preferred-backup-window", "", "", "Preferred Backup Window")
	_docdbCmd.Flags().StringVarP(&_docdbPreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_docdbCmd.Flags().StringVarP(&_docdbPromotionTier, "promotion-tier", "", "", "Promotion Tier")
	_docdbCmd.Flags().StringVarP(&_docdbResetAllParameters, "reset-all-parameters", "", "", "Reset All Parameters")
	_docdbCmd.Flags().StringVarP(&_docdbResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbResourceName, "resource-name", "", "", "Resource Name")
	_docdbCmd.Flags().StringVarP(&_docdbRestoreToTime, "restore-to-time", "", "", "Restore To Time")
	_docdbCmd.Flags().StringVarP(&_docdbRestoreType, "restore-type", "", "", "Restore Type")
	_docdbCmd.Flags().StringVarP(&_docdbRotateMasterUserPassword, "rotate-master-user-password", "", "", "Rotate Master User Password")
	_docdbCmd.Flags().StringVarP(&_docdbServerlessV2ScalingConfiguration, "serverless-v2-scaling-configuration", "", "", "Serverless V2 Scaling Configuration")
	_docdbCmd.Flags().StringVarP(&_docdbSkipFinalSnapshot, "skip-final-snapshot", "", "", "Skip Final Snapshot")
	_docdbCmd.Flags().StringVarP(&_docdbSnapshotIdentifier, "snapshot-identifier", "", "", "Snapshot Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbSnapshotType, "snapshot-type", "", "", "Snapshot Type")
	_docdbCmd.Flags().StringVarP(&_docdbSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_docdbCmd.Flags().StringVarP(&_docdbSource, "source", "", "", "Source")
	_docdbCmd.Flags().StringVarP(&_docdbSourceDBClusterIdentifier, "source-db-cluster-identifier", "", "", "Source DB Cluster Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbSourceDBClusterParameterGroupIdentifier, "source-db-cluster-parameter-group-identifier", "", "", "Source DB Cluster Parameter Group Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbSourceDBClusterSnapshotIdentifier, "source-db-cluster-snapshot-identifier", "", "", "Source DB Cluster Snapshot Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbSourceIdentifier, "source-identifier", "", "", "Source Identifier")
	_docdbCmd.Flags().StringSliceVarP(&_docdbSourceIds, "source-ids", "", nil, "Source Ids")
	_docdbCmd.Flags().StringVarP(&_docdbSourceRegion, "source-region", "", "", "Source Region")
	_docdbCmd.Flags().StringVarP(&_docdbSourceType, "source-type", "", "", "Source Type")
	_docdbCmd.Flags().StringVarP(&_docdbStartTime, "start-time", "", "", "Start Time")
	_docdbCmd.Flags().StringVarP(&_docdbStorageEncrypted, "storage-encrypted", "", "", "Storage Encrypted")
	_docdbCmd.Flags().StringVarP(&_docdbStorageType, "storage-type", "", "", "Storage Type")
	_docdbCmd.Flags().StringSliceVarP(&_docdbSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_docdbCmd.Flags().StringVarP(&_docdbSubscriptionName, "subscription-name", "", "", "Subscription Name")
	_docdbCmd.Flags().StringVarP(&_docdbSwitchover, "switchover", "", "", "Switchover")
	_docdbCmd.Flags().StringSliceVarP(&_docdbTagKeys, "tag-keys", "", nil, "Tag Keys")
	_docdbCmd.Flags().StringVarP(&_docdbTags, "tags", "", "", "Tags")
	_docdbCmd.Flags().StringVarP(&_docdbTargetDbClusterIdentifier, "target-db-cluster-identifier", "", "", "Target DB Cluster Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbTargetDBClusterParameterGroupDescription, "target-db-cluster-parameter-group-description", "", "", "Target DB Cluster Parameter Group Description")
	_docdbCmd.Flags().StringVarP(&_docdbTargetDBClusterParameterGroupIdentifier, "target-db-cluster-parameter-group-identifier", "", "", "Target DB Cluster Parameter Group Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbTargetDBClusterSnapshotIdentifier, "target-db-cluster-snapshot-identifier", "", "", "Target DB Cluster Snapshot Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbTargetDBInstanceIdentifier, "target-db-instance-identifier", "", "", "Target DB Instance Identifier")
	_docdbCmd.Flags().StringVarP(&_docdbUseLatestRestorableTime, "use-latest-restorable-time", "", "", "Use Latest Restorable Time")
	_docdbCmd.Flags().StringSliceVarP(&_docdbValuesToAdd, "values-to-add", "", nil, "Values To Add")
	_docdbCmd.Flags().StringSliceVarP(&_docdbValuesToRemove, "values-to-remove", "", nil, "Values To Remove")
	_docdbCmd.Flags().StringVarP(&_docdbVpc, "vpc", "", "", "VPC")
	_docdbCmd.Flags().StringSliceVarP(&_docdbVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")

	_docdbCmd.Flags().BoolVarP(&_docdbAddSourceIdentifierToSubscription, "add-source-identifier-to-subscription", "", false, "Add Source Identifier To Subscription")
	_docdbCmd.Flags().BoolVarP(&_docdbAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_docdbCmd.Flags().BoolVarP(&_docdbApplyPendingMaintenanceAction, "apply-pending-maintenance-action", "", false, "Apply Pending Maintenance Action")
	_docdbCmd.Flags().BoolVarP(&_docdbCopyDBClusterParameterGroup, "copy-db-cluster-parameter-group", "", false, "Copy DB Cluster Parameter Group")
	_docdbCmd.Flags().BoolVarP(&_docdbCopyDBClusterSnapshot, "copy-db-cluster-snapshot", "", false, "Copy DB Cluster Snapshot")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateDBCluster, "create-db-cluster", "", false, "Create DB Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateDBClusterParameterGroup, "create-db-cluster-parameter-group", "", false, "Create DB Cluster Parameter Group")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateDBClusterSnapshot, "create-db-cluster-snapshot", "", false, "Create DB Cluster Snapshot")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateDBInstance, "create-db-instance", "", false, "Create DB Instance")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateDBSubnetGroup, "create-db-subnet-group", "", false, "Create DB Subnet Group")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateEventSubscription, "create-event-subscription", "", false, "Create Event Subscription")
	_docdbCmd.Flags().BoolVarP(&_docdbCreateGlobalCluster, "create-global-cluster", "", false, "Create Global Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteDBCluster, "delete-db-cluster", "", false, "Delete DB Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteDBClusterParameterGroup, "delete-db-cluster-parameter-group", "", false, "Delete DB Cluster Parameter Group")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteDBClusterSnapshot, "delete-db-cluster-snapshot", "", false, "Delete DB Cluster Snapshot")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteDBInstance, "delete-db-instance", "", false, "Delete DB Instance")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteDBSubnetGroup, "delete-db-subnet-group", "", false, "Delete DB Subnet Group")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteEventSubscription, "delete-event-subscription", "", false, "Delete Event Subscription")
	_docdbCmd.Flags().BoolVarP(&_docdbDeleteGlobalCluster, "delete-global-cluster", "", false, "Delete Global Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeCertificates, "describe-certificates", "", false, "Describe Certificates")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBClusterParameterGroups, "describe-db-cluster-parameter-groups", "", false, "Describe DB Cluster Parameter Groups")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBClusterParameters, "describe-db-cluster-parameters", "", false, "Describe DB Cluster Parameters")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBClusterSnapshotAttributes, "describe-db-cluster-snapshot-attributes", "", false, "Describe DB Cluster Snapshot Attributes")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBClusterSnapshots, "describe-db-cluster-snapshots", "", false, "Describe DB Cluster Snapshots")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBClusters, "describe-db-clusters", "", false, "Describe DB Clusters")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBEngineVersions, "describe-db-engine-versions", "", false, "Describe DB Engine Versions")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBInstances, "describe-db-instances", "", false, "Describe DB Instances")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeDBSubnetGroups, "describe-db-subnet-groups", "", false, "Describe DB Subnet Groups")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeEngineDefaultClusterParameters, "describe-engine-default-cluster-parameters", "", false, "Describe Engine Default Cluster Parameters")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeEventCategories, "describe-event-categories", "", false, "Describe Event Categories")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeEventSubscriptions, "describe-event-subscriptions", "", false, "Describe Event Subscriptions")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeEvents, "describe-events", "", false, "Describe Events")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeGlobalClusters, "describe-global-clusters", "", false, "Describe Global Clusters")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribeOrderableDBInstanceOptions, "describe-orderable-db-instance-options", "", false, "Describe Orderable DB Instance Options")
	_docdbCmd.Flags().BoolVarP(&_docdbDescribePendingMaintenanceActions, "describe-pending-maintenance-actions", "", false, "Describe Pending Maintenance Actions")
	_docdbCmd.Flags().BoolVarP(&_docdbFailoverDBCluster, "failover-db-cluster", "", false, "Failover DB Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbFailoverGlobalCluster, "failover-global-cluster", "", false, "Failover Global Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyDBCluster, "modify-db-cluster", "", false, "Modify DB Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyDBClusterParameterGroup, "modify-db-cluster-parameter-group", "", false, "Modify DB Cluster Parameter Group")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyDBClusterSnapshotAttribute, "modify-db-cluster-snapshot-attribute", "", false, "Modify DB Cluster Snapshot Attribute")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyDBInstance, "modify-db-instance", "", false, "Modify DB Instance")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyDBSubnetGroup, "modify-db-subnet-group", "", false, "Modify DB Subnet Group")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyEventSubscription, "modify-event-subscription", "", false, "Modify Event Subscription")
	_docdbCmd.Flags().BoolVarP(&_docdbModifyGlobalCluster, "modify-global-cluster", "", false, "Modify Global Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbRebootDBInstance, "reboot-db-instance", "", false, "Reboot DB Instance")
	_docdbCmd.Flags().BoolVarP(&_docdbRemoveFromGlobalCluster, "remove-from-global-cluster", "", false, "Remove From Global Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbRemoveSourceIdentifierFromSubscription, "remove-source-identifier-from-subscription", "", false, "Remove Source Identifier From Subscription")
	_docdbCmd.Flags().BoolVarP(&_docdbRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_docdbCmd.Flags().BoolVarP(&_docdbResetDBClusterParameterGroup, "reset-db-cluster-parameter-group", "", false, "Reset DB Cluster Parameter Group")
	_docdbCmd.Flags().BoolVarP(&_docdbRestoreDBClusterFromSnapshot, "restore-db-cluster-from-snapshot", "", false, "Restore DB Cluster From Snapshot")
	_docdbCmd.Flags().BoolVarP(&_docdbRestoreDBClusterToPointInTime, "restore-db-cluster-to-point-in-time", "", false, "Restore DB Cluster To Point In Time")
	_docdbCmd.Flags().BoolVarP(&_docdbStartDBCluster, "start-db-cluster", "", false, "Start DB Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbStopDBCluster, "stop-db-cluster", "", false, "Stop DB Cluster")
	_docdbCmd.Flags().BoolVarP(&_docdbSwitchoverGlobalCluster, "switchover-global-cluster", "", false, "Switchover Global Cluster")

}
