package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/partnercentralchannel"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// partnercentralchannelCmd represents the partnercentralchannel command
var _partnercentralchannelCmd = &cobra.Command{
	Use:   "partnercentralchannel",
	Short: "AWS partnercentralchannel CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := partnercentralchannel.NewFromConfig(cfg)
		if _partnercentralchannelAcceptChannelHandshake {
			partnercentralchannel_AcceptChannelHandshake(cfg, client)
			return
		}
		if _partnercentralchannelCancelChannelHandshake {
			partnercentralchannel_CancelChannelHandshake(cfg, client)
			return
		}
		if _partnercentralchannelCreateChannelHandshake {
			partnercentralchannel_CreateChannelHandshake(cfg, client)
			return
		}
		if _partnercentralchannelCreateProgramManagementAccount {
			partnercentralchannel_CreateProgramManagementAccount(cfg, client)
			return
		}
		if _partnercentralchannelCreateRelationship {
			partnercentralchannel_CreateRelationship(cfg, client)
			return
		}
		if _partnercentralchannelDeleteProgramManagementAccount {
			partnercentralchannel_DeleteProgramManagementAccount(cfg, client)
			return
		}
		if _partnercentralchannelDeleteRelationship {
			partnercentralchannel_DeleteRelationship(cfg, client)
			return
		}
		if _partnercentralchannelGetRelationship {
			partnercentralchannel_GetRelationship(cfg, client)
			return
		}
		if _partnercentralchannelListChannelHandshakes {
			partnercentralchannel_ListChannelHandshakes(cfg, client)
			return
		}
		if _partnercentralchannelListProgramManagementAccounts {
			partnercentralchannel_ListProgramManagementAccounts(cfg, client)
			return
		}
		if _partnercentralchannelListRelationships {
			partnercentralchannel_ListRelationships(cfg, client)
			return
		}
		if _partnercentralchannelListTagsForResource {
			partnercentralchannel_ListTagsForResource(cfg, client)
			return
		}
		if _partnercentralchannelRejectChannelHandshake {
			partnercentralchannel_RejectChannelHandshake(cfg, client)
			return
		}
		if _partnercentralchannelTagResource {
			partnercentralchannel_TagResource(cfg, client)
			return
		}
		if _partnercentralchannelUntagResource {
			partnercentralchannel_UntagResource(cfg, client)
			return
		}
		if _partnercentralchannelUpdateProgramManagementAccount {
			partnercentralchannel_UpdateProgramManagementAccount(cfg, client)
			return
		}
		if _partnercentralchannelUpdateRelationship {
			partnercentralchannel_UpdateRelationship(cfg, client)
			return
		}

	},
}

var (
	_partnercentralchannelAcceptChannelHandshake         bool
	_partnercentralchannelCancelChannelHandshake         bool
	_partnercentralchannelCreateChannelHandshake         bool
	_partnercentralchannelCreateProgramManagementAccount bool
	_partnercentralchannelCreateRelationship             bool
	_partnercentralchannelDeleteProgramManagementAccount bool
	_partnercentralchannelDeleteRelationship             bool
	_partnercentralchannelGetRelationship                bool
	_partnercentralchannelListChannelHandshakes          bool
	_partnercentralchannelListProgramManagementAccounts  bool
	_partnercentralchannelListRelationships              bool
	_partnercentralchannelListTagsForResource            bool
	_partnercentralchannelRejectChannelHandshake         bool
	_partnercentralchannelTagResource                    bool
	_partnercentralchannelUntagResource                  bool
	_partnercentralchannelUpdateProgramManagementAccount bool
	_partnercentralchannelUpdateRelationship             bool

	_partnercentralchannelAccountId                           string
	_partnercentralchannelAccountIds                          []string
	_partnercentralchannelAssociatedAccountId                 string
	_partnercentralchannelAssociatedAccountIds                []string
	_partnercentralchannelAssociatedResourceIdentifier        string
	_partnercentralchannelAssociatedResourceIdentifiers       []string
	_partnercentralchannelAssociationType                     string
	_partnercentralchannelAssociationTypes                    string
	_partnercentralchannelCatalog                             string
	_partnercentralchannelClientToken                         string
	_partnercentralchannelDisplayName                         string
	_partnercentralchannelDisplayNames                        []string
	_partnercentralchannelHandshakeType                       string
	_partnercentralchannelHandshakeTypeFilters                string
	_partnercentralchannelHandshakeTypeSort                   string
	_partnercentralchannelIdentifier                          string
	_partnercentralchannelMaxResults                          string
	_partnercentralchannelNextToken                           string
	_partnercentralchannelParticipantType                     string
	_partnercentralchannelPayload                             string
	_partnercentralchannelProgram                             string
	_partnercentralchannelProgramManagementAccountIdentifier  string
	_partnercentralchannelProgramManagementAccountIdentifiers []string
	_partnercentralchannelPrograms                            string
	_partnercentralchannelRequestedSupportPlan                string
	_partnercentralchannelResaleAccountModel                  string
	_partnercentralchannelResourceArn                         string
	_partnercentralchannelRevision                            string
	_partnercentralchannelSector                              string
	_partnercentralchannelSort                                string
	_partnercentralchannelStatuses                            string
	_partnercentralchannelTagKeys                             []string
	_partnercentralchannelTags                                string
)

