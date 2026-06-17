package acmpca

// GetCertificateAuthorityCsr is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Retrieves the certificate signing request (CSR) for your private certificate
// authority (CA). The CSR is created when you call the [CreateCertificateAuthority]action. Sign the CSR with
// your Amazon Web Services Private CA-hosted or on-premises root or subordinate
// CA. Then import the signed certificate back into Amazon Web Services Private CA
// by calling the [ImportCertificateAuthorityCertificate]action. The CSR is returned as a base64 PEM-encoded string.
//
// [ImportCertificateAuthorityCertificate]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_ImportCertificateAuthorityCertificate.html
// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
