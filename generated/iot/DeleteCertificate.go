package iot

// DeleteCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Deletes the specified certificate.
//
// A certificate cannot be deleted if it has a policy or IoT thing attached to it
// or if its status is set to ACTIVE. To delete a certificate, first use the DetachPolicy
// action to detach all policies. Next, use the UpdateCertificateaction to set the certificate to
// the INACTIVE status.
//
// Requires permission to access the [DeleteCertificate] action.
//
// [DeleteCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
