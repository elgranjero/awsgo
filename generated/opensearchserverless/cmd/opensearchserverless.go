package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearchserverless"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// opensearchserverlessCmd represents the opensearchserverless command
var _opensearchserverlessCmd = &cobra.Command{
	Use:   "opensearchserverless",
	Short: "AWS opensearchserverless CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := opensearchserverless.NewFromConfig(cfg)
		if _opensearchserverlessBatchGetCollection {
			opensearchserverless_BatchGetCollection(cfg, client)
			return
		}
		if _opensearchserverlessBatchGetCollectionGroup {
			opensearchserverless_BatchGetCollectionGroup(cfg, client)
			return
		}
		if _opensearchserverlessBatchGetEffectiveLifecyclePolicy {
			opensearchserverless_BatchGetEffectiveLifecyclePolicy(cfg, client)
			return
		}
		if _opensearchserverlessBatchGetLifecyclePolicy {
			opensearchserverless_BatchGetLifecyclePolicy(cfg, client)
			return
		}
		if _opensearchserverlessBatchGetVpcEndpoint {
			opensearchserverless_BatchGetVpcEndpoint(cfg, client)
			return
		}
		if _opensearchserverlessCreateAccessPolicy {
			opensearchserverless_CreateAccessPolicy(cfg, client)
			return
		}
		if _opensearchserverlessCreateCollection {
			opensearchserverless_CreateCollection(cfg, client)
			return
		}
		if _opensearchserverlessCreateCollectionGroup {
			opensearchserverless_CreateCollectionGroup(cfg, client)
			return
		}
		if _opensearchserverlessCreateIndex {
			opensearchserverless_CreateIndex(cfg, client)
			return
		}
		if _opensearchserverlessCreateLifecyclePolicy {
			opensearchserverless_CreateLifecyclePolicy(cfg, client)
			return
		}
		if _opensearchserverlessCreateSecurityConfig {
			opensearchserverless_CreateSecurityConfig(cfg, client)
			return
		}
		if _opensearchserverlessCreateSecurityPolicy {
			opensearchserverless_CreateSecurityPolicy(cfg, client)
			return
		}
		if _opensearchserverlessCreateVpcEndpoint {
			opensearchserverless_CreateVpcEndpoint(cfg, client)
			return
		}
		if _opensearchserverlessDeleteAccessPolicy {
			opensearchserverless_DeleteAccessPolicy(cfg, client)
			return
		}
		if _opensearchserverlessDeleteCollection {
			opensearchserverless_DeleteCollection(cfg, client)
			return
		}
		if _opensearchserverlessDeleteCollectionGroup {
			opensearchserverless_DeleteCollectionGroup(cfg, client)
			return
		}
		if _opensearchserverlessDeleteIndex {
			opensearchserverless_DeleteIndex(cfg, client)
			return
		}
		if _opensearchserverlessDeleteLifecyclePolicy {
			opensearchserverless_DeleteLifecyclePolicy(cfg, client)
			return
		}
		if _opensearchserverlessDeleteSecurityConfig {
			opensearchserverless_DeleteSecurityConfig(cfg, client)
			return
		}
		if _opensearchserverlessDeleteSecurityPolicy {
			opensearchserverless_DeleteSecurityPolicy(cfg, client)
			return
		}
		if _opensearchserverlessDeleteVpcEndpoint {
			opensearchserverless_DeleteVpcEndpoint(cfg, client)
			return
		}
		if _opensearchserverlessGetAccessPolicy {
			opensearchserverless_GetAccessPolicy(cfg, client)
			return
		}
		if _opensearchserverlessGetAccountSettings {
			opensearchserverless_GetAccountSettings(cfg, client)
			return
		}
		if _opensearchserverlessGetIndex {
			opensearchserverless_GetIndex(cfg, client)
			return
		}
		if _opensearchserverlessGetPoliciesStats {
			opensearchserverless_GetPoliciesStats(cfg, client)
			return
		}
		if _opensearchserverlessGetSecurityConfig {
			opensearchserverless_GetSecurityConfig(cfg, client)
			return
		}
		if _opensearchserverlessGetSecurityPolicy {
			opensearchserverless_GetSecurityPolicy(cfg, client)
			return
		}
		if _opensearchserverlessListAccessPolicies {
			opensearchserverless_ListAccessPolicies(cfg, client)
			return
		}
		if _opensearchserverlessListCollectionGroups {
			opensearchserverless_ListCollectionGroups(cfg, client)
			return
		}
		if _opensearchserverlessListCollections {
			opensearchserverless_ListCollections(cfg, client)
			return
		}
		if _opensearchserverlessListLifecyclePolicies {
			opensearchserverless_ListLifecyclePolicies(cfg, client)
			return
		}
		if _opensearchserverlessListSecurityConfigs {
			opensearchserverless_ListSecurityConfigs(cfg, client)
			return
		}
		if _opensearchserverlessListSecurityPolicies {
			opensearchserverless_ListSecurityPolicies(cfg, client)
			return
		}
		if _opensearchserverlessListTagsForResource {
			opensearchserverless_ListTagsForResource(cfg, client)
			return
		}
		if _opensearchserverlessListVpcEndpoints {
			opensearchserverless_ListVpcEndpoints(cfg, client)
			return
		}
		if _opensearchserverlessTagResource {
			opensearchserverless_TagResource(cfg, client)
			return
		}
		if _opensearchserverlessUntagResource {
			opensearchserverless_UntagResource(cfg, client)
			return
		}
		if _opensearchserverlessUpdateAccessPolicy {
			opensearchserverless_UpdateAccessPolicy(cfg, client)
			return
		}
		if _opensearchserverlessUpdateAccountSettings {
			opensearchserverless_UpdateAccountSettings(cfg, client)
			return
		}
		if _opensearchserverlessUpdateCollection {
			opensearchserverless_UpdateCollection(cfg, client)
			return
		}
		if _opensearchserverlessUpdateCollectionGroup {
			opensearchserverless_UpdateCollectionGroup(cfg, client)
			return
		}
		if _opensearchserverlessUpdateIndex {
			opensearchserverless_UpdateIndex(cfg, client)
			return
		}
		if _opensearchserverlessUpdateLifecyclePolicy {
			opensearchserverless_UpdateLifecyclePolicy(cfg, client)
			return
		}
		if _opensearchserverlessUpdateSecurityConfig {
			opensearchserverless_UpdateSecurityConfig(cfg, client)
			return
		}
		if _opensearchserverlessUpdateSecurityPolicy {
			opensearchserverless_UpdateSecurityPolicy(cfg, client)
			return
		}
		if _opensearchserverlessUpdateVpcEndpoint {
			opensearchserverless_UpdateVpcEndpoint(cfg, client)
			return
		}

	},
}

