package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/iotfleetwise"
)

var fields_associate_vehicle_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_batch_create_vehicle = []leanruntime.Field{
	{Name: "Vehicles", Flag: "vehicles", Type: "[]types.CreateVehicleRequestItem", Required: true},
}

var fields_batch_update_vehicle = []leanruntime.Field{
	{Name: "Vehicles", Flag: "vehicles", Type: "[]types.UpdateVehicleRequestItem", Required: true},
}

var fields_create_campaign = []leanruntime.Field{
	{Name: "CollectionScheme", Flag: "collection-scheme", Type: "types.CollectionScheme", Required: true},
	{Name: "Compression", Flag: "compression", Type: "types.Compression", Required: false},
	{Name: "DataDestinationConfigs", Flag: "data-destination-configs", Type: "[]types.DataDestinationConfig", Required: false},
	{Name: "DataExtraDimensions", Flag: "data-extra-dimensions", Type: "[]string", Required: false},
	{Name: "DataPartitions", Flag: "data-partitions", Type: "[]types.DataPartition", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "DiagnosticsMode", Flag: "diagnostics-mode", Type: "types.DiagnosticsMode", Required: false},
	{Name: "ExpiryTime", Flag: "expiry-time", Type: "*time.Time", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "PostTriggerCollectionDuration", Flag: "post-trigger-collection-duration", Type: "*int64", Required: false},
	{Name: "Priority", Flag: "priority", Type: "*int32", Required: false},
	{Name: "SignalCatalogArn", Flag: "signal-catalog-arn", Type: "*string", Required: true},
	{Name: "SignalsToCollect", Flag: "signals-to-collect", Type: "[]types.SignalInformation", Required: false},
	{Name: "SignalsToFetch", Flag: "signals-to-fetch", Type: "[]types.SignalFetchInformation", Required: false},
	{Name: "SpoolingMode", Flag: "spooling-mode", Type: "types.SpoolingMode", Required: false},
	{Name: "StartTime", Flag: "start-time", Type: "*time.Time", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "TargetArn", Flag: "target-arn", Type: "*string", Required: true},
}

