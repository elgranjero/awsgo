package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/signer"
)

var fields_add_profile_permission = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "*string", Required: true},
	{Name: "Principal", Flag: "principal", Type: "*string", Required: true},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "ProfileVersion", Flag: "profile-version", Type: "*string", Required: false},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: false},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_cancel_signing_profile = []leanruntime.Field{
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
}

var fields_describe_signing_job = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_get_revocation_status = []leanruntime.Field{
	{Name: "CertificateHashes", Flag: "certificate-hashes", Type: "[]string", Required: true},
	{Name: "JobArn", Flag: "job-arn", Type: "*string", Required: true},
	{Name: "PlatformId", Flag: "platform-id", Type: "*string", Required: true},
	{Name: "ProfileVersionArn", Flag: "profile-version-arn", Type: "*string", Required: true},
	{Name: "SignatureTimestamp", Flag: "signature-timestamp", Type: "*time.Time", Required: true},
}

var fields_get_signing_platform = []leanruntime.Field{
	{Name: "PlatformId", Flag: "platform-id", Type: "*string", Required: true},
}

var fields_get_signing_profile = []leanruntime.Field{
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "ProfileOwner", Flag: "profile-owner", Type: "*string", Required: false},
}

var fields_list_profile_permissions = []leanruntime.Field{
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
}

