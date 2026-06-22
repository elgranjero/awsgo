package main

import (
	"context"
	"fmt"
	"os"

	"aws/awsgo/leanruntime"
	"github.com/aws/aws-sdk-go-v2/aws"
	svc "github.com/aws/aws-sdk-go-v2/service/georoutes"
)

var fields_calculate_isolines = []leanruntime.Field{
	{Name: "Allow", Flag: "allow", Type: "*types.IsolineAllowOptions", Required: false},
	{Name: "ArrivalTime", Flag: "arrival-time", Type: "*string", Required: false},
	{Name: "Avoid", Flag: "avoid", Type: "*types.IsolineAvoidanceOptions", Required: false},
	{Name: "DepartNow", Flag: "depart-now", Type: "*bool", Required: false},
	{Name: "DepartureTime", Flag: "departure-time", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "[]float64", Required: false},
	{Name: "DestinationOptions", Flag: "destination-options", Type: "*types.IsolineDestinationOptions", Required: false},
	{Name: "IsolineGeometryFormat", Flag: "isoline-geometry-format", Type: "types.GeometryFormat", Required: false},
	{Name: "IsolineGranularity", Flag: "isoline-granularity", Type: "*types.IsolineGranularityOptions", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "OptimizeIsolineFor", Flag: "optimize-isoline-for", Type: "types.IsolineOptimizationObjective", Required: false},
	{Name: "OptimizeRoutingFor", Flag: "optimize-routing-for", Type: "types.RoutingObjective", Required: false},
	{Name: "Origin", Flag: "origin", Type: "[]float64", Required: false},
	{Name: "OriginOptions", Flag: "origin-options", Type: "*types.IsolineOriginOptions", Required: false},
	{Name: "Thresholds", Flag: "thresholds", Type: "*types.IsolineThresholds", Required: true},
	{Name: "Traffic", Flag: "traffic", Type: "*types.IsolineTrafficOptions", Required: false},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.IsolineTravelMode", Required: false},
	{Name: "TravelModeOptions", Flag: "travel-mode-options", Type: "*types.IsolineTravelModeOptions", Required: false},
}

var fields_calculate_route_matrix = []leanruntime.Field{
	{Name: "Allow", Flag: "allow", Type: "*types.RouteMatrixAllowOptions", Required: false},
	{Name: "Avoid", Flag: "avoid", Type: "*types.RouteMatrixAvoidanceOptions", Required: false},
	{Name: "DepartNow", Flag: "depart-now", Type: "*bool", Required: false},
	{Name: "DepartureTime", Flag: "departure-time", Type: "*string", Required: false},
	{Name: "Destinations", Flag: "destinations", Type: "[]types.RouteMatrixDestination", Required: true},
	{Name: "Exclude", Flag: "exclude", Type: "*types.RouteMatrixExclusionOptions", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "OptimizeRoutingFor", Flag: "optimize-routing-for", Type: "types.RoutingObjective", Required: false},
	{Name: "Origins", Flag: "origins", Type: "[]types.RouteMatrixOrigin", Required: true},
	{Name: "RoutingBoundary", Flag: "routing-boundary", Type: "*types.RouteMatrixBoundary", Required: true},
	{Name: "Traffic", Flag: "traffic", Type: "*types.RouteMatrixTrafficOptions", Required: false},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.RouteMatrixTravelMode", Required: false},
	{Name: "TravelModeOptions", Flag: "travel-mode-options", Type: "*types.RouteMatrixTravelModeOptions", Required: false},
}

var fields_calculate_routes = []leanruntime.Field{
	{Name: "Allow", Flag: "allow", Type: "*types.RouteAllowOptions", Required: false},
	{Name: "ArrivalTime", Flag: "arrival-time", Type: "*string", Required: false},
	{Name: "Avoid", Flag: "avoid", Type: "*types.RouteAvoidanceOptions", Required: false},
	{Name: "DepartNow", Flag: "depart-now", Type: "*bool", Required: false},
	{Name: "DepartureTime", Flag: "departure-time", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "[]float64", Required: true},
	{Name: "DestinationOptions", Flag: "destination-options", Type: "*types.RouteDestinationOptions", Required: false},
	{Name: "Driver", Flag: "driver", Type: "*types.RouteDriverOptions", Required: false},
	{Name: "Exclude", Flag: "exclude", Type: "*types.RouteExclusionOptions", Required: false},
	{Name: "InstructionsMeasurementSystem", Flag: "instructions-measurement-system", Type: "types.MeasurementSystem", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "Languages", Flag: "languages", Type: "[]string", Required: false},
	{Name: "LegAdditionalFeatures", Flag: "leg-additional-features", Type: "[]types.RouteLegAdditionalFeature", Required: false},
	{Name: "LegGeometryFormat", Flag: "leg-geometry-format", Type: "types.GeometryFormat", Required: false},
	{Name: "MaxAlternatives", Flag: "max-alternatives", Type: "*int32", Required: false},
	{Name: "OptimizeRoutingFor", Flag: "optimize-routing-for", Type: "types.RoutingObjective", Required: false},
	{Name: "Origin", Flag: "origin", Type: "[]float64", Required: true},
	{Name: "OriginOptions", Flag: "origin-options", Type: "*types.RouteOriginOptions", Required: false},
	{Name: "SpanAdditionalFeatures", Flag: "span-additional-features", Type: "[]types.RouteSpanAdditionalFeature", Required: false},
	{Name: "Tolls", Flag: "tolls", Type: "*types.RouteTollOptions", Required: false},
	{Name: "Traffic", Flag: "traffic", Type: "*types.RouteTrafficOptions", Required: false},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.RouteTravelMode", Required: false},
	{Name: "TravelModeOptions", Flag: "travel-mode-options", Type: "*types.RouteTravelModeOptions", Required: false},
	{Name: "TravelStepType", Flag: "travel-step-type", Type: "types.RouteTravelStepType", Required: false},
	{Name: "Waypoints", Flag: "waypoints", Type: "[]types.RouteWaypoint", Required: false},
}

