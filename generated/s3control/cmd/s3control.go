package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3control"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// s3controlCmd represents the s3control command
var _s3controlCmd = &cobra.Command{
	Use:   "s3control",
	Short: "AWS s3control CLI",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := LoadAWSConfigWithMiddleware(_awsProfile)
		if err != nil {
			log.Errorf("Failed to load configuration: %s", err.Error())
			return
		}
		if len(_awsRegion) > 0 {
			cfg.Region = _awsRegion
		}
		client := s3control.NewFromConfig(cfg)
		if _s3controlAssociateAccessGrantsIdentityCenter {
			s3control_AssociateAccessGrantsIdentityCenter(cfg, client)
			return
		}
		if _s3controlCreateAccessGrant {
			s3control_CreateAccessGrant(cfg, client)
			return
		}
		if _s3controlCreateAccessGrantsInstance {
			s3control_CreateAccessGrantsInstance(cfg, client)
			return
		}
		if _s3controlCreateAccessGrantsLocation {
			s3control_CreateAccessGrantsLocation(cfg, client)
			return
		}
		if _s3controlCreateAccessPoint {
			s3control_CreateAccessPoint(cfg, client)
			return
		}
		if _s3controlCreateAccessPointForObjectLambda {
			s3control_CreateAccessPointForObjectLambda(cfg, client)
			return
		}
		if _s3controlCreateBucket {
			s3control_CreateBucket(cfg, client)
			return
		}
		if _s3controlCreateJob {
			s3control_CreateJob(cfg, client)
			return
		}
		if _s3controlCreateMultiRegionAccessPoint {
			s3control_CreateMultiRegionAccessPoint(cfg, client)
			return
		}
		if _s3controlCreateStorageLensGroup {
			s3control_CreateStorageLensGroup(cfg, client)
			return
		}
		if _s3controlDeleteAccessGrant {
			s3control_DeleteAccessGrant(cfg, client)
			return
		}
		if _s3controlDeleteAccessGrantsInstance {
			s3control_DeleteAccessGrantsInstance(cfg, client)
			return
		}
		if _s3controlDeleteAccessGrantsInstanceResourcePolicy {
			s3control_DeleteAccessGrantsInstanceResourcePolicy(cfg, client)
			return
		}
		if _s3controlDeleteAccessGrantsLocation {
			s3control_DeleteAccessGrantsLocation(cfg, client)
			return
		}
		if _s3controlDeleteAccessPoint {
			s3control_DeleteAccessPoint(cfg, client)
			return
		}
		if _s3controlDeleteAccessPointForObjectLambda {
			s3control_DeleteAccessPointForObjectLambda(cfg, client)
			return
		}
		if _s3controlDeleteAccessPointPolicy {
			s3control_DeleteAccessPointPolicy(cfg, client)
			return
		}
		if _s3controlDeleteAccessPointPolicyForObjectLambda {
			s3control_DeleteAccessPointPolicyForObjectLambda(cfg, client)
			return
		}
		if _s3controlDeleteAccessPointScope {
			s3control_DeleteAccessPointScope(cfg, client)
			return
		}
		if _s3controlDeleteBucket {
			s3control_DeleteBucket(cfg, client)
			return
		}
		if _s3controlDeleteBucketLifecycleConfiguration {
			s3control_DeleteBucketLifecycleConfiguration(cfg, client)
			return
		}
		if _s3controlDeleteBucketPolicy {
			s3control_DeleteBucketPolicy(cfg, client)
			return
		}
		if _s3controlDeleteBucketReplication {
			s3control_DeleteBucketReplication(cfg, client)
			return
		}
		if _s3controlDeleteBucketTagging {
			s3control_DeleteBucketTagging(cfg, client)
			return
		}
		if _s3controlDeleteJobTagging {
			s3control_DeleteJobTagging(cfg, client)
			return
		}
		if _s3controlDeleteMultiRegionAccessPoint {
			s3control_DeleteMultiRegionAccessPoint(cfg, client)
			return
		}
		if _s3controlDeletePublicAccessBlock {
			s3control_DeletePublicAccessBlock(cfg, client)
			return
		}
		if _s3controlDeleteStorageLensConfiguration {
			s3control_DeleteStorageLensConfiguration(cfg, client)
			return
		}
		if _s3controlDeleteStorageLensConfigurationTagging {
			s3control_DeleteStorageLensConfigurationTagging(cfg, client)
			return
		}
		if _s3controlDeleteStorageLensGroup {
			s3control_DeleteStorageLensGroup(cfg, client)
			return
		}
		if _s3controlDescribeJob {
			s3control_DescribeJob(cfg, client)
			return
		}
		if _s3controlDescribeMultiRegionAccessPointOperation {
			s3control_DescribeMultiRegionAccessPointOperation(cfg, client)
			return
		}
		if _s3controlDissociateAccessGrantsIdentityCenter {
			s3control_DissociateAccessGrantsIdentityCenter(cfg, client)
			return
		}
		if _s3controlGetAccessGrant {
			s3control_GetAccessGrant(cfg, client)
			return
		}
		if _s3controlGetAccessGrantsInstance {
			s3control_GetAccessGrantsInstance(cfg, client)
			return
		}
		if _s3controlGetAccessGrantsInstanceForPrefix {
			s3control_GetAccessGrantsInstanceForPrefix(cfg, client)
			return
		}
		if _s3controlGetAccessGrantsInstanceResourcePolicy {
			s3control_GetAccessGrantsInstanceResourcePolicy(cfg, client)
			return
		}
		if _s3controlGetAccessGrantsLocation {
			s3control_GetAccessGrantsLocation(cfg, client)
			return
		}
		if _s3controlGetAccessPoint {
			s3control_GetAccessPoint(cfg, client)
			return
		}
		if _s3controlGetAccessPointConfigurationForObjectLambda {
			s3control_GetAccessPointConfigurationForObjectLambda(cfg, client)
			return
		}
		if _s3controlGetAccessPointForObjectLambda {
			s3control_GetAccessPointForObjectLambda(cfg, client)
			return
		}
		if _s3controlGetAccessPointPolicy {
			s3control_GetAccessPointPolicy(cfg, client)
			return
		}
		if _s3controlGetAccessPointPolicyForObjectLambda {
			s3control_GetAccessPointPolicyForObjectLambda(cfg, client)
			return
		}
		if _s3controlGetAccessPointPolicyStatus {
			s3control_GetAccessPointPolicyStatus(cfg, client)
			return
		}
		if _s3controlGetAccessPointPolicyStatusForObjectLambda {
			s3control_GetAccessPointPolicyStatusForObjectLambda(cfg, client)
			return
		}
		if _s3controlGetAccessPointScope {
			s3control_GetAccessPointScope(cfg, client)
			return
		}
		if _s3controlGetBucket {
			s3control_GetBucket(cfg, client)
			return
		}
		if _s3controlGetBucketLifecycleConfiguration {
			s3control_GetBucketLifecycleConfiguration(cfg, client)
			return
		}
		if _s3controlGetBucketPolicy {
			s3control_GetBucketPolicy(cfg, client)
			return
		}
		if _s3controlGetBucketReplication {
			s3control_GetBucketReplication(cfg, client)
			return
		}
		if _s3controlGetBucketTagging {
			s3control_GetBucketTagging(cfg, client)
			return
		}
		if _s3controlGetBucketVersioning {
			s3control_GetBucketVersioning(cfg, client)
			return
		}
		if _s3controlGetDataAccess {
			s3control_GetDataAccess(cfg, client)
			return
		}
		if _s3controlGetJobTagging {
			s3control_GetJobTagging(cfg, client)
			return
		}
		if _s3controlGetMultiRegionAccessPoint {
			s3control_GetMultiRegionAccessPoint(cfg, client)
			return
		}
		if _s3controlGetMultiRegionAccessPointPolicy {
			s3control_GetMultiRegionAccessPointPolicy(cfg, client)
			return
		}
		if _s3controlGetMultiRegionAccessPointPolicyStatus {
			s3control_GetMultiRegionAccessPointPolicyStatus(cfg, client)
			return
		}
		if _s3controlGetMultiRegionAccessPointRoutes {
			s3control_GetMultiRegionAccessPointRoutes(cfg, client)
			return
		}
		if _s3controlGetPublicAccessBlock {
			s3control_GetPublicAccessBlock(cfg, client)
			return
		}
		if _s3controlGetStorageLensConfiguration {
			s3control_GetStorageLensConfiguration(cfg, client)
			return
		}
		if _s3controlGetStorageLensConfigurationTagging {
			s3control_GetStorageLensConfigurationTagging(cfg, client)
			return
		}
		if _s3controlGetStorageLensGroup {
			s3control_GetStorageLensGroup(cfg, client)
			return
		}
		if _s3controlListAccessGrants {
			s3control_ListAccessGrants(cfg, client)
			return
		}
		if _s3controlListAccessGrantsInstances {
			s3control_ListAccessGrantsInstances(cfg, client)
			return
		}
		if _s3controlListAccessGrantsLocations {
			s3control_ListAccessGrantsLocations(cfg, client)
			return
		}
		if _s3controlListAccessPoints {
			s3control_ListAccessPoints(cfg, client)
			return
		}
		if _s3controlListAccessPointsForDirectoryBuckets {
			s3control_ListAccessPointsForDirectoryBuckets(cfg, client)
			return
		}
		if _s3controlListAccessPointsForObjectLambda {
			s3control_ListAccessPointsForObjectLambda(cfg, client)
			return
		}
		if _s3controlListCallerAccessGrants {
			s3control_ListCallerAccessGrants(cfg, client)
			return
		}
		if _s3controlListJobs {
			s3control_ListJobs(cfg, client)
			return
		}
		if _s3controlListMultiRegionAccessPoints {
			s3control_ListMultiRegionAccessPoints(cfg, client)
			return
		}
		if _s3controlListRegionalBuckets {
			s3control_ListRegionalBuckets(cfg, client)
			return
		}
		if _s3controlListStorageLensConfigurations {
			s3control_ListStorageLensConfigurations(cfg, client)
			return
		}
		if _s3controlListStorageLensGroups {
			s3control_ListStorageLensGroups(cfg, client)
			return
		}
		if _s3controlListTagsForResource {
			s3control_ListTagsForResource(cfg, client)
			return
		}
		if _s3controlPutAccessGrantsInstanceResourcePolicy {
			s3control_PutAccessGrantsInstanceResourcePolicy(cfg, client)
			return
		}
		if _s3controlPutAccessPointConfigurationForObjectLambda {
			s3control_PutAccessPointConfigurationForObjectLambda(cfg, client)
			return
		}
		if _s3controlPutAccessPointPolicy {
			s3control_PutAccessPointPolicy(cfg, client)
			return
		}
		if _s3controlPutAccessPointPolicyForObjectLambda {
			s3control_PutAccessPointPolicyForObjectLambda(cfg, client)
			return
		}
		if _s3controlPutAccessPointScope {
			s3control_PutAccessPointScope(cfg, client)
			return
		}
		if _s3controlPutBucketLifecycleConfiguration {
			s3control_PutBucketLifecycleConfiguration(cfg, client)
			return
		}
		if _s3controlPutBucketPolicy {
			s3control_PutBucketPolicy(cfg, client)
			return
		}
		if _s3controlPutBucketReplication {
			s3control_PutBucketReplication(cfg, client)
			return
		}
		if _s3controlPutBucketTagging {
			s3control_PutBucketTagging(cfg, client)
			return
		}
		if _s3controlPutBucketVersioning {
			s3control_PutBucketVersioning(cfg, client)
			return
		}
		if _s3controlPutJobTagging {
			s3control_PutJobTagging(cfg, client)
			return
		}
		if _s3controlPutMultiRegionAccessPointPolicy {
			s3control_PutMultiRegionAccessPointPolicy(cfg, client)
			return
		}
		if _s3controlPutPublicAccessBlock {
			s3control_PutPublicAccessBlock(cfg, client)
			return
		}
		if _s3controlPutStorageLensConfiguration {
			s3control_PutStorageLensConfiguration(cfg, client)
			return
		}
		if _s3controlPutStorageLensConfigurationTagging {
			s3control_PutStorageLensConfigurationTagging(cfg, client)
			return
		}
		if _s3controlSubmitMultiRegionAccessPointRoutes {
			s3control_SubmitMultiRegionAccessPointRoutes(cfg, client)
			return
		}
		if _s3controlTagResource {
			s3control_TagResource(cfg, client)
			return
		}
		if _s3controlUntagResource {
			s3control_UntagResource(cfg, client)
			return
		}
		if _s3controlUpdateAccessGrantsLocation {
			s3control_UpdateAccessGrantsLocation(cfg, client)
			return
		}
		if _s3controlUpdateJobPriority {
			s3control_UpdateJobPriority(cfg, client)
			return
		}
		if _s3controlUpdateJobStatus {
			s3control_UpdateJobStatus(cfg, client)
			return
		}
		if _s3controlUpdateStorageLensGroup {
			s3control_UpdateStorageLensGroup(cfg, client)
			return
		}

	},
}

