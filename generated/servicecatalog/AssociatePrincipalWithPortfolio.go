package servicecatalog

// AssociatePrincipalWithPortfolio is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Associates the specified principal ARN with the specified portfolio.
//
// If you share the portfolio with principal name sharing enabled, the PrincipalARN
// association is included in the share.
//
// The PortfolioID , PrincipalARN , and PrincipalType parameters are required.
//
// You can associate a maximum of 10 Principals with a portfolio using
// PrincipalType as IAM_PATTERN .
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
