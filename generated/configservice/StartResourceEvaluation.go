package configservice

// StartResourceEvaluation is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Runs an on-demand evaluation for the specified resource to determine whether
// the resource details will comply with configured Config rules. You can also use
// it for evaluation purposes. Config recommends using an evaluation context. It
// runs an execution against the resource details with all of the Config rules in
// your account that match with the specified proactive mode and resource type.
//
// Ensure you have the cloudformation:DescribeType role setup to validate the
// resource type schema.
//
// You can find the [Resource type schema] in "Amazon Web Services public extensions" within the
// CloudFormation registry or with the following CLI commmand: aws cloudformation
// describe-type --type-name "AWS::S3::Bucket" --type RESOURCE .
//
// For more information, see [Managing extensions through the CloudFormation registry] and [Amazon Web Services resource and property types reference] in the CloudFormation User Guide.
//
// [Resource type schema]: https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/resource-type-schema.html
// [Amazon Web Services resource and property types reference]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-template-resource-type-ref.html
// [Managing extensions through the CloudFormation registry]: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/registry.html#registry-view
