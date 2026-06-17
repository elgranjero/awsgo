package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// efsCmd represents the efs command
var _efsCmd = &cobra.Command{
	Use:   "efs",
	Short: "AWS efs CLI",
	Run: func(cmd *cobra.Command, args []string) {
		_awsOutput = resolveAWSOutput(_awsProfile, cmd.Flags().Changed("output"))
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := efs.NewFromConfig(cfg)
		if _efsCreateAccessPoint {
			efs_CreateAccessPoint(cfg, client)
			return
		}
		if _efsCreateFileSystem {
			efs_CreateFileSystem(cfg, client)
			return
		}
		if _efsCreateMountTarget {
			efs_CreateMountTarget(cfg, client)
			return
		}
		if _efsCreateReplicationConfiguration {
			efs_CreateReplicationConfiguration(cfg, client)
			return
		}
		if _efsCreateTags {
			efs_CreateTags(cfg, client)
			return
		}
		if _efsDeleteAccessPoint {
			efs_DeleteAccessPoint(cfg, client)
			return
		}
		if _efsDeleteFileSystem {
			efs_DeleteFileSystem(cfg, client)
			return
		}
		if _efsDeleteFileSystemPolicy {
			efs_DeleteFileSystemPolicy(cfg, client)
			return
		}
		if _efsDeleteMountTarget {
			efs_DeleteMountTarget(cfg, client)
			return
		}
		if _efsDeleteReplicationConfiguration {
			efs_DeleteReplicationConfiguration(cfg, client)
			return
		}
		if _efsDeleteTags {
			efs_DeleteTags(cfg, client)
			return
		}
		if _efsDescribeAccessPoints {
			efs_DescribeAccessPoints(cfg, client)
			return
		}
		if _efsDescribeAccountPreferences {
			efs_DescribeAccountPreferences(cfg, client)
			return
		}
		if _efsDescribeBackupPolicy {
			efs_DescribeBackupPolicy(cfg, client)
			return
		}
		if _efsDescribeFileSystemPolicy {
			efs_DescribeFileSystemPolicy(cfg, client)
			return
		}
		if _efsDescribeFileSystems {
			efs_DescribeFileSystems(cfg, client)
			return
		}
		if _efsDescribeLifecycleConfiguration {
			efs_DescribeLifecycleConfiguration(cfg, client)
			return
		}
		if _efsDescribeMountTargetSecurityGroups {
			efs_DescribeMountTargetSecurityGroups(cfg, client)
			return
		}
		if _efsDescribeMountTargets {
			efs_DescribeMountTargets(cfg, client)
			return
		}
		if _efsDescribeReplicationConfigurations {
			efs_DescribeReplicationConfigurations(cfg, client)
			return
		}
		if _efsDescribeTags {
			efs_DescribeTags(cfg, client)
			return
		}
		if _efsListTagsForResource {
			efs_ListTagsForResource(cfg, client)
			return
		}
		if _efsModifyMountTargetSecurityGroups {
			efs_ModifyMountTargetSecurityGroups(cfg, client)
			return
		}
		if _efsPutAccountPreferences {
			efs_PutAccountPreferences(cfg, client)
			return
		}
		if _efsPutBackupPolicy {
			efs_PutBackupPolicy(cfg, client)
			return
		}
		if _efsPutFileSystemPolicy {
			efs_PutFileSystemPolicy(cfg, client)
			return
		}
		if _efsPutLifecycleConfiguration {
			efs_PutLifecycleConfiguration(cfg, client)
			return
		}
		if _efsTagResource {
			efs_TagResource(cfg, client)
			return
		}
		if _efsUntagResource {
			efs_UntagResource(cfg, client)
			return
		}
		if _efsUpdateFileSystem {
			efs_UpdateFileSystem(cfg, client)
			return
		}
		if _efsUpdateFileSystemProtection {
			efs_UpdateFileSystemProtection(cfg, client)
			return
		}

	},
}

var (
	_efsCreateAccessPoint                 bool
	_efsCreateFileSystem                  bool
	_efsCreateMountTarget                 bool
	_efsCreateReplicationConfiguration    bool
	_efsCreateTags                        bool
	_efsDeleteAccessPoint                 bool
	_efsDeleteFileSystem                  bool
	_efsDeleteFileSystemPolicy            bool
	_efsDeleteMountTarget                 bool
	_efsDeleteReplicationConfiguration    bool
	_efsDeleteTags                        bool
	_efsDescribeAccessPoints              bool
	_efsDescribeAccountPreferences        bool
	_efsDescribeBackupPolicy              bool
	_efsDescribeFileSystemPolicy          bool
	_efsDescribeFileSystems               bool
	_efsDescribeLifecycleConfiguration    bool
	_efsDescribeMountTargetSecurityGroups bool
	_efsDescribeMountTargets              bool
	_efsDescribeReplicationConfigurations bool
	_efsDescribeTags                      bool
	_efsListTagsForResource               bool
	_efsModifyMountTargetSecurityGroups   bool
	_efsPutAccountPreferences             bool
	_efsPutBackupPolicy                   bool
	_efsPutFileSystemPolicy               bool
	_efsPutLifecycleConfiguration         bool
	_efsTagResource                       bool
	_efsUntagResource                     bool
	_efsUpdateFileSystem                  bool
	_efsUpdateFileSystemProtection        bool

	_efsAccessPointId                  string
	_efsAvailabilityZoneName           string
	_efsBackup                         string
	_efsBackupPolicy                   string
	_efsBypassPolicyLockoutSafetyCheck string
	_efsClientToken                    string
	_efsCreationToken                  string
	_efsDeletionMode                   string
	_efsDestinations                   string
	_efsEncrypted                      string
	_efsFileSystemId                   string
	_efsIpAddress                      string
	_efsIpAddressType                  string
	_efsIpv6Address                    string
	_efsKmsKeyId                       string
	_efsLifecyclePolicies              string
	_efsMarker                         string
	_efsMaxItems                       string
	_efsMaxResults                     string
	_efsMountTargetId                  string
	_efsNextToken                      string
	_efsPerformanceMode                string
	_efsPolicy                         string
	_efsPosixUser                      string
	_efsProvisionedThroughputInMibps   string
	_efsReplicationOverwriteProtection string
	_efsResourceId                     string
	_efsResourceIdType                 string
	_efsRootDirectory                  string
	_efsSecurityGroups                 []string
	_efsSourceFileSystemId             string
	_efsSubnetId                       string
	_efsTagKeys                        []string
	_efsTags                           string
	_efsThroughputMode                 string
)

