package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/timestreaminfluxdb"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// timestreaminfluxdbCmd represents the timestreaminfluxdb command
var _timestreaminfluxdbCmd = &cobra.Command{
	Use:   "timestreaminfluxdb",
	Short: "AWS timestreaminfluxdb CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := timestreaminfluxdb.NewFromConfig(cfg)
		if _timestreaminfluxdbCreateDbCluster {
			timestreaminfluxdb_CreateDbCluster(cfg, client)
			return
		}
		if _timestreaminfluxdbCreateDbInstance {
			timestreaminfluxdb_CreateDbInstance(cfg, client)
			return
		}
		if _timestreaminfluxdbCreateDbParameterGroup {
			timestreaminfluxdb_CreateDbParameterGroup(cfg, client)
			return
		}
		if _timestreaminfluxdbDeleteDbCluster {
			timestreaminfluxdb_DeleteDbCluster(cfg, client)
			return
		}
		if _timestreaminfluxdbDeleteDbInstance {
			timestreaminfluxdb_DeleteDbInstance(cfg, client)
			return
		}
		if _timestreaminfluxdbGetDbCluster {
			timestreaminfluxdb_GetDbCluster(cfg, client)
			return
		}
		if _timestreaminfluxdbGetDbInstance {
			timestreaminfluxdb_GetDbInstance(cfg, client)
			return
		}
		if _timestreaminfluxdbGetDbParameterGroup {
			timestreaminfluxdb_GetDbParameterGroup(cfg, client)
			return
		}
		if _timestreaminfluxdbListDbClusters {
			timestreaminfluxdb_ListDbClusters(cfg, client)
			return
		}
		if _timestreaminfluxdbListDbInstances {
			timestreaminfluxdb_ListDbInstances(cfg, client)
			return
		}
		if _timestreaminfluxdbListDbInstancesForCluster {
			timestreaminfluxdb_ListDbInstancesForCluster(cfg, client)
			return
		}
		if _timestreaminfluxdbListDbParameterGroups {
			timestreaminfluxdb_ListDbParameterGroups(cfg, client)
			return
		}
		if _timestreaminfluxdbListTagsForResource {
			timestreaminfluxdb_ListTagsForResource(cfg, client)
			return
		}
		if _timestreaminfluxdbRebootDbCluster {
			timestreaminfluxdb_RebootDbCluster(cfg, client)
			return
		}
		if _timestreaminfluxdbRebootDbInstance {
			timestreaminfluxdb_RebootDbInstance(cfg, client)
			return
		}
		if _timestreaminfluxdbTagResource {
			timestreaminfluxdb_TagResource(cfg, client)
			return
		}
		if _timestreaminfluxdbUntagResource {
			timestreaminfluxdb_UntagResource(cfg, client)
			return
		}
		if _timestreaminfluxdbUpdateDbCluster {
			timestreaminfluxdb_UpdateDbCluster(cfg, client)
			return
		}
		if _timestreaminfluxdbUpdateDbInstance {
			timestreaminfluxdb_UpdateDbInstance(cfg, client)
			return
		}

	},
}

