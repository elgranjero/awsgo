package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// licensemanagerCmd represents the licensemanager command
var _licensemanagerCmd = &cobra.Command{
	Use:   "licensemanager",
	Short: "AWS licensemanager CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := licensemanager.NewFromConfig(cfg)
		if _licensemanagerAcceptGrant {
			licensemanager_AcceptGrant(cfg, client)
			return
		}
		if _licensemanagerCheckInLicense {
			licensemanager_CheckInLicense(cfg, client)
			return
		}
		if _licensemanagerCheckoutBorrowLicense {
			licensemanager_CheckoutBorrowLicense(cfg, client)
			return
		}
		if _licensemanagerCheckoutLicense {
			licensemanager_CheckoutLicense(cfg, client)
			return
		}
		if _licensemanagerCreateGrant {
			licensemanager_CreateGrant(cfg, client)
			return
		}
		if _licensemanagerCreateGrantVersion {
			licensemanager_CreateGrantVersion(cfg, client)
			return
		}
		if _licensemanagerCreateLicense {
			licensemanager_CreateLicense(cfg, client)
			return
		}
		if _licensemanagerCreateLicenseAssetGroup {
			licensemanager_CreateLicenseAssetGroup(cfg, client)
			return
		}
		if _licensemanagerCreateLicenseAssetRuleset {
			licensemanager_CreateLicenseAssetRuleset(cfg, client)
			return
		}
		if _licensemanagerCreateLicenseConfiguration {
			licensemanager_CreateLicenseConfiguration(cfg, client)
			return
		}
		if _licensemanagerCreateLicenseConversionTaskForResource {
			licensemanager_CreateLicenseConversionTaskForResource(cfg, client)
			return
		}
		if _licensemanagerCreateLicenseManagerReportGenerator {
			licensemanager_CreateLicenseManagerReportGenerator(cfg, client)
			return
		}
		if _licensemanagerCreateLicenseVersion {
			licensemanager_CreateLicenseVersion(cfg, client)
			return
		}
		if _licensemanagerCreateToken {
			licensemanager_CreateToken(cfg, client)
			return
		}
		if _licensemanagerDeleteGrant {
			licensemanager_DeleteGrant(cfg, client)
			return
		}
		if _licensemanagerDeleteLicense {
			licensemanager_DeleteLicense(cfg, client)
			return
		}
		if _licensemanagerDeleteLicenseAssetGroup {
			licensemanager_DeleteLicenseAssetGroup(cfg, client)
			return
		}
		if _licensemanagerDeleteLicenseAssetRuleset {
			licensemanager_DeleteLicenseAssetRuleset(cfg, client)
			return
		}
		if _licensemanagerDeleteLicenseConfiguration {
			licensemanager_DeleteLicenseConfiguration(cfg, client)
			return
		}
		if _licensemanagerDeleteLicenseManagerReportGenerator {
			licensemanager_DeleteLicenseManagerReportGenerator(cfg, client)
			return
		}
		if _licensemanagerDeleteToken {
			licensemanager_DeleteToken(cfg, client)
			return
		}
		if _licensemanagerExtendLicenseConsumption {
			licensemanager_ExtendLicenseConsumption(cfg, client)
			return
		}
		if _licensemanagerGetAccessToken {
			licensemanager_GetAccessToken(cfg, client)
			return
		}
		if _licensemanagerGetGrant {
			licensemanager_GetGrant(cfg, client)
			return
		}
		if _licensemanagerGetLicense {
			licensemanager_GetLicense(cfg, client)
			return
		}
		if _licensemanagerGetLicenseAssetGroup {
			licensemanager_GetLicenseAssetGroup(cfg, client)
			return
		}
		if _licensemanagerGetLicenseAssetRuleset {
			licensemanager_GetLicenseAssetRuleset(cfg, client)
			return
		}
		if _licensemanagerGetLicenseConfiguration {
			licensemanager_GetLicenseConfiguration(cfg, client)
			return
		}
		if _licensemanagerGetLicenseConversionTask {
			licensemanager_GetLicenseConversionTask(cfg, client)
			return
		}
		if _licensemanagerGetLicenseManagerReportGenerator {
			licensemanager_GetLicenseManagerReportGenerator(cfg, client)
			return
		}
		if _licensemanagerGetLicenseUsage {
			licensemanager_GetLicenseUsage(cfg, client)
			return
		}
		if _licensemanagerGetServiceSettings {
			licensemanager_GetServiceSettings(cfg, client)
			return
		}
		if _licensemanagerListAssetsForLicenseAssetGroup {
			licensemanager_ListAssetsForLicenseAssetGroup(cfg, client)
			return
		}
		if _licensemanagerListAssociationsForLicenseConfiguration {
			licensemanager_ListAssociationsForLicenseConfiguration(cfg, client)
			return
		}
		if _licensemanagerListDistributedGrants {
			licensemanager_ListDistributedGrants(cfg, client)
			return
		}
		if _licensemanagerListFailuresForLicenseConfigurationOperations {
			licensemanager_ListFailuresForLicenseConfigurationOperations(cfg, client)
			return
		}
		if _licensemanagerListLicenseAssetGroups {
			licensemanager_ListLicenseAssetGroups(cfg, client)
			return
		}
		if _licensemanagerListLicenseAssetRulesets {
			licensemanager_ListLicenseAssetRulesets(cfg, client)
			return
		}
		if _licensemanagerListLicenseConfigurations {
			licensemanager_ListLicenseConfigurations(cfg, client)
			return
		}
		if _licensemanagerListLicenseConfigurationsForOrganization {
			licensemanager_ListLicenseConfigurationsForOrganization(cfg, client)
			return
		}
		if _licensemanagerListLicenseConversionTasks {
			licensemanager_ListLicenseConversionTasks(cfg, client)
			return
		}
		if _licensemanagerListLicenseManagerReportGenerators {
			licensemanager_ListLicenseManagerReportGenerators(cfg, client)
			return
		}
		if _licensemanagerListLicenseSpecificationsForResource {
			licensemanager_ListLicenseSpecificationsForResource(cfg, client)
			return
		}
		if _licensemanagerListLicenseVersions {
			licensemanager_ListLicenseVersions(cfg, client)
			return
		}
		if _licensemanagerListLicenses {
			licensemanager_ListLicenses(cfg, client)
			return
		}
		if _licensemanagerListReceivedGrants {
			licensemanager_ListReceivedGrants(cfg, client)
			return
		}
		if _licensemanagerListReceivedGrantsForOrganization {
			licensemanager_ListReceivedGrantsForOrganization(cfg, client)
			return
		}
		if _licensemanagerListReceivedLicenses {
			licensemanager_ListReceivedLicenses(cfg, client)
			return
		}
		if _licensemanagerListReceivedLicensesForOrganization {
			licensemanager_ListReceivedLicensesForOrganization(cfg, client)
			return
		}
		if _licensemanagerListResourceInventory {
			licensemanager_ListResourceInventory(cfg, client)
			return
		}
		if _licensemanagerListTagsForResource {
			licensemanager_ListTagsForResource(cfg, client)
			return
		}
		if _licensemanagerListTokens {
			licensemanager_ListTokens(cfg, client)
			return
		}
		if _licensemanagerListUsageForLicenseConfiguration {
			licensemanager_ListUsageForLicenseConfiguration(cfg, client)
			return
		}
		if _licensemanagerRejectGrant {
			licensemanager_RejectGrant(cfg, client)
			return
		}
		if _licensemanagerTagResource {
			licensemanager_TagResource(cfg, client)
			return
		}
		if _licensemanagerUntagResource {
			licensemanager_UntagResource(cfg, client)
			return
		}
		if _licensemanagerUpdateLicenseAssetGroup {
			licensemanager_UpdateLicenseAssetGroup(cfg, client)
			return
		}
		if _licensemanagerUpdateLicenseAssetRuleset {
			licensemanager_UpdateLicenseAssetRuleset(cfg, client)
			return
		}
		if _licensemanagerUpdateLicenseConfiguration {
			licensemanager_UpdateLicenseConfiguration(cfg, client)
			return
		}
		if _licensemanagerUpdateLicenseManagerReportGenerator {
			licensemanager_UpdateLicenseManagerReportGenerator(cfg, client)
			return
		}
		if _licensemanagerUpdateLicenseSpecificationsForResource {
			licensemanager_UpdateLicenseSpecificationsForResource(cfg, client)
			return
		}
		if _licensemanagerUpdateServiceSettings {
			licensemanager_UpdateServiceSettings(cfg, client)
			return
		}

	},
}

