package gamelift

// CreateAlias is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Creates an alias for a fleet. In most situations, you can use an alias ID in
// place of a fleet ID. An alias provides a level of abstraction for a fleet that
// is useful when redirecting player traffic from one fleet to another, such as
// when updating your game build.
//
// Amazon GameLift Servers supports two types of routing strategies for aliases:
// simple and terminal. A simple alias points to an active fleet. A terminal alias
// is used to display messaging or link to a URL instead of routing players to an
// active fleet. For example, you might use a terminal alias when a game version is
// no longer supported and you want to direct players to an upgrade site.
//
// To create a fleet alias, specify an alias name, routing strategy, and optional
// description. Each simple alias can point to only one fleet, but a fleet can have
// multiple aliases. If successful, a new alias record is returned, including an
// alias ID and an ARN. You can reassign an alias to another fleet by calling
// UpdateAlias .
//
// # Related actions
//
// [All APIs by task]
//
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
