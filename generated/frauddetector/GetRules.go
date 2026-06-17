package frauddetector

// GetRules is generated as a reference stub.
// Executable command wiring lives under cmd/frauddetector.go.
//
// Get all rules for a detector (paginated) if ruleId and ruleVersion are not
// specified. Gets all rules for the detector and the ruleId if present
// (paginated). Gets a specific rule if both the ruleId and the ruleVersion are
// specified.
//
// This is a paginated API. Providing null maxResults results in retrieving
// maximum of 100 records per page. If you provide maxResults the value must be
// between 50 and 100. To get the next page result, a provide a pagination token
// from GetRulesResult as part of your request. Null pagination token fetches the
// records from the beginning.
