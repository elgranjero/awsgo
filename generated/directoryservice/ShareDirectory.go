package directoryservice

// ShareDirectory is generated as a reference stub.
// Executable command wiring lives under cmd/directoryservice.go.
//
// Shares a specified directory ( DirectoryId ) in your Amazon Web Services account
// (directory owner) with another Amazon Web Services account (directory consumer).
// With this operation you can use your directory from any Amazon Web Services
// account and from any Amazon VPC within an Amazon Web Services Region.
//
// When you share your Managed Microsoft AD directory, Directory Service creates a
// shared directory in the directory consumer account. This shared directory
// contains the metadata to provide access to the directory within the directory
// owner account. The shared directory is visible in all VPCs in the directory
// consumer account.
//
// The ShareMethod parameter determines whether the specified directory can be
// shared between Amazon Web Services accounts inside the same Amazon Web Services
// organization ( ORGANIZATIONS ). It also determines whether you can share the
// directory with any other Amazon Web Services account either inside or outside of
// the organization ( HANDSHAKE ).
//
// The ShareNotes parameter is only used when HANDSHAKE is called, which sends a
// directory sharing request to the directory consumer.
