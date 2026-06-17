package servicecatalog

// UpdatePortfolioShare is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Updates the specified portfolio share. You can use this API to enable or
// disable TagOptions sharing or Principal sharing for an existing portfolio
// share.
//
// The portfolio share cannot be updated if the CreatePortfolioShare operation is
// IN_PROGRESS , as the share is not available to recipient entities. In this case,
// you must wait for the portfolio share to be completed.
//
// You must provide the accountId or organization node in the input, but not both.
//
// If the portfolio is shared to both an external account and an organization
// node, and both shares need to be updated, you must invoke UpdatePortfolioShare
// separately for each share type.
//
// This API cannot be used for removing the portfolio share. You must use
// DeletePortfolioShare API for that action.
//
// When you associate a principal with portfolio, a potential privilege escalation
// path may occur when that portfolio is then shared with other accounts. For a
// user in a recipient account who is not an Service Catalog Admin, but still has
// the ability to create Principals (Users/Groups/Roles), that user could create a
// role that matches a principal name association for the portfolio. Although this
// user may not know which principal names are associated through Service Catalog,
// they may be able to guess the user. If this potential escalation path is a
// concern, then Service Catalog recommends using PrincipalType as IAM . With this
// configuration, the PrincipalARN must already exist in the recipient account
// before it can be associated.
