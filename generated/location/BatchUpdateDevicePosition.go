package location

// BatchUpdateDevicePosition is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
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
