package gamelift

// DescribeMatchmaking is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves one or more matchmaking tickets. Use this operation to retrieve
// ticket information, including--after a successful match is made--connection
// information for the resulting new game session.
//
// To request matchmaking tickets, provide a list of up to 10 ticket IDs. If the
// request is successful, a ticket object is returned for each requested ID that
// currently exists.
//
// This operation is not designed to be continually called to track matchmaking
// ticket status. This practice can cause you to exceed your API limit, which
// results in errors. Instead, as a best practice, set up an Amazon Simple
// Notification Service to receive notifications, and provide the topic ARN in the
// matchmaking configuration.
//
// # Learn more
//
// [Add FlexMatch to a game client]
//
// [Set Up FlexMatch event notification]
//
// [Set Up FlexMatch event notification]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-notification.html
// [Add FlexMatch to a game client]: https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-client.html
