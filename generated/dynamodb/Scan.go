package dynamodb

// Scan is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// The Scan operation returns one or more items and item attributes by accessing
// every item in a table or a secondary index. To have DynamoDB return fewer items,
// you can provide a FilterExpression operation.
//
// If the total size of scanned items exceeds the maximum dataset size limit of 1
// MB, the scan completes and results are returned to the user. The
// LastEvaluatedKey value is also returned and the requestor can use the
// LastEvaluatedKey to continue the scan in a subsequent operation. Each scan
// response also includes number of items that were scanned (ScannedCount) as part
// of the request. If using a FilterExpression , a scan result can result in no
// items meeting the criteria and the Count will result in zero. If you did not
// use a FilterExpression in the scan request, then Count is the same as
// ScannedCount .
//
// Count and ScannedCount only return the count of items specific to a single scan
// request and, unless the table is less than 1MB, do not represent the total
// number of items in the table.
//
// A single Scan operation first reads up to the maximum number of items set (if
// using the Limit parameter) or a maximum of 1 MB of data and then applies any
// filtering to the results if a FilterExpression is provided. If LastEvaluatedKey
// is present in the response, pagination is required to complete the full table
// scan. For more information, see [Paginating the Results]in the Amazon DynamoDB Developer Guide.
//
// Scan operations proceed sequentially; however, for faster performance on a
// large table or secondary index, applications can request a parallel Scan
// operation by providing the Segment and TotalSegments parameters. For more
// information, see [Parallel Scan]in the Amazon DynamoDB Developer Guide.
//
// By default, a Scan uses eventually consistent reads when accessing the items in
// a table. Therefore, the results from an eventually consistent Scan may not
// include the latest item changes at the time the scan iterates through each item
// in the table. If you require a strongly consistent read of each item as the scan
// iterates through the items in the table, you can set the ConsistentRead
// parameter to true. Strong consistency only relates to the consistency of the
// read at the item level.
//
// DynamoDB does not provide snapshot isolation for a scan operation when the
// ConsistentRead parameter is set to true. Thus, a DynamoDB scan operation does
// not guarantee that all reads in a scan see a consistent snapshot of the table
// when the scan operation was requested.
//
// [Paginating the Results]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Scan.html#Scan.Pagination
// [Parallel Scan]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Scan.html#Scan.ParallelScan
