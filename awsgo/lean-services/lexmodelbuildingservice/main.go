package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/lexmodelbuildingservice"
)

var fields_create_bot_version = []leanruntime.Field{
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_intent_version = []leanruntime.Field{
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_create_slot_type_version = []leanruntime.Field{
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_bot = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_bot_alias = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_bot_channel_association = []leanruntime.Field{
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_bot_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_delete_intent = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_intent_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_delete_slot_type = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_slot_type_version = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_delete_utterances = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "UserId", Flag: "user-id", Type: "*string", Required: true},
}

var fields_get_bot = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "VersionOrAlias", Flag: "version-or-alias", Type: "*string", Required: true},
}

var fields_get_bot_alias = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_bot_aliases = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_bot_channel_association = []leanruntime.Field{
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_bot_channel_associations = []leanruntime.Field{
	{Name: "BotAlias", Flag: "bot-alias", Type: "*string", Required: true},
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_bot_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_bots = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_builtin_intent = []leanruntime.Field{
	{Name: "Signature", Flag: "signature", Type: "*string", Required: true},
}

var fields_get_builtin_intents = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SignatureContains", Flag: "signature-contains", Type: "*string", Required: false},
}

var fields_get_builtin_slot_types = []leanruntime.Field{
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SignatureContains", Flag: "signature-contains", Type: "*string", Required: false},
}

var fields_get_export = []leanruntime.Field{
	{Name: "ExportType", Flag: "export-type", Type: "types.ExportType", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_get_import = []leanruntime.Field{
	{Name: "ImportId", Flag: "import-id", Type: "*string", Required: true},
}

var fields_get_intent = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_get_intent_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_intents = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_migration = []leanruntime.Field{
	{Name: "MigrationId", Flag: "migration-id", Type: "*string", Required: true},
}

var fields_get_migrations = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "MigrationStatusEquals", Flag: "migration-status-equals", Type: "types.MigrationStatus", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SortByAttribute", Flag: "sort-by-attribute", Type: "types.MigrationSortAttribute", Required: false},
	{Name: "SortByOrder", Flag: "sort-by-order", Type: "types.SortOrder", Required: false},
	{Name: "V1BotNameContains", Flag: "v1-bot-name-contains", Type: "*string", Required: false},
}

var fields_get_slot_type = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Version", Flag: "version", Type: "*string", Required: true},
}

var fields_get_slot_type_versions = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_slot_types = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NameContains", Flag: "name-contains", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_get_utterances_view = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "BotVersions", Flag: "bot-versions", Type: "[]string", Required: true},
	{Name: "StatusType", Flag: "status-type", Type: "types.StatusType", Required: true},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_put_bot = []leanruntime.Field{
	{Name: "AbortStatement", Flag: "abort-statement", Type: "*types.Statement", Required: false},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "ChildDirected", Flag: "child-directed", Type: "*bool", Required: true},
	{Name: "ClarificationPrompt", Flag: "clarification-prompt", Type: "*types.Prompt", Required: false},
	{Name: "CreateVersion", Flag: "create-version", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DetectSentiment", Flag: "detect-sentiment", Type: "*bool", Required: false},
	{Name: "EnableModelImprovements", Flag: "enable-model-improvements", Type: "*bool", Required: false},
	{Name: "IdleSessionTTLInSeconds", Flag: "idle-session-ttlin-seconds", Type: "*int32", Required: false},
	{Name: "Intents", Flag: "intents", Type: "[]types.Intent", Required: false},
	{Name: "Locale", Flag: "locale", Type: "types.Locale", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NluIntentConfidenceThreshold", Flag: "nlu-intent-confidence-threshold", Type: "*float64", Required: false},
	{Name: "ProcessBehavior", Flag: "process-behavior", Type: "types.ProcessBehavior", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VoiceId", Flag: "voice-id", Type: "*string", Required: false},
}

var fields_put_bot_alias = []leanruntime.Field{
	{Name: "BotName", Flag: "bot-name", Type: "*string", Required: true},
	{Name: "BotVersion", Flag: "bot-version", Type: "*string", Required: true},
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "ConversationLogs", Flag: "conversation-logs", Type: "*types.ConversationLogsRequest", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_put_intent = []leanruntime.Field{
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "ConclusionStatement", Flag: "conclusion-statement", Type: "*types.Statement", Required: false},
	{Name: "ConfirmationPrompt", Flag: "confirmation-prompt", Type: "*types.Prompt", Required: false},
	{Name: "CreateVersion", Flag: "create-version", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DialogCodeHook", Flag: "dialog-code-hook", Type: "*types.CodeHook", Required: false},
	{Name: "FollowUpPrompt", Flag: "follow-up-prompt", Type: "*types.FollowUpPrompt", Required: false},
	{Name: "FulfillmentActivity", Flag: "fulfillment-activity", Type: "*types.FulfillmentActivity", Required: false},
	{Name: "InputContexts", Flag: "input-contexts", Type: "[]types.InputContext", Required: false},
	{Name: "KendraConfiguration", Flag: "kendra-configuration", Type: "*types.KendraConfiguration", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputContexts", Flag: "output-contexts", Type: "[]types.OutputContext", Required: false},
	{Name: "ParentIntentSignature", Flag: "parent-intent-signature", Type: "*string", Required: false},
	{Name: "RejectionStatement", Flag: "rejection-statement", Type: "*types.Statement", Required: false},
	{Name: "SampleUtterances", Flag: "sample-utterances", Type: "[]string", Required: false},
	{Name: "Slots", Flag: "slots", Type: "[]types.Slot", Required: false},
}

var fields_put_slot_type = []leanruntime.Field{
	{Name: "Checksum", Flag: "checksum", Type: "*string", Required: false},
	{Name: "CreateVersion", Flag: "create-version", Type: "*bool", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EnumerationValues", Flag: "enumeration-values", Type: "[]types.EnumerationValue", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "ParentSlotTypeSignature", Flag: "parent-slot-type-signature", Type: "*string", Required: false},
	{Name: "SlotTypeConfigurations", Flag: "slot-type-configurations", Type: "[]types.SlotTypeConfiguration", Required: false},
	{Name: "ValueSelectionStrategy", Flag: "value-selection-strategy", Type: "types.SlotValueSelectionStrategy", Required: false},
}

var fields_start_import = []leanruntime.Field{
	{Name: "MergeStrategy", Flag: "merge-strategy", Type: "types.MergeStrategy", Required: true},
	{Name: "Payload", Flag: "payload", Type: "[]byte", Required: true},
	{Name: "ResourceType", Flag: "resource-type", Type: "types.ResourceType", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_start_migration = []leanruntime.Field{
	{Name: "MigrationStrategy", Flag: "migration-strategy", Type: "types.MigrationStrategy", Required: true},
	{Name: "V1BotName", Flag: "v1-bot-name", Type: "*string", Required: true},
	{Name: "V1BotVersion", Flag: "v1-bot-version", Type: "*string", Required: true},
	{Name: "V2BotName", Flag: "v2-bot-name", Type: "*string", Required: true},
	{Name: "V2BotRole", Flag: "v2-bot-role", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-bot-version": {
			Name:   "create-bot-version",
			Fields: fields_create_bot_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateBotVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_bot_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateBotVersion(ctx, input)
			},
		},
		"create-intent-version": {
			Name:   "create-intent-version",
			Fields: fields_create_intent_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateIntentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_intent_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateIntentVersion(ctx, input)
			},
		},
		"create-slot-type-version": {
			Name:   "create-slot-type-version",
			Fields: fields_create_slot_type_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSlotTypeVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_slot_type_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSlotTypeVersion(ctx, input)
			},
		},
		"delete-bot": {
			Name:   "delete-bot",
			Fields: fields_delete_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBot(ctx, input)
			},
		},
		"delete-bot-alias": {
			Name:   "delete-bot-alias",
			Fields: fields_delete_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotAlias(ctx, input)
			},
		},
		"delete-bot-channel-association": {
			Name:   "delete-bot-channel-association",
			Fields: fields_delete_bot_channel_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotChannelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_channel_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotChannelAssociation(ctx, input)
			},
		},
		"delete-bot-version": {
			Name:   "delete-bot-version",
			Fields: fields_delete_bot_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteBotVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_bot_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteBotVersion(ctx, input)
			},
		},
		"delete-intent": {
			Name:   "delete-intent",
			Fields: fields_delete_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntent(ctx, input)
			},
		},
		"delete-intent-version": {
			Name:   "delete-intent-version",
			Fields: fields_delete_intent_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteIntentVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_intent_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteIntentVersion(ctx, input)
			},
		},
		"delete-slot-type": {
			Name:   "delete-slot-type",
			Fields: fields_delete_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlotType(ctx, input)
			},
		},
		"delete-slot-type-version": {
			Name:   "delete-slot-type-version",
			Fields: fields_delete_slot_type_version,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSlotTypeVersionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_slot_type_version, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSlotTypeVersion(ctx, input)
			},
		},
		"delete-utterances": {
			Name:   "delete-utterances",
			Fields: fields_delete_utterances,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteUtterancesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_utterances, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteUtterances(ctx, input)
			},
		},
		"get-bot": {
			Name:   "get-bot",
			Fields: fields_get_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBot(ctx, input)
			},
		},
		"get-bot-alias": {
			Name:   "get-bot-alias",
			Fields: fields_get_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBotAlias(ctx, input)
			},
		},
		"get-bot-aliases": {
			Name:   "get-bot-aliases",
			Fields: fields_get_bot_aliases,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotAliasesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_bot_aliases, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBotAliases(ctx, input)
				}
				var results []*svc.GetBotAliasesOutput
				p := svc.NewGetBotAliasesPaginator(client, input)
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
		"get-bot-channel-association": {
			Name:   "get-bot-channel-association",
			Fields: fields_get_bot_channel_association,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotChannelAssociationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_bot_channel_association, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBotChannelAssociation(ctx, input)
			},
		},
		"get-bot-channel-associations": {
			Name:   "get-bot-channel-associations",
			Fields: fields_get_bot_channel_associations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotChannelAssociationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_bot_channel_associations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBotChannelAssociations(ctx, input)
				}
				var results []*svc.GetBotChannelAssociationsOutput
				p := svc.NewGetBotChannelAssociationsPaginator(client, input)
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
		"get-bot-versions": {
			Name:   "get-bot-versions",
			Fields: fields_get_bot_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_bot_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBotVersions(ctx, input)
				}
				var results []*svc.GetBotVersionsOutput
				p := svc.NewGetBotVersionsPaginator(client, input)
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
		"get-bots": {
			Name:   "get-bots",
			Fields: fields_get_bots,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBotsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_bots, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBots(ctx, input)
				}
				var results []*svc.GetBotsOutput
				p := svc.NewGetBotsPaginator(client, input)
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
		"get-builtin-intent": {
			Name:   "get-builtin-intent",
			Fields: fields_get_builtin_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBuiltinIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_builtin_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetBuiltinIntent(ctx, input)
			},
		},
		"get-builtin-intents": {
			Name:   "get-builtin-intents",
			Fields: fields_get_builtin_intents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBuiltinIntentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_builtin_intents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBuiltinIntents(ctx, input)
				}
				var results []*svc.GetBuiltinIntentsOutput
				p := svc.NewGetBuiltinIntentsPaginator(client, input)
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
		"get-builtin-slot-types": {
			Name:   "get-builtin-slot-types",
			Fields: fields_get_builtin_slot_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetBuiltinSlotTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_builtin_slot_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetBuiltinSlotTypes(ctx, input)
				}
				var results []*svc.GetBuiltinSlotTypesOutput
				p := svc.NewGetBuiltinSlotTypesPaginator(client, input)
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
		"get-export": {
			Name:   "get-export",
			Fields: fields_get_export,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetExportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_export, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetExport(ctx, input)
			},
		},
		"get-import": {
			Name:   "get-import",
			Fields: fields_get_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetImport(ctx, input)
			},
		},
		"get-intent": {
			Name:   "get-intent",
			Fields: fields_get_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetIntent(ctx, input)
			},
		},
		"get-intent-versions": {
			Name:   "get-intent-versions",
			Fields: fields_get_intent_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntentVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_intent_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIntentVersions(ctx, input)
				}
				var results []*svc.GetIntentVersionsOutput
				p := svc.NewGetIntentVersionsPaginator(client, input)
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
		"get-intents": {
			Name:   "get-intents",
			Fields: fields_get_intents,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetIntentsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_intents, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetIntents(ctx, input)
				}
				var results []*svc.GetIntentsOutput
				p := svc.NewGetIntentsPaginator(client, input)
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
		"get-migration": {
			Name:   "get-migration",
			Fields: fields_get_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMigration(ctx, input)
			},
		},
		"get-migrations": {
			Name:   "get-migrations",
			Fields: fields_get_migrations,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMigrationsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_migrations, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetMigrations(ctx, input)
				}
				var results []*svc.GetMigrationsOutput
				p := svc.NewGetMigrationsPaginator(client, input)
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
		"get-slot-type": {
			Name:   "get-slot-type",
			Fields: fields_get_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSlotType(ctx, input)
			},
		},
		"get-slot-type-versions": {
			Name:   "get-slot-type-versions",
			Fields: fields_get_slot_type_versions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSlotTypeVersionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_slot_type_versions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSlotTypeVersions(ctx, input)
				}
				var results []*svc.GetSlotTypeVersionsOutput
				p := svc.NewGetSlotTypeVersionsPaginator(client, input)
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
		"get-slot-types": {
			Name:   "get-slot-types",
			Fields: fields_get_slot_types,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSlotTypesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_slot_types, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetSlotTypes(ctx, input)
				}
				var results []*svc.GetSlotTypesOutput
				p := svc.NewGetSlotTypesPaginator(client, input)
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
		"get-utterances-view": {
			Name:   "get-utterances-view",
			Fields: fields_get_utterances_view,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetUtterancesViewInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_utterances_view, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetUtterancesView(ctx, input)
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
		"put-bot": {
			Name:   "put-bot",
			Fields: fields_put_bot,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBotInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bot, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBot(ctx, input)
			},
		},
		"put-bot-alias": {
			Name:   "put-bot-alias",
			Fields: fields_put_bot_alias,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutBotAliasInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_bot_alias, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutBotAlias(ctx, input)
			},
		},
		"put-intent": {
			Name:   "put-intent",
			Fields: fields_put_intent,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutIntentInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_intent, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutIntent(ctx, input)
			},
		},
		"put-slot-type": {
			Name:   "put-slot-type",
			Fields: fields_put_slot_type,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutSlotTypeInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_slot_type, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutSlotType(ctx, input)
			},
		},
		"start-import": {
			Name:   "start-import",
			Fields: fields_start_import,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartImportInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_import, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartImport(ctx, input)
			},
		},
		"start-migration": {
			Name:   "start-migration",
			Fields: fields_start_migration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartMigrationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_migration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartMigration(ctx, input)
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
	if err := leanruntime.Execute("lexmodelbuildingservice", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
