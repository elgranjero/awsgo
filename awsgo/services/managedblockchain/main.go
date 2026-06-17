package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/managedblockchain/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"create-accessor", "create-member", "create-network", "create-node", "create-proposal", "delete-accessor", "delete-member", "delete-node", "get-accessor", "get-member", "get-network", "get-node", "get-proposal", "list-accessors", "list-invitations", "list-members", "list-networks", "list-nodes", "list-proposal-votes", "list-proposals", "list-tags-for-resource", "reject-invitation", "tag-resource", "untag-resource", "update-member", "update-node", "vote-on-proposal"},
		OperationSet: map[string]bool{"create-accessor": true, "create-member": true, "create-network": true, "create-node": true, "create-proposal": true, "delete-accessor": true, "delete-member": true, "delete-node": true, "get-accessor": true, "get-member": true, "get-network": true, "get-node": true, "get-proposal": true, "list-accessors": true, "list-invitations": true, "list-members": true, "list-networks": true, "list-nodes": true, "list-proposal-votes": true, "list-proposals": true, "list-tags-for-resource": true, "reject-invitation": true, "tag-resource": true, "untag-resource": true, "update-member": true, "update-node": true, "vote-on-proposal": true},
		OperationInputs: map[string][]string{
			"create-accessor":        {"AccessorType", "ClientRequestToken", "NetworkType", "Tags"},
			"create-member":          {"ClientRequestToken", "InvitationId", "MemberConfiguration", "NetworkId"},
			"create-network":         {"ClientRequestToken", "Description", "Framework", "FrameworkConfiguration", "FrameworkVersion", "MemberConfiguration", "Name", "Tags", "VotingPolicy"},
			"create-node":            {"ClientRequestToken", "MemberId", "NetworkId", "NodeConfiguration", "Tags"},
			"create-proposal":        {"Actions", "ClientRequestToken", "Description", "MemberId", "NetworkId", "Tags"},
			"delete-accessor":        {"AccessorId"},
			"delete-member":          {"MemberId", "NetworkId"},
			"delete-node":            {"MemberId", "NetworkId", "NodeId"},
			"get-accessor":           {"AccessorId"},
			"get-member":             {"MemberId", "NetworkId"},
			"get-network":            {"NetworkId"},
			"get-node":               {"MemberId", "NetworkId", "NodeId"},
			"get-proposal":           {"NetworkId", "ProposalId"},
			"list-accessors":         {"MaxResults", "NetworkType", "NextToken"},
			"list-invitations":       {"MaxResults", "NextToken"},
			"list-members":           {"IsOwned", "MaxResults", "Name", "NetworkId", "NextToken", "Status"},
			"list-networks":          {"Framework", "MaxResults", "Name", "NextToken", "Status"},
			"list-nodes":             {"MaxResults", "MemberId", "NetworkId", "NextToken", "Status"},
			"list-proposal-votes":    {"MaxResults", "NetworkId", "NextToken", "ProposalId"},
			"list-proposals":         {"MaxResults", "NetworkId", "NextToken"},
			"list-tags-for-resource": {"ResourceArn"},
			"reject-invitation":      {"InvitationId"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-member":          {"LogPublishingConfiguration", "MemberId", "NetworkId"},
			"update-node":            {"LogPublishingConfiguration", "MemberId", "NetworkId", "NodeId"},
			"vote-on-proposal":       {"NetworkId", "ProposalId", "Vote", "VoterMemberId"},
		},
		OperationInputTypes: map[string]map[string]string{
			"create-accessor":        {"AccessorType": "types.AccessorType", "ClientRequestToken": "*string", "NetworkType": "types.AccessorNetworkType", "Tags": "map[string]string"},
			"create-member":          {"ClientRequestToken": "*string", "InvitationId": "*string", "MemberConfiguration": "*types.MemberConfiguration", "NetworkId": "*string"},
			"create-network":         {"ClientRequestToken": "*string", "Description": "*string", "Framework": "types.Framework", "FrameworkConfiguration": "*types.NetworkFrameworkConfiguration", "FrameworkVersion": "*string", "MemberConfiguration": "*types.MemberConfiguration", "Name": "*string", "Tags": "map[string]string", "VotingPolicy": "*types.VotingPolicy"},
			"create-node":            {"ClientRequestToken": "*string", "MemberId": "*string", "NetworkId": "*string", "NodeConfiguration": "*types.NodeConfiguration", "Tags": "map[string]string"},
			"create-proposal":        {"Actions": "*types.ProposalActions", "ClientRequestToken": "*string", "Description": "*string", "MemberId": "*string", "NetworkId": "*string", "Tags": "map[string]string"},
			"delete-accessor":        {"AccessorId": "*string"},
			"delete-member":          {"MemberId": "*string", "NetworkId": "*string"},
			"delete-node":            {"MemberId": "*string", "NetworkId": "*string", "NodeId": "*string"},
			"get-accessor":           {"AccessorId": "*string"},
			"get-member":             {"MemberId": "*string", "NetworkId": "*string"},
			"get-network":            {"NetworkId": "*string"},
			"get-node":               {"MemberId": "*string", "NetworkId": "*string", "NodeId": "*string"},
			"get-proposal":           {"NetworkId": "*string", "ProposalId": "*string"},
			"list-accessors":         {"MaxResults": "*int32", "NetworkType": "types.AccessorNetworkType", "NextToken": "*string"},
			"list-invitations":       {"MaxResults": "*int32", "NextToken": "*string"},
			"list-members":           {"IsOwned": "*bool", "MaxResults": "*int32", "Name": "*string", "NetworkId": "*string", "NextToken": "*string", "Status": "types.MemberStatus"},
			"list-networks":          {"Framework": "types.Framework", "MaxResults": "*int32", "Name": "*string", "NextToken": "*string", "Status": "types.NetworkStatus"},
			"list-nodes":             {"MaxResults": "*int32", "MemberId": "*string", "NetworkId": "*string", "NextToken": "*string", "Status": "types.NodeStatus"},
			"list-proposal-votes":    {"MaxResults": "*int32", "NetworkId": "*string", "NextToken": "*string", "ProposalId": "*string"},
			"list-proposals":         {"MaxResults": "*int32", "NetworkId": "*string", "NextToken": "*string"},
			"list-tags-for-resource": {"ResourceArn": "*string"},
			"reject-invitation":      {"InvitationId": "*string"},
			"tag-resource":           {"ResourceArn": "*string", "Tags": "map[string]string"},
			"untag-resource":         {"ResourceArn": "*string", "TagKeys": "[]string"},
			"update-member":          {"LogPublishingConfiguration": "*types.MemberLogPublishingConfiguration", "MemberId": "*string", "NetworkId": "*string"},
			"update-node":            {"LogPublishingConfiguration": "*types.NodeLogPublishingConfiguration", "MemberId": "*string", "NetworkId": "*string", "NodeId": "*string"},
			"vote-on-proposal":       {"NetworkId": "*string", "ProposalId": "*string", "Vote": "types.VoteValue", "VoterMemberId": "*string"},
		},
		OperationInputRequired: map[string][]string{
			"create-accessor":        {"AccessorType", "ClientRequestToken"},
			"create-member":          {"ClientRequestToken", "InvitationId", "MemberConfiguration", "NetworkId"},
			"create-network":         {"ClientRequestToken", "Framework", "FrameworkVersion", "MemberConfiguration", "Name", "VotingPolicy"},
			"create-node":            {"ClientRequestToken", "NetworkId", "NodeConfiguration"},
			"create-proposal":        {"Actions", "ClientRequestToken", "MemberId", "NetworkId"},
			"delete-accessor":        {"AccessorId"},
			"delete-member":          {"MemberId", "NetworkId"},
			"delete-node":            {"NetworkId", "NodeId"},
			"get-accessor":           {"AccessorId"},
			"get-member":             {"MemberId", "NetworkId"},
			"get-network":            {"NetworkId"},
			"get-node":               {"NetworkId", "NodeId"},
			"get-proposal":           {"NetworkId", "ProposalId"},
			"list-accessors":         {},
			"list-invitations":       {},
			"list-members":           {"NetworkId"},
			"list-networks":          {},
			"list-nodes":             {"NetworkId"},
			"list-proposal-votes":    {"NetworkId", "ProposalId"},
			"list-proposals":         {"NetworkId"},
			"list-tags-for-resource": {"ResourceArn"},
			"reject-invitation":      {"InvitationId"},
			"tag-resource":           {"ResourceArn", "Tags"},
			"untag-resource":         {"ResourceArn", "TagKeys"},
			"update-member":          {"MemberId", "NetworkId"},
			"update-node":            {"NetworkId", "NodeId"},
			"vote-on-proposal":       {"NetworkId", "ProposalId", "Vote", "VoterMemberId"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("managedblockchain", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
