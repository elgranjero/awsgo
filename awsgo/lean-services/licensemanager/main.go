package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/licensemanager"
)

var fields_accept_grant = []leanruntime.Field{
	{Name: "GrantArn", Flag: "grant-arn", Type: "*string", Required: true},
}

var fields_check_in_license = []leanruntime.Field{
	{Name: "Beneficiary", Flag: "beneficiary", Type: "*string", Required: false},
	{Name: "LicenseConsumptionToken", Flag: "license-consumption-token", Type: "*string", Required: true},
}

var fields_checkout_borrow_license = []leanruntime.Field{
	{Name: "CheckoutMetadata", Flag: "checkout-metadata", Type: "[]types.Metadata", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "DigitalSignatureMethod", Flag: "digital-signature-method", Type: "types.DigitalSignatureMethod", Required: true},
	{Name: "Entitlements", Flag: "entitlements", Type: "[]types.EntitlementData", Required: true},
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: false},
}

var fields_checkout_license = []leanruntime.Field{
	{Name: "Beneficiary", Flag: "beneficiary", Type: "*string", Required: false},
	{Name: "CheckoutType", Flag: "checkout-type", Type: "types.CheckoutType", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Entitlements", Flag: "entitlements", Type: "[]types.EntitlementData", Required: true},
	{Name: "KeyFingerprint", Flag: "key-fingerprint", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: false},
	{Name: "ProductSKU", Flag: "product-sku", Type: "*string", Required: true},
}

var fields_create_grant = []leanruntime.Field{
	{Name: "AllowedOperations", Flag: "allowed-operations", Type: "[]types.AllowedOperation", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GrantName", Flag: "grant-name", Type: "*string", Required: true},
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: true},
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "Principals", Flag: "principals", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_grant_version = []leanruntime.Field{
	{Name: "AllowedOperations", Flag: "allowed-operations", Type: "[]types.AllowedOperation", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "GrantArn", Flag: "grant-arn", Type: "*string", Required: true},
	{Name: "GrantName", Flag: "grant-name", Type: "*string", Required: false},
	{Name: "Options", Flag: "options", Type: "*types.Options", Required: false},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.GrantStatus", Required: false},
	{Name: "StatusReason", Flag: "status-reason", Type: "*string", Required: false},
}

var fields_create_license = []leanruntime.Field{
	{Name: "Beneficiary", Flag: "beneficiary", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConsumptionConfiguration", Flag: "consumption-configuration", Type: "*types.ConsumptionConfiguration", Required: true},
	{Name: "Entitlements", Flag: "entitlements", Type: "[]types.Entitlement", Required: true},
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: true},
	{Name: "Issuer", Flag: "issuer", Type: "*types.Issuer", Required: true},
	{Name: "LicenseMetadata", Flag: "license-metadata", Type: "[]types.Metadata", Required: false},
	{Name: "LicenseName", Flag: "license-name", Type: "*string", Required: true},
	{Name: "ProductName", Flag: "product-name", Type: "*string", Required: true},
	{Name: "ProductSKU", Flag: "product-sku", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Validity", Flag: "validity", Type: "*types.DatetimeRange", Required: true},
}

var fields_create_license_asset_group = []leanruntime.Field{
	{Name: "AssociatedLicenseAssetRulesetARNs", Flag: "associated-license-asset-ruleset-arns", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LicenseAssetGroupConfigurations", Flag: "license-asset-group-configurations", Type: "[]types.LicenseAssetGroupConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Properties", Flag: "properties", Type: "[]types.LicenseAssetGroupProperty", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_license_asset_ruleset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Rules", Flag: "rules", Type: "[]types.LicenseAssetRule", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_license_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisassociateWhenNotFound", Flag: "disassociate-when-not-found", Type: "*bool", Required: false},
	{Name: "LicenseCount", Flag: "license-count", Type: "*int64", Required: false},
	{Name: "LicenseCountHardLimit", Flag: "license-count-hard-limit", Type: "*bool", Required: false},
	{Name: "LicenseCountingType", Flag: "license-counting-type", Type: "types.LicenseCountingType", Required: true},
	{Name: "LicenseExpiry", Flag: "license-expiry", Type: "*int64", Required: false},
	{Name: "LicenseRules", Flag: "license-rules", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProductInformationList", Flag: "product-information-list", Type: "[]types.ProductInformation", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_license_conversion_task_for_resource = []leanruntime.Field{
	{Name: "DestinationLicenseContext", Flag: "destination-license-context", Type: "*types.LicenseConversionContext", Required: true},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "SourceLicenseContext", Flag: "source-license-context", Type: "*types.LicenseConversionContext", Required: true},
}

var fields_create_license_manager_report_generator = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ReportContext", Flag: "report-context", Type: "*types.ReportContext", Required: true},
	{Name: "ReportFrequency", Flag: "report-frequency", Type: "*types.ReportFrequency", Required: true},
	{Name: "ReportGeneratorName", Flag: "report-generator-name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "[]types.ReportType", Required: true},
}

var fields_create_license_version = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConsumptionConfiguration", Flag: "consumption-configuration", Type: "*types.ConsumptionConfiguration", Required: true},
	{Name: "Entitlements", Flag: "entitlements", Type: "[]types.Entitlement", Required: true},
	{Name: "HomeRegion", Flag: "home-region", Type: "*string", Required: true},
	{Name: "Issuer", Flag: "issuer", Type: "*types.Issuer", Required: true},
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "LicenseMetadata", Flag: "license-metadata", Type: "[]types.Metadata", Required: false},
	{Name: "LicenseName", Flag: "license-name", Type: "*string", Required: true},
	{Name: "ProductName", Flag: "product-name", Type: "*string", Required: true},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.LicenseStatus", Required: true},
	{Name: "Validity", Flag: "validity", Type: "*types.DatetimeRange", Required: true},
}

var fields_create_token = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ExpirationInDays", Flag: "expiration-in-days", Type: "*int32", Required: false},
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "RoleArns", Flag: "role-arns", Type: "[]string", Required: false},
	{Name: "TokenProperties", Flag: "token-properties", Type: "[]string", Required: false},
}

