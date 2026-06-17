package codedeploy

// ContinueDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/codedeploy.go.
//
// For a blue/green deployment, starts the process of rerouting traffic from
// instances in the original environment to instances in the replacement
// environment without waiting for a specified wait time to elapse. (Traffic
// rerouting, which is achieved by registering instances in the replacement
// environment with the load balancer, can start as soon as all instances have a
// status of Ready.)
