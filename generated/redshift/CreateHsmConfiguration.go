package redshift

// CreateHsmConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/redshift.go.
//
// Creates an HSM configuration that contains the information required by an
// Amazon Redshift cluster to store and use database encryption keys in a Hardware
// Security Module (HSM). After creating the HSM configuration, you can specify it
// as a parameter when creating a cluster. The cluster will then store its
// encryption keys in the HSM.
//
// In addition to creating an HSM configuration, you must also create an HSM
// client certificate. For more information, go to [Hardware Security Modules]in the Amazon Redshift Cluster
// Management Guide.
//
// [Hardware Security Modules]: https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-HSM.html
