package acm

// GetCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acm.go.
//
// Retrieves a certificate and its certificate chain. The certificate may be
// either a public or private certificate issued using the ACM RequestCertificate
// action, or a certificate imported into ACM using the ImportCertificate action.
// The chain consists of the certificate of the issuing CA and the intermediate
// certificates of any other subordinate CAs. All of the certificates are base64
// encoded. You can use [OpenSSL]to decode the certificates and inspect individual fields.
//
// [OpenSSL]: https://wiki.openssl.org/index.php/Command_Line_Utilities
