package paymentcryptography

// GetParametersForImport is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Gets the import token and the wrapping key certificate in PEM format (base64
// encoded) to initiate a TR-34 WrappedKeyBlock or a RSA WrappedKeyCryptogram
// import into Amazon Web Services Payment Cryptography.
//
// The wrapping key certificate wraps the key under import. The import token and
// wrapping key certificate must be in place and operational before calling [ImportKey]. The
// import token expires in 30 days. You can use the same import token to import
// multiple keys into your service account.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [GetParametersForExport]
//
// [ImportKey]
//
// [GetParametersForExport]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetParametersForExport.html
// [ImportKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ImportKey.html
