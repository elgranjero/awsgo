package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/neptune"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// neptuneCmd represents the neptune command
var _neptuneCmd = &cobra.Command{
	Use:   "neptune",
	Short: "AWS neptune CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := neptune.NewFromConfig(cfg)
		if _neptuneAddRoleToDBCluster {
			neptune_AddRoleToDBCluster(cfg, client)
			return
		}
		if _neptuneAddSourceIdentifierToSubscription {
			neptune_AddSourceIdentifierToSubscription(cfg, client)
			return
		}
		if _neptuneAddTagsToResource {
			neptune_AddTagsToResource(cfg, client)
			return
		}
		if _neptuneApplyPendingMaintenanceAction {
			neptune_ApplyPendingMaintenanceAction(cfg, client)
			return
		}
		if _neptuneCopyDBClusterParameterGroup {
			neptune_CopyDBClusterParameterGroup(cfg, client)
			return
		}
		if _neptuneCopyDBClusterSnapshot {
			neptune_CopyDBClusterSnapshot(cfg, client)
			return
		}
		if _neptuneCopyDBParameterGroup {
			neptune_CopyDBParameterGroup(cfg, client)
			return
		}
		if _neptuneCreateDBCluster {
			neptune_CreateDBCluster(cfg, client)
			return
		}
		if _neptuneCreateDBClusterEndpoint {
			neptune_CreateDBClusterEndpoint(cfg, client)
			return
		}
		if _neptuneCreateDBClusterParameterGroup {
			neptune_CreateDBClusterParameterGroup(cfg, client)
			return
		}
		if _neptuneCreateDBClusterSnapshot {
			neptune_CreateDBClusterSnapshot(cfg, client)
			return
		}
		if _neptuneCreateDBInstance {
			neptune_CreateDBInstance(cfg, client)
			return
		}
		if _neptuneCreateDBParameterGroup {
			neptune_CreateDBParameterGroup(cfg, client)
			return
		}
		if _neptuneCreateDBSubnetGroup {
			neptune_CreateDBSubnetGroup(cfg, client)
			return
		}
		if _neptuneCreateEventSubscription {
			neptune_CreateEventSubscription(cfg, client)
			return
		}
		if _neptuneCreateGlobalCluster {
			neptune_CreateGlobalCluster(cfg, client)
			return
		}
		if _neptuneDeleteDBCluster {
			neptune_DeleteDBCluster(cfg, client)
			return
		}
		if _neptuneDeleteDBClusterEndpoint {
			neptune_DeleteDBClusterEndpoint(cfg, client)
			return
		}
		if _neptuneDeleteDBClusterParameterGroup {
			neptune_DeleteDBClusterParameterGroup(cfg, client)
			return
		}
		if _neptuneDeleteDBClusterSnapshot {
			neptune_DeleteDBClusterSnapshot(cfg, client)
			return
		}
		if _neptuneDeleteDBInstance {
			neptune_DeleteDBInstance(cfg, client)
			return
		}
		if _neptuneDeleteDBParameterGroup {
			neptune_DeleteDBParameterGroup(cfg, client)
			return
		}
		if _neptuneDeleteDBSubnetGroup {
			neptune_DeleteDBSubnetGroup(cfg, client)
			return
		}
		if _neptuneDeleteEventSubscription {
			neptune_DeleteEventSubscription(cfg, client)
			return
		}
		if _neptuneDeleteGlobalCluster {
			neptune_DeleteGlobalCluster(cfg, client)
			return
		}
		if _neptuneDescribeDBClusterEndpoints {
			neptune_DescribeDBClusterEndpoints(cfg, client)
			return
		}
		if _neptuneDescribeDBClusterParameterGroups {
			neptune_DescribeDBClusterParameterGroups(cfg, client)
			return
		}
		if _neptuneDescribeDBClusterParameters {
			neptune_DescribeDBClusterParameters(cfg, client)
			return
		}
		if _neptuneDescribeDBClusterSnapshotAttributes {
			neptune_DescribeDBClusterSnapshotAttributes(cfg, client)
			return
		}
		if _neptuneDescribeDBClusterSnapshots {
			neptune_DescribeDBClusterSnapshots(cfg, client)
			return
		}
		if _neptuneDescribeDBClusters {
			neptune_DescribeDBClusters(cfg, client)
			return
		}
		if _neptuneDescribeDBEngineVersions {
			neptune_DescribeDBEngineVersions(cfg, client)
			return
		}
		if _neptuneDescribeDBInstances {
			neptune_DescribeDBInstances(cfg, client)
			return
		}
		if _neptuneDescribeDBParameterGroups {
			neptune_DescribeDBParameterGroups(cfg, client)
			return
		}
		if _neptuneDescribeDBParameters {
			neptune_DescribeDBParameters(cfg, client)
			return
		}
		if _neptuneDescribeDBSubnetGroups {
			neptune_DescribeDBSubnetGroups(cfg, client)
			return
		}
		if _neptuneDescribeEngineDefaultClusterParameters {
			neptune_DescribeEngineDefaultClusterParameters(cfg, client)
			return
		}
		if _neptuneDescribeEngineDefaultParameters {
			neptune_DescribeEngineDefaultParameters(cfg, client)
			return
		}
		if _neptuneDescribeEventCategories {
			neptune_DescribeEventCategories(cfg, client)
			return
		}
		if _neptuneDescribeEventSubscriptions {
			neptune_DescribeEventSubscriptions(cfg, client)
			return
		}
		if _neptuneDescribeEvents {
			neptune_DescribeEvents(cfg, client)
			return
		}
		if _neptuneDescribeGlobalClusters {
			neptune_DescribeGlobalClusters(cfg, client)
			return
		}
		if _neptuneDescribeOrderableDBInstanceOptions {
			neptune_DescribeOrderableDBInstanceOptions(cfg, client)
			return
		}
		if _neptuneDescribePendingMaintenanceActions {
			neptune_DescribePendingMaintenanceActions(cfg, client)
			return
		}
		if _neptuneDescribeValidDBInstanceModifications {
			neptune_DescribeValidDBInstanceModifications(cfg, client)
			return
		}
		if _neptuneFailoverDBCluster {
			neptune_FailoverDBCluster(cfg, client)
			return
		}
		if _neptuneFailoverGlobalCluster {
			neptune_FailoverGlobalCluster(cfg, client)
			return
		}
		if _neptuneListTagsForResource {
			neptune_ListTagsForResource(cfg, client)
			return
		}
		if _neptuneModifyDBCluster {
			neptune_ModifyDBCluster(cfg, client)
			return
		}
		if _neptuneModifyDBClusterEndpoint {
			neptune_ModifyDBClusterEndpoint(cfg, client)
			return
		}
		if _neptuneModifyDBClusterParameterGroup {
			neptune_ModifyDBClusterParameterGroup(cfg, client)
			return
		}
		if _neptuneModifyDBClusterSnapshotAttribute {
			neptune_ModifyDBClusterSnapshotAttribute(cfg, client)
			return
		}
		if _neptuneModifyDBInstance {
			neptune_ModifyDBInstance(cfg, client)
			return
		}
		if _neptuneModifyDBParameterGroup {
			neptune_ModifyDBParameterGroup(cfg, client)
			return
		}
		if _neptuneModifyDBSubnetGroup {
			neptune_ModifyDBSubnetGroup(cfg, client)
			return
		}
		if _neptuneModifyEventSubscription {
			neptune_ModifyEventSubscription(cfg, client)
			return
		}
		if _neptuneModifyGlobalCluster {
			neptune_ModifyGlobalCluster(cfg, client)
			return
		}
		if _neptunePromoteReadReplicaDBCluster {
			neptune_PromoteReadReplicaDBCluster(cfg, client)
			return
		}
		if _neptuneRebootDBInstance {
			neptune_RebootDBInstance(cfg, client)
			return
		}
		if _neptuneRemoveFromGlobalCluster {
			neptune_RemoveFromGlobalCluster(cfg, client)
			return
		}
		if _neptuneRemoveRoleFromDBCluster {
			neptune_RemoveRoleFromDBCluster(cfg, client)
			return
		}
		if _neptuneRemoveSourceIdentifierFromSubscription {
			neptune_RemoveSourceIdentifierFromSubscription(cfg, client)
			return
		}
		if _neptuneRemoveTagsFromResource {
			neptune_RemoveTagsFromResource(cfg, client)
			return
		}
		if _neptuneResetDBClusterParameterGroup {
			neptune_ResetDBClusterParameterGroup(cfg, client)
			return
		}
		if _neptuneResetDBParameterGroup {
			neptune_ResetDBParameterGroup(cfg, client)
			return
		}
		if _neptuneRestoreDBClusterFromSnapshot {
			neptune_RestoreDBClusterFromSnapshot(cfg, client)
			return
		}
		if _neptuneRestoreDBClusterToPointInTime {
			neptune_RestoreDBClusterToPointInTime(cfg, client)
			return
		}
		if _neptuneStartDBCluster {
			neptune_StartDBCluster(cfg, client)
			return
		}
		if _neptuneStopDBCluster {
			neptune_StopDBCluster(cfg, client)
			return
		}
		if _neptuneSwitchoverGlobalCluster {
			neptune_SwitchoverGlobalCluster(cfg, client)
			return
		}

	},
}

