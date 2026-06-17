package storagegateway

// DescribeTapes is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Returns a description of virtual tapes that correspond to the specified Amazon
// Resource Names (ARNs). If TapeARN is not specified, returns a description of
// the virtual tapes associated with the specified gateway. This operation is only
// supported for the tape gateway type.
//
// The operation supports pagination. By default, the operation returns a maximum
// of up to 100 tapes. You can optionally specify the Limit field in the body to
// limit the number of tapes in the response. If the number of tapes returned in
// the response is truncated, the response includes a Marker field. You can use
// this Marker value in your subsequent request to retrieve the next set of tapes.
