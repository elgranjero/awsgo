package paymentcryptography

// GetParametersForExport is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Gets the export token and the signing key certificate to initiate a TR-34 key
// export from Amazon Web Services Payment Cryptography.
//
// The signing key certificate signs the wrapped key under export within the TR-34
// key payload. The export token and signing key certificate must be in place and
// operational before calling [ExportKey]. The export token expires in 30 days. You can use
// the same export token to export multiple keys from your service account.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [ExportKey]
//
// [GetParametersForImport]
//
// [ExportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ExportKey.html
// [GetParametersForImport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForImport.html
