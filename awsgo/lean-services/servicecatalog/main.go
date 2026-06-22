package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/servicecatalog"
)

var fields_accept_portfolio_share = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "PortfolioShareType", Flag: "portfolio-share-type", Type: "types.PortfolioShareType", Required: false},
}

var fields_associate_budget_with_resource = []leanruntime.Field{
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_associate_principal_with_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "PrincipalARN", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: true},
}

var fields_associate_product_with_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "SourcePortfolioId", Flag: "source-portfolio-id", Type: "*string", Required: false},
}

var fields_associate_service_action_with_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
	{Name: "ServiceActionId", Flag: "service-action-id", Type: "*string", Required: true},
}

var fields_associate_tag_option_with_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagOptionId", Flag: "tag-option-id", Type: "*string", Required: true},
}

var fields_batch_associate_service_action_with_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "ServiceActionAssociations", Flag: "service-action-associations", Type: "[]types.ServiceActionAssociation", Required: true},
}

var fields_batch_disassociate_service_action_from_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "ServiceActionAssociations", Flag: "service-action-associations", Type: "[]types.ServiceActionAssociation", Required: true},
}

var fields_copy_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "CopyOptions", Flag: "copy-options", Type: "[]types.CopyOption", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "SourceProductArn", Flag: "source-product-arn", Type: "*string", Required: true},
	{Name: "SourceProvisioningArtifactIdentifiers", Flag: "source-provisioning-artifact-identifiers", Type: "[]map[string]string", Required: false},
	{Name: "TargetProductId", Flag: "target-product-id", Type: "*string", Required: false},
	{Name: "TargetProductName", Flag: "target-product-name", Type: "*string", Required: false},
}

var fields_create_constraint = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*string", Required: true},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
}

var fields_create_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_portfolio_share = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "OrganizationNode", Flag: "organization-node", Type: "*types.OrganizationNode", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "SharePrincipals", Flag: "share-principals", Type: "bool", Required: false},
	{Name: "ShareTagOptions", Flag: "share-tag-options", Type: "bool", Required: false},
}

var fields_create_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Distributor", Flag: "distributor", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: true},
	{Name: "ProductType", Flag: "product-type", Type: "types.ProductType", Required: true},
	{Name: "ProvisioningArtifactParameters", Flag: "provisioning-artifact-parameters", Type: "*types.ProvisioningArtifactProperties", Required: false},
	{Name: "SourceConnection", Flag: "source-connection", Type: "*types.SourceConnection", Required: false},
	{Name: "SupportDescription", Flag: "support-description", Type: "*string", Required: false},
	{Name: "SupportEmail", Flag: "support-email", Type: "*string", Required: false},
	{Name: "SupportUrl", Flag: "support-url", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_provisioned_product_plan = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "NotificationArns", Flag: "notification-arns", Type: "[]string", Required: false},
	{Name: "PathId", Flag: "path-id", Type: "*string", Required: false},
	{Name: "PlanName", Flag: "plan-name", Type: "*string", Required: true},
	{Name: "PlanType", Flag: "plan-type", Type: "types.ProvisionedProductPlanType", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisionedProductName", Flag: "provisioned-product-name", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
	{Name: "ProvisioningParameters", Flag: "provisioning-parameters", Type: "[]types.UpdateProvisioningParameter", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*types.ProvisioningArtifactProperties", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_create_service_action = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "map[string]string", Required: true},
	{Name: "DefinitionType", Flag: "definition-type", Type: "types.ServiceActionDefinitionType", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_tag_option = []leanruntime.Field{
	{Name: "Key", Flag: "key", Type: "*string", Required: true},
	{Name: "Value", Flag: "value", Type: "*string", Required: true},
}

var fields_delete_constraint = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_portfolio_share = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "OrganizationNode", Flag: "organization-node", Type: "*types.OrganizationNode", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
}

var fields_delete_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_provisioned_product_plan = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IgnoreErrors", Flag: "ignore-errors", Type: "bool", Required: false},
	{Name: "PlanId", Flag: "plan-id", Type: "*string", Required: true},
}

var fields_delete_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
}

var fields_delete_service_action = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
}

var fields_delete_tag_option = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_constraint = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_copy_product_status = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "CopyProductToken", Flag: "copy-product-token", Type: "*string", Required: true},
}

var fields_describe_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_portfolio_share_status = []leanruntime.Field{
	{Name: "PortfolioShareToken", Flag: "portfolio-share-token", Type: "*string", Required: true},
}

var fields_describe_portfolio_shares = []leanruntime.Field{
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.DescribePortfolioShareType", Required: true},
}

var fields_describe_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_describe_product_as_admin = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "SourcePortfolioId", Flag: "source-portfolio-id", Type: "*string", Required: false},
}

