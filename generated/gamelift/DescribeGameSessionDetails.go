package gamelift

// DescribeGameSessionDetails is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves additional game session properties, including the game session
// protection policy in force, a set of one or more game sessions in a specific
// fleet location. You can optionally filter the results by current game session
// status.
//
// This operation can be used in the following ways:
//
// - To retrieve details for all game sessions that are currently running on all
// locations in a fleet, provide a fleet or alias ID, with an optional status
// filter. This approach returns details from the fleet's home Region and all
// remote locations.
//
// - To retrieve details for all game sessions that are currently running on a
// specific fleet location, provide a fleet or alias ID and a location name, with
// optional status filter. The location can be the fleet's home Region or any
// remote location.
//
// - To retrieve details for a specific game session, provide the game session
// ID. This approach looks for the game session ID in all fleets that reside in the
// Amazon Web Services Region defined in the request.
//
// Use the pagination parameters to retrieve results as a set of sequential pages.
//
// If successful, a GameSessionDetail object is returned for each game session
// that matches the request.
//
// # Learn more
//
// [Find a game session]
//
// [All APIs by task]
//
// [Find a game session]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-sdk-client-api.html#gamelift-sdk-client-api-find
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
