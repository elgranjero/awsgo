package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ram"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ramCmd represents the ram command
var _ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "AWS ram CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := ram.NewFromConfig(cfg)
		if _ramAcceptResourceShareInvitation {
			ram_AcceptResourceShareInvitation(cfg, client)
			return
		}
		if _ramAssociateResourceShare {
			ram_AssociateResourceShare(cfg, client)
			return
		}
		if _ramAssociateResourceSharePermission {
			ram_AssociateResourceSharePermission(cfg, client)
			return
		}
		if _ramCreatePermission {
			ram_CreatePermission(cfg, client)
			return
		}
		if _ramCreatePermissionVersion {
			ram_CreatePermissionVersion(cfg, client)
			return
		}
		if _ramCreateResourceShare {
			ram_CreateResourceShare(cfg, client)
			return
		}
		if _ramDeletePermission {
			ram_DeletePermission(cfg, client)
			return
		}
		if _ramDeletePermissionVersion {
			ram_DeletePermissionVersion(cfg, client)
			return
		}
		if _ramDeleteResourceShare {
			ram_DeleteResourceShare(cfg, client)
			return
		}
		if _ramDisassociateResourceShare {
			ram_DisassociateResourceShare(cfg, client)
			return
		}
		if _ramDisassociateResourceSharePermission {
			ram_DisassociateResourceSharePermission(cfg, client)
			return
		}
		if _ramEnableSharingWithAwsOrganization {
			ram_EnableSharingWithAwsOrganization(cfg, client)
			return
		}
		if _ramGetPermission {
			ram_GetPermission(cfg, client)
			return
		}
		if _ramGetResourcePolicies {
			ram_GetResourcePolicies(cfg, client)
			return
		}
		if _ramGetResourceShareAssociations {
			ram_GetResourceShareAssociations(cfg, client)
			return
		}
		if _ramGetResourceShareInvitations {
			ram_GetResourceShareInvitations(cfg, client)
			return
		}
		if _ramGetResourceShares {
			ram_GetResourceShares(cfg, client)
			return
		}
		if _ramListPendingInvitationResources {
			ram_ListPendingInvitationResources(cfg, client)
			return
		}
		if _ramListPermissionAssociations {
			ram_ListPermissionAssociations(cfg, client)
			return
		}
		if _ramListPermissionVersions {
			ram_ListPermissionVersions(cfg, client)
			return
		}
		if _ramListPermissions {
			ram_ListPermissions(cfg, client)
			return
		}
		if _ramListPrincipals {
			ram_ListPrincipals(cfg, client)
			return
		}
		if _ramListReplacePermissionAssociationsWork {
			ram_ListReplacePermissionAssociationsWork(cfg, client)
			return
		}
		if _ramListResourceSharePermissions {
			ram_ListResourceSharePermissions(cfg, client)
			return
		}
		if _ramListResourceTypes {
			ram_ListResourceTypes(cfg, client)
			return
		}
		if _ramListResources {
			ram_ListResources(cfg, client)
			return
		}
		if _ramListSourceAssociations {
			ram_ListSourceAssociations(cfg, client)
			return
		}
		if _ramPromotePermissionCreatedFromPolicy {
			ram_PromotePermissionCreatedFromPolicy(cfg, client)
			return
		}
		if _ramPromoteResourceShareCreatedFromPolicy {
			ram_PromoteResourceShareCreatedFromPolicy(cfg, client)
			return
		}
		if _ramRejectResourceShareInvitation {
			ram_RejectResourceShareInvitation(cfg, client)
			return
		}
		if _ramReplacePermissionAssociations {
			ram_ReplacePermissionAssociations(cfg, client)
			return
		}
		if _ramSetDefaultPermissionVersion {
			ram_SetDefaultPermissionVersion(cfg, client)
			return
		}
		if _ramTagResource {
			ram_TagResource(cfg, client)
			return
		}
		if _ramUntagResource {
			ram_UntagResource(cfg, client)
			return
		}
		if _ramUpdateResourceShare {
			ram_UpdateResourceShare(cfg, client)
			return
		}

	},
}

var (
	_ramAcceptResourceShareInvitation         bool
	_ramAssociateResourceShare                bool
	_ramAssociateResourceSharePermission      bool
	_ramCreatePermission                      bool
	_ramCreatePermissionVersion               bool
	_ramCreateResourceShare                   bool
	_ramDeletePermission                      bool
	_ramDeletePermissionVersion               bool
	_ramDeleteResourceShare                   bool
	_ramDisassociateResourceShare             bool
	_ramDisassociateResourceSharePermission   bool
	_ramEnableSharingWithAwsOrganization      bool
	_ramGetPermission                         bool
	_ramGetResourcePolicies                   bool
	_ramGetResourceShareAssociations          bool
	_ramGetResourceShareInvitations           bool
	_ramGetResourceShares                     bool
	_ramListPendingInvitationResources        bool
	_ramListPermissionAssociations            bool
	_ramListPermissionVersions                bool
	_ramListPermissions                       bool
	_ramListPrincipals                        bool
	_ramListReplacePermissionAssociationsWork bool
	_ramListResourceSharePermissions          bool
	_ramListResourceTypes                     bool
	_ramListResources                         bool
	_ramListSourceAssociations                bool
	_ramPromotePermissionCreatedFromPolicy    bool
	_ramPromoteResourceShareCreatedFromPolicy bool
	_ramRejectResourceShareInvitation         bool
	_ramReplacePermissionAssociations         bool
	_ramSetDefaultPermissionVersion           bool
	_ramTagResource                           bool
	_ramUntagResource                         bool
	_ramUpdateResourceShare                   bool

	_ramAllowExternalPrincipals     string
	_ramAssociationStatus           string
	_ramAssociationType             string
	_ramClientToken                 string
	_ramDefaultVersion              string
	_ramFeatureSet                  string
	_ramFromPermissionArn           string
	_ramFromPermissionVersion       string
	_ramMaxResults                  string
	_ramName                        string
	_ramNextToken                   string
	_ramPermissionArn               string
	_ramPermissionArns              []string
	_ramPermissionType              string
	_ramPermissionVersion           string
	_ramPolicyTemplate              string
	_ramPrincipal                   string
	_ramPrincipals                  []string
	_ramReplace                     string
	_ramResourceArn                 string
	_ramResourceArns                []string
	_ramResourceOwner               string
	_ramResourceRegionScope         string
	_ramResourceShareArn            string
	_ramResourceShareArns           []string
	_ramResourceShareConfiguration  string
	_ramResourceShareInvitationArn  string
	_ramResourceShareInvitationArns []string
	_ramResourceShareStatus         string
	_ramResourceType                string
	_ramSourceId                    string
	_ramSourceType                  string
	_ramSources                     []string
	_ramStatus                      string
	_ramTagFilters                  string
	_ramTagKeys                     []string
	_ramTags                        string
	_ramToPermissionArn             string
	_ramWorkIds                     []string
)