// Accepts a pending channel handshake request from another AWS account.
func partnercentralchannel_AcceptChannelHandshake(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.AcceptChannelHandshakeInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}

	if resp, err := client.AcceptChannelHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Cancels a pending channel handshake request.
func partnercentralchannel_CancelChannelHandshake(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.CancelChannelHandshakeInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}

	if resp, err := client.CancelChannelHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new channel handshake request to establish a partnership with another
// AWS account.
func partnercentralchannel_CreateChannelHandshake(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.CreateChannelHandshakeInput{
		// AssociatedResourceIdentifier: *string, // Required
		// Catalog: *string, // Required
		// HandshakeType: types.HandshakeType, // Required
	}

	if len(_partnercentralchannelAssociatedResourceIdentifier) > 0 {
		input.AssociatedResourceIdentifier = aws.String(_partnercentralchannelAssociatedResourceIdentifier)
	}
	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelHandshakeType) > 0 {
		if err := assignInputField(input, "HandshakeType", _partnercentralchannelHandshakeType); err != nil {
			log.Errorf("invalid --handshake-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralchannelClientToken)
	}
	if len(_partnercentralchannelPayload) > 0 {
		if err := assignInputField(input, "Payload", _partnercentralchannelPayload); err != nil {
			log.Errorf("invalid --payload: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralchannelTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateChannelHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new program management account for managing partner relationships.
func partnercentralchannel_CreateProgramManagementAccount(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.CreateProgramManagementAccountInput{
		// AccountId: *string, // Required
		// Catalog: *string, // Required
		// DisplayName: *string, // Required
		// Program: types.Program, // Required
	}

	if len(_partnercentralchannelAccountId) > 0 {
		input.AccountId = aws.String(_partnercentralchannelAccountId)
	}
	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelDisplayName) > 0 {
		input.DisplayName = aws.String(_partnercentralchannelDisplayName)
	}
	if len(_partnercentralchannelProgram) > 0 {
		if err := assignInputField(input, "Program", _partnercentralchannelProgram); err != nil {
			log.Errorf("invalid --program: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralchannelClientToken)
	}
	if len(_partnercentralchannelTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralchannelTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProgramManagementAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new partner relationship between accounts.
func partnercentralchannel_CreateRelationship(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.CreateRelationshipInput{
		// AssociatedAccountId: *string, // Required
		// AssociationType: types.AssociationType, // Required
		// Catalog: *string, // Required
		// DisplayName: *string, // Required
		// ProgramManagementAccountIdentifier: *string, // Required
		// Sector: types.Sector, // Required
	}

	if len(_partnercentralchannelAssociatedAccountId) > 0 {
		input.AssociatedAccountId = aws.String(_partnercentralchannelAssociatedAccountId)
	}
	if len(_partnercentralchannelAssociationType) > 0 {
		if err := assignInputField(input, "AssociationType", _partnercentralchannelAssociationType); err != nil {
			log.Errorf("invalid --association-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelDisplayName) > 0 {
		input.DisplayName = aws.String(_partnercentralchannelDisplayName)
	}
	if len(_partnercentralchannelProgramManagementAccountIdentifier) > 0 {
		input.ProgramManagementAccountIdentifier = aws.String(_partnercentralchannelProgramManagementAccountIdentifier)
	}
	if len(_partnercentralchannelSector) > 0 {
		if err := assignInputField(input, "Sector", _partnercentralchannelSector); err != nil {
			log.Errorf("invalid --sector: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralchannelClientToken)
	}
	if len(_partnercentralchannelRequestedSupportPlan) > 0 {
		if err := assignInputField(input, "RequestedSupportPlan", _partnercentralchannelRequestedSupportPlan); err != nil {
			log.Errorf("invalid --requested-support-plan: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelResaleAccountModel) > 0 {
		if err := assignInputField(input, "ResaleAccountModel", _partnercentralchannelResaleAccountModel); err != nil {
			log.Errorf("invalid --resale-account-model: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralchannelTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a program management account.
func partnercentralchannel_DeleteProgramManagementAccount(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.DeleteProgramManagementAccountInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}
	if len(_partnercentralchannelClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralchannelClientToken)
	}

	if resp, err := client.DeleteProgramManagementAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a partner relationship.
func partnercentralchannel_DeleteRelationship(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.DeleteRelationshipInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// ProgramManagementAccountIdentifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}
	if len(_partnercentralchannelProgramManagementAccountIdentifier) > 0 {
		input.ProgramManagementAccountIdentifier = aws.String(_partnercentralchannelProgramManagementAccountIdentifier)
	}
	if len(_partnercentralchannelClientToken) > 0 {
		input.ClientToken = aws.String(_partnercentralchannelClientToken)
	}

	if resp, err := client.DeleteRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves details of a specific partner relationship.
func partnercentralchannel_GetRelationship(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.GetRelationshipInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// ProgramManagementAccountIdentifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}
	if len(_partnercentralchannelProgramManagementAccountIdentifier) > 0 {
		input.ProgramManagementAccountIdentifier = aws.String(_partnercentralchannelProgramManagementAccountIdentifier)
	}

	if resp, err := client.GetRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists channel handshakes based on specified criteria.
func partnercentralchannel_ListChannelHandshakes(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.ListChannelHandshakesInput{
		// Catalog: *string, // Required
		// HandshakeType: types.HandshakeType, // Required
		// ParticipantType: types.ParticipantType, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelHandshakeType) > 0 {
		if err := assignInputField(input, "HandshakeType", _partnercentralchannelHandshakeType); err != nil {
			log.Errorf("invalid --handshake-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelParticipantType) > 0 {
		if err := assignInputField(input, "ParticipantType", _partnercentralchannelParticipantType); err != nil {
			log.Errorf("invalid --participant-type: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelAssociatedResourceIdentifiers) > 0 {
		input.AssociatedResourceIdentifiers = append([]string(nil), _partnercentralchannelAssociatedResourceIdentifiers...)
	}
	if len(_partnercentralchannelHandshakeTypeFilters) > 0 {
		if err := assignInputField(input, "HandshakeTypeFilters", _partnercentralchannelHandshakeTypeFilters); err != nil {
			log.Errorf("invalid --handshake-type-filters: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelHandshakeTypeSort) > 0 {
		if err := assignInputField(input, "HandshakeTypeSort", _partnercentralchannelHandshakeTypeSort); err != nil {
			log.Errorf("invalid --handshake-type-sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralchannelMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralchannelNextToken)
	}
	if len(_partnercentralchannelStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _partnercentralchannelStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListChannelHandshakes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralchannel.ListChannelHandshakesOutput
	p := partnercentralchannel.NewListChannelHandshakesPaginator(client, input)
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

// Lists program management accounts based on specified criteria.
func partnercentralchannel_ListProgramManagementAccounts(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.ListProgramManagementAccountsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelAccountIds) > 0 {
		input.AccountIds = append([]string(nil), _partnercentralchannelAccountIds...)
	}
	if len(_partnercentralchannelDisplayNames) > 0 {
		input.DisplayNames = append([]string(nil), _partnercentralchannelDisplayNames...)
	}
	if len(_partnercentralchannelMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralchannelMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralchannelNextToken)
	}
	if len(_partnercentralchannelPrograms) > 0 {
		if err := assignInputField(input, "Programs", _partnercentralchannelPrograms); err != nil {
			log.Errorf("invalid --programs: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralchannelSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelStatuses) > 0 {
		if err := assignInputField(input, "Statuses", _partnercentralchannelStatuses); err != nil {
			log.Errorf("invalid --statuses: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListProgramManagementAccounts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralchannel.ListProgramManagementAccountsOutput
	p := partnercentralchannel.NewListProgramManagementAccountsPaginator(client, input)
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

// Lists partner relationships based on specified criteria.
func partnercentralchannel_ListRelationships(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.ListRelationshipsInput{
		// Catalog: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelAssociatedAccountIds) > 0 {
		input.AssociatedAccountIds = append([]string(nil), _partnercentralchannelAssociatedAccountIds...)
	}
	if len(_partnercentralchannelAssociationTypes) > 0 {
		if err := assignInputField(input, "AssociationTypes", _partnercentralchannelAssociationTypes); err != nil {
			log.Errorf("invalid --association-types: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelDisplayNames) > 0 {
		input.DisplayNames = append([]string(nil), _partnercentralchannelDisplayNames...)
	}
	if len(_partnercentralchannelMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _partnercentralchannelMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelNextToken) > 0 {
		input.NextToken = aws.String(_partnercentralchannelNextToken)
	}
	if len(_partnercentralchannelProgramManagementAccountIdentifiers) > 0 {
		input.ProgramManagementAccountIdentifiers = append([]string(nil), _partnercentralchannelProgramManagementAccountIdentifiers...)
	}
	if len(_partnercentralchannelSort) > 0 {
		if err := assignInputField(input, "Sort", _partnercentralchannelSort); err != nil {
			log.Errorf("invalid --sort: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListRelationships(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*partnercentralchannel.ListRelationshipsOutput
	p := partnercentralchannel.NewListRelationshipsPaginator(client, input)
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

// Lists tags associated with a specific resource.
func partnercentralchannel_ListTagsForResource(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_partnercentralchannelResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralchannelResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects a pending channel handshake request.
func partnercentralchannel_RejectChannelHandshake(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.RejectChannelHandshakeInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}

	if resp, err := client.RejectChannelHandshake(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or updates tags for a specified resource.
func partnercentralchannel_TagResource(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_partnercentralchannelResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralchannelResourceArn)
	}
	if len(_partnercentralchannelTags) > 0 {
		if err := assignInputField(input, "Tags", _partnercentralchannelTags); err != nil {
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

// Removes tags from a specified resource.
func partnercentralchannel_UntagResource(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_partnercentralchannelResourceArn) > 0 {
		input.ResourceArn = aws.String(_partnercentralchannelResourceArn)
	}
	if len(_partnercentralchannelTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _partnercentralchannelTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a program management account.
func partnercentralchannel_UpdateProgramManagementAccount(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.UpdateProgramManagementAccountInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}
	if len(_partnercentralchannelDisplayName) > 0 {
		input.DisplayName = aws.String(_partnercentralchannelDisplayName)
	}
	if len(_partnercentralchannelRevision) > 0 {
		input.Revision = aws.String(_partnercentralchannelRevision)
	}

	if resp, err := client.UpdateProgramManagementAccount(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the properties of a partner relationship.
func partnercentralchannel_UpdateRelationship(cfg aws.Config, client *partnercentralchannel.Client) {
	input := &partnercentralchannel.UpdateRelationshipInput{
		// Catalog: *string, // Required
		// Identifier: *string, // Required
		// ProgramManagementAccountIdentifier: *string, // Required
	}

	if len(_partnercentralchannelCatalog) > 0 {
		input.Catalog = aws.String(_partnercentralchannelCatalog)
	}
	if len(_partnercentralchannelIdentifier) > 0 {
		input.Identifier = aws.String(_partnercentralchannelIdentifier)
	}
	if len(_partnercentralchannelProgramManagementAccountIdentifier) > 0 {
		input.ProgramManagementAccountIdentifier = aws.String(_partnercentralchannelProgramManagementAccountIdentifier)
	}
	if len(_partnercentralchannelDisplayName) > 0 {
		input.DisplayName = aws.String(_partnercentralchannelDisplayName)
	}
	if len(_partnercentralchannelRequestedSupportPlan) > 0 {
		if err := assignInputField(input, "RequestedSupportPlan", _partnercentralchannelRequestedSupportPlan); err != nil {
			log.Errorf("invalid --requested-support-plan: %s", err.Error())
			return
		}
	}
	if len(_partnercentralchannelRevision) > 0 {
		input.Revision = aws.String(_partnercentralchannelRevision)
	}

	if resp, err := client.UpdateRelationship(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_partnercentralchannelCmd)
	_partnercentralchannelCmd.Flags().SortFlags = false

	_partnercentralchannelCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_partnercentralchannelCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_partnercentralchannelCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelAccountId, "account-id", "", "", "Account ID")
	_partnercentralchannelCmd.Flags().StringSliceVarP(&_partnercentralchannelAccountIds, "account-ids", "", nil, "Account Ids")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelAssociatedAccountId, "associated-account-id", "", "", "Associated Account ID")
	_partnercentralchannelCmd.Flags().StringSliceVarP(&_partnercentralchannelAssociatedAccountIds, "associated-account-ids", "", nil, "Associated Account Ids")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelAssociatedResourceIdentifier, "associated-resource-identifier", "", "", "Associated Resource Identifier")
	_partnercentralchannelCmd.Flags().StringSliceVarP(&_partnercentralchannelAssociatedResourceIdentifiers, "associated-resource-identifiers", "", nil, "Associated Resource Identifiers")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelAssociationType, "association-type", "", "", "Association Type")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelAssociationTypes, "association-types", "", "", "Association Types")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelCatalog, "catalog", "", "", "Catalog")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelClientToken, "client-token", "", "", "Client Token")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelDisplayName, "display-name", "", "", "Display Name")
	_partnercentralchannelCmd.Flags().StringSliceVarP(&_partnercentralchannelDisplayNames, "display-names", "", nil, "Display Names")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelHandshakeType, "handshake-type", "", "", "Handshake Type")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelHandshakeTypeFilters, "handshake-type-filters", "", "", "Handshake Type Filters")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelHandshakeTypeSort, "handshake-type-sort", "", "", "Handshake Type Sort")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelIdentifier, "identifier", "", "", "Identifier")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelMaxResults, "max-results", "", "", "Max Results")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelNextToken, "next-token", "", "", "Next Token")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelParticipantType, "participant-type", "", "", "Participant Type")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelPayload, "payload", "", "", "Payload")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelProgram, "program", "", "", "Program")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelProgramManagementAccountIdentifier, "program-management-account-identifier", "", "", "Program Management Account Identifier")
	_partnercentralchannelCmd.Flags().StringSliceVarP(&_partnercentralchannelProgramManagementAccountIdentifiers, "program-management-account-identifiers", "", nil, "Program Management Account Identifiers")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelPrograms, "programs", "", "", "Programs")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelRequestedSupportPlan, "requested-support-plan", "", "", "Requested Support Plan")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelResaleAccountModel, "resale-account-model", "", "", "Resale Account Model")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelResourceArn, "resource-arn", "", "", "Resource ARN")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelRevision, "revision", "", "", "Revision")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelSector, "sector", "", "", "Sector")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelSort, "sort", "", "", "Sort")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelStatuses, "statuses", "", "", "Statuses")
	_partnercentralchannelCmd.Flags().StringSliceVarP(&_partnercentralchannelTagKeys, "tag-keys", "", nil, "Tag Keys")
	_partnercentralchannelCmd.Flags().StringVarP(&_partnercentralchannelTags, "tags", "", "", "Tags")

	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelAcceptChannelHandshake, "accept-channel-handshake", "", false, "Accept Channel Handshake")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelCancelChannelHandshake, "cancel-channel-handshake", "", false, "Cancel Channel Handshake")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelCreateChannelHandshake, "create-channel-handshake", "", false, "Create Channel Handshake")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelCreateProgramManagementAccount, "create-program-management-account", "", false, "Create Program Management Account")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelCreateRelationship, "create-relationship", "", false, "Create Relationship")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelDeleteProgramManagementAccount, "delete-program-management-account", "", false, "Delete Program Management Account")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelDeleteRelationship, "delete-relationship", "", false, "Delete Relationship")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelGetRelationship, "get-relationship", "", false, "Get Relationship")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelListChannelHandshakes, "list-channel-handshakes", "", false, "List Channel Handshakes")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelListProgramManagementAccounts, "list-program-management-accounts", "", false, "List Program Management Accounts")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelListRelationships, "list-relationships", "", false, "List Relationships")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelRejectChannelHandshake, "reject-channel-handshake", "", false, "Reject Channel Handshake")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelTagResource, "tag-resource", "", false, "Tag Resource")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelUntagResource, "untag-resource", "", false, "Untag Resource")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelUpdateProgramManagementAccount, "update-program-management-account", "", false, "Update Program Management Account")
	_partnercentralchannelCmd.Flags().BoolVarP(&_partnercentralchannelUpdateRelationship, "update-relationship", "", false, "Update Relationship")

}
