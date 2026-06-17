package redshift

// CreateHsmClientCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Creates an HSM client certificate that an Amazon Redshift cluster will use to
// connect to the client's HSM in order to store and retrieve the keys used to
// encrypt the cluster databases.
//
// The command returns a public key, which you must store in the HSM. In addition
// to creating the HSM certificate, you must create an Amazon Redshift HSM
// configuration that provides a cluster the information needed to store and use
// encryption keys in the HSM. For more information, go to [Hardware Security Modules]in the Amazon Redshift
// Cluster Management Guide.
//
// [Hardware Security Modules]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-db-encryption.html#working-with-HSM
