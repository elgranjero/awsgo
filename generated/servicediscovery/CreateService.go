package servicediscovery

// CreateService is generated as a reference stub.
// Executable command wiring lives under cmd/servicediscovery.go.
//
// Creates a service. This action defines the configuration for the following
// entities:
//
// - For public and private DNS namespaces, one of the following combinations of
// DNS records in Amazon Route 53:
//
// - A
//
// - AAAA
//
// - A and AAAA
//
// - SRV
//
// - CNAME
//
// - Optionally, a health check
//
// After you create the service, you can submit a [RegisterInstance] request, and Cloud Map uses the
// values in the configuration to create the specified entities.
//
// For the current quota on the number of instances that you can register using
// the same namespace and using the same service, see [Cloud Map quotas]in the Cloud Map Developer
// Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
// [RegisterInstance]: https://docs.aws.amazon.com/cloud-map/latest/api/API_RegisterInstance.html
