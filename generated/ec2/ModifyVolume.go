package ec2

// ModifyVolume is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// You can modify several parameters of an existing EBS volume, including volume
// size, volume type, and IOPS capacity. If your EBS volume is attached to a
// current-generation EC2 instance type, you might be able to apply these changes
// without stopping the instance or detaching the volume from it. For more
// information about modifying EBS volumes, see [Amazon EBS Elastic Volumes]in the Amazon EBS User Guide.
//
// When you complete a resize operation on your volume, you need to extend the
// volume's file-system size to take advantage of the new storage capacity. For
// more information, see [Extend the file system].
//
// For more information, see [Monitor the progress of volume modifications] in the Amazon EBS User Guide.
//
// With previous-generation instance types, resizing an EBS volume might require
// detaching and reattaching the volume or stopping and restarting the instance.
//
// After you initiate a volume modification, you must wait for that modification
// to reach the completed state before you can initiate another modification for
// the same volume. You can modify a volume up to four times within a rolling
// 24-hour period, as long as the volume is in the in-use or available state, and
// all previous modifications for that volume are completed . If you exceed this
// limit, you get an error message that indicates when you can perform your next
// modification.
//
// [Monitor the progress of volume modifications]: https://docs.aws.amazon.com/ebs/latest/userguide/monitoring-volume-modifications.html
// [Amazon EBS Elastic Volumes]: https://docs.aws.amazon.com/ebs/latest/userguide/ebs-modify-volume.html
// [Extend the file system]: https://docs.aws.amazon.com/ebs/latest/userguide/recognize-expanded-volume-linux.html
