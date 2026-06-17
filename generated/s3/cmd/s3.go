package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// s3Cmd represents the s3 command
var _s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "AWS s3 CLI",
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
		client := s3.NewFromConfig(cfg)
		if _s3AbortMultipartUpload {
			s3_AbortMultipartUpload(cfg, client)
			return
		}
		if _s3CompleteMultipartUpload {
			s3_CompleteMultipartUpload(cfg, client)
			return
		}
		if _s3CopyObject {
			s3_CopyObject(cfg, client)
			return
		}
		if _s3CreateBucket {
			s3_CreateBucket(cfg, client)
			return
		}
		if _s3CreateBucketMetadataConfiguration {
			s3_CreateBucketMetadataConfiguration(cfg, client)
			return
		}
		if _s3CreateBucketMetadataTableConfiguration {
			s3_CreateBucketMetadataTableConfiguration(cfg, client)
			return
		}
		if _s3CreateMultipartUpload {
			s3_CreateMultipartUpload(cfg, client)
			return
		}
		if _s3CreateSession {
			s3_CreateSession(cfg, client)
			return
		}
		if _s3DeleteBucket {
			s3_DeleteBucket(cfg, client)
			return
		}
		if _s3DeleteBucketAnalyticsConfiguration {
			s3_DeleteBucketAnalyticsConfiguration(cfg, client)
			return
		}
		if _s3DeleteBucketCors {
			s3_DeleteBucketCors(cfg, client)
			return
		}
		if _s3DeleteBucketEncryption {
			s3_DeleteBucketEncryption(cfg, client)
			return
		}
		if _s3DeleteBucketIntelligentTieringConfiguration {
			s3_DeleteBucketIntelligentTieringConfiguration(cfg, client)
			return
		}
		if _s3DeleteBucketInventoryConfiguration {
			s3_DeleteBucketInventoryConfiguration(cfg, client)
			return
		}
		if _s3DeleteBucketLifecycle {
			s3_DeleteBucketLifecycle(cfg, client)
			return
		}
		if _s3DeleteBucketMetadataConfiguration {
			s3_DeleteBucketMetadataConfiguration(cfg, client)
			return
		}
		if _s3DeleteBucketMetadataTableConfiguration {
			s3_DeleteBucketMetadataTableConfiguration(cfg, client)
			return
		}
		if _s3DeleteBucketMetricsConfiguration {
			s3_DeleteBucketMetricsConfiguration(cfg, client)
			return
		}
		if _s3DeleteBucketOwnershipControls {
			s3_DeleteBucketOwnershipControls(cfg, client)
			return
		}
		if _s3DeleteBucketPolicy {
			s3_DeleteBucketPolicy(cfg, client)
			return
		}
		if _s3DeleteBucketReplication {
			s3_DeleteBucketReplication(cfg, client)
			return
		}
		if _s3DeleteBucketTagging {
			s3_DeleteBucketTagging(cfg, client)
			return
		}
		if _s3DeleteBucketWebsite {
			s3_DeleteBucketWebsite(cfg, client)
			return
		}
		if _s3DeleteObject {
			s3_DeleteObject(cfg, client)
			return
		}
		if _s3DeleteObjectTagging {
			s3_DeleteObjectTagging(cfg, client)
			return
		}
		if _s3DeleteObjects {
			s3_DeleteObjects(cfg, client)
			return
		}
		if _s3DeletePublicAccessBlock {
			s3_DeletePublicAccessBlock(cfg, client)
			return
		}
		if _s3GetBucketAbac {
			s3_GetBucketAbac(cfg, client)
			return
		}
		if _s3GetBucketAccelerateConfiguration {
			s3_GetBucketAccelerateConfiguration(cfg, client)
			return
		}
		if _s3GetBucketAcl {
			s3_GetBucketAcl(cfg, client)
			return
		}
		if _s3GetBucketAnalyticsConfiguration {
			s3_GetBucketAnalyticsConfiguration(cfg, client)
			return
		}
		if _s3GetBucketCors {
			s3_GetBucketCors(cfg, client)
			return
		}
		if _s3GetBucketEncryption {
			s3_GetBucketEncryption(cfg, client)
			return
		}
		if _s3GetBucketIntelligentTieringConfiguration {
			s3_GetBucketIntelligentTieringConfiguration(cfg, client)
			return
		}
		if _s3GetBucketInventoryConfiguration {
			s3_GetBucketInventoryConfiguration(cfg, client)
			return
		}
		if _s3GetBucketLifecycleConfiguration {
			s3_GetBucketLifecycleConfiguration(cfg, client)
			return
		}
		if _s3GetBucketLocation {
			s3_GetBucketLocation(cfg, client)
			return
		}
		if _s3GetBucketLogging {
			s3_GetBucketLogging(cfg, client)
			return
		}
		if _s3GetBucketMetadataConfiguration {
			s3_GetBucketMetadataConfiguration(cfg, client)
			return
		}
		if _s3GetBucketMetadataTableConfiguration {
			s3_GetBucketMetadataTableConfiguration(cfg, client)
			return
		}
		if _s3GetBucketMetricsConfiguration {
			s3_GetBucketMetricsConfiguration(cfg, client)
			return
		}
		if _s3GetBucketNotificationConfiguration {
			s3_GetBucketNotificationConfiguration(cfg, client)
			return
		}
		if _s3GetBucketOwnershipControls {
			s3_GetBucketOwnershipControls(cfg, client)
			return
		}
		if _s3GetBucketPolicy {
			s3_GetBucketPolicy(cfg, client)
			return
		}
		if _s3GetBucketPolicyStatus {
			s3_GetBucketPolicyStatus(cfg, client)
			return
		}
		if _s3GetBucketReplication {
			s3_GetBucketReplication(cfg, client)
			return
		}
		if _s3GetBucketRequestPayment {
			s3_GetBucketRequestPayment(cfg, client)
			return
		}
		if _s3GetBucketTagging {
			s3_GetBucketTagging(cfg, client)
			return
		}
		if _s3GetBucketVersioning {
			s3_GetBucketVersioning(cfg, client)
			return
		}
		if _s3GetBucketWebsite {
			s3_GetBucketWebsite(cfg, client)
			return
		}
		if _s3GetObject {
			s3_GetObject(cfg, client)
			return
		}
		if _s3GetObjectAcl {
			s3_GetObjectAcl(cfg, client)
			return
		}
		if _s3GetObjectAttributes {
			s3_GetObjectAttributes(cfg, client)
			return
		}
		if _s3GetObjectLegalHold {
			s3_GetObjectLegalHold(cfg, client)
			return
		}
		if _s3GetObjectLockConfiguration {
			s3_GetObjectLockConfiguration(cfg, client)
			return
		}
		if _s3GetObjectRetention {
			s3_GetObjectRetention(cfg, client)
			return
		}
		if _s3GetObjectTagging {
			s3_GetObjectTagging(cfg, client)
			return
		}
		if _s3GetObjectTorrent {
			s3_GetObjectTorrent(cfg, client)
			return
		}
		if _s3GetPublicAccessBlock {
			s3_GetPublicAccessBlock(cfg, client)
			return
		}
		if _s3HeadBucket {
			s3_HeadBucket(cfg, client)
			return
		}
		if _s3HeadObject {
			s3_HeadObject(cfg, client)
			return
		}
		if _s3ListBucketAnalyticsConfigurations {
			s3_ListBucketAnalyticsConfigurations(cfg, client)
			return
		}
		if _s3ListBucketIntelligentTieringConfigurations {
			s3_ListBucketIntelligentTieringConfigurations(cfg, client)
			return
		}
		if _s3ListBucketInventoryConfigurations {
			s3_ListBucketInventoryConfigurations(cfg, client)
			return
		}
		if _s3ListBucketMetricsConfigurations {
			s3_ListBucketMetricsConfigurations(cfg, client)
			return
		}
		if _s3ListBuckets {
			s3_ListBuckets(cfg, client)
			return
		}
		if _s3ListDirectoryBuckets {
			s3_ListDirectoryBuckets(cfg, client)
			return
		}
		if _s3ListMultipartUploads {
			s3_ListMultipartUploads(cfg, client)
			return
		}
		if _s3ListObjectVersions {
			s3_ListObjectVersions(cfg, client)
			return
		}
		if _s3ListObjects {
			s3_ListObjects(cfg, client)
			return
		}
		if _s3ListObjectsV2 {
			s3_ListObjectsV2(cfg, client)
			return
		}
		if _s3ListParts {
			s3_ListParts(cfg, client)
			return
		}
		if _s3PutBucketAbac {
			s3_PutBucketAbac(cfg, client)
			return
		}
		if _s3PutBucketAccelerateConfiguration {
			s3_PutBucketAccelerateConfiguration(cfg, client)
			return
		}
		if _s3PutBucketAcl {
			s3_PutBucketAcl(cfg, client)
			return
		}
		if _s3PutBucketAnalyticsConfiguration {
			s3_PutBucketAnalyticsConfiguration(cfg, client)
			return
		}
		if _s3PutBucketCors {
			s3_PutBucketCors(cfg, client)
			return
		}
		if _s3PutBucketEncryption {
			s3_PutBucketEncryption(cfg, client)
			return
		}
		if _s3PutBucketIntelligentTieringConfiguration {
			s3_PutBucketIntelligentTieringConfiguration(cfg, client)
			return
		}
		if _s3PutBucketInventoryConfiguration {
			s3_PutBucketInventoryConfiguration(cfg, client)
			return
		}
		if _s3PutBucketLifecycleConfiguration {
			s3_PutBucketLifecycleConfiguration(cfg, client)
			return
		}
		if _s3PutBucketLogging {
			s3_PutBucketLogging(cfg, client)
			return
		}
		if _s3PutBucketMetricsConfiguration {
			s3_PutBucketMetricsConfiguration(cfg, client)
			return
		}
		if _s3PutBucketNotificationConfiguration {
			s3_PutBucketNotificationConfiguration(cfg, client)
			return
		}
		if _s3PutBucketOwnershipControls {
			s3_PutBucketOwnershipControls(cfg, client)
			return
		}
		if _s3PutBucketPolicy {
			s3_PutBucketPolicy(cfg, client)
			return
		}
		if _s3PutBucketReplication {
			s3_PutBucketReplication(cfg, client)
			return
		}
		if _s3PutBucketRequestPayment {
			s3_PutBucketRequestPayment(cfg, client)
			return
		}
		if _s3PutBucketTagging {
			s3_PutBucketTagging(cfg, client)
			return
		}
		if _s3PutBucketVersioning {
			s3_PutBucketVersioning(cfg, client)
			return
		}
		if _s3PutBucketWebsite {
			s3_PutBucketWebsite(cfg, client)
			return
		}
		if _s3PutObject {
			s3_PutObject(cfg, client)
			return
		}
		if _s3PutObjectAcl {
			s3_PutObjectAcl(cfg, client)
			return
		}
		if _s3PutObjectLegalHold {
			s3_PutObjectLegalHold(cfg, client)
			return
		}
		if _s3PutObjectLockConfiguration {
			s3_PutObjectLockConfiguration(cfg, client)
			return
		}
		if _s3PutObjectRetention {
			s3_PutObjectRetention(cfg, client)
			return
		}
		if _s3PutObjectTagging {
			s3_PutObjectTagging(cfg, client)
			return
		}
		if _s3PutPublicAccessBlock {
			s3_PutPublicAccessBlock(cfg, client)
			return
		}
		if _s3RenameObject {
			s3_RenameObject(cfg, client)
			return
		}
		if _s3RestoreObject {
			s3_RestoreObject(cfg, client)
			return
		}
		if _s3SelectObjectContent {
			s3_SelectObjectContent(cfg, client)
			return
		}
		if _s3UpdateBucketMetadataInventoryTableConfiguration {
			s3_UpdateBucketMetadataInventoryTableConfiguration(cfg, client)
			return
		}
		if _s3UpdateBucketMetadataJournalTableConfiguration {
			s3_UpdateBucketMetadataJournalTableConfiguration(cfg, client)
			return
		}
		if _s3UpdateObjectEncryption {
			s3_UpdateObjectEncryption(cfg, client)
			return
		}
		if _s3UploadPart {
			s3_UploadPart(cfg, client)
			return
		}
		if _s3UploadPartCopy {
			s3_UploadPartCopy(cfg, client)
			return
		}
		if _s3WriteGetObjectResponse {
			s3_WriteGetObjectResponse(cfg, client)
			return
		}

	},
}

var (
	_s3AbortMultipartUpload                            bool
	_s3CompleteMultipartUpload                         bool
	_s3CopyObject                                      bool
	_s3CreateBucket                                    bool
	_s3CreateBucketMetadataConfiguration               bool
	_s3CreateBucketMetadataTableConfiguration          bool
	_s3CreateMultipartUpload                           bool
	_s3CreateSession                                   bool
	_s3DeleteBucket                                    bool
	_s3DeleteBucketAnalyticsConfiguration              bool
	_s3DeleteBucketCors                                bool
	_s3DeleteBucketEncryption                          bool
	_s3DeleteBucketIntelligentTieringConfiguration     bool
	_s3DeleteBucketInventoryConfiguration              bool
	_s3DeleteBucketLifecycle                           bool
	_s3DeleteBucketMetadataConfiguration               bool
	_s3DeleteBucketMetadataTableConfiguration          bool
	_s3DeleteBucketMetricsConfiguration                bool
	_s3DeleteBucketOwnershipControls                   bool
	_s3DeleteBucketPolicy                              bool
	_s3DeleteBucketReplication                         bool
	_s3DeleteBucketTagging                             bool
	_s3DeleteBucketWebsite                             bool
	_s3DeleteObject                                    bool
	_s3DeleteObjectTagging                             bool
	_s3DeleteObjects                                   bool
	_s3DeletePublicAccessBlock                         bool
	_s3GetBucketAbac                                   bool
	_s3GetBucketAccelerateConfiguration                bool
	_s3GetBucketAcl                                    bool
	_s3GetBucketAnalyticsConfiguration                 bool
	_s3GetBucketCors                                   bool
	_s3GetBucketEncryption                             bool
	_s3GetBucketIntelligentTieringConfiguration        bool
	_s3GetBucketInventoryConfiguration                 bool
	_s3GetBucketLifecycleConfiguration                 bool
	_s3GetBucketLocation                               bool
	_s3GetBucketLogging                                bool
	_s3GetBucketMetadataConfiguration                  bool
	_s3GetBucketMetadataTableConfiguration             bool
	_s3GetBucketMetricsConfiguration                   bool
	_s3GetBucketNotificationConfiguration              bool
	_s3GetBucketOwnershipControls                      bool
	_s3GetBucketPolicy                                 bool
	_s3GetBucketPolicyStatus                           bool
	_s3GetBucketReplication                            bool
	_s3GetBucketRequestPayment                         bool
	_s3GetBucketTagging                                bool
	_s3GetBucketVersioning                             bool
	_s3GetBucketWebsite                                bool
	_s3GetObject                                       bool
	_s3GetObjectAcl                                    bool
	_s3GetObjectAttributes                             bool
	_s3GetObjectLegalHold                              bool
	_s3GetObjectLockConfiguration                      bool
	_s3GetObjectRetention                              bool
	_s3GetObjectTagging                                bool
	_s3GetObjectTorrent                                bool
	_s3GetPublicAccessBlock                            bool
	_s3HeadBucket                                      bool
	_s3HeadObject                                      bool
	_s3ListBucketAnalyticsConfigurations               bool
	_s3ListBucketIntelligentTieringConfigurations      bool
	_s3ListBucketInventoryConfigurations               bool
	_s3ListBucketMetricsConfigurations                 bool
	_s3ListBuckets                                     bool
	_s3ListDirectoryBuckets                            bool
	_s3ListMultipartUploads                            bool
	_s3ListObjectVersions                              bool
	_s3ListObjects                                     bool
	_s3ListObjectsV2                                   bool
	_s3ListParts                                       bool
	_s3PutBucketAbac                                   bool
	_s3PutBucketAccelerateConfiguration                bool
	_s3PutBucketAcl                                    bool
	_s3PutBucketAnalyticsConfiguration                 bool
	_s3PutBucketCors                                   bool
	_s3PutBucketEncryption                             bool
	_s3PutBucketIntelligentTieringConfiguration        bool
	_s3PutBucketInventoryConfiguration                 bool
	_s3PutBucketLifecycleConfiguration                 bool
	_s3PutBucketLogging                                bool
	_s3PutBucketMetricsConfiguration                   bool
	_s3PutBucketNotificationConfiguration              bool
	_s3PutBucketOwnershipControls                      bool
	_s3PutBucketPolicy                                 bool
	_s3PutBucketReplication                            bool
	_s3PutBucketRequestPayment                         bool
	_s3PutBucketTagging                                bool
	_s3PutBucketVersioning                             bool
	_s3PutBucketWebsite                                bool
	_s3PutObject                                       bool
	_s3PutObjectAcl                                    bool
	_s3PutObjectLegalHold                              bool
	_s3PutObjectLockConfiguration                      bool
	_s3PutObjectRetention                              bool
	_s3PutObjectTagging                                bool
	_s3PutPublicAccessBlock                            bool
	_s3RenameObject                                    bool
	_s3RestoreObject                                   bool
	_s3SelectObjectContent                             bool
	_s3UpdateBucketMetadataInventoryTableConfiguration bool
	_s3UpdateBucketMetadataJournalTableConfiguration   bool
	_s3UpdateObjectEncryption                          bool
	_s3UploadPart                                      bool
	_s3UploadPartCopy                                  bool
	_s3WriteGetObjectResponse                          bool

	_s3AbacStatus                         string
	_s3AccelerateConfiguration            string
	_s3AcceptRanges                       string
	_s3AccessControlPolicy                string
	_s3ACL                                string
	_s3AnalyticsConfiguration             string
	_s3Body                               string
	_s3Bucket                             string
	_s3BucketKeyEnabled                   string
	_s3BucketLoggingStatus                string
	_s3BucketRegion                       string
	_s3BypassGovernanceRetention          string
	_s3CacheControl                       string
	_s3ChecksumAlgorithm                  string
	_s3ChecksumCRC32                      string
	_s3ChecksumCRC32C                     string
	_s3ChecksumCRC64NVME                  string
	_s3ChecksumMode                       string
	_s3ChecksumSHA1                       string
	_s3ChecksumSHA256                     string
	_s3ChecksumType                       string
	_s3ClientToken                        string
	_s3ConfirmRemoveSelfBucketAccess      string
	_s3ContentDisposition                 string
	_s3ContentEncoding                    string
	_s3ContentLanguage                    string
	_s3ContentLength                      string
	_s3ContentMD5                         string
	_s3ContentRange                       string
	_s3ContentType                        string
	_s3ContinuationToken                  string
	_s3CopySource                         string
	_s3CopySourceIfMatch                  string
	_s3CopySourceIfModifiedSince          string
	_s3CopySourceIfNoneMatch              string
	_s3CopySourceIfUnmodifiedSince        string
	_s3CopySourceRange                    string
	_s3CopySourceSSECustomerAlgorithm     string
	_s3CopySourceSSECustomerKey           string
	_s3CopySourceSSECustomerKeyMD5        string
	_s3CORSConfiguration                  string
	_s3CreateBucketConfiguration          string
	_s3Delete                             string
	_s3DeleteMarker                       string
	_s3Delimiter                          string
	_s3DestinationIfMatch                 string
	_s3DestinationIfModifiedSince         string
	_s3DestinationIfNoneMatch             string
	_s3DestinationIfUnmodifiedSince       string
	_s3EncodingType                       string
	_s3ErrorCode                          string
	_s3ErrorMessage                       string
	_s3ETag                               string
	_s3ExpectedBucketOwner                string
	_s3ExpectedSourceBucketOwner          string
	_s3Expiration                         string
	_s3Expires                            string
	_s3Expression                         string
	_s3ExpressionType                     string
	_s3FetchOwner                         string
	_s3GrantFullControl                   string
	_s3GrantRead                          string
	_s3GrantReadACP                       string
	_s3GrantWrite                         string
	_s3GrantWriteACP                      string
	_s3Id                                 string
	_s3IfMatch                            string
	_s3IfMatchInitiatedTime               string
	_s3IfMatchLastModifiedTime            string
	_s3IfMatchSize                        string
	_s3IfModifiedSince                    string
	_s3IfNoneMatch                        string
	_s3IfUnmodifiedSince                  string
	_s3InputSerialization                 string
	_s3IntelligentTieringConfiguration    string
	_s3InventoryConfiguration             string
	_s3InventoryTableConfiguration        string
	_s3JournalTableConfiguration          string
	_s3Key                                string
	_s3KeyMarker                          string
	_s3LastModified                       string
	_s3LegalHold                          string
	_s3LifecycleConfiguration             string
	_s3Marker                             string
	_s3MaxBuckets                         string
	_s3MaxDirectoryBuckets                string
	_s3MaxKeys                            string
	_s3MaxParts                           string
	_s3MaxUploads                         string
	_s3Metadata                           string
	_s3MetadataConfiguration              string
	_s3MetadataDirective                  string
	_s3MetadataTableConfiguration         string
	_s3MetricsConfiguration               string
	_s3MFA                                string
	_s3MissingMeta                        string
	_s3MpuObjectSize                      string
	_s3MultipartUpload                    string
	_s3NotificationConfiguration          string
	_s3ObjectAttributes                   string
	_s3ObjectEncryption                   string
	_s3ObjectLockConfiguration            string
	_s3ObjectLockEnabledForBucket         string
	_s3ObjectLockLegalHoldStatus          string
	_s3ObjectLockMode                     string
	_s3ObjectLockRetainUntilDate          string
	_s3ObjectOwnership                    string
	_s3OptionalObjectAttributes           string
	_s3OutputSerialization                string
	_s3OwnershipControls                  string
	_s3PartNumber                         string
	_s3PartNumberMarker                   string
	_s3PartsCount                         string
	_s3Policy                             string
	_s3Prefix                             string
	_s3PublicAccessBlockConfiguration     string
	_s3Range                              string
	_s3RenameSource                       string
	_s3ReplicationConfiguration           string
	_s3ReplicationStatus                  string
	_s3RequestCharged                     string
	_s3RequestPayer                       string
	_s3RequestPaymentConfiguration        string
	_s3RequestProgress                    string
	_s3RequestRoute                       string
	_s3RequestToken                       string
	_s3ResponseCacheControl               string
	_s3ResponseContentDisposition         string
	_s3ResponseContentEncoding            string
	_s3ResponseContentLanguage            string
	_s3ResponseContentType                string
	_s3ResponseExpires                    string
	_s3Restore                            string
	_s3RestoreRequest                     string
	_s3Retention                          string
	_s3ScanRange                          string
	_s3ServerSideEncryption               string
	_s3ServerSideEncryptionConfiguration  string
	_s3SessionMode                        string
	_s3SkipDestinationValidation          string
	_s3SourceIfMatch                      string
	_s3SourceIfModifiedSince              string
	_s3SourceIfNoneMatch                  string
	_s3SourceIfUnmodifiedSince            string
	_s3SSECustomerAlgorithm               string
	_s3SSECustomerKey                     string
	_s3SSECustomerKeyMD5                  string
	_s3SSEKMSEncryptionContext            string
	_s3SSEKMSKeyId                        string
	_s3StartAfter                         string
	_s3StatusCode                         string
	_s3StorageClass                       string
	_s3TagCount                           string
	_s3Tagging                            string
	_s3TaggingDirective                   string
	_s3Token                              string
	_s3TransitionDefaultMinimumObjectSize string
	_s3UploadId                           string
	_s3UploadIdMarker                     string
	_s3VersionId                          string
	_s3VersionIdMarker                    string
	_s3VersioningConfiguration            string
	_s3WebsiteConfiguration               string
	_s3WebsiteRedirectLocation            string
	_s3WriteOffsetBytes                   string
)

