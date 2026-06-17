package elasticbeanstalk

// DeleteEnvironmentConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/elasticbeanstalk.go.
//
// Deletes the draft configuration associated with the running environment.
//
// Updating a running environment with any configuration changes creates a draft
// configuration set. You can get the draft configuration using DescribeConfigurationSettingswhile the update
// is in progress or if the update fails. The DeploymentStatus for the draft
// configuration indicates whether the deployment is in process or has failed. The
// draft configuration remains in existence until it is deleted with this action.
