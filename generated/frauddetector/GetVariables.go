package frauddetector

// GetVariables is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Gets all of the variables or the specific variable. This is a paginated API.
// Providing null maxSizePerPage results in retrieving maximum of 100 records per
// page. If you provide maxSizePerPage the value must be between 50 and 100. To
// get the next page result, a provide a pagination token from GetVariablesResult
// as part of your request. Null pagination token fetches the records from the
// beginning.
