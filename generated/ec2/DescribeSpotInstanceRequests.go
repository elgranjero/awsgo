package ec2

// DescribeSpotInstanceRequests is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the specified Spot Instance requests.
//
// You can use DescribeSpotInstanceRequests to find a running Spot Instance by
// examining the response. If the status of the Spot Instance is fulfilled , the
// instance ID appears in the response and contains the identifier of the instance.
// Alternatively, you can use [DescribeInstances]with a filter to look for instances where the
// instance lifecycle is spot .
//
// We recommend that you set MaxResults to a value between 5 and 1000 to limit the
// number of items returned. This paginates the output, which makes the list more
// manageable and returns the items faster. If the list of items exceeds your
// MaxResults value, then that number of items is returned along with a NextToken
// value that can be passed to a subsequent DescribeSpotInstanceRequests request
// to retrieve the remaining items.
//
// Spot Instance requests are deleted four hours after they are canceled and their
// instances are terminated.
//
// [DescribeInstances]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances
