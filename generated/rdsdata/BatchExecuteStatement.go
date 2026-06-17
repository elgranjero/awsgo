package rdsdata

// BatchExecuteStatement is generated as a reference stub.
// Executable command wiring lives under cmd/rdsdata.go.
//
// Runs a batch SQL statement over an array of data.
//
// You can run bulk update and insert operations for multiple records using a DML
// statement with different parameter sets. Bulk operations can provide a
// significant performance improvement over individual insert and update
// operations.
//
// If a call isn't part of a transaction because it doesn't include the
// transactionID parameter, changes that result from the call are committed
// automatically.
//
// There isn't a fixed upper limit on the number of parameter sets. However, the
// maximum size of the HTTP request submitted through the Data API is 4 MiB. If the
// request exceeds this limit, the Data API returns an error and doesn't process
// the request. This 4-MiB limit includes the size of the HTTP headers and the JSON
// notation in the request. Thus, the number of parameter sets that you can include
// depends on a combination of factors, such as the size of the SQL statement and
// the size of each parameter set.
//
// The response size limit is 1 MiB. If the call returns more than 1 MiB of
// response data, the call is terminated.
