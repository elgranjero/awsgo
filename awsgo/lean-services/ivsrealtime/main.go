package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/ivsrealtime"
)

var fields_create_encoder_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "Video", Flag: "video", Type: "*types.Video", Required: false},
}

var fields_create_ingest_configuration = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "IngestProtocol", Flag: "ingest-protocol", Type: "types.IngestProtocol", Required: true},
	{Name: "InsecureIngest", Flag: "insecure-ingest", Type: "bool", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_create_participant_token = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "Capabilities", Flag: "capabilities", Type: "[]types.ParticipantTokenCapability", Required: false},
	{Name: "Duration", Flag: "duration", Type: "*int32", Required: false},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: false},
}

var fields_create_stage = []leanruntime.Field{
	{Name: "AutoParticipantRecordingConfiguration", Flag: "auto-participant-recording-configuration", Type: "*types.AutoParticipantRecordingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "ParticipantTokenConfigurations", Flag: "participant-token-configurations", Type: "[]types.ParticipantTokenConfiguration", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_storage_configuration = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "S3", Flag: "s3", Type: "*types.S3StorageConfiguration", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_delete_encoder_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_ingest_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "Force", Flag: "force", Type: "bool", Required: false},
}

var fields_delete_public_key = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_stage = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_delete_storage_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_disconnect_participant = []leanruntime.Field{
	{Name: "ParticipantId", Flag: "participant-id", Type: "*string", Required: true},
	{Name: "Reason", Flag: "reason", Type: "*string", Required: false},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
}

var fields_get_composition = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_encoder_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_ingest_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_participant = []leanruntime.Field{
	{Name: "ParticipantId", Flag: "participant-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
}

var fields_get_public_key = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_stage = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_get_stage_session = []leanruntime.Field{
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
}

var fields_get_storage_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_import_public_key = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PublicKeyMaterial", Flag: "public-key-material", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_list_compositions = []leanruntime.Field{
	{Name: "FilterByEncoderConfigurationArn", Flag: "filter-by-encoder-configuration-arn", Type: "*string", Required: false},
	{Name: "FilterByStageArn", Flag: "filter-by-stage-arn", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_encoder_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_ingest_configurations = []leanruntime.Field{
	{Name: "FilterByStageArn", Flag: "filter-by-stage-arn", Type: "*string", Required: false},
	{Name: "FilterByState", Flag: "filter-by-state", Type: "types.IngestConfigurationState", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_participant_events = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParticipantId", Flag: "participant-id", Type: "*string", Required: true},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
}

var fields_list_participant_replicas = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ParticipantId", Flag: "participant-id", Type: "*string", Required: true},
	{Name: "SourceStageArn", Flag: "source-stage-arn", Type: "*string", Required: true},
}

var fields_list_participants = []leanruntime.Field{
	{Name: "FilterByPublished", Flag: "filter-by-published", Type: "bool", Required: false},
	{Name: "FilterByRecordingState", Flag: "filter-by-recording-state", Type: "types.ParticipantRecordingFilterByRecordingState", Required: false},
	{Name: "FilterByState", Flag: "filter-by-state", Type: "types.ParticipantState", Required: false},
	{Name: "FilterByUserId", Flag: "filter-by-user-id", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SessionId", Flag: "session-id", Type: "*string", Required: true},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
}

var fields_list_public_keys = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_stage_sessions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
}

var fields_list_stages = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_storage_configurations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_start_composition = []leanruntime.Field{
	{Name: "Destinations", Flag: "destinations", Type: "[]types.DestinationConfiguration", Required: true},
	{Name: "IdempotencyToken", Flag: "idempotency-token", Type: "*string", Required: false},
	{Name: "Layout", Flag: "layout", Type: "*types.LayoutConfiguration", Required: false},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_start_participant_replication = []leanruntime.Field{
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "DestinationStageArn", Flag: "destination-stage-arn", Type: "*string", Required: true},
	{Name: "ParticipantId", Flag: "participant-id", Type: "*string", Required: true},
	{Name: "ReconnectWindowSeconds", Flag: "reconnect-window-seconds", Type: "*int32", Required: false},
	{Name: "SourceStageArn", Flag: "source-stage-arn", Type: "*string", Required: true},
}

var fields_stop_composition = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
}

var fields_stop_participant_replication = []leanruntime.Field{
	{Name: "DestinationStageArn", Flag: "destination-stage-arn", Type: "*string", Required: true},
	{Name: "ParticipantId", Flag: "participant-id", Type: "*string", Required: true},
	{Name: "SourceStageArn", Flag: "source-stage-arn", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_ingest_configuration = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "StageArn", Flag: "stage-arn", Type: "*string", Required: false},
}

var fields_update_stage = []leanruntime.Field{
	{Name: "Arn", Flag: "arn", Type: "*string", Required: true},
	{Name: "AutoParticipantRecordingConfiguration", Flag: "auto-participant-recording-configuration", Type: "*types.AutoParticipantRecordingConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-encoder-configuration": {
			Name:   "create-encoder-configuration",
			Fields: fields_create_encoder_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateEncoderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_encoder_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateEncoderConfiguration(ctx, input)
			},
		},
		"create-ingest-configuration": {
			Name:   "create-ingest-configuration",
			Fields: fields_create_ingest_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIngestConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_ingest_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIngestConfiguration(ctx, input)
			},
		},
		"create-participant-token": {
			Name:   "create-participant-token",
			Fields: fields_create_participant_token,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateParticipantTokenInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_participant_token, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateParticipantToken(ctx, input)
			},
		},
		"create-stage": {
			Name:   "create-stage",
			Fields: fields_create_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStage(ctx, input)
			},
		},
		"create-storage-configuration": {
			Name:   "create-storage-configuration",
			Fields: fields_create_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStorageConfiguration(ctx, input)
			},
		},
		"delete-encoder-configuration": {
			Name:   "delete-encoder-configuration",
			Fields: fields_delete_encoder_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteEncoderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_encoder_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteEncoderConfiguration(ctx, input)
			},
		},
		"delete-ingest-configuration": {
			Name:   "delete-ingest-configuration",
			Fields: fields_delete_ingest_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIngestConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_ingest_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIngestConfiguration(ctx, input)
			},
		},
		"delete-public-key": {
			Name:   "delete-public-key",
			Fields: fields_delete_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePublicKey(ctx, input)
			},
		},
		"delete-stage": {
			Name:   "delete-stage",
			Fields: fields_delete_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStage(ctx, input)
			},
		},
		"delete-storage-configuration": {
			Name:   "delete-storage-configuration",
			Fields: fields_delete_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStorageConfiguration(ctx, input)
			},
		},
		"disconnect-participant": {
			Name:   "disconnect-participant",
			Fields: fields_disconnect_participant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisconnectParticipantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disconnect_participant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisconnectParticipant(ctx, input)
			},
		},
		"get-composition": {
			Name:   "get-composition",
			Fields: fields_get_composition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCompositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_composition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetComposition(ctx, input)
			},
		},
		"get-encoder-configuration": {
			Name:   "get-encoder-configuration",
			Fields: fields_get_encoder_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEncoderConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_encoder_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEncoderConfiguration(ctx, input)
			},
		},
		"get-ingest-configuration": {
			Name:   "get-ingest-configuration",
			Fields: fields_get_ingest_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIngestConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_ingest_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIngestConfiguration(ctx, input)
			},
		},
		"get-participant": {
			Name:   "get-participant",
			Fields: fields_get_participant,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetParticipantInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_participant, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetParticipant(ctx, input)
			},
		},
		"get-public-key": {
			Name:   "get-public-key",
			Fields: fields_get_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPublicKey(ctx, input)
			},
		},
		"get-stage": {
			Name:   "get-stage",
			Fields: fields_get_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStage(ctx, input)
			},
		},
		"get-stage-session": {
			Name:   "get-stage-session",
			Fields: fields_get_stage_session,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStageSessionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_stage_session, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStageSession(ctx, input)
			},
		},
		"get-storage-configuration": {
			Name:   "get-storage-configuration",
			Fields: fields_get_storage_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStorageConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_storage_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStorageConfiguration(ctx, input)
			},
		},
		"import-public-key": {
			Name:   "import-public-key",
			Fields: fields_import_public_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportPublicKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_public_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportPublicKey(ctx, input)
			},
		},
		"list-compositions": {
			Name:   "list-compositions",
			Fields: fields_list_compositions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCompositionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_compositions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCompositions(ctx, input)
				}
				var results []*svc.ListCompositionsOutput
				p := svc.NewListCompositionsPaginator(client, input)
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
		"list-encoder-configurations": {
			Name:   "list-encoder-configurations",
			Fields: fields_list_encoder_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListEncoderConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_encoder_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListEncoderConfigurations(ctx, input)
				}
				var results []*svc.ListEncoderConfigurationsOutput
				p := svc.NewListEncoderConfigurationsPaginator(client, input)
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
		"list-ingest-configurations": {
			Name:   "list-ingest-configurations",
			Fields: fields_list_ingest_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListIngestConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_ingest_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListIngestConfigurations(ctx, input)
				}
				var results []*svc.ListIngestConfigurationsOutput
				p := svc.NewListIngestConfigurationsPaginator(client, input)
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
		"list-participant-events": {
			Name:   "list-participant-events",
			Fields: fields_list_participant_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListParticipantEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_participant_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParticipantEvents(ctx, input)
				}
				var results []*svc.ListParticipantEventsOutput
				p := svc.NewListParticipantEventsPaginator(client, input)
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
		"list-participant-replicas": {
			Name:   "list-participant-replicas",
			Fields: fields_list_participant_replicas,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListParticipantReplicasInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_participant_replicas, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParticipantReplicas(ctx, input)
				}
				var results []*svc.ListParticipantReplicasOutput
				p := svc.NewListParticipantReplicasPaginator(client, input)
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
		"list-participants": {
			Name:   "list-participants",
			Fields: fields_list_participants,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListParticipantsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_participants, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListParticipants(ctx, input)
				}
				var results []*svc.ListParticipantsOutput
				p := svc.NewListParticipantsPaginator(client, input)
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
		"list-public-keys": {
			Name:   "list-public-keys",
			Fields: fields_list_public_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPublicKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_public_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPublicKeys(ctx, input)
				}
				var results []*svc.ListPublicKeysOutput
				p := svc.NewListPublicKeysPaginator(client, input)
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
		"list-stage-sessions": {
			Name:   "list-stage-sessions",
			Fields: fields_list_stage_sessions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStageSessionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stage_sessions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStageSessions(ctx, input)
				}
				var results []*svc.ListStageSessionsOutput
				p := svc.NewListStageSessionsPaginator(client, input)
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
		"list-stages": {
			Name:   "list-stages",
			Fields: fields_list_stages,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStagesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_stages, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStages(ctx, input)
				}
				var results []*svc.ListStagesOutput
				p := svc.NewListStagesPaginator(client, input)
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
		"list-storage-configurations": {
			Name:   "list-storage-configurations",
			Fields: fields_list_storage_configurations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStorageConfigurationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_storage_configurations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStorageConfigurations(ctx, input)
				}
				var results []*svc.ListStorageConfigurationsOutput
				p := svc.NewListStorageConfigurationsPaginator(client, input)
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
		"start-composition": {
			Name:   "start-composition",
			Fields: fields_start_composition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartCompositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_composition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartComposition(ctx, input)
			},
		},
		"start-participant-replication": {
			Name:   "start-participant-replication",
			Fields: fields_start_participant_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartParticipantReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_participant_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartParticipantReplication(ctx, input)
			},
		},
		"stop-composition": {
			Name:   "stop-composition",
			Fields: fields_stop_composition,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopCompositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_composition, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopComposition(ctx, input)
			},
		},
		"stop-participant-replication": {
			Name:   "stop-participant-replication",
			Fields: fields_stop_participant_replication,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StopParticipantReplicationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_stop_participant_replication, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StopParticipantReplication(ctx, input)
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
		"update-ingest-configuration": {
			Name:   "update-ingest-configuration",
			Fields: fields_update_ingest_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateIngestConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_ingest_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateIngestConfiguration(ctx, input)
			},
		},
		"update-stage": {
			Name:   "update-stage",
			Fields: fields_update_stage,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStageInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_stage, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStage(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("ivsrealtime", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
