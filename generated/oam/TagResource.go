package oam

// TagResource is generated as a reference stub.
// Executable command wiring lives under cmd/oam.go.
//
// Assigns one or more tags (key-value pairs) to the specified resource. Both
// sinks and links can be tagged.
//
// Tags can help you organize and categorize your resources. You can also use them
// to scope user permissions by granting a user permission to access or change only
// resources with certain tag values.
//
// Tags don't have any semantic meaning to Amazon Web Services and are interpreted
// strictly as strings of characters.
//
// You can use the TagResource action with a resource that already has tags. If
// you specify a new tag key for the alarm, this tag is appended to the list of
// tags associated with the alarm. If you specify a tag key that is already
// associated with the alarm, the new tag value that you specify replaces the
// previous value for that tag.
//
// You can associate as many as 50 tags with a resource.
//
// Unlike tagging permissions in other Amazon Web Services services, to tag or
// untag links and sinks you must have the oam:ResourceTag permission. The
// iam:ResourceTag permission does not allow you to tag and untag links and sinks.
