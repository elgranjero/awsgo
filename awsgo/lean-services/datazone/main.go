package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/datazone"
)

var fields_accept_predictions = []leanruntime.Field{
	{Name: "AcceptChoices", Flag: "accept-choices", Type: "[]types.AcceptChoice", Required: false},
	{Name: "AcceptRule", Flag: "accept-rule", Type: "*types.AcceptRule", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_accept_subscription_request = []leanruntime.Field{
	{Name: "AssetPermissions", Flag: "asset-permissions", Type: "[]types.AssetPermission", Required: false},
	{Name: "AssetScopes", Flag: "asset-scopes", Type: "[]types.AcceptedAssetScope", Required: false},
	{Name: "DecisionComment", Flag: "decision-comment", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_add_entity_owner = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.DataZoneEntityType", Required: true},
	{Name: "Owner", Flag: "owner", Type: "types.OwnerProperties", Required: true},
}

var fields_add_policy_grant = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Detail", Flag: "detail", Type: "types.PolicyGrantDetail", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TargetEntityType", Required: true},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.ManagedPolicyType", Required: true},
	{Name: "Principal", Flag: "principal", Type: "types.PolicyGrantPrincipal", Required: true},
}

var fields_associate_environment_role = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentRoleArn", Flag: "environment-role-arn", Type: "*string", Required: true},
}

var fields_associate_governed_terms = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.GovernedEntityType", Required: true},
	{Name: "GovernedGlossaryTerms", Flag: "governed-glossary-terms", Type: "[]string", Required: true},
}

var fields_batch_get_attributes_metadata = []leanruntime.Field{
	{Name: "AttributeIdentifiers", Flag: "attribute-identifiers", Type: "[]string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityRevision", Flag: "entity-revision", Type: "*string", Required: false},
	{Name: "EntityType", Flag: "entity-type", Type: "types.AttributeEntityType", Required: true},
}

var fields_batch_put_attributes_metadata = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "[]types.AttributeInput", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.AttributeEntityType", Required: true},
}

var fields_cancel_metadata_generation_run = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_cancel_subscription = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_create_account_pool = []leanruntime.Field{
	{Name: "AccountSource", Flag: "account-source", Type: "types.AccountSource", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResolutionStrategy", Flag: "resolution-strategy", Type: "types.ResolutionStrategy", Required: true},
}

var fields_create_asset = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "ExternalIdentifier", Flag: "external-identifier", Type: "*string", Required: false},
	{Name: "FormsInput", Flag: "forms-input", Type: "[]types.FormInput", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: true},
	{Name: "PredictionConfiguration", Flag: "prediction-configuration", Type: "*types.PredictionConfiguration", Required: false},
	{Name: "TypeIdentifier", Flag: "type-identifier", Type: "*string", Required: true},
	{Name: "TypeRevision", Flag: "type-revision", Type: "*string", Required: false},
}

var fields_create_asset_filter = []leanruntime.Field{
	{Name: "AssetIdentifier", Flag: "asset-identifier", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.AssetFilterConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_asset_revision = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FormsInput", Flag: "forms-input", Type: "[]types.FormInput", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PredictionConfiguration", Flag: "prediction-configuration", Type: "*types.PredictionConfiguration", Required: false},
	{Name: "TypeRevision", Flag: "type-revision", Type: "*string", Required: false},
}

var fields_create_asset_type = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FormsInput", Flag: "forms-input", Type: "map[string]types.FormEntryInput", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: true},
}

var fields_create_connection = []leanruntime.Field{
	{Name: "AwsLocation", Flag: "aws-location", Type: "*types.AwsLocation", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnableTrustedIdentityPropagation", Flag: "enable-trusted-identity-propagation", Type: "*bool", Required: false},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Props", Flag: "props", Type: "types.ConnectionPropertiesInput", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.ConnectionScope", Required: false},
}

var fields_create_data_product = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FormsInput", Flag: "forms-input", Type: "[]types.FormInput", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Items", Flag: "items", Type: "[]types.DataProductItem", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: true},
}

var fields_create_data_product_revision = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FormsInput", Flag: "forms-input", Type: "[]types.FormInput", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Items", Flag: "items", Type: "[]types.DataProductItem", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_data_source = []leanruntime.Field{
	{Name: "AssetFormsInput", Flag: "asset-forms-input", Type: "[]types.FormInput", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.DataSourceConfigurationInput", Required: false},
	{Name: "ConnectionIdentifier", Flag: "connection-identifier", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnableSetting", Flag: "enable-setting", Type: "types.EnableSetting", Required: false},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
	{Name: "PublishOnImport", Flag: "publish-on-import", Type: "*bool", Required: false},
	{Name: "Recommendation", Flag: "recommendation", Type: "*types.RecommendationConfiguration", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.ScheduleConfiguration", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainExecutionRole", Flag: "domain-execution-role", Type: "*string", Required: true},
	{Name: "DomainVersion", Flag: "domain-version", Type: "types.DomainVersion", Required: false},
	{Name: "KmsKeyIdentifier", Flag: "kms-key-identifier", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: false},
	{Name: "SingleSignOn", Flag: "single-sign-on", Type: "*types.SingleSignOn", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_domain_unit = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentDomainUnitIdentifier", Flag: "parent-domain-unit-identifier", Type: "*string", Required: true},
}

var fields_create_environment = []leanruntime.Field{
	{Name: "DeploymentOrder", Flag: "deployment-order", Type: "*int32", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentAccountIdentifier", Flag: "environment-account-identifier", Type: "*string", Required: false},
	{Name: "EnvironmentAccountRegion", Flag: "environment-account-region", Type: "*string", Required: false},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: false},
	{Name: "EnvironmentConfigurationId", Flag: "environment-configuration-id", Type: "*string", Required: false},
	{Name: "EnvironmentProfileIdentifier", Flag: "environment-profile-identifier", Type: "*string", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.EnvironmentParameter", Required: false},
}

var fields_create_environment_action = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Parameters", Flag: "parameters", Type: "types.ActionParameters", Required: true},
}

var fields_create_environment_blueprint = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProvisioningProperties", Flag: "provisioning-properties", Type: "types.ProvisioningProperties", Required: true},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.CustomParameter", Required: false},
}

var fields_create_environment_profile = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsAccountRegion", Flag: "aws-account-region", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.EnvironmentParameter", Required: false},
}

var fields_create_form_type = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Model", Flag: "model", Type: "types.Model", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.FormTypeStatus", Required: false},
}

var fields_create_glossary = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.GlossaryStatus", Required: false},
	{Name: "UsageRestrictions", Flag: "usage-restrictions", Type: "[]types.GlossaryUsageRestriction", Required: false},
}

var fields_create_glossary_term = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GlossaryIdentifier", Flag: "glossary-identifier", Type: "*string", Required: true},
	{Name: "LongDescription", Flag: "long-description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ShortDescription", Flag: "short-description", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.GlossaryTermStatus", Required: false},
	{Name: "TermRelations", Flag: "term-relations", Type: "*types.TermRelations", Required: false},
}

var fields_create_group_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
}

var fields_create_listing_change_set = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.ChangeAction", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityRevision", Flag: "entity-revision", Type: "*string", Required: false},
	{Name: "EntityType", Flag: "entity-type", Type: "types.EntityType", Required: true},
}

var fields_create_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "DomainUnitId", Flag: "domain-unit-id", Type: "*string", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectProfileId", Flag: "project-profile-id", Type: "*string", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "map[string]string", Required: false},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.EnvironmentConfigurationUserParameter", Required: false},
}

