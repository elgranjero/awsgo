package kinesisvideo

// GetDataEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/kinesisvideo.go.
//
// Gets an endpoint for a specified stream for either reading or writing. Use this
// endpoint in your application to read from the specified stream (using the
// GetMedia or GetMediaForFragmentList operations) or write to it (using the
// PutMedia operation).
//
// The returned endpoint does not have the API name appended. The client needs to
// add the API name to the returned endpoint.
//
// In the request, specify the stream either by StreamName or StreamARN .
