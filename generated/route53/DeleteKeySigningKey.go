package route53

// DeleteKeySigningKey is generated as a reference stub.
// Executable command wiring lives under cmd/route53.go.
//
// Deletes a key-signing key (KSK). Before you can delete a KSK, you must
// deactivate it. The KSK must be deactivated before you can delete it regardless
// of whether the hosted zone is enabled for DNSSEC signing.
//
// You can use [DeactivateKeySigningKey] to deactivate the key before you delete it.
//
// Use [GetDNSSEC] to verify that the KSK is in an INACTIVE status.
//
// [DeactivateKeySigningKey]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_DeactivateKeySigningKey.html
// [GetDNSSEC]: https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetDNSSEC.html
