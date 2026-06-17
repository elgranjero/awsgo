package main

import (
	"fmt"
	"os"

	runtime "aws/awsgo/svcruntime"
	servicecmd "aws/generated/georoutes/cmd"
)

func main() {
	svc := runtime.ServiceDef{
		Operations:   []string{"calculate-isolines", "calculate-route-matrix", "calculate-routes", "optimize-waypoints", "snap-to-roads"},
		OperationSet: map[string]bool{"calculate-isolines": true, "calculate-route-matrix": true, "calculate-routes": true, "optimize-waypoints": true, "snap-to-roads": true},
		OperationInputs: map[string][]string{
			"calculate-isolines":     {"Allow", "ArrivalTime", "Avoid", "DepartNow", "DepartureTime", "Destination", "DestinationOptions", "IsolineGeometryFormat", "IsolineGranularity", "Key", "OptimizeIsolineFor", "OptimizeRoutingFor", "Origin", "OriginOptions", "Thresholds", "Traffic", "TravelMode", "TravelModeOptions"},
			"calculate-route-matrix": {"Allow", "Avoid", "DepartNow", "DepartureTime", "Destinations", "Exclude", "Key", "OptimizeRoutingFor", "Origins", "RoutingBoundary", "Traffic", "TravelMode", "TravelModeOptions"},
			"calculate-routes":       {"Allow", "ArrivalTime", "Avoid", "DepartNow", "DepartureTime", "Destination", "DestinationOptions", "Driver", "Exclude", "InstructionsMeasurementSystem", "Key", "Languages", "LegAdditionalFeatures", "LegGeometryFormat", "MaxAlternatives", "OptimizeRoutingFor", "Origin", "OriginOptions", "SpanAdditionalFeatures", "Tolls", "Traffic", "TravelMode", "TravelModeOptions", "TravelStepType", "Waypoints"},
			"optimize-waypoints":     {"Avoid", "Clustering", "DepartureTime", "Destination", "DestinationOptions", "Driver", "Exclude", "Key", "OptimizeSequencingFor", "Origin", "OriginOptions", "Traffic", "TravelMode", "TravelModeOptions", "Waypoints"},
			"snap-to-roads":          {"Key", "SnapRadius", "SnappedGeometryFormat", "TracePoints", "TravelMode", "TravelModeOptions"},
		},
		OperationInputTypes: map[string]map[string]string{
			"calculate-isolines":     {"Allow": "*types.IsolineAllowOptions", "ArrivalTime": "*string", "Avoid": "*types.IsolineAvoidanceOptions", "DepartNow": "*bool", "DepartureTime": "*string", "Destination": "[]float64", "DestinationOptions": "*types.IsolineDestinationOptions", "IsolineGeometryFormat": "types.GeometryFormat", "IsolineGranularity": "*types.IsolineGranularityOptions", "Key": "*string", "OptimizeIsolineFor": "types.IsolineOptimizationObjective", "OptimizeRoutingFor": "types.RoutingObjective", "Origin": "[]float64", "OriginOptions": "*types.IsolineOriginOptions", "Thresholds": "*types.IsolineThresholds", "Traffic": "*types.IsolineTrafficOptions", "TravelMode": "types.IsolineTravelMode", "TravelModeOptions": "*types.IsolineTravelModeOptions"},
			"calculate-route-matrix": {"Allow": "*types.RouteMatrixAllowOptions", "Avoid": "*types.RouteMatrixAvoidanceOptions", "DepartNow": "*bool", "DepartureTime": "*string", "Destinations": "[]types.RouteMatrixDestination", "Exclude": "*types.RouteMatrixExclusionOptions", "Key": "*string", "OptimizeRoutingFor": "types.RoutingObjective", "Origins": "[]types.RouteMatrixOrigin", "RoutingBoundary": "*types.RouteMatrixBoundary", "Traffic": "*types.RouteMatrixTrafficOptions", "TravelMode": "types.RouteMatrixTravelMode", "TravelModeOptions": "*types.RouteMatrixTravelModeOptions"},
			"calculate-routes":       {"Allow": "*types.RouteAllowOptions", "ArrivalTime": "*string", "Avoid": "*types.RouteAvoidanceOptions", "DepartNow": "*bool", "DepartureTime": "*string", "Destination": "[]float64", "DestinationOptions": "*types.RouteDestinationOptions", "Driver": "*types.RouteDriverOptions", "Exclude": "*types.RouteExclusionOptions", "InstructionsMeasurementSystem": "types.MeasurementSystem", "Key": "*string", "Languages": "[]string", "LegAdditionalFeatures": "[]types.RouteLegAdditionalFeature", "LegGeometryFormat": "types.GeometryFormat", "MaxAlternatives": "*int32", "OptimizeRoutingFor": "types.RoutingObjective", "Origin": "[]float64", "OriginOptions": "*types.RouteOriginOptions", "SpanAdditionalFeatures": "[]types.RouteSpanAdditionalFeature", "Tolls": "*types.RouteTollOptions", "Traffic": "*types.RouteTrafficOptions", "TravelMode": "types.RouteTravelMode", "TravelModeOptions": "*types.RouteTravelModeOptions", "TravelStepType": "types.RouteTravelStepType", "Waypoints": "[]types.RouteWaypoint"},
			"optimize-waypoints":     {"Avoid": "*types.WaypointOptimizationAvoidanceOptions", "Clustering": "*types.WaypointOptimizationClusteringOptions", "DepartureTime": "*string", "Destination": "[]float64", "DestinationOptions": "*types.WaypointOptimizationDestinationOptions", "Driver": "*types.WaypointOptimizationDriverOptions", "Exclude": "*types.WaypointOptimizationExclusionOptions", "Key": "*string", "OptimizeSequencingFor": "types.WaypointOptimizationSequencingObjective", "Origin": "[]float64", "OriginOptions": "*types.WaypointOptimizationOriginOptions", "Traffic": "*types.WaypointOptimizationTrafficOptions", "TravelMode": "types.WaypointOptimizationTravelMode", "TravelModeOptions": "*types.WaypointOptimizationTravelModeOptions", "Waypoints": "[]types.WaypointOptimizationWaypoint"},
			"snap-to-roads":          {"Key": "*string", "SnapRadius": "int64", "SnappedGeometryFormat": "types.GeometryFormat", "TracePoints": "[]types.RoadSnapTracePoint", "TravelMode": "types.RoadSnapTravelMode", "TravelModeOptions": "*types.RoadSnapTravelModeOptions"},
		},
		OperationInputRequired: map[string][]string{
			"calculate-isolines":     {"Thresholds"},
			"calculate-route-matrix": {"Destinations", "Origins", "RoutingBoundary"},
			"calculate-routes":       {"Destination", "Origin"},
			"optimize-waypoints":     {"Origin"},
			"snap-to-roads":          {"TracePoints"},
		},
		Run: servicecmd.Execute,
	}
	if err := runtime.ExecuteService("georoutes", svc, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
