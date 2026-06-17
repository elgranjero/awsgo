package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/detective"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// detectiveCmd represents the detective command
var _detectiveCmd = &cobra.Command{
	Use:   "detective",
	Short: "AWS detective CLI",
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
		client := detective.NewFromConfig(cfg)
		if _detectiveAcceptInvitation {
			detective_AcceptInvitation(cfg, client)
			return
		}
		if _detectiveBatchGetGraphMemberDatasources {
			detective_BatchGetGraphMemberDatasources(cfg, client)
			return
		}
		if _detectiveBatchGetMembershipDatasources {
			detective_BatchGetMembershipDatasources(cfg, client)
			return
		}
		if _detectiveCreateGraph {
			detective_CreateGraph(cfg, client)
			return
		}
		if _detectiveCreateMembers {
			detective_CreateMembers(cfg, client)
			return
		}
		if _detectiveDeleteGraph {
			detective_DeleteGraph(cfg, client)
			return
		}
		if _detectiveDeleteMembers {
			detective_DeleteMembers(cfg, client)
			return
		}
		if _detectiveDescribeOrganizationConfiguration {
			detective_DescribeOrganizationConfiguration(cfg, client)
			return
		}
		if _detectiveDisableOrganizationAdminAccount {
			detective_DisableOrganizationAdminAccount(cfg, client)
			return
		}
		if _detectiveDisassociateMembership {
			detective_DisassociateMembership(cfg, client)
			return
		}
		if _detectiveEnableOrganizationAdminAccount {
			detective_EnableOrganizationAdminAccount(cfg, client)
			return
		}
		if _detectiveGetInvestigation {
			detective_GetInvestigation(cfg, client)
			return
		}
		if _detectiveGetMembers {
			detective_GetMembers(cfg, client)
			return
		}
		if _detectiveListDatasourcePackages {
			detective_ListDatasourcePackages(cfg, client)
			return
		}
		if _detectiveListGraphs {
			detective_ListGraphs(cfg, client)
			return
		}
		if _detectiveListIndicators {
			detective_ListIndicators(cfg, client)
			return
		}
		if _detectiveListInvestigations {
			detective_ListInvestigations(cfg, client)
			return
		}
		if _detectiveListInvitations {
			detective_ListInvitations(cfg, client)
			return
		}
		if _detectiveListMembers {
			detective_ListMembers(cfg, client)
			return
		}
		if _detectiveListOrganizationAdminAccounts {
			detective_ListOrganizationAdminAccounts(cfg, client)
			return
		}
		if _detectiveListTagsForResource {
			detective_ListTagsForResource(cfg, client)
			return
		}
		if _detectiveRejectInvitation {
			detective_RejectInvitation(cfg, client)
			return
		}
		if _detectiveStartInvestigation {
			detective_StartInvestigation(cfg, client)
			return
		}
		if _detectiveStartMonitoringMember {
			detective_StartMonitoringMember(cfg, client)
			return
		}
		if _detectiveTagResource {
			detective_TagResource(cfg, client)
			return
		}
		if _detectiveUntagResource {
			detective_UntagResource(cfg, client)
			return
		}
		if _detectiveUpdateDatasourcePackages {
			detective_UpdateDatasourcePackages(cfg, client)
			return
		}
		if _detectiveUpdateInvestigationState {
			detective_UpdateInvestigationState(cfg, client)
			return
		}
		if _detectiveUpdateOrganizationConfiguration {
			detective_UpdateOrganizationConfiguration(cfg, client)
			return
		}

	},
}

