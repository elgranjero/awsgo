package textract

// AnalyzeDocument is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Analyzes an input document for relationships between detected items.
//
// The types of information returned are as follows:
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
// (including text that doesn't have a relationship with the value of
// FeatureTypes ).
//
// - Signatures. A SIGNATURE Block object contains the location information of a
// signature in a document. If used in conjunction with forms or tables, a
// signature can be given a Key-Value pairing or be detected in the cell of a
// table.
//
// - Query. A QUERY Block object contains the query text, alias and link to the
// associated Query results block object.
//
// - Query Result. A QUERY_RESULT Block object contains the answer to the query
// and an ID that connects it to the query asked. This Block also contains a
// confidence score.
//
// Selection elements such as check boxes and option buttons (radio buttons) can
// be detected in form data and in tables. A SELECTION_ELEMENT Block object
// contains information about a selection element, including the selection status.
//
// You can choose which type of analysis to perform by specifying the FeatureTypes
// list.
//
// The output is returned in a list of Block objects.
//
// AnalyzeDocument is a synchronous operation. To analyze documents
// asynchronously, use StartDocumentAnalysis.
//
// For more information, see [Document Text Analysis].
//
// [Document Text Analysis]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-analyzing.html
