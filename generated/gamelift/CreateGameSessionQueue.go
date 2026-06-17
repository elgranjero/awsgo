package gamelift

// CreateGameSessionQueue is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Creates a placement queue that processes requests for new game sessions. A
// queue uses FleetIQ algorithms to locate the best available placement locations
// for a new game session, and then prompts the game server process to start a new
// game session.
//
// A game session queue is configured with a set of destinations (Amazon GameLift
// Servers fleets or aliases) that determine where the queue can place new game
// sessions. These destinations can span multiple Amazon Web Services Regions, can
// use different instance types, and can include both Spot and On-Demand fleets. If
// the queue includes multi-location fleets, the queue can place game sessions in
// any of a fleet's remote locations.
//
// You can configure a queue to determine how it selects the best available
// placement for a new game session. Queues can prioritize placement decisions
// based on a combination of location, hosting cost, and player latency. You can
// set up the queue to use the default prioritization or provide alternate
// instructions using PriorityConfiguration .
//
// # Request options
//
// Use this operation to make these common types of requests.
//
// - Create a queue with the minimum required parameters.
//
// - Name
//
// - Destinations (This parameter isn't required, but a queue can't make
// placements without at least one destination.)
//
// - Create a queue with placement notification. Queues that have high placement
// activity must use a notification system, such as with Amazon Simple Notification
// Service (Amazon SNS) or Amazon CloudWatch.
//
// - Required parameters Name and Destinations
//
// - NotificationTarget
//
// - Create a queue with custom prioritization settings. These custom settings
// replace the default prioritization configuration for a queue.
//
// - Required parameters Name and Destinations
//
// - PriorityConfiguration
//
// - Create a queue with special rules for processing player latency data.
//
// - Required parameters Name and Destinations
//
// - PlayerLatencyPolicies
//
// # Results
//
// If successful, this operation returns a new GameSessionQueue object with an
// assigned queue ARN. Use the queue's name or ARN when submitting new game session
// requests with [StartGameSessionPlacement]or [StartMatchmaking].
//
// # Learn more
//
// [Design a game session queue]
//
// [Create a game session queue]
//
// # Related actions
//
// [CreateGameSessionQueue]| [DescribeGameSessionQueues] | [UpdateGameSessionQueue] | [DeleteGameSessionQueue] | [All APIs by task]
//
// [StartMatchmaking]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StartMatchmaking.html
// [DescribeGameSessionQueues]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeGameSessionQueues.html
// [StartGameSessionPlacement]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_StartGameSessionPlacement.html
// [Create a game session queue]: https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-creating.html
// [CreateGameSessionQueue]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateGameSessionQueue.html
// [UpdateGameSessionQueue]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_UpdateGameSessionQueue.html
// [Design a game session queue]: https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-design.html
// [DeleteGameSessionQueue]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DeleteGameSessionQueue.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
