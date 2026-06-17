package securityir

// UpdateCaseStatus is generated as a reference stub.
// Executable command wiring lives under cmd/securityir.go.
//
// Updates the state transitions for a designated cases.
//
// Self-managed: the following states are available for self-managed cases.
//
// - Submitted → Detection and Analysis
//
// - Detection and Analysis → Containment, Eradication, and Recovery
//
// - Detection and Analysis → Post-incident Activities
//
// - Containment, Eradication, and Recovery → Detection and Analysis
//
// - Containment, Eradication, and Recovery → Post-incident Activities
//
// - Post-incident Activities → Containment, Eradication, and Recovery
//
// - Post-incident Activities → Detection and Analysis
//
// - Any → Closed
//
// AWS supported: You must use the CloseCase API to close.
