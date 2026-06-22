package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/b2bi"
)

var fields_create_capability = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Configuration", Flag: "configuration", Type: "types.CapabilityConfiguration", Required: true},
	{Name: "InstructionsDocuments", Flag: "instructions-documents", Type: "[]types.S3Location", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Type", Flag: "type", Type: "types.CapabilityType", Required: true},
}

var fields_create_partnership = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]string", Required: true},
	{Name: "CapabilityOptions", Flag: "capability-options", Type: "*types.CapabilityOptions", Required: false},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Phone", Flag: "phone", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_profile = []leanruntime.Field{
	{Name: "BusinessName", Flag: "business-name", Type: "*string", Required: true},
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "Logging", Flag: "logging", Type: "types.Logging", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Phone", Flag: "phone", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_starter_mapping_template = []leanruntime.Field{
	{Name: "MappingType", Flag: "mapping-type", Type: "types.MappingType", Required: true},
	{Name: "OutputSampleLocation", Flag: "output-sample-location", Type: "*types.S3Location", Required: false},
	{Name: "TemplateDetails", Flag: "template-details", Type: "types.TemplateDetails", Required: true},
}

var fields_create_transformer = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "EdiType", Flag: "edi-type", Type: "types.EdiType", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "InputConversion", Flag: "input-conversion", Type: "*types.InputConversion", Required: false},
	{Name: "Mapping", Flag: "mapping", Type: "*types.Mapping", Required: false},
	{Name: "MappingTemplate", Flag: "mapping-template", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "OutputConversion", Flag: "output-conversion", Type: "*types.OutputConversion", Required: false},
	{Name: "SampleDocument", Flag: "sample-document", Type: "*string", Required: false},
	{Name: "SampleDocuments", Flag: "sample-documents", Type: "*types.SampleDocuments", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_delete_capability = []leanruntime.Field{
	{Name: "CapabilityId", Flag: "capability-id", Type: "*string", Required: true},
}

var fields_delete_partnership = []leanruntime.Field{
	{Name: "PartnershipId", Flag: "partnership-id", Type: "*string", Required: true},
}

var fields_delete_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_delete_transformer = []leanruntime.Field{
	{Name: "TransformerId", Flag: "transformer-id", Type: "*string", Required: true},
}

var fields_generate_mapping = []leanruntime.Field{
	{Name: "InputFileContent", Flag: "input-file-content", Type: "*string", Required: true},
	{Name: "MappingType", Flag: "mapping-type", Type: "types.MappingType", Required: true},
	{Name: "OutputFileContent", Flag: "output-file-content", Type: "*string", Required: true},
}

var fields_get_capability = []leanruntime.Field{
	{Name: "CapabilityId", Flag: "capability-id", Type: "*string", Required: true},
}

var fields_get_partnership = []leanruntime.Field{
	{Name: "PartnershipId", Flag: "partnership-id", Type: "*string", Required: true},
}

var fields_get_profile = []leanruntime.Field{
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_get_transformer = []leanruntime.Field{
	{Name: "TransformerId", Flag: "transformer-id", Type: "*string", Required: true},
}

var fields_get_transformer_job = []leanruntime.Field{
	{Name: "TransformerId", Flag: "transformer-id", Type: "*string", Required: true},
	{Name: "TransformerJobId", Flag: "transformer-job-id", Type: "*string", Required: true},
}

var fields_list_capabilities = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_partnerships = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: false},
}

