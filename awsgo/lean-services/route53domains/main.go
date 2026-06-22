package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/route53domains"
)

var fields_accept_domain_transfer_from_another_aws_account = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Password", Flag: "password", Type: "*string", Required: true},
}

var fields_associate_delegation_signer_to_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "SigningAttributes", Flag: "signing-attributes", Type: "*types.DnssecSigningAttributes", Required: true},
}

var fields_cancel_domain_transfer_to_another_aws_account = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_check_domain_availability = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "IdnLangCode", Flag: "idn-lang-code", Type: "*string", Required: false},
}

var fields_check_domain_transferability = []leanruntime.Field{
	{Name: "AuthCode", Flag: "auth-code", Type: "*string", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_delete_tags_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "TagsToDelete", Flag: "tags-to-delete", Type: "[]string", Required: true},
}

var fields_disable_domain_auto_renew = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_disable_domain_transfer_lock = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_disassociate_delegation_signer_from_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_enable_domain_auto_renew = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_enable_domain_transfer_lock = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_contact_reachability_status = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
}

var fields_get_domain_detail = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_get_domain_suggestions = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "OnlyAvailable", Flag: "only-available", Type: "*bool", Required: true},
	{Name: "SuggestionCount", Flag: "suggestion-count", Type: "int32", Required: true},
}

var fields_get_operation_detail = []leanruntime.Field{
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "FilterConditions", Flag: "filter-conditions", Type: "[]types.FilterCondition", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "SortCondition", Flag: "sort-condition", Type: "*types.SortCondition", Required: false},
}

var fields_list_operations = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "SortBy", Flag: "sort-by", Type: "types.ListOperationsSortAttributeName", Required: false},
	{Name: "SortOrder", Flag: "sort-order", Type: "types.SortOrder", Required: false},
	{Name: "Status", Flag: "status", Type: "[]types.OperationStatus", Required: false},
	{Name: "SubmittedSince", Flag: "submitted-since", Type: "*time.Time", Required: false},
	{Name: "Type", Flag: "type", Type: "[]types.OperationType", Required: false},
}

var fields_list_prices = []leanruntime.Field{
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Tld", Flag: "tld", Type: "*string", Required: false},
}

var fields_list_tags_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_push_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "Target", Flag: "target", Type: "*string", Required: true},
}

var fields_register_domain = []leanruntime.Field{
	{Name: "AdminContact", Flag: "admin-contact", Type: "*types.ContactDetail", Required: true},
	{Name: "AutoRenew", Flag: "auto-renew", Type: "*bool", Required: false},
	{Name: "BillingContact", Flag: "billing-contact", Type: "*types.ContactDetail", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DurationInYears", Flag: "duration-in-years", Type: "*int32", Required: true},
	{Name: "IdnLangCode", Flag: "idn-lang-code", Type: "*string", Required: false},
	{Name: "PrivacyProtectAdminContact", Flag: "privacy-protect-admin-contact", Type: "*bool", Required: false},
	{Name: "PrivacyProtectBillingContact", Flag: "privacy-protect-billing-contact", Type: "*bool", Required: false},
	{Name: "PrivacyProtectRegistrantContact", Flag: "privacy-protect-registrant-contact", Type: "*bool", Required: false},
	{Name: "PrivacyProtectTechContact", Flag: "privacy-protect-tech-contact", Type: "*bool", Required: false},
	{Name: "RegistrantContact", Flag: "registrant-contact", Type: "*types.ContactDetail", Required: true},
	{Name: "TechContact", Flag: "tech-contact", Type: "*types.ContactDetail", Required: true},
}

var fields_reject_domain_transfer_from_another_aws_account = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_renew_domain = []leanruntime.Field{
	{Name: "CurrentExpiryYear", Flag: "current-expiry-year", Type: "int32", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DurationInYears", Flag: "duration-in-years", Type: "*int32", Required: false},
}

var fields_resend_contact_reachability_email = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: false},
}

