package paymentcryptography

// CreateAlias is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Creates an alias, or a friendly name, for an Amazon Web Services Payment
// Cryptography key. You can use an alias to identify a key in the console and when
// you call cryptographic operations such as [EncryptData]or [DecryptData].
//
// You can associate the alias with any key in the same Amazon Web Services
// Region. Each alias is associated with only one key at a time, but a key can have
// multiple aliases. You can't create an alias without a key. The alias must be
// unique in the account and Amazon Web Services Region, but you can create another
// alias with the same name in a different Amazon Web Services Region.
//
// To change the key that's associated with the alias, call [UpdateAlias]. To delete the alias,
// call [DeleteAlias]. These operations don't affect the underlying key. To get the alias that
// you created, call [ListAliases].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [DeleteAlias]
//
// [GetAlias]
//
// [ListAliases]
//
// [UpdateAlias]
//
// [ListAliases]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_ListAliases.html
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
// [EncryptData]: https://docs.aws.amazon.com/payment-cryptography/latest/DataAPIReference/API_EncryptData.html
// [DecryptData]: https://docs.aws.amazon.com/payment-cryptography/latest/DataAPIReference/API_DecryptData.html
