package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/opensearch"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// opensearchCmd represents the opensearch command
var _opensearchCmd = &cobra.Command{
	Use:   "opensearch",
	Short: "AWS opensearch CLI",
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
		client := opensearch.NewFromConfig(cfg)
		if _opensearchAcceptInboundConnection {
			opensearch_AcceptInboundConnection(cfg, client)
			return
		}
		if _opensearchAddDataSource {
			opensearch_AddDataSource(cfg, client)
			return
		}
		if _opensearchAddDirectQueryDataSource {
			opensearch_AddDirectQueryDataSource(cfg, client)
			return
		}
		if _opensearchAddTags {
			opensearch_AddTags(cfg, client)
			return
		}
		if _opensearchAssociatePackage {
			opensearch_AssociatePackage(cfg, client)
			return
		}
		if _opensearchAssociatePackages {
			opensearch_AssociatePackages(cfg, client)
			return
		}
		if _opensearchAuthorizeVpcEndpointAccess {
			opensearch_AuthorizeVpcEndpointAccess(cfg, client)
			return
		}
		if _opensearchCancelDomainConfigChange {
			opensearch_CancelDomainConfigChange(cfg, client)
			return
		}
		if _opensearchCancelServiceSoftwareUpdate {
			opensearch_CancelServiceSoftwareUpdate(cfg, client)
			return
		}
		if _opensearchCreateApplication {
			opensearch_CreateApplication(cfg, client)
			return
		}
		if _opensearchCreateDomain {
			opensearch_CreateDomain(cfg, client)
			return
		}
		if _opensearchCreateIndex {
			opensearch_CreateIndex(cfg, client)
			return
		}
		if _opensearchCreateOutboundConnection {
			opensearch_CreateOutboundConnection(cfg, client)
			return
		}
		if _opensearchCreatePackage {
			opensearch_CreatePackage(cfg, client)
			return
		}
		if _opensearchCreateVpcEndpoint {
			opensearch_CreateVpcEndpoint(cfg, client)
			return
		}
		if _opensearchDeleteApplication {
			opensearch_DeleteApplication(cfg, client)
			return
		}
		if _opensearchDeleteDataSource {
			opensearch_DeleteDataSource(cfg, client)
			return
		}
		if _opensearchDeleteDirectQueryDataSource {
			opensearch_DeleteDirectQueryDataSource(cfg, client)
			return
		}
		if _opensearchDeleteDomain {
			opensearch_DeleteDomain(cfg, client)
			return
		}
		if _opensearchDeleteInboundConnection {
			opensearch_DeleteInboundConnection(cfg, client)
			return
		}
		if _opensearchDeleteIndex {
			opensearch_DeleteIndex(cfg, client)
			return
		}
		if _opensearchDeleteOutboundConnection {
			opensearch_DeleteOutboundConnection(cfg, client)
			return
		}
		if _opensearchDeletePackage {
			opensearch_DeletePackage(cfg, client)
			return
		}
		if _opensearchDeleteVpcEndpoint {
			opensearch_DeleteVpcEndpoint(cfg, client)
			return
		}
		if _opensearchDescribeDomain {
			opensearch_DescribeDomain(cfg, client)
			return
		}
		if _opensearchDescribeDomainAutoTunes {
			opensearch_DescribeDomainAutoTunes(cfg, client)
			return
		}
		if _opensearchDescribeDomainChangeProgress {
			opensearch_DescribeDomainChangeProgress(cfg, client)
			return
		}
		if _opensearchDescribeDomainConfig {
			opensearch_DescribeDomainConfig(cfg, client)
			return
		}
		if _opensearchDescribeDomainHealth {
			opensearch_DescribeDomainHealth(cfg, client)
			return
		}
		if _opensearchDescribeDomainNodes {
			opensearch_DescribeDomainNodes(cfg, client)
			return
		}
		if _opensearchDescribeDomains {
			opensearch_DescribeDomains(cfg, client)
			return
		}
		if _opensearchDescribeDryRunProgress {
			opensearch_DescribeDryRunProgress(cfg, client)
			return
		}
		if _opensearchDescribeInboundConnections {
			opensearch_DescribeInboundConnections(cfg, client)
			return
		}
		if _opensearchDescribeInstanceTypeLimits {
			opensearch_DescribeInstanceTypeLimits(cfg, client)
			return
		}
		if _opensearchDescribeOutboundConnections {
			opensearch_DescribeOutboundConnections(cfg, client)
			return
		}
		if _opensearchDescribePackages {
			opensearch_DescribePackages(cfg, client)
			return
		}
		if _opensearchDescribeReservedInstanceOfferings {
			opensearch_DescribeReservedInstanceOfferings(cfg, client)
			return
		}
		if _opensearchDescribeReservedInstances {
			opensearch_DescribeReservedInstances(cfg, client)
			return
		}
		if _opensearchDescribeVpcEndpoints {
			opensearch_DescribeVpcEndpoints(cfg, client)
			return
		}
		if _opensearchDissociatePackage {
			opensearch_DissociatePackage(cfg, client)
			return
		}
		if _opensearchDissociatePackages {
			opensearch_DissociatePackages(cfg, client)
			return
		}
		if _opensearchGetApplication {
			opensearch_GetApplication(cfg, client)
			return
		}
		if _opensearchGetCompatibleVersions {
			opensearch_GetCompatibleVersions(cfg, client)
			return
		}
		if _opensearchGetDataSource {
			opensearch_GetDataSource(cfg, client)
			return
		}
		if _opensearchGetDefaultApplicationSetting {
			opensearch_GetDefaultApplicationSetting(cfg, client)
			return
		}
		if _opensearchGetDirectQueryDataSource {
			opensearch_GetDirectQueryDataSource(cfg, client)
			return
		}
		if _opensearchGetDomainMaintenanceStatus {
			opensearch_GetDomainMaintenanceStatus(cfg, client)
			return
		}
		if _opensearchGetIndex {
			opensearch_GetIndex(cfg, client)
			return
		}
		if _opensearchGetPackageVersionHistory {
			opensearch_GetPackageVersionHistory(cfg, client)
			return
		}
		if _opensearchGetUpgradeHistory {
			opensearch_GetUpgradeHistory(cfg, client)
			return
		}
		if _opensearchGetUpgradeStatus {
			opensearch_GetUpgradeStatus(cfg, client)
			return
		}
		if _opensearchListApplications {
			opensearch_ListApplications(cfg, client)
			return
		}
		if _opensearchListDataSources {
			opensearch_ListDataSources(cfg, client)
			return
		}
		if _opensearchListDirectQueryDataSources {
			opensearch_ListDirectQueryDataSources(cfg, client)
			return
		}
		if _opensearchListDomainMaintenances {
			opensearch_ListDomainMaintenances(cfg, client)
			return
		}
		if _opensearchListDomainNames {
			opensearch_ListDomainNames(cfg, client)
			return
		}
		if _opensearchListDomainsForPackage {
			opensearch_ListDomainsForPackage(cfg, client)
			return
		}
		if _opensearchListInstanceTypeDetails {
			opensearch_ListInstanceTypeDetails(cfg, client)
			return
		}
		if _opensearchListPackagesForDomain {
			opensearch_ListPackagesForDomain(cfg, client)
			return
		}
		if _opensearchListScheduledActions {
			opensearch_ListScheduledActions(cfg, client)
			return
		}
		if _opensearchListTags {
			opensearch_ListTags(cfg, client)
			return
		}
		if _opensearchListVersions {
			opensearch_ListVersions(cfg, client)
			return
		}
		if _opensearchListVpcEndpointAccess {
			opensearch_ListVpcEndpointAccess(cfg, client)
			return
		}
		if _opensearchListVpcEndpoints {
			opensearch_ListVpcEndpoints(cfg, client)
			return
		}
		if _opensearchListVpcEndpointsForDomain {
			opensearch_ListVpcEndpointsForDomain(cfg, client)
			return
		}
		if _opensearchPurchaseReservedInstanceOffering {
			opensearch_PurchaseReservedInstanceOffering(cfg, client)
			return
		}
		if _opensearchPutDefaultApplicationSetting {
			opensearch_PutDefaultApplicationSetting(cfg, client)
			return
		}
		if _opensearchRejectInboundConnection {
			opensearch_RejectInboundConnection(cfg, client)
			return
		}
		if _opensearchRemoveTags {
			opensearch_RemoveTags(cfg, client)
			return
		}
		if _opensearchRevokeVpcEndpointAccess {
			opensearch_RevokeVpcEndpointAccess(cfg, client)
			return
		}
		if _opensearchStartDomainMaintenance {
			opensearch_StartDomainMaintenance(cfg, client)
			return
		}
		if _opensearchStartServiceSoftwareUpdate {
			opensearch_StartServiceSoftwareUpdate(cfg, client)
			return
		}
		if _opensearchUpdateApplication {
			opensearch_UpdateApplication(cfg, client)
			return
		}
		if _opensearchUpdateDataSource {
			opensearch_UpdateDataSource(cfg, client)
			return
		}
		if _opensearchUpdateDirectQueryDataSource {
			opensearch_UpdateDirectQueryDataSource(cfg, client)
			return
		}
		if _opensearchUpdateDomainConfig {
			opensearch_UpdateDomainConfig(cfg, client)
			return
		}
		if _opensearchUpdateIndex {
			opensearch_UpdateIndex(cfg, client)
			return
		}
		if _opensearchUpdatePackage {
			opensearch_UpdatePackage(cfg, client)
			return
		}
		if _opensearchUpdatePackageScope {
			opensearch_UpdatePackageScope(cfg, client)
			return
		}
		if _opensearchUpdateScheduledAction {
			opensearch_UpdateScheduledAction(cfg, client)
			return
		}
		if _opensearchUpdateVpcEndpoint {
			opensearch_UpdateVpcEndpoint(cfg, client)
			return
		}
		if _opensearchUpgradeDomain {
			opensearch_UpgradeDomain(cfg, client)
			return
		}

	},
}

