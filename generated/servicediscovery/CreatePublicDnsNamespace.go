package servicediscovery

// CreatePublicDnsNamespace is generated as a reference stub.
// Executable command wiring lives under cmd/servicediscovery.go.
//
// Creates a public namespace based on DNS, which is visible on the internet. The
// namespace defines your service naming scheme. For example, if you name your
// namespace example.com and name your service backend , the resulting DNS name for
// the service is backend.example.com . You can discover instances that were
// registered with a public DNS namespace by using either a DiscoverInstances
// request or using DNS. For the current quota on the number of namespaces that you
// can create using the same Amazon Web Services account, see [Cloud Map quotas]in the Cloud Map
// Developer Guide.
//
// The CreatePublicDnsNamespace API operation is not supported in the Amazon Web
// Services GovCloud (US) Regions.
//
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
