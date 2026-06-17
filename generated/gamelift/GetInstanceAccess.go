package gamelift

// GetInstanceAccess is generated as a reference stub.
// Executable command wiring lives under cmd/gamelift.go.
//
// This API works with the following fleet types: EC2
//
// Requests authorization to remotely connect to an instance in an Amazon GameLift
// Servers managed fleet. Use this operation to connect to instances with game
// servers that use Amazon GameLift Servers server SDK 4.x or earlier. To connect
// to instances with game servers that use server SDK 5.x or later, call [https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetComputeAccess].
//
// To request access to an instance, specify IDs for the instance and the fleet it
// belongs to. You can retrieve instance IDs for a fleet by calling [DescribeInstances]with the fleet
// ID.
//
// If successful, this operation returns an IP address and credentials. The
// returned credentials match the operating system of the instance, as follows:
//
// - For a Windows instance: returns a user name and secret (password) for use
// with a Windows Remote Desktop client.
//
// - For a Linux instance: returns a user name and secret (RSA private key) for
// use with an SSH client. You must save the secret to a .pem file. If you're
// using the CLI, see the example [Get credentials for a Linux instance]for tips on automatically saving the secret to
// a .pem file.
//
// # Learn more
//
// [Remotely connect to fleet instances]
//
// [Debug fleet issues]
//
// # Related actions
//
// [All APIs by task]
//
// [Remotely connect to fleet instances]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-remote-access.html
// [https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetComputeAccess]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetComputeAccess
// [DescribeInstances]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_DescribeInstances.html
// [Get credentials for a Linux instance]: https://docs.aws.amazon.com/gamelift/latest/apireference/API_GetInstanceAccess.html#API_GetInstanceAccess_Examples
// [Debug fleet issues]: https://docs.aws.amazon.com/gamelift/latest/developerguide/fleets-creating-debug.html
// [All APIs by task]: https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets
