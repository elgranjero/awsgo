package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/securityir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// securityirCmd represents the securityir command
var _securityirCmd = &cobra.Command{
	Use:   "securityir",
	Short: "AWS securityir CLI",
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
		client := securityir.NewFromConfig(cfg)
		if _securityirBatchGetMemberAccountDetails {
			securityir_BatchGetMemberAccountDetails(cfg, client)
			return
		}
		if _securityirCancelMembership {
			securityir_CancelMembership(cfg, client)
			return
		}
		if _securityirCloseCase {
			securityir_CloseCase(cfg, client)
			return
		}
		if _securityirCreateCase {
			securityir_CreateCase(cfg, client)
			return
		}
		if _securityirCreateCaseComment {
			securityir_CreateCaseComment(cfg, client)
			return
		}
		if _securityirCreateMembership {
			securityir_CreateMembership(cfg, client)
			return
		}
		if _securityirGetCase {
			securityir_GetCase(cfg, client)
			return
		}
		if _securityirGetCaseAttachmentDownloadUrl {
			securityir_GetCaseAttachmentDownloadUrl(cfg, client)
			return
		}
		if _securityirGetCaseAttachmentUploadUrl {
			securityir_GetCaseAttachmentUploadUrl(cfg, client)
			return
		}
		if _securityirGetMembership {
			securityir_GetMembership(cfg, client)
			return
		}
		if _securityirListCaseEdits {
			securityir_ListCaseEdits(cfg, client)
			return
		}
		if _securityirListCases {
			securityir_ListCases(cfg, client)
			return
		}
		if _securityirListComments {
			securityir_ListComments(cfg, client)
			return
		}
		if _securityirListInvestigations {
			securityir_ListInvestigations(cfg, client)
			return
		}
		if _securityirListMemberships {
			securityir_ListMemberships(cfg, client)
			return
		}
		if _securityirListTagsForResource {
			securityir_ListTagsForResource(cfg, client)
			return
		}
		if _securityirSendFeedback {
			securityir_SendFeedback(cfg, client)
			return
		}
		if _securityirTagResource {
			securityir_TagResource(cfg, client)
			return
		}
		if _securityirUntagResource {
			securityir_UntagResource(cfg, client)
			return
		}
		if _securityirUpdateCase {
			securityir_UpdateCase(cfg, client)
			return
		}
		if _securityirUpdateCaseComment {
			securityir_UpdateCaseComment(cfg, client)
			return
		}
		if _securityirUpdateCaseStatus {
			securityir_UpdateCaseStatus(cfg, client)
			return
		}
		if _securityirUpdateMembership {
			securityir_UpdateMembership(cfg, client)
			return
		}
		if _securityirUpdateResolverType {
			securityir_UpdateResolverType(cfg, client)
			return
		}

	},
}

var (
	_securityirBatchGetMemberAccountDetails bool
	_securityirCancelMembership             bool
	_securityirCloseCase                    bool
	_securityirCreateCase                   bool
	_securityirCreateCaseComment            bool
	_securityirCreateMembership             bool
	_securityirGetCase                      bool
	_securityirGetCaseAttachmentDownloadUrl bool
	_securityirGetCaseAttachmentUploadUrl   bool
	_securityirGetMembership                bool
	_securityirListCaseEdits                bool
	_securityirListCases                    bool
	_securityirListComments                 bool
	_securityirListInvestigations           bool
	_securityirListMemberships              bool
	_securityirListTagsForResource          bool
	_securityirSendFeedback                 bool
	_securityirTagResource                  bool
	_securityirUntagResource                bool
	_securityirUpdateCase                   bool
	_securityirUpdateCaseComment            bool
	_securityirUpdateCaseStatus             bool
	_securityirUpdateMembership             bool
	_securityirUpdateResolverType           bool

	_securityirAccountIds                             []string
	_securityirActualIncidentStartDate                string
	_securityirAttachmentId                           string
	_securityirBody                                   string
	_securityirCaseId                                 string
	_securityirCaseMetadata                           string
	_securityirCaseStatus                             string
	_securityirClientToken                            string
	_securityirComment                                string
	_securityirCommentId                              string
	_securityirContentLength                          string
	_securityirCoverEntireOrganization                string
	_securityirDescription                            string
	_securityirEngagementType                         string
	_securityirFileName                               string
	_securityirImpactedAccounts                       []string
	_securityirImpactedAccountsToAdd                  []string
	_securityirImpactedAccountsToDelete               []string
	_securityirImpactedAwsRegions                     string
	_securityirImpactedAwsRegionsToAdd                string
	_securityirImpactedAwsRegionsToDelete             string
	_securityirImpactedServices                       []string
	_securityirImpactedServicesToAdd                  []string
	_securityirImpactedServicesToDelete               []string
	_securityirIncidentResponseTeam                   string
	_securityirMaxResults                             string
	_securityirMembershipAccountsConfigurationsUpdate string
	_securityirMembershipId                           string
	_securityirMembershipName                         string
	_securityirNextToken                              string
	_securityirOptInFeatures                          string
	_securityirReportedIncidentStartDate              string
	_securityirResolverType                           string
	_securityirResourceArn                            string
	_securityirResultId                               string
	_securityirTagKeys                                []string
	_securityirTags                                   string
	_securityirThreatActorIpAddresses                 string
	_securityirThreatActorIpAddressesToAdd            string
	_securityirThreatActorIpAddressesToDelete         string
	_securityirTitle                                  string
	_securityirUndoMembershipCancellation             string
	_securityirUsefulness                             string
	_securityirWatchers                               string
	_securityirWatchersToAdd                          string
	_securityirWatchersToDelete                       string
)