var (
	_s3controlAssociateAccessGrantsIdentityCenter        bool
	_s3controlCreateAccessGrant                          bool
	_s3controlCreateAccessGrantsInstance                 bool
	_s3controlCreateAccessGrantsLocation                 bool
	_s3controlCreateAccessPoint                          bool
	_s3controlCreateAccessPointForObjectLambda           bool
	_s3controlCreateBucket                               bool
	_s3controlCreateJob                                  bool
	_s3controlCreateMultiRegionAccessPoint               bool
	_s3controlCreateStorageLensGroup                     bool
	_s3controlDeleteAccessGrant                          bool
	_s3controlDeleteAccessGrantsInstance                 bool
	_s3controlDeleteAccessGrantsInstanceResourcePolicy   bool
	_s3controlDeleteAccessGrantsLocation                 bool
	_s3controlDeleteAccessPoint                          bool
	_s3controlDeleteAccessPointForObjectLambda           bool
	_s3controlDeleteAccessPointPolicy                    bool
	_s3controlDeleteAccessPointPolicyForObjectLambda     bool
	_s3controlDeleteAccessPointScope                     bool
	_s3controlDeleteBucket                               bool
	_s3controlDeleteBucketLifecycleConfiguration         bool
	_s3controlDeleteBucketPolicy                         bool
	_s3controlDeleteBucketReplication                    bool
	_s3controlDeleteBucketTagging                        bool
	_s3controlDeleteJobTagging                           bool
	_s3controlDeleteMultiRegionAccessPoint               bool
	_s3controlDeletePublicAccessBlock                    bool
	_s3controlDeleteStorageLensConfiguration             bool
	_s3controlDeleteStorageLensConfigurationTagging      bool
	_s3controlDeleteStorageLensGroup                     bool
	_s3controlDescribeJob                                bool
	_s3controlDescribeMultiRegionAccessPointOperation    bool
	_s3controlDissociateAccessGrantsIdentityCenter       bool
	_s3controlGetAccessGrant                             bool
	_s3controlGetAccessGrantsInstance                    bool
	_s3controlGetAccessGrantsInstanceForPrefix           bool
	_s3controlGetAccessGrantsInstanceResourcePolicy      bool
	_s3controlGetAccessGrantsLocation                    bool
	_s3controlGetAccessPoint                             bool
	_s3controlGetAccessPointConfigurationForObjectLambda bool
	_s3controlGetAccessPointForObjectLambda              bool
	_s3controlGetAccessPointPolicy                       bool
	_s3controlGetAccessPointPolicyForObjectLambda        bool
	_s3controlGetAccessPointPolicyStatus                 bool
	_s3controlGetAccessPointPolicyStatusForObjectLambda  bool
	_s3controlGetAccessPointScope                        bool
	_s3controlGetBucket                                  bool
	_s3controlGetBucketLifecycleConfiguration            bool
	_s3controlGetBucketPolicy                            bool
	_s3controlGetBucketReplication                       bool
	_s3controlGetBucketTagging                           bool
	_s3controlGetBucketVersioning                        bool
	_s3controlGetDataAccess                              bool
	_s3controlGetJobTagging                              bool
	_s3controlGetMultiRegionAccessPoint                  bool
	_s3controlGetMultiRegionAccessPointPolicy            bool
	_s3controlGetMultiRegionAccessPointPolicyStatus      bool
	_s3controlGetMultiRegionAccessPointRoutes            bool
	_s3controlGetPublicAccessBlock                       bool
	_s3controlGetStorageLensConfiguration                bool
	_s3controlGetStorageLensConfigurationTagging         bool
	_s3controlGetStorageLensGroup                        bool
	_s3controlListAccessGrants                           bool
	_s3controlListAccessGrantsInstances                  bool
	_s3controlListAccessGrantsLocations                  bool
	_s3controlListAccessPoints                           bool
	_s3controlListAccessPointsForDirectoryBuckets        bool
	_s3controlListAccessPointsForObjectLambda            bool
	_s3controlListCallerAccessGrants                     bool
	_s3controlListJobs                                   bool
	_s3controlListMultiRegionAccessPoints                bool
	_s3controlListRegionalBuckets                        bool
	_s3controlListStorageLensConfigurations              bool
	_s3controlListStorageLensGroups                      bool
	_s3controlListTagsForResource                        bool
	_s3controlPutAccessGrantsInstanceResourcePolicy      bool
	_s3controlPutAccessPointConfigurationForObjectLambda bool
	_s3controlPutAccessPointPolicy                       bool
	_s3controlPutAccessPointPolicyForObjectLambda        bool
	_s3controlPutAccessPointScope                        bool
	_s3controlPutBucketLifecycleConfiguration            bool
	_s3controlPutBucketPolicy                            bool
	_s3controlPutBucketReplication                       bool
	_s3controlPutBucketTagging                           bool
	_s3controlPutBucketVersioning                        bool
	_s3controlPutJobTagging                              bool
	_s3controlPutMultiRegionAccessPointPolicy            bool
	_s3controlPutPublicAccessBlock                       bool
	_s3controlPutStorageLensConfiguration                bool
	_s3controlPutStorageLensConfigurationTagging         bool
	_s3controlSubmitMultiRegionAccessPointRoutes         bool
	_s3controlTagResource                                bool
	_s3controlUntagResource                              bool
	_s3controlUpdateAccessGrantsLocation                 bool
	_s3controlUpdateJobPriority                          bool
	_s3controlUpdateJobStatus                            bool
	_s3controlUpdateStorageLensGroup                     bool

	_s3controlAccessGrantId                     string
	_s3controlAccessGrantsLocationConfiguration string
	_s3controlAccessGrantsLocationId            string
	_s3controlAccountId                         string
	_s3controlACL                               string
	_s3controlAllowedByApplication              string
	_s3controlApplicationArn                    string
	_s3controlBucket                            string
	_s3controlBucketAccountId                   string
	_s3controlClientRequestToken                string
	_s3controlClientToken                       string
	_s3controlConfigId                          string
	_s3controlConfiguration                     string
	_s3controlConfirmRemoveSelfBucketAccess     string
	_s3controlConfirmationRequired              string
	_s3controlCreateBucketConfiguration         string
	_s3controlDataSourceId                      string
	_s3controlDataSourceType                    string
	_s3controlDescription                       string
	_s3controlDetails                           string
	_s3controlDirectoryBucket                   string
	_s3controlDurationSeconds                   string
	_s3controlGrantFullControl                  string
	_s3controlGrantRead                         string
	_s3controlGrantReadACP                      string
	_s3controlGrantScope                        string
	_s3controlGrantWrite                        string
	_s3controlGrantWriteACP                     string
	_s3controlGrantee                           string
	_s3controlGranteeIdentifier                 string
	_s3controlGranteeType                       string
	_s3controlIAMRoleArn                        string
	_s3controlIdentityCenterArn                 string
	_s3controlJobId                             string
	_s3controlJobStatuses                       string
	_s3controlLifecycleConfiguration            string
	_s3controlLocationScope                     string
	_s3controlManifest                          string
	_s3controlManifestGenerator                 string
	_s3controlMaxResults                        string
	_s3controlMFA                               string
	_s3controlMrap                              string
	_s3controlName                              string
	_s3controlNextToken                         string
	_s3controlObjectLockEnabledForBucket        string
	_s3controlOperation                         string
	_s3controlOrganization                      string
	_s3controlOutpostId                         string
	_s3controlPermission                        string
	_s3controlPolicy                            string
	_s3controlPriority                          string
	_s3controlPrivilege                         string
	_s3controlPublicAccessBlockConfiguration    string
	_s3controlReplicationConfiguration          string
	_s3controlReport                            string
	_s3controlRequestTokenARN                   string
	_s3controlRequestedJobStatus                string
	_s3controlResourceArn                       string
	_s3controlRoleArn                           string
	_s3controlRouteUpdates                      string
	_s3controlS3Prefix                          string
	_s3controlS3PrefixType                      string
	_s3controlScope                             string
	_s3controlStatusUpdateReason                string
	_s3controlStorageLensConfiguration          string
	_s3controlStorageLensGroup                  string
	_s3controlTagKeys                           []string
	_s3controlTagging                           string
	_s3controlTags                              string
	_s3controlTarget                            string
	_s3controlTargetType                        string
	_s3controlVersioningConfiguration           string
	_s3controlVpcConfiguration                  string
)

