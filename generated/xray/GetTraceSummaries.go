package xray

// GetTraceSummaries is generated as a reference stub.
// Executable command wiring lives under cmd/xray.go.
//
// Retrieves IDs and annotations for traces available for a specified time frame
// using an optional filter. To get the full traces, pass the trace IDs to
// BatchGetTraces .
//
// A filter expression can target traced requests that hit specific service nodes
// or edges, have errors, or come from a known user. For example, the following
// filter expression targets traces that pass through api.example.com :
//
// service("api.example.com")
//
// This filter expression finds traces that have an annotation named account with
// the value 12345 :
//
// annotation.account = "12345"
//
// For a full list of indexed fields and keywords that you can use in filter
// expressions, see [Use filter expressions]in the Amazon Web Services X-Ray Developer Guide.
//
// [Use filter expressions]: https://docs.aws.amazon.com/xray/latest/devguide/aws-xray-interface-console.html#xray-console-filters
