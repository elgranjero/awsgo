package codedeploy

// BatchGetDeploymentTargets is generated as a reference stub.
// Executable command wiring lives under cmd/codedeploy.go.
//
// Returns an array of one or more targets associated with a deployment. This
//
// method works with all compute types and should be used instead of the deprecated
// BatchGetDeploymentInstances . The maximum number of targets that can be returned
// is 25.
//
// The type of targets returned depends on the deployment's compute platform or
// deployment method:
//
// - EC2/On-premises: Information about Amazon EC2 instance targets.
//
// - Lambda: Information about Lambda functions targets.
//
// - Amazon ECS: Information about Amazon ECS service targets.
//
// - CloudFormation: Information about targets of blue/green deployments
// initiated by a CloudFormation stack update.
