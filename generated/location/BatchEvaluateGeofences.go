package location

// BatchEvaluateGeofences is generated as a reference stub.
// Executable command wiring lives under cmd/location.go.
//
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
