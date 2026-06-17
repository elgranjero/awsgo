package iot

// CancelCertificateTransfer is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Cancels a pending transfer for the specified certificate.
//
// Note Only the transfer source account can use this operation to cancel a
// transfer. (Transfer destinations can use RejectCertificateTransferinstead.) After transfer, IoT returns
// the certificate to the source account in the INACTIVE state. After the
// destination account has accepted the transfer, the transfer cannot be cancelled.
//
// After a certificate transfer is cancelled, the status of the certificate
// changes from PENDING_TRANSFER to INACTIVE.
//
// Requires permission to access the [CancelCertificateTransfer] action.
//
// [CancelCertificateTransfer]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