var fields_describe_product_view = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_provisioned_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_describe_provisioned_product_plan = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PlanId", Flag: "plan-id", Type: "*string", Required: true},
}

var fields_describe_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IncludeProvisioningArtifactParameters", Flag: "include-provisioning-artifact-parameters", Type: "bool", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: false},
	{Name: "ProductName", Flag: "product-name", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactName", Flag: "provisioning-artifact-name", Type: "*string", Required: false},
	{Name: "Verbose", Flag: "verbose", Type: "bool", Required: false},
}

var fields_describe_provisioning_parameters = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PathId", Flag: "path-id", Type: "*string", Required: false},
	{Name: "PathName", Flag: "path-name", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: false},
	{Name: "ProductName", Flag: "product-name", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactName", Flag: "provisioning-artifact-name", Type: "*string", Required: false},
}

var fields_describe_record = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_describe_service_action = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_describe_service_action_execution_parameters = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: true},
	{Name: "ServiceActionId", Flag: "service-action-id", Type: "*string", Required: true},
}

var fields_describe_tag_option = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_disable_aws_organizations_access = []leanruntime.Field{}

var fields_disassociate_budget_from_resource = []leanruntime.Field{
	{Name: "BudgetName", Flag: "budget-name", Type: "*string", Required: true},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_disassociate_principal_from_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "PrincipalARN", Flag: "principal-arn", Type: "*string", Required: true},
	{Name: "PrincipalType", Flag: "principal-type", Type: "types.PrincipalType", Required: false},
}

var fields_disassociate_product_from_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_disassociate_service_action_from_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
	{Name: "ServiceActionId", Flag: "service-action-id", Type: "*string", Required: true},
}

var fields_disassociate_tag_option_from_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagOptionId", Flag: "tag-option-id", Type: "*string", Required: true},
}

var fields_enable_aws_organizations_access = []leanruntime.Field{}

var fields_execute_provisioned_product_plan = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "PlanId", Flag: "plan-id", Type: "*string", Required: true},
}

var fields_execute_provisioned_product_service_action = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "ExecuteToken", Flag: "execute-token", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "map[string][]string", Required: false},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: true},
	{Name: "ServiceActionId", Flag: "service-action-id", Type: "*string", Required: true},
}

var fields_get_aws_organizations_access_status = []leanruntime.Field{}

var fields_get_provisioned_product_outputs = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "OutputKeys", Flag: "output-keys", Type: "[]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: false},
	{Name: "ProvisionedProductName", Flag: "provisioned-product-name", Type: "*string", Required: false},
}

var fields_import_as_provisioned_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "PhysicalId", Flag: "physical-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisionedProductName", Flag: "provisioned-product-name", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
}

var fields_list_accepted_portfolio_shares = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioShareType", Flag: "portfolio-share-type", Type: "types.PortfolioShareType", Required: false},
}

var fields_list_budgets_for_resource = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_list_constraints_for_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: false},
}

var fields_list_launch_paths = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_list_organization_portfolio_access = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "OrganizationNodeType", Flag: "organization-node-type", Type: "types.OrganizationNodeType", Required: true},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
}

var fields_list_portfolio_access = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "OrganizationParentId", Flag: "organization-parent-id", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
}

var fields_list_portfolios = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_list_portfolios_for_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_list_principals_for_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
}

var fields_list_provisioned_product_plans = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccessLevelFilter", Flag: "access-level-filter", Type: "*types.AccessLevelFilter", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ProvisionProductId", Flag: "provision-product-id", Type: "*string", Required: false},
}

var fields_list_provisioning_artifacts = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
}

var fields_list_provisioning_artifacts_for_service_action = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ServiceActionId", Flag: "service-action-id", Type: "*string", Required: true},
}

var fields_list_record_history = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccessLevelFilter", Flag: "access-level-filter", Type: "*types.AccessLevelFilter", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "SearchFilter", Flag: "search-filter", Type: "*types.ListRecordHistorySearchFilter", Required: false},
}

var fields_list_resources_for_tag_option = []leanruntime.Field{
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ResourceType", Flag: "resource-type", Type: "*string", Required: false},
	{Name: "TagOptionId", Flag: "tag-option-id", Type: "*string", Required: true},
}

var fields_list_service_actions = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_list_service_actions_for_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
}

var fields_list_stack_instances_for_provisioned_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: true},
}

var fields_list_tag_options = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.ListTagOptionsFilters", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_notify_provision_product_engine_workflow_result = []leanruntime.Field{
	{Name: "FailureReason", Flag: "failure-reason", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.RecordOutput", Required: false},
	{Name: "RecordId", Flag: "record-id", Type: "*string", Required: true},
	{Name: "ResourceIdentifier", Flag: "resource-identifier", Type: "*types.EngineWorkflowResourceIdentifier", Required: false},
	{Name: "Status", Flag: "status", Type: "types.EngineWorkflowStatus", Required: true},
	{Name: "WorkflowToken", Flag: "workflow-token", Type: "*string", Required: true},
}

