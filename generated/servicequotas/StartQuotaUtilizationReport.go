package servicequotas

// StartQuotaUtilizationReport is generated as a reference stub.
// Executable command wiring lives under cmd/servicequotas.go.
//
// Initiates the generation of a quota utilization report for your Amazon Web
// Services account. This asynchronous operation analyzes your quota usage across
// all Amazon Web Services services and returns a unique report identifier that you
// can use to retrieve the results.
//
// The report generation process may take several seconds to complete, depending
// on the number of quotas in your account. Use the GetQuotaUtilizationReport
// operation to check the status and retrieve the results when the report is ready.
