package apprunner

// CreateObservabilityConfiguration is generated as a reference stub.
// Executable command wiring lives under cmd/apprunner.go.
//
// Create an App Runner observability configuration resource. App Runner requires
// this resource when you create or update App Runner services and you want to
// enable non-default observability features. You can share an observability
// configuration across multiple services.
//
// Create multiple revisions of a configuration by calling this action multiple
// times using the same ObservabilityConfigurationName . The call returns
// incremental ObservabilityConfigurationRevision values. When you create a
// service and configure an observability configuration resource, the service uses
// the latest active revision of the observability configuration by default. You
// can optionally configure the service to use a specific revision.
//
// The observability configuration resource is designed to configure multiple
// features (currently one feature, tracing). This action takes optional parameters
// that describe the configuration of these features (currently one parameter,
// TraceConfiguration ). If you don't specify a feature parameter, App Runner
// doesn't enable the feature.
