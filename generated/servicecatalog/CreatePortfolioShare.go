package servicecatalog

// CreatePortfolioShare is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Shares the specified portfolio with the specified account or organization node.
// Shares to an organization node can only be created by the management account of
// an organization or by a delegated administrator. You can share portfolios to an
// organization, an organizational unit, or a specific account.
//
// Note that if a delegated admin is de-registered, they can no longer create
// portfolio shares.
//
// AWSOrganizationsAccess must be enabled in order to create a portfolio share to
// an organization node.
//
// You can't share a shared resource, including portfolios that contain a shared
// product.
//
// If the portfolio share with the specified account or organization node already
// exists, this action will have no effect and will not return an error. To update
// an existing share, you must use the UpdatePortfolioShare API instead.
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
