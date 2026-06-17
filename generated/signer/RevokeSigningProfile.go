package signer

// RevokeSigningProfile is generated as a reference stub.
// Executable command wiring lives under cmd/signer.go.
//
// Changes the state of a signing profile to REVOKED . This indicates that
// signatures generated using the signing profile after an effective start date are
// no longer valid. A revoked profile is still viewable with the
// ListSigningProfiles operation, but it cannot perform new signing jobs. See [Data Retention] for
// more information on scheduled deletion of a revoked signing profile.
//
// [Data Retention]: https://docs.aws.amazon.com/signer/latest/developerguide/retention.html
