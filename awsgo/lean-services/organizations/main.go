package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/organizations"
)

var fields_accept_handshake = []leanruntime.Field{
	{Name: "HandshakeId", Flag: "handshake-id", Type: "*string", Required: true},
}

var fields_attach_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_cancel_handshake = []leanruntime.Field{
	{Name: "HandshakeId", Flag: "handshake-id", Type: "*string", Required: true},
}

var fields_close_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_create_account = []leanruntime.Field{
	{Name: "AccountName", Flag: "account-name", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "IamUserAccessToBilling", Flag: "iam-user-access-to-billing", Type: "types.IAMUserAccessToBilling", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_gov_cloud_account = []leanruntime.Field{
	{Name: "AccountName", Flag: "account-name", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "IamUserAccessToBilling", Flag: "iam-user-access-to-billing", Type: "types.IAMUserAccessToBilling", Required: false},
	{Name: "RoleName", Flag: "role-name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_organization = []leanruntime.Field{
	{Name: "FeatureSet", Flag: "feature-set", Type: "types.OrganizationFeatureSet", Required: false},
}

var fields_create_organizational_unit = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentId", Flag: "parent-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_policy = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.PolicyType", Required: true},
}

var fields_decline_handshake = []leanruntime.Field{
	{Name: "HandshakeId", Flag: "handshake-id", Type: "*string", Required: true},
}

var fields_delete_organization = []leanruntime.Field{}

var fields_delete_organizational_unit = []leanruntime.Field{
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: true},
}

var fields_delete_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_delete_resource_policy = []leanruntime.Field{}

var fields_deregister_delegated_administrator = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: true},
}

var fields_describe_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_describe_create_account_status = []leanruntime.Field{
	{Name: "CreateAccountRequestId", Flag: "create-account-request-id", Type: "*string", Required: true},
}

var fields_describe_effective_policy = []leanruntime.Field{
	{Name: "PolicyType", Flag: "policy-type", Type: "types.EffectivePolicyType", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: false},
}

var fields_describe_handshake = []leanruntime.Field{
	{Name: "HandshakeId", Flag: "handshake-id", Type: "*string", Required: true},
}

var fields_describe_organization = []leanruntime.Field{}

var fields_describe_organizational_unit = []leanruntime.Field{
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: true},
}

var fields_describe_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_describe_resource_policy = []leanruntime.Field{}

var fields_describe_responsibility_transfer = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_detach_policy = []leanruntime.Field{
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_disable_aws_service_access = []leanruntime.Field{
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: true},
}

var fields_disable_policy_type = []leanruntime.Field{
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
	{Name: "RootId", Flag: "root-id", Type: "*string", Required: true},
}

var fields_enable_all_features = []leanruntime.Field{}

var fields_enable_aws_service_access = []leanruntime.Field{
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: true},
}

var fields_enable_policy_type = []leanruntime.Field{
	{Name: "PolicyType", Flag: "policy-type", Type: "types.PolicyType", Required: true},
	{Name: "RootId", Flag: "root-id", Type: "*string", Required: true},
}

var fields_invite_account_to_organization = []leanruntime.Field{
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.HandshakeParty", Required: true},
}

var fields_invite_organization_to_transfer_responsibility = []leanruntime.Field{
	{Name: "Notes", Flag: "notes", Type: "*string", Required: false},
	{Name: "SourceName", Flag: "source-name", Type: "*string", Required: true},
	{Name: "StartTimestamp", Flag: "start-timestamp", Type: "*time.Time", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Target", Flag: "target", Type: "*types.HandshakeParty", Required: true},
	{Name: "Type", Flag: "type", Type: "types.ResponsibilityTransferType", Required: true},
}

var fields_leave_organization = []leanruntime.Field{}

var fields_list_accounts = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_accounts_for_parent = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentId", Flag: "parent-id", Type: "*string", Required: true},
}

var fields_list_accounts_with_invalid_effective_policy = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.EffectivePolicyType", Required: true},
}

var fields_list_aws_service_access_for_organization = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_children = []leanruntime.Field{
	{Name: "ChildType", Flag: "child-type", Type: "types.ChildType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentId", Flag: "parent-id", Type: "*string", Required: true},
}