var (
	_licensemanagerAcceptGrant                                   bool
	_licensemanagerCheckInLicense                                bool
	_licensemanagerCheckoutBorrowLicense                         bool
	_licensemanagerCheckoutLicense                               bool
	_licensemanagerCreateGrant                                   bool
	_licensemanagerCreateGrantVersion                            bool
	_licensemanagerCreateLicense                                 bool
	_licensemanagerCreateLicenseAssetGroup                       bool
	_licensemanagerCreateLicenseAssetRuleset                     bool
	_licensemanagerCreateLicenseConfiguration                    bool
	_licensemanagerCreateLicenseConversionTaskForResource        bool
	_licensemanagerCreateLicenseManagerReportGenerator           bool
	_licensemanagerCreateLicenseVersion                          bool
	_licensemanagerCreateToken                                   bool
	_licensemanagerDeleteGrant                                   bool
	_licensemanagerDeleteLicense                                 bool
	_licensemanagerDeleteLicenseAssetGroup                       bool
	_licensemanagerDeleteLicenseAssetRuleset                     bool
	_licensemanagerDeleteLicenseConfiguration                    bool
	_licensemanagerDeleteLicenseManagerReportGenerator           bool
	_licensemanagerDeleteToken                                   bool
	_licensemanagerExtendLicenseConsumption                      bool
	_licensemanagerGetAccessToken                                bool
	_licensemanagerGetGrant                                      bool
	_licensemanagerGetLicense                                    bool
	_licensemanagerGetLicenseAssetGroup                          bool
	_licensemanagerGetLicenseAssetRuleset                        bool
	_licensemanagerGetLicenseConfiguration                       bool
	_licensemanagerGetLicenseConversionTask                      bool
	_licensemanagerGetLicenseManagerReportGenerator              bool
	_licensemanagerGetLicenseUsage                               bool
	_licensemanagerGetServiceSettings                            bool
	_licensemanagerListAssetsForLicenseAssetGroup                bool
	_licensemanagerListAssociationsForLicenseConfiguration       bool
	_licensemanagerListDistributedGrants                         bool
	_licensemanagerListFailuresForLicenseConfigurationOperations bool
	_licensemanagerListLicenseAssetGroups                        bool
	_licensemanagerListLicenseAssetRulesets                      bool
	_licensemanagerListLicenseConfigurations                     bool
	_licensemanagerListLicenseConfigurationsForOrganization      bool
	_licensemanagerListLicenseConversionTasks                    bool
	_licensemanagerListLicenseManagerReportGenerators            bool
	_licensemanagerListLicenseSpecificationsForResource          bool
	_licensemanagerListLicenseVersions                           bool
	_licensemanagerListLicenses                                  bool
	_licensemanagerListReceivedGrants                            bool
	_licensemanagerListReceivedGrantsForOrganization             bool
	_licensemanagerListReceivedLicenses                          bool
	_licensemanagerListReceivedLicensesForOrganization           bool
	_licensemanagerListResourceInventory                         bool
	_licensemanagerListTagsForResource                           bool
	_licensemanagerListTokens                                    bool
	_licensemanagerListUsageForLicenseConfiguration              bool
	_licensemanagerRejectGrant                                   bool
	_licensemanagerTagResource                                   bool
	_licensemanagerUntagResource                                 bool
	_licensemanagerUpdateLicenseAssetGroup                       bool
	_licensemanagerUpdateLicenseAssetRuleset                     bool
	_licensemanagerUpdateLicenseConfiguration                    bool
	_licensemanagerUpdateLicenseManagerReportGenerator           bool
	_licensemanagerUpdateLicenseSpecificationsForResource        bool
	_licensemanagerUpdateServiceSettings                         bool

	_licensemanagerAddLicenseSpecifications           string
	_licensemanagerAllowedOperations                  string
	_licensemanagerAssetType                          string
	_licensemanagerAssociatedLicenseAssetRulesetARNs  []string
	_licensemanagerBeneficiary                        string
	_licensemanagerCheckoutMetadata                   string
	_licensemanagerCheckoutType                       string
	_licensemanagerClientToken                        string
	_licensemanagerConsumptionConfiguration           string
	_licensemanagerDescription                        string
	_licensemanagerDestinationLicenseContext          string
	_licensemanagerDigitalSignatureMethod             string
	_licensemanagerDisassociateWhenNotFound           string
	_licensemanagerDryRun                             string
	_licensemanagerEnableCrossAccountsDiscovery       string
	_licensemanagerEnabledDiscoverySourceRegions      []string
	_licensemanagerEntitlements                       string
	_licensemanagerExpirationInDays                   string
	_licensemanagerFilters                            string
	_licensemanagerGrantArn                           string
	_licensemanagerGrantArns                          []string
	_licensemanagerGrantName                          string
	_licensemanagerHomeRegion                         string
	_licensemanagerIssuer                             string
	_licensemanagerKeyFingerprint                     string
	_licensemanagerLicenseArn                         string
	_licensemanagerLicenseArns                        []string
	_licensemanagerLicenseAssetGroupArn               string
	_licensemanagerLicenseAssetGroupConfigurations    string
	_licensemanagerLicenseAssetRulesetArn             string
	_licensemanagerLicenseConfigurationArn            string
	_licensemanagerLicenseConfigurationArns           []string
	_licensemanagerLicenseConfigurationStatus         string
	_licensemanagerLicenseConsumptionToken            string
	_licensemanagerLicenseConversionTaskId            string
	_licensemanagerLicenseCount                       string
	_licensemanagerLicenseCountHardLimit              string
	_licensemanagerLicenseCountingType                string
	_licensemanagerLicenseExpiry                      string
	_licensemanagerLicenseManagerReportGeneratorArn   string
	_licensemanagerLicenseMetadata                    string
	_licensemanagerLicenseName                        string
	_licensemanagerLicenseRules                       []string
	_licensemanagerMaxResults                         string
	_licensemanagerName                               string
	_licensemanagerNextToken                          string
	_licensemanagerNodeId                             string
	_licensemanagerOptions                            string
	_licensemanagerOrganizationConfiguration          string
	_licensemanagerPrincipals                         []string
	_licensemanagerProductInformationList             string
	_licensemanagerProductName                        string
	_licensemanagerProductSKU                         string
	_licensemanagerProperties                         string
	_licensemanagerRemoveLicenseSpecifications        string
	_licensemanagerReportContext                      string
	_licensemanagerReportFrequency                    string
	_licensemanagerReportGeneratorName                string
	_licensemanagerResourceArn                        string
	_licensemanagerRoleArns                           []string
	_licensemanagerRules                              string
	_licensemanagerS3BucketArn                        string
	_licensemanagerShowAWSManagedLicenseAssetRulesets string
	_licensemanagerSnsTopicArn                        string
	_licensemanagerSourceLicenseContext               string
	_licensemanagerSourceVersion                      string
	_licensemanagerStatus                             string
	_licensemanagerStatusReason                       string
	_licensemanagerTagKeys                            []string
	_licensemanagerTags                               string
	_licensemanagerToken                              string
	_licensemanagerTokenId                            string
	_licensemanagerTokenIds                           []string
	_licensemanagerTokenProperties                    []string
	_licensemanagerType                               string
	_licensemanagerValidity                           string
	_licensemanagerVersion                            string
)

