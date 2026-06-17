package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralaccount"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// partnercentralaccountCmd represents the partnercentralaccount command
var _partnercentralaccountCmd = &cobra.Command{
	Use:   "partnercentralaccount",
	Short: "AWS partnercentralaccount CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := partnercentralaccount.NewFromConfig(cfg)
		if _partnercentralaccountAcceptConnectionInvitation {
			partnercentralaccount_AcceptConnectionInvitation(cfg, client)
			return
		}
		if _partnercentralaccountAssociateAwsTrainingCertificationEmailDomain {
			partnercentralaccount_AssociateAwsTrainingCertificationEmailDomain(cfg, client)
			return
		}
		if _partnercentralaccountCancelConnection {
			partnercentralaccount_CancelConnection(cfg, client)
			return
		}
		if _partnercentralaccountCancelConnectionInvitation {
			partnercentralaccount_CancelConnectionInvitation(cfg, client)
			return
		}
		if _partnercentralaccountCancelProfileUpdateTask {
			partnercentralaccount_CancelProfileUpdateTask(cfg, client)
			return
		}
		if _partnercentralaccountCreateConnectionInvitation {
			partnercentralaccount_CreateConnectionInvitation(cfg, client)
			return
		}
		if _partnercentralaccountCreatePartner {
			partnercentralaccount_CreatePartner(cfg, client)
			return
		}
		if _partnercentralaccountDisassociateAwsTrainingCertificationEmailDomain {
			partnercentralaccount_DisassociateAwsTrainingCertificationEmailDomain(cfg, client)
			return
		}
		if _partnercentralaccountGetAllianceLeadContact {
			partnercentralaccount_GetAllianceLeadContact(cfg, client)
			return
		}
		if _partnercentralaccountGetConnection {
			partnercentralaccount_GetConnection(cfg, client)
			return
		}
		if _partnercentralaccountGetConnectionInvitation {
			partnercentralaccount_GetConnectionInvitation(cfg, client)
			return
		}
		if _partnercentralaccountGetConnectionPreferences {
			partnercentralaccount_GetConnectionPreferences(cfg, client)
			return
		}
		if _partnercentralaccountGetPartner {
			partnercentralaccount_GetPartner(cfg, client)
			return
		}
		if _partnercentralaccountGetProfileUpdateTask {
			partnercentralaccount_GetProfileUpdateTask(cfg, client)
			return
		}
		if _partnercentralaccountGetProfileVisibility {
			partnercentralaccount_GetProfileVisibility(cfg, client)
			return
		}
		if _partnercentralaccountGetVerification {
			partnercentralaccount_GetVerification(cfg, client)
			return
		}
		if _partnercentralaccountListConnectionInvitations {
			partnercentralaccount_ListConnectionInvitations(cfg, client)
			return
		}
		if _partnercentralaccountListConnections {
			partnercentralaccount_ListConnections(cfg, client)
			return
		}
		if _partnercentralaccountListPartners {
			partnercentralaccount_ListPartners(cfg, client)
			return
		}
		if _partnercentralaccountListTagsForResource {
			partnercentralaccount_ListTagsForResource(cfg, client)
			return
		}
		if _partnercentralaccountPutAllianceLeadContact {
			partnercentralaccount_PutAllianceLeadContact(cfg, client)
			return
		}
		if _partnercentralaccountPutProfileVisibility {
			partnercentralaccount_PutProfileVisibility(cfg, client)
			return
		}
		if _partnercentralaccountRejectConnectionInvitation {
			partnercentralaccount_RejectConnectionInvitation(cfg, client)
			return
		}
		if _partnercentralaccountSendEmailVerificationCode {
			partnercentralaccount_SendEmailVerificationCode(cfg, client)
			return
		}
		if _partnercentralaccountStartProfileUpdateTask {
			partnercentralaccount_StartProfileUpdateTask(cfg, client)
			return
		}
		if _partnercentralaccountStartVerification {
			partnercentralaccount_StartVerification(cfg, client)
			return
		}
		if _partnercentralaccountTagResource {
			partnercentralaccount_TagResource(cfg, client)
			return
		}
		if _partnercentralaccountUntagResource {
			partnercentralaccount_UntagResource(cfg, client)
			return
		}
		if _partnercentralaccountUpdateConnectionPreferences {
			partnercentralaccount_UpdateConnectionPreferences(cfg, client)
			return
		}

	},
}

