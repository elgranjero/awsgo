package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/directoryservicedata"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// directoryservicedataCmd represents the directoryservicedata command
var _directoryservicedataCmd = &cobra.Command{
	Use:   "directoryservicedata",
	Short: "AWS directoryservicedata CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := directoryservicedata.NewFromConfig(cfg)
		if _directoryservicedataAddGroupMember {
			directoryservicedata_AddGroupMember(cfg, client)
			return
		}
		if _directoryservicedataCreateGroup {
			directoryservicedata_CreateGroup(cfg, client)
			return
		}
		if _directoryservicedataCreateUser {
			directoryservicedata_CreateUser(cfg, client)
			return
		}
		if _directoryservicedataDeleteGroup {
			directoryservicedata_DeleteGroup(cfg, client)
			return
		}
		if _directoryservicedataDeleteUser {
			directoryservicedata_DeleteUser(cfg, client)
			return
		}
		if _directoryservicedataDescribeGroup {
			directoryservicedata_DescribeGroup(cfg, client)
			return
		}
		if _directoryservicedataDescribeUser {
			directoryservicedata_DescribeUser(cfg, client)
			return
		}
		if _directoryservicedataDisableUser {
			directoryservicedata_DisableUser(cfg, client)
			return
		}
		if _directoryservicedataListGroupMembers {
			directoryservicedata_ListGroupMembers(cfg, client)
			return
		}
		if _directoryservicedataListGroups {
			directoryservicedata_ListGroups(cfg, client)
			return
		}
		if _directoryservicedataListGroupsForMember {
			directoryservicedata_ListGroupsForMember(cfg, client)
			return
		}
		if _directoryservicedataListUsers {
			directoryservicedata_ListUsers(cfg, client)
			return
		}
		if _directoryservicedataRemoveGroupMember {
			directoryservicedata_RemoveGroupMember(cfg, client)
			return
		}
		if _directoryservicedataSearchGroups {
			directoryservicedata_SearchGroups(cfg, client)
			return
		}
		if _directoryservicedataSearchUsers {
			directoryservicedata_SearchUsers(cfg, client)
			return
		}
		if _directoryservicedataUpdateGroup {
			directoryservicedata_UpdateGroup(cfg, client)
			return
		}
		if _directoryservicedataUpdateUser {
			directoryservicedata_UpdateUser(cfg, client)
			return
		}

	},
}

var (
	_directoryservicedataAddGroupMember      bool
	_directoryservicedataCreateGroup         bool
	_directoryservicedataCreateUser          bool
	_directoryservicedataDeleteGroup         bool
	_directoryservicedataDeleteUser          bool
	_directoryservicedataDescribeGroup       bool
	_directoryservicedataDescribeUser        bool
	_directoryservicedataDisableUser         bool
	_directoryservicedataListGroupMembers    bool
	_directoryservicedataListGroups          bool
	_directoryservicedataListGroupsForMember bool
	_directoryservicedataListUsers           bool
	_directoryservicedataRemoveGroupMember   bool
	_directoryservicedataSearchGroups        bool
	_directoryservicedataSearchUsers         bool
	_directoryservicedataUpdateGroup         bool
	_directoryservicedataUpdateUser          bool

	_directoryservicedataClientToken      string
	_directoryservicedataDirectoryId      string
	_directoryservicedataEmailAddress     string
	_directoryservicedataGivenName        string
	_directoryservicedataGroupName        string
	_directoryservicedataGroupScope       string
	_directoryservicedataGroupType        string
	_directoryservicedataMaxResults       string
	_directoryservicedataMemberName       string
	_directoryservicedataMemberRealm      string
	_directoryservicedataNextToken        string
	_directoryservicedataOtherAttributes  string
	_directoryservicedataRealm            string
	_directoryservicedataSAMAccountName   string
	_directoryservicedataSearchAttributes []string
	_directoryservicedataSearchString     string
	_directoryservicedataSurname          string
	_directoryservicedataUpdateType       string
)

