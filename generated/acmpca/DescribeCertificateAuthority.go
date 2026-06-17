package acmpca

// DescribeCertificateAuthority is generated as a reference stub.
// Executable command wiring lives under cmd/acmpca.go.
//
// Lists information about your private certificate authority (CA) or one that has
// been shared with you. You specify the private CA on input by its ARN (Amazon
// Resource Name). The output contains the status of your CA. This can be any of
// the following:
//
// - CREATING - Amazon Web Services Private CA is creating your private
// certificate authority.
//
// - PENDING_CERTIFICATE - The certificate is pending. You must use your Amazon
// Web Services Private CA-hosted or on-premises root or subordinate CA to sign
// your private CA CSR and then import it into Amazon Web Services Private CA.
//
// - ACTIVE - Your private CA is active.
//
// - DISABLED - Your private CA has been disabled.
//
// - EXPIRED - Your private CA certificate has expired.
//
// - FAILED - Your private CA has failed. Your CA can fail because of problems
// such a network outage or back-end Amazon Web Services failure or other errors. A
// failed CA can never return to the pending state. You must create a new CA.
//
// - DELETED - Your private CA is within the restoration period, after which it
// is permanently deleted. The length of time remaining in the CA's restoration
// period is also included in this action's output.
