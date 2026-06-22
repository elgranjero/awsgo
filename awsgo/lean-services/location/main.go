package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/location"
)

var fields_associate_tracker_consumer = []leanruntime.Field{
	{Name: "ConsumerArn", Flag: "consumer-arn", Type: "*string", Required: true},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_batch_delete_device_position_history = []leanruntime.Field{
	{Name: "DeviceIds", Flag: "device-ids", Type: "[]string", Required: true},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_batch_delete_geofence = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "GeofenceIds", Flag: "geofence-ids", Type: "[]string", Required: true},
}

var fields_batch_evaluate_geofences = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "DevicePositionUpdates", Flag: "device-position-updates", Type: "[]types.DevicePositionUpdate", Required: true},
}

var fields_batch_get_device_position = []leanruntime.Field{
	{Name: "DeviceIds", Flag: "device-ids", Type: "[]string", Required: true},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_batch_put_geofence = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "Entries", Flag: "entries", Type: "[]types.BatchPutGeofenceRequestEntry", Required: true},
}

var fields_batch_update_device_position = []leanruntime.Field{
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
	{Name: "Updates", Flag: "updates", Type: "[]types.DevicePositionUpdate", Required: true},
}

var fields_calculate_route = []leanruntime.Field{
	{Name: "ArrivalTime", Flag: "arrival-time", Type: "*time.Time", Required: false},
	{Name: "CalculatorName", Flag: "calculator-name", Type: "*string", Required: true},
	{Name: "CarModeOptions", Flag: "car-mode-options", Type: "*types.CalculateRouteCarModeOptions", Required: false},
	{Name: "DepartNow", Flag: "depart-now", Type: "*bool", Required: false},
	{Name: "DeparturePosition", Flag: "departure-position", Type: "[]float64", Required: true},
	{Name: "DepartureTime", Flag: "departure-time", Type: "*time.Time", Required: false},
	{Name: "DestinationPosition", Flag: "destination-position", Type: "[]float64", Required: true},
	{Name: "DistanceUnit", Flag: "distance-unit", Type: "types.DistanceUnit", Required: false},
	{Name: "IncludeLegGeometry", Flag: "include-leg-geometry", Type: "*bool", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "OptimizeFor", Flag: "optimize-for", Type: "types.OptimizationMode", Required: false},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.TravelMode", Required: false},
	{Name: "TruckModeOptions", Flag: "truck-mode-options", Type: "*types.CalculateRouteTruckModeOptions", Required: false},
	{Name: "WaypointPositions", Flag: "waypoint-positions", Type: "[][]float64", Required: false},
}

