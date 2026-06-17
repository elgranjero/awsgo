package dynamodb

// ExecuteStatement is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// This operation allows you to perform reads and singleton writes on data stored
// in DynamoDB, using PartiQL.
//
// For PartiQL reads ( SELECT statement), if the total number of processed items
// exceeds the maximum dataset size limit of 1 MB, the read stops and results are
// returned to the user as a LastEvaluatedKey value to continue the read in a
// subsequent operation. If the filter criteria in WHERE clause does not match any
// data, the read will return an empty result set.
//
// A single SELECT statement response can return up to the maximum number of items
// (if using the Limit parameter) or a maximum of 1 MB of data (and then apply any
// filtering to the results using WHERE clause). If LastEvaluatedKey is present in
// the response, you need to paginate the result set. If NextToken is present, you
// need to paginate the result set and include NextToken .
