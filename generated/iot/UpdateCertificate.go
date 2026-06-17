package iot

// UpdateCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Updates the status of the specified certificate. This operation is idempotent.
//
// Requires permission to access the [UpdateCertificate] action.
//
// Certificates must be in the ACTIVE state to authenticate devices that use a
// certificate to connect to IoT.
//
// Within a few minutes of updating a certificate from the ACTIVE state to any
// other state, IoT disconnects all devices that used that certificate to connect.
// Devices cannot use a certificate that is not in the ACTIVE state to reconnect.
//
// [UpdateCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
