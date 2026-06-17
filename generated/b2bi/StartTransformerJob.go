package b2bi

// StartTransformerJob is generated as a reference stub.
// Executable command wiring lives under cmd/b2bi.go.
//
// Runs a job, using a transformer, to parse input EDI (electronic data
// interchange) file into the output structures used by Amazon Web Services B2B
// Data Interchange.
//
// If you only want to transform EDI (electronic data interchange) documents, you
// don't need to create profiles, partnerships or capabilities. Just create and
// configure a transformer, and then run the StartTransformerJob API to process
// your files.
//
// The system stores transformer jobs for 30 days. During that period, you can run [GetTransformerJob]
// and supply its transformerId and transformerJobId to return details of the job.
//
// [GetTransformerJob]: https://docs.aws.amazon.com/b2bi/latest/APIReference/API_GetTransformerJob.html
