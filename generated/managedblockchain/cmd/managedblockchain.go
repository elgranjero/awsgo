package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/managedblockchain"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// managedblockchainCmd represents the managedblockchain command
var _managedblockchainCmd = &cobra.Command{
	Use:   "managedblockchain",
	Short: "AWS managedblockchain CLI",
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
		client := managedblockchain.NewFromConfig(cfg)
		if _managedblockchainCreateAccessor {
			managedblockchain_CreateAccessor(cfg, client)
			return
		}
		if _managedblockchainCreateMember {
			managedblockchain_CreateMember(cfg, client)
			return
		}
		if _managedblockchainCreateNetwork {
			managedblockchain_CreateNetwork(cfg, client)
			return
		}
		if _managedblockchainCreateNode {
			managedblockchain_CreateNode(cfg, client)
			return
		}
		if _managedblockchainCreateProposal {
			managedblockchain_CreateProposal(cfg, client)
			return
		}
		if _managedblockchainDeleteAccessor {
			managedblockchain_DeleteAccessor(cfg, client)
			return
		}
		if _managedblockchainDeleteMember {
			managedblockchain_DeleteMember(cfg, client)
			return
		}
		if _managedblockchainDeleteNode {
			managedblockchain_DeleteNode(cfg, client)
			return
		}
		if _managedblockchainGetAccessor {
			managedblockchain_GetAccessor(cfg, client)
			return
		}
		if _managedblockchainGetMember {
			managedblockchain_GetMember(cfg, client)
			return
		}
		if _managedblockchainGetNetwork {
			managedblockchain_GetNetwork(cfg, client)
			return
		}
		if _managedblockchainGetNode {
			managedblockchain_GetNode(cfg, client)
			return
		}
		if _managedblockchainGetProposal {
			managedblockchain_GetProposal(cfg, client)
			return
		}
		if _managedblockchainListAccessors {
			managedblockchain_ListAccessors(cfg, client)
			return
		}
		if _managedblockchainListInvitations {
			managedblockchain_ListInvitations(cfg, client)
			return
		}
		if _managedblockchainListMembers {
			managedblockchain_ListMembers(cfg, client)
			return
		}
		if _managedblockchainListNetworks {
			managedblockchain_ListNetworks(cfg, client)
			return
		}
		if _managedblockchainListNodes {
			managedblockchain_ListNodes(cfg, client)
			return
		}
		if _managedblockchainListProposalVotes {
			managedblockchain_ListProposalVotes(cfg, client)
			return
		}
		if _managedblockchainListProposals {
			managedblockchain_ListProposals(cfg, client)
			return
		}
		if _managedblockchainListTagsForResource {
			managedblockchain_ListTagsForResource(cfg, client)
			return
		}
		if _managedblockchainRejectInvitation {
			managedblockchain_RejectInvitation(cfg, client)
			return
		}
		if _managedblockchainTagResource {
			managedblockchain_TagResource(cfg, client)
			return
		}
		if _managedblockchainUntagResource {
			managedblockchain_UntagResource(cfg, client)
			return
		}
		if _managedblockchainUpdateMember {
			managedblockchain_UpdateMember(cfg, client)
			return
		}
		if _managedblockchainUpdateNode {
			managedblockchain_UpdateNode(cfg, client)
			return
		}
		if _managedblockchainVoteOnProposal {
			managedblockchain_VoteOnProposal(cfg, client)
			return
		}

	},
}

