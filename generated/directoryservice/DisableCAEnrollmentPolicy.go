package directoryservice

// DisableCAEnrollmentPolicy is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Disables the certificate authority (CA) enrollment policy for the specified
// directory. This stops automatic certificate enrollment and management for
// domain-joined clients, but does not affect existing certificates.
//
// Disabling the CA enrollment policy prevents new certificates from being
// automatically enrolled, but existing certificates remain valid and functional
// until they expire.
