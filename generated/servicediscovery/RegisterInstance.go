package servicediscovery

// RegisterInstance is generated as a reference stub.
// Executable command wiring lives under cmd/servicediscovery.go.
//
// Creates or updates one or more records and, optionally, creates a health check
// based on the settings in a specified service. When you submit a RegisterInstance
// request, the following occurs:
//
// - For each DNS record that you define in the service that's specified by
// ServiceId , a record is created or updated in the hosted zone that's
// associated with the corresponding namespace.
//
// - If the service includes HealthCheckConfig , a health check is created based
// on the settings in the health check configuration.
//
// - The health check, if any, is associated with each of the new or updated
// records.
//
// One RegisterInstance request must complete before you can submit another
// request and specify the same service ID and instance ID.
//
// For more information, see [CreateService].
//
// When Cloud Map receives a DNS query for the specified DNS name, it returns the
// applicable value:
//
// - If the health check is healthy: returns all the records
//
// - If the health check is unhealthy: returns the applicable value for the last
// healthy instance
//
// - If you didn't specify a health check configuration: returns all the records
//
// For the current quota on the number of instances that you can register using
// the same namespace and using the same service, see [Cloud Map quotas]in the Cloud Map Developer
// Guide.
//
// [CreateService]: https://docs.aws.amazon.com/cloud-map/latest/api/API_CreateService.html
// [Cloud Map quotas]: https://docs.aws.amazon.com/cloud-map/latest/dg/cloud-map-limits.html