var (
	_managedblockchainCreateAccessor      bool
	_managedblockchainCreateMember        bool
	_managedblockchainCreateNetwork       bool
	_managedblockchainCreateNode          bool
	_managedblockchainCreateProposal      bool
	_managedblockchainDeleteAccessor      bool
	_managedblockchainDeleteMember        bool
	_managedblockchainDeleteNode          bool
	_managedblockchainGetAccessor         bool
	_managedblockchainGetMember           bool
	_managedblockchainGetNetwork          bool
	_managedblockchainGetNode             bool
	_managedblockchainGetProposal         bool
	_managedblockchainListAccessors       bool
	_managedblockchainListInvitations     bool
	_managedblockchainListMembers         bool
	_managedblockchainListNetworks        bool
	_managedblockchainListNodes           bool
	_managedblockchainListProposalVotes   bool
	_managedblockchainListProposals       bool
	_managedblockchainListTagsForResource bool
	_managedblockchainRejectInvitation    bool
	_managedblockchainTagResource         bool
	_managedblockchainUntagResource       bool
	_managedblockchainUpdateMember        bool
	_managedblockchainUpdateNode          bool
	_managedblockchainVoteOnProposal      bool

	_managedblockchainAccessorId                 string
	_managedblockchainAccessorType               string
	_managedblockchainActions                    string
	_managedblockchainClientRequestToken         string
	_managedblockchainDescription                string
	_managedblockchainFramework                  string
	_managedblockchainFrameworkConfiguration     string
	_managedblockchainFrameworkVersion           string
	_managedblockchainInvitationId               string
	_managedblockchainIsOwned                    string
	_managedblockchainLogPublishingConfiguration string
	_managedblockchainMaxResults                 string
	_managedblockchainMemberConfiguration        string
	_managedblockchainMemberId                   string
	_managedblockchainName                       string
	_managedblockchainNetworkId                  string
	_managedblockchainNetworkType                string
	_managedblockchainNextToken                  string
	_managedblockchainNodeConfiguration          string
	_managedblockchainNodeId                     string
	_managedblockchainProposalId                 string
	_managedblockchainResourceArn                string
	_managedblockchainStatus                     string
	_managedblockchainTagKeys                    []string
	_managedblockchainTags                       string
	_managedblockchainVote                       string
	_managedblockchainVoterMemberId              string
	_managedblockchainVotingPolicy               string
)

