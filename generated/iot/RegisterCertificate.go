package iot

// RegisterCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Registers a device certificate with IoT in the same [certificate mode] as the signing CA. If you
// have more than one CA certificate that has the same subject field, you must
// specify the CA certificate that was used to sign the device certificate being
// registered.
//
// Requires permission to access the [RegisterCertificate] action.
//
// [RegisterCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [certificate mode]: https://docs.aws.amazon.com/iot/latest/apireference/API_CertificateDescription.html#iot-Type-CertificateDescription-certificateMode
