package route53domains

// AcceptDomainTransferFromAnotherAwsAccount is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// Accepts the transfer of a domain from another Amazon Web Services account to
// the currentAmazon Web Services account. You initiate a transfer between Amazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// If you use the CLI command at [accept-domain-transfer-from-another-aws-account], use JSON format as input instead of text
// because otherwise CLI will throw an error from domain transfer input that
// includes single quotes.
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [accept-domain-transfer-from-another-aws-account]: https://docs.aws.amazon.com/cli/latest/reference/route53domains/accept-domain-transfer-from-another-aws-account.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