var fields_notify_terminate_provisioned_product_engine_workflow_result = []leanruntime.Field{
	{Name: "FailureReason", Flag: "failure-reason", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "RecordId", Flag: "record-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.EngineWorkflowStatus", Required: true},
	{Name: "WorkflowToken", Flag: "workflow-token", Type: "*string", Required: true},
}

var fields_notify_update_provisioned_product_engine_workflow_result = []leanruntime.Field{
	{Name: "FailureReason", Flag: "failure-reason", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "Outputs", Flag: "outputs", Type: "[]types.RecordOutput", Required: false},
	{Name: "RecordId", Flag: "record-id", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.EngineWorkflowStatus", Required: true},
	{Name: "WorkflowToken", Flag: "workflow-token", Type: "*string", Required: true},
}

var fields_provision_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "NotificationArns", Flag: "notification-arns", Type: "[]string", Required: false},
	{Name: "PathId", Flag: "path-id", Type: "*string", Required: false},
	{Name: "PathName", Flag: "path-name", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: false},
	{Name: "ProductName", Flag: "product-name", Type: "*string", Required: false},
	{Name: "ProvisionToken", Flag: "provision-token", Type: "*string", Required: true},
	{Name: "ProvisionedProductName", Flag: "provisioned-product-name", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactName", Flag: "provisioning-artifact-name", Type: "*string", Required: false},
	{Name: "ProvisioningParameters", Flag: "provisioning-parameters", Type: "[]types.ProvisioningParameter", Required: false},
	{Name: "ProvisioningPreferences", Flag: "provisioning-preferences", Type: "*types.ProvisioningPreferences", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_reject_portfolio_share = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "PortfolioShareType", Flag: "portfolio-share-type", Type: "types.PortfolioShareType", Required: false},
}

var fields_scan_provisioned_products = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccessLevelFilter", Flag: "access-level-filter", Type: "*types.AccessLevelFilter", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
}

var fields_search_products = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ProductViewSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_search_products_as_admin = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: false},
	{Name: "ProductSource", Flag: "product-source", Type: "types.ProductSource", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ProductViewSortBy", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_search_provisioned_products = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccessLevelFilter", Flag: "access-level-filter", Type: "*types.AccessLevelFilter", Required: false},
	{Name: "Filters", Flag: "filters", Type: "map[string][]string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "int32", Required: false},
	{Name: "PageToken", Flag: "page-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_terminate_provisioned_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IgnoreErrors", Flag: "ignore-errors", Type: "bool", Required: false},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: false},
	{Name: "ProvisionedProductName", Flag: "provisioned-product-name", Type: "*string", Required: false},
	{Name: "RetainPhysicalResources", Flag: "retain-physical-resources", Type: "bool", Required: false},
	{Name: "TerminateToken", Flag: "terminate-token", Type: "*string", Required: true},
}

var fields_update_constraint = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "*string", Required: false},
}

var fields_update_portfolio = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AddTags", Flag: "add-tags", Type: "[]types.Tag", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisplayName", Flag: "display-name", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "ProviderName", Flag: "provider-name", Type: "*string", Required: false},
	{Name: "RemoveTags", Flag: "remove-tags", Type: "[]string", Required: false},
}

var fields_update_portfolio_share = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: false},
	{Name: "OrganizationNode", Flag: "organization-node", Type: "*types.OrganizationNode", Required: false},
	{Name: "PortfolioId", Flag: "portfolio-id", Type: "*string", Required: true},
	{Name: "SharePrincipals", Flag: "share-principals", Type: "*bool", Required: false},
	{Name: "ShareTagOptions", Flag: "share-tag-options", Type: "*bool", Required: false},
}

var fields_update_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "AddTags", Flag: "add-tags", Type: "[]types.Tag", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Distributor", Flag: "distributor", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Owner", Flag: "owner", Type: "*string", Required: false},
	{Name: "RemoveTags", Flag: "remove-tags", Type: "[]string", Required: false},
	{Name: "SourceConnection", Flag: "source-connection", Type: "*types.SourceConnection", Required: false},
	{Name: "SupportDescription", Flag: "support-description", Type: "*string", Required: false},
	{Name: "SupportEmail", Flag: "support-email", Type: "*string", Required: false},
	{Name: "SupportUrl", Flag: "support-url", Type: "*string", Required: false},
}

