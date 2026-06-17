package cloudtrail

// AddTags is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Adds one or more tags to a trail, event data store, dashboard, or channel, up
// to a limit of 50. Overwrites an existing tag's value when a new value is
// specified for an existing tag key. Tag key names must be unique; you cannot have
// two keys with the same name but different values. If you specify a key without a
// value, the tag will be created with the specified key and a value of null. You
// can tag a trail or event data store that applies to all Amazon Web Services
// Regions only from the Region in which the trail or event data store was created
// (also known as its home Region).