var fields_calculate_route_matrix = []leanruntime.Field{
	{Name: "CalculatorName", Flag: "calculator-name", Type: "*string", Required: true},
	{Name: "CarModeOptions", Flag: "car-mode-options", Type: "*types.CalculateRouteCarModeOptions", Required: false},
	{Name: "DepartNow", Flag: "depart-now", Type: "*bool", Required: false},
	{Name: "DeparturePositions", Flag: "departure-positions", Type: "[][]float64", Required: true},
	{Name: "DepartureTime", Flag: "departure-time", Type: "*time.Time", Required: false},
	{Name: "DestinationPositions", Flag: "destination-positions", Type: "[][]float64", Required: true},
	{Name: "DistanceUnit", Flag: "distance-unit", Type: "types.DistanceUnit", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.TravelMode", Required: false},
	{Name: "TruckModeOptions", Flag: "truck-mode-options", Type: "*types.CalculateRouteTruckModeOptions", Required: false},
}

var fields_create_geofence_collection = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "PricingPlanDataSource", Flag: "pricing-plan-data-source", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_key = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpireTime", Flag: "expire-time", Type: "*time.Time", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "NoExpiry", Flag: "no-expiry", Type: "*bool", Required: false},
	{Name: "Restrictions", Flag: "restrictions", Type: "*types.ApiKeyRestrictions", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_map = []leanruntime.Field{
	{Name: "Configuration", Flag: "configuration", Type: "*types.MapConfiguration", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_place_index = []leanruntime.Field{
	{Name: "DataSource", Flag: "data-source", Type: "*string", Required: true},
	{Name: "DataSourceConfiguration", Flag: "data-source-configuration", Type: "*types.DataSourceConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_route_calculator = []leanruntime.Field{
	{Name: "CalculatorName", Flag: "calculator-name", Type: "*string", Required: true},
	{Name: "DataSource", Flag: "data-source", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
}

var fields_create_tracker = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBridgeEnabled", Flag: "event-bridge-enabled", Type: "*bool", Required: false},
	{Name: "KmsKeyEnableGeospatialQueries", Flag: "kms-key-enable-geospatial-queries", Type: "*bool", Required: false},
	{Name: "KmsKeyId", Flag: "kms-key-id", Type: "*string", Required: false},
	{Name: "PositionFiltering", Flag: "position-filtering", Type: "types.PositionFiltering", Required: false},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "PricingPlanDataSource", Flag: "pricing-plan-data-source", Type: "*string", Required: false},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: false},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_delete_geofence_collection = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
}

var fields_delete_key = []leanruntime.Field{
	{Name: "ForceDelete", Flag: "force-delete", Type: "*bool", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
}

var fields_delete_map = []leanruntime.Field{
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
}

var fields_delete_place_index = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_delete_route_calculator = []leanruntime.Field{
	{Name: "CalculatorName", Flag: "calculator-name", Type: "*string", Required: true},
}

var fields_delete_tracker = []leanruntime.Field{
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_describe_geofence_collection = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
}

var fields_describe_key = []leanruntime.Field{
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
}

var fields_describe_map = []leanruntime.Field{
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
}

var fields_describe_place_index = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
}

var fields_describe_route_calculator = []leanruntime.Field{
	{Name: "CalculatorName", Flag: "calculator-name", Type: "*string", Required: true},
}

var fields_describe_tracker = []leanruntime.Field{
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_disassociate_tracker_consumer = []leanruntime.Field{
	{Name: "ConsumerArn", Flag: "consumer-arn", Type: "*string", Required: true},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_forecast_geofence_events = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "DeviceState", Flag: "device-state", Type: "*types.ForecastGeofenceEventsDeviceState", Required: true},
	{Name: "DistanceUnit", Flag: "distance-unit", Type: "types.DistanceUnit", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "SpeedUnit", Flag: "speed-unit", Type: "types.SpeedUnit", Required: false},
	{Name: "TimeHorizonMinutes", Flag: "time-horizon-minutes", Type: "*float64", Required: false},
}

var fields_get_device_position = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_get_device_position_history = []leanruntime.Field{
	{Name: "DeviceId", Flag: "device-id", Type: "*string", Required: true},
	{Name: "EndTimeExclusive", Flag: "end-time-exclusive", Type: "*time.Time", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "StartTimeInclusive", Flag: "start-time-inclusive", Type: "*time.Time", Required: false},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_get_geofence = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "GeofenceId", Flag: "geofence-id", Type: "*string", Required: true},
}

var fields_get_map_glyphs = []leanruntime.Field{
	{Name: "FontStack", Flag: "font-stack", Type: "*string", Required: true},
	{Name: "FontUnicodeRange", Flag: "font-unicode-range", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
}

var fields_get_map_sprites = []leanruntime.Field{
	{Name: "FileName", Flag: "file-name", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
}

var fields_get_map_style_descriptor = []leanruntime.Field{
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
}

var fields_get_map_tile = []leanruntime.Field{
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
	{Name: "X", Flag: "x", Type: "*string", Required: true},
	{Name: "Y", Flag: "y", Type: "*string", Required: true},
	{Name: "Z", Flag: "z", Type: "*string", Required: true},
}

var fields_get_place = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "PlaceId", Flag: "place-id", Type: "*string", Required: true},
}

var fields_list_device_positions = []leanruntime.Field{
	{Name: "FilterGeometry", Flag: "filter-geometry", Type: "*types.TrackingFilterGeometry", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_list_geofence_collections = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_geofences = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_keys = []leanruntime.Field{
	{Name: "Filter", Flag: "filter", Type: "*types.ApiKeyFilter", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_maps = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_place_indexes = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_route_calculators = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_list_tags_for_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
}

var fields_list_tracker_consumers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_list_trackers = []leanruntime.Field{
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "NextToken", Flag: "next-token", Type: "*string", Required: false},
}

var fields_put_geofence = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "GeofenceId", Flag: "geofence-id", Type: "*string", Required: true},
	{Name: "GeofenceProperties", Flag: "geofence-properties", Type: "map[string]string", Required: false},
	{Name: "Geometry", Flag: "geometry", Type: "*types.GeofenceGeometry", Required: true},
}

var fields_search_place_index_for_position = []leanruntime.Field{
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Position", Flag: "position", Type: "[]float64", Required: true},
}

var fields_search_place_index_for_suggestions = []leanruntime.Field{
	{Name: "BiasPosition", Flag: "bias-position", Type: "[]float64", Required: false},
	{Name: "FilterBBox", Flag: "filter-bbox", Type: "[]float64", Required: false},
	{Name: "FilterCategories", Flag: "filter-categories", Type: "[]string", Required: false},
	{Name: "FilterCountries", Flag: "filter-countries", Type: "[]string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_search_place_index_for_text = []leanruntime.Field{
	{Name: "BiasPosition", Flag: "bias-position", Type: "[]float64", Required: false},
	{Name: "FilterBBox", Flag: "filter-bbox", Type: "[]float64", Required: false},
	{Name: "FilterCategories", Flag: "filter-categories", Type: "[]string", Required: false},
	{Name: "FilterCountries", Flag: "filter-countries", Type: "[]string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Language", Flag: "language", Type: "*string", Required: false},
	{Name: "MaxResults", Flag: "max-results", Type: "*int32", Required: false},
	{Name: "Text", Flag: "text", Type: "*string", Required: true},
}

var fields_tag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "Tags", Flag: "tags", Type: "map[string]string", Required: true},
}

var fields_untag_resource = []leanruntime.Field{
	{Name: "ResourceArn", Flag: "resource-arn", Type: "*string", Required: true},
	{Name: "TagKeys", Flag: "tag-keys", Type: "[]string", Required: true},
}

var fields_update_geofence_collection = []leanruntime.Field{
	{Name: "CollectionName", Flag: "collection-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "PricingPlanDataSource", Flag: "pricing-plan-data-source", Type: "*string", Required: false},
}

var fields_update_key = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "ExpireTime", Flag: "expire-time", Type: "*time.Time", Required: false},
	{Name: "ForceUpdate", Flag: "force-update", Type: "*bool", Required: false},
	{Name: "KeyName", Flag: "key-name", Type: "*string", Required: true},
	{Name: "NoExpiry", Flag: "no-expiry", Type: "*bool", Required: false},
	{Name: "Restrictions", Flag: "restrictions", Type: "*types.ApiKeyRestrictions", Required: false},
}

