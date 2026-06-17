package transfer

// ImportCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Imports the signing and encryption certificates that you need to create local
// (AS2) profiles and partner profiles.
//
// You can import both the certificate and its chain in the Certificate parameter.
//
// After importing a certificate, Transfer Family automatically creates a Amazon
// CloudWatch metric called DaysUntilExpiry that tracks the number of days until
// the certificate expires. The metric is based on the InactiveDate parameter and
// is published daily in the AWS/Transfer namespace.
//
// It can take up to a full day after importing a certificate for Transfer Family
// to emit the DaysUntilExpiry metric to your account.
//
// If you use the Certificate parameter to upload both the certificate and its
// chain, don't use the CertificateChain parameter.
//
// # CloudWatch monitoring
//
// The DaysUntilExpiry metric includes the following specifications:
//
// - Units: Count (days)
//
// - Dimensions: CertificateId (always present), Description (if provided during
// certificate import)
//
// - Statistics: Minimum, Maximum, Average
//
// - Frequency: Published daily