var fields_delete_grant = []leanruntime.Field{
	{Name: "GrantArn", Flag: "grant-arn", Type: "*string", Required: true},
	{Name: "StatusReason", Flag: "status-reason", Type: "*string", Required: false},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_delete_license = []leanruntime.Field{
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "SourceVersion", Flag: "source-version", Type: "*string", Required: true},
}

var fields_delete_license_asset_group = []leanruntime.Field{
	{Name: "LicenseAssetGroupArn", Flag: "license-asset-group-arn", Type: "*string", Required: true},
}

var fields_delete_license_asset_ruleset = []leanruntime.Field{
	{Name: "LicenseAssetRulesetArn", Flag: "license-asset-ruleset-arn", Type: "*string", Required: true},
}

var fields_delete_license_configuration = []leanruntime.Field{
	{Name: "LicenseConfigurationArn", Flag: "license-configuration-arn", Type: "*string", Required: true},
}

var fields_delete_license_manager_report_generator = []leanruntime.Field{
	{Name: "LicenseManagerReportGeneratorArn", Flag: "license-manager-report-generator-arn", Type: "*string", Required: true},
}

var fields_delete_token = []leanruntime.Field{
	{Name: "TokenId", Flag: "token-id", Type: "*string", Required: true},
}

var fields_extend_license_consumption = []leanruntime.Field{
	{Name: "DryRun", Flag: "dry-run", Type: "bool", Required: false},
	{Name: "LicenseConsumptionToken", Flag: "license-consumption-token", Type: "*string", Required: true},
}

var fields_get_access_token = []leanruntime.Field{
	{Name: "Token", Flag: "token", Type: "*string", Required: true},
	{Name: "TokenProperties", Flag: "token-properties", Type: "[]string", Required: false},
}

