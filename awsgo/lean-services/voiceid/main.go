package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/voiceid"
)

var fields_associate_fraudster = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FraudsterId", Flag: "fraudster-id", Type: "*string", Required: true},
	{Name: "WatchlistId", Flag: "watchlist-id", Type: "*string", Required: true},
}

var fields_create_domain = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_watchlist = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_domain = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_delete_fraudster = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FraudsterId", Flag: "fraudster-id", Type: "*string", Required: true},
}

var fields_delete_speaker = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpeakerId", Flag: "speaker-id", Type: "*string", Required: true},
}

var fields_delete_watchlist = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "WatchlistId", Flag: "watchlist-id", Type: "*string", Required: true},
}

var fields_describe_domain = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
}

var fields_describe_fraudster = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FraudsterId", Flag: "fraudster-id", Type: "*string", Required: true},
}

var fields_describe_fraudster_registration_job = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_speaker = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpeakerId", Flag: "speaker-id", Type: "*string", Required: true},
}

var fields_describe_speaker_enrollment_job = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "JobId", Flag: "job-id", Type: "*string", Required: true},
}

var fields_describe_watchlist = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "WatchlistId", Flag: "watchlist-id", Type: "*string", Required: true},
}

var fields_disassociate_fraudster = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "FraudsterId", Flag: "fraudster-id", Type: "*string", Required: true},
	{Name: "WatchlistId", Flag: "watchlist-id", Type: "*string", Required: true},
}

var fields_evaluate_session = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SessionNameOrId", Flag: "session-name-or-id", Type: "*string", Required: true},
}

var fields_list_domains = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fraudster_registration_jobs = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "JobStatus", Flag: "job-status", Type: "types.FraudsterRegistrationJobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fraudsters = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "WatchlistId", Flag: "watchlist-id", Type: "*string", Required: false},
}

var fields_list_speaker_enrollment_jobs = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "JobStatus", Flag: "job-status", Type: "types.SpeakerEnrollmentJobStatus", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_speakers = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_watchlists = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_opt_out_speaker = []leanruntime.Field{
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "SpeakerId", Flag: "speaker-id", Type: "*string", Required: true},
}

var fields_start_fraudster_registration_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
	{Name: "RegistrationConfig", Flag: "registration-config", Type: "*types.RegistrationConfig", Required: false},
}

var fields_start_speaker_enrollment_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "DataAccessRoleArn", Flag: "data-access-role-arn", Type: "*string", Required: true},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "EnrollmentConfig", Flag: "enrollment-config", Type: "*types.EnrollmentConfig", Required: false},
	{Name: "InputDataConfig", Flag: "input-data-config", Type: "*types.InputDataConfig", Required: true},
	{Name: "JobName", Flag: "job-name", Type: "*string", Required: false},
	{Name: "OutputDataConfig", Flag: "output-data-config", Type: "*types.OutputDataConfig", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_domain = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ServerSideEncryptionConfiguration", Flag: "server-side-encryption-configuration", Type: "*types.ServerSideEncryptionConfiguration", Required: true},
}

