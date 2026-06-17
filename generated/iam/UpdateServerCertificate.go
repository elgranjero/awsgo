package iam

// UpdateServerCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Updates the name and/or the path of the specified server certificate stored in
// IAM.
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// You should understand the implications of changing a server certificate's path
// or name. For more information, see [Renaming a server certificate]in the IAM User Guide.
//
// The person making the request (the principal), must have permission to change
// the server certificate with the old name and the new name. For example, to
// change the certificate named ProductionCert to ProdCert , the principal must
// have a policy that allows them to update both certificates. If the principal has
// permission to update the ProductionCert group, but not the ProdCert
// certificate, then the update fails. For more information about permissions, see [Access management]
// in the IAM User Guide.
//
// [Renaming a server certificate]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs_manage.html#RenamingServerCerts
// [Access management]: https://docs.aws.amazon.com/IAM/latest/UserGuide/access.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
