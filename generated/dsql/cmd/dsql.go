package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dsql"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// dsqlCmd represents the dsql command
var _dsqlCmd = &cobra.Command{
	Use:   "dsql",
	Short: "AWS dsql CLI",
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
		client := dsql.NewFromConfig(cfg)
		if _dsqlCreateCluster {
			dsql_CreateCluster(cfg, client)
			return
		}
		if _dsqlDeleteCluster {
			dsql_DeleteCluster(cfg, client)
			return
		}
		if _dsqlDeleteClusterPolicy {
			dsql_DeleteClusterPolicy(cfg, client)
			return
		}
		if _dsqlGetCluster {
			dsql_GetCluster(cfg, client)
			return
		}
		if _dsqlGetClusterPolicy {
			dsql_GetClusterPolicy(cfg, client)
			return
		}
		if _dsqlGetVpcEndpointServiceName {
			dsql_GetVpcEndpointServiceName(cfg, client)
			return
		}
		if _dsqlListClusters {
			dsql_ListClusters(cfg, client)
			return
		}
		if _dsqlListTagsForResource {
			dsql_ListTagsForResource(cfg, client)
			return
		}
		if _dsqlPutClusterPolicy {
			dsql_PutClusterPolicy(cfg, client)
			return
		}
		if _dsqlTagResource {
			dsql_TagResource(cfg, client)
			return
		}
		if _dsqlUntagResource {
			dsql_UntagResource(cfg, client)
			return
		}
		if _dsqlUpdateCluster {
			dsql_UpdateCluster(cfg, client)
			return
		}

	},
}

var (
	_dsqlCreateCluster             bool
	_dsqlDeleteCluster             bool
	_dsqlDeleteClusterPolicy       bool
	_dsqlGetCluster                bool
	_dsqlGetClusterPolicy          bool
	_dsqlGetVpcEndpointServiceName bool
	_dsqlListClusters              bool
	_dsqlListTagsForResource       bool
	_dsqlPutClusterPolicy          bool
	_dsqlTagResource               bool
	_dsqlUntagResource             bool
	_dsqlUpdateCluster             bool

	_dsqlBypassPolicyLockoutSafetyCheck string
	_dsqlClientToken                    string
	_dsqlDeletionProtectionEnabled      string
	_dsqlExpectedPolicyVersion          string
	_dsqlIdentifier                     string
	_dsqlKmsEncryptionKey               string
	_dsqlMaxResults                     string
	_dsqlMultiRegionProperties          string
	_dsqlNextToken                      string
	_dsqlPolicy                         string
	_dsqlResourceArn                    string
	_dsqlTagKeys                        []string
	_dsqlTags                           string
)

