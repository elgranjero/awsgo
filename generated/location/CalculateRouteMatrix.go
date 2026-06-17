package location

// CalculateRouteMatrix is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
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
