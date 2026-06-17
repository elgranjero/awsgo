package paymentcryptography

// DeleteKey is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Deletes the key material and metadata associated with Amazon Web Services
// Payment Cryptography key.
//
// Key deletion is irreversible. After a key is deleted, you can't perform
// cryptographic operations using the key. For example, you can't decrypt data that
// was encrypted by a deleted Amazon Web Services Payment Cryptography key, and the
// data may become unrecoverable. Because key deletion is destructive, Amazon Web
// Services Payment Cryptography has a safety mechanism to prevent accidental
// deletion of a key. When you call this operation, Amazon Web Services Payment
// Cryptography disables the specified key but doesn't delete it until after a
// waiting period set using DeleteKeyInDays . The default waiting period is 7 days.
// During the waiting period, the KeyState is DELETE_PENDING . After the key is
// deleted, the KeyState is DELETE_COMPLETE .
//
// You should delete a key only when you are sure that you don't need to use it
// anymore and no other parties are utilizing this key. If you aren't sure,
// consider deactivating it instead by calling [StopKeyUsage].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [RestoreKey]
//
// [StartKeyUsage]
//
// [StopKeyUsage]
//
// [StartKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StartKeyUsage.html
// [StopKeyUsage]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_StopKeyUsage.html
// [RestoreKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_RestoreKey.html
