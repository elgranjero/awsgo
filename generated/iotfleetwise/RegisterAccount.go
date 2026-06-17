package iotfleetwise

// RegisterAccount is generated as a reference stub.
// Executable command wiring lives under cmd/iotfleetwise.go.
//
// This API operation contains deprecated parameters. Register your account again
// without the Timestream resources parameter so that Amazon Web Services IoT
// FleetWise can remove the Timestream metadata stored. You should then pass the
// data destination into the [CreateCampaign]API operation.
//
// You must delete any existing campaigns that include an empty data destination
// before you register your account again. For more information, see the [DeleteCampaign]API
// operation.
//
// If you want to delete the Timestream inline policy from the service-linked
// role, such as to mitigate an overly permissive policy, you must first delete any
// existing campaigns. Then delete the service-linked role and register your
// account again to enable CloudWatch metrics. For more information, see [DeleteServiceLinkedRole]in the
// Identity and Access Management API Reference.
//
// Registers your Amazon Web Services account, IAM, and Amazon Timestream
// resources so Amazon Web Services IoT FleetWise can transfer your vehicle data to
// the Amazon Web Services Cloud. For more information, including step-by-step
// procedures, see [Setting up Amazon Web Services IoT FleetWise].
//
// An Amazon Web Services account is not the same thing as a "user." An [Amazon Web Services user] is an
// identity that you create using Identity and Access Management (IAM) and takes
// the form of either an [IAM user]or an [IAM role, both with credentials]. A single Amazon Web Services account can, and
// typically does, contain many users and roles.
//
// [CreateCampaign]: https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_CreateCampaign.html
// [DeleteServiceLinkedRole]: https://docs.aws.amazon.com/IAM/latest/APIReference/API_DeleteServiceLinkedRole.html
// [Amazon Web Services user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/introduction_identity-management.html#intro-identity-users
// [Setting up Amazon Web Services IoT FleetWise]: https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/setting-up.html
// [DeleteCampaign]: https://docs.aws.amazon.com/iot-fleetwise/latest/APIReference/API_DeleteCampaign.html
// [IAM user]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users.html
// [IAM role, both with credentials]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html
