package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// servicecatalogCmd represents the servicecatalog command
var _servicecatalogCmd = &cobra.Command{
	Use:   "servicecatalog",
	Short: "AWS servicecatalog CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := servicecatalog.NewFromConfig(cfg)
		if _servicecatalogAcceptPortfolioShare {
			servicecatalog_AcceptPortfolioShare(cfg, client)
			return
		}
		if _servicecatalogAssociateBudgetWithResource {
			servicecatalog_AssociateBudgetWithResource(cfg, client)
			return
		}
		if _servicecatalogAssociatePrincipalWithPortfolio {
			servicecatalog_AssociatePrincipalWithPortfolio(cfg, client)
			return
		}
		if _servicecatalogAssociateProductWithPortfolio {
			servicecatalog_AssociateProductWithPortfolio(cfg, client)
			return
		}
		if _servicecatalogAssociateServiceActionWithProvisioningArtifact {
			servicecatalog_AssociateServiceActionWithProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogAssociateTagOptionWithResource {
			servicecatalog_AssociateTagOptionWithResource(cfg, client)
			return
		}
		if _servicecatalogBatchAssociateServiceActionWithProvisioningArtifact {
			servicecatalog_BatchAssociateServiceActionWithProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogBatchDisassociateServiceActionFromProvisioningArtifact {
			servicecatalog_BatchDisassociateServiceActionFromProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogCopyProduct {
			servicecatalog_CopyProduct(cfg, client)
			return
		}
		if _servicecatalogCreateConstraint {
			servicecatalog_CreateConstraint(cfg, client)
			return
		}
		if _servicecatalogCreatePortfolio {
			servicecatalog_CreatePortfolio(cfg, client)
			return
		}
		if _servicecatalogCreatePortfolioShare {
			servicecatalog_CreatePortfolioShare(cfg, client)
			return
		}
		if _servicecatalogCreateProduct {
			servicecatalog_CreateProduct(cfg, client)
			return
		}
		if _servicecatalogCreateProvisionedProductPlan {
			servicecatalog_CreateProvisionedProductPlan(cfg, client)
			return
		}
		if _servicecatalogCreateProvisioningArtifact {
			servicecatalog_CreateProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogCreateServiceAction {
			servicecatalog_CreateServiceAction(cfg, client)
			return
		}
		if _servicecatalogCreateTagOption {
			servicecatalog_CreateTagOption(cfg, client)
			return
		}
		if _servicecatalogDeleteConstraint {
			servicecatalog_DeleteConstraint(cfg, client)
			return
		}
		if _servicecatalogDeletePortfolio {
			servicecatalog_DeletePortfolio(cfg, client)
			return
		}
		if _servicecatalogDeletePortfolioShare {
			servicecatalog_DeletePortfolioShare(cfg, client)
			return
		}
		if _servicecatalogDeleteProduct {
			servicecatalog_DeleteProduct(cfg, client)
			return
		}
		if _servicecatalogDeleteProvisionedProductPlan {
			servicecatalog_DeleteProvisionedProductPlan(cfg, client)
			return
		}
		if _servicecatalogDeleteProvisioningArtifact {
			servicecatalog_DeleteProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogDeleteServiceAction {
			servicecatalog_DeleteServiceAction(cfg, client)
			return
		}
		if _servicecatalogDeleteTagOption {
			servicecatalog_DeleteTagOption(cfg, client)
			return
		}
		if _servicecatalogDescribeConstraint {
			servicecatalog_DescribeConstraint(cfg, client)
			return
		}
		if _servicecatalogDescribeCopyProductStatus {
			servicecatalog_DescribeCopyProductStatus(cfg, client)
			return
		}
		if _servicecatalogDescribePortfolio {
			servicecatalog_DescribePortfolio(cfg, client)
			return
		}
		if _servicecatalogDescribePortfolioShareStatus {
			servicecatalog_DescribePortfolioShareStatus(cfg, client)
			return
		}
		if _servicecatalogDescribePortfolioShares {
			servicecatalog_DescribePortfolioShares(cfg, client)
			return
		}
		if _servicecatalogDescribeProduct {
			servicecatalog_DescribeProduct(cfg, client)
			return
		}
		if _servicecatalogDescribeProductAsAdmin {
			servicecatalog_DescribeProductAsAdmin(cfg, client)
			return
		}
		if _servicecatalogDescribeProductView {
			servicecatalog_DescribeProductView(cfg, client)
			return
		}
		if _servicecatalogDescribeProvisionedProduct {
			servicecatalog_DescribeProvisionedProduct(cfg, client)
			return
		}
		if _servicecatalogDescribeProvisionedProductPlan {
			servicecatalog_DescribeProvisionedProductPlan(cfg, client)
			return
		}
		if _servicecatalogDescribeProvisioningArtifact {
			servicecatalog_DescribeProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogDescribeProvisioningParameters {
			servicecatalog_DescribeProvisioningParameters(cfg, client)
			return
		}
		if _servicecatalogDescribeRecord {
			servicecatalog_DescribeRecord(cfg, client)
			return
		}
		if _servicecatalogDescribeServiceAction {
			servicecatalog_DescribeServiceAction(cfg, client)
			return
		}
		if _servicecatalogDescribeServiceActionExecutionParameters {
			servicecatalog_DescribeServiceActionExecutionParameters(cfg, client)
			return
		}
		if _servicecatalogDescribeTagOption {
			servicecatalog_DescribeTagOption(cfg, client)
			return
		}
		if _servicecatalogDisableAWSOrganizationsAccess {
			servicecatalog_DisableAWSOrganizationsAccess(cfg, client)
			return
		}
		if _servicecatalogDisassociateBudgetFromResource {
			servicecatalog_DisassociateBudgetFromResource(cfg, client)
			return
		}
		if _servicecatalogDisassociatePrincipalFromPortfolio {
			servicecatalog_DisassociatePrincipalFromPortfolio(cfg, client)
			return
		}
		if _servicecatalogDisassociateProductFromPortfolio {
			servicecatalog_DisassociateProductFromPortfolio(cfg, client)
			return
		}
		if _servicecatalogDisassociateServiceActionFromProvisioningArtifact {
			servicecatalog_DisassociateServiceActionFromProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogDisassociateTagOptionFromResource {
			servicecatalog_DisassociateTagOptionFromResource(cfg, client)
			return
		}
		if _servicecatalogEnableAWSOrganizationsAccess {
			servicecatalog_EnableAWSOrganizationsAccess(cfg, client)
			return
		}
		if _servicecatalogExecuteProvisionedProductPlan {
			servicecatalog_ExecuteProvisionedProductPlan(cfg, client)
			return
		}
		if _servicecatalogExecuteProvisionedProductServiceAction {
			servicecatalog_ExecuteProvisionedProductServiceAction(cfg, client)
			return
		}
		if _servicecatalogGetAWSOrganizationsAccessStatus {
			servicecatalog_GetAWSOrganizationsAccessStatus(cfg, client)
			return
		}
		if _servicecatalogGetProvisionedProductOutputs {
			servicecatalog_GetProvisionedProductOutputs(cfg, client)
			return
		}
		if _servicecatalogImportAsProvisionedProduct {
			servicecatalog_ImportAsProvisionedProduct(cfg, client)
			return
		}
		if _servicecatalogListAcceptedPortfolioShares {
			servicecatalog_ListAcceptedPortfolioShares(cfg, client)
			return
		}
		if _servicecatalogListBudgetsForResource {
			servicecatalog_ListBudgetsForResource(cfg, client)
			return
		}
		if _servicecatalogListConstraintsForPortfolio {
			servicecatalog_ListConstraintsForPortfolio(cfg, client)
			return
		}
		if _servicecatalogListLaunchPaths {
			servicecatalog_ListLaunchPaths(cfg, client)
			return
		}
		if _servicecatalogListOrganizationPortfolioAccess {
			servicecatalog_ListOrganizationPortfolioAccess(cfg, client)
			return
		}
		if _servicecatalogListPortfolioAccess {
			servicecatalog_ListPortfolioAccess(cfg, client)
			return
		}
		if _servicecatalogListPortfolios {
			servicecatalog_ListPortfolios(cfg, client)
			return
		}
		if _servicecatalogListPortfoliosForProduct {
			servicecatalog_ListPortfoliosForProduct(cfg, client)
			return
		}
		if _servicecatalogListPrincipalsForPortfolio {
			servicecatalog_ListPrincipalsForPortfolio(cfg, client)
			return
		}
		if _servicecatalogListProvisionedProductPlans {
			servicecatalog_ListProvisionedProductPlans(cfg, client)
			return
		}
		if _servicecatalogListProvisioningArtifacts {
			servicecatalog_ListProvisioningArtifacts(cfg, client)
			return
		}
		if _servicecatalogListProvisioningArtifactsForServiceAction {
			servicecatalog_ListProvisioningArtifactsForServiceAction(cfg, client)
			return
		}
		if _servicecatalogListRecordHistory {
			servicecatalog_ListRecordHistory(cfg, client)
			return
		}
		if _servicecatalogListResourcesForTagOption {
			servicecatalog_ListResourcesForTagOption(cfg, client)
			return
		}
		if _servicecatalogListServiceActions {
			servicecatalog_ListServiceActions(cfg, client)
			return
		}
		if _servicecatalogListServiceActionsForProvisioningArtifact {
			servicecatalog_ListServiceActionsForProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogListStackInstancesForProvisionedProduct {
			servicecatalog_ListStackInstancesForProvisionedProduct(cfg, client)
			return
		}
		if _servicecatalogListTagOptions {
			servicecatalog_ListTagOptions(cfg, client)
			return
		}
		if _servicecatalogNotifyProvisionProductEngineWorkflowResult {
			servicecatalog_NotifyProvisionProductEngineWorkflowResult(cfg, client)
			return
		}
		if _servicecatalogNotifyTerminateProvisionedProductEngineWorkflowResult {
			servicecatalog_NotifyTerminateProvisionedProductEngineWorkflowResult(cfg, client)
			return
		}
		if _servicecatalogNotifyUpdateProvisionedProductEngineWorkflowResult {
			servicecatalog_NotifyUpdateProvisionedProductEngineWorkflowResult(cfg, client)
			return
		}
		if _servicecatalogProvisionProduct {
			servicecatalog_ProvisionProduct(cfg, client)
			return
		}
		if _servicecatalogRejectPortfolioShare {
			servicecatalog_RejectPortfolioShare(cfg, client)
			return
		}
		if _servicecatalogScanProvisionedProducts {
			servicecatalog_ScanProvisionedProducts(cfg, client)
			return
		}
		if _servicecatalogSearchProducts {
			servicecatalog_SearchProducts(cfg, client)
			return
		}
		if _servicecatalogSearchProductsAsAdmin {
			servicecatalog_SearchProductsAsAdmin(cfg, client)
			return
		}
		if _servicecatalogSearchProvisionedProducts {
			servicecatalog_SearchProvisionedProducts(cfg, client)
			return
		}
		if _servicecatalogTerminateProvisionedProduct {
			servicecatalog_TerminateProvisionedProduct(cfg, client)
			return
		}
		if _servicecatalogUpdateConstraint {
			servicecatalog_UpdateConstraint(cfg, client)
			return
		}
		if _servicecatalogUpdatePortfolio {
			servicecatalog_UpdatePortfolio(cfg, client)
			return
		}
		if _servicecatalogUpdatePortfolioShare {
			servicecatalog_UpdatePortfolioShare(cfg, client)
			return
		}
		if _servicecatalogUpdateProduct {
			servicecatalog_UpdateProduct(cfg, client)
			return
		}
		if _servicecatalogUpdateProvisionedProduct {
			servicecatalog_UpdateProvisionedProduct(cfg, client)
			return
		}
		if _servicecatalogUpdateProvisionedProductProperties {
			servicecatalog_UpdateProvisionedProductProperties(cfg, client)
			return
		}
		if _servicecatalogUpdateProvisioningArtifact {
			servicecatalog_UpdateProvisioningArtifact(cfg, client)
			return
		}
		if _servicecatalogUpdateServiceAction {
			servicecatalog_UpdateServiceAction(cfg, client)
			return
		}
		if _servicecatalogUpdateTagOption {
			servicecatalog_UpdateTagOption(cfg, client)
			return
		}

	},
}

var (
	_servicecatalogAcceptPortfolioShare                                   bool
	_servicecatalogAssociateBudgetWithResource                            bool
	_servicecatalogAssociatePrincipalWithPortfolio                        bool
	_servicecatalogAssociateProductWithPortfolio                          bool
	_servicecatalogAssociateServiceActionWithProvisioningArtifact         bool
	_servicecatalogAssociateTagOptionWithResource                         bool
	_servicecatalogBatchAssociateServiceActionWithProvisioningArtifact    bool
	_servicecatalogBatchDisassociateServiceActionFromProvisioningArtifact bool
	_servicecatalogCopyProduct                                            bool
	_servicecatalogCreateConstraint                                       bool
	_servicecatalogCreatePortfolio                                        bool
	_servicecatalogCreatePortfolioShare                                   bool
	_servicecatalogCreateProduct                                          bool
	_servicecatalogCreateProvisionedProductPlan                           bool
	_servicecatalogCreateProvisioningArtifact                             bool
	_servicecatalogCreateServiceAction                                    bool
	_servicecatalogCreateTagOption                                        bool
	_servicecatalogDeleteConstraint                                       bool
	_servicecatalogDeletePortfolio                                        bool
	_servicecatalogDeletePortfolioShare                                   bool
	_servicecatalogDeleteProduct                                          bool
	_servicecatalogDeleteProvisionedProductPlan                           bool
	_servicecatalogDeleteProvisioningArtifact                             bool
	_servicecatalogDeleteServiceAction                                    bool
	_servicecatalogDeleteTagOption                                        bool
	_servicecatalogDescribeConstraint                                     bool
	_servicecatalogDescribeCopyProductStatus                              bool
	_servicecatalogDescribePortfolio                                      bool
	_servicecatalogDescribePortfolioShareStatus                           bool
	_servicecatalogDescribePortfolioShares                                bool
	_servicecatalogDescribeProduct                                        bool
	_servicecatalogDescribeProductAsAdmin                                 bool
	_servicecatalogDescribeProductView                                    bool
	_servicecatalogDescribeProvisionedProduct                             bool
	_servicecatalogDescribeProvisionedProductPlan                         bool
	_servicecatalogDescribeProvisioningArtifact                           bool
	_servicecatalogDescribeProvisioningParameters                         bool
	_servicecatalogDescribeRecord                                         bool
	_servicecatalogDescribeServiceAction                                  bool
	_servicecatalogDescribeServiceActionExecutionParameters               bool
	_servicecatalogDescribeTagOption                                      bool
	_servicecatalogDisableAWSOrganizationsAccess                          bool
	_servicecatalogDisassociateBudgetFromResource                         bool
	_servicecatalogDisassociatePrincipalFromPortfolio                     bool
	_servicecatalogDisassociateProductFromPortfolio                       bool
	_servicecatalogDisassociateServiceActionFromProvisioningArtifact      bool
	_servicecatalogDisassociateTagOptionFromResource                      bool
	_servicecatalogEnableAWSOrganizationsAccess                           bool
	_servicecatalogExecuteProvisionedProductPlan                          bool
	_servicecatalogExecuteProvisionedProductServiceAction                 bool
	_servicecatalogGetAWSOrganizationsAccessStatus                        bool
	_servicecatalogGetProvisionedProductOutputs                           bool
	_servicecatalogImportAsProvisionedProduct                             bool
	_servicecatalogListAcceptedPortfolioShares                            bool
	_servicecatalogListBudgetsForResource                                 bool
	_servicecatalogListConstraintsForPortfolio                            bool
	_servicecatalogListLaunchPaths                                        bool
	_servicecatalogListOrganizationPortfolioAccess                        bool
	_servicecatalogListPortfolioAccess                                    bool
	_servicecatalogListPortfolios                                         bool
	_servicecatalogListPortfoliosForProduct                               bool
	_servicecatalogListPrincipalsForPortfolio                             bool
	_servicecatalogListProvisionedProductPlans                            bool
	_servicecatalogListProvisioningArtifacts                              bool
	_servicecatalogListProvisioningArtifactsForServiceAction              bool
	_servicecatalogListRecordHistory                                      bool
	_servicecatalogListResourcesForTagOption                              bool
	_servicecatalogListServiceActions                                     bool
	_servicecatalogListServiceActionsForProvisioningArtifact              bool
	_servicecatalogListStackInstancesForProvisionedProduct                bool
	_servicecatalogListTagOptions                                         bool
	_servicecatalogNotifyProvisionProductEngineWorkflowResult             bool
	_servicecatalogNotifyTerminateProvisionedProductEngineWorkflowResult  bool
	_servicecatalogNotifyUpdateProvisionedProductEngineWorkflowResult     bool
	_servicecatalogProvisionProduct                                       bool
	_servicecatalogRejectPortfolioShare                                   bool
	_servicecatalogScanProvisionedProducts                                bool
	_servicecatalogSearchProducts                                         bool
	_servicecatalogSearchProductsAsAdmin                                  bool
	_servicecatalogSearchProvisionedProducts                              bool
	_servicecatalogTerminateProvisionedProduct                            bool
	_servicecatalogUpdateConstraint                                       bool
	_servicecatalogUpdatePortfolio                                        bool
	_servicecatalogUpdatePortfolioShare                                   bool
	_servicecatalogUpdateProduct                                          bool
	_servicecatalogUpdateProvisionedProduct                               bool
	_servicecatalogUpdateProvisionedProductProperties                     bool
	_servicecatalogUpdateProvisioningArtifact                             bool
	_servicecatalogUpdateServiceAction                                    bool
	_servicecatalogUpdateTagOption                                        bool

	_servicecatalogAcceptLanguage                        string
	_servicecatalogAccessLevelFilter                     string
	_servicecatalogAccountId                             string
	_servicecatalogActive                                string
	_servicecatalogAddTags                               string
	_servicecatalogBudgetName                            string
	_servicecatalogCopyOptions                           string
	_servicecatalogCopyProductToken                      string
	_servicecatalogDefinition                            string
	_servicecatalogDefinitionType                        string
	_servicecatalogDescription                           string
	_servicecatalogDisplayName                           string
	_servicecatalogDistributor                           string
	_servicecatalogExecuteToken                          string
	_servicecatalogFailureReason                         string
	_servicecatalogFilters                               string
	_servicecatalogGuidance                              string
	_servicecatalogId                                    string
	_servicecatalogIdempotencyToken                      string
	_servicecatalogIgnoreErrors                          string
	_servicecatalogIncludeProvisioningArtifactParameters string
	_servicecatalogKey                                   string
	_servicecatalogName                                  string
	_servicecatalogNotificationArns                      []string
	_servicecatalogOrganizationNode                      string
	_servicecatalogOrganizationNodeType                  string
	_servicecatalogOrganizationParentId                  string
	_servicecatalogOutputKeys                            []string
	_servicecatalogOutputs                               string
	_servicecatalogOwner                                 string
	_servicecatalogPageSize                              string
	_servicecatalogPageToken                             string
	_servicecatalogParameters                            string
	_servicecatalogPathId                                string
	_servicecatalogPathName                              string
	_servicecatalogPhysicalId                            string
	_servicecatalogPlanId                                string
	_servicecatalogPlanName                              string
	_servicecatalogPlanType                              string
	_servicecatalogPortfolioId                           string
	_servicecatalogPortfolioShareToken                   string
	_servicecatalogPortfolioShareType                    string
	_servicecatalogPrincipalARN                          string
	_servicecatalogPrincipalType                         string
	_servicecatalogProductId                             string
	_servicecatalogProductName                           string
	_servicecatalogProductSource                         string
	_servicecatalogProductType                           string
	_servicecatalogProviderName                          string
	_servicecatalogProvisionProductId                    string
	_servicecatalogProvisionToken                        string
	_servicecatalogProvisionedProductId                  string
	_servicecatalogProvisionedProductName                string
	_servicecatalogProvisionedProductProperties          string
	_servicecatalogProvisioningArtifactId                string
	_servicecatalogProvisioningArtifactName              string
	_servicecatalogProvisioningArtifactParameters        string
	_servicecatalogProvisioningParameters                string
	_servicecatalogProvisioningPreferences               string
	_servicecatalogRecordId                              string
	_servicecatalogRemoveTags                            []string
	_servicecatalogResourceId                            string
	_servicecatalogResourceIdentifier                    string
	_servicecatalogResourceType                          string
	_servicecatalogRetainPhysicalResources               string
	_servicecatalogSearchFilter                          string
	_servicecatalogServiceActionAssociations             string
	_servicecatalogServiceActionId                       string
	_servicecatalogSharePrincipals                       string
	_servicecatalogShareTagOptions                       string
	_servicecatalogSortBy                                string
	_servicecatalogSortOrder                             string
	_servicecatalogSourceConnection                      string
	_servicecatalogSourcePortfolioId                     string
	_servicecatalogSourceProductArn                      string
	_servicecatalogSourceProvisioningArtifactIdentifiers string
	_servicecatalogStatus                                string
	_servicecatalogSupportDescription                    string
	_servicecatalogSupportEmail                          string
	_servicecatalogSupportUrl                            string
	_servicecatalogTagOptionId                           string
	_servicecatalogTags                                  string
	_servicecatalogTargetProductId                       string
	_servicecatalogTargetProductName                     string
	_servicecatalogTerminateToken                        string
	_servicecatalogType                                  string
	_servicecatalogUpdateToken                           string
	_servicecatalogValue                                 string
	_servicecatalogVerbose                               string
	_servicecatalogWorkflowToken                         string
)

// Accepts an offer to share the specified portfolio.
func servicecatalog_AcceptPortfolioShare(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.AcceptPortfolioShareInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPortfolioShareType) > 0 {
		if err := assignInputField(input, "PortfolioShareType", _servicecatalogPortfolioShareType); err != nil {
			log.Errorf("invalid --portfolio-share-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.AcceptPortfolioShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified budget with the specified resource.
func servicecatalog_AssociateBudgetWithResource(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.AssociateBudgetWithResourceInput{
		// BudgetName: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_servicecatalogBudgetName) > 0 {
		input.BudgetName = aws.String(_servicecatalogBudgetName)
	}
	if len(_servicecatalogResourceId) > 0 {
		input.ResourceId = aws.String(_servicecatalogResourceId)
	}

	if resp, err := client.AssociateBudgetWithResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified principal ARN with the specified portfolio.
// If you share the portfolio with principal name sharing enabled, the PrincipalARN
// association is included in the share.
//
// The PortfolioID , PrincipalARN , and PrincipalType parameters are required.
//
// You can associate a maximum of 10 Principals with a portfolio using
// PrincipalType as IAM_PATTERN .
//
// When you associate a principal with portfolio, a potential privilege escalation
// path may occur when that portfolio is then shared with other accounts. For a
// user in a recipient account who is not an Service Catalog Admin, but still has
// the ability to create Principals (Users/Groups/Roles), that user could create a
// role that matches a principal name association for the portfolio. Although this
// user may not know which principal names are associated through Service Catalog,
// they may be able to guess the user. If this potential escalation path is a
// concern, then Service Catalog recommends using PrincipalType as IAM . With this
// configuration, the PrincipalARN must already exist in the recipient account
// before it can be associated.
func servicecatalog_AssociatePrincipalWithPortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.AssociatePrincipalWithPortfolioInput{
		// PortfolioId: *string, // Required
		// PrincipalARN: *string, // Required
		// PrincipalType: types.PrincipalType, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogPrincipalARN) > 0 {
		input.PrincipalARN = aws.String(_servicecatalogPrincipalARN)
	}
	if len(_servicecatalogPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _servicecatalogPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.AssociatePrincipalWithPortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates the specified product with the specified portfolio.
// A delegated admin is authorized to invoke this command.
func servicecatalog_AssociateProductWithPortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.AssociateProductWithPortfolioInput{
		// PortfolioId: *string, // Required
		// ProductId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogSourcePortfolioId) > 0 {
		input.SourcePortfolioId = aws.String(_servicecatalogSourcePortfolioId)
	}

	if resp, err := client.AssociateProductWithPortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates a self-service action with a provisioning artifact.
func servicecatalog_AssociateServiceActionWithProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.AssociateServiceActionWithProvisioningArtifactInput{
		// ProductId: *string, // Required
		// ProvisioningArtifactId: *string, // Required
		// ServiceActionId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogServiceActionId) > 0 {
		input.ServiceActionId = aws.String(_servicecatalogServiceActionId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}

	if resp, err := client.AssociateServiceActionWithProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associate the specified TagOption with the specified portfolio or product.
func servicecatalog_AssociateTagOptionWithResource(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.AssociateTagOptionWithResourceInput{
		// ResourceId: *string, // Required
		// TagOptionId: *string, // Required
	}

	if len(_servicecatalogResourceId) > 0 {
		input.ResourceId = aws.String(_servicecatalogResourceId)
	}
	if len(_servicecatalogTagOptionId) > 0 {
		input.TagOptionId = aws.String(_servicecatalogTagOptionId)
	}

	if resp, err := client.AssociateTagOptionWithResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates multiple self-service actions with provisioning artifacts.
func servicecatalog_BatchAssociateServiceActionWithProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.BatchAssociateServiceActionWithProvisioningArtifactInput{
		// ServiceActionAssociations: []types.ServiceActionAssociation, // Required
	}

	if len(_servicecatalogServiceActionAssociations) > 0 {
		if err := assignInputField(input, "ServiceActionAssociations", _servicecatalogServiceActionAssociations); err != nil {
			log.Errorf("invalid --service-action-associations: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.BatchAssociateServiceActionWithProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a batch of self-service actions from the specified provisioning
// artifact.
func servicecatalog_BatchDisassociateServiceActionFromProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.BatchDisassociateServiceActionFromProvisioningArtifactInput{
		// ServiceActionAssociations: []types.ServiceActionAssociation, // Required
	}

	if len(_servicecatalogServiceActionAssociations) > 0 {
		if err := assignInputField(input, "ServiceActionAssociations", _servicecatalogServiceActionAssociations); err != nil {
			log.Errorf("invalid --service-action-associations: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.BatchDisassociateServiceActionFromProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Copies the specified source product to the specified target product or a new
// product.
//
// You can copy a product to the same account or another account. You can copy a
// product to the same Region or another Region. If you copy a product to another
// account, you must first share the product in a portfolio using CreatePortfolioShare.
//
// This operation is performed asynchronously. To track the progress of the
// operation, use DescribeCopyProductStatus.
func servicecatalog_CopyProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CopyProductInput{
		// IdempotencyToken: *string, // Required
		// SourceProductArn: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogSourceProductArn) > 0 {
		input.SourceProductArn = aws.String(_servicecatalogSourceProductArn)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogCopyOptions) > 0 {
		if err := assignInputField(input, "CopyOptions", _servicecatalogCopyOptions); err != nil {
			log.Errorf("invalid --copy-options: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSourceProvisioningArtifactIdentifiers) > 0 {
		if err := assignInputField(input, "SourceProvisioningArtifactIdentifiers", _servicecatalogSourceProvisioningArtifactIdentifiers); err != nil {
			log.Errorf("invalid --source-provisioning-artifact-identifiers: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogTargetProductId) > 0 {
		input.TargetProductId = aws.String(_servicecatalogTargetProductId)
	}
	if len(_servicecatalogTargetProductName) > 0 {
		input.TargetProductName = aws.String(_servicecatalogTargetProductName)
	}

	if resp, err := client.CopyProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a constraint.
// A delegated admin is authorized to invoke this command.
func servicecatalog_CreateConstraint(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreateConstraintInput{
		// IdempotencyToken: *string, // Required
		// Parameters: *string, // Required
		// PortfolioId: *string, // Required
		// ProductId: *string, // Required
		// Type: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogParameters) > 0 {
		input.Parameters = aws.String(_servicecatalogParameters)
	}
	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogType) > 0 {
		input.Type = aws.String(_servicecatalogType)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}

	if resp, err := client.CreateConstraint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a portfolio.
// A delegated admin is authorized to invoke this command.
func servicecatalog_CreatePortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreatePortfolioInput{
		// DisplayName: *string, // Required
		// IdempotencyToken: *string, // Required
		// ProviderName: *string, // Required
	}

	if len(_servicecatalogDisplayName) > 0 {
		input.DisplayName = aws.String(_servicecatalogDisplayName)
	}
	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogProviderName) > 0 {
		input.ProviderName = aws.String(_servicecatalogProviderName)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Shares the specified portfolio with the specified account or organization node.
// Shares to an organization node can only be created by the management account of
// an organization or by a delegated administrator. You can share portfolios to an
// organization, an organizational unit, or a specific account.
//
// Note that if a delegated admin is de-registered, they can no longer create
// portfolio shares.
//
// AWSOrganizationsAccess must be enabled in order to create a portfolio share to
// an organization node.
//
// You can't share a shared resource, including portfolios that contain a shared
// product.
//
// If the portfolio share with the specified account or organization node already
// exists, this action will have no effect and will not return an error. To update
// an existing share, you must use the UpdatePortfolioShare API instead.
//
// When you associate a principal with portfolio, a potential privilege escalation
// path may occur when that portfolio is then shared with other accounts. For a
// user in a recipient account who is not an Service Catalog Admin, but still has
// the ability to create Principals (Users/Groups/Roles), that user could create a
// role that matches a principal name association for the portfolio. Although this
// user may not know which principal names are associated through Service Catalog,
// they may be able to guess the user. If this potential escalation path is a
// concern, then Service Catalog recommends using PrincipalType as IAM . With this
// configuration, the PrincipalARN must already exist in the recipient account
// before it can be associated.
func servicecatalog_CreatePortfolioShare(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreatePortfolioShareInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccountId) > 0 {
		input.AccountId = aws.String(_servicecatalogAccountId)
	}
	if len(_servicecatalogOrganizationNode) > 0 {
		if err := assignInputField(input, "OrganizationNode", _servicecatalogOrganizationNode); err != nil {
			log.Errorf("invalid --organization-node: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSharePrincipals) > 0 {
		if err := assignInputField(input, "SharePrincipals", _servicecatalogSharePrincipals); err != nil {
			log.Errorf("invalid --share-principals: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogShareTagOptions) > 0 {
		if err := assignInputField(input, "ShareTagOptions", _servicecatalogShareTagOptions); err != nil {
			log.Errorf("invalid --share-tag-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePortfolioShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a product.
// A delegated admin is authorized to invoke this command.
//
// The user or role that performs this operation must have the
// cloudformation:GetTemplate IAM policy permission. This policy permission is
// required when using the ImportFromPhysicalId template source in the information
// data section.
func servicecatalog_CreateProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreateProductInput{
		// IdempotencyToken: *string, // Required
		// Name: *string, // Required
		// Owner: *string, // Required
		// ProductType: types.ProductType, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}
	if len(_servicecatalogOwner) > 0 {
		input.Owner = aws.String(_servicecatalogOwner)
	}
	if len(_servicecatalogProductType) > 0 {
		if err := assignInputField(input, "ProductType", _servicecatalogProductType); err != nil {
			log.Errorf("invalid --product-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogDistributor) > 0 {
		input.Distributor = aws.String(_servicecatalogDistributor)
	}
	if len(_servicecatalogProvisioningArtifactParameters) > 0 {
		if err := assignInputField(input, "ProvisioningArtifactParameters", _servicecatalogProvisioningArtifactParameters); err != nil {
			log.Errorf("invalid --provisioning-artifact-parameters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSourceConnection) > 0 {
		if err := assignInputField(input, "SourceConnection", _servicecatalogSourceConnection); err != nil {
			log.Errorf("invalid --source-connection: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSupportDescription) > 0 {
		input.SupportDescription = aws.String(_servicecatalogSupportDescription)
	}
	if len(_servicecatalogSupportEmail) > 0 {
		input.SupportEmail = aws.String(_servicecatalogSupportEmail)
	}
	if len(_servicecatalogSupportUrl) > 0 {
		input.SupportUrl = aws.String(_servicecatalogSupportUrl)
	}
	if len(_servicecatalogTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a plan.
// A plan includes the list of resources to be created (when provisioning a new
// product) or modified (when updating a provisioned product) when the plan is
// executed.
//
// You can create one plan for each provisioned product. To create a plan for an
// existing provisioned product, the product status must be AVAILABLE or TAINTED.
//
// To view the resource changes in the change set, use DescribeProvisionedProductPlan. To create or modify the
// provisioned product, use ExecuteProvisionedProductPlan.
func servicecatalog_CreateProvisionedProductPlan(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreateProvisionedProductPlanInput{
		// IdempotencyToken: *string, // Required
		// PlanName: *string, // Required
		// PlanType: types.ProvisionedProductPlanType, // Required
		// ProductId: *string, // Required
		// ProvisionedProductName: *string, // Required
		// ProvisioningArtifactId: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogPlanName) > 0 {
		input.PlanName = aws.String(_servicecatalogPlanName)
	}
	if len(_servicecatalogPlanType) > 0 {
		if err := assignInputField(input, "PlanType", _servicecatalogPlanType); err != nil {
			log.Errorf("invalid --plan-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisionedProductName) > 0 {
		input.ProvisionedProductName = aws.String(_servicecatalogProvisionedProductName)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogNotificationArns) > 0 {
		input.NotificationArns = append([]string(nil), _servicecatalogNotificationArns...)
	}
	if len(_servicecatalogPathId) > 0 {
		input.PathId = aws.String(_servicecatalogPathId)
	}
	if len(_servicecatalogProvisioningParameters) > 0 {
		if err := assignInputField(input, "ProvisioningParameters", _servicecatalogProvisioningParameters); err != nil {
			log.Errorf("invalid --provisioning-parameters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProvisionedProductPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a provisioning artifact (also known as a version) for the specified
// product.
//
// You cannot create a provisioning artifact for a product that was shared with
// you.
//
// The user or role that performs this operation must have the
// cloudformation:GetTemplate IAM policy permission. This policy permission is
// required when using the ImportFromPhysicalId template source in the information
// data section.
func servicecatalog_CreateProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreateProvisioningArtifactInput{
		// IdempotencyToken: *string, // Required
		// Parameters: *types.ProvisioningArtifactProperties, // Required
		// ProductId: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogParameters) > 0 {
		if err := assignInputField(input, "Parameters", _servicecatalogParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.CreateProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a self-service action.
func servicecatalog_CreateServiceAction(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreateServiceActionInput{
		// Definition: map[string]string, // Required
		// DefinitionType: types.ServiceActionDefinitionType, // Required
		// IdempotencyToken: *string, // Required
		// Name: *string, // Required
	}

	if len(_servicecatalogDefinition) > 0 {
		if err := assignInputField(input, "Definition", _servicecatalogDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogDefinitionType) > 0 {
		if err := assignInputField(input, "DefinitionType", _servicecatalogDefinitionType); err != nil {
			log.Errorf("invalid --definition-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}

	if resp, err := client.CreateServiceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a TagOption.
func servicecatalog_CreateTagOption(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.CreateTagOptionInput{
		// Key: *string, // Required
		// Value: *string, // Required
	}

	if len(_servicecatalogKey) > 0 {
		input.Key = aws.String(_servicecatalogKey)
	}
	if len(_servicecatalogValue) > 0 {
		input.Value = aws.String(_servicecatalogValue)
	}

	if resp, err := client.CreateTagOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified constraint.
// A delegated admin is authorized to invoke this command.
func servicecatalog_DeleteConstraint(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeleteConstraintInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DeleteConstraint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified portfolio.
// You cannot delete a portfolio if it was shared with you or if it has associated
// products, users, constraints, or shared accounts.
//
// A delegated admin is authorized to invoke this command.
func servicecatalog_DeletePortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeletePortfolioInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DeletePortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Stops sharing the specified portfolio with the specified account or
// organization node. Shares to an organization node can only be deleted by the
// management account of an organization or by a delegated administrator.
//
// Note that if a delegated admin is de-registered, portfolio shares created from
// that account are removed.
func servicecatalog_DeletePortfolioShare(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeletePortfolioShareInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccountId) > 0 {
		input.AccountId = aws.String(_servicecatalogAccountId)
	}
	if len(_servicecatalogOrganizationNode) > 0 {
		if err := assignInputField(input, "OrganizationNode", _servicecatalogOrganizationNode); err != nil {
			log.Errorf("invalid --organization-node: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeletePortfolioShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified product.
// You cannot delete a product if it was shared with you or is associated with a
// portfolio.
//
// A delegated admin is authorized to invoke this command.
func servicecatalog_DeleteProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeleteProductInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DeleteProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified plan.
func servicecatalog_DeleteProvisionedProductPlan(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeleteProvisionedProductPlanInput{
		// PlanId: *string, // Required
	}

	if len(_servicecatalogPlanId) > 0 {
		input.PlanId = aws.String(_servicecatalogPlanId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogIgnoreErrors) > 0 {
		if err := assignInputField(input, "IgnoreErrors", _servicecatalogIgnoreErrors); err != nil {
			log.Errorf("invalid --ignore-errors: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteProvisionedProductPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified provisioning artifact (also known as a version) for the
// specified product.
//
// You cannot delete a provisioning artifact associated with a product that was
// shared with you. You cannot delete the last provisioning artifact for a product,
// because a product must have at least one provisioning artifact.
func servicecatalog_DeleteProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeleteProvisioningArtifactInput{
		// ProductId: *string, // Required
		// ProvisioningArtifactId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DeleteProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a self-service action.
func servicecatalog_DeleteServiceAction(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeleteServiceActionInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}

	if resp, err := client.DeleteServiceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified TagOption.
// You cannot delete a TagOption if it is associated with a product or portfolio.
func servicecatalog_DeleteTagOption(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DeleteTagOptionInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}

	if resp, err := client.DeleteTagOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified constraint.
func servicecatalog_DescribeConstraint(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeConstraintInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DescribeConstraint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of the specified copy product operation.
func servicecatalog_DescribeCopyProductStatus(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeCopyProductStatusInput{
		// CopyProductToken: *string, // Required
	}

	if len(_servicecatalogCopyProductToken) > 0 {
		input.CopyProductToken = aws.String(_servicecatalogCopyProductToken)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DescribeCopyProductStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified portfolio.
// A delegated admin is authorized to invoke this command.
func servicecatalog_DescribePortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribePortfolioInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DescribePortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets the status of the specified portfolio share operation. This API can only
// be called by the management account in the organization or by a delegated admin.
func servicecatalog_DescribePortfolioShareStatus(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribePortfolioShareStatusInput{
		// PortfolioShareToken: *string, // Required
	}

	if len(_servicecatalogPortfolioShareToken) > 0 {
		input.PortfolioShareToken = aws.String(_servicecatalogPortfolioShareToken)
	}

	if resp, err := client.DescribePortfolioShareStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a summary of each of the portfolio shares that were created for the
// specified portfolio.
//
// You can use this API to determine which accounts or organizational nodes this
// portfolio have been shared, whether the recipient entity has imported the share,
// and whether TagOptions are included with the share.
//
// The PortfolioId and Type parameters are both required.
func servicecatalog_DescribePortfolioShares(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribePortfolioSharesInput{
		// PortfolioId: *string, // Required
		// Type: types.DescribePortfolioShareType, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogType) > 0 {
		if err := assignInputField(input, "Type", _servicecatalogType); err != nil {
			log.Errorf("invalid --type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribePortfolioShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.DescribePortfolioSharesOutput
	p := servicecatalog.NewDescribePortfolioSharesPaginator(client, input)
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

// Gets information about the specified product.
// Running this operation with administrator access results in a failure. DescribeProductAsAdmin should
// be used instead.
func servicecatalog_DescribeProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProductInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}

	if resp, err := client.DescribeProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified product. This operation is run with
// administrator access.
func servicecatalog_DescribeProductAsAdmin(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProductAsAdminInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}
	if len(_servicecatalogSourcePortfolioId) > 0 {
		input.SourcePortfolioId = aws.String(_servicecatalogSourcePortfolioId)
	}

	if resp, err := client.DescribeProductAsAdmin(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified product.
func servicecatalog_DescribeProductView(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProductViewInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DescribeProductView(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified provisioned product.
func servicecatalog_DescribeProvisionedProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProvisionedProductInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}

	if resp, err := client.DescribeProvisionedProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the resource changes for the specified plan.
func servicecatalog_DescribeProvisionedProductPlan(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProvisionedProductPlanInput{
		// PlanId: *string, // Required
	}

	if len(_servicecatalogPlanId) > 0 {
		input.PlanId = aws.String(_servicecatalogPlanId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if resp, err := client.DescribeProvisionedProductPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified provisioning artifact (also known as a
// version) for the specified product.
func servicecatalog_DescribeProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProvisioningArtifactInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogIncludeProvisioningArtifactParameters) > 0 {
		if err := assignInputField(input, "IncludeProvisioningArtifactParameters", _servicecatalogIncludeProvisioningArtifactParameters); err != nil {
			log.Errorf("invalid --include-provisioning-artifact-parameters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProductName) > 0 {
		input.ProductName = aws.String(_servicecatalogProductName)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogProvisioningArtifactName) > 0 {
		input.ProvisioningArtifactName = aws.String(_servicecatalogProvisioningArtifactName)
	}
	if len(_servicecatalogVerbose) > 0 {
		if err := assignInputField(input, "Verbose", _servicecatalogVerbose); err != nil {
			log.Errorf("invalid --verbose: %s", err.Error())
			return
		}
	}

	if resp, err := client.DescribeProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the configuration required to provision the specified
// product using the specified provisioning artifact.
//
// If the output contains a TagOption key with an empty list of values, there is a
// TagOption conflict for that key. The end user cannot take action to fix the
// conflict, and launch is not blocked. In subsequent calls to ProvisionProduct, do not include
// conflicted TagOption keys as tags, or this causes the error "Parameter
// validation failed: Missing required parameter in Tags[N]:Value". Tag the
// provisioned product with the value sc-tagoption-conflict-portfolioId-productId .
func servicecatalog_DescribeProvisioningParameters(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeProvisioningParametersInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPathId) > 0 {
		input.PathId = aws.String(_servicecatalogPathId)
	}
	if len(_servicecatalogPathName) > 0 {
		input.PathName = aws.String(_servicecatalogPathName)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProductName) > 0 {
		input.ProductName = aws.String(_servicecatalogProductName)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogProvisioningArtifactName) > 0 {
		input.ProvisioningArtifactName = aws.String(_servicecatalogProvisioningArtifactName)
	}

	if resp, err := client.DescribeProvisioningParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified request operation.
// Use this operation after calling a request operation (for example, ProvisionProduct, TerminateProvisionedProduct, or UpdateProvisionedProduct).
//
// If a provisioned product was transferred to a new owner using UpdateProvisionedProductProperties, the new owner
// will be able to describe all past records for that product. The previous owner
// will no longer be able to describe the records, but will be able to use ListRecordHistoryto see
// the product's history from when he was the owner.
func servicecatalog_DescribeRecord(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeRecordInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if resp, err := client.DescribeRecord(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Describes a self-service action.
func servicecatalog_DescribeServiceAction(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeServiceActionInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DescribeServiceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Finds the default parameters for a specific self-service action on a specific
// provisioned product and returns a map of the results to the user.
func servicecatalog_DescribeServiceActionExecutionParameters(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeServiceActionExecutionParametersInput{
		// ProvisionedProductId: *string, // Required
		// ServiceActionId: *string, // Required
	}

	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogServiceActionId) > 0 {
		input.ServiceActionId = aws.String(_servicecatalogServiceActionId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DescribeServiceActionExecutionParameters(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the specified TagOption.
func servicecatalog_DescribeTagOption(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DescribeTagOptionInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}

	if resp, err := client.DescribeTagOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disable portfolio sharing through the Organizations service. This command will
// not delete your current shares, but prevents you from creating new shares
// throughout your organization. Current shares are not kept in sync with your
// organization structure if the structure changes after calling this API. Only the
// management account in the organization can call this API.
//
// You cannot call this API if there are active delegated administrators in the
// organization.
//
// Note that a delegated administrator is not authorized to invoke
// DisableAWSOrganizationsAccess .
//
// If you share an Service Catalog portfolio in an organization within
// Organizations, and then disable Organizations access for Service Catalog, the
// portfolio access permissions will not sync with the latest changes to the
// organization structure. Specifically, accounts that you removed from the
// organization after disabling Service Catalog access will retain access to the
// previously shared portfolio.
func servicecatalog_DisableAWSOrganizationsAccess(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DisableAWSOrganizationsAccessInput{}

	if resp, err := client.DisableAWSOrganizationsAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified budget from the specified resource.
func servicecatalog_DisassociateBudgetFromResource(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DisassociateBudgetFromResourceInput{
		// BudgetName: *string, // Required
		// ResourceId: *string, // Required
	}

	if len(_servicecatalogBudgetName) > 0 {
		input.BudgetName = aws.String(_servicecatalogBudgetName)
	}
	if len(_servicecatalogResourceId) > 0 {
		input.ResourceId = aws.String(_servicecatalogResourceId)
	}

	if resp, err := client.DisassociateBudgetFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates a previously associated principal ARN from a specified portfolio.
// The PrincipalType and PrincipalARN must match the
// AssociatePrincipalWithPortfolio call request details. For example, to
// disassociate an association created with a PrincipalARN of PrincipalType IAM
// you must use the PrincipalType IAM when calling
// DisassociatePrincipalFromPortfolio .
//
// For portfolios that have been shared with principal name sharing enabled: after
// disassociating a principal, share recipient accounts will no longer be able to
// provision products in this portfolio using a role matching the name of the
// associated principal.
//
// For more information, review [associate-principal-with-portfolio] in the Amazon Web Services CLI Command Reference.
//
// If you disassociate a principal from a portfolio, with PrincipalType as IAM ,
// the same principal will still have access to the portfolio if it matches one of
// the associated principals of type IAM_PATTERN . To fully remove access for a
// principal, verify all the associated Principals of type IAM_PATTERN , and then
// ensure you disassociate any IAM_PATTERN principals that match the principal
// whose access you are removing.
//
// [associate-principal-with-portfolio]: https://docs.aws.amazon.com/cli/latest/reference/servicecatalog/associate-principal-with-portfolio.html#options
func servicecatalog_DisassociatePrincipalFromPortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DisassociatePrincipalFromPortfolioInput{
		// PortfolioId: *string, // Required
		// PrincipalARN: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogPrincipalARN) > 0 {
		input.PrincipalARN = aws.String(_servicecatalogPrincipalARN)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPrincipalType) > 0 {
		if err := assignInputField(input, "PrincipalType", _servicecatalogPrincipalType); err != nil {
			log.Errorf("invalid --principal-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.DisassociatePrincipalFromPortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified product from the specified portfolio.
// A delegated admin is authorized to invoke this command.
func servicecatalog_DisassociateProductFromPortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DisassociateProductFromPortfolioInput{
		// PortfolioId: *string, // Required
		// ProductId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.DisassociateProductFromPortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified self-service action association from the specified
// provisioning artifact.
func servicecatalog_DisassociateServiceActionFromProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DisassociateServiceActionFromProvisioningArtifactInput{
		// ProductId: *string, // Required
		// ProvisioningArtifactId: *string, // Required
		// ServiceActionId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogServiceActionId) > 0 {
		input.ServiceActionId = aws.String(_servicecatalogServiceActionId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}

	if resp, err := client.DisassociateServiceActionFromProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disassociates the specified TagOption from the specified resource.
func servicecatalog_DisassociateTagOptionFromResource(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.DisassociateTagOptionFromResourceInput{
		// ResourceId: *string, // Required
		// TagOptionId: *string, // Required
	}

	if len(_servicecatalogResourceId) > 0 {
		input.ResourceId = aws.String(_servicecatalogResourceId)
	}
	if len(_servicecatalogTagOptionId) > 0 {
		input.TagOptionId = aws.String(_servicecatalogTagOptionId)
	}

	if resp, err := client.DisassociateTagOptionFromResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enable portfolio sharing feature through Organizations. This API will allow
// Service Catalog to receive updates on your organization in order to sync your
// shares with the current structure. This API can only be called by the management
// account in the organization.
//
// When you call this API, Service Catalog calls
// organizations:EnableAWSServiceAccess on your behalf so that your shares stay in
// sync with any changes in your Organizations structure.
//
// Note that a delegated administrator is not authorized to invoke
// EnableAWSOrganizationsAccess .
//
// If you have previously disabled Organizations access for Service Catalog, and
// then enable access again, the portfolio access permissions might not sync with
// the latest changes to the organization structure. Specifically, accounts that
// you removed from the organization after disabling Service Catalog access, and
// before you enabled access again, can retain access to the previously shared
// portfolio. As a result, an account that has been removed from the organization
// might still be able to create or manage Amazon Web Services resources when it is
// no longer authorized to do so. Amazon Web Services is working to resolve this
// issue.
func servicecatalog_EnableAWSOrganizationsAccess(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.EnableAWSOrganizationsAccessInput{}

	if resp, err := client.EnableAWSOrganizationsAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions or modifies a product based on the resource changes for the
// specified plan.
func servicecatalog_ExecuteProvisionedProductPlan(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ExecuteProvisionedProductPlanInput{
		// IdempotencyToken: *string, // Required
		// PlanId: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogPlanId) > 0 {
		input.PlanId = aws.String(_servicecatalogPlanId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.ExecuteProvisionedProductPlan(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Executes a self-service action against a provisioned product.
func servicecatalog_ExecuteProvisionedProductServiceAction(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ExecuteProvisionedProductServiceActionInput{
		// ExecuteToken: *string, // Required
		// ProvisionedProductId: *string, // Required
		// ServiceActionId: *string, // Required
	}

	if len(_servicecatalogExecuteToken) > 0 {
		input.ExecuteToken = aws.String(_servicecatalogExecuteToken)
	}
	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogServiceActionId) > 0 {
		input.ServiceActionId = aws.String(_servicecatalogServiceActionId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogParameters) > 0 {
		if err := assignInputField(input, "Parameters", _servicecatalogParameters); err != nil {
			log.Errorf("invalid --parameters: %s", err.Error())
			return
		}
	}

	if resp, err := client.ExecuteProvisionedProductServiceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the Access Status for Organizations portfolio share feature. This API can
// only be called by the management account in the organization or by a delegated
// admin.
func servicecatalog_GetAWSOrganizationsAccessStatus(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.GetAWSOrganizationsAccessStatusInput{}

	if resp, err := client.GetAWSOrganizationsAccessStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This API takes either a ProvisonedProductId or a ProvisionedProductName , along
// with a list of one or more output keys, and responds with the key/value pairs of
// those outputs.
func servicecatalog_GetProvisionedProductOutputs(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.GetProvisionedProductOutputsInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogOutputKeys) > 0 {
		input.OutputKeys = append([]string(nil), _servicecatalogOutputKeys...)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogProvisionedProductName) > 0 {
		input.ProvisionedProductName = aws.String(_servicecatalogProvisionedProductName)
	}

	if disablePaginator() {
		if resp, err := client.GetProvisionedProductOutputs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.GetProvisionedProductOutputsOutput
	p := servicecatalog.NewGetProvisionedProductOutputsPaginator(client, input)
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

// Requests the import of a resource as an Service Catalog provisioned product
// that is associated to an Service Catalog product and provisioning artifact. Once
// imported, all supported governance actions are supported on the provisioned
// product.
//
// Resource import only supports CloudFormation stack ARNs. CloudFormation
// StackSets, and non-root nested stacks, are not supported.
//
// The CloudFormation stack must have one of the following statuses to be
// imported: CREATE_COMPLETE , UPDATE_COMPLETE , UPDATE_ROLLBACK_COMPLETE ,
// IMPORT_COMPLETE , and IMPORT_ROLLBACK_COMPLETE .
//
// Import of the resource requires that the CloudFormation stack template matches
// the associated Service Catalog product provisioning artifact.
//
// When you import an existing CloudFormation stack into a portfolio, Service
// Catalog does not apply the product's associated constraints during the import
// process. Service Catalog applies the constraints after you call
// UpdateProvisionedProduct for the provisioned product.
//
// The user or role that performs this operation must have the
// cloudformation:GetTemplate and cloudformation:DescribeStacks IAM policy
// permissions.
//
// You can only import one provisioned product at a time. The product's
// CloudFormation stack must have the IMPORT_COMPLETE status before you import
// another.
func servicecatalog_ImportAsProvisionedProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ImportAsProvisionedProductInput{
		// IdempotencyToken: *string, // Required
		// PhysicalId: *string, // Required
		// ProductId: *string, // Required
		// ProvisionedProductName: *string, // Required
		// ProvisioningArtifactId: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogPhysicalId) > 0 {
		input.PhysicalId = aws.String(_servicecatalogPhysicalId)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisionedProductName) > 0 {
		input.ProvisionedProductName = aws.String(_servicecatalogProvisionedProductName)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.ImportAsProvisionedProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all imported portfolios for which account-to-account shares were accepted
// by this account. By specifying the PortfolioShareType , you can list portfolios
// for which organizational shares were accepted by this account.
func servicecatalog_ListAcceptedPortfolioShares(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListAcceptedPortfolioSharesInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogPortfolioShareType) > 0 {
		if err := assignInputField(input, "PortfolioShareType", _servicecatalogPortfolioShareType); err != nil {
			log.Errorf("invalid --portfolio-share-type: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAcceptedPortfolioShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListAcceptedPortfolioSharesOutput
	p := servicecatalog.NewListAcceptedPortfolioSharesPaginator(client, input)
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

// Lists all the budgets associated to the specified resource.
func servicecatalog_ListBudgetsForResource(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListBudgetsForResourceInput{
		// ResourceId: *string, // Required
	}

	if len(_servicecatalogResourceId) > 0 {
		input.ResourceId = aws.String(_servicecatalogResourceId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListBudgetsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListBudgetsForResourceOutput
	p := servicecatalog.NewListBudgetsForResourcePaginator(client, input)
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

// Lists the constraints for the specified portfolio and product.
func servicecatalog_ListConstraintsForPortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListConstraintsForPortfolioInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}

	if disablePaginator() {
		if resp, err := client.ListConstraintsForPortfolio(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListConstraintsForPortfolioOutput
	p := servicecatalog.NewListConstraintsForPortfolioPaginator(client, input)
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

// Lists the paths to the specified product. A path describes how the user gets
// access to a specified product and is necessary when provisioning a product. A
// path also determines the constraints that are put on a product. A path is
// dependent on a specific product, porfolio, and principal.
//
// When provisioning a product that's been added to a portfolio, you must grant
// your user, group, or role access to the portfolio. For more information, see [Granting users access]in
// the Service Catalog User Guide.
//
// [Granting users access]: https://docs.aws.amazon.com/servicecatalog/latest/adminguide/catalogs_portfolios_users.html
func servicecatalog_ListLaunchPaths(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListLaunchPathsInput{
		// ProductId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListLaunchPaths(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListLaunchPathsOutput
	p := servicecatalog.NewListLaunchPathsPaginator(client, input)
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

// Lists the organization nodes that have access to the specified portfolio. This
// API can only be called by the management account in the organization or by a
// delegated admin.
//
// If a delegated admin is de-registered, they can no longer perform this
// operation.
func servicecatalog_ListOrganizationPortfolioAccess(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListOrganizationPortfolioAccessInput{
		// OrganizationNodeType: types.OrganizationNodeType, // Required
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogOrganizationNodeType) > 0 {
		if err := assignInputField(input, "OrganizationNodeType", _servicecatalogOrganizationNodeType); err != nil {
			log.Errorf("invalid --organization-node-type: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationPortfolioAccess(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListOrganizationPortfolioAccessOutput
	p := servicecatalog.NewListOrganizationPortfolioAccessPaginator(client, input)
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

// Lists the account IDs that have access to the specified portfolio.
// A delegated admin can list the accounts that have access to the shared
// portfolio. Note that if a delegated admin is de-registered, they can no longer
// perform this operation.
func servicecatalog_ListPortfolioAccess(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListPortfolioAccessInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogOrganizationParentId) > 0 {
		input.OrganizationParentId = aws.String(_servicecatalogOrganizationParentId)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPortfolioAccess(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListPortfolioAccessOutput
	p := servicecatalog.NewListPortfolioAccessPaginator(client, input)
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

// Lists all portfolios in the catalog.
func servicecatalog_ListPortfolios(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListPortfoliosInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPortfolios(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListPortfoliosOutput
	p := servicecatalog.NewListPortfoliosPaginator(client, input)
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

// Lists all portfolios that the specified product is associated with.
func servicecatalog_ListPortfoliosForProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListPortfoliosForProductInput{
		// ProductId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPortfoliosForProduct(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListPortfoliosForProductOutput
	p := servicecatalog.NewListPortfoliosForProductPaginator(client, input)
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

// Lists all PrincipalARN s and corresponding PrincipalType s associated with the
// specified portfolio.
func servicecatalog_ListPrincipalsForPortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListPrincipalsForPortfolioInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPrincipalsForPortfolio(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListPrincipalsForPortfolioOutput
	p := servicecatalog.NewListPrincipalsForPortfolioPaginator(client, input)
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

// Lists the plans for the specified provisioned product or all plans to which the
// user has access.
func servicecatalog_ListProvisionedProductPlans(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListProvisionedProductPlansInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccessLevelFilter) > 0 {
		if err := assignInputField(input, "AccessLevelFilter", _servicecatalogAccessLevelFilter); err != nil {
			log.Errorf("invalid --access-level-filter: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogProvisionProductId) > 0 {
		input.ProvisionProductId = aws.String(_servicecatalogProvisionProductId)
	}

	if resp, err := client.ListProvisionedProductPlans(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all provisioning artifacts (also known as versions) for the specified
// product.
func servicecatalog_ListProvisioningArtifacts(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListProvisioningArtifactsInput{
		// ProductId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.ListProvisioningArtifacts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists all provisioning artifacts (also known as versions) for the specified
// self-service action.
func servicecatalog_ListProvisioningArtifactsForServiceAction(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListProvisioningArtifactsForServiceActionInput{
		// ServiceActionId: *string, // Required
	}

	if len(_servicecatalogServiceActionId) > 0 {
		input.ServiceActionId = aws.String(_servicecatalogServiceActionId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProvisioningArtifactsForServiceAction(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListProvisioningArtifactsForServiceActionOutput
	p := servicecatalog.NewListProvisioningArtifactsForServiceActionPaginator(client, input)
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

// Lists the specified requests or all performed requests.
func servicecatalog_ListRecordHistory(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListRecordHistoryInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccessLevelFilter) > 0 {
		if err := assignInputField(input, "AccessLevelFilter", _servicecatalogAccessLevelFilter); err != nil {
			log.Errorf("invalid --access-level-filter: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogSearchFilter) > 0 {
		if err := assignInputField(input, "SearchFilter", _servicecatalogSearchFilter); err != nil {
			log.Errorf("invalid --search-filter: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListRecordHistory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the resources associated with the specified TagOption.
func servicecatalog_ListResourcesForTagOption(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListResourcesForTagOptionInput{
		// TagOptionId: *string, // Required
	}

	if len(_servicecatalogTagOptionId) > 0 {
		input.TagOptionId = aws.String(_servicecatalogTagOptionId)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogResourceType) > 0 {
		input.ResourceType = aws.String(_servicecatalogResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListResourcesForTagOption(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListResourcesForTagOptionOutput
	p := servicecatalog.NewListResourcesForTagOptionPaginator(client, input)
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

// Lists all self-service actions.
func servicecatalog_ListServiceActions(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListServiceActionsInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceActions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListServiceActionsOutput
	p := servicecatalog.NewListServiceActionsPaginator(client, input)
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

// Returns a paginated list of self-service actions associated with the specified
// Product ID and Provisioning Artifact ID.
func servicecatalog_ListServiceActionsForProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListServiceActionsForProvisioningArtifactInput{
		// ProductId: *string, // Required
		// ProvisioningArtifactId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListServiceActionsForProvisioningArtifact(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListServiceActionsForProvisioningArtifactOutput
	p := servicecatalog.NewListServiceActionsForProvisioningArtifactPaginator(client, input)
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

// Returns summary information about stack instances that are associated with the
// specified CFN_STACKSET type provisioned product. You can filter for stack
// instances that are associated with a specific Amazon Web Services account name
// or Region.
func servicecatalog_ListStackInstancesForProvisionedProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListStackInstancesForProvisionedProductInput{
		// ProvisionedProductId: *string, // Required
	}

	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if resp, err := client.ListStackInstancesForProvisionedProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the specified TagOptions or all TagOptions.
func servicecatalog_ListTagOptions(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ListTagOptionsInput{}

	if len(_servicecatalogFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicecatalogFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagOptions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.ListTagOptionsOutput
	p := servicecatalog.NewListTagOptionsPaginator(client, input)
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

// Notifies the result of the provisioning engine execution.
func servicecatalog_NotifyProvisionProductEngineWorkflowResult(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.NotifyProvisionProductEngineWorkflowResultInput{
		// IdempotencyToken: *string, // Required
		// RecordId: *string, // Required
		// Status: types.EngineWorkflowStatus, // Required
		// WorkflowToken: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogRecordId) > 0 {
		input.RecordId = aws.String(_servicecatalogRecordId)
	}
	if len(_servicecatalogStatus) > 0 {
		if err := assignInputField(input, "Status", _servicecatalogStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogWorkflowToken) > 0 {
		input.WorkflowToken = aws.String(_servicecatalogWorkflowToken)
	}
	if len(_servicecatalogFailureReason) > 0 {
		input.FailureReason = aws.String(_servicecatalogFailureReason)
	}
	if len(_servicecatalogOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _servicecatalogOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogResourceIdentifier) > 0 {
		if err := assignInputField(input, "ResourceIdentifier", _servicecatalogResourceIdentifier); err != nil {
			log.Errorf("invalid --resource-identifier: %s", err.Error())
			return
		}
	}

	if resp, err := client.NotifyProvisionProductEngineWorkflowResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies the result of the terminate engine execution.
func servicecatalog_NotifyTerminateProvisionedProductEngineWorkflowResult(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.NotifyTerminateProvisionedProductEngineWorkflowResultInput{
		// IdempotencyToken: *string, // Required
		// RecordId: *string, // Required
		// Status: types.EngineWorkflowStatus, // Required
		// WorkflowToken: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogRecordId) > 0 {
		input.RecordId = aws.String(_servicecatalogRecordId)
	}
	if len(_servicecatalogStatus) > 0 {
		if err := assignInputField(input, "Status", _servicecatalogStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogWorkflowToken) > 0 {
		input.WorkflowToken = aws.String(_servicecatalogWorkflowToken)
	}
	if len(_servicecatalogFailureReason) > 0 {
		input.FailureReason = aws.String(_servicecatalogFailureReason)
	}

	if resp, err := client.NotifyTerminateProvisionedProductEngineWorkflowResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Notifies the result of the update engine execution.
func servicecatalog_NotifyUpdateProvisionedProductEngineWorkflowResult(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.NotifyUpdateProvisionedProductEngineWorkflowResultInput{
		// IdempotencyToken: *string, // Required
		// RecordId: *string, // Required
		// Status: types.EngineWorkflowStatus, // Required
		// WorkflowToken: *string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogRecordId) > 0 {
		input.RecordId = aws.String(_servicecatalogRecordId)
	}
	if len(_servicecatalogStatus) > 0 {
		if err := assignInputField(input, "Status", _servicecatalogStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogWorkflowToken) > 0 {
		input.WorkflowToken = aws.String(_servicecatalogWorkflowToken)
	}
	if len(_servicecatalogFailureReason) > 0 {
		input.FailureReason = aws.String(_servicecatalogFailureReason)
	}
	if len(_servicecatalogOutputs) > 0 {
		if err := assignInputField(input, "Outputs", _servicecatalogOutputs); err != nil {
			log.Errorf("invalid --outputs: %s", err.Error())
			return
		}
	}

	if resp, err := client.NotifyUpdateProvisionedProductEngineWorkflowResult(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Provisions the specified product.
// A provisioned product is a resourced instance of a product. For example,
// provisioning a product that's based on an CloudFormation template launches an
// CloudFormation stack and its underlying resources. You can check the status of
// this request using DescribeRecord.
//
// If the request contains a tag key with an empty list of values, there's a tag
// conflict for that key. Don't include conflicted keys as tags, or this will cause
// the error "Parameter validation failed: Missing required parameter in
// Tags[N]:Value".
//
// When provisioning a product that's been added to a portfolio, you must grant
// your user, group, or role access to the portfolio. For more information, see [Granting users access]in
// the Service Catalog User Guide.
//
// [Granting users access]: https://docs.aws.amazon.com/servicecatalog/latest/adminguide/catalogs_portfolios_users.html
func servicecatalog_ProvisionProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ProvisionProductInput{
		// ProvisionToken: *string, // Required
		// ProvisionedProductName: *string, // Required
	}

	if len(_servicecatalogProvisionToken) > 0 {
		input.ProvisionToken = aws.String(_servicecatalogProvisionToken)
	}
	if len(_servicecatalogProvisionedProductName) > 0 {
		input.ProvisionedProductName = aws.String(_servicecatalogProvisionedProductName)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogNotificationArns) > 0 {
		input.NotificationArns = append([]string(nil), _servicecatalogNotificationArns...)
	}
	if len(_servicecatalogPathId) > 0 {
		input.PathId = aws.String(_servicecatalogPathId)
	}
	if len(_servicecatalogPathName) > 0 {
		input.PathName = aws.String(_servicecatalogPathName)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProductName) > 0 {
		input.ProductName = aws.String(_servicecatalogProductName)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogProvisioningArtifactName) > 0 {
		input.ProvisioningArtifactName = aws.String(_servicecatalogProvisioningArtifactName)
	}
	if len(_servicecatalogProvisioningParameters) > 0 {
		if err := assignInputField(input, "ProvisioningParameters", _servicecatalogProvisioningParameters); err != nil {
			log.Errorf("invalid --provisioning-parameters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogProvisioningPreferences) > 0 {
		if err := assignInputField(input, "ProvisioningPreferences", _servicecatalogProvisioningPreferences); err != nil {
			log.Errorf("invalid --provisioning-preferences: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.ProvisionProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects an offer to share the specified portfolio.
func servicecatalog_RejectPortfolioShare(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.RejectPortfolioShareInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPortfolioShareType) > 0 {
		if err := assignInputField(input, "PortfolioShareType", _servicecatalogPortfolioShareType); err != nil {
			log.Errorf("invalid --portfolio-share-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.RejectPortfolioShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the provisioned products that are available (not terminated).
// To use additional filtering, see SearchProvisionedProducts.
func servicecatalog_ScanProvisionedProducts(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.ScanProvisionedProductsInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccessLevelFilter) > 0 {
		if err := assignInputField(input, "AccessLevelFilter", _servicecatalogAccessLevelFilter); err != nil {
			log.Errorf("invalid --access-level-filter: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}

	if resp, err := client.ScanProvisionedProducts(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information about the products to which the caller has access.
func servicecatalog_SearchProducts(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.SearchProductsInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicecatalogFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _servicecatalogSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _servicecatalogSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchProducts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.SearchProductsOutput
	p := servicecatalog.NewSearchProductsPaginator(client, input)
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

// Gets information about the products for the specified portfolio or all products.
func servicecatalog_SearchProductsAsAdmin(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.SearchProductsAsAdminInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicecatalogFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogProductSource) > 0 {
		if err := assignInputField(input, "ProductSource", _servicecatalogProductSource); err != nil {
			log.Errorf("invalid --product-source: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSortBy) > 0 {
		if err := assignInputField(input, "SortBy", _servicecatalogSortBy); err != nil {
			log.Errorf("invalid --sort-by: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _servicecatalogSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchProductsAsAdmin(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.SearchProductsAsAdminOutput
	p := servicecatalog.NewSearchProductsAsAdminPaginator(client, input)
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

// Gets information about the provisioned products that meet the specified
// criteria.
func servicecatalog_SearchProvisionedProducts(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.SearchProvisionedProductsInput{}

	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccessLevelFilter) > 0 {
		if err := assignInputField(input, "AccessLevelFilter", _servicecatalogAccessLevelFilter); err != nil {
			log.Errorf("invalid --access-level-filter: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogFilters) > 0 {
		if err := assignInputField(input, "Filters", _servicecatalogFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageSize) > 0 {
		if err := assignInputField(input, "PageSize", _servicecatalogPageSize); err != nil {
			log.Errorf("invalid --page-size: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogPageToken) > 0 {
		input.PageToken = aws.String(_servicecatalogPageToken)
	}
	if len(_servicecatalogSortBy) > 0 {
		input.SortBy = aws.String(_servicecatalogSortBy)
	}
	if len(_servicecatalogSortOrder) > 0 {
		if err := assignInputField(input, "SortOrder", _servicecatalogSortOrder); err != nil {
			log.Errorf("invalid --sort-order: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.SearchProvisionedProducts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*servicecatalog.SearchProvisionedProductsOutput
	p := servicecatalog.NewSearchProvisionedProductsPaginator(client, input)
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

// Terminates the specified provisioned product.
// This operation does not delete any records associated with the provisioned
// product.
//
// You can check the status of this request using DescribeRecord.
func servicecatalog_TerminateProvisionedProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.TerminateProvisionedProductInput{
		// TerminateToken: *string, // Required
	}

	if len(_servicecatalogTerminateToken) > 0 {
		input.TerminateToken = aws.String(_servicecatalogTerminateToken)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogIgnoreErrors) > 0 {
		if err := assignInputField(input, "IgnoreErrors", _servicecatalogIgnoreErrors); err != nil {
			log.Errorf("invalid --ignore-errors: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogProvisionedProductName) > 0 {
		input.ProvisionedProductName = aws.String(_servicecatalogProvisionedProductName)
	}
	if len(_servicecatalogRetainPhysicalResources) > 0 {
		if err := assignInputField(input, "RetainPhysicalResources", _servicecatalogRetainPhysicalResources); err != nil {
			log.Errorf("invalid --retain-physical-resources: %s", err.Error())
			return
		}
	}

	if resp, err := client.TerminateProvisionedProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified constraint.
func servicecatalog_UpdateConstraint(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateConstraintInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogParameters) > 0 {
		input.Parameters = aws.String(_servicecatalogParameters)
	}

	if resp, err := client.UpdateConstraint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified portfolio.
// You cannot update a product that was shared with you.
func servicecatalog_UpdatePortfolio(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdatePortfolioInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAddTags) > 0 {
		if err := assignInputField(input, "AddTags", _servicecatalogAddTags); err != nil {
			log.Errorf("invalid --add-tags: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogDisplayName) > 0 {
		input.DisplayName = aws.String(_servicecatalogDisplayName)
	}
	if len(_servicecatalogProviderName) > 0 {
		input.ProviderName = aws.String(_servicecatalogProviderName)
	}
	if len(_servicecatalogRemoveTags) > 0 {
		input.RemoveTags = append([]string(nil), _servicecatalogRemoveTags...)
	}

	if resp, err := client.UpdatePortfolio(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified portfolio share. You can use this API to enable or
// disable TagOptions sharing or Principal sharing for an existing portfolio
// share.
//
// The portfolio share cannot be updated if the CreatePortfolioShare operation is
// IN_PROGRESS , as the share is not available to recipient entities. In this case,
// you must wait for the portfolio share to be completed.
//
// You must provide the accountId or organization node in the input, but not both.
//
// If the portfolio is shared to both an external account and an organization
// node, and both shares need to be updated, you must invoke UpdatePortfolioShare
// separately for each share type.
//
// This API cannot be used for removing the portfolio share. You must use
// DeletePortfolioShare API for that action.
//
// When you associate a principal with portfolio, a potential privilege escalation
// path may occur when that portfolio is then shared with other accounts. For a
// user in a recipient account who is not an Service Catalog Admin, but still has
// the ability to create Principals (Users/Groups/Roles), that user could create a
// role that matches a principal name association for the portfolio. Although this
// user may not know which principal names are associated through Service Catalog,
// they may be able to guess the user. If this potential escalation path is a
// concern, then Service Catalog recommends using PrincipalType as IAM . With this
// configuration, the PrincipalARN must already exist in the recipient account
// before it can be associated.
func servicecatalog_UpdatePortfolioShare(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdatePortfolioShareInput{
		// PortfolioId: *string, // Required
	}

	if len(_servicecatalogPortfolioId) > 0 {
		input.PortfolioId = aws.String(_servicecatalogPortfolioId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAccountId) > 0 {
		input.AccountId = aws.String(_servicecatalogAccountId)
	}
	if len(_servicecatalogOrganizationNode) > 0 {
		if err := assignInputField(input, "OrganizationNode", _servicecatalogOrganizationNode); err != nil {
			log.Errorf("invalid --organization-node: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSharePrincipals) > 0 {
		if err := assignInputField(input, "SharePrincipals", _servicecatalogSharePrincipals); err != nil {
			log.Errorf("invalid --share-principals: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogShareTagOptions) > 0 {
		if err := assignInputField(input, "ShareTagOptions", _servicecatalogShareTagOptions); err != nil {
			log.Errorf("invalid --share-tag-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePortfolioShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified product.
func servicecatalog_UpdateProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateProductInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogAddTags) > 0 {
		if err := assignInputField(input, "AddTags", _servicecatalogAddTags); err != nil {
			log.Errorf("invalid --add-tags: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogDistributor) > 0 {
		input.Distributor = aws.String(_servicecatalogDistributor)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}
	if len(_servicecatalogOwner) > 0 {
		input.Owner = aws.String(_servicecatalogOwner)
	}
	if len(_servicecatalogRemoveTags) > 0 {
		input.RemoveTags = append([]string(nil), _servicecatalogRemoveTags...)
	}
	if len(_servicecatalogSourceConnection) > 0 {
		if err := assignInputField(input, "SourceConnection", _servicecatalogSourceConnection); err != nil {
			log.Errorf("invalid --source-connection: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogSupportDescription) > 0 {
		input.SupportDescription = aws.String(_servicecatalogSupportDescription)
	}
	if len(_servicecatalogSupportEmail) > 0 {
		input.SupportEmail = aws.String(_servicecatalogSupportEmail)
	}
	if len(_servicecatalogSupportUrl) > 0 {
		input.SupportUrl = aws.String(_servicecatalogSupportUrl)
	}

	if resp, err := client.UpdateProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests updates to the configuration of the specified provisioned product.
// If there are tags associated with the object, they cannot be updated or added.
// Depending on the specific updates requested, this operation can update with no
// interruption, with some interruption, or replace the provisioned product
// entirely.
//
// You can check the status of this request using DescribeRecord.
func servicecatalog_UpdateProvisionedProduct(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateProvisionedProductInput{
		// UpdateToken: *string, // Required
	}

	if len(_servicecatalogUpdateToken) > 0 {
		input.UpdateToken = aws.String(_servicecatalogUpdateToken)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogPathId) > 0 {
		input.PathId = aws.String(_servicecatalogPathId)
	}
	if len(_servicecatalogPathName) > 0 {
		input.PathName = aws.String(_servicecatalogPathName)
	}
	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProductName) > 0 {
		input.ProductName = aws.String(_servicecatalogProductName)
	}
	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogProvisionedProductName) > 0 {
		input.ProvisionedProductName = aws.String(_servicecatalogProvisionedProductName)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogProvisioningArtifactName) > 0 {
		input.ProvisioningArtifactName = aws.String(_servicecatalogProvisioningArtifactName)
	}
	if len(_servicecatalogProvisioningParameters) > 0 {
		if err := assignInputField(input, "ProvisioningParameters", _servicecatalogProvisioningParameters); err != nil {
			log.Errorf("invalid --provisioning-parameters: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogProvisioningPreferences) > 0 {
		if err := assignInputField(input, "ProvisioningPreferences", _servicecatalogProvisioningPreferences); err != nil {
			log.Errorf("invalid --provisioning-preferences: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogTags) > 0 {
		if err := assignInputField(input, "Tags", _servicecatalogTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateProvisionedProduct(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Requests updates to the properties of the specified provisioned product.
func servicecatalog_UpdateProvisionedProductProperties(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateProvisionedProductPropertiesInput{
		// IdempotencyToken: *string, // Required
		// ProvisionedProductId: *string, // Required
		// ProvisionedProductProperties: map[string]string, // Required
	}

	if len(_servicecatalogIdempotencyToken) > 0 {
		input.IdempotencyToken = aws.String(_servicecatalogIdempotencyToken)
	}
	if len(_servicecatalogProvisionedProductId) > 0 {
		input.ProvisionedProductId = aws.String(_servicecatalogProvisionedProductId)
	}
	if len(_servicecatalogProvisionedProductProperties) > 0 {
		if err := assignInputField(input, "ProvisionedProductProperties", _servicecatalogProvisionedProductProperties); err != nil {
			log.Errorf("invalid --provisioned-product-properties: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}

	if resp, err := client.UpdateProvisionedProductProperties(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified provisioning artifact (also known as a version) for the
// specified product.
//
// You cannot update a provisioning artifact for a product that was shared with
// you.
func servicecatalog_UpdateProvisioningArtifact(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateProvisioningArtifactInput{
		// ProductId: *string, // Required
		// ProvisioningArtifactId: *string, // Required
	}

	if len(_servicecatalogProductId) > 0 {
		input.ProductId = aws.String(_servicecatalogProductId)
	}
	if len(_servicecatalogProvisioningArtifactId) > 0 {
		input.ProvisioningArtifactId = aws.String(_servicecatalogProvisioningArtifactId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogActive) > 0 {
		if err := assignInputField(input, "Active", _servicecatalogActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogGuidance) > 0 {
		if err := assignInputField(input, "Guidance", _servicecatalogGuidance); err != nil {
			log.Errorf("invalid --guidance: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}

	if resp, err := client.UpdateProvisioningArtifact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a self-service action.
func servicecatalog_UpdateServiceAction(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateServiceActionInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogAcceptLanguage) > 0 {
		input.AcceptLanguage = aws.String(_servicecatalogAcceptLanguage)
	}
	if len(_servicecatalogDefinition) > 0 {
		if err := assignInputField(input, "Definition", _servicecatalogDefinition); err != nil {
			log.Errorf("invalid --definition: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogDescription) > 0 {
		input.Description = aws.String(_servicecatalogDescription)
	}
	if len(_servicecatalogName) > 0 {
		input.Name = aws.String(_servicecatalogName)
	}

	if resp, err := client.UpdateServiceAction(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified TagOption.
func servicecatalog_UpdateTagOption(cfg aws.Config, client *servicecatalog.Client) {
	input := &servicecatalog.UpdateTagOptionInput{
		// Id: *string, // Required
	}

	if len(_servicecatalogId) > 0 {
		input.Id = aws.String(_servicecatalogId)
	}
	if len(_servicecatalogActive) > 0 {
		if err := assignInputField(input, "Active", _servicecatalogActive); err != nil {
			log.Errorf("invalid --active: %s", err.Error())
			return
		}
	}
	if len(_servicecatalogValue) > 0 {
		input.Value = aws.String(_servicecatalogValue)
	}

	if resp, err := client.UpdateTagOption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_servicecatalogCmd)
	_servicecatalogCmd.Flags().SortFlags = false

	_servicecatalogCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_servicecatalogCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_servicecatalogCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogAcceptLanguage, "accept-language", "", "", "Accept Language")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogAccessLevelFilter, "access-level-filter", "", "", "Access Level Filter")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogAccountId, "account-id", "", "", "Account ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogActive, "active", "", "", "Active")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogAddTags, "add-tags", "", "", "Add Tags")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogBudgetName, "budget-name", "", "", "Budget Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogCopyOptions, "copy-options", "", "", "Copy Options")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogCopyProductToken, "copy-product-token", "", "", "Copy Product Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogDefinition, "definition", "", "", "Definition")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogDefinitionType, "definition-type", "", "", "Definition Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogDescription, "description", "", "", "Description")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogDisplayName, "display-name", "", "", "Display Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogDistributor, "distributor", "", "", "Distributor")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogExecuteToken, "execute-token", "", "", "Execute Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogFailureReason, "failure-reason", "", "", "Failure Reason")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogFilters, "filters", "", "", "Filters")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogGuidance, "guidance", "", "", "Guidance")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogId, "id", "", "", "ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogIdempotencyToken, "idempotency-token", "", "", "Idempotency Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogIgnoreErrors, "ignore-errors", "", "", "Ignore Errors")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogIncludeProvisioningArtifactParameters, "include-provisioning-artifact-parameters", "", "", "Include Provisioning Artifact Parameters")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogKey, "key", "", "", "Key")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogName, "name", "", "", "Name")
	_servicecatalogCmd.Flags().StringSliceVarP(&_servicecatalogNotificationArns, "notification-arns", "", nil, "Notification Arns")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogOrganizationNode, "organization-node", "", "", "Organization Node")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogOrganizationNodeType, "organization-node-type", "", "", "Organization Node Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogOrganizationParentId, "organization-parent-id", "", "", "Organization Parent ID")
	_servicecatalogCmd.Flags().StringSliceVarP(&_servicecatalogOutputKeys, "output-keys", "", nil, "Output Keys")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogOutputs, "outputs", "", "", "Outputs")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogOwner, "owner", "", "", "Owner")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPageSize, "page-size", "", "", "Page Size")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPageToken, "page-token", "", "", "Page Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogParameters, "parameters", "", "", "Parameters")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPathId, "path-id", "", "", "Path ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPathName, "path-name", "", "", "Path Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPhysicalId, "physical-id", "", "", "Physical ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPlanId, "plan-id", "", "", "Plan ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPlanName, "plan-name", "", "", "Plan Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPlanType, "plan-type", "", "", "Plan Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPortfolioId, "portfolio-id", "", "", "Portfolio ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPortfolioShareToken, "portfolio-share-token", "", "", "Portfolio Share Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPortfolioShareType, "portfolio-share-type", "", "", "Portfolio Share Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPrincipalARN, "principal-arn", "", "", "Principal ARN")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogPrincipalType, "principal-type", "", "", "Principal Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProductId, "product-id", "", "", "Product ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProductName, "product-name", "", "", "Product Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProductSource, "product-source", "", "", "Product Source")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProductType, "product-type", "", "", "Product Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProviderName, "provider-name", "", "", "Provider Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisionProductId, "provision-product-id", "", "", "Provision Product ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisionToken, "provision-token", "", "", "Provision Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisionedProductId, "provisioned-product-id", "", "", "Provisioned Product ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisionedProductName, "provisioned-product-name", "", "", "Provisioned Product Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisionedProductProperties, "provisioned-product-properties", "", "", "Provisioned Product Properties")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisioningArtifactId, "provisioning-artifact-id", "", "", "Provisioning Artifact ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisioningArtifactName, "provisioning-artifact-name", "", "", "Provisioning Artifact Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisioningArtifactParameters, "provisioning-artifact-parameters", "", "", "Provisioning Artifact Parameters")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisioningParameters, "provisioning-parameters", "", "", "Provisioning Parameters")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogProvisioningPreferences, "provisioning-preferences", "", "", "Provisioning Preferences")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogRecordId, "record-id", "", "", "Record ID")
	_servicecatalogCmd.Flags().StringSliceVarP(&_servicecatalogRemoveTags, "remove-tags", "", nil, "Remove Tags")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogResourceId, "resource-id", "", "", "Resource ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogResourceIdentifier, "resource-identifier", "", "", "Resource Identifier")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogResourceType, "resource-type", "", "", "Resource Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogRetainPhysicalResources, "retain-physical-resources", "", "", "Retain Physical Resources")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSearchFilter, "search-filter", "", "", "Search Filter")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogServiceActionAssociations, "service-action-associations", "", "", "Service Action Associations")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogServiceActionId, "service-action-id", "", "", "Service Action ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSharePrincipals, "share-principals", "", "", "Share Principals")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogShareTagOptions, "share-tag-options", "", "", "Share Tag Options")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSortBy, "sort-by", "", "", "Sort By")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSortOrder, "sort-order", "", "", "Sort Order")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSourceConnection, "source-connection", "", "", "Source Connection")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSourcePortfolioId, "source-portfolio-id", "", "", "Source Portfolio ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSourceProductArn, "source-product-arn", "", "", "Source Product ARN")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSourceProvisioningArtifactIdentifiers, "source-provisioning-artifact-identifiers", "", "", "Source Provisioning Artifact Identifiers")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogStatus, "status", "", "", "Status")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSupportDescription, "support-description", "", "", "Support Description")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSupportEmail, "support-email", "", "", "Support Email")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogSupportUrl, "support-url", "", "", "Support URL")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogTagOptionId, "tag-option-id", "", "", "Tag Option ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogTags, "tags", "", "", "Tags")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogTargetProductId, "target-product-id", "", "", "Target Product ID")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogTargetProductName, "target-product-name", "", "", "Target Product Name")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogTerminateToken, "terminate-token", "", "", "Terminate Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogType, "type", "", "", "Type")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogUpdateToken, "update-token", "", "", "Update Token")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogValue, "value", "", "", "Value")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogVerbose, "verbose", "", "", "Verbose")
	_servicecatalogCmd.Flags().StringVarP(&_servicecatalogWorkflowToken, "workflow-token", "", "", "Workflow Token")

	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogAcceptPortfolioShare, "accept-portfolio-share", "", false, "Accept Portfolio Share")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogAssociateBudgetWithResource, "associate-budget-with-resource", "", false, "Associate Budget With Resource")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogAssociatePrincipalWithPortfolio, "associate-principal-with-portfolio", "", false, "Associate Principal With Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogAssociateProductWithPortfolio, "associate-product-with-portfolio", "", false, "Associate Product With Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogAssociateServiceActionWithProvisioningArtifact, "associate-service-action-with-provisioning-artifact", "", false, "Associate Service Action With Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogAssociateTagOptionWithResource, "associate-tag-option-with-resource", "", false, "Associate Tag Option With Resource")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogBatchAssociateServiceActionWithProvisioningArtifact, "batch-associate-service-action-with-provisioning-artifact", "", false, "Batch Associate Service Action With Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogBatchDisassociateServiceActionFromProvisioningArtifact, "batch-disassociate-service-action-from-provisioning-artifact", "", false, "Batch Disassociate Service Action From Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCopyProduct, "copy-product", "", false, "Copy Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreateConstraint, "create-constraint", "", false, "Create Constraint")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreatePortfolio, "create-portfolio", "", false, "Create Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreatePortfolioShare, "create-portfolio-share", "", false, "Create Portfolio Share")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreateProduct, "create-product", "", false, "Create Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreateProvisionedProductPlan, "create-provisioned-product-plan", "", false, "Create Provisioned Product Plan")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreateProvisioningArtifact, "create-provisioning-artifact", "", false, "Create Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreateServiceAction, "create-service-action", "", false, "Create Service Action")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogCreateTagOption, "create-tag-option", "", false, "Create Tag Option")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeleteConstraint, "delete-constraint", "", false, "Delete Constraint")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeletePortfolio, "delete-portfolio", "", false, "Delete Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeletePortfolioShare, "delete-portfolio-share", "", false, "Delete Portfolio Share")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeleteProduct, "delete-product", "", false, "Delete Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeleteProvisionedProductPlan, "delete-provisioned-product-plan", "", false, "Delete Provisioned Product Plan")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeleteProvisioningArtifact, "delete-provisioning-artifact", "", false, "Delete Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeleteServiceAction, "delete-service-action", "", false, "Delete Service Action")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDeleteTagOption, "delete-tag-option", "", false, "Delete Tag Option")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeConstraint, "describe-constraint", "", false, "Describe Constraint")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeCopyProductStatus, "describe-copy-product-status", "", false, "Describe Copy Product Status")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribePortfolio, "describe-portfolio", "", false, "Describe Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribePortfolioShareStatus, "describe-portfolio-share-status", "", false, "Describe Portfolio Share Status")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribePortfolioShares, "describe-portfolio-shares", "", false, "Describe Portfolio Shares")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProduct, "describe-product", "", false, "Describe Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProductAsAdmin, "describe-product-as-admin", "", false, "Describe Product As Admin")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProductView, "describe-product-view", "", false, "Describe Product View")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProvisionedProduct, "describe-provisioned-product", "", false, "Describe Provisioned Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProvisionedProductPlan, "describe-provisioned-product-plan", "", false, "Describe Provisioned Product Plan")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProvisioningArtifact, "describe-provisioning-artifact", "", false, "Describe Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeProvisioningParameters, "describe-provisioning-parameters", "", false, "Describe Provisioning Parameters")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeRecord, "describe-record", "", false, "Describe Record")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeServiceAction, "describe-service-action", "", false, "Describe Service Action")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeServiceActionExecutionParameters, "describe-service-action-execution-parameters", "", false, "Describe Service Action Execution Parameters")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDescribeTagOption, "describe-tag-option", "", false, "Describe Tag Option")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDisableAWSOrganizationsAccess, "disable-aws-organizations-access", "", false, "Disable AWS Organizations Access")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDisassociateBudgetFromResource, "disassociate-budget-from-resource", "", false, "Disassociate Budget From Resource")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDisassociatePrincipalFromPortfolio, "disassociate-principal-from-portfolio", "", false, "Disassociate Principal From Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDisassociateProductFromPortfolio, "disassociate-product-from-portfolio", "", false, "Disassociate Product From Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDisassociateServiceActionFromProvisioningArtifact, "disassociate-service-action-from-provisioning-artifact", "", false, "Disassociate Service Action From Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogDisassociateTagOptionFromResource, "disassociate-tag-option-from-resource", "", false, "Disassociate Tag Option From Resource")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogEnableAWSOrganizationsAccess, "enable-aws-organizations-access", "", false, "Enable AWS Organizations Access")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogExecuteProvisionedProductPlan, "execute-provisioned-product-plan", "", false, "Execute Provisioned Product Plan")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogExecuteProvisionedProductServiceAction, "execute-provisioned-product-service-action", "", false, "Execute Provisioned Product Service Action")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogGetAWSOrganizationsAccessStatus, "get-aws-organizations-access-status", "", false, "Get AWS Organizations Access Status")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogGetProvisionedProductOutputs, "get-provisioned-product-outputs", "", false, "Get Provisioned Product Outputs")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogImportAsProvisionedProduct, "import-as-provisioned-product", "", false, "Import As Provisioned Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListAcceptedPortfolioShares, "list-accepted-portfolio-shares", "", false, "List Accepted Portfolio Shares")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListBudgetsForResource, "list-budgets-for-resource", "", false, "List Budgets For Resource")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListConstraintsForPortfolio, "list-constraints-for-portfolio", "", false, "List Constraints For Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListLaunchPaths, "list-launch-paths", "", false, "List Launch Paths")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListOrganizationPortfolioAccess, "list-organization-portfolio-access", "", false, "List Organization Portfolio Access")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListPortfolioAccess, "list-portfolio-access", "", false, "List Portfolio Access")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListPortfolios, "list-portfolios", "", false, "List Portfolios")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListPortfoliosForProduct, "list-portfolios-for-product", "", false, "List Portfolios For Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListPrincipalsForPortfolio, "list-principals-for-portfolio", "", false, "List Principals For Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListProvisionedProductPlans, "list-provisioned-product-plans", "", false, "List Provisioned Product Plans")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListProvisioningArtifacts, "list-provisioning-artifacts", "", false, "List Provisioning Artifacts")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListProvisioningArtifactsForServiceAction, "list-provisioning-artifacts-for-service-action", "", false, "List Provisioning Artifacts For Service Action")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListRecordHistory, "list-record-history", "", false, "List Record History")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListResourcesForTagOption, "list-resources-for-tag-option", "", false, "List Resources For Tag Option")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListServiceActions, "list-service-actions", "", false, "List Service Actions")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListServiceActionsForProvisioningArtifact, "list-service-actions-for-provisioning-artifact", "", false, "List Service Actions For Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListStackInstancesForProvisionedProduct, "list-stack-instances-for-provisioned-product", "", false, "List Stack Instances For Provisioned Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogListTagOptions, "list-tag-options", "", false, "List Tag Options")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogNotifyProvisionProductEngineWorkflowResult, "notify-provision-product-engine-workflow-result", "", false, "Notify Provision Product Engine Workflow Result")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogNotifyTerminateProvisionedProductEngineWorkflowResult, "notify-terminate-provisioned-product-engine-workflow-result", "", false, "Notify Terminate Provisioned Product Engine Workflow Result")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogNotifyUpdateProvisionedProductEngineWorkflowResult, "notify-update-provisioned-product-engine-workflow-result", "", false, "Notify Update Provisioned Product Engine Workflow Result")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogProvisionProduct, "provision-product", "", false, "Provision Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogRejectPortfolioShare, "reject-portfolio-share", "", false, "Reject Portfolio Share")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogScanProvisionedProducts, "scan-provisioned-products", "", false, "Scan Provisioned Products")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogSearchProducts, "search-products", "", false, "Search Products")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogSearchProductsAsAdmin, "search-products-as-admin", "", false, "Search Products As Admin")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogSearchProvisionedProducts, "search-provisioned-products", "", false, "Search Provisioned Products")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogTerminateProvisionedProduct, "terminate-provisioned-product", "", false, "Terminate Provisioned Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateConstraint, "update-constraint", "", false, "Update Constraint")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdatePortfolio, "update-portfolio", "", false, "Update Portfolio")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdatePortfolioShare, "update-portfolio-share", "", false, "Update Portfolio Share")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateProduct, "update-product", "", false, "Update Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateProvisionedProduct, "update-provisioned-product", "", false, "Update Provisioned Product")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateProvisionedProductProperties, "update-provisioned-product-properties", "", false, "Update Provisioned Product Properties")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateProvisioningArtifact, "update-provisioning-artifact", "", false, "Update Provisioning Artifact")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateServiceAction, "update-service-action", "", false, "Update Service Action")
	_servicecatalogCmd.Flags().BoolVarP(&_servicecatalogUpdateTagOption, "update-tag-option", "", false, "Update Tag Option")

}
