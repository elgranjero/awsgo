package transfer

// CreateAgreement is generated as a reference stub.
// Executable command wiring lives under cmd/transfer.go.
//
// Creates an agreement. An agreement is a bilateral trading partner agreement, or
// partnership, between an Transfer Family server and an AS2 process. The agreement
// defines the file and message transfer relationship between the server and the
// AS2 process. To define an agreement, Transfer Family combines a server, local
// profile, partner profile, certificate, and other attributes.
//
// The partner is identified with the PartnerProfileId , and the AS2 process is
// identified with the LocalProfileId .
//
// Specify either BaseDirectory or CustomDirectories , but not both. Specifying
// both causes the command to fail.
