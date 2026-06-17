package gamelift

// DescribeEC2InstanceLimits is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Retrieves the instance limits and current utilization for an Amazon Web
// Services Region or location. Instance limits control the number of instances,
// per instance type, per location, that your Amazon Web Services account can use.
// Learn more at [Amazon EC2 Instance Types]. The information returned includes the maximum number of
// instances allowed and your account's current usage across all fleets. This
// information can affect your ability to scale your Amazon GameLift Servers
// fleets. You can request a limit increase for your account by using the Service
// limits page in the Amazon GameLift Servers console.
//
// Instance limits differ based on whether the instances are deployed in a fleet's
// home Region or in a remote location. For remote locations, limits also differ
// based on the combination of home Region and remote location. All requests must
// specify an Amazon Web Services Region (either explicitly or as your default
// settings). To get the limit for a remote location, you must also specify the
// location. For example, the following requests all return different results:
//
// - Request specifies the Region ap-northeast-1 with no location. The result is
// limits and usage data on all instance types that are deployed in us-east-2 ,
// by all of the fleets that reside in ap-northeast-1 .
//
// - Request specifies the Region us-east-1 with location ca-central-1 . The
// result is limits and usage data on all instance types that are deployed in
// ca-central-1 , by all of the fleets that reside in us-east-2 . These limits do
// not affect fleets in any other Regions that deploy instances to ca-central-1 .
//
// - Request specifies the Region eu-west-1 with location ca-central-1 . The
// result is limits and usage data on all instance types that are deployed in
// ca-central-1 , by all of the fleets that reside in eu-west-1 .
//
// This operation can be used in the following ways:
//
// - To get limit and usage data for all instance types that are deployed in an
// Amazon Web Services Region by fleets that reside in the same Region: Specify the
// Region only. Optionally, specify a single instance type to retrieve information
// for.
//
// - To get limit and usage data for all instance types that are deployed to a
// remote location by fleets that reside in different Amazon Web Services Region:
// Provide both the Amazon Web Services Region and the remote location. Optionally,
// specify a single instance type to retrieve information for.
//
// If successful, an EC2InstanceLimits object is returned with limits and usage
// data for each requested instance type.
//
// # Learn more
//
// [Setting up Amazon GameLift Servers fleets]
//
// [Amazon EC2 Instance Types]: http://aws.amazon.com/ec2/instance-types/
// [Setting up Amazon GameLift Servers fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
