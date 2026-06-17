package ses

// ListConfigurationSets is generated as a reference stub.
// Executable command wiring lives under cmd/ses.go.
//
// Provides a list of the configuration sets associated with your Amazon SES
// account in the current Amazon Web Services Region. For information about using
// configuration sets, see [Monitoring Your Amazon SES Sending Activity]in the Amazon SES Developer Guide.
//
// You can execute this operation no more than once per second. This operation
// returns up to 1,000 configuration sets each time it is run. If your Amazon SES
// account has more than 1,000 configuration sets, this operation also returns
// NextToken . You can then execute the ListConfigurationSets operation again,
// passing the NextToken parameter and the value of the NextToken element to
// retrieve additional results.
//
// [Monitoring Your Amazon SES Sending Activity]: https://docs.aws.amazon.com/ses/latest/dg/monitor-sending-activity.html