var (
	_partnercentralaccountAcceptConnectionInvitation                      bool
	_partnercentralaccountAssociateAwsTrainingCertificationEmailDomain    bool
	_partnercentralaccountCancelConnection                                bool
	_partnercentralaccountCancelConnectionInvitation                      bool
	_partnercentralaccountCancelProfileUpdateTask                         bool
	_partnercentralaccountCreateConnectionInvitation                      bool
	_partnercentralaccountCreatePartner                                   bool
	_partnercentralaccountDisassociateAwsTrainingCertificationEmailDomain bool
	_partnercentralaccountGetAllianceLeadContact                          bool
	_partnercentralaccountGetConnection                                   bool
	_partnercentralaccountGetConnectionInvitation                         bool
	_partnercentralaccountGetConnectionPreferences                        bool
	_partnercentralaccountGetPartner                                      bool
	_partnercentralaccountGetProfileUpdateTask                            bool
	_partnercentralaccountGetProfileVisibility                            bool
	_partnercentralaccountGetVerification                                 bool
	_partnercentralaccountListConnectionInvitations                       bool
	_partnercentralaccountListConnections                                 bool
	_partnercentralaccountListPartners                                    bool
	_partnercentralaccountListTagsForResource                             bool
	_partnercentralaccountPutAllianceLeadContact                          bool
	_partnercentralaccountPutProfileVisibility                            bool
	_partnercentralaccountRejectConnectionInvitation                      bool
	_partnercentralaccountSendEmailVerificationCode                       bool
	_partnercentralaccountStartProfileUpdateTask                          bool
	_partnercentralaccountStartVerification                               bool
	_partnercentralaccountTagResource                                     bool
	_partnercentralaccountUntagResource                                   bool
	_partnercentralaccountUpdateConnectionPreferences                     bool

	_partnercentralaccountAccessType                     string
	_partnercentralaccountAllianceLeadContact            string
	_partnercentralaccountCatalog                        string
	_partnercentralaccountClientToken                    string
	_partnercentralaccountConnectionType                 string
	_partnercentralaccountDomainName                     string
	_partnercentralaccountEmail                          string
	_partnercentralaccountEmailVerificationCode          string
	_partnercentralaccountExcludedParticipantIdentifiers []string
	_partnercentralaccountIdentifier                     string
	_partnercentralaccountLegalName                      string
	_partnercentralaccountMaxResults                     string
	_partnercentralaccountMessage                        string
	_partnercentralaccountName                           string
	_partnercentralaccountNextToken                      string
	_partnercentralaccountOtherParticipantIdentifiers    []string
	_partnercentralaccountParticipantType                string
	_partnercentralaccountPrimarySolutionType            string
	_partnercentralaccountReason                         string
	_partnercentralaccountReceiverIdentifier             string
	_partnercentralaccountResourceArn                    string
	_partnercentralaccountRevision                       string
	_partnercentralaccountStatus                         string
	_partnercentralaccountTagKeys                        []string
	_partnercentralaccountTags                           string
	_partnercentralaccountTaskDetails                    string
	_partnercentralaccountTaskId                         string
	_partnercentralaccountVerificationDetails            string
	_partnercentralaccountVerificationType               string
	_partnercentralaccountVisibility                     string
)

