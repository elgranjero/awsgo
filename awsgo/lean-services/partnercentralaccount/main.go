package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/partnercentralaccount"
)

var fields_accept_connection_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_associate_aws_training_certification_email_domain = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "EmailVerificationCode", Flag: "email-verification-code", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_cancel_connection = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_cancel_connection_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_cancel_profile_update_task = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "TaskId", Flag: "task-id", Type: "*string", Required: true},
}

var fields_create_connection_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "Message", Flag: "message", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ReceiverIdentifier", Flag: "receiver-identifier", Type: "*string", Required: true},
}

var fields_create_partner = []leanruntime.Field{
	{Name: "AllianceLeadContact", Flag: "alliance-lead-contact", Type: "*types.AllianceLeadContact", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EmailVerificationCode", Flag: "email-verification-code", Type: "*string", Required: true},
	{Name: "LegalName", Flag: "legal-name", Type: "*string", Required: true},
	{Name: "PrimarySolutionType", Flag: "primary-solution-type", Type: "types.PrimarySolutionType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_disassociate_aws_training_certification_email_domain = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_alliance_lead_contact = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_connection = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_connection_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_connection_preferences = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
}

var fields_get_partner = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_profile_update_task = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_profile_visibility = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_verification = []leanruntime.Field{
	{Name: "VerificationType", Flag: "verification-type", Type: "types.VerificationType", Required: true},
}

var fields_list_connection_invitations = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ConnectionType", Flag: "connection-type", Type: "types.ConnectionType", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OtherParticipantIdentifiers", Flag: "other-participant-identifiers", Type: "[]string", Required: false},
	{Name: "ParticipantType", Flag: "participant-type", Type: "types.ParticipantType", Required: false},
	{Name: "Status", Flag: "status", Type: "types.InvitationStatus", Required: false},
}

var fields_list_connections = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ConnectionType", Flag: "connection-type", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "OtherParticipantIdentifiers", Flag: "other-participant-identifiers", Type: "[]string", Required: false},
}

var fields_list_partners = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_alliance_lead_contact = []leanruntime.Field{
	{Name: "AllianceLeadContact", Flag: "alliance-lead-contact", Type: "*types.AllianceLeadContact", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "EmailVerificationCode", Flag: "email-verification-code", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_put_profile_visibility = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Visibility", Flag: "visibility", Type: "types.ProfileVisibility", Required: true},
}

var fields_reject_connection_invitation = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: true},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
}

var fields_send_email_verification_code = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
}

var fields_start_profile_update_task = []leanruntime.Field{
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "TaskDetails", Flag: "task-details", Type: "*types.TaskDetails", Required: true},
}

