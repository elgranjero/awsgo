package cloudtrail

// ListPublicKeys is generated as a reference stub.
// Executable command wiring lives under cmd/cloudtrail.go.
//
// Returns all public keys whose private keys were used to sign the digest files
// within the specified time range. The public key is needed to validate digest
// files that were signed with its corresponding private key.
//
// CloudTrail uses different private and public key pairs per Region. Each digest
// file is signed with a private key unique to its Region. When you validate a
// digest file from a specific Region, you must look in the same Region for its
// corresponding public key.