// Accepts a connection invitation from another partner, establishing a formal
// partnership connection between the two parties.
func partnercentralaccount_AcceptConnectionInvitation(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.AcceptConnectionInvitationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.AcceptConnectionInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an email domain with AWS training and certification for the partner
// account, enabling automatic verification of employee certifications.
func partnercentralaccount_AssociateAwsTrainingCertificationEmailDomain(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.AssociateAwsTrainingCertificationEmailDomainInput{
		// Catalog: *string, // Required
		// Email: *string, // Required
		// EmailVerificationCode: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountEmail) > 0 {
		input.Email = aws.String(_partnercentralaccountEmail)
	}
	if len(_partnercentralaccountEmailVerificationCode) > 0 {
		input.EmailVerificationCode = aws.String(_partnercentralaccountEmailVerificationCode)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}

	if resp, err := client.AssociateAwsTrainingCertificationEmailDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an existing connection between partners, terminating the partnership
// relationship.
func partnercentralaccount_CancelConnection(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.CancelConnectionInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// ConnectionType: types.ConnectionType, // Required
		// Identifier: *string, // Required
		// Reason: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _partnercentralaccountConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountReason) > 0 {
		input.Reason = aws.String(_partnercentralaccountReason)
	}

	if resp, err := client.CancelConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a pending connection invitation before it has been accepted or rejected.
func partnercentralaccount_CancelConnectionInvitation(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.CancelConnectionInvitationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.CancelConnectionInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels an in-progress profile update task, stopping any pending changes to the
// partner profile.
func partnercentralaccount_CancelProfileUpdateTask(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.CancelProfileUpdateTaskInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// TaskId: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountTaskId) > 0 {
		input.TaskId = aws.String(_partnercentralaccountTaskId)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}

	if resp, err := client.CancelProfileUpdateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new connection invitation to establish a partnership with another
// organization.
func partnercentralaccount_CreateConnectionInvitation(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.CreateConnectionInvitationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// ConnectionType: types.ConnectionType, // Required
		// Email: *string, // Required
		// Message: *string, // Required
		// Name: *string, // Required
		// ReceiverIdentifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _partnercentralaccountConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountEmail) > 0 {
		input.Email = aws.String(_partnercentralaccountEmail)
	}
	if len(_partnercentralaccountMessage) > 0 {
		input.Message = aws.String(_partnercentralaccountMessage)
	}
	if len(_partnercentralaccountName) > 0 {
		input.Name = aws.String(_partnercentralaccountName)
	}
	if len(_partnercentralaccountReceiverIdentifier) > 0 {
		input.ReceiverIdentifier = aws.String(_partnercentralaccountReceiverIdentifier)
	}

	if resp, err := client.CreateConnectionInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new partner account in the AWS Partner Network with the specified
// details and configuration.
func partnercentralaccount_CreatePartner(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.CreatePartnerInput{
		// AllianceLeadContact: *types.AllianceLeadContact, // Required
		// Catalog: *string, // Required
		// EmailVerificationCode: *string, // Required
		// LegalName: *string, // Required
		// PrimarySolutionType: types.PrimarySolutionType, // Required
	}

	if len(_partnercentralaccountAllianceLeadContact) > 0 {
		if err := assignInputField(input, "AllianceLeadContact", _partnercentralaccountAllianceLeadContact); err != nil {
			log.Errorf("invalid --alliance-lead-contact: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountEmailVerificationCode) > 0 {
		input.EmailVerificationCode = aws.String(_partnercentralaccountEmailVerificationCode)
	}
	if len(_partnercentralaccountLegalName) > 0 {
		input.LegalName = aws.String(_partnercentralaccountLegalName)
	}
	if len(_partnercentralaccountPrimarySolutionType) > 0 {
		if err := assignInputField(input, "PrimarySolutionType", _partnercentralaccountPrimarySolutionType); err != nil {
			log.Errorf("invalid --primary-solution-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralaccountTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePartner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between an email domain and AWS training and
// certification for the partner account.
func partnercentralaccount_DisassociateAwsTrainingCertificationEmailDomain(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.DisassociateAwsTrainingCertificationEmailDomainInput{
		// Catalog: *string, // Required
		// DomainName: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountDomainName) > 0 {
		input.DomainName = aws.String(_partnercentralaccountDomainName)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}

	if resp, err := client.DisassociateAwsTrainingCertificationEmailDomain(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the alliance lead contact information for a partner account.
func partnercentralaccount_GetAllianceLeadContact(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetAllianceLeadContactInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.GetAllianceLeadContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific connection between partners.
func partnercentralaccount_GetConnection(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetConnectionInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.GetConnection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific connection invitation.
func partnercentralaccount_GetConnectionInvitation(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetConnectionInvitationInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.GetConnectionInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the connection preferences for a partner account, including access
// settings and exclusions.
func partnercentralaccount_GetConnectionPreferences(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetConnectionPreferencesInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}

	if resp, err := client.GetConnectionPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves detailed information about a specific partner account.
func partnercentralaccount_GetPartner(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetPartnerInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.GetPartner(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves information about a specific profile update task.
func partnercentralaccount_GetProfileUpdateTask(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetProfileUpdateTaskInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.GetProfileUpdateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the visibility settings for a partner profile, determining who can
// see the profile information.
func partnercentralaccount_GetProfileVisibility(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetProfileVisibilityInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}

	if resp, err := client.GetProfileVisibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the current status and details of a verification process for a
// partner account. This operation allows partners to check the progress and
// results of business or registrant verification processes.
func partnercentralaccount_GetVerification(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.GetVerificationInput{
		// VerificationType: types.VerificationType, // Required
	}

	if len(_partnercentralaccountVerificationType) > 0 {
		if err := assignInputField(input, "VerificationType", _partnercentralaccountVerificationType); err != nil {
			log.Errorf("invalid --verification-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetVerification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists connection invitations for the partner account, with optional filtering
// by status, type, and other criteria.
func partnercentralaccount_ListConnectionInvitations(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.ListConnectionInvitationsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountConnectionType) > 0 {
		if err := assignInputField(input, "ConnectionType", _partnercentralaccountConnectionType); err != nil {
			log.Errorf("invalid --connection-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralaccountMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralaccountNextToken)
	}
	if len(_partnercentralaccountOtherParticipantIdentifiers) > 0 {
		input.OtherParticipantIdentifiers = append([]string(nil), _partnercentralaccountOtherParticipantIdentifiers...)
	}
	if len(_partnercentralaccountParticipantType) > 0 {
		if err := assignInputField(input, "ParticipantType", _partnercentralaccountParticipantType); err != nil {
			log.Errorf("invalid --participant-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountStatus) > 0 {
		if err := assignInputField(input, "Status", _partnercentralaccountStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListConnectionInvitations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralaccount.ListConnectionInvitationsOutput
	p := partnercentralaccount.NewListConnectionInvitationsPaginator(client, input)
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

// Lists active connections for the partner account, with optional filtering by
// connection type and participant.
func partnercentralaccount_ListConnections(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.ListConnectionsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountConnectionType) > 0 {
		input.ConnectionType = aws.String(_partnercentralaccountConnectionType)
	}
	if len(_partnercentralaccountMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralaccountMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralaccountNextToken)
	}
	if len(_partnercentralaccountOtherParticipantIdentifiers) > 0 {
		input.OtherParticipantIdentifiers = append([]string(nil), _partnercentralaccountOtherParticipantIdentifiers...)
	}

	if disablePaginator() {
		if resp, err := client.ListConnections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralaccount.ListConnectionsOutput
	p := partnercentralaccount.NewListConnectionsPaginator(client, input)
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

// Lists partner accounts in the catalog, providing a summary view of all partners.
func partnercentralaccount_ListPartners(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.ListPartnersInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralaccountNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPartners(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralaccount.ListPartnersOutput
	p := partnercentralaccount.NewListPartnersPaginator(client, input)
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

// Lists all tags associated with a specific AWS Partner Central Account resource.
func partnercentralaccount_ListTagsForResource(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_partnercentralaccountResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralaccountResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or updates the alliance lead contact information for a partner account.
func partnercentralaccount_PutAllianceLeadContact(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.PutAllianceLeadContactInput{
		// AllianceLeadContact: *types.AllianceLeadContact, // Required
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountAllianceLeadContact) > 0 {
		if err := assignInputField(input, "AllianceLeadContact", _partnercentralaccountAllianceLeadContact); err != nil {
			log.Errorf("invalid --alliance-lead-contact: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountEmailVerificationCode) > 0 {
		input.EmailVerificationCode = aws.String(_partnercentralaccountEmailVerificationCode)
	}

	if resp, err := client.PutAllianceLeadContact(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the visibility level for a partner profile, controlling who can view the
// profile information.
func partnercentralaccount_PutProfileVisibility(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.PutProfileVisibilityInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// Visibility: types.ProfileVisibility, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountVisibility) > 0 {
		if err := assignInputField(input, "Visibility", _partnercentralaccountVisibility); err != nil {
			log.Errorf("invalid --visibility: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutProfileVisibility(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a connection invitation from another partner, declining the partnership
// request.
func partnercentralaccount_RejectConnectionInvitation(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.RejectConnectionInvitationInput{
		// Catalog: *string, // Required
		// ClientToken: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountReason) > 0 {
		input.Reason = aws.String(_partnercentralaccountReason)
	}

	if resp, err := client.RejectConnectionInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sends an email verification code to the specified email address for account
// verification purposes.
func partnercentralaccount_SendEmailVerificationCode(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.SendEmailVerificationCodeInput{
		// Catalog: *string, // Required
		// Email: *string, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountEmail) > 0 {
		input.Email = aws.String(_partnercentralaccountEmail)
	}

	if resp, err := client.SendEmailVerificationCode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a profile update task to modify partner profile information
// asynchronously.
func partnercentralaccount_StartProfileUpdateTask(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.StartProfileUpdateTaskInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// TaskDetails: *types.TaskDetails, // Required
	}

	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralaccountIdentifier)
	}
	if len(_partnercentralaccountTaskDetails) > 0 {
		if err := assignInputField(input, "TaskDetails", _partnercentralaccountTaskDetails); err != nil {
			log.Errorf("invalid --task-details: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}

	if resp, err := client.StartProfileUpdateTask(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Initiates a new verification process for a partner account. This operation
// begins the verification workflow for either business registration or individual
// registrant identity verification as required by AWS Partner Central.
func partnercentralaccount_StartVerification(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.StartVerificationInput{}

	if len(_partnercentralaccountClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralaccountClientToken)
	}
	if len(_partnercentralaccountVerificationDetails) > 0 {
		if err := assignInputField(input, "VerificationDetails", _partnercentralaccountVerificationDetails); err != nil {
			log.Errorf("invalid --verification-details: %s", err.Error())
			return
		}
	}

	if resp, err := client.StartVerification(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a specified AWS Partner Central Account resource.
func partnercentralaccount_TagResource(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_partnercentralaccountResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralaccountResourceArn)
	}
	if len(_partnercentralaccountTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralaccountTags); err != nil {
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

// Removes specified tags from an AWS Partner Central Account resource.
func partnercentralaccount_UntagResource(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_partnercentralaccountResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralaccountResourceArn)
	}
	if len(_partnercentralaccountTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _partnercentralaccountTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the connection preferences for a partner account, modifying access
// settings and exclusions.
func partnercentralaccount_UpdateConnectionPreferences(cfg aws.Config, client *partnercentralaccount.Client) {
	input := &partnercentralaccount.UpdateConnectionPreferencesInput{
		// AccessType: types.AccessType, // Required
		// Catalog: *string, // Required
		// Revision: *int64, // Required
	}

	if len(_partnercentralaccountAccessType) > 0 {
		if err := assignInputField(input, "AccessType", _partnercentralaccountAccessType); err != nil {
			log.Errorf("invalid --access-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralaccountCatalog)
	}
	if len(_partnercentralaccountRevision) > 0 {
		if err := assignInputField(input, "Revision", _partnercentralaccountRevision); err != nil {
			log.Errorf("invalid --revision: %s", err.Error())
			return
		}
	}
	if len(_partnercentralaccountExcludedParticipantIdentifiers) > 0 {
		input.ExcludedParticipantIdentifiers = append([]string(nil), _partnercentralaccountExcludedParticipantIdentifiers...)
	}

	if resp, err := client.UpdateConnectionPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_partnercentralaccountCmd)
	_partnercentralaccountCmd.Flags().SortFlags = false

	_partnercentralaccountCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_partnercentralaccountCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_partnercentralaccountCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountAccessType, "access-type", "", "", "Access Type")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountAllianceLeadContact, "alliance-lead-contact", "", "", "Alliance Lead Contact")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountCatalog, "catalog", "", "", "Catalog")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountClientToken, "client-token", "", "", "Client Token")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountConnectionType, "connection-type", "", "", "Connection Type")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountDomainName, "domain-name", "", "", "Domain Name")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountEmail, "email", "", "", "Email")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountEmailVerificationCode, "email-verification-code", "", "", "Email Verification Code")
	_partnercentralaccountCmd.Flags().StringSliceVarP(&_partnercentralaccountExcludedParticipantIdentifiers, "excluded-participant-identifiers", "", nil, "Excluded Participant Identifiers")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountIdentifier, "identifier", "", "", "Identifier")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountLegalName, "legal-name", "", "", "Legal Name")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountMaxResults, "max-results", "", "", "Max Results")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountMessage, "message", "", "", "Message")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountName, "name", "", "", "Name")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountNextToken, "next-token", "", "", "Next Token")
	_partnercentralaccountCmd.Flags().StringSliceVarP(&_partnercentralaccountOtherParticipantIdentifiers, "other-participant-identifiers", "", nil, "Other Participant Identifiers")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountParticipantType, "participant-type", "", "", "Participant Type")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountPrimarySolutionType, "primary-solution-type", "", "", "Primary Solution Type")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountReason, "reason", "", "", "Reason")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountReceiverIdentifier, "receiver-identifier", "", "", "Receiver Identifier")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountResourceArn, "resource-arn", "", "", "Resource ARN")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountRevision, "revision", "", "", "Revision")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountStatus, "status", "", "", "Status")
	_partnercentralaccountCmd.Flags().StringSliceVarP(&_partnercentralaccountTagKeys, "tag-keys", "", nil, "Tag Keys")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountTags, "tags", "", "", "Tags")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountTaskDetails, "task-details", "", "", "Task Details")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountTaskId, "task-id", "", "", "Task ID")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountVerificationDetails, "verification-details", "", "", "Verification Details")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountVerificationType, "verification-type", "", "", "Verification Type")
	_partnercentralaccountCmd.Flags().StringVarP(&_partnercentralaccountVisibility, "visibility", "", "", "Visibility")

	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountAcceptConnectionInvitation, "accept-connection-invitation", "", false, "Accept Connection Invitation")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountAssociateAwsTrainingCertificationEmailDomain, "associate-aws-training-certification-email-domain", "", false, "Associate AWS Training Certification Email Domain")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountCancelConnection, "cancel-connection", "", false, "Cancel Connection")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountCancelConnectionInvitation, "cancel-connection-invitation", "", false, "Cancel Connection Invitation")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountCancelProfileUpdateTask, "cancel-profile-update-task", "", false, "Cancel Profile Update Task")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountCreateConnectionInvitation, "create-connection-invitation", "", false, "Create Connection Invitation")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountCreatePartner, "create-partner", "", false, "Create Partner")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountDisassociateAwsTrainingCertificationEmailDomain, "disassociate-aws-training-certification-email-domain", "", false, "Disassociate AWS Training Certification Email Domain")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetAllianceLeadContact, "get-alliance-lead-contact", "", false, "Get Alliance Lead Contact")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetConnection, "get-connection", "", false, "Get Connection")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetConnectionInvitation, "get-connection-invitation", "", false, "Get Connection Invitation")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetConnectionPreferences, "get-connection-preferences", "", false, "Get Connection Preferences")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetPartner, "get-partner", "", false, "Get Partner")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetProfileUpdateTask, "get-profile-update-task", "", false, "Get Profile Update Task")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetProfileVisibility, "get-profile-visibility", "", false, "Get Profile Visibility")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountGetVerification, "get-verification", "", false, "Get Verification")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountListConnectionInvitations, "list-connection-invitations", "", false, "List Connection Invitations")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountListConnections, "list-connections", "", false, "List Connections")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountListPartners, "list-partners", "", false, "List Partners")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountPutAllianceLeadContact, "put-alliance-lead-contact", "", false, "Put Alliance Lead Contact")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountPutProfileVisibility, "put-profile-visibility", "", false, "Put Profile Visibility")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountRejectConnectionInvitation, "reject-connection-invitation", "", false, "Reject Connection Invitation")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountSendEmailVerificationCode, "send-email-verification-code", "", false, "Send Email Verification Code")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountStartProfileUpdateTask, "start-profile-update-task", "", false, "Start Profile Update Task")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountStartVerification, "start-verification", "", false, "Start Verification")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountTagResource, "tag-resource", "", false, "Tag Resource")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountUntagResource, "untag-resource", "", false, "Untag Resource")
	_partnercentralaccountCmd.Flags().BoolVarP(&_partnercentralaccountUpdateConnectionPreferences, "update-connection-preferences", "", false, "Update Connection Preferences")

}
