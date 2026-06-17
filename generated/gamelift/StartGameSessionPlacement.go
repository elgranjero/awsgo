package gamelift

// StartGameSessionPlacement is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Makes a request to start a new game session using a game session queue. When
// processing a placement request, Amazon GameLift Servers looks for the best
// possible available resource to host the game session, based on how the queue is
// configured to prioritize factors such as resource cost, latency, and location.
// After selecting an available resource, Amazon GameLift Servers prompts the
// resource to start a game session. A placement request can include a list of
// players to create a set of player sessions. The request can also include
// information to pass to the new game session, such as to specify a game map or
// other options.
//
// # Request options
//
// Use this operation to make the following types of requests.
//
// - Request a placement using the queue's default prioritization process (see
// the default prioritization described in [PriorityConfiguration]). Include these required parameters:
//
// - GameSessionQueueName
//
// - MaximumPlayerSessionCount
//
// - PlacementID
//
// - Request a placement and prioritize based on latency. Include these
// parameters:
//
// - Required parameters GameSessionQueueName , MaximumPlayerSessionCount ,
// PlacementID .
//
// - PlayerLatencies . Include a set of latency values for destinations in the
// queue. When a request includes latency data, Amazon GameLift Servers
// automatically reorder the queue's locations priority list based on lowest
// available latency values. If a request includes latency data for multiple
// players, Amazon GameLift Servers calculates each location's average latency for
// all players and reorders to find the lowest latency across all players.
//
// - Don't include PriorityConfigurationOverride .
//
// - Prioritize based on a custom list of locations. If you're using a queue
// that's configured to prioritize location first (see [PriorityConfiguration]for game session queues),
// you can optionally use the PriorityConfigurationOverride parameter to substitute
// a different location priority list for this placement request. Amazon GameLift
// Servers searches each location on the priority override list to find an
// available hosting resource for the new game session. Specify a fallback strategy
// to use in the event that Amazon GameLift Servers fails to place the game session
// in any of the locations on the override list.
//
// - Request a placement and prioritized based on a custom list of locations.
//
// - You can request new player sessions for a group of players. Include the
// DesiredPlayerSessions parameter and include at minimum a unique player ID for
// each. You can also include player-specific data to pass to the new game session.
//
// # Result
//
// If successful, this operation generates a new game session placement request
// and adds it to the game session queue for processing. You can track the status
// of individual placement requests by calling [DescribeGameSessionPlacement]or by monitoring queue
// notifications. When the request status is FULFILLED , a new game session has
// started and the placement request is updated with connection information for the
// game session (IP address and port). If the request included player session data,
// Amazon GameLift Servers creates a player session for each player ID in the
// request.
//
// The request results in a InvalidRequestException in the following situations:
//
// - If the request includes both PlayerLatencies and
// PriorityConfigurationOverride parameters.
//
// - If the request includes the PriorityConfigurationOverride parameter and
// specifies a queue that doesn't prioritize locations.
//
// Amazon GameLift Servers continues to retry each placement request until it
// reaches the queue's timeout setting. If a request times out, you can resubmit
// the request to the same queue or try a different queue.
//
// [DescribeGameSessionPlacement]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeGameSessionPlacement.html
// [PriorityConfiguration]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_PriorityConfiguration.html