var fields_create_project_membership = []leanruntime.Field{
	{Name: "Designation", Flag: "designation", Type: "types.UserDesignation", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Member", Flag: "member", Type: "types.Member", Required: true},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
}

var fields_create_project_profile = []leanruntime.Field{
	{Name: "AllowCustomProjectResourceTags", Flag: "allow-custom-project-resource-tags", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "DomainUnitIdentifier", Flag: "domain-unit-identifier", Type: "*string", Required: false},
	{Name: "EnvironmentConfigurations", Flag: "environment-configurations", Type: "[]types.EnvironmentConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ProjectResourceTags", Flag: "project-resource-tags", Type: "[]types.ResourceTagParameter", Required: false},
	{Name: "ProjectResourceTagsDescription", Flag: "project-resource-tags-description", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.Status", Required: false},
}

var fields_create_rule = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.RuleAction", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Detail", Flag: "detail", Type: "types.RuleDetail", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Scope", Flag: "scope", Type: "*types.RuleScope", Required: true},
	{Name: "Target", Flag: "target", Type: "types.RuleTarget", Required: true},
}

var fields_create_subscription_grant = []leanruntime.Field{
	{Name: "AssetTargetNames", Flag: "asset-target-names", Type: "[]types.AssetTargetNameMap", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "GrantedEntity", Flag: "granted-entity", Type: "types.GrantedEntityInput", Required: true},
	{Name: "SubscriptionTargetIdentifier", Flag: "subscription-target-identifier", Type: "*string", Required: false},
}

var fields_create_subscription_request = []leanruntime.Field{
	{Name: "AssetPermissions", Flag: "asset-permissions", Type: "[]types.AssetPermission", Required: false},
	{Name: "AssetScopes", Flag: "asset-scopes", Type: "[]types.AcceptedAssetScope", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MetadataForms", Flag: "metadata-forms", Type: "[]types.FormInput", Required: false},
	{Name: "RequestReason", Flag: "request-reason", Type: "*string", Required: true},
	{Name: "SubscribedListings", Flag: "subscribed-listings", Type: "[]types.SubscribedListingInput", Required: true},
	{Name: "SubscribedPrincipals", Flag: "subscribed-principals", Type: "[]types.SubscribedPrincipalInput", Required: true},
}

var fields_create_subscription_target = []leanruntime.Field{
	{Name: "ApplicableAssetTypes", Flag: "applicable-asset-types", Type: "[]string", Required: true},
	{Name: "AuthorizedPrincipals", Flag: "authorized-principals", Type: "[]string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "ManageAccessRole", Flag: "manage-access-role", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: false},
	{Name: "SubscriptionGrantCreationMode", Flag: "subscription-grant-creation-mode", Type: "types.SubscriptionGrantCreationMode", Required: false},
	{Name: "SubscriptionTargetConfig", Flag: "subscription-target-config", Type: "[]types.SubscriptionTargetForm", Required: true},
	{Name: "Type", Flag: "type", Type: "*string", Required: true},
}

var fields_create_user_profile = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "UserIdentifier", Flag: "user-identifier", Type: "*string", Required: true},
	{Name: "UserType", Flag: "user-type", Type: "types.UserType", Required: false},
}

var fields_delete_account_pool = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_asset = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_asset_filter = []leanruntime.Field{
	{Name: "AssetIdentifier", Flag: "asset-identifier", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_asset_type = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_connection = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_data_export_configuration = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
}

var fields_delete_data_product = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_data_source = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RetainPermissionsOnRevokeFailure", Flag: "retain-permissions-on-revoke-failure", Type: "*bool", Required: false},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "SkipDeletionCheck", Flag: "skip-deletion-check", Type: "*bool", Required: false},
}

var fields_delete_domain_unit = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_environment = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_environment_action = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_environment_blueprint = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_environment_blueprint_configuration = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: true},
}

var fields_delete_environment_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_form_type = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FormTypeIdentifier", Flag: "form-type-identifier", Type: "*string", Required: true},
}

var fields_delete_glossary = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_glossary_term = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_listing = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_project = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "SkipDeletionCheck", Flag: "skip-deletion-check", Type: "*bool", Required: false},
}

var fields_delete_project_membership = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Member", Flag: "member", Type: "types.Member", Required: true},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
}

var fields_delete_project_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_rule = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_subscription_grant = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_subscription_request = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_subscription_target = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_time_series_data_points = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TimeSeriesEntityType", Required: true},
	{Name: "FormName", Flag: "form-name", Type: "*string", Required: true},
}

var fields_disassociate_environment_role = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentRoleArn", Flag: "environment-role-arn", Type: "*string", Required: true},
}

var fields_disassociate_governed_terms = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.GovernedEntityType", Required: true},
	{Name: "GovernedGlossaryTerms", Flag: "governed-glossary-terms", Type: "[]string", Required: true},
}

var fields_get_account_pool = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_asset = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_get_asset_filter = []leanruntime.Field{
	{Name: "AssetIdentifier", Flag: "asset-identifier", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_asset_type = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_get_connection = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "WithSecret", Flag: "with-secret", Type: "*bool", Required: false},
}

var fields_get_data_export_configuration = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
}

var fields_get_data_product = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_get_data_source = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_data_source_run = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_domain = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_domain_unit = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_environment = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_environment_action = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_environment_blueprint = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_environment_blueprint_configuration = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: true},
}

var fields_get_environment_credentials = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
}

var fields_get_environment_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_form_type = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FormTypeIdentifier", Flag: "form-type-identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_get_glossary = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_glossary_term = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_group_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
}

var fields_get_iam_portal_login_url = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
}

var fields_get_job_run = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_lineage_event = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_lineage_node = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EventTimestamp", Flag: "event-timestamp", Type: "*time.Time", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_listing = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ListingRevision", Flag: "listing-revision", Type: "*string", Required: false},
}

var fields_get_metadata_generation_run = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.MetadataGenerationRunType", Required: false},
}

var fields_get_project = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_project_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_rule = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_get_subscription = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_subscription_grant = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_subscription_request_details = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_subscription_target = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_time_series_data_point = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TimeSeriesEntityType", Required: true},
	{Name: "FormName", Flag: "form-name", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_user_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Type", Flag: "type", Type: "types.UserProfileType", Required: false},
	{Name: "UserIdentifier", Flag: "user-identifier", Type: "*string", Required: true},
}

var fields_list_account_pools = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortFieldAccountPool", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_accounts_in_account_pool = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_asset_filters = []leanruntime.Field{
	{Name: "AssetIdentifier", Flag: "asset-identifier", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.FilterStatus", Required: false},
}

var fields_list_asset_revisions = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_connections = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "types.ConnectionScope", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortFieldConnection", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ConnectionType", Required: false},
}

var fields_list_data_product_revisions = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_data_source_run_activities = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.DataAssetActivityStatus", Required: false},
}

var fields_list_data_source_runs = []leanruntime.Field{
	{Name: "DataSourceIdentifier", Flag: "data-source-identifier", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.DataSourceRunStatus", Required: false},
}

var fields_list_data_sources = []leanruntime.Field{
	{Name: "ConnectionIdentifier", Flag: "connection-identifier", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.DataSourceStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "*string", Required: false},
}

var fields_list_domain_units_for_parent = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentDomainUnitIdentifier", Flag: "parent-domain-unit-identifier", Type: "*string", Required: true},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.DomainStatus", Required: false},
}

var fields_list_entity_owners = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.DataZoneEntityType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_actions = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_blueprint_configurations = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_blueprints = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Managed", Flag: "managed", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_environment_profiles = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsAccountRegion", Flag: "aws-account-region", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: false},
}

var fields_list_environments = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsAccountRegion", Flag: "aws-account-region", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: false},
	{Name: "EnvironmentProfileIdentifier", Flag: "environment-profile-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.EnvironmentStatus", Required: false},
}

var fields_list_job_runs = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "JobIdentifier", Flag: "job-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.JobRunStatus", Required: false},
}

var fields_list_lineage_events = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProcessingStatus", Flag: "processing-status", Type: "types.LineageEventProcessingStatus", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "TimestampAfter", Flag: "timestamp-after", Type: "*time.Time", Required: false},
	{Name: "TimestampBefore", Flag: "timestamp-before", Type: "*time.Time", Required: false},
}

