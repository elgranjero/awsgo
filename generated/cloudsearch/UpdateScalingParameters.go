package cloudsearch

// UpdateScalingParameters is generated as a reference stub.
// Executable command wiring lives under cmd/cloudsearch.go.
//
// Configures scaling parameters for a domain. A domain's scaling parameters
// specify the desired search instance type and replication count. Amazon
// CloudSearch will still automatically scale your domain based on the volume of
// data and traffic, but not below the desired instance type and replication count.
// If the Multi-AZ option is enabled, these values control the resources used per
// Availability Zone. For more information, see [Configuring Scaling Options]in the Amazon CloudSearch
// Developer Guide.
//
// [Configuring Scaling Options]: http://docs.aws.amazon.com/cloudsearch/latest/developerguide/configuring-scaling-options.html
