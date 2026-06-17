package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/mpa"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// mpaCmd represents the mpa command
var _mpaCmd = &cobra.Command{
	Use:   "mpa",
	Short: "AWS mpa CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := mpa.NewFromConfig(cfg)
		if _mpaCancelSession {
			mpa_CancelSession(cfg, client)
			return
		}
		if _mpaCreateApprovalTeam {
			mpa_CreateApprovalTeam(cfg, client)
			return
		}
		if _mpaCreateIdentitySource {
			mpa_CreateIdentitySource(cfg, client)
			return
		}
		if _mpaDeleteIdentitySource {
			mpa_DeleteIdentitySource(cfg, client)
			return
		}
		if _mpaDeleteInactiveApprovalTeamVersion {
			mpa_DeleteInactiveApprovalTeamVersion(cfg, client)
			return
		}
		if _mpaGetApprovalTeam {
			mpa_GetApprovalTeam(cfg, client)
			return
		}
		if _mpaGetIdentitySource {
			mpa_GetIdentitySource(cfg, client)
			return
		}
		if _mpaGetPolicyVersion {
			mpa_GetPolicyVersion(cfg, client)
			return
		}
		if _mpaGetResourcePolicy {
			mpa_GetResourcePolicy(cfg, client)
			return
		}
		if _mpaGetSession {
			mpa_GetSession(cfg, client)
			return
		}
		if _mpaListApprovalTeams {
			mpa_ListApprovalTeams(cfg, client)
			return
		}
		if _mpaListIdentitySources {
			mpa_ListIdentitySources(cfg, client)
			return
		}
		if _mpaListPolicies {
			mpa_ListPolicies(cfg, client)
			return
		}
		if _mpaListPolicyVersions {
			mpa_ListPolicyVersions(cfg, client)
			return
		}
		if _mpaListResourcePolicies {
			mpa_ListResourcePolicies(cfg, client)
			return
		}
		if _mpaListSessions {
			mpa_ListSessions(cfg, client)
			return
		}
		if _mpaListTagsForResource {
			mpa_ListTagsForResource(cfg, client)
			return
		}
		if _mpaStartActiveApprovalTeamDeletion {
			mpa_StartActiveApprovalTeamDeletion(cfg, client)
			return
		}
		if _mpaTagResource {
			mpa_TagResource(cfg, client)
			return
		}
		if _mpaUntagResource {
			mpa_UntagResource(cfg, client)
			return
		}
		if _mpaUpdateApprovalTeam {
			mpa_UpdateApprovalTeam(cfg, client)
			return
		}

	},
}

var (
	_mpaCancelSession                     bool
	_mpaCreateApprovalTeam                bool
	_mpaCreateIdentitySource              bool
	_mpaDeleteIdentitySource              bool
	_mpaDeleteInactiveApprovalTeamVersion bool
	_mpaGetApprovalTeam                   bool
	_mpaGetIdentitySource                 bool
	_mpaGetPolicyVersion                  bool
	_mpaGetResourcePolicy                 bool
	_mpaGetSession                        bool
	_mpaListApprovalTeams                 bool
	_mpaListIdentitySources               bool
	_mpaListPolicies                      bool
	_mpaListPolicyVersions                bool
	_mpaListResourcePolicies              bool
	_mpaListSessions                      bool
	_mpaListTagsForResource               bool
	_mpaStartActiveApprovalTeamDeletion   bool
	_mpaTagResource                       bool
	_mpaUntagResource                     bool
	_mpaUpdateApprovalTeam                bool

	_mpaApprovalStrategy         string
	_mpaApprovalTeamArn          string
	_mpaApprovers                string
	_mpaArn                      string
	_mpaClientToken              string
	_mpaDescription              string
	_mpaFilters                  string
	_mpaIdentitySourceArn        string
	_mpaIdentitySourceParameters string
	_mpaMaxResults               string
	_mpaName                     string
	_mpaNextToken                string
	_mpaPendingWindowDays        string
	_mpaPolicies                 string
	_mpaPolicyArn                string
	_mpaPolicyName               string
	_mpaPolicyType               string
	_mpaPolicyVersionArn         string
	_mpaResourceArn              string
	_mpaSessionArn               string
	_mpaTagKeys                  []string
	_mpaTags                     string
	_mpaUpdateActions            string
	_mpaVersionId                string
)

// Cancels an approval session. For more information, see [Session] in the Multi-party
// approval User Guide.
//
// [Session]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_CancelSession(cfg aws.Config, client *mpa.Client) {
	input := &mpa.CancelSessionInput{
		// SessionArn: *string, // Required
	}

	if len(_mpaSessionArn) > 0 {
		input.SessionArn = aws.String(_mpaSessionArn)
	}

	if resp, err := client.CancelSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new approval team. For more information, see [Approval team] in the Multi-party
