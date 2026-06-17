package appconfigdata

// GetLatestConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/appconfigdata.go.
//
// Retrieves the latest deployed configuration. This API may return empty
// configuration data if the client already has the latest version. For more
// information about this API action and to view example CLI commands that show how
// to use it with the StartConfigurationSessionAPI action, see [Retrieving the configuration] in the AppConfig User Guide.
//
// Note the following important information.
//
// - Each configuration token is only valid for one call to
// GetLatestConfiguration . The GetLatestConfiguration response includes a
// NextPollConfigurationToken that should always replace the token used for the
// just-completed call in preparation for the next one.
//
// - GetLatestConfiguration is a priced call. For more information, see [Pricing].
//
// [Pricing]: https://aws.amazon.com/systems-manager/pricing/
// [Retrieving the configuration]: http://docs.aws.amazon.com/appconfig/latest/userguide/appconfig-retrieving-the-configuration
