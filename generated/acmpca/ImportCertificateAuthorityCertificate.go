package acmpca

// ImportCertificateAuthorityCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Imports a signed private CA certificate into Amazon Web Services Private CA.
// This action is used when you are using a chain of trust whose root is located
// outside Amazon Web Services Private CA. Before you can call this action, the
// following preparations must in place:
//
// - In Amazon Web Services Private CA, call the [CreateCertificateAuthority]action to create the private CA
// that you plan to back with the imported certificate.
//
// - Call the [GetCertificateAuthorityCsr]action to generate a certificate signing request (CSR).
//
// - Sign the CSR using a root or intermediate CA hosted by either an
// on-premises PKI hierarchy or by a commercial CA.
//
// - Create a certificate chain and copy the signed certificate and the
// certificate chain to your working directory.
//
// Amazon Web Services Private CA supports three scenarios for installing a CA
// certificate:
//
// - Installing a certificate for a root CA hosted by Amazon Web Services
// Private CA.
//
// - Installing a subordinate CA certificate whose parent authority is hosted by
// Amazon Web Services Private CA.
//
// - Installing a subordinate CA certificate whose parent authority is
// externally hosted.
//
// The following additional requirements apply when you import a CA certificate.
//
// - Only a self-signed certificate can be imported as a root CA.
//
// - A self-signed certificate cannot be imported as a subordinate CA.
//
// - Your certificate chain must not include the private CA certificate that you
// are importing.
//
// - Your root CA must be the last certificate in your chain. The subordinate
// certificate, if any, that your root CA signed must be next to last. The
// subordinate certificate signed by the preceding subordinate CA must come next,
// and so on until your chain is built.
//
// - The chain must be PEM-encoded.
//
// - The maximum allowed size of a certificate is 32 KB.
//
// - The maximum allowed size of a certificate chain is 2 MB.
//
// # Enforcement of Critical Constraints
//
// Amazon Web Services Private CA allows the following extensions to be marked
// critical in the imported CA certificate or chain.
//
// - Authority key identifier
//
// - Basic constraints (must be marked critical)
//
// - Certificate policies
//
// - Extended key usage
//
// - Inhibit anyPolicy
//
// - Issuer alternative name
//
// - Key usage
//
// - Name constraints
//
// - Policy mappings
//
// - Subject alternative name
//
// - Subject directory attributes
//
// - Subject key identifier
//
// - Subject information access
//
// Amazon Web Services Private CA rejects the following extensions when they are
// marked critical in an imported CA certificate or chain.
//
// - Authority information access
//
// - CRL distribution points
//
// - Freshest CRL
//
// - Policy constraints
//
// Amazon Web Services Private Certificate Authority will also reject any other
// extension marked as critical not contained on the preceding list of allowed
// extensions.
//
// [GetCertificateAuthorityCsr]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_GetCertificateAuthorityCsr.html
// [CreateCertificateAuthority]: https://docs.aws.amazon.com/privateca/latest/APIReference/API_CreateCertificateAuthority.html