// approval User Guide.
//
// [Approval team]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_CreateApprovalTeam(cfg aws.Config, client *mpa.Client) {
	input := &mpa.CreateApprovalTeamInput{
		// ApprovalStrategy: types.ApprovalStrategy, // Required
		// Approvers: []types.ApprovalTeamRequestApprover, // Required
		// Description: *string, // Required
		// Name: *string, // Required
		// Policies: []types.PolicyReference, // Required
	}

	if len(_mpaApprovalStrategy) > 0 {
		if err := assignInputField(input, "ApprovalStrategy", _mpaApprovalStrategy); err != nil {
			log.Errorf("invalid --approval-strategy: %s", err.Error())
			return
		}
	}
	if len(_mpaApprovers) > 0 {
		if err := assignInputField(input, "Approvers", _mpaApprovers); err != nil {
			log.Errorf("invalid --approvers: %s", err.Error())
			return
		}
	}
	if len(_mpaDescription) > 0 {
		input.Description = aws.String(_mpaDescription)
	}
	if len(_mpaName) > 0 {
		input.Name = aws.String(_mpaName)
	}
	if len(_mpaPolicies) > 0 {
		if err := assignInputField(input, "Policies", _mpaPolicies); err != nil {
			log.Errorf("invalid --policies: %s", err.Error())
			return
		}
	}
	if len(_mpaClientToken) > 0 {
		input.ClientToken = aws.String(_mpaClientToken)
	}
	if len(_mpaTags) > 0 {
		if err := assignInputField(input, "Tags", _mpaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateApprovalTeam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new identity source. For more information, see [Identity Source] in the Multi-party
// approval User Guide.
//
// [Identity Source]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_CreateIdentitySource(cfg aws.Config, client *mpa.Client) {
	input := &mpa.CreateIdentitySourceInput{
		// IdentitySourceParameters: *types.IdentitySourceParameters, // Required
	}

	if len(_mpaIdentitySourceParameters) > 0 {
		if err := assignInputField(input, "IdentitySourceParameters", _mpaIdentitySourceParameters); err != nil {
			log.Errorf("invalid --identity-source-parameters: %s", err.Error())
			return
		}
	}
	if len(_mpaClientToken) > 0 {
		input.ClientToken = aws.String(_mpaClientToken)
	}
	if len(_mpaTags) > 0 {
		if err := assignInputField(input, "Tags", _mpaTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an identity source. For more information, see [Identity Source] in the Multi-party
// approval User Guide.
//
// [Identity Source]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_DeleteIdentitySource(cfg aws.Config, client *mpa.Client) {
	input := &mpa.DeleteIdentitySourceInput{
		// IdentitySourceArn: *string, // Required
	}

	if len(_mpaIdentitySourceArn) > 0 {
		input.IdentitySourceArn = aws.String(_mpaIdentitySourceArn)
	}

	if resp, err := client.DeleteIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an inactive approval team. For more information, see [Team health] in the
// Multi-party approval User Guide.
//
// You can also use this operation to delete a team draft. For more information,
// see [Interacting with drafts]in the Multi-party approval User Guide.
//
// [Interacting with drafts]: https://docs.aws.amazon.com/mpa/latest/userguide/update-team.html#update-team-draft-status
// [Team health]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-health.html
func mpa_DeleteInactiveApprovalTeamVersion(cfg aws.Config, client *mpa.Client) {
	input := &mpa.DeleteInactiveApprovalTeamVersionInput{
		// Arn: *string, // Required
		// VersionId: *string, // Required
	}

	if len(_mpaArn) > 0 {
		input.Arn = aws.String(_mpaArn)
	}
	if len(_mpaVersionId) > 0 {
		input.VersionId = aws.String(_mpaVersionId)
	}

	if resp, err := client.DeleteInactiveApprovalTeamVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for an approval team.
func mpa_GetApprovalTeam(cfg aws.Config, client *mpa.Client) {
	input := &mpa.GetApprovalTeamInput{
		// Arn: *string, // Required
	}

	if len(_mpaArn) > 0 {
		input.Arn = aws.String(_mpaArn)
	}

	if resp, err := client.GetApprovalTeam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for an identity source. For more information, see [Identity Source] in the
// Multi-party approval User Guide.
//
// [Identity Source]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_GetIdentitySource(cfg aws.Config, client *mpa.Client) {
	input := &mpa.GetIdentitySourceInput{
		// IdentitySourceArn: *string, // Required
	}

	if len(_mpaIdentitySourceArn) > 0 {
		input.IdentitySourceArn = aws.String(_mpaIdentitySourceArn)
	}

	if resp, err := client.GetIdentitySource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for the version of a policy. Policies define the permissions
// for team resources.
func mpa_GetPolicyVersion(cfg aws.Config, client *mpa.Client) {
	input := &mpa.GetPolicyVersionInput{
		// PolicyVersionArn: *string, // Required
	}

	if len(_mpaPolicyVersionArn) > 0 {
		input.PolicyVersionArn = aws.String(_mpaPolicyVersionArn)
	}

	if resp, err := client.GetPolicyVersion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details about a policy for a resource.
func mpa_GetResourcePolicy(cfg aws.Config, client *mpa.Client) {
	input := &mpa.GetResourcePolicyInput{
		// PolicyName: *string, // Required
		// PolicyType: types.PolicyType, // Required
		// ResourceArn: *string, // Required
	}

	if len(_mpaPolicyName) > 0 {
		input.PolicyName = aws.String(_mpaPolicyName)
	}
	if len(_mpaPolicyType) > 0 {
		if err := assignInputField(input, "PolicyType", _mpaPolicyType); err != nil {
			log.Errorf("invalid --policy-type: %s", err.Error())
			return
		}
	}
	if len(_mpaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mpaResourceArn)
	}

	if resp, err := client.GetResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns details for an approval session. For more information, see [Session] in the
// Multi-party approval User Guide.
//
// [Session]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_GetSession(cfg aws.Config, client *mpa.Client) {
	input := &mpa.GetSessionInput{
		// SessionArn: *string, // Required
	}

	if len(_mpaSessionArn) > 0 {
		input.SessionArn = aws.String(_mpaSessionArn)
	}

	if resp, err := client.GetSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of approval teams.
func mpa_ListApprovalTeams(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListApprovalTeamsInput{}

	if len(_mpaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mpaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mpaNextToken) > 0 {
		input.NextToken = aws.String(_mpaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListApprovalTeams(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mpa.ListApprovalTeamsOutput
	p := mpa.NewListApprovalTeamsPaginator(client, input)
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

// Returns a list of identity sources. For more information, see [Identity Source] in the
// Multi-party approval User Guide.
//
// [Identity Source]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_ListIdentitySources(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListIdentitySourcesInput{}

	if len(_mpaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mpaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mpaNextToken) > 0 {
		input.NextToken = aws.String(_mpaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListIdentitySources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mpa.ListIdentitySourcesOutput
	p := mpa.NewListIdentitySourcesPaginator(client, input)
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

// Returns a list of policies. Policies define the permissions for team resources.
func mpa_ListPolicies(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListPoliciesInput{}

	if len(_mpaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mpaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mpaNextToken) > 0 {
		input.NextToken = aws.String(_mpaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mpa.ListPoliciesOutput
	p := mpa.NewListPoliciesPaginator(client, input)
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

// Returns a list of the versions for policies. Policies define the permissions
// for team resources.
func mpa_ListPolicyVersions(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListPolicyVersionsInput{
		// PolicyArn: *string, // Required
	}

	if len(_mpaPolicyArn) > 0 {
		input.PolicyArn = aws.String(_mpaPolicyArn)
	}
	if len(_mpaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mpaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mpaNextToken) > 0 {
		input.NextToken = aws.String(_mpaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPolicyVersions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mpa.ListPolicyVersionsOutput
	p := mpa.NewListPolicyVersionsPaginator(client, input)
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

// Returns a list of policies for a resource.
func mpa_ListResourcePolicies(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListResourcePoliciesInput{
		// ResourceArn: *string, // Required
	}

	if len(_mpaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mpaResourceArn)
	}
	if len(_mpaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mpaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mpaNextToken) > 0 {
		input.NextToken = aws.String(_mpaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListResourcePolicies(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mpa.ListResourcePoliciesOutput
	p := mpa.NewListResourcePoliciesPaginator(client, input)
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

// Returns a list of approval sessions. For more information, see [Session] in the
// Multi-party approval User Guide.
//
// [Session]: https://docs.aws.amazon.com/mpa/latest/userguide/mpa-concepts.html
func mpa_ListSessions(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListSessionsInput{
		// ApprovalTeamArn: *string, // Required
	}

	if len(_mpaApprovalTeamArn) > 0 {
		input.ApprovalTeamArn = aws.String(_mpaApprovalTeamArn)
	}
	if len(_mpaFilters) > 0 {
		if err := assignInputField(input, "Filters", _mpaFilters); err != nil {
			log.Errorf("invalid --filters: %s", err.Error())
			return
		}
	}
	if len(_mpaMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _mpaMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_mpaNextToken) > 0 {
		input.NextToken = aws.String(_mpaNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListSessions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*mpa.ListSessionsOutput
	p := mpa.NewListSessionsPaginator(client, input)
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

// Returns a list of the tags for a resource.
func mpa_ListTagsForResource(cfg aws.Config, client *mpa.Client) {
	input := &mpa.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_mpaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mpaResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Starts the deletion process for an active approval team.
// # Deletions require team approval
//
// Requests to delete an active team must be approved by the team.
func mpa_StartActiveApprovalTeamDeletion(cfg aws.Config, client *mpa.Client) {
	input := &mpa.StartActiveApprovalTeamDeletionInput{
		// Arn: *string, // Required
	}

	if len(_mpaArn) > 0 {
		input.Arn = aws.String(_mpaArn)
	}
	if len(_mpaPendingWindowDays) > 0 {
		if err := assignInputField(input, "PendingWindowDays", _mpaPendingWindowDays); err != nil {
			log.Errorf("invalid --pending-window-days: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartActiveApprovalTeamDeletion(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates a resource tag. Each tag is a label consisting of a
// user-defined key and value. Tags can help you manage, identify, organize, search
// for, and filter resources.
func mpa_TagResource(cfg aws.Config, client *mpa.Client) {
	input := &mpa.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_mpaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mpaResourceArn)
	}
	if len(_mpaTags) > 0 {
		if err := assignInputField(input, "Tags", _mpaTags); err != nil {
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

// Removes a resource tag. Each tag is a label consisting of a user-defined key
// and value. Tags can help you manage, identify, organize, search for, and filter
// resources.
func mpa_UntagResource(cfg aws.Config, client *mpa.Client) {
	input := &mpa.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_mpaResourceArn) > 0 {
		input.ResourceArn = aws.String(_mpaResourceArn)
	}
	if len(_mpaTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _mpaTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an approval team. You can request to update the team description,
// approval threshold, and approvers in the team.
//
// # Updates require team approval
//
// Updates to an active team must be approved by the team.
func mpa_UpdateApprovalTeam(cfg aws.Config, client *mpa.Client) {
	input := &mpa.UpdateApprovalTeamInput{
		// Arn: *string, // Required
	}

	if len(_mpaArn) > 0 {
		input.Arn = aws.String(_mpaArn)
	}
	if len(_mpaApprovalStrategy) > 0 {
		if err := assignInputField(input, "ApprovalStrategy", _mpaApprovalStrategy); err != nil {
			log.Errorf("invalid --approval-strategy: %s", err.Error())
			return
		}
	}
	if len(_mpaApprovers) > 0 {
		if err := assignInputField(input, "Approvers", _mpaApprovers); err != nil {
			log.Errorf("invalid --approvers: %s", err.Error())
			return
		}
	}
	if len(_mpaDescription) > 0 {
		input.Description = aws.String(_mpaDescription)
	}
	if len(_mpaUpdateActions) > 0 {
		if err := assignInputField(input, "UpdateActions", _mpaUpdateActions); err != nil {
			log.Errorf("invalid --update-actions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateApprovalTeam(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_mpaCmd)
	_mpaCmd.Flags().SortFlags = false

	_mpaCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_mpaCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_mpaCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_mpaCmd.Flags().StringVarP(&_mpaApprovalStrategy, "approval-strategy", "", "", "Approval Strategy")
	_mpaCmd.Flags().StringVarP(&_mpaApprovalTeamArn, "approval-team-arn", "", "", "Approval Team ARN")
	_mpaCmd.Flags().StringVarP(&_mpaApprovers, "approvers", "", "", "Approvers")
	_mpaCmd.Flags().StringVarP(&_mpaArn, "arn", "", "", "ARN")
	_mpaCmd.Flags().StringVarP(&_mpaClientToken, "client-token", "", "", "Client Token")
	_mpaCmd.Flags().StringVarP(&_mpaDescription, "description", "", "", "Description")
	_mpaCmd.Flags().StringVarP(&_mpaFilters, "filters", "", "", "Filters")
	_mpaCmd.Flags().StringVarP(&_mpaIdentitySourceArn, "identity-source-arn", "", "", "Identity Source ARN")
	_mpaCmd.Flags().StringVarP(&_mpaIdentitySourceParameters, "identity-source-parameters", "", "", "Identity Source Parameters")
	_mpaCmd.Flags().StringVarP(&_mpaMaxResults, "max-results", "", "", "Max Results")
	_mpaCmd.Flags().StringVarP(&_mpaName, "name", "", "", "Name")
	_mpaCmd.Flags().StringVarP(&_mpaNextToken, "next-token", "", "", "Next Token")
	_mpaCmd.Flags().StringVarP(&_mpaPendingWindowDays, "pending-window-days", "", "", "Pending Window Days")
	_mpaCmd.Flags().StringVarP(&_mpaPolicies, "policies", "", "", "Policies")
	_mpaCmd.Flags().StringVarP(&_mpaPolicyArn, "policy-arn", "", "", "Policy ARN")
	_mpaCmd.Flags().StringVarP(&_mpaPolicyName, "policy-name", "", "", "Policy Name")
	_mpaCmd.Flags().StringVarP(&_mpaPolicyType, "policy-type", "", "", "Policy Type")
	_mpaCmd.Flags().StringVarP(&_mpaPolicyVersionArn, "policy-version-arn", "", "", "Policy Version ARN")
	_mpaCmd.Flags().StringVarP(&_mpaResourceArn, "resource-arn", "", "", "Resource ARN")
	_mpaCmd.Flags().StringVarP(&_mpaSessionArn, "session-arn", "", "", "Session ARN")
	_mpaCmd.Flags().StringSliceVarP(&_mpaTagKeys, "tag-keys", "", nil, "Tag Keys")
	_mpaCmd.Flags().StringVarP(&_mpaTags, "tags", "", "", "Tags")
	_mpaCmd.Flags().StringVarP(&_mpaUpdateActions, "update-actions", "", "", "Update Actions")
	_mpaCmd.Flags().StringVarP(&_mpaVersionId, "version-id", "", "", "Version ID")

	_mpaCmd.Flags().BoolVarP(&_mpaCancelSession, "cancel-session", "", false, "Cancel Session")
	_mpaCmd.Flags().BoolVarP(&_mpaCreateApprovalTeam, "create-approval-team", "", false, "Create Approval Team")
	_mpaCmd.Flags().BoolVarP(&_mpaCreateIdentitySource, "create-identity-source", "", false, "Create Identity Source")
	_mpaCmd.Flags().BoolVarP(&_mpaDeleteIdentitySource, "delete-identity-source", "", false, "Delete Identity Source")
	_mpaCmd.Flags().BoolVarP(&_mpaDeleteInactiveApprovalTeamVersion, "delete-inactive-approval-team-version", "", false, "Delete Inactive Approval Team Version")
	_mpaCmd.Flags().BoolVarP(&_mpaGetApprovalTeam, "get-approval-team", "", false, "Get Approval Team")
	_mpaCmd.Flags().BoolVarP(&_mpaGetIdentitySource, "get-identity-source", "", false, "Get Identity Source")
	_mpaCmd.Flags().BoolVarP(&_mpaGetPolicyVersion, "get-policy-version", "", false, "Get Policy Version")
	_mpaCmd.Flags().BoolVarP(&_mpaGetResourcePolicy, "get-resource-policy", "", false, "Get Resource Policy")
	_mpaCmd.Flags().BoolVarP(&_mpaGetSession, "get-session", "", false, "Get Session")
	_mpaCmd.Flags().BoolVarP(&_mpaListApprovalTeams, "list-approval-teams", "", false, "List Approval Teams")
	_mpaCmd.Flags().BoolVarP(&_mpaListIdentitySources, "list-identity-sources", "", false, "List Identity Sources")
	_mpaCmd.Flags().BoolVarP(&_mpaListPolicies, "list-policies", "", false, "List Policies")
	_mpaCmd.Flags().BoolVarP(&_mpaListPolicyVersions, "list-policy-versions", "", false, "List Policy Versions")
	_mpaCmd.Flags().BoolVarP(&_mpaListResourcePolicies, "list-resource-policies", "", false, "List Resource Policies")
	_mpaCmd.Flags().BoolVarP(&_mpaListSessions, "list-sessions", "", false, "List Sessions")
	_mpaCmd.Flags().BoolVarP(&_mpaListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_mpaCmd.Flags().BoolVarP(&_mpaStartActiveApprovalTeamDeletion, "start-active-approval-team-deletion", "", false, "Start Active Approval Team Deletion")
	_mpaCmd.Flags().BoolVarP(&_mpaTagResource, "tag-resource", "", false, "Tag Resource")
	_mpaCmd.Flags().BoolVarP(&_mpaUntagResource, "untag-resource", "", false, "Untag Resource")
	_mpaCmd.Flags().BoolVarP(&_mpaUpdateApprovalTeam, "update-approval-team", "", false, "Update Approval Team")

}