// The CreateCluster API allows you to create both single-Region clusters and
// multi-Region clusters. With the addition of the multiRegionProperties parameter,
// you can create a cluster with witness Region support and establish peer
// relationships with clusters in other Regions during creation.
//
// Creating multi-Region clusters requires additional IAM permissions beyond those
// needed for single-Region clusters, as detailed in the Required permissions
// section below.
//
// # Required permissions
//
// dsql:CreateCluster Required to create a cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// dsql:TagResource Permission to add tags to a resource.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// dsql:PutMultiRegionProperties Permission to configure multi-Region properties
// for a cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// dsql:AddPeerCluster When specifying multiRegionProperties.clusters , permission
// to add peer clusters.
//
// Resources:
//
// - Local cluster: arn:aws:dsql:region:account-id:cluster/*
//
// - Each peer cluster: exact ARN of each specified peer cluster
//
// dsql:PutWitnessRegion When specifying multiRegionProperties.witnessRegion ,
// permission to set a witness Region. This permission is checked both in the
// cluster Region and in the witness Region.
//
// Resources: arn:aws:dsql:region:account-id:cluster/*
//
// Condition Keys: dsql:WitnessRegion (matching the specified witness region)
//
// - The witness Region specified in multiRegionProperties.witnessRegion cannot
// be the same as the cluster's Region.
func dsql_CreateCluster(cfg aws.Config, client *dsql.Client) {
	input := &dsql.CreateClusterInput{}

	if len(_dsqlBypassPolicyLockoutSafetyCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutSafetyCheck", _dsqlBypassPolicyLockoutSafetyCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-safety-check: %s", err.Error())
			return
		}
	}
	if len(_dsqlClientToken) > 0 {
		input.ClientToken = aws.String(_dsqlClientToken)
	}
	if len(_dsqlDeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _dsqlDeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_dsqlKmsEncryptionKey) > 0 {
		input.KmsEncryptionKey = aws.String(_dsqlKmsEncryptionKey)
	}
	if len(_dsqlMultiRegionProperties) > 0 {
		if err := assignInputField(input, "MultiRegionProperties", _dsqlMultiRegionProperties); err != nil {
			log.Errorf("invalid --multi-region-properties: %s", err.Error())
			return
		}
	}
	if len(_dsqlPolicy) > 0 {
		input.Policy = aws.String(_dsqlPolicy)
	}
	if len(_dsqlTags) > 0 {
		if err := assignInputField(input, "Tags", _dsqlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
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

// Deletes a cluster in Amazon Aurora DSQL.
func dsql_DeleteCluster(cfg aws.Config, client *dsql.Client) {
	input := &dsql.DeleteClusterInput{
		// Identifier: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}
	if len(_dsqlClientToken) > 0 {
		input.ClientToken = aws.String(_dsqlClientToken)
	}

	if resp, err := client.DeleteCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource-based policy attached to a cluster. This removes all
// access permissions defined by the policy, reverting to default access controls.
func dsql_DeleteClusterPolicy(cfg aws.Config, client *dsql.Client) {
	input := &dsql.DeleteClusterPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}
	if len(_dsqlClientToken) > 0 {
		input.ClientToken = aws.String(_dsqlClientToken)
	}
	if len(_dsqlExpectedPolicyVersion) > 0 {
		input.ExpectedPolicyVersion = aws.String(_dsqlExpectedPolicyVersion)
	}

	if resp, err := client.DeleteClusterPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a cluster.
func dsql_GetCluster(cfg aws.Config, client *dsql.Client) {
	input := &dsql.GetClusterInput{
		// Identifier: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}

	if resp, err := client.GetCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource-based policy document attached to a cluster. This policy
// defines the access permissions and conditions for the cluster.
func dsql_GetClusterPolicy(cfg aws.Config, client *dsql.Client) {
	input := &dsql.GetClusterPolicyInput{
		// Identifier: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}

	if resp, err := client.GetClusterPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the VPC endpoint service name.
func dsql_GetVpcEndpointServiceName(cfg aws.Config, client *dsql.Client) {
	input := &dsql.GetVpcEndpointServiceNameInput{
		// Identifier: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}

	if resp, err := client.GetVpcEndpointServiceName(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a list of clusters.
func dsql_ListClusters(cfg aws.Config, client *dsql.Client) {
	input := &dsql.ListClustersInput{}

	if len(_dsqlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _dsqlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_dsqlNextToken) > 0 {
		input.NextToken = aws.String(_dsqlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListClusters(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*dsql.ListClustersOutput
	p := dsql.NewListClustersPaginator(client, input)
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

// Lists all of the tags for a resource.
func dsql_ListTagsForResource(cfg aws.Config, client *dsql.Client) {
	input := &dsql.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_dsqlResourceArn) > 0 {
		input.ResourceArn = aws.String(_dsqlResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches a resource-based policy to a cluster. This policy defines access
// permissions and conditions for the cluster, allowing you to control which
// principals can perform actions on the cluster.
func dsql_PutClusterPolicy(cfg aws.Config, client *dsql.Client) {
	input := &dsql.PutClusterPolicyInput{
		// Identifier: *string, // Required
		// Policy: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}
	if len(_dsqlPolicy) > 0 {
		input.Policy = aws.String(_dsqlPolicy)
	}
	if len(_dsqlBypassPolicyLockoutSafetyCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutSafetyCheck", _dsqlBypassPolicyLockoutSafetyCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-safety-check: %s", err.Error())
			return
		}
	}
	if len(_dsqlClientToken) > 0 {
		input.ClientToken = aws.String(_dsqlClientToken)
	}
	if len(_dsqlExpectedPolicyVersion) > 0 {
		input.ExpectedPolicyVersion = aws.String(_dsqlExpectedPolicyVersion)
	}

	if resp, err := client.PutClusterPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Tags a resource with a map of key and value pairs.
func dsql_TagResource(cfg aws.Config, client *dsql.Client) {
	input := &dsql.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_dsqlResourceArn) > 0 {
		input.ResourceArn = aws.String(_dsqlResourceArn)
	}
	if len(_dsqlTags) > 0 {
		if err := assignInputField(input, "Tags", _dsqlTags); err != nil {
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

// Removes a tag from a resource.
func dsql_UntagResource(cfg aws.Config, client *dsql.Client) {
	input := &dsql.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_dsqlResourceArn) > 0 {
		input.ResourceArn = aws.String(_dsqlResourceArn)
	}
	if len(_dsqlTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _dsqlTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The UpdateCluster API allows you to modify both single-Region and multi-Region
// cluster configurations. With the multiRegionProperties parameter, you can add or
// modify witness Region support and manage peer relationships with clusters in
// other Regions.
//
// Note that updating multi-Region clusters requires additional IAM permissions
// beyond those needed for standard cluster updates, as detailed in the Permissions
// section.
//
// # Required permissions
//
// dsql:UpdateCluster Permission to update a DSQL cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// dsql:PutMultiRegionProperties Permission to configure multi-Region properties
// for a cluster.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// dsql:GetCluster Permission to retrieve cluster information.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// dsql:AddPeerCluster Permission to add peer clusters.
//
// Resources:
//
// - Local cluster: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// - Each peer cluster: exact ARN of each specified peer cluster
//
// dsql:RemovePeerCluster Permission to remove peer clusters. The
// dsql:RemovePeerCluster permission uses a wildcard ARN pattern to simplify
// permission management during updates.
//
// Resources: arn:aws:dsql:*:account-id:cluster/*
//
// dsql:PutWitnessRegion Permission to set a witness Region.
//
// Resources: arn:aws:dsql:region:account-id:cluster/cluster-id
//
// Condition Keys: dsql:WitnessRegion (matching the specified witness Region)
//
// This permission is checked both in the cluster Region and in the witness
// Region.
//
// - The witness region specified in multiRegionProperties.witnessRegion cannot
// be the same as the cluster's Region.
//
// - When updating clusters with peer relationships, permissions are checked for
// both adding and removing peers.
//
// - The dsql:RemovePeerCluster permission uses a wildcard ARN pattern to
// simplify permission management during updates.
func dsql_UpdateCluster(cfg aws.Config, client *dsql.Client) {
	input := &dsql.UpdateClusterInput{
		// Identifier: *string, // Required
	}

	if len(_dsqlIdentifier) > 0 {
		input.Identifier = aws.String(_dsqlIdentifier)
	}
	if len(_dsqlClientToken) > 0 {
		input.ClientToken = aws.String(_dsqlClientToken)
	}
	if len(_dsqlDeletionProtectionEnabled) > 0 {
		if err := assignInputField(input, "DeletionProtectionEnabled", _dsqlDeletionProtectionEnabled); err != nil {
			log.Errorf("invalid --deletion-protection-enabled: %s", err.Error())
			return
		}
	}
	if len(_dsqlKmsEncryptionKey) > 0 {
		input.KmsEncryptionKey = aws.String(_dsqlKmsEncryptionKey)
	}
	if len(_dsqlMultiRegionProperties) > 0 {
		if err := assignInputField(input, "MultiRegionProperties", _dsqlMultiRegionProperties); err != nil {
			log.Errorf("invalid --multi-region-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCluster(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_dsqlCmd)
	_dsqlCmd.Flags().SortFlags = false

	_dsqlCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_dsqlCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_dsqlCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_dsqlCmd.Flags().StringVarP(&_dsqlBypassPolicyLockoutSafetyCheck, "bypass-policy-lockout-safety-check", "", "", "Bypass Policy Lockout Safety Check")
	_dsqlCmd.Flags().StringVarP(&_dsqlClientToken, "client-token", "", "", "Client Token")
	_dsqlCmd.Flags().StringVarP(&_dsqlDeletionProtectionEnabled, "deletion-protection-enabled", "", "", "Deletion Protection Enabled")
	_dsqlCmd.Flags().StringVarP(&_dsqlExpectedPolicyVersion, "expected-policy-version", "", "", "Expected Policy Version")
	_dsqlCmd.Flags().StringVarP(&_dsqlIdentifier, "identifier", "", "", "Identifier")
	_dsqlCmd.Flags().StringVarP(&_dsqlKmsEncryptionKey, "kms-encryption-key", "", "", "KMS Encryption Key")
	_dsqlCmd.Flags().StringVarP(&_dsqlMaxResults, "max-results", "", "", "Max Results")
	_dsqlCmd.Flags().StringVarP(&_dsqlMultiRegionProperties, "multi-region-properties", "", "", "Multi Region Properties")
	_dsqlCmd.Flags().StringVarP(&_dsqlNextToken, "next-token", "", "", "Next Token")
	_dsqlCmd.Flags().StringVarP(&_dsqlPolicy, "policy", "", "", "Policy")
	_dsqlCmd.Flags().StringVarP(&_dsqlResourceArn, "resource-arn", "", "", "Resource ARN")
	_dsqlCmd.Flags().StringSliceVarP(&_dsqlTagKeys, "tag-keys", "", nil, "Tag Keys")
	_dsqlCmd.Flags().StringVarP(&_dsqlTags, "tags", "", "", "Tags")

	_dsqlCmd.Flags().BoolVarP(&_dsqlCreateCluster, "create-cluster", "", false, "Create Cluster")
	_dsqlCmd.Flags().BoolVarP(&_dsqlDeleteCluster, "delete-cluster", "", false, "Delete Cluster")
	_dsqlCmd.Flags().BoolVarP(&_dsqlDeleteClusterPolicy, "delete-cluster-policy", "", false, "Delete Cluster Policy")
	_dsqlCmd.Flags().BoolVarP(&_dsqlGetCluster, "get-cluster", "", false, "Get Cluster")
	_dsqlCmd.Flags().BoolVarP(&_dsqlGetClusterPolicy, "get-cluster-policy", "", false, "Get Cluster Policy")
	_dsqlCmd.Flags().BoolVarP(&_dsqlGetVpcEndpointServiceName, "get-vpc-endpoint-service-name", "", false, "Get VPC Endpoint Service Name")
	_dsqlCmd.Flags().BoolVarP(&_dsqlListClusters, "list-clusters", "", false, "List Clusters")
	_dsqlCmd.Flags().BoolVarP(&_dsqlListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_dsqlCmd.Flags().BoolVarP(&_dsqlPutClusterPolicy, "put-cluster-policy", "", false, "Put Cluster Policy")
	_dsqlCmd.Flags().BoolVarP(&_dsqlTagResource, "tag-resource", "", false, "Tag Resource")
	_dsqlCmd.Flags().BoolVarP(&_dsqlUntagResource, "untag-resource", "", false, "Untag Resource")
	_dsqlCmd.Flags().BoolVarP(&_dsqlUpdateCluster, "update-cluster", "", false, "Update Cluster")

}
