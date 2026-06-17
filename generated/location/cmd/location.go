package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/location"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// locationCmd represents the location command
var _locationCmd = &cobra.Command{
	Use:   "location",
	Short: "AWS location CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := location.NewFromConfig(cfg)
		if _locationAssociateTrackerConsumer {
			location_AssociateTrackerConsumer(cfg, client)
			return
		}
		if _locationBatchDeleteDevicePositionHistory {
			location_BatchDeleteDevicePositionHistory(cfg, client)
			return
		}
		if _locationBatchDeleteGeofence {
			location_BatchDeleteGeofence(cfg, client)
			return
		}
		if _locationBatchEvaluateGeofences {
			location_BatchEvaluateGeofences(cfg, client)
			return
		}
		if _locationBatchGetDevicePosition {
			location_BatchGetDevicePosition(cfg, client)
			return
		}
		if _locationBatchPutGeofence {
			location_BatchPutGeofence(cfg, client)
			return
		}
		if _locationBatchUpdateDevicePosition {
			location_BatchUpdateDevicePosition(cfg, client)
			return
		}
		if _locationCalculateRoute {
			location_CalculateRoute(cfg, client)
			return
		}
		if _locationCalculateRouteMatrix {
			location_CalculateRouteMatrix(cfg, client)
			return
		}
		if _locationCreateGeofenceCollection {
			location_CreateGeofenceCollection(cfg, client)
			return
		}
		if _locationCreateKey {
			location_CreateKey(cfg, client)
			return
		}
		if _locationCreateMap {
			location_CreateMap(cfg, client)
			return
		}
		if _locationCreatePlaceIndex {
			location_CreatePlaceIndex(cfg, client)
			return
		}
		if _locationCreateRouteCalculator {
			location_CreateRouteCalculator(cfg, client)
			return
		}
		if _locationCreateTracker {
			location_CreateTracker(cfg, client)
			return
		}
		if _locationDeleteGeofenceCollection {
			location_DeleteGeofenceCollection(cfg, client)
			return
		}
		if _locationDeleteKey {
			location_DeleteKey(cfg, client)
			return
		}
		if _locationDeleteMap {
			location_DeleteMap(cfg, client)
			return
		}
		if _locationDeletePlaceIndex {
			location_DeletePlaceIndex(cfg, client)
			return
		}
		if _locationDeleteRouteCalculator {
			location_DeleteRouteCalculator(cfg, client)
			return
		}
		if _locationDeleteTracker {
			location_DeleteTracker(cfg, client)
			return
		}
		if _locationDescribeGeofenceCollection {
			location_DescribeGeofenceCollection(cfg, client)
			return
		}
		if _locationDescribeKey {
			location_DescribeKey(cfg, client)
			return
		}
		if _locationDescribeMap {
			location_DescribeMap(cfg, client)
			return
		}
		if _locationDescribePlaceIndex {
			location_DescribePlaceIndex(cfg, client)
			return
		}
		if _locationDescribeRouteCalculator {
			location_DescribeRouteCalculator(cfg, client)
			return
		}
		if _locationDescribeTracker {
			location_DescribeTracker(cfg, client)
			return
		}
		if _locationDisassociateTrackerConsumer {
			location_DisassociateTrackerConsumer(cfg, client)
			return
		}
		if _locationForecastGeofenceEvents {
			location_ForecastGeofenceEvents(cfg, client)
			return
		}
		if _locationGetDevicePosition {
			location_GetDevicePosition(cfg, client)
			return
		}
		if _locationGetDevicePositionHistory {
			location_GetDevicePositionHistory(cfg, client)
			return
		}
		if _locationGetGeofence {
			location_GetGeofence(cfg, client)
			return
		}
		if _locationGetMapGlyphs {
			location_GetMapGlyphs(cfg, client)
			return
		}
		if _locationGetMapSprites {
			location_GetMapSprites(cfg, client)
			return
		}
		if _locationGetMapStyleDescriptor {
			location_GetMapStyleDescriptor(cfg, client)
			return
		}
		if _locationGetMapTile {
			location_GetMapTile(cfg, client)
			return
		}
		if _locationGetPlace {
			location_GetPlace(cfg, client)
			return
		}
		if _locationListDevicePositions {
			location_ListDevicePositions(cfg, client)
			return
		}
		if _locationListGeofenceCollections {
			location_ListGeofenceCollections(cfg, client)
			return
		}
		if _locationListGeofences {
			location_ListGeofences(cfg, client)
			return
		}
		if _locationListKeys {
			location_ListKeys(cfg, client)
			return
		}
		if _locationListMaps {
			location_ListMaps(cfg, client)
			return
		}
		if _locationListPlaceIndexes {
			location_ListPlaceIndexes(cfg, client)
			return
		}
		if _locationListRouteCalculators {
			location_ListRouteCalculators(cfg, client)
			return
		}
		if _locationListTagsForResource {
			location_ListTagsForResource(cfg, client)
			return
		}
		if _locationListTrackerConsumers {
			location_ListTrackerConsumers(cfg, client)
			return
		}
		if _locationListTrackers {
			location_ListTrackers(cfg, client)
			return
		}
		if _locationPutGeofence {
			location_PutGeofence(cfg, client)
			return
		}
		if _locationSearchPlaceIndexForPosition {
			location_SearchPlaceIndexForPosition(cfg, client)
			return
		}
		if _locationSearchPlaceIndexForSuggestions {
			location_SearchPlaceIndexForSuggestions(cfg, client)
			return
		}
		if _locationSearchPlaceIndexForText {
			location_SearchPlaceIndexForText(cfg, client)
			return
		}
		if _locationTagResource {
			location_TagResource(cfg, client)
			return
		}
		if _locationUntagResource {
			location_UntagResource(cfg, client)
			return
		}
		if _locationUpdateGeofenceCollection {
			location_UpdateGeofenceCollection(cfg, client)
			return
		}
		if _locationUpdateKey {
			location_UpdateKey(cfg, client)
			return
		}
		if _locationUpdateMap {
			location_UpdateMap(cfg, client)
			return
		}
		if _locationUpdatePlaceIndex {
			location_UpdatePlaceIndex(cfg, client)
			return
		}
		if _locationUpdateRouteCalculator {
			location_UpdateRouteCalculator(cfg, client)
			return
		}
		if _locationUpdateTracker {
			location_UpdateTracker(cfg, client)
			return
		}
		if _locationVerifyDevicePosition {
			location_VerifyDevicePosition(cfg, client)
			return
		}

	},
}

var (
	_locationAssociateTrackerConsumer         bool
	_locationBatchDeleteDevicePositionHistory bool
	_locationBatchDeleteGeofence              bool
	_locationBatchEvaluateGeofences           bool
	_locationBatchGetDevicePosition           bool
	_locationBatchPutGeofence                 bool
	_locationBatchUpdateDevicePosition        bool
	_locationCalculateRoute                   bool
	_locationCalculateRouteMatrix             bool
	_locationCreateGeofenceCollection         bool
	_locationCreateKey                        bool
	_locationCreateMap                        bool
	_locationCreatePlaceIndex                 bool
	_locationCreateRouteCalculator            bool
	_locationCreateTracker                    bool
	_locationDeleteGeofenceCollection         bool
	_locationDeleteKey                        bool
	_locationDeleteMap                        bool
	_locationDeletePlaceIndex                 bool
	_locationDeleteRouteCalculator            bool
	_locationDeleteTracker                    bool
	_locationDescribeGeofenceCollection       bool
	_locationDescribeKey                      bool
	_locationDescribeMap                      bool
	_locationDescribePlaceIndex               bool
	_locationDescribeRouteCalculator          bool
	_locationDescribeTracker                  bool
	_locationDisassociateTrackerConsumer      bool
	_locationForecastGeofenceEvents           bool
	_locationGetDevicePosition                bool
	_locationGetDevicePositionHistory         bool
	_locationGetGeofence                      bool
	_locationGetMapGlyphs                     bool
	_locationGetMapSprites                    bool
	_locationGetMapStyleDescriptor            bool
	_locationGetMapTile                       bool
	_locationGetPlace                         bool
	_locationListDevicePositions              bool
	_locationListGeofenceCollections          bool
	_locationListGeofences                    bool
	_locationListKeys                         bool
	_locationListMaps                         bool
	_locationListPlaceIndexes                 bool
	_locationListRouteCalculators             bool
	_locationListTagsForResource              bool
	_locationListTrackerConsumers             bool
	_locationListTrackers                     bool
	_locationPutGeofence                      bool
	_locationSearchPlaceIndexForPosition      bool
	_locationSearchPlaceIndexForSuggestions   bool
	_locationSearchPlaceIndexForText          bool
	_locationTagResource                      bool
	_locationUntagResource                    bool
	_locationUpdateGeofenceCollection         bool
	_locationUpdateKey                        bool
	_locationUpdateMap                        bool
	_locationUpdatePlaceIndex                 bool
	_locationUpdateRouteCalculator            bool
	_locationUpdateTracker                    bool
	_locationVerifyDevicePosition             bool

	_locationArrivalTime                   string
	_locationBiasPosition                  string
	_locationCalculatorName                string
	_locationCarModeOptions                string
	_locationCollectionName                string
	_locationConfiguration                 string
	_locationConfigurationUpdate           string
	_locationConsumerArn                   string
	_locationDataSource                    string
	_locationDataSourceConfiguration       string
	_locationDepartNow                     string
	_locationDeparturePosition             string
	_locationDeparturePositions            string
	_locationDepartureTime                 string
	_locationDescription                   string
	_locationDestinationPosition           string
	_locationDestinationPositions          string
	_locationDeviceId                      string
	_locationDeviceIds                     []string
	_locationDevicePositionUpdates         string
	_locationDeviceState                   string
	_locationDistanceUnit                  string
	_locationEndTimeExclusive              string
	_locationEntries                       string
	_locationEventBridgeEnabled            string
	_locationExpireTime                    string
	_locationFileName                      string
	_locationFilter                        string
	_locationFilterBBox                    string
	_locationFilterCategories              []string
	_locationFilterCountries               []string
	_locationFilterGeometry                string
	_locationFontStack                     string
	_locationFontUnicodeRange              string
	_locationForceDelete                   string
	_locationForceUpdate                   string
	_locationGeofenceId                    string
	_locationGeofenceIds                   []string
	_locationGeofenceProperties            string
	_locationGeometry                      string
	_locationIncludeLegGeometry            string
	_locationIndexName                     string
	_locationKey                           string
	_locationKeyName                       string
	_locationKmsKeyEnableGeospatialQueries string
	_locationKmsKeyId                      string
	_locationLanguage                      string
	_locationMapName                       string
	_locationMaxResults                    string
	_locationNextToken                     string
	_locationNoExpiry                      string
	_locationOptimizeFor                   string
	_locationPlaceId                       string
	_locationPosition                      string
	_locationPositionFiltering             string
	_locationPricingPlan                   string
	_locationPricingPlanDataSource         string
	_locationResourceArn                   string
	_locationRestrictions                  string
	_locationSpeedUnit                     string
	_locationStartTimeInclusive            string
	_locationTagKeys                       []string
	_locationTags                          string
	_locationText                          string
	_locationTimeHorizonMinutes            string
	_locationTrackerName                   string
	_locationTravelMode                    string
	_locationTruckModeOptions              string
	_locationUpdates                       string
	_locationWaypointPositions             string
	_locationX                             string
	_locationY                             string
	_locationZ                             string
)

