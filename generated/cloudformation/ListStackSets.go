package cloudformation

// ListStackSets is generated as a reference stub.
// Executable command wiring lives under cmd/cloudformation.go.
//
// Returns summary information about StackSets that are associated with the user.
//
// This API provides strongly consistent reads meaning it will always return the
// most up-to-date data.
//
// - [Self-managed permissions] If you set the CallAs parameter to SELF while
// signed in to your Amazon Web Services account, ListStackSets returns all
// self-managed StackSets in your Amazon Web Services account.
//
// - [Service-managed permissions] If you set the CallAs parameter to SELF while
// signed in to the organization's management account, ListStackSets returns all
// StackSets in the management account.
//
// - [Service-managed permissions] If you set the CallAs parameter to
// DELEGATED_ADMIN while signed in to your member account, ListStackSets returns
// all StackSets with service-managed permissions in the management account.
