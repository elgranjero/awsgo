package cloudfront

// UpdateRealtimeLogConfig is generated as a reference stub.
// Executable command wiring lives under cmd/cloudfront.go.
//
// Updates a real-time log configuration.
//
// When you update a real-time log configuration, all the parameters are updated
// with the values provided in the request. You cannot update some parameters
// independent of others. To update a real-time log configuration:
//
// - Call GetRealtimeLogConfig to get the current real-time log configuration.
//
// - Locally modify the parameters in the real-time log configuration that you
// want to update.
//
// - Call this API ( UpdateRealtimeLogConfig ) by providing the entire real-time
// log configuration, including the parameters that you modified and those that you
// didn't.
//
// You cannot update a real-time log configuration's Name or ARN .
