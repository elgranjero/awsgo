package paymentcryptography

// ListKeys is generated as a reference stub.
// Executable command wiring lives under cmd/paymentcryptography.go.
//
// Lists the keys in the caller's Amazon Web Services account and Amazon Web
// Services Region. You can filter the list of keys.
//
// This is a paginated operation, which means that each response might contain
// only a subset of all the keys. When the response contains only a subset of keys,
// it includes a NextToken value. Use this value in a subsequent ListKeys request
// to get more keys. When you receive a response with no NextToken (or an empty or
// null value), that means there are no more keys to get.
//
// Cross-account use: This operation can't be used across different Amazon Web
// Services accounts.
//
// Related operations:
//
// [CreateKey]
//
// [DeleteKey]
//
// [GetKey]
//
// [DeleteKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_DeleteKey.html
// [GetKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_GetKey.html
// [CreateKey]: https://docs.aws.amazon.com/payment-cryptography/latest/APIReference/API_CreateKey.html
