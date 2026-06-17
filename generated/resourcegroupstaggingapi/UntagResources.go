package resourcegroupstaggingapi

// UntagResources is generated as a reference stub.
// Executable command wiring lives under cmd/resourcegroupstaggingapi.go.
//
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
