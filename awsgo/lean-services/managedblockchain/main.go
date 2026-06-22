package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/managedblockchain"
)

var fields_create_accessor = []leanruntime.Field{
	{Name: "AccessorType", Flag: "accessor-type", Type: "types.AccessorType", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "NetworkType", Flag: "network-type", Type: "types.AccessorNetworkType", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_member = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
	{Name: "MemberConfiguration", Flag: "member-configuration", Type: "*types.MemberConfiguration", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_create_network = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Framework", Flag: "framework", Type: "types.Framework", Required: true},
	{Name: "FrameworkConfiguration", Flag: "framework-configuration", Type: "*types.NetworkFrameworkConfiguration", Required: false},
	{Name: "FrameworkVersion", Flag: "framework-version", Type: "*string", Required: true},
	{Name: "MemberConfiguration", Flag: "member-configuration", Type: "*types.MemberConfiguration", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "VotingPolicy", Flag: "voting-policy", Type: "*types.VotingPolicy", Required: true},
}

var fields_create_node = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NodeConfiguration", Flag: "node-configuration", Type: "*types.NodeConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_proposal = []leanruntime.Field{
	{Name: "Actions", Flag: "actions", Type: "*types.ProposalActions", Required: true},
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_accessor = []leanruntime.Field{
	{Name: "AccessorId", Flag: "accessor-id", Type: "*string", Required: true},
}

var fields_delete_member = []leanruntime.Field{
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_delete_node = []leanruntime.Field{
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_get_accessor = []leanruntime.Field{
	{Name: "AccessorId", Flag: "accessor-id", Type: "*string", Required: true},
}

var fields_get_member = []leanruntime.Field{
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_network = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_get_node = []leanruntime.Field{
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_get_proposal = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "ProposalId", Flag: "proposal-id", Type: "*string", Required: true},
}

var fields_list_accessors = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkType", Flag: "network-type", Type: "types.AccessorNetworkType", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_invitations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_members = []leanruntime.Field{
	{Name: "IsOwned", Flag: "is-owned", Type: "*bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.MemberStatus", Required: false},
}

var fields_list_networks = []leanruntime.Field{
	{Name: "Framework", Flag: "framework", Type: "types.Framework", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.NetworkStatus", Required: false},
}

var fields_list_nodes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.NodeStatus", Required: false},
}

var fields_list_proposal_votes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProposalId", Flag: "proposal-id", Type: "*string", Required: true},
}

var fields_list_proposals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_reject_invitation = []leanruntime.Field{
	{Name: "InvitationId", Flag: "invitation-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_member = []leanruntime.Field{
	{Name: "LogPublishingConfiguration", Flag: "log-publishing-configuration", Type: "*types.MemberLogPublishingConfiguration", Required: false},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: true},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
}

var fields_update_node = []leanruntime.Field{
	{Name: "LogPublishingConfiguration", Flag: "log-publishing-configuration", Type: "*types.NodeLogPublishingConfiguration", Required: false},
	{Name: "MemberId", Flag: "member-id", Type: "*string", Required: false},
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "NodeId", Flag: "node-id", Type: "*string", Required: true},
}

var fields_vote_on_proposal = []leanruntime.Field{
	{Name: "NetworkId", Flag: "network-id", Type: "*string", Required: true},
	{Name: "ProposalId", Flag: "proposal-id", Type: "*string", Required: true},
	{Name: "Vote", Flag: "vote", Type: "types.VoteValue", Required: true},
	{Name: "VoterMemberId", Flag: "voter-member-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-accessor": {
			Name:   "create-accessor",
			Fields: fields_create_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccessor(ctx, input)
			},
		},
		"create-member": {
			Name:   "create-member",
			Fields: fields_create_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMember(ctx, input)
			},
		},
		"create-network": {
			Name:   "create-network",
			Fields: fields_create_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNetwork(ctx, input)
			},
		},
		"create-node": {
			Name:   "create-node",
			Fields: fields_create_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateNode(ctx, input)
			},
		},
		"create-proposal": {
			Name:   "create-proposal",
			Fields: fields_create_proposal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProposalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_proposal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProposal(ctx, input)
			},
		},
		"delete-accessor": {
			Name:   "delete-accessor",
			Fields: fields_delete_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAccessor(ctx, input)
			},
		},
		"delete-member": {
			Name:   "delete-member",
			Fields: fields_delete_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMember(ctx, input)
			},
		},
		"delete-node": {
			Name:   "delete-node",
			Fields: fields_delete_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteNode(ctx, input)
			},
		},
		"get-accessor": {
			Name:   "get-accessor",
			Fields: fields_get_accessor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAccessorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_accessor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAccessor(ctx, input)
			},
		},
		"get-member": {
			Name:   "get-member",
			Fields: fields_get_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMember(ctx, input)
			},
		},
		"get-network": {
			Name:   "get-network",
			Fields: fields_get_network,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNetworkInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_network, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNetwork(ctx, input)
			},
		},
		"get-node": {
			Name:   "get-node",
			Fields: fields_get_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetNode(ctx, input)
			},
		},
		"get-proposal": {
			Name:   "get-proposal",
			Fields: fields_get_proposal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProposalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_proposal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProposal(ctx, input)
			},
		},
		"list-accessors": {
			Name:   "list-accessors",
			Fields: fields_list_accessors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccessorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accessors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccessors(ctx, input)
				}
				var results []*svc.ListAccessorsOutput
				p := svc.NewListAccessorsPaginator(client, input)
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
		"list-invitations": {
			Name:   "list-invitations",
			Fields: fields_list_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListInvitations(ctx, input)
				}
				var results []*svc.ListInvitationsOutput
				p := svc.NewListInvitationsPaginator(client, input)
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
		"list-members": {
			Name:   "list-members",
			Fields: fields_list_members,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMembersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_members, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMembers(ctx, input)
				}
				var results []*svc.ListMembersOutput
				p := svc.NewListMembersPaginator(client, input)
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
		"list-networks": {
			Name:   "list-networks",
			Fields: fields_list_networks,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNetworksInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_networks, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNetworks(ctx, input)
				}
				var results []*svc.ListNetworksOutput
				p := svc.NewListNetworksPaginator(client, input)
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
		"list-nodes": {
			Name:   "list-nodes",
			Fields: fields_list_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListNodes(ctx, input)
				}
				var results []*svc.ListNodesOutput
				p := svc.NewListNodesPaginator(client, input)
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
		"list-proposal-votes": {
			Name:   "list-proposal-votes",
			Fields: fields_list_proposal_votes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProposalVotesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_proposal_votes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProposalVotes(ctx, input)
				}
				var results []*svc.ListProposalVotesOutput
				p := svc.NewListProposalVotesPaginator(client, input)
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
		"list-proposals": {
			Name:   "list-proposals",
			Fields: fields_list_proposals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProposalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_proposals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProposals(ctx, input)
				}
				var results []*svc.ListProposalsOutput
				p := svc.NewListProposalsPaginator(client, input)
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
		"reject-invitation": {
			Name:   "reject-invitation",
			Fields: fields_reject_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectInvitation(ctx, input)
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
		"update-member": {
			Name:   "update-member",
			Fields: fields_update_member,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMemberInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_member, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMember(ctx, input)
			},
		},
		"update-node": {
			Name:   "update-node",
			Fields: fields_update_node,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateNodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_node, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateNode(ctx, input)
			},
		},
		"vote-on-proposal": {
			Name:   "vote-on-proposal",
			Fields: fields_vote_on_proposal,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VoteOnProposalInput{}
				if _, err := leanruntime.ApplyInput(input, fields_vote_on_proposal, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VoteOnProposal(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("managedblockchain", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