var fields_resend_operation_authorization = []leanruntime.Field{
	{Name: "OperationId", Flag: "operation-id", Type: "*string", Required: true},
}

var fields_retrieve_domain_auth_code = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_transfer_domain = []leanruntime.Field{
	{Name: "AdminContact", Flag: "admin-contact", Type: "*types.ContactDetail", Required: true},
	{Name: "AuthCode", Flag: "auth-code", Type: "*string", Required: false},
	{Name: "AutoRenew", Flag: "auto-renew", Type: "*bool", Required: false},
	{Name: "BillingContact", Flag: "billing-contact", Type: "*types.ContactDetail", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "DurationInYears", Flag: "duration-in-years", Type: "*int32", Required: true},
	{Name: "IdnLangCode", Flag: "idn-lang-code", Type: "*string", Required: false},
	{Name: "Nameservers", Flag: "nameservers", Type: "[]types.Nameserver", Required: false},
	{Name: "PrivacyProtectAdminContact", Flag: "privacy-protect-admin-contact", Type: "*bool", Required: false},
	{Name: "PrivacyProtectBillingContact", Flag: "privacy-protect-billing-contact", Type: "*bool", Required: false},
	{Name: "PrivacyProtectRegistrantContact", Flag: "privacy-protect-registrant-contact", Type: "*bool", Required: false},
	{Name: "PrivacyProtectTechContact", Flag: "privacy-protect-tech-contact", Type: "*bool", Required: false},
	{Name: "RegistrantContact", Flag: "registrant-contact", Type: "*types.ContactDetail", Required: true},
	{Name: "TechContact", Flag: "tech-contact", Type: "*types.ContactDetail", Required: true},
}

var fields_transfer_domain_to_another_aws_account = []leanruntime.Field{
	{Name: "AccountId", Flag: "account-id", Type: "*string", Required: true},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
}

var fields_update_domain_contact = []leanruntime.Field{
	{Name: "AdminContact", Flag: "admin-contact", Type: "*types.ContactDetail", Required: false},
	{Name: "BillingContact", Flag: "billing-contact", Type: "*types.ContactDetail", Required: false},
	{Name: "Consent", Flag: "consent", Type: "*types.Consent", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RegistrantContact", Flag: "registrant-contact", Type: "*types.ContactDetail", Required: false},
	{Name: "TechContact", Flag: "tech-contact", Type: "*types.ContactDetail", Required: false},
}

var fields_update_domain_contact_privacy = []leanruntime.Field{
	{Name: "AdminPrivacy", Flag: "admin-privacy", Type: "*bool", Required: false},
	{Name: "BillingPrivacy", Flag: "billing-privacy", Type: "*bool", Required: false},
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "RegistrantPrivacy", Flag: "registrant-privacy", Type: "*bool", Required: false},
	{Name: "TechPrivacy", Flag: "tech-privacy", Type: "*bool", Required: false},
}

var fields_update_domain_nameservers = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "FIAuthKey", Flag: "fi-auth-key", Type: "*string", Required: false},
	{Name: "Nameservers", Flag: "nameservers", Type: "[]types.Nameserver", Required: true},
}

var fields_update_tags_for_domain = []leanruntime.Field{
	{Name: "DomainName", Flag: "domain-name", Type: "*string", Required: true},
	{Name: "TagsToUpdate", Flag: "tags-to-update", Type: "[]types.Tag", Required: false},
}

