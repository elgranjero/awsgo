package apprunner

// CreateAutoScalingConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Create an App Runner automatic scaling configuration resource. App Runner
// requires this resource when you create or update App Runner services and you
// require non-default auto scaling settings. You can share an auto scaling
// configuration across multiple services.
//
// Create multiple revisions of a configuration by calling this action multiple
// times using the same AutoScalingConfigurationName . The call returns incremental
// AutoScalingConfigurationRevision values. When you create a service and configure
// an auto scaling configuration resource, the service uses the latest active
// revision of the auto scaling configuration by default. You can optionally
// configure the service to use a specific revision.
//
// Configure a higher MinSize to increase the spread of your App Runner service
// over more Availability Zones in the Amazon Web Services Region. The tradeoff is
// a higher minimal cost.
//
// Configure a lower MaxSize to control your cost. The tradeoff is lower
// responsiveness during peak demand.
