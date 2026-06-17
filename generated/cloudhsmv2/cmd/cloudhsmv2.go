package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudhsmv2"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cloudhsmv2Cmd represents the cloudhsmv2 command
var _cloudhsmv2Cmd = &cobra.Command{
	Use:   "cloudhsmv2",
	Short: "AWS cloudhsmv2 CLI",
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
		client := cloudhsmv2.NewFromConfig(cfg)
		if _cloudhsmv2CopyBackupToRegion {
			cloudhsmv2_CopyBackupToRegion(cfg, client)
			return
		}
		if _cloudhsmv2CreateCluster {
			cloudhsmv2_CreateCluster(cfg, client)
			return
		}
		if _cloudhsmv2CreateHsm {
			cloudhsmv2_CreateHsm(cfg, client)
			return
		}
		if _cloudhsmv2DeleteBackup {
			cloudhsmv2_DeleteBackup(cfg, client)
			return
		}
		if _cloudhsmv2DeleteCluster {
			cloudhsmv2_DeleteCluster(cfg, client)
			return
		}
		if _cloudhsmv2DeleteHsm {
			cloudhsmv2_DeleteHsm(cfg, client)
			return
		}
		if _cloudhsmv2DeleteResourcePolicy {
			cloudhsmv2_DeleteResourcePolicy(cfg, client)
			return
		}
		if _cloudhsmv2DescribeBackups {
			cloudhsmv2_DescribeBackups(cfg, client)
			return
		}
		if _cloudhsmv2DescribeClusters {
			cloudhsmv2_DescribeClusters(cfg, client)
			return
		}
		if _cloudhsmv2GetResourcePolicy {
			cloudhsmv2_GetResourcePolicy(cfg, client)
			return
		}
		if _cloudhsmv2InitializeCluster {
			cloudhsmv2_InitializeCluster(cfg, client)
			return
		}
		if _cloudhsmv2ListTags {
			cloudhsmv2_ListTags(cfg, client)
			return
		}
		if _cloudhsmv2ModifyBackupAttributes {
			cloudhsmv2_ModifyBackupAttributes(cfg, client)
			return
		}
		if _cloudhsmv2ModifyCluster {
			cloudhsmv2_ModifyCluster(cfg, client)
			return
		}
		if _cloudhsmv2PutResourcePolicy {
			cloudhsmv2_PutResourcePolicy(cfg, client)
			return
		}
		if _cloudhsmv2RestoreBackup {
			cloudhsmv2_RestoreBackup(cfg, client)
			return
		}
		if _cloudhsmv2TagResource {
			cloudhsmv2_TagResource(cfg, client)
			return
		}
		if _cloudhsmv2UntagResource {
			cloudhsmv2_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_cloudhsmv2CopyBackupToRegion     bool
	_cloudhsmv2CreateCluster          bool
	_cloudhsmv2CreateHsm              bool
	_cloudhsmv2DeleteBackup           bool
	_cloudhsmv2DeleteCluster          bool
	_cloudhsmv2DeleteHsm              bool
	_cloudhsmv2DeleteResourcePolicy   bool
	_cloudhsmv2DescribeBackups        bool
	_cloudhsmv2DescribeClusters       bool
	_cloudhsmv2GetResourcePolicy      bool
	_cloudhsmv2InitializeCluster      bool
	_cloudhsmv2ListTags               bool
	_cloudhsmv2ModifyBackupAttributes bool
	_cloudhsmv2ModifyCluster          bool
	_cloudhsmv2PutResourcePolicy      bool
	_cloudhsmv2RestoreBackup          bool
	_cloudhsmv2TagResource            bool
	_cloudhsmv2UntagResource          bool

	_cloudhsmv2AvailabilityZone      string
	_cloudhsmv2BackupId              string
	_cloudhsmv2BackupRetentionPolicy string
	_cloudhsmv2ClusterId             string
	_cloudhsmv2DestinationRegion     string
	_cloudhsmv2EniId                 string
	_cloudhsmv2EniIp                 string
	_cloudhsmv2Filters               string
	_cloudhsmv2HsmId                 string
	_cloudhsmv2HsmType               string
	_cloudhsmv2IpAddress             string
	_cloudhsmv2MaxResults            string
	_cloudhsmv2Mode                  string
	_cloudhsmv2NetworkType           string
	_cloudhsmv2NeverExpires          string
	_cloudhsmv2NextToken             string
	_cloudhsmv2Policy                string
	_cloudhsmv2ResourceArn           string
	_cloudhsmv2ResourceId            string
	_cloudhsmv2Shared                string
	_cloudhsmv2SignedCert            string
	_cloudhsmv2SortAscending         string
	_cloudhsmv2SourceBackupId        string
	_cloudhsmv2SubnetIds             []string
	_cloudhsmv2TagKeyList            []string
	_cloudhsmv2TagList               string
	_cloudhsmv2TrustAnchor           string
)

// Copy an CloudHSM cluster backup to a different region.
// Cross-account use: No. You cannot perform this operation on an CloudHSM backup
// in a different Amazon Web Services account.
func cloudhsmv2_CopyBackupToRegion(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.CopyBackupToRegionInput{
		// BackupId: *string, // Required
		// DestinationRegion: *string, // Required
	}

	if len(_cloudhsmv2BackupId) > 0 {
		input.BackupId = aws.String(_cloudhsmv2BackupId)
	}
	if len(_cloudhsmv2DestinationRegion) > 0 {
		input.DestinationRegion = aws.String(_cloudhsmv2DestinationRegion)
	}
	if len(_cloudhsmv2TagList) > 0 {
		if err := assignInputField(input, "TagList", _cloudhsmv2TagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CopyBackupToRegion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new CloudHSM cluster.
// Cross-account use: Yes. To perform this operation with an CloudHSM backup in a
// different AWS account, specify the full backup ARN in the value of the
// SourceBackupId parameter.
func cloudhsmv2_CreateCluster(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.CreateClusterInput{
		// HsmType: *string, // Required
		// SubnetIds: []string, // Required
	}

	if len(_cloudhsmv2HsmType) > 0 {
		input.HsmType = aws.String(_cloudhsmv2HsmType)
	}
	if len(_cloudhsmv2SubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _cloudhsmv2SubnetIds...)
	}
	if len(_cloudhsmv2BackupRetentionPolicy) > 0 {
		if err := assignInputField(input, "BackupRetentionPolicy", _cloudhsmv2BackupRetentionPolicy); err != nil {
			log.Errorf("invalid --backup-retention-policy: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2Mode) > 0 {
		if err := assignInputField(input, "Mode", _cloudhsmv2Mode); err != nil {
			log.Errorf("invalid --mode: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2NetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _cloudhsmv2NetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2SourceBackupId) > 0 {
		input.SourceBackupId = aws.String(_cloudhsmv2SourceBackupId)
	}
	if len(_cloudhsmv2TagList) > 0 {
		if err := assignInputField(input, "TagList", _cloudhsmv2TagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new hardware security module (HSM) in the specified CloudHSM cluster.
// Cross-account use: No. You cannot perform this operation on an CloudHSM cluster
// in a different Amazon Web Service account.
func cloudhsmv2_CreateHsm(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.CreateHsmInput{
		// AvailabilityZone: *string, // Required
		// ClusterId: *string, // Required
	}

	if len(_cloudhsmv2AvailabilityZone) > 0 {
		input.AvailabilityZone = aws.String(_cloudhsmv2AvailabilityZone)
	}
	if len(_cloudhsmv2ClusterId) > 0 {
		input.ClusterId = aws.String(_cloudhsmv2ClusterId)
	}
	if len(_cloudhsmv2IpAddress) > 0 {
		input.IpAddress = aws.String(_cloudhsmv2IpAddress)
	}

	if resp, err := client.CreateHsm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified CloudHSM backup. A backup can be restored up to 7 days
// after the DeleteBackup request is made. For more information on restoring a
// backup, see RestoreBackup.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM backup
// in a different Amazon Web Services account.
func cloudhsmv2_DeleteBackup(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.DeleteBackupInput{
		// BackupId: *string, // Required
	}

	if len(_cloudhsmv2BackupId) > 0 {
		input.BackupId = aws.String(_cloudhsmv2BackupId)
	}

	if resp, err := client.DeleteBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified CloudHSM cluster. Before you can delete a cluster, you
// must delete all HSMs in the cluster. To see if the cluster contains any HSMs,
// use DescribeClusters. To delete an HSM, use DeleteHsm.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM cluster
// in a different Amazon Web Services account.
func cloudhsmv2_DeleteCluster(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.DeleteClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_cloudhsmv2ClusterId) > 0 {
		input.ClusterId = aws.String(_cloudhsmv2ClusterId)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified HSM. To specify an HSM, you can use its identifier (ID),
// the IP address of the HSM's elastic network interface (ENI), or the ID of the
// HSM's ENI. You need to specify only one of these values. To find these values,
// use DescribeClusters.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM hsm in
// a different Amazon Web Services account.
func cloudhsmv2_DeleteHsm(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.DeleteHsmInput{
		// ClusterId: *string, // Required
	}

	if len(_cloudhsmv2ClusterId) > 0 {
		input.ClusterId = aws.String(_cloudhsmv2ClusterId)
	}
	if len(_cloudhsmv2EniId) > 0 {
		input.EniId = aws.String(_cloudhsmv2EniId)
	}
	if len(_cloudhsmv2EniIp) > 0 {
		input.EniIp = aws.String(_cloudhsmv2EniIp)
	}
	if len(_cloudhsmv2HsmId) > 0 {
		input.HsmId = aws.String(_cloudhsmv2HsmId)
	}

	if resp, err := client.DeleteHsm(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an CloudHSM resource policy. Deleting a resource policy will result in
// the resource being unshared and removed from any RAM resource shares. Deleting
// the resource policy attached to a backup will not impact any clusters created
// from that backup.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
func cloudhsmv2_DeleteResourcePolicy(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.DeleteResourcePolicyInput{}

	if len(_cloudhsmv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudhsmv2ResourceArn)
	}

	if resp, err := client.DeleteResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about backups of CloudHSM clusters. Lists either the backups
// you own or the backups shared with you when the Shared parameter is true.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the backups. When the response contains only a subset of
// backups, it includes a NextToken value. Use this value in a subsequent
// DescribeBackups request to get more backups. When you receive a response with no
// NextToken (or an empty or null value), that means there are no more backups to
// get.
//
// Cross-account use: Yes. Customers can describe backups in other Amazon Web
// Services accounts that are shared with them.
func cloudhsmv2_DescribeBackups(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.DescribeBackupsInput{}

	if len(_cloudhsmv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _cloudhsmv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudhsmv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2NextToken) > 0 {
		input.NextToken = aws.String(_cloudhsmv2NextToken)
	}
	if len(_cloudhsmv2Shared) > 0 {
		if err := assignInputField(input, "Shared", _cloudhsmv2Shared); err != nil {
			log.Errorf("invalid --shared: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2SortAscending) > 0 {
		if err := assignInputField(input, "SortAscending", _cloudhsmv2SortAscending); err != nil {
			log.Errorf("invalid --sort-ascending: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeBackups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudhsmv2.DescribeBackupsOutput
	p := cloudhsmv2.NewDescribeBackupsPaginator(client, input)
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

// Gets information about CloudHSM clusters.
// This is a paginated operation, which means that each response might contain
// only a subset of all the clusters. When the response contains only a subset of
// clusters, it includes a NextToken value. Use this value in a subsequent
// DescribeClusters request to get more clusters. When you receive a response with
// no NextToken (or an empty or null value), that means there are no more clusters
// to get.
//
// Cross-account use: No. You cannot perform this operation on CloudHSM clusters
// in a different Amazon Web Services account.
func cloudhsmv2_DescribeClusters(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.DescribeClustersInput{}

	if len(_cloudhsmv2Filters) > 0 {
		if err := assignInputField(input, "Filters", _cloudhsmv2Filters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudhsmv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2NextToken) > 0 {
		input.NextToken = aws.String(_cloudhsmv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudhsmv2.DescribeClustersOutput
	p := cloudhsmv2.NewDescribeClustersPaginator(client, input)
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

// Retrieves the resource policy document attached to a given resource.
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
func cloudhsmv2_GetResourcePolicy(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.GetResourcePolicyInput{}

	if len(_cloudhsmv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudhsmv2ResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Claims an CloudHSM cluster by submitting the cluster certificate issued by your
// issuing certificate authority (CA) and the CA's root certificate. Before you can
// claim a cluster, you must sign the cluster's certificate signing request (CSR)
// with your issuing CA. To get the cluster's CSR, use DescribeClusters.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM cluster
// in a different Amazon Web Services account.
func cloudhsmv2_InitializeCluster(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.InitializeClusterInput{
		// ClusterId: *string, // Required
		// SignedCert: *string, // Required
		// TrustAnchor: *string, // Required
	}

	if len(_cloudhsmv2ClusterId) > 0 {
		input.ClusterId = aws.String(_cloudhsmv2ClusterId)
	}
	if len(_cloudhsmv2SignedCert) > 0 {
		input.SignedCert = aws.String(_cloudhsmv2SignedCert)
	}
	if len(_cloudhsmv2TrustAnchor) > 0 {
		input.TrustAnchor = aws.String(_cloudhsmv2TrustAnchor)
	}

	if resp, err := client.InitializeCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a list of tags for the specified CloudHSM cluster.
// This is a paginated operation, which means that each response might contain
// only a subset of all the tags. When the response contains only a subset of tags,
// it includes a NextToken value. Use this value in a subsequent ListTags request
// to get more tags. When you receive a response with no NextToken (or an empty or
// null value), that means there are no more tags to get.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
func cloudhsmv2_ListTags(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.ListTagsInput{
		// ResourceId: *string, // Required
	}

	if len(_cloudhsmv2ResourceId) > 0 {
		input.ResourceId = aws.String(_cloudhsmv2ResourceId)
	}
	if len(_cloudhsmv2MaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _cloudhsmv2MaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2NextToken) > 0 {
		input.NextToken = aws.String(_cloudhsmv2NextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*cloudhsmv2.ListTagsOutput
	p := cloudhsmv2.NewListTagsPaginator(client, input)
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

// Modifies attributes for CloudHSM backup.
// Cross-account use: No. You cannot perform this operation on an CloudHSM backup
// in a different Amazon Web Services account.
func cloudhsmv2_ModifyBackupAttributes(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.ModifyBackupAttributesInput{
		// BackupId: *string, // Required
		// NeverExpires: *bool, // Required
	}

	if len(_cloudhsmv2BackupId) > 0 {
		input.BackupId = aws.String(_cloudhsmv2BackupId)
	}
	if len(_cloudhsmv2NeverExpires) > 0 {
		if err := assignInputField(input, "NeverExpires", _cloudhsmv2NeverExpires); err != nil {
			log.Errorf("invalid --never-expires: %s", err.Error())
			return
		}
	}

	if resp, err := client.ModifyBackupAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies CloudHSM cluster.
// Cross-account use: No. You cannot perform this operation on an CloudHSM cluster
// in a different Amazon Web Services account.
func cloudhsmv2_ModifyCluster(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.ModifyClusterInput{
		// ClusterId: *string, // Required
	}

	if len(_cloudhsmv2ClusterId) > 0 {
		input.ClusterId = aws.String(_cloudhsmv2ClusterId)
	}
	if len(_cloudhsmv2BackupRetentionPolicy) > 0 {
		if err := assignInputField(input, "BackupRetentionPolicy", _cloudhsmv2BackupRetentionPolicy); err != nil {
			log.Errorf("invalid --backup-retention-policy: %s", err.Error())
			return
		}
	}
	if len(_cloudhsmv2HsmType) > 0 {
		input.HsmType = aws.String(_cloudhsmv2HsmType)
	}

	if resp, err := client.ModifyCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates an CloudHSM resource policy. A resource policy helps you to
// define the IAM entity (for example, an Amazon Web Services account) that can
// manage your CloudHSM resources. The following resources support CloudHSM
// resource policies:
//
// - Backup - The resource policy allows you to describe the backup and restore
// a cluster from the backup in another Amazon Web Services account.
//
// In order to share a backup, it must be in a 'READY' state and you must own it.
//
// While you can share a backup using the CloudHSM PutResourcePolicy operation, we
// recommend using Resource Access Manager (RAM) instead. Using RAM provides
// multiple benefits as it creates the policy for you, allows multiple resources to
// be shared at one time, and increases the discoverability of shared resources. If
// you use PutResourcePolicy and want consumers to be able to describe the backups
// you share with them, you must promote the backup to a standard RAM Resource
// Share using the RAM PromoteResourceShareCreatedFromPolicy API operation.
//
// For more information, see [Working with shared backups] in the CloudHSM User Guide
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
//
// [Working with shared backups]: https://docs.aws.amazon.com/cloudhsm/latest/userguide/sharing.html
func cloudhsmv2_PutResourcePolicy(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.PutResourcePolicyInput{}

	if len(_cloudhsmv2Policy) > 0 {
		input.Policy = aws.String(_cloudhsmv2Policy)
	}
	if len(_cloudhsmv2ResourceArn) > 0 {
		input.ResourceArn = aws.String(_cloudhsmv2ResourceArn)
	}

	if resp, err := client.PutResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Restores a specified CloudHSM backup that is in the PENDING_DELETION state. For
// more information on deleting a backup, see DeleteBackup.
//
// Cross-account use: No. You cannot perform this operation on an CloudHSM backup
// in a different Amazon Web Services account.
func cloudhsmv2_RestoreBackup(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.RestoreBackupInput{
		// BackupId: *string, // Required
	}

	if len(_cloudhsmv2BackupId) > 0 {
		input.BackupId = aws.String(_cloudhsmv2BackupId)
	}

	if resp, err := client.RestoreBackup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites one or more tags for the specified CloudHSM cluster.
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
func cloudhsmv2_TagResource(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.TagResourceInput{
		// ResourceId: *string, // Required
		// TagList: []types.Tag, // Required
	}

	if len(_cloudhsmv2ResourceId) > 0 {
		input.ResourceId = aws.String(_cloudhsmv2ResourceId)
	}
	if len(_cloudhsmv2TagList) > 0 {
		if err := assignInputField(input, "TagList", _cloudhsmv2TagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
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

// Removes the specified tag or tags from the specified CloudHSM cluster.
// Cross-account use: No. You cannot perform this operation on an CloudHSM
// resource in a different Amazon Web Services account.
func cloudhsmv2_UntagResource(cfg aws.Config, client *cloudhsmv2.Client) {
	input := &cloudhsmv2.UntagResourceInput{
		// ResourceId: *string, // Required
		// TagKeyList: []string, // Required
	}

	if len(_cloudhsmv2ResourceId) > 0 {
		input.ResourceId = aws.String(_cloudhsmv2ResourceId)
	}
	if len(_cloudhsmv2TagKeyList) > 0 {
		input.TagKeyList = append([]string(nil), _cloudhsmv2TagKeyList...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_cloudhsmv2Cmd)
	_cloudhsmv2Cmd.Flags().SortFlags = false

	_cloudhsmv2Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_cloudhsmv2Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_cloudhsmv2Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2AvailabilityZone, "availability-zone", "", "", "Availability Zone")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2BackupId, "backup-id", "", "", "Backup ID")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2BackupRetentionPolicy, "backup-retention-policy", "", "", "Backup Retention Policy")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2ClusterId, "cluster-id", "", "", "Cluster ID")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2DestinationRegion, "destination-region", "", "", "Destination Region")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2EniId, "eni-id", "", "", "Eni ID")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2EniIp, "eni-ip", "", "", "Eni IP")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2Filters, "filters", "", "", "Filters")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2HsmId, "hsm-id", "", "", "Hsm ID")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2HsmType, "hsm-type", "", "", "Hsm Type")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2IpAddress, "ip-address", "", "", "IP Address")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2MaxResults, "max-results", "", "", "Max Results")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2Mode, "mode", "", "", "Mode")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2NetworkType, "network-type", "", "", "Network Type")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2NeverExpires, "never-expires", "", "", "Never Expires")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2NextToken, "next-token", "", "", "Next Token")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2Policy, "policy", "", "", "Policy")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2ResourceArn, "resource-arn", "", "", "Resource ARN")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2ResourceId, "resource-id", "", "", "Resource ID")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2Shared, "shared", "", "", "Shared")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2SignedCert, "signed-cert", "", "", "Signed Cert")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2SortAscending, "sort-ascending", "", "", "Sort Ascending")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2SourceBackupId, "source-backup-id", "", "", "Source Backup ID")
	_cloudhsmv2Cmd.Flags().StringSliceVarP(&_cloudhsmv2SubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_cloudhsmv2Cmd.Flags().StringSliceVarP(&_cloudhsmv2TagKeyList, "tag-key-list", "", nil, "Tag Key List")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2TagList, "tag-list", "", "", "Tag List")
	_cloudhsmv2Cmd.Flags().StringVarP(&_cloudhsmv2TrustAnchor, "trust-anchor", "", "", "Trust Anchor")

	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2CopyBackupToRegion, "copy-backup-to-region", "", false, "Copy Backup To Region")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2CreateCluster, "create-cluster", "", false, "Create Cluster")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2CreateHsm, "create-hsm", "", false, "Create Hsm")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2DeleteBackup, "delete-backup", "", false, "Delete Backup")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2DeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2DeleteHsm, "delete-hsm", "", false, "Delete Hsm")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2DeleteResourcePolicy, "delete-resource-policy", "", false, "Delete Resource Policy")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2DescribeBackups, "describe-backups", "", false, "Describe Backups")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2DescribeClusters, "describe-clusters", "", false, "Describe Clusters")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2GetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2InitializeCluster, "initialize-cluster", "", false, "Initialize Cluster")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2ListTags, "list-tags", "", false, "List Tags")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2ModifyBackupAttributes, "modify-backup-attributes", "", false, "Modify Backup Attributes")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2ModifyCluster, "modify-cluster", "", false, "Modify Cluster")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2PutResourcePolicy, "put-resource-policy", "", false, "Put Resource Policy")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2RestoreBackup, "restore-backup", "", false, "Restore Backup")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2TagResource, "tag-resource", "", false, "Tag Resource")
	_cloudhsmv2Cmd.Flags().BoolVarP(&_cloudhsmv2UntagResource, "untag-resource", "", false, "Untag Resource")

}