var (
	_detectiveAcceptInvitation                  bool
	_detectiveBatchGetGraphMemberDatasources    bool
	_detectiveBatchGetMembershipDatasources     bool
	_detectiveCreateGraph                       bool
	_detectiveCreateMembers                     bool
	_detectiveDeleteGraph                       bool
	_detectiveDeleteMembers                     bool
	_detectiveDescribeOrganizationConfiguration bool
	_detectiveDisableOrganizationAdminAccount   bool
	_detectiveDisassociateMembership            bool
	_detectiveEnableOrganizationAdminAccount    bool
	_detectiveGetInvestigation                  bool
	_detectiveGetMembers                        bool
	_detectiveListDatasourcePackages            bool
	_detectiveListGraphs                        bool
	_detectiveListIndicators                    bool
	_detectiveListInvestigations                bool
	_detectiveListInvitations                   bool
	_detectiveListMembers                       bool
	_detectiveListOrganizationAdminAccounts     bool
	_detectiveListTagsForResource               bool
	_detectiveRejectInvitation                  bool
	_detectiveStartInvestigation                bool
	_detectiveStartMonitoringMember             bool
	_detectiveTagResource                       bool
	_detectiveUntagResource                     bool
	_detectiveUpdateDatasourcePackages          bool
	_detectiveUpdateInvestigationState          bool
	_detectiveUpdateOrganizationConfiguration   bool

	_detectiveAccountId                string
	_detectiveAccountIds               []string
	_detectiveAccounts                 string
	_detectiveAutoEnable               string
	_detectiveDatasourcePackages       string
	_detectiveDisableEmailNotification string
	_detectiveEntityArn                string
	_detectiveFilterCriteria           string
	_detectiveGraphArn                 string
	_detectiveGraphArns                []string
	_detectiveIndicatorType            string
	_detectiveInvestigationId          string
	_detectiveMaxResults               string
	_detectiveMessage                  string
	_detectiveNextToken                string
	_detectiveResourceArn              string
	_detectiveScopeEndTime             string
	_detectiveScopeStartTime           string
	_detectiveSortCriteria             string
	_detectiveState                    string
	_detectiveTagKeys                  []string
	_detectiveTags                     string
)

