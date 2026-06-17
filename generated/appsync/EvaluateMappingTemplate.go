package appsync

// EvaluateMappingTemplate is generated as a reference stub.
// Executable command wiring lives under cmd/appsync.go.
//
// Evaluates a given template and returns the response. The mapping template can
// be a request or response template.
//
// Request templates take the incoming request after a GraphQL operation is parsed
// and convert it into a request configuration for the selected data source
// operation. Response templates interpret responses from the data source and map
// it to the shape of the GraphQL field output type.
//
// Mapping templates are written in the Apache Velocity Template Language (VTL).
