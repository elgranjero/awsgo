package customerprofiles

// MergeProfiles is generated as a reference stub.
// Executable command wiring lives under cmd/customerprofiles.go.
//
// Runs an AWS Lambda job that does the following:
//
// - All the profileKeys in the ProfileToBeMerged will be moved to the main
// profile.
//
// - All the objects in the ProfileToBeMerged will be moved to the main profile.
//
// - All the ProfileToBeMerged will be deleted at the end.
//
// - All the profileKeys in the ProfileIdsToBeMerged will be moved to the main
// profile.
//
// - Standard fields are merged as follows:
//
// - Fields are always "union"-ed if there are no conflicts in standard fields
// or attributeKeys.
//
// - When there are conflicting fields:
//
// - If no SourceProfileIds entry is specified, the main Profile value is always
// taken.
//
// - If a SourceProfileIds entry is specified, the specified profileId is always
// taken, even if it is a NULL value.
//
// You can use MergeProfiles together with [GetMatches], which returns potentially matching
// profiles, or use it with the results of another matching system. After profiles
// have been merged, they cannot be separated (unmerged).
//
// [GetMatches]: https://docs.aws.amazon.com/customerprofiles/latest/APIReference/API_GetMatches.html
