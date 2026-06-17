package elasticbeanstalk

// UpdateEnvironment is generated as a reference stub.
// Executable command wiring lives under cmd/elasticbeanstalk.go.
//
// Updates the environment description, deploys a new application version, updates
// the configuration settings to an entirely new configuration template, or updates
// select configuration option values in the running environment.
//
// Attempting to update both the release and configuration is not allowed and AWS
// Elastic Beanstalk returns an InvalidParameterCombination error.
//
// When updating the configuration settings to a new template or individual
// settings, a draft configuration is created and DescribeConfigurationSettingsfor this environment returns two
// setting descriptions with different DeploymentStatus values.