var fields_get_grant = []leanruntime.Field{
	{Name: "GrantArn", Flag: "grant-arn", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_license = []leanruntime.Field{
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: false},
}

var fields_get_license_asset_group = []leanruntime.Field{
	{Name: "LicenseAssetGroupArn", Flag: "license-asset-group-arn", Type: "*string", Required: true},
}

var fields_get_license_asset_ruleset = []leanruntime.Field{
	{Name: "LicenseAssetRulesetArn", Flag: "license-asset-ruleset-arn", Type: "*string", Required: true},
}

var fields_get_license_configuration = []leanruntime.Field{
	{Name: "LicenseConfigurationArn", Flag: "license-configuration-arn", Type: "*string", Required: true},
}

var fields_get_license_conversion_task = []leanruntime.Field{
	{Name: "LicenseConversionTaskId", Flag: "license-conversion-task-id", Type: "*string", Required: true},
}

var fields_get_license_manager_report_generator = []leanruntime.Field{
	{Name: "LicenseManagerReportGeneratorArn", Flag: "license-manager-report-generator-arn", Type: "*string", Required: true},
}

var fields_get_license_usage = []leanruntime.Field{
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
}

var fields_get_service_settings = []leanruntime.Field{}

var fields_list_assets_for_license_asset_group = []leanruntime.Field{
	{Name: "AssetType", Flag: "asset-type", Type: "*string", Required: true},
	{Name: "LicenseAssetGroupArn", Flag: "license-asset-group-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_associations_for_license_configuration = []leanruntime.Field{
	{Name: "LicenseConfigurationArn", Flag: "license-configuration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_distributed_grants = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "GrantArns", Flag: "grant-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_failures_for_license_configuration_operations = []leanruntime.Field{
	{Name: "LicenseConfigurationArn", Flag: "license-configuration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_asset_groups = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_asset_rulesets = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ShowAWSManagedLicenseAssetRulesets", Flag: "show-aws-managed-license-asset-rulesets", Type: "bool", Required: false},
}

var fields_list_license_configurations = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseConfigurationArns", Flag: "license-configuration-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_configurations_for_organization = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseConfigurationArns", Flag: "license-configuration-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_conversion_tasks = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_manager_report_generators = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_license_specifications_for_resource = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_license_versions = []leanruntime.Field{
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_licenses = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseArns", Flag: "license-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_received_grants = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "GrantArns", Flag: "grant-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_received_grants_for_organization = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseArn", Flag: "license-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_received_licenses = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseArns", Flag: "license-arns", Type: "[]string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_received_licenses_for_organization = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_resource_inventory = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.InventoryFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tokens = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TokenIds", Flag: "token-ids", Type: "[]string", Required: false},
}

var fields_list_usage_for_license_configuration = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "[]types.Filter", Required: false},
	{Name: "LicenseConfigurationArn", Flag: "license-configuration-arn", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_reject_grant = []leanruntime.Field{
	{Name: "GrantArn", Flag: "grant-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_license_asset_group = []leanruntime.Field{
	{Name: "AssociatedLicenseAssetRulesetARNs", Flag: "associated-license-asset-ruleset-arns", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LicenseAssetGroupArn", Flag: "license-asset-group-arn", Type: "*string", Required: true},
	{Name: "LicenseAssetGroupConfigurations", Flag: "license-asset-group-configurations", Type: "[]types.LicenseAssetGroupConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Properties", Flag: "properties", Type: "[]types.LicenseAssetGroupProperty", Required: false},
	{Name: "Status", Flag: "status", Type: "types.LicenseAssetGroupStatus", Required: false},
}

var fields_update_license_asset_ruleset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LicenseAssetRulesetArn", Flag: "license-asset-ruleset-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Rules", Flag: "rules", Type: "[]types.LicenseAssetRule", Required: true},
}

var fields_update_license_configuration = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DisassociateWhenNotFound", Flag: "disassociate-when-not-found", Type: "*bool", Required: false},
	{Name: "LicenseConfigurationArn", Flag: "license-configuration-arn", Type: "*string", Required: true},
	{Name: "LicenseConfigurationStatus", Flag: "license-configuration-status", Type: "types.LicenseConfigurationStatus", Required: false},
	{Name: "LicenseCount", Flag: "license-count", Type: "*int64", Required: false},
	{Name: "LicenseCountHardLimit", Flag: "license-count-hard-limit", Type: "*bool", Required: false},
	{Name: "LicenseExpiry", Flag: "license-expiry", Type: "*int64", Required: false},
	{Name: "LicenseRules", Flag: "license-rules", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProductInformationList", Flag: "product-information-list", Type: "[]types.ProductInformation", Required: false},
}

var fields_update_license_manager_report_generator = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "LicenseManagerReportGeneratorArn", Flag: "license-manager-report-generator-arn", Type: "*string", Required: true},
	{Name: "ReportContext", Flag: "report-context", Type: "*types.ReportContext", Required: true},
	{Name: "ReportFrequency", Flag: "report-frequency", Type: "*types.ReportFrequency", Required: true},
	{Name: "ReportGeneratorName", Flag: "report-generator-name", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "[]types.ReportType", Required: true},
}

var fields_update_license_specifications_for_resource = []leanruntime.Field{
	{Name: "AddLicenseSpecifications", Flag: "add-license-specifications", Type: "[]types.LicenseSpecification", Required: false},
	{Name: "RemoveLicenseSpecifications", Flag: "remove-license-specifications", Type: "[]types.LicenseSpecification", Required: false},
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_update_service_settings = []leanruntime.Field{
	{Name: "EnableCrossAccountsDiscovery", Flag: "enable-cross-accounts-discovery", Type: "*bool", Required: false},
	{Name: "EnabledDiscoverySourceRegions", Flag: "enabled-discovery-source-regions", Type: "[]string", Required: false},
	{Name: "OrganizationConfiguration", Flag: "organization-configuration", Type: "*types.OrganizationConfiguration", Required: false},
	{Name: "S3BucketArn", Flag: "s3-bucket-arn", Type: "*string", Required: false},
	{Name: "SnsTopicArn", Flag: "sns-topic-arn", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-grant": {
			Name:   "accept-grant",
			Fields: fields_accept_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptGrant(ctx, input)
			},
		},
		"check-in-license": {
			Name:   "check-in-license",
			Fields: fields_check_in_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckInLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_in_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckInLicense(ctx, input)
			},
		},
		"checkout-borrow-license": {
			Name:   "checkout-borrow-license",
			Fields: fields_checkout_borrow_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckoutBorrowLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_checkout_borrow_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckoutBorrowLicense(ctx, input)
			},
		},
		"checkout-license": {
			Name:   "checkout-license",
			Fields: fields_checkout_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckoutLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_checkout_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckoutLicense(ctx, input)
			},
		},
		"create-grant": {
			Name:   "create-grant",
			Fields: fields_create_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGrant(ctx, input)
			},
		},
		"create-grant-version": {
			Name:   "create-grant-version",
			Fields: fields_create_grant_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGrantVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_grant_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGrantVersion(ctx, input)
			},
		},
		"create-license": {
			Name:   "create-license",
			Fields: fields_create_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicense(ctx, input)
			},
		},
		"create-license-asset-group": {
			Name:   "create-license-asset-group",
			Fields: fields_create_license_asset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseAssetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_asset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseAssetGroup(ctx, input)
			},
		},
		"create-license-asset-ruleset": {
			Name:   "create-license-asset-ruleset",
			Fields: fields_create_license_asset_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseAssetRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_asset_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseAssetRuleset(ctx, input)
			},
		},
		"create-license-configuration": {
			Name:   "create-license-configuration",
			Fields: fields_create_license_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseConfiguration(ctx, input)
			},
		},
		"create-license-conversion-task-for-resource": {
			Name:   "create-license-conversion-task-for-resource",
			Fields: fields_create_license_conversion_task_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseConversionTaskForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_conversion_task_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseConversionTaskForResource(ctx, input)
			},
		},
		"create-license-manager-report-generator": {
			Name:   "create-license-manager-report-generator",
			Fields: fields_create_license_manager_report_generator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseManagerReportGeneratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_manager_report_generator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseManagerReportGenerator(ctx, input)
			},
		},
		"create-license-version": {
			Name:   "create-license-version",
			Fields: fields_create_license_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateLicenseVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_license_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateLicenseVersion(ctx, input)
			},
		},
		"create-token": {
			Name:   "create-token",
			Fields: fields_create_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateToken(ctx, input)
			},
		},
		"delete-grant": {
			Name:   "delete-grant",
			Fields: fields_delete_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGrant(ctx, input)
			},
		},
		"delete-license": {
			Name:   "delete-license",
			Fields: fields_delete_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicense(ctx, input)
			},
		},
		"delete-license-asset-group": {
			Name:   "delete-license-asset-group",
			Fields: fields_delete_license_asset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseAssetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license_asset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicenseAssetGroup(ctx, input)
			},
		},
		"delete-license-asset-ruleset": {
			Name:   "delete-license-asset-ruleset",
			Fields: fields_delete_license_asset_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseAssetRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license_asset_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicenseAssetRuleset(ctx, input)
			},
		},
		"delete-license-configuration": {
			Name:   "delete-license-configuration",
			Fields: fields_delete_license_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicenseConfiguration(ctx, input)
			},
		},
		"delete-license-manager-report-generator": {
			Name:   "delete-license-manager-report-generator",
			Fields: fields_delete_license_manager_report_generator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteLicenseManagerReportGeneratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_license_manager_report_generator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteLicenseManagerReportGenerator(ctx, input)
			},
		},
		"delete-token": {
			Name:   "delete-token",
			Fields: fields_delete_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteToken(ctx, input)
			},
		},
		"extend-license-consumption": {
			Name:   "extend-license-consumption",
			Fields: fields_extend_license_consumption,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ExtendLicenseConsumptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_extend_license_consumption, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ExtendLicenseConsumption(ctx, input)
			},
		},
		"get-access-token": {
			Name:   "get-access-token",
			Fields: fields_get_access_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_access_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessToken(ctx, input)
			},
		},
		"get-grant": {
			Name:   "get-grant",
			Fields: fields_get_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGrant(ctx, input)
			},
		},
		"get-license": {
			Name:   "get-license",
			Fields: fields_get_license,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicense(ctx, input)
			},
		},
		"get-license-asset-group": {
			Name:   "get-license-asset-group",
			Fields: fields_get_license_asset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseAssetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_asset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseAssetGroup(ctx, input)
			},
		},
		"get-license-asset-ruleset": {
			Name:   "get-license-asset-ruleset",
			Fields: fields_get_license_asset_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseAssetRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_asset_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseAssetRuleset(ctx, input)
			},
		},
		"get-license-configuration": {
			Name:   "get-license-configuration",
			Fields: fields_get_license_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseConfiguration(ctx, input)
			},
		},
		"get-license-conversion-task": {
			Name:   "get-license-conversion-task",
			Fields: fields_get_license_conversion_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseConversionTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_conversion_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseConversionTask(ctx, input)
			},
		},
		"get-license-manager-report-generator": {
			Name:   "get-license-manager-report-generator",
			Fields: fields_get_license_manager_report_generator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseManagerReportGeneratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_manager_report_generator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseManagerReportGenerator(ctx, input)
			},
		},
		"get-license-usage": {
			Name:   "get-license-usage",
			Fields: fields_get_license_usage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLicenseUsageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_license_usage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLicenseUsage(ctx, input)
			},
		},
		"get-service-settings": {
			Name:   "get-service-settings",
			Fields: fields_get_service_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetServiceSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_service_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetServiceSettings(ctx, input)
			},
		},
		"list-assets-for-license-asset-group": {
			Name:   "list-assets-for-license-asset-group",
			Fields: fields_list_assets_for_license_asset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetsForLicenseAssetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_assets_for_license_asset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAssetsForLicenseAssetGroup(ctx, input)
			},
		},
		"list-associations-for-license-configuration": {
			Name:   "list-associations-for-license-configuration",
			Fields: fields_list_associations_for_license_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssociationsForLicenseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_associations_for_license_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListAssociationsForLicenseConfiguration(ctx, input)
			},
		},
		"list-distributed-grants": {
			Name:   "list-distributed-grants",
			Fields: fields_list_distributed_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDistributedGrantsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_distributed_grants, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListDistributedGrants(ctx, input)
			},
		},
		"list-failures-for-license-configuration-operations": {
			Name:   "list-failures-for-license-configuration-operations",
			Fields: fields_list_failures_for_license_configuration_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFailuresForLicenseConfigurationOperationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_failures_for_license_configuration_operations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListFailuresForLicenseConfigurationOperations(ctx, input)
			},
		},
		"list-license-asset-groups": {
			Name:   "list-license-asset-groups",
			Fields: fields_list_license_asset_groups,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseAssetGroupsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_asset_groups, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseAssetGroups(ctx, input)
			},
		},
		"list-license-asset-rulesets": {
			Name:   "list-license-asset-rulesets",
			Fields: fields_list_license_asset_rulesets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseAssetRulesetsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_asset_rulesets, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseAssetRulesets(ctx, input)
			},
		},
		"list-license-configurations": {
			Name:   "list-license-configurations",
			Fields: fields_list_license_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseConfigurationsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_configurations, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseConfigurations(ctx, input)
			},
		},
		"list-license-configurations-for-organization": {
			Name:   "list-license-configurations-for-organization",
			Fields: fields_list_license_configurations_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseConfigurationsForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_configurations_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseConfigurationsForOrganization(ctx, input)
			},
		},
		"list-license-conversion-tasks": {
			Name:   "list-license-conversion-tasks",
			Fields: fields_list_license_conversion_tasks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseConversionTasksInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_conversion_tasks, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseConversionTasks(ctx, input)
			},
		},
		"list-license-manager-report-generators": {
			Name:   "list-license-manager-report-generators",
			Fields: fields_list_license_manager_report_generators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseManagerReportGeneratorsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_manager_report_generators, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseManagerReportGenerators(ctx, input)
			},
		},
		"list-license-specifications-for-resource": {
			Name:   "list-license-specifications-for-resource",
			Fields: fields_list_license_specifications_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseSpecificationsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_specifications_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseSpecificationsForResource(ctx, input)
			},
		},
		"list-license-versions": {
			Name:   "list-license-versions",
			Fields: fields_list_license_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicenseVersionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_license_versions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenseVersions(ctx, input)
			},
		},
		"list-licenses": {
			Name:   "list-licenses",
			Fields: fields_list_licenses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLicensesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_licenses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListLicenses(ctx, input)
			},
		},
		"list-received-grants": {
			Name:   "list-received-grants",
			Fields: fields_list_received_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceivedGrantsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_received_grants, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReceivedGrants(ctx, input)
			},
		},
		"list-received-grants-for-organization": {
			Name:   "list-received-grants-for-organization",
			Fields: fields_list_received_grants_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceivedGrantsForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_received_grants_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReceivedGrantsForOrganization(ctx, input)
			},
		},
		"list-received-licenses": {
			Name:   "list-received-licenses",
			Fields: fields_list_received_licenses,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceivedLicensesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_received_licenses, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReceivedLicenses(ctx, input)
			},
		},
		"list-received-licenses-for-organization": {
			Name:   "list-received-licenses-for-organization",
			Fields: fields_list_received_licenses_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListReceivedLicensesForOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_received_licenses_for_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListReceivedLicensesForOrganization(ctx, input)
			},
		},
		"list-resource-inventory": {
			Name:   "list-resource-inventory",
			Fields: fields_list_resource_inventory,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListResourceInventoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_resource_inventory, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListResourceInventory(ctx, input)
			},
		},
		"list-tags-for-resource": {
			Name:   "list-tags-for-resource",
			Fields: fields_list_tags_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForResource(ctx, input)
			},
		},
		"list-tokens": {
			Name:   "list-tokens",
			Fields: fields_list_tokens,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTokensInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tokens, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTokens(ctx, input)
			},
		},
		"list-usage-for-license-configuration": {
			Name:   "list-usage-for-license-configuration",
			Fields: fields_list_usage_for_license_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListUsageForLicenseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_usage_for_license_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListUsageForLicenseConfiguration(ctx, input)
			},
		},
		"reject-grant": {
			Name:   "reject-grant",
			Fields: fields_reject_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectGrant(ctx, input)
			},
		},
		"tag-resource": {
			Name:   "tag-resource",
			Fields: fields_tag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_tag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TagResource(ctx, input)
			},
		},
		"untag-resource": {
			Name:   "untag-resource",
			Fields: fields_untag_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UntagResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_untag_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UntagResource(ctx, input)
			},
		},
		"update-license-asset-group": {
			Name:   "update-license-asset-group",
			Fields: fields_update_license_asset_group,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLicenseAssetGroupInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_license_asset_group, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLicenseAssetGroup(ctx, input)
			},
		},
		"update-license-asset-ruleset": {
			Name:   "update-license-asset-ruleset",
			Fields: fields_update_license_asset_ruleset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLicenseAssetRulesetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_license_asset_ruleset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLicenseAssetRuleset(ctx, input)
			},
		},
		"update-license-configuration": {
			Name:   "update-license-configuration",
			Fields: fields_update_license_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLicenseConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_license_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLicenseConfiguration(ctx, input)
			},
		},
		"update-license-manager-report-generator": {
			Name:   "update-license-manager-report-generator",
			Fields: fields_update_license_manager_report_generator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLicenseManagerReportGeneratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_license_manager_report_generator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLicenseManagerReportGenerator(ctx, input)
			},
		},
		"update-license-specifications-for-resource": {
			Name:   "update-license-specifications-for-resource",
			Fields: fields_update_license_specifications_for_resource,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateLicenseSpecificationsForResourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_license_specifications_for_resource, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateLicenseSpecificationsForResource(ctx, input)
			},
		},
		"update-service-settings": {
			Name:   "update-service-settings",
			Fields: fields_update_service_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateServiceSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_service_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateServiceSettings(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("licensemanager", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
