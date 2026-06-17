package rds

// CreateBlueGreenDeployment is generated as a reference stub.
// Executable command wiring lives under cmd/rds.go.
//
// Creates a blue/green deployment.
//
// A blue/green deployment creates a staging environment that copies the
// production environment. In a blue/green deployment, the blue environment is the
// current production environment. The green environment is the staging
// environment, and it stays in sync with the current production environment.
//
// You can make changes to the databases in the green environment without
// affecting production workloads. For example, you can upgrade the major or minor
// DB engine version, change database parameters, or make schema changes in the
// staging environment. You can thoroughly test changes in the green environment.
// When ready, you can switch over the environments to promote the green
// environment to be the new production environment. The switchover typically takes
// under a minute.
//
// For more information, see [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon RDS User Guide and [Using Amazon RDS Blue/Green Deployments for database updates] in the Amazon
// Aurora User Guide.
//
// [Using Amazon RDS Blue/Green Deployments for database updates]: https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/blue-green-deployments.html
