package iot

// CreateCertificateProvider is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Creates an Amazon Web Services IoT Core certificate provider. You can use
// Amazon Web Services IoT Core certificate provider to customize how to sign a
// certificate signing request (CSR) in IoT fleet provisioning. For more
// information, see [Customizing certificate signing using Amazon Web Services IoT Core certificate provider]from Amazon Web Services IoT Core Developer Guide.
//
// Requires permission to access the [CreateCertificateProvider] action.
//
// After you create a certificate provider, the behavior of [CreateCertificateFromCsr API for fleet provisioning]
// CreateCertificateFromCsr will change and all API calls to
// CreateCertificateFromCsr will invoke the certificate provider to create the
// certificates. It can take up to a few minutes for this behavior to change after
// a certificate provider is created.
//
// [CreateCertificateProvider]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [CreateCertificateFromCsr API for fleet provisioning]: https://docs.aws.amazon.com/iot/latest/developerguide/fleet-provision-api.html#create-cert-csr
// [Customizing certificate signing using Amazon Web Services IoT Core certificate provider]: https://docs.aws.amazon.com/iot/latest/developerguide/provisioning-cert-provider.html