var fields_list_signing_jobs = []leanruntime.Field{
	{Name: "IsRevoked", Flag: "is-revoked", Type: "bool", Required: false},
	{Name: "JobInvoker", Flag: "job-invoker", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlatformId", Flag: "platform-id", Type: "*string", Required: false},
	{Name: "RequestedBy", Flag: "requested-by", Type: "*string", Required: false},
	{Name: "SignatureExpiresAfter", Flag: "signature-expires-after", Type: "*time.Time", Required: false},
	{Name: "SignatureExpiresBefore", Flag: "signature-expires-before", Type: "*time.Time", Required: false},
	{Name: "Status", Flag: "status", Type: "types.SigningStatus", Required: false},
}

var fields_list_signing_platforms = []leanruntime.Field{
	{Name: "Category", Flag: "category", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Partner", Flag: "partner", Type: "*string", Required: false},
	{Name: "Target", Flag: "target", Type: "*string", Required: false},
}

var fields_list_signing_profiles = []leanruntime.Field{
	{Name: "IncludeCanceled", Flag: "include-canceled", Type: "bool", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "PlatformId", Flag: "platform-id", Type: "*string", Required: false},
	{Name: "Statuses", Flag: "statuses", Type: "[]types.SigningProfileStatus", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_signing_profile = []leanruntime.Field{
	{Name: "Overrides", Flag: "overrides", Type: "*types.SigningPlatformOverrides", Required: false},
	{Name: "PlatformId", Flag: "platform-id", Type: "*string", Required: true},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "SignatureValidityPeriod", Flag: "signature-validity-period", Type: "*types.SignatureValidityPeriod", Required: false},
	{Name: "SigningMaterial", Flag: "signing-material", Type: "*types.SigningMaterial", Required: false},
	{Name: "SigningParameters", Flag: "signing-parameters", Type: "map[string]string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_remove_profile_permission = []leanruntime.Field{
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "RevisionId", Flag: "revision-id", Type: "*string", Required: true},
	{Name: "StatementId", Flag: "statement-id", Type: "*string", Required: true},
}

var fields_revoke_signature = []leanruntime.Field{
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
	{Name: "JobOwner", Flag: "job-owner", Type: "*string", Required: false},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_revoke_signing_profile = []leanruntime.Field{
	{Name: "EffectiveTime", Flag: "effective-time", Type: "*time.Time", Required: true},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "ProfileVersion", Flag: "profile-version", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: true},
}

var fields_sign_payload = []leanruntime.Field{
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: true},
	{Name: "PayloadFormat", Flag: "payload-format", Type: "*string", Required: true},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "ProfileOwner", Flag: "profile-owner", Type: "*string", Required: false},
}

var fields_start_signing_job = []leanruntime.Field{
	{Name: "ClientRequestToken", Flag: "client-request-token", Type: "*string", Required: true},
	{Name: "Destination", Flag: "destination", Type: "*types.Destination", Required: true},
	{Name: "ProfileName", Flag: "profile-name", Type: "*string", Required: true},
	{Name: "ProfileOwner", Flag: "profile-owner", Type: "*string", Required: false},
	{Name: "Source", Flag: "source", Type: "*types.Source", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"add-profile-permission": {
			Name:   "add-profile-permission",
			Fields: fields_add_profile_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AddProfilePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_add_profile_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AddProfilePermission(ctx, input)
			},
		},
		"cancel-signing-profile": {
			Name:   "cancel-signing-profile",
			Fields: fields_cancel_signing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CancelSigningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_cancel_signing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CancelSigningProfile(ctx, input)
			},
		},
		"describe-signing-job": {
			Name:   "describe-signing-job",
			Fields: fields_describe_signing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSigningJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_signing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSigningJob(ctx, input)
			},
		},
		"get-revocation-status": {
			Name:   "get-revocation-status",
			Fields: fields_get_revocation_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRevocationStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_revocation_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRevocationStatus(ctx, input)
			},
		},
		"get-signing-platform": {
			Name:   "get-signing-platform",
			Fields: fields_get_signing_platform,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSigningPlatformInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signing_platform, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSigningPlatform(ctx, input)
			},
		},
		"get-signing-profile": {
			Name:   "get-signing-profile",
			Fields: fields_get_signing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSigningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSigningProfile(ctx, input)
			},
		},
		"list-profile-permissions": {
			Name:   "list-profile-permissions",
			Fields: fields_list_profile_permissions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListProfilePermissionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_list_profile_permissions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ListProfilePermissions(ctx, input)
			},
		},
		"list-signing-jobs": {
			Name:   "list-signing-jobs",
			Fields: fields_list_signing_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSigningJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signing_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSigningJobs(ctx, input)
				}
				var results []*svc.ListSigningJobsOutput
				p := svc.NewListSigningJobsPaginator(client, input)
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
		"list-signing-platforms": {
			Name:   "list-signing-platforms",
			Fields: fields_list_signing_platforms,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSigningPlatformsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signing_platforms, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSigningPlatforms(ctx, input)
				}
				var results []*svc.ListSigningPlatformsOutput
				p := svc.NewListSigningPlatformsPaginator(client, input)
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
		"list-signing-profiles": {
			Name:   "list-signing-profiles",
			Fields: fields_list_signing_profiles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSigningProfilesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signing_profiles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSigningProfiles(ctx, input)
				}
				var results []*svc.ListSigningProfilesOutput
				p := svc.NewListSigningProfilesPaginator(client, input)
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
		"put-signing-profile": {
			Name:   "put-signing-profile",
			Fields: fields_put_signing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSigningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_signing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSigningProfile(ctx, input)
			},
		},
		"remove-profile-permission": {
			Name:   "remove-profile-permission",
			Fields: fields_remove_profile_permission,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RemoveProfilePermissionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_remove_profile_permission, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RemoveProfilePermission(ctx, input)
			},
		},
		"revoke-signature": {
			Name:   "revoke-signature",
			Fields: fields_revoke_signature,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeSignatureInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_signature, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeSignature(ctx, input)
			},
		},
		"revoke-signing-profile": {
			Name:   "revoke-signing-profile",
			Fields: fields_revoke_signing_profile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RevokeSigningProfileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_revoke_signing_profile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RevokeSigningProfile(ctx, input)
			},
		},
		"sign-payload": {
			Name:   "sign-payload",
			Fields: fields_sign_payload,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SignPayloadInput{}
				if _, err := leanruntime.ApplyInput(input, fields_sign_payload, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SignPayload(ctx, input)
			},
		},
		"start-signing-job": {
			Name:   "start-signing-job",
			Fields: fields_start_signing_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSigningJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_signing_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSigningJob(ctx, input)
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
	}
	if err := leanruntime.Execute("signer", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