var (
	_opensearchAcceptInboundConnection           bool
	_opensearchAddDataSource                     bool
	_opensearchAddDirectQueryDataSource          bool
	_opensearchAddTags                           bool
	_opensearchAssociatePackage                  bool
	_opensearchAssociatePackages                 bool
	_opensearchAuthorizeVpcEndpointAccess        bool
	_opensearchCancelDomainConfigChange          bool
	_opensearchCancelServiceSoftwareUpdate       bool
	_opensearchCreateApplication                 bool
	_opensearchCreateDomain                      bool
	_opensearchCreateIndex                       bool
	_opensearchCreateOutboundConnection          bool
	_opensearchCreatePackage                     bool
	_opensearchCreateVpcEndpoint                 bool
	_opensearchDeleteApplication                 bool
	_opensearchDeleteDataSource                  bool
	_opensearchDeleteDirectQueryDataSource       bool
	_opensearchDeleteDomain                      bool
	_opensearchDeleteInboundConnection           bool
	_opensearchDeleteIndex                       bool
	_opensearchDeleteOutboundConnection          bool
	_opensearchDeletePackage                     bool
	_opensearchDeleteVpcEndpoint                 bool
	_opensearchDescribeDomain                    bool
	_opensearchDescribeDomainAutoTunes           bool
	_opensearchDescribeDomainChangeProgress      bool
	_opensearchDescribeDomainConfig              bool
	_opensearchDescribeDomainHealth              bool
	_opensearchDescribeDomainNodes               bool
	_opensearchDescribeDomains                   bool
	_opensearchDescribeDryRunProgress            bool
	_opensearchDescribeInboundConnections        bool
	_opensearchDescribeInstanceTypeLimits        bool
	_opensearchDescribeOutboundConnections       bool
	_opensearchDescribePackages                  bool
	_opensearchDescribeReservedInstanceOfferings bool
	_opensearchDescribeReservedInstances         bool
	_opensearchDescribeVpcEndpoints              bool
	_opensearchDissociatePackage                 bool
	_opensearchDissociatePackages                bool
	_opensearchGetApplication                    bool
	_opensearchGetCompatibleVersions             bool
	_opensearchGetDataSource                     bool
	_opensearchGetDefaultApplicationSetting      bool
	_opensearchGetDirectQueryDataSource          bool
	_opensearchGetDomainMaintenanceStatus        bool
	_opensearchGetIndex                          bool
	_opensearchGetPackageVersionHistory          bool
	_opensearchGetUpgradeHistory                 bool
	_opensearchGetUpgradeStatus                  bool
	_opensearchListApplications                  bool
	_opensearchListDataSources                   bool
	_opensearchListDirectQueryDataSources        bool
	_opensearchListDomainMaintenances            bool
	_opensearchListDomainNames                   bool
	_opensearchListDomainsForPackage             bool
	_opensearchListInstanceTypeDetails           bool
	_opensearchListPackagesForDomain             bool
	_opensearchListScheduledActions              bool
	_opensearchListTags                          bool
	_opensearchListVersions                      bool
	_opensearchListVpcEndpointAccess             bool
	_opensearchListVpcEndpoints                  bool
	_opensearchListVpcEndpointsForDomain         bool
	_opensearchPurchaseReservedInstanceOffering  bool
	_opensearchPutDefaultApplicationSetting      bool
	_opensearchRejectInboundConnection           bool
	_opensearchRemoveTags                        bool
	_opensearchRevokeVpcEndpointAccess           bool
	_opensearchStartDomainMaintenance            bool
	_opensearchStartServiceSoftwareUpdate        bool
	_opensearchUpdateApplication                 bool
	_opensearchUpdateDataSource                  bool
	_opensearchUpdateDirectQueryDataSource       bool
	_opensearchUpdateDomainConfig                bool
	_opensearchUpdateIndex                       bool
	_opensearchUpdatePackage                     bool
	_opensearchUpdatePackageScope                bool
	_opensearchUpdateScheduledAction             bool
	_opensearchUpdateVpcEndpoint                 bool
	_opensearchUpgradeDomain                     bool

	_opensearchAccessPolicies              string
	_opensearchAccount                     string
	_opensearchAction                      string
	_opensearchActionID                    string
	_opensearchActionType                  string
	_opensearchAdvancedOptions             string
	_opensearchAdvancedSecurityOptions     string
	_opensearchAIMLOptions                 string
	_opensearchAppConfigs                  string
	_opensearchApplicationArn              string
	_opensearchARN                         string
	_opensearchAssociationConfiguration    string
	_opensearchAutoTuneOptions             string
	_opensearchChangeId                    string
	_opensearchClientToken                 string
	_opensearchClusterConfig               string
	_opensearchCognitoOptions              string
	_opensearchCommitMessage               string
	_opensearchConnectionAlias             string
	_opensearchConnectionId                string
	_opensearchConnectionMode              string
	_opensearchConnectionProperties        string
	_opensearchDataSourceName              string
	_opensearchDataSourceType              string
	_opensearchDataSources                 string
	_opensearchDescription                 string
	_opensearchDesiredStartTime            string
	_opensearchDomainArn                   string
	_opensearchDomainEndpointOptions       string
	_opensearchDomainName                  string
	_opensearchDomainNames                 []string
	_opensearchDryRun                      string
	_opensearchDryRunId                    string
	_opensearchDryRunMode                  string
	_opensearchEBSOptions                  string
	_opensearchEncryptionAtRestOptions     string
	_opensearchEngineType                  string
	_opensearchEngineVersion               string
	_opensearchFilters                     string
	_opensearchIamIdentityCenterOptions    string
	_opensearchId                          string
	_opensearchIdentityCenterOptions       string
	_opensearchIndexName                   string
	_opensearchIndexSchema                 string
	_opensearchInstanceCount               string
	_opensearchInstanceType                string
	_opensearchIPAddressType               string
	_opensearchKmsKeyArn                   string
	_opensearchLoadDryRunConfig            string
	_opensearchLocalDomainInfo             string
	_opensearchLogPublishingOptions        string
	_opensearchMaintenanceId               string
	_opensearchMaxResults                  string
	_opensearchName                        string
	_opensearchNextToken                   string
	_opensearchNodeId                      string
	_opensearchNodeToNodeEncryptionOptions string
	_opensearchOffPeakWindowOptions        string
	_opensearchOpenSearchArns              []string
	_opensearchOperation                   string
	_opensearchPackageConfiguration        string
	_opensearchPackageDescription          string
	_opensearchPackageEncryptionOptions    string
	_opensearchPackageID                   string
	_opensearchPackageList                 []string
	_opensearchPackageName                 string
	_opensearchPackageSource               string
	_opensearchPackageType                 string
	_opensearchPackageUserList             []string
	_opensearchPackageVendingOptions       string
	_opensearchPerformCheckOnly            string
	_opensearchPrerequisitePackageIDList   []string
	_opensearchRemoteDomainInfo            string
	_opensearchReservationName             string
	_opensearchReservedInstanceId          string
	_opensearchReservedInstanceOfferingId  string
	_opensearchRetrieveAZs                 string
	_opensearchScheduleAt                  string
	_opensearchService                     string
	_opensearchSetAsDefault                string
	_opensearchSnapshotOptions             string
	_opensearchSoftwareUpdateOptions       string
	_opensearchStatus                      string
	_opensearchStatuses                    string
	_opensearchTagKeys                     []string
	_opensearchTagList                     string
	_opensearchTargetVersion               string
	_opensearchVpcEndpointId               string
	_opensearchVpcEndpointIds              []string
	_opensearchVpcOptions                  string
)

// Allows the destination Amazon OpenSearch Service domain owner to accept an
// inbound cross-cluster search connection request. For more information, see [Cross-cluster search for Amazon OpenSearch Service].
//
// [Cross-cluster search for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html
func opensearch_AcceptInboundConnection(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AcceptInboundConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_opensearchConnectionId) > 0 {
		input.ConnectionId = aws.String(_opensearchConnectionId)
	}

	if resp, err := client.AcceptInboundConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new direct-query data source to the specified domain. For more
