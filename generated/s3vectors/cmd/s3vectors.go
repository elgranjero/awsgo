package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3vectors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// s3vectorsCmd represents the s3vectors command
var _s3vectorsCmd = &cobra.Command{
	Use:   "s3vectors",
	Short: "AWS s3vectors CLI",
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
		client := s3vectors.NewFromConfig(cfg)
		if _s3vectorsCreateIndex {
			s3vectors_CreateIndex(cfg, client)
			return
		}
		if _s3vectorsCreateVectorBucket {
			s3vectors_CreateVectorBucket(cfg, client)
			return
		}
		if _s3vectorsDeleteIndex {
			s3vectors_DeleteIndex(cfg, client)
			return
		}
		if _s3vectorsDeleteVectorBucket {
			s3vectors_DeleteVectorBucket(cfg, client)
			return
		}
		if _s3vectorsDeleteVectorBucketPolicy {
			s3vectors_DeleteVectorBucketPolicy(cfg, client)
			return
		}
		if _s3vectorsDeleteVectors {
			s3vectors_DeleteVectors(cfg, client)
			return
		}
		if _s3vectorsGetIndex {
			s3vectors_GetIndex(cfg, client)
			return
		}
		if _s3vectorsGetVectorBucket {
			s3vectors_GetVectorBucket(cfg, client)
			return
		}
		if _s3vectorsGetVectorBucketPolicy {
			s3vectors_GetVectorBucketPolicy(cfg, client)
			return
		}
		if _s3vectorsGetVectors {
			s3vectors_GetVectors(cfg, client)
			return
		}
		if _s3vectorsListIndexes {
			s3vectors_ListIndexes(cfg, client)
			return
		}
		if _s3vectorsListTagsForResource {
			s3vectors_ListTagsForResource(cfg, client)
			return
		}
		if _s3vectorsListVectorBuckets {
			s3vectors_ListVectorBuckets(cfg, client)
			return
		}
		if _s3vectorsListVectors {
			s3vectors_ListVectors(cfg, client)
			return
		}
		if _s3vectorsPutVectorBucketPolicy {
			s3vectors_PutVectorBucketPolicy(cfg, client)
			return
		}
		if _s3vectorsPutVectors {
			s3vectors_PutVectors(cfg, client)
			return
		}
		if _s3vectorsQueryVectors {
			s3vectors_QueryVectors(cfg, client)
			return
		}
		if _s3vectorsTagResource {
			s3vectors_TagResource(cfg, client)
			return
		}
		if _s3vectorsUntagResource {
			s3vectors_UntagResource(cfg, client)
			return
		}

	},
}

var (
	_s3vectorsCreateIndex              bool
	_s3vectorsCreateVectorBucket       bool
	_s3vectorsDeleteIndex              bool
	_s3vectorsDeleteVectorBucket       bool
	_s3vectorsDeleteVectorBucketPolicy bool
	_s3vectorsDeleteVectors            bool
	_s3vectorsGetIndex                 bool
	_s3vectorsGetVectorBucket          bool
	_s3vectorsGetVectorBucketPolicy    bool
	_s3vectorsGetVectors               bool
	_s3vectorsListIndexes              bool
	_s3vectorsListTagsForResource      bool
	_s3vectorsListVectorBuckets        bool
	_s3vectorsListVectors              bool
	_s3vectorsPutVectorBucketPolicy    bool
	_s3vectorsPutVectors               bool
	_s3vectorsQueryVectors             bool
	_s3vectorsTagResource              bool
	_s3vectorsUntagResource            bool

	_s3vectorsDataType                string
	_s3vectorsDimension               string
	_s3vectorsDistanceMetric          string
	_s3vectorsEncryptionConfiguration string
	_s3vectorsFilter                  string
	_s3vectorsIndexArn                string
	_s3vectorsIndexName               string
	_s3vectorsKeys                    []string
	_s3vectorsMaxResults              string
	_s3vectorsMetadataConfiguration   string
	_s3vectorsNextToken               string
	_s3vectorsPolicy                  string
	_s3vectorsPrefix                  string
	_s3vectorsQueryVector             string
	_s3vectorsResourceArn             string
	_s3vectorsReturnData              string
	_s3vectorsReturnDistance          string
	_s3vectorsReturnMetadata          string
	_s3vectorsSegmentCount            string
	_s3vectorsSegmentIndex            string
	_s3vectorsTagKeys                 []string
	_s3vectorsTags                    string
	_s3vectorsTopK                    string
	_s3vectorsVectorBucketArn         string
	_s3vectorsVectorBucketName        string
	_s3vectorsVectors                 string
)

