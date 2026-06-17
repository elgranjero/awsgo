package route53

// AssociateVPCWithHostedZone is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Associates an Amazon VPC with a private hosted zone.
//
// To perform the association, the VPC and the private hosted zone must already
// exist. You can't convert a public hosted zone into a private hosted zone.
//
// If you want to associate a VPC that was created by using one Amazon Web
// Services account with a private hosted zone that was created by using a
// different account, the Amazon Web Services account that created the private
// hosted zone must first submit a CreateVPCAssociationAuthorization request. Then
// the account that created the VPC must submit an AssociateVPCWithHostedZone
// request.
//
// When granting access, the hosted zone and the Amazon VPC must belong to the
// same partition. A partition is a group of Amazon Web Services Regions. Each
// Amazon Web Services account is scoped to one partition.
//
// The following are the supported partitions:
//
// - aws - Amazon Web Services Regions
//
// - aws-cn - China Regions
//
// - aws-us-gov - Amazon Web Services GovCloud (US) Region
//
// For more information, see [Access Management] in the Amazon Web Services General Reference.
//
// [Access Management]: https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html
