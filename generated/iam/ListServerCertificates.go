package iam

// ListServerCertificates is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Lists the server certificates stored in IAM that have the specified path
// prefix. If none exist, the operation returns an empty list.
//
// You can paginate the results using the MaxItems and Marker parameters.
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// IAM resource-listing operations return a subset of the available attributes for
// the resource. For example, this operation does not return tags, even though they
// are an attribute of the returned object. To view all of the information for a
// servercertificate, see [GetServerCertificate].
//
// [GetServerCertificate]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetServerCertificate.html
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