// Adds an existing user, group, or computer as a group member.
func directoryservicedata_AddGroupMember(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.AddGroupMemberInput{
		// DirectoryId: *string, // Required
		// GroupName: *string, // Required
		// MemberName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataGroupName) > 0 {
		input.GroupName = aws.String(_directoryservicedataGroupName)
	}
	if len(_directoryservicedataMemberName) > 0 {
		input.MemberName = aws.String(_directoryservicedataMemberName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}
	if len(_directoryservicedataMemberRealm) > 0 {
		input.MemberRealm = aws.String(_directoryservicedataMemberRealm)
	}

	if resp, err := client.AddGroupMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new group.
func directoryservicedata_CreateGroup(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.CreateGroupInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}
	if len(_directoryservicedataGroupScope) > 0 {
		if err := assignInputField(input, "GroupScope", _directoryservicedataGroupScope); err != nil {
			log.Errorf("invalid --group-scope: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataGroupType) > 0 {
		if err := assignInputField(input, "GroupType", _directoryservicedataGroupType); err != nil {
			log.Errorf("invalid --group-type: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataOtherAttributes) > 0 {
		if err := assignInputField(input, "OtherAttributes", _directoryservicedataOtherAttributes); err != nil {
			log.Errorf("invalid --other-attributes: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new user.
func directoryservicedata_CreateUser(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.CreateUserInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}
	if len(_directoryservicedataEmailAddress) > 0 {
		input.EmailAddress = aws.String(_directoryservicedataEmailAddress)
	}
	if len(_directoryservicedataGivenName) > 0 {
		input.GivenName = aws.String(_directoryservicedataGivenName)
	}
	if len(_directoryservicedataOtherAttributes) > 0 {
		if err := assignInputField(input, "OtherAttributes", _directoryservicedataOtherAttributes); err != nil {
			log.Errorf("invalid --other-attributes: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataSurname) > 0 {
		input.Surname = aws.String(_directoryservicedataSurname)
	}

	if resp, err := client.CreateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a group.
func directoryservicedata_DeleteGroup(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.DeleteGroupInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}

	if resp, err := client.DeleteGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a user.
func directoryservicedata_DeleteUser(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.DeleteUserInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}

	if resp, err := client.DeleteUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific group.
func directoryservicedata_DescribeGroup(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.DescribeGroupInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataOtherAttributes) > 0 {
		input.OtherAttributes = []string{_directoryservicedataOtherAttributes}
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
	}

	if resp, err := client.DescribeGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about a specific user.
func directoryservicedata_DescribeUser(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.DescribeUserInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataOtherAttributes) > 0 {
		input.OtherAttributes = []string{_directoryservicedataOtherAttributes}
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
	}

	if resp, err := client.DescribeUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deactivates an active user account. For information about how to enable an
// inactive user account, see [ResetUserPassword]in the Directory Service API Reference.
//
// [ResetUserPassword]: https://docs.aws.amazon.com/directoryservice/latest/devguide/API_ResetUserPassword.html
func directoryservicedata_DisableUser(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.DisableUserInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}

	if resp, err := client.DisableUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns member information for the specified group.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// ListGroupMembers.NextToken member contains a token that you pass in the next
// call to ListGroupMembers . This retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
func directoryservicedata_ListGroupMembers(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.ListGroupMembersInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directoryservicedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataMemberRealm) > 0 {
		input.MemberRealm = aws.String(_directoryservicedataMemberRealm)
	}
	if len(_directoryservicedataNextToken) > 0 {
		input.NextToken = aws.String(_directoryservicedataNextToken)
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservicedata.ListGroupMembersOutput
	p := directoryservicedata.NewListGroupMembersPaginator(client, input)
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

// Returns group information for the specified directory.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the ListGroups.NextToken
// member contains a token that you pass in the next call to ListGroups . This
// retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
func directoryservicedata_ListGroups(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.ListGroupsInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directoryservicedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataNextToken) > 0 {
		input.NextToken = aws.String(_directoryservicedataNextToken)
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
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

	var results []*directoryservicedata.ListGroupsOutput
	p := directoryservicedata.NewListGroupsPaginator(client, input)
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

// Returns group information for the specified member.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the
// ListGroupsForMember.NextToken member contains a token that you pass in the next
// call to ListGroupsForMember . This retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
func directoryservicedata_ListGroupsForMember(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.ListGroupsForMemberInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directoryservicedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataMemberRealm) > 0 {
		input.MemberRealm = aws.String(_directoryservicedataMemberRealm)
	}
	if len(_directoryservicedataNextToken) > 0 {
		input.NextToken = aws.String(_directoryservicedataNextToken)
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
	}

	if disablePaginator() {
		if resp, err := client.ListGroupsForMember(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservicedata.ListGroupsForMemberOutput
	p := directoryservicedata.NewListGroupsForMemberPaginator(client, input)
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

// Returns user information for the specified directory.
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the ListUsers.NextToken
// member contains a token that you pass in the next call to ListUsers . This
// retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
func directoryservicedata_ListUsers(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.ListUsersInput{
		// DirectoryId: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directoryservicedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataNextToken) > 0 {
		input.NextToken = aws.String(_directoryservicedataNextToken)
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
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

	var results []*directoryservicedata.ListUsersOutput
	p := directoryservicedata.NewListUsersPaginator(client, input)
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

// Removes a member from a group.
func directoryservicedata_RemoveGroupMember(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.RemoveGroupMemberInput{
		// DirectoryId: *string, // Required
		// GroupName: *string, // Required
		// MemberName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataGroupName) > 0 {
		input.GroupName = aws.String(_directoryservicedataGroupName)
	}
	if len(_directoryservicedataMemberName) > 0 {
		input.MemberName = aws.String(_directoryservicedataMemberName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}
	if len(_directoryservicedataMemberRealm) > 0 {
		input.MemberRealm = aws.String(_directoryservicedataMemberRealm)
	}

	if resp, err := client.RemoveGroupMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Searches the specified directory for a group. You can find groups that match
// the SearchString parameter with the value of their attributes included in the
// SearchString parameter.
//
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the SearchGroups.NextToken
// member contains a token that you pass in the next call to SearchGroups . This
// retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
func directoryservicedata_SearchGroups(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.SearchGroupsInput{
		// DirectoryId: *string, // Required
		// SearchAttributes: []string, // Required
		// SearchString: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSearchAttributes) > 0 {
		input.SearchAttributes = append([]string(nil), _directoryservicedataSearchAttributes...)
	}
	if len(_directoryservicedataSearchString) > 0 {
		input.SearchString = aws.String(_directoryservicedataSearchString)
	}
	if len(_directoryservicedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directoryservicedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataNextToken) > 0 {
		input.NextToken = aws.String(_directoryservicedataNextToken)
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
	}

	if disablePaginator() {
		if resp, err := client.SearchGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservicedata.SearchGroupsOutput
	p := directoryservicedata.NewSearchGroupsPaginator(client, input)
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

// Searches the specified directory for a user. You can find users that match the
// SearchString parameter with the value of their attributes included in the
// SearchString parameter.
//
// This operation supports pagination with the use of the NextToken request and
// response parameters. If more results are available, the SearchUsers.NextToken
// member contains a token that you pass in the next call to SearchUsers . This
// retrieves the next set of items.
//
// You can also specify a maximum number of return results with the MaxResults
// parameter.
func directoryservicedata_SearchUsers(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.SearchUsersInput{
		// DirectoryId: *string, // Required
		// SearchAttributes: []string, // Required
		// SearchString: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSearchAttributes) > 0 {
		input.SearchAttributes = append([]string(nil), _directoryservicedataSearchAttributes...)
	}
	if len(_directoryservicedataSearchString) > 0 {
		input.SearchString = aws.String(_directoryservicedataSearchString)
	}
	if len(_directoryservicedataMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _directoryservicedataMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataNextToken) > 0 {
		input.NextToken = aws.String(_directoryservicedataNextToken)
	}
	if len(_directoryservicedataRealm) > 0 {
		input.Realm = aws.String(_directoryservicedataRealm)
	}

	if disablePaginator() {
		if resp, err := client.SearchUsers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*directoryservicedata.SearchUsersOutput
	p := directoryservicedata.NewSearchUsersPaginator(client, input)
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

// Updates group information.
func directoryservicedata_UpdateGroup(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.UpdateGroupInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}
	if len(_directoryservicedataGroupScope) > 0 {
		if err := assignInputField(input, "GroupScope", _directoryservicedataGroupScope); err != nil {
			log.Errorf("invalid --group-scope: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataGroupType) > 0 {
		if err := assignInputField(input, "GroupType", _directoryservicedataGroupType); err != nil {
			log.Errorf("invalid --group-type: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataOtherAttributes) > 0 {
		if err := assignInputField(input, "OtherAttributes", _directoryservicedataOtherAttributes); err != nil {
			log.Errorf("invalid --other-attributes: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataUpdateType) > 0 {
		if err := assignInputField(input, "UpdateType", _directoryservicedataUpdateType); err != nil {
			log.Errorf("invalid --update-type: %s", err.Error())
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

// Updates user information.
func directoryservicedata_UpdateUser(cfg aws.Config, client *directoryservicedata.Client) {
	input := &directoryservicedata.UpdateUserInput{
		// DirectoryId: *string, // Required
		// SAMAccountName: *string, // Required
	}

	if len(_directoryservicedataDirectoryId) > 0 {
		input.DirectoryId = aws.String(_directoryservicedataDirectoryId)
	}
	if len(_directoryservicedataSAMAccountName) > 0 {
		input.SAMAccountName = aws.String(_directoryservicedataSAMAccountName)
	}
	if len(_directoryservicedataClientToken) > 0 {
		input.ClientToken = aws.String(_directoryservicedataClientToken)
	}
	if len(_directoryservicedataEmailAddress) > 0 {
		input.EmailAddress = aws.String(_directoryservicedataEmailAddress)
	}
	if len(_directoryservicedataGivenName) > 0 {
		input.GivenName = aws.String(_directoryservicedataGivenName)
	}
	if len(_directoryservicedataOtherAttributes) > 0 {
		if err := assignInputField(input, "OtherAttributes", _directoryservicedataOtherAttributes); err != nil {
			log.Errorf("invalid --other-attributes: %s", err.Error())
			return
		}
	}
	if len(_directoryservicedataSurname) > 0 {
		input.Surname = aws.String(_directoryservicedataSurname)
	}
	if len(_directoryservicedataUpdateType) > 0 {
		if err := assignInputField(input, "UpdateType", _directoryservicedataUpdateType); err != nil {
			log.Errorf("invalid --update-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateUser(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_directoryservicedataCmd)
	_directoryservicedataCmd.Flags().SortFlags = false

	_directoryservicedataCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_directoryservicedataCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_directoryservicedataCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataClientToken, "client-token", "", "", "Client Token")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataDirectoryId, "directory-id", "", "", "Directory ID")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataEmailAddress, "email-address", "", "", "Email Address")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataGivenName, "given-name", "", "", "Given Name")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataGroupName, "group-name", "", "", "Group Name")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataGroupScope, "group-scope", "", "", "Group Scope")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataGroupType, "group-type", "", "", "Group Type")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataMaxResults, "max-results", "", "", "Max Results")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataMemberName, "member-name", "", "", "Member Name")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataMemberRealm, "member-realm", "", "", "Member Realm")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataNextToken, "next-token", "", "", "Next Token")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataOtherAttributes, "other-attributes", "", "", "Other Attributes")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataRealm, "realm", "", "", "Realm")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataSAMAccountName, "sam-account-name", "", "", "Sam Account Name")
	_directoryservicedataCmd.Flags().StringSliceVarP(&_directoryservicedataSearchAttributes, "search-attributes", "", nil, "Search Attributes")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataSearchString, "search-string", "", "", "Search String")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataSurname, "surname", "", "", "Surname")
	_directoryservicedataCmd.Flags().StringVarP(&_directoryservicedataUpdateType, "update-type", "", "", "Update Type")

	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataAddGroupMember, "add-group-member", "", false, "Add Group Member")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataCreateGroup, "create-group", "", false, "Create Group")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataCreateUser, "create-user", "", false, "Create User")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataDeleteGroup, "delete-group", "", false, "Delete Group")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataDeleteUser, "delete-user", "", false, "Delete User")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataDescribeGroup, "describe-group", "", false, "Describe Group")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataDescribeUser, "describe-user", "", false, "Describe User")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataDisableUser, "disable-user", "", false, "Disable User")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataListGroupMembers, "list-group-members", "", false, "List Group Members")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataListGroups, "list-groups", "", false, "List Groups")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataListGroupsForMember, "list-groups-for-member", "", false, "List Groups For Member")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataListUsers, "list-users", "", false, "List Users")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataRemoveGroupMember, "remove-group-member", "", false, "Remove Group Member")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataSearchGroups, "search-groups", "", false, "Search Groups")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataSearchUsers, "search-users", "", false, "Search Users")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataUpdateGroup, "update-group", "", false, "Update Group")
	_directoryservicedataCmd.Flags().BoolVarP(&_directoryservicedataUpdateUser, "update-user", "", false, "Update User")

}
