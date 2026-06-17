package servicecatalog

// DisassociatePrincipalFromPortfolio is generated as a reference stub.
// Executable command wiring lives under cmd/servicecatalog.go.
//
// Disassociates a previously associated principal ARN from a specified portfolio.
//
// The PrincipalType and PrincipalARN must match the
// AssociatePrincipalWithPortfolio call request details. For example, to
// disassociate an association created with a PrincipalARN of PrincipalType IAM
// you must use the PrincipalType IAM when calling
// DisassociatePrincipalFromPortfolio .
//
// For portfolios that have been shared with principal name sharing enabled: after
// disassociating a principal, share recipient accounts will no longer be able to
// provision products in this portfolio using a role matching the name of the
// associated principal.
//
// For more information, review [associate-principal-with-portfolio] in the Amazon Web Services CLI Command Reference.
//
// If you disassociate a principal from a portfolio, with PrincipalType as IAM ,
// the same principal will still have access to the portfolio if it matches one of
// the associated principals of type IAM_PATTERN . To fully remove access for a
// principal, verify all the associated Principals of type IAM_PATTERN , and then
// ensure you disassociate any IAM_PATTERN principals that match the principal
// whose access you are removing.
//
// [associate-principal-with-portfolio]: https://docs.aws.amazon.com/cli/latest/reference/servicecatalog/associate-principal-with-portfolio.html#options
