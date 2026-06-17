package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/identitystore"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// identitystoreCmd represents the identitystore command
var _identitystoreCmd = &cobra.Command{
	Use:   "identitystore",
	Short: "AWS identitystore CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := identitystore.NewFromConfig(cfg)
		if _identitystoreCreateGroup {
			identitystore_CreateGroup(cfg, client)
			return
		}
		if _identitystoreCreateGroupMembership {
			identitystore_CreateGroupMembership(cfg, client)
			return
		}
		if _identitystoreCreateUser {
			identitystore_CreateUser(cfg, client)
			return
		}
		if _identitystoreDeleteGroup {
			identitystore_DeleteGroup(cfg, client)
			return
		}
		if _identitystoreDeleteGroupMembership {
			identitystore_DeleteGroupMembership(cfg, client)
			return
		}
		if _identitystoreDeleteUser {
			identitystore_DeleteUser(cfg, client)
			return
		}
		if _identitystoreDescribeGroup {
			identitystore_DescribeGroup(cfg, client)
			return
		}
		if _identitystoreDescribeGroupMembership {
			identitystore_DescribeGroupMembership(cfg, client)
			return
		}
		if _identitystoreDescribeUser {
			identitystore_DescribeUser(cfg, client)
			return
		}
		if _identitystoreGetGroupId {
			identitystore_GetGroupId(cfg, client)
			return
		}
		if _identitystoreGetGroupMembershipId {
			identitystore_GetGroupMembershipId(cfg, client)
			return
		}
		if _identitystoreGetUserId {
			identitystore_GetUserId(cfg, client)
			return
		}
		if _identitystoreIsMemberInGroups {
			identitystore_IsMemberInGroups(cfg, client)
			return
		}
		if _identitystoreListGroupMemberships {
			identitystore_ListGroupMemberships(cfg, client)
			return
		}
		if _identitystoreListGroupMembershipsForMember {
			identitystore_ListGroupMembershipsForMember(cfg, client)
			return
		}
		if _identitystoreListGroups {
			identitystore_ListGroups(cfg, client)
			return
		}
		if _identitystoreListUsers {
			identitystore_ListUsers(cfg, client)
			return
		}
		if _identitystoreUpdateGroup {
			identitystore_UpdateGroup(cfg, client)
			return
		}
		if _identitystoreUpdateUser {
			identitystore_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_identitystoreCreateGroup                   bool
	_identitystoreCreateGroupMembership         bool
	_identitystoreCreateUser                    bool
	_identitystoreDeleteGroup                   bool
	_identitystoreDeleteGroupMembership         bool
	_identitystoreDeleteUser                    bool
	_identitystoreDescribeGroup                 bool
	_identitystoreDescribeGroupMembership       bool
	_identitystoreDescribeUser                  bool
	_identitystoreGetGroupId                    bool
	_identitystoreGetGroupMembershipId          bool
	_identitystoreGetUserId                     bool
	_identitystoreIsMemberInGroups              bool
	_identitystoreListGroupMemberships          bool
	_identitystoreListGroupMembershipsForMember bool
	_identitystoreListGroups                    bool
	_identitystoreListUsers                     bool
	_identitystoreUpdateGroup                   bool
	_identitystoreUpdateUser                    bool

	_identitystoreAddresses           string
	_identitystoreAlternateIdentifier string
	_identitystoreBirthdate           string
	_identitystoreDescription         string
	_identitystoreDisplayName         string
	_identitystoreEmails              string
	_identitystoreExtensions          []string
	_identitystoreFilters             string
	_identitystoreGroupId             string
	_identitystoreGroupIds            []string
	_identitystoreIdentityStoreId     string
	_identitystoreLocale              string
	_identitystoreMaxResults          string
	_identitystoreMemberId            string
	_identitystoreMembershipId        string
	_identitystoreName                string
	_identitystoreNextToken           string
	_identitystoreNickName            string
	_identitystoreOperations          string
	_identitystorePhoneNumbers        string
	_identitystorePhotos              string
	_identitystorePreferredLanguage   string
	_identitystoreProfileUrl          string
	_identitystoreRoles               string
	_identitystoreTimezone            string
	_identitystoreTitle               string
	_identitystoreUserId              string
	_identitystoreUserName            string
	_identitystoreUserType            string
	_identitystoreWebsite             string
)

// Creates a group within the specified identity store.
func identitystore_CreateGroup(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.CreateGroupInput{
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreDescription) > 0 {
		input.Description = aws.String(_identitystoreDescription)
	}
	if len(_identitystoreDisplayName) > 0 {
		input.DisplayName = aws.String(_identitystoreDisplayName)
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a relationship between a member and a group. The following identifiers
// must be specified: GroupId , IdentityStoreId , and MemberId .
func identitystore_CreateGroupMembership(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.CreateGroupMembershipInput{
		// GroupId: *string, // Required
		// IdentityStoreId: *string, // Required
		// MemberId: types.MemberId, // Required
	}

	if len(_identitystoreGroupId) > 0 {
		input.GroupId = aws.String(_identitystoreGroupId)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMemberId) > 0 {
		if err := assignInputField(input, "MemberId", _identitystoreMemberId); err != nil {
			log.Errorf("invalid --member-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroupMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a user within the specified identity store.
func identitystore_CreateUser(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.CreateUserInput{
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreAddresses) > 0 {
		if err := assignInputField(input, "Addresses", _identitystoreAddresses); err != nil {
			log.Errorf("invalid --addresses: %s", err.Error())
			return
		}
	}
	if len(_identitystoreBirthdate) > 0 {
		input.Birthdate = aws.String(_identitystoreBirthdate)
	}
	if len(_identitystoreDisplayName) > 0 {
		input.DisplayName = aws.String(_identitystoreDisplayName)
	}
	if len(_identitystoreEmails) > 0 {
		if err := assignInputField(input, "Emails", _identitystoreEmails); err != nil {
			log.Errorf("invalid --emails: %s", err.Error())
			return
		}
	}
	if len(_identitystoreExtensions) > 0 {
		if err := assignInputField(input, "Extensions", _identitystoreExtensions[0]); err != nil {
			log.Errorf("invalid --extensions: %s", err.Error())
			return
		}
	}
	if len(_identitystoreLocale) > 0 {
		input.Locale = aws.String(_identitystoreLocale)
	}
	if len(_identitystoreName) > 0 {
		if err := assignInputField(input, "Name", _identitystoreName); err != nil {
			log.Errorf("invalid --name: %s", err.Error())
			return
		}
	}
	if len(_identitystoreNickName) > 0 {
		input.NickName = aws.String(_identitystoreNickName)
	}
	if len(_identitystorePhoneNumbers) > 0 {
		if err := assignInputField(input, "PhoneNumbers", _identitystorePhoneNumbers); err != nil {
			log.Errorf("invalid --phone-numbers: %s", err.Error())
			return
		}
	}
	if len(_identitystorePhotos) > 0 {
		if err := assignInputField(input, "Photos", _identitystorePhotos); err != nil {
			log.Errorf("invalid --photos: %s", err.Error())
			return
		}
	}
	if len(_identitystorePreferredLanguage) > 0 {
		input.PreferredLanguage = aws.String(_identitystorePreferredLanguage)
	}
	if len(_identitystoreProfileUrl) > 0 {
		input.ProfileUrl = aws.String(_identitystoreProfileUrl)
	}
	if len(_identitystoreRoles) > 0 {
		if err := assignInputField(input, "Roles", _identitystoreRoles); err != nil {
			log.Errorf("invalid --roles: %s", err.Error())
			return
		}
	}
	if len(_identitystoreTimezone) > 0 {
		input.Timezone = aws.String(_identitystoreTimezone)
	}
	if len(_identitystoreTitle) > 0 {
		input.Title = aws.String(_identitystoreTitle)
	}
	if len(_identitystoreUserName) > 0 {
		input.UserName = aws.String(_identitystoreUserName)
	}
	if len(_identitystoreUserType) > 0 {
		input.UserType = aws.String(_identitystoreUserType)
	}
	if len(_identitystoreWebsite) > 0 {
		input.Website = aws.String(_identitystoreWebsite)
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a group within an identity store given GroupId .
func identitystore_DeleteGroup(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.DeleteGroupInput{
		// GroupId: *string, // Required
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreGroupId) > 0 {
		input.GroupId = aws.String(_identitystoreGroupId)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Delete a membership within a group given MembershipId .
func identitystore_DeleteGroupMembership(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.DeleteGroupMembershipInput{
		// IdentityStoreId: *string, // Required
		// MembershipId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMembershipId) > 0 {
		input.MembershipId = aws.String(_identitystoreMembershipId)
	}

	if resp, err := client.DeleteGroupMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user within an identity store given UserId .
func identitystore_DeleteUser(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.DeleteUserInput{
		// IdentityStoreId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreUserId) > 0 {
		input.UserId = aws.String(_identitystoreUserId)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the group metadata and attributes from GroupId in an identity store.
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_DescribeGroup(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.DescribeGroupInput{
		// GroupId: *string, // Required
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreGroupId) > 0 {
		input.GroupId = aws.String(_identitystoreGroupId)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}

	if resp, err := client.DescribeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves membership metadata and attributes from MembershipId in an identity
// store.
//
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_DescribeGroupMembership(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.DescribeGroupMembershipInput{
		// IdentityStoreId: *string, // Required
		// MembershipId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMembershipId) > 0 {
		input.MembershipId = aws.String(_identitystoreMembershipId)
	}

	if resp, err := client.DescribeGroupMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the user metadata and attributes from the UserId in an identity store.
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_DescribeUser(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.DescribeUserInput{
		// IdentityStoreId: *string, // Required
		// UserId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreUserId) > 0 {
		input.UserId = aws.String(_identitystoreUserId)
	}
	if len(_identitystoreExtensions) > 0 {
		input.Extensions = append([]string(nil), _identitystoreExtensions...)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves GroupId in an identity store.
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_GetGroupId(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.GetGroupIdInput{
		// AlternateIdentifier: types.AlternateIdentifier, // Required
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreAlternateIdentifier) > 0 {
		if err := assignInputField(input, "AlternateIdentifier", _identitystoreAlternateIdentifier); err != nil {
			log.Errorf("invalid --alternate-identifier: %s", err.Error())
			return
		}
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}

	if resp, err := client.GetGroupId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the MembershipId in an identity store.
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_GetGroupMembershipId(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.GetGroupMembershipIdInput{
		// GroupId: *string, // Required
		// IdentityStoreId: *string, // Required
		// MemberId: types.MemberId, // Required
	}

	if len(_identitystoreGroupId) > 0 {
		input.GroupId = aws.String(_identitystoreGroupId)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMemberId) > 0 {
		if err := assignInputField(input, "MemberId", _identitystoreMemberId); err != nil {
			log.Errorf("invalid --member-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetGroupMembershipId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the UserId in an identity store.
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_GetUserId(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.GetUserIdInput{
		// AlternateIdentifier: types.AlternateIdentifier, // Required
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreAlternateIdentifier) > 0 {
		if err := assignInputField(input, "AlternateIdentifier", _identitystoreAlternateIdentifier); err != nil {
			log.Errorf("invalid --alternate-identifier: %s", err.Error())
			return
		}
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}

	if resp, err := client.GetUserId(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Checks the user's membership in all requested groups and returns if the member
// exists in all queried groups.
//
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_IsMemberInGroups(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.IsMemberInGroupsInput{
		// GroupIds: []string, // Required
		// IdentityStoreId: *string, // Required
		// MemberId: types.MemberId, // Required
	}

	if len(_identitystoreGroupIds) > 0 {
		input.GroupIds = append([]string(nil), _identitystoreGroupIds...)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMemberId) > 0 {
		if err := assignInputField(input, "MemberId", _identitystoreMemberId); err != nil {
			log.Errorf("invalid --member-id: %s", err.Error())
			return
		}
	}

	if resp, err := client.IsMemberInGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// For the specified group in the specified identity store, returns the list of
// all GroupMembership objects and returns results in paginated form.
//
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_ListGroupMemberships(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.ListGroupMembershipsInput{
		// GroupId: *string, // Required
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreGroupId) > 0 {
		input.GroupId = aws.String(_identitystoreGroupId)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _identitystoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_identitystoreNextToken) > 0 {
		input.NextToken = aws.String(_identitystoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*identitystore.ListGroupMembershipsOutput
	p := identitystore.NewListGroupMembershipsPaginator(client, input)
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

// For the specified member in the specified identity store, returns the list of
// all GroupMembership objects and returns results in paginated form.
//
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_ListGroupMembershipsForMember(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.ListGroupMembershipsForMemberInput{
		// IdentityStoreId: *string, // Required
		// MemberId: types.MemberId, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreMemberId) > 0 {
		if err := assignInputField(input, "MemberId", _identitystoreMemberId); err != nil {
			log.Errorf("invalid --member-id: %s", err.Error())
			return
		}
	}
	if len(_identitystoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _identitystoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_identitystoreNextToken) > 0 {
		input.NextToken = aws.String(_identitystoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupMembershipsForMember(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*identitystore.ListGroupMembershipsForMemberOutput
	p := identitystore.NewListGroupMembershipsForMemberPaginator(client, input)
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

// Lists all groups in the identity store. Returns a paginated list of complete
// Group objects. Filtering for a Group by the DisplayName attribute is
// deprecated. Instead, use the GetGroupId API action.
//
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_ListGroups(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.ListGroupsInput{
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreFilters) > 0 {
		if err := assignInputField(input, "Filters", _identitystoreFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_identitystoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _identitystoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_identitystoreNextToken) > 0 {
		input.NextToken = aws.String(_identitystoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*identitystore.ListGroupsOutput
	p := identitystore.NewListGroupsPaginator(client, input)
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

// Lists all users in the identity store. Returns a paginated list of complete User
// objects. Filtering for a User by the UserName attribute is deprecated. Instead,
// use the GetUserId API action.
//
// If you have access to a member account, you can use this API operation from the
// member account. For more information, see [Limiting access to the identity store from member accounts]in the IAM Identity Center User Guide.
//
// [Limiting access to the identity store from member accounts]: https://docs.aws.amazon.com/singlesignon/latest/userguide/manage-your-accounts.html#limiting-access-from-member-accounts
func identitystore_ListUsers(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.ListUsersInput{
		// IdentityStoreId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreExtensions) > 0 {
		input.Extensions = append([]string(nil), _identitystoreExtensions...)
	}
	if len(_identitystoreFilters) > 0 {
		if err := assignInputField(input, "Filters", _identitystoreFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_identitystoreMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _identitystoreMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_identitystoreNextToken) > 0 {
		input.NextToken = aws.String(_identitystoreNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*identitystore.ListUsersOutput
	p := identitystore.NewListUsersPaginator(client, input)
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

// Updates the specified group metadata and attributes in the specified identity
// store.
func identitystore_UpdateGroup(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.UpdateGroupInput{
		// GroupId: *string, // Required
		// IdentityStoreId: *string, // Required
		// Operations: []types.AttributeOperation, // Required
	}

	if len(_identitystoreGroupId) > 0 {
		input.GroupId = aws.String(_identitystoreGroupId)
	}
	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreOperations) > 0 {
		if err := assignInputField(input, "Operations", _identitystoreOperations); err != nil {
			log.Errorf("invalid --operations: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified user metadata and attributes in the specified identity
// store.
func identitystore_UpdateUser(cfg aws.Config, client *identitystore.Client) {
	input := &identitystore.UpdateUserInput{
		// IdentityStoreId: *string, // Required
		// Operations: []types.AttributeOperation, // Required
		// UserId: *string, // Required
	}

	if len(_identitystoreIdentityStoreId) > 0 {
		input.IdentityStoreId = aws.String(_identitystoreIdentityStoreId)
	}
	if len(_identitystoreOperations) > 0 {
		if err := assignInputField(input, "Operations", _identitystoreOperations); err != nil {
			log.Errorf("invalid --operations: %s", err.Error())
			return
		}
	}
	if len(_identitystoreUserId) > 0 {
		input.UserId = aws.String(_identitystoreUserId)
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_identitystoreCmd)
	_identitystoreCmd.Flags().SortFlags = false

	_identitystoreCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_identitystoreCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_identitystoreCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_identitystoreCmd.Flags().StringVarP(&_identitystoreAddresses, "addresses", "", "", "Addresses")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreAlternateIdentifier, "alternate-identifier", "", "", "Alternate Identifier")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreBirthdate, "birthdate", "", "", "Birthdate")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreDescription, "description", "", "", "Description")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreDisplayName, "display-name", "", "", "Display Name")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreEmails, "emails", "", "", "Emails")
	_identitystoreCmd.Flags().StringSliceVarP(&_identitystoreExtensions, "extensions", "", nil, "Extensions")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreFilters, "filters", "", "", "Filters")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreGroupId, "group-id", "", "", "Group ID")
	_identitystoreCmd.Flags().StringSliceVarP(&_identitystoreGroupIds, "group-ids", "", nil, "Group Ids")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreIdentityStoreId, "identity-store-id", "", "", "Identity Store ID")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreLocale, "locale", "", "", "Locale")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreMaxResults, "max-results", "", "", "Max Results")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreMemberId, "member-id", "", "", "Member ID")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreMembershipId, "membership-id", "", "", "Membership ID")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreName, "name", "", "", "Name")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreNextToken, "next-token", "", "", "Next Token")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreNickName, "nick-name", "", "", "Nick Name")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreOperations, "operations", "", "", "Operations")
	_identitystoreCmd.Flags().StringVarP(&_identitystorePhoneNumbers, "phone-numbers", "", "", "Phone Numbers")
	_identitystoreCmd.Flags().StringVarP(&_identitystorePhotos, "photos", "", "", "Photos")
	_identitystoreCmd.Flags().StringVarP(&_identitystorePreferredLanguage, "preferred-language", "", "", "Preferred Language")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreProfileUrl, "profile-url", "", "", "Profile URL")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreRoles, "roles", "", "", "Roles")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreTimezone, "timezone", "", "", "Timezone")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreTitle, "title", "", "", "Title")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreUserId, "user-id", "", "", "User ID")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreUserName, "user-name", "", "", "User Name")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreUserType, "user-type", "", "", "User Type")
	_identitystoreCmd.Flags().StringVarP(&_identitystoreWebsite, "website", "", "", "Website")

	_identitystoreCmd.Flags().BoolVarP(&_identitystoreCreateGroup, "create-group", "", false, "Create Group")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreCreateGroupMembership, "create-group-membership", "", false, "Create Group Membership")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreCreateUser, "create-user", "", false, "Create User")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreDeleteGroup, "delete-group", "", false, "Delete Group")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreDeleteGroupMembership, "delete-group-membership", "", false, "Delete Group Membership")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreDeleteUser, "delete-user", "", false, "Delete User")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreDescribeGroup, "describe-group", "", false, "Describe Group")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreDescribeGroupMembership, "describe-group-membership", "", false, "Describe Group Membership")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreDescribeUser, "describe-user", "", false, "Describe User")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreGetGroupId, "get-group-id", "", false, "Get Group ID")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreGetGroupMembershipId, "get-group-membership-id", "", false, "Get Group Membership ID")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreGetUserId, "get-user-id", "", false, "Get User ID")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreIsMemberInGroups, "is-member-in-groups", "", false, "Is Member In Groups")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreListGroupMemberships, "list-group-memberships", "", false, "List Group Memberships")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreListGroupMembershipsForMember, "list-group-memberships-for-member", "", false, "List Group Memberships For Member")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreListGroups, "list-groups", "", false, "List Groups")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreListUsers, "list-users", "", false, "List Users")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreUpdateGroup, "update-group", "", false, "Update Group")
	_identitystoreCmd.Flags().BoolVarP(&_identitystoreUpdateUser, "update-user", "", false, "Update User")

}
