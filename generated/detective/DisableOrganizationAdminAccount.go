package detective

// DisableOrganizationAdminAccount is generated as a reference stub.
// Executable command wiring lives under cmd/detective.go.
//
// Removes the Detective administrator account in the current Region. Deletes the
// organization behavior graph.
//
// Can only be called by the organization management account.
//
// Removing the Detective administrator account does not affect the delegated
// administrator account for Detective in Organizations.
//
// To remove the delegated administrator account in Organizations, use the
// Organizations API. Removing the delegated administrator account also removes the
// Detective administrator account in all Regions, except for Regions where the
// Detective administrator account is the organization management account.
