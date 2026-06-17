package route53recoverycontrolconfig

// CreateSafetyRule is generated as a reference stub.
// Executable command wiring lives under cmd/route53recoverycontrolconfig.go.
//
// Creates a safety rule in a control panel. Safety rules let you add safeguards
// around changing routing control states, and for enabling and disabling routing
// controls, to help prevent unexpected outcomes.
//
// There are two types of safety rules: assertion rules and gating rules.
//
// Assertion rule: An assertion rule enforces that, when you change a routing
// control state, that a certain criteria is met. For example, the criteria might
// be that at least one routing control state is On after the transaction so that
// traffic continues to flow to at least one cell for the application. This ensures
// that you avoid a fail-open scenario.
//
// Gating rule: A gating rule lets you configure a gating routing control as an
// overall "on/off" switch for a group of routing controls. Or, you can configure
// more complex gating scenarios, for example by configuring multiple gating
// routing controls.
//
// For more information, see [Safety rules] in the Amazon Route 53 Application Recovery
// Controller Developer Guide.
//
// [Safety rules]: https://docs.aws.amazon.com/r53recovery/latest/dg/routing-control.safety-rules.html