// Creates an EFS access point. An access point is an application-specific view
// into an EFS file system that applies an operating system user and group, and a
// file system path, to any file system request made through the access point. The
// operating system user and group override any identity information provided by
// the NFS client. The file system path is exposed as the access point's root
// directory. Applications using the access point can only access data in the
// application's own directory and any subdirectories. A file system can have a
// maximum of 10,000 access points unless you request an increase. To learn more,
// see [Mounting a file system using EFS access points].
//
// If multiple requests to create access points on the same file system are sent
// in quick succession, and the file system is near the limit of access points, you
// may experience a throttling response for these requests. This is to ensure that
// the file system does not exceed the stated access point limit.
//
// This operation requires permissions for the elasticfilesystem:CreateAccessPoint
// action.
//
// Access points can be tagged on creation. If tags are specified in the creation
// action, IAM performs additional authorization on the
// elasticfilesystem:TagResource action to verify if users have permissions to
// create tags. Therefore, you must grant explicit permissions to use the
// elasticfilesystem:TagResource action. For more information, see [Granting permissions to tag resources during creation].
//
// [Mounting a file system using EFS access points]: https://docs.aws.amazon.com/efs/latest/ug/efs-access-points.html
// [Granting permissions to tag resources during creation]: https://docs.aws.amazon.com/efs/latest/ug/using-tags-efs.html#supported-iam-actions-tagging.html
func efs_CreateAccessPoint(cfg aws.Config, client *efs.Client) {
	input := &efs.CreateAccessPointInput{
		// ClientToken: *string, // Required
		// FileSystemId: *string, // Required
	}

	if len(_efsClientToken) > 0 {
		input.ClientToken = aws.String(_efsClientToken)
	}
	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsPosixUser) > 0 {
		if err := assignInputField(input, "PosixUser", _efsPosixUser); err != nil {
			log.Errorf("invalid --posix-user: %s", err.Error())
			return
		}
	}
	if len(_efsRootDirectory) > 0 {
		if err := assignInputField(input, "RootDirectory", _efsRootDirectory); err != nil {
			log.Errorf("invalid --root-directory: %s", err.Error())
			return
		}
	}
	if len(_efsTags) > 0 {
		if err := assignInputField(input, "Tags", _efsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new, empty file system. The operation requires a creation token in
// the request that Amazon EFS uses to ensure idempotent creation (calling the
// operation with same creation token has no effect). If a file system does not
// currently exist that is owned by the caller's Amazon Web Services account with
// the specified creation token, this operation does the following:
//
// - Creates a new, empty file system. The file system will have an Amazon EFS
// assigned ID, and an initial lifecycle state creating .
//
// - Returns with the description of the created file system.
//
// Otherwise, this operation returns a FileSystemAlreadyExists error with the ID
// of the existing file system.
//
// For basic use cases, you can use a randomly generated UUID for the creation
// token.
//
// The idempotent operation allows you to retry a CreateFileSystem call without
// risk of creating an extra file system. This can happen when an initial call
// fails in a way that leaves it uncertain whether or not a file system was
// actually created. An example might be that a transport level timeout occurred or
// your connection was reset. As long as you use the same creation token, if the
// initial call had succeeded in creating a file system, the client can learn of
// its existence from the FileSystemAlreadyExists error.
//
// For more information, see [Creating a file system] in the Amazon EFS User Guide.
//
// The CreateFileSystem call returns while the file system's lifecycle state is
// still creating . You can check the file system creation status by calling the DescribeFileSystems
// operation, which among other things returns the file system state.
//
// This operation accepts an optional PerformanceMode parameter that you choose
// for your file system. We recommend generalPurpose PerformanceMode for all file
// systems. The maxIO mode is a previous generation performance type that is
// designed for highly parallelized workloads that can tolerate higher latencies
// than the generalPurpose mode. MaxIO mode is not supported for One Zone file
// systems or file systems that use Elastic throughput.
//
// The PerformanceMode can't be changed after the file system has been created.
// For more information, see [Amazon EFS performance modes].
//
// You can set the throughput mode for the file system using the ThroughputMode
// parameter.
//
// After the file system is fully created, Amazon EFS sets its lifecycle state to
// available , at which point you can create one or more mount targets for the file
// system in your VPC. For more information, see CreateMountTarget. You mount your Amazon EFS file
// system on an EC2 instances in your VPC by using the mount target. For more
// information, see [Amazon EFS: How it Works].
//
// This operation requires permissions for the elasticfilesystem:CreateFileSystem
// action.
//
// File systems can be tagged on creation. If tags are specified in the creation
// action, IAM performs additional authorization on the
// elasticfilesystem:TagResource action to verify if users have permissions to
// create tags. Therefore, you must grant explicit permissions to use the
// elasticfilesystem:TagResource action. For more information, see [Granting permissions to tag resources during creation].
//
// [Creating a file system]: https://docs.aws.amazon.com/efs/latest/ug/creating-using-create-fs.html#creating-using-create-fs-part1
// [Amazon EFS: How it Works]: https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html
// [Amazon EFS performance modes]: https://docs.aws.amazon.com/efs/latest/ug/performance.html#performancemodes.html
// [Granting permissions to tag resources during creation]: https://docs.aws.amazon.com/efs/latest/ug/using-tags-efs.html#supported-iam-actions-tagging.html
func efs_CreateFileSystem(cfg aws.Config, client *efs.Client) {
	input := &efs.CreateFileSystemInput{
		// CreationToken: *string, // Required
	}

	if len(_efsCreationToken) > 0 {
		input.CreationToken = aws.String(_efsCreationToken)
	}
	if len(_efsAvailabilityZoneName) > 0 {
		input.AvailabilityZoneName = aws.String(_efsAvailabilityZoneName)
	}
	if len(_efsBackup) > 0 {
		if err := assignInputField(input, "Backup", _efsBackup); err != nil {
			log.Errorf("invalid --backup: %s", err.Error())
			return
		}
	}
	if len(_efsEncrypted) > 0 {
		if err := assignInputField(input, "Encrypted", _efsEncrypted); err != nil {
			log.Errorf("invalid --encrypted: %s", err.Error())
			return
		}
	}
	if len(_efsKmsKeyId) > 0 {
		input.KmsKeyId = aws.String(_efsKmsKeyId)
	}
	if len(_efsPerformanceMode) > 0 {
		if err := assignInputField(input, "PerformanceMode", _efsPerformanceMode); err != nil {
			log.Errorf("invalid --performance-mode: %s", err.Error())
			return
		}
	}
	if len(_efsProvisionedThroughputInMibps) > 0 {
		if err := assignInputField(input, "ProvisionedThroughputInMibps", _efsProvisionedThroughputInMibps); err != nil {
			log.Errorf("invalid --provisioned-throughput-in-mibps: %s", err.Error())
			return
		}
	}
	if len(_efsTags) > 0 {
		if err := assignInputField(input, "Tags", _efsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_efsThroughputMode) > 0 {
		if err := assignInputField(input, "ThroughputMode", _efsThroughputMode); err != nil {
			log.Errorf("invalid --throughput-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a mount target for a file system. You can then mount the file system on
// EC2 instances by using the mount target.
//
// You can create one mount target in each Availability Zone in your VPC. All EC2
// instances in a VPC within a given Availability Zone share a single mount target
// for a given file system. If you have multiple subnets in an Availability Zone,
// you create a mount target in one of the subnets. EC2 instances do not need to be
// in the same subnet as the mount target in order to access their file system.
//
// You can create only one mount target for a One Zone file system. You must
// create that mount target in the same Availability Zone in which the file system
// is located. Use the AvailabilityZoneName and AvailabiltyZoneId properties in
// the DescribeFileSystemsresponse object to get this information. Use the subnetId associated with
// the file system's Availability Zone when creating the mount target.
//
// For more information, see [Amazon EFS: How it Works].
//
// To create a mount target for a file system, the file system's lifecycle state
// must be available . For more information, see DescribeFileSystems.
//
// In the request, provide the following:
//
// - The file system ID for which you are creating the mount target.
//
// - A subnet ID, which determines the following:
//
// - The VPC in which Amazon EFS creates the mount target
//
// - The Availability Zone in which Amazon EFS creates the mount target
//
// - The IP address range from which Amazon EFS selects the IP address of the
// mount target (if you don't specify an IP address in the request)
//
// After creating the mount target, Amazon EFS returns a response that includes, a
// MountTargetId and an IpAddress . You use this IP address when mounting the file
// system in an EC2 instance. You can also use the mount target's DNS name when
// mounting the file system. The EC2 instance on which you mount the file system by
// using the mount target can resolve the mount target's DNS name to its IP
// address. For more information, see [How it Works: Implementation Overview].
//
// Note that you can create mount targets for a file system in only one VPC, and
// there can be only one mount target per Availability Zone. That is, if the file
// system already has one or more mount targets created for it, the subnet
// specified in the request to add another mount target must meet the following
// requirements:
//
// - Must belong to the same VPC as the subnets of the existing mount targets
//
// - Must not be in the same Availability Zone as any of the subnets of the
// existing mount targets
//
// If the request satisfies the requirements, Amazon EFS does the following:
//
// - Creates a new mount target in the specified subnet.
//
// - Also creates a new network interface in the subnet as follows:
//
// - If the request provides an IpAddress , Amazon EFS assigns that IP address to
// the network interface. Otherwise, Amazon EFS assigns a free address in the
// subnet (in the same way that the Amazon EC2 CreateNetworkInterface call does
// when a request does not specify a primary private IP address).
//
// - If the request provides SecurityGroups , this network interface is
// associated with those security groups. Otherwise, it belongs to the default
// security group for the subnet's VPC.
//
// - Assigns the description Mount target fsmt-id for file system fs-id where
// fsmt-id is the mount target ID, and fs-id is the FileSystemId .
//
// - Sets the requesterManaged property of the network interface to true , and
// the requesterId value to EFS .
//
// # Each Amazon EFS mount target has one corresponding requester-managed EC2
//
// network interface. After the network interface is created, Amazon EFS sets the
// NetworkInterfaceId field in the mount target's description to the network
// interface ID, and the IpAddress field to its address. If network interface
// creation fails, the entire CreateMountTarget operation fails.
//
// The CreateMountTarget call returns only after creating the network interface,
// but while the mount target state is still creating , you can check the mount
// target creation status by calling the DescribeMountTargetsoperation, which among other things
// returns the mount target state.
//
// We recommend that you create a mount target in each of the Availability Zones.
// There are cost considerations for using a file system in an Availability Zone
// through a mount target created in another Availability Zone. For more
// information, see [Amazon EFS pricing]. In addition, by always using a mount target local to the
// instance's Availability Zone, you eliminate a partial failure scenario. If the
// Availability Zone in which your mount target is created goes down, then you
// can't access your file system through that mount target.
//
// This operation requires permissions for the following action on the file system:
//
// - elasticfilesystem:CreateMountTarget
//
// This operation also requires permissions for the following Amazon EC2 actions:
//
// - ec2:DescribeSubnets
//
// - ec2:DescribeNetworkInterfaces
//
// - ec2:CreateNetworkInterface
//
// [Amazon EFS: How it Works]: https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html
// [Amazon EFS pricing]: http://aws.amazon.com/efs/pricing/
// [How it Works: Implementation Overview]: https://docs.aws.amazon.com/efs/latest/ug/how-it-works.html#how-it-works-implementation
func efs_CreateMountTarget(cfg aws.Config, client *efs.Client) {
	input := &efs.CreateMountTargetInput{
		// FileSystemId: *string, // Required
		// SubnetId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsSubnetId) > 0 {
		input.SubnetId = aws.String(_efsSubnetId)
	}
	if len(_efsIpAddress) > 0 {
		input.IpAddress = aws.String(_efsIpAddress)
	}
	if len(_efsIpAddressType) > 0 {
		if err := assignInputField(input, "IpAddressType", _efsIpAddressType); err != nil {
			log.Errorf("invalid --ip-address-type: %s", err.Error())
			return
		}
	}
	if len(_efsIpv6Address) > 0 {
		input.Ipv6Address = aws.String(_efsIpv6Address)
	}
	if len(_efsSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _efsSecurityGroups...)
	}

	if resp, err := client.CreateMountTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a replication conﬁguration to either a new or existing EFS file system.
// For more information, see [Amazon EFS replication]in the Amazon EFS User Guide. The replication
// configuration specifies the following:
//
// - Source file system – The EFS file system that you want to replicate.
//
// - Destination file system – The destination file system to which the source
// file system is replicated. There can only be one destination file system in a
// replication configuration.
//
// A file system can be part of only one replication configuration.
//
// # The destination parameters for the replication configuration depend on whether
//
// you are replicating to a new file system or to an existing file system, and if
// you are replicating across Amazon Web Services accounts. See DestinationToCreatefor more
// information.
//
// This operation requires permissions for the
// elasticfilesystem:CreateReplicationConfiguration action. Additionally, other
// permissions are required depending on how you are replicating file systems. For
// more information, see [Required permissions for replication]in the Amazon EFS User Guide.
//
// [Required permissions for replication]: https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html#efs-replication-permissions
// [Amazon EFS replication]: https://docs.aws.amazon.com/efs/latest/ug/efs-replication.html
func efs_CreateReplicationConfiguration(cfg aws.Config, client *efs.Client) {
	input := &efs.CreateReplicationConfigurationInput{
		// Destinations: []types.DestinationToCreate, // Required
		// SourceFileSystemId: *string, // Required
	}

	if len(_efsDestinations) > 0 {
		if err := assignInputField(input, "Destinations", _efsDestinations); err != nil {
			log.Errorf("invalid --destinations: %s", err.Error())
			return
		}
	}
	if len(_efsSourceFileSystemId) > 0 {
		input.SourceFileSystemId = aws.String(_efsSourceFileSystemId)
	}

	if resp, err := client.CreateReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// DEPRECATED - CreateTags is deprecated and not maintained. To create tags for
// EFS resources, use the API action.
//
// Creates or overwrites tags associated with a file system. Each tag is a
// key-value pair. If a tag key specified in the request already exists on the file
// system, this operation overwrites its value with the value provided in the
// request. If you add the Name tag to your file system, Amazon EFS returns it in
// the response to the DescribeFileSystemsoperation.
//
// This operation requires permission for the elasticfilesystem:CreateTags action.
//
// Deprecated: Use TagResource.
func efs_CreateTags(cfg aws.Config, client *efs.Client) {
	input := &efs.CreateTagsInput{
		// FileSystemId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsTags) > 0 {
		if err := assignInputField(input, "Tags", _efsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified access point. After deletion is complete, new clients can
// no longer connect to the access points. Clients connected to the access point at
// the time of deletion will continue to function until they terminate their
// connection.
//
// This operation requires permissions for the elasticfilesystem:DeleteAccessPoint
// action.
func efs_DeleteAccessPoint(cfg aws.Config, client *efs.Client) {
	input := &efs.DeleteAccessPointInput{
		// AccessPointId: *string, // Required
	}

	if len(_efsAccessPointId) > 0 {
		input.AccessPointId = aws.String(_efsAccessPointId)
	}

	if resp, err := client.DeleteAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a file system, permanently severing access to its contents. Upon
// return, the file system no longer exists and you can't access any contents of
// the deleted file system.
//
// You need to manually delete mount targets attached to a file system before you
// can delete an EFS file system. This step is performed for you when you use the
// Amazon Web Services console to delete a file system.
//
// You cannot delete a file system that is part of an EFS replication
// configuration. You need to delete the replication configuration first.
//
// You can't delete a file system that is in use. That is, if the file system has
// any mount targets, you must first delete them. For more information, see DescribeMountTargetsand DeleteMountTarget.
//
// The DeleteFileSystem call returns while the file system state is still deleting
// . You can check the file system deletion status by calling the DescribeFileSystemsoperation, which
// returns a list of file systems in your account. If you pass file system ID or
// creation token for the deleted file system, the DescribeFileSystemsreturns a 404 FileSystemNotFound
// error.
//
// This operation requires permissions for the elasticfilesystem:DeleteFileSystem
// action.
func efs_DeleteFileSystem(cfg aws.Config, client *efs.Client) {
	input := &efs.DeleteFileSystemInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}

	if resp, err := client.DeleteFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the FileSystemPolicy for the specified file system. The default
// FileSystemPolicy goes into effect once the existing policy is deleted. For more
// information about the default file system policy, see [Using Resource-based Policies with EFS].
//
// This operation requires permissions for the
// elasticfilesystem:DeleteFileSystemPolicy action.
//
// [Using Resource-based Policies with EFS]: https://docs.aws.amazon.com/efs/latest/ug/res-based-policies-efs.html
func efs_DeleteFileSystemPolicy(cfg aws.Config, client *efs.Client) {
	input := &efs.DeleteFileSystemPolicyInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}

	if resp, err := client.DeleteFileSystemPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified mount target.
// This operation forcibly breaks any mounts of the file system by using the mount
// target that is being deleted, which might disrupt instances or applications
// using those mounts. To avoid applications getting cut off abruptly, you might
// consider unmounting any mounts of the mount target, if feasible. The operation
// also deletes the associated network interface. Uncommitted writes might be lost,
// but breaking a mount target using this operation does not corrupt the file
// system itself. The file system you created remains. You can mount an EC2
// instance in your VPC by using another mount target.
//
// This operation requires permissions for the following action on the file system:
//
// - elasticfilesystem:DeleteMountTarget
//
// The DeleteMountTarget call returns while the mount target state is still
// deleting . You can check the mount target deletion by calling the DescribeMountTargets operation,
// which returns a list of mount target descriptions for the given file system.
//
// The operation also requires permissions for the following Amazon EC2 action on
// the mount target's network interface:
//
// - ec2:DeleteNetworkInterface
func efs_DeleteMountTarget(cfg aws.Config, client *efs.Client) {
	input := &efs.DeleteMountTargetInput{
		// MountTargetId: *string, // Required
	}

	if len(_efsMountTargetId) > 0 {
		input.MountTargetId = aws.String(_efsMountTargetId)
	}

	if resp, err := client.DeleteMountTarget(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a replication configuration. Deleting a replication configuration ends
// the replication process. After a replication configuration is deleted, the
// destination file system becomes Writeable and its replication overwrite
// protection is re-enabled. For more information, see [Delete a replication configuration].
//
// This operation requires permissions for the
// elasticfilesystem:DeleteReplicationConfiguration action.
//
// [Delete a replication configuration]: https://docs.aws.amazon.com/efs/latest/ug/delete-replications.html
func efs_DeleteReplicationConfiguration(cfg aws.Config, client *efs.Client) {
	input := &efs.DeleteReplicationConfigurationInput{
		// SourceFileSystemId: *string, // Required
	}

	if len(_efsSourceFileSystemId) > 0 {
		input.SourceFileSystemId = aws.String(_efsSourceFileSystemId)
	}
	if len(_efsDeletionMode) > 0 {
		if err := assignInputField(input, "DeletionMode", _efsDeletionMode); err != nil {
			log.Errorf("invalid --deletion-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteReplicationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// DEPRECATED - DeleteTags is deprecated and not maintained. To remove tags from
// EFS resources, use the API action.
//
// Deletes the specified tags from a file system. If the DeleteTags request
// includes a tag key that doesn't exist, Amazon EFS ignores it and doesn't cause
// an error. For more information about tags and related restrictions, see [Tag restrictions]in the
// Billing and Cost Management User Guide.
//
// This operation requires permissions for the elasticfilesystem:DeleteTags action.
//
// Deprecated: Use UntagResource.
//
// [Tag restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
func efs_DeleteTags(cfg aws.Config, client *efs.Client) {
	input := &efs.DeleteTagsInput{
		// FileSystemId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _efsTagKeys...)
	}

	if resp, err := client.DeleteTags(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of a specific Amazon EFS access point if the
// AccessPointId is provided. If you provide an EFS FileSystemId , it returns
// descriptions of all access points for that file system. You can provide either
// an AccessPointId or a FileSystemId in the request, but not both.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeAccessPoints action.
func efs_DescribeAccessPoints(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeAccessPointsInput{}

	if len(_efsAccessPointId) > 0 {
		input.AccessPointId = aws.String(_efsAccessPointId)
	}
	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _efsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_efsNextToken) > 0 {
		input.NextToken = aws.String(_efsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeAccessPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*efs.DescribeAccessPointsOutput
	p := efs.NewDescribeAccessPointsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns the account preferences settings for the Amazon Web Services account
// associated with the user making the request, in the current Amazon Web Services
// Region.
func efs_DescribeAccountPreferences(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeAccountPreferencesInput{}

	if len(_efsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _efsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_efsNextToken) > 0 {
		input.NextToken = aws.String(_efsNextToken)
	}

	if resp, err := client.DescribeAccountPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the backup policy for the specified EFS file system.
func efs_DescribeBackupPolicy(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeBackupPolicyInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}

	if resp, err := client.DescribeBackupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the FileSystemPolicy for the specified EFS file system.
// This operation requires permissions for the
// elasticfilesystem:DescribeFileSystemPolicy action.
func efs_DescribeFileSystemPolicy(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeFileSystemPolicyInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}

	if resp, err := client.DescribeFileSystemPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the description of a specific Amazon EFS file system if either the file
// system CreationToken or the FileSystemId is provided. Otherwise, it returns
// descriptions of all file systems owned by the caller's Amazon Web Services
// account in the Amazon Web Services Region of the endpoint that you're calling.
//
// When retrieving all file system descriptions, you can optionally specify the
// MaxItems parameter to limit the number of descriptions in a response. This
// number is automatically set to 100. If more file system descriptions remain,
// Amazon EFS returns a NextMarker , an opaque token, in the response. In this
// case, you should send a subsequent request with the Marker request parameter
// set to the value of NextMarker .
//
// To retrieve a list of your file system descriptions, this operation is used in
// an iterative process, where DescribeFileSystems is called first without the
// Marker and then the operation continues to call it with the Marker parameter
// set to the value of the NextMarker from the previous response until the
// response has no NextMarker .
//
// The order of file systems returned in the response of one DescribeFileSystems
// call and the order of file systems returned across the responses of a multi-call
// iteration is unspecified.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeFileSystems action.
func efs_DescribeFileSystems(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeFileSystemsInput{}

	if len(_efsCreationToken) > 0 {
		input.CreationToken = aws.String(_efsCreationToken)
	}
	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsMarker) > 0 {
		input.Marker = aws.String(_efsMarker)
	}
	if len(_efsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _efsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeFileSystems(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*efs.DescribeFileSystemsOutput
	p := efs.NewDescribeFileSystemsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Returns the current LifecycleConfiguration object for the specified EFS file
// system. Lifecycle management uses the LifecycleConfiguration object to identify
// when to move files between storage classes. For a file system without a
// LifecycleConfiguration object, the call returns an empty array in the response.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeLifecycleConfiguration operation.
func efs_DescribeLifecycleConfiguration(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeLifecycleConfigurationInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}

	if resp, err := client.DescribeLifecycleConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the security groups currently in effect for a mount target. This
// operation requires that the network interface of the mount target has been
// created and the lifecycle state of the mount target is not deleted .
//
// This operation requires permissions for the following actions:
//
// - elasticfilesystem:DescribeMountTargetSecurityGroups action on the mount
// target's file system.
//
// - ec2:DescribeNetworkInterfaceAttribute action on the mount target's network
// interface.
func efs_DescribeMountTargetSecurityGroups(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeMountTargetSecurityGroupsInput{
		// MountTargetId: *string, // Required
	}

	if len(_efsMountTargetId) > 0 {
		input.MountTargetId = aws.String(_efsMountTargetId)
	}

	if resp, err := client.DescribeMountTargetSecurityGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the descriptions of all the current mount targets, or a specific mount
// target, for a file system. When requesting all of the current mount targets, the
// order of mount targets returned in the response is unspecified.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeMountTargets action, on either the file system ID that
// you specify in FileSystemId , or on the file system of the mount target that you
// specify in MountTargetId .
func efs_DescribeMountTargets(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeMountTargetsInput{}

	if len(_efsAccessPointId) > 0 {
		input.AccessPointId = aws.String(_efsAccessPointId)
	}
	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsMarker) > 0 {
		input.Marker = aws.String(_efsMarker)
	}
	if len(_efsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _efsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}
	if len(_efsMountTargetId) > 0 {
		input.MountTargetId = aws.String(_efsMountTargetId)
	}

	if disablePaginator() {
		if resp, err := client.DescribeMountTargets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*efs.DescribeMountTargetsOutput
	p := efs.NewDescribeMountTargetsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Retrieves the replication configuration for a specific file system. If a file
// system is not specified, all of the replication configurations for the Amazon
// Web Services account in an Amazon Web Services Region are retrieved.
func efs_DescribeReplicationConfigurations(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeReplicationConfigurationsInput{}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _efsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_efsNextToken) > 0 {
		input.NextToken = aws.String(_efsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.DescribeReplicationConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*efs.DescribeReplicationConfigurationsOutput
	p := efs.NewDescribeReplicationConfigurationsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// DEPRECATED - The DescribeTags action is deprecated and not maintained. To view
// tags associated with EFS resources, use the ListTagsForResource API action.
//
// Returns the tags associated with a file system. The order of tags returned in
// the response of one DescribeTags call and the order of tags returned across the
// responses of a multiple-call iteration (when using pagination) is unspecified.
//
// This operation requires permissions for the elasticfilesystem:DescribeTags
// action.
//
// Deprecated: Use ListTagsForResource.
func efs_DescribeTags(cfg aws.Config, client *efs.Client) {
	input := &efs.DescribeTagsInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsMarker) > 0 {
		input.Marker = aws.String(_efsMarker)
	}
	if len(_efsMaxItems) > 0 {
		if err := assignInputField(input, "MaxItems", _efsMaxItems); err != nil {
			log.Errorf("invalid --max-items: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.DescribeTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*efs.DescribeTagsOutput
	p := efs.NewDescribeTagsPaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Lists all tags for a top-level EFS resource. You must provide the ID of the
// resource that you want to retrieve the tags for.
//
// This operation requires permissions for the
// elasticfilesystem:DescribeAccessPoints action.
func efs_ListTagsForResource(cfg aws.Config, client *efs.Client) {
	input := &efs.ListTagsForResourceInput{
		// ResourceId: *string, // Required
	}

	if len(_efsResourceId) > 0 {
		input.ResourceId = aws.String(_efsResourceId)
	}
	if len(_efsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _efsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_efsNextToken) > 0 {
		input.NextToken = aws.String(_efsNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*efs.ListTagsForResourceOutput
	p := efs.NewListTagsForResourcePaginator(client, input)
	for p.HasMorePages() {
		if resp, err := p.NextPage(context.TODO()); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			results = append(results, resp)
		}
	}
	writeOutput(nil, nil, results, _awsOutput)
}

// Modifies the set of security groups in effect for a mount target.
// When you create a mount target, Amazon EFS also creates a new network
// interface. For more information, see CreateMountTarget. This operation replaces the security
// groups in effect for the network interface associated with a mount target, with
// the SecurityGroups provided in the request. This operation requires that the
// network interface of the mount target has been created and the lifecycle state
// of the mount target is not deleted .
//
// The operation requires permissions for the following actions:
//
// - elasticfilesystem:ModifyMountTargetSecurityGroups action on the mount
// target's file system.
//
// - ec2:ModifyNetworkInterfaceAttribute action on the mount target's network
// interface.
func efs_ModifyMountTargetSecurityGroups(cfg aws.Config, client *efs.Client) {
	input := &efs.ModifyMountTargetSecurityGroupsInput{
		// MountTargetId: *string, // Required
	}

	if len(_efsMountTargetId) > 0 {
		input.MountTargetId = aws.String(_efsMountTargetId)
	}
	if len(_efsSecurityGroups) > 0 {
		input.SecurityGroups = append([]string(nil), _efsSecurityGroups...)
	}

	if resp, err := client.ModifyMountTargetSecurityGroups(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this operation to set the account preference in the current Amazon Web
// Services Region to use long 17 character (63 bit) or short 8 character (32 bit)
// resource IDs for new EFS file system and mount target resources. All existing
// resource IDs are not affected by any changes you make. You can set the ID
// preference during the opt-in period as EFS transitions to long resource IDs. For
// more information, see [Managing Amazon EFS resource IDs].
//
// Starting in October, 2021, you will receive an error if you try to set the
// account preference to use the short 8 character format resource ID. Contact
// Amazon Web Services support if you receive an error and must use short IDs for
// file system and mount target resources.
//
// [Managing Amazon EFS resource IDs]: https://docs.aws.amazon.com/efs/latest/ug/manage-efs-resource-ids.html
func efs_PutAccountPreferences(cfg aws.Config, client *efs.Client) {
	input := &efs.PutAccountPreferencesInput{
		// ResourceIdType: types.ResourceIdType, // Required
	}

	if len(_efsResourceIdType) > 0 {
		if err := assignInputField(input, "ResourceIdType", _efsResourceIdType); err != nil {
			log.Errorf("invalid --resource-id-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccountPreferences(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the file system's backup policy. Use this action to start or stop
// automatic backups of the file system.
func efs_PutBackupPolicy(cfg aws.Config, client *efs.Client) {
	input := &efs.PutBackupPolicyInput{
		// BackupPolicy: *types.BackupPolicy, // Required
		// FileSystemId: *string, // Required
	}

	if len(_efsBackupPolicy) > 0 {
		if err := assignInputField(input, "BackupPolicy", _efsBackupPolicy); err != nil {
			log.Errorf("invalid --backup-policy: %s", err.Error())
			return
		}
	}
	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}

	if resp, err := client.PutBackupPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies an Amazon EFS FileSystemPolicy to an Amazon EFS file system. A file
// system policy is an IAM resource-based policy and can contain multiple policy
// statements. A file system always has exactly one file system policy, which can
// be the default policy or an explicit policy set or updated using this API
// operation. EFS file system policies have a 20,000 character limit. When an
// explicit policy is set, it overrides the default policy. For more information
// about the default file system policy, see [Default EFS file system policy].
//
// EFS file system policies have a 20,000 character limit.
//
// This operation requires permissions for the
// elasticfilesystem:PutFileSystemPolicy action.
//
// [Default EFS file system policy]: https://docs.aws.amazon.com/efs/latest/ug/iam-access-control-nfs-efs.html#default-filesystempolicy
func efs_PutFileSystemPolicy(cfg aws.Config, client *efs.Client) {
	input := &efs.PutFileSystemPolicyInput{
		// FileSystemId: *string, // Required
		// Policy: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsPolicy) > 0 {
		input.Policy = aws.String(_efsPolicy)
	}
	if len(_efsBypassPolicyLockoutSafetyCheck) > 0 {
		if err := assignInputField(input, "BypassPolicyLockoutSafetyCheck", _efsBypassPolicyLockoutSafetyCheck); err != nil {
			log.Errorf("invalid --bypass-policy-lockout-safety-check: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutFileSystemPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Use this action to manage storage for your file system. A LifecycleConfiguration
// consists of one or more LifecyclePolicy objects that define the following:
//
// - TransitionToIA – When to move files in the file system from primary storage
// (Standard storage class) into the Infrequent Access (IA) storage.
//
// - TransitionToArchive – When to move files in the file system from their
// current storage class (either IA or Standard storage) into the Archive storage.
//
// # File systems cannot transition into Archive storage before transitioning into
//
// IA storage. Therefore, TransitionToArchive must either not be set or must be
// later than TransitionToIA.
//
// # The Archive storage class is available only for file systems that use the
//
// Elastic throughput mode and the General Purpose performance mode.
//
// - TransitionToPrimaryStorageClass – Whether to move files in the file system
// back to primary storage (Standard storage class) after they are accessed in IA
// or Archive storage.
//
// For more information, see [Managing file system storage].
//
// Each Amazon EFS file system supports one lifecycle configuration, which applies
// to all files in the file system. If a LifecycleConfiguration object already
// exists for the specified file system, a PutLifecycleConfiguration call modifies
// the existing configuration. A PutLifecycleConfiguration call with an empty
// LifecyclePolicies array in the request body deletes any existing
// LifecycleConfiguration . In the request, specify the following:
//
// - The ID for the file system for which you are enabling, disabling, or
// modifying lifecycle management.
//
// - A LifecyclePolicies array of LifecyclePolicy objects that define when to
// move files to IA storage, to Archive storage, and back to primary storage.
//
// # Amazon EFS requires that each LifecyclePolicy object have only have a single
//
// transition, so the LifecyclePolicies array needs to be structured with
// separate LifecyclePolicy objects. See the example requests in the following
// section for more information.
//
// This operation requires permissions for the
// elasticfilesystem:PutLifecycleConfiguration operation.
//
// To apply a LifecycleConfiguration object to an encrypted file system, you need
// the same Key Management Service permissions as when you created the encrypted
// file system.
//
// [Managing file system storage]: https://docs.aws.amazon.com/efs/latest/ug/lifecycle-management-efs.html
func efs_PutLifecycleConfiguration(cfg aws.Config, client *efs.Client) {
	input := &efs.PutLifecycleConfigurationInput{
		// FileSystemId: *string, // Required
		// LifecyclePolicies: []types.LifecyclePolicy, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsLifecyclePolicies) > 0 {
		if err := assignInputField(input, "LifecyclePolicies", _efsLifecyclePolicies); err != nil {
			log.Errorf("invalid --lifecycle-policies: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutLifecycleConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a tag for an EFS resource. You can create tags for EFS file systems and
// access points using this API operation.
//
// This operation requires permissions for the elasticfilesystem:TagResource
// action.
func efs_TagResource(cfg aws.Config, client *efs.Client) {
	input := &efs.TagResourceInput{
		// ResourceId: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_efsResourceId) > 0 {
		input.ResourceId = aws.String(_efsResourceId)
	}
	if len(_efsTags) > 0 {
		if err := assignInputField(input, "Tags", _efsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes tags from an EFS resource. You can remove tags from EFS file systems
// and access points using this API operation.
//
// This operation requires permissions for the elasticfilesystem:UntagResource
// action.
func efs_UntagResource(cfg aws.Config, client *efs.Client) {
	input := &efs.UntagResourceInput{
		// ResourceId: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_efsResourceId) > 0 {
		input.ResourceId = aws.String(_efsResourceId)
	}
	if len(_efsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _efsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the throughput mode or the amount of provisioned throughput of an
// existing file system.
func efs_UpdateFileSystem(cfg aws.Config, client *efs.Client) {
	input := &efs.UpdateFileSystemInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsProvisionedThroughputInMibps) > 0 {
		if err := assignInputField(input, "ProvisionedThroughputInMibps", _efsProvisionedThroughputInMibps); err != nil {
			log.Errorf("invalid --provisioned-throughput-in-mibps: %s", err.Error())
			return
		}
	}
	if len(_efsThroughputMode) > 0 {
		if err := assignInputField(input, "ThroughputMode", _efsThroughputMode); err != nil {
			log.Errorf("invalid --throughput-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFileSystem(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates protection on the file system.
// This operation requires permissions for the
// elasticfilesystem:UpdateFileSystemProtection action.
func efs_UpdateFileSystemProtection(cfg aws.Config, client *efs.Client) {
	input := &efs.UpdateFileSystemProtectionInput{
		// FileSystemId: *string, // Required
	}

	if len(_efsFileSystemId) > 0 {
		input.FileSystemId = aws.String(_efsFileSystemId)
	}
	if len(_efsReplicationOverwriteProtection) > 0 {
		if err := assignInputField(input, "ReplicationOverwriteProtection", _efsReplicationOverwriteProtection); err != nil {
			log.Errorf("invalid --replication-overwrite-protection: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateFileSystemProtection(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_efsCmd)
	_efsCmd.Flags().SortFlags = false

	_efsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_efsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_efsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_efsCmd.Flags().StringVarP(&_efsAccessPointId, "access-point-id", "", "", "Access Point ID")
	_efsCmd.Flags().StringVarP(&_efsAvailabilityZoneName, "availability-zone-name", "", "", "Availability Zone Name")
	_efsCmd.Flags().StringVarP(&_efsBackup, "backup", "", "", "Backup")
	_efsCmd.Flags().StringVarP(&_efsBackupPolicy, "backup-policy", "", "", "Backup Policy")
	_efsCmd.Flags().StringVarP(&_efsBypassPolicyLockoutSafetyCheck, "bypass-policy-lockout-safety-check", "", "", "Bypass Policy Lockout Safety Check")
	_efsCmd.Flags().StringVarP(&_efsClientToken, "client-token", "", "", "Client Token")
	_efsCmd.Flags().StringVarP(&_efsCreationToken, "creation-token", "", "", "Creation Token")
	_efsCmd.Flags().StringVarP(&_efsDeletionMode, "deletion-mode", "", "", "Deletion Mode")
	_efsCmd.Flags().StringVarP(&_efsDestinations, "destinations", "", "", "Destinations")
	_efsCmd.Flags().StringVarP(&_efsEncrypted, "encrypted", "", "", "Encrypted")
	_efsCmd.Flags().StringVarP(&_efsFileSystemId, "file-system-id", "", "", "File System ID")
	_efsCmd.Flags().StringVarP(&_efsIpAddress, "ip-address", "", "", "IP Address")
	_efsCmd.Flags().StringVarP(&_efsIpAddressType, "ip-address-type", "", "", "IP Address Type")
	_efsCmd.Flags().StringVarP(&_efsIpv6Address, "ipv6-address", "", "", "IPV6 Address")
	_efsCmd.Flags().StringVarP(&_efsKmsKeyId, "kms-key-id", "", "", "KMS Key ID")
	_efsCmd.Flags().StringVarP(&_efsLifecyclePolicies, "lifecycle-policies", "", "", "Lifecycle Policies")
	_efsCmd.Flags().StringVarP(&_efsMarker, "marker", "", "", "Marker")
	_efsCmd.Flags().StringVarP(&_efsMaxItems, "max-items", "", "", "Max Items")
	_efsCmd.Flags().StringVarP(&_efsMaxResults, "max-results", "", "", "Max Results")
	_efsCmd.Flags().StringVarP(&_efsMountTargetId, "mount-target-id", "", "", "Mount Target ID")
	_efsCmd.Flags().StringVarP(&_efsNextToken, "next-token", "", "", "Next Token")
	_efsCmd.Flags().StringVarP(&_efsPerformanceMode, "performance-mode", "", "", "Performance Mode")
	_efsCmd.Flags().StringVarP(&_efsPolicy, "policy", "", "", "Policy")
	_efsCmd.Flags().StringVarP(&_efsPosixUser, "posix-user", "", "", "Posix User")
	_efsCmd.Flags().StringVarP(&_efsProvisionedThroughputInMibps, "provisioned-throughput-in-mibps", "", "", "Provisioned Throughput In Mibps")
	_efsCmd.Flags().StringVarP(&_efsReplicationOverwriteProtection, "replication-overwrite-protection", "", "", "Replication Overwrite Protection")
	_efsCmd.Flags().StringVarP(&_efsResourceId, "resource-id", "", "", "Resource ID")
	_efsCmd.Flags().StringVarP(&_efsResourceIdType, "resource-id-type", "", "", "Resource ID Type")
	_efsCmd.Flags().StringVarP(&_efsRootDirectory, "root-directory", "", "", "Root Directory")
	_efsCmd.Flags().StringSliceVarP(&_efsSecurityGroups, "security-groups", "", nil, "Security Groups")
	_efsCmd.Flags().StringVarP(&_efsSourceFileSystemId, "source-file-system-id", "", "", "Source File System ID")
	_efsCmd.Flags().StringVarP(&_efsSubnetId, "subnet-id", "", "", "Subnet ID")
	_efsCmd.Flags().StringSliceVarP(&_efsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_efsCmd.Flags().StringVarP(&_efsTags, "tags", "", "", "Tags")
	_efsCmd.Flags().StringVarP(&_efsThroughputMode, "throughput-mode", "", "", "Throughput Mode")

	_efsCmd.Flags().BoolVarP(&_efsCreateAccessPoint, "create-access-point", "", false, "Create Access Point")
	_efsCmd.Flags().BoolVarP(&_efsCreateFileSystem, "create-file-system", "", false, "Create File System")
	_efsCmd.Flags().BoolVarP(&_efsCreateMountTarget, "create-mount-target", "", false, "Create Mount Target")
	_efsCmd.Flags().BoolVarP(&_efsCreateReplicationConfiguration, "create-replication-configuration", "", false, "Create Replication Configuration")
	_efsCmd.Flags().BoolVarP(&_efsCreateTags, "create-tags", "", false, "Create Tags")
	_efsCmd.Flags().BoolVarP(&_efsDeleteAccessPoint, "delete-access-point", "", false, "Delete Access Point")
	_efsCmd.Flags().BoolVarP(&_efsDeleteFileSystem, "delete-file-system", "", false, "Delete File System")
	_efsCmd.Flags().BoolVarP(&_efsDeleteFileSystemPolicy, "delete-file-system-policy", "", false, "Delete File System Policy")
	_efsCmd.Flags().BoolVarP(&_efsDeleteMountTarget, "delete-mount-target", "", false, "Delete Mount Target")
	_efsCmd.Flags().BoolVarP(&_efsDeleteReplicationConfiguration, "delete-replication-configuration", "", false, "Delete Replication Configuration")
	_efsCmd.Flags().BoolVarP(&_efsDeleteTags, "delete-tags", "", false, "Delete Tags")
	_efsCmd.Flags().BoolVarP(&_efsDescribeAccessPoints, "describe-access-points", "", false, "Describe Access Points")
	_efsCmd.Flags().BoolVarP(&_efsDescribeAccountPreferences, "describe-account-preferences", "", false, "Describe Account Preferences")
	_efsCmd.Flags().BoolVarP(&_efsDescribeBackupPolicy, "describe-backup-policy", "", false, "Describe Backup Policy")
	_efsCmd.Flags().BoolVarP(&_efsDescribeFileSystemPolicy, "describe-file-system-policy", "", false, "Describe File System Policy")
	_efsCmd.Flags().BoolVarP(&_efsDescribeFileSystems, "describe-file-systems", "", false, "Describe File Systems")
	_efsCmd.Flags().BoolVarP(&_efsDescribeLifecycleConfiguration, "describe-lifecycle-configuration", "", false, "Describe Lifecycle Configuration")
	_efsCmd.Flags().BoolVarP(&_efsDescribeMountTargetSecurityGroups, "describe-mount-target-security-groups", "", false, "Describe Mount Target Security Groups")
	_efsCmd.Flags().BoolVarP(&_efsDescribeMountTargets, "describe-mount-targets", "", false, "Describe Mount Targets")
	_efsCmd.Flags().BoolVarP(&_efsDescribeReplicationConfigurations, "describe-replication-configurations", "", false, "Describe Replication Configurations")
	_efsCmd.Flags().BoolVarP(&_efsDescribeTags, "describe-tags", "", false, "Describe Tags")
	_efsCmd.Flags().BoolVarP(&_efsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_efsCmd.Flags().BoolVarP(&_efsModifyMountTargetSecurityGroups, "modify-mount-target-security-groups", "", false, "Modify Mount Target Security Groups")
	_efsCmd.Flags().BoolVarP(&_efsPutAccountPreferences, "put-account-preferences", "", false, "Put Account Preferences")
	_efsCmd.Flags().BoolVarP(&_efsPutBackupPolicy, "put-backup-policy", "", false, "Put Backup Policy")
	_efsCmd.Flags().BoolVarP(&_efsPutFileSystemPolicy, "put-file-system-policy", "", false, "Put File System Policy")
	_efsCmd.Flags().BoolVarP(&_efsPutLifecycleConfiguration, "put-lifecycle-configuration", "", false, "Put Lifecycle Configuration")
	_efsCmd.Flags().BoolVarP(&_efsTagResource, "tag-resource", "", false, "Tag Resource")
	_efsCmd.Flags().BoolVarP(&_efsUntagResource, "untag-resource", "", false, "Untag Resource")
	_efsCmd.Flags().BoolVarP(&_efsUpdateFileSystem, "update-file-system", "", false, "Update File System")
	_efsCmd.Flags().BoolVarP(&_efsUpdateFileSystemProtection, "update-file-system-protection", "", false, "Update File System Protection")

}