var fields_list_profiles = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_transformers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_start_transformer_job = []leanruntime.Field{
	{Name: "ClientToken", Flag: "client-token", Type: "*string", Required: false},
	{Name: "InputFile", Flag: "input-file", Type: "*types.S3Location", Required: true},
	{Name: "OutputLocation", Flag: "output-location", Type: "*types.S3Location", Required: true},
	{Name: "TransformerId", Flag: "transformer-id", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_test_conversion = []leanruntime.Field{
	{Name: "Source", Flag: "source", Type: "*types.ConversionSource", Required: true},
	{Name: "Target", Flag: "target", Type: "*types.ConversionTarget", Required: true},
}

var fields_test_mapping = []leanruntime.Field{
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: true},
	{Name: "InputFileContent", Flag: "input-file-content", Type: "*string", Required: true},
	{Name: "MappingTemplate", Flag: "mapping-template", Type: "*string", Required: true},
}

var fields_test_parsing = []leanruntime.Field{
	{Name: "AdvancedOptions", Flag: "advanced-options", Type: "*types.AdvancedOptions", Required: false},
	{Name: "EdiType", Flag: "edi-type", Type: "types.EdiType", Required: true},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: true},
	{Name: "InputFile", Flag: "input-file", Type: "*types.S3Location", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_capability = []leanruntime.Field{
	{Name: "CapabilityId", Flag: "capability-id", Type: "*string", Required: true},
	{Name: "Configuration", Flag: "configuration", Type: "types.CapabilityConfiguration", Required: false},
	{Name: "InstructionsDocuments", Flag: "instructions-documents", Type: "[]types.S3Location", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
}

var fields_update_partnership = []leanruntime.Field{
	{Name: "Capabilities", Flag: "capabilities", Type: "[]string", Required: false},
	{Name: "CapabilityOptions", Flag: "capability-options", Type: "*types.CapabilityOptions", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "PartnershipId", Flag: "partnership-id", Type: "*string", Required: true},
}

var fields_update_profile = []leanruntime.Field{
	{Name: "BusinessName", Flag: "business-name", Type: "*string", Required: false},
	{Name: "Email", Flag: "email", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "Phone", Flag: "phone", Type: "*string", Required: false},
	{Name: "ProfileId", Flag: "profile-id", Type: "*string", Required: true},
}

var fields_update_transformer = []leanruntime.Field{
	{Name: "EdiType", Flag: "edi-type", Type: "types.EdiType", Required: false},
	{Name: "FileFormat", Flag: "file-format", Type: "types.FileFormat", Required: false},
	{Name: "InputConversion", Flag: "input-conversion", Type: "*types.InputConversion", Required: false},
	{Name: "Mapping", Flag: "mapping", Type: "*types.Mapping", Required: false},
	{Name: "MappingTemplate", Flag: "mapping-template", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: false},
	{Name: "OutputConversion", Flag: "output-conversion", Type: "*types.OutputConversion", Required: false},
	{Name: "SampleDocument", Flag: "sample-document", Type: "*string", Required: false},
	{Name: "SampleDocuments", Flag: "sample-documents", Type: "*types.SampleDocuments", Required: false},
	{Name: "Status", Flag: "status", Type: "types.TransformerStatus", Required: false},
	{Name: "TransformerId", Flag: "transformer-id", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"create-capability": {
			Name:   "create-capability",
			Fields: fields_create_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateCapability(ctx, input)
			},
		},
		"create-partnership": {
			Name:   "create-partnership",
			Fields: fields_create_partnership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePartnershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_partnership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePartnership(ctx, input)
			},
		},
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
		"create-starter-mapping-template": {
			Name:   "create-starter-mapping-template",
			Fields: fields_create_starter_mapping_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStarterMappingTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_starter_mapping_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStarterMappingTemplate(ctx, input)
			},
		},
		"create-transformer": {
			Name:   "create-transformer",
			Fields: fields_create_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTransformer(ctx, input)
			},
		},
		"delete-capability": {
			Name:   "delete-capability",
			Fields: fields_delete_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteCapability(ctx, input)
			},
		},
		"delete-partnership": {
			Name:   "delete-partnership",
			Fields: fields_delete_partnership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePartnershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_partnership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePartnership(ctx, input)
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
		"delete-transformer": {
			Name:   "delete-transformer",
			Fields: fields_delete_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTransformer(ctx, input)
			},
		},
		"generate-mapping": {
			Name:   "generate-mapping",
			Fields: fields_generate_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GenerateMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_generate_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GenerateMapping(ctx, input)
			},
		},
		"get-capability": {
			Name:   "get-capability",
			Fields: fields_get_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCapability(ctx, input)
			},
		},
		"get-partnership": {
			Name:   "get-partnership",
			Fields: fields_get_partnership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPartnershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_partnership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPartnership(ctx, input)
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
		"get-transformer": {
			Name:   "get-transformer",
			Fields: fields_get_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransformer(ctx, input)
			},
		},
		"get-transformer-job": {
			Name:   "get-transformer-job",
			Fields: fields_get_transformer_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetTransformerJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_transformer_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetTransformerJob(ctx, input)
			},
		},
		"list-capabilities": {
			Name:   "list-capabilities",
			Fields: fields_list_capabilities,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListCapabilitiesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_capabilities, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListCapabilities(ctx, input)
				}
				var results []*svc.ListCapabilitiesOutput
				p := svc.NewListCapabilitiesPaginator(client, input)
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
		"list-partnerships": {
			Name:   "list-partnerships",
			Fields: fields_list_partnerships,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPartnershipsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_partnerships, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPartnerships(ctx, input)
				}
				var results []*svc.ListPartnershipsOutput
				p := svc.NewListPartnershipsPaginator(client, input)
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
		"list-transformers": {
			Name:   "list-transformers",
			Fields: fields_list_transformers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTransformersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_transformers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTransformers(ctx, input)
				}
				var results []*svc.ListTransformersOutput
				p := svc.NewListTransformersPaginator(client, input)
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
		"start-transformer-job": {
			Name:   "start-transformer-job",
			Fields: fields_start_transformer_job,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.StartTransformerJobInput{}
				if _, err := leanruntime.ApplyInput(input, fields_start_transformer_job, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.StartTransformerJob(ctx, input)
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
		"test-conversion": {
			Name:   "test-conversion",
			Fields: fields_test_conversion,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestConversionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_conversion, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestConversion(ctx, input)
			},
		},
		"test-mapping": {
			Name:   "test-mapping",
			Fields: fields_test_mapping,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestMappingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_mapping, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestMapping(ctx, input)
			},
		},
		"test-parsing": {
			Name:   "test-parsing",
			Fields: fields_test_parsing,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.TestParsingInput{}
				if _, err := leanruntime.ApplyInput(input, fields_test_parsing, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.TestParsing(ctx, input)
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
		"update-capability": {
			Name:   "update-capability",
			Fields: fields_update_capability,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCapabilityInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_capability, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCapability(ctx, input)
			},
		},
		"update-partnership": {
			Name:   "update-partnership",
			Fields: fields_update_partnership,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePartnershipInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_partnership, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePartnership(ctx, input)
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
		"update-transformer": {
			Name:   "update-transformer",
			Fields: fields_update_transformer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTransformerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_transformer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTransformer(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("b2bi", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
