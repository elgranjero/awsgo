package textract

// GetDocumentTextDetection is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Gets the results for an Amazon Textract asynchronous operation that detects
// text in a document. Amazon Textract can detect lines of text and the words that
// make up a line of text.
//
// You start asynchronous text detection by calling StartDocumentTextDetection, which returns a job
// identifier ( JobId ). When the text detection operation finishes, Amazon
// Textract publishes a completion status to the Amazon Simple Notification Service
// (Amazon SNS) topic that's registered in the initial call to
// StartDocumentTextDetection . To get the results of the text-detection operation,
// first check that the status value published to the Amazon SNS topic is SUCCEEDED
// . If so, call GetDocumentTextDetection , and pass the job identifier ( JobId )
// from the initial call to StartDocumentTextDetection .
//
// GetDocumentTextDetection returns an array of Block objects.
//
// Each document page has as an associated Block of type PAGE. Each PAGE Block
// object is the parent of LINE Block objects that represent the lines of detected
// text on a page. A LINE Block object is a parent for each word that makes up the
// line. Words are represented by Block objects of type WORD.
//
// Use the MaxResults parameter to limit the number of blocks that are returned.
// If there are more results than specified in MaxResults , the value of NextToken
// in the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetDocumentTextDetection , and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetDocumentTextDetection .
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