var fields_list_lineage_node_history = []leanruntime.Field{
	{Name: "Direction", Flag: "direction", Type: "types.EdgeDirection", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EventTimestampGTE", Flag: "event-timestamp-gte", Type: "*time.Time", Required: false},
	{Name: "EventTimestampLTE", Flag: "event-timestamp-lte", Type: "*time.Time", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_metadata_generation_runs = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MetadataGenerationRunStatus", Required: false},
	{Name: "TargetIdentifier", Flag: "target-identifier", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.MetadataGenerationRunType", Required: false},
}

var fields_list_notifications = []leanruntime.Field{
	{Name: "AfterTimestamp", Flag: "after-timestamp", Type: "*time.Time", Required: false},
	{Name: "BeforeTimestamp", Flag: "before-timestamp", Type: "*time.Time", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Subjects", Flag: "subjects", Type: "[]string", Required: false},
	{Name: "TaskStatus", Flag: "task-status", Type: "types.TaskStatus", Required: false},
	{Name: "Type", Flag: "type", Type: "types.NotificationType", Required: true},
}

var fields_list_policy_grants = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TargetEntityType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.ManagedPolicyType", Required: true},
}

var fields_list_project_memberships = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectIdentifier", Flag: "project-identifier", Type: "*string", Required: true},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortFieldProject", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_project_profiles = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortFieldProject", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_projects = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "UserIdentifier", Flag: "user-identifier", Type: "*string", Required: false},
}

var fields_list_rules = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.RuleAction", Required: false},
	{Name: "AssetTypes", Flag: "asset-types", Type: "[]string", Required: false},
	{Name: "DataProduct", Flag: "data-product", Type: "*bool", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "IncludeCascaded", Flag: "include-cascaded", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProjectIds", Flag: "project-ids", Type: "[]string", Required: false},
	{Name: "RuleType", Flag: "rule-type", Type: "types.RuleType", Required: false},
	{Name: "TargetIdentifier", Flag: "target-identifier", Type: "*string", Required: true},
	{Name: "TargetType", Flag: "target-type", Type: "types.RuleTargetType", Required: true},
}

var fields_list_subscription_grants = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentId", Flag: "environment-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwningGroupId", Flag: "owning-group-id", Type: "*string", Required: false},
	{Name: "OwningIamPrincipalArn", Flag: "owning-iam-principal-arn", Type: "*string", Required: false},
	{Name: "OwningProjectId", Flag: "owning-project-id", Type: "*string", Required: false},
	{Name: "OwningUserId", Flag: "owning-user-id", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "SubscribedListingId", Flag: "subscribed-listing-id", Type: "*string", Required: false},
	{Name: "SubscriptionId", Flag: "subscription-id", Type: "*string", Required: false},
	{Name: "SubscriptionTargetId", Flag: "subscription-target-id", Type: "*string", Required: false},
}

var fields_list_subscription_requests = []leanruntime.Field{
	{Name: "ApproverProjectId", Flag: "approver-project-id", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwningGroupId", Flag: "owning-group-id", Type: "*string", Required: false},
	{Name: "OwningIamPrincipalArn", Flag: "owning-iam-principal-arn", Type: "*string", Required: false},
	{Name: "OwningProjectId", Flag: "owning-project-id", Type: "*string", Required: false},
	{Name: "OwningUserId", Flag: "owning-user-id", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.SubscriptionRequestStatus", Required: false},
	{Name: "SubscribedListingId", Flag: "subscribed-listing-id", Type: "*string", Required: false},
}

var fields_list_subscription_targets = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
}

var fields_list_subscriptions = []leanruntime.Field{
	{Name: "ApproverProjectId", Flag: "approver-project-id", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwningGroupId", Flag: "owning-group-id", Type: "*string", Required: false},
	{Name: "OwningIamPrincipalArn", Flag: "owning-iam-principal-arn", Type: "*string", Required: false},
	{Name: "OwningProjectId", Flag: "owning-project-id", Type: "*string", Required: false},
	{Name: "OwningUserId", Flag: "owning-user-id", Type: "*string", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.SortKey", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.SubscriptionStatus", Required: false},
	{Name: "SubscribedListingId", Flag: "subscribed-listing-id", Type: "*string", Required: false},
	{Name: "SubscriptionRequestIdentifier", Flag: "subscription-request-identifier", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_time_series_data_points = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EndedAt", Flag: "ended-at", Type: "*time.Time", Required: false},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TimeSeriesEntityType", Required: true},
	{Name: "FormName", Flag: "form-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartedAt", Flag: "started-at", Type: "*time.Time", Required: false},
}

var fields_post_lineage_event = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Event", Flag: "event", Type: "[]byte", Required: true},
}

var fields_post_time_series_data_points = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TimeSeriesEntityType", Required: true},
	{Name: "Forms", Flag: "forms", Type: "[]types.TimeSeriesDataPointFormInput", Required: true},
}

var fields_put_data_export_configuration = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnableExport", Flag: "enable-export", Type: "*bool", Required: true},
	{Name: "EncryptionConfiguration", Flag: "encryption-configuration", Type: "*types.EncryptionConfiguration", Required: false},
}

var fields_put_environment_blueprint_configuration = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnabledRegions", Flag: "enabled-regions", Type: "[]string", Required: true},
	{Name: "EnvironmentBlueprintIdentifier", Flag: "environment-blueprint-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentRolePermissionBoundary", Flag: "environment-role-permission-boundary", Type: "*string", Required: false},
	{Name: "GlobalParameters", Flag: "global-parameters", Type: "map[string]string", Required: false},
	{Name: "ManageAccessRoleArn", Flag: "manage-access-role-arn", Type: "*string", Required: false},
	{Name: "ProvisioningConfigurations", Flag: "provisioning-configurations", Type: "[]types.ProvisioningConfiguration", Required: false},
	{Name: "ProvisioningRoleArn", Flag: "provisioning-role-arn", Type: "*string", Required: false},
	{Name: "RegionalParameters", Flag: "regional-parameters", Type: "map[string]map[string]string", Required: false},
}

var fields_query_graph = []leanruntime.Field{
	{Name: "AdditionalAttributes", Flag: "additional-attributes", Type: "*types.AdditionalAttributes", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Match", Flag: "match", Type: "[]types.MatchClause", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_reject_predictions = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RejectChoices", Flag: "reject-choices", Type: "[]types.RejectChoice", Required: false},
	{Name: "RejectRule", Flag: "reject-rule", Type: "*types.RejectRule", Required: false},
	{Name: "Revision", Flag: "revision", Type: "*string", Required: false},
}

var fields_reject_subscription_request = []leanruntime.Field{
	{Name: "DecisionComment", Flag: "decision-comment", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_remove_entity_owner = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.DataZoneEntityType", Required: true},
	{Name: "Owner", Flag: "owner", Type: "types.OwnerProperties", Required: true},
}

var fields_remove_policy_grant = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EntityIdentifier", Flag: "entity-identifier", Type: "*string", Required: true},
	{Name: "EntityType", Flag: "entity-type", Type: "types.TargetEntityType", Required: true},
	{Name: "GrantIdentifier", Flag: "grant-identifier", Type: "*string", Required: false},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.ManagedPolicyType", Required: true},
	{Name: "Principal", Flag: "principal", Type: "types.PolicyGrantPrincipal", Required: true},
}

var fields_revoke_subscription = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RetainPermissions", Flag: "retain-permissions", Type: "*bool", Required: false},
}

var fields_search = []leanruntime.Field{
	{Name: "AdditionalAttributes", Flag: "additional-attributes", Type: "[]types.SearchOutputAdditionalAttribute", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "types.FilterClause", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: false},
	{Name: "SearchIn", Flag: "search-in", Type: "[]types.SearchInItem", Required: false},
	{Name: "SearchScope", Flag: "search-scope", Type: "types.InventorySearchScope", Required: true},
	{Name: "SearchText", Flag: "search-text", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SearchSort", Required: false},
}

var fields_search_group_profiles = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GroupType", Flag: "group-type", Type: "types.GroupSearchType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchText", Flag: "search-text", Type: "*string", Required: false},
}

var fields_search_listings = []leanruntime.Field{
	{Name: "AdditionalAttributes", Flag: "additional-attributes", Type: "[]types.SearchOutputAdditionalAttribute", Required: false},
	{Name: "Aggregations", Flag: "aggregations", Type: "[]types.AggregationListItem", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "types.FilterClause", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchIn", Flag: "search-in", Type: "[]types.SearchInItem", Required: false},
	{Name: "SearchText", Flag: "search-text", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SearchSort", Required: false},
}

var fields_search_types = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Filters", Flag: "filters", Type: "types.FilterClause", Required: false},
	{Name: "Managed", Flag: "managed", Type: "*bool", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchIn", Flag: "search-in", Type: "[]types.SearchInItem", Required: false},
	{Name: "SearchScope", Flag: "search-scope", Type: "types.TypesSearchScope", Required: true},
	{Name: "SearchText", Flag: "search-text", Type: "*string", Required: false},
	{Name: "Sort", Flag: "sort", Type: "*types.SearchSort", Required: false},
}

var fields_search_user_profiles = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SearchText", Flag: "search-text", Type: "*string", Required: false},
	{Name: "UserType", Flag: "user-type", Type: "types.UserSearchType", Required: true},
}

var fields_start_data_source_run = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataSourceIdentifier", Flag: "data-source-identifier", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
}

