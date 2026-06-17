package sagemaker

// UpdateEndpoint is generated as a reference stub.
// Executable command wiring lives under cmd/sagemaker.go.
//
// Deploys the EndpointConfig specified in the request to a new fleet of
// instances. SageMaker shifts endpoint traffic to the new instances with the
// updated endpoint configuration and then deletes the old instances using the
// previous EndpointConfig (there is no availability loss). For more information
// about how to control the update and traffic shifting process, see [Update models in production].
//
// When SageMaker receives the request, it sets the endpoint status to Updating .
// After updating the endpoint, it sets the status to InService . To check the
// status of an endpoint, use the [DescribeEndpoint]API.
//
// You must not delete an EndpointConfig in use by an endpoint that is live or
// while the UpdateEndpoint or CreateEndpoint operations are being performed on
// the endpoint. To update an endpoint, you must create a new EndpointConfig .
//
// If you delete the EndpointConfig of an endpoint that is active or being created
// or updated you may lose visibility into the instance type the endpoint is using.
// The endpoint must be deleted in order to stop incurring charges.
//
// [DescribeEndpoint]: https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeEndpoint.html
// [Update models in production]: https://docs.aws.amazon.com/sagemaker/latest/dg/deployment-guardrails.html