var (
	_neptuneAddRoleToDBCluster                     bool
	_neptuneAddSourceIdentifierToSubscription      bool
	_neptuneAddTagsToResource                      bool
	_neptuneApplyPendingMaintenanceAction          bool
	_neptuneCopyDBClusterParameterGroup            bool
	_neptuneCopyDBClusterSnapshot                  bool
	_neptuneCopyDBParameterGroup                   bool
	_neptuneCreateDBCluster                        bool
	_neptuneCreateDBClusterEndpoint                bool
	_neptuneCreateDBClusterParameterGroup          bool
	_neptuneCreateDBClusterSnapshot                bool
	_neptuneCreateDBInstance                       bool
	_neptuneCreateDBParameterGroup                 bool
	_neptuneCreateDBSubnetGroup                    bool
	_neptuneCreateEventSubscription                bool
	_neptuneCreateGlobalCluster                    bool
	_neptuneDeleteDBCluster                        bool
	_neptuneDeleteDBClusterEndpoint                bool
	_neptuneDeleteDBClusterParameterGroup          bool
	_neptuneDeleteDBClusterSnapshot                bool
	_neptuneDeleteDBInstance                       bool
	_neptuneDeleteDBParameterGroup                 bool
	_neptuneDeleteDBSubnetGroup                    bool
	_neptuneDeleteEventSubscription                bool
	_neptuneDeleteGlobalCluster                    bool
	_neptuneDescribeDBClusterEndpoints             bool
	_neptuneDescribeDBClusterParameterGroups       bool
	_neptuneDescribeDBClusterParameters            bool
	_neptuneDescribeDBClusterSnapshotAttributes    bool
	_neptuneDescribeDBClusterSnapshots             bool
	_neptuneDescribeDBClusters                     bool
	_neptuneDescribeDBEngineVersions               bool
	_neptuneDescribeDBInstances                    bool
	_neptuneDescribeDBParameterGroups              bool
	_neptuneDescribeDBParameters                   bool
	_neptuneDescribeDBSubnetGroups                 bool
	_neptuneDescribeEngineDefaultClusterParameters bool
	_neptuneDescribeEngineDefaultParameters        bool
	_neptuneDescribeEventCategories                bool
	_neptuneDescribeEventSubscriptions             bool
	_neptuneDescribeEvents                         bool
	_neptuneDescribeGlobalClusters                 bool
	_neptuneDescribeOrderableDBInstanceOptions     bool
	_neptuneDescribePendingMaintenanceActions      bool
	_neptuneDescribeValidDBInstanceModifications   bool
	_neptuneFailoverDBCluster                      bool
	_neptuneFailoverGlobalCluster                  bool
	_neptuneListTagsForResource                    bool
	_neptuneModifyDBCluster                        bool
	_neptuneModifyDBClusterEndpoint                bool
	_neptuneModifyDBClusterParameterGroup          bool
	_neptuneModifyDBClusterSnapshotAttribute       bool
	_neptuneModifyDBInstance                       bool
	_neptuneModifyDBParameterGroup                 bool
	_neptuneModifyDBSubnetGroup                    bool
	_neptuneModifyEventSubscription                bool
	_neptuneModifyGlobalCluster                    bool
	_neptunePromoteReadReplicaDBCluster            bool
	_neptuneRebootDBInstance                       bool
	_neptuneRemoveFromGlobalCluster                bool
	_neptuneRemoveRoleFromDBCluster                bool
	_neptuneRemoveSourceIdentifierFromSubscription bool
	_neptuneRemoveTagsFromResource                 bool
	_neptuneResetDBClusterParameterGroup           bool
	_neptuneResetDBParameterGroup                  bool
	_neptuneRestoreDBClusterFromSnapshot           bool
	_neptuneRestoreDBClusterToPointInTime          bool
	_neptuneStartDBCluster                         bool
	_neptuneStopDBCluster                          bool
	_neptuneSwitchoverGlobalCluster                bool

	_neptuneAllocatedStorage                         string
	_neptuneAllowDataLoss                            string
	_neptuneAllowMajorVersionUpgrade                 string
	_neptuneApplyAction                              string
	_neptuneApplyImmediately                         string
	_neptuneAttributeName                            string
	_neptuneAutoMinorVersionUpgrade                  string
	_neptuneAvailabilityZone                         string
	_neptuneAvailabilityZones                        []string
	_neptuneBackupRetentionPeriod                    string
	_neptuneCACertificateIdentifier                  string
	_neptuneCharacterSetName                         string
	_neptuneCloudwatchLogsExportConfiguration        string
	_neptuneCopyTags                                 string
	_neptuneCopyTagsToSnapshot                       string
	_neptuneDatabaseName                             string
	_neptuneDBClusterEndpointIdentifier              string
	_neptuneDBClusterIdentifier                      string
	_neptuneDBClusterParameterGroupName              string
	_neptuneDBClusterSnapshotIdentifier              string
	_neptuneDBInstanceClass                          string
	_neptuneDBInstanceIdentifier                     string
	_neptuneDBInstanceParameterGroupName             string
	_neptuneDBName                                   string
	_neptuneDBParameterGroupFamily                   string
	_neptuneDBParameterGroupName                     string
	_neptuneDBPortNumber                             string
	_neptuneDBSecurityGroups                         []string
	_neptuneDBSubnetGroupDescription                 string
	_neptuneDBSubnetGroupName                        string
	_neptuneDefaultOnly                              string
	_neptuneDeletionProtection                       string
	_neptuneDescription                              string
	_neptuneDomain                                   string
	_neptuneDomainIAMRoleName                        string
	_neptuneDuration                                 string
	_neptuneEnableCloudwatchLogsExports              []string
	_neptuneEnableIAMDatabaseAuthentication          string
	_neptuneEnablePerformanceInsights                string
	_neptuneEnabled                                  string
	_neptuneEndTime                                  string
	_neptuneEndpointType                             string
	_neptuneEngine                                   string
	_neptuneEngineVersion                            string
	_neptuneEventCategories                          []string
	_neptuneExcludedMembers                          []string
	_neptuneFeatureName                              string
	_neptuneFilters                                  string
	_neptuneFinalDBSnapshotIdentifier                string
	_neptuneForceFailover                            string
	_neptuneGlobalClusterIdentifier                  string
	_neptuneIncludePublic                            string
	_neptuneIncludeShared                            string
	_neptuneIops                                     string
	_neptuneKmsKeyId                                 string
	_neptuneLicenseModel                             string
	_neptuneListSupportedCharacterSets               string
	_neptuneListSupportedTimezones                   string
	_neptuneMarker                                   string
	_neptuneMasterUserPassword                       string
	_neptuneMasterUsername                           string
	_neptuneMaxRecords                               string
	_neptuneMonitoringInterval                       string
	_neptuneMonitoringRoleArn                        string
	_neptuneMultiAZ                                  string
	_neptuneNewDBClusterIdentifier                   string
	_neptuneNewDBInstanceIdentifier                  string
	_neptuneNewGlobalClusterIdentifier               string
	_neptuneOptInType                                string
	_neptuneOptionGroupName                          string
	_neptuneParameters                               string
	_neptunePerformanceInsightsKMSKeyId              string
	_neptunePort                                     string
	_neptunePreSignedUrl                             string
	_neptunePreferredBackupWindow                    string
	_neptunePreferredMaintenanceWindow               string
	_neptunePromotionTier                            string
	_neptunePubliclyAccessible                       string
	_neptuneReplicationSourceIdentifier              string
	_neptuneResetAllParameters                       string
	_neptuneResourceIdentifier                       string
	_neptuneResourceName                             string
	_neptuneRestoreToTime                            string
	_neptuneRestoreType                              string
	_neptuneRoleArn                                  string
	_neptuneServerlessV2ScalingConfiguration         string
	_neptuneSkipFinalSnapshot                        string
	_neptuneSnapshotIdentifier                       string
	_neptuneSnapshotType                             string
	_neptuneSnsTopicArn                              string
	_neptuneSource                                   string
	_neptuneSourceDBClusterIdentifier                string
	_neptuneSourceDBClusterParameterGroupIdentifier  string
	_neptuneSourceDBClusterSnapshotIdentifier        string
	_neptuneSourceDBParameterGroupIdentifier         string
	_neptuneSourceIdentifier                         string
	_neptuneSourceIds                                []string
	_neptuneSourceRegion                             string
	_neptuneSourceType                               string
	_neptuneStartTime                                string
	_neptuneStaticMembers                            []string
	_neptuneStorageEncrypted                         string
	_neptuneStorageType                              string
	_neptuneSubnetIds                                []string
	_neptuneSubscriptionName                         string
	_neptuneSwitchover                               string
	_neptuneTagKeys                                  []string
	_neptuneTags                                     string
	_neptuneTargetDbClusterIdentifier                string
	_neptuneTargetDBClusterParameterGroupDescription string
	_neptuneTargetDBClusterParameterGroupIdentifier  string
	_neptuneTargetDBClusterSnapshotIdentifier        string
	_neptuneTargetDBInstanceIdentifier               string
	_neptuneTargetDBParameterGroupDescription        string
	_neptuneTargetDBParameterGroupIdentifier         string
	_neptuneTdeCredentialArn                         string
	_neptuneTdeCredentialPassword                    string
	_neptuneTimezone                                 string
	_neptuneUseLatestRestorableTime                  string
	_neptuneValuesToAdd                              []string
	_neptuneValuesToRemove                           []string
	_neptuneVpc                                      string
	_neptuneVpcSecurityGroupIds                      []string
)

// Associates an Identity and Access Management (IAM) role with an Neptune DB
// cluster.
func neptune_AddRoleToDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.AddRoleToDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneRoleArn) > 0 {
		input.RoleArn = aws.String(_neptuneRoleArn)
	}
	if len(_neptuneFeatureName) > 0 {
		input.FeatureName = aws.String(_neptuneFeatureName)
	}

	if resp, err := client.AddRoleToDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a source identifier to an existing event notification subscription.
func neptune_AddSourceIdentifierToSubscription(cfg aws.Config, client *neptune.Client) {
	input := &neptune.AddSourceIdentifierToSubscriptionInput{
		// SourceIdentifier: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_neptuneSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_neptuneSourceIdentifier)
	}
	if len(_neptuneSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_neptuneSubscriptionName)
	}

	if resp, err := client.AddSourceIdentifierToSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds metadata tags to an Amazon Neptune resource. These tags can also be used