var (
	_timestreaminfluxdbCreateDbCluster           bool
	_timestreaminfluxdbCreateDbInstance          bool
	_timestreaminfluxdbCreateDbParameterGroup    bool
	_timestreaminfluxdbDeleteDbCluster           bool
	_timestreaminfluxdbDeleteDbInstance          bool
	_timestreaminfluxdbGetDbCluster              bool
	_timestreaminfluxdbGetDbInstance             bool
	_timestreaminfluxdbGetDbParameterGroup       bool
	_timestreaminfluxdbListDbClusters            bool
	_timestreaminfluxdbListDbInstances           bool
	_timestreaminfluxdbListDbInstancesForCluster bool
	_timestreaminfluxdbListDbParameterGroups     bool
	_timestreaminfluxdbListTagsForResource       bool
	_timestreaminfluxdbRebootDbCluster           bool
	_timestreaminfluxdbRebootDbInstance          bool
	_timestreaminfluxdbTagResource               bool
	_timestreaminfluxdbUntagResource             bool
	_timestreaminfluxdbUpdateDbCluster           bool
	_timestreaminfluxdbUpdateDbInstance          bool

	_timestreaminfluxdbAllocatedStorage           string
	_timestreaminfluxdbBucket                     string
	_timestreaminfluxdbDbClusterId                string
	_timestreaminfluxdbDbInstanceType             string
	_timestreaminfluxdbDbParameterGroupIdentifier string
	_timestreaminfluxdbDbStorageType              string
	_timestreaminfluxdbDeploymentType             string
	_timestreaminfluxdbDescription                string
	_timestreaminfluxdbFailoverMode               string
	_timestreaminfluxdbIdentifier                 string
	_timestreaminfluxdbInstanceIds                []string
	_timestreaminfluxdbLogDeliveryConfiguration   string
	_timestreaminfluxdbMaxResults                 string
	_timestreaminfluxdbName                       string
	_timestreaminfluxdbNetworkType                string
	_timestreaminfluxdbNextToken                  string
	_timestreaminfluxdbOrganization               string
	_timestreaminfluxdbParameters                 string
	_timestreaminfluxdbPassword                   string
	_timestreaminfluxdbPort                       string
	_timestreaminfluxdbPubliclyAccessible         string
	_timestreaminfluxdbResourceArn                string
	_timestreaminfluxdbTagKeys                    []string
	_timestreaminfluxdbTags                       string
	_timestreaminfluxdbUsername                   string
	_timestreaminfluxdbVpcSecurityGroupIds        []string
	_timestreaminfluxdbVpcSubnetIds               []string
)

