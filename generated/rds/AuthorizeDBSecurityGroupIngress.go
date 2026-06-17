package rds

// AuthorizeDBSecurityGroupIngress is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Enables ingress to a DBSecurityGroup using one of two forms of authorization.
// First, EC2 or VPC security groups can be added to the DBSecurityGroup if the
// application using the database is running on EC2 or VPC instances. Second, IP
// ranges are available if the application accessing your database is running on
// the internet. Required parameters for this API are one of CIDR range,
// EC2SecurityGroupId for VPC, or (EC2SecurityGroupOwnerId and either
// EC2SecurityGroupName or EC2SecurityGroupId for non-VPC).
//
// You can't authorize ingress from an EC2 security group in one Amazon Web
// Services Region to an Amazon RDS DB instance in another. You can't authorize
// ingress from a VPC security group in one VPC to an Amazon RDS DB instance in
// another.
//
// For an overview of CIDR ranges, go to the [Wikipedia Tutorial].
//
// EC2-Classic was retired on August 15, 2022. If you haven't migrated from
// EC2-Classic to a VPC, we recommend that you migrate as soon as possible. For
// more information, see [Migrate from EC2-Classic to a VPC]in the Amazon EC2 User Guide, the blog [EC2-Classic Networking is Retiring – Here’s How to Prepare], and [Moving a DB instance not in a VPC into a VPC] in the
// Amazon RDS User Guide.
//
// [Migrate from EC2-Classic to a VPC]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html
// [Wikipedia Tutorial]: http://en.wikipedia.org/wiki/Classless_Inter-Domain_Routing
// [EC2-Classic Networking is Retiring – Here’s How to Prepare]: http://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare/
// [Moving a DB instance not in a VPC into a VPC]: https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html
