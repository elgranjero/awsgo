package cmd

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroupstaggingapi"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// resourcegroupstaggingapiCmd represents the resourcegroupstaggingapi command
var _resourcegroupstaggingapiCmd = &cobra.Command{
	Use:   "resourcegroupstaggingapi",
	Short: "AWS resourcegroupstaggingapi CLI",
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
		client := resourcegroupstaggingapi.NewFromConfig(cfg)
		if _resourcegroupstaggingapiDescribeReportCreation {
			resourcegroupstaggingapi_DescribeReportCreation(cfg, client)
			return
		}
		if _resourcegroupstaggingapiGetComplianceSummary {
			resourcegroupstaggingapi_GetComplianceSummary(cfg, client)
			return
		}
		if _resourcegroupstaggingapiGetResources {
			resourcegroupstaggingapi_GetResources(cfg, client)
			return
		}
		if _resourcegroupstaggingapiGetTagKeys {
			resourcegroupstaggingapi_GetTagKeys(cfg, client)
			return
		}
		if _resourcegroupstaggingapiGetTagValues {
			resourcegroupstaggingapi_GetTagValues(cfg, client)
			return
		}
		if _resourcegroupstaggingapiListRequiredTags {
			resourcegroupstaggingapi_ListRequiredTags(cfg, client)
			return
		}
		if _resourcegroupstaggingapiStartReportCreation {
			resourcegroupstaggingapi_StartReportCreation(cfg, client)
			return
		}
		if _resourcegroupstaggingapiTagResources {
			resourcegroupstaggingapi_TagResources(cfg, client)
			return
		}
		if _resourcegroupstaggingapiUntagResources {
			resourcegroupstaggingapi_UntagResources(cfg, client)
			return
		}

	},
}

var (
	_resourcegroupstaggingapiDescribeReportCreation bool
	_resourcegroupstaggingapiGetComplianceSummary   bool
	_resourcegroupstaggingapiGetResources           bool
	_resourcegroupstaggingapiGetTagKeys             bool
	_resourcegroupstaggingapiGetTagValues           bool
	_resourcegroupstaggingapiListRequiredTags       bool
	_resourcegroupstaggingapiStartReportCreation    bool
	_resourcegroupstaggingapiTagResources           bool
	_resourcegroupstaggingapiUntagResources         bool

	_resourcegroupstaggingapiExcludeCompliantResources string
	_resourcegroupstaggingapiGroupBy                   string
	_resourcegroupstaggingapiIncludeComplianceDetails  string
	_resourcegroupstaggingapiKey                       string
	_resourcegroupstaggingapiMaxResults                string
	_resourcegroupstaggingapiNextToken                 string
	_resourcegroupstaggingapiPaginationToken           string
	_resourcegroupstaggingapiRegionFilters             []string
	_resourcegroupstaggingapiResourceARNList           []string
	_resourcegroupstaggingapiResourceTypeFilters       []string
	_resourcegroupstaggingapiResourcesPerPage          string
	_resourcegroupstaggingapiS3Bucket                  string
	_resourcegroupstaggingapiTagFilters                string
	_resourcegroupstaggingapiTagKeyFilters             []string
	_resourcegroupstaggingapiTagKeys                   []string
	_resourcegroupstaggingapiTags                      string
	_resourcegroupstaggingapiTagsPerPage               string
	_resourcegroupstaggingapiTargetIdFilters           []string
)