// Associate your S3 Access Grants instance with an Amazon Web Services IAM
// Identity Center instance. Use this action if you want to create access grants
// for users or groups from your corporate identity directory. First, you must add
// your corporate identity directory to Amazon Web Services IAM Identity Center.
// Then, you can associate this IAM Identity Center instance with your S3 Access
// Grants instance.
//
// Permissions You must have the s3:AssociateAccessGrantsIdentityCenter permission
// to use this operation.
//
// Additional Permissions You must also have the following permissions:
// sso:CreateApplication , sso:PutApplicationGrant , and
// sso:PutApplicationAuthenticationMethod .
func s3control_AssociateAccessGrantsIdentityCenter(cfg aws.Config, client *s3control.Client) {
	input := &s3control.AssociateAccessGrantsIdentityCenterInput{
		// AccountId: *string, // Required
		// IdentityCenterArn: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlIdentityCenterArn) > 0 {
		input.IdentityCenterArn = aws.String(_s3controlIdentityCenterArn)
	}

	if resp, err := client.AssociateAccessGrantsIdentityCenter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access grant that gives a grantee access to your S3 data. The
// grantee can be an IAM user or role or a directory user, or group. Before you can
// create a grant, you must have an S3 Access Grants instance in the same Region as
// the S3 data. You can create an S3 Access Grants instance using the [CreateAccessGrantsInstance]. You must
// also have registered at least one S3 data location in your S3 Access Grants
// instance using [CreateAccessGrantsLocation].
//
// Permissions You must have the s3:CreateAccessGrant permission to use this
// operation.
//
// Additional Permissions For any directory identity - sso:DescribeInstance and
// sso:DescribeApplication
//
// For directory users - identitystore:DescribeUser
//
// For directory groups - identitystore:DescribeGroup
//
// [CreateAccessGrantsLocation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessGrantsLocation.html
// [CreateAccessGrantsInstance]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessGrantsInstance.html
func s3control_CreateAccessGrant(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateAccessGrantInput{
		// AccessGrantsLocationId: *string, // Required
		// AccountId: *string, // Required
		// Grantee: *types.Grantee, // Required
		// Permission: types.Permission, // Required
	}

	if len(_s3controlAccessGrantsLocationId) > 0 {
		input.AccessGrantsLocationId = aws.String(_s3controlAccessGrantsLocationId)
	}
	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlGrantee) > 0 {
		if err := assignInputField(input, "Grantee", _s3controlGrantee); err != nil {
			log.Errorf("invalid --grantee: %s", err.Error())
			return
		}
	}
	if len(_s3controlPermission) > 0 {
		if err := assignInputField(input, "Permission", _s3controlPermission); err != nil {
			log.Errorf("invalid --permission: %s", err.Error())
			return
		}
	}
	if len(_s3controlAccessGrantsLocationConfiguration) > 0 {
		if err := assignInputField(input, "AccessGrantsLocationConfiguration", _s3controlAccessGrantsLocationConfiguration); err != nil {
			log.Errorf("invalid --access-grants-location-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_s3controlApplicationArn)
	}
	if len(_s3controlS3PrefixType) > 0 {
		if err := assignInputField(input, "S3PrefixType", _s3controlS3PrefixType); err != nil {
			log.Errorf("invalid --s3-prefix-type: %s", err.Error())
			return
		}
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an S3 Access Grants instance, which serves as a logical grouping for
// access grants. You can create one S3 Access Grants instance per Region per
// account.
//
// Permissions You must have the s3:CreateAccessGrantsInstance permission to use
// this operation.
//
// Additional Permissions To associate an IAM Identity Center instance with your
// S3 Access Grants instance, you must also have the sso:DescribeInstance ,
// sso:CreateApplication , sso:PutApplicationGrant , and
// sso:PutApplicationAuthenticationMethod permissions.
func s3control_CreateAccessGrantsInstance(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateAccessGrantsInstanceInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlIdentityCenterArn) > 0 {
		input.IdentityCenterArn = aws.String(_s3controlIdentityCenterArn)
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessGrantsInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The S3 data location that you would like to register in your S3 Access Grants
// instance. Your S3 data must be in the same Region as your S3 Access Grants
// instance. The location can be one of the following:
//
// - The default S3 location s3://
//
// - A bucket - S3://
//
// - A bucket and prefix - S3:///
//
// When you register a location, you must include the IAM role that has permission
// to manage the S3 location that you are registering. Give S3 Access Grants
// permission to assume this role [using a policy]. S3 Access Grants assumes this role to manage
// access to the location and to vend temporary credentials to grantees or client
// applications.
//
// Permissions You must have the s3:CreateAccessGrantsLocation permission to use
// this operation.
//
// Additional Permissions You must also have the following permission for the
// specified IAM role: iam:PassRole
//
// [using a policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-location.html
func s3control_CreateAccessGrantsLocation(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateAccessGrantsLocationInput{
		// AccountId: *string, // Required
		// IAMRoleArn: *string, // Required
		// LocationScope: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlIAMRoleArn) > 0 {
		input.IAMRoleArn = aws.String(_s3controlIAMRoleArn)
	}
	if len(_s3controlLocationScope) > 0 {
		input.LocationScope = aws.String(_s3controlLocationScope)
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateAccessGrantsLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an access point and associates it to a specified bucket. For more
// information, see [Managing access to shared datasets with access points]or [Managing access to shared datasets in directory buckets with access points] in the Amazon S3 User Guide.
//
// To create an access point and attach it to a volume on an Amazon FSx file
// system, see [CreateAndAttachS3AccessPoint]in the Amazon FSx API Reference.
//
// S3 on Outposts only supports VPC-style access points.
//
// For more information, see [Accessing Amazon S3 on Outposts using virtual private cloud (VPC) only access points] in the Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to CreateAccessPoint :
//
// [GetAccessPoint]
//
// [DeleteAccessPoint]
//
// [ListAccessPoints]
//
// [ListAccessPointsForDirectoryBuckets]
//
// [ListAccessPointsForDirectoryBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForDirectoryBuckets.html
// [ListAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
// [CreateAndAttachS3AccessPoint]: https://docs.aws.amazon.com/fsx/latest/APIReference/API_CreateAndAttachS3AccessPoint.html
// [Managing access to shared datasets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [Accessing Amazon S3 on Outposts using virtual private cloud (VPC) only access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Managing access to shared datasets in directory buckets with access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-directory-buckets.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html#API_control_CreateAccessPoint_Examples
func s3control_CreateAccessPoint(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateAccessPointInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}
	if len(_s3controlBucketAccountId) > 0 {
		input.BucketAccountId = aws.String(_s3controlBucketAccountId)
	}
	if len(_s3controlPublicAccessBlockConfiguration) > 0 {
		if err := assignInputField(input, "PublicAccessBlockConfiguration", _s3controlPublicAccessBlockConfiguration); err != nil {
			log.Errorf("invalid --public-access-block-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlScope) > 0 {
		if err := assignInputField(input, "Scope", _s3controlScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_s3controlVpcConfiguration) > 0 {
		if err := assignInputField(input, "VpcConfiguration", _s3controlVpcConfiguration); err != nil {
			log.Errorf("invalid --vpc-configuration: %s", err.Error())
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

// This operation is not supported by directory buckets.
// Creates an Object Lambda Access Point. For more information, see [Transforming objects with Object Lambda Access Points] in the Amazon
// S3 User Guide.
//
// The following actions are related to CreateAccessPointForObjectLambda :
//
// [DeleteAccessPointForObjectLambda]
//
// [GetAccessPointForObjectLambda]
//
// [ListAccessPointsForObjectLambda]
//
// [ListAccessPointsForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForObjectLambda.html
// [Transforming objects with Object Lambda Access Points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/transforming-objects.html
// [DeleteAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html
// [GetAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html
func s3control_CreateAccessPointForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateAccessPointForObjectLambdaInput{
		// AccountId: *string, // Required
		// Configuration: *types.ObjectLambdaConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _s3controlConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.CreateAccessPointForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action creates an Amazon S3 on Outposts bucket. To create an S3 bucket,
// see [Create Bucket]in the Amazon S3 API Reference.
//
// Creates a new Outposts bucket. By creating the bucket, you become the bucket
// owner. To create an Outposts bucket, you must have S3 on Outposts. For more
// information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// Not every string is an acceptable bucket name. For information on bucket naming
// restrictions, see [Working with Amazon S3 Buckets].
//
// S3 on Outposts buckets support:
//
// - Tags
//
// - LifecycleConfigurations for deleting expired objects
//
// For a complete list of restrictions and Amazon S3 feature limitations on S3 on
// Outposts, see [Amazon S3 on Outposts Restrictions and Limitations].
//
// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
// on Outposts endpoint hostname prefix and x-amz-outpost-id in your API request,
// see the [Examples]section.
//
// The following actions are related to CreateBucket for Amazon S3 on Outposts:
//
// [PutObject]
//
// [GetBucket]
//
// [DeleteBucket]
//
// [CreateAccessPoint]
//
// [PutAccessPointPolicy]
//
// [GetBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [Working with Amazon S3 Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/BucketRestrictions.html#bucketnamingrules
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html
// [Create Bucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [PutAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html#API_control_CreateBucket_Examples
// [Amazon S3 on Outposts Restrictions and Limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OnOutpostsRestrictionsLimitations.html
func s3control_CreateBucket(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateBucketInput{
		// Bucket: *string, // Required
	}

	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlACL) > 0 {
		if err := assignInputField(input, "ACL", _s3controlACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3controlCreateBucketConfiguration) > 0 {
		if err := assignInputField(input, "CreateBucketConfiguration", _s3controlCreateBucketConfiguration); err != nil {
			log.Errorf("invalid --create-bucket-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlGrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3controlGrantFullControl)
	}
	if len(_s3controlGrantRead) > 0 {
		input.GrantRead = aws.String(_s3controlGrantRead)
	}
	if len(_s3controlGrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3controlGrantReadACP)
	}
	if len(_s3controlGrantWrite) > 0 {
		input.GrantWrite = aws.String(_s3controlGrantWrite)
	}
	if len(_s3controlGrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3controlGrantWriteACP)
	}
	if len(_s3controlObjectLockEnabledForBucket) > 0 {
		if err := assignInputField(input, "ObjectLockEnabledForBucket", _s3controlObjectLockEnabledForBucket); err != nil {
			log.Errorf("invalid --object-lock-enabled-for-bucket: %s", err.Error())
			return
		}
	}
	if len(_s3controlOutpostId) > 0 {
		input.OutpostId = aws.String(_s3controlOutpostId)
	}

	if resp, err := client.CreateBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation creates an S3 Batch Operations job.
// You can use S3 Batch Operations to perform large-scale batch actions on Amazon
// S3 objects. Batch Operations can run a single action on lists of Amazon S3
// objects that you specify. For more information, see [S3 Batch Operations]in the Amazon S3 User Guide.
//
// Permissions For information about permissions required to use the Batch
// Operations, see [Granting permissions for S3 Batch Operations]in the Amazon S3 User Guide.
//
// Related actions include:
//
// [DescribeJob]
//
// [ListJobs]
//
// [UpdateJobPriority]
//
// [UpdateJobStatus]
//
// [JobOperation]
//
// [DescribeJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html
// [S3 Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html
// [Granting permissions for S3 Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops-iam-role-policies.html
// [JobOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_JobOperation.html
// [UpdateJobPriority]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html
// [UpdateJobStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html
// [ListJobs]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html
func s3control_CreateJob(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateJobInput{
		// AccountId: *string, // Required
		// ClientRequestToken: *string, // Required
		// Operation: *types.JobOperation, // Required
		// Priority: *int32, // Required
		// Report: *types.JobReport, // Required
		// RoleArn: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlClientRequestToken) > 0 {
		input.ClientRequestToken = aws.String(_s3controlClientRequestToken)
	}
	if len(_s3controlOperation) > 0 {
		if err := assignInputField(input, "Operation", _s3controlOperation); err != nil {
			log.Errorf("invalid --operation: %s", err.Error())
			return
		}
	}
	if len(_s3controlPriority) > 0 {
		if err := assignInputField(input, "Priority", _s3controlPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}
	if len(_s3controlReport) > 0 {
		if err := assignInputField(input, "Report", _s3controlReport); err != nil {
			log.Errorf("invalid --report: %s", err.Error())
			return
		}
	}
	if len(_s3controlRoleArn) > 0 {
		input.RoleArn = aws.String(_s3controlRoleArn)
	}
	if len(_s3controlConfirmationRequired) > 0 {
		if err := assignInputField(input, "ConfirmationRequired", _s3controlConfirmationRequired); err != nil {
			log.Errorf("invalid --confirmation-required: %s", err.Error())
			return
		}
	}
	if len(_s3controlDescription) > 0 {
		input.Description = aws.String(_s3controlDescription)
	}
	if len(_s3controlManifest) > 0 {
		if err := assignInputField(input, "Manifest", _s3controlManifest); err != nil {
			log.Errorf("invalid --manifest: %s", err.Error())
			return
		}
	}
	if len(_s3controlManifestGenerator) > 0 {
		if err := assignInputField(input, "ManifestGenerator", _s3controlManifestGenerator); err != nil {
			log.Errorf("invalid --manifest-generator: %s", err.Error())
			return
		}
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Creates a Multi-Region Access Point and associates it with the specified
// buckets. For more information about creating Multi-Region Access Points, see [Creating Multi-Region Access Points]in
// the Amazon S3 User Guide.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// This request is asynchronous, meaning that you might receive a response before
// the command has completed. When this request provides a response, it provides a
// token that you can use to monitor the status of the request with
// DescribeMultiRegionAccessPointOperation .
//
// The following actions are related to CreateMultiRegionAccessPoint :
//
// [DeleteMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [GetMultiRegionAccessPoint]
//
// [ListMultiRegionAccessPoints]
//
// [Creating Multi-Region Access Points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CreatingMultiRegionAccessPoints.html
// [DeleteMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [ListMultiRegionAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_CreateMultiRegionAccessPoint(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateMultiRegionAccessPointInput{
		// AccountId: *string, // Required
		// ClientToken: *string, // Required
		// Details: *types.CreateMultiRegionAccessPointInput, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlClientToken) > 0 {
		input.ClientToken = aws.String(_s3controlClientToken)
	}
	if len(_s3controlDetails) > 0 {
		if err := assignInputField(input, "Details", _s3controlDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateMultiRegionAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new S3 Storage Lens group and associates it with the specified
// Amazon Web Services account ID. An S3 Storage Lens group is a custom grouping of
// objects based on prefix, suffix, object tags, object size, object age, or a
// combination of these filters. For each Storage Lens group that you’ve created,
// you can also optionally add Amazon Web Services resource tags. For more
// information about S3 Storage Lens groups, see [Working with S3 Storage Lens groups].
//
// To use this operation, you must have the permission to perform the
// s3:CreateStorageLensGroup action. If you’re trying to create a Storage Lens
// group with Amazon Web Services resource tags, you must also have permission to
// perform the s3:TagResource action. For more information about the required
// Storage Lens Groups permissions, see [Setting account permissions to use S3 Storage Lens groups].
//
// For information about Storage Lens groups errors, see [List of Amazon S3 Storage Lens error codes].
//
// [Setting account permissions to use S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions
// [List of Amazon S3 Storage Lens error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList
// [Working with S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups-overview.html
func s3control_CreateStorageLensGroup(cfg aws.Config, client *s3control.Client) {
	input := &s3control.CreateStorageLensGroupInput{
		// AccountId: *string, // Required
		// StorageLensGroup: *types.StorageLensGroup, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlStorageLensGroup) > 0 {
		if err := assignInputField(input, "StorageLensGroup", _s3controlStorageLensGroup); err != nil {
			log.Errorf("invalid --storage-lens-group: %s", err.Error())
			return
		}
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateStorageLensGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the access grant from the S3 Access Grants instance. You cannot undo an
// access grant deletion and the grantee will no longer have access to the S3 data.
//
// Permissions You must have the s3:DeleteAccessGrant permission to use this
// operation.
func s3control_DeleteAccessGrant(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessGrantInput{
		// AccessGrantId: *string, // Required
		// AccountId: *string, // Required
	}

	if len(_s3controlAccessGrantId) > 0 {
		input.AccessGrantId = aws.String(_s3controlAccessGrantId)
	}
	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.DeleteAccessGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes your S3 Access Grants instance. You must first delete the access grants
// and locations before S3 Access Grants can delete the instance. See [DeleteAccessGrant]and [DeleteAccessGrantsLocation]. If you
// have associated an IAM Identity Center instance with your S3 Access Grants
// instance, you must first dissassociate the Identity Center instance from the S3
// Access Grants instance before you can delete the S3 Access Grants instance. See [AssociateAccessGrantsIdentityCenter]
// and [DissociateAccessGrantsIdentityCenter].
//
// Permissions You must have the s3:DeleteAccessGrantsInstance permission to use
// this operation.
//
// [DeleteAccessGrant]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessGrant.html
// [AssociateAccessGrantsIdentityCenter]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_AssociateAccessGrantsIdentityCenter.html
// [DeleteAccessGrantsLocation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessGrantsLocation.html
// [DissociateAccessGrantsIdentityCenter]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DissociateAccessGrantsIdentityCenter.html
func s3control_DeleteAccessGrantsInstance(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessGrantsInstanceInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.DeleteAccessGrantsInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the resource policy of the S3 Access Grants instance. The resource
// policy is used to manage cross-account access to your S3 Access Grants instance.
// By deleting the resource policy, you delete any cross-account permissions to
// your S3 Access Grants instance.
//
// Permissions You must have the s3:DeleteAccessGrantsInstanceResourcePolicy
// permission to use this operation.
func s3control_DeleteAccessGrantsInstanceResourcePolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessGrantsInstanceResourcePolicyInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.DeleteAccessGrantsInstanceResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deregisters a location from your S3 Access Grants instance. You can only delete
// a location registration from an S3 Access Grants instance if there are no grants
// associated with this location. See [Delete a grant]for information on how to delete grants. You
// need to have at least one registered location in your S3 Access Grants instance
// in order to create access grants.
//
// Permissions You must have the s3:DeleteAccessGrantsLocation permission to use
// this operation.
//
// [Delete a grant]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessGrant.html
func s3control_DeleteAccessGrantsLocation(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessGrantsLocationInput{
		// AccessGrantsLocationId: *string, // Required
		// AccountId: *string, // Required
	}

	if len(_s3controlAccessGrantsLocationId) > 0 {
		input.AccessGrantsLocationId = aws.String(_s3controlAccessGrantsLocationId)
	}
	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.DeleteAccessGrantsLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the specified access point.
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to DeleteAccessPoint :
//
// [CreateAccessPoint]
//
// [GetAccessPoint]
//
// [ListAccessPoints]
//
// [ListAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html#API_control_DeleteAccessPoint_Examples
func s3control_DeleteAccessPoint(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessPointInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.DeleteAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Deletes the specified Object Lambda Access Point.
//
// The following actions are related to DeleteAccessPointForObjectLambda :
//
// [CreateAccessPointForObjectLambda]
//
// [GetAccessPointForObjectLambda]
//
// [ListAccessPointsForObjectLambda]
//
// [CreateAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html
// [ListAccessPointsForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForObjectLambda.html
// [GetAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html
func s3control_DeleteAccessPointForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessPointForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.DeleteAccessPointForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the access point policy for the specified access point.
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to DeleteAccessPointPolicy :
//
// [PutAccessPointPolicy]
//
// [GetAccessPointPolicy]
//
// [GetAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html
// [PutAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html#API_control_DeleteAccessPointPolicy_Examples
func s3control_DeleteAccessPointPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessPointPolicyInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.DeleteAccessPointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Removes the resource policy for an Object Lambda Access Point.
//
// The following actions are related to DeleteAccessPointPolicyForObjectLambda :
//
// [GetAccessPointPolicyForObjectLambda]
//
// [PutAccessPointPolicyForObjectLambda]
//
// [PutAccessPointPolicyForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicyForObjectLambda.html
// [GetAccessPointPolicyForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicyForObjectLambda.html
func s3control_DeleteAccessPointPolicyForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessPointPolicyForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.DeleteAccessPointPolicyForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing access point scope for a directory bucket.
// When you delete the scope of an access point, all prefixes and permissions are
// deleted.
//
// To use this operation, you must have the permission to perform the
// s3express:DeleteAccessPointScope action.
//
// For information about REST API errors, see [REST error responses].
//
// [REST error responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#RESTErrorResponses
func s3control_DeleteAccessPointScope(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteAccessPointScopeInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.DeleteAccessPointScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action deletes an Amazon S3 on Outposts bucket. To delete an S3 bucket,
// see [DeleteBucket]in the Amazon S3 API Reference.
//
// Deletes the Amazon S3 on Outposts bucket. All objects (including all object
// versions and delete markers) in the bucket must be deleted before the bucket
// itself can be deleted. For more information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// # Related Resources
//
// [CreateBucket]
//
// [GetBucket]
//
// [DeleteObject]
//
// [GetBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html#API_control_DeleteBucket_Examples
func s3control_DeleteBucket(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteBucketInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.DeleteBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action deletes an Amazon S3 on Outposts bucket's lifecycle configuration.
// To delete an S3 bucket's lifecycle configuration, see [DeleteBucketLifecycle]in the Amazon S3 API
// Reference.
//
// Deletes the lifecycle configuration from the specified Outposts bucket. Amazon
// S3 on Outposts removes all the lifecycle configuration rules in the lifecycle
// subresource associated with the bucket. Your objects never expire, and Amazon S3
// on Outposts no longer automatically deletes any objects on the basis of rules
// contained in the deleted lifecycle configuration. For more information, see [Using Amazon S3 on Outposts]in
// Amazon S3 User Guide.
//
// To use this operation, you must have permission to perform the
// s3-outposts:PutLifecycleConfiguration action. By default, the bucket owner has
// this permission and the Outposts bucket owner can grant this permission to
// others.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// For more information about object expiration, see [Elements to Describe Lifecycle Actions].
//
// Related actions include:
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [Elements to Describe Lifecycle Actions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#intro-lifecycle-rules-actions
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html#API_control_DeleteBucketLifecycleConfiguration_Examples
func s3control_DeleteBucketLifecycleConfiguration(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteBucketLifecycleConfigurationInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.DeleteBucketLifecycleConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action deletes an Amazon S3 on Outposts bucket policy. To delete an S3
// bucket policy, see [DeleteBucketPolicy]in the Amazon S3 API Reference.
//
// This implementation of the DELETE action uses the policy subresource to delete
// the policy of a specified Amazon S3 on Outposts bucket. If you are using an
// identity other than the root user of the Amazon Web Services account that owns
// the bucket, the calling identity must have the s3-outposts:DeleteBucketPolicy
// permissions on the specified Outposts bucket and belong to the bucket owner's
// account to use this action. For more information, see [Using Amazon S3 on Outposts]in Amazon S3 User Guide.
//
// If you don't have DeleteBucketPolicy permissions, Amazon S3 returns a 403
// Access Denied error. If you have the correct permissions, but you're not using
// an identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
//
// As a security precaution, the root user of the Amazon Web Services account that
// owns a bucket can always use this action, even if the policy explicitly denies
// the root user the ability to perform this action.
//
// For more information about bucket policies, see [Using Bucket Policies and User Policies].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to DeleteBucketPolicy :
//
// [GetBucketPolicy]
//
// [PutBucketPolicy]
//
// [PutBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html
// [DeleteBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketPolicy.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [GetBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html#API_control_DeleteBucketPolicy_Examples
func s3control_DeleteBucketPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteBucketPolicyInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.DeleteBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation deletes an Amazon S3 on Outposts bucket's replication
// configuration. To delete an S3 bucket's replication configuration, see [DeleteBucketReplication]in the
// Amazon S3 API Reference.
//
// Deletes the replication configuration from the specified S3 on Outposts bucket.
//
// To use this operation, you must have permissions to perform the
// s3-outposts:PutReplicationConfiguration action. The Outposts bucket owner has
// this permission by default and can grant it to others. For more information
// about permissions, see [Setting up IAM with S3 on Outposts]and [Managing access to S3 on Outposts buckets] in the Amazon S3 User Guide.
//
// It can take a while to propagate PUT or DELETE requests for a replication
// configuration to all S3 on Outposts systems. Therefore, the replication
// configuration that's returned by a GET request soon after a PUT or DELETE
// request might return a more recent result than what's on the Outpost. If an
// Outpost is offline, the delay in updating the replication configuration on that
// Outpost can be significant.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// For information about S3 replication on Outposts configuration, see [Replicating objects for S3 on Outposts] in the
// Amazon S3 User Guide.
//
// The following operations are related to DeleteBucketReplication :
//
// [PutBucketReplication]
//
// [GetBucketReplication]
//
// [Replicating objects for S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html
// [Setting up IAM with S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html
// [Managing access to S3 on Outposts buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html#API_control_DeleteBucketReplication_Examples
func s3control_DeleteBucketReplication(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteBucketReplicationInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.DeleteBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action deletes an Amazon S3 on Outposts bucket's tags. To delete an S3
// bucket tags, see [DeleteBucketTagging]in the Amazon S3 API Reference.
//
// Deletes the tags from the Outposts bucket. For more information, see [Using Amazon S3 on Outposts] in Amazon
// S3 User Guide.
//
// To use this action, you must have permission to perform the PutBucketTagging
// action. By default, the bucket owner has this permission and can grant this
// permission to others.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to DeleteBucketTagging :
//
// [GetBucketTagging]
//
// [PutBucketTagging]
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html#API_control_DeleteBucketTagging_Examples
func s3control_DeleteBucketTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteBucketTaggingInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.DeleteBucketTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the entire tag set from the specified S3 Batch Operations job.
// Permissions To use the DeleteJobTagging operation, you must have permission to
// perform the s3:DeleteJobTagging action. For more information, see [Controlling access and labeling jobs using tags] in the
// Amazon S3 User Guide.
//
// Related actions include:
//
// [CreateJob]
//
// [GetJobTagging]
//
// [PutJobTagging]
//
// [Controlling access and labeling jobs using tags]: https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags
// [GetJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html
// [PutJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutJobTagging.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
func s3control_DeleteJobTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteJobTaggingInput{
		// AccountId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobId) > 0 {
		input.JobId = aws.String(_s3controlJobId)
	}

	if resp, err := client.DeleteJobTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Deletes a Multi-Region Access Point. This action does not delete the buckets
// associated with the Multi-Region Access Point, only the Multi-Region Access
// Point itself.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// This request is asynchronous, meaning that you might receive a response before
// the command has completed. When this request provides a response, it provides a
// token that you can use to monitor the status of the request with
// DescribeMultiRegionAccessPointOperation .
//
// The following actions are related to DeleteMultiRegionAccessPoint :
//
// [CreateMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [GetMultiRegionAccessPoint]
//
// [ListMultiRegionAccessPoints]
//
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [ListMultiRegionAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [CreateMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_DeleteMultiRegionAccessPoint(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteMultiRegionAccessPointInput{
		// AccountId: *string, // Required
		// ClientToken: *string, // Required
		// Details: *types.DeleteMultiRegionAccessPointInput, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlClientToken) > 0 {
		input.ClientToken = aws.String(_s3controlClientToken)
	}
	if len(_s3controlDetails) > 0 {
		if err := assignInputField(input, "Details", _s3controlDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteMultiRegionAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Removes the PublicAccessBlock configuration for an Amazon Web Services account.
// This operation might be restricted when the account is managed by
// organization-level Block Public Access policies. You’ll get an Access Denied
// (403) error when the account is managed by organization-level Block Public
// Access policies. Organization-level policies override account-level settings,
// preventing direct account-level modifications. For more information, see [Using Amazon S3 block public access].
//
// Related actions include:
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html
// [Using Amazon S3 block public access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
func s3control_DeletePublicAccessBlock(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeletePublicAccessBlockInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.DeletePublicAccessBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Deletes the Amazon S3 Storage Lens configuration. For more information about S3
// Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:DeleteStorageLensConfiguration action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the
// Amazon S3 User Guide.
//
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func s3control_DeleteStorageLensConfiguration(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteStorageLensConfigurationInput{
		// AccountId: *string, // Required
		// ConfigId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfigId) > 0 {
		input.ConfigId = aws.String(_s3controlConfigId)
	}

	if resp, err := client.DeleteStorageLensConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Deletes the Amazon S3 Storage Lens configuration tags. For more information
// about S3 Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:DeleteStorageLensConfigurationTagging action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in
// the Amazon S3 User Guide.
//
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func s3control_DeleteStorageLensConfigurationTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteStorageLensConfigurationTaggingInput{
		// AccountId: *string, // Required
		// ConfigId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfigId) > 0 {
		input.ConfigId = aws.String(_s3controlConfigId)
	}

	if resp, err := client.DeleteStorageLensConfigurationTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an existing S3 Storage Lens group.
// To use this operation, you must have the permission to perform the
// s3:DeleteStorageLensGroup action. For more information about the required
// Storage Lens Groups permissions, see [Setting account permissions to use S3 Storage Lens groups].
//
// For information about Storage Lens groups errors, see [List of Amazon S3 Storage Lens error codes].
//
// [Setting account permissions to use S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions
// [List of Amazon S3 Storage Lens error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList
func s3control_DeleteStorageLensGroup(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DeleteStorageLensGroupInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.DeleteStorageLensGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the configuration parameters and status for a Batch Operations job.
// For more information, see [S3 Batch Operations]in the Amazon S3 User Guide.
//
// Permissions To use the DescribeJob operation, you must have permission to
// perform the s3:DescribeJob action.
//
// Related actions include:
//
// [CreateJob]
//
// [ListJobs]
//
// [UpdateJobPriority]
//
// [UpdateJobStatus]
//
// [S3 Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html
// [UpdateJobPriority]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
// [UpdateJobStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html
// [ListJobs]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html
func s3control_DescribeJob(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DescribeJobInput{
		// AccountId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobId) > 0 {
		input.JobId = aws.String(_s3controlJobId)
	}

	if resp, err := client.DescribeJob(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Retrieves the status of an asynchronous request to manage a Multi-Region Access
// Point. For more information about managing Multi-Region Access Points and how
// asynchronous requests work, see [Using Multi-Region Access Points]in the Amazon S3 User Guide.
//
// The following actions are related to GetMultiRegionAccessPoint :
//
// [CreateMultiRegionAccessPoint]
//
// [DeleteMultiRegionAccessPoint]
//
// [GetMultiRegionAccessPoint]
//
// [ListMultiRegionAccessPoints]
//
// [Using Multi-Region Access Points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MrapOperations.html
// [DeleteMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [ListMultiRegionAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html
// [CreateMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html
func s3control_DescribeMultiRegionAccessPointOperation(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DescribeMultiRegionAccessPointOperationInput{
		// AccountId: *string, // Required
		// RequestTokenARN: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlRequestTokenARN) > 0 {
		input.RequestTokenARN = aws.String(_s3controlRequestTokenARN)
	}

	if resp, err := client.DescribeMultiRegionAccessPointOperation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Dissociates the Amazon Web Services IAM Identity Center instance from the S3
// Access Grants instance.
//
// Permissions You must have the s3:DissociateAccessGrantsIdentityCenter
// permission to use this operation.
//
// Additional Permissions You must have the sso:DeleteApplication permission to
// use this operation.
func s3control_DissociateAccessGrantsIdentityCenter(cfg aws.Config, client *s3control.Client) {
	input := &s3control.DissociateAccessGrantsIdentityCenterInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.DissociateAccessGrantsIdentityCenter(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Get the details of an access grant from your S3 Access Grants instance.
// Permissions You must have the s3:GetAccessGrant permission to use this
// operation.
func s3control_GetAccessGrant(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessGrantInput{
		// AccessGrantId: *string, // Required
		// AccountId: *string, // Required
	}

	if len(_s3controlAccessGrantId) > 0 {
		input.AccessGrantId = aws.String(_s3controlAccessGrantId)
	}
	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.GetAccessGrant(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the S3 Access Grants instance for a Region in your account.
// Permissions You must have the s3:GetAccessGrantsInstance permission to use this
// operation.
//
// GetAccessGrantsInstance is not supported for cross-account access. You can only
// call the API from the account that owns the S3 Access Grants instance.
func s3control_GetAccessGrantsInstance(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessGrantsInstanceInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.GetAccessGrantsInstance(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieve the S3 Access Grants instance that contains a particular prefix.
// Permissions You must have the s3:GetAccessGrantsInstanceForPrefix permission
// for the caller account to use this operation.
//
// Additional Permissions The prefix owner account must grant you the following
// permissions to their S3 Access Grants instance:
// s3:GetAccessGrantsInstanceForPrefix .
func s3control_GetAccessGrantsInstanceForPrefix(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessGrantsInstanceForPrefixInput{
		// AccountId: *string, // Required
		// S3Prefix: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlS3Prefix) > 0 {
		input.S3Prefix = aws.String(_s3controlS3Prefix)
	}

	if resp, err := client.GetAccessGrantsInstanceForPrefix(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the resource policy of the S3 Access Grants instance.
// Permissions You must have the s3:GetAccessGrantsInstanceResourcePolicy
// permission to use this operation.
func s3control_GetAccessGrantsInstanceResourcePolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessGrantsInstanceResourcePolicyInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.GetAccessGrantsInstanceResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the details of a particular location registered in your S3 Access
// Grants instance.
//
// Permissions You must have the s3:GetAccessGrantsLocation permission to use this
// operation.
func s3control_GetAccessGrantsLocation(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessGrantsLocationInput{
		// AccessGrantsLocationId: *string, // Required
		// AccountId: *string, // Required
	}

	if len(_s3controlAccessGrantsLocationId) > 0 {
		input.AccessGrantsLocationId = aws.String(_s3controlAccessGrantsLocationId)
	}
	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.GetAccessGrantsLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns configuration information about the specified access point.
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to GetAccessPoint :
//
// [CreateAccessPoint]
//
// [DeleteAccessPoint]
//
// [ListAccessPoints]
//
// [ListAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples
func s3control_GetAccessPoint(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Returns configuration for an Object Lambda Access Point.
//
// The following actions are related to GetAccessPointConfigurationForObjectLambda :
//
// [PutAccessPointConfigurationForObjectLambda]
//
// [PutAccessPointConfigurationForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointConfigurationForObjectLambda.html
func s3control_GetAccessPointConfigurationForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointConfigurationForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointConfigurationForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// # Returns configuration information about the specified Object Lambda Access Point
//
// The following actions are related to GetAccessPointForObjectLambda :
//
// [CreateAccessPointForObjectLambda]
//
// [DeleteAccessPointForObjectLambda]
//
// [ListAccessPointsForObjectLambda]
//
// [CreateAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html
// [ListAccessPointsForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPointsForObjectLambda.html
// [DeleteAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html
func s3control_GetAccessPointForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the access point policy associated with the specified access point.
// The following actions are related to GetAccessPointPolicy :
//
// [PutAccessPointPolicy]
//
// [DeleteAccessPointPolicy]
//
// [DeleteAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html
// [PutAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html
func s3control_GetAccessPointPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointPolicyInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Returns the resource policy for an Object Lambda Access Point.
//
// The following actions are related to GetAccessPointPolicyForObjectLambda :
//
// [DeleteAccessPointPolicyForObjectLambda]
//
// [PutAccessPointPolicyForObjectLambda]
//
// [PutAccessPointPolicyForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicyForObjectLambda.html
// [DeleteAccessPointPolicyForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicyForObjectLambda.html
func s3control_GetAccessPointPolicyForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointPolicyForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointPolicyForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Indicates whether the specified access point currently has a policy that allows
// public access. For more information about public access through access points,
// see [Managing Data Access with Amazon S3 access points]in the Amazon S3 User Guide.
//
// [Managing Data Access with Amazon S3 access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points.html
func s3control_GetAccessPointPolicyStatus(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointPolicyStatusInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointPolicyStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Returns the status of the resource policy associated with an Object Lambda
// Access Point.
func s3control_GetAccessPointPolicyStatusForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointPolicyStatusForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointPolicyStatusForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the access point scope for a directory bucket.
// To use this operation, you must have the permission to perform the
// s3express:GetAccessPointScope action.
//
// For information about REST API errors, see [REST error responses].
//
// [REST error responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#RESTErrorResponses
func s3control_GetAccessPointScope(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetAccessPointScopeInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetAccessPointScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets an Amazon S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts] in the Amazon
// S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the Outposts bucket, the calling identity must have
// the s3-outposts:GetBucket permissions on the specified Outposts bucket and
// belong to the Outposts bucket owner's account in order to use this action. Only
// users from Outposts bucket owner account with the right permissions can perform
// actions on an Outposts bucket.
//
// If you don't have s3-outposts:GetBucket permissions or you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 403
// Access Denied error.
//
// The following actions are related to GetBucket for Amazon S3 on Outposts:
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// [PutObject]
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucket.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucket.html#API_control_GetBucket_Examples
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
func s3control_GetBucket(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetBucketInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.GetBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action gets an Amazon S3 on Outposts bucket's lifecycle configuration. To
// get an S3 bucket's lifecycle configuration, see [GetBucketLifecycleConfiguration]in the Amazon S3 API Reference.
//
// Returns the lifecycle configuration information set on the Outposts bucket. For
// more information, see [Using Amazon S3 on Outposts]and for information about lifecycle configuration, see [Object Lifecycle Management]
// in Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3-outposts:GetLifecycleConfiguration action. The Outposts bucket owner has this
// permission, by default. The bucket owner can grant this permission to others.
// For more information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// GetBucketLifecycleConfiguration has the following special error:
//
// - Error code: NoSuchLifecycleConfiguration
//
// - Description: The lifecycle configuration does not exist.
//
// - HTTP Status Code: 404 Not Found
//
// - SOAP Fault Code Prefix: Client
//
// The following actions are related to GetBucketLifecycleConfiguration :
//
// [PutBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [DeleteBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html#API_control_GetBucketLifecycleConfiguration_Examples
func s3control_GetBucketLifecycleConfiguration(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetBucketLifecycleConfigurationInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.GetBucketLifecycleConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action gets a bucket policy for an Amazon S3 on Outposts bucket. To get a
// policy for an S3 bucket, see [GetBucketPolicy]in the Amazon S3 API Reference.
//
// Returns the policy of a specified Outposts bucket. For more information, see [Using Amazon S3 on Outposts]
// in the Amazon S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the bucket, the calling identity must have the
// GetBucketPolicy permissions on the specified bucket and belong to the bucket
// owner's account in order to use this action.
//
// Only users from Outposts bucket owner account with the right permissions can
// perform actions on an Outposts bucket. If you don't have
// s3-outposts:GetBucketPolicy permissions or you're not using an identity that
// belongs to the bucket owner's account, Amazon S3 returns a 403 Access Denied
// error.
//
// As a security precaution, the root user of the Amazon Web Services account that
// owns a bucket can always use this action, even if the policy explicitly denies
// the root user the ability to perform this action.
//
// For more information about bucket policies, see [Using Bucket Policies and User Policies].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to GetBucketPolicy :
//
// [GetObject]
//
// [PutBucketPolicy]
//
// [DeleteBucketPolicy]
//
// [PutBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [DeleteBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html
// [GetBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicy.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html#API_control_GetBucketPolicy_Examples
func s3control_GetBucketPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetBucketPolicyInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.GetBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation gets an Amazon S3 on Outposts bucket's replication
// configuration. To get an S3 bucket's replication configuration, see [GetBucketReplication]in the
// Amazon S3 API Reference.
//
// Returns the replication configuration of an S3 on Outposts bucket. For more
// information about S3 on Outposts, see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide. For
// information about S3 replication on Outposts configuration, see [Replicating objects for S3 on Outposts]in the Amazon
// S3 User Guide.
//
// It can take a while to propagate PUT or DELETE requests for a replication
// configuration to all S3 on Outposts systems. Therefore, the replication
// configuration that's returned by a GET request soon after a PUT or DELETE
// request might return a more recent result than what's on the Outpost. If an
// Outpost is offline, the delay in updating the replication configuration on that
// Outpost can be significant.
//
// This action requires permissions for the s3-outposts:GetReplicationConfiguration
// action. The Outposts bucket owner has this permission by default and can grant
// it to others. For more information about permissions, see [Setting up IAM with S3 on Outposts]and [Managing access to S3 on Outposts bucket] in the Amazon S3
// User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// If you include the Filter element in a replication configuration, you must also
// include the DeleteMarkerReplication , Status , and Priority elements. The
// response also returns those elements.
//
// For information about S3 on Outposts replication failure reasons, see [Replication failure reasons] in the
// Amazon S3 User Guide.
//
// The following operations are related to GetBucketReplication :
//
// [PutBucketReplication]
//
// [DeleteBucketReplication]
//
// [Replicating objects for S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html
// [Replication failure reasons]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [Setting up IAM with S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html
// [Managing access to S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html#API_control_GetBucketReplication_Examples
func s3control_GetBucketReplication(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetBucketReplicationInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.GetBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action gets an Amazon S3 on Outposts bucket's tags. To get an S3 bucket
// tags, see [GetBucketTagging]in the Amazon S3 API Reference.
//
// Returns the tag set associated with the Outposts bucket. For more information,
// see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the GetBucketTagging
// action. By default, the bucket owner has this permission and can grant this
// permission to others.
//
// GetBucketTagging has the following special error:
//
// - Error code: NoSuchTagSetError
//
// - Description: There is no tag set associated with the bucket.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to GetBucketTagging :
//
// [PutBucketTagging]
//
// [DeleteBucketTagging]
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html#API_control_GetBucketTagging_Examples
func s3control_GetBucketTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetBucketTaggingInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.GetBucketTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation returns the versioning state for S3 on Outposts buckets only. To
// return the versioning state for an S3 bucket, see [GetBucketVersioning]in the Amazon S3 API
// Reference.
//
// Returns the versioning state for an S3 on Outposts bucket. With S3 Versioning,
// you can save multiple distinct copies of your objects and recover from
// unintended user actions and application failures.
//
// If you've never set versioning on your bucket, it has no versioning state. In
// that case, the GetBucketVersioning request does not return a versioning state
// value.
//
// For more information about versioning, see [Versioning] in the Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following operations are related to GetBucketVersioning for S3 on Outposts.
//
// [PutBucketVersioning]
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketLifecycleConfiguration]
//
// [Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [PutBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketVersioning.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html#API_control_GetBucketVersioning_Examples
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html
func s3control_GetBucketVersioning(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetBucketVersioningInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}

	if resp, err := client.GetBucketVersioning(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a temporary access credential from S3 Access Grants to the grantee or
// client application. The [temporary credential]is an Amazon Web Services STS token that grants them
// access to the S3 data.
//
// Permissions You must have the s3:GetDataAccess permission to use this
// operation.
//
// Additional Permissions The IAM role that S3 Access Grants assumes must have the
// following permissions specified in the trust policy when registering the
// location: sts:AssumeRole , for directory users or groups sts:SetContext , and
// for IAM users or roles sts:SetSourceIdentity .
//
// [temporary credential]: https://docs.aws.amazon.com/STS/latest/APIReference/API_Credentials.html
func s3control_GetDataAccess(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetDataAccessInput{
		// AccountId: *string, // Required
		// Permission: types.Permission, // Required
		// Target: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlPermission) > 0 {
		if err := assignInputField(input, "Permission", _s3controlPermission); err != nil {
			log.Errorf("invalid --permission: %s", err.Error())
			return
		}
	}
	if len(_s3controlTarget) > 0 {
		input.Target = aws.String(_s3controlTarget)
	}
	if len(_s3controlDurationSeconds) > 0 {
		if err := assignInputField(input, "DurationSeconds", _s3controlDurationSeconds); err != nil {
			log.Errorf("invalid --duration-seconds: %s", err.Error())
			return
		}
	}
	if len(_s3controlPrivilege) > 0 {
		if err := assignInputField(input, "Privilege", _s3controlPrivilege); err != nil {
			log.Errorf("invalid --privilege: %s", err.Error())
			return
		}
	}
	if len(_s3controlTargetType) > 0 {
		if err := assignInputField(input, "TargetType", _s3controlTargetType); err != nil {
			log.Errorf("invalid --target-type: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetDataAccess(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the tags on an S3 Batch Operations job.
// Permissions To use the GetJobTagging operation, you must have permission to
// perform the s3:GetJobTagging action. For more information, see [Controlling access and labeling jobs using tags] in the Amazon
// S3 User Guide.
//
// Related actions include:
//
// [CreateJob]
//
// [PutJobTagging]
//
// [DeleteJobTagging]
//
// [DeleteJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html
// [Controlling access and labeling jobs using tags]: https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags
// [PutJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutJobTagging.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
func s3control_GetJobTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetJobTaggingInput{
		// AccountId: *string, // Required
		// JobId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobId) > 0 {
		input.JobId = aws.String(_s3controlJobId)
	}

	if resp, err := client.GetJobTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Returns configuration information about the specified Multi-Region Access Point.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to GetMultiRegionAccessPoint :
//
// [CreateMultiRegionAccessPoint]
//
// [DeleteMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [ListMultiRegionAccessPoints]
//
// [DeleteMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html
// [ListMultiRegionAccessPoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListMultiRegionAccessPoints.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [CreateMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_GetMultiRegionAccessPoint(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetMultiRegionAccessPointInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetMultiRegionAccessPoint(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Returns the access control policy of the specified Multi-Region Access Point.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to GetMultiRegionAccessPointPolicy :
//
// [GetMultiRegionAccessPointPolicyStatus]
//
// [PutMultiRegionAccessPointPolicy]
//
// [PutMultiRegionAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutMultiRegionAccessPointPolicy.html
// [GetMultiRegionAccessPointPolicyStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicyStatus.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_GetMultiRegionAccessPointPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetMultiRegionAccessPointPolicyInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetMultiRegionAccessPointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Indicates whether the specified Multi-Region Access Point has an access control
// policy that allows public access.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to GetMultiRegionAccessPointPolicyStatus :
//
// [GetMultiRegionAccessPointPolicy]
//
// [PutMultiRegionAccessPointPolicy]
//
// [PutMultiRegionAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutMultiRegionAccessPointPolicy.html
// [GetMultiRegionAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicy.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_GetMultiRegionAccessPointPolicyStatus(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetMultiRegionAccessPointPolicyStatusInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetMultiRegionAccessPointPolicyStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Returns the routing configuration for a Multi-Region Access Point, indicating
// which Regions are active or passive.
//
// To obtain routing control changes and failover requests, use the Amazon S3
// failover control infrastructure endpoints in these five Amazon Web Services
// Regions:
//
// - us-east-1
//
// - us-west-2
//
// - ap-southeast-2
//
// - ap-northeast-1
//
// - eu-west-1
func s3control_GetMultiRegionAccessPointRoutes(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetMultiRegionAccessPointRoutesInput{
		// AccountId: *string, // Required
		// Mrap: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlMrap) > 0 {
		input.Mrap = aws.String(_s3controlMrap)
	}

	if resp, err := client.GetMultiRegionAccessPointRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Retrieves the PublicAccessBlock configuration for an Amazon Web Services
// account. This operation returns the effective account-level configuration, which
// may inherit from organization-level policies. For more information, see [Using Amazon S3 block public access].
//
// Related actions include:
//
// [DeletePublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html
// [Using Amazon S3 block public access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
func s3control_GetPublicAccessBlock(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetPublicAccessBlockInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}

	if resp, err := client.GetPublicAccessBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Gets the Amazon S3 Storage Lens configuration. For more information, see [Assessing your storage activity and usage with Amazon S3 Storage Lens] in
// the Amazon S3 User Guide. For a complete list of S3 Storage Lens metrics, see [S3 Storage Lens metrics glossary]
// in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:GetStorageLensConfiguration action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the Amazon
// S3 User Guide.
//
// [S3 Storage Lens metrics glossary]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_metrics_glossary.html
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func s3control_GetStorageLensConfiguration(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetStorageLensConfigurationInput{
		// AccountId: *string, // Required
		// ConfigId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfigId) > 0 {
		input.ConfigId = aws.String(_s3controlConfigId)
	}

	if resp, err := client.GetStorageLensConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Gets the tags of Amazon S3 Storage Lens configuration. For more information
// about S3 Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:GetStorageLensConfigurationTagging action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the
// Amazon S3 User Guide.
//
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func s3control_GetStorageLensConfigurationTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetStorageLensConfigurationTaggingInput{
		// AccountId: *string, // Required
		// ConfigId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfigId) > 0 {
		input.ConfigId = aws.String(_s3controlConfigId)
	}

	if resp, err := client.GetStorageLensConfigurationTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the Storage Lens group configuration details.
// To use this operation, you must have the permission to perform the
// s3:GetStorageLensGroup action. For more information about the required Storage
// Lens Groups permissions, see [Setting account permissions to use S3 Storage Lens groups].
//
// For information about Storage Lens groups errors, see [List of Amazon S3 Storage Lens error codes].
//
// [Setting account permissions to use S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions
// [List of Amazon S3 Storage Lens error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList
func s3control_GetStorageLensGroup(cfg aws.Config, client *s3control.Client) {
	input := &s3control.GetStorageLensGroupInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.GetStorageLensGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the list of access grants in your S3 Access Grants instance.
// Permissions You must have the s3:ListAccessGrants permission to use this
// operation.
func s3control_ListAccessGrants(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListAccessGrantsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlApplicationArn) > 0 {
		input.ApplicationArn = aws.String(_s3controlApplicationArn)
	}
	if len(_s3controlGrantScope) > 0 {
		input.GrantScope = aws.String(_s3controlGrantScope)
	}
	if len(_s3controlGranteeIdentifier) > 0 {
		input.GranteeIdentifier = aws.String(_s3controlGranteeIdentifier)
	}
	if len(_s3controlGranteeType) > 0 {
		if err := assignInputField(input, "GranteeType", _s3controlGranteeType); err != nil {
			log.Errorf("invalid --grantee-type: %s", err.Error())
			return
		}
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}
	if len(_s3controlPermission) > 0 {
		if err := assignInputField(input, "Permission", _s3controlPermission); err != nil {
			log.Errorf("invalid --permission: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListAccessGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListAccessGrantsOutput
	p := s3control.NewListAccessGrantsPaginator(client, input)
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

// Returns a list of S3 Access Grants instances. An S3 Access Grants instance
// serves as a logical grouping for your individual access grants. You can only
// have one S3 Access Grants instance per Region per account.
//
// Permissions You must have the s3:ListAccessGrantsInstances permission to use
// this operation.
func s3control_ListAccessGrantsInstances(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListAccessGrantsInstancesInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessGrantsInstances(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListAccessGrantsInstancesOutput
	p := s3control.NewListAccessGrantsInstancesPaginator(client, input)
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

// Returns a list of the locations registered in your S3 Access Grants instance.
// Permissions You must have the s3:ListAccessGrantsLocations permission to use
// this operation.
func s3control_ListAccessGrantsLocations(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListAccessGrantsLocationsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlLocationScope) > 0 {
		input.LocationScope = aws.String(_s3controlLocationScope)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessGrantsLocations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListAccessGrantsLocationsOutput
	p := s3control.NewListAccessGrantsLocationsPaginator(client, input)
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

// This operation is not supported by directory buckets.
// Returns a list of the access points. You can retrieve up to 1,000 access points
// per call. If the call returns more than 1,000 access points (or the number
// specified in maxResults , whichever is less), the response will include a
// continuation token that you can use to list the additional access points.
//
// Returns only access points attached to S3 buckets by default. To return all
// access points specify DataSourceType as ALL .
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to ListAccessPoints :
//
// [CreateAccessPoint]
//
// [DeleteAccessPoint]
//
// [GetAccessPoint]
//
// [CreateAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPoint.html
// [GetAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html
// [DeleteAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPoint.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPoint.html#API_control_GetAccessPoint_Examples
func s3control_ListAccessPoints(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListAccessPointsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlDataSourceId) > 0 {
		input.DataSourceId = aws.String(_s3controlDataSourceId)
	}
	if len(_s3controlDataSourceType) > 0 {
		input.DataSourceType = aws.String(_s3controlDataSourceType)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListAccessPointsOutput
	p := s3control.NewListAccessPointsPaginator(client, input)
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

// Returns a list of the access points that are owned by the Amazon Web Services
// account and that are associated with the specified directory bucket.
//
// To list access points for general purpose buckets, see [ListAccesspoints].
//
// To use this operation, you must have the permission to perform the
// s3express:ListAccessPointsForDirectoryBuckets action.
//
// For information about REST API errors, see [REST error responses].
//
// [REST error responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#RESTErrorResponses
// [ListAccesspoints]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListAccessPoints.html
func s3control_ListAccessPointsForDirectoryBuckets(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListAccessPointsForDirectoryBucketsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlDirectoryBucket) > 0 {
		input.DirectoryBucket = aws.String(_s3controlDirectoryBucket)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPointsForDirectoryBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListAccessPointsForDirectoryBucketsOutput
	p := s3control.NewListAccessPointsForDirectoryBucketsPaginator(client, input)
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

// This operation is not supported by directory buckets.
// Returns some or all (up to 1,000) access points associated with the Object
// Lambda Access Point per call. If there are more access points than what can be
// returned in one call, the response will include a continuation token that you
// can use to list the additional access points.
//
// The following actions are related to ListAccessPointsForObjectLambda :
//
// [CreateAccessPointForObjectLambda]
//
// [DeleteAccessPointForObjectLambda]
//
// [GetAccessPointForObjectLambda]
//
// [CreateAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateAccessPointForObjectLambda.html
// [DeleteAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointForObjectLambda.html
// [GetAccessPointForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointForObjectLambda.html
func s3control_ListAccessPointsForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListAccessPointsForObjectLambdaInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListAccessPointsForObjectLambda(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListAccessPointsForObjectLambdaOutput
	p := s3control.NewListAccessPointsForObjectLambdaPaginator(client, input)
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

// Use this API to list the access grants that grant the caller access to Amazon
// S3 data through S3 Access Grants. The caller (grantee) can be an Identity and
// Access Management (IAM) identity or Amazon Web Services Identity Center
// corporate directory identity. You must pass the Amazon Web Services account of
// the S3 data owner (grantor) in the request. You can, optionally, narrow the
// results by GrantScope , using a fragment of the data's S3 path, and S3 Access
// Grants will return only the grants with a path that contains the path fragment.
// You can also pass the AllowedByApplication filter in the request, which returns
// only the grants authorized for applications, whether the application is the
// caller's Identity Center application or any other application ( ALL ). For more
// information, see [List the caller's access grants]in the Amazon S3 User Guide.
//
// Permissions You must have the s3:ListCallerAccessGrants permission to use this
// operation.
//
// [List the caller's access grants]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-list-grants.html
func s3control_ListCallerAccessGrants(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListCallerAccessGrantsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlAllowedByApplication) > 0 {
		if err := assignInputField(input, "AllowedByApplication", _s3controlAllowedByApplication); err != nil {
			log.Errorf("invalid --allowed-by-application: %s", err.Error())
			return
		}
	}
	if len(_s3controlGrantScope) > 0 {
		input.GrantScope = aws.String(_s3controlGrantScope)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListCallerAccessGrants(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListCallerAccessGrantsOutput
	p := s3control.NewListCallerAccessGrantsPaginator(client, input)
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

// Lists current S3 Batch Operations jobs as well as the jobs that have ended
// within the last 90 days for the Amazon Web Services account making the request.
// For more information, see [S3 Batch Operations]in the Amazon S3 User Guide.
//
// Permissions To use the ListJobs operation, you must have permission to perform
// the s3:ListJobs action.
//
// Related actions include:
//
// [CreateJob]
//
// [DescribeJob]
//
// [UpdateJobPriority]
//
// [UpdateJobStatus]
//
// [DescribeJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html
// [S3 Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html
// [UpdateJobPriority]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobPriority.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
// [UpdateJobStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html
func s3control_ListJobs(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListJobsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobStatuses) > 0 {
		if err := assignInputField(input, "JobStatuses", _s3controlJobStatuses); err != nil {
			log.Errorf("invalid --job-statuses: %s", err.Error())
			return
		}
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListJobs(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListJobsOutput
	p := s3control.NewListJobsPaginator(client, input)
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

// This operation is not supported by directory buckets.
// Returns a list of the Multi-Region Access Points currently associated with the
// specified Amazon Web Services account. Each call can return up to 100
// Multi-Region Access Points, the maximum number of Multi-Region Access Points
// that can be associated with a single account.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to ListMultiRegionAccessPoint :
//
// [CreateMultiRegionAccessPoint]
//
// [DeleteMultiRegionAccessPoint]
//
// [DescribeMultiRegionAccessPointOperation]
//
// [GetMultiRegionAccessPoint]
//
// [DeleteMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteMultiRegionAccessPoint.html
// [GetMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPoint.html
// [DescribeMultiRegionAccessPointOperation]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeMultiRegionAccessPointOperation.html
// [CreateMultiRegionAccessPoint]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateMultiRegionAccessPoint.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_ListMultiRegionAccessPoints(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListMultiRegionAccessPointsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListMultiRegionAccessPoints(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListMultiRegionAccessPointsOutput
	p := s3control.NewListMultiRegionAccessPointsPaginator(client, input)
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

// This operation is not supported by directory buckets.
// Returns a list of all Outposts buckets in an Outpost that are owned by the
// authenticated sender of the request. For more information, see [Using Amazon S3 on Outposts]in the Amazon S3
// User Guide.
//
// For an example of the request syntax for Amazon S3 on Outposts that uses the S3
// on Outposts endpoint hostname prefix and x-amz-outpost-id in your request, see
// the [Examples]section.
//
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListRegionalBuckets.html#API_control_ListRegionalBuckets_Examples
func s3control_ListRegionalBuckets(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListRegionalBucketsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3controlMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}
	if len(_s3controlOutpostId) > 0 {
		input.OutpostId = aws.String(_s3controlOutpostId)
	}

	if disablePaginator() {
		if resp, err := client.ListRegionalBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListRegionalBucketsOutput
	p := s3control.NewListRegionalBucketsPaginator(client, input)
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

// This operation is not supported by directory buckets.
// Gets a list of Amazon S3 Storage Lens configurations. For more information
// about S3 Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:ListStorageLensConfigurations action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the
// Amazon S3 User Guide.
//
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func s3control_ListStorageLensConfigurations(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListStorageLensConfigurationsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStorageLensConfigurations(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListStorageLensConfigurationsOutput
	p := s3control.NewListStorageLensConfigurationsPaginator(client, input)
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

// Lists all the Storage Lens groups in the specified home Region.
// To use this operation, you must have the permission to perform the
// s3:ListStorageLensGroups action. For more information about the required Storage
// Lens Groups permissions, see [Setting account permissions to use S3 Storage Lens groups].
//
// For information about Storage Lens groups errors, see [List of Amazon S3 Storage Lens error codes].
//
// [Setting account permissions to use S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions
// [List of Amazon S3 Storage Lens error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList
func s3control_ListStorageLensGroups(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListStorageLensGroupsInput{
		// AccountId: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlNextToken) > 0 {
		input.NextToken = aws.String(_s3controlNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListStorageLensGroups(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3control.ListStorageLensGroupsOutput
	p := s3control.NewListStorageLensGroupsPaginator(client, input)
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

// This operation allows you to list all of the tags for a specified resource.
// Each tag is a label consisting of a key and value. Tags can help you organize,
// track costs for, and control access to resources.
//
// This operation is only supported for the following Amazon S3 resources:
//
// [General purpose buckets]
//
// [Access Points for directory buckets]
//
// [Access Points for general purpose buckets]
//
// [Directory buckets]
//
// [S3 Storage Lens groups]
//
// [S3 Access Grants instances, registered locations, and grants]
// - .
//
// Permissions For general purpose buckets, access points for general purpose
// buckets, Storage Lens groups, and S3 Access Grants, you must have the
// s3:ListTagsForResource permission to use this operation.
//
// Directory bucket permissions For directory buckets, you must have the
// s3express:ListTagsForResource permission to use this operation. For more
// information about directory buckets policies and permissions, see [Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon
// S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// For information about S3 Tagging errors, see [List of Amazon S3 Tagging error codes].
//
// [Access Points for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-db-tagging.html
// [General purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging.html
// [S3 Access Grants instances, registered locations, and grants]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-tagging.html
// [List of Amazon S3 Tagging error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3TaggingErrorCodeList
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-tagging.html
// [Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-permissions.html
// [Access Points for general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-tagging.html
// [S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups.html
func s3control_ListTagsForResource(cfg aws.Config, client *s3control.Client) {
	input := &s3control.ListTagsForResourceInput{
		// AccountId: *string, // Required
		// ResourceArn: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3controlResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the resource policy of the S3 Access Grants instance.
// Permissions You must have the s3:PutAccessGrantsInstanceResourcePolicy
// permission to use this operation.
func s3control_PutAccessGrantsInstanceResourcePolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutAccessGrantsInstanceResourcePolicyInput{
		// AccountId: *string, // Required
		// Policy: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlPolicy) > 0 {
		input.Policy = aws.String(_s3controlPolicy)
	}
	if len(_s3controlOrganization) > 0 {
		input.Organization = aws.String(_s3controlOrganization)
	}

	if resp, err := client.PutAccessGrantsInstanceResourcePolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Replaces configuration for an Object Lambda Access Point.
//
// The following actions are related to PutAccessPointConfigurationForObjectLambda :
//
// [GetAccessPointConfigurationForObjectLambda]
//
// [GetAccessPointConfigurationForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointConfigurationForObjectLambda.html
func s3control_PutAccessPointConfigurationForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutAccessPointConfigurationForObjectLambdaInput{
		// AccountId: *string, // Required
		// Configuration: *types.ObjectLambdaConfiguration, // Required
		// Name: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfiguration) > 0 {
		if err := assignInputField(input, "Configuration", _s3controlConfiguration); err != nil {
			log.Errorf("invalid --configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}

	if resp, err := client.PutAccessPointConfigurationForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Associates an access policy with the specified access point. Each access point
// can have only one policy, so a request made to this API replaces any existing
// policy associated with the specified access point.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutAccessPointPolicy :
//
// [GetAccessPointPolicy]
//
// [DeleteAccessPointPolicy]
//
// [GetAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicy.html
// [DeleteAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicy.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutAccessPointPolicy.html#API_control_PutAccessPointPolicy_Examples
func s3control_PutAccessPointPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutAccessPointPolicyInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
		// Policy: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}
	if len(_s3controlPolicy) > 0 {
		input.Policy = aws.String(_s3controlPolicy)
	}

	if resp, err := client.PutAccessPointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Creates or replaces resource policy for an Object Lambda Access Point. For an
// example policy, see [Creating Object Lambda Access Points]in the Amazon S3 User Guide.
//
// The following actions are related to PutAccessPointPolicyForObjectLambda :
//
// [DeleteAccessPointPolicyForObjectLambda]
//
// [GetAccessPointPolicyForObjectLambda]
//
// [Creating Object Lambda Access Points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-create.html#olap-create-cli
// [GetAccessPointPolicyForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetAccessPointPolicyForObjectLambda.html
// [DeleteAccessPointPolicyForObjectLambda]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteAccessPointPolicyForObjectLambda.html
func s3control_PutAccessPointPolicyForObjectLambda(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutAccessPointPolicyForObjectLambdaInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
		// Policy: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}
	if len(_s3controlPolicy) > 0 {
		input.Policy = aws.String(_s3controlPolicy)
	}

	if resp, err := client.PutAccessPointPolicyForObjectLambda(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates or replaces the access point scope for a directory bucket. You can use
// the access point scope to restrict access to specific prefixes, API operations,
// or a combination of both.
//
// You can specify any amount of prefixes, but the total length of characters of
// all prefixes must be less than 256 bytes in size.
//
// To use this operation, you must have the permission to perform the
// s3express:PutAccessPointScope action.
//
// For information about REST API errors, see [REST error responses].
//
// [REST error responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#RESTErrorResponses
func s3control_PutAccessPointScope(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutAccessPointScopeInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
		// Scope: *types.Scope, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}
	if len(_s3controlScope) > 0 {
		if err := assignInputField(input, "Scope", _s3controlScope); err != nil {
			log.Errorf("invalid --scope: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutAccessPointScope(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action puts a lifecycle configuration to an Amazon S3 on Outposts bucket.
// To put a lifecycle configuration to an S3 bucket, see [PutBucketLifecycleConfiguration]in the Amazon S3 API
// Reference.
//
// Creates a new lifecycle configuration for the S3 on Outposts bucket or replaces
// an existing lifecycle configuration. Outposts buckets only support lifecycle
// configurations that delete/expire objects after a certain period of time and
// abort incomplete multipart uploads.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutBucketLifecycleConfiguration :
//
// [GetBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [DeleteBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html#API_control_PutBucketLifecycleConfiguration_Examples
func s3control_PutBucketLifecycleConfiguration(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutBucketLifecycleConfigurationInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlLifecycleConfiguration) > 0 {
		if err := assignInputField(input, "LifecycleConfiguration", _s3controlLifecycleConfiguration); err != nil {
			log.Errorf("invalid --lifecycle-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBucketLifecycleConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action puts a bucket policy to an Amazon S3 on Outposts bucket. To put a
// policy on an S3 bucket, see [PutBucketPolicy]in the Amazon S3 API Reference.
//
// Applies an Amazon S3 bucket policy to an Outposts bucket. For more information,
// see [Using Amazon S3 on Outposts]in the Amazon S3 User Guide.
//
// If you are using an identity other than the root user of the Amazon Web
// Services account that owns the Outposts bucket, the calling identity must have
// the PutBucketPolicy permissions on the specified Outposts bucket and belong to
// the bucket owner's account in order to use this action.
//
// If you don't have PutBucketPolicy permissions, Amazon S3 returns a 403 Access
// Denied error. If you have the correct permissions, but you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
//
// As a security precaution, the root user of the Amazon Web Services account that
// owns a bucket can always use this action, even if the policy explicitly denies
// the root user the ability to perform this action.
//
// For more information about bucket policies, see [Using Bucket Policies and User Policies].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutBucketPolicy :
//
// [GetBucketPolicy]
//
// [DeleteBucketPolicy]
//
// [PutBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketPolicy.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [DeleteBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketPolicy.html
// [GetBucketPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketPolicy.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketPolicy.html#API_control_PutBucketPolicy_Examples
func s3control_PutBucketPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutBucketPolicyInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
		// Policy: *string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlPolicy) > 0 {
		input.Policy = aws.String(_s3controlPolicy)
	}
	if len(_s3controlConfirmRemoveSelfBucketAccess) > 0 {
		if err := assignInputField(input, "ConfirmRemoveSelfBucketAccess", _s3controlConfirmRemoveSelfBucketAccess); err != nil {
			log.Errorf("invalid --confirm-remove-self-bucket-access: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action creates an Amazon S3 on Outposts bucket's replication
// configuration. To create an S3 bucket's replication configuration, see [PutBucketReplication]in the
// Amazon S3 API Reference.
//
// Creates a replication configuration or replaces an existing one. For
// information about S3 replication on Outposts configuration, see [Replicating objects for S3 on Outposts]in the Amazon
// S3 User Guide.
//
// It can take a while to propagate PUT or DELETE requests for a replication
// configuration to all S3 on Outposts systems. Therefore, the replication
// configuration that's returned by a GET request soon after a PUT or DELETE
// request might return a more recent result than what's on the Outpost. If an
// Outpost is offline, the delay in updating the replication configuration on that
// Outpost can be significant.
//
// Specify the replication configuration in the request body. In the replication
// configuration, you provide the following information:
//
// - The name of the destination bucket or buckets where you want S3 on Outposts
// to replicate objects
//
// - The Identity and Access Management (IAM) role that S3 on Outposts can
// assume to replicate objects on your behalf
//
// - Other relevant information, such as replication rules
//
// A replication configuration must include at least one rule and can contain a
// maximum of 100. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source Outposts bucket. To choose additional
// subsets of objects to replicate, add a rule for each subset.
//
// To specify a subset of the objects in the source Outposts bucket to apply a
// replication rule to, add the Filter element as a child of the Rule element. You
// can filter objects based on an object key prefix, one or more object tags, or
// both. When you add the Filter element in the configuration, you must also add
// the following elements: DeleteMarkerReplication , Status , and Priority .
//
// Using PutBucketReplication on Outposts requires that both the source and
// destination buckets must have versioning enabled. For information about enabling
// versioning on a bucket, see [Managing S3 Versioning for your S3 on Outposts bucket].
//
// For information about S3 on Outposts replication failure reasons, see [Replication failure reasons] in the
// Amazon S3 User Guide.
//
// # Handling Replication of Encrypted Objects
//
// Outposts buckets are encrypted at all times. All the objects in the source
// Outposts bucket are encrypted and can be replicated. Also, all the replicas in
// the destination Outposts bucket are encrypted with the same encryption key as
// the objects in the source Outposts bucket.
//
// # Permissions
//
// To create a PutBucketReplication request, you must have
// s3-outposts:PutReplicationConfiguration permissions for the bucket. The Outposts
// bucket owner has this permission by default and can grant it to others. For more
// information about permissions, see [Setting up IAM with S3 on Outposts]and [Managing access to S3 on Outposts buckets].
//
// To perform this operation, the user or role must also have the iam:CreateRole
// and iam:PassRole permissions. For more information, see [Granting a user permissions to pass a role to an Amazon Web Services service].
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following operations are related to PutBucketReplication :
//
// [GetBucketReplication]
//
// [DeleteBucketReplication]
//
// [Replicating objects for S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsReplication.html
// [Replication failure reasons]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/outposts-replication-eventbridge.html#outposts-replication-failure-codes
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketReplication.html
// [Setting up IAM with S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsIAM.html
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html
// [Managing access to S3 on Outposts buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsBucketPolicy.html
// [Managing S3 Versioning for your S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsManagingVersioning.html
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketReplication.html
// [Granting a user permissions to pass a role to an Amazon Web Services service]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketReplication.html#API_control_PutBucketReplication_Examples
func s3control_PutBucketReplication(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutBucketReplicationInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
		// ReplicationConfiguration: *types.ReplicationConfiguration, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlReplicationConfiguration) > 0 {
		if err := assignInputField(input, "ReplicationConfiguration", _s3controlReplicationConfiguration); err != nil {
			log.Errorf("invalid --replication-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action puts tags on an Amazon S3 on Outposts bucket. To put tags on an S3
// bucket, see [PutBucketTagging]in the Amazon S3 API Reference.
//
// Sets the tags for an S3 on Outposts bucket. For more information, see [Using Amazon S3 on Outposts] in the
// Amazon S3 User Guide.
//
// Use tags to organize your Amazon Web Services bill to reflect your own cost
// structure. To do this, sign up to get your Amazon Web Services account bill with
// tag key values included. Then, to see the cost of combined resources, organize
// your billing information according to resources with the same tag key values.
// For example, you can tag several resources with a specific application name, and
// then organize your billing information to see the total cost of that application
// across several services. For more information, see [Cost allocation and tagging].
//
// Within a bucket, if you add a tag that has the same key as an existing tag, the
// new value overwrites the old value. For more information, see [Using cost allocation in Amazon S3 bucket tags].
//
// To use this action, you must have permissions to perform the
// s3-outposts:PutBucketTagging action. The Outposts bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing access permissions to your Amazon S3 resources].
//
// PutBucketTagging has the following special errors:
//
// - Error code: InvalidTagError
//
// - Description: The tag provided was not a valid tag. This error can occur if
// the tag did not pass input validation. For information about tag restrictions,
// see [User-Defined Tag Restrictions]and [Amazon Web Services-Generated Cost Allocation Tag Restrictions].
//
// - Error code: MalformedXMLError
//
// - Description: The XML provided does not match the schema.
//
// - Error code: OperationAbortedError
//
// - Description: A conflicting conditional action is currently in progress
// against this resource. Try again.
//
// - Error code: InternalError
//
// - Description: The service was unable to apply the provided tag to the bucket.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following actions are related to PutBucketTagging :
//
// [GetBucketTagging]
//
// [DeleteBucketTagging]
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [Cost allocation and tagging]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteBucketTagging.html
// [Using cost allocation in Amazon S3 bucket tags]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html
// [User-Defined Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
// [Managing access permissions to your Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Using Amazon S3 on Outposts]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3onOutposts.html
// [Amazon Web Services-Generated Cost Allocation Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/aws-tag-restrictions.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketTagging.html#API_control_PutBucketTagging_Examples
func s3control_PutBucketTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutBucketTaggingInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
		// Tagging: *types.Tagging, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlTagging) > 0 {
		if err := assignInputField(input, "Tagging", _s3controlTagging); err != nil {
			log.Errorf("invalid --tagging: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBucketTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation sets the versioning state for S3 on Outposts buckets only. To
// set the versioning state for an S3 bucket, see [PutBucketVersioning]in the Amazon S3 API Reference.
//
// Sets the versioning state for an S3 on Outposts bucket. With S3 Versioning, you
// can save multiple distinct copies of your objects and recover from unintended
// user actions and application failures.
//
// You can set the versioning state to one of the following:
//
// - Enabled - Enables versioning for the objects in the bucket. All objects
// added to the bucket receive a unique version ID.
//
// - Suspended - Suspends versioning for the objects in the bucket. All objects
// added to the bucket receive the version ID null .
//
// If you've never set versioning on your bucket, it has no versioning state. In
// that case, a [GetBucketVersioning]request does not return a versioning state value.
//
// When you enable S3 Versioning, for each object in your bucket, you have a
// current version and zero or more noncurrent versions. You can configure your
// bucket S3 Lifecycle rules to expire noncurrent versions after a specified time
// period. For more information, see [Creating and managing a lifecycle configuration for your S3 on Outposts bucket]in the Amazon S3 User Guide.
//
// If you have an object expiration lifecycle configuration in your non-versioned
// bucket and you want to maintain the same permanent delete behavior when you
// enable versioning, you must add a noncurrent expiration policy. The noncurrent
// expiration lifecycle configuration will manage the deletes of the noncurrent
// object versions in the version-enabled bucket. For more information, see [Versioning]in the
// Amazon S3 User Guide.
//
// All Amazon S3 on Outposts REST API requests for this action require an
// additional parameter of x-amz-outpost-id to be passed with the request. In
// addition, you must use an S3 on Outposts endpoint hostname prefix instead of
// s3-control . For an example of the request syntax for Amazon S3 on Outposts that
// uses the S3 on Outposts endpoint hostname prefix and the x-amz-outpost-id
// derived by using the access point ARN, see the [Examples]section.
//
// The following operations are related to PutBucketVersioning for S3 on Outposts.
//
// [GetBucketVersioning]
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketLifecycleConfiguration]
//
// [Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/Versioning.html
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketLifecycleConfiguration.html
// [PutBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketVersioning.html
// [Creating and managing a lifecycle configuration for your S3 on Outposts bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/S3OutpostsLifecycleManaging.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketLifecycleConfiguration.html
// [Examples]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_PutBucketVersioning.html#API_control_PutBucketVersioning_Examples
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetBucketVersioning.html
func s3control_PutBucketVersioning(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutBucketVersioningInput{
		// AccountId: *string, // Required
		// Bucket: *string, // Required
		// VersioningConfiguration: *types.VersioningConfiguration, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlBucket) > 0 {
		input.Bucket = aws.String(_s3controlBucket)
	}
	if len(_s3controlVersioningConfiguration) > 0 {
		if err := assignInputField(input, "VersioningConfiguration", _s3controlVersioningConfiguration); err != nil {
			log.Errorf("invalid --versioning-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlMFA) > 0 {
		input.MFA = aws.String(_s3controlMFA)
	}

	if resp, err := client.PutBucketVersioning(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Sets the supplied tag-set on an S3 Batch Operations job.
// A tag is a key-value pair. You can associate S3 Batch Operations tags with any
// job by sending a PUT request against the tagging subresource that is associated
// with the job. To modify the existing tag set, you can either replace the
// existing tag set entirely, or make changes within the existing tag set by
// retrieving the existing tag set using [GetJobTagging], modify that tag set, and use this
// operation to replace the tag set with the one you modified. For more
// information, see [Controlling access and labeling jobs using tags]in the Amazon S3 User Guide.
//
// - If you send this request with an empty tag set, Amazon S3 deletes the
// existing tag set on the Batch Operations job. If you use this method, you are
// charged for a Tier 1 Request (PUT). For more information, see [Amazon S3 pricing].
//
// - For deleting existing tags for your Batch Operations job, a [DeleteJobTagging]request is
// preferred because it achieves the same result without incurring charges.
//
// - A few things to consider about using tags:
//
// - Amazon S3 limits the maximum number of tags to 50 tags per job.
//
// - You can associate up to 50 tags with a job as long as they have unique tag
// keys.
//
// - A tag key can be up to 128 Unicode characters in length, and tag values can
// be up to 256 Unicode characters in length.
//
// - The key and values are case sensitive.
//
// - For tagging-related restrictions related to characters and encodings, see [User-Defined Tag Restrictions]
// in the Billing and Cost Management User Guide.
//
// Permissions To use the PutJobTagging operation, you must have permission to
// perform the s3:PutJobTagging action.
//
// Related actions include:
//
// [CreateJob]
//
// [GetJobTagging]
//
// [DeleteJobTagging]
//
// [DeleteJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeleteJobTagging.html
// [Controlling access and labeling jobs using tags]: https://docs.aws.amazon.com/AmazonS3/latest/dev/batch-ops-managing-jobs.html#batch-ops-job-tags
// [User-Defined Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
// [GetJobTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetJobTagging.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
// [Amazon S3 pricing]: http://aws.amazon.com/s3/pricing/
func s3control_PutJobTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutJobTaggingInput{
		// AccountId: *string, // Required
		// JobId: *string, // Required
		// Tags: []types.S3Tag, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobId) > 0 {
		input.JobId = aws.String(_s3controlJobId)
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutJobTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Associates an access control policy with the specified Multi-Region Access
// Point. Each Multi-Region Access Point can have only one policy, so a request
// made to this action replaces any existing policy that is associated with the
// specified Multi-Region Access Point.
//
// This action will always be routed to the US West (Oregon) Region. For more
// information about the restrictions around working with Multi-Region Access
// Points, see [Multi-Region Access Point restrictions and limitations]in the Amazon S3 User Guide.
//
// The following actions are related to PutMultiRegionAccessPointPolicy :
//
// [GetMultiRegionAccessPointPolicy]
//
// [GetMultiRegionAccessPointPolicyStatus]
//
// [GetMultiRegionAccessPointPolicyStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicyStatus.html
// [GetMultiRegionAccessPointPolicy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetMultiRegionAccessPointPolicy.html
// [Multi-Region Access Point restrictions and limitations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/MultiRegionAccessPointRestrictions.html
func s3control_PutMultiRegionAccessPointPolicy(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutMultiRegionAccessPointPolicyInput{
		// AccountId: *string, // Required
		// ClientToken: *string, // Required
		// Details: *types.PutMultiRegionAccessPointPolicyInput, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlClientToken) > 0 {
		input.ClientToken = aws.String(_s3controlClientToken)
	}
	if len(_s3controlDetails) > 0 {
		if err := assignInputField(input, "Details", _s3controlDetails); err != nil {
			log.Errorf("invalid --details: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutMultiRegionAccessPointPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Creates or modifies the PublicAccessBlock configuration for an Amazon Web
// Services account. This operation may be restricted when the account is managed
// by organization-level Block Public Access policies. You might get an Access
// Denied (403) error when the account is managed by organization-level Block
// Public Access policies. Organization-level policies override account-level
// settings, preventing direct account-level modifications. For this operation,
// users must have the s3:PutAccountPublicAccessBlock permission. For more
// information, see [Using Amazon S3 block public access].
//
// Related actions include:
//
// [GetPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_GetPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DeletePublicAccessBlock.html
// [Using Amazon S3 block public access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
func s3control_PutPublicAccessBlock(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutPublicAccessBlockInput{
		// AccountId: *string, // Required
		// PublicAccessBlockConfiguration: *types.PublicAccessBlockConfiguration, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlPublicAccessBlockConfiguration) > 0 {
		if err := assignInputField(input, "PublicAccessBlockConfiguration", _s3controlPublicAccessBlockConfiguration); err != nil {
			log.Errorf("invalid --public-access-block-configuration: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutPublicAccessBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Puts an Amazon S3 Storage Lens configuration. For more information about S3
// Storage Lens, see [Working with Amazon S3 Storage Lens]in the Amazon S3 User Guide. For a complete list of S3
// Storage Lens metrics, see [S3 Storage Lens metrics glossary]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:PutStorageLensConfiguration action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the Amazon
// S3 User Guide.
//
// [Working with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
// [S3 Storage Lens metrics glossary]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_metrics_glossary.html
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
func s3control_PutStorageLensConfiguration(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutStorageLensConfigurationInput{
		// AccountId: *string, // Required
		// ConfigId: *string, // Required
		// StorageLensConfiguration: *types.StorageLensConfiguration, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfigId) > 0 {
		input.ConfigId = aws.String(_s3controlConfigId)
	}
	if len(_s3controlStorageLensConfiguration) > 0 {
		if err := assignInputField(input, "StorageLensConfiguration", _s3controlStorageLensConfiguration); err != nil {
			log.Errorf("invalid --storage-lens-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutStorageLensConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Put or replace tags on an existing Amazon S3 Storage Lens configuration. For
// more information about S3 Storage Lens, see [Assessing your storage activity and usage with Amazon S3 Storage Lens]in the Amazon S3 User Guide.
//
// To use this action, you must have permission to perform the
// s3:PutStorageLensConfigurationTagging action. For more information, see [Setting permissions to use Amazon S3 Storage Lens] in the
// Amazon S3 User Guide.
//
// [Setting permissions to use Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens_iam_permissions.html
// [Assessing your storage activity and usage with Amazon S3 Storage Lens]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage_lens.html
func s3control_PutStorageLensConfigurationTagging(cfg aws.Config, client *s3control.Client) {
	input := &s3control.PutStorageLensConfigurationTaggingInput{
		// AccountId: *string, // Required
		// ConfigId: *string, // Required
		// Tags: []types.StorageLensTag, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlConfigId) > 0 {
		input.ConfigId = aws.String(_s3controlConfigId)
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutStorageLensConfigurationTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported by directory buckets.
// Submits an updated route configuration for a Multi-Region Access Point. This
// API operation updates the routing status for the specified Regions from active
// to passive, or from passive to active. A value of 0 indicates a passive status,
// which means that traffic won't be routed to the specified Region. A value of 100
// indicates an active status, which means that traffic will be routed to the
// specified Region. At least one Region must be active at all times.
//
// When the routing configuration is changed, any in-progress operations (uploads,
// copies, deletes, and so on) to formerly active Regions will continue to run to
// their final completion state (success or failure). The routing configurations of
// any Regions that aren’t specified remain unchanged.
//
// Updated routing configurations might not be immediately applied. It can take up
// to 2 minutes for your changes to take effect.
//
// To submit routing control changes and failover requests, use the Amazon S3
// failover control infrastructure endpoints in these five Amazon Web Services
// Regions:
//
// - us-east-1
//
// - us-west-2
//
// - ap-southeast-2
//
// - ap-northeast-1
//
// - eu-west-1
func s3control_SubmitMultiRegionAccessPointRoutes(cfg aws.Config, client *s3control.Client) {
	input := &s3control.SubmitMultiRegionAccessPointRoutesInput{
		// AccountId: *string, // Required
		// Mrap: *string, // Required
		// RouteUpdates: []types.MultiRegionAccessPointRoute, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlMrap) > 0 {
		input.Mrap = aws.String(_s3controlMrap)
	}
	if len(_s3controlRouteUpdates) > 0 {
		if err := assignInputField(input, "RouteUpdates", _s3controlRouteUpdates); err != nil {
			log.Errorf("invalid --route-updates: %s", err.Error())
			return
		}
	}

	if resp, err := client.SubmitMultiRegionAccessPointRoutes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new user-defined tag or updates an existing tag. Each tag is a label
// consisting of a key and value that is applied to your resource. Tags can help
// you organize, track costs for, and control access to your resources. You can add
// up to 50 Amazon Web Services resource tags for each S3 resource.
//
// This operation is only supported for the following Amazon S3 resource:
//
// [General purpose buckets]
//
// [Access Points for directory buckets]
//
// [Access Points for general purpose buckets]
//
// [Directory buckets]
//
// [S3 Storage Lens groups]
//
// [S3 Access Grants instances, registered locations, or grants]
// - .
//
// Permissions For general purpose buckets, access points for general purpose
// buckets, Storage Lens groups, and S3 Access Grants, you must have the
// s3:TagResource permission to use this operation.
//
// Directory bucket permissions For directory buckets, you must have the
// s3express:TagResource permission to use this operation. For more information
// about directory buckets policies and permissions, see [Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User
// Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// For information about S3 Tagging errors, see [List of Amazon S3 Tagging error codes].
//
// [Access Points for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-db-tagging.html
// [General purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging.html
// [List of Amazon S3 Tagging error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3TaggingErrorCodeList
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-tagging.html
// [Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-permissions.html
// [Access Points for general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-tagging.html
// [S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups.html
// [S3 Access Grants instances, registered locations, or grants]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-tagging.html
func s3control_TagResource(cfg aws.Config, client *s3control.Client) {
	input := &s3control.TagResourceInput{
		// AccountId: *string, // Required
		// ResourceArn: *string, // Required
		// Tags: []types.Tag, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3controlResourceArn)
	}
	if len(_s3controlTags) > 0 {
		if err := assignInputField(input, "Tags", _s3controlTags); err != nil {
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

// This operation removes the specified user-defined tags from an S3 resource. You
// can pass one or more tag keys.
//
// This operation is only supported for the following Amazon S3 resources:
//
// [General purpose buckets]
//
// [Access Points for directory buckets]
//
// [Access Points for general purpose buckets]
//
// [Directory buckets]
//
// [S3 Storage Lens groups]
//
// [S3 Access Grants instances, registered locations, and grants]
// - .
//
// Permissions For general purpose buckets, access points for general purpose
// buckets, Storage Lens groups, and S3 Access Grants, you must have the
// s3:UntagResource permission to use this operation.
//
// Directory bucket permissions For directory buckets, you must have the
// s3express:UntagResource permission to use this operation. For more information
// about directory buckets policies and permissions, see [Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User
// Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// For information about S3 Tagging errors, see [List of Amazon S3 Tagging error codes].
//
// [Access Points for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-db-tagging.html
// [General purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging.html
// [S3 Access Grants instances, registered locations, and grants]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-grants-tagging.html
// [List of Amazon S3 Tagging error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3TaggingErrorCodeList
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-tagging.html
// [Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-permissions.html
// [Access Points for general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-points-tagging.html
// [S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-lens-groups.html
func s3control_UntagResource(cfg aws.Config, client *s3control.Client) {
	input := &s3control.UntagResourceInput{
		// AccountId: *string, // Required
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3controlResourceArn)
	}
	if len(_s3controlTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _s3controlTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the IAM role of a registered location in your S3 Access Grants instance.
// Permissions You must have the s3:UpdateAccessGrantsLocation permission to use
// this operation.
//
// Additional Permissions You must also have the following permission: iam:PassRole
func s3control_UpdateAccessGrantsLocation(cfg aws.Config, client *s3control.Client) {
	input := &s3control.UpdateAccessGrantsLocationInput{
		// AccessGrantsLocationId: *string, // Required
		// AccountId: *string, // Required
		// IAMRoleArn: *string, // Required
	}

	if len(_s3controlAccessGrantsLocationId) > 0 {
		input.AccessGrantsLocationId = aws.String(_s3controlAccessGrantsLocationId)
	}
	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlIAMRoleArn) > 0 {
		input.IAMRoleArn = aws.String(_s3controlIAMRoleArn)
	}

	if resp, err := client.UpdateAccessGrantsLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates an existing S3 Batch Operations job's priority. For more information,
// see [S3 Batch Operations]in the Amazon S3 User Guide.
//
// Permissions To use the UpdateJobPriority operation, you must have permission to
// perform the s3:UpdateJobPriority action.
//
// Related actions include:
//
// [CreateJob]
//
// [ListJobs]
//
// [DescribeJob]
//
// [UpdateJobStatus]
//
// [DescribeJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html
// [S3 Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
// [UpdateJobStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html
// [ListJobs]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html
func s3control_UpdateJobPriority(cfg aws.Config, client *s3control.Client) {
	input := &s3control.UpdateJobPriorityInput{
		// AccountId: *string, // Required
		// JobId: *string, // Required
		// Priority: int32, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobId) > 0 {
		input.JobId = aws.String(_s3controlJobId)
	}
	if len(_s3controlPriority) > 0 {
		if err := assignInputField(input, "Priority", _s3controlPriority); err != nil {
			log.Errorf("invalid --priority: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateJobPriority(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the status for the specified job. Use this operation to confirm that
// you want to run a job or to cancel an existing job. For more information, see [S3 Batch Operations]
// in the Amazon S3 User Guide.
//
// Permissions To use the UpdateJobStatus operation, you must have permission to
// perform the s3:UpdateJobStatus action.
//
// Related actions include:
//
// [CreateJob]
//
// [ListJobs]
//
// [DescribeJob]
//
// [UpdateJobStatus]
//
// [DescribeJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_DescribeJob.html
// [S3 Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/batch-ops.html
// [CreateJob]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateJob.html
// [UpdateJobStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UpdateJobStatus.html
// [ListJobs]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListJobs.html
func s3control_UpdateJobStatus(cfg aws.Config, client *s3control.Client) {
	input := &s3control.UpdateJobStatusInput{
		// AccountId: *string, // Required
		// JobId: *string, // Required
		// RequestedJobStatus: types.RequestedJobStatus, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlJobId) > 0 {
		input.JobId = aws.String(_s3controlJobId)
	}
	if len(_s3controlRequestedJobStatus) > 0 {
		if err := assignInputField(input, "RequestedJobStatus", _s3controlRequestedJobStatus); err != nil {
			log.Errorf("invalid --requested-job-status: %s", err.Error())
			return
		}
	}
	if len(_s3controlStatusUpdateReason) > 0 {
		input.StatusUpdateReason = aws.String(_s3controlStatusUpdateReason)
	}

	if resp, err := client.UpdateJobStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Updates the existing Storage Lens group.
// To use this operation, you must have the permission to perform the
// s3:UpdateStorageLensGroup action. For more information about the required
// Storage Lens Groups permissions, see [Setting account permissions to use S3 Storage Lens groups].
//
// For information about Storage Lens groups errors, see [List of Amazon S3 Storage Lens error codes].
//
// [Setting account permissions to use S3 Storage Lens groups]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage_lens_iam_permissions.html#storage_lens_groups_permissions
// [List of Amazon S3 Storage Lens error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#S3LensErrorCodeList
func s3control_UpdateStorageLensGroup(cfg aws.Config, client *s3control.Client) {
	input := &s3control.UpdateStorageLensGroupInput{
		// AccountId: *string, // Required
		// Name: *string, // Required
		// StorageLensGroup: *types.StorageLensGroup, // Required
	}

	if len(_s3controlAccountId) > 0 {
		input.AccountId = aws.String(_s3controlAccountId)
	}
	if len(_s3controlName) > 0 {
		input.Name = aws.String(_s3controlName)
	}
	if len(_s3controlStorageLensGroup) > 0 {
		if err := assignInputField(input, "StorageLensGroup", _s3controlStorageLensGroup); err != nil {
			log.Errorf("invalid --storage-lens-group: %s", err.Error())
			return
		}
	}

	if resp, err := client.UpdateStorageLensGroup(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_s3controlCmd)
	_s3controlCmd.Flags().SortFlags = false

	_s3controlCmd.Flags().StringVarP(&_awsProfile, "profile", "", "default", "Use Profile from ~/.aws/creds")
	_s3controlCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_s3controlCmd.Flags().StringVarP(&_awsOutput, "output", "o", "json", "Output format: json|yaml|text|table|csv|markdown|html")

	_s3controlCmd.Flags().StringVarP(&_s3controlAccessGrantId, "access-grant-id", "", "", "Access Grant ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlAccessGrantsLocationConfiguration, "access-grants-location-configuration", "", "", "Access Grants Location Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlAccessGrantsLocationId, "access-grants-location-id", "", "", "Access Grants Location ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlAccountId, "account-id", "", "", "Account ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlACL, "acl", "", "", "ACL")
	_s3controlCmd.Flags().StringVarP(&_s3controlAllowedByApplication, "allowed-by-application", "", "", "Allowed By Application")
	_s3controlCmd.Flags().StringVarP(&_s3controlApplicationArn, "application-arn", "", "", "Application ARN")
	_s3controlCmd.Flags().StringVarP(&_s3controlBucket, "bucket", "", "", "Bucket")
	_s3controlCmd.Flags().StringVarP(&_s3controlBucketAccountId, "bucket-account-id", "", "", "Bucket Account ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlClientRequestToken, "client-request-token", "", "", "Client Request Token")
	_s3controlCmd.Flags().StringVarP(&_s3controlClientToken, "client-token", "", "", "Client Token")
	_s3controlCmd.Flags().StringVarP(&_s3controlConfigId, "config-id", "", "", "Config ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlConfiguration, "configuration", "", "", "Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlConfirmRemoveSelfBucketAccess, "confirm-remove-self-bucket-access", "", "", "Confirm Remove Self Bucket Access")
	_s3controlCmd.Flags().StringVarP(&_s3controlConfirmationRequired, "confirmation-required", "", "", "Confirmation Required")
	_s3controlCmd.Flags().StringVarP(&_s3controlCreateBucketConfiguration, "create-bucket-configuration", "", "", "Create Bucket Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlDataSourceId, "data-source-id", "", "", "Data Source ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlDataSourceType, "data-source-type", "", "", "Data Source Type")
	_s3controlCmd.Flags().StringVarP(&_s3controlDescription, "description", "", "", "Description")
	_s3controlCmd.Flags().StringVarP(&_s3controlDetails, "details", "", "", "Details")
	_s3controlCmd.Flags().StringVarP(&_s3controlDirectoryBucket, "directory-bucket", "", "", "Directory Bucket")
	_s3controlCmd.Flags().StringVarP(&_s3controlDurationSeconds, "duration-seconds", "", "", "Duration Seconds")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantFullControl, "grant-full-control", "", "", "Grant Full Control")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantRead, "grant-read", "", "", "Grant Read")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantReadACP, "grant-read-acp", "", "", "Grant Read Acp")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantScope, "grant-scope", "", "", "Grant Scope")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantWrite, "grant-write", "", "", "Grant Write")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantWriteACP, "grant-write-acp", "", "", "Grant Write Acp")
	_s3controlCmd.Flags().StringVarP(&_s3controlGrantee, "grantee", "", "", "Grantee")
	_s3controlCmd.Flags().StringVarP(&_s3controlGranteeIdentifier, "grantee-identifier", "", "", "Grantee Identifier")
	_s3controlCmd.Flags().StringVarP(&_s3controlGranteeType, "grantee-type", "", "", "Grantee Type")
	_s3controlCmd.Flags().StringVarP(&_s3controlIAMRoleArn, "iam-role-arn", "", "", "IAM Role ARN")
	_s3controlCmd.Flags().StringVarP(&_s3controlIdentityCenterArn, "identity-center-arn", "", "", "Identity Center ARN")
	_s3controlCmd.Flags().StringVarP(&_s3controlJobId, "job-id", "", "", "Job ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlJobStatuses, "job-statuses", "", "", "Job Statuses")
	_s3controlCmd.Flags().StringVarP(&_s3controlLifecycleConfiguration, "lifecycle-configuration", "", "", "Lifecycle Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlLocationScope, "location-scope", "", "", "Location Scope")
	_s3controlCmd.Flags().StringVarP(&_s3controlManifest, "manifest", "", "", "Manifest")
	_s3controlCmd.Flags().StringVarP(&_s3controlManifestGenerator, "manifest-generator", "", "", "Manifest Generator")
	_s3controlCmd.Flags().StringVarP(&_s3controlMaxResults, "max-results", "", "", "Max Results")
	_s3controlCmd.Flags().StringVarP(&_s3controlMFA, "mfa", "", "", "MFA")
	_s3controlCmd.Flags().StringVarP(&_s3controlMrap, "mrap", "", "", "Mrap")
	_s3controlCmd.Flags().StringVarP(&_s3controlName, "name", "", "", "Name")
	_s3controlCmd.Flags().StringVarP(&_s3controlNextToken, "next-token", "", "", "Next Token")
	_s3controlCmd.Flags().StringVarP(&_s3controlObjectLockEnabledForBucket, "object-lock-enabled-for-bucket", "", "", "Object Lock Enabled For Bucket")
	_s3controlCmd.Flags().StringVarP(&_s3controlOperation, "operation", "", "", "Operation")
	_s3controlCmd.Flags().StringVarP(&_s3controlOrganization, "organization", "", "", "Organization")
	_s3controlCmd.Flags().StringVarP(&_s3controlOutpostId, "outpost-id", "", "", "Outpost ID")
	_s3controlCmd.Flags().StringVarP(&_s3controlPermission, "permission", "", "", "Permission")
	_s3controlCmd.Flags().StringVarP(&_s3controlPolicy, "policy", "", "", "Policy")
	_s3controlCmd.Flags().StringVarP(&_s3controlPriority, "priority", "", "", "Priority")
	_s3controlCmd.Flags().StringVarP(&_s3controlPrivilege, "privilege", "", "", "Privilege")
	_s3controlCmd.Flags().StringVarP(&_s3controlPublicAccessBlockConfiguration, "public-access-block-configuration", "", "", "Public Access Block Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlReplicationConfiguration, "replication-configuration", "", "", "Replication Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlReport, "report", "", "", "Report")
	_s3controlCmd.Flags().StringVarP(&_s3controlRequestTokenARN, "request-token-arn", "", "", "Request Token ARN")
	_s3controlCmd.Flags().StringVarP(&_s3controlRequestedJobStatus, "requested-job-status", "", "", "Requested Job Status")
	_s3controlCmd.Flags().StringVarP(&_s3controlResourceArn, "resource-arn", "", "", "Resource ARN")
	_s3controlCmd.Flags().StringVarP(&_s3controlRoleArn, "role-arn", "", "", "Role ARN")
	_s3controlCmd.Flags().StringVarP(&_s3controlRouteUpdates, "route-updates", "", "", "Route Updates")
	_s3controlCmd.Flags().StringVarP(&_s3controlS3Prefix, "s3-prefix", "", "", "S3 Prefix")
	_s3controlCmd.Flags().StringVarP(&_s3controlS3PrefixType, "s3-prefix-type", "", "", "S3 Prefix Type")
	_s3controlCmd.Flags().StringVarP(&_s3controlScope, "scope", "", "", "Scope")
	_s3controlCmd.Flags().StringVarP(&_s3controlStatusUpdateReason, "status-update-reason", "", "", "Status Update Reason")
	_s3controlCmd.Flags().StringVarP(&_s3controlStorageLensConfiguration, "storage-lens-configuration", "", "", "Storage Lens Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlStorageLensGroup, "storage-lens-group", "", "", "Storage Lens Group")
	_s3controlCmd.Flags().StringSliceVarP(&_s3controlTagKeys, "tag-keys", "", nil, "Tag Keys")
	_s3controlCmd.Flags().StringVarP(&_s3controlTagging, "tagging", "", "", "Tagging")
	_s3controlCmd.Flags().StringVarP(&_s3controlTags, "tags", "", "", "Tags")
	_s3controlCmd.Flags().StringVarP(&_s3controlTarget, "target", "", "", "Target")
	_s3controlCmd.Flags().StringVarP(&_s3controlTargetType, "target-type", "", "", "Target Type")
	_s3controlCmd.Flags().StringVarP(&_s3controlVersioningConfiguration, "versioning-configuration", "", "", "Versioning Configuration")
	_s3controlCmd.Flags().StringVarP(&_s3controlVpcConfiguration, "vpc-configuration", "", "", "VPC Configuration")

	_s3controlCmd.Flags().BoolVarP(&_s3controlAssociateAccessGrantsIdentityCenter, "associate-access-grants-identity-center", "", false, "Associate Access Grants Identity Center")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateAccessGrant, "create-access-grant", "", false, "Create Access Grant")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateAccessGrantsInstance, "create-access-grants-instance", "", false, "Create Access Grants Instance")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateAccessGrantsLocation, "create-access-grants-location", "", false, "Create Access Grants Location")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateAccessPoint, "create-access-point", "", false, "Create Access Point")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateAccessPointForObjectLambda, "create-access-point-for-object-lambda", "", false, "Create Access Point For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateBucket, "create-bucket", "", false, "Create Bucket")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateJob, "create-job", "", false, "Create Job")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateMultiRegionAccessPoint, "create-multi-region-access-point", "", false, "Create Multi Region Access Point")
	_s3controlCmd.Flags().BoolVarP(&_s3controlCreateStorageLensGroup, "create-storage-lens-group", "", false, "Create Storage Lens Group")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessGrant, "delete-access-grant", "", false, "Delete Access Grant")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessGrantsInstance, "delete-access-grants-instance", "", false, "Delete Access Grants Instance")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessGrantsInstanceResourcePolicy, "delete-access-grants-instance-resource-policy", "", false, "Delete Access Grants Instance Resource Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessGrantsLocation, "delete-access-grants-location", "", false, "Delete Access Grants Location")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessPoint, "delete-access-point", "", false, "Delete Access Point")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessPointForObjectLambda, "delete-access-point-for-object-lambda", "", false, "Delete Access Point For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessPointPolicy, "delete-access-point-policy", "", false, "Delete Access Point Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessPointPolicyForObjectLambda, "delete-access-point-policy-for-object-lambda", "", false, "Delete Access Point Policy For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteAccessPointScope, "delete-access-point-scope", "", false, "Delete Access Point Scope")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteBucket, "delete-bucket", "", false, "Delete Bucket")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteBucketLifecycleConfiguration, "delete-bucket-lifecycle-configuration", "", false, "Delete Bucket Lifecycle Configuration")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteBucketPolicy, "delete-bucket-policy", "", false, "Delete Bucket Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteBucketReplication, "delete-bucket-replication", "", false, "Delete Bucket Replication")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteBucketTagging, "delete-bucket-tagging", "", false, "Delete Bucket Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteJobTagging, "delete-job-tagging", "", false, "Delete Job Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteMultiRegionAccessPoint, "delete-multi-region-access-point", "", false, "Delete Multi Region Access Point")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeletePublicAccessBlock, "delete-public-access-block", "", false, "Delete Public Access Block")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteStorageLensConfiguration, "delete-storage-lens-configuration", "", false, "Delete Storage Lens Configuration")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteStorageLensConfigurationTagging, "delete-storage-lens-configuration-tagging", "", false, "Delete Storage Lens Configuration Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDeleteStorageLensGroup, "delete-storage-lens-group", "", false, "Delete Storage Lens Group")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDescribeJob, "describe-job", "", false, "Describe Job")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDescribeMultiRegionAccessPointOperation, "describe-multi-region-access-point-operation", "", false, "Describe Multi Region Access Point Operation")
	_s3controlCmd.Flags().BoolVarP(&_s3controlDissociateAccessGrantsIdentityCenter, "dissociate-access-grants-identity-center", "", false, "Dissociate Access Grants Identity Center")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessGrant, "get-access-grant", "", false, "Get Access Grant")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessGrantsInstance, "get-access-grants-instance", "", false, "Get Access Grants Instance")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessGrantsInstanceForPrefix, "get-access-grants-instance-for-prefix", "", false, "Get Access Grants Instance For Prefix")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessGrantsInstanceResourcePolicy, "get-access-grants-instance-resource-policy", "", false, "Get Access Grants Instance Resource Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessGrantsLocation, "get-access-grants-location", "", false, "Get Access Grants Location")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPoint, "get-access-point", "", false, "Get Access Point")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointConfigurationForObjectLambda, "get-access-point-configuration-for-object-lambda", "", false, "Get Access Point Configuration For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointForObjectLambda, "get-access-point-for-object-lambda", "", false, "Get Access Point For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointPolicy, "get-access-point-policy", "", false, "Get Access Point Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointPolicyForObjectLambda, "get-access-point-policy-for-object-lambda", "", false, "Get Access Point Policy For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointPolicyStatus, "get-access-point-policy-status", "", false, "Get Access Point Policy Status")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointPolicyStatusForObjectLambda, "get-access-point-policy-status-for-object-lambda", "", false, "Get Access Point Policy Status For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetAccessPointScope, "get-access-point-scope", "", false, "Get Access Point Scope")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetBucket, "get-bucket", "", false, "Get Bucket")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetBucketLifecycleConfiguration, "get-bucket-lifecycle-configuration", "", false, "Get Bucket Lifecycle Configuration")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetBucketPolicy, "get-bucket-policy", "", false, "Get Bucket Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetBucketReplication, "get-bucket-replication", "", false, "Get Bucket Replication")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetBucketTagging, "get-bucket-tagging", "", false, "Get Bucket Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetBucketVersioning, "get-bucket-versioning", "", false, "Get Bucket Versioning")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetDataAccess, "get-data-access", "", false, "Get Data Access")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetJobTagging, "get-job-tagging", "", false, "Get Job Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetMultiRegionAccessPoint, "get-multi-region-access-point", "", false, "Get Multi Region Access Point")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetMultiRegionAccessPointPolicy, "get-multi-region-access-point-policy", "", false, "Get Multi Region Access Point Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetMultiRegionAccessPointPolicyStatus, "get-multi-region-access-point-policy-status", "", false, "Get Multi Region Access Point Policy Status")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetMultiRegionAccessPointRoutes, "get-multi-region-access-point-routes", "", false, "Get Multi Region Access Point Routes")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetPublicAccessBlock, "get-public-access-block", "", false, "Get Public Access Block")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetStorageLensConfiguration, "get-storage-lens-configuration", "", false, "Get Storage Lens Configuration")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetStorageLensConfigurationTagging, "get-storage-lens-configuration-tagging", "", false, "Get Storage Lens Configuration Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlGetStorageLensGroup, "get-storage-lens-group", "", false, "Get Storage Lens Group")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListAccessGrants, "list-access-grants", "", false, "List Access Grants")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListAccessGrantsInstances, "list-access-grants-instances", "", false, "List Access Grants Instances")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListAccessGrantsLocations, "list-access-grants-locations", "", false, "List Access Grants Locations")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListAccessPoints, "list-access-points", "", false, "List Access Points")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListAccessPointsForDirectoryBuckets, "list-access-points-for-directory-buckets", "", false, "List Access Points For Directory Buckets")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListAccessPointsForObjectLambda, "list-access-points-for-object-lambda", "", false, "List Access Points For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListCallerAccessGrants, "list-caller-access-grants", "", false, "List Caller Access Grants")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListJobs, "list-jobs", "", false, "List Jobs")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListMultiRegionAccessPoints, "list-multi-region-access-points", "", false, "List Multi Region Access Points")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListRegionalBuckets, "list-regional-buckets", "", false, "List Regional Buckets")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListStorageLensConfigurations, "list-storage-lens-configurations", "", false, "List Storage Lens Configurations")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListStorageLensGroups, "list-storage-lens-groups", "", false, "List Storage Lens Groups")
	_s3controlCmd.Flags().BoolVarP(&_s3controlListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutAccessGrantsInstanceResourcePolicy, "put-access-grants-instance-resource-policy", "", false, "Put Access Grants Instance Resource Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutAccessPointConfigurationForObjectLambda, "put-access-point-configuration-for-object-lambda", "", false, "Put Access Point Configuration For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutAccessPointPolicy, "put-access-point-policy", "", false, "Put Access Point Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutAccessPointPolicyForObjectLambda, "put-access-point-policy-for-object-lambda", "", false, "Put Access Point Policy For Object Lambda")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutAccessPointScope, "put-access-point-scope", "", false, "Put Access Point Scope")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutBucketLifecycleConfiguration, "put-bucket-lifecycle-configuration", "", false, "Put Bucket Lifecycle Configuration")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutBucketPolicy, "put-bucket-policy", "", false, "Put Bucket Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutBucketReplication, "put-bucket-replication", "", false, "Put Bucket Replication")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutBucketTagging, "put-bucket-tagging", "", false, "Put Bucket Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutBucketVersioning, "put-bucket-versioning", "", false, "Put Bucket Versioning")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutJobTagging, "put-job-tagging", "", false, "Put Job Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutMultiRegionAccessPointPolicy, "put-multi-region-access-point-policy", "", false, "Put Multi Region Access Point Policy")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutPublicAccessBlock, "put-public-access-block", "", false, "Put Public Access Block")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutStorageLensConfiguration, "put-storage-lens-configuration", "", false, "Put Storage Lens Configuration")
	_s3controlCmd.Flags().BoolVarP(&_s3controlPutStorageLensConfigurationTagging, "put-storage-lens-configuration-tagging", "", false, "Put Storage Lens Configuration Tagging")
	_s3controlCmd.Flags().BoolVarP(&_s3controlSubmitMultiRegionAccessPointRoutes, "submit-multi-region-access-point-routes", "", false, "Submit Multi Region Access Point Routes")
	_s3controlCmd.Flags().BoolVarP(&_s3controlTagResource, "tag-resource", "", false, "Tag Resource")
	_s3controlCmd.Flags().BoolVarP(&_s3controlUntagResource, "untag-resource", "", false, "Untag Resource")
	_s3controlCmd.Flags().BoolVarP(&_s3controlUpdateAccessGrantsLocation, "update-access-grants-location", "", false, "Update Access Grants Location")
	_s3controlCmd.Flags().BoolVarP(&_s3controlUpdateJobPriority, "update-job-priority", "", false, "Update Job Priority")
	_s3controlCmd.Flags().BoolVarP(&_s3controlUpdateJobStatus, "update-job-status", "", false, "Update Job Status")
	_s3controlCmd.Flags().BoolVarP(&_s3controlUpdateStorageLensGroup, "update-storage-lens-group", "", false, "Update Storage Lens Group")

}
