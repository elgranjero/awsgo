package sagemaker

// DeleteEndpointConfig is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deletes an endpoint configuration. The DeleteEndpointConfig API deletes only
// the specified configuration. It does not delete endpoints created using the
// configuration.
//
// You must not delete an EndpointConfig in use by an endpoint that is live or
// while the UpdateEndpoint or CreateEndpoint operations are being performed on
// the endpoint. If you delete the EndpointConfig of an endpoint that is active or
// being created or updated you may lose visibility into the instance type the
// endpoint is using. The endpoint must be deleted in order to stop incurring
// charges.
