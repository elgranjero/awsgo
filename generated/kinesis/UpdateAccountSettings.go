package kinesis

// UpdateAccountSettings is generated as a reference stub.
// Executable command wiring lives under cmd/kinesis.go.
//
// Updates the account-level settings for Amazon Kinesis Data Streams.
//
// Updating account settings is a synchronous operation. Upon receiving the
// request, Kinesis Data Streams will return immediately with your account’s
// updated settings.
//
// API limits
//
// - Certain account configurations have minimum commitment windows. Attempting
// to update your settings prior to the end of the minimum commitment window might
// have certain restrictions.
//
// - This API has a call limit of 5 transactions per second (TPS) for each
// Amazon Web Services account. TPS over 5 will initiate the
// LimitExceededException .
