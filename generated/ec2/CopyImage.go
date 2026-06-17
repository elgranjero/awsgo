package ec2

// CopyImage is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Initiates an AMI copy operation. You must specify the source AMI ID and both
// the source and destination locations. The copy operation must be initiated in
// the destination Region.
//
// CopyImage supports the following source to destination copies:
//
// - Region to Region
//
// - Region to Outpost
//
// - Parent Region to Local Zone
//
// - Local Zone to parent Region
//
// - Between Local Zones with the same parent Region (only supported for certain
// Local Zones)
//
// CopyImage does not support the following source to destination copies:
//
// - Local Zone to non-parent Regions
//
// - Between Local Zones with different parent Regions
//
// - Local Zone to Outpost
//
// - Outpost to Local Zone
//
// - Outpost to Region
//
// - Between Outposts
//
// - Within same Outpost
//
// - Cross-partition copies (use [CreateStoreImageTask]instead)
//
// Destination specification
//
// - Region to Region: The destination Region is the Region in which you
// initiate the copy operation.
//
// - Region to Outpost: Specify the destination using the DestinationOutpostArn
// parameter (the ARN of the Outpost)
//
// - Region to Local Zone, and Local Zone to Local Zone copies: Specify the
// destination using the DestinationAvailabilityZone parameter (the name of the
// destination Local Zone) or DestinationAvailabilityZoneId parameter (the ID of
// the destination Local Zone).
//
// Snapshot encryption
//
// - Region to Outpost: Backing snapshots copied to an Outpost are encrypted by
// default using the default encryption key for the Region or the key that you
// specify. Outposts do not support unencrypted snapshots.
//
// - Region to Local Zone, and Local Zone to Local Zone: Not all Local Zones
// require encrypted snapshots. In Local Zones that require encrypted snapshots,
// backing snapshots are automatically encrypted during copy. In Local Zones where
// encryption is not required, snapshots retain their original encryption state
// (encrypted or unencrypted) by default.
//
// For more information, including the required permissions for copying an AMI,
// see [Copy an Amazon EC2 AMI]in the Amazon EC2 User Guide.
//
// [CreateStoreImageTask]: https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateStoreImageTask.html
// [Copy an Amazon EC2 AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/CopyingAMIs.html
