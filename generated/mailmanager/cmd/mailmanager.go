package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mailmanager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mailmanagerCmd represents the mailmanager command
var _mailmanagerCmd = &cobra.Command{
	Use:   "mailmanager",
	Short: "AWS mailmanager CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mailmanager.NewFromConfig(cfg)
		if _mailmanagerCreateAddonInstance {
			mailmanager_CreateAddonInstance(cfg, client)
			return
		}
		if _mailmanagerCreateAddonSubscription {
			mailmanager_CreateAddonSubscription(cfg, client)
			return
		}
		if _mailmanagerCreateAddressList {
			mailmanager_CreateAddressList(cfg, client)
			return
		}
		if _mailmanagerCreateAddressListImportJob {
			mailmanager_CreateAddressListImportJob(cfg, client)
			return
		}
		if _mailmanagerCreateArchive {
			mailmanager_CreateArchive(cfg, client)
			return
		}
		if _mailmanagerCreateIngressPoint {
			mailmanager_CreateIngressPoint(cfg, client)
			return
		}
		if _mailmanagerCreateRelay {
			mailmanager_CreateRelay(cfg, client)
			return
		}
		if _mailmanagerCreateRuleSet {
			mailmanager_CreateRuleSet(cfg, client)
			return
		}
		if _mailmanagerCreateTrafficPolicy {
			mailmanager_CreateTrafficPolicy(cfg, client)
			return
		}
		if _mailmanagerDeleteAddonInstance {
			mailmanager_DeleteAddonInstance(cfg, client)
			return
		}
		if _mailmanagerDeleteAddonSubscription {
			mailmanager_DeleteAddonSubscription(cfg, client)
			return
		}
		if _mailmanagerDeleteAddressList {
			mailmanager_DeleteAddressList(cfg, client)
			return
		}
		if _mailmanagerDeleteArchive {
			mailmanager_DeleteArchive(cfg, client)
			return
		}
		if _mailmanagerDeleteIngressPoint {
			mailmanager_DeleteIngressPoint(cfg, client)
			return
		}
		if _mailmanagerDeleteRelay {
			mailmanager_DeleteRelay(cfg, client)
			return
		}
		if _mailmanagerDeleteRuleSet {
			mailmanager_DeleteRuleSet(cfg, client)
			return
		}
		if _mailmanagerDeleteTrafficPolicy {
			mailmanager_DeleteTrafficPolicy(cfg, client)
			return
		}
		if _mailmanagerDeregisterMemberFromAddressList {
			mailmanager_DeregisterMemberFromAddressList(cfg, client)
			return
		}
		if _mailmanagerGetAddonInstance {
			mailmanager_GetAddonInstance(cfg, client)
			return
		}
		if _mailmanagerGetAddonSubscription {
			mailmanager_GetAddonSubscription(cfg, client)
			return
		}
		if _mailmanagerGetAddressList {
			mailmanager_GetAddressList(cfg, client)
			return
		}
		if _mailmanagerGetAddressListImportJob {
			mailmanager_GetAddressListImportJob(cfg, client)
			return
		}
		if _mailmanagerGetArchive {
			mailmanager_GetArchive(cfg, client)
			return
		}
		if _mailmanagerGetArchiveExport {
			mailmanager_GetArchiveExport(cfg, client)
			return
		}
		if _mailmanagerGetArchiveMessage {
			mailmanager_GetArchiveMessage(cfg, client)
			return
		}
		if _mailmanagerGetArchiveMessageContent {
			mailmanager_GetArchiveMessageContent(cfg, client)
			return
		}
		if _mailmanagerGetArchiveSearch {
			mailmanager_GetArchiveSearch(cfg, client)
			return
		}
		if _mailmanagerGetArchiveSearchResults {
			mailmanager_GetArchiveSearchResults(cfg, client)
			return
		}
		if _mailmanagerGetIngressPoint {
			mailmanager_GetIngressPoint(cfg, client)
			return
		}
		if _mailmanagerGetMemberOfAddressList {
			mailmanager_GetMemberOfAddressList(cfg, client)
			return
		}
		if _mailmanagerGetRelay {
			mailmanager_GetRelay(cfg, client)
			return
		}
		if _mailmanagerGetRuleSet {
			mailmanager_GetRuleSet(cfg, client)
			return
		}
		if _mailmanagerGetTrafficPolicy {
			mailmanager_GetTrafficPolicy(cfg, client)
			return
		}
		if _mailmanagerListAddonInstances {
			mailmanager_ListAddonInstances(cfg, client)
			return
		}
		if _mailmanagerListAddonSubscriptions {
			mailmanager_ListAddonSubscriptions(cfg, client)
			return
		}
		if _mailmanagerListAddressListImportJobs {
			mailmanager_ListAddressListImportJobs(cfg, client)
			return
		}
		if _mailmanagerListAddressLists {
			mailmanager_ListAddressLists(cfg, client)
			return
		}
		if _mailmanagerListArchiveExports {
			mailmanager_ListArchiveExports(cfg, client)
			return
		}
		if _mailmanagerListArchiveSearches {
			mailmanager_ListArchiveSearches(cfg, client)
			return
		}
		if _mailmanagerListArchives {
			mailmanager_ListArchives(cfg, client)
			return
		}
		if _mailmanagerListIngressPoints {
			mailmanager_ListIngressPoints(cfg, client)
			return
		}
		if _mailmanagerListMembersOfAddressList {
			mailmanager_ListMembersOfAddressList(cfg, client)
			return
		}
		if _mailmanagerListRelays {
			mailmanager_ListRelays(cfg, client)
			return
		}
		if _mailmanagerListRuleSets {
			mailmanager_ListRuleSets(cfg, client)
			return
		}
		if _mailmanagerListTagsForResource {
			mailmanager_ListTagsForResource(cfg, client)
			return
		}
		if _mailmanagerListTrafficPolicies {
			mailmanager_ListTrafficPolicies(cfg, client)
			return
		}
		if _mailmanagerRegisterMemberToAddressList {
			mailmanager_RegisterMemberToAddressList(cfg, client)
			return
		}
		if _mailmanagerStartAddressListImportJob {
			mailmanager_StartAddressListImportJob(cfg, client)
			return
		}
		if _mailmanagerStartArchiveExport {
			mailmanager_StartArchiveExport(cfg, client)
			return
		}
		if _mailmanagerStartArchiveSearch {
			mailmanager_StartArchiveSearch(cfg, client)
			return
		}
		if _mailmanagerStopAddressListImportJob {
			mailmanager_StopAddressListImportJob(cfg, client)
			return
		}
		if _mailmanagerStopArchiveExport {
			mailmanager_StopArchiveExport(cfg, client)
			return
		}
		if _mailmanagerStopArchiveSearch {
			mailmanager_StopArchiveSearch(cfg, client)
			return
		}
		if _mailmanagerTagResource {
			mailmanager_TagResource(cfg, client)
			return
		}
		if _mailmanagerUntagResource {
			mailmanager_UntagResource(cfg, client)
			return
		}
		if _mailmanagerUpdateArchive {
			mailmanager_UpdateArchive(cfg, client)
			return
		}
		if _mailmanagerUpdateIngressPoint {
			mailmanager_UpdateIngressPoint(cfg, client)
			return
		}
		if _mailmanagerUpdateRelay {
			mailmanager_UpdateRelay(cfg, client)
			return
		}
		if _mailmanagerUpdateRuleSet {
			mailmanager_UpdateRuleSet(cfg, client)
			return
		}
		if _mailmanagerUpdateTrafficPolicy {
			mailmanager_UpdateTrafficPolicy(cfg, client)
			return
		}

	},
}

