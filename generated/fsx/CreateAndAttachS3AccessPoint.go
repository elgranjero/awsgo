package fsx

// CreateAndAttachS3AccessPoint is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Creates an S3 access point and attaches it to an Amazon FSx volume. For FSx for
// OpenZFS file systems, the volume must be hosted on a high-availability file
// system, either Single-AZ or Multi-AZ. For more information, see [Accessing your data using Amazon S3 access points]. in the Amazon
// FSx for OpenZFS User Guide.
//
// The requester requires the following permissions to perform these actions:
//
// - fsx:CreateAndAttachS3AccessPoint
//
// - s3:CreateAccessPoint
//
// - s3:GetAccessPoint
//
// - s3:PutAccessPointPolicy
//
// - s3:DeleteAccessPoint
//
// The following actions are related to CreateAndAttachS3AccessPoint :
//
// # DescribeS3AccessPointAttachments
//
// # DetachAndDeleteS3AccessPoint
//
// [Accessing your data using Amazon S3 access points]: https://docs.aws.amazon.com/fsx/latest/OpenZFSGuide/s3accesspoints-for-FSx.html
