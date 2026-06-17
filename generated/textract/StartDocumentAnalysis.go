package textract

// StartDocumentAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Starts the asynchronous analysis of an input document for relationships between
// detected items such as key-value pairs, tables, and selection elements.
//
// StartDocumentAnalysis can analyze text in documents that are in JPEG, PNG,
// TIFF, and PDF format. The documents are stored in an Amazon S3 bucket. Use DocumentLocationto
// specify the bucket name and file name of the document.
//
// StartDocumentAnalysis returns a job identifier ( JobId ) that you use to get the
// results of the operation. When text analysis is finished, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that you specify in NotificationChannel . To get the results of the
// text analysis operation, first check that the status value published to the
// Amazon SNS topic is SUCCEEDED . If so, call GetDocumentAnalysis, and pass the job identifier ( JobId
// ) from the initial call to StartDocumentAnalysis .
//
// For more information, see [Document Text Analysis].
//
// [Document Text Analysis]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html