var fields_start_metadata_generation_run = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "OwningProjectIdentifier", Flag: "owning-project-identifier", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "*types.MetadataGenerationRunTarget", Required: true},
	{Name: "Type", Flag: "type", Type: "types.MetadataGenerationRunType", Required: false},
	{Name: "Types", Flag: "types", Type: "[]types.MetadataGenerationRunType", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_account_pool = []leanruntime.Field{
	{Name: "AccountSource", Flag: "account-source", Type: "types.AccountSource", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ResolutionStrategy", Flag: "resolution-strategy", Type: "types.ResolutionStrategy", Required: false},
}

var fields_update_asset_filter = []leanruntime.Field{
	{Name: "AssetIdentifier", Flag: "asset-identifier", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.AssetFilterConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_connection = []leanruntime.Field{
	{Name: "AwsLocation", Flag: "aws-location", Type: "*types.AwsLocation", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Props", Flag: "props", Type: "types.ConnectionPropertiesPatch", Required: false},
}

var fields_update_data_source = []leanruntime.Field{
	{Name: "AssetFormsInput", Flag: "asset-forms-input", Type: "[]types.FormInput", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.DataSourceConfigurationInput", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnableSetting", Flag: "enable-setting", Type: "types.EnableSetting", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PublishOnImport", Flag: "publish-on-import", Type: "*bool", Required: false},
	{Name: "Recommendation", Flag: "recommendation", Type: "*types.RecommendationConfiguration", Required: false},
	{Name: "RetainPermissionsOnRevokeFailure", Flag: "retain-permissions-on-revoke-failure", Type: "*bool", Required: false},
	{Name: "Schedule", Flag: "schedule", Type: "*types.ScheduleConfiguration", Required: false},
}

var fields_update_domain = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainExecutionRole", Flag: "domain-execution-role", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ServiceRole", Flag: "service-role", Type: "*string", Required: false},
	{Name: "SingleSignOn", Flag: "single-sign-on", Type: "*types.SingleSignOn", Required: false},
}

var fields_update_domain_unit = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_environment = []leanruntime.Field{
	{Name: "BlueprintVersion", Flag: "blueprint-version", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.EnvironmentParameter", Required: false},
}

var fields_update_environment_action = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Parameters", Flag: "parameters", Type: "types.ActionParameters", Required: false},
}

var fields_update_environment_blueprint = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ProvisioningProperties", Flag: "provisioning-properties", Type: "types.ProvisioningProperties", Required: false},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.CustomParameter", Required: false},
}

var fields_update_environment_profile = []leanruntime.Field{
	{Name: "AwsAccountId", Flag: "aws-account-id", Type: "*string", Required: false},
	{Name: "AwsAccountRegion", Flag: "aws-account-region", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.EnvironmentParameter", Required: false},
}

var fields_update_glossary = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.GlossaryStatus", Required: false},
}

var fields_update_glossary_term = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GlossaryIdentifier", Flag: "glossary-identifier", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "LongDescription", Flag: "long-description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ShortDescription", Flag: "short-description", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.GlossaryTermStatus", Required: false},
	{Name: "TermRelations", Flag: "term-relations", Type: "*types.TermRelations", Required: false},
}

var fields_update_group_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "GroupIdentifier", Flag: "group-identifier", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.GroupProfileStatus", Required: true},
}

var fields_update_project = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "DomainUnitId", Flag: "domain-unit-id", Type: "*string", Required: false},
	{Name: "EnvironmentDeploymentDetails", Flag: "environment-deployment-details", Type: "*types.EnvironmentDeploymentDetails", Required: false},
	{Name: "GlossaryTerms", Flag: "glossary-terms", Type: "[]string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProjectProfileVersion", Flag: "project-profile-version", Type: "*string", Required: false},
	{Name: "ResourceTags", Flag: "resource-tags", Type: "map[string]string", Required: false},
	{Name: "UserParameters", Flag: "user-parameters", Type: "[]types.EnvironmentConfigurationUserParameter", Required: false},
}

var fields_update_project_profile = []leanruntime.Field{
	{Name: "AllowCustomProjectResourceTags", Flag: "allow-custom-project-resource-tags", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "DomainUnitIdentifier", Flag: "domain-unit-identifier", Type: "*string", Required: false},
	{Name: "EnvironmentConfigurations", Flag: "environment-configurations", Type: "[]types.EnvironmentConfiguration", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProjectResourceTags", Flag: "project-resource-tags", Type: "[]types.ResourceTagParameter", Required: false},
	{Name: "ProjectResourceTagsDescription", Flag: "project-resource-tags-description", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.Status", Required: false},
}

var fields_update_root_domain_unit_owner = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "CurrentOwner", Flag: "current-owner", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "NewOwner", Flag: "new-owner", Type: "*string", Required: true},
}

var fields_update_rule = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Detail", Flag: "detail", Type: "types.RuleDetail", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "IncludeChildDomainUnits", Flag: "include-child-domain-units", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Scope", Flag: "scope", Type: "*types.RuleScope", Required: false},
}

var fields_update_subscription_grant_status = []leanruntime.Field{
	{Name: "AssetIdentifier", Flag: "asset-identifier", Type: "*string", Required: true},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "FailureCause", Flag: "failure-cause", Type: "*types.FailureCause", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.SubscriptionGrantStatus", Required: true},
	{Name: "TargetName", Flag: "target-name", Type: "*string", Required: false},
}

var fields_update_subscription_request = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "RequestReason", Flag: "request-reason", Type: "*string", Required: true},
}

var fields_update_subscription_target = []leanruntime.Field{
	{Name: "ApplicableAssetTypes", Flag: "applicable-asset-types", Type: "[]string", Required: false},
	{Name: "AuthorizedPrincipals", Flag: "authorized-principals", Type: "[]string", Required: false},
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "EnvironmentIdentifier", Flag: "environment-identifier", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "ManageAccessRole", Flag: "manage-access-role", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Provider", Flag: "provider", Type: "*string", Required: false},
	{Name: "SubscriptionGrantCreationMode", Flag: "subscription-grant-creation-mode", Type: "types.SubscriptionGrantCreationMode", Required: false},
	{Name: "SubscriptionTargetConfig", Flag: "subscription-target-config", Type: "[]types.SubscriptionTargetForm", Required: false},
}

