package route53

// ListTrafficPolicyInstancesByPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Gets information about the traffic policy instances that you created by using a
// specify traffic policy version.
//
// After you submit a CreateTrafficPolicyInstance or an UpdateTrafficPolicyInstance
// request, there's a brief delay while Amazon Route 53 creates the resource record
// sets that are specified in the traffic policy definition. For more information,
// see the State response element.
//
// Route 53 returns a maximum of 100 items in each response. If you have a lot of
// traffic policy instances, you can use the MaxItems parameter to list them in
// groups of up to 100.
