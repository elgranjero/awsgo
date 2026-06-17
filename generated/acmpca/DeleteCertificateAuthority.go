package acmpca

// DeleteCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Deletes a private certificate authority (CA). You must provide the Amazon
// Resource Name (ARN) of the private CA that you want to delete. You can find the
// ARN by calling the [ListCertificateAuthorities]action.
//
// Deleting a CA will invalidate other CAs and certificates below it in your CA
// hierarchy.
//
// Before you can delete a CA that you have created and activated, you must
// disable it. To do this, call the [UpdateCertificateAuthority]action and set the CertificateAuthorityStatus
// parameter to DISABLED .
//
// Additionally, you can delete a CA if you are waiting for it to be created (that
// is, the status of the CA is CREATING ). You can also delete it if the CA has
// been created but you haven't yet imported the signed certificate into Amazon Web
// Services Private CA (that is, the status of the CA is PENDING_CERTIFICATE ).
//
// When you successfully call [DeleteCertificateAuthority], the CA's status changes to DELETED . However, the
// CA won't be permanently deleted until the restoration period has passed. By
// default, if you do not set the PermanentDeletionTimeInDays parameter, the CA
// remains restorable for 30 days. You can set the parameter from 7 to 30 days. The
// [DescribeCertificateAuthority]action returns the time remaining in the restoration window of a private CA in
// the DELETED state. To restore an eligible CA, call the [RestoreCertificateAuthority] action.
//
// A private CA can be deleted if it is in the PENDING_CERTIFICATE , CREATING ,
// EXPIRED , DISABLED , or FAILED state. To delete a CA in the ACTIVE state, you
// must first disable it, or else the delete request results in an exception. If
// you are deleting a private CA in the PENDING_CERTIFICATE or DISABLED state, you
// can set the length of its restoration period to 7-30 days. The default is 30.
// During this time, the status is set to DELETED and the CA can be restored. A
// private CA deleted in the CREATING or FAILED state has no assigned restoration
// period and cannot be restored.
//
// [ListCertificateAuthorities]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListCertificateAuthorities.html
// [RestoreCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_RestoreCertificateAuthority.html
// [UpdateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_UpdateCertificateAuthority.html
// [DeleteCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeleteCertificateAuthority.html
// [DescribeCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DescribeCertificateAuthority.html