var fields_list_create_account_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "States", Flag: "states", Type: "[]types.CreateAccountState", Required: false},
}

var fields_list_delegated_administrators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: false},
}

var fields_list_delegated_services_for_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_effective_policy_validation_errors = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyType", Flag: "policy-type", Type: "types.EffectivePolicyType", Required: true},
}

var fields_list_handshakes_for_account = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.HandshakeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_handshakes_for_organization = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.HandshakeFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_inbound_responsibility_transfers = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ResponsibilityTransferType", Required: true},
}

var fields_list_organizational_units_for_parent = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParentId", Flag: "parent-id", Type: "*string", Required: true},
}

var fields_list_outbound_responsibility_transfers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Type", Flag: "type", Type: "types.ResponsibilityTransferType", Required: true},
}

var fields_list_parents = []leanruntime.Field{
	{Name: "ChildId", Flag: "child-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policies = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "types.PolicyType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_policies_for_target = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "types.PolicyType", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TargetId", Flag: "target-id", Type: "*string", Required: true},
}

var fields_list_roots = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
}

var fields_list_targets_for_policy = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_move_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "DestinationParentId", Flag: "destination-parent-id", Type: "*string", Required: true},
	{Name: "SourceParentId", Flag: "source-parent-id", Type: "*string", Required: true},
}

var fields_put_resource_policy = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_register_delegated_administrator = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "ServicePrincipal", Flag: "service-principal", Type: "*string", Required: true},
}

var fields_remove_account_from_organization = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_terminate_responsibility_transfer = []leanruntime.Field{
	{Name: "EndTimestamp", Flag: "end-timestamp", Type: "*time.Time", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceId", Flag: "resource-id", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_organizational_unit = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OrganizationalUnitId", Flag: "organizational-unit-id", Type: "*string", Required: true},
}

