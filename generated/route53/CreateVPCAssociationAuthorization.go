package route53

// CreateVPCAssociationAuthorization is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Authorizes the Amazon Web Services account that created a specified VPC to
// submit an AssociateVPCWithHostedZone request to associate the VPC with a
// specified hosted zone that was created by a different account. To submit a
// CreateVPCAssociationAuthorization request, you must use the account that created
// the hosted zone. After you authorize the association, use the account that
// created the VPC to submit an AssociateVPCWithHostedZone request.
//
// If you want to associate multiple VPCs that you created by using one account
// with a hosted zone that you created by using a different account, you must
// submit one authorization request for each VPC.
