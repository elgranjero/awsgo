package translate

// StartTextTranslationJob is generated as a reference stub.
// Executable command wiring lives under cmd/translate.go.
//
// Starts an asynchronous batch translation job. Use batch translation jobs to
// translate large volumes of text across multiple documents at once. For batch
// translation, you can input documents with different source languages (specify
// auto as the source language). You can specify one or more target languages.
// Batch translation translates each input document into each of the target
// languages. For more information, see [Asynchronous batch processing].
//
// Batch translation jobs can be described with the DescribeTextTranslationJob operation, listed with the ListTextTranslationJobs
// operation, and stopped with the StopTextTranslationJoboperation.
//
// [Asynchronous batch processing]: https://docs.aws.amazon.com/translate/latest/dg/async.html
