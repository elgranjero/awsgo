package storagegateway

// ListGateways is generated as a reference stub.
// Executable command wiring lives under cmd/storagegateway.go.
//
// Lists gateways owned by an Amazon Web Services account in an Amazon Web
// Services Region specified in the request. The returned list is ordered by
// gateway Amazon Resource Name (ARN).
//
// By default, the operation returns a maximum of 100 gateways. This operation
// supports pagination that allows you to optionally reduce the number of gateways
// returned in a response.
//
// If you have more gateways than are returned in a response (that is, the
// response returns only a truncated list of your gateways), the response contains
// a marker that you can specify in your next request to fetch the next page of
// gateways.
