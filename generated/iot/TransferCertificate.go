package iot

// TransferCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Transfers the specified certificate to the specified Amazon Web Services
// account.
//
// Requires permission to access the [TransferCertificate] action.
//
// You can cancel the transfer until it is accepted by the recipient.
//
// No notification is sent to the transfer destination's account. The caller is
// responsible for notifying the transfer target.
//
// The certificate being transferred must not be in the ACTIVE state. You can use
// the UpdateCertificateaction to deactivate it.
//
// The certificate must not have any policies attached to it. You can use the DetachPolicy
// action to detach them.
//
// Customer managed key behavior: When you use a customer managed key to encrypt
// your data and then transfer the certificate to a customer in a different account
// using the TransferCertificate operation, the certificates will no longer be
// encrypted by their customer managed key configuration. During the transfer
// process, certificates are encrypted using Amazon Web Services IoT Core owned
// keys.
//
// While a certificate is in the PENDING_TRANSFER state, it's always protected by
// Amazon Web Services IoT Core owned keys, regardless of the customer managed key
// configuration of either the source or destination account.
//
// Once the transfer is completed through AcceptCertificateTransfer, RejectCertificateTransfer, or CancelCertificateTransfer, the certificate will be
// protected by the customer managed key configuration of the account that owns the
// certificate after the transfer operation:
//
// - If the transfer is accepted: The certificate is encrypted by the target
// account's customer managed key configuration.
//
// - If the transfer is rejected or cancelled: The certificate is protected by
// the source account's customer managed key configuration.
//
// [TransferCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