// Accepts an invitation for the member account to contribute data to a behavior
// graph. This operation can only be called by an invited member account.
//
// The request provides the ARN of behavior graph.
//
// The member account status in the graph must be INVITED .
func detective_AcceptInvitation(cfg aws.Config, client *detective.Client) {
	input := &detective.AcceptInvitationInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.AcceptInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets data source package information for the behavior graph.
func detective_BatchGetGraphMemberDatasources(cfg aws.Config, client *detective.Client) {
	input := &detective.BatchGetGraphMemberDatasourcesInput{
		// AccountIds: []string, // Required
		// GraphArn: *string, // Required
	}

	if len(_detectiveAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _detectiveAccountIds...)
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.BatchGetGraphMemberDatasources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets information on the data source package history for an account.
func detective_BatchGetMembershipDatasources(cfg aws.Config, client *detective.Client) {
	input := &detective.BatchGetMembershipDatasourcesInput{
		// GraphArns: []string, // Required
	}

	if len(_detectiveGraphArns) > 0 {
		input.GraphArns = append([]string(nil), _detectiveGraphArns...)
	}

	if resp, err := client.BatchGetMembershipDatasources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new behavior graph for the calling account, and sets that account as
// the administrator account. This operation is called by the account that is
// enabling Detective.
//
// The operation also enables Detective for the calling account in the currently
// selected Region. It returns the ARN of the new behavior graph.
//
// CreateGraph triggers a process to create the corresponding data tables for the
// new behavior graph.
//
// An account can only be the administrator account for one behavior graph within
// a Region. If the same account calls CreateGraph with the same administrator
// account, it always returns the same behavior graph ARN. It does not create a new
// behavior graph.
func detective_CreateGraph(cfg aws.Config, client *detective.Client) {
	input := &detective.CreateGraphInput{}

	if len(_detectiveTags) > 0 {
		if err := assignInputField(input, "Tags", _detectiveTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// CreateMembers is used to send invitations to accounts. For the organization
// behavior graph, the Detective administrator account uses CreateMembers to
// enable organization accounts as member accounts.
//
// For invited accounts, CreateMembers sends a request to invite the specified
// Amazon Web Services accounts to be member accounts in the behavior graph. This
// operation can only be called by the administrator account for a behavior graph.
//
// CreateMembers verifies the accounts and then invites the verified accounts. The
// administrator can optionally specify to not send invitation emails to the member
// accounts. This would be used when the administrator manages their member
// accounts centrally.
//
// For organization accounts in the organization behavior graph, CreateMembers
// attempts to enable the accounts. The organization accounts do not receive
// invitations.
//
// The request provides the behavior graph ARN and the list of accounts to invite
// or to enable.
//
// The response separates the requested accounts into two lists:
//
// - The accounts that CreateMembers was able to process. For invited accounts,
// includes member accounts that are being verified, that have passed verification
// and are to be invited, and that have failed verification. For organization
// accounts in the organization behavior graph, includes accounts that can be
// enabled and that cannot be enabled.
//
// - The accounts that CreateMembers was unable to process. This list includes
// accounts that were already invited to be member accounts in the behavior graph.
func detective_CreateMembers(cfg aws.Config, client *detective.Client) {
	input := &detective.CreateMembersInput{
		// Accounts: []types.Account, // Required
		// GraphArn: *string, // Required
	}

	if len(_detectiveAccounts) > 0 {
		if err := assignInputField(input, "Accounts", _detectiveAccounts); err != nil {
			log.Errorf("invalid --accounts: %s", err.Error())
			return
		}
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveDisableEmailNotification) > 0 {
		if err := assignInputField(input, "DisableEmailNotification", _detectiveDisableEmailNotification); err != nil {
			log.Errorf("invalid --disable-email-notification: %s", err.Error())
			return
		}
	}
	if len(_detectiveMessage) > 0 {
		input.Message = aws.String(_detectiveMessage)
	}

	if resp, err := client.CreateMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Disables the specified behavior graph and queues it to be deleted. This
// operation removes the behavior graph from each member account's list of behavior
// graphs.
//
// DeleteGraph can only be called by the administrator account for a behavior
// graph.
func detective_DeleteGraph(cfg aws.Config, client *detective.Client) {
	input := &detective.DeleteGraphInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.DeleteGraph(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified member accounts from the behavior graph. The removed
// accounts no longer contribute data to the behavior graph. This operation can
// only be called by the administrator account for the behavior graph.
//
// For invited accounts, the removed accounts are deleted from the list of
// accounts in the behavior graph. To restore the account, the administrator
// account must send another invitation.
//
// For organization accounts in the organization behavior graph, the Detective
// administrator account can always enable the organization account again.
// Organization accounts that are not enabled as member accounts are not included
// in the ListMembers results for the organization behavior graph.
//
// An administrator account cannot use DeleteMembers to remove their own account
// from the behavior graph. To disable a behavior graph, the administrator account
// uses the DeleteGraph API method.
func detective_DeleteMembers(cfg aws.Config, client *detective.Client) {
	input := &detective.DeleteMembersInput{
		// AccountIds: []string, // Required
		// GraphArn: *string, // Required
	}

	if len(_detectiveAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _detectiveAccountIds...)
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.DeleteMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns information about the configuration for the organization behavior
// graph. Currently indicates whether to automatically enable new organization
// accounts as member accounts.
//
// Can only be called by the Detective administrator account for the organization.
func detective_DescribeOrganizationConfiguration(cfg aws.Config, client *detective.Client) {
	input := &detective.DescribeOrganizationConfigurationInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.DescribeOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the Detective administrator account in the current Region. Deletes the
// organization behavior graph.
//
// Can only be called by the organization management account.
//
// Removing the Detective administrator account does not affect the delegated
// administrator account for Detective in Organizations.
//
// To remove the delegated administrator account in Organizations, use the
// Organizations API. Removing the delegated administrator account also removes the
// Detective administrator account in all Regions, except for Regions where the
// Detective administrator account is the organization management account.
func detective_DisableOrganizationAdminAccount(cfg aws.Config, client *detective.Client) {
	input := &detective.DisableOrganizationAdminAccountInput{}

	if resp, err := client.DisableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the member account from the specified behavior graph. This operation
// can only be called by an invited member account that has the ENABLED status.
//
// DisassociateMembership cannot be called by an organization account in the
// organization behavior graph. For the organization behavior graph, the Detective
// administrator account determines which organization accounts to enable or
// disable as member accounts.
func detective_DisassociateMembership(cfg aws.Config, client *detective.Client) {
	input := &detective.DisassociateMembershipInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.DisassociateMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Designates the Detective administrator account for the organization in the
// current Region.
//
// If the account does not have Detective enabled, then enables Detective for that
// account and creates a new behavior graph.
//
// Can only be called by the organization management account.
//
// If the organization has a delegated administrator account in Organizations,
// then the Detective administrator account must be either the delegated
// administrator account or the organization management account.
//
// If the organization does not have a delegated administrator account in
// Organizations, then you can choose any account in the organization. If you
// choose an account other than the organization management account, Detective
// calls Organizations to make that account the delegated administrator account for
// Detective. The organization management account cannot be the delegated
// administrator account.
func detective_EnableOrganizationAdminAccount(cfg aws.Config, client *detective.Client) {
	input := &detective.EnableOrganizationAdminAccountInput{
		// AccountId: *string, // Required
	}

	if len(_detectiveAccountId) > 0 {
		input.AccountId = aws.String(_detectiveAccountId)
	}

	if resp, err := client.EnableOrganizationAdminAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detective investigations lets you investigate IAM users and IAM roles using
// indicators of compromise. An indicator of compromise (IOC) is an artifact
// observed in or on a network, system, or environment that can (with a high level
// of confidence) identify malicious activity or a security incident.
// GetInvestigation returns the investigation results of an investigation for a
// behavior graph.
func detective_GetInvestigation(cfg aws.Config, client *detective.Client) {
	input := &detective.GetInvestigationInput{
		// GraphArn: *string, // Required
		// InvestigationId: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveInvestigationId) > 0 {
		input.InvestigationId = aws.String(_detectiveInvestigationId)
	}

	if resp, err := client.GetInvestigation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the membership details for specified member accounts for a behavior
// graph.
func detective_GetMembers(cfg aws.Config, client *detective.Client) {
	input := &detective.GetMembersInput{
		// AccountIds: []string, // Required
		// GraphArn: *string, // Required
	}

	if len(_detectiveAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _detectiveAccountIds...)
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.GetMembers(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists data source packages in the behavior graph.
func detective_ListDatasourcePackages(cfg aws.Config, client *detective.Client) {
	input := &detective.ListDatasourcePackagesInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDatasourcePackages(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*detective.ListDatasourcePackagesOutput
	p := detective.NewListDatasourcePackagesPaginator(client, input)
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

// Returns the list of behavior graphs that the calling account is an
// administrator account of. This operation can only be called by an administrator
// account.
//
// Because an account can currently only be the administrator of one behavior
// graph within a Region, the results always contain a single behavior graph.
func detective_ListGraphs(cfg aws.Config, client *detective.Client) {
	input := &detective.ListGraphsInput{}

	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGraphs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*detective.ListGraphsOutput
	p := detective.NewListGraphsPaginator(client, input)
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

// Gets the indicators from an investigation. You can use the information from the
// indicators to determine if an IAM user and/or IAM role is involved in an unusual
// activity that could indicate malicious behavior and its impact.
func detective_ListIndicators(cfg aws.Config, client *detective.Client) {
	input := &detective.ListIndicatorsInput{
		// GraphArn: *string, // Required
		// InvestigationId: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveInvestigationId) > 0 {
		input.InvestigationId = aws.String(_detectiveInvestigationId)
	}
	if len(_detectiveIndicatorType) > 0 {
		if err := assignInputField(input, "IndicatorType", _detectiveIndicatorType); err != nil {
			log.Errorf("invalid --indicator-type: %s", err.Error())
			return
		}
	}
	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}

	if resp, err := client.ListIndicators(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detective investigations lets you investigate IAM users and IAM roles using
// indicators of compromise. An indicator of compromise (IOC) is an artifact
// observed in or on a network, system, or environment that can (with a high level
// of confidence) identify malicious activity or a security incident.
// ListInvestigations lists all active Detective investigations.
func detective_ListInvestigations(cfg aws.Config, client *detective.Client) {
	input := &detective.ListInvestigationsInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveFilterCriteria) > 0 {
		if err := assignInputField(input, "FilterCriteria", _detectiveFilterCriteria); err != nil {
			log.Errorf("invalid --filter-criteria: %s", err.Error())
			return
		}
	}
	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}
	if len(_detectiveSortCriteria) > 0 {
		if err := assignInputField(input, "SortCriteria", _detectiveSortCriteria); err != nil {
			log.Errorf("invalid --sort-criteria: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListInvestigations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the list of open and accepted behavior graph invitations for the
// member account. This operation can only be called by an invited member account.
//
// Open invitations are invitations that the member account has not responded to.
//
// The results do not include behavior graphs for which the member account
// declined the invitation. The results also do not include behavior graphs that
// the member account resigned from or was removed from.
func detective_ListInvitations(cfg aws.Config, client *detective.Client) {
	input := &detective.ListInvitationsInput{}

	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*detective.ListInvitationsOutput
	p := detective.NewListInvitationsPaginator(client, input)
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

// Retrieves the list of member accounts for a behavior graph.
// For invited accounts, the results do not include member accounts that were
// removed from the behavior graph.
//
// For the organization behavior graph, the results do not include organization
// accounts that the Detective administrator account has not enabled as member
// accounts.
func detective_ListMembers(cfg aws.Config, client *detective.Client) {
	input := &detective.ListMembersInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMembers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*detective.ListMembersOutput
	p := detective.NewListMembersPaginator(client, input)
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

// Returns information about the Detective administrator account for an
// organization. Can only be called by the organization management account.
func detective_ListOrganizationAdminAccounts(cfg aws.Config, client *detective.Client) {
	input := &detective.ListOrganizationAdminAccountsInput{}

	if len(_detectiveMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _detectiveMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_detectiveNextToken) > 0 {
		input.NextToken = aws.String(_detectiveNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListOrganizationAdminAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*detective.ListOrganizationAdminAccountsOutput
	p := detective.NewListOrganizationAdminAccountsPaginator(client, input)
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

// Returns the tag values that are assigned to a behavior graph.
func detective_ListTagsForResource(cfg aws.Config, client *detective.Client) {
	input := &detective.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_detectiveResourceArn) > 0 {
		input.ResourceArn = aws.String(_detectiveResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects an invitation to contribute the account data to a behavior graph. This
// operation must be called by an invited member account that has the INVITED
// status.
//
// RejectInvitation cannot be called by an organization account in the
// organization behavior graph. In the organization behavior graph, organization
// accounts do not receive an invitation.
func detective_RejectInvitation(cfg aws.Config, client *detective.Client) {
	input := &detective.RejectInvitationInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.RejectInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Detective investigations lets you investigate IAM users and IAM roles using
// indicators of compromise. An indicator of compromise (IOC) is an artifact
// observed in or on a network, system, or environment that can (with a high level
// of confidence) identify malicious activity or a security incident.
// StartInvestigation initiates an investigation on an entity in a behavior graph.
func detective_StartInvestigation(cfg aws.Config, client *detective.Client) {
	input := &detective.StartInvestigationInput{
		// EntityArn: *string, // Required
		// GraphArn: *string, // Required
		// ScopeEndTime: *time.Time, // Required
		// ScopeStartTime: *time.Time, // Required
	}

	if len(_detectiveEntityArn) > 0 {
		input.EntityArn = aws.String(_detectiveEntityArn)
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveScopeEndTime) > 0 {
		if err := assignInputField(input, "ScopeEndTime", _detectiveScopeEndTime); err != nil {
			log.Errorf("invalid --scope-end-time: %s", err.Error())
			return
		}
	}
	if len(_detectiveScopeStartTime) > 0 {
		if err := assignInputField(input, "ScopeStartTime", _detectiveScopeStartTime); err != nil {
			log.Errorf("invalid --scope-start-time: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartInvestigation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends a request to enable data ingest for a member account that has a status of
// ACCEPTED_BUT_DISABLED .
//
// For valid member accounts, the status is updated as follows.
//
// - If Detective enabled the member account, then the new status is ENABLED .
//
// - If Detective cannot enable the member account, the status remains
// ACCEPTED_BUT_DISABLED .
func detective_StartMonitoringMember(cfg aws.Config, client *detective.Client) {
	input := &detective.StartMonitoringMemberInput{
		// AccountId: *string, // Required
		// GraphArn: *string, // Required
	}

	if len(_detectiveAccountId) > 0 {
		input.AccountId = aws.String(_detectiveAccountId)
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.StartMonitoringMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies tag values to a behavior graph.
func detective_TagResource(cfg aws.Config, client *detective.Client) {
	input := &detective.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_detectiveResourceArn) > 0 {
		input.ResourceArn = aws.String(_detectiveResourceArn)
	}
	if len(_detectiveTags) > 0 {
		if err := assignInputField(input, "Tags", _detectiveTags); err != nil {
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

// Removes tags from a behavior graph.
func detective_UntagResource(cfg aws.Config, client *detective.Client) {
	input := &detective.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_detectiveResourceArn) > 0 {
		input.ResourceArn = aws.String(_detectiveResourceArn)
	}
	if len(_detectiveTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _detectiveTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts a data source package for the Detective behavior graph.
func detective_UpdateDatasourcePackages(cfg aws.Config, client *detective.Client) {
	input := &detective.UpdateDatasourcePackagesInput{
		// DatasourcePackages: []types.DatasourcePackage, // Required
		// GraphArn: *string, // Required
	}

	if len(_detectiveDatasourcePackages) > 0 {
		if err := assignInputField(input, "DatasourcePackages", _detectiveDatasourcePackages); err != nil {
			log.Errorf("invalid --datasource-packages: %s", err.Error())
			return
		}
	}
	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}

	if resp, err := client.UpdateDatasourcePackages(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the state of an investigation.
func detective_UpdateInvestigationState(cfg aws.Config, client *detective.Client) {
	input := &detective.UpdateInvestigationStateInput{
		// GraphArn: *string, // Required
		// InvestigationId: *string, // Required
		// State: types.State, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveInvestigationId) > 0 {
		input.InvestigationId = aws.String(_detectiveInvestigationId)
	}
	if len(_detectiveState) > 0 {
		if err := assignInputField(input, "State", _detectiveState); err != nil {
			log.Errorf("invalid --state: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateInvestigationState(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the configuration for the Organizations integration in the current
// Region. Can only be called by the Detective administrator account for the
// organization.
func detective_UpdateOrganizationConfiguration(cfg aws.Config, client *detective.Client) {
	input := &detective.UpdateOrganizationConfigurationInput{
		// GraphArn: *string, // Required
	}

	if len(_detectiveGraphArn) > 0 {
		input.GraphArn = aws.String(_detectiveGraphArn)
	}
	if len(_detectiveAutoEnable) > 0 {
		if err := assignInputField(input, "AutoEnable", _detectiveAutoEnable); err != nil {
			log.Errorf("invalid --auto-enable: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateOrganizationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_detectiveCmd)
	_detectiveCmd.Flags().SortFlags = false

	_detectiveCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_detectiveCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_detectiveCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_detectiveCmd.Flags().StringVarP(&_detectiveAccountId, "account-id", "", "", "Account ID")
	_detectiveCmd.Flags().StringSliceVarP(&_detectiveAccountIds, "account-ids", "", nil, "Account Ids")
	_detectiveCmd.Flags().StringVarP(&_detectiveAccounts, "accounts", "", "", "Accounts")
	_detectiveCmd.Flags().StringVarP(&_detectiveAutoEnable, "auto-enable", "", "", "Auto Enable")
	_detectiveCmd.Flags().StringVarP(&_detectiveDatasourcePackages, "datasource-packages", "", "", "Datasource Packages")
	_detectiveCmd.Flags().StringVarP(&_detectiveDisableEmailNotification, "disable-email-notification", "", "", "Disable Email Notification")
	_detectiveCmd.Flags().StringVarP(&_detectiveEntityArn, "entity-arn", "", "", "Entity ARN")
	_detectiveCmd.Flags().StringVarP(&_detectiveFilterCriteria, "filter-criteria", "", "", "Filter Criteria")
	_detectiveCmd.Flags().StringVarP(&_detectiveGraphArn, "graph-arn", "", "", "Graph ARN")
	_detectiveCmd.Flags().StringSliceVarP(&_detectiveGraphArns, "graph-arns", "", nil, "Graph Arns")
	_detectiveCmd.Flags().StringVarP(&_detectiveIndicatorType, "indicator-type", "", "", "Indicator Type")
	_detectiveCmd.Flags().StringVarP(&_detectiveInvestigationId, "investigation-id", "", "", "Investigation ID")
	_detectiveCmd.Flags().StringVarP(&_detectiveMaxResults, "max-results", "", "", "Max Results")
	_detectiveCmd.Flags().StringVarP(&_detectiveMessage, "message", "", "", "Message")
	_detectiveCmd.Flags().StringVarP(&_detectiveNextToken, "next-token", "", "", "Next Token")
	_detectiveCmd.Flags().StringVarP(&_detectiveResourceArn, "resource-arn", "", "", "Resource ARN")
	_detectiveCmd.Flags().StringVarP(&_detectiveScopeEndTime, "scope-end-time", "", "", "Scope End Time")
	_detectiveCmd.Flags().StringVarP(&_detectiveScopeStartTime, "scope-start-time", "", "", "Scope Start Time")
	_detectiveCmd.Flags().StringVarP(&_detectiveSortCriteria, "sort-criteria", "", "", "Sort Criteria")
	_detectiveCmd.Flags().StringVarP(&_detectiveState, "state", "", "", "State")
	_detectiveCmd.Flags().StringSliceVarP(&_detectiveTagKeys, "tag-keys", "", nil, "Tag Keys")
	_detectiveCmd.Flags().StringVarP(&_detectiveTags, "tags", "", "", "Tags")

	_detectiveCmd.Flags().BoolVarP(&_detectiveAcceptInvitation, "accept-invitation", "", false, "Accept Invitation")
	_detectiveCmd.Flags().BoolVarP(&_detectiveBatchGetGraphMemberDatasources, "batch-get-graph-member-datasources", "", false, "Batch Get Graph Member Datasources")
	_detectiveCmd.Flags().BoolVarP(&_detectiveBatchGetMembershipDatasources, "batch-get-membership-datasources", "", false, "Batch Get Membership Datasources")
	_detectiveCmd.Flags().BoolVarP(&_detectiveCreateGraph, "create-graph", "", false, "Create Graph")
	_detectiveCmd.Flags().BoolVarP(&_detectiveCreateMembers, "create-members", "", false, "Create Members")
	_detectiveCmd.Flags().BoolVarP(&_detectiveDeleteGraph, "delete-graph", "", false, "Delete Graph")
	_detectiveCmd.Flags().BoolVarP(&_detectiveDeleteMembers, "delete-members", "", false, "Delete Members")
	_detectiveCmd.Flags().BoolVarP(&_detectiveDescribeOrganizationConfiguration, "describe-organization-configuration", "", false, "Describe Organization Configuration")
	_detectiveCmd.Flags().BoolVarP(&_detectiveDisableOrganizationAdminAccount, "disable-organization-admin-account", "", false, "Disable Organization Admin Account")
	_detectiveCmd.Flags().BoolVarP(&_detectiveDisassociateMembership, "disassociate-membership", "", false, "Disassociate Membership")
	_detectiveCmd.Flags().BoolVarP(&_detectiveEnableOrganizationAdminAccount, "enable-organization-admin-account", "", false, "Enable Organization Admin Account")
	_detectiveCmd.Flags().BoolVarP(&_detectiveGetInvestigation, "get-investigation", "", false, "Get Investigation")
	_detectiveCmd.Flags().BoolVarP(&_detectiveGetMembers, "get-members", "", false, "Get Members")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListDatasourcePackages, "list-datasource-packages", "", false, "List Datasource Packages")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListGraphs, "list-graphs", "", false, "List Graphs")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListIndicators, "list-indicators", "", false, "List Indicators")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListInvestigations, "list-investigations", "", false, "List Investigations")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListInvitations, "list-invitations", "", false, "List Invitations")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListMembers, "list-members", "", false, "List Members")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListOrganizationAdminAccounts, "list-organization-admin-accounts", "", false, "List Organization Admin Accounts")
	_detectiveCmd.Flags().BoolVarP(&_detectiveListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_detectiveCmd.Flags().BoolVarP(&_detectiveRejectInvitation, "reject-invitation", "", false, "Reject Invitation")
	_detectiveCmd.Flags().BoolVarP(&_detectiveStartInvestigation, "start-investigation", "", false, "Start Investigation")
	_detectiveCmd.Flags().BoolVarP(&_detectiveStartMonitoringMember, "start-monitoring-member", "", false, "Start Monitoring Member")
	_detectiveCmd.Flags().BoolVarP(&_detectiveTagResource, "tag-resource", "", false, "Tag Resource")
	_detectiveCmd.Flags().BoolVarP(&_detectiveUntagResource, "untag-resource", "", false, "Untag Resource")
	_detectiveCmd.Flags().BoolVarP(&_detectiveUpdateDatasourcePackages, "update-datasource-packages", "", false, "Update Datasource Packages")
	_detectiveCmd.Flags().BoolVarP(&_detectiveUpdateInvestigationState, "update-investigation-state", "", false, "Update Investigation State")
	_detectiveCmd.Flags().BoolVarP(&_detectiveUpdateOrganizationConfiguration, "update-organization-configuration", "", false, "Update Organization Configuration")

}