// Creates a new accessor for use with Amazon Managed Blockchain service that
// supports token based access. The accessor contains information required for
// token based access.
func managedblockchain_CreateAccessor(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.CreateAccessorInput{
		// AccessorType: types.AccessorType, // Required
		// ClientRequestToken: *string, // Required
	}

	if len(_managedblockchainAccessorType) > 0 {
		if err := assignInputField(input, "AccessorType", _managedblockchainAccessorType); err != nil {
			log.Errorf("invalid --accessor-type: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_managedblockchainClientRequestToken)
	}
	if len(_managedblockchainNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _managedblockchainNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainTags) > 0 {
		if err := assignInputField(input, "Tags", _managedblockchainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a member within a Managed Blockchain network.
// Applies only to Hyperledger Fabric.
func managedblockchain_CreateMember(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.CreateMemberInput{
		// ClientRequestToken: *string, // Required
		// InvitationId: *string, // Required
		// MemberConfiguration: *types.MemberConfiguration, // Required
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_managedblockchainClientRequestToken)
	}
	if len(_managedblockchainInvitationId) > 0 {
		input.InvitationId = aws.String(_managedblockchainInvitationId)
	}
	if len(_managedblockchainMemberConfiguration) > 0 {
		if err := assignInputField(input, "MemberConfiguration", _managedblockchainMemberConfiguration); err != nil {
			log.Errorf("invalid --member-configuration: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}

	if resp, err := client.CreateMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new blockchain network using Amazon Managed Blockchain.
// Applies only to Hyperledger Fabric.
func managedblockchain_CreateNetwork(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.CreateNetworkInput{
		// ClientRequestToken: *string, // Required
		// Framework: types.Framework, // Required
		// FrameworkVersion: *string, // Required
		// MemberConfiguration: *types.MemberConfiguration, // Required
		// Name: *string, // Required
		// VotingPolicy: *types.VotingPolicy, // Required
	}

	if len(_managedblockchainClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_managedblockchainClientRequestToken)
	}
	if len(_managedblockchainFramework) > 0 {
		if err := assignInputField(input, "Framework", _managedblockchainFramework); err != nil {
			log.Errorf("invalid --framework: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainFrameworkVersion) > 0 {
		input.FrameworkVersion = aws.String(_managedblockchainFrameworkVersion)
	}
	if len(_managedblockchainMemberConfiguration) > 0 {
		if err := assignInputField(input, "MemberConfiguration", _managedblockchainMemberConfiguration); err != nil {
			log.Errorf("invalid --member-configuration: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainName) > 0 {
		input.Name = aws.String(_managedblockchainName)
	}
	if len(_managedblockchainVotingPolicy) > 0 {
		if err := assignInputField(input, "VotingPolicy", _managedblockchainVotingPolicy); err != nil {
			log.Errorf("invalid --voting-policy: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainDescription) > 0 {
		input.Description = aws.String(_managedblockchainDescription)
	}
	if len(_managedblockchainFrameworkConfiguration) > 0 {
		if err := assignInputField(input, "FrameworkConfiguration", _managedblockchainFrameworkConfiguration); err != nil {
			log.Errorf("invalid --framework-configuration: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainTags) > 0 {
		if err := assignInputField(input, "Tags", _managedblockchainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a node on the specified blockchain network.
// Applies to Hyperledger Fabric and Ethereum.
func managedblockchain_CreateNode(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.CreateNodeInput{
		// ClientRequestToken: *string, // Required
		// NetworkId: *string, // Required
		// NodeConfiguration: *types.NodeConfiguration, // Required
	}

	if len(_managedblockchainClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_managedblockchainClientRequestToken)
	}
	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainNodeConfiguration) > 0 {
		if err := assignInputField(input, "NodeConfiguration", _managedblockchainNodeConfiguration); err != nil {
			log.Errorf("invalid --node-configuration: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}
	if len(_managedblockchainTags) > 0 {
		if err := assignInputField(input, "Tags", _managedblockchainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a proposal for a change to the network that other members of the
// network can vote on, for example, a proposal to add a new member to the network.
// Any member can create a proposal.
//
// Applies only to Hyperledger Fabric.
func managedblockchain_CreateProposal(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.CreateProposalInput{
		// Actions: *types.ProposalActions, // Required
		// ClientRequestToken: *string, // Required
		// MemberId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainActions) > 0 {
		if err := assignInputField(input, "Actions", _managedblockchainActions); err != nil {
			log.Errorf("invalid --actions: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_managedblockchainClientRequestToken)
	}
	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}
	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainDescription) > 0 {
		input.Description = aws.String(_managedblockchainDescription)
	}
	if len(_managedblockchainTags) > 0 {
		if err := assignInputField(input, "Tags", _managedblockchainTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateProposal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an accessor that your Amazon Web Services account owns. An accessor
// object is a container that has the information required for token based access
// to your Ethereum nodes including, the BILLING_TOKEN . After an accessor is
// deleted, the status of the accessor changes from AVAILABLE to PENDING_DELETION .
// An accessor in the PENDING_DELETION state can’t be used for new WebSocket
// requests or HTTP requests. However, WebSocket connections that were initiated
// while the accessor was in the AVAILABLE state remain open until they expire (up
// to 2 hours).
func managedblockchain_DeleteAccessor(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.DeleteAccessorInput{
		// AccessorId: *string, // Required
	}

	if len(_managedblockchainAccessorId) > 0 {
		input.AccessorId = aws.String(_managedblockchainAccessorId)
	}

	if resp, err := client.DeleteAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a member. Deleting a member removes the member and all associated
// resources from the network. DeleteMember can only be called for a specified
// MemberId if the principal performing the action is associated with the Amazon
// Web Services account that owns the member. In all other cases, the DeleteMember
// action is carried out as the result of an approved proposal to remove a member.
// If MemberId is the last member in a network specified by the last Amazon Web
// Services account, the network is deleted also.
//
// Applies only to Hyperledger Fabric.
func managedblockchain_DeleteMember(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.DeleteMemberInput{
		// MemberId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}
	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}

	if resp, err := client.DeleteMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a node that your Amazon Web Services account owns. All data on the node
// is lost and cannot be recovered.
//
// Applies to Hyperledger Fabric and Ethereum.
func managedblockchain_DeleteNode(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.DeleteNodeInput{
		// NetworkId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainNodeId) > 0 {
		input.NodeId = aws.String(_managedblockchainNodeId)
	}
	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}

	if resp, err := client.DeleteNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about an accessor. An accessor object is a
// container that has the information required for token based access to your
// Ethereum nodes.
func managedblockchain_GetAccessor(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.GetAccessorInput{
		// AccessorId: *string, // Required
	}

	if len(_managedblockchainAccessorId) > 0 {
		input.AccessorId = aws.String(_managedblockchainAccessorId)
	}

	if resp, err := client.GetAccessor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a member.
// Applies only to Hyperledger Fabric.
func managedblockchain_GetMember(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.GetMemberInput{
		// MemberId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}
	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}

	if resp, err := client.GetMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a network.
// Applies to Hyperledger Fabric and Ethereum.
func managedblockchain_GetNetwork(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.GetNetworkInput{
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}

	if resp, err := client.GetNetwork(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a node.
// Applies to Hyperledger Fabric and Ethereum.
func managedblockchain_GetNode(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.GetNodeInput{
		// NetworkId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainNodeId) > 0 {
		input.NodeId = aws.String(_managedblockchainNodeId)
	}
	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}

	if resp, err := client.GetNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns detailed information about a proposal.
// Applies only to Hyperledger Fabric.
func managedblockchain_GetProposal(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.GetProposalInput{
		// NetworkId: *string, // Required
		// ProposalId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainProposalId) > 0 {
		input.ProposalId = aws.String(_managedblockchainProposalId)
	}

	if resp, err := client.GetProposal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of the accessors and their properties. Accessor objects are
// containers that have the information required for token based access to your
// Ethereum nodes.
func managedblockchain_ListAccessors(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListAccessorsInput{}

	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainNetworkType) > 0 {
		if err := assignInputField(input, "NetworkType", _managedblockchainNetworkType); err != nil {
			log.Errorf("invalid --network-type: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchain.ListAccessorsOutput
	p := managedblockchain.NewListAccessorsPaginator(client, input)
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

// Returns a list of all invitations for the current Amazon Web Services account.
// Applies only to Hyperledger Fabric.
func managedblockchain_ListInvitations(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListInvitationsInput{}

	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
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

	var results []*managedblockchain.ListInvitationsOutput
	p := managedblockchain.NewListInvitationsPaginator(client, input)
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

// Returns a list of the members in a network and properties of their
// configurations.
//
// Applies only to Hyperledger Fabric.
func managedblockchain_ListMembers(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListMembersInput{
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainIsOwned) > 0 {
		if err := assignInputField(input, "IsOwned", _managedblockchainIsOwned); err != nil {
			log.Errorf("invalid --is-owned: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainName) > 0 {
		input.Name = aws.String(_managedblockchainName)
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
	}
	if len(_managedblockchainStatus) > 0 {
		if err := assignInputField(input, "Status", _managedblockchainStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
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

	var results []*managedblockchain.ListMembersOutput
	p := managedblockchain.NewListMembersPaginator(client, input)
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

// Returns information about the networks in which the current Amazon Web Services
// account participates.
//
// Applies to Hyperledger Fabric and Ethereum.
func managedblockchain_ListNetworks(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListNetworksInput{}

	if len(_managedblockchainFramework) > 0 {
		if err := assignInputField(input, "Framework", _managedblockchainFramework); err != nil {
			log.Errorf("invalid --framework: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainName) > 0 {
		input.Name = aws.String(_managedblockchainName)
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
	}
	if len(_managedblockchainStatus) > 0 {
		if err := assignInputField(input, "Status", _managedblockchainStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNetworks(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchain.ListNetworksOutput
	p := managedblockchain.NewListNetworksPaginator(client, input)
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

// Returns information about the nodes within a network.
// Applies to Hyperledger Fabric and Ethereum.
func managedblockchain_ListNodes(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListNodesInput{
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
	}
	if len(_managedblockchainStatus) > 0 {
		if err := assignInputField(input, "Status", _managedblockchainStatus); err != nil {
			log.Errorf("invalid --status: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListNodes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchain.ListNodesOutput
	p := managedblockchain.NewListNodesPaginator(client, input)
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

// Returns the list of votes for a specified proposal, including the value of each
// vote and the unique identifier of the member that cast the vote.
//
// Applies only to Hyperledger Fabric.
func managedblockchain_ListProposalVotes(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListProposalVotesInput{
		// NetworkId: *string, // Required
		// ProposalId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainProposalId) > 0 {
		input.ProposalId = aws.String(_managedblockchainProposalId)
	}
	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProposalVotes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchain.ListProposalVotesOutput
	p := managedblockchain.NewListProposalVotesPaginator(client, input)
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

// Returns a list of proposals for the network.
// Applies only to Hyperledger Fabric.
func managedblockchain_ListProposals(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListProposalsInput{
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _managedblockchainMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainNextToken) > 0 {
		input.NextToken = aws.String(_managedblockchainNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListProposals(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*managedblockchain.ListProposalsOutput
	p := managedblockchain.NewListProposalsPaginator(client, input)
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

// Returns a list of tags for the specified resource. Each tag consists of a key
// and optional value.
//
// For more information about tags, see [Tagging Resources] in the Amazon Managed Blockchain Ethereum
// Developer Guide, or [Tagging Resources]in the Amazon Managed Blockchain Hyperledger Fabric
// Developer Guide.
//
// [Tagging Resources]: https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html
func managedblockchain_ListTagsForResource(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_managedblockchainResourceArn) > 0 {
		input.ResourceArn = aws.String(_managedblockchainResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Rejects an invitation to join a network. This action can be called by a
// principal in an Amazon Web Services account that has received an invitation to
// create a member and join a network.
//
// Applies only to Hyperledger Fabric.
func managedblockchain_RejectInvitation(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.RejectInvitationInput{
		// InvitationId: *string, // Required
	}

	if len(_managedblockchainInvitationId) > 0 {
		input.InvitationId = aws.String(_managedblockchainInvitationId)
	}

	if resp, err := client.RejectInvitation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds or overwrites the specified tags for the specified Amazon Managed
// Blockchain resource. Each tag consists of a key and optional value.
//
// When you specify a tag key that already exists, the tag value is overwritten
// with the new value. Use UntagResource to remove tag keys.
//
// A resource can have up to 50 tags. If you try to create more than 50 tags for a
// resource, your request fails and returns an error.
//
// For more information about tags, see [Tagging Resources] in the Amazon Managed Blockchain Ethereum
// Developer Guide, or [Tagging Resources]in the Amazon Managed Blockchain Hyperledger Fabric
// Developer Guide.
//
// [Tagging Resources]: https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html
func managedblockchain_TagResource(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_managedblockchainResourceArn) > 0 {
		input.ResourceArn = aws.String(_managedblockchainResourceArn)
	}
	if len(_managedblockchainTags) > 0 {
		if err := assignInputField(input, "Tags", _managedblockchainTags); err != nil {
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

// Removes the specified tags from the Amazon Managed Blockchain resource.
// For more information about tags, see [Tagging Resources] in the Amazon Managed Blockchain Ethereum
// Developer Guide, or [Tagging Resources]in the Amazon Managed Blockchain Hyperledger Fabric
// Developer Guide.
//
// [Tagging Resources]: https://docs.aws.amazon.com/managed-blockchain/latest/hyperledger-fabric-dev/tagging-resources.html
func managedblockchain_UntagResource(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_managedblockchainResourceArn) > 0 {
		input.ResourceArn = aws.String(_managedblockchainResourceArn)
	}
	if len(_managedblockchainTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _managedblockchainTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a member configuration with new parameters.
// Applies only to Hyperledger Fabric.
func managedblockchain_UpdateMember(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.UpdateMemberInput{
		// MemberId: *string, // Required
		// NetworkId: *string, // Required
	}

	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}
	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainLogPublishingConfiguration) > 0 {
		if err := assignInputField(input, "LogPublishingConfiguration", _managedblockchainLogPublishingConfiguration); err != nil {
			log.Errorf("invalid --log-publishing-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMember(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates a node configuration with new parameters.
// Applies only to Hyperledger Fabric.
func managedblockchain_UpdateNode(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.UpdateNodeInput{
		// NetworkId: *string, // Required
		// NodeId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainNodeId) > 0 {
		input.NodeId = aws.String(_managedblockchainNodeId)
	}
	if len(_managedblockchainLogPublishingConfiguration) > 0 {
		if err := assignInputField(input, "LogPublishingConfiguration", _managedblockchainLogPublishingConfiguration); err != nil {
			log.Errorf("invalid --log-publishing-configuration: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainMemberId) > 0 {
		input.MemberId = aws.String(_managedblockchainMemberId)
	}

	if resp, err := client.UpdateNode(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Casts a vote for a specified ProposalId on behalf of a member. The member to
// vote as, specified by VoterMemberId , must be in the same Amazon Web Services
// account as the principal that calls the action.
//
// Applies only to Hyperledger Fabric.
func managedblockchain_VoteOnProposal(cfg aws.Config, client *managedblockchain.Client) {
	input := &managedblockchain.VoteOnProposalInput{
		// NetworkId: *string, // Required
		// ProposalId: *string, // Required
		// Vote: types.VoteValue, // Required
		// VoterMemberId: *string, // Required
	}

	if len(_managedblockchainNetworkId) > 0 {
		input.NetworkId = aws.String(_managedblockchainNetworkId)
	}
	if len(_managedblockchainProposalId) > 0 {
		input.ProposalId = aws.String(_managedblockchainProposalId)
	}
	if len(_managedblockchainVote) > 0 {
		if err := assignInputField(input, "Vote", _managedblockchainVote); err != nil {
			log.Errorf("invalid --vote: %s", err.Error())
			return
		}
	}
	if len(_managedblockchainVoterMemberId) > 0 {
		input.VoterMemberId = aws.String(_managedblockchainVoterMemberId)
	}

	if resp, err := client.VoteOnProposal(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_managedblockchainCmd)
	_managedblockchainCmd.Flags().SortFlags = false

	_managedblockchainCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_managedblockchainCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_managedblockchainCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainAccessorId, "accessor-id", "", "", "Accessor ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainAccessorType, "accessor-type", "", "", "Accessor Type")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainActions, "actions", "", "", "Actions")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainDescription, "description", "", "", "Description")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainFramework, "framework", "", "", "Framework")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainFrameworkConfiguration, "framework-configuration", "", "", "Framework Configuration")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainFrameworkVersion, "framework-version", "", "", "Framework Version")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainInvitationId, "invitation-id", "", "", "Invitation ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainIsOwned, "is-owned", "", "", "Is Owned")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainLogPublishingConfiguration, "log-publishing-configuration", "", "", "Log Publishing Configuration")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainMaxResults, "max-results", "", "", "Max Results")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainMemberConfiguration, "member-configuration", "", "", "Member Configuration")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainMemberId, "member-id", "", "", "Member ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainName, "name", "", "", "Name")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainNetworkId, "network-id", "", "", "Network ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainNetworkType, "network-type", "", "", "Network Type")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainNextToken, "next-token", "", "", "Next Token")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainNodeConfiguration, "node-configuration", "", "", "Node Configuration")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainNodeId, "node-id", "", "", "Node ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainProposalId, "proposal-id", "", "", "Proposal ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainResourceArn, "resource-arn", "", "", "Resource ARN")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainStatus, "status", "", "", "Status")
	_managedblockchainCmd.Flags().StringSliceVarP(&_managedblockchainTagKeys, "tag-keys", "", nil, "Tag Keys")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainTags, "tags", "", "", "Tags")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainVote, "vote", "", "", "Vote")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainVoterMemberId, "voter-member-id", "", "", "Voter Member ID")
	_managedblockchainCmd.Flags().StringVarP(&_managedblockchainVotingPolicy, "voting-policy", "", "", "Voting Policy")

	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainCreateAccessor, "create-accessor", "", false, "Create Accessor")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainCreateMember, "create-member", "", false, "Create Member")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainCreateNetwork, "create-network", "", false, "Create Network")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainCreateNode, "create-node", "", false, "Create Node")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainCreateProposal, "create-proposal", "", false, "Create Proposal")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainDeleteAccessor, "delete-accessor", "", false, "Delete Accessor")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainDeleteMember, "delete-member", "", false, "Delete Member")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainDeleteNode, "delete-node", "", false, "Delete Node")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainGetAccessor, "get-accessor", "", false, "Get Accessor")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainGetMember, "get-member", "", false, "Get Member")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainGetNetwork, "get-network", "", false, "Get Network")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainGetNode, "get-node", "", false, "Get Node")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainGetProposal, "get-proposal", "", false, "Get Proposal")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListAccessors, "list-accessors", "", false, "List Accessors")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListInvitations, "list-invitations", "", false, "List Invitations")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListMembers, "list-members", "", false, "List Members")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListNetworks, "list-networks", "", false, "List Networks")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListNodes, "list-nodes", "", false, "List Nodes")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListProposalVotes, "list-proposal-votes", "", false, "List Proposal Votes")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListProposals, "list-proposals", "", false, "List Proposals")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainRejectInvitation, "reject-invitation", "", false, "Reject Invitation")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainTagResource, "tag-resource", "", false, "Tag Resource")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainUntagResource, "untag-resource", "", false, "Untag Resource")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainUpdateMember, "update-member", "", false, "Update Member")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainUpdateNode, "update-node", "", false, "Update Node")
	_managedblockchainCmd.Flags().BoolVarP(&_managedblockchainVoteOnProposal, "vote-on-proposal", "", false, "Vote On Proposal")

}