// Creates a new Timestream for InfluxDB cluster.
func timestreaminfluxdb_CreateDbCluster(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.CreateDbClusterInput{
		// DbInstanceType: types.DbInstanceType, // Required
		// Name: *string, // Required
		// VpcSecurityGroupIds: []string, // Required
		// VpcSubnetIds: []string, // Required
	}

	if len(_timestreaminfluxdbDbInstanceType) > 0 {
		if err := assignInputField(input, "DbInstanceType", _timestreaminfluxdbDbInstanceType); err != nil {
			log.Errorf("invalid --db-instance-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbName) > 0 {
		input.Name = aws.String(_timestreaminfluxdbName)
	}
	if len(_timestreaminfluxdbVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _timestreaminfluxdbVpcSecurityGroupIds...)
	}
	if len(_timestreaminfluxdbVpcSubnetIds) > 0 {
		input.VpcSubnetIds = append([]string(nil), _timestreaminfluxdbVpcSubnetIds...)
	}
	if len(_timestreaminfluxdbAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _timestreaminfluxdbAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbBucket) > 0 {
		input.Bucket = aws.String(_timestreaminfluxdbBucket)
	}
	if len(_timestreaminfluxdbDbParameterGroupIdentifier) > 0 {
		input.DbParameterGroupIdentifier = aws.String(_timestreaminfluxdbDbParameterGroupIdentifier)
	}
	if len(_timestreaminfluxdbDbStorageType) > 0 {
		if err := assignInputField(input, "DbStorageType", _timestreaminfluxdbDbStorageType); err != nil {
			log.Errorf("invalid --db-storage-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _timestreaminfluxdbDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbFailoverMode) > 0 {
		if err := assignInputField(input, "FailoverMode", _timestreaminfluxdbFailoverMode); err != nil {
			log.Errorf("invalid --failover-mode: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbLogDeliveryConfiguration) > 0 {
		if err := assignInputField(input, "LogDeliveryConfiguration", _timestreaminfluxdbLogDeliveryConfiguration); err != nil {
			log.Errorf("invalid --log-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _timestreaminfluxdbNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbOrganization) > 0 {
		input.Organization = aws.String(_timestreaminfluxdbOrganization)
	}
	if len(_timestreaminfluxdbPassword) > 0 {
		input.Password = aws.String(_timestreaminfluxdbPassword)
	}
	if len(_timestreaminfluxdbPort) > 0 {
		if err := assignInputField(input, "Port", _timestreaminfluxdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _timestreaminfluxdbPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreaminfluxdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbUsername) > 0 {
		input.Username = aws.String(_timestreaminfluxdbUsername)
	}

	if resp, err := client.CreateDbCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Timestream for InfluxDB DB instance.
func timestreaminfluxdb_CreateDbInstance(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.CreateDbInstanceInput{
		// AllocatedStorage: *int32, // Required
		// DbInstanceType: types.DbInstanceType, // Required
		// Name: *string, // Required
		// Password: *string, // Required
		// VpcSecurityGroupIds: []string, // Required
		// VpcSubnetIds: []string, // Required
	}

	if len(_timestreaminfluxdbAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _timestreaminfluxdbAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDbInstanceType) > 0 {
		if err := assignInputField(input, "DbInstanceType", _timestreaminfluxdbDbInstanceType); err != nil {
			log.Errorf("invalid --db-instance-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbName) > 0 {
		input.Name = aws.String(_timestreaminfluxdbName)
	}
	if len(_timestreaminfluxdbPassword) > 0 {
		input.Password = aws.String(_timestreaminfluxdbPassword)
	}
	if len(_timestreaminfluxdbVpcSecurityGroupIds) > 0 {
		input.VpcSecurityGroupIds = append([]string(nil), _timestreaminfluxdbVpcSecurityGroupIds...)
	}
	if len(_timestreaminfluxdbVpcSubnetIds) > 0 {
		input.VpcSubnetIds = append([]string(nil), _timestreaminfluxdbVpcSubnetIds...)
	}
	if len(_timestreaminfluxdbBucket) > 0 {
		input.Bucket = aws.String(_timestreaminfluxdbBucket)
	}
	if len(_timestreaminfluxdbDbParameterGroupIdentifier) > 0 {
		input.DbParameterGroupIdentifier = aws.String(_timestreaminfluxdbDbParameterGroupIdentifier)
	}
	if len(_timestreaminfluxdbDbStorageType) > 0 {
		if err := assignInputField(input, "DbStorageType", _timestreaminfluxdbDbStorageType); err != nil {
			log.Errorf("invalid --db-storage-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _timestreaminfluxdbDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbLogDeliveryConfiguration) > 0 {
		if err := assignInputField(input, "LogDeliveryConfiguration", _timestreaminfluxdbLogDeliveryConfiguration); err != nil {
			log.Errorf("invalid --log-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _timestreaminfluxdbNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbOrganization) > 0 {
		input.Organization = aws.String(_timestreaminfluxdbOrganization)
	}
	if len(_timestreaminfluxdbPort) > 0 {
		if err := assignInputField(input, "Port", _timestreaminfluxdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbPubliclyAccessible) > 0 {
		if err := assignInputField(input, "PubliclyAccessible", _timestreaminfluxdbPubliclyAccessible); err != nil {
			log.Errorf("invalid --publicly-accessible: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreaminfluxdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbUsername) > 0 {
		input.Username = aws.String(_timestreaminfluxdbUsername)
	}

	if resp, err := client.CreateDbInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Timestream for InfluxDB DB parameter group to associate with DB
// instances.
func timestreaminfluxdb_CreateDbParameterGroup(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.CreateDbParameterGroupInput{
		// Name: *string, // Required
	}

	if len(_timestreaminfluxdbName) > 0 {
		input.Name = aws.String(_timestreaminfluxdbName)
	}
	if len(_timestreaminfluxdbDescription) > 0 {
		input.Description = aws.String(_timestreaminfluxdbDescription)
	}
	if len(_timestreaminfluxdbParameters) > 0 {
		if err := assignInputField(input, "Parameters", _timestreaminfluxdbParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreaminfluxdbTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDbParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Timestream for InfluxDB cluster.
func timestreaminfluxdb_DeleteDbCluster(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.DeleteDbClusterInput{
		// DbClusterId: *string, // Required
	}

	if len(_timestreaminfluxdbDbClusterId) > 0 {
		input.DbClusterId = aws.String(_timestreaminfluxdbDbClusterId)
	}

	if resp, err := client.DeleteDbCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a Timestream for InfluxDB DB instance.
func timestreaminfluxdb_DeleteDbInstance(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.DeleteDbInstanceInput{
		// Identifier: *string, // Required
	}

	if len(_timestreaminfluxdbIdentifier) > 0 {
		input.Identifier = aws.String(_timestreaminfluxdbIdentifier)
	}

	if resp, err := client.DeleteDbInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a Timestream for InfluxDB cluster.
func timestreaminfluxdb_GetDbCluster(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.GetDbClusterInput{
		// DbClusterId: *string, // Required
	}

	if len(_timestreaminfluxdbDbClusterId) > 0 {
		input.DbClusterId = aws.String(_timestreaminfluxdbDbClusterId)
	}

	if resp, err := client.GetDbCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a Timestream for InfluxDB DB instance.
func timestreaminfluxdb_GetDbInstance(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.GetDbInstanceInput{
		// Identifier: *string, // Required
	}

	if len(_timestreaminfluxdbIdentifier) > 0 {
		input.Identifier = aws.String(_timestreaminfluxdbIdentifier)
	}

	if resp, err := client.GetDbInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a Timestream for InfluxDB DB parameter group.
func timestreaminfluxdb_GetDbParameterGroup(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.GetDbParameterGroupInput{
		// Identifier: *string, // Required
	}

	if len(_timestreaminfluxdbIdentifier) > 0 {
		input.Identifier = aws.String(_timestreaminfluxdbIdentifier)
	}

	if resp, err := client.GetDbParameterGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Timestream for InfluxDB DB clusters.
func timestreaminfluxdb_ListDbClusters(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.ListDbClustersInput{}

	if len(_timestreaminfluxdbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreaminfluxdbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbNextToken) > 0 {
		input.NextToken = aws.String(_timestreaminfluxdbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*timestreaminfluxdb.ListDbClustersOutput
	p := timestreaminfluxdb.NewListDbClustersPaginator(client, input)
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

// Returns a list of Timestream for InfluxDB DB instances.
func timestreaminfluxdb_ListDbInstances(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.ListDbInstancesInput{}

	if len(_timestreaminfluxdbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreaminfluxdbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbNextToken) > 0 {
		input.NextToken = aws.String(_timestreaminfluxdbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*timestreaminfluxdb.ListDbInstancesOutput
	p := timestreaminfluxdb.NewListDbInstancesPaginator(client, input)
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

// Returns a list of Timestream for InfluxDB clusters.
func timestreaminfluxdb_ListDbInstancesForCluster(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.ListDbInstancesForClusterInput{
		// DbClusterId: *string, // Required
	}

	if len(_timestreaminfluxdbDbClusterId) > 0 {
		input.DbClusterId = aws.String(_timestreaminfluxdbDbClusterId)
	}
	if len(_timestreaminfluxdbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreaminfluxdbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbNextToken) > 0 {
		input.NextToken = aws.String(_timestreaminfluxdbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbInstancesForCluster(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*timestreaminfluxdb.ListDbInstancesForClusterOutput
	p := timestreaminfluxdb.NewListDbInstancesForClusterPaginator(client, input)
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

// Returns a list of Timestream for InfluxDB DB parameter groups.
func timestreaminfluxdb_ListDbParameterGroups(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.ListDbParameterGroupsInput{}

	if len(_timestreaminfluxdbMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _timestreaminfluxdbMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbNextToken) > 0 {
		input.NextToken = aws.String(_timestreaminfluxdbNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDbParameterGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*timestreaminfluxdb.ListDbParameterGroupsOutput
	p := timestreaminfluxdb.NewListDbParameterGroupsPaginator(client, input)
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

// A list of tags applied to the resource.
func timestreaminfluxdb_ListTagsForResource(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_timestreaminfluxdbResourceArn) > 0 {
		input.ResourceArn = aws.String(_timestreaminfluxdbResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots a Timestream for InfluxDB cluster.
func timestreaminfluxdb_RebootDbCluster(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.RebootDbClusterInput{
		// DbClusterId: *string, // Required
	}

	if len(_timestreaminfluxdbDbClusterId) > 0 {
		input.DbClusterId = aws.String(_timestreaminfluxdbDbClusterId)
	}
	if len(_timestreaminfluxdbInstanceIds) > 0 {
		input.InstanceIds = append([]string(nil), _timestreaminfluxdbInstanceIds...)
	}

	if resp, err := client.RebootDbCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reboots a Timestream for InfluxDB instance.
func timestreaminfluxdb_RebootDbInstance(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.RebootDbInstanceInput{
		// Identifier: *string, // Required
	}

	if len(_timestreaminfluxdbIdentifier) > 0 {
		input.Identifier = aws.String(_timestreaminfluxdbIdentifier)
	}

	if resp, err := client.RebootDbInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags are composed of a Key/Value pairs. You can use tags to categorize and
// track your Timestream for InfluxDB resources.
func timestreaminfluxdb_TagResource(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_timestreaminfluxdbResourceArn) > 0 {
		input.ResourceArn = aws.String(_timestreaminfluxdbResourceArn)
	}
	if len(_timestreaminfluxdbTags) > 0 {
		if err := assignInputField(input, "Tags", _timestreaminfluxdbTags); err != nil {
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

// Removes the tag from the specified resource.
func timestreaminfluxdb_UntagResource(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_timestreaminfluxdbResourceArn) > 0 {
		input.ResourceArn = aws.String(_timestreaminfluxdbResourceArn)
	}
	if len(_timestreaminfluxdbTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _timestreaminfluxdbTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Timestream for InfluxDB cluster.
func timestreaminfluxdb_UpdateDbCluster(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.UpdateDbClusterInput{
		// DbClusterId: *string, // Required
	}

	if len(_timestreaminfluxdbDbClusterId) > 0 {
		input.DbClusterId = aws.String(_timestreaminfluxdbDbClusterId)
	}
	if len(_timestreaminfluxdbDbInstanceType) > 0 {
		if err := assignInputField(input, "DbInstanceType", _timestreaminfluxdbDbInstanceType); err != nil {
			log.Errorf("invalid --db-instance-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDbParameterGroupIdentifier) > 0 {
		input.DbParameterGroupIdentifier = aws.String(_timestreaminfluxdbDbParameterGroupIdentifier)
	}
	if len(_timestreaminfluxdbFailoverMode) > 0 {
		if err := assignInputField(input, "FailoverMode", _timestreaminfluxdbFailoverMode); err != nil {
			log.Errorf("invalid --failover-mode: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbLogDeliveryConfiguration) > 0 {
		if err := assignInputField(input, "LogDeliveryConfiguration", _timestreaminfluxdbLogDeliveryConfiguration); err != nil {
			log.Errorf("invalid --log-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbPort) > 0 {
		if err := assignInputField(input, "Port", _timestreaminfluxdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDbCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a Timestream for InfluxDB DB instance.
func timestreaminfluxdb_UpdateDbInstance(cfg aws.Config, client *timestreaminfluxdb.Client) {
	input := &timestreaminfluxdb.UpdateDbInstanceInput{
		// Identifier: *string, // Required
	}

	if len(_timestreaminfluxdbIdentifier) > 0 {
		input.Identifier = aws.String(_timestreaminfluxdbIdentifier)
	}
	if len(_timestreaminfluxdbAllocatedStorage) > 0 {
		if err := assignInputField(input, "AllocatedStorage", _timestreaminfluxdbAllocatedStorage); err != nil {
			log.Errorf("invalid --allocated-storage: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDbInstanceType) > 0 {
		if err := assignInputField(input, "DbInstanceType", _timestreaminfluxdbDbInstanceType); err != nil {
			log.Errorf("invalid --db-instance-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDbParameterGroupIdentifier) > 0 {
		input.DbParameterGroupIdentifier = aws.String(_timestreaminfluxdbDbParameterGroupIdentifier)
	}
	if len(_timestreaminfluxdbDbStorageType) > 0 {
		if err := assignInputField(input, "DbStorageType", _timestreaminfluxdbDbStorageType); err != nil {
			log.Errorf("invalid --db-storage-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbDeploymentType) > 0 {
		if err := assignInputField(input, "DeploymentType", _timestreaminfluxdbDeploymentType); err != nil {
			log.Errorf("invalid --deployment-type: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbLogDeliveryConfiguration) > 0 {
		if err := assignInputField(input, "LogDeliveryConfiguration", _timestreaminfluxdbLogDeliveryConfiguration); err != nil {
			log.Errorf("invalid --log-delivery-configuration: %s", err.Error())
			return
		}
	}
	if len(_timestreaminfluxdbPort) > 0 {
		if err := assignInputField(input, "Port", _timestreaminfluxdbPort); err != nil {
			log.Errorf("invalid --port: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDbInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_timestreaminfluxdbCmd)
	_timestreaminfluxdbCmd.Flags().SortFlags = false

	_timestreaminfluxdbCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_timestreaminfluxdbCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbAllocatedStorage, "allocated-storage", "", "", "Allocated Storage")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbBucket, "bucket", "", "", "Bucket")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbDbClusterId, "db-cluster-id", "", "", "DB Cluster ID")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbDbInstanceType, "db-instance-type", "", "", "DB Instance Type")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbDbParameterGroupIdentifier, "db-parameter-group-identifier", "", "", "DB Parameter Group Identifier")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbDbStorageType, "db-storage-type", "", "", "DB Storage Type")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbDeploymentType, "deployment-type", "", "", "Deployment Type")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbDescription, "description", "", "", "Description")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbFailoverMode, "failover-mode", "", "", "Failover Mode")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbIdentifier, "identifier", "", "", "Identifier")
	_timestreaminfluxdbCmd.Flags().StringSliceVarP(&_timestreaminfluxdbInstanceIds, "instance-ids", "", nil, "Instance Ids")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbLogDeliveryConfiguration, "log-delivery-configuration", "", "", "Log Delivery Configuration")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbMaxResults, "max-results", "", "", "Max Results")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbName, "name", "", "", "Name")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbNetworkType, "network-type", "", "", "Network Type")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbNextToken, "next-token", "", "", "Next Token")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbOrganization, "organization", "", "", "Organization")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbParameters, "parameters", "", "", "Parameters")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbPassword, "password", "", "", "Password")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbPort, "port", "", "", "Port")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbPubliclyAccessible, "publicly-accessible", "", "", "Publicly Accessible")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbResourceArn, "resource-arn", "", "", "Resource ARN")
	_timestreaminfluxdbCmd.Flags().StringSliceVarP(&_timestreaminfluxdbTagKeys, "tag-keys", "", nil, "Tag Keys")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbTags, "tags", "", "", "Tags")
	_timestreaminfluxdbCmd.Flags().StringVarP(&_timestreaminfluxdbUsername, "username", "", "", "Username")
	_timestreaminfluxdbCmd.Flags().StringSliceVarP(&_timestreaminfluxdbVpcSecurityGroupIds, "vpc-security-group-ids", "", nil, "VPC Security Group Ids")
	_timestreaminfluxdbCmd.Flags().StringSliceVarP(&_timestreaminfluxdbVpcSubnetIds, "vpc-subnet-ids", "", nil, "VPC Subnet Ids")

	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbCreateDbCluster, "create-db-cluster", "", false, "Create DB Cluster")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbCreateDbInstance, "create-db-instance", "", false, "Create DB Instance")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbCreateDbParameterGroup, "create-db-parameter-group", "", false, "Create DB Parameter Group")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbDeleteDbCluster, "delete-db-cluster", "", false, "Delete DB Cluster")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbDeleteDbInstance, "delete-db-instance", "", false, "Delete DB Instance")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbGetDbCluster, "get-db-cluster", "", false, "Get DB Cluster")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbGetDbInstance, "get-db-instance", "", false, "Get DB Instance")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbGetDbParameterGroup, "get-db-parameter-group", "", false, "Get DB Parameter Group")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbListDbClusters, "list-db-clusters", "", false, "List DB Clusters")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbListDbInstances, "list-db-instances", "", false, "List DB Instances")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbListDbInstancesForCluster, "list-db-instances-for-cluster", "", false, "List DB Instances For Cluster")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbListDbParameterGroups, "list-db-parameter-groups", "", false, "List DB Parameter Groups")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbRebootDbCluster, "reboot-db-cluster", "", false, "Reboot DB Cluster")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbRebootDbInstance, "reboot-db-instance", "", false, "Reboot DB Instance")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbTagResource, "tag-resource", "", false, "Tag Resource")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbUntagResource, "untag-resource", "", false, "Untag Resource")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbUpdateDbCluster, "update-db-cluster", "", false, "Update DB Cluster")
	_timestreaminfluxdbCmd.Flags().BoolVarP(&_timestreaminfluxdbUpdateDbInstance, "update-db-instance", "", false, "Update DB Instance")

}
