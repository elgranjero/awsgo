package redshift

// DescribeHsmClientCertificates is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Returns information about the specified HSM client certificate. If no
// certificate ID is specified, returns information about all the HSM certificates
// owned by your Amazon Web Services account.
//
// If you specify both tag keys and tag values in the same request, Amazon
// Redshift returns all HSM client certificates that match any combination of the
// specified keys and values. For example, if you have owner and environment for
// tag keys, and admin and test for tag values, all HSM client certificates that
// have any combination of those values are returned.
//
// If both tag keys and values are omitted from the request, HSM client
// certificates are returned regardless of whether they have tag keys or values
// associated with them.
