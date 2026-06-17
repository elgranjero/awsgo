package quicksight

// DeleteAccountSubscription is generated as a reference stub.
// Executable command wiring lives under cmd/quicksight.go.
//
// Deleting your Quick Sight account subscription has permanent, irreversible
// consequences across all Amazon Web Services regions:
//
// - Global deletion – Running this operation from any single region will delete
// your Quick Sight account and all data in every Amazon Web Services region where
// you have Quick Sight resources.
//
// - Complete data loss – All dashboards, analyses, datasets, data sources, and
// custom visuals will be permanently deleted across all regions.
//
// - Embedded content failure – All embedded dashboards and visuals in your
// applications will immediately stop working and display errors to end users.
//
// - Shared resources removed – All shared dashboards, folders, and resources
// will become inaccessible to other users and external recipients.
//
// - User access terminated – All Quick Sight users in your account will lose
// access immediately, including authors, readers, and administrators.
//
// - No recovery possible – Once deleted, your Quick Sight account and all
// associated data cannot be restored.
//
// Consider exporting critical dashboards and data before proceeding with account
// deletion.
//
// Use the DeleteAccountSubscription operation to delete an Quick Sight account.
// This operation will result in an error message if you have configured your
// account termination protection settings to True . To change this setting and
// delete your account, call the UpdateAccountSettings API and set the value of
// the TerminationProtectionEnabled parameter to False , then make another call to
// the DeleteAccountSubscription API.
