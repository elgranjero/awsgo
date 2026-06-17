package gamelift

// CreateFleetLocations is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2, Container
//
// Adds remote locations to an EC2 and begins populating the new locations with
// instances. The new instances conform to the fleet's instance type, auto-scaling,
// and other configuration settings.
//
// You can't add remote locations to a fleet that resides in an Amazon Web
// Services Region that doesn't support multiple locations. Fleets created prior to
// March 2021 can't support multiple locations.
//
// To add fleet locations, specify the fleet to be updated and provide a list of
// one or more locations.
//
// If successful, this operation returns the list of added locations with their
// status set to NEW . Amazon GameLift Servers initiates the process of starting an
// instance in each added location. You can track the status of each new location
// by monitoring location creation events using [DescribeFleetEvents].
//
// # Learn more
//
// [Setting up fleets]
//
// [Update fleet locations]
//
// [Amazon GameLift Servers service locations]for managed hosting.
//
// [DescribeFleetEvents]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeFleetEvents.html
// [Amazon GameLift Servers service locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-regions.html
// [Update fleet locations]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-editing.html#fleets-update-locations
// [Setting up fleets]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-intro.html
