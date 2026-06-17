package route53

// ChangeResourceRecordSets is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Creates, changes, or deletes a resource record set, which contains
// authoritative DNS information for a specified domain name or subdomain name. For
// example, you can use ChangeResourceRecordSets to create a resource record set
// that routes traffic for test.example.com to a web server that has an IP address
// of 192.0.2.44.
//
// # Deleting Resource Record Sets
//
// To delete a resource record set, you must specify all the same values that you
// specified when you created it.
//
// # Change Batches and Transactional Changes
//
// The request body must include a document with a ChangeResourceRecordSetsRequest
// element. The request body contains a list of change items, known as a change
// batch. Change batches are considered transactional changes. Route 53 validates
// the changes in the request and then either makes all or none of the changes in
// the change batch request. This ensures that DNS routing isn't adversely affected
// by partial changes to the resource record sets in a hosted zone.
//
// For example, suppose a change batch request contains two changes: it deletes
// the CNAME resource record set for www.example.com and creates an alias resource
// record set for www.example.com. If validation for both records succeeds, Route
// 53 deletes the first resource record set and creates the second resource record
// set in a single operation. If validation for either the DELETE or the CREATE
// action fails, then the request is canceled, and the original CNAME record
// continues to exist.
//
// If you try to delete the same resource record set more than once in a single
// change batch, Route 53 returns an InvalidChangeBatch error.
//
// # Traffic Flow
//
// To create resource record sets for complex routing configurations, use either
// the traffic flow visual editor in the Route 53 console or the API actions for
// traffic policies and traffic policy instances. Save the configuration as a
// traffic policy, then associate the traffic policy with one or more domain names
// (such as example.com) or subdomain names (such as www.example.com), in the same
// hosted zone or in multiple hosted zones. You can roll back the updates if the
// new configuration isn't performing as expected. For more information, see [Using Traffic Flow to Route DNS Traffic]in
// the Amazon Route 53 Developer Guide.
//
// # Create, Delete, and Upsert
//
// Use ChangeResourceRecordsSetsRequest to perform the following actions:
//
// - CREATE : Creates a resource record set that has the specified values.
//
// - DELETE : Deletes an existing resource record set that has the specified
// values.
//
// - UPSERT : If a resource set doesn't exist, Route 53 creates it. If a resource
// set exists Route 53 updates it with the values in the request.
//
// # Syntaxes for Creating, Updating, and Deleting Resource Record Sets
//
// The syntax for a request depends on the type of resource record set that you
// want to create, delete, or update, such as weighted, alias, or failover. The XML
// elements in your request must appear in the order listed in the syntax.
//
// For an example for each type of resource record set, see "Examples."
//
// Don't refer to the syntax in the "Parameter Syntax" section, which includes all
// of the elements for every kind of resource record set that you can create,
// delete, or update by using ChangeResourceRecordSets .
//
// # Change Propagation to Route 53 DNS Servers
//
// When you submit a ChangeResourceRecordSets request, Route 53 propagates your
// changes to all of the Route 53 authoritative DNS servers managing the hosted
// zone. While your changes are propagating, GetChange returns a status of PENDING
// . When propagation is complete, GetChange returns a status of INSYNC . Changes
// generally propagate to all Route 53 name servers managing the hosted zone within
// 60 seconds. For more information, see [GetChange].
//
// # Limits on ChangeResourceRecordSets Requests
//
// For information about the limits on a ChangeResourceRecordSets request, see [Limits] in
// the Amazon Route 53 Developer Guide.
//
// [Limits]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/DNSLimitations.html
// [GetChange]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetChange.html
// [Using Traffic Flow to Route DNS Traffic]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/traffic-flow.html