var fields_update_user_profile = []leanruntime.Field{
	{Name: "DomainIdentifier", Flag: "domain-identifier", Type: "*string", Required: true},
	{Name: "Status", Flag: "status", Type: "types.UserProfileStatus", Required: true},
	{Name: "Type", Flag: "type", Type: "types.UserProfileType", Required: false},
	{Name: "UserIdentifier", Flag: "user-identifier", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-predictions": {
			Name:   "accept-predictions",
			Fields: fields_accept_predictions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptPredictionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_predictions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptPredictions(ctx, input)
			},
		},
		"accept-subscription-request": {
			Name:   "accept-subscription-request",
			Fields: fields_accept_subscription_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptSubscriptionRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_subscription_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptSubscriptionRequest(ctx, input)
			},
		},
		"add-entity-owner": {
			Name:   "add-entity-owner",
			Fields: fields_add_entity_owner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddEntityOwnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_entity_owner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddEntityOwner(ctx, input)
			},
		},
		"add-policy-grant": {
			Name:   "add-policy-grant",
			Fields: fields_add_policy_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddPolicyGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_policy_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddPolicyGrant(ctx, input)
			},
		},
		"associate-environment-role": {
			Name:   "associate-environment-role",
			Fields: fields_associate_environment_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateEnvironmentRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_environment_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateEnvironmentRole(ctx, input)
			},
		},
		"associate-governed-terms": {
			Name:   "associate-governed-terms",
			Fields: fields_associate_governed_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateGovernedTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_governed_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateGovernedTerms(ctx, input)
			},
		},
		"batch-get-attributes-metadata": {
			Name:   "batch-get-attributes-metadata",
			Fields: fields_batch_get_attributes_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetAttributesMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_attributes_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetAttributesMetadata(ctx, input)
			},
		},
		"batch-put-attributes-metadata": {
			Name:   "batch-put-attributes-metadata",
			Fields: fields_batch_put_attributes_metadata,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutAttributesMetadataInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_attributes_metadata, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutAttributesMetadata(ctx, input)
			},
		},
		"cancel-metadata-generation-run": {
			Name:   "cancel-metadata-generation-run",
			Fields: fields_cancel_metadata_generation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelMetadataGenerationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_metadata_generation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelMetadataGenerationRun(ctx, input)
			},
		},
		"cancel-subscription": {
			Name:   "cancel-subscription",
			Fields: fields_cancel_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSubscription(ctx, input)
			},
		},
		"create-account-pool": {
			Name:   "create-account-pool",
			Fields: fields_create_account_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccountPool(ctx, input)
			},
		},
		"create-asset": {
			Name:   "create-asset",
			Fields: fields_create_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAsset(ctx, input)
			},
		},
		"create-asset-filter": {
			Name:   "create-asset-filter",
			Fields: fields_create_asset_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssetFilter(ctx, input)
			},
		},
		"create-asset-revision": {
			Name:   "create-asset-revision",
			Fields: fields_create_asset_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssetRevision(ctx, input)
			},
		},
		"create-asset-type": {
			Name:   "create-asset-type",
			Fields: fields_create_asset_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAssetTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_asset_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAssetType(ctx, input)
			},
		},
		"create-connection": {
			Name:   "create-connection",
			Fields: fields_create_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnection(ctx, input)
			},
		},
		"create-data-product": {
			Name:   "create-data-product",
			Fields: fields_create_data_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataProduct(ctx, input)
			},
		},
		"create-data-product-revision": {
			Name:   "create-data-product-revision",
			Fields: fields_create_data_product_revision,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataProductRevisionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_product_revision, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataProductRevision(ctx, input)
			},
		},
		"create-data-source": {
			Name:   "create-data-source",
			Fields: fields_create_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDataSource(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-domain-unit": {
			Name:   "create-domain-unit",
			Fields: fields_create_domain_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomainUnit(ctx, input)
			},
		},
		"create-environment": {
			Name:   "create-environment",
			Fields: fields_create_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironment(ctx, input)
			},
		},
		"create-environment-action": {
			Name:   "create-environment-action",
			Fields: fields_create_environment_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentAction(ctx, input)
			},
		},
		"create-environment-blueprint": {
			Name:   "create-environment-blueprint",
			Fields: fields_create_environment_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentBlueprint(ctx, input)
			},
		},
		"create-environment-profile": {
			Name:   "create-environment-profile",
			Fields: fields_create_environment_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEnvironmentProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_environment_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEnvironmentProfile(ctx, input)
			},
		},
		"create-form-type": {
			Name:   "create-form-type",
			Fields: fields_create_form_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFormTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_form_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFormType(ctx, input)
			},
		},
		"create-glossary": {
			Name:   "create-glossary",
			Fields: fields_create_glossary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlossaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_glossary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlossary(ctx, input)
			},
		},
		"create-glossary-term": {
			Name:   "create-glossary-term",
			Fields: fields_create_glossary_term,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGlossaryTermInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_glossary_term, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGlossaryTerm(ctx, input)
			},
		},
		"create-group-profile": {
			Name:   "create-group-profile",
			Fields: fields_create_group_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGroupProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_group_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGroupProfile(ctx, input)
			},
		},
		"create-listing-change-set": {
			Name:   "create-listing-change-set",
			Fields: fields_create_listing_change_set,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateListingChangeSetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_listing_change_set, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateListingChangeSet(ctx, input)
			},
		},
		"create-project": {
			Name:   "create-project",
			Fields: fields_create_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProject(ctx, input)
			},
		},
		"create-project-membership": {
			Name:   "create-project-membership",
			Fields: fields_create_project_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProjectMembership(ctx, input)
			},
		},
		"create-project-profile": {
			Name:   "create-project-profile",
			Fields: fields_create_project_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProjectProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_project_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProjectProfile(ctx, input)
			},
		},
		"create-rule": {
			Name:   "create-rule",
			Fields: fields_create_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRule(ctx, input)
			},
		},
		"create-subscription-grant": {
			Name:   "create-subscription-grant",
			Fields: fields_create_subscription_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriptionGrant(ctx, input)
			},
		},
		"create-subscription-request": {
			Name:   "create-subscription-request",
			Fields: fields_create_subscription_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriptionRequest(ctx, input)
			},
		},
		"create-subscription-target": {
			Name:   "create-subscription-target",
			Fields: fields_create_subscription_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSubscriptionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_subscription_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSubscriptionTarget(ctx, input)
			},
		},
		"create-user-profile": {
			Name:   "create-user-profile",
			Fields: fields_create_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateUserProfile(ctx, input)
			},
		},
		"delete-account-pool": {
			Name:   "delete-account-pool",
			Fields: fields_delete_account_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccountPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_account_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccountPool(ctx, input)
			},
		},
		"delete-asset": {
			Name:   "delete-asset",
			Fields: fields_delete_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAsset(ctx, input)
			},
		},
		"delete-asset-filter": {
			Name:   "delete-asset-filter",
			Fields: fields_delete_asset_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssetFilter(ctx, input)
			},
		},
		"delete-asset-type": {
			Name:   "delete-asset-type",
			Fields: fields_delete_asset_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAssetTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_asset_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAssetType(ctx, input)
			},
		},
		"delete-connection": {
			Name:   "delete-connection",
			Fields: fields_delete_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnection(ctx, input)
			},
		},
		"delete-data-export-configuration": {
			Name:   "delete-data-export-configuration",
			Fields: fields_delete_data_export_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataExportConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_export_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataExportConfiguration(ctx, input)
			},
		},
		"delete-data-product": {
			Name:   "delete-data-product",
			Fields: fields_delete_data_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataProduct(ctx, input)
			},
		},
		"delete-data-source": {
			Name:   "delete-data-source",
			Fields: fields_delete_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDataSource(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-domain-unit": {
			Name:   "delete-domain-unit",
			Fields: fields_delete_domain_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomainUnit(ctx, input)
			},
		},
		"delete-environment": {
			Name:   "delete-environment",
			Fields: fields_delete_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironment(ctx, input)
			},
		},
		"delete-environment-action": {
			Name:   "delete-environment-action",
			Fields: fields_delete_environment_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentAction(ctx, input)
			},
		},
		"delete-environment-blueprint": {
			Name:   "delete-environment-blueprint",
			Fields: fields_delete_environment_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentBlueprint(ctx, input)
			},
		},
		"delete-environment-blueprint-configuration": {
			Name:   "delete-environment-blueprint-configuration",
			Fields: fields_delete_environment_blueprint_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentBlueprintConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_blueprint_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentBlueprintConfiguration(ctx, input)
			},
		},
		"delete-environment-profile": {
			Name:   "delete-environment-profile",
			Fields: fields_delete_environment_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEnvironmentProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_environment_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEnvironmentProfile(ctx, input)
			},
		},
		"delete-form-type": {
			Name:   "delete-form-type",
			Fields: fields_delete_form_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFormTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_form_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFormType(ctx, input)
			},
		},
		"delete-glossary": {
			Name:   "delete-glossary",
			Fields: fields_delete_glossary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlossaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_glossary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlossary(ctx, input)
			},
		},
		"delete-glossary-term": {
			Name:   "delete-glossary-term",
			Fields: fields_delete_glossary_term,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGlossaryTermInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_glossary_term, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGlossaryTerm(ctx, input)
			},
		},
		"delete-listing": {
			Name:   "delete-listing",
			Fields: fields_delete_listing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteListingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_listing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteListing(ctx, input)
			},
		},
		"delete-project": {
			Name:   "delete-project",
			Fields: fields_delete_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProject(ctx, input)
			},
		},
		"delete-project-membership": {
			Name:   "delete-project-membership",
			Fields: fields_delete_project_membership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectMembershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project_membership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProjectMembership(ctx, input)
			},
		},
		"delete-project-profile": {
			Name:   "delete-project-profile",
			Fields: fields_delete_project_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProjectProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_project_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProjectProfile(ctx, input)
			},
		},
		"delete-rule": {
			Name:   "delete-rule",
			Fields: fields_delete_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRule(ctx, input)
			},
		},
		"delete-subscription-grant": {
			Name:   "delete-subscription-grant",
			Fields: fields_delete_subscription_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriptionGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscription_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriptionGrant(ctx, input)
			},
		},
		"delete-subscription-request": {
			Name:   "delete-subscription-request",
			Fields: fields_delete_subscription_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriptionRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscription_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriptionRequest(ctx, input)
			},
		},
		"delete-subscription-target": {
			Name:   "delete-subscription-target",
			Fields: fields_delete_subscription_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSubscriptionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_subscription_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSubscriptionTarget(ctx, input)
			},
		},
		"delete-time-series-data-points": {
			Name:   "delete-time-series-data-points",
			Fields: fields_delete_time_series_data_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTimeSeriesDataPointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_time_series_data_points, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTimeSeriesDataPoints(ctx, input)
			},
		},
		"disassociate-environment-role": {
			Name:   "disassociate-environment-role",
			Fields: fields_disassociate_environment_role,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateEnvironmentRoleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_environment_role, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateEnvironmentRole(ctx, input)
			},
		},
		"disassociate-governed-terms": {
			Name:   "disassociate-governed-terms",
			Fields: fields_disassociate_governed_terms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateGovernedTermsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_governed_terms, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateGovernedTerms(ctx, input)
			},
		},
		"get-account-pool": {
			Name:   "get-account-pool",
			Fields: fields_get_account_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccountPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_account_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccountPool(ctx, input)
			},
		},
		"get-asset": {
			Name:   "get-asset",
			Fields: fields_get_asset,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_asset, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAsset(ctx, input)
			},
		},
		"get-asset-filter": {
			Name:   "get-asset-filter",
			Fields: fields_get_asset_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_asset_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssetFilter(ctx, input)
			},
		},
		"get-asset-type": {
			Name:   "get-asset-type",
			Fields: fields_get_asset_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAssetTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_asset_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAssetType(ctx, input)
			},
		},
		"get-connection": {
			Name:   "get-connection",
			Fields: fields_get_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnection(ctx, input)
			},
		},
		"get-data-export-configuration": {
			Name:   "get-data-export-configuration",
			Fields: fields_get_data_export_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataExportConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_export_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataExportConfiguration(ctx, input)
			},
		},
		"get-data-product": {
			Name:   "get-data-product",
			Fields: fields_get_data_product,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataProductInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_product, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataProduct(ctx, input)
			},
		},
		"get-data-source": {
			Name:   "get-data-source",
			Fields: fields_get_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSource(ctx, input)
			},
		},
		"get-data-source-run": {
			Name:   "get-data-source-run",
			Fields: fields_get_data_source_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDataSourceRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_data_source_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDataSourceRun(ctx, input)
			},
		},
		"get-domain": {
			Name:   "get-domain",
			Fields: fields_get_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomain(ctx, input)
			},
		},
		"get-domain-unit": {
			Name:   "get-domain-unit",
			Fields: fields_get_domain_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainUnit(ctx, input)
			},
		},
		"get-environment": {
			Name:   "get-environment",
			Fields: fields_get_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironment(ctx, input)
			},
		},
		"get-environment-action": {
			Name:   "get-environment-action",
			Fields: fields_get_environment_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentAction(ctx, input)
			},
		},
		"get-environment-blueprint": {
			Name:   "get-environment-blueprint",
			Fields: fields_get_environment_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentBlueprint(ctx, input)
			},
		},
		"get-environment-blueprint-configuration": {
			Name:   "get-environment-blueprint-configuration",
			Fields: fields_get_environment_blueprint_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentBlueprintConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_blueprint_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentBlueprintConfiguration(ctx, input)
			},
		},
		"get-environment-credentials": {
			Name:   "get-environment-credentials",
			Fields: fields_get_environment_credentials,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentCredentialsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_credentials, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentCredentials(ctx, input)
			},
		},
		"get-environment-profile": {
			Name:   "get-environment-profile",
			Fields: fields_get_environment_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEnvironmentProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_environment_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEnvironmentProfile(ctx, input)
			},
		},
		"get-form-type": {
			Name:   "get-form-type",
			Fields: fields_get_form_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFormTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_form_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFormType(ctx, input)
			},
		},
		"get-glossary": {
			Name:   "get-glossary",
			Fields: fields_get_glossary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlossaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_glossary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlossary(ctx, input)
			},
		},
		"get-glossary-term": {
			Name:   "get-glossary-term",
			Fields: fields_get_glossary_term,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGlossaryTermInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_glossary_term, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGlossaryTerm(ctx, input)
			},
		},
		"get-group-profile": {
			Name:   "get-group-profile",
			Fields: fields_get_group_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGroupProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_group_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGroupProfile(ctx, input)
			},
		},
		"get-iam-portal-login-url": {
			Name:   "get-iam-portal-login-url",
			Fields: fields_get_iam_portal_login_url,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIamPortalLoginUrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_iam_portal_login_url, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIamPortalLoginUrl(ctx, input)
			},
		},
		"get-job-run": {
			Name:   "get-job-run",
			Fields: fields_get_job_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetJobRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_job_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetJobRun(ctx, input)
			},
		},
		"get-lineage-event": {
			Name:   "get-lineage-event",
			Fields: fields_get_lineage_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLineageEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lineage_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLineageEvent(ctx, input)
			},
		},
		"get-lineage-node": {
			Name:   "get-lineage-node",
			Fields: fields_get_lineage_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLineageNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_lineage_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLineageNode(ctx, input)
			},
		},
		"get-listing": {
			Name:   "get-listing",
			Fields: fields_get_listing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetListingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_listing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetListing(ctx, input)
			},
		},
		"get-metadata-generation-run": {
			Name:   "get-metadata-generation-run",
			Fields: fields_get_metadata_generation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMetadataGenerationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_metadata_generation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMetadataGenerationRun(ctx, input)
			},
		},
		"get-project": {
			Name:   "get-project",
			Fields: fields_get_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProject(ctx, input)
			},
		},
		"get-project-profile": {
			Name:   "get-project-profile",
			Fields: fields_get_project_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProjectProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_project_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProjectProfile(ctx, input)
			},
		},
		"get-rule": {
			Name:   "get-rule",
			Fields: fields_get_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRule(ctx, input)
			},
		},
		"get-subscription": {
			Name:   "get-subscription",
			Fields: fields_get_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscription(ctx, input)
			},
		},
		"get-subscription-grant": {
			Name:   "get-subscription-grant",
			Fields: fields_get_subscription_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionGrant(ctx, input)
			},
		},
		"get-subscription-request-details": {
			Name:   "get-subscription-request-details",
			Fields: fields_get_subscription_request_details,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionRequestDetailsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_request_details, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionRequestDetails(ctx, input)
			},
		},
		"get-subscription-target": {
			Name:   "get-subscription-target",
			Fields: fields_get_subscription_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubscriptionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subscription_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubscriptionTarget(ctx, input)
			},
		},
		"get-time-series-data-point": {
			Name:   "get-time-series-data-point",
			Fields: fields_get_time_series_data_point,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTimeSeriesDataPointInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_time_series_data_point, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTimeSeriesDataPoint(ctx, input)
			},
		},
		"get-user-profile": {
			Name:   "get-user-profile",
			Fields: fields_get_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUserProfile(ctx, input)
			},
		},
		"list-account-pools": {
			Name:   "list-account-pools",
			Fields: fields_list_account_pools,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountPoolsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_account_pools, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountPools(ctx, input)
				}
				var results []*svc.ListAccountPoolsOutput
				p := svc.NewListAccountPoolsPaginator(client, input)
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
		"list-accounts-in-account-pool": {
			Name:   "list-accounts-in-account-pool",
			Fields: fields_list_accounts_in_account_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsInAccountPoolInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts_in_account_pool, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountsInAccountPool(ctx, input)
				}
				var results []*svc.ListAccountsInAccountPoolOutput
				p := svc.NewListAccountsInAccountPoolPaginator(client, input)
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
		"list-asset-filters": {
			Name:   "list-asset-filters",
			Fields: fields_list_asset_filters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetFiltersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_filters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetFilters(ctx, input)
				}
				var results []*svc.ListAssetFiltersOutput
				p := svc.NewListAssetFiltersPaginator(client, input)
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
		"list-asset-revisions": {
			Name:   "list-asset-revisions",
			Fields: fields_list_asset_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAssetRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_asset_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAssetRevisions(ctx, input)
				}
				var results []*svc.ListAssetRevisionsOutput
				p := svc.NewListAssetRevisionsPaginator(client, input)
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
		"list-connections": {
			Name:   "list-connections",
			Fields: fields_list_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnections(ctx, input)
				}
				var results []*svc.ListConnectionsOutput
				p := svc.NewListConnectionsPaginator(client, input)
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
		"list-data-product-revisions": {
			Name:   "list-data-product-revisions",
			Fields: fields_list_data_product_revisions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataProductRevisionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_product_revisions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataProductRevisions(ctx, input)
				}
				var results []*svc.ListDataProductRevisionsOutput
				p := svc.NewListDataProductRevisionsPaginator(client, input)
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
		"list-data-source-run-activities": {
			Name:   "list-data-source-run-activities",
			Fields: fields_list_data_source_run_activities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourceRunActivitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_source_run_activities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSourceRunActivities(ctx, input)
				}
				var results []*svc.ListDataSourceRunActivitiesOutput
				p := svc.NewListDataSourceRunActivitiesPaginator(client, input)
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
		"list-data-source-runs": {
			Name:   "list-data-source-runs",
			Fields: fields_list_data_source_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourceRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_source_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSourceRuns(ctx, input)
				}
				var results []*svc.ListDataSourceRunsOutput
				p := svc.NewListDataSourceRunsPaginator(client, input)
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
		"list-data-sources": {
			Name:   "list-data-sources",
			Fields: fields_list_data_sources,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDataSourcesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_data_sources, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDataSources(ctx, input)
				}
				var results []*svc.ListDataSourcesOutput
				p := svc.NewListDataSourcesPaginator(client, input)
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
		"list-domain-units-for-parent": {
			Name:   "list-domain-units-for-parent",
			Fields: fields_list_domain_units_for_parent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainUnitsForParentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domain_units_for_parent, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomainUnitsForParent(ctx, input)
				}
				var results []*svc.ListDomainUnitsForParentOutput
				p := svc.NewListDomainUnitsForParentPaginator(client, input)
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
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-entity-owners": {
			Name:   "list-entity-owners",
			Fields: fields_list_entity_owners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEntityOwnersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_entity_owners, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEntityOwners(ctx, input)
				}
				var results []*svc.ListEntityOwnersOutput
				p := svc.NewListEntityOwnersPaginator(client, input)
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
		"list-environment-actions": {
			Name:   "list-environment-actions",
			Fields: fields_list_environment_actions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentActionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_actions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentActions(ctx, input)
				}
				var results []*svc.ListEnvironmentActionsOutput
				p := svc.NewListEnvironmentActionsPaginator(client, input)
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
		"list-environment-blueprint-configurations": {
			Name:   "list-environment-blueprint-configurations",
			Fields: fields_list_environment_blueprint_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentBlueprintConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_blueprint_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentBlueprintConfigurations(ctx, input)
				}
				var results []*svc.ListEnvironmentBlueprintConfigurationsOutput
				p := svc.NewListEnvironmentBlueprintConfigurationsPaginator(client, input)
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
		"list-environment-blueprints": {
			Name:   "list-environment-blueprints",
			Fields: fields_list_environment_blueprints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentBlueprintsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_blueprints, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentBlueprints(ctx, input)
				}
				var results []*svc.ListEnvironmentBlueprintsOutput
				p := svc.NewListEnvironmentBlueprintsPaginator(client, input)
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
		"list-environment-profiles": {
			Name:   "list-environment-profiles",
			Fields: fields_list_environment_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environment_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironmentProfiles(ctx, input)
				}
				var results []*svc.ListEnvironmentProfilesOutput
				p := svc.NewListEnvironmentProfilesPaginator(client, input)
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
		"list-environments": {
			Name:   "list-environments",
			Fields: fields_list_environments,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEnvironmentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_environments, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEnvironments(ctx, input)
				}
				var results []*svc.ListEnvironmentsOutput
				p := svc.NewListEnvironmentsPaginator(client, input)
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
		"list-job-runs": {
			Name:   "list-job-runs",
			Fields: fields_list_job_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListJobRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_job_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListJobRuns(ctx, input)
				}
				var results []*svc.ListJobRunsOutput
				p := svc.NewListJobRunsPaginator(client, input)
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
		"list-lineage-events": {
			Name:   "list-lineage-events",
			Fields: fields_list_lineage_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLineageEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lineage_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLineageEvents(ctx, input)
				}
				var results []*svc.ListLineageEventsOutput
				p := svc.NewListLineageEventsPaginator(client, input)
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
		"list-lineage-node-history": {
			Name:   "list-lineage-node-history",
			Fields: fields_list_lineage_node_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListLineageNodeHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_lineage_node_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListLineageNodeHistory(ctx, input)
				}
				var results []*svc.ListLineageNodeHistoryOutput
				p := svc.NewListLineageNodeHistoryPaginator(client, input)
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
		"list-metadata-generation-runs": {
			Name:   "list-metadata-generation-runs",
			Fields: fields_list_metadata_generation_runs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMetadataGenerationRunsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_metadata_generation_runs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMetadataGenerationRuns(ctx, input)
				}
				var results []*svc.ListMetadataGenerationRunsOutput
				p := svc.NewListMetadataGenerationRunsPaginator(client, input)
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
		"list-notifications": {
			Name:   "list-notifications",
			Fields: fields_list_notifications,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNotificationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_notifications, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNotifications(ctx, input)
				}
				var results []*svc.ListNotificationsOutput
				p := svc.NewListNotificationsPaginator(client, input)
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
		"list-policy-grants": {
			Name:   "list-policy-grants",
			Fields: fields_list_policy_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPolicyGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policy_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicyGrants(ctx, input)
				}
				var results []*svc.ListPolicyGrantsOutput
				p := svc.NewListPolicyGrantsPaginator(client, input)
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
		"list-project-memberships": {
			Name:   "list-project-memberships",
			Fields: fields_list_project_memberships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectMembershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_project_memberships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjectMemberships(ctx, input)
				}
				var results []*svc.ListProjectMembershipsOutput
				p := svc.NewListProjectMembershipsPaginator(client, input)
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
		"list-project-profiles": {
			Name:   "list-project-profiles",
			Fields: fields_list_project_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_project_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjectProfiles(ctx, input)
				}
				var results []*svc.ListProjectProfilesOutput
				p := svc.NewListProjectProfilesPaginator(client, input)
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
		"list-projects": {
			Name:   "list-projects",
			Fields: fields_list_projects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_projects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProjects(ctx, input)
				}
				var results []*svc.ListProjectsOutput
				p := svc.NewListProjectsPaginator(client, input)
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
		"list-rules": {
			Name:   "list-rules",
			Fields: fields_list_rules,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRulesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_rules, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRules(ctx, input)
				}
				var results []*svc.ListRulesOutput
				p := svc.NewListRulesPaginator(client, input)
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
		"list-subscription-grants": {
			Name:   "list-subscription-grants",
			Fields: fields_list_subscription_grants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionGrantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscription_grants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptionGrants(ctx, input)
				}
				var results []*svc.ListSubscriptionGrantsOutput
				p := svc.NewListSubscriptionGrantsPaginator(client, input)
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
		"list-subscription-requests": {
			Name:   "list-subscription-requests",
			Fields: fields_list_subscription_requests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionRequestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscription_requests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptionRequests(ctx, input)
				}
				var results []*svc.ListSubscriptionRequestsOutput
				p := svc.NewListSubscriptionRequestsPaginator(client, input)
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
		"list-subscription-targets": {
			Name:   "list-subscription-targets",
			Fields: fields_list_subscription_targets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionTargetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscription_targets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptionTargets(ctx, input)
				}
				var results []*svc.ListSubscriptionTargetsOutput
				p := svc.NewListSubscriptionTargetsPaginator(client, input)
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
		"list-subscriptions": {
			Name:   "list-subscriptions",
			Fields: fields_list_subscriptions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubscriptionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subscriptions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubscriptions(ctx, input)
				}
				var results []*svc.ListSubscriptionsOutput
				p := svc.NewListSubscriptionsPaginator(client, input)
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
		"list-time-series-data-points": {
			Name:   "list-time-series-data-points",
			Fields: fields_list_time_series_data_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTimeSeriesDataPointsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_time_series_data_points, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTimeSeriesDataPoints(ctx, input)
				}
				var results []*svc.ListTimeSeriesDataPointsOutput
				p := svc.NewListTimeSeriesDataPointsPaginator(client, input)
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
		"post-lineage-event": {
			Name:   "post-lineage-event",
			Fields: fields_post_lineage_event,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostLineageEventInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_lineage_event, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostLineageEvent(ctx, input)
			},
		},
		"post-time-series-data-points": {
			Name:   "post-time-series-data-points",
			Fields: fields_post_time_series_data_points,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PostTimeSeriesDataPointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_post_time_series_data_points, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PostTimeSeriesDataPoints(ctx, input)
			},
		},
		"put-data-export-configuration": {
			Name:   "put-data-export-configuration",
			Fields: fields_put_data_export_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDataExportConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_data_export_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDataExportConfiguration(ctx, input)
			},
		},
		"put-environment-blueprint-configuration": {
			Name:   "put-environment-blueprint-configuration",
			Fields: fields_put_environment_blueprint_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEnvironmentBlueprintConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_environment_blueprint_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEnvironmentBlueprintConfiguration(ctx, input)
			},
		},
		"query-graph": {
			Name:   "query-graph",
			Fields: fields_query_graph,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.QueryGraphInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_query_graph, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.QueryGraph(ctx, input)
				}
				var results []*svc.QueryGraphOutput
				p := svc.NewQueryGraphPaginator(client, input)
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
		"reject-predictions": {
			Name:   "reject-predictions",
			Fields: fields_reject_predictions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectPredictionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_predictions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectPredictions(ctx, input)
			},
		},
		"reject-subscription-request": {
			Name:   "reject-subscription-request",
			Fields: fields_reject_subscription_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectSubscriptionRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_subscription_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectSubscriptionRequest(ctx, input)
			},
		},
		"remove-entity-owner": {
			Name:   "remove-entity-owner",
			Fields: fields_remove_entity_owner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveEntityOwnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_entity_owner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveEntityOwner(ctx, input)
			},
		},
		"remove-policy-grant": {
			Name:   "remove-policy-grant",
			Fields: fields_remove_policy_grant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemovePolicyGrantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_policy_grant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemovePolicyGrant(ctx, input)
			},
		},
		"revoke-subscription": {
			Name:   "revoke-subscription",
			Fields: fields_revoke_subscription,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeSubscriptionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_subscription, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeSubscription(ctx, input)
			},
		},
		"search": {
			Name:   "search",
			Fields: fields_search,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.Search(ctx, input)
				}
				var results []*svc.SearchOutput
				p := svc.NewSearchPaginator(client, input)
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
		"search-group-profiles": {
			Name:   "search-group-profiles",
			Fields: fields_search_group_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchGroupProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_group_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchGroupProfiles(ctx, input)
				}
				var results []*svc.SearchGroupProfilesOutput
				p := svc.NewSearchGroupProfilesPaginator(client, input)
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
		"search-listings": {
			Name:   "search-listings",
			Fields: fields_search_listings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchListingsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_listings, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchListings(ctx, input)
				}
				var results []*svc.SearchListingsOutput
				p := svc.NewSearchListingsPaginator(client, input)
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
		"search-types": {
			Name:   "search-types",
			Fields: fields_search_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchTypes(ctx, input)
				}
				var results []*svc.SearchTypesOutput
				p := svc.NewSearchTypesPaginator(client, input)
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
		"search-user-profiles": {
			Name:   "search-user-profiles",
			Fields: fields_search_user_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchUserProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_search_user_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.SearchUserProfiles(ctx, input)
				}
				var results []*svc.SearchUserProfilesOutput
				p := svc.NewSearchUserProfilesPaginator(client, input)
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
		"start-data-source-run": {
			Name:   "start-data-source-run",
			Fields: fields_start_data_source_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartDataSourceRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_data_source_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartDataSourceRun(ctx, input)
			},
		},
		"start-metadata-generation-run": {
			Name:   "start-metadata-generation-run",
			Fields: fields_start_metadata_generation_run,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMetadataGenerationRunInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_metadata_generation_run, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMetadataGenerationRun(ctx, input)
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
		"update-account-pool": {
			Name:   "update-account-pool",
			Fields: fields_update_account_pool,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAccountPoolInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_account_pool, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAccountPool(ctx, input)
			},
		},
		"update-asset-filter": {
			Name:   "update-asset-filter",
			Fields: fields_update_asset_filter,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateAssetFilterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_asset_filter, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateAssetFilter(ctx, input)
			},
		},
		"update-connection": {
			Name:   "update-connection",
			Fields: fields_update_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnection(ctx, input)
			},
		},
		"update-data-source": {
			Name:   "update-data-source",
			Fields: fields_update_data_source,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDataSourceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_data_source, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDataSource(ctx, input)
			},
		},
		"update-domain": {
			Name:   "update-domain",
			Fields: fields_update_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomain(ctx, input)
			},
		},
		"update-domain-unit": {
			Name:   "update-domain-unit",
			Fields: fields_update_domain_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainUnit(ctx, input)
			},
		},
		"update-environment": {
			Name:   "update-environment",
			Fields: fields_update_environment,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironment(ctx, input)
			},
		},
		"update-environment-action": {
			Name:   "update-environment-action",
			Fields: fields_update_environment_action,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentActionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_action, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentAction(ctx, input)
			},
		},
		"update-environment-blueprint": {
			Name:   "update-environment-blueprint",
			Fields: fields_update_environment_blueprint,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentBlueprintInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_blueprint, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentBlueprint(ctx, input)
			},
		},
		"update-environment-profile": {
			Name:   "update-environment-profile",
			Fields: fields_update_environment_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateEnvironmentProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_environment_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateEnvironmentProfile(ctx, input)
			},
		},
		"update-glossary": {
			Name:   "update-glossary",
			Fields: fields_update_glossary,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlossaryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_glossary, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlossary(ctx, input)
			},
		},
		"update-glossary-term": {
			Name:   "update-glossary-term",
			Fields: fields_update_glossary_term,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGlossaryTermInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_glossary_term, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGlossaryTerm(ctx, input)
			},
		},
		"update-group-profile": {
			Name:   "update-group-profile",
			Fields: fields_update_group_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGroupProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_group_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGroupProfile(ctx, input)
			},
		},
		"update-project": {
			Name:   "update-project",
			Fields: fields_update_project,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProject(ctx, input)
			},
		},
		"update-project-profile": {
			Name:   "update-project-profile",
			Fields: fields_update_project_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProjectProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_project_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProjectProfile(ctx, input)
			},
		},
		"update-root-domain-unit-owner": {
			Name:   "update-root-domain-unit-owner",
			Fields: fields_update_root_domain_unit_owner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRootDomainUnitOwnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_root_domain_unit_owner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRootDomainUnitOwner(ctx, input)
			},
		},
		"update-rule": {
			Name:   "update-rule",
			Fields: fields_update_rule,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRuleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_rule, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRule(ctx, input)
			},
		},
		"update-subscription-grant-status": {
			Name:   "update-subscription-grant-status",
			Fields: fields_update_subscription_grant_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionGrantStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscription_grant_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriptionGrantStatus(ctx, input)
			},
		},
		"update-subscription-request": {
			Name:   "update-subscription-request",
			Fields: fields_update_subscription_request,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionRequestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscription_request, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriptionRequest(ctx, input)
			},
		},
		"update-subscription-target": {
			Name:   "update-subscription-target",
			Fields: fields_update_subscription_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSubscriptionTargetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_subscription_target, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSubscriptionTarget(ctx, input)
			},
		},
		"update-user-profile": {
			Name:   "update-user-profile",
			Fields: fields_update_user_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateUserProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_user_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateUserProfile(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("datazone", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