var (
	_mailmanagerCreateAddonInstance             bool
	_mailmanagerCreateAddonSubscription         bool
	_mailmanagerCreateAddressList               bool
	_mailmanagerCreateAddressListImportJob      bool
	_mailmanagerCreateArchive                   bool
	_mailmanagerCreateIngressPoint              bool
	_mailmanagerCreateRelay                     bool
	_mailmanagerCreateRuleSet                   bool
	_mailmanagerCreateTrafficPolicy             bool
	_mailmanagerDeleteAddonInstance             bool
	_mailmanagerDeleteAddonSubscription         bool
	_mailmanagerDeleteAddressList               bool
	_mailmanagerDeleteArchive                   bool
	_mailmanagerDeleteIngressPoint              bool
	_mailmanagerDeleteRelay                     bool
	_mailmanagerDeleteRuleSet                   bool
	_mailmanagerDeleteTrafficPolicy             bool
	_mailmanagerDeregisterMemberFromAddressList bool
	_mailmanagerGetAddonInstance                bool
	_mailmanagerGetAddonSubscription            bool
	_mailmanagerGetAddressList                  bool
	_mailmanagerGetAddressListImportJob         bool
	_mailmanagerGetArchive                      bool
	_mailmanagerGetArchiveExport                bool
	_mailmanagerGetArchiveMessage               bool
	_mailmanagerGetArchiveMessageContent        bool
	_mailmanagerGetArchiveSearch                bool
	_mailmanagerGetArchiveSearchResults         bool
	_mailmanagerGetIngressPoint                 bool
	_mailmanagerGetMemberOfAddressList          bool
	_mailmanagerGetRelay                        bool
	_mailmanagerGetRuleSet                      bool
	_mailmanagerGetTrafficPolicy                bool
	_mailmanagerListAddonInstances              bool
	_mailmanagerListAddonSubscriptions          bool
	_mailmanagerListAddressListImportJobs       bool
	_mailmanagerListAddressLists                bool
	_mailmanagerListArchiveExports              bool
	_mailmanagerListArchiveSearches             bool
	_mailmanagerListArchives                    bool
	_mailmanagerListIngressPoints               bool
	_mailmanagerListMembersOfAddressList        bool
	_mailmanagerListRelays                      bool
	_mailmanagerListRuleSets                    bool
	_mailmanagerListTagsForResource             bool
	_mailmanagerListTrafficPolicies             bool
	_mailmanagerRegisterMemberToAddressList     bool
	_mailmanagerStartAddressListImportJob       bool
	_mailmanagerStartArchiveExport              bool
	_mailmanagerStartArchiveSearch              bool
	_mailmanagerStopAddressListImportJob        bool
	_mailmanagerStopArchiveExport               bool
	_mailmanagerStopArchiveSearch               bool
	_mailmanagerTagResource                     bool
	_mailmanagerUntagResource                   bool
	_mailmanagerUpdateArchive                   bool
	_mailmanagerUpdateIngressPoint              bool
	_mailmanagerUpdateRelay                     bool
	_mailmanagerUpdateRuleSet                   bool
	_mailmanagerUpdateTrafficPolicy             bool

	_mailmanagerAddonInstanceId                string
	_mailmanagerAddonName                      string
	_mailmanagerAddonSubscriptionId            string
	_mailmanagerAddress                        string
	_mailmanagerAddressListId                  string
	_mailmanagerAddressListName                string
	_mailmanagerArchiveId                      string
	_mailmanagerArchiveName                    string
	_mailmanagerArchivedMessageId              string
	_mailmanagerAuthentication                 string
	_mailmanagerClientToken                    string
	_mailmanagerDefaultAction                  string
	_mailmanagerExportDestinationConfiguration string
	_mailmanagerExportId                       string
	_mailmanagerFilter                         string
	_mailmanagerFilters                        string
	_mailmanagerFromTimestamp                  string
	_mailmanagerImportDataFormat               string
	_mailmanagerIncludeMetadata                string
	_mailmanagerIngressPointConfiguration      string
	_mailmanagerIngressPointId                 string
	_mailmanagerIngressPointName               string
	_mailmanagerJobId                          string
	_mailmanagerKmsKeyArn                      string
	_mailmanagerMaxMessageSizeBytes            string
	_mailmanagerMaxResults                     string
	_mailmanagerName                           string
	_mailmanagerNetworkConfiguration           string
	_mailmanagerNextToken                      string
	_mailmanagerPageSize                       string
	_mailmanagerPolicyStatements               string
	_mailmanagerRelayId                        string
	_mailmanagerRelayName                      string
	_mailmanagerResourceArn                    string
	_mailmanagerRetention                      string
	_mailmanagerRuleSetId                      string
	_mailmanagerRuleSetName                    string
	_mailmanagerRules                          string
	_mailmanagerSearchId                       string
	_mailmanagerServerName                     string
	_mailmanagerServerPort                     string
	_mailmanagerStatusToUpdate                 string
	_mailmanagerTagKeys                        []string
	_mailmanagerTags                           string
	_mailmanagerToTimestamp                    string
	_mailmanagerTrafficPolicyId                string
	_mailmanagerTrafficPolicyName              string
	_mailmanagerType                           string
)

