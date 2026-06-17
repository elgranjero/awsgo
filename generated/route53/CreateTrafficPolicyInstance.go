package route53

// CreateTrafficPolicyInstance is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Creates resource record sets in a specified hosted zone based on the settings
// in a specified traffic policy version. In addition, CreateTrafficPolicyInstance
// associates the resource record sets with a specified domain name (such as
// example.com) or subdomain name (such as www.example.com). Amazon Route 53
// responds to DNS queries for the domain or subdomain name by using the resource
// record sets that CreateTrafficPolicyInstance created.
//
// After you submit an CreateTrafficPolicyInstance request, there's a brief delay
// while Amazon Route 53 creates the resource record sets that are specified in the
// traffic policy definition. Use GetTrafficPolicyInstance with the id of new
// traffic policy instance to confirm that the CreateTrafficPolicyInstance request
// completed successfully. For more information, see the State response element.
