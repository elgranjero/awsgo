package iam

// DeleteServerCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Deletes the specified server certificate.
//
// For more information about working with server certificates, see [Working with server certificates] in the IAM
// User Guide. This topic also includes a list of Amazon Web Services services that
// can use the server certificates that you manage with IAM.
//
// If you are using a server certificate with Elastic Load Balancing, deleting the
// certificate could have implications for your application. If Elastic Load
// Balancing doesn't detect the deletion of bound certificates, it may continue to
// use the certificates. This could cause Elastic Load Balancing to stop accepting
// traffic. We recommend that you remove the reference to the certificate from
// Elastic Load Balancing before using this command to delete the certificate. For
// more information, see [DeleteLoadBalancerListeners]in the Elastic Load Balancing API Reference.
//
// [Working with server certificates]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html
// [DeleteLoadBalancerListeners]: https://docs.aws.amazon.com/ElasticLoadBalancing/latest/APIReference/API_DeleteLoadBalancerListeners.html
