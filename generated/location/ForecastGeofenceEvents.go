package location

// ForecastGeofenceEvents is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
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
