package textract

// StartExpenseAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Starts the asynchronous analysis of invoices or receipts for data like contact
// information, items purchased, and vendor names.
//
// StartExpenseAnalysis can analyze text in documents that are in JPEG, PNG, and
// PDF format. The documents must be stored in an Amazon S3 bucket. Use the DocumentLocation
// parameter to specify the name of your S3 bucket and the name of the document in
// that bucket.
//
// StartExpenseAnalysis returns a job identifier ( JobId ) that you will provide to
// GetExpenseAnalysis to retrieve the results of the operation. When the analysis
// of the input invoices/receipts is finished, Amazon Textract publishes a
// completion status to the Amazon Simple Notification Service (Amazon SNS) topic
// that you provide to the NotificationChannel . To obtain the results of the
// invoice and receipt analysis operation, ensure that the status value published
// to the Amazon SNS topic is SUCCEEDED . If so, call GetExpenseAnalysis, and pass the job
// identifier ( JobId ) that was returned by your call to StartExpenseAnalysis .
//
// For more information, see [Analyzing Invoices and Receipts].
//
// [Analyzing Invoices and Receipts]: https://docs.aws.amazon.com/textract/latest/dg/invoice-receipts.html
