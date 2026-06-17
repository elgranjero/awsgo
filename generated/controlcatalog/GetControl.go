package controlcatalog

// GetControl is generated as a reference stub.
// Executable command wiring lives under cmd/controlcatalog.go.
//
// Returns details about a specific control, most notably a list of Amazon Web
// Services Regions where this control is supported. Input a value for the
// ControlArn parameter, in ARN form. GetControl accepts controltower or
// controlcatalog control ARNs as input. Returns a controlcatalog ARN format.
//
// In the API response, controls that have the value GLOBAL in the Scope field do
// not show the DeployableRegions field, because it does not apply. Controls that
// have the value REGIONAL in the Scope field return a value for the
// DeployableRegions field, as shown in the example.
