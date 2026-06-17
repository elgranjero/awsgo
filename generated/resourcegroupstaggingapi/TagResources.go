package resourcegroupstaggingapi

// TagResources is generated as a reference stub.
// Executable command wiring lives under cmd/resourcegroupstaggingapi.go.
//
// Applies one or more tags to the specified resources. Note the following:
//
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
