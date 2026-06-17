package route53domains

// TransferDomainToAnotherAwsAccount is generated as a reference stub.
// Executable command wiring lives under cmd/route53domains.go.
//
// Transfers a domain from the current Amazon Web Services account to another
// Amazon Web Services account. Note the following:
//
// - The Amazon Web Services account that you're transferring the domain to must
// accept the transfer. If the other account doesn't accept the transfer within 3
// days, we cancel the transfer. See [AcceptDomainTransferFromAnotherAwsAccount].
//
// - You can cancel the transfer before the other account accepts it. See [CancelDomainTransferToAnotherAwsAccount].
//
// - The other account can reject the transfer. See [RejectDomainTransferFromAnotherAwsAccount].
//
// When you transfer a domain from one Amazon Web Services account to another,
// Route 53 doesn't transfer the hosted zone that is associated with the domain.
// DNS resolution isn't affected if the domain and the hosted zone are owned by
// separate accounts, so transferring the hosted zone is optional. For information
// about transferring the hosted zone to another Amazon Web Services account, see [Migrating a Hosted Zone to a Different Amazon Web Services Account]
// in the Amazon Route 53 Developer Guide.
//
// Use either [ListOperations] or [GetOperationDetail] to determine whether the operation succeeded. [GetOperationDetail] provides
// additional information, for example, Domain Transfer from Aws Account
// 111122223333 has been cancelled .
//
// [RejectDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_RejectDomainTransferFromAnotherAwsAccount.html
// [AcceptDomainTransferFromAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_AcceptDomainTransferFromAnotherAwsAccount.html
// [Migrating a Hosted Zone to a Different Amazon Web Services Account]: https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/hosted-zones-migrating.html
// [CancelDomainTransferToAnotherAwsAccount]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_CancelDomainTransferToAnotherAwsAccount.html
// [ListOperations]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_ListOperations.html
// [GetOperationDetail]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_GetOperationDetail.html
