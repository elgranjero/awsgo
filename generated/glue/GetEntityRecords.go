package glue

// GetEntityRecords is generated as a reference stub.
// Executable command wiring lives under cmd/glue.go.
//
// This API is used to query preview data from a given connection type or from a
// native Amazon S3 based Glue Data Catalog.
//
// Returns records as an array of JSON blobs. Each record is formatted using
// Jackson JsonNode based on the field type defined by the DescribeEntity API.
//
// Spark connectors generate schemas according to the same data type mapping as in
// the DescribeEntity API. Spark connectors convert data to the appropriate data
// types matching the schema when returning rows.
