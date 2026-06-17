package elasticbeanstalk

// UpdateTagsForResource is generated as a reference stub.
// Executable command wiring lives under cmd/elasticbeanstalk.go.
//
// Update the list of tags applied to an AWS Elastic Beanstalk resource. Two lists
// can be passed: TagsToAdd for tags to add or update, and TagsToRemove .
//
// Elastic Beanstalk supports tagging of all of its resources. For details about
// resource tagging, see [Tagging Application Resources].
//
// If you create a custom IAM user policy to control permission to this operation,
// specify one of the following two virtual actions (or both) instead of the API
// operation name:
//
// elasticbeanstalk:AddTags Controls permission to call UpdateTagsForResource and
// pass a list of tags to add in the TagsToAdd parameter.
//
// elasticbeanstalk:RemoveTags Controls permission to call UpdateTagsForResource
// and pass a list of tag keys to remove in the TagsToRemove parameter.
//
// For details about creating a custom user policy, see [Creating a Custom User Policy].
//
// [Creating a Custom User Policy]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.iam.managed-policies.html#AWSHowTo.iam.policies
// [Tagging Application Resources]: https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/applications-tagging-resources.html
