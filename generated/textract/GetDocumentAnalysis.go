package textract

// GetDocumentAnalysis is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Gets the results for an Amazon Textract asynchronous operation that analyzes
// text in a document.
//
// You start asynchronous text analysis by calling StartDocumentAnalysis, which returns a job
// identifier ( JobId ). When the text analysis operation finishes, Amazon Textract
// publishes a completion status to the Amazon Simple Notification Service (Amazon
// SNS) topic that's registered in the initial call to StartDocumentAnalysis . To
// get the results of the text-detection operation, first check that the status
// value published to the Amazon SNS topic is SUCCEEDED . If so, call
// GetDocumentAnalysis , and pass the job identifier ( JobId ) from the initial
// call to StartDocumentAnalysis .
//
// GetDocumentAnalysis returns an array of Block objects. The following types of
// information are returned:
//
// - Form data (key-value pairs). The related information is returned in two Block
// objects, each of type KEY_VALUE_SET : a KEY Block object and a VALUE Block
// object. For example, Name: Ana Silva Carolina contains a key and value. Name: is
// the key. Ana Silva Carolina is the value.
//
// - Table and table cell data. A TABLE Block object contains information about a
// detected table. A CELL Block object is returned for each cell in a table.
//
// - Lines and words of text. A LINE Block object contains one or more WORD Block
// objects. All lines and words that are detected in the document are returned
// (including text that doesn't have a relationship with the value of the
// StartDocumentAnalysis FeatureTypes input parameter).
//
// - Query. A QUERY Block object contains the query text, alias and link to the
// associated Query results block object.
//
// - Query Results. A QUERY_RESULT Block object contains the answer to the query
// and an ID that connects it to the query asked. This Block also contains a
// confidence score.
//
// While processing a document with queries, look out for
// INVALID_REQUEST_PARAMETERS output. This indicates that either the per page query
// limit has been exceeded or that the operation is trying to query a page in the
// document which doesn’t exist.
//
// Selection elements such as check boxes and option buttons (radio buttons) can
// be detected in form data and in tables. A SELECTION_ELEMENT Block object
// contains information about a selection element, including the selection status.
//
// Use the MaxResults parameter to limit the number of blocks that are returned.
// If there are more results than specified in MaxResults , the value of NextToken
// in the operation response contains a pagination token for getting the next set
// of results. To get the next page of results, call GetDocumentAnalysis , and
// populate the NextToken request parameter with the token value that's returned
// from the previous call to GetDocumentAnalysis .
//
// For more information, see [Document Text Analysis].
//
// [Document Text Analysis]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html
