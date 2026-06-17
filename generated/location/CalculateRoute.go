package location

// CalculateRoute is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
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
// If you specify walking for the travel mode and your data provider is Esri, the
//
// start and destination must be within 40km.
//
// [Specifying a departure time]: https://docs.aws.amazon.com/location/previous/developerguide/departure-time.html
// [Specifying a travel mode]: https://docs.aws.amazon.com/location/previous/developerguide/travel-mode.html
// [Calculates a route]: https://docs.aws.amazon.com/location/previous/developerguide/calculate-route.html
// [create a route calculator resource]: https://docs.aws.amazon.com/location-routes/latest/APIReference/API_CreateRouteCalculator.html
