package gamelift

// DescribeMatchmakingConfigurations is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves the details of FlexMatch matchmaking configurations.
//
// This operation offers the following options: (1) retrieve all matchmaking
// configurations, (2) retrieve configurations for a specified list, or (3)
// retrieve all configurations that use a specified rule set name. When requesting
// multiple items, use the pagination parameters to retrieve results as a set of
// sequential pages.
//
// If successful, a configuration is returned for each requested name. When
// specifying a list of names, only configurations that currently exist are
// returned.
//
// # Learn more
//
// [Setting up FlexMatch matchmakers]
//
// [Setting up FlexMatch matchmakers]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/matchmaker-build.html
