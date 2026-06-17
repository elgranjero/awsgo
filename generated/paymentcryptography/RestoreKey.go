package paymentcryptography

// RestoreKey is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Cancels a scheduled key deletion during the waiting period. Use this operation
// to restore a Key that is scheduled for deletion.
//
// During the waiting period, the KeyState is DELETE_PENDING and
// deletePendingTimestamp contains the date and time after which the Key will be
// deleted. After Key is restored, the KeyState is CREATE_COMPLETE , and the value
// for deletePendingTimestamp is removed.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteKey]
//
// [StartKeyUsage]
//
// [StopKeyUsage]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [StartKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StartKeyUsage.html
// [StopKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StopKeyUsage.html
