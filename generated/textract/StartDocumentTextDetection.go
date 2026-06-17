package textract

// StartDocumentTextDetection is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Starts the asynchronous detection of text in a document. Amazon Textract can
// detect lines of text and the words that make up a line of text.
//
// StartDocumentTextDetection can analyze text in documents that are in JPEG, PNG,
// TIFF, and PDF format. The documents are stored in an Amazon S3 bucket. Use DocumentLocationto
// specify the bucket name and file name of the document.
//
// StartDocumentTextDetection returns a job identifier ( JobId ) that you use to
// get the results of the operation. When text detection is finished, Amazon
// Textract publishes a completion status to the Amazon Simple Notification Service
// (Amazon SNS) topic that you specify in NotificationChannel . To get the results
// of the text detection operation, first check that the status value published to
// the Amazon SNS topic is SUCCEEDED . If so, call GetDocumentTextDetection, and pass the job identifier (
// JobId ) from the initial call to StartDocumentTextDetection .
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
