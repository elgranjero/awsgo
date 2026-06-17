package acmpca

// RestoreCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Restores a certificate authority (CA) that is in the DELETED state. You can
// restore a CA during the period that you defined in the
// PermanentDeletionTimeInDays parameter of the [DeleteCertificateAuthority]action. Currently, you can specify
// 7 to 30 days. If you did not specify a PermanentDeletionTimeInDays value, by
// default you can restore the CA at any time in a 30 day period. You can check the
// time remaining in the restoration period of a private CA in the DELETED state
// by calling the [DescribeCertificateAuthority]or [ListCertificateAuthorities] actions. The status of a restored CA is set to its
// pre-deletion status when the RestoreCertificateAuthority action returns. To
// change its status to ACTIVE , call the [UpdateCertificateAuthority] action. If the private CA was in the
// PENDING_CERTIFICATE state at deletion, you must use the [ImportCertificateAuthorityCertificate] action to import a
// certificate authority into the private CA before it can be activated. You cannot
// restore a CA after the restoration period has ended.
//
// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
// [ImportCertificateAuthorityCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html
// [UpdateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html
// [DeleteCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthority.html
// [DescribeCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DescribeCertificateAuthority.html