var fields_update_provisioned_product = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "PathId", Flag: "path-id", Type: "*string", Required: false},
	{Name: "PathName", Flag: "path-name", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: false},
	{Name: "ProductName", Flag: "product-name", Type: "*string", Required: false},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: false},
	{Name: "ProvisionedProductName", Flag: "provisioned-product-name", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: false},
	{Name: "ProvisioningArtifactName", Flag: "provisioning-artifact-name", Type: "*string", Required: false},
	{Name: "ProvisioningParameters", Flag: "provisioning-parameters", Type: "[]types.UpdateProvisioningParameter", Required: false},
	{Name: "ProvisioningPreferences", Flag: "provisioning-preferences", Type: "*types.UpdateProvisioningPreferences", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "UpdateToken", Flag: "update-token", Type: "*string", Required: true},
}

var fields_update_provisioned_product_properties = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: true},
	{Name: "ProvisionedProductId", Flag: "provisioned-product-id", Type: "*string", Required: true},
	{Name: "ProvisionedProductProperties", Flag: "provisioned-product-properties", Type: "map[string]string", Required: true},
}

var fields_update_provisioning_artifact = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Active", Flag: "active", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Guidance", Flag: "guidance", Type: "types.ProvisioningArtifactGuidance", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProductId", Flag: "product-id", Type: "*string", Required: true},
	{Name: "ProvisioningArtifactId", Flag: "provisioning-artifact-id", Type: "*string", Required: true},
}