// with cost allocation reporting to track cost associated with Amazon Neptune
// resources, or used in a Condition statement in an IAM policy for Amazon Neptune.
func neptune_AddTagsToResource(cfg aws.Config, client *neptune.Client) {
	input := &neptune.AddTagsToResourceInput{
		// ResourceName: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_neptuneResourceName) > 0 {
		input.ResourceName = aws.String(_neptuneResourceName)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Applies a pending maintenance action to a resource (for example, to a DB
// instance).
func neptune_ApplyPendingMaintenanceAction(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ApplyPendingMaintenanceActionInput{
		// ApplyAction: *string, // Required
		// OptInType: *string, // Required
		// ResourceIdentifier: *string, // Required
	}

	if len(_neptuneApplyAction) > 0 {
		input.ApplyAction = aws.String(_neptuneApplyAction)
	}
	if len(_neptuneOptInType) > 0 {
		input.OptInType = aws.String(_neptuneOptInType)
	}
	if len(_neptuneResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_neptuneResourceIdentifier)
	}

	if resp, err := client.ApplyPendingMaintenanceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified DB cluster parameter group.
func neptune_CopyDBClusterParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CopyDBClusterParameterGroupInput{
		// SourceDBClusterParameterGroupIdentifier: *string, // Required
		// TargetDBClusterParameterGroupDescription: *string, // Required
		// TargetDBClusterParameterGroupIdentifier: *string, // Required
	}

	if len(_neptuneSourceDBClusterParameterGroupIdentifier) > 0 {
		input.SourceDBClusterParameterGroupIdentifier = aws.String(_neptuneSourceDBClusterParameterGroupIdentifier)
	}
	if len(_neptuneTargetDBClusterParameterGroupDescription) > 0 {
		input.TargetDBClusterParameterGroupDescription = aws.String(_neptuneTargetDBClusterParameterGroupDescription)
	}
	if len(_neptuneTargetDBClusterParameterGroupIdentifier) > 0 {
		input.TargetDBClusterParameterGroupIdentifier = aws.String(_neptuneTargetDBClusterParameterGroupIdentifier)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Copies a snapshot of a DB cluster.
// To copy a DB cluster snapshot from a shared manual DB cluster snapshot,
// SourceDBClusterSnapshotIdentifier must be the Amazon Resource Name (ARN) of the
// shared DB cluster snapshot.
func neptune_CopyDBClusterSnapshot(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CopyDBClusterSnapshotInput{
		// SourceDBClusterSnapshotIdentifier: *string, // Required
		// TargetDBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_neptuneSourceDBClusterSnapshotIdentifier) > 0 {
		input.SourceDBClusterSnapshotIdentifier = aws.String(_neptuneSourceDBClusterSnapshotIdentifier)
	}
	if len(_neptuneTargetDBClusterSnapshotIdentifier) > 0 {
		input.TargetDBClusterSnapshotIdentifier = aws.String(_neptuneTargetDBClusterSnapshotIdentifier)
	}
	if len(_neptuneCopyTags) > 0 {
		if err := assignInputField(input, "CopyTags", _neptuneCopyTags); err != nil {
			log.Errorf("invalid --copy-tags: %s", err.Error())
			return
		}
	}
	if len(_neptuneKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_neptuneKmsKeyId)
	}
	if len(_neptunePreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_neptunePreSignedUrl)
	}
	if len(_neptuneSourceRegion) > 0 {
		input.SourceRegion = aws.String(_neptuneSourceRegion)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Copies the specified DB parameter group.
func neptune_CopyDBParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CopyDBParameterGroupInput{
		// SourceDBParameterGroupIdentifier: *string, // Required
		// TargetDBParameterGroupDescription: *string, // Required
		// TargetDBParameterGroupIdentifier: *string, // Required
	}

	if len(_neptuneSourceDBParameterGroupIdentifier) > 0 {
		input.SourceDBParameterGroupIdentifier = aws.String(_neptuneSourceDBParameterGroupIdentifier)
	}
	if len(_neptuneTargetDBParameterGroupDescription) > 0 {
		input.TargetDBParameterGroupDescription = aws.String(_neptuneTargetDBParameterGroupDescription)
	}
	if len(_neptuneTargetDBParameterGroupIdentifier) > 0 {
		input.TargetDBParameterGroupIdentifier = aws.String(_neptuneTargetDBParameterGroupIdentifier)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Amazon Neptune DB cluster.
// You can use the ReplicationSourceIdentifier parameter to create the DB cluster
// as a Read Replica of another DB cluster or Amazon Neptune DB instance.
//
// Note that when you create a new cluster using CreateDBCluster directly,
// deletion protection is disabled by default (when you create a new production
// cluster in the console, deletion protection is enabled by default). You can only
// delete a DB cluster if its DeletionProtection field is set to false .
func neptune_CreateDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneEngine) > 0 {
		input.Engine = aws.String(_neptuneEngine)
	}
	if len(_neptuneAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _neptuneAvailabilityZones...)
	}
	if len(_neptuneBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _neptuneBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_neptuneCharacterSetName) > 0 {
		input.CharacterSetName = aws.String(_neptuneCharacterSetName)
	}
	if len(_neptuneCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _neptuneCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneDatabaseName) > 0 {
		input.DatabaseName = aws.String(_neptuneDatabaseName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _neptuneEnableCloudwatchLogsExports...)
	}
	if len(_neptuneEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _neptuneEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}
	if len(_neptuneKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_neptuneKmsKeyId)
	}
	if len(_neptuneMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_neptuneMasterUserPassword)
	}
	if len(_neptuneMasterUsername) > 0 {
		input.MasterUsername = aws.String(_neptuneMasterUsername)
	}
	if len(_neptuneOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_neptuneOptionGroupName)
	}
	if len(_neptunePort) > 0 {
		if err := assignInputField(input, "Port", _neptunePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_neptunePreSignedUrl) > 0 {
		input.PreSignedUrl = aws.String(_neptunePreSignedUrl)
	}
	if len(_neptunePreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_neptunePreferredBackupWindow)
	}
	if len(_neptunePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_neptunePreferredMaintenanceWindow)
	}
	if len(_neptuneReplicationSourceIdentifier) > 0 {
		input.ReplicationSourceIdentifier = aws.String(_neptuneReplicationSourceIdentifier)
	}
	if len(_neptuneServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _neptuneServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptuneSourceRegion) > 0 {
		input.SourceRegion = aws.String(_neptuneSourceRegion)
	}
	if len(_neptuneStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _neptuneStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageType) > 0 {
		input.StorageType = aws.String(_neptuneStorageType)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_neptuneVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptuneVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new custom endpoint and associates it with an Amazon Neptune DB
// cluster.
func neptune_CreateDBClusterEndpoint(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBClusterEndpointInput{
		// DBClusterEndpointIdentifier: *string, // Required
		// DBClusterIdentifier: *string, // Required
		// EndpointType: *string, // Required
	}

	if len(_neptuneDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_neptuneDBClusterEndpointIdentifier)
	}
	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneEndpointType) > 0 {
		input.EndpointType = aws.String(_neptuneEndpointType)
	}
	if len(_neptuneExcludedMembers) > 0 {
		input.ExcludedMembers = append([]string(nil), _neptuneExcludedMembers...)
	}
	if len(_neptuneStaticMembers) > 0 {
		input.StaticMembers = append([]string(nil), _neptuneStaticMembers...)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBClusterEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB cluster parameter group.
// Parameters in a DB cluster parameter group apply to all of the instances in a
// DB cluster.
//
// A DB cluster parameter group is initially created with the default parameters
// for the database engine used by instances in the DB cluster. To provide custom
// values for any of the parameters, you must modify the group after creating it
// using ModifyDBClusterParameterGroup. Once you've created a DB cluster parameter group, you need to associate
// it with your DB cluster using ModifyDBCluster. When you associate a new DB cluster parameter
// group with a running DB cluster, you need to reboot the DB instances in the DB
// cluster without failover for the new DB cluster parameter group and associated
// settings to take effect.
//
// After you create a DB cluster parameter group, you should wait at least 5
// minutes before creating your first DB cluster that uses that DB cluster
// parameter group as the default parameter group. This allows Amazon Neptune to
// fully complete the create action before the DB cluster parameter group is used
// as the default for a new DB cluster. This is especially important for parameters
// that are critical when creating the default database for a DB cluster, such as
// the character set for the default database defined by the character_set_database
// parameter. You can use the Parameter Groups option of the [Amazon Neptune console]or the DescribeDBClusterParameters command to
// verify that your DB cluster parameter group has been created or modified.
//
// [Amazon Neptune console]: https://console.aws.amazon.com/rds/
func neptune_CreateDBClusterParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
		// DBParameterGroupFamily: *string, // Required
		// Description: *string, // Required
	}

	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_neptuneDBParameterGroupFamily)
	}
	if len(_neptuneDescription) > 0 {
		input.Description = aws.String(_neptuneDescription)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Creates a snapshot of a DB cluster.
func neptune_CreateDBClusterSnapshot(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBClusterSnapshotInput{
		// DBClusterIdentifier: *string, // Required
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_neptuneDBClusterSnapshotIdentifier)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Creates a new DB instance.
func neptune_CreateDBInstance(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBInstanceInput{
		// DBClusterIdentifier: *string, // Required
		// DBInstanceClass: *string, // Required
		// DBInstanceIdentifier: *string, // Required
		// Engine: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_neptuneDBInstanceClass)
	}
	if len(_neptuneDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_neptuneDBInstanceIdentifier)
	}
	if len(_neptuneEngine) > 0 {
		input.Engine = aws.String(_neptuneEngine)
	}
	if len(_neptuneAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _neptuneAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_neptuneAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _neptuneAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_neptuneAvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_neptuneAvailabilityZone)
	}
	if len(_neptuneBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _neptuneBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_neptuneCharacterSetName) > 0 {
		input.CharacterSetName = aws.String(_neptuneCharacterSetName)
	}
	if len(_neptuneCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _neptuneCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_neptuneDBName) > 0 {
		input.DBName = aws.String(_neptuneDBName)
	}
	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneDBSecurityGroups) > 0 {
		input.DBSecurityGroups = append([]string(nil), _neptuneDBSecurityGroups...)
	}
	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneDomain) > 0 {
		input.Domain = aws.String(_neptuneDomain)
	}
	if len(_neptuneDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_neptuneDomainIAMRoleName)
	}
	if len(_neptuneEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _neptuneEnableCloudwatchLogsExports...)
	}
	if len(_neptuneEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _neptuneEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_neptuneEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _neptuneEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneIops) > 0 {
		if err := assignInputField(input, "Iops", _neptuneIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_neptuneKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_neptuneKmsKeyId)
	}
	if len(_neptuneLicenseModel) > 0 {
		input.LicenseModel = aws.String(_neptuneLicenseModel)
	}
	if len(_neptuneMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_neptuneMasterUserPassword)
	}
	if len(_neptuneMasterUsername) > 0 {
		input.MasterUsername = aws.String(_neptuneMasterUsername)
	}
	if len(_neptuneMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _neptuneMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_neptuneMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_neptuneMonitoringRoleArn)
	}
	if len(_neptuneMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _neptuneMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_neptuneOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_neptuneOptionGroupName)
	}
	if len(_neptunePerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_neptunePerformanceInsightsKMSKeyId)
	}
	if len(_neptunePort) > 0 {
		if err := assignInputField(input, "Port", _neptunePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_neptunePreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_neptunePreferredBackupWindow)
	}
	if len(_neptunePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_neptunePreferredMaintenanceWindow)
	}
	if len(_neptunePromotionTier) > 0 {
		if err := assignInputField(input, "PromotionTier", _neptunePromotionTier); err != nil {
			log.Errorf("invalid --promotion-tier: %s", err.Error())
			return
		}
	}
	if len(_neptunePubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _neptunePubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _neptuneStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageType) > 0 {
		input.StorageType = aws.String(_neptuneStorageType)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_neptuneTdeCredentialArn) > 0 {
		input.TdeCredentialArn = aws.String(_neptuneTdeCredentialArn)
	}
	if len(_neptuneTdeCredentialPassword) > 0 {
		input.TdeCredentialPassword = aws.String(_neptuneTdeCredentialPassword)
	}
	if len(_neptuneTimezone) > 0 {
		input.Timezone = aws.String(_neptuneTimezone)
	}
	if len(_neptuneVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptuneVpcSecurityGroupIds...)
	}

	if resp, err := client.CreateDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB parameter group.
// A DB parameter group is initially created with the default parameters for the
// database engine used by the DB instance. To provide custom values for any of the
// parameters, you must modify the group after creating it using
// ModifyDBParameterGroup. Once you've created a DB parameter group, you need to
// associate it with your DB instance using ModifyDBInstance. When you associate a
// new DB parameter group with a running DB instance, you need to reboot the DB
// instance without failover for the new DB parameter group and associated settings
// to take effect.
//
// After you create a DB parameter group, you should wait at least 5 minutes
// before creating your first DB instance that uses that DB parameter group as the
// default parameter group. This allows Amazon Neptune to fully complete the create
// action before the parameter group is used as the default for a new DB instance.
// This is especially important for parameters that are critical when creating the
// default database for a DB instance, such as the character set for the default
// database defined by the character_set_database parameter. You can use the
// Parameter Groups option of the Amazon Neptune console or the
// DescribeDBParameters command to verify that your DB parameter group has been
// created or modified.
func neptune_CreateDBParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBParameterGroupInput{
		// DBParameterGroupFamily: *string, // Required
		// DBParameterGroupName: *string, // Required
		// Description: *string, // Required
	}

	if len(_neptuneDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_neptuneDBParameterGroupFamily)
	}
	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneDescription) > 0 {
		input.Description = aws.String(_neptuneDescription)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB subnet group. DB subnet groups must contain at least one
// subnet in at least two AZs in the Amazon Region.
func neptune_CreateDBSubnetGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateDBSubnetGroupInput{
		// DBSubnetGroupDescription: *string, // Required
		// DBSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_neptuneDBSubnetGroupDescription) > 0 {
		input.DBSubnetGroupDescription = aws.String(_neptuneDBSubnetGroupDescription)
	}
	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _neptuneSubnetIds...)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Creates an event notification subscription. This action requires a topic ARN
// (Amazon Resource Name) created by either the Neptune console, the SNS console,
// or the SNS API. To obtain an ARN with SNS, you must create a topic in Amazon SNS
// and subscribe to the topic. The ARN is displayed in the SNS console.
//
// You can specify the type of source (SourceType) you want to be notified of,
// provide a list of Neptune sources (SourceIds) that triggers the events, and
// provide a list of event categories (EventCategories) for events you want to be
// notified of. For example, you can specify SourceType = db-instance, SourceIds =
// mydbinstance1, mydbinstance2 and EventCategories = Availability, Backup.
//
// If you specify both the SourceType and SourceIds, such as SourceType =
// db-instance and SourceIdentifier = myDBInstance1, you are notified of all the
// db-instance events for the specified source. If you specify a SourceType but do
// not specify a SourceIdentifier, you receive notice of the events for that source
// type for all your Neptune sources. If you do not specify either the SourceType
// nor the SourceIdentifier, you are notified of events generated from all Neptune
// sources belonging to your customer account.
func neptune_CreateEventSubscription(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateEventSubscriptionInput{
		// SnsTopicArn: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_neptuneSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_neptuneSnsTopicArn)
	}
	if len(_neptuneSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_neptuneSubscriptionName)
	}
	if len(_neptuneEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _neptuneEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_neptuneEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _neptuneEventCategories...)
	}
	if len(_neptuneSourceIds) > 0 {
		input.SourceIds = append([]string(nil), _neptuneSourceIds...)
	}
	if len(_neptuneSourceType) > 0 {
		input.SourceType = aws.String(_neptuneSourceType)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
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

// Creates a Neptune global database spread across multiple Amazon Regions. The
// global database contains a single primary cluster with read-write capability,
// and read-only secondary clusters that receive data from the primary cluster
// through high-speed replication performed by the Neptune storage subsystem.
//
// You can create a global database that is initially empty, and then add a
// primary cluster and secondary clusters to it, or you can specify an existing
// Neptune cluster during the create operation to become the primary cluster of the
// global database.
func neptune_CreateGlobalCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.CreateGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}
	if len(_neptuneDatabaseName) > 0 {
		input.DatabaseName = aws.String(_neptuneDatabaseName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngine) > 0 {
		input.Engine = aws.String(_neptuneEngine)
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_neptuneSourceDBClusterIdentifier)
	}
	if len(_neptuneStorageEncrypted) > 0 {
		if err := assignInputField(input, "StorageEncrypted", _neptuneStorageEncrypted); err != nil {
			log.Errorf("invalid --storage-encrypted: %s", err.Error())
			return
		}
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// The DeleteDBCluster action deletes a previously provisioned DB cluster. When
// you delete a DB cluster, all automated backups for that DB cluster are deleted
// and can't be recovered. Manual DB cluster snapshots of the specified DB cluster
// are not deleted.
//
// Note that the DB Cluster cannot be deleted if deletion protection is enabled.
// To delete it, you must first set its DeletionProtection field to False .
func neptune_DeleteDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneFinalDBSnapshotIdentifier) > 0 {
		input.FinalDBSnapshotIdentifier = aws.String(_neptuneFinalDBSnapshotIdentifier)
	}
	if len(_neptuneSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _neptuneSkipFinalSnapshot); err != nil {
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

// Deletes a custom endpoint and removes it from an Amazon Neptune DB cluster.
func neptune_DeleteDBClusterEndpoint(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBClusterEndpointInput{
		// DBClusterEndpointIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_neptuneDBClusterEndpointIdentifier)
	}

	if resp, err := client.DeleteDBClusterEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified DB cluster parameter group. The DB cluster parameter group
// to be deleted can't be associated with any DB clusters.
func neptune_DeleteDBClusterParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}

	if resp, err := client.DeleteDBClusterParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DB cluster snapshot. If the snapshot is being copied, the copy
// operation is terminated.
//
// The DB cluster snapshot must be in the available state to be deleted.
func neptune_DeleteDBClusterSnapshot(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBClusterSnapshotInput{
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_neptuneDBClusterSnapshotIdentifier)
	}

	if resp, err := client.DeleteDBClusterSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The DeleteDBInstance action deletes a previously provisioned DB instance. When
// you delete a DB instance, all automated backups for that instance are deleted
// and can't be recovered. Manual DB snapshots of the DB instance to be deleted by
// DeleteDBInstance are not deleted.
//
// If you request a final DB snapshot the status of the Amazon Neptune DB instance
// is deleting until the DB snapshot is created. The API action DescribeDBInstance
// is used to monitor the status of this operation. The action can't be canceled or
// reverted once submitted.
//
// Note that when a DB instance is in a failure state and has a status of failed ,
// incompatible-restore , or incompatible-network , you can only delete it when the
// SkipFinalSnapshot parameter is set to true .
//
// You can't delete a DB instance if it is the only instance in the DB cluster, or
// if it has deletion protection enabled.
func neptune_DeleteDBInstance(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_neptuneDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_neptuneDBInstanceIdentifier)
	}
	if len(_neptuneFinalDBSnapshotIdentifier) > 0 {
		input.FinalDBSnapshotIdentifier = aws.String(_neptuneFinalDBSnapshotIdentifier)
	}
	if len(_neptuneSkipFinalSnapshot) > 0 {
		if err := assignInputField(input, "SkipFinalSnapshot", _neptuneSkipFinalSnapshot); err != nil {
			log.Errorf("invalid --skip-final-snapshot: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified DBParameterGroup. The DBParameterGroup to be deleted can't
// be associated with any DB instances.
func neptune_DeleteDBParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBParameterGroupInput{
		// DBParameterGroupName: *string, // Required
	}

	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}

	if resp, err := client.DeleteDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a DB subnet group.
// The specified database subnet group must not be associated with any DB
// instances.
func neptune_DeleteDBSubnetGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteDBSubnetGroupInput{
		// DBSubnetGroupName: *string, // Required
	}

	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}

	if resp, err := client.DeleteDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an event notification subscription.
func neptune_DeleteEventSubscription(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_neptuneSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_neptuneSubscriptionName)
	}

	if resp, err := client.DeleteEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a global database. The primary and all secondary clusters must already
