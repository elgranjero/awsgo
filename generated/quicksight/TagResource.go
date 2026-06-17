package quicksight

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Assigns one or more tags (key-value pairs) to the specified Amazon Quick Sight
// resource.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions, by granting a user permission to access or change
// only resources with certain tag values. You can use the TagResource operation
// with a resource that already has tags. If you specify a new tag key for the
// resource, this tag is appended to the list of tags associated with the resource.
// If you specify a tag key that is already associated with the resource, the new
// tag value that you specify replaces the previous value for that tag.
//
// You can associate as many as 50 tags with a resource. Amazon Quick Sight
// supports tagging on data set, data source, dashboard, template, topic, and user.
//
// Tagging for Amazon Quick Sight works in a similar way to tagging for other
// Amazon Web Services services, except for the following:
//
// - Tags are used to track costs for users in Amazon Quick Sight. You can't tag
// other resources that Amazon Quick Sight costs are based on, such as storage
// capacoty (SPICE), session usage, alert consumption, or reporting units.
//
// - Amazon Quick Sight doesn't currently support the tag editor for Resource
// Groups.
