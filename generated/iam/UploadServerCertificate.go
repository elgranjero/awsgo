package iam

// UploadServerCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Uploads a server certificate entity for the Amazon Web Services account. The
// server certificate entity includes a public key certificate, a private key, and
// an optional certificate chain, which should all be PEM-encoded.
//
// We recommend that you use [Certificate Manager] to provision, manage, and deploy your server
// certificates. With ACM you can request a certificate, deploy it to Amazon Web
// Services resources, and let ACM handle certificate renewals for you.
// Certificates provided by ACM are free. For more information about using ACM, see
// the [Certificate Manager User Guide].
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic includes a list of Amazon Web Services services that can
// use the server certificates that you manage with IAM.
//
// For information about the number of server certificates you can upload, see [IAM and STS quotas] in
// the IAM User Guide.
//
// Because the body of the public key certificate, private key, and the
// certificate chain can be large, you should use POST rather than GET when calling
// UploadServerCertificate . For information about setting up signatures and
// authorization through the API, see [Signing Amazon Web Services API requests]in the Amazon Web Services General
// Reference. For general information about using the Query API with IAM, see [Calling the API by making HTTP query requests]in
// the IAM User Guide.
//
// [Certificate Manager]: https://docs.aws.amazon.com/acm/
// [Certificate Manager User Guide]: https://docs.aws.amazon.com/acm/latest/userguide/
// [IAM and STS quotas]: https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_iam-quotas.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [Signing Amazon Web Services API requests]: https://docs.aws.amazon.com/general/latest/gr/signing_aws_api_requests.html
// [Calling the API by making HTTP query requests]: https://docs.aws.amazon.com/IAM/latest/UserGuide/programming.html