// be detached or deleted first.
func neptune_DeleteGlobalCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DeleteGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}

	if resp, err := client.DeleteGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about endpoints for an Amazon Neptune DB cluster.
// This operation can also return information for Amazon RDS clusters and Amazon
// DocDB clusters.
func neptune_DescribeDBClusterEndpoints(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBClusterEndpointsInput{}

	if len(_neptuneDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_neptuneDBClusterEndpointIdentifier)
	}
	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBClusterEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptune.DescribeDBClusterEndpointsOutput
	p := neptune.NewDescribeDBClusterEndpointsPaginator(client, input)
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
// DBClusterParameterGroupName parameter is specified, the list will contain only
// the description of the specified DB cluster parameter group.
func neptune_DescribeDBClusterParameterGroups(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBClusterParameterGroupsInput{}

	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

	var results []*neptune.DescribeDBClusterParameterGroupsOutput
	p := neptune.NewDescribeDBClusterParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular DB cluster parameter group.
func neptune_DescribeDBClusterParameters(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBClusterParametersInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneSource) > 0 {
		input.Source = aws.String(_neptuneSource)
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

	var results []*neptune.DescribeDBClusterParametersOutput
	p := neptune.NewDescribeDBClusterParametersPaginator(client, input)
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

// Returns a list of DB cluster snapshot attribute names and values for a manual
// DB cluster snapshot.
//
// When sharing snapshots with other Amazon accounts,
// DescribeDBClusterSnapshotAttributes returns the restore attribute and a list of
// IDs for the Amazon accounts that are authorized to copy or restore the manual DB
// cluster snapshot. If all is included in the list of values for the restore
// attribute, then the manual DB cluster snapshot is public and can be copied or
// restored by all Amazon accounts.
//
// To add or remove access for an Amazon account to copy or restore a manual DB
// cluster snapshot, or to make the manual DB cluster snapshot public or private,
// use the ModifyDBClusterSnapshotAttributeAPI action.
func neptune_DescribeDBClusterSnapshotAttributes(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBClusterSnapshotAttributesInput{
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_neptuneDBClusterSnapshotIdentifier)
	}

	if resp, err := client.DescribeDBClusterSnapshotAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about DB cluster snapshots. This API action supports
// pagination.
func neptune_DescribeDBClusterSnapshots(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBClusterSnapshotsInput{}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_neptuneDBClusterSnapshotIdentifier)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneIncludePublic) > 0 {
		if err := assignInputField(input, "IncludePublic", _neptuneIncludePublic); err != nil {
			log.Errorf("invalid --include-public: %s", err.Error())
			return
		}
	}
	if len(_neptuneIncludeShared) > 0 {
		if err := assignInputField(input, "IncludeShared", _neptuneIncludeShared); err != nil {
			log.Errorf("invalid --include-shared: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneSnapshotType) > 0 {
		input.SnapshotType = aws.String(_neptuneSnapshotType)
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

	var results []*neptune.DescribeDBClusterSnapshotsOutput
	p := neptune.NewDescribeDBClusterSnapshotsPaginator(client, input)
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

// Returns information about provisioned DB clusters, and supports pagination.
// This operation can also return information for Amazon RDS clusters and Amazon
// DocDB clusters.
func neptune_DescribeDBClusters(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBClustersInput{}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

	var results []*neptune.DescribeDBClustersOutput
	p := neptune.NewDescribeDBClustersPaginator(client, input)
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

// Returns a list of the available DB engines.
func neptune_DescribeDBEngineVersions(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBEngineVersionsInput{}

	if len(_neptuneDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_neptuneDBParameterGroupFamily)
	}
	if len(_neptuneDefaultOnly) > 0 {
		if err := assignInputField(input, "DefaultOnly", _neptuneDefaultOnly); err != nil {
			log.Errorf("invalid --default-only: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngine) > 0 {
		input.Engine = aws.String(_neptuneEngine)
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneListSupportedCharacterSets) > 0 {
		if err := assignInputField(input, "ListSupportedCharacterSets", _neptuneListSupportedCharacterSets); err != nil {
			log.Errorf("invalid --list-supported-character-sets: %s", err.Error())
			return
		}
	}
	if len(_neptuneListSupportedTimezones) > 0 {
		if err := assignInputField(input, "ListSupportedTimezones", _neptuneListSupportedTimezones); err != nil {
			log.Errorf("invalid --list-supported-timezones: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

	var results []*neptune.DescribeDBEngineVersionsOutput
	p := neptune.NewDescribeDBEngineVersionsPaginator(client, input)
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

// Returns information about provisioned instances, and supports pagination.
// This operation can also return information for Amazon RDS instances and Amazon
// DocDB instances.
func neptune_DescribeDBInstances(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBInstancesInput{}

	if len(_neptuneDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_neptuneDBInstanceIdentifier)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

	var results []*neptune.DescribeDBInstancesOutput
	p := neptune.NewDescribeDBInstancesPaginator(client, input)
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

// Returns a list of DBParameterGroup descriptions. If a DBParameterGroupName is
// specified, the list will contain only the description of the specified DB
// parameter group.
func neptune_DescribeDBParameterGroups(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBParameterGroupsInput{}

	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptune.DescribeDBParameterGroupsOutput
	p := neptune.NewDescribeDBParameterGroupsPaginator(client, input)
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

// Returns the detailed parameter list for a particular DB parameter group.
func neptune_DescribeDBParameters(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBParametersInput{
		// DBParameterGroupName: *string, // Required
	}

	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneSource) > 0 {
		input.Source = aws.String(_neptuneSource)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDBParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptune.DescribeDBParametersOutput
	p := neptune.NewDescribeDBParametersPaginator(client, input)
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
// DBSubnetGroup.
//
// For an overview of CIDR ranges, go to the [Wikipedia Tutorial].
//
// [Wikipedia Tutorial]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
func neptune_DescribeDBSubnetGroups(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeDBSubnetGroupsInput{}

	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

	var results []*neptune.DescribeDBSubnetGroupsOutput
	p := neptune.NewDescribeDBSubnetGroupsPaginator(client, input)
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
func neptune_DescribeEngineDefaultClusterParameters(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeEngineDefaultClusterParametersInput{
		// DBParameterGroupFamily: *string, // Required
	}

	if len(_neptuneDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_neptuneDBParameterGroupFamily)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

// Returns the default engine and system parameter information for the specified
// database engine.
func neptune_DescribeEngineDefaultParameters(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeEngineDefaultParametersInput{
		// DBParameterGroupFamily: *string, // Required
	}

	if len(_neptuneDBParameterGroupFamily) > 0 {
		input.DBParameterGroupFamily = aws.String(_neptuneDBParameterGroupFamily)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeEngineDefaultParameters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*neptune.DescribeEngineDefaultParametersOutput
	p := neptune.NewDescribeEngineDefaultParametersPaginator(client, input)
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

// Displays a list of categories for all event source types, or, if specified, for
// a specified source type.
func neptune_DescribeEventCategories(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeEventCategoriesInput{}

	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneSourceType) > 0 {
		input.SourceType = aws.String(_neptuneSourceType)
	}

	if resp, err := client.DescribeEventCategories(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the subscription descriptions for a customer account. The description
// for a subscription includes SubscriptionName, SNSTopicARN, CustomerID,
// SourceType, SourceID, CreationTime, and Status.
//
// If you specify a SubscriptionName, lists the description for that subscription.
func neptune_DescribeEventSubscriptions(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeEventSubscriptionsInput{}

	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_neptuneSubscriptionName)
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

	var results []*neptune.DescribeEventSubscriptionsOutput
	p := neptune.NewDescribeEventSubscriptionsPaginator(client, input)
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

// Returns events related to DB instances, DB security groups, DB snapshots, and
// DB parameter groups for the past 14 days. Events specific to a particular DB
// instance, DB security group, database snapshot, or DB parameter group can be
// obtained by providing the name as a parameter. By default, the past hour of
// events are returned.
func neptune_DescribeEvents(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeEventsInput{}

	if len(_neptuneDuration) > 0 {
		if err := assignInputField(input, "Duration", _neptuneDuration); err != nil {
			log.Errorf("invalid --duration: %s", err.Error())
			return
		}
	}
	if len(_neptuneEndTime) > 0 {
		if err := assignInputField(input, "EndTime", _neptuneEndTime); err != nil {
			log.Errorf("invalid --end-time: %s", err.Error())
			return
		}
	}
	if len(_neptuneEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _neptuneEventCategories...)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_neptuneSourceIdentifier)
	}
	if len(_neptuneSourceType) > 0 {
		if err := assignInputField(input, "SourceType", _neptuneSourceType); err != nil {
			log.Errorf("invalid --source-type: %s", err.Error())
			return
		}
	}
	if len(_neptuneStartTime) > 0 {
		if err := assignInputField(input, "StartTime", _neptuneStartTime); err != nil {
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

	var results []*neptune.DescribeEventsOutput
	p := neptune.NewDescribeEventsPaginator(client, input)
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

// Returns information about Neptune global database clusters. This API supports
// pagination.
func neptune_DescribeGlobalClusters(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeGlobalClustersInput{}

	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
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

	var results []*neptune.DescribeGlobalClustersOutput
	p := neptune.NewDescribeGlobalClustersPaginator(client, input)
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

// Returns a list of orderable DB instance options for the specified engine.
func neptune_DescribeOrderableDBInstanceOptions(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeOrderableDBInstanceOptionsInput{
		// Engine: *string, // Required
	}

	if len(_neptuneEngine) > 0 {
		input.Engine = aws.String(_neptuneEngine)
	}
	if len(_neptuneDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_neptuneDBInstanceClass)
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneLicenseModel) > 0 {
		input.LicenseModel = aws.String(_neptuneLicenseModel)
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneVpc) > 0 {
		if err := assignInputField(input, "Vpc", _neptuneVpc); err != nil {
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

	var results []*neptune.DescribeOrderableDBInstanceOptionsOutput
	p := neptune.NewDescribeOrderableDBInstanceOptionsPaginator(client, input)
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

// Returns a list of resources (for example, DB instances) that have at least one
// pending maintenance action.
func neptune_DescribePendingMaintenanceActions(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribePendingMaintenanceActionsInput{}

	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_neptuneMarker) > 0 {
		input.Marker = aws.String(_neptuneMarker)
	}
	if len(_neptuneMaxRecords) > 0 {
		if err := assignInputField(input, "MaxRecords", _neptuneMaxRecords); err != nil {
			log.Errorf("invalid --max-records: %s", err.Error())
			return
		}
	}
	if len(_neptuneResourceIdentifier) > 0 {
		input.ResourceIdentifier = aws.String(_neptuneResourceIdentifier)
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

	var results []*neptune.DescribePendingMaintenanceActionsOutput
	p := neptune.NewDescribePendingMaintenanceActionsPaginator(client, input)
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

// You can call DescribeValidDBInstanceModifications to learn what modifications you can make to your DB instance. You
// can use this information when you call ModifyDBInstance.
func neptune_DescribeValidDBInstanceModifications(cfg aws.Config, client *neptune.Client) {
	input := &neptune.DescribeValidDBInstanceModificationsInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_neptuneDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_neptuneDBInstanceIdentifier)
	}

	if resp, err := client.DescribeValidDBInstanceModifications(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Forces a failover for a DB cluster.
// A failover for a DB cluster promotes one of the Read Replicas (read-only
// instances) in the DB cluster to be the primary instance (the cluster writer).
//
// Amazon Neptune will automatically fail over to a Read Replica, if one exists,
// when the primary instance fails. You can force a failover when you want to
// simulate a failure of a primary instance for testing. Because each instance in a
// DB cluster has its own endpoint address, you will need to clean up and
// re-establish any existing connections that use those endpoint addresses when the
// failover is complete.
func neptune_FailoverDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.FailoverDBClusterInput{}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneTargetDBInstanceIdentifier) > 0 {
		input.TargetDBInstanceIdentifier = aws.String(_neptuneTargetDBInstanceIdentifier)
	}

	if resp, err := client.FailoverDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates the failover process for a Neptune global database.
// A failover for a Neptune global database promotes one of secondary read-only DB
// clusters to be the primary DB cluster and demotes the primary DB cluster to
// being a secondary (read-only) DB cluster. In other words, the role of the
// current primary DB cluster and the selected target secondary DB cluster are
// switched. The selected secondary DB cluster assumes full read/write capabilities
// for the Neptune global database.
//
// This action applies only to Neptune global databases. This action is only
// intended for use on healthy Neptune global databases with healthy Neptune DB
// clusters and no region-wide outages, to test disaster recovery scenarios or to
// reconfigure the global database topology.
func neptune_FailoverGlobalCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.FailoverGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
		// TargetDbClusterIdentifier: *string, // Required
	}

	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}
	if len(_neptuneTargetDbClusterIdentifier) > 0 {
		input.TargetDbClusterIdentifier = aws.String(_neptuneTargetDbClusterIdentifier)
	}
	if len(_neptuneAllowDataLoss) > 0 {
		if err := assignInputField(input, "AllowDataLoss", _neptuneAllowDataLoss); err != nil {
			log.Errorf("invalid --allow-data-loss: %s", err.Error())
			return
		}
	}
	if len(_neptuneSwitchover) > 0 {
		if err := assignInputField(input, "Switchover", _neptuneSwitchover); err != nil {
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

// Lists all tags on an Amazon Neptune resource.
func neptune_ListTagsForResource(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ListTagsForResourceInput{
		// ResourceName: *string, // Required
	}

	if len(_neptuneResourceName) > 0 {
		input.ResourceName = aws.String(_neptuneResourceName)
	}
	if len(_neptuneFilters) > 0 {
		if err := assignInputField(input, "Filters", _neptuneFilters); err != nil {
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

// Modify a setting for a DB cluster. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request.
func neptune_ModifyDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _neptuneAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_neptuneApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _neptuneApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_neptuneBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _neptuneBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_neptuneCloudwatchLogsExportConfiguration) > 0 {
		if err := assignInputField(input, "CloudwatchLogsExportConfiguration", _neptuneCloudwatchLogsExportConfiguration); err != nil {
			log.Errorf("invalid --cloudwatch-logs-export-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptuneCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _neptuneCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneDBInstanceParameterGroupName) > 0 {
		input.DBInstanceParameterGroupName = aws.String(_neptuneDBInstanceParameterGroupName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _neptuneEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_neptuneMasterUserPassword)
	}
	if len(_neptuneNewDBClusterIdentifier) > 0 {
		input.NewDBClusterIdentifier = aws.String(_neptuneNewDBClusterIdentifier)
	}
	if len(_neptuneOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_neptuneOptionGroupName)
	}
	if len(_neptunePort) > 0 {
		if err := assignInputField(input, "Port", _neptunePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_neptunePreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_neptunePreferredBackupWindow)
	}
	if len(_neptunePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_neptunePreferredMaintenanceWindow)
	}
	if len(_neptuneServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _neptuneServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageType) > 0 {
		input.StorageType = aws.String(_neptuneStorageType)
	}
	if len(_neptuneVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptuneVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the properties of an endpoint in an Amazon Neptune DB cluster.
func neptune_ModifyDBClusterEndpoint(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBClusterEndpointInput{
		// DBClusterEndpointIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterEndpointIdentifier) > 0 {
		input.DBClusterEndpointIdentifier = aws.String(_neptuneDBClusterEndpointIdentifier)
	}
	if len(_neptuneEndpointType) > 0 {
		input.EndpointType = aws.String(_neptuneEndpointType)
	}
	if len(_neptuneExcludedMembers) > 0 {
		input.ExcludedMembers = append([]string(nil), _neptuneExcludedMembers...)
	}
	if len(_neptuneStaticMembers) > 0 {
		input.StaticMembers = append([]string(nil), _neptuneStaticMembers...)
	}

	if resp, err := client.ModifyDBClusterEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB cluster parameter group. To modify more than
// one parameter, submit a list of the following: ParameterName , ParameterValue ,
// and ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot without failover to the DB cluster associated with
// the parameter group before the change can take effect.
//
// After you create a DB cluster parameter group, you should wait at least 5
// minutes before creating your first DB cluster that uses that DB cluster
// parameter group as the default parameter group. This allows Amazon Neptune to
// fully complete the create action before the parameter group is used as the
// default for a new DB cluster. This is especially important for parameters that
// are critical when creating the default database for a DB cluster, such as the
// character set for the default database defined by the character_set_database
// parameter. You can use the Parameter Groups option of the Amazon Neptune console
// or the DescribeDBClusterParameterscommand to verify that your DB cluster parameter group has been created
// or modified.
func neptune_ModifyDBClusterParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
		// Parameters: []types.Parameter, // Required
	}

	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _neptuneParameters); err != nil {
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
// manual DB cluster snapshot.
//
// To share a manual DB cluster snapshot with other Amazon accounts, specify
// restore as the AttributeName and use the ValuesToAdd parameter to add a list of
// IDs of the Amazon accounts that are authorized to restore the manual DB cluster
// snapshot. Use the value all to make the manual DB cluster snapshot public,
// which means that it can be copied or restored by all Amazon accounts. Do not add
// the all value for any manual DB cluster snapshots that contain private
// information that you don't want available to all Amazon accounts. If a manual DB
// cluster snapshot is encrypted, it can be shared, but only by specifying a list
// of authorized Amazon account IDs for the ValuesToAdd parameter. You can't use
// all as a value for that parameter in this case.
//
// To view which Amazon accounts have access to copy or restore a manual DB
// cluster snapshot, or whether a manual DB cluster snapshot public or private, use
// the DescribeDBClusterSnapshotAttributesAPI action.
func neptune_ModifyDBClusterSnapshotAttribute(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBClusterSnapshotAttributeInput{
		// AttributeName: *string, // Required
		// DBClusterSnapshotIdentifier: *string, // Required
	}

	if len(_neptuneAttributeName) > 0 {
		input.AttributeName = aws.String(_neptuneAttributeName)
	}
	if len(_neptuneDBClusterSnapshotIdentifier) > 0 {
		input.DBClusterSnapshotIdentifier = aws.String(_neptuneDBClusterSnapshotIdentifier)
	}
	if len(_neptuneValuesToAdd) > 0 {
		input.ValuesToAdd = append([]string(nil), _neptuneValuesToAdd...)
	}
	if len(_neptuneValuesToRemove) > 0 {
		input.ValuesToRemove = append([]string(nil), _neptuneValuesToRemove...)
	}

	if resp, err := client.ModifyDBClusterSnapshotAttribute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies settings for a DB instance. You can change one or more database
// configuration parameters by specifying these parameters and the new values in
// the request. To learn what modifications you can make to your DB instance, call DescribeValidDBInstanceModifications
// before you call ModifyDBInstance.
func neptune_ModifyDBInstance(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_neptuneDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_neptuneDBInstanceIdentifier)
	}
	if len(_neptuneAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _neptuneAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_neptuneAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _neptuneAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_neptuneApplyImmediately) > 0 {
		if err := assignInputField(input, "ApplyImmediately", _neptuneApplyImmediately); err != nil {
			log.Errorf("invalid --apply-immediately: %s", err.Error())
			return
		}
	}
	if len(_neptuneAutoMinorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AutoMinorVersionUpgrade", _neptuneAutoMinorVersionUpgrade); err != nil {
			log.Errorf("invalid --auto-minor-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_neptuneBackupRetentionPeriod) > 0 {
		if err := assignInputField(input, "BackupRetentionPeriod", _neptuneBackupRetentionPeriod); err != nil {
			log.Errorf("invalid --backup-retention-period: %s", err.Error())
			return
		}
	}
	if len(_neptuneCACertificateIdentifier) > 0 {
		input.CACertificateIdentifier = aws.String(_neptuneCACertificateIdentifier)
	}
	if len(_neptuneCloudwatchLogsExportConfiguration) > 0 {
		if err := assignInputField(input, "CloudwatchLogsExportConfiguration", _neptuneCloudwatchLogsExportConfiguration); err != nil {
			log.Errorf("invalid --cloudwatch-logs-export-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptuneCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _neptuneCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_neptuneDBInstanceClass) > 0 {
		input.DBInstanceClass = aws.String(_neptuneDBInstanceClass)
	}
	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneDBPortNumber) > 0 {
		if err := assignInputField(input, "DBPortNumber", _neptuneDBPortNumber); err != nil {
			log.Errorf("invalid --db-port-number: %s", err.Error())
			return
		}
	}
	if len(_neptuneDBSecurityGroups) > 0 {
		input.DBSecurityGroups = append([]string(nil), _neptuneDBSecurityGroups...)
	}
	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneDomain) > 0 {
		input.Domain = aws.String(_neptuneDomain)
	}
	if len(_neptuneDomainIAMRoleName) > 0 {
		input.DomainIAMRoleName = aws.String(_neptuneDomainIAMRoleName)
	}
	if len(_neptuneEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _neptuneEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_neptuneEnablePerformanceInsights) > 0 {
		if err := assignInputField(input, "EnablePerformanceInsights", _neptuneEnablePerformanceInsights); err != nil {
			log.Errorf("invalid --enable-performance-insights: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneIops) > 0 {
		if err := assignInputField(input, "Iops", _neptuneIops); err != nil {
			log.Errorf("invalid --iops: %s", err.Error())
			return
		}
	}
	if len(_neptuneLicenseModel) > 0 {
		input.LicenseModel = aws.String(_neptuneLicenseModel)
	}
	if len(_neptuneMasterUserPassword) > 0 {
		input.MasterUserPassword = aws.String(_neptuneMasterUserPassword)
	}
	if len(_neptuneMonitoringInterval) > 0 {
		if err := assignInputField(input, "MonitoringInterval", _neptuneMonitoringInterval); err != nil {
			log.Errorf("invalid --monitoring-interval: %s", err.Error())
			return
		}
	}
	if len(_neptuneMonitoringRoleArn) > 0 {
		input.MonitoringRoleArn = aws.String(_neptuneMonitoringRoleArn)
	}
	if len(_neptuneMultiAZ) > 0 {
		if err := assignInputField(input, "MultiAZ", _neptuneMultiAZ); err != nil {
			log.Errorf("invalid --multi-az: %s", err.Error())
			return
		}
	}
	if len(_neptuneNewDBInstanceIdentifier) > 0 {
		input.NewDBInstanceIdentifier = aws.String(_neptuneNewDBInstanceIdentifier)
	}
	if len(_neptuneOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_neptuneOptionGroupName)
	}
	if len(_neptunePerformanceInsightsKMSKeyId) > 0 {
		input.PerformanceInsightsKMSKeyId = aws.String(_neptunePerformanceInsightsKMSKeyId)
	}
	if len(_neptunePreferredBackupWindow) > 0 {
		input.PreferredBackupWindow = aws.String(_neptunePreferredBackupWindow)
	}
	if len(_neptunePreferredMaintenanceWindow) > 0 {
		input.PreferredMaintenanceWindow = aws.String(_neptunePreferredMaintenanceWindow)
	}
	if len(_neptunePromotionTier) > 0 {
		if err := assignInputField(input, "PromotionTier", _neptunePromotionTier); err != nil {
			log.Errorf("invalid --promotion-tier: %s", err.Error())
			return
		}
	}
	if len(_neptunePubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _neptunePubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageType) > 0 {
		input.StorageType = aws.String(_neptuneStorageType)
	}
	if len(_neptuneTdeCredentialArn) > 0 {
		input.TdeCredentialArn = aws.String(_neptuneTdeCredentialArn)
	}
	if len(_neptuneTdeCredentialPassword) > 0 {
		input.TdeCredentialPassword = aws.String(_neptuneTdeCredentialPassword)
	}
	if len(_neptuneVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptuneVpcSecurityGroupIds...)
	}

	if resp, err := client.ModifyDBInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB parameter group. To modify more than one
// parameter, submit a list of the following: ParameterName , ParameterValue , and
// ApplyMethod . A maximum of 20 parameters can be modified in a single request.
//
// Changes to dynamic parameters are applied immediately. Changes to static
// parameters require a reboot without failover to the DB instance associated with
// the parameter group before the change can take effect.
//
// After you modify a DB parameter group, you should wait at least 5 minutes
// before creating your first DB instance that uses that DB parameter group as the
// default parameter group. This allows Amazon Neptune to fully complete the modify
// action before the parameter group is used as the default for a new DB instance.
// This is especially important for parameters that are critical when creating the
// default database for a DB instance, such as the character set for the default
// database defined by the character_set_database parameter. You can use the
// Parameter Groups option of the Amazon Neptune console or the
// DescribeDBParameters command to verify that your DB parameter group has been
// created or modified.
func neptune_ModifyDBParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBParameterGroupInput{
		// DBParameterGroupName: *string, // Required
		// Parameters: []types.Parameter, // Required
	}

	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _neptuneParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing DB subnet group. DB subnet groups must contain at least
// one subnet in at least two AZs in the Amazon Region.
func neptune_ModifyDBSubnetGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyDBSubnetGroupInput{
		// DBSubnetGroupName: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _neptuneSubnetIds...)
	}
	if len(_neptuneDBSubnetGroupDescription) > 0 {
		input.DBSubnetGroupDescription = aws.String(_neptuneDBSubnetGroupDescription)
	}

	if resp, err := client.ModifyDBSubnetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an existing event notification subscription. Note that you can't
// modify the source identifiers using this call; to change source identifiers for
// a subscription, use the AddSourceIdentifierToSubscriptionand RemoveSourceIdentifierFromSubscription calls.
//
// You can see a list of the event categories for a given SourceType by using the
// DescribeEventCategories action.
func neptune_ModifyEventSubscription(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyEventSubscriptionInput{
		// SubscriptionName: *string, // Required
	}

	if len(_neptuneSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_neptuneSubscriptionName)
	}
	if len(_neptuneEnabled) > 0 {
		if err := assignInputField(input, "Enabled", _neptuneEnabled); err != nil {
			log.Errorf("invalid --enabled: %s", err.Error())
			return
		}
	}
	if len(_neptuneEventCategories) > 0 {
		input.EventCategories = append([]string(nil), _neptuneEventCategories...)
	}
	if len(_neptuneSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_neptuneSnsTopicArn)
	}
	if len(_neptuneSourceType) > 0 {
		input.SourceType = aws.String(_neptuneSourceType)
	}

	if resp, err := client.ModifyEventSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modify a setting for an Amazon Neptune global cluster. You can change one or
// more database configuration parameters by specifying these parameters and their
// new values in the request.
func neptune_ModifyGlobalCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ModifyGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}
	if len(_neptuneAllowMajorVersionUpgrade) > 0 {
		if err := assignInputField(input, "AllowMajorVersionUpgrade", _neptuneAllowMajorVersionUpgrade); err != nil {
			log.Errorf("invalid --allow-major-version-upgrade: %s", err.Error())
			return
		}
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneNewGlobalClusterIdentifier) > 0 {
		input.NewGlobalClusterIdentifier = aws.String(_neptuneNewGlobalClusterIdentifier)
	}

	if resp, err := client.ModifyGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Not supported.
func neptune_PromoteReadReplicaDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.PromoteReadReplicaDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}

	if resp, err := client.PromoteReadReplicaDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You might need to reboot your DB instance, usually for maintenance reasons. For
// example, if you make certain modifications, or if you change the DB parameter
// group associated with the DB instance, you must reboot the instance for the
// changes to take effect.
//
// Rebooting a DB instance restarts the database engine service. Rebooting a DB
// instance results in a momentary outage, during which the DB instance status is
// set to rebooting.
func neptune_RebootDBInstance(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RebootDBInstanceInput{
		// DBInstanceIdentifier: *string, // Required
	}

	if len(_neptuneDBInstanceIdentifier) > 0 {
		input.DBInstanceIdentifier = aws.String(_neptuneDBInstanceIdentifier)
	}
	if len(_neptuneForceFailover) > 0 {
		if err := assignInputField(input, "ForceFailover", _neptuneForceFailover); err != nil {
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

// Detaches a Neptune DB cluster from a Neptune global database. A secondary
// cluster becomes a normal standalone cluster with read-write capability instead
// of being read-only, and no longer receives data from a the primary cluster.
func neptune_RemoveFromGlobalCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RemoveFromGlobalClusterInput{
		// DbClusterIdentifier: *string, // Required
		// GlobalClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DbClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}

	if resp, err := client.RemoveFromGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates an Identity and Access Management (IAM) role from a DB cluster.
func neptune_RemoveRoleFromDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RemoveRoleFromDBClusterInput{
		// DBClusterIdentifier: *string, // Required
		// RoleArn: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneRoleArn) > 0 {
		input.RoleArn = aws.String(_neptuneRoleArn)
	}
	if len(_neptuneFeatureName) > 0 {
		input.FeatureName = aws.String(_neptuneFeatureName)
	}

	if resp, err := client.RemoveRoleFromDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a source identifier from an existing event notification subscription.
func neptune_RemoveSourceIdentifierFromSubscription(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RemoveSourceIdentifierFromSubscriptionInput{
		// SourceIdentifier: *string, // Required
		// SubscriptionName: *string, // Required
	}

	if len(_neptuneSourceIdentifier) > 0 {
		input.SourceIdentifier = aws.String(_neptuneSourceIdentifier)
	}
	if len(_neptuneSubscriptionName) > 0 {
		input.SubscriptionName = aws.String(_neptuneSubscriptionName)
	}

	if resp, err := client.RemoveSourceIdentifierFromSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes metadata tags from an Amazon Neptune resource.
func neptune_RemoveTagsFromResource(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RemoveTagsFromResourceInput{
		// ResourceName: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_neptuneResourceName) > 0 {
		input.ResourceName = aws.String(_neptuneResourceName)
	}
	if len(_neptuneTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _neptuneTagKeys...)
	}

	if resp, err := client.RemoveTagsFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the parameters of a DB cluster parameter group to the default value.
// To reset specific parameters submit a list of the following: ParameterName and
// ApplyMethod . To reset the entire DB cluster parameter group, specify the
// DBClusterParameterGroupName and ResetAllParameters parameters.
//
// When resetting the entire group, dynamic parameters are updated immediately and
// static parameters are set to pending-reboot to take effect on the next DB
// instance restart or RebootDBInstancerequest. You must call RebootDBInstance for every DB instance in your DB
// cluster that you want the updated static parameter to apply to.
func neptune_ResetDBClusterParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ResetDBClusterParameterGroupInput{
		// DBClusterParameterGroupName: *string, // Required
	}

	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _neptuneParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_neptuneResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _neptuneResetAllParameters); err != nil {
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

// Modifies the parameters of a DB parameter group to the engine/system default
// value. To reset specific parameters, provide a list of the following:
// ParameterName and ApplyMethod . To reset the entire DB parameter group, specify
// the DBParameterGroup name and ResetAllParameters parameters. When resetting the
// entire group, dynamic parameters are updated immediately and static parameters
// are set to pending-reboot to take effect on the next DB instance restart or
// RebootDBInstance request.
func neptune_ResetDBParameterGroup(cfg aws.Config, client *neptune.Client) {
	input := &neptune.ResetDBParameterGroupInput{
		// DBParameterGroupName: *string, // Required
	}

	if len(_neptuneDBParameterGroupName) > 0 {
		input.DBParameterGroupName = aws.String(_neptuneDBParameterGroupName)
	}
	if len(_neptuneParameters) > 0 {
		if err := assignInputField(input, "Parameters", _neptuneParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_neptuneResetAllParameters) > 0 {
		if err := assignInputField(input, "ResetAllParameters", _neptuneResetAllParameters); err != nil {
			log.Errorf("invalid --reset-all-parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ResetDBParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new DB cluster from a DB snapshot or DB cluster snapshot.
// If a DB snapshot is specified, the target DB cluster is created from the source
// DB snapshot with a default configuration and default security group.
//
// If a DB cluster snapshot is specified, the target DB cluster is created from
// the source DB cluster restore point with the same configuration as the original
// source DB cluster, except that the new DB cluster is created with the default
// security group.
func neptune_RestoreDBClusterFromSnapshot(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RestoreDBClusterFromSnapshotInput{
		// DBClusterIdentifier: *string, // Required
		// Engine: *string, // Required
		// SnapshotIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneEngine) > 0 {
		input.Engine = aws.String(_neptuneEngine)
	}
	if len(_neptuneSnapshotIdentifier) > 0 {
		input.SnapshotIdentifier = aws.String(_neptuneSnapshotIdentifier)
	}
	if len(_neptuneAvailabilityZones) > 0 {
		input.AvailabilityZones = append([]string(nil), _neptuneAvailabilityZones...)
	}
	if len(_neptuneCopyTagsToSnapshot) > 0 {
		if err := assignInputField(input, "CopyTagsToSnapshot", _neptuneCopyTagsToSnapshot); err != nil {
			log.Errorf("invalid --copy-tags-to-snapshot: %s", err.Error())
			return
		}
	}
	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneDatabaseName) > 0 {
		input.DatabaseName = aws.String(_neptuneDatabaseName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _neptuneEnableCloudwatchLogsExports...)
	}
	if len(_neptuneEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _neptuneEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_neptuneEngineVersion) > 0 {
		input.EngineVersion = aws.String(_neptuneEngineVersion)
	}
	if len(_neptuneKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_neptuneKmsKeyId)
	}
	if len(_neptuneOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_neptuneOptionGroupName)
	}
	if len(_neptunePort) > 0 {
		if err := assignInputField(input, "Port", _neptunePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_neptuneServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _neptuneServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageType) > 0 {
		input.StorageType = aws.String(_neptuneStorageType)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_neptuneVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptuneVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterFromSnapshot(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a DB cluster to an arbitrary point in time. Users can restore to any
// point in time before LatestRestorableTime for up to BackupRetentionPeriod days.
// The target DB cluster is created from the source DB cluster with the same
// configuration as the original DB cluster, except that the new DB cluster is
// created with the default DB security group.
//
// This action only restores the DB cluster, not the DB instances for that DB
// cluster. You must invoke the CreateDBInstanceaction to create DB instances for the restored DB
// cluster, specifying the identifier of the restored DB cluster in
// DBClusterIdentifier . You can create DB instances only after the
// RestoreDBClusterToPointInTime action has completed and the DB cluster is
// available.
func neptune_RestoreDBClusterToPointInTime(cfg aws.Config, client *neptune.Client) {
	input := &neptune.RestoreDBClusterToPointInTimeInput{
		// DBClusterIdentifier: *string, // Required
		// SourceDBClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}
	if len(_neptuneSourceDBClusterIdentifier) > 0 {
		input.SourceDBClusterIdentifier = aws.String(_neptuneSourceDBClusterIdentifier)
	}
	if len(_neptuneDBClusterParameterGroupName) > 0 {
		input.DBClusterParameterGroupName = aws.String(_neptuneDBClusterParameterGroupName)
	}
	if len(_neptuneDBSubnetGroupName) > 0 {
		input.DBSubnetGroupName = aws.String(_neptuneDBSubnetGroupName)
	}
	if len(_neptuneDeletionProtection) > 0 {
		if err := assignInputField(input, "DeletionProtection", _neptuneDeletionProtection); err != nil {
			log.Errorf("invalid --deletion-protection: %s", err.Error())
			return
		}
	}
	if len(_neptuneEnableCloudwatchLogsExports) > 0 {
		input.EnableCloudwatchLogsExports = append([]string(nil), _neptuneEnableCloudwatchLogsExports...)
	}
	if len(_neptuneEnableIAMDatabaseAuthentication) > 0 {
		if err := assignInputField(input, "EnableIAMDatabaseAuthentication", _neptuneEnableIAMDatabaseAuthentication); err != nil {
			log.Errorf("invalid --enable-iam-database-authentication: %s", err.Error())
			return
		}
	}
	if len(_neptuneKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_neptuneKmsKeyId)
	}
	if len(_neptuneOptionGroupName) > 0 {
		input.OptionGroupName = aws.String(_neptuneOptionGroupName)
	}
	if len(_neptunePort) > 0 {
		if err := assignInputField(input, "Port", _neptunePort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_neptuneRestoreToTime) > 0 {
		if err := assignInputField(input, "RestoreToTime", _neptuneRestoreToTime); err != nil {
			log.Errorf("invalid --restore-to-time: %s", err.Error())
			return
		}
	}
	if len(_neptuneRestoreType) > 0 {
		input.RestoreType = aws.String(_neptuneRestoreType)
	}
	if len(_neptuneServerlessV2ScalingConfiguration) > 0 {
		if err := assignInputField(input, "ServerlessV2ScalingConfiguration", _neptuneServerlessV2ScalingConfiguration); err != nil {
			log.Errorf("invalid --serverless-v2-scaling-configuration: %s", err.Error())
			return
		}
	}
	if len(_neptuneStorageType) > 0 {
		input.StorageType = aws.String(_neptuneStorageType)
	}
	if len(_neptuneTags) > 0 {
		if err := assignInputField(input, "Tags", _neptuneTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_neptuneUseLatestRestorableTime) > 0 {
		if err := assignInputField(input, "UseLatestRestorableTime", _neptuneUseLatestRestorableTime); err != nil {
			log.Errorf("invalid --use-latest-restorable-time: %s", err.Error())
			return
		}
	}
	if len(_neptuneVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _neptuneVpcSecurityGroupIds...)
	}

	if resp, err := client.RestoreDBClusterToPointInTime(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an Amazon Neptune DB cluster that was stopped using the Amazon console,
// the Amazon CLI stop-db-cluster command, or the StopDBCluster API.
func neptune_StartDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.StartDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}

	if resp, err := client.StartDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an Amazon Neptune DB cluster. When you stop a DB cluster, Neptune retains
// the DB cluster's metadata, including its endpoints and DB parameter groups.
//
// Neptune also retains the transaction logs so you can do a point-in-time restore
// if necessary.
func neptune_StopDBCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.StopDBClusterInput{
		// DBClusterIdentifier: *string, // Required
	}

	if len(_neptuneDBClusterIdentifier) > 0 {
		input.DBClusterIdentifier = aws.String(_neptuneDBClusterIdentifier)
	}

	if resp, err := client.StopDBCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Switches over the specified secondary DB cluster to be the new primary DB
// cluster in the global database cluster. Switchover operations were previously
// called "managed planned failovers."
//
// Promotes the specified secondary cluster to assume full read/write capabilities
// and demotes the current primary cluster to a secondary (read-only) cluster,
// maintaining the original replication topology. All secondary clusters are
// synchronized with the primary at the beginning of the process so the new primary
// continues operations for the global database without losing any data. Your
// database is unavailable for a short time while the primary and selected
// secondary clusters are assuming their new roles.
//
// This operation is intended for controlled environments, for operations such as
// "regional rotation" or to fall back to the original primary after a global
// database failover.
func neptune_SwitchoverGlobalCluster(cfg aws.Config, client *neptune.Client) {
	input := &neptune.SwitchoverGlobalClusterInput{
		// GlobalClusterIdentifier: *string, // Required
		// TargetDbClusterIdentifier: *string, // Required
	}

	if len(_neptuneGlobalClusterIdentifier) > 0 {
		input.GlobalClusterIdentifier = aws.String(_neptuneGlobalClusterIdentifier)
	}
	if len(_neptuneTargetDbClusterIdentifier) > 0 {
		input.TargetDbClusterIdentifier = aws.String(_neptuneTargetDbClusterIdentifier)
	}

	if resp, err := client.SwitchoverGlobalCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_neptuneCmd)
	_neptuneCmd.Flags().SortFlags = false

	_neptuneCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_neptuneCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_neptuneCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_neptuneCmd.Flags().StringVarP(&_neptuneAllocatedStorage, "allocated-storage", "", "", "Allocated Storage")
	_neptuneCmd.Flags().StringVarP(&_neptuneAllowDataLoss, "allow-data-loss", "", "", "Allow Data Loss")
	_neptuneCmd.Flags().StringVarP(&_neptuneAllowMajorVersionUpgrade, "allow-major-version-upgrade", "", "", "Allow Major Version Upgrade")
	_neptuneCmd.Flags().StringVarP(&_neptuneApplyAction, "apply-action", "", "", "Apply Action")
	_neptuneCmd.Flags().StringVarP(&_neptuneApplyImmediately, "apply-immediately", "", "", "Apply Immediately")
	_neptuneCmd.Flags().StringVarP(&_neptuneAttributeName, "attribute-name", "", "", "Attribute Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneAutoMinorVersionUpgrade, "auto-minor-version-upgrade", "", "", "Auto Minor Version Upgrade")
	_neptuneCmd.Flags().StringVarP(&_neptuneAvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneAvailabilityZones, "availability-zones", "", nil, "Availability Zones")
	_neptuneCmd.Flags().StringVarP(&_neptuneBackupRetentionPeriod, "backup-retention-period", "", "", "Backup Retention Period")
	_neptuneCmd.Flags().StringVarP(&_neptuneCACertificateIdentifier, "ca-certificate-identifier", "", "", "Ca Certificate Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneCharacterSetName, "character-set-name", "", "", "Character Set Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneCloudwatchLogsExportConfiguration, "cloudwatch-logs-export-configuration", "", "", "Cloudwatch Logs Export Configuration")
	_neptuneCmd.Flags().StringVarP(&_neptuneCopyTags, "copy-tags", "", "", "Copy Tags")
	_neptuneCmd.Flags().StringVarP(&_neptuneCopyTagsToSnapshot, "copy-tags-to-snapshot", "", "", "Copy Tags To Snapshot")
	_neptuneCmd.Flags().StringVarP(&_neptuneDatabaseName, "database-name", "", "", "Database Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBClusterEndpointIdentifier, "db-cluster-endpoint-identifier", "", "", "DB Cluster Endpoint Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBClusterIdentifier, "db-cluster-identifier", "", "", "DB Cluster Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBClusterParameterGroupName, "db-cluster-parameter-group-name", "", "", "DB Cluster Parameter Group Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBClusterSnapshotIdentifier, "db-cluster-snapshot-identifier", "", "", "DB Cluster Snapshot Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBInstanceClass, "db-instance-class", "", "", "DB Instance Class")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBInstanceIdentifier, "db-instance-identifier", "", "", "DB Instance Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBInstanceParameterGroupName, "db-instance-parameter-group-name", "", "", "DB Instance Parameter Group Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBName, "db-name", "", "", "DB Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBParameterGroupFamily, "db-parameter-group-family", "", "", "DB Parameter Group Family")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBParameterGroupName, "db-parameter-group-name", "", "", "DB Parameter Group Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBPortNumber, "db-port-number", "", "", "DB Port Number")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneDBSecurityGroups, "db-security-groups", "", nil, "DB Security Groups")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBSubnetGroupDescription, "db-subnet-group-description", "", "", "DB Subnet Group Description")
	_neptuneCmd.Flags().StringVarP(&_neptuneDBSubnetGroupName, "db-subnet-group-name", "", "", "DB Subnet Group Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDefaultOnly, "default-only", "", "", "Default Only")
	_neptuneCmd.Flags().StringVarP(&_neptuneDeletionProtection, "deletion-protection", "", "", "Deletion Protection")
	_neptuneCmd.Flags().StringVarP(&_neptuneDescription, "description", "", "", "Description")
	_neptuneCmd.Flags().StringVarP(&_neptuneDomain, "domain", "", "", "Domain")
	_neptuneCmd.Flags().StringVarP(&_neptuneDomainIAMRoleName, "domain-iam-role-name", "", "", "Domain IAM Role Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneDuration, "duration", "", "", "Duration")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneEnableCloudwatchLogsExports, "enable-cloudwatch-logs-exports", "", nil, "Enable Cloudwatch Logs Exports")
	_neptuneCmd.Flags().StringVarP(&_neptuneEnableIAMDatabaseAuthentication, "enable-iam-database-authentication", "", "", "Enable IAM Database Authentication")
	_neptuneCmd.Flags().StringVarP(&_neptuneEnablePerformanceInsights, "enable-performance-insights", "", "", "Enable Performance Insights")
	_neptuneCmd.Flags().StringVarP(&_neptuneEnabled, "enabled", "", "", "Enabled")
	_neptuneCmd.Flags().StringVarP(&_neptuneEndTime, "end-time", "", "", "End Time")
	_neptuneCmd.Flags().StringVarP(&_neptuneEndpointType, "endpoint-type", "", "", "Endpoint Type")
	_neptuneCmd.Flags().StringVarP(&_neptuneEngine, "engine", "", "", "Engine")
	_neptuneCmd.Flags().StringVarP(&_neptuneEngineVersion, "engine-version", "", "", "Engine Version")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneEventCategories, "event-categories", "", nil, "Event Categories")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneExcludedMembers, "excluded-members", "", nil, "Excluded Members")
	_neptuneCmd.Flags().StringVarP(&_neptuneFeatureName, "feature-name", "", "", "Feature Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneFilters, "filters", "", "", "Filters")
	_neptuneCmd.Flags().StringVarP(&_neptuneFinalDBSnapshotIdentifier, "final-db-snapshot-identifier", "", "", "Final DB Snapshot Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneForceFailover, "force-failover", "", "", "Force Failover")
	_neptuneCmd.Flags().StringVarP(&_neptuneGlobalClusterIdentifier, "global-cluster-identifier", "", "", "Global Cluster Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneIncludePublic, "include-public", "", "", "Include Public")
	_neptuneCmd.Flags().StringVarP(&_neptuneIncludeShared, "include-shared", "", "", "Include Shared")
	_neptuneCmd.Flags().StringVarP(&_neptuneIops, "iops", "", "", "IOPS")
	_neptuneCmd.Flags().StringVarP(&_neptuneKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_neptuneCmd.Flags().StringVarP(&_neptuneLicenseModel, "license-model", "", "", "License Model")
	_neptuneCmd.Flags().StringVarP(&_neptuneListSupportedCharacterSets, "list-supported-character-sets", "", "", "List Supported Character Sets")
	_neptuneCmd.Flags().StringVarP(&_neptuneListSupportedTimezones, "list-supported-timezones", "", "", "List Supported Timezones")
	_neptuneCmd.Flags().StringVarP(&_neptuneMarker, "marker", "", "", "Marker")
	_neptuneCmd.Flags().StringVarP(&_neptuneMasterUserPassword, "master-user-password", "", "", "Master User Password")
	_neptuneCmd.Flags().StringVarP(&_neptuneMasterUsername, "master-username", "", "", "Master Username")
	_neptuneCmd.Flags().StringVarP(&_neptuneMaxRecords, "max-records", "", "", "Max Records")
	_neptuneCmd.Flags().StringVarP(&_neptuneMonitoringInterval, "monitoring-interval", "", "", "Monitoring Interval")
	_neptuneCmd.Flags().StringVarP(&_neptuneMonitoringRoleArn, "monitoring-role-arn", "", "", "Monitoring Role ARN")
	_neptuneCmd.Flags().StringVarP(&_neptuneMultiAZ, "multi-az", "", "", "Multi AZ")
	_neptuneCmd.Flags().StringVarP(&_neptuneNewDBClusterIdentifier, "new-db-cluster-identifier", "", "", "New DB Cluster Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneNewDBInstanceIdentifier, "new-db-instance-identifier", "", "", "New DB Instance Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneNewGlobalClusterIdentifier, "new-global-cluster-identifier", "", "", "New Global Cluster Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneOptInType, "opt-in-type", "", "", "Opt In Type")
	_neptuneCmd.Flags().StringVarP(&_neptuneOptionGroupName, "option-group-name", "", "", "Option Group Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneParameters, "parameters", "", "", "Parameters")
	_neptuneCmd.Flags().StringVarP(&_neptunePerformanceInsightsKMSKeyId, "performance-insights-kms-key-id", "", "", "Performance Insights KMS Key ID")
	_neptuneCmd.Flags().StringVarP(&_neptunePort, "port", "", "", "Port")
	_neptuneCmd.Flags().StringVarP(&_neptunePreSignedUrl, "pre-signed-url", "", "", "Pre Signed URL")
	_neptuneCmd.Flags().StringVarP(&_neptunePreferredBackupWindow, "preferred-backup-window", "", "", "Preferred Backup Window")
	_neptuneCmd.Flags().StringVarP(&_neptunePreferredMaintenanceWindow, "preferred-maintenance-window", "", "", "Preferred Maintenance Window")
	_neptuneCmd.Flags().StringVarP(&_neptunePromotionTier, "promotion-tier", "", "", "Promotion Tier")
	_neptuneCmd.Flags().StringVarP(&_neptunePubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_neptuneCmd.Flags().StringVarP(&_neptuneReplicationSourceIdentifier, "replication-source-identifier", "", "", "Replication Source Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneResetAllParameters, "reset-all-parameters", "", "", "Reset All Parameters")
	_neptuneCmd.Flags().StringVarP(&_neptuneResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneResourceName, "resource-name", "", "", "Resource Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneRestoreToTime, "restore-to-time", "", "", "Restore To Time")
	_neptuneCmd.Flags().StringVarP(&_neptuneRestoreType, "restore-type", "", "", "Restore Type")
	_neptuneCmd.Flags().StringVarP(&_neptuneRoleArn, "role-arn", "", "", "Role ARN")
	_neptuneCmd.Flags().StringVarP(&_neptuneServerlessV2ScalingConfiguration, "serverless-v2-scaling-configuration", "", "", "Serverless V2 Scaling Configuration")
	_neptuneCmd.Flags().StringVarP(&_neptuneSkipFinalSnapshot, "skip-final-snapshot", "", "", "Skip Final Snapshot")
	_neptuneCmd.Flags().StringVarP(&_neptuneSnapshotIdentifier, "snapshot-identifier", "", "", "Snapshot Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneSnapshotType, "snapshot-type", "", "", "Snapshot Type")
	_neptuneCmd.Flags().StringVarP(&_neptuneSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_neptuneCmd.Flags().StringVarP(&_neptuneSource, "source", "", "", "Source")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceDBClusterIdentifier, "source-db-cluster-identifier", "", "", "Source DB Cluster Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceDBClusterParameterGroupIdentifier, "source-db-cluster-parameter-group-identifier", "", "", "Source DB Cluster Parameter Group Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceDBClusterSnapshotIdentifier, "source-db-cluster-snapshot-identifier", "", "", "Source DB Cluster Snapshot Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceDBParameterGroupIdentifier, "source-db-parameter-group-identifier", "", "", "Source DB Parameter Group Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceIdentifier, "source-identifier", "", "", "Source Identifier")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneSourceIds, "source-ids", "", nil, "Source Ids")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceRegion, "source-region", "", "", "Source Region")
	_neptuneCmd.Flags().StringVarP(&_neptuneSourceType, "source-type", "", "", "Source Type")
	_neptuneCmd.Flags().StringVarP(&_neptuneStartTime, "start-time", "", "", "Start Time")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneStaticMembers, "static-members", "", nil, "Static Members")
	_neptuneCmd.Flags().StringVarP(&_neptuneStorageEncrypted, "storage-encrypted", "", "", "Storage Encrypted")
	_neptuneCmd.Flags().StringVarP(&_neptuneStorageType, "storage-type", "", "", "Storage Type")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_neptuneCmd.Flags().StringVarP(&_neptuneSubscriptionName, "subscription-name", "", "", "Subscription Name")
	_neptuneCmd.Flags().StringVarP(&_neptuneSwitchover, "switchover", "", "", "Switchover")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneTagKeys, "tag-keys", "", nil, "Tag Keys")
	_neptuneCmd.Flags().StringVarP(&_neptuneTags, "tags", "", "", "Tags")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDbClusterIdentifier, "target-db-cluster-identifier", "", "", "Target DB Cluster Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDBClusterParameterGroupDescription, "target-db-cluster-parameter-group-description", "", "", "Target DB Cluster Parameter Group Description")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDBClusterParameterGroupIdentifier, "target-db-cluster-parameter-group-identifier", "", "", "Target DB Cluster Parameter Group Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDBClusterSnapshotIdentifier, "target-db-cluster-snapshot-identifier", "", "", "Target DB Cluster Snapshot Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDBInstanceIdentifier, "target-db-instance-identifier", "", "", "Target DB Instance Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDBParameterGroupDescription, "target-db-parameter-group-description", "", "", "Target DB Parameter Group Description")
	_neptuneCmd.Flags().StringVarP(&_neptuneTargetDBParameterGroupIdentifier, "target-db-parameter-group-identifier", "", "", "Target DB Parameter Group Identifier")
	_neptuneCmd.Flags().StringVarP(&_neptuneTdeCredentialArn, "tde-credential-arn", "", "", "Tde Credential ARN")
	_neptuneCmd.Flags().StringVarP(&_neptuneTdeCredentialPassword, "tde-credential-password", "", "", "Tde Credential Password")
	_neptuneCmd.Flags().StringVarP(&_neptuneTimezone, "timezone", "", "", "Timezone")
	_neptuneCmd.Flags().StringVarP(&_neptuneUseLatestRestorableTime, "use-latest-restorable-time", "", "", "Use Latest Restorable Time")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneValuesToAdd, "values-to-add", "", nil, "Values To Add")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneValuesToRemove, "values-to-remove", "", nil, "Values To Remove")
	_neptuneCmd.Flags().StringVarP(&_neptuneVpc, "vpc", "", "", "VPC")
	_neptuneCmd.Flags().StringSliceVarP(&_neptuneVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")

	_neptuneCmd.Flags().BoolVarP(&_neptuneAddRoleToDBCluster, "add-role-to-db-cluster", "", false, "Add Role To DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneAddSourceIdentifierToSubscription, "add-source-identifier-to-subscription", "", false, "Add Source Identifier To Subscription")
	_neptuneCmd.Flags().BoolVarP(&_neptuneAddTagsToResource, "add-tags-to-resource", "", false, "Add Tags To Resource")
	_neptuneCmd.Flags().BoolVarP(&_neptuneApplyPendingMaintenanceAction, "apply-pending-maintenance-action", "", false, "Apply Pending Maintenance Action")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCopyDBClusterParameterGroup, "copy-db-cluster-parameter-group", "", false, "Copy DB Cluster Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCopyDBClusterSnapshot, "copy-db-cluster-snapshot", "", false, "Copy DB Cluster Snapshot")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCopyDBParameterGroup, "copy-db-parameter-group", "", false, "Copy DB Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBCluster, "create-db-cluster", "", false, "Create DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBClusterEndpoint, "create-db-cluster-endpoint", "", false, "Create DB Cluster Endpoint")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBClusterParameterGroup, "create-db-cluster-parameter-group", "", false, "Create DB Cluster Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBClusterSnapshot, "create-db-cluster-snapshot", "", false, "Create DB Cluster Snapshot")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBInstance, "create-db-instance", "", false, "Create DB Instance")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBParameterGroup, "create-db-parameter-group", "", false, "Create DB Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateDBSubnetGroup, "create-db-subnet-group", "", false, "Create DB Subnet Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateEventSubscription, "create-event-subscription", "", false, "Create Event Subscription")
	_neptuneCmd.Flags().BoolVarP(&_neptuneCreateGlobalCluster, "create-global-cluster", "", false, "Create Global Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBCluster, "delete-db-cluster", "", false, "Delete DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBClusterEndpoint, "delete-db-cluster-endpoint", "", false, "Delete DB Cluster Endpoint")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBClusterParameterGroup, "delete-db-cluster-parameter-group", "", false, "Delete DB Cluster Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBClusterSnapshot, "delete-db-cluster-snapshot", "", false, "Delete DB Cluster Snapshot")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBInstance, "delete-db-instance", "", false, "Delete DB Instance")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBParameterGroup, "delete-db-parameter-group", "", false, "Delete DB Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteDBSubnetGroup, "delete-db-subnet-group", "", false, "Delete DB Subnet Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteEventSubscription, "delete-event-subscription", "", false, "Delete Event Subscription")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDeleteGlobalCluster, "delete-global-cluster", "", false, "Delete Global Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBClusterEndpoints, "describe-db-cluster-endpoints", "", false, "Describe DB Cluster Endpoints")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBClusterParameterGroups, "describe-db-cluster-parameter-groups", "", false, "Describe DB Cluster Parameter Groups")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBClusterParameters, "describe-db-cluster-parameters", "", false, "Describe DB Cluster Parameters")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBClusterSnapshotAttributes, "describe-db-cluster-snapshot-attributes", "", false, "Describe DB Cluster Snapshot Attributes")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBClusterSnapshots, "describe-db-cluster-snapshots", "", false, "Describe DB Cluster Snapshots")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBClusters, "describe-db-clusters", "", false, "Describe DB Clusters")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBEngineVersions, "describe-db-engine-versions", "", false, "Describe DB Engine Versions")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBInstances, "describe-db-instances", "", false, "Describe DB Instances")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBParameterGroups, "describe-db-parameter-groups", "", false, "Describe DB Parameter Groups")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBParameters, "describe-db-parameters", "", false, "Describe DB Parameters")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeDBSubnetGroups, "describe-db-subnet-groups", "", false, "Describe DB Subnet Groups")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeEngineDefaultClusterParameters, "describe-engine-default-cluster-parameters", "", false, "Describe Engine Default Cluster Parameters")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeEngineDefaultParameters, "describe-engine-default-parameters", "", false, "Describe Engine Default Parameters")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeEventCategories, "describe-event-categories", "", false, "Describe Event Categories")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeEventSubscriptions, "describe-event-subscriptions", "", false, "Describe Event Subscriptions")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeEvents, "describe-events", "", false, "Describe Events")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeGlobalClusters, "describe-global-clusters", "", false, "Describe Global Clusters")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeOrderableDBInstanceOptions, "describe-orderable-db-instance-options", "", false, "Describe Orderable DB Instance Options")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribePendingMaintenanceActions, "describe-pending-maintenance-actions", "", false, "Describe Pending Maintenance Actions")
	_neptuneCmd.Flags().BoolVarP(&_neptuneDescribeValidDBInstanceModifications, "describe-valid-db-instance-modifications", "", false, "Describe Valid DB Instance Modifications")
	_neptuneCmd.Flags().BoolVarP(&_neptuneFailoverDBCluster, "failover-db-cluster", "", false, "Failover DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneFailoverGlobalCluster, "failover-global-cluster", "", false, "Failover Global Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBCluster, "modify-db-cluster", "", false, "Modify DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBClusterEndpoint, "modify-db-cluster-endpoint", "", false, "Modify DB Cluster Endpoint")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBClusterParameterGroup, "modify-db-cluster-parameter-group", "", false, "Modify DB Cluster Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBClusterSnapshotAttribute, "modify-db-cluster-snapshot-attribute", "", false, "Modify DB Cluster Snapshot Attribute")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBInstance, "modify-db-instance", "", false, "Modify DB Instance")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBParameterGroup, "modify-db-parameter-group", "", false, "Modify DB Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyDBSubnetGroup, "modify-db-subnet-group", "", false, "Modify DB Subnet Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyEventSubscription, "modify-event-subscription", "", false, "Modify Event Subscription")
	_neptuneCmd.Flags().BoolVarP(&_neptuneModifyGlobalCluster, "modify-global-cluster", "", false, "Modify Global Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptunePromoteReadReplicaDBCluster, "promote-read-replica-db-cluster", "", false, "Promote Read Replica DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRebootDBInstance, "reboot-db-instance", "", false, "Reboot DB Instance")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRemoveFromGlobalCluster, "remove-from-global-cluster", "", false, "Remove From Global Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRemoveRoleFromDBCluster, "remove-role-from-db-cluster", "", false, "Remove Role From DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRemoveSourceIdentifierFromSubscription, "remove-source-identifier-from-subscription", "", false, "Remove Source Identifier From Subscription")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRemoveTagsFromResource, "remove-tags-from-resource", "", false, "Remove Tags From Resource")
	_neptuneCmd.Flags().BoolVarP(&_neptuneResetDBClusterParameterGroup, "reset-db-cluster-parameter-group", "", false, "Reset DB Cluster Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneResetDBParameterGroup, "reset-db-parameter-group", "", false, "Reset DB Parameter Group")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRestoreDBClusterFromSnapshot, "restore-db-cluster-from-snapshot", "", false, "Restore DB Cluster From Snapshot")
	_neptuneCmd.Flags().BoolVarP(&_neptuneRestoreDBClusterToPointInTime, "restore-db-cluster-to-point-in-time", "", false, "Restore DB Cluster To Point In Time")
	_neptuneCmd.Flags().BoolVarP(&_neptuneStartDBCluster, "start-db-cluster", "", false, "Start DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneStopDBCluster, "stop-db-cluster", "", false, "Stop DB Cluster")
	_neptuneCmd.Flags().BoolVarP(&_neptuneSwitchoverGlobalCluster, "switchover-global-cluster", "", false, "Switchover Global Cluster")

}