// Creates a vector index within a vector bucket. To specify the vector bucket,
// you must use either the vector bucket name or the vector bucket Amazon Resource
// Name (ARN).
//
// Permissions You must have the s3vectors:CreateIndex permission to use this
// operation.
//
// You must have the s3vectors:TagResource permission in addition to
// s3vectors:CreateIndex permission to create a vector index with tags.
func s3vectors_CreateIndex(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.CreateIndexInput{
		// DataType: types.DataType, // Required
		// Dimension: *int32, // Required
		// DistanceMetric: types.DistanceMetric, // Required
		// IndexName: *string, // Required
	}

	if len(_s3vectorsDataType) > 0 {
		if err := assignInputField(input, "DataType", _s3vectorsDataType); err != nil {
			log.Errorf("invalid --data-type: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsDimension) > 0 {
		if err := assignInputField(input, "Dimension", _s3vectorsDimension); err != nil {
			log.Errorf("invalid --dimension: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsDistanceMetric) > 0 {
		if err := assignInputField(input, "DistanceMetric", _s3vectorsDistanceMetric); err != nil {
			log.Errorf("invalid --distance-metric: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _s3vectorsEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsMetadataConfiguration) > 0 {
		if err := assignInputField(input, "MetadataConfiguration", _s3vectorsMetadataConfiguration); err != nil {
			log.Errorf("invalid --metadata-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsTags) > 0 {
		if err := assignInputField(input, "Tags", _s3vectorsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.CreateIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Creates a vector bucket in the Amazon Web Services Region that you want your
// bucket to be in.
//
// Permissions You must have the s3vectors:CreateVectorBucket permission to use
// this operation.
//
// You must have the s3vectors:TagResource permission in addition to
// s3vectors:CreateVectorBucket permission to create a vector bucket with tags.
func s3vectors_CreateVectorBucket(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.CreateVectorBucketInput{
		// VectorBucketName: *string, // Required
	}

	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}
	if len(_s3vectorsEncryptionConfiguration) > 0 {
		if err := assignInputField(input, "EncryptionConfiguration", _s3vectorsEncryptionConfiguration); err != nil {
			log.Errorf("invalid --encryption-configuration: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsTags) > 0 {
		if err := assignInputField(input, "Tags", _s3vectorsTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.CreateVectorBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a vector index. To specify the vector index, you can either use both
// the vector bucket name and vector index name, or use the vector index Amazon
// Resource Name (ARN).
//
// Permissions You must have the s3vectors:DeleteIndex permission to use this
// operation.
func s3vectors_DeleteIndex(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.DeleteIndexInput{}

	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.DeleteIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a vector bucket. All vector indexes in the vector bucket must be
// deleted before the vector bucket can be deleted. To perform this operation, you
// must use either the vector bucket name or the vector bucket Amazon Resource Name
// (ARN).
//
// Permissions You must have the s3vectors:DeleteVectorBucket permission to use
// this operation.
func s3vectors_DeleteVectorBucket(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.DeleteVectorBucketInput{}

	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.DeleteVectorBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes a vector bucket policy. To specify the bucket, you must use either the
// vector bucket name or the vector bucket Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:DeleteVectorBucketPolicy permission to
// use this operation.
func s3vectors_DeleteVectorBucketPolicy(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.DeleteVectorBucketPolicyInput{}

	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.DeleteVectorBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Deletes one or more vectors in a vector index. To specify the vector index, you
// can either use both the vector bucket name and vector index name, or use the
// vector index Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:DeleteVectors permission to use this
// operation.
func s3vectors_DeleteVectors(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.DeleteVectorsInput{
		// Keys: []string, // Required
	}

	if len(_s3vectorsKeys) > 0 {
		input.Keys = append([]string(nil), _s3vectorsKeys...)
	}
	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.DeleteVectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns vector index attributes. To specify the vector index, you can either
// use both the vector bucket name and the vector index name, or use the vector
// index Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:GetIndex permission to use this
// operation.
func s3vectors_GetIndex(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.GetIndexInput{}

	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.GetIndex(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns vector bucket attributes. To specify the bucket, you must use either
// the vector bucket name or the vector bucket Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:GetVectorBucket permission to use this
// operation.
func s3vectors_GetVectorBucket(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.GetVectorBucketInput{}

	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.GetVectorBucket(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Gets details about a vector bucket policy. To specify the bucket, you must use
// either the vector bucket name or the vector bucket Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:GetVectorBucketPolicy permission to use
// this operation.
func s3vectors_GetVectorBucketPolicy(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.GetVectorBucketPolicyInput{}

	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.GetVectorBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns vector attributes. To specify the vector index, you can either use both
// the vector bucket name and the vector index name, or use the vector index Amazon
// Resource Name (ARN).
//
// Permissions You must have the s3vectors:GetVectors permission to use this
// operation.
func s3vectors_GetVectors(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.GetVectorsInput{
		// Keys: []string, // Required
	}

	if len(_s3vectorsKeys) > 0 {
		input.Keys = append([]string(nil), _s3vectorsKeys...)
	}
	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsReturnData) > 0 {
		if err := assignInputField(input, "ReturnData", _s3vectorsReturnData); err != nil {
			log.Errorf("invalid --return-data: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsReturnMetadata) > 0 {
		if err := assignInputField(input, "ReturnMetadata", _s3vectorsReturnMetadata); err != nil {
			log.Errorf("invalid --return-metadata: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.GetVectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all the vector indexes within the specified vector bucket. To
// specify the bucket, you must use either the vector bucket name or the vector
// bucket Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:ListIndexes permission to use this
// operation.
func s3vectors_ListIndexes(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.ListIndexesInput{}

	if len(_s3vectorsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3vectorsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsNextToken) > 0 {
		input.NextToken = aws.String(_s3vectorsNextToken)
	}
	if len(_s3vectorsPrefix) > 0 {
		input.Prefix = aws.String(_s3vectorsPrefix)
	}
	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if disablePaginator() {
		if resp, err := client.ListIndexes(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3vectors.ListIndexesOutput
	p := s3vectors.NewListIndexesPaginator(client, input)
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

// Lists all of the tags applied to a specified Amazon S3 Vectors resource. Each
// tag is a label consisting of a key and value pair. Tags can help you organize,
// track costs for, and control access to resources.
//
// For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources].
//
// Permissions For vector buckets and vector indexes, you must have the
// s3vectors:ListTagsForResource permission to use this operation.
//
// [Managing tags for Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#manage-tags
func s3vectors_ListTagsForResource(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.ListTagsForResourceInput{
		// ResourceArn: *string, // Required
	}

	if len(_s3vectorsResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3vectorsResourceArn)
	}

	if resp, err := client.ListTagsForResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a list of all the vector buckets that are owned by the authenticated
// sender of the request.
//
// Permissions You must have the s3vectors:ListVectorBuckets permission to use
// this operation.
func s3vectors_ListVectorBuckets(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.ListVectorBucketsInput{}

	if len(_s3vectorsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3vectorsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsNextToken) > 0 {
		input.NextToken = aws.String(_s3vectorsNextToken)
	}
	if len(_s3vectorsPrefix) > 0 {
		input.Prefix = aws.String(_s3vectorsPrefix)
	}

	if disablePaginator() {
		if resp, err := client.ListVectorBuckets(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3vectors.ListVectorBucketsOutput
	p := s3vectors.NewListVectorBucketsPaginator(client, input)
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

// List vectors in the specified vector index. To specify the vector index, you
// can either use both the vector bucket name and the vector index name, or use the
// vector index Amazon Resource Name (ARN).
//
// ListVectors operations proceed sequentially; however, for faster performance on
// a large number of vectors in a vector index, applications can request a parallel
// ListVectors operation by providing the segmentCount and segmentIndex parameters.
//
// Permissions You must have the s3vectors:ListVectors permission to use this
// operation. Additional permissions are required based on the request parameters
// you specify:
//
// - With only s3vectors:ListVectors permission, you can list vector keys when
// returnData and returnMetadata are both set to false or not specified..
//
// - If you set returnData or returnMetadata to true, you must have both
// s3vectors:ListVectors and s3vectors:GetVectors permissions. The request fails
// with a 403 Forbidden error if you request vector data or metadata without the
// s3vectors:GetVectors permission.
func s3vectors_ListVectors(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.ListVectorsInput{}

	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _s3vectorsMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsNextToken) > 0 {
		input.NextToken = aws.String(_s3vectorsNextToken)
	}
	if len(_s3vectorsReturnData) > 0 {
		if err := assignInputField(input, "ReturnData", _s3vectorsReturnData); err != nil {
			log.Errorf("invalid --return-data: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsReturnMetadata) > 0 {
		if err := assignInputField(input, "ReturnMetadata", _s3vectorsReturnMetadata); err != nil {
			log.Errorf("invalid --return-metadata: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsSegmentCount) > 0 {
		if err := assignInputField(input, "SegmentCount", _s3vectorsSegmentCount); err != nil {
			log.Errorf("invalid --segment-count: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsSegmentIndex) > 0 {
		if err := assignInputField(input, "SegmentIndex", _s3vectorsSegmentIndex); err != nil {
			log.Errorf("invalid --segment-index: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if disablePaginator() {
		if resp, err := client.ListVectors(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*s3vectors.ListVectorsOutput
	p := s3vectors.NewListVectorsPaginator(client, input)
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

// Creates a bucket policy for a vector bucket. To specify the bucket, you must
// use either the vector bucket name or the vector bucket Amazon Resource Name
// (ARN).
//
// Permissions You must have the s3vectors:PutVectorBucketPolicy permission to use
// this operation.
func s3vectors_PutVectorBucketPolicy(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.PutVectorBucketPolicyInput{
		// Policy: *string, // Required
	}

	if len(_s3vectorsPolicy) > 0 {
		input.Policy = aws.String(_s3vectorsPolicy)
	}
	if len(_s3vectorsVectorBucketArn) > 0 {
		input.VectorBucketArn = aws.String(_s3vectorsVectorBucketArn)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.PutVectorBucketPolicy(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Adds one or more vectors to a vector index. To specify the vector index, you
// can either use both the vector bucket name and the vector index name, or use the
// vector index Amazon Resource Name (ARN).
//
// For more information about limits, see [Limitations and restrictions] in the Amazon S3 User Guide.
//
// When inserting vector data into your vector index, you must provide the vector
// data as float32 (32-bit floating point) values. If you pass higher-precision
// values to an Amazon Web Services SDK, S3 Vectors converts the values to 32-bit
// floating point before storing them, and GetVectors , ListVectors , and
// QueryVectors operations return the float32 values. Different Amazon Web Services
// SDKs may have different default numeric types, so ensure your vectors are
// properly formatted as float32 values regardless of which SDK you're using. For
// example, in Python, use numpy.float32 or explicitly cast your values.
//
// Permissions You must have the s3vectors:PutVectors permission to use this
// operation.
//
// [Limitations and restrictions]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-vectors-limitations.html
func s3vectors_PutVectors(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.PutVectorsInput{
		// Vectors: []types.PutInputVector, // Required
	}

	if len(_s3vectorsVectors) > 0 {
		if err := assignInputField(input, "Vectors", _s3vectorsVectors); err != nil {
			log.Errorf("invalid --vectors: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.PutVectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Performs an approximate nearest neighbor search query in a vector index using a
// query vector. By default, it returns the keys of approximate nearest neighbors.
// You can optionally include the computed distance (between the query vector and
// each vector in the response), the vector data, and metadata of each vector in
// the response.
//
// To specify the vector index, you can either use both the vector bucket name and
// the vector index name, or use the vector index Amazon Resource Name (ARN).
//
// Permissions You must have the s3vectors:QueryVectors permission to use this
// operation. Additional permissions are required based on the request parameters
// you specify:
//
// - With only s3vectors:QueryVectors permission, you can retrieve vector keys of
// approximate nearest neighbors and computed distances between these vectors. This
// permission is sufficient only when you don't set any metadata filters and don't
// request vector data or metadata (by keeping the returnMetadata parameter set
// to false or not specified).
//
// - If you specify a metadata filter or set returnMetadata to true, you must
// have both s3vectors:QueryVectors and s3vectors:GetVectors permissions. The
// request fails with a 403 Forbidden error if you request metadata filtering,
// vector data, or metadata without the s3vectors:GetVectors permission.
func s3vectors_QueryVectors(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.QueryVectorsInput{
		// QueryVector: types.VectorData, // Required
		// TopK: *int32, // Required
	}

	if len(_s3vectorsQueryVector) > 0 {
		if err := assignInputField(input, "QueryVector", _s3vectorsQueryVector); err != nil {
			log.Errorf("invalid --query-vector: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsTopK) > 0 {
		if err := assignInputField(input, "TopK", _s3vectorsTopK); err != nil {
			log.Errorf("invalid --top-k: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsFilter) > 0 {
		if err := assignInputField(input, "Filter", _s3vectorsFilter); err != nil {
			log.Errorf("invalid --filter: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsIndexArn) > 0 {
		input.IndexArn = aws.String(_s3vectorsIndexArn)
	}
	if len(_s3vectorsIndexName) > 0 {
		input.IndexName = aws.String(_s3vectorsIndexName)
	}
	if len(_s3vectorsReturnDistance) > 0 {
		if err := assignInputField(input, "ReturnDistance", _s3vectorsReturnDistance); err != nil {
			log.Errorf("invalid --return-distance: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsReturnMetadata) > 0 {
		if err := assignInputField(input, "ReturnMetadata", _s3vectorsReturnMetadata); err != nil {
			log.Errorf("invalid --return-metadata: %s", err.Error())
			return
		}
	}
	if len(_s3vectorsVectorBucketName) > 0 {
		input.VectorBucketName = aws.String(_s3vectorsVectorBucketName)
	}

	if resp, err := client.QueryVectors(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies one or more user-defined tags to an Amazon S3 Vectors resource or
// updates existing tags. Each tag is a label consisting of a key and value pair.
// Tags can help you organize, track costs for, and control access to your
// resources. You can add up to 50 tags for each resource.
//
// For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources].
//
// Permissions For vector buckets and vector indexes, you must have the
// s3vectors:TagResource permission to use this operation.
//
// [Managing tags for Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#manage-tags
func s3vectors_TagResource(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.TagResourceInput{
		// ResourceArn: *string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_s3vectorsResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3vectorsResourceArn)
	}
	if len(_s3vectorsTags) > 0 {
		if err := assignInputField(input, "Tags", _s3vectorsTags); err != nil {
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

// Removes the specified user-defined tags from an Amazon S3 Vectors resource. You
// can pass one or more tag keys.
//
// For a list of S3 resources that support tagging, see [Managing tags for Amazon S3 resources].
//
// Permissions For vector buckets and vector indexes, you must have the
// s3vectors:UntagResource permission to use this operation.
//
// [Managing tags for Amazon S3 resources]: https://docs.aws.amazon.com/AmazonS3/latest/userguide/tagging.html#manage-tags
func s3vectors_UntagResource(cfg aws.Config, client *s3vectors.Client) {
	input := &s3vectors.UntagResourceInput{
		// ResourceArn: *string, // Required
		// TagKeys: []string, // Required
	}

	if len(_s3vectorsResourceArn) > 0 {
		input.ResourceArn = aws.String(_s3vectorsResourceArn)
	}
	if len(_s3vectorsTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _s3vectorsTagKeys...)
	}

	if resp, err := client.UntagResource(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_s3vectorsCmd)
	_s3vectorsCmd.Flags().SortFlags = false

	_s3vectorsCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_s3vectorsCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_s3vectorsCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsDataType, "data-type", "", "", "Data Type")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsDimension, "dimension", "", "", "Dimension")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsDistanceMetric, "distance-metric", "", "", "Distance Metric")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsEncryptionConfiguration, "encryption-configuration", "", "", "Encryption Configuration")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsFilter, "filter", "", "", "Filter")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsIndexArn, "index-arn", "", "", "Index ARN")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsIndexName, "index-name", "", "", "Index Name")
	_s3vectorsCmd.Flags().StringSliceVarP(&_s3vectorsKeys, "keys", "", nil, "Keys")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsMaxResults, "max-results", "", "", "Max Results")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsMetadataConfiguration, "metadata-configuration", "", "", "Metadata Configuration")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsNextToken, "next-token", "", "", "Next Token")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsPolicy, "policy", "", "", "Policy")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsPrefix, "prefix", "", "", "Prefix")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsQueryVector, "query-vector", "", "", "Query Vector")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsResourceArn, "resource-arn", "", "", "Resource ARN")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsReturnData, "return-data", "", "", "Return Data")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsReturnDistance, "return-distance", "", "", "Return Distance")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsReturnMetadata, "return-metadata", "", "", "Return Metadata")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsSegmentCount, "segment-count", "", "", "Segment Count")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsSegmentIndex, "segment-index", "", "", "Segment Index")
	_s3vectorsCmd.Flags().StringSliceVarP(&_s3vectorsTagKeys, "tag-keys", "", nil, "Tag Keys")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsTags, "tags", "", "", "Tags")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsTopK, "top-k", "", "", "Top K")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsVectorBucketArn, "vector-bucket-arn", "", "", "Vector Bucket ARN")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsVectorBucketName, "vector-bucket-name", "", "", "Vector Bucket Name")
	_s3vectorsCmd.Flags().StringVarP(&_s3vectorsVectors, "vectors", "", "", "Vectors")

	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsCreateIndex, "create-index", "", false, "Create Index")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsCreateVectorBucket, "create-vector-bucket", "", false, "Create Vector Bucket")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsDeleteIndex, "delete-index", "", false, "Delete Index")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsDeleteVectorBucket, "delete-vector-bucket", "", false, "Delete Vector Bucket")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsDeleteVectorBucketPolicy, "delete-vector-bucket-policy", "", false, "Delete Vector Bucket Policy")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsDeleteVectors, "delete-vectors", "", false, "Delete Vectors")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsGetIndex, "get-index", "", false, "Get Index")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsGetVectorBucket, "get-vector-bucket", "", false, "Get Vector Bucket")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsGetVectorBucketPolicy, "get-vector-bucket-policy", "", false, "Get Vector Bucket Policy")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsGetVectors, "get-vectors", "", false, "Get Vectors")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsListIndexes, "list-indexes", "", false, "List Indexes")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsListTagsForResource, "list-tags-for-resource", "", false, "List Tags For Resource")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsListVectorBuckets, "list-vector-buckets", "", false, "List Vector Buckets")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsListVectors, "list-vectors", "", false, "List Vectors")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsPutVectorBucketPolicy, "put-vector-bucket-policy", "", false, "Put Vector Bucket Policy")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsPutVectors, "put-vectors", "", false, "Put Vectors")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsQueryVectors, "query-vectors", "", false, "Query Vectors")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsTagResource, "tag-resource", "", false, "Tag Resource")
	_s3vectorsCmd.Flags().BoolVarP(&_s3vectorsUntagResource, "untag-resource", "", false, "Untag Resource")

}