var fields_update_map = []leanruntime.Field{
	{Name: "ConfigurationUpdate", Flag: "configuration-update", Type: "*types.MapConfigurationUpdate", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "MapName", Flag: "map-name", Type: "*string", Required: true},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
}

var fields_update_place_index = []leanruntime.Field{
	{Name: "DataSourceConfiguration", Flag: "data-source-configuration", Type: "*types.DataSourceConfiguration", Required: false},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "IndexName", Flag: "index-name", Type: "*string", Required: true},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
}

var fields_update_route_calculator = []leanruntime.Field{
	{Name: "CalculatorName", Flag: "calculator-name", Type: "*string", Required: true},
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
}

var fields_update_tracker = []leanruntime.Field{
	{Name: "Description", Flag: "description", Type: "*string", Required: false},
	{Name: "EventBridgeEnabled", Flag: "event-bridge-enabled", Type: "*bool", Required: false},
	{Name: "KmsKeyEnableGeospatialQueries", Flag: "kms-key-enable-geospatial-queries", Type: "*bool", Required: false},
	{Name: "PositionFiltering", Flag: "position-filtering", Type: "types.PositionFiltering", Required: false},
	{Name: "PricingPlan", Flag: "pricing-plan", Type: "types.PricingPlan", Required: false},
	{Name: "PricingPlanDataSource", Flag: "pricing-plan-data-source", Type: "*string", Required: false},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

var fields_verify_device_position = []leanruntime.Field{
	{Name: "DeviceState", Flag: "device-state", Type: "*types.DeviceState", Required: true},
	{Name: "DistanceUnit", Flag: "distance-unit", Type: "types.DistanceUnit", Required: false},
	{Name: "TrackerName", Flag: "tracker-name", Type: "*string", Required: true},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"associate-tracker-consumer": {
			Name:   "associate-tracker-consumer",
			Fields: fields_associate_tracker_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.AssociateTrackerConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_associate_tracker_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.AssociateTrackerConsumer(ctx, input)
			},
		},
		"batch-delete-device-position-history": {
			Name:   "batch-delete-device-position-history",
			Fields: fields_batch_delete_device_position_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteDevicePositionHistoryInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_device_position_history, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteDevicePositionHistory(ctx, input)
			},
		},
		"batch-delete-geofence": {
			Name:   "batch-delete-geofence",
			Fields: fields_batch_delete_geofence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchDeleteGeofenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_delete_geofence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchDeleteGeofence(ctx, input)
			},
		},
		"batch-evaluate-geofences": {
			Name:   "batch-evaluate-geofences",
			Fields: fields_batch_evaluate_geofences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchEvaluateGeofencesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_evaluate_geofences, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchEvaluateGeofences(ctx, input)
			},
		},
		"batch-get-device-position": {
			Name:   "batch-get-device-position",
			Fields: fields_batch_get_device_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchGetDevicePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_get_device_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchGetDevicePosition(ctx, input)
			},
		},
		"batch-put-geofence": {
			Name:   "batch-put-geofence",
			Fields: fields_batch_put_geofence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchPutGeofenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_put_geofence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchPutGeofence(ctx, input)
			},
		},
		"batch-update-device-position": {
			Name:   "batch-update-device-position",
			Fields: fields_batch_update_device_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.BatchUpdateDevicePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_batch_update_device_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.BatchUpdateDevicePosition(ctx, input)
			},
		},
		"calculate-route": {
			Name:   "calculate-route",
			Fields: fields_calculate_route,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CalculateRouteInput{}
				if _, err := leanruntime.ApplyInput(input, fields_calculate_route, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CalculateRoute(ctx, input)
			},
		},
		"calculate-route-matrix": {
			Name:   "calculate-route-matrix",
			Fields: fields_calculate_route_matrix,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CalculateRouteMatrixInput{}
				if _, err := leanruntime.ApplyInput(input, fields_calculate_route_matrix, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CalculateRouteMatrix(ctx, input)
			},
		},
		"create-geofence-collection": {
			Name:   "create-geofence-collection",
			Fields: fields_create_geofence_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateGeofenceCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_geofence_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateGeofenceCollection(ctx, input)
			},
		},
		"create-key": {
			Name:   "create-key",
			Fields: fields_create_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateKey(ctx, input)
			},
		},
		"create-map": {
			Name:   "create-map",
			Fields: fields_create_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateMap(ctx, input)
			},
		},
		"create-place-index": {
			Name:   "create-place-index",
			Fields: fields_create_place_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreatePlaceIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_place_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreatePlaceIndex(ctx, input)
			},
		},
		"create-route-calculator": {
			Name:   "create-route-calculator",
			Fields: fields_create_route_calculator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateRouteCalculatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_route_calculator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateRouteCalculator(ctx, input)
			},
		},
		"create-tracker": {
			Name:   "create-tracker",
			Fields: fields_create_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CreateTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_create_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CreateTracker(ctx, input)
			},
		},
		"delete-geofence-collection": {
			Name:   "delete-geofence-collection",
			Fields: fields_delete_geofence_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteGeofenceCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_geofence_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteGeofenceCollection(ctx, input)
			},
		},
		"delete-key": {
			Name:   "delete-key",
			Fields: fields_delete_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteKey(ctx, input)
			},
		},
		"delete-map": {
			Name:   "delete-map",
			Fields: fields_delete_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteMap(ctx, input)
			},
		},
		"delete-place-index": {
			Name:   "delete-place-index",
			Fields: fields_delete_place_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeletePlaceIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_place_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeletePlaceIndex(ctx, input)
			},
		},
		"delete-route-calculator": {
			Name:   "delete-route-calculator",
			Fields: fields_delete_route_calculator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteRouteCalculatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_route_calculator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteRouteCalculator(ctx, input)
			},
		},
		"delete-tracker": {
			Name:   "delete-tracker",
			Fields: fields_delete_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DeleteTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_delete_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DeleteTracker(ctx, input)
			},
		},
		"describe-geofence-collection": {
			Name:   "describe-geofence-collection",
			Fields: fields_describe_geofence_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeGeofenceCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_geofence_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeGeofenceCollection(ctx, input)
			},
		},
		"describe-key": {
			Name:   "describe-key",
			Fields: fields_describe_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeKey(ctx, input)
			},
		},
		"describe-map": {
			Name:   "describe-map",
			Fields: fields_describe_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeMap(ctx, input)
			},
		},
		"describe-place-index": {
			Name:   "describe-place-index",
			Fields: fields_describe_place_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribePlaceIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_place_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribePlaceIndex(ctx, input)
			},
		},
		"describe-route-calculator": {
			Name:   "describe-route-calculator",
			Fields: fields_describe_route_calculator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeRouteCalculatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_route_calculator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeRouteCalculator(ctx, input)
			},
		},
		"describe-tracker": {
			Name:   "describe-tracker",
			Fields: fields_describe_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DescribeTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_describe_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DescribeTracker(ctx, input)
			},
		},
		"disassociate-tracker-consumer": {
			Name:   "disassociate-tracker-consumer",
			Fields: fields_disassociate_tracker_consumer,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.DisassociateTrackerConsumerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_disassociate_tracker_consumer, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.DisassociateTrackerConsumer(ctx, input)
			},
		},
		"forecast-geofence-events": {
			Name:   "forecast-geofence-events",
			Fields: fields_forecast_geofence_events,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ForecastGeofenceEventsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_forecast_geofence_events, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ForecastGeofenceEvents(ctx, input)
				}
				var results []*svc.ForecastGeofenceEventsOutput
				p := svc.NewForecastGeofenceEventsPaginator(client, input)
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
		"get-device-position": {
			Name:   "get-device-position",
			Fields: fields_get_device_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevicePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_device_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetDevicePosition(ctx, input)
			},
		},
		"get-device-position-history": {
			Name:   "get-device-position-history",
			Fields: fields_get_device_position_history,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetDevicePositionHistoryInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_get_device_position_history, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.GetDevicePositionHistory(ctx, input)
				}
				var results []*svc.GetDevicePositionHistoryOutput
				p := svc.NewGetDevicePositionHistoryPaginator(client, input)
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
		"get-geofence": {
			Name:   "get-geofence",
			Fields: fields_get_geofence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetGeofenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_geofence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetGeofence(ctx, input)
			},
		},
		"get-map-glyphs": {
			Name:   "get-map-glyphs",
			Fields: fields_get_map_glyphs,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMapGlyphsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_map_glyphs, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMapGlyphs(ctx, input)
			},
		},
		"get-map-sprites": {
			Name:   "get-map-sprites",
			Fields: fields_get_map_sprites,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMapSpritesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_map_sprites, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMapSprites(ctx, input)
			},
		},
		"get-map-style-descriptor": {
			Name:   "get-map-style-descriptor",
			Fields: fields_get_map_style_descriptor,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMapStyleDescriptorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_map_style_descriptor, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMapStyleDescriptor(ctx, input)
			},
		},
		"get-map-tile": {
			Name:   "get-map-tile",
			Fields: fields_get_map_tile,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetMapTileInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_map_tile, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetMapTile(ctx, input)
			},
		},
		"get-place": {
			Name:   "get-place",
			Fields: fields_get_place,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.GetPlaceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_get_place, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.GetPlace(ctx, input)
			},
		},
		"list-device-positions": {
			Name:   "list-device-positions",
			Fields: fields_list_device_positions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListDevicePositionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_device_positions, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListDevicePositions(ctx, input)
				}
				var results []*svc.ListDevicePositionsOutput
				p := svc.NewListDevicePositionsPaginator(client, input)
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
		"list-geofence-collections": {
			Name:   "list-geofence-collections",
			Fields: fields_list_geofence_collections,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGeofenceCollectionsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_geofence_collections, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGeofenceCollections(ctx, input)
				}
				var results []*svc.ListGeofenceCollectionsOutput
				p := svc.NewListGeofenceCollectionsPaginator(client, input)
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
		"list-geofences": {
			Name:   "list-geofences",
			Fields: fields_list_geofences,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListGeofencesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_geofences, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListGeofences(ctx, input)
				}
				var results []*svc.ListGeofencesOutput
				p := svc.NewListGeofencesPaginator(client, input)
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
		"list-keys": {
			Name:   "list-keys",
			Fields: fields_list_keys,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListKeysInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_keys, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListKeys(ctx, input)
				}
				var results []*svc.ListKeysOutput
				p := svc.NewListKeysPaginator(client, input)
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
		"list-maps": {
			Name:   "list-maps",
			Fields: fields_list_maps,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListMapsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_maps, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListMaps(ctx, input)
				}
				var results []*svc.ListMapsOutput
				p := svc.NewListMapsPaginator(client, input)
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
		"list-place-indexes": {
			Name:   "list-place-indexes",
			Fields: fields_list_place_indexes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListPlaceIndexesInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_place_indexes, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListPlaceIndexes(ctx, input)
				}
				var results []*svc.ListPlaceIndexesOutput
				p := svc.NewListPlaceIndexesPaginator(client, input)
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
		"list-route-calculators": {
			Name:   "list-route-calculators",
			Fields: fields_list_route_calculators,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListRouteCalculatorsInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_route_calculators, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListRouteCalculators(ctx, input)
				}
				var results []*svc.ListRouteCalculatorsOutput
				p := svc.NewListRouteCalculatorsPaginator(client, input)
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
		"list-tracker-consumers": {
			Name:   "list-tracker-consumers",
			Fields: fields_list_tracker_consumers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrackerConsumersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_tracker_consumers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrackerConsumers(ctx, input)
				}
				var results []*svc.ListTrackerConsumersOutput
				p := svc.NewListTrackerConsumersPaginator(client, input)
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
		"list-trackers": {
			Name:   "list-trackers",
			Fields: fields_list_trackers,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.ListTrackersInput{}
				disablePaginator, err := leanruntime.ApplyInput(input, fields_list_trackers, values)
				if err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				if disablePaginator || leanruntime.PaginatorDisabled() {
					return client.ListTrackers(ctx, input)
				}
				var results []*svc.ListTrackersOutput
				p := svc.NewListTrackersPaginator(client, input)
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
		"put-geofence": {
			Name:   "put-geofence",
			Fields: fields_put_geofence,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.PutGeofenceInput{}
				if _, err := leanruntime.ApplyInput(input, fields_put_geofence, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.PutGeofence(ctx, input)
			},
		},
		"search-place-index-for-position": {
			Name:   "search-place-index-for-position",
			Fields: fields_search_place_index_for_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchPlaceIndexForPositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_place_index_for_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchPlaceIndexForPosition(ctx, input)
			},
		},
		"search-place-index-for-suggestions": {
			Name:   "search-place-index-for-suggestions",
			Fields: fields_search_place_index_for_suggestions,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchPlaceIndexForSuggestionsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_place_index_for_suggestions, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchPlaceIndexForSuggestions(ctx, input)
			},
		},
		"search-place-index-for-text": {
			Name:   "search-place-index-for-text",
			Fields: fields_search_place_index_for_text,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SearchPlaceIndexForTextInput{}
				if _, err := leanruntime.ApplyInput(input, fields_search_place_index_for_text, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SearchPlaceIndexForText(ctx, input)
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
		"update-geofence-collection": {
			Name:   "update-geofence-collection",
			Fields: fields_update_geofence_collection,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateGeofenceCollectionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_geofence_collection, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateGeofenceCollection(ctx, input)
			},
		},
		"update-key": {
			Name:   "update-key",
			Fields: fields_update_key,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateKeyInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_key, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateKey(ctx, input)
			},
		},
		"update-map": {
			Name:   "update-map",
			Fields: fields_update_map,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateMapInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_map, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateMap(ctx, input)
			},
		},
		"update-place-index": {
			Name:   "update-place-index",
			Fields: fields_update_place_index,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdatePlaceIndexInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_place_index, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdatePlaceIndex(ctx, input)
			},
		},
		"update-route-calculator": {
			Name:   "update-route-calculator",
			Fields: fields_update_route_calculator,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateRouteCalculatorInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_route_calculator, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateRouteCalculator(ctx, input)
			},
		},
		"update-tracker": {
			Name:   "update-tracker",
			Fields: fields_update_tracker,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.UpdateTrackerInput{}
				if _, err := leanruntime.ApplyInput(input, fields_update_tracker, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.UpdateTracker(ctx, input)
			},
		},
		"verify-device-position": {
			Name:   "verify-device-position",
			Fields: fields_verify_device_position,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.VerifyDevicePositionInput{}
				if _, err := leanruntime.ApplyInput(input, fields_verify_device_position, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.VerifyDevicePosition(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("location", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
