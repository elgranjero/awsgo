package gamelift

// CreateGameServerGroup is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2 (FleetIQ)
//
// Creates a Amazon GameLift Servers FleetIQ game server group for managing game
// hosting on a collection of Amazon Elastic Compute Cloud instances for game
// hosting. This operation creates the game server group, creates an Auto Scaling
// group in your Amazon Web Services account, and establishes a link between the
// two groups. You can view the status of your game server groups in the Amazon
// GameLift Servers console. Game server group metrics and events are emitted to
// Amazon CloudWatch.
//
// Before creating a new game server group, you must have the following:
//
// - An Amazon Elastic Compute Cloud launch template that specifies how to
// launch Amazon Elastic Compute Cloud instances with your game server build. For
// more information, see [Launching an Instance from a Launch Template]in the Amazon Elastic Compute Cloud User Guide.
//
// - An IAM role that extends limited access to your Amazon Web Services account
// to allow Amazon GameLift Servers FleetIQ to create and interact with the Auto
// Scaling group. For more information, see [Create IAM roles for cross-service interaction]in the Amazon GameLift Servers
// FleetIQ Developer Guide.
//
// To create a new game server group, specify a unique group name, IAM role and
// Amazon Elastic Compute Cloud launch template, and provide a list of instance
// types that can be used in the group. You must also set initial maximum and
// minimum limits on the group's instance count. You can optionally set an Auto
// Scaling policy with target tracking based on a Amazon GameLift Servers FleetIQ
// metric.
//
// Once the game server group and corresponding Auto Scaling group are created,
// you have full access to change the Auto Scaling group's configuration as needed.
// Several properties that are set when creating a game server group, including
// maximum/minimum size and auto-scaling policy settings, must be updated directly
// in the Auto Scaling group. Keep in mind that some Auto Scaling group properties
// are periodically updated by Amazon GameLift Servers FleetIQ as part of its
// balancing activities to optimize for availability and cost.
//
// # Learn more
//
// [Amazon GameLift Servers FleetIQ Guide]
//
// [Amazon GameLift Servers FleetIQ Guide]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-intro.html
// [Create IAM roles for cross-service interaction]: https://docs.aws.amazon.com/gamelift/latest/fleetiqguide/gsg-iam-permissions-roles.html
// [Launching an Instance from a Launch Template]: https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-launch-templates.html
