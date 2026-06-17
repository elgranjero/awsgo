package synthetics

// CreateGroup is generated as a reference stub.
// Executable command wiring lives under cmd/synthetics.go.
//
// Creates a group which you can use to associate canaries with each other,
// including cross-Region canaries. Using groups can help you with managing and
// automating your canaries, and you can also view aggregated run results and
// statistics for all canaries in a group.
//
// Groups are global resources. When you create a group, it is replicated across
// Amazon Web Services Regions, and you can view it and add canaries to it from any
// Region. Although the group ARN format reflects the Region name where it was
// created, a group is not constrained to any Region. This means that you can put
// canaries from multiple Regions into the same group, and then use that group to
// view and manage all of those canaries in a single view.
//
// Groups are supported in all Regions except the Regions that are disabled by
// default. For more information about these Regions, see [Enabling a Region].
//
// Each group can contain as many as 10 canaries. You can have as many as 20
// groups in your account. Any single canary can be a member of up to 10 groups.
//
// [Enabling a Region]: https://docs.aws.amazon.com/general/latest/gr/rande-manage.html#rande-manage-enable
