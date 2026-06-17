package appstream

// UpdateFleet is generated as a reference stub.
// Executable command wiring lives under cmd/appstream.go.
//
// Updates the specified fleet.
//
// If the fleet is in the STOPPED state, you can update any attribute except the
// fleet name.
//
// If the fleet is in the RUNNING state, you can update the following based on the
// fleet type:
//
// - Always-On and On-Demand fleet types
//
// You can update the DisplayName , ComputeCapacity , ImageARN , ImageName ,
//
// IdleDisconnectTimeoutInSeconds , and DisconnectTimeoutInSeconds attributes.
//
// - Elastic fleet type
//
// You can update the DisplayName , IdleDisconnectTimeoutInSeconds ,
//
// DisconnectTimeoutInSeconds , MaxConcurrentSessions , SessionScriptS3Location
// and UsbDeviceFilterStrings attributes.
//
// If the fleet is in the STARTING or STOPPED state, you can't update it.
