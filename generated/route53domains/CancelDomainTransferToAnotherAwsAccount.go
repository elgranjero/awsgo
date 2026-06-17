package route53domains

// CancelDomainTransferToAnotherAwsAccount is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// Cancels the transfer of a domain from the current Amazon Web Services account
// to another Amazon Web Services account. You initiate a transfer betweenAmazon
// Web Services accounts using [TransferDomainToAnotherAwsAccount].
//
// You must cancel the transfer before the other Amazon Web Services account
// accepts the transfer using [AcceptDomainTransferFromAnotherAwsAccount].
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [TransferDomainToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html
// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
