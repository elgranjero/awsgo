package gamelift

// CreateMatchmakingRuleSet is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Creates a new rule set for FlexMatch matchmaking. A rule set describes the type
// of match to create, such as the number and size of teams. It also sets the
// parameters for acceptable player matches, such as minimum skill level or
// character type.
//
// To create a matchmaking rule set, provide unique rule set name and the rule set
// body in JSON format. Rule sets must be defined in the same Region as the
// matchmaking configuration they are used with.
//
// Since matchmaking rule sets cannot be edited, it is a good idea to check the
// rule set syntax using [ValidateMatchmakingRuleSet]before creating a new rule set.
//
// # Learn more
//
// [Build a rule set]
//
// [Design a matchmaker]
//
// [Matchmaking with FlexMatch]
//
// [Matchmaking with FlexMatch]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-intro.html
// [Build a rule set]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html
// [Design a matchmaker]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html
// [ValidateMatchmakingRuleSet]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_ValidateMatchmakingRuleSet.html
