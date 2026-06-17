package ec2

// DescribeInstanceEventWindows is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Describes the specified event windows or all event windows.
//
// If you specify event window IDs, the output includes information for only the
// specified event windows. If you specify filters, the output includes information
// for only those event windows that meet the filter criteria. If you do not
// specify event windows IDs or filters, the output includes information for all
// event windows, which can affect performance. We recommend that you use
// pagination to ensure that the operation returns quickly and successfully.
//
// For more information, see [Define event windows for scheduled events] in the Amazon EC2 User Guide.
//
// [Define event windows for scheduled events]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/event-windows.html