var fields_view_billing = []leanruntime.Field{
	{Name: "End", Flag: "end", Type: "*time.Time", Required: false},
	{Name: "Marker", Flag: "marker", Type: "*string", Required: false},
	{Name: "MaxItems", Flag: "max-items", Type: "*int32", Required: false},
	{Name: "Start", Flag: "start", Type: "*time.Time", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"accept-domain-transfer-from-another-aws-account": {
			Name:   "accept-domain-transfer-from-another-aws-account",
			Fields: fields_accept_domain_transfer_from_another_aws_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AcceptDomainTransferFromAnotherAwsAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_accept_domain_transfer_from_another_aws_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AcceptDomainTransferFromAnotherAwsAccount(ctx, input)
			},
		},
		"associate-delegation-signer-to-domain": {
			Name:   "associate-delegation-signer-to-domain",
			Fields: fields_associate_delegation_signer_to_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateDelegationSignerToDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_delegation_signer_to_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateDelegationSignerToDomain(ctx, input)
			},
		},
		"cancel-domain-transfer-to-another-aws-account": {
			Name:   "cancel-domain-transfer-to-another-aws-account",
			Fields: fields_cancel_domain_transfer_to_another_aws_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelDomainTransferToAnotherAwsAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_domain_transfer_to_another_aws_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelDomainTransferToAnotherAwsAccount(ctx, input)
			},
		},
		"check-domain-availability": {
			Name:   "check-domain-availability",
			Fields: fields_check_domain_availability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckDomainAvailabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_domain_availability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckDomainAvailability(ctx, input)
			},
		},
		"check-domain-transferability": {
			Name:   "check-domain-transferability",
			Fields: fields_check_domain_transferability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CheckDomainTransferabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_check_domain_transferability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CheckDomainTransferability(ctx, input)
			},
		},
		"delete-domain": {
			Name:   "delete-domain",
			Fields: fields_delete_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDomain(ctx, input)
			},
		},
		"delete-tags-for-domain": {
			Name:   "delete-tags-for-domain",
			Fields: fields_delete_tags_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTagsForDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tags_for_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTagsForDomain(ctx, input)
			},
		},
		"disable-domain-auto-renew": {
			Name:   "disable-domain-auto-renew",
			Fields: fields_disable_domain_auto_renew,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDomainAutoRenewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_domain_auto_renew, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDomainAutoRenew(ctx, input)
			},
		},
		"disable-domain-transfer-lock": {
			Name:   "disable-domain-transfer-lock",
			Fields: fields_disable_domain_transfer_lock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableDomainTransferLockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_domain_transfer_lock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableDomainTransferLock(ctx, input)
			},
		},
		"disassociate-delegation-signer-from-domain": {
			Name:   "disassociate-delegation-signer-from-domain",
			Fields: fields_disassociate_delegation_signer_from_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateDelegationSignerFromDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_delegation_signer_from_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateDelegationSignerFromDomain(ctx, input)
			},
		},
		"enable-domain-auto-renew": {
			Name:   "enable-domain-auto-renew",
			Fields: fields_enable_domain_auto_renew,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDomainAutoRenewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_domain_auto_renew, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDomainAutoRenew(ctx, input)
			},
		},
		"enable-domain-transfer-lock": {
			Name:   "enable-domain-transfer-lock",
			Fields: fields_enable_domain_transfer_lock,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableDomainTransferLockInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_domain_transfer_lock, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableDomainTransferLock(ctx, input)
			},
		},
		"get-contact-reachability-status": {
			Name:   "get-contact-reachability-status",
			Fields: fields_get_contact_reachability_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetContactReachabilityStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_contact_reachability_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetContactReachabilityStatus(ctx, input)
			},
		},
		"get-domain-detail": {
			Name:   "get-domain-detail",
			Fields: fields_get_domain_detail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainDetailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_detail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainDetail(ctx, input)
			},
		},
		"get-domain-suggestions": {
			Name:   "get-domain-suggestions",
			Fields: fields_get_domain_suggestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDomainSuggestionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_domain_suggestions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDomainSuggestions(ctx, input)
			},
		},
		"get-operation-detail": {
			Name:   "get-operation-detail",
			Fields: fields_get_operation_detail,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetOperationDetailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_operation_detail, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetOperationDetail(ctx, input)
			},
		},
		"list-domains": {
			Name:   "list-domains",
			Fields: fields_list_domains,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDomainsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_domains, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDomains(ctx, input)
				}
				var results []*svc.ListDomainsOutput
				p := svc.NewListDomainsPaginator(client, input)
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
		"list-operations": {
			Name:   "list-operations",
			Fields: fields_list_operations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListOperationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_operations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListOperations(ctx, input)
				}
				var results []*svc.ListOperationsOutput
				p := svc.NewListOperationsPaginator(client, input)
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
		"list-prices": {
			Name:   "list-prices",
			Fields: fields_list_prices,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPricesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_prices, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPrices(ctx, input)
				}
				var results []*svc.ListPricesOutput
				p := svc.NewListPricesPaginator(client, input)
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
		"list-tags-for-domain": {
			Name:   "list-tags-for-domain",
			Fields: fields_list_tags_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTagsForDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_tags_for_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListTagsForDomain(ctx, input)
			},
		},
		"push-domain": {
			Name:   "push-domain",
			Fields: fields_push_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PushDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_push_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PushDomain(ctx, input)
			},
		},
		"register-domain": {
			Name:   "register-domain",
			Fields: fields_register_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterDomain(ctx, input)
			},
		},
		"reject-domain-transfer-from-another-aws-account": {
			Name:   "reject-domain-transfer-from-another-aws-account",
			Fields: fields_reject_domain_transfer_from_another_aws_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RejectDomainTransferFromAnotherAwsAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reject_domain_transfer_from_another_aws_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RejectDomainTransferFromAnotherAwsAccount(ctx, input)
			},
		},
		"renew-domain": {
			Name:   "renew-domain",
			Fields: fields_renew_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RenewDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_renew_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RenewDomain(ctx, input)
			},
		},
		"resend-contact-reachability-email": {
			Name:   "resend-contact-reachability-email",
			Fields: fields_resend_contact_reachability_email,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResendContactReachabilityEmailInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resend_contact_reachability_email, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResendContactReachabilityEmail(ctx, input)
			},
		},
		"resend-operation-authorization": {
			Name:   "resend-operation-authorization",
			Fields: fields_resend_operation_authorization,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResendOperationAuthorizationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resend_operation_authorization, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResendOperationAuthorization(ctx, input)
			},
		},
		"retrieve-domain-auth-code": {
			Name:   "retrieve-domain-auth-code",
			Fields: fields_retrieve_domain_auth_code,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RetrieveDomainAuthCodeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_retrieve_domain_auth_code, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RetrieveDomainAuthCode(ctx, input)
			},
		},
		"transfer-domain": {
			Name:   "transfer-domain",
			Fields: fields_transfer_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransferDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transfer_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransferDomain(ctx, input)
			},
		},
		"transfer-domain-to-another-aws-account": {
			Name:   "transfer-domain-to-another-aws-account",
			Fields: fields_transfer_domain_to_another_aws_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TransferDomainToAnotherAwsAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_transfer_domain_to_another_aws_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TransferDomainToAnotherAwsAccount(ctx, input)
			},
		},
		"update-domain-contact": {
			Name:   "update-domain-contact",
			Fields: fields_update_domain_contact,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainContactInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_contact, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainContact(ctx, input)
			},
		},
		"update-domain-contact-privacy": {
			Name:   "update-domain-contact-privacy",
			Fields: fields_update_domain_contact_privacy,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainContactPrivacyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_contact_privacy, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainContactPrivacy(ctx, input)
			},
		},
		"update-domain-nameservers": {
			Name:   "update-domain-nameservers",
			Fields: fields_update_domain_nameservers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainNameserversInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain_nameservers, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomainNameservers(ctx, input)
			},
		},
		"update-tags-for-domain": {
			Name:   "update-tags-for-domain",
			Fields: fields_update_tags_for_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTagsForDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_tags_for_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTagsForDomain(ctx, input)
			},
		},
		"view-billing": {
			Name:   "view-billing",
			Fields: fields_view_billing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ViewBillingInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_view_billing, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ViewBilling(ctx, input)
				}
				var results []*svc.ViewBillingOutput
				p := svc.NewViewBillingPaginator(client, input)
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
	}
	if err := leanruntime.Execute("route53domains", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
