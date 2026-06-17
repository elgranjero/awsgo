package paymentcryptography

// DeleteAlias is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Deletes the alias, but doesn't affect the underlying key.
//
// Each key can have multiple aliases. To get the aliases of all keys, use the [UpdateAlias]
// operation. To change the alias of a key, first use [DeleteAlias]to delete the current alias
// and then use [CreateAlias]to create a new alias. To associate an existing alias with a
// different key, call [UpdateAlias].
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
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
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
