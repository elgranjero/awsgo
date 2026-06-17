package gamelift

// CreateMatchmakingConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Defines a new matchmaking configuration for use with FlexMatch. Whether your
// are using FlexMatch with Amazon GameLift Servers hosting or as a standalone
// matchmaking service, the matchmaking configuration sets out rules for matching
// players and forming teams. If you're also using Amazon GameLift Servers hosting,
// it defines how to start game sessions for each match. Your matchmaking system
// can use multiple configurations to handle different game scenarios. All
// matchmaking requests identify the matchmaking configuration to use and provide
// player attributes consistent with that configuration.
//
// To create a matchmaking configuration, you must provide the following:
// configuration name and FlexMatch mode (with or without Amazon GameLift Servers
// hosting); a rule set that specifies how to evaluate players and find acceptable
// matches; whether player acceptance is required; and the maximum time allowed for
// a matchmaking attempt. When using FlexMatch with Amazon GameLift Servers
// hosting, you also need to identify the game session queue to use when starting a
// game session for the match.
//
// In addition, you must set up an Amazon Simple Notification Service topic to
// receive matchmaking notifications. Provide the topic ARN in the matchmaking
// configuration.
//
// # Learn more
//
// [Design a FlexMatch matchmaker]
//
// [Set up FlexMatch event notification]
//
// [Design a FlexMatch matchmaker]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-configuration.html
// [Set up FlexMatch event notification]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html
