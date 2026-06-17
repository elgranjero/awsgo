package mediapackagev2

// ResetOriginEndpointState is generated as a reference stub.
// Executable command wiring lives under cmd/mediapackagev2.go.
//
// Resetting the origin endpoint can help to resolve unexpected behavior and other
// content packaging issues. It also helps to preserve special events when you
// don't want the previous content to be available for viewing. A reset clears out
// all previous content from the origin endpoint.
//
// MediaPackage might return old content from this endpoint in the first 30
// seconds after the endpoint reset. For best results, when possible, wait 30
// seconds from endpoint reset to send playback requests to this endpoint.
