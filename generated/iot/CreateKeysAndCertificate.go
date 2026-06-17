package iot

// CreateKeysAndCertificate is generated as a reference stub.
// Executable command wiring lives under cmd/iot.go.
//
// Creates a 2048-bit RSA key pair and issues an X.509 certificate using the
// issued public key. You can also call CreateKeysAndCertificate over MQTT from a
// device, for more information, see [Provisioning MQTT API].
//
// Note This is the only time IoT issues the private key for this certificate, so
// it is important to keep it in a secure location.
//
// Requires permission to access the [CreateKeysAndCertificate] action.
//
// [CreateKeysAndCertificate]: https://docs.aws.amazon.com/service-authorization/latest/reference/list_awsiot.html#awsiot-actions-as-permissions
// [Provisioning MQTT API]: https://docs.aws.amazon.com/iot/latest/developerguide/provision-wo-cert.html#provision-mqtt-api
