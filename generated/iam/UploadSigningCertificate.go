package iam

// UploadSigningCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Uploads an X.509 signing certificate and associates it with the specified IAM
// user. Some Amazon Web Services services require you to use certificates to
// validate requests that are signed with a corresponding private key. When you
// upload the certificate, its default status is Active .
//
// For information about when you would use an X.509 signing certificate, see [Managing server certificates in IAM] in
// the IAM User Guide.
//
// If the UserName is not specified, the IAM user name is determined implicitly
// based on the Amazon Web Services access key ID used to sign the request. This
// operation works for access keys under the Amazon Web Services account.
// Consequently, you can use this operation to manage Amazon Web Services account
// root user credentials even if the Amazon Web Services account has no associated
// users.
//
// Because the body of an X.509 certificate can be large, you should use POST
// rather than GET when calling UploadSigningCertificate . For information about
// setting up signatures and authorization through the API, see [Signing Amazon Web Services API requests]in the Amazon Web
// Services General Reference. For general information about using the Query API
// with IAM, see [Making query requests]in the IAM User Guide.
//
// [Managing server certificates in IAM]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Making query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/IAM_UsingQueryAPI.html
// [Signing Amazon Web Services API requests]: https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html
