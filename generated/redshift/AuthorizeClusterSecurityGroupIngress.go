package redshift

// AuthorizeClusterSecurityGroupIngress is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Adds an inbound (ingress) rule to an Amazon Redshift security group. Depending
// on whether the application accessing your cluster is running on the Internet or
// an Amazon EC2 instance, you can authorize inbound access to either a Classless
// Interdomain Routing (CIDR)/Internet Protocol (IP) range or to an Amazon EC2
// security group. You can add as many as 20 ingress rules to an Amazon Redshift
// security group.
//
// If you authorize access to an Amazon EC2 security group, specify
// EC2SecurityGroupName and EC2SecurityGroupOwnerId. The Amazon EC2 security group
// and Amazon Redshift cluster must be in the same Amazon Web Services Region.
//
// If you authorize access to a CIDR/IP address range, specify CIDRIP. For an
// overview of CIDR blocks, see the Wikipedia article on [Classless Inter-Domain Routing].
//
// You must also associate the security group with a cluster so that clients
// running on these IP addresses or the EC2 instance are authorized to connect to
// the cluster. For information about managing security groups, go to [Working with Security Groups]in the
// Amazon Redshift Cluster Management Guide.
//
// [Classless Inter-Domain Routing]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
// [Working with Security Groups]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-security-groups.html
