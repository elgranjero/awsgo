package textract

// GetExpenseAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Gets the results for an Amazon Textract asynchronous operation that analyzes
// invoices and receipts. Amazon Textract finds contact information, items
// purchased, and vendor name, from input invoices and receipts.
//
// You start asynchronous invoice/receipt analysis by calling StartExpenseAnalysis, which returns a
// job identifier ( JobId ). Upon completion of the invoice/receipt analysis,
// Amazon Textract publishes the completion status to the Amazon Simple
// Notification Service (Amazon SNS) topic. This topic must be registered in the
// initial call to StartExpenseAnalysis . To get the results of the invoice/receipt
// analysis operation, first ensure that the status value published to the Amazon
// SNS topic is SUCCEEDED . If so, call GetExpenseAnalysis , and pass the job
// identifier ( JobId ) from the initial call to StartExpenseAnalysis .
//
// Use the MaxResults parameter to limit the number of blocks that are returned.
// If there are more results than specified in MaxResults , the value of NextToken
// in the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetExpenseAnalysis , and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetExpenseAnalysis .
//
// For more information, see [Analyzing Invoices and Receipts].
//
// [Analyzing Invoices and Receipts]: https://docs.aws.amazon.com/textract/latest/dg/invoices-receipts.html
