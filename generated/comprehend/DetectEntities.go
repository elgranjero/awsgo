package comprehend

// DetectEntities is generated as a reference stub.
// Executable command wiring lives under cmd/comprehend.go.
//
// Detects named entities in input text when you use the pre-trained model.
// Detects custom entities if you have a custom entity recognition model.
//
// When detecting named entities using the pre-trained model, use plain text as
// the input. For more information about named entities, see [Entities]in the Comprehend
// Developer Guide.
//
// When you use a custom entity recognition model, you can input plain text or you
// can upload a single-page input document (text, PDF, Word, or image).
//
// If the system detects errors while processing a page in the input document, the
// API response includes an entry in Errors for each error.
//
// If the system detects a document-level error in your input document, the API
// returns an InvalidRequestException error response. For details about this
// exception, see [Errors in semi-structured documents]in the Comprehend Developer Guide.
//
// [Errors in semi-structured documents]: https://docs.aws.amazon.com/comprehend/latest/dg/idp-inputs-sync-err.html
// [Entities]: https://docs.aws.amazon.com/comprehend/latest/dg/how-entities.html
