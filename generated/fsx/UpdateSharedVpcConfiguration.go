package fsx

// UpdateSharedVpcConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/fsx.go.
//
// Configures whether participant accounts in your organization can create Amazon
// FSx for NetApp ONTAP Multi-AZ file systems in subnets that are shared by a
// virtual private cloud (VPC) owner. For more information, see the [Amazon FSx for NetApp ONTAP User Guide].
//
// We strongly recommend that participant-created Multi-AZ file systems in the
// shared VPC are deleted before you disable this feature. Once the feature is
// disabled, these file systems will enter a MISCONFIGURED state and behave like
// Single-AZ file systems. For more information, see [Important considerations before disabling shared VPC support for Multi-AZ file systems].
//
// [Amazon FSx for NetApp ONTAP User Guide]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/maz-shared-vpc.html
// [Important considerations before disabling shared VPC support for Multi-AZ file systems]: https://docs.aws.amazon.com/fsx/latest/ONTAPGuide/maz-shared-vpc.html#disabling-maz-vpc-sharing
