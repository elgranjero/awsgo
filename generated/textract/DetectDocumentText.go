package textract

// DetectDocumentText is generated as a reference stub.
// Executable command wiring lives under cmd/textract.go.
//
// Detects text in the input document. Amazon Textract can detect lines of text
// and the words that make up a line of text. The input document must be in one of
// the following image formats: JPEG, PNG, PDF, or TIFF. DetectDocumentText
// returns the detected text in an array of Blockobjects.
//
// Each document page has as an associated Block of type PAGE. Each PAGE Block
// object is the parent of LINE Block objects that represent the lines of detected
// text on a page. A LINE Block object is a parent for each word that makes up the
// line. Words are represented by Block objects of type WORD.
//
// DetectDocumentText is a synchronous operation. To analyze documents
// asynchronously, use StartDocumentTextDetection.
//
// For more information, see [Document Text Detection].
//
// [Document Text Detection]: https://docs.aws.amazon.com/textract/latest/dg/how-it-works-detecting.html
