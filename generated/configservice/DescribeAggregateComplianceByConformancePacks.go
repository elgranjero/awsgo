package configservice

// DescribeAggregateComplianceByConformancePacks is generated as a reference stub.
// Executable command wiring lives under cmd/configservice.go.
//
// Returns a list of the existing and deleted conformance packs and their
// associated compliance status with the count of compliant and noncompliant Config
// rules within each conformance pack. Also returns the total rule count which
// includes compliant rules, noncompliant rules, and rules that cannot be evaluated
// due to insufficient data.
//
// The results can return an empty result page, but if you have a nextToken , the
// results are displayed on the next page.
