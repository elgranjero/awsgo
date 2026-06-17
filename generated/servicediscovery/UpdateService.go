package servicediscovery

// UpdateService is generated as a reference stub.
// Executable command wiring lives under cmd/servicediscovery.go.
//
// Submits a request to perform the following operations:
//
// - Update the TTL setting for existing DnsRecords configurations
//
// - Add, update, or delete HealthCheckConfig for a specified service
//
// You can't add, update, or delete a HealthCheckCustomConfig configuration.
//
// For public and private DNS namespaces, note the following:
//
// - If you omit any existing DnsRecords or HealthCheckConfig configurations from
// an UpdateService request, the configurations are deleted from the service.
//
// - If you omit an existing HealthCheckCustomConfig configuration from an
// UpdateService request, the configuration isn't deleted from the service.
//
// You can't call UpdateService and update settings in the following scenarios:
//
// - When the service is associated with an HTTP namespace
//
// - When the service is associated with a shared namespace and contains
// instances that were registered by Amazon Web Services accounts other than the
// account making the UpdateService call
//
// When you update settings for a service, Cloud Map also updates the
// corresponding settings in all the records and health checks that were created by
// using the specified service.
