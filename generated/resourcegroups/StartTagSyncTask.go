package resourcegroups

// StartTagSyncTask is generated as a reference stub.
// Executable command wiring lives under cmd/resourcegroups.go.
//
// Creates a new tag-sync task to onboard and sync resources tagged with a
// specific tag key-value pair to an application. To start a tag-sync task, you
// need a [resource tagging role]. The resource tagging role grants permissions to tag and untag
// applications resources and must include a trust policy that allows Resource
// Groups to assume the role and perform resource tagging tasks on your behalf.
//
// For instructions on creating a tag-sync task, see [Create a tag-sync using the Resource Groups API] in the Amazon Web Services
// Service Catalog AppRegistry Administrator Guide.
//
// # Minimum permissions
//
// To run this command, you must have the following permissions:
//
// - resource-groups:StartTagSyncTask on the application group
//
// - resource-groups:CreateGroup
//
// - iam:PassRole on the role provided in the request
//
// [resource tagging role]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/app-tag-sync.html#tag-sync-role
// [Create a tag-sync using the Resource Groups API]: https://docs.aws.amazon.com/servicecatalog/latest/arguide/app-tag-sync.html#create-tag-sync
