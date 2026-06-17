package paymentcryptography

// ListAliases is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Lists the aliases for all keys in the caller's Amazon Web Services account and
// Amazon Web Services Region. You can filter the aliases by keyARN . For more
// information, see [Using aliases]in the Amazon Web Services Payment Cryptography User Guide.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the aliases. When the response contains only a subset of
// aliases, it includes a NextToken value. Use this value in a subsequent
// ListAliases request to get more aliases. When you receive a response with no
// NextToken (or an empty or null value), that means there are no more aliases to
// get.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateAlias]
//
// [DeleteAlias]
//
// [GetAlias]
//
// [UpdateAlias]
//
// [DeleteAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteAlias.html
// [UpdateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_UpdateAlias.html
// [Using aliases]: https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-managealias.html
// [CreateAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateAlias.html
// [GetAlias]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetAlias.html