var (
	_opensearchserverlessBatchGetCollection               bool
	_opensearchserverlessBatchGetCollectionGroup          bool
	_opensearchserverlessBatchGetEffectiveLifecyclePolicy bool
	_opensearchserverlessBatchGetLifecyclePolicy          bool
	_opensearchserverlessBatchGetVpcEndpoint              bool
	_opensearchserverlessCreateAccessPolicy               bool
	_opensearchserverlessCreateCollection                 bool
	_opensearchserverlessCreateCollectionGroup            bool
	_opensearchserverlessCreateIndex                      bool
	_opensearchserverlessCreateLifecyclePolicy            bool
	_opensearchserverlessCreateSecurityConfig             bool
	_opensearchserverlessCreateSecurityPolicy             bool
	_opensearchserverlessCreateVpcEndpoint                bool
	_opensearchserverlessDeleteAccessPolicy               bool
	_opensearchserverlessDeleteCollection                 bool
	_opensearchserverlessDeleteCollectionGroup            bool
	_opensearchserverlessDeleteIndex                      bool
	_opensearchserverlessDeleteLifecyclePolicy            bool
	_opensearchserverlessDeleteSecurityConfig             bool
	_opensearchserverlessDeleteSecurityPolicy             bool
	_opensearchserverlessDeleteVpcEndpoint                bool
	_opensearchserverlessGetAccessPolicy                  bool
	_opensearchserverlessGetAccountSettings               bool
	_opensearchserverlessGetIndex                         bool
	_opensearchserverlessGetPoliciesStats                 bool
	_opensearchserverlessGetSecurityConfig                bool
	_opensearchserverlessGetSecurityPolicy                bool
	_opensearchserverlessListAccessPolicies               bool
	_opensearchserverlessListCollectionGroups             bool
	_opensearchserverlessListCollections                  bool
	_opensearchserverlessListLifecyclePolicies            bool
	_opensearchserverlessListSecurityConfigs              bool
	_opensearchserverlessListSecurityPolicies             bool
	_opensearchserverlessListTagsForResource              bool
	_opensearchserverlessListVpcEndpoints                 bool
	_opensearchserverlessTagResource                      bool
	_opensearchserverlessUntagResource                    bool
	_opensearchserverlessUpdateAccessPolicy               bool
	_opensearchserverlessUpdateAccountSettings            bool
	_opensearchserverlessUpdateCollection                 bool
	_opensearchserverlessUpdateCollectionGroup            bool
	_opensearchserverlessUpdateIndex                      bool
	_opensearchserverlessUpdateLifecyclePolicy            bool
	_opensearchserverlessUpdateSecurityConfig             bool
	_opensearchserverlessUpdateSecurityPolicy             bool
	_opensearchserverlessUpdateVpcEndpoint                bool

	_opensearchserverlessAddSecurityGroupIds             []string
	_opensearchserverlessAddSubnetIds                    []string
	_opensearchserverlessCapacityLimits                  string
	_opensearchserverlessClientToken                     string
	_opensearchserverlessCollectionFilters               string
	_opensearchserverlessCollectionGroupName             string
	_opensearchserverlessConfigVersion                   string
	_opensearchserverlessDescription                     string
	_opensearchserverlessEncryptionConfig                string
	_opensearchserverlessIamFederationOptions            string
	_opensearchserverlessIamIdentityCenterOptions        string
	_opensearchserverlessIamIdentityCenterOptionsUpdates string
	_opensearchserverlessId                              string
	_opensearchserverlessIdentifiers                     string
	_opensearchserverlessIds                             []string
	_opensearchserverlessIndexName                       string
	_opensearchserverlessIndexSchema                     string
	_opensearchserverlessMaxResults                      string
	_opensearchserverlessName                            string
	_opensearchserverlessNames                           []string
	_opensearchserverlessNextToken                       string
	_opensearchserverlessPolicy                          string
	_opensearchserverlessPolicyVersion                   string
	_opensearchserverlessRemoveSecurityGroupIds          []string
	_opensearchserverlessRemoveSubnetIds                 []string
	_opensearchserverlessResource                        []string
	_opensearchserverlessResourceArn                     string
	_opensearchserverlessResourceIdentifiers             string
	_opensearchserverlessResources                       []string
	_opensearchserverlessSamlOptions                     string
	_opensearchserverlessSecurityGroupIds                []string
	_opensearchserverlessStandbyReplicas                 string
	_opensearchserverlessSubnetIds                       []string
	_opensearchserverlessTagKeys                         []string
	_opensearchserverlessTags                            string
	_opensearchserverlessType                            string
	_opensearchserverlessVectorOptions                   string
	_opensearchserverlessVpcEndpointFilters              string
	_opensearchserverlessVpcId                           string
)

