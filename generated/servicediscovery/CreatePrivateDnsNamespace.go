package servicediscovery

// CreatePrivateDnsNamespace is generated as a reference stub.
// Executable command wiring lives under cmd/servicediscovery.go.
//
// Creates a private namespace based on DNS, which is visible only inside a
// specified Amazon VPC. The namespace defines your service naming scheme. For
// example, if you name your namespace example.com and name your service backend ,
// the resulting DNS name for the service is backend.example.com . Service
// instances that are registered using a private DNS namespace can be discovered
// using either a DiscoverInstances request or using DNS. For the current quota on
// the number of namespaces that you can create using the same Amazon Web Services
// account, see [Cloud Map quotas]in the Cloud Map Developer Guide.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
