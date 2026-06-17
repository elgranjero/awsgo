package acmpca

// CreatePermission is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Grants one or more permissions on a private CA to the Certificate Manager (ACM)
// service principal ( acm.amazonaws.com ). These permissions allow ACM to issue
// and renew ACM certificates that reside in the same Amazon Web Services account
// as the CA.
//
// You can list current permissions with the [ListPermissions] action and revoke them with the [DeletePermission]
// action.
//
// About Permissions
//
// - If the private CA and the certificates it issues reside in the same
// account, you can use CreatePermission to grant permissions for ACM to carry
// out automatic certificate renewals.
//
// - For automatic certificate renewal to succeed, the ACM service principal
// needs permissions to create, retrieve, and list certificates.
//
// - If the private CA and the ACM certificates reside in different accounts,
// then permissions cannot be used to enable automatic renewals. Instead, the ACM
// certificate owner must set up a resource-based policy to enable cross-account
// issuance and renewals. For more information, see [Using a Resource Based Policy with Amazon Web Services Private CA].
//
// [Using a Resource Based Policy with Amazon Web Services Private CA]: https://docs.aws.amazon.com/privateca/latest/userguide/pca-rbp.html
// [ListPermissions]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ListPermissions.html
// [DeletePermission]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_DeletePermission.html
