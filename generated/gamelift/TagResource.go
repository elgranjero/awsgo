package gamelift

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Assigns a tag to an Amazon GameLift Servers resource. You can use tags to
// organize resources, create IAM permissions policies to manage access to groups
// of resources, customize Amazon Web Services cost breakdowns, and more. This
// operation handles the permissions necessary to manage tags for Amazon GameLift
// Servers resources that support tagging.
//
// To add a tag to a resource, specify the unique ARN value for the resource and
// provide a tag list containing one or more tags. The operation succeeds even if
// the list includes tags that are already assigned to the resource.
//
// # Learn more
//
// [Tagging Amazon Web Services Resources]in the Amazon Web Services General Reference
//
// [Amazon Web Services Tagging Strategies]
//
// # Related actions
//
// [All APIs by task]
//
// [Tagging Amazon Web Services Resources]: https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html
// [Amazon Web Services Tagging Strategies]: http://aws.amazon.com/answers/account-management/aws-tagging-strategies/
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
