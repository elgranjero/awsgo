package detective

// EnableOrganizationAdminAccount is generated as a reference stub.
// Executable command wiring lives under cmd/detective.go.
//
// Designates the Detective administrator account for the organization in the
// current Region.
//
// If the account does not have Detective enabled, then enables Detective for that
// account and creates a new behavior graph.
//
// Can only be called by the organization management account.
//
// If the organization has a delegated administrator account in Organizations,
// then the Detective administrator account must be either the delegated
// administrator account or the organization management account.
//
// If the organization does not have a delegated administrator account in
// Organizations, then you can choose any account in the organization. If you
// choose an account other than the organization management account, Detective
// calls Organizations to make that account the delegated administrator account for
// Detective. The organization management account cannot be the delegated
// administrator account.
