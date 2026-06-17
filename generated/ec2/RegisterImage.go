package ec2

// RegisterImage is generated as a reference stub.
// Executable command wiring lives under cmd/ec2.go.
//
// Registers an AMI. When you're creating an instance-store backed AMI,
// registering the AMI is the final step in the creation process. For more
// information about creating AMIs, see [Create an AMI from a snapshot]and [Create an instance-store backed AMI] in the Amazon EC2 User Guide.
//
// If needed, you can deregister an AMI at any time. Any modifications you make to
// an AMI backed by an instance store volume invalidates its registration. If you
// make changes to an image, deregister the previous image and register the new
// image.
//
// # Register a snapshot of a root device volume
//
// You can use RegisterImage to create an Amazon EBS-backed Linux AMI from a
// snapshot of a root device volume. You specify the snapshot using a block device
// mapping. You can't set the encryption state of the volume using the block device
// mapping. If the snapshot is encrypted, or encryption by default is enabled, the
// root volume of an instance launched from the AMI is encrypted.
//
// For more information, see [Create an AMI from a snapshot] and [Use encryption with EBS-backed AMIs] in the Amazon EC2 User Guide.
//
// # Amazon Web Services Marketplace product codes
//
// If any snapshots have Amazon Web Services Marketplace product codes, they are
// copied to the new AMI.
//
// In most cases, AMIs for Windows, RedHat, SUSE, and SQL Server require correct
// licensing information to be present on the AMI. For more information, see [Understand AMI billing information]in
// the Amazon EC2 User Guide. When creating an AMI from a snapshot, the
// RegisterImage operation derives the correct billing information from the
// snapshot's metadata, but this requires the appropriate metadata to be present.
// To verify if the correct billing information was applied, check the
// PlatformDetails field on the new AMI. If the field is empty or doesn't match the
// expected operating system code (for example, Windows, RedHat, SUSE, or SQL), the
// AMI creation was unsuccessful, and you should discard the AMI and instead create
// the AMI from an instance. For more information, see [Create an AMI from an instance]in the Amazon EC2 User
// Guide.
//
// If you purchase a Reserved Instance to apply to an On-Demand Instance that was
// launched from an AMI with a billing product code, make sure that the Reserved
// Instance has the matching billing product code. If you purchase a Reserved
// Instance without the matching billing product code, the Reserved Instance is not
// applied to the On-Demand Instance. For information about how to obtain the
// platform details and billing information of an AMI, see [Understand AMI billing information]in the Amazon EC2 User
// Guide.
//
// [Use encryption with EBS-backed AMIs]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/AMIEncryption.html
// [Understand AMI billing information]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-billing-info.html
// [Create an instance-store backed AMI]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-instance-store.html
// [Create an AMI from an instance]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html#how-to-create-ebs-ami
// [Create an AMI from a snapshot]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/creating-an-ami-ebs.html#creating-launching-ami-from-snapshot
