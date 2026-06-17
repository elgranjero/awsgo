package lightsail

// ExportSnapshot is generated as a reference stub.
// Executable command wiring lives under cmd/lightsail.go.
//
// Exports an Amazon Lightsail instance or block storage disk snapshot to Amazon
// Elastic Compute Cloud (Amazon EC2). This operation results in an export snapshot
// record that can be used with the create cloud formation stack operation to
// create new Amazon EC2 instances.
//
// Exported instance snapshots appear in Amazon EC2 as Amazon Machine Images
// (AMIs), and the instance system disk appears as an Amazon Elastic Block Store
// (Amazon EBS) volume. Exported disk snapshots appear in Amazon EC2 as Amazon EBS
// volumes. Snapshots are exported to the same Amazon Web Services Region in Amazon
// EC2 as the source Lightsail snapshot.
//
// The export snapshot operation supports tag-based access control via resource
// tags applied to the resource identified by source snapshot name . For more
// information, see the [Amazon Lightsail Developer Guide].
//
// Use the get instance snapshots or get disk snapshots operations to get a list
// of snapshots that you can export to Amazon EC2.
//
// [Amazon Lightsail Developer Guide]: https://docs.aws.amazon.com/lightsail/latest/userguide/amazon-lightsail-controlling-access-using-tags