// Accepts the specified grant.
func licensemanager_AcceptGrant(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.AcceptGrantInput{
		// GrantArn: *string, // Required
	}

	if len(_licensemanagerGrantArn) > 0 {
		input.GrantArn = aws.String(_licensemanagerGrantArn)
	}

	if resp, err := client.AcceptGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks in the specified license. Check in a license when it is no longer in use.
func licensemanager_CheckInLicense(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CheckInLicenseInput{
		// LicenseConsumptionToken: *string, // Required
	}

	if len(_licensemanagerLicenseConsumptionToken) > 0 {
		input.LicenseConsumptionToken = aws.String(_licensemanagerLicenseConsumptionToken)
	}
	if len(_licensemanagerBeneficiary) > 0 {
		input.Beneficiary = aws.String(_licensemanagerBeneficiary)
	}

	if resp, err := client.CheckInLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks out the specified license for offline use.
func licensemanager_CheckoutBorrowLicense(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CheckoutBorrowLicenseInput{
		// ClientToken: *string, // Required
		// DigitalSignatureMethod: types.DigitalSignatureMethod, // Required
		// Entitlements: []types.EntitlementData, // Required
		// LicenseArn: *string, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerDigitalSignatureMethod) > 0 {
		if err := assignInputField(input, "DigitalSignatureMethod", _licensemanagerDigitalSignatureMethod); err != nil {
			log.Errorf("invalid --digital-signature-method: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerEntitlements) > 0 {
		if err := assignInputField(input, "Entitlements", _licensemanagerEntitlements); err != nil {
			log.Errorf("invalid --entitlements: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerCheckoutMetadata) > 0 {
		if err := assignInputField(input, "CheckoutMetadata", _licensemanagerCheckoutMetadata); err != nil {
			log.Errorf("invalid --checkout-metadata: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNodeId) > 0 {
		input.NodeId = aws.String(_licensemanagerNodeId)
	}

	if resp, err := client.CheckoutBorrowLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks out the specified license.
// If the account that created the license is the same that is performing the
// check out, you must specify the account as the beneficiary.
func licensemanager_CheckoutLicense(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CheckoutLicenseInput{
		// CheckoutType: types.CheckoutType, // Required
		// ClientToken: *string, // Required
		// Entitlements: []types.EntitlementData, // Required
		// KeyFingerprint: *string, // Required
		// ProductSKU: *string, // Required
	}

	if len(_licensemanagerCheckoutType) > 0 {
		if err := assignInputField(input, "CheckoutType", _licensemanagerCheckoutType); err != nil {
			log.Errorf("invalid --checkout-type: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerEntitlements) > 0 {
		if err := assignInputField(input, "Entitlements", _licensemanagerEntitlements); err != nil {
			log.Errorf("invalid --entitlements: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerKeyFingerprint) > 0 {
		input.KeyFingerprint = aws.String(_licensemanagerKeyFingerprint)
	}
	if len(_licensemanagerProductSKU) > 0 {
		input.ProductSKU = aws.String(_licensemanagerProductSKU)
	}
	if len(_licensemanagerBeneficiary) > 0 {
		input.Beneficiary = aws.String(_licensemanagerBeneficiary)
	}
	if len(_licensemanagerNodeId) > 0 {
		input.NodeId = aws.String(_licensemanagerNodeId)
	}

	if resp, err := client.CheckoutLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a grant for the specified license. A grant shares the use of license
// entitlements with a specific Amazon Web Services account, an organization, or an
// organizational unit (OU). For more information, see [Granted licenses in License Manager]in the License Manager User
// Guide.
//
// [Granted licenses in License Manager]: https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html
func licensemanager_CreateGrant(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateGrantInput{
		// AllowedOperations: []types.AllowedOperation, // Required
		// ClientToken: *string, // Required
		// GrantName: *string, // Required
		// HomeRegion: *string, // Required
		// LicenseArn: *string, // Required
		// Principals: []string, // Required
	}

	if len(_licensemanagerAllowedOperations) > 0 {
		if err := assignInputField(input, "AllowedOperations", _licensemanagerAllowedOperations); err != nil {
			log.Errorf("invalid --allowed-operations: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerGrantName) > 0 {
		input.GrantName = aws.String(_licensemanagerGrantName)
	}
	if len(_licensemanagerHomeRegion) > 0 {
		input.HomeRegion = aws.String(_licensemanagerHomeRegion)
	}
	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerPrincipals) > 0 {
		input.Principals = append([]string(nil), _licensemanagerPrincipals...)
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of the specified grant. For more information, see [Granted licenses in License Manager] in the
// License Manager User Guide.
//
// [Granted licenses in License Manager]: https://docs.aws.amazon.com/license-manager/latest/userguide/granted-licenses.html
func licensemanager_CreateGrantVersion(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateGrantVersionInput{
		// ClientToken: *string, // Required
		// GrantArn: *string, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerGrantArn) > 0 {
		input.GrantArn = aws.String(_licensemanagerGrantArn)
	}
	if len(_licensemanagerAllowedOperations) > 0 {
		if err := assignInputField(input, "AllowedOperations", _licensemanagerAllowedOperations); err != nil {
			log.Errorf("invalid --allowed-operations: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerGrantName) > 0 {
		input.GrantName = aws.String(_licensemanagerGrantName)
	}
	if len(_licensemanagerOptions) > 0 {
		if err := assignInputField(input, "Options", _licensemanagerOptions); err != nil {
			log.Errorf("invalid --options: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerSourceVersion) > 0 {
		input.SourceVersion = aws.String(_licensemanagerSourceVersion)
	}
	if len(_licensemanagerStatus) > 0 {
		if err := assignInputField(input, "Status", _licensemanagerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerStatusReason) > 0 {
		input.StatusReason = aws.String(_licensemanagerStatusReason)
	}

	if resp, err := client.CreateGrantVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a license.
func licensemanager_CreateLicense(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseInput{
		// Beneficiary: *string, // Required
		// ClientToken: *string, // Required
		// ConsumptionConfiguration: *types.ConsumptionConfiguration, // Required
		// Entitlements: []types.Entitlement, // Required
		// HomeRegion: *string, // Required
		// Issuer: *types.Issuer, // Required
		// LicenseName: *string, // Required
		// ProductName: *string, // Required
		// ProductSKU: *string, // Required
		// Validity: *types.DatetimeRange, // Required
	}

	if len(_licensemanagerBeneficiary) > 0 {
		input.Beneficiary = aws.String(_licensemanagerBeneficiary)
	}
	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerConsumptionConfiguration) > 0 {
		if err := assignInputField(input, "ConsumptionConfiguration", _licensemanagerConsumptionConfiguration); err != nil {
			log.Errorf("invalid --consumption-configuration: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerEntitlements) > 0 {
		if err := assignInputField(input, "Entitlements", _licensemanagerEntitlements); err != nil {
			log.Errorf("invalid --entitlements: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerHomeRegion) > 0 {
		input.HomeRegion = aws.String(_licensemanagerHomeRegion)
	}
	if len(_licensemanagerIssuer) > 0 {
		if err := assignInputField(input, "Issuer", _licensemanagerIssuer); err != nil {
			log.Errorf("invalid --issuer: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseName) > 0 {
		input.LicenseName = aws.String(_licensemanagerLicenseName)
	}
	if len(_licensemanagerProductName) > 0 {
		input.ProductName = aws.String(_licensemanagerProductName)
	}
	if len(_licensemanagerProductSKU) > 0 {
		input.ProductSKU = aws.String(_licensemanagerProductSKU)
	}
	if len(_licensemanagerValidity) > 0 {
		if err := assignInputField(input, "Validity", _licensemanagerValidity); err != nil {
			log.Errorf("invalid --validity: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseMetadata) > 0 {
		if err := assignInputField(input, "LicenseMetadata", _licensemanagerLicenseMetadata); err != nil {
			log.Errorf("invalid --license-metadata: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a license asset group.
func licensemanager_CreateLicenseAssetGroup(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseAssetGroupInput{
		// AssociatedLicenseAssetRulesetARNs: []string, // Required
		// ClientToken: *string, // Required
		// LicenseAssetGroupConfigurations: []types.LicenseAssetGroupConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_licensemanagerAssociatedLicenseAssetRulesetARNs) > 0 {
		input.AssociatedLicenseAssetRulesetARNs = append([]string(nil), _licensemanagerAssociatedLicenseAssetRulesetARNs...)
	}
	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerLicenseAssetGroupConfigurations) > 0 {
		if err := assignInputField(input, "LicenseAssetGroupConfigurations", _licensemanagerLicenseAssetGroupConfigurations); err != nil {
			log.Errorf("invalid --license-asset-group-configurations: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerName) > 0 {
		input.Name = aws.String(_licensemanagerName)
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerProperties) > 0 {
		if err := assignInputField(input, "Properties", _licensemanagerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseAssetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a license asset ruleset.
func licensemanager_CreateLicenseAssetRuleset(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseAssetRulesetInput{
		// ClientToken: *string, // Required
		// Name: *string, // Required
		// Rules: []types.LicenseAssetRule, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerName) > 0 {
		input.Name = aws.String(_licensemanagerName)
	}
	if len(_licensemanagerRules) > 0 {
		if err := assignInputField(input, "Rules", _licensemanagerRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseAssetRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a license configuration.
// A license configuration is an abstraction of a customer license agreement that
// can be consumed and enforced by License Manager. Components include
// specifications for the license type (licensing by instance, socket, CPU, or
// vCPU), allowed tenancy (shared tenancy, Dedicated Instance, Dedicated Host, or
// all of these), license affinity to host (how long a license must be associated
// with a host), and the number of licenses purchased and used.
func licensemanager_CreateLicenseConfiguration(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseConfigurationInput{
		// LicenseCountingType: types.LicenseCountingType, // Required
		// Name: *string, // Required
	}

	if len(_licensemanagerLicenseCountingType) > 0 {
		if err := assignInputField(input, "LicenseCountingType", _licensemanagerLicenseCountingType); err != nil {
			log.Errorf("invalid --license-counting-type: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerName) > 0 {
		input.Name = aws.String(_licensemanagerName)
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerDisassociateWhenNotFound) > 0 {
		if err := assignInputField(input, "DisassociateWhenNotFound", _licensemanagerDisassociateWhenNotFound); err != nil {
			log.Errorf("invalid --disassociate-when-not-found: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseCount) > 0 {
		if err := assignInputField(input, "LicenseCount", _licensemanagerLicenseCount); err != nil {
			log.Errorf("invalid --license-count: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseCountHardLimit) > 0 {
		if err := assignInputField(input, "LicenseCountHardLimit", _licensemanagerLicenseCountHardLimit); err != nil {
			log.Errorf("invalid --license-count-hard-limit: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseExpiry) > 0 {
		if err := assignInputField(input, "LicenseExpiry", _licensemanagerLicenseExpiry); err != nil {
			log.Errorf("invalid --license-expiry: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseRules) > 0 {
		input.LicenseRules = append([]string(nil), _licensemanagerLicenseRules...)
	}
	if len(_licensemanagerProductInformationList) > 0 {
		if err := assignInputField(input, "ProductInformationList", _licensemanagerProductInformationList); err != nil {
			log.Errorf("invalid --product-information-list: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new license conversion task.
func licensemanager_CreateLicenseConversionTaskForResource(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseConversionTaskForResourceInput{
		// DestinationLicenseContext: *types.LicenseConversionContext, // Required
		// ResourceArn: *string, // Required
		// SourceLicenseContext: *types.LicenseConversionContext, // Required
	}

	if len(_licensemanagerDestinationLicenseContext) > 0 {
		if err := assignInputField(input, "DestinationLicenseContext", _licensemanagerDestinationLicenseContext); err != nil {
			log.Errorf("invalid --destination-license-context: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerResourceArn)
	}
	if len(_licensemanagerSourceLicenseContext) > 0 {
		if err := assignInputField(input, "SourceLicenseContext", _licensemanagerSourceLicenseContext); err != nil {
			log.Errorf("invalid --source-license-context: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseConversionTaskForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a report generator.
func licensemanager_CreateLicenseManagerReportGenerator(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseManagerReportGeneratorInput{
		// ClientToken: *string, // Required
		// ReportContext: *types.ReportContext, // Required
		// ReportFrequency: *types.ReportFrequency, // Required
		// ReportGeneratorName: *string, // Required
		// Type: []types.ReportType, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerReportContext) > 0 {
		if err := assignInputField(input, "ReportContext", _licensemanagerReportContext); err != nil {
			log.Errorf("invalid --report-context: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerReportFrequency) > 0 {
		if err := assignInputField(input, "ReportFrequency", _licensemanagerReportFrequency); err != nil {
			log.Errorf("invalid --report-frequency: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerReportGeneratorName) > 0 {
		input.ReportGeneratorName = aws.String(_licensemanagerReportGeneratorName)
	}
	if len(_licensemanagerType) > 0 {
		if err := assignInputField(input, "Type", _licensemanagerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateLicenseManagerReportGenerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of the specified license.
func licensemanager_CreateLicenseVersion(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateLicenseVersionInput{
		// ClientToken: *string, // Required
		// ConsumptionConfiguration: *types.ConsumptionConfiguration, // Required
		// Entitlements: []types.Entitlement, // Required
		// HomeRegion: *string, // Required
		// Issuer: *types.Issuer, // Required
		// LicenseArn: *string, // Required
		// LicenseName: *string, // Required
		// ProductName: *string, // Required
		// Status: types.LicenseStatus, // Required
		// Validity: *types.DatetimeRange, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerConsumptionConfiguration) > 0 {
		if err := assignInputField(input, "ConsumptionConfiguration", _licensemanagerConsumptionConfiguration); err != nil {
			log.Errorf("invalid --consumption-configuration: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerEntitlements) > 0 {
		if err := assignInputField(input, "Entitlements", _licensemanagerEntitlements); err != nil {
			log.Errorf("invalid --entitlements: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerHomeRegion) > 0 {
		input.HomeRegion = aws.String(_licensemanagerHomeRegion)
	}
	if len(_licensemanagerIssuer) > 0 {
		if err := assignInputField(input, "Issuer", _licensemanagerIssuer); err != nil {
			log.Errorf("invalid --issuer: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerLicenseName) > 0 {
		input.LicenseName = aws.String(_licensemanagerLicenseName)
	}
	if len(_licensemanagerProductName) > 0 {
		input.ProductName = aws.String(_licensemanagerProductName)
	}
	if len(_licensemanagerStatus) > 0 {
		if err := assignInputField(input, "Status", _licensemanagerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerValidity) > 0 {
		if err := assignInputField(input, "Validity", _licensemanagerValidity); err != nil {
			log.Errorf("invalid --validity: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseMetadata) > 0 {
		if err := assignInputField(input, "LicenseMetadata", _licensemanagerLicenseMetadata); err != nil {
			log.Errorf("invalid --license-metadata: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerSourceVersion) > 0 {
		input.SourceVersion = aws.String(_licensemanagerSourceVersion)
	}

	if resp, err := client.CreateLicenseVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a long-lived token.
// A refresh token is a JWT token used to get an access token. With an access
// token, you can call AssumeRoleWithWebIdentity to get role credentials that you
// can use to call License Manager to manage the specified license.
func licensemanager_CreateToken(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.CreateTokenInput{
		// ClientToken: *string, // Required
		// LicenseArn: *string, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerExpirationInDays) > 0 {
		if err := assignInputField(input, "ExpirationInDays", _licensemanagerExpirationInDays); err != nil {
			log.Errorf("invalid --expiration-in-days: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerRoleArns) > 0 {
		input.RoleArns = append([]string(nil), _licensemanagerRoleArns...)
	}
	if len(_licensemanagerTokenProperties) > 0 {
		input.TokenProperties = append([]string(nil), _licensemanagerTokenProperties...)
	}

	if resp, err := client.CreateToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified grant.
func licensemanager_DeleteGrant(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteGrantInput{
		// GrantArn: *string, // Required
		// Version: *string, // Required
	}

	if len(_licensemanagerGrantArn) > 0 {
		input.GrantArn = aws.String(_licensemanagerGrantArn)
	}
	if len(_licensemanagerVersion) > 0 {
		input.Version = aws.String(_licensemanagerVersion)
	}
	if len(_licensemanagerStatusReason) > 0 {
		input.StatusReason = aws.String(_licensemanagerStatusReason)
	}

	if resp, err := client.DeleteGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified license.
func licensemanager_DeleteLicense(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteLicenseInput{
		// LicenseArn: *string, // Required
		// SourceVersion: *string, // Required
	}

	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerSourceVersion) > 0 {
		input.SourceVersion = aws.String(_licensemanagerSourceVersion)
	}

	if resp, err := client.DeleteLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a license asset group.
func licensemanager_DeleteLicenseAssetGroup(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteLicenseAssetGroupInput{
		// LicenseAssetGroupArn: *string, // Required
	}

	if len(_licensemanagerLicenseAssetGroupArn) > 0 {
		input.LicenseAssetGroupArn = aws.String(_licensemanagerLicenseAssetGroupArn)
	}

	if resp, err := client.DeleteLicenseAssetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a license asset ruleset.
func licensemanager_DeleteLicenseAssetRuleset(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteLicenseAssetRulesetInput{
		// LicenseAssetRulesetArn: *string, // Required
	}

	if len(_licensemanagerLicenseAssetRulesetArn) > 0 {
		input.LicenseAssetRulesetArn = aws.String(_licensemanagerLicenseAssetRulesetArn)
	}

	if resp, err := client.DeleteLicenseAssetRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified license configuration.
// You cannot delete a license configuration that is in use.
func licensemanager_DeleteLicenseConfiguration(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteLicenseConfigurationInput{
		// LicenseConfigurationArn: *string, // Required
	}

	if len(_licensemanagerLicenseConfigurationArn) > 0 {
		input.LicenseConfigurationArn = aws.String(_licensemanagerLicenseConfigurationArn)
	}

	if resp, err := client.DeleteLicenseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified report generator.
// This action deletes the report generator, which stops it from generating future
// reports. The action cannot be reversed. It has no effect on the previous reports
// from this generator.
func licensemanager_DeleteLicenseManagerReportGenerator(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteLicenseManagerReportGeneratorInput{
		// LicenseManagerReportGeneratorArn: *string, // Required
	}

	if len(_licensemanagerLicenseManagerReportGeneratorArn) > 0 {
		input.LicenseManagerReportGeneratorArn = aws.String(_licensemanagerLicenseManagerReportGeneratorArn)
	}

	if resp, err := client.DeleteLicenseManagerReportGenerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified token. Must be called in the license home Region.
func licensemanager_DeleteToken(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.DeleteTokenInput{
		// TokenId: *string, // Required
	}

	if len(_licensemanagerTokenId) > 0 {
		input.TokenId = aws.String(_licensemanagerTokenId)
	}

	if resp, err := client.DeleteToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Extends the expiration date for license consumption.
func licensemanager_ExtendLicenseConsumption(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ExtendLicenseConsumptionInput{
		// LicenseConsumptionToken: *string, // Required
	}

	if len(_licensemanagerLicenseConsumptionToken) > 0 {
		input.LicenseConsumptionToken = aws.String(_licensemanagerLicenseConsumptionToken)
	}
	if len(_licensemanagerDryRun) > 0 {
		if err := assignInputField(input, "DryRun", _licensemanagerDryRun); err != nil {
			log.Errorf("invalid --dry-run: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExtendLicenseConsumption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a temporary access token to use with AssumeRoleWithWebIdentity. Access
// tokens are valid for one hour.
func licensemanager_GetAccessToken(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetAccessTokenInput{
		// Token: *string, // Required
	}

	if len(_licensemanagerToken) > 0 {
		input.Token = aws.String(_licensemanagerToken)
	}
	if len(_licensemanagerTokenProperties) > 0 {
		input.TokenProperties = append([]string(nil), _licensemanagerTokenProperties...)
	}

	if resp, err := client.GetAccessToken(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about the specified grant.
func licensemanager_GetGrant(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetGrantInput{
		// GrantArn: *string, // Required
	}

	if len(_licensemanagerGrantArn) > 0 {
		input.GrantArn = aws.String(_licensemanagerGrantArn)
	}
	if len(_licensemanagerVersion) > 0 {
		input.Version = aws.String(_licensemanagerVersion)
	}

	if resp, err := client.GetGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about the specified license.
func licensemanager_GetLicense(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseInput{
		// LicenseArn: *string, // Required
	}

	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerVersion) > 0 {
		input.Version = aws.String(_licensemanagerVersion)
	}

	if resp, err := client.GetLicense(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a license asset group.
func licensemanager_GetLicenseAssetGroup(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseAssetGroupInput{
		// LicenseAssetGroupArn: *string, // Required
	}

	if len(_licensemanagerLicenseAssetGroupArn) > 0 {
		input.LicenseAssetGroupArn = aws.String(_licensemanagerLicenseAssetGroupArn)
	}

	if resp, err := client.GetLicenseAssetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets a license asset ruleset.
func licensemanager_GetLicenseAssetRuleset(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseAssetRulesetInput{
		// LicenseAssetRulesetArn: *string, // Required
	}

	if len(_licensemanagerLicenseAssetRulesetArn) > 0 {
		input.LicenseAssetRulesetArn = aws.String(_licensemanagerLicenseAssetRulesetArn)
	}

	if resp, err := client.GetLicenseAssetRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about the specified license configuration.
func licensemanager_GetLicenseConfiguration(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseConfigurationInput{
		// LicenseConfigurationArn: *string, // Required
	}

	if len(_licensemanagerLicenseConfigurationArn) > 0 {
		input.LicenseConfigurationArn = aws.String(_licensemanagerLicenseConfigurationArn)
	}

	if resp, err := client.GetLicenseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified license type conversion task.
func licensemanager_GetLicenseConversionTask(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseConversionTaskInput{
		// LicenseConversionTaskId: *string, // Required
	}

	if len(_licensemanagerLicenseConversionTaskId) > 0 {
		input.LicenseConversionTaskId = aws.String(_licensemanagerLicenseConversionTaskId)
	}

	if resp, err := client.GetLicenseConversionTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified report generator.
func licensemanager_GetLicenseManagerReportGenerator(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseManagerReportGeneratorInput{
		// LicenseManagerReportGeneratorArn: *string, // Required
	}

	if len(_licensemanagerLicenseManagerReportGeneratorArn) > 0 {
		input.LicenseManagerReportGeneratorArn = aws.String(_licensemanagerLicenseManagerReportGeneratorArn)
	}

	if resp, err := client.GetLicenseManagerReportGenerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets detailed information about the usage of the specified license.
func licensemanager_GetLicenseUsage(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetLicenseUsageInput{
		// LicenseArn: *string, // Required
	}

	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}

	if resp, err := client.GetLicenseUsage(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the License Manager settings for the current Region.
func licensemanager_GetServiceSettings(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.GetServiceSettingsInput{}

	if resp, err := client.GetServiceSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists assets for a license asset group.
func licensemanager_ListAssetsForLicenseAssetGroup(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListAssetsForLicenseAssetGroupInput{
		// AssetType: *string, // Required
		// LicenseAssetGroupArn: *string, // Required
	}

	if len(_licensemanagerAssetType) > 0 {
		input.AssetType = aws.String(_licensemanagerAssetType)
	}
	if len(_licensemanagerLicenseAssetGroupArn) > 0 {
		input.LicenseAssetGroupArn = aws.String(_licensemanagerLicenseAssetGroupArn)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListAssetsForLicenseAssetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the resource associations for the specified license configuration.
// Resource associations need not consume licenses from a license configuration.
// For example, an AMI or a stopped instance might not consume a license (depending
// on the license rules).
func licensemanager_ListAssociationsForLicenseConfiguration(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListAssociationsForLicenseConfigurationInput{
		// LicenseConfigurationArn: *string, // Required
	}

	if len(_licensemanagerLicenseConfigurationArn) > 0 {
		input.LicenseConfigurationArn = aws.String(_licensemanagerLicenseConfigurationArn)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListAssociationsForLicenseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the grants distributed for the specified license.
func licensemanager_ListDistributedGrants(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListDistributedGrantsInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerGrantArns) > 0 {
		input.GrantArns = append([]string(nil), _licensemanagerGrantArns...)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListDistributedGrants(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the license configuration operations that failed.
func licensemanager_ListFailuresForLicenseConfigurationOperations(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListFailuresForLicenseConfigurationOperationsInput{
		// LicenseConfigurationArn: *string, // Required
	}

	if len(_licensemanagerLicenseConfigurationArn) > 0 {
		input.LicenseConfigurationArn = aws.String(_licensemanagerLicenseConfigurationArn)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListFailuresForLicenseConfigurationOperations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists license asset groups.
func licensemanager_ListLicenseAssetGroups(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseAssetGroupsInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseAssetGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists license asset rulesets.
func licensemanager_ListLicenseAssetRulesets(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseAssetRulesetsInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}
	if len(_licensemanagerShowAWSManagedLicenseAssetRulesets) > 0 {
		if err := assignInputField(input, "ShowAWSManagedLicenseAssetRulesets", _licensemanagerShowAWSManagedLicenseAssetRulesets); err != nil {
			log.Errorf("invalid --show-aws-managed-license-asset-rulesets: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListLicenseAssetRulesets(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the license configurations for your account.
func licensemanager_ListLicenseConfigurations(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseConfigurationsInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseConfigurationArns) > 0 {
		input.LicenseConfigurationArns = append([]string(nil), _licensemanagerLicenseConfigurationArns...)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists license configurations for an organization.
func licensemanager_ListLicenseConfigurationsForOrganization(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseConfigurationsForOrganizationInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseConfigurationArns) > 0 {
		input.LicenseConfigurationArns = append([]string(nil), _licensemanagerLicenseConfigurationArns...)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseConfigurationsForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the license type conversion tasks for your account.
func licensemanager_ListLicenseConversionTasks(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseConversionTasksInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseConversionTasks(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the report generators for your account.
func licensemanager_ListLicenseManagerReportGenerators(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseManagerReportGeneratorsInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseManagerReportGenerators(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes the license configurations for the specified resource.
func licensemanager_ListLicenseSpecificationsForResource(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseSpecificationsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_licensemanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerResourceArn)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseSpecificationsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all versions of the specified license.
func licensemanager_ListLicenseVersions(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicenseVersionsInput{
		// LicenseArn: *string, // Required
	}

	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenseVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the licenses for your account.
func licensemanager_ListLicenses(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListLicensesInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseArns) > 0 {
		input.LicenseArns = append([]string(nil), _licensemanagerLicenseArns...)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListLicenses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists grants that are received. Received grants are grants created while
// specifying the recipient as this Amazon Web Services account, your organization,
// or an organizational unit (OU) to which this member account belongs.
func licensemanager_ListReceivedGrants(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListReceivedGrantsInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerGrantArns) > 0 {
		input.GrantArns = append([]string(nil), _licensemanagerGrantArns...)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListReceivedGrants(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the grants received for all accounts in the organization.
func licensemanager_ListReceivedGrantsForOrganization(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListReceivedGrantsForOrganizationInput{
		// LicenseArn: *string, // Required
	}

	if len(_licensemanagerLicenseArn) > 0 {
		input.LicenseArn = aws.String(_licensemanagerLicenseArn)
	}
	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListReceivedGrantsForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists received licenses.
func licensemanager_ListReceivedLicenses(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListReceivedLicensesInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseArns) > 0 {
		input.LicenseArns = append([]string(nil), _licensemanagerLicenseArns...)
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListReceivedLicenses(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the licenses received for all accounts in the organization.
func licensemanager_ListReceivedLicensesForOrganization(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListReceivedLicensesForOrganizationInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListReceivedLicensesForOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists resources managed using Systems Manager inventory.
func licensemanager_ListResourceInventory(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListResourceInventoryInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListResourceInventory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the tags for the specified resource. For more information about tagging
// support in License Manager, see the [TagResource]operation.
//
// [TagResource]: https://docs.aws.amazon.com/license-manager/latest/APIReference/API_TagResource.html
func licensemanager_ListTagsForResource(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_licensemanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists your tokens.
func licensemanager_ListTokens(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListTokensInput{}

	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}
	if len(_licensemanagerTokenIds) > 0 {
		input.TokenIds = append([]string(nil), _licensemanagerTokenIds...)
	}

	if resp, err := client.ListTokens(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all license usage records for a license configuration, displaying license
// consumption details by resource at a selected point in time. Use this action to
// audit the current license consumption for any license inventory and
// configuration.
func licensemanager_ListUsageForLicenseConfiguration(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.ListUsageForLicenseConfigurationInput{
		// LicenseConfigurationArn: *string, // Required
	}

	if len(_licensemanagerLicenseConfigurationArn) > 0 {
		input.LicenseConfigurationArn = aws.String(_licensemanagerLicenseConfigurationArn)
	}
	if len(_licensemanagerFilters) > 0 {
		if err := assignInputField(input, "Filters", _licensemanagerFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _licensemanagerMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerNextToken) > 0 {
		input.NextToken = aws.String(_licensemanagerNextToken)
	}

	if resp, err := client.ListUsageForLicenseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects the specified grant.
func licensemanager_RejectGrant(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.RejectGrantInput{
		// GrantArn: *string, // Required
	}

	if len(_licensemanagerGrantArn) > 0 {
		input.GrantArn = aws.String(_licensemanagerGrantArn)
	}

	if resp, err := client.RejectGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tags to the specified resource. The following resources
// support tagging in License Manager:
//
// - Licenses
//
// - Grants
//
// - License configurations
//
// - Report generators
func licensemanager_TagResource(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_licensemanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerResourceArn)
	}
	if len(_licensemanagerTags) > 0 {
		if err := assignInputField(input, "Tags", _licensemanagerTags); err != nil {
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

// Removes the specified tags from the specified resource.
func licensemanager_UntagResource(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_licensemanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerResourceArn)
	}
	if len(_licensemanagerTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _licensemanagerTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a license asset group.
func licensemanager_UpdateLicenseAssetGroup(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UpdateLicenseAssetGroupInput{
		// AssociatedLicenseAssetRulesetARNs: []string, // Required
		// ClientToken: *string, // Required
		// LicenseAssetGroupArn: *string, // Required
	}

	if len(_licensemanagerAssociatedLicenseAssetRulesetARNs) > 0 {
		input.AssociatedLicenseAssetRulesetARNs = append([]string(nil), _licensemanagerAssociatedLicenseAssetRulesetARNs...)
	}
	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerLicenseAssetGroupArn) > 0 {
		input.LicenseAssetGroupArn = aws.String(_licensemanagerLicenseAssetGroupArn)
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerLicenseAssetGroupConfigurations) > 0 {
		if err := assignInputField(input, "LicenseAssetGroupConfigurations", _licensemanagerLicenseAssetGroupConfigurations); err != nil {
			log.Errorf("invalid --license-asset-group-configurations: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerName) > 0 {
		input.Name = aws.String(_licensemanagerName)
	}
	if len(_licensemanagerProperties) > 0 {
		if err := assignInputField(input, "Properties", _licensemanagerProperties); err != nil {
			log.Errorf("invalid --properties: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerStatus) > 0 {
		if err := assignInputField(input, "Status", _licensemanagerStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLicenseAssetGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a license asset ruleset.
func licensemanager_UpdateLicenseAssetRuleset(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UpdateLicenseAssetRulesetInput{
		// ClientToken: *string, // Required
		// LicenseAssetRulesetArn: *string, // Required
		// Rules: []types.LicenseAssetRule, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerLicenseAssetRulesetArn) > 0 {
		input.LicenseAssetRulesetArn = aws.String(_licensemanagerLicenseAssetRulesetArn)
	}
	if len(_licensemanagerRules) > 0 {
		if err := assignInputField(input, "Rules", _licensemanagerRules); err != nil {
			log.Errorf("invalid --rules: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerName) > 0 {
		input.Name = aws.String(_licensemanagerName)
	}

	if resp, err := client.UpdateLicenseAssetRuleset(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies the attributes of an existing license configuration.
func licensemanager_UpdateLicenseConfiguration(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UpdateLicenseConfigurationInput{
		// LicenseConfigurationArn: *string, // Required
	}

	if len(_licensemanagerLicenseConfigurationArn) > 0 {
		input.LicenseConfigurationArn = aws.String(_licensemanagerLicenseConfigurationArn)
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}
	if len(_licensemanagerDisassociateWhenNotFound) > 0 {
		if err := assignInputField(input, "DisassociateWhenNotFound", _licensemanagerDisassociateWhenNotFound); err != nil {
			log.Errorf("invalid --disassociate-when-not-found: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseConfigurationStatus) > 0 {
		if err := assignInputField(input, "LicenseConfigurationStatus", _licensemanagerLicenseConfigurationStatus); err != nil {
			log.Errorf("invalid --license-configuration-status: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseCount) > 0 {
		if err := assignInputField(input, "LicenseCount", _licensemanagerLicenseCount); err != nil {
			log.Errorf("invalid --license-count: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseCountHardLimit) > 0 {
		if err := assignInputField(input, "LicenseCountHardLimit", _licensemanagerLicenseCountHardLimit); err != nil {
			log.Errorf("invalid --license-count-hard-limit: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseExpiry) > 0 {
		if err := assignInputField(input, "LicenseExpiry", _licensemanagerLicenseExpiry); err != nil {
			log.Errorf("invalid --license-expiry: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerLicenseRules) > 0 {
		input.LicenseRules = append([]string(nil), _licensemanagerLicenseRules...)
	}
	if len(_licensemanagerName) > 0 {
		input.Name = aws.String(_licensemanagerName)
	}
	if len(_licensemanagerProductInformationList) > 0 {
		if err := assignInputField(input, "ProductInformationList", _licensemanagerProductInformationList); err != nil {
			log.Errorf("invalid --product-information-list: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLicenseConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a report generator.
// After you make changes to a report generator, it starts generating new reports
// within 60 minutes of being updated.
func licensemanager_UpdateLicenseManagerReportGenerator(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UpdateLicenseManagerReportGeneratorInput{
		// ClientToken: *string, // Required
		// LicenseManagerReportGeneratorArn: *string, // Required
		// ReportContext: *types.ReportContext, // Required
		// ReportFrequency: *types.ReportFrequency, // Required
		// ReportGeneratorName: *string, // Required
		// Type: []types.ReportType, // Required
	}

	if len(_licensemanagerClientToken) > 0 {
		input.ClientToken = aws.String(_licensemanagerClientToken)
	}
	if len(_licensemanagerLicenseManagerReportGeneratorArn) > 0 {
		input.LicenseManagerReportGeneratorArn = aws.String(_licensemanagerLicenseManagerReportGeneratorArn)
	}
	if len(_licensemanagerReportContext) > 0 {
		if err := assignInputField(input, "ReportContext", _licensemanagerReportContext); err != nil {
			log.Errorf("invalid --report-context: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerReportFrequency) > 0 {
		if err := assignInputField(input, "ReportFrequency", _licensemanagerReportFrequency); err != nil {
			log.Errorf("invalid --report-frequency: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerReportGeneratorName) > 0 {
		input.ReportGeneratorName = aws.String(_licensemanagerReportGeneratorName)
	}
	if len(_licensemanagerType) > 0 {
		if err := assignInputField(input, "Type", _licensemanagerType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerDescription) > 0 {
		input.Description = aws.String(_licensemanagerDescription)
	}

	if resp, err := client.UpdateLicenseManagerReportGenerator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or removes the specified license configurations for the specified Amazon
// Web Services resource.
//
// You can update the license specifications of AMIs, instances, and hosts. You
// cannot update the license specifications for launch templates and CloudFormation
// templates, as they send license configurations to the operation that creates the
// resource.
func licensemanager_UpdateLicenseSpecificationsForResource(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UpdateLicenseSpecificationsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_licensemanagerResourceArn) > 0 {
		input.ResourceArn = aws.String(_licensemanagerResourceArn)
	}
	if len(_licensemanagerAddLicenseSpecifications) > 0 {
		if err := assignInputField(input, "AddLicenseSpecifications", _licensemanagerAddLicenseSpecifications); err != nil {
			log.Errorf("invalid --add-license-specifications: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerRemoveLicenseSpecifications) > 0 {
		if err := assignInputField(input, "RemoveLicenseSpecifications", _licensemanagerRemoveLicenseSpecifications); err != nil {
			log.Errorf("invalid --remove-license-specifications: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateLicenseSpecificationsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates License Manager settings for the current Region.
func licensemanager_UpdateServiceSettings(cfg aws.Config, client *licensemanager.Client) {
	input := &licensemanager.UpdateServiceSettingsInput{}

	if len(_licensemanagerEnableCrossAccountsDiscovery) > 0 {
		if err := assignInputField(input, "EnableCrossAccountsDiscovery", _licensemanagerEnableCrossAccountsDiscovery); err != nil {
			log.Errorf("invalid --enable-cross-accounts-discovery: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerEnabledDiscoverySourceRegions) > 0 {
		input.EnabledDiscoverySourceRegions = append([]string(nil), _licensemanagerEnabledDiscoverySourceRegions...)
	}
	if len(_licensemanagerOrganizationConfiguration) > 0 {
		if err := assignInputField(input, "OrganizationConfiguration", _licensemanagerOrganizationConfiguration); err != nil {
			log.Errorf("invalid --organization-configuration: %s", err.Error())
			return
		}
	}
	if len(_licensemanagerS3BucketArn) > 0 {
		input.S3BucketArn = aws.String(_licensemanagerS3BucketArn)
	}
	if len(_licensemanagerSnsTopicArn) > 0 {
		input.SnsTopicArn = aws.String(_licensemanagerSnsTopicArn)
	}

	if resp, err := client.UpdateServiceSettings(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_licensemanagerCmd)
	_licensemanagerCmd.Flags().SortFlags = false

	_licensemanagerCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_licensemanagerCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_licensemanagerCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerAddLicenseSpecifications, "add-license-specifications", "", "", "Add License Specifications")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerAllowedOperations, "allowed-operations", "", "", "Allowed Operations")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerAssetType, "asset-type", "", "", "Asset Type")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerAssociatedLicenseAssetRulesetARNs, "associated-license-asset-ruleset-arns", "", nil, "Associated License Asset Ruleset Arns")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerBeneficiary, "beneficiary", "", "", "Beneficiary")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerCheckoutMetadata, "checkout-metadata", "", "", "Checkout Metadata")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerCheckoutType, "checkout-type", "", "", "Checkout Type")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerClientToken, "client-token", "", "", "Client Token")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerConsumptionConfiguration, "consumption-configuration", "", "", "Consumption Configuration")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerDescription, "description", "", "", "Description")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerDestinationLicenseContext, "destination-license-context", "", "", "Destination License Context")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerDigitalSignatureMethod, "digital-signature-method", "", "", "Digital Signature Method")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerDisassociateWhenNotFound, "disassociate-when-not-found", "", "", "Disassociate When Not Found")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerDryRun, "dry-run", "", "", "Dry Run")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerEnableCrossAccountsDiscovery, "enable-cross-accounts-discovery", "", "", "Enable Cross Accounts Discovery")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerEnabledDiscoverySourceRegions, "enabled-discovery-source-regions", "", nil, "Enabled Discovery Source Regions")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerEntitlements, "entitlements", "", "", "Entitlements")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerExpirationInDays, "expiration-in-days", "", "", "Expiration In Days")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerFilters, "filters", "", "", "Filters")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerGrantArn, "grant-arn", "", "", "Grant ARN")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerGrantArns, "grant-arns", "", nil, "Grant Arns")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerGrantName, "grant-name", "", "", "Grant Name")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerHomeRegion, "home-region", "", "", "Home Region")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerIssuer, "issuer", "", "", "Issuer")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerKeyFingerprint, "key-fingerprint", "", "", "Key Fingerprint")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseArn, "license-arn", "", "", "License ARN")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerLicenseArns, "license-arns", "", nil, "License Arns")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseAssetGroupArn, "license-asset-group-arn", "", "", "License Asset Group ARN")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseAssetGroupConfigurations, "license-asset-group-configurations", "", "", "License Asset Group Configurations")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseAssetRulesetArn, "license-asset-ruleset-arn", "", "", "License Asset Ruleset ARN")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseConfigurationArn, "license-configuration-arn", "", "", "License Configuration ARN")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerLicenseConfigurationArns, "license-configuration-arns", "", nil, "License Configuration Arns")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseConfigurationStatus, "license-configuration-status", "", "", "License Configuration Status")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseConsumptionToken, "license-consumption-token", "", "", "License Consumption Token")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseConversionTaskId, "license-conversion-task-id", "", "", "License Conversion Task ID")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseCount, "license-count", "", "", "License Count")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseCountHardLimit, "license-count-hard-limit", "", "", "License Count Hard Limit")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseCountingType, "license-counting-type", "", "", "License Counting Type")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseExpiry, "license-expiry", "", "", "License Expiry")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseManagerReportGeneratorArn, "license-manager-report-generator-arn", "", "", "License Manager Report Generator ARN")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseMetadata, "license-metadata", "", "", "License Metadata")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerLicenseName, "license-name", "", "", "License Name")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerLicenseRules, "license-rules", "", nil, "License Rules")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerMaxResults, "max-results", "", "", "Max Results")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerName, "name", "", "", "Name")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerNextToken, "next-token", "", "", "Next Token")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerNodeId, "node-id", "", "", "Node ID")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerOptions, "options", "", "", "Options")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerOrganizationConfiguration, "organization-configuration", "", "", "Organization Configuration")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerPrincipals, "principals", "", nil, "Principals")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerProductInformationList, "product-information-list", "", "", "Product Information List")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerProductName, "product-name", "", "", "Product Name")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerProductSKU, "product-sku", "", "", "Product Sku")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerProperties, "properties", "", "", "Properties")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerRemoveLicenseSpecifications, "remove-license-specifications", "", "", "Remove License Specifications")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerReportContext, "report-context", "", "", "Report Context")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerReportFrequency, "report-frequency", "", "", "Report Frequency")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerReportGeneratorName, "report-generator-name", "", "", "Report Generator Name")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerResourceArn, "resource-arn", "", "", "Resource ARN")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerRoleArns, "role-arns", "", nil, "Role Arns")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerRules, "rules", "", "", "Rules")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerS3BucketArn, "s3-bucket-arn", "", "", "S3 Bucket ARN")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerShowAWSManagedLicenseAssetRulesets, "show-aws-managed-license-asset-rulesets", "", "", "Show AWS Managed License Asset Rulesets")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerSnsTopicArn, "sns-topic-arn", "", "", "SNS Topic ARN")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerSourceLicenseContext, "source-license-context", "", "", "Source License Context")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerSourceVersion, "source-version", "", "", "Source Version")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerStatus, "status", "", "", "Status")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerStatusReason, "status-reason", "", "", "Status Reason")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerTagKeys, "tag-keys", "", nil, "Tag Keys")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerTags, "tags", "", "", "Tags")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerToken, "token", "", "", "Token")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerTokenId, "token-id", "", "", "Token ID")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerTokenIds, "token-ids", "", nil, "Token Ids")
	_licensemanagerCmd.Flags().StringSliceVarP(&_licensemanagerTokenProperties, "token-properties", "", nil, "Token Properties")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerType, "type", "", "", "Type")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerValidity, "validity", "", "", "Validity")
	_licensemanagerCmd.Flags().StringVarP(&_licensemanagerVersion, "version", "", "", "Version")

	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerAcceptGrant, "accept-grant", "", false, "Accept Grant")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCheckInLicense, "check-in-license", "", false, "Check In License")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCheckoutBorrowLicense, "checkout-borrow-license", "", false, "Checkout Borrow License")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCheckoutLicense, "checkout-license", "", false, "Checkout License")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateGrant, "create-grant", "", false, "Create Grant")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateGrantVersion, "create-grant-version", "", false, "Create Grant Version")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicense, "create-license", "", false, "Create License")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicenseAssetGroup, "create-license-asset-group", "", false, "Create License Asset Group")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicenseAssetRuleset, "create-license-asset-ruleset", "", false, "Create License Asset Ruleset")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicenseConfiguration, "create-license-configuration", "", false, "Create License Configuration")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicenseConversionTaskForResource, "create-license-conversion-task-for-resource", "", false, "Create License Conversion Task For Resource")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicenseManagerReportGenerator, "create-license-manager-report-generator", "", false, "Create License Manager Report Generator")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateLicenseVersion, "create-license-version", "", false, "Create License Version")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerCreateToken, "create-token", "", false, "Create Token")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteGrant, "delete-grant", "", false, "Delete Grant")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteLicense, "delete-license", "", false, "Delete License")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteLicenseAssetGroup, "delete-license-asset-group", "", false, "Delete License Asset Group")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteLicenseAssetRuleset, "delete-license-asset-ruleset", "", false, "Delete License Asset Ruleset")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteLicenseConfiguration, "delete-license-configuration", "", false, "Delete License Configuration")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteLicenseManagerReportGenerator, "delete-license-manager-report-generator", "", false, "Delete License Manager Report Generator")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerDeleteToken, "delete-token", "", false, "Delete Token")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerExtendLicenseConsumption, "extend-license-consumption", "", false, "Extend License Consumption")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetAccessToken, "get-access-token", "", false, "Get Access Token")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetGrant, "get-grant", "", false, "Get Grant")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicense, "get-license", "", false, "Get License")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicenseAssetGroup, "get-license-asset-group", "", false, "Get License Asset Group")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicenseAssetRuleset, "get-license-asset-ruleset", "", false, "Get License Asset Ruleset")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicenseConfiguration, "get-license-configuration", "", false, "Get License Configuration")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicenseConversionTask, "get-license-conversion-task", "", false, "Get License Conversion Task")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicenseManagerReportGenerator, "get-license-manager-report-generator", "", false, "Get License Manager Report Generator")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetLicenseUsage, "get-license-usage", "", false, "Get License Usage")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerGetServiceSettings, "get-service-settings", "", false, "Get Service Settings")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListAssetsForLicenseAssetGroup, "list-assets-for-license-asset-group", "", false, "List Assets For License Asset Group")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListAssociationsForLicenseConfiguration, "list-associations-for-license-configuration", "", false, "List Associations For License Configuration")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListDistributedGrants, "list-distributed-grants", "", false, "List Distributed Grants")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListFailuresForLicenseConfigurationOperations, "list-failures-for-license-configuration-operations", "", false, "List Failures For License Configuration Operations")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseAssetGroups, "list-license-asset-groups", "", false, "List License Asset Groups")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseAssetRulesets, "list-license-asset-rulesets", "", false, "List License Asset Rulesets")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseConfigurations, "list-license-configurations", "", false, "List License Configurations")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseConfigurationsForOrganization, "list-license-configurations-for-organization", "", false, "List License Configurations For Organization")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseConversionTasks, "list-license-conversion-tasks", "", false, "List License Conversion Tasks")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseManagerReportGenerators, "list-license-manager-report-generators", "", false, "List License Manager Report Generators")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseSpecificationsForResource, "list-license-specifications-for-resource", "", false, "List License Specifications For Resource")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenseVersions, "list-license-versions", "", false, "List License Versions")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListLicenses, "list-licenses", "", false, "List Licenses")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListReceivedGrants, "list-received-grants", "", false, "List Received Grants")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListReceivedGrantsForOrganization, "list-received-grants-for-organization", "", false, "List Received Grants For Organization")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListReceivedLicenses, "list-received-licenses", "", false, "List Received Licenses")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListReceivedLicensesForOrganization, "list-received-licenses-for-organization", "", false, "List Received Licenses For Organization")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListResourceInventory, "list-resource-inventory", "", false, "List Resource Inventory")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListTokens, "list-tokens", "", false, "List Tokens")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerListUsageForLicenseConfiguration, "list-usage-for-license-configuration", "", false, "List Usage For License Configuration")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerRejectGrant, "reject-grant", "", false, "Reject Grant")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerTagResource, "tag-resource", "", false, "Tag Resource")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUntagResource, "untag-resource", "", false, "Untag Resource")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUpdateLicenseAssetGroup, "update-license-asset-group", "", false, "Update License Asset Group")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUpdateLicenseAssetRuleset, "update-license-asset-ruleset", "", false, "Update License Asset Ruleset")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUpdateLicenseConfiguration, "update-license-configuration", "", false, "Update License Configuration")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUpdateLicenseManagerReportGenerator, "update-license-manager-report-generator", "", false, "Update License Manager Report Generator")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUpdateLicenseSpecificationsForResource, "update-license-specifications-for-resource", "", false, "Update License Specifications For Resource")
	_licensemanagerCmd.Flags().BoolVarP(&_licensemanagerUpdateServiceSettings, "update-service-settings", "", false, "Update Service Settings")

}
