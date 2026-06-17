package comprehend

// ClassifyDocument is generated as a reference stub.
// Executable command wiring lives under cmd/comprehend.go.
//
// Creates a classification request to analyze a single document in real-time.
// ClassifyDocument supports the following model types:
//
// - Custom classifier - a custom model that you have created and trained. For
// input, you can provide plain text, a single-page document (PDF, Word, or image),
// or Amazon Textract API output. For more information, see [Custom classification]in the Amazon
// Comprehend Developer Guide.
//
// - Prompt safety classifier - Amazon Comprehend provides a pre-trained model
// for classifying input prompts for generative AI applications. For input, you
// provide English plain text input. For prompt safety classification, the response
// includes only the Classes field. For more information about prompt safety
// classifiers, see [Prompt safety classification]in the Amazon Comprehend Developer Guide.
//
// If the system detects errors while processing a page in the input document, the
// API response includes an Errors field that describes the errors.
//
// If the system detects a document-level error in your input document, the API
// returns an InvalidRequestException error response. For details about this
// exception, see [Errors in semi-structured documents]in the Comprehend Developer Guide.
//
// [Custom classification]: https://docs.aws.amazon.com/comprehend/latest/dg/how-document-classification.html
// [Prompt safety classification]: https://docs.aws.amazon.com/comprehend/latest/dg/trust-safety.html#prompt-classification
// [Errors in semi-structured documents]: https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync-err.html