var fields_update_policy = []leanruntime.Field{
	{Name: "Content", Flag: "content", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PolicyId", Flag: "policy-id", Type: "*string", Required: true},
}

var fields_update_responsibility_transfer = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-handshake": {
			Name:   "accept-handshake",
			Fields: fields_accept_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptHandshake(ctx, input)
			},
		},
		"attach-policy": {
			Name:   "attach-policy",
			Fields: fields_attach_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AttachPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_attach_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AttachPolicy(ctx, input)
			},
		},
		"cancel-handshake": {
			Name:   "cancel-handshake",
			Fields: fields_cancel_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelHandshake(ctx, input)
			},
		},
		"close-account": {
			Name:   "close-account",
			Fields: fields_close_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CloseAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_close_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CloseAccount(ctx, input)
			},
		},
		"create-account": {
			Name:   "create-account",
			Fields: fields_create_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateAccount(ctx, input)
			},
		},
		"create-gov-cloud-account": {
			Name:   "create-gov-cloud-account",
			Fields: fields_create_gov_cloud_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGovCloudAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_gov_cloud_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGovCloudAccount(ctx, input)
			},
		},
		"create-organization": {
			Name:   "create-organization",
			Fields: fields_create_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOrganization(ctx, input)
			},
		},
		"create-organizational-unit": {
			Name:   "create-organizational-unit",
			Fields: fields_create_organizational_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateOrganizationalUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_organizational_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateOrganizationalUnit(ctx, input)
			},
		},
		"create-policy": {
			Name:   "create-policy",
			Fields: fields_create_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePolicy(ctx, input)
			},
		},
		"decline-handshake": {
			Name:   "decline-handshake",
			Fields: fields_decline_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeclineHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_decline_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeclineHandshake(ctx, input)
			},
		},
		"delete-organization": {
			Name:   "delete-organization",
			Fields: fields_delete_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOrganization(ctx, input)
			},
		},
		"delete-organizational-unit": {
			Name:   "delete-organizational-unit",
			Fields: fields_delete_organizational_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteOrganizationalUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_organizational_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteOrganizationalUnit(ctx, input)
			},
		},
		"delete-policy": {
			Name:   "delete-policy",
			Fields: fields_delete_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePolicy(ctx, input)
			},
		},
		"delete-resource-policy": {
			Name:   "delete-resource-policy",
			Fields: fields_delete_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteResourcePolicy(ctx, input)
			},
		},
		"deregister-delegated-administrator": {
			Name:   "deregister-delegated-administrator",
			Fields: fields_deregister_delegated_administrator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeregisterDelegatedAdministratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_deregister_delegated_administrator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeregisterDelegatedAdministrator(ctx, input)
			},
		},
		"describe-account": {
			Name:   "describe-account",
			Fields: fields_describe_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeAccount(ctx, input)
			},
		},
		"describe-create-account-status": {
			Name:   "describe-create-account-status",
			Fields: fields_describe_create_account_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCreateAccountStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_create_account_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCreateAccountStatus(ctx, input)
			},
		},
		"describe-effective-policy": {
			Name:   "describe-effective-policy",
			Fields: fields_describe_effective_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeEffectivePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_effective_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeEffectivePolicy(ctx, input)
			},
		},
		"describe-handshake": {
			Name:   "describe-handshake",
			Fields: fields_describe_handshake,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeHandshakeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_handshake, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeHandshake(ctx, input)
			},
		},
		"describe-organization": {
			Name:   "describe-organization",
			Fields: fields_describe_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganization(ctx, input)
			},
		},
		"describe-organizational-unit": {
			Name:   "describe-organizational-unit",
			Fields: fields_describe_organizational_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeOrganizationalUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_organizational_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeOrganizationalUnit(ctx, input)
			},
		},
		"describe-policy": {
			Name:   "describe-policy",
			Fields: fields_describe_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePolicy(ctx, input)
			},
		},
		"describe-resource-policy": {
			Name:   "describe-resource-policy",
			Fields: fields_describe_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResourcePolicy(ctx, input)
			},
		},
		"describe-responsibility-transfer": {
			Name:   "describe-responsibility-transfer",
			Fields: fields_describe_responsibility_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeResponsibilityTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_responsibility_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeResponsibilityTransfer(ctx, input)
			},
		},
		"detach-policy": {
			Name:   "detach-policy",
			Fields: fields_detach_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DetachPolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_detach_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DetachPolicy(ctx, input)
			},
		},
		"disable-aws-service-access": {
			Name:   "disable-aws-service-access",
			Fields: fields_disable_aws_service_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableAWSServiceAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_aws_service_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableAWSServiceAccess(ctx, input)
			},
		},
		"disable-policy-type": {
			Name:   "disable-policy-type",
			Fields: fields_disable_policy_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisablePolicyTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_policy_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisablePolicyType(ctx, input)
			},
		},
		"enable-all-features": {
			Name:   "enable-all-features",
			Fields: fields_enable_all_features,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAllFeaturesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_all_features, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAllFeatures(ctx, input)
			},
		},
		"enable-aws-service-access": {
			Name:   "enable-aws-service-access",
			Fields: fields_enable_aws_service_access,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableAWSServiceAccessInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_aws_service_access, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableAWSServiceAccess(ctx, input)
			},
		},
		"enable-policy-type": {
			Name:   "enable-policy-type",
			Fields: fields_enable_policy_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnablePolicyTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_policy_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnablePolicyType(ctx, input)
			},
		},
		"invite-account-to-organization": {
			Name:   "invite-account-to-organization",
			Fields: fields_invite_account_to_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InviteAccountToOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invite_account_to_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InviteAccountToOrganization(ctx, input)
			},
		},
		"invite-organization-to-transfer-responsibility": {
			Name:   "invite-organization-to-transfer-responsibility",
			Fields: fields_invite_organization_to_transfer_responsibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.InviteOrganizationToTransferResponsibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_invite_organization_to_transfer_responsibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.InviteOrganizationToTransferResponsibility(ctx, input)
			},
		},
		"leave-organization": {
			Name:   "leave-organization",
			Fields: fields_leave_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.LeaveOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_leave_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.LeaveOrganization(ctx, input)
			},
		},
		"list-accounts": {
			Name:   "list-accounts",
			Fields: fields_list_accounts,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccounts(ctx, input)
				}
				var results []*svc.ListAccountsOutput
				p := svc.NewListAccountsPaginator(client, input)
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
		"list-accounts-for-parent": {
			Name:   "list-accounts-for-parent",
			Fields: fields_list_accounts_for_parent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsForParentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts_for_parent, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountsForParent(ctx, input)
				}
				var results []*svc.ListAccountsForParentOutput
				p := svc.NewListAccountsForParentPaginator(client, input)
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
		"list-accounts-with-invalid-effective-policy": {
			Name:   "list-accounts-with-invalid-effective-policy",
			Fields: fields_list_accounts_with_invalid_effective_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAccountsWithInvalidEffectivePolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_accounts_with_invalid_effective_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAccountsWithInvalidEffectivePolicy(ctx, input)
				}
				var results []*svc.ListAccountsWithInvalidEffectivePolicyOutput
				p := svc.NewListAccountsWithInvalidEffectivePolicyPaginator(client, input)
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
		"list-aws-service-access-for-organization": {
			Name:   "list-aws-service-access-for-organization",
			Fields: fields_list_aws_service_access_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListAWSServiceAccessForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_aws_service_access_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListAWSServiceAccessForOrganization(ctx, input)
				}
				var results []*svc.ListAWSServiceAccessForOrganizationOutput
				p := svc.NewListAWSServiceAccessForOrganizationPaginator(client, input)
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
		"list-children": {
			Name:   "list-children",
			Fields: fields_list_children,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListChildrenInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_children, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListChildren(ctx, input)
				}
				var results []*svc.ListChildrenOutput
				p := svc.NewListChildrenPaginator(client, input)
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
		"list-create-account-status": {
			Name:   "list-create-account-status",
			Fields: fields_list_create_account_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCreateAccountStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_create_account_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCreateAccountStatus(ctx, input)
				}
				var results []*svc.ListCreateAccountStatusOutput
				p := svc.NewListCreateAccountStatusPaginator(client, input)
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
		"list-delegated-administrators": {
			Name:   "list-delegated-administrators",
			Fields: fields_list_delegated_administrators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDelegatedAdministratorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_delegated_administrators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDelegatedAdministrators(ctx, input)
				}
				var results []*svc.ListDelegatedAdministratorsOutput
				p := svc.NewListDelegatedAdministratorsPaginator(client, input)
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
		"list-delegated-services-for-account": {
			Name:   "list-delegated-services-for-account",
			Fields: fields_list_delegated_services_for_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDelegatedServicesForAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_delegated_services_for_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDelegatedServicesForAccount(ctx, input)
				}
				var results []*svc.ListDelegatedServicesForAccountOutput
				p := svc.NewListDelegatedServicesForAccountPaginator(client, input)
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
		"list-effective-policy-validation-errors": {
			Name:   "list-effective-policy-validation-errors",
			Fields: fields_list_effective_policy_validation_errors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEffectivePolicyValidationErrorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_effective_policy_validation_errors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEffectivePolicyValidationErrors(ctx, input)
				}
				var results []*svc.ListEffectivePolicyValidationErrorsOutput
				p := svc.NewListEffectivePolicyValidationErrorsPaginator(client, input)
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
		"list-handshakes-for-account": {
			Name:   "list-handshakes-for-account",
			Fields: fields_list_handshakes_for_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHandshakesForAccountInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_handshakes_for_account, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHandshakesForAccount(ctx, input)
				}
				var results []*svc.ListHandshakesForAccountOutput
				p := svc.NewListHandshakesForAccountPaginator(client, input)
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
		"list-handshakes-for-organization": {
			Name:   "list-handshakes-for-organization",
			Fields: fields_list_handshakes_for_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListHandshakesForOrganizationInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_handshakes_for_organization, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListHandshakesForOrganization(ctx, input)
				}
				var results []*svc.ListHandshakesForOrganizationOutput
				p := svc.NewListHandshakesForOrganizationPaginator(client, input)
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
		"list-inbound-responsibility-transfers": {
			Name:   "list-inbound-responsibility-transfers",
			Fields: fields_list_inbound_responsibility_transfers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListInboundResponsibilityTransfersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_inbound_responsibility_transfers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListInboundResponsibilityTransfers(ctx, input)
			},
		},
		"list-organizational-units-for-parent": {
			Name:   "list-organizational-units-for-parent",
			Fields: fields_list_organizational_units_for_parent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOrganizationalUnitsForParentInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_organizational_units_for_parent, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOrganizationalUnitsForParent(ctx, input)
				}
				var results []*svc.ListOrganizationalUnitsForParentOutput
				p := svc.NewListOrganizationalUnitsForParentPaginator(client, input)
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
		"list-outbound-responsibility-transfers": {
			Name:   "list-outbound-responsibility-transfers",
			Fields: fields_list_outbound_responsibility_transfers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOutboundResponsibilityTransfersInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_outbound_responsibility_transfers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListOutboundResponsibilityTransfers(ctx, input)
			},
		},
		"list-parents": {
			Name:   "list-parents",
			Fields: fields_list_parents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListParentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_parents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParents(ctx, input)
				}
				var results []*svc.ListParentsOutput
				p := svc.NewListParentsPaginator(client, input)
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
		"list-policies": {
			Name:   "list-policies",
			Fields: fields_list_policies,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policies, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPolicies(ctx, input)
				}
				var results []*svc.ListPoliciesOutput
				p := svc.NewListPoliciesPaginator(client, input)
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
		"list-policies-for-target": {
			Name:   "list-policies-for-target",
			Fields: fields_list_policies_for_target,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPoliciesForTargetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_policies_for_target, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPoliciesForTarget(ctx, input)
				}
				var results []*svc.ListPoliciesForTargetOutput
				p := svc.NewListPoliciesForTargetPaginator(client, input)
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
		"list-roots": {
			Name:   "list-roots",
			Fields: fields_list_roots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRootsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_roots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRoots(ctx, input)
				}
				var results []*svc.ListRootsOutput
				p := svc.NewListRootsPaginator(client, input)
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
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tags_for_resource, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTagsForResource(ctx, input)
				}
				var results []*svc.ListTagsForResourceOutput
				p := svc.NewListTagsForResourcePaginator(client, input)
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
		"list-targets-for-policy": {
			Name:   "list-targets-for-policy",
			Fields: fields_list_targets_for_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTargetsForPolicyInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_targets_for_policy, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTargetsForPolicy(ctx, input)
				}
				var results []*svc.ListTargetsForPolicyOutput
				p := svc.NewListTargetsForPolicyPaginator(client, input)
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
		"move-account": {
			Name:   "move-account",
			Fields: fields_move_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.MoveAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_move_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.MoveAccount(ctx, input)
			},
		},
		"put-resource-policy": {
			Name:   "put-resource-policy",
			Fields: fields_put_resource_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutResourcePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_resource_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutResourcePolicy(ctx, input)
			},
		},
		"register-delegated-administrator": {
			Name:   "register-delegated-administrator",
			Fields: fields_register_delegated_administrator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDelegatedAdministratorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_delegated_administrator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDelegatedAdministrator(ctx, input)
			},
		},
		"remove-account-from-organization": {
			Name:   "remove-account-from-organization",
			Fields: fields_remove_account_from_organization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveAccountFromOrganizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_account_from_organization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveAccountFromOrganization(ctx, input)
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
		"terminate-responsibility-transfer": {
			Name:   "terminate-responsibility-transfer",
			Fields: fields_terminate_responsibility_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TerminateResponsibilityTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_terminate_responsibility_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TerminateResponsibilityTransfer(ctx, input)
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
		"update-organizational-unit": {
			Name:   "update-organizational-unit",
			Fields: fields_update_organizational_unit,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateOrganizationalUnitInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_organizational_unit, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateOrganizationalUnit(ctx, input)
			},
		},
		"update-policy": {
			Name:   "update-policy",
			Fields: fields_update_policy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePolicyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_policy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePolicy(ctx, input)
			},
		},
		"update-responsibility-transfer": {
			Name:   "update-responsibility-transfer",
			Fields: fields_update_responsibility_transfer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateResponsibilityTransferInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_responsibility_transfer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateResponsibilityTransfer(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("organizations", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