var fields_start_verification = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "VerificationDetails", Flag: "verification-details", Type: "types.VerificationDetails", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_connection_preferences = []leanruntime.Field{
	{Name: "AccessType", Flag: "access-type", Type: "types.AccessType", Required: true},
	{Name: "Catalog", Flag: "catalog", Type: "*string", Required: true},
	{Name: "ExcludedParticipantIdentifiers", Flag: "excluded-participant-identifiers", Type: "[]string", Required: false},
	{Name: "Revision", Flag: "revision", Type: "*int64", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-connection-invitation": {
			Name:   "accept-connection-invitation",
			Fields: fields_accept_connection_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptConnectionInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_connection_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptConnectionInvitation(ctx, input)
			},
		},
		"associate-aws-training-certification-email-domain": {
			Name:   "associate-aws-training-certification-email-domain",
			Fields: fields_associate_aws_training_certification_email_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateAwsTrainingCertificationEmailDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_aws_training_certification_email_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateAwsTrainingCertificationEmailDomain(ctx, input)
			},
		},
		"cancel-connection": {
			Name:   "cancel-connection",
			Fields: fields_cancel_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelConnection(ctx, input)
			},
		},
		"cancel-connection-invitation": {
			Name:   "cancel-connection-invitation",
			Fields: fields_cancel_connection_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelConnectionInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_connection_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelConnectionInvitation(ctx, input)
			},
		},
		"cancel-profile-update-task": {
			Name:   "cancel-profile-update-task",
			Fields: fields_cancel_profile_update_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelProfileUpdateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_profile_update_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelProfileUpdateTask(ctx, input)
			},
		},
		"create-connection-invitation": {
			Name:   "create-connection-invitation",
			Fields: fields_create_connection_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateConnectionInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_connection_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateConnectionInvitation(ctx, input)
			},
		},
		"create-partner": {
			Name:   "create-partner",
			Fields: fields_create_partner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartner(ctx, input)
			},
		},
		"disassociate-aws-training-certification-email-domain": {
			Name:   "disassociate-aws-training-certification-email-domain",
			Fields: fields_disassociate_aws_training_certification_email_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateAwsTrainingCertificationEmailDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_aws_training_certification_email_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateAwsTrainingCertificationEmailDomain(ctx, input)
			},
		},
		"get-alliance-lead-contact": {
			Name:   "get-alliance-lead-contact",
			Fields: fields_get_alliance_lead_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetAllianceLeadContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_alliance_lead_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetAllianceLeadContact(ctx, input)
			},
		},
		"get-connection": {
			Name:   "get-connection",
			Fields: fields_get_connection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnection(ctx, input)
			},
		},
		"get-connection-invitation": {
			Name:   "get-connection-invitation",
			Fields: fields_get_connection_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionInvitation(ctx, input)
			},
		},
		"get-connection-preferences": {
			Name:   "get-connection-preferences",
			Fields: fields_get_connection_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectionPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connection_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectionPreferences(ctx, input)
			},
		},
		"get-partner": {
			Name:   "get-partner",
			Fields: fields_get_partner,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPartnerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_partner, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPartner(ctx, input)
			},
		},
		"get-profile-update-task": {
			Name:   "get-profile-update-task",
			Fields: fields_get_profile_update_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileUpdateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_update_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileUpdateTask(ctx, input)
			},
		},
		"get-profile-visibility": {
			Name:   "get-profile-visibility",
			Fields: fields_get_profile_visibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileVisibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile_visibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfileVisibility(ctx, input)
			},
		},
		"get-verification": {
			Name:   "get-verification",
			Fields: fields_get_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVerification(ctx, input)
			},
		},
		"list-connection-invitations": {
			Name:   "list-connection-invitations",
			Fields: fields_list_connection_invitations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionInvitationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connection_invitations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnectionInvitations(ctx, input)
				}
				var results []*svc.ListConnectionInvitationsOutput
				p := svc.NewListConnectionInvitationsPaginator(client, input)
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
		"list-connections": {
			Name:   "list-connections",
			Fields: fields_list_connections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListConnectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_connections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListConnections(ctx, input)
				}
				var results []*svc.ListConnectionsOutput
				p := svc.NewListConnectionsPaginator(client, input)
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
		"list-partners": {
			Name:   "list-partners",
			Fields: fields_list_partners,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartnersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_partners, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPartners(ctx, input)
				}
				var results []*svc.ListPartnersOutput
				p := svc.NewListPartnersPaginator(client, input)
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
		"put-alliance-lead-contact": {
			Name:   "put-alliance-lead-contact",
			Fields: fields_put_alliance_lead_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAllianceLeadContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_alliance_lead_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAllianceLeadContact(ctx, input)
			},
		},
		"put-profile-visibility": {
			Name:   "put-profile-visibility",
			Fields: fields_put_profile_visibility,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutProfileVisibilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_profile_visibility, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutProfileVisibility(ctx, input)
			},
		},
		"reject-connection-invitation": {
			Name:   "reject-connection-invitation",
			Fields: fields_reject_connection_invitation,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectConnectionInvitationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_connection_invitation, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectConnectionInvitation(ctx, input)
			},
		},
		"send-email-verification-code": {
			Name:   "send-email-verification-code",
			Fields: fields_send_email_verification_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SendEmailVerificationCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_send_email_verification_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SendEmailVerificationCode(ctx, input)
			},
		},
		"start-profile-update-task": {
			Name:   "start-profile-update-task",
			Fields: fields_start_profile_update_task,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartProfileUpdateTaskInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_profile_update_task, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartProfileUpdateTask(ctx, input)
			},
		},
		"start-verification": {
			Name:   "start-verification",
			Fields: fields_start_verification,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartVerificationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_verification, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartVerification(ctx, input)
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
		"update-connection-preferences": {
			Name:   "update-connection-preferences",
			Fields: fields_update_connection_preferences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateConnectionPreferencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_connection_preferences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateConnectionPreferences(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("partnercentralaccount", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
