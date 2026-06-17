package route53

// ChangeCidrCollection is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Creates, changes, or deletes CIDR blocks within a collection. Contains
// authoritative IP information mapping blocks to one or multiple locations.
//
// A change request can update multiple locations in a collection at a time, which
// is helpful if you want to move one or more CIDR blocks from one location to
// another in one transaction, without downtime.
//
// # Limits
//
// The max number of CIDR blocks included in the request is 1000. As a result, big
// updates require multiple API calls.
//
// PUT and DELETE_IF_EXISTS
//
// Use ChangeCidrCollection to perform the following actions:
//
// - PUT : Create a CIDR block within the specified collection.
//
// - DELETE_IF_EXISTS : Delete an existing CIDR block from the collection.