var fields_update_service_action = []leanruntime.Field{
	{Name: "AcceptLanguage", Flag: "accept-language", Type: "*string", Required: false},
	{Name: "Definition", Flag: "definition", Type: "map[string]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_tag_option = []leanruntime.Field{
	{Name: "Active", Flag: "active", Type: "*bool", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Value", Flag: "value", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-portfolio-share": {
			Name:   "accept-portfolio-share",
			Fields: fields_accept_portfolio_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptPortfolioShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_portfolio_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptPortfolioShare(ctx, input)
			},
		},
		"associate-budget-with-resource": {
			Name:   "associate-budget-with-resource",
			Fields: fields_associate_budget_with_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateBudgetWithResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_budget_with_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateBudgetWithResource(ctx, input)
			},
		},
		"associate-principal-with-portfolio": {
			Name:   "associate-principal-with-portfolio",
			Fields: fields_associate_principal_with_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociatePrincipalWithPortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_principal_with_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociatePrincipalWithPortfolio(ctx, input)
			},
		},
		"associate-product-with-portfolio": {
			Name:   "associate-product-with-portfolio",
			Fields: fields_associate_product_with_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateProductWithPortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_product_with_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateProductWithPortfolio(ctx, input)
			},
		},
		"associate-service-action-with-provisioning-artifact": {
			Name:   "associate-service-action-with-provisioning-artifact",
			Fields: fields_associate_service_action_with_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateServiceActionWithProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_service_action_with_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateServiceActionWithProvisioningArtifact(ctx, input)
			},
		},
		"associate-tag-option-with-resource": {
			Name:   "associate-tag-option-with-resource",
			Fields: fields_associate_tag_option_with_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTagOptionWithResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_tag_option_with_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTagOptionWithResource(ctx, input)
			},
		},
		"batch-associate-service-action-with-provisioning-artifact": {
			Name:   "batch-associate-service-action-with-provisioning-artifact",
			Fields: fields_batch_associate_service_action_with_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchAssociateServiceActionWithProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_associate_service_action_with_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchAssociateServiceActionWithProvisioningArtifact(ctx, input)
			},
		},
		"batch-disassociate-service-action-from-provisioning-artifact": {
			Name:   "batch-disassociate-service-action-from-provisioning-artifact",
			Fields: fields_batch_disassociate_service_action_from_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDisassociateServiceActionFromProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_disassociate_service_action_from_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDisassociateServiceActionFromProvisioningArtifact(ctx, input)
			},
		},
		"copy-product": {
			Name:   "copy-product",
			Fields: fields_copy_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CopyProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_copy_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CopyProduct(ctx, input)
			},
		},
		"create-constraint": {
			Name:   "create-constraint",
			Fields: fields_create_constraint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConstraintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_constraint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConstraint(ctx, input)
			},
		},
		"create-portfolio": {
			Name:   "create-portfolio",
			Fields: fields_create_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePortfolio(ctx, input)
			},
		},
		"create-portfolio-share": {
			Name:   "create-portfolio-share",
			Fields: fields_create_portfolio_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePortfolioShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_portfolio_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePortfolioShare(ctx, input)
			},
		},
		"create-product": {
			Name:   "create-product",
			Fields: fields_create_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProduct(ctx, input)
			},
		},
		"create-provisioned-product-plan": {
			Name:   "create-provisioned-product-plan",
			Fields: fields_create_provisioned_product_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisionedProductPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioned_product_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisionedProductPlan(ctx, input)
			},
		},
		"create-provisioning-artifact": {
			Name:   "create-provisioning-artifact",
			Fields: fields_create_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProvisioningArtifact(ctx, input)
			},
		},
		"create-service-action": {
			Name:   "create-service-action",
			Fields: fields_create_service_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateServiceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_service_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateServiceAction(ctx, input)
			},
		},
		"create-tag-option": {
			Name:   "create-tag-option",
			Fields: fields_create_tag_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTagOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tag_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTagOption(ctx, input)
			},
		},
		"delete-constraint": {
			Name:   "delete-constraint",
			Fields: fields_delete_constraint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConstraintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_constraint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConstraint(ctx, input)
			},
		},
		"delete-portfolio": {
			Name:   "delete-portfolio",
			Fields: fields_delete_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortfolio(ctx, input)
			},
		},
		"delete-portfolio-share": {
			Name:   "delete-portfolio-share",
			Fields: fields_delete_portfolio_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePortfolioShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_portfolio_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePortfolioShare(ctx, input)
			},
		},
		"delete-product": {
			Name:   "delete-product",
			Fields: fields_delete_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProduct(ctx, input)
			},
		},
		"delete-provisioned-product-plan": {
			Name:   "delete-provisioned-product-plan",
			Fields: fields_delete_provisioned_product_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisionedProductPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioned_product_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisionedProductPlan(ctx, input)
			},
		},
		"delete-provisioning-artifact": {
			Name:   "delete-provisioning-artifact",
			Fields: fields_delete_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProvisioningArtifact(ctx, input)
			},
		},
		"delete-service-action": {
			Name:   "delete-service-action",
			Fields: fields_delete_service_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteServiceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_service_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteServiceAction(ctx, input)
			},
		},
		"delete-tag-option": {
			Name:   "delete-tag-option",
			Fields: fields_delete_tag_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tag_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTagOption(ctx, input)
			},
		},
		"describe-constraint": {
			Name:   "describe-constraint",
			Fields: fields_describe_constraint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeConstraintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_constraint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeConstraint(ctx, input)
			},
		},
		"describe-copy-product-status": {
			Name:   "describe-copy-product-status",
			Fields: fields_describe_copy_product_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCopyProductStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_copy_product_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCopyProductStatus(ctx, input)
			},
		},
		"describe-portfolio": {
			Name:   "describe-portfolio",
			Fields: fields_describe_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePortfolio(ctx, input)
			},
		},
		"describe-portfolio-share-status": {
			Name:   "describe-portfolio-share-status",
			Fields: fields_describe_portfolio_share_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePortfolioShareStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_portfolio_share_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePortfolioShareStatus(ctx, input)
			},
		},
		"describe-portfolio-shares": {
			Name:   "describe-portfolio-shares",
			Fields: fields_describe_portfolio_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePortfolioSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_describe_portfolio_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.DescribePortfolioShares(ctx, input)
				}
				var results []*svc.DescribePortfolioSharesOutput
				p := svc.NewDescribePortfolioSharesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"describe-product": {
			Name:   "describe-product",
			Fields: fields_describe_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProduct(ctx, input)
			},
		},
		"describe-product-as-admin": {
			Name:   "describe-product-as-admin",
			Fields: fields_describe_product_as_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProductAsAdminInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_product_as_admin, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProductAsAdmin(ctx, input)
			},
		},
		"describe-product-view": {
			Name:   "describe-product-view",
			Fields: fields_describe_product_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProductViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_product_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProductView(ctx, input)
			},
		},
		"describe-provisioned-product": {
			Name:   "describe-provisioned-product",
			Fields: fields_describe_provisioned_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProvisionedProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_provisioned_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProvisionedProduct(ctx, input)
			},
		},
		"describe-provisioned-product-plan": {
			Name:   "describe-provisioned-product-plan",
			Fields: fields_describe_provisioned_product_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProvisionedProductPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_provisioned_product_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProvisionedProductPlan(ctx, input)
			},
		},
		"describe-provisioning-artifact": {
			Name:   "describe-provisioning-artifact",
			Fields: fields_describe_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProvisioningArtifact(ctx, input)
			},
		},
		"describe-provisioning-parameters": {
			Name:   "describe-provisioning-parameters",
			Fields: fields_describe_provisioning_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeProvisioningParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_provisioning_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeProvisioningParameters(ctx, input)
			},
		},
		"describe-record": {
			Name:   "describe-record",
			Fields: fields_describe_record,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRecordInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_record, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRecord(ctx, input)
			},
		},
		"describe-service-action": {
			Name:   "describe-service-action",
			Fields: fields_describe_service_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceAction(ctx, input)
			},
		},
		"describe-service-action-execution-parameters": {
			Name:   "describe-service-action-execution-parameters",
			Fields: fields_describe_service_action_execution_parameters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeServiceActionExecutionParametersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_service_action_execution_parameters, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeServiceActionExecutionParameters(ctx, input)
			},
		},
		"describe-tag-option": {
			Name:   "describe-tag-option",
			Fields: fields_describe_tag_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTagOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tag_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTagOption(ctx, input)
			},
		},
		"disable-aws-organizations-access": {
			Name:   "disable-aws-organizations-access",
			Fields: fields_disable_aws_organizations_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAWSOrganizationsAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_aws_organizations_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAWSOrganizationsAccess(ctx, input)
			},
		},
		"disassociate-budget-from-resource": {
			Name:   "disassociate-budget-from-resource",
			Fields: fields_disassociate_budget_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateBudgetFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_budget_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateBudgetFromResource(ctx, input)
			},
		},
		"disassociate-principal-from-portfolio": {
			Name:   "disassociate-principal-from-portfolio",
			Fields: fields_disassociate_principal_from_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociatePrincipalFromPortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_principal_from_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociatePrincipalFromPortfolio(ctx, input)
			},
		},
		"disassociate-product-from-portfolio": {
			Name:   "disassociate-product-from-portfolio",
			Fields: fields_disassociate_product_from_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateProductFromPortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_product_from_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateProductFromPortfolio(ctx, input)
			},
		},
		"disassociate-service-action-from-provisioning-artifact": {
			Name:   "disassociate-service-action-from-provisioning-artifact",
			Fields: fields_disassociate_service_action_from_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateServiceActionFromProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_service_action_from_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateServiceActionFromProvisioningArtifact(ctx, input)
			},
		},
		"disassociate-tag-option-from-resource": {
			Name:   "disassociate-tag-option-from-resource",
			Fields: fields_disassociate_tag_option_from_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTagOptionFromResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_tag_option_from_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTagOptionFromResource(ctx, input)
			},
		},
		"enable-aws-organizations-access": {
			Name:   "enable-aws-organizations-access",
			Fields: fields_enable_aws_organizations_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAWSOrganizationsAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_aws_organizations_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAWSOrganizationsAccess(ctx, input)
			},
		},
		"execute-provisioned-product-plan": {
			Name:   "execute-provisioned-product-plan",
			Fields: fields_execute_provisioned_product_plan,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteProvisionedProductPlanInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_provisioned_product_plan, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteProvisionedProductPlan(ctx, input)
			},
		},
		"execute-provisioned-product-service-action": {
			Name:   "execute-provisioned-product-service-action",
			Fields: fields_execute_provisioned_product_service_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExecuteProvisionedProductServiceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_execute_provisioned_product_service_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExecuteProvisionedProductServiceAction(ctx, input)
			},
		},
		"get-aws-organizations-access-status": {
			Name:   "get-aws-organizations-access-status",
			Fields: fields_get_aws_organizations_access_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAWSOrganizationsAccessStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_aws_organizations_access_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAWSOrganizationsAccessStatus(ctx, input)
			},
		},
		"get-provisioned-product-outputs": {
			Name:   "get-provisioned-product-outputs",
			Fields: fields_get_provisioned_product_outputs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProvisionedProductOutputsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_provisioned_product_outputs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetProvisionedProductOutputs(ctx, input)
				}
				var results []*svc.GetProvisionedProductOutputsOutput
				p := svc.NewGetProvisionedProductOutputsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"import-as-provisioned-product": {
			Name:   "import-as-provisioned-product",
			Fields: fields_import_as_provisioned_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportAsProvisionedProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_as_provisioned_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportAsProvisionedProduct(ctx, input)
			},
		},
		"list-accepted-portfolio-shares": {
			Name:   "list-accepted-portfolio-shares",
			Fields: fields_list_accepted_portfolio_shares,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAcceptedPortfolioSharesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accepted_portfolio_shares, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAcceptedPortfolioShares(ctx, input)
				}
				var results []*svc.ListAcceptedPortfolioSharesOutput
				p := svc.NewListAcceptedPortfolioSharesPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-budgets-for-resource": {
			Name:   "list-budgets-for-resource",
			Fields: fields_list_budgets_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListBudgetsForResourceInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_budgets_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListBudgetsForResource(ctx, input)
				}
				var results []*svc.ListBudgetsForResourceOutput
				p := svc.NewListBudgetsForResourcePaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-constraints-for-portfolio": {
			Name:   "list-constraints-for-portfolio",
			Fields: fields_list_constraints_for_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConstraintsForPortfolioInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_constraints_for_portfolio, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConstraintsForPortfolio(ctx, input)
				}
				var results []*svc.ListConstraintsForPortfolioOutput
				p := svc.NewListConstraintsForPortfolioPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-launch-paths": {
			Name:   "list-launch-paths",
			Fields: fields_list_launch_paths,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLaunchPathsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_launch_paths, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLaunchPaths(ctx, input)
				}
				var results []*svc.ListLaunchPathsOutput
				p := svc.NewListLaunchPathsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-organization-portfolio-access": {
			Name:   "list-organization-portfolio-access",
			Fields: fields_list_organization_portfolio_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationPortfolioAccessInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organization_portfolio_access, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationPortfolioAccess(ctx, input)
				}
				var results []*svc.ListOrganizationPortfolioAccessOutput
				p := svc.NewListOrganizationPortfolioAccessPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-portfolio-access": {
			Name:   "list-portfolio-access",
			Fields: fields_list_portfolio_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortfolioAccessInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_portfolio_access, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPortfolioAccess(ctx, input)
				}
				var results []*svc.ListPortfolioAccessOutput
				p := svc.NewListPortfolioAccessPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-portfolios": {
			Name:   "list-portfolios",
			Fields: fields_list_portfolios,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortfoliosInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_portfolios, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPortfolios(ctx, input)
				}
				var results []*svc.ListPortfoliosOutput
				p := svc.NewListPortfoliosPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-portfolios-for-product": {
			Name:   "list-portfolios-for-product",
			Fields: fields_list_portfolios_for_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPortfoliosForProductInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_portfolios_for_product, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPortfoliosForProduct(ctx, input)
				}
				var results []*svc.ListPortfoliosForProductOutput
				p := svc.NewListPortfoliosForProductPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-principals-for-portfolio": {
			Name:   "list-principals-for-portfolio",
			Fields: fields_list_principals_for_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPrincipalsForPortfolioInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_principals_for_portfolio, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrincipalsForPortfolio(ctx, input)
				}
				var results []*svc.ListPrincipalsForPortfolioOutput
				p := svc.NewListPrincipalsForPortfolioPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-provisioned-product-plans": {
			Name:   "list-provisioned-product-plans",
			Fields: fields_list_provisioned_product_plans,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisionedProductPlansInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_provisioned_product_plans, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProvisionedProductPlans(ctx, input)
			},
		},
		"list-provisioning-artifacts": {
			Name:   "list-provisioning-artifacts",
			Fields: fields_list_provisioning_artifacts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisioningArtifactsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_provisioning_artifacts, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProvisioningArtifacts(ctx, input)
			},
		},
		"list-provisioning-artifacts-for-service-action": {
			Name:   "list-provisioning-artifacts-for-service-action",
			Fields: fields_list_provisioning_artifacts_for_service_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProvisioningArtifactsForServiceActionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_provisioning_artifacts_for_service_action, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProvisioningArtifactsForServiceAction(ctx, input)
				}
				var results []*svc.ListProvisioningArtifactsForServiceActionOutput
				p := svc.NewListProvisioningArtifactsForServiceActionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-record-history": {
			Name:   "list-record-history",
			Fields: fields_list_record_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRecordHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_record_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListRecordHistory(ctx, input)
			},
		},
		"list-resources-for-tag-option": {
			Name:   "list-resources-for-tag-option",
			Fields: fields_list_resources_for_tag_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourcesForTagOptionInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_resources_for_tag_option, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListResourcesForTagOption(ctx, input)
				}
				var results []*svc.ListResourcesForTagOptionOutput
				p := svc.NewListResourcesForTagOptionPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-service-actions": {
			Name:   "list-service-actions",
			Fields: fields_list_service_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceActions(ctx, input)
				}
				var results []*svc.ListServiceActionsOutput
				p := svc.NewListServiceActionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-service-actions-for-provisioning-artifact": {
			Name:   "list-service-actions-for-provisioning-artifact",
			Fields: fields_list_service_actions_for_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListServiceActionsForProvisioningArtifactInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_service_actions_for_provisioning_artifact, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListServiceActionsForProvisioningArtifact(ctx, input)
				}
				var results []*svc.ListServiceActionsForProvisioningArtifactOutput
				p := svc.NewListServiceActionsForProvisioningArtifactPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"list-stack-instances-for-provisioned-product": {
			Name:   "list-stack-instances-for-provisioned-product",
			Fields: fields_list_stack_instances_for_provisioned_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStackInstancesForProvisionedProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_stack_instances_for_provisioned_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListStackInstancesForProvisionedProduct(ctx, input)
			},
		},
		"list-tag-options": {
			Name:   "list-tag-options",
			Fields: fields_list_tag_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagOptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tag_options, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagOptions(ctx, input)
				}
				var results []*svc.ListTagOptionsOutput
				p := svc.NewListTagOptionsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"notify-provision-product-engine-workflow-result": {
			Name:   "notify-provision-product-engine-workflow-result",
			Fields: fields_notify_provision_product_engine_workflow_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyProvisionProductEngineWorkflowResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_provision_product_engine_workflow_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyProvisionProductEngineWorkflowResult(ctx, input)
			},
		},
		"notify-terminate-provisioned-product-engine-workflow-result": {
			Name:   "notify-terminate-provisioned-product-engine-workflow-result",
			Fields: fields_notify_terminate_provisioned_product_engine_workflow_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyTerminateProvisionedProductEngineWorkflowResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_terminate_provisioned_product_engine_workflow_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyTerminateProvisionedProductEngineWorkflowResult(ctx, input)
			},
		},
		"notify-update-provisioned-product-engine-workflow-result": {
			Name:   "notify-update-provisioned-product-engine-workflow-result",
			Fields: fields_notify_update_provisioned_product_engine_workflow_result,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.NotifyUpdateProvisionedProductEngineWorkflowResultInput{}
				if _, err := leanruntime.ApplyInput(input, fields_notify_update_provisioned_product_engine_workflow_result, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.NotifyUpdateProvisionedProductEngineWorkflowResult(ctx, input)
			},
		},
		"provision-product": {
			Name:   "provision-product",
			Fields: fields_provision_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ProvisionProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_provision_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ProvisionProduct(ctx, input)
			},
		},
		"reject-portfolio-share": {
			Name:   "reject-portfolio-share",
			Fields: fields_reject_portfolio_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectPortfolioShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_portfolio_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectPortfolioShare(ctx, input)
			},
		},
		"scan-provisioned-products": {
			Name:   "scan-provisioned-products",
			Fields: fields_scan_provisioned_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ScanProvisionedProductsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_scan_provisioned_products, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ScanProvisionedProducts(ctx, input)
			},
		},
		"search-products": {
			Name:   "search-products",
			Fields: fields_search_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchProductsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_products, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchProducts(ctx, input)
				}
				var results []*svc.SearchProductsOutput
				p := svc.NewSearchProductsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"search-products-as-admin": {
			Name:   "search-products-as-admin",
			Fields: fields_search_products_as_admin,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchProductsAsAdminInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_products_as_admin, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchProductsAsAdmin(ctx, input)
				}
				var results []*svc.SearchProductsAsAdminOutput
				p := svc.NewSearchProductsAsAdminPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"search-provisioned-products": {
			Name:   "search-provisioned-products",
			Fields: fields_search_provisioned_products,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchProvisionedProductsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_provisioned_products, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchProvisionedProducts(ctx, input)
				}
				var results []*svc.SearchProvisionedProductsOutput
				p := svc.NewSearchProvisionedProductsPaginator(client, input)
				for p.HasMorePages() {
					resp, err := p.NextPage(ctx)
					if err != nil {
						return nil, err
					}
					results = append(results, resp)
				}
				return results, nil
			},
		},
		"terminate-provisioned-product": {
			Name:   "terminate-provisioned-product",
			Fields: fields_terminate_provisioned_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateProvisionedProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_provisioned_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateProvisionedProduct(ctx, input)
			},
		},
		"update-constraint": {
			Name:   "update-constraint",
			Fields: fields_update_constraint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConstraintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_constraint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConstraint(ctx, input)
			},
		},
		"update-portfolio": {
			Name:   "update-portfolio",
			Fields: fields_update_portfolio,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePortfolioInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_portfolio, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePortfolio(ctx, input)
			},
		},
		"update-portfolio-share": {
			Name:   "update-portfolio-share",
			Fields: fields_update_portfolio_share,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePortfolioShareInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_portfolio_share, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePortfolioShare(ctx, input)
			},
		},
		"update-product": {
			Name:   "update-product",
			Fields: fields_update_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProduct(ctx, input)
			},
		},
		"update-provisioned-product": {
			Name:   "update-provisioned-product",
			Fields: fields_update_provisioned_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProvisionedProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_provisioned_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProvisionedProduct(ctx, input)
			},
		},
		"update-provisioned-product-properties": {
			Name:   "update-provisioned-product-properties",
			Fields: fields_update_provisioned_product_properties,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProvisionedProductPropertiesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_provisioned_product_properties, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProvisionedProductProperties(ctx, input)
			},
		},
		"update-provisioning-artifact": {
			Name:   "update-provisioning-artifact",
			Fields: fields_update_provisioning_artifact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProvisioningArtifactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_provisioning_artifact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProvisioningArtifact(ctx, input)
			},
		},
		"update-service-action": {
			Name:   "update-service-action",
			Fields: fields_update_service_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceAction(ctx, input)
			},
		},
		"update-tag-option": {
			Name:   "update-tag-option",
			Fields: fields_update_tag_option,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTagOptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_tag_option, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTagOption(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("servicecatalog", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
