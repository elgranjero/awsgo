package costexplorer

// GetReservationCoverage is generated as a reference stub.
// Executable command wiring lives under cmd/costexplorer.go.
//
// Retrieves the reservation coverage for your account, which you can use to see
// how much of your Amazon Elastic Compute Cloud, Amazon ElastiCache, Amazon
// Relational Database Service, or Amazon Redshift usage is covered by a
// reservation. An organization's management account can see the coverage of the
// associated member accounts. This supports dimensions, cost categories, and
// nested expressions. For any time period, you can filter data about reservation
// usage by the following dimensions:
//
// - AZ
//
// - CACHE_ENGINE
//
// - DATABASE_ENGINE
//
// - DEPLOYMENT_OPTION
//
// - INSTANCE_TYPE
//
// - LINKED_ACCOUNT
//
// - OPERATING_SYSTEM
//
// - PLATFORM
//
// - REGION
//
// - SERVICE
//
// - TAG
//
// - TENANCY
//
// To determine valid values for a dimension, use the GetDimensionValues
// operation.
