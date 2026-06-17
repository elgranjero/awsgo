package ssm

// AddTagsToResource is generated as a reference stub.
// Executable command wiring lives under cmd/ssm.go.
//
// Adds or overwrites one or more tags for the specified resource. Tags are
// metadata that you can assign to your automations, documents, managed nodes,
// maintenance windows, Parameter Store parameters, and patch baselines. Tags
// enable you to categorize your resources in different ways, for example, by
// purpose, owner, or environment. Each tag consists of a key and an optional
// value, both of which you define. For example, you could define a set of tags for
// your account's managed nodes that helps you track each node's owner and stack
// level. For example:
//
// - Key=Owner,Value=DbAdmin
//
// - Key=Owner,Value=SysAdmin
//
// - Key=Owner,Value=Dev
//
// - Key=Stack,Value=Production
//
// - Key=Stack,Value=Pre-Production
//
// - Key=Stack,Value=Test
//
// Most resources can have a maximum of 50 tags. Automations can have a maximum of
// 5 tags.
//
// We recommend that you devise a set of tag keys that meets your needs for each
// resource type. Using a consistent set of tag keys makes it easier for you to
// manage your resources. You can search and filter the resources based on the tags
// you add. Tags don't have any semantic meaning to and are interpreted strictly as
// a string of characters.
//
// For more information about using tags with Amazon Elastic Compute Cloud (Amazon
// EC2) instances, see [Tag your Amazon EC2 resources]in the Amazon EC2 User Guide.
//
// [Tag your Amazon EC2 resources]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Tags.html
