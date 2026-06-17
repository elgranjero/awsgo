package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/georoutes"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// georoutesCmd represents the georoutes command
var _georoutesCmd = &cobra.Command{
	Use:   "georoutes",
	Short: "AWS georoutes CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := georoutes.NewFromConfig(cfg)
		if _georoutesCalculateIsolines {
			georoutes_CalculateIsolines(cfg, client)
			return
		}
		if _georoutesCalculateRouteMatrix {
			georoutes_CalculateRouteMatrix(cfg, client)
			return
		}
		if _georoutesCalculateRoutes {
			georoutes_CalculateRoutes(cfg, client)
			return
		}
		if _georoutesOptimizeWaypoints {
			georoutes_OptimizeWaypoints(cfg, client)
			return
		}
		if _georoutesSnapToRoads {
			georoutes_SnapToRoads(cfg, client)
			return
		}

	},
}

var (
	_georoutesCalculateIsolines    bool
	_georoutesCalculateRouteMatrix bool
	_georoutesCalculateRoutes      bool
	_georoutesOptimizeWaypoints    bool
	_georoutesSnapToRoads          bool

	_georoutesAllow                         string
	_georoutesArrivalTime                   string
	_georoutesAvoid                         string
	_georoutesClustering                    string
	_georoutesDepartNow                     string
	_georoutesDepartureTime                 string
	_georoutesDestination                   string
	_georoutesDestinationOptions            string
	_georoutesDestinations                  string
	_georoutesDriver                        string
	_georoutesExclude                       string
	_georoutesInstructionsMeasurementSystem string
	_georoutesIsolineGeometryFormat         string
	_georoutesIsolineGranularity            string
	_georoutesKey                           string
	_georoutesLanguages                     []string
	_georoutesLegAdditionalFeatures         string
	_georoutesLegGeometryFormat             string
	_georoutesMaxAlternatives               string
	_georoutesOptimizeIsolineFor            string
	_georoutesOptimizeRoutingFor            string
	_georoutesOptimizeSequencingFor         string
	_georoutesOrigin                        string
	_georoutesOriginOptions                 string
	_georoutesOrigins                       string
	_georoutesRoutingBoundary               string
	_georoutesSnapRadius                    string
	_georoutesSnappedGeometryFormat         string
	_georoutesSpanAdditionalFeatures        string
	_georoutesThresholds                    string
	_georoutesTolls                         string
	_georoutesTracePoints                   string
	_georoutesTraffic                       string
	_georoutesTravelMode                    string
	_georoutesTravelModeOptions             string
	_georoutesTravelStepType                string
	_georoutesWaypoints                     string
)