var fields_optimize_waypoints = []leanruntime.Field{
	{Name: "Avoid", Flag: "avoid", Type: "*types.WaypointOptimizationAvoidanceOptions", Required: false},
	{Name: "Clustering", Flag: "clustering", Type: "*types.WaypointOptimizationClusteringOptions", Required: false},
	{Name: "DepartureTime", Flag: "departure-time", Type: "*string", Required: false},
	{Name: "Destination", Flag: "destination", Type: "[]float64", Required: false},
	{Name: "DestinationOptions", Flag: "destination-options", Type: "*types.WaypointOptimizationDestinationOptions", Required: false},
	{Name: "Driver", Flag: "driver", Type: "*types.WaypointOptimizationDriverOptions", Required: false},
	{Name: "Exclude", Flag: "exclude", Type: "*types.WaypointOptimizationExclusionOptions", Required: false},
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "OptimizeSequencingFor", Flag: "optimize-sequencing-for", Type: "types.WaypointOptimizationSequencingObjective", Required: false},
	{Name: "Origin", Flag: "origin", Type: "[]float64", Required: true},
	{Name: "OriginOptions", Flag: "origin-options", Type: "*types.WaypointOptimizationOriginOptions", Required: false},
	{Name: "Traffic", Flag: "traffic", Type: "*types.WaypointOptimizationTrafficOptions", Required: false},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.WaypointOptimizationTravelMode", Required: false},
	{Name: "TravelModeOptions", Flag: "travel-mode-options", Type: "*types.WaypointOptimizationTravelModeOptions", Required: false},
	{Name: "Waypoints", Flag: "waypoints", Type: "[]types.WaypointOptimizationWaypoint", Required: false},
}

var fields_snap_to_roads = []leanruntime.Field{
	{Name: "Key", Flag: "key", Type: "*string", Required: false},
	{Name: "SnapRadius", Flag: "snap-radius", Type: "int64", Required: false},
	{Name: "SnappedGeometryFormat", Flag: "snapped-geometry-format", Type: "types.GeometryFormat", Required: false},
	{Name: "TracePoints", Flag: "trace-points", Type: "[]types.RoadSnapTracePoint", Required: true},
	{Name: "TravelMode", Flag: "travel-mode", Type: "types.RoadSnapTravelMode", Required: false},
	{Name: "TravelModeOptions", Flag: "travel-mode-options", Type: "*types.RoadSnapTravelModeOptions", Required: false},
}

func main() {
	ops := map[string]leanruntime.Operation{
		"calculate-isolines": {
			Name:   "calculate-isolines",
			Fields: fields_calculate_isolines,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CalculateIsolinesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_calculate_isolines, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CalculateIsolines(ctx, input)
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
		"calculate-routes": {
			Name:   "calculate-routes",
			Fields: fields_calculate_routes,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.CalculateRoutesInput{}
				if _, err := leanruntime.ApplyInput(input, fields_calculate_routes, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.CalculateRoutes(ctx, input)
			},
		},
		"optimize-waypoints": {
			Name:   "optimize-waypoints",
			Fields: fields_optimize_waypoints,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.OptimizeWaypointsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_optimize_waypoints, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.OptimizeWaypoints(ctx, input)
			},
		},
		"snap-to-roads": {
			Name:   "snap-to-roads",
			Fields: fields_snap_to_roads,
			Run: func(ctx context.Context, cfg aws.Config, values leanruntime.Values) (any, error) {
				input := &svc.SnapToRoadsInput{}
				if _, err := leanruntime.ApplyInput(input, fields_snap_to_roads, values); err != nil {
					return nil, err
				}
				client := svc.NewFromConfig(cfg)
				return client.SnapToRoads(ctx, input)
			},
		},
	}
	if err := leanruntime.Execute("georoutes", ops, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
