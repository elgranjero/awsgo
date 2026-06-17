package dynamodb

// BatchWriteItem is generated as a reference stub.
// Executable command wiring lives under cmd/dynamodb.go.
//
// The BatchWriteItem operation puts or deletes multiple items in one or more
// tables. A single call to BatchWriteItem can transmit up to 16MB of data over
// the network, consisting of up to 25 item put or delete operations. While
// individual items can be up to 400 KB once stored, it's important to note that an
// item's representation might be greater than 400KB while being sent in DynamoDB's
// JSON format for the API call. For more details on this distinction, see [Naming Rules and Data Types].
//
// BatchWriteItem cannot update items. If you perform a BatchWriteItem operation
// on an existing item, that item's values will be overwritten by the operation and
// it will appear like it was updated. To update items, we recommend you use the
// UpdateItem action.
//
// The individual PutItem and DeleteItem operations specified in BatchWriteItem
// are atomic; however BatchWriteItem as a whole is not. If any requested
// operations fail because the table's provisioned throughput is exceeded or an
// internal processing failure occurs, the failed operations are returned in the
// UnprocessedItems response parameter. You can investigate and optionally resend
// the requests. Typically, you would call BatchWriteItem in a loop. Each
// iteration would check for unprocessed items and submit a new BatchWriteItem
// request with those unprocessed items until all items have been processed.
//
// For tables and indexes with provisioned capacity, if none of the items can be
// processed due to insufficient provisioned throughput on all of the tables in the
// request, then BatchWriteItem returns a ProvisionedThroughputExceededException .
// For all tables and indexes, if none of the items can be processed due to other
// throttling scenarios (such as exceeding partition level limits), then
// BatchWriteItem returns a ThrottlingException .
//
// If DynamoDB returns any unprocessed items, you should retry the batch operation
// on those items. However, we strongly recommend that you use an exponential
// backoff algorithm. If you retry the batch operation immediately, the underlying
// read or write requests can still fail due to throttling on the individual
// tables. If you delay the batch operation using exponential backoff, the
// individual requests in the batch are much more likely to succeed.
//
// For more information, see [Batch Operations and Error Handling] in the Amazon DynamoDB Developer Guide.
//
// With BatchWriteItem , you can efficiently write or delete large amounts of data,
// such as from Amazon EMR, or copy data from another database into DynamoDB. In
// order to improve performance with these large-scale operations, BatchWriteItem
// does not behave in the same way as individual PutItem and DeleteItem calls
// would. For example, you cannot specify conditions on individual put and delete
// requests, and BatchWriteItem does not return deleted items in the response.
//
// If you use a programming language that supports concurrency, you can use
// threads to write items in parallel. Your application must include the necessary
// logic to manage the threads. With languages that don't support threading, you
// must update or delete the specified items one at a time. In both situations,
// BatchWriteItem performs the specified put and delete operations in parallel,
// giving you the power of the thread pool approach without having to introduce
// complexity into your application.
//
// Parallel processing reduces latency, but each specified put and delete request
// consumes the same number of write capacity units whether it is processed in
// parallel or not. Delete operations on nonexistent items consume one write
// capacity unit.
//
// If one or more of the following is true, DynamoDB rejects the entire batch
// write operation:
//
// - One or more tables specified in the BatchWriteItem request does not exist.
//
// - Primary key attributes specified on an item in the request do not match
// those in the corresponding table's primary key schema.
//
// - You try to perform multiple operations on the same item in the same
// BatchWriteItem request. For example, you cannot put and delete the same item
// in the same BatchWriteItem request.
//
// - Your request contains at least two items with identical hash and range keys
// (which essentially is two put operations).
//
// - There are more than 25 requests in the batch.
//
// - Any individual item in a batch exceeds 400 KB.
//
// - The total request size exceeds 16 MB.
//
// - Any individual items with keys exceeding the key length limits. For a
// partition key, the limit is 2048 bytes and for a sort key, the limit is 1024
// bytes.
//
// [Batch Operations and Error Handling]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#Programming.Errors.BatchOperations
// [Naming Rules and Data Types]: https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html
