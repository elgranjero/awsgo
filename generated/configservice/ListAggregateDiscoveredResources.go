package configservice

// ListAggregateDiscoveredResources is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Accepts a resource type and returns a list of resource identifiers that are
// aggregated for a specific resource type across accounts and regions. A resource
// identifier includes the resource type, ID, (if available) the custom resource
// name, source account, and source region. You can narrow the results to include
// only resources that have specific resource IDs, or a resource name, or source
// account ID, or source region.
//
// For example, if the input consists of accountID 12345678910 and the region is
// us-east-1 for resource type AWS::EC2::Instance then the API returns all the EC2
// instance identifiers of accountID 12345678910 and region us-east-1.
