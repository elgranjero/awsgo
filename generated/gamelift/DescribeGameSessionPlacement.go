package gamelift

// DescribeGameSessionPlacement is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Anywhere, Container
//
// Retrieves information, including current status, about a game session placement
// request.
//
// To get game session placement details, specify the placement ID.
//
// This operation is not designed to be continually called to track game session
// status. This practice can cause you to exceed your API limit, which results in
// errors. Instead, you must configure an Amazon Simple Notification Service (SNS)
// topic to receive notifications from FlexMatch or queues. Continuously polling
// with DescribeGameSessionPlacement should only be used for games in development
// with low game session usage.
