package securitylake

// CreateDataLake is generated as a reference stub.
// Executable command wiring lives under cmd/securitylake.go.
//
// Initializes an Amazon Security Lake instance with the provided (or default)
// configuration. You can enable Security Lake in Amazon Web Services Regions with
// customized settings before enabling log collection in Regions. To specify
// particular Regions, configure these Regions using the configurations parameter.
// If you have already enabled Security Lake in a Region when you call this
// command, the command will update the Region if you provide new configuration
// parameters. If you have not already enabled Security Lake in the Region when you
// call this API, it will set up the data lake in the Region with the specified
// configurations.
//
// When you enable Security Lake, it starts ingesting security data after the
// CreateAwsLogSource call and after you create subscribers using the
// CreateSubscriber API. This includes ingesting security data from sources,
// storing data, and making data accessible to subscribers. Security Lake also
// enables all the existing settings and resources that it stores or maintains for
// your Amazon Web Services account in the current Region, including security log
// and event data. For more information, see the [Amazon Security Lake User Guide].
//
// [Amazon Security Lake User Guide]: https://docs.aws.amazon.com/security-lake/latest/userguide/what-is-security-lake.html