var fields_update_watchlist = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DomainId", Flag: "domain-id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "WatchlistId", Flag: "watchlist-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-fraudster": {
			Name:   "associate-fraudster",
			Fields: fields_associate_fraudster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateFraudsterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_fraudster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateFraudster(ctx, input)
			},
		},
		"create-domain": {
			Name:   "create-domain",
			Fields: fields_create_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDomain(ctx, input)
			},
		},
		"create-watchlist": {
			Name:   "create-watchlist",
			Fields: fields_create_watchlist,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateWatchlistInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_watchlist, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateWatchlist(ctx, input)
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
		"delete-fraudster": {
			Name:   "delete-fraudster",
			Fields: fields_delete_fraudster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFraudsterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fraudster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFraudster(ctx, input)
			},
		},
		"delete-speaker": {
			Name:   "delete-speaker",
			Fields: fields_delete_speaker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSpeakerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_speaker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSpeaker(ctx, input)
			},
		},
		"delete-watchlist": {
			Name:   "delete-watchlist",
			Fields: fields_delete_watchlist,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteWatchlistInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_watchlist, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteWatchlist(ctx, input)
			},
		},
		"describe-domain": {
			Name:   "describe-domain",
			Fields: fields_describe_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeDomain(ctx, input)
			},
		},
		"describe-fraudster": {
			Name:   "describe-fraudster",
			Fields: fields_describe_fraudster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFraudsterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fraudster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFraudster(ctx, input)
			},
		},
		"describe-fraudster-registration-job": {
			Name:   "describe-fraudster-registration-job",
			Fields: fields_describe_fraudster_registration_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeFraudsterRegistrationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_fraudster_registration_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeFraudsterRegistrationJob(ctx, input)
			},
		},
		"describe-speaker": {
			Name:   "describe-speaker",
			Fields: fields_describe_speaker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpeakerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_speaker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSpeaker(ctx, input)
			},
		},
		"describe-speaker-enrollment-job": {
			Name:   "describe-speaker-enrollment-job",
			Fields: fields_describe_speaker_enrollment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeSpeakerEnrollmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_speaker_enrollment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeSpeakerEnrollmentJob(ctx, input)
			},
		},
		"describe-watchlist": {
			Name:   "describe-watchlist",
			Fields: fields_describe_watchlist,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeWatchlistInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_watchlist, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeWatchlist(ctx, input)
			},
		},
		"disassociate-fraudster": {
			Name:   "disassociate-fraudster",
			Fields: fields_disassociate_fraudster,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateFraudsterInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_fraudster, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateFraudster(ctx, input)
			},
		},
		"evaluate-session": {
			Name:   "evaluate-session",
			Fields: fields_evaluate_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.EvaluateSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_evaluate_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.EvaluateSession(ctx, input)
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
		"list-fraudster-registration-jobs": {
			Name:   "list-fraudster-registration-jobs",
			Fields: fields_list_fraudster_registration_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFraudsterRegistrationJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fraudster_registration_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFraudsterRegistrationJobs(ctx, input)
				}
				var results []*svc.ListFraudsterRegistrationJobsOutput
				p := svc.NewListFraudsterRegistrationJobsPaginator(client, input)
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
		"list-fraudsters": {
			Name:   "list-fraudsters",
			Fields: fields_list_fraudsters,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFraudstersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fraudsters, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFraudsters(ctx, input)
				}
				var results []*svc.ListFraudstersOutput
				p := svc.NewListFraudstersPaginator(client, input)
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
		"list-speaker-enrollment-jobs": {
			Name:   "list-speaker-enrollment-jobs",
			Fields: fields_list_speaker_enrollment_jobs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpeakerEnrollmentJobsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_speaker_enrollment_jobs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpeakerEnrollmentJobs(ctx, input)
				}
				var results []*svc.ListSpeakerEnrollmentJobsOutput
				p := svc.NewListSpeakerEnrollmentJobsPaginator(client, input)
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
		"list-speakers": {
			Name:   "list-speakers",
			Fields: fields_list_speakers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSpeakersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_speakers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSpeakers(ctx, input)
				}
				var results []*svc.ListSpeakersOutput
				p := svc.NewListSpeakersPaginator(client, input)
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
		"list-watchlists": {
			Name:   "list-watchlists",
			Fields: fields_list_watchlists,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListWatchlistsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_watchlists, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListWatchlists(ctx, input)
				}
				var results []*svc.ListWatchlistsOutput
				p := svc.NewListWatchlistsPaginator(client, input)
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
		"opt-out-speaker": {
			Name:   "opt-out-speaker",
			Fields: fields_opt_out_speaker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OptOutSpeakerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_opt_out_speaker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OptOutSpeaker(ctx, input)
			},
		},
		"start-fraudster-registration-job": {
			Name:   "start-fraudster-registration-job",
			Fields: fields_start_fraudster_registration_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartFraudsterRegistrationJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_fraudster_registration_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartFraudsterRegistrationJob(ctx, input)
			},
		},
		"start-speaker-enrollment-job": {
			Name:   "start-speaker-enrollment-job",
			Fields: fields_start_speaker_enrollment_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartSpeakerEnrollmentJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_speaker_enrollment_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartSpeakerEnrollmentJob(ctx, input)
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
		"update-domain": {
			Name:   "update-domain",
			Fields: fields_update_domain,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDomainInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_domain, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDomain(ctx, input)
			},
		},
		"update-watchlist": {
			Name:   "update-watchlist",
			Fields: fields_update_watchlist,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateWatchlistInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_watchlist, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateWatchlist(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("voiceid", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