// Accepts an invitation to a resource share from another Amazon Web Services
// account. After you accept the invitation, the resources included in the resource
// share are available to interact with in the relevant Amazon Web Services
// Management Consoles and tools.
func ram_AcceptResourceShareInvitation(cfg aws.Config, client *ram.Client) {
	input := &ram.AcceptResourceShareInvitationInput{
		// ResourceShareInvitationArn: *string, // Required
	}

	if len(_ramResourceShareInvitationArn) > 0 {
		input.ResourceShareInvitationArn = aws.String(_ramResourceShareInvitationArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.AcceptResourceShareInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified list of principals, resources, and source constraints to a
// resource share. Principals that already have access to this resource share
// immediately receive access to the added resources. Newly added principals
// immediately receive access to the resources shared in this resource share.
func ram_AssociateResourceShare(cfg aws.Config, client *ram.Client) {
	input := &ram.AssociateResourceShareInput{
		// ResourceShareArn: *string, // Required
	}

	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramPrincipals) > 0 {
		input.Principals = append([]string(nil), _ramPrincipals...)
	}
	if len(_ramResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _ramResourceArns...)
	}
	if len(_ramSources) > 0 {
		input.Sources = append([]string(nil), _ramSources...)
	}

	if resp, err := client.AssociateResourceShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or replaces the RAM permission for a resource type included in a resource
// share. You can have exactly one permission associated with each resource type in
// the resource share. You can add a new RAM permission only if there are currently
// no resources of that resource type currently in the resource share.
func ram_AssociateResourceSharePermission(cfg aws.Config, client *ram.Client) {
	input := &ram.AssociateResourceSharePermissionInput{
		// PermissionArn: *string, // Required
		// ResourceShareArn: *string, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramPermissionVersion) > 0 {
		if err := assignInputField(input, "PermissionVersion", _ramPermissionVersion); err != nil {
			log.Errorf("invalid --permission-version: %s", err.Error())
			return
		}
	}
	if len(_ramReplace) > 0 {
		if err := assignInputField(input, "Replace", _ramReplace); err != nil {
			log.Errorf("invalid --replace: %s", err.Error())
			return
		}
	}

	if resp, err := client.AssociateResourceSharePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a customer managed permission for a specified resource type that you
// can attach to resource shares. It is created in the Amazon Web Services Region
// in which you call the operation.
func ram_CreatePermission(cfg aws.Config, client *ram.Client) {
	input := &ram.CreatePermissionInput{
		// Name: *string, // Required
		// PolicyTemplate: *string, // Required
		// ResourceType: *string, // Required
	}

	if len(_ramName) > 0 {
		input.Name = aws.String(_ramName)
	}
	if len(_ramPolicyTemplate) > 0 {
		input.PolicyTemplate = aws.String(_ramPolicyTemplate)
	}
	if len(_ramResourceType) > 0 {
		input.ResourceType = aws.String(_ramResourceType)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramTags) > 0 {
		if err := assignInputField(input, "Tags", _ramTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new version of the specified customer managed permission. The new
// version is automatically set as the default version of the customer managed
// permission. New resource shares automatically use the default permission.
// Existing resource shares continue to use their original permission versions, but
// you can use ReplacePermissionAssociationsto update them.
//
// If the specified customer managed permission already has the maximum of 5
// versions, then you must delete one of the existing versions before you can
// create a new one.
func ram_CreatePermissionVersion(cfg aws.Config, client *ram.Client) {
	input := &ram.CreatePermissionVersionInput{
		// PermissionArn: *string, // Required
		// PolicyTemplate: *string, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramPolicyTemplate) > 0 {
		input.PolicyTemplate = aws.String(_ramPolicyTemplate)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.CreatePermissionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a resource share. You can provide a list of the [Amazon Resource Names (ARNs)] for the resources that
// you want to share, a list of principals you want to share the resources with,
// the permissions to grant those principals, and optionally source constraints to
// enhance security for service principal sharing.
//
// Sharing a resource makes it available for use by principals outside of the
// Amazon Web Services account that created the resource. Sharing doesn't change
// any permissions or quotas that apply to the resource in the account that created
// it.
//
// [Amazon Resource Names (ARNs)]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
func ram_CreateResourceShare(cfg aws.Config, client *ram.Client) {
	input := &ram.CreateResourceShareInput{
		// Name: *string, // Required
	}

	if len(_ramName) > 0 {
		input.Name = aws.String(_ramName)
	}
	if len(_ramAllowExternalPrincipals) > 0 {
		if err := assignInputField(input, "AllowExternalPrincipals", _ramAllowExternalPrincipals); err != nil {
			log.Errorf("invalid --allow-external-principals: %s", err.Error())
			return
		}
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramPermissionArns) > 0 {
		input.PermissionArns = append([]string(nil), _ramPermissionArns...)
	}
	if len(_ramPrincipals) > 0 {
		input.Principals = append([]string(nil), _ramPrincipals...)
	}
	if len(_ramResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _ramResourceArns...)
	}
	if len(_ramResourceShareConfiguration) > 0 {
		if err := assignInputField(input, "ResourceShareConfiguration", _ramResourceShareConfiguration); err != nil {
			log.Errorf("invalid --resource-share-configuration: %s", err.Error())
			return
		}
	}
	if len(_ramSources) > 0 {
		input.Sources = append([]string(nil), _ramSources...)
	}
	if len(_ramTags) > 0 {
		if err := assignInputField(input, "Tags", _ramTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateResourceShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified customer managed permission in the Amazon Web Services
// Region in which you call this operation. You can delete a customer managed
// permission only if it isn't attached to any resource share. The operation
// deletes all versions associated with the customer managed permission.
func ram_DeletePermission(cfg aws.Config, client *ram.Client) {
	input := &ram.DeletePermissionInput{
		// PermissionArn: *string, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.DeletePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one version of a customer managed permission. The version you specify
// must not be attached to any resource share and must not be the default version
// for the permission.
//
// If a customer managed permission has the maximum of 5 versions, then you must
// delete at least one version before you can create another.
func ram_DeletePermissionVersion(cfg aws.Config, client *ram.Client) {
	input := &ram.DeletePermissionVersionInput{
		// PermissionArn: *string, // Required
		// PermissionVersion: *int32, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramPermissionVersion) > 0 {
		if err := assignInputField(input, "PermissionVersion", _ramPermissionVersion); err != nil {
			log.Errorf("invalid --permission-version: %s", err.Error())
			return
		}
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.DeletePermissionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified resource share.
// This doesn't delete any of the resources that were associated with the resource
// share; it only stops the sharing of those resources through this resource share.
func ram_DeleteResourceShare(cfg aws.Config, client *ram.Client) {
	input := &ram.DeleteResourceShareInput{
		// ResourceShareArn: *string, // Required
	}

	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.DeleteResourceShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified principals, resources, or source constraints from
// participating in the specified resource share.
func ram_DisassociateResourceShare(cfg aws.Config, client *ram.Client) {
	input := &ram.DisassociateResourceShareInput{
		// ResourceShareArn: *string, // Required
	}

	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramPrincipals) > 0 {
		input.Principals = append([]string(nil), _ramPrincipals...)
	}
	if len(_ramResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _ramResourceArns...)
	}
	if len(_ramSources) > 0 {
		input.Sources = append([]string(nil), _ramSources...)
	}

	if resp, err := client.DisassociateResourceShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes a managed permission from a resource share. Permission changes take
// effect immediately. You can remove a managed permission from a resource share
// only if there are currently no resources of the relevant resource type currently
// attached to the resource share.
func ram_DisassociateResourceSharePermission(cfg aws.Config, client *ram.Client) {
	input := &ram.DisassociateResourceSharePermissionInput{
		// PermissionArn: *string, // Required
		// ResourceShareArn: *string, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.DisassociateResourceSharePermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables resource sharing within your organization in Organizations. This
// operation creates a service-linked role called
// AWSServiceRoleForResourceAccessManager that has the IAM managed policy named
// AWSResourceAccessManagerServiceRolePolicy attached. This role permits RAM to
// retrieve information about the organization and its structure. This lets you
// share resources with all of the accounts in the calling account's organization
// by specifying the organization ID, or all of the accounts in an organizational
// unit (OU) by specifying the OU ID. Until you enable sharing within the
// organization, you can specify only individual Amazon Web Services accounts, or
// for supported resource types, IAM roles and users.
//
// You must call this operation from an IAM role or user in the organization's
// management account.
func ram_EnableSharingWithAwsOrganization(cfg aws.Config, client *ram.Client) {
	input := &ram.EnableSharingWithAwsOrganizationInput{}

	if resp, err := client.EnableSharingWithAwsOrganization(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the contents of a managed permission in JSON format.
func ram_GetPermission(cfg aws.Config, client *ram.Client) {
	input := &ram.GetPermissionInput{
		// PermissionArn: *string, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramPermissionVersion) > 0 {
		if err := assignInputField(input, "PermissionVersion", _ramPermissionVersion); err != nil {
			log.Errorf("invalid --permission-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetPermission(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the resource policies for the specified resources that you own and
// have shared.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_GetResourcePolicies(cfg aws.Config, client *ram.Client) {
	input := &ram.GetResourcePoliciesInput{
		// ResourceArns: []string, // Required
	}

	if len(_ramResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _ramResourceArns...)
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPrincipal) > 0 {
		input.Principal = aws.String(_ramPrincipal)
	}

	if disablePaginator() {
		if resp, err := client.GetResourcePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.GetResourcePoliciesOutput
	p := ram.NewGetResourcePoliciesPaginator(client, input)
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

// Retrieves the lists of resources and principals that associated for resource
// shares that you own.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_GetResourceShareAssociations(cfg aws.Config, client *ram.Client) {
	input := &ram.GetResourceShareAssociationsInput{
		// AssociationType: types.ResourceShareAssociationType, // Required
	}

	if len(_ramAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _ramAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}
	if len(_ramAssociationStatus) > 0 {
		if err := assignInputField(input, "AssociationStatus", _ramAssociationStatus); err != nil {
			log.Errorf("invalid --association-status: %s", err.Error())
			return
		}
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPrincipal) > 0 {
		input.Principal = aws.String(_ramPrincipal)
	}
	if len(_ramResourceArn) > 0 {
		input.ResourceArn = aws.String(_ramResourceArn)
	}
	if len(_ramResourceShareArns) > 0 {
		input.ResourceShareArns = append([]string(nil), _ramResourceShareArns...)
	}

	if disablePaginator() {
		if resp, err := client.GetResourceShareAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.GetResourceShareAssociationsOutput
	p := ram.NewGetResourceShareAssociationsPaginator(client, input)
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

// Retrieves details about invitations that you have received for resource shares.
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_GetResourceShareInvitations(cfg aws.Config, client *ram.Client) {
	input := &ram.GetResourceShareInvitationsInput{}

	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramResourceShareArns) > 0 {
		input.ResourceShareArns = append([]string(nil), _ramResourceShareArns...)
	}
	if len(_ramResourceShareInvitationArns) > 0 {
		input.ResourceShareInvitationArns = append([]string(nil), _ramResourceShareInvitationArns...)
	}

	if disablePaginator() {
		if resp, err := client.GetResourceShareInvitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.GetResourceShareInvitationsOutput
	p := ram.NewGetResourceShareInvitationsPaginator(client, input)
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

// Retrieves details about the resource shares that you own or that are shared
// with you.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_GetResourceShares(cfg aws.Config, client *ram.Client) {
	input := &ram.GetResourceSharesInput{
		// ResourceOwner: types.ResourceOwner, // Required
	}

	if len(_ramResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _ramResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramName) > 0 {
		input.Name = aws.String(_ramName)
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramPermissionVersion) > 0 {
		if err := assignInputField(input, "PermissionVersion", _ramPermissionVersion); err != nil {
			log.Errorf("invalid --permission-version: %s", err.Error())
			return
		}
	}
	if len(_ramResourceShareArns) > 0 {
		input.ResourceShareArns = append([]string(nil), _ramResourceShareArns...)
	}
	if len(_ramResourceShareStatus) > 0 {
		if err := assignInputField(input, "ResourceShareStatus", _ramResourceShareStatus); err != nil {
			log.Errorf("invalid --resource-share-status: %s", err.Error())
			return
		}
	}
	if len(_ramTagFilters) > 0 {
		if err := assignInputField(input, "TagFilters", _ramTagFilters); err != nil {
			log.Errorf("invalid --tag-filters: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetResourceShares(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.GetResourceSharesOutput
	p := ram.NewGetResourceSharesPaginator(client, input)
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

// Lists the resources in a resource share that is shared with you but for which
// the invitation is still PENDING . That means that you haven't accepted or
// rejected the invitation and the invitation hasn't expired.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListPendingInvitationResources(cfg aws.Config, client *ram.Client) {
	input := &ram.ListPendingInvitationResourcesInput{
		// ResourceShareInvitationArn: *string, // Required
	}

	if len(_ramResourceShareInvitationArn) > 0 {
		input.ResourceShareInvitationArn = aws.String(_ramResourceShareInvitationArn)
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramResourceRegionScope) > 0 {
		if err := assignInputField(input, "ResourceRegionScope", _ramResourceRegionScope); err != nil {
			log.Errorf("invalid --resource-region-scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListPendingInvitationResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListPendingInvitationResourcesOutput
	p := ram.NewListPendingInvitationResourcesPaginator(client, input)
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

// Lists information about the managed permission and its associations to any
// resource shares that use this managed permission. This lets you see which
// resource shares use which versions of the specified managed permission.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListPermissionAssociations(cfg aws.Config, client *ram.Client) {
	input := &ram.ListPermissionAssociationsInput{}

	if len(_ramAssociationStatus) > 0 {
		if err := assignInputField(input, "AssociationStatus", _ramAssociationStatus); err != nil {
			log.Errorf("invalid --association-status: %s", err.Error())
			return
		}
	}
	if len(_ramDefaultVersion) > 0 {
		if err := assignInputField(input, "DefaultVersion", _ramDefaultVersion); err != nil {
			log.Errorf("invalid --default-version: %s", err.Error())
			return
		}
	}
	if len(_ramFeatureSet) > 0 {
		if err := assignInputField(input, "FeatureSet", _ramFeatureSet); err != nil {
			log.Errorf("invalid --feature-set: %s", err.Error())
			return
		}
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramPermissionVersion) > 0 {
		if err := assignInputField(input, "PermissionVersion", _ramPermissionVersion); err != nil {
			log.Errorf("invalid --permission-version: %s", err.Error())
			return
		}
	}
	if len(_ramResourceType) > 0 {
		input.ResourceType = aws.String(_ramResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissionAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListPermissionAssociationsOutput
	p := ram.NewListPermissionAssociationsPaginator(client, input)
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

// Lists the available versions of the specified RAM permission.
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListPermissionVersions(cfg aws.Config, client *ram.Client) {
	input := &ram.ListPermissionVersionsInput{
		// PermissionArn: *string, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissionVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListPermissionVersionsOutput
	p := ram.NewListPermissionVersionsPaginator(client, input)
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

// Retrieves a list of available RAM permissions that you can use for the
// supported resource types.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListPermissions(cfg aws.Config, client *ram.Client) {
	input := &ram.ListPermissionsInput{}

	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPermissionType) > 0 {
		if err := assignInputField(input, "PermissionType", _ramPermissionType); err != nil {
			log.Errorf("invalid --permission-type: %s", err.Error())
			return
		}
	}
	if len(_ramResourceType) > 0 {
		input.ResourceType = aws.String(_ramResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListPermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListPermissionsOutput
	p := ram.NewListPermissionsPaginator(client, input)
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

// Lists the principals that you are sharing resources with or that are sharing
// resources with you.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListPrincipals(cfg aws.Config, client *ram.Client) {
	input := &ram.ListPrincipalsInput{
		// ResourceOwner: types.ResourceOwner, // Required
	}

	if len(_ramResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _ramResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPrincipals) > 0 {
		input.Principals = append([]string(nil), _ramPrincipals...)
	}
	if len(_ramResourceArn) > 0 {
		input.ResourceArn = aws.String(_ramResourceArn)
	}
	if len(_ramResourceShareArns) > 0 {
		input.ResourceShareArns = append([]string(nil), _ramResourceShareArns...)
	}
	if len(_ramResourceType) > 0 {
		input.ResourceType = aws.String(_ramResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListPrincipals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListPrincipalsOutput
	p := ram.NewListPrincipalsPaginator(client, input)
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

// Retrieves the current status of the asynchronous tasks performed by RAM when
// you perform the ReplacePermissionAssociationsWorkoperation.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListReplacePermissionAssociationsWork(cfg aws.Config, client *ram.Client) {
	input := &ram.ListReplacePermissionAssociationsWorkInput{}

	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramStatus) > 0 {
		if err := assignInputField(input, "Status", _ramStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}
	if len(_ramWorkIds) > 0 {
		input.WorkIds = append([]string(nil), _ramWorkIds...)
	}

	if disablePaginator() {
		if resp, err := client.ListReplacePermissionAssociationsWork(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListReplacePermissionAssociationsWorkOutput
	p := ram.NewListReplacePermissionAssociationsWorkPaginator(client, input)
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

// Lists the RAM permissions that are associated with a resource share.
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListResourceSharePermissions(cfg aws.Config, client *ram.Client) {
	input := &ram.ListResourceSharePermissionsInput{
		// ResourceShareArn: *string, // Required
	}

	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourceSharePermissions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListResourceSharePermissionsOutput
	p := ram.NewListResourceSharePermissionsPaginator(client, input)
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

// Lists the resource types that can be shared by RAM.
func ram_ListResourceTypes(cfg aws.Config, client *ram.Client) {
	input := &ram.ListResourceTypesInput{}

	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramResourceRegionScope) > 0 {
		if err := assignInputField(input, "ResourceRegionScope", _ramResourceRegionScope); err != nil {
			log.Errorf("invalid --resource-region-scope: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListResourceTypes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListResourceTypesOutput
	p := ram.NewListResourceTypesPaginator(client, input)
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

// Lists the resources that you added to a resource share or the resources that
// are shared with you.
//
// Always check the NextToken response parameter for a null value when calling a
// paginated operation. These operations can occasionally return an empty set of
// results even when there are more results available. The NextToken response
// parameter value is null only when there are no more results to display.
func ram_ListResources(cfg aws.Config, client *ram.Client) {
	input := &ram.ListResourcesInput{
		// ResourceOwner: types.ResourceOwner, // Required
	}

	if len(_ramResourceOwner) > 0 {
		if err := assignInputField(input, "ResourceOwner", _ramResourceOwner); err != nil {
			log.Errorf("invalid --resource-owner: %s", err.Error())
			return
		}
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramPrincipal) > 0 {
		input.Principal = aws.String(_ramPrincipal)
	}
	if len(_ramResourceArns) > 0 {
		input.ResourceArns = append([]string(nil), _ramResourceArns...)
	}
	if len(_ramResourceRegionScope) > 0 {
		if err := assignInputField(input, "ResourceRegionScope", _ramResourceRegionScope); err != nil {
			log.Errorf("invalid --resource-region-scope: %s", err.Error())
			return
		}
	}
	if len(_ramResourceShareArns) > 0 {
		input.ResourceShareArns = append([]string(nil), _ramResourceShareArns...)
	}
	if len(_ramResourceType) > 0 {
		input.ResourceType = aws.String(_ramResourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListResourcesOutput
	p := ram.NewListResourcesPaginator(client, input)
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

// Lists source associations for resource shares. Source associations control
// which sources can be used with service principals in resource shares. This
// operation provides visibility into source associations for resource share
// owners.
//
// You can filter the results by resource share Amazon Resource Name (ARN), source
// ID, source type, or association status. We recommend using pagination to ensure
// that the operation returns quickly and successfully.
func ram_ListSourceAssociations(cfg aws.Config, client *ram.Client) {
	input := &ram.ListSourceAssociationsInput{}

	if len(_ramAssociationStatus) > 0 {
		if err := assignInputField(input, "AssociationStatus", _ramAssociationStatus); err != nil {
			log.Errorf("invalid --association-status: %s", err.Error())
			return
		}
	}
	if len(_ramMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _ramMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_ramNextToken) > 0 {
		input.NextToken = aws.String(_ramNextToken)
	}
	if len(_ramResourceShareArns) > 0 {
		input.ResourceShareArns = append([]string(nil), _ramResourceShareArns...)
	}
	if len(_ramSourceId) > 0 {
		input.SourceId = aws.String(_ramSourceId)
	}
	if len(_ramSourceType) > 0 {
		input.SourceType = aws.String(_ramSourceType)
	}

	if disablePaginator() {
		if resp, err := client.ListSourceAssociations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*ram.ListSourceAssociationsOutput
	p := ram.NewListSourceAssociationsPaginator(client, input)
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

// When you attach a resource-based policy to a resource, RAM automatically
// creates a resource share of featureSet = CREATED_FROM_POLICY with a managed
// permission that has the same IAM permissions as the original resource-based
// policy. However, this type of managed permission is visible to only the resource
// share owner, and the associated resource share can't be modified by using RAM.
//
// This operation creates a separate, fully manageable customer managed permission
// that has the same IAM permissions as the original resource-based policy. You can
// associate this customer managed permission to any resource shares.
//
// Before you use PromoteResourceShareCreatedFromPolicy, you should first run this operation to ensure that you have an
// appropriate customer managed permission that can be associated with the promoted
// resource share.
//
// - The original CREATED_FROM_POLICY policy isn't deleted, and resource shares
// using that original policy aren't automatically updated.
//
// - You can't modify a CREATED_FROM_POLICY resource share so you can't associate
// the new customer managed permission by using ReplacePermsissionAssociations .
// However, if you use PromoteResourceShareCreatedFromPolicy, that operation automatically associates the fully
// manageable customer managed permission to the newly promoted STANDARD resource
// share.
//
// - After you promote a resource share, if the original CREATED_FROM_POLICY
// managed permission has no other associations to A resource share, then RAM
// automatically deletes it.
func ram_PromotePermissionCreatedFromPolicy(cfg aws.Config, client *ram.Client) {
	input := &ram.PromotePermissionCreatedFromPolicyInput{
		// Name: *string, // Required
		// PermissionArn: *string, // Required
	}

	if len(_ramName) > 0 {
		input.Name = aws.String(_ramName)
	}
	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.PromotePermissionCreatedFromPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// When you attach a resource-based policy to a resource, RAM automatically
// creates a resource share of featureSet = CREATED_FROM_POLICY with a managed
// permission that has the same IAM permissions as the original resource-based
// policy. However, this type of managed permission is visible to only the resource
// share owner, and the associated resource share can't be modified by using RAM.
//
// This operation promotes the resource share to a STANDARD resource share that is
// fully manageable in RAM. When you promote a resource share, you can then manage
// the resource share in RAM and it becomes visible to all of the principals you
// shared it with.
//
// Before you perform this operation, you should first run PromotePermissionCreatedFromPolicyto ensure that you have
// an appropriate customer managed permission that can be associated with this
// resource share after its is promoted. If this operation can't find a managed
// permission that exactly matches the existing CREATED_FROM_POLICY permission,
// then this operation fails.
func ram_PromoteResourceShareCreatedFromPolicy(cfg aws.Config, client *ram.Client) {
	input := &ram.PromoteResourceShareCreatedFromPolicyInput{
		// ResourceShareArn: *string, // Required
	}

	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}

	if resp, err := client.PromoteResourceShareCreatedFromPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects an invitation to a resource share from another Amazon Web Services
// account.
func ram_RejectResourceShareInvitation(cfg aws.Config, client *ram.Client) {
	input := &ram.RejectResourceShareInvitationInput{
		// ResourceShareInvitationArn: *string, // Required
	}

	if len(_ramResourceShareInvitationArn) > 0 {
		input.ResourceShareInvitationArn = aws.String(_ramResourceShareInvitationArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.RejectResourceShareInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates all resource shares that use a managed permission to a different
// managed permission. This operation always applies the default version of the
// target managed permission. You can optionally specify that the update applies to
// only resource shares that currently use a specified version. This enables you to
// update to the latest version, without changing the which managed permission is
// used.
//
// You can use this operation to update all of your resource shares to use the
// current default version of the permission by specifying the same value for the
// fromPermissionArn and toPermissionArn parameters.
//
// You can use the optional fromPermissionVersion parameter to update only those
// resources that use a specified version of the managed permission to the new
// managed permission.
//
// To successfully perform this operation, you must have permission to update the
// resource-based policy on all affected resource types.
func ram_ReplacePermissionAssociations(cfg aws.Config, client *ram.Client) {
	input := &ram.ReplacePermissionAssociationsInput{
		// FromPermissionArn: *string, // Required
		// ToPermissionArn: *string, // Required
	}

	if len(_ramFromPermissionArn) > 0 {
		input.FromPermissionArn = aws.String(_ramFromPermissionArn)
	}
	if len(_ramToPermissionArn) > 0 {
		input.ToPermissionArn = aws.String(_ramToPermissionArn)
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramFromPermissionVersion) > 0 {
		if err := assignInputField(input, "FromPermissionVersion", _ramFromPermissionVersion); err != nil {
			log.Errorf("invalid --from-permission-version: %s", err.Error())
			return
		}
	}

	if resp, err := client.ReplacePermissionAssociations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates the specified version number as the default version for the
// specified customer managed permission. New resource shares automatically use
// this new default permission. Existing resource shares continue to use their
// original permission version, but you can use ReplacePermissionAssociationsto update them.
func ram_SetDefaultPermissionVersion(cfg aws.Config, client *ram.Client) {
	input := &ram.SetDefaultPermissionVersionInput{
		// PermissionArn: *string, // Required
		// PermissionVersion: *int32, // Required
	}

	if len(_ramPermissionArn) > 0 {
		input.PermissionArn = aws.String(_ramPermissionArn)
	}
	if len(_ramPermissionVersion) > 0 {
		if err := assignInputField(input, "PermissionVersion", _ramPermissionVersion); err != nil {
			log.Errorf("invalid --permission-version: %s", err.Error())
			return
		}
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}

	if resp, err := client.SetDefaultPermissionVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds the specified tag keys and values to a resource share or managed
// permission. If you choose a resource share, the tags are attached to only the
// resource share, not to the resources that are in the resource share.
//
// The tags on a managed permission are the same for all versions of the managed
// permission.
func ram_TagResource(cfg aws.Config, client *ram.Client) {
	input := &ram.TagResourceInput{
		// Tags: []types.Tag, // Required
	}

	if len(_ramTags) > 0 {
		if err := assignInputField(input, "Tags", _ramTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_ramResourceArn) > 0 {
		input.ResourceArn = aws.String(_ramResourceArn)
	}
	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tag key and value pairs from the specified resource share
// or managed permission.
func ram_UntagResource(cfg aws.Config, client *ram.Client) {
	input := &ram.UntagResourceInput{
		// TagKeys: []string, // Required
	}

	if len(_ramTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _ramTagKeys...)
	}
	if len(_ramResourceArn) > 0 {
		input.ResourceArn = aws.String(_ramResourceArn)
	}
	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Modifies some of the properties of the specified resource share.
func ram_UpdateResourceShare(cfg aws.Config, client *ram.Client) {
	input := &ram.UpdateResourceShareInput{
		// ResourceShareArn: *string, // Required
	}

	if len(_ramResourceShareArn) > 0 {
		input.ResourceShareArn = aws.String(_ramResourceShareArn)
	}
	if len(_ramAllowExternalPrincipals) > 0 {
		if err := assignInputField(input, "AllowExternalPrincipals", _ramAllowExternalPrincipals); err != nil {
			log.Errorf("invalid --allow-external-principals: %s", err.Error())
			return
		}
	}
	if len(_ramClientToken) > 0 {
		input.ClientToken = aws.String(_ramClientToken)
	}
	if len(_ramName) > 0 {
		input.Name = aws.String(_ramName)
	}

	if resp, err := client.UpdateResourceShare(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_ramCmd)
	_ramCmd.Flags().SortFlags = false

	_ramCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_ramCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_ramCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_ramCmd.Flags().StringVarP(&_ramAllowExternalPrincipals, "allow-external-principals", "", "", "Allow External Principals")
	_ramCmd.Flags().StringVarP(&_ramAssociationStatus, "association-status", "", "", "Association Status")
	_ramCmd.Flags().StringVarP(&_ramAssociationType, "association-type", "", "", "Association Type")
	_ramCmd.Flags().StringVarP(&_ramClientToken, "client-token", "", "", "Client Token")
	_ramCmd.Flags().StringVarP(&_ramDefaultVersion, "default-version", "", "", "Default Version")
	_ramCmd.Flags().StringVarP(&_ramFeatureSet, "feature-set", "", "", "Feature Set")
	_ramCmd.Flags().StringVarP(&_ramFromPermissionArn, "from-permission-arn", "", "", "From Permission ARN")
	_ramCmd.Flags().StringVarP(&_ramFromPermissionVersion, "from-permission-version", "", "", "From Permission Version")
	_ramCmd.Flags().StringVarP(&_ramMaxResults, "max-results", "", "", "Max Results")
	_ramCmd.Flags().StringVarP(&_ramName, "name", "", "", "Name")
	_ramCmd.Flags().StringVarP(&_ramNextToken, "next-token", "", "", "Next Token")
	_ramCmd.Flags().StringVarP(&_ramPermissionArn, "permission-arn", "", "", "Permission ARN")
	_ramCmd.Flags().StringSliceVarP(&_ramPermissionArns, "permission-arns", "", nil, "Permission Arns")
	_ramCmd.Flags().StringVarP(&_ramPermissionType, "permission-type", "", "", "Permission Type")
	_ramCmd.Flags().StringVarP(&_ramPermissionVersion, "permission-version", "", "", "Permission Version")
	_ramCmd.Flags().StringVarP(&_ramPolicyTemplate, "policy-template", "", "", "Policy Template")
	_ramCmd.Flags().StringVarP(&_ramPrincipal, "principal", "", "", "Principal")
	_ramCmd.Flags().StringSliceVarP(&_ramPrincipals, "principals", "", nil, "Principals")
	_ramCmd.Flags().StringVarP(&_ramReplace, "replace", "", "", "Replace")
	_ramCmd.Flags().StringVarP(&_ramResourceArn, "resource-arn", "", "", "Resource ARN")
	_ramCmd.Flags().StringSliceVarP(&_ramResourceArns, "resource-arns", "", nil, "Resource Arns")
	_ramCmd.Flags().StringVarP(&_ramResourceOwner, "resource-owner", "", "", "Resource Owner")
	_ramCmd.Flags().StringVarP(&_ramResourceRegionScope, "resource-region-scope", "", "", "Resource Region Scope")
	_ramCmd.Flags().StringVarP(&_ramResourceShareArn, "resource-share-arn", "", "", "Resource Share ARN")
	_ramCmd.Flags().StringSliceVarP(&_ramResourceShareArns, "resource-share-arns", "", nil, "Resource Share Arns")
	_ramCmd.Flags().StringVarP(&_ramResourceShareConfiguration, "resource-share-configuration", "", "", "Resource Share Configuration")
	_ramCmd.Flags().StringVarP(&_ramResourceShareInvitationArn, "resource-share-invitation-arn", "", "", "Resource Share Invitation ARN")
	_ramCmd.Flags().StringSliceVarP(&_ramResourceShareInvitationArns, "resource-share-invitation-arns", "", nil, "Resource Share Invitation Arns")
	_ramCmd.Flags().StringVarP(&_ramResourceShareStatus, "resource-share-status", "", "", "Resource Share Status")
	_ramCmd.Flags().StringVarP(&_ramResourceType, "resource-type", "", "", "Resource Type")
	_ramCmd.Flags().StringVarP(&_ramSourceId, "source-id", "", "", "Source ID")
	_ramCmd.Flags().StringVarP(&_ramSourceType, "source-type", "", "", "Source Type")
	_ramCmd.Flags().StringSliceVarP(&_ramSources, "sources", "", nil, "Sources")
	_ramCmd.Flags().StringVarP(&_ramStatus, "status", "", "", "Status")
	_ramCmd.Flags().StringVarP(&_ramTagFilters, "tag-filters", "", "", "Tag Filters")
	_ramCmd.Flags().StringSliceVarP(&_ramTagKeys, "tag-keys", "", nil, "Tag Keys")
	_ramCmd.Flags().StringVarP(&_ramTags, "tags", "", "", "Tags")
	_ramCmd.Flags().StringVarP(&_ramToPermissionArn, "to-permission-arn", "", "", "To Permission ARN")
	_ramCmd.Flags().StringSliceVarP(&_ramWorkIds, "work-ids", "", nil, "Work Ids")

	_ramCmd.Flags().BoolVarP(&_ramAcceptResourceShareInvitation, "accept-resource-share-invitation", "", false, "Accept Resource Share Invitation")
	_ramCmd.Flags().BoolVarP(&_ramAssociateResourceShare, "associate-resource-share", "", false, "Associate Resource Share")
	_ramCmd.Flags().BoolVarP(&_ramAssociateResourceSharePermission, "associate-resource-share-permission", "", false, "Associate Resource Share Permission")
	_ramCmd.Flags().BoolVarP(&_ramCreatePermission, "create-permission", "", false, "Create Permission")
	_ramCmd.Flags().BoolVarP(&_ramCreatePermissionVersion, "create-permission-version", "", false, "Create Permission Version")
	_ramCmd.Flags().BoolVarP(&_ramCreateResourceShare, "create-resource-share", "", false, "Create Resource Share")
	_ramCmd.Flags().BoolVarP(&_ramDeletePermission, "delete-permission", "", false, "Delete Permission")
	_ramCmd.Flags().BoolVarP(&_ramDeletePermissionVersion, "delete-permission-version", "", false, "Delete Permission Version")
	_ramCmd.Flags().BoolVarP(&_ramDeleteResourceShare, "delete-resource-share", "", false, "Delete Resource Share")
	_ramCmd.Flags().BoolVarP(&_ramDisassociateResourceShare, "disassociate-resource-share", "", false, "Disassociate Resource Share")
	_ramCmd.Flags().BoolVarP(&_ramDisassociateResourceSharePermission, "disassociate-resource-share-permission", "", false, "Disassociate Resource Share Permission")
	_ramCmd.Flags().BoolVarP(&_ramEnableSharingWithAwsOrganization, "enable-sharing-with-aws-organization", "", false, "Enable Sharing With AWS Organization")
	_ramCmd.Flags().BoolVarP(&_ramGetPermission, "get-permission", "", false, "Get Permission")
	_ramCmd.Flags().BoolVarP(&_ramGetResourcePolicies, "get-resource-policies", "", false, "Get Resource Policies")
	_ramCmd.Flags().BoolVarP(&_ramGetResourceShareAssociations, "get-resource-share-associations", "", false, "Get Resource Share Associations")
	_ramCmd.Flags().BoolVarP(&_ramGetResourceShareInvitations, "get-resource-share-invitations", "", false, "Get Resource Share Invitations")
	_ramCmd.Flags().BoolVarP(&_ramGetResourceShares, "get-resource-shares", "", false, "Get Resource Shares")
	_ramCmd.Flags().BoolVarP(&_ramListPendingInvitationResources, "list-pending-invitation-resources", "", false, "List Pending Invitation Resources")
	_ramCmd.Flags().BoolVarP(&_ramListPermissionAssociations, "list-permission-associations", "", false, "List Permission Associations")
	_ramCmd.Flags().BoolVarP(&_ramListPermissionVersions, "list-permission-versions", "", false, "List Permission Versions")
	_ramCmd.Flags().BoolVarP(&_ramListPermissions, "list-permissions", "", false, "List Permissions")
	_ramCmd.Flags().BoolVarP(&_ramListPrincipals, "list-principals", "", false, "List Principals")
	_ramCmd.Flags().BoolVarP(&_ramListReplacePermissionAssociationsWork, "list-replace-permission-associations-work", "", false, "List Replace Permission Associations Work")
	_ramCmd.Flags().BoolVarP(&_ramListResourceSharePermissions, "list-resource-share-permissions", "", false, "List Resource Share Permissions")
	_ramCmd.Flags().BoolVarP(&_ramListResourceTypes, "list-resource-types", "", false, "List Resource Types")
	_ramCmd.Flags().BoolVarP(&_ramListResources, "list-resources", "", false, "List Resources")
	_ramCmd.Flags().BoolVarP(&_ramListSourceAssociations, "list-source-associations", "", false, "List Source Associations")
	_ramCmd.Flags().BoolVarP(&_ramPromotePermissionCreatedFromPolicy, "promote-permission-created-from-policy", "", false, "Promote Permission Created From Policy")
	_ramCmd.Flags().BoolVarP(&_ramPromoteResourceShareCreatedFromPolicy, "promote-resource-share-created-from-policy", "", false, "Promote Resource Share Created From Policy")
	_ramCmd.Flags().BoolVarP(&_ramRejectResourceShareInvitation, "reject-resource-share-invitation", "", false, "Reject Resource Share Invitation")
	_ramCmd.Flags().BoolVarP(&_ramReplacePermissionAssociations, "replace-permission-associations", "", false, "Replace Permission Associations")
	_ramCmd.Flags().BoolVarP(&_ramSetDefaultPermissionVersion, "set-default-permission-version", "", false, "Set Default Permission Version")
	_ramCmd.Flags().BoolVarP(&_ramTagResource, "tag-resource", "", false, "Tag Resource")
	_ramCmd.Flags().BoolVarP(&_ramUntagResource, "untag-resource", "", false, "Untag Resource")
	_ramCmd.Flags().BoolVarP(&_ramUpdateResourceShare, "update-resource-share", "", false, "Update Resource Share")

}
