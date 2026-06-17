package configservice

// SelectAggregateResourceConfig is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Accepts a structured query language (SQL) SELECT command and an aggregator to
// query configuration state of Amazon Web Services resources across multiple
// accounts and regions, performs the corresponding search, and returns resource
// configurations matching the properties.
//
// For more information about query components, see the [Query Components] section in the Config
// Developer Guide.
//
// If you run an aggregation query (i.e., using GROUP BY or using aggregate
// functions such as COUNT ; e.g., SELECT resourceId, COUNT(*) WHERE resourceType
// = 'AWS::IAM::Role' GROUP BY resourceId ) and do not specify the MaxResults or
// the Limit query parameters, the default page size is set to 500.
//
// If you run a non-aggregation query (i.e., not using GROUP BY or aggregate
// function; e.g., SELECT * WHERE resourceType = 'AWS::IAM::Role' ) and do not
// specify the MaxResults or the Limit query parameters, the default page size is
// set to 25.
//
// [Query Components]: https://docs.aws.amazon.com/config/latest/developerguide/query-components.html
