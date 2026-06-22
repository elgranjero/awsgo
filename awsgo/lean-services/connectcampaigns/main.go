package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/connectcampaigns"
)

var fields_create_campaign = []leanruntime.Field{
	{Name: "ConnectInstanceId", Flag: "connect-instance-id", Type: "*string", Required: true},
	{Name: "DialerConfig", Flag: "dialer-config", Type: "types.DialerConfig", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutboundCallConfig", Flag: "outbound-call-config", Type: "*types.OutboundCallConfig", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_campaign = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_delete_connect_instance_config = []leanruntime.Field{
	{Name: "ConnectInstanceId", Flag: "connect-instance-id", Type: "*string", Required: true},
}

var fields_delete_instance_onboarding_job = []leanruntime.Field{
	{Name: "ConnectInstanceId", Flag: "connect-instance-id", Type: "*string", Required: true},
}

var fields_describe_campaign = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_campaign_state = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_get_campaign_state_batch = []leanruntime.Field{
	{Name: "CampaignIds", Flag: "campaign-ids", Type: "[]string", Required: true},
}

var fields_get_connect_instance_config = []leanruntime.Field{
	{Name: "ConnectInstanceId", Flag: "connect-instance-id", Type: "*string", Required: true},
}

var fields_get_instance_onboarding_job_status = []leanruntime.Field{
	{Name: "ConnectInstanceId", Flag: "connect-instance-id", Type: "*string", Required: true},
}

var fields_list_campaigns = []leanruntime.Field{
	{Name: "Filters", Flag: "filters", Type: "*types.CampaignFilters", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_pause_campaign = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_put_dial_request_batch = []leanruntime.Field{
	{Name: "DialRequests", Flag: "dial-requests", Type: "[]types.DialRequest", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_resume_campaign = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_start_campaign = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_start_instance_onboarding_job = []leanruntime.Field{
	{Name: "ConnectInstanceId", Flag: "connect-instance-id", Type: "*string", Required: true},
	{Name: "EncryptionConfig", Flag: "encryption-config", Type: "*types.EncryptionConfig", Required: true},
}

var fields_stop_campaign = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_campaign_dialer_config = []leanruntime.Field{
	{Name: "DialerConfig", Flag: "dialer-config", Type: "types.DialerConfig", Required: true},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

var fields_update_campaign_name = []leanruntime.Field{
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_campaign_outbound_call_config = []leanruntime.Field{
	{Name: "AnswerMachineDetectionConfig", Flag: "answer-machine-detection-config", Type: "*types.AnswerMachineDetectionConfig", Required: false},
	{Name: "ConnectContactFlowId", Flag: "connect-contact-flow-id", Type: "*string", Required: false},
	{Name: "ConnectSourcePhoneNumber", Flag: "connect-source-phone-number", Type: "*string", Required: false},
	{Name: "Id", Flag: "id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-campaign": {
			Name:   "create-campaign",
			Fields: fields_create_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCampaign(ctx, input)
			},
		},
		"delete-campaign": {
			Name:   "delete-campaign",
			Fields: fields_delete_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCampaign(ctx, input)
			},
		},
		"delete-connect-instance-config": {
			Name:   "delete-connect-instance-config",
			Fields: fields_delete_connect_instance_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteConnectInstanceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_connect_instance_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteConnectInstanceConfig(ctx, input)
			},
		},
		"delete-instance-onboarding-job": {
			Name:   "delete-instance-onboarding-job",
			Fields: fields_delete_instance_onboarding_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteInstanceOnboardingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_instance_onboarding_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteInstanceOnboardingJob(ctx, input)
			},
		},
		"describe-campaign": {
			Name:   "describe-campaign",
			Fields: fields_describe_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeCampaign(ctx, input)
			},
		},
		"get-campaign-state": {
			Name:   "get-campaign-state",
			Fields: fields_get_campaign_state,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignStateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign_state, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaignState(ctx, input)
			},
		},
		"get-campaign-state-batch": {
			Name:   "get-campaign-state-batch",
			Fields: fields_get_campaign_state_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignStateBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign_state_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaignStateBatch(ctx, input)
			},
		},
		"get-connect-instance-config": {
			Name:   "get-connect-instance-config",
			Fields: fields_get_connect_instance_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetConnectInstanceConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_connect_instance_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetConnectInstanceConfig(ctx, input)
			},
		},
		"get-instance-onboarding-job-status": {
			Name:   "get-instance-onboarding-job-status",
			Fields: fields_get_instance_onboarding_job_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetInstanceOnboardingJobStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_instance_onboarding_job_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetInstanceOnboardingJobStatus(ctx, input)
			},
		},
		"list-campaigns": {
			Name:   "list-campaigns",
			Fields: fields_list_campaigns,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCampaignsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_campaigns, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCampaigns(ctx, input)
				}
				var results []*svc.ListCampaignsOutput
				p := svc.NewListCampaignsPaginator(client, input)
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
		"pause-campaign": {
			Name:   "pause-campaign",
			Fields: fields_pause_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PauseCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_pause_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PauseCampaign(ctx, input)
			},
		},
		"put-dial-request-batch": {
			Name:   "put-dial-request-batch",
			Fields: fields_put_dial_request_batch,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutDialRequestBatchInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_dial_request_batch, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutDialRequestBatch(ctx, input)
			},
		},
		"resume-campaign": {
			Name:   "resume-campaign",
			Fields: fields_resume_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ResumeCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_resume_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ResumeCampaign(ctx, input)
			},
		},
		"start-campaign": {
			Name:   "start-campaign",
			Fields: fields_start_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartCampaign(ctx, input)
			},
		},
		"start-instance-onboarding-job": {
			Name:   "start-instance-onboarding-job",
			Fields: fields_start_instance_onboarding_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartInstanceOnboardingJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_instance_onboarding_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartInstanceOnboardingJob(ctx, input)
			},
		},
		"stop-campaign": {
			Name:   "stop-campaign",
			Fields: fields_stop_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopCampaign(ctx, input)
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
		"update-campaign-dialer-config": {
			Name:   "update-campaign-dialer-config",
			Fields: fields_update_campaign_dialer_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCampaignDialerConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_campaign_dialer_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCampaignDialerConfig(ctx, input)
			},
		},
		"update-campaign-name": {
			Name:   "update-campaign-name",
			Fields: fields_update_campaign_name,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCampaignNameInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_campaign_name, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCampaignName(ctx, input)
			},
		},
		"update-campaign-outbound-call-config": {
			Name:   "update-campaign-outbound-call-config",
			Fields: fields_update_campaign_outbound_call_config,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCampaignOutboundCallConfigInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_campaign_outbound_call_config, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCampaignOutboundCallConfig(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("connectcampaigns", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