var fields_create_decoder_manifest = []leanruntime.Field{
	{Name: "DefaultForUnmappedSignals", Flag: "default-for-unmapped-signals", Type: "types.DefaultForUnmappedSignalsType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ModelManifestArn", Flag: "model-manifest-arn", Type: "*string", Required: true},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkInterfaces", Flag: "network-interfaces", Type: "[]types.NetworkInterface", Required: false},
	{Name: "SignalDecoders", Flag: "signal-decoders", Type: "[]types.SignalDecoder", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_fleet = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "SignalCatalogArn", Flag: "signal-catalog-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_model_manifest = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Nodes", Flag: "nodes", Type: "[]string", Required: true},
	{Name: "SignalCatalogArn", Flag: "signal-catalog-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_signal_catalog = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Nodes", Flag: "nodes", Type: "[]types.Node", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_state_template = []leanruntime.Field{
	{Name: "DataExtraDimensions", Flag: "data-extra-dimensions", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MetadataExtraDimensions", Flag: "metadata-extra-dimensions", Type: "[]string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "SignalCatalogArn", Flag: "signal-catalog-arn", Type: "*string", Required: true},
	{Name: "StateTemplateProperties", Flag: "state-template-properties", Type: "[]string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
}

var fields_create_vehicle = []leanruntime.Field{
	{Name: "AssociationBehavior", Flag: "association-behavior", Type: "types.VehicleAssociationBehavior", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "DecoderManifestArn", Flag: "decoder-manifest-arn", Type: "*string", Required: true},
	{Name: "ModelManifestArn", Flag: "model-manifest-arn", Type: "*string", Required: true},
	{Name: "StateTemplates", Flag: "state-templates", Type: "[]types.StateTemplateAssociation", Required: false},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_delete_campaign = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_decoder_manifest = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_delete_model_manifest = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_signal_catalog = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_delete_state_template = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_delete_vehicle = []leanruntime.Field{
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_disassociate_vehicle_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_get_campaign = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_decoder_manifest = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_encryption_configuration = []leanruntime.Field{}

var fields_get_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_get_logging_options = []leanruntime.Field{}

var fields_get_model_manifest = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_register_account_status = []leanruntime.Field{}

var fields_get_signal_catalog = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_get_state_template = []leanruntime.Field{
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
}

var fields_get_vehicle = []leanruntime.Field{
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_get_vehicle_status = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_import_decoder_manifest = []leanruntime.Field{
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkFileDefinitions", Flag: "network-file-definitions", Type: "[]types.NetworkFileDefinition", Required: true},
}

var fields_import_signal_catalog = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: false},
	{Name: "Vss", Flag: "vss", Type: "types.FormattedVss", Required: false},
}

var fields_list_campaigns = []leanruntime.Field{
	{Name: "ListResponseScope", Flag: "list-response-scope", Type: "types.ListResponseScope", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "Status", Flag: "status", Type: "*string", Required: false},
}

var fields_list_decoder_manifest_network_interfaces = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_decoder_manifest_signals = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_decoder_manifests = []leanruntime.Field{
	{Name: "ListResponseScope", Flag: "list-response-scope", Type: "types.ListResponseScope", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelManifestArn", Flag: "model-manifest-arn", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fleets = []leanruntime.Field{
	{Name: "ListResponseScope", Flag: "list-response-scope", Type: "types.ListResponseScope", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_fleets_for_vehicle = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

var fields_list_model_manifest_nodes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_model_manifests = []leanruntime.Field{
	{Name: "ListResponseScope", Flag: "list-response-scope", Type: "types.ListResponseScope", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SignalCatalogArn", Flag: "signal-catalog-arn", Type: "*string", Required: false},
}

var fields_list_signal_catalog_nodes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SignalNodeType", Flag: "signal-node-type", Type: "types.SignalNodeType", Required: false},
}

var fields_list_signal_catalogs = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_state_templates = []leanruntime.Field{
	{Name: "ListResponseScope", Flag: "list-response-scope", Type: "types.ListResponseScope", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_vehicles = []leanruntime.Field{
	{Name: "AttributeNames", Flag: "attribute-names", Type: "[]string", Required: false},
	{Name: "AttributeValues", Flag: "attribute-values", Type: "[]string", Required: false},
	{Name: "ListResponseScope", Flag: "list-response-scope", Type: "types.ListResponseScope", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "ModelManifestArn", Flag: "model-manifest-arn", Type: "*string", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_vehicles_in_fleet = []leanruntime.Field{
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_encryption_configuration = []leanruntime.Field{
	{Name: "EncryptionType", Flag: "encryption-type", Type: "types.EncryptionType", Required: true},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
}

var fields_put_logging_options = []leanruntime.Field{
	{Name: "CloudWatchLogDelivery", Flag: "cloud-watch-log-delivery", Type: "*types.CloudWatchLogDeliveryOptions", Required: true},
}

var fields_register_account = []leanruntime.Field{
	{Name: "IamResources", Flag: "iam-resources", Type: "*types.IamResources", Required: false},
	{Name: "TimestreamResources", Flag: "timestream-resources", Type: "*types.TimestreamResources", Required: false},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "[]types.Tag", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceARN", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_campaign = []leanruntime.Field{
	{Name: "Action", Flag: "action", Type: "types.UpdateCampaignAction", Required: true},
	{Name: "DataExtraDimensions", Flag: "data-extra-dimensions", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
}

var fields_update_decoder_manifest = []leanruntime.Field{
	{Name: "DefaultForUnmappedSignals", Flag: "default-for-unmapped-signals", Type: "types.DefaultForUnmappedSignalsType", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NetworkInterfacesToAdd", Flag: "network-interfaces-to-add", Type: "[]types.NetworkInterface", Required: false},
	{Name: "NetworkInterfacesToRemove", Flag: "network-interfaces-to-remove", Type: "[]string", Required: false},
	{Name: "NetworkInterfacesToUpdate", Flag: "network-interfaces-to-update", Type: "[]types.NetworkInterface", Required: false},
	{Name: "SignalDecodersToAdd", Flag: "signal-decoders-to-add", Type: "[]types.SignalDecoder", Required: false},
	{Name: "SignalDecodersToRemove", Flag: "signal-decoders-to-remove", Type: "[]string", Required: false},
	{Name: "SignalDecodersToUpdate", Flag: "signal-decoders-to-update", Type: "[]types.SignalDecoder", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ManifestStatus", Required: false},
}

var fields_update_fleet = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "FleetId", Flag: "fleet-id", Type: "*string", Required: true},
}

var fields_update_model_manifest = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NodesToAdd", Flag: "nodes-to-add", Type: "[]string", Required: false},
	{Name: "NodesToRemove", Flag: "nodes-to-remove", Type: "[]string", Required: false},
	{Name: "Status", Flag: "status", Type: "types.ManifestStatus", Required: false},
}

var fields_update_signal_catalog = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Name", Flag: "name", Type: "*string", Required: true},
	{Name: "NodesToAdd", Flag: "nodes-to-add", Type: "[]types.Node", Required: false},
	{Name: "NodesToRemove", Flag: "nodes-to-remove", Type: "[]string", Required: false},
	{Name: "NodesToUpdate", Flag: "nodes-to-update", Type: "[]types.Node", Required: false},
}

var fields_update_state_template = []leanruntime.Field{
	{Name: "DataExtraDimensions", Flag: "data-extra-dimensions", Type: "[]string", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "Identifier", Flag: "identifier", Type: "*string", Required: true},
	{Name: "MetadataExtraDimensions", Flag: "metadata-extra-dimensions", Type: "[]string", Required: false},
	{Name: "StateTemplatePropertiesToAdd", Flag: "state-template-properties-to-add", Type: "[]string", Required: false},
	{Name: "StateTemplatePropertiesToRemove", Flag: "state-template-properties-to-remove", Type: "[]string", Required: false},
}

var fields_update_vehicle = []leanruntime.Field{
	{Name: "AttributeUpdateMode", Flag: "attribute-update-mode", Type: "types.UpdateMode", Required: false},
	{Name: "Attributes", Flag: "attributes", Type: "map[string]string", Required: false},
	{Name: "DecoderManifestArn", Flag: "decoder-manifest-arn", Type: "*string", Required: false},
	{Name: "ModelManifestArn", Flag: "model-manifest-arn", Type: "*string", Required: false},
	{Name: "StateTemplatesToAdd", Flag: "state-templates-to-add", Type: "[]types.StateTemplateAssociation", Required: false},
	{Name: "StateTemplatesToRemove", Flag: "state-templates-to-remove", Type: "[]string", Required: false},
	{Name: "StateTemplatesToUpdate", Flag: "state-templates-to-update", Type: "[]types.StateTemplateAssociation", Required: false},
	{Name: "VehicleName", Flag: "vehicle-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-vehicle-fleet": {
			Name:   "associate-vehicle-fleet",
			Fields: fields_associate_vehicle_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateVehicleFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_vehicle_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateVehicleFleet(ctx, input)
			},
		},
		"batch-create-vehicle": {
			Name:   "batch-create-vehicle",
			Fields: fields_batch_create_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchCreateVehicleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_create_vehicle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchCreateVehicle(ctx, input)
			},
		},
		"batch-update-vehicle": {
			Name:   "batch-update-vehicle",
			Fields: fields_batch_update_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateVehicleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_vehicle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateVehicle(ctx, input)
			},
		},
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
		"create-decoder-manifest": {
			Name:   "create-decoder-manifest",
			Fields: fields_create_decoder_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateDecoderManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_decoder_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateDecoderManifest(ctx, input)
			},
		},
		"create-fleet": {
			Name:   "create-fleet",
			Fields: fields_create_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateFleet(ctx, input)
			},
		},
		"create-model-manifest": {
			Name:   "create-model-manifest",
			Fields: fields_create_model_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateModelManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_model_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateModelManifest(ctx, input)
			},
		},
		"create-signal-catalog": {
			Name:   "create-signal-catalog",
			Fields: fields_create_signal_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateSignalCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_signal_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateSignalCatalog(ctx, input)
			},
		},
		"create-state-template": {
			Name:   "create-state-template",
			Fields: fields_create_state_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateStateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_state_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateStateTemplate(ctx, input)
			},
		},
		"create-vehicle": {
			Name:   "create-vehicle",
			Fields: fields_create_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateVehicleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_vehicle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateVehicle(ctx, input)
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
		"delete-decoder-manifest": {
			Name:   "delete-decoder-manifest",
			Fields: fields_delete_decoder_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteDecoderManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_decoder_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteDecoderManifest(ctx, input)
			},
		},
		"delete-fleet": {
			Name:   "delete-fleet",
			Fields: fields_delete_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteFleet(ctx, input)
			},
		},
		"delete-model-manifest": {
			Name:   "delete-model-manifest",
			Fields: fields_delete_model_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteModelManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_model_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteModelManifest(ctx, input)
			},
		},
		"delete-signal-catalog": {
			Name:   "delete-signal-catalog",
			Fields: fields_delete_signal_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteSignalCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_signal_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteSignalCatalog(ctx, input)
			},
		},
		"delete-state-template": {
			Name:   "delete-state-template",
			Fields: fields_delete_state_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteStateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_state_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteStateTemplate(ctx, input)
			},
		},
		"delete-vehicle": {
			Name:   "delete-vehicle",
			Fields: fields_delete_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteVehicleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_vehicle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteVehicle(ctx, input)
			},
		},
		"disassociate-vehicle-fleet": {
			Name:   "disassociate-vehicle-fleet",
			Fields: fields_disassociate_vehicle_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateVehicleFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_vehicle_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateVehicleFleet(ctx, input)
			},
		},
		"get-campaign": {
			Name:   "get-campaign",
			Fields: fields_get_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetCampaign(ctx, input)
			},
		},
		"get-decoder-manifest": {
			Name:   "get-decoder-manifest",
			Fields: fields_get_decoder_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDecoderManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_decoder_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDecoderManifest(ctx, input)
			},
		},
		"get-encryption-configuration": {
			Name:   "get-encryption-configuration",
			Fields: fields_get_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetEncryptionConfiguration(ctx, input)
			},
		},
		"get-fleet": {
			Name:   "get-fleet",
			Fields: fields_get_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetFleet(ctx, input)
			},
		},
		"get-logging-options": {
			Name:   "get-logging-options",
			Fields: fields_get_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetLoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetLoggingOptions(ctx, input)
			},
		},
		"get-model-manifest": {
			Name:   "get-model-manifest",
			Fields: fields_get_model_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetModelManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_model_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetModelManifest(ctx, input)
			},
		},
		"get-register-account-status": {
			Name:   "get-register-account-status",
			Fields: fields_get_register_account_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetRegisterAccountStatusInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_register_account_status, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetRegisterAccountStatus(ctx, input)
			},
		},
		"get-signal-catalog": {
			Name:   "get-signal-catalog",
			Fields: fields_get_signal_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetSignalCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_signal_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetSignalCatalog(ctx, input)
			},
		},
		"get-state-template": {
			Name:   "get-state-template",
			Fields: fields_get_state_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetStateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_state_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetStateTemplate(ctx, input)
			},
		},
		"get-vehicle": {
			Name:   "get-vehicle",
			Fields: fields_get_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVehicleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_vehicle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetVehicle(ctx, input)
			},
		},
		"get-vehicle-status": {
			Name:   "get-vehicle-status",
			Fields: fields_get_vehicle_status,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetVehicleStatusInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_vehicle_status, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetVehicleStatus(ctx, input)
				}
				var results []*svc.GetVehicleStatusOutput
				p := svc.NewGetVehicleStatusPaginator(client, input)
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
		"import-decoder-manifest": {
			Name:   "import-decoder-manifest",
			Fields: fields_import_decoder_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportDecoderManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_decoder_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportDecoderManifest(ctx, input)
			},
		},
		"import-signal-catalog": {
			Name:   "import-signal-catalog",
			Fields: fields_import_signal_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ImportSignalCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_import_signal_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.ImportSignalCatalog(ctx, input)
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
		"list-decoder-manifest-network-interfaces": {
			Name:   "list-decoder-manifest-network-interfaces",
			Fields: fields_list_decoder_manifest_network_interfaces,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDecoderManifestNetworkInterfacesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_decoder_manifest_network_interfaces, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDecoderManifestNetworkInterfaces(ctx, input)
				}
				var results []*svc.ListDecoderManifestNetworkInterfacesOutput
				p := svc.NewListDecoderManifestNetworkInterfacesPaginator(client, input)
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
		"list-decoder-manifest-signals": {
			Name:   "list-decoder-manifest-signals",
			Fields: fields_list_decoder_manifest_signals,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDecoderManifestSignalsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_decoder_manifest_signals, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDecoderManifestSignals(ctx, input)
				}
				var results []*svc.ListDecoderManifestSignalsOutput
				p := svc.NewListDecoderManifestSignalsPaginator(client, input)
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
		"list-decoder-manifests": {
			Name:   "list-decoder-manifests",
			Fields: fields_list_decoder_manifests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDecoderManifestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_decoder_manifests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDecoderManifests(ctx, input)
				}
				var results []*svc.ListDecoderManifestsOutput
				p := svc.NewListDecoderManifestsPaginator(client, input)
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
		"list-fleets": {
			Name:   "list-fleets",
			Fields: fields_list_fleets,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleets, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleets(ctx, input)
				}
				var results []*svc.ListFleetsOutput
				p := svc.NewListFleetsPaginator(client, input)
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
		"list-fleets-for-vehicle": {
			Name:   "list-fleets-for-vehicle",
			Fields: fields_list_fleets_for_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListFleetsForVehicleInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_fleets_for_vehicle, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListFleetsForVehicle(ctx, input)
				}
				var results []*svc.ListFleetsForVehicleOutput
				p := svc.NewListFleetsForVehiclePaginator(client, input)
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
		"list-model-manifest-nodes": {
			Name:   "list-model-manifest-nodes",
			Fields: fields_list_model_manifest_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelManifestNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_manifest_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelManifestNodes(ctx, input)
				}
				var results []*svc.ListModelManifestNodesOutput
				p := svc.NewListModelManifestNodesPaginator(client, input)
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
		"list-model-manifests": {
			Name:   "list-model-manifests",
			Fields: fields_list_model_manifests,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListModelManifestsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_model_manifests, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListModelManifests(ctx, input)
				}
				var results []*svc.ListModelManifestsOutput
				p := svc.NewListModelManifestsPaginator(client, input)
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
		"list-signal-catalog-nodes": {
			Name:   "list-signal-catalog-nodes",
			Fields: fields_list_signal_catalog_nodes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSignalCatalogNodesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signal_catalog_nodes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSignalCatalogNodes(ctx, input)
				}
				var results []*svc.ListSignalCatalogNodesOutput
				p := svc.NewListSignalCatalogNodesPaginator(client, input)
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
		"list-signal-catalogs": {
			Name:   "list-signal-catalogs",
			Fields: fields_list_signal_catalogs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListSignalCatalogsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_signal_catalogs, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListSignalCatalogs(ctx, input)
				}
				var results []*svc.ListSignalCatalogsOutput
				p := svc.NewListSignalCatalogsPaginator(client, input)
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
		"list-state-templates": {
			Name:   "list-state-templates",
			Fields: fields_list_state_templates,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListStateTemplatesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_state_templates, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListStateTemplates(ctx, input)
				}
				var results []*svc.ListStateTemplatesOutput
				p := svc.NewListStateTemplatesPaginator(client, input)
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
		"list-vehicles": {
			Name:   "list-vehicles",
			Fields: fields_list_vehicles,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVehiclesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vehicles, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVehicles(ctx, input)
				}
				var results []*svc.ListVehiclesOutput
				p := svc.NewListVehiclesPaginator(client, input)
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
		"list-vehicles-in-fleet": {
			Name:   "list-vehicles-in-fleet",
			Fields: fields_list_vehicles_in_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListVehiclesInFleetInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_vehicles_in_fleet, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListVehiclesInFleet(ctx, input)
				}
				var results []*svc.ListVehiclesInFleetOutput
				p := svc.NewListVehiclesInFleetPaginator(client, input)
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
		"put-encryption-configuration": {
			Name:   "put-encryption-configuration",
			Fields: fields_put_encryption_configuration,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutEncryptionConfigurationInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_encryption_configuration, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutEncryptionConfiguration(ctx, input)
			},
		},
		"put-logging-options": {
			Name:   "put-logging-options",
			Fields: fields_put_logging_options,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutLoggingOptionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_logging_options, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutLoggingOptions(ctx, input)
			},
		},
		"register-account": {
			Name:   "register-account",
			Fields: fields_register_account,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.RegisterAccountInput{}
				if _, err := leanruntime.ApplyInput(input, fields_register_account, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.RegisterAccount(ctx, input)
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
		"update-campaign": {
			Name:   "update-campaign",
			Fields: fields_update_campaign,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateCampaignInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_campaign, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateCampaign(ctx, input)
			},
		},
		"update-decoder-manifest": {
			Name:   "update-decoder-manifest",
			Fields: fields_update_decoder_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateDecoderManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_decoder_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateDecoderManifest(ctx, input)
			},
		},
		"update-fleet": {
			Name:   "update-fleet",
			Fields: fields_update_fleet,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateFleetInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_fleet, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateFleet(ctx, input)
			},
		},
		"update-model-manifest": {
			Name:   "update-model-manifest",
			Fields: fields_update_model_manifest,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateModelManifestInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_model_manifest, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateModelManifest(ctx, input)
			},
		},
		"update-signal-catalog": {
			Name:   "update-signal-catalog",
			Fields: fields_update_signal_catalog,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateSignalCatalogInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_signal_catalog, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateSignalCatalog(ctx, input)
			},
		},
		"update-state-template": {
			Name:   "update-state-template",
			Fields: fields_update_state_template,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateStateTemplateInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_state_template, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateStateTemplate(ctx, input)
			},
		},
		"update-vehicle": {
			Name:   "update-vehicle",
			Fields: fields_update_vehicle,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateVehicleInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_vehicle, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateVehicle(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("iotfleetwise", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
