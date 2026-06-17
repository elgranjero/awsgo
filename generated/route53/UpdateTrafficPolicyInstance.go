package route53

// UpdateTrafficPolicyInstance is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// After you submit a UpdateTrafficPolicyInstance request, there's a brief delay
// while Route 53 creates the resource record sets that are specified in the
// traffic policy definition. Use GetTrafficPolicyInstance with the id of updated
// traffic policy instance confirm that the UpdateTrafficPolicyInstance request
// completed successfully. For more information, see the State response element.
//
// Updates the resource record sets in a specified hosted zone that were created
// based on the settings in a specified traffic policy version.
//
// When you update a traffic policy instance, Amazon Route 53 continues to respond
// to DNS queries for the root resource record set name (such as example.com) while
// it replaces one group of resource record sets with another. Route 53 performs
// the following operations:
//
// - Route 53 creates a new group of resource record sets based on the specified
// traffic policy. This is true regardless of how significant the differences are
// between the existing resource record sets and the new resource record sets.
//
// - When all of the new resource record sets have been created, Route 53 starts
// to respond to DNS queries for the root resource record set name (such as
// example.com) by using the new resource record sets.
//
// - Route 53 deletes the old group of resource record sets that are associated
// with the root resource record set name.