// Returns attributes for one or more collections, including the collection
// endpoint, the OpenSearch Dashboards endpoint, and FIPS-compliant endpoints. For
// more information, see [Creating and managing Amazon OpenSearch Serverless collections].
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_BatchGetCollection(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.BatchGetCollectionInput{}

	if len(_opensearchserverlessIds) > 0 {
		input.Ids = append([]string(nil), _opensearchserverlessIds...)
	}
	if len(_opensearchserverlessNames) > 0 {
		input.Names = append([]string(nil), _opensearchserverlessNames...)
	}

	if resp, err := client.BatchGetCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns attributes for one or more collection groups, including capacity limits
// and the number of collections in each group. For more information, see [Creating and managing Amazon OpenSearch Serverless collections].
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_BatchGetCollectionGroup(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.BatchGetCollectionGroupInput{}

	if len(_opensearchserverlessIds) > 0 {
		input.Ids = append([]string(nil), _opensearchserverlessIds...)
	}
	if len(_opensearchserverlessNames) > 0 {
		input.Names = append([]string(nil), _opensearchserverlessNames...)
	}

	if resp, err := client.BatchGetCollectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of successful and failed retrievals for the OpenSearch
// Serverless indexes. For more information, see [Viewing data lifecycle policies].
//
// [Viewing data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-list
func opensearchserverless_BatchGetEffectiveLifecyclePolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.BatchGetEffectiveLifecyclePolicyInput{
		// ResourceIdentifiers: []types.LifecyclePolicyResourceIdentifier, // Required
	}

	if len(_opensearchserverlessResourceIdentifiers) > 0 {
		if err := assignInputField(input, "ResourceIdentifiers", _opensearchserverlessResourceIdentifiers); err != nil {
			log.Errorf("invalid --resource-identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetEffectiveLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns one or more configured OpenSearch Serverless lifecycle policies. For
// more information, see [Viewing data lifecycle policies].
//
// [Viewing data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-list
func opensearchserverless_BatchGetLifecyclePolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.BatchGetLifecyclePolicyInput{
		// Identifiers: []types.LifecyclePolicyIdentifier, // Required
	}

	if len(_opensearchserverlessIdentifiers) > 0 {
		if err := assignInputField(input, "Identifiers", _opensearchserverlessIdentifiers); err != nil {
			log.Errorf("invalid --identifiers: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchGetLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns attributes for one or more VPC endpoints associated with the current
// account. For more information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func opensearchserverless_BatchGetVpcEndpoint(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.BatchGetVpcEndpointInput{
		// Ids: []string, // Required
	}

	if len(_opensearchserverlessIds) > 0 {
		input.Ids = append([]string(nil), _opensearchserverlessIds...)
	}

	if resp, err := client.BatchGetVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a data access policy for OpenSearch Serverless. Access policies limit
// access to collections and the resources within them, and allow a user to access
// that data irrespective of the access mechanism or network source. For more
// information, see [Data access control for Amazon OpenSearch Serverless].
//
// [Data access control for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html
func opensearchserverless_CreateAccessPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateAccessPolicyInput{
		// Name: *string, // Required
		// Policy: *string, // Required
		// Type: types.AccessPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessPolicy) > 0 {
		input.Policy = aws.String(_opensearchserverlessPolicy)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}

	if resp, err := client.CreateAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new OpenSearch Serverless collection. For more information, see [Creating and managing Amazon OpenSearch Serverless collections].
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_CreateCollection(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateCollectionInput{
		// Name: *string, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessCollectionGroupName) > 0 {
		input.CollectionGroupName = aws.String(_opensearchserverlessCollectionGroupName)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessEncryptionConfig) > 0 {
		if err := assignInputField(input, "EncryptionConfig", _opensearchserverlessEncryptionConfig); err != nil {
			log.Errorf("invalid --encryption-config: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessStandbyReplicas) > 0 {
		if err := assignInputField(input, "StandbyReplicas", _opensearchserverlessStandbyReplicas); err != nil {
			log.Errorf("invalid --standby-replicas: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _opensearchserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessVectorOptions) > 0 {
		if err := assignInputField(input, "VectorOptions", _opensearchserverlessVectorOptions); err != nil {
			log.Errorf("invalid --vector-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a collection group within OpenSearch Serverless. Collection groups let
// you manage OpenSearch Compute Units (OCUs) at a group level, with multiple
// collections sharing the group's capacity limits.
//
// For more information, see [Managing collection groups].
//
// [Managing collection groups]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-collection-groups.html
func opensearchserverless_CreateCollectionGroup(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateCollectionGroupInput{
		// Name: *string, // Required
		// StandbyReplicas: types.StandbyReplicas, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessStandbyReplicas) > 0 {
		if err := assignInputField(input, "StandbyReplicas", _opensearchserverlessStandbyReplicas); err != nil {
			log.Errorf("invalid --standby-replicas: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessCapacityLimits) > 0 {
		if err := assignInputField(input, "CapacityLimits", _opensearchserverlessCapacityLimits); err != nil {
			log.Errorf("invalid --capacity-limits: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _opensearchserverlessTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCollectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an index within an OpenSearch Serverless collection. Unlike other
// OpenSearch indexes, indexes created by this API are automatically configured to
// conduct automatic semantic enrichment ingestion and search. For more
// information, see [About automatic semantic enrichment]in the OpenSearch User Guide.
//
// [About automatic semantic enrichment]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html#serverless-semantic-enrichment
func opensearchserverless_CreateIndex(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateIndexInput{
		// Id: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessIndexName) > 0 {
		input.IndexName = aws.String(_opensearchserverlessIndexName)
	}
	if len(_opensearchserverlessIndexSchema) > 0 {
		if err := assignInputField(input, "IndexSchema", _opensearchserverlessIndexSchema); err != nil {
			log.Errorf("invalid --index-schema: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a lifecyle policy to be applied to OpenSearch Serverless indexes.
// Lifecycle policies define the number of days or hours to retain the data on an
// OpenSearch Serverless index. For more information, see [Creating data lifecycle policies].
//
// [Creating data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-create
func opensearchserverless_CreateLifecyclePolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateLifecyclePolicyInput{
		// Name: *string, // Required
		// Policy: *string, // Required
		// Type: types.LifecyclePolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessPolicy) > 0 {
		input.Policy = aws.String(_opensearchserverlessPolicy)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}

	if resp, err := client.CreateLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Specifies a security configuration for OpenSearch Serverless. For more
// information, see [SAML authentication for Amazon OpenSearch Serverless].
//
// [SAML authentication for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html
func opensearchserverless_CreateSecurityConfig(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateSecurityConfigInput{
		// Name: *string, // Required
		// Type: types.SecurityConfigType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessIamFederationOptions) > 0 {
		if err := assignInputField(input, "IamFederationOptions", _opensearchserverlessIamFederationOptions); err != nil {
			log.Errorf("invalid --iam-federation-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessIamIdentityCenterOptions) > 0 {
		if err := assignInputField(input, "IamIdentityCenterOptions", _opensearchserverlessIamIdentityCenterOptions); err != nil {
			log.Errorf("invalid --iam-identity-center-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessSamlOptions) > 0 {
		if err := assignInputField(input, "SamlOptions", _opensearchserverlessSamlOptions); err != nil {
			log.Errorf("invalid --saml-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSecurityConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a security policy to be used by one or more OpenSearch Serverless
// collections. Security policies provide access to a collection and its OpenSearch
// Dashboards endpoint from public networks or specific VPC endpoints. They also
// allow you to secure a collection with a KMS encryption key. For more
// information, see [Network access for Amazon OpenSearch Serverless]and [Encryption at rest for Amazon OpenSearch Serverless].
//
// [Network access for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html
// [Encryption at rest for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html
func opensearchserverless_CreateSecurityPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateSecurityPolicyInput{
		// Name: *string, // Required
		// Policy: *string, // Required
		// Type: types.SecurityPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessPolicy) > 0 {
		input.Policy = aws.String(_opensearchserverlessPolicy)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}

	if resp, err := client.CreateSecurityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an OpenSearch Serverless-managed interface VPC endpoint. For more
// information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func opensearchserverless_CreateVpcEndpoint(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.CreateVpcEndpointInput{
		// Name: *string, // Required
		// SubnetIds: []string, // Required
		// VpcId: *string, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessSubnetIds) > 0 {
		input.SubnetIds = append([]string(nil), _opensearchserverlessSubnetIds...)
	}
	if len(_opensearchserverlessVpcId) > 0 {
		input.VpcId = aws.String(_opensearchserverlessVpcId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessSecurityGroupIds) > 0 {
		input.SecurityGroupIds = append([]string(nil), _opensearchserverlessSecurityGroupIds...)
	}

	if resp, err := client.CreateVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch Serverless access policy. For more information, see [Data access control for Amazon OpenSearch Serverless].
//
// [Data access control for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html
func opensearchserverless_DeleteAccessPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteAccessPolicyInput{
		// Name: *string, // Required
		// Type: types.AccessPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch Serverless collection. For more information, see [Creating and managing Amazon OpenSearch Serverless collections].
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_DeleteCollection(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteCollectionInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a collection group. You can only delete empty collection groups that
// contain no collections. For more information, see [Creating and managing Amazon OpenSearch Serverless collections].
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_DeleteCollectionGroup(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteCollectionGroupInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteCollectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an index from an OpenSearch Serverless collection. Be aware that the
// index might be configured to conduct automatic semantic enrichment ingestion and
// search. For more information, see [About automatic semantic enrichment].
//
// [About automatic semantic enrichment]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html#serverless-semantic-enrichment
func opensearchserverless_DeleteIndex(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteIndexInput{
		// Id: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessIndexName) > 0 {
		input.IndexName = aws.String(_opensearchserverlessIndexName)
	}

	if resp, err := client.DeleteIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch Serverless lifecycle policy. For more information, see [Deleting data lifecycle policies].
//
// [Deleting data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-delete
func opensearchserverless_DeleteLifecyclePolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteLifecyclePolicyInput{
		// Name: *string, // Required
		// Type: types.LifecyclePolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a security configuration for OpenSearch Serverless. For more
// information, see [SAML authentication for Amazon OpenSearch Serverless].
//
// [SAML authentication for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html
func opensearchserverless_DeleteSecurityConfig(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteSecurityConfigInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteSecurityConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch Serverless security policy.
func opensearchserverless_DeleteSecurityPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteSecurityPolicyInput{
		// Name: *string, // Required
		// Type: types.SecurityPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteSecurityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch Serverless-managed interface endpoint. For more
// information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func opensearchserverless_DeleteVpcEndpoint(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.DeleteVpcEndpointInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}

	if resp, err := client.DeleteVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns an OpenSearch Serverless access policy. For more information, see [Data access control for Amazon OpenSearch Serverless].
//
// [Data access control for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html
func opensearchserverless_GetAccessPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.GetAccessPolicyInput{
		// Name: *string, // Required
		// Type: types.AccessPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns account-level settings related to OpenSearch Serverless.
func opensearchserverless_GetAccountSettings(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.GetAccountSettingsInput{}

	if resp, err := client.GetAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an index in an OpenSearch Serverless collection,
// including its schema definition. The index might be configured to conduct
// automatic semantic enrichment ingestion and search. For more information, see [About automatic semantic enrichment].
//
// [About automatic semantic enrichment]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html#serverless-semantic-enrichment
func opensearchserverless_GetIndex(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.GetIndexInput{
		// Id: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessIndexName) > 0 {
		input.IndexName = aws.String(_opensearchserverlessIndexName)
	}

	if resp, err := client.GetIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns statistical information about your OpenSearch Serverless access
// policies, security configurations, and security policies.
func opensearchserverless_GetPoliciesStats(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.GetPoliciesStatsInput{}

	if resp, err := client.GetPoliciesStats(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about an OpenSearch Serverless security configuration. For
// more information, see [SAML authentication for Amazon OpenSearch Serverless].
//
// [SAML authentication for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html
func opensearchserverless_GetSecurityConfig(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.GetSecurityConfigInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}

	if resp, err := client.GetSecurityConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a configured OpenSearch Serverless security policy.
// For more information, see [Network access for Amazon OpenSearch Serverless]and [Encryption at rest for Amazon OpenSearch Serverless].
//
// [Network access for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html
// [Encryption at rest for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html
func opensearchserverless_GetSecurityPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.GetSecurityPolicyInput{
		// Name: *string, // Required
		// Type: types.SecurityPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetSecurityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a list of OpenSearch Serverless access policies.
func opensearchserverless_ListAccessPolicies(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListAccessPoliciesInput{
		// Type: types.AccessPolicyType, // Required
	}

	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}
	if len(_opensearchserverlessResource) > 0 {
		input.Resource = append([]string(nil), _opensearchserverlessResource...)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListAccessPoliciesOutput
	p := opensearchserverless.NewListAccessPoliciesPaginator(client, input)
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

// Returns a list of collection groups. For more information, see [Creating and managing Amazon OpenSearch Serverless collections].
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_ListCollectionGroups(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListCollectionGroupsInput{}

	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollectionGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListCollectionGroupsOutput
	p := opensearchserverless.NewListCollectionGroupsPaginator(client, input)
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

// Lists all OpenSearch Serverless collections. For more information, see [Creating and managing Amazon OpenSearch Serverless collections].
// Make sure to include an empty request body {} if you don't include any
// collection filters in the request.
//
// [Creating and managing Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html
func opensearchserverless_ListCollections(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListCollectionsInput{}

	if len(_opensearchserverlessCollectionFilters) > 0 {
		if err := assignInputField(input, "CollectionFilters", _opensearchserverlessCollectionFilters); err != nil {
			log.Errorf("invalid --collection-filters: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCollections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListCollectionsOutput
	p := opensearchserverless.NewListCollectionsPaginator(client, input)
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

// Returns a list of OpenSearch Serverless lifecycle policies. For more
// information, see [Viewing data lifecycle policies].
//
// [Viewing data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-list
func opensearchserverless_ListLifecyclePolicies(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListLifecyclePoliciesInput{
		// Type: types.LifecyclePolicyType, // Required
	}

	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}
	if len(_opensearchserverlessResources) > 0 {
		input.Resources = append([]string(nil), _opensearchserverlessResources...)
	}

	if disablePaginator() {
		if resp, err := client.ListLifecyclePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListLifecyclePoliciesOutput
	p := opensearchserverless.NewListLifecyclePoliciesPaginator(client, input)
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

// Returns information about configured OpenSearch Serverless security
// configurations. For more information, see [SAML authentication for Amazon OpenSearch Serverless].
//
// [SAML authentication for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html
func opensearchserverless_ListSecurityConfigs(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListSecurityConfigsInput{
		// Type: types.SecurityConfigType, // Required
	}

	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityConfigs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListSecurityConfigsOutput
	p := opensearchserverless.NewListSecurityConfigsPaginator(client, input)
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

// Returns information about configured OpenSearch Serverless security policies.
func opensearchserverless_ListSecurityPolicies(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListSecurityPoliciesInput{
		// Type: types.SecurityPolicyType, // Required
	}

	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}
	if len(_opensearchserverlessResource) > 0 {
		input.Resource = append([]string(nil), _opensearchserverlessResource...)
	}

	if disablePaginator() {
		if resp, err := client.ListSecurityPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListSecurityPoliciesOutput
	p := opensearchserverless.NewListSecurityPoliciesPaginator(client, input)
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

// Returns the tags for an OpenSearch Serverless resource. For more information,
// see [Tagging Amazon OpenSearch Serverless collections].
//
// [Tagging Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/tag-collection.html
func opensearchserverless_ListTagsForResource(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_opensearchserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_opensearchserverlessResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the OpenSearch Serverless-managed interface VPC endpoints associated
// with the current account. For more information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func opensearchserverless_ListVpcEndpoints(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.ListVpcEndpointsInput{}

	if len(_opensearchserverlessMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchserverlessMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessNextToken) > 0 {
		input.NextToken = aws.String(_opensearchserverlessNextToken)
	}
	if len(_opensearchserverlessVpcEndpointFilters) > 0 {
		if err := assignInputField(input, "VpcEndpointFilters", _opensearchserverlessVpcEndpointFilters); err != nil {
			log.Errorf("invalid --vpc-endpoint-filters: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListVpcEndpoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearchserverless.ListVpcEndpointsOutput
	p := opensearchserverless.NewListVpcEndpointsPaginator(client, input)
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

// Associates tags with an OpenSearch Serverless resource. For more information,
// see [Tagging Amazon OpenSearch Serverless collections].
//
// [Tagging Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/tag-collection.html
func opensearchserverless_TagResource(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_opensearchserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_opensearchserverlessResourceArn)
	}
	if len(_opensearchserverlessTags) > 0 {
		if err := assignInputField(input, "Tags", _opensearchserverlessTags); err != nil {
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

// Removes a tag or set of tags from an OpenSearch Serverless resource. For more
// information, see [Tagging Amazon OpenSearch Serverless collections].
//
// [Tagging Amazon OpenSearch Serverless collections]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/tag-collection.html
func opensearchserverless_UntagResource(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_opensearchserverlessResourceArn) > 0 {
		input.ResourceArn = aws.String(_opensearchserverlessResourceArn)
	}
	if len(_opensearchserverlessTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _opensearchserverlessTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an OpenSearch Serverless access policy. For more information, see [Data access control for Amazon OpenSearch Serverless].
//
// [Data access control for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-data-access.html
func opensearchserverless_UpdateAccessPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateAccessPolicyInput{
		// Name: *string, // Required
		// PolicyVersion: *string, // Required
		// Type: types.AccessPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessPolicyVersion) > 0 {
		input.PolicyVersion = aws.String(_opensearchserverlessPolicyVersion)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessPolicy) > 0 {
		input.Policy = aws.String(_opensearchserverlessPolicy)
	}

	if resp, err := client.UpdateAccessPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update the OpenSearch Serverless settings for the current Amazon Web Services
// account. For more information, see [Managing capacity limits for Amazon OpenSearch Serverless].
//
// [Managing capacity limits for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-scaling.html
func opensearchserverless_UpdateAccountSettings(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateAccountSettingsInput{}

	if len(_opensearchserverlessCapacityLimits) > 0 {
		if err := assignInputField(input, "CapacityLimits", _opensearchserverlessCapacityLimits); err != nil {
			log.Errorf("invalid --capacity-limits: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateAccountSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an OpenSearch Serverless collection.
func opensearchserverless_UpdateCollection(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateCollectionInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}

	if resp, err := client.UpdateCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the description and capacity limits of a collection group.
func opensearchserverless_UpdateCollectionGroup(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateCollectionGroupInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessCapacityLimits) > 0 {
		if err := assignInputField(input, "CapacityLimits", _opensearchserverlessCapacityLimits); err != nil {
			log.Errorf("invalid --capacity-limits: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}

	if resp, err := client.UpdateCollectionGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing index in an OpenSearch Serverless collection. This
// operation allows you to modify the index schema, including adding new fields or
// changing field mappings. You can also enable automatic semantic enrichment
// ingestion and search. For more information, see [About automatic semantic enrichment].
//
// [About automatic semantic enrichment]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-manage.html#serverless-semantic-enrichment
func opensearchserverless_UpdateIndex(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateIndexInput{
		// Id: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessIndexName) > 0 {
		input.IndexName = aws.String(_opensearchserverlessIndexName)
	}
	if len(_opensearchserverlessIndexSchema) > 0 {
		if err := assignInputField(input, "IndexSchema", _opensearchserverlessIndexSchema); err != nil {
			log.Errorf("invalid --index-schema: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an OpenSearch Serverless access policy. For more information, see [Updating data lifecycle policies].
//
// [Updating data lifecycle policies]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-lifecycle.html#serverless-lifecycle-update
func opensearchserverless_UpdateLifecyclePolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateLifecyclePolicyInput{
		// Name: *string, // Required
		// PolicyVersion: *string, // Required
		// Type: types.LifecyclePolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessPolicyVersion) > 0 {
		input.PolicyVersion = aws.String(_opensearchserverlessPolicyVersion)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessPolicy) > 0 {
		input.Policy = aws.String(_opensearchserverlessPolicy)
	}

	if resp, err := client.UpdateLifecyclePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a security configuration for OpenSearch Serverless. For more
// information, see [SAML authentication for Amazon OpenSearch Serverless].
//
// [SAML authentication for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-saml.html
func opensearchserverless_UpdateSecurityConfig(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateSecurityConfigInput{
		// ConfigVersion: *string, // Required
		// Id: *string, // Required
	}

	if len(_opensearchserverlessConfigVersion) > 0 {
		input.ConfigVersion = aws.String(_opensearchserverlessConfigVersion)
	}
	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessIamFederationOptions) > 0 {
		if err := assignInputField(input, "IamFederationOptions", _opensearchserverlessIamFederationOptions); err != nil {
			log.Errorf("invalid --iam-federation-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessIamIdentityCenterOptionsUpdates) > 0 {
		if err := assignInputField(input, "IamIdentityCenterOptionsUpdates", _opensearchserverlessIamIdentityCenterOptionsUpdates); err != nil {
			log.Errorf("invalid --iam-identity-center-options-updates: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessSamlOptions) > 0 {
		if err := assignInputField(input, "SamlOptions", _opensearchserverlessSamlOptions); err != nil {
			log.Errorf("invalid --saml-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateSecurityConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an OpenSearch Serverless security policy. For more information, see [Network access for Amazon OpenSearch Serverless]
// and [Encryption at rest for Amazon OpenSearch Serverless].
//
// [Encryption at rest for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-encryption.html
// [Network access for Amazon OpenSearch Serverless]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-network.html
func opensearchserverless_UpdateSecurityPolicy(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateSecurityPolicyInput{
		// Name: *string, // Required
		// PolicyVersion: *string, // Required
		// Type: types.SecurityPolicyType, // Required
	}

	if len(_opensearchserverlessName) > 0 {
		input.Name = aws.String(_opensearchserverlessName)
	}
	if len(_opensearchserverlessPolicyVersion) > 0 {
		input.PolicyVersion = aws.String(_opensearchserverlessPolicyVersion)
	}
	if len(_opensearchserverlessType) > 0 {
		if err := assignInputField(input, "Type", _opensearchserverlessType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessDescription) > 0 {
		input.Description = aws.String(_opensearchserverlessDescription)
	}
	if len(_opensearchserverlessPolicy) > 0 {
		input.Policy = aws.String(_opensearchserverlessPolicy)
	}

	if resp, err := client.UpdateSecurityPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an OpenSearch Serverless-managed interface endpoint. For more
// information, see [Access Amazon OpenSearch Serverless using an interface endpoint].
//
// [Access Amazon OpenSearch Serverless using an interface endpoint]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/serverless-vpc.html
func opensearchserverless_UpdateVpcEndpoint(cfg aws.Config, client *opensearchserverless.Client) {
	input := &opensearchserverless.UpdateVpcEndpointInput{
		// Id: *string, // Required
	}

	if len(_opensearchserverlessId) > 0 {
		input.Id = aws.String(_opensearchserverlessId)
	}
	if len(_opensearchserverlessAddSecurityGroupIds) > 0 {
		input.AddSecurityGroupIds = append([]string(nil), _opensearchserverlessAddSecurityGroupIds...)
	}
	if len(_opensearchserverlessAddSubnetIds) > 0 {
		input.AddSubnetIds = append([]string(nil), _opensearchserverlessAddSubnetIds...)
	}
	if len(_opensearchserverlessClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchserverlessClientToken)
	}
	if len(_opensearchserverlessRemoveSecurityGroupIds) > 0 {
		input.RemoveSecurityGroupIds = append([]string(nil), _opensearchserverlessRemoveSecurityGroupIds...)
	}
	if len(_opensearchserverlessRemoveSubnetIds) > 0 {
		input.RemoveSubnetIds = append([]string(nil), _opensearchserverlessRemoveSubnetIds...)
	}

	if resp, err := client.UpdateVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_opensearchserverlessCmd)
	_opensearchserverlessCmd.Flags().SortFlags = false

	_opensearchserverlessCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_opensearchserverlessCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_opensearchserverlessCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessAddSecurityGroupIds, "add-security-group-ids", "", nil, "Add Security Group Ids")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessAddSubnetIds, "add-subnet-ids", "", nil, "Add Subnet Ids")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessCapacityLimits, "capacity-limits", "", "", "Capacity Limits")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessClientToken, "client-token", "", "", "Client Token")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessCollectionFilters, "collection-filters", "", "", "Collection Filters")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessCollectionGroupName, "collection-group-name", "", "", "Collection Group Name")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessConfigVersion, "config-version", "", "", "Config Version")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessDescription, "description", "", "", "Description")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessEncryptionConfig, "encryption-config", "", "", "Encryption Config")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessIamFederationOptions, "iam-federation-options", "", "", "IAM Federation Options")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessIamIdentityCenterOptions, "iam-identity-center-options", "", "", "IAM Identity Center Options")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessIamIdentityCenterOptionsUpdates, "iam-identity-center-options-updates", "", "", "IAM Identity Center Options Updates")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessId, "id", "", "", "ID")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessIdentifiers, "identifiers", "", "", "Identifiers")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessIds, "ids", "", nil, "Ids")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessIndexName, "index-name", "", "", "Index Name")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessIndexSchema, "index-schema", "", "", "Index Schema")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessMaxResults, "max-results", "", "", "Max Results")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessName, "name", "", "", "Name")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessNames, "names", "", nil, "Names")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessNextToken, "next-token", "", "", "Next Token")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessPolicy, "policy", "", "", "Policy")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessPolicyVersion, "policy-version", "", "", "Policy Version")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessRemoveSecurityGroupIds, "remove-security-group-ids", "", nil, "Remove Security Group Ids")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessRemoveSubnetIds, "remove-subnet-ids", "", nil, "Remove Subnet Ids")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessResource, "resource", "", nil, "Resource")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessResourceArn, "resource-arn", "", "", "Resource ARN")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessResourceIdentifiers, "resource-identifiers", "", "", "Resource Identifiers")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessResources, "resources", "", nil, "Resources")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessSamlOptions, "saml-options", "", "", "Saml Options")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessSecurityGroupIds, "security-group-ids", "", nil, "Security Group Ids")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessStandbyReplicas, "standby-replicas", "", "", "Standby Replicas")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessSubnetIds, "subnet-ids", "", nil, "Subnet Ids")
	_opensearchserverlessCmd.Flags().StringSliceVarP(&_opensearchserverlessTagKeys, "tag-keys", "", nil, "Tag Keys")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessTags, "tags", "", "", "Tags")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessType, "type", "", "", "Type")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessVectorOptions, "vector-options", "", "", "Vector Options")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessVpcEndpointFilters, "vpc-endpoint-filters", "", "", "VPC Endpoint Filters")
	_opensearchserverlessCmd.Flags().StringVarP(&_opensearchserverlessVpcId, "vpc-id", "", "", "VPC ID")

	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessBatchGetCollection, "batch-get-collection", "", false, "Batch Get Collection")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessBatchGetCollectionGroup, "batch-get-collection-group", "", false, "Batch Get Collection Group")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessBatchGetEffectiveLifecyclePolicy, "batch-get-effective-lifecycle-policy", "", false, "Batch Get Effective Lifecycle Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessBatchGetLifecyclePolicy, "batch-get-lifecycle-policy", "", false, "Batch Get Lifecycle Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessBatchGetVpcEndpoint, "batch-get-vpc-endpoint", "", false, "Batch Get VPC Endpoint")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateAccessPolicy, "create-access-policy", "", false, "Create Access Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateCollection, "create-collection", "", false, "Create Collection")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateCollectionGroup, "create-collection-group", "", false, "Create Collection Group")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateIndex, "create-index", "", false, "Create Index")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateLifecyclePolicy, "create-lifecycle-policy", "", false, "Create Lifecycle Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateSecurityConfig, "create-security-config", "", false, "Create Security Config")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateSecurityPolicy, "create-security-policy", "", false, "Create Security Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessCreateVpcEndpoint, "create-vpc-endpoint", "", false, "Create VPC Endpoint")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteAccessPolicy, "delete-access-policy", "", false, "Delete Access Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteCollection, "delete-collection", "", false, "Delete Collection")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteCollectionGroup, "delete-collection-group", "", false, "Delete Collection Group")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteIndex, "delete-index", "", false, "Delete Index")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteLifecyclePolicy, "delete-lifecycle-policy", "", false, "Delete Lifecycle Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteSecurityConfig, "delete-security-config", "", false, "Delete Security Config")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteSecurityPolicy, "delete-security-policy", "", false, "Delete Security Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessDeleteVpcEndpoint, "delete-vpc-endpoint", "", false, "Delete VPC Endpoint")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessGetAccessPolicy, "get-access-policy", "", false, "Get Access Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessGetAccountSettings, "get-account-settings", "", false, "Get Account Settings")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessGetIndex, "get-index", "", false, "Get Index")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessGetPoliciesStats, "get-policies-stats", "", false, "Get Policies Stats")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessGetSecurityConfig, "get-security-config", "", false, "Get Security Config")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessGetSecurityPolicy, "get-security-policy", "", false, "Get Security Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListAccessPolicies, "list-access-policies", "", false, "List Access Policies")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListCollectionGroups, "list-collection-groups", "", false, "List Collection Groups")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListCollections, "list-collections", "", false, "List Collections")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListLifecyclePolicies, "list-lifecycle-policies", "", false, "List Lifecycle Policies")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListSecurityConfigs, "list-security-configs", "", false, "List Security Configs")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListSecurityPolicies, "list-security-policies", "", false, "List Security Policies")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessListVpcEndpoints, "list-vpc-endpoints", "", false, "List VPC Endpoints")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessTagResource, "tag-resource", "", false, "Tag Resource")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUntagResource, "untag-resource", "", false, "Untag Resource")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateAccessPolicy, "update-access-policy", "", false, "Update Access Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateAccountSettings, "update-account-settings", "", false, "Update Account Settings")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateCollection, "update-collection", "", false, "Update Collection")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateCollectionGroup, "update-collection-group", "", false, "Update Collection Group")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateIndex, "update-index", "", false, "Update Index")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateLifecyclePolicy, "update-lifecycle-policy", "", false, "Update Lifecycle Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateSecurityConfig, "update-security-config", "", false, "Update Security Config")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateSecurityPolicy, "update-security-policy", "", false, "Update Security Policy")
	_opensearchserverlessCmd.Flags().BoolVarP(&_opensearchserverlessUpdateVpcEndpoint, "update-vpc-endpoint", "", false, "Update VPC Endpoint")

}