// information, see [Creating Amazon OpenSearch Service data source integrations with Amazon S3].
//
// [Creating Amazon OpenSearch Service data source integrations with Amazon S3]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/direct-query-s3-creating.html
func opensearch_AddDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AddDataSourceInput{
		// DataSourceType: types.DataSourceType, // Required
		// DomainName: *string, // Required
		// Name: *string, // Required
	}

	if len(_opensearchDataSourceType) > 0 {
		if err := assignInputField(input, "DataSourceType", _opensearchDataSourceType); err != nil {
			log.Errorf("invalid --data-source-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchName) > 0 {
		input.Name = aws.String(_opensearchName)
	}
	if len(_opensearchDescription) > 0 {
		input.Description = aws.String(_opensearchDescription)
	}

	if resp, err := client.AddDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a new data source in Amazon OpenSearch Service so that you can perform
// direct queries on external data.
func opensearch_AddDirectQueryDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AddDirectQueryDataSourceInput{
		// DataSourceName: *string, // Required
		// DataSourceType: types.DirectQueryDataSourceType, // Required
		// OpenSearchArns: []string, // Required
	}

	if len(_opensearchDataSourceName) > 0 {
		input.DataSourceName = aws.String(_opensearchDataSourceName)
	}
	if len(_opensearchDataSourceType) > 0 {
		if err := assignInputField(input, "DataSourceType", _opensearchDataSourceType); err != nil {
			log.Errorf("invalid --data-source-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchOpenSearchArns) > 0 {
		input.OpenSearchArns = append([]string(nil), _opensearchOpenSearchArns...)
	}
	if len(_opensearchDescription) > 0 {
		input.Description = aws.String(_opensearchDescription)
	}
	if len(_opensearchTagList) > 0 {
		if err := assignInputField(input, "TagList", _opensearchTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AddDirectQueryDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Attaches tags to an existing Amazon OpenSearch Service domain, data source, or
// application.
//
// Tags are a set of case-sensitive key-value pairs. A domain, data source, or
// application can have up to 10 tags. For more information, see [Tagging Amazon OpenSearch Service resources].
//
// [Tagging Amazon OpenSearch Service resources]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-awsresourcetagging.html
func opensearch_AddTags(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AddTagsInput{
		// ARN: *string, // Required
		// TagList: []types.Tag, // Required
	}

	if len(_opensearchARN) > 0 {
		input.ARN = aws.String(_opensearchARN)
	}
	if len(_opensearchTagList) > 0 {
		if err := assignInputField(input, "TagList", _opensearchTagList); err != nil {
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

// Associates a package with an Amazon OpenSearch Service domain. For more
// information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_AssociatePackage(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AssociatePackageInput{
		// DomainName: *string, // Required
		// PackageID: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}
	if len(_opensearchAssociationConfiguration) > 0 {
		if err := assignInputField(input, "AssociationConfiguration", _opensearchAssociationConfiguration); err != nil {
			log.Errorf("invalid --association-configuration: %s", err.Error())
			return
		}
	}
	if len(_opensearchPrerequisitePackageIDList) > 0 {
		input.PrerequisitePackageIDList = append([]string(nil), _opensearchPrerequisitePackageIDList...)
	}

	if resp, err := client.AssociatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Operation in the Amazon OpenSearch Service API for associating multiple
// packages with a domain simultaneously.
func opensearch_AssociatePackages(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AssociatePackagesInput{
		// DomainName: *string, // Required
		// PackageList: []types.PackageDetailsForAssociation, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchPackageList) > 0 {
		if err := assignInputField(input, "PackageList", _opensearchPackageList[0]); err != nil {
			log.Errorf("invalid --package-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociatePackages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provides access to an Amazon OpenSearch Service domain through the use of an
// interface VPC endpoint.
func opensearch_AuthorizeVpcEndpointAccess(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.AuthorizeVpcEndpointAccessInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchAccount) > 0 {
		input.Account = aws.String(_opensearchAccount)
	}
	if len(_opensearchService) > 0 {
		if err := assignInputField(input, "Service", _opensearchService); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}

	if resp, err := client.AuthorizeVpcEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a pending configuration change on an Amazon OpenSearch Service domain.
func opensearch_CancelDomainConfigChange(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CancelDomainConfigChangeInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _opensearchDryRun); err != nil {
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

// Cancels a scheduled service software update for an Amazon OpenSearch Service
// domain. You can only perform this operation before the AutomatedUpdateDate and
// when the domain's UpdateStatus is PENDING_UPDATE . For more information, see [Service software updates in Amazon OpenSearch Service].
//
// [Service software updates in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html
func opensearch_CancelServiceSoftwareUpdate(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CancelServiceSoftwareUpdateInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.CancelServiceSoftwareUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an OpenSearch UI application. For more information, see [Using the OpenSearch user interface in Amazon OpenSearch Service].
//
// [Using the OpenSearch user interface in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/application.html
func opensearch_CreateApplication(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CreateApplicationInput{
		// Name: *string, // Required
	}

	if len(_opensearchName) > 0 {
		input.Name = aws.String(_opensearchName)
	}
	if len(_opensearchAppConfigs) > 0 {
		if err := assignInputField(input, "AppConfigs", _opensearchAppConfigs); err != nil {
			log.Errorf("invalid --app-configs: %s", err.Error())
			return
		}
	}
	if len(_opensearchClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchClientToken)
	}
	if len(_opensearchDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _opensearchDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}
	if len(_opensearchIamIdentityCenterOptions) > 0 {
		if err := assignInputField(input, "IamIdentityCenterOptions", _opensearchIamIdentityCenterOptions); err != nil {
			log.Errorf("invalid --iam-identity-center-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_opensearchKmsKeyArn)
	}
	if len(_opensearchTagList) > 0 {
		if err := assignInputField(input, "TagList", _opensearchTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon OpenSearch Service domain. For more information, see [Creating and managing Amazon OpenSearch Service domains].
//
// [Creating and managing Amazon OpenSearch Service domains]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html
func opensearch_CreateDomain(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CreateDomainInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchAIMLOptions) > 0 {
		if err := assignInputField(input, "AIMLOptions", _opensearchAIMLOptions); err != nil {
			log.Errorf("invalid --aiml-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchAccessPolicies) > 0 {
		input.AccessPolicies = aws.String(_opensearchAccessPolicies)
	}
	if len(_opensearchAdvancedOptions) > 0 {
		if err := assignInputField(input, "AdvancedOptions", _opensearchAdvancedOptions); err != nil {
			log.Errorf("invalid --advanced-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchAdvancedSecurityOptions) > 0 {
		if err := assignInputField(input, "AdvancedSecurityOptions", _opensearchAdvancedSecurityOptions); err != nil {
			log.Errorf("invalid --advanced-security-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchAutoTuneOptions) > 0 {
		if err := assignInputField(input, "AutoTuneOptions", _opensearchAutoTuneOptions); err != nil {
			log.Errorf("invalid --auto-tune-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchClusterConfig) > 0 {
		if err := assignInputField(input, "ClusterConfig", _opensearchClusterConfig); err != nil {
			log.Errorf("invalid --cluster-config: %s", err.Error())
			return
		}
	}
	if len(_opensearchCognitoOptions) > 0 {
		if err := assignInputField(input, "CognitoOptions", _opensearchCognitoOptions); err != nil {
			log.Errorf("invalid --cognito-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainEndpointOptions) > 0 {
		if err := assignInputField(input, "DomainEndpointOptions", _opensearchDomainEndpointOptions); err != nil {
			log.Errorf("invalid --domain-endpoint-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchEBSOptions) > 0 {
		if err := assignInputField(input, "EBSOptions", _opensearchEBSOptions); err != nil {
			log.Errorf("invalid --ebs-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchEncryptionAtRestOptions) > 0 {
		if err := assignInputField(input, "EncryptionAtRestOptions", _opensearchEncryptionAtRestOptions); err != nil {
			log.Errorf("invalid --encryption-at-rest-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchEngineVersion) > 0 {
		input.EngineVersion = aws.String(_opensearchEngineVersion)
	}
	if len(_opensearchIPAddressType) > 0 {
		if err := assignInputField(input, "IPAddressType", _opensearchIPAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchIdentityCenterOptions) > 0 {
		if err := assignInputField(input, "IdentityCenterOptions", _opensearchIdentityCenterOptions); err != nil {
			log.Errorf("invalid --identity-center-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchLogPublishingOptions) > 0 {
		if err := assignInputField(input, "LogPublishingOptions", _opensearchLogPublishingOptions); err != nil {
			log.Errorf("invalid --log-publishing-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchNodeToNodeEncryptionOptions) > 0 {
		if err := assignInputField(input, "NodeToNodeEncryptionOptions", _opensearchNodeToNodeEncryptionOptions); err != nil {
			log.Errorf("invalid --node-to-node-encryption-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchOffPeakWindowOptions) > 0 {
		if err := assignInputField(input, "OffPeakWindowOptions", _opensearchOffPeakWindowOptions); err != nil {
			log.Errorf("invalid --off-peak-window-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchSnapshotOptions) > 0 {
		if err := assignInputField(input, "SnapshotOptions", _opensearchSnapshotOptions); err != nil {
			log.Errorf("invalid --snapshot-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchSoftwareUpdateOptions) > 0 {
		if err := assignInputField(input, "SoftwareUpdateOptions", _opensearchSoftwareUpdateOptions); err != nil {
			log.Errorf("invalid --software-update-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchTagList) > 0 {
		if err := assignInputField(input, "TagList", _opensearchTagList); err != nil {
			log.Errorf("invalid --tag-list: %s", err.Error())
			return
		}
	}
	if len(_opensearchVpcOptions) > 0 {
		if err := assignInputField(input, "VPCOptions", _opensearchVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an OpenSearch index with optional automatic semantic enrichment for
// specified text fields. Automatic semantic enrichment enables semantic search
// capabilities without requiring machine learning expertise, improving search
// relevance by up to 20% by understanding search intent and contextual meaning
// beyond keyword matching. The semantic enrichment process has zero impact on
// search latency as sparse encodings are stored directly within the index during
// indexing. For more information, see [Automatic semantic enrichment].
//
// [Automatic semantic enrichment]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/opensearch-semantic-enrichment.html
func opensearch_CreateIndex(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CreateIndexInput{
		// DomainName: *string, // Required
		// IndexName: *string, // Required
		// IndexSchema: document.Interface, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchIndexName) > 0 {
		input.IndexName = aws.String(_opensearchIndexName)
	}
	if len(_opensearchIndexSchema) > 0 {
		if err := assignInputField(input, "IndexSchema", _opensearchIndexSchema); err != nil {
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

// Creates a new cross-cluster search connection from a source Amazon OpenSearch
// Service domain to a destination domain. For more information, see [Cross-cluster search for Amazon OpenSearch Service].
//
// [Cross-cluster search for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html
func opensearch_CreateOutboundConnection(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CreateOutboundConnectionInput{
		// ConnectionAlias: *string, // Required
		// LocalDomainInfo: *types.DomainInformationContainer, // Required
		// RemoteDomainInfo: *types.DomainInformationContainer, // Required
	}

	if len(_opensearchConnectionAlias) > 0 {
		input.ConnectionAlias = aws.String(_opensearchConnectionAlias)
	}
	if len(_opensearchLocalDomainInfo) > 0 {
		if err := assignInputField(input, "LocalDomainInfo", _opensearchLocalDomainInfo); err != nil {
			log.Errorf("invalid --local-domain-info: %s", err.Error())
			return
		}
	}
	if len(_opensearchRemoteDomainInfo) > 0 {
		if err := assignInputField(input, "RemoteDomainInfo", _opensearchRemoteDomainInfo); err != nil {
			log.Errorf("invalid --remote-domain-info: %s", err.Error())
			return
		}
	}
	if len(_opensearchConnectionMode) > 0 {
		if err := assignInputField(input, "ConnectionMode", _opensearchConnectionMode); err != nil {
			log.Errorf("invalid --connection-mode: %s", err.Error())
			return
		}
	}
	if len(_opensearchConnectionProperties) > 0 {
		if err := assignInputField(input, "ConnectionProperties", _opensearchConnectionProperties); err != nil {
			log.Errorf("invalid --connection-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateOutboundConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a package for use with Amazon OpenSearch Service domains. For more
// information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_CreatePackage(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CreatePackageInput{
		// PackageName: *string, // Required
		// PackageSource: *types.PackageSource, // Required
		// PackageType: types.PackageType, // Required
	}

	if len(_opensearchPackageName) > 0 {
		input.PackageName = aws.String(_opensearchPackageName)
	}
	if len(_opensearchPackageSource) > 0 {
		if err := assignInputField(input, "PackageSource", _opensearchPackageSource); err != nil {
			log.Errorf("invalid --package-source: %s", err.Error())
			return
		}
	}
	if len(_opensearchPackageType) > 0 {
		if err := assignInputField(input, "PackageType", _opensearchPackageType); err != nil {
			log.Errorf("invalid --package-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchEngineVersion) > 0 {
		input.EngineVersion = aws.String(_opensearchEngineVersion)
	}
	if len(_opensearchPackageConfiguration) > 0 {
		if err := assignInputField(input, "PackageConfiguration", _opensearchPackageConfiguration); err != nil {
			log.Errorf("invalid --package-configuration: %s", err.Error())
			return
		}
	}
	if len(_opensearchPackageDescription) > 0 {
		input.PackageDescription = aws.String(_opensearchPackageDescription)
	}
	if len(_opensearchPackageEncryptionOptions) > 0 {
		if err := assignInputField(input, "PackageEncryptionOptions", _opensearchPackageEncryptionOptions); err != nil {
			log.Errorf("invalid --package-encryption-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchPackageVendingOptions) > 0 {
		if err := assignInputField(input, "PackageVendingOptions", _opensearchPackageVendingOptions); err != nil {
			log.Errorf("invalid --package-vending-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an Amazon OpenSearch Service-managed VPC endpoint.
func opensearch_CreateVpcEndpoint(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.CreateVpcEndpointInput{
		// DomainArn: *string, // Required
		// VpcOptions: *types.VPCOptions, // Required
	}

	if len(_opensearchDomainArn) > 0 {
		input.DomainArn = aws.String(_opensearchDomainArn)
	}
	if len(_opensearchVpcOptions) > 0 {
		if err := assignInputField(input, "VpcOptions", _opensearchVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchClientToken) > 0 {
		input.ClientToken = aws.String(_opensearchClientToken)
	}

	if resp, err := client.CreateVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a specified OpenSearch application.
func opensearch_DeleteApplication(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteApplicationInput{
		// Id: *string, // Required
	}

	if len(_opensearchId) > 0 {
		input.Id = aws.String(_opensearchId)
	}

	if resp, err := client.DeleteApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a direct-query data source. For more information, see [Deleting an Amazon OpenSearch Service data source with Amazon S3].
//
// [Deleting an Amazon OpenSearch Service data source with Amazon S3]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/direct-query-s3-delete.html
func opensearch_DeleteDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteDataSourceInput{
		// DomainName: *string, // Required
		// Name: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchName) > 0 {
		input.Name = aws.String(_opensearchName)
	}

	if resp, err := client.DeleteDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a previously configured direct query data source from Amazon
// OpenSearch Service.
func opensearch_DeleteDirectQueryDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteDirectQueryDataSourceInput{
		// DataSourceName: *string, // Required
	}

	if len(_opensearchDataSourceName) > 0 {
		input.DataSourceName = aws.String(_opensearchDataSourceName)
	}

	if resp, err := client.DeleteDirectQueryDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon OpenSearch Service domain and all of its data. You can't
// recover a domain after you delete it.
func opensearch_DeleteDomain(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteDomainInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.DeleteDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the destination Amazon OpenSearch Service domain owner to delete an
// existing inbound cross-cluster search connection. For more information, see [Cross-cluster search for Amazon OpenSearch Service].
//
// [Cross-cluster search for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html
func opensearch_DeleteInboundConnection(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteInboundConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_opensearchConnectionId) > 0 {
		input.ConnectionId = aws.String(_opensearchConnectionId)
	}

	if resp, err := client.DeleteInboundConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an OpenSearch index. This operation permanently removes the index and
// cannot be undone.
func opensearch_DeleteIndex(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteIndexInput{
		// DomainName: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchIndexName) > 0 {
		input.IndexName = aws.String(_opensearchIndexName)
	}

	if resp, err := client.DeleteIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the source Amazon OpenSearch Service domain owner to delete an existing
// outbound cross-cluster search connection. For more information, see [Cross-cluster search for Amazon OpenSearch Service].
//
// [Cross-cluster search for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html
func opensearch_DeleteOutboundConnection(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteOutboundConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_opensearchConnectionId) > 0 {
		input.ConnectionId = aws.String(_opensearchConnectionId)
	}

	if resp, err := client.DeleteOutboundConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon OpenSearch Service package. For more information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_DeletePackage(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeletePackageInput{
		// PackageID: *string, // Required
	}

	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}

	if resp, err := client.DeletePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Amazon OpenSearch Service-managed interface VPC endpoint.
func opensearch_DeleteVpcEndpoint(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DeleteVpcEndpointInput{
		// VpcEndpointId: *string, // Required
	}

	if len(_opensearchVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_opensearchVpcEndpointId)
	}

	if resp, err := client.DeleteVpcEndpoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the domain configuration for the specified Amazon OpenSearch Service
// domain, including the domain ID, domain service endpoint, and domain ARN.
func opensearch_DescribeDomain(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.DescribeDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of optimizations that Auto-Tune has made to an Amazon
// OpenSearch Service domain. For more information, see [Auto-Tune for Amazon OpenSearch Service].
//
// [Auto-Tune for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html
func opensearch_DescribeDomainAutoTunes(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainAutoTunesInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.DescribeDomainAutoTunesOutput
	p := opensearch.NewDescribeDomainAutoTunesPaginator(client, input)
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

// Returns information about the current blue/green deployment happening on an
// Amazon OpenSearch Service domain. For more information, see [Making configuration changes in Amazon OpenSearch Service].
//
// [Making configuration changes in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-configuration-changes.html
func opensearch_DescribeDomainChangeProgress(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainChangeProgressInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchChangeId) > 0 {
		input.ChangeId = aws.String(_opensearchChangeId)
	}

	if resp, err := client.DescribeDomainChangeProgress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the configuration of an Amazon OpenSearch Service domain.
func opensearch_DescribeDomainConfig(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainConfigInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.DescribeDomainConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about domain and node health, the standby Availability
// Zone, number of nodes per Availability Zone, and shard count per node.
func opensearch_DescribeDomainHealth(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainHealthInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.DescribeDomainHealth(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about domain and nodes, including data nodes, master nodes,
// ultrawarm nodes, Availability Zone(s), standby nodes, node configurations, and
// node states.
func opensearch_DescribeDomainNodes(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainNodesInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.DescribeDomainNodes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns domain configuration information about the specified Amazon OpenSearch
// Service domains.
func opensearch_DescribeDomains(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDomainsInput{
		// DomainNames: []string, // Required
	}

	if len(_opensearchDomainNames) > 0 {
		input.DomainNames = append([]string(nil), _opensearchDomainNames...)
	}

	if resp, err := client.DescribeDomains(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the progress of a pre-update dry run analysis on an Amazon OpenSearch
// Service domain. For more information, see [Determining whether a change will cause a blue/green deployment].
//
// [Determining whether a change will cause a blue/green deployment]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-configuration-changes#dryrun
func opensearch_DescribeDryRunProgress(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeDryRunProgressInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchDryRunId) > 0 {
		input.DryRunId = aws.String(_opensearchDryRunId)
	}
	if len(_opensearchLoadDryRunConfig) > 0 {
		if err := assignInputField(input, "LoadDryRunConfig", _opensearchLoadDryRunConfig); err != nil {
			log.Errorf("invalid --load-dry-run-config: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeDryRunProgress(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the inbound cross-cluster search connections for a destination
// (remote) Amazon OpenSearch Service domain. For more information, see [Cross-cluster search for Amazon OpenSearch Service].
//
// [Cross-cluster search for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html
func opensearch_DescribeInboundConnections(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeInboundConnectionsInput{}

	if len(_opensearchFilters) > 0 {
		if err := assignInputField(input, "Filters", _opensearchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeInboundConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.DescribeInboundConnectionsOutput
	p := opensearch.NewDescribeInboundConnectionsPaginator(client, input)
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

// Describes the instance count, storage, and master node limits for a given
// OpenSearch or Elasticsearch version and instance type.
func opensearch_DescribeInstanceTypeLimits(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeInstanceTypeLimitsInput{
		// EngineVersion: *string, // Required
		// InstanceType: types.OpenSearchPartitionInstanceType, // Required
	}

	if len(_opensearchEngineVersion) > 0 {
		input.EngineVersion = aws.String(_opensearchEngineVersion)
	}
	if len(_opensearchInstanceType) > 0 {
		if err := assignInputField(input, "InstanceType", _opensearchInstanceType); err != nil {
			log.Errorf("invalid --instance-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.DescribeInstanceTypeLimits(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all the outbound cross-cluster connections for a local (source) Amazon
// OpenSearch Service domain. For more information, see [Cross-cluster search for Amazon OpenSearch Service].
//
// [Cross-cluster search for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cross-cluster-search.html
func opensearch_DescribeOutboundConnections(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeOutboundConnectionsInput{}

	if len(_opensearchFilters) > 0 {
		if err := assignInputField(input, "Filters", _opensearchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeOutboundConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.DescribeOutboundConnectionsOutput
	p := opensearch.NewDescribeOutboundConnectionsPaginator(client, input)
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

// Describes all packages available to OpenSearch Service. For more information,
// see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_DescribePackages(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribePackagesInput{}

	if len(_opensearchFilters) > 0 {
		if err := assignInputField(input, "Filters", _opensearchFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.DescribePackagesOutput
	p := opensearch.NewDescribePackagesPaginator(client, input)
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

// Describes the available Amazon OpenSearch Service Reserved Instance offerings
// for a given Region. For more information, see [Reserved Instances in Amazon OpenSearch Service].
//
// [Reserved Instances in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ri.html
func opensearch_DescribeReservedInstanceOfferings(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeReservedInstanceOfferingsInput{}

	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}
	if len(_opensearchReservedInstanceOfferingId) > 0 {
		input.ReservedInstanceOfferingId = aws.String(_opensearchReservedInstanceOfferingId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedInstanceOfferings(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.DescribeReservedInstanceOfferingsOutput
	p := opensearch.NewDescribeReservedInstanceOfferingsPaginator(client, input)
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

// Describes the Amazon OpenSearch Service instances that you have reserved in a
// given Region. For more information, see [Reserved Instances in Amazon OpenSearch Service].
//
// [Reserved Instances in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/ri.html
func opensearch_DescribeReservedInstances(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeReservedInstancesInput{}

	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}
	if len(_opensearchReservedInstanceId) > 0 {
		input.ReservedInstanceId = aws.String(_opensearchReservedInstanceId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReservedInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.DescribeReservedInstancesOutput
	p := opensearch.NewDescribeReservedInstancesPaginator(client, input)
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
func opensearch_DescribeVpcEndpoints(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DescribeVpcEndpointsInput{
		// VpcEndpointIds: []string, // Required
	}

	if len(_opensearchVpcEndpointIds) > 0 {
		input.VpcEndpointIds = append([]string(nil), _opensearchVpcEndpointIds...)
	}

	if resp, err := client.DescribeVpcEndpoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a package from the specified Amazon OpenSearch Service domain. The
// package can't be in use with any OpenSearch index for the dissociation to
// succeed. The package is still available in OpenSearch Service for association
// later. For more information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_DissociatePackage(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DissociatePackageInput{
		// DomainName: *string, // Required
		// PackageID: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}

	if resp, err := client.DissociatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dissociates multiple packages from a domain simultaneously.
func opensearch_DissociatePackages(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.DissociatePackagesInput{
		// DomainName: *string, // Required
		// PackageList: []string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchPackageList) > 0 {
		input.PackageList = append([]string(nil), _opensearchPackageList...)
	}

	if resp, err := client.DissociatePackages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration and status of an existing OpenSearch application.
func opensearch_GetApplication(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetApplicationInput{
		// Id: *string, // Required
	}

	if len(_opensearchId) > 0 {
		input.Id = aws.String(_opensearchId)
	}

	if resp, err := client.GetApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a map of OpenSearch or Elasticsearch versions and the versions you can
// upgrade them to.
func opensearch_GetCompatibleVersions(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetCompatibleVersionsInput{}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.GetCompatibleVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a direct query data source.
func opensearch_GetDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetDataSourceInput{
		// DomainName: *string, // Required
		// Name: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchName) > 0 {
		input.Name = aws.String(_opensearchName)
	}

	if resp, err := client.GetDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the ARN of the current default application.
// If the default application isn't set, the operation returns a resource not
// found error.
func opensearch_GetDefaultApplicationSetting(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetDefaultApplicationSettingInput{}

	if resp, err := client.GetDefaultApplicationSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed configuration information for a specific direct query data
// source in Amazon OpenSearch Service.
func opensearch_GetDirectQueryDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetDirectQueryDataSourceInput{
		// DataSourceName: *string, // Required
	}

	if len(_opensearchDataSourceName) > 0 {
		input.DataSourceName = aws.String(_opensearchDataSourceName)
	}

	if resp, err := client.GetDirectQueryDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The status of the maintenance action.
func opensearch_GetDomainMaintenanceStatus(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetDomainMaintenanceStatusInput{
		// DomainName: *string, // Required
		// MaintenanceId: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchMaintenanceId) > 0 {
		input.MaintenanceId = aws.String(_opensearchMaintenanceId)
	}

	if resp, err := client.GetDomainMaintenanceStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about an OpenSearch index including its schema and
// semantic enrichment configuration. Use this operation to view the current index
// structure and semantic search settings.
func opensearch_GetIndex(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetIndexInput{
		// DomainName: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchIndexName) > 0 {
		input.IndexName = aws.String(_opensearchIndexName)
	}

	if resp, err := client.GetIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of Amazon OpenSearch Service package versions, along with their
// creation time, commit message, and plugin properties (if the package is a zip
// plugin package). For more information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_GetPackageVersionHistory(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetPackageVersionHistoryInput{
		// PackageID: *string, // Required
	}

	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.GetPackageVersionHistoryOutput
	p := opensearch.NewGetPackageVersionHistoryPaginator(client, input)
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

// Retrieves the complete history of the last 10 upgrades performed on an Amazon
// OpenSearch Service domain.
func opensearch_GetUpgradeHistory(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetUpgradeHistoryInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.GetUpgradeHistoryOutput
	p := opensearch.NewGetUpgradeHistoryPaginator(client, input)
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

// Returns the most recent status of the last upgrade or upgrade eligibility check
// performed on an Amazon OpenSearch Service domain.
func opensearch_GetUpgradeStatus(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.GetUpgradeStatusInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.GetUpgradeStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all OpenSearch applications under your account.
func opensearch_ListApplications(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListApplicationsInput{}

	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}
	if len(_opensearchStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _opensearchStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListApplications(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.ListApplicationsOutput
	p := opensearch.NewListApplicationsPaginator(client, input)
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

// Lists direct-query data sources for a specific domain. For more information,
// see For more information, see [Working with Amazon OpenSearch Service direct queries with Amazon S3].
//
// [Working with Amazon OpenSearch Service direct queries with Amazon S3]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/direct-query-s3.html
func opensearch_ListDataSources(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListDataSourcesInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}

	if resp, err := client.ListDataSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists an inventory of all the direct query data sources that you have
// configured within Amazon OpenSearch Service.
func opensearch_ListDirectQueryDataSources(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListDirectQueryDataSourcesInput{}

	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}

	if resp, err := client.ListDirectQueryDataSources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A list of maintenance actions for the domain.
func opensearch_ListDomainMaintenances(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListDomainMaintenancesInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchAction) > 0 {
		if err := assignInputField(input, "Action", _opensearchAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}
	if len(_opensearchStatus) > 0 {
		if err := assignInputField(input, "Status", _opensearchStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDomainMaintenances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.ListDomainMaintenancesOutput
	p := opensearch.NewListDomainMaintenancesPaginator(client, input)
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

// Returns the names of all Amazon OpenSearch Service domains owned by the current
// user in the active Region.
func opensearch_ListDomainNames(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListDomainNamesInput{}

	if len(_opensearchEngineType) > 0 {
		if err := assignInputField(input, "EngineType", _opensearchEngineType); err != nil {
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

// Lists all Amazon OpenSearch Service domains associated with a given package.
// For more information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_ListDomainsForPackage(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListDomainsForPackageInput{
		// PackageID: *string, // Required
	}

	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.ListDomainsForPackageOutput
	p := opensearch.NewListDomainsForPackagePaginator(client, input)
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

// Lists all instance types and available features for a given OpenSearch or
// Elasticsearch version.
func opensearch_ListInstanceTypeDetails(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListInstanceTypeDetailsInput{
		// EngineVersion: *string, // Required
	}

	if len(_opensearchEngineVersion) > 0 {
		input.EngineVersion = aws.String(_opensearchEngineVersion)
	}
	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchInstanceType) > 0 {
		input.InstanceType = aws.String(_opensearchInstanceType)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}
	if len(_opensearchRetrieveAZs) > 0 {
		if err := assignInputField(input, "RetrieveAZs", _opensearchRetrieveAZs); err != nil {
			log.Errorf("invalid --retrieve-azs: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListInstanceTypeDetails(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.ListInstanceTypeDetailsOutput
	p := opensearch.NewListInstanceTypeDetailsPaginator(client, input)
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

// Lists all packages associated with an Amazon OpenSearch Service domain. For
// more information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_ListPackagesForDomain(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListPackagesForDomainInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.ListPackagesForDomainOutput
	p := opensearch.NewListPackagesForDomainPaginator(client, input)
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

// Retrieves a list of configuration changes that are scheduled for a domain.
// These changes can be [service software updates]or [blue/green Auto-Tune enhancements].
//
// [blue/green Auto-Tune enhancements]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html#auto-tune-types
// [service software updates]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html
func opensearch_ListScheduledActions(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListScheduledActionsInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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

	var results []*opensearch.ListScheduledActionsOutput
	p := opensearch.NewListScheduledActionsPaginator(client, input)
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

// Returns all resource tags for an Amazon OpenSearch Service domain, data source,
// or application. For more information, see [Tagging Amazon OpenSearch Service resources].
//
// [Tagging Amazon OpenSearch Service resources]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains-awsresourcetagging.html
func opensearch_ListTags(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListTagsInput{
		// ARN: *string, // Required
	}

	if len(_opensearchARN) > 0 {
		input.ARN = aws.String(_opensearchARN)
	}

	if resp, err := client.ListTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all versions of OpenSearch and Elasticsearch that Amazon OpenSearch
// Service supports.
func opensearch_ListVersions(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListVersionsInput{}

	if len(_opensearchMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _opensearchMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*opensearch.ListVersionsOutput
	p := opensearch.NewListVersionsPaginator(client, input)
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

// Retrieves information about each Amazon Web Services principal that is allowed
// to access a given Amazon OpenSearch Service domain through the use of an
// interface VPC endpoint.
func opensearch_ListVpcEndpointAccess(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListVpcEndpointAccessInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}

	if resp, err := client.ListVpcEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all Amazon OpenSearch Service-managed VPC endpoints in the current
// Amazon Web Services account and Region.
func opensearch_ListVpcEndpoints(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListVpcEndpointsInput{}

	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
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
func opensearch_ListVpcEndpointsForDomain(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.ListVpcEndpointsForDomainInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchNextToken) > 0 {
		input.NextToken = aws.String(_opensearchNextToken)
	}

	if resp, err := client.ListVpcEndpointsForDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows you to purchase Amazon OpenSearch Service Reserved Instances.
func opensearch_PurchaseReservedInstanceOffering(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.PurchaseReservedInstanceOfferingInput{
		// ReservationName: *string, // Required
		// ReservedInstanceOfferingId: *string, // Required
	}

	if len(_opensearchReservationName) > 0 {
		input.ReservationName = aws.String(_opensearchReservationName)
	}
	if len(_opensearchReservedInstanceOfferingId) > 0 {
		input.ReservedInstanceOfferingId = aws.String(_opensearchReservedInstanceOfferingId)
	}
	if len(_opensearchInstanceCount) > 0 {
		if err := assignInputField(input, "InstanceCount", _opensearchInstanceCount); err != nil {
			log.Errorf("invalid --instance-count: %s", err.Error())
			return
		}
	}

	if resp, err := client.PurchaseReservedInstanceOffering(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the default application to the application with the specified ARN.
// To remove the default application, use the GetDefaultApplicationSetting
// operation to get the current default and then call the
// PutDefaultApplicationSetting with the current applications ARN and the
// setAsDefault parameter set to false .
func opensearch_PutDefaultApplicationSetting(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.PutDefaultApplicationSettingInput{
		// ApplicationArn: *string, // Required
		// SetAsDefault: *bool, // Required
	}

	if len(_opensearchApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_opensearchApplicationArn)
	}
	if len(_opensearchSetAsDefault) > 0 {
		if err := assignInputField(input, "SetAsDefault", _opensearchSetAsDefault); err != nil {
			log.Errorf("invalid --set-as-default: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutDefaultApplicationSetting(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Allows the remote Amazon OpenSearch Service domain owner to reject an inbound
// cross-cluster connection request.
func opensearch_RejectInboundConnection(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.RejectInboundConnectionInput{
		// ConnectionId: *string, // Required
	}

	if len(_opensearchConnectionId) > 0 {
		input.ConnectionId = aws.String(_opensearchConnectionId)
	}

	if resp, err := client.RejectInboundConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified set of tags from an Amazon OpenSearch Service domain,
// data source, or application. For more information, see [Tagging Amazon OpenSearch Service resources].
//
// [Tagging Amazon OpenSearch Service resources]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/managedomains.html#managedomains-awsresorcetagging
func opensearch_RemoveTags(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.RemoveTagsInput{
		// ARN: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_opensearchARN) > 0 {
		input.ARN = aws.String(_opensearchARN)
	}
	if len(_opensearchTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _opensearchTagKeys...)
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
func opensearch_RevokeVpcEndpointAccess(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.RevokeVpcEndpointAccessInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchAccount) > 0 {
		input.Account = aws.String(_opensearchAccount)
	}
	if len(_opensearchService) > 0 {
		if err := assignInputField(input, "Service", _opensearchService); err != nil {
			log.Errorf("invalid --service: %s", err.Error())
			return
		}
	}

	if resp, err := client.RevokeVpcEndpointAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the node maintenance process on the data node. These processes can
// include a node reboot, an Opensearch or Elasticsearch process restart, or a
// Dashboard or Kibana restart.
func opensearch_StartDomainMaintenance(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.StartDomainMaintenanceInput{
		// Action: types.MaintenanceType, // Required
		// DomainName: *string, // Required
	}

	if len(_opensearchAction) > 0 {
		if err := assignInputField(input, "Action", _opensearchAction); err != nil {
			log.Errorf("invalid --action: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchNodeId) > 0 {
		input.NodeId = aws.String(_opensearchNodeId)
	}

	if resp, err := client.StartDomainMaintenance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Schedules a service software update for an Amazon OpenSearch Service domain.
// For more information, see [Service software updates in Amazon OpenSearch Service].
//
// [Service software updates in Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html
func opensearch_StartServiceSoftwareUpdate(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.StartServiceSoftwareUpdateInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchDesiredStartTime) > 0 {
		if err := assignInputField(input, "DesiredStartTime", _opensearchDesiredStartTime); err != nil {
			log.Errorf("invalid --desired-start-time: %s", err.Error())
			return
		}
	}
	if len(_opensearchScheduleAt) > 0 {
		if err := assignInputField(input, "ScheduleAt", _opensearchScheduleAt); err != nil {
			log.Errorf("invalid --schedule-at: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartServiceSoftwareUpdate(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration and settings of an existing OpenSearch application.
func opensearch_UpdateApplication(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateApplicationInput{
		// Id: *string, // Required
	}

	if len(_opensearchId) > 0 {
		input.Id = aws.String(_opensearchId)
	}
	if len(_opensearchAppConfigs) > 0 {
		if err := assignInputField(input, "AppConfigs", _opensearchAppConfigs); err != nil {
			log.Errorf("invalid --app-configs: %s", err.Error())
			return
		}
	}
	if len(_opensearchDataSources) > 0 {
		if err := assignInputField(input, "DataSources", _opensearchDataSources); err != nil {
			log.Errorf("invalid --data-sources: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a direct-query data source. For more information, see [Working with Amazon OpenSearch Service data source integrations with Amazon S3].
//
// [Working with Amazon OpenSearch Service data source integrations with Amazon S3]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/direct-query-s3-creating.html
func opensearch_UpdateDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateDataSourceInput{
		// DataSourceType: types.DataSourceType, // Required
		// DomainName: *string, // Required
		// Name: *string, // Required
	}

	if len(_opensearchDataSourceType) > 0 {
		if err := assignInputField(input, "DataSourceType", _opensearchDataSourceType); err != nil {
			log.Errorf("invalid --data-source-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchName) > 0 {
		input.Name = aws.String(_opensearchName)
	}
	if len(_opensearchDescription) > 0 {
		input.Description = aws.String(_opensearchDescription)
	}
	if len(_opensearchStatus) > 0 {
		if err := assignInputField(input, "Status", _opensearchStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration or properties of an existing direct query data
// source in Amazon OpenSearch Service.
func opensearch_UpdateDirectQueryDataSource(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateDirectQueryDataSourceInput{
		// DataSourceName: *string, // Required
		// DataSourceType: types.DirectQueryDataSourceType, // Required
		// OpenSearchArns: []string, // Required
	}

	if len(_opensearchDataSourceName) > 0 {
		input.DataSourceName = aws.String(_opensearchDataSourceName)
	}
	if len(_opensearchDataSourceType) > 0 {
		if err := assignInputField(input, "DataSourceType", _opensearchDataSourceType); err != nil {
			log.Errorf("invalid --data-source-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchOpenSearchArns) > 0 {
		input.OpenSearchArns = append([]string(nil), _opensearchOpenSearchArns...)
	}
	if len(_opensearchDescription) > 0 {
		input.Description = aws.String(_opensearchDescription)
	}

	if resp, err := client.UpdateDirectQueryDataSource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the cluster configuration of the specified Amazon OpenSearch Service
// domain.
func opensearch_UpdateDomainConfig(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateDomainConfigInput{
		// DomainName: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchAIMLOptions) > 0 {
		if err := assignInputField(input, "AIMLOptions", _opensearchAIMLOptions); err != nil {
			log.Errorf("invalid --aiml-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchAccessPolicies) > 0 {
		input.AccessPolicies = aws.String(_opensearchAccessPolicies)
	}
	if len(_opensearchAdvancedOptions) > 0 {
		if err := assignInputField(input, "AdvancedOptions", _opensearchAdvancedOptions); err != nil {
			log.Errorf("invalid --advanced-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchAdvancedSecurityOptions) > 0 {
		if err := assignInputField(input, "AdvancedSecurityOptions", _opensearchAdvancedSecurityOptions); err != nil {
			log.Errorf("invalid --advanced-security-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchAutoTuneOptions) > 0 {
		if err := assignInputField(input, "AutoTuneOptions", _opensearchAutoTuneOptions); err != nil {
			log.Errorf("invalid --auto-tune-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchClusterConfig) > 0 {
		if err := assignInputField(input, "ClusterConfig", _opensearchClusterConfig); err != nil {
			log.Errorf("invalid --cluster-config: %s", err.Error())
			return
		}
	}
	if len(_opensearchCognitoOptions) > 0 {
		if err := assignInputField(input, "CognitoOptions", _opensearchCognitoOptions); err != nil {
			log.Errorf("invalid --cognito-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainEndpointOptions) > 0 {
		if err := assignInputField(input, "DomainEndpointOptions", _opensearchDomainEndpointOptions); err != nil {
			log.Errorf("invalid --domain-endpoint-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _opensearchDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}
	if len(_opensearchDryRunMode) > 0 {
		if err := assignInputField(input, "DryRunMode", _opensearchDryRunMode); err != nil {
			log.Errorf("invalid --dry-run-mode: %s", err.Error())
			return
		}
	}
	if len(_opensearchEBSOptions) > 0 {
		if err := assignInputField(input, "EBSOptions", _opensearchEBSOptions); err != nil {
			log.Errorf("invalid --ebs-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchEncryptionAtRestOptions) > 0 {
		if err := assignInputField(input, "EncryptionAtRestOptions", _opensearchEncryptionAtRestOptions); err != nil {
			log.Errorf("invalid --encryption-at-rest-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchIPAddressType) > 0 {
		if err := assignInputField(input, "IPAddressType", _opensearchIPAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchIdentityCenterOptions) > 0 {
		if err := assignInputField(input, "IdentityCenterOptions", _opensearchIdentityCenterOptions); err != nil {
			log.Errorf("invalid --identity-center-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchLogPublishingOptions) > 0 {
		if err := assignInputField(input, "LogPublishingOptions", _opensearchLogPublishingOptions); err != nil {
			log.Errorf("invalid --log-publishing-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchNodeToNodeEncryptionOptions) > 0 {
		if err := assignInputField(input, "NodeToNodeEncryptionOptions", _opensearchNodeToNodeEncryptionOptions); err != nil {
			log.Errorf("invalid --node-to-node-encryption-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchOffPeakWindowOptions) > 0 {
		if err := assignInputField(input, "OffPeakWindowOptions", _opensearchOffPeakWindowOptions); err != nil {
			log.Errorf("invalid --off-peak-window-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchSnapshotOptions) > 0 {
		if err := assignInputField(input, "SnapshotOptions", _opensearchSnapshotOptions); err != nil {
			log.Errorf("invalid --snapshot-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchSoftwareUpdateOptions) > 0 {
		if err := assignInputField(input, "SoftwareUpdateOptions", _opensearchSoftwareUpdateOptions); err != nil {
			log.Errorf("invalid --software-update-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchVpcOptions) > 0 {
		if err := assignInputField(input, "VPCOptions", _opensearchVpcOptions); err != nil {
			log.Errorf("invalid --vpc-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateDomainConfig(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing OpenSearch index schema and semantic enrichment
// configuration. This operation allows modification of field mappings and semantic
// search settings for text fields. Changes to semantic enrichment configuration
// will apply to newly ingested documents.
func opensearch_UpdateIndex(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateIndexInput{
		// DomainName: *string, // Required
		// IndexName: *string, // Required
		// IndexSchema: document.Interface, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchIndexName) > 0 {
		input.IndexName = aws.String(_opensearchIndexName)
	}
	if len(_opensearchIndexSchema) > 0 {
		if err := assignInputField(input, "IndexSchema", _opensearchIndexSchema); err != nil {
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

// Updates a package for use with Amazon OpenSearch Service domains. For more
// information, see [Custom packages for Amazon OpenSearch Service].
//
// [Custom packages for Amazon OpenSearch Service]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/custom-packages.html
func opensearch_UpdatePackage(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdatePackageInput{
		// PackageID: *string, // Required
		// PackageSource: *types.PackageSource, // Required
	}

	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}
	if len(_opensearchPackageSource) > 0 {
		if err := assignInputField(input, "PackageSource", _opensearchPackageSource); err != nil {
			log.Errorf("invalid --package-source: %s", err.Error())
			return
		}
	}
	if len(_opensearchCommitMessage) > 0 {
		input.CommitMessage = aws.String(_opensearchCommitMessage)
	}
	if len(_opensearchPackageConfiguration) > 0 {
		if err := assignInputField(input, "PackageConfiguration", _opensearchPackageConfiguration); err != nil {
			log.Errorf("invalid --package-configuration: %s", err.Error())
			return
		}
	}
	if len(_opensearchPackageDescription) > 0 {
		input.PackageDescription = aws.String(_opensearchPackageDescription)
	}
	if len(_opensearchPackageEncryptionOptions) > 0 {
		if err := assignInputField(input, "PackageEncryptionOptions", _opensearchPackageEncryptionOptions); err != nil {
			log.Errorf("invalid --package-encryption-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePackage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the scope of a package. Scope of the package defines users who can view
// and associate a package.
func opensearch_UpdatePackageScope(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdatePackageScopeInput{
		// Operation: types.PackageScopeOperationEnum, // Required
		// PackageID: *string, // Required
		// PackageUserList: []string, // Required
	}

	if len(_opensearchOperation) > 0 {
		if err := assignInputField(input, "Operation", _opensearchOperation); err != nil {
			log.Errorf("invalid --operation: %s", err.Error())
			return
		}
	}
	if len(_opensearchPackageID) > 0 {
		input.PackageID = aws.String(_opensearchPackageID)
	}
	if len(_opensearchPackageUserList) > 0 {
		input.PackageUserList = append([]string(nil), _opensearchPackageUserList...)
	}

	if resp, err := client.UpdatePackageScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Reschedules a planned domain configuration change for a later time. This change
// can be a scheduled [service software update]or a [blue/green Auto-Tune enhancement].
//
// [service software update]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/service-software.html
// [blue/green Auto-Tune enhancement]: https://docs.aws.amazon.com/opensearch-service/latest/developerguide/auto-tune.html#auto-tune-types
func opensearch_UpdateScheduledAction(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateScheduledActionInput{
		// ActionID: *string, // Required
		// ActionType: types.ActionType, // Required
		// DomainName: *string, // Required
		// ScheduleAt: types.ScheduleAt, // Required
	}

	if len(_opensearchActionID) > 0 {
		input.ActionID = aws.String(_opensearchActionID)
	}
	if len(_opensearchActionType) > 0 {
		if err := assignInputField(input, "ActionType", _opensearchActionType); err != nil {
			log.Errorf("invalid --action-type: %s", err.Error())
			return
		}
	}
	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchScheduleAt) > 0 {
		if err := assignInputField(input, "ScheduleAt", _opensearchScheduleAt); err != nil {
			log.Errorf("invalid --schedule-at: %s", err.Error())
			return
		}
	}
	if len(_opensearchDesiredStartTime) > 0 {
		if err := assignInputField(input, "DesiredStartTime", _opensearchDesiredStartTime); err != nil {
			log.Errorf("invalid --desired-start-time: %s", err.Error())
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

// Modifies an Amazon OpenSearch Service-managed interface VPC endpoint.
func opensearch_UpdateVpcEndpoint(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpdateVpcEndpointInput{
		// VpcEndpointId: *string, // Required
		// VpcOptions: *types.VPCOptions, // Required
	}

	if len(_opensearchVpcEndpointId) > 0 {
		input.VpcEndpointId = aws.String(_opensearchVpcEndpointId)
	}
	if len(_opensearchVpcOptions) > 0 {
		if err := assignInputField(input, "VpcOptions", _opensearchVpcOptions); err != nil {
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

// Allows you to either upgrade your Amazon OpenSearch Service domain or perform
// an upgrade eligibility check to a compatible version of OpenSearch or
// Elasticsearch.
func opensearch_UpgradeDomain(cfg aws.Config, client *opensearch.Client) {
	input := &opensearch.UpgradeDomainInput{
		// DomainName: *string, // Required
		// TargetVersion: *string, // Required
	}

	if len(_opensearchDomainName) > 0 {
		input.DomainName = aws.String(_opensearchDomainName)
	}
	if len(_opensearchTargetVersion) > 0 {
		input.TargetVersion = aws.String(_opensearchTargetVersion)
	}
	if len(_opensearchAdvancedOptions) > 0 {
		if err := assignInputField(input, "AdvancedOptions", _opensearchAdvancedOptions); err != nil {
			log.Errorf("invalid --advanced-options: %s", err.Error())
			return
		}
	}
	if len(_opensearchPerformCheckOnly) > 0 {
		if err := assignInputField(input, "PerformCheckOnly", _opensearchPerformCheckOnly); err != nil {
			log.Errorf("invalid --perform-check-only: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpgradeDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_opensearchCmd)
	_opensearchCmd.Flags().SortFlags = false

	_opensearchCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_opensearchCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_opensearchCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_opensearchCmd.Flags().StringVarP(&_opensearchAccessPolicies, "access-policies", "", "", "Access Policies")
	_opensearchCmd.Flags().StringVarP(&_opensearchAccount, "account", "", "", "Account")
	_opensearchCmd.Flags().StringVarP(&_opensearchAction, "action", "", "", "Action")
	_opensearchCmd.Flags().StringVarP(&_opensearchActionID, "action-id", "", "", "Action ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchActionType, "action-type", "", "", "Action Type")
	_opensearchCmd.Flags().StringVarP(&_opensearchAdvancedOptions, "advanced-options", "", "", "Advanced Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchAdvancedSecurityOptions, "advanced-security-options", "", "", "Advanced Security Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchAIMLOptions, "aiml-options", "", "", "Aiml Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchAppConfigs, "app-configs", "", "", "App Configs")
	_opensearchCmd.Flags().StringVarP(&_opensearchApplicationArn, "application-arn", "", "", "Application ARN")
	_opensearchCmd.Flags().StringVarP(&_opensearchARN, "arn", "", "", "ARN")
	_opensearchCmd.Flags().StringVarP(&_opensearchAssociationConfiguration, "association-configuration", "", "", "Association Configuration")
	_opensearchCmd.Flags().StringVarP(&_opensearchAutoTuneOptions, "auto-tune-options", "", "", "Auto Tune Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchChangeId, "change-id", "", "", "Change ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchClientToken, "client-token", "", "", "Client Token")
	_opensearchCmd.Flags().StringVarP(&_opensearchClusterConfig, "cluster-config", "", "", "Cluster Config")
	_opensearchCmd.Flags().StringVarP(&_opensearchCognitoOptions, "cognito-options", "", "", "Cognito Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchCommitMessage, "commit-message", "", "", "Commit Message")
	_opensearchCmd.Flags().StringVarP(&_opensearchConnectionAlias, "connection-alias", "", "", "Connection Alias")
	_opensearchCmd.Flags().StringVarP(&_opensearchConnectionId, "connection-id", "", "", "Connection ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchConnectionMode, "connection-mode", "", "", "Connection Mode")
	_opensearchCmd.Flags().StringVarP(&_opensearchConnectionProperties, "connection-properties", "", "", "Connection Properties")
	_opensearchCmd.Flags().StringVarP(&_opensearchDataSourceName, "data-source-name", "", "", "Data Source Name")
	_opensearchCmd.Flags().StringVarP(&_opensearchDataSourceType, "data-source-type", "", "", "Data Source Type")
	_opensearchCmd.Flags().StringVarP(&_opensearchDataSources, "data-sources", "", "", "Data Sources")
	_opensearchCmd.Flags().StringVarP(&_opensearchDescription, "description", "", "", "Description")
	_opensearchCmd.Flags().StringVarP(&_opensearchDesiredStartTime, "desired-start-time", "", "", "Desired Start Time")
	_opensearchCmd.Flags().StringVarP(&_opensearchDomainArn, "domain-arn", "", "", "Domain ARN")
	_opensearchCmd.Flags().StringVarP(&_opensearchDomainEndpointOptions, "domain-endpoint-options", "", "", "Domain Endpoint Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchDomainName, "domain-name", "", "", "Domain Name")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchDomainNames, "domain-names", "", nil, "Domain Names")
	_opensearchCmd.Flags().StringVarP(&_opensearchDryRun, "dry-run", "", "", "Dry Run")
	_opensearchCmd.Flags().StringVarP(&_opensearchDryRunId, "dry-run-id", "", "", "Dry Run ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchDryRunMode, "dry-run-mode", "", "", "Dry Run Mode")
	_opensearchCmd.Flags().StringVarP(&_opensearchEBSOptions, "ebs-options", "", "", "Ebs Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchEncryptionAtRestOptions, "encryption-at-rest-options", "", "", "Encryption At Rest Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchEngineType, "engine-type", "", "", "Engine Type")
	_opensearchCmd.Flags().StringVarP(&_opensearchEngineVersion, "engine-version", "", "", "Engine Version")
	_opensearchCmd.Flags().StringVarP(&_opensearchFilters, "filters", "", "", "Filters")
	_opensearchCmd.Flags().StringVarP(&_opensearchIamIdentityCenterOptions, "iam-identity-center-options", "", "", "IAM Identity Center Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchId, "id", "", "", "ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchIdentityCenterOptions, "identity-center-options", "", "", "Identity Center Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchIndexName, "index-name", "", "", "Index Name")
	_opensearchCmd.Flags().StringVarP(&_opensearchIndexSchema, "index-schema", "", "", "Index Schema")
	_opensearchCmd.Flags().StringVarP(&_opensearchInstanceCount, "instance-count", "", "", "Instance Count")
	_opensearchCmd.Flags().StringVarP(&_opensearchInstanceType, "instance-type", "", "", "Instance Type")
	_opensearchCmd.Flags().StringVarP(&_opensearchIPAddressType, "ip-address-type", "", "", "IP Address Type")
	_opensearchCmd.Flags().StringVarP(&_opensearchKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_opensearchCmd.Flags().StringVarP(&_opensearchLoadDryRunConfig, "load-dry-run-config", "", "", "Load Dry Run Config")
	_opensearchCmd.Flags().StringVarP(&_opensearchLocalDomainInfo, "local-domain-info", "", "", "Local Domain Info")
	_opensearchCmd.Flags().StringVarP(&_opensearchLogPublishingOptions, "log-publishing-options", "", "", "Log Publishing Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchMaintenanceId, "maintenance-id", "", "", "Maintenance ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchMaxResults, "max-results", "", "", "Max Results")
	_opensearchCmd.Flags().StringVarP(&_opensearchName, "name", "", "", "Name")
	_opensearchCmd.Flags().StringVarP(&_opensearchNextToken, "next-token", "", "", "Next Token")
	_opensearchCmd.Flags().StringVarP(&_opensearchNodeId, "node-id", "", "", "Node ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchNodeToNodeEncryptionOptions, "node-to-node-encryption-options", "", "", "Node To Node Encryption Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchOffPeakWindowOptions, "off-peak-window-options", "", "", "Off Peak Window Options")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchOpenSearchArns, "open-search-arns", "", nil, "Open Search Arns")
	_opensearchCmd.Flags().StringVarP(&_opensearchOperation, "operation", "", "", "Operation")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageConfiguration, "package-configuration", "", "", "Package Configuration")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageDescription, "package-description", "", "", "Package Description")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageEncryptionOptions, "package-encryption-options", "", "", "Package Encryption Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageID, "package-id", "", "", "Package ID")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchPackageList, "package-list", "", nil, "Package List")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageName, "package-name", "", "", "Package Name")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageSource, "package-source", "", "", "Package Source")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageType, "package-type", "", "", "Package Type")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchPackageUserList, "package-user-list", "", nil, "Package User List")
	_opensearchCmd.Flags().StringVarP(&_opensearchPackageVendingOptions, "package-vending-options", "", "", "Package Vending Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchPerformCheckOnly, "perform-check-only", "", "", "Perform Check Only")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchPrerequisitePackageIDList, "prerequisite-package-id-list", "", nil, "Prerequisite Package ID List")
	_opensearchCmd.Flags().StringVarP(&_opensearchRemoteDomainInfo, "remote-domain-info", "", "", "Remote Domain Info")
	_opensearchCmd.Flags().StringVarP(&_opensearchReservationName, "reservation-name", "", "", "Reservation Name")
	_opensearchCmd.Flags().StringVarP(&_opensearchReservedInstanceId, "reserved-instance-id", "", "", "Reserved Instance ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchReservedInstanceOfferingId, "reserved-instance-offering-id", "", "", "Reserved Instance Offering ID")
	_opensearchCmd.Flags().StringVarP(&_opensearchRetrieveAZs, "retrieve-azs", "", "", "Retrieve Azs")
	_opensearchCmd.Flags().StringVarP(&_opensearchScheduleAt, "schedule-at", "", "", "Schedule At")
	_opensearchCmd.Flags().StringVarP(&_opensearchService, "service", "", "", "Service")
	_opensearchCmd.Flags().StringVarP(&_opensearchSetAsDefault, "set-as-default", "", "", "Set As Default")
	_opensearchCmd.Flags().StringVarP(&_opensearchSnapshotOptions, "snapshot-options", "", "", "Snapshot Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchSoftwareUpdateOptions, "software-update-options", "", "", "Software Update Options")
	_opensearchCmd.Flags().StringVarP(&_opensearchStatus, "status", "", "", "Status")
	_opensearchCmd.Flags().StringVarP(&_opensearchStatuses, "statuses", "", "", "Statuses")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchTagKeys, "tag-keys", "", nil, "Tag Keys")
	_opensearchCmd.Flags().StringVarP(&_opensearchTagList, "tag-list", "", "", "Tag List")
	_opensearchCmd.Flags().StringVarP(&_opensearchTargetVersion, "target-version", "", "", "Target Version")
	_opensearchCmd.Flags().StringVarP(&_opensearchVpcEndpointId, "vpc-endpoint-id", "", "", "VPC Endpoint ID")
	_opensearchCmd.Flags().StringSliceVarP(&_opensearchVpcEndpointIds, "vpc-endpoint-ids", "", nil, "VPC Endpoint Ids")
	_opensearchCmd.Flags().StringVarP(&_opensearchVpcOptions, "vpc-options", "", "", "VPC Options")

	_opensearchCmd.Flags().BoolVarP(&_opensearchAcceptInboundConnection, "accept-inbound-connection", "", false, "Accept Inbound Connection")
	_opensearchCmd.Flags().BoolVarP(&_opensearchAddDataSource, "add-data-source", "", false, "Add Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchAddDirectQueryDataSource, "add-direct-query-data-source", "", false, "Add Direct Query Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchAddTags, "add-tags", "", false, "Add Tags")
	_opensearchCmd.Flags().BoolVarP(&_opensearchAssociatePackage, "associate-package", "", false, "Associate Package")
	_opensearchCmd.Flags().BoolVarP(&_opensearchAssociatePackages, "associate-packages", "", false, "Associate Packages")
	_opensearchCmd.Flags().BoolVarP(&_opensearchAuthorizeVpcEndpointAccess, "authorize-vpc-endpoint-access", "", false, "Authorize VPC Endpoint Access")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCancelDomainConfigChange, "cancel-domain-config-change", "", false, "Cancel Domain Config Change")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCancelServiceSoftwareUpdate, "cancel-service-software-update", "", false, "Cancel Service Software Update")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCreateApplication, "create-application", "", false, "Create Application")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCreateDomain, "create-domain", "", false, "Create Domain")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCreateIndex, "create-index", "", false, "Create Index")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCreateOutboundConnection, "create-outbound-connection", "", false, "Create Outbound Connection")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCreatePackage, "create-package", "", false, "Create Package")
	_opensearchCmd.Flags().BoolVarP(&_opensearchCreateVpcEndpoint, "create-vpc-endpoint", "", false, "Create VPC Endpoint")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteApplication, "delete-application", "", false, "Delete Application")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteDataSource, "delete-data-source", "", false, "Delete Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteDirectQueryDataSource, "delete-direct-query-data-source", "", false, "Delete Direct Query Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteDomain, "delete-domain", "", false, "Delete Domain")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteInboundConnection, "delete-inbound-connection", "", false, "Delete Inbound Connection")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteIndex, "delete-index", "", false, "Delete Index")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteOutboundConnection, "delete-outbound-connection", "", false, "Delete Outbound Connection")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeletePackage, "delete-package", "", false, "Delete Package")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDeleteVpcEndpoint, "delete-vpc-endpoint", "", false, "Delete VPC Endpoint")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomain, "describe-domain", "", false, "Describe Domain")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomainAutoTunes, "describe-domain-auto-tunes", "", false, "Describe Domain Auto Tunes")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomainChangeProgress, "describe-domain-change-progress", "", false, "Describe Domain Change Progress")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomainConfig, "describe-domain-config", "", false, "Describe Domain Config")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomainHealth, "describe-domain-health", "", false, "Describe Domain Health")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomainNodes, "describe-domain-nodes", "", false, "Describe Domain Nodes")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDomains, "describe-domains", "", false, "Describe Domains")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeDryRunProgress, "describe-dry-run-progress", "", false, "Describe Dry Run Progress")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeInboundConnections, "describe-inbound-connections", "", false, "Describe Inbound Connections")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeInstanceTypeLimits, "describe-instance-type-limits", "", false, "Describe Instance Type Limits")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeOutboundConnections, "describe-outbound-connections", "", false, "Describe Outbound Connections")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribePackages, "describe-packages", "", false, "Describe Packages")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeReservedInstanceOfferings, "describe-reserved-instance-offerings", "", false, "Describe Reserved Instance Offerings")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeReservedInstances, "describe-reserved-instances", "", false, "Describe Reserved Instances")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDescribeVpcEndpoints, "describe-vpc-endpoints", "", false, "Describe VPC Endpoints")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDissociatePackage, "dissociate-package", "", false, "Dissociate Package")
	_opensearchCmd.Flags().BoolVarP(&_opensearchDissociatePackages, "dissociate-packages", "", false, "Dissociate Packages")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetApplication, "get-application", "", false, "Get Application")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetCompatibleVersions, "get-compatible-versions", "", false, "Get Compatible Versions")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetDataSource, "get-data-source", "", false, "Get Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetDefaultApplicationSetting, "get-default-application-setting", "", false, "Get Default Application Setting")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetDirectQueryDataSource, "get-direct-query-data-source", "", false, "Get Direct Query Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetDomainMaintenanceStatus, "get-domain-maintenance-status", "", false, "Get Domain Maintenance Status")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetIndex, "get-index", "", false, "Get Index")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetPackageVersionHistory, "get-package-version-history", "", false, "Get Package Version History")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetUpgradeHistory, "get-upgrade-history", "", false, "Get Upgrade History")
	_opensearchCmd.Flags().BoolVarP(&_opensearchGetUpgradeStatus, "get-upgrade-status", "", false, "Get Upgrade Status")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListApplications, "list-applications", "", false, "List Applications")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListDataSources, "list-data-sources", "", false, "List Data Sources")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListDirectQueryDataSources, "list-direct-query-data-sources", "", false, "List Direct Query Data Sources")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListDomainMaintenances, "list-domain-maintenances", "", false, "List Domain Maintenances")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListDomainNames, "list-domain-names", "", false, "List Domain Names")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListDomainsForPackage, "list-domains-for-package", "", false, "List Domains For Package")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListInstanceTypeDetails, "list-instance-type-details", "", false, "List Instance Type Details")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListPackagesForDomain, "list-packages-for-domain", "", false, "List Packages For Domain")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListScheduledActions, "list-scheduled-actions", "", false, "List Scheduled Actions")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListTags, "list-tags", "", false, "List Tags")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListVersions, "list-versions", "", false, "List Versions")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListVpcEndpointAccess, "list-vpc-endpoint-access", "", false, "List VPC Endpoint Access")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListVpcEndpoints, "list-vpc-endpoints", "", false, "List VPC Endpoints")
	_opensearchCmd.Flags().BoolVarP(&_opensearchListVpcEndpointsForDomain, "list-vpc-endpoints-for-domain", "", false, "List VPC Endpoints For Domain")
	_opensearchCmd.Flags().BoolVarP(&_opensearchPurchaseReservedInstanceOffering, "purchase-reserved-instance-offering", "", false, "Purchase Reserved Instance Offering")
	_opensearchCmd.Flags().BoolVarP(&_opensearchPutDefaultApplicationSetting, "put-default-application-setting", "", false, "Put Default Application Setting")
	_opensearchCmd.Flags().BoolVarP(&_opensearchRejectInboundConnection, "reject-inbound-connection", "", false, "Reject Inbound Connection")
	_opensearchCmd.Flags().BoolVarP(&_opensearchRemoveTags, "remove-tags", "", false, "Remove Tags")
	_opensearchCmd.Flags().BoolVarP(&_opensearchRevokeVpcEndpointAccess, "revoke-vpc-endpoint-access", "", false, "Revoke VPC Endpoint Access")
	_opensearchCmd.Flags().BoolVarP(&_opensearchStartDomainMaintenance, "start-domain-maintenance", "", false, "Start Domain Maintenance")
	_opensearchCmd.Flags().BoolVarP(&_opensearchStartServiceSoftwareUpdate, "start-service-software-update", "", false, "Start Service Software Update")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateApplication, "update-application", "", false, "Update Application")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateDataSource, "update-data-source", "", false, "Update Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateDirectQueryDataSource, "update-direct-query-data-source", "", false, "Update Direct Query Data Source")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateDomainConfig, "update-domain-config", "", false, "Update Domain Config")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateIndex, "update-index", "", false, "Update Index")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdatePackage, "update-package", "", false, "Update Package")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdatePackageScope, "update-package-scope", "", false, "Update Package Scope")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateScheduledAction, "update-scheduled-action", "", false, "Update Scheduled Action")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpdateVpcEndpoint, "update-vpc-endpoint", "", false, "Update VPC Endpoint")
	_opensearchCmd.Flags().BoolVarP(&_opensearchUpgradeDomain, "upgrade-domain", "", false, "Upgrade Domain")

}
