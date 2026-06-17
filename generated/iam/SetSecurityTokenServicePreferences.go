package iam

// SetSecurityTokenServicePreferences is generated as a reference stub.
// Executable command wiring lives under cmd/iam.go.
//
// Sets the specified version of the global endpoint token as the token version
// used for the Amazon Web Services account.
//
// By default, Security Token Service (STS) is available as a global service, and
// all STS requests go to a single endpoint at https://sts.amazonaws.com . Amazon
// Web Services recommends using Regional STS endpoints to reduce latency, build in
// redundancy, and increase session token availability. For information about
// Regional endpoints for STS, see [Security Token Service endpoints and quotas]in the Amazon Web Services General Reference.
//
// If you make an STS call to the global endpoint, the resulting session tokens
// might be valid in some Regions but not others. It depends on the version that is
// set in this operation. Version 1 tokens are valid only in Amazon Web Services
// Regions that are available by default. These tokens do not work in manually
// enabled Regions, such as Asia Pacific (Hong Kong). Version 2 tokens are valid in
// all Regions. However, version 2 tokens are longer and might affect systems where
// you temporarily store tokens. For information, see [Activating and deactivating STS in an Amazon Web Services Region]in the IAM User Guide.
//
// To view the current session token version, see the GlobalEndpointTokenVersion
// entry in the response of the [GetAccountSummary]operation.
//
// [Activating and deactivating STS in an Amazon Web Services Region]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_enable-regions.html
// [GetAccountSummary]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_GetAccountSummary.html
// [Security Token Service endpoints and quotas]: https://docs.aws.amazon.com/general/latest/gr/sts.html
