package customerprofiles

// GetAutoMergingPreview is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Tests the auto-merging settings of your Identity Resolution Job without merging
// your data. It randomly selects a sample of matching groups from the existing
// matching results, and applies the automerging settings that you provided. You
// can then view the number of profiles in the sample, the number of matches, and
// the number of profiles identified to be merged. This enables you to evaluate the
// accuracy of the attributes in your matching list.
//
// You can't view which profiles are matched and would be merged.
//
// We strongly recommend you use this API to do a dry run of the automerging
// process before running the Identity Resolution Job. Include at least two
// matching attributes. If your matching list includes too few attributes (such as
// only FirstName or only LastName ), there may be a large number of matches. This
// increases the chances of erroneous merges.
