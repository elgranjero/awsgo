package resiliencehub

// CreateAppVersionResource is generated as a reference stub.
// Executable command wiring lives under cmd/resiliencehub.go.
//
// Adds a resource to the Resilience Hub application and assigns it to the
// specified Application Components. If you specify a new Application Component,
// Resilience Hub will automatically create the Application Component.
//
// - This action has no effect outside Resilience Hub.
//
// - This API updates the Resilience Hub application draft version. To use this
// resource for running resiliency assessments, you must publish the Resilience Hub
// application using the PublishAppVersion API.
//
// - To update application version with new physicalResourceID , you must call
// ResolveAppVersionResources API.
