package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// elasticsearchserviceCmd represents the elasticsearchservice command
var _elasticsearchserviceCmd = &cobra.Command{
	Use:   "elasticsearchservice",
	Short: "AWS elasticsearchservice CLI",
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
		client := elasticsearchservice.NewFromConfig(cfg)
		if _elasticsearchserviceAcceptInboundCrossClusterSearchConnection {
			elasticsearchservice_AcceptInboundCrossClusterSearchConnection(cfg, client)
			return
		}
		if _elasticsearchserviceAddTags {
			elasticsearchservice_AddTags(cfg, client)
			return
		}
		if _elasticsearchserviceAssociatePackage {
			elasticsearchservice_AssociatePackage(cfg, client)
			return
		}
		if _elasticsearchserviceAuthorizeVpcEndpointAccess {
			elasticsearchservice_AuthorizeVpcEndpointAccess(cfg, client)
			return
		}
		if _elasticsearchserviceCancelDomainConfigChange {
			elasticsearchservice_CancelDomainConfigChange(cfg, client)
			return
		}
		if _elasticsearchserviceCancelElasticsearchServiceSoftwareUpdate {
			elasticsearchservice_CancelElasticsearchServiceSoftwareUpdate(cfg, client)
			return
		}
		if _elasticsearchserviceCreateElasticsearchDomain {
			elasticsearchservice_CreateElasticsearchDomain(cfg, client)
			return
		}
		if _elasticsearchserviceCreateOutboundCrossClusterSearchConnection {
			elasticsearchservice_CreateOutboundCrossClusterSearchConnection(cfg, client)
			return
		}
		if _elasticsearchserviceCreatePackage {
			elasticsearchservice_CreatePackage(cfg, client)
			return
		}
		if _elasticsearchserviceCreateVpcEndpoint {
			elasticsearchservice_CreateVpcEndpoint(cfg, client)
			return
		}
		if _elasticsearchserviceDeleteElasticsearchDomain {
			elasticsearchservice_DeleteElasticsearchDomain(cfg, client)
			return
		}
		if _elasticsearchserviceDeleteElasticsearchServiceRole {
			elasticsearchservice_DeleteElasticsearchServiceRole(cfg, client)
			return
		}
		if _elasticsearchserviceDeleteInboundCrossClusterSearchConnection {
			elasticsearchservice_DeleteInboundCrossClusterSearchConnection(cfg, client)
			return
		}
		if _elasticsearchserviceDeleteOutboundCrossClusterSearchConnection {
			elasticsearchservice_DeleteOutboundCrossClusterSearchConnection(cfg, client)
			return
		}
		if _elasticsearchserviceDeletePackage {
			elasticsearchservice_DeletePackage(cfg, client)
			return
		}
		if _elasticsearchserviceDeleteVpcEndpoint {
			elasticsearchservice_DeleteVpcEndpoint(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeDomainAutoTunes {
			elasticsearchservice_DescribeDomainAutoTunes(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeDomainChangeProgress {
			elasticsearchservice_DescribeDomainChangeProgress(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeElasticsearchDomain {
			elasticsearchservice_DescribeElasticsearchDomain(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeElasticsearchDomainConfig {
			elasticsearchservice_DescribeElasticsearchDomainConfig(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeElasticsearchDomains {
			elasticsearchservice_DescribeElasticsearchDomains(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeElasticsearchInstanceTypeLimits {
			elasticsearchservice_DescribeElasticsearchInstanceTypeLimits(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeInboundCrossClusterSearchConnections {
			elasticsearchservice_DescribeInboundCrossClusterSearchConnections(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeOutboundCrossClusterSearchConnections {
			elasticsearchservice_DescribeOutboundCrossClusterSearchConnections(cfg, client)
			return
		}
		if _elasticsearchserviceDescribePackages {
			elasticsearchservice_DescribePackages(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeReservedElasticsearchInstanceOfferings {
			elasticsearchservice_DescribeReservedElasticsearchInstanceOfferings(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeReservedElasticsearchInstances {
			elasticsearchservice_DescribeReservedElasticsearchInstances(cfg, client)
			return
		}
		if _elasticsearchserviceDescribeVpcEndpoints {
			elasticsearchservice_DescribeVpcEndpoints(cfg, client)
			return
		}
		if _elasticsearchserviceDissociatePackage {
			elasticsearchservice_DissociatePackage(cfg, client)
			return
		}
		if _elasticsearchserviceGetCompatibleElasticsearchVersions {
			elasticsearchservice_GetCompatibleElasticsearchVersions(cfg, client)
			return
		}
		if _elasticsearchserviceGetPackageVersionHistory {
			elasticsearchservice_GetPackageVersionHistory(cfg, client)
			return
		}
		if _elasticsearchserviceGetUpgradeHistory {
			elasticsearchservice_GetUpgradeHistory(cfg, client)
			return
		}
		if _elasticsearchserviceGetUpgradeStatus {
			elasticsearchservice_GetUpgradeStatus(cfg, client)
			return
		}
		if _elasticsearchserviceListDomainNames {
			elasticsearchservice_ListDomainNames(cfg, client)
			return
		}
		if _elasticsearchserviceListDomainsForPackage {
			elasticsearchservice_ListDomainsForPackage(cfg, client)
			return
		}
		if _elasticsearchserviceListElasticsearchInstanceTypes {
			elasticsearchservice_ListElasticsearchInstanceTypes(cfg, client)
			return
		}
		if _elasticsearchserviceListElasticsearchVersions {
			elasticsearchservice_ListElasticsearchVersions(cfg, client)
			return
		}
		if _elasticsearchserviceListPackagesForDomain {
			elasticsearchservice_ListPackagesForDomain(cfg, client)
			return
		}
		if _elasticsearchserviceListTags {
			elasticsearchservice_ListTags(cfg, client)
			return
		}
		if _elasticsearchserviceListVpcEndpointAccess {
			elasticsearchservice_ListVpcEndpointAccess(cfg, client)
			return
		}
		if _elasticsearchserviceListVpcEndpoints {
			elasticsearchservice_ListVpcEndpoints(cfg, client)
			return
		}
		if _elasticsearchserviceListVpcEndpointsForDomain {
			elasticsearchservice_ListVpcEndpointsForDomain(cfg, client)
			return
		}
		if _elasticsearchservicePurchaseReservedElasticsearchInstanceOffering {
			elasticsearchservice_PurchaseReservedElasticsearchInstanceOffering(cfg, client)
			return
		}
		if _elasticsearchserviceRejectInboundCrossClusterSearchConnection {
			elasticsearchservice_RejectInboundCrossClusterSearchConnection(cfg, client)
			return
		}
		if _elasticsearchserviceRemoveTags {
			elasticsearchservice_RemoveTags(cfg, client)
			return
		}
		if _elasticsearchserviceRevokeVpcEndpointAccess {
			elasticsearchservice_RevokeVpcEndpointAccess(cfg, client)
			return
		}
		if _elasticsearchserviceStartElasticsearchServiceSoftwareUpdate {
			elasticsearchservice_StartElasticsearchServiceSoftwareUpdate(cfg, client)
			return
		}
		if _elasticsearchserviceUpdateElasticsearchDomainConfig {
			elasticsearchservice_UpdateElasticsearchDomainConfig(cfg, client)
			return
		}
		if _elasticsearchserviceUpdatePackage {
			elasticsearchservice_UpdatePackage(cfg, client)
			return
		}
		if _elasticsearchserviceUpdateVpcEndpoint {
			elasticsearchservice_UpdateVpcEndpoint(cfg, client)
			return
		}
		if _elasticsearchserviceUpgradeElasticsearchDomain {
			elasticsearchservice_UpgradeElasticsearchDomain(cfg, client)
			return
		}

	},
}

var (
	_elasticsearchserviceAcceptInboundCrossClusterSearchConnection      bool
	_elasticsearchserviceAddTags                                        bool
	_elasticsearchserviceAssociatePackage                               bool
	_elasticsearchserviceAuthorizeVpcEndpointAccess                     bool
	_elasticsearchserviceCancelDomainConfigChange                       bool
	_elasticsearchserviceCancelElasticsearchServiceSoftwareUpdate       bool
	_elasticsearchserviceCreateElasticsearchDomain                      bool
	_elasticsearchserviceCreateOutboundCrossClusterSearchConnection     bool
	_elasticsearchserviceCreatePackage                                  bool
	_elasticsearchserviceCreateVpcEndpoint                              bool
	_elasticsearchserviceDeleteElasticsearchDomain                      bool
	_elasticsearchserviceDeleteElasticsearchServiceRole                 bool
	_elasticsearchserviceDeleteInboundCrossClusterSearchConnection      bool
	_elasticsearchserviceDeleteOutboundCrossClusterSearchConnection     bool
	_elasticsearchserviceDeletePackage                                  bool
	_elasticsearchserviceDeleteVpcEndpoint                              bool
	_elasticsearchserviceDescribeDomainAutoTunes                        bool
	_elasticsearchserviceDescribeDomainChangeProgress                   bool
	_elasticsearchserviceDescribeElasticsearchDomain                    bool
	_elasticsearchserviceDescribeElasticsearchDomainConfig              bool
	_elasticsearchserviceDescribeElasticsearchDomains                   bool
	_elasticsearchserviceDescribeElasticsearchInstanceTypeLimits        bool
	_elasticsearchserviceDescribeInboundCrossClusterSearchConnections   bool
	_elasticsearchserviceDescribeOutboundCrossClusterSearchConnections  bool
	_elasticsearchserviceDescribePackages                               bool
	_elasticsearchserviceDescribeReservedElasticsearchInstanceOfferings bool
	_elasticsearchserviceDescribeReservedElasticsearchInstances         bool
	_elasticsearchserviceDescribeVpcEndpoints                           bool
	_elasticsearchserviceDissociatePackage                              bool
	_elasticsearchserviceGetCompatibleElasticsearchVersions             bool
	_elasticsearchserviceGetPackageVersionHistory                       bool
	_elasticsearchserviceGetUpgradeHistory                              bool
	_elasticsearchserviceGetUpgradeStatus                               bool
	_elasticsearchserviceListDomainNames                                bool
	_elasticsearchserviceListDomainsForPackage                          bool
	_elasticsearchserviceListElasticsearchInstanceTypes                 bool
	_elasticsearchserviceListElasticsearchVersions                      bool
	_elasticsearchserviceListPackagesForDomain                          bool
	_elasticsearchserviceListTags                                       bool
	_elasticsearchserviceListVpcEndpointAccess                          bool
	_elasticsearchserviceListVpcEndpoints                               bool
	_elasticsearchserviceListVpcEndpointsForDomain                      bool
	_elasticsearchservicePurchaseReservedElasticsearchInstanceOffering  bool
	_elasticsearchserviceRejectInboundCrossClusterSearchConnection      bool
	_elasticsearchserviceRemoveTags                                     bool
	_elasticsearchserviceRevokeVpcEndpointAccess                        bool
	_elasticsearchserviceStartElasticsearchServiceSoftwareUpdate        bool
	_elasticsearchserviceUpdateElasticsearchDomainConfig                bool
	_elasticsearchserviceUpdatePackage                                  bool
	_elasticsearchserviceUpdateVpcEndpoint                              bool
	_elasticsearchserviceUpgradeElasticsearchDomain                     bool

	_elasticsearchserviceAccessPolicies                          string
	_elasticsearchserviceAccount                                 string
	_elasticsearchserviceAdvancedOptions                         string
	_elasticsearchserviceAdvancedSecurityOptions                 string
	_elasticsearchserviceARN                                     string
	_elasticsearchserviceAutoTuneOptions                         string
	_elasticsearchserviceChangeId                                string
	_elasticsearchserviceClientToken                             string
	_elasticsearchserviceCognitoOptions                          string
	_elasticsearchserviceCommitMessage                           string
	_elasticsearchserviceConnectionAlias                         string
	_elasticsearchserviceCrossClusterSearchConnectionId          string
	_elasticsearchserviceDestinationDomainInfo                   string
	_elasticsearchserviceDomainArn                               string
	_elasticsearchserviceDomainEndpointOptions                   string
	_elasticsearchserviceDomainName                              string
	_elasticsearchserviceDomainNames                             []string
	_elasticsearchserviceDryRun                                  string
	_elasticsearchserviceEBSOptions                              string
	_elasticsearchserviceElasticsearchClusterConfig              string
	_elasticsearchserviceElasticsearchVersion                    string
	_elasticsearchserviceEncryptionAtRestOptions                 string
	_elasticsearchserviceEngineType                              string
	_elasticsearchserviceFilters                                 string
	_elasticsearchserviceInstanceCount                           string
	_elasticsearchserviceInstanceType                            string
	_elasticsearchserviceLogPublishingOptions                    string
	_elasticsearchserviceMaxResults                              string
	_elasticsearchserviceNextToken                               string
	_elasticsearchserviceNodeToNodeEncryptionOptions             string
	_elasticsearchservicePackageDescription                      string
	_elasticsearchservicePackageID                               string
	_elasticsearchservicePackageName                             string
	_elasticsearchservicePackageSource                           string
	_elasticsearchservicePackageType                             string
	_elasticsearchservicePerformCheckOnly                        string
	_elasticsearchserviceReservationName                         string
	_elasticsearchserviceReservedElasticsearchInstanceId         string
	_elasticsearchserviceReservedElasticsearchInstanceOfferingId string
	_elasticsearchserviceSnapshotOptions                         string
	_elasticsearchserviceSourceDomainInfo                        string
	_elasticsearchserviceTagKeys                                 []string
	_elasticsearchserviceTagList                                 string
	_elasticsearchserviceTargetVersion                           string
	_elasticsearchserviceVpcEndpointId                           string
	_elasticsearchserviceVpcEndpointIds                          []string
	_elasticsearchserviceVpcOptions                              string
)

// Allows the destination domain owner to accept an inbound cross-cluster search
// connection request.
func elasticsearchservice_AcceptInboundCrossClusterSearchConnection(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.AcceptInboundCrossClusterSearchConnectionInput{
		// CrossClusterSearchConnectionId: *string, // Required
	}

	if len(_elasticsearchserviceCrossClusterSearchConnectionId) > 0 {
		input.CrossClusterSearchConnectionId = aws.String(_elasticsearchserviceCrossClusterSearchConnectionId)
	}

	if resp, err := client.AcceptInboundCrossClusterSearchConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches tags to an existing Elasticsearch domain. Tags are a set of
// case-sensitive key value pairs. An Elasticsearch domain may have up to 10 tags.
// See [Tagging Amazon Elasticsearch Service Domains for more information.]
//
// [Tagging Amazon Elasticsearch Service Domains for more information.]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-managedomains.html#es-managedomains-awsresorcetagging
func elasticsearchservice_AddTags(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.AddTagsInput{
		// ARN: *string, // Required
		// TagList: []types.Tag, // Required
	}

	if len(_elasticsearchserviceARN) > 0 {
		input.ARN = aws.String(_elasticsearchserviceARN)
	}
	if len(_elasticsearchserviceTagList) > 0 {
		if err := assignInputField(input, "TagList", _elasticsearchserviceTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a package with an Amazon ES domain.
func elasticsearchservice_AssociatePackage(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.AssociatePackageInput{
		// DomainName: *string, // Required
		// PackageID: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchservicePackageID) > 0 {
		input.PackageID = aws.String(_elasticsearchservicePackageID)
	}

	if resp, err := client.AssociatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides access to an Amazon OpenSearch Service domain through the use of an
// interface VPC endpoint.
func elasticsearchservice_AuthorizeVpcEndpointAccess(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.AuthorizeVpcEndpointAccessInput{
		// Account: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceAccount) > 0 {
		input.Account = aws.String(_elasticsearchserviceAccount)
	}
	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.AuthorizeVpcEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a pending configuration change on an Amazon OpenSearch Service domain.
func elasticsearchservice_CancelDomainConfigChange(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.CancelDomainConfigChangeInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _elasticsearchserviceDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.CancelDomainConfigChange(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a scheduled service software update for an Amazon ES domain. You can
// only perform this operation before the AutomatedUpdateDate and when the
// UpdateStatus is in the PENDING_UPDATE state.
func elasticsearchservice_CancelElasticsearchServiceSoftwareUpdate(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.CancelElasticsearchServiceSoftwareUpdateInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.CancelElasticsearchServiceSoftwareUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new Elasticsearch domain. For more information, see [Creating Elasticsearch Domains] in the Amazon
// Elasticsearch Service Developer Guide.
//
// [Creating Elasticsearch Domains]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomains
func elasticsearchservice_CreateElasticsearchDomain(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.CreateElasticsearchDomainInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceAccessPolicies) > 0 {
		input.AccessPolicies = aws.String(_elasticsearchserviceAccessPolicies)
	}
	if len(_elasticsearchserviceAdvancedOptions) > 0 {
		if err := assignInputField(input, "AdvancedOptions", _elasticsearchserviceAdvancedOptions); err != nil {
			log.Errorf("invalid --advanced-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceAdvancedSecurityOptions) > 0 {
		if err := assignInputField(input, "AdvancedSecurityOptions", _elasticsearchserviceAdvancedSecurityOptions); err != nil {
			log.Errorf("invalid --advanced-security-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceAutoTuneOptions) > 0 {
		if err := assignInputField(input, "AutoTuneOptions", _elasticsearchserviceAutoTuneOptions); err != nil {
			log.Errorf("invalid --auto-tune-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceCognitoOptions) > 0 {
		if err := assignInputField(input, "CognitoOptions", _elasticsearchserviceCognitoOptions); err != nil {
			log.Errorf("invalid --cognito-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceDomainEndpointOptions) > 0 {
		if err := assignInputField(input, "DomainEndpointOptions", _elasticsearchserviceDomainEndpointOptions); err != nil {
			log.Errorf("invalid --domain-endpoint-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceEBSOptions) > 0 {
		if err := assignInputField(input, "EBSOptions", _elasticsearchserviceEBSOptions); err != nil {
			log.Errorf("invalid --ebs-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceElasticsearchClusterConfig) > 0 {
		if err := assignInputField(input, "ElasticsearchClusterConfig", _elasticsearchserviceElasticsearchClusterConfig); err != nil {
			log.Errorf("invalid --elasticsearch-cluster-config: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceElasticsearchVersion) > 0 {
		input.ElasticsearchVersion = aws.String(_elasticsearchserviceElasticsearchVersion)
	}
	if len(_elasticsearchserviceEncryptionAtRestOptions) > 0 {
		if err := assignInputField(input, "EncryptionAtRestOptions", _elasticsearchserviceEncryptionAtRestOptions); err != nil {
			log.Errorf("invalid --encryption-at-rest-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceLogPublishingOptions) > 0 {
		if err := assignInputField(input, "LogPublishingOptions", _elasticsearchserviceLogPublishingOptions); err != nil {
			log.Errorf("invalid --log-publishing-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNodeToNodeEncryptionOptions) > 0 {
		if err := assignInputField(input, "NodeToNodeEncryptionOptions", _elasticsearchserviceNodeToNodeEncryptionOptions); err != nil {
			log.Errorf("invalid --node-to-node-encryption-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceSnapshotOptions) > 0 {
		if err := assignInputField(input, "SnapshotOptions", _elasticsearchserviceSnapshotOptions); err != nil {
			log.Errorf("invalid --snapshot-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceTagList) > 0 {
		if err := assignInputField(input, "TagList", _elasticsearchserviceTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceVpcOptions) > 0 {
		if err := assignInputField(input, "VPCOptions", _elasticsearchserviceVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateElasticsearchDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new cross-cluster search connection from a source domain to a
// destination domain.
func elasticsearchservice_CreateOutboundCrossClusterSearchConnection(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.CreateOutboundCrossClusterSearchConnectionInput{
		// ConnectionAlias: *string, // Required
		// DestinationDomainInfo: *types.DomainInformation, // Required
		// SourceDomainInfo: *types.DomainInformation, // Required
	}

	if len(_elasticsearchserviceConnectionAlias) > 0 {
		input.ConnectionAlias = aws.String(_elasticsearchserviceConnectionAlias)
	}
	if len(_elasticsearchserviceDestinationDomainInfo) > 0 {
		if err := assignInputField(input, "DestinationDomainInfo", _elasticsearchserviceDestinationDomainInfo); err != nil {
			log.Errorf("invalid --destination-domain-info: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceSourceDomainInfo) > 0 {
		if err := assignInputField(input, "SourceDomainInfo", _elasticsearchserviceSourceDomainInfo); err != nil {
			log.Errorf("invalid --source-domain-info: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOutboundCrossClusterSearchConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Create a package for use with Amazon ES domains.
func elasticsearchservice_CreatePackage(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.CreatePackageInput{
		// PackageName: *string, // Required
		// PackageSource: *types.PackageSource, // Required
		// PackageType: types.PackageType, // Required
	}

	if len(_elasticsearchservicePackageName) > 0 {
		input.PackageName = aws.String(_elasticsearchservicePackageName)
	}
	if len(_elasticsearchservicePackageSource) > 0 {
		if err := assignInputField(input, "PackageSource", _elasticsearchservicePackageSource); err != nil {
			log.Errorf("invalid --package-source: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchservicePackageType) > 0 {
		if err := assignInputField(input, "PackageType", _elasticsearchservicePackageType); err != nil {
			log.Errorf("invalid --package-type: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchservicePackageDescription) > 0 {
		input.PackageDescription = aws.String(_elasticsearchservicePackageDescription)
	}

	if resp, err := client.CreatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon OpenSearch Service-managed VPC endpoint.
func elasticsearchservice_CreateVpcEndpoint(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.CreateVpcEndpointInput{
		// DomainArn: *string, // Required
		// VpcOptions: *types.VPCOptions, // Required
	}

	if len(_elasticsearchserviceDomainArn) > 0 {
		input.DomainArn = aws.String(_elasticsearchserviceDomainArn)
	}
	if len(_elasticsearchserviceVpcOptions) > 0 {
		if err := assignInputField(input, "VpcOptions", _elasticsearchserviceVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceClientToken) > 0 {
		input.ClientToken = aws.String(_elasticsearchserviceClientToken)
	}

	if resp, err := client.CreateVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Permanently deletes the specified Elasticsearch domain and all of its data.
// Once a domain is deleted, it cannot be recovered.
func elasticsearchservice_DeleteElasticsearchDomain(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DeleteElasticsearchDomainInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.DeleteElasticsearchDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the service-linked role that Elasticsearch Service uses to manage and
// maintain VPC domains. Role deletion will fail if any existing VPC domains use
// the role. You must delete any such Elasticsearch domains before deleting the
// role. See [Deleting Elasticsearch Service Role]in VPC Endpoints for Amazon Elasticsearch Service Domains.
//
// [Deleting Elasticsearch Service Role]: http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-enabling-slr
func elasticsearchservice_DeleteElasticsearchServiceRole(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DeleteElasticsearchServiceRoleInput{}

	if resp, err := client.DeleteElasticsearchServiceRole(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the destination domain owner to delete an existing inbound cross-cluster
// search connection.
func elasticsearchservice_DeleteInboundCrossClusterSearchConnection(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DeleteInboundCrossClusterSearchConnectionInput{
		// CrossClusterSearchConnectionId: *string, // Required
	}

	if len(_elasticsearchserviceCrossClusterSearchConnectionId) > 0 {
		input.CrossClusterSearchConnectionId = aws.String(_elasticsearchserviceCrossClusterSearchConnectionId)
	}

	if resp, err := client.DeleteInboundCrossClusterSearchConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the source domain owner to delete an existing outbound cross-cluster
// search connection.
func elasticsearchservice_DeleteOutboundCrossClusterSearchConnection(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DeleteOutboundCrossClusterSearchConnectionInput{
		// CrossClusterSearchConnectionId: *string, // Required
	}

	if len(_elasticsearchserviceCrossClusterSearchConnectionId) > 0 {
		input.CrossClusterSearchConnectionId = aws.String(_elasticsearchserviceCrossClusterSearchConnectionId)
	}

	if resp, err := client.DeleteOutboundCrossClusterSearchConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete the package.
func elasticsearchservice_DeletePackage(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DeletePackageInput{
		// PackageID: *string, // Required
	}

	if len(_elasticsearchservicePackageID) > 0 {
		input.PackageID = aws.String(_elasticsearchservicePackageID)
	}

	if resp, err := client.DeletePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon OpenSearch Service-managed interface VPC endpoint.
func elasticsearchservice_DeleteVpcEndpoint(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DeleteVpcEndpointInput{
		// VpcEndpointId: *string, // Required
	}

	if len(_elasticsearchserviceVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_elasticsearchserviceVpcEndpointId)
	}

	if resp, err := client.DeleteVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides scheduled Auto-Tune action details for the Elasticsearch domain, such
// as Auto-Tune action type, description, severity, and scheduled date.
func elasticsearchservice_DescribeDomainAutoTunes(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeDomainAutoTunesInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeDomainAutoTunes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.DescribeDomainAutoTunesOutput
	p := elasticsearchservice.NewDescribeDomainAutoTunesPaginator(client, input)
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

// Returns information about the current blue/green deployment happening on a
// domain, including a change ID, status, and progress stages.
func elasticsearchservice_DescribeDomainChangeProgress(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeDomainChangeProgressInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceChangeId) > 0 {
		input.ChangeId = aws.String(_elasticsearchserviceChangeId)
	}

	if resp, err := client.DescribeDomainChangeProgress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns domain configuration information about the specified Elasticsearch
// domain, including the domain ID, domain endpoint, and domain ARN.
func elasticsearchservice_DescribeElasticsearchDomain(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeElasticsearchDomainInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.DescribeElasticsearchDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides cluster configuration information about the specified Elasticsearch
// domain, such as the state, creation date, update version, and update date for
// cluster options.
func elasticsearchservice_DescribeElasticsearchDomainConfig(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeElasticsearchDomainConfigInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.DescribeElasticsearchDomainConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns domain configuration information about the specified Elasticsearch
// domains, including the domain ID, domain endpoint, and domain ARN.
func elasticsearchservice_DescribeElasticsearchDomains(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeElasticsearchDomainsInput{
		// DomainNames: []string, // Required
	}

	if len(_elasticsearchserviceDomainNames) > 0 {
		input.DomainNames = append([]string(nil), _elasticsearchserviceDomainNames...)
	}

	if resp, err := client.DescribeElasticsearchDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describe Elasticsearch Limits for a given InstanceType and
// ElasticsearchVersion. When modifying existing Domain, specify the DomainNameto know what
// Limits are supported for modifying.
func elasticsearchservice_DescribeElasticsearchInstanceTypeLimits(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeElasticsearchInstanceTypeLimitsInput{
		// ElasticsearchVersion: *string, // Required
		// InstanceType: types.ESPartitionInstanceType, // Required
	}

	if len(_elasticsearchserviceElasticsearchVersion) > 0 {
		input.ElasticsearchVersion = aws.String(_elasticsearchserviceElasticsearchVersion)
	}
	if len(_elasticsearchserviceInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _elasticsearchserviceInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.DescribeElasticsearchInstanceTypeLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the inbound cross-cluster search connections for a destination domain.
func elasticsearchservice_DescribeInboundCrossClusterSearchConnections(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsInput{}

	if len(_elasticsearchserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _elasticsearchserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInboundCrossClusterSearchConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.DescribeInboundCrossClusterSearchConnectionsOutput
	p := elasticsearchservice.NewDescribeInboundCrossClusterSearchConnectionsPaginator(client, input)
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

// Lists all the outbound cross-cluster search connections for a source domain.
func elasticsearchservice_DescribeOutboundCrossClusterSearchConnections(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsInput{}

	if len(_elasticsearchserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _elasticsearchserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOutboundCrossClusterSearchConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.DescribeOutboundCrossClusterSearchConnectionsOutput
	p := elasticsearchservice.NewDescribeOutboundCrossClusterSearchConnectionsPaginator(client, input)
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

// Describes all packages available to Amazon ES. Includes options for filtering,
// limiting the number of results, and pagination.
func elasticsearchservice_DescribePackages(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribePackagesInput{}

	if len(_elasticsearchserviceFilters) > 0 {
		if err := assignInputField(input, "Filters", _elasticsearchserviceFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribePackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.DescribePackagesOutput
	p := elasticsearchservice.NewDescribePackagesPaginator(client, input)
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

// Lists available reserved Elasticsearch instance offerings.
func elasticsearchservice_DescribeReservedElasticsearchInstanceOfferings(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsInput{}

	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}
	if len(_elasticsearchserviceReservedElasticsearchInstanceOfferingId) > 0 {
		input.ReservedElasticsearchInstanceOfferingId = aws.String(_elasticsearchserviceReservedElasticsearchInstanceOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedElasticsearchInstanceOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.DescribeReservedElasticsearchInstanceOfferingsOutput
	p := elasticsearchservice.NewDescribeReservedElasticsearchInstanceOfferingsPaginator(client, input)
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

// Returns information about reserved Elasticsearch instances for this account.
func elasticsearchservice_DescribeReservedElasticsearchInstances(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeReservedElasticsearchInstancesInput{}

	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}
	if len(_elasticsearchserviceReservedElasticsearchInstanceId) > 0 {
		input.ReservedElasticsearchInstanceId = aws.String(_elasticsearchserviceReservedElasticsearchInstanceId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedElasticsearchInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.DescribeReservedElasticsearchInstancesOutput
	p := elasticsearchservice.NewDescribeReservedElasticsearchInstancesPaginator(client, input)
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

// Describes one or more Amazon OpenSearch Service-managed VPC endpoints.
func elasticsearchservice_DescribeVpcEndpoints(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DescribeVpcEndpointsInput{
		// VpcEndpointIds: []string, // Required
	}

	if len(_elasticsearchserviceVpcEndpointIds) > 0 {
		input.VpcEndpointIds = append([]string(nil), _elasticsearchserviceVpcEndpointIds...)
	}

	if resp, err := client.DescribeVpcEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dissociates a package from the Amazon ES domain.
func elasticsearchservice_DissociatePackage(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.DissociatePackageInput{
		// DomainName: *string, // Required
		// PackageID: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchservicePackageID) > 0 {
		input.PackageID = aws.String(_elasticsearchservicePackageID)
	}

	if resp, err := client.DissociatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of upgrade compatible Elastisearch versions. You can optionally
// pass a DomainNameto get all upgrade compatible Elasticsearch versions for that specific
// domain.
func elasticsearchservice_GetCompatibleElasticsearchVersions(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.GetCompatibleElasticsearchVersionsInput{}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.GetCompatibleElasticsearchVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of versions of the package, along with their creation time and
// commit message.
func elasticsearchservice_GetPackageVersionHistory(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.GetPackageVersionHistoryInput{
		// PackageID: *string, // Required
	}

	if len(_elasticsearchservicePackageID) > 0 {
		input.PackageID = aws.String(_elasticsearchservicePackageID)
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetPackageVersionHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.GetPackageVersionHistoryOutput
	p := elasticsearchservice.NewGetPackageVersionHistoryPaginator(client, input)
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

// Retrieves the complete history of the last 10 upgrades that were performed on
// the domain.
func elasticsearchservice_GetUpgradeHistory(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.GetUpgradeHistoryInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.GetUpgradeHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.GetUpgradeHistoryOutput
	p := elasticsearchservice.NewGetUpgradeHistoryPaginator(client, input)
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

// Retrieves the latest status of the last upgrade or upgrade eligibility check
// that was performed on the domain.
func elasticsearchservice_GetUpgradeStatus(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.GetUpgradeStatusInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.GetUpgradeStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the name of all Elasticsearch domains owned by the current user's
// account.
func elasticsearchservice_ListDomainNames(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListDomainNamesInput{}

	if len(_elasticsearchserviceEngineType) > 0 {
		if err := assignInputField(input, "EngineType", _elasticsearchserviceEngineType); err != nil {
			log.Errorf("invalid --engine-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListDomainNames(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all Amazon ES domains associated with the package.
func elasticsearchservice_ListDomainsForPackage(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListDomainsForPackageInput{
		// PackageID: *string, // Required
	}

	if len(_elasticsearchservicePackageID) > 0 {
		input.PackageID = aws.String(_elasticsearchservicePackageID)
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDomainsForPackage(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.ListDomainsForPackageOutput
	p := elasticsearchservice.NewListDomainsForPackagePaginator(client, input)
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

// List all Elasticsearch instance types that are supported for given
// ElasticsearchVersion
func elasticsearchservice_ListElasticsearchInstanceTypes(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListElasticsearchInstanceTypesInput{
		// ElasticsearchVersion: *string, // Required
	}

	if len(_elasticsearchserviceElasticsearchVersion) > 0 {
		input.ElasticsearchVersion = aws.String(_elasticsearchserviceElasticsearchVersion)
	}
	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListElasticsearchInstanceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.ListElasticsearchInstanceTypesOutput
	p := elasticsearchservice.NewListElasticsearchInstanceTypesPaginator(client, input)
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

// List all supported Elasticsearch versions
func elasticsearchservice_ListElasticsearchVersions(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListElasticsearchVersionsInput{}

	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListElasticsearchVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.ListElasticsearchVersionsOutput
	p := elasticsearchservice.NewListElasticsearchVersionsPaginator(client, input)
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

// Lists all packages associated with the Amazon ES domain.
func elasticsearchservice_ListPackagesForDomain(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListPackagesForDomainInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _elasticsearchserviceMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPackagesForDomain(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*elasticsearchservice.ListPackagesForDomainOutput
	p := elasticsearchservice.NewListPackagesForDomainPaginator(client, input)
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

// Returns all tags for the given Elasticsearch domain.
func elasticsearchservice_ListTags(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListTagsInput{
		// ARN: *string, // Required
	}

	if len(_elasticsearchserviceARN) > 0 {
		input.ARN = aws.String(_elasticsearchserviceARN)
	}

	if resp, err := client.ListTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about each principal that is allowed to access a given
// Amazon OpenSearch Service domain through the use of an interface VPC endpoint.
func elasticsearchservice_ListVpcEndpointAccess(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListVpcEndpointAccessInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if resp, err := client.ListVpcEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all Amazon OpenSearch Service-managed VPC endpoints in the current
// account and Region.
func elasticsearchservice_ListVpcEndpoints(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListVpcEndpointsInput{}

	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if resp, err := client.ListVpcEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all Amazon OpenSearch Service-managed VPC endpoints associated with a
// particular domain.
func elasticsearchservice_ListVpcEndpointsForDomain(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.ListVpcEndpointsForDomainInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceNextToken) > 0 {
		input.NextToken = aws.String(_elasticsearchserviceNextToken)
	}

	if resp, err := client.ListVpcEndpointsForDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to purchase reserved Elasticsearch instances.
func elasticsearchservice_PurchaseReservedElasticsearchInstanceOffering(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.PurchaseReservedElasticsearchInstanceOfferingInput{
		// ReservationName: *string, // Required
		// ReservedElasticsearchInstanceOfferingId: *string, // Required
	}

	if len(_elasticsearchserviceReservationName) > 0 {
		input.ReservationName = aws.String(_elasticsearchserviceReservationName)
	}
	if len(_elasticsearchserviceReservedElasticsearchInstanceOfferingId) > 0 {
		input.ReservedElasticsearchInstanceOfferingId = aws.String(_elasticsearchserviceReservedElasticsearchInstanceOfferingId)
	}
	if len(_elasticsearchserviceInstanceCount) > 0 {
		if err := assignInputField(input, "InstanceCount", _elasticsearchserviceInstanceCount); err != nil {
			log.Errorf("invalid --instance-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseReservedElasticsearchInstanceOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the destination domain owner to reject an inbound cross-cluster search
// connection request.
func elasticsearchservice_RejectInboundCrossClusterSearchConnection(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.RejectInboundCrossClusterSearchConnectionInput{
		// CrossClusterSearchConnectionId: *string, // Required
	}

	if len(_elasticsearchserviceCrossClusterSearchConnectionId) > 0 {
		input.CrossClusterSearchConnectionId = aws.String(_elasticsearchserviceCrossClusterSearchConnectionId)
	}

	if resp, err := client.RejectInboundCrossClusterSearchConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified set of tags from the specified Elasticsearch domain.
func elasticsearchservice_RemoveTags(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.RemoveTagsInput{
		// ARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_elasticsearchserviceARN) > 0 {
		input.ARN = aws.String(_elasticsearchserviceARN)
	}
	if len(_elasticsearchserviceTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _elasticsearchserviceTagKeys...)
	}

	if resp, err := client.RemoveTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Revokes access to an Amazon OpenSearch Service domain that was provided through
// an interface VPC endpoint.
func elasticsearchservice_RevokeVpcEndpointAccess(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.RevokeVpcEndpointAccessInput{
		// Account: *string, // Required
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceAccount) > 0 {
		input.Account = aws.String(_elasticsearchserviceAccount)
	}
	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.RevokeVpcEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Schedules a service software update for an Amazon ES domain.
func elasticsearchservice_StartElasticsearchServiceSoftwareUpdate(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.StartElasticsearchServiceSoftwareUpdateInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}

	if resp, err := client.StartElasticsearchServiceSoftwareUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the cluster configuration of the specified Elasticsearch domain,
// setting as setting the instance type and the number of instances.
func elasticsearchservice_UpdateElasticsearchDomainConfig(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.UpdateElasticsearchDomainConfigInput{
		// DomainName: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceAccessPolicies) > 0 {
		input.AccessPolicies = aws.String(_elasticsearchserviceAccessPolicies)
	}
	if len(_elasticsearchserviceAdvancedOptions) > 0 {
		if err := assignInputField(input, "AdvancedOptions", _elasticsearchserviceAdvancedOptions); err != nil {
			log.Errorf("invalid --advanced-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceAdvancedSecurityOptions) > 0 {
		if err := assignInputField(input, "AdvancedSecurityOptions", _elasticsearchserviceAdvancedSecurityOptions); err != nil {
			log.Errorf("invalid --advanced-security-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceAutoTuneOptions) > 0 {
		if err := assignInputField(input, "AutoTuneOptions", _elasticsearchserviceAutoTuneOptions); err != nil {
			log.Errorf("invalid --auto-tune-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceCognitoOptions) > 0 {
		if err := assignInputField(input, "CognitoOptions", _elasticsearchserviceCognitoOptions); err != nil {
			log.Errorf("invalid --cognito-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceDomainEndpointOptions) > 0 {
		if err := assignInputField(input, "DomainEndpointOptions", _elasticsearchserviceDomainEndpointOptions); err != nil {
			log.Errorf("invalid --domain-endpoint-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _elasticsearchserviceDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceEBSOptions) > 0 {
		if err := assignInputField(input, "EBSOptions", _elasticsearchserviceEBSOptions); err != nil {
			log.Errorf("invalid --ebs-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceElasticsearchClusterConfig) > 0 {
		if err := assignInputField(input, "ElasticsearchClusterConfig", _elasticsearchserviceElasticsearchClusterConfig); err != nil {
			log.Errorf("invalid --elasticsearch-cluster-config: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceEncryptionAtRestOptions) > 0 {
		if err := assignInputField(input, "EncryptionAtRestOptions", _elasticsearchserviceEncryptionAtRestOptions); err != nil {
			log.Errorf("invalid --encryption-at-rest-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceLogPublishingOptions) > 0 {
		if err := assignInputField(input, "LogPublishingOptions", _elasticsearchserviceLogPublishingOptions); err != nil {
			log.Errorf("invalid --log-publishing-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceNodeToNodeEncryptionOptions) > 0 {
		if err := assignInputField(input, "NodeToNodeEncryptionOptions", _elasticsearchserviceNodeToNodeEncryptionOptions); err != nil {
			log.Errorf("invalid --node-to-node-encryption-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceSnapshotOptions) > 0 {
		if err := assignInputField(input, "SnapshotOptions", _elasticsearchserviceSnapshotOptions); err != nil {
			log.Errorf("invalid --snapshot-options: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceVpcOptions) > 0 {
		if err := assignInputField(input, "VPCOptions", _elasticsearchserviceVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateElasticsearchDomainConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a package for use with Amazon ES domains.
func elasticsearchservice_UpdatePackage(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.UpdatePackageInput{
		// PackageID: *string, // Required
		// PackageSource: *types.PackageSource, // Required
	}

	if len(_elasticsearchservicePackageID) > 0 {
		input.PackageID = aws.String(_elasticsearchservicePackageID)
	}
	if len(_elasticsearchservicePackageSource) > 0 {
		if err := assignInputField(input, "PackageSource", _elasticsearchservicePackageSource); err != nil {
			log.Errorf("invalid --package-source: %s", err.Error())
			return
		}
	}
	if len(_elasticsearchserviceCommitMessage) > 0 {
		input.CommitMessage = aws.String(_elasticsearchserviceCommitMessage)
	}
	if len(_elasticsearchservicePackageDescription) > 0 {
		input.PackageDescription = aws.String(_elasticsearchservicePackageDescription)
	}

	if resp, err := client.UpdatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies an Amazon OpenSearch Service-managed interface VPC endpoint.
func elasticsearchservice_UpdateVpcEndpoint(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.UpdateVpcEndpointInput{
		// VpcEndpointId: *string, // Required
		// VpcOptions: *types.VPCOptions, // Required
	}

	if len(_elasticsearchserviceVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_elasticsearchserviceVpcEndpointId)
	}
	if len(_elasticsearchserviceVpcOptions) > 0 {
		if err := assignInputField(input, "VpcOptions", _elasticsearchserviceVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to either upgrade your domain or perform an Upgrade eligibility
// check to a compatible Elasticsearch version.
func elasticsearchservice_UpgradeElasticsearchDomain(cfg aws.Config, client *elasticsearchservice.Client) {
	input := &elasticsearchservice.UpgradeElasticsearchDomainInput{
		// DomainName: *string, // Required
		// TargetVersion: *string, // Required
	}

	if len(_elasticsearchserviceDomainName) > 0 {
		input.DomainName = aws.String(_elasticsearchserviceDomainName)
	}
	if len(_elasticsearchserviceTargetVersion) > 0 {
		input.TargetVersion = aws.String(_elasticsearchserviceTargetVersion)
	}
	if len(_elasticsearchservicePerformCheckOnly) > 0 {
		if err := assignInputField(input, "PerformCheckOnly", _elasticsearchservicePerformCheckOnly); err != nil {
			log.Errorf("invalid --perform-check-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpgradeElasticsearchDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_elasticsearchserviceCmd)
	_elasticsearchserviceCmd.Flags().SortFlags = false

	_elasticsearchserviceCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_elasticsearchserviceCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_elasticsearchserviceCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceAccessPolicies, "access-policies", "", "", "Access Policies")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceAccount, "account", "", "", "Account")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceAdvancedOptions, "advanced-options", "", "", "Advanced Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceAdvancedSecurityOptions, "advanced-security-options", "", "", "Advanced Security Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceARN, "arn", "", "", "ARN")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceAutoTuneOptions, "auto-tune-options", "", "", "Auto Tune Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceChangeId, "change-id", "", "", "Change ID")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceClientToken, "client-token", "", "", "Client Token")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceCognitoOptions, "cognito-options", "", "", "Cognito Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceCommitMessage, "commit-message", "", "", "Commit Message")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceConnectionAlias, "connection-alias", "", "", "Connection Alias")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceCrossClusterSearchConnectionId, "cross-cluster-search-connection-id", "", "", "Cross Cluster Search Connection ID")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceDestinationDomainInfo, "destination-domain-info", "", "", "Destination Domain Info")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceDomainArn, "domain-arn", "", "", "Domain ARN")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceDomainEndpointOptions, "domain-endpoint-options", "", "", "Domain Endpoint Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceDomainName, "domain-name", "", "", "Domain Name")
	_elasticsearchserviceCmd.Flags().StringSliceVarP(&_elasticsearchserviceDomainNames, "domain-names", "", nil, "Domain Names")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceDryRun, "dry-run", "", "", "Dry Run")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceEBSOptions, "ebs-options", "", "", "Ebs Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceElasticsearchClusterConfig, "elasticsearch-cluster-config", "", "", "Elasticsearch Cluster Config")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceElasticsearchVersion, "elasticsearch-version", "", "", "Elasticsearch Version")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceEncryptionAtRestOptions, "encryption-at-rest-options", "", "", "Encryption At Rest Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceEngineType, "engine-type", "", "", "Engine Type")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceFilters, "filters", "", "", "Filters")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceInstanceCount, "instance-count", "", "", "Instance Count")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceInstanceType, "instance-type", "", "", "Instance Type")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceLogPublishingOptions, "log-publishing-options", "", "", "Log Publishing Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceMaxResults, "max-results", "", "", "Max Results")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceNextToken, "next-token", "", "", "Next Token")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceNodeToNodeEncryptionOptions, "node-to-node-encryption-options", "", "", "Node To Node Encryption Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchservicePackageDescription, "package-description", "", "", "Package Description")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchservicePackageID, "package-id", "", "", "Package ID")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchservicePackageName, "package-name", "", "", "Package Name")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchservicePackageSource, "package-source", "", "", "Package Source")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchservicePackageType, "package-type", "", "", "Package Type")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchservicePerformCheckOnly, "perform-check-only", "", "", "Perform Check Only")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceReservationName, "reservation-name", "", "", "Reservation Name")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceReservedElasticsearchInstanceId, "reserved-elasticsearch-instance-id", "", "", "Reserved Elasticsearch Instance ID")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceReservedElasticsearchInstanceOfferingId, "reserved-elasticsearch-instance-offering-id", "", "", "Reserved Elasticsearch Instance Offering ID")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceSnapshotOptions, "snapshot-options", "", "", "Snapshot Options")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceSourceDomainInfo, "source-domain-info", "", "", "Source Domain Info")
	_elasticsearchserviceCmd.Flags().StringSliceVarP(&_elasticsearchserviceTagKeys, "tag-keys", "", nil, "Tag Keys")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceTagList, "tag-list", "", "", "Tag List")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceTargetVersion, "target-version", "", "", "Target Version")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceVpcEndpointId, "vpc-endpoint-id", "", "", "VPC Endpoint ID")
	_elasticsearchserviceCmd.Flags().StringSliceVarP(&_elasticsearchserviceVpcEndpointIds, "vpc-endpoint-ids", "", nil, "VPC Endpoint Ids")
	_elasticsearchserviceCmd.Flags().StringVarP(&_elasticsearchserviceVpcOptions, "vpc-options", "", "", "VPC Options")

	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceAcceptInboundCrossClusterSearchConnection, "accept-inbound-cross-cluster-search-connection", "", false, "Accept Inbound Cross Cluster Search Connection")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceAddTags, "add-tags", "", false, "Add Tags")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceAssociatePackage, "associate-package", "", false, "Associate Package")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceAuthorizeVpcEndpointAccess, "authorize-vpc-endpoint-access", "", false, "Authorize VPC Endpoint Access")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceCancelDomainConfigChange, "cancel-domain-config-change", "", false, "Cancel Domain Config Change")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceCancelElasticsearchServiceSoftwareUpdate, "cancel-elasticsearch-service-software-update", "", false, "Cancel Elasticsearch Service Software Update")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceCreateElasticsearchDomain, "create-elasticsearch-domain", "", false, "Create Elasticsearch Domain")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceCreateOutboundCrossClusterSearchConnection, "create-outbound-cross-cluster-search-connection", "", false, "Create Outbound Cross Cluster Search Connection")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceCreatePackage, "create-package", "", false, "Create Package")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceCreateVpcEndpoint, "create-vpc-endpoint", "", false, "Create VPC Endpoint")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDeleteElasticsearchDomain, "delete-elasticsearch-domain", "", false, "Delete Elasticsearch Domain")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDeleteElasticsearchServiceRole, "delete-elasticsearch-service-role", "", false, "Delete Elasticsearch Service Role")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDeleteInboundCrossClusterSearchConnection, "delete-inbound-cross-cluster-search-connection", "", false, "Delete Inbound Cross Cluster Search Connection")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDeleteOutboundCrossClusterSearchConnection, "delete-outbound-cross-cluster-search-connection", "", false, "Delete Outbound Cross Cluster Search Connection")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDeletePackage, "delete-package", "", false, "Delete Package")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDeleteVpcEndpoint, "delete-vpc-endpoint", "", false, "Delete VPC Endpoint")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeDomainAutoTunes, "describe-domain-auto-tunes", "", false, "Describe Domain Auto Tunes")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeDomainChangeProgress, "describe-domain-change-progress", "", false, "Describe Domain Change Progress")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeElasticsearchDomain, "describe-elasticsearch-domain", "", false, "Describe Elasticsearch Domain")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeElasticsearchDomainConfig, "describe-elasticsearch-domain-config", "", false, "Describe Elasticsearch Domain Config")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeElasticsearchDomains, "describe-elasticsearch-domains", "", false, "Describe Elasticsearch Domains")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeElasticsearchInstanceTypeLimits, "describe-elasticsearch-instance-type-limits", "", false, "Describe Elasticsearch Instance Type Limits")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeInboundCrossClusterSearchConnections, "describe-inbound-cross-cluster-search-connections", "", false, "Describe Inbound Cross Cluster Search Connections")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeOutboundCrossClusterSearchConnections, "describe-outbound-cross-cluster-search-connections", "", false, "Describe Outbound Cross Cluster Search Connections")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribePackages, "describe-packages", "", false, "Describe Packages")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeReservedElasticsearchInstanceOfferings, "describe-reserved-elasticsearch-instance-offerings", "", false, "Describe Reserved Elasticsearch Instance Offerings")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeReservedElasticsearchInstances, "describe-reserved-elasticsearch-instances", "", false, "Describe Reserved Elasticsearch Instances")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDescribeVpcEndpoints, "describe-vpc-endpoints", "", false, "Describe VPC Endpoints")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceDissociatePackage, "dissociate-package", "", false, "Dissociate Package")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceGetCompatibleElasticsearchVersions, "get-compatible-elasticsearch-versions", "", false, "Get Compatible Elasticsearch Versions")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceGetPackageVersionHistory, "get-package-version-history", "", false, "Get Package Version History")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceGetUpgradeHistory, "get-upgrade-history", "", false, "Get Upgrade History")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceGetUpgradeStatus, "get-upgrade-status", "", false, "Get Upgrade Status")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListDomainNames, "list-domain-names", "", false, "List Domain Names")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListDomainsForPackage, "list-domains-for-package", "", false, "List Domains For Package")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListElasticsearchInstanceTypes, "list-elasticsearch-instance-types", "", false, "List Elasticsearch Instance Types")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListElasticsearchVersions, "list-elasticsearch-versions", "", false, "List Elasticsearch Versions")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListPackagesForDomain, "list-packages-for-domain", "", false, "List Packages For Domain")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListTags, "list-tags", "", false, "List Tags")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListVpcEndpointAccess, "list-vpc-endpoint-access", "", false, "List VPC Endpoint Access")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListVpcEndpoints, "list-vpc-endpoints", "", false, "List VPC Endpoints")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceListVpcEndpointsForDomain, "list-vpc-endpoints-for-domain", "", false, "List VPC Endpoints For Domain")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchservicePurchaseReservedElasticsearchInstanceOffering, "purchase-reserved-elasticsearch-instance-offering", "", false, "Purchase Reserved Elasticsearch Instance Offering")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceRejectInboundCrossClusterSearchConnection, "reject-inbound-cross-cluster-search-connection", "", false, "Reject Inbound Cross Cluster Search Connection")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceRemoveTags, "remove-tags", "", false, "Remove Tags")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceRevokeVpcEndpointAccess, "revoke-vpc-endpoint-access", "", false, "Revoke VPC Endpoint Access")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceStartElasticsearchServiceSoftwareUpdate, "start-elasticsearch-service-software-update", "", false, "Start Elasticsearch Service Software Update")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceUpdateElasticsearchDomainConfig, "update-elasticsearch-domain-config", "", false, "Update Elasticsearch Domain Config")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceUpdatePackage, "update-package", "", false, "Update Package")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceUpdateVpcEndpoint, "update-vpc-endpoint", "", false, "Update VPC Endpoint")
	_elasticsearchserviceCmd.Flags().BoolVarP(&_elasticsearchserviceUpgradeElasticsearchDomain, "upgrade-elasticsearch-domain", "", false, "Upgrade Elasticsearch Domain")

}