// Use the CalculateIsolines action to find service areas that can be reached in a
// given threshold of time, distance.
func georoutes_CalculateIsolines(cfg aws.Config, client *georoutes.Client) {
	input := &georoutes.CalculateIsolinesInput{
		// Thresholds: *types.IsolineThresholds, // Required
	}

	if len(_georoutesThresholds) > 0 {
		if err := assignInputField(input, "Thresholds", _georoutesThresholds); err != nil {
			log.Errorf("invalid --thresholds: %s", err.Error())
			return
		}
	}
	if len(_georoutesAllow) > 0 {
		if err := assignInputField(input, "Allow", _georoutesAllow); err != nil {
			log.Errorf("invalid --allow: %s", err.Error())
			return
		}
	}
	if len(_georoutesArrivalTime) > 0 {
		input.ArrivalTime = aws.String(_georoutesArrivalTime)
	}
	if len(_georoutesAvoid) > 0 {
		if err := assignInputField(input, "Avoid", _georoutesAvoid); err != nil {
			log.Errorf("invalid --avoid: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartNow) > 0 {
		if err := assignInputField(input, "DepartNow", _georoutesDepartNow); err != nil {
			log.Errorf("invalid --depart-now: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartureTime) > 0 {
		input.DepartureTime = aws.String(_georoutesDepartureTime)
	}
	if len(_georoutesDestination) > 0 {
		if err := assignInputField(input, "Destination", _georoutesDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_georoutesDestinationOptions) > 0 {
		if err := assignInputField(input, "DestinationOptions", _georoutesDestinationOptions); err != nil {
			log.Errorf("invalid --destination-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesIsolineGeometryFormat) > 0 {
		if err := assignInputField(input, "IsolineGeometryFormat", _georoutesIsolineGeometryFormat); err != nil {
			log.Errorf("invalid --isoline-geometry-format: %s", err.Error())
			return
		}
	}
	if len(_georoutesIsolineGranularity) > 0 {
		if err := assignInputField(input, "IsolineGranularity", _georoutesIsolineGranularity); err != nil {
			log.Errorf("invalid --isoline-granularity: %s", err.Error())
			return
		}
	}
	if len(_georoutesKey) > 0 {
		input.Key = aws.String(_georoutesKey)
	}
	if len(_georoutesOptimizeIsolineFor) > 0 {
		if err := assignInputField(input, "OptimizeIsolineFor", _georoutesOptimizeIsolineFor); err != nil {
			log.Errorf("invalid --optimize-isoline-for: %s", err.Error())
			return
		}
	}
	if len(_georoutesOptimizeRoutingFor) > 0 {
		if err := assignInputField(input, "OptimizeRoutingFor", _georoutesOptimizeRoutingFor); err != nil {
			log.Errorf("invalid --optimize-routing-for: %s", err.Error())
			return
		}
	}
	if len(_georoutesOrigin) > 0 {
		if err := assignInputField(input, "Origin", _georoutesOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_georoutesOriginOptions) > 0 {
		if err := assignInputField(input, "OriginOptions", _georoutesOriginOptions); err != nil {
			log.Errorf("invalid --origin-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesTraffic) > 0 {
		if err := assignInputField(input, "Traffic", _georoutesTraffic); err != nil {
			log.Errorf("invalid --traffic: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _georoutesTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelModeOptions) > 0 {
		if err := assignInputField(input, "TravelModeOptions", _georoutesTravelModeOptions); err != nil {
			log.Errorf("invalid --travel-mode-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CalculateIsolines(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use CalculateRouteMatrix to compute results for all pairs of Origins to
// Destinations. Each row corresponds to one entry in Origins. Each entry in the
// row corresponds to the route from that entry in Origins to an entry in
// Destinations positions.
func georoutes_CalculateRouteMatrix(cfg aws.Config, client *georoutes.Client) {
	input := &georoutes.CalculateRouteMatrixInput{
		// Destinations: []types.RouteMatrixDestination, // Required
		// Origins: []types.RouteMatrixOrigin, // Required
		// RoutingBoundary: *types.RouteMatrixBoundary, // Required
	}

	if len(_georoutesDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _georoutesDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_georoutesOrigins) > 0 {
		if err := assignInputField(input, "Origins", _georoutesOrigins); err != nil {
			log.Errorf("invalid --origins: %s", err.Error())
			return
		}
	}
	if len(_georoutesRoutingBoundary) > 0 {
		if err := assignInputField(input, "RoutingBoundary", _georoutesRoutingBoundary); err != nil {
			log.Errorf("invalid --routing-boundary: %s", err.Error())
			return
		}
	}
	if len(_georoutesAllow) > 0 {
		if err := assignInputField(input, "Allow", _georoutesAllow); err != nil {
			log.Errorf("invalid --allow: %s", err.Error())
			return
		}
	}
	if len(_georoutesAvoid) > 0 {
		if err := assignInputField(input, "Avoid", _georoutesAvoid); err != nil {
			log.Errorf("invalid --avoid: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartNow) > 0 {
		if err := assignInputField(input, "DepartNow", _georoutesDepartNow); err != nil {
			log.Errorf("invalid --depart-now: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartureTime) > 0 {
		input.DepartureTime = aws.String(_georoutesDepartureTime)
	}
	if len(_georoutesExclude) > 0 {
		if err := assignInputField(input, "Exclude", _georoutesExclude); err != nil {
			log.Errorf("invalid --exclude: %s", err.Error())
			return
		}
	}
	if len(_georoutesKey) > 0 {
		input.Key = aws.String(_georoutesKey)
	}
	if len(_georoutesOptimizeRoutingFor) > 0 {
		if err := assignInputField(input, "OptimizeRoutingFor", _georoutesOptimizeRoutingFor); err != nil {
			log.Errorf("invalid --optimize-routing-for: %s", err.Error())
			return
		}
	}
	if len(_georoutesTraffic) > 0 {
		if err := assignInputField(input, "Traffic", _georoutesTraffic); err != nil {
			log.Errorf("invalid --traffic: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _georoutesTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelModeOptions) > 0 {
		if err := assignInputField(input, "TravelModeOptions", _georoutesTravelModeOptions); err != nil {
			log.Errorf("invalid --travel-mode-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.CalculateRouteMatrix(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// CalculateRoutes computes routes given the following required parameters: Origin
// and Destination .
func georoutes_CalculateRoutes(cfg aws.Config, client *georoutes.Client) {
	input := &georoutes.CalculateRoutesInput{
		// Destination: []float64, // Required
		// Origin: []float64, // Required
	}

	if len(_georoutesDestination) > 0 {
		if err := assignInputField(input, "Destination", _georoutesDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_georoutesOrigin) > 0 {
		if err := assignInputField(input, "Origin", _georoutesOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_georoutesAllow) > 0 {
		if err := assignInputField(input, "Allow", _georoutesAllow); err != nil {
			log.Errorf("invalid --allow: %s", err.Error())
			return
		}
	}
	if len(_georoutesArrivalTime) > 0 {
		input.ArrivalTime = aws.String(_georoutesArrivalTime)
	}
	if len(_georoutesAvoid) > 0 {
		if err := assignInputField(input, "Avoid", _georoutesAvoid); err != nil {
			log.Errorf("invalid --avoid: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartNow) > 0 {
		if err := assignInputField(input, "DepartNow", _georoutesDepartNow); err != nil {
			log.Errorf("invalid --depart-now: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartureTime) > 0 {
		input.DepartureTime = aws.String(_georoutesDepartureTime)
	}
	if len(_georoutesDestinationOptions) > 0 {
		if err := assignInputField(input, "DestinationOptions", _georoutesDestinationOptions); err != nil {
			log.Errorf("invalid --destination-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesDriver) > 0 {
		if err := assignInputField(input, "Driver", _georoutesDriver); err != nil {
			log.Errorf("invalid --driver: %s", err.Error())
			return
		}
	}
	if len(_georoutesExclude) > 0 {
		if err := assignInputField(input, "Exclude", _georoutesExclude); err != nil {
			log.Errorf("invalid --exclude: %s", err.Error())
			return
		}
	}
	if len(_georoutesInstructionsMeasurementSystem) > 0 {
		if err := assignInputField(input, "InstructionsMeasurementSystem", _georoutesInstructionsMeasurementSystem); err != nil {
			log.Errorf("invalid --instructions-measurement-system: %s", err.Error())
			return
		}
	}
	if len(_georoutesKey) > 0 {
		input.Key = aws.String(_georoutesKey)
	}
	if len(_georoutesLanguages) > 0 {
		input.Languages = append([]string(nil), _georoutesLanguages...)
	}
	if len(_georoutesLegAdditionalFeatures) > 0 {
		if err := assignInputField(input, "LegAdditionalFeatures", _georoutesLegAdditionalFeatures); err != nil {
			log.Errorf("invalid --leg-additional-features: %s", err.Error())
			return
		}
	}
	if len(_georoutesLegGeometryFormat) > 0 {
		if err := assignInputField(input, "LegGeometryFormat", _georoutesLegGeometryFormat); err != nil {
			log.Errorf("invalid --leg-geometry-format: %s", err.Error())
			return
		}
	}
	if len(_georoutesMaxAlternatives) > 0 {
		if err := assignInputField(input, "MaxAlternatives", _georoutesMaxAlternatives); err != nil {
			log.Errorf("invalid --max-alternatives: %s", err.Error())
			return
		}
	}
	if len(_georoutesOptimizeRoutingFor) > 0 {
		if err := assignInputField(input, "OptimizeRoutingFor", _georoutesOptimizeRoutingFor); err != nil {
			log.Errorf("invalid --optimize-routing-for: %s", err.Error())
			return
		}
	}
	if len(_georoutesOriginOptions) > 0 {
		if err := assignInputField(input, "OriginOptions", _georoutesOriginOptions); err != nil {
			log.Errorf("invalid --origin-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesSpanAdditionalFeatures) > 0 {
		if err := assignInputField(input, "SpanAdditionalFeatures", _georoutesSpanAdditionalFeatures); err != nil {
			log.Errorf("invalid --span-additional-features: %s", err.Error())
			return
		}
	}
	if len(_georoutesTolls) > 0 {
		if err := assignInputField(input, "Tolls", _georoutesTolls); err != nil {
			log.Errorf("invalid --tolls: %s", err.Error())
			return
		}
	}
	if len(_georoutesTraffic) > 0 {
		if err := assignInputField(input, "Traffic", _georoutesTraffic); err != nil {
			log.Errorf("invalid --traffic: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _georoutesTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelModeOptions) > 0 {
		if err := assignInputField(input, "TravelModeOptions", _georoutesTravelModeOptions); err != nil {
			log.Errorf("invalid --travel-mode-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelStepType) > 0 {
		if err := assignInputField(input, "TravelStepType", _georoutesTravelStepType); err != nil {
			log.Errorf("invalid --travel-step-type: %s", err.Error())
			return
		}
	}
	if len(_georoutesWaypoints) > 0 {
		if err := assignInputField(input, "Waypoints", _georoutesWaypoints); err != nil {
			log.Errorf("invalid --waypoints: %s", err.Error())
			return
		}
	}

	if resp, err := client.CalculateRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// OptimizeWaypoints calculates the optimal order to travel between a set of
// waypoints to minimize either the travel time or the distance travelled during
// the journey, based on road network restrictions and the traffic pattern data.
func georoutes_OptimizeWaypoints(cfg aws.Config, client *georoutes.Client) {
	input := &georoutes.OptimizeWaypointsInput{
		// Origin: []float64, // Required
	}

	if len(_georoutesOrigin) > 0 {
		if err := assignInputField(input, "Origin", _georoutesOrigin); err != nil {
			log.Errorf("invalid --origin: %s", err.Error())
			return
		}
	}
	if len(_georoutesAvoid) > 0 {
		if err := assignInputField(input, "Avoid", _georoutesAvoid); err != nil {
			log.Errorf("invalid --avoid: %s", err.Error())
			return
		}
	}
	if len(_georoutesClustering) > 0 {
		if err := assignInputField(input, "Clustering", _georoutesClustering); err != nil {
			log.Errorf("invalid --clustering: %s", err.Error())
			return
		}
	}
	if len(_georoutesDepartureTime) > 0 {
		input.DepartureTime = aws.String(_georoutesDepartureTime)
	}
	if len(_georoutesDestination) > 0 {
		if err := assignInputField(input, "Destination", _georoutesDestination); err != nil {
			log.Errorf("invalid --destination: %s", err.Error())
			return
		}
	}
	if len(_georoutesDestinationOptions) > 0 {
		if err := assignInputField(input, "DestinationOptions", _georoutesDestinationOptions); err != nil {
			log.Errorf("invalid --destination-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesDriver) > 0 {
		if err := assignInputField(input, "Driver", _georoutesDriver); err != nil {
			log.Errorf("invalid --driver: %s", err.Error())
			return
		}
	}
	if len(_georoutesExclude) > 0 {
		if err := assignInputField(input, "Exclude", _georoutesExclude); err != nil {
			log.Errorf("invalid --exclude: %s", err.Error())
			return
		}
	}
	if len(_georoutesKey) > 0 {
		input.Key = aws.String(_georoutesKey)
	}
	if len(_georoutesOptimizeSequencingFor) > 0 {
		if err := assignInputField(input, "OptimizeSequencingFor", _georoutesOptimizeSequencingFor); err != nil {
			log.Errorf("invalid --optimize-sequencing-for: %s", err.Error())
			return
		}
	}
	if len(_georoutesOriginOptions) > 0 {
		if err := assignInputField(input, "OriginOptions", _georoutesOriginOptions); err != nil {
			log.Errorf("invalid --origin-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesTraffic) > 0 {
		if err := assignInputField(input, "Traffic", _georoutesTraffic); err != nil {
			log.Errorf("invalid --traffic: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _georoutesTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelModeOptions) > 0 {
		if err := assignInputField(input, "TravelModeOptions", _georoutesTravelModeOptions); err != nil {
			log.Errorf("invalid --travel-mode-options: %s", err.Error())
			return
		}
	}
	if len(_georoutesWaypoints) > 0 {
		if err := assignInputField(input, "Waypoints", _georoutesWaypoints); err != nil {
			log.Errorf("invalid --waypoints: %s", err.Error())
			return
		}
	}

	if resp, err := client.OptimizeWaypoints(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// SnapToRoads matches GPS trace to roads most likely traveled on.
func georoutes_SnapToRoads(cfg aws.Config, client *georoutes.Client) {
	input := &georoutes.SnapToRoadsInput{
		// TracePoints: []types.RoadSnapTracePoint, // Required
	}

	if len(_georoutesTracePoints) > 0 {
		if err := assignInputField(input, "TracePoints", _georoutesTracePoints); err != nil {
			log.Errorf("invalid --trace-points: %s", err.Error())
			return
		}
	}
	if len(_georoutesKey) > 0 {
		input.Key = aws.String(_georoutesKey)
	}
	if len(_georoutesSnapRadius) > 0 {
		if err := assignInputField(input, "SnapRadius", _georoutesSnapRadius); err != nil {
			log.Errorf("invalid --snap-radius: %s", err.Error())
			return
		}
	}
	if len(_georoutesSnappedGeometryFormat) > 0 {
		if err := assignInputField(input, "SnappedGeometryFormat", _georoutesSnappedGeometryFormat); err != nil {
			log.Errorf("invalid --snapped-geometry-format: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _georoutesTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_georoutesTravelModeOptions) > 0 {
		if err := assignInputField(input, "TravelModeOptions", _georoutesTravelModeOptions); err != nil {
			log.Errorf("invalid --travel-mode-options: %s", err.Error())
			return
		}
	}

	if resp, err := client.SnapToRoads(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_georoutesCmd)
	_georoutesCmd.Flags().SortFlags = false

	_georoutesCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_georoutesCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_georoutesCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_georoutesCmd.Flags().StringVarP(&_georoutesAllow, "allow", "", "", "Allow")
	_georoutesCmd.Flags().StringVarP(&_georoutesArrivalTime, "arrival-time", "", "", "Arrival Time")
	_georoutesCmd.Flags().StringVarP(&_georoutesAvoid, "avoid", "", "", "Avoid")
	_georoutesCmd.Flags().StringVarP(&_georoutesClustering, "clustering", "", "", "Clustering")
	_georoutesCmd.Flags().StringVarP(&_georoutesDepartNow, "depart-now", "", "", "Depart Now")
	_georoutesCmd.Flags().StringVarP(&_georoutesDepartureTime, "departure-time", "", "", "Departure Time")
	_georoutesCmd.Flags().StringVarP(&_georoutesDestination, "destination", "", "", "Destination")
	_georoutesCmd.Flags().StringVarP(&_georoutesDestinationOptions, "destination-options", "", "", "Destination Options")
	_georoutesCmd.Flags().StringVarP(&_georoutesDestinations, "destinations", "", "", "Destinations")
	_georoutesCmd.Flags().StringVarP(&_georoutesDriver, "driver", "", "", "Driver")
	_georoutesCmd.Flags().StringVarP(&_georoutesExclude, "exclude", "", "", "Exclude")
	_georoutesCmd.Flags().StringVarP(&_georoutesInstructionsMeasurementSystem, "instructions-measurement-system", "", "", "Instructions Measurement System")
	_georoutesCmd.Flags().StringVarP(&_georoutesIsolineGeometryFormat, "isoline-geometry-format", "", "", "Isoline Geometry Format")
	_georoutesCmd.Flags().StringVarP(&_georoutesIsolineGranularity, "isoline-granularity", "", "", "Isoline Granularity")
	_georoutesCmd.Flags().StringVarP(&_georoutesKey, "key", "", "", "Key")
	_georoutesCmd.Flags().StringSliceVarP(&_georoutesLanguages, "languages", "", nil, "Languages")
	_georoutesCmd.Flags().StringVarP(&_georoutesLegAdditionalFeatures, "leg-additional-features", "", "", "Leg Additional Features")
	_georoutesCmd.Flags().StringVarP(&_georoutesLegGeometryFormat, "leg-geometry-format", "", "", "Leg Geometry Format")
	_georoutesCmd.Flags().StringVarP(&_georoutesMaxAlternatives, "max-alternatives", "", "", "Max Alternatives")
	_georoutesCmd.Flags().StringVarP(&_georoutesOptimizeIsolineFor, "optimize-isoline-for", "", "", "Optimize Isoline For")
	_georoutesCmd.Flags().StringVarP(&_georoutesOptimizeRoutingFor, "optimize-routing-for", "", "", "Optimize Routing For")
	_georoutesCmd.Flags().StringVarP(&_georoutesOptimizeSequencingFor, "optimize-sequencing-for", "", "", "Optimize Sequencing For")
	_georoutesCmd.Flags().StringVarP(&_georoutesOrigin, "origin", "", "", "Origin")
	_georoutesCmd.Flags().StringVarP(&_georoutesOriginOptions, "origin-options", "", "", "Origin Options")
	_georoutesCmd.Flags().StringVarP(&_georoutesOrigins, "origins", "", "", "Origins")
	_georoutesCmd.Flags().StringVarP(&_georoutesRoutingBoundary, "routing-boundary", "", "", "Routing Boundary")
	_georoutesCmd.Flags().StringVarP(&_georoutesSnapRadius, "snap-radius", "", "", "Snap Radius")
	_georoutesCmd.Flags().StringVarP(&_georoutesSnappedGeometryFormat, "snapped-geometry-format", "", "", "Snapped Geometry Format")
	_georoutesCmd.Flags().StringVarP(&_georoutesSpanAdditionalFeatures, "span-additional-features", "", "", "Span Additional Features")
	_georoutesCmd.Flags().StringVarP(&_georoutesThresholds, "thresholds", "", "", "Thresholds")
	_georoutesCmd.Flags().StringVarP(&_georoutesTolls, "tolls", "", "", "Tolls")
	_georoutesCmd.Flags().StringVarP(&_georoutesTracePoints, "trace-points", "", "", "Trace Points")
	_georoutesCmd.Flags().StringVarP(&_georoutesTraffic, "traffic", "", "", "Traffic")
	_georoutesCmd.Flags().StringVarP(&_georoutesTravelMode, "travel-mode", "", "", "Travel Mode")
	_georoutesCmd.Flags().StringVarP(&_georoutesTravelModeOptions, "travel-mode-options", "", "", "Travel Mode Options")
	_georoutesCmd.Flags().StringVarP(&_georoutesTravelStepType, "travel-step-type", "", "", "Travel Step Type")
	_georoutesCmd.Flags().StringVarP(&_georoutesWaypoints, "waypoints", "", "", "Waypoints")

	_georoutesCmd.Flags().BoolVarP(&_georoutesCalculateIsolines, "calculate-isolines", "", false, "Calculate Isolines")
	_georoutesCmd.Flags().BoolVarP(&_georoutesCalculateRouteMatrix, "calculate-route-matrix", "", false, "Calculate Route Matrix")
	_georoutesCmd.Flags().BoolVarP(&_georoutesCalculateRoutes, "calculate-routes", "", false, "Calculate Routes")
	_georoutesCmd.Flags().BoolVarP(&_georoutesOptimizeWaypoints, "optimize-waypoints", "", false, "Optimize Waypoints")
	_georoutesCmd.Flags().BoolVarP(&_georoutesSnapToRoads, "snap-to-roads", "", false, "Snap To Roads")

}
