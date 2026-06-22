package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/rolesanywhere"
)

var fields_create_profile = []leanruntime.Field{
	{Name: "AcceptRoleSessionName", Flag: "accept-role-session-name", Type: "*bool", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "ManagedPolicyArns", Flag: "managed-policy-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "RequireInstanceProperties", Flag: "require-instance-properties", Type: "*bool", Required: false},
	{Name: "RoleArns", Flag: "role-arns", Type: "[]string", Required: true},
	{Name: "SessionPolicy", Flag: "session-policy", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_trust_anchor = []leanruntime.Field{
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NotificationSettings", Flag: "notification-settings", Type: "[]types.NotificationSetting", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.Source", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_attribute_mapping = []leanruntime.Field{
	{Name: "CertificateField", Flag: "certificate-field", Type: "types.CertificateField", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "Specifiers", Flag: "specifiers", Type: "[]string", Required: false},
}

var fields_delete_crl = []leanruntime.Field{
	{Name: "CrlId", Flag: "crl-id", Type: "*string", Required: true},
}

var fields_delete_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_delete_trust_anchor = []leanruntime.Field{
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

var fields_disable_crl = []leanruntime.Field{
	{Name: "CrlId", Flag: "crl-id", Type: "*string", Required: true},
}

var fields_disable_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_disable_trust_anchor = []leanruntime.Field{
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

var fields_enable_crl = []leanruntime.Field{
	{Name: "CrlId", Flag: "crl-id", Type: "*string", Required: true},
}

var fields_enable_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_enable_trust_anchor = []leanruntime.Field{
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

var fields_get_crl = []leanruntime.Field{
	{Name: "CrlId", Flag: "crl-id", Type: "*string", Required: true},
}

var fields_get_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_get_subject = []leanruntime.Field{
	{Name: "SubjectId", Flag: "subject-id", Type: "*string", Required: true},
}

var fields_get_trust_anchor = []leanruntime.Field{
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

var fields_import_crl = []leanruntime.Field{
	{Name: "CrlData", Flag: "crl-data", Type: "[]byte", Required: true},
	{Name: "Enabled", Flag: "enabled", Type: "*bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TrustAnchorArn", Flag: "trust-anchor-arn", Type: "*string", Required: true},
}

var fields_list_crls = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_profiles = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_subjects = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_trust_anchors = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PageSize", Flag: "page-size", Type: "*int32", Required: false},
}

var fields_put_attribute_mapping = []leanruntime.Field{
	{Name: "CertificateField", Flag: "certificate-field", Type: "types.CertificateField", Required: true},
	{Name: "MappingRules", Flag: "mapping-rules", Type: "[]types.MappingRule", Required: true},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_put_notification_settings = []leanruntime.Field{
	{Name: "NotificationSettings", Flag: "notification-settings", Type: "[]types.NotificationSetting", Required: true},
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

var fields_reset_notification_settings = []leanruntime.Field{
	{Name: "NotificationSettingKeys", Flag: "notification-setting-keys", Type: "[]types.NotificationSettingKey", Required: true},
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_crl = []leanruntime.Field{
	{Name: "CrlData", Flag: "crl-data", Type: "[]byte", Required: false},
	{Name: "CrlId", Flag: "crl-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_profile = []leanruntime.Field{
	{Name: "AcceptRoleSessionName", Flag: "accept-role-session-name", Type: "*bool", Required: false},
	{Name: "DurationSeconds", Flag: "duration-seconds", Type: "*int32", Required: false},
	{Name: "ManagedPolicyArns", Flag: "managed-policy-arns", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "RoleArns", Flag: "role-arns", Type: "[]string", Required: false},
	{Name: "SessionPolicy", Flag: "session-policy", Type: "*string", Required: false},
}

var fields_update_trust_anchor = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.Source", Required: false},
	{Name: "TrustAnchorId", Flag: "trust-anchor-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-profile": {
			Name:   "create-profile",
			Fields: fields_create_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateProfile(ctx, input)
			},
		},
		"create-trust-anchor": {
			Name:   "create-trust-anchor",
			Fields: fields_create_trust_anchor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrustAnchorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_trust_anchor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTrustAnchor(ctx, input)
			},
		},
		"delete-attribute-mapping": {
			Name:   "delete-attribute-mapping",
			Fields: fields_delete_attribute_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteAttributeMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_attribute_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteAttributeMapping(ctx, input)
			},
		},
		"delete-crl": {
			Name:   "delete-crl",
			Fields: fields_delete_crl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_crl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCrl(ctx, input)
			},
		},
		"delete-profile": {
			Name:   "delete-profile",
			Fields: fields_delete_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteProfile(ctx, input)
			},
		},
		"delete-trust-anchor": {
			Name:   "delete-trust-anchor",
			Fields: fields_delete_trust_anchor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrustAnchorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_trust_anchor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTrustAnchor(ctx, input)
			},
		},
		"disable-crl": {
			Name:   "disable-crl",
			Fields: fields_disable_crl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableCrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_crl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableCrl(ctx, input)
			},
		},
		"disable-profile": {
			Name:   "disable-profile",
			Fields: fields_disable_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableProfile(ctx, input)
			},
		},
		"disable-trust-anchor": {
			Name:   "disable-trust-anchor",
			Fields: fields_disable_trust_anchor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisableTrustAnchorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disable_trust_anchor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisableTrustAnchor(ctx, input)
			},
		},
		"enable-crl": {
			Name:   "enable-crl",
			Fields: fields_enable_crl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableCrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_crl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableCrl(ctx, input)
			},
		},
		"enable-profile": {
			Name:   "enable-profile",
			Fields: fields_enable_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableProfile(ctx, input)
			},
		},
		"enable-trust-anchor": {
			Name:   "enable-trust-anchor",
			Fields: fields_enable_trust_anchor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EnableTrustAnchorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_enable_trust_anchor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EnableTrustAnchor(ctx, input)
			},
		},
		"get-crl": {
			Name:   "get-crl",
			Fields: fields_get_crl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_crl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCrl(ctx, input)
			},
		},
		"get-profile": {
			Name:   "get-profile",
			Fields: fields_get_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetProfile(ctx, input)
			},
		},
		"get-subject": {
			Name:   "get-subject",
			Fields: fields_get_subject,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSubjectInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_subject, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSubject(ctx, input)
			},
		},
		"get-trust-anchor": {
			Name:   "get-trust-anchor",
			Fields: fields_get_trust_anchor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTrustAnchorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_trust_anchor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTrustAnchor(ctx, input)
			},
		},
		"import-crl": {
			Name:   "import-crl",
			Fields: fields_import_crl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportCrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_crl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportCrl(ctx, input)
			},
		},
		"list-crls": {
			Name:   "list-crls",
			Fields: fields_list_crls,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCrlsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_crls, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCrls(ctx, input)
				}
				var results []*svc.ListCrlsOutput
				p := svc.NewListCrlsPaginator(client, input)
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
		"list-profiles": {
			Name:   "list-profiles",
			Fields: fields_list_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListProfiles(ctx, input)
				}
				var results []*svc.ListProfilesOutput
				p := svc.NewListProfilesPaginator(client, input)
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
		"list-subjects": {
			Name:   "list-subjects",
			Fields: fields_list_subjects,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSubjectsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_subjects, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSubjects(ctx, input)
				}
				var results []*svc.ListSubjectsOutput
				p := svc.NewListSubjectsPaginator(client, input)
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
		"list-trust-anchors": {
			Name:   "list-trust-anchors",
			Fields: fields_list_trust_anchors,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrustAnchorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trust_anchors, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrustAnchors(ctx, input)
				}
				var results []*svc.ListTrustAnchorsOutput
				p := svc.NewListTrustAnchorsPaginator(client, input)
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
		"put-attribute-mapping": {
			Name:   "put-attribute-mapping",
			Fields: fields_put_attribute_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutAttributeMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_attribute_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutAttributeMapping(ctx, input)
			},
		},
		"put-notification-settings": {
			Name:   "put-notification-settings",
			Fields: fields_put_notification_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutNotificationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_notification_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutNotificationSettings(ctx, input)
			},
		},
		"reset-notification-settings": {
			Name:   "reset-notification-settings",
			Fields: fields_reset_notification_settings,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResetNotificationSettingsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_reset_notification_settings, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResetNotificationSettings(ctx, input)
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
		"update-crl": {
			Name:   "update-crl",
			Fields: fields_update_crl,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCrlInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_crl, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCrl(ctx, input)
			},
		},
		"update-profile": {
			Name:   "update-profile",
			Fields: fields_update_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateProfile(ctx, input)
			},
		},
		"update-trust-anchor": {
			Name:   "update-trust-anchor",
			Fields: fields_update_trust_anchor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrustAnchorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_trust_anchor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTrustAnchor(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("rolesanywhere", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