// Provides information on whether the supplied account IDs are associated with a
// membership.
//
// AWS account ID's may appear less than 12 characters and need to be
// zero-prepended. An example would be 123123123 which is nine digits, and with
// zero-prepend would be 000123123123 . Not zero-prepending to 12 digits could
// result in errors.
func securityir_BatchGetMemberAccountDetails(cfg aws.Config, client *securityir.Client) {
	input := &securityir.BatchGetMemberAccountDetailsInput{
		// AccountIds: []string, // Required
		// MembershipId: *string, // Required
	}

	if len(_securityirAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _securityirAccountIds...)
	}
	if len(_securityirMembershipId) > 0 {
		input.MembershipId = aws.String(_securityirMembershipId)
	}

	if resp, err := client.BatchGetMemberAccountDetails(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an existing membership.
func securityir_CancelMembership(cfg aws.Config, client *securityir.Client) {
	input := &securityir.CancelMembershipInput{
		// MembershipId: *string, // Required
	}

	if len(_securityirMembershipId) > 0 {
		input.MembershipId = aws.String(_securityirMembershipId)
	}

	if resp, err := client.CancelMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Closes an existing case.
func securityir_CloseCase(cfg aws.Config, client *securityir.Client) {
	input := &securityir.CloseCaseInput{
		// CaseId: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}

	if resp, err := client.CloseCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new case.
func securityir_CreateCase(cfg aws.Config, client *securityir.Client) {
	input := &securityir.CreateCaseInput{
		// Description: *string, // Required
		// EngagementType: types.EngagementType, // Required
		// ImpactedAccounts: []string, // Required
		// ReportedIncidentStartDate: *time.Time, // Required
		// ResolverType: types.ResolverType, // Required
		// Title: *string, // Required
		// Watchers: []types.Watcher, // Required
	}

	if len(_securityirDescription) > 0 {
		input.Description = aws.String(_securityirDescription)
	}
	if len(_securityirEngagementType) > 0 {
		if err := assignInputField(input, "EngagementType", _securityirEngagementType); err != nil {
			log.Errorf("invalid --engagement-type: %s", err.Error())
			return
		}
	}
	if len(_securityirImpactedAccounts) > 0 {
		input.ImpactedAccounts = append([]string(nil), _securityirImpactedAccounts...)
	}
	if len(_securityirReportedIncidentStartDate) > 0 {
		if err := assignInputField(input, "ReportedIncidentStartDate", _securityirReportedIncidentStartDate); err != nil {
			log.Errorf("invalid --reported-incident-start-date: %s", err.Error())
			return
		}
	}
	if len(_securityirResolverType) > 0 {
		if err := assignInputField(input, "ResolverType", _securityirResolverType); err != nil {
			log.Errorf("invalid --resolver-type: %s", err.Error())
			return
		}
	}
	if len(_securityirTitle) > 0 {
		input.Title = aws.String(_securityirTitle)
	}
	if len(_securityirWatchers) > 0 {
		if err := assignInputField(input, "Watchers", _securityirWatchers); err != nil {
			log.Errorf("invalid --watchers: %s", err.Error())
			return
		}
	}
	if len(_securityirClientToken) > 0 {
		input.ClientToken = aws.String(_securityirClientToken)
	}
	if len(_securityirImpactedAwsRegions) > 0 {
		if err := assignInputField(input, "ImpactedAwsRegions", _securityirImpactedAwsRegions); err != nil {
			log.Errorf("invalid --impacted-aws-regions: %s", err.Error())
			return
		}
	}
	if len(_securityirImpactedServices) > 0 {
		input.ImpactedServices = append([]string(nil), _securityirImpactedServices...)
	}
	if len(_securityirTags) > 0 {
		if err := assignInputField(input, "Tags", _securityirTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_securityirThreatActorIpAddresses) > 0 {
		if err := assignInputField(input, "ThreatActorIpAddresses", _securityirThreatActorIpAddresses); err != nil {
			log.Errorf("invalid --threat-actor-ip-addresses: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a comment to an existing case.
func securityir_CreateCaseComment(cfg aws.Config, client *securityir.Client) {
	input := &securityir.CreateCaseCommentInput{
		// Body: *string, // Required
		// CaseId: *string, // Required
	}

	if len(_securityirBody) > 0 {
		input.Body = aws.String(_securityirBody)
	}
	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirClientToken) > 0 {
		input.ClientToken = aws.String(_securityirClientToken)
	}

	if resp, err := client.CreateCaseComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new membership.
func securityir_CreateMembership(cfg aws.Config, client *securityir.Client) {
	input := &securityir.CreateMembershipInput{
		// IncidentResponseTeam: []types.IncidentResponder, // Required
		// MembershipName: *string, // Required
	}

	if len(_securityirIncidentResponseTeam) > 0 {
		if err := assignInputField(input, "IncidentResponseTeam", _securityirIncidentResponseTeam); err != nil {
			log.Errorf("invalid --incident-response-team: %s", err.Error())
			return
		}
	}
	if len(_securityirMembershipName) > 0 {
		input.MembershipName = aws.String(_securityirMembershipName)
	}
	if len(_securityirClientToken) > 0 {
		input.ClientToken = aws.String(_securityirClientToken)
	}
	if len(_securityirCoverEntireOrganization) > 0 {
		if err := assignInputField(input, "CoverEntireOrganization", _securityirCoverEntireOrganization); err != nil {
			log.Errorf("invalid --cover-entire-organization: %s", err.Error())
			return
		}
	}
	if len(_securityirOptInFeatures) > 0 {
		if err := assignInputField(input, "OptInFeatures", _securityirOptInFeatures); err != nil {
			log.Errorf("invalid --opt-in-features: %s", err.Error())
			return
		}
	}
	if len(_securityirTags) > 0 {
		if err := assignInputField(input, "Tags", _securityirTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the attributes of a case.
func securityir_GetCase(cfg aws.Config, client *securityir.Client) {
	input := &securityir.GetCaseInput{
		// CaseId: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}

	if resp, err := client.GetCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a Pre-Signed URL for uploading attachments into a case.
func securityir_GetCaseAttachmentDownloadUrl(cfg aws.Config, client *securityir.Client) {
	input := &securityir.GetCaseAttachmentDownloadUrlInput{
		// AttachmentId: *string, // Required
		// CaseId: *string, // Required
	}

	if len(_securityirAttachmentId) > 0 {
		input.AttachmentId = aws.String(_securityirAttachmentId)
	}
	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}

	if resp, err := client.GetCaseAttachmentDownloadUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads an attachment to a case.
func securityir_GetCaseAttachmentUploadUrl(cfg aws.Config, client *securityir.Client) {
	input := &securityir.GetCaseAttachmentUploadUrlInput{
		// CaseId: *string, // Required
		// ContentLength: *int64, // Required
		// FileName: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirContentLength) > 0 {
		if err := assignInputField(input, "ContentLength", _securityirContentLength); err != nil {
			log.Errorf("invalid --content-length: %s", err.Error())
			return
		}
	}
	if len(_securityirFileName) > 0 {
		input.FileName = aws.String(_securityirFileName)
	}
	if len(_securityirClientToken) > 0 {
		input.ClientToken = aws.String(_securityirClientToken)
	}

	if resp, err := client.GetCaseAttachmentUploadUrl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the attributes of a membership.
func securityir_GetMembership(cfg aws.Config, client *securityir.Client) {
	input := &securityir.GetMembershipInput{
		// MembershipId: *string, // Required
	}

	if len(_securityirMembershipId) > 0 {
		input.MembershipId = aws.String(_securityirMembershipId)
	}

	if resp, err := client.GetMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Views the case history for edits made to a designated case.
func securityir_ListCaseEdits(cfg aws.Config, client *securityir.Client) {
	input := &securityir.ListCaseEditsInput{
		// CaseId: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityirMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityirNextToken) > 0 {
		input.NextToken = aws.String(_securityirNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCaseEdits(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityir.ListCaseEditsOutput
	p := securityir.NewListCaseEditsPaginator(client, input)
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

// Lists all cases the requester has access to.
func securityir_ListCases(cfg aws.Config, client *securityir.Client) {
	input := &securityir.ListCasesInput{}

	if len(_securityirMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityirMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityirNextToken) > 0 {
		input.NextToken = aws.String(_securityirNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCases(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityir.ListCasesOutput
	p := securityir.NewListCasesPaginator(client, input)
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

// Returns comments for a designated case.
func securityir_ListComments(cfg aws.Config, client *securityir.Client) {
	input := &securityir.ListCommentsInput{
		// CaseId: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityirMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityirNextToken) > 0 {
		input.NextToken = aws.String(_securityirNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListComments(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityir.ListCommentsOutput
	p := securityir.NewListCommentsPaginator(client, input)
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

// Investigation performed by an agent for a security incident...
func securityir_ListInvestigations(cfg aws.Config, client *securityir.Client) {
	input := &securityir.ListInvestigationsInput{
		// CaseId: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityirMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityirNextToken) > 0 {
		input.NextToken = aws.String(_securityirNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListInvestigations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityir.ListInvestigationsOutput
	p := securityir.NewListInvestigationsPaginator(client, input)
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

// Returns the memberships that the calling principal can access.
func securityir_ListMemberships(cfg aws.Config, client *securityir.Client) {
	input := &securityir.ListMembershipsInput{}

	if len(_securityirMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _securityirMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_securityirNextToken) > 0 {
		input.NextToken = aws.String(_securityirNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMemberships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*securityir.ListMembershipsOutput
	p := securityir.NewListMembershipsPaginator(client, input)
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

// Returns currently configured tags on a resource.
func securityir_ListTagsForResource(cfg aws.Config, client *securityir.Client) {
	input := &securityir.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_securityirResourceArn) > 0 {
		input.ResourceArn = aws.String(_securityirResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Send feedback based on response investigation action
func securityir_SendFeedback(cfg aws.Config, client *securityir.Client) {
	input := &securityir.SendFeedbackInput{
		// CaseId: *string, // Required
		// ResultId: *string, // Required
		// Usefulness: types.UsefulnessRating, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirResultId) > 0 {
		input.ResultId = aws.String(_securityirResultId)
	}
	if len(_securityirUsefulness) > 0 {
		if err := assignInputField(input, "Usefulness", _securityirUsefulness); err != nil {
			log.Errorf("invalid --usefulness: %s", err.Error())
			return
		}
	}
	if len(_securityirComment) > 0 {
		input.Comment = aws.String(_securityirComment)
	}

	if resp, err := client.SendFeedback(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds a tag(s) to a designated resource.
func securityir_TagResource(cfg aws.Config, client *securityir.Client) {
	input := &securityir.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_securityirResourceArn) > 0 {
		input.ResourceArn = aws.String(_securityirResourceArn)
	}
	if len(_securityirTags) > 0 {
		if err := assignInputField(input, "Tags", _securityirTags); err != nil {
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

// Removes a tag(s) from a designate resource.
func securityir_UntagResource(cfg aws.Config, client *securityir.Client) {
	input := &securityir.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_securityirResourceArn) > 0 {
		input.ResourceArn = aws.String(_securityirResourceArn)
	}
	if len(_securityirTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _securityirTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing case.
func securityir_UpdateCase(cfg aws.Config, client *securityir.Client) {
	input := &securityir.UpdateCaseInput{
		// CaseId: *string, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirActualIncidentStartDate) > 0 {
		if err := assignInputField(input, "ActualIncidentStartDate", _securityirActualIncidentStartDate); err != nil {
			log.Errorf("invalid --actual-incident-start-date: %s", err.Error())
			return
		}
	}
	if len(_securityirCaseMetadata) > 0 {
		if err := assignInputField(input, "CaseMetadata", _securityirCaseMetadata); err != nil {
			log.Errorf("invalid --case-metadata: %s", err.Error())
			return
		}
	}
	if len(_securityirDescription) > 0 {
		input.Description = aws.String(_securityirDescription)
	}
	if len(_securityirEngagementType) > 0 {
		if err := assignInputField(input, "EngagementType", _securityirEngagementType); err != nil {
			log.Errorf("invalid --engagement-type: %s", err.Error())
			return
		}
	}
	if len(_securityirImpactedAccountsToAdd) > 0 {
		input.ImpactedAccountsToAdd = append([]string(nil), _securityirImpactedAccountsToAdd...)
	}
	if len(_securityirImpactedAccountsToDelete) > 0 {
		input.ImpactedAccountsToDelete = append([]string(nil), _securityirImpactedAccountsToDelete...)
	}
	if len(_securityirImpactedAwsRegionsToAdd) > 0 {
		if err := assignInputField(input, "ImpactedAwsRegionsToAdd", _securityirImpactedAwsRegionsToAdd); err != nil {
			log.Errorf("invalid --impacted-aws-regions-to-add: %s", err.Error())
			return
		}
	}
	if len(_securityirImpactedAwsRegionsToDelete) > 0 {
		if err := assignInputField(input, "ImpactedAwsRegionsToDelete", _securityirImpactedAwsRegionsToDelete); err != nil {
			log.Errorf("invalid --impacted-aws-regions-to-delete: %s", err.Error())
			return
		}
	}
	if len(_securityirImpactedServicesToAdd) > 0 {
		input.ImpactedServicesToAdd = append([]string(nil), _securityirImpactedServicesToAdd...)
	}
	if len(_securityirImpactedServicesToDelete) > 0 {
		input.ImpactedServicesToDelete = append([]string(nil), _securityirImpactedServicesToDelete...)
	}
	if len(_securityirReportedIncidentStartDate) > 0 {
		if err := assignInputField(input, "ReportedIncidentStartDate", _securityirReportedIncidentStartDate); err != nil {
			log.Errorf("invalid --reported-incident-start-date: %s", err.Error())
			return
		}
	}
	if len(_securityirThreatActorIpAddressesToAdd) > 0 {
		if err := assignInputField(input, "ThreatActorIpAddressesToAdd", _securityirThreatActorIpAddressesToAdd); err != nil {
			log.Errorf("invalid --threat-actor-ip-addresses-to-add: %s", err.Error())
			return
		}
	}
	if len(_securityirThreatActorIpAddressesToDelete) > 0 {
		if err := assignInputField(input, "ThreatActorIpAddressesToDelete", _securityirThreatActorIpAddressesToDelete); err != nil {
			log.Errorf("invalid --threat-actor-ip-addresses-to-delete: %s", err.Error())
			return
		}
	}
	if len(_securityirTitle) > 0 {
		input.Title = aws.String(_securityirTitle)
	}
	if len(_securityirWatchersToAdd) > 0 {
		if err := assignInputField(input, "WatchersToAdd", _securityirWatchersToAdd); err != nil {
			log.Errorf("invalid --watchers-to-add: %s", err.Error())
			return
		}
	}
	if len(_securityirWatchersToDelete) > 0 {
		if err := assignInputField(input, "WatchersToDelete", _securityirWatchersToDelete); err != nil {
			log.Errorf("invalid --watchers-to-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCase(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing case comment.
func securityir_UpdateCaseComment(cfg aws.Config, client *securityir.Client) {
	input := &securityir.UpdateCaseCommentInput{
		// Body: *string, // Required
		// CaseId: *string, // Required
		// CommentId: *string, // Required
	}

	if len(_securityirBody) > 0 {
		input.Body = aws.String(_securityirBody)
	}
	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirCommentId) > 0 {
		input.CommentId = aws.String(_securityirCommentId)
	}

	if resp, err := client.UpdateCaseComment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the state transitions for a designated cases.
// Self-managed: the following states are available for self-managed cases.
//
// - Submitted → Detection and Analysis
//
// - Detection and Analysis → Containment, Eradication, and Recovery
//
// - Detection and Analysis → Post-incident Activities
//
// - Containment, Eradication, and Recovery → Detection and Analysis
//
// - Containment, Eradication, and Recovery → Post-incident Activities
//
// - Post-incident Activities → Containment, Eradication, and Recovery
//
// - Post-incident Activities → Detection and Analysis
//
// - Any → Closed
//
// AWS supported: You must use the CloseCase API to close.
func securityir_UpdateCaseStatus(cfg aws.Config, client *securityir.Client) {
	input := &securityir.UpdateCaseStatusInput{
		// CaseId: *string, // Required
		// CaseStatus: types.SelfManagedCaseStatus, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirCaseStatus) > 0 {
		if err := assignInputField(input, "CaseStatus", _securityirCaseStatus); err != nil {
			log.Errorf("invalid --case-status: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateCaseStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates membership configuration.
func securityir_UpdateMembership(cfg aws.Config, client *securityir.Client) {
	input := &securityir.UpdateMembershipInput{
		// MembershipId: *string, // Required
	}

	if len(_securityirMembershipId) > 0 {
		input.MembershipId = aws.String(_securityirMembershipId)
	}
	if len(_securityirIncidentResponseTeam) > 0 {
		if err := assignInputField(input, "IncidentResponseTeam", _securityirIncidentResponseTeam); err != nil {
			log.Errorf("invalid --incident-response-team: %s", err.Error())
			return
		}
	}
	if len(_securityirMembershipAccountsConfigurationsUpdate) > 0 {
		if err := assignInputField(input, "MembershipAccountsConfigurationsUpdate", _securityirMembershipAccountsConfigurationsUpdate); err != nil {
			log.Errorf("invalid --membership-accounts-configurations-update: %s", err.Error())
			return
		}
	}
	if len(_securityirMembershipName) > 0 {
		input.MembershipName = aws.String(_securityirMembershipName)
	}
	if len(_securityirOptInFeatures) > 0 {
		if err := assignInputField(input, "OptInFeatures", _securityirOptInFeatures); err != nil {
			log.Errorf("invalid --opt-in-features: %s", err.Error())
			return
		}
	}
	if len(_securityirUndoMembershipCancellation) > 0 {
		if err := assignInputField(input, "UndoMembershipCancellation", _securityirUndoMembershipCancellation); err != nil {
			log.Errorf("invalid --undo-membership-cancellation: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMembership(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resolver type for a case.
// This is a one-way action and cannot be reversed.
func securityir_UpdateResolverType(cfg aws.Config, client *securityir.Client) {
	input := &securityir.UpdateResolverTypeInput{
		// CaseId: *string, // Required
		// ResolverType: types.ResolverType, // Required
	}

	if len(_securityirCaseId) > 0 {
		input.CaseId = aws.String(_securityirCaseId)
	}
	if len(_securityirResolverType) > 0 {
		if err := assignInputField(input, "ResolverType", _securityirResolverType); err != nil {
			log.Errorf("invalid --resolver-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateResolverType(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_securityirCmd)
	_securityirCmd.Flags().SortFlags = false

	_securityirCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_securityirCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_securityirCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_securityirCmd.Flags().StringSliceVarP(&_securityirAccountIds, "account-ids", "", nil, "Account Ids")
	_securityirCmd.Flags().StringVarP(&_securityirActualIncidentStartDate, "actual-incident-start-date", "", "", "Actual Incident Start Date")
	_securityirCmd.Flags().StringVarP(&_securityirAttachmentId, "attachment-id", "", "", "Attachment ID")
	_securityirCmd.Flags().StringVarP(&_securityirBody, "body", "", "", "Body")
	_securityirCmd.Flags().StringVarP(&_securityirCaseId, "case-id", "", "", "Case ID")
	_securityirCmd.Flags().StringVarP(&_securityirCaseMetadata, "case-metadata", "", "", "Case Metadata")
	_securityirCmd.Flags().StringVarP(&_securityirCaseStatus, "case-status", "", "", "Case Status")
	_securityirCmd.Flags().StringVarP(&_securityirClientToken, "client-token", "", "", "Client Token")
	_securityirCmd.Flags().StringVarP(&_securityirComment, "comment", "", "", "Comment")
	_securityirCmd.Flags().StringVarP(&_securityirCommentId, "comment-id", "", "", "Comment ID")
	_securityirCmd.Flags().StringVarP(&_securityirContentLength, "content-length", "", "", "Content Length")
	_securityirCmd.Flags().StringVarP(&_securityirCoverEntireOrganization, "cover-entire-organization", "", "", "Cover Entire Organization")
	_securityirCmd.Flags().StringVarP(&_securityirDescription, "description", "", "", "Description")
	_securityirCmd.Flags().StringVarP(&_securityirEngagementType, "engagement-type", "", "", "Engagement Type")
	_securityirCmd.Flags().StringVarP(&_securityirFileName, "file-name", "", "", "File Name")
	_securityirCmd.Flags().StringSliceVarP(&_securityirImpactedAccounts, "impacted-accounts", "", nil, "Impacted Accounts")
	_securityirCmd.Flags().StringSliceVarP(&_securityirImpactedAccountsToAdd, "impacted-accounts-to-add", "", nil, "Impacted Accounts To Add")
	_securityirCmd.Flags().StringSliceVarP(&_securityirImpactedAccountsToDelete, "impacted-accounts-to-delete", "", nil, "Impacted Accounts To Delete")
	_securityirCmd.Flags().StringVarP(&_securityirImpactedAwsRegions, "impacted-aws-regions", "", "", "Impacted AWS Regions")
	_securityirCmd.Flags().StringVarP(&_securityirImpactedAwsRegionsToAdd, "impacted-aws-regions-to-add", "", "", "Impacted AWS Regions To Add")
	_securityirCmd.Flags().StringVarP(&_securityirImpactedAwsRegionsToDelete, "impacted-aws-regions-to-delete", "", "", "Impacted AWS Regions To Delete")
	_securityirCmd.Flags().StringSliceVarP(&_securityirImpactedServices, "impacted-services", "", nil, "Impacted Services")
	_securityirCmd.Flags().StringSliceVarP(&_securityirImpactedServicesToAdd, "impacted-services-to-add", "", nil, "Impacted Services To Add")
	_securityirCmd.Flags().StringSliceVarP(&_securityirImpactedServicesToDelete, "impacted-services-to-delete", "", nil, "Impacted Services To Delete")
	_securityirCmd.Flags().StringVarP(&_securityirIncidentResponseTeam, "incident-response-team", "", "", "Incident Response Team")
	_securityirCmd.Flags().StringVarP(&_securityirMaxResults, "max-results", "", "", "Max Results")
	_securityirCmd.Flags().StringVarP(&_securityirMembershipAccountsConfigurationsUpdate, "membership-accounts-configurations-update", "", "", "Membership Accounts Configurations Update")
	_securityirCmd.Flags().StringVarP(&_securityirMembershipId, "membership-id", "", "", "Membership ID")
	_securityirCmd.Flags().StringVarP(&_securityirMembershipName, "membership-name", "", "", "Membership Name")
	_securityirCmd.Flags().StringVarP(&_securityirNextToken, "next-token", "", "", "Next Token")
	_securityirCmd.Flags().StringVarP(&_securityirOptInFeatures, "opt-in-features", "", "", "Opt In Features")
	_securityirCmd.Flags().StringVarP(&_securityirReportedIncidentStartDate, "reported-incident-start-date", "", "", "Reported Incident Start Date")
	_securityirCmd.Flags().StringVarP(&_securityirResolverType, "resolver-type", "", "", "Resolver Type")
	_securityirCmd.Flags().StringVarP(&_securityirResourceArn, "resource-arn", "", "", "Resource ARN")
	_securityirCmd.Flags().StringVarP(&_securityirResultId, "result-id", "", "", "Result ID")
	_securityirCmd.Flags().StringSliceVarP(&_securityirTagKeys, "tag-keys", "", nil, "Tag Keys")
	_securityirCmd.Flags().StringVarP(&_securityirTags, "tags", "", "", "Tags")
	_securityirCmd.Flags().StringVarP(&_securityirThreatActorIpAddresses, "threat-actor-ip-addresses", "", "", "Threat Actor IP Addresses")
	_securityirCmd.Flags().StringVarP(&_securityirThreatActorIpAddressesToAdd, "threat-actor-ip-addresses-to-add", "", "", "Threat Actor IP Addresses To Add")
	_securityirCmd.Flags().StringVarP(&_securityirThreatActorIpAddressesToDelete, "threat-actor-ip-addresses-to-delete", "", "", "Threat Actor IP Addresses To Delete")
	_securityirCmd.Flags().StringVarP(&_securityirTitle, "title", "", "", "Title")
	_securityirCmd.Flags().StringVarP(&_securityirUndoMembershipCancellation, "undo-membership-cancellation", "", "", "Undo Membership Cancellation")
	_securityirCmd.Flags().StringVarP(&_securityirUsefulness, "usefulness", "", "", "Usefulness")
	_securityirCmd.Flags().StringVarP(&_securityirWatchers, "watchers", "", "", "Watchers")
	_securityirCmd.Flags().StringVarP(&_securityirWatchersToAdd, "watchers-to-add", "", "", "Watchers To Add")
	_securityirCmd.Flags().StringVarP(&_securityirWatchersToDelete, "watchers-to-delete", "", "", "Watchers To Delete")

	_securityirCmd.Flags().BoolVarP(&_securityirBatchGetMemberAccountDetails, "batch-get-member-account-details", "", false, "Batch Get Member Account Details")
	_securityirCmd.Flags().BoolVarP(&_securityirCancelMembership, "cancel-membership", "", false, "Cancel Membership")
	_securityirCmd.Flags().BoolVarP(&_securityirCloseCase, "close-case", "", false, "Close Case")
	_securityirCmd.Flags().BoolVarP(&_securityirCreateCase, "create-case", "", false, "Create Case")
	_securityirCmd.Flags().BoolVarP(&_securityirCreateCaseComment, "create-case-comment", "", false, "Create Case Comment")
	_securityirCmd.Flags().BoolVarP(&_securityirCreateMembership, "create-membership", "", false, "Create Membership")
	_securityirCmd.Flags().BoolVarP(&_securityirGetCase, "get-case", "", false, "Get Case")
	_securityirCmd.Flags().BoolVarP(&_securityirGetCaseAttachmentDownloadUrl, "get-case-attachment-download-url", "", false, "Get Case Attachment Download URL")
	_securityirCmd.Flags().BoolVarP(&_securityirGetCaseAttachmentUploadUrl, "get-case-attachment-upload-url", "", false, "Get Case Attachment Upload URL")
	_securityirCmd.Flags().BoolVarP(&_securityirGetMembership, "get-membership", "", false, "Get Membership")
	_securityirCmd.Flags().BoolVarP(&_securityirListCaseEdits, "list-case-edits", "", false, "List Case Edits")
	_securityirCmd.Flags().BoolVarP(&_securityirListCases, "list-cases", "", false, "List Cases")
	_securityirCmd.Flags().BoolVarP(&_securityirListComments, "list-comments", "", false, "List Comments")
	_securityirCmd.Flags().BoolVarP(&_securityirListInvestigations, "list-investigations", "", false, "List Investigations")
	_securityirCmd.Flags().BoolVarP(&_securityirListMemberships, "list-memberships", "", false, "List Memberships")
	_securityirCmd.Flags().BoolVarP(&_securityirListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_securityirCmd.Flags().BoolVarP(&_securityirSendFeedback, "send-feedback", "", false, "Send Feedback")
	_securityirCmd.Flags().BoolVarP(&_securityirTagResource, "tag-resource", "", false, "Tag Resource")
	_securityirCmd.Flags().BoolVarP(&_securityirUntagResource, "untag-resource", "", false, "Untag Resource")
	_securityirCmd.Flags().BoolVarP(&_securityirUpdateCase, "update-case", "", false, "Update Case")
	_securityirCmd.Flags().BoolVarP(&_securityirUpdateCaseComment, "update-case-comment", "", false, "Update Case Comment")
	_securityirCmd.Flags().BoolVarP(&_securityirUpdateCaseStatus, "update-case-status", "", false, "Update Case Status")
	_securityirCmd.Flags().BoolVarP(&_securityirUpdateMembership, "update-membership", "", false, "Update Membership")
	_securityirCmd.Flags().BoolVarP(&_securityirUpdateResolverType, "update-resolver-type", "", false, "Update Resolver Type")

}