// Describes the status of the StartReportCreation operation.
// You can call this operation only from the organization's management account and
// from the us-east-1 Region.
func resourcegroupstaggingapi_DescribeReportCreation(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.DescribeReportCreationInput{}

	if resp, err := client.DescribeReportCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Returns a table that shows counts of resources that are noncompliant with their
// tag policies.
//
// For more information on tag policies, see [Tag Policies] in the Organizations User Guide.
//
// You can call this operation only from the organization's management account and
// from the us-east-1 Region.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
//
// [Tag Policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
func resourcegroupstaggingapi_GetComplianceSummary(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.GetComplianceSummaryInput{}

	if len(_resourcegroupstaggingapiGroupBy) > 0 {
		if err := assignInputField(input, "GroupBy", _resourcegroupstaggingapiGroupBy); err != nil {
			log.Errorf("invalid --group-by: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupstaggingapiMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiPaginationToken) > 0 {
		input.PaginationToken = aws.String(_resourcegroupstaggingapiPaginationToken)
	}
	if len(_resourcegroupstaggingapiRegionFilters) > 0 {
		input.RegionFilters = append([]string(nil), _resourcegroupstaggingapiRegionFilters...)
	}
	if len(_resourcegroupstaggingapiResourceTypeFilters) > 0 {
		input.ResourceTypeFilters = append([]string(nil), _resourcegroupstaggingapiResourceTypeFilters...)
	}
	if len(_resourcegroupstaggingapiTagKeyFilters) > 0 {
		input.TagKeyFilters = append([]string(nil), _resourcegroupstaggingapiTagKeyFilters...)
	}
	if len(_resourcegroupstaggingapiTargetIdFilters) > 0 {
		input.TargetIdFilters = append([]string(nil), _resourcegroupstaggingapiTargetIdFilters...)
	}

	if disablePaginator() {
		if resp, err := client.GetComplianceSummary(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroupstaggingapi.GetComplianceSummaryOutput
	p := resourcegroupstaggingapi.NewGetComplianceSummaryPaginator(client, input)
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

// Returns all the tagged or previously tagged resources that are located in the
// specified Amazon Web Services Region for the account.
//
// Depending on what information you want returned, you can also specify the
// following:
//
// - Filters that specify what tags and resource types you want returned. The
// response includes all tags that are associated with the requested resources.
//
// - Information about compliance with the account's effective tag policy. For
// more information on tag policies, see [Tag Policies]in the Organizations User Guide.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
//
// GetResources does not return untagged resources.
//
// To find untagged resources in your account, use Amazon Web Services Resource
// Explorer with a query that uses tag:none . For more information, see [Search query syntax reference for Resource Explorer].
//
// [Search query syntax reference for Resource Explorer]: https://docs.aws.amazon.com/resource-explorer/latest/userguide/using-search-query-syntax.html
// [Tag Policies]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_tag-policies.html
func resourcegroupstaggingapi_GetResources(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.GetResourcesInput{}

	if len(_resourcegroupstaggingapiExcludeCompliantResources) > 0 {
		if err := assignInputField(input, "ExcludeCompliantResources", _resourcegroupstaggingapiExcludeCompliantResources); err != nil {
			log.Errorf("invalid --exclude-compliant-resources: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiIncludeComplianceDetails) > 0 {
		if err := assignInputField(input, "IncludeComplianceDetails", _resourcegroupstaggingapiIncludeComplianceDetails); err != nil {
			log.Errorf("invalid --include-compliance-details: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiPaginationToken) > 0 {
		input.PaginationToken = aws.String(_resourcegroupstaggingapiPaginationToken)
	}
	if len(_resourcegroupstaggingapiResourceARNList) > 0 {
		input.ResourceARNList = append([]string(nil), _resourcegroupstaggingapiResourceARNList...)
	}
	if len(_resourcegroupstaggingapiResourceTypeFilters) > 0 {
		input.ResourceTypeFilters = append([]string(nil), _resourcegroupstaggingapiResourceTypeFilters...)
	}
	if len(_resourcegroupstaggingapiResourcesPerPage) > 0 {
		if err := assignInputField(input, "ResourcesPerPage", _resourcegroupstaggingapiResourcesPerPage); err != nil {
			log.Errorf("invalid --resources-per-page: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiTagFilters) > 0 {
		if err := assignInputField(input, "TagFilters", _resourcegroupstaggingapiTagFilters); err != nil {
			log.Errorf("invalid --tag-filters: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiTagsPerPage) > 0 {
		if err := assignInputField(input, "TagsPerPage", _resourcegroupstaggingapiTagsPerPage); err != nil {
			log.Errorf("invalid --tags-per-page: %s", err.Error())
			return
		}
	}

	if disablePaginator() {
		if resp, err := client.GetResources(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroupstaggingapi.GetResourcesOutput
	p := resourcegroupstaggingapi.NewGetResourcesPaginator(client, input)
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

// Returns all tag keys currently in use in the specified Amazon Web Services
// Region for the calling account.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
func resourcegroupstaggingapi_GetTagKeys(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.GetTagKeysInput{}

	if len(_resourcegroupstaggingapiPaginationToken) > 0 {
		input.PaginationToken = aws.String(_resourcegroupstaggingapiPaginationToken)
	}

	if disablePaginator() {
		if resp, err := client.GetTagKeys(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroupstaggingapi.GetTagKeysOutput
	p := resourcegroupstaggingapi.NewGetTagKeysPaginator(client, input)
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

// Returns all tag values for the specified key that are used in the specified
// Amazon Web Services Region for the calling account.
//
// This operation supports pagination, where the response can be sent in multiple
// pages. You should check the PaginationToken response parameter to determine if
// there are additional results available to return. Repeat the query, passing the
// PaginationToken response parameter value as an input to the next request until
// you recieve a null value. A null value for PaginationToken indicates that there
// are no more results waiting to be returned.
func resourcegroupstaggingapi_GetTagValues(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.GetTagValuesInput{
		// Key: *string, // Required
	}

	if len(_resourcegroupstaggingapiKey) > 0 {
		input.Key = aws.String(_resourcegroupstaggingapiKey)
	}
	if len(_resourcegroupstaggingapiPaginationToken) > 0 {
		input.PaginationToken = aws.String(_resourcegroupstaggingapiPaginationToken)
	}

	if disablePaginator() {
		if resp, err := client.GetTagValues(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroupstaggingapi.GetTagValuesOutput
	p := resourcegroupstaggingapi.NewGetTagValuesPaginator(client, input)
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

// Lists the required tags for supported resource types in an Amazon Web Services
// account.
func resourcegroupstaggingapi_ListRequiredTags(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.ListRequiredTagsInput{}

	if len(_resourcegroupstaggingapiMaxResults) > 0 {
		if err := assignInputField(input, "MaxResults", _resourcegroupstaggingapiMaxResults); err != nil {
			log.Errorf("invalid --max-results: %s", err.Error())
			return
		}
	}
	if len(_resourcegroupstaggingapiNextToken) > 0 {
		input.NextToken = aws.String(_resourcegroupstaggingapiNextToken)
	}

	if disablePaginator() {
		if resp, err := client.ListRequiredTags(context.TODO(), input); err != nil {
			log.Errorf("%s", err.Error())
			return
		} else {
			writeOutput(nil, nil, resp, _awsOutput)
		}
		return
	}

	var results []*resourcegroupstaggingapi.ListRequiredTagsOutput
	p := resourcegroupstaggingapi.NewListRequiredTagsPaginator(client, input)
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

// Generates a report that lists all tagged resources in the accounts across your
// organization and tells whether each resource is compliant with the effective tag
// policy. Compliance data is refreshed daily. The report is generated
// asynchronously.
//
// The generated report is saved to the following location:
//
// s3://amzn-s3-demo-bucket/AwsTagPolicies/o-exampleorgid/YYYY-MM-ddTHH:mm:ssZ/report.csv
//
// For more information about evaluating resource compliance with tag policies,
// including the required permissions, review [Permissions for evaluating organization-wide compliance]in the Tagging Amazon Web Services
// Resources and Tag Editor user guide.
//
// You can call this operation only from the organization's management account and
// from the us-east-1 Region.
//
// If the account associated with the identity used to call StartReportCreation is
// different from the account that owns the Amazon S3 bucket, there must be a
// bucket policy attached to the bucket to provide access. For more information,
// review [Amazon S3 bucket policy for report storage]in the Tagging Amazon Web Services Resources and Tag Editor user guide.
//
// [Amazon S3 bucket policy for report storage]: https://docs.aws.amazon.com/tag-editor/latest/userguide/tag-policies-orgs.html#bucket-policy
// [Permissions for evaluating organization-wide compliance]: https://docs.aws.amazon.com/tag-editor/latest/userguide/tag-policies-orgs.html#tag-policies-permissions-org
func resourcegroupstaggingapi_StartReportCreation(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.StartReportCreationInput{
		// S3Bucket: *string, // Required
	}

	if len(_resourcegroupstaggingapiS3Bucket) > 0 {
		input.S3Bucket = aws.String(_resourcegroupstaggingapiS3Bucket)
	}

	if resp, err := client.StartReportCreation(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Applies one or more tags to the specified resources. Note the following:
// - Not all resources can have tags. For a list of services with resources that
// support tagging using this operation, see [Services that support the Resource Groups Tagging API]. If the resource doesn't yet
// support this operation, the resource's service might support tagging using its
// own API operations. For more information, refer to the documentation for that
// service.
//
// - Each resource can have up to 50 tags. For other limits, see [Tag Naming and Usage Conventions]in the Amazon
// Web Services General Reference.
//
// - You can only tag resources that are located in the specified Amazon Web
// Services Region for the Amazon Web Services account.
//
// - To add tags to a resource, you need the necessary permissions for the
// service that the resource belongs to as well as permissions for adding tags. For
// more information, see the documentation for each service.
//
// - When you use the [Amazon Web Services Resource Groups Tagging API]to update tags for Amazon Web Services CloudFormation
// stack sets, Amazon Web Services calls the [Amazon Web Services CloudFormation UpdateStack]UpdateStack operation. This
// operation may initiate additional resource property updates in addition to the
// desired tag updates. To avoid unexpected resource updates, Amazon Web Services
// recommends that you only apply or update tags to your CloudFormation stack sets
// using Amazon Web Services CloudFormation.
//
// Do not store personally identifiable information (PII) or other confidential or
// sensitive information in tags. We use tags to provide you with billing and
// administration services. Tags are not intended to be used for private or
// sensitive data.
//
// # Minimum permissions
//
// In addition to the tag:TagResources permission required by this operation, you
// must also have the tagging permission defined by the service that created the
// resource. For example, to tag an Amazon EC2 instance using the TagResources
// operation, you must have both of the following permissions:
//
// - tag:TagResources
//
// - ec2:CreateTags
//
// In addition, some services might have specific requirements for tagging some
// types of resources. For example, to tag an Amazon S3 bucket, you must also have
// the s3:GetBucketTagging permission. If the expected minimum permissions don't
// work, check the documentation for that service's tagging APIs for more
// information.
//
// [Amazon Web Services CloudFormation UpdateStack]: https://docs.aws.amazon.com/AWSCloudFormation/latest/APIReference/API_UpdateStack.html
// [Amazon Web Services Resource Groups Tagging API]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/overview.html
// [Services that support the Resource Groups Tagging API]: https://docs.aws.amazon.com/resourcegroupstagging/latest/APIReference/supported-services.html
// [Tag Naming and Usage Conventions]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html#tag-conventions
func resourcegroupstaggingapi_TagResources(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.TagResourcesInput{
		// ResourceARNList: []string, // Required
		// Tags: map[string]string, // Required
	}

	if len(_resourcegroupstaggingapiResourceARNList) > 0 {
		input.ResourceARNList = append([]string(nil), _resourcegroupstaggingapiResourceARNList...)
	}
	if len(_resourcegroupstaggingapiTags) > 0 {
		if err := assignInputField(input, "Tags", _resourcegroupstaggingapiTags); err != nil {
			log.Errorf("invalid --tags: %s", err.Error())
			return
		}
	}

	if resp, err := client.TagResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

// Removes the specified tags from the specified resources. When you specify a tag
// key, the action removes both that key and its associated value. The operation
// succeeds even if you attempt to remove tags from a resource that were already
// removed. Note the following:
//
// - To remove tags from a resource, you need the necessary permissions for the
// service that the resource belongs to as well as permissions for removing tags.
// For more information, see the documentation for the service whose resource you
// want to untag.
//
// - You can only tag resources that are located in the specified Amazon Web
// Services Region for the calling Amazon Web Services account.
//
// # Minimum permissions
//
// In addition to the tag:UntagResources permission required by this operation,
// you must also have the remove tags permission defined by the service that
// created the resource. For example, to remove the tags from an Amazon EC2
// instance using the UntagResources operation, you must have both of the
// following permissions:
//
// - tag:UntagResources
//
// - ec2:DeleteTags
//
// In addition, some services might have specific requirements for untagging some
// types of resources. For example, to untag Amazon Web Services Glue Connection,
// you must also have the glue:GetConnection permission. If the expected minimum
// permissions don't work, check the documentation for that service's tagging APIs
// for more information.
func resourcegroupstaggingapi_UntagResources(cfg aws.Config, client *resourcegroupstaggingapi.Client) {
	input := &resourcegroupstaggingapi.UntagResourcesInput{
		// ResourceARNList: []string, // Required
		// TagKeys: []string, // Required
	}

	if len(_resourcegroupstaggingapiResourceARNList) > 0 {
		input.ResourceARNList = append([]string(nil), _resourcegroupstaggingapiResourceARNList...)
	}
	if len(_resourcegroupstaggingapiTagKeys) > 0 {
		input.TagKeys = append([]string(nil), _resourcegroupstaggingapiTagKeys...)
	}

	if resp, err := client.UntagResources(context.TODO(), input); err != nil {
		log.Errorf("%s", err.Error())
		return
	} else {
		writeOutput(nil, nil, resp, _awsOutput)
	}
}

func init() {
	_rootCmd.AddCommand(_resourcegroupstaggingapiCmd)
	_resourcegroupstaggingapiCmd.Flags().SortFlags = false

	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_awsProfile, "profile", "", "", "AWS shared config profile")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_awsRegion, "region", "", "", "Set AWS Region")

	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_awsOutput, "output", "o", "", "Output format: json|yaml|text|table|csv|markdown|html")

	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiExcludeCompliantResources, "exclude-compliant-resources", "", "", "Exclude Compliant Resources")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiGroupBy, "group-by", "", "", "Group By")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiIncludeComplianceDetails, "include-compliance-details", "", "", "Include Compliance Details")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiKey, "key", "", "", "Key")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiMaxResults, "max-results", "", "", "Max Results")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiNextToken, "next-token", "", "", "Next Token")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiPaginationToken, "pagination-token", "", "", "Pagination Token")
	_resourcegroupstaggingapiCmd.Flags().StringSliceVarP(&_resourcegroupstaggingapiRegionFilters, "region-filters", "", nil, "Region Filters")
	_resourcegroupstaggingapiCmd.Flags().StringSliceVarP(&_resourcegroupstaggingapiResourceARNList, "resource-arn-list", "", nil, "Resource ARN List")
	_resourcegroupstaggingapiCmd.Flags().StringSliceVarP(&_resourcegroupstaggingapiResourceTypeFilters, "resource-type-filters", "", nil, "Resource Type Filters")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiResourcesPerPage, "resources-per-page", "", "", "Resources Per Page")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiS3Bucket, "s3-bucket", "", "", "S3 Bucket")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiTagFilters, "tag-filters", "", "", "Tag Filters")
	_resourcegroupstaggingapiCmd.Flags().StringSliceVarP(&_resourcegroupstaggingapiTagKeyFilters, "tag-key-filters", "", nil, "Tag Key Filters")
	_resourcegroupstaggingapiCmd.Flags().StringSliceVarP(&_resourcegroupstaggingapiTagKeys, "tag-keys", "", nil, "Tag Keys")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiTags, "tags", "", "", "Tags")
	_resourcegroupstaggingapiCmd.Flags().StringVarP(&_resourcegroupstaggingapiTagsPerPage, "tags-per-page", "", "", "Tags Per Page")
	_resourcegroupstaggingapiCmd.Flags().StringSliceVarP(&_resourcegroupstaggingapiTargetIdFilters, "target-id-filters", "", nil, "Target ID Filters")

	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiDescribeReportCreation, "describe-report-creation", "", false, "Describe Report Creation")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiGetComplianceSummary, "get-compliance-summary", "", false, "Get Compliance Summary")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiGetResources, "get-resources", "", false, "Get Resources")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiGetTagKeys, "get-tag-keys", "", false, "Get Tag Keys")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiGetTagValues, "get-tag-values", "", false, "Get Tag Values")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiListRequiredTags, "list-required-tags", "", false, "List Required Tags")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiStartReportCreation, "start-report-creation", "", false, "Start Report Creation")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiTagResources, "tag-resources", "", false, "Tag Resources")
	_resourcegroupstaggingapiCmd.Flags().BoolVarP(&_resourcegroupstaggingapiUntagResources, "untag-resources", "", false, "Untag Resources")

}