// This operation aborts a multipart upload. After a multipart upload is aborted,
// no additional parts can be uploaded using that upload ID. The storage consumed
// by any previously uploaded parts will be freed. However, if any part uploads are
// currently in progress, those part uploads might or might not succeed. As a
// result, it might be necessary to abort a given multipart upload multiple times
// in order to completely free all storage consumed by all parts.
//
// To verify that all parts have been removed and prevent getting charged for the
// part storage, you should call the [ListParts]API operation and ensure that the parts list
// is empty.
//
// - Directory buckets - If multipart uploads in a directory bucket are in
// progress, you can't delete the bucket until all the in-progress multipart
// uploads are aborted or completed. To delete these in-progress multipart uploads,
// use the ListMultipartUploads operation to list the in-progress multipart
// uploads in the bucket and use the AbortMultipartUpload operation to abort all
// the in-progress multipart uploads.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// # Permissions
//
// - General purpose bucket permissions - For information about permissions
// required to use the multipart upload, see [Multipart Upload and Permissions]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to AbortMultipartUpload :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart Upload and Permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
func s3_AbortMultipartUpload(cfg aws.Config, client *s3.Client) {
	input := &s3.AbortMultipartUploadInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3UploadId) > 0 {
		input.UploadId = aws.String(_s3UploadId)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3IfMatchInitiatedTime) > 0 {
		if err := assignInputField(input, "IfMatchInitiatedTime", _s3IfMatchInitiatedTime); err != nil {
			log.Errorf("invalid --if-match-initiated-time: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}

	if resp, err := client.AbortMultipartUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Completes a multipart upload by assembling previously uploaded parts.
// You first initiate the multipart upload and then upload all parts using the [UploadPart]
// operation or the [UploadPartCopy]operation. After successfully uploading all relevant parts of
// an upload, you call this CompleteMultipartUpload operation to complete the
// upload. Upon receiving this request, Amazon S3 concatenates all the parts in
// ascending order by part number to create a new object. In the
// CompleteMultipartUpload request, you must provide the parts list and ensure that
// the parts list is complete. The CompleteMultipartUpload API operation
// concatenates the parts that you provide in the list. For each part in the list,
// you must provide the PartNumber value and the ETag value that are returned
// after that part was uploaded.
//
// The processing of a CompleteMultipartUpload request could take several minutes
// to finalize. After Amazon S3 begins processing the request, it sends an HTTP
// response header that specifies a 200 OK response. While processing is in
// progress, Amazon S3 periodically sends white space characters to keep the
// connection from timing out. A request could fail after the initial 200 OK
// response has been sent. This means that a 200 OK response can contain either a
// success or an error. The error response might be embedded in the 200 OK
// response. If you call this API operation directly, make sure to design your
// application to parse the contents of the response and handle it appropriately.
// If you use Amazon Web Services SDKs, SDKs handle this condition. The SDKs detect
// the embedded error and apply error handling per your configuration settings
// (including automatically retrying the request as appropriate). If the condition
// persists, the SDKs throw an exception (or, for the SDKs that don't use
// exceptions, they return an error).
//
// Note that if CompleteMultipartUpload fails, applications should be prepared to
// retry any failed requests (including 500 error responses). For more information,
// see [Amazon S3 Error Best Practices].
//
// You can't use Content-Type: application/x-www-form-urlencoded for the
// CompleteMultipartUpload requests. Also, if you don't provide a Content-Type
// header, CompleteMultipartUpload can still return a 200 OK response.
//
// For more information about multipart uploads, see [Uploading Objects Using Multipart Upload] in the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - For information about permissions
// required to use the multipart upload API, see [Multipart Upload and Permissions]in the Amazon S3 User Guide.
//
// If you provide an [additional checksum value]in your MultipartUpload requests and the object is encrypted
//
// with Key Management Service, you must have permission to use the kms:Decrypt
// action for the CompleteMultipartUpload request to succeed.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// # Special errors
//
// - Error Code: EntityTooSmall
//
// - Description: Your proposed upload is smaller than the minimum allowed
// object size. Each part must be at least 5 MB in size, except the last part.
//
// - HTTP Status Code: 400 Bad Request
//
// - Error Code: InvalidPart
//
// - Description: One or more of the specified parts could not be found. The
// part might not have been uploaded, or the specified ETag might not have matched
// the uploaded part's ETag.
//
// - HTTP Status Code: 400 Bad Request
//
// - Error Code: InvalidPartOrder
//
// - Description: The list of parts was not in ascending order. The parts list
// must be specified in order by part number.
//
// - HTTP Status Code: 400 Bad Request
//
// - Error Code: NoSuchUpload
//
// - Description: The specified multipart upload does not exist. The upload ID
// might be invalid, or the multipart upload might have been aborted or completed.
//
// - HTTP Status Code: 404 Not Found
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to CompleteMultipartUpload :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [AbortMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Amazon S3 Error Best Practices]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ErrorBestPractices.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [additional checksum value]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_Checksum.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Multipart Upload and Permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_CompleteMultipartUpload(cfg aws.Config, client *s3.Client) {
	input := &s3.CompleteMultipartUploadInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3UploadId) > 0 {
		input.UploadId = aws.String(_s3UploadId)
	}
	if len(_s3ChecksumCRC32) > 0 {
		input.ChecksumCRC32 = aws.String(_s3ChecksumCRC32)
	}
	if len(_s3ChecksumCRC32C) > 0 {
		input.ChecksumCRC32C = aws.String(_s3ChecksumCRC32C)
	}
	if len(_s3ChecksumCRC64NVME) > 0 {
		input.ChecksumCRC64NVME = aws.String(_s3ChecksumCRC64NVME)
	}
	if len(_s3ChecksumSHA1) > 0 {
		input.ChecksumSHA1 = aws.String(_s3ChecksumSHA1)
	}
	if len(_s3ChecksumSHA256) > 0 {
		input.ChecksumSHA256 = aws.String(_s3ChecksumSHA256)
	}
	if len(_s3ChecksumType) > 0 {
		if err := assignInputField(input, "ChecksumType", _s3ChecksumType); err != nil {
			log.Errorf("invalid --checksum-type: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3IfMatch) > 0 {
		input.IfMatch = aws.String(_s3IfMatch)
	}
	if len(_s3IfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_s3IfNoneMatch)
	}
	if len(_s3MpuObjectSize) > 0 {
		if err := assignInputField(input, "MpuObjectSize", _s3MpuObjectSize); err != nil {
			log.Errorf("invalid --mpu-object-size: %s", err.Error())
			return
		}
	}
	if len(_s3MultipartUpload) > 0 {
		if err := assignInputField(input, "MultipartUpload", _s3MultipartUpload); err != nil {
			log.Errorf("invalid --multipart-upload: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}

	if resp, err := client.CompleteMultipartUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a copy of an object that is already stored in Amazon S3.
// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// You can store individual objects of up to 50 TB in Amazon S3. You create a copy
// of your object up to 5 GB in size in a single atomic action using this API.
// However, to copy an object greater than 5 GB, you must use the multipart upload
// Upload Part - Copy (UploadPartCopy) API. For more information, see [Copy Object Using the REST Multipart Upload API].
//
// You can copy individual objects between general purpose buckets, between
// directory buckets, and between general purpose buckets and directory buckets.
//
// - Amazon S3 supports copy operations using Multi-Region Access Points only as
// a destination when using the Multi-Region Access Point ARN.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// - VPC endpoints don't support cross-Region requests (including copies). If
// you're using VPC endpoints, your source and destination buckets should be in the
// same Amazon Web Services Region as your VPC endpoint.
//
// Both the Region that you want to copy the object from and the Region that you
// want to copy the object to must be enabled for your account. For more
// information about how to enable a Region for your account, see [Enable or disable a Region for standalone accounts]in the Amazon
// Web Services Account Management Guide.
//
// Amazon S3 transfer acceleration does not support cross-Region copies. If you
// request a cross-Region copy using a transfer acceleration endpoint, you get a
// 400 Bad Request error. For more information, see [Transfer Acceleration].
//
// Authentication and authorization All CopyObject requests must be authenticated
// and signed by using IAM credentials (access key ID and secret access key for the
// IAM identities). All headers with the x-amz- prefix, including x-amz-copy-source
// , must be signed. For more information, see [REST Authentication].
//
// Directory buckets - You must use the IAM credentials to authenticate and
// authorize your access to the CopyObject API operation, instead of using the
// temporary security credentials through the CreateSession API operation.
//
// Amazon Web Services CLI or SDKs handles authentication and authorization on
// your behalf.
//
// Permissions You must have read access to the source object and write access to
// the destination bucket.
//
// - General purpose bucket permissions - You must have permissions in an IAM
// policy based on the source and destination bucket types in a CopyObject
// operation.
//
// - If the source object is in a general purpose bucket, you must have
// s3:GetObject permission to read the source object that is being copied.
//
// - If the destination bucket is a general purpose bucket, you must have
// s3:PutObject permission to write the object copy to the destination bucket.
//
// - Directory bucket permissions - You must have permissions in a bucket policy
// or an IAM identity-based policy based on the source and destination bucket types
// in a CopyObject operation.
//
// - If the source object that you want to copy is in a directory bucket, you
// must have the s3express:CreateSession permission in the Action element of a
// policy to read the object. By default, the session is in the ReadWrite mode.
// If you want to restrict the access, you can explicitly set the
// s3express:SessionMode condition key to ReadOnly on the copy source bucket.
//
// - If the copy destination is a directory bucket, you must have the
// s3express:CreateSession permission in the Action element of a policy to write
// the object to the destination. The s3express:SessionMode condition key can't
// be set to ReadOnly on the copy destination bucket.
//
// # If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// For example policies, see [Example bucket policies for S3 Express One Zone]and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]in the Amazon S3 User Guide.
//
// Response and special errors When the request is an HTTP 1.1 request, the
// response is chunk encoded. When the request is not an HTTP 1.1 request, the
// response would not contain the Content-Length . You always need to read the
// entire response body to check if the copy succeeds.
//
// - If the copy is successful, you receive a response with information about
// the copied object.
//
// - A copy request might return an error when Amazon S3 receives the copy
// request or while Amazon S3 is copying the files. A 200 OK response can contain
// either a success or an error.
//
// - If the error occurs before the copy action starts, you receive a standard
// Amazon S3 error.
//
// - If the error occurs during the copy operation, the error response is
// embedded in the 200 OK response. For example, in a cross-region copy, you may
// encounter throttling and receive a 200 OK response. For more information, see [Resolve the Error 200 response when copying objects to Amazon S3]
// . The 200 OK status code means the copy was accepted, but it doesn't mean the
// copy is complete. Another example is when you disconnect from Amazon S3 before
// the copy is complete, Amazon S3 might cancel the copy and you may receive a
// 200 OK response. You must stay connected to Amazon S3 until the entire
// response is successfully received and processed.
//
// # If you call this API operation directly, make sure to design your application
//
// to parse the content of the response and handle it appropriately. If you use
// Amazon Web Services SDKs, SDKs handle this condition. The SDKs detect the
// embedded error and apply error handling per your configuration settings
// (including automatically retrying the request as appropriate). If the condition
// persists, the SDKs throw an exception (or, for the SDKs that don't use
// exceptions, they return an error).
//
// Charge The copy request charge is based on the storage class and Region that
// you specify for the destination object. The request can also result in a data
// retrieval charge for the source if the source storage class bills for data
// retrieval. If the copy source is in a different region, the data transfer is
// billed to the copy source account. For pricing information, see [Amazon S3 pricing].
//
// # HTTP Host header syntax
//
// - Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// - Amazon S3 on Outposts - When you use this action with S3 on Outposts
// through the REST API, you must direct requests to the S3 on Outposts hostname.
// The S3 on Outposts hostname takes the form
// AccessPointName-AccountId.outpostID.s3-outposts.Region.amazonaws.com . The
// hostname isn't required when you use the Amazon Web Services CLI or SDKs.
//
// The following operations are related to CopyObject :
//
// [PutObject]
//
// [GetObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [Resolve the Error 200 response when copying objects to Amazon S3]: https://repost.aws/knowledge-center/s3-resolve-200-internalerror
// [Copy Object Using the REST Multipart Upload API]: https://docs.aws.amazon.com/AmazonS3/latest/dev/CopyingObjctsUsingRESTMPUapi.html
// [REST Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [Enable or disable a Region for standalone accounts]: https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-regions.html#manage-acct-regions-enable-standalone
// [Transfer Acceleration]: https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon S3 pricing]: http://aws.amazon.com/s3/pricing/
func s3_CopyObject(cfg aws.Config, client *s3.Client) {
	input := &s3.CopyObjectInput{
		// Bucket: *string, // Required
		// CopySource: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3CopySource) > 0 {
		input.CopySource = aws.String(_s3CopySource)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ACL) > 0 {
		if err := assignInputField(input, "ACL", _s3ACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3BucketKeyEnabled) > 0 {
		if err := assignInputField(input, "BucketKeyEnabled", _s3BucketKeyEnabled); err != nil {
			log.Errorf("invalid --bucket-key-enabled: %s", err.Error())
			return
		}
	}
	if len(_s3CacheControl) > 0 {
		input.CacheControl = aws.String(_s3CacheControl)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentDisposition) > 0 {
		input.ContentDisposition = aws.String(_s3ContentDisposition)
	}
	if len(_s3ContentEncoding) > 0 {
		input.ContentEncoding = aws.String(_s3ContentEncoding)
	}
	if len(_s3ContentLanguage) > 0 {
		input.ContentLanguage = aws.String(_s3ContentLanguage)
	}
	if len(_s3ContentType) > 0 {
		input.ContentType = aws.String(_s3ContentType)
	}
	if len(_s3CopySourceIfMatch) > 0 {
		input.CopySourceIfMatch = aws.String(_s3CopySourceIfMatch)
	}
	if len(_s3CopySourceIfModifiedSince) > 0 {
		if err := assignInputField(input, "CopySourceIfModifiedSince", _s3CopySourceIfModifiedSince); err != nil {
			log.Errorf("invalid --copy-source-if-modified-since: %s", err.Error())
			return
		}
	}
	if len(_s3CopySourceIfNoneMatch) > 0 {
		input.CopySourceIfNoneMatch = aws.String(_s3CopySourceIfNoneMatch)
	}
	if len(_s3CopySourceIfUnmodifiedSince) > 0 {
		if err := assignInputField(input, "CopySourceIfUnmodifiedSince", _s3CopySourceIfUnmodifiedSince); err != nil {
			log.Errorf("invalid --copy-source-if-unmodified-since: %s", err.Error())
			return
		}
	}
	if len(_s3CopySourceSSECustomerAlgorithm) > 0 {
		input.CopySourceSSECustomerAlgorithm = aws.String(_s3CopySourceSSECustomerAlgorithm)
	}
	if len(_s3CopySourceSSECustomerKey) > 0 {
		input.CopySourceSSECustomerKey = aws.String(_s3CopySourceSSECustomerKey)
	}
	if len(_s3CopySourceSSECustomerKeyMD5) > 0 {
		input.CopySourceSSECustomerKeyMD5 = aws.String(_s3CopySourceSSECustomerKeyMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3ExpectedSourceBucketOwner) > 0 {
		input.ExpectedSourceBucketOwner = aws.String(_s3ExpectedSourceBucketOwner)
	}
	if len(_s3Expires) > 0 {
		if err := assignInputField(input, "Expires", _s3Expires); err != nil {
			log.Errorf("invalid --expires: %s", err.Error())
			return
		}
	}
	if len(_s3GrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3GrantFullControl)
	}
	if len(_s3GrantRead) > 0 {
		input.GrantRead = aws.String(_s3GrantRead)
	}
	if len(_s3GrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3GrantReadACP)
	}
	if len(_s3GrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3GrantWriteACP)
	}
	if len(_s3IfMatch) > 0 {
		input.IfMatch = aws.String(_s3IfMatch)
	}
	if len(_s3IfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_s3IfNoneMatch)
	}
	if len(_s3Metadata) > 0 {
		if err := assignInputField(input, "Metadata", _s3Metadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_s3MetadataDirective) > 0 {
		if err := assignInputField(input, "MetadataDirective", _s3MetadataDirective); err != nil {
			log.Errorf("invalid --metadata-directive: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockLegalHoldStatus) > 0 {
		if err := assignInputField(input, "ObjectLockLegalHoldStatus", _s3ObjectLockLegalHoldStatus); err != nil {
			log.Errorf("invalid --object-lock-legal-hold-status: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockMode) > 0 {
		if err := assignInputField(input, "ObjectLockMode", _s3ObjectLockMode); err != nil {
			log.Errorf("invalid --object-lock-mode: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockRetainUntilDate) > 0 {
		if err := assignInputField(input, "ObjectLockRetainUntilDate", _s3ObjectLockRetainUntilDate); err != nil {
			log.Errorf("invalid --object-lock-retain-until-date: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3SSEKMSEncryptionContext) > 0 {
		input.SSEKMSEncryptionContext = aws.String(_s3SSEKMSEncryptionContext)
	}
	if len(_s3SSEKMSKeyId) > 0 {
		input.SSEKMSKeyId = aws.String(_s3SSEKMSKeyId)
	}
	if len(_s3ServerSideEncryption) > 0 {
		if err := assignInputField(input, "ServerSideEncryption", _s3ServerSideEncryption); err != nil {
			log.Errorf("invalid --server-side-encryption: %s", err.Error())
			return
		}
	}
	if len(_s3StorageClass) > 0 {
		if err := assignInputField(input, "StorageClass", _s3StorageClass); err != nil {
			log.Errorf("invalid --storage-class: %s", err.Error())
			return
		}
	}
	if len(_s3Tagging) > 0 {
		input.Tagging = aws.String(_s3Tagging)
	}
	if len(_s3TaggingDirective) > 0 {
		if err := assignInputField(input, "TaggingDirective", _s3TaggingDirective); err != nil {
			log.Errorf("invalid --tagging-directive: %s", err.Error())
			return
		}
	}
	if len(_s3WebsiteRedirectLocation) > 0 {
		input.WebsiteRedirectLocation = aws.String(_s3WebsiteRedirectLocation)
	}

	if resp, err := client.CopyObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This action creates an Amazon S3 bucket. To create an Amazon S3 on Outposts
// bucket, see [CreateBucket]CreateBucket .
//
// Creates a new S3 bucket. To create a bucket, you must set up Amazon S3 and have
// a valid Amazon Web Services Access Key ID to authenticate requests. Anonymous
// requests are never allowed to create buckets. By creating the bucket, you become
// the bucket owner.
//
// There are two types of buckets: general purpose buckets and directory buckets.
// For more information about these bucket types, see [Creating, configuring, and working with Amazon S3 buckets]in the Amazon S3 User Guide.
//
// - General purpose buckets - If you send your CreateBucket request to the
// s3.amazonaws.com global endpoint, the request goes to the us-east-1 Region. So
// the signature calculations in Signature Version 4 must use us-east-1 as the
// Region, even if the location constraint in the request specifies another Region
// where the bucket is to be created. If you create a bucket in a Region other than
// US East (N. Virginia), your application must be able to handle 307 redirect. For
// more information, see [Virtual hosting of buckets]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Regional endpoint. These endpoints support path-style
// requests in the format
// https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// # Permissions
//
// - General purpose bucket permissions - In addition to the s3:CreateBucket
// permission, the following permissions are required in a policy when your
// CreateBucket request includes specific headers:
//
// - Access control lists (ACLs) - In your CreateBucket request, if you specify
// an access control list (ACL) and set it to public-read , public-read-write ,
// authenticated-read , or if you explicitly specify any other custom ACLs, both
// s3:CreateBucket and s3:PutBucketAcl permissions are required. In your
// CreateBucket request, if you set the ACL to private , or if you don't specify
// any ACLs, only the s3:CreateBucket permission is required.
//
// - Object Lock - In your CreateBucket request, if you set
// x-amz-bucket-object-lock-enabled to true, the
// s3:PutBucketObjectLockConfiguration and s3:PutBucketVersioning permissions are
// required.
//
// - S3 Object Ownership - If your CreateBucket request includes the
// x-amz-object-ownership header, then the s3:PutBucketOwnershipControls
// permission is required.
//
// # To set an ACL on a bucket as part of a CreateBucket request, you must explicitly
//
// set S3 Object Ownership for the bucket to a different value than the default,
// BucketOwnerEnforced . Additionally, if your desired bucket ACL grants public
// access, you must first create the bucket (without the bucket ACL) and then
// explicitly disable Block Public Access on the bucket before using PutBucketAcl
// to set the ACL. If you try to create a bucket with a public ACL, the request
// will fail.
//
// # For the majority of modern use cases in S3, we recommend that you keep all
//
// Block Public Access settings enabled and keep ACLs disabled. If you would like
// to share data with users outside of your account, you can use bucket policies as
// needed. For more information, see [Controlling ownership of objects and disabling ACLs for your bucket]and [Blocking public access to your Amazon S3 storage]in the Amazon S3 User Guide.
//
// - S3 Block Public Access - If your specific use case requires granting public
// access to your S3 resources, you can disable Block Public Access. Specifically,
// you can create a new bucket with Block Public Access enabled, then separately
// call the [DeletePublicAccessBlock]DeletePublicAccessBlock API. To use this operation, you must have the
// s3:PutBucketPublicAccessBlock permission. For more information about S3 Block
// Public Access, see [Blocking public access to your Amazon S3 storage]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - You must have the s3express:CreateBucket
// permission in an IAM identity-based policy instead of a bucket policy.
// Cross-account access to this API operation isn't supported. This operation can
// only be performed by the Amazon Web Services account that owns the resource. For
// more information about directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the
// Amazon S3 User Guide.
//
// # The permissions for ACLs, Object Lock, S3 Object Ownership, and S3 Block Public
//
// Access are not supported for directory buckets. For directory buckets, all Block
// Public Access settings are enabled at the bucket level and S3 Object Ownership
// is set to Bucket owner enforced (ACLs disabled). These settings can't be
// modified.
//
// # For more information about permissions for creating and working with directory
//
// buckets, see [Directory buckets]in the Amazon S3 User Guide. For more information about
// supported S3 features for directory buckets, see [Features of S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to CreateBucket :
//
// [PutObject]
//
// [DeleteBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Creating, configuring, and working with Amazon S3 buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html
// [Virtual hosting of buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-overview.html
// [Features of S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-one-zone.html#s3-express-features
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [Controlling ownership of objects and disabling ACLs for your bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
// [Blocking public access to your Amazon S3 storage]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html
func s3_CreateBucket(cfg aws.Config, client *s3.Client) {
	input := &s3.CreateBucketInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ACL) > 0 {
		if err := assignInputField(input, "ACL", _s3ACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3CreateBucketConfiguration) > 0 {
		if err := assignInputField(input, "CreateBucketConfiguration", _s3CreateBucketConfiguration); err != nil {
			log.Errorf("invalid --create-bucket-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3GrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3GrantFullControl)
	}
	if len(_s3GrantRead) > 0 {
		input.GrantRead = aws.String(_s3GrantRead)
	}
	if len(_s3GrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3GrantReadACP)
	}
	if len(_s3GrantWrite) > 0 {
		input.GrantWrite = aws.String(_s3GrantWrite)
	}
	if len(_s3GrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3GrantWriteACP)
	}
	if len(_s3ObjectLockEnabledForBucket) > 0 {
		if err := assignInputField(input, "ObjectLockEnabledForBucket", _s3ObjectLockEnabledForBucket); err != nil {
			log.Errorf("invalid --object-lock-enabled-for-bucket: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectOwnership) > 0 {
		if err := assignInputField(input, "ObjectOwnership", _s3ObjectOwnership); err != nil {
			log.Errorf("invalid --object-ownership: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates an S3 Metadata V2 metadata configuration for a general purpose bucket.
// For more information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the following permissions. For
// more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you want to encrypt your metadata tables with server-side encryption with
// Key Management Service (KMS) keys (SSE-KMS), you need additional permissions in
// your KMS key policy. For more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you also want to integrate your table bucket with Amazon Web Services
// analytics services so that you can query your metadata table, you need
// additional permissions. For more information, see [Integrating Amazon S3 Tables with Amazon Web Services analytics services]in the Amazon S3 User Guide.
//
// To query your metadata tables, you need additional permissions. For more
// information, see [Permissions for querying metadata tables]in the Amazon S3 User Guide.
//
// - s3:CreateBucketMetadataTableConfiguration
//
// The IAM policy action name is the same for the V1 and V2 API operations.
//
// - s3tables:CreateTableBucket
//
// - s3tables:CreateNamespace
//
// - s3tables:GetTable
//
// - s3tables:CreateTable
//
// - s3tables:PutTablePolicy
//
// - s3tables:PutTableEncryption
//
// - kms:DescribeKey
//
// The following operations are related to CreateBucketMetadataConfiguration :
//
// [DeleteBucketMetadataConfiguration]
//
// [GetBucketMetadataConfiguration]
//
// [UpdateBucketMetadataInventoryTableConfiguration]
//
// [UpdateBucketMetadataJournalTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [UpdateBucketMetadataJournalTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [Permissions for querying metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-bucket-query-permissions.html
// [UpdateBucketMetadataInventoryTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html
// [DeleteBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html
// [Integrating Amazon S3 Tables with Amazon Web Services analytics services]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-aws.html
func s3_CreateBucketMetadataConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.CreateBucketMetadataConfigurationInput{
		// Bucket: *string, // Required
		// MetadataConfiguration: *types.MetadataConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3MetadataConfiguration) > 0 {
		if err := assignInputField(input, "MetadataConfiguration", _s3MetadataConfiguration); err != nil {
			log.Errorf("invalid --metadata-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.CreateBucketMetadataConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend that you create your S3 Metadata configurations by using the V2 [CreateBucketMetadataConfiguration]
// API operation. We no longer recommend using the V1
// CreateBucketMetadataTableConfiguration API operation.
//
// If you created your S3 Metadata configuration before July 15, 2025, we
// recommend that you delete and re-create your configuration by using [CreateBucketMetadataConfiguration]so that you
// can expire journal table records and create a live inventory table.
//
// Creates a V1 S3 Metadata configuration for a general purpose bucket. For more
// information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the following permissions. For
// more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you want to encrypt your metadata tables with server-side encryption with
// Key Management Service (KMS) keys (SSE-KMS), you need additional permissions.
// For more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you also want to integrate your table bucket with Amazon Web Services
// analytics services so that you can query your metadata table, you need
// additional permissions. For more information, see [Integrating Amazon S3 Tables with Amazon Web Services analytics services]in the Amazon S3 User Guide.
//
// - s3:CreateBucketMetadataTableConfiguration
//
// - s3tables:CreateNamespace
//
// - s3tables:GetTable
//
// - s3tables:CreateTable
//
// - s3tables:PutTablePolicy
//
// The following operations are related to CreateBucketMetadataTableConfiguration :
//
// [DeleteBucketMetadataTableConfiguration]
//
// [GetBucketMetadataTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [GetBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html
// [DeleteBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataTableConfiguration.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [Integrating Amazon S3 Tables with Amazon Web Services analytics services]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-tables-integrating-aws.html
func s3_CreateBucketMetadataTableConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.CreateBucketMetadataTableConfigurationInput{
		// Bucket: *string, // Required
		// MetadataTableConfiguration: *types.MetadataTableConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3MetadataTableConfiguration) > 0 {
		if err := assignInputField(input, "MetadataTableConfiguration", _s3MetadataTableConfiguration); err != nil {
			log.Errorf("invalid --metadata-table-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.CreateBucketMetadataTableConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// This action initiates a multipart upload and returns an upload ID. This upload
// ID is used to associate all of the parts in the specific multipart upload. You
// specify this upload ID in each of your subsequent upload part requests (see [UploadPart]).
// You also include this upload ID in the final request to either complete or abort
// the multipart upload request. For more information about multipart uploads, see [Multipart Upload Overview]
// in the Amazon S3 User Guide.
//
// After you initiate a multipart upload and upload one or more parts, to stop
// being charged for storing the uploaded parts, you must either complete or abort
// the multipart upload. Amazon S3 frees up the space used to store the parts and
// stops charging you for storing them only after you either complete or abort a
// multipart upload.
//
// If you have configured a lifecycle rule to abort incomplete multipart uploads,
// the created multipart upload must be completed within the number of days
// specified in the bucket lifecycle configuration. Otherwise, the incomplete
// multipart upload becomes eligible for an abort action and Amazon S3 aborts the
// multipart upload. For more information, see [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration].
//
// - Directory buckets - S3 Lifecycle is not supported by directory buckets.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Request signing For request signing, multipart upload is just a series of
// regular requests. You initiate a multipart upload, send one or more requests to
// upload parts, and then complete the multipart upload process. You sign each
// request individually. There is nothing special about signing multipart upload
// requests. For more information about signing, see [Authenticating Requests (Amazon Web Services Signature Version 4)]in the Amazon S3 User Guide.
//
// # Permissions
//
// - General purpose bucket permissions - To perform a multipart upload with
// encryption using an Key Management Service (KMS) KMS key, the requester must
// have permission to the kms:Decrypt and kms:GenerateDataKey actions on the key.
// The requester must also have permissions for the kms:GenerateDataKey action
// for the CreateMultipartUpload API. Then, the requester needs permissions for
// the kms:Decrypt action on the UploadPart and UploadPartCopy APIs. These
// permissions are required because Amazon S3 must decrypt and read data from the
// encrypted file parts before it completes the multipart upload. For more
// information, see [Multipart upload API and permissions]and [Protecting data using server-side encryption with Amazon Web Services KMS]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # Encryption
//
// - General purpose buckets - Server-side encryption is for data encryption at
// rest. Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts it when you access it. Amazon S3 automatically encrypts all new
// objects that are uploaded to an S3 bucket. When doing a multipart upload, if you
// don't specify encryption information in your request, the encryption setting of
// the uploaded parts is set to the default encryption configuration of the
// destination bucket. By default, all buckets have a base level of encryption
// configuration that uses server-side encryption with Amazon S3 managed keys
// (SSE-S3). If the destination bucket has a default encryption configuration that
// uses server-side encryption with an Key Management Service (KMS) key (SSE-KMS),
// or a customer-provided encryption key (SSE-C), Amazon S3 uses the corresponding
// KMS key, or a customer-provided key to encrypt the uploaded parts. When you
// perform a CreateMultipartUpload operation, if you want to use a different type
// of encryption setting for the uploaded parts, you can request that Amazon S3
// encrypts the object with a different encryption key (such as an Amazon S3
// managed key, a KMS key, or a customer-provided key). When the encryption setting
// in your request is different from the default encryption configuration of the
// destination bucket, the encryption setting in your request takes precedence. If
// you choose to provide your own encryption key, the request headers you provide
// in [UploadPart]and [UploadPartCopy]requests must match the headers you used in the CreateMultipartUpload
// request.
//
// - Use KMS keys (SSE-KMS) that include the Amazon Web Services managed key (
// aws/s3 ) and KMS customer managed keys stored in Key Management Service (KMS)
// – If you want Amazon Web Services to manage the keys used to encrypt data,
// specify the following headers in the request.
//
// - x-amz-server-side-encryption
//
// - x-amz-server-side-encryption-aws-kms-key-id
//
// - x-amz-server-side-encryption-context
//
// - If you specify x-amz-server-side-encryption:aws:kms , but don't provide
// x-amz-server-side-encryption-aws-kms-key-id , Amazon S3 uses the Amazon Web
// Services managed key ( aws/s3 key) in KMS to protect the data.
//
// - To perform a multipart upload with encryption by using an Amazon Web
// Services KMS key, the requester must have permission to the kms:Decrypt and
// kms:GenerateDataKey* actions on the key. These permissions are required
// because Amazon S3 must decrypt and read data from the encrypted file parts
// before it completes the multipart upload. For more information, see [Multipart upload API and permissions]and [Protecting data using server-side encryption with Amazon Web Services KMS]in
// the Amazon S3 User Guide.
//
// - If your Identity and Access Management (IAM) user or role is in the same
// Amazon Web Services account as the KMS key, then you must have these permissions
// on the key policy. If your IAM user or role is in a different account from the
// key, then you must have the permissions on both the key policy and your IAM user
// or role.
//
// - All GET and PUT requests for an object protected by KMS fail if you don't
// make them by using Secure Sockets Layer (SSL), Transport Layer Security (TLS),
// or Signature Version 4. For information about configuring any of the officially
// supported Amazon Web Services SDKs and Amazon Web Services CLI, see [Specifying the Signature Version in Request Authentication]in the
// Amazon S3 User Guide.
//
// For more information about server-side encryption with KMS keys (SSE-KMS), see [Protecting Data Using Server-Side Encryption with KMS keys]
//
// in the Amazon S3 User Guide.
//
// - Use customer-provided encryption keys (SSE-C) – If you want to manage your
// own encryption keys, provide all the following headers in the request.
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// # For more information about server-side encryption with customer-provided
//
// encryption keys (SSE-C), see [Protecting data using server-side encryption with customer-provided encryption keys (SSE-C)]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: server-side encryption with Amazon S3
// managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ). We recommend that the bucket's default encryption uses
// the desired encryption configuration and you don't override the bucket default
// encryption in your CreateSession requests or PUT object requests. Then, new
// objects are automatically encrypted with the desired encryption settings. For
// more information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide. For more information about
// the encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads].
//
// In the Zonal endpoint API calls (except [CopyObject]and [UploadPartCopy]) using the REST API, the
//
// encryption request headers must match the encryption settings that are specified
// in the CreateSession request. You can't override the values of the encryption
// settings ( x-amz-server-side-encryption ,
// x-amz-server-side-encryption-aws-kms-key-id ,
// x-amz-server-side-encryption-context , and
// x-amz-server-side-encryption-bucket-key-enabled ) that are specified in the
// CreateSession request. You don't need to explicitly specify these encryption
// settings values in Zonal endpoint API calls, and Amazon S3 will use the
// encryption settings values from the CreateSession request to protect new
// objects in the directory bucket.
//
// # When you use the CLI or the Amazon Web Services SDKs, for CreateSession , the
//
// session token refreshes automatically to avoid service interruptions when a
// session expires. The CLI or the Amazon Web Services SDKs use the bucket's
// default encryption configuration for the CreateSession request. It's not
// supported to override the encryption settings values in the CreateSession
// request. So in the Zonal endpoint API calls (except [CopyObject]and [UploadPartCopy]), the encryption
// request headers must match the default encryption configuration of the directory
// bucket.
//
// # For directory buckets, when you perform a CreateMultipartUpload operation and an
//
// UploadPartCopy operation, the request headers you provide in the
// CreateMultipartUpload request must match the default encryption configuration
// of the destination bucket.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to CreateMultipartUpload :
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [AbortMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [Protecting Data Using Server-Side Encryption with KMS keys]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [Specifying the Signature Version in Request Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingAWSSDK.html#specify-signature-version
// [Aborting Incomplete Multipart Uploads Using a Bucket Lifecycle Configuration]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html#mpu-abort-incomplete-mpu-lifecycle-config
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart upload API and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpuAndPermissions
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [Multipart Upload Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html
// [Protecting data using server-side encryption with Amazon Web Services KMS]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [Protecting data using server-side encryption with customer-provided encryption keys (SSE-C)]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerSideEncryptionCustomerKeys.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
func s3_CreateMultipartUpload(cfg aws.Config, client *s3.Client) {
	input := &s3.CreateMultipartUploadInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ACL) > 0 {
		if err := assignInputField(input, "ACL", _s3ACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3BucketKeyEnabled) > 0 {
		if err := assignInputField(input, "BucketKeyEnabled", _s3BucketKeyEnabled); err != nil {
			log.Errorf("invalid --bucket-key-enabled: %s", err.Error())
			return
		}
	}
	if len(_s3CacheControl) > 0 {
		input.CacheControl = aws.String(_s3CacheControl)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumType) > 0 {
		if err := assignInputField(input, "ChecksumType", _s3ChecksumType); err != nil {
			log.Errorf("invalid --checksum-type: %s", err.Error())
			return
		}
	}
	if len(_s3ContentDisposition) > 0 {
		input.ContentDisposition = aws.String(_s3ContentDisposition)
	}
	if len(_s3ContentEncoding) > 0 {
		input.ContentEncoding = aws.String(_s3ContentEncoding)
	}
	if len(_s3ContentLanguage) > 0 {
		input.ContentLanguage = aws.String(_s3ContentLanguage)
	}
	if len(_s3ContentType) > 0 {
		input.ContentType = aws.String(_s3ContentType)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3Expires) > 0 {
		if err := assignInputField(input, "Expires", _s3Expires); err != nil {
			log.Errorf("invalid --expires: %s", err.Error())
			return
		}
	}
	if len(_s3GrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3GrantFullControl)
	}
	if len(_s3GrantRead) > 0 {
		input.GrantRead = aws.String(_s3GrantRead)
	}
	if len(_s3GrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3GrantReadACP)
	}
	if len(_s3GrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3GrantWriteACP)
	}
	if len(_s3Metadata) > 0 {
		if err := assignInputField(input, "Metadata", _s3Metadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockLegalHoldStatus) > 0 {
		if err := assignInputField(input, "ObjectLockLegalHoldStatus", _s3ObjectLockLegalHoldStatus); err != nil {
			log.Errorf("invalid --object-lock-legal-hold-status: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockMode) > 0 {
		if err := assignInputField(input, "ObjectLockMode", _s3ObjectLockMode); err != nil {
			log.Errorf("invalid --object-lock-mode: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockRetainUntilDate) > 0 {
		if err := assignInputField(input, "ObjectLockRetainUntilDate", _s3ObjectLockRetainUntilDate); err != nil {
			log.Errorf("invalid --object-lock-retain-until-date: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3SSEKMSEncryptionContext) > 0 {
		input.SSEKMSEncryptionContext = aws.String(_s3SSEKMSEncryptionContext)
	}
	if len(_s3SSEKMSKeyId) > 0 {
		input.SSEKMSKeyId = aws.String(_s3SSEKMSKeyId)
	}
	if len(_s3ServerSideEncryption) > 0 {
		if err := assignInputField(input, "ServerSideEncryption", _s3ServerSideEncryption); err != nil {
			log.Errorf("invalid --server-side-encryption: %s", err.Error())
			return
		}
	}
	if len(_s3StorageClass) > 0 {
		if err := assignInputField(input, "StorageClass", _s3StorageClass); err != nil {
			log.Errorf("invalid --storage-class: %s", err.Error())
			return
		}
	}
	if len(_s3Tagging) > 0 {
		input.Tagging = aws.String(_s3Tagging)
	}
	if len(_s3WebsiteRedirectLocation) > 0 {
		input.WebsiteRedirectLocation = aws.String(_s3WebsiteRedirectLocation)
	}

	if resp, err := client.CreateMultipartUpload(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a session that establishes temporary security credentials to support
// fast authentication and authorization for the Zonal endpoint API operations on
// directory buckets. For more information about Zonal endpoint API operations that
// include the Availability Zone in the request endpoint, see [S3 Express One Zone APIs]in the Amazon S3
// User Guide.
//
// To make Zonal endpoint API requests on a directory bucket, use the CreateSession
// API operation. Specifically, you grant s3express:CreateSession permission to a
// bucket in a bucket policy or an IAM identity-based policy. Then, you use IAM
// credentials to make the CreateSession API request on the bucket, which returns
// temporary security credentials that include the access key ID, secret access
// key, session token, and expiration. These credentials have associated
// permissions to access the Zonal endpoint API operations. After the session is
// created, you don’t need to use other policies to grant permissions to each Zonal
// endpoint API individually. Instead, in your Zonal endpoint API requests, you
// sign your requests by applying the temporary security credentials of the session
// to the request headers and following the SigV4 protocol for authentication. You
// also apply the session token to the x-amz-s3session-token request header for
// authorization. Temporary security credentials are scoped to the bucket and
// expire after 5 minutes. After the expiration time, any calls that you make with
// those credentials will fail. You must use IAM credentials again to make a
// CreateSession API request that generates a new set of temporary credentials for
// use. Temporary credentials cannot be extended or refreshed beyond the original
// specified interval.
//
// If you use Amazon Web Services SDKs, SDKs handle the session token refreshes
// automatically to avoid service interruptions when a session expires. We
// recommend that you use the Amazon Web Services SDKs to initiate and manage
// requests to the CreateSession API. For more information, see [Performance guidelines and design patterns]in the Amazon S3
// User Guide.
//
// - You must make requests for this API operation to the Zonal endpoint. These
// endpoints support virtual-hosted-style requests in the format
// https://bucket-name.s3express-zone-id.region-code.amazonaws.com . Path-style
// requests are not supported. For more information about endpoints in Availability
// Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information about endpoints
// in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// - CopyObject API operation - Unlike other Zonal endpoint API operations, the
// CopyObject API operation doesn't use the temporary security credentials
// returned from the CreateSession API operation for authentication and
// authorization. For information about authentication and authorization of the
// CopyObject API operation on directory buckets, see [CopyObject].
//
// - HeadBucket API operation - Unlike other Zonal endpoint API operations, the
// HeadBucket API operation doesn't use the temporary security credentials
// returned from the CreateSession API operation for authentication and
// authorization. For information about authentication and authorization of the
// HeadBucket API operation on directory buckets, see [HeadBucket].
//
// Permissions To obtain temporary security credentials, you must create a bucket
// policy or an IAM identity-based policy that grants s3express:CreateSession
// permission to the bucket. In a policy, you can have the s3express:SessionMode
// condition key to control who can create a ReadWrite or ReadOnly session. For
// more information about ReadWrite or ReadOnly sessions, see [x-amz-create-session-mode]
// x-amz-create-session-mode . For example policies, see [Example bucket policies for S3 Express One Zone] and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone] in the Amazon S3
// User Guide.
//
// To grant cross-account access to Zonal endpoint API operations, the bucket
// policy should also grant both accounts the s3express:CreateSession permission.
//
// If you want to encrypt objects with SSE-KMS, you must also have the
// kms:GenerateDataKey and the kms:Decrypt permissions in IAM identity-based
// policies and KMS key policies for the target KMS key.
//
// Encryption For directory buckets, there are only two supported options for
// server-side encryption: server-side encryption with Amazon S3 managed keys
// (SSE-S3) ( AES256 ) and server-side encryption with KMS keys (SSE-KMS) ( aws:kms
// ). We recommend that the bucket's default encryption uses the desired encryption
// configuration and you don't override the bucket default encryption in your
// CreateSession requests or PUT object requests. Then, new objects are
// automatically encrypted with the desired encryption settings. For more
// information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide. For more information about the
// encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads].
//
// For [Zonal endpoint (object-level) API operations] except [CopyObject] and [UploadPartCopy], you authenticate and authorize requests through [CreateSession] for low
// latency. To encrypt new objects in a directory bucket with SSE-KMS, you must
// specify SSE-KMS as the directory bucket's default encryption configuration with
// a KMS key (specifically, a [customer managed key]). Then, when a session is created for Zonal
// endpoint API operations, new objects are automatically encrypted and decrypted
// with SSE-KMS and S3 Bucket Keys during the session.
//
// Only 1 [customer managed key] is supported per directory bucket for the lifetime of the bucket. The [Amazon Web Services managed key] (
// aws/s3 ) isn't supported. After you specify SSE-KMS as your bucket's default
// encryption configuration with a customer managed key, you can't change the
// customer managed key for the bucket's SSE-KMS configuration.
//
// In the Zonal endpoint API calls (except [CopyObject] and [UploadPartCopy]) using the REST API, you can't
// override the values of the encryption settings ( x-amz-server-side-encryption ,
// x-amz-server-side-encryption-aws-kms-key-id ,
// x-amz-server-side-encryption-context , and
// x-amz-server-side-encryption-bucket-key-enabled ) from the CreateSession
// request. You don't need to explicitly specify these encryption settings values
// in Zonal endpoint API calls, and Amazon S3 will use the encryption settings
// values from the CreateSession request to protect new objects in the directory
// bucket.
//
// When you use the CLI or the Amazon Web Services SDKs, for CreateSession , the
// session token refreshes automatically to avoid service interruptions when a
// session expires. The CLI or the Amazon Web Services SDKs use the bucket's
// default encryption configuration for the CreateSession request. It's not
// supported to override the encryption settings values in the CreateSession
// request. Also, in the Zonal endpoint API calls (except [CopyObject]and [UploadPartCopy]), it's not
// supported to override the values of the encryption settings from the
// CreateSession request.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Performance guidelines and design patterns]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-optimizing-performance-guidelines-design-patterns.html#s3-express-optimizing-performance-session-authentication
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [S3 Express One Zone APIs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-APIs.html
// [HeadBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [x-amz-create-session-mode]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html#API_CreateSession_RequestParameters
// [Zonal endpoint (object-level) API operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-differences.html#s3-express-differences-api-operations
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_CreateSession(cfg aws.Config, client *s3.Client) {
	input := &s3.CreateSessionInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3BucketKeyEnabled) > 0 {
		if err := assignInputField(input, "BucketKeyEnabled", _s3BucketKeyEnabled); err != nil {
			log.Errorf("invalid --bucket-key-enabled: %s", err.Error())
			return
		}
	}
	if len(_s3SSEKMSEncryptionContext) > 0 {
		input.SSEKMSEncryptionContext = aws.String(_s3SSEKMSEncryptionContext)
	}
	if len(_s3SSEKMSKeyId) > 0 {
		input.SSEKMSKeyId = aws.String(_s3SSEKMSKeyId)
	}
	if len(_s3ServerSideEncryption) > 0 {
		if err := assignInputField(input, "ServerSideEncryption", _s3ServerSideEncryption); err != nil {
			log.Errorf("invalid --server-side-encryption: %s", err.Error())
			return
		}
	}
	if len(_s3SessionMode) > 0 {
		if err := assignInputField(input, "SessionMode", _s3SessionMode); err != nil {
			log.Errorf("invalid --session-mode: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateSession(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the S3 bucket. All objects (including all object versions and delete
// markers) in the bucket must be deleted before the bucket itself can be deleted.
//
// - Directory buckets - If multipart uploads in a directory bucket are in
// progress, you can't delete the bucket until all the in-progress multipart
// uploads are aborted or completed.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Regional endpoint. These endpoints support path-style
// requests in the format
// https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// # Permissions
//
// - General purpose bucket permissions - You must have the s3:DeleteBucket
// permission on the specified bucket in a policy.
//
// - Directory bucket permissions - You must have the s3express:DeleteBucket
// permission in an IAM identity-based policy instead of a bucket policy.
// Cross-account access to this API operation isn't supported. This operation can
// only be performed by the Amazon Web Services account that owns the resource. For
// more information about directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the
// Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to DeleteBucket :
//
// [CreateBucket]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_DeleteBucket(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes an analytics configuration for the bucket (specified by the analytics
// configuration ID).
//
// To use this operation, you must have permissions to perform the
// s3:PutAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about the Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class Analysis].
//
// The following operations are related to DeleteBucketAnalyticsConfiguration :
//
// [GetBucketAnalyticsConfiguration]
//
// [ListBucketAnalyticsConfigurations]
//
// [PutBucketAnalyticsConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Analytics – Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [GetBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAnalyticsConfiguration.html
// [ListBucketAnalyticsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketAnalyticsConfigurations.html
// [PutBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_DeleteBucketAnalyticsConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketAnalyticsConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketAnalyticsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes the cors configuration information set for the bucket.
//
// To use this operation, you must have permission to perform the s3:PutBucketCORS
// action. The bucket owner has this permission by default and can grant this
// permission to others.
//
// For information about cors , see [Enabling Cross-Origin Resource Sharing] in the Amazon S3 User Guide.
//
// # Related Resources
//
// [PutBucketCors]
//
// [RESTOPTIONSobject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [RESTOPTIONSobject]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html
func s3_DeleteBucketCors(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketCorsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketCors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This implementation of the DELETE action resets the default encryption for the
// bucket as server-side encryption with Amazon S3 managed keys (SSE-S3).
//
// - General purpose buckets - For information about the bucket default
// encryption feature, see [Amazon S3 Bucket Default Encryption]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: SSE-S3 and SSE-KMS. For information about
// the default encryption configuration in directory buckets, see [Setting default server-side encryption behavior for directory buckets].
//
// # Permissions
//
// - General purpose bucket permissions - The s3:PutEncryptionConfiguration
// permission is required in a policy. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:PutEncryptionConfiguration permission in an IAM
// identity-based policy instead of a bucket policy. Cross-account access to this
// API operation isn't supported. This operation can only be performed by the
// Amazon Web Services account that owns the resource. For more information about
// directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to DeleteBucketEncryption :
//
// [PutBucketEncryption]
//
// [GetBucketEncryption]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html
// [PutBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html
// [Setting default server-side encryption behavior for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-bucket-encryption.html
// [Amazon S3 Bucket Default Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Permissions Related to Bucket Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_DeleteBucketEncryption(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketEncryptionInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes the S3 Intelligent-Tiering configuration from the specified bucket.
//
// The S3 Intelligent-Tiering storage class is designed to optimize storage costs
// by automatically moving data to the most cost-effective storage access tier,
// without performance impact or operational overhead. S3 Intelligent-Tiering
// delivers automatic cost savings in three low latency and high throughput access
// tiers. To get the lowest storage cost on data that can be accessed in minutes to
// hours, you can choose to activate additional archiving capabilities.
//
// The S3 Intelligent-Tiering storage class is the ideal storage class for data
// with unknown, changing, or unpredictable access patterns, independent of object
// size or retention period. If the size of an object is less than 128 KB, it is
// not monitored and not eligible for auto-tiering. Smaller objects can be stored,
// but they are always charged at the Frequent Access tier rates in the S3
// Intelligent-Tiering storage class.
//
// For more information, see [Storage class for automatically optimizing frequently and infrequently accessed objects].
//
// Operations related to DeleteBucketIntelligentTieringConfiguration include:
//
// [GetBucketIntelligentTieringConfiguration]
//
// [PutBucketIntelligentTieringConfiguration]
//
// [ListBucketIntelligentTieringConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListBucketIntelligentTieringConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketIntelligentTieringConfigurations.html
// [GetBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html
// [PutBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html
// [Storage class for automatically optimizing frequently and infrequently accessed objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access
func s3_DeleteBucketIntelligentTieringConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketIntelligentTieringConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketIntelligentTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes an S3 Inventory configuration (identified by the inventory ID) from the
// bucket.
//
// To use this operation, you must have permissions to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory].
//
// Operations related to DeleteBucketInventoryConfiguration include:
//
// [GetBucketInventoryConfiguration]
//
// [PutBucketInventoryConfiguration]
//
// [ListBucketInventoryConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [ListBucketInventoryConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [PutBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
func s3_DeleteBucketInventoryConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketInventoryConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketInventoryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the lifecycle configuration from the specified bucket. Amazon S3
// removes all the lifecycle configuration rules in the lifecycle subresource
// associated with the bucket. Your objects never expire, and Amazon S3 no longer
// automatically deletes any objects on the basis of rules contained in the deleted
// lifecycle configuration.
//
// Permissions
// - General purpose bucket permissions - By default, all Amazon S3 resources
// are private, including buckets, objects, and related subresources (for example,
// lifecycle configuration and website configuration). Only the resource owner
// (that is, the Amazon Web Services account that created it) can access the
// resource. The resource owner can optionally grant access permissions to others
// by writing an access policy. For this operation, a user must have the
// s3:PutLifecycleConfiguration permission.
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - You must have the
// s3express:PutLifecycleConfiguration permission in an IAM identity-based policy
// to use this operation. Cross-account access to this API operation isn't
// supported. The resource owner can optionally grant access permissions to others
// by creating a role or user for them as long as they are within the same account
// as the owner and resource.
//
// For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM]in
//
// the Amazon S3 User Guide.
//
// # Directory buckets - For directory buckets, you must make requests for this API
//
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name
// . Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// For more information about the object expiration, see [Elements to Describe Lifecycle Actions].
//
// Related actions include:
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketLifecycleConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [Elements to Describe Lifecycle Actions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html#intro-lifecycle-rules-actions
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_DeleteBucketLifecycle(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketLifecycleInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketLifecycle(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes an S3 Metadata configuration from a general purpose bucket. For more
// information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// You can use the V2 DeleteBucketMetadataConfiguration API operation with V1 or
// V2 metadata configurations. However, if you try to use the V1
// DeleteBucketMetadataTableConfiguration API operation with V2 configurations, you
// will receive an HTTP 405 Method Not Allowed error.
//
// Permissions To use this operation, you must have the
// s3:DeleteBucketMetadataTableConfiguration permission. For more information, see [Setting up permissions for configuring metadata tables]
// in the Amazon S3 User Guide.
//
// The IAM policy action name is the same for the V1 and V2 API operations.
//
// The following operations are related to DeleteBucketMetadataConfiguration :
//
// [CreateBucketMetadataConfiguration]
//
// [GetBucketMetadataConfiguration]
//
// [UpdateBucketMetadataInventoryTableConfiguration]
//
// [UpdateBucketMetadataJournalTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [UpdateBucketMetadataJournalTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [UpdateBucketMetadataInventoryTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html
func s3_DeleteBucketMetadataConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketMetadataConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketMetadataConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend that you delete your S3 Metadata configurations by using the V2 [DeleteBucketMetadataTableConfiguration]
// API operation. We no longer recommend using the V1
// DeleteBucketMetadataTableConfiguration API operation.
//
// If you created your S3 Metadata configuration before July 15, 2025, we
// recommend that you delete and re-create your configuration by using [CreateBucketMetadataConfiguration]so that you
// can expire journal table records and create a live inventory table.
//
// Deletes a V1 S3 Metadata configuration from a general purpose bucket. For more
// information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// You can use the V2 DeleteBucketMetadataConfiguration API operation with V1 or
// V2 metadata table configurations. However, if you try to use the V1
// DeleteBucketMetadataTableConfiguration API operation with V2 configurations, you
// will receive an HTTP 405 Method Not Allowed error.
//
// Make sure that you update your processes to use the new V2 API operations (
// CreateBucketMetadataConfiguration , GetBucketMetadataConfiguration , and
// DeleteBucketMetadataConfiguration ) instead of the V1 API operations.
//
// Permissions To use this operation, you must have the
// s3:DeleteBucketMetadataTableConfiguration permission. For more information, see [Setting up permissions for configuring metadata tables]
// in the Amazon S3 User Guide.
//
// The following operations are related to DeleteBucketMetadataTableConfiguration :
//
// [CreateBucketMetadataTableConfiguration]
//
// [GetBucketMetadataTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [GetBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html
// [CreateBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataTableConfiguration.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [DeleteBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataTableConfiguration.html
func s3_DeleteBucketMetadataTableConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketMetadataTableConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketMetadataTableConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes a metrics configuration for the Amazon CloudWatch request metrics
// (specified by the metrics configuration ID) from the bucket. Note that this
// doesn't include the daily storage metrics.
//
// To use this operation, you must have permissions to perform the
// s3:PutMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about CloudWatch request metrics for Amazon S3, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to DeleteBucketMetricsConfiguration :
//
// [GetBucketMetricsConfiguration]
//
// [PutBucketMetricsConfiguration]
//
// [ListBucketMetricsConfigurations]
//
// [Monitoring Metrics with Amazon CloudWatch]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [GetBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html
// [ListBucketMetricsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html
// [PutBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_DeleteBucketMetricsConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketMetricsConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketMetricsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Removes OwnershipControls for an Amazon S3 bucket. To use this operation, you
// must have the s3:PutBucketOwnershipControls permission. For more information
// about Amazon S3 permissions, see [Specifying Permissions in a Policy].
//
// For information about Amazon S3 Object Ownership, see [Using Object Ownership].
//
// The following operations are related to DeleteBucketOwnershipControls :
//
// # GetBucketOwnershipControls
//
// # PutBucketOwnershipControls
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Using Object Ownership]: https://docs.aws.amazon.com/AmazonS3/latest/dev/about-object-ownership.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
func s3_DeleteBucketOwnershipControls(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketOwnershipControlsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketOwnershipControls(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes the policy of a specified bucket.
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions If you are using an identity other than the root user of the Amazon
// Web Services account that owns the bucket, the calling identity must both have
// the DeleteBucketPolicy permissions on the specified bucket and belong to the
// bucket owner's account in order to use this operation.
//
// If you don't have DeleteBucketPolicy permissions, Amazon S3 returns a 403
// Access Denied error. If you have the correct permissions, but you're not using
// an identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
//
// To ensure that bucket owners don't inadvertently lock themselves out of their
// own buckets, the root principal in a bucket owner's Amazon Web Services account
// can perform the GetBucketPolicy , PutBucketPolicy , and DeleteBucketPolicy API
// actions, even if their bucket policy explicitly denies the root principal's
// access. Bucket owner root principals can only be blocked from performing these
// API actions by VPC endpoint policies and Amazon Web Services Organizations
// policies.
//
// - General purpose bucket permissions - The s3:DeleteBucketPolicy permission is
// required in a policy. For more information about general purpose buckets bucket
// policies, see [Using Bucket Policies and User Policies]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:DeleteBucketPolicy permission in an IAM identity-based
// policy instead of a bucket policy. Cross-account access to this API operation
// isn't supported. This operation can only be performed by the Amazon Web Services
// account that owns the resource. For more information about directory bucket
// policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// # The following operations are related to DeleteBucketPolicy
//
// [CreateBucket]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_DeleteBucketPolicy(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketPolicyInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes the replication configuration from the bucket.
//
// To use this operation, you must have permissions to perform the
// s3:PutReplicationConfiguration action. The bucket owner has these permissions by
// default and can grant it to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations]
// and [Managing Access Permissions to Your Amazon S3 Resources].
//
// It can take a while for the deletion of a replication configuration to fully
// propagate.
//
// For information about replication configuration, see [Replication] in the Amazon S3 User
// Guide.
//
// The following operations are related to DeleteBucketReplication :
//
// [PutBucketReplication]
//
// [GetBucketReplication]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html
// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_DeleteBucketReplication(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketReplicationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Deletes tags from the general purpose bucket if attribute based access control
// (ABAC) is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use [UntagResource]instead.
//
// if ABAC is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use [UntagResource]instead.
//
// To use this operation, you must have permission to perform the
// s3:PutBucketTagging action. By default, the bucket owner has this permission and
// can grant this permission to others.
//
// The following operations are related to DeleteBucketTagging :
//
// [GetBucketTagging]
//
// [PutBucketTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [enable ABAC for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UntagResource.html
func s3_DeleteBucketTagging(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketTaggingInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// This action removes the website configuration for a bucket. Amazon S3 returns a
// 200 OK response upon successfully deleting a website configuration on the
// specified bucket. You will get a 200 OK response if the website configuration
// you are trying to delete does not exist on the bucket. Amazon S3 returns a 404
// response if the bucket specified in the request does not exist.
//
// This DELETE action requires the S3:DeleteBucketWebsite permission. By default,
// only the bucket owner can delete the website configuration attached to a bucket.
// However, bucket owners can grant other users permission to delete the website
// configuration by writing a bucket policy granting them the
// S3:DeleteBucketWebsite permission.
//
// For more information about hosting websites, see [Hosting Websites on Amazon S3].
//
// The following operations are related to DeleteBucketWebsite :
//
// [GetBucketWebsite]
//
// [PutBucketWebsite]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketWebsite.html
// [PutBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html
// [Hosting Websites on Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
func s3_DeleteBucketWebsite(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteBucketWebsiteInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeleteBucketWebsite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes an object from a bucket. The behavior depends on the bucket's
// versioning state:
//
// - If bucket versioning is not enabled, the operation permanently deletes the
// object.
//
// - If bucket versioning is enabled, the operation inserts a delete marker,
// which becomes the current version of the object. To permanently delete an object
// in a versioned bucket, you must include the object’s versionId in the request.
// For more information about versioning-enabled buckets, see [Deleting object versions from a versioning-enabled bucket].
//
// - If bucket versioning is suspended, the operation removes the object that
// has a null versionId , if there is one, and inserts a delete marker that
// becomes the current version of the object. If there isn't an object with a null
// versionId , and all versions of the object have a versionId , Amazon S3 does
// not remove the object and only inserts a delete marker. To permanently delete an
// object that has a versionId , you must include the object’s versionId in the
// request. For more information about versioning-suspended buckets, see [Deleting objects from versioning-suspended buckets].
//
// - Directory buckets - S3 Versioning isn't enabled and supported for directory
// buckets. For this API operation, only the null value of the version ID is
// supported by directory buckets. You can only specify null to the versionId
// query parameter in the request.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// To remove a specific version, you must use the versionId query parameter. Using
// this query parameter permanently deletes the version. If the object deleted is a
// delete marker, Amazon S3 sets the response header x-amz-delete-marker to true.
//
// If the object you want to delete is in a bucket where the bucket versioning
// configuration is MFA Delete enabled, you must include the x-amz-mfa request
// header in the DELETE versionId request. Requests that include x-amz-mfa must
// use HTTPS. For more information about MFA Delete, see [Using MFA Delete]in the Amazon S3 User
// Guide. To see sample requests that use versioning, see [Sample Request].
//
// Directory buckets - MFA delete is not supported by directory buckets.
//
// You can delete objects by explicitly calling DELETE Object or calling ([PutBucketLifecycle] ) to
// enable Amazon S3 to remove them for you. If you want to block users or accounts
// from removing or deleting objects from your bucket, you must deny them the
// s3:DeleteObject , s3:DeleteObjectVersion , and s3:PutLifeCycleConfiguration
// actions.
//
// Directory buckets - S3 Lifecycle is not supported by directory buckets.
//
// # Permissions
//
// - General purpose bucket permissions - The following permissions are required
// in your policies when your DeleteObjects request includes specific headers.
//
// - s3:DeleteObject - To delete an object from a bucket, you must always have
// the s3:DeleteObject permission.
//
// - s3:DeleteObjectVersion - To delete a specific version of an object from a
// versioning-enabled bucket, you must have the s3:DeleteObjectVersion permission.
//
// If the s3:DeleteObject or s3:DeleteObjectVersion permissions are explicitly
//
// denied in your bucket policy, attempts to delete any unversioned objects result
// in a 403 Access Denied error.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following action is related to DeleteObject :
//
// [PutObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// The If-Match header is supported for both general purpose and directory
// buckets. IfMatchLastModifiedTime and IfMatchSize is only supported for
// directory buckets.
//
// [Sample Request]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectDELETE.html#ExampleVersionObjectDelete
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Deleting objects from versioning-suspended buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/DeletingObjectsfromVersioningSuspendedBuckets.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Deleting object versions from a versioning-enabled bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/DeletingObjectVersions.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Using MFA Delete]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingMFADelete.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_DeleteObject(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteObjectInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3BypassGovernanceRetention) > 0 {
		if err := assignInputField(input, "BypassGovernanceRetention", _s3BypassGovernanceRetention); err != nil {
			log.Errorf("invalid --bypass-governance-retention: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3IfMatch) > 0 {
		input.IfMatch = aws.String(_s3IfMatch)
	}
	if len(_s3IfMatchLastModifiedTime) > 0 {
		if err := assignInputField(input, "IfMatchLastModifiedTime", _s3IfMatchLastModifiedTime); err != nil {
			log.Errorf("invalid --if-match-last-modified-time: %s", err.Error())
			return
		}
	}
	if len(_s3IfMatchSize) > 0 {
		if err := assignInputField(input, "IfMatchSize", _s3IfMatchSize); err != nil {
			log.Errorf("invalid --if-match-size: %s", err.Error())
			return
		}
	}
	if len(_s3MFA) > 0 {
		input.MFA = aws.String(_s3MFA)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.DeleteObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Removes the entire tag set from the specified object. For more information
// about managing object tags, see [Object Tagging].
//
// To use this operation, you must have permission to perform the
// s3:DeleteObjectTagging action.
//
// To delete tags of a specific object version, add the versionId query parameter
// in the request. You will need permission for the s3:DeleteObjectVersionTagging
// action.
//
// The following operations are related to DeleteObjectTagging :
//
// [PutObjectTagging]
//
// [GetObjectTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html
// [Object Tagging]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html
// [GetObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html
func s3_DeleteObjectTagging(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteObjectTaggingInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.DeleteObjectTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation enables you to delete multiple objects from a bucket using a
// single HTTP request. If you know the object keys that you want to delete, then
// this operation provides a suitable alternative to sending individual delete
// requests, reducing per-request overhead.
//
// The request can contain a list of up to 1,000 keys that you want to delete. In
// the XML, you provide the object key names, and optionally, version IDs if you
// want to delete a specific version of the object from a versioning-enabled
// bucket. For each key, Amazon S3 performs a delete operation and returns the
// result of that delete, success or failure, in the response. If the object
// specified in the request isn't found, Amazon S3 confirms the deletion by
// returning the result as deleted.
//
// - Directory buckets - S3 Versioning isn't enabled and supported for directory
// buckets.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// The operation supports two modes for the response: verbose and quiet. By
// default, the operation uses verbose mode in which the response includes the
// result of deletion of each key in your request. In quiet mode the response
// includes only keys where the delete operation encountered an error. For a
// successful deletion in a quiet mode, the operation does not return any
// information about the delete in the response body.
//
// When performing this action on an MFA Delete enabled bucket, that attempts to
// delete any versioned objects, you must include an MFA token. If you do not
// provide one, the entire request will fail, even if there are non-versioned
// objects you are trying to delete. If you provide an invalid token, whether there
// are versioned keys in the request or not, the entire Multi-Object Delete request
// will fail. For information about MFA Delete, see [MFA Delete]in the Amazon S3 User Guide.
//
// Directory buckets - MFA delete is not supported by directory buckets.
//
// # Permissions
//
// - General purpose bucket permissions - The following permissions are required
// in your policies when your DeleteObjects request includes specific headers.
//
// - s3:DeleteObject - To delete an object from a bucket, you must always specify
// the s3:DeleteObject permission.
//
// - s3:DeleteObjectVersion - To delete a specific version of an object from a
// versioning-enabled bucket, you must specify the s3:DeleteObjectVersion
// permission.
//
// If the s3:DeleteObject or s3:DeleteObjectVersion permissions are explicitly
//
// denied in your bucket policy, attempts to delete any unversioned objects result
// in a 403 Access Denied error.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # Content-MD5 request header
//
// - General purpose bucket - The Content-MD5 request header is required for all
// Multi-Object Delete requests. Amazon S3 uses the header value to ensure that
// your request body has not been altered in transit.
//
// - Directory bucket - The Content-MD5 request header or a additional checksum
// request header (including x-amz-checksum-crc32 , x-amz-checksum-crc32c ,
// x-amz-checksum-sha1 , or x-amz-checksum-sha256 ) is required for all
// Multi-Object Delete requests.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to DeleteObjects :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [ListParts]
//
// [AbortMultipartUpload]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [MFA Delete]: https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html#MultiFactorAuthenticationDelete
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_DeleteObjects(cfg aws.Config, client *s3.Client) {
	input := &s3.DeleteObjectsInput{
		// Bucket: *string, // Required
		// Delete: *types.Delete, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Delete) > 0 {
		if err := assignInputField(input, "Delete", _s3Delete); err != nil {
			log.Errorf("invalid --delete: %s", err.Error())
			return
		}
	}
	if len(_s3BypassGovernanceRetention) > 0 {
		if err := assignInputField(input, "BypassGovernanceRetention", _s3BypassGovernanceRetention); err != nil {
			log.Errorf("invalid --bypass-governance-retention: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3MFA) > 0 {
		input.MFA = aws.String(_s3MFA)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}

	if resp, err := client.DeleteObjects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Removes the PublicAccessBlock configuration for an Amazon S3 bucket. This
// operation removes the bucket-level configuration only. The effective public
// access behavior will still be governed by account-level settings (which may
// inherit from organization-level policies). To use this operation, you must have
// the s3:PutBucketPublicAccessBlock permission. For more information about
// permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// The following operations are related to DeletePublicAccessBlock :
//
// [Using Amazon S3 Block Public Access]
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [GetBucketPolicyStatus]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [GetBucketPolicyStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html
func s3_DeletePublicAccessBlock(cfg aws.Config, client *s3.Client) {
	input := &s3.DeletePublicAccessBlockInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.DeletePublicAccessBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the attribute-based access control (ABAC) property of the general
// purpose bucket. If ABAC is enabled on your bucket, you can use tags on the
// bucket for access control. For more information, see [Enabling ABAC in general purpose buckets].
//
// [Enabling ABAC in general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
func s3_GetBucketAbac(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketAbacInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketAbac(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// This implementation of the GET action uses the accelerate subresource to return
// the Transfer Acceleration state of a bucket, which is either Enabled or
// Suspended . Amazon S3 Transfer Acceleration is a bucket-level feature that
// enables you to perform faster data transfers to and from Amazon S3.
//
// To use this operation, you must have permission to perform the
// s3:GetAccelerateConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to your Amazon S3 Resources] in the Amazon S3 User Guide.
//
// You set the Transfer Acceleration state of an existing bucket to Enabled or
// Suspended by using the [PutBucketAccelerateConfiguration] operation.
//
// A GET accelerate request does not return a state value for a bucket that has no
// transfer acceleration state. A bucket has no Transfer Acceleration state if a
// state has never been set on the bucket.
//
// For more information about transfer acceleration, see [Transfer Acceleration] in the Amazon S3 User
// Guide.
//
// The following operations are related to GetBucketAccelerateConfiguration :
//
// [PutBucketAccelerateConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketAccelerateConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAccelerateConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Managing Access Permissions to your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Transfer Acceleration]: https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html
func s3_GetBucketAccelerateConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketAccelerateConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetBucketAccelerateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// This implementation of the GET action uses the acl subresource to return the
// access control list (ACL) of a bucket. To use GET to return the ACL of the
// bucket, you must have the READ_ACP access to the bucket. If READ_ACP permission
// is granted to the anonymous user, you can return the ACL of the bucket without
// using an authorization header.
//
// When you use this API operation with an access point, provide the alias of the
// access point in place of the bucket name.
//
// When you use this API operation with an Object Lambda access point, provide the
// alias of the Object Lambda access point in place of the bucket name. If the
// Object Lambda access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see [List of Error Codes].
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// requests to read ACLs are still supported and return the
// bucket-owner-full-control ACL with the owner being the account that created the
// bucket. For more information, see [Controlling object ownership and disabling ACLs]in the Amazon S3 User Guide.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// The following operations are related to GetBucketAcl :
//
// [ListObjects]
//
// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [Controlling object ownership and disabling ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
func s3_GetBucketAcl(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketAclInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketAcl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// This implementation of the GET action returns an analytics configuration
// (identified by the analytics configuration ID) from the bucket.
//
// To use this operation, you must have permissions to perform the
// s3:GetAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources] in the Amazon S3 User Guide.
//
// For information about Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class Analysis] in the Amazon S3 User
// Guide.
//
// The following operations are related to GetBucketAnalyticsConfiguration :
//
// [DeleteBucketAnalyticsConfiguration]
//
// [ListBucketAnalyticsConfigurations]
//
// [PutBucketAnalyticsConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Analytics – Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html
// [DeleteBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [ListBucketAnalyticsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketAnalyticsConfigurations.html
// [PutBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_GetBucketAnalyticsConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketAnalyticsConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketAnalyticsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the Cross-Origin Resource Sharing (CORS) configuration information set
// for the bucket.
//
// To use this operation, you must have permission to perform the s3:GetBucketCORS
// action. By default, the bucket owner has this permission and can grant it to
// others.
//
// When you use this API operation with an access point, provide the alias of the
// access point in place of the bucket name.
//
// When you use this API operation with an Object Lambda access point, provide the
// alias of the Object Lambda access point in place of the bucket name. If the
// Object Lambda access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see [List of Error Codes].
//
// For more information about CORS, see [Enabling Cross-Origin Resource Sharing].
//
// The following operations are related to GetBucketCors :
//
// [PutBucketCors]
//
// [DeleteBucketCors]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [DeleteBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html
func s3_GetBucketCors(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketCorsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketCors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the default encryption configuration for an Amazon S3 bucket. By
// default, all buckets have a default encryption configuration that uses
// server-side encryption with Amazon S3 managed keys (SSE-S3). This operation also
// returns the [BucketKeyEnabled]and [BlockedEncryptionTypes] statuses.
//
// - General purpose buckets - For information about the bucket default
// encryption feature, see [Amazon S3 Bucket Default Encryption]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: SSE-S3 and SSE-KMS. For information about
// the default encryption configuration in directory buckets, see [Setting default server-side encryption behavior for directory buckets].
//
// # Permissions
//
// - General purpose bucket permissions - The s3:GetEncryptionConfiguration
// permission is required in a policy. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:GetEncryptionConfiguration permission in an IAM
// identity-based policy instead of a bucket policy. Cross-account access to this
// API operation isn't supported. This operation can only be performed by the
// Amazon Web Services account that owns the resource. For more information about
// directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to GetBucketEncryption :
//
// [PutBucketEncryption]
//
// [DeleteBucketEncryption]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [BucketKeyEnabled]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html#AmazonS3-Type-ServerSideEncryptionRule-BucketKeyEnabled
// [BlockedEncryptionTypes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ServerSideEncryptionRule.html#AmazonS3-Type-ServerSideEncryptionRule-BlockedEncryptionTypes
// [DeleteBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html
// [PutBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketEncryption.html
// [Setting default server-side encryption behavior for directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-bucket-encryption.html
// [Amazon S3 Bucket Default Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucket-encryption.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Permissions Related to Bucket Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_GetBucketEncryption(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketEncryptionInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Gets the S3 Intelligent-Tiering configuration from the specified bucket.
//
// The S3 Intelligent-Tiering storage class is designed to optimize storage costs
// by automatically moving data to the most cost-effective storage access tier,
// without performance impact or operational overhead. S3 Intelligent-Tiering
// delivers automatic cost savings in three low latency and high throughput access
// tiers. To get the lowest storage cost on data that can be accessed in minutes to
// hours, you can choose to activate additional archiving capabilities.
//
// The S3 Intelligent-Tiering storage class is the ideal storage class for data
// with unknown, changing, or unpredictable access patterns, independent of object
// size or retention period. If the size of an object is less than 128 KB, it is
// not monitored and not eligible for auto-tiering. Smaller objects can be stored,
// but they are always charged at the Frequent Access tier rates in the S3
// Intelligent-Tiering storage class.
//
// For more information, see [Storage class for automatically optimizing frequently and infrequently accessed objects].
//
// Operations related to GetBucketIntelligentTieringConfiguration include:
//
// [DeleteBucketIntelligentTieringConfiguration]
//
// [PutBucketIntelligentTieringConfiguration]
//
// [ListBucketIntelligentTieringConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListBucketIntelligentTieringConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketIntelligentTieringConfigurations.html
// [PutBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html
// [Storage class for automatically optimizing frequently and infrequently accessed objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access
// [DeleteBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html
func s3_GetBucketIntelligentTieringConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketIntelligentTieringConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketIntelligentTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns an S3 Inventory configuration (identified by the inventory
// configuration ID) from the bucket.
//
// To use this operation, you must have permissions to perform the
// s3:GetInventoryConfiguration action. The bucket owner has this permission by
// default and can grant this permission to others. For more information about
// permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory].
//
// The following operations are related to GetBucketInventoryConfiguration :
//
// [DeleteBucketInventoryConfiguration]
//
// [ListBucketInventoryConfigurations]
//
// [PutBucketInventoryConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [ListBucketInventoryConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [PutBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html
func s3_GetBucketInventoryConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketInventoryConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketInventoryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the lifecycle configuration information set on the bucket. For
// information about lifecycle configuration, see [Object Lifecycle Management].
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, object size, or any
// combination of these. Accordingly, this section describes the latest API, which
// is compatible with the new functionality. The previous version of the API
// supported filtering based only on an object key name prefix, which is supported
// for general purpose buckets for backward compatibility. For the related API
// description, see [GetBucketLifecycle].
//
// Lifecyle configurations for directory buckets only support expiring objects and
// cancelling multipart uploads. Expiring of versioned objects, transitions and tag
// filters are not supported.
//
// Permissions
// - General purpose bucket permissions - By default, all Amazon S3 resources
// are private, including buckets, objects, and related subresources (for example,
// lifecycle configuration and website configuration). Only the resource owner
// (that is, the Amazon Web Services account that created it) can access the
// resource. The resource owner can optionally grant access permissions to others
// by writing an access policy. For this operation, a user must have the
// s3:GetLifecycleConfiguration permission.
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - You must have the
// s3express:GetLifecycleConfiguration permission in an IAM identity-based policy
// to use this operation. Cross-account access to this API operation isn't
// supported. The resource owner can optionally grant access permissions to others
// by creating a role or user for them as long as they are within the same account
// as the owner and resource.
//
// For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM]in
//
// the Amazon S3 User Guide.
//
// # Directory buckets - For directory buckets, you must make requests for this API
//
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name
// . Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
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
// The following operations are related to GetBucketLifecycleConfiguration :
//
// [GetBucketLifecycle]
//
// [PutBucketLifecycle]
//
// [DeleteBucketLifecycle]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycle.html
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_GetBucketLifecycleConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketLifecycleConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketLifecycleConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Using the GetBucketLocation operation is no longer a best practice. To return
// the Region that a bucket resides in, we recommend that you use the [HeadBucket]operation
// instead. For backward compatibility, Amazon S3 continues to support the
// GetBucketLocation operation.
//
// Returns the Region the bucket resides in. You set the bucket's Region using the
// LocationConstraint request parameter in a CreateBucket request. For more
// information, see [CreateBucket].
//
// In a bucket's home Region, calls to the GetBucketLocation operation are
// governed by the bucket's policy. In other Regions, the bucket policy doesn't
// apply, which means that cross-account access won't be authorized. However, calls
// to the HeadBucket operation always return the bucket’s location through an HTTP
// response header, whether access to the bucket is authorized or not. Therefore,
// we recommend using the HeadBucket operation for bucket Region discovery and to
// avoid using the GetBucketLocation operation.
//
// When you use this API operation with an access point, provide the alias of the
// access point in place of the bucket name.
//
// When you use this API operation with an Object Lambda access point, provide the
// alias of the Object Lambda access point in place of the bucket name. If the
// Object Lambda access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see [List of Error Codes].
//
// This operation is not supported for directory buckets.
//
// The following operations are related to GetBucketLocation :
//
// [GetObject]
//
// [CreateBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [HeadBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadBucket.html
func s3_GetBucketLocation(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketLocationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketLocation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the logging status of a bucket and the permissions users have to view
// and modify that status.
//
// The following operations are related to GetBucketLogging :
//
// [CreateBucket]
//
// [PutBucketLogging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketLogging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLogging.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
func s3_GetBucketLogging(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketLoggingInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves the S3 Metadata configuration for a general purpose bucket. For more
// information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// You can use the V2 GetBucketMetadataConfiguration API operation with V1 or V2
// metadata configurations. However, if you try to use the V1
// GetBucketMetadataTableConfiguration API operation with V2 configurations, you
// will receive an HTTP 405 Method Not Allowed error.
//
// Permissions To use this operation, you must have the
// s3:GetBucketMetadataTableConfiguration permission. For more information, see [Setting up permissions for configuring metadata tables]
// in the Amazon S3 User Guide.
//
// The IAM policy action name is the same for the V1 and V2 API operations.
//
// The following operations are related to GetBucketMetadataConfiguration :
//
// [CreateBucketMetadataConfiguration]
//
// [DeleteBucketMetadataConfiguration]
//
// [UpdateBucketMetadataInventoryTableConfiguration]
//
// [UpdateBucketMetadataJournalTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [UpdateBucketMetadataJournalTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [UpdateBucketMetadataInventoryTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html
// [DeleteBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html
func s3_GetBucketMetadataConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketMetadataConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketMetadataConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// We recommend that you retrieve your S3 Metadata configurations by using the V2 [GetBucketMetadataTableConfiguration]
// API operation. We no longer recommend using the V1
// GetBucketMetadataTableConfiguration API operation.
//
// If you created your S3 Metadata configuration before July 15, 2025, we
// recommend that you delete and re-create your configuration by using [CreateBucketMetadataConfiguration]so that you
// can expire journal table records and create a live inventory table.
//
// Retrieves the V1 S3 Metadata configuration for a general purpose bucket. For
// more information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// You can use the V2 GetBucketMetadataConfiguration API operation with V1 or V2
// metadata table configurations. However, if you try to use the V1
// GetBucketMetadataTableConfiguration API operation with V2 configurations, you
// will receive an HTTP 405 Method Not Allowed error.
//
// Make sure that you update your processes to use the new V2 API operations (
// CreateBucketMetadataConfiguration , GetBucketMetadataConfiguration , and
// DeleteBucketMetadataConfiguration ) instead of the V1 API operations.
//
// Permissions To use this operation, you must have the
// s3:GetBucketMetadataTableConfiguration permission. For more information, see [Setting up permissions for configuring metadata tables]
// in the Amazon S3 User Guide.
//
// The following operations are related to GetBucketMetadataTableConfiguration :
//
// [CreateBucketMetadataTableConfiguration]
//
// [DeleteBucketMetadataTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [CreateBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataTableConfiguration.html
// [DeleteBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataTableConfiguration.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [GetBucketMetadataTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataTableConfiguration.html
func s3_GetBucketMetadataTableConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketMetadataTableConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketMetadataTableConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Gets a metrics configuration (specified by the metrics configuration ID) from
// the bucket. Note that this doesn't include the daily storage metrics.
//
// To use this operation, you must have permissions to perform the
// s3:GetMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about CloudWatch request metrics for Amazon S3, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to GetBucketMetricsConfiguration :
//
// [PutBucketMetricsConfiguration]
//
// [DeleteBucketMetricsConfiguration]
//
// [ListBucketMetricsConfigurations]
//
// [Monitoring Metrics with Amazon CloudWatch]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [ListBucketMetricsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html
// [PutBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html
// [DeleteBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_GetBucketMetricsConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketMetricsConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketMetricsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the notification configuration of a bucket.
//
// If notifications are not enabled on the bucket, the action returns an empty
// NotificationConfiguration element.
//
// By default, you must be the bucket owner to read the notification configuration
// of a bucket. However, the bucket owner can use a bucket policy to grant
// permission to other users to read this configuration with the
// s3:GetBucketNotification permission.
//
// When you use this API operation with an access point, provide the alias of the
// access point in place of the bucket name.
//
// When you use this API operation with an Object Lambda access point, provide the
// alias of the Object Lambda access point in place of the bucket name. If the
// Object Lambda access point alias in a request is not valid, the error code
// InvalidAccessPointAliasError is returned. For more information about
// InvalidAccessPointAliasError , see [List of Error Codes].
//
// For more information about setting and reading the notification configuration
// on a bucket, see [Setting Up Notification of Bucket Events]. For more information about bucket policies, see [Using Bucket Policies].
//
// The following action is related to GetBucketNotification :
//
// [PutBucketNotification]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Using Bucket Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [Setting Up Notification of Bucket Events]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
// [List of Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ErrorCodeList
// [PutBucketNotification]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketNotification.html
func s3_GetBucketNotificationConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketNotificationConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Retrieves OwnershipControls for an Amazon S3 bucket. To use this operation, you
// must have the s3:GetBucketOwnershipControls permission. For more information
// about Amazon S3 permissions, see [Specifying permissions in a policy].
//
// A bucket doesn't have OwnershipControls settings in the following cases:
//
// - The bucket was created before the BucketOwnerEnforced ownership setting was
// introduced and you've never explicitly applied this value
//
// - You've manually deleted the bucket ownership control value using the
// DeleteBucketOwnershipControls API operation.
//
// By default, Amazon S3 sets OwnershipControls for all newly created buckets.
//
// For information about Amazon S3 Object Ownership, see [Using Object Ownership].
//
// The following operations are related to GetBucketOwnershipControls :
//
// # PutBucketOwnershipControls
//
// # DeleteBucketOwnershipControls
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Using Object Ownership]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
// [Specifying permissions in a policy]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html
func s3_GetBucketOwnershipControls(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketOwnershipControlsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketOwnershipControls(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns the policy of a specified bucket.
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions If you are using an identity other than the root user of the Amazon
// Web Services account that owns the bucket, the calling identity must both have
// the GetBucketPolicy permissions on the specified bucket and belong to the
// bucket owner's account in order to use this operation.
//
// If you don't have GetBucketPolicy permissions, Amazon S3 returns a 403 Access
// Denied error. If you have the correct permissions, but you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
//
// To ensure that bucket owners don't inadvertently lock themselves out of their
// own buckets, the root principal in a bucket owner's Amazon Web Services account
// can perform the GetBucketPolicy , PutBucketPolicy , and DeleteBucketPolicy API
// actions, even if their bucket policy explicitly denies the root principal's
// access. Bucket owner root principals can only be blocked from performing these
// API actions by VPC endpoint policies and Amazon Web Services Organizations
// policies.
//
// - General purpose bucket permissions - The s3:GetBucketPolicy permission is
// required in a policy. For more information about general purpose buckets bucket
// policies, see [Using Bucket Policies and User Policies]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:GetBucketPolicy permission in an IAM identity-based
// policy instead of a bucket policy. Cross-account access to this API operation
// isn't supported. This operation can only be performed by the Amazon Web Services
// account that owns the resource. For more information about directory bucket
// policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// Example bucket policies  General purpose buckets example bucket policies - See [Bucket policy examples]
// in the Amazon S3 User Guide.
//
// Directory bucket example bucket policies - See [Example bucket policies for S3 Express One Zone] in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following action is related to GetBucketPolicy :
//
// [GetObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Bucket policy examples]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_GetBucketPolicy(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketPolicyInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Retrieves the policy status for an Amazon S3 bucket, indicating whether the
// bucket is public. In order to use this operation, you must have the
// s3:GetBucketPolicyStatus permission. For more information about Amazon S3
// permissions, see [Specifying Permissions in a Policy].
//
// For more information about when Amazon S3 considers a bucket public, see [The Meaning of "Public"].
//
// The following operations are related to GetBucketPolicyStatus :
//
// [Using Amazon S3 Block Public Access]
//
// [GetPublicAccessBlock]
//
// [PutPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
func s3_GetBucketPolicyStatus(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketPolicyStatusInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketPolicyStatus(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the replication configuration of a bucket.
//
// It can take a while to propagate the put or delete a replication configuration
// to all Amazon S3 systems. Therefore, a get request soon after put or delete can
// return a wrong result.
//
// For information about replication configuration, see [Replication] in the Amazon S3 User
// Guide.
//
// This action requires permissions for the s3:GetReplicationConfiguration action.
// For more information about permissions, see [Using Bucket Policies and User Policies].
//
// If you include the Filter element in a replication configuration, you must also
// include the DeleteMarkerReplication and Priority elements. The response also
// returns those elements.
//
// For information about GetBucketReplication errors, see [List of replication-related error codes]
//
// The following operations are related to GetBucketReplication :
//
// [PutBucketReplication]
//
// [DeleteBucketReplication]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketReplication.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html
// [List of replication-related error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ReplicationErrorCodeList
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html
func s3_GetBucketReplication(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketReplicationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the request payment configuration of a bucket. To use this version of
// the operation, you must be the bucket owner. For more information, see [Requester Pays Buckets].
//
// The following operations are related to GetBucketRequestPayment :
//
// [ListObjects]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
// [Requester Pays Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html
func s3_GetBucketRequestPayment(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketRequestPaymentInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketRequestPayment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the tag set associated with the general purpose bucket.
//
// if ABAC is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use [ListTagsForResource]instead.
//
// To use this operation, you must have permission to perform the
// s3:GetBucketTagging action. By default, the bucket owner has this permission and
// can grant this permission to others.
//
// GetBucketTagging has the following special error:
//
// - Error code: NoSuchTagSet
//
// - Description: There is no tag set associated with the bucket.
//
// The following operations are related to GetBucketTagging :
//
// [PutBucketTagging]
//
// [DeleteBucketTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [enable ABAC for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [ListTagsForResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_ListTagsForResource.html
func s3_GetBucketTagging(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketTaggingInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the versioning state of a bucket.
//
// To retrieve the versioning state of a bucket, you must be the bucket owner.
//
// This implementation also returns the MFA Delete status of the versioning state.
// If the MFA Delete status is enabled , the bucket owner must use an
// authentication device to change the versioning state of the bucket.
//
// The following operations are related to GetBucketVersioning :
//
// [GetObject]
//
// [PutObject]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
func s3_GetBucketVersioning(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketVersioningInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketVersioning(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the website configuration for a bucket. To host website on Amazon S3,
// you can configure a bucket as website by adding a website configuration. For
// more information about hosting websites, see [Hosting Websites on Amazon S3].
//
// This GET action requires the S3:GetBucketWebsite permission. By default, only
// the bucket owner can read the bucket website configuration. However, bucket
// owners can allow other users to read the website configuration by writing a
// bucket policy granting them the S3:GetBucketWebsite permission.
//
// The following operations are related to GetBucketWebsite :
//
// [DeleteBucketWebsite]
//
// [PutBucketWebsite]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketWebsite.html
// [Hosting Websites on Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
// [DeleteBucketWebsite]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketWebsite.html
func s3_GetBucketWebsite(cfg aws.Config, client *s3.Client) {
	input := &s3.GetBucketWebsiteInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetBucketWebsite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves an object from Amazon S3.
// In the GetObject request, specify the full key name for the object.
//
// General purpose buckets - Both the virtual-hosted-style requests and the
// path-style requests are supported. For a virtual hosted-style request example,
// if you have the object photos/2006/February/sample.jpg , specify the object key
// name as /photos/2006/February/sample.jpg . For a path-style request example, if
// you have the object photos/2006/February/sample.jpg in the bucket named
// examplebucket , specify the object key name as
// /examplebucket/photos/2006/February/sample.jpg . For more information about
// request types, see [HTTP Host Header Bucket Specification]in the Amazon S3 User Guide.
//
// Directory buckets - Only virtual-hosted-style requests are supported. For a
// virtual hosted-style request example, if you have the object
// photos/2006/February/sample.jpg in the bucket named
// amzn-s3-demo-bucket--usw2-az1--x-s3 , specify the object key name as
// /photos/2006/February/sample.jpg . Also, when you make requests to this API
// operation, your requests are sent to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://bucket-name.s3express-zone-id.region-code.amazonaws.com/key-name .
// Path-style requests are not supported. For more information about endpoints in
// Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information about
// endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - You must have the required permissions
// in a policy. To use GetObject , you must have the READ access to the object
// (or version). If you grant READ access to the anonymous user, the GetObject
// operation returns the object without using an authorization header. For more
// information, see [Specifying permissions in a policy]in the Amazon S3 User Guide.
//
// # If you include a versionId in your request header, you must have the
//
// s3:GetObjectVersion permission to access a specific version of an object. The
// s3:GetObject permission is not required in this scenario.
//
// # If you request the current version of an object without a specific versionId in
//
// the request header, only the s3:GetObject permission is required. The
// s3:GetObjectVersion permission is not required in this scenario.
//
// # If the object that you request doesn’t exist, the error that Amazon S3 returns
//
// depends on whether you also have the s3:ListBucket permission.
//
// - If you have the s3:ListBucket permission on the bucket, Amazon S3 returns an
// HTTP status code 404 Not Found error.
//
// - If you don’t have the s3:ListBucket permission, Amazon S3 returns an HTTP
// status code 403 Access Denied error.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # If the object is encrypted using SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// Storage classes If the object you are retrieving is stored in the S3 Glacier
// Flexible Retrieval storage class, the S3 Glacier Deep Archive storage class, the
// S3 Intelligent-Tiering Archive Access tier, or the S3 Intelligent-Tiering Deep
// Archive Access tier, before you can retrieve the object you must first restore a
// copy using [RestoreObject]. Otherwise, this operation returns an InvalidObjectState error. For
// information about restoring archived objects, see [Restoring Archived Objects]in the Amazon S3 User Guide.
//
// Directory buckets - Directory buckets only support EXPRESS_ONEZONE (the S3
// Express One Zone storage class) in Availability Zones and ONEZONE_IA (the S3
// One Zone-Infrequent Access storage class) in Dedicated Local Zones. Unsupported
// storage class values won't write a destination object and will respond with the
// HTTP status code 400 Bad Request .
//
// Encryption Encryption request headers, like x-amz-server-side-encryption ,
// should not be sent for the GetObject requests, if your object uses server-side
// encryption with Amazon S3 managed encryption keys (SSE-S3), server-side
// encryption with Key Management Service (KMS) keys (SSE-KMS), or dual-layer
// server-side encryption with Amazon Web Services KMS keys (DSSE-KMS). If you
// include the header in your GetObject requests for the object that uses these
// types of keys, you’ll get an HTTP 400 Bad Request error.
//
// Directory buckets - For directory buckets, there are only two supported options
// for server-side encryption: SSE-S3 and SSE-KMS. SSE-C isn't supported. For more
// information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide.
//
// Overriding response header values through the request There are times when you
// want to override certain response header values of a GetObject response. For
// example, you might override the Content-Disposition response header value
// through your GetObject request.
//
// You can override values for a set of response headers. These modified response
// header values are included only in a successful response, that is, when the HTTP
// status code 200 OK is returned. The headers you can override using the
// following query parameters in the request are a subset of the headers that
// Amazon S3 accepts when you create an object.
//
// The response headers that you can override for the GetObject response are
// Cache-Control , Content-Disposition , Content-Encoding , Content-Language ,
// Content-Type , and Expires .
//
// To override values for a set of response headers in the GetObject response, you
// can use the following query parameters in the request.
//
// - response-cache-control
//
// - response-content-disposition
//
// - response-content-encoding
//
// - response-content-language
//
// - response-content-type
//
// - response-expires
//
// When you use these parameters, you must sign the request by using either an
// Authorization header or a presigned URL. These parameters cannot be used with an
// unsigned (anonymous) request.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to GetObject :
//
// [ListBuckets]
//
// [GetObjectAcl]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [RestoreObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_RestoreObject.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [ListBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html
// [HTTP Host Header Bucket Specification]: https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html#VirtualHostingSpecifyBucket
// [Restoring Archived Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html
// [GetObjectAcl]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html
// [Specifying permissions in a policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_GetObject(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ChecksumMode) > 0 {
		if err := assignInputField(input, "ChecksumMode", _s3ChecksumMode); err != nil {
			log.Errorf("invalid --checksum-mode: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3IfMatch) > 0 {
		input.IfMatch = aws.String(_s3IfMatch)
	}
	if len(_s3IfModifiedSince) > 0 {
		if err := assignInputField(input, "IfModifiedSince", _s3IfModifiedSince); err != nil {
			log.Errorf("invalid --if-modified-since: %s", err.Error())
			return
		}
	}
	if len(_s3IfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_s3IfNoneMatch)
	}
	if len(_s3IfUnmodifiedSince) > 0 {
		if err := assignInputField(input, "IfUnmodifiedSince", _s3IfUnmodifiedSince); err != nil {
			log.Errorf("invalid --if-unmodified-since: %s", err.Error())
			return
		}
	}
	if len(_s3PartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _s3PartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_s3Range) > 0 {
		input.Range = aws.String(_s3Range)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3ResponseCacheControl) > 0 {
		input.ResponseCacheControl = aws.String(_s3ResponseCacheControl)
	}
	if len(_s3ResponseContentDisposition) > 0 {
		input.ResponseContentDisposition = aws.String(_s3ResponseContentDisposition)
	}
	if len(_s3ResponseContentEncoding) > 0 {
		input.ResponseContentEncoding = aws.String(_s3ResponseContentEncoding)
	}
	if len(_s3ResponseContentLanguage) > 0 {
		input.ResponseContentLanguage = aws.String(_s3ResponseContentLanguage)
	}
	if len(_s3ResponseContentType) > 0 {
		input.ResponseContentType = aws.String(_s3ResponseContentType)
	}
	if len(_s3ResponseExpires) > 0 {
		if err := assignInputField(input, "ResponseExpires", _s3ResponseExpires); err != nil {
			log.Errorf("invalid --response-expires: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.GetObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the access control list (ACL) of an object. To use this operation, you
// must have s3:GetObjectAcl permissions or READ_ACP access to the object. For
// more information, see [Mapping of ACL permissions and access policy permissions]in the Amazon S3 User Guide
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// By default, GET returns ACL information about the current version of an object.
// To return ACL information about a different version, use the versionId
// subresource.
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// requests to read ACLs are still supported and return the
// bucket-owner-full-control ACL with the owner being the account that created the
// bucket. For more information, see [Controlling object ownership and disabling ACLs]in the Amazon S3 User Guide.
//
// The following operations are related to GetObjectAcl :
//
// [GetObject]
//
// [GetObjectAttributes]
//
// [DeleteObject]
//
// [PutObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [Mapping of ACL permissions and access policy permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/acl-overview.html#acl-access-policy-permission-mapping
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Controlling object ownership and disabling ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
func s3_GetObjectAcl(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectAclInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.GetObjectAcl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Retrieves all of the metadata from an object without returning the object
// itself. This operation is useful if you're interested only in an object's
// metadata.
//
// GetObjectAttributes combines the functionality of HeadObject and ListParts . All
// of the data returned with both of those individual calls can be returned with a
// single call to GetObjectAttributes .
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - To use GetObjectAttributes , you must
// have READ access to the object.
//
// # The other permissions that you need to use this operation depend on whether the
//
// bucket is versioned and if a version ID is passed in the GetObjectAttributes
// request.
//
// - If you pass a version ID in your request, you need both the
// s3:GetObjectVersion and s3:GetObjectVersionAttributes permissions.
//
// - If you do not pass a version ID in your request, you need the s3:GetObject
// and s3:GetObjectAttributes permissions.
//
// For more information, see [Specifying Permissions in a Policy]in the Amazon S3 User Guide.
//
// # If the object that you request does not exist, the error Amazon S3 returns
//
// depends on whether you also have the s3:ListBucket permission.
//
// - If you have the s3:ListBucket permission on the bucket, Amazon S3 returns an
// HTTP status code 404 Not Found ("no such key") error.
//
// - If you don't have the s3:ListBucket permission, Amazon S3 returns an HTTP
// status code 403 Forbidden ("access denied") error.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// Encryption Encryption request headers, like x-amz-server-side-encryption ,
// should not be sent for HEAD requests if your object uses server-side encryption
// with Key Management Service (KMS) keys (SSE-KMS), dual-layer server-side
// encryption with Amazon Web Services KMS keys (DSSE-KMS), or server-side
// encryption with Amazon S3 managed encryption keys (SSE-S3). The
// x-amz-server-side-encryption header is used when you PUT an object to S3 and
// want to specify the encryption method. If you include this header in a GET
// request for an object that uses these types of keys, you’ll get an HTTP 400 Bad
// Request error. It's because the encryption method can't be changed when you
// retrieve the object.
//
// If you encrypted an object when you stored the object in Amazon S3 by using
// server-side encryption with customer-provided encryption keys (SSE-C), then when
// you retrieve the metadata from the object, you must use the following headers.
// These headers provide the server with the encryption key required to retrieve
// the object's metadata. The headers are:
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// For more information about SSE-C, see [Server-Side Encryption (Using Customer-Provided Encryption Keys)] in the Amazon S3 User Guide.
//
// Directory bucket permissions - For directory buckets, there are only two
// supported options for server-side encryption: server-side encryption with Amazon
// S3 managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ). We recommend that the bucket's default encryption uses
// the desired encryption configuration and you don't override the bucket default
// encryption in your CreateSession requests or PUT object requests. Then, new
// objects are automatically encrypted with the desired encryption settings. For
// more information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide. For more information about
// the encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads].
//
// Versioning  Directory buckets - S3 Versioning isn't enabled and supported for
// directory buckets. For this API operation, only the null value of the version
// ID is supported by directory buckets. You can only specify null to the versionId
// query parameter in the request.
//
// Conditional request headers Consider the following when using request headers:
//
// - If both of the If-Match and If-Unmodified-Since headers are present in the
// request as follows, then Amazon S3 returns the HTTP status code 200 OK and the
// data requested:
//
// - If-Match condition evaluates to true .
//
// - If-Unmodified-Since condition evaluates to false .
//
// For more information about conditional requests, see [RFC 7232].
//
// - If both of the If-None-Match and If-Modified-Since headers are present in
// the request as follows, then Amazon S3 returns the HTTP status code 304 Not
// Modified :
//
// - If-None-Match condition evaluates to false .
//
// - If-Modified-Since condition evaluates to true .
//
// For more information about conditional requests, see [RFC 7232].
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following actions are related to GetObjectAttributes :
//
// [GetObject]
//
// [GetObjectAcl]
//
// [GetObjectLegalHold]
//
// [GetObjectLockConfiguration]
//
// [GetObjectRetention]
//
// [GetObjectTagging]
//
// [HeadObject]
//
// [ListParts]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [GetObjectLegalHold]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLegalHold.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [Server-Side Encryption (Using Customer-Provided Encryption Keys)]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html
// [GetObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [RFC 7232]: https://tools.ietf.org/html/rfc7232
// [HeadObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_HeadObject.html
// [GetObjectLockConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectLockConfiguration.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [GetObjectAcl]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html
// [GetObjectRetention]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectRetention.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_GetObjectAttributes(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectAttributesInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// ObjectAttributes: []types.ObjectAttributes, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ObjectAttributes) > 0 {
		if err := assignInputField(input, "ObjectAttributes", _s3ObjectAttributes); err != nil {
			log.Errorf("invalid --object-attributes: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3MaxParts) > 0 {
		if err := assignInputField(input, "MaxParts", _s3MaxParts); err != nil {
			log.Errorf("invalid --max-parts: %s", err.Error())
			return
		}
	}
	if len(_s3PartNumberMarker) > 0 {
		input.PartNumberMarker = aws.String(_s3PartNumberMarker)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.GetObjectAttributes(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Gets an object's current legal hold status. For more information, see [Locking Objects].
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// The following action is related to GetObjectLegalHold :
//
// [GetObjectAttributes]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func s3_GetObjectLegalHold(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectLegalHoldInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.GetObjectLegalHold(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Gets the Object Lock configuration for a bucket. The rule specified in the
// Object Lock configuration will be applied by default to every new object placed
// in the specified bucket. For more information, see [Locking Objects].
//
// The following action is related to GetObjectLockConfiguration :
//
// [GetObjectAttributes]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func s3_GetObjectLockConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectLockConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetObjectLockConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Retrieves an object's retention settings. For more information, see [Locking Objects].
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// The following action is related to GetObjectRetention :
//
// [GetObjectAttributes]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func s3_GetObjectRetention(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectRetentionInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.GetObjectRetention(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns the tag-set of an object. You send the GET request against the tagging
// subresource associated with the object.
//
// To use this operation, you must have permission to perform the
// s3:GetObjectTagging action. By default, the GET action returns information about
// current version of an object. For a versioned bucket, you can have multiple
// versions of an object in your bucket. To retrieve tags of any other version, use
// the versionId query parameter. You also need permission for the
// s3:GetObjectVersionTagging action.
//
// By default, the bucket owner has this permission and can grant this permission
// to others.
//
// For information about the Amazon S3 object tagging feature, see [Object Tagging].
//
// The following actions are related to GetObjectTagging :
//
// [DeleteObjectTagging]
//
// [GetObjectAttributes]
//
// [PutObjectTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html
// [PutObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObjectTagging.html
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Object Tagging]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-tagging.html
func s3_GetObjectTagging(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectTaggingInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.GetObjectTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns torrent files from a bucket. BitTorrent can save you bandwidth when
// you're distributing large files.
//
// You can get torrent only for objects that are less than 5 GB in size, and that
// are not encrypted using server-side encryption with a customer-provided
// encryption key.
//
// To use GET, you must have READ access to the object.
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// The following action is related to GetObjectTorrent :
//
// [GetObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
func s3_GetObjectTorrent(cfg aws.Config, client *s3.Client) {
	input := &s3.GetObjectTorrentInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}

	if resp, err := client.GetObjectTorrent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Retrieves the PublicAccessBlock configuration for an Amazon S3 bucket. This
// operation returns the bucket-level configuration only. To understand the
// effective public access behavior, you must also consider account-level settings
// (which may inherit from organization-level policies). To use this operation, you
// must have the s3:GetBucketPublicAccessBlock permission. For more information
// about Amazon S3 permissions, see [Specifying Permissions in a Policy].
//
// When Amazon S3 evaluates the PublicAccessBlock configuration for a bucket or an
// object, it checks the PublicAccessBlock configuration for both the bucket (or
// the bucket that contains the object) and the bucket owner's account.
// Account-level settings automatically inherit from organization-level policies
// when present. If the PublicAccessBlock settings are different between the
// bucket and the account, Amazon S3 uses the most restrictive combination of the
// bucket-level and account-level settings.
//
// For more information about when Amazon S3 considers a bucket or an object
// public, see [The Meaning of "Public"].
//
// The following operations are related to GetPublicAccessBlock :
//
// [Using Amazon S3 Block Public Access]
//
// [PutPublicAccessBlock]
//
// [GetPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [PutPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
func s3_GetPublicAccessBlock(cfg aws.Config, client *s3.Client) {
	input := &s3.GetPublicAccessBlockInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.GetPublicAccessBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// You can use this operation to determine if a bucket exists and if you have
// permission to access it. The action returns a 200 OK HTTP status code if the
// bucket exists and you have permission to access it. You can make a HeadBucket
// call on any bucket name to any Region in the partition, and regardless of the
// permissions on the bucket, you will receive a response header with the correct
// bucket location so that you can then make a proper, signed request to the
// appropriate Regional endpoint.
//
// If the bucket doesn't exist or you don't have permission to access it, the HEAD
// request returns a generic 400 Bad Request , 403 Forbidden , or 404 Not Found
// HTTP status code. A message body isn't included, so you can't determine the
// exception beyond these HTTP response codes.
//
// Authentication and authorization  General purpose buckets - Request to public
// buckets that grant the s3:ListBucket permission publicly do not need to be
// signed. All other HeadBucket requests must be authenticated and signed by using
// IAM credentials (access key ID and secret access key for the IAM identities).
// All headers with the x-amz- prefix, including x-amz-copy-source , must be
// signed. For more information, see [REST Authentication].
//
// Directory buckets - You must use IAM credentials to authenticate and authorize
// your access to the HeadBucket API operation, instead of using the temporary
// security credentials through the CreateSession API operation.
//
// Amazon Web Services CLI or SDKs handles authentication and authorization on
// your behalf.
//
// # Permissions
//
// - General purpose bucket permissions - To use this operation, you must have
// permissions to perform the s3:ListBucket action. The bucket owner has this
// permission by default and can grant this permission to others. For more
// information about permissions, see [Managing access permissions to your Amazon S3 resources]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - You must have the s3express:CreateSession
// permission in the Action element of a policy. By default, the session is in
// the ReadWrite mode. If you want to restrict the access, you can explicitly set
// the s3express:SessionMode condition key to ReadOnly on the bucket.
//
// For more information about example bucket policies, see [Example bucket policies for S3 Express One Zone]and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]in the Amazon S3
//
// User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// You must make requests for this API operation to the Zonal endpoint. These
// endpoints support virtual-hosted-style requests in the format
// https://bucket-name.s3express-zone-id.region-code.amazonaws.com . Path-style
// requests are not supported. For more information about endpoints in Availability
// Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information about endpoints in
// Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [REST Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [Managing access permissions to your Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_HeadBucket(cfg aws.Config, client *s3.Client) {
	input := &s3.HeadBucketInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.HeadBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// The HEAD operation retrieves metadata from an object without returning the
// object itself. This operation is useful if you're interested only in an object's
// metadata.
//
// A HEAD request has the same options as a GET operation on an object. The
// response is identical to the GET response except that there is no response
// body. Because of this, if the HEAD request generates an error, it returns a
// generic code, such as 400 Bad Request , 403 Forbidden , 404 Not Found , 405
// Method Not Allowed , 412 Precondition Failed , or 304 Not Modified . It's not
// possible to retrieve the exact exception of these error codes.
//
// Request headers are limited to 8 KB in size. For more information, see [Common Request Headers].
//
// # Permissions
//
// - General purpose bucket permissions - To use HEAD , you must have the
// s3:GetObject permission. You need the relevant read object (or version)
// permission for this operation. For more information, see [Actions, resources, and condition keys for Amazon S3]in the Amazon S3
// User Guide. For more information about the permissions to S3 API operations by
// S3 resource types, see Required permissions for Amazon S3 API operationsin the Amazon S3 User Guide.
//
// If the object you request doesn't exist, the error that Amazon S3 returns
//
// depends on whether you also have the s3:ListBucket permission.
//
// - If you have the s3:ListBucket permission on the bucket, Amazon S3 returns an
// HTTP status code 404 Not Found error.
//
// - If you don’t have the s3:ListBucket permission, Amazon S3 returns an HTTP
// status code 403 Forbidden error.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # If you enable x-amz-checksum-mode in the request and the object is encrypted
//
// with Amazon Web Services Key Management Service (Amazon Web Services KMS), you
// must also have the kms:GenerateDataKey and kms:Decrypt permissions in IAM
// identity-based policies and KMS key policies for the KMS key to retrieve the
// checksum of the object.
//
// Encryption Encryption request headers, like x-amz-server-side-encryption ,
// should not be sent for HEAD requests if your object uses server-side encryption
// with Key Management Service (KMS) keys (SSE-KMS), dual-layer server-side
// encryption with Amazon Web Services KMS keys (DSSE-KMS), or server-side
// encryption with Amazon S3 managed encryption keys (SSE-S3). The
// x-amz-server-side-encryption header is used when you PUT an object to S3 and
// want to specify the encryption method. If you include this header in a HEAD
// request for an object that uses these types of keys, you’ll get an HTTP 400 Bad
// Request error. It's because the encryption method can't be changed when you
// retrieve the object.
//
// If you encrypt an object by using server-side encryption with customer-provided
// encryption keys (SSE-C) when you store the object in Amazon S3, then when you
// retrieve the metadata from the object, you must use the following headers to
// provide the encryption key for the server to be able to retrieve the object's
// metadata. The headers are:
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// For more information about SSE-C, see [Server-Side Encryption (Using Customer-Provided Encryption Keys)] in the Amazon S3 User Guide.
//
// Directory bucket - For directory buckets, there are only two supported options
// for server-side encryption: SSE-S3 and SSE-KMS. SSE-C isn't supported. For more
// information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide.
//
// # Versioning
//
// - If the current version of the object is a delete marker, Amazon S3 behaves
// as if the object was deleted and includes x-amz-delete-marker: true in the
// response.
//
// - If the specified version is a delete marker, the response returns a 405
// Method Not Allowed error and the Last-Modified: timestamp response header.
//
// - Directory buckets - Delete marker is not supported for directory buckets.
//
// - Directory buckets - S3 Versioning isn't enabled and supported for directory
// buckets. For this API operation, only the null value of the version ID is
// supported by directory buckets. You can only specify null to the versionId
// query parameter in the request.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// For directory buckets, you must make requests for this API operation to the
// Zonal endpoint. These endpoints support virtual-hosted-style requests in the
// format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// The following actions are related to HeadObject :
//
// [GetObject]
//
// [GetObjectAttributes]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Server-Side Encryption (Using Customer-Provided Encryption Keys)]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
// [Actions, resources, and condition keys for Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/list_amazons3.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Common Request Headers]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTCommonRequestHeaders.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_HeadObject(cfg aws.Config, client *s3.Client) {
	input := &s3.HeadObjectInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ChecksumMode) > 0 {
		if err := assignInputField(input, "ChecksumMode", _s3ChecksumMode); err != nil {
			log.Errorf("invalid --checksum-mode: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3IfMatch) > 0 {
		input.IfMatch = aws.String(_s3IfMatch)
	}
	if len(_s3IfModifiedSince) > 0 {
		if err := assignInputField(input, "IfModifiedSince", _s3IfModifiedSince); err != nil {
			log.Errorf("invalid --if-modified-since: %s", err.Error())
			return
		}
	}
	if len(_s3IfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_s3IfNoneMatch)
	}
	if len(_s3IfUnmodifiedSince) > 0 {
		if err := assignInputField(input, "IfUnmodifiedSince", _s3IfUnmodifiedSince); err != nil {
			log.Errorf("invalid --if-unmodified-since: %s", err.Error())
			return
		}
	}
	if len(_s3PartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _s3PartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_s3Range) > 0 {
		input.Range = aws.String(_s3Range)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3ResponseCacheControl) > 0 {
		input.ResponseCacheControl = aws.String(_s3ResponseCacheControl)
	}
	if len(_s3ResponseContentDisposition) > 0 {
		input.ResponseContentDisposition = aws.String(_s3ResponseContentDisposition)
	}
	if len(_s3ResponseContentEncoding) > 0 {
		input.ResponseContentEncoding = aws.String(_s3ResponseContentEncoding)
	}
	if len(_s3ResponseContentLanguage) > 0 {
		input.ResponseContentLanguage = aws.String(_s3ResponseContentLanguage)
	}
	if len(_s3ResponseContentType) > 0 {
		input.ResponseContentType = aws.String(_s3ResponseContentType)
	}
	if len(_s3ResponseExpires) > 0 {
		if err := assignInputField(input, "ResponseExpires", _s3ResponseExpires); err != nil {
			log.Errorf("invalid --response-expires: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.HeadObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Lists the analytics configurations for the bucket. You can have up to 1,000
// analytics configurations per bucket.
//
// This action supports list pagination and does not return more than 100
// configurations at a time. You should always check the IsTruncated element in
// the response. If there are no more configurations to list, IsTruncated is set
// to false. If there are more configurations to list, IsTruncated is set to true,
// and there will be a value in NextContinuationToken . You use the
// NextContinuationToken value to continue the pagination of the list by passing
// the value in continuation-token in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the
// s3:GetAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about Amazon S3 analytics feature, see [Amazon S3 Analytics – Storage Class Analysis].
//
// The following operations are related to ListBucketAnalyticsConfigurations :
//
// [GetBucketAnalyticsConfiguration]
//
// [DeleteBucketAnalyticsConfiguration]
//
// [PutBucketAnalyticsConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Analytics – Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html
// [DeleteBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [GetBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAnalyticsConfiguration.html
// [PutBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketAnalyticsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_ListBucketAnalyticsConfigurations(cfg aws.Config, client *s3.Client) {
	input := &s3.ListBucketAnalyticsConfigurationsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.ListBucketAnalyticsConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Lists the S3 Intelligent-Tiering configuration from the specified bucket.
//
// The S3 Intelligent-Tiering storage class is designed to optimize storage costs
// by automatically moving data to the most cost-effective storage access tier,
// without performance impact or operational overhead. S3 Intelligent-Tiering
// delivers automatic cost savings in three low latency and high throughput access
// tiers. To get the lowest storage cost on data that can be accessed in minutes to
// hours, you can choose to activate additional archiving capabilities.
//
// The S3 Intelligent-Tiering storage class is the ideal storage class for data
// with unknown, changing, or unpredictable access patterns, independent of object
// size or retention period. If the size of an object is less than 128 KB, it is
// not monitored and not eligible for auto-tiering. Smaller objects can be stored,
// but they are always charged at the Frequent Access tier rates in the S3
// Intelligent-Tiering storage class.
//
// For more information, see [Storage class for automatically optimizing frequently and infrequently accessed objects].
//
// Operations related to ListBucketIntelligentTieringConfigurations include:
//
// [DeleteBucketIntelligentTieringConfiguration]
//
// [PutBucketIntelligentTieringConfiguration]
//
// [GetBucketIntelligentTieringConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html
// [PutBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketIntelligentTieringConfiguration.html
// [Storage class for automatically optimizing frequently and infrequently accessed objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access
// [DeleteBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html
func s3_ListBucketIntelligentTieringConfigurations(cfg aws.Config, client *s3.Client) {
	input := &s3.ListBucketIntelligentTieringConfigurationsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.ListBucketIntelligentTieringConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns a list of S3 Inventory configurations for the bucket. You can have up
// to 1,000 inventory configurations per bucket.
//
// This action supports list pagination and does not return more than 100
// configurations at a time. Always check the IsTruncated element in the response.
// If there are no more configurations to list, IsTruncated is set to false. If
// there are more configurations to list, IsTruncated is set to true, and there is
// a value in NextContinuationToken . You use the NextContinuationToken value to
// continue the pagination of the list by passing the value in continuation-token
// in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the
// s3:GetInventoryConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about the Amazon S3 inventory feature, see [Amazon S3 Inventory]
//
// The following operations are related to ListBucketInventoryConfigurations :
//
// [GetBucketInventoryConfiguration]
//
// [DeleteBucketInventoryConfiguration]
//
// [PutBucketInventoryConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [PutBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
func s3_ListBucketInventoryConfigurations(cfg aws.Config, client *s3.Client) {
	input := &s3.ListBucketInventoryConfigurationsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.ListBucketInventoryConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Lists the metrics configurations for the bucket. The metrics configurations are
// only for the request metrics of the bucket and do not provide information on
// daily storage metrics. You can have up to 1,000 configurations per bucket.
//
// This action supports list pagination and does not return more than 100
// configurations at a time. Always check the IsTruncated element in the response.
// If there are no more configurations to list, IsTruncated is set to false. If
// there are more configurations to list, IsTruncated is set to true, and there is
// a value in NextContinuationToken . You use the NextContinuationToken value to
// continue the pagination of the list by passing the value in continuation-token
// in the request to GET the next page.
//
// To use this operation, you must have permissions to perform the
// s3:GetMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For more information about metrics configurations and CloudWatch request
// metrics, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to ListBucketMetricsConfigurations :
//
// [PutBucketMetricsConfiguration]
//
// [GetBucketMetricsConfiguration]
//
// [DeleteBucketMetricsConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [GetBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html
// [PutBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketMetricsConfiguration.html
// [DeleteBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_ListBucketMetricsConfigurations(cfg aws.Config, client *s3.Client) {
	input := &s3.ListBucketMetricsConfigurationsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.ListBucketMetricsConfigurations(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns a list of all buckets owned by the authenticated sender of the request.
// To grant IAM permission to use this operation, you must add the
// s3:ListAllMyBuckets policy action.
//
// For information about Amazon S3 buckets, see [Creating, configuring, and working with Amazon S3 buckets].
//
// We strongly recommend using only paginated ListBuckets requests. Unpaginated
// ListBuckets requests are only supported for Amazon Web Services accounts set to
// the default general purpose bucket quota of 10,000. If you have an approved
// general purpose bucket quota above 10,000, you must send paginated ListBuckets
// requests to list your account’s buckets. All unpaginated ListBuckets requests
// will be rejected for Amazon Web Services accounts with a general purpose bucket
// quota greater than 10,000.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Creating, configuring, and working with Amazon S3 buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/creating-buckets-s3.html
func s3_ListBuckets(cfg aws.Config, client *s3.Client) {
	input := &s3.ListBucketsInput{}

	if len(_s3BucketRegion) > 0 {
		input.BucketRegion = aws.String(_s3BucketRegion)
	}
	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3MaxBuckets) > 0 {
		if err := assignInputField(input, "MaxBuckets", _s3MaxBuckets); err != nil {
			log.Errorf("invalid --max-buckets: %s", err.Error())
			return
		}
	}
	if len(_s3Prefix) > 0 {
		input.Prefix = aws.String(_s3Prefix)
	}

	if disablePaginator() {
		if resp, err := client.ListBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3.ListBucketsOutput
	p := s3.NewListBucketsPaginator(client, input)
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

// Returns a list of all Amazon S3 directory buckets owned by the authenticated
// sender of the request. For more information about directory buckets, see [Directory buckets]in the
// Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions You must have the s3express:ListAllMyDirectoryBuckets permission in
// an IAM identity-based policy instead of a bucket policy. Cross-account access to
// this API operation isn't supported. This operation can only be performed by the
// Amazon Web Services account that owns the resource. For more information about
// directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// The BucketRegion response element is not part of the ListDirectoryBuckets
// Response Syntax.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Directory buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-overview.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_ListDirectoryBuckets(cfg aws.Config, client *s3.Client) {
	input := &s3.ListDirectoryBucketsInput{}

	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3MaxDirectoryBuckets) > 0 {
		if err := assignInputField(input, "MaxDirectoryBuckets", _s3MaxDirectoryBuckets); err != nil {
			log.Errorf("invalid --max-directory-buckets: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.ListDirectoryBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3.ListDirectoryBucketsOutput
	p := s3.NewListDirectoryBucketsPaginator(client, input)
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

// This operation lists in-progress multipart uploads in a bucket. An in-progress
// multipart upload is a multipart upload that has been initiated by the
// CreateMultipartUpload request, but has not yet been completed or aborted.
//
// Directory buckets - If multipart uploads in a directory bucket are in progress,
// you can't delete the bucket until all the in-progress multipart uploads are
// aborted or completed. To delete these in-progress multipart uploads, use the
// ListMultipartUploads operation to list the in-progress multipart uploads in the
// bucket and use the AbortMultipartUpload operation to abort all the in-progress
// multipart uploads.
//
// The ListMultipartUploads operation returns a maximum of 1,000 multipart uploads
// in the response. The limit of 1,000 multipart uploads is also the default value.
// You can further limit the number of uploads in a response by specifying the
// max-uploads request parameter. If there are more than 1,000 multipart uploads
// that satisfy your ListMultipartUploads request, the response returns an
// IsTruncated element with the value of true , a NextKeyMarker element, and a
// NextUploadIdMarker element. To list the remaining multipart uploads, you need to
// make subsequent ListMultipartUploads requests. In these requests, include two
// query parameters: key-marker and upload-id-marker . Set the value of key-marker
// to the NextKeyMarker value from the previous response. Similarly, set the value
// of upload-id-marker to the NextUploadIdMarker value from the previous response.
//
// Directory buckets - The upload-id-marker element and the NextUploadIdMarker
// element aren't supported by directory buckets. To list the additional multipart
// uploads, you only need to set the value of key-marker to the NextKeyMarker
// value from the previous response.
//
// For more information about multipart uploads, see [Uploading Objects Using Multipart Upload] in the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// # Permissions
//
// - General purpose bucket permissions - For information about permissions
// required to use the multipart upload API, see [Multipart Upload and Permissions]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # Sorting of multipart uploads in response
//
// - General purpose bucket - In the ListMultipartUploads response, the multipart
// uploads are sorted based on two criteria:
//
// - Key-based sorting - Multipart uploads are initially sorted in ascending
// order based on their object keys.
//
// - Time-based sorting - For uploads that share the same object key, they are
// further sorted in ascending order based on the upload initiation time. Among
// uploads with the same key, the one that was initiated first will appear before
// the ones that were initiated later.
//
// - Directory bucket - In the ListMultipartUploads response, the multipart
// uploads aren't sorted lexicographically based on the object keys.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to ListMultipartUploads :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [ListParts]
//
// [AbortMultipartUpload]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart Upload and Permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
func s3_ListMultipartUploads(cfg aws.Config, client *s3.Client) {
	input := &s3.ListMultipartUploadsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Delimiter) > 0 {
		input.Delimiter = aws.String(_s3Delimiter)
	}
	if len(_s3EncodingType) > 0 {
		if err := assignInputField(input, "EncodingType", _s3EncodingType); err != nil {
			log.Errorf("invalid --encoding-type: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3KeyMarker) > 0 {
		input.KeyMarker = aws.String(_s3KeyMarker)
	}
	if len(_s3MaxUploads) > 0 {
		if err := assignInputField(input, "MaxUploads", _s3MaxUploads); err != nil {
			log.Errorf("invalid --max-uploads: %s", err.Error())
			return
		}
	}
	if len(_s3Prefix) > 0 {
		input.Prefix = aws.String(_s3Prefix)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3UploadIdMarker) > 0 {
		input.UploadIdMarker = aws.String(_s3UploadIdMarker)
	}

	if resp, err := client.ListMultipartUploads(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns metadata about all versions of the objects in a bucket. You can also
// use request parameters as selection criteria to return metadata about a subset
// of all the object versions.
//
// To use this operation, you must have permission to perform the
// s3:ListBucketVersions action. Be aware of the name difference.
//
// A 200 OK response can contain valid or invalid XML. Make sure to design your
// application to parse the contents of the response and handle it appropriately.
//
// To use this operation, you must have READ access to the bucket.
//
// The following operations are related to ListObjectVersions :
//
// [ListObjectsV2]
//
// [GetObject]
//
// [PutObject]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [ListObjectsV2]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html
func s3_ListObjectVersions(cfg aws.Config, client *s3.Client) {
	input := &s3.ListObjectVersionsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Delimiter) > 0 {
		input.Delimiter = aws.String(_s3Delimiter)
	}
	if len(_s3EncodingType) > 0 {
		if err := assignInputField(input, "EncodingType", _s3EncodingType); err != nil {
			log.Errorf("invalid --encoding-type: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3KeyMarker) > 0 {
		input.KeyMarker = aws.String(_s3KeyMarker)
	}
	if len(_s3MaxKeys) > 0 {
		if err := assignInputField(input, "MaxKeys", _s3MaxKeys); err != nil {
			log.Errorf("invalid --max-keys: %s", err.Error())
			return
		}
	}
	if len(_s3OptionalObjectAttributes) > 0 {
		if err := assignInputField(input, "OptionalObjectAttributes", _s3OptionalObjectAttributes); err != nil {
			log.Errorf("invalid --optional-object-attributes: %s", err.Error())
			return
		}
	}
	if len(_s3Prefix) > 0 {
		input.Prefix = aws.String(_s3Prefix)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionIdMarker) > 0 {
		input.VersionIdMarker = aws.String(_s3VersionIdMarker)
	}

	if resp, err := client.ListObjectVersions(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Returns some or all (up to 1,000) of the objects in a bucket. You can use the
// request parameters as selection criteria to return a subset of the objects in a
// bucket. A 200 OK response can contain valid or invalid XML. Be sure to design
// your application to parse the contents of the response and handle it
// appropriately.
//
// This action has been revised. We recommend that you use the newer version, [ListObjectsV2],
// when developing applications. For backward compatibility, Amazon S3 continues to
// support ListObjects .
//
// The following operations are related to ListObjects :
//
// [ListObjectsV2]
//
// [GetObject]
//
// [PutObject]
//
// [CreateBucket]
//
// [ListBuckets]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [ListObjectsV2]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjectsV2.html
func s3_ListObjects(cfg aws.Config, client *s3.Client) {
	input := &s3.ListObjectsInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Delimiter) > 0 {
		input.Delimiter = aws.String(_s3Delimiter)
	}
	if len(_s3EncodingType) > 0 {
		if err := assignInputField(input, "EncodingType", _s3EncodingType); err != nil {
			log.Errorf("invalid --encoding-type: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3Marker) > 0 {
		input.Marker = aws.String(_s3Marker)
	}
	if len(_s3MaxKeys) > 0 {
		if err := assignInputField(input, "MaxKeys", _s3MaxKeys); err != nil {
			log.Errorf("invalid --max-keys: %s", err.Error())
			return
		}
	}
	if len(_s3OptionalObjectAttributes) > 0 {
		if err := assignInputField(input, "OptionalObjectAttributes", _s3OptionalObjectAttributes); err != nil {
			log.Errorf("invalid --optional-object-attributes: %s", err.Error())
			return
		}
	}
	if len(_s3Prefix) > 0 {
		input.Prefix = aws.String(_s3Prefix)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}

	if resp, err := client.ListObjects(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns some or all (up to 1,000) of the objects in a bucket with each request.
// You can use the request parameters as selection criteria to return a subset of
// the objects in a bucket. A 200 OK response can contain valid or invalid XML.
// Make sure to design your application to parse the contents of the response and
// handle it appropriately. For more information about listing objects, see [Listing object keys programmatically]in the
// Amazon S3 User Guide. To get a list of your buckets, see [ListBuckets].
//
// - General purpose bucket - For general purpose buckets, ListObjectsV2 doesn't
// return prefixes that are related only to in-progress multipart uploads.
//
// - Directory buckets - For directory buckets, ListObjectsV2 response includes
// the prefixes that are related only to in-progress multipart uploads.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// # Permissions
//
// - General purpose bucket permissions - To use this operation, you must have
// READ access to the bucket. You must have permission to perform the
// s3:ListBucket action. The bucket owner has this permission by default and can
// grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations]
// and [Managing Access Permissions to Your Amazon S3 Resources]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # Sorting order of returned objects
//
// - General purpose bucket - For general purpose buckets, ListObjectsV2 returns
// objects in lexicographical order based on their key names.
//
// - Directory bucket - For directory buckets, ListObjectsV2 does not return
// objects in lexicographical order.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// This section describes the latest revision of this action. We recommend that
// you use this revised API operation for application development. For backward
// compatibility, Amazon S3 continues to support the prior version of this API
// operation, [ListObjects].
//
// The following operations are related to ListObjectsV2 :
//
// [GetObject]
//
// [PutObject]
//
// [CreateBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListObjects]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListObjects.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Listing object keys programmatically]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ListingKeysUsingAPIs.html
// [ListBuckets]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBuckets.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_ListObjectsV2(cfg aws.Config, client *s3.Client) {
	input := &s3.ListObjectsV2Input{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ContinuationToken) > 0 {
		input.ContinuationToken = aws.String(_s3ContinuationToken)
	}
	if len(_s3Delimiter) > 0 {
		input.Delimiter = aws.String(_s3Delimiter)
	}
	if len(_s3EncodingType) > 0 {
		if err := assignInputField(input, "EncodingType", _s3EncodingType); err != nil {
			log.Errorf("invalid --encoding-type: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3FetchOwner) > 0 {
		if err := assignInputField(input, "FetchOwner", _s3FetchOwner); err != nil {
			log.Errorf("invalid --fetch-owner: %s", err.Error())
			return
		}
	}
	if len(_s3MaxKeys) > 0 {
		if err := assignInputField(input, "MaxKeys", _s3MaxKeys); err != nil {
			log.Errorf("invalid --max-keys: %s", err.Error())
			return
		}
	}
	if len(_s3OptionalObjectAttributes) > 0 {
		if err := assignInputField(input, "OptionalObjectAttributes", _s3OptionalObjectAttributes); err != nil {
			log.Errorf("invalid --optional-object-attributes: %s", err.Error())
			return
		}
	}
	if len(_s3Prefix) > 0 {
		input.Prefix = aws.String(_s3Prefix)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3StartAfter) > 0 {
		input.StartAfter = aws.String(_s3StartAfter)
	}

	if disablePaginator() {
		if resp, err := client.ListObjectsV2(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3.ListObjectsV2Output
	p := s3.NewListObjectsV2Paginator(client, input)
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

// Lists the parts that have been uploaded for a specific multipart upload.
// To use this operation, you must provide the upload ID in the request. You
// obtain this uploadID by sending the initiate multipart upload request through [CreateMultipartUpload].
//
// The ListParts request returns a maximum of 1,000 uploaded parts. The limit of
// 1,000 parts is also the default value. You can restrict the number of parts in a
// response by specifying the max-parts request parameter. If your multipart
// upload consists of more than 1,000 parts, the response returns an IsTruncated
// field with the value of true , and a NextPartNumberMarker element. To list
// remaining uploaded parts, in subsequent ListParts requests, include the
// part-number-marker query string parameter and set its value to the
// NextPartNumberMarker field value from the previous response.
//
// For more information on multipart uploads, see [Uploading Objects Using Multipart Upload] in the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - For information about permissions
// required to use the multipart upload API, see [Multipart Upload and Permissions]in the Amazon S3 User Guide.
//
// # If the upload was created using server-side encryption with Key Management
//
// Service (KMS) keys (SSE-KMS) or dual-layer server-side encryption with Amazon
// Web Services KMS keys (DSSE-KMS), you must have permission to the kms:Decrypt
// action for the ListParts request to succeed.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to ListParts :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [AbortMultipartUpload]
//
// [GetObjectAttributes]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [GetObjectAttributes]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAttributes.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Multipart Upload and Permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
func s3_ListParts(cfg aws.Config, client *s3.Client) {
	input := &s3.ListPartsInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// UploadId: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3UploadId) > 0 {
		input.UploadId = aws.String(_s3UploadId)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3MaxParts) > 0 {
		if err := assignInputField(input, "MaxParts", _s3MaxParts); err != nil {
			log.Errorf("invalid --max-parts: %s", err.Error())
			return
		}
	}
	if len(_s3PartNumberMarker) > 0 {
		input.PartNumberMarker = aws.String(_s3PartNumberMarker)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}

	if disablePaginator() {
		if resp, err := client.ListParts(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3.ListPartsOutput
	p := s3.NewListPartsPaginator(client, input)
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

// Sets the attribute-based access control (ABAC) property of the general purpose
// bucket. You must have s3:PutBucketABAC permission to perform this action. When
// you enable ABAC, you can use tags for access control on your buckets.
// Additionally, when ABAC is enabled, you must use the [TagResource]and [UntagResource] actions to manage
// tags on your buckets. You can nolonger use the [PutBucketTagging]and [DeleteBucketTagging] actions to tag your bucket.
// For more information, see [Enabling ABAC in general purpose buckets].
//
// [PutBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketTagging.html
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [TagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UntagResource.html
// [Enabling ABAC in general purpose buckets]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
func s3_PutBucketAbac(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketAbacInput{
		// AbacStatus: *types.AbacStatus, // Required
		// Bucket: *string, // Required
	}

	if len(_s3AbacStatus) > 0 {
		if err := assignInputField(input, "AbacStatus", _s3AbacStatus); err != nil {
			log.Errorf("invalid --abac-status: %s", err.Error())
			return
		}
	}
	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketAbac(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets the accelerate configuration of an existing bucket. Amazon S3 Transfer
// Acceleration is a bucket-level feature that enables you to perform faster data
// transfers to Amazon S3.
//
// To use this operation, you must have permission to perform the
// s3:PutAccelerateConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// The Transfer Acceleration state of a bucket can be set to one of the following
// two values:
//
// - Enabled – Enables accelerated data transfers to the bucket.
//
// - Suspended – Disables accelerated data transfers to the bucket.
//
// The [GetBucketAccelerateConfiguration] action returns the transfer acceleration state of a bucket.
//
// After setting the Transfer Acceleration state of a bucket to Enabled, it might
// take up to thirty minutes before the data transfer rates to the bucket increase.
//
// The name of the bucket used for Transfer Acceleration must be DNS-compliant and
// must not contain periods (".").
//
// For more information about transfer acceleration, see [Transfer Acceleration].
//
// The following operations are related to PutBucketAccelerateConfiguration :
//
// [GetBucketAccelerateConfiguration]
//
// [CreateBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Transfer Acceleration]: https://docs.aws.amazon.com/AmazonS3/latest/dev/transfer-acceleration.html
// [GetBucketAccelerateConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAccelerateConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
func s3_PutBucketAccelerateConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketAccelerateConfigurationInput{
		// AccelerateConfiguration: *types.AccelerateConfiguration, // Required
		// Bucket: *string, // Required
	}

	if len(_s3AccelerateConfiguration) > 0 {
		if err := assignInputField(input, "AccelerateConfiguration", _s3AccelerateConfiguration); err != nil {
			log.Errorf("invalid --accelerate-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketAccelerateConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// This operation is not supported for directory buckets.
//
// Sets the permissions on an existing bucket using access control lists (ACL).
// For more information, see [Using ACLs]. To set the ACL of a bucket, you must have the
// WRITE_ACP permission.
//
// You can use one of the following two ways to set a bucket's permissions:
//
// - Specify the ACL in the request body
//
// - Specify permissions using request headers
//
// You cannot specify access permission using both the body and the request
// headers.
//
// Depending on your application needs, you may choose to set the ACL on a bucket
// using either the request body or the headers. For example, if you have an
// existing application that updates a bucket ACL using the request body, then you
// can continue to use that approach.
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// ACLs are disabled and no longer affect permissions. You must use policies to
// grant access to your bucket and the objects in it. Requests to set ACLs or
// update ACLs fail and return the AccessControlListNotSupported error code.
// Requests to read ACLs are still supported. For more information, see [Controlling object ownership]in the
// Amazon S3 User Guide.
//
// Permissions You can set access permissions by using one of the following
// methods:
//
// - Specify a canned ACL with the x-amz-acl request header. Amazon S3 supports a
// set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined
// set of grantees and permissions. Specify the canned ACL name as the value of
// x-amz-acl . If you use this header, you cannot use other access
// control-specific headers in your request. For more information, see [Canned ACL].
//
// - Specify access permissions explicitly with the x-amz-grant-read ,
// x-amz-grant-read-acp , x-amz-grant-write-acp , and x-amz-grant-full-control
// headers. When using these headers, you specify explicit access permissions and
// grantees (Amazon Web Services accounts or Amazon S3 groups) who will receive the
// permission. If you use these ACL-specific headers, you cannot use the
// x-amz-acl header to set a canned ACL. These parameters map to the set of
// permissions that Amazon S3 supports in an ACL. For more information, see [Access Control List (ACL) Overview].
//
// You specify each grantee as a type=value pair, where the type is one of the
//
// following:
//
// - id – if the value specified is the canonical user ID of an Amazon Web
// Services account
//
// - uri – if you are granting permissions to a predefined group
//
// - emailAddress – if the value specified is the email address of an Amazon Web
// Services account
//
// # Using email addresses to specify a grantee is only supported in the following
//
// Amazon Web Services Regions:
//
// - US East (N. Virginia)
//
// - US West (N. California)
//
// - US West (Oregon)
//
// - Asia Pacific (Singapore)
//
// - Asia Pacific (Sydney)
//
// - Asia Pacific (Tokyo)
//
// - Europe (Ireland)
//
// - South America (São Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see [Regions and Endpoints]in the
//
// Amazon Web Services General Reference.
//
// For example, the following x-amz-grant-write header grants create, overwrite,
//
// and delete objects permission to LogDelivery group predefined by Amazon S3 and
// two Amazon Web Services accounts identified by their email addresses.
//
// x-amz-grant-write: uri="http://acs.amazonaws.com/groups/s3/LogDelivery",
//
// id="111122223333", id="555566667777"
//
// You can use either a canned ACL or specify access permissions explicitly. You
// cannot do both.
//
// Grantee Values You can specify the person (grantee) to whom you're assigning
// access rights (using request elements) in the following ways. For examples of
// how to specify these grantee values in JSON format, see the Amazon Web Services
// CLI example in [Enabling Amazon S3 server access logging]in the Amazon S3 User Guide.
//
// - By the person's ID:
//
// <>ID<><>GranteesEmail<>
//
// # DisplayName is optional and ignored in the request
//
// - By URI:
//
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// - By Email address:
//
// <>Grantees(at)email.com<>&
//
// # The grantee is resolved to the CanonicalUser and, in a response to a GET Object
//
// acl request, appears as the CanonicalUser.
//
// # Using email addresses to specify a grantee is only supported in the following
//
// Amazon Web Services Regions:
//
// - US East (N. Virginia)
//
// - US West (N. California)
//
// - US West (Oregon)
//
// - Asia Pacific (Singapore)
//
// - Asia Pacific (Sydney)
//
// - Asia Pacific (Tokyo)
//
// - Europe (Ireland)
//
// - South America (São Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see [Regions and Endpoints]in the
//
// Amazon Web Services General Reference.
//
// The following operations are related to PutBucketAcl :
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [GetObjectAcl]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Regions and Endpoints]: https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
// [Access Control List (ACL) Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
// [Controlling object ownership]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [Using ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html
// [Canned ACL]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL
// [GetObjectAcl]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectAcl.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Enabling Amazon S3 server access logging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html
func s3_PutBucketAcl(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketAclInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ACL) > 0 {
		if err := assignInputField(input, "ACL", _s3ACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3AccessControlPolicy) > 0 {
		if err := assignInputField(input, "AccessControlPolicy", _s3AccessControlPolicy); err != nil {
			log.Errorf("invalid --access-control-policy: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3GrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3GrantFullControl)
	}
	if len(_s3GrantRead) > 0 {
		input.GrantRead = aws.String(_s3GrantRead)
	}
	if len(_s3GrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3GrantReadACP)
	}
	if len(_s3GrantWrite) > 0 {
		input.GrantWrite = aws.String(_s3GrantWrite)
	}
	if len(_s3GrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3GrantWriteACP)
	}

	if resp, err := client.PutBucketAcl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets an analytics configuration for the bucket (specified by the analytics
// configuration ID). You can have up to 1,000 analytics configurations per bucket.
//
// You can choose to have storage class analysis export analysis reports sent to a
// comma-separated values (CSV) flat file. See the DataExport request element.
// Reports are updated daily and are based on the object filters that you
// configure. When selecting data export, you specify a destination bucket and an
// optional destination prefix where the file is written. You can export the data
// to a destination bucket in a different account. However, the destination bucket
// must be in the same Region as the bucket that you are making the PUT analytics
// configuration to. For more information, see [Amazon S3 Analytics – Storage Class Analysis].
//
// You must create a bucket policy on the destination bucket where the exported
// file is written to grant permissions to Amazon S3 to write objects to the
// bucket. For an example policy, see [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis].
//
// To use this operation, you must have permissions to perform the
// s3:PutAnalyticsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// PutBucketAnalyticsConfiguration has the following special errors:
//
// - HTTP Error: HTTP 400 Bad Request
//
// - Code: InvalidArgument
//
// - Cause: Invalid argument.
//
// - HTTP Error: HTTP 400 Bad Request
//
// - Code: TooManyConfigurations
//
// - Cause: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// - HTTP Error: HTTP 403 Forbidden
//
// - Code: AccessDenied
//
// - Cause: You are not the owner of the specified bucket, or you do not have
// the s3:PutAnalyticsConfiguration bucket permission to set the configuration on
// the bucket.
//
// The following operations are related to PutBucketAnalyticsConfiguration :
//
// [GetBucketAnalyticsConfiguration]
//
// [DeleteBucketAnalyticsConfiguration]
//
// [ListBucketAnalyticsConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Amazon S3 Analytics – Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html
// [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-9
// [DeleteBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketAnalyticsConfiguration.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [GetBucketAnalyticsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketAnalyticsConfiguration.html
// [ListBucketAnalyticsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketAnalyticsConfigurations.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_PutBucketAnalyticsConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketAnalyticsConfigurationInput{
		// AnalyticsConfiguration: *types.AnalyticsConfiguration, // Required
		// Bucket: *string, // Required
		// Id: *string, // Required
	}

	if len(_s3AnalyticsConfiguration) > 0 {
		if err := assignInputField(input, "AnalyticsConfiguration", _s3AnalyticsConfiguration); err != nil {
			log.Errorf("invalid --analytics-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketAnalyticsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets the cors configuration for your bucket. If the configuration exists,
// Amazon S3 replaces it.
//
// To use this operation, you must be allowed to perform the s3:PutBucketCORS
// action. By default, the bucket owner has this permission and can grant it to
// others.
//
// You set this configuration on a bucket so that the bucket can service
// cross-origin requests. For example, you might want to enable a request whose
// origin is http://www.example.com to access your Amazon S3 bucket at
// my.example.bucket.com by using the browser's XMLHttpRequest capability.
//
// To enable cross-origin resource sharing (CORS) on a bucket, you add the cors
// subresource to the bucket. The cors subresource is an XML document in which you
// configure rules that identify origins and the HTTP methods that can be executed
// on your bucket. The document is limited to 64 KB in size.
//
// When Amazon S3 receives a cross-origin request (or a pre-flight OPTIONS
// request) against a bucket, it evaluates the cors configuration on the bucket
// and uses the first CORSRule rule that matches the incoming browser request to
// enable a cross-origin request. For a rule to match, the following conditions
// must be met:
//
// - The request's Origin header must match AllowedOrigin elements.
//
// - The request method (for example, GET, PUT, HEAD, and so on) or the
// Access-Control-Request-Method header in case of a pre-flight OPTIONS request
// must be one of the AllowedMethod elements.
//
// - Every header specified in the Access-Control-Request-Headers request header
// of a pre-flight request must match an AllowedHeader element.
//
// For more information about CORS, go to [Enabling Cross-Origin Resource Sharing] in the Amazon S3 User Guide.
//
// The following operations are related to PutBucketCors :
//
// [GetBucketCors]
//
// [DeleteBucketCors]
//
// [RESTOPTIONSobject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketCors.html
// [Enabling Cross-Origin Resource Sharing]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cors.html
// [RESTOPTIONSobject]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTOPTIONSobject.html
// [DeleteBucketCors]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketCors.html
func s3_PutBucketCors(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketCorsInput{
		// Bucket: *string, // Required
		// CORSConfiguration: *types.CORSConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3CORSConfiguration) > 0 {
		if err := assignInputField(input, "CORSConfiguration", _s3CORSConfiguration); err != nil {
			log.Errorf("invalid --cors-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketCors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation configures default encryption and Amazon S3 Bucket Keys for an
// existing bucket. You can also [block encryption types]using this operation.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// By default, all buckets have a default encryption configuration that uses
// server-side encryption with Amazon S3 managed keys (SSE-S3).
//
// - General purpose buckets
//
// - You can optionally configure default encryption for a bucket by using
// server-side encryption with Key Management Service (KMS) keys (SSE-KMS) or
// dual-layer server-side encryption with Amazon Web Services KMS keys (DSSE-KMS).
// If you specify default encryption by using SSE-KMS, you can also configure [Amazon S3 Bucket Keys].
// For information about the bucket default encryption feature, see [Amazon S3 Bucket Default Encryption]in the
// Amazon S3 User Guide.
//
// - If you use PutBucketEncryption to set your [default bucket encryption]to SSE-KMS, you should verify
// that your KMS key ID is correct. Amazon S3 doesn't validate the KMS key ID
// provided in PutBucketEncryption requests.
//
// - Directory buckets - You can optionally configure default encryption for a
// bucket by using server-side encryption with Key Management Service (KMS) keys
// (SSE-KMS).
//
// - We recommend that the bucket's default encryption uses the desired
// encryption configuration and you don't override the bucket default encryption in
// your CreateSession requests or PUT object requests. Then, new objects are
// automatically encrypted with the desired encryption settings. For more
// information about the encryption overriding behaviors in directory buckets, see [Specifying server-side encryption with KMS for new object uploads]
// .
//
// - Your SSE-KMS configuration can only support 1 [customer managed key]per directory bucket's
// lifetime. The [Amazon Web Services managed key]( aws/s3 ) isn't supported.
//
// - S3 Bucket Keys are always enabled for GET and PUT operations in a directory
// bucket and can’t be disabled. S3 Bucket Keys aren't supported, when you copy
// SSE-KMS encrypted objects from general purpose buckets to directory buckets,
// from directory buckets to general purpose buckets, or between directory buckets,
// through [CopyObject], [UploadPartCopy], [the Copy operation in Batch Operations], or [the import jobs]. In this case, Amazon S3 makes a call to KMS every time a
// copy request is made for a KMS-encrypted object.
//
// - When you specify an [KMS customer managed key]for encryption in your directory bucket, only use the
// key ID or key ARN. The key alias format of the KMS key isn't supported.
//
// - For directory buckets, if you use PutBucketEncryption to set your [default bucket encryption]to
// SSE-KMS, Amazon S3 validates the KMS key ID provided in PutBucketEncryption
// requests.
//
// If you're specifying a customer managed KMS key, we recommend using a fully
// qualified KMS key ARN. If you use a KMS key alias instead, then KMS resolves the
// key within the requester’s account. This behavior can result in data that's
// encrypted with a KMS key that belongs to the requester, and not the bucket
// owner.
//
// Also, this action requires Amazon Web Services Signature Version 4. For more
// information, see [Authenticating Requests (Amazon Web Services Signature Version 4)].
//
// # Permissions
//
// - General purpose bucket permissions - The s3:PutEncryptionConfiguration
// permission is required in a policy. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Operations]and [Managing Access Permissions to Your Amazon S3 Resources]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:PutEncryptionConfiguration permission in an IAM
// identity-based policy instead of a bucket policy. Cross-account access to this
// API operation isn't supported. This operation can only be performed by the
// Amazon Web Services account that owns the resource. For more information about
// directory bucket policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// # To set a directory bucket default encryption with SSE-KMS, you must also have
//
// the kms:GenerateDataKey and the kms:Decrypt permissions in IAM identity-based
// policies and KMS key policies for the target KMS key.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to PutBucketEncryption :
//
// [GetBucketEncryption]
//
// [DeleteBucketEncryption]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Specifying server-side encryption with KMS for new object uploads]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-specifying-kms-encryption.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [KMS customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
// [Amazon S3 Bucket Default Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Permissions Related to Bucket Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [Amazon Web Services managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#aws-managed-cmk
// [Authenticating Requests (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-authenticating-requests.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [Amazon S3 Bucket Keys]: https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-key.html
// [GetBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketEncryption.html
// [DeleteBucketEncryption]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketEncryption.html
// [customer managed key]: https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#customer-cmk
// [block encryption types]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_BlockedEncryptionTypes.html
// [default bucket encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/bucket-encryption.html
// [the import jobs]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/create-import-job
// [the Copy operation in Batch Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/directory-buckets-objects-Batch-Ops
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_PutBucketEncryption(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketEncryptionInput{
		// Bucket: *string, // Required
		// ServerSideEncryptionConfiguration: *types.ServerSideEncryptionConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ServerSideEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "ServerSideEncryptionConfiguration", _s3ServerSideEncryptionConfiguration); err != nil {
			log.Errorf("invalid --server-side-encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Puts a S3 Intelligent-Tiering configuration to the specified bucket. You can
// have up to 1,000 S3 Intelligent-Tiering configurations per bucket.
//
// The S3 Intelligent-Tiering storage class is designed to optimize storage costs
// by automatically moving data to the most cost-effective storage access tier,
// without performance impact or operational overhead. S3 Intelligent-Tiering
// delivers automatic cost savings in three low latency and high throughput access
// tiers. To get the lowest storage cost on data that can be accessed in minutes to
// hours, you can choose to activate additional archiving capabilities.
//
// The S3 Intelligent-Tiering storage class is the ideal storage class for data
// with unknown, changing, or unpredictable access patterns, independent of object
// size or retention period. If the size of an object is less than 128 KB, it is
// not monitored and not eligible for auto-tiering. Smaller objects can be stored,
// but they are always charged at the Frequent Access tier rates in the S3
// Intelligent-Tiering storage class.
//
// For more information, see [Storage class for automatically optimizing frequently and infrequently accessed objects].
//
// Operations related to PutBucketIntelligentTieringConfiguration include:
//
// [DeleteBucketIntelligentTieringConfiguration]
//
// [GetBucketIntelligentTieringConfiguration]
//
// [ListBucketIntelligentTieringConfigurations]
//
// You only need S3 Intelligent-Tiering enabled on a bucket if you want to
// automatically move objects stored in the S3 Intelligent-Tiering storage class to
// the Archive Access or Deep Archive Access tier.
//
// PutBucketIntelligentTieringConfiguration has the following special errors:
//
// HTTP 400 Bad Request Error  Code: InvalidArgument
//
// Cause: Invalid Argument
//
// HTTP 400 Bad Request Error  Code: TooManyConfigurations
//
// Cause: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// HTTP 403 Forbidden Error  Cause: You are not the owner of the specified bucket,
// or you do not have the s3:PutIntelligentTieringConfiguration bucket permission
// to set the configuration on the bucket.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [ListBucketIntelligentTieringConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketIntelligentTieringConfigurations.html
// [GetBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketIntelligentTieringConfiguration.html
// [Storage class for automatically optimizing frequently and infrequently accessed objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-class-intro.html#sc-dynamic-data-access
// [DeleteBucketIntelligentTieringConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketIntelligentTieringConfiguration.html
func s3_PutBucketIntelligentTieringConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketIntelligentTieringConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
		// IntelligentTieringConfiguration: *types.IntelligentTieringConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3IntelligentTieringConfiguration) > 0 {
		if err := assignInputField(input, "IntelligentTieringConfiguration", _s3IntelligentTieringConfiguration); err != nil {
			log.Errorf("invalid --intelligent-tiering-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketIntelligentTieringConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// This implementation of the PUT action adds an S3 Inventory configuration
// (identified by the inventory ID) to the bucket. You can have up to 1,000
// inventory configurations per bucket.
//
// Amazon S3 inventory generates inventories of the objects in the bucket on a
// daily or weekly basis, and the results are published to a flat file. The bucket
// that is inventoried is called the source bucket, and the bucket where the
// inventory flat file is stored is called the destination bucket. The destination
// bucket must be in the same Amazon Web Services Region as the source bucket.
//
// When you configure an inventory for a source bucket, you specify the
// destination bucket where you want the inventory to be stored, and whether to
// generate the inventory daily or weekly. You can also configure what object
// metadata to include and whether to inventory all object versions or only current
// versions. For more information, see [Amazon S3 Inventory]in the Amazon S3 User Guide.
//
// You must create a bucket policy on the destination bucket to grant permissions
// to Amazon S3 to write objects to the bucket in the defined location. For an
// example policy, see [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis].
//
// Permissions To use this operation, you must have permission to perform the
// s3:PutInventoryConfiguration action. The bucket owner has this permission by
// default and can grant this permission to others.
//
// The s3:PutInventoryConfiguration permission allows a user to create an [S3 Inventory] report
// that includes all object metadata fields available and to specify the
// destination bucket to store the inventory. A user with read access to objects in
// the destination bucket can also access all object metadata fields that are
// available in the inventory report.
//
// To restrict access to an inventory report, see [Restricting access to an Amazon S3 Inventory report] in the Amazon S3 User Guide.
// For more information about the metadata fields available in S3 Inventory, see [Amazon S3 Inventory lists]
// in the Amazon S3 User Guide. For more information about permissions, see [Permissions related to bucket subresource operations]and [Identity and access management in Amazon S3]
// in the Amazon S3 User Guide.
//
// PutBucketInventoryConfiguration has the following special errors:
//
// HTTP 400 Bad Request Error  Code: InvalidArgument
//
// Cause: Invalid Argument
//
// HTTP 400 Bad Request Error  Code: TooManyConfigurations
//
// Cause: You are attempting to create a new configuration but have already
// reached the 1,000-configuration limit.
//
// HTTP 403 Forbidden Error  Cause: You are not the owner of the specified bucket,
// or you do not have the s3:PutInventoryConfiguration bucket permission to set
// the configuration on the bucket.
//
// The following operations are related to PutBucketInventoryConfiguration :
//
// [GetBucketInventoryConfiguration]
//
// [DeleteBucketInventoryConfiguration]
//
// [ListBucketInventoryConfigurations]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Granting Permissions for Amazon S3 Inventory and Storage Class Analysis]: https://docs.aws.amazon.com/AmazonS3/latest/dev/example-bucket-policies.html#example-bucket-policies-use-case-9
// [Amazon S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html
// [ListBucketInventoryConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html
// [S3 Inventory]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html
// [Permissions related to bucket subresource operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html
// [Identity and access management in Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Restricting access to an Amazon S3 Inventory report]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html#example-bucket-policies-use-case-10
// [Amazon S3 Inventory lists]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-inventory.html#storage-inventory-contents
// [GetBucketInventoryConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketInventoryConfiguration.html
func s3_PutBucketInventoryConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketInventoryConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
		// InventoryConfiguration: *types.InventoryConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3InventoryConfiguration) > 0 {
		if err := assignInputField(input, "InventoryConfiguration", _s3InventoryConfiguration); err != nil {
			log.Errorf("invalid --inventory-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketInventoryConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a new lifecycle configuration for the bucket or replaces an existing
// lifecycle configuration. Keep in mind that this will overwrite an existing
// lifecycle configuration, so if you want to retain any configuration details,
// they must be included in the new lifecycle configuration. For information about
// lifecycle configuration, see [Managing your storage lifecycle].
//
// Bucket lifecycle configuration now supports specifying a lifecycle rule using
// an object key name prefix, one or more object tags, object size, or any
// combination of these. Accordingly, this section describes the latest API. The
// previous version of the API supported filtering based only on an object key name
// prefix, which is supported for backward compatibility. For the related API
// description, see [PutBucketLifecycle].
//
// Rules You specify the lifecycle configuration in your request body. The
// lifecycle configuration is specified as XML consisting of one or more rules. An
// Amazon S3 Lifecycle configuration can have up to 1,000 rules. This limit is not
// adjustable.
//
// Bucket lifecycle configuration supports specifying a lifecycle rule using an
// object key name prefix, one or more object tags, object size, or any combination
// of these. Accordingly, this section describes the latest API. The previous
// version of the API supported filtering based only on an object key name prefix,
// which is supported for backward compatibility for general purpose buckets. For
// the related API description, see [PutBucketLifecycle].
//
// Lifecyle configurations for directory buckets only support expiring objects and
// cancelling multipart uploads. Expiring of versioned objects,transitions and tag
// filters are not supported.
//
// A lifecycle rule consists of the following:
//
// - A filter identifying a subset of objects to which the rule applies. The
// filter can be based on a key name prefix, object tags, object size, or any
// combination of these.
//
// - A status indicating whether the rule is in effect.
//
// - One or more lifecycle transition and expiration actions that you want
// Amazon S3 to perform on the objects identified by the filter. If the state of
// your bucket is versioning-enabled or versioning-suspended, you can have many
// versions of the same object (one current version and zero or more noncurrent
// versions). Amazon S3 provides predefined actions that you can specify for
// current and noncurrent object versions.
//
// For more information, see [Object Lifecycle Management] and [Lifecycle Configuration Elements].
//
// Permissions
// - General purpose bucket permissions - By default, all Amazon S3 resources
// are private, including buckets, objects, and related subresources (for example,
// lifecycle configuration and website configuration). Only the resource owner
// (that is, the Amazon Web Services account that created it) can access the
// resource. The resource owner can optionally grant access permissions to others
// by writing an access policy. For this operation, a user must have the
// s3:PutLifecycleConfiguration permission.
//
// You can also explicitly deny permissions. An explicit deny also supersedes any
//
// other permissions. If you want to block users or accounts from removing or
// deleting objects from your bucket, you must deny them permissions for the
// following actions:
//
// - s3:DeleteObject
//
// - s3:DeleteObjectVersion
//
// - s3:PutLifecycleConfiguration
//
// For more information about permissions, see [Managing Access Permissions to Your Amazon S3 Resources].
//
// - Directory bucket permissions - You must have the
// s3express:PutLifecycleConfiguration permission in an IAM identity-based policy
// to use this operation. Cross-account access to this API operation isn't
// supported. The resource owner can optionally grant access permissions to others
// by creating a role or user for them as long as they are within the same account
// as the owner and resource.
//
// For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM]in
//
// the Amazon S3 User Guide.
//
// # Directory buckets - For directory buckets, you must make requests for this API
//
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name
// . Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region.amazonaws.com .
//
// The following operations are related to PutBucketLifecycleConfiguration :
//
// [GetBucketLifecycleConfiguration]
//
// [DeleteBucketLifecycle]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Lifecycle Configuration Elements]: https://docs.aws.amazon.com/AmazonS3/latest/dev/intro-lifecycle-rules.html
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [Authorizing Regional endpoint APIs with IAM]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
// [PutBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycle.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [DeleteBucketLifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketLifecycle.html
// [Managing your storage lifecycle]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lifecycle-mgmt.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
func s3_PutBucketLifecycleConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketLifecycleConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3LifecycleConfiguration) > 0 {
		if err := assignInputField(input, "LifecycleConfiguration", _s3LifecycleConfiguration); err != nil {
			log.Errorf("invalid --lifecycle-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3TransitionDefaultMinimumObjectSize) > 0 {
		if err := assignInputField(input, "TransitionDefaultMinimumObjectSize", _s3TransitionDefaultMinimumObjectSize); err != nil {
			log.Errorf("invalid --transition-default-minimum-object-size: %s", err.Error())
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

// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// This operation is not supported for directory buckets.
//
// Set the logging parameters for a bucket and to specify permissions for who can
// view and modify the logging parameters. All logs are saved to buckets in the
// same Amazon Web Services Region as the source bucket. To set the logging status
// of a bucket, you must be the bucket owner.
//
// The bucket owner is automatically granted FULL_CONTROL to all logs. You use the
// Grantee request element to grant access to other people. The Permissions
// request element specifies the kind of access the grantee has to the logs.
//
// If the target bucket for log delivery uses the bucket owner enforced setting
// for S3 Object Ownership, you can't use the Grantee request element to grant
// access to others. Permissions can only be granted using policies. For more
// information, see [Permissions for server access log delivery]in the Amazon S3 User Guide.
//
// Grantee Values You can specify the person (grantee) to whom you're assigning
// access rights (by using request elements) in the following ways. For examples of
// how to specify these grantee values in JSON format, see the Amazon Web Services
// CLI example in [Enabling Amazon S3 server access logging]in the Amazon S3 User Guide.
//
// - By the person's ID:
//
// <>ID<><>GranteesEmail<>
//
// DisplayName is optional and ignored in the request.
//
// - By Email address:
//
// <>Grantees(at)email.com<>
//
// # The grantee is resolved to the CanonicalUser and, in a response to a
//
// GETObjectAcl request, appears as the CanonicalUser.
//
// - By URI:
//
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// To enable logging, you use LoggingEnabled and its children request elements. To
// disable logging, you use an empty BucketLoggingStatus request element:
//
// For more information about server access logging, see [Server Access Logging] in the Amazon S3 User
// Guide.
//
// For more information about creating a bucket, see [CreateBucket]. For more information about
// returning the logging status of a bucket, see [GetBucketLogging].
//
// The following operations are related to PutBucketLogging :
//
// [PutObject]
//
// [DeleteBucket]
//
// [CreateBucket]
//
// [GetBucketLogging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions for server access log delivery]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html#grant-log-delivery-permissions-general
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [GetBucketLogging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLogging.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Enabling Amazon S3 server access logging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html
// [Server Access Logging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/ServerLogs.html
func s3_PutBucketLogging(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketLoggingInput{
		// Bucket: *string, // Required
		// BucketLoggingStatus: *types.BucketLoggingStatus, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3BucketLoggingStatus) > 0 {
		if err := assignInputField(input, "BucketLoggingStatus", _s3BucketLoggingStatus); err != nil {
			log.Errorf("invalid --bucket-logging-status: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketLogging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets a metrics configuration (specified by the metrics configuration ID) for
// the bucket. You can have up to 1,000 metrics configurations per bucket. If
// you're updating an existing metrics configuration, note that this is a full
// replacement of the existing metrics configuration. If you don't include the
// elements you want to keep, they are erased.
//
// To use this operation, you must have permissions to perform the
// s3:PutMetricsConfiguration action. The bucket owner has this permission by
// default. The bucket owner can grant this permission to others. For more
// information about permissions, see [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// For information about CloudWatch request metrics for Amazon S3, see [Monitoring Metrics with Amazon CloudWatch].
//
// The following operations are related to PutBucketMetricsConfiguration :
//
// [DeleteBucketMetricsConfiguration]
//
// [GetBucketMetricsConfiguration]
//
// [ListBucketMetricsConfigurations]
//
// PutBucketMetricsConfiguration has the following special error:
//
// - Error code: TooManyConfigurations
//
// - Description: You are attempting to create a new configuration but have
// already reached the 1,000-configuration limit.
//
// - HTTP Status Code: HTTP 400 Bad Request
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Monitoring Metrics with Amazon CloudWatch]: https://docs.aws.amazon.com/AmazonS3/latest/dev/cloudwatch-monitoring.html
// [GetBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetricsConfiguration.html
// [ListBucketMetricsConfigurations]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketMetricsConfigurations.html
// [DeleteBucketMetricsConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetricsConfiguration.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_PutBucketMetricsConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketMetricsConfigurationInput{
		// Bucket: *string, // Required
		// Id: *string, // Required
		// MetricsConfiguration: *types.MetricsConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Id) > 0 {
		input.Id = aws.String(_s3Id)
	}
	if len(_s3MetricsConfiguration) > 0 {
		if err := assignInputField(input, "MetricsConfiguration", _s3MetricsConfiguration); err != nil {
			log.Errorf("invalid --metrics-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketMetricsConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Enables notifications of specified events for a bucket. For more information
// about event notifications, see [Configuring Event Notifications].
//
// Using this API, you can replace an existing notification configuration. The
// configuration is an XML file that defines the event types that you want Amazon
// S3 to publish and the destination where you want Amazon S3 to publish an event
// notification when it detects an event of the specified type.
//
// By default, your bucket has no event notifications configured. That is, the
// notification configuration will be an empty NotificationConfiguration .
//
// This action replaces the existing notification configuration with the
// configuration you include in the request body.
//
// After Amazon S3 receives this request, it first verifies that any Amazon Simple
// Notification Service (Amazon SNS) or Amazon Simple Queue Service (Amazon SQS)
// destination exists, and that the bucket owner has permission to publish to it by
// sending a test notification. In the case of Lambda destinations, Amazon S3
// verifies that the Lambda function permissions grant Amazon S3 permission to
// invoke the function from the Amazon S3 bucket. For more information, see [Configuring Notifications for Amazon S3 Events].
//
// You can disable notifications by adding the empty NotificationConfiguration
// element.
//
// For more information about the number of event notification configurations that
// you can create per bucket, see [Amazon S3 service quotas]in Amazon Web Services General Reference.
//
// By default, only the bucket owner can configure notifications on a bucket.
// However, bucket owners can use a bucket policy to grant permission to other
// users to set this configuration with the required s3:PutBucketNotification
// permission.
//
// The PUT notification is an atomic operation. For example, suppose your
// notification configuration includes SNS topic, SQS queue, and Lambda function
// configurations. When you send a PUT request with this configuration, Amazon S3
// sends test messages to your SNS topic. If the message fails, the entire PUT
// action will fail, and Amazon S3 will not add the configuration to your bucket.
//
// If the configuration in the request body includes only one TopicConfiguration
// specifying only the s3:ReducedRedundancyLostObject event type, the response
// will also include the x-amz-sns-test-message-id header containing the message
// ID of the test notification sent to the topic.
//
// The following action is related to PutBucketNotificationConfiguration :
//
// [GetBucketNotificationConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Configuring Notifications for Amazon S3 Events]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
// [Amazon S3 service quotas]: https://docs.aws.amazon.com/general/latest/gr/s3.html#limits_s3
// [GetBucketNotificationConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketNotificationConfiguration.html
// [Configuring Event Notifications]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
func s3_PutBucketNotificationConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketNotificationConfigurationInput{
		// Bucket: *string, // Required
		// NotificationConfiguration: *types.NotificationConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3NotificationConfiguration) > 0 {
		if err := assignInputField(input, "NotificationConfiguration", _s3NotificationConfiguration); err != nil {
			log.Errorf("invalid --notification-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3SkipDestinationValidation) > 0 {
		if err := assignInputField(input, "SkipDestinationValidation", _s3SkipDestinationValidation); err != nil {
			log.Errorf("invalid --skip-destination-validation: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutBucketNotificationConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Creates or modifies OwnershipControls for an Amazon S3 bucket. To use this
// operation, you must have the s3:PutBucketOwnershipControls permission. For more
// information about Amazon S3 permissions, see [Specifying permissions in a policy].
//
// For information about Amazon S3 Object Ownership, see [Using object ownership].
//
// The following operations are related to PutBucketOwnershipControls :
//
// # GetBucketOwnershipControls
//
// # DeleteBucketOwnershipControls
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Specifying permissions in a policy]: https://docs.aws.amazon.com/AmazonS3/latest/user-guide/using-with-s3-actions.html
// [Using object ownership]: https://docs.aws.amazon.com/AmazonS3/latest/user-guide/about-object-ownership.html
func s3_PutBucketOwnershipControls(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketOwnershipControlsInput{
		// Bucket: *string, // Required
		// OwnershipControls: *types.OwnershipControls, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3OwnershipControls) > 0 {
		if err := assignInputField(input, "OwnershipControls", _s3OwnershipControls); err != nil {
			log.Errorf("invalid --ownership-controls: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketOwnershipControls(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies an Amazon S3 bucket policy to an Amazon S3 bucket.
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Regional endpoint. These endpoints support path-style requests
// in the format https://s3express-control.region-code.amazonaws.com/bucket-name .
// Virtual-hosted-style requests aren't supported. For more information about
// endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more
// information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions If you are using an identity other than the root user of the Amazon
// Web Services account that owns the bucket, the calling identity must both have
// the PutBucketPolicy permissions on the specified bucket and belong to the
// bucket owner's account in order to use this operation.
//
// If you don't have PutBucketPolicy permissions, Amazon S3 returns a 403 Access
// Denied error. If you have the correct permissions, but you're not using an
// identity that belongs to the bucket owner's account, Amazon S3 returns a 405
// Method Not Allowed error.
//
// To ensure that bucket owners don't inadvertently lock themselves out of their
// own buckets, the root principal in a bucket owner's Amazon Web Services account
// can perform the GetBucketPolicy , PutBucketPolicy , and DeleteBucketPolicy API
// actions, even if their bucket policy explicitly denies the root principal's
// access. Bucket owner root principals can only be blocked from performing these
// API actions by VPC endpoint policies and Amazon Web Services Organizations
// policies.
//
// - General purpose bucket permissions - The s3:PutBucketPolicy permission is
// required in a policy. For more information about general purpose buckets bucket
// policies, see [Using Bucket Policies and User Policies]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation, you
// must have the s3express:PutBucketPolicy permission in an IAM identity-based
// policy instead of a bucket policy. Cross-account access to this API operation
// isn't supported. This operation can only be performed by the Amazon Web Services
// account that owns the resource. For more information about directory bucket
// policies and permissions, see [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]in the Amazon S3 User Guide.
//
// Example bucket policies  General purpose buckets example bucket policies - See [Bucket policy examples]
// in the Amazon S3 User Guide.
//
// Directory bucket example bucket policies - See [Example bucket policies for S3 Express One Zone] in the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// s3express-control.region-code.amazonaws.com .
//
// The following operations are related to PutBucketPolicy :
//
// [CreateBucket]
//
// [DeleteBucket]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Bucket policy examples]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/example-bucket-policies.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [Using Bucket Policies and User Policies]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-iam-policies.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Amazon Web Services Identity and Access Management (IAM) for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam.html
func s3_PutBucketPolicy(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketPolicyInput{
		// Bucket: *string, // Required
		// Policy: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Policy) > 0 {
		input.Policy = aws.String(_s3Policy)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ConfirmRemoveSelfBucketAccess) > 0 {
		if err := assignInputField(input, "ConfirmRemoveSelfBucketAccess", _s3ConfirmRemoveSelfBucketAccess); err != nil {
			log.Errorf("invalid --confirm-remove-self-bucket-access: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Creates a replication configuration or replaces an existing one. For more
// information, see [Replication]in the Amazon S3 User Guide.
//
// Specify the replication configuration in the request body. In the replication
// configuration, you provide the name of the destination bucket or buckets where
// you want Amazon S3 to replicate objects, the IAM role that Amazon S3 can assume
// to replicate objects on your behalf, and other relevant information. You can
// invoke this request for a specific Amazon Web Services Region by using the [aws:RequestedRegion]
// aws:RequestedRegion condition key.
//
// A replication configuration must include at least one rule, and can contain a
// maximum of 1,000. Each rule identifies a subset of objects to replicate by
// filtering the objects in the source bucket. To choose additional subsets of
// objects to replicate, add a rule for each subset.
//
// To specify a subset of the objects in the source bucket to apply a replication
// rule to, add the Filter element as a child of the Rule element. You can filter
// objects based on an object key prefix, one or more object tags, or both. When
// you add the Filter element in the configuration, you must also add the following
// elements: DeleteMarkerReplication , Status , and Priority .
//
// If you are using an earlier version of the replication configuration, Amazon S3
// handles replication of delete markers differently. For more information, see [Backward Compatibility].
//
// For information about enabling versioning on a bucket, see [Using Versioning].
//
// Handling Replication of Encrypted Objects By default, Amazon S3 doesn't
// replicate objects that are stored at rest using server-side encryption with KMS
// keys. To replicate Amazon Web Services KMS-encrypted objects, add the following:
// SourceSelectionCriteria , SseKmsEncryptedObjects , Status ,
// EncryptionConfiguration , and ReplicaKmsKeyID . For information about
// replication configuration, see [Replicating Objects Created with SSE Using KMS keys].
//
// For information on PutBucketReplication errors, see [List of replication-related error codes]
//
// Permissions To create a PutBucketReplication request, you must have
// s3:PutReplicationConfiguration permissions for the bucket.
//
// By default, a resource owner, in this case the Amazon Web Services account that
// created the bucket, can perform this operation. The resource owner can also
// grant others permissions to perform the operation. For more information about
// permissions, see [Specifying Permissions in a Policy]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// To perform this operation, the user or role performing the action must have the [iam:PassRole]
// permission.
//
// The following operations are related to PutBucketReplication :
//
// [GetBucketReplication]
//
// [DeleteBucketReplication]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [iam:PassRole]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use_passrole.html
// [GetBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketReplication.html
// [aws:RequestedRegion]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_condition-keys.html#condition-keys-requestedregion
// [Replicating Objects Created with SSE Using KMS keys]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-config-for-kms-objects.html
// [Using Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/dev/Versioning.html
// [Replication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication.html
// [List of replication-related error codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#ReplicationErrorCodeList
// [Backward Compatibility]: https://docs.aws.amazon.com/AmazonS3/latest/dev/replication-add-config.html#replication-backward-compat-considerations
// [DeleteBucketReplication]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketReplication.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
func s3_PutBucketReplication(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketReplicationInput{
		// Bucket: *string, // Required
		// ReplicationConfiguration: *types.ReplicationConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ReplicationConfiguration) > 0 {
		if err := assignInputField(input, "ReplicationConfiguration", _s3ReplicationConfiguration); err != nil {
			log.Errorf("invalid --replication-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3Token) > 0 {
		input.Token = aws.String(_s3Token)
	}

	if resp, err := client.PutBucketReplication(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets the request payment configuration for a bucket. By default, the bucket
// owner pays for downloads from the bucket. This configuration parameter enables
// the bucket owner (only) to specify that the person requesting the download will
// be charged for the download. For more information, see [Requester Pays Buckets].
//
// The following operations are related to PutBucketRequestPayment :
//
// [CreateBucket]
//
// [GetBucketRequestPayment]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketRequestPayment]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketRequestPayment.html
// [Requester Pays Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RequesterPaysBuckets.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
func s3_PutBucketRequestPayment(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketRequestPaymentInput{
		// Bucket: *string, // Required
		// RequestPaymentConfiguration: *types.RequestPaymentConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3RequestPaymentConfiguration) > 0 {
		if err := assignInputField(input, "RequestPaymentConfiguration", _s3RequestPaymentConfiguration); err != nil {
			log.Errorf("invalid --request-payment-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketRequestPayment(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets the tags for a general purpose bucket if attribute based access control
// (ABAC) is not enabled for the bucket. When you [enable ABAC for a general purpose bucket], you can no longer use this
// operation for that bucket and must use the [TagResource]or [UntagResource] operations instead.
//
// Use tags to organize your Amazon Web Services bill to reflect your own cost
// structure. To do this, sign up to get your Amazon Web Services account bill with
// tag key values included. Then, to see the cost of combined resources, organize
// your billing information according to resources with the same tag key values.
// For example, you can tag several resources with a specific application name, and
// then organize your billing information to see the total cost of that application
// across several services. For more information, see [Cost Allocation and Tagging]and [Using Cost Allocation in Amazon S3 Bucket Tags].
//
// When this operation sets the tags for a bucket, it will overwrite any current
// tags the bucket already has. You cannot use this operation to add tags to an
// existing list of tags.
//
// To use this operation, you must have permissions to perform the
// s3:PutBucketTagging action. The bucket owner has this permission by default and
// can grant this permission to others. For more information about permissions, see
// [Permissions Related to Bucket Subresource Operations]and [Managing Access Permissions to Your Amazon S3 Resources].
//
// PutBucketTagging has the following special errors. For more Amazon S3 errors
// see, [Error Responses].
//
// - InvalidTag - The tag provided was not a valid tag. This error can occur if
// the tag did not pass input validation. For more information, see [Using Cost Allocation in Amazon S3 Bucket Tags].
//
// - MalformedXML - The XML provided does not match the schema.
//
// - OperationAborted - A conflicting conditional action is currently in progress
// against this resource. Please try again.
//
// - InternalError - The service was unable to apply the provided tag to the
// bucket.
//
// The following operations are related to PutBucketTagging :
//
// [GetBucketTagging]
//
// [DeleteBucketTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Error Responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
// [GetBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketTagging.html
// [Cost Allocation and Tagging]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/cost-alloc-tags.html
// [enable ABAC for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [DeleteBucketTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketTagging.html
// [TagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_TagResource.html
// [UntagResource]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_UntagResource.html
// [Using Cost Allocation in Amazon S3 Bucket Tags]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/CostAllocTagging.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
func s3_PutBucketTagging(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketTaggingInput{
		// Bucket: *string, // Required
		// Tagging: *types.Tagging, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Tagging) > 0 {
		if err := assignInputField(input, "Tagging", _s3Tagging); err != nil {
			log.Errorf("invalid --tagging: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// When you enable versioning on a bucket for the first time, it might take a
// short amount of time for the change to be fully propagated. While this change is
// propagating, you might encounter intermittent HTTP 404 NoSuchKey errors for
// requests to objects created or updated after enabling versioning. We recommend
// that you wait for 15 minutes after enabling versioning before issuing write
// operations ( PUT or DELETE ) on objects in the bucket.
//
// Sets the versioning state of an existing bucket.
//
// You can set the versioning state with one of the following values:
//
// Enabled—Enables versioning for the objects in the bucket. All objects added to
// the bucket receive a unique version ID.
//
// Suspended—Disables versioning for the objects in the bucket. All objects added
// to the bucket receive the version ID null.
//
// If the versioning state has never been set on a bucket, it has no versioning
// state; a [GetBucketVersioning]request does not return a versioning state value.
//
// In order to enable MFA Delete, you must be the bucket owner. If you are the
// bucket owner and want to enable MFA Delete in the bucket versioning
// configuration, you must include the x-amz-mfa request header and the Status and
// the MfaDelete request elements in a request to set the versioning state of the
// bucket.
//
// If you have an object expiration lifecycle configuration in your non-versioned
// bucket and you want to maintain the same permanent delete behavior when you
// enable versioning, you must add a noncurrent expiration policy. The noncurrent
// expiration lifecycle configuration will manage the deletes of the noncurrent
// object versions in the version-enabled bucket. (A version-enabled bucket
// maintains one current and zero or more noncurrent object versions.) For more
// information, see [Lifecycle and Versioning].
//
// The following operations are related to PutBucketVersioning :
//
// [CreateBucket]
//
// [DeleteBucket]
//
// [GetBucketVersioning]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [DeleteBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html
// [CreateBucket]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucket.html
// [Lifecycle and Versioning]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html#lifecycle-and-other-bucket-config
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html
func s3_PutBucketVersioning(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketVersioningInput{
		// Bucket: *string, // Required
		// VersioningConfiguration: *types.VersioningConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3VersioningConfiguration) > 0 {
		if err := assignInputField(input, "VersioningConfiguration", _s3VersioningConfiguration); err != nil {
			log.Errorf("invalid --versioning-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3MFA) > 0 {
		input.MFA = aws.String(_s3MFA)
	}

	if resp, err := client.PutBucketVersioning(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets the configuration of the website that is specified in the website
// subresource. To configure a bucket as a website, you can add this subresource on
// the bucket with website configuration information such as the file name of the
// index document and any redirect rules. For more information, see [Hosting Websites on Amazon S3].
//
// This PUT action requires the S3:PutBucketWebsite permission. By default, only
// the bucket owner can configure the website attached to a bucket; however, bucket
// owners can allow other users to set the website configuration by writing a
// bucket policy that grants them the S3:PutBucketWebsite permission.
//
// To redirect all website requests sent to the bucket's website endpoint, you add
// a website configuration with the following elements. Because all requests are
// sent to another website, you don't need to provide index document name for the
// bucket.
//
// - WebsiteConfiguration
//
// - RedirectAllRequestsTo
//
// - HostName
//
// - Protocol
//
// If you want granular control over redirects, you can use the following elements
// to add routing rules that describe conditions for redirecting requests and
// information about the redirect destination. In this case, the website
// configuration must provide an index document for the bucket, because some
// requests might not be redirected.
//
// - WebsiteConfiguration
//
// - IndexDocument
//
// - Suffix
//
// - ErrorDocument
//
// - Key
//
// - RoutingRules
//
// - RoutingRule
//
// - Condition
//
// - HttpErrorCodeReturnedEquals
//
// - KeyPrefixEquals
//
// - Redirect
//
// - Protocol
//
// - HostName
//
// - ReplaceKeyPrefixWith
//
// - ReplaceKeyWith
//
// - HttpRedirectCode
//
// Amazon S3 has a limitation of 50 routing rules per website configuration. If
// you require more than 50 routing rules, you can use object redirect. For more
// information, see [Configuring an Object Redirect]in the Amazon S3 User Guide.
//
// The maximum request length is limited to 128 KB.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Hosting Websites on Amazon S3]: https://docs.aws.amazon.com/AmazonS3/latest/dev/WebsiteHosting.html
// [Configuring an Object Redirect]: https://docs.aws.amazon.com/AmazonS3/latest/dev/how-to-page-redirect.html
func s3_PutBucketWebsite(cfg aws.Config, client *s3.Client) {
	input := &s3.PutBucketWebsiteInput{
		// Bucket: *string, // Required
		// WebsiteConfiguration: *types.WebsiteConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3WebsiteConfiguration) > 0 {
		if err := assignInputField(input, "WebsiteConfiguration", _s3WebsiteConfiguration); err != nil {
			log.Errorf("invalid --website-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutBucketWebsite(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// Adds an object to a bucket.
//
// - Amazon S3 never adds partial objects; if you receive a success response,
// Amazon S3 added the entire object to the bucket. You cannot use PutObject to
// only update a single piece of metadata for an existing object. You must put the
// entire object with updated metadata if you want to update some values.
//
// - If your bucket uses the bucket owner enforced setting for Object Ownership,
// ACLs are disabled and no longer affect permissions. All objects written to the
// bucket by any account will be owned by the bucket owner.
//
// - Directory buckets - For directory buckets, you must make requests for this
// API operation to the Zonal endpoint. These endpoints support
// virtual-hosted-style requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Amazon S3 is a distributed system. If it receives multiple write requests for
// the same object simultaneously, it overwrites all but the last object written.
// However, Amazon S3 provides features that can modify this behavior:
//
// - S3 Object Lock - To prevent objects from being deleted or overwritten, you
// can use [Amazon S3 Object Lock]in the Amazon S3 User Guide.
//
// This functionality is not supported for directory buckets.
//
// - If-None-Match - Uploads the object only if the object key name does not
// already exist in the specified bucket. Otherwise, Amazon S3 returns a 412
// Precondition Failed error. If a conflicting operation occurs during the
// upload, S3 returns a 409 ConditionalRequestConflict response. On a 409
// failure, retry the upload.
//
// Expects the * character (asterisk).
//
// For more information, see [Add preconditions to S3 operations with conditional requests]in the Amazon S3 User Guide or [RFC 7232].
//
// This functionality is not supported for S3 on Outposts.
//
// - S3 Versioning - When you enable versioning for a bucket, if Amazon S3
// receives multiple write requests for the same object simultaneously, it stores
// all versions of the objects. For each write request that is made to the same
// object, Amazon S3 automatically generates a unique version ID of that object
// being stored in Amazon S3. You can retrieve, replace, or delete any version of
// the object. For more information about versioning, see [Adding Objects to Versioning-Enabled Buckets]in the Amazon S3 User
// Guide. For information about returning the versioning state of a bucket, see [GetBucketVersioning]
// .
//
// This functionality is not supported for directory buckets.
//
// # Permissions
//
// - General purpose bucket permissions - The following permissions are required
// in your policies when your PutObject request includes specific headers.
//
// - s3:PutObject - To successfully complete the PutObject request, you must
// always have the s3:PutObject permission on a bucket to add an object to it.
//
// - s3:PutObjectAcl - To successfully change the objects ACL of your PutObject
// request, you must have the s3:PutObjectAcl .
//
// - s3:PutObjectTagging - To successfully set the tag-set with your PutObject
// request, you must have the s3:PutObjectTagging .
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// # Data integrity with Content-MD5
//
// - General purpose bucket - To ensure that data is not corrupted traversing
// the network, use the Content-MD5 header. When you use this header, Amazon S3
// checks the object against the provided MD5 value and, if they do not match,
// Amazon S3 returns an error. Alternatively, when the object's ETag is its MD5
// digest, you can calculate the MD5 while putting the object to Amazon S3 and
// compare the returned ETag to the calculated MD5 value.
//
// - Directory bucket - This functionality is not supported for directory
// buckets.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// # Errors
//
// - You might receive an InvalidRequest error for several reasons. Depending on
// the reason for the error, you might receive one of the following messages:
//
// - Cannot specify both a write offset value and user-defined object metadata
// for existing objects.
//
// - Checksum Type mismatch occurred, expected checksum Type: sha1, actual
// checksum Type: crc32c.
//
// - Request body cannot be empty when 'write offset' is specified.
//
// For more information about related Amazon S3 APIs, see the following:
//
// [CopyObject]
//
// [DeleteObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [Amazon S3 Object Lock]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock.html
// [DeleteObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObject.html
// [Adding Objects to Versioning-Enabled Buckets]: https://docs.aws.amazon.com/AmazonS3/latest/dev/AddingObjectstoVersioningEnabledBuckets.html
// [Add preconditions to S3 operations with conditional requests]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/conditional-requests.html
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [RFC 7232]: https://datatracker.ietf.org/doc/rfc7232/
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [GetBucketVersioning]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketVersioning.html
func s3_PutObject(cfg aws.Config, client *s3.Client) {
	input := &s3.PutObjectInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ACL) > 0 {
		if err := assignInputField(input, "ACL", _s3ACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3Body) > 0 {
		if err := assignInputField(input, "Body", _s3Body); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_s3BucketKeyEnabled) > 0 {
		if err := assignInputField(input, "BucketKeyEnabled", _s3BucketKeyEnabled); err != nil {
			log.Errorf("invalid --bucket-key-enabled: %s", err.Error())
			return
		}
	}
	if len(_s3CacheControl) > 0 {
		input.CacheControl = aws.String(_s3CacheControl)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumCRC32) > 0 {
		input.ChecksumCRC32 = aws.String(_s3ChecksumCRC32)
	}
	if len(_s3ChecksumCRC32C) > 0 {
		input.ChecksumCRC32C = aws.String(_s3ChecksumCRC32C)
	}
	if len(_s3ChecksumCRC64NVME) > 0 {
		input.ChecksumCRC64NVME = aws.String(_s3ChecksumCRC64NVME)
	}
	if len(_s3ChecksumSHA1) > 0 {
		input.ChecksumSHA1 = aws.String(_s3ChecksumSHA1)
	}
	if len(_s3ChecksumSHA256) > 0 {
		input.ChecksumSHA256 = aws.String(_s3ChecksumSHA256)
	}
	if len(_s3ContentDisposition) > 0 {
		input.ContentDisposition = aws.String(_s3ContentDisposition)
	}
	if len(_s3ContentEncoding) > 0 {
		input.ContentEncoding = aws.String(_s3ContentEncoding)
	}
	if len(_s3ContentLanguage) > 0 {
		input.ContentLanguage = aws.String(_s3ContentLanguage)
	}
	if len(_s3ContentLength) > 0 {
		if err := assignInputField(input, "ContentLength", _s3ContentLength); err != nil {
			log.Errorf("invalid --content-length: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ContentType) > 0 {
		input.ContentType = aws.String(_s3ContentType)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3Expires) > 0 {
		if err := assignInputField(input, "Expires", _s3Expires); err != nil {
			log.Errorf("invalid --expires: %s", err.Error())
			return
		}
	}
	if len(_s3GrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3GrantFullControl)
	}
	if len(_s3GrantRead) > 0 {
		input.GrantRead = aws.String(_s3GrantRead)
	}
	if len(_s3GrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3GrantReadACP)
	}
	if len(_s3GrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3GrantWriteACP)
	}
	if len(_s3IfMatch) > 0 {
		input.IfMatch = aws.String(_s3IfMatch)
	}
	if len(_s3IfNoneMatch) > 0 {
		input.IfNoneMatch = aws.String(_s3IfNoneMatch)
	}
	if len(_s3Metadata) > 0 {
		if err := assignInputField(input, "Metadata", _s3Metadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockLegalHoldStatus) > 0 {
		if err := assignInputField(input, "ObjectLockLegalHoldStatus", _s3ObjectLockLegalHoldStatus); err != nil {
			log.Errorf("invalid --object-lock-legal-hold-status: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockMode) > 0 {
		if err := assignInputField(input, "ObjectLockMode", _s3ObjectLockMode); err != nil {
			log.Errorf("invalid --object-lock-mode: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockRetainUntilDate) > 0 {
		if err := assignInputField(input, "ObjectLockRetainUntilDate", _s3ObjectLockRetainUntilDate); err != nil {
			log.Errorf("invalid --object-lock-retain-until-date: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3SSEKMSEncryptionContext) > 0 {
		input.SSEKMSEncryptionContext = aws.String(_s3SSEKMSEncryptionContext)
	}
	if len(_s3SSEKMSKeyId) > 0 {
		input.SSEKMSKeyId = aws.String(_s3SSEKMSKeyId)
	}
	if len(_s3ServerSideEncryption) > 0 {
		if err := assignInputField(input, "ServerSideEncryption", _s3ServerSideEncryption); err != nil {
			log.Errorf("invalid --server-side-encryption: %s", err.Error())
			return
		}
	}
	if len(_s3StorageClass) > 0 {
		if err := assignInputField(input, "StorageClass", _s3StorageClass); err != nil {
			log.Errorf("invalid --storage-class: %s", err.Error())
			return
		}
	}
	if len(_s3Tagging) > 0 {
		input.Tagging = aws.String(_s3Tagging)
	}
	if len(_s3WebsiteRedirectLocation) > 0 {
		input.WebsiteRedirectLocation = aws.String(_s3WebsiteRedirectLocation)
	}
	if len(_s3WriteOffsetBytes) > 0 {
		if err := assignInputField(input, "WriteOffsetBytes", _s3WriteOffsetBytes); err != nil {
			log.Errorf("invalid --write-offset-bytes: %s", err.Error())
			return
		}
	}

	if resp, err := client.PutObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// End of support notice: As of October 1, 2025, Amazon S3 has discontinued
// support for Email Grantee Access Control Lists (ACLs). If you attempt to use an
// Email Grantee ACL in a request after October 1, 2025, the request will receive
// an HTTP 405 (Method Not Allowed) error.
//
// This change affects the following Amazon Web Services Regions: US East (N.
// Virginia), US West (N. California), US West (Oregon), Asia Pacific (Singapore),
// Asia Pacific (Sydney), Asia Pacific (Tokyo), Europe (Ireland), and South America
// (São Paulo).
//
// This operation is not supported for directory buckets.
//
// Uses the acl subresource to set the access control list (ACL) permissions for a
// new or existing object in an S3 bucket. You must have the WRITE_ACP permission
// to set the ACL of an object. For more information, see [What permissions can I grant?]in the Amazon S3 User
// Guide.
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// Depending on your application needs, you can choose to set the ACL on an object
// using either the request body or the headers. For example, if you have an
// existing application that updates a bucket ACL using the request body, you can
// continue to use that approach. For more information, see [Access Control List (ACL) Overview]in the Amazon S3 User
// Guide.
//
// If your bucket uses the bucket owner enforced setting for S3 Object Ownership,
// ACLs are disabled and no longer affect permissions. You must use policies to
// grant access to your bucket and the objects in it. Requests to set ACLs or
// update ACLs fail and return the AccessControlListNotSupported error code.
// Requests to read ACLs are still supported. For more information, see [Controlling object ownership]in the
// Amazon S3 User Guide.
//
// Permissions You can set access permissions using one of the following methods:
//
// - Specify a canned ACL with the x-amz-acl request header. Amazon S3 supports a
// set of predefined ACLs, known as canned ACLs. Each canned ACL has a predefined
// set of grantees and permissions. Specify the canned ACL name as the value of
// x-amz-ac l. If you use this header, you cannot use other access
// control-specific headers in your request. For more information, see [Canned ACL].
//
// - Specify access permissions explicitly with the x-amz-grant-read ,
// x-amz-grant-read-acp , x-amz-grant-write-acp , and x-amz-grant-full-control
// headers. When using these headers, you specify explicit access permissions and
// grantees (Amazon Web Services accounts or Amazon S3 groups) who will receive the
// permission. If you use these ACL-specific headers, you cannot use x-amz-acl
// header to set a canned ACL. These parameters map to the set of permissions that
// Amazon S3 supports in an ACL. For more information, see [Access Control List (ACL) Overview].
//
// You specify each grantee as a type=value pair, where the type is one of the
//
// following:
//
// - id – if the value specified is the canonical user ID of an Amazon Web
// Services account
//
// - uri – if you are granting permissions to a predefined group
//
// - emailAddress – if the value specified is the email address of an Amazon Web
// Services account
//
// # Using email addresses to specify a grantee is only supported in the following
//
// Amazon Web Services Regions:
//
// - US East (N. Virginia)
//
// - US West (N. California)
//
// - US West (Oregon)
//
// - Asia Pacific (Singapore)
//
// - Asia Pacific (Sydney)
//
// - Asia Pacific (Tokyo)
//
// - Europe (Ireland)
//
// - South America (São Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see [Regions and Endpoints]in the
//
// Amazon Web Services General Reference.
//
// # For example, the following x-amz-grant-read header grants list objects
//
// permission to the two Amazon Web Services accounts identified by their email
// addresses.
//
// x-amz-grant-read: emailAddress="xyz(at)amazon.com", emailAddress="abc(at)amazon.com"
//
// You can use either a canned ACL or specify access permissions explicitly. You
// cannot do both.
//
// Grantee Values You can specify the person (grantee) to whom you're assigning
// access rights (using request elements) in the following ways. For examples of
// how to specify these grantee values in JSON format, see the Amazon Web Services
// CLI example in [Enabling Amazon S3 server access logging]in the Amazon S3 User Guide.
//
// - By the person's ID:
//
// <>ID<><>GranteesEmail<>
//
// DisplayName is optional and ignored in the request.
//
// - By URI:
//
// <>http://acs.amazonaws.com/groups/global/AuthenticatedUsers<>
//
// - By Email address:
//
// <>Grantees(at)email.com<>lt;/Grantee>
//
// # The grantee is resolved to the CanonicalUser and, in a response to a GET Object
//
// acl request, appears as the CanonicalUser.
//
// # Using email addresses to specify a grantee is only supported in the following
//
// Amazon Web Services Regions:
//
// - US East (N. Virginia)
//
// - US West (N. California)
//
// - US West (Oregon)
//
// - Asia Pacific (Singapore)
//
// - Asia Pacific (Sydney)
//
// - Asia Pacific (Tokyo)
//
// - Europe (Ireland)
//
// - South America (São Paulo)
//
// For a list of all the Amazon S3 supported Regions and endpoints, see [Regions and Endpoints]in the
//
// Amazon Web Services General Reference.
//
// Versioning The ACL of an object is set at the object version level. By default,
// PUT sets the ACL of the current version of an object. To set the ACL of a
// different version, use the versionId subresource.
//
// The following operations are related to PutObjectAcl :
//
// [CopyObject]
//
// [GetObject]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Regions and Endpoints]: https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region
// [Access Control List (ACL) Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html
// [Controlling object ownership]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/about-object-ownership.html
// [Canned ACL]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#CannedACL
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [What permissions can I grant?]: https://docs.aws.amazon.com/AmazonS3/latest/dev/acl-overview.html#permissions
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Enabling Amazon S3 server access logging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/enable-server-access-logging.html
func s3_PutObjectAcl(cfg aws.Config, client *s3.Client) {
	input := &s3.PutObjectAclInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ACL) > 0 {
		if err := assignInputField(input, "ACL", _s3ACL); err != nil {
			log.Errorf("invalid --acl: %s", err.Error())
			return
		}
	}
	if len(_s3AccessControlPolicy) > 0 {
		if err := assignInputField(input, "AccessControlPolicy", _s3AccessControlPolicy); err != nil {
			log.Errorf("invalid --access-control-policy: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3GrantFullControl) > 0 {
		input.GrantFullControl = aws.String(_s3GrantFullControl)
	}
	if len(_s3GrantRead) > 0 {
		input.GrantRead = aws.String(_s3GrantRead)
	}
	if len(_s3GrantReadACP) > 0 {
		input.GrantReadACP = aws.String(_s3GrantReadACP)
	}
	if len(_s3GrantWrite) > 0 {
		input.GrantWrite = aws.String(_s3GrantWrite)
	}
	if len(_s3GrantWriteACP) > 0 {
		input.GrantWriteACP = aws.String(_s3GrantWriteACP)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.PutObjectAcl(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Applies a legal hold configuration to the specified object. For more
// information, see [Locking Objects].
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func s3_PutObjectLegalHold(cfg aws.Config, client *s3.Client) {
	input := &s3.PutObjectLegalHoldInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3LegalHold) > 0 {
		if err := assignInputField(input, "LegalHold", _s3LegalHold); err != nil {
			log.Errorf("invalid --legal-hold: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.PutObjectLegalHold(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Places an Object Lock configuration on the specified bucket. The rule specified
// in the Object Lock configuration will be applied by default to every new object
// placed in the specified bucket. For more information, see [Locking Objects].
//
// - The DefaultRetention settings require both a mode and a period.
//
// - The DefaultRetention period can be either Days or Years but you must select
// one. You cannot specify Days and Years at the same time.
//
// - You can enable Object Lock for new or existing buckets. For more
// information, see [Configuring Object Lock].
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Configuring Object Lock]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-lock-configure.html
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func s3_PutObjectLockConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.PutObjectLockConfigurationInput{
		// Bucket: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3ObjectLockConfiguration) > 0 {
		if err := assignInputField(input, "ObjectLockConfiguration", _s3ObjectLockConfiguration); err != nil {
			log.Errorf("invalid --object-lock-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3Token) > 0 {
		input.Token = aws.String(_s3Token)
	}

	if resp, err := client.PutObjectLockConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Places an Object Retention configuration on an object. For more information,
// see [Locking Objects]. Users or accounts require the s3:PutObjectRetention permission in order
// to place an Object Retention configuration on objects. Bypassing a Governance
// Retention configuration requires the s3:BypassGovernanceRetention permission.
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Locking Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lock.html
func s3_PutObjectRetention(cfg aws.Config, client *s3.Client) {
	input := &s3.PutObjectRetentionInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3BypassGovernanceRetention) > 0 {
		if err := assignInputField(input, "BypassGovernanceRetention", _s3BypassGovernanceRetention); err != nil {
			log.Errorf("invalid --bypass-governance-retention: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3Retention) > 0 {
		if err := assignInputField(input, "Retention", _s3Retention); err != nil {
			log.Errorf("invalid --retention: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.PutObjectRetention(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Sets the supplied tag-set to an object that already exists in a bucket. A tag
// is a key-value pair. For more information, see [Object Tagging].
//
// You can associate tags with an object by sending a PUT request against the
// tagging subresource that is associated with the object. You can retrieve tags by
// sending a GET request. For more information, see [GetObjectTagging].
//
// For tagging-related restrictions related to characters and encodings, see [Tag Restrictions].
// Note that Amazon S3 limits the maximum number of tags to 10 tags per object.
//
// To use this operation, you must have permission to perform the
// s3:PutObjectTagging action. By default, the bucket owner has this permission and
// can grant this permission to others.
//
// To put tags of any other version, use the versionId query parameter. You also
// need permission for the s3:PutObjectVersionTagging action.
//
// PutObjectTagging has the following special errors. For more Amazon S3 errors
// see, [Error Responses].
//
// - InvalidTag - The tag provided was not a valid tag. This error can occur if
// the tag did not pass input validation. For more information, see [Object Tagging].
//
// - MalformedXML - The XML provided does not match the schema.
//
// - OperationAborted - A conflicting conditional action is currently in progress
// against this resource. Please try again.
//
// - InternalError - The service was unable to apply the provided tag to the
// object.
//
// The following operations are related to PutObjectTagging :
//
// [GetObjectTagging]
//
// [DeleteObjectTagging]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Error Responses]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html
// [DeleteObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteObjectTagging.html
// [Object Tagging]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/object-tagging.html
// [Tag Restrictions]: https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/allocation-tag-restrictions.html
// [GetObjectTagging]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObjectTagging.html
func s3_PutObjectTagging(cfg aws.Config, client *s3.Client) {
	input := &s3.PutObjectTaggingInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// Tagging: *types.Tagging, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3Tagging) > 0 {
		if err := assignInputField(input, "Tagging", _s3Tagging); err != nil {
			log.Errorf("invalid --tagging: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.PutObjectTagging(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Creates or modifies the PublicAccessBlock configuration for an Amazon S3
// bucket. To use this operation, you must have the s3:PutBucketPublicAccessBlock
// permission. For more information about Amazon S3 permissions, see [Specifying Permissions in a Policy].
//
// When Amazon S3 evaluates the PublicAccessBlock configuration for a bucket or an
// object, it checks the PublicAccessBlock configuration for both the bucket (or
// the bucket that contains the object) and the bucket owner's account.
// Account-level settings automatically inherit from organization-level policies
// when present. If the PublicAccessBlock configurations are different between the
// bucket and the account, Amazon S3 uses the most restrictive combination of the
// bucket-level and account-level settings.
//
// For more information about when Amazon S3 considers a bucket or an object
// public, see [The Meaning of "Public"].
//
// The following operations are related to PutPublicAccessBlock :
//
// [GetPublicAccessBlock]
//
// [DeletePublicAccessBlock]
//
// [GetBucketPolicyStatus]
//
// [Using Amazon S3 Block Public Access]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetPublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetPublicAccessBlock.html
// [DeletePublicAccessBlock]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeletePublicAccessBlock.html
// [Using Amazon S3 Block Public Access]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html
// [GetBucketPolicyStatus]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketPolicyStatus.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [The Meaning of "Public"]: https://docs.aws.amazon.com/AmazonS3/latest/dev/access-control-block-public-access.html#access-control-block-public-access-policy-status
func s3_PutPublicAccessBlock(cfg aws.Config, client *s3.Client) {
	input := &s3.PutPublicAccessBlockInput{
		// Bucket: *string, // Required
		// PublicAccessBlockConfiguration: *types.PublicAccessBlockConfiguration, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3PublicAccessBlockConfiguration) > 0 {
		if err := assignInputField(input, "PublicAccessBlockConfiguration", _s3PublicAccessBlockConfiguration); err != nil {
			log.Errorf("invalid --public-access-block-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.PutPublicAccessBlock(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Renames an existing object in a directory bucket that uses the S3 Express One
// Zone storage class. You can use RenameObject by specifying an existing object’s
// name as the source and the new name of the object as the destination within the
// same directory bucket.
//
// RenameObject is only supported for objects stored in the S3 Express One Zone
// storage class.
//
// To prevent overwriting an object, you can use the If-None-Match conditional
// header.
//
// - If-None-Match - Renames the object only if an object with the specified
// name does not already exist in the directory bucket. If you don't want to
// overwrite an existing object, you can add the If-None-Match conditional header
// with the value ‘*’ in the RenameObject request. Amazon S3 then returns a 412
// Precondition Failed error if the object with the specified name already
// exists. For more information, see [RFC 7232].
//
// Permissions  To grant access to the RenameObject operation on a directory
// bucket, we recommend that you use the CreateSession operation for session-based
// authorization. Specifically, you grant the s3express:CreateSession permission
// to the directory bucket in a bucket policy or an IAM identity-based policy.
// Then, you make the CreateSession API call on the directory bucket to obtain a
// session token. With the session token in your request header, you can make API
// requests to this operation. After the session token expires, you make another
// CreateSession API call to generate a new session token for use. The Amazon Web
// Services CLI and SDKs will create and manage your session including refreshing
// the session token automatically to avoid service interruptions when a session
// expires. In your bucket policy, you can specify the s3express:SessionMode
// condition key to control who can create a ReadWrite or ReadOnly session. A
// ReadWrite session is required for executing all the Zonal endpoint API
// operations, including RenameObject . For more information about authorization,
// see [CreateSession]CreateSession . To learn more about Zonal endpoint API operations, see [Authorizing Zonal endpoint API operations with CreateSession] in
// the Amazon S3 User Guide.
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [RFC 7232]: https://datatracker.ietf.org/doc/rfc7232/
// [Authorizing Zonal endpoint API operations with CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-create-session.html
func s3_RenameObject(cfg aws.Config, client *s3.Client) {
	input := &s3.RenameObjectInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// RenameSource: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3RenameSource) > 0 {
		input.RenameSource = aws.String(_s3RenameSource)
	}
	if len(_s3ClientToken) > 0 {
		input.ClientToken = aws.String(_s3ClientToken)
	}
	if len(_s3DestinationIfMatch) > 0 {
		input.DestinationIfMatch = aws.String(_s3DestinationIfMatch)
	}
	if len(_s3DestinationIfModifiedSince) > 0 {
		if err := assignInputField(input, "DestinationIfModifiedSince", _s3DestinationIfModifiedSince); err != nil {
			log.Errorf("invalid --destination-if-modified-since: %s", err.Error())
			return
		}
	}
	if len(_s3DestinationIfNoneMatch) > 0 {
		input.DestinationIfNoneMatch = aws.String(_s3DestinationIfNoneMatch)
	}
	if len(_s3DestinationIfUnmodifiedSince) > 0 {
		if err := assignInputField(input, "DestinationIfUnmodifiedSince", _s3DestinationIfUnmodifiedSince); err != nil {
			log.Errorf("invalid --destination-if-unmodified-since: %s", err.Error())
			return
		}
	}
	if len(_s3SourceIfMatch) > 0 {
		input.SourceIfMatch = aws.String(_s3SourceIfMatch)
	}
	if len(_s3SourceIfModifiedSince) > 0 {
		if err := assignInputField(input, "SourceIfModifiedSince", _s3SourceIfModifiedSince); err != nil {
			log.Errorf("invalid --source-if-modified-since: %s", err.Error())
			return
		}
	}
	if len(_s3SourceIfNoneMatch) > 0 {
		input.SourceIfNoneMatch = aws.String(_s3SourceIfNoneMatch)
	}
	if len(_s3SourceIfUnmodifiedSince) > 0 {
		if err := assignInputField(input, "SourceIfUnmodifiedSince", _s3SourceIfUnmodifiedSince); err != nil {
			log.Errorf("invalid --source-if-unmodified-since: %s", err.Error())
			return
		}
	}

	if resp, err := client.RenameObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// # Restores an archived copy of an object back into Amazon S3
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// This action performs the following types of requests:
//
// - restore an archive - Restore an archived object
//
// For more information about the S3 structure in the request body, see the
// following:
//
// [PutObject]
//
// [Managing Access with ACLs]
// - in the Amazon S3 User Guide
//
// [Protecting Data Using Server-Side Encryption]
// - in the Amazon S3 User Guide
//
// Permissions To use this operation, you must have permissions to perform the
// s3:RestoreObject action. The bucket owner has this permission by default and can
// grant this permission to others. For more information about permissions, see [Permissions Related to Bucket Subresource Operations]
// and [Managing Access Permissions to Your Amazon S3 Resources]in the Amazon S3 User Guide.
//
// Restoring objects Objects that you archive to the S3 Glacier Flexible Retrieval
// or S3 Glacier Deep Archive storage class, and S3 Intelligent-Tiering Archive or
// S3 Intelligent-Tiering Deep Archive tiers, are not accessible in real time. For
// objects in the S3 Glacier Flexible Retrieval or S3 Glacier Deep Archive storage
// classes, you must first initiate a restore request, and then wait until a
// temporary copy of the object is available. If you want a permanent copy of the
// object, create a copy of it in the Amazon S3 Standard storage class in your S3
// bucket. To access an archived object, you must restore the object for the
// duration (number of days) that you specify. For objects in the Archive Access or
// Deep Archive Access tiers of S3 Intelligent-Tiering, you must first initiate a
// restore request, and then wait until the object is moved into the Frequent
// Access tier.
//
// To restore a specific object version, you can provide a version ID. If you
// don't provide a version ID, Amazon S3 restores the current version.
//
// When restoring an archived object, you can specify one of the following data
// access tier options in the Tier element of the request body:
//
// - Expedited - Expedited retrievals allow you to quickly access your data
// stored in the S3 Glacier Flexible Retrieval storage class or S3
// Intelligent-Tiering Archive tier when occasional urgent requests for restoring
// archives are required. For all but the largest archived objects (250 MB+), data
// accessed using Expedited retrievals is typically made available within 1–5
// minutes. Provisioned capacity ensures that retrieval capacity for Expedited
// retrievals is available when you need it. Expedited retrievals and provisioned
// capacity are not available for objects stored in the S3 Glacier Deep Archive
// storage class or S3 Intelligent-Tiering Deep Archive tier.
//
// - Standard - Standard retrievals allow you to access any of your archived
// objects within several hours. This is the default option for retrieval requests
// that do not specify the retrieval option. Standard retrievals typically finish
// within 3–5 hours for objects stored in the S3 Glacier Flexible Retrieval storage
// class or S3 Intelligent-Tiering Archive tier. They typically finish within 12
// hours for objects stored in the S3 Glacier Deep Archive storage class or S3
// Intelligent-Tiering Deep Archive tier. Standard retrievals are free for objects
// stored in S3 Intelligent-Tiering.
//
// - Bulk - Bulk retrievals free for objects stored in the S3 Glacier Flexible
// Retrieval and S3 Intelligent-Tiering storage classes, enabling you to retrieve
// large amounts, even petabytes, of data at no cost. Bulk retrievals typically
// finish within 5–12 hours for objects stored in the S3 Glacier Flexible Retrieval
// storage class or S3 Intelligent-Tiering Archive tier. Bulk retrievals are also
// the lowest-cost retrieval option when restoring objects from S3 Glacier Deep
// Archive. They typically finish within 48 hours for objects stored in the S3
// Glacier Deep Archive storage class or S3 Intelligent-Tiering Deep Archive tier.
//
// For more information about archive retrieval options and provisioned capacity
// for Expedited data access, see [Restoring Archived Objects] in the Amazon S3 User Guide.
//
// You can use Amazon S3 restore speed upgrade to change the restore speed to a
// faster speed while it is in progress. For more information, see [Upgrading the speed of an in-progress restore]in the Amazon
// S3 User Guide.
//
// To get the status of object restoration, you can send a HEAD request.
// Operations return the x-amz-restore header, which provides information about
// the restoration status, in the response. You can use Amazon S3 event
// notifications to notify you when a restore is initiated or completed. For more
// information, see [Configuring Amazon S3 Event Notifications]in the Amazon S3 User Guide.
//
// After restoring an archived object, you can update the restoration period by
// reissuing the request with a new period. Amazon S3 updates the restoration
// period relative to the current time and charges only for the request-there are
// no data transfer charges. You cannot update the restoration period when Amazon
// S3 is actively processing your current restore request for the object.
//
// If your bucket has a lifecycle configuration with a rule that includes an
// expiration action, the object expiration overrides the life span that you
// specify in a restore request. For example, if you restore an object copy for 10
// days, but the object is scheduled to expire in 3 days, Amazon S3 deletes the
// object in 3 days. For more information about lifecycle configuration, see [PutBucketLifecycleConfiguration]and [Object Lifecycle Management]
// in Amazon S3 User Guide.
//
// Responses A successful action returns either the 200 OK or 202 Accepted status
// code.
//
// - If the object is not previously restored, then Amazon S3 returns 202
// Accepted in the response.
//
// - If the object is previously restored, Amazon S3 returns 200 OK in the
// response.
//
// - Special errors:
//
// - Code: RestoreAlreadyInProgress
//
// - Cause: Object restore is already in progress.
//
// - HTTP Status Code: 409 Conflict
//
// - SOAP Fault Code Prefix: Client
//
// - Code: GlacierExpeditedRetrievalNotAvailable
//
// - Cause: expedited retrievals are currently not available. Try again later.
// (Returned if there is insufficient capacity to process the Expedited request.
// This error applies only to Expedited retrievals and not to S3 Standard or Bulk
// retrievals.)
//
// - HTTP Status Code: 503
//
// - SOAP Fault Code Prefix: N/A
//
// The following operations are related to RestoreObject :
//
// [PutBucketLifecycleConfiguration]
//
// [GetBucketNotificationConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [Object Lifecycle Management]: https://docs.aws.amazon.com/AmazonS3/latest/dev/object-lifecycle-mgmt.html
// [Permissions Related to Bucket Subresource Operations]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources
// [Configuring Amazon S3 Event Notifications]: https://docs.aws.amazon.com/AmazonS3/latest/dev/NotificationHowTo.html
// [Managing Access with ACLs]: https://docs.aws.amazon.com/AmazonS3/latest/dev/S3_ACLs_UsingACLs.html
// [Protecting Data Using Server-Side Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html
// [GetBucketNotificationConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketNotificationConfiguration.html
// [PutObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutObject.html
// [Restoring Archived Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html
// [Managing Access Permissions to Your Amazon S3 Resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-access-control.html
// [Upgrading the speed of an in-progress restore]: https://docs.aws.amazon.com/AmazonS3/latest/dev/restoring-objects.html#restoring-objects-upgrade-tier.title.html
func s3_RestoreObject(cfg aws.Config, client *s3.Client) {
	input := &s3.RestoreObjectInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3RestoreRequest) > 0 {
		if err := assignInputField(input, "RestoreRequest", _s3RestoreRequest); err != nil {
			log.Errorf("invalid --restore-request: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.RestoreObject(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// This action filters the contents of an Amazon S3 object based on a simple
// structured query language (SQL) statement. In the request, along with the SQL
// expression, you must also specify a data serialization format (JSON, CSV, or
// Apache Parquet) of the object. Amazon S3 uses this format to parse object data
// into records, and returns only records that match the specified SQL expression.
// You must also specify the data serialization format for the response.
//
// This functionality is not supported for Amazon S3 on Outposts.
//
// For more information about Amazon S3 Select, see [Selecting Content from Objects] and [SELECT Command] in the Amazon S3 User
// Guide.
//
// Permissions You must have the s3:GetObject permission for this operation.
// Amazon S3 Select does not support anonymous access. For more information about
// permissions, see [Specifying Permissions in a Policy]in the Amazon S3 User Guide.
//
// Object Data Formats You can use Amazon S3 Select to query objects that have the
// following format properties:
//
// - CSV, JSON, and Parquet - Objects must be in CSV, JSON, or Parquet format.
//
// - UTF-8 - UTF-8 is the only encoding type Amazon S3 Select supports.
//
// - GZIP or BZIP2 - CSV and JSON files can be compressed using GZIP or BZIP2.
// GZIP and BZIP2 are the only compression formats that Amazon S3 Select supports
// for CSV and JSON files. Amazon S3 Select supports columnar compression for
// Parquet using GZIP or Snappy. Amazon S3 Select does not support whole-object
// compression for Parquet objects.
//
// - Server-side encryption - Amazon S3 Select supports querying objects that
// are protected with server-side encryption.
//
// For objects that are encrypted with customer-provided encryption keys (SSE-C),
//
// you must use HTTPS, and you must use the headers that are documented in the [GetObject].
// For more information about SSE-C, see [Server-Side Encryption (Using Customer-Provided Encryption Keys)]in the Amazon S3 User Guide.
//
// # For objects that are encrypted with Amazon S3 managed keys (SSE-S3) and Amazon
//
// Web Services KMS keys (SSE-KMS), server-side encryption is handled
// transparently, so you don't need to specify anything. For more information about
// server-side encryption, including SSE-S3 and SSE-KMS, see [Protecting Data Using Server-Side Encryption]in the Amazon S3
// User Guide.
//
// Working with the Response Body Given the response size is unknown, Amazon S3
// Select streams the response as a series of messages and includes a
// Transfer-Encoding header with chunked as its value in the response. For more
// information, see [Appendix: SelectObjectContent Response].
//
// GetObject Support The SelectObjectContent action does not support the following
// GetObject functionality. For more information, see [GetObject].
//
// - Range : Although you can specify a scan range for an Amazon S3 Select
// request (see [SelectObjectContentRequest - ScanRange]in the request parameters), you cannot specify the range of
// bytes of an object to return.
//
// - The GLACIER , DEEP_ARCHIVE , and REDUCED_REDUNDANCY storage classes, or the
// ARCHIVE_ACCESS and DEEP_ARCHIVE_ACCESS access tiers of the INTELLIGENT_TIERING
// storage class: You cannot query objects in the GLACIER , DEEP_ARCHIVE , or
// REDUCED_REDUNDANCY storage classes, nor objects in the ARCHIVE_ACCESS or
// DEEP_ARCHIVE_ACCESS access tiers of the INTELLIGENT_TIERING storage class. For
// more information about storage classes, see [Using Amazon S3 storage classes]in the Amazon S3 User Guide.
//
// Special Errors For a list of special errors for this operation, see [List of SELECT Object Content Error Codes]
//
// The following operations are related to SelectObjectContent :
//
// [GetObject]
//
// [GetBucketLifecycleConfiguration]
//
// [PutBucketLifecycleConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Appendix: SelectObjectContent Response]: https://docs.aws.amazon.com/AmazonS3/latest/API/RESTSelectObjectAppendix.html
// [Selecting Content from Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/selecting-content-from-objects.html
// [PutBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketLifecycleConfiguration.html
// [SelectObjectContentRequest - ScanRange]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_SelectObjectContent.html#AmazonS3-SelectObjectContent-request-ScanRange
// [List of SELECT Object Content Error Codes]: https://docs.aws.amazon.com/AmazonS3/latest/API/ErrorResponses.html#SelectObjectContentErrorCodeList
// [GetBucketLifecycleConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketLifecycleConfiguration.html
// [Using Amazon S3 storage classes]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/storage-class-intro.html
// [SELECT Command]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-glacier-select-sql-reference-select.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
// [Specifying Permissions in a Policy]: https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html
// [Server-Side Encryption (Using Customer-Provided Encryption Keys)]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ServerSideEncryptionCustomerKeys.html
// [Protecting Data Using Server-Side Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/serv-side-encryption.html
func s3_SelectObjectContent(cfg aws.Config, client *s3.Client) {
	input := &s3.SelectObjectContentInput{
		// Bucket: *string, // Required
		// Expression: *string, // Required
		// ExpressionType: types.ExpressionType, // Required
		// InputSerialization: *types.InputSerialization, // Required
		// Key: *string, // Required
		// OutputSerialization: *types.OutputSerialization, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Expression) > 0 {
		input.Expression = aws.String(_s3Expression)
	}
	if len(_s3ExpressionType) > 0 {
		if err := assignInputField(input, "ExpressionType", _s3ExpressionType); err != nil {
			log.Errorf("invalid --expression-type: %s", err.Error())
			return
		}
	}
	if len(_s3InputSerialization) > 0 {
		if err := assignInputField(input, "InputSerialization", _s3InputSerialization); err != nil {
			log.Errorf("invalid --input-serialization: %s", err.Error())
			return
		}
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3OutputSerialization) > 0 {
		if err := assignInputField(input, "OutputSerialization", _s3OutputSerialization); err != nil {
			log.Errorf("invalid --output-serialization: %s", err.Error())
			return
		}
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestProgress) > 0 {
		if err := assignInputField(input, "RequestProgress", _s3RequestProgress); err != nil {
			log.Errorf("invalid --request-progress: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3ScanRange) > 0 {
		if err := assignInputField(input, "ScanRange", _s3ScanRange); err != nil {
			log.Errorf("invalid --scan-range: %s", err.Error())
			return
		}
	}

	if resp, err := client.SelectObjectContent(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables a live inventory table for an S3 Metadata configuration on
// a general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata]in the Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the following permissions. For
// more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// If you want to encrypt your inventory table with server-side encryption with
// Key Management Service (KMS) keys (SSE-KMS), you need additional permissions in
// your KMS key policy. For more information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// - s3:UpdateBucketMetadataInventoryTableConfiguration
//
// - s3tables:CreateTableBucket
//
// - s3tables:CreateNamespace
//
// - s3tables:GetTable
//
// - s3tables:CreateTable
//
// - s3tables:PutTablePolicy
//
// - s3tables:PutTableEncryption
//
// - kms:DescribeKey
//
// The following operations are related to
// UpdateBucketMetadataInventoryTableConfiguration :
//
// [CreateBucketMetadataConfiguration]
//
// [DeleteBucketMetadataConfiguration]
//
// [GetBucketMetadataConfiguration]
//
// [UpdateBucketMetadataJournalTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [UpdateBucketMetadataJournalTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataJournalTableConfiguration.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [DeleteBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html
func s3_UpdateBucketMetadataInventoryTableConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.UpdateBucketMetadataInventoryTableConfigurationInput{
		// Bucket: *string, // Required
		// InventoryTableConfiguration: *types.InventoryTableConfigurationUpdates, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3InventoryTableConfiguration) > 0 {
		if err := assignInputField(input, "InventoryTableConfiguration", _s3InventoryTableConfiguration); err != nil {
			log.Errorf("invalid --inventory-table-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.UpdateBucketMetadataInventoryTableConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Enables or disables journal table record expiration for an S3 Metadata
// configuration on a general purpose bucket. For more information, see [Accelerating data discovery with S3 Metadata]in the
// Amazon S3 User Guide.
//
// Permissions To use this operation, you must have the
// s3:UpdateBucketMetadataJournalTableConfiguration permission. For more
// information, see [Setting up permissions for configuring metadata tables]in the Amazon S3 User Guide.
//
// The following operations are related to
// UpdateBucketMetadataJournalTableConfiguration :
//
// [CreateBucketMetadataConfiguration]
//
// [DeleteBucketMetadataConfiguration]
//
// [GetBucketMetadataConfiguration]
//
// [UpdateBucketMetadataInventoryTableConfiguration]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [GetBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetBucketMetadataConfiguration.html
// [Setting up permissions for configuring metadata tables]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-permissions.html
// [Accelerating data discovery with S3 Metadata]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/metadata-tables-overview.html
// [CreateBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateBucketMetadataConfiguration.html
// [UpdateBucketMetadataInventoryTableConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UpdateBucketMetadataInventoryTableConfiguration.html
// [DeleteBucketMetadataConfiguration]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketMetadataConfiguration.html
func s3_UpdateBucketMetadataJournalTableConfiguration(cfg aws.Config, client *s3.Client) {
	input := &s3.UpdateBucketMetadataJournalTableConfigurationInput{
		// Bucket: *string, // Required
		// JournalTableConfiguration: *types.JournalTableConfigurationUpdates, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3JournalTableConfiguration) > 0 {
		if err := assignInputField(input, "JournalTableConfiguration", _s3JournalTableConfiguration); err != nil {
			log.Errorf("invalid --journal-table-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}

	if resp, err := client.UpdateBucketMetadataJournalTableConfiguration(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets or Amazon S3 on Outposts
// buckets.
//
// Updates the server-side encryption type of an existing encrypted object in a
// general purpose bucket. You can use the UpdateObjectEncryption operation to
// change encrypted objects from server-side encryption with Amazon S3 managed keys
// (SSE-S3) to server-side encryption with Key Management Service (KMS) keys
// (SSE-KMS), or to apply S3 Bucket Keys. You can also use the
// UpdateObjectEncryption operation to change the customer-managed KMS key used to
// encrypt your data so that you can comply with custom key-rotation standards.
//
// Using the UpdateObjectEncryption operation, you can atomically update the
// server-side encryption type of an existing object in a general purpose bucket
// without any data movement. The UpdateObjectEncryption operation uses envelope
// encryption to re-encrypt the data key used to encrypt and decrypt your object
// with your newly specified server-side encryption type. In other words, when you
// use the UpdateObjectEncryption operation, your data isn't copied, archived
// objects in the S3 Glacier Flexible Retrieval and S3 Glacier Deep Archive storage
// classes aren't restored, and objects in the S3 Intelligent-Tiering storage class
// aren't moved between tiers. Additionally, the UpdateObjectEncryption operation
// preserves all object metadata properties, including the storage class, creation
// date, last modified date, ETag, and checksum properties. For more information,
// see [Updating server-side encryption for existing objects]in the Amazon S3 User Guide.
//
// By default, all UpdateObjectEncryption requests that specify a customer-managed
// KMS key are restricted to KMS keys that are owned by the bucket owner's Amazon
// Web Services account. If you're using Organizations, you can request the ability
// to use KMS keys owned by other member accounts within your organization by
// contacting Amazon Web Services Support.
//
// Source objects that are unencrypted, or encrypted with either dual-layer
// server-side encryption with KMS keys (DSSE-KMS) or server-side encryption with
// customer-provided keys (SSE-C) aren't supported by this operation. Additionally,
// you cannot specify SSE-S3 encryption as the requested new encryption type
// UpdateObjectEncryption request.
//
// # Permissions
//
// - To use the UpdateObjectEncryption operation, you must have the following
// permissions:
//
// - s3:PutObject
//
// - s3:UpdateObjectEncryption
//
// - kms:Encrypt
//
// - kms:Decrypt
//
// - kms:GenerateDataKey
//
// - kms:ReEncrypt*
//
// - If you're using Organizations, to use this operation with customer-managed
// KMS keys from other Amazon Web Services accounts within your organization, you
// must have the organizations:DescribeAccount permission.
//
// # Errors
//
// - You might receive an InvalidRequest error for several reasons. Depending on
// the reason for the error, you might receive one of the following messages:
//
// - The UpdateObjectEncryption operation doesn't supported unencrypted source
// objects. Only source objects encrypted with SSE-S3 or SSE-KMS are supported.
//
// - The UpdateObjectEncryption operation doesn't support source objects with the
// encryption type DSSE-KMS or SSE-C. Only source objects encrypted with SSE-S3 or
// SSE-KMS are supported.
//
// - The UpdateObjectEncryption operation doesn't support updating the encryption
// type to DSSE-KMS or SSE-C. Modify the request to specify SSE-KMS for the updated
// encryption type, and then try again.
//
// - Requests that modify an object encryption configuration require Amazon Web
// Services Signature Version 4. Modify the request to use Amazon Web Services
// Signature Version 4, and then try again.
//
// - Requests that modify an object encryption configuration require a valid new
// encryption type. Valid values are SSEKMS . Modify the request to specify
// SSE-KMS for the updated encryption type, and then try again.
//
// - Requests that modify an object's encryption type to SSE-KMS require an
// Amazon Web Services KMS key Amazon Resource Name (ARN). Modify the request to
// specify a KMS key ARN, and then try again.
//
// - Requests that modify an object's encryption type to SSE-KMS require a valid
// Amazon Web Services KMS key Amazon Resource Name (ARN). Confirm that you have a
// correctly formatted KMS key ARN in your request, and then try again.
//
// - The BucketKeyEnabled value isn't valid. Valid values are true or false .
// Modify the request to specify a valid value, and then try again.
//
// - You might receive an AccessDenied error for several reasons. Depending on
// the reason for the error, you might receive one of the following messages:
//
// - The Amazon Web Services KMS key in the request must be owned by the same
// account as the bucket. Modify the request to specify a KMS key from the same
// account, and then try again.
//
// - The bucket owner's account was approved to make UpdateObjectEncryption
// requests that use any Amazon Web Services KMS key in their organization, but the
// bucket owner's account isn't part of an organization in Organizations. Make sure
// that the bucket owner's account and the specified KMS key belong to the same
// organization, and then try again.
//
// - The specified Amazon Web Services KMS key must be from the same
// organization in Organizations as the bucket. Specify a KMS key that belongs to
// the same organization as the bucket, and then try again.
//
// - The encryption type for the specified object can’t be updated because that
// object is protected by S3 Object Lock. If the object has a governance-mode
// retention period or a legal hold, you must first remove the Object Lock status
// on the object before you issue your UpdateObjectEncryption request. You can't
// use the UpdateObjectEncryption operation with objects that have an Object Lock
// compliance mode retention period applied to them.
//
// [Updating server-side encryption for existing objects]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/update-sse-encryption.html
func s3_UpdateObjectEncryption(cfg aws.Config, client *s3.Client) {
	input := &s3.UpdateObjectEncryptionInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// ObjectEncryption: types.ObjectEncryption, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3ObjectEncryption) > 0 {
		if err := assignInputField(input, "ObjectEncryption", _s3ObjectEncryption); err != nil {
			log.Errorf("invalid --object-encryption: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.UpdateObjectEncryption(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a part in a multipart upload.
// In this operation, you provide new data as a part of an object in your request.
// However, you have an option to specify your existing Amazon S3 object as a data
// source for the part you are uploading. To upload a part from an existing object,
// you use the [UploadPartCopy]operation.
//
// You must initiate a multipart upload (see [CreateMultipartUpload]) before you can upload any part. In
// response to your initiate request, Amazon S3 returns an upload ID, a unique
// identifier that you must include in your upload part request.
//
// Part numbers can be any number from 1 to 10,000, inclusive. A part number
// uniquely identifies a part and also defines its position within the object being
// created. If you upload a new part using the same part number that was used with
// a previous part, the previously uploaded part is overwritten.
//
// For information about maximum and minimum part sizes and other multipart upload
// specifications, see [Multipart upload limits]in the Amazon S3 User Guide.
//
// After you initiate multipart upload and upload one or more parts, you must
// either complete or abort multipart upload in order to stop getting charged for
// storage of the uploaded parts. Only after you either complete or abort multipart
// upload, Amazon S3 frees up the parts storage and stops charging you for the
// parts storage.
//
// For more information on multipart uploads, go to [Multipart Upload Overview] in the Amazon S3 User Guide .
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Permissions
// - General purpose bucket permissions - To perform a multipart upload with
// encryption using an Key Management Service key, the requester must have
// permission to the kms:Decrypt and kms:GenerateDataKey actions on the key. The
// requester must also have permissions for the kms:GenerateDataKey action for
// the CreateMultipartUpload API. Then, the requester needs permissions for the
// kms:Decrypt action on the UploadPart and UploadPartCopy APIs.
//
// # These permissions are required because Amazon S3 must decrypt and read data
//
// from the encrypted file parts before it completes the multipart upload. For more
// information about KMS permissions, see [Protecting data using server-side encryption with KMS]in the Amazon S3 User Guide. For
// information about the permissions required to use the multipart upload API, see [Multipart upload and permissions]
// and [Multipart upload API and permissions]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - To grant access to this API operation on a
// directory bucket, we recommend that you use the [CreateSession]CreateSession API operation
// for session-based authorization. Specifically, you grant the
// s3express:CreateSession permission to the directory bucket in a bucket policy
// or an IAM identity-based policy. Then, you make the CreateSession API call on
// the bucket to obtain a session token. With the session token in your request
// header, you can make API requests to this operation. After the session token
// expires, you make another CreateSession API call to generate a new session
// token for use. Amazon Web Services CLI or SDKs create session and refresh the
// session token automatically to avoid service interruptions when a session
// expires. For more information about authorization, see [CreateSession]CreateSession .
//
// # If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// Data integrity  General purpose bucket - To ensure that data is not corrupted
// traversing the network, specify the Content-MD5 header in the upload part
// request. Amazon S3 checks the part data against the provided MD5 value. If they
// do not match, Amazon S3 returns an error. If the upload request is signed with
// Signature Version 4, then Amazon Web Services S3 uses the x-amz-content-sha256
// header as a checksum instead of Content-MD5 . For more information see [Authenticating Requests: Using the Authorization Header (Amazon Web Services Signature Version 4)].
//
// Directory buckets - MD5 is not supported by directory buckets. You can use
// checksum algorithms to check object integrity.
//
// Encryption
// - General purpose bucket - Server-side encryption is for data encryption at
// rest. Amazon S3 encrypts your data as it writes it to disks in its data centers
// and decrypts it when you access it. You have mutually exclusive options to
// protect data using server-side encryption in Amazon S3, depending on how you
// choose to manage the encryption keys. Specifically, the encryption key options
// are Amazon S3 managed keys (SSE-S3), Amazon Web Services KMS keys (SSE-KMS), and
// Customer-Provided Keys (SSE-C). Amazon S3 encrypts data with server-side
// encryption using Amazon S3 managed keys (SSE-S3) by default. You can optionally
// tell Amazon S3 to encrypt data at rest using server-side encryption with other
// key options. The option you use depends on whether you want to use KMS keys
// (SSE-KMS) or provide your own encryption key (SSE-C).
//
// Server-side encryption is supported by the S3 Multipart Upload operations.
//
// Unless you are using a customer-provided encryption key (SSE-C), you don't need
// to specify the encryption parameters in each UploadPart request. Instead, you
// only need to specify the server-side encryption parameters in the initial
// Initiate Multipart request. For more information, see [CreateMultipartUpload].
//
// # If you have server-side encryption with customer-provided keys (SSE-C) blocked
//
// for your general purpose bucket, you will get an HTTP 403 Access Denied error
// when you specify the SSE-C request headers while writing new data to your
// bucket. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket].
//
// # If you request server-side encryption using a customer-provided encryption key
//
// (SSE-C) in your initiate multipart upload request, you must provide identical
// encryption information in each part upload using the following request headers.
//
// - x-amz-server-side-encryption-customer-algorithm
//
// - x-amz-server-side-encryption-customer-key
//
// - x-amz-server-side-encryption-customer-key-MD5
//
// For more information, see [Using Server-Side Encryption]in the Amazon S3 User Guide.
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: server-side encryption with Amazon S3
// managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ).
//
// # Special errors
//
// - Error Code: NoSuchUpload
//
// - Description: The specified multipart upload does not exist. The upload ID
// might be invalid, or the multipart upload might have been aborted or completed.
//
// - HTTP Status Code: 404 Not Found
//
// - SOAP Fault Code Prefix: Client
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to UploadPart :
//
// [CreateMultipartUpload]
//
// [CompleteMultipartUpload]
//
// [AbortMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [Authenticating Requests: Using the Authorization Header (Amazon Web Services Signature Version 4)]: https://docs.aws.amazon.com/AmazonS3/latest/API/sigv4-auth-using-authorization-header.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [Using Server-Side Encryption]: https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingServerSideEncryption.html
// [Multipart upload limits]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/qfacts.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [Multipart Upload Overview]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuoverview.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Protecting data using server-side encryption with KMS]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [Multipart upload and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [CreateSession]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateSession.html
// [Multipart upload API and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpuAndPermissions
// [Blocking or unblocking SSE-C for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.html
func s3_UploadPart(cfg aws.Config, client *s3.Client) {
	input := &s3.UploadPartInput{
		// Bucket: *string, // Required
		// Key: *string, // Required
		// PartNumber: *int32, // Required
		// UploadId: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3PartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _s3PartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_s3UploadId) > 0 {
		input.UploadId = aws.String(_s3UploadId)
	}
	if len(_s3Body) > 0 {
		if err := assignInputField(input, "Body", _s3Body); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumAlgorithm) > 0 {
		if err := assignInputField(input, "ChecksumAlgorithm", _s3ChecksumAlgorithm); err != nil {
			log.Errorf("invalid --checksum-algorithm: %s", err.Error())
			return
		}
	}
	if len(_s3ChecksumCRC32) > 0 {
		input.ChecksumCRC32 = aws.String(_s3ChecksumCRC32)
	}
	if len(_s3ChecksumCRC32C) > 0 {
		input.ChecksumCRC32C = aws.String(_s3ChecksumCRC32C)
	}
	if len(_s3ChecksumCRC64NVME) > 0 {
		input.ChecksumCRC64NVME = aws.String(_s3ChecksumCRC64NVME)
	}
	if len(_s3ChecksumSHA1) > 0 {
		input.ChecksumSHA1 = aws.String(_s3ChecksumSHA1)
	}
	if len(_s3ChecksumSHA256) > 0 {
		input.ChecksumSHA256 = aws.String(_s3ChecksumSHA256)
	}
	if len(_s3ContentLength) > 0 {
		if err := assignInputField(input, "ContentLength", _s3ContentLength); err != nil {
			log.Errorf("invalid --content-length: %s", err.Error())
			return
		}
	}
	if len(_s3ContentMD5) > 0 {
		input.ContentMD5 = aws.String(_s3ContentMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}

	if resp, err := client.UploadPart(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Uploads a part by copying data from an existing object as data source. To
// specify the data source, you add the request header x-amz-copy-source in your
// request. To specify a byte range, you add the request header
// x-amz-copy-source-range in your request.
//
// For information about maximum and minimum part sizes and other multipart upload
// specifications, see [Multipart upload limits]in the Amazon S3 User Guide.
//
// Instead of copying data from an existing object as part data, you might use the [UploadPart]
// action to upload new data as a part of an object in your request.
//
// You must initiate a multipart upload before you can upload any part. In
// response to your initiate request, Amazon S3 returns the upload ID, a unique
// identifier that you must include in your upload part request.
//
// For conceptual information about multipart uploads, see [Uploading Objects Using Multipart Upload] in the Amazon S3 User
// Guide. For information about copying objects using a single atomic action vs. a
// multipart upload, see [Operations on Objects]in the Amazon S3 User Guide.
//
// Directory buckets - For directory buckets, you must make requests for this API
// operation to the Zonal endpoint. These endpoints support virtual-hosted-style
// requests in the format
// https://amzn-s3-demo-bucket.s3express-zone-id.region-code.amazonaws.com/key-name
// . Path-style requests are not supported. For more information about endpoints
// in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones]in the Amazon S3 User Guide. For more information
// about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones]in the Amazon S3 User Guide.
//
// Authentication and authorization All UploadPartCopy requests must be
// authenticated and signed by using IAM credentials (access key ID and secret
// access key for the IAM identities). All headers with the x-amz- prefix,
// including x-amz-copy-source , must be signed. For more information, see [REST Authentication].
//
// Directory buckets - You must use IAM credentials to authenticate and authorize
// your access to the UploadPartCopy API operation, instead of using the temporary
// security credentials through the CreateSession API operation.
//
// Amazon Web Services CLI or SDKs handles authentication and authorization on
// your behalf.
//
// Permissions You must have READ access to the source object and WRITE access to
// the destination bucket.
//
// - General purpose bucket permissions - You must have the permissions in a
// policy based on the bucket types of your source bucket and destination bucket in
// an UploadPartCopy operation.
//
// - If the source object is in a general purpose bucket, you must have the
// s3:GetObject permission to read the source object that is being copied.
//
// - If the destination bucket is a general purpose bucket, you must have the
// s3:PutObject permission to write the object copy to the destination bucket.
//
// - To perform a multipart upload with encryption using an Key Management
// Service key, the requester must have permission to the kms:Decrypt and
// kms:GenerateDataKey actions on the key. The requester must also have
// permissions for the kms:GenerateDataKey action for the CreateMultipartUpload
// API. Then, the requester needs permissions for the kms:Decrypt action on the
// UploadPart and UploadPartCopy APIs. These permissions are required because
// Amazon S3 must decrypt and read data from the encrypted file parts before it
// completes the multipart upload. For more information about KMS permissions, see [Protecting data using server-side encryption with KMS]
// in the Amazon S3 User Guide. For information about the permissions required to
// use the multipart upload API, see [Multipart upload and permissions]and [Multipart upload API and permissions]in the Amazon S3 User Guide.
//
// - Directory bucket permissions - You must have permissions in a bucket policy
// or an IAM identity-based policy based on the source and destination bucket types
// in an UploadPartCopy operation.
//
// - If the source object that you want to copy is in a directory bucket, you
// must have the s3express:CreateSession permission in the Action element of a
// policy to read the object. By default, the session is in the ReadWrite mode.
// If you want to restrict the access, you can explicitly set the
// s3express:SessionMode condition key to ReadOnly on the copy source bucket.
//
// - If the copy destination is a directory bucket, you must have the
// s3express:CreateSession permission in the Action element of a policy to write
// the object to the destination. The s3express:SessionMode condition key cannot
// be set to ReadOnly on the copy destination.
//
// # If the object is encrypted with SSE-KMS, you must also have the
//
// kms:GenerateDataKey and kms:Decrypt permissions in IAM identity-based policies
// and KMS key policies for the KMS key.
//
// For example policies, see [Example bucket policies for S3 Express One Zone]and [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]in the Amazon S3 User Guide.
//
// Encryption
// - General purpose buckets - For information about using server-side
// encryption with customer-provided encryption keys with the UploadPartCopy
// operation, see [CopyObject]and [UploadPart].
//
// # If you have server-side encryption with customer-provided keys (SSE-C) blocked
//
// for your general purpose bucket, you will get an HTTP 403 Access Denied error
// when you specify the SSE-C request headers while writing new data to your
// bucket. For more information, see [Blocking or unblocking SSE-C for a general purpose bucket].
//
// - Directory buckets - For directory buckets, there are only two supported
// options for server-side encryption: server-side encryption with Amazon S3
// managed keys (SSE-S3) ( AES256 ) and server-side encryption with KMS keys
// (SSE-KMS) ( aws:kms ). For more information, see [Protecting data with server-side encryption]in the Amazon S3 User Guide.
//
// # For directory buckets, when you perform a CreateMultipartUpload operation and an
//
// UploadPartCopy operation, the request headers you provide in the
// CreateMultipartUpload request must match the default encryption configuration
// of the destination bucket.
//
// S3 Bucket Keys aren't supported, when you copy SSE-KMS encrypted objects from
//
// general purpose buckets to directory buckets, from directory buckets to general
// purpose buckets, or between directory buckets, through [UploadPartCopy]. In this case, Amazon
// S3 makes a call to KMS every time a copy request is made for a KMS-encrypted
// object.
//
// # Special errors
//
// - Error Code: NoSuchUpload
//
// - Description: The specified multipart upload does not exist. The upload ID
// might be invalid, or the multipart upload might have been aborted or completed.
//
// - HTTP Status Code: 404 Not Found
//
// - Error Code: InvalidRequest
//
// - Description: The specified copy source is not supported as a byte-range
// copy source.
//
// - HTTP Status Code: 400 Bad Request
//
// HTTP Host header syntax  Directory buckets - The HTTP Host header syntax is
// Bucket-name.s3express-zone-id.region-code.amazonaws.com .
//
// The following operations are related to UploadPartCopy :
//
// [CreateMultipartUpload]
//
// [UploadPart]
//
// [CompleteMultipartUpload]
//
// [AbortMultipartUpload]
//
// [ListParts]
//
// [ListMultipartUploads]
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Uploading Objects Using Multipart Upload]: https://docs.aws.amazon.com/AmazonS3/latest/dev/uploadobjusingmpu.html
// [Concepts for directory buckets in Local Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-lzs-for-directory-buckets.html
// [ListParts]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListParts.html
// [UploadPart]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPart.html
// [Protecting data using server-side encryption with KMS]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/UsingKMSEncryption.html
// [CopyObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html
// [Multipart upload and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/dev/mpuAndPermissions.html
// [Multipart upload API and permissions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/mpuoverview.html#mpuAndPermissions
// [CompleteMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CompleteMultipartUpload.html
// [CreateMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_CreateMultipartUpload.html
// [Multipart upload limits]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/qfacts.html
// [Amazon Web Services Identity and Access Management (IAM) identity-based policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-identity-policies.html
// [AbortMultipartUpload]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_AbortMultipartUpload.html
// [REST Authentication]: https://docs.aws.amazon.com/AmazonS3/latest/dev/RESTAuthentication.html
// [Example bucket policies for S3 Express One Zone]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-security-iam-example-bucket-policies.html
// [Operations on Objects]: https://docs.aws.amazon.com/AmazonS3/latest/dev/ObjectOperations.html
// [ListMultipartUploads]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListMultipartUploads.html
// [Regional and Zonal endpoints for directory buckets in Availability Zones]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/endpoint-directory-buckets-AZ.html
// [Blocking or unblocking SSE-C for a general purpose bucket]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/blocking-unblocking-s3-c-encryption-gpb.html
// [UploadPartCopy]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html
// [Protecting data with server-side encryption]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-serv-side-encryption.html
func s3_UploadPartCopy(cfg aws.Config, client *s3.Client) {
	input := &s3.UploadPartCopyInput{
		// Bucket: *string, // Required
		// CopySource: *string, // Required
		// Key: *string, // Required
		// PartNumber: *int32, // Required
		// UploadId: *string, // Required
	}

	if len(_s3Bucket) > 0 {
		input.Bucket = aws.String(_s3Bucket)
	}
	if len(_s3CopySource) > 0 {
		input.CopySource = aws.String(_s3CopySource)
	}
	if len(_s3Key) > 0 {
		input.Key = aws.String(_s3Key)
	}
	if len(_s3PartNumber) > 0 {
		if err := assignInputField(input, "PartNumber", _s3PartNumber); err != nil {
			log.Errorf("invalid --part-number: %s", err.Error())
			return
		}
	}
	if len(_s3UploadId) > 0 {
		input.UploadId = aws.String(_s3UploadId)
	}
	if len(_s3CopySourceIfMatch) > 0 {
		input.CopySourceIfMatch = aws.String(_s3CopySourceIfMatch)
	}
	if len(_s3CopySourceIfModifiedSince) > 0 {
		if err := assignInputField(input, "CopySourceIfModifiedSince", _s3CopySourceIfModifiedSince); err != nil {
			log.Errorf("invalid --copy-source-if-modified-since: %s", err.Error())
			return
		}
	}
	if len(_s3CopySourceIfNoneMatch) > 0 {
		input.CopySourceIfNoneMatch = aws.String(_s3CopySourceIfNoneMatch)
	}
	if len(_s3CopySourceIfUnmodifiedSince) > 0 {
		if err := assignInputField(input, "CopySourceIfUnmodifiedSince", _s3CopySourceIfUnmodifiedSince); err != nil {
			log.Errorf("invalid --copy-source-if-unmodified-since: %s", err.Error())
			return
		}
	}
	if len(_s3CopySourceRange) > 0 {
		input.CopySourceRange = aws.String(_s3CopySourceRange)
	}
	if len(_s3CopySourceSSECustomerAlgorithm) > 0 {
		input.CopySourceSSECustomerAlgorithm = aws.String(_s3CopySourceSSECustomerAlgorithm)
	}
	if len(_s3CopySourceSSECustomerKey) > 0 {
		input.CopySourceSSECustomerKey = aws.String(_s3CopySourceSSECustomerKey)
	}
	if len(_s3CopySourceSSECustomerKeyMD5) > 0 {
		input.CopySourceSSECustomerKeyMD5 = aws.String(_s3CopySourceSSECustomerKeyMD5)
	}
	if len(_s3ExpectedBucketOwner) > 0 {
		input.ExpectedBucketOwner = aws.String(_s3ExpectedBucketOwner)
	}
	if len(_s3ExpectedSourceBucketOwner) > 0 {
		input.ExpectedSourceBucketOwner = aws.String(_s3ExpectedSourceBucketOwner)
	}
	if len(_s3RequestPayer) > 0 {
		if err := assignInputField(input, "RequestPayer", _s3RequestPayer); err != nil {
			log.Errorf("invalid --request-payer: %s", err.Error())
			return
		}
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKey) > 0 {
		input.SSECustomerKey = aws.String(_s3SSECustomerKey)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}

	if resp, err := client.UploadPartCopy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// This operation is not supported for directory buckets.
// Passes transformed objects to a GetObject operation when using Object Lambda
// access points. For information about Object Lambda access points, see [Transforming objects with Object Lambda access points]in the
// Amazon S3 User Guide.
//
// This operation supports metadata that can be returned by [GetObject], in addition to
// RequestRoute , RequestToken , StatusCode , ErrorCode , and ErrorMessage . The
// GetObject response metadata is supported so that the WriteGetObjectResponse
// caller, typically an Lambda function, can provide the same metadata when it
// internally invokes GetObject . When WriteGetObjectResponse is called by a
// customer-owned Lambda function, the metadata returned to the end user GetObject
// call might differ from what Amazon S3 would normally return.
//
// You can include any number of metadata headers. When including a metadata
// header, it should be prefaced with x-amz-meta . For example,
// x-amz-meta-my-custom-header: MyCustomValue . The primary use case for this is to
// forward GetObject metadata.
//
// Amazon Web Services provides some prebuilt Lambda functions that you can use
// with S3 Object Lambda to detect and redact personally identifiable information
// (PII) and decompress S3 objects. These Lambda functions are available in the
// Amazon Web Services Serverless Application Repository, and can be selected
// through the Amazon Web Services Management Console when you create your Object
// Lambda access point.
//
// Example 1: PII Access Control - This Lambda function uses Amazon Comprehend, a
// natural language processing (NLP) service using machine learning to find
// insights and relationships in text. It automatically detects personally
// identifiable information (PII) such as names, addresses, dates, credit card
// numbers, and social security numbers from documents in your Amazon S3 bucket.
//
// Example 2: PII Redaction - This Lambda function uses Amazon Comprehend, a
// natural language processing (NLP) service using machine learning to find
// insights and relationships in text. It automatically redacts personally
// identifiable information (PII) such as names, addresses, dates, credit card
// numbers, and social security numbers from documents in your Amazon S3 bucket.
//
// Example 3: Decompression - The Lambda function S3ObjectLambdaDecompression, is
// equipped to decompress objects stored in S3 in one of six compressed file
// formats including bzip2, gzip, snappy, zlib, zstandard and ZIP.
//
// For information on how to view and use these functions, see [Using Amazon Web Services built Lambda functions] in the Amazon S3
// User Guide.
//
// You must URL encode any signed header values that contain spaces. For example,
// if your header value is my file.txt , containing two spaces after my , you must
// URL encode this value to my%20%20file.txt .
//
// [Transforming objects with Object Lambda access points]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/transforming-objects.html
// [Using Amazon Web Services built Lambda functions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/olap-examples.html
// [GetObject]: https://docs.aws.amazon.com/AmazonS3/latest/API/API_GetObject.html
func s3_WriteGetObjectResponse(cfg aws.Config, client *s3.Client) {
	input := &s3.WriteGetObjectResponseInput{
		// RequestRoute: *string, // Required
		// RequestToken: *string, // Required
	}

	if len(_s3RequestRoute) > 0 {
		input.RequestRoute = aws.String(_s3RequestRoute)
	}
	if len(_s3RequestToken) > 0 {
		input.RequestToken = aws.String(_s3RequestToken)
	}
	if len(_s3AcceptRanges) > 0 {
		input.AcceptRanges = aws.String(_s3AcceptRanges)
	}
	if len(_s3Body) > 0 {
		if err := assignInputField(input, "Body", _s3Body); err != nil {
			log.Errorf("invalid --body: %s", err.Error())
			return
		}
	}
	if len(_s3BucketKeyEnabled) > 0 {
		if err := assignInputField(input, "BucketKeyEnabled", _s3BucketKeyEnabled); err != nil {
			log.Errorf("invalid --bucket-key-enabled: %s", err.Error())
			return
		}
	}
	if len(_s3CacheControl) > 0 {
		input.CacheControl = aws.String(_s3CacheControl)
	}
	if len(_s3ChecksumCRC32) > 0 {
		input.ChecksumCRC32 = aws.String(_s3ChecksumCRC32)
	}
	if len(_s3ChecksumCRC32C) > 0 {
		input.ChecksumCRC32C = aws.String(_s3ChecksumCRC32C)
	}
	if len(_s3ChecksumCRC64NVME) > 0 {
		input.ChecksumCRC64NVME = aws.String(_s3ChecksumCRC64NVME)
	}
	if len(_s3ChecksumSHA1) > 0 {
		input.ChecksumSHA1 = aws.String(_s3ChecksumSHA1)
	}
	if len(_s3ChecksumSHA256) > 0 {
		input.ChecksumSHA256 = aws.String(_s3ChecksumSHA256)
	}
	if len(_s3ContentDisposition) > 0 {
		input.ContentDisposition = aws.String(_s3ContentDisposition)
	}
	if len(_s3ContentEncoding) > 0 {
		input.ContentEncoding = aws.String(_s3ContentEncoding)
	}
	if len(_s3ContentLanguage) > 0 {
		input.ContentLanguage = aws.String(_s3ContentLanguage)
	}
	if len(_s3ContentLength) > 0 {
		if err := assignInputField(input, "ContentLength", _s3ContentLength); err != nil {
			log.Errorf("invalid --content-length: %s", err.Error())
			return
		}
	}
	if len(_s3ContentRange) > 0 {
		input.ContentRange = aws.String(_s3ContentRange)
	}
	if len(_s3ContentType) > 0 {
		input.ContentType = aws.String(_s3ContentType)
	}
	if len(_s3DeleteMarker) > 0 {
		if err := assignInputField(input, "DeleteMarker", _s3DeleteMarker); err != nil {
			log.Errorf("invalid --delete-marker: %s", err.Error())
			return
		}
	}
	if len(_s3ETag) > 0 {
		input.ETag = aws.String(_s3ETag)
	}
	if len(_s3ErrorCode) > 0 {
		input.ErrorCode = aws.String(_s3ErrorCode)
	}
	if len(_s3ErrorMessage) > 0 {
		input.ErrorMessage = aws.String(_s3ErrorMessage)
	}
	if len(_s3Expiration) > 0 {
		input.Expiration = aws.String(_s3Expiration)
	}
	if len(_s3Expires) > 0 {
		if err := assignInputField(input, "Expires", _s3Expires); err != nil {
			log.Errorf("invalid --expires: %s", err.Error())
			return
		}
	}
	if len(_s3LastModified) > 0 {
		if err := assignInputField(input, "LastModified", _s3LastModified); err != nil {
			log.Errorf("invalid --last-modified: %s", err.Error())
			return
		}
	}
	if len(_s3Metadata) > 0 {
		if err := assignInputField(input, "Metadata", _s3Metadata); err != nil {
			log.Errorf("invalid --metadata: %s", err.Error())
			return
		}
	}
	if len(_s3MissingMeta) > 0 {
		if err := assignInputField(input, "MissingMeta", _s3MissingMeta); err != nil {
			log.Errorf("invalid --missing-meta: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockLegalHoldStatus) > 0 {
		if err := assignInputField(input, "ObjectLockLegalHoldStatus", _s3ObjectLockLegalHoldStatus); err != nil {
			log.Errorf("invalid --object-lock-legal-hold-status: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockMode) > 0 {
		if err := assignInputField(input, "ObjectLockMode", _s3ObjectLockMode); err != nil {
			log.Errorf("invalid --object-lock-mode: %s", err.Error())
			return
		}
	}
	if len(_s3ObjectLockRetainUntilDate) > 0 {
		if err := assignInputField(input, "ObjectLockRetainUntilDate", _s3ObjectLockRetainUntilDate); err != nil {
			log.Errorf("invalid --object-lock-retain-until-date: %s", err.Error())
			return
		}
	}
	if len(_s3PartsCount) > 0 {
		if err := assignInputField(input, "PartsCount", _s3PartsCount); err != nil {
			log.Errorf("invalid --parts-count: %s", err.Error())
			return
		}
	}
	if len(_s3ReplicationStatus) > 0 {
		if err := assignInputField(input, "ReplicationStatus", _s3ReplicationStatus); err != nil {
			log.Errorf("invalid --replication-status: %s", err.Error())
			return
		}
	}
	if len(_s3RequestCharged) > 0 {
		if err := assignInputField(input, "RequestCharged", _s3RequestCharged); err != nil {
			log.Errorf("invalid --request-charged: %s", err.Error())
			return
		}
	}
	if len(_s3Restore) > 0 {
		input.Restore = aws.String(_s3Restore)
	}
	if len(_s3SSECustomerAlgorithm) > 0 {
		input.SSECustomerAlgorithm = aws.String(_s3SSECustomerAlgorithm)
	}
	if len(_s3SSECustomerKeyMD5) > 0 {
		input.SSECustomerKeyMD5 = aws.String(_s3SSECustomerKeyMD5)
	}
	if len(_s3SSEKMSKeyId) > 0 {
		input.SSEKMSKeyId = aws.String(_s3SSEKMSKeyId)
	}
	if len(_s3ServerSideEncryption) > 0 {
		if err := assignInputField(input, "ServerSideEncryption", _s3ServerSideEncryption); err != nil {
			log.Errorf("invalid --server-side-encryption: %s", err.Error())
			return
		}
	}
	if len(_s3StatusCode) > 0 {
		if err := assignInputField(input, "StatusCode", _s3StatusCode); err != nil {
			log.Errorf("invalid --status-code: %s", err.Error())
			return
		}
	}
	if len(_s3StorageClass) > 0 {
		if err := assignInputField(input, "StorageClass", _s3StorageClass); err != nil {
			log.Errorf("invalid --storage-class: %s", err.Error())
			return
		}
	}
	if len(_s3TagCount) > 0 {
		if err := assignInputField(input, "TagCount", _s3TagCount); err != nil {
			log.Errorf("invalid --tag-count: %s", err.Error())
			return
		}
	}
	if len(_s3VersionId) > 0 {
		input.VersionId = aws.String(_s3VersionId)
	}

	if resp, err := client.WriteGetObjectResponse(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_s3Cmd)
	_s3Cmd.Flags().SortFlags = false

	_s3Cmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_s3Cmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_s3Cmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_s3Cmd.Flags().StringVarP(&_s3AbacStatus, "abac-status", "", "", "Abac Status")
	_s3Cmd.Flags().StringVarP(&_s3AccelerateConfiguration, "accelerate-configuration", "", "", "Accelerate Configuration")
	_s3Cmd.Flags().StringVarP(&_s3AcceptRanges, "accept-ranges", "", "", "Accept Ranges")
	_s3Cmd.Flags().StringVarP(&_s3AccessControlPolicy, "access-control-policy", "", "", "Access Control Policy")
	_s3Cmd.Flags().StringVarP(&_s3ACL, "acl", "", "", "ACL")
	_s3Cmd.Flags().StringVarP(&_s3AnalyticsConfiguration, "analytics-configuration", "", "", "Analytics Configuration")
	_s3Cmd.Flags().StringVarP(&_s3Body, "body", "", "", "Body")
	_s3Cmd.Flags().StringVarP(&_s3Bucket, "bucket", "", "", "Bucket")
	_s3Cmd.Flags().StringVarP(&_s3BucketKeyEnabled, "bucket-key-enabled", "", "", "Bucket Key Enabled")
	_s3Cmd.Flags().StringVarP(&_s3BucketLoggingStatus, "bucket-logging-status", "", "", "Bucket Logging Status")
	_s3Cmd.Flags().StringVarP(&_s3BucketRegion, "bucket-region", "", "", "Bucket Region")
	_s3Cmd.Flags().StringVarP(&_s3BypassGovernanceRetention, "bypass-governance-retention", "", "", "Bypass Governance Retention")
	_s3Cmd.Flags().StringVarP(&_s3CacheControl, "cache-control", "", "", "Cache Control")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumAlgorithm, "checksum-algorithm", "", "", "Checksum Algorithm")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumCRC32, "checksum-crc32", "", "", "Checksum CRC32")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumCRC32C, "checksum-crc32c", "", "", "Checksum CRC32C")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumCRC64NVME, "checksum-crc64nvme", "", "", "Checksum CRC64NVME")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumMode, "checksum-mode", "", "", "Checksum Mode")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumSHA1, "checksum-sha1", "", "", "Checksum SHA1")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumSHA256, "checksum-sha256", "", "", "Checksum SHA256")
	_s3Cmd.Flags().StringVarP(&_s3ChecksumType, "checksum-type", "", "", "Checksum Type")
	_s3Cmd.Flags().StringVarP(&_s3ClientToken, "client-token", "", "", "Client Token")
	_s3Cmd.Flags().StringVarP(&_s3ConfirmRemoveSelfBucketAccess, "confirm-remove-self-bucket-access", "", "", "Confirm Remove Self Bucket Access")
	_s3Cmd.Flags().StringVarP(&_s3ContentDisposition, "content-disposition", "", "", "Content Disposition")
	_s3Cmd.Flags().StringVarP(&_s3ContentEncoding, "content-encoding", "", "", "Content Encoding")
	_s3Cmd.Flags().StringVarP(&_s3ContentLanguage, "content-language", "", "", "Content Language")
	_s3Cmd.Flags().StringVarP(&_s3ContentLength, "content-length", "", "", "Content Length")
	_s3Cmd.Flags().StringVarP(&_s3ContentMD5, "content-md5", "", "", "Content MD5")
	_s3Cmd.Flags().StringVarP(&_s3ContentRange, "content-range", "", "", "Content Range")
	_s3Cmd.Flags().StringVarP(&_s3ContentType, "content-type", "", "", "Content Type")
	_s3Cmd.Flags().StringVarP(&_s3ContinuationToken, "continuation-token", "", "", "Continuation Token")
	_s3Cmd.Flags().StringVarP(&_s3CopySource, "copy-source", "", "", "Copy Source")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceIfMatch, "copy-source-if-match", "", "", "Copy Source If Match")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceIfModifiedSince, "copy-source-if-modified-since", "", "", "Copy Source If Modified Since")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceIfNoneMatch, "copy-source-if-none-match", "", "", "Copy Source If None Match")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceIfUnmodifiedSince, "copy-source-if-unmodified-since", "", "", "Copy Source If Unmodified Since")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceRange, "copy-source-range", "", "", "Copy Source Range")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceSSECustomerAlgorithm, "copy-source-sse-customer-algorithm", "", "", "Copy Source SSE Customer Algorithm")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceSSECustomerKey, "copy-source-sse-customer-key", "", "", "Copy Source SSE Customer Key")
	_s3Cmd.Flags().StringVarP(&_s3CopySourceSSECustomerKeyMD5, "copy-source-sse-customer-key-md5", "", "", "Copy Source SSE Customer Key MD5")
	_s3Cmd.Flags().StringVarP(&_s3CORSConfiguration, "cors-configuration", "", "", "Cors Configuration")
	_s3Cmd.Flags().StringVarP(&_s3CreateBucketConfiguration, "create-bucket-configuration", "", "", "Create Bucket Configuration")
	_s3Cmd.Flags().StringVarP(&_s3Delete, "delete", "", "", "Delete")
	_s3Cmd.Flags().StringVarP(&_s3DeleteMarker, "delete-marker", "", "", "Delete Marker")
	_s3Cmd.Flags().StringVarP(&_s3Delimiter, "delimiter", "", "", "Delimiter")
	_s3Cmd.Flags().StringVarP(&_s3DestinationIfMatch, "destination-if-match", "", "", "Destination If Match")
	_s3Cmd.Flags().StringVarP(&_s3DestinationIfModifiedSince, "destination-if-modified-since", "", "", "Destination If Modified Since")
	_s3Cmd.Flags().StringVarP(&_s3DestinationIfNoneMatch, "destination-if-none-match", "", "", "Destination If None Match")
	_s3Cmd.Flags().StringVarP(&_s3DestinationIfUnmodifiedSince, "destination-if-unmodified-since", "", "", "Destination If Unmodified Since")
	_s3Cmd.Flags().StringVarP(&_s3EncodingType, "encoding-type", "", "", "Encoding Type")
	_s3Cmd.Flags().StringVarP(&_s3ErrorCode, "error-code", "", "", "Error Code")
	_s3Cmd.Flags().StringVarP(&_s3ErrorMessage, "error-message", "", "", "Error Message")
	_s3Cmd.Flags().StringVarP(&_s3ETag, "etag", "", "", "Etag")
	_s3Cmd.Flags().StringVarP(&_s3ExpectedBucketOwner, "expected-bucket-owner", "", "", "Expected Bucket Owner")
	_s3Cmd.Flags().StringVarP(&_s3ExpectedSourceBucketOwner, "expected-source-bucket-owner", "", "", "Expected Source Bucket Owner")
	_s3Cmd.Flags().StringVarP(&_s3Expiration, "expiration", "", "", "Expiration")
	_s3Cmd.Flags().StringVarP(&_s3Expires, "expires", "", "", "Expires")
	_s3Cmd.Flags().StringVarP(&_s3Expression, "expression", "", "", "Expression")
	_s3Cmd.Flags().StringVarP(&_s3ExpressionType, "expression-type", "", "", "Expression Type")
	_s3Cmd.Flags().StringVarP(&_s3FetchOwner, "fetch-owner", "", "", "Fetch Owner")
	_s3Cmd.Flags().StringVarP(&_s3GrantFullControl, "grant-full-control", "", "", "Grant Full Control")
	_s3Cmd.Flags().StringVarP(&_s3GrantRead, "grant-read", "", "", "Grant Read")
	_s3Cmd.Flags().StringVarP(&_s3GrantReadACP, "grant-read-acp", "", "", "Grant Read Acp")
	_s3Cmd.Flags().StringVarP(&_s3GrantWrite, "grant-write", "", "", "Grant Write")
	_s3Cmd.Flags().StringVarP(&_s3GrantWriteACP, "grant-write-acp", "", "", "Grant Write Acp")
	_s3Cmd.Flags().StringVarP(&_s3Id, "id", "", "", "ID")
	_s3Cmd.Flags().StringVarP(&_s3IfMatch, "if-match", "", "", "If Match")
	_s3Cmd.Flags().StringVarP(&_s3IfMatchInitiatedTime, "if-match-initiated-time", "", "", "If Match Initiated Time")
	_s3Cmd.Flags().StringVarP(&_s3IfMatchLastModifiedTime, "if-match-last-modified-time", "", "", "If Match Last Modified Time")
	_s3Cmd.Flags().StringVarP(&_s3IfMatchSize, "if-match-size", "", "", "If Match Size")
	_s3Cmd.Flags().StringVarP(&_s3IfModifiedSince, "if-modified-since", "", "", "If Modified Since")
	_s3Cmd.Flags().StringVarP(&_s3IfNoneMatch, "if-none-match", "", "", "If None Match")
	_s3Cmd.Flags().StringVarP(&_s3IfUnmodifiedSince, "if-unmodified-since", "", "", "If Unmodified Since")
	_s3Cmd.Flags().StringVarP(&_s3InputSerialization, "input-serialization", "", "", "Input Serialization")
	_s3Cmd.Flags().StringVarP(&_s3IntelligentTieringConfiguration, "intelligent-tiering-configuration", "", "", "Intelligent Tiering Configuration")
	_s3Cmd.Flags().StringVarP(&_s3InventoryConfiguration, "inventory-configuration", "", "", "Inventory Configuration")
	_s3Cmd.Flags().StringVarP(&_s3InventoryTableConfiguration, "inventory-table-configuration", "", "", "Inventory Table Configuration")
	_s3Cmd.Flags().StringVarP(&_s3JournalTableConfiguration, "journal-table-configuration", "", "", "Journal Table Configuration")
	_s3Cmd.Flags().StringVarP(&_s3Key, "key", "", "", "Key")
	_s3Cmd.Flags().StringVarP(&_s3KeyMarker, "key-marker", "", "", "Key Marker")
	_s3Cmd.Flags().StringVarP(&_s3LastModified, "last-modified", "", "", "Last Modified")
	_s3Cmd.Flags().StringVarP(&_s3LegalHold, "legal-hold", "", "", "Legal Hold")
	_s3Cmd.Flags().StringVarP(&_s3LifecycleConfiguration, "lifecycle-configuration", "", "", "Lifecycle Configuration")
	_s3Cmd.Flags().StringVarP(&_s3Marker, "marker", "", "", "Marker")
	_s3Cmd.Flags().StringVarP(&_s3MaxBuckets, "max-buckets", "", "", "Max Buckets")
	_s3Cmd.Flags().StringVarP(&_s3MaxDirectoryBuckets, "max-directory-buckets", "", "", "Max Directory Buckets")
	_s3Cmd.Flags().StringVarP(&_s3MaxKeys, "max-keys", "", "", "Max Keys")
	_s3Cmd.Flags().StringVarP(&_s3MaxParts, "max-parts", "", "", "Max Parts")
	_s3Cmd.Flags().StringVarP(&_s3MaxUploads, "max-uploads", "", "", "Max Uploads")
	_s3Cmd.Flags().StringVarP(&_s3Metadata, "metadata", "", "", "Metadata")
	_s3Cmd.Flags().StringVarP(&_s3MetadataConfiguration, "metadata-configuration", "", "", "Metadata Configuration")
	_s3Cmd.Flags().StringVarP(&_s3MetadataDirective, "metadata-directive", "", "", "Metadata Directive")
	_s3Cmd.Flags().StringVarP(&_s3MetadataTableConfiguration, "metadata-table-configuration", "", "", "Metadata Table Configuration")
	_s3Cmd.Flags().StringVarP(&_s3MetricsConfiguration, "metrics-configuration", "", "", "Metrics Configuration")
	_s3Cmd.Flags().StringVarP(&_s3MFA, "mfa", "", "", "MFA")
	_s3Cmd.Flags().StringVarP(&_s3MissingMeta, "missing-meta", "", "", "Missing Meta")
	_s3Cmd.Flags().StringVarP(&_s3MpuObjectSize, "mpu-object-size", "", "", "Mpu Object Size")
	_s3Cmd.Flags().StringVarP(&_s3MultipartUpload, "multipart-upload", "", "", "Multipart Upload")
	_s3Cmd.Flags().StringVarP(&_s3NotificationConfiguration, "notification-configuration", "", "", "Notification Configuration")
	_s3Cmd.Flags().StringVarP(&_s3ObjectAttributes, "object-attributes", "", "", "Object Attributes")
	_s3Cmd.Flags().StringVarP(&_s3ObjectEncryption, "object-encryption", "", "", "Object Encryption")
	_s3Cmd.Flags().StringVarP(&_s3ObjectLockConfiguration, "object-lock-configuration", "", "", "Object Lock Configuration")
	_s3Cmd.Flags().StringVarP(&_s3ObjectLockEnabledForBucket, "object-lock-enabled-for-bucket", "", "", "Object Lock Enabled For Bucket")
	_s3Cmd.Flags().StringVarP(&_s3ObjectLockLegalHoldStatus, "object-lock-legal-hold-status", "", "", "Object Lock Legal Hold Status")
	_s3Cmd.Flags().StringVarP(&_s3ObjectLockMode, "object-lock-mode", "", "", "Object Lock Mode")
	_s3Cmd.Flags().StringVarP(&_s3ObjectLockRetainUntilDate, "object-lock-retain-until-date", "", "", "Object Lock Retain Until Date")
	_s3Cmd.Flags().StringVarP(&_s3ObjectOwnership, "object-ownership", "", "", "Object Ownership")
	_s3Cmd.Flags().StringVarP(&_s3OptionalObjectAttributes, "optional-object-attributes", "", "", "Optional Object Attributes")
	_s3Cmd.Flags().StringVarP(&_s3OutputSerialization, "output-serialization", "", "", "Output Serialization")
	_s3Cmd.Flags().StringVarP(&_s3OwnershipControls, "ownership-controls", "", "", "Ownership Controls")
	_s3Cmd.Flags().StringVarP(&_s3PartNumber, "part-number", "", "", "Part Number")
	_s3Cmd.Flags().StringVarP(&_s3PartNumberMarker, "part-number-marker", "", "", "Part Number Marker")
	_s3Cmd.Flags().StringVarP(&_s3PartsCount, "parts-count", "", "", "Parts Count")
	_s3Cmd.Flags().StringVarP(&_s3Policy, "policy", "", "", "Policy")
	_s3Cmd.Flags().StringVarP(&_s3Prefix, "prefix", "", "", "Prefix")
	_s3Cmd.Flags().StringVarP(&_s3PublicAccessBlockConfiguration, "public-access-block-configuration", "", "", "Public Access Block Configuration")
	_s3Cmd.Flags().StringVarP(&_s3Range, "range", "", "", "Range")
	_s3Cmd.Flags().StringVarP(&_s3RenameSource, "rename-source", "", "", "Rename Source")
	_s3Cmd.Flags().StringVarP(&_s3ReplicationConfiguration, "replication-configuration", "", "", "Replication Configuration")
	_s3Cmd.Flags().StringVarP(&_s3ReplicationStatus, "replication-status", "", "", "Replication Status")
	_s3Cmd.Flags().StringVarP(&_s3RequestCharged, "request-charged", "", "", "Request Charged")
	_s3Cmd.Flags().StringVarP(&_s3RequestPayer, "request-payer", "", "", "Request Payer")
	_s3Cmd.Flags().StringVarP(&_s3RequestPaymentConfiguration, "request-payment-configuration", "", "", "Request Payment Configuration")
	_s3Cmd.Flags().StringVarP(&_s3RequestProgress, "request-progress", "", "", "Request Progress")
	_s3Cmd.Flags().StringVarP(&_s3RequestRoute, "request-route", "", "", "Request Route")
	_s3Cmd.Flags().StringVarP(&_s3RequestToken, "request-token", "", "", "Request Token")
	_s3Cmd.Flags().StringVarP(&_s3ResponseCacheControl, "response-cache-control", "", "", "Response Cache Control")
	_s3Cmd.Flags().StringVarP(&_s3ResponseContentDisposition, "response-content-disposition", "", "", "Response Content Disposition")
	_s3Cmd.Flags().StringVarP(&_s3ResponseContentEncoding, "response-content-encoding", "", "", "Response Content Encoding")
	_s3Cmd.Flags().StringVarP(&_s3ResponseContentLanguage, "response-content-language", "", "", "Response Content Language")
	_s3Cmd.Flags().StringVarP(&_s3ResponseContentType, "response-content-type", "", "", "Response Content Type")
	_s3Cmd.Flags().StringVarP(&_s3ResponseExpires, "response-expires", "", "", "Response Expires")
	_s3Cmd.Flags().StringVarP(&_s3Restore, "restore", "", "", "Restore")
	_s3Cmd.Flags().StringVarP(&_s3RestoreRequest, "restore-request", "", "", "Restore Request")
	_s3Cmd.Flags().StringVarP(&_s3Retention, "retention", "", "", "Retention")
	_s3Cmd.Flags().StringVarP(&_s3ScanRange, "scan-range", "", "", "Scan Range")
	_s3Cmd.Flags().StringVarP(&_s3ServerSideEncryption, "server-side-encryption", "", "", "Server Side Encryption")
	_s3Cmd.Flags().StringVarP(&_s3ServerSideEncryptionConfiguration, "server-side-encryption-configuration", "", "", "Server Side Encryption Configuration")
	_s3Cmd.Flags().StringVarP(&_s3SessionMode, "session-mode", "", "", "Session Mode")
	_s3Cmd.Flags().StringVarP(&_s3SkipDestinationValidation, "skip-destination-validation", "", "", "Skip Destination Validation")
	_s3Cmd.Flags().StringVarP(&_s3SourceIfMatch, "source-if-match", "", "", "Source If Match")
	_s3Cmd.Flags().StringVarP(&_s3SourceIfModifiedSince, "source-if-modified-since", "", "", "Source If Modified Since")
	_s3Cmd.Flags().StringVarP(&_s3SourceIfNoneMatch, "source-if-none-match", "", "", "Source If None Match")
	_s3Cmd.Flags().StringVarP(&_s3SourceIfUnmodifiedSince, "source-if-unmodified-since", "", "", "Source If Unmodified Since")
	_s3Cmd.Flags().StringVarP(&_s3SSECustomerAlgorithm, "sse-customer-algorithm", "", "", "SSE Customer Algorithm")
	_s3Cmd.Flags().StringVarP(&_s3SSECustomerKey, "sse-customer-key", "", "", "SSE Customer Key")
	_s3Cmd.Flags().StringVarP(&_s3SSECustomerKeyMD5, "sse-customer-key-md5", "", "", "SSE Customer Key MD5")
	_s3Cmd.Flags().StringVarP(&_s3SSEKMSEncryptionContext, "ssekms-encryption-context", "", "", "Ssekms Encryption Context")
	_s3Cmd.Flags().StringVarP(&_s3SSEKMSKeyId, "ssekms-key-id", "", "", "Ssekms Key ID")
	_s3Cmd.Flags().StringVarP(&_s3StartAfter, "start-after", "", "", "Start After")
	_s3Cmd.Flags().StringVarP(&_s3StatusCode, "status-code", "", "", "Status Code")
	_s3Cmd.Flags().StringVarP(&_s3StorageClass, "storage-class", "", "", "Storage Class")
	_s3Cmd.Flags().StringVarP(&_s3TagCount, "tag-count", "", "", "Tag Count")
	_s3Cmd.Flags().StringVarP(&_s3Tagging, "tagging", "", "", "Tagging")
	_s3Cmd.Flags().StringVarP(&_s3TaggingDirective, "tagging-directive", "", "", "Tagging Directive")
	_s3Cmd.Flags().StringVarP(&_s3Token, "token", "", "", "Token")
	_s3Cmd.Flags().StringVarP(&_s3TransitionDefaultMinimumObjectSize, "transition-default-minimum-object-size", "", "", "Transition Default Minimum Object Size")
	_s3Cmd.Flags().StringVarP(&_s3UploadId, "upload-id", "", "", "Upload ID")
	_s3Cmd.Flags().StringVarP(&_s3UploadIdMarker, "upload-id-marker", "", "", "Upload ID Marker")
	_s3Cmd.Flags().StringVarP(&_s3VersionId, "version-id", "", "", "Version ID")
	_s3Cmd.Flags().StringVarP(&_s3VersionIdMarker, "version-id-marker", "", "", "Version ID Marker")
	_s3Cmd.Flags().StringVarP(&_s3VersioningConfiguration, "versioning-configuration", "", "", "Versioning Configuration")
	_s3Cmd.Flags().StringVarP(&_s3WebsiteConfiguration, "website-configuration", "", "", "Website Configuration")
	_s3Cmd.Flags().StringVarP(&_s3WebsiteRedirectLocation, "website-redirect-location", "", "", "Website Redirect Location")
	_s3Cmd.Flags().StringVarP(&_s3WriteOffsetBytes, "write-offset-bytes", "", "", "Write Offset Bytes")

	_s3Cmd.Flags().BoolVarP(&_s3AbortMultipartUpload, "abort-multipart-upload", "", false, "Abort Multipart Upload")
	_s3Cmd.Flags().BoolVarP(&_s3CompleteMultipartUpload, "complete-multipart-upload", "", false, "Complete Multipart Upload")
	_s3Cmd.Flags().BoolVarP(&_s3CopyObject, "copy-object", "", false, "Copy Object")
	_s3Cmd.Flags().BoolVarP(&_s3CreateBucket, "create-bucket", "", false, "Create Bucket")
	_s3Cmd.Flags().BoolVarP(&_s3CreateBucketMetadataConfiguration, "create-bucket-metadata-configuration", "", false, "Create Bucket Metadata Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3CreateBucketMetadataTableConfiguration, "create-bucket-metadata-table-configuration", "", false, "Create Bucket Metadata Table Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3CreateMultipartUpload, "create-multipart-upload", "", false, "Create Multipart Upload")
	_s3Cmd.Flags().BoolVarP(&_s3CreateSession, "create-session", "", false, "Create Session")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucket, "delete-bucket", "", false, "Delete Bucket")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketAnalyticsConfiguration, "delete-bucket-analytics-configuration", "", false, "Delete Bucket Analytics Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketCors, "delete-bucket-cors", "", false, "Delete Bucket Cors")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketEncryption, "delete-bucket-encryption", "", false, "Delete Bucket Encryption")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketIntelligentTieringConfiguration, "delete-bucket-intelligent-tiering-configuration", "", false, "Delete Bucket Intelligent Tiering Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketInventoryConfiguration, "delete-bucket-inventory-configuration", "", false, "Delete Bucket Inventory Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketLifecycle, "delete-bucket-lifecycle", "", false, "Delete Bucket Lifecycle")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketMetadataConfiguration, "delete-bucket-metadata-configuration", "", false, "Delete Bucket Metadata Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketMetadataTableConfiguration, "delete-bucket-metadata-table-configuration", "", false, "Delete Bucket Metadata Table Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketMetricsConfiguration, "delete-bucket-metrics-configuration", "", false, "Delete Bucket Metrics Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketOwnershipControls, "delete-bucket-ownership-controls", "", false, "Delete Bucket Ownership Controls")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketPolicy, "delete-bucket-policy", "", false, "Delete Bucket Policy")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketReplication, "delete-bucket-replication", "", false, "Delete Bucket Replication")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketTagging, "delete-bucket-tagging", "", false, "Delete Bucket Tagging")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteBucketWebsite, "delete-bucket-website", "", false, "Delete Bucket Website")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteObject, "delete-object", "", false, "Delete Object")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteObjectTagging, "delete-object-tagging", "", false, "Delete Object Tagging")
	_s3Cmd.Flags().BoolVarP(&_s3DeleteObjects, "delete-objects", "", false, "Delete Objects")
	_s3Cmd.Flags().BoolVarP(&_s3DeletePublicAccessBlock, "delete-public-access-block", "", false, "Delete Public Access Block")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketAbac, "get-bucket-abac", "", false, "Get Bucket Abac")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketAccelerateConfiguration, "get-bucket-accelerate-configuration", "", false, "Get Bucket Accelerate Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketAcl, "get-bucket-acl", "", false, "Get Bucket ACL")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketAnalyticsConfiguration, "get-bucket-analytics-configuration", "", false, "Get Bucket Analytics Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketCors, "get-bucket-cors", "", false, "Get Bucket Cors")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketEncryption, "get-bucket-encryption", "", false, "Get Bucket Encryption")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketIntelligentTieringConfiguration, "get-bucket-intelligent-tiering-configuration", "", false, "Get Bucket Intelligent Tiering Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketInventoryConfiguration, "get-bucket-inventory-configuration", "", false, "Get Bucket Inventory Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketLifecycleConfiguration, "get-bucket-lifecycle-configuration", "", false, "Get Bucket Lifecycle Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketLocation, "get-bucket-location", "", false, "Get Bucket Location")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketLogging, "get-bucket-logging", "", false, "Get Bucket Logging")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketMetadataConfiguration, "get-bucket-metadata-configuration", "", false, "Get Bucket Metadata Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketMetadataTableConfiguration, "get-bucket-metadata-table-configuration", "", false, "Get Bucket Metadata Table Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketMetricsConfiguration, "get-bucket-metrics-configuration", "", false, "Get Bucket Metrics Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketNotificationConfiguration, "get-bucket-notification-configuration", "", false, "Get Bucket Notification Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketOwnershipControls, "get-bucket-ownership-controls", "", false, "Get Bucket Ownership Controls")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketPolicy, "get-bucket-policy", "", false, "Get Bucket Policy")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketPolicyStatus, "get-bucket-policy-status", "", false, "Get Bucket Policy Status")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketReplication, "get-bucket-replication", "", false, "Get Bucket Replication")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketRequestPayment, "get-bucket-request-payment", "", false, "Get Bucket Request Payment")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketTagging, "get-bucket-tagging", "", false, "Get Bucket Tagging")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketVersioning, "get-bucket-versioning", "", false, "Get Bucket Versioning")
	_s3Cmd.Flags().BoolVarP(&_s3GetBucketWebsite, "get-bucket-website", "", false, "Get Bucket Website")
	_s3Cmd.Flags().BoolVarP(&_s3GetObject, "get-object", "", false, "Get Object")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectAcl, "get-object-acl", "", false, "Get Object ACL")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectAttributes, "get-object-attributes", "", false, "Get Object Attributes")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectLegalHold, "get-object-legal-hold", "", false, "Get Object Legal Hold")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectLockConfiguration, "get-object-lock-configuration", "", false, "Get Object Lock Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectRetention, "get-object-retention", "", false, "Get Object Retention")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectTagging, "get-object-tagging", "", false, "Get Object Tagging")
	_s3Cmd.Flags().BoolVarP(&_s3GetObjectTorrent, "get-object-torrent", "", false, "Get Object Torrent")
	_s3Cmd.Flags().BoolVarP(&_s3GetPublicAccessBlock, "get-public-access-block", "", false, "Get Public Access Block")
	_s3Cmd.Flags().BoolVarP(&_s3HeadBucket, "head-bucket", "", false, "Head Bucket")
	_s3Cmd.Flags().BoolVarP(&_s3HeadObject, "head-object", "", false, "Head Object")
	_s3Cmd.Flags().BoolVarP(&_s3ListBucketAnalyticsConfigurations, "list-bucket-analytics-configurations", "", false, "List Bucket Analytics Configurations")
	_s3Cmd.Flags().BoolVarP(&_s3ListBucketIntelligentTieringConfigurations, "list-bucket-intelligent-tiering-configurations", "", false, "List Bucket Intelligent Tiering Configurations")
	_s3Cmd.Flags().BoolVarP(&_s3ListBucketInventoryConfigurations, "list-bucket-inventory-configurations", "", false, "List Bucket Inventory Configurations")
	_s3Cmd.Flags().BoolVarP(&_s3ListBucketMetricsConfigurations, "list-bucket-metrics-configurations", "", false, "List Bucket Metrics Configurations")
	_s3Cmd.Flags().BoolVarP(&_s3ListBuckets, "list-buckets", "", false, "List Buckets")
	_s3Cmd.Flags().BoolVarP(&_s3ListDirectoryBuckets, "list-directory-buckets", "", false, "List Directory Buckets")
	_s3Cmd.Flags().BoolVarP(&_s3ListMultipartUploads, "list-multipart-uploads", "", false, "List Multipart Uploads")
	_s3Cmd.Flags().BoolVarP(&_s3ListObjectVersions, "list-object-versions", "", false, "List Object Versions")
	_s3Cmd.Flags().BoolVarP(&_s3ListObjects, "list-objects", "", false, "List Objects")
	_s3Cmd.Flags().BoolVarP(&_s3ListObjectsV2, "list-objects-v2", "", false, "List Objects V2")
	_s3Cmd.Flags().BoolVarP(&_s3ListParts, "list-parts", "", false, "List Parts")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketAbac, "put-bucket-abac", "", false, "Put Bucket Abac")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketAccelerateConfiguration, "put-bucket-accelerate-configuration", "", false, "Put Bucket Accelerate Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketAcl, "put-bucket-acl", "", false, "Put Bucket ACL")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketAnalyticsConfiguration, "put-bucket-analytics-configuration", "", false, "Put Bucket Analytics Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketCors, "put-bucket-cors", "", false, "Put Bucket Cors")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketEncryption, "put-bucket-encryption", "", false, "Put Bucket Encryption")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketIntelligentTieringConfiguration, "put-bucket-intelligent-tiering-configuration", "", false, "Put Bucket Intelligent Tiering Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketInventoryConfiguration, "put-bucket-inventory-configuration", "", false, "Put Bucket Inventory Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketLifecycleConfiguration, "put-bucket-lifecycle-configuration", "", false, "Put Bucket Lifecycle Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketLogging, "put-bucket-logging", "", false, "Put Bucket Logging")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketMetricsConfiguration, "put-bucket-metrics-configuration", "", false, "Put Bucket Metrics Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketNotificationConfiguration, "put-bucket-notification-configuration", "", false, "Put Bucket Notification Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketOwnershipControls, "put-bucket-ownership-controls", "", false, "Put Bucket Ownership Controls")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketPolicy, "put-bucket-policy", "", false, "Put Bucket Policy")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketReplication, "put-bucket-replication", "", false, "Put Bucket Replication")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketRequestPayment, "put-bucket-request-payment", "", false, "Put Bucket Request Payment")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketTagging, "put-bucket-tagging", "", false, "Put Bucket Tagging")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketVersioning, "put-bucket-versioning", "", false, "Put Bucket Versioning")
	_s3Cmd.Flags().BoolVarP(&_s3PutBucketWebsite, "put-bucket-website", "", false, "Put Bucket Website")
	_s3Cmd.Flags().BoolVarP(&_s3PutObject, "put-object", "", false, "Put Object")
	_s3Cmd.Flags().BoolVarP(&_s3PutObjectAcl, "put-object-acl", "", false, "Put Object ACL")
	_s3Cmd.Flags().BoolVarP(&_s3PutObjectLegalHold, "put-object-legal-hold", "", false, "Put Object Legal Hold")
	_s3Cmd.Flags().BoolVarP(&_s3PutObjectLockConfiguration, "put-object-lock-configuration", "", false, "Put Object Lock Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3PutObjectRetention, "put-object-retention", "", false, "Put Object Retention")
	_s3Cmd.Flags().BoolVarP(&_s3PutObjectTagging, "put-object-tagging", "", false, "Put Object Tagging")
	_s3Cmd.Flags().BoolVarP(&_s3PutPublicAccessBlock, "put-public-access-block", "", false, "Put Public Access Block")
	_s3Cmd.Flags().BoolVarP(&_s3RenameObject, "rename-object", "", false, "Rename Object")
	_s3Cmd.Flags().BoolVarP(&_s3RestoreObject, "restore-object", "", false, "Restore Object")
	_s3Cmd.Flags().BoolVarP(&_s3SelectObjectContent, "select-object-content", "", false, "Select Object Content")
	_s3Cmd.Flags().BoolVarP(&_s3UpdateBucketMetadataInventoryTableConfiguration, "update-bucket-metadata-inventory-table-configuration", "", false, "Update Bucket Metadata Inventory Table Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3UpdateBucketMetadataJournalTableConfiguration, "update-bucket-metadata-journal-table-configuration", "", false, "Update Bucket Metadata Journal Table Configuration")
	_s3Cmd.Flags().BoolVarP(&_s3UpdateObjectEncryption, "update-object-encryption", "", false, "Update Object Encryption")
	_s3Cmd.Flags().BoolVarP(&_s3UploadPart, "upload-part", "", false, "Upload Part")
	_s3Cmd.Flags().BoolVarP(&_s3UploadPartCopy, "upload-part-copy", "", false, "Upload Part Copy")
	_s3Cmd.Flags().BoolVarP(&_s3WriteGetObjectResponse, "write-get-object-response", "", false, "Write Get Object Response")

}
