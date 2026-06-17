package configservice

// PutResourceConfig is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Records the configuration state for the resource provided in the request.
//
// The configuration state of a resource is represented in Config as Configuration
// Items. Once this API records the configuration item, you can retrieve the list
// of configuration items for the custom resource type using existing Config APIs.
//
// The custom resource type must be registered with CloudFormation. This API
// accepts the configuration item registered with CloudFormation.
//
// When you call this API, Config only stores configuration state of the resource
// provided in the request. This API does not change or remediate the configuration
// of the resource.
//
// Write-only schema properites are not recorded as part of the published
// configuration item.