// Creates an association between a geofence collection and a tracker resource.
// This allows the tracker resource to communicate location data to the linked
// geofence collection.
//
// You can associate up to five geofence collections to each tracker resource.
//
// Currently not supported — Cross-account configurations, such as creating
// associations between a tracker resource in one account and a geofence collection
// in another account.
func location_AssociateTrackerConsumer(cfg aws.Config, client *location.Client) {
	input := &location.AssociateTrackerConsumerInput{
		// ConsumerArn: *string, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationConsumerArn) > 0 {
		input.ConsumerArn = aws.String(_locationConsumerArn)
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.AssociateTrackerConsumer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the position history of one or more devices from a tracker resource.
func location_BatchDeleteDevicePositionHistory(cfg aws.Config, client *location.Client) {
	input := &location.BatchDeleteDevicePositionHistoryInput{
		// DeviceIds: []string, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationDeviceIds) > 0 {
		input.DeviceIds = append([]string(nil), _locationDeviceIds...)
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.BatchDeleteDevicePositionHistory(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a batch of geofences from a geofence collection.
// This operation deletes the resource permanently.
func location_BatchDeleteGeofence(cfg aws.Config, client *location.Client) {
	input := &location.BatchDeleteGeofenceInput{
		// CollectionName: *string, // Required
		// GeofenceIds: []string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationGeofenceIds) > 0 {
		input.GeofenceIds = append([]string(nil), _locationGeofenceIds...)
	}

	if resp, err := client.BatchDeleteGeofence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Evaluates device positions against the geofence geometries from a given
// geofence collection.
//
// This operation always returns an empty response because geofences are
// asynchronously evaluated. The evaluation determines if the device has entered or
// exited a geofenced area, and then publishes one of the following events to
// Amazon EventBridge:
//
// - ENTER if Amazon Location determines that the tracked device has entered a
// geofenced area.
//
// - EXIT if Amazon Location determines that the tracked device has exited a
// geofenced area.
//
// The last geofence that a device was observed within is tracked for 30 days
// after the most recent device position update.
//
// Geofence evaluation uses the given device position. It does not account for the
// optional Accuracy of a DevicePositionUpdate .
//
// The DeviceID is used as a string to represent the device. You do not need to
// have a Tracker associated with the DeviceID .
func location_BatchEvaluateGeofences(cfg aws.Config, client *location.Client) {
	input := &location.BatchEvaluateGeofencesInput{
		// CollectionName: *string, // Required
		// DevicePositionUpdates: []types.DevicePositionUpdate, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationDevicePositionUpdates) > 0 {
		if err := assignInputField(input, "DevicePositionUpdates", _locationDevicePositionUpdates); err != nil {
			log.Errorf("invalid --device-position-updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchEvaluateGeofences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists the latest device positions for requested devices.
func location_BatchGetDevicePosition(cfg aws.Config, client *location.Client) {
	input := &location.BatchGetDevicePositionInput{
		// DeviceIds: []string, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationDeviceIds) > 0 {
		input.DeviceIds = append([]string(nil), _locationDeviceIds...)
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.BatchGetDevicePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A batch request for storing geofence geometries into a given geofence
// collection, or updates the geometry of an existing geofence if a geofence ID is
// included in the request.
func location_BatchPutGeofence(cfg aws.Config, client *location.Client) {
	input := &location.BatchPutGeofenceInput{
		// CollectionName: *string, // Required
		// Entries: []types.BatchPutGeofenceRequestEntry, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationEntries) > 0 {
		if err := assignInputField(input, "Entries", _locationEntries); err != nil {
			log.Errorf("invalid --entries: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchPutGeofence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads position update data for one or more devices to a tracker resource (up
// to 10 devices per batch). Amazon Location uses the data when it reports the last
// known device position and position history. Amazon Location retains location
// data for 30 days.
//
// Position updates are handled based on the PositionFiltering property of the
// tracker. When PositionFiltering is set to TimeBased , updates are evaluated
// against linked geofence collections, and location data is stored at a maximum of
// one position per 30 second interval. If your update frequency is more often than
// every 30 seconds, only one update per 30 seconds is stored for each unique
// device ID.
//
// When PositionFiltering is set to DistanceBased filtering, location data is
// stored and evaluated against linked geofence collections only if the device has
// moved more than 30 m (98.4 ft).
//
// When PositionFiltering is set to AccuracyBased filtering, location data is
// stored and evaluated against linked geofence collections only if the device has
// moved more than the measured accuracy. For example, if two consecutive updates
// from a device have a horizontal accuracy of 5 m and 10 m, the second update is
// neither stored or evaluated if the device has moved less than 15 m. If
// PositionFiltering is set to AccuracyBased filtering, Amazon Location uses the
// default value { "Horizontal": 0} when accuracy is not provided on a
// DevicePositionUpdate .
func location_BatchUpdateDevicePosition(cfg aws.Config, client *location.Client) {
	input := &location.BatchUpdateDevicePositionInput{
		// TrackerName: *string, // Required
		// Updates: []types.DevicePositionUpdate, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationUpdates) > 0 {
		if err := assignInputField(input, "Updates", _locationUpdates); err != nil {
			log.Errorf("invalid --updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.BatchUpdateDevicePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to CalculateRoutesCalculateRoutes or CalculateIsolinesCalculateIsolines unless you
// require Grab data.
//
// - CalculateRoute is part of a previous Amazon Location Service Routes API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The version 2 CalculateRoutes operation gives better results for
// point-to-point routing, while the version 2 CalculateIsolines operation adds
// support for calculating service areas and travel time envelopes.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// [Calculates a route]given the following required parameters: DeparturePosition and
// DestinationPosition . Requires that you first [create a route calculator resource].
//
// By default, a request that doesn't specify a departure time uses the best time
// of day to travel with the best traffic conditions when calculating the route.
//
// Additional options include:
//
// [Specifying a departure time]
// - using either DepartureTime or DepartNow . This calculates a route based on
// predictive traffic data at the given time.
//
// You can't specify both DepartureTime and DepartNow in a single request.
//
// Specifying both parameters returns a validation error.
//
// [Specifying a travel mode]
// - using TravelMode sets the transportation mode used to calculate the routes.
// This also lets you specify additional route preferences in CarModeOptions if
// traveling by Car , or TruckModeOptions if traveling by Truck .
//
// # If you specify walking for the travel mode and your data provider is Esri, the
//
// start and destination must be within 40km.
//
// [Specifying a departure time]: https://docs.aws.amazon.com/location/previous/developerguide/departure-time.html
// [Specifying a travel mode]: https://docs.aws.amazon.com/location/previous/developerguide/travel-mode.html
// [Calculates a route]: https://docs.aws.amazon.com/location/previous/developerguide/calculate-route.html
// [create a route calculator resource]: https://docs.aws.amazon.com/location-routes/latest/APIReference/API_CreateRouteCalculator.html
func location_CalculateRoute(cfg aws.Config, client *location.Client) {
	input := &location.CalculateRouteInput{
		// CalculatorName: *string, // Required
		// DeparturePosition: []float64, // Required
		// DestinationPosition: []float64, // Required
	}

	if len(_locationCalculatorName) > 0 {
		input.CalculatorName = aws.String(_locationCalculatorName)
	}
	if len(_locationDeparturePosition) > 0 {
		if err := assignInputField(input, "DeparturePosition", _locationDeparturePosition); err != nil {
			log.Errorf("invalid --departure-position: %s", err.Error())
			return
		}
	}
	if len(_locationDestinationPosition) > 0 {
		if err := assignInputField(input, "DestinationPosition", _locationDestinationPosition); err != nil {
			log.Errorf("invalid --destination-position: %s", err.Error())
			return
		}
	}
	if len(_locationArrivalTime) > 0 {
		if err := assignInputField(input, "ArrivalTime", _locationArrivalTime); err != nil {
			log.Errorf("invalid --arrival-time: %s", err.Error())
			return
		}
	}
	if len(_locationCarModeOptions) > 0 {
		if err := assignInputField(input, "CarModeOptions", _locationCarModeOptions); err != nil {
			log.Errorf("invalid --car-mode-options: %s", err.Error())
			return
		}
	}
	if len(_locationDepartNow) > 0 {
		if err := assignInputField(input, "DepartNow", _locationDepartNow); err != nil {
			log.Errorf("invalid --depart-now: %s", err.Error())
			return
		}
	}
	if len(_locationDepartureTime) > 0 {
		if err := assignInputField(input, "DepartureTime", _locationDepartureTime); err != nil {
			log.Errorf("invalid --departure-time: %s", err.Error())
			return
		}
	}
	if len(_locationDistanceUnit) > 0 {
		if err := assignInputField(input, "DistanceUnit", _locationDistanceUnit); err != nil {
			log.Errorf("invalid --distance-unit: %s", err.Error())
			return
		}
	}
	if len(_locationIncludeLegGeometry) > 0 {
		if err := assignInputField(input, "IncludeLegGeometry", _locationIncludeLegGeometry); err != nil {
			log.Errorf("invalid --include-leg-geometry: %s", err.Error())
			return
		}
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}
	if len(_locationOptimizeFor) > 0 {
		if err := assignInputField(input, "OptimizeFor", _locationOptimizeFor); err != nil {
			log.Errorf("invalid --optimize-for: %s", err.Error())
			return
		}
	}
	if len(_locationTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _locationTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_locationTruckModeOptions) > 0 {
		if err := assignInputField(input, "TruckModeOptions", _locationTruckModeOptions); err != nil {
			log.Errorf("invalid --truck-mode-options: %s", err.Error())
			return
		}
	}
	if len(_locationWaypointPositions) > 0 {
		if err := assignInputField(input, "WaypointPositions", _locationWaypointPositions); err != nil {
			log.Errorf("invalid --waypoint-positions: %s", err.Error())
			return
		}
	}

	if resp, err := client.CalculateRoute(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the V2 CalculateRouteMatrixCalculateRouteMatrix unless you require Grab data.
//
// - This version of CalculateRouteMatrix is part of a previous Amazon Location
// Service Routes API (version 1) which has been superseded by a more intuitive,
// powerful, and complete API (version 2).
//
// - The version 2 CalculateRouteMatrix operation gives better results for matrix
// routing calculations.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Routes V2 API Referenceor the Developer Guide.
//
// [Calculates a route matrix]given the following required parameters: DeparturePositions and
// DestinationPositions . CalculateRouteMatrix calculates routes and returns the
// travel time and travel distance from each departure position to each destination
// position in the request. For example, given departure positions A and B, and
// destination positions X and Y, CalculateRouteMatrix will return time and
// distance for routes from A to X, A to Y, B to X, and B to Y (in that order). The
// number of results returned (and routes calculated) will be the number of
// DeparturePositions times the number of DestinationPositions .
//
// Your account is charged for each route calculated, not the number of requests.
//
// Requires that you first [create a route calculator resource].
//
// By default, a request that doesn't specify a departure time uses the best time
// of day to travel with the best traffic conditions when calculating routes.
//
// Additional options include:
//
// [Specifying a departure time]
// - using either DepartureTime or DepartNow . This calculates routes based on
// predictive traffic data at the given time.
//
// You can't specify both DepartureTime and DepartNow in a single request.
//
// Specifying both parameters returns a validation error.
//
// [Specifying a travel mode]
// - using TravelMode sets the transportation mode used to calculate the routes.
// This also lets you specify additional route preferences in CarModeOptions if
// traveling by Car , or TruckModeOptions if traveling by Truck .
//
// [Specifying a departure time]: https://docs.aws.amazon.com/location/previous/developerguide/departure-time.html
// [Specifying a travel mode]: https://docs.aws.amazon.com/location/previous/developerguide/travel-mode.html
// [Calculates a route matrix]: https://docs.aws.amazon.com/location/previous/developerguide/calculate-route-matrix.html
// [create a route calculator resource]: https://docs.aws.amazon.com/location-routes/latest/APIReference/API_CreateRouteCalculator.html
func location_CalculateRouteMatrix(cfg aws.Config, client *location.Client) {
	input := &location.CalculateRouteMatrixInput{
		// CalculatorName: *string, // Required
		// DeparturePositions: [][]float64, // Required
		// DestinationPositions: [][]float64, // Required
	}

	if len(_locationCalculatorName) > 0 {
		input.CalculatorName = aws.String(_locationCalculatorName)
	}
	if len(_locationDeparturePositions) > 0 {
		if err := assignInputField(input, "DeparturePositions", _locationDeparturePositions); err != nil {
			log.Errorf("invalid --departure-positions: %s", err.Error())
			return
		}
	}
	if len(_locationDestinationPositions) > 0 {
		if err := assignInputField(input, "DestinationPositions", _locationDestinationPositions); err != nil {
			log.Errorf("invalid --destination-positions: %s", err.Error())
			return
		}
	}
	if len(_locationCarModeOptions) > 0 {
		if err := assignInputField(input, "CarModeOptions", _locationCarModeOptions); err != nil {
			log.Errorf("invalid --car-mode-options: %s", err.Error())
			return
		}
	}
	if len(_locationDepartNow) > 0 {
		if err := assignInputField(input, "DepartNow", _locationDepartNow); err != nil {
			log.Errorf("invalid --depart-now: %s", err.Error())
			return
		}
	}
	if len(_locationDepartureTime) > 0 {
		if err := assignInputField(input, "DepartureTime", _locationDepartureTime); err != nil {
			log.Errorf("invalid --departure-time: %s", err.Error())
			return
		}
	}
	if len(_locationDistanceUnit) > 0 {
		if err := assignInputField(input, "DistanceUnit", _locationDistanceUnit); err != nil {
			log.Errorf("invalid --distance-unit: %s", err.Error())
			return
		}
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}
	if len(_locationTravelMode) > 0 {
		if err := assignInputField(input, "TravelMode", _locationTravelMode); err != nil {
			log.Errorf("invalid --travel-mode: %s", err.Error())
			return
		}
	}
	if len(_locationTruckModeOptions) > 0 {
		if err := assignInputField(input, "TruckModeOptions", _locationTruckModeOptions); err != nil {
			log.Errorf("invalid --truck-mode-options: %s", err.Error())
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

// Creates a geofence collection, which manages and stores geofences.
func location_CreateGeofenceCollection(cfg aws.Config, client *location.Client) {
	input := &location.CreateGeofenceCollectionInput{
		// CollectionName: *string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_locationKmsKeyId)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationPricingPlanDataSource) > 0 {
		input.PricingPlanDataSource = aws.String(_locationPricingPlanDataSource)
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateGeofenceCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an API key resource in your Amazon Web Services account, which lets you
// grant actions for Amazon Location resources to the API key bearer.
//
// For more information, see [Use API keys to authenticate] in the Amazon Location Service Developer Guide.
//
// [Use API keys to authenticate]: https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html
func location_CreateKey(cfg aws.Config, client *location.Client) {
	input := &location.CreateKeyInput{
		// KeyName: *string, // Required
		// Restrictions: *types.ApiKeyRestrictions, // Required
	}

	if len(_locationKeyName) > 0 {
		input.KeyName = aws.String(_locationKeyName)
	}
	if len(_locationRestrictions) > 0 {
		if err := assignInputField(input, "Restrictions", _locationRestrictions); err != nil {
			log.Errorf("invalid --restrictions: %s", err.Error())
			return
		}
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationExpireTime) > 0 {
		if err := assignInputField(input, "ExpireTime", _locationExpireTime); err != nil {
			log.Errorf("invalid --expire-time: %s", err.Error())
			return
		}
	}
	if len(_locationNoExpiry) > 0 {
		if err := assignInputField(input, "NoExpiry", _locationNoExpiry); err != nil {
			log.Errorf("invalid --no-expiry: %s", err.Error())
			return
		}
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to the Maps API V2 unless you require Grab data.
//
// - CreateMap is part of a previous Amazon Location Service Maps API (version 1)
// which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The Maps API version 2 has a simplified interface that can be used without
// creating or managing map resources.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Creates a map resource in your Amazon Web Services account, which provides map
// tiles of different styles sourced from global location data providers.
//
// If your application is tracking or routing assets you use in your business,
// such as delivery vehicles or employees, you must not use Esri as your
// geolocation provider. See section 82 of the [Amazon Web Services service terms]for more details.
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
func location_CreateMap(cfg aws.Config, client *location.Client) {
	input := &location.CreateMapInput{
		// Configuration: *types.MapConfiguration, // Required
		// MapName: *string, // Required
	}

	if len(_locationConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _locationConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Places API V2 unless you require Grab data.
//
// - CreatePlaceIndex is part of a previous Amazon Location Service Places API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Places API version 2 has a simplified interface that can be used
// without creating or managing place index resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Creates a place index resource in your Amazon Web Services account. Use a place
// index resource to geocode addresses and other text queries by using the
// SearchPlaceIndexForText operation, and reverse geocode coordinates by using the
// SearchPlaceIndexForPosition operation, and enable autosuggestions by using the
// SearchPlaceIndexForSuggestions operation.
//
// If your application is tracking or routing assets you use in your business,
// such as delivery vehicles or employees, you must not use Esri as your
// geolocation provider. See section 82 of the [Amazon Web Services service terms]for more details.
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms
func location_CreatePlaceIndex(cfg aws.Config, client *location.Client) {
	input := &location.CreatePlaceIndexInput{
		// DataSource: *string, // Required
		// IndexName: *string, // Required
	}

	if len(_locationDataSource) > 0 {
		input.DataSource = aws.String(_locationDataSource)
	}
	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}
	if len(_locationDataSourceConfiguration) > 0 {
		if err := assignInputField(input, "DataSourceConfiguration", _locationDataSourceConfiguration); err != nil {
			log.Errorf("invalid --data-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreatePlaceIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Routes API V2 unless you require Grab data.
//
// - CreateRouteCalculator is part of a previous Amazon Location Service Routes
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Routes API version 2 has a simplified interface that can be used
// without creating or managing route calculator resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Routes V2 API Referenceor the Developer Guide.
//
// Creates a route calculator resource in your Amazon Web Services account.
//
// You can send requests to a route calculator resource to estimate travel time,
// distance, and get directions. A route calculator sources traffic and road
// network data from your chosen data provider.
//
// If your application is tracking or routing assets you use in your business,
// such as delivery vehicles or employees, you must not use Esri as your
// geolocation provider. See section 82 of the [Amazon Web Services service terms]for more details.
//
// [Amazon Web Services service terms]: http://aws.amazon.com/service-terms
func location_CreateRouteCalculator(cfg aws.Config, client *location.Client) {
	input := &location.CreateRouteCalculatorInput{
		// CalculatorName: *string, // Required
		// DataSource: *string, // Required
	}

	if len(_locationCalculatorName) > 0 {
		input.CalculatorName = aws.String(_locationCalculatorName)
	}
	if len(_locationDataSource) > 0 {
		input.DataSource = aws.String(_locationDataSource)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateRouteCalculator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a tracker resource in your Amazon Web Services account, which lets you
// retrieve current and historical location of devices.
func location_CreateTracker(cfg aws.Config, client *location.Client) {
	input := &location.CreateTrackerInput{
		// TrackerName: *string, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationEventBridgeEnabled) > 0 {
		if err := assignInputField(input, "EventBridgeEnabled", _locationEventBridgeEnabled); err != nil {
			log.Errorf("invalid --event-bridge-enabled: %s", err.Error())
			return
		}
	}
	if len(_locationKmsKeyEnableGeospatialQueries) > 0 {
		if err := assignInputField(input, "KmsKeyEnableGeospatialQueries", _locationKmsKeyEnableGeospatialQueries); err != nil {
			log.Errorf("invalid --kms-key-enable-geospatial-queries: %s", err.Error())
			return
		}
	}
	if len(_locationKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_locationKmsKeyId)
	}
	if len(_locationPositionFiltering) > 0 {
		if err := assignInputField(input, "PositionFiltering", _locationPositionFiltering); err != nil {
			log.Errorf("invalid --position-filtering: %s", err.Error())
			return
		}
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationPricingPlanDataSource) > 0 {
		input.PricingPlanDataSource = aws.String(_locationPricingPlanDataSource)
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a geofence collection from your Amazon Web Services account.
// This operation deletes the resource permanently. If the geofence collection is
// the target of a tracker resource, the devices will no longer be monitored.
func location_DeleteGeofenceCollection(cfg aws.Config, client *location.Client) {
	input := &location.DeleteGeofenceCollectionInput{
		// CollectionName: *string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}

	if resp, err := client.DeleteGeofenceCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified API key. The API key must have been deactivated more than
// 90 days previously.
//
// For more information, see [Use API keys to authenticate] in the Amazon Location Service Developer Guide.
//
// [Use API keys to authenticate]: https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html
func location_DeleteKey(cfg aws.Config, client *location.Client) {
	input := &location.DeleteKeyInput{
		// KeyName: *string, // Required
	}

	if len(_locationKeyName) > 0 {
		input.KeyName = aws.String(_locationKeyName)
	}
	if len(_locationForceDelete) > 0 {
		if err := assignInputField(input, "ForceDelete", _locationForceDelete); err != nil {
			log.Errorf("invalid --force-delete: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to the Maps API V2 unless you require Grab data.
//
// - DeleteMap is part of a previous Amazon Location Service Maps API (version 1)
// which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The Maps API version 2 has a simplified interface that can be used without
// creating or managing map resources.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Deletes a map resource from your Amazon Web Services account.
//
// This operation deletes the resource permanently. If the map is being used in an
// application, the map may not render.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
func location_DeleteMap(cfg aws.Config, client *location.Client) {
	input := &location.DeleteMapInput{
		// MapName: *string, // Required
	}

	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}

	if resp, err := client.DeleteMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Places API V2 unless you require Grab data.
//
// - DeletePlaceIndex is part of a previous Amazon Location Service Places API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Places API version 2 has a simplified interface that can be used
// without creating or managing place index resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Deletes a place index resource from your Amazon Web Services account.
//
// This operation deletes the resource permanently.
func location_DeletePlaceIndex(cfg aws.Config, client *location.Client) {
	input := &location.DeletePlaceIndexInput{
		// IndexName: *string, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}

	if resp, err := client.DeletePlaceIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Routes API V2 unless you require Grab data.
//
// - DeleteRouteCalculator is part of a previous Amazon Location Service Routes
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Routes API version 2 has a simplified interface that can be used
// without creating or managing route calculator resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Routes V2 API Referenceor the Developer Guide.
//
// Deletes a route calculator resource from your Amazon Web Services account.
//
// This operation deletes the resource permanently.
func location_DeleteRouteCalculator(cfg aws.Config, client *location.Client) {
	input := &location.DeleteRouteCalculatorInput{
		// CalculatorName: *string, // Required
	}

	if len(_locationCalculatorName) > 0 {
		input.CalculatorName = aws.String(_locationCalculatorName)
	}

	if resp, err := client.DeleteRouteCalculator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a tracker resource from your Amazon Web Services account.
// This operation deletes the resource permanently. If the tracker resource is in
// use, you may encounter an error. Make sure that the target resource isn't a
// dependency for your applications.
func location_DeleteTracker(cfg aws.Config, client *location.Client) {
	input := &location.DeleteTrackerInput{
		// TrackerName: *string, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.DeleteTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the geofence collection details.
func location_DescribeGeofenceCollection(cfg aws.Config, client *location.Client) {
	input := &location.DescribeGeofenceCollectionInput{
		// CollectionName: *string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}

	if resp, err := client.DescribeGeofenceCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the API key resource details.
// For more information, see [Use API keys to authenticate] in the Amazon Location Service Developer Guide.
//
// [Use API keys to authenticate]: https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html
func location_DescribeKey(cfg aws.Config, client *location.Client) {
	input := &location.DescribeKeyInput{
		// KeyName: *string, // Required
	}

	if len(_locationKeyName) > 0 {
		input.KeyName = aws.String(_locationKeyName)
	}

	if resp, err := client.DescribeKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to the Maps API V2 unless you require Grab data.
//
// - DescribeMap is part of a previous Amazon Location Service Maps API (version
// 1) which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The Maps API version 2 has a simplified interface that can be used without
// creating or managing map resources.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Retrieves the map resource details.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
func location_DescribeMap(cfg aws.Config, client *location.Client) {
	input := &location.DescribeMapInput{
		// MapName: *string, // Required
	}

	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}

	if resp, err := client.DescribeMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Places API V2 unless you require Grab data.
//
// - DescribePlaceIndex is part of a previous Amazon Location Service Places API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Places API version 2 has a simplified interface that can be used
// without creating or managing place index resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Retrieves the place index resource details.
func location_DescribePlaceIndex(cfg aws.Config, client *location.Client) {
	input := &location.DescribePlaceIndexInput{
		// IndexName: *string, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}

	if resp, err := client.DescribePlaceIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Routes API V2 unless you require Grab data.
//
// - DescribeRouteCalculator is part of a previous Amazon Location Service Routes
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Routes API version 2 has a simplified interface that can be used
// without creating or managing route calculator resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Routes V2 API Referenceor the Developer Guide.
//
// Retrieves the route calculator resource details.
func location_DescribeRouteCalculator(cfg aws.Config, client *location.Client) {
	input := &location.DescribeRouteCalculatorInput{
		// CalculatorName: *string, // Required
	}

	if len(_locationCalculatorName) > 0 {
		input.CalculatorName = aws.String(_locationCalculatorName)
	}

	if resp, err := client.DescribeRouteCalculator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the tracker resource details.
func location_DescribeTracker(cfg aws.Config, client *location.Client) {
	input := &location.DescribeTrackerInput{
		// TrackerName: *string, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.DescribeTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the association between a tracker resource and a geofence collection.
// Once you unlink a tracker resource from a geofence collection, the tracker
// positions will no longer be automatically evaluated against geofences.
func location_DisassociateTrackerConsumer(cfg aws.Config, client *location.Client) {
	input := &location.DisassociateTrackerConsumerInput{
		// ConsumerArn: *string, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationConsumerArn) > 0 {
		input.ConsumerArn = aws.String(_locationConsumerArn)
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.DisassociateTrackerConsumer(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action forecasts future geofence events that are likely to occur within a
// specified time horizon if a device continues moving at its current speed. Each
// forecasted event is associated with a geofence from a provided geofence
// collection. A forecast event can have one of the following states:
//
// ENTER : The device position is outside the referenced geofence, but the device
// may cross into the geofence during the forecasting time horizon if it maintains
// its current speed.
//
// EXIT : The device position is inside the referenced geofence, but the device may
// leave the geofence during the forecasted time horizon if the device maintains
// it's current speed.
//
// IDLE :The device is inside the geofence, and it will remain inside the geofence
// through the end of the time horizon if the device maintains it's current speed.
//
// Heading direction is not considered in the current version. The API takes a
// conservative approach and includes events that can occur for any heading.
func location_ForecastGeofenceEvents(cfg aws.Config, client *location.Client) {
	input := &location.ForecastGeofenceEventsInput{
		// CollectionName: *string, // Required
		// DeviceState: *types.ForecastGeofenceEventsDeviceState, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationDeviceState) > 0 {
		if err := assignInputField(input, "DeviceState", _locationDeviceState); err != nil {
			log.Errorf("invalid --device-state: %s", err.Error())
			return
		}
	}
	if len(_locationDistanceUnit) > 0 {
		if err := assignInputField(input, "DistanceUnit", _locationDistanceUnit); err != nil {
			log.Errorf("invalid --distance-unit: %s", err.Error())
			return
		}
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}
	if len(_locationSpeedUnit) > 0 {
		if err := assignInputField(input, "SpeedUnit", _locationSpeedUnit); err != nil {
			log.Errorf("invalid --speed-unit: %s", err.Error())
			return
		}
	}
	if len(_locationTimeHorizonMinutes) > 0 {
		if err := assignInputField(input, "TimeHorizonMinutes", _locationTimeHorizonMinutes); err != nil {
			log.Errorf("invalid --time-horizon-minutes: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ForecastGeofenceEvents(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ForecastGeofenceEventsOutput
	p := location.NewForecastGeofenceEventsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves a device's most recent position according to its sample time.
// Device positions are deleted after 30 days.
func location_GetDevicePosition(cfg aws.Config, client *location.Client) {
	input := &location.GetDevicePositionInput{
		// DeviceId: *string, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationDeviceId) > 0 {
		input.DeviceId = aws.String(_locationDeviceId)
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}

	if resp, err := client.GetDevicePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the device position history from a tracker resource within a
// specified range of time.
//
// Device positions are deleted after 30 days.
func location_GetDevicePositionHistory(cfg aws.Config, client *location.Client) {
	input := &location.GetDevicePositionHistoryInput{
		// DeviceId: *string, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationDeviceId) > 0 {
		input.DeviceId = aws.String(_locationDeviceId)
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationEndTimeExclusive) > 0 {
		if err := assignInputField(input, "EndTimeExclusive", _locationEndTimeExclusive); err != nil {
			log.Errorf("invalid --end-time-exclusive: %s", err.Error())
			return
		}
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}
	if len(_locationStartTimeInclusive) > 0 {
		if err := assignInputField(input, "StartTimeInclusive", _locationStartTimeInclusive); err != nil {
			log.Errorf("invalid --start-time-inclusive: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetDevicePositionHistory(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.GetDevicePositionHistoryOutput
	p := location.NewGetDevicePositionHistoryPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the geofence details from a geofence collection.
// The returned geometry will always match the geometry format used when the
// geofence was created.
func location_GetGeofence(cfg aws.Config, client *location.Client) {
	input := &location.GetGeofenceInput{
		// CollectionName: *string, // Required
		// GeofenceId: *string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationGeofenceId) > 0 {
		input.GeofenceId = aws.String(_locationGeofenceId)
	}

	if resp, err := client.GetGeofence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to [GetGlyphs]GetGlyphs unless you require Grab data.
//
// - GetMapGlyphs is part of a previous Amazon Location Service Maps API (version
// 1) which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The version 2 GetGlyphs operation gives a better user experience and is
// compatible with the remainder of the V2 Maps API.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Retrieves glyphs used to display labels on a map.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
// [GetGlyphs]: https://docs.aws.amazon.com/location/latest/APIReference/API_geomaps_GetGlyphs.html
func location_GetMapGlyphs(cfg aws.Config, client *location.Client) {
	input := &location.GetMapGlyphsInput{
		// FontStack: *string, // Required
		// FontUnicodeRange: *string, // Required
		// MapName: *string, // Required
	}

	if len(_locationFontStack) > 0 {
		input.FontStack = aws.String(_locationFontStack)
	}
	if len(_locationFontUnicodeRange) > 0 {
		input.FontUnicodeRange = aws.String(_locationFontUnicodeRange)
	}
	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}

	if resp, err := client.GetMapGlyphs(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to [GetSprites]GetSprites unless you require Grab data.
//
// - GetMapSprites is part of a previous Amazon Location Service Maps API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The version 2 GetSprites operation gives a better user experience and is
// compatible with the remainder of the V2 Maps API.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Retrieves the sprite sheet corresponding to a map resource. The sprite sheet is
// a PNG image paired with a JSON document describing the offsets of individual
// icons that will be displayed on a rendered map.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
// [GetSprites]: https://docs.aws.amazon.com/location/latest/APIReference/API_geomaps_GetSprites.html
func location_GetMapSprites(cfg aws.Config, client *location.Client) {
	input := &location.GetMapSpritesInput{
		// FileName: *string, // Required
		// MapName: *string, // Required
	}

	if len(_locationFileName) > 0 {
		input.FileName = aws.String(_locationFileName)
	}
	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}

	if resp, err := client.GetMapSprites(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to [GetStyleDescriptor]GetStyleDescriptor unless you require Grab data.
//
// - GetMapStyleDescriptor is part of a previous Amazon Location Service Maps API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The version 2 GetStyleDescriptor operation gives a better user experience
// and is compatible with the remainder of the V2 Maps API.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Retrieves the map style descriptor from a map resource.
//
// The style descriptor contains speciﬁcations on how features render on a map.
// For example, what data to display, what order to display the data in, and the
// style for the data. Style descriptors follow the Mapbox Style Specification.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
// [GetStyleDescriptor]: https://docs.aws.amazon.com/location/latest/APIReference/API_geomaps_GetStyleDescriptor.html
func location_GetMapStyleDescriptor(cfg aws.Config, client *location.Client) {
	input := &location.GetMapStyleDescriptorInput{
		// MapName: *string, // Required
	}

	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}

	if resp, err := client.GetMapStyleDescriptor(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to [GetTile]GetTile unless you require Grab data.
//
// - GetMapTile is part of a previous Amazon Location Service Maps API (version
// 1) which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The version 2 GetTile operation gives a better user experience and is
// compatible with the remainder of the V2 Maps API.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Retrieves a vector data tile from the map resource. Map tiles are used by
// clients to render a map. they're addressed using a grid arrangement with an X
// coordinate, Y coordinate, and Z (zoom) level.
//
// The origin (0, 0) is the top left of the map. Increasing the zoom level by 1
// doubles both the X and Y dimensions, so a tile containing data for the entire
// world at (0/0/0) will be split into 4 tiles at zoom 1 (1/0/0, 1/0/1, 1/1/0,
// 1/1/1).
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
// [GetTile]: https://docs.aws.amazon.com/location/latest/APIReference/API_geomaps_GetTile.html
func location_GetMapTile(cfg aws.Config, client *location.Client) {
	input := &location.GetMapTileInput{
		// MapName: *string, // Required
		// X: *string, // Required
		// Y: *string, // Required
		// Z: *string, // Required
	}

	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}
	if len(_locationX) > 0 {
		input.X = aws.String(_locationX)
	}
	if len(_locationY) > 0 {
		input.Y = aws.String(_locationY)
	}
	if len(_locationZ) > 0 {
		input.Z = aws.String(_locationZ)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}

	if resp, err := client.GetMapTile(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the V2 GetPlaceGetPlace operation unless you require Grab data.
//
// - This version of GetPlace is part of a previous Amazon Location Service
// Places API (version 1) which has been superseded by a more intuitive, powerful,
// and complete API (version 2).
//
// - Version 2 of the GetPlace operation interoperates with the rest of the
// Places V2 API, while this version does not.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Finds a place by its unique ID. A PlaceId is returned by other search
// operations.
//
// A PlaceId is valid only if all of the following are the same in the original
// search request and the call to GetPlace .
//
// - Customer Amazon Web Services account
//
// - Amazon Web Services Region
//
// - Data provider specified in the place index resource
//
// If your Place index resource is configured with Grab as your geolocation
// provider and Storage as Intended use, the GetPlace operation is unavailable. For
// more information, see [AWS service terms].
//
// [AWS service terms]: http://aws.amazon.com/service-terms
func location_GetPlace(cfg aws.Config, client *location.Client) {
	input := &location.GetPlaceInput{
		// IndexName: *string, // Required
		// PlaceId: *string, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}
	if len(_locationPlaceId) > 0 {
		input.PlaceId = aws.String(_locationPlaceId)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}
	if len(_locationLanguage) > 0 {
		input.Language = aws.String(_locationLanguage)
	}

	if resp, err := client.GetPlace(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// A batch request to retrieve all device positions.
func location_ListDevicePositions(cfg aws.Config, client *location.Client) {
	input := &location.ListDevicePositionsInput{
		// TrackerName: *string, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationFilterGeometry) > 0 {
		if err := assignInputField(input, "FilterGeometry", _locationFilterGeometry); err != nil {
			log.Errorf("invalid --filter-geometry: %s", err.Error())
			return
		}
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListDevicePositions(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListDevicePositionsOutput
	p := location.NewListDevicePositionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists geofence collections in your Amazon Web Services account.
func location_ListGeofenceCollections(cfg aws.Config, client *location.Client) {
	input := &location.ListGeofenceCollectionsInput{}

	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGeofenceCollections(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListGeofenceCollectionsOutput
	p := location.NewListGeofenceCollectionsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists geofences stored in a given geofence collection.
func location_ListGeofences(cfg aws.Config, client *location.Client) {
	input := &location.ListGeofencesInput{
		// CollectionName: *string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListGeofences(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListGeofencesOutput
	p := location.NewListGeofencesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists API key resources in your Amazon Web Services account.
// For more information, see [Use API keys to authenticate] in the Amazon Location Service Developer Guide.
//
// [Use API keys to authenticate]: https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html
func location_ListKeys(cfg aws.Config, client *location.Client) {
	input := &location.ListKeysInput{}

	if len(_locationFilter) > 0 {
		if err := assignInputField(input, "Filter", _locationFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListKeysOutput
	p := location.NewListKeysPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to the Maps API V2 unless you require Grab data.
//
// - ListMaps is part of a previous Amazon Location Service Maps API (version 1)
// which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The Maps API version 2 has a simplified interface that can be used without
// creating or managing map resources.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Lists map resources in your Amazon Web Services account.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
func location_ListMaps(cfg aws.Config, client *location.Client) {
	input := &location.ListMapsInput{}

	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMaps(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListMapsOutput
	p := location.NewListMapsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Places API V2 unless you require Grab data.
//
// - ListPlaceIndexes is part of a previous Amazon Location Service Places API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Places API version 2 has a simplified interface that can be used
// without creating or managing place index resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Lists place index resources in your Amazon Web Services account.
func location_ListPlaceIndexes(cfg aws.Config, client *location.Client) {
	input := &location.ListPlaceIndexesInput{}

	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListPlaceIndexes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListPlaceIndexesOutput
	p := location.NewListPlaceIndexesPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Routes API V2 unless you require Grab data.
//
// - ListRouteCalculators is part of a previous Amazon Location Service Routes
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Routes API version 2 has a simplified interface that can be used
// without creating or managing route calculator resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Routes V2 API Referenceor the Developer Guide.
//
// Lists route calculator resources in your Amazon Web Services account.
func location_ListRouteCalculators(cfg aws.Config, client *location.Client) {
	input := &location.ListRouteCalculatorsInput{}

	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRouteCalculators(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListRouteCalculatorsOutput
	p := location.NewListRouteCalculatorsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns a list of tags that are applied to the specified Amazon Location
// resource.
func location_ListTagsForResource(cfg aws.Config, client *location.Client) {
	input := &location.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_locationResourceArn) > 0 {
		input.ResourceArn = aws.String(_locationResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Lists geofence collections currently associated to the given tracker resource.
func location_ListTrackerConsumers(cfg aws.Config, client *location.Client) {
	input := &location.ListTrackerConsumersInput{
		// TrackerName: *string, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrackerConsumers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListTrackerConsumersOutput
	p := location.NewListTrackerConsumersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists tracker resources in your Amazon Web Services account.
func location_ListTrackers(cfg aws.Config, client *location.Client) {
	input := &location.ListTrackersInput{}

	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_locationNextToken) > 0 {
		input.NextToken = aws.String(_locationNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTrackers(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*location.ListTrackersOutput
	p := location.NewListTrackersPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Stores a geofence geometry in a given geofence collection, or updates the
// geometry of an existing geofence if a geofence ID is included in the request.
func location_PutGeofence(cfg aws.Config, client *location.Client) {
	input := &location.PutGeofenceInput{
		// CollectionName: *string, // Required
		// GeofenceId: *string, // Required
		// Geometry: *types.GeofenceGeometry, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationGeofenceId) > 0 {
		input.GeofenceId = aws.String(_locationGeofenceId)
	}
	if len(_locationGeometry) > 0 {
		if err := assignInputField(input, "Geometry", _locationGeometry); err != nil {
			log.Errorf("invalid --geometry: %s", err.Error())
			return
		}
	}
	if len(_locationGeofenceProperties) > 0 {
		if err := assignInputField(input, "GeofenceProperties", _locationGeofenceProperties); err != nil {
			log.Errorf("invalid --geofence-properties: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutGeofence(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to ReverseGeocodeReverseGeocode or SearchNearbySearchNearby unless you require Grab
// data.
//
// - SearchPlaceIndexForPosition is part of a previous Amazon Location Service
// Places API (version 1) which has been superseded by a more intuitive, powerful,
// and complete API (version 2).
//
// - The version 2 ReverseGeocode operation gives better results in the address
// reverse-geocoding use case, while the version 2 SearchNearby operation gives
// better results when searching for businesses and points of interest near a
// specific location.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// Reverse geocodes a given coordinate and returns a legible address. Allows you
// to search for Places or points of interest near a given position.
func location_SearchPlaceIndexForPosition(cfg aws.Config, client *location.Client) {
	input := &location.SearchPlaceIndexForPositionInput{
		// IndexName: *string, // Required
		// Position: []float64, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}
	if len(_locationPosition) > 0 {
		if err := assignInputField(input, "Position", _locationPosition); err != nil {
			log.Errorf("invalid --position: %s", err.Error())
			return
		}
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}
	if len(_locationLanguage) > 0 {
		input.Language = aws.String(_locationLanguage)
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchPlaceIndexForPosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to SuggestSuggest or AutocompleteAutocomplete unless you require Grab data.
//
// - SearchPlaceIndexForSuggestions is part of a previous Amazon Location Service
// Places API (version 1) which has been superseded by a more intuitive, powerful,
// and complete API (version 2).
//
// - The version 2 Suggest operation gives better results for typeahead place
// search suggestions with fuzzy matching, while the version 2 Autocomplete
// operation gives better results for address completion based on partial input.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// Generates suggestions for addresses and points of interest based on partial or
// misspelled free-form text. This operation is also known as autocomplete,
// autosuggest, or fuzzy matching.
//
// Optional parameters let you narrow your search results by bounding box or
// country, or bias your search toward a specific position on the globe.
//
// You can search for suggested place names near a specified position by using
// BiasPosition , or filter results within a bounding box by using FilterBBox .
// These parameters are mutually exclusive; using both BiasPosition and FilterBBox
// in the same command returns an error.
func location_SearchPlaceIndexForSuggestions(cfg aws.Config, client *location.Client) {
	input := &location.SearchPlaceIndexForSuggestionsInput{
		// IndexName: *string, // Required
		// Text: *string, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}
	if len(_locationText) > 0 {
		input.Text = aws.String(_locationText)
	}
	if len(_locationBiasPosition) > 0 {
		if err := assignInputField(input, "BiasPosition", _locationBiasPosition); err != nil {
			log.Errorf("invalid --bias-position: %s", err.Error())
			return
		}
	}
	if len(_locationFilterBBox) > 0 {
		if err := assignInputField(input, "FilterBBox", _locationFilterBBox); err != nil {
			log.Errorf("invalid --filter-bbox: %s", err.Error())
			return
		}
	}
	if len(_locationFilterCategories) > 0 {
		input.FilterCategories = append([]string(nil), _locationFilterCategories...)
	}
	if len(_locationFilterCountries) > 0 {
		input.FilterCountries = append([]string(nil), _locationFilterCountries...)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}
	if len(_locationLanguage) > 0 {
		input.Language = aws.String(_locationLanguage)
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchPlaceIndexForSuggestions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to GeocodeGeocode or SearchTextSearchText unless you require Grab data.
//
// - SearchPlaceIndexForText is part of a previous Amazon Location Service Places
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The version 2 Geocode operation gives better results in the address
// geocoding use case, while the version 2 SearchText operation gives better
// results when searching for businesses and points of interest.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// Geocodes free-form text, such as an address, name, city, or region to allow you
// to search for Places or points of interest.
//
// Optional parameters let you narrow your search results by bounding box or
// country, or bias your search toward a specific position on the globe.
//
// You can search for places near a given position using BiasPosition , or filter
// results within a bounding box using FilterBBox . Providing both parameters
// simultaneously returns an error.
//
// Search results are returned in order of highest to lowest relevance.
func location_SearchPlaceIndexForText(cfg aws.Config, client *location.Client) {
	input := &location.SearchPlaceIndexForTextInput{
		// IndexName: *string, // Required
		// Text: *string, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}
	if len(_locationText) > 0 {
		input.Text = aws.String(_locationText)
	}
	if len(_locationBiasPosition) > 0 {
		if err := assignInputField(input, "BiasPosition", _locationBiasPosition); err != nil {
			log.Errorf("invalid --bias-position: %s", err.Error())
			return
		}
	}
	if len(_locationFilterBBox) > 0 {
		if err := assignInputField(input, "FilterBBox", _locationFilterBBox); err != nil {
			log.Errorf("invalid --filter-bbox: %s", err.Error())
			return
		}
	}
	if len(_locationFilterCategories) > 0 {
		input.FilterCategories = append([]string(nil), _locationFilterCategories...)
	}
	if len(_locationFilterCountries) > 0 {
		input.FilterCountries = append([]string(nil), _locationFilterCountries...)
	}
	if len(_locationKey) > 0 {
		input.Key = aws.String(_locationKey)
	}
	if len(_locationLanguage) > 0 {
		input.Language = aws.String(_locationLanguage)
	}
	if len(_locationMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _locationMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}

	if resp, err := client.SearchPlaceIndexForText(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Assigns one or more tags (key-value pairs) to the specified Amazon Location
// Service resource.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions, by granting a user permission to access or change
// only resources with certain tag values.
//
// You can use the TagResource operation with an Amazon Location Service resource
// that already has tags. If you specify a new tag key for the resource, this tag
// is appended to the tags already associated with the resource. If you specify a
// tag key that's already associated with the resource, the new tag value that you
// specify replaces the previous value for that tag.
//
// You can associate up to 50 tags with a resource.
func location_TagResource(cfg aws.Config, client *location.Client) {
	input := &location.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_locationResourceArn) > 0 {
		input.ResourceArn = aws.String(_locationResourceArn)
	}
	if len(_locationTags) > 0 {
		if err := assignInputField(input, "Tags", _locationTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes one or more tags from the specified Amazon Location resource.
func location_UntagResource(cfg aws.Config, client *location.Client) {
	input := &location.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_locationResourceArn) > 0 {
		input.ResourceArn = aws.String(_locationResourceArn)
	}
	if len(_locationTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _locationTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified properties of a given geofence collection.
func location_UpdateGeofenceCollection(cfg aws.Config, client *location.Client) {
	input := &location.UpdateGeofenceCollectionInput{
		// CollectionName: *string, // Required
	}

	if len(_locationCollectionName) > 0 {
		input.CollectionName = aws.String(_locationCollectionName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationPricingPlanDataSource) > 0 {
		input.PricingPlanDataSource = aws.String(_locationPricingPlanDataSource)
	}

	if resp, err := client.UpdateGeofenceCollection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified properties of a given API key resource.
func location_UpdateKey(cfg aws.Config, client *location.Client) {
	input := &location.UpdateKeyInput{
		// KeyName: *string, // Required
	}

	if len(_locationKeyName) > 0 {
		input.KeyName = aws.String(_locationKeyName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationExpireTime) > 0 {
		if err := assignInputField(input, "ExpireTime", _locationExpireTime); err != nil {
			log.Errorf("invalid --expire-time: %s", err.Error())
			return
		}
	}
	if len(_locationForceUpdate) > 0 {
		if err := assignInputField(input, "ForceUpdate", _locationForceUpdate); err != nil {
			log.Errorf("invalid --force-update: %s", err.Error())
			return
		}
	}
	if len(_locationNoExpiry) > 0 {
		if err := assignInputField(input, "NoExpiry", _locationNoExpiry); err != nil {
			log.Errorf("invalid --no-expiry: %s", err.Error())
			return
		}
	}
	if len(_locationRestrictions) > 0 {
		if err := assignInputField(input, "Restrictions", _locationRestrictions); err != nil {
			log.Errorf("invalid --restrictions: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateKey(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend upgrading to the Maps API V2 unless you require Grab data.
//
// - UpdateMap is part of a previous Amazon Location Service Maps API (version 1)
// which has been superseded by a more intuitive, powerful, and complete API
// (version 2).
//
// - The Maps API version 2 has a simplified interface that can be used without
// creating or managing map resources.
//
// - If you are using an AWS SDK or the AWS CLI, note that the Maps API version
// 2 is found under geo-maps or geo_maps , not under location .
//
// - Since Grab is not yet fully supported in Maps API version 2, we recommend
// you continue using API version 1 when using Grab .
//
// - Start your version 2 API journey with the [Maps V2 API Reference]or the [Developer Guide].
//
// Updates the specified properties of a given map resource.
//
// [Maps V2 API Reference]: https://docs.aws.amazon.com/location/latest/APIReference/API_Operations_Amazon_Location_Service_Maps_V2.html
// [Developer Guide]: https://docs.aws.amazon.com/location/latest/developerguide/maps.html
func location_UpdateMap(cfg aws.Config, client *location.Client) {
	input := &location.UpdateMapInput{
		// MapName: *string, // Required
	}

	if len(_locationMapName) > 0 {
		input.MapName = aws.String(_locationMapName)
	}
	if len(_locationConfigurationUpdate) > 0 {
		if err := assignInputField(input, "ConfigurationUpdate", _locationConfigurationUpdate); err != nil {
			log.Errorf("invalid --configuration-update: %s", err.Error())
			return
		}
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateMap(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Places API V2 unless you require Grab data.
//
// - UpdatePlaceIndex is part of a previous Amazon Location Service Places API
// (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Places API version 2 has a simplified interface that can be used
// without creating or managing place index resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Places API version 2 is found under geo-places or geo_places ,
// not under location .
//
// - Since Grab is not yet fully supported in Places API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Places V2 API Referenceor the Developer Guide.
//
// Updates the specified properties of a given place index resource.
func location_UpdatePlaceIndex(cfg aws.Config, client *location.Client) {
	input := &location.UpdatePlaceIndexInput{
		// IndexName: *string, // Required
	}

	if len(_locationIndexName) > 0 {
		input.IndexName = aws.String(_locationIndexName)
	}
	if len(_locationDataSourceConfiguration) > 0 {
		if err := assignInputField(input, "DataSourceConfiguration", _locationDataSourceConfiguration); err != nil {
			log.Errorf("invalid --data-source-configuration: %s", err.Error())
			return
		}
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdatePlaceIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is no longer current and may be deprecated in the future. We
// recommend you upgrade to the Routes API V2 unless you require Grab data.
//
// - UpdateRouteCalculator is part of a previous Amazon Location Service Routes
// API (version 1) which has been superseded by a more intuitive, powerful, and
// complete API (version 2).
//
// - The Routes API version 2 has a simplified interface that can be used
// without creating or managing route calculator resources.
//
// - If you are using an Amazon Web Services SDK or the Amazon Web Services CLI,
// note that the Routes API version 2 is found under geo-routes or geo_routes ,
// not under location .
//
// - Since Grab is not yet fully supported in Routes API version 2, we recommend
// you continue using API version 1 when using Grab.
//
// - Start your version 2 API journey with the Routes V2 API Referenceor the Developer Guide.
//
// Updates the specified properties for a given route calculator resource.
func location_UpdateRouteCalculator(cfg aws.Config, client *location.Client) {
	input := &location.UpdateRouteCalculatorInput{
		// CalculatorName: *string, // Required
	}

	if len(_locationCalculatorName) > 0 {
		input.CalculatorName = aws.String(_locationCalculatorName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateRouteCalculator(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the specified properties of a given tracker resource.
func location_UpdateTracker(cfg aws.Config, client *location.Client) {
	input := &location.UpdateTrackerInput{
		// TrackerName: *string, // Required
	}

	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationDescription) > 0 {
		input.Description = aws.String(_locationDescription)
	}
	if len(_locationEventBridgeEnabled) > 0 {
		if err := assignInputField(input, "EventBridgeEnabled", _locationEventBridgeEnabled); err != nil {
			log.Errorf("invalid --event-bridge-enabled: %s", err.Error())
			return
		}
	}
	if len(_locationKmsKeyEnableGeospatialQueries) > 0 {
		if err := assignInputField(input, "KmsKeyEnableGeospatialQueries", _locationKmsKeyEnableGeospatialQueries); err != nil {
			log.Errorf("invalid --kms-key-enable-geospatial-queries: %s", err.Error())
			return
		}
	}
	if len(_locationPositionFiltering) > 0 {
		if err := assignInputField(input, "PositionFiltering", _locationPositionFiltering); err != nil {
			log.Errorf("invalid --position-filtering: %s", err.Error())
			return
		}
	}
	if len(_locationPricingPlan) > 0 {
		if err := assignInputField(input, "PricingPlan", _locationPricingPlan); err != nil {
			log.Errorf("invalid --pricing-plan: %s", err.Error())
			return
		}
	}
	if len(_locationPricingPlanDataSource) > 0 {
		input.PricingPlanDataSource = aws.String(_locationPricingPlanDataSource)
	}

	if resp, err := client.UpdateTracker(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Verifies the integrity of the device's position by determining if it was
// reported behind a proxy, and by comparing it to an inferred position estimated
// based on the device's state.
//
// The Location Integrity SDK provides enhanced features related to device
// verification, and it is available for use by request. To get access to the SDK,
// contact [Sales Support].
//
// [Sales Support]: https://aws.amazon.com/contact-us/sales-support/?pg=locationprice&cta=herobtn
func location_VerifyDevicePosition(cfg aws.Config, client *location.Client) {
	input := &location.VerifyDevicePositionInput{
		// DeviceState: *types.DeviceState, // Required
		// TrackerName: *string, // Required
	}

	if len(_locationDeviceState) > 0 {
		if err := assignInputField(input, "DeviceState", _locationDeviceState); err != nil {
			log.Errorf("invalid --device-state: %s", err.Error())
			return
		}
	}
	if len(_locationTrackerName) > 0 {
		input.TrackerName = aws.String(_locationTrackerName)
	}
	if len(_locationDistanceUnit) > 0 {
		if err := assignInputField(input, "DistanceUnit", _locationDistanceUnit); err != nil {
			log.Errorf("invalid --distance-unit: %s", err.Error())
			return
		}
	}

	if resp, err := client.VerifyDevicePosition(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_locationCmd)
	_locationCmd.Flags().SortFlags = false

	_locationCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_locationCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_locationCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_locationCmd.Flags().StringVarP(&_locationArrivalTime, "arrival-time", "", "", "Arrival Time")
	_locationCmd.Flags().StringVarP(&_locationBiasPosition, "bias-position", "", "", "Bias Position")
	_locationCmd.Flags().StringVarP(&_locationCalculatorName, "calculator-name", "", "", "Calculator Name")
	_locationCmd.Flags().StringVarP(&_locationCarModeOptions, "car-mode-options", "", "", "Car Mode Options")
	_locationCmd.Flags().StringVarP(&_locationCollectionName, "collection-name", "", "", "Collection Name")
	_locationCmd.Flags().StringVarP(&_locationConfiguration, "configuration", "", "", "Configuration")
	_locationCmd.Flags().StringVarP(&_locationConfigurationUpdate, "configuration-update", "", "", "Configuration Update")
	_locationCmd.Flags().StringVarP(&_locationConsumerArn, "consumer-arn", "", "", "Consumer ARN")
	_locationCmd.Flags().StringVarP(&_locationDataSource, "data-source", "", "", "Data Source")
	_locationCmd.Flags().StringVarP(&_locationDataSourceConfiguration, "data-source-configuration", "", "", "Data Source Configuration")
	_locationCmd.Flags().StringVarP(&_locationDepartNow, "depart-now", "", "", "Depart Now")
	_locationCmd.Flags().StringVarP(&_locationDeparturePosition, "departure-position", "", "", "Departure Position")
	_locationCmd.Flags().StringVarP(&_locationDeparturePositions, "departure-positions", "", "", "Departure Positions")
	_locationCmd.Flags().StringVarP(&_locationDepartureTime, "departure-time", "", "", "Departure Time")
	_locationCmd.Flags().StringVarP(&_locationDescription, "description", "", "", "Description")
	_locationCmd.Flags().StringVarP(&_locationDestinationPosition, "destination-position", "", "", "Destination Position")
	_locationCmd.Flags().StringVarP(&_locationDestinationPositions, "destination-positions", "", "", "Destination Positions")
	_locationCmd.Flags().StringVarP(&_locationDeviceId, "device-id", "", "", "Device ID")
	_locationCmd.Flags().StringSliceVarP(&_locationDeviceIds, "device-ids", "", nil, "Device Ids")
	_locationCmd.Flags().StringVarP(&_locationDevicePositionUpdates, "device-position-updates", "", "", "Device Position Updates")
	_locationCmd.Flags().StringVarP(&_locationDeviceState, "device-state", "", "", "Device State")
	_locationCmd.Flags().StringVarP(&_locationDistanceUnit, "distance-unit", "", "", "Distance Unit")
	_locationCmd.Flags().StringVarP(&_locationEndTimeExclusive, "end-time-exclusive", "", "", "End Time Exclusive")
	_locationCmd.Flags().StringVarP(&_locationEntries, "entries", "", "", "Entries")
	_locationCmd.Flags().StringVarP(&_locationEventBridgeEnabled, "event-bridge-enabled", "", "", "Event Bridge Enabled")
	_locationCmd.Flags().StringVarP(&_locationExpireTime, "expire-time", "", "", "Expire Time")
	_locationCmd.Flags().StringVarP(&_locationFileName, "file-name", "", "", "File Name")
	_locationCmd.Flags().StringVarP(&_locationFilter, "filter", "", "", "Filter")
	_locationCmd.Flags().StringVarP(&_locationFilterBBox, "filter-bbox", "", "", "Filter Bbox")
	_locationCmd.Flags().StringSliceVarP(&_locationFilterCategories, "filter-categories", "", nil, "Filter Categories")
	_locationCmd.Flags().StringSliceVarP(&_locationFilterCountries, "filter-countries", "", nil, "Filter Countries")
	_locationCmd.Flags().StringVarP(&_locationFilterGeometry, "filter-geometry", "", "", "Filter Geometry")
	_locationCmd.Flags().StringVarP(&_locationFontStack, "font-stack", "", "", "Font Stack")
	_locationCmd.Flags().StringVarP(&_locationFontUnicodeRange, "font-unicode-range", "", "", "Font Unicode Range")
	_locationCmd.Flags().StringVarP(&_locationForceDelete, "force-delete", "", "", "Force Delete")
	_locationCmd.Flags().StringVarP(&_locationForceUpdate, "force-update", "", "", "Force Update")
	_locationCmd.Flags().StringVarP(&_locationGeofenceId, "geofence-id", "", "", "Geofence ID")
	_locationCmd.Flags().StringSliceVarP(&_locationGeofenceIds, "geofence-ids", "", nil, "Geofence Ids")
	_locationCmd.Flags().StringVarP(&_locationGeofenceProperties, "geofence-properties", "", "", "Geofence Properties")
	_locationCmd.Flags().StringVarP(&_locationGeometry, "geometry", "", "", "Geometry")
	_locationCmd.Flags().StringVarP(&_locationIncludeLegGeometry, "include-leg-geometry", "", "", "Include Leg Geometry")
	_locationCmd.Flags().StringVarP(&_locationIndexName, "index-name", "", "", "Index Name")
	_locationCmd.Flags().StringVarP(&_locationKey, "key", "", "", "Key")
	_locationCmd.Flags().StringVarP(&_locationKeyName, "key-name", "", "", "Key Name")
	_locationCmd.Flags().StringVarP(&_locationKmsKeyEnableGeospatialQueries, "kms-key-enable-geospatial-queries", "", "", "KMS Key Enable Geospatial Queries")
	_locationCmd.Flags().StringVarP(&_locationKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_locationCmd.Flags().StringVarP(&_locationLanguage, "language", "", "", "Language")
	_locationCmd.Flags().StringVarP(&_locationMapName, "map-name", "", "", "Map Name")
	_locationCmd.Flags().StringVarP(&_locationMaxResults, "max-results", "", "", "Max Results")
	_locationCmd.Flags().StringVarP(&_locationNextToken, "next-token", "", "", "Next Token")
	_locationCmd.Flags().StringVarP(&_locationNoExpiry, "no-expiry", "", "", "No Expiry")
	_locationCmd.Flags().StringVarP(&_locationOptimizeFor, "optimize-for", "", "", "Optimize For")
	_locationCmd.Flags().StringVarP(&_locationPlaceId, "place-id", "", "", "Place ID")
	_locationCmd.Flags().StringVarP(&_locationPosition, "position", "", "", "Position")
	_locationCmd.Flags().StringVarP(&_locationPositionFiltering, "position-filtering", "", "", "Position Filtering")
	_locationCmd.Flags().StringVarP(&_locationPricingPlan, "pricing-plan", "", "", "Pricing Plan")
	_locationCmd.Flags().StringVarP(&_locationPricingPlanDataSource, "pricing-plan-data-source", "", "", "Pricing Plan Data Source")
	_locationCmd.Flags().StringVarP(&_locationResourceArn, "resource-arn", "", "", "Resource ARN")
	_locationCmd.Flags().StringVarP(&_locationRestrictions, "restrictions", "", "", "Restrictions")
	_locationCmd.Flags().StringVarP(&_locationSpeedUnit, "speed-unit", "", "", "Speed Unit")
	_locationCmd.Flags().StringVarP(&_locationStartTimeInclusive, "start-time-inclusive", "", "", "Start Time Inclusive")
	_locationCmd.Flags().StringSliceVarP(&_locationTagKeys, "tag-keys", "", nil, "Tag Keys")
	_locationCmd.Flags().StringVarP(&_locationTags, "tags", "", "", "Tags")
	_locationCmd.Flags().StringVarP(&_locationText, "text", "", "", "Text")
	_locationCmd.Flags().StringVarP(&_locationTimeHorizonMinutes, "time-horizon-minutes", "", "", "Time Horizon Minutes")
	_locationCmd.Flags().StringVarP(&_locationTrackerName, "tracker-name", "", "", "Tracker Name")
	_locationCmd.Flags().StringVarP(&_locationTravelMode, "travel-mode", "", "", "Travel Mode")
	_locationCmd.Flags().StringVarP(&_locationTruckModeOptions, "truck-mode-options", "", "", "Truck Mode Options")
	_locationCmd.Flags().StringVarP(&_locationUpdates, "updates", "", "", "Updates")
	_locationCmd.Flags().StringVarP(&_locationWaypointPositions, "waypoint-positions", "", "", "Waypoint Positions")
	_locationCmd.Flags().StringVarP(&_locationX, "x", "", "", "X")
	_locationCmd.Flags().StringVarP(&_locationY, "y", "", "", "Y")
	_locationCmd.Flags().StringVarP(&_locationZ, "z", "", "", "Z")

	_locationCmd.Flags().BoolVarP(&_locationAssociateTrackerConsumer, "associate-tracker-consumer", "", false, "Associate Tracker Consumer")
	_locationCmd.Flags().BoolVarP(&_locationBatchDeleteDevicePositionHistory, "batch-delete-device-position-history", "", false, "Batch Delete Device Position History")
	_locationCmd.Flags().BoolVarP(&_locationBatchDeleteGeofence, "batch-delete-geofence", "", false, "Batch Delete Geofence")
	_locationCmd.Flags().BoolVarP(&_locationBatchEvaluateGeofences, "batch-evaluate-geofences", "", false, "Batch Evaluate Geofences")
	_locationCmd.Flags().BoolVarP(&_locationBatchGetDevicePosition, "batch-get-device-position", "", false, "Batch Get Device Position")
	_locationCmd.Flags().BoolVarP(&_locationBatchPutGeofence, "batch-put-geofence", "", false, "Batch Put Geofence")
	_locationCmd.Flags().BoolVarP(&_locationBatchUpdateDevicePosition, "batch-update-device-position", "", false, "Batch Update Device Position")
	_locationCmd.Flags().BoolVarP(&_locationCalculateRoute, "calculate-route", "", false, "Calculate Route")
	_locationCmd.Flags().BoolVarP(&_locationCalculateRouteMatrix, "calculate-route-matrix", "", false, "Calculate Route Matrix")
	_locationCmd.Flags().BoolVarP(&_locationCreateGeofenceCollection, "create-geofence-collection", "", false, "Create Geofence Collection")
	_locationCmd.Flags().BoolVarP(&_locationCreateKey, "create-key", "", false, "Create Key")
	_locationCmd.Flags().BoolVarP(&_locationCreateMap, "create-map", "", false, "Create Map")
	_locationCmd.Flags().BoolVarP(&_locationCreatePlaceIndex, "create-place-index", "", false, "Create Place Index")
	_locationCmd.Flags().BoolVarP(&_locationCreateRouteCalculator, "create-route-calculator", "", false, "Create Route Calculator")
	_locationCmd.Flags().BoolVarP(&_locationCreateTracker, "create-tracker", "", false, "Create Tracker")
	_locationCmd.Flags().BoolVarP(&_locationDeleteGeofenceCollection, "delete-geofence-collection", "", false, "Delete Geofence Collection")
	_locationCmd.Flags().BoolVarP(&_locationDeleteKey, "delete-key", "", false, "Delete Key")
	_locationCmd.Flags().BoolVarP(&_locationDeleteMap, "delete-map", "", false, "Delete Map")
	_locationCmd.Flags().BoolVarP(&_locationDeletePlaceIndex, "delete-place-index", "", false, "Delete Place Index")
	_locationCmd.Flags().BoolVarP(&_locationDeleteRouteCalculator, "delete-route-calculator", "", false, "Delete Route Calculator")
	_locationCmd.Flags().BoolVarP(&_locationDeleteTracker, "delete-tracker", "", false, "Delete Tracker")
	_locationCmd.Flags().BoolVarP(&_locationDescribeGeofenceCollection, "describe-geofence-collection", "", false, "Describe Geofence Collection")
	_locationCmd.Flags().BoolVarP(&_locationDescribeKey, "describe-key", "", false, "Describe Key")
	_locationCmd.Flags().BoolVarP(&_locationDescribeMap, "describe-map", "", false, "Describe Map")
	_locationCmd.Flags().BoolVarP(&_locationDescribePlaceIndex, "describe-place-index", "", false, "Describe Place Index")
	_locationCmd.Flags().BoolVarP(&_locationDescribeRouteCalculator, "describe-route-calculator", "", false, "Describe Route Calculator")
	_locationCmd.Flags().BoolVarP(&_locationDescribeTracker, "describe-tracker", "", false, "Describe Tracker")
	_locationCmd.Flags().BoolVarP(&_locationDisassociateTrackerConsumer, "disassociate-tracker-consumer", "", false, "Disassociate Tracker Consumer")
	_locationCmd.Flags().BoolVarP(&_locationForecastGeofenceEvents, "forecast-geofence-events", "", false, "Forecast Geofence Events")
	_locationCmd.Flags().BoolVarP(&_locationGetDevicePosition, "get-device-position", "", false, "Get Device Position")
	_locationCmd.Flags().BoolVarP(&_locationGetDevicePositionHistory, "get-device-position-history", "", false, "Get Device Position History")
	_locationCmd.Flags().BoolVarP(&_locationGetGeofence, "get-geofence", "", false, "Get Geofence")
	_locationCmd.Flags().BoolVarP(&_locationGetMapGlyphs, "get-map-glyphs", "", false, "Get Map Glyphs")
	_locationCmd.Flags().BoolVarP(&_locationGetMapSprites, "get-map-sprites", "", false, "Get Map Sprites")
	_locationCmd.Flags().BoolVarP(&_locationGetMapStyleDescriptor, "get-map-style-descriptor", "", false, "Get Map Style Descriptor")
	_locationCmd.Flags().BoolVarP(&_locationGetMapTile, "get-map-tile", "", false, "Get Map Tile")
	_locationCmd.Flags().BoolVarP(&_locationGetPlace, "get-place", "", false, "Get Place")
	_locationCmd.Flags().BoolVarP(&_locationListDevicePositions, "list-device-positions", "", false, "List Device Positions")
	_locationCmd.Flags().BoolVarP(&_locationListGeofenceCollections, "list-geofence-collections", "", false, "List Geofence Collections")
	_locationCmd.Flags().BoolVarP(&_locationListGeofences, "list-geofences", "", false, "List Geofences")
	_locationCmd.Flags().BoolVarP(&_locationListKeys, "list-keys", "", false, "List Keys")
	_locationCmd.Flags().BoolVarP(&_locationListMaps, "list-maps", "", false, "List Maps")
	_locationCmd.Flags().BoolVarP(&_locationListPlaceIndexes, "list-place-indexes", "", false, "List Place Indexes")
	_locationCmd.Flags().BoolVarP(&_locationListRouteCalculators, "list-route-calculators", "", false, "List Route Calculators")
	_locationCmd.Flags().BoolVarP(&_locationListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_locationCmd.Flags().BoolVarP(&_locationListTrackerConsumers, "list-tracker-consumers", "", false, "List Tracker Consumers")
	_locationCmd.Flags().BoolVarP(&_locationListTrackers, "list-trackers", "", false, "List Trackers")
	_locationCmd.Flags().BoolVarP(&_locationPutGeofence, "put-geofence", "", false, "Put Geofence")
	_locationCmd.Flags().BoolVarP(&_locationSearchPlaceIndexForPosition, "search-place-index-for-position", "", false, "Search Place Index For Position")
	_locationCmd.Flags().BoolVarP(&_locationSearchPlaceIndexForSuggestions, "search-place-index-for-suggestions", "", false, "Search Place Index For Suggestions")
	_locationCmd.Flags().BoolVarP(&_locationSearchPlaceIndexForText, "search-place-index-for-text", "", false, "Search Place Index For Text")
	_locationCmd.Flags().BoolVarP(&_locationTagResource, "tag-resource", "", false, "Tag Resource")
	_locationCmd.Flags().BoolVarP(&_locationUntagResource, "untag-resource", "", false, "Untag Resource")
	_locationCmd.Flags().BoolVarP(&_locationUpdateGeofenceCollection, "update-geofence-collection", "", false, "Update Geofence Collection")
	_locationCmd.Flags().BoolVarP(&_locationUpdateKey, "update-key", "", false, "Update Key")
	_locationCmd.Flags().BoolVarP(&_locationUpdateMap, "update-map", "", false, "Update Map")
	_locationCmd.Flags().BoolVarP(&_locationUpdatePlaceIndex, "update-place-index", "", false, "Update Place Index")
	_locationCmd.Flags().BoolVarP(&_locationUpdateRouteCalculator, "update-route-calculator", "", false, "Update Route Calculator")
	_locationCmd.Flags().BoolVarP(&_locationUpdateTracker, "update-tracker", "", false, "Update Tracker")
	_locationCmd.Flags().BoolVarP(&_locationVerifyDevicePosition, "verify-device-position", "", false, "Verify Device Position")

}