// Creates an Add On instance for the subscription indicated in the request. The
// resulting Amazon Resource Name (ARN) can be used in a conditional statement for
// a rule set or traffic policy.
func mailmanager_CreateAddonInstance(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateAddonInstanceInput{
		// AddonSubscriptionId: *string, // Required
	}

	if len(_mailmanagerAddonSubscriptionId) > 0 {
		input.AddonSubscriptionId = aws.String(_mailmanagerAddonSubscriptionId)
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAddonInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a subscription for an Add On representing the acceptance of its terms
// of use and additional pricing. The subscription can then be used to create an
// instance for use in rule sets or traffic policies.
func mailmanager_CreateAddonSubscription(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateAddonSubscriptionInput{
		// AddonName: *string, // Required
	}

	if len(_mailmanagerAddonName) > 0 {
		input.AddonName = aws.String(_mailmanagerAddonName)
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAddonSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new address list.
func mailmanager_CreateAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateAddressListInput{
		// AddressListName: *string, // Required
	}

	if len(_mailmanagerAddressListName) > 0 {
		input.AddressListName = aws.String(_mailmanagerAddressListName)
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAddressList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an import job for an address list.
func mailmanager_CreateAddressListImportJob(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateAddressListImportJobInput{
		// AddressListId: *string, // Required
		// ImportDataFormat: *types.ImportDataFormat, // Required
		// Name: *string, // Required
	}

	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}
	if len(_mailmanagerImportDataFormat) > 0 {
		if err := assignInputField(input, "ImportDataFormat", _mailmanagerImportDataFormat); err != nil {
			log.Errorf("invalid --import-data-format: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerName) > 0 {
		input.Name = aws.String(_mailmanagerName)
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}

	if resp, err := client.CreateAddressListImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new email archive resource for storing and retaining emails.
func mailmanager_CreateArchive(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateArchiveInput{
		// ArchiveName: *string, // Required
	}

	if len(_mailmanagerArchiveName) > 0 {
		input.ArchiveName = aws.String(_mailmanagerArchiveName)
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerKmsKeyArn) > 0 {
		input.KmsKeyArn = aws.String(_mailmanagerKmsKeyArn)
	}
	if len(_mailmanagerRetention) > 0 {
		if err := assignInputField(input, "Retention", _mailmanagerRetention); err != nil {
			log.Errorf("invalid --retention: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provision a new ingress endpoint resource.
func mailmanager_CreateIngressPoint(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateIngressPointInput{
		// IngressPointName: *string, // Required
		// RuleSetId: *string, // Required
		// TrafficPolicyId: *string, // Required
		// Type: types.IngressPointType, // Required
	}

	if len(_mailmanagerIngressPointName) > 0 {
		input.IngressPointName = aws.String(_mailmanagerIngressPointName)
	}
	if len(_mailmanagerRuleSetId) > 0 {
		input.RuleSetId = aws.String(_mailmanagerRuleSetId)
	}
	if len(_mailmanagerTrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_mailmanagerTrafficPolicyId)
	}
	if len(_mailmanagerType) > 0 {
		if err := assignInputField(input, "Type", _mailmanagerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerIngressPointConfiguration) > 0 {
		if err := assignInputField(input, "IngressPointConfiguration", _mailmanagerIngressPointConfiguration); err != nil {
			log.Errorf("invalid --ingress-point-configuration: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerNetworkConfiguration) > 0 {
		if err := assignInputField(input, "NetworkConfiguration", _mailmanagerNetworkConfiguration); err != nil {
			log.Errorf("invalid --network-configuration: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIngressPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a relay resource which can be used in rules to relay incoming emails to
// defined relay destinations.
func mailmanager_CreateRelay(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateRelayInput{
		// Authentication: types.RelayAuthentication, // Required
		// RelayName: *string, // Required
		// ServerName: *string, // Required
		// ServerPort: *int32, // Required
	}

	if len(_mailmanagerAuthentication) > 0 {
		if err := assignInputField(input, "Authentication", _mailmanagerAuthentication); err != nil {
			log.Errorf("invalid --authentication: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerRelayName) > 0 {
		input.RelayName = aws.String(_mailmanagerRelayName)
	}
	if len(_mailmanagerServerName) > 0 {
		input.ServerName = aws.String(_mailmanagerServerName)
	}
	if len(_mailmanagerServerPort) > 0 {
		if err := assignInputField(input, "ServerPort", _mailmanagerServerPort); err != nil {
			log.Errorf("invalid --server-port: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRelay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provision a new rule set.
func mailmanager_CreateRuleSet(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateRuleSetInput{
		// RuleSetName: *string, // Required
		// Rules: []types.Rule, // Required
	}

	if len(_mailmanagerRuleSetName) > 0 {
		input.RuleSetName = aws.String(_mailmanagerRuleSetName)
	}
	if len(_mailmanagerRules) > 0 {
		if err := assignInputField(input, "Rules", _mailmanagerRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provision a new traffic policy resource.
func mailmanager_CreateTrafficPolicy(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.CreateTrafficPolicyInput{
		// DefaultAction: types.AcceptAction, // Required
		// PolicyStatements: []types.PolicyStatement, // Required
		// TrafficPolicyName: *string, // Required
	}

	if len(_mailmanagerDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _mailmanagerDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerPolicyStatements) > 0 {
		if err := assignInputField(input, "PolicyStatements", _mailmanagerPolicyStatements); err != nil {
			log.Errorf("invalid --policy-statements: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerTrafficPolicyName) > 0 {
		input.TrafficPolicyName = aws.String(_mailmanagerTrafficPolicyName)
	}
	if len(_mailmanagerClientToken) > 0 {
		input.ClientToken = aws.String(_mailmanagerClientToken)
	}
	if len(_mailmanagerMaxMessageSizeBytes) > 0 {
		if err := assignInputField(input, "MaxMessageSizeBytes", _mailmanagerMaxMessageSizeBytes); err != nil {
			log.Errorf("invalid --max-message-size-bytes: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Add On instance.
func mailmanager_DeleteAddonInstance(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteAddonInstanceInput{
		// AddonInstanceId: *string, // Required
	}

	if len(_mailmanagerAddonInstanceId) > 0 {
		input.AddonInstanceId = aws.String(_mailmanagerAddonInstanceId)
	}

	if resp, err := client.DeleteAddonInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an Add On subscription.
func mailmanager_DeleteAddonSubscription(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteAddonSubscriptionInput{
		// AddonSubscriptionId: *string, // Required
	}

	if len(_mailmanagerAddonSubscriptionId) > 0 {
		input.AddonSubscriptionId = aws.String(_mailmanagerAddonSubscriptionId)
	}

	if resp, err := client.DeleteAddonSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an address list.
func mailmanager_DeleteAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteAddressListInput{
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}

	if resp, err := client.DeleteAddressList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates deletion of an email archive. This changes the archive state to
// pending deletion. In this state, no new emails can be added, and existing
// archived emails become inaccessible (search, export, download). The archive and
// all of its contents will be permanently deleted 30 days after entering the
// pending deletion state, regardless of the configured retention period.
func mailmanager_DeleteArchive(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteArchiveInput{
		// ArchiveId: *string, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}

	if resp, err := client.DeleteArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete an ingress endpoint resource.
func mailmanager_DeleteIngressPoint(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteIngressPointInput{
		// IngressPointId: *string, // Required
	}

	if len(_mailmanagerIngressPointId) > 0 {
		input.IngressPointId = aws.String(_mailmanagerIngressPointId)
	}

	if resp, err := client.DeleteIngressPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing relay resource.
func mailmanager_DeleteRelay(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteRelayInput{
		// RelayId: *string, // Required
	}

	if len(_mailmanagerRelayId) > 0 {
		input.RelayId = aws.String(_mailmanagerRelayId)
	}

	if resp, err := client.DeleteRelay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a rule set.
func mailmanager_DeleteRuleSet(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteRuleSetInput{
		// RuleSetId: *string, // Required
	}

	if len(_mailmanagerRuleSetId) > 0 {
		input.RuleSetId = aws.String(_mailmanagerRuleSetId)
	}

	if resp, err := client.DeleteRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a traffic policy resource.
func mailmanager_DeleteTrafficPolicy(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeleteTrafficPolicyInput{
		// TrafficPolicyId: *string, // Required
	}

	if len(_mailmanagerTrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_mailmanagerTrafficPolicyId)
	}

	if resp, err := client.DeleteTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a member from an address list.
func mailmanager_DeregisterMemberFromAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.DeregisterMemberFromAddressListInput{
		// Address: *string, // Required
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddress) > 0 {
		input.Address = aws.String(_mailmanagerAddress)
	}
	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}

	if resp, err := client.DeregisterMemberFromAddressList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about an Add On instance.
func mailmanager_GetAddonInstance(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetAddonInstanceInput{
		// AddonInstanceId: *string, // Required
	}

	if len(_mailmanagerAddonInstanceId) > 0 {
		input.AddonInstanceId = aws.String(_mailmanagerAddonInstanceId)
	}

	if resp, err := client.GetAddonInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about an Add On subscription.
func mailmanager_GetAddonSubscription(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetAddonSubscriptionInput{
		// AddonSubscriptionId: *string, // Required
	}

	if len(_mailmanagerAddonSubscriptionId) > 0 {
		input.AddonSubscriptionId = aws.String(_mailmanagerAddonSubscriptionId)
	}

	if resp, err := client.GetAddonSubscription(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch attributes of an address list.
func mailmanager_GetAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetAddressListInput{
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}

	if resp, err := client.GetAddressList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch attributes of an import job.
func mailmanager_GetAddressListImportJob(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetAddressListImportJobInput{
		// JobId: *string, // Required
	}

	if len(_mailmanagerJobId) > 0 {
		input.JobId = aws.String(_mailmanagerJobId)
	}

	if resp, err := client.GetAddressListImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the full details and current state of a specified email archive.
func mailmanager_GetArchive(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetArchiveInput{
		// ArchiveId: *string, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}

	if resp, err := client.GetArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details and current status of a specific email archive export job.
func mailmanager_GetArchiveExport(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetArchiveExportInput{
		// ExportId: *string, // Required
	}

	if len(_mailmanagerExportId) > 0 {
		input.ExportId = aws.String(_mailmanagerExportId)
	}

	if resp, err := client.GetArchiveExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a pre-signed URL that provides temporary download access to the
// specific email message stored in the archive.
func mailmanager_GetArchiveMessage(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetArchiveMessageInput{
		// ArchivedMessageId: *string, // Required
	}

	if len(_mailmanagerArchivedMessageId) > 0 {
		input.ArchivedMessageId = aws.String(_mailmanagerArchivedMessageId)
	}

	if resp, err := client.GetArchiveMessage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the textual content of a specific email message stored in the archive.
// Attachments are not included.
func mailmanager_GetArchiveMessageContent(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetArchiveMessageContentInput{
		// ArchivedMessageId: *string, // Required
	}

	if len(_mailmanagerArchivedMessageId) > 0 {
		input.ArchivedMessageId = aws.String(_mailmanagerArchivedMessageId)
	}

	if resp, err := client.GetArchiveMessageContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details and current status of a specific email archive search job.
func mailmanager_GetArchiveSearch(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetArchiveSearchInput{
		// SearchId: *string, // Required
	}

	if len(_mailmanagerSearchId) > 0 {
		input.SearchId = aws.String(_mailmanagerSearchId)
	}

	if resp, err := client.GetArchiveSearch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the results of a completed email archive search job.
func mailmanager_GetArchiveSearchResults(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetArchiveSearchResultsInput{
		// SearchId: *string, // Required
	}

	if len(_mailmanagerSearchId) > 0 {
		input.SearchId = aws.String(_mailmanagerSearchId)
	}

	if resp, err := client.GetArchiveSearchResults(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch ingress endpoint resource attributes.
func mailmanager_GetIngressPoint(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetIngressPointInput{
		// IngressPointId: *string, // Required
	}

	if len(_mailmanagerIngressPointId) > 0 {
		input.IngressPointId = aws.String(_mailmanagerIngressPointId)
	}

	if resp, err := client.GetIngressPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch attributes of a member in an address list.
func mailmanager_GetMemberOfAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetMemberOfAddressListInput{
		// Address: *string, // Required
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddress) > 0 {
		input.Address = aws.String(_mailmanagerAddress)
	}
	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}

	if resp, err := client.GetMemberOfAddressList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch the relay resource and it's attributes.
func mailmanager_GetRelay(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetRelayInput{
		// RelayId: *string, // Required
	}

	if len(_mailmanagerRelayId) > 0 {
		input.RelayId = aws.String(_mailmanagerRelayId)
	}

	if resp, err := client.GetRelay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch attributes of a rule set.
func mailmanager_GetRuleSet(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetRuleSetInput{
		// RuleSetId: *string, // Required
	}

	if len(_mailmanagerRuleSetId) > 0 {
		input.RuleSetId = aws.String(_mailmanagerRuleSetId)
	}

	if resp, err := client.GetRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Fetch attributes of a traffic policy resource.
func mailmanager_GetTrafficPolicy(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.GetTrafficPolicyInput{
		// TrafficPolicyId: *string, // Required
	}

	if len(_mailmanagerTrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_mailmanagerTrafficPolicyId)
	}

	if resp, err := client.GetTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all Add On instances in your account.
func mailmanager_ListAddonInstances(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListAddonInstancesInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAddonInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListAddonInstancesOutput
	p := mailmanager.NewListAddonInstancesPaginator(client, input)
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

// Lists all Add On subscriptions in your account.
func mailmanager_ListAddonSubscriptions(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListAddonSubscriptionsInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAddonSubscriptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListAddonSubscriptionsOutput
	p := mailmanager.NewListAddonSubscriptionsPaginator(client, input)
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

// Lists jobs for an address list.
func mailmanager_ListAddressListImportJobs(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListAddressListImportJobsInput{
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}
	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAddressListImportJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListAddressListImportJobsOutput
	p := mailmanager.NewListAddressListImportJobsPaginator(client, input)
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

// Lists address lists for this account.
func mailmanager_ListAddressLists(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListAddressListsInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAddressLists(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListAddressListsOutput
	p := mailmanager.NewListAddressListsPaginator(client, input)
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

// Returns a list of email archive export jobs.
func mailmanager_ListArchiveExports(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListArchiveExportsInput{
		// ArchiveId: *string, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}
	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListArchiveExports(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListArchiveExportsOutput
	p := mailmanager.NewListArchiveExportsPaginator(client, input)
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

// Returns a list of email archive search jobs.
func mailmanager_ListArchiveSearches(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListArchiveSearchesInput{
		// ArchiveId: *string, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}
	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListArchiveSearches(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListArchiveSearchesOutput
	p := mailmanager.NewListArchiveSearchesPaginator(client, input)
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

// Returns a list of all email archives in your account.
func mailmanager_ListArchives(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListArchivesInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListArchives(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListArchivesOutput
	p := mailmanager.NewListArchivesPaginator(client, input)
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

// List all ingress endpoint resources.
func mailmanager_ListIngressPoints(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListIngressPointsInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListIngressPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListIngressPointsOutput
	p := mailmanager.NewListIngressPointsPaginator(client, input)
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

// Lists members of an address list.
func mailmanager_ListMembersOfAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListMembersOfAddressListInput{
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}
	if len(_mailmanagerFilter) > 0 {
		if err := assignInputField(input, "Filter", _mailmanagerFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListMembersOfAddressList(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListMembersOfAddressListOutput
	p := mailmanager.NewListMembersOfAddressListPaginator(client, input)
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

// Lists all the existing relay resources.
func mailmanager_ListRelays(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListRelaysInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRelays(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListRelaysOutput
	p := mailmanager.NewListRelaysPaginator(client, input)
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

// List rule sets for this account.
func mailmanager_ListRuleSets(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListRuleSetsInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRuleSets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListRuleSetsOutput
	p := mailmanager.NewListRuleSetsPaginator(client, input)
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

// Retrieves the list of tags (keys and values) assigned to the resource.
func mailmanager_ListTagsForResource(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mailmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_mailmanagerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// List traffic policy resources.
func mailmanager_ListTrafficPolicies(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.ListTrafficPoliciesInput{}

	if len(_mailmanagerNextToken) > 0 {
		input.NextToken = aws.String(_mailmanagerNextToken)
	}
	if len(_mailmanagerPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _mailmanagerPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListTrafficPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mailmanager.ListTrafficPoliciesOutput
	p := mailmanager.NewListTrafficPoliciesPaginator(client, input)
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

// Adds a member to an address list.
func mailmanager_RegisterMemberToAddressList(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.RegisterMemberToAddressListInput{
		// Address: *string, // Required
		// AddressListId: *string, // Required
	}

	if len(_mailmanagerAddress) > 0 {
		input.Address = aws.String(_mailmanagerAddress)
	}
	if len(_mailmanagerAddressListId) > 0 {
		input.AddressListId = aws.String(_mailmanagerAddressListId)
	}

	if resp, err := client.RegisterMemberToAddressList(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts an import job for an address list.
func mailmanager_StartAddressListImportJob(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.StartAddressListImportJobInput{
		// JobId: *string, // Required
	}

	if len(_mailmanagerJobId) > 0 {
		input.JobId = aws.String(_mailmanagerJobId)
	}

	if resp, err := client.StartAddressListImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates an export of emails from the specified archive.
func mailmanager_StartArchiveExport(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.StartArchiveExportInput{
		// ArchiveId: *string, // Required
		// ExportDestinationConfiguration: types.ExportDestinationConfiguration, // Required
		// FromTimestamp: *time.Time, // Required
		// ToTimestamp: *time.Time, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}
	if len(_mailmanagerExportDestinationConfiguration) > 0 {
		if err := assignInputField(input, "ExportDestinationConfiguration", _mailmanagerExportDestinationConfiguration); err != nil {
			log.Errorf("invalid --export-destination-configuration: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerFromTimestamp) > 0 {
		if err := assignInputField(input, "FromTimestamp", _mailmanagerFromTimestamp); err != nil {
			log.Errorf("invalid --from-timestamp: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerToTimestamp) > 0 {
		if err := assignInputField(input, "ToTimestamp", _mailmanagerToTimestamp); err != nil {
			log.Errorf("invalid --to-timestamp: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _mailmanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerIncludeMetadata) > 0 {
		if err := assignInputField(input, "IncludeMetadata", _mailmanagerIncludeMetadata); err != nil {
			log.Errorf("invalid --include-metadata: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mailmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartArchiveExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a search across emails in the specified archive.
func mailmanager_StartArchiveSearch(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.StartArchiveSearchInput{
		// ArchiveId: *string, // Required
		// FromTimestamp: *time.Time, // Required
		// MaxResults: *int32, // Required
		// ToTimestamp: *time.Time, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}
	if len(_mailmanagerFromTimestamp) > 0 {
		if err := assignInputField(input, "FromTimestamp", _mailmanagerFromTimestamp); err != nil {
			log.Errorf("invalid --from-timestamp: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mailmanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerToTimestamp) > 0 {
		if err := assignInputField(input, "ToTimestamp", _mailmanagerToTimestamp); err != nil {
			log.Errorf("invalid --to-timestamp: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _mailmanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartArchiveSearch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an ongoing import job for an address list.
func mailmanager_StopAddressListImportJob(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.StopAddressListImportJobInput{
		// JobId: *string, // Required
	}

	if len(_mailmanagerJobId) > 0 {
		input.JobId = aws.String(_mailmanagerJobId)
	}

	if resp, err := client.StopAddressListImportJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an in-progress export of emails from an archive.
func mailmanager_StopArchiveExport(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.StopArchiveExportInput{
		// ExportId: *string, // Required
	}

	if len(_mailmanagerExportId) > 0 {
		input.ExportId = aws.String(_mailmanagerExportId)
	}

	if resp, err := client.StopArchiveExport(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops an in-progress archive search job.
func mailmanager_StopArchiveSearch(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.StopArchiveSearchInput{
		// SearchId: *string, // Required
	}

	if len(_mailmanagerSearchId) > 0 {
		input.SearchId = aws.String(_mailmanagerSearchId)
	}

	if resp, err := client.StopArchiveSearch(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more tags (keys and values) to a specified resource.
func mailmanager_TagResource(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_mailmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_mailmanagerResourceArn)
	}
	if len(_mailmanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _mailmanagerTags); err != nil {
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

// Remove one or more tags (keys and values) from a specified resource.
func mailmanager_UntagResource(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mailmanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_mailmanagerResourceArn)
	}
	if len(_mailmanagerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mailmanagerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the attributes of an existing email archive.
func mailmanager_UpdateArchive(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.UpdateArchiveInput{
		// ArchiveId: *string, // Required
	}

	if len(_mailmanagerArchiveId) > 0 {
		input.ArchiveId = aws.String(_mailmanagerArchiveId)
	}
	if len(_mailmanagerArchiveName) > 0 {
		input.ArchiveName = aws.String(_mailmanagerArchiveName)
	}
	if len(_mailmanagerRetention) > 0 {
		if err := assignInputField(input, "Retention", _mailmanagerRetention); err != nil {
			log.Errorf("invalid --retention: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateArchive(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update attributes of a provisioned ingress endpoint resource.
func mailmanager_UpdateIngressPoint(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.UpdateIngressPointInput{
		// IngressPointId: *string, // Required
	}

	if len(_mailmanagerIngressPointId) > 0 {
		input.IngressPointId = aws.String(_mailmanagerIngressPointId)
	}
	if len(_mailmanagerIngressPointConfiguration) > 0 {
		if err := assignInputField(input, "IngressPointConfiguration", _mailmanagerIngressPointConfiguration); err != nil {
			log.Errorf("invalid --ingress-point-configuration: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerIngressPointName) > 0 {
		input.IngressPointName = aws.String(_mailmanagerIngressPointName)
	}
	if len(_mailmanagerRuleSetId) > 0 {
		input.RuleSetId = aws.String(_mailmanagerRuleSetId)
	}
	if len(_mailmanagerStatusToUpdate) > 0 {
		if err := assignInputField(input, "StatusToUpdate", _mailmanagerStatusToUpdate); err != nil {
			log.Errorf("invalid --status-to-update: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerTrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_mailmanagerTrafficPolicyId)
	}

	if resp, err := client.UpdateIngressPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the attributes of an existing relay resource.
func mailmanager_UpdateRelay(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.UpdateRelayInput{
		// RelayId: *string, // Required
	}

	if len(_mailmanagerRelayId) > 0 {
		input.RelayId = aws.String(_mailmanagerRelayId)
	}
	if len(_mailmanagerAuthentication) > 0 {
		if err := assignInputField(input, "Authentication", _mailmanagerAuthentication); err != nil {
			log.Errorf("invalid --authentication: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerRelayName) > 0 {
		input.RelayName = aws.String(_mailmanagerRelayName)
	}
	if len(_mailmanagerServerName) > 0 {
		input.ServerName = aws.String(_mailmanagerServerName)
	}
	if len(_mailmanagerServerPort) > 0 {
		if err := assignInputField(input, "ServerPort", _mailmanagerServerPort); err != nil {
			log.Errorf("invalid --server-port: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRelay(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update attributes of an already provisioned rule set.
func mailmanager_UpdateRuleSet(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.UpdateRuleSetInput{
		// RuleSetId: *string, // Required
	}

	if len(_mailmanagerRuleSetId) > 0 {
		input.RuleSetId = aws.String(_mailmanagerRuleSetId)
	}
	if len(_mailmanagerRuleSetName) > 0 {
		input.RuleSetName = aws.String(_mailmanagerRuleSetName)
	}
	if len(_mailmanagerRules) > 0 {
		if err := assignInputField(input, "Rules", _mailmanagerRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRuleSet(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Update attributes of an already provisioned traffic policy resource.
func mailmanager_UpdateTrafficPolicy(cfg aws.Config, client *mailmanager.Client) {
	input := &mailmanager.UpdateTrafficPolicyInput{
		// TrafficPolicyId: *string, // Required
	}

	if len(_mailmanagerTrafficPolicyId) > 0 {
		input.TrafficPolicyId = aws.String(_mailmanagerTrafficPolicyId)
	}
	if len(_mailmanagerDefaultAction) > 0 {
		if err := assignInputField(input, "DefaultAction", _mailmanagerDefaultAction); err != nil {
			log.Errorf("invalid --default-action: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerMaxMessageSizeBytes) > 0 {
		if err := assignInputField(input, "MaxMessageSizeBytes", _mailmanagerMaxMessageSizeBytes); err != nil {
			log.Errorf("invalid --max-message-size-bytes: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerPolicyStatements) > 0 {
		if err := assignInputField(input, "PolicyStatements", _mailmanagerPolicyStatements); err != nil {
			log.Errorf("invalid --policy-statements: %s", err.Error())
			return
		}
	}
	if len(_mailmanagerTrafficPolicyName) > 0 {
		input.TrafficPolicyName = aws.String(_mailmanagerTrafficPolicyName)
	}

	if resp, err := client.UpdateTrafficPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mailmanagerCmd)
	_mailmanagerCmd.Flags().SortFlags = false

	_mailmanagerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mailmanagerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mailmanagerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAddonInstanceId, "addon-instance-id", "", "", "Addon Instance ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAddonName, "addon-name", "", "", "Addon Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAddonSubscriptionId, "addon-subscription-id", "", "", "Addon Subscription ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAddress, "address", "", "", "Address")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAddressListId, "address-list-id", "", "", "Address List ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAddressListName, "address-list-name", "", "", "Address List Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerArchiveId, "archive-id", "", "", "Archive ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerArchiveName, "archive-name", "", "", "Archive Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerArchivedMessageId, "archived-message-id", "", "", "Archived Message ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerAuthentication, "authentication", "", "", "Authentication")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerClientToken, "client-token", "", "", "Client Token")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerDefaultAction, "default-action", "", "", "Default Action")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerExportDestinationConfiguration, "export-destination-configuration", "", "", "Export Destination Configuration")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerExportId, "export-id", "", "", "Export ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerFilter, "filter", "", "", "Filter")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerFilters, "filters", "", "", "Filters")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerFromTimestamp, "from-timestamp", "", "", "From Timestamp")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerImportDataFormat, "import-data-format", "", "", "Import Data Format")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerIncludeMetadata, "include-metadata", "", "", "Include Metadata")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerIngressPointConfiguration, "ingress-point-configuration", "", "", "Ingress Point Configuration")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerIngressPointId, "ingress-point-id", "", "", "Ingress Point ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerIngressPointName, "ingress-point-name", "", "", "Ingress Point Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerJobId, "job-id", "", "", "Job ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerKmsKeyArn, "kms-key-arn", "", "", "KMS Key ARN")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerMaxMessageSizeBytes, "max-message-size-bytes", "", "", "Max Message Size Bytes")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerMaxResults, "max-results", "", "", "Max Results")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerName, "name", "", "", "Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerNetworkConfiguration, "network-configuration", "", "", "Network Configuration")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerNextToken, "next-token", "", "", "Next Token")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerPageSize, "page-size", "", "", "Page Size")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerPolicyStatements, "policy-statements", "", "", "Policy Statements")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerRelayId, "relay-id", "", "", "Relay ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerRelayName, "relay-name", "", "", "Relay Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerResourceArn, "resource-arn", "", "", "Resource ARN")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerRetention, "retention", "", "", "Retention")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerRuleSetId, "rule-set-id", "", "", "Rule Set ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerRuleSetName, "rule-set-name", "", "", "Rule Set Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerRules, "rules", "", "", "Rules")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerSearchId, "search-id", "", "", "Search ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerServerName, "server-name", "", "", "Server Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerServerPort, "server-port", "", "", "Server Port")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerStatusToUpdate, "status-to-update", "", "", "Status To Update")
	_mailmanagerCmd.Flags().StringSliceVarP(&_mailmanagerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerTags, "tags", "", "", "Tags")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerToTimestamp, "to-timestamp", "", "", "To Timestamp")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerTrafficPolicyId, "traffic-policy-id", "", "", "Traffic Policy ID")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerTrafficPolicyName, "traffic-policy-name", "", "", "Traffic Policy Name")
	_mailmanagerCmd.Flags().StringVarP(&_mailmanagerType, "type", "", "", "Type")

	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateAddonInstance, "create-addon-instance", "", false, "Create Addon Instance")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateAddonSubscription, "create-addon-subscription", "", false, "Create Addon Subscription")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateAddressList, "create-address-list", "", false, "Create Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateAddressListImportJob, "create-address-list-import-job", "", false, "Create Address List Import Job")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateArchive, "create-archive", "", false, "Create Archive")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateIngressPoint, "create-ingress-point", "", false, "Create Ingress Point")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateRelay, "create-relay", "", false, "Create Relay")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateRuleSet, "create-rule-set", "", false, "Create Rule Set")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerCreateTrafficPolicy, "create-traffic-policy", "", false, "Create Traffic Policy")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteAddonInstance, "delete-addon-instance", "", false, "Delete Addon Instance")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteAddonSubscription, "delete-addon-subscription", "", false, "Delete Addon Subscription")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteAddressList, "delete-address-list", "", false, "Delete Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteArchive, "delete-archive", "", false, "Delete Archive")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteIngressPoint, "delete-ingress-point", "", false, "Delete Ingress Point")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteRelay, "delete-relay", "", false, "Delete Relay")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteRuleSet, "delete-rule-set", "", false, "Delete Rule Set")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeleteTrafficPolicy, "delete-traffic-policy", "", false, "Delete Traffic Policy")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerDeregisterMemberFromAddressList, "deregister-member-from-address-list", "", false, "Deregister Member From Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetAddonInstance, "get-addon-instance", "", false, "Get Addon Instance")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetAddonSubscription, "get-addon-subscription", "", false, "Get Addon Subscription")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetAddressList, "get-address-list", "", false, "Get Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetAddressListImportJob, "get-address-list-import-job", "", false, "Get Address List Import Job")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetArchive, "get-archive", "", false, "Get Archive")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetArchiveExport, "get-archive-export", "", false, "Get Archive Export")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetArchiveMessage, "get-archive-message", "", false, "Get Archive Message")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetArchiveMessageContent, "get-archive-message-content", "", false, "Get Archive Message Content")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetArchiveSearch, "get-archive-search", "", false, "Get Archive Search")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetArchiveSearchResults, "get-archive-search-results", "", false, "Get Archive Search Results")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetIngressPoint, "get-ingress-point", "", false, "Get Ingress Point")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetMemberOfAddressList, "get-member-of-address-list", "", false, "Get Member Of Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetRelay, "get-relay", "", false, "Get Relay")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetRuleSet, "get-rule-set", "", false, "Get Rule Set")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerGetTrafficPolicy, "get-traffic-policy", "", false, "Get Traffic Policy")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListAddonInstances, "list-addon-instances", "", false, "List Addon Instances")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListAddonSubscriptions, "list-addon-subscriptions", "", false, "List Addon Subscriptions")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListAddressListImportJobs, "list-address-list-import-jobs", "", false, "List Address List Import Jobs")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListAddressLists, "list-address-lists", "", false, "List Address Lists")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListArchiveExports, "list-archive-exports", "", false, "List Archive Exports")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListArchiveSearches, "list-archive-searches", "", false, "List Archive Searches")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListArchives, "list-archives", "", false, "List Archives")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListIngressPoints, "list-ingress-points", "", false, "List Ingress Points")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListMembersOfAddressList, "list-members-of-address-list", "", false, "List Members Of Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListRelays, "list-relays", "", false, "List Relays")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListRuleSets, "list-rule-sets", "", false, "List Rule Sets")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerListTrafficPolicies, "list-traffic-policies", "", false, "List Traffic Policies")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerRegisterMemberToAddressList, "register-member-to-address-list", "", false, "Register Member To Address List")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerStartAddressListImportJob, "start-address-list-import-job", "", false, "Start Address List Import Job")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerStartArchiveExport, "start-archive-export", "", false, "Start Archive Export")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerStartArchiveSearch, "start-archive-search", "", false, "Start Archive Search")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerStopAddressListImportJob, "stop-address-list-import-job", "", false, "Stop Address List Import Job")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerStopArchiveExport, "stop-archive-export", "", false, "Stop Archive Export")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerStopArchiveSearch, "stop-archive-search", "", false, "Stop Archive Search")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerTagResource, "tag-resource", "", false, "Tag Resource")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerUntagResource, "untag-resource", "", false, "Untag Resource")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerUpdateArchive, "update-archive", "", false, "Update Archive")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerUpdateIngressPoint, "update-ingress-point", "", false, "Update Ingress Point")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerUpdateRelay, "update-relay", "", false, "Update Relay")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerUpdateRuleSet, "update-rule-set", "", false, "Update Rule Set")
	_mailmanagerCmd.Flags().BoolVarP(&_mailmanagerUpdateTrafficPolicy, "update-traffic-policy", "", false, "Update Traffic Policy")

}
