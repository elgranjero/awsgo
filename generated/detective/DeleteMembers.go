package detective

// DeleteMembers is generated as a reference stub.
// Executable command wiring lives under cmd/detective.go.
//
// Removes the specified member accounts from the behavior graph. The removed
// accounts no longer contribute data to the behavior graph. This operation can
// only be called by the administrator account for the behavior graph.
//
// For invited accounts, the removed accounts are deleted from the list of
// accounts in the behavior graph. To restore the account, the administrator
// account must send another invitation.
//
// For organization accounts in the organization behavior graph, the Detective
// administrator account can always enable the organization account again.
// Organization accounts that are not enabled as member accounts are not included
// in the ListMembers results for the organization behavior graph.
//
// An administrator account cannot use DeleteMembers to remove their own account
// from the behavior graph. To disable a behavior graph, the administrator account
// uses the DeleteGraph API method.
